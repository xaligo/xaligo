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

// textWidth estimates the rendered width of a string in pixels.
// charW: approximate pixel width per rune (font-specific).
func textWidthV1EngineSceneItem(s string, charW float64) float64 {
	return math.Ceil(displayColumnsV1EngineSceneItem(s)*charW) + 8
}

func itemLabelHeightV1EngineSceneItem(label string) float64 {
	lines := 1
	for _, line := range strings.Split(label, "\n") {
		wrapped := int(math.Ceil(displayColumnsV1EngineSceneItem(line) * itemLabelCharWV1EngineSceneTypes / itemLabelWV1EngineSceneTypes))
		if wrapped < 1 {
			wrapped = 1
		}
		lines += wrapped - 1
	}
	lineH := itemLabelFontPxV1EngineSceneTypes * 1.25
	return math.Max(itemLabelHV1EngineSceneTypes, math.Ceil(float64(lines)*lineH))
}

func displayColumnsV1EngineSceneItem(s string) float64 {
	cols := 0.0
	for _, r := range s {
		cols += runeColumnsV1EngineSceneItem(r)
	}
	return cols
}

func runeColumnsV1EngineSceneItem(r rune) float64 {
	switch {
	case r == '\t':
		return 4
	case r < 0x20:
		return 0
	case r >= 0x1100 && (r <= 0x115F ||
		r == 0x2329 || r == 0x232A ||
		(r >= 0x2E80 && r <= 0xA4CF && r != 0x303F) ||
		(r >= 0xAC00 && r <= 0xD7A3) ||
		(r >= 0xF900 && r <= 0xFAFF) ||
		(r >= 0xFE10 && r <= 0xFE19) ||
		(r >= 0xFE30 && r <= 0xFE6F) ||
		(r >= 0xFF00 && r <= 0xFF60) ||
		(r >= 0xFFE0 && r <= 0xFFE6)):
		return 2
	default:
		return 1
	}
}

// parseItemAlign parses an align attribute value (e.g. "top-left", "middle-center")
// into vertical ("top"/"middle"/"bottom") and horizontal ("left"/"center"/"right") parts.
// Defaults to "middle" / "center" when absent or unrecognised.
func parseItemAlignV1EngineSceneItem(align string) (vert, horiz string) {
	vert, horiz = "middle", "center"
	parts := strings.SplitN(strings.ToLower(strings.TrimSpace(align)), "-", 2)
	if len(parts) == 2 {
		if parts[0] == "top" || parts[0] == "middle" || parts[0] == "bottom" {
			vert = parts[0]
		}
		if parts[1] == "left" || parts[1] == "center" || parts[1] == "right" || parts[1] == "spread" {
			horiz = parts[1]
		}
	}
	return
}

// renderItemGrid lays out all items collected under the same visibleAncestor as
// a compact grid within the ancestor's content area.
func renderItemGridV1EngineSceneItem(items []*entity.Box, ancestor, metadataFrame *entity.Box, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, maxSize float64, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, abbrevMap map[int]string, deps SceneDependenciesV1EngineSceneTypes) error {
	if catalogCSV == "" || len(items) == 0 || ancestor == nil {
		return nil
	}
	nItems := len(items)
	labelBoxH := estimateMaxItemLabelHeightV1EngineSceneItem(items, catalogCSV, fsys, abbrevMap, deps)
	grid, err := resolveItemGridV1EngineLayoutItemGrid(nItems, ancestor, maxSize, labelBoxH)
	if err != nil {
		return err
	}

	for i, item := range items {
		iconX, iconY := grid.iconPositionV1EngineLayoutItemGrid(i)
		if strings.TrimSpace(item.Attrs["id"]) != "" {
			dx, dy, err := itemIconOffsetV1EngineSceneItem(item)
			if err != nil {
				return err
			}
			iconX += dx
			iconY += dy
			if err := validateItemVisualBoundsV1EngineSceneItem(item, ancestor, metadataFrame, iconX, iconY, grid.IconSize, labelBoxH); err != nil {
				return err
			}
		}
		connectionKey := strings.TrimSpace(item.Attrs[internalConnectionKeyAttrV1EngineParseDocument])
		if connectionKey == "" {
			connectionKey = item.ID
		}
		renderIconAtV1EngineSceneItem(item.ID, connectionKey, item.Attrs["id"], iconX, iconY, grid.IconSize, elements, files, catalogCSV, projectRoot, fsys, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps)
	}
	return nil
}

