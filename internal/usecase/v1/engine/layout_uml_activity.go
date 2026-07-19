package engine

import (
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type umlActivityLaneV1EngineLayoutUmlActivity struct {
	id    string
	nodes []*entity.Node
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
	case "decision", "merge":
		return 78, 78
	case "fork", "join":
		return math.Min(maxW, 120), 24
	default:
		return math.Min(maxW, 180), 58
	}
}

func umlActivityNodeStepV1EngineLayoutUmlActivity(maxNodes int, usableH float64) float64 {
	if maxNodes <= 1 {
		return 0
	}
	return math.Max(82, math.Min(132, usableH/float64(maxNodes)))
}
