package pptxplan

import (
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
)

// plan.go — builds a fully-resolved PPTX draw plan from an Excalidraw scene.
// This is the Go port of the calculation half of the former TS pptx.ts.

const (
	defaultPxPerInch = 96.0
	anchorGrid       = 5
	// Keep the mask smaller than the default 8 px lane gap so a jump does not
	// accidentally erase a nearby parallel lane.
	lineJumpSizePx = 6.0
)

// paperSizeIn lists portrait paper dimensions in inches (width × height).
// Landscape swaps them.
var paperSizeIn = map[string][2]float64{
	"A5":      {5.83, 8.27},
	"A4":      {8.27, 11.69},
	"A3":      {11.69, 16.54},
	"A2":      {16.54, 23.39},
	"A1":      {23.39, 33.11},
	"Letter":  {8.5, 11},
	"Legal":   {8.5, 14},
	"Tabloid": {11, 17},
}

var hexFull = regexp.MustCompile(`^#?[0-9a-fA-F]{6}$`)

// BuildPlanJSON parses an Excalidraw JSON scene and returns the PPTX draw plan
// as JSON, applying every geometry option in opt.
func BuildPlanJSON(sceneJSON string, opt Options) ([]byte, error) {
	var scene Scene
	if err := json.Unmarshal([]byte(sceneJSON), &scene); err != nil {
		return nil, err
	}
	plan := BuildPlan(&scene, opt)
	return json.Marshal(plan)
}

// connectorStyle is the resolved arrowhead + width for connectors.
type connectorStyle struct {
	Head     string
	Width    float64
	HasWidth bool
}

func resolveConnectorStyle(style string) connectorStyle {
	switch style {
	case "standard":
		return connectorStyle{Head: "triangle", Width: 1.5, HasWidth: true}
	case "triangle":
		return connectorStyle{Head: "triangle"}
	case "stealth":
		return connectorStyle{Head: "stealth"}
	case "arrow":
		return connectorStyle{Head: "arrow"}
	case "diamond":
		return connectorStyle{Head: "diamond"}
	case "oval":
		return connectorStyle{Head: "oval"}
	case "none":
		return connectorStyle{Head: "none"}
	default: // "thin" and unset → slender stealth head on a thin line
		return connectorStyle{Head: "stealth", Width: 1, HasWidth: true}
	}
}

// resolvePaper returns the target slide size (inches) for a named paper size,
// auto-selecting the orientation that fits the diagram best when not forced.
func resolvePaper(size, orientation string, contentWIn, contentHIn float64) (w, h float64, ok bool) {
	base, found := paperSizeIn[size]
	if !found {
		return 0, 0, false
	}
	pw, ph := base[0], base[1]
	lw, lh := base[1], base[0]
	switch orientation {
	case "portrait":
		return pw, ph, true
	case "landscape":
		return lw, lh, true
	}
	scaleP := math.Min(pw/contentWIn, ph/contentHIn)
	scaleL := math.Min(lw/contentWIn, lh/contentHIn)
	if scaleL >= scaleP {
		return lw, lh, true
	}
	return pw, ph, true
}

