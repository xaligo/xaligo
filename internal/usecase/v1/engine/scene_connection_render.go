package engine

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

// renderConnections generates elbowed arrow elements for each <connection> node and
// updates the boundElements of the bound source/destination elements — required by
// Excalidraw so that the application recognises the binding relationship.
//
// src/dst are resolved item connection keys; the corresponding item rects and element IDs
// must already be populated in itemImgRects/itemLblRects/itemImgIDs/itemLblIDs by renderIconAt.
// Arrows start/end at the actual element edge; when the connection exits/enters from the
// bottom the label text element is used instead of the image element.
func renderConnectionsV1EngineSceneConnectionRender(connections []*entity.Node, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, frameRects map[string][4]float64, elements *[]map[string]any) {
	if len(connections) == 0 {
		return
	}
	updated := excalidrawUpdatedV1EngineSceneTypes

	// boundMap accumulates the arrow binding entries that must be written back
	// into each referenced element's boundElements array.
	// key = element ID, value = slice of {"type":"arrow","id":<arrowID>}
	boundMap := map[string][]map[string]any{}
	type junctionCandidate struct {
		edge  [2]float64
		side  string
		color string
		count int
		seed  int
	}
	junctionCandidates := map[string]*junctionCandidate{}

	orderedConnections := append([]*entity.Node(nil), connections...)
	sort.SliceStable(orderedConnections, func(i, j int) bool {
		return connectionKindPriorityV1EngineSceneConnectionRoute(connectionKindV1EngineSceneConnectionRoute(orderedConnections[i])) < connectionKindPriorityV1EngineSceneConnectionRoute(connectionKindV1EngineSceneConnectionRoute(orderedConnections[j]))
	})
	obstacles := excalidrawRouteObstaclesV1EngineSceneConnectionRoute(*elements)
	placed := [][]segmentV1EngineRouteTypes{}
	routePaths := map[string][]ptV1EngineRouteTypes{}

	for i, conn := range orderedConnections {
		srcIDStr := strings.TrimSpace(conn.Attrs["src"])
		dstIDStr := strings.TrimSpace(conn.Attrs["dst"])
		srcKey := strings.TrimSpace(conn.Attrs[internalConnectionSrcKeyAttrV1EngineParseDocument])
		dstKey := strings.TrimSpace(conn.Attrs[internalConnectionDstKeyAttrV1EngineParseDocument])
		if srcKey == "" || dstKey == "" {
			loggerV1EngineSharedLogging.WARN(IUESRC001V1EngineSceneTypes, "invalid connection source or destination", map[string]any{"src": srcIDStr, "dst": dstIDStr, "srcKey": srcKey, "dstKey": dstKey})
			continue
		}
		srcImgRect, srcOk := itemImgRects[srcKey]
		dstImgRect, dstOk := itemImgRects[dstKey]
		if !srcOk {
			loggerV1EngineSharedLogging.WARN(IUESRC002V1EngineSceneTypes, "source item not found or not rendered", map[string]any{"src": srcIDStr, "key": srcKey})
			continue
		}
		if !dstOk {
			loggerV1EngineSharedLogging.WARN(IUESRC003V1EngineSceneTypes, "destination item not found or not rendered", map[string]any{"dst": dstIDStr, "key": dstKey})
			continue
		}

		// Determine exit/entry side from image-center to image-center.
		srcCx := srcImgRect[0] + srcImgRect[2]/2
		srcCy := srcImgRect[1] + srcImgRect[3]/2
		dstCx := dstImgRect[0] + dstImgRect[2]/2
		dstCy := dstImgRect[1] + dstImgRect[3]/2
		srcSide, dstSide := connectionSideV1EngineSceneConnection(srcCx, srcCy, dstCx, dstCy)
		if conn.Attr("uml-diagram-kind") == "sequence-diagram" && srcIDStr == dstIDStr {
			srcSide, dstSide = "right", "right"
		}
		bends := connectionBendPointsV1EngineSceneConnection(conn)
		srcAnchor, hasSrcAnchor := connectionEndpointAnchorV1EngineSceneConnectionRoute(conn, "src")
		dstAnchor, hasDstAnchor := connectionEndpointAnchorV1EngineSceneConnectionRoute(conn, "dst")
		if hasSrcAnchor {
			srcSide = string(srcAnchor.side)
		} else if explicit, ok := connectionEndpointSideV1EngineSceneConnectionRoute(conn, "src"); ok {
			srcSide = string(explicit)
		} else if len(bends) > 0 {
			srcSide = sideTowardPointV1EngineSceneConnection(srcImgRect, bends[0])
		}
		if hasDstAnchor {
			dstSide = string(dstAnchor.side)
		} else if explicit, ok := connectionEndpointSideV1EngineSceneConnectionRoute(conn, "dst"); ok {
			dstSide = string(explicit)
		} else if len(bends) > 0 {
			dstSide = sideTowardPointV1EngineSceneConnection(dstImgRect, bends[len(bends)-1])
		}
		if _, ok := umlSequencePositionV1EngineSceneConnectionRoute(conn, "src"); ok {
			srcSide = umlSequenceVerticalSideV1EngineSceneConnectionRoute(srcSide)
			dstSide = umlSequenceVerticalSideV1EngineSceneConnectionRoute(dstSide)
		}

		// Choose element: bottom edge → label text box; other edges → image element.
		var srcElemID string
		var srcRect [4]float64
		if srcSide == "bottom" {
			if lblRect, ok := itemLblRects[srcKey]; ok {
				srcRect = lblRect
				srcElemID = itemLblIDs[srcKey]
			} else {
				srcRect = srcImgRect
				srcElemID = itemImgIDs[srcKey]
			}
		} else {
			srcRect = srcImgRect
			srcElemID = itemImgIDs[srcKey]
		}

		var dstElemID string
		var dstRect [4]float64
		if dstSide == "bottom" {
			if lblRect, ok := itemLblRects[dstKey]; ok {
				dstRect = lblRect
				dstElemID = itemLblIDs[dstKey]
			} else {
				dstRect = dstImgRect
				dstElemID = itemImgIDs[dstKey]
			}
		} else {
			dstRect = dstImgRect
			dstElemID = itemImgIDs[dstKey]
		}

		srcFP := fixedPointForSideV1EngineSceneConnection(srcSide)
		if hasSrcAnchor {
			srcFP = fixedPointForAnchorV1EngineSceneConnection(srcAnchor)
		}
		dstFP := fixedPointForSideV1EngineSceneConnection(dstSide)
		if hasDstAnchor {
			dstFP = fixedPointForAnchorV1EngineSceneConnection(dstAnchor)
		}
		if fp, ok := umlSequenceFixedPointV1EngineSceneConnectionRoute(conn, "src", srcSide); ok {
			srcFP = fp
		}
		if fp, ok := umlSequenceFixedPointV1EngineSceneConnectionRoute(conn, "dst", dstSide); ok {
			dstFP = fp
		}
		srcEdge := rectFixedPointV1EngineSceneConnection(srcRect, srcFP)
		dstEdge := rectFixedPointV1EngineSceneConnection(dstRect, dstFP)
		dx := dstEdge[0] - srcEdge[0]
		dy := dstEdge[1] - srcEdge[1]
		style := resolveConnectionStyleV1EngineSceneConnectionRoute(conn)
		if conn.Attrs[internalConnectionCrossFrameAttrV1EngineParseDocument] == "true" {
			srcFrameID := strings.TrimSpace(conn.Attrs[internalConnectionSrcFrameAttrV1EngineParseDocument])
			dstFrameID := strings.TrimSpace(conn.Attrs[internalConnectionDstFrameAttrV1EngineParseDocument])
			srcFrameRect, srcFrameOK := frameRects[srcFrameID]
			dstFrameRect, dstFrameOK := frameRects[dstFrameID]
			if srcFrameOK && dstFrameOK {
				seed := stableConnectionSeedV1EngineSceneConnectionRoute(srcKey, dstKey, i)
				renderCrossFrameConnectionV1EngineSceneConnectionPage(conn, srcKey, dstKey, srcElemID, dstElemID, srcRect, dstRect, srcFP, dstFP, srcFrameID, dstFrameID, srcFrameRect, dstFrameRect, style, seed, i, updated, elements, boundMap)
				continue
			}
		}
		routePoints := excalidrawConnectionPointsV1EngineSceneConnectionRoute(conn, srcRect, dstRect, srcSide, dstSide, style.Kind, obstacles, placed, routePaths)
		if style.Kind == "route" {
			routePaths[routePairKeyV1EngineRouteBuild(excalidrawRouteRequestV1EngineSceneConnectionRoute(conn, srcRect, dstRect, srcSide, dstSide, style.Kind), false)] = append([]ptV1EngineRouteTypes(nil), routePoints...)
		}
		placed = append(placed, toSegmentsV1EngineRouteGeometry(routePoints))
		minX, minY, maxX, maxY := srcEdge[0], srcEdge[1], srcEdge[0], srcEdge[1]
		points := make([][]float64, 0, len(routePoints))
		for _, p := range routePoints {
			minX = math.Min(minX, p.X)
			minY = math.Min(minY, p.Y)
			maxX = math.Max(maxX, p.X)
			maxY = math.Max(maxY, p.Y)
			points = append(points, []float64{p.X - srcEdge[0], p.Y - srcEdge[1]})
		}
		// seed は src/dst/index から決定論的に計算し、再生成しても描画ばらつきが出ないようにする。
		seed := stableConnectionSeedV1EngineSceneConnectionRoute(srcKey, dstKey, i)
		connID := fmt.Sprintf("conn-%s-%s-%d", sanitizeElementIDV1EngineSceneConnectionRoute(srcKey), sanitizeElementIDV1EngineSceneConnectionRoute(dstKey), i)

		if style.Kind == "route" {
			for _, endpoint := range []struct {
				id   string
				edge [2]float64
				side string
				seed int
			}{{srcElemID, srcEdge, srcSide, seed}, {dstElemID, dstEdge, dstSide, seed + 1}} {
				key := endpoint.id + "|" + endpoint.side
				candidate := junctionCandidates[key]
				if candidate == nil {
					candidate = &junctionCandidate{edge: endpoint.edge, side: endpoint.side, color: style.Color, seed: endpoint.seed}
					junctionCandidates[key] = candidate
				}
				candidate.count++
			}
		}

		customData := map[string]any{
			"xaligoConnectorKind":             style.Kind,
			"xaligoConnectorStartArrowhead":   style.StartArrowhead,
			"xaligoConnectorEndArrowhead":     style.EndArrowhead,
			"xaligoConnectorStyleSourceKnown": true,
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
		if hasSrcAnchor {
			customData["xaligoConnectorSrcAnchor"] = true
		}
		if hasDstAnchor {
			customData["xaligoConnectorDstAnchor"] = true
		}
		if bends := strings.TrimSpace(connectionBendsV1EngineSceneConnectionRoute(conn)); bends != "" {
			customData["xaligoConnectorBends"] = bends
		}
		if scale, ok := positiveFloatAttrV1EngineSceneConnectionRoute(conn, "coordinate-scale", "scale"); ok {
			customData["xaligoConnectorScale"] = scale
		}
		if grid, ok := positiveFloatAttrV1EngineSceneConnectionRoute(conn, "grid"); ok {
			customData["xaligoConnectorGrid"] = grid
		}
		if hasSrcAnchor {
			customData["xaligoConnectorStartAnchor"] = true
		}
		if hasDstAnchor {
			customData["xaligoConnectorEndAnchor"] = true
		}
		applyDatabaseConnectionMetadataV1EngineSceneConnectionRender(customData, conn)
		applyUMLConnectionMetadataV1EngineSceneConnectionRender(customData, conn)
		applyConnectionDiffStatusV1EngineSceneDiffHighlight(customData, conn)

		*elements = append(*elements, map[string]any{
			"id": connID, "type": "arrow",
			"x": srcEdge[0], "y": srcEdge[1],
			"width": math.Max(math.Abs(dx), maxX-minX), "height": math.Max(math.Abs(dy), maxY-minY),
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
			"startBinding": map[string]any{
				"elementId":  srcElemID,
				"focus":      0.0,
				"gap":        5.0,
				"fixedPoint": []float64{srcFP[0], srcFP[1]},
			},
			"endBinding": map[string]any{
				"elementId":  dstElemID,
				"focus":      0.0,
				"gap":        5.0,
				"fixedPoint": []float64{dstFP[0], dstFP[1]},
			},
			"startArrowhead":     style.ExcalidrawStartArrowhead,
			"endArrowhead":       style.ExcalidrawEndArrowhead,
			"endArrowheadSize":   "s",
			"startArrowheadSize": "s",
			"elbowed":            true,
			"customData":         customData,
		})
		appendUMLRelationLabelV1EngineSceneConnectionRender(elements, conn, connID, routePoints, style.Color, updated, seed)

		// Register this arrow in boundMap for both endpoints.
		entry := map[string]any{"type": "arrow", "id": connID}
		boundMap[srcElemID] = append(boundMap[srcElemID], entry)
		boundMap[dstElemID] = append(boundMap[dstElemID], entry)
	}

	junctionKeys := make([]string, 0, len(junctionCandidates))
	for key, candidate := range junctionCandidates {
		if candidate.count >= 2 {
			junctionKeys = append(junctionKeys, key)
		}
	}
	sort.Strings(junctionKeys)
	for i, key := range junctionKeys {
		candidate := junctionCandidates[key]
		point := extendConnectionPointV1EngineSceneConnectionRoute(candidate.edge, candidate.side, 25)
		const diameter = 8.0
		*elements = append(*elements, map[string]any{
			"id": fmt.Sprintf("junction-%d", i), "type": "ellipse",
			"x": point[0] - diameter/2, "y": point[1] - diameter/2,
			"width": diameter, "height": diameter, "angle": 0,
			"strokeColor": candidate.color, "backgroundColor": candidate.color,
			"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100, "groupIds": []string{}, "roundness": nil,
			"seed": candidate.seed, "version": 1, "versionNonce": candidate.seed,
			"isDeleted": false, "boundElements": nil, "updated": updated,
			"link": nil, "locked": false, "frameId": nil,
			"customData": map[string]any{"xaligoJunction": true},
		})
	}

	// Second pass: write back boundElements into each referenced element so that
	// Excalidraw recognises the binding relationship.
	if len(boundMap) == 0 {
		return
	}
	for idx := range *elements {
		elem := (*elements)[idx]
		id, _ := elem["id"].(string)
		if entries, ok := boundMap[id]; ok {
			// Merge with any existing bound elements (e.g. text containerId refs).
			existing, _ := elem["boundElements"].([]map[string]any)
			elem["boundElements"] = append(existing, entries...)
			(*elements)[idx] = elem
		}
	}
}

func applyUMLConnectionMetadataV1EngineSceneConnectionRender(customData map[string]any, conn *entity.Node) {
	if customData == nil || conn == nil {
		return
	}
	for source, target := range map[string]string{
		"uml-id":               "xaligoUmlId",
		"uml-diagram-kind":     "xaligoUmlDiagramKind",
		"uml-relation-kind":    "xaligoUmlRelationKind",
		"uml-relation-label":   "xaligoUmlRelationLabel",
		"uml-src-ref":          "xaligoUmlRelationSourceReference",
		"uml-dst-ref":          "xaligoUmlRelationDestinationReference",
		"uml-order":            "xaligoUmlMessageOrder",
		"uml-guard":            "xaligoUmlGuard",
		"uml-src-multiplicity": "xaligoUmlSourceMultiplicity",
		"uml-dst-multiplicity": "xaligoUmlDestinationMultiplicity",
		"uml-at":               "xaligoUmlOccurrenceAt",
		"uml-from":             "xaligoUmlDurationFrom",
		"uml-to":               "xaligoUmlDurationTo",
	} {
		if value := strings.TrimSpace(conn.Attr(source)); value != "" {
			customData[target] = value
		}
	}
}

func appendUMLRelationLabelV1EngineSceneConnectionRender(elements *[]map[string]any, conn *entity.Node, connID string, routePoints []ptV1EngineRouteTypes, color string, updated int64, seed int) {
	label := strings.TrimSpace(conn.Attr("uml-relation-label"))
	if label == "" || len(routePoints) == 0 {
		return
	}
	point := umlRelationLabelPointV1EngineSceneConnectionRender(routePoints, seed)
	width := math.Max(80, math.Min(220, textWidthV1EngineSceneItem(label, 6)+16))
	const height = 20.0
	labelID := connID + "-uml-label"
	customData := map[string]any{}
	applyUMLConnectionMetadataV1EngineSceneConnectionRender(customData, conn)
	customData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(entity.TextRoleConnectorLabel, true, 1.2)
	*elements = append(*elements, map[string]any{
		"id": labelID, "type": "text", "x": point.X - width/2, "y": point.Y - height/2,
		"width": width, "height": height, "angle": 0,
		"strokeColor": color, "backgroundColor": "#ffffff",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100, "groupIds": []string{}, "roundness": nil,
		"seed": seed + 1000, "version": 1, "versionNonce": seed + 1000,
		"isDeleted": false, "boundElements": nil, "updated": updated, "link": nil, "locked": false,
		"text": label, "fontSize": 12.0, "fontFamily": 2,
		"textAlign": "center", "verticalAlign": "middle", "containerId": nil,
		"originalText": label, "lineHeight": 1.2,
		"customData": customData,
	})
}

func umlRelationLabelPointV1EngineSceneConnectionRender(routePoints []ptV1EngineRouteTypes, seed int) ptV1EngineRouteTypes {
	if len(routePoints) < 2 {
		return routePoints[0]
	}
	bestStart, bestEnd := routePoints[0], routePoints[1]
	bestLength := math.Hypot(bestEnd.X-bestStart.X, bestEnd.Y-bestStart.Y)
	for index := 1; index < len(routePoints)-1; index++ {
		start, end := routePoints[index], routePoints[index+1]
		length := math.Hypot(end.X-start.X, end.Y-start.Y)
		if length > bestLength {
			bestStart, bestEnd, bestLength = start, end, length
		}
	}
	point := ptV1EngineRouteTypes{X: (bestStart.X + bestEnd.X) / 2, Y: (bestStart.Y + bestEnd.Y) / 2}
	offset := 12.0 + float64(seed%3)*6
	if seed%2 != 0 {
		offset = -offset
	}
	if math.Abs(bestEnd.X-bestStart.X) >= math.Abs(bestEnd.Y-bestStart.Y) {
		point.Y += offset
	} else {
		point.X += offset
	}
	return point
}

func applyDatabaseConnectionMetadataV1EngineSceneConnectionRender(customData map[string]any, conn *entity.Node) {
	if customData == nil || conn == nil {
		return
	}
	for source, target := range map[string]string{
		"_xaligoDatabaseForeignKey": "xaligoDatabaseForeignKey",
		"_xaligoDatabaseReferences": "xaligoDatabaseReferences",
		"_xaligoDatabaseOnDelete":   "xaligoDatabaseOnDelete",
		"_xaligoDatabaseOnUpdate":   "xaligoDatabaseOnUpdate",
	} {
		if value := strings.TrimSpace(conn.Attr(source)); value != "" {
			customData[target] = value
		}
	}
}
