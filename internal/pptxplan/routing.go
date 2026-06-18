package pptxplan

import (
	"math"
	"sort"
)

// routing.go — orthogonal arrow routing core (ported from the former TS
// routing.ts). Pure, deterministic geometry in layout-pixel space: given anchor
// rectangles, exit/entry sides and obstacle rectangles, it produces a polyline
// that avoids obstacles, prefers empty gutters, minimises crossings and
// near-parallel intrusions, and snaps bends to the grid.

type pt struct{ X, Y float64 }

type rect struct{ X, Y, W, H float64 }

type side string

const (
	sideTop    side = "top"
	sideBottom side = "bottom"
	sideLeft   side = "left"
	sideRight  side = "right"
)

type routeRequest struct {
	ID        string
	Src       rect
	Dst       rect
	SrcSide   side
	DstSide   side
	SrcAnchor *pt
	DstAnchor *pt
	SrcGap    float64
	DstGap    float64
}

type routedPath struct {
	ID     string
	Points []pt
}

type segment struct{ A, B pt }

type routerOptions struct {
	Grid       float64
	Clearance  float64
	Stub       float64
	LaneGap    float64
	LineMargin float64
}

func defaultRouterOptions() routerOptions {
	return routerOptions{Grid: 8, Clearance: 12, Stub: 20, LaneGap: 8, LineMargin: 8}
}

const (
	alignTol = 4.0
	eps      = 1e-6

	wObstacle  = 1000.0
	wCross     = 20.0
	wOverlap   = 40.0
	wProximity = 24.0
	wLen       = 5.0
	wBend      = 8.0
)

// routeConnections routes every request against a shared obstacle set. Requests
// are processed in order; each finalised path informs later ones.
func routeConnections(requests []routeRequest, obstacles []rect, opt routerOptions) []routedPath {
	placed := [][]segment{}
	results := make([]routedPath, 0, len(requests))
	for _, req := range requests {
		local := filterObstacles(obstacles, req)
		path := routeOne(req, local, placed, opt)
		results = append(results, path)
		placed = append(placed, toSegments(path.Points))
	}
	return results
}

func filterObstacles(obstacles []rect, req routeRequest) []rect {
	out := make([]rect, 0, len(obstacles))
	for _, o := range obstacles {
		if sameRect(o, req.Src) || sameRect(o, req.Dst) {
			continue
		}
		out = append(out, o)
	}
	return out
}

func routeOne(req routeRequest, obstacles []rect, placed [][]segment, opt routerOptions) routedPath {
	inflated := make([]rect, len(obstacles))
	for i, o := range obstacles {
		inflated[i] = inflate(o, opt.Clearance)
	}
	s := edgeMidpoint(req.Src, req.SrcSide)
	if req.SrcAnchor != nil {
		s = *req.SrcAnchor
	}
	d := edgeMidpoint(req.Dst, req.DstSide)
	if req.DstAnchor != nil {
		d = *req.DstAnchor
	}
	if req.SrcGap > 0 {
		s = extend(s, req.SrcSide, req.SrcGap)
	}
	if req.DstGap > 0 {
		d = extend(d, req.DstSide, req.DstGap)
	}
	s2 := extend(s, req.SrcSide, opt.Stub)
	d2 := extend(d, req.DstSide, opt.Stub)

	candidates := buildCandidates(s, d, s2, d2, inflated, placed, opt)

	var best []pt
	bestCost := math.Inf(1)
	foundClean := false
	for _, raw := range candidates {
		points := simplify(raw)
		if len(points) < 2 {
			continue
		}
		hits := obstacleHitCount(points, inflated)
		if hits > 0 && foundClean {
			continue
		}
		cost := scorePath(points, inflated, placed, opt.LineMargin)
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
		return routedPath{ID: req.ID, Points: []pt{s, d}}
	}
	return routedPath{ID: req.ID, Points: best}
}

