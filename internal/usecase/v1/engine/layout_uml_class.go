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
	cols := int(math.Ceil(math.Sqrt(float64(len(children)) * 1.25)))
	if cols < 1 {
		cols = 1
	}
	if cols > 3 {
		cols = 3
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
		width := math.Min(math.Max(360, maxW*0.86), math.Max(260, maxW-16))
		height := math.Min(math.Max(220, 120+math.Ceil(childCount/2)*132), math.Max(180, maxH-16))
		return width, height
	}
	lines := strings.Count(labelOfV1EngineLayoutAttributes(node), "\n") + 1
	width := math.Min(math.Max(170, maxW*0.58), 230)
	height := math.Min(math.Max(78, 32+float64(lines)*18), 160)
	if maxW > 0 {
		width = math.Min(width, math.Max(MinBoxWidthV1EngineLayoutFlow, maxW-24))
	}
	if maxH > 0 {
		height = math.Min(height, math.Max(MinBoxHeightV1EngineLayoutFlow, maxH-24))
	}
	return width, height
}
