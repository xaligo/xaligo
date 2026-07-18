package engine

import "math"

// separateExactOverlaps moves only an internal trunk segment onto an adjacent
// lane when the chosen route would otherwise share the exact same coordinates
// with an earlier connector. Endpoints and endpoint stubs remain fixed.
func separateExactOverlapsV1EngineRouteOverlap(points []ptV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []ptV1EngineRouteTypes {
	if len(points) < 2 || len(placed) == 0 || opt.LaneGap <= 0 {
		return points
	}
	inflated := make([]rectV1EngineRouteTypes, len(obstacles))
	for i, obstacle := range obstacles {
		inflated[i] = inflateV1EngineRouteGeometry(obstacle, opt.Clearance)
	}
	best := append([]ptV1EngineRouteTypes(nil), points...)
	bestOverlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(best), placed)
	if bestOverlap <= epsV1EngineRouteTypes {
		return best
	}
	bestScore := scorePathV1EngineRouteCandidate(best, inflated, placed, opt.LineMargin)
	// An overlapping endpoint stub is escaped with a short local jog. Because the
	// jog is brief and hugs the endpoint, it only needs a light visual gap from
	// the actual icon rectangles rather than the full routing clearance halo,
	// which is too strict in crowded corners (e.g. WAF at the top row sharing the
	// IAM approach lane).
	stubMargin := math.Min(opt.LineMargin, opt.Clearance) / 2
	stubObstacles := make([]rectV1EngineRouteTypes, len(obstacles))
	for i, obstacle := range obstacles {
		stubObstacles[i] = inflateV1EngineRouteGeometry(obstacle, stubMargin)
	}
	if candidate, ok := offsetFirstStubV1EngineRouteOverlap(best, placed, stubObstacles, opt.LaneGap); ok {
		overlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(candidate), placed)
		if overlap < bestOverlap-epsV1EngineRouteTypes {
			best, bestOverlap = candidate, overlap
			bestScore = scorePathV1EngineRouteCandidate(candidate, inflated, placed, opt.LineMargin)
		}
	}
	for segmentIndex := 1; segmentIndex < len(best)-2; segmentIndex++ {
		segment := segmentV1EngineRouteTypes{A: best[segmentIndex], B: best[segmentIndex+1]}
		for _, direction := range []float64{-1, 1} {
			candidate := append([]ptV1EngineRouteTypes(nil), best...)
			offset := direction * opt.LaneGap
			if isHorizontalV1EngineRouteGeometry(segment) {
				candidate[segmentIndex].Y += offset
				candidate[segmentIndex+1].Y += offset
			} else {
				candidate[segmentIndex].X += offset
				candidate[segmentIndex+1].X += offset
			}
			candidate = simplifyV1EngineRouteGeometry(candidate)
			if obstacleHitCountV1EngineRouteCandidate(candidate, inflated) > 0 {
				continue
			}
			overlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(candidate), placed)
			score := scorePathV1EngineRouteCandidate(candidate, inflated, placed, opt.LineMargin)
			if overlap < bestOverlap-epsV1EngineRouteTypes || (math.Abs(overlap-bestOverlap) < epsV1EngineRouteTypes && score < bestScore) {
				best, bestOverlap, bestScore = candidate, overlap, score
			}
		}
	}
	return best
}

// separateObstacleHits reuses the lane-offset post-processing for routes that
// still clip a visual obstacle after candidate search. This is intentionally
// limited to internal segments so endpoint bindings and their short stubs stay
// stable; it is mainly used for crowded group-header tag bands.
func separateObstacleHitsV1EngineRouteOverlap(points []ptV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []ptV1EngineRouteTypes {
	if len(points) < 4 || len(obstacles) == 0 || opt.LaneGap <= 0 {
		return points
	}
	best := append([]ptV1EngineRouteTypes(nil), points...)
	bestHits := obstacleHitCountV1EngineRouteCandidate(best, obstacles)
	if bestHits <= epsV1EngineRouteTypes {
		return best
	}
	bestOverlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(best), placed)
	bestScore := scorePathV1EngineRouteCandidate(best, obstacles, placed, opt.LineMargin)
	for segmentIndex := 1; segmentIndex < len(best)-2; segmentIndex++ {
		base := append([]ptV1EngineRouteTypes(nil), best...)
		seg := segmentV1EngineRouteTypes{A: base[segmentIndex], B: base[segmentIndex+1]}
		for _, direction := range []float64{-1, 1} {
			for _, mult := range []float64{1, 2, 3} {
				candidate := append([]ptV1EngineRouteTypes(nil), base...)
				offset := direction * mult * opt.LaneGap
				if isHorizontalV1EngineRouteGeometry(seg) {
					candidate[segmentIndex].Y += offset
					candidate[segmentIndex+1].Y += offset
				} else {
					candidate[segmentIndex].X += offset
					candidate[segmentIndex+1].X += offset
				}
				candidate = simplifyV1EngineRouteGeometry(candidate)
				hits := obstacleHitCountV1EngineRouteCandidate(candidate, obstacles)
				overlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(candidate), placed)
				score := scorePathV1EngineRouteCandidate(candidate, obstacles, placed, opt.LineMargin)
				if hits < bestHits-epsV1EngineRouteTypes || (math.Abs(hits-bestHits) < epsV1EngineRouteTypes && (overlap < bestOverlap-epsV1EngineRouteTypes || (math.Abs(overlap-bestOverlap) < epsV1EngineRouteTypes && score < bestScore))) {
					best, bestHits, bestOverlap, bestScore = candidate, hits, overlap, score
				}
			}
		}
	}
	return best
}

