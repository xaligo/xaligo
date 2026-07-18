package engine

import (
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

func connectorIDLabelOpV1EnginePlanConnectorLabel(id string, path routedPathV1EngineRouteTypes, allPaths []routedPathV1EngineRouteTypes, obstacles, hardObstacles []rectV1EngineRouteTypes, placedLabels []rectV1EngineRouteTypes, frame rectV1EngineRouteTypes, ppi float64, line entity.LineStyle) (entity.DrawOp, rectV1EngineRouteTypes, bool) {
	if len(path.Points) < 2 {
		return entity.DrawOp{}, rectV1EngineRouteTypes{}, false
	}
	p, ok := connectorLabelPointV1EnginePlanConnectorLabel(path, allPaths, obstacles, hardObstacles, placedLabels)
	if !ok {
		return entity.DrawOp{}, rectV1EngineRouteTypes{}, false
	}
	w := connectorLabelWidthPxV1EnginePlanBuild / ppi
	h := connectorLabelHeightPxV1EnginePlanBuild / ppi
	labelRect := connectorIDLabelRectV1EnginePlanConnectorLabel(p)
	return entity.DrawOp{
		ID:         id + "-label",
		FrontLayer: true,
		Kind:       "text",
		X:          (p.X-frame.X)/ppi - w/2,
		Y:          (p.Y-frame.Y)/ppi - h/2,
		W:          w,
		H:          h,
		Text:       id,
		Color:      line.Color,
		FontFace:   "Helvetica",
		FontSize:   math.Max(1, pxToPtV1EnginePlanStyle(connectorLabelFontPxV1EnginePlanBuild, ppi)),
		Bold:       true,
		Align:      "center",
		Valign:     "middle",
		TextLayout: defaultTextLayoutV1EnginePlanText(entity.TextRoleConnectorLabel, false, 1.0),
	}, labelRect, true
}

func connectorIDLabelRectV1EnginePlanConnectorLabel(center ptV1EngineRouteTypes) rectV1EngineRouteTypes {
	return rectV1EngineRouteTypes{
		X: center.X - connectorLabelWidthPxV1EnginePlanBuild/2,
		Y: center.Y - connectorLabelHeightPxV1EnginePlanBuild/2,
		W: connectorLabelWidthPxV1EnginePlanBuild,
		H: connectorLabelHeightPxV1EnginePlanBuild,
	}
}

type connectorLabelCandidateV1EnginePlanConnectorLabel struct {
	Point    ptV1EngineRouteTypes
	Priority float64
}

func connectorLabelPointV1EnginePlanConnectorLabel(path routedPathV1EngineRouteTypes, allPaths []routedPathV1EngineRouteTypes, obstacles, hardObstacles []rectV1EngineRouteTypes, placedLabels []rectV1EngineRouteTypes) (ptV1EngineRouteTypes, bool) {
	candidates := connectorLabelCandidatesV1EnginePlanConnectorLabel(path.Points)
	if len(candidates) == 0 {
		return ptV1EngineRouteTypes{}, false
	}
	best := ptV1EngineRouteTypes{}
	bestScore := math.Inf(1)
	found := false
	for _, candidate := range candidates {
		label := connectorIDLabelRectV1EnginePlanConnectorLabel(candidate.Point)
		blocked := false
		for _, obstacle := range hardObstacles {
			if rectsOverlapV1EnginePlanConnectorLabel(label, inflateV1EngineRouteGeometry(obstacle, 1)) {
				blocked = true
				break
			}
		}
		if blocked {
			continue
		}
		score := candidate.Priority + connectorLabelScoreV1EnginePlanConnectorLabel(path.ID, candidate.Point, allPaths, obstacles, placedLabels)
		if score < bestScore {
			best = candidate.Point
			bestScore = score
			found = true
		}
	}
	return best, found
}

func connectorLabelCandidatesV1EnginePlanConnectorLabel(points []ptV1EngineRouteTypes) []connectorLabelCandidateV1EnginePlanConnectorLabel {
	if len(points) < 2 {
		return nil
	}
	candidates := []connectorLabelCandidateV1EnginePlanConnectorLabel{}
	add := func(point ptV1EngineRouteTypes, priority float64) {
		candidates = append(candidates, connectorLabelCandidateV1EnginePlanConnectorLabel{Point: point, Priority: priority})
	}
	addEndpointCandidates := func(anchor, neighbor ptV1EngineRouteTypes) {
		dx := neighbor.X - anchor.X
		dy := neighbor.Y - anchor.Y
		length := math.Abs(dx) + math.Abs(dy)
		if length <= epsV1EngineRouteTypes {
			return
		}
		ux, uy := dx/length, dy/length
		px, py := -uy, ux
		for _, alongOffset := range []float64{12, 18, 26, 34, 42} {
			for _, sideOffset := range []float64{-8, 8, -14, 14, -22, 22, -30, 30, -40, 40} {
				point := ptV1EngineRouteTypes{X: anchor.X + ux*alongOffset + px*sideOffset, Y: anchor.Y + uy*alongOffset + py*sideOffset}
				add(point, alongOffset*0.12+math.Abs(sideOffset)*0.22)
			}
		}
	}
	addEndpointCandidates(points[0], points[1])
	addEndpointCandidates(points[len(points)-1], points[len(points)-2])
	for i := 1; i < len(points)-1; i++ {
		p := points[i]
		prev := points[i-1]
		next := points[i+1]
		if (math.Abs(prev.X-p.X) < epsV1EngineRouteTypes && math.Abs(next.X-p.X) < epsV1EngineRouteTypes) || (math.Abs(prev.Y-p.Y) < epsV1EngineRouteTypes && math.Abs(next.Y-p.Y) < epsV1EngineRouteTypes) {
			continue
		}
		for _, dx := range []float64{-8, 8, -14, 14, -22, 22, -30, 30, -40, 40} {
			for _, dy := range []float64{-8, 8, -14, 14, -22, 22, -30, 30, -40, 40} {
				point := ptV1EngineRouteTypes{X: p.X + dx, Y: p.Y + dy}
				add(point, 2.0+math.Hypot(dx, dy)*0.18)
			}
		}
	}
	return candidates
}

func connectorLabelScoreV1EnginePlanConnectorLabel(pathID string, center ptV1EngineRouteTypes, allPaths []routedPathV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, placedLabels []rectV1EngineRouteTypes) float64 {
	label := connectorIDLabelRectV1EnginePlanConnectorLabel(center)
	score := 0.0
	for _, obstacle := range obstacles {
		if rectsOverlapV1EnginePlanConnectorLabel(label, inflateV1EngineRouteGeometry(obstacle, 2)) {
			score += 1000
		}
		score += proximityPenaltyV1EnginePlanConnectorLabel(distanceRectToRectV1EnginePlanConnectorLabel(label, obstacle), 18, 18)
	}
	for _, path := range allPaths {
		pathPenalty := 160.0
		if path.ID == pathID {
			pathPenalty = 6.0
		}
		for _, seg := range toSegmentsV1EngineRouteGeometry(path.Points) {
			if segIntersectsRectV1EngineRouteGeometry(seg, label) {
				score += pathPenalty
			}
			if path.ID != pathID {
				score += proximityPenaltyV1EnginePlanConnectorLabel(distancePointToSegmentV1EnginePlanConnectorLabel(center, seg), 14, pathPenalty/10)
			}
		}
	}
	for _, placed := range placedLabels {
		if rectsOverlapV1EnginePlanConnectorLabel(label, inflateV1EngineRouteGeometry(placed, 2)) {
			score += 1200
		}
		score += proximityPenaltyV1EnginePlanConnectorLabel(distanceRectToRectV1EnginePlanConnectorLabel(label, placed), 24, 40)
	}
	return score
}

func proximityPenaltyV1EnginePlanConnectorLabel(distance, threshold, weight float64) float64 {
	if distance >= threshold {
		return 0
	}
	return weight * (threshold - distance) / threshold
}

func rectsOverlapV1EnginePlanConnectorLabel(a, b rectV1EngineRouteTypes) bool {
	return a.X < b.X+b.W && a.X+a.W > b.X && a.Y < b.Y+b.H && a.Y+a.H > b.Y
}

func distanceRectToRectV1EnginePlanConnectorLabel(a, b rectV1EngineRouteTypes) float64 {
	dx := math.Max(math.Max(b.X-(a.X+a.W), a.X-(b.X+b.W)), 0)
	dy := math.Max(math.Max(b.Y-(a.Y+a.H), a.Y-(b.Y+b.H)), 0)
	return math.Hypot(dx, dy)
}

func distancePointToSegmentV1EnginePlanConnectorLabel(p ptV1EngineRouteTypes, seg segmentV1EngineRouteTypes) float64 {
	x1, y1 := seg.A.X, seg.A.Y
	x2, y2 := seg.B.X, seg.B.Y
	dx, dy := x2-x1, y2-y1
	lengthSq := dx*dx + dy*dy
	if lengthSq <= epsV1EngineRouteTypes {
		return math.Hypot(p.X-x1, p.Y-y1)
	}
	t := ((p.X-x1)*dx + (p.Y-y1)*dy) / lengthSq
	t = math.Max(0, math.Min(1, t))
	closest := ptV1EngineRouteTypes{X: x1 + t*dx, Y: y1 + t*dy}
	return math.Hypot(p.X-closest.X, p.Y-closest.Y)
}
