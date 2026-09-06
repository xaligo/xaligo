package v2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
)

const (
	frontendV1FrameMetadataContentTop = 31.0
	frontendV1FrameMetadataPageInset  = 4.0
	frontendV1FrameMetadataHeight     = 19.0
	frontendV1FrameMetadataGap        = 8.0
	frontendV1FrameMetadataFontSize   = 12.0
	frontendV1FrameMetadataLineHeight = 1.2
	frontendV1FrameMetadataTextPadX   = 4.0
	frontendV1FrameMetadataTextPadY   = 2.0
	frontendV1GroupTopInset           = 44.0
	frontendV1GroupSideInset          = 12.0
	frontendV1ItemGroupIconTopInset   = 26.0
	frontendV1ItemGroupTextTopInset   = 20.0
	frontendV1ItemGroupBottomInset    = 8.0
	frontendV1ItemGridGap             = 8.0
	frontendV1ItemVisualPad           = 6.0
	frontendV1ItemLabelGap            = 4.0
	frontendV1ItemLabelWidth          = 56.0
	frontendV1ItemLabelMinimumHeight  = 14.0
	frontendV1ItemLabelFontSize       = 8.0 * 96.0 / 72.0
	frontendV1ItemLabelLineHeight     = 1.25
	frontendV1PortWidth               = 56.0
	frontendV1PortHeight              = 22.0
)

type frontendV1GroupProfile struct {
	stroke      string
	strokeWidth float64
	strokeStyle entity.EngineLineStyle
	icon        string
}

type frontendV1MetadataTag struct {
	key   string
	value string
}

// frontendV1GroupProfiles is the authoring-profile boundary. AWS names and
// assets stop here; the Rust engine receives only generic group, visual, and
// icon parameters.
var frontendV1GroupProfiles = frontendAWSGroupProfiles()

func applyFrontendV1GeometryProfile(node *frontendNode, element *entity.EngineElementSpec) {
	if node == nil || element == nil {
		return
	}
	if element.Concept == entity.EngineConceptFrame {
		if frontendHasFrameMetadata(node) && !frontendHasExplicitInsets(node, "padding") {
			frontendSetInsetMinimum(&element.Padding.Top, frontendV1FrameMetadataContentTop)
		}
		return
	}
	if definition, ok := frontendAWSBoundaryAttachment(node.tag); ok {
		frontendSetNumberDefault(&element.Width, definition.DefaultSize)
		frontendSetNumberDefault(&element.Height, definition.DefaultSize)
		return
	}
	if element.Concept == entity.EngineConceptPort {
		frontendSetNumberDefault(&element.Width, frontendV1PortWidth)
		frontendSetNumberDefault(&element.Height, frontendV1PortHeight)
		return
	}
	if element.Concept != entity.EngineConceptGroup && element.Concept != entity.EngineConceptCapture {
		return
	}
	if frontendIsLayoutOnlyTag(node.tag) || frontendHasExplicitInsets(node, "padding") {
		return
	}
	if frontendHasOnlyItemChildren(node) {
		topInset := frontendV1ItemGroupTextTopInset
		if frontendHasGroupHeaderIcon(node) {
			topInset = frontendV1ItemGroupIconTopInset
		}
		frontendAddInset(&element.Padding.Top, topInset)
		frontendAddInset(&element.Padding.Right, frontendV1GroupSideInset)
		frontendAddInset(&element.Padding.Bottom, frontendV1ItemGroupBottomInset)
		frontendAddInset(&element.Padding.Left, frontendV1GroupSideInset)
		frontendSetNumberDefault(&element.Gap, frontendV1ItemGridGap)
		if strings.TrimSpace(node.attrs["align"]) == "" {
			element.Align = entity.EngineAlignCenter
			element.Justify = entity.EngineJustifyCenter
		}
		if strings.TrimSpace(node.attrs["overflow"]) == "" {
			element.Overflow = entity.EngineOverflowVisible
		}
		return
	}
	frontendAddInset(&element.Padding.Top, frontendV1GroupTopInset)
	frontendAddInset(&element.Padding.Right, frontendV1GroupSideInset)
	frontendAddInset(&element.Padding.Bottom, frontendV1GroupSideInset)
	frontendAddInset(&element.Padding.Left, frontendV1GroupSideInset)
}