// offsetFirstStub escapes an overlapping source stub with a short local jog onto
// an adjacent lane. The endpoint (points[0]) is preserved; only the run leaving
// it is shifted. Both perpendicular directions and a few lane multiples are
// tried, and the obstacle-free candidate that removes the most stub overlap is
// returned. obstacles are pre-inflated by the caller's stub margin.
func offsetFirstStubV1EngineRouteOverlap(points []ptV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, laneGap float64) ([]ptV1EngineRouteTypes, bool) {
	if len(points) < 2 || laneGap <= 0 {
		return nil, false
	}
	first := segmentV1EngineRouteTypes{A: points[0], B: points[1]}
	baseOverlap := exactOverlapLengthV1EngineRouteOverlap([]segmentV1EngineRouteTypes{first}, placed)
	if baseOverlap <= epsV1EngineRouteTypes {
		return nil, false
	}
	const exitPx = 2.0
	var bestCandidate []ptV1EngineRouteTypes
	bestOverlap := baseOverlap
	for _, direction := range []float64{-1, 1} {
		for _, mult := range []float64{1, 2, 3} {
			offset := direction * mult * laneGap
			near := first.A
			offsetNear := first.A
			offsetEnd := first.B
			if isHorizontalV1EngineRouteGeometry(first) {
				near.X += math.Copysign(math.Min(exitPx, math.Abs(first.B.X-first.A.X)/2), first.B.X-first.A.X)
				offsetNear = near
				offsetNear.Y += offset
				offsetEnd.Y += offset
			} else {
				near.Y += math.Copysign(math.Min(exitPx, math.Abs(first.B.Y-first.A.Y)/2), first.B.Y-first.A.Y)
				offsetNear = near
				offsetNear.X += offset
				offsetEnd.X += offset
			}
			candidate := []ptV1EngineRouteTypes{first.A, near, offsetNear, offsetEnd, first.B}
			candidate = append(candidate, points[2:]...)
			candidate = simplifyV1EngineRouteGeometry(candidate)
			if obstacleHitCountV1EngineRouteCandidate(candidate, obstacles) > 0 {
				continue
			}
			overlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(candidate), placed)
			if overlap < bestOverlap-epsV1EngineRouteTypes {
				bestCandidate, bestOverlap = candidate, overlap
			}
		}
	}
	if bestCandidate == nil {
		return nil, false
	}
	return bestCandidate, true
}

func exactOverlapLengthV1EngineRouteOverlap(path []segmentV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes) float64 {
	total := 0.0
	for _, current := range path {
		for _, otherPath := range placed {
			for _, other := range otherPath {
				total += collinearOverlapV1EngineRouteGeometry(current, other)
			}
		}
	}
	return total
}

func inflateRectsV1EngineRouteOverlap(rects []rectV1EngineRouteTypes, margin float64) []rectV1EngineRouteTypes {
	if margin <= 0 {
		return rects
	}
	out := make([]rectV1EngineRouteTypes, len(rects))
	for i, r := range rects {
		out[i] = inflateV1EngineRouteGeometry(r, margin)
	}
	return out
}

func filterObstaclesV1EngineRouteOverlap(obstacles []rectV1EngineRouteTypes, req routeRequestV1EngineRouteTypes) []rectV1EngineRouteTypes {
	out := make([]rectV1EngineRouteTypes, 0, len(obstacles))
	for _, o := range obstacles {
		if sameRectV1EngineRouteGeometry(o, req.Src) || sameRectV1EngineRouteGeometry(o, req.Dst) {
			continue
		}
		out = append(out, o)
	}
	return out
}
