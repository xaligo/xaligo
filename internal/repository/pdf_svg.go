//go:build !js

package repository

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"image"
	stdcolor "image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-fonts/latin-modern/lmsans10bold"
	"github.com/go-fonts/latin-modern/lmsans10boldoblique"
	"github.com/go-fonts/latin-modern/lmsans10oblique"
	"github.com/go-fonts/latin-modern/lmsans10regular"
	"github.com/tdewolff/canvas"
)

const (
	pdfSVGPointsPerMillimeter = 72.0 / 25.4
	pdfSVGMaxImageDepth       = 8
)

var pdfSVGNumberPattern = regexp.MustCompile(`[-+]?(?:\d+(?:\.\d*)?|\.\d+)(?:[eE][-+]?\d+)?`)

// pdfSVGComposer complements canvas.ParseSVG with the two SVG features used by
// Xaligo output that the canvas parser currently omits: tspan text and images.
// Fonts are loaded from embedded bytes so PDF export never depends on fonts
// installed on the host running Xaligo.
type pdfSVGComposer struct {
	fontFamily *canvas.FontFamily
}

func newPDFSVGComposer() (*pdfSVGComposer, error) {
	family := canvas.NewFontFamily("Xaligo PDF Sans")
	fonts := []struct {
		data  []byte
		style canvas.FontStyle
	}{
		{data: lmsans10regular.TTF, style: canvas.FontRegular},
		{data: lmsans10bold.TTF, style: canvas.FontBold},
		{data: lmsans10oblique.TTF, style: canvas.FontRegular | canvas.FontItalic},
		{data: lmsans10boldoblique.TTF, style: canvas.FontBold | canvas.FontItalic},
	}
	for _, font := range fonts {
		if err := family.LoadFont(font.data, 0, font.style); err != nil {
			return nil, fmt.Errorf("load embedded PDF font %s: %w", font.style, err)
		}
	}
	return &pdfSVGComposer{fontFamily: family}, nil
}

func (composer *pdfSVGComposer) parse(data []byte) (*canvas.Canvas, error) {
	return composer.parseDepth(data, 0)
}

func (composer *pdfSVGComposer) parseDepth(data []byte, depth int) (*canvas.Canvas, error) {
	if depth > pdfSVGMaxImageDepth {
		return nil, fmt.Errorf("embedded SVG image nesting exceeds %d levels", pdfSVGMaxImageDepth)
	}

	root, err := parsePDFSVGTree(data)
	if err != nil {
		return nil, err
	}
	sanitized, err := pdfSVGWithoutOverlays(data)
	if err != nil {
		return nil, fmt.Errorf("sanitize SVG for PDF: %w", err)
	}
	base, err := parseCanvasSVGSafely(sanitized)
	if err != nil {
		return nil, err
	}
	width, height := base.Size()
	if !positiveFinitePDFDimension(width) || !positiveFinitePDFDimension(height) {
		return nil, fmt.Errorf("parsed SVG dimensions must be positive and finite")
	}

	result := canvas.New(width, height)
	base.RenderTo(result)
	view, err := pdfSVGRootView(root, width, height)
	if err != nil {
		return nil, err
	}
	state := pdfSVGState{
		fill:        "#000000",
		fillOpacity: 1,
		opacity:     1,
		fontSize:    16,
		fontStyle:   canvas.FontRegular,
		textAnchor:  "start",
		visible:     true,
	}
	clips := collectPDFSVGClips(root)
	if err := composer.renderChildren(result, root, state, canvas.Identity, view, clips, depth); err != nil {
		return nil, err
	}
	return result, nil
}

func parseCanvasSVGSafely(data []byte) (drawing *canvas.Canvas, err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("parse SVG: %v", recovered)
			drawing = nil
		}
	}()
	drawing, err = canvas.ParseSVG(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("parse SVG: %w", err)
	}
	return drawing, nil
}

type pdfSVGNode struct {
	name    xml.Name
	attrs   map[string]string
	content []pdfSVGContent
}

type pdfSVGContent struct {
	text string
	node *pdfSVGNode
}