// BuildPlan converts a parsed scene into a draw plan.
func BuildPlan(scene *Scene, opt Options) Plan {
	elements := make([]*Element, 0, len(scene.Elements))
	for i := range scene.Elements {
		if !scene.Elements[i].IsDeleted {
			elements = append(elements, &scene.Elements[i])
		}
	}

	basePPI := opt.PxPerInch
	if basePPI <= 0 {
		basePPI = defaultPxPerInch
	}
	contentFrame := findPaperFrame(elements)
	if contentFrame == nil {
		cb := contentBounds(elements)
		contentFrame = &cb
	}

	contentWIn := contentFrame.W / basePPI
	contentHIn := contentFrame.H / basePPI
	paperW, paperH, hasPaper := resolvePaper(opt.PaperSize, opt.Orientation, contentWIn, contentHIn)

	frame := *contentFrame
	ppi := basePPI
	layoutW := contentWIn
	layoutH := contentHIn
	if hasPaper {
		scale := math.Min(paperW/contentWIn, paperH/contentHIn)
		ppi = basePPI / scale
		offsetXIn := (paperW - contentWIn*scale) / 2
		offsetYIn := (paperH - contentHIn*scale) / 2
		frame = rect{
			X: contentFrame.X - offsetXIn*ppi,
			Y: contentFrame.Y - offsetYIn*ppi,
			W: contentFrame.W,
			H: contentFrame.H,
		}
		layoutW = paperW
		layoutH = paperH
	}

	style := resolveConnectorStyle(opt.ArrowStyle)
	background := normalizeColor(backgroundColor(scene), "FFFFFF")
	stubPx := 20.0
	if opt.ArrowStubPx > 0 {
		stubPx = opt.ArrowStubPx
	}
	marginPx := 8.0
	if opt.ArrowMargin > 0 {
		marginPx = opt.ArrowMargin
	}

	elementsByID := map[string]*Element{}
	for _, el := range elements {
		if el.ID != "" {
			elementsByID[el.ID] = el
		}
	}

	obstacles := collectObstacles(elements)

	connectors := []*Element{}
	for _, el := range elements {
		if el.Type == "arrow" || el.Type == "line" {
			connectors = append(connectors, el)
		}
	}
	prepared := prepareConnectors(connectors, elementsByID)

	ops := []DrawOp{}

	// 1) Anchor grids first → behind the icons drawn on top.
	gridIDs := make([]string, 0, len(prepared.gridRects))
	for id := range prepared.gridRects {
		gridIDs = append(gridIDs, id)
	}
	sort.Strings(gridIDs)
	for _, id := range gridIDs {
		ops = append(ops, anchorGridOps(prepared.gridRects[id], frame, ppi, background)...)
	}

	// 2) Containers/shapes in scene order.
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		if el.CustomData != nil && el.CustomData.Junction {
			continue
		}
		switch el.Type {
		case "frame", "rectangle", "ellipse":
			if op, ok := shapeOp(el, frame, ppi); ok {
				ops = append(ops, op)
			}
		}
	}

	// 3) Connectors above containers but below icons/labels.
	for _, el := range prepared.raw {
		if op, ok := rawLineOp(el, frame, ppi, style); ok {
			ops = append(ops, op)
		}
	}

	ordered := make([]preparedConnector, len(prepared.routed))
	copy(ordered, prepared.routed)
	sort.Slice(ordered, func(i, j int) bool {
		pi := connectorKindPriority(ordered[i].req.Kind)
		pj := connectorKindPriority(ordered[j].req.Kind)
		if pi != pj {
			return pi < pj
		}
		return ordered[i].req.ID < ordered[j].req.ID
	})
	reqs := make([]routeRequest, len(ordered))
	for i, pc := range ordered {
		reqs[i] = pc.req
	}
	junctions := applyRouteJunctions(reqs, stubPx)
	rOpt := defaultRouterOptions()
	rOpt.Stub = stubPx
	rOpt.LineMargin = marginPx
	rOpt.Reserved = collectContainerBorderPaths(elements)
	routed := routeConnections(reqs, obstacles, rOpt)
	elByConn := map[string]*Element{}
	for _, pc := range ordered {
		elByConn[pc.req.ID] = pc.el
	}
	for i, path := range routed {
		el := elByConn[path.ID]
		if el == nil {
			continue
		}
		for _, crossing := range pathCrossings(path, routed[:i]) {
			maskColor := lineJumpBackground(crossing, elements, background)
			ops = append(ops, lineJumpMaskOp(crossing, frame, ppi, maskColor))
		}
		if op, ok := polylineOp(el, path.Points, frame, ppi, style); ok {
			ops = append(ops, op)
		}
	}
	for _, junction := range junctions {
		el := elByConn[junction.ConnectorID]
		if el == nil {
			continue
		}
		ops = append(ops, junctionOp(junction.Point, frame, ppi, connectorLine(el, style)))
	}

	// 4) Icons and labels on top so routed lines never visually cover them.
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		switch el.Type {
		case "text":
			if op, ok := textOp(el, frame, ppi); ok {
				ops = append(ops, op)
			}
		case "image":
			if op, ok := imageOp(el, scene.Files, frame, ppi); ok {
				ops = append(ops, op)
			}
		}
	}

	return Plan{
		Slide: PlanSlide{
			W:          layoutW,
			H:          layoutH,
			Background: background,
		},
		Ops:    ops,
		Legend: buildLegend(scene, opt.LegendEntries),
	}
}

