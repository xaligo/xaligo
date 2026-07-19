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
			strokeWidth := 1.0
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
			activityContainer := isUMLActivityContainerV1EngineSceneWalk(b)
			hiddenUMLContainer := isUMLContainerV1EngineSceneWalk(b)
			if activityContainer {
				genStroke = "#052d6e"
				strokeWidth = 1.5
				appendUMLActivityPartitionsV1EngineSceneWalk(b, elements, updated)
			}
			if isXaligoActivityShapeV1EngineSceneWalk(b) {
				genStroke = "#052d6e"
				strokeWidth = 1.35
				backgroundColor = umlActivityShapeBackgroundV1EngineSceneWalk(b)
				fillStyle = "solid"
			}
			classShape := isXaligoClassShapeV1EngineSceneWalk(b)
			if classShape {
				genStroke = "#052d6e"
				strokeWidth = 1.35
				backgroundColor = umlClassShapeBackgroundV1EngineSceneWalk(b)
				fillStyle = "solid"
			}
			componentShape := isXaligoComponentDiagramShapeV1EngineSceneWalk(b)
			if componentShape {
				genStroke = "#052d6e"
				strokeWidth = 1.35
				backgroundColor = "#ffffff"
				fillStyle = "solid"
			}
			stateMachineShape := isXaligoStateMachineShapeV1EngineSceneWalk(b)
			if stateMachineShape {
				genStroke = "#052d6e"
				strokeWidth = 1.35
				backgroundColor = umlStateMachineShapeBackgroundV1EngineSceneWalk(b)
				fillStyle = "solid"
			}
			sequenceLifeline := isXaligoSequenceLifelineV1EngineSceneWalk(b)
			if sequenceLifeline {
				appendUMLSequenceLifelineV1EngineSceneWalk(b, elements, itemImgRects, itemImgIDs, updated)
			} else if !hiddenUMLContainer {
				boundElements := any(nil)
				shapeCustomData := umlShapeCustomDataV1EngineSceneWalk(b)
				if b.Label != "" {
					boundElements = []map[string]any{{"type": "text", "id": textID}}
				}
				if classShape {
					boundElements = []map[string]any{{"type": "text", "id": fmt.Sprintf("%s-class-header-text", b.ID)}, {"type": "text", "id": fmt.Sprintf("%s-class-body-text", b.ID)}}
				}
				if isXaligoComponentShapeV1EngineSceneWalk(b) {
					boundElements = []map[string]any{{"type": "text", "id": fmt.Sprintf("%s-component-header-text", b.ID)}}
				}
				shapeType := umlShapeTypeV1EngineSceneWalk(b)
				*elements = append(*elements, map[string]any{
					"id": rectID, "type": shapeType,
					"x": b.X, "y": b.Y, "width": b.W, "height": b.H,
					"angle": 0, "strokeColor": genStroke, "backgroundColor": backgroundColor,
					"fillStyle": fillStyle, "strokeWidth": strokeWidth, "strokeStyle": "solid",
					"roughness": 0, "opacity": 100,
					"groupIds": []string{}, "roundness": roundness,
					"seed": stableSceneSeedV1EngineSceneTypes(rectID), "version": 1,
					"versionNonce":  stableSceneSeedV1EngineSceneTypes(rectID),
					"isDeleted":     false,
					"boundElements": boundElements,
					"updated":       updated, "link": nil, "locked": false,
					"customData": shapeCustomData,
				})
				if isXaligoFinalWithDotV1EngineSceneWalk(b) {
					appendUMLFinalDotV1EngineSceneWalk(b, elements, updated)
				}
				if classShape {
					appendUMLClassCompartmentsV1EngineSceneWalk(b, elements, updated)
				}
				if isXaligoStateMachineStateShapeV1EngineSceneWalk(b) {
					appendUMLStateMachineCompartmentsV1EngineSceneWalk(b, elements, updated)
				}
				if isXaligoComponentShapeV1EngineSceneWalk(b) {
					appendUMLComponentHeaderV1EngineSceneWalk(b, elements, updated)
					appendUMLComponentBoundaryInterfacesV1EngineSceneWalk(b, elements, updated)
				}
				if b.Tag == "entity" {
					registerConnectionEndpointV1EngineSceneWalk(b, rectID, [4]float64{b.X, b.Y, b.W, b.H}, itemImgRects, itemImgIDs)
				}
				if b.Tag == "rectangle" || b.Tag == "port" {
					registerConnectionEndpointV1EngineSceneWalk(b, rectID, [4]float64{b.X, b.Y, b.W, b.H}, itemImgRects, itemImgIDs)
				}
			}
			if !hiddenUMLContainer && !sequenceLifeline && b.Label != "" && !isXaligoClassShapeV1EngineSceneWalk(b) && !isXaligoStateMachineStateShapeV1EngineSceneWalk(b) && !isXaligoComponentShapeV1EngineSceneWalk(b) {
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
				if isXaligoStateMachineShapeV1EngineSceneWalk(b) {
					textX, textY = b.X+8, b.Y+6
					textW, textH = math.Max(1, b.W-16), math.Max(1, b.H-12)
					textAlign, verticalAlign = "center", "middle"
					if strings.Contains(b.Label, "\n") {
						textAlign, verticalAlign = "left", "top"
					}
				}
				textCustomData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(role, true, 1.2)
				*elements = append(*elements, map[string]any{
					"id": textID, "type": "text",
					"x": textX, "y": textY,
					"width": textW, "height": textH,
					"angle":       0,
					"strokeColor": umlShapeTextColorV1EngineSceneWalk(b), "backgroundColor": "transparent",
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

func appendUMLActivityPartitionsV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64) {
	partitions := umlActivityPartitionsV1EngineSceneWalk(box)
	if len(partitions) == 0 {
		return
	}
	if strings.TrimSpace(box.Attrs["layout"]) == "horizontal" || strings.TrimSpace(box.Attrs["lanes"]) == "horizontal" {
		appendUMLActivityHorizontalPartitionsV1EngineSceneWalk(box, elements, updated, partitions)
		return
	}
	innerX, innerY := box.X+8, box.Y+8
	innerW, innerH := math.Max(1, box.W-16), math.Max(1, box.H-16)
	headerH := math.Min(44, innerH*0.2)
	laneW := innerW / float64(len(partitions))
	for index, partition := range partitions {
		x := innerX + float64(index)*laneW
		width := laneW
		if index == len(partitions)-1 {
			width = innerX + innerW - x
		}
		backgroundID := fmt.Sprintf("%s-partition-%s-bg", box.ID, partition.id)
		backgroundSeed := stableSceneSeedV1EngineSceneTypes(backgroundID)
		backgroundColor := "#ffffff"
		if index%2 == 1 {
			backgroundColor = "#f8fbff"
		}
		*elements = append(*elements, map[string]any{
			"id": backgroundID, "type": "rectangle",
			"x": x, "y": innerY, "width": width, "height": innerH,
			"angle": 0, "strokeColor": "#052d6e", "backgroundColor": backgroundColor,
			"fillStyle": "solid", "strokeWidth": 1.15, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": backgroundSeed, "version": 1, "versionNonce": backgroundSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"customData": map[string]any{"xaligoUmlPartition": true, "xaligoUmlPartitionId": partition.id, "xaligoUmlPartitionTitle": partition.title},
		})
		headerID := fmt.Sprintf("%s-partition-%s-header", box.ID, partition.id)
		headerSeed := stableSceneSeedV1EngineSceneTypes(headerID)
		*elements = append(*elements, map[string]any{
			"id": headerID, "type": "rectangle",
			"x": x, "y": innerY, "width": width, "height": headerH,
			"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#08b8ea",
			"fillStyle": "solid", "strokeWidth": 1.25, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": headerSeed, "version": 1, "versionNonce": headerSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"customData": map[string]any{"xaligoUmlPartitionHeader": true, "xaligoUmlPartitionId": partition.id, "xaligoUmlPartitionTitle": partition.title},
		})
		textID := fmt.Sprintf("%s-partition-%s-title", box.ID, partition.id)
		textSeed := stableSceneSeedV1EngineSceneTypes(textID)
		*elements = append(*elements, map[string]any{
			"id": textID, "type": "text",
			"x": x + 4, "y": innerY + 8,
			"width": math.Max(1, width-8), "height": math.Max(1, headerH-12),
			"angle": 0, "strokeColor": "#ffffff", "backgroundColor": "transparent",
			"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": textSeed, "version": 1, "versionNonce": textSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"text": partition.title, "fontSize": 18, "fontFamily": 2,
			"textAlign": "center", "verticalAlign": "middle",
			"containerId": nil, "originalText": partition.title, "lineHeight": 1.2,
			"customData": map[string]any{"xaligoUmlPartitionHeaderContent": true, "xaligoTextLayout": sceneTextLayoutV1EngineSceneBuild(entity.TextRoleGroupHeader, false, 1.2)},
		})
	}
}

func appendUMLActivityHorizontalPartitionsV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64, partitions []umlActivityPartitionV1EngineSceneWalk) {
	innerX, innerY := box.X+8, box.Y+8
	innerW, innerH := math.Max(1, box.W-16), math.Max(1, box.H-16)
	headerW := math.Min(132, innerW*0.24)
	laneH := innerH / float64(len(partitions))
	for index, partition := range partitions {
		y := innerY + float64(index)*laneH
		height := laneH
		if index == len(partitions)-1 {
			height = innerY + innerH - y
		}
		backgroundID := fmt.Sprintf("%s-partition-%s-bg", box.ID, partition.id)
		backgroundSeed := stableSceneSeedV1EngineSceneTypes(backgroundID)
		backgroundColor := "#ffffff"
		if index%2 == 1 {
			backgroundColor = "#f8fbff"
		}
		*elements = append(*elements, map[string]any{
			"id": backgroundID, "type": "rectangle",
			"x": innerX, "y": y, "width": innerW, "height": height,
			"angle": 0, "strokeColor": "#052d6e", "backgroundColor": backgroundColor,
			"fillStyle": "solid", "strokeWidth": 1.15, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": backgroundSeed, "version": 1, "versionNonce": backgroundSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"customData": map[string]any{"xaligoUmlPartition": true, "xaligoUmlPartitionId": partition.id, "xaligoUmlPartitionTitle": partition.title},
		})
		headerID := fmt.Sprintf("%s-partition-%s-header", box.ID, partition.id)
		headerSeed := stableSceneSeedV1EngineSceneTypes(headerID)
		*elements = append(*elements, map[string]any{
			"id": headerID, "type": "rectangle",
			"x": innerX, "y": y, "width": headerW, "height": height,
			"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#08b8ea",
			"fillStyle": "solid", "strokeWidth": 1.25, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": headerSeed, "version": 1, "versionNonce": headerSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"customData": map[string]any{"xaligoUmlPartitionHeader": true, "xaligoUmlPartitionId": partition.id, "xaligoUmlPartitionTitle": partition.title},
		})
		textID := fmt.Sprintf("%s-partition-%s-title", box.ID, partition.id)
		textSeed := stableSceneSeedV1EngineSceneTypes(textID)
		*elements = append(*elements, map[string]any{
			"id": textID, "type": "text",
			"x": innerX + 4, "y": y + 8,
			"width": math.Max(1, headerW-8), "height": math.Max(1, height-16),
			"angle": 0, "strokeColor": "#ffffff", "backgroundColor": "transparent",
			"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": textSeed, "version": 1, "versionNonce": textSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"text": partition.title, "fontSize": 18, "fontFamily": 2,
			"textAlign": "center", "verticalAlign": "middle",
			"containerId": nil, "originalText": partition.title, "lineHeight": 1.2,
			"customData": map[string]any{"xaligoUmlPartitionHeaderContent": true, "xaligoTextLayout": sceneTextLayoutV1EngineSceneBuild(entity.TextRoleGroupHeader, false, 1.2)},
		})
	}
}

type umlActivityPartitionV1EngineSceneWalk struct {
	id    string
	title string
}

func umlActivityPartitionsV1EngineSceneWalk(box *entity.Box) []umlActivityPartitionV1EngineSceneWalk {
	seen := map[string]bool{}
	var partitions []umlActivityPartitionV1EngineSceneWalk
	for _, child := range box.Children {
		id := strings.TrimSpace(child.Attrs["uml-partition-id"])
		title := strings.TrimSpace(child.Attrs["uml-partition-title"])
		if id == "" || title == "" || seen[id] {
			continue
		}
		seen[id] = true
		partitions = append(partitions, umlActivityPartitionV1EngineSceneWalk{id: id, title: title})
	}
	return partitions
}

func isXaligoActivityShapeV1EngineSceneWalk(box *entity.Box) bool {
	return box != nil && box.Tag == "rectangle" && box.Attrs["uml-diagram-kind"] == "activity-diagram"
}

func isXaligoClassShapeV1EngineSceneWalk(box *entity.Box) bool {
	return box != nil && box.Tag == "rectangle" && box.Attrs["uml-diagram-kind"] == "class-diagram"
}

func isXaligoComponentDiagramShapeV1EngineSceneWalk(box *entity.Box) bool {
	return box != nil && box.Tag == "rectangle" && box.Attrs["uml-diagram-kind"] == "component-diagram"
}

func isXaligoComponentShapeV1EngineSceneWalk(box *entity.Box) bool {
	return isXaligoComponentDiagramShapeV1EngineSceneWalk(box) && strings.TrimSpace(box.Attrs["uml-element-kind"]) == "component"
}

func isXaligoStateMachineShapeV1EngineSceneWalk(box *entity.Box) bool {
	return box != nil && box.Tag == "rectangle" && box.Attrs["uml-diagram-kind"] == "state-machine-diagram"
}

func isXaligoStateMachineStateShapeV1EngineSceneWalk(box *entity.Box) bool {
	return isXaligoStateMachineShapeV1EngineSceneWalk(box) && strings.TrimSpace(box.Attrs["uml-element-kind"]) == "state"
}

func isXaligoSequenceLifelineV1EngineSceneWalk(box *entity.Box) bool {
	if box == nil || box.Tag != "rectangle" || box.Attrs["uml-diagram-kind"] != "sequence-diagram" {
		return false
	}
	switch strings.TrimSpace(box.Attrs["uml-element-kind"]) {
	case "participant", "lifeline":
		return true
	default:
		return false
	}
}

func appendUMLSequenceLifelineV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, endpointRects map[string][4]float64, endpointIDs map[string]string, updated int64) {
	if box == nil || elements == nil {
		return
	}
	label := strings.TrimSpace(box.Label)
	if label == "" {
		label = strings.TrimSpace(box.Attrs["uml-local-id"])
	}
	headerHeight := math.Min(30, math.Max(24, box.H*0.12))
	headerWidth := math.Min(math.Max(92, textWidthV1EngineSceneItem(label, 7.5)+24), math.Max(1, box.W-16))
	headerX := box.X + (box.W-headerWidth)/2
	headerY := box.Y
	centerX := box.X + box.W/2
	lineY := headerY + headerHeight + 8
	lineHeight := math.Max(1, box.Y+box.H-lineY)
	owner := strings.TrimSpace(box.Attrs["uml-ref"])
	if owner == "" {
		owner = strings.TrimSpace(box.Attrs["ref"])
	}
	headerID := fmt.Sprintf("%s-sequence-header", box.ID)
	headerCustomData := umlShapeCustomDataV1EngineSceneWalk(box)
	if headerCustomData == nil {
		headerCustomData = map[string]any{}
	}
	headerCustomData["xaligoUmlSequenceLifelineHeader"] = true
	headerCustomData["xaligoUmlSequenceLifelineOwner"] = owner
	*elements = append(*elements, map[string]any{
		"id": headerID, "type": "rectangle",
		"x": headerX, "y": headerY, "width": headerWidth, "height": headerHeight,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#08b8ea",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": map[string]any{"type": 3},
		"seed": stableSceneSeedV1EngineSceneTypes(headerID), "version": 1, "versionNonce": stableSceneSeedV1EngineSceneTypes(headerID),
		"isDeleted": false, "boundElements": []map[string]any{{"type": "text", "id": fmt.Sprintf("%s-sequence-title", box.ID)}},
		"updated": updated, "link": nil, "locked": false,
		"customData": headerCustomData,
	})
	titleID := fmt.Sprintf("%s-sequence-title", box.ID)
	*elements = append(*elements, map[string]any{
		"id": titleID, "type": "text",
		"x": headerX + 4, "y": headerY + 2, "width": math.Max(1, headerWidth-8), "height": math.Max(1, headerHeight-4),
		"angle": 0, "strokeColor": "#ffffff", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": stableSceneSeedV1EngineSceneTypes(titleID), "version": 1, "versionNonce": stableSceneSeedV1EngineSceneTypes(titleID),
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"text": label, "fontSize": attrFloatV1EngineLayoutAttributes(box.Attrs["font-size"], 14), "fontFamily": fontFamilyV1EngineSceneWalk(box.Attrs["font-family"]),
		"textAlign": "center", "verticalAlign": "middle",
		"containerId": headerID, "originalText": label, "lineHeight": 1.2,
		"customData": map[string]any{"xaligoUmlSequenceLifelineHeaderContent": true, "xaligoTextLayout": sceneTextLayoutV1EngineSceneBuild(entity.TextRoleLabel, true, 1.2)},
	})
	lineID := fmt.Sprintf("%s-sequence-lifeline", box.ID)
	lineCustomData := map[string]any{
		"xaligoUmlDiagramKind":                  "sequence-diagram",
		"xaligoUmlElementKind":                  "lifeline-axis",
		"xaligoUmlSequenceLifeline":             true,
		"xaligoUmlSequenceLifelineOwner":        owner,
		"xaligoUmlRelationDestinationReference": owner,
	}
	*elements = append(*elements, map[string]any{
		"id": lineID, "type": "line",
		"x": centerX, "y": lineY, "width": 0, "height": lineHeight,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "dashed",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": stableSceneSeedV1EngineSceneTypes(lineID), "version": 1, "versionNonce": stableSceneSeedV1EngineSceneTypes(lineID),
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"points":     [][]float64{{0, 0}, {0, lineHeight}},
		"customData": lineCustomData,
	})
	registerConnectionEndpointV1EngineSceneWalk(box, lineID, [4]float64{centerX - 3, lineY, 6, lineHeight}, endpointRects, endpointIDs)
}