func parsePDFSVGTree(data []byte) (*pdfSVGNode, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return nil, fmt.Errorf("parse SVG: expected SVG tag")
		}
		if err != nil {
			return nil, fmt.Errorf("parse SVG: %w", err)
		}
		start, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		if strings.ToLower(start.Name.Local) != "svg" {
			return nil, fmt.Errorf("parse SVG: expected SVG tag, got <%s>", start.Name.Local)
		}
		return parsePDFSVGNode(decoder, start)
	}
}

func parsePDFSVGNode(decoder *xml.Decoder, start xml.StartElement) (*pdfSVGNode, error) {
	node := &pdfSVGNode{name: start.Name, attrs: make(map[string]string, len(start.Attr))}
	for _, attr := range start.Attr {
		node.attrs[attr.Name.Local] = attr.Value
	}
	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("parse SVG <%s>: %w", start.Name.Local, err)
		}
		switch value := token.(type) {
		case xml.StartElement:
			child, err := parsePDFSVGNode(decoder, value)
			if err != nil {
				return nil, err
			}
			node.content = append(node.content, pdfSVGContent{node: child})
		case xml.CharData:
			node.content = append(node.content, pdfSVGContent{text: string(value)})
		case xml.EndElement:
			return node, nil
		}
	}
}

// pdfSVGWithoutOverlays removes all text and image subtrees before handing the
// document to canvas.ParseSVG. Besides preventing duplicate output, this avoids
// canvas.ParseSVG's system-font lookup (and its panic on a missing font).
func pdfSVGWithoutOverlays(data []byte) ([]byte, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	var output bytes.Buffer
	encoder := xml.NewEncoder(&output)
	skipDepth := 0
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch value := token.(type) {
		case xml.StartElement:
			if skipDepth > 0 {
				skipDepth++
				continue
			}
			name := strings.ToLower(value.Name.Local)
			if name == "text" || name == "image" {
				skipDepth = 1
				continue
			}
		case xml.EndElement:
			if skipDepth > 0 {
				skipDepth--
				continue
			}
		default:
			if skipDepth > 0 {
				continue
			}
		}
		if err := encoder.EncodeToken(token); err != nil {
			return nil, err
		}
	}
	if err := encoder.Flush(); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

type pdfSVGState struct {
	fill        string
	fillOpacity float64
	opacity     float64
	fontSize    float64
	fontStyle   canvas.FontStyle
	textAnchor  string
	visible     bool
}

type pdfSVGRect struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func (composer *pdfSVGComposer) renderChildren(
	target *canvas.Canvas,
	parent *pdfSVGNode,
	state pdfSVGState,
	transform canvas.Matrix,
	rootView canvas.Matrix,
	clips map[string]pdfSVGRect,
	depth int,
) error {
	for _, content := range parent.content {
		if content.node == nil {
			continue
		}
		child := content.node
		name := strings.ToLower(child.name.Local)
		if name == "defs" {
			continue
		}
		childState := applyPDFSVGState(state, child.attrs)
		childTransform := transform
		if value := strings.TrimSpace(child.attrs["transform"]); value != "" {
			matrix, err := parsePDFSVGTransform(value)
			if err != nil {
				return fmt.Errorf("render PDF SVG <%s> transform: %w", child.name.Local, err)
			}
			childTransform = childTransform.Mul(matrix)
		}
		if !childState.visible {
			continue
		}
		switch name {
		case "text":
			if err := composer.renderText(target, child, childState, childTransform, rootView, clips); err != nil {
				return err
			}
		case "image":
			if err := composer.renderImage(target, child, childTransform, rootView, depth); err != nil {
				return err
			}
		default:
			if err := composer.renderChildren(target, child, childState, childTransform, rootView, clips, depth); err != nil {
				return err
			}
		}
	}
	return nil
}

