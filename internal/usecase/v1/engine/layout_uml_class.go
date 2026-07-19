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
	componentWidth := attrFloatV1EngineLayoutAttributes(node.Attr("component-width"), 0)
	componentHeight := attrFloatV1EngineLayoutAttributes(node.Attr("component-height"), 0)
	packageGrid := umlClassPackageGridV1EngineLayoutUmlClass(node, children)
	cols, configuredGrid := umlClassGridColumnsV1EngineLayoutUmlClass(node, len(children), w, h, packageGrid)
	maxCols := 3
	if packageGrid {
		maxCols = len(children)
	} else if node.Attr("uml-diagram-kind") == "class-diagram" && node.Attr("uml-element-kind") == "package" {
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
	cellHeights := make([]float64, rows)
	componentDiagram := node.Attr("uml-kind") == "component-diagram"
	if componentDiagram {
		for index, child := range children {
			_, nodeH := umlClassNodeSizeV1EngineLayoutUmlClass(child, cellW, 0, componentWidth, componentHeight)
			row := index / cols
			cellHeights[row] = math.Max(cellHeights[row], nodeH+24)
		}
	} else {
		for row := range cellHeights {
			cellHeights[row] = h / float64(rows)
		}
	}
	rowOffsets := make([]float64, rows)
	for row := 1; row < rows; row++ {
		rowOffsets[row] = rowOffsets[row-1] + cellHeights[row-1]
	}
	for index, child := range children {
		row := index / cols
		col := index % cols
		cellH := cellHeights[row]
		maxNodeH := cellH
		if componentDiagram {
			maxNodeH = 0
		}
		nodeW, nodeH := umlClassNodeSizeV1EngineLayoutUmlClass(child, cellW, maxNodeH, componentWidth, componentHeight)
		if componentDiagram {
			if err := validateUMLComponentInterfaceWidthV1EngineLayoutUmlClass(child, nodeW); err != nil {
				return err
			}
		}
		if packageGrid && isUMLClassPackageNodeV1EngineLayoutUmlClass(child) {
			nodeW = math.Max(1, cellW-16)
			nodeH = math.Max(1, cellH-16)
		}
		nodeX := x + float64(col)*cellW + (cellW-nodeW)/2
		nodeY := y + rowOffsets[row] + (cellH-nodeH)/2
		box := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, box, nodeX, nodeY, nodeW, nodeH); err != nil {
			return err
		}
		target.Children = append(target.Children, box)
	}
	return nil
}

func validateUMLComponentInterfaceWidthV1EngineLayoutUmlClass(node *entity.Node, componentWidth float64) error {
	interfaceWidth := attrFloatV1EngineLayoutAttributes(node.Attr("interface-width"), 0)
	if interfaceWidth <= 0 {
		return nil
	}
	availableWidth := componentWidth + 21 - 10 - 14 - 4
	minimumRemainder := 0.0
	for _, line := range strings.Split(node.Attr("uml-component-interfaces"), "\n") {
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) > 1 && strings.TrimSpace(parts[1]) != "" {
			minimumRemainder = 36
			break
		}
	}
	maximumWidth := availableWidth - minimumRemainder
	if interfaceWidth > maximumWidth {
		return newLayoutErrorV1EngineLayoutValidation(node, "interface-width %.6g exceeds the available component interface-name width %.6g", interfaceWidth, maximumWidth)
	}
	return nil
}

func umlClassNodeSizeV1EngineLayoutUmlClass(node *entity.Node, maxW, maxH, componentWidth, componentHeight float64) (float64, float64) {
	if node.Attr("uml-diagram-kind") == "class-diagram" && node.Attr("uml-element-kind") == "package" {
		childCount := math.Max(1, float64(len(layoutKidsV1EngineLayoutNode(node))))
		rows := math.Ceil(childCount / 2)
		width := math.Max(260, maxW-16)
		height := math.Min(math.Max(260, 118+rows*170), math.Max(220, maxH-16))
		return width, height
	}
	width := math.Min(math.Max(240, maxW*0.82), 300)
	if node.Attr("uml-diagram-kind") == "component-diagram" && node.Attr("uml-element-kind") == "component" {
		width = umlComponentNodeWidthV1EngineLayoutUmlClass(node, width, componentWidth)
	}
	heightLimit := 220.0
	if node.Attr("uml-diagram-kind") == "component-diagram" && node.Attr("uml-element-kind") == "component" {
		if maxH > 0 {
			heightLimit = math.Max(heightLimit, maxH-24)
		} else {
			heightLimit = math.Inf(1)
		}
	}
	height := math.Min(umlClassNodeHeightV1EngineLayoutUmlClass(node, componentHeight), heightLimit)
	if maxW > 0 {
		width = math.Min(width, math.Max(MinBoxWidthV1EngineLayoutFlow, maxW-24))
	}
	if maxH > 0 {
		height = math.Min(height, math.Max(MinBoxHeightV1EngineLayoutFlow, maxH-24))
	}
	return width, height
}

