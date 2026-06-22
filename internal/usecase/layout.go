package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
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

var (
	IULB001   = share.NewMCode("IULB-001", "Build nil root branch")
	IULB002   = share.NewMCode("IULB-002", "Build layout root branch")
	IULLK001  = share.NewMCode("IULLK-001", "Layout kids skip connection branch")
	IULN001   = share.NewMCode("IULN-001", "Layout node frame margin branch")
	IULN002   = share.NewMCode("IULN-002", "Layout node explicit width branch")
	IULN003   = share.NewMCode("IULN-003", "Layout node explicit height branch")
	IULN004   = share.NewMCode("IULN-004", "Layout node frame inner margin branch")
	IULN005   = share.NewMCode("IULN-005", "Layout node frame horizontal branch")
	IULN006   = share.NewMCode("IULN-006", "Layout node frame stack branch")
	IULN007   = share.NewMCode("IULN-007", "Layout node row branch")
	IULN008   = share.NewMCode("IULN-008", "Layout node col horizontal branch")
	IULN009   = share.NewMCode("IULN-009", "Layout node col stack branch")
	IULN010   = share.NewMCode("IULN-010", "Layout node default container branch")
	IULN011   = share.NewMCode("IULN-011", "Layout node default all items branch")
	IULN012   = share.NewMCode("IULN-012", "Layout node default staggered branch")
	IULN013   = share.NewMCode("IULN-013", "Layout node default horizontal branch")
	IULN014   = share.NewMCode("IULN-014", "Layout node default stack branch")
	IULN015   = share.NewMCode("IULN-015", "Layout node leaf branch")
	IULS001   = share.NewMCode("IULS-001", "Layout stack empty children branch")
	IULFH001  = share.NewMCode("IULFH-001", "Layout flex horizontal empty children branch")
	IULR001   = share.NewMCode("IULR-001", "Layout row empty children branch")
	IULST001  = share.NewMCode("IULST-001", "Layout stagger fallback branch")
	IULST002  = share.NewMCode("IULST-002", "Layout stagger minimum width branch")
	IULST003  = share.NewMCode("IULST-003", "Layout stagger minimum height branch")
	IULACA001 = share.NewMCode("IULACA-001", "Align content area width clamp branch")
	IULACA002 = share.NewMCode("IULACA-002", "Align content area height clamp branch")
	IULACA003 = share.NewMCode("IULACA-003", "Align content area right branch")
	IULACA004 = share.NewMCode("IULACA-004", "Align content area center branch")
	IULACA005 = share.NewMCode("IULACA-005", "Align content area bottom branch")
	IULACA006 = share.NewMCode("IULACA-006", "Align content area middle branch")
	IULLO001  = share.NewMCode("IULLO-001", "Label of title branch")
	IULLO002  = share.NewMCode("IULLO-002", "Label of text branch")
	IULLO003  = share.NewMCode("IULLO-003", "Label of tag branch")
	IULAF001  = share.NewMCode("IULAF-001", "Attribute float fallback empty branch")
	IULAF002  = share.NewMCode("IULAF-002", "Attribute float parse failed branch")
	IULPCS001 = share.NewMCode("IULPCS-001", "Parse class spacing padding all branch")
	IULPCS002 = share.NewMCode("IULPCS-002", "Parse class spacing margin all branch")
	IULPCS003 = share.NewMCode("IULPCS-003", "Parse class spacing padding x branch")
	IULPCS004 = share.NewMCode("IULPCS-004", "Parse class spacing padding y branch")
	IULPCS005 = share.NewMCode("IULPCS-005", "Parse class spacing margin x branch")
	IULPCS006 = share.NewMCode("IULPCS-006", "Parse class spacing margin y branch")
	IULPCS007 = share.NewMCode("IULPCS-007", "Parse class spacing padding top branch")
	IULPCS008 = share.NewMCode("IULPCS-008", "Parse class spacing padding right branch")
	IULPCS009 = share.NewMCode("IULPCS-009", "Parse class spacing padding bottom branch")
	IULPCS010 = share.NewMCode("IULPCS-010", "Parse class spacing padding left branch")
	IULPCS011 = share.NewMCode("IULPCS-011", "Parse class spacing margin top branch")
	IULPCS012 = share.NewMCode("IULPCS-012", "Parse class spacing margin right branch")
	IULPCS013 = share.NewMCode("IULPCS-013", "Parse class spacing margin bottom branch")
	IULPCS014 = share.NewMCode("IULPCS-014", "Parse class spacing margin left branch")
	IULPA001  = share.NewMCode("IULPA-001", "Parse align invalid branch")
	IULPA002  = share.NewMCode("IULPA-002", "Parse align vertical branch")
	IULPA003  = share.NewMCode("IULPA-003", "Parse align horizontal branch")
	IULSV001  = share.NewMCode("IULSV-001", "Spacing value parse failed branch")
	IULPAM001 = share.NewMCode("IULPAM-001", "Parse attribute margin empty branch")
	IULPAM002 = share.NewMCode("IULPAM-002", "Parse attribute margin all branch")
	IULPAM003 = share.NewMCode("IULPAM-003", "Parse attribute margin top branch")
	IULPAM004 = share.NewMCode("IULPAM-004", "Parse attribute margin right branch")
	IULPAM005 = share.NewMCode("IULPAM-005", "Parse attribute margin bottom branch")
	IULPAM006 = share.NewMCode("IULPAM-006", "Parse attribute margin left branch")
)

