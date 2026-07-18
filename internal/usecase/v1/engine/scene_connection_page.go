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

func renderCrossFrameConnectionV1EngineSceneConnectionPage(conn *entity.Node, srcKey, dstKey, srcElemID, dstElemID string, srcRect, dstRect [4]float64, srcFP, dstFP [2]float64, srcFrameID, dstFrameID string, srcFrameRect, dstFrameRect [4]float64, style resolvedConnectionStyleV1EngineSceneTypes, seed, index int, updated int64, elements *[]map[string]any, boundMap map[string][]map[string]any) {
	srcEdge := rectFixedPointV1EngineSceneConnection(srcRect, srcFP)
	dstEdge := rectFixedPointV1EngineSceneConnection(dstRect, dstFP)
	srcTerminal, srcSide := pageTerminalPointV1EngineSceneConnectionPage(srcFrameRect, srcEdge, dstFrameRect)
	dstTerminal, dstSide := pageTerminalPointV1EngineSceneConnectionPage(dstFrameRect, dstEdge, srcFrameRect)
	sourceID := fmt.Sprintf("conn-%s-to-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(srcKey), sanitizeElementIDV1EngineSceneConnectionRoute(dstFrameID), index)
	destID := fmt.Sprintf("conn-%s-from-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(dstKey), sanitizeElementIDV1EngineSceneConnectionRoute(srcFrameID), index)
	metadata := crossFrameConnectorMetadataV1EngineSceneConnectionPage{
		logicalID:            fmt.Sprintf("conn-%s-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(srcKey), sanitizeElementIDV1EngineSceneConnectionRoute(dstKey), index),
		sourceElementID:      srcElemID,
		destinationElementID: dstElemID,
	}
	appendCrossFrameArrowV1EngineSceneConnectionPage(elements, sourceID, srcEdge, srcTerminal, style, seed, updated, map[string]any{
		"elementId":  srcElemID,
		"focus":      0.0,
		"gap":        5.0,
		"fixedPoint": []float64{srcFP[0], srcFP[1]},
	}, nil, srcFrameID, dstFrameID, conn, metadata)
	appendCrossFrameLabelV1EngineSceneConnectionPage(elements, sourceID+"-label", "to "+dstFrameID, srcTerminal, srcSide, style.Color, seed+1, updated)
	appendCrossFrameArrowV1EngineSceneConnectionPage(elements, destID, dstTerminal, dstEdge, style, seed+2, updated, nil, map[string]any{
		"elementId":  dstElemID,
		"focus":      0.0,
		"gap":        5.0,
		"fixedPoint": []float64{dstFP[0], dstFP[1]},
	}, srcFrameID, dstFrameID, conn, metadata)
	appendCrossFrameLabelV1EngineSceneConnectionPage(elements, destID+"-label", "from "+srcFrameID, dstTerminal, dstSide, style.Color, seed+3, updated)
	boundMap[srcElemID] = append(boundMap[srcElemID], map[string]any{"type": "arrow", "id": sourceID})
	boundMap[dstElemID] = append(boundMap[dstElemID], map[string]any{"type": "arrow", "id": destID})
}

func appendCrossFrameArrowV1EngineSceneConnectionPage(elements *[]map[string]any, id string, start, end [2]float64, style resolvedConnectionStyleV1EngineSceneTypes, seed int, updated int64, startBinding, endBinding any, srcFrameID, dstFrameID string, conn *entity.Node, metadata crossFrameConnectorMetadataV1EngineSceneConnectionPage) {
	dx := end[0] - start[0]
	dy := end[1] - start[1]
	customData := map[string]any{
		"xaligoConnectorKind":                 style.Kind,
		"xaligoConnectorStartArrowhead":       style.StartArrowhead,
		"xaligoConnectorEndArrowhead":         style.EndArrowhead,
		"xaligoConnectorStyleSourceKnown":     true,
		"xaligoCrossFrame":                    true,
		"xaligoSourceFrame":                   srcFrameID,
		"xaligoDestinationFrame":              dstFrameID,
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
		"points":             [][]float64{{0, 0}, {dx, dy}},
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

func appendCrossFrameLabelV1EngineSceneConnectionPage(elements *[]map[string]any, id, label string, terminal [2]float64, side, color string, seed int, updated int64) {
	fontSize := 12.0
	w := textWidthV1EngineSceneItem(label, fontSize*0.5)
	h := math.Ceil(fontSize * 1.2)
	x := terminal[0] - w/2
	y := terminal[1] - h/2
	switch side {
	case "left":
		x = terminal[0] + 6
	case "right":
		x = terminal[0] - w - 6
	case "top":
		y = terminal[1] + 6
	case "bottom":
		y = terminal[1] - h - 6
	}
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
			"xaligoTextLayout":      sceneTextLayoutV1EngineSceneBuild(entity.TextRoleConnectorLabel, true, 1.2),
		},
	})
}

func pageTerminalPointV1EngineSceneConnectionPage(frameRect [4]float64, anchor [2]float64, otherFrameRect [4]float64) ([2]float64, string) {
	ownCx := frameRect[0] + frameRect[2]/2
	ownCy := frameRect[1] + frameRect[3]/2
	otherCx := otherFrameRect[0] + otherFrameRect[2]/2
	otherCy := otherFrameRect[1] + otherFrameRect[3]/2
	side, _ := connectionSideV1EngineSceneConnection(ownCx, ownCy, otherCx, otherCy)
	const margin = 24.0
	minX := frameRect[0] + margin
	maxX := frameRect[0] + frameRect[2] - margin
	minY := frameRect[1] + margin
	maxY := frameRect[1] + frameRect[3] - margin
	switch side {
	case "left":
		return [2]float64{minX, clampFloatV1EngineLayoutPort(anchor[1], minY, maxY)}, side
	case "right":
		return [2]float64{maxX, clampFloatV1EngineLayoutPort(anchor[1], minY, maxY)}, side
	case "top":
		return [2]float64{clampFloatV1EngineLayoutPort(anchor[0], minX, maxX), minY}, side
	default:
		return [2]float64{clampFloatV1EngineLayoutPort(anchor[0], minX, maxX), maxY}, side
	}
}
