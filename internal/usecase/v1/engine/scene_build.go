package engine

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

// BuildJSONWithFSV1EngineSceneBuild is a convenience wrapper for WASM / embedded builds.
// It uses fsys (typically an embed.FS) for all asset reads instead of the OS
// filesystem.  catalogCSV and svgGroupDir are resolved relative to the root
// of fsys (e.g. "service-catalog.csv" and "svg/Architecture-Group-Icons").
func BuildJSONWithFSV1EngineSceneBuild(root *entity.Box, fsys fs.FS, catalogCSV, svgGroupDir string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, deps SceneDependenciesV1EngineSceneTypes) ([]byte, error) {
	return BuildJSONV1EngineSceneBuild(root, svgGroupDir, catalogCSV, "", itemIconSize, connections, abbrevMap, fsys, deps)
}

// BuildJSONV1EngineSceneBuild converts a entity.Box layout tree into Excalidraw JSON.
// svgGroupDir:  absolute path to Architecture-Group-Icons/ (or FS-relative path when fsys≠nil)
// catalogCSV:   absolute path to service-catalog.csv (or FS-relative path when fsys≠nil)
// projectRoot:  project root directory (used to resolve rel_path from catalog; ignored when fsys≠nil)
// itemIconSize: default maximum icon size (px) for <item> elements.
// connections:  <connection> nodes extracted from the DSL (may be nil).
// abbrevMap:    optional catalog-ID → abbreviation map derived from services.csv.
// fsys:         when non-nil, all asset reads go through this fs.FS (WASM / embedded mode).
func BuildJSONV1EngineSceneBuild(root *entity.Box, svgGroupDir string, catalogCSV string, projectRoot string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, fsys fs.FS, deps SceneDependenciesV1EngineSceneTypes) ([]byte, error) {
	if root == nil {
		return nil, fmt.Errorf("root layout is nil")
	}
	updated := excalidrawUpdatedV1EngineSceneTypes
	frameSeed := stableSceneSeedV1EngineSceneTypes("paper-frame")

	// Outermost Excalidraw frame element representing the paper size.
	frameElem := map[string]any{
		"id": "paper-frame", "type": "frame",
		"x": root.X, "y": root.Y, "width": root.W, "height": root.H,
		"angle":       0,
		"name":        detectPaperNameV1EngineSceneTypes(root.W, root.H),
		"strokeColor": "transparent", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": frameSeed, "version": 1,
		"versionNonce": frameSeed,
		"isDeleted":    false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
	}

	var elements []map[string]any
	elements = append(elements, frameElem)
	files := map[string]any{}

	// 2パス: 1) item を visibleAncestorID ごとに収集, 2) グリッド一括描画
	itemGroups := map[string][]*entity.Box{}
	ancestorBoxes := map[string]*entity.Box{}
	itemFrames := map[string]*entity.Box{}
	// <frame item-size="N"> overrides the global itemIconSize.
	if v := root.Attrs["item-size"]; v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil && f > 0 {
			itemIconSize = f
		}
	}

	// itemImgRects / itemLblRects / itemImgIDs / itemLblIDs:
	// connection key → bounding rect (x,y,w,h) and element ID of the image / label elements.
	// Populated during renderItemGrid → renderIconAt, used for edge-based connections.
	itemImgRects := map[string][4]float64{}
	itemLblRects := map[string][4]float64{}
	itemImgIDs := map[string]string{}
	itemLblIDs := map[string]string{}
	frameRects := map[string][4]float64{}
	frameElementIDs := map[string]string{}
	if root.Tag == "frames" {
		for _, child := range root.Children {
			if child.Tag != "frame" {
				continue
			}
			frameID := strings.TrimSpace(child.Attrs["id"])
			if frameID == "" {
				frameID = child.ID
			}
			pageFrameID := pageFrameElementIDV1EngineSceneWalk(child)
			pageFrameSeed := stableSceneSeedV1EngineSceneTypes(pageFrameID)
			pageFrame := map[string]any{
				"id": pageFrameID, "type": "frame",
				"x": child.X, "y": child.Y, "width": child.W, "height": child.H,
				"angle":       0,
				"name":        frameID + " (" + detectPaperNameV1EngineSceneTypes(child.W, child.H) + ")",
				"strokeColor": "transparent", "backgroundColor": "transparent",
				"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": pageFrameSeed, "version": 1,
				"versionNonce": pageFrameSeed,
				"isDeleted":    false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"customData": map[string]any{"xaligoPageFrame": true, "xaligoFrameID": frameID},
			}
			elements = append(elements, pageFrame)
			frameRects[frameID] = [4]float64{child.X, child.Y, child.W, child.H}
			frameElementIDs[frameID] = pageFrameID
			registerConnectionEndpointV1EngineSceneWalk(child, pageFrameID, frameRects[frameID], itemImgRects, itemImgIDs)
		}
	} else if root.Tag == "frame" {
		frameID := strings.TrimSpace(root.Attrs["id"])
		frameRects[frameID] = [4]float64{root.X, root.Y, root.W, root.H}
		frameElementIDs[frameID] = "paper-frame"
		registerConnectionEndpointV1EngineSceneWalk(root, "paper-frame", frameRects[frameID], itemImgRects, itemImgIDs)
	}

	walkV1EngineSceneWalk(root, &elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, root, itemGroups, ancestorBoxes, itemFrames, itemImgRects, itemImgIDs, nil, deps)
	ancestorIDs := make([]string, 0, len(itemGroups))
	for ancID := range itemGroups {
		ancestorIDs = append(ancestorIDs, ancID)
	}
	sort.Strings(ancestorIDs)
	for _, ancID := range ancestorIDs {
		items := itemGroups[ancID]
		if err := renderItemGridV1EngineSceneItem(items, ancestorBoxes[ancID], itemFrames[ancID], &elements, files, catalogCSV, projectRoot, fsys, itemIconSize, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps); err != nil {
			return nil, err
		}
	}
	frameMetadata := appendFrameMetadataV1EngineSceneFrameMetadata(root, &elements)
	applySemanticElementMetadataV1EngineSceneWalk(elements, collectSemanticElementMetadataV1EngineSceneWalk(root))
	componentInterfaces := registerUMLComponentInterfaceEndpointsV1EngineSceneBuild(elements, itemImgRects, itemImgIDs)
	expandUMLComponentInterfaceEndpointsV1EngineSceneBuild(&elements, &componentInterfaces, itemImgRects, itemImgIDs, connections)
	connectionsForRender := bindUMLComponentInterfaceConnectionsV1EngineSceneBuild(connections, componentInterfaces)
	renderConnectionsV1EngineSceneConnectionRender(connectionsForRender, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, frameRects, frameElementIDs, frameMetadata, &elements)
	appendUMLComponentCallerSocketsV1EngineSceneBuild(&elements)
	appendDiffBoxHighlightsV1EngineSceneDiffHighlight(root, &elements)
	elements = orderSceneLayersV1EngineSceneBuild(elements)

	out := fileV1EngineSceneTypes{
		Type:     "excalidraw",
		Version:  2,
		Source:   "https://github.com/xaligo/xaligo",
		Elements: elements,
		AppState: map[string]any{
			"gridSize":            20,
			"viewBackgroundColor": "#ffffff",
		},
		Files: files,
	}
	return json.MarshalIndent(out, "", "  ")
}