func applyFrontendV1VisualProfile(node *frontendNode, element *entity.EngineElementSpec) {
	if node == nil || element == nil {
		return
	}
	if _, ok := frontendAWSBoundaryAttachment(node.tag); ok {
		element.Visual.Shape = entity.EngineShapeNone
		return
	}
	if element.Concept == entity.EngineConceptPort {
		frontendSetStringDefault(&element.Visual.Fill, "#ffffff")
		frontendSetStringDefault(&element.Visual.Stroke, "#334155")
		frontendSetNumberDefault(&element.Visual.StrokeWidth, 1)
		return
	}
	profile, known := frontendV1GroupProfiles[node.tag]
	if !known && (element.Concept != entity.EngineConceptGroup || frontendIsLayoutOnlyTag(node.tag)) {
		return
	}
	if !known {
		profile = frontendV1GroupProfile{stroke: "#1e1e1e", strokeWidth: 1}
	}
	frontendSetStringDefault(&element.Visual.Fill, "#ffffff")
	frontendSetStringDefault(&element.Visual.Stroke, profile.stroke)
	frontendSetNumberDefault(&element.Visual.StrokeWidth, profile.strokeWidth)
	frontendSetNumberDefault(&element.Visual.CornerRadius, 0)
	style := profile.strokeStyle
	if strings.TrimSpace(node.attrs["stroke-style"]) != "" || strings.TrimSpace(node.attrs["style"]) != "" {
		style = engineLineStyleForFrontendNode(node)
	}
	if style != "" && style != entity.EngineLineSolid {
		element.Line = &entity.EngineLineSpec{Style: style}
	}
}

func applyFrontendV1TextProfile(node *frontendNode, element *entity.EngineElementSpec) {
	if node == nil || element == nil || element.Text == nil {
		return
	}
	if element.Concept != entity.EngineConceptGroup && element.Concept != entity.EngineConceptCapture {
		return
	}
	if profile, ok := frontendV1GroupProfiles[node.tag]; ok {
		frontendSetStringDefault(&element.Text.Color, profile.stroke)
		element.Text.Role = string(entity.TextRoleGroupHeader)
	} else if !frontendIsLayoutOnlyTag(node.tag) {
		frontendSetStringDefault(&element.Text.Color, "#1e1e1e")
	}
	frontendSetStringDefault(&element.Text.FontFamily, "Helvetica")
	frontendSetNumberDefault(&element.Text.FontSize, 14)
	frontendSetNumberDefault(&element.Text.LineHeight, 1.25)
}

func applyFrontendV1IconProfile(node *frontendNode, element *entity.EngineElementSpec) {
	if node == nil || element == nil || element.Icon == nil {
		return
	}
	if definition, ok := frontendAWSBoundaryAttachment(node.tag); ok {
		width, height := definition.DefaultSize, definition.DefaultSize
		if element.Width != nil {
			width = *element.Width
		}
		if element.Height != nil {
			height = *element.Height
		}
		frontendSetNumberDefault(&element.Icon.Width, width)
		frontendSetNumberDefault(&element.Icon.Height, height)
		return
	}
	if element.Concept == entity.EngineConceptGroup || element.Concept == entity.EngineConceptCapture {
		frontendSetNumberDefault(&element.Icon.Width, 32)
		frontendSetNumberDefault(&element.Icon.Height, 32)
	}
}

