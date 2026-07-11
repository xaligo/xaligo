package engine

import "math"

func routeOneV1EngineRoutePath(req routeRequestV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) routedPathV1EngineRouteTypes {
	if req.Grid > 0 {
		opt.Grid = req.Grid
	}
	inflated := make([]rectV1EngineRouteTypes, len(obstacles))
	for i, o := range obstacles {
		inflated[i] = inflateV1EngineRouteGeometry(o, opt.Clearance)
	}
	s := edgeMidpointV1EngineRouteGeometry(req.Src, req.SrcSide)
	if req.SrcAnchor != nil {
		s = *req.SrcAnchor
	}
	d := edgeMidpointV1EngineRouteGeometry(req.Dst, req.DstSide)
	if req.DstAnchor != nil {
		d = *req.DstAnchor
	}
	if req.SrcGap > 0 {
		s = extendV1EngineRouteGeometry(s, req.SrcSide, req.SrcGap)
	}
	if req.DstGap > 0 {
		d = extendV1EngineRouteGeometry(d, req.DstSide, req.DstGap)
	}
	s2 := extendV1EngineRouteGeometry(s, req.SrcSide, math.Max(opt.LaneGap, opt.Stub+req.SrcLane*opt.LaneGap))
	d2 := extendV1EngineRouteGeometry(d, req.DstSide, math.Max(opt.LaneGap, opt.Stub+req.DstLane*opt.LaneGap))

	if len(req.Bends) > 0 {
		points := routeViaBendsV1EngineRoutePath(s, d, s2, d2, req.Bends, opt)
		return routedPathV1EngineRouteTypes{ID: req.ID, Points: points}
	}

	candidates := buildCandidatesV1EngineRouteCandidate(s, d, s2, d2, inflated, placed, opt)

	var best []ptV1EngineRouteTypes
	bestCost := math.Inf(1)
	foundClean := false
	for _, raw := range candidates {
		if reversesEndpointStubV1EngineRoutePath(raw, s, s2, d2, d) {
			continue
		}
		points := simplifyRouteCandidateV1EngineRoutePath(raw)
		if len(points) < 2 {
			continue
		}
		if endpointApproachHitsTargetV1EngineRoutePath(points, req) {
			continue
		}
		hits := obstacleHitCountV1EngineRouteCandidate(points, inflated)
		if hits > 0 && foundClean {
			continue
		}
		cost := scorePathV1EngineRouteCandidate(points, inflated, placed, opt.LineMargin)
		if hits == 0 && !foundClean {
			foundClean = true
			bestCost = math.Inf(1)
			best = nil
		}
		if cost < bestCost {
			bestCost = cost
			best = points
		}
	}
	if best == nil {
		return routedPathV1EngineRouteTypes{ID: req.ID, Points: fallbackOrthogonalRouteV1EngineRoutePath(s, d, s2, d2)}
	}
	return routedPathV1EngineRouteTypes{ID: req.ID, Points: best}
}

