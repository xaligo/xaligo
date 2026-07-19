package engine

import (
	"math"
	"strconv"
)

// routeConnections routes every request against a shared obstacle set. Requests
// are processed in order; each finalised path informs later ones.
func routeConnectionsV1EngineRouteBuild(requests []routeRequestV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []routedPathV1EngineRouteTypes {
	placed := make([][]segmentV1EngineRouteTypes, len(opt.Reserved))
	for i := range opt.Reserved {
		placed[i] = append([]segmentV1EngineRouteTypes(nil), opt.Reserved[i]...)
	}
	results := make([]routedPathV1EngineRouteTypes, 0, len(requests))
	routePaths := map[string][]ptV1EngineRouteTypes{}
	for _, req := range requests {
		local := filterObstaclesV1EngineRouteOverlap(obstacles, req)
		requestOpt := opt
		if req.HardAvoid {
			requestOpt.HardObstacles = append(append([]rectV1EngineRouteTypes(nil), opt.HardObstacles...), local...)
		}
		path := routeOneV1EngineRoutePath(req, local, placed, requestOpt)
		followedRoute := false
		if req.Kind == "traffic" {
			if base, ok := matchingRoutePathV1EngineRouteBuild(req, routePaths); ok {
				path.Points = trafficAlongsideRouteV1EngineRouteBuild(base, path.Points, opt.LaneGap)
				followedRoute = true
			} else {
				path.Points = separateExactOverlapsV1EngineRouteOverlap(path.Points, placed[len(opt.Reserved):], local, opt)
			}
		} else if req.Kind != "route" {
			path.Points = separateExactOverlapsV1EngineRouteOverlap(path.Points, placed[len(opt.Reserved):], local, opt)
		}
		visualMargin := math.Min(opt.LineMargin, opt.Clearance) / 2
		path.Points = separateObstacleHitsV1EngineRouteOverlap(path.Points, placed[len(opt.Reserved):], inflateRectsV1EngineRouteOverlap(local, visualMargin), opt)
		if len(req.Bends) == 0 {
			path.Points = rerouteEndpointApproachV1EngineRoutePath(path.Points, req, opt)
		}
		if followedRoute {
			path.Points = restoreDestinationApproachV1EngineRouteBuild(path.Points, req.DstSide, opt.Stub)
		}
		path.Points = enforceOrthogonalPolylineV1EngineRoutePath(path.Points)
		path.Points = enforceHardObstacleExclusionV1EngineRouteBuild(req, path.Points, local, placed, opt)
		path.Points = enforceHardObstacleExclusionV1EngineRouteBuild(req, path.Points, local, placed, requestOpt)
		results = append(results, path)
		if req.Kind == "route" && len(path.Points) >= 2 {
			routePaths[routePairKeyV1EngineRouteBuild(req, false)] = append([]ptV1EngineRouteTypes(nil), path.Points...)
		}
		if len(path.Points) >= 2 {
			placed = append(placed, toSegmentsV1EngineRouteGeometry(path.Points))
		}
	}
	return results
}

func enforceHardObstacleExclusionV1EngineRouteBuild(req routeRequestV1EngineRouteTypes, points []ptV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []ptV1EngineRouteTypes {
	hardObstacles := inflateRectsV1EngineRouteOverlap(opt.HardObstacles, 1)
	if len(hardObstacles) == 0 || obstacleHitCountV1EngineRouteCandidate(points, hardObstacles) == 0 {
		return points
	}
	retry := req
	retry.Bends = nil
	retryObstacles := append([]rectV1EngineRouteTypes(nil), obstacles...)
	retryObstacles = append(retryObstacles, opt.HardObstacles...)
	path := routeOneV1EngineRoutePath(retry, retryObstacles, placed, opt)
	visualMargin := math.Min(opt.LineMargin, opt.Clearance) / 2
	path.Points = separateObstacleHitsV1EngineRouteOverlap(path.Points, placed, inflateRectsV1EngineRouteOverlap(retryObstacles, visualMargin), opt)
	path.Points = rerouteEndpointApproachV1EngineRoutePath(path.Points, retry, opt)
	path.Points = enforceOrthogonalPolylineV1EngineRoutePath(path.Points)
	if obstacleHitCountV1EngineRouteCandidate(path.Points, hardObstacles) == 0 {
		return path.Points
	}
	if fallback, ok := horizontalStripSafeFallbackV1EngineRouteBuild(retry, hardObstacles, opt); ok && obstacleHitCountV1EngineRouteCandidate(fallback, hardObstacles) == 0 {
		return fallback
	}
	if fallback, ok := visibilityGridFallbackV1EngineRouteHardObstacle(retry, hardObstacles, opt); ok {
		return fallback
	}
	// A point-only path is the defensive terminal state when malformed input
	// places an endpoint inside an exclusion zone. Callers do not draw it, and
	// the hard no-entry postcondition remains stronger than connector recovery.
	source, _, _, _ := routeEndpointStubsV1EngineRouteHardObstacle(retry, opt)
	return []ptV1EngineRouteTypes{source}
}

func horizontalStripSafeFallbackV1EngineRouteBuild(req routeRequestV1EngineRouteTypes, hardObstacles []rectV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) ([]ptV1EngineRouteTypes, bool) {
	source, sourceStub, destination, destinationStub := routeEndpointStubsV1EngineRouteHardObstacle(req, opt)

	lowerBound := math.Inf(-1)
	upperBound := math.Inf(1)
	matched := false
	for _, obstacle := range hardObstacles {
		right := obstacle.X + obstacle.W
		if source.X < obstacle.X || source.X > right || destination.X < obstacle.X || destination.X > right {
			continue
		}
		bottom := obstacle.Y + obstacle.H
		switch {
		case source.Y >= bottom && destination.Y >= bottom:
			lowerBound = math.Max(lowerBound, bottom)
			matched = true
		case source.Y <= obstacle.Y && destination.Y <= obstacle.Y:
			upperBound = math.Min(upperBound, obstacle.Y)
			matched = true
		default:
			return nil, false
		}
	}
	if !matched || lowerBound > upperBound {
		return nil, false
	}
	railY := (sourceStub.Y + destinationStub.Y) / 2
	if !math.IsInf(lowerBound, -1) {
		if sourceStub.Y < lowerBound || destinationStub.Y < lowerBound {
			return nil, false
		}
		railY = math.Max(math.Max(sourceStub.Y, destinationStub.Y), lowerBound)
	}
	if !math.IsInf(upperBound, 1) {
		if sourceStub.Y > upperBound || destinationStub.Y > upperBound {
			return nil, false
		}
		railY = math.Min(math.Min(sourceStub.Y, destinationStub.Y), upperBound)
	}
	points := []ptV1EngineRouteTypes{source, sourceStub}
	points = appendOrthogonalLegV1EngineRoutePath(points, sourceStub, ptV1EngineRouteTypes{X: sourceStub.X, Y: railY})
	points = appendOrthogonalLegV1EngineRoutePath(points, points[len(points)-1], ptV1EngineRouteTypes{X: destinationStub.X, Y: railY})
	points = appendOrthogonalLegV1EngineRoutePath(points, points[len(points)-1], destinationStub)
	points = append(points, destination)
	return enforceOrthogonalPolylineV1EngineRoutePath(points), true
}

func matchingRoutePathV1EngineRouteBuild(req routeRequestV1EngineRouteTypes, routePaths map[string][]ptV1EngineRouteTypes) ([]ptV1EngineRouteTypes, bool) {
	if points, ok := routePaths[routePairKeyV1EngineRouteBuild(req, false)]; ok {
		return points, true
	}
	if points, ok := routePaths[routePairKeyV1EngineRouteBuild(req, true)]; ok {
		return reversePointsV1EngineRoutePath(points), true
	}
	return nil, false
}

func routePairKeyV1EngineRouteBuild(req routeRequestV1EngineRouteTypes, reversed bool) string {
	src, dst := req.Src, req.Dst
	srcSide, dstSide := req.SrcSide, req.DstSide
	if reversed {
		src, dst = dst, src
		srcSide, dstSide = dstSide, srcSide
	}
	return rectKeyV1EngineRouteBuild(src) + "|" + string(srcSide) + ">" + rectKeyV1EngineRouteBuild(dst) + "|" + string(dstSide)
}

func rectKeyV1EngineRouteBuild(r rectV1EngineRouteTypes) string {
	return fmtFloatV1EngineRouteBuild(r.X) + "," + fmtFloatV1EngineRouteBuild(r.Y) + "," + fmtFloatV1EngineRouteBuild(r.W) + "," + fmtFloatV1EngineRouteBuild(r.H)
}

func fmtFloatV1EngineRouteBuild(v float64) string {
	return strconv.FormatFloat(math.Round(v*1000)/1000, 'f', 3, 64)
}

func trafficAlongsideRouteV1EngineRouteBuild(route, current []ptV1EngineRouteTypes, laneGap float64) []ptV1EngineRouteTypes {
	if len(route) < 2 || len(current) < 2 || laneGap <= 0 {
		return current
	}
	shifted := offsetPolylineV1EngineRouteBuild(route, laneGap)
	out := make([]ptV1EngineRouteTypes, 0, len(shifted)+2)
	out = append(out, current[0])
	if len(shifted) > 0 {
		out = appendOrthogonalLegV1EngineRoutePath(out, current[0], shifted[0])
		out = append(out, shifted[1:]...)
		out = appendOrthogonalLegV1EngineRoutePath(out, shifted[len(shifted)-1], current[len(current)-1])
	}
	return simplifyRouteCandidateV1EngineRoutePath(out)
}

func restoreDestinationApproachV1EngineRouteBuild(points []ptV1EngineRouteTypes, dstSide sideV1EngineRouteTypes, distance float64) []ptV1EngineRouteTypes {
	if len(points) < 2 || distance <= 0 {
		return points
	}
	end := points[len(points)-1]
	prev := points[len(points)-2]
	if approachesEndpointFromSideV1EngineRouteBuild(prev, end, dstSide) {
		return points
	}
	approach := extendV1EngineRouteGeometry(end, dstSide, distance)
	out := append([]ptV1EngineRouteTypes(nil), points[:len(points)-1]...)
	out = appendOrthogonalLegV1EngineRoutePath(out, out[len(out)-1], approach)
	out = append(out, end)
	return simplifyRouteCandidateV1EngineRoutePath(out)
}

func approachesEndpointFromSideV1EngineRouteBuild(prev, end ptV1EngineRouteTypes, dstSide sideV1EngineRouteTypes) bool {
	switch dstSide {
	case sideTopV1EngineRouteTypes:
		return math.Abs(prev.X-end.X) < epsV1EngineRouteTypes && prev.Y < end.Y-epsV1EngineRouteTypes
	case sideBottomV1EngineRouteTypes:
		return math.Abs(prev.X-end.X) < epsV1EngineRouteTypes && prev.Y > end.Y+epsV1EngineRouteTypes
	case sideLeftV1EngineRouteTypes:
		return math.Abs(prev.Y-end.Y) < epsV1EngineRouteTypes && prev.X < end.X-epsV1EngineRouteTypes
	case sideRightV1EngineRouteTypes:
		return math.Abs(prev.Y-end.Y) < epsV1EngineRouteTypes && prev.X > end.X+epsV1EngineRouteTypes
	default:
		return false
	}
}

func offsetPolylineV1EngineRouteBuild(points []ptV1EngineRouteTypes, offset float64) []ptV1EngineRouteTypes {
	out := make([]ptV1EngineRouteTypes, len(points))
	for i := range points {
		normal := vertexNormalV1EngineRouteBuild(points, i)
		out[i] = ptV1EngineRouteTypes{X: points[i].X + normal.X*offset, Y: points[i].Y + normal.Y*offset}
	}
	return out
}

func vertexNormalV1EngineRouteBuild(points []ptV1EngineRouteTypes, i int) ptV1EngineRouteTypes {
	var seg segmentV1EngineRouteTypes
	switch {
	case i == 0:
		seg = segmentV1EngineRouteTypes{A: points[0], B: points[1]}
	case i == len(points)-1:
		seg = segmentV1EngineRouteTypes{A: points[len(points)-2], B: points[len(points)-1]}
	default:
		before := segmentV1EngineRouteTypes{A: points[i-1], B: points[i]}
		after := segmentV1EngineRouteTypes{A: points[i], B: points[i+1]}
		if segmentLengthV1EngineRouteBuild(after) >= segmentLengthV1EngineRouteBuild(before) {
			seg = after
		} else {
			seg = before
		}
	}
	dx := seg.B.X - seg.A.X
	dy := seg.B.Y - seg.A.Y
	if math.Abs(dx) >= math.Abs(dy) {
		return ptV1EngineRouteTypes{X: 0, Y: 1}
	}
	return ptV1EngineRouteTypes{X: -1, Y: 0}
}

func segmentLengthV1EngineRouteBuild(s segmentV1EngineRouteTypes) float64 {
	return math.Hypot(s.B.X-s.A.X, s.B.Y-s.A.Y)
}
