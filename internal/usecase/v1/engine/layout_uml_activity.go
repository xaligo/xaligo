package engine

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type umlActivityLaneV1EngineLayoutUmlActivity struct {
	id    string
	nodes []*entity.Node
}

type umlStateMachineRowV1EngineLayoutUmlActivity struct {
	row   int
	nodes []*entity.Node
}

type umlStateMachineRelationV1EngineLayoutUmlActivity struct {
	src string
	dst string
}

func hasUMLActivityPartitionsV1EngineLayoutUmlActivity(node *entity.Node) bool {
	if node == nil {
		return false
	}
	for _, child := range layoutKidsV1EngineLayoutNode(node) {
		if strings.TrimSpace(child.Attr("uml-partition-id")) != "" {
			return true
		}
	}
	return false
}

func layoutUMLActivityPartitionsV1EngineLayoutUmlActivity(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	lanes := umlActivityLanesV1EngineLayoutUmlActivity(node)
	if len(lanes) == 0 {
		return layoutStackV1EngineLayoutFlow(node, target, x, y, w, h)
	}
	if umlActivityHorizontalLanesV1EngineLayoutUmlActivity(node) {
		return layoutUMLActivityHorizontalPartitionsV1EngineLayoutUmlActivity(node, target, x, y, w, h, lanes)
	}
	laneW := w / float64(len(lanes))
	usableY := y + 52
	usableH := math.Max(MinBoxHeightV1EngineLayoutFlow, h-60)
	maxNodes := 1
	for _, lane := range lanes {
		if len(lane.nodes) > maxNodes {
			maxNodes = len(lane.nodes)
		}
	}
	step := umlActivityNodeStepV1EngineLayoutUmlActivity(maxNodes, usableH)
	for laneIndex, lane := range lanes {
		laneX := x + float64(laneIndex)*laneW
		laneInnerW := math.Max(MinBoxWidthV1EngineLayoutFlow, laneW-24)
		centerX := laneX + laneW/2
		for nodeIndex, child := range lane.nodes {
			nodeW, nodeH := umlActivityNodeSizeV1EngineLayoutUmlActivity(child, laneInnerW)
			nodeX := centerX - nodeW/2
			nodeY := usableY + float64(nodeIndex)*step
			box := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
			if err := layoutNodeV1EngineLayoutNode(child, box, nodeX, nodeY, nodeW, nodeH); err != nil {
				return err
			}
			target.Children = append(target.Children, box)
		}
	}
	return nil
}

func layoutUMLActivityHorizontalPartitionsV1EngineLayoutUmlActivity(node *entity.Node, target *entity.Box, x, y, w, h float64, lanes []umlActivityLaneV1EngineLayoutUmlActivity) error {
	laneY := y + 8
	laneH := math.Max(MinBoxHeightV1EngineLayoutFlow, h-16) / float64(len(lanes))
	usableX := x + 156
	usableW := math.Max(MinBoxWidthV1EngineLayoutFlow, w-172)
	maxNodes := 1
	for _, lane := range lanes {
		if len(lane.nodes) > maxNodes {
			maxNodes = len(lane.nodes)
		}
	}
	step := umlActivityNodeStepHorizontalV1EngineLayoutUmlActivity(maxNodes, usableW)
	for laneIndex, lane := range lanes {
		currentLaneY := laneY + float64(laneIndex)*laneH
		laneInnerH := math.Max(MinBoxHeightV1EngineLayoutFlow, laneH-24)
		centerY := currentLaneY + laneH/2
		for nodeIndex, child := range lane.nodes {
			nodeW, nodeH := umlActivityNodeSizeV1EngineLayoutUmlActivity(child, usableW)
			if nodeH > laneInnerH {
				nodeH = laneInnerH
			}
			nodeX := usableX + float64(nodeIndex)*step
			nodeY := centerY - nodeH/2
			box := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
			if err := layoutNodeV1EngineLayoutNode(child, box, nodeX, nodeY, nodeW, nodeH); err != nil {
				return err
			}
			target.Children = append(target.Children, box)
		}
	}
	return nil
}