type umlComponentInterfaceEndpointsV1EngineSceneBuild struct {
	byOwner        map[string]map[string]*umlComponentInterfaceEndpointGroupV1EngineSceneBuild
	byKey          map[string]*umlComponentInterfaceEndpointGroupV1EngineSceneBuild
	sideByEndpoint map[string]string
}

type umlComponentInterfaceEndpointGroupV1EngineSceneBuild struct {
	ownerKey     string
	label        string
	baseKey      string
	baseID       string
	baseRect     [4]float64
	ownerRect    [4]float64
	portRect     [4]float64
	endpointKeys []string
}

func registerUMLComponentInterfaceEndpointsV1EngineSceneBuild(elements []map[string]any, endpointRects map[string][4]float64, endpointIDs map[string]string) umlComponentInterfaceEndpointsV1EngineSceneBuild {
	result := umlComponentInterfaceEndpointsV1EngineSceneBuild{byOwner: map[string]map[string]*umlComponentInterfaceEndpointGroupV1EngineSceneBuild{}, byKey: map[string]*umlComponentInterfaceEndpointGroupV1EngineSceneBuild{}, sideByEndpoint: map[string]string{}}
	portRects := map[string][4]float64{}
	for _, element := range elements {
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlComponentInterfacePort"] != true {
			continue
		}
		ownerKey, _ := customData["xaligoUmlComponentOwnerConnectionKey"].(string)
		label, _ := customData["xaligoUmlComponentInterfaceLabel"].(string)
		x, okX := element["x"].(float64)
		y, okY := element["y"].(float64)
		w, okW := element["width"].(float64)
		h, okH := element["height"].(float64)
		if ownerKey == "" || label == "" || !okX || !okY || !okW || !okH || w <= 0 || h <= 0 {
			continue
		}
		portRects[umlComponentInterfaceEndpointKeyV1EngineSceneBuild(ownerKey, label)] = [4]float64{x, y, w, h}
	}
	for _, element := range elements {
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlComponentInterfaceSymbol"] != true {
			continue
		}
		ownerKey, _ := customData["xaligoUmlComponentOwnerConnectionKey"].(string)
		label, _ := customData["xaligoUmlComponentInterfaceLabel"].(string)
		id, _ := element["id"].(string)
		if ownerKey == "" || label == "" || id == "" {
			continue
		}
		x, okX := element["x"].(float64)
		y, okY := element["y"].(float64)
		w, okW := element["width"].(float64)
		h, okH := element["height"].(float64)
		if !okX || !okY || !okW || !okH || w <= 0 || h <= 0 {
			continue
		}
		endpointKey := umlComponentInterfaceEndpointKeyV1EngineSceneBuild(ownerKey, label)
		endpointRects[endpointKey] = [4]float64{x, y, w, h}
		endpointIDs[endpointKey] = id
		result.sideByEndpoint[endpointKey] = umlComponentInterfaceConnectionSideV1EngineSceneBuild([4]float64{x, y, w, h}, endpointRects[ownerKey])
		if result.byOwner[ownerKey] == nil {
			result.byOwner[ownerKey] = map[string]*umlComponentInterfaceEndpointGroupV1EngineSceneBuild{}
		}
		group := &umlComponentInterfaceEndpointGroupV1EngineSceneBuild{ownerKey: ownerKey, label: label, baseKey: endpointKey, baseID: id, baseRect: [4]float64{x, y, w, h}, ownerRect: endpointRects[ownerKey], portRect: portRects[endpointKey], endpointKeys: []string{endpointKey}}
		result.byOwner[ownerKey][label] = group
		result.byKey[endpointKey] = group
	}
	return result
}

