package engine

import (
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

// MinBoxWidthV1EngineLayoutFlow / MinBoxHeightV1EngineLayoutFlow are the smallest dimensions at which a staggered
// box can be meaningfully rendered. A strict parent reports insufficient
// space; overflow="visible" preserves the explicit minimum-size opt-out.
const (
	MinBoxWidthV1EngineLayoutFlow  = 60.0
	MinBoxHeightV1EngineLayoutFlow = 48.0
)

var (
	IULS001V1EngineLayoutFlow  = share.NewMCode("IULS-001", "Layout stack empty children branch")
	IULFH001V1EngineLayoutFlow = share.NewMCode("IULFH-001", "Layout flex horizontal empty children branch")
	IULR001V1EngineLayoutFlow  = share.NewMCode("IULR-001", "Layout row empty children branch")
	IULST001V1EngineLayoutFlow = share.NewMCode("IULST-001", "Layout stagger fallback branch")
	IULST002V1EngineLayoutFlow = share.NewMCode("IULST-002", "Layout stagger minimum width branch")
	IULST003V1EngineLayoutFlow = share.NewMCode("IULST-003", "Layout stagger minimum height branch")
)

func layoutStackV1EngineLayoutFlow(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	children := layoutKidsV1EngineLayoutNode(node)
	if len(children) == 0 {
		loggerV1EngineSharedLogging.DEBUG(IULS001V1EngineLayoutFlow, "branch empty children", map[string]any{"tag": node.Tag})
		return nil
	}
	gap := attrFloatV1EngineLayoutAttributes(node.Attr("gap"), 16)

	// 各子要素の margin を事前に読み取り、縦方向の余白合計を算出する。
	// これにより margin が sibling 間スペースとして機能する (CSS ライク)。
	// row 属性は flex-grow スタイルの高さ比率。デフォルト 1.0 (均等)。
	totalMarginH := 0.0
	totalRow := 0.0
	totalFixedH := 0.0
	flexCount := 0
	for _, child := range children {
		childMar := effectiveMarginV1EngineLayoutAttributes(child)
		totalMarginH += childMar.Top + childMar.Bottom
		if fixedH, fixed := explicitSizeV1EngineLayoutConstraints(child, "height"); fixed {
			totalFixedH += fixedH
		} else {
			totalRow += attrFloatV1EngineLayoutAttributes(child.Attr("row"), 1.0)
			flexCount++
		}
	}
	availH := h - gap*float64(len(children)-1) - totalMarginH
	flexH, err := flexibleAxisSpaceV1EngineLayoutConstraints(node, "height", availH, h, totalFixedH, flexCount)
	if err != nil {
		return err
	}

	curY := y
	for i, child := range children {
		childMar := effectiveMarginV1EngineLayoutAttributes(child)
		childH, fixed := explicitSizeV1EngineLayoutConstraints(child, "height")
		if !fixed {
			row := attrFloatV1EngineLayoutAttributes(child.Attr("row"), 1.0)
			childH = flexH * (row / totalRow)
		}
		if !isPositiveFiniteV1EngineLayoutConstraints(childH) {
			return newLayoutErrorV1EngineLayoutValidation(child, "allocated height must be positive and finite, got %.6g", childH)
		}
		alloc := childH + childMar.Top + childMar.Bottom
		cb := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, i), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, cb, x, curY, w, alloc); err != nil {
			return err
		}
		target.Children = append(target.Children, cb)
		curY += alloc + gap
	}
	return nil
}

// layoutFlexH lays out children horizontally with free ratio weights.
// Each child's width share is determined by its `col` attribute (default 1.0).
// This mirrors layoutStack but in the horizontal direction.
func layoutFlexHV1EngineLayoutFlow(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	children := layoutKidsV1EngineLayoutNode(node)
	if len(children) == 0 {
		loggerV1EngineSharedLogging.DEBUG(IULFH001V1EngineLayoutFlow, "branch empty children", map[string]any{"tag": node.Tag})
		return nil
	}
	gap := attrFloatV1EngineLayoutAttributes(node.Attr("gap"), 16)

	// 各子要素の水平 margin を事前集計し利用可能幅を算出する。
	// col 属性は flex-grow スタイルの幅比率。デフォルト 1.0 (均等)。
	totalMarginW := 0.0
	totalCol := 0.0
	totalFixedW := 0.0
	flexCount := 0
	for _, child := range children {
		childMar := effectiveMarginV1EngineLayoutAttributes(child)
		totalMarginW += childMar.Left + childMar.Right
		if fixedW, fixed := explicitSizeV1EngineLayoutConstraints(child, "width"); fixed {
			totalFixedW += fixedW
		} else {
			totalCol += attrFloatV1EngineLayoutAttributes(child.Attr("col"), 1.0)
			flexCount++
		}
	}
	availW := w - gap*float64(len(children)-1) - totalMarginW
	flexW, err := flexibleAxisSpaceV1EngineLayoutConstraints(node, "width", availW, w, totalFixedW, flexCount)
	if err != nil {
		return err
	}

	curX := x
	for i, child := range children {
		childMar := effectiveMarginV1EngineLayoutAttributes(child)
		childW, fixed := explicitSizeV1EngineLayoutConstraints(child, "width")
		if !fixed {
			col := attrFloatV1EngineLayoutAttributes(child.Attr("col"), 1.0)
			childW = flexW * (col / totalCol)
		}
		if !isPositiveFiniteV1EngineLayoutConstraints(childW) {
			return newLayoutErrorV1EngineLayoutValidation(child, "allocated width must be positive and finite, got %.6g", childW)
		}
		alloc := childW + childMar.Left + childMar.Right
		cb := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, i), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, cb, curX, y, alloc, h); err != nil {
			return err
		}
		target.Children = append(target.Children, cb)
		curX += alloc + gap
	}
	return nil
}

