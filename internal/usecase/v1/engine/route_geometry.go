package engine

import "math"

// ── Geometry helpers ─────────────────────────────────────────────────────────

func edgeMidpointV1EngineRouteGeometry(r rectV1EngineRouteTypes, s sideV1EngineRouteTypes) ptV1EngineRouteTypes {
	cx := r.X + r.W/2
	cy := r.Y + r.H/2
	switch s {
	case sideTopV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: cx, Y: r.Y}
	case sideBottomV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: cx, Y: r.Y + r.H}
	case sideLeftV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: r.X, Y: cy}
	default:
		return ptV1EngineRouteTypes{X: r.X + r.W, Y: cy}
	}
}

func extendV1EngineRouteGeometry(p ptV1EngineRouteTypes, s sideV1EngineRouteTypes, d float64) ptV1EngineRouteTypes {
	switch s {
	case sideTopV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: p.X, Y: p.Y - d}
	case sideBottomV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: p.X, Y: p.Y + d}
	case sideLeftV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: p.X - d, Y: p.Y}
	default:
		return ptV1EngineRouteTypes{X: p.X + d, Y: p.Y}
	}
}

func inflateV1EngineRouteGeometry(r rectV1EngineRouteTypes, c float64) rectV1EngineRouteTypes {
	return rectV1EngineRouteTypes{X: r.X - c, Y: r.Y - c, W: r.W + 2*c, H: r.H + 2*c}
}

func sameRectV1EngineRouteGeometry(a, b rectV1EngineRouteTypes) bool {
	return math.Abs(a.X-b.X) < epsV1EngineRouteTypes &&
		math.Abs(a.Y-b.Y) < epsV1EngineRouteTypes &&
		math.Abs(a.W-b.W) < epsV1EngineRouteTypes &&
		math.Abs(a.H-b.H) < epsV1EngineRouteTypes
}

func toSegmentsV1EngineRouteGeometry(points []ptV1EngineRouteTypes) []segmentV1EngineRouteTypes {
	segs := make([]segmentV1EngineRouteTypes, 0, len(points))
	for i := 0; i < len(points)-1; i++ {
		segs = append(segs, segmentV1EngineRouteTypes{A: points[i], B: points[i+1]})
	}
	return segs
}

func snapV1EngineRouteGeometry(v, grid float64) float64 {
	return math.Round(v/grid) * grid
}