func expandUMLComponentInterfaceEndpointsV1EngineSceneBuild(elements *[]map[string]any, endpoints *umlComponentInterfaceEndpointsV1EngineSceneBuild, endpointRects map[string][4]float64, endpointIDs map[string]string, connections []*entity.Node) {
	if elements == nil || endpoints == nil || len(connections) == 0 || len(endpoints.byOwner) == 0 {
		return
	}
	counts := map[string]int{}
	for _, conn := range connections {
		if conn == nil || conn.Attr("uml-diagram-kind") != "component-diagram" || conn.Attr("uml-relation-kind") != "association" || conn.Attr("uml-src-kind") != "component" || conn.Attr("uml-dst-kind") != "component" {
			continue
		}
		srcKey := strings.TrimSpace(conn.Attrs[internalConnectionSrcKeyAttrV1EngineParseDocument])
		dstKey := strings.TrimSpace(conn.Attrs[internalConnectionDstKeyAttrV1EngineParseDocument])
		boundDst, ok := matchingUMLComponentInterfaceEndpointKeyV1EngineSceneBuild(*endpoints, srcKey, dstKey)
		if ok {
			counts[boundDst]++
		}
	}
	updated := excalidrawUpdatedV1EngineSceneTypes
	for baseKey, count := range counts {
		if count <= 1 {
			continue
		}
		group := endpoints.byKey[baseKey]
		if group == nil {
			continue
		}
		rect := group.baseRect
		port := group.portRect
		if port[2] <= 0 || port[3] <= 0 {
			continue
		}
		diameter := rect[2]
		step := math.Max(umlComponentCallerSocketRadiusForCircleV1EngineSceneBuild(rect)*2+2, 16)
		baseCenterY := rect[1] + rect[3]/2
		startY := baseCenterY - step*float64(count-1)/2
		circleLeft := rect[0]
		circleRight := circleLeft + diameter
		trunkX := circleRight + math.Max(1.5, (port[0]-circleRight)/2)
		for index := 0; index < count; index++ {
			cy := startY + step*float64(index)
			endpointKey := baseKey
			circleID := group.baseID
			if index == 0 {
				moveUMLComponentInterfaceElementV1EngineSceneBuild(*elements, circleID, circleLeft, cy-diameter/2)
			} else {
				endpointKey = fmt.Sprintf("%s|connection-%d", baseKey, index)
				circleID = fmt.Sprintf("%s-connection-%d", group.baseID, index)
				appendUMLComponentInterfaceCircleV1EngineSceneBuild(elements, circleID, circleLeft, cy-diameter/2, diameter, updated, group)
				group.endpointKeys = append(group.endpointKeys, endpointKey)
			}
			endpointRects[endpointKey] = [4]float64{circleLeft, cy - diameter/2, diameter, diameter}
			endpointIDs[endpointKey] = circleID
			endpoints.sideByEndpoint[endpointKey] = "left"
		}
		removeUMLComponentInterfaceStemsV1EngineSceneBuild(elements, group)
		appendUMLComponentInterfaceLineV1EngineSceneWalk(elements, group.baseID+"-multi-trunk", trunkX, startY, trunkX, startY+step*float64(count-1), updated, map[string]any{"xaligoUmlComponentInterfaceStem": true})
		appendUMLComponentInterfaceLineV1EngineSceneWalk(elements, group.baseID+"-multi-port-stem", trunkX, baseCenterY, port[0], baseCenterY, updated, map[string]any{"xaligoUmlComponentInterfaceStem": true})
		for index := 0; index < count; index++ {
			cy := startY + step*float64(index)
			appendUMLComponentInterfaceLineV1EngineSceneWalk(elements, fmt.Sprintf("%s-multi-circle-stem-%d", group.baseID, index), circleLeft+diameter, cy, trunkX, cy, updated, map[string]any{"xaligoUmlComponentInterfaceStem": true})
		}
	}
}

