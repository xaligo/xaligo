package engine

import (
	"math"
	"regexp"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

var hexFullV1EnginePlanStyle = regexp.MustCompile(`^#?[0-9a-fA-F]{6}$`)

// ── Styling helpers ──────────────────────────────────────────────────────────

func linePropsV1EnginePlanStyle(el *entity.Element, ppi float64) entity.LineStyle {
	color := normalizeColorV1EnginePlanStyle(el.StrokeColor, "1E1E1E")
	dash := "solid"
	if el.StrokeStyle == "dashed" {
		dash = "dash"
	} else if el.StrokeStyle == "dotted" {
		dash = "dot"
	}
	transparency := opacityToTransparencyV1EnginePlanStyle(el.Opacity)
	if strings.EqualFold(strings.TrimSpace(el.StrokeColor), "transparent") || strings.EqualFold(strings.TrimSpace(el.StrokeColor), "#00000000") {
		transparency = 100
	}
	width := el.StrokeWidth
	if width == 0 {
		width = 1
	}
	width = pxToPtV1EnginePlanStyle(math.Max(0.25, width), ppi)
	return entity.LineStyle{Color: color, Width: width, Dash: dash, Transparency: transparency}
}

func fillPropsV1EnginePlanStyle(color string, transparency float64) entity.FillStyle {
	if color == "" || color == "transparent" {
		return entity.FillStyle{Color: "FFFFFF", Transparency: 100}
	}
	return entity.FillStyle{Color: normalizeColorV1EnginePlanStyle(color, "FFFFFF"), Transparency: transparency}
}

func normalizeColorV1EnginePlanStyle(color, fallback string) string {
	if color == "" || color == "transparent" {
		return fallback
	}
	trimmed := strings.TrimSpace(color)
	if hexFullV1EnginePlanStyle.MatchString(trimmed) {
		trimmed = strings.TrimPrefix(trimmed, "#")
		return strings.ToUpper(trimmed)
	}
	return fallback
}

func normalizeAlignV1EnginePlanStyle(align string) string {
	if align == "center" || align == "right" {
		return align
	}
	return "left"
}

func normalizeValignV1EnginePlanStyle(align string) string {
	if align == "middle" || align == "bottom" {
		return align
	}
	return "top"
}

func fontFaceV1EnginePlanStyle(fontFamily *int) string {
	if fontFamily == nil {
		return "Helvetica"
	}
	switch *fontFamily {
	case 1:
		return "Virgil"
	case 3:
		return "Cascadia Code"
	case 4:
		return "Assistant"
	case 5:
		return "Excalifont"
	case 6:
		return "Nunito"
	case 7:
		return "Lilita One"
	case 8:
		return "Comic Shanns"
	case 9:
		return "Liberation Sans"
	default:
		return "Helvetica"
	}
}

func opacityToTransparencyV1EnginePlanStyle(opacity *float64) float64 {
	value := 100.0
	if opacity != nil {
		value = *opacity
	}
	return math.Max(0, math.Min(100, 100-value))
}

func pxToPtV1EnginePlanStyle(px, ppi float64) float64 {
	if ppi <= 0 {
		ppi = defaultPxPerInchV1EnginePlanBuild
	}
	return px * 72.0 / ppi
}
