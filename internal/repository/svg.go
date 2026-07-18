package repository

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"math"
	"strings"
	"unicode"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

const svgDefaultPxPerInch = 96.0

type SVGRepository interface {
	Render(plan entity.Plan, pxPerInch float64, legendPosition string) ([]byte, error)
}

type svgRepository struct{}

func NewSVGRepository() SVGRepository { return &svgRepository{} }

// RenderPlan converts the shared draw plan into an SVG document.
func (rcvr *svgRepository) Render(plan entity.Plan, pxPerInch float64, legendPosition string) ([]byte, error) {
	if pxPerInch <= 0 {
		pxPerInch = svgDefaultPxPerInch
	}
	if plan.Slide.W <= 0 || plan.Slide.H <= 0 {
		return nil, fmt.Errorf("SVG slide size must be positive")
	}

	w := plan.Slide.W * pxPerInch
	h := plan.Slide.H * pxPerInch
	bounds := svgOpsBounds(plan.Ops, pxPerInch, w, h)
	const diagramPad = 24.0
	pad := 0.0
	if bounds.hasPath || bounds.minX < 0 || bounds.minY < 0 || bounds.maxX > w || bounds.maxY > h {
		pad = diagramPad
	}
	diagramOffsetX := pad - bounds.minX
	diagramOffsetY := pad - bounds.minY
	diagramW := bounds.maxX - bounds.minX + pad*2
	diagramH := bounds.maxY - bounds.minY + pad*2
	layout := svgLegendLayout(diagramW, diagramH, plan.Legend, legendPosition)

	var b bytes.Buffer
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%s" height="%s" viewBox="0 0 %s %s">`+"\n", num(layout.canvasW), num(layout.canvasH), num(layout.canvasW), num(layout.canvasH))
	writeMarkerDefinitions(&b)
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%s" height="%s" fill="#%s"/>`+"\n", num(layout.canvasW), num(layout.canvasH), color(plan.Slide.Background, "FFFFFF"))

	fmt.Fprintf(&b, `<g transform="translate(%s %s)">`+"\n", num(layout.diagramX+diagramOffsetX), num(layout.diagramY+diagramOffsetY))
	for _, op := range plan.Ops {
		writeOp(&b, op, pxPerInch)
	}
	b.WriteString("</g>\n")
	writeLegend(&b, plan.Legend, layout)
	b.WriteString("</svg>\n")
	return b.Bytes(), nil
}

type svgBounds struct {
	minX    float64
	minY    float64
	maxX    float64
	maxY    float64
	hasPath bool
}

func svgOpsBounds(ops []entity.DrawOp, ppi, w, h float64) svgBounds {
	bounds := svgBounds{minX: 0, minY: 0, maxX: w, maxY: h}
	pathPad := 0.0
	expand := func(x, y float64) {
		bounds.minX = math.Min(bounds.minX, x)
		bounds.minY = math.Min(bounds.minY, y)
		bounds.maxX = math.Max(bounds.maxX, x)
		bounds.maxY = math.Max(bounds.maxY, y)
	}
	for _, op := range ops {
		x, y, ow, oh := op.X*ppi, op.Y*ppi, op.W*ppi, op.H*ppi
		switch op.Kind {
		case "line", "polygon":
			bounds.hasPath = true
			pathPad = math.Max(pathPad, svgPathBoundsPad(op, ppi))
			for _, p := range absolutePoints(op, x, y, ow, oh, ppi) {
				expand(p.x, p.y)
			}
		case "text":
			expand(x, y)
			expand(x+ow, y+oh)
			layout := resolvedTextLayout(op)
			if layout.Clip {
				continue
			}
			metrics, ok := resolveSVGTextMetrics(op, x, y, ow, oh, ppi)
			if !ok {
				continue
			}
			minX, minY, maxX, maxY := svgTextGlyphBounds(op, metrics, x, y, ow, oh)
			expand(minX, minY)
			expand(maxX, maxY)
		default:
			expand(x, y)
			expand(x+ow, y+oh)
		}
	}
	if bounds.hasPath {
		pathPad = math.Max(18, pathPad)
		bounds.minX -= pathPad
		bounds.minY -= pathPad
		bounds.maxX += pathPad
		bounds.maxY += pathPad
	}
	return bounds
}

