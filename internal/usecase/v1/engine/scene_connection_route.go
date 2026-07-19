package engine

import (
	"math"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func firstNonEmptyAttrV1EngineSceneConnectionRoute(node *entity.Node, names ...string) string {
	if node == nil {
		return ""
	}
	for _, name := range names {
		if value := strings.TrimSpace(node.Attrs[name]); value != "" {
			return value
		}
	}
	return ""
}

func connectionEndpointSideV1EngineSceneConnectionRoute(conn *entity.Node, endpoint string) (sideV1EngineRouteTypes, bool) {
	if conn == nil {
		return "", false
	}
	return normalizeConnectionSideV1EngineParseConnection(conn.Attrs[endpoint+"-side"])
}

func connectionEndpointAnchorV1EngineSceneConnectionRoute(conn *entity.Node, endpoint string) (connectionAnchorSpecV1EngineParseConnection, bool) {
	if conn == nil {
		return connectionAnchorSpecV1EngineParseConnection{}, false
	}
	spec, ok, err := parseConnectionAnchorSpecV1EngineParseConnection(conn.Attrs[endpoint+"-side"], conn.Attrs[endpoint+"-anchor"])
	if err != nil || !ok || !spec.hasSlot {
		return connectionAnchorSpecV1EngineParseConnection{}, false
	}
	return spec, true
}

func connectionFrameAnchorV1EngineSceneConnectionRoute(conn *entity.Node, endpoint string) (connectionAnchorSpecV1EngineParseConnection, bool) {
	if conn == nil {
		return connectionAnchorSpecV1EngineParseConnection{}, false
	}
	spec, ok, err := parseConnectionAnchorSpecV1EngineParseConnection(conn.Attrs[endpoint+"-frame-side"], conn.Attrs[endpoint+"-frame-anchor"])
	if err != nil || !ok {
		return connectionAnchorSpecV1EngineParseConnection{}, false
	}
	return spec, true
}

func excalidrawConnectionPointsV1EngineSceneConnectionRoute(conn *entity.Node, srcRect, dstRect [4]float64, srcSide, dstSide, kind string, obstacles, hardObstacles []rectV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, routePaths map[string][]ptV1EngineRouteTypes) []ptV1EngineRouteTypes {
	if points, ok := umlSequenceSelfMessagePointsV1EngineSceneConnectionRoute(conn, srcRect); ok {
		return points
	}
	req := excalidrawRouteRequestV1EngineSceneConnectionRoute(conn, srcRect, dstRect, srcSide, dstSide, kind)
	opt := defaultRouterOptionsV1EngineRouteTypes()
	opt.HardObstacles = hardObstacles
	local := filterObstaclesV1EngineRouteOverlap(obstacles, req)
	path := routeOneV1EngineRoutePath(req, local, placed, opt)
	followedRoute := false
	if req.Kind == "traffic" {
		if base, ok := matchingRoutePathV1EngineRouteBuild(req, routePaths); ok {
			path.Points = trafficAlongsideRouteV1EngineRouteBuild(base, path.Points, opt.LaneGap)
			path.Points = separateExactOverlapsV1EngineRouteOverlap(path.Points, placed, local, opt)
			followedRoute = true
		} else {
			path.Points = separateExactOverlapsV1EngineRouteOverlap(path.Points, placed, local, opt)
		}
	} else if req.Kind != "route" {
		path.Points = separateExactOverlapsV1EngineRouteOverlap(path.Points, placed, local, opt)
	}
	visualMargin := math.Min(opt.LineMargin, opt.Clearance) / 2
	path.Points = separateObstacleHitsV1EngineRouteOverlap(path.Points, placed, inflateRectsV1EngineRouteOverlap(local, visualMargin), opt)
	if len(req.Bends) == 0 {
		path.Points = rerouteEndpointApproachV1EngineRoutePath(path.Points, req, opt)
	}
	path.Points = separatePinnedExactOverlapsV1EngineSceneConnectionRoute(path.Points, placed, local, opt)
	if followedRoute {
		path.Points = restoreDestinationApproachV1EngineRouteBuild(path.Points, req.DstSide, opt.Stub)
	}
	path.Points = enforceOrthogonalPolylineV1EngineRoutePath(path.Points)
	return enforceHardObstacleExclusionV1EngineRouteBuild(req, path.Points, local, placed, opt)
}

func umlSequenceSelfMessagePointsV1EngineSceneConnectionRoute(conn *entity.Node, rect [4]float64) ([]ptV1EngineRouteTypes, bool) {
	if conn == nil || conn.Attr("uml-diagram-kind") != "sequence-diagram" || conn.Attr("src") != conn.Attr("dst") {
		return nil, false
	}
	srcFP, ok := umlSequenceFixedPointV1EngineSceneConnectionRoute(conn, "src", "right")
	if !ok {
		return nil, false
	}
	dstFP, ok := umlSequenceFixedPointV1EngineSceneConnectionRoute(conn, "dst", "right")
	if !ok {
		return nil, false
	}
	start := ptV1EngineRouteTypes{X: rect[0] + rect[2]*srcFP[0], Y: rect[1] + rect[3]*srcFP[1]}
	end := ptV1EngineRouteTypes{X: rect[0] + rect[2]*dstFP[0], Y: rect[1] + rect[3]*dstFP[1]}
	visualX := rect[0] + rect[2]/2 + 8
	start.X = visualX
	end.X = visualX
	loopWidth := math.Max(math.Max(72, rect[2]*3), umlSequenceSelfMessageLabelWidthV1EngineSceneConnectionRoute(conn)+24)
	return []ptV1EngineRouteTypes{
		start,
		{X: start.X + loopWidth, Y: start.Y},
		{X: start.X + loopWidth, Y: end.Y},
		end,
	}, true
}

func umlSequenceSelfMessageLabelWidthV1EngineSceneConnectionRoute(conn *entity.Node) float64 {
	label := strings.TrimSpace(conn.Attr("uml-relation-label"))
	if label == "" {
		return 0
	}
	return math.Max(80, math.Min(220, textWidthV1EngineSceneItem(label, 6)+16))
}

func separatePinnedExactOverlapsV1EngineSceneConnectionRoute(points []ptV1EngineRouteTypes, placed [][]segmentV1EngineRouteTypes, obstacles []rectV1EngineRouteTypes, opt routerOptionsV1EngineRouteTypes) []ptV1EngineRouteTypes {
	if len(points) < 3 || len(placed) == 0 || opt.LaneGap <= 0 {
		return points
	}
	best := append([]ptV1EngineRouteTypes(nil), points...)
	bestOverlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(best), placed)
	if bestOverlap <= epsV1EngineRouteTypes {
		return best
	}
	inflated := inflateRectsV1EngineRouteOverlap(obstacles, math.Min(opt.LineMargin, opt.Clearance)/2)
	bestScore := scorePathV1EngineRouteCandidate(best, inflated, placed, opt.LineMargin)
	for _, offset := range []float64{opt.LaneGap, -opt.LaneGap, opt.LaneGap * 2, -opt.LaneGap * 2} {
		shifted := offsetPolylineV1EngineRouteBuild(points, offset)
		candidate := []ptV1EngineRouteTypes{points[0]}
		appendTarget := shifted[1]
		candidate = appendOrthogonalLegV1EngineRoutePath(candidate, points[0], appendTarget)
		if len(shifted) > 3 {
			candidate = append(candidate, shifted[2:len(shifted)-1]...)
		}
		candidate = appendOrthogonalLegV1EngineRoutePath(candidate, shifted[len(shifted)-2], points[len(points)-1])
		candidate = simplifyRouteCandidateV1EngineRoutePath(candidate)
		candidate = enforceOrthogonalPolylineV1EngineRoutePath(candidate)
		if obstacleHitCountV1EngineRouteCandidate(candidate, inflated) > 0 {
			continue
		}
		overlap := exactOverlapLengthV1EngineRouteOverlap(toSegmentsV1EngineRouteGeometry(candidate), placed)
		score := scorePathV1EngineRouteCandidate(candidate, inflated, placed, opt.LineMargin)
		if overlap < bestOverlap-epsV1EngineRouteTypes || (math.Abs(overlap-bestOverlap) < epsV1EngineRouteTypes && score < bestScore) {
			best, bestOverlap, bestScore = candidate, overlap, score
		}
	}
	return best
}

