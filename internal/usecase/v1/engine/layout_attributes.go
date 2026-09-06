package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

const spacingUnitV1EngineLayoutAttributes = 8

var (
	IULACA001V1EngineLayoutAttributes = share.NewMCode("IULACA-001", "Align content area width clamp branch")
	IULACA002V1EngineLayoutAttributes = share.NewMCode("IULACA-002", "Align content area height clamp branch")
	IULACA003V1EngineLayoutAttributes = share.NewMCode("IULACA-003", "Align content area right branch")
	IULACA004V1EngineLayoutAttributes = share.NewMCode("IULACA-004", "Align content area center branch")
	IULACA005V1EngineLayoutAttributes = share.NewMCode("IULACA-005", "Align content area bottom branch")
	IULACA006V1EngineLayoutAttributes = share.NewMCode("IULACA-006", "Align content area middle branch")
	IULLO001V1EngineLayoutAttributes  = share.NewMCode("IULLO-001", "Label of title branch")
	IULLO002V1EngineLayoutAttributes  = share.NewMCode("IULLO-002", "Label of text branch")
	IULLO003V1EngineLayoutAttributes  = share.NewMCode("IULLO-003", "Label of tag branch")
	IULAF001V1EngineLayoutAttributes  = share.NewMCode("IULAF-001", "Attribute float fallback empty branch")
	IULAF002V1EngineLayoutAttributes  = share.NewMCode("IULAF-002", "Attribute float parse failed branch")
	IULPCS001V1EngineLayoutAttributes = share.NewMCode("IULPCS-001", "Parse class spacing padding all branch")
	IULPCS002V1EngineLayoutAttributes = share.NewMCode("IULPCS-002", "Parse class spacing margin all branch")
	IULPCS003V1EngineLayoutAttributes = share.NewMCode("IULPCS-003", "Parse class spacing padding x branch")
	IULPCS004V1EngineLayoutAttributes = share.NewMCode("IULPCS-004", "Parse class spacing padding y branch")
	IULPCS005V1EngineLayoutAttributes = share.NewMCode("IULPCS-005", "Parse class spacing margin x branch")
	IULPCS006V1EngineLayoutAttributes = share.NewMCode("IULPCS-006", "Parse class spacing margin y branch")
	IULPCS007V1EngineLayoutAttributes = share.NewMCode("IULPCS-007", "Parse class spacing padding top branch")
	IULPCS008V1EngineLayoutAttributes = share.NewMCode("IULPCS-008", "Parse class spacing padding right branch")
	IULPCS009V1EngineLayoutAttributes = share.NewMCode("IULPCS-009", "Parse class spacing padding bottom branch")
	IULPCS010V1EngineLayoutAttributes = share.NewMCode("IULPCS-010", "Parse class spacing padding left branch")
	IULPCS011V1EngineLayoutAttributes = share.NewMCode("IULPCS-011", "Parse class spacing margin top branch")
	IULPCS012V1EngineLayoutAttributes = share.NewMCode("IULPCS-012", "Parse class spacing margin right branch")
	IULPCS013V1EngineLayoutAttributes = share.NewMCode("IULPCS-013", "Parse class spacing margin bottom branch")
	IULPCS014V1EngineLayoutAttributes = share.NewMCode("IULPCS-014", "Parse class spacing margin left branch")
	IULPA001V1EngineLayoutAttributes  = share.NewMCode("IULPA-001", "Parse align invalid branch")
	IULPA002V1EngineLayoutAttributes  = share.NewMCode("IULPA-002", "Parse align vertical branch")
	IULPA003V1EngineLayoutAttributes  = share.NewMCode("IULPA-003", "Parse align horizontal branch")
	IULSV001V1EngineLayoutAttributes  = share.NewMCode("IULSV-001", "Spacing value parse failed branch")
	IULPAM001V1EngineLayoutAttributes = share.NewMCode("IULPAM-001", "Parse attribute margin empty branch")
	IULPAM002V1EngineLayoutAttributes = share.NewMCode("IULPAM-002", "Parse attribute margin all branch")
	IULPAM003V1EngineLayoutAttributes = share.NewMCode("IULPAM-003", "Parse attribute margin top branch")
	IULPAM004V1EngineLayoutAttributes = share.NewMCode("IULPAM-004", "Parse attribute margin right branch")
	IULPAM005V1EngineLayoutAttributes = share.NewMCode("IULPAM-005", "Parse attribute margin bottom branch")
	IULPAM006V1EngineLayoutAttributes = share.NewMCode("IULPAM-006", "Parse attribute margin left branch")
)

