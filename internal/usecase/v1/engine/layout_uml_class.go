package engine

import (
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func layoutUMLClassDiagramV1EngineLayoutUmlClass(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	children := layoutKidsV1EngineLayoutNode(node)
	if len(children) == 0 {
		return nil
	}
	cols, configuredGrid := umlClassGridColumnsV1EngineLayoutUmlClass(node, len(children))
	maxCols := 3
	if node.Attr("uml-diagram-kind") == "class-diagram" && node.Attr("uml-element-kind") == "package" {
		maxCols = 2
		if w < 560 {
			maxCols = 1
		}
	}
	if cols < 1 {
		cols = 1
	}
	if !configuredGrid && cols > maxCols {
		cols = maxCols
	}
	rows := int(math.Ceil(float64(len(children)) / float64(cols)))
	cellW := w / float64(cols)
	cellH := h / float64(rows)
	for index, child := range children {
		row := index / cols
		col := index % cols
		nodeW, nodeH := umlClassNodeSizeV1EngineLayoutUmlClass(child, cellW, cellH)
		nodeX := x + float64(col)*cellW + (cellW-nodeW)/2
		nodeY := y + float64(row)*cellH + (cellH-nodeH)/2
		box := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, box, nodeX, nodeY, nodeW, nodeH); err != nil {
			return err
		}
		target.Children = append(target.Children, box)
	}
	return nil
}

func umlClassNodeSizeV1EngineLayoutUmlClass(node *entity.Node, maxW, maxH float64) (float64, float64) {
	if node.Attr("uml-diagram-kind") == "class-diagram" && node.Attr("uml-element-kind") == "package" {
		childCount := math.Max(1, float64(len(layoutKidsV1EngineLayoutNode(node))))
		rows := math.Ceil(childCount / 2)
		width := math.Max(260, maxW-16)
		height := math.Min(math.Max(260, 118+rows*170), math.Max(220, maxH-16))
		return width, height
	}
	width := math.Min(math.Max(240, maxW*0.82), 300)
	height := math.Min(umlClassNodeHeightV1EngineLayoutUmlClass(node), 220)
	if maxW > 0 {
		width = math.Min(width, math.Max(MinBoxWidthV1EngineLayoutFlow, maxW-24))
	}
	if maxH > 0 {
		height = math.Min(height, math.Max(MinBoxHeightV1EngineLayoutFlow, maxH-24))
	}
	return width, height
}

func umlClassGridColumnsV1EngineLayoutUmlClass(node *entity.Node, childCount int) (int, bool) {
	if childCount <= 0 {
		return 1, false
	}
	if configured := attrFloatV1EngineLayoutAttributes(node.Attr("grid"), 0); configured > 0 {
		cols := int(math.Round(configured))
		if cols < 1 {
			return 1, true
		}
		if cols > childCount {
			return childCount, true
		}
		return cols, true
	}
	return int(math.Ceil(math.Sqrt(float64(childCount) * 1.25))), false
}

func umlClassNodeHeightV1EngineLayoutUmlClass(node *entity.Node) float64 {
	lines := strings.Count(labelOfV1EngineLayoutAttributes(node), "\n") + 1
	headerLines := attrFloatV1EngineLayoutAttributes(node.Attr("uml-class-header-lines"), 1)
	attributeLines := attrFloatV1EngineLayoutAttributes(node.Attr("uml-class-attribute-lines"), 0)
	operationLines := attrFloatV1EngineLayoutAttributes(node.Attr("uml-class-operation-lines"), 0)
	if attributeLines == 0 && operationLines == 0 {
		return math.Max(96, 38+float64(lines)*20)
	}
	headerH := math.Max(34, 10+headerLines*16.8)
	bodyH := 18 + attributeLines*18
	if attributeLines > 0 && operationLines > 0 {
		bodyH += 8
	}
	bodyH += operationLines * 18
	return math.Max(108, headerH+bodyH)
}