func buildCandidates(s, d, s2, d2 pt, inflated []rect, placed [][]segment, opt routerOptions) [][]pt {
	candidates := [][]pt{}

	// 1) Straight only when endpoints share an axis exactly.
	if math.Abs(s.X-d.X) < eps || math.Abs(s.Y-d.Y) < eps {
		candidates = append(candidates, []pt{s, d})
	}

	// 2) L-shaped (single bend) between the stub endpoints.
	candidates = append(candidates, []pt{s, s2, {X: d2.X, Y: s2.Y}, d2, d})
	candidates = append(candidates, []pt{s, s2, {X: s2.X, Y: d2.Y}, d2, d})

	// 3) Z-shaped (two bends) through candidate trunk lines.
	midX := snap((s2.X+d2.X)/2, opt.Grid)
	midY := snap((s2.Y+d2.Y)/2, opt.Grid)

	obstacleXEdges := []float64{}
	obstacleYEdges := []float64{}
	for _, r := range inflated {
		obstacleXEdges = append(obstacleXEdges, r.X, r.X+r.W)
		obstacleYEdges = append(obstacleYEdges, r.Y, r.Y+r.H)
	}

	xRaw := []float64{midX, midX - opt.LaneGap, midX + opt.LaneGap, midX - 2*opt.LaneGap, midX + 2*opt.LaneGap}
	xRaw = append(xRaw, gutterCenters(projectX(inflated), math.Min(s2.X, d2.X), math.Max(s2.X, d2.X))...)
	xRaw = append(xRaw, obstacleXEdges...)
	xRaw = append(xRaw, placedLaneCoords(placed, false, opt)...)
	for i := range xRaw {
		xRaw[i] = snap(xRaw[i], opt.Grid)
	}
	xLines := dedupe(xRaw)

	yRaw := []float64{midY, midY - opt.LaneGap, midY + opt.LaneGap, midY - 2*opt.LaneGap, midY + 2*opt.LaneGap}
	yRaw = append(yRaw, gutterCenters(projectY(inflated), math.Min(s2.Y, d2.Y), math.Max(s2.Y, d2.Y))...)
	yRaw = append(yRaw, obstacleYEdges...)
	yRaw = append(yRaw, placedLaneCoords(placed, true, opt)...)
	for i := range yRaw {
		yRaw[i] = snap(yRaw[i], opt.Grid)
	}
	yLines := dedupe(yRaw)

	for _, x := range xLines {
		candidates = append(candidates, []pt{s, s2, {X: x, Y: s2.Y}, {X: x, Y: d2.Y}, d2, d})
	}
	for _, y := range yLines {
		candidates = append(candidates, []pt{s, s2, {X: s2.X, Y: y}, {X: d2.X, Y: y}, d2, d})
	}
	return candidates
}

