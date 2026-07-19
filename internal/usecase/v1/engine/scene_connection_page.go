package engine

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type crossFrameConnectorMetadataV1EngineSceneConnectionPage struct {
	logicalID            string
	sourceElementID      string
	destinationElementID string
}

const (
	pageLinkLabelGapV1EngineSceneConnectionPage = 4.0
	pageLinkApproachV1EngineSceneConnectionPage = 24.0
)

func renderCrossFrameConnectionV1EngineSceneConnectionPage(conn *entity.Node, srcKey, dstKey, srcElemID, dstElemID string, srcRect, dstRect, srcVisualRect, dstVisualRect [4]float64, srcFP, dstFP [2]float64, srcSide, dstSide, srcFrameSide, dstFrameSide, srcFrameID, dstFrameID string, srcFrameRect, dstFrameRect [4]float64, srcFrameEndpoint, dstFrameEndpoint bool, srcMetadata, dstMetadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata, style resolvedConnectionStyleV1EngineSceneTypes, seed, index int, updated int64, elements *[]map[string]any, boundMap map[string][]map[string]any) {
	srcEdge := rectFixedPointV1EngineSceneConnection(srcRect, srcFP)
	dstEdge := rectFixedPointV1EngineSceneConnection(dstRect, dstFP)
	srcTerminal := pageTerminalPointV1EngineSceneConnectionPage(srcFrameRect, srcEdge, srcFrameSide)
	srcFrameAnchorExplicit := false
	if anchor, ok := connectionFrameAnchorV1EngineSceneConnectionRoute(conn, "src"); ok && anchor.hasSlot && string(anchor.side) == srcFrameSide {
		srcTerminal = pageTerminalPointForAnchorV1EngineSceneConnectionPage(srcFrameRect, anchor)
		srcFrameAnchorExplicit = true
	}
	dstTerminal := pageTerminalPointV1EngineSceneConnectionPage(dstFrameRect, dstEdge, dstFrameSide)
	dstFrameAnchorExplicit := false
	if anchor, ok := connectionFrameAnchorV1EngineSceneConnectionRoute(conn, "dst"); ok && anchor.hasSlot && string(anchor.side) == dstFrameSide {
		dstTerminal = pageTerminalPointForAnchorV1EngineSceneConnectionPage(dstFrameRect, anchor)
		dstFrameAnchorExplicit = true
	}
	if !srcFrameAnchorExplicit {
		srcTerminal = avoidFrameMetadataTerminalV1EngineSceneConnectionPage(srcTerminal, srcFrameRect, srcFrameSide, srcMetadata)
	}
	if !dstFrameAnchorExplicit {
		dstTerminal = avoidFrameMetadataTerminalV1EngineSceneConnectionPage(dstTerminal, dstFrameRect, dstFrameSide, dstMetadata)
	}
	srcTerminal = insetPageTerminalV1EngineSceneConnectionPage(srcTerminal, srcFrameRect, srcFrameSide, srcMetadata)
	dstTerminal = insetPageTerminalV1EngineSceneConnectionPage(dstTerminal, dstFrameRect, dstFrameSide, dstMetadata)
	if !srcFrameAnchorExplicit {
		srcTerminal = shiftCoincidentInsetPageTerminalV1EngineSceneConnectionPage(srcTerminal, srcEdge, srcFrameRect, srcFrameSide, srcMetadata)
	}
	if !dstFrameAnchorExplicit {
		dstTerminal = shiftCoincidentInsetPageTerminalV1EngineSceneConnectionPage(dstTerminal, dstEdge, dstFrameRect, dstFrameSide, dstMetadata)
	}
	sourceID := fmt.Sprintf("conn-%s-to-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(srcKey), sanitizeElementIDV1EngineSceneConnectionRoute(dstFrameID), index)
	destID := fmt.Sprintf("conn-%s-from-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(dstKey), sanitizeElementIDV1EngineSceneConnectionRoute(srcFrameID), index)
	metadata := crossFrameConnectorMetadataV1EngineSceneConnectionPage{
		logicalID:            fmt.Sprintf("conn-%s-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(srcKey), sanitizeElementIDV1EngineSceneConnectionRoute(dstKey), index),
		sourceElementID:      srcElemID,
		destinationElementID: dstElemID,
	}
	appendCrossFrameArrowV1EngineSceneConnectionPage(elements, sourceID, srcEdge, srcTerminal, srcSide, srcFrameSide, false, srcFrameRect, srcFrameEndpoint, srcMetadata, style, seed, updated, map[string]any{
		"elementId":  srcElemID,
		"focus":      0.0,
		"gap":        5.0,
		"fixedPoint": []float64{srcFP[0], srcFP[1]},
	}, nil, srcFrameID, dstFrameID, srcFrameID, conn, metadata)
	appendCrossFrameLabelV1EngineSceneConnectionPage(elements, sourceID+"-label", "to <"+dstFrameID+">", srcTerminal, srcFrameRect, srcVisualRect, srcFrameSide, style.Color, srcFrameID, frameMetadataReservedRectsV1EngineSceneConnectionPage(srcMetadata), seed+1, updated)
	appendCrossFrameArrowV1EngineSceneConnectionPage(elements, destID, dstTerminal, dstEdge, dstSide, dstFrameSide, true, dstFrameRect, dstFrameEndpoint, dstMetadata, style, seed+2, updated, nil, map[string]any{
		"elementId":  dstElemID,
		"focus":      0.0,
		"gap":        5.0,
		"fixedPoint": []float64{dstFP[0], dstFP[1]},
	}, srcFrameID, dstFrameID, dstFrameID, conn, metadata)
	appendCrossFrameLabelV1EngineSceneConnectionPage(elements, destID+"-label", "from <"+srcFrameID+">", dstTerminal, dstFrameRect, dstVisualRect, dstFrameSide, style.Color, dstFrameID, frameMetadataReservedRectsV1EngineSceneConnectionPage(dstMetadata), seed+3, updated)
	boundMap[srcElemID] = append(boundMap[srcElemID], map[string]any{"type": "arrow", "id": sourceID})
	boundMap[dstElemID] = append(boundMap[dstElemID], map[string]any{"type": "arrow", "id": destID})
}

func appendCrossFrameArrowV1EngineSceneConnectionPage(elements *[]map[string]any, id string, start, end [2]float64, endpointSide, frameSide string, frameAtStart bool, frameRect [4]float64, frameEndpoint bool, frameMetadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata, style resolvedConnectionStyleV1EngineSceneTypes, seed int, updated int64, startBinding, endBinding any, srcFrameID, dstFrameID, ownerFrameID string, conn *entity.Node, metadata crossFrameConnectorMetadataV1EngineSceneConnectionPage) {
	dx := end[0] - start[0]
	dy := end[1] - start[1]
	points := crossFrameArrowPointsAvoidingMetadataV1EngineSceneConnectionPage(start, end, endpointSide, frameSide, frameAtStart, frameRect, frameEndpoint, frameMetadata)
	width, height := math.Abs(dx), math.Abs(dy)
	if len(points) > 0 {
		minX, maxX := points[0][0], points[0][0]
		minY, maxY := points[0][1], points[0][1]
		for _, point := range points[1:] {
			minX, maxX = math.Min(minX, point[0]), math.Max(maxX, point[0])
			minY, maxY = math.Min(minY, point[1]), math.Max(maxY, point[1])
		}
		width, height = math.Max(width, maxX-minX), math.Max(height, maxY-minY)
	}
	customData := map[string]any{
		"xaligoConnectorKind":                 style.Kind,
		"xaligoConnectorStartArrowhead":       style.StartArrowhead,
		"xaligoConnectorEndArrowhead":         style.EndArrowhead,
		"xaligoConnectorStyleSourceKnown":     true,
		"xaligoCrossFrame":                    true,
		"xaligoSourceFrame":                   srcFrameID,
		"xaligoDestinationFrame":              dstFrameID,
		"xaligoFrameID":                       ownerFrameID,
		"xaligoConnectorLogicalId":            metadata.logicalID,
		"xaligoConnectorSourceElementId":      metadata.sourceElementID,
		"xaligoConnectorDestinationElementId": metadata.destinationElementID,
	}
	if style.StartArrowheadExplicit {
		customData["xaligoConnectorStartArrowheadExplicit"] = true
	}
	if style.EndArrowheadExplicit {
		customData["xaligoConnectorEndArrowheadExplicit"] = true
	}
	if style.WidthExplicit {
		customData["xaligoConnectorStrokeWidthExplicit"] = true
	}
	if grid, ok := positiveFloatAttrV1EngineSceneConnectionRoute(conn, "grid"); ok {
		customData["xaligoConnectorGrid"] = grid
	}
	if bends := strings.TrimSpace(connectionBendsV1EngineSceneConnectionRoute(conn)); bends != "" {
		customData["xaligoConnectorBends"] = bends
	}
	if scale, ok := positiveFloatAttrV1EngineSceneConnectionRoute(conn, "coordinate-scale", "scale"); ok {
		customData["xaligoConnectorScale"] = scale
	}
	if anchor, ok := connectionEndpointAnchorV1EngineSceneConnectionRoute(conn, "src"); ok && anchor.hasSlot {
		customData["xaligoConnectorSrcAnchor"] = true
	}
	if anchor, ok := connectionEndpointAnchorV1EngineSceneConnectionRoute(conn, "dst"); ok && anchor.hasSlot {
		customData["xaligoConnectorDstAnchor"] = true
	}
	if anchor, ok := connectionFrameAnchorV1EngineSceneConnectionRoute(conn, "src"); ok {
		customData["xaligoConnectorSourceFrameSide"] = string(anchor.side)
		if anchor.hasSlot {
			customData["xaligoConnectorSourceFrameAnchor"] = anchor.StringV1EngineParseConnection()
		}
	}
	if anchor, ok := connectionFrameAnchorV1EngineSceneConnectionRoute(conn, "dst"); ok {
		customData["xaligoConnectorDestinationFrameSide"] = string(anchor.side)
		if anchor.hasSlot {
			customData["xaligoConnectorDestinationFrameAnchor"] = anchor.StringV1EngineParseConnection()
		}
	}
	applyDatabaseConnectionMetadataV1EngineSceneConnectionRender(customData, conn)
	applyConnectionDiffStatusV1EngineSceneDiffHighlight(customData, conn)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "arrow",
		"x": start[0], "y": start[1],
		"width": width, "height": height,
		"angle":       0,
		"strokeColor": style.Color, "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": style.Width, "strokeStyle": style.StrokeStyle,
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": map[string]any{"type": 2},
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false, "frameId": nil,
		"points":             points,
		"lastCommittedPoint": nil,
		"startBinding":       startBinding,
		"endBinding":         endBinding,
		"startArrowhead":     style.ExcalidrawStartArrowhead,
		"endArrowhead":       style.ExcalidrawEndArrowhead,
		"endArrowheadSize":   "s",
		"startArrowheadSize": "s",
		"elbowed":            false,
		"customData":         customData,
	})
}

func crossFrameArrowPointsAvoidingMetadataV1EngineSceneConnectionPage(start, end [2]float64, endpointSide, frameSide string, frameAtStart bool, frameRect [4]float64, frameEndpoint bool, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) [][]float64 {
	endpoint, terminal := start, end
	if frameAtStart {
		endpoint, terminal = end, start
	}
	const epsilon = 1e-9
	coincident := math.Abs(endpoint[0]-terminal[0]) <= epsilon && math.Abs(endpoint[1]-terminal[1]) <= epsilon
	endpointHorizontal := endpointSide == "left" || endpointSide == "right"
	frameHorizontal := frameSide == "left" || frameSide == "right"
	if !coincident && endpointHorizontal == frameHorizontal {
		aligned := (endpointHorizontal && math.Abs(endpoint[1]-terminal[1]) <= epsilon) ||
			(!endpointHorizontal && math.Abs(endpoint[0]-terminal[0]) <= epsilon)
		if aligned {
			return [][]float64{{0, 0}, {end[0] - start[0], end[1] - start[1]}}
		}
	}
	endpointApproach := pageLinkApproachForSideV1EngineSceneConnectionPage(frameRect, endpointSide)
	frameApproach := pageLinkApproachForSideV1EngineSceneConnectionPage(frameRect, frameSide)
	endpointDirection := pageLinkSideVectorV1EngineSceneConnectionPage(endpointSide)
	frameDirection := pageLinkSideVectorV1EngineSceneConnectionPage(frameSide)
	if frameEndpoint {
		endpointDirection[0] *= -1
		endpointDirection[1] *= -1
	}
	endpointStub := [2]float64{
		endpoint[0] + endpointDirection[0]*endpointApproach,
		endpoint[1] + endpointDirection[1]*endpointApproach,
	}
	terminalStub := [2]float64{
		terminal[0] - frameDirection[0]*frameApproach,
		terminal[1] - frameDirection[1]*frameApproach,
	}

	frameRight := frameRect[0] + frameRect[2]
	frameBottom := frameRect[1] + frameRect[3]
	if endpointStub[0] < frameRect[0] || endpointStub[0] > frameRight || endpointStub[1] < frameRect[1] || endpointStub[1] > frameBottom {
		// An item may touch a page edge. In that case the nominal outward
		// approach would leave the physical page, so approach the same item
		// side from the page interior while keeping the segment perpendicular.
		endpointStub = [2]float64{
			endpoint[0] - endpointDirection[0]*endpointApproach,
			endpoint[1] - endpointDirection[1]*endpointApproach,
		}
	}
	endpointStub[0] = clampFloatV1EngineLayoutPort(endpointStub[0], frameRect[0], frameRight)
	endpointStub[1] = clampFloatV1EngineLayoutPort(endpointStub[1], frameRect[1], frameBottom)
	terminalStub[0] = clampFloatV1EngineLayoutPort(terminalStub[0], frameRect[0], frameRight)
	safeTop, safeBottom := frameRect[1], frameBottom
	if metadata.Reserved[2] > 0 && metadata.Reserved[3] > 0 {
		const clearance = 8.0
		if metadata.Position == "bottom" {
			safeBottom = math.Min(safeBottom, metadata.Reserved[1]-clearance)
		} else {
			safeTop = math.Max(safeTop, metadata.Reserved[1]+metadata.Reserved[3]+clearance)
		}
		if safeBottom < safeTop {
			safeTop, safeBottom = frameRect[1], frameBottom
		}
	}
	// A left/right terminal's Y coordinate has already passed reservation
	// handling. Moving only the adjacent stub's tangent coordinate would turn
	// the final segment diagonal instead of adding clearance.
	if !frameHorizontal {
		terminalStub[1] = clampFloatV1EngineLayoutPort(terminalStub[1], safeTop, safeBottom)
	}

	absolute := make([][2]float64, 0, 6)
	appendPoint := func(point [2]float64) {
		if len(absolute) > 0 && math.Abs(absolute[len(absolute)-1][0]-point[0]) <= epsilon && math.Abs(absolute[len(absolute)-1][1]-point[1]) <= epsilon {
			return
		}
		if len(absolute) >= 2 {
			previous := absolute[len(absolute)-1]
			beforePrevious := absolute[len(absolute)-2]
			firstDirection := [2]float64{previous[0] - beforePrevious[0], previous[1] - beforePrevious[1]}
			secondDirection := [2]float64{point[0] - previous[0], point[1] - previous[1]}
			collinear := math.Abs(firstDirection[0]*secondDirection[1]-firstDirection[1]*secondDirection[0]) <= epsilon
			sameDirection := firstDirection[0]*secondDirection[0]+firstDirection[1]*secondDirection[1] >= 0
			if collinear && sameDirection {
				absolute[len(absolute)-1] = point
				return
			}
		}
		absolute = append(absolute, point)
	}
	appendPoint(endpoint)
	appendPoint(endpointStub)
	if endpointSide == "left" || endpointSide == "right" {
		appendPoint([2]float64{endpointStub[0], terminalStub[1]})
	} else {
		appendPoint([2]float64{terminalStub[0], endpointStub[1]})
	}
	appendPoint(terminalStub)
	appendPoint(terminal)
	if frameAtStart {
		for left, right := 0, len(absolute)-1; left < right; left, right = left+1, right-1 {
			absolute[left], absolute[right] = absolute[right], absolute[left]
		}
	}

	points := make([][]float64, 0, len(absolute))
	for _, point := range absolute {
		points = append(points, []float64{point[0] - start[0], point[1] - start[1]})
	}
	return points
}

func pageLinkApproachForSideV1EngineSceneConnectionPage(frameRect [4]float64, side string) float64 {
	size := frameRect[2]
	if side == "top" || side == "bottom" {
		size = frameRect[3]
	}
	return math.Min(pageLinkApproachV1EngineSceneConnectionPage, math.Max(1, size/4))
}

func pageLinkSideVectorV1EngineSceneConnectionPage(side string) [2]float64 {
	switch side {
	case "top":
		return [2]float64{0, -1}
	case "bottom":
		return [2]float64{0, 1}
	case "left":
		return [2]float64{-1, 0}
	default:
		return [2]float64{1, 0}
	}
}

func appendCrossFrameLabelV1EngineSceneConnectionPage(elements *[]map[string]any, id, label string, terminal [2]float64, frameRect, endpointRect [4]float64, side, color, frameID string, metadataRects [][4]float64, seed int, updated int64) {
	fontSize := 12.0
	safeTop, safeBottom, hasMetadataReservation := frameMetadataSafeVerticalIntervalV1EngineSceneConnectionPage(frameRect, metadataRects)
	if hasMetadataReservation {
		fontSize = math.Min(fontSize, (safeBottom-safeTop)/1.2)
	}
	w := textWidthV1EngineSceneItem(label, fontSize*0.5)
	h := math.Ceil(fontSize * 1.2)
	availableW := frameRect[2] - 2*pageLinkLabelGapV1EngineSceneConnectionPage
	if availableW <= 0 {
		availableW = frameRect[2]
	}
	availableH := frameRect[3] - 2*pageLinkLabelGapV1EngineSceneConnectionPage
	if availableH <= 0 {
		availableH = frameRect[3]
	}
	frameRight := frameRect[0] + frameRect[2]
	frameBottom := frameRect[1] + frameRect[3]
	switch side {
	case "left":
		availableW = frameRight - terminal[0] - 2*pageLinkLabelGapV1EngineSceneConnectionPage
	case "right":
		availableW = terminal[0] - frameRect[0] - 2*pageLinkLabelGapV1EngineSceneConnectionPage
	case "top":
		availableH = frameBottom - terminal[1] - 2*pageLinkLabelGapV1EngineSceneConnectionPage
	case "bottom":
		availableH = terminal[1] - frameRect[1] - 2*pageLinkLabelGapV1EngineSceneConnectionPage
	}
	w = math.Min(w, math.Max(0.1, availableW))
	h = math.Min(h, math.Max(0.1, availableH))
	if hasMetadataReservation {
		h = math.Min(h, safeBottom-safeTop)
	}
	x, y := pageLinkLabelPositionV1EngineSceneConnectionPage(terminal, frameRect, endpointRect, w, h, side)
	x, y = pageLinkLabelPositionAvoidingMetadataV1EngineSceneConnectionPage(x, y, w, h, terminal, frameRect, endpointRect, side, metadataRects)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "text",
		"x": x, "y": y, "width": w, "height": h,
		"angle":       0,
		"strokeColor": color, "backgroundColor": "#ffffff",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"text": label, "fontSize": fontSize, "fontFamily": 1,
		"textAlign": "center", "verticalAlign": "middle",
		"containerId": nil, "originalText": label, "lineHeight": 1.2,
		"customData": map[string]any{
			"xaligoCrossFrameLabel": true,
			"xaligoFrameID":         frameID,
			"xaligoTextLayout":      sceneTextLayoutV1EngineSceneBuild(entity.TextRoleConnectorLabel, true, 1.2),
		},
	})
}