func isUMLActivityContainerV1EngineSceneWalk(box *entity.Box) bool {
	return box != nil && box.Tag == "uml" && box.Attrs["uml-kind"] == "activity-diagram"
}

func isUMLContainerV1EngineSceneWalk(box *entity.Box) bool {
	return box != nil && box.Tag == "uml" && strings.TrimSpace(box.Attrs["uml-kind"]) != ""
}

func umlActivityShapeBackgroundV1EngineSceneWalk(box *entity.Box) string {
	if strings.TrimSpace(box.Attrs["tone"]) == "primary" {
		return "#08b8ea"
	}
	switch strings.TrimSpace(box.Attrs["uml-element-kind"]) {
	case "initial":
		return "#052d6e"
	case "final":
		return "#ffffff"
	case "object-node":
		return "#e6fbf7"
	default:
		return "#e8f7fd"
	}
}

func umlClassShapeBackgroundV1EngineSceneWalk(box *entity.Box) string {
	return "#ffffff"
}

func umlStateMachineShapeBackgroundV1EngineSceneWalk(box *entity.Box) string {
	if configured := strings.TrimSpace(box.Attrs["background-color"]); configured != "" {
		return configured
	}
	switch strings.TrimSpace(box.Attrs["uml-element-kind"]) {
	case "initial":
		return "#052d6e"
	default:
		return "#ffffff"
	}
}