func layoutRowV1EngineLayoutFlow(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	children := layoutKidsV1EngineLayoutNode(node)
	if len(children) == 0 {
		loggerV1EngineSharedLogging.DEBUG(IULR001V1EngineLayoutFlow, "branch empty children", map[string]any{"tag": node.Tag})
		return nil
	}
	gap := attrFloatV1EngineLayoutAttributes(node.Attr("gap"), 16)

	// 各子要素の水平 margin を事前に読み取り、幅方向の合計を算出する。
	totalMarginW := 0.0
	totalFixedW := 0.0
	flexCount := 0
	for _, child := range children {
		childMar := effectiveMarginV1EngineLayoutAttributes(child)
		totalMarginW += childMar.Left + childMar.Right
		if fixedW, fixed := explicitSizeV1EngineLayoutConstraints(child, "width"); fixed {
			totalFixedW += fixedW
		} else {
			flexCount++
		}
	}
	remainingW := w - gap*float64(len(children)-1) - totalMarginW
	flexW, err := flexibleAxisSpaceV1EngineLayoutConstraints(node, "width", remainingW, w, totalFixedW, flexCount)
	if err != nil {
		return err
	}
	curX := x

	for i, child := range children {
		childMar := effectiveMarginV1EngineLayoutAttributes(child)
		childW, fixed := explicitSizeV1EngineLayoutConstraints(child, "width")
		if !fixed {
			span := attrFloatV1EngineLayoutAttributes(child.Attr("span"), 12/float64(flexCount))
			childW = flexW * (span / 12.0)
		}
		if !isPositiveFiniteV1EngineLayoutConstraints(childW) {
			return newLayoutErrorV1EngineLayoutValidation(child, "allocated width must be positive and finite, got %.6g", childW)
		}
		cw := childW + childMar.Left + childMar.Right
		cb := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, i), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, cb, curX, y, cw, h); err != nil {
			return err
		}
		target.Children = append(target.Children, cb)
		curX += cw + gap
	}
	return nil
}

// layoutStagger places children in staggered depth-overlap mode.
// Each child is offset staggerOffset px right-and-down from the previous.
// Children are appended to target.Children in back-to-front render order
// (highest StaggerDepth first = rendered behind, depth 0 last = on top).
// Falls back to layoutStack when fewer than 2 children.
func layoutStaggerV1EngineLayoutFlow(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	children := layoutKidsV1EngineLayoutNode(node)
	n := len(children)
	if n < 2 {
		loggerV1EngineSharedLogging.DEBUG(IULST001V1EngineLayoutFlow, "branch fallback stack", map[string]any{"tag": node.Tag, "children": n})
		return layoutStackV1EngineLayoutFlow(node, target, x, y, w, h)
	}
	const staggerOffset = 16.0
	childW := w - staggerOffset*float64(n-1)
	childH := h - staggerOffset*float64(n-1)
	if childW < MinBoxWidthV1EngineLayoutFlow {
		loggerV1EngineSharedLogging.DEBUG(IULST002V1EngineLayoutFlow, "branch min width", map[string]any{"width": childW, "minWidth": MinBoxWidthV1EngineLayoutFlow})
		if normalizedOverflowV1EngineLayoutConstraints(node) != entity.OverflowVisible {
			return newLayoutErrorV1EngineLayoutValidation(node, "staggered children require at least %.0fpx width, but %.6gpx is available", MinBoxWidthV1EngineLayoutFlow, childW)
		}
		childW = MinBoxWidthV1EngineLayoutFlow
	}
	if childH < MinBoxHeightV1EngineLayoutFlow {
		loggerV1EngineSharedLogging.DEBUG(IULST003V1EngineLayoutFlow, "branch min height", map[string]any{"height": childH, "minHeight": MinBoxHeightV1EngineLayoutFlow})
		if normalizedOverflowV1EngineLayoutConstraints(node) != entity.OverflowVisible {
			return newLayoutErrorV1EngineLayoutValidation(node, "staggered children require at least %.0fpx height, but %.6gpx is available", MinBoxHeightV1EngineLayoutFlow, childH)
		}
		childH = MinBoxHeightV1EngineLayoutFlow
	}
	// Render back-to-front: highest depth first → behind, depth 0 last → front.
	for i := n - 1; i >= 0; i-- {
		child := children[i]
		cX := x + float64(i)*staggerOffset
		cY := y + float64(i)*staggerOffset
		cb := &entity.Box{
			ID:           childIDV1EngineLayoutAttributes(target.ID, i),
			Tag:          child.Tag,
			Label:        labelOfV1EngineLayoutAttributes(child),
			Position:     child.Position,
			StaggerDepth: i,
			IsStaggerBg:  i > 0,
			InStagger:    true,
		}
		if err := layoutNodeV1EngineLayoutNode(child, cb, cX, cY, childW, childH); err != nil {
			return err
		}
		target.Children = append(target.Children, cb)
	}
	return nil
}