func avoidFrameMetadataTerminalV1EngineSceneConnectionPage(terminal [2]float64, frameRect [4]float64, side string, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) [2]float64 {
	if (side == "left" || side == "right") && metadata.Reserved[2] > 0 && metadata.Reserved[3] > 0 {
		const clearance = 8.0
		if metadata.Position == "bottom" {
			terminal[1] = math.Min(terminal[1], metadata.Reserved[1]-clearance)
		} else {
			terminal[1] = math.Max(terminal[1], metadata.Reserved[1]+metadata.Reserved[3]+clearance)
		}
		terminal[1] = clampFloatV1EngineLayoutPort(terminal[1], frameRect[1], frameRect[1]+frameRect[3])
		return terminal
	}
	if len(metadata.Rects) == 0 || metadata.Position != side || side != "top" && side != "bottom" {
		return terminal
	}
	const clearance = 8.0
	const cornerGutter = 24.0
	adaptiveGutter := func(size float64) float64 { return math.Min(cornerGutter, math.Max(0, size/4)) }
	minimum := frameRect[0] + adaptiveGutter(frameRect[2])
	maximum := frameRect[0] + frameRect[2] - adaptiveGutter(frameRect[2])
	intervals := make([][2]float64, 0, len(metadata.Rects))
	for _, rect := range metadata.Rects {
		start := rect[0] - clearance
		end := rect[0] + rect[2] + clearance
		if end > minimum && start < maximum {
			intervals = append(intervals, [2]float64{start, end})
		}
	}
	if len(intervals) == 0 {
		return terminal
	}
	if coordinate, ok := nearestFreeCoordinateV1EngineSceneConnectionPage(terminal[0], minimum, maximum, intervals); ok {
		terminal[0] = coordinate
		return terminal
	}
	if coordinate, ok := nearestFreeCoordinateV1EngineSceneConnectionPage(terminal[0], frameRect[0]+1, frameRect[0]+frameRect[2]-1, intervals); ok {
		terminal[0] = coordinate
		return terminal
	}
	if math.Abs(terminal[0]-frameRect[0]) <= math.Abs(terminal[0]-(frameRect[0]+frameRect[2])) {
		terminal[0] = frameRect[0]
	} else {
		terminal[0] = frameRect[0] + frameRect[2]
	}
	return terminal
}