func appendUMLStateMachineCompartmentsV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64) {
	fontSize := attrFloatV1EngineLayoutAttributes(box.Attrs["font-size"], 14)
	headerH := math.Min(math.Max(30, fontSize*2.4), math.Max(30, box.H*0.38))
	headerID := fmt.Sprintf("%s-state-header", box.ID)
	headerSeed := stableSceneSeedV1EngineSceneTypes(headerID)
	*elements = append(*elements, map[string]any{
		"id": headerID, "type": "rectangle",
		"x": box.X, "y": box.Y, "width": box.W, "height": headerH,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#08b8ea",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": headerSeed, "version": 1, "versionNonce": headerSeed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"customData": map[string]any{"xaligoUmlStateHeader": true},
	})
	headerText := strings.TrimSpace(box.Attrs["uml-state-header-text"])
	if headerText == "" {
		headerText = strings.Split(strings.TrimSpace(box.Label), "\n")[0]
	}
	if headerText != "" {
		appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-state-header-text", box.ID), box.X+6, box.Y+4, math.Max(1, box.W-12), math.Max(1, headerH-8), headerText, "#ffffff", "center", "middle", fontSize, box, updated, map[string]any{"xaligoUmlStateHeaderContent": true})
	}
	dividerY := box.Y + headerH
	appendUMLStateMachineDividerV1EngineSceneWalk(elements, fmt.Sprintf("%s-state-header-divider", box.ID), box.X, dividerY, box.W, 0, updated, map[string]any{"xaligoUmlStateHeaderDivider": true})
	keys := splitNonEmptyLinesV1EngineSceneWalk(box.Attrs["uml-state-compartment-keys"])
	values := splitNonEmptyLinesV1EngineSceneWalk(box.Attrs["uml-state-compartment-values"])
	if len(keys) == 0 || len(values) == 0 {
		return
	}
	rowCount := len(keys)
	if len(values) < rowCount {
		rowCount = len(values)
	}
	bodyH := math.Max(1, box.H-headerH)
	rowH := bodyH / float64(rowCount)
	hideKeys := box.Attrs["uml-state-compartment-keys-hidden"] == "true"
	keyW := 0.0
	if !hideKeys {
		keyW = math.Min(76, math.Max(66, box.W*0.4))
		appendUMLStateMachineDividerV1EngineSceneWalk(elements, fmt.Sprintf("%s-state-column-divider", box.ID), box.X+keyW, dividerY, 0, bodyH, updated, map[string]any{"xaligoUmlStateColumnDivider": true})
	}
	for index := 0; index < rowCount; index++ {
		rowY := dividerY + float64(index)*rowH
		if index > 0 {
			appendUMLStateMachineDividerV1EngineSceneWalk(elements, fmt.Sprintf("%s-state-row-divider-%d", box.ID, index), box.X, rowY, box.W, 0, updated, map[string]any{"xaligoUmlStateRowDivider": true})
		}
		if !hideKeys {
			appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-state-key-%d", box.ID, index), box.X+4, rowY+2, math.Max(1, keyW-8), math.Max(1, rowH-4), keys[index], "#052d6e", "center", "middle", math.Min(fontSize, 12), box, updated, map[string]any{"xaligoUmlStateCompartmentKey": true})
		}
		appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-state-value-%d", box.ID, index), box.X+keyW+6, rowY+2, math.Max(1, box.W-keyW-12), math.Max(1, rowH-4), values[index], "#052d6e", "left", "middle", math.Min(fontSize, 12), box, updated, map[string]any{"xaligoUmlStateCompartmentValue": true})
	}
}