func itemIconOffsetV1EngineSceneItem(item *entity.Box) (float64, float64, error) {
	if item == nil {
		return 0, 0, nil
	}
	dx, err := parseOptionalFloatAttrV1EngineSceneItem(item, "dx")
	if err != nil {
		return 0, 0, err
	}
	dy, err := parseOptionalFloatAttrV1EngineSceneItem(item, "dy")
	if err != nil {
		return 0, 0, err
	}
	return dx, dy, nil
}

func parseOptionalFloatAttrV1EngineSceneItem(item *entity.Box, attr string) (float64, error) {
	value := strings.TrimSpace(item.Attrs[attr])
	if value == "" {
		return 0, nil
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("<item %s=%q> must be a number", attr, value)
	}
	return parsed, nil
}

func validateItemVisualBoundsV1EngineSceneItem(item, ancestor, metadataFrame *entity.Box, x, y, size, labelHeight float64) error {
	const epsilon = 1e-6
	if item == nil || ancestor == nil {
		return nil
	}
	if !isPositiveFiniteV1EngineLayoutConstraints(size) || !isPositiveFiniteV1EngineLayoutConstraints(labelHeight) || math.IsNaN(x) || math.IsInf(x, 0) || math.IsNaN(y) || math.IsInf(y, 0) {
		return newResolvedLayoutErrorV1EngineLayoutValidation(item, "item visual geometry must be finite with positive icon and label sizes")
	}
	minX := ancestor.X
	minY := ancestor.Y
	maxX := ancestor.X + ancestor.W
	maxY := ancestor.Y + ancestor.H
	if ancestor.Tag == "frame" && isPositiveFiniteV1EngineLayoutConstraints(ancestor.ContentW) && isPositiveFiniteV1EngineLayoutConstraints(ancestor.ContentH) {
		minX = ancestor.ContentX
		minY = ancestor.ContentY
		maxX = ancestor.ContentX + ancestor.ContentW
		maxY = ancestor.ContentY + ancestor.ContentH
	}
	visualLeft := math.Min(x-itemAnchorGridVisualPadPxV1EngineSceneTypes, x+(size-itemLabelWV1EngineSceneTypes)/2)
	visualTop := y - itemAnchorGridVisualPadPxV1EngineSceneTypes
	visualRight := math.Max(x+size+itemAnchorGridVisualPadPxV1EngineSceneTypes, x+(size+itemLabelWV1EngineSceneTypes)/2)
	visualBottom := math.Max(y+size+itemAnchorGridVisualPadPxV1EngineSceneTypes, y+size+4+labelHeight)
	if metadataFrame != nil && metadataFrame.FrameMetadata != nil {
		metadata := metadataFrame.FrameMetadata
		if metadata.ReservedW > 0 && metadata.ReservedH > 0 &&
			visualLeft < metadata.ReservedX+metadata.ReservedW-epsilon && visualRight > metadata.ReservedX+epsilon &&
			visualTop < metadata.ReservedY+metadata.ReservedH-epsilon && visualBottom > metadata.ReservedY+epsilon {
			return newResolvedLayoutErrorV1EngineLayoutValidation(item, "<item id=%q> icon, label, or connection anchor enters frame metadata reserved strip; overflow=\"visible\" cannot override this page-decoration exclusion zone", strings.TrimSpace(item.Attrs["id"]))
		}
	}
	if visualLeft+epsilon < minX || visualTop+epsilon < minY || visualRight > maxX+epsilon || visualBottom > maxY+epsilon {
		if ancestor.Overflow == entity.OverflowVisible {
			ancestor.Overflowed = true
			return nil
		}
		return newResolvedLayoutErrorV1EngineLayoutValidation(item, "<item id=%q> offset moves its icon, label, or connection anchor outside parent %q content bounds: visual=(%.1f,%.1f,%.1f,%.1f), parent=(%.1f,%.1f,%.1f,%.1f)",
			strings.TrimSpace(item.Attrs["id"]), ancestor.Tag, visualLeft, visualTop, visualRight-visualLeft, visualBottom-visualTop, minX, minY, maxX-minX, maxY-minY)
	}
	return nil
}