func insetPageTerminalV1EngineSceneConnectionPage(terminal [2]float64, frameRect [4]float64, side string, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) [2]float64 {
	inset := defaultFrameMetadataRowGapV1EngineLayoutFrameMetadata
	if metadata.HasPageInset {
		inset = metadata.PageInset
	}
	if inset <= 0 {
		return terminal
	}
	switch side {
	case "top":
		terminal[1] = frameRect[1] + inset
	case "right":
		terminal[0] = frameRect[0] + frameRect[2] - inset
	case "bottom":
		terminal[1] = frameRect[1] + frameRect[3] - inset
	case "left":
		terminal[0] = frameRect[0] + inset
	}
	return terminal
}

func shiftCoincidentInsetPageTerminalV1EngineSceneConnectionPage(terminal, endpoint [2]float64, frameRect [4]float64, side string, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) [2]float64 {
	const cornerGutter = 24.0
	adaptiveGutter := func(size float64) float64 { return math.Min(cornerGutter, math.Max(0, size/4)) }
	frameTop := frameRect[1]
	frameBottom := frameRect[1] + frameRect[3]
	minX := frameRect[0] + adaptiveGutter(frameRect[2])
	maxX := frameRect[0] + frameRect[2] - adaptiveGutter(frameRect[2])
	minY := frameTop + adaptiveGutter(frameRect[3])
	maxY := frameBottom - adaptiveGutter(frameRect[3])
	if (side == "left" || side == "right") && metadata.Reserved[2] > 0 && metadata.Reserved[3] > 0 {
		const clearance = 8.0
		if metadata.Position == "bottom" {
			maxY = math.Min(maxY, metadata.Reserved[1]-clearance)
		} else {
			minY = math.Max(minY, metadata.Reserved[1]+metadata.Reserved[3]+clearance)
		}
		if minY > maxY {
			// Tiny metadata-free intervals cannot retain both the normal corner
			// gutter and the preferred reservation clearance. Fall back to the
			// full non-reserved interval so the tangent shift stays inside the
			// frame and never enters the metadata strip.
			if metadata.Position == "bottom" {
				minY = frameTop
				maxY = clampFloatV1EngineLayoutPort(metadata.Reserved[1], frameTop, frameBottom)
			} else {
				minY = clampFloatV1EngineLayoutPort(metadata.Reserved[1]+metadata.Reserved[3], frameTop, frameBottom)
				maxY = frameBottom
			}
		}
	}
	return shiftCoincidentPageTerminalV1EngineSceneConnectionPage(terminal, endpoint, side, minX, maxX, minY, maxY, cornerGutter)
}