func appendUMLStateMachineDividerV1EngineSceneWalk(elements *[]map[string]any, id string, x, y, w, h float64, updated int64, customData map[string]any) {
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "line",
		"x": x, "y": y, "width": w, "height": h,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 70,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"points":     [][]float64{{0, 0}, {w, h}},
		"customData": customData,
	})
}

func splitNonEmptyLinesV1EngineSceneWalk(value string) []string {
	var result []string
	for _, line := range strings.Split(value, "\n") {
		if line = strings.TrimSpace(line); line != "" {
			result = append(result, line)
		}
	}
	return result
}

func appendUMLClassCompartmentsV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64) {
	fontSize := attrFloatV1EngineLayoutAttributes(box.Attrs["font-size"], 14)
	lineHeight := fontSize * 1.2
	headerText, attributeText, operationText, attributeLines, operationLines := umlClassTextSectionsV1EngineSceneWalk(box)
	headerH := umlClassHeaderHeightV1EngineSceneWalk(box, lineHeight, headerText)
	headerID := fmt.Sprintf("%s-class-header", box.ID)
	headerSeed := stableSceneSeedV1EngineSceneTypes(headerID)
	*elements = append(*elements, map[string]any{
		"id": headerID, "type": "rectangle",
		"x": box.X, "y": box.Y, "width": box.W, "height": headerH,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#08b8ea",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": headerSeed, "version": 1, "versionNonce": headerSeed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"customData": map[string]any{"xaligoUmlClassHeader": true},
	})
	appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-class-header-text", box.ID), box.X+6, box.Y+4, math.Max(1, box.W-12), math.Max(1, headerH-8), headerText, "#ffffff", "center", "middle", fontSize, box, updated, map[string]any{"xaligoUmlClassHeaderContent": true})
	appendUMLClassHeaderDividerV1EngineSceneWalk(box, elements, updated, box.Y+headerH)
	if attributeText == "" && operationText == "" {
		return
	}
	bodyY := box.Y + headerH + 7
	bottomY := box.Y + box.H - 6
	textX := box.X + 10
	textW := math.Max(1, box.W-20)
	if attributeLines > 0 && operationLines > 0 {
		attributeH := math.Max(lineHeight, float64(attributeLines)*lineHeight+6)
		dividerY := bodyY + attributeH
		appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-class-attribute-text", box.ID), textX, bodyY, textW, math.Max(1, attributeH-3), attributeText, "#052d6e", "left", "top", fontSize, box, updated, map[string]any{"xaligoUmlClassAttributeContent": true})
		appendUMLClassBodyDividerV1EngineSceneWalk(box, elements, updated, dividerY)
		operationY := dividerY + 6
		appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-class-operation-text", box.ID), textX, operationY, textW, math.Max(1, bottomY-operationY), operationText, "#052d6e", "left", "top", fontSize, box, updated, map[string]any{"xaligoUmlClassOperationContent": true})
		return
	}
	bodyText := attributeText
	customData := map[string]any{"xaligoUmlClassAttributeContent": true}
	if bodyText == "" {
		bodyText = operationText
		customData = map[string]any{"xaligoUmlClassOperationContent": true}
	}
	appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-class-body-text", box.ID), textX, bodyY, textW, math.Max(1, bottomY-bodyY), bodyText, "#052d6e", "left", "top", fontSize, box, updated, customData)
}