func groupHeaderHeightForItemsV1EngineSceneItem(ancestor *entity.Box) float64 {
	headerH := float64(groupTextHeightV1EngineSceneTypes + groupHeaderTextPadYV1EngineSceneTypes*2)
	if ancestor == nil {
		return headerH
	}
	if ancestor.Tag == "generic-group" && strings.TrimSpace(ancestor.Attrs["icon-id"]) != "" {
		return float64(groupIconSizeV1EngineSceneTypes)
	}
	if gd, ok := awsGroupsV1EngineSceneTypes[ancestor.Tag]; ok && gd.IconFile != "" {
		return float64(groupIconSizeV1EngineSceneTypes)
	}
	return headerH
}

func estimateMaxItemLabelHeightV1EngineSceneItem(items []*entity.Box, catalogCSV string, fsys fs.FS, abbrevMap map[int]string, deps SceneDependenciesV1EngineSceneTypes) float64 {
	maxH := itemLabelHV1EngineSceneTypes
	for _, item := range items {
		id, err := strconv.Atoi(strings.TrimSpace(item.Attrs["id"]))
		if err != nil {
			continue
		}
		label := ""
		if abbrevMap != nil {
			label = abbrevMap[id]
		}
		if label == "" {
			ce, lookupErr := deps.lookupCatalogByIDV1EngineSceneTypes(fsys, catalogCSV, id)
			err = lookupErr
			if err != nil {
				continue
			}
			label = entity.ItemShortName(ce.Service)
		}
		maxH = math.Max(maxH, itemLabelHeightV1EngineSceneItem(label))
	}
	return maxH
}

func chooseItemGridV1EngineSceneItem(n int, areaW, areaH, maxSize float64, labelBoxH float64) (cols int, rows int, iconSize float64) {
	if n <= 0 || areaW <= 0 || areaH <= 0 {
		return 0, 0, 0
	}
	bestScore := -1.0
	for c := 1; c <= n; c++ {
		r := int(math.Ceil(float64(n) / float64(c)))
		cellW := (areaW - itemGapV1EngineSceneTypes*float64(c-1)) / float64(c)
		cellH := (areaH - itemGapV1EngineSceneTypes*float64(r-1)) / float64(r)
		size := math.Min(cellW-itemAnchorGridVisualPadPxV1EngineSceneTypes*2, cellH-4-labelBoxH)
		size = math.Min(size, maxSize)
		if size < itemMinSizeV1EngineSceneTypes {
			continue
		}
		slotW := math.Max(itemLabelWV1EngineSceneTypes, size+itemAnchorGridVisualPadPxV1EngineSceneTypes*2)
		usedW := slotW*float64(c) + itemGapV1EngineSceneTypes*float64(c-1)
		usedH := (size+4+labelBoxH)*float64(r) + itemGapV1EngineSceneTypes*float64(r-1)
		if usedW-areaW > 1e-6 || usedH-areaH > 1e-6 {
			continue
		}
		aspectPenalty := math.Abs(float64(c)/float64(r) - areaW/math.Max(1, areaH))
		score := size*100 - aspectPenalty
		if score > bestScore {
			bestScore = score
			cols = c
			rows = r
			iconSize = size
		}
	}
	return cols, rows, iconSize
}

func gridAxisV1EngineSceneItem(areaStart, areaSize, totalSize, cellSize float64, count int, align string) (start, step float64) {
	if count <= 1 {
		return areaStart + math.Max(0, (areaSize-cellSize)/2), 0
	}
	switch align {
	case "left", "top":
		return areaStart, cellSize + itemGapV1EngineSceneTypes
	case "right", "bottom":
		return areaStart + math.Max(0, areaSize-totalSize), cellSize + itemGapV1EngineSceneTypes
	case "spread":
		gap := (areaSize - cellSize*float64(count)) / float64(count+1)
		if gap < itemGapV1EngineSceneTypes {
			gap = itemGapV1EngineSceneTypes
		}
		return areaStart + gap, cellSize + gap
	default:
		return areaStart + math.Max(0, (areaSize-totalSize)/2), cellSize + itemGapV1EngineSceneTypes
	}
}