func Build(doc entity.Document) (*entity.Box, error) {
	if doc.Root == nil {
		logger.DEBUG(IULB001, "branch nil root")
		return nil, fmt.Errorf("document root is nil")
	}
	w := attrFloat(doc.Root.Attr("width"), 1280)
	h := attrFloat(doc.Root.Attr("height"), 720)
	root := &entity.Box{ID: "frame", Tag: "frame", Label: "frame", X: 0, Y: 0, W: w, H: h}
	logger.DEBUG(IULB002, "branch layout root", map[string]any{"width": w, "height": h})
	layoutNode(doc.Root, root, 0, 0, w, h)
	return root, nil
}

// layoutKids returns node's children that participate in layout,
// filtering out meta-nodes such as <connection> which are handled separately.
func layoutKids(node *entity.Node) []*entity.Node {
	var kids []*entity.Node
	for _, c := range node.Children {
		if c.Tag == "connection" {
			logger.DEBUG(IULLK001, "branch skip connection")
			continue
		}
		kids = append(kids, c)
	}
	return kids
}

func layoutNode(node *entity.Node, target *entity.Box, x, y, w, h float64) {
	target.Attrs = node.Attrs
	pad, classMar := parseClassSpacing(node.Attr("class"))

	// 直接 px 指定マージン属性をクラスベースマージンに加算する
	attrMar := parseAttrMargin(node.Attrs)
	mar := entity.Spacing{
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
		logger.DEBUG(IULN001, "branch frame margin")
		boxX = x
		boxY = y
		boxW = w
		boxH = h
	}

	// width 属性が指定されていれば親計算値を上書きする (frame ルートは除く)
	if node.Tag != "frame" {
		if wv := node.Attr("width"); wv != "" {
			if ew := attrFloat(wv, 0); ew > 0 {
				logger.DEBUG(IULN002, "branch explicit width", map[string]any{"tag": node.Tag, "width": ew})
				boxW = ew
			}
		}
		if hv := node.Attr("height"); hv != "" {
			if eh := attrFloat(hv, 0); eh > 0 {
				logger.DEBUG(IULN003, "branch explicit height", map[string]any{"tag": node.Tag, "height": eh})
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
		logger.DEBUG(IULN004, "branch frame inner margin")
		innerX += mar.Left
		innerY += mar.Top
		innerW -= mar.Left + mar.Right
		innerH -= mar.Top + mar.Bottom
	}
	innerX, innerY, innerW, innerH = alignContentArea(node, innerX, innerY, innerW, innerH)

	switch node.Tag {
	case "frame", "container":
		if node.Attr("layout") == "horizontal" {
			logger.DEBUG(IULN005, "branch frame/container horizontal", map[string]any{"tag": node.Tag})
			layoutFlexH(node, target, innerX, innerY, innerW, innerH)
		} else {
			logger.DEBUG(IULN006, "branch frame/container stack", map[string]any{"tag": node.Tag})
			layoutStack(node, target, innerX, innerY, innerW, innerH)
		}
	case "row":
		logger.DEBUG(IULN007, "branch row")
		layoutRow(node, target, innerX, innerY, innerW, innerH)
	case "col":
		if node.Attr("layout") == "horizontal" {
			logger.DEBUG(IULN008, "branch col horizontal")
			layoutFlexH(node, target, innerX, innerY, innerW, innerH)
		} else {
			logger.DEBUG(IULN009, "branch col stack")
			layoutStack(node, target, innerX, innerY, innerW, innerH)
		}
	default:
		// AWS グループタグおよびその他の未知タグ:
		// 子要素があればコンテナ, なければリーフとして扱う。
		kids := layoutKids(node)
		if len(kids) > 0 {
			logger.DEBUG(IULN010, "branch default container", map[string]any{"tag": node.Tag, "children": len(kids)})
			// <item> / <spacer> のみの親はグループアイコン/ラベルがないので topInset を適用しない
			allItems := true
			for _, ch := range kids {
				if !IsItemLike(ch.Tag) {
					allItems = false
					break
				}
			}
			if allItems {
				logger.DEBUG(IULN011, "branch all item children", map[string]any{"tag": node.Tag})
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
				logger.DEBUG(IULN012, "branch staggered", map[string]any{"tag": node.Tag})
				layoutStagger(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			} else if node.Attr("layout") == "horizontal" {
				logger.DEBUG(IULN013, "branch horizontal", map[string]any{"tag": node.Tag})
				layoutFlexH(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			} else {
				logger.DEBUG(IULN014, "branch stack", map[string]any{"tag": node.Tag})
				layoutStack(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			}
		} else {
			logger.DEBUG(IULN015, "branch leaf", map[string]any{"tag": node.Tag})
			layoutLeaf(node, target, innerX, innerY, innerW, innerH)
		}
	}
}

func layoutStack(node *entity.Node, target *entity.Box, x, y, w, h float64) {
	children := layoutKids(node)
	if len(children) == 0 {
		logger.DEBUG(IULS001, "branch empty children", map[string]any{"tag": node.Tag})
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
		cb := &entity.Box{ID: childID(target.ID, i), Tag: child.Tag, Label: labelOf(child)}
		layoutNode(child, cb, x, curY, w, alloc)
		target.Children = append(target.Children, cb)
		curY += alloc + gap
	}
}

// layoutFlexH lays out children horizontally with free ratio weights.
// Each child's width share is determined by its `col` attribute (default 1.0).
// This mirrors layoutStack but in the horizontal direction.
func layoutFlexH(node *entity.Node, target *entity.Box, x, y, w, h float64) {
	children := layoutKids(node)
	if len(children) == 0 {
		logger.DEBUG(IULFH001, "branch empty children", map[string]any{"tag": node.Tag})
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
		cb := &entity.Box{ID: childID(target.ID, i), Tag: child.Tag, Label: labelOf(child)}
		layoutNode(child, cb, curX, y, alloc, h)
		target.Children = append(target.Children, cb)
		curX += alloc + gap
	}
}

func layoutRow(node *entity.Node, target *entity.Box, x, y, w, h float64) {
	children := layoutKids(node)
	if len(children) == 0 {
		logger.DEBUG(IULR001, "branch empty children", map[string]any{"tag": node.Tag})
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
		cb := &entity.Box{ID: childID(target.ID, i), Tag: child.Tag, Label: labelOf(child)}
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
func layoutStagger(node *entity.Node, target *entity.Box, x, y, w, h float64) {
	children := layoutKids(node)
	n := len(children)
	if n < 2 {
		logger.DEBUG(IULST001, "branch fallback stack", map[string]any{"tag": node.Tag, "children": n})
		layoutStack(node, target, x, y, w, h)
		return
	}
	const staggerOffset = 16.0
	childW := w - staggerOffset*float64(n-1)
	childH := h - staggerOffset*float64(n-1)
	if childW < MinBoxWidth {
		logger.DEBUG(IULST002, "branch min width", map[string]any{"width": childW, "minWidth": MinBoxWidth})
		childW = MinBoxWidth
	}
	if childH < MinBoxHeight {
		logger.DEBUG(IULST003, "branch min height", map[string]any{"height": childH, "minHeight": MinBoxHeight})
		childH = MinBoxHeight
	}
	// Render back-to-front: highest depth first → behind, depth 0 last → front.
	for i := n - 1; i >= 0; i-- {
		child := children[i]
		cX := x + float64(i)*staggerOffset
		cY := y + float64(i)*staggerOffset
		cb := &entity.Box{
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

func layoutLeaf(node *entity.Node, target *entity.Box, x, y, w, h float64) {
	target.X = x
	target.Y = y
	target.W = w
	target.H = h
}

func alignContentArea(node *entity.Node, x, y, w, h float64) (float64, float64, float64, float64) {
	contentW := attrFloat(node.Attr("content-width"), w)
	contentH := attrFloat(node.Attr("content-height"), h)
	if contentW <= 0 || contentW > w {
		logger.DEBUG(IULACA001, "branch clamp width", map[string]any{"tag": node.Tag, "contentWidth": contentW, "width": w})
		contentW = w
	}
	if contentH <= 0 || contentH > h {
		logger.DEBUG(IULACA002, "branch clamp height", map[string]any{"tag": node.Tag, "contentHeight": contentH, "height": h})
		contentH = h
	}
	vert, horiz := parseAlign(node.Attr("align"))
	switch horiz {
	case "right":
		logger.DEBUG(IULACA003, "branch right", map[string]any{"tag": node.Tag})
		x += w - contentW
	case "center":
		logger.DEBUG(IULACA004, "branch center", map[string]any{"tag": node.Tag})
		x += (w - contentW) / 2
	}
	switch vert {
	case "bottom":
		logger.DEBUG(IULACA005, "branch bottom", map[string]any{"tag": node.Tag})
		y += h - contentH
	case "middle":
		logger.DEBUG(IULACA006, "branch middle", map[string]any{"tag": node.Tag})
		y += (h - contentH) / 2
	}
	return x, y, contentW, contentH
}

func childID(parent string, index int) string {
	return fmt.Sprintf("%s-%d", parent, index)
}

func labelOf(n *entity.Node) string {
	if title := n.Attr("title"); title != "" {
		logger.DEBUG(IULLO001, "branch title", map[string]any{"tag": n.Tag})
		return title
	}
	if n.Text != "" {
		logger.DEBUG(IULLO002, "branch text", map[string]any{"tag": n.Tag})
		return n.Text
	}
	logger.DEBUG(IULLO003, "branch tag", map[string]any{"tag": n.Tag})
	return n.Tag
}

func attrFloat(v string, fallback float64) float64 {
	if strings.TrimSpace(v) == "" {
		logger.DEBUG(IULAF001, "branch fallback empty")
		return fallback
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		logger.WARN(IULAF002, "branch parse failed", map[string]any{"value": v, "error": err})
		return fallback
	}
	return f
}

func parseClassSpacing(class string) (entity.Spacing, entity.Spacing) {
	pad := entity.Spacing{}
	mar := entity.Spacing{}
	for _, tok := range strings.Fields(class) {
		switch {
		case strings.HasPrefix(tok, "pa-"):
			logger.DEBUG(IULPCS001, "branch padding all", map[string]any{"token": tok})
			v := spacingValue(tok[3:])
			pad = entity.Spacing{Top: v, Right: v, Bottom: v, Left: v}
		case strings.HasPrefix(tok, "ma-"):
			logger.DEBUG(IULPCS002, "branch margin all", map[string]any{"token": tok})
			v := spacingValue(tok[3:])
			mar = entity.Spacing{Top: v, Right: v, Bottom: v, Left: v}
		// 軸別一括: px=左右, py=上下
		case strings.HasPrefix(tok, "px-"):
			logger.DEBUG(IULPCS003, "branch padding x", map[string]any{"token": tok})
			v := spacingValue(tok[3:])
			pad.Left = v
			pad.Right = v
		case strings.HasPrefix(tok, "py-"):
			logger.DEBUG(IULPCS004, "branch padding y", map[string]any{"token": tok})
			v := spacingValue(tok[3:])
			pad.Top = v
			pad.Bottom = v
		case strings.HasPrefix(tok, "mx-"):
			logger.DEBUG(IULPCS005, "branch margin x", map[string]any{"token": tok})
			v := spacingValue(tok[3:])
			mar.Left = v
			mar.Right = v
		case strings.HasPrefix(tok, "my-"):
			logger.DEBUG(IULPCS006, "branch margin y", map[string]any{"token": tok})
			v := spacingValue(tok[3:])
			mar.Top = v
			mar.Bottom = v
		// 個別方向
		case strings.HasPrefix(tok, "pt-"):
			logger.DEBUG(IULPCS007, "branch padding top", map[string]any{"token": tok})
			pad.Top = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "pr-"):
			logger.DEBUG(IULPCS008, "branch padding right", map[string]any{"token": tok})
			pad.Right = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "pb-"):
			logger.DEBUG(IULPCS009, "branch padding bottom", map[string]any{"token": tok})
			pad.Bottom = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "pl-"):
			logger.DEBUG(IULPCS010, "branch padding left", map[string]any{"token": tok})
			pad.Left = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "mt-"):
			logger.DEBUG(IULPCS011, "branch margin top", map[string]any{"token": tok})
			mar.Top = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "mr-"):
			logger.DEBUG(IULPCS012, "branch margin right", map[string]any{"token": tok})
			mar.Right = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "mb-"):
			logger.DEBUG(IULPCS013, "branch margin bottom", map[string]any{"token": tok})
			mar.Bottom = spacingValue(tok[3:])
		case strings.HasPrefix(tok, "ml-"):
			logger.DEBUG(IULPCS014, "branch margin left", map[string]any{"token": tok})
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
		logger.WARN(IULPA001, "branch invalid align", map[string]any{"align": align})
		return vert, horiz
	}
	if parts[0] == "top" || parts[0] == "middle" || parts[0] == "bottom" {
		logger.DEBUG(IULPA002, "branch vertical", map[string]any{"vertical": parts[0]})
		vert = parts[0]
	}
	if parts[1] == "left" || parts[1] == "center" || parts[1] == "right" {
		logger.DEBUG(IULPA003, "branch horizontal", map[string]any{"horizontal": parts[1]})
		horiz = parts[1]
	}
	return vert, horiz
}

func spacingValue(s string) float64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		logger.WARN(IULSV001, "branch parse failed", map[string]any{"value": s, "error": err})
		return 0
	}
	return float64(n * spacingUnit)
}

// parseAttrMargin reads direct pixel-value margin attributes from a DSL node's
// attribute map. Supported attributes: margin, margin-top, margin-right,
// margin-bottom, margin-left. Values are in pixels (floats).
// When `margin` and a directional key (e.g. `margin-top`) are both present,
// the directional key overrides the corresponding side from `margin`.
func parseAttrMargin(attrs map[string]string) entity.Spacing {
	if len(attrs) == 0 {
		logger.DEBUG(IULPAM001, "branch empty attrs")
		return entity.Spacing{}
	}
	m := entity.Spacing{}
	if v := attrs["margin"]; v != "" {
		logger.DEBUG(IULPAM002, "branch margin all", map[string]any{"value": v})
		val := attrFloat(v, 0)
		m = entity.Spacing{Top: val, Right: val, Bottom: val, Left: val}
	}
	if v := attrs["margin-top"]; v != "" {
		logger.DEBUG(IULPAM003, "branch margin top", map[string]any{"value": v})
		m.Top = attrFloat(v, 0)
	}
	if v := attrs["margin-right"]; v != "" {
		logger.DEBUG(IULPAM004, "branch margin right", map[string]any{"value": v})
		m.Right = attrFloat(v, 0)
	}
	if v := attrs["margin-bottom"]; v != "" {
		logger.DEBUG(IULPAM005, "branch margin bottom", map[string]any{"value": v})
		m.Bottom = attrFloat(v, 0)
	}
	if v := attrs["margin-left"]; v != "" {
		logger.DEBUG(IULPAM006, "branch margin left", map[string]any{"value": v})
		m.Left = attrFloat(v, 0)
	}
	return m
}

// effectiveMargin returns the combined margin for a node by summing
// class-based spacing (ma-N, mt-N …) and direct px-value attributes
// (margin, margin-top …).
func effectiveMargin(node *entity.Node) entity.Spacing {
	_, classMar := parseClassSpacing(node.Attr("class"))
	attrMar := parseAttrMargin(node.Attrs)
	return entity.Spacing{
		Top:    classMar.Top + attrMar.Top,
		Right:  classMar.Right + attrMar.Right,
		Bottom: classMar.Bottom + attrMar.Bottom,
		Left:   classMar.Left + attrMar.Left,
	}
}