func appendUMLClassHeaderDividerV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64, y float64) {
	id := fmt.Sprintf("%s-class-divider", box.ID)
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "line",
		"x": box.X, "y": y, "width": box.W, "height": 0,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 70,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"points":     [][]float64{{0, 0}, {box.W, 0}},
		"customData": map[string]any{"xaligoUmlClassHeaderDivider": true},
	})
}

func appendUMLClassBodyDividerV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64, y float64) {
	id := fmt.Sprintf("%s-class-body-divider", box.ID)
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "line",
		"x": box.X, "y": y, "width": box.W, "height": 0,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 70,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"points":     [][]float64{{0, 0}, {box.W, 0}},
		"customData": map[string]any{"xaligoUmlClassBodyDivider": true},
	})
}

func appendUMLClassTextV1EngineSceneWalk(elements *[]map[string]any, id string, x, y, w, h float64, text, color, align, valign string, fontSize float64, box *entity.Box, updated int64, customData map[string]any) {
	if strings.TrimSpace(text) == "" {
		return
	}
	customData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(entity.TextRoleLabel, true, 1.2)
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "text",
		"x": x, "y": y, "width": w, "height": h,
		"angle": 0, "strokeColor": color, "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"text": text, "fontSize": fontSize, "fontFamily": fontFamilyV1EngineSceneWalk(box.Attrs["font-family"]),
		"textAlign": align, "verticalAlign": valign,
		"containerId": fmt.Sprintf("%s-rect", box.ID), "originalText": text, "lineHeight": 1.2,
		"customData": customData,
	})
}