func svgPathBoundsPad(op entity.DrawOp, ppi float64) float64 {
	if op.Line == nil || op.Line.Width <= 0 || ppi <= 0 {
		return 0
	}
	strokeWidth := op.Line.Width * ppi / 72.0
	if math.IsNaN(strokeWidth) || math.IsInf(strokeWidth, 0) || strokeWidth <= 0 {
		return 0
	}
	// SVG's default miter join can extend beyond half the stroke width. Markers
	// use markerUnits="strokeWidth", so their perpendicular extent grows with
	// the resolved output stroke instead of remaining inside a fixed pixel pad.
	pad := strokeWidth * 2
	if op.Kind == "line" {
		pad = math.Max(pad, strokeWidth*(markerReachInStrokeWidths(op.Line.BeginArrowType)+0.5))
		pad = math.Max(pad, strokeWidth*(markerReachInStrokeWidths(op.Line.EndArrowType)+0.5))
	}
	return pad
}

func markerReachInStrokeWidths(arrowType string) float64 {
	switch arrowType {
	case "", "none":
		return 0
	case "diamond":
		return math.Sqrt(11*11 + 4*4)
	case "oval":
		return 3.5
	case "triangle":
		return math.Sqrt(7*7 + 5*5)
	default:
		// arrow, stealth, and the legacy unknown-value fallback have a 9-unit
		// rear reach and a 5-unit half-height from the marker reference point.
		return math.Sqrt(9*9 + 5*5)
	}
}

type svgLegendBox struct {
	canvasW  float64
	canvasH  float64
	diagramX float64
	diagramY float64
	legendX  float64
	legendY  float64
	legendW  float64
	legendH  float64
	visible  bool
}

func svgLegendLayout(w, h float64, entries []entity.LegendEntry, position string) svgLegendBox {
	layout := svgLegendBox{canvasW: w, canvasH: h}
	if len(entries) == 0 {
		return layout
	}
	position = strings.ToLower(strings.TrimSpace(position))
	if position == "" {
		position = "bottom"
	}
	const (
		gap       = 24.0
		pad       = 16.0
		rowH      = 38.0
		sideW     = 280.0
		bottomMin = 360.0
	)
	rows := float64(len(entries))
	layout.visible = true
	switch position {
	case "top":
		layout.legendW = math.Max(bottomMin, w)
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = math.Max(w, layout.legendW)
		layout.canvasH = h + gap + layout.legendH
		layout.diagramY = layout.legendH + gap
	case "right":
		layout.legendW = sideW
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = w + gap + layout.legendW
		layout.canvasH = math.Max(h, layout.legendH)
		layout.legendX = w + gap
	case "left":
		layout.legendW = sideW
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = w + gap + layout.legendW
		layout.canvasH = math.Max(h, layout.legendH)
		layout.diagramX = layout.legendW + gap
	default:
		layout.legendW = math.Max(bottomMin, w)
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = math.Max(w, layout.legendW)
		layout.canvasH = h + gap + layout.legendH
		layout.legendY = h + gap
	}
	return layout
}

func writeLegend(b *bytes.Buffer, entries []entity.LegendEntry, layout svgLegendBox) {
	if !layout.visible {
		return
	}
	const (
		pad      = 16.0
		rowH     = 38.0
		iconSize = 24.0
	)
	fmt.Fprintf(b, `<g id="xaligo-svg-legend" transform="translate(%s %s)">`+"\n", num(layout.legendX), num(layout.legendY))
	fmt.Fprintf(b, `<rect x="0" y="0" width="%s" height="%s" fill="#FFFFFF" stroke="#CBD5E1" stroke-width="1"/>`+"\n", num(layout.legendW), num(layout.legendH))
	for i, entry := range entries {
		y := pad + float64(i)*rowH
		if entry.Data != "" {
			fmt.Fprintf(b, `<image x="%s" y="%s" width="%s" height="%s" href="%s"/>`+"\n", num(pad), num(y), num(iconSize), num(iconSize), attr(entry.Data))
		}
		textX := pad + iconSize + 10
		fmt.Fprintf(b, `<text x="%s" y="%s" fill="#0F172A" font-family="Arial" font-size="12" font-weight="700">%s</text>`+"\n", num(textX), num(y+11), text(entry.Abbreviation))
		fmt.Fprintf(b, `<text x="%s" y="%s" fill="#475569" font-family="Arial" font-size="10">%s</text>`+"\n", num(textX), num(y+25), text(entry.OfficialName))
	}
	b.WriteString("</g>\n")
}

