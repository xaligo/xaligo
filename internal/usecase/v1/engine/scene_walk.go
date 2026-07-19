package engine

import (
	"fmt"
	"io/fs"
	"math"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type semanticElementMetadataV1EngineSceneWalk struct {
	ParentElementID string
	Kind            string
}

func walkV1EngineSceneWalk(b *entity.Box, elements *[]map[string]any, files map[string]any, svgGroupDir string, catalogCSV string, projectRoot string, fsys fs.FS, visibleAncestor *entity.Box, itemGroups map[string][]*entity.Box, ancestorBoxes, itemFrames map[string]*entity.Box, itemImgRects map[string][4]float64, itemImgIDs map[string]string, activeFrame *entity.Box, deps SceneDependenciesV1EngineSceneTypes) {
	if b != nil && b.Tag == "frame" {
		activeFrame = b
	}
	if IsItemLikeV1EngineLayoutAttributes(b.Tag) {
		// 描画はしない: visibleAncestor に結び付けて収集のみ (<item> / <spacer> 共通)
		key := visibleAncestor.ID
		itemGroups[key] = append(itemGroups[key], b)
		ancestorBoxes[key] = visibleAncestor
		itemFrames[key] = activeFrame
		return
	}

	// selfVisible=false のとき: 自身の描画 (枠・アイコン・ラベル) はスキップするが
	// 子要素の描画は継続する (親子関係なく個別に制御可能)。
	selfVisible := b.Attrs["visible"] != "false"

	if b.Tag != "frame" && selfVisible {
		updated := excalidrawUpdatedV1EngineSceneTypes

		noBorder := b.Attrs["border"] == "none"

		if gd, isGroup := awsGroupsV1EngineSceneTypes[b.Tag]; isGroup {
			// ── AWS group border ────────────────────────────────────
			rectID := fmt.Sprintf("%s-rect", b.ID)
			groupStroke := gd.StrokeColor
			if noBorder {
				groupStroke = "transparent"
			}
			*elements = append(*elements, map[string]any{
				"id": rectID, "type": "rectangle",
				"x": b.X, "y": b.Y, "width": b.W, "height": b.H,
				"angle":       0,
				"strokeColor": groupStroke, "backgroundColor": staggerBGColorV1EngineSceneTypes(b),
				"fillStyle":   "solid",
				"strokeWidth": gd.StrokeWidth, "strokeStyle": gd.StrokeStyle,
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": stableSceneSeedV1EngineSceneTypes(rectID), "version": 1,
				"versionNonce": stableSceneSeedV1EngineSceneTypes(rectID),
				"isDeleted":    false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"customData": map[string]any{"xaligoGroupBorder": true},
			})
			registerConnectionEndpointV1EngineSceneWalk(b, rectID, [4]float64{b.X, b.Y, b.W, b.H}, itemImgRects, itemImgIDs)

			// ── Group icon ──────────────────────────────────────────
			headerX := b.X - groupHeaderLeftOverflowV1EngineSceneTypes
			textX := headerX + groupHeaderTextInsetV1EngineSceneTypes
			var iconDataURL, iconFileID, iconBackground string
			if b.Tag == "generic-group" && strings.TrimSpace(b.Attrs["icon-id"]) != "" {
				catalogID, _ := strconv.Atoi(strings.TrimSpace(b.Attrs["icon-id"]))
				entry, err := deps.lookupCatalogByIDV1EngineSceneTypes(fsys, catalogCSV, catalogID)
				if err == nil && entry.DataURL == "" && entry.RelPath != "" && projectRoot != "" {
					entry.DataURL, err = deps.svgToDataURLV1EngineSceneTypes(nil, filepath.Join(projectRoot, entry.RelPath))
				}
				if err != nil {
					loggerV1EngineSharedLogging.WARN(IUESW002V1EngineSceneTypes, "generic group icon lookup failed", map[string]any{"catalogID": catalogID, "error": err})
				} else {
					iconDataURL = entry.DataURL
					iconFileID = fmt.Sprintf("group-cat-%d", catalogID)
					iconBackground = deps.svgBGColorV1EngineSceneTypes(entry.DataURL)
				}
			} else if gd.IconFile != "" && svgGroupDir != "" {
				iconPath := filepath.Join(svgGroupDir, gd.IconFile)
				if fsys != nil {
					// In embedded mode, use forward slashes even on Windows.
					iconPath = svgGroupDir + "/" + gd.IconFile
				}
				loadedDataURL, err := deps.svgToDataURLV1EngineSceneTypes(fsys, iconPath)
				iconDataURL = loadedDataURL
				if err != nil {
					iconDataURL = ""
				}
				iconFileID = deps.fileIDV1EngineSceneTypes(gd.IconFile)
				iconBackground = "transparent"
			}
			if iconDataURL != "" {
				iconDataURL = tintSVGDataURLV1EngineSceneTypes(iconDataURL, gd.StrokeColor)
				iconBackground = "transparent"
				textX = headerX + float64(groupIconSizeV1EngineSceneTypes) + groupHeaderTextInsetV1EngineSceneTypes
			}
			lblW := textWidthV1EngineSceneItem(b.Label, groupLabelCharWV1EngineSceneTypes)
			headerBackground := staggerBGColorV1EngineSceneTypes(b)
			if headerBackground == "transparent" {
				headerBackground = "#ffffff"
			}
			// Extend the opaque header mask beyond the group's left border so the
			// vertical border cannot show through beside a catalog icon.
			headerH := float64(groupTextHeightV1EngineSceneTypes + groupHeaderTextPadYV1EngineSceneTypes*2)
			if iconDataURL != "" {
				headerH = float64(groupIconSizeV1EngineSceneTypes)
			}
			headerTip := math.Min(groupHeaderTipMaxV1EngineSceneTypes, headerH/2)
			headerW := textX + lblW + groupHeaderPadEndV1EngineSceneTypes + headerTip - headerX
			headerY := avoidGroupHeaderOverlapV1EngineSceneBuild(headerX, b.Y-headerH/2, headerW, headerH, rectID, *elements)
			headerY = groupHeaderYAvoidingFrameMetadataV1EngineSceneWalk(headerY, headerH, activeFrame)
			alignGroupBorderTopToHeaderV1EngineSceneBuild(rectID, headerY+headerH/2, b.Y+b.H, *elements)
			headerID := fmt.Sprintf("%s-header-bg", b.ID)
			headerSeed := stableSceneSeedV1EngineSceneTypes(headerID)
			*elements = append(*elements, map[string]any{
				"id": headerID, "type": "line",
				"x": headerX, "y": headerY,
				"width": headerW, "height": headerH,
				"points": [][]float64{{0, 0}, {headerW - headerTip, 0}, {headerW, headerH / 2}, {headerW - headerTip, headerH}, {0, headerH}, {0, 0}},
				"angle":  0, "strokeColor": gd.StrokeColor, "backgroundColor": headerBackground,
				"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": headerSeed, "version": 1, "versionNonce": headerSeed,
				"isDeleted": false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"customData": map[string]any{"xaligoGroupHeader": true},
			})
			if iconDataURL != "" {
				iconID := fmt.Sprintf("%s-icon", b.ID)
				iconSeed := stableSceneSeedV1EngineSceneTypes(iconID)
				*elements = append(*elements, map[string]any{
					"id": iconID, "type": "image",
					"x": headerX, "y": headerY + (headerH-float64(groupIconSizeV1EngineSceneTypes))/2,
					"width": float64(groupIconSizeV1EngineSceneTypes), "height": float64(groupIconSizeV1EngineSceneTypes),
					"fileId": iconFileID, "status": "saved", "scale": []int{1, 1},
					"strokeColor": "transparent", "backgroundColor": iconBackground,
					"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
					"roughness": 0, "opacity": 100, "angle": 0,
					"seed": iconSeed, "version": 1, "versionNonce": iconSeed,
					"isDeleted": false, "groupIds": []string{},
					"frameId": nil, "boundElements": nil,
					"updated": updated, "link": nil, "locked": false,
					"customData": map[string]any{"xaligoGroupHeaderContent": true},
				})
				if _, exists := files[iconFileID]; !exists {
					files[iconFileID] = map[string]any{
						"mimeType": "image/svg+xml", "id": iconFileID, "dataURL": iconDataURL,
						"created": updated, "lastRetrieved": updated,
					}
				}
			}

			// ── AWS group label ─────────────────────────────────────
			textY := headerY + (headerH-float64(groupTextHeightV1EngineSceneTypes))/2
			labelID := fmt.Sprintf("%s-label", b.ID)
			labelSeed := stableSceneSeedV1EngineSceneTypes(labelID)
			// groupFontFamily=2 (Helvetica 14px): ~7.5px/rune
			*elements = append(*elements, map[string]any{
				"id": labelID, "type": "text",
				"x": textX, "y": textY,
				"width": lblW, "height": float64(groupTextHeightV1EngineSceneTypes),
				"angle":       0,
				"strokeColor": gd.StrokeColor, "backgroundColor": "transparent",
				"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": labelSeed, "version": 1,
				"versionNonce": labelSeed,
				"isDeleted":    false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"text": b.Label, "fontSize": groupFontSizeV1EngineSceneTypes, "fontFamily": groupFontFamilyV1EngineSceneTypes,
				"textAlign": "left", "verticalAlign": "middle",
				"containerId": nil, "originalText": b.Label, "lineHeight": 1.25,
				"customData": map[string]any{
					"xaligoGroupHeaderContent": true,
					"xaligoTextLayout":         sceneTextLayoutV1EngineSceneBuild(entity.TextRoleGroupHeader, false, 1.25),
				},
			})
		} else if !isLayoutTagV1EngineSceneWalk(b.Tag) {
			// ── Generic tag: rectangle + label ──────────────────────
			rectID := fmt.Sprintf("%s-rect", b.ID)
			textID := fmt.Sprintf("%s-text", b.ID)
			genStroke := "#1e1e1e"
			if configured := strings.TrimSpace(b.Attrs["border-color"]); configured != "" {
				genStroke = configured
			}
			if noBorder {
				genStroke = "transparent"
			}
			backgroundColor := "transparent"
			fillStyle := "hachure"
			roundness := map[string]any{"type": 3}
			if b.Tag == "rectangle" || b.Tag == "table" || b.Tag == "entity" || b.Tag == "table-cell" {
				fillStyle = "solid"
			}
			if b.Tag == "table" || b.Tag == "entity" || b.Tag == "table-cell" {
				backgroundColor = "#ffffff"
				roundness = nil
			}
			if b.Tag == "table-cell" && b.Attrs["_xaligoTableHeader"] == "true" {
				backgroundColor = "#f1f5f9"
			}
			if b.Tag == "port" {
				backgroundColor = "#ffffff"
				fillStyle = "solid"
				roundness = nil
			}
			if configured, exists := b.Attrs["background-color"]; exists {
				backgroundColor = configured
			}
			boundElements := any(nil)
			shapeCustomData := umlShapeCustomDataV1EngineSceneWalk(b)
			if b.Label != "" {
				boundElements = []map[string]any{{"type": "text", "id": textID}}
			}
			shapeType := umlShapeTypeV1EngineSceneWalk(b)
			*elements = append(*elements, map[string]any{
				"id": rectID, "type": shapeType,
				"x": b.X, "y": b.Y, "width": b.W, "height": b.H,
				"angle": 0, "strokeColor": genStroke, "backgroundColor": backgroundColor,
				"fillStyle": fillStyle, "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": roundness,
				"seed": stableSceneSeedV1EngineSceneTypes(rectID), "version": 1,
				"versionNonce":  stableSceneSeedV1EngineSceneTypes(rectID),
				"isDeleted":     false,
				"boundElements": boundElements,
				"updated":       updated, "link": nil, "locked": false,
				"customData": shapeCustomData,
			})
			if b.Tag == "entity" {
				registerConnectionEndpointV1EngineSceneWalk(b, rectID, [4]float64{b.X, b.Y, b.W, b.H}, itemImgRects, itemImgIDs)
			}
			if b.Tag == "rectangle" || b.Tag == "port" {
				registerConnectionEndpointV1EngineSceneWalk(b, rectID, [4]float64{b.X, b.Y, b.W, b.H}, itemImgRects, itemImgIDs)
			}
			if b.Label != "" {
				fontSize := attrFloatV1EngineLayoutAttributes(b.Attrs["font-size"], 20)
				textX, textY := b.X+4, b.Y+2
				textW, textH := math.Max(1, b.W-8), math.Max(1, math.Min(math.Ceil(fontSize*1.2), b.H-4))
				textAlign, verticalAlign := "left", "top"
				role := entity.TextRoleLabel
				textCustomData := map[string]any{}
				if b.Tag == "rectangle" || b.Tag == "port" || b.Tag == "table-cell" {
					textX, textY = b.X+4, b.Y+2
					textW, textH = math.Max(1, b.W-8), math.Max(1, b.H-4)
					textAlign, verticalAlign = "center", "middle"
				}
				if b.Tag == "table-cell" {
					switch {
					case strings.HasSuffix(b.Attrs["align"], "-right") || b.Attrs["align"] == "right":
						textAlign = "right"
					case strings.HasSuffix(b.Attrs["align"], "-center") || b.Attrs["align"] == "center":
						textAlign = "center"
					default:
						textAlign = "left"
					}
				}
				if b.Tag == "port" {
					role = entity.TextRolePortLabel
					textCustomData["xaligoPortLabel"] = true
				}
				textCustomData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(role, true, 1.2)
				*elements = append(*elements, map[string]any{
					"id": textID, "type": "text",
					"x": textX, "y": textY,
					"width": textW, "height": textH,
					"angle":       0,
					"strokeColor": tableTextColorV1EngineSceneWalk(b), "backgroundColor": "transparent",
					"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
					"roughness": 0, "opacity": 100,
					"groupIds": []string{}, "roundness": nil,
					"seed": stableSceneSeedV1EngineSceneTypes(textID), "version": 1,
					"versionNonce": stableSceneSeedV1EngineSceneTypes(textID),
					"isDeleted":    false, "boundElements": nil,
					"updated": updated, "link": nil, "locked": false,
					"text": b.Label, "fontSize": fontSize, "fontFamily": fontFamilyV1EngineSceneWalk(b.Attrs["font-family"]),
					"textAlign": textAlign, "verticalAlign": verticalAlign,
					"containerId": rectID, "originalText": b.Label, "lineHeight": 1.2,
					"customData": textCustomData,
				})
			}
		}
	}

	// Stagger background layers: render border + label only, skip children.
	if b.IsStaggerBg {
		return
	}
	// 非表示要素は visibleAncestor を引き継ぐ (子の item が正しい親に紐付くよう)
	nextVisible := b
	if !selfVisible {
		nextVisible = visibleAncestor
	}
	for _, c := range b.Children {
		walkV1EngineSceneWalk(c, elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, nextVisible, itemGroups, ancestorBoxes, itemFrames, itemImgRects, itemImgIDs, activeFrame, deps)
	}
}

func groupHeaderYAvoidingFrameMetadataV1EngineSceneWalk(y, height float64, frame *entity.Box) float64 {
	if frame == nil || frame.FrameMetadata == nil || frame.FrameMetadata.ReservedW <= 0 || frame.FrameMetadata.ReservedH <= 0 {
		return y
	}
	metadata := frame.FrameMetadata
	if metadata.Position == "bottom" {
		return math.Min(y, metadata.ReservedY-height)
	}
	return math.Max(y, metadata.ReservedY+metadata.ReservedH)
}

func umlShapeCustomDataV1EngineSceneWalk(box *entity.Box) map[string]any {
	customData := map[string]any{}
	if box == nil {
		return customData
	}
	for source, target := range map[string]string{
		"uml-id":                "xaligoUmlId",
		"uml-local-id":          "xaligoUmlLocalId",
		"uml-ref":               "xaligoUmlReference",
		"uml-kind":              "xaligoUmlDiagramKind",
		"uml-diagram-kind":      "xaligoUmlDiagramKind",
		"uml-element-kind":      "xaligoUmlElementKind",
		"uml-owner-id":          "xaligoUmlOwnerId",
		"uml-owner-ref":         "xaligoUmlOwnerReference",
		"uml-compartment-kinds": "xaligoUmlCompartmentKinds",
		"from":                  "xaligoUmlTimeFrom",
		"to":                    "xaligoUmlTimeTo",
	} {
		if value := strings.TrimSpace(box.Attrs[source]); value != "" {
			customData[target] = value
		}
	}
	if len(customData) == 0 {
		return nil
	}
	return customData
}

func umlShapeTypeV1EngineSceneWalk(box *entity.Box) string {
	if box == nil {
		return "rectangle"
	}
	switch strings.TrimSpace(box.Attrs["uml-element-kind"]) {
	case "use-case", "initial", "final":
		return "ellipse"
	case "decision", "merge", "choice", "history":
		return "diamond"
	default:
		return "rectangle"
	}
}

func tableTextColorV1EngineSceneWalk(box *entity.Box) string {
	if box != nil && strings.TrimSpace(box.Attrs["color"]) != "" {
		return box.Attrs["color"]
	}
	return "#1e1e1e"
}

func fontFamilyV1EngineSceneWalk(name string) int {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "helvetica":
		return 2
	case "cascadia":
		return 3
	case "assistant":
		return 4
	case "excalifont":
		return 5
	case "nunito":
		return 6
	case "lilita-one":
		return 7
	case "comic-shanns":
		return 8
	case "liberation-sans":
		return 9
	default:
		return 1
	}
}

