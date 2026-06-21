package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
)

const spacingUnit = 8

// MinBoxWidth / MinBoxHeight are the smallest dimensions at which a box
// can be meaningfully rendered. layoutStack and layoutRow clamp child
// sizes to these values so boxes are never invisible.
const (
	MinBoxWidth  = 60.0
	MinBoxHeight = 48.0
)

// defaultGroupInset is the automatic padding applied to container nodes
// (AWS group tags and unknown tags with children) when no explicit class
// padding is specified.  The top inset reserves room for the 32 px icon +
// label row; the side inset keeps children clear of the border line.
const (
	defaultGroupTopInset  = 44.0
	defaultGroupSideInset = 12.0

	// GroupTopInset / GroupSideInset are the exported equivalents used by
	// the excalidraw renderer to position item icons below the header row.
	GroupTopInset  = defaultGroupTopInset
	GroupSideInset = defaultGroupSideInset
)

type Box = entity.Box
type Spacing = entity.Spacing

func Build(doc entity.Document) (*Box, error) {
	if doc.Root == nil {
		return nil, fmt.Errorf("document root is nil")
	}
	w := attrFloat(doc.Root.Attr("width"), 1280)
	h := attrFloat(doc.Root.Attr("height"), 720)
	root := &Box{ID: "frame", Tag: "frame", Label: "frame", X: 0, Y: 0, W: w, H: h}
	layoutNode(doc.Root, root, 0, 0, w, h)
	return root, nil
}

// layoutKids returns node's children that participate in layout,
// filtering out meta-nodes such as <connection> which are handled separately.
func layoutKids(node *entity.Node) []*entity.Node {
	var kids []*entity.Node
	for _, c := range node.Children {
		if c.Tag == "connection" {
			continue
		}
		kids = append(kids, c)
	}
	return kids
}