// renderIconAt draws a single service icon (image + label) at an explicit position.
// itemImgRects/itemLblRects/itemImgIDs/itemLblIDs are populated with the bounding rect
// and element ID of the image and label elements, keyed by the unique item connection key.
func renderIconAtV1EngineSceneItem(boxID, connectionKey, idAttr string, iconX, iconY, iconSize float64, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, abbrevMap map[int]string, deps SceneDependenciesV1EngineSceneTypes) {
	renderIconAtWithLabelV1EngineSceneItem(boxID, connectionKey, idAttr, iconX, iconY, iconSize, elements, files, catalogCSV, projectRoot, fsys, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps, nil)
}

func renderIconAtWithLabelV1EngineSceneItem(boxID, connectionKey, idAttr string, iconX, iconY, iconSize float64, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, abbrevMap map[int]string, deps SceneDependenciesV1EngineSceneTypes, labelOverride *string) {
	renderIconWithTextBoxV1EngineSceneItem(boxID, connectionKey, idAttr, iconX, iconY, iconSize, elements, files, catalogCSV, projectRoot, fsys, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps, labelOverride, itemLabelWV1EngineSceneTypes)
}

func renderIconWithTextBoxV1EngineSceneItem(boxID, connectionKey, idAttr string, iconX, iconY, iconSize float64, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, abbrevMap map[int]string, deps SceneDependenciesV1EngineSceneTypes, labelOverride *string, labelWidth float64) {
	if catalogCSV == "" {
		return
	}
	idAttr = strings.TrimSpace(idAttr)
	if idAttr == "" {
		return
	}

	// 1:1 — id は単一の整数
	id, err := strconv.Atoi(idAttr)
	if err != nil {
		loggerV1EngineSharedLogging.WARN(IUESRIA001V1EngineSceneTypes, "item ID must be a single integer", map[string]any{"id": idAttr, "error": err})
		return
	}
	ce, err := deps.lookupCatalogByIDV1EngineSceneTypes(fsys, catalogCSV, id)
	if err != nil {
		loggerV1EngineSharedLogging.WARN(IUESRIA002V1EngineSceneTypes, "catalog lookup failed", map[string]any{"id": id, "error": err})
		return
	}
	if ce.DataURL == "" && ce.RelPath != "" && projectRoot != "" {
		svgPath := filepath.Join(projectRoot, ce.RelPath)
		if du, err2 := deps.svgToDataURLV1EngineSceneTypes(nil, svgPath); err2 == nil {
			ce.DataURL = du
		} else {
			loggerV1EngineSharedLogging.WARN(IUESRIA003V1EngineSceneTypes, "cannot load SVG", map[string]any{"id": id, "path": svgPath, "error": err2})
		}
	}
	if ce.DataURL == "" {
		return
	}
	ce.DataURL = normalizeItemSVGDataURLV1EngineSceneTypes(ce.DataURL)

	updated := excalidrawUpdatedV1EngineSceneTypes
	fid := fmt.Sprintf("item-cat-%d", id)
	if _, exists := files[fid]; !exists {
		files[fid] = map[string]any{
			"mimeType": "image/svg+xml", "id": fid,
			"dataURL": ce.DataURL,
			"created": updated, "lastRetrieved": updated,
		}
	}
	iconID := fmt.Sprintf("%s-item", boxID)
	seed := stableSceneSeedV1EngineSceneTypes(iconID)
	anchorGroupID := fmt.Sprintf("%s-anchor", boxID)
	label := ""
	if labelOverride != nil {
		label = strings.TrimSpace(*labelOverride)
	} else {
		if abbrevMap != nil {
			label = abbrevMap[id]
		}
		if label == "" {
			label = entity.ItemShortName(ce.Service)
		}
	}
	labelH := itemLabelHeightV1EngineSceneItem(label)
	if labelWidth != itemLabelWV1EngineSceneTypes {
		labelH = math.Ceil(float64(strings.Count(label, "\n")+1) * itemLabelFontPxV1EngineSceneTypes * 1.25)
	}
	labelY := iconY + iconSize + 4
	labelX := iconX + (iconSize-labelWidth)/2 // centre label on icon
	anchorX := iconX - excalidrawAnchorPadPxV1EngineSceneTypes
	anchorY := iconY - excalidrawAnchorPadPxV1EngineSceneTypes
	anchorW := iconSize + excalidrawAnchorPadPxV1EngineSceneTypes*2
	anchorH := iconSize + excalidrawAnchorPadPxV1EngineSceneTypes*2
	// Record bounding rects and element IDs for edge-based connection arrows.
	if itemImgRects != nil {
		itemImgRects[connectionKey] = [4]float64{iconX, iconY, iconSize, iconSize}
		itemImgIDs[connectionKey] = iconID
	}
	appendExcalidrawAnchorGridV1EngineSceneItem(elements, boxID, anchorGroupID, anchorX, anchorY, anchorW, anchorH, updated)
	*elements = append(*elements, map[string]any{
		"id": iconID, "type": "image",
		"x": iconX, "y": iconY,
		"width": iconSize, "height": iconSize,
		"fileId": fid, "status": "saved",
		"scale":       []int{1, 1},
		"strokeColor": "transparent", "backgroundColor": deps.svgBGColorV1EngineSceneTypes(ce.DataURL),
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100, "angle": 0,
		"groupIds": []string{anchorGroupID}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false, "frameId": nil,
		"customData": map[string]any{"xaligoAnchorContent": true},
	})
	if label == "" {
		return
	}
	// Record label bounding rect for bottom-side connection binding.
	if itemLblRects != nil {
		itemLblRects[connectionKey] = [4]float64{labelX, labelY, labelWidth, labelH}
		itemLblIDs[connectionKey] = iconID + "-lbl"
	}
	textSeed := stableSceneSeedV1EngineSceneTypes(iconID + "-lbl")
	*elements = append(*elements, map[string]any{
		"id": iconID + "-lbl", "type": "text",
		"x": labelX, "y": labelY,
		"width": labelWidth, "height": labelH,
		"angle":       0,
		"strokeColor": "#1e1e1e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{anchorGroupID}, "roundness": nil,
		"seed": textSeed, "version": 1, "versionNonce": textSeed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false, "frameId": nil,
		"text": label, "rawText": label, "originalText": label,
		"fontSize": itemLabelFontPxV1EngineSceneTypes, "fontFamily": 4,
		"textAlign": "center", "verticalAlign": "top",
		"containerId": nil, "lineHeight": 1.25,
		"customData": map[string]any{
			"xaligoAnchorContent": true,
			"xaligoTextLayout":    sceneTextLayoutV1EngineSceneBuild(entity.TextRoleItemLabel, true, 1.25),
		},
	})
}