func frontendHasFrameMetadata(node *frontendNode) bool {
	if strings.TrimSpace(node.attrs["title"]) != "" || strings.TrimSpace(node.attrs["version"]) != "" {
		return true
	}
	for _, child := range node.children {
		if child.tag == "metadata" {
			return true
		}
	}
	return false
}

func applyFrontendV1FrameMetadataComposition(node *frontendNode, frame *entity.EngineElementSpec) {
	if node == nil || frame == nil || frame.Concept != entity.EngineConceptFrame || !frontendHasFrameMetadata(node) {
		return
	}
	tags := []frontendV1MetadataTag{
		{key: "id", value: strings.TrimSpace(node.attrs["id"])},
		{key: "title", value: strings.TrimSpace(node.attrs["title"])},
		{key: "version", value: strings.TrimSpace(node.attrs["version"])},
	}
	for _, child := range node.children {
		if child.tag != "metadata" {
			continue
		}
		for _, entry := range child.children {
			if entry.tag == "entry" {
				tags = append(tags, frontendV1MetadataTag{
					key: strings.TrimSpace(entry.attrs["key"]), value: strings.TrimSpace(entry.attrs["value"]),
				})
			}
		}
		break
	}
	filtered := tags[:0]
	for _, tag := range tags {
		if tag.key != "" && tag.value != "" {
			filtered = append(filtered, tag)
		}
	}
	if len(filtered) == 0 {
		return
	}

	originalLayout := frame.Layout
	originalGap := frame.Gap
	originalAlign := frame.Align
	originalJustify := frame.Justify
	originalOverflow := frame.Overflow
	originalPadding := frame.Padding
	originalChildren := frame.Children

	metadataRow := entity.EngineElementSpec{
		ID: frame.ID + "-metadata", Concept: entity.EngineConceptGroup, SpanID: frame.SpanID,
		Layout: entity.EngineLayoutHorizontal, Overflow: entity.EngineOverflowError,
		Height: frontendFloatPointer(frontendV1FrameMetadataHeight), Gap: frontendFloatPointer(0),
		Margin: entity.EngineInsets{
			Left: frontendFloatPointer(frontendV1FrameMetadataPageInset), Right: frontendFloatPointer(frontendV1FrameMetadataPageInset),
		},
		Align: entity.EngineAlignStretch, Justify: entity.EngineJustifyStart,
		Visual: entity.EngineVisualSpec{Shape: entity.EngineShapeNone, Visible: frontendBoolPointer(false)},
	}
	metadataRow.Children = make([]entity.EngineElementSpec, 0, len(filtered)*2)
	for index, tag := range filtered {
		metadataRow.Children = append(metadataRow.Children,
			frontendV1MetadataCell(frame, index, "key", tag.key, true, 0),
			frontendV1MetadataCell(frame, index, "value", tag.value, false, frontendMetadataValueMargin(index, len(filtered))),
		)
	}

	contentTopMargin := frontendInsetValue(originalPadding.Top) - frontendV1FrameMetadataContentTop
	if contentTopMargin < 0 {
		contentTopMargin = 0
	}
	content := entity.EngineElementSpec{
		ID: frame.ID + "-content", Concept: entity.EngineConceptGroup, SpanID: frame.SpanID,
		Layout: originalLayout, Gap: originalGap, Weight: frontendFloatPointer(1),
		Margin: entity.EngineInsets{
			Top:  frontendOptionalPositivePointer(contentTopMargin),
			Left: originalPadding.Left, Right: originalPadding.Right,
		},
		Align: originalAlign, Justify: originalJustify, Overflow: originalOverflow,
		Visual: entity.EngineVisualSpec{Shape: entity.EngineShapeNone, Visible: frontendBoolPointer(false)}, Children: originalChildren,
	}

	frame.Text = nil
	frame.Layout = entity.EngineLayoutVertical
	frame.Gap = frontendFloatPointer(frontendV1FrameMetadataGap)
	frame.Align = entity.EngineAlignStretch
	frame.Justify = entity.EngineJustifyStart
	frame.Padding = entity.EngineInsets{
		Top: frontendFloatPointer(frontendV1FrameMetadataPageInset), Bottom: originalPadding.Bottom,
	}
	frame.Children = []entity.EngineElementSpec{metadataRow, content}
}