func bindUMLComponentInterfaceConnectionsV1EngineSceneBuild(connections []*entity.Node, endpoints umlComponentInterfaceEndpointsV1EngineSceneBuild) []*entity.Node {
	if len(connections) == 0 || len(endpoints.byOwner) == 0 {
		return connections
	}
	bound := make([]*entity.Node, 0, len(connections))
	usedByBase := map[string]int{}
	for _, conn := range connections {
		if conn == nil || conn.Attr("uml-diagram-kind") != "component-diagram" || conn.Attr("uml-relation-kind") != "association" || conn.Attr("uml-src-kind") != "component" || conn.Attr("uml-dst-kind") != "component" {
			bound = append(bound, conn)
			continue
		}
		srcKey := strings.TrimSpace(conn.Attrs[internalConnectionSrcKeyAttrV1EngineParseDocument])
		dstKey := strings.TrimSpace(conn.Attrs[internalConnectionDstKeyAttrV1EngineParseDocument])
		boundDst, ok := matchingUMLComponentInterfaceEndpointKeyV1EngineSceneBuild(endpoints, srcKey, dstKey)
		if !ok {
			bound = append(bound, conn)
			continue
		}
		group := endpoints.byKey[boundDst]
		if group == nil || len(group.endpointKeys) == 0 {
			bound = append(bound, conn)
			continue
		}
		srcSide, srcAnchor := umlComponentInterfaceSourceAnchorV1EngineSceneBuild(endpoints, srcKey, group.ownerRect)
		endpointIndex := usedByBase[group.baseKey]
		if endpointIndex >= len(group.endpointKeys) {
			endpointIndex = len(group.endpointKeys) - 1
		}
		boundDst = group.endpointKeys[endpointIndex]
		usedByBase[group.baseKey]++
		clone := *conn
		clone.Attrs = cloneAttrsV1EngineParseTable(conn.Attrs)
		clone.Attrs[internalConnectionSrcKeyAttrV1EngineParseDocument] = srcKey
		clone.Attrs[internalConnectionDstKeyAttrV1EngineParseDocument] = boundDst
		clone.Attrs["src-side"] = srcSide
		clone.Attrs["src-anchor"] = srcAnchor
		clone.Attrs["dst-side"] = endpoints.sideByEndpoint[boundDst]
		clone.Attrs["uml-component-interface-dst"] = "true"
		clone.Attrs["uml-component-caller-socket"] = "true"
		bound = append(bound, &clone)
	}
	return bound
}