func layoutUMLStateMachineDiagramV1EngineLayoutUmlActivity(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	children := layoutKidsV1EngineLayoutNode(node)
	if len(children) == 0 {
		return nil
	}
	if rows, ok := umlStateMachineRowsV1EngineLayoutUmlActivity(children); ok {
		return layoutUMLStateMachineRowsV1EngineLayoutUmlActivity(node, target, x, y, w, h, rows)
	}
	if strings.TrimSpace(node.Attr("direction")) == "down" {
		return layoutUMLStateMachineDiagramDownV1EngineLayoutUmlActivity(node, target, x, y, w, h, children)
	}
	step := umlActivityNodeStepHorizontalV1EngineLayoutUmlActivity(len(children), w)
	startX := x
	if len(children) > 1 {
		usedW := step*float64(len(children)-1) + 180
		if usedW < w {
			startX = x + (w-usedW)/2
		}
	}
	centerY := y + h/2
	for childIndex, child := range children {
		nodeW, nodeH := umlActivityNodeSizeV1EngineLayoutUmlActivity(child, 180)
		nodeX := startX + float64(childIndex)*step
		nodeY := centerY - nodeH/2
		box := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, box, nodeX, nodeY, nodeW, nodeH); err != nil {
			return err
		}
		target.Children = append(target.Children, box)
	}
	return nil
}

func umlStateMachineRowsV1EngineLayoutUmlActivity(children []*entity.Node) ([]umlStateMachineRowV1EngineLayoutUmlActivity, bool) {
	indices := map[int]int{}
	var rows []umlStateMachineRowV1EngineLayoutUmlActivity
	hasRow := false
	for _, child := range children {
		row := int(attrFloatV1EngineLayoutAttributes(child.Attr("row"), 1))
		if row < 1 {
			row = 1
		}
		if strings.TrimSpace(child.Attr("row")) != "" {
			hasRow = true
		}
		index, ok := indices[row]
		if !ok {
			index = len(rows)
			indices[row] = index
			rows = append(rows, umlStateMachineRowV1EngineLayoutUmlActivity{row: row})
		}
		rows[index].nodes = append(rows[index].nodes, child)
	}
	if !hasRow {
		return nil, false
	}
	sort.SliceStable(rows, func(left, right int) bool {
		return rows[left].row < rows[right].row
	})
	return rows, true
}

func layoutUMLStateMachineRowsV1EngineLayoutUmlActivity(node *entity.Node, target *entity.Box, x, y, w, h float64, rows []umlStateMachineRowV1EngineLayoutUmlActivity) error {
	rowH := math.Max(MinBoxHeightV1EngineLayoutFlow, h/float64(len(rows)))
	columns := umlStateMachineGridColumnsV1EngineLayoutUmlActivity(node, rows)
	if columns < 1 {
		columns = 1
	}
	cellW := w / float64(columns)
	for rowIndex, row := range rows {
		centerY := y + float64(rowIndex)*rowH + rowH/2
		for nodeIndex, child := range row.nodes {
			col := umlStateMachineNodeColumnV1EngineLayoutUmlActivity(child)
			if col < 1 || col > columns {
				col = nodeIndex + 1
			}
			nodeW, nodeH := umlActivityNodeSizeV1EngineLayoutUmlActivity(child, math.Max(40, cellW-28))
			nodeX := x + (float64(col)-0.5)*cellW - nodeW/2
			nodeY := centerY - nodeH/2
			box := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
			if err := layoutNodeV1EngineLayoutNode(child, box, nodeX, nodeY, nodeW, nodeH); err != nil {
				return err
			}
			target.Children = append(target.Children, box)
		}
	}
	return nil
}