func excalidrawRouteRequestV1EngineSceneConnectionRoute(conn *entity.Node, srcRect, dstRect [4]float64, srcSide, dstSide, kind string) routeRequestV1EngineRouteTypes {
	src := rectV1EngineRouteTypes{X: srcRect[0], Y: srcRect[1], W: srcRect[2], H: srcRect[3]}
	dst := rectV1EngineRouteTypes{X: dstRect[0], Y: dstRect[1], W: dstRect[2], H: dstRect[3]}
	req := routeRequestV1EngineRouteTypes{
		ID:      firstNonEmptyAttrV1EngineSceneConnectionRoute(conn, "src") + "-" + firstNonEmptyAttrV1EngineSceneConnectionRoute(conn, "dst"),
		Kind:    kind,
		Src:     src,
		Dst:     dst,
		SrcSide: sideV1EngineRouteTypes(srcSide),
		DstSide: sideV1EngineRouteTypes(dstSide),
		SrcGap:  5,
		DstGap:  5,
	}
	if anchor, ok := connectionEndpointAnchorV1EngineSceneConnectionRoute(conn, "src"); ok {
		fp := fixedPointForAnchorV1EngineSceneConnection(anchor)
		req.SrcAnchor = &ptV1EngineRouteTypes{X: src.X + src.W*fp[0], Y: src.Y + src.H*fp[1]}
	}
	if anchor, ok := connectionEndpointAnchorV1EngineSceneConnectionRoute(conn, "dst"); ok {
		fp := fixedPointForAnchorV1EngineSceneConnection(anchor)
		req.DstAnchor = &ptV1EngineRouteTypes{X: dst.X + dst.W*fp[0], Y: dst.Y + dst.H*fp[1]}
	}
	if fp, ok := umlSequenceFixedPointV1EngineSceneConnectionRoute(conn, "src", srcSide); ok {
		req.SrcAnchor = &ptV1EngineRouteTypes{X: src.X + src.W*fp[0], Y: src.Y + src.H*fp[1]}
	}
	if fp, ok := umlSequenceFixedPointV1EngineSceneConnectionRoute(conn, "dst", dstSide); ok {
		req.DstAnchor = &ptV1EngineRouteTypes{X: dst.X + dst.W*fp[0], Y: dst.Y + dst.H*fp[1]}
	}
	if scale, ok := positiveFloatAttrV1EngineSceneConnectionRoute(conn, "coordinate-scale", "scale"); ok {
		req.Bends = parseConnectorBendsV1EnginePlanConnectorPrepare(connectionBendsV1EngineSceneConnectionRoute(conn), scale)
	} else {
		req.Bends = parseConnectorBendsV1EnginePlanConnectorPrepare(connectionBendsV1EngineSceneConnectionRoute(conn), 1)
	}
	if grid, ok := positiveFloatAttrV1EngineSceneConnectionRoute(conn, "grid"); ok {
		req.Grid = grid
	}
	return req
}