func (composer *pdfSVGComposer) renderText(
	target *canvas.Canvas,
	node *pdfSVGNode,
	state pdfSVGState,
	transform canvas.Matrix,
	rootView canvas.Matrix,
	clips map[string]pdfSVGRect,
) error {
	x, err := pdfSVGDimension(node.attrs["x"], 0)
	if err != nil {
		return fmt.Errorf("render PDF SVG <text> x: %w", err)
	}
	y, err := pdfSVGDimension(node.attrs["y"], 0)
	if err != nil {
		return fmt.Errorf("render PDF SVG <text> y: %w", err)
	}
	var clip *pdfSVGRect
	if clipID := pdfSVGURLFragment(node.attrs["clip-path"]); clipID != "" {
		if rect, ok := clips[clipID]; ok {
			clip = &rect
		}
	}

	for _, content := range node.content {
		if content.node == nil {
			text := normalizePDFSVGText(content.text)
			if text == "" {
				continue
			}
			advance, err := composer.drawTextLine(target, text, x, y, state, transform, rootView, clip)
			if err != nil {
				return err
			}
			x += advance
			continue
		}
		span := content.node
		if strings.ToLower(span.name.Local) != "tspan" {
			continue
		}
		spanState := applyPDFSVGState(state, span.attrs)
		spanTransform := transform
		if value := strings.TrimSpace(span.attrs["transform"]); value != "" {
			matrix, err := parsePDFSVGTransform(value)
			if err != nil {
				return fmt.Errorf("render PDF SVG <tspan> transform: %w", err)
			}
			spanTransform = spanTransform.Mul(matrix)
		}
		if value := strings.TrimSpace(span.attrs["x"]); value != "" {
			x, err = pdfSVGDimension(value, 0)
			if err != nil {
				return fmt.Errorf("render PDF SVG <tspan> x: %w", err)
			}
		}
		if value := strings.TrimSpace(span.attrs["y"]); value != "" {
			y, err = pdfSVGDimension(value, 0)
			if err != nil {
				return fmt.Errorf("render PDF SVG <tspan> y: %w", err)
			}
		}
		if value := strings.TrimSpace(span.attrs["dx"]); value != "" {
			delta, deltaErr := pdfSVGDimension(value, 0)
			if deltaErr != nil {
				return fmt.Errorf("render PDF SVG <tspan> dx: %w", deltaErr)
			}
			x += delta
		}
		if value := strings.TrimSpace(span.attrs["dy"]); value != "" {
			delta, deltaErr := pdfSVGDimension(value, 0)
			if deltaErr != nil {
				return fmt.Errorf("render PDF SVG <tspan> dy: %w", deltaErr)
			}
			y += delta
		}
		text := normalizePDFSVGText(pdfSVGNodeText(span))
		if text == "" || !spanState.visible {
			continue
		}
		advance, err := composer.drawTextLine(target, text, x, y, spanState, spanTransform, rootView, clip)
		if err != nil {
			return err
		}
		x += advance
	}
	return nil
}

func (composer *pdfSVGComposer) drawTextLine(
	target *canvas.Canvas,
	text string,
	x, y float64,
	state pdfSVGState,
	transform canvas.Matrix,
	rootView canvas.Matrix,
	clip *pdfSVGRect,
) (float64, error) {
	if state.fontSize <= 0 || math.IsNaN(state.fontSize) || math.IsInf(state.fontSize, 0) {
		return 0, fmt.Errorf("render PDF SVG text: font-size must be positive and finite")
	}
	fill := parsePDFSVGColor(state.fill)
	fill = pdfSVGColorOpacity(fill, state.fillOpacity*state.opacity)
	if fill.A == 0 {
		return 0, nil
	}
	align := canvas.Left
	switch strings.ToLower(strings.TrimSpace(state.textAnchor)) {
	case "middle":
		align = canvas.Center
	case "end":
		align = canvas.Right
	}
	face := composer.fontFamily.Face(state.fontSize*pdfSVGPointsPerMillimeter, fill, state.fontStyle)
	text = clipPDFSVGText(text, face, x, state.textAnchor, clip)
	if text == "" {
		return 0, nil
	}
	line := canvas.NewTextLine(face, text, align)
	context := canvas.NewContext(target)
	context.SetCoordSystem(canvas.CartesianIV)
	context.SetView(rootView.Mul(transform))
	context.DrawText(x, y, line)
	return face.TextWidth(text), nil
}