func umlClassTextSectionsV1EngineSceneWalk(box *entity.Box) (string, string, string, int, int) {
	lines := strings.Split(box.Label, "\n")
	headerLines := int(attrFloatV1EngineLayoutAttributes(box.Attrs["uml-class-header-lines"], 1))
	attributeLines := int(attrFloatV1EngineLayoutAttributes(box.Attrs["uml-class-attribute-lines"], 0))
	operationLines := int(attrFloatV1EngineLayoutAttributes(box.Attrs["uml-class-operation-lines"], 0))
	if headerLines < 1 || headerLines > len(lines) {
		headerLines = 1
	}
	headerText := strings.Join(lines[:headerLines], "\n")
	bodyLines := lines[headerLines:]
	attributeText := strings.TrimSpace(box.Attrs["uml-class-attribute-text"])
	operationText := strings.TrimSpace(box.Attrs["uml-class-operation-text"])
	if attributeText == "" && operationText == "" {
		attributeText = strings.Join(bodyLines, "\n")
	}
	if attributeLines+operationLines > len(bodyLines) {
		attributeLines = len(bodyLines)
		operationLines = 0
	}
	return headerText, attributeText, operationText, attributeLines, operationLines
}

func umlClassHeaderHeightV1EngineSceneWalk(box *entity.Box, lineHeight float64, headerText string) float64 {
	lines := strings.Count(headerText, "\n") + 1
	return math.Min(math.Max(34, 10+float64(lines)*lineHeight), math.Max(34, box.H*0.48))
}

func isXaligoFinalWithDotV1EngineSceneWalk(box *entity.Box) bool {
	if strings.TrimSpace(box.Attrs["uml-element-kind"]) != "final" {
		return false
	}
	switch strings.TrimSpace(box.Attrs["uml-diagram-kind"]) {
	case "activity-diagram", "state-machine-diagram":
		return true
	default:
		return false
	}
}