func frameMetadataReservedRectsV1EngineSceneConnectionPage(metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) [][4]float64 {
	if metadata.Reserved[2] <= 0 || metadata.Reserved[3] <= 0 {
		return nil
	}
	return [][4]float64{metadata.Reserved}
}

func nearestFreeCoordinateV1EngineSceneConnectionPage(preferred, minimum, maximum float64, intervals [][2]float64) (float64, bool) {
	if maximum <= minimum {
		return preferred, false
	}
	clipped := make([][2]float64, 0, len(intervals))
	for _, interval := range intervals {
		start := math.Max(minimum, interval[0])
		end := math.Min(maximum, interval[1])
		if end > start {
			clipped = append(clipped, [2]float64{start, end})
		}
	}
	if len(clipped) == 0 {
		return clampFloatV1EngineLayoutPort(preferred, minimum, maximum), true
	}
	sort.Slice(clipped, func(left, right int) bool { return clipped[left][0] < clipped[right][0] })
	merged := make([][2]float64, 0, len(intervals))
	for _, interval := range clipped {
		if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
			continue
		}
		merged[len(merged)-1][1] = math.Max(merged[len(merged)-1][1], interval[1])
	}
	free := make([][2]float64, 0, len(merged)+1)
	cursor := minimum
	for _, interval := range merged {
		if interval[0] > cursor {
			free = append(free, [2]float64{cursor, interval[0]})
		}
		cursor = math.Max(cursor, interval[1])
	}
	if cursor < maximum {
		free = append(free, [2]float64{cursor, maximum})
	}
	if len(free) == 0 {
		return preferred, false
	}
	preferred = clampFloatV1EngineLayoutPort(preferred, minimum, maximum)
	best := preferred
	bestDistance := math.Inf(1)
	for _, gap := range free {
		candidate := clampFloatV1EngineLayoutPort(preferred, gap[0], gap[1])
		if distance := math.Abs(candidate - preferred); distance < bestDistance {
			best = candidate
			bestDistance = distance
		}
	}
	return best, true
}