func (composer *pdfSVGComposer) renderImage(
	target *canvas.Canvas,
	node *pdfSVGNode,
	transform canvas.Matrix,
	rootView canvas.Matrix,
	depth int,
) error {
	href := strings.TrimSpace(node.attrs["href"])
	if href == "" {
		return fmt.Errorf("render PDF SVG <image>: href is required")
	}
	x, err := pdfSVGDimension(node.attrs["x"], 0)
	if err != nil {
		return fmt.Errorf("render PDF SVG <image> x: %w", err)
	}
	y, err := pdfSVGDimension(node.attrs["y"], 0)
	if err != nil {
		return fmt.Errorf("render PDF SVG <image> y: %w", err)
	}
	width, err := pdfSVGDimension(node.attrs["width"], 0)
	if err != nil {
		return fmt.Errorf("render PDF SVG <image> width: %w", err)
	}
	height, err := pdfSVGDimension(node.attrs["height"], 0)
	if err != nil {
		return fmt.Errorf("render PDF SVG <image> height: %w", err)
	}
	if !positiveFinitePDFDimension(width) || !positiveFinitePDFDimension(height) {
		return fmt.Errorf("render PDF SVG <image>: width and height must be positive and finite")
	}

	mediaType, payload, err := decodePDFSVGDataURL(href)
	if err != nil {
		return fmt.Errorf("render PDF SVG <image>: %w", err)
	}
	if mediaType == "image/svg+xml" || bytes.HasPrefix(bytes.TrimSpace(payload), []byte("<svg")) || bytes.HasPrefix(bytes.TrimSpace(payload), []byte("<?xml")) {
		imageCanvas, err := composer.parseDepth(payload, depth+1)
		if err != nil {
			return fmt.Errorf("render PDF SVG embedded image: %w", err)
		}
		imageWidth, imageHeight := imageCanvas.Size()
		if !positiveFinitePDFDimension(imageWidth) || !positiveFinitePDFDimension(imageHeight) {
			return fmt.Errorf("render PDF SVG embedded image: dimensions must be positive and finite")
		}
		scale := math.Min(width/imageWidth, height/imageHeight)
		dx := (width - imageWidth*scale) / 2
		dy := (height - imageHeight*scale) / 2
		parentFlip := canvas.Identity.ReflectYAbout(target.H / 2)
		imageFlip := canvas.Identity.ReflectYAbout(imageHeight / 2)
		view := parentFlip.Mul(rootView).Mul(transform).
			Translate(x+dx, y+dy).
			Scale(scale, scale).
			Mul(imageFlip.Inv())
		imageCanvas.RenderViewTo(target, view)
		return nil
	}

	raster, _, err := image.Decode(bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("decode %q image: %w", mediaType, err)
	}
	context := canvas.NewContext(target)
	context.SetCoordSystem(canvas.CartesianIV)
	context.SetView(rootView.Mul(transform))
	context.FitImage(raster, canvas.Rect{X0: x, Y0: y, X1: x + width, Y1: y + height}, canvas.ImageContain)
	return nil
}

func decodePDFSVGDataURL(value string) (string, []byte, error) {
	if !strings.HasPrefix(strings.ToLower(value), "data:") {
		return "", nil, fmt.Errorf("only data: image URLs are supported")
	}
	header, encoded, found := strings.Cut(value[5:], ",")
	if !found {
		return "", nil, fmt.Errorf("invalid data URL")
	}
	parts := strings.Split(header, ";")
	mediaType := strings.ToLower(strings.TrimSpace(parts[0]))
	if mediaType == "" {
		mediaType = "text/plain"
	}
	base64Encoded := false
	for _, parameter := range parts[1:] {
		if strings.EqualFold(strings.TrimSpace(parameter), "base64") {
			base64Encoded = true
		}
	}
	if base64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			return "", nil, fmt.Errorf("decode base64 data URL: %w", err)
		}
		return mediaType, decoded, nil
	}
	decoded, err := url.PathUnescape(encoded)
	if err != nil {
		return "", nil, fmt.Errorf("decode data URL: %w", err)
	}
	return mediaType, []byte(decoded), nil
}