func umlSequenceFixedPointV1EngineSceneConnectionRoute(conn *entity.Node, endpoint, side string) ([2]float64, bool) {
	position, ok := umlSequencePositionV1EngineSceneConnectionRoute(conn, endpoint)
	if !ok {
		return [2]float64{}, false
	}
	switch umlSequenceVerticalSideV1EngineSceneConnectionRoute(side) {
	case "left":
		return [2]float64{0, position}, true
	default:
		return [2]float64{1, position}, true
	}
}

func umlSequencePositionV1EngineSceneConnectionRoute(conn *entity.Node, endpoint string) (float64, bool) {
	if conn == nil || conn.Attr("uml-diagram-kind") != "sequence-diagram" {
		return 0, false
	}
	position, err := strconv.ParseFloat(strings.TrimSpace(conn.Attr("uml-sequence-position")), 64)
	if err != nil || math.IsNaN(position) || math.IsInf(position, 0) || position <= 0 || position >= 1 {
		return 0, false
	}
	if endpoint == "dst" && conn.Attr("src") == conn.Attr("dst") {
		position = math.Min(0.98, position+0.04)
	}
	return position, true
}

func umlSequenceVerticalSideV1EngineSceneConnectionRoute(side string) string {
	switch side {
	case "top":
		return "left"
	case "bottom":
		return "right"
	default:
		return side
	}
}

