package engine

import (
	"fmt"
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type crossFrameConnectorMetadataV1EngineSceneConnectionPage struct {
	logicalID            string
	sourceElementID      string
	destinationElementID string
}

func renderCrossFrameConnectionV1EngineSceneConnectionPage(conn *entity.Node, srcKey, dstKey, srcElemID, dstElemID string, srcRect, dstRect, srcVisualRect, dstVisualRect [4]float64, srcFP, dstFP [2]float64, srcSide, dstSide, srcFrameID, dstFrameID string, srcFrameRect, dstFrameRect [4]float64, style resolvedConnectionStyleV1EngineSceneTypes, seed, index int, updated int64, elements *[]map[string]any, boundMap map[string][]map[string]any) {
	srcEdge := rectFixedPointV1EngineSceneConnection(srcRect, srcFP)
	dstEdge := rectFixedPointV1EngineSceneConnection(dstRect, dstFP)
	srcTerminal := pageTerminalPointV1EngineSceneConnectionPage(srcFrameRect, srcEdge, srcSide)
	dstTerminal := pageTerminalPointV1EngineSceneConnectionPage(dstFrameRect, dstEdge, dstSide)
	sourceID := fmt.Sprintf("conn-%s-to-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(srcKey), sanitizeElementIDV1EngineSceneConnectionRoute(dstFrameID), index)
	destID := fmt.Sprintf("conn-%s-from-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(dstKey), sanitizeElementIDV1EngineSceneConnectionRoute(srcFrameID), index)
	metadata := crossFrameConnectorMetadataV1EngineSceneConnectionPage{
		logicalID:            fmt.Sprintf("conn-%s-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(srcKey), sanitizeElementIDV1EngineSceneConnectionRoute(dstKey), index),
		sourceElementID:      srcElemID,
		destinationElementID: dstElemID,
	}
	appendCrossFrameArrowV1EngineSceneConnectionPage(elements, sourceID, srcEdge, srcTerminal, srcSide, style, seed, updated, map[string]any{
		"elementId":  srcElemID,
		"focus":      0.0,
		"gap":        5.0,
		"fixedPoint": []float64{srcFP[0], srcFP[1]},
	}, nil, srcFrameID, dstFrameID, srcFrameID, conn, metadata)
	appendCrossFrameLabelV1EngineSceneConnectionPage(elements, sourceID+"-label", "to <"+dstFrameID+">", srcTerminal, srcFrameRect, srcVisualRect, srcSide, style.Color, srcFrameID, seed+1, updated)
	appendCrossFrameArrowV1EngineSceneConnectionPage(elements, destID, dstTerminal, dstEdge, dstSide, style, seed+2, updated, nil, map[string]any{
		"elementId":  dstElemID,
		"focus":      0.0,
		"gap":        5.0,
		"fixedPoint": []float64{dstFP[0], dstFP[1]},
	}, srcFrameID, dstFrameID, dstFrameID, conn, metadata)
	appendCrossFrameLabelV1EngineSceneConnectionPage(elements, destID+"-label", "from <"+srcFrameID+">", dstTerminal, dstFrameRect, dstVisualRect, dstSide, style.Color, dstFrameID, seed+3, updated)
	boundMap[srcElemID] = append(boundMap[srcElemID], map[string]any{"type": "arrow", "id": sourceID})
	boundMap[dstElemID] = append(boundMap[dstElemID], map[string]any{"type": "arrow", "id": destID})
}

func appendCrossFrameArrowV1EngineSceneConnectionPage(elements *[]map[string]any, id string, start, end [2]float64, side string, style resolvedConnectionStyleV1EngineSceneTypes, seed int, updated int64, startBinding, endBinding any, srcFrameID, dstFrameID, ownerFrameID string, conn *entity.Node, metadata crossFrameConnectorMetadataV1EngineSceneConnectionPage) {
	dx := end[0] - start[0]
	dy := end[1] - start[1]
	points := crossFrameArrowPointsV1EngineSceneConnectionPage(dx, dy, side)
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
	applyDatabaseConnectionMetadataV1EngineSceneConnectionRender(customData, conn)
	applyConnectionDiffStatusV1EngineSceneDiffHighlight(customData, conn)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "arrow",
		"x": start[0], "y": start[1],
		"width": math.Abs(dx), "height": math.Abs(dy),
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

func crossFrameArrowPointsV1EngineSceneConnectionPage(dx, dy float64, side string) [][]float64 {
	const epsilon = 1e-9
	points := [][]float64{{0, 0}}
	if math.Abs(dx) > epsilon && math.Abs(dy) > epsilon {
		if side == "left" || side == "right" {
			midX := dx / 2
			points = append(points, []float64{midX, 0}, []float64{midX, dy})
		} else {
			midY := dy / 2
			points = append(points, []float64{0, midY}, []float64{dx, midY})
		}
	}
	return append(points, []float64{dx, dy})
}

func appendCrossFrameLabelV1EngineSceneConnectionPage(elements *[]map[string]any, id, label string, terminal [2]float64, frameRect, endpointRect [4]float64, side, color, frameID string, seed int, updated int64) {
	fontSize := 12.0
	w := textWidthV1EngineSceneItem(label, fontSize*0.5)
	h := math.Ceil(fontSize * 1.2)
	availableW := frameRect[2] - 12
	if availableW <= 0 {
		availableW = frameRect[2]
	}
	availableH := frameRect[3] - 12
	if availableH <= 0 {
		availableH = frameRect[3]
	}
	w = math.Min(w, math.Max(0.1, availableW))
	h = math.Min(h, math.Max(0.1, availableH))
	x, y := pageLinkLabelPositionV1EngineSceneConnectionPage(terminal, frameRect, endpointRect, w, h, side)
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

func pageLinkLabelPositionV1EngineSceneConnectionPage(terminal [2]float64, frameRect, endpointRect [4]float64, w, h float64, side string) (float64, float64) {
	const gap = 6.0
	frameRight := frameRect[0] + frameRect[2]
	frameBottom := frameRect[1] + frameRect[3]
	x := terminal[0] - w/2
	y := terminal[1] - h/2
	switch side {
	case "left":
		x = terminal[0] + 6
		y = endpointRect[1] - h - gap
		if y < frameRect[1]+gap {
			y = endpointRect[1] + endpointRect[3] + gap
		}
	case "right":
		x = terminal[0] - w - 6
		y = endpointRect[1] - h - gap
		if y < frameRect[1]+gap {
			y = endpointRect[1] + endpointRect[3] + gap
		}
	case "top":
		y = terminal[1] + 6
		x = endpointRect[0] + endpointRect[2] + gap
		if x+w > frameRight-gap {
			x = endpointRect[0] - w - gap
		}
	case "bottom":
		y = terminal[1] - h - 6
		x = endpointRect[0] + endpointRect[2] + gap
		if x+w > frameRight-gap {
			x = endpointRect[0] - w - gap
		}
	}
	clamp := func(value, minimum, maximum float64) float64 {
		if minimum > maximum {
			return (minimum + maximum) / 2
		}
		return clampFloatV1EngineLayoutPort(value, minimum, maximum)
	}
	return clamp(x, frameRect[0]+gap, frameRight-w-gap), clamp(y, frameRect[1]+gap, frameBottom-h-gap)
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
	distances := map[string]float64{
		"top":    math.Max(0, endpointRect[1]-frameRect[1]),
		"right":  math.Max(0, frameRect[0]+frameRect[2]-(endpointRect[0]+endpointRect[2])),
		"bottom": math.Max(0, frameRect[1]+frameRect[3]-(endpointRect[1]+endpointRect[3])),
		"left":   math.Max(0, endpointRect[0]-frameRect[0]),
	}
	minimum := math.Min(math.Min(distances["top"], distances["right"]), math.Min(distances["bottom"], distances["left"]))
	const epsilon = 1e-9
	ownCx := frameRect[0] + frameRect[2]/2
	ownCy := frameRect[1] + frameRect[3]/2
	otherCx := otherFrameRect[0] + otherFrameRect[2]/2
	otherCy := otherFrameRect[1] + otherFrameRect[3]/2
	remoteSide, _ := connectionSideV1EngineSceneConnection(ownCx, ownCy, otherCx, otherCy)
	if distances[remoteSide] <= minimum+epsilon {
		return remoteSide
	}
	for _, side := range []string{"top", "right", "bottom", "left"} {
		if distances[side] <= minimum+epsilon {
			return side
		}
	}
	return "top"
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