func pdfSVGRootView(root *pdfSVGNode, width, height float64) (canvas.Matrix, error) {
	viewBox := strings.TrimSpace(root.attrs["viewBox"])
	if viewBox == "" {
		return canvas.Identity, nil
	}
	values, err := pdfSVGNumbers(viewBox)
	if err != nil || len(values) != 4 {
		return canvas.Identity, fmt.Errorf("parse SVG: invalid viewBox %q", viewBox)
	}
	if !positiveFinitePDFDimension(values[2]) || !positiveFinitePDFDimension(values[3]) {
		return canvas.Identity, fmt.Errorf("parse SVG: viewBox dimensions must be positive and finite")
	}
	return canvas.Identity.Scale(width/values[2], height/values[3]).Translate(-values[0], -values[1]), nil
}

func collectPDFSVGClips(root *pdfSVGNode) map[string]pdfSVGRect {
	clips := map[string]pdfSVGRect{}
	var visit func(*pdfSVGNode)
	visit = func(node *pdfSVGNode) {
		if strings.EqualFold(node.name.Local, "clipPath") {
			id := strings.TrimSpace(node.attrs["id"])
			if id != "" {
				for _, content := range node.content {
					if content.node == nil || !strings.EqualFold(content.node.name.Local, "rect") {
						continue
					}
					x, xErr := pdfSVGDimension(content.node.attrs["x"], 0)
					y, yErr := pdfSVGDimension(content.node.attrs["y"], 0)
					width, widthErr := pdfSVGDimension(content.node.attrs["width"], 0)
					height, heightErr := pdfSVGDimension(content.node.attrs["height"], 0)
					if xErr == nil && yErr == nil && widthErr == nil && heightErr == nil &&
						positiveFinitePDFDimension(width) && positiveFinitePDFDimension(height) {
						clips[id] = pdfSVGRect{x: x, y: y, width: width, height: height}
					}
					break
				}
			}
		}
		for _, content := range node.content {
			if content.node != nil {
				visit(content.node)
			}
		}
	}
	visit(root)
	return clips
}

func pdfSVGURLFragment(value string) string {
	value = strings.TrimSpace(value)
	if !strings.HasPrefix(value, "url(") || !strings.HasSuffix(value, ")") {
		return ""
	}
	fragment := strings.TrimSpace(value[4 : len(value)-1])
	fragment = strings.Trim(fragment, `"'`)
	return strings.TrimPrefix(fragment, "#")
}

// clipPDFSVGText enforces Xaligo's rectangular text clip paths at glyph
// boundaries. Canvas' renderer interface has no clipping primitive; limiting
// the rendered run keeps fallback-font metrics from leaking labels outside the
// generated text box while retaining PDF text objects and searchability.
func clipPDFSVGText(text string, face *canvas.FontFace, x float64, anchor string, clip *pdfSVGRect) string {
	if clip == nil || text == "" {
		return text
	}
	available := 0.0
	switch strings.ToLower(strings.TrimSpace(anchor)) {
	case "middle":
		available = 2 * math.Min(x-clip.x, clip.x+clip.width-x)
	case "end":
		available = x - clip.x
	default:
		available = clip.x + clip.width - x
	}
	if available <= 0 {
		return ""
	}
	if face.TextWidth(text) <= available {
		return text
	}
	runes := []rune(text)
	keepSuffix := strings.EqualFold(strings.TrimSpace(anchor), "end")
	low, high := 0, len(runes)
	for low < high {
		middle := (low + high + 1) / 2
		candidate := string(runes[:middle])
		if keepSuffix {
			candidate = string(runes[len(runes)-middle:])
		}
		if face.TextWidth(candidate) <= available {
			low = middle
		} else {
			high = middle - 1
		}
	}
	if keepSuffix {
		return string(runes[len(runes)-low:])
	}
	return string(runes[:low])
}

