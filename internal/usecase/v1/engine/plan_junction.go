package engine

import (
	"fmt"
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

type routeJunctionV1EnginePlanJunction struct {
	Point       ptV1EngineRouteTypes
	ConnectorID string
}

type junctionEndpointV1EnginePlanJunction struct {
	requestIndex int
	rect         rectV1EngineRouteTypes
	side         sideV1EngineRouteTypes
	gap          float64
	source       bool
	profile      string
}

// applyRouteJunctions makes route fan-out/fan-in connections share a centered
// anchor and stub. The returned points are drawn after the route lines, making
// the branch visually explicit. Traffic and ordinary connections stay
// independent.
func applyRouteJunctionsV1EnginePlanJunction(requests []routeRequestV1EngineRouteTypes, stub float64) []routeJunctionV1EnginePlanJunction {
	groups := map[string][]junctionEndpointV1EnginePlanJunction{}
	keys := []string{}
	add := func(key string, endpoint junctionEndpointV1EnginePlanJunction) {
		if _, exists := groups[key]; !exists {
			keys = append(keys, key)
		}
		groups[key] = append(groups[key], endpoint)
	}
	for i, req := range requests {
		if req.Kind != "route" {
			continue
		}
		add(junctionGroupKeyV1EnginePlanJunction("src", req.Src, req.SrcSide), junctionEndpointV1EnginePlanJunction{i, req.Src, req.SrcSide, req.SrcGap, true, req.SrcProfile})
		add(junctionGroupKeyV1EnginePlanJunction("dst", req.Dst, req.DstSide), junctionEndpointV1EnginePlanJunction{i, req.Dst, req.DstSide, req.DstGap, false, req.DstProfile})
	}

	junctions := []routeJunctionV1EnginePlanJunction{}
	seen := map[string]bool{}
	for _, key := range keys {
		group := groups[key]
		if len(group) < 2 {
			continue
		}
		first := group[0]
		anchor := anchorPointForProfileV1EnginePlanConnectorPrepare(first.rect, first.side, anchorGridV1EnginePlanBuild/2, first.profile)
		for _, endpoint := range group {
			copyPoint := anchor
			if endpoint.source {
				requests[endpoint.requestIndex].SrcAnchor = &copyPoint
				requests[endpoint.requestIndex].SrcLane = 0
			} else {
				requests[endpoint.requestIndex].DstAnchor = &copyPoint
				requests[endpoint.requestIndex].DstLane = 0
			}
		}
		point := anchor
		if first.gap > 0 {
			point = extendV1EngineRouteGeometry(point, first.side, first.gap)
		}
		point = extendV1EngineRouteGeometry(point, first.side, stub)
		pointKey := fmt.Sprintf("%.4f|%.4f", point.X, point.Y)
		if !seen[pointKey] {
			junctions = append(junctions, routeJunctionV1EnginePlanJunction{Point: point, ConnectorID: requests[first.requestIndex].ID})
			seen[pointKey] = true
		}
	}
	return junctions
}

func junctionGroupKeyV1EnginePlanJunction(prefix string, r rectV1EngineRouteTypes, side sideV1EngineRouteTypes) string {
	return fmt.Sprintf("%s|%.4f|%.4f|%.4f|%.4f|%s", prefix, r.X, r.Y, r.W, r.H, side)
}

func junctionOpV1EnginePlanJunction(id string, point ptV1EngineRouteTypes, frame rectV1EngineRouteTypes, ppi float64, line entity.LineStyle) entity.DrawOp {
	const diameterPx = 8.0
	diameter := diameterPx / ppi
	return entity.DrawOp{
		ID:         id + "-junction",
		FrontLayer: true,
		Kind:       "ellipse",
		X:          (point.X-frame.X)/ppi - diameter/2,
		Y:          (point.Y-frame.Y)/ppi - diameter/2,
		W:          diameter,
		H:          diameter,
		Fill:       &entity.FillStyle{Color: line.Color, Transparency: line.Transparency},
		Line:       &entity.LineStyle{Color: line.Color, Width: math.Max(pxToPtV1EnginePlanStyle(0.75, ppi), line.Width), Dash: "solid", Transparency: line.Transparency},
	}
}

func lineJumpMaskOpV1EnginePlanJunction(id string, crossing ptV1EngineRouteTypes, frame rectV1EngineRouteTypes, ppi float64, background string) entity.DrawOp {
	size := lineJumpSizePxV1EnginePlanBuild / ppi
	return entity.DrawOp{
		ID:         id,
		FrontLayer: true,
		Kind:       "rect",
		X:          (crossing.X-frame.X)/ppi - size/2,
		Y:          (crossing.Y-frame.Y)/ppi - size/2,
		W:          size,
		H:          size,
		Fill:       &entity.FillStyle{Color: background, Transparency: 0},
		Line:       &entity.LineStyle{Color: background, Width: pxToPtV1EnginePlanStyle(0.25, ppi), Transparency: 100},
	}
}

func groupBorderMaskOpV1EnginePlanJunction(id string, crossing ptV1EngineRouteTypes, frame rectV1EngineRouteTypes, ppi float64) entity.DrawOp {
	size := groupBorderMaskSizePxV1EnginePlanBuild / ppi
	return entity.DrawOp{
		ID:         id,
		FrontLayer: true,
		Kind:       "rect",
		X:          (crossing.X-frame.X)/ppi - size/2,
		Y:          (crossing.Y-frame.Y)/ppi - size/2,
		W:          size,
		H:          size,
		Fill:       &entity.FillStyle{Color: "FFFFFF", Transparency: 0},
		Line:       &entity.LineStyle{Color: "FFFFFF", Width: pxToPtV1EnginePlanStyle(0.25, ppi), Transparency: 100},
	}
}

// lineJumpBackground returns the uppermost opaque shape fill beneath a crossing.
// Transparent or partially transparent fills fall back to the slide background,
// as reproducing their composited color would require renderer-specific blending.
func lineJumpBackgroundV1EnginePlanJunction(crossing ptV1EngineRouteTypes, elements []*entity.Element, fallback string) string {
	color := fallback
	for _, el := range elements {
		if el.Type != "frame" && el.Type != "rectangle" && el.Type != "ellipse" {
			continue
		}
		if el.BackgroundColor == "" || el.BackgroundColor == "transparent" {
			continue
		}
		if el.Opacity != nil && *el.Opacity < 100 {
			continue
		}
		inside := crossing.X >= el.X && crossing.X <= el.X+el.Width &&
			crossing.Y >= el.Y && crossing.Y <= el.Y+el.Height
		if el.Type == "ellipse" && inside {
			rx, ry := el.Width/2, el.Height/2
			if rx <= 0 || ry <= 0 {
				inside = false
			} else {
				cx, cy := el.X+rx, el.Y+ry
				dx, dy := (crossing.X-cx)/rx, (crossing.Y-cy)/ry
				inside = dx*dx+dy*dy <= 1
			}
		}
		if inside {
			color = normalizeColorV1EnginePlanStyle(el.BackgroundColor, color)
		}
	}
	return color
}