func appendExcalidrawAnchorGridV1EngineSceneItem(elements *[]map[string]any, boxID, groupID string, x, y, w, h float64, updated int64) {
	if elements == nil || excalidrawAnchorGridV1EngineSceneTypes <= 0 {
		return
	}
	cellW := w / float64(excalidrawAnchorGridV1EngineSceneTypes)
	cellH := h / float64(excalidrawAnchorGridV1EngineSceneTypes)
	for row := 0; row < excalidrawAnchorGridV1EngineSceneTypes; row++ {
		for col := 0; col < excalidrawAnchorGridV1EngineSceneTypes; col++ {
			cellX := x + float64(col)*cellW + excalidrawAnchorCellGapPxV1EngineSceneTypes
			cellY := y + float64(row)*cellH + excalidrawAnchorCellGapPxV1EngineSceneTypes
			cellWidth := math.Max(1, cellW-excalidrawAnchorCellGapPxV1EngineSceneTypes*2)
			cellHeight := math.Max(1, cellH-excalidrawAnchorCellGapPxV1EngineSceneTypes*2)
			cellID := fmt.Sprintf("%s-anchor-bg-%02d-%02d", boxID, row, col)
			cellSeed := stableSceneSeedV1EngineSceneTypes(cellID)
			*elements = append(*elements, map[string]any{
				"id": cellID, "type": "rectangle",
				"x": cellX, "y": cellY,
				"width": cellWidth, "height": cellHeight,
				"angle":       0,
				"strokeColor": "#ffffff", "backgroundColor": "#ffffff",
				"fillStyle": "solid", "strokeWidth": 0, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{groupID}, "roundness": nil,
				"seed": cellSeed, "version": 1, "versionNonce": cellSeed,
				"isDeleted": false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false, "frameId": nil,
				"customData": map[string]any{"xaligoAnchorBackground": true},
			})
		}
	}
}
