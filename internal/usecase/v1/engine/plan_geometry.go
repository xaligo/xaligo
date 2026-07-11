package engine

import (
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

func rectOfV1EnginePlanGeometry(el *entity.Element) (rectV1EngineRouteTypes, bool) {
	if el == nil {
		return rectV1EngineRouteTypes{}, false
	}
	if el.Width <= 0 || el.Height <= 0 {
		return rectV1EngineRouteTypes{}, false
	}
	return rectV1EngineRouteTypes{X: el.X, Y: el.Y, W: el.Width, H: el.Height}, true
}

func inferSidesV1EnginePlanGeometry(src, dst rectV1EngineRouteTypes) (srcSide, dstSide sideV1EngineRouteTypes) {
	dx := dst.X + dst.W/2 - (src.X + src.W/2)
	dy := dst.Y + dst.H/2 - (src.Y + src.H/2)
	if math.Abs(dx) >= math.Abs(dy) {
		if dx >= 0 {
			return sideRightV1EngineRouteTypes, sideLeftV1EngineRouteTypes
		}
		return sideLeftV1EngineRouteTypes, sideRightV1EngineRouteTypes
	}
	if dy >= 0 {
		return sideBottomV1EngineRouteTypes, sideTopV1EngineRouteTypes
	}
	return sideTopV1EngineRouteTypes, sideBottomV1EngineRouteTypes
}

// ── Op builders (pixel → inch) ───────────────────────────────────────────────

func findPaperFrameV1EnginePlanGeometry(elements []*entity.Element) *rectV1EngineRouteTypes {
	for _, el := range elements {
		if el.ID == "paper-frame" || el.Type == "frame" {
			return &rectV1EngineRouteTypes{
				X: el.X,
				Y: el.Y,
				W: math.Max(1, el.Width),
				H: math.Max(1, el.Height),
			}
		}
	}
	return nil
}

func contentBoundsV1EnginePlanGeometry(elements []*entity.Element) rectV1EngineRouteTypes {
	visible := []*entity.Element{}
	for _, el := range elements {
		if el.Type != "arrow" && el.Type != "line" {
			visible = append(visible, el)
		}
	}
	if len(visible) == 0 {
		return rectV1EngineRouteTypes{X: 0, Y: 0, W: 1280, H: 720}
	}
	minX, minY := math.Inf(1), math.Inf(1)
	maxX, maxY := math.Inf(-1), math.Inf(-1)
	for _, el := range visible {
		minX = math.Min(minX, el.X)
		minY = math.Min(minY, el.Y)
		maxX = math.Max(maxX, el.X+el.Width)
		maxY = math.Max(maxY, el.Y+el.Height)
	}
	return rectV1EngineRouteTypes{X: minX, Y: minY, W: math.Max(1, maxX-minX), H: math.Max(1, maxY-minY)}
}

type posV1EnginePlanGeometry struct{ X, Y, W, H float64 }

func toPosV1EnginePlanGeometry(el *entity.Element, frame rectV1EngineRouteTypes, ppi float64) (posV1EnginePlanGeometry, bool) {
	w := el.Width / ppi
	h := el.Height / ppi
	if w <= 0 || h <= 0 {
		return posV1EnginePlanGeometry{}, false
	}
	return posV1EnginePlanGeometry{
		X: (el.X - frame.X) / ppi,
		Y: (el.Y - frame.Y) / ppi,
		W: w,
		H: h,
	}, true
}