type routeJunction struct {
	Point       pt
	ConnectorID string
}

type junctionEndpoint struct {
	requestIndex int
	rect         rect
	side         side
	gap          float64
	source       bool
}

// applyRouteJunctions makes route fan-out/fan-in connections share a centered
// anchor and stub. The returned points are drawn after the route lines, making
// the branch visually explicit. Traffic and ordinary connections stay
// independent.
func applyRouteJunctions(requests []routeRequest, stub float64) []routeJunction {
	groups := map[string][]junctionEndpoint{}
	keys := []string{}
	add := func(key string, endpoint junctionEndpoint) {
		if _, exists := groups[key]; !exists {
			keys = append(keys, key)
		}
		groups[key] = append(groups[key], endpoint)
	}
	for i, req := range requests {
		if req.Kind != "route" {
			continue
		}
		add(junctionGroupKey("src", req.Src, req.SrcSide), junctionEndpoint{i, req.Src, req.SrcSide, req.SrcGap, true})
		add(junctionGroupKey("dst", req.Dst, req.DstSide), junctionEndpoint{i, req.Dst, req.DstSide, req.DstGap, false})
	}

	junctions := []routeJunction{}
	seen := map[string]bool{}
	for _, key := range keys {
		group := groups[key]
		if len(group) < 2 {
			continue
		}
		first := group[0]
		anchor := anchorPoint(first.rect, first.side, anchorGrid/2)
		for _, endpoint := range group {
			copyPoint := anchor
			if endpoint.source {
				requests[endpoint.requestIndex].SrcAnchor = &copyPoint
			} else {
				requests[endpoint.requestIndex].DstAnchor = &copyPoint
			}
		}
		point := anchor
		if first.gap > 0 {
			point = extend(point, first.side, first.gap)
		}
		point = extend(point, first.side, stub)
		pointKey := fmt.Sprintf("%.4f|%.4f", point.X, point.Y)
		if !seen[pointKey] {
			junctions = append(junctions, routeJunction{Point: point, ConnectorID: requests[first.requestIndex].ID})
			seen[pointKey] = true
		}
	}
	return junctions
}

func junctionGroupKey(prefix string, r rect, side side) string {
	return fmt.Sprintf("%s|%.4f|%.4f|%.4f|%.4f|%s", prefix, r.X, r.Y, r.W, r.H, side)
}

func junctionOp(point pt, frame rect, ppi float64, line LineStyle) DrawOp {
	const diameterPx = 8.0
	diameter := diameterPx / ppi
	return DrawOp{
		Kind: "ellipse",
		X:    (point.X-frame.X)/ppi - diameter/2,
		Y:    (point.Y-frame.Y)/ppi - diameter/2,
		W:    diameter,
		H:    diameter,
		Fill: &FillStyle{Color: line.Color, Transparency: line.Transparency},
		Line: &LineStyle{Color: line.Color, Width: math.Max(0.75, line.Width), Dash: "solid", Transparency: line.Transparency},
	}
}

func lineJumpMaskOp(crossing pt, frame rect, ppi float64, background string) DrawOp {
	size := lineJumpSizePx / ppi
	return DrawOp{
		Kind: "rect",
		X:    (crossing.X-frame.X)/ppi - size/2,
		Y:    (crossing.Y-frame.Y)/ppi - size/2,
		W:    size,
		H:    size,
		Fill: &FillStyle{Color: background, Transparency: 0},
		Line: &LineStyle{Color: background, Width: 0.25, Transparency: 100},
	}
}

// lineJumpBackground returns the uppermost opaque shape fill beneath a crossing.
// Transparent or partially transparent fills fall back to the slide background,
// as reproducing their composited color would require renderer-specific blending.
func lineJumpBackground(crossing pt, elements []*Element, fallback string) string {
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
			color = normalizeColor(el.BackgroundColor, color)
		}
	}
	return color
}