func frontendV1MetadataCell(frame *entity.EngineElementSpec, index int, kind, value string, key bool, marginRight float64) entity.EngineElementSpec {
	fill := "transparent"
	role := "frame-metadata-value"
	if key {
		fill = "#f8fafc"
		role = "frame-metadata-key"
	}
	return entity.EngineElementSpec{
		ID:      frame.ID + "-metadata-" + strconv.Itoa(index) + "-" + kind,
		Concept: entity.EngineConceptText, SpanID: frame.SpanID,
		Layout: entity.EngineLayoutNone, Overflow: entity.EngineOverflowError,
		Margin: entity.EngineInsets{Right: frontendOptionalPositivePointer(marginRight)},
		Align:  entity.EngineAlignCenter, Justify: entity.EngineJustifyCenter,
		Visual: entity.EngineVisualSpec{
			Shape: entity.EngineShapeRectangle, Fill: fill, Stroke: "#cbd5e1",
			StrokeWidth: frontendFloatPointer(0.75), CornerRadius: frontendFloatPointer(0),
		},
		Text: &entity.EngineTextSpec{
			Value: value, FontFamily: "Virgil", Color: "#64748b", Role: role,
			FontSize:   frontendFloatPointer(frontendV1FrameMetadataFontSize),
			LineHeight: frontendFloatPointer(frontendV1FrameMetadataLineHeight),
			Padding: entity.EngineInsets{
				Top: frontendFloatPointer(frontendV1FrameMetadataTextPadY), Right: frontendFloatPointer(frontendV1FrameMetadataTextPadX),
				Bottom: frontendFloatPointer(frontendV1FrameMetadataTextPadY), Left: frontendFloatPointer(frontendV1FrameMetadataTextPadX),
			},
		},
	}
}

func frontendMetadataValueMargin(index, count int) float64 {
	if index+1 < count {
		return frontendV1FrameMetadataGap
	}
	return 0
}

func frontendFloatPointer(value float64) *float64 {
	return &value
}

func frontendBoolPointer(value bool) *bool {
	return &value
}

func frontendOptionalPositivePointer(value float64) *float64 {
	if value <= 0 {
		return nil
	}
	return &value
}

func frontendInsetValue(value *float64) float64 {
	if value == nil {
		return 0
	}
	return *value
}

func frontendHasOnlyItemChildren(node *frontendNode) bool {
	if node == nil || len(node.children) == 0 {
		return false
	}
	found := false
	for _, child := range node.children {
		if frontendAWSComponent(child) {
			return false
		}
		if _, ok := frontendAWSBoundaryAttachment(child.tag); ok {
			continue
		}
		if awsprofile.IsResourceTag(child.tag) {
			found = true
			continue
		}
		switch child.tag {
		case "item", "spacer", "blank":
			found = true
		case "connections":
			continue
		default:
			return false
		}
	}
	return found
}

func frontendAWSBoundaryAttachment(tag string) (awsprofile.BoundaryAttachment, bool) {
	return awsprofile.BoundaryAttachmentForTag(tag)
}

func isFrontendPortTag(tag string) bool {
	if tag == "port" {
		return true
	}
	_, ok := frontendAWSBoundaryAttachment(tag)
	return ok
}