func umlComponentInterfaceSourceAnchorV1EngineSceneBuild(endpoints umlComponentInterfaceEndpointsV1EngineSceneBuild, srcOwnerKey string, dstOwnerRect [4]float64) (string, string) {
	srcByLabel := endpoints.byOwner[srcOwnerKey]
	for _, group := range srcByLabel {
		if group == nil || group.ownerRect[2] <= 0 || group.ownerRect[3] <= 0 || dstOwnerRect[2] <= 0 || dstOwnerRect[3] <= 0 {
			continue
		}
		anchor := nearestUMLComponentCallerAnchorV1EngineSceneBuild(group.ownerRect, [2]float64{dstOwnerRect[0] + dstOwnerRect[2]/2, dstOwnerRect[1] + dstOwnerRect[3]/2})
		return string(anchor.side), anchor.StringV1EngineParseConnection()
	}
	return "right", "right-3"
}

func nearestUMLComponentCallerAnchorV1EngineSceneBuild(rect [4]float64, target [2]float64) connectionAnchorSpecV1EngineParseConnection {
	best := connectionAnchorSpecV1EngineParseConnection{side: sideRightV1EngineRouteTypes, slot: 2, hasSlot: true}
	bestDistance := math.Inf(1)
	for _, side := range []sideV1EngineRouteTypes{sideTopV1EngineRouteTypes, sideRightV1EngineRouteTypes, sideBottomV1EngineRouteTypes} {
		for slot := 0; slot < anchorGridV1EnginePlanBuild; slot++ {
			fp := fixedPointForSideSlotV1EngineSceneConnection(string(side), slot)
			x := rect[0] + rect[2]*fp[0]
			y := rect[1] + rect[3]*fp[1]
			distance := math.Hypot(target[0]-x, target[1]-y)
			if distance >= bestDistance {
				continue
			}
			best = connectionAnchorSpecV1EngineParseConnection{side: side, slot: slot, hasSlot: true}
			bestDistance = distance
		}
	}
	return best
}

func umlComponentInterfaceConnectionSideV1EngineSceneBuild(symbolRect, ownerRect [4]float64) string {
	if ownerRect[2] <= 0 || ownerRect[3] <= 0 {
		return "right"
	}
	symbolCenterX := symbolRect[0] + symbolRect[2]/2
	symbolCenterY := symbolRect[1] + symbolRect[3]/2
	ownerLeft := ownerRect[0]
	ownerRight := ownerRect[0] + ownerRect[2]
	ownerTop := ownerRect[1]
	ownerBottom := ownerRect[1] + ownerRect[3]
	if symbolCenterX < ownerLeft {
		return "left"
	}
	if symbolCenterX > ownerRight {
		return "right"
	}
	if symbolCenterY < ownerTop {
		return "bottom"
	}
	if symbolCenterY > ownerBottom {
		return "top"
	}
	if symbolCenterX < ownerLeft+ownerRect[2]/2 {
		return "left"
	}
	return "right"
}