func layoutNode(node *entity.Node, target *Box, x, y, w, h float64) {
	target.Attrs = node.Attrs
	pad, classMar := parseClassSpacing(node.Attr("class"))

	// 直接 px 指定マージン属性をクラスベースマージンに加算する
	attrMar := parseAttrMargin(node.Attrs)
	mar := Spacing{
		Top:    classMar.Top + attrMar.Top,
		Right:  classMar.Right + attrMar.Right,
		Bottom: classMar.Bottom + attrMar.Bottom,
		Left:   classMar.Left + attrMar.Left,
	}

	// margin は親から渡された割り当て領域を削る (sibling spacing)。
	// Root frame の margin は紙フレーム自体を縮めず、内側コンテンツの外側余白として扱う。
	boxX := x + mar.Left
	boxY := y + mar.Top
	boxW := w - mar.Left - mar.Right
	boxH := h - mar.Top - mar.Bottom
	if node.Tag == "frame" {
		boxX = x
		boxY = y
		boxW = w
		boxH = h
	}

	// width 属性が指定されていれば親計算値を上書きする (frame ルートは除く)
	if node.Tag != "frame" {
		if wv := node.Attr("width"); wv != "" {
			if ew := attrFloat(wv, 0); ew > 0 {
				boxW = ew
			}
		}
		if hv := node.Attr("height"); hv != "" {
			if eh := attrFloat(hv, 0); eh > 0 {
				boxH = eh
			}
		}
	}

	target.X = boxX
	target.Y = boxY
	target.W = boxW
	target.H = boxH

	// padding は box 内側の余白 (子要素の配置開始点)
	innerX := boxX + pad.Left
	innerY := boxY + pad.Top
	innerW := boxW - pad.Left - pad.Right
	innerH := boxH - pad.Top - pad.Bottom
	if node.Tag == "frame" {
		innerX += mar.Left
		innerY += mar.Top
		innerW -= mar.Left + mar.Right
		innerH -= mar.Top + mar.Bottom
	}
	innerX, innerY, innerW, innerH = alignContentArea(node, innerX, innerY, innerW, innerH)

	switch node.Tag {
	case "frame", "container":
		if node.Attr("layout") == "horizontal" {
			layoutFlexH(node, target, innerX, innerY, innerW, innerH)
		} else {
			layoutStack(node, target, innerX, innerY, innerW, innerH)
		}
	case "row":
		layoutRow(node, target, innerX, innerY, innerW, innerH)
	case "col":
		if node.Attr("layout") == "horizontal" {
			layoutFlexH(node, target, innerX, innerY, innerW, innerH)
		} else {
			layoutStack(node, target, innerX, innerY, innerW, innerH)
		}
	default:
		// AWS グループタグおよびその他の未知タグ:
		// 子要素があればコンテナ, なければリーフとして扱う。
		kids := layoutKids(node)
		if len(kids) > 0 {
			// <item> / <spacer> のみの親はグループアイコン/ラベルがないので topInset を適用しない
			allItems := true
			for _, ch := range kids {
				if !IsItemLike(ch.Tag) {
					allItems = false
					break
				}
			}
			if allItems {
				layoutRow(node, target, innerX, innerY, innerW, innerH)
				break
			}
			// グループ inset は常に適用。ユーザー指定 padding はその上に加算する。
			// これにより class="pa-2" でヘッダー行と子要素が重なることを防ぐ。
			gInnerX := boxX + defaultGroupSideInset + pad.Left
			gInnerY := boxY + defaultGroupTopInset + pad.Top
			gInnerW := boxW - defaultGroupSideInset*2 - pad.Left - pad.Right
			gInnerH := boxH - defaultGroupTopInset - defaultGroupSideInset - pad.Top - pad.Bottom
			gInnerX, gInnerY, gInnerW, gInnerH = alignContentArea(node, gInnerX, gInnerY, gInnerW, gInnerH)
			if node.Attr("layout") == "staggered" {
				layoutStagger(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			} else if node.Attr("layout") == "horizontal" {
				layoutFlexH(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			} else {
				layoutStack(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			}
		} else {
			layoutLeaf(node, target, innerX, innerY, innerW, innerH)
		}
	}
}

func layoutStack(node *entity.Node, target *Box, x, y, w, h float64) {
	children := layoutKids(node)
	if len(children) == 0 {
		return
	}
	gap := attrFloat(node.Attr("gap"), 16)

	// 各子要素の margin を事前に読み取り、縦方向の余白合計を算出する。
	// これにより margin が sibling 間スペースとして機能する (CSS ライク)。
	// row 属性は flex-grow スタイルの高さ比率。デフォルト 1.0 (均等)。
	totalMarginH := 0.0
	totalRow := 0.0
	for _, child := range children {
		childMar := effectiveMargin(child)
		totalMarginH += childMar.Top + childMar.Bottom
		totalRow += attrFloat(child.Attr("row"), 1.0)
	}
	availH := h - gap*float64(len(children)-1) - totalMarginH

	curY := y
	for i, child := range children {
		childMar := effectiveMargin(child)
		row := attrFloat(child.Attr("row"), 1.0)
		// 子への割り当て = 比率に応じた content 高さ + その子自身の上下 margin
		childH := availH * (row / totalRow)
		alloc := childH + childMar.Top + childMar.Bottom
		cb := &Box{ID: childID(target.ID, i), Tag: child.Tag, Label: labelOf(child)}
		layoutNode(child, cb, x, curY, w, alloc)
		target.Children = append(target.Children, cb)
		curY += alloc + gap
	}
}

// layoutFlexH lays out children horizontally with free ratio weights.
// Each child's width share is determined by its `col` attribute (default 1.0).
// This mirrors layoutStack but in the horizontal direction.
func layoutFlexH(node *entity.Node, target *Box, x, y, w, h float64) {
	children := layoutKids(node)
	if len(children) == 0 {
		return
	}
	gap := attrFloat(node.Attr("gap"), 16)

	// 各子要素の水平 margin を事前集計し利用可能幅を算出する。
	// col 属性は flex-grow スタイルの幅比率。デフォルト 1.0 (均等)。
	totalMarginW := 0.0
	totalCol := 0.0
	for _, child := range children {
		childMar := effectiveMargin(child)
		totalMarginW += childMar.Left + childMar.Right
		totalCol += attrFloat(child.Attr("col"), 1.0)
	}
	availW := w - gap*float64(len(children)-1) - totalMarginW

	curX := x
	for i, child := range children {
		childMar := effectiveMargin(child)
		col := attrFloat(child.Attr("col"), 1.0)
		// 子への割り当て = 比率に応じた content 幅 + その子自身の左右 margin
		childW := availW * (col / totalCol)
		alloc := childW + childMar.Left + childMar.Right
		cb := &Box{ID: childID(target.ID, i), Tag: child.Tag, Label: labelOf(child)}
		layoutNode(child, cb, curX, y, alloc, h)
		target.Children = append(target.Children, cb)
		curX += alloc + gap
	}
}

func layoutRow(node *entity.Node, target *Box, x, y, w, h float64) {
	children := layoutKids(node)
	if len(children) == 0 {
		return
	}
	gap := attrFloat(node.Attr("gap"), 16)

	// 各子要素の水平 margin を事前に読み取り、幅方向の合計を算出する。
	totalMarginW := 0.0
	for _, child := range children {
		childMar := effectiveMargin(child)
		totalMarginW += childMar.Left + childMar.Right
	}
	remainingW := w - gap*float64(len(children)-1) - totalMarginW
	curX := x

	for i, child := range children {
		childMar := effectiveMargin(child)
		span := attrFloat(child.Attr("span"), 12/float64(len(children)))
		cw := remainingW*(span/12.0) + childMar.Left + childMar.Right
		cb := &Box{ID: childID(target.ID, i), Tag: child.Tag, Label: labelOf(child)}
		layoutNode(child, cb, curX, y, cw, h)
		target.Children = append(target.Children, cb)
		curX += cw + gap
	}
}

// layoutStagger places children in staggered depth-overlap mode.
// Each child is offset staggerOffset px right-and-down from the previous.
// Children are appended to target.Children in back-to-front render order
// (highest StaggerDepth first = rendered behind, depth 0 last = on top).
// Falls back to layoutStack when fewer than 2 children.
func layoutStagger(node *entity.Node, target *Box, x, y, w, h float64) {
	children := layoutKids(node)
	n := len(children)
	if n < 2 {
		layoutStack(node, target, x, y, w, h)
		return
	}
	const staggerOffset = 16.0
	childW := w - staggerOffset*float64(n-1)
	childH := h - staggerOffset*float64(n-1)
	if childW < MinBoxWidth {
		childW = MinBoxWidth
	}
	if childH < MinBoxHeight {
		childH = MinBoxHeight
	}
	// Render back-to-front: highest depth first → behind, depth 0 last → front.
	for i := n - 1; i >= 0; i-- {
		child := children[i]
		cX := x + float64(i)*staggerOffset
		cY := y + float64(i)*staggerOffset
		cb := &Box{
			ID:           childID(target.ID, i),
			Tag:          child.Tag,
			Label:        labelOf(child),
			StaggerDepth: i,
			IsStaggerBg:  i > 0,
			InStagger:    true,
		}
		layoutNode(child, cb, cX, cY, childW, childH)
		target.Children = append(target.Children, cb)
	}
}

func layoutLeaf(node *entity.Node, target *Box, x, y, w, h float64) {
	target.X = x
	target.Y = y
	target.W = w
	target.H = h
}

func alignContentArea(node *entity.Node, x, y, w, h float64) (float64, float64, float64, float64) {
	contentW := attrFloat(node.Attr("content-width"), w)
	contentH := attrFloat(node.Attr("content-height"), h)
	if contentW <= 0 || contentW > w {
		contentW = w
	}
	if contentH <= 0 || contentH > h {
		contentH = h
	}
	vert, horiz := parseAlign(node.Attr("align"))
	switch horiz {
	case "right":
		x += w - contentW
	case "center":
		x += (w - contentW) / 2
	}
	switch vert {
	case "bottom":
		y += h - contentH
	case "middle":
		y += (h - contentH) / 2
	}
	return x, y, contentW, contentH
}

func childID(parent string, index int) string {
	return fmt.Sprintf("%s-%d", parent, index)
}

func labelOf(n *entity.Node) string {
	if title := n.Attr("title"); title != "" {
		return title
	}
	if n.Text != "" {
		return n.Text
	}
	return n.Tag
}

func attrFloat(v string, fallback float64) float64 {
	if strings.TrimSpace(v) == "" {
		return fallback
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return fallback
	}
	return f
}

func parseClassSpacing(class string) (Spacing, Spacing) {
	pad := Spacing{}
	mar := Spacing{}
	for _, tok := range strings.Fields(class) {
		switch {
		case strings.HasPrefix(tok, "pa-"):
			v := spacingValue(tok[3:])
			pad = Spacing{Top: v, Right: v, Bottom: v, Left: v}
		case strings.HasPrefix(tok, "ma-"):
			v := spacingValue(tok[3:])
			mar = Spacing{Top: v, Right: v, Bottom: v, Left: v}
		// 軸別一括: px=左右, py=上下
		case strings.HasPrefix(tok, "px-"):
			v := spacingValue(tok[3:])
			pad.Left = v
			pad.Right = v
		case strings.HasPrefix(tok, "py-"):
			v := spacingValue(tok[3:])
			pad.Top = v
			pad.Bottom = v
		case strings.HasPrefix(tok, "mx-"):
			v := spacingValue(tok[3:])
			mar.Left = v
			mar.Right = v
		case strings.HasPrefix(tok, "my-"):
			v := spacingValue(tok[3:])
			mar.Top = v
			mar.Bottom = v
		// 個別方向
		case strings.HasPrefix(tok, "pt-"):
			pad.Top = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "pr-"):
			pad.Right = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "pb-"):
			pad.Bottom = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "pl-"):
			pad.Left = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "mt-"):
			mar.Top = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "mr-"):
			mar.Right = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "mb-"):
			mar.Bottom = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "ml-"):
			mar.Left = spacingValue(tok[3:])
		}
	}
	return pad, mar
}

