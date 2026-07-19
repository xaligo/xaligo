package engine

import "math"

func buildCandidatesV1EngineRouteCandidate(s, d, s2, d2 ptV1EngineRouteTypes, inflated []rectV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) [][]ptV1EngineRouteTypes {
	candidates := [][]ptV1EngineRouteTypes{}

	// 1) Straight only when endpoints share an axis exactly.
	if math.Abs(s.X-d.X) < epsV1EngineRouteTypes || math.Abs(s.Y-d.Y) < epsV1EngineRouteTypes {
		candidates = append(candidates, []ptV1EngineRouteTypes{s, d})
	}

	// 2) L-shaped (single bend) between the stub endpoints.
	candidates = append(candidates, []ptV1EngineRouteTypes{s, s2, {X: d2.X, Y: s2.Y}, d2, d})
	candidates = append(candidates, []ptV1EngineRouteTypes{s, s2, {X: s2.X, Y: d2.Y}, d2, d})

	// 3) Z-shaped (two bends) through candidate trunk lines.
	midX := snapV1EngineRouteGeometry((s2.X+d2.X)/2, opt.Grid)
	midY := snapV1EngineRouteGeometry((s2.Y+d2.Y)/2, opt.Grid)

	obstacleXEdges := []float64{}
	obstacleYEdges := []float64{}
	for _, r := range inflated {
		obstacleXEdges = append(obstacleXEdges, r.X-opt.LaneGap, r.X, r.X+r.W, r.X+r.W+opt.LaneGap)
		obstacleYEdges = append(obstacleYEdges, r.Y-opt.LaneGap, r.Y, r.Y+r.H, r.Y+r.H+opt.LaneGap)
	}

	xRaw := []float64{midX, midX - opt.LaneGap, midX + opt.LaneGap, midX - 2*opt.LaneGap, midX + 2*opt.LaneGap}
	xRaw = append(xRaw, gutterCentersV1EngineRouteGutter(projectXV1EngineRouteGutter(inflated), math.Min(s2.X, d2.X), math.Max(s2.X, d2.X))...)
	xRaw = append(xRaw, obstacleXEdges...)
	xRaw = append(xRaw, placedLaneCoordsV1EngineRouteCandidate(placed, false, opt)...)
	for i := range xRaw {
		xRaw[i] = snapV1EngineRouteGeometry(xRaw[i], opt.Grid)
	}
	xLines := dedupeV1EngineRouteGutter(xRaw)

	yRaw := []float64{midY, midY - opt.LaneGap, midY + opt.LaneGap, midY - 2*opt.LaneGap, midY + 2*opt.LaneGap}
	yRaw = append(yRaw, gutterCentersV1EngineRouteGutter(projectYV1EngineRouteGutter(inflated), math.Min(s2.Y, d2.Y), math.Max(s2.Y, d2.Y))...)
	yRaw = append(yRaw, obstacleYEdges...)
	yRaw = append(yRaw, placedLaneCoordsV1EngineRouteCandidate(placed, true, opt)...)
	for i := range yRaw {
		yRaw[i] = snapV1EngineRouteGeometry(yRaw[i], opt.Grid)
	}
	yLines := dedupeV1EngineRouteGutter(yRaw)

	for _, x := range xLines {
		candidates = append(candidates, []ptV1EngineRouteTypes{s, s2, {X: x, Y: s2.Y}, {X: x, Y: d2.Y}, d2, d})
	}
	for _, y := range yLines {
		candidates = append(candidates, []ptV1EngineRouteTypes{s, s2, {X: s2.X, Y: y}, {X: d2.X, Y: y}, d2, d})
	}
	return candidates
}

func placedLaneCoordsV1EngineRouteCandidate(placed [][]segmentV1EngineRouteTypes, horizontal bool, opt routerOptionsV1EngineRouteTypes) []float64 {
	offsets := make([]float64, 0, 3)
	for lane := 0; lane < 3; lane++ {
		offsets = append(offsets, opt.LineMargin+float64(lane)*opt.LaneGap)
	}
	out := []float64{}
	for _, path := range placed {
		for _, seg := range path {
			if isHorizontalV1EngineRouteGeometry(seg) != horizontal {
				continue
			}
			base := seg.A.X
			if horizontal {
				base = seg.A.Y
			}
			for _, off := range offsets {
				if off <= 0 {
					continue
				}
				out = append(out, base-off, base+off)
			}
		}
	}
	return out
}

// ── Scoring ──────────────────────────────────────────────────────────────────

func scorePathV1EngineRouteCandidate(points []ptV1EngineRouteTypes, inflated []rectV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, margin float64) float64 {
	segs := toSegmentsV1EngineRouteGeometry(points)

	obstacleHits := obstacleHitCountV1EngineRouteCandidate(points, inflated)

	crossings := 0.0
	overlap := 0.0
	proximity := 0.0
	for _, other := range placed {
		for _, seg := range segs {
			for _, oseg := range other {
				if segmentsCrossV1EngineRouteGeometry(seg, oseg) {
					crossings++
				}
				overlap += collinearOverlapV1EngineRouteGeometry(seg, oseg)
				proximity += nearParallelV1EngineRouteGeometry(seg, oseg, margin)
			}
		}
	}

	length := pathLengthV1EngineRouteGeometry(points)
	bends := math.Max(0, float64(len(points))-2)

	overlapTerm := 0.0
	if overlap > 0 {
		overlapTerm = 1 + overlap/math.Max(1, margin)
	}
	proximityTerm := 0.0
	if proximity > 0 {
		proximityTerm = 1 + proximity/100
	}

	return wObstacleV1EngineRouteTypes*obstacleHits +
		wCrossV1EngineRouteTypes*crossings +
		wOverlapV1EngineRouteTypes*overlapTerm +
		wProximityV1EngineRouteTypes*proximityTerm +
		wLenV1EngineRouteTypes*(length/1000) +
		wBendV1EngineRouteTypes*bends
}

func obstacleHitCountV1EngineRouteCandidate(points []ptV1EngineRouteTypes, inflated []rectV1EngineRouteTypes) float64 {
	hits := 0.0
	for _, seg := range toSegmentsV1EngineRouteGeometry(points) {
		for _, r := range inflated {
			if segIntersectsRectV1EngineRouteGeometry(seg, r) {
				hits++
			}
		}
	}
	return hits
}