func buildLegend(scene *Scene, entries []LegendEntry) []LegendEntry {
	if scene == nil || len(entries) == 0 {
		return nil
	}
	out := make([]LegendEntry, 0, len(entries))
	seen := map[int]bool{}
	for _, entry := range entries {
		if entry.CatalogID <= 0 || seen[entry.CatalogID] {
			continue
		}
		file, ok := scene.Files[fmt.Sprintf("item-cat-%d", entry.CatalogID)]
		if !ok || file.DataURL == "" {
			continue
		}
		if entry.Abbreviation == "" {
			entry.Abbreviation = entry.OfficialName
		}
		entry.Data = file.DataURL
		out = append(out, entry)
		seen[entry.CatalogID] = true
	}
	return out
}

func backgroundColor(scene *Scene) string {
	if scene.AppState != nil {
		return scene.AppState.ViewBackgroundColor
	}
	return ""
}

// ── Connector preparation (anchors + sides) ──────────────────────────────────

type preparedConnector struct {
	el  *Element
	req routeRequest
}

type endpoint struct {
	connID    string
	rect      rect
	side      side
	oppCenter pt
	isSrc     bool
}

type preparedResult struct {
	routed    []preparedConnector
	raw       []*Element
	gridRects map[string]rect
}

func prepareConnectors(connectors []*Element, byID map[string]*Element) preparedResult {
	raw := []*Element{}
	gridRects := map[string]rect{}
	groupKeys := []string{}
	groups := map[string][]endpoint{}
	type item struct {
		el      *Element
		src     rect
		dst     rect
		srcSide side
		dstSide side
		srcGap  float64
		dstGap  float64
	}
	itemKeys := []string{}
	items := map[string]item{}

	pushGroup := func(iconID string, s side, ep endpoint) {
		key := iconID + "|" + string(s)
		if _, ok := groups[key]; !ok {
			groupKeys = append(groupKeys, key)
		}
		groups[key] = append(groups[key], ep)
	}

	for _, el := range connectors {
		srcIconID := ""
		if el.StartBinding != nil {
			srcIconID = el.StartBinding.ElementID
		}
		dstIconID := ""
		if el.EndBinding != nil {
			dstIconID = el.EndBinding.ElementID
		}
		src, srcOK := rectOf(byID[srcIconID])
		dst, dstOK := rectOf(byID[dstIconID])
		if !srcOK || !dstOK || el.ID == "" {
			raw = append(raw, el)
			continue
		}
		srcSide, dstSide := inferSides(src, dst)
		if el.StartBinding != nil {
			if s, ok := sideFromFixedPoint(el.StartBinding.FixedPoint); ok {
				srcSide = s
			}
		}
		if el.EndBinding != nil {
			if s, ok := sideFromFixedPoint(el.EndBinding.FixedPoint); ok {
				dstSide = s
			}
		}
		srcCenter := pt{X: src.X + src.W/2, Y: src.Y + src.H/2}
		dstCenter := pt{X: dst.X + dst.W/2, Y: dst.Y + dst.H/2}

		if _, ok := items[el.ID]; !ok {
			itemKeys = append(itemKeys, el.ID)
		}
		srcGap := 0.0
		if el.StartBinding != nil {
			srcGap = el.StartBinding.Gap
		}
		dstGap := 0.0
		if el.EndBinding != nil {
			dstGap = el.EndBinding.Gap
		}
		items[el.ID] = item{el: el, src: src, dst: dst, srcSide: srcSide, dstSide: dstSide, srcGap: srcGap, dstGap: dstGap}
		gridRects[srcIconID] = src
		gridRects[dstIconID] = dst
		pushGroup(srcIconID, srcSide, endpoint{connID: el.ID, rect: src, side: srcSide, oppCenter: dstCenter, isSrc: true})
		pushGroup(dstIconID, dstSide, endpoint{connID: el.ID, rect: dst, side: dstSide, oppCenter: srcCenter, isSrc: false})
	}

	anchors := map[string]*anchorPair{}
	for _, key := range groupKeys {
		assignGroupAnchors(groups[key], anchors)
	}

	routed := make([]preparedConnector, 0, len(itemKeys))
	for _, id := range itemKeys {
		it := items[id]
		req := routeRequest{
			ID:      id,
			Kind:    connectorKind(it.el),
			Src:     it.src,
			Dst:     it.dst,
			SrcSide: it.srcSide,
			DstSide: it.dstSide,
			SrcGap:  it.srcGap,
			DstGap:  it.dstGap,
		}
		if a := anchors[id]; a != nil {
			req.SrcAnchor = a.src
			req.DstAnchor = a.dst
		}
		routed = append(routed, preparedConnector{el: it.el, req: req})
	}
	return preparedResult{routed: routed, raw: raw, gridRects: gridRects}
}