func writeOp(b *bytes.Buffer, op entity.DrawOp, ppi float64) {
	x, y, w, h := op.X*ppi, op.Y*ppi, op.W*ppi, op.H*ppi
	transform := rotateAttr(op.Rotate, x+w/2, y+h/2)
	switch op.Kind {
	case "rect":
		fmt.Fprintf(b, `<rect x="%s" y="%s" width="%s" height="%s"%s%s%s/>`+"\n", num(x), num(y), num(w), num(h), fillAttrs(op.Fill), lineAttrs(op.Line, ppi), transform)
	case "ellipse":
		fmt.Fprintf(b, `<ellipse cx="%s" cy="%s" rx="%s" ry="%s"%s%s%s/>`+"\n", num(x+w/2), num(y+h/2), num(w/2), num(h/2), fillAttrs(op.Fill), lineAttrs(op.Line, ppi), transform)
	case "polygon":
		writePolygon(b, op, x, y, w, h, ppi, transform)
	case "text":
		writeText(b, op, x, y, w, h, ppi, transform)
	case "image":
		if op.Data == "" {
			return
		}
		fmt.Fprintf(b, `<image x="%s" y="%s" width="%s" height="%s" href="%s" opacity="%s"%s/>`+"\n", num(x), num(y), num(w), num(h), attr(op.Data), num(opacity(op.Transparency)), transform)
	case "line":
		writeLine(b, op, x, y, w, h, ppi)
	}
}

func writePolygon(b *bytes.Buffer, op entity.DrawOp, x, y, w, h, ppi float64, transform string) {
	points := absolutePoints(op, x, y, w, h, ppi)
	if len(points) < 3 {
		return
	}
	var value strings.Builder
	for i, p := range points {
		if i > 0 {
			value.WriteByte(' ')
		}
		fmt.Fprintf(&value, "%s,%s", num(p.x), num(p.y))
	}
	fmt.Fprintf(b, `<polygon points="%s"%s%s%s/>`+"\n", value.String(), fillAttrs(op.Fill), lineAttrs(op.Line, ppi), transform)
}

func writeText(b *bytes.Buffer, op entity.DrawOp, x, y, w, h, ppi float64, transform string) {
	if op.Text == "" {
		return
	}
	metrics, ok := resolveSVGTextMetrics(op, x, y, w, h, ppi)
	if !ok {
		return
	}
	weight := ""
	if op.Bold {
		weight = ` font-weight="700"`
	}
	clip := ""
	if metrics.layout.Clip {
		clipID := svgTextClipID(op, x, y, w, h)
		fmt.Fprintf(b, `<defs><clipPath id="%s" clipPathUnits="userSpaceOnUse"><rect x="%s" y="%s" width="%s" height="%s"/></clipPath></defs>`+"\n",
			clipID, num(x), num(y), num(w), num(h))
		clip = ` clip-path="url(#` + clipID + `)"`
	}
	fmt.Fprintf(b, `<text x="%s" y="%s" fill="#%s" font-family="%s" font-size="%s" text-anchor="%s"%s%s%s>`,
		num(metrics.textX), num(metrics.textY), color(op.Color, "1E1E1E"), attr(op.FontFace), num(metrics.fontSize), metrics.anchor, weight, clip, transform)
	for i, line := range metrics.lines {
		if i == 0 {
			fmt.Fprintf(b, `<tspan x="%s" y="%s">%s</tspan>`, num(metrics.textX), num(metrics.textY), text(line))
			continue
		}
		fmt.Fprintf(b, `<tspan x="%s" dy="%s">%s</tspan>`, num(metrics.textX), num(metrics.lineStep), text(line))
	}
	b.WriteString("</text>\n")
}

type svgTextMetrics struct {
	layout                     entity.TextLayout
	lines                      []string
	fontSize, lineStep, textH  float64
	textX, textY, maxTextWidth float64
	anchor                     string
}

