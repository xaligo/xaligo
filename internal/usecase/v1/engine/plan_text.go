package engine

import (
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func textOpV1EnginePlanText(el *entity.Element, frame rectV1EngineRouteTypes, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPosV1EnginePlanGeometry(el, frame, ppi)
	if !ok {
		return entity.DrawOp{}, false
	}
	text := el.Text
	if text == "" {
		text = el.RawText
	}
	if text == "" {
		return entity.DrawOp{}, false
	}
	if strings.EqualFold(strings.TrimSpace(el.StrokeColor), "transparent") || strings.EqualFold(strings.TrimSpace(el.StrokeColor), "#00000000") {
		return entity.DrawOp{}, false
	}
	fontSize := 12.0
	if el.FontSize != nil {
		fontSize = *el.FontSize
	}
	role := textRoleV1EnginePlanText(el)
	lineHeight := 1.2
	if el.LineHeight != nil && *el.LineHeight > 0 && !math.IsNaN(*el.LineHeight) && !math.IsInf(*el.LineHeight, 0) {
		lineHeight = *el.LineHeight
	}
	return entity.DrawOp{
		ID:         el.ID,
		FrontLayer: el.CustomData != nil && el.CustomData.FrameMetadata,
		Kind:       "text",
		X:          p.X,
		Y:          p.Y,
		W:          p.W,
		H:          p.H,
		Rotate:     el.Angle,
		Text:       text,
		Color:      normalizeColorV1EnginePlanStyle(el.StrokeColor, "1E1E1E"),
		FontFace:   fontFaceV1EnginePlanStyle(el.FontFamily),
		FontSize:   math.Max(1, pxToPtV1EnginePlanStyle(fontSize, ppi)),
		Bold:       el.FontStyle == "bold",
		Align:      normalizeAlignV1EnginePlanStyle(el.TextAlign),
		Valign:     normalizeValignV1EnginePlanStyle(el.VerticalAlign),
		TextLayout: planTextLayoutV1EnginePlanText(el, role, role != entity.TextRoleGroupHeader, lineHeight, ppi),
	}, true
}

func textRoleV1EnginePlanText(el *entity.Element) entity.TextRole {
	if el == nil {
		return entity.TextRoleLabel
	}
	if el.CustomData != nil {
		if el.CustomData.TextLayout != nil && el.CustomData.TextLayout.Role != "" {
			return el.CustomData.TextLayout.Role
		}
		switch {
		case el.CustomData.GroupHeaderContent:
			return entity.TextRoleGroupHeader
		case el.CustomData.AnchorContent:
			return entity.TextRoleItemLabel
		case el.CustomData.PortLabel:
			return entity.TextRolePortLabel
		case el.CustomData.CrossFrameLabel:
			return entity.TextRoleConnectorLabel
		}
	}
	// Older scene JSON did not expose semantic customData in the typed scene.
	// Keep its item-label behavior while new plans use the explicit role above.
	if strings.HasSuffix(el.ID, "-item-lbl") {
		return entity.TextRoleItemLabel
	}
	return entity.TextRoleLabel
}

func defaultTextLayoutV1EnginePlanText(role entity.TextRole, wrap bool, lineHeight float64) *entity.TextLayout {
	return &entity.TextLayout{
		Role:       role,
		Wrap:       wrap,
		Fit:        entity.TextFitShrink,
		Overflow:   entity.TextOverflowClip,
		Clip:       true,
		LineHeight: lineHeight,
		Padding:    entity.TextPadding{},
	}
}

func planTextLayoutV1EnginePlanText(el *entity.Element, role entity.TextRole, wrap bool, lineHeight, ppi float64) *entity.TextLayout {
	if el == nil || el.CustomData == nil || el.CustomData.TextLayout == nil {
		return defaultTextLayoutV1EnginePlanText(role, wrap, lineHeight)
	}
	layout := *el.CustomData.TextLayout
	if layout.Role == "" {
		layout.Role = role
	}
	if layout.Fit != entity.TextFitNone && layout.Fit != entity.TextFitShrink {
		layout.Fit = entity.TextFitShrink
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
	if layout.LineHeight <= 0 || math.IsNaN(layout.LineHeight) || math.IsInf(layout.LineHeight, 0) {
		layout.LineHeight = lineHeight
	}
	if ppi <= 0 {
		ppi = defaultPxPerInchV1EnginePlanBuild
	}
	layout.Padding.Top /= ppi
	layout.Padding.Right /= ppi
	layout.Padding.Bottom /= ppi
	layout.Padding.Left /= ppi
	return &layout
}