func appendUMLFinalDotV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64) {
	diameter := math.Max(8, math.Min(box.W, box.H)*0.46)
	x := box.X + (box.W-diameter)/2
	y := box.Y + (box.H-diameter)/2
	id := fmt.Sprintf("%s-final-dot", box.ID)
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "ellipse",
		"x": x, "y": y, "width": diameter, "height": diameter,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#052d6e",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"customData": map[string]any{"xaligoUmlFinalDot": true},
	})
}

func umlShapeTextColorV1EngineSceneWalk(box *entity.Box) string {
	if isXaligoActivityShapeV1EngineSceneWalk(box) && strings.TrimSpace(box.Attrs["tone"]) == "primary" {
		return "#ffffff"
	}
	if isXaligoClassShapeV1EngineSceneWalk(box) && strings.TrimSpace(box.Attrs["tone"]) == "primary" {
		return "#ffffff"
	}
	if isXaligoActivityShapeV1EngineSceneWalk(box) {
		return "#052d6e"
	}
	if isXaligoClassShapeV1EngineSceneWalk(box) {
		return "#052d6e"
	}
	if isXaligoComponentDiagramShapeV1EngineSceneWalk(box) {
		return "#052d6e"
	}
	if isXaligoStateMachineShapeV1EngineSceneWalk(box) {
		return "#052d6e"
	}
	return tableTextColorV1EngineSceneWalk(box)
}

func appendUMLComponentHeaderV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64) {
	fontSize := attrFloatV1EngineLayoutAttributes(box.Attrs["font-size"], 14)
	headerH := math.Min(math.Max(30, fontSize*2.4), math.Max(30, box.H*0.38))
	headerID := fmt.Sprintf("%s-component-header", box.ID)
	headerSeed := stableSceneSeedV1EngineSceneTypes(headerID)
	*elements = append(*elements, map[string]any{
		"id": headerID, "type": "rectangle",
		"x": box.X, "y": box.Y, "width": box.W, "height": headerH,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#08b8ea",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": headerSeed, "version": 1, "versionNonce": headerSeed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"customData": map[string]any{"xaligoUmlComponentHeader": true},
	})
	if headerText := strings.TrimSpace(strings.Split(box.Label, "\n")[0]); headerText != "" {
		appendUMLClassTextV1EngineSceneWalk(elements, fmt.Sprintf("%s-component-header-text", box.ID), box.X+10, box.Y+4, math.Max(1, box.W-20), math.Max(1, headerH-8), headerText, "#ffffff", "left", "middle", fontSize, box, updated, map[string]any{"xaligoUmlComponentHeaderContent": true})
	}
	appendUMLClassHeaderDividerV1EngineSceneWalk(box, elements, updated, box.Y+headerH)
}

func appendUMLComponentBoundaryInterfacesV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64) {
	interfaces := splitUMLComponentInterfaceLabelsV1EngineSceneWalk(box.Attrs["uml-component-interfaces"])
	appendUMLComponentInterfacesV1EngineSceneWalk(box, elements, updated, interfaces)
}

func splitUMLComponentInterfaceLabelsV1EngineSceneWalk(value string) []string {
	var labels []string
	for _, line := range strings.Split(value, "\n") {
		if label := strings.TrimSpace(line); label != "" {
			labels = append(labels, label)
		}
	}
	return labels
}

func appendUMLComponentInterfacesV1EngineSceneWalk(box *entity.Box, elements *[]map[string]any, updated int64, labels []string) {
	if len(labels) == 0 {
		return
	}
	diameter := math.Min(18, math.Max(12, box.H*0.12))
	stem := math.Min(5, math.Max(3, box.W*0.015))
	portW := math.Min(72, math.Max(52, box.W*0.22))
	portH := math.Min(22, math.Max(18, box.H*0.17))
	for index, label := range labels {
		cy := umlComponentInterfaceYV1EngineSceneWalk(box, len(labels), index, 0.36, portH)
		portX := box.X - portW*0.12
		portY := cy - portH/2
		portID := fmt.Sprintf("%s-interface-%d-port", box.ID, index)
		portSeed := stableSceneSeedV1EngineSceneTypes(portID)
		*elements = append(*elements, map[string]any{
			"id": portID, "type": "rectangle",
			"x": portX, "y": portY, "width": portW, "height": portH,
			"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#ffffff",
			"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": portSeed, "version": 1, "versionNonce": portSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"customData": umlComponentInterfaceCustomDataV1EngineSceneWalk(box, label, map[string]any{"xaligoUmlComponentInterfacePort": true}),
		})
		appendUMLComponentInterfacePortLabelV1EngineSceneWalk(elements, fmt.Sprintf("%s-interface-%d-port-label", box.ID, index), portX+3, cy, portW-6, portH-4, label, box, updated, map[string]any{"xaligoUmlComponentInterfacePortLabel": true})

		circleX := portX - stem - diameter
		circleY := cy - diameter/2
		circleID := fmt.Sprintf("%s-interface-%d", box.ID, index)
		circleSeed := stableSceneSeedV1EngineSceneTypes(circleID)
		*elements = append(*elements, map[string]any{
			"id": circleID, "type": "ellipse",
			"x": circleX, "y": circleY, "width": diameter, "height": diameter,
			"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "#ffffff",
			"fillStyle": "solid", "strokeWidth": 1.2, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": nil,
			"seed": circleSeed, "version": 1, "versionNonce": circleSeed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false,
			"customData": umlComponentInterfaceCustomDataV1EngineSceneWalk(box, label, map[string]any{"xaligoUmlComponentInterfaceSymbol": true, "xaligoUmlComponentInterfaceCircle": true}),
		})
		appendUMLComponentInterfaceLineV1EngineSceneWalk(elements, fmt.Sprintf("%s-interface-%d-stem", box.ID, index), circleX+diameter, cy, portX, cy, updated, map[string]any{"xaligoUmlComponentInterfaceStem": true})
	}
}