func resolveSVGTextMetrics(op entity.DrawOp, x, y, w, h, ppi float64) (svgTextMetrics, bool) {
	layout := resolvedTextLayout(op)
	contentX := x + math.Max(0, layout.Padding.Left*ppi)
	contentY := y + math.Max(0, layout.Padding.Top*ppi)
	contentW := w - math.Max(0, (layout.Padding.Left+layout.Padding.Right)*ppi)
	contentH := h - math.Max(0, (layout.Padding.Top+layout.Padding.Bottom)*ppi)
	if contentW <= 0 || contentH <= 0 {
		return svgTextMetrics{}, false
	}
	fontSize := op.FontSize * ppi / 72.0
	if fontSize <= 0 {
		fontSize = 1
	}
	lineHeight := layout.LineHeight
	if lineHeight <= 0 || math.IsNaN(lineHeight) || math.IsInf(lineHeight, 0) {
		lineHeight = 1.2
	}
	lines := svgTextLines(op.Text, contentW, fontSize, op.Bold, layout.Wrap)
	maxTextW := 0.0
	for _, line := range lines {
		maxTextW = math.Max(maxTextW, svgTextWidth(line, fontSize, op.Bold))
	}
	if layout.Fit == entity.TextFitShrink {
		textH := float64(len(lines)) * fontSize * lineHeight
		scale := 1.0
		if maxTextW > contentW {
			scale = math.Min(scale, contentW/maxTextW)
		}
		if textH > contentH {
			scale = math.Min(scale, contentH/textH)
		}
		fontSize = math.Max(0.1, fontSize*scale)
		maxTextW = 0
		for _, line := range lines {
			maxTextW = math.Max(maxTextW, svgTextWidth(line, fontSize, op.Bold))
		}
	}
	lineStep := fontSize * lineHeight
	textH := float64(len(lines)) * lineStep
	anchor := "start"
	textX := contentX
	if op.Align == "center" {
		anchor = "middle"
		textX = contentX + contentW/2
	} else if op.Align == "right" {
		anchor = "end"
		textX = contentX + contentW
	}
	textY := contentY + fontSize
	if op.Valign == "middle" {
		textY = contentY + (contentH-textH)/2 + fontSize
	} else if op.Valign == "bottom" {
		textY = contentY + contentH - textH + fontSize
	}
	return svgTextMetrics{
		layout: layout, lines: lines, fontSize: fontSize, lineStep: lineStep,
		textH: textH, textX: textX, textY: textY, maxTextWidth: maxTextW, anchor: anchor,
	}, true
}

func svgTextGlyphBounds(op entity.DrawOp, metrics svgTextMetrics, x, y, w, h float64) (float64, float64, float64, float64) {
	// The SVG encoder cannot query the eventual viewer's font metrics. Use the
	// wrapping estimate for placement and a wider per-rune estimate for visible
	// overflow bounds so alternate/fallback fonts are not cropped by the viewport.
	boundsWidth := metrics.maxTextWidth
	for _, line := range metrics.lines {
		boundsWidth = math.Max(boundsWidth, svgConservativeTextWidth(line, metrics.fontSize, op.Bold))
	}
	minX := metrics.textX
	switch metrics.anchor {
	case "middle":
		minX -= boundsWidth / 2
	case "end":
		minX -= boundsWidth
	}
	minY := metrics.textY - metrics.fontSize
	maxX := minX + boundsWidth
	maxY := minY + metrics.textH
	if math.Abs(op.Rotate) < 0.0001 {
		return minX, minY, maxX, maxY
	}
	return rotatedRectBounds(minX, minY, maxX, maxY, op.Rotate, x+w/2, y+h/2)
}

func svgConservativeTextWidth(value string, fontSize float64, bold bool) float64 {
	units := 0.0
	for _, r := range value {
		switch {
		case unicode.Is(unicode.Mn, r), unicode.Is(unicode.Me, r):
			continue
		case unicode.IsSpace(r):
			units += 0.5
		default:
			units += 1.1
		}
	}
	if bold {
		units *= 1.05
	}
	return units * fontSize
}

func rotatedRectBounds(minX, minY, maxX, maxY, degrees, centerX, centerY float64) (float64, float64, float64, float64) {
	radians := degrees * math.Pi / 180
	sin, cos := math.Sin(radians), math.Cos(radians)
	outMinX, outMinY := math.Inf(1), math.Inf(1)
	outMaxX, outMaxY := math.Inf(-1), math.Inf(-1)
	for _, point := range [][2]float64{{minX, minY}, {maxX, minY}, {maxX, maxY}, {minX, maxY}} {
		dx, dy := point[0]-centerX, point[1]-centerY
		px := centerX + dx*cos - dy*sin
		py := centerY + dx*sin + dy*cos
		outMinX = math.Min(outMinX, px)
		outMinY = math.Min(outMinY, py)
		outMaxX = math.Max(outMaxX, px)
		outMaxY = math.Max(outMaxY, py)
	}
	return outMinX, outMinY, outMaxX, outMaxY
}