func pageLinkLabelPositionAvoidingMetadataV1EngineSceneConnectionPage(x, y, w, h float64, terminal [2]float64, frameRect, endpointRect [4]float64, side string, metadataRects [][4]float64) (float64, float64) {
	if len(metadataRects) == 0 || !labelOverlapsMetadataV1EngineSceneConnectionPage(x, y, w, h, metadataRects) {
		return x, y
	}
	const gap = pageLinkLabelGapV1EngineSceneConnectionPage
	frameRight := frameRect[0] + frameRect[2]
	frameBottom := frameRect[1] + frameRect[3]
	minimumY := frameRect[1] + gap
	maximumY := math.Max(minimumY, frameBottom-h-gap)
	if safeTop, safeBottom, ok := frameMetadataSafeVerticalIntervalV1EngineSceneConnectionPage(frameRect, metadataRects); ok {
		verticalInset := math.Min(gap, math.Max(0, (safeBottom-safeTop-h)/2))
		minimumY = safeTop + verticalInset
		maximumY = math.Max(minimumY, safeBottom-h-verticalInset)
	}
	candidates := [][2]float64{}
	appendCandidate := func(candidateX, candidateY float64) {
		candidateX = clampFloatV1EngineLayoutPort(candidateX, frameRect[0]+gap, math.Max(frameRect[0]+gap, frameRight-w-gap))
		candidateY = clampFloatV1EngineLayoutPort(candidateY, minimumY, maximumY)
		candidates = append(candidates, [2]float64{candidateX, candidateY})
	}
	switch side {
	case "top", "bottom":
		edgeY := terminal[1] + gap
		if side == "bottom" {
			edgeY = terminal[1] - h - gap
		}
		appendCandidate(terminal[0]+gap, edgeY)
		appendCandidate(terminal[0]-w-gap, edgeY)
		appendCandidate(frameRect[0]+gap, edgeY)
		appendCandidate(frameRight-w-gap, edgeY)
		for _, rect := range metadataRects {
			appendCandidate(rect[0]+rect[2]+gap, edgeY)
			appendCandidate(rect[0]-w-gap, edgeY)
		}
		if side == "top" {
			metadataBottom := frameRect[1]
			for _, rect := range metadataRects {
				metadataBottom = math.Max(metadataBottom, rect[1]+rect[3])
			}
			appendCandidate(terminal[0]+gap, metadataBottom+gap)
			appendCandidate(terminal[0]-w-gap, metadataBottom+gap)
		} else {
			metadataTop := frameBottom
			for _, rect := range metadataRects {
				metadataTop = math.Min(metadataTop, rect[1])
			}
			appendCandidate(terminal[0]+gap, metadataTop-h-gap)
			appendCandidate(terminal[0]-w-gap, metadataTop-h-gap)
		}
	case "left", "right":
		edgeX := terminal[0] + gap
		if side == "right" {
			edgeX = terminal[0] - w - gap
		}
		appendCandidate(edgeX, terminal[1]+gap)
		appendCandidate(edgeX, terminal[1]-h-gap)
		for _, rect := range metadataRects {
			appendCandidate(edgeX, rect[1]+rect[3]+gap)
			appendCandidate(edgeX, rect[1]-h-gap)
		}
	}
	for _, candidate := range candidates {
		if !labelOverlapsMetadataV1EngineSceneConnectionPage(candidate[0], candidate[1], w, h, metadataRects) &&
			!pageLinkRectOverlapsV1EngineSceneConnectionPage(candidate[0], candidate[1], w, h, endpointRect) {
			return candidate[0], candidate[1]
		}
	}
	if safeTop, safeBottom, ok := frameMetadataSafeVerticalIntervalV1EngineSceneConnectionPage(frameRect, metadataRects); ok && safeBottom-safeTop >= h {
		return clampFloatV1EngineLayoutPort(x, frameRect[0], math.Max(frameRect[0], frameRight-w)),
			clampFloatV1EngineLayoutPort(y, safeTop, safeBottom-h)
	}
	return x, y
}