// IsItemLike reports whether a tag behaves as a layout item slot.
func IsItemLike(tag string) bool {
	return tag == "item" || IsBlank(tag)
}

// IsBlank reports whether a tag participates in layout without rendering.
func IsBlank(tag string) bool {
	return tag == "spacer" || tag == "blank"
}

func parseAlign(align string) (vert, horiz string) {
	vert, horiz = "top", "left"
	parts := strings.SplitN(strings.ToLower(strings.TrimSpace(align)), "-", 2)
	if len(parts) != 2 {
		return vert, horiz
	}
	if parts[0] == "top" || parts[0] == "middle" || parts[0] == "bottom" {
		vert = parts[0]
	}
	if parts[1] == "left" || parts[1] == "center" || parts[1] == "right" {
		horiz = parts[1]
	}
	return vert, horiz
}

func spacingValue(s string) float64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return float64(n * spacingUnit)
}

// parseAttrMargin reads direct pixel-value margin attributes from a DSL node's
// attribute map. Supported attributes: margin, margin-top, margin-right,
// margin-bottom, margin-left. Values are in pixels (floats).
// When `margin` and a directional key (e.g. `margin-top`) are both present,
// the directional key overrides the corresponding side from `margin`.
func parseAttrMargin(attrs map[string]string) Spacing {
	if len(attrs) == 0 {
		return Spacing{}
	}
	m := Spacing{}
	if v := attrs["margin"]; v != "" {
		val := attrFloat(v, 0)
		m = Spacing{Top: val, Right: val, Bottom: val, Left: val}
	}
	if v := attrs["margin-top"]; v != "" {
		m.Top = attrFloat(v, 0)
	}
	if v := attrs["margin-right"]; v != "" {
		m.Right = attrFloat(v, 0)
	}
	if v := attrs["margin-bottom"]; v != "" {
		m.Bottom = attrFloat(v, 0)
	}
	if v := attrs["margin-left"]; v != "" {
		m.Left = attrFloat(v, 0)
	}
	return m
}

// effectiveMargin returns the combined margin for a node by summing
// class-based spacing (ma-N, mt-N …) and direct px-value attributes
// (margin, margin-top …).
func effectiveMargin(node *entity.Node) Spacing {
	_, classMar := parseClassSpacing(node.Attr("class"))
	attrMar := parseAttrMargin(node.Attrs)
	return Spacing{
		Top:    classMar.Top + attrMar.Top,
		Right:  classMar.Right + attrMar.Right,
		Bottom: classMar.Bottom + attrMar.Bottom,
		Left:   classMar.Left + attrMar.Left,
	}
}