func appendUMLComponentInterfaceCircleV1EngineSceneBuild(elements *[]map[string]any, id string, x, y, diameter float64, updated int64, group *umlComponentInterfaceEndpointGroupV1EngineSceneBuild) {
	if elements == nil || group == nil {
		return
	}
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "ellipse",
		"x": x, "y": y, "width": diameter, "height": diameter,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#ffffff",
		"fillStyle": "solid", "strokeWidth": 1.2, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"customData": map[string]any{
			"xaligoUmlComponentInterfaceLabel":     group.label,
			"xaligoUmlComponentOwnerConnectionKey": group.ownerKey,
			"xaligoUmlComponentInterfaceSymbol":    true,
			"xaligoUmlComponentInterfaceCircle":    true,
		},
	})
}

func moveUMLComponentInterfaceElementV1EngineSceneBuild(elements []map[string]any, id string, x, y float64) {
	for _, element := range elements {
		elementID, _ := element["id"].(string)
		if elementID != id {
			continue
		}
		element["x"] = x
		element["y"] = y
		return
	}
}

func removeUMLComponentInterfaceStemsV1EngineSceneBuild(elements *[]map[string]any, group *umlComponentInterfaceEndpointGroupV1EngineSceneBuild) {
	if elements == nil || group == nil {
		return
	}
	filtered := (*elements)[:0]
	for _, element := range *elements {
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlComponentInterfaceStem"] == true && customData["xaligoUmlComponentOwnerConnectionKey"] == group.ownerKey && customData["xaligoUmlComponentInterfaceLabel"] == group.label {
			continue
		}
		filtered = append(filtered, element)
	}
	*elements = filtered
}

func matchingUMLComponentInterfaceEndpointKeyV1EngineSceneBuild(endpoints umlComponentInterfaceEndpointsV1EngineSceneBuild, srcOwnerKey, dstOwnerKey string) (string, bool) {
	srcByLabel := endpoints.byOwner[srcOwnerKey]
	dstByLabel := endpoints.byOwner[dstOwnerKey]
	if len(srcByLabel) == 0 || len(dstByLabel) == 0 {
		return "", false
	}
	labels := make([]string, 0, len(srcByLabel))
	for label := range srcByLabel {
		if _, ok := dstByLabel[label]; ok {
			labels = append(labels, label)
		}
	}
	if len(labels) == 0 {
		return "", false
	}
	sort.Strings(labels)
	label := labels[0]
	return dstByLabel[label].baseKey, true
}

func umlComponentInterfaceEndpointKeyV1EngineSceneBuild(ownerKey, label string) string {
	return ownerKey + "|uml-component-interface|" + label
}

const umlComponentCallerSocketRadiusPaddingV1EngineSceneBuild = 2.0

func umlComponentCallerSocketRadiusForCircleV1EngineSceneBuild(circleRect [4]float64) float64 {
	circleRadius := math.Max(circleRect[2], circleRect[3]) / 2
	return circleRadius + umlComponentCallerSocketRadiusPaddingV1EngineSceneBuild
}

func umlComponentCallerSocketGapForCircleV1EngineSceneBuild(circleRect [4]float64) float64 {
	circleRadius := math.Max(circleRect[2], circleRect[3]) / 2
	return math.Max(0, umlComponentCallerSocketRadiusForCircleV1EngineSceneBuild(circleRect)-circleRadius)
}

