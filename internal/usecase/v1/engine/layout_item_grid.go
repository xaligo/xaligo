package engine

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

// resolvedItemGrid is the shared geometry chosen for a set of item slots.
// Scene construction may supply a catalog-derived label height, while Build
// uses the minimum label box to reject source-only failures before encoding.
type resolvedItemGridV1EngineLayoutItemGrid struct {
	Cols     int
	Rows     int
	IconSize float64
	CellW    float64
	CellH    float64
	StartX   float64
	StartY   float64
	StepX    float64
	StepY    float64
}

func (grid resolvedItemGridV1EngineLayoutItemGrid) iconPositionV1EngineLayoutItemGrid(index int) (float64, float64) {
	col := index % grid.Cols
	row := index / grid.Cols
	x := grid.StartX + float64(col)*grid.StepX + math.Max(0, (grid.CellW-grid.IconSize)/2)
	y := grid.StartY + float64(row)*grid.StepY
	return x, y
}

func resolveItemGridV1EngineLayoutItemGrid(count int, ancestor *entity.Box, maxSize, labelBoxH float64) (resolvedItemGridV1EngineLayoutItemGrid, error) {
	if ancestor == nil {
		return resolvedItemGridV1EngineLayoutItemGrid{}, fmt.Errorf("item grid has no containing box")
	}
	if count <= 0 {
		return resolvedItemGridV1EngineLayoutItemGrid{}, nil
	}
	if !isPositiveFiniteV1EngineLayoutConstraints(maxSize) || !isPositiveFiniteV1EngineLayoutConstraints(labelBoxH) {
		return resolvedItemGridV1EngineLayoutItemGrid{}, newResolvedLayoutErrorV1EngineLayoutValidation(ancestor, "item grid sizes must be positive and finite")
	}

	areaX, areaY, areaW, areaH := itemGridAreaV1EngineLayoutItemGrid(ancestor)
	if !isPositiveFiniteV1EngineLayoutConstraints(areaW) || !isPositiveFiniteV1EngineLayoutConstraints(areaH) {
		return resolvedItemGridV1EngineLayoutItemGrid{}, newResolvedLayoutErrorV1EngineLayoutValidation(ancestor, "item grid has no usable area (%.6gx%.6g)", areaW, areaH)
	}
	cols, rows, iconSize := chooseItemGridV1EngineSceneItem(count, areaW, areaH, maxSize, labelBoxH)
	if cols <= 0 || rows <= 0 || !isPositiveFiniteV1EngineLayoutConstraints(iconSize) {
		return resolvedItemGridV1EngineLayoutItemGrid{}, newResolvedLayoutErrorV1EngineLayoutValidation(ancestor, "cannot fit %d item slots in %.6gx%.6g", count, areaW, areaH)
	}
	cellW := math.Max(itemLabelWV1EngineSceneTypes, iconSize+itemAnchorGridVisualPadPxV1EngineSceneTypes*2)
	cellH := iconSize + 4 + labelBoxH
	totalW := cellW*float64(cols) + itemGapV1EngineSceneTypes*float64(cols-1)
	totalH := cellH*float64(rows) + itemGapV1EngineSceneTypes*float64(rows-1)
	if totalW > areaW+geometryEpsilonV1EngineLayoutValidation || totalH > areaH+geometryEpsilonV1EngineLayoutValidation {
		return resolvedItemGridV1EngineLayoutItemGrid{}, newResolvedLayoutErrorV1EngineLayoutValidation(ancestor, "cannot fit %d item slots in %.6gx%.6g; requires %.6gx%.6g", count, areaW, areaH, totalW, totalH)
	}
	vert, horiz := parseItemAlignV1EngineSceneItem(ancestor.Attrs["align"])
	startX, stepX := gridAxisV1EngineSceneItem(areaX, areaW, totalW, cellW, cols, horiz)
	startY, stepY := gridAxisV1EngineSceneItem(areaY, areaH, totalH, cellH, rows, vert)
	return resolvedItemGridV1EngineLayoutItemGrid{
		Cols: cols, Rows: rows, IconSize: iconSize,
		CellW: cellW, CellH: cellH,
		StartX: startX, StartY: startY, StepX: stepX, StepY: stepY,
	}, nil
}