func excalidrawRouteObstaclesV1EngineSceneConnectionRoute(elements []map[string]any) []rectV1EngineRouteTypes {
	obstacles := make([]rectV1EngineRouteTypes, 0)
	for _, el := range elements {
		custom, _ := el["customData"].(map[string]any)
		isAnchorContent, _ := custom["xaligoAnchorContent"].(bool)
		isHeader, _ := custom["xaligoGroupHeader"].(bool)
		isHeaderContent, _ := custom["xaligoGroupHeaderContent"].(bool)
		isFrameMetadata, _ := custom["xaligoFrameMetadata"].(bool)
		isFrameMetadataReserved, _ := custom["xaligoFrameMetadataReserved"].(bool)
		if !isAnchorContent && !isHeader && !isHeaderContent && !isFrameMetadata && !isFrameMetadataReserved {
			continue
		}
		r, ok := elementRectV1EngineSceneConnectionRoute(el)
		if !ok {
			continue
		}
		obstacles = append(obstacles, r)
	}
	return obstacles
}

func elementRectV1EngineSceneConnectionRoute(el map[string]any) (rectV1EngineRouteTypes, bool) {
	x, okX := el["x"].(float64)
	y, okY := el["y"].(float64)
	w, okW := el["width"].(float64)
	h, okH := el["height"].(float64)
	if !okX || !okY || !okW || !okH || w <= 0 || h <= 0 {
		return rectV1EngineRouteTypes{}, false
	}
	return rectV1EngineRouteTypes{X: x, Y: y, W: w, H: h}, true
}

func connectionBendsV1EngineSceneConnectionRoute(conn *entity.Node) string {
	if conn == nil {
		return ""
	}
	points := connectionChildBendsV1EngineSceneConnectionRoute(conn)
	if len(points) > 0 {
		return strings.Join(points, " ")
	}
	return firstNonEmptyAttrV1EngineSceneConnectionRoute(conn, "bends", "points", "via")
}

func connectionChildBendsV1EngineSceneConnectionRoute(node *entity.Node) []string {
	if node == nil {
		return nil
	}
	points := []string{}
	for _, child := range node.Children {
		switch strings.ToLower(strings.TrimSpace(child.Tag)) {
		case "bend", "point", "via", "waypoint":
			if point, ok := connectionPointStringV1EngineSceneConnectionRoute(child); ok {
				points = append(points, point)
			}
		case "bends", "points", "path":
			points = append(points, connectionChildBendsV1EngineSceneConnectionRoute(child)...)
		}
	}
	return points
}

func connectionPointStringV1EngineSceneConnectionRoute(node *entity.Node) (string, bool) {
	x, xOK := floatAttrV1EngineSceneConnectionRoute(node, "x")
	y, yOK := floatAttrV1EngineSceneConnectionRoute(node, "y")
	if xOK && yOK {
		return fmtFloatV1EngineRouteBuild(x) + "," + fmtFloatV1EngineRouteBuild(y), true
	}
	parts := strings.Split(strings.TrimSpace(node.Text), ",")
	if len(parts) != 2 {
		return "", false
	}
	x, xErr := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	y, yErr := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if xErr != nil || yErr != nil {
		return "", false
	}
	return fmtFloatV1EngineRouteBuild(x) + "," + fmtFloatV1EngineRouteBuild(y), true
}

func floatAttrV1EngineSceneConnectionRoute(node *entity.Node, name string) (float64, bool) {
	if node == nil {
		return 0, false
	}
	value := strings.TrimSpace(node.Attrs[name])
	if value == "" {
		return 0, false
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return parsed, true
}

func positiveFloatAttrV1EngineSceneConnectionRoute(node *entity.Node, names ...string) (float64, bool) {
	if node == nil {
		return 0, false
	}
	for _, name := range names {
		value := strings.TrimSpace(node.Attrs[name])
		if value == "" {
			continue
		}
		parsed, err := strconv.ParseFloat(value, 64)
		if err == nil && parsed > 0 {
			return parsed, true
		}
	}
	return 0, false
}

func stableConnectionSeedV1EngineSceneConnectionRoute(srcKey, dstKey string, index int) int {
	// Use the FNV-1a uint32 domain explicitly so connector metadata is
	// identical on every Go target, regardless of the native int width.
	seed := uint32(2166136261)
	for _, r := range srcKey + "|" + dstKey + "|" + strconv.Itoa(index) {
		seed ^= uint32(r)
		seed *= 16777619
	}
	return int(seed%99999999) + 1
}

func sanitizeElementIDV1EngineSceneConnectionRoute(s string) string {
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			b.WriteRune(r)
			continue
		}
		b.WriteByte('-')
	}
	out := b.String()
	if out == "" {
		return "endpoint"
	}
	return out
}