// simplify removes duplicate and collinear interior vertices.
func simplifyV1EngineRouteGeometry(points []ptV1EngineRouteTypes) []ptV1EngineRouteTypes {
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

func pathLengthV1EngineRouteGeometry(points []ptV1EngineRouteTypes) float64 {
	total := 0.0
	for i := 0; i < len(points)-1; i++ {
		total += math.Abs(points[i+1].X-points[i].X) + math.Abs(points[i+1].Y-points[i].Y)
	}
	return total
}

func isHorizontalV1EngineRouteGeometry(s segmentV1EngineRouteTypes) bool {
	return math.Abs(s.A.Y-s.B.Y) < epsV1EngineRouteTypes
}

func segIntersectsRectV1EngineRouteGeometry(seg segmentV1EngineRouteTypes, r rectV1EngineRouteTypes) bool {
	rx1, rx2 := r.X, r.X+r.W
	ry1, ry2 := r.Y, r.Y+r.H
	if isHorizontalV1EngineRouteGeometry(seg) {
		y := seg.A.Y
		if y < ry1 || y > ry2 {
			return false
		}
		x1 := math.Min(seg.A.X, seg.B.X)
		x2 := math.Max(seg.A.X, seg.B.X)
		return x2 >= rx1 && x1 <= rx2
	}
	x := seg.A.X
	if x < rx1 || x > rx2 {
		return false
	}
	y1 := math.Min(seg.A.Y, seg.B.Y)
	y2 := math.Max(seg.A.Y, seg.B.Y)
	return y2 >= ry1 && y1 <= ry2
}

func segmentsCrossV1EngineRouteGeometry(p, q segmentV1EngineRouteTypes) bool {
	pH := isHorizontalV1EngineRouteGeometry(p)
	qH := isHorizontalV1EngineRouteGeometry(q)
	if pH == qH {
		return false
	}
	h, v := p, q
	if !pH {
		h, v = q, p
	}
	hy := h.A.Y
	vx := v.A.X
	hx1 := math.Min(h.A.X, h.B.X)
	hx2 := math.Max(h.A.X, h.B.X)
	vy1 := math.Min(v.A.Y, v.B.Y)
	vy2 := math.Max(v.A.Y, v.B.Y)
	return vx > hx1+epsV1EngineRouteTypes && vx < hx2-epsV1EngineRouteTypes && hy > vy1+epsV1EngineRouteTypes && hy < vy2-epsV1EngineRouteTypes
}

// crossingPoint returns the interior intersection of two orthogonal segments.
// Endpoint touches and collinear overlaps are deliberately not line jumps.
func crossingPointV1EngineRouteGeometry(p, q segmentV1EngineRouteTypes) (ptV1EngineRouteTypes, bool) {
	if !segmentsCrossV1EngineRouteGeometry(p, q) {
		return ptV1EngineRouteTypes{}, false
	}
	h, v := p, q
	if !isHorizontalV1EngineRouteGeometry(p) {
		h, v = q, p
	}
	return ptV1EngineRouteTypes{X: v.A.X, Y: h.A.Y}, true
}

// pathCrossings returns deduplicated intersections between the path drawn on
// top and all previously drawn paths below it.
func pathCrossingsV1EngineRouteGeometry(top routedPathV1EngineRouteTypes, below []routedPathV1EngineRouteTypes) []ptV1EngineRouteTypes {
	var out []ptV1EngineRouteTypes
	for _, topSeg := range toSegmentsV1EngineRouteGeometry(top.Points) {
		for _, lower := range below {
			for _, lowerSeg := range toSegmentsV1EngineRouteGeometry(lower.Points) {
				p, ok := crossingPointV1EngineRouteGeometry(topSeg, lowerSeg)
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
	}
	return out
}

func collinearOverlapV1EngineRouteGeometry(p, q segmentV1EngineRouteTypes) float64 {
	if isHorizontalV1EngineRouteGeometry(p) != isHorizontalV1EngineRouteGeometry(q) {
		return 0
	}
	if isHorizontalV1EngineRouteGeometry(p) {
		if math.Abs(p.A.Y-q.A.Y) > epsV1EngineRouteTypes {
			return 0
		}
		a1 := math.Min(p.A.X, p.B.X)
		a2 := math.Max(p.A.X, p.B.X)
		b1 := math.Min(q.A.X, q.B.X)
		b2 := math.Max(q.A.X, q.B.X)
		return math.Max(0, math.Min(a2, b2)-math.Max(a1, b1))
	}
	if math.Abs(p.A.X-q.A.X) > epsV1EngineRouteTypes {
		return 0
	}
	a1 := math.Min(p.A.Y, p.B.Y)
	a2 := math.Max(p.A.Y, p.B.Y)
	b1 := math.Min(q.A.Y, q.B.Y)
	b2 := math.Max(q.A.Y, q.B.Y)
	return math.Max(0, math.Min(a2, b2)-math.Max(a1, b1))
}

// nearParallel penalises two parallel co-axial segments running within margin of
// each other (but not exactly collinear, which collinearOverlap handles).
func nearParallelV1EngineRouteGeometry(p, q segmentV1EngineRouteTypes, margin float64) float64 {
	if margin <= 0 {
		return 0
	}
	if isHorizontalV1EngineRouteGeometry(p) != isHorizontalV1EngineRouteGeometry(q) {
		return 0
	}
	var gap, axialOverlap float64
	if isHorizontalV1EngineRouteGeometry(p) {
		gap = math.Abs(p.A.Y - q.A.Y)
		a1 := math.Min(p.A.X, p.B.X)
		a2 := math.Max(p.A.X, p.B.X)
		b1 := math.Min(q.A.X, q.B.X)
		b2 := math.Max(q.A.X, q.B.X)
		axialOverlap = math.Max(0, math.Min(a2, b2)-math.Max(a1, b1))
	} else {
		gap = math.Abs(p.A.X - q.A.X)
		a1 := math.Min(p.A.Y, p.B.Y)
		a2 := math.Max(p.A.Y, p.B.Y)
		b1 := math.Min(q.A.Y, q.B.Y)
		b2 := math.Max(q.A.Y, q.B.Y)
		axialOverlap = math.Max(0, math.Min(a2, b2)-math.Max(a1, b1))
	}
	if axialOverlap <= 0 {
		return 0
	}
	if gap <= epsV1EngineRouteTypes || gap >= margin {
		return 0
	}
	return axialOverlap * (1 - gap/margin)
}