func umlStateMachineGridColumnsV1EngineLayoutUmlActivity(node *entity.Node, rows []umlStateMachineRowV1EngineLayoutUmlActivity) int {
	columns := 0
	assignments := map[string]int{}
	for _, row := range rows {
		if len(row.nodes) > columns {
			columns = len(row.nodes)
		}
		for _, child := range row.nodes {
			if col := umlStateMachineExplicitNodeColumnV1EngineLayoutUmlActivity(child); col > 0 {
				assignments[umlStateMachineNodeReferenceV1EngineLayoutUmlActivity(child)] = col
				if col > columns {
					columns = col
				}
			}
		}
	}
	relations := umlStateMachineRelationsV1EngineLayoutUmlActivity(node)
	for changed := true; changed; {
		changed = false
		for _, relation := range relations {
			if relation.src == "" || relation.dst == "" {
				continue
			}
			srcCol, hasSrc := assignments[relation.src]
			dstCol, hasDst := assignments[relation.dst]
			switch {
			case hasSrc && !hasDst:
				assignments[relation.dst] = srcCol
				changed = true
			case hasDst && !hasSrc:
				assignments[relation.src] = dstCol
				changed = true
			}
		}
	}
	for _, row := range rows {
		used := map[int]bool{}
		for _, child := range row.nodes {
			ref := umlStateMachineNodeReferenceV1EngineLayoutUmlActivity(child)
			if col := assignments[ref]; col > 0 && !used[col] {
				child.Attrs["col"] = strconv.Itoa(col)
				used[col] = true
				continue
			}
			col := umlStateMachinePreferredColumnV1EngineLayoutUmlActivity(ref, assignments, relations)
			if col < 1 || used[col] {
				col = 1
			}
			for used[col] {
				col++
			}
			child.Attrs["col"] = strconv.Itoa(col)
			assignments[ref] = col
			used[col] = true
			if col > columns {
				columns = col
			}
		}
	}
	return columns
}

func umlStateMachinePreferredColumnV1EngineLayoutUmlActivity(ref string, assignments map[string]int, relations []umlStateMachineRelationV1EngineLayoutUmlActivity) int {
	if ref == "" {
		return 0
	}
	columns := 0
	count := 0
	for _, relation := range relations {
		switch ref {
		case relation.src:
			if col := assignments[relation.dst]; col > 0 {
				columns += col
				count++
			}
		case relation.dst:
			if col := assignments[relation.src]; col > 0 {
				columns += col
				count++
			}
		}
	}
	if count == 0 {
		return 0
	}
	return int(math.Round(float64(columns) / float64(count)))
}

func umlStateMachineNodeColumnV1EngineLayoutUmlActivity(node *entity.Node) int {
	if col := umlStateMachineExplicitNodeColumnV1EngineLayoutUmlActivity(node); col > 0 {
		return col
	}
	return int(attrFloatV1EngineLayoutAttributes(node.Attr("col"), 0))
}

func umlStateMachineExplicitNodeColumnV1EngineLayoutUmlActivity(node *entity.Node) int {
	if strings.TrimSpace(node.Attr("col")) == "" {
		return 0
	}
	col := int(attrFloatV1EngineLayoutAttributes(node.Attr("col"), 0))
	if col < 1 {
		return 0
	}
	return col
}

func umlStateMachineNodeReferenceV1EngineLayoutUmlActivity(node *entity.Node) string {
	if ref := strings.TrimSpace(node.Attr("ref")); ref != "" {
		return ref
	}
	return strings.TrimSpace(node.Attr("id"))
}

func umlStateMachineRelationsV1EngineLayoutUmlActivity(node *entity.Node) []umlStateMachineRelationV1EngineLayoutUmlActivity {
	if node == nil {
		return nil
	}
	var relations []umlStateMachineRelationV1EngineLayoutUmlActivity
	for _, child := range node.Children {
		if child.Tag != "connection" || strings.TrimSpace(child.Attr("uml-diagram-kind")) != "state-machine-diagram" {
			continue
		}
		relations = append(relations, umlStateMachineRelationV1EngineLayoutUmlActivity{
			src: strings.TrimSpace(child.Attr("uml-src-ref")),
			dst: strings.TrimSpace(child.Attr("uml-dst-ref")),
		})
	}
	for _, child := range layoutKidsV1EngineLayoutNode(node) {
		src := umlStateMachineNodeReferenceV1EngineLayoutUmlActivity(child)
		for _, dst := range splitCSVV1EngineLayoutUmlActivity(child.Attr("uml-related-refs")) {
			relations = append(relations, umlStateMachineRelationV1EngineLayoutUmlActivity{src: src, dst: dst})
		}
	}
	return relations
}