func applyPDFSVGState(parent pdfSVGState, attrs map[string]string) pdfSVGState {
	state := parent
	properties := make(map[string]string, len(attrs))
	for key, value := range attrs {
		properties[strings.ToLower(strings.TrimSpace(key))] = strings.TrimSpace(value)
	}
	if style := properties["style"]; style != "" {
		for _, declaration := range strings.Split(style, ";") {
			key, value, found := strings.Cut(declaration, ":")
			if found {
				properties[strings.ToLower(strings.TrimSpace(key))] = strings.TrimSpace(value)
			}
		}
	}
	if value := properties["fill"]; value != "" {
		state.fill = value
	}
	if value := properties["fill-opacity"]; value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			state.fillOpacity = clampPDFSVGUnit(parsed)
		}
	}
	if value := properties["opacity"]; value != "" {
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			state.opacity *= clampPDFSVGUnit(parsed)
		}
	}
	if value := properties["font-size"]; value != "" {
		if parsed, err := pdfSVGDimension(value, state.fontSize); err == nil {
			state.fontSize = parsed
		}
	}
	if value := strings.ToLower(properties["font-weight"]); value != "" {
		weight := canvas.FontRegular
		if value == "bold" || value == "bolder" {
			weight = canvas.FontBold
		} else if numeric, err := strconv.Atoi(value); err == nil && numeric >= 600 {
			weight = canvas.FontBold
		}
		state.fontStyle = state.fontStyle&canvas.FontItalic | weight
	}
	if value := strings.ToLower(properties["font-style"]); value != "" {
		state.fontStyle &^= canvas.FontItalic
		if value == "italic" || value == "oblique" {
			state.fontStyle |= canvas.FontItalic
		}
	}
	if value := properties["text-anchor"]; value != "" {
		state.textAnchor = value
	}
	if strings.EqualFold(properties["display"], "none") || strings.EqualFold(properties["visibility"], "hidden") {
		state.visible = false
	}
	return state
}

func parsePDFSVGTransform(value string) (canvas.Matrix, error) {
	remaining := strings.TrimSpace(value)
	matrix := canvas.Identity
	for remaining != "" {
		open := strings.IndexByte(remaining, '(')
		if open <= 0 {
			return canvas.Identity, fmt.Errorf("invalid transform %q", value)
		}
		close := strings.IndexByte(remaining[open+1:], ')')
		if close < 0 {
			return canvas.Identity, fmt.Errorf("invalid transform %q", value)
		}
		close += open + 1
		name := strings.ToLower(strings.TrimSpace(remaining[:open]))
		numbers, err := pdfSVGNumbers(remaining[open+1 : close])
		if err != nil {
			return canvas.Identity, err
		}
		switch name {
		case "matrix":
			if len(numbers) != 6 {
				return canvas.Identity, fmt.Errorf("matrix() requires 6 numbers")
			}
			matrix = matrix.Mul(canvas.Matrix{{numbers[0], numbers[2], numbers[4]}, {numbers[1], numbers[3], numbers[5]}})
		case "translate":
			if len(numbers) != 1 && len(numbers) != 2 {
				return canvas.Identity, fmt.Errorf("translate() requires 1 or 2 numbers")
			}
			y := 0.0
			if len(numbers) == 2 {
				y = numbers[1]
			}
			matrix = matrix.Translate(numbers[0], y)
		case "scale":
			if len(numbers) != 1 && len(numbers) != 2 {
				return canvas.Identity, fmt.Errorf("scale() requires 1 or 2 numbers")
			}
			y := numbers[0]
			if len(numbers) == 2 {
				y = numbers[1]
			}
			matrix = matrix.Scale(numbers[0], y)
		case "rotate":
			if len(numbers) != 1 && len(numbers) != 3 {
				return canvas.Identity, fmt.Errorf("rotate() requires 1 or 3 numbers")
			}
			if len(numbers) == 1 {
				matrix = matrix.Rotate(numbers[0])
			} else {
				matrix = matrix.RotateAbout(numbers[0], numbers[1], numbers[2])
			}
		case "skewx":
			if len(numbers) != 1 {
				return canvas.Identity, fmt.Errorf("skewX() requires 1 number")
			}
			matrix = matrix.Shear(math.Tan(numbers[0]*math.Pi/180), 0)
		case "skewy":
			if len(numbers) != 1 {
				return canvas.Identity, fmt.Errorf("skewY() requires 1 number")
			}
			matrix = matrix.Shear(0, math.Tan(numbers[0]*math.Pi/180))
		default:
			return canvas.Identity, fmt.Errorf("unsupported transform %q", name)
		}
		remaining = strings.TrimSpace(strings.TrimLeft(remaining[close+1:], ","))
	}
	return matrix, nil
}