func frameMetadataSafeVerticalIntervalV1EngineSceneConnectionPage(frameRect [4]float64, metadataRects [][4]float64) (float64, float64, bool) {
	frameTop := frameRect[1]
	frameBottom := frameRect[1] + frameRect[3]
	safeTop, safeBottom := frameTop, frameBottom
	reserved := false
	for _, rect := range metadataRects {
		if rect[2] <= 0 || rect[3] <= 0 {
			continue
		}
		reserved = true
		rectBottom := rect[1] + rect[3]
		switch {
		case rect[1] <= frameTop+geometryEpsilonV1EngineLayoutValidation:
			safeTop = math.Max(safeTop, rectBottom)
		case rectBottom >= frameBottom-geometryEpsilonV1EngineLayoutValidation:
			safeBottom = math.Min(safeBottom, rect[1])
		case rect[1]+rect[3]/2 <= frameTop+frameRect[3]/2:
			safeTop = math.Max(safeTop, rectBottom)
		default:
			safeBottom = math.Min(safeBottom, rect[1])
		}
	}
	return safeTop, safeBottom, reserved && safeBottom > safeTop
}

func labelOverlapsMetadataV1EngineSceneConnectionPage(x, y, w, h float64, metadataRects [][4]float64) bool {
	for _, rect := range metadataRects {
		if pageLinkRectOverlapsV1EngineSceneConnectionPage(x, y, w, h, rect) {
			return true
		}
	}
	return false
}

