package v2

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type FrontendUsecase interface {
	Lower([]byte) (entity.EngineDocumentSpec, string, error)
}

type frontendUsecase struct{}

func NewFrontendUsecase() FrontendUsecase {
	return &frontendUsecase{}
}

type frontendNode struct {
	tag      string
	attrs    map[string]string
	text     string
	position entity.Position
	children []*frontendNode
}

func (rcvr *frontendUsecase) Lower(source []byte) (entity.EngineDocumentSpec, string, error) {
	root, err := parseFrontendDocument(source)
	if err != nil {
		return entity.EngineDocumentSpec{}, "", err
	}
	version := strings.TrimSpace(root.attrs["version"])
	if version == "" {
		version = "1"
	}
	if version != "1" && version != "2" {
		return entity.EngineDocumentSpec{}, "", fmt.Errorf("unsupported XAL document version %q", version)
	}
	content := root
	if root.tag == "xaligo" {
		content = firstFrontendContent(root)
		if content == nil {
			return entity.EngineDocumentSpec{}, version, fmt.Errorf("<xaligo version=%q> has no renderable content", version)
		}
	}
	width, err := frontendNumber(content, "width", 1280)
	if err != nil {
		return entity.EngineDocumentSpec{}, version, err
	}
	height, err := frontendNumber(content, "height", 720)
	if err != nil {
		return entity.EngineDocumentSpec{}, version, err
	}
	gap, err := frontendNumber(content, "gap", 16)
	if err != nil {
		return entity.EngineDocumentSpec{}, version, err
	}
	state := frontendLowerState{source: source, nextID: 1, nextSpanID: 1}
	elements := make([]entity.EngineElementSpec, 0, len(content.children))
	if conceptForFrontendTag(content.tag, len(content.children)) == entity.EngineConceptFrame {
		element, lowerErr := state.lower(content)
		if lowerErr != nil {
			return entity.EngineDocumentSpec{}, version, lowerErr
		}
		elements = append(elements, element)
	} else {
		for _, child := range content.children {
			element, lowerErr := state.lower(child)
			if lowerErr != nil {
				return entity.EngineDocumentSpec{}, version, lowerErr
			}
			elements = append(elements, element)
		}
	}
	if len(elements) == 0 {
		return entity.EngineDocumentSpec{}, version, fmt.Errorf("XAL document has no calculable elements")
	}
	return entity.EngineDocumentSpec{
		Layout: engineLayoutForFrontendNode(content), Width: width, Height: height,
		Gap: gap, Overflow: engineOverflowForFrontendNode(content), Elements: elements,
		Spans: state.spans,
	}, version, nil
}

type frontendLowerState struct {
	source     []byte
	nextID     int
	nextSpanID uint32
	spans      []entity.EngineSourceSpan
}