func pdfSVGDimension(value string, percentageReference float64) (float64, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, nil
	}
	match := pdfSVGNumberPattern.FindStringIndex(value)
	if match == nil || match[0] != 0 {
		return 0, fmt.Errorf("invalid dimension %q", value)
	}
	number, err := strconv.ParseFloat(value[:match[1]], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid dimension %q: %w", value, err)
	}
	unit := strings.ToLower(strings.TrimSpace(value[match[1]:]))
	if separator := strings.IndexAny(unit, " ,"); separator >= 0 {
		unit = unit[:separator]
	}
	switch unit {
	case "", "px":
		return number, nil
	case "%":
		return number * percentageReference / 100, nil
	case "cm":
		return number * 10 * 96 / 25.4, nil
	case "mm":
		return number * 96 / 25.4, nil
	case "q":
		return number * 0.25 * 96 / 25.4, nil
	case "in":
		return number * 96, nil
	case "pc":
		return number * 96 / 6, nil
	case "pt":
		return number * 96 / 72, nil
	default:
		return 0, fmt.Errorf("unsupported dimension unit %q", unit)
	}
}

func pdfSVGNumbers(value string) ([]float64, error) {
	matches := pdfSVGNumberPattern.FindAllString(value, -1)
	if len(matches) == 0 && strings.TrimSpace(value) != "" {
		return nil, fmt.Errorf("invalid number list %q", value)
	}
	values := make([]float64, 0, len(matches))
	for _, match := range matches {
		value, err := strconv.ParseFloat(match, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number %q: %w", match, err)
		}
		values = append(values, value)
	}
	return values, nil
}

func parsePDFSVGColor(value string) stdcolor.RGBA {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case "", "black", "currentcolor":
		return canvas.Black
	case "white":
		return canvas.White
	case "transparent", "none":
		return canvas.Transparent
	case "red":
		return canvas.Red
	case "green":
		return canvas.Green
	case "blue":
		return canvas.Blue
	}
	if strings.HasPrefix(value, "#") {
		return canvas.Hex(value)
	}
	if strings.HasPrefix(value, "rgb(") && strings.HasSuffix(value, ")") {
		parts := strings.Split(value[4:len(value)-1], ",")
		if len(parts) == 3 {
			components := [3]float64{}
			for index, part := range parts {
				part = strings.TrimSpace(part)
				if strings.HasSuffix(part, "%") {
					parsed, err := strconv.ParseFloat(strings.TrimSuffix(part, "%"), 64)
					if err != nil {
						return canvas.Black
					}
					components[index] = parsed / 100
				} else {
					parsed, err := strconv.ParseFloat(part, 64)
					if err != nil {
						return canvas.Black
					}
					components[index] = parsed / 255
				}
			}
			return canvas.RGB(clampPDFSVGUnit(components[0]), clampPDFSVGUnit(components[1]), clampPDFSVGUnit(components[2]))
		}
	}
	return canvas.Black
}

func pdfSVGColorOpacity(input stdcolor.RGBA, opacity float64) stdcolor.RGBA {
	opacity = clampPDFSVGUnit(opacity)
	if opacity == 1 || input.A == 0 {
		return input
	}
	alpha := float64(input.A) / 255
	red := float64(input.R) / float64(input.A)
	green := float64(input.G) / float64(input.A)
	blue := float64(input.B) / float64(input.A)
	return canvas.RGBA(red, green, blue, alpha*opacity)
}

func clampPDFSVGUnit(value float64) float64 {
	return math.Max(0, math.Min(1, value))
}

func pdfSVGNodeText(node *pdfSVGNode) string {
	var builder strings.Builder
	for _, content := range node.content {
		if content.node != nil {
			builder.WriteString(pdfSVGNodeText(content.node))
		} else {
			builder.WriteString(content.text)
		}
	}
	return builder.String()
}

func normalizePDFSVGText(value string) string {
	return strings.Join(strings.Fields(value), " ")
}