func alignContentAreaV1EngineLayoutAttributes(node *entity.Node, x, y, w, h float64) (float64, float64, float64, float64, error) {
	if !isPositiveFiniteV1EngineLayoutConstraints(w) || !isPositiveFiniteV1EngineLayoutConstraints(h) {
		return 0, 0, 0, 0, newLayoutErrorV1EngineLayoutValidation(node, "padding and margins leave a non-positive content area %.6gx%.6g", w, h)
	}
	contentW := attrFloatV1EngineLayoutAttributes(node.Attr("content-width"), w)
	contentH := attrFloatV1EngineLayoutAttributes(node.Attr("content-height"), h)
	if contentW > w+geometryEpsilonV1EngineLayoutValidation {
		loggerV1EngineSharedLogging.DEBUG(IULACA001V1EngineLayoutAttributes, "branch clamp width", map[string]any{"tag": node.Tag, "contentWidth": contentW, "width": w})
		return 0, 0, 0, 0, newLayoutErrorV1EngineLayoutValidation(node, "content-width %.6g exceeds the available width %.6g", contentW, w)
	}
	if contentH > h+geometryEpsilonV1EngineLayoutValidation {
		loggerV1EngineSharedLogging.DEBUG(IULACA002V1EngineLayoutAttributes, "branch clamp height", map[string]any{"tag": node.Tag, "contentHeight": contentH, "height": h})
		return 0, 0, 0, 0, newLayoutErrorV1EngineLayoutValidation(node, "content-height %.6g exceeds the available height %.6g", contentH, h)
	}
	vert, horiz := parseAlignV1EngineLayoutAttributes(node.Attr("align"))
	switch horiz {
	case "right":
		loggerV1EngineSharedLogging.DEBUG(IULACA003V1EngineLayoutAttributes, "branch right", map[string]any{"tag": node.Tag})
		x += w - contentW
	case "center":
		loggerV1EngineSharedLogging.DEBUG(IULACA004V1EngineLayoutAttributes, "branch center", map[string]any{"tag": node.Tag})
		x += (w - contentW) / 2
	}
	switch vert {
	case "bottom":
		loggerV1EngineSharedLogging.DEBUG(IULACA005V1EngineLayoutAttributes, "branch bottom", map[string]any{"tag": node.Tag})
		y += h - contentH
	case "middle":
		loggerV1EngineSharedLogging.DEBUG(IULACA006V1EngineLayoutAttributes, "branch middle", map[string]any{"tag": node.Tag})
		y += (h - contentH) / 2
	}
	return x, y, contentW, contentH, nil
}

func childIDV1EngineLayoutAttributes(parent string, index int) string {
	return fmt.Sprintf("%s-%d", parent, index)
}

func labelOfV1EngineLayoutAttributes(n *entity.Node) string {
	if label, ok := awsGroupLabelV1EngineAwsResource(n); ok {
		return label
	}
	if title := n.Attr("title"); title != "" {
		loggerV1EngineSharedLogging.DEBUG(IULLO001V1EngineLayoutAttributes, "branch title", map[string]any{"tag": n.Tag})
		return title
	}
	if n.Text != "" {
		loggerV1EngineSharedLogging.DEBUG(IULLO002V1EngineLayoutAttributes, "branch text", map[string]any{"tag": n.Tag})
		return n.Text
	}
	if n.Tag == "rectangle" || n.Tag == "port" {
		return ""
	}
	loggerV1EngineSharedLogging.DEBUG(IULLO003V1EngineLayoutAttributes, "branch tag", map[string]any{"tag": n.Tag})
	return n.Tag
}

func attrFloatV1EngineLayoutAttributes(v string, fallback float64) float64 {
	if strings.TrimSpace(v) == "" {
		loggerV1EngineSharedLogging.DEBUG(IULAF001V1EngineLayoutAttributes, "branch fallback empty")
		return fallback
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		loggerV1EngineSharedLogging.WARN(IULAF002V1EngineLayoutAttributes, "branch parse failed", map[string]any{"value": v, "error": err})
		return fallback
	}
	return f
}

func attrFloatOKV1EngineLayoutAttributes(v string) (float64, bool) {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0, false
	}
	f, err := strconv.ParseFloat(v, 64)
	return f, err == nil
}