func resolvedTextLayout(op entity.DrawOp) entity.TextLayout {
	if op.TextLayout != nil {
		layout := *op.TextLayout
		if layout.Fit != entity.TextFitShrink {
			layout.Fit = entity.TextFitNone
		}
		switch layout.Overflow {
		case entity.TextOverflowVisible:
			layout.Clip = false
		case entity.TextOverflowClip:
			layout.Clip = true
		default:
			if layout.Clip {
				layout.Overflow = entity.TextOverflowClip
			} else {
				layout.Overflow = entity.TextOverflowVisible
			}
		}
		return layout
	}
	// Plans produced before textLayout used shrink-to-fit everywhere and a
	// no-wrap ID convention for group headers. Preserve that behavior while
	// containing text inside its declared box in SVG.
	role := entity.TextRoleLabel
	wrap := true
	if legacyGroupHeaderLabelID(op.ID) {
		role = entity.TextRoleGroupHeader
		wrap = false
	}
	return entity.TextLayout{
		Role:       role,
		Wrap:       wrap,
		Fit:        entity.TextFitShrink,
		Overflow:   entity.TextOverflowClip,
		Clip:       true,
		LineHeight: 1.2,
	}
}

func legacyGroupHeaderLabelID(id string) bool {
	return id != "" && strings.HasSuffix(id, "-label") &&
		!strings.HasSuffix(id, "-item-lbl") && !isConnectorLabelID(id)
}