func umlComponentInterfaceCustomDataV1EngineSceneWalk(box *entity.Box, label string, customData map[string]any) map[string]any {
	customData["xaligoUmlComponentInterfaceLabel"] = label
	customData["xaligoUmlComponentOwnerConnectionKey"] = strings.TrimSpace(box.Attrs[internalConnectionKeyAttrV1EngineParseDocument])
	customData["xaligoUmlComponentOwnerLocalId"] = strings.TrimSpace(box.Attrs["uml-local-id"])
	return customData
}

func umlComponentInterfaceYV1EngineSceneWalk(box *entity.Box, count, index int, base, portH float64) float64 {
	if count < 1 {
		count = 1
	}
	fontSize := attrFloatV1EngineLayoutAttributes(box.Attrs["font-size"], 14)
	headerH := math.Min(math.Max(30, fontSize*2.4), math.Max(30, box.H*0.38))
	minCenter := box.Y + headerH + math.Max(8, fontSize*0.7) + portH/2
	if count == 1 {
		return math.Max(box.Y+box.H*base, minCenter)
	}
	minStep := math.Min(34, math.Max(26, box.H*0.22))
	span := minStep * float64(count-1)
	top := box.Y + box.H*base - span/2
	maxBottom := box.Y + box.H - math.Max(16, box.H*0.12)
	if top < minCenter {
		top = minCenter
	}
	if bottom := top + span; bottom > maxBottom {
		top = math.Max(minCenter, maxBottom-span)
	}
	return top + minStep*float64(index)
}

func appendUMLComponentInterfaceLineV1EngineSceneWalk(elements *[]map[string]any, id string, x1, y1, x2, y2 float64, updated int64, customData map[string]any) {
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "line",
		"x": x1, "y": y1, "width": x2 - x1, "height": y2 - y1,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1.2, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"points":     [][]float64{{0, 0}, {x2 - x1, y2 - y1}},
		"customData": customData,
	})
}

func appendUMLComponentInterfaceLabelV1EngineSceneWalk(elements *[]map[string]any, id string, x, centerY, width float64, label, align string, box *entity.Box, updated int64, customData map[string]any) {
	fontSize := math.Max(11, attrFloatV1EngineLayoutAttributes(box.Attrs["font-size"], 14)-2)
	height := math.Max(14, fontSize*1.4)
	customData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(entity.TextRoleLabel, true, 1.2)
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "text",
		"x": x, "y": centerY - height/2, "width": width, "height": height,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"text": label, "fontSize": fontSize, "fontFamily": fontFamilyV1EngineSceneWalk(box.Attrs["font-family"]),
		"textAlign": align, "verticalAlign": "middle",
		"containerId": nil, "originalText": label, "lineHeight": 1.2,
		"customData": customData,
	})
}

func appendUMLComponentInterfacePortLabelV1EngineSceneWalk(elements *[]map[string]any, id string, x, centerY, width, height float64, label string, box *entity.Box, updated int64, customData map[string]any) {
	fontSize := math.Min(9, math.Max(7, attrFloatV1EngineLayoutAttributes(box.Attrs["font-size"], 14)-5))
	customData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(entity.TextRoleLabel, true, 1.0)
	seed := stableSceneSeedV1EngineSceneTypes(id)
	*elements = append(*elements, map[string]any{
		"id": id, "type": "text",
		"x": x, "y": centerY - height/2, "width": width, "height": height,
		"angle": 0, "strokeColor": "#052d6e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
		"text": label, "fontSize": fontSize, "fontFamily": fontFamilyV1EngineSceneWalk(box.Attrs["font-family"]),
		"textAlign": "center", "verticalAlign": "middle",
		"containerId": nil, "originalText": label, "lineHeight": 1.0,
		"customData": customData,
	})
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
		"uml-stereotype":        "xaligoUmlStereotype",
		"uml-abstract":          "xaligoUmlAbstract",
		"uml-static":            "xaligoUmlStatic",
		"uml-partition-id":      "xaligoUmlPartitionId",
		"uml-partition-title":   "xaligoUmlPartitionTitle",
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