func splitCSVV1EngineLayoutUmlActivity(value string) []string {
	var result []string
	for _, part := range strings.Split(value, ",") {
		if part = strings.TrimSpace(part); part != "" {
			result = append(result, part)
		}
	}
	return result
}

func layoutUMLStateMachineDiagramDownV1EngineLayoutUmlActivity(node *entity.Node, target *entity.Box, x, y, w, h float64, children []*entity.Node) error {
	step := umlActivityNodeStepV1EngineLayoutUmlActivity(len(children), h)
	startY := y
	if len(children) > 1 {
		usedH := step*float64(len(children)-1) + 58
		if usedH < h {
			startY = y + (h-usedH)/2
		}
	}
	centerX := x + w/2
	for childIndex, child := range children {
		nodeW, nodeH := umlActivityNodeSizeV1EngineLayoutUmlActivity(child, math.Min(w, 180))
		nodeX := centerX - nodeW/2
		nodeY := startY + float64(childIndex)*step
		box := &entity.Box{ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, box, nodeX, nodeY, nodeW, nodeH); err != nil {
			return err
		}
		target.Children = append(target.Children, box)
	}
	return nil
}

func umlActivityHorizontalLanesV1EngineLayoutUmlActivity(node *entity.Node) bool {
	return strings.TrimSpace(node.Attr("layout")) == "horizontal" || strings.TrimSpace(node.Attr("lanes")) == "horizontal"
}

func umlActivityLanesV1EngineLayoutUmlActivity(node *entity.Node) []umlActivityLaneV1EngineLayoutUmlActivity {
	indices := map[string]int{}
	var lanes []umlActivityLaneV1EngineLayoutUmlActivity
	for _, child := range layoutKidsV1EngineLayoutNode(node) {
		partition := strings.TrimSpace(child.Attr("uml-partition-id"))
		if partition == "" {
			partition = "_unpartitioned"
		}
		index, ok := indices[partition]
		if !ok {
			index = len(lanes)
			indices[partition] = index
			lanes = append(lanes, umlActivityLaneV1EngineLayoutUmlActivity{id: partition})
		}
		lanes[index].nodes = append(lanes[index].nodes, child)
	}
	return lanes
}

func umlActivityNodeSizeV1EngineLayoutUmlActivity(node *entity.Node, maxW float64) (float64, float64) {
	switch strings.TrimSpace(node.Attr("uml-element-kind")) {
	case "initial", "final":
		return 34, 34
	case "decision", "merge", "choice", "history":
		return 78, 78
	case "fork", "join":
		return math.Min(maxW, 120), 24
	default:
		if strings.TrimSpace(node.Attr("uml-diagram-kind")) == "state-machine-diagram" && strings.TrimSpace(node.Attr("uml-element-kind")) == "state" {
			return math.Min(maxW, 190), 112
		}
		return math.Min(maxW, 180), 58
	}
}

func umlActivityNodeStepV1EngineLayoutUmlActivity(maxNodes int, usableH float64) float64 {
	if maxNodes <= 1 {
		return 0
	}
	return math.Max(82, math.Min(132, usableH/float64(maxNodes)))
}

func umlActivityNodeStepHorizontalV1EngineLayoutUmlActivity(maxNodes int, usableW float64) float64 {
	if maxNodes <= 1 {
		return 0
	}
	return math.Max(132, math.Min(220, usableW/float64(maxNodes)))
}
