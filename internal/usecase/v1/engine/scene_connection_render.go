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
func renderConnectionsV1EngineSceneConnectionRender(connections []*entity.Node, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, frameRects map[string][4]float64, frameElementIDs map[string]string, frameMetadata map[string]frameMetadataSceneGeometryV1EngineSceneFrameMetadata, elements *[]map[string]any) {
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
	hardObstacles := frameMetadataReservedObstaclesV1EngineSceneConnectionRender(frameMetadata)
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
		srcSideExplicit := hasSrcAnchor
		dstSideExplicit := hasDstAnchor
		if hasSrcAnchor {
			srcSide = string(srcAnchor.side)
		} else if explicit, ok := connectionEndpointSideV1EngineSceneConnectionRoute(conn, "src"); ok {
			srcSide = string(explicit)
			srcSideExplicit = true
		} else if len(bends) > 0 {
			srcSide = sideTowardPointV1EngineSceneConnection(srcImgRect, bends[0])
		}
		if hasDstAnchor {
			dstSide = string(dstAnchor.side)
		} else if explicit, ok := connectionEndpointSideV1EngineSceneConnectionRoute(conn, "dst"); ok {
			dstSide = string(explicit)
			dstSideExplicit = true
		} else if len(bends) > 0 {
			dstSide = sideTowardPointV1EngineSceneConnection(dstImgRect, bends[len(bends)-1])
		}
		if _, ok := umlSequencePositionV1EngineSceneConnectionRoute(conn, "src"); ok {
			srcSide = umlSequenceVerticalSideV1EngineSceneConnectionRoute(srcSide)
			dstSide = umlSequenceVerticalSideV1EngineSceneConnectionRoute(dstSide)
		}
		srcFrameSide, dstFrameSide := srcSide, dstSide

		crossFrame := conn.Attrs[internalConnectionCrossFrameAttrV1EngineParseDocument] == "true"
		srcFrameID := strings.TrimSpace(conn.Attrs[internalConnectionSrcFrameAttrV1EngineParseDocument])
		dstFrameID := strings.TrimSpace(conn.Attrs[internalConnectionDstFrameAttrV1EngineParseDocument])
		var srcFrameRect, dstFrameRect [4]float64
		var srcVisualRect, dstVisualRect [4]float64
		if crossFrame {
			var srcFrameOK, dstFrameOK bool
			srcFrameRect, srcFrameOK = frameRects[srcFrameID]
			dstFrameRect, dstFrameOK = frameRects[dstFrameID]
			if !srcFrameOK || !dstFrameOK {
				loggerV1EngineSharedLogging.WARN(IUESRC004V1EngineSceneTypes, "cross-frame connection frame not found or not rendered", map[string]any{"srcFrame": srcFrameID, "dstFrame": dstFrameID, "srcFrameRendered": srcFrameOK, "dstFrameRendered": dstFrameOK})
				continue
			}
			// Page-link stubs select their endpoint and page side from the local
			// endpoint geometry. Manual bends remain logical metadata for graph
			// adapters and do not steer either page-local stub.
			srcVisualRect = endpointVisualRectV1EngineSceneConnectionPage(srcImgRect, itemLblRects[srcKey])
			dstVisualRect = endpointVisualRectV1EngineSceneConnectionPage(dstImgRect, itemLblRects[dstKey])
			srcFrameAnchor, srcFrameSideExplicit := connectionFrameAnchorV1EngineSceneConnectionRoute(conn, "src")
			dstFrameAnchor, dstFrameSideExplicit := connectionFrameAnchorV1EngineSceneConnectionRoute(conn, "dst")
			if !srcSideExplicit {
				srcSide = nearestFrameSideV1EngineSceneConnectionPage(srcFrameRect, srcVisualRect, dstFrameRect)
			}
			if !dstSideExplicit {
				dstSide = nearestFrameSideV1EngineSceneConnectionPage(dstFrameRect, dstVisualRect, srcFrameRect)
			}
			if srcFrameSideExplicit {
				srcFrameSide = string(srcFrameAnchor.side)
			} else {
				srcFrameSide = srcSide
			}
			if dstFrameSideExplicit {
				dstFrameSide = string(dstFrameAnchor.side)
			} else {
				dstFrameSide = dstSide
			}
			if adjustedEndpointSide := pageLinkSideAvoidingFrameMetadataV1EngineSceneConnectionPage(srcSide, srcFrameRect, srcVisualRect, dstFrameRect, frameMetadata[srcFrameID]); adjustedEndpointSide != srcSide {
				srcSide = adjustedEndpointSide
				hasSrcAnchor = false
			}
			if adjustedEndpointSide := pageLinkSideAvoidingFrameMetadataV1EngineSceneConnectionPage(dstSide, dstFrameRect, dstVisualRect, srcFrameRect, frameMetadata[dstFrameID]); adjustedEndpointSide != dstSide {
				dstSide = adjustedEndpointSide
				hasDstAnchor = false
			}
			if adjustedFrameSide := pageLinkSideAvoidingFrameMetadataV1EngineSceneConnectionPage(srcFrameSide, srcFrameRect, srcVisualRect, dstFrameRect, frameMetadata[srcFrameID]); adjustedFrameSide != srcFrameSide {
				srcFrameSide = adjustedFrameSide
				if !srcFrameSideExplicit {
					srcSide = adjustedFrameSide
					hasSrcAnchor = false
				}
			}
			if adjustedFrameSide := pageLinkSideAvoidingFrameMetadataV1EngineSceneConnectionPage(dstFrameSide, dstFrameRect, dstVisualRect, srcFrameRect, frameMetadata[dstFrameID]); adjustedFrameSide != dstFrameSide {
				dstFrameSide = adjustedFrameSide
				if !dstFrameSideExplicit {
					dstSide = adjustedFrameSide
					hasDstAnchor = false
				}
			}
		} else {
			srcLocalVisualRect := endpointVisualRectV1EngineSceneConnectionPage(srcImgRect, itemLblRects[srcKey])
			dstLocalVisualRect := endpointVisualRectV1EngineSceneConnectionPage(dstImgRect, itemLblRects[dstKey])
			if adjusted := localConnectionSideAvoidingFrameMetadataV1EngineSceneConnectionRender(srcSide, srcLocalVisualRect, dstLocalVisualRect, frameRects[srcFrameID], frameMetadata[srcFrameID]); adjusted != srcSide {
				srcSide = adjusted
				hasSrcAnchor = false
			}
			if adjusted := localConnectionSideAvoidingFrameMetadataV1EngineSceneConnectionRender(dstSide, dstLocalVisualRect, srcLocalVisualRect, frameRects[dstFrameID], frameMetadata[dstFrameID]); adjusted != dstSide {
				dstSide = adjusted
				hasDstAnchor = false
			}
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
		style := resolveConnectionStyleV1EngineSceneConnectionRoute(conn)
		if crossFrame {
			seed := stableConnectionSeedV1EngineSceneConnectionRoute(srcKey, dstKey, i)
			renderCrossFrameConnectionV1EngineSceneConnectionPage(conn, srcKey, dstKey, srcElemID, dstElemID, srcRect, dstRect, srcVisualRect, dstVisualRect, srcFP, dstFP, srcSide, dstSide, srcFrameSide, dstFrameSide, srcFrameID, dstFrameID, srcFrameRect, dstFrameRect, srcElemID == frameElementIDs[srcFrameID], dstElemID == frameElementIDs[dstFrameID], frameMetadata[srcFrameID], frameMetadata[dstFrameID], style, seed, i, updated, elements, boundMap)
			continue
		}
		srcEdge := rectFixedPointV1EngineSceneConnection(srcRect, srcFP)
		dstEdge := rectFixedPointV1EngineSceneConnection(dstRect, dstFP)
		dx := dstEdge[0] - srcEdge[0]
		dy := dstEdge[1] - srcEdge[1]
		routePoints := excalidrawConnectionPointsV1EngineSceneConnectionRoute(conn, srcRect, dstRect, srcSide, dstSide, style.Kind, obstacles, hardObstacles, placed, routePaths)
		if len(routePoints) < 2 {
			continue
		}
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
		if srcFrameID != "" && srcFrameID == dstFrameID {
			customData["xaligoFrameID"] = srcFrameID
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
		appendUMLRelationLabelV1EngineSceneConnectionRender(elements, conn, connID, routePoints, style.Color, updated, seed, frameRects[srcFrameID], frameMetadata[srcFrameID])

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

func frameMetadataReservedObstaclesV1EngineSceneConnectionRender(metadata map[string]frameMetadataSceneGeometryV1EngineSceneFrameMetadata) []rectV1EngineRouteTypes {
	keys := make([]string, 0, len(metadata))
	for key := range metadata {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	obstacles := make([]rectV1EngineRouteTypes, 0, len(keys))
	for _, key := range keys {
		reserved := metadata[key].Reserved
		if reserved[2] > 0 && reserved[3] > 0 {
			obstacles = append(obstacles, rectV1EngineRouteTypes{X: reserved[0], Y: reserved[1], W: reserved[2], H: reserved[3]})
		}
	}
	return obstacles
}

func localConnectionSideAvoidingFrameMetadataV1EngineSceneConnectionRender(side string, endpointRect, otherRect, frameRect [4]float64, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) string {
	reserved := metadata.Reserved
	if reserved[2] <= 0 || reserved[3] <= 0 || side != metadata.Position || frameRect[2] <= 0 || frameRect[3] <= 0 {
		return side
	}
	options := defaultRouterOptionsV1EngineRouteTypes()
	minimumClearance := 5.0 + math.Max(options.Stub, options.LaneGap) + 1.0
	clearance := math.Inf(1)
	if side == "top" {
		clearance = endpointRect[1] - (reserved[1] + reserved[3])
	} else if side == "bottom" {
		clearance = reserved[1] - (endpointRect[1] + endpointRect[3])
	}
	if clearance >= minimumClearance {
		return side
	}
	return nearestFrameSideExcludingV1EngineSceneConnectionPage(frameRect, endpointRect, otherRect, metadata.Position)
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

func appendUMLRelationLabelV1EngineSceneConnectionRender(elements *[]map[string]any, conn *entity.Node, connID string, routePoints []ptV1EngineRouteTypes, color string, updated int64, seed int, frameRect [4]float64, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) {
	label := strings.TrimSpace(conn.Attr("uml-relation-label"))
	if label == "" || len(routePoints) == 0 {
		return
	}
	point := umlRelationLabelPointV1EngineSceneConnectionRender(routePoints, seed)
	fontSize := 12.0
	height := 20.0
	if safeTop, safeBottom, ok := frameMetadataSafeVerticalIntervalV1EngineSceneConnectionPage(frameRect, frameMetadataReservedRectsV1EngineSceneConnectionPage(metadata)); ok {
		height = math.Min(height, safeBottom-safeTop)
		fontSize = math.Min(fontSize, height/1.2)
	}
	width := math.Max(80, math.Min(220, textWidthV1EngineSceneItem(label, fontSize*0.5)+16))
	if frameRect[2] > 0 {
		width = math.Min(width, frameRect[2])
	}
	x, y := point.X-width/2, point.Y-height/2
	x, y = frameMetadataLabelPositionV1EngineSceneConnectionRender(x, y, width, height, frameRect, metadata)
	labelID := connID + "-uml-label"
	customData := map[string]any{}
	applyUMLConnectionMetadataV1EngineSceneConnectionRender(customData, conn)
	customData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(entity.TextRoleConnectorLabel, true, 1.2)
	*elements = append(*elements, map[string]any{
		"id": labelID, "type": "text", "x": x, "y": y,
		"width": width, "height": height, "angle": 0,
		"strokeColor": color, "backgroundColor": "#ffffff",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100, "groupIds": []string{}, "roundness": nil,
		"seed": seed + 1000, "version": 1, "versionNonce": seed + 1000,
		"isDeleted": false, "boundElements": nil, "updated": updated, "link": nil, "locked": false,
		"text": label, "fontSize": fontSize, "fontFamily": 2,
		"textAlign": "center", "verticalAlign": "middle", "containerId": nil,
		"originalText": label, "lineHeight": 1.2,
		"customData": customData,
	})
}

func frameMetadataLabelPositionV1EngineSceneConnectionRender(x, y, width, height float64, frameRect [4]float64, metadata frameMetadataSceneGeometryV1EngineSceneFrameMetadata) (float64, float64) {
	reserved := metadata.Reserved
	if reserved[2] <= 0 || reserved[3] <= 0 {
		return x, y
	}
	const gap = 6.0
	if frameRect[2] <= 0 || frameRect[3] <= 0 {
		return x, y
	}
	horizontalInset := math.Min(gap, math.Max(0, (frameRect[2]-width)/2))
	x = clampFloatV1EngineLayoutPort(x, frameRect[0]+horizontalInset, math.Max(frameRect[0]+horizontalInset, frameRect[0]+frameRect[2]-width-horizontalInset))
	if safeTop, safeBottom, ok := frameMetadataSafeVerticalIntervalV1EngineSceneConnectionPage(frameRect, [][4]float64{reserved}); ok {
		verticalInset := math.Min(gap, math.Max(0, (safeBottom-safeTop-height)/2))
		y = clampFloatV1EngineLayoutPort(y, safeTop+verticalInset, math.Max(safeTop+verticalInset, safeBottom-height-verticalInset))
	}
	return x, y
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