func routeViaBendsV1EngineRoutePath(s, d, s2, d2 ptV1EngineRouteTypes, bends []ptV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []ptV1EngineRouteTypes {
	points := []ptV1EngineRouteTypes{s, s2}
	current := s2
	for _, bend := range bends {
		bend = snapPointV1EngineRoutePath(bend, opt.Grid)
		points = appendOrthogonalLegV1EngineRoutePath(points, current, bend)
		current = bend
	}
	points = appendOrthogonalLegV1EngineRoutePath(points, current, d2)
	points = append(points, d)
	return simplifyRouteCandidateV1EngineRoutePath(points)
}

func appendOrthogonalLegV1EngineRoutePath(points []ptV1EngineRouteTypes, from, to ptV1EngineRouteTypes) []ptV1EngineRouteTypes {
	if math.Abs(from.X-to.X) < epsV1EngineRouteTypes || math.Abs(from.Y-to.Y) < epsV1EngineRouteTypes {
		return append(points, to)
	}
	return append(points, ptV1EngineRouteTypes{X: to.X, Y: from.Y}, to)
}

func enforceOrthogonalPolylineV1EngineRoutePath(points []ptV1EngineRouteTypes) []ptV1EngineRouteTypes {
	if len(points) < 2 {
		return points
	}
	out := []ptV1EngineRouteTypes{points[0]}
	for i := 1; i < len(points); i++ {
		out = appendOrthogonalLegV1EngineRoutePath(out, out[len(out)-1], points[i])
	}
	return simplifyRouteCandidateV1EngineRoutePath(out)
}

func snapPointV1EngineRoutePath(p ptV1EngineRouteTypes, grid float64) ptV1EngineRouteTypes {
	return ptV1EngineRouteTypes{X: snapV1EngineRouteGeometry(p.X, grid), Y: snapV1EngineRouteGeometry(p.Y, grid)}
}

func fallbackOrthogonalRouteV1EngineRoutePath(s, d, s2, d2 ptV1EngineRouteTypes) []ptV1EngineRouteTypes {
	viaHorizontal := []ptV1EngineRouteTypes{s, s2, {X: d2.X, Y: s2.Y}, d2, d}
	viaVertical := []ptV1EngineRouteTypes{s, s2, {X: s2.X, Y: d2.Y}, d2, d}
	if pathLengthV1EngineRouteGeometry(viaVertical) < pathLengthV1EngineRouteGeometry(viaHorizontal) {
		return simplifyRouteCandidateV1EngineRoutePath(viaVertical)
	}
	return simplifyRouteCandidateV1EngineRoutePath(viaHorizontal)
}

func endpointApproachHitsTargetV1EngineRoutePath(points []ptV1EngineRouteTypes, req routeRequestV1EngineRouteTypes) bool {
	if len(points) >= 2 {
		if segIntersectsRectV1EngineRouteGeometry(segmentV1EngineRouteTypes{A: points[0], B: points[1]}, req.Src) {
			return true
		}
		last := len(points) - 1
		if segIntersectsRectV1EngineRouteGeometry(segmentV1EngineRouteTypes{A: points[last-1], B: points[last]}, req.Dst) {
			return true
		}
	}
	return false
}

func rerouteEndpointApproachV1EngineRoutePath(points []ptV1EngineRouteTypes, req routeRequestV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []ptV1EngineRouteTypes {
	points = rerouteDstApproachV1EngineRoutePath(points, req, opt)
	points = reversePointsV1EngineRoutePath(rerouteDstApproachV1EngineRoutePath(reversePointsV1EngineRoutePath(points), reverseRequestV1EngineRoutePath(req), opt))
	return simplifyRouteCandidateV1EngineRoutePath(points)
}

func rerouteDstApproachV1EngineRoutePath(points []ptV1EngineRouteTypes, req routeRequestV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []ptV1EngineRouteTypes {
	if len(points) < 4 || !pathIntersectsRectV1EngineRoutePath(points, req.Dst) {
		return points
	}
	end := points[len(points)-1]
	anchorIndex := len(points) - 4
	if anchorIndex < 0 {
		anchorIndex = 0
	}
	anchor := points[anchorIndex]
	out := append([]ptV1EngineRouteTypes(nil), points[:anchorIndex+1]...)
	margin := math.Max(opt.LineMargin, opt.LaneGap)
	switch req.DstSide {
	case sideTopV1EngineRouteTypes, sideBottomV1EngineRouteTypes:
		railX := req.Dst.X - margin
		if anchor.X >= req.Dst.X+req.Dst.W/2 {
			railX = req.Dst.X + req.Dst.W + margin
		}
		out = append(out, ptV1EngineRouteTypes{X: railX, Y: anchor.Y}, ptV1EngineRouteTypes{X: railX, Y: end.Y}, end)
	case sideLeftV1EngineRouteTypes, sideRightV1EngineRouteTypes:
		railY := req.Dst.Y - margin
		if anchor.Y >= req.Dst.Y+req.Dst.H/2 {
			railY = req.Dst.Y + req.Dst.H + margin
		}
		out = append(out, ptV1EngineRouteTypes{X: anchor.X, Y: railY}, ptV1EngineRouteTypes{X: end.X, Y: railY}, end)
	default:
		return points
	}
	return simplifyRouteCandidateV1EngineRoutePath(out)
}

func pathIntersectsRectV1EngineRoutePath(points []ptV1EngineRouteTypes, r rectV1EngineRouteTypes) bool {
	for _, seg := range toSegmentsV1EngineRouteGeometry(points) {
		if segIntersectsRectV1EngineRouteGeometry(seg, r) {
			return true
		}
	}
	return false
}

func reversePointsV1EngineRoutePath(points []ptV1EngineRouteTypes) []ptV1EngineRouteTypes {
	out := make([]ptV1EngineRouteTypes, len(points))
	for i := range points {
		out[i] = points[len(points)-1-i]
	}
	return out
}

func reverseRequestV1EngineRoutePath(req routeRequestV1EngineRouteTypes) routeRequestV1EngineRouteTypes {
	return routeRequestV1EngineRouteTypes{ID: req.ID, Kind: req.Kind, Src: req.Dst, Dst: req.Src, SrcSide: req.DstSide, DstSide: req.SrcSide, SrcGap: req.DstGap, DstGap: req.SrcGap, Grid: req.Grid}
}

func simplifyRouteCandidateV1EngineRoutePath(points []ptV1EngineRouteTypes) []ptV1EngineRouteTypes {
	out := []ptV1EngineRouteTypes{}
	for _, p := range points {
		if len(out) > 0 {
			last := out[len(out)-1]
			if math.Abs(last.X-p.X) < epsV1EngineRouteTypes && math.Abs(last.Y-p.Y) < epsV1EngineRouteTypes {
				continue
			}
		}
		out = append(out, p)
	}
	i := 1
	for i < len(out)-1 {
		if i == 1 || i == len(out)-2 {
			i++
			continue
		}
		a, b, c := out[i-1], out[i], out[i+1]
		collinearH := math.Abs(a.Y-b.Y) < epsV1EngineRouteTypes && math.Abs(b.Y-c.Y) < epsV1EngineRouteTypes
		collinearV := math.Abs(a.X-b.X) < epsV1EngineRouteTypes && math.Abs(b.X-c.X) < epsV1EngineRouteTypes
		if collinearH || collinearV {
			out = append(out[:i], out[i+1:]...)
		} else {
			i++
		}
	}
	return out
}

// reversesEndpointStub rejects a zero-width U-turn immediately after the
// source stub or immediately before the destination stub. Besides producing an
// unnecessary detour, simplification would erase the shared junction point.
func reversesEndpointStubV1EngineRoutePath(points []ptV1EngineRouteTypes, s, s2, d2, d ptV1EngineRouteTypes) bool {
	if len(points) < 3 {
		return false
	}
	if vectorsReverseV1EngineRoutePath(s2.X-s.X, s2.Y-s.Y, points[2].X-s2.X, points[2].Y-s2.Y) {
		return true
	}
	previous := points[len(points)-3]
	return vectorsReverseV1EngineRoutePath(d2.X-previous.X, d2.Y-previous.Y, d.X-d2.X, d.Y-d2.Y)
}

func vectorsReverseV1EngineRoutePath(ax, ay, bx, by float64) bool {
	if math.Abs(ax*by-ay*bx) > epsV1EngineRouteTypes {
		return false
	}
	return ax*bx+ay*by < -epsV1EngineRouteTypes
}