func umlComponentNodeWidthV1EngineLayoutUmlClass(node *entity.Node, fallback, diagramWidth float64) float64 {
	if width := attrFloatV1EngineLayoutAttributes(node.Attr("width"), 0); width > 0 {
		return width
	}
	if width := attrFloatV1EngineLayoutAttributes(node.Attr("component-width"), 0); width > 0 {
		return width
	}
	if diagramWidth > 0 {
		return diagramWidth
	}
	return fallback
}

func umlClassGridColumnsV1EngineLayoutUmlClass(node *entity.Node, childCount int, width, height float64, balanceToFrame bool) (int, bool) {
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
	if balanceToFrame {
		return balancedUMLClassGridColumnsV1EngineLayoutUmlClass(childCount, width, height), false
	}
	return int(math.Ceil(math.Sqrt(float64(childCount) * 1.25))), false
}

func balancedUMLClassGridColumnsV1EngineLayoutUmlClass(childCount int, width, height float64) int {
	if childCount <= 1 || width <= 0 || height <= 0 {
		return 1
	}
	idealAspect := 1.15
	bestCols := 1
	bestScore := math.Inf(1)
	bestEmpty := childCount
	for cols := 1; cols <= childCount; cols++ {
		rows := int(math.Ceil(float64(childCount) / float64(cols)))
		cellW := width / float64(cols)
		cellH := height / float64(rows)
		if cellW <= 0 || cellH <= 0 {
			continue
		}
		empty := cols*rows - childCount
		score := math.Abs(math.Log((cellW/cellH)/idealAspect)) + float64(empty)*0.35
		if score < bestScore || (math.Abs(score-bestScore) < 0.001 && empty < bestEmpty) {
			bestCols = cols
			bestScore = score
			bestEmpty = empty
		}
	}
	return bestCols
}

func umlClassPackageGridV1EngineLayoutUmlClass(node *entity.Node, children []*entity.Node) bool {
	if node == nil || node.Tag != "uml" || node.Attr("uml-kind") != "class-diagram" || len(children) == 0 {
		return false
	}
	for _, child := range children {
		if !isUMLClassPackageNodeV1EngineLayoutUmlClass(child) {
			return false
		}
	}
	return true
}

func isUMLClassPackageNodeV1EngineLayoutUmlClass(node *entity.Node) bool {
	return node != nil && node.Attr("uml-diagram-kind") == "class-diagram" && node.Attr("uml-element-kind") == "package"
}

func umlClassNodeHeightV1EngineLayoutUmlClass(node *entity.Node, componentHeight float64) float64 {
	lines := strings.Count(labelOfV1EngineLayoutAttributes(node), "\n") + 1
	if node.Attr("uml-diagram-kind") == "component-diagram" && node.Attr("uml-element-kind") == "component" {
		if height := attrFloatV1EngineLayoutAttributes(node.Attr("height"), 0); height > 0 {
			return height
		}
		if height := attrFloatV1EngineLayoutAttributes(node.Attr("component-height"), 0); height > 0 {
			return height
		}
		if componentHeight > 0 {
			return componentHeight
		}
		interfaceLines := strings.Count(strings.TrimSpace(node.Attr("uml-component-interfaces")), "\n") + 1
		if strings.TrimSpace(node.Attr("uml-component-interfaces")) == "" {
			interfaceLines = 0
		}
		connectionLines := int(attrFloatV1EngineLayoutAttributes(node.Attr("uml-component-connection-count"), 0))
		if interfaceLines > 0 {
			fanoutExtra := attrFloatV1EngineLayoutAttributes(node.Attr("uml-component-interface-fanout-extra"), 0)
			return math.Max(88, 64+float64(interfaceLines)*24+fanoutExtra*22)
		}
		if connectionLines > 0 {
			return math.Max(88, 64+float64(connectionLines)*18)
		}
	}
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
