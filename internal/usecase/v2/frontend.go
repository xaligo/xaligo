package v2

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
)

type FrontendUsecase interface {
	Lower([]byte) (entity.EngineDocumentSpec, string, error)
	LowerWithProvenance([]byte) (entity.EngineDocumentSpec, string, error)
}

type frontendUsecase struct{}

func NewFrontendUsecase() FrontendUsecase {
	return &frontendUsecase{}
}

type frontendNode struct {
	tag      string
	path     string
	attrs    map[string]string
	text     string
	position entity.Position
	children []*frontendNode
}

func (rcvr *frontendUsecase) Lower(source []byte) (entity.EngineDocumentSpec, string, error) {
	return rcvr.lowerDocument(source, false)
}

func (rcvr *frontendUsecase) LowerWithProvenance(source []byte) (entity.EngineDocumentSpec, string, error) {
	return rcvr.lowerDocument(source, true)
}

func (rcvr *frontendUsecase) lowerDocument(source []byte, preserveProvenance bool) (entity.EngineDocumentSpec, string, error) {
	root, err := parseFrontendDocument(source)
	if err != nil {
		return entity.EngineDocumentSpec{}, "", err
	}
	version, err := frontendDocumentVersion(root)
	if err != nil {
		return entity.EngineDocumentSpec{}, "", err
	}
	if err := validateFrontendAWSBoundaryAttachments(root); err != nil {
		return entity.EngineDocumentSpec{}, version, err
	}
	if preserveProvenance {
		assignFrontendSourcePaths(root, "", 0)
	}
	content := root
	if root.tag == "xaligo" {
		content, err = firstFrontendContent(root, version)
		if err != nil {
			return entity.EngineDocumentSpec{}, version, err
		}
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
	var padding entity.EngineInsets
	var columns *uint16
	if conceptForFrontendTag(content.tag, len(content.children)) != entity.EngineConceptFrame {
		padding, err = frontendInsets(content, "padding", frontendUsesV1AuthoringProfile(version))
		if err != nil {
			return entity.EngineDocumentSpec{}, version, err
		}
		columns, err = frontendOptionalUint16(content, "columns")
		if err != nil {
			return entity.EngineDocumentSpec{}, version, err
		}
	}
	state := frontendLowerState{source: source, version: version, preserveProvenance: preserveProvenance, nextID: 1, nextSpanID: 1}
	elements := make([]entity.EngineElementSpec, 0, len(content.children))
	if conceptForFrontendTag(content.tag, len(content.children)) == entity.EngineConceptFrame {
		element, lowerErr := state.lower(content, frontendDefaults{}, nil)
		if lowerErr != nil {
			return entity.EngineDocumentSpec{}, version, lowerErr
		}
		elements = append(elements, element)
	} else {
		defaults, defaultsErr := frontendDefaultsForNode(content, frontendDefaults{})
		if defaultsErr != nil {
			return entity.EngineDocumentSpec{}, version, defaultsErr
		}
		for _, child := range content.children {
			element, lowerErr := state.lower(child, defaults, nil)
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
		Gap: gap, Padding: padding, Columns: columns,
		Overflow: engineOverflowForFrontendNode(content), Elements: elements,
		Spans: state.spans,
	}, version, nil
}

type frontendLowerState struct {
	source             []byte
	version            string
	preserveProvenance bool
	nextID             int
	nextSpanID         uint32
	spans              []entity.EngineSourceSpan
}

type frontendDefaults struct {
	itemSize *float64
}

func (rcvr *frontendLowerState) lower(node *frontendNode, inherited frontendDefaults, portAnchor *float64) (entity.EngineElementSpec, error) {
	defaults, err := frontendDefaultsForNode(node, inherited)
	if err != nil {
		return entity.EngineElementSpec{}, err
	}
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
	if rcvr.preserveProvenance {
		element.Provenance = frontendElementProvenance(node)
	}
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
	if element.Height != nil {
		element.Sources = append(element.Sources, entity.EngineParameterSource{Parameter: "height", Origin: "explicit", SpanID: spanID})
	}
	if element.Gap, err = frontendOptionalNumber(node, "gap"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Weight, err = frontendOptionalNumber(node, "weight"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Weight == nil && frontendUsesV1AuthoringProfile(rcvr.version) {
		if element.Weight, err = frontendOptionalNumber(node, "row"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Weight == nil {
			if element.Weight, err = frontendOptionalNumber(node, "col"); err != nil {
				return entity.EngineElementSpec{}, err
			}
		}
	}
	if element.Weight == nil && frontendUsesV1AuthoringProfile(rcvr.version) && node.tag == "col" {
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
	if raw := strings.TrimSpace(node.attrs["justify"]); raw != "" {
		element.Justify = frontendJustification(raw)
	}
	if element.Margin, err = frontendInsets(node, "margin", frontendUsesV1AuthoringProfile(rcvr.version)); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Padding, err = frontendInsets(node, "padding", frontendUsesV1AuthoringProfile(rcvr.version)); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Columns, err = frontendOptionalUint16(node, "columns"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.ColumnSpan, err = frontendOptionalUint16(node, "column-span"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.RowSpan, err = frontendOptionalUint16(node, "row-span"); err != nil {
		return entity.EngineElementSpec{}, err
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
	if element.Visual.Visible, err = frontendOptionalBool(node, "visible"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	if element.Visual.Layer, err = frontendOptionalInt32(node, "layer"); err != nil {
		return entity.EngineElementSpec{}, err
	}
	applyFrontendV1GeometryProfile(node, &element)
	applyFrontendV1VisualProfile(node, &element)
	text := strings.TrimSpace(node.text)
	catalogFallbackLabel := false
	if label := frontendLabel(node, element.Concept, rcvr.version); label != "" {
		text = label
	}
	if text == "" && frontendUsesV1AuthoringProfile(rcvr.version) && node.tag == "item" {
		if _, parseErr := strconv.Atoi(strings.TrimSpace(node.attrs["id"])); parseErr == nil {
			text = strings.TrimSpace(node.attrs["name"])
			catalogFallbackLabel = text != ""
		}
	}
	if text != "" {
		role := "label"
		if catalogFallbackLabel {
			role = "catalog-label"
		}
		element.Text = &entity.EngineTextSpec{Value: text, Role: role}
		if element.Text.FontSize, err = frontendOptionalNumber(node, "font-size"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		applyFrontendV1TextProfile(node, &element)
	}
	if element.Concept == entity.EngineConceptPort {
		portLabel := firstNonEmpty(node.attrs["label"], node.attrs["title"])
		if frontendUsesV1AuthoringProfile(rcvr.version) {
			portLabel = firstNonEmpty(portLabel, node.attrs["name"])
		}
		defaultSide := string(entity.EngineSideAuto)
		if definition, ok := frontendAWSBoundaryAttachment(node.tag); ok {
			defaultSide = definition.DefaultSide
			portLabel = ""
		}
		element.Port = &entity.EnginePortSpec{Side: entity.EngineSide(firstNonEmpty(node.attrs["side"], defaultSide)), Anchor: portAnchor, Label: portLabel}
		if element.Port.Anchor, err = frontendOptionalNumberFallback(node, "anchor", element.Port.Anchor); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Port.Offset, err = frontendOptionalNumber(node, "offset"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Port.Size, err = frontendOptionalNumber(node, "size"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Port.Visible, err = frontendOptionalBool(node, "port-visible"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		applyFrontendAWSBoundaryPortProfile(node, &element)
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
		if element.Line.SourceSide, element.Line.SourceAnchor, err = frontendLineEndpoint(node, "source", "src", frontendUsesV1AuthoringProfile(rcvr.version)); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Line.TargetSide, element.Line.TargetAnchor, err = frontendLineEndpoint(node, "target", "dst", frontendUsesV1AuthoringProfile(rcvr.version)); err != nil {
			return entity.EngineElementSpec{}, err
		}
		element.Line.SourceDecoration = engineDecorationForFrontendValue(firstNonEmpty(node.attrs["source-decoration"], node.attrs["src-arrow"], node.attrs["source-arrow"]))
		element.Line.TargetDecoration = engineDecorationForFrontendValue(firstNonEmpty(node.attrs["target-decoration"], node.attrs["dst-arrow"], node.attrs["target-arrow"], node.attrs["arrow"]))
		if frontendUsesV1AuthoringProfile(rcvr.version) && strings.EqualFold(strings.TrimSpace(node.attrs["kind"]), "traffic") && element.Line.TargetDecoration == entity.EngineDecorationNone {
			element.Line.TargetDecoration = entity.EngineDecorationArrow
		}
	}
	if iconRef := frontendIconRef(node, rcvr.version); iconRef != "" {
		element.Icon = &entity.EngineIconSpec{Ref: iconRef, MissingPolicy: entity.EngineIconMissingFallback}
		if element.Concept == entity.EngineConceptItem && strings.TrimSpace(node.attrs["shape"]) == "" {
			element.Visual.Shape = entity.EngineShapeNone
		}
		if element.Icon.Width, err = frontendOptionalNumber(node, "icon-width"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Icon.Height, err = frontendOptionalNumber(node, "icon-height"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Icon.Scale, err = frontendOptionalNumber(node, "icon-scale"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Icon.OffsetX, err = frontendOptionalNumber(node, "icon-offset-x"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		if element.Icon.OffsetY, err = frontendOptionalNumber(node, "icon-offset-y"); err != nil {
			return entity.EngineElementSpec{}, err
		}
		applyFrontendV1IconProfile(node, &element)
		if element.Concept == entity.EngineConceptItem && defaults.itemSize != nil && !awsprofile.IsResourceTag(node.tag) {
			if element.Icon.Width == nil {
				element.Icon.Width = defaults.itemSize
			}
			if element.Icon.Height == nil {
				element.Icon.Height = defaults.itemSize
			}
			if element.Text != nil {
				frontendSetNumberDefault(&element.Text.FontSize, frontendV1ItemLabelFontSize)
				frontendSetNumberDefault(&element.Text.LineHeight, frontendV1ItemLabelLineHeight)
			}
			if element.Width == nil {
				width := math.Max(frontendV1ItemLabelWidth, *defaults.itemSize+frontendV1ItemVisualPad*2)
				element.Width = &width
			}
			if element.Height == nil {
				height := *defaults.itemSize
				if element.Text != nil {
					height += frontendV1ItemLabelGap + frontendV1ItemTextHeight(element.Text)
				}
				element.Height = &height
			}
		}
	}
	if err := applyFrontendAWSResource(node, &element); err != nil {
		return entity.EngineElementSpec{}, err
	}
	var portAnchors map[*frontendNode]*float64
	if frontendUsesV1AuthoringProfile(rcvr.version) {
		portAnchors = frontendPortAnchors(node.children)
	}
	for _, child := range node.children {
		if element.Concept == entity.EngineConceptFrame && child.tag == "metadata" {
			continue
		}
		if child.tag == "connections" {
			for _, connection := range child.children {
				lowered, lowerErr := rcvr.lower(connection, defaults, nil)
				if lowerErr != nil {
					return entity.EngineElementSpec{}, lowerErr
				}
				element.Children = append(element.Children, lowered)
			}
			continue
		}
		lowered, lowerErr := rcvr.lower(child, defaults, portAnchors[child])
		if lowerErr != nil {
			return entity.EngineElementSpec{}, lowerErr
		}
		element.Children = append(element.Children, lowered)
	}
	if frontendUsesV1AuthoringProfile(rcvr.version) {
		applyFrontendV1FrameMetadataComposition(node, &element)
	}
	return element, nil
}

func frontendAlignment(value string) (entity.EngineAlignment, entity.EngineJustification) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "middle", "middle-center":
		return entity.EngineAlignCenter, entity.EngineJustifyCenter
	case "center":
		return entity.EngineAlignCenter, entity.EngineJustifyStart
	case "middle-spread", "center-spread":
		return entity.EngineAlignCenter, entity.EngineJustifySpaceEvenly
	case "stretch":
		return entity.EngineAlignStretch, entity.EngineJustifyStart
	case "end", "right", "bottom":
		return entity.EngineAlignEnd, entity.EngineJustifyEnd
	default:
		return entity.EngineAlignStart, entity.EngineJustifyStart
	}
}

func frontendJustification(value string) entity.EngineJustification {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "center", "middle":
		return entity.EngineJustifyCenter
	case "end", "right", "bottom":
		return entity.EngineJustifyEnd
	case "space-between", "between":
		return entity.EngineJustifySpaceBetween
	case "space-evenly", "evenly", "spread":
		return entity.EngineJustifySpaceEvenly
	default:
		return entity.EngineJustifyStart
	}
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
	anchor := math.Min(0.9, math.Max(0.1, (float64(index)*2.0-1.0)/10.0))
	return side, &anchor
}

func frontendDocumentVersion(root *frontendNode) (string, error) {
	version := strings.TrimSpace(root.attrs["version"])
	switch root.tag {
	case "scene":
		return "", fmt.Errorf("unsupported XAL document root <scene>; use <xaligo version=\"2\">")
	case "xaligo":
		if version == "" {
			return "1", nil
		}
		if version != "1" && version != "2" {
			return "", fmt.Errorf("<xaligo> version must be \"1\" or \"2\", got %q", version)
		}
		return version, nil
	case "frame", "frames":
		if version != "" && version != "1" {
			return "", fmt.Errorf("legacy <%s> root accepts only version=\"1\"", root.tag)
		}
		return "1", nil
	default:
		return "", fmt.Errorf("unsupported XAL document root <%s>", root.tag)
	}
}

func frontendUsesV1AuthoringProfile(version string) bool {
	return version == "1" || version == "2"
}

func frontendDefaultsForNode(node *frontendNode, inherited frontendDefaults) (frontendDefaults, error) {
	defaults := inherited
	itemSize, err := frontendOptionalNumber(node, "item-size")
	if err != nil {
		return frontendDefaults{}, err
	}
	if itemSize != nil {
		defaults.itemSize = itemSize
	}
	return defaults, nil
}

func frontendLabel(node *frontendNode, concept entity.EngineConcept, version string) string {
	if concept == entity.EngineConceptLine {
		return ""
	}
	if _, ok := frontendAWSBoundaryAttachment(node.tag); ok {
		return ""
	}
	if definition, ok := awsprofile.DefinitionForTag(node.tag); ok {
		if definition.Group == nil {
			return definition.Label(node.attrs)
		}
		for _, parameter := range definition.Parameters {
			if node.attrs[parameter.Name] != "" {
				return definition.Label(node.attrs)
			}
		}
		if node.attrs["detail"] != "" {
			return definition.Label(node.attrs)
		}
	}
	if label := firstNonEmpty(node.attrs["label"], node.attrs["title"]); label != "" {
		return label
	}
	if concept == entity.EngineConceptText {
		return strings.TrimSpace(node.attrs["name"])
	}
	return ""
}

func frontendElementProvenance(node *frontendNode) *entity.EngineElementProvenance {
	identity := firstNonEmpty(node.attrs["id"], node.attrs["name"], node.attrs["ref"])
	name := firstNonEmpty(node.attrs["title"], node.attrs["label"], node.attrs["name"], node.attrs["ref"], node.attrs["id"])
	text := strings.TrimSpace(node.text)
	if name == "" {
		name = text
	}
	keys := make([]string, 0, len(node.attrs))
	for key := range node.attrs {
		if !strings.HasPrefix(key, "_") {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys)+2)
	parts = append(parts, node.tag)
	for _, key := range keys {
		if value := strings.TrimSpace(node.attrs[key]); value != "" {
			parts = append(parts, key+"="+value)
		}
	}
	if text != "" {
		parts = append(parts, text)
	}
	return &entity.EngineElementProvenance{
		Tag: node.tag, Path: node.path, Identity: identity, Name: name, Detail: strings.Join(parts, " "),
		SourceRef: firstNonEmpty(node.attrs["source"], node.attrs["src"]),
		TargetRef: firstNonEmpty(node.attrs["target"], node.attrs["dst"]),
		Position:  node.position,
	}
}

func assignFrontendSourcePaths(node *frontendNode, parent string, sibling int) {
	if node == nil {
		return
	}
	identity := firstNonEmpty(node.attrs["id"], node.attrs["name"], node.attrs["ref"])
	segment := fmt.Sprintf("%s[%d]", node.tag, sibling)
	if identity != "" {
		segment = node.tag + "#" + identity
	}
	node.path = segment
	if parent != "" {
		node.path = parent + "/" + segment
	}
	for index, child := range node.children {
		assignFrontendSourcePaths(child, node.path, index)
	}
}

func frontendLineEndpoint(node *frontendNode, canonical, compatibility string, allowCompatibility bool) (entity.EngineSide, *float64, error) {
	side := entity.EngineSide(firstNonEmpty(node.attrs[canonical+"-side"], string(entity.EngineSideAuto)))
	raw := strings.TrimSpace(node.attrs[canonical+"-anchor"])
	if raw == "" && allowCompatibility {
		raw = strings.TrimSpace(node.attrs[compatibility+"-anchor"])
	}
	if raw == "" {
		return side, nil, nil
	}
	if value, err := strconv.ParseFloat(raw, 64); err == nil {
		return side, &value, nil
	}
	if !allowCompatibility {
		return entity.EngineSideAuto, nil, fmt.Errorf("<%s> %s-anchor=%q is not numeric", node.tag, canonical, raw)
	}
	aliasSide, anchor := frontendConnectionAnchor(raw)
	if anchor == nil {
		return entity.EngineSideAuto, nil, fmt.Errorf("<%s> %s-anchor=%q is not numeric or a supported side anchor", node.tag, canonical, raw)
	}
	return aliasSide, anchor, nil
}

func frontendPortAnchors(children []*frontendNode) map[*frontendNode]*float64 {
	var bySide map[string][]*frontendNode
	for _, child := range children {
		if !isFrontendPortTag(child.tag) || strings.TrimSpace(child.attrs["anchor"]) != "" {
			continue
		}
		if bySide == nil {
			bySide = make(map[string][]*frontendNode)
		}
		defaultSide := string(entity.EngineSideAuto)
		if definition, ok := frontendAWSBoundaryAttachment(child.tag); ok {
			defaultSide = definition.DefaultSide
		}
		side := strings.ToLower(firstNonEmpty(child.attrs["side"], defaultSide))
		bySide[side] = append(bySide[side], child)
	}
	if len(bySide) == 0 {
		return nil
	}
	anchors := make(map[*frontendNode]*float64)
	for _, ports := range bySide {
		for index, port := range ports {
			anchor := float64(index+1) / float64(len(ports)+1)
			anchors[port] = &anchor
		}
	}
	return anchors
}

func parseFrontendDocument(source []byte) (*frontendNode, error) {
	decoder := xml.NewDecoder(bytes.NewReader(source))
	positions := frontendPositionTracker{source: source, line: 1}
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
			node := &frontendNode{tag: value.Name.Local, attrs: make(map[string]string), position: positions.at(offset)}
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

func firstFrontendContent(root *frontendNode, version string) (*frontendNode, error) {
	var content *frontendNode
	for _, child := range root.children {
		if child.tag == "data" {
			continue
		}
		if content != nil {
			return nil, fmt.Errorf("<xaligo version=%q> currently requires exactly one renderable content root", version)
		}
		content = child
	}
	if content != nil && content.tag == "frames" {
		if len(content.children) != 1 {
			return nil, fmt.Errorf("<xaligo version=%q> currently requires exactly one <frame>, got %d", version, len(content.children))
		}
		return content.children[0], nil
	}
	return content, nil
}

func conceptForFrontendTag(tag string, childCount int) entity.EngineConcept {
	if _, ok := frontendAWSBoundaryAttachment(tag); ok {
		return entity.EngineConceptPort
	}
	if definition, ok := awsprofile.DefinitionForTag(tag); ok {
		if definition.Group != nil {
			return entity.EngineConceptGroup
		}
		return entity.EngineConceptItem
	}
	switch tag {
	case "frame", "frames":
		return entity.EngineConceptFrame
	case "group", "row", "col", "container", "connections":
		return entity.EngineConceptGroup
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
			if frontendHasOnlyItemChildren(node) {
				return entity.EngineLayoutAdaptiveGrid
			}
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

func frontendIconRef(node *frontendNode, version string) string {
	if definition, ok := frontendAWSBoundaryAttachment(node.tag); ok {
		return "catalog:" + strconv.Itoa(definition.CatalogID)
	}
	if awsprofile.IsResourceTag(node.tag) {
		definition, _ := awsprofile.DefinitionForTag(node.tag)
		return "catalog:" + strconv.Itoa(definition.CatalogID)
	}
	if ref := firstNonEmpty(node.attrs["icon"], node.attrs["icon-ref"]); ref != "" {
		return ref
	}
	if frontendUsesV1AuthoringProfile(version) {
		if iconID := strings.TrimSpace(node.attrs["icon-id"]); iconID != "" {
			return "catalog:" + iconID
		}
	}
	if frontendUsesV1AuthoringProfile(version) && node.tag == "item" {
		id := strings.TrimSpace(node.attrs["id"])
		if _, err := strconv.Atoi(id); err == nil {
			return "catalog:" + id
		}
	}
	if frontendUsesV1AuthoringProfile(version) {
		if profile, ok := frontendV1GroupProfiles[node.tag]; ok && profile.icon != "" {
			return "group:" + profile.icon
		}
	}
	return ""
}

func engineOverflowForFrontendNode(node *frontendNode) entity.EngineOverflow {
	if strings.TrimSpace(node.attrs["overflow"]) == "visible" {
		return entity.EngineOverflowVisible
	}
	return entity.EngineOverflowError
}

func engineShapeForFrontendNode(node *frontendNode) entity.EngineShape {
	if _, ok := frontendAWSBoundaryAttachment(node.tag); ok {
		return entity.EngineShapeNone
	}
	switch strings.TrimSpace(node.attrs["shape"]) {
	case "ellipse":
		return entity.EngineShapeEllipse
	case "none":
		return entity.EngineShapeNone
	default:
		if node.tag == "spacer" || node.tag == "blank" || node.tag == "line" || node.tag == "connection" || node.tag == "row" || node.tag == "col" || node.tag == "container" || node.tag == "connections" {
			return entity.EngineShapeNone
		}
		return entity.EngineShapeRectangle
	}
}

func frontendInsets(node *frontendNode, prefix string, includeClasses bool) (entity.EngineInsets, error) {
	var insets entity.EngineInsets
	if includeClasses {
		for _, className := range strings.Fields(node.attrs["class"]) {
			if err := applyFrontendSpacingClass(&insets, prefix, className); err != nil {
				return entity.EngineInsets{}, fmt.Errorf("<%s> class %q: %w", node.tag, className, err)
			}
		}
	}
	all, err := frontendOptionalNumber(node, prefix)
	if err != nil {
		return entity.EngineInsets{}, err
	}
	if all != nil {
		insets.Top, insets.Right, insets.Bottom, insets.Left = all, all, all, all
	}
	x, err := frontendOptionalNumber(node, prefix+"-x")
	if err != nil {
		return entity.EngineInsets{}, err
	}
	if x != nil {
		insets.Left, insets.Right = x, x
	}
	y, err := frontendOptionalNumber(node, prefix+"-y")
	if err != nil {
		return entity.EngineInsets{}, err
	}
	if y != nil {
		insets.Top, insets.Bottom = y, y
	}
	for _, side := range []struct {
		name   string
		target **float64
	}{
		{name: "top", target: &insets.Top},
		{name: "right", target: &insets.Right},
		{name: "bottom", target: &insets.Bottom},
		{name: "left", target: &insets.Left},
	} {
		value, valueErr := frontendOptionalNumber(node, prefix+"-"+side.name)
		if valueErr != nil {
			return entity.EngineInsets{}, valueErr
		}
		if value != nil {
			*side.target = value
		}
	}
	return insets, nil
}

func applyFrontendSpacingClass(insets *entity.EngineInsets, prefix, className string) error {
	classPrefix := "p"
	if prefix == "margin" {
		classPrefix = "m"
	}
	parts := strings.Split(className, "-")
	if len(parts) != 2 || !strings.HasPrefix(parts[0], classPrefix) {
		return nil
	}
	level, err := strconv.ParseFloat(parts[1], 64)
	if err != nil || level < 0 {
		return fmt.Errorf("spacing level must be a non-negative number")
	}
	value := level * 8
	switch parts[0] {
	case classPrefix + "a":
		insets.Top, insets.Right, insets.Bottom, insets.Left = &value, &value, &value, &value
	case classPrefix + "x":
		insets.Left, insets.Right = &value, &value
	case classPrefix + "y":
		insets.Top, insets.Bottom = &value, &value
	case classPrefix + "t":
		insets.Top = &value
	case classPrefix + "r":
		insets.Right = &value
	case classPrefix + "b":
		insets.Bottom = &value
	case classPrefix + "l":
		insets.Left = &value
	}
	return nil
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

func frontendOptionalNumberFallback(node *frontendNode, name string, fallback *float64) (*float64, error) {
	value, err := frontendOptionalNumber(node, name)
	if err != nil || value != nil {
		return value, err
	}
	return fallback, nil
}

func frontendOptionalUint16(node *frontendNode, name string) (*uint16, error) {
	raw, ok := node.attrs[name]
	if !ok || strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	value, err := strconv.ParseUint(strings.TrimSpace(raw), 10, 16)
	if err != nil {
		return nil, fmt.Errorf("<%s> %s=%q is not an unsigned 16-bit integer", node.tag, name, raw)
	}
	result := uint16(value)
	return &result, nil
}

func frontendOptionalInt32(node *frontendNode, name string) (*int32, error) {
	raw, ok := node.attrs[name]
	if !ok || strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	value, err := strconv.ParseInt(strings.TrimSpace(raw), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("<%s> %s=%q is not a signed 32-bit integer", node.tag, name, raw)
	}
	result := int32(value)
	return &result, nil
}

func frontendOptionalBool(node *frontendNode, name string) (*bool, error) {
	raw, ok := node.attrs[name]
	if !ok || strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	value, err := strconv.ParseBool(strings.TrimSpace(raw))
	if err != nil {
		return nil, fmt.Errorf("<%s> %s=%q is not boolean", node.tag, name, raw)
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

type frontendPositionTracker struct {
	source    []byte
	offset    int
	line      int
	lineStart int
}

// at advances through source once across all calls. XML tokens arrive in
// source order, so source-position tracking stays linear for large documents.
func (tracker *frontendPositionTracker) at(offset int) entity.Position {
	if offset < 0 {
		offset = 0
	}
	if offset > len(tracker.source) {
		offset = len(tracker.source)
	}
	if offset < tracker.offset {
		tracker.offset, tracker.line, tracker.lineStart = 0, 1, 0
	}
	for index := tracker.offset; index < offset; index++ {
		if tracker.source[index] == '\n' {
			tracker.line++
			tracker.lineStart = index + 1
		}
	}
	tracker.offset = offset
	return entity.Position{Offset: offset, Line: tracker.line, Column: offset - tracker.lineStart + 1}
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
