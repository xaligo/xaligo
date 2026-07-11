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
		"strokeColor": "#bbb", "backgroundColor": "transparent",
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
				"strokeColor": "#bbb", "backgroundColor": "transparent",
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
			registerConnectionEndpointV1EngineSceneWalk(child, pageFrameID, frameRects[frameID], itemImgRects, itemImgIDs)
		}
	}

	walkV1EngineSceneWalk(root, &elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, root, itemGroups, ancestorBoxes, itemImgRects, itemImgIDs, deps)
	ancestorIDs := make([]string, 0, len(itemGroups))
	for ancID := range itemGroups {
		ancestorIDs = append(ancestorIDs, ancID)
	}
	sort.Strings(ancestorIDs)
	for _, ancID := range ancestorIDs {
		items := itemGroups[ancID]
		if err := renderItemGridV1EngineSceneItem(items, ancestorBoxes[ancID], &elements, files, catalogCSV, projectRoot, fsys, itemIconSize, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps); err != nil {
			return nil, err
		}
	}
	applySemanticElementMetadataV1EngineSceneWalk(elements, collectSemanticElementMetadataV1EngineSceneWalk(root))
	renderConnectionsV1EngineSceneConnectionRender(connections, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, frameRects, &elements)
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

// orderSceneLayers keeps connectors below readable content while preserving
// group headers above nested borders.
func orderSceneLayersV1EngineSceneBuild(elements []map[string]any) []map[string]any {
	base := make([]map[string]any, 0, len(elements))
	headShapes := make([]map[string]any, 0)
	connectors := make([]map[string]any, 0)
	anchorBackgrounds := make([]map[string]any, 0)
	anchorContent := make([]map[string]any, 0)
	headContent := make([]map[string]any, 0)
	for _, el := range elements {
		custom, _ := el["customData"].(map[string]any)
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
	return append(ordered, headContent...)
}

func avoidGroupHeaderBorderOverlapV1EngineSceneBuild(x, y, w, h float64, ownBorderID string, elements []map[string]any) float64 {
	adjustedY := y
	for pass := 0; pass < 4; pass++ {
		nextY := adjustedY
		for _, el := range elements {
			if id, _ := el["id"].(string); id == ownBorderID {
				continue
			}
			custom, _ := el["customData"].(map[string]any)
			if isBorder, _ := custom["xaligoGroupBorder"].(bool); !isBorder {
				continue
			}
			bx, okX := el["x"].(float64)
			by, okY := el["y"].(float64)
			bw, okW := el["width"].(float64)
			bh, okH := el["height"].(float64)
			if !okX || !okY || !okW || !okH || horizontalOverlapV1EngineSceneBuild(x, x+w, bx, bx+bw) <= 0 {
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