// collectSemanticElementMetadataV1EngineSceneWalk projects the resolved Box
// tree into the semantic nodes emitted by scene construction. Pure layout and
// invisible boxes deliberately do not become parents: their descendants keep
// the nearest rendered semantic ancestor instead.
func collectSemanticElementMetadataV1EngineSceneWalk(root *entity.Box) map[string]semanticElementMetadataV1EngineSceneWalk {
	metadata := map[string]semanticElementMetadataV1EngineSceneWalk{}
	var visit func(*entity.Box, string, bool)
	visit = func(box *entity.Box, parentElementID string, isRoot bool) {
		if box == nil {
			return
		}

		if IsItemLikeV1EngineLayoutAttributes(box.Tag) {
			if box.Tag == "item" && strings.TrimSpace(box.Attrs["id"]) != "" {
				metadata[box.ID+"-item"] = semanticElementMetadataV1EngineSceneWalk{
					ParentElementID: parentElementID,
					Kind:            "item",
				}
			}
			return
		}

		semanticParentElementID := parentElementID
		semanticElementID := ""
		semanticKind := ""
		selfVisible := box.Attrs["visible"] != "false"
		switch {
		case box.Tag == "frame" && !isRoot:
			semanticElementID = pageFrameElementIDV1EngineSceneWalk(box)
			semanticKind = "frame"
		case box.Tag != "frame" && selfVisible && !isLayoutTagV1EngineSceneWalk(box.Tag):
			semanticElementID = box.ID + "-rect"
			semanticKind = semanticElementKindV1EngineSceneWalk(box)
		}
		if semanticElementID != "" {
			metadata[semanticElementID] = semanticElementMetadataV1EngineSceneWalk{
				ParentElementID: parentElementID,
				Kind:            semanticKind,
			}
			semanticParentElementID = semanticElementID
		}

		if box.IsStaggerBg {
			return
		}
		for _, child := range box.Children {
			visit(child, semanticParentElementID, false)
		}
	}
	visit(root, "", true)
	return metadata
}