func pageLinkRectOverlapsV1EngineSceneConnectionPage(x, y, w, h float64, rect [4]float64) bool {
	return x < rect[0]+rect[2] && x+w > rect[0] && y < rect[1]+rect[3] && y+h > rect[1]
}

func pageLinkLabelPositionV1EngineSceneConnectionPage(terminal [2]float64, frameRect, endpointRect [4]float64, w, h float64, side string) (float64, float64) {
	const gap = pageLinkLabelGapV1EngineSceneConnectionPage
	frameRight := frameRect[0] + frameRect[2]
	frameBottom := frameRect[1] + frameRect[3]
	edgeX := terminal[0] + gap
	if side == "right" {
		edgeX = terminal[0] - w - gap
	}
	edgeY := terminal[1] + gap
	if side == "bottom" {
		edgeY = terminal[1] - h - gap
	}
	candidates := make([][2]float64, 0, 6)
	switch side {
	case "left", "right":
		candidates = append(candidates,
			[2]float64{edgeX, terminal[1] + gap},
			[2]float64{edgeX, terminal[1] - h - gap},
			[2]float64{edgeX, endpointRect[1] + endpointRect[3] + gap},
			[2]float64{edgeX, endpointRect[1] - h - gap},
		)
	case "top", "bottom":
		candidates = append(candidates,
			[2]float64{terminal[0] + gap, edgeY},
			[2]float64{terminal[0] - w - gap, edgeY},
			[2]float64{endpointRect[0] + endpointRect[2] + gap, edgeY},
			[2]float64{endpointRect[0] - w - gap, edgeY},
		)
	}
	inside := func(point [2]float64) bool {
		return point[0] >= frameRect[0]+gap && point[0]+w <= frameRight-gap && point[1] >= frameRect[1]+gap && point[1]+h <= frameBottom-gap
	}
	overlapsEndpoint := func(point [2]float64) bool {
		return pageLinkRectOverlapsV1EngineSceneConnectionPage(point[0], point[1], w, h, endpointRect)
	}
	for _, candidate := range candidates {
		if inside(candidate) && !overlapsEndpoint(candidate) {
			return candidate[0], candidate[1]
		}
	}
	if len(candidates) == 0 {
		candidates = append(candidates, [2]float64{terminal[0] - w/2, terminal[1] - h/2})
	}
	return clampFloatV1EngineLayoutPort(candidates[0][0], frameRect[0]+gap, math.Max(frameRect[0]+gap, frameRight-w-gap)),
		clampFloatV1EngineLayoutPort(candidates[0][1], frameRect[1]+gap, math.Max(frameRect[1]+gap, frameBottom-h-gap))
}

func endpointVisualRectV1EngineSceneConnectionPage(imageRect, labelRect [4]float64) [4]float64 {
	if labelRect[2] <= 0 || labelRect[3] <= 0 {
		return imageRect
	}
	minX := math.Min(imageRect[0], labelRect[0])
	minY := math.Min(imageRect[1], labelRect[1])
	maxX := math.Max(imageRect[0]+imageRect[2], labelRect[0]+labelRect[2])
	maxY := math.Max(imageRect[1]+imageRect[3], labelRect[1]+labelRect[3])
	return [4]float64{minX, minY, maxX - minX, maxY - minY}
}

func nearestFrameSideV1EngineSceneConnectionPage(frameRect, endpointRect, otherFrameRect [4]float64) string {
	return nearestFrameSideExcludingV1EngineSceneConnectionPage(frameRect, endpointRect, otherFrameRect, "")
}

func pageLinkSideAvoidingFrameMetadataV1EngineSceneConnectionPage(side string, frameRect, endpointRect, otherFrameRect [4]float64, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) string {
	if metadata.Reserved[2] <= 0 || metadata.Reserved[3] <= 0 || side != metadata.Position {
		return side
	}
	return nearestFrameSideExcludingV1EngineSceneConnectionPage(frameRect, endpointRect, otherFrameRect, metadata.Position)
}