func validateFrontendAWSBoundaryAttachments(root *frontendNode) error {
	var walk func(*frontendNode, *frontendNode) error
	walk = func(node, parent *frontendNode) error {
		if node == nil {
			return nil
		}
		if root.attrs["version"] != "2" && frontendAWSComponent(node) {
			return fmt.Errorf("<%s> component requires XAL version 2", node.tag)
		}
		if err := validateFrontendAWSResource(node); err != nil {
			return err
		}
		if node.tag == "aws-listener" && (parent == nil || frontendAWSLoadBalancer(parent.tag) == "") {
			return fmt.Errorf("<aws-listener> must be a direct child of an ALB or NLB component")
		}
		if definition, ok := frontendAWSBoundaryAttachment(node.tag); ok {
			if parent == nil || parent.tag != definition.ParentTag {
				return fmt.Errorf("<%s> must be a direct child of <%s>", node.tag, definition.ParentTag)
			}
			id := strings.TrimSpace(node.attrs["id"])
			if id == "" {
				return fmt.Errorf("<%s> requires a non-empty id attribute", node.tag)
			}
			if strings.ContainsAny(id, " \t\r\n") {
				return fmt.Errorf("<%s id=%q> must not contain whitespace", node.tag, id)
			}
			if len(node.children) != 0 || strings.TrimSpace(node.text) != "" {
				return fmt.Errorf("<%s> must be empty", node.tag)
			}
			if raw := strings.TrimSpace(node.attrs["side"]); raw != "" {
				side := strings.ToLower(raw)
				switch side {
				case "top", "right", "bottom", "left":
					node.attrs["side"] = side
				default:
					return fmt.Errorf("<%s> side=%q must be top, right, bottom, or left", node.tag, raw)
				}
			}
		}
		for _, child := range node.children {
			if err := walk(child, node); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(root, nil)
}

func applyFrontendAWSBoundaryPortProfile(node *frontendNode, element *entity.EngineElementSpec) {
	definition, ok := frontendAWSBoundaryAttachment(node.tag)
	if !ok || element == nil || element.Port == nil {
		return
	}
	if element.Port.Size != nil {
		if strings.TrimSpace(node.attrs["width"]) == "" {
			element.Width = element.Port.Size
		}
		if strings.TrimSpace(node.attrs["height"]) == "" {
			element.Height = element.Port.Size
		}
	}
	width, height := definition.DefaultSize, definition.DefaultSize
	if element.Width != nil {
		width = *element.Width
	}
	if element.Height != nil {
		height = *element.Height
	}
	switch element.Port.Side {
	case entity.EngineSideTop:
		frontendAddNumber(&element.OffsetY, -height/2)
	case entity.EngineSideBottom:
		frontendAddNumber(&element.OffsetY, height/2)
	case entity.EngineSideLeft:
		frontendAddNumber(&element.OffsetX, -width/2)
	default:
		frontendAddNumber(&element.OffsetX, width/2)
	}
}

func frontendAddNumber(target **float64, amount float64) {
	value := amount
	if *target != nil {
		value += **target
	}
	*target = &value
}

func frontendHasGroupHeaderIcon(node *frontendNode) bool {
	if node == nil {
		return false
	}
	if firstNonEmpty(node.attrs["icon"], node.attrs["icon-ref"], node.attrs["icon-id"]) != "" {
		return true
	}
	profile, ok := frontendV1GroupProfiles[node.tag]
	return ok && profile.icon != ""
}

func frontendIsLayoutOnlyTag(tag string) bool {
	switch tag {
	case "row", "col", "container", "connections", "frames":
		return true
	default:
		return false
	}
}

func frontendHasExplicitInsets(node *frontendNode, prefix string) bool {
	for _, suffix := range []string{"", "-x", "-y", "-top", "-right", "-bottom", "-left"} {
		if strings.TrimSpace(node.attrs[prefix+suffix]) != "" {
			return true
		}
	}
	return false
}

func frontendSetInsetMinimum(target **float64, minimum float64) {
	if *target != nil && **target >= minimum {
		return
	}
	value := minimum
	*target = &value
}

func frontendAddInset(target **float64, amount float64) {
	value := amount
	if *target != nil {
		value += **target
	}
	*target = &value
}

func frontendSetNumberDefault(target **float64, value float64) {
	if *target == nil {
		*target = &value
	}
}

func frontendSetStringDefault(target *string, value string) {
	if strings.TrimSpace(*target) == "" {
		*target = value
	}
}

func frontendV1ItemTextHeight(text *entity.EngineTextSpec) float64 {
	if text == nil || strings.TrimSpace(text.Value) == "" {
		return 0
	}
	fontSize := frontendV1ItemLabelFontSize
	if text.FontSize != nil {
		fontSize = *text.FontSize
	}
	lineHeight := frontendV1ItemLabelLineHeight
	if text.LineHeight != nil {
		lineHeight = *text.LineHeight
	}
	lineCount := strings.Count(text.Value, "\n") + 1
	return math.Max(frontendV1ItemLabelMinimumHeight, math.Ceil(float64(lineCount)*fontSize*lineHeight))
}

// ApplyCatalogLabels completes catalog-backed item labels after the caller has
// loaded its asset catalog. The traversal is linear and does not reparse XAL
// or invoke the V1 renderer.
func ApplyCatalogLabels(document *entity.EngineDocumentSpec, labels map[int]string) {
	if document == nil || len(labels) == 0 {
		return
	}
	var apply func([]entity.EngineElementSpec)
	apply = func(elements []entity.EngineElementSpec) {
		for index := range elements {
			element := &elements[index]
			if element.Concept == entity.EngineConceptItem && element.Icon != nil && strings.HasPrefix(element.Icon.Ref, "catalog:") && (element.Text == nil || element.Text.Role == "catalog-label") {
				catalogID, err := strconv.Atoi(strings.TrimPrefix(element.Icon.Ref, "catalog:"))
				label := labels[catalogID]
				if err == nil && strings.TrimSpace(label) != "" {
					wrapped := frontendWrapCatalogLabel(label)
					if element.Text == nil {
						element.Text = &entity.EngineTextSpec{Role: "catalog-label"}
					}
					element.Text.Value = wrapped
					frontendSetNumberDefault(&element.Text.FontSize, 10)
					frontendReserveCatalogLabelHeight(element, wrapped)
				}
			}
			apply(element.Children)
		}
	}
	apply(document.Elements)
}

func frontendReserveCatalogLabelHeight(element *entity.EngineElementSpec, label string) {
	if element == nil || element.Icon == nil || element.Icon.Height == nil || strings.TrimSpace(label) == "" {
		return
	}
	frontendSetNumberDefault(&element.Text.FontSize, frontendV1ItemLabelFontSize)
	frontendSetNumberDefault(&element.Text.LineHeight, frontendV1ItemLabelLineHeight)
	height := *element.Icon.Height + frontendV1ItemLabelGap + frontendV1ItemTextHeight(element.Text)
	if !frontendHasExplicitParameter(element, "height") {
		element.Height = &height
	}
}

func frontendHasExplicitParameter(element *entity.EngineElementSpec, parameter string) bool {
	for _, source := range element.Sources {
		if source.Parameter == parameter && source.Origin == "explicit" {
			return true
		}
	}
	return false
}

func frontendWrapCatalogLabel(label string) string {
	const maxColumns = 9
	words := strings.Fields(label)
	lines := make([]string, 0, len(words))
	current := ""
	for _, word := range words {
		candidate := word
		if current != "" {
			candidate = current + " " + word
		}
		if utf8.RuneCountInString(candidate) <= maxColumns {
			current = candidate
			continue
		}
		if current != "" {
			lines = append(lines, current)
			current = ""
		}
		runes := []rune(word)
		for len(runes) > maxColumns {
			lines = append(lines, string(runes[:maxColumns]))
			runes = runes[maxColumns:]
		}
		current = string(runes)
	}
	if current != "" {
		lines = append(lines, current)
	}
	return strings.Join(lines, "\n")
}
