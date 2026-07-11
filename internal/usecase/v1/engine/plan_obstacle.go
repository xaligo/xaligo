package engine

import (
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func collectObstaclesV1EnginePlanObstacle(elements []*entity.Element) []rectV1EngineRouteTypes {
	rects := []rectV1EngineRouteTypes{}
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		isHeader := el.CustomData != nil && el.CustomData.GroupHeader
		if el.Type != "image" && el.Type != "text" && !isHeader {
			continue
		}
		r, ok := rectOfV1EnginePlanGeometry(el)
		if !ok {
			continue
		}
		rects = append(rects, r)
	}
	return rects
}

// collectContainerBorderPaths reserves a clear lane beside visible container
// strokes. Borders are routing guides rather than solid obstacles: connectors
// can cross them to move between nested groups, but parallel overlap and paths
// inside LineMargin are penalised by the normal lane-scoring logic.
func collectContainerBorderPathsV1EnginePlanObstacle(elements []*entity.Element) [][]segmentV1EngineRouteTypes {
	paths := make([][]segmentV1EngineRouteTypes, 0)
	for _, el := range elements {
		if el.ID == "paper-frame" || (el.Type != "frame" && el.Type != "rectangle") {
			continue
		}
		stroke := strings.ToLower(strings.TrimSpace(el.StrokeColor))
		if stroke == "" || stroke == "transparent" || stroke == "#00000000" {
			continue
		}
		r, ok := rectOfV1EnginePlanGeometry(el)
		if !ok {
			continue
		}
		topLeft := ptV1EngineRouteTypes{X: r.X, Y: r.Y}
		topRight := ptV1EngineRouteTypes{X: r.X + r.W, Y: r.Y}
		bottomLeft := ptV1EngineRouteTypes{X: r.X, Y: r.Y + r.H}
		bottomRight := ptV1EngineRouteTypes{X: r.X + r.W, Y: r.Y + r.H}
		paths = append(paths,
			[]segmentV1EngineRouteTypes{{A: topLeft, B: topRight}},
			[]segmentV1EngineRouteTypes{{A: bottomLeft, B: bottomRight}},
			[]segmentV1EngineRouteTypes{{A: topLeft, B: bottomLeft}},
			[]segmentV1EngineRouteTypes{{A: topRight, B: bottomRight}},
		)
	}
	return paths
}

func collectGroupBorderPathsV1EnginePlanObstacle(elements []*entity.Element) []segmentV1EngineRouteTypes {
	var paths []segmentV1EngineRouteTypes
	for _, el := range elements {
		if el.CustomData == nil || !el.CustomData.GroupBorder {
			continue
		}
		stroke := strings.ToLower(strings.TrimSpace(el.StrokeColor))
		if stroke == "" || stroke == "transparent" || stroke == "#00000000" {
			continue
		}
		r, ok := rectOfV1EnginePlanGeometry(el)
		if !ok {
			continue
		}
		tl := ptV1EngineRouteTypes{X: r.X, Y: r.Y}
		tr := ptV1EngineRouteTypes{X: r.X + r.W, Y: r.Y}
		bl := ptV1EngineRouteTypes{X: r.X, Y: r.Y + r.H}
		br := ptV1EngineRouteTypes{X: r.X + r.W, Y: r.Y + r.H}
		paths = append(paths, segmentV1EngineRouteTypes{A: tl, B: tr}, segmentV1EngineRouteTypes{A: tr, B: br}, segmentV1EngineRouteTypes{A: br, B: bl}, segmentV1EngineRouteTypes{A: bl, B: tl})
	}
	return paths
}

func pathBorderCrossingsV1EnginePlanObstacle(path routedPathV1EngineRouteTypes, borders []segmentV1EngineRouteTypes) []ptV1EngineRouteTypes {
	var out []ptV1EngineRouteTypes
	for _, pathSeg := range toSegmentsV1EngineRouteGeometry(path.Points) {
		for _, border := range borders {
			p, ok := crossingPointV1EngineRouteGeometry(pathSeg, border)
			if !ok {
				continue
			}
			duplicate := false
			for _, existing := range out {
				if math.Abs(existing.X-p.X) < epsV1EngineRouteTypes && math.Abs(existing.Y-p.Y) < epsV1EngineRouteTypes {
					duplicate = true
					break
				}
			}
			if !duplicate {
				out = append(out, p)
			}
		}
	}
	return out
}

func sideFromFixedPointV1EnginePlanObstacle(fp []float64) (sideV1EngineRouteTypes, bool) {
	if len(fp) < 2 {
		return "", false
	}
	fx, fy := fp[0], fp[1]
	if math.Abs(fy-0) < 0.01 {
		return sideTopV1EngineRouteTypes, true
	}
	if math.Abs(fy-1) < 0.01 {
		return sideBottomV1EngineRouteTypes, true
	}
	if math.Abs(fx-0) < 0.01 {
		return sideLeftV1EngineRouteTypes, true
	}
	if math.Abs(fx-1) < 0.01 {
		return sideRightV1EngineRouteTypes, true
	}
	return "", false
}

func anchorFromFixedPointV1EnginePlanObstacle(r rectV1EngineRouteTypes, s sideV1EngineRouteTypes, fp []float64, explicit bool) (ptV1EngineRouteTypes, bool) {
	if len(fp) < 2 {
		return ptV1EngineRouteTypes{}, false
	}
	fx := math.Max(0, math.Min(1, fp[0]))
	fy := math.Max(0, math.Min(1, fp[1]))
	switch s {
	case sideTopV1EngineRouteTypes, sideBottomV1EngineRouteTypes:
		if !explicit && math.Abs(fx-0.5) < 0.01 {
			return ptV1EngineRouteTypes{}, false
		}
	case sideLeftV1EngineRouteTypes, sideRightV1EngineRouteTypes:
		if !explicit && math.Abs(fy-0.5) < 0.01 {
			return ptV1EngineRouteTypes{}, false
		}
	default:
		return ptV1EngineRouteTypes{}, false
	}
	return ptV1EngineRouteTypes{X: r.X + r.W*fx, Y: r.Y + r.H*fy}, true
}

func pointFromFixedPointV1EnginePlanObstacle(r rectV1EngineRouteTypes, fp []float64) ptV1EngineRouteTypes {
	if len(fp) < 2 {
		return ptV1EngineRouteTypes{X: r.X + r.W/2, Y: r.Y + r.H/2}
	}
	fx := math.Max(0, math.Min(1, fp[0]))
	fy := math.Max(0, math.Min(1, fp[1]))
	return ptV1EngineRouteTypes{X: r.X + r.W*fx, Y: r.Y + r.H*fy}
}