func (rcvr *frontendLowerState) lower(node *frontendNode) (entity.EngineElementSpec, error) {
	spanID := rcvr.nextSpanID
	rcvr.nextSpanID++
	rcvr.spans = append(rcvr.spans, entity.EngineSourceSpan{
		ID: spanID, Offset: node.position.Offset, Line: node.position.Line,
		Column: node.position.Column, Length: len(node.tag) + 2,
	})
	id := strings.TrimSpace(node.attrs["id"])
	if id == "" {
		id = fmt.Sprintf("v2-%s-%d", node.tag, rcvr.nextID)
	}
	rcvr.nextID++
	element := entity.EngineElementSpec{
		ID: id, Concept: conceptForFrontendTag(node.tag, len(node.children)), SpanID: spanID,
		Layout: engineLayoutForFrontendNode(node), Overflow: engineOverflowForFrontendNode(node),
		Visual:  entity.EngineVisualSpec{Shape: engineShapeForFrontendNode(node)},
		Sources: []entity.EngineParameterSource{{Parameter: "concept", Origin: "profile", SpanID: spanID}},
	}
	var err error
	if element.X, err = frontendOptionalNumber(node, "x"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Y, err = frontendOptionalNumber(node, "y"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Width, err = frontendOptionalNumber(node, "width"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Height, err = frontendOptionalNumber(node, "height"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Concept == entity.EngineConceptItem && len(node.children) == 0 && strings.TrimSpace(node.attrs["weight"]) == "" && strings.TrimSpace(node.attrs["span"]) == "" {
		const compatibilityItemSize = 72.0
		if element.Width == nil {
			width := compatibilityItemSize
			element.Width = &width
		}
		if element.Height == nil {
			height := compatibilityItemSize
			element.Height = &height
		}
	}
	if element.Gap, err = frontendOptionalNumber(node, "gap"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Weight, err = frontendOptionalNumber(node, "weight"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Weight == nil {
		if element.Weight, err = frontendOptionalNumber(node, "span"); err != nil {
			return entity.EngineElementSpec{}, err
		}
	}
	if element.OffsetX, err = frontendOptionalNumber(node, "dx"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.OffsetY, err = frontendOptionalNumber(node, "dy"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	element.Align, element.Justify = frontendAlignment(node.attrs["align"])
	element.Padding = frontendClassPadding(node.attrs["class"])
	if row, rowErr := frontendOptionalNumber(node, "row"); rowErr != nil {
		return entity.EngineElementSpec{}, rowErr
	} else if row != nil && *row > 1 {
		top := (*row - 1) * 4
		element.Margin.Top = &top
	}
	if raw := strings.TrimSpace(node.attrs["fill"]); raw != "" {
		element.Visual.Fill = raw
	}
	if raw := strings.TrimSpace(node.attrs["stroke"]); raw != "" {
		element.Visual.Stroke = raw
	}
	if element.Visual.Stroke == "" {
		element.Visual.Stroke = strings.TrimSpace(node.attrs["color"])
	}
	if element.Visual.StrokeWidth, err = frontendOptionalNumber(node, "stroke-width"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Visual.CornerRadius, err = frontendOptionalNumber(node, "corner-radius"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Visual.Opacity, err = frontendOptionalNumber(node, "opacity"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	text := strings.TrimSpace(node.text)
	if label := firstNonEmpty(node.attrs["label"], node.attrs["title"], node.attrs["name"]); label != "" {
		text = label
	}
	if text != "" {
		element.Text = &entity.EngineTextSpec{Value: text, Role: "label"}
		if element.Text.FontSize, err = frontendOptionalNumber(node, "font-size"); err != nil {
			return entity.EngineElementSpec{}, err
		}
	}
	if element.Concept == entity.EngineConceptPort {
		element.Port = &entity.EnginePortSpec{Side: entity.EngineSide(node.attrs["side"]), Label: firstNonEmpty(node.attrs["label"], node.attrs["title"], node.attrs["name"])}
	}
	if element.Concept == entity.EngineConceptLine {
		element.Line = &entity.EngineLineSpec{
			Source:  firstNonEmpty(node.attrs["source"], node.attrs["src"]),
			Target:  firstNonEmpty(node.attrs["target"], node.attrs["dst"]),
			Routing: entity.EngineRoutingPolicy(firstNonEmpty(node.attrs["routing"], node.attrs["route"])),
			Label:   firstNonEmpty(node.attrs["label"], node.attrs["title"], node.attrs["name"]),
		}
		if element.Line.Routing == "" {
			element.Line.Routing = entity.EngineRoutingOrthogonal
		}
		element.Line.Style = engineLineStyleForFrontendNode(node)
		element.Line.SourceSide, element.Line.SourceAnchor = frontendConnectionAnchor(node.attrs["src-anchor"])
		element.Line.TargetSide, element.Line.TargetAnchor = frontendConnectionAnchor(node.attrs["dst-anchor"])
		element.Line.SourceDecoration = engineDecorationForFrontendValue(firstNonEmpty(node.attrs["src-arrow"], node.attrs["source-arrow"]))
		element.Line.TargetDecoration = engineDecorationForFrontendValue(firstNonEmpty(node.attrs["dst-arrow"], node.attrs["target-arrow"], node.attrs["arrow"]))
	}
	if iconRef := frontendIconRef(node); iconRef != "" {
		element.Icon = &entity.EngineIconSpec{Ref: iconRef, MissingPolicy: entity.EngineIconMissingFallback}
	}
	for _, child := range node.children {
		if child.tag == "connections" {
			for _, connection := range child.children {
				lowered, lowerErr := rcvr.lower(connection)
				if lowerErr != nil {
					return entity.EngineElementSpec{}, lowerErr
				}
				element.Children = append(element.Children, lowered)
			}
			continue
		}
		lowered, lowerErr := rcvr.lower(child)
		if lowerErr != nil {
			return entity.EngineElementSpec{}, lowerErr
		}
		element.Children = append(element.Children, lowered)
	}
	return element, nil
}

func frontendAlignment(value string) (entity.EngineAlignment, entity.EngineJustification) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "middle", "center":
		return entity.EngineAlignCenter, entity.EngineJustifyCenter
	case "middle-spread", "center-spread":
		return entity.EngineAlignCenter, entity.EngineJustifySpaceBetween
	case "end", "right", "bottom":
		return entity.EngineAlignEnd, entity.EngineJustifyEnd
	default:
		return entity.EngineAlignStart, entity.EngineJustifyStart
	}
}

func frontendClassPadding(value string) entity.EngineInsets {
	for _, className := range strings.Fields(value) {
		if strings.HasPrefix(className, "pa-") {
			level, err := strconv.ParseFloat(strings.TrimPrefix(className, "pa-"), 64)
			if err == nil && level >= 0 {
				padding := level * 4
				return entity.EngineInsets{Top: &padding, Right: &padding, Bottom: &padding, Left: &padding}
			}
		}
	}
	return entity.EngineInsets{}
}

func frontendConnectionAnchor(value string) (entity.EngineSide, *float64) {
	parts := strings.Split(strings.ToLower(strings.TrimSpace(value)), "-")
	if len(parts) == 0 || parts[0] == "" {
		return entity.EngineSideAuto, nil
	}
	side := entity.EngineSide(parts[0])
	if len(parts) == 1 {
		return side, nil
	}
	index, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil || index < 1 {
		return side, nil
	}
	anchor := math.Min(0.9, float64(index)/4.0)
	return side, &anchor
}

func parseFrontendDocument(source []byte) (*frontendNode, error) {
	decoder := xml.NewDecoder(bytes.NewReader(source))
	var stack []*frontendNode
	var root *frontendNode
	for {
		offset := int(decoder.InputOffset())
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("parse XAL frontend: %w", err)
		}
		switch value := token.(type) {
		case xml.StartElement:
			node := &frontendNode{tag: value.Name.Local, attrs: make(map[string]string), position: frontendPosition(source, offset)}
			for _, attr := range value.Attr {
				node.attrs[attr.Name.Local] = attr.Value
			}
			if len(stack) == 0 {
				root = node
			} else {
				stack[len(stack)-1].children = append(stack[len(stack)-1].children, node)
			}
			stack = append(stack, node)
		case xml.CharData:
			if len(stack) > 0 {
				stack[len(stack)-1].text += string(value)
			}
		case xml.EndElement:
			if len(stack) == 0 {
				return nil, fmt.Errorf("unexpected closing tag </%s>", value.Name.Local)
			}
			stack = stack[:len(stack)-1]
		}
	}
	if root == nil {
		return nil, fmt.Errorf("empty XAL document")
	}
	return root, nil
}

func firstFrontendContent(root *frontendNode) *frontendNode {
	for _, child := range root.children {
		if child.tag == "data" {
			continue
		}
		if child.tag == "frames" && len(child.children) == 1 {
			return child.children[0]
		}
		return child
	}
	return nil
}

func conceptForFrontendTag(tag string, childCount int) entity.EngineConcept {
	switch tag {
	case "frame", "frames":
		return entity.EngineConceptFrame
	case "capture":
		return entity.EngineConceptCapture
	case "port":
		return entity.EngineConceptPort
	case "connection", "line", "route", "traffic":
		return entity.EngineConceptLine
	case "text", "label":
		return entity.EngineConceptText
	case "spacer", "blank":
		return entity.EngineConceptSpacer
	case "item", "rectangle":
		return entity.EngineConceptItem
	default:
		if childCount > 0 {
			return entity.EngineConceptGroup
		}
		return entity.EngineConceptItem
	}
}

func engineLayoutForFrontendNode(node *frontendNode) entity.EngineLayoutPolicy {
	if raw := strings.TrimSpace(node.attrs["layout"]); raw != "" {
		return entity.EngineLayoutPolicy(raw)
	}
	switch node.tag {
	case "row":
		return entity.EngineLayoutHorizontal
	case "col", "container", "frame", "frames":
		return entity.EngineLayoutVertical
	case "connections":
		return entity.EngineLayoutNone
	default:
		if len(node.children) > 0 {
			if strings.Contains(strings.ToLower(node.attrs["align"]), "spread") {
				return entity.EngineLayoutHorizontal
			}
			return entity.EngineLayoutVertical
		}
		return entity.EngineLayoutNone
	}
}

func engineLineStyleForFrontendNode(node *frontendNode) entity.EngineLineStyle {
	switch strings.ToLower(firstNonEmpty(node.attrs["stroke-style"], node.attrs["style"])) {
	case "dashed", "dash":
		return entity.EngineLineDashed
	case "dotted", "dot":
		return entity.EngineLineDotted
	default:
		return entity.EngineLineSolid
	}
}

func engineDecorationForFrontendValue(value string) entity.EngineDecoration {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "arrow", "thin", "standard", "stealth":
		return entity.EngineDecorationArrow
	case "triangle":
		return entity.EngineDecorationTriangle
	case "diamond":
		return entity.EngineDecorationDiamond
	case "circle", "oval":
		return entity.EngineDecorationCircle
	default:
		return entity.EngineDecorationNone
	}
}

func frontendIconRef(node *frontendNode) string {
	if ref := firstNonEmpty(node.attrs["icon"], node.attrs["icon-ref"]); ref != "" {
		return ref
	}
	if iconID := strings.TrimSpace(node.attrs["icon-id"]); iconID != "" {
		return "catalog:" + iconID
	}
	if node.tag == "item" {
		if id := strings.TrimSpace(node.attrs["id"]); id != "" {
			return "catalog:" + id
		}
	}
	return ""
}

func engineOverflowForFrontendNode(node *frontendNode) entity.EngineOverflow {
	if strings.TrimSpace(node.attrs["overflow"]) == "visible" {
		return entity.EngineOverflowVisible
	}
	if len(node.children) > 0 && node.tag != "xaligo" && node.tag != "frame" && node.tag != "frames" {
		return entity.EngineOverflowVisible
	}
	return entity.EngineOverflowError
}

func engineShapeForFrontendNode(node *frontendNode) entity.EngineShape {
	switch strings.TrimSpace(node.attrs["shape"]) {
	case "ellipse":
		return entity.EngineShapeEllipse
	case "none":
		return entity.EngineShapeNone
	default:
		if node.tag == "spacer" || node.tag == "blank" || node.tag == "line" || node.tag == "connection" {
			return entity.EngineShapeNone
		}
		return entity.EngineShapeRectangle
	}
}

func frontendOptionalNumber(node *frontendNode, name string) (*float64, error) {
	raw, ok := node.attrs[name]
	if !ok || strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	value, err := strconv.ParseFloat(strings.TrimSpace(raw), 64)
	if err != nil {
		return nil, fmt.Errorf("<%s> %s=%q is not numeric", node.tag, name, raw)
	}
	return &value, nil
}

func frontendNumber(node *frontendNode, name string, fallback float64) (float64, error) {
	value, err := frontendOptionalNumber(node, name)
	if err != nil {
		return 0, err
	}
	if value == nil {
		return fallback, nil
	}
	return *value, nil
}

func frontendPosition(source []byte, offset int) entity.Position {
	if offset < 0 {
		offset = 0
	}
	if offset > len(source) {
		offset = len(source)
	}
	prefix := source[:offset]
	line := bytes.Count(prefix, []byte{'\n'}) + 1
	last := bytes.LastIndexByte(prefix, '\n')
	return entity.Position{Offset: offset, Line: line, Column: offset - last}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