func itemGridAreaV1EngineLayoutItemGrid(ancestor *entity.Box) (x, y, w, h float64) {
	if ancestor == nil {
		return 0, 0, 0, 0
	}
	if _, isGroup := awsGroupsV1EngineSceneTypes[ancestor.Tag]; isGroup {
		allItemChildren := true
		for _, child := range ancestor.Children {
			if !IsItemLikeV1EngineLayoutAttributes(child.Tag) {
				allItemChildren = false
				break
			}
		}
		if allItemChildren {
			topClearance := groupHeaderHeightForItemsV1EngineSceneItem(ancestor)/2 + itemAnchorGridVisualPadPxV1EngineSceneTypes + float64(groupHeaderBorderGapV1EngineSceneTypes)
			if topClearance < itemGapV1EngineSceneTypes {
				topClearance = itemGapV1EngineSceneTypes
			}
			return ancestor.X + GroupSideInsetV1EngineLayoutNode, ancestor.Y + topClearance,
				ancestor.W - GroupSideInsetV1EngineLayoutNode*2, ancestor.H - topClearance - itemGapV1EngineSceneTypes
		}
		return ancestor.X + GroupSideInsetV1EngineLayoutNode, ancestor.Y + GroupTopInsetV1EngineLayoutNode + itemGapV1EngineSceneTypes,
			ancestor.W - GroupSideInsetV1EngineLayoutNode*2, ancestor.H - GroupTopInsetV1EngineLayoutNode - itemGapV1EngineSceneTypes*2
	}
	if ancestor.Tag == "frame" && isPositiveFiniteV1EngineLayoutConstraints(ancestor.ContentW) && isPositiveFiniteV1EngineLayoutConstraints(ancestor.ContentH) {
		return ancestor.ContentX + itemGapV1EngineSceneTypes, ancestor.ContentY + itemGapV1EngineSceneTypes,
			ancestor.ContentW - itemGapV1EngineSceneTypes*2, ancestor.ContentH - itemGapV1EngineSceneTypes*2
	}
	return ancestor.X + itemGapV1EngineSceneTypes, ancestor.Y + itemGapV1EngineSceneTypes,
		ancestor.W - itemGapV1EngineSceneTypes*2, ancestor.H - itemGapV1EngineSceneTypes*2
}

// validateResolvedItemGrids catches source-only item geometry failures in the
// same Build stage used by validate and render. Catalog-derived label metrics
// are checked again by resolveItemGrid during scene construction.
func validateResolvedItemGridsV1EngineLayoutItemGrid(root *entity.Box) error {
	if root == nil {
		return nil
	}
	groups, ancestors, frames := collectResolvedItemGroupsV1EngineLayoutItemGrid(root)
	maxSize := itemMaxSizeV1EngineSceneTypes
	if value := strings.TrimSpace(root.Attrs["item-size"]); value != "" {
		maxSize = attrFloatV1EngineLayoutAttributes(value, itemMaxSizeV1EngineSceneTypes)
	}
	keys := make([]string, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		items := groups[key]
		ancestor := ancestors[key]
		grid, err := resolveItemGridV1EngineLayoutItemGrid(len(items), ancestor, maxSize, itemLabelHV1EngineSceneTypes)
		if err != nil {
			return err
		}
		for index, item := range items {
			if strings.TrimSpace(item.Attrs["id"]) == "" {
				continue
			}
			x, y := grid.iconPositionV1EngineLayoutItemGrid(index)
			dx, dy, err := itemIconOffsetV1EngineSceneItem(item)
			if err != nil {
				return newResolvedLayoutErrorV1EngineLayoutValidation(item, "%v", err)
			}
			if err := validateItemVisualBoundsV1EngineSceneItem(item, ancestor, frames[key], x+dx, y+dy, grid.IconSize, itemLabelHV1EngineSceneTypes); err != nil {
				return err
			}
		}
	}
	return nil
}

func collectResolvedItemGroupsV1EngineLayoutItemGrid(root *entity.Box) (map[string][]*entity.Box, map[string]*entity.Box, map[string]*entity.Box) {
	groups := map[string][]*entity.Box{}
	ancestors := map[string]*entity.Box{}
	frames := map[string]*entity.Box{}
	var walk func(*entity.Box, *entity.Box, *entity.Box)
	walk = func(box, visibleAncestor, activeFrame *entity.Box) {
		if box == nil {
			return
		}
		if box.Tag == "frame" {
			activeFrame = box
		}
		if IsItemLikeV1EngineLayoutAttributes(box.Tag) {
			if visibleAncestor != nil {
				groups[visibleAncestor.ID] = append(groups[visibleAncestor.ID], box)
				ancestors[visibleAncestor.ID] = visibleAncestor
				frames[visibleAncestor.ID] = activeFrame
			}
			return
		}
		if box.IsStaggerBg {
			return
		}
		nextVisible := box
		if box.Attrs["visible"] == "false" {
			nextVisible = visibleAncestor
		}
		for _, child := range box.Children {
			walk(child, nextVisible, activeFrame)
		}
	}
	walk(root, root, nil)
	return groups, ancestors, frames
}