func applySemanticElementMetadataV1EngineSceneWalk(elements []map[string]any, metadata map[string]semanticElementMetadataV1EngineSceneWalk) {
	for _, element := range elements {
		id, _ := element["id"].(string)
		semantic, ok := metadata[id]
		if !ok {
			continue
		}
		customData, _ := element["customData"].(map[string]any)
		if customData == nil {
			customData = map[string]any{}
			element["customData"] = customData
		}
		customData["xaligoSemanticElementKind"] = semantic.Kind
		if semantic.ParentElementID != "" {
			customData["xaligoSemanticParentElementId"] = semantic.ParentElementID
		}
	}
}

func pageFrameElementIDV1EngineSceneWalk(box *entity.Box) string {
	frameID := strings.TrimSpace(box.Attrs["id"])
	if frameID == "" {
		frameID = box.ID
	}
	// This ID is also the semantic-parent key used for page projection. It
	// must therefore preserve the validated frame ID losslessly; a sanitized
	// display form is not a safe identity key because distinct punctuation can
	// collapse to the same value.
	return "paper-frame-" + frameID
}

func semanticElementKindV1EngineSceneWalk(box *entity.Box) string {
	if box == nil {
		return ""
	}
	switch box.Tag {
	case "table":
		return "table"
	case "database":
		return "database"
	case "entity":
		return "entity"
	case "table-cell":
		return "table-cell"
	case "rectangle":
		return "rectangle"
	case "port":
		return "port"
	}
	if _, isGroup := awsGroupsV1EngineSceneTypes[box.Tag]; isGroup || len(box.Children) > 0 {
		return "group"
	}
	return "rectangle"
}

func registerConnectionEndpointV1EngineSceneWalk(b *entity.Box, elementID string, rect [4]float64, endpointRects map[string][4]float64, endpointIDs map[string]string) {
	if b == nil || endpointRects == nil || endpointIDs == nil || elementID == "" {
		return
	}
	key := strings.TrimSpace(b.Attrs[internalConnectionKeyAttrV1EngineParseDocument])
	if key == "" {
		return
	}
	endpointRects[key] = rect
	endpointIDs[key] = elementID
}

// isLayoutTag reports whether a tag is a pure layout container
// (<row>, <col>, <container>) that should not render any visible border or label.
func isLayoutTagV1EngineSceneWalk(tag string) bool {
	return tag == "frames" || tag == "row" || tag == "col" || tag == "container" || tag == "table-header" || tag == "table-row" || IsBlankV1EngineLayoutAttributes(tag)
}