func extendConnectionPointV1EngineSceneConnectionRoute(point [2]float64, side string, distance float64) [2]float64 {
	switch side {
	case "top":
		point[1] -= distance
	case "bottom":
		point[1] += distance
	case "left":
		point[0] -= distance
	default:
		point[0] += distance
	}
	return point
}

func connectionKindV1EngineSceneConnectionRoute(conn *entity.Node) string {
	kind := strings.ToLower(strings.TrimSpace(conn.Attrs["kind"]))
	if kind == "route" || kind == "traffic" {
		return kind
	}
	return "connection"
}

func connectionKindPriorityV1EngineSceneConnectionRoute(kind string) int {
	switch kind {
	case "route":
		return 0
	case "traffic":
		return 2
	default:
		return 1
	}
}

func resolveConnectionStyleV1EngineSceneConnectionRoute(conn *entity.Node) resolvedConnectionStyleV1EngineSceneTypes {
	kind := connectionKindV1EngineSceneConnectionRoute(conn)
	style := resolvedConnectionStyleV1EngineSceneTypes{
		Kind: kind, Color: "#1e1e1e", Width: 1, StrokeStyle: "solid",
		StartArrowhead: "none", EndArrowhead: "stealth",
		ExcalidrawStartArrowhead: nil, ExcalidrawEndArrowhead: "arrow",
	}
	switch kind {
	case "route":
		style.Color = "#64748b"
		style.EndArrowhead = "none"
		style.ExcalidrawEndArrowhead = nil
	case "traffic":
		style.Color = "#2563eb"
	}
	if color := strings.TrimSpace(conn.Attrs["color"]); color != "" {
		style.Color = color
	}
	widthValue := strings.TrimSpace(conn.Attrs["stroke-width"])
	if widthValue == "" {
		widthValue = strings.TrimSpace(conn.Attrs["width"])
	}
	if width, err := strconv.ParseFloat(widthValue, 64); err == nil && width > 0 {
		style.Width = width
		style.WidthExplicit = true
	}
	if strokeStyle := strings.ToLower(strings.TrimSpace(conn.Attrs["stroke-style"])); strokeStyle == "solid" || strokeStyle == "dashed" || strokeStyle == "dotted" {
		style.StrokeStyle = strokeStyle
	}
	startArrowhead := strings.TrimSpace(conn.Attrs["start-arrowhead"])
	endArrowhead := strings.TrimSpace(conn.Attrs["end-arrowhead"])
	if endArrowhead == "" {
		endArrowhead = strings.TrimSpace(conn.Attrs["arrowhead"])
	}
	style.StartArrowheadExplicit = startArrowhead != ""
	style.EndArrowheadExplicit = endArrowhead != ""
	style.StartArrowhead, style.ExcalidrawStartArrowhead = resolveArrowheadV1EngineSceneConnectionRoute(startArrowhead, style.StartArrowhead, style.ExcalidrawStartArrowhead)
	style.EndArrowhead, style.ExcalidrawEndArrowhead = resolveArrowheadV1EngineSceneConnectionRoute(endArrowhead, style.EndArrowhead, style.ExcalidrawEndArrowhead)
	return style
}

func resolveArrowheadV1EngineSceneConnectionRoute(value, current string, currentExcalidraw any) (string, any) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "none":
		return "none", nil
	case "arrow", "triangle", "diamond":
		value = strings.ToLower(strings.TrimSpace(value))
		return value, value
	case "stealth":
		return "stealth", "arrow"
	case "oval":
		return "oval", "dot"
	default:
		return current, currentExcalidraw
	}
}
