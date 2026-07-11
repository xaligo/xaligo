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
		path := routeOneV1EngineRoutePath(req, local, placed, opt)
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
		results = append(results, path)
		if req.Kind == "route" {
			routePaths[routePairKeyV1EngineRouteBuild(req, false)] = append([]ptV1EngineRouteTypes(nil), path.Points...)
		}
		placed = append(placed, toSegmentsV1EngineRouteGeometry(path.Points))
	}
	return results
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
