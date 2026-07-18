package engine

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const geometryEpsilonV1EngineLayoutValidation = 1e-7

type layoutNumberRuleV1EngineLayoutValidation struct {
	name          string
	allowZero     bool
	allowNegative bool
	maximum       float64
}

// validateLayoutDocument normalizes the public numeric contract before any
// layout arithmetic takes place. strconv.ParseFloat accepts NaN and Inf, so a
// successful parse alone is not sufficient for geometry values.
func validateLayoutDocumentV1EngineLayoutValidation(root *entity.Node) error {
	if root == nil {
		return fmt.Errorf("document root is nil")
	}

	var walk func(*entity.Node) error
	walk = func(node *entity.Node) error {
		if err := validateLayoutNodeAttributesV1EngineLayoutValidation(node); err != nil {
			return err
		}
		if err := validateLayoutRatiosV1EngineLayoutValidation(node); err != nil {
			return err
		}
		for _, child := range node.Children {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(root)
}

func validateLayoutNodeAttributesV1EngineLayoutValidation(node *entity.Node) error {
	if node == nil {
		return nil
	}

	rules := []layoutNumberRuleV1EngineLayoutValidation{
		{name: "width"},
		{name: "height"},
		{name: "content-width"},
		{name: "content-height"},
		{name: "row"},
		{name: "col"},
		{name: "span", maximum: 12},
		{name: "gap", allowZero: true},
		{name: "margin", allowZero: true},
		{name: "margin-top", allowZero: true},
		{name: "margin-right", allowZero: true},
		{name: "margin-bottom", allowZero: true},
		{name: "margin-left", allowZero: true},
		{name: "font-size"},
		{name: "item-size"},
		{name: "scale"},
		{name: "coordinate-scale"},
		{name: "grid"},
		{name: "stroke-width"},
	}
	if node.Tag == "metadata" {
		rules = append(rules,
			layoutNumberRuleV1EngineLayoutValidation{name: "row-gap", allowZero: true},
			layoutNumberRuleV1EngineLayoutValidation{name: "key-width"},
		)
	}
	if node.Tag == "entry" {
		rules = append(rules, layoutNumberRuleV1EngineLayoutValidation{name: "key-width"})
	}
	if node.Tag == "port" {
		rules = append(rules,
			layoutNumberRuleV1EngineLayoutValidation{name: "w"},
			layoutNumberRuleV1EngineLayoutValidation{name: "h"},
			layoutNumberRuleV1EngineLayoutValidation{name: "x", allowZero: true, allowNegative: true},
			layoutNumberRuleV1EngineLayoutValidation{name: "y", allowZero: true, allowNegative: true},
		)
	}
	if node.Tag == "item" {
		rules = append(rules,
			layoutNumberRuleV1EngineLayoutValidation{name: "dx", allowZero: true, allowNegative: true},
			layoutNumberRuleV1EngineLayoutValidation{name: "dy", allowZero: true, allowNegative: true},
		)
	}
	if isBendNodeV1EngineLayoutValidation(node.Tag) {
		rules = append(rules,
			layoutNumberRuleV1EngineLayoutValidation{name: "x", allowZero: true, allowNegative: true},
			layoutNumberRuleV1EngineLayoutValidation{name: "y", allowZero: true, allowNegative: true},
		)
	}

	for _, rule := range rules {
		value, exists := node.Attrs[rule.name]
		if !exists {
			continue
		}
		if strings.TrimSpace(value) == "" {
			return newLayoutErrorV1EngineLayoutValidation(node, "%s must not be empty", rule.name)
		}
		parsed, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil || math.IsNaN(parsed) || math.IsInf(parsed, 0) {
			return newLayoutErrorV1EngineLayoutValidation(node, "%s=%q must be a finite number", rule.name, value)
		}
		if !rule.allowNegative {
			if parsed < 0 || (!rule.allowZero && parsed == 0) {
				qualifier := "greater than zero"
				if rule.allowZero {
					qualifier = "zero or greater"
				}
				return newLayoutErrorV1EngineLayoutValidation(node, "%s=%q must be %s", rule.name, value, qualifier)
			}
		}
		if rule.maximum > 0 && parsed > rule.maximum {
			return newLayoutErrorV1EngineLayoutValidation(node, "%s=%q must not exceed %.0f", rule.name, value, rule.maximum)
		}
	}

	if overflow := strings.ToLower(strings.TrimSpace(node.Attr("overflow"))); overflow != "" && overflow != "error" && overflow != "visible" {
		return newLayoutErrorV1EngineLayoutValidation(node, "overflow=%q must be error or visible", node.Attr("overflow"))
	}
	if err := validateLayoutAttributeV1EngineLayoutValidation(node); err != nil {
		return err
	}
	if err := validatePortSideV1EngineLayoutValidation(node); err != nil {
		return err
	}
	if err := validateSpacingClassV1EngineLayoutValidation(node); err != nil {
		return err
	}
	if isBendNodeV1EngineLayoutValidation(node.Tag) {
		if err := validateBendNodeV1EngineLayoutValidation(node); err != nil {
			return err
		}
	}
	for _, name := range []string{"bends", "points", "via"} {
		if value, exists := node.Attrs[name]; exists {
			if err := validateCoordinateListV1EngineLayoutValidation(node, name, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateLayoutAttributeV1EngineLayoutValidation(node *entity.Node) error {
	raw, exists := node.Attrs["layout"]
	if !exists || strings.TrimSpace(raw) == "" {
		return nil
	}
	value := strings.TrimSpace(raw)
	node.Attrs["layout"] = value

	if _, isGroup := awsGroupsV1EngineSceneTypes[node.Tag]; isGroup {
		switch value {
		case "vertical", "horizontal", "staggered":
			return nil
		default:
			return newLayoutErrorV1EngineLayoutValidation(node, "layout=%q must be vertical, horizontal, or staggered on <%s>", raw, node.Tag)
		}
	}

	switch node.Tag {
	case "frames", "frame", "container", "col", "database":
		if value == "vertical" || value == "horizontal" {
			return nil
		}
		if value == "staggered" {
			return newLayoutErrorV1EngineLayoutValidation(node, "layout=%q is only supported on AWS/group tags", raw)
		}
		return newLayoutErrorV1EngineLayoutValidation(node, "layout=%q must be vertical or horizontal on <%s>", raw, node.Tag)
	}

	if isUnknownContainerV1EngineLayoutValidation(node) {
		if value == "vertical" || value == "horizontal" || value == "staggered" {
			return nil
		}
		return newLayoutErrorV1EngineLayoutValidation(node, "layout=%q must be vertical, horizontal, or staggered on <%s>", raw, node.Tag)
	}
	return newLayoutErrorV1EngineLayoutValidation(node, "layout=%q is not supported on <%s>", raw, node.Tag)
}

func isUnknownContainerV1EngineLayoutValidation(node *entity.Node) bool {
	if node == nil || len(layoutKidsV1EngineLayoutNode(node)) == 0 {
		return false
	}
	switch node.Tag {
	case "frames", "frame", "container", "row", "col", "database", "entity", "table", "table-header", "table-row", "table-cell", "rectangle", "port", "item", "spacer", "blank",
		"connection", "connections", "src", "dst", "bend", "point", "via", "waypoint", "bends", "points", "path":
		return false
	default:
		_, isGroup := awsGroupsV1EngineSceneTypes[node.Tag]
		return !isGroup
	}
}

func validatePortSideV1EngineLayoutValidation(node *entity.Node) error {
	if node.Tag != "port" {
		return nil
	}
	raw, exists := node.Attrs["side"]
	if !exists || strings.TrimSpace(raw) == "" {
		return nil
	}
	side := strings.ToLower(strings.TrimSpace(raw))
	switch side {
	case "top", "right", "bottom", "left":
		node.Attrs["side"] = side
		return nil
	default:
		return newLayoutErrorV1EngineLayoutValidation(node, "side=%q must be top, right, bottom, or left", raw)
	}
}

func isBendNodeV1EngineLayoutValidation(tag string) bool {
	switch strings.ToLower(strings.TrimSpace(tag)) {
	case "bend", "point", "via", "waypoint":
		return true
	default:
		return false
	}
}

func validateBendNodeV1EngineLayoutValidation(node *entity.Node) error {
	_, hasX := node.Attrs["x"]
	_, hasY := node.Attrs["y"]
	if hasX != hasY {
		return newLayoutErrorV1EngineLayoutValidation(node, "bend coordinates require both x and y")
	}
	if hasX {
		return nil
	}
	return validateCoordinateListV1EngineLayoutValidation(node, "bend", node.Text)
}

func validateCoordinateListV1EngineLayoutValidation(node *entity.Node, name, value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return newLayoutErrorV1EngineLayoutValidation(node, "%s coordinates must not be empty", name)
	}
	replacer := strings.NewReplacer(";", " ", "|", " ", "\n", " ", "\t", " ")
	tokens := strings.Fields(replacer.Replace(value))
	if len(tokens) == 0 {
		return newLayoutErrorV1EngineLayoutValidation(node, "%s coordinates must contain x,y pairs", name)
	}
	for _, token := range tokens {
		parts := strings.Split(token, ",")
		if len(parts) != 2 {
			return newLayoutErrorV1EngineLayoutValidation(node, "%s coordinate %q must be an x,y pair", name, token)
		}
		for _, part := range parts {
			parsed, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
			if err != nil || math.IsNaN(parsed) || math.IsInf(parsed, 0) {
				return newLayoutErrorV1EngineLayoutValidation(node, "%s coordinate %q must contain finite numbers", name, token)
			}
		}
	}
	return nil
}

func validateSpacingClassV1EngineLayoutValidation(node *entity.Node) error {
	for _, token := range strings.Fields(node.Attr("class")) {
		if len(token) < 3 || token[2] != '-' || token[0] != 'p' && token[0] != 'm' || !strings.ContainsRune("axytrbl", rune(token[1])) {
			continue
		}
		value, err := strconv.Atoi(token[3:])
		// Unknown/malformed class tokens have historically been ignored by the
		// spacing parser. Preserve that compatibility, while rejecting an
		// explicitly numeric negative spacing value.
		if err == nil && value < 0 {
			return newLayoutErrorV1EngineLayoutValidation(node, "spacing class %q must use a non-negative integer", token)
		}
	}
	return nil
}

func validateLayoutRatiosV1EngineLayoutValidation(node *entity.Node) error {
	children := layoutKidsV1EngineLayoutNode(node)
	if len(children) == 0 || node.Tag == "rectangle" || node.Tag == "frames" || node.Tag == "metadata" || node.Tag == "entry" {
		return nil
	}

	if nodeUsesRowLayoutV1EngineLayoutValidation(node, children) {
		total := 0.0
		flexCount := 0
		for _, child := range children {
			if _, fixed := explicitSizeV1EngineLayoutConstraints(child, "width"); !fixed {
				flexCount++
			}
		}
		if flexCount == 0 {
			return nil
		}
		defaultSpan := 12 / float64(flexCount)
		for _, child := range children {
			if _, fixed := explicitSizeV1EngineLayoutConstraints(child, "width"); fixed {
				continue
			}
			total += attrFloatV1EngineLayoutAttributes(child.Attr("span"), defaultSpan)
			if math.IsInf(total, 0) || math.IsNaN(total) {
				return newLayoutErrorV1EngineLayoutValidation(node, "span ratio total must be finite")
			}
		}
		if total > 12+geometryEpsilonV1EngineLayoutValidation {
			return newLayoutErrorV1EngineLayoutValidation(node, "child span total %.6g exceeds the 12-column grid", total)
		}
		return nil
	}

	attribute := "row"
	if node.Attr("layout") == "horizontal" {
		attribute = "col"
	} else if node.Attr("layout") == "staggered" {
		return nil
	}
	total := 0.0
	fixedAttribute := "height"
	if attribute == "col" {
		fixedAttribute = "width"
	}
	for _, child := range children {
		if _, fixed := explicitSizeV1EngineLayoutConstraints(child, fixedAttribute); fixed {
			continue
		}
		total += attrFloatV1EngineLayoutAttributes(child.Attr(attribute), 1)
		if math.IsInf(total, 0) || math.IsNaN(total) {
			return newLayoutErrorV1EngineLayoutValidation(node, "%s ratio total must be finite", attribute)
		}
	}
	return nil
}

func nodeUsesRowLayoutV1EngineLayoutValidation(node *entity.Node, children []*entity.Node) bool {
	if node.Tag == "row" {
		return true
	}
	switch node.Tag {
	case "frame", "container", "col", "rectangle":
		return false
	}
	for _, child := range children {
		if !IsItemLikeV1EngineLayoutAttributes(child.Tag) {
			return false
		}
	}
	return true
}

func newLayoutErrorV1EngineLayoutValidation(node *entity.Node, format string, args ...any) error {
	err := fmt.Errorf(format, args...)
	if node == nil {
		return err
	}
	return &entity.ParseError{Position: node.Position, Err: err}
}

func validateResolvedGeometryV1EngineLayoutValidation(root *entity.Box) error {
	if root == nil {
		return fmt.Errorf("resolved layout root is nil")
	}
	return validateResolvedBoxV1EngineLayoutValidation(root, nil)
}

func validateResolvedBoxV1EngineLayoutValidation(box, parent *entity.Box) error {
	for name, value := range map[string]float64{
		"x": box.X, "y": box.Y, "width": box.W, "height": box.H,
		"content-x": box.ContentX, "content-y": box.ContentY,
		"content-width": box.ContentW, "content-height": box.ContentH,
		"intrinsic-width": box.IntrinsicW, "intrinsic-height": box.IntrinsicH,
	} {
		if math.IsNaN(value) || math.IsInf(value, 0) {
			return newResolvedLayoutErrorV1EngineLayoutValidation(box, "%s resolved to a non-finite value", name)
		}
	}
	if box.W <= 0 || box.H <= 0 {
		return newResolvedLayoutErrorV1EngineLayoutValidation(box, "resolved size must be positive, got %.6gx%.6g", box.W, box.H)
	}
	if box.ContentW < 0 || box.ContentH < 0 {
		return newResolvedLayoutErrorV1EngineLayoutValidation(box, "resolved content size must not be negative, got %.6gx%.6g", box.ContentW, box.ContentH)
	}
	if box.IntrinsicW < 0 || box.IntrinsicH < 0 {
		return newResolvedLayoutErrorV1EngineLayoutValidation(box, "intrinsic size must not be negative, got %.6gx%.6g", box.IntrinsicW, box.IntrinsicH)
	}
	if !containsRectV1EngineLayoutValidation(box.X, box.Y, box.W, box.H, box.ContentX, box.ContentY, box.ContentW, box.ContentH) {
		return newResolvedLayoutErrorV1EngineLayoutValidation(box, "content box overflows its border box")
	}
	if err := validateResolvedFrameMetadataV1EngineLayoutValidation(box); err != nil {
		return err
	}
	if parent != nil && !containsRectV1EngineLayoutValidation(parent.ContentX, parent.ContentY, parent.ContentW, parent.ContentH, box.X, box.Y, box.W, box.H) {
		if parent.Overflow != entity.OverflowVisible {
			return newResolvedLayoutErrorV1EngineLayoutValidation(box, "resolved box overflows parent <%s> content box; set overflow=\"visible\" on the parent to allow it explicitly", parent.Tag)
		}
		parent.Overflowed = true
	}
	for _, child := range box.Children {
		if err := validateResolvedBoxV1EngineLayoutValidation(child, box); err != nil {
			return err
		}
	}
	return nil
}

func validateResolvedFrameMetadataV1EngineLayoutValidation(box *entity.Box) error {
	if box == nil || box.FrameMetadata == nil {
		return nil
	}
	metadata := box.FrameMetadata
	if !isPositiveFiniteV1EngineLayoutConstraints(metadata.FontSize) {
		return newResolvedLayoutErrorV1EngineLayoutValidation(box, "frame metadata font size must be positive and finite")
	}
	for index, tag := range metadata.Tags {
		for name, value := range map[string]float64{
			"x": tag.X, "y": tag.Y, "width": tag.W, "height": tag.H, "key-width": tag.KeyW,
		} {
			if math.IsNaN(value) || math.IsInf(value, 0) {
				return newResolvedLayoutErrorV1EngineLayoutValidation(box, "frame metadata tag %d %s resolved to a non-finite value", index+1, name)
			}
		}
		if tag.W <= 0 || tag.H <= 0 || tag.KeyW <= 0 || tag.KeyW >= tag.W {
			return newResolvedLayoutErrorV1EngineLayoutValidation(box, "frame metadata tag %d has invalid resolved cell geometry", index+1)
		}
		if !containsRectV1EngineLayoutValidation(box.X, box.Y, box.W, box.H, tag.X, tag.Y, tag.W, tag.H) {
			return newResolvedLayoutErrorV1EngineLayoutValidation(box, "frame metadata tag %d overflows the frame border box", index+1)
		}
	}
	return nil
}

func containsRectV1EngineLayoutValidation(outerX, outerY, outerW, outerH, innerX, innerY, innerW, innerH float64) bool {
	return innerX >= outerX-geometryEpsilonV1EngineLayoutValidation &&
		innerY >= outerY-geometryEpsilonV1EngineLayoutValidation &&
		innerX+innerW <= outerX+outerW+geometryEpsilonV1EngineLayoutValidation &&
		innerY+innerH <= outerY+outerH+geometryEpsilonV1EngineLayoutValidation
}

func newResolvedLayoutErrorV1EngineLayoutValidation(box *entity.Box, format string, args ...any) error {
	err := fmt.Errorf("layout <%s> %q: %s", box.Tag, box.Label, fmt.Sprintf(format, args...))
	return &entity.ParseError{Position: box.Position, Err: err}
}