func placedLaneCoords(placed [][]segment, horizontal bool, opt routerOptions) []float64 {
	offsets := []float64{opt.LineMargin, opt.LineMargin + opt.LaneGap, opt.LineMargin + 2*opt.LaneGap}
	out := []float64{}
	for _, path := range placed {
		for _, seg := range path {
			if isHorizontal(seg) != horizontal {
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

func scorePath(points []pt, inflated []rect, placed [][]segment, margin float64) float64 {
	segs := toSegments(points)

	obstacleHits := obstacleHitCount(points, inflated)

	crossings := 0.0
	overlap := 0.0
	proximity := 0.0
	for _, other := range placed {
		for _, seg := range segs {
			for _, oseg := range other {
				if segmentsCross(seg, oseg) {
					crossings++
				}
				overlap += collinearOverlap(seg, oseg)
				proximity += nearParallel(seg, oseg, margin)
			}
		}
	}

	length := pathLength(points)
	bends := math.Max(0, float64(len(points))-2)

	overlapTerm := 0.0
	if overlap > 0 {
		overlapTerm = 1 + overlap/math.Max(1, margin)
	}
	proximityTerm := 0.0
	if proximity > 0 {
		proximityTerm = 1 + proximity/100
	}

	return wObstacle*obstacleHits +
		wCross*crossings +
		wOverlap*overlapTerm +
		wProximity*proximityTerm +
		wLen*(length/1000) +
		wBend*bends
}

func obstacleHitCount(points []pt, inflated []rect) float64 {
	hits := 0.0
	for _, seg := range toSegments(points) {
		for _, r := range inflated {
			if segIntersectsRect(seg, r) {
				hits++
			}
		}
	}
	return hits
}

// ── Geometry helpers ─────────────────────────────────────────────────────────

func edgeMidpoint(r rect, s side) pt {
	cx := r.X + r.W/2
	cy := r.Y + r.H/2
	switch s {
	case sideTop:
		return pt{X: cx, Y: r.Y}
	case sideBottom:
		return pt{X: cx, Y: r.Y + r.H}
	case sideLeft:
		return pt{X: r.X, Y: cy}
	default:
		return pt{X: r.X + r.W, Y: cy}
	}
}

func extend(p pt, s side, d float64) pt {
	switch s {
	case sideTop:
		return pt{X: p.X, Y: p.Y - d}
	case sideBottom:
		return pt{X: p.X, Y: p.Y + d}
	case sideLeft:
		return pt{X: p.X - d, Y: p.Y}
	default:
		return pt{X: p.X + d, Y: p.Y}
	}
}

func inflate(r rect, c float64) rect {
	return rect{X: r.X - c, Y: r.Y - c, W: r.W + 2*c, H: r.H + 2*c}
}

func sameRect(a, b rect) bool {
	return math.Abs(a.X-b.X) < eps &&
		math.Abs(a.Y-b.Y) < eps &&
		math.Abs(a.W-b.W) < eps &&
		math.Abs(a.H-b.H) < eps
}

func toSegments(points []pt) []segment {
	segs := make([]segment, 0, len(points))
	for i := 0; i < len(points)-1; i++ {
		segs = append(segs, segment{A: points[i], B: points[i+1]})
	}
	return segs
}

func snap(v, grid float64) float64 {
	return math.Round(v/grid) * grid
}

// simplify removes duplicate and collinear interior vertices.
func simplify(points []pt) []pt {
	out := []pt{}
	for _, p := range points {
		if len(out) > 0 {
			last := out[len(out)-1]
			if math.Abs(last.X-p.X) < eps && math.Abs(last.Y-p.Y) < eps {
				continue
			}
		}
		out = append(out, p)
	}
	i := 1
	for i < len(out)-1 {
		a, b, c := out[i-1], out[i], out[i+1]
		collinearH := math.Abs(a.Y-b.Y) < eps && math.Abs(b.Y-c.Y) < eps
		collinearV := math.Abs(a.X-b.X) < eps && math.Abs(b.X-c.X) < eps
		if collinearH || collinearV {
			out = append(out[:i], out[i+1:]...)
		} else {
			i++
		}
	}
	return out
}

func pathLength(points []pt) float64 {
	total := 0.0
	for i := 0; i < len(points)-1; i++ {
		total += math.Abs(points[i+1].X-points[i].X) + math.Abs(points[i+1].Y-points[i].Y)
	}
	return total
}

func isHorizontal(s segment) bool {
	return math.Abs(s.A.Y-s.B.Y) < eps
}

func segIntersectsRect(seg segment, r rect) bool {
	rx1, rx2 := r.X, r.X+r.W
	ry1, ry2 := r.Y, r.Y+r.H
	if isHorizontal(seg) {
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

func segmentsCross(p, q segment) bool {
	pH := isHorizontal(p)
	qH := isHorizontal(q)
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
	return vx > hx1+eps && vx < hx2-eps && hy > vy1+eps && hy < vy2-eps
}

func collinearOverlap(p, q segment) float64 {
	if isHorizontal(p) != isHorizontal(q) {
		return 0
	}
	if isHorizontal(p) {
		if math.Abs(p.A.Y-q.A.Y) > eps {
			return 0
		}
		a1 := math.Min(p.A.X, p.B.X)
		a2 := math.Max(p.A.X, p.B.X)
		b1 := math.Min(q.A.X, q.B.X)
		b2 := math.Max(q.A.X, q.B.X)
		return math.Max(0, math.Min(a2, b2)-math.Max(a1, b1))
	}
	if math.Abs(p.A.X-q.A.X) > eps {
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
func nearParallel(p, q segment, margin float64) float64 {
	if margin <= 0 {
		return 0
	}
	if isHorizontal(p) != isHorizontal(q) {
		return 0
	}
	var gap, axialOverlap float64
	if isHorizontal(p) {
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
	if gap <= eps || gap >= margin {
		return 0
	}
	return axialOverlap * (1 - gap/margin)
}

// ── Gutter extraction ────────────────────────────────────────────────────────

func projectX(rects []rect) [][2]float64 {
	out := make([][2]float64, len(rects))
	for i, r := range rects {
		out[i] = [2]float64{r.X, r.X + r.W}
	}
	return out
}

func projectY(rects []rect) [][2]float64 {
	out := make([][2]float64, len(rects))
	for i, r := range rects {
		out[i] = [2]float64{r.Y, r.Y + r.H}
	}
	return out
}

func gutterCenters(intervals [][2]float64, lo, hi float64) []float64 {
	if hi-lo < eps {
		return nil
	}
	sorted := make([][2]float64, len(intervals))
	copy(sorted, intervals)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i][0] < sorted[j][0] })
	merged := [][2]float64{}
	for _, iv := range sorted {
		if len(merged) > 0 && iv[0] <= merged[len(merged)-1][1] {
			if iv[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = iv[1]
			}
		} else {
			merged = append(merged, iv)
		}
	}
	centers := []float64{}
	cursor := lo
	for _, m := range merged {
		start, end := m[0], m[1]
		if start > cursor {
			gapStart := math.Max(cursor, lo)
			gapEnd := math.Min(start, hi)
			if gapEnd > gapStart {
				centers = append(centers, (gapStart+gapEnd)/2)
			}
		}
		cursor = math.Max(cursor, end)
	}
	if cursor < hi {
		centers = append(centers, (math.Max(cursor, lo)+hi)/2)
	}
	return centers
}

func dedupe(values []float64) []float64 {
	out := []float64{}
	for _, v := range values {
		dup := false
		for _, u := range out {
			if math.Abs(u-v) < alignTol {
				dup = true
				break
			}
		}
		if !dup {
			out = append(out, v)
		}
	}
	return out
}