func isConnectorLabelID(id string) bool {
	if !strings.HasPrefix(id, "L") || !strings.HasSuffix(id, "-label") {
		return false
	}
	digits := strings.TrimSuffix(strings.TrimPrefix(id, "L"), "-label")
	if digits == "" {
		return false
	}
	for _, r := range digits {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func svgTextLines(value string, maxWidth, fontSize float64, bold, wrap bool) []string {
	paragraphs := strings.Split(strings.ReplaceAll(value, "\r\n", "\n"), "\n")
	lines := make([]string, 0, len(paragraphs))
	for _, paragraph := range paragraphs {
		if !wrap || strings.TrimSpace(paragraph) == "" {
			lines = append(lines, paragraph)
			continue
		}
		words := strings.Fields(paragraph)
		current := ""
		for _, word := range words {
			candidate := word
			if current != "" {
				candidate = current + " " + word
			}
			if svgTextWidth(candidate, fontSize, bold) <= maxWidth {
				current = candidate
				continue
			}
			if current != "" {
				lines = append(lines, current)
				current = ""
			}
			parts := svgBreakTextToken(word, maxWidth, fontSize, bold)
			if len(parts) > 1 {
				lines = append(lines, parts[:len(parts)-1]...)
			}
			if len(parts) > 0 {
				current = parts[len(parts)-1]
			}
		}
		if current != "" {
			lines = append(lines, current)
		}
	}
	if len(lines) == 0 {
		return []string{""}
	}
	return lines
}

func svgBreakTextToken(value string, maxWidth, fontSize float64, bold bool) []string {
	if value == "" {
		return []string{""}
	}
	parts := []string{}
	var current strings.Builder
	for _, r := range value {
		candidate := current.String() + string(r)
		if current.Len() > 0 && svgTextWidth(candidate, fontSize, bold) > maxWidth {
			parts = append(parts, current.String())
			current.Reset()
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}
	return parts
}

func svgTextWidth(value string, fontSize float64, bold bool) float64 {
	return share.PresentationTextWidth(value, fontSize, bold)
}

func svgTextClipID(op entity.DrawOp, x, y, w, h float64) string {
	key := fmt.Sprintf("%s|%.6f|%.6f|%.6f|%.6f", op.ID, x, y, w, h)
	var hash uint64 = 1469598103934665603
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= 1099511628211
	}
	return fmt.Sprintf("xaligo-text-clip-%x", hash)
}

func writeLine(b *bytes.Buffer, op entity.DrawOp, x, y, w, h, ppi float64) {
	points := absolutePoints(op, x, y, w, h, ppi)
	if len(points) < 2 {
		return
	}
	var d strings.Builder
	for i, p := range points {
		cmd := "L"
		if i == 0 {
			cmd = "M"
		}
		fmt.Fprintf(&d, "%s %s %s ", cmd, num(p.x), num(p.y))
	}
	marker := ""
	if op.Line != nil {
		if id := markerID(op.Line.BeginArrowType); id != "" {
			marker += ` marker-start="url(#` + id + `)"`
		}
		if id := markerID(op.Line.EndArrowType); id != "" {
			marker += ` marker-end="url(#` + id + `)"`
		}
	}
	fmt.Fprintf(b, `<path d="%s" fill="none"%s%s/>`+"\n", strings.TrimSpace(d.String()), lineAttrs(op.Line, ppi), marker)
}

func markerID(arrowType string) string {
	switch arrowType {
	case "", "none":
		return ""
	case "triangle":
		return "xaligo-triangle"
	case "stealth":
		return "xaligo-stealth"
	case "diamond":
		return "xaligo-diamond"
	case "oval":
		return "xaligo-oval"
	case "arrow":
		return "xaligo-arrow"
	default:
		// Plans produced before arrowhead validation treated unknown values as
		// the default arrow marker. Keep that renderer-level compatibility.
		return "xaligo-arrow"
	}
}

func writeMarkerDefinitions(b *bytes.Buffer) {
	b.WriteString(`<defs>` +
		`<marker id="xaligo-arrow" markerWidth="10" markerHeight="10" refX="9" refY="5" orient="auto-start-reverse" markerUnits="strokeWidth"><path d="M 0 0 L 10 5 L 0 10 z" fill="context-stroke"/></marker>` +
		`<marker id="xaligo-triangle" markerWidth="8" markerHeight="10" refX="7" refY="5" orient="auto-start-reverse" markerUnits="strokeWidth"><path d="M 0 0 L 8 5 L 0 10 z" fill="context-stroke"/></marker>` +
		`<marker id="xaligo-stealth" markerWidth="10" markerHeight="10" refX="9" refY="5" orient="auto-start-reverse" markerUnits="strokeWidth"><path d="M 0 0 L 10 5 L 0 10 L 3 5 z" fill="context-stroke"/></marker>` +
		`<marker id="xaligo-diamond" markerWidth="12" markerHeight="8" refX="11" refY="4" orient="auto-start-reverse" markerUnits="strokeWidth"><path d="M 0 4 L 5.5 0 L 11 4 L 5.5 8 z" fill="context-stroke"/></marker>` +
		`<marker id="xaligo-oval" markerWidth="7" markerHeight="7" refX="3.5" refY="3.5" orient="auto" markerUnits="strokeWidth"><circle cx="3.5" cy="3.5" r="2.5" fill="context-stroke"/></marker>` +
		`</defs>` + "\n")
}

type point struct{ x, y float64 }

func absolutePoints(op entity.DrawOp, x, y, w, h, ppi float64) []point {
	if len(op.Points) > 0 {
		out := make([]point, 0, len(op.Points))
		for _, p := range op.Points {
			out = append(out, point{x + p.X*ppi, y + p.Y*ppi})
		}
		return out
	}
	x1, x2 := x, x+w
	y1, y2 := y, y+h
	if op.FlipH {
		x1, x2 = x2, x1
	}
	if op.FlipV {
		y1, y2 = y2, y1
	}
	return []point{{x1, y1}, {x2, y2}}
}

func fillAttrs(fill *entity.FillStyle) string {
	if fill == nil || fill.Transparency >= 100 {
		return ` fill="none"`
	}
	return fmt.Sprintf(` fill="#%s" fill-opacity="%s"`, color(fill.Color, "FFFFFF"), num(opacity(fill.Transparency)))
}

func lineAttrs(line *entity.LineStyle, ppi float64) string {
	if line == nil || line.Transparency >= 100 || line.Width <= 0 {
		return ` stroke="none"`
	}
	dash := ""
	if line.Dash == "dash" {
		dash = ` stroke-dasharray="8 6"`
	} else if line.Dash == "dot" {
		dash = ` stroke-dasharray="2 5" stroke-linecap="round"`
	}
	widthPx := line.Width * ppi / 72.0
	return fmt.Sprintf(` stroke="#%s" stroke-width="%s" stroke-opacity="%s"%s`, color(line.Color, "1E1E1E"), num(widthPx), num(opacity(line.Transparency)), dash)
}

func rotateAttr(deg, cx, cy float64) string {
	if math.Abs(deg) < 0.0001 {
		return ""
	}
	return fmt.Sprintf(` transform="rotate(%s %s %s)"`, num(deg), num(cx), num(cy))
}

func opacity(transparency float64) float64 {
	return math.Max(0, math.Min(1, 1-transparency/100.0))
}

func color(v, fallback string) string {
	v = strings.TrimPrefix(strings.TrimSpace(v), "#")
	if len(v) == 6 {
		return strings.ToUpper(v)
	}
	return fallback
}

func num(v float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", v), "0"), ".")
}

func attr(v string) string {
	var b bytes.Buffer
	_ = xml.EscapeText(&b, []byte(v))
	return b.String()
}

func text(v string) string {
	return attr(v)
}