func appendUMLComponentCallerSocketsV1EngineSceneBuild(elements *[]map[string]any) {
	if elements == nil {
		return
	}
	updated := excalidrawUpdatedV1EngineSceneTypes
	for _, element := range *elements {
		customData, _ := element["customData"].(map[string]any)
		if !boolishV1EngineSceneBuild(customData["xaligoUmlComponentCallerSocket"]) {
			continue
		}
		id, _ := element["id"].(string)
		x, okX := element["x"].(float64)
		y, okY := element["y"].(float64)
		points, ok := element["points"].([][]float64)
		if !ok || len(points) < 2 || id == "" || !okX || !okY {
			continue
		}
		absolutePoints := make([][2]float64, 0, len(points))
		for _, point := range points {
			if len(point) < 2 {
				continue
			}
			absolutePoints = append(absolutePoints, [2]float64{x + point[0], y + point[1]})
		}
		if len(absolutePoints) < 2 {
			continue
		}
		endpoint := absolutePoints[len(absolutePoints)-1]
		radius := positiveNumberV1EngineSceneBuild(customData["xaligoUmlComponentCallerSocketRadius"])
		if radius <= 0 {
			radius = 7
		}
		appendUMLComponentCallerSocketV1EngineSceneBuild(elements, id+"-caller-socket", endpoint, radius, updated)
	}
}

func appendUMLComponentCallerSocketV1EngineSceneBuild(elements *[]map[string]any, id string, endpoint [2]float64, radius float64, updated int64) {
	absolute := make([][2]float64, 0, 13)
	minX, minY := math.Inf(1), math.Inf(1)
	maxX, maxY := math.Inf(-1), math.Inf(-1)
	for index := 0; index <= 12; index++ {
		theta := -math.Pi/2 - math.Pi*float64(index)/12
		point := [2]float64{endpoint[0] + radius + math.Cos(theta)*radius, endpoint[1] + math.Sin(theta)*radius}
		absolute = append(absolute, point)
		minX = math.Min(minX, point[0])
		minY = math.Min(minY, point[1])
		maxX = math.Max(maxX, point[0])
		maxY = math.Max(maxY, point[1])
	}
	points := make([][]float64, 0, len(absolute))
	for _, point := range absolute {
		points = append(points, []float64{point[0] - minX, point[1] - minY})
	}
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "line",
		"x": minX, "y": minY, "width": maxX - minX, "height": maxY - minY,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1.35, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"points":     points,
		"customData": map[string]any{"xaligoUmlComponentCallerSocket": true},
	})
}

func boolishV1EngineSceneBuild(value any) bool {
	switch typed := value.(type) {
	case bool:
		return typed
	case string:
		return strings.EqualFold(strings.TrimSpace(typed), "true")
	default:
		return false
	}
}

func positiveNumberV1EngineSceneBuild(value any) float64 {
	switch typed := value.(type) {
	case float64:
		if typed > 0 {
			return typed
		}
	case int:
		if typed > 0 {
			return float64(typed)
		}
	case string:
		parsed, err := strconv.ParseFloat(strings.TrimSpace(typed), 64)
		if err == nil && parsed > 0 {
			return parsed
		}
	}
	return 0
}