type anchorPair = struct {
	src *pt
	dst *pt
}

func assignGroupAnchors(eps []endpoint, anchors map[string]*anchorPair) {
	if len(eps) == 0 {
		return
	}
	s := eps[0].side
	horizontal := s == sideTop || s == sideBottom
	sort.SliceStable(eps, func(i, j int) bool {
		if horizontal {
			return eps[i].oppCenter.X < eps[j].oppCenter.X
		}
		return eps[i].oppCenter.Y < eps[j].oppCenter.Y
	})
	meanFrac := 0.0
	for _, ep := range eps {
		meanFrac += edgeFraction(ep, horizontal)
	}
	meanFrac /= float64(len(eps))
	slots := assignSlots(len(eps), meanFrac)
	for k, ep := range eps {
		p := anchorPoint(ep.rect, s, slots[k])
		entry := anchors[ep.connID]
		if entry == nil {
			entry = &anchorPair{}
			anchors[ep.connID] = entry
		}
		pc := p
		if ep.isSrc {
			entry.src = &pc
		} else {
			entry.dst = &pc
		}
	}
}

func edgeFraction(ep endpoint, horizontal bool) float64 {
	var f float64
	if horizontal {
		w := ep.rect.W
		if w == 0 {
			w = 1
		}
		f = (ep.oppCenter.X - ep.rect.X) / w
	} else {
		h := ep.rect.H
		if h == 0 {
			h = 1
		}
		f = (ep.oppCenter.Y - ep.rect.Y) / h
	}
	return math.Max(0, math.Min(1, f))
}

func assignSlots(n int, frac float64) []int {
	slots := anchorGrid
	if n <= 0 {
		return nil
	}
	if n >= slots {
		out := make([]int, n)
		for k := 0; k < n; k++ {
			out[k] = int(math.Round(float64(k*(slots-1)) / float64(n-1)))
		}
		return out
	}
	center := frac * float64(slots-1)
	start := int(math.Round(center - float64(n-1)/2))
	if start < 0 {
		start = 0
	}
	if start > slots-n {
		start = slots - n
	}
	out := make([]int, n)
	for k := 0; k < n; k++ {
		out[k] = start + k
	}
	return out
}

func anchorPoint(r rect, s side, slot int) pt {
	colX := r.X + (float64(slot)+0.5)*(r.W/float64(anchorGrid))
	rowY := r.Y + (float64(slot)+0.5)*(r.H/float64(anchorGrid))
	switch s {
	case sideTop:
		return pt{X: colX, Y: r.Y}
	case sideBottom:
		return pt{X: colX, Y: r.Y + r.H}
	case sideLeft:
		return pt{X: r.X, Y: rowY}
	default:
		return pt{X: r.X + r.W, Y: rowY}
	}
}

// ── Obstacles + side inference ───────────────────────────────────────────────

