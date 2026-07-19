package engine

import (
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

// defaultGroupInset is the automatic padding applied to container nodes
// (AWS group tags and unknown tags with children) when no explicit class
// padding is specified.  The top inset reserves room for the 32 px icon +
// label row; the side inset keeps children clear of the border line.
const (
	defaultGroupTopInsetV1EngineLayoutNode  = 44.0
	defaultGroupSideInsetV1EngineLayoutNode = 12.0

	// GroupTopInsetV1EngineLayoutNode / GroupSideInset are the exported equivalents used by
	// the excalidraw renderer to position item icons below the header row.
	GroupTopInsetV1EngineLayoutNode  = defaultGroupTopInsetV1EngineLayoutNode
	GroupSideInsetV1EngineLayoutNode = defaultGroupSideInsetV1EngineLayoutNode
)

var (
	IULLK001V1EngineLayoutNode = share.NewMCode("IULLK-001", "Layout kids skip connection branch")
	IULN001V1EngineLayoutNode  = share.NewMCode("IULN-001", "Layout node frame margin branch")
	IULN002V1EngineLayoutNode  = share.NewMCode("IULN-002", "Layout node explicit width branch")
	IULN003V1EngineLayoutNode  = share.NewMCode("IULN-003", "Layout node explicit height branch")
	IULN004V1EngineLayoutNode  = share.NewMCode("IULN-004", "Layout node frame inner margin branch")
	IULN005V1EngineLayoutNode  = share.NewMCode("IULN-005", "Layout node frame horizontal branch")
	IULN006V1EngineLayoutNode  = share.NewMCode("IULN-006", "Layout node frame stack branch")
	IULN007V1EngineLayoutNode  = share.NewMCode("IULN-007", "Layout node row branch")
	IULN008V1EngineLayoutNode  = share.NewMCode("IULN-008", "Layout node col horizontal branch")
	IULN009V1EngineLayoutNode  = share.NewMCode("IULN-009", "Layout node col stack branch")
	IULN010V1EngineLayoutNode  = share.NewMCode("IULN-010", "Layout node default container branch")
	IULN011V1EngineLayoutNode  = share.NewMCode("IULN-011", "Layout node default all items branch")
	IULN012V1EngineLayoutNode  = share.NewMCode("IULN-012", "Layout node default staggered branch")
	IULN013V1EngineLayoutNode  = share.NewMCode("IULN-013", "Layout node default horizontal branch")
	IULN014V1EngineLayoutNode  = share.NewMCode("IULN-014", "Layout node default stack branch")
	IULN015V1EngineLayoutNode  = share.NewMCode("IULN-015", "Layout node leaf branch")
)

func layoutKidsV1EngineLayoutNode(node *entity.Node) []*entity.Node {
	var kids []*entity.Node
	for _, c := range node.Children {
		if c.Tag == "connection" || c.Tag == "connections" || c.Tag == "metadata" {
			loggerV1EngineSharedLogging.DEBUG(IULLK001V1EngineLayoutNode, "branch skip connection")
			continue
		}
		kids = append(kids, c)
	}
	return kids
}

func layoutNodeV1EngineLayoutNode(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	target.Attrs = node.Attrs
	target.Position = node.Position
	target.Overflow = normalizedOverflowV1EngineLayoutConstraints(node)
	if intrinsicW, ok := explicitSizeV1EngineLayoutConstraints(node, "width"); ok {
		target.IntrinsicW = intrinsicW
	}
	if intrinsicH, ok := explicitSizeV1EngineLayoutConstraints(node, "height"); ok {
		target.IntrinsicH = intrinsicH
	}
	pad, classMar := parseClassSpacingV1EngineLayoutAttributes(node.Attr("class"))

	// 直接 px 指定マージン属性をクラスベースマージンに加算する
	attrMar := parseAttrMarginV1EngineLayoutAttributes(node.Attrs)
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
		loggerV1EngineSharedLogging.DEBUG(IULN001V1EngineLayoutNode, "branch frame margin")
		boxX = x
		boxY = y
		boxW = w
		boxH = h
	}

	// width 属性が指定されていれば親計算値を上書きする (frame ルートは除く)
	if node.Tag != "frame" {
		if wv := node.Attr("width"); wv != "" {
			if ew := attrFloatV1EngineLayoutAttributes(wv, 0); ew > 0 {
				loggerV1EngineSharedLogging.DEBUG(IULN002V1EngineLayoutNode, "branch explicit width", map[string]any{"tag": node.Tag, "width": ew})
				boxW = ew
			}
		}
		if hv := node.Attr("height"); hv != "" {
			if eh := attrFloatV1EngineLayoutAttributes(hv, 0); eh > 0 {
				loggerV1EngineSharedLogging.DEBUG(IULN003V1EngineLayoutNode, "branch explicit height", map[string]any{"tag": node.Tag, "height": eh})
				boxH = eh
			}
		}
	}

	target.X = boxX
	target.Y = boxY
	target.W = boxW
	target.H = boxH
	if !isPositiveFiniteV1EngineLayoutConstraints(boxW) || !isPositiveFiniteV1EngineLayoutConstraints(boxH) {
		return newLayoutErrorV1EngineLayoutValidation(node, "resolved border size must be positive and finite, got %.6gx%.6g", boxW, boxH)
	}

	// padding は box 内側の余白 (子要素の配置開始点)
	zoneX := boxX + pad.Left
	zoneY := boxY + pad.Top
	zoneW := boxW - pad.Left - pad.Right
	zoneH := boxH - pad.Top - pad.Bottom
	innerX := zoneX
	innerY := zoneY
	innerW := zoneW
	innerH := zoneH
	if node.Tag == "frame" {
		loggerV1EngineSharedLogging.DEBUG(IULN004V1EngineLayoutNode, "branch frame inner margin")
		innerX += mar.Left
		innerY += mar.Top
		innerW -= mar.Left + mar.Right
		innerH -= mar.Top + mar.Bottom
	}
	var err error
	if node.Tag == "frame" {
		innerX, innerY, innerW, innerH, err = layoutFrameMetadataV1EngineLayoutFrameMetadata(node, target, innerX, innerY, innerW, innerH)
		if err != nil {
			return err
		}
	}
	innerX, innerY, innerW, innerH, err = alignContentAreaV1EngineLayoutAttributes(node, innerX, innerY, innerW, innerH)
	if err != nil {
		return err
	}
	if node.Tag == "frame" {
		finalizeFrameMetadataReservedStripV1EngineLayoutFrameMetadata(target, innerY, innerH)
	}
	setContentBoxV1EngineLayoutConstraints(target, innerX, innerY, innerW, innerH)

	switch node.Tag {
	case "frame", "container":
		if node.Attr("layout") == "horizontal" {
			loggerV1EngineSharedLogging.DEBUG(IULN005V1EngineLayoutNode, "branch frame/container horizontal", map[string]any{"tag": node.Tag})
			return layoutFlexHV1EngineLayoutFlow(node, target, innerX, innerY, innerW, innerH)
		} else {
			loggerV1EngineSharedLogging.DEBUG(IULN006V1EngineLayoutNode, "branch frame/container stack", map[string]any{"tag": node.Tag})
			return layoutStackV1EngineLayoutFlow(node, target, innerX, innerY, innerW, innerH)
		}
	case "row":
		loggerV1EngineSharedLogging.DEBUG(IULN007V1EngineLayoutNode, "branch row")
		return layoutRowV1EngineLayoutFlow(node, target, innerX, innerY, innerW, innerH)
	case "table", "entity":
		gInnerX := boxX + defaultGroupSideInsetV1EngineLayoutNode + pad.Left
		gInnerY := boxY + defaultGroupTopInsetV1EngineLayoutNode + pad.Top
		gInnerW := boxW - defaultGroupSideInsetV1EngineLayoutNode*2 - pad.Left - pad.Right
		gInnerH := boxH - defaultGroupTopInsetV1EngineLayoutNode - defaultGroupSideInsetV1EngineLayoutNode - pad.Top - pad.Bottom
		setContentBoxV1EngineLayoutConstraints(target, gInnerX, gInnerY, gInnerW, gInnerH)
		return layoutStackV1EngineLayoutFlow(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
	case "table-header", "table-row":
		return layoutFlexHV1EngineLayoutFlow(node, target, innerX, innerY, innerW, innerH)
	case "col":
		if node.Attr("layout") == "horizontal" {
			loggerV1EngineSharedLogging.DEBUG(IULN008V1EngineLayoutNode, "branch col horizontal")
			return layoutFlexHV1EngineLayoutFlow(node, target, innerX, innerY, innerW, innerH)
		} else {
			loggerV1EngineSharedLogging.DEBUG(IULN009V1EngineLayoutNode, "branch col stack")
			return layoutStackV1EngineLayoutFlow(node, target, innerX, innerY, innerW, innerH)
		}
	case "rectangle":
		setContentBoxV1EngineLayoutConstraints(target, boxX, boxY, boxW, boxH)
		return layoutRectangleV1EngineLayoutRectangle(node, target, boxX, boxY, boxW, boxH)
	default:
		// AWS グループタグおよびその他の未知タグ:
		// 子要素があればコンテナ, なければリーフとして扱う。
		kids := layoutKidsV1EngineLayoutNode(node)
		if len(kids) > 0 {
			loggerV1EngineSharedLogging.DEBUG(IULN010V1EngineLayoutNode, "branch default container", map[string]any{"tag": node.Tag, "children": len(kids)})
			// <item> / <spacer> のみの親はグループアイコン/ラベルがないので topInset を適用しない
			allItems := true
			for _, ch := range kids {
				if !IsItemLikeV1EngineLayoutAttributes(ch.Tag) {
					allItems = false
					break
				}
			}
			if allItems {
				loggerV1EngineSharedLogging.DEBUG(IULN011V1EngineLayoutNode, "branch all item children", map[string]any{"tag": node.Tag})
				return layoutRowV1EngineLayoutFlow(node, target, innerX, innerY, innerW, innerH)
			}
			// グループ inset は常に適用。ユーザー指定 padding はその上に加算する。
			// これにより class="pa-2" でヘッダー行と子要素が重なることを防ぐ。
			gInnerX := boxX + defaultGroupSideInsetV1EngineLayoutNode + pad.Left
			gInnerY := boxY + defaultGroupTopInsetV1EngineLayoutNode + pad.Top
			gInnerW := boxW - defaultGroupSideInsetV1EngineLayoutNode*2 - pad.Left - pad.Right
			gInnerH := boxH - defaultGroupTopInsetV1EngineLayoutNode - defaultGroupSideInsetV1EngineLayoutNode - pad.Top - pad.Bottom
			gInnerX, gInnerY, gInnerW, gInnerH, err = alignContentAreaV1EngineLayoutAttributes(node, gInnerX, gInnerY, gInnerW, gInnerH)
			if err != nil {
				return err
			}
			setContentBoxV1EngineLayoutConstraints(target, gInnerX, gInnerY, gInnerW, gInnerH)
			if node.Tag == "uml" && node.Attr("uml-kind") == "activity-diagram" && hasUMLActivityPartitionsV1EngineLayoutUmlActivity(node) {
				return layoutUMLActivityPartitionsV1EngineLayoutUmlActivity(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			}
			if node.Tag == "uml" && node.Attr("uml-kind") == "class-diagram" {
				return layoutUMLClassDiagramV1EngineLayoutUmlClass(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			}
			if node.Tag == "uml" && node.Attr("uml-kind") == "component-diagram" {
				return layoutUMLClassDiagramV1EngineLayoutUmlClass(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			}
			if node.Tag == "uml" && node.Attr("uml-kind") == "state-machine-diagram" {
				return layoutUMLStateMachineDiagramV1EngineLayoutUmlActivity(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			}
			if node.Attr("layout") == "staggered" {
				loggerV1EngineSharedLogging.DEBUG(IULN012V1EngineLayoutNode, "branch staggered", map[string]any{"tag": node.Tag})
				return layoutStaggerV1EngineLayoutFlow(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			} else if node.Attr("layout") == "horizontal" {
				loggerV1EngineSharedLogging.DEBUG(IULN013V1EngineLayoutNode, "branch horizontal", map[string]any{"tag": node.Tag})
				return layoutFlexHV1EngineLayoutFlow(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			} else {
				loggerV1EngineSharedLogging.DEBUG(IULN014V1EngineLayoutNode, "branch stack", map[string]any{"tag": node.Tag})
				return layoutStackV1EngineLayoutFlow(node, target, gInnerX, gInnerY, gInnerW, gInnerH)
			}
		} else {
			loggerV1EngineSharedLogging.DEBUG(IULN015V1EngineLayoutNode, "branch leaf", map[string]any{"tag": node.Tag})
			// The border box was resolved above. A leaf has no child layout, but its
			// content box still reflects padding/content alignment; do not replace
			// the border box with that smaller content box.
		}
	}
	return nil
}
