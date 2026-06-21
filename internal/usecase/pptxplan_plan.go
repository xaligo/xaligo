package usecase

import (
	"encoding/json"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
)

// plan.go — builds a fully-resolved PPTX draw plan from an Excalidraw scene.
// This is the Go port of the calculation half of the former TS pptx.ts.

const (
	defaultPxPerInch        = 96.0
	anchorGrid              = 5
	anchorGridPadPx         = 4.0
	anchorGridOuterMarginPx = 2.0
	anchorGridVisualPadPx   = anchorGridPadPx + anchorGridOuterMarginPx
	// Keep the mask smaller than the default 8 px lane gap so a jump does not
	// accidentally erase a nearby parallel lane.
	lineJumpSizePx        = 6.0
	groupBorderMaskSizePx = 8.0
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

type paperMargins struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

// BuildPlanJSON parses an Excalidraw JSON scene and returns the PPTX draw plan
// as JSON, applying every geometry option in opt.
func BuildPlanJSON(sceneJSON string, opt entity.PptxOptions) ([]byte, error) {
	var scene entity.PptxScene
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
func resolvePaper(size, orientation string, contentWIn, contentHIn float64, margins paperMargins) (w, h float64, ok bool) {
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
	availPW, availPH := margins.available(pw, ph)
	availLW, availLH := margins.available(lw, lh)
	scaleP := math.Min(availPW/contentWIn, availPH/contentHIn)
	scaleL := math.Min(availLW/contentWIn, availLH/contentHIn)
	if scaleL >= scaleP {
		return lw, lh, true
	}
	return pw, ph, true
}

func resolvePaperMargins(opt entity.PptxOptions) paperMargins {
	all := math.Max(0, opt.PaperMargin)
	margins := paperMargins{Top: all, Right: all, Bottom: all, Left: all}
	if opt.PaperMarginTop > 0 {
		margins.Top = opt.PaperMarginTop
	}
	if opt.PaperMarginRight > 0 {
		margins.Right = opt.PaperMarginRight
	}
	if opt.PaperMarginBottom > 0 {
		margins.Bottom = opt.PaperMarginBottom
	}
	if opt.PaperMarginLeft > 0 {
		margins.Left = opt.PaperMarginLeft
	}
	return margins
}

func (m paperMargins) available(w, h float64) (float64, float64) {
	const minPaperContentIn = 0.01
	return math.Max(minPaperContentIn, w-m.Left-m.Right), math.Max(minPaperContentIn, h-m.Top-m.Bottom)
}

// BuildPlan converts a parsed scene into a draw plan.
func BuildPlan(scene *entity.PptxScene, opt entity.PptxOptions) entity.Plan {
	elements := make([]*entity.Element, 0, len(scene.Elements))
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
	paperMargins := resolvePaperMargins(opt)
	paperW, paperH, hasPaper := resolvePaper(opt.PaperSize, opt.Orientation, contentWIn, contentHIn, paperMargins)

	frame := *contentFrame
	ppi := basePPI
	layoutW := contentWIn
	layoutH := contentHIn
	if hasPaper {
		availableW, availableH := paperMargins.available(paperW, paperH)
		scale := math.Min(availableW/contentWIn, availableH/contentHIn)
		ppi = basePPI / scale
		offsetXIn := paperMargins.Left + (availableW-contentWIn*scale)/2
		offsetYIn := paperMargins.Top + (availableH-contentHIn*scale)/2
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

	elementsByID := map[string]*entity.Element{}
	for _, el := range elements {
		if el.ID != "" {
			elementsByID[el.ID] = el
		}
	}

	obstacles := collectObstacles(elements)

	connectors := []*entity.Element{}
	for _, el := range elements {
		if (el.Type == "arrow" || el.Type == "line") && (el.CustomData == nil || !el.CustomData.GroupHeader) {
			connectors = append(connectors, el)
		}
	}
	prepared := prepareConnectors(connectors, elementsByID)

	ops := []entity.DrawOp{}

	// 1) Anchor grids first → behind the icons drawn on top.
	gridIDs := make([]string, 0, len(prepared.gridRects))
	for id := range prepared.gridRects {
		gridIDs = append(gridIDs, id)
	}
	sort.Strings(gridIDs)
	for _, id := range gridIDs {
		grid := prepared.gridRects[id]
		ops = append(ops, anchorGridOps(id, grid.rect, frame, ppi, background)...)
	}

	// 2) Containers/shapes in scene order. Group title tags are deferred until
	// after every group border so a nested child border cannot cover a parent tag.
	headerShapes := []*entity.Element{}
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		if el.CustomData != nil && el.CustomData.Junction {
			continue
		}
		if el.CustomData != nil && el.CustomData.GroupHeader {
			headerShapes = append(headerShapes, el)
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
	groupBorders := collectGroupBorderPaths(elements)
	routed := routeConnections(reqs, obstacles, rOpt)
	elByConn := map[string]*entity.Element{}
	for _, pc := range ordered {
		elByConn[pc.req.ID] = pc.el
	}
	connectorLabels := []entity.DrawOp{}
	connectorLabelRects := []rect{}
	connectorLegend := []entity.ConnectorLegendEntry{}
	for i, path := range routed {
		el := elByConn[path.ID]
		if el == nil {
			continue
		}
		connectorID := fmt.Sprintf("L%02d", i+1)
		for maskIndex, crossing := range pathBorderCrossings(path, groupBorders) {
			ops = append(ops, groupBorderMaskOp(fmt.Sprintf("%s-border-mask-%02d", path.ID, maskIndex), crossing, frame, ppi))
		}
		for maskIndex, crossing := range pathCrossings(path, routed[:i]) {
			maskColor := lineJumpBackground(crossing, elements, background)
			ops = append(ops, lineJumpMaskOp(fmt.Sprintf("%s-jump-mask-%02d", path.ID, maskIndex), crossing, frame, ppi, maskColor))
		}
		if op, ok := polylineOp(el, path.Points, frame, ppi, style); ok {
			ops = append(ops, op)
		}
		line := connectorLine(el, style)
		if op, labelRect, ok := connectorIDLabelOp(connectorID, path, routed, obstacles, connectorLabelRects, frame, ppi, line); ok {
			connectorLabels = append(connectorLabels, op)
			connectorLabelRects = append(connectorLabelRects, labelRect)
		}
		connectorLegend = append(connectorLegend, connectorLegendEntry(connectorID, el, line))
	}
	for _, junction := range junctions {
		el := elByConn[junction.ConnectorID]
		if el == nil {
			continue
		}
		ops = append(ops, junctionOp(junction.ConnectorID, junction.Point, frame, ppi, connectorLine(el, style)))
	}
	for _, el := range headerShapes {
		if op, ok := polygonOp(el, frame, ppi); ok {
			ops = append(ops, op)
		}
	}

	anchorGroupIDs := anchorGroups(prepared.gridRects)

	// 4) Icons and labels on top so routed lines never visually cover them.
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		switch el.Type {
		case "text":
			if op, ok := textOp(el, frame, ppi); ok {
				applyAnchorGroup(&op, el.ID, anchorGroupIDs)
				ops = append(ops, op)
			}
		case "image":
			if op, ok := imageOp(el, scene.Files, frame, ppi); ok {
				applyAnchorGroup(&op, el.ID, anchorGroupIDs)
				ops = append(ops, op)
			}
		}
	}
	ops = append(ops, connectorLabels...)

	return entity.Plan{
		Slide: entity.PlanSlide{
			W:          layoutW,
			H:          layoutH,
			Background: background,
		},
		Ops:             ops,
		Legend:          buildLegend(scene, opt.LegendEntries),
		ConnectorLegend: connectorLegend,
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
				requests[endpoint.requestIndex].SrcLane = 0
			} else {
				requests[endpoint.requestIndex].DstAnchor = &copyPoint
				requests[endpoint.requestIndex].DstLane = 0
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

func junctionOp(id string, point pt, frame rect, ppi float64, line entity.LineStyle) entity.DrawOp {
	const diameterPx = 8.0
	diameter := diameterPx / ppi
	return entity.DrawOp{
		ID:         id + "-junction",
		FrontLayer: true,
		Kind:       "ellipse",
		X:          (point.X-frame.X)/ppi - diameter/2,
		Y:          (point.Y-frame.Y)/ppi - diameter/2,
		W:          diameter,
		H:          diameter,
		Fill:       &entity.FillStyle{Color: line.Color, Transparency: line.Transparency},
		Line:       &entity.LineStyle{Color: line.Color, Width: math.Max(0.75, line.Width), Dash: "solid", Transparency: line.Transparency},
	}
}

func lineJumpMaskOp(id string, crossing pt, frame rect, ppi float64, background string) entity.DrawOp {
	size := lineJumpSizePx / ppi
	return entity.DrawOp{
		ID:         id,
		FrontLayer: true,
		Kind:       "rect",
		X:          (crossing.X-frame.X)/ppi - size/2,
		Y:          (crossing.Y-frame.Y)/ppi - size/2,
		W:          size,
		H:          size,
		Fill:       &entity.FillStyle{Color: background, Transparency: 0},
		Line:       &entity.LineStyle{Color: background, Width: 0.25, Transparency: 100},
	}
}

func groupBorderMaskOp(id string, crossing pt, frame rect, ppi float64) entity.DrawOp {
	size := groupBorderMaskSizePx / ppi
	return entity.DrawOp{
		ID:         id,
		FrontLayer: true,
		Kind:       "rect",
		X:          (crossing.X-frame.X)/ppi - size/2,
		Y:          (crossing.Y-frame.Y)/ppi - size/2,
		W:          size,
		H:          size,
		Fill:       &entity.FillStyle{Color: "FFFFFF", Transparency: 0},
		Line:       &entity.LineStyle{Color: "FFFFFF", Width: 0.25, Transparency: 100},
	}
}

// lineJumpBackground returns the uppermost opaque shape fill beneath a crossing.
// Transparent or partially transparent fills fall back to the slide background,
// as reproducing their composited color would require renderer-specific blending.
func lineJumpBackground(crossing pt, elements []*entity.Element, fallback string) string {
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

func buildLegend(scene *entity.PptxScene, entries []entity.LegendEntry) []entity.LegendEntry {
	if scene == nil || len(entries) == 0 {
		return nil
	}
	out := make([]entity.LegendEntry, 0, len(entries))
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

func connectorLegendEntry(id string, el *entity.Element, line entity.LineStyle) entity.ConnectorLegendEntry {
	kind := connectorKind(el)
	entry := entity.ConnectorLegendEntry{ID: id, Kind: kind, Line: line, Source: bindingElementID(el.StartBinding), Target: bindingElementID(el.EndBinding)}
	switch kind {
	case "route":
		entry.Label = "Route line"
		entry.Description = "Network route, path, or logical reachability"
	case "traffic":
		entry.Label = "Traffic line"
		entry.Description = "Application, data, or operational communication"
	default:
		entry.Label = "Connection line"
		entry.Description = "Custom connector"
	}
	return entry
}

func bindingElementID(binding *entity.Binding) string {
	if binding == nil {
		return ""
	}
	return binding.ElementID
}

func connectorIDLabelOp(id string, path routedPath, allPaths []routedPath, obstacles []rect, placedLabels []rect, frame rect, ppi float64, line entity.LineStyle) (entity.DrawOp, rect, bool) {
	if len(path.Points) < 2 {
		return entity.DrawOp{}, rect{}, false
	}
	p, ok := connectorLabelPoint(path, allPaths, obstacles, placedLabels)
	if !ok {
		return entity.DrawOp{}, rect{}, false
	}
	w := 0.22
	h := 0.12
	labelRect := connectorIDLabelRect(p)
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
		FontSize:   5.5,
		Bold:       true,
		Align:      "center",
		Valign:     "middle",
	}, labelRect, true
}

func connectorIDLabelRect(center pt) rect {
	const labelW = 22.0
	const labelH = 12.0
	return rect{X: center.X - labelW/2, Y: center.Y - labelH/2, W: labelW, H: labelH}
}

type connectorLabelCandidate struct {
	Point    pt
	Priority float64
}

func connectorLabelPoint(path routedPath, allPaths []routedPath, obstacles []rect, placedLabels []rect) (pt, bool) {
	candidates := connectorLabelCandidates(path.Points)
	if len(candidates) == 0 {
		return pt{}, false
	}
	best := candidates[0].Point
	bestScore := connectorLabelScore(path.ID, best, allPaths, obstacles, placedLabels)
	bestScore += candidates[0].Priority
	for _, candidate := range candidates[1:] {
		score := candidate.Priority + connectorLabelScore(path.ID, candidate.Point, allPaths, obstacles, placedLabels)
		if score < bestScore {
			best = candidate.Point
			bestScore = score
		}
	}
	return best, true
}

func connectorLabelCandidates(points []pt) []connectorLabelCandidate {
	if len(points) < 2 {
		return nil
	}
	candidates := []connectorLabelCandidate{}
	add := func(point pt, priority float64) {
		candidates = append(candidates, connectorLabelCandidate{Point: point, Priority: priority})
	}
	addEndpointCandidates := func(anchor, neighbor pt) {
		dx := neighbor.X - anchor.X
		dy := neighbor.Y - anchor.Y
		length := math.Abs(dx) + math.Abs(dy)
		if length <= eps {
			return
		}
		ux, uy := dx/length, dy/length
		px, py := -uy, ux
		for _, alongOffset := range []float64{12, 18, 26, 34, 42} {
			for _, sideOffset := range []float64{-8, 8, -14, 14, -22, 22, -30, 30, -40, 40} {
				point := pt{X: anchor.X + ux*alongOffset + px*sideOffset, Y: anchor.Y + uy*alongOffset + py*sideOffset}
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
		if (math.Abs(prev.X-p.X) < eps && math.Abs(next.X-p.X) < eps) || (math.Abs(prev.Y-p.Y) < eps && math.Abs(next.Y-p.Y) < eps) {
			continue
		}
		for _, dx := range []float64{-8, 8, -14, 14, -22, 22, -30, 30, -40, 40} {
			for _, dy := range []float64{-8, 8, -14, 14, -22, 22, -30, 30, -40, 40} {
				point := pt{X: p.X + dx, Y: p.Y + dy}
				add(point, 2.0+math.Hypot(dx, dy)*0.18)
			}
		}
	}
	return candidates
}

func connectorLabelScore(pathID string, center pt, allPaths []routedPath, obstacles []rect, placedLabels []rect) float64 {
	label := connectorIDLabelRect(center)
	score := 0.0
	for _, obstacle := range obstacles {
		if rectsOverlap(label, inflate(obstacle, 2)) {
			score += 1000
		}
		score += proximityPenalty(distanceRectToRect(label, obstacle), 18, 18)
	}
	for _, path := range allPaths {
		pathPenalty := 160.0
		if path.ID == pathID {
			pathPenalty = 6.0
		}
		for _, seg := range toSegments(path.Points) {
			if segIntersectsRect(seg, label) {
				score += pathPenalty
			}
			if path.ID != pathID {
				score += proximityPenalty(distancePointToSegment(center, seg), 14, pathPenalty/10)
			}
		}
	}
	for _, placed := range placedLabels {
		if rectsOverlap(label, inflate(placed, 2)) {
			score += 1200
		}
		score += proximityPenalty(distanceRectToRect(label, placed), 24, 40)
	}
	return score
}

func proximityPenalty(distance, threshold, weight float64) float64 {
	if distance >= threshold {
		return 0
	}
	return weight * (threshold - distance) / threshold
}

func rectsOverlap(a, b rect) bool {
	return a.X < b.X+b.W && a.X+a.W > b.X && a.Y < b.Y+b.H && a.Y+a.H > b.Y
}

func distanceRectToRect(a, b rect) float64 {
	dx := math.Max(math.Max(b.X-(a.X+a.W), a.X-(b.X+b.W)), 0)
	dy := math.Max(math.Max(b.Y-(a.Y+a.H), a.Y-(b.Y+b.H)), 0)
	return math.Hypot(dx, dy)
}

func distancePointToSegment(p pt, seg segment) float64 {
	x1, y1 := seg.A.X, seg.A.Y
	x2, y2 := seg.B.X, seg.B.Y
	dx, dy := x2-x1, y2-y1
	lengthSq := dx*dx + dy*dy
	if lengthSq <= eps {
		return math.Hypot(p.X-x1, p.Y-y1)
	}
	t := ((p.X-x1)*dx + (p.Y-y1)*dy) / lengthSq
	t = math.Max(0, math.Min(1, t))
	closest := pt{X: x1 + t*dx, Y: y1 + t*dy}
	return math.Hypot(p.X-closest.X, p.Y-closest.Y)
}

func backgroundColor(scene *entity.PptxScene) string {
	if scene.AppState != nil {
		return scene.AppState.ViewBackgroundColor
	}
	return ""
}

// ── Connector preparation (anchors + sides) ──────────────────────────────────

type preparedConnector struct {
	el  *entity.Element
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
	raw       []*entity.Element
	gridRects map[string]anchorGridRect
}

type anchorGridRect struct {
	rect       rect
	background string
}

func prepareConnectors(connectors []*entity.Element, byID map[string]*entity.Element) preparedResult {
	raw := []*entity.Element{}
	gridRects := map[string]anchorGridRect{}
	groupKeys := []string{}
	groups := map[string][]endpoint{}
	type item struct {
		el      *entity.Element
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
		srcEl := byID[srcIconID]
		dstEl := byID[dstIconID]
		src, srcOK := rectOf(srcEl)
		dst, dstOK := rectOf(dstEl)
		if !srcOK || !dstOK || el.ID == "" {
			raw = append(raw, el)
			continue
		}
		srcGrid := anchorGridForElement(srcIconID, srcEl, src, byID)
		dstGrid := anchorGridForElement(dstIconID, dstEl, dst, byID)
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
		if srcSide == sideBottom && strings.HasSuffix(srcIconID, "-lbl") {
			src = inflateRect(src, anchorGridVisualPadPx)
		}
		if dstSide == sideBottom && strings.HasSuffix(dstIconID, "-lbl") {
			dst = inflateRect(dst, anchorGridVisualPadPx)
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
		gridRects[srcIconID] = srcGrid
		gridRects[dstIconID] = dstGrid
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
			req.SrcLane = a.srcLane
			req.DstLane = a.dstLane
		}
		routed = append(routed, preparedConnector{el: it.el, req: req})
	}
	return preparedResult{routed: routed, raw: raw, gridRects: gridRects}
}

func anchorGridForElement(id string, el *entity.Element, base rect, byID map[string]*entity.Element) anchorGridRect {
	grid := anchorGridRect{rect: base, background: el.BackgroundColor}
	if strings.HasSuffix(id, "-lbl") {
		if imageEl := byID[strings.TrimSuffix(id, "-lbl")]; imageEl != nil {
			if imageRect, ok := rectOf(imageEl); ok {
				grid.rect = unionRect(grid.rect, imageRect)
				if grid.background == "" || grid.background == "transparent" {
					grid.background = imageEl.BackgroundColor
				}
			}
		}
		grid.rect = inflateRect(grid.rect, anchorGridVisualPadPx)
		return grid
	}
	if labelEl := byID[id+"-lbl"]; labelEl != nil {
		if labelRect, ok := rectOf(labelEl); ok {
			grid.rect = unionRect(grid.rect, labelRect)
		}
	}
	grid.rect = inflateRect(grid.rect, anchorGridVisualPadPx)
	return grid
}

func inflateRect(r rect, pad float64) rect {
	return rect{X: r.X - pad, Y: r.Y - pad, W: r.W + pad*2, H: r.H + pad*2}
}

func unionRect(a, b rect) rect {
	minX := math.Min(a.X, b.X)
	minY := math.Min(a.Y, b.Y)
	maxX := math.Max(a.X+a.W, b.X+b.W)
	maxY := math.Max(a.Y+a.H, b.Y+b.H)
	return rect{X: minX, Y: minY, W: maxX - minX, H: maxY - minY}
}

type anchorPair struct {
	src     *pt
	dst     *pt
	srcLane float64
	dstLane float64
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
		lane := 0.0
		if len(eps) > 1 {
			lane = float64(k) - float64(len(eps)-1)/2
		}
		if ep.isSrc {
			entry.src = &pc
			entry.srcLane = lane
		} else {
			entry.dst = &pc
			entry.dstLane = lane
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

func collectObstacles(elements []*entity.Element) []rect {
	rects := []rect{}
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		isHeader := el.CustomData != nil && el.CustomData.GroupHeader
		if el.Type != "image" && el.Type != "text" && !isHeader {
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
func collectContainerBorderPaths(elements []*entity.Element) [][]segment {
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

func collectGroupBorderPaths(elements []*entity.Element) []segment {
	var paths []segment
	for _, el := range elements {
		if el.CustomData == nil || !el.CustomData.GroupBorder {
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
		tl := pt{X: r.X, Y: r.Y}
		tr := pt{X: r.X + r.W, Y: r.Y}
		bl := pt{X: r.X, Y: r.Y + r.H}
		br := pt{X: r.X + r.W, Y: r.Y + r.H}
		paths = append(paths, segment{A: tl, B: tr}, segment{A: tr, B: br}, segment{A: br, B: bl}, segment{A: bl, B: tl})
	}
	return paths
}

func pathBorderCrossings(path routedPath, borders []segment) []pt {
	var out []pt
	for _, pathSeg := range toSegments(path.Points) {
		for _, border := range borders {
			p, ok := crossingPoint(pathSeg, border)
			if !ok {
				continue
			}
			duplicate := false
			for _, existing := range out {
				if math.Abs(existing.X-p.X) < eps && math.Abs(existing.Y-p.Y) < eps {
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

func rectOf(el *entity.Element) (rect, bool) {
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

func findPaperFrame(elements []*entity.Element) *rect {
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

func contentBounds(elements []*entity.Element) rect {
	visible := []*entity.Element{}
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

func toPos(el *entity.Element, frame rect, ppi float64) (pos, bool) {
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

func shapeOp(el *entity.Element, frame rect, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPos(el, frame, ppi)
	if !ok {
		return entity.DrawOp{}, false
	}
	kind := "rect"
	if el.Type == "ellipse" {
		kind = "ellipse"
	}
	ln := lineProps(el)
	fl := fillProps(el.BackgroundColor, opacityToTransparency(el.Opacity))
	return entity.DrawOp{
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

func polygonOp(el *entity.Element, frame rect, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPos(el, frame, ppi)
	if !ok || len(el.Points) < 3 {
		return entity.DrawOp{}, false
	}
	points := make([]entity.PtIn, 0, len(el.Points))
	for i, point := range el.Points {
		if len(point) < 2 {
			continue
		}
		points = append(points, entity.PtIn{X: point[0] / ppi, Y: point[1] / ppi, MoveTo: i == 0})
	}
	if len(points) < 3 {
		return entity.DrawOp{}, false
	}
	ln := lineProps(el)
	fl := fillProps(el.BackgroundColor, opacityToTransparency(el.Opacity))
	return entity.DrawOp{Kind: "polygon", X: p.X, Y: p.Y, W: p.W, H: p.H, Rotate: el.Angle, Points: points, Line: &ln, Fill: &fl}, true
}

func textOp(el *entity.Element, frame rect, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPos(el, frame, ppi)
	if !ok {
		return entity.DrawOp{}, false
	}
	text := el.Text
	if text == "" {
		text = el.RawText
	}
	if text == "" {
		return entity.DrawOp{}, false
	}
	fontSize := 12.0
	if el.FontSize != nil {
		fontSize = *el.FontSize
	}
	return entity.DrawOp{
		ID:       el.ID,
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

func imageOp(el *entity.Element, files map[string]entity.SceneFile, frame rect, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPos(el, frame, ppi)
	if !ok || el.FileID == "" {
		return entity.DrawOp{}, false
	}
	f, ok := files[el.FileID]
	if !ok || f.DataURL == "" {
		return entity.DrawOp{}, false
	}
	return entity.DrawOp{
		ID:           el.ID,
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

func anchorGridOps(id string, r rect, frame rect, ppi float64, background string) []entity.DrawOp {
	cellW := r.W / float64(anchorGrid)
	cellH := r.H / float64(anchorGrid)
	ops := make([]entity.DrawOp, 0, anchorGrid*anchorGrid)
	baseID := anchorBaseID(id)
	groupID := anchorGroupID(baseID)
	for i := 0; i < anchorGrid; i++ {
		for j := 0; j < anchorGrid; j++ {
			cx := r.X + float64(i)*cellW
			cy := r.Y + float64(j)*cellH
			ops = append(ops, entity.DrawOp{
				ID:      fmt.Sprintf("%s-grid-%02d-%02d", baseID, i, j),
				GroupID: groupID,
				Kind:    "rect",
				X:       (cx - frame.X) / ppi,
				Y:       (cy - frame.Y) / ppi,
				W:       cellW / ppi,
				H:       cellH / ppi,
				Fill:    &entity.FillStyle{Color: background, Transparency: 0},
				Line:    &entity.LineStyle{Color: background, Width: 0.25, Dash: "solid", Transparency: 0},
			})
		}
	}
	return ops
}

func anchorGroups(grids map[string]anchorGridRect) map[string]string {
	out := map[string]string{}
	for id := range grids {
		baseID := anchorBaseID(id)
		out[baseID] = anchorGroupID(baseID)
	}
	return out
}

func applyAnchorGroup(op *entity.DrawOp, elementID string, groups map[string]string) {
	if op == nil || elementID == "" {
		return
	}
	baseID := anchorBaseID(elementID)
	groupID, ok := groups[baseID]
	if !ok {
		return
	}
	op.GroupID = groupID
}

func anchorBaseID(id string) string {
	return strings.TrimSuffix(id, "-lbl")
}

func anchorGroupID(baseID string) string {
	return "xaligo-anchor-" + baseID
}

func polylineOp(el *entity.Element, points []pt, frame rect, ppi float64, style connectorStyle) (entity.DrawOp, bool) {
	if len(points) < 2 {
		return entity.DrawOp{}, false
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
	rel := make([]entity.PtIn, len(inch))
	for i, p := range inch {
		rel[i] = entity.PtIn{X: p.X - minX, Y: p.Y - minY, MoveTo: i == 0}
	}
	ln := connectorLine(el, style)
	return entity.DrawOp{
		ID:         el.ID,
		FrontLayer: true,
		Kind:       "line",
		X:          minX,
		Y:          minY,
		W:          w,
		H:          h,
		Points:     rel,
		Line:       &ln,
	}, true
}

func rawLineOp(el *entity.Element, frame rect, ppi float64, style connectorStyle) (entity.DrawOp, bool) {
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
		return entity.DrawOp{}, false
	}
	ln := connectorLine(el, style)
	return entity.DrawOp{
		ID:         el.ID,
		FrontLayer: true,
		Kind:       "line",
		X:          x,
		Y:          y,
		W:          w,
		H:          h,
		FlipH:      dx < 0,
		FlipV:      dy < 0,
		Line:       &ln,
	}, true
}

func connectorLine(el *entity.Element, style connectorStyle) entity.LineStyle {
	base := lineProps(el)
	kind := connectorKind(el)
	beginHead, endHead := connectorArrowheads(el)
	width := base.Width
	switch kind {
	case "route":
		if endHead == "" {
			endHead = style.Head
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
	if style.HasWidth {
		width = style.Width
	}
	base.Width = width
	base.BeginArrowType = beginHead
	base.EndArrowType = endHead
	return base
}

func connectorKind(el *entity.Element) string {
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

func connectorArrowheads(el *entity.Element) (string, string) {
	if el.CustomData == nil {
		return "", ""
	}
	return el.CustomData.ConnectorStartArrowhead, el.CustomData.ConnectorEndArrowhead
}

// ── Styling helpers ──────────────────────────────────────────────────────────

func lineProps(el *entity.Element) entity.LineStyle {
	color := normalizeColor(el.StrokeColor, "1E1E1E")
	dash := "solid"
	if el.StrokeStyle == "dashed" {
		dash = "dash"
	} else if el.StrokeStyle == "dotted" {
		dash = "dot"
	}
	transparency := opacityToTransparency(el.Opacity)
	if strings.EqualFold(strings.TrimSpace(el.StrokeColor), "transparent") || strings.EqualFold(strings.TrimSpace(el.StrokeColor), "#00000000") {
		transparency = 100
	}
	width := el.StrokeWidth
	if width == 0 {
		width = 1
	}
	width = math.Max(0.25, width)
	return entity.LineStyle{Color: color, Width: width, Dash: dash, Transparency: transparency}
}

func fillProps(color string, transparency float64) entity.FillStyle {
	if color == "" || color == "transparent" {
		return entity.FillStyle{Color: "FFFFFF", Transparency: 100}
	}
	return entity.FillStyle{Color: normalizeColor(color, "FFFFFF"), Transparency: transparency}
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