func pageLinkSideAvoidingUnsafeInsetV1EngineSceneConnectionPage(side string, frameRect, endpointRect, otherFrameRect [4]float64, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) string {
	pageInset := defaultFrameMetadataRowGapV1EngineLayoutFrameMetadata
	if metadata.HasPageInset {
		pageInset = metadata.PageInset
	}
	if pageLinkTerminalSideIsSafeV1EngineLayoutFrameMetadata(side, frameRect, pageInset, metadata.Position, metadata.Reserved) {
		return side
	}
	if safeSide, ok := nearestFrameSideMatchingV1EngineSceneConnectionPage(frameRect, endpointRect, otherFrameRect, func(candidate string) bool {
		return pageLinkTerminalSideIsSafeV1EngineLayoutFrameMetadata(candidate, frameRect, pageInset, metadata.Position, metadata.Reserved)
	}); ok {
		return safeSide
	}
	return side
}

func nearestFrameSideExcludingV1EngineSceneConnectionPage(frameRect, endpointRect, otherFrameRect [4]float64, excluded string) string {
	side, ok := nearestFrameSideMatchingV1EngineSceneConnectionPage(frameRect, endpointRect, otherFrameRect, func(candidate string) bool {
		return candidate != excluded
	})
	if ok {
		return side
	}
	return "top"
}

func nearestFrameSideMatchingV1EngineSceneConnectionPage(frameRect, endpointRect, otherFrameRect [4]float64, allowed func(string) bool) (string, bool) {
	distances := map[string]float64{
		"top":    math.Max(0, endpointRect[1]-frameRect[1]),
		"right":  math.Max(0, frameRect[0]+frameRect[2]-(endpointRect[0]+endpointRect[2])),
		"bottom": math.Max(0, frameRect[1]+frameRect[3]-(endpointRect[1]+endpointRect[3])),
		"left":   math.Max(0, endpointRect[0]-frameRect[0]),
	}
	minimum := math.Inf(1)
	for _, side := range []string{"top", "right", "bottom", "left"} {
		if allowed(side) {
			minimum = math.Min(minimum, distances[side])
		}
	}
	if math.IsInf(minimum, 1) {
		return "", false
	}
	const epsilon = 1e-9
	ownCx := frameRect[0] + frameRect[2]/2
	ownCy := frameRect[1] + frameRect[3]/2
	otherCx := otherFrameRect[0] + otherFrameRect[2]/2
	otherCy := otherFrameRect[1] + otherFrameRect[3]/2
	remoteSide, _ := connectionSideV1EngineSceneConnection(ownCx, ownCy, otherCx, otherCy)
	if allowed(remoteSide) && distances[remoteSide] <= minimum+epsilon {
		return remoteSide, true
	}
	for _, side := range []string{"top", "right", "bottom", "left"} {
		if allowed(side) && distances[side] <= minimum+epsilon {
			return side, true
		}
	}
	return "", false
}

func pageTerminalPointV1EngineSceneConnectionPage(frameRect [4]float64, anchor [2]float64, side string) [2]float64 {
	const cornerGutter = 24.0
	adaptiveGutter := func(size float64) float64 { return math.Min(cornerGutter, math.Max(0, size/4)) }
	gutterX := adaptiveGutter(frameRect[2])
	gutterY := adaptiveGutter(frameRect[3])
	minX, maxX := frameRect[0]+gutterX, frameRect[0]+frameRect[2]-gutterX
	minY, maxY := frameRect[1]+gutterY, frameRect[1]+frameRect[3]-gutterY
	var terminal [2]float64
	switch side {
	case "left":
		terminal = [2]float64{frameRect[0], clampFloatV1EngineLayoutPort(anchor[1], minY, maxY)}
	case "right":
		terminal = [2]float64{frameRect[0] + frameRect[2], clampFloatV1EngineLayoutPort(anchor[1], minY, maxY)}
	case "top":
		terminal = [2]float64{clampFloatV1EngineLayoutPort(anchor[0], minX, maxX), frameRect[1]}
	default:
		terminal = [2]float64{clampFloatV1EngineLayoutPort(anchor[0], minX, maxX), frameRect[1] + frameRect[3]}
	}
	return shiftCoincidentPageTerminalV1EngineSceneConnectionPage(terminal, anchor, side, minX, maxX, minY, maxY, cornerGutter)
}

func pageTerminalPointForAnchorV1EngineSceneConnectionPage(frameRect [4]float64, anchor connectionAnchorSpecV1EngineParseConnection) [2]float64 {
	return rectFixedPointV1EngineSceneConnection(frameRect, fixedPointForAnchorV1EngineSceneConnection(anchor))
}

func shiftCoincidentPageTerminalV1EngineSceneConnectionPage(terminal, anchor [2]float64, side string, minX, maxX, minY, maxY, distance float64) [2]float64 {
	const epsilon = 1e-9
	if math.Abs(terminal[0]-anchor[0]) > epsilon || math.Abs(terminal[1]-anchor[1]) > epsilon {
		return terminal
	}
	shift := func(value, minimum, maximum float64) float64 {
		if value+distance <= maximum {
			return value + distance
		}
		if value-distance >= minimum {
			return value - distance
		}
		if maximum-value >= value-minimum {
			return maximum
		}
		return minimum
	}
	if side == "top" || side == "bottom" {
		terminal[0] = shift(terminal[0], minX, maxX)
	} else {
		terminal[1] = shift(terminal[1], minY, maxY)
	}
	return terminal
}