func collectObstacles(elements []*Element) []rect {
	rects := []rect{}
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		if el.Type != "image" && el.Type != "text" {
			continue
		}
		r, ok := rectOf(el)
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
func collectContainerBorderPaths(elements []*Element) [][]segment {
	paths := make([][]segment, 0)
	for _, el := range elements {
		if el.ID == "paper-frame" || (el.Type != "frame" && el.Type != "rectangle") {
			continue
		}
		stroke := strings.ToLower(strings.TrimSpace(el.StrokeColor))
		if stroke == "" || stroke == "transparent" || stroke == "#00000000" {
			continue
		}
		r, ok := rectOf(el)
		if !ok {
			continue
		}
		topLeft := pt{X: r.X, Y: r.Y}
		topRight := pt{X: r.X + r.W, Y: r.Y}
		bottomLeft := pt{X: r.X, Y: r.Y + r.H}
		bottomRight := pt{X: r.X + r.W, Y: r.Y + r.H}
		paths = append(paths,
			[]segment{{A: topLeft, B: topRight}},
			[]segment{{A: bottomLeft, B: bottomRight}},
			[]segment{{A: topLeft, B: bottomLeft}},
			[]segment{{A: topRight, B: bottomRight}},
		)
	}
	return paths
}

func sideFromFixedPoint(fp []float64) (side, bool) {
	if len(fp) < 2 {
		return "", false
	}
	fx, fy := fp[0], fp[1]
	if math.Abs(fy-0) < 0.01 {
		return sideTop, true
	}
	if math.Abs(fy-1) < 0.01 {
		return sideBottom, true
	}
	if math.Abs(fx-0) < 0.01 {
		return sideLeft, true
	}
	if math.Abs(fx-1) < 0.01 {
		return sideRight, true
	}
	return "", false
}

func rectOf(el *Element) (rect, bool) {
	if el == nil {
		return rect{}, false
	}
	if el.Width <= 0 || el.Height <= 0 {
		return rect{}, false
	}
	return rect{X: el.X, Y: el.Y, W: el.Width, H: el.Height}, true
}

func inferSides(src, dst rect) (srcSide, dstSide side) {
	dx := dst.X + dst.W/2 - (src.X + src.W/2)
	dy := dst.Y + dst.H/2 - (src.Y + src.H/2)
	if math.Abs(dx) >= math.Abs(dy) {
		if dx >= 0 {
			return sideRight, sideLeft
		}
		return sideLeft, sideRight
	}
	if dy >= 0 {
		return sideBottom, sideTop
	}
	return sideTop, sideBottom
}

// ── Op builders (pixel → inch) ───────────────────────────────────────────────

func findPaperFrame(elements []*Element) *rect {
	for _, el := range elements {
		if el.ID == "paper-frame" || el.Type == "frame" {
			return &rect{
				X: el.X,
				Y: el.Y,
				W: math.Max(1, el.Width),
				H: math.Max(1, el.Height),
			}
		}
	}
	return nil
}

func contentBounds(elements []*Element) rect {
	visible := []*Element{}
	for _, el := range elements {
		if el.Type != "arrow" && el.Type != "line" {
			visible = append(visible, el)
		}
	}
	if len(visible) == 0 {
		return rect{X: 0, Y: 0, W: 1280, H: 720}
	}
	minX, minY := math.Inf(1), math.Inf(1)
	maxX, maxY := math.Inf(-1), math.Inf(-1)
	for _, el := range visible {
		minX = math.Min(minX, el.X)
		minY = math.Min(minY, el.Y)
		maxX = math.Max(maxX, el.X+el.Width)
		maxY = math.Max(maxY, el.Y+el.Height)
	}
	return rect{X: minX, Y: minY, W: math.Max(1, maxX-minX), H: math.Max(1, maxY-minY)}
}

type pos struct{ X, Y, W, H float64 }

func toPos(el *Element, frame rect, ppi float64) (pos, bool) {
	w := el.Width / ppi
	h := el.Height / ppi
	if w <= 0 || h <= 0 {
		return pos{}, false
	}
	return pos{
		X: (el.X - frame.X) / ppi,
		Y: (el.Y - frame.Y) / ppi,
		W: w,
		H: h,
	}, true
}

func shapeOp(el *Element, frame rect, ppi float64) (DrawOp, bool) {
	p, ok := toPos(el, frame, ppi)
	if !ok {
		return DrawOp{}, false
	}
	kind := "rect"
	if el.Type == "ellipse" {
		kind = "ellipse"
	}
	ln := lineProps(el)
	fl := fillProps(el.BackgroundColor, opacityToTransparency(el.Opacity))
	return DrawOp{
		Kind:   kind,
		X:      p.X,
		Y:      p.Y,
		W:      p.W,
		H:      p.H,
		Rotate: el.Angle,
		Line:   &ln,
		Fill:   &fl,
	}, true
}

func textOp(el *Element, frame rect, ppi float64) (DrawOp, bool) {
	p, ok := toPos(el, frame, ppi)
	if !ok {
		return DrawOp{}, false
	}
	text := el.Text
	if text == "" {
		text = el.RawText
	}
	if text == "" {
		return DrawOp{}, false
	}
	fontSize := 12.0
	if el.FontSize != nil {
		fontSize = *el.FontSize
	}
	return DrawOp{
		Kind:     "text",
		X:        p.X,
		Y:        p.Y,
		W:        p.W,
		H:        p.H,
		Rotate:   el.Angle,
		Text:     text,
		Color:    normalizeColor(el.StrokeColor, "1E1E1E"),
		FontFace: fontFace(el.FontFamily),
		FontSize: math.Max(1, pxToPt(fontSize)),
		Bold:     el.FontStyle == "bold",
		Align:    normalizeAlign(el.TextAlign),
		Valign:   normalizeValign(el.VerticalAlign),
	}, true
}

func imageOp(el *Element, files map[string]SceneFile, frame rect, ppi float64) (DrawOp, bool) {
	p, ok := toPos(el, frame, ppi)
	if !ok || el.FileID == "" {
		return DrawOp{}, false
	}
	f, ok := files[el.FileID]
	if !ok || f.DataURL == "" {
		return DrawOp{}, false
	}
	return DrawOp{
		Kind:         "image",
		X:            p.X,
		Y:            p.Y,
		W:            p.W,
		H:            p.H,
		Rotate:       el.Angle,
		Data:         f.DataURL,
		Transparency: opacityToTransparency(el.Opacity),
	}, true
}

func anchorGridOps(r rect, frame rect, ppi float64, background string) []DrawOp {
	cellW := r.W / float64(anchorGrid)
	cellH := r.H / float64(anchorGrid)
	ops := make([]DrawOp, 0, anchorGrid*anchorGrid)
	for i := 0; i < anchorGrid; i++ {
		for j := 0; j < anchorGrid; j++ {
			cx := r.X + float64(i)*cellW
			cy := r.Y + float64(j)*cellH
			ops = append(ops, DrawOp{
				Kind: "rect",
				X:    (cx - frame.X) / ppi,
				Y:    (cy - frame.Y) / ppi,
				W:    cellW / ppi,
				H:    cellH / ppi,
				Fill: &FillStyle{Color: background, Transparency: 0},
				Line: &LineStyle{Color: background, Width: 0.25, Dash: "solid", Transparency: 0},
			})
		}
	}
	return ops
}

func polylineOp(el *Element, points []pt, frame rect, ppi float64, style connectorStyle) (DrawOp, bool) {
	if len(points) < 2 {
		return DrawOp{}, false
	}
	inch := make([]pt, len(points))
	for i, p := range points {
		inch[i] = pt{X: (p.X - frame.X) / ppi, Y: (p.Y - frame.Y) / ppi}
	}
	minX, minY := inch[0].X, inch[0].Y
	maxX, maxY := inch[0].X, inch[0].Y
	for _, p := range inch {
		minX = math.Min(minX, p.X)
		minY = math.Min(minY, p.Y)
		maxX = math.Max(maxX, p.X)
		maxY = math.Max(maxY, p.Y)
	}
	w := math.Max(maxX-minX, 0.0001)
	h := math.Max(maxY-minY, 0.0001)
	rel := make([]PtIn, len(inch))
	for i, p := range inch {
		rel[i] = PtIn{X: p.X - minX, Y: p.Y - minY, MoveTo: i == 0}
	}
	ln := connectorLine(el, style)
	return DrawOp{
		Kind:   "line",
		X:      minX,
		Y:      minY,
		W:      w,
		H:      h,
		Points: rel,
		Line:   &ln,
	}, true
}

func rawLineOp(el *Element, frame rect, ppi float64, style connectorStyle) (DrawOp, bool) {
	startX := el.X - frame.X
	startY := el.Y - frame.Y
	points := el.Points
	if len(points) == 0 {
		points = [][]float64{{0, 0}, {el.Width, el.Height}}
	}
	endPoint := points[len(points)-1]
	dx, dy := 0.0, 0.0
	if len(endPoint) >= 2 {
		dx, dy = endPoint[0], endPoint[1]
	}
	x := math.Min(startX, startX+dx) / ppi
	y := math.Min(startY, startY+dy) / ppi
	w := math.Abs(dx) / ppi
	h := math.Abs(dy) / ppi
	if w <= 0 && h <= 0 {
		return DrawOp{}, false
	}
	ln := connectorLine(el, style)
	return DrawOp{
		Kind:  "line",
		X:     x,
		Y:     y,
		W:     w,
		H:     h,
		FlipH: dx < 0,
		FlipV: dy < 0,
		Line:  &ln,
	}, true
}

func connectorLine(el *Element, style connectorStyle) LineStyle {
	base := lineProps(el)
	kind := connectorKind(el)
	beginHead, endHead := connectorArrowheads(el)
	width := base.Width
	switch kind {
	case "route":
		if beginHead == "" {
			beginHead = "oval"
		}
		if endHead == "" {
			endHead = "oval"
		}
	case "traffic":
		if endHead == "" {
			endHead = style.Head
		}
	default:
		if endHead == "" && el.Type == "arrow" {
			endHead = style.Head
		}
		if style.HasWidth {
			width = style.Width
		}
	}
	if beginHead == "" {
		beginHead = "none"
	}
	if endHead == "" {
		endHead = "none"
	}
	base.Width = width
	base.BeginArrowType = beginHead
	base.EndArrowType = endHead
	return base
}

func connectorKind(el *Element) string {
	if el.CustomData == nil {
		return "connection"
	}
	switch el.CustomData.ConnectorKind {
	case "route", "traffic":
		return el.CustomData.ConnectorKind
	default:
		return "connection"
	}
}

func connectorKindPriority(kind string) int {
	switch kind {
	case "route":
		return 0
	case "traffic":
		return 2
	default:
		return 1
	}
}

func connectorArrowheads(el *Element) (string, string) {
	if el.CustomData == nil {
		return "", ""
	}
	return el.CustomData.ConnectorStartArrowhead, el.CustomData.ConnectorEndArrowhead
}

// ── Styling helpers ──────────────────────────────────────────────────────────

func lineProps(el *Element) LineStyle {
	color := normalizeColor(el.StrokeColor, "1E1E1E")
	dash := "solid"
	if el.StrokeStyle == "dashed" {
		dash = "dash"
	} else if el.StrokeStyle == "dotted" {
		dash = "dot"
	}
	transparency := opacityToTransparency(el.Opacity)
	if color == "FFFFFF" && el.StrokeColor == "transparent" {
		transparency = 100
	}
	width := el.StrokeWidth
	if width == 0 {
		width = 1
	}
	width = math.Max(0.25, width)
	return LineStyle{Color: color, Width: width, Dash: dash, Transparency: transparency}
}

func fillProps(color string, transparency float64) FillStyle {
	if color == "" || color == "transparent" {
		return FillStyle{Color: "FFFFFF", Transparency: 100}
	}
	return FillStyle{Color: normalizeColor(color, "FFFFFF"), Transparency: transparency}
}

func normalizeColor(color, fallback string) string {
	if color == "" || color == "transparent" {
		return fallback
	}
	trimmed := strings.TrimSpace(color)
	if hexFull.MatchString(trimmed) {
		trimmed = strings.TrimPrefix(trimmed, "#")
		return strings.ToUpper(trimmed)
	}
	return fallback
}

func normalizeAlign(align string) string {
	if align == "center" || align == "right" {
		return align
	}
	return "left"
}

func normalizeValign(align string) string {
	if align == "middle" || align == "bottom" {
		return align
	}
	return "top"
}

func fontFace(fontFamily *int) string {
	if fontFamily != nil && *fontFamily == 1 {
		return "Virgil"
	}
	return "Helvetica"
}

func opacityToTransparency(opacity *float64) float64 {
	value := 100.0
	if opacity != nil {
		value = *opacity
	}
	return math.Max(0, math.Min(100, 100-value))
}

func pxToPt(px float64) float64 {
	return px * 0.75
}