func firstAttrV1EngineLayoutAttributes(node *entity.Node, names ...string) string {
	if node == nil {
		return ""
	}
	for _, name := range names {
		if value := strings.TrimSpace(node.Attr(name)); value != "" {
			return value
		}
	}
	return ""
}

func parseClassSpacingV1EngineLayoutAttributes(class string) (entity.Spacing, entity.Spacing) {
	pad := entity.Spacing{}
	mar := entity.Spacing{}
	for _, tok := range strings.Fields(class) {
		switch {
		case strings.HasPrefix(tok, "pa-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS001V1EngineLayoutAttributes, "branch padding all", map[string]any{"token": tok})
			v := spacingValueV1EngineLayoutAttributes(tok[3:])
			pad = entity.Spacing{Top: v, Right: v, Bottom: v, Left: v}
		case strings.HasPrefix(tok, "ma-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS002V1EngineLayoutAttributes, "branch margin all", map[string]any{"token": tok})
			v := spacingValueV1EngineLayoutAttributes(tok[3:])
			mar = entity.Spacing{Top: v, Right: v, Bottom: v, Left: v}
		// 軸別一括: px=左右, py=上下
		case strings.HasPrefix(tok, "px-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS003V1EngineLayoutAttributes, "branch padding x", map[string]any{"token": tok})
			v := spacingValueV1EngineLayoutAttributes(tok[3:])
			pad.Left = v
			pad.Right = v
		case strings.HasPrefix(tok, "py-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS004V1EngineLayoutAttributes, "branch padding y", map[string]any{"token": tok})
			v := spacingValueV1EngineLayoutAttributes(tok[3:])
			pad.Top = v
			pad.Bottom = v
		case strings.HasPrefix(tok, "mx-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS005V1EngineLayoutAttributes, "branch margin x", map[string]any{"token": tok})
			v := spacingValueV1EngineLayoutAttributes(tok[3:])
			mar.Left = v
			mar.Right = v
		case strings.HasPrefix(tok, "my-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS006V1EngineLayoutAttributes, "branch margin y", map[string]any{"token": tok})
			v := spacingValueV1EngineLayoutAttributes(tok[3:])
			mar.Top = v
			mar.Bottom = v
		// 個別方向
		case strings.HasPrefix(tok, "pt-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS007V1EngineLayoutAttributes, "branch padding top", map[string]any{"token": tok})
			pad.Top = spacingValueV1EngineLayoutAttributes(tok[3:])
		case strings.HasPrefix(tok, "pr-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS008V1EngineLayoutAttributes, "branch padding right", map[string]any{"token": tok})
			pad.Right = spacingValueV1EngineLayoutAttributes(tok[3:])
		case strings.HasPrefix(tok, "pb-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS009V1EngineLayoutAttributes, "branch padding bottom", map[string]any{"token": tok})
			pad.Bottom = spacingValueV1EngineLayoutAttributes(tok[3:])
		case strings.HasPrefix(tok, "pl-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS010V1EngineLayoutAttributes, "branch padding left", map[string]any{"token": tok})
			pad.Left = spacingValueV1EngineLayoutAttributes(tok[3:])
		case strings.HasPrefix(tok, "mt-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS011V1EngineLayoutAttributes, "branch margin top", map[string]any{"token": tok})
			mar.Top = spacingValueV1EngineLayoutAttributes(tok[3:])
		case strings.HasPrefix(tok, "mr-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS012V1EngineLayoutAttributes, "branch margin right", map[string]any{"token": tok})
			mar.Right = spacingValueV1EngineLayoutAttributes(tok[3:])
		case strings.HasPrefix(tok, "mb-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS013V1EngineLayoutAttributes, "branch margin bottom", map[string]any{"token": tok})
			mar.Bottom = spacingValueV1EngineLayoutAttributes(tok[3:])
		case strings.HasPrefix(tok, "ml-"):
			loggerV1EngineSharedLogging.DEBUG(IULPCS014V1EngineLayoutAttributes, "branch margin left", map[string]any{"token": tok})
			mar.Left = spacingValueV1EngineLayoutAttributes(tok[3:])
		}
	}
	return pad, mar
}

// IsItemLikeV1EngineLayoutAttributes reports whether a tag behaves as a layout item slot.
func IsItemLikeV1EngineLayoutAttributes(tag string) bool {
	return tag == "item" || IsBlankV1EngineLayoutAttributes(tag)
}

// IsBlankV1EngineLayoutAttributes reports whether a tag participates in layout without rendering.
func IsBlankV1EngineLayoutAttributes(tag string) bool {
	return tag == "spacer" || tag == "blank"
}

func parseAlignV1EngineLayoutAttributes(align string) (vert, horiz string) {
	vert, horiz = "top", "left"
	normalized := strings.ToLower(strings.TrimSpace(align))
	if normalized == "" {
		return vert, horiz
	}
	parts := strings.SplitN(normalized, "-", 2)
	if len(parts) != 2 {
		loggerV1EngineSharedLogging.WARN(IULPA001V1EngineLayoutAttributes, "branch invalid align", map[string]any{"align": align})
		return vert, horiz
	}
	if parts[0] == "top" || parts[0] == "middle" || parts[0] == "bottom" {
		loggerV1EngineSharedLogging.DEBUG(IULPA002V1EngineLayoutAttributes, "branch vertical", map[string]any{"vertical": parts[0]})
		vert = parts[0]
	}
	if parts[1] == "left" || parts[1] == "center" || parts[1] == "right" {
		loggerV1EngineSharedLogging.DEBUG(IULPA003V1EngineLayoutAttributes, "branch horizontal", map[string]any{"horizontal": parts[1]})
		horiz = parts[1]
	}
	return vert, horiz
}

func spacingValueV1EngineLayoutAttributes(s string) float64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		loggerV1EngineSharedLogging.WARN(IULSV001V1EngineLayoutAttributes, "branch parse failed", map[string]any{"value": s, "error": err})
		return 0
	}
	return float64(n * spacingUnitV1EngineLayoutAttributes)
}

// parseAttrMarginV1EngineLayoutAttributes reads direct pixel-value margin attributes from a DSL node's
// attribute map. Supported attributes: margin, margin-top, margin-right,
// margin-bottom, margin-left. Values are in pixels (floats).
// When `margin` and a directional key (e.g. `margin-top`) are both present,
// the directional key overrides the corresponding side from `margin`.
func parseAttrMarginV1EngineLayoutAttributes(attrs map[string]string) entity.Spacing {
	if len(attrs) == 0 {
		loggerV1EngineSharedLogging.DEBUG(IULPAM001V1EngineLayoutAttributes, "branch empty attrs")
		return entity.Spacing{}
	}
	m := entity.Spacing{}
	if v := attrs["margin"]; v != "" {
		loggerV1EngineSharedLogging.DEBUG(IULPAM002V1EngineLayoutAttributes, "branch margin all", map[string]any{"value": v})
		val := attrFloatV1EngineLayoutAttributes(v, 0)
		m = entity.Spacing{Top: val, Right: val, Bottom: val, Left: val}
	}
	if v := attrs["margin-top"]; v != "" {
		loggerV1EngineSharedLogging.DEBUG(IULPAM003V1EngineLayoutAttributes, "branch margin top", map[string]any{"value": v})
		m.Top = attrFloatV1EngineLayoutAttributes(v, 0)
	}
	if v := attrs["margin-right"]; v != "" {
		loggerV1EngineSharedLogging.DEBUG(IULPAM004V1EngineLayoutAttributes, "branch margin right", map[string]any{"value": v})
		m.Right = attrFloatV1EngineLayoutAttributes(v, 0)
	}
	if v := attrs["margin-bottom"]; v != "" {
		loggerV1EngineSharedLogging.DEBUG(IULPAM005V1EngineLayoutAttributes, "branch margin bottom", map[string]any{"value": v})
		m.Bottom = attrFloatV1EngineLayoutAttributes(v, 0)
	}
	if v := attrs["margin-left"]; v != "" {
		loggerV1EngineSharedLogging.DEBUG(IULPAM006V1EngineLayoutAttributes, "branch margin left", map[string]any{"value": v})
		m.Left = attrFloatV1EngineLayoutAttributes(v, 0)
	}
	return m
}

// effectiveMarginV1EngineLayoutAttributes returns the combined margin for a node by summing
// class-based spacing (ma-N, mt-N …) and direct px-value attributes
// (margin, margin-top …).
func effectiveMarginV1EngineLayoutAttributes(node *entity.Node) entity.Spacing {
	_, classMar := parseClassSpacingV1EngineLayoutAttributes(node.Attr("class"))
	attrMar := parseAttrMarginV1EngineLayoutAttributes(node.Attrs)
	return entity.Spacing{
		Top:    classMar.Top + attrMar.Top,
		Right:  classMar.Right + attrMar.Right,
		Bottom: classMar.Bottom + attrMar.Bottom,
		Left:   classMar.Left + attrMar.Left,
	}
}