// orderSceneLayers keeps connectors below readable content while preserving
// group headers above nested borders.
func orderSceneLayersV1EngineSceneBuild(elements []map[string]any) []map[string]any {
	base := make([]map[string]any, 0, len(elements))
	headShapes := make([]map[string]any, 0)
	connectors := make([]map[string]any, 0)
	anchorBackgrounds := make([]map[string]any, 0)
	anchorContent := make([]map[string]any, 0)
	headContent := make([]map[string]any, 0)
	metadataBackgrounds := make([]map[string]any, 0)
	metadataContent := make([]map[string]any, 0)
	for _, el := range elements {
		custom, _ := el["customData"].(map[string]any)
		if isMetadata, _ := custom["xaligoFrameMetadata"].(bool); isMetadata {
			if isContent, _ := custom["xaligoFrameMetadataContent"].(bool); isContent {
				metadataContent = append(metadataContent, el)
			} else {
				metadataBackgrounds = append(metadataBackgrounds, el)
			}
			continue
		}
		if isHeader, _ := custom["xaligoGroupHeader"].(bool); isHeader {
			headShapes = append(headShapes, el)
			continue
		}
		if isContent, _ := custom["xaligoGroupHeaderContent"].(bool); isContent {
			headContent = append(headContent, el)
			continue
		}
		if isAnchorBackground, _ := custom["xaligoAnchorBackground"].(bool); isAnchorBackground {
			anchorBackgrounds = append(anchorBackgrounds, el)
			continue
		}
		if isAnchorContent, _ := custom["xaligoAnchorContent"].(bool); isAnchorContent {
			anchorContent = append(anchorContent, el)
			continue
		}
		typ, _ := el["type"].(string)
		if isJunction, _ := custom["xaligoJunction"].(bool); typ == "arrow" || typ == "line" || isJunction {
			connectors = append(connectors, el)
			continue
		}
		base = append(base, el)
	}
	ordered := append(base, headShapes...)
	ordered = append(ordered, connectors...)
	ordered = append(ordered, anchorBackgrounds...)
	ordered = append(ordered, anchorContent...)
	ordered = append(ordered, headContent...)
	ordered = append(ordered, metadataBackgrounds...)
	return append(ordered, metadataContent...)
}

func avoidGroupHeaderOverlapV1EngineSceneBuild(x, y, w, h float64, ownBorderID string, elements []map[string]any) float64 {
	adjustedY := y
	for pass := 0; pass < 4; pass++ {
		nextY := adjustedY
		for _, el := range elements {
			if id, _ := el["id"].(string); id == ownBorderID {
				continue
			}
			custom, _ := el["customData"].(map[string]any)
			isBorder, _ := custom["xaligoGroupBorder"].(bool)
			isHeader, _ := custom["xaligoGroupHeader"].(bool)
			if !isBorder && !isHeader {
				continue
			}
			bx, okX := el["x"].(float64)
			by, okY := el["y"].(float64)
			bw, okW := el["width"].(float64)
			bh, okH := el["height"].(float64)
			if !okX || !okY || !okW || !okH || horizontalOverlapV1EngineSceneBuild(x, x+w, bx, bx+bw) <= 0 {
				continue
			}
			if isHeader {
				gap := float64(groupHeaderBorderGapV1EngineSceneTypes)
				if by < adjustedY+h+gap && by+bh > adjustedY-gap {
					nextY = math.Max(nextY, by+bh+gap)
				}
				continue
			}
			for _, lineY := range []float64{by, by + bh} {
				if lineY >= adjustedY-float64(groupHeaderBorderGapV1EngineSceneTypes) && lineY <= adjustedY+h+float64(groupHeaderBorderGapV1EngineSceneTypes) {
					nextY = math.Max(nextY, lineY+float64(groupHeaderBorderGapV1EngineSceneTypes))
				}
			}
		}
		if math.Abs(nextY-adjustedY) < 0.01 {
			break
		}
		adjustedY = nextY
	}
	return adjustedY
}

func horizontalOverlapV1EngineSceneBuild(a0, a1, b0, b1 float64) float64 {
	return math.Max(0, math.Min(math.Max(a0, a1), math.Max(b0, b1))-math.Max(math.Min(a0, a1), math.Min(b0, b1)))
}

func sceneTextLayoutV1EngineSceneBuild(role entity.TextRole, wrap bool, lineHeight float64) *entity.TextLayout {
	return &entity.TextLayout{
		Role:       role,
		Wrap:       wrap,
		Fit:        entity.TextFitShrink,
		Overflow:   entity.TextOverflowClip,
		Clip:       true,
		LineHeight: lineHeight,
	}
}

func alignGroupBorderTopToHeaderV1EngineSceneBuild(borderID string, topY, bottomY float64, elements []map[string]any) {
	for i := range elements {
		id, _ := elements[i]["id"].(string)
		if id != borderID {
			continue
		}
		if topY <= bottomY-MinBoxHeightV1EngineLayoutFlow {
			elements[i]["y"] = topY
			elements[i]["height"] = bottomY - topY
		}
		return
	}
}
