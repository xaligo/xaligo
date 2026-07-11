package repository

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

const genericEndpointIconID = "xaligo-generic-endpoint"

type isoflowConnectorInput struct {
	id              string
	sourceElementID string
	targetElementID string
	element         entity.Element
	crossFrame      bool
}

// isoflowConnectorInputs recovers the logical endpoint pair represented by a
// scene connector. Cross-frame V1 connections are serialized as two local
// stubs, so their shared logical ID and endpoint metadata are used to emit one
// Isoflow connector.
func isoflowConnectorInputs(elements []entity.Element) []isoflowConnectorInput {
	inputs := make([]isoflowConnectorInput, 0)
	seenCrossFrame := map[string]bool{}
	for _, element := range elements {
		if element.IsDeleted || (element.Type != "arrow" && element.Type != "line") {
			continue
		}
		if element.CustomData != nil && element.CustomData.ConnectorCrossFrame {
			source := strings.TrimSpace(element.CustomData.ConnectorSourceElementID)
			target := strings.TrimSpace(element.CustomData.ConnectorDestinationElementID)
			if source == "" || target == "" {
				continue
			}
			logicalID := strings.TrimSpace(element.CustomData.ConnectorLogicalID)
			if logicalID == "" {
				logicalID = element.ID
			}
			key := logicalID + "\x00" + source + "\x00" + target
			if seenCrossFrame[key] {
				continue
			}
			seenCrossFrame[key] = true
			inputs = append(inputs, isoflowConnectorInput{id: logicalID, sourceElementID: source, targetElementID: target, element: element, crossFrame: true})
			continue
		}
		if element.StartBinding == nil || element.EndBinding == nil {
			continue
		}
		inputs = append(inputs, isoflowConnectorInput{
			id:              element.ID,
			sourceElementID: element.StartBinding.ElementID,
			targetElementID: element.EndBinding.ElementID,
			element:         element,
		})
	}
	return inputs
}

// appendIsoflowEndpointNodes emits only non-item scene elements that are
// actually referenced by a connector. Items have already reserved their
// positions, which preserves the existing item ordering and tile placement.
func appendIsoflowEndpointNodes(
	elements []entity.Element,
	inputs []isoflowConnectorInput,
	placer *tilePlacer,
	nodeIDByElementID map[string]string,
	modelItems *[]entity.IsoflowModelItem,
	viewItems *[]entity.IsoflowViewItem,
) bool {
	appended := false
	referenced := map[string]bool{}
	for _, input := range inputs {
		referenced[input.sourceElementID] = true
		referenced[input.targetElementID] = true
	}
	usedNodeIDs := map[string]bool{}
	for _, id := range nodeIDByElementID {
		usedNodeIDs[id] = true
	}
	labels := isoflowEndpointLabels(elements)
	for _, element := range elements {
		if !referenced[element.ID] || element.IsDeleted || !isoflowEndpointElement(element) {
			continue
		}
		if _, exists := nodeIDByElementID[element.ID]; exists {
			continue
		}
		nodeID := uniqueIsoflowEndpointNodeID(element.ID, usedNodeIDs)
		usedNodeIDs[nodeID] = true
		nodeIDByElementID[element.ID] = nodeID
		*modelItems = append(*modelItems, entity.IsoflowModelItem{
			ID: nodeID, Name: isoflowEndpointName(element, labels), IsoflowIcon: genericEndpointIconID,
		})
		*viewItems = append(*viewItems, entity.IsoflowViewItem{
			ID:          nodeID,
			Tile:        placer.place(pixelToTile(element.X+element.Width/2, element.Y+element.Height/2)),
			LabelHeight: defaultLabelHeight,
		})
		appended = true
	}
	return appended
}

func genericIsoflowEndpointIcon() entity.IsoflowIcon {
	// Isoflow requires every model item to reference a defined icon. Keep one
	// deterministic built-in icon for V1 shapes that have no target icon model.
	const svg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 96 96"><path d="M48 8 84 28 48 48 12 28Z" fill="#e2e8f0" stroke="#475569" stroke-width="4"/><path d="M12 28v40l36 20V48Z" fill="#cbd5e1" stroke="#475569" stroke-width="4"/><path d="M84 28v40L48 88V48Z" fill="#94a3b8" stroke="#475569" stroke-width="4"/></svg>`
	return entity.IsoflowIcon{
		ID: genericEndpointIconID, Name: "Generic endpoint", URL: share.SVGDataURLFromBytes([]byte(svg)),
		Collection: "xaligo", IsIsometric: true,
	}
}

func isoflowEndpointElement(element entity.Element) bool {
	if element.Width <= 0 || element.Height <= 0 {
		return false
	}
	switch element.Type {
	case "rectangle":
		return true
	case "frame":
		return element.ID != "paper-frame"
	default:
		return false
	}
}

func isoflowEndpointLabels(elements []entity.Element) map[string]string {
	labels := map[string]string{}
	for _, element := range elements {
		if element.IsDeleted || element.Type != "text" || strings.TrimSpace(element.Text) == "" {
			continue
		}
		switch {
		case strings.HasSuffix(element.ID, "-label"):
			labels[strings.TrimSuffix(element.ID, "-label")+"-rect"] = strings.TrimSpace(element.Text)
		case strings.HasSuffix(element.ID, "-text"):
			labels[strings.TrimSuffix(element.ID, "-text")+"-rect"] = strings.TrimSpace(element.Text)
		}
	}
	return labels
}

func isoflowEndpointName(element entity.Element, labels map[string]string) string {
	if label := strings.TrimSpace(labels[element.ID]); label != "" {
		return label
	}
	if strings.HasPrefix(element.ID, "paper-frame-") {
		return strings.TrimPrefix(element.ID, "paper-frame-")
	}
	if name := strings.TrimSuffix(element.ID, "-rect"); name != "" {
		return name
	}
	return element.ID
}

func uniqueIsoflowEndpointNodeID(elementID string, used map[string]bool) string {
	base := elementID + "-node"
	if !used[base] {
		return base
	}
	for index := 2; ; index++ {
		candidate := fmt.Sprintf("%s-%d", base, index)
		if !used[candidate] {
			return candidate
		}
	}
}

func resolveIsoflowNodeID(nodeIDByElementID map[string]string, elementID string) (string, bool) {
	if nodeID, ok := nodeIDByElementID[elementID]; ok {
		return nodeID, true
	}
	nodeID, ok := nodeIDByElementID[share.ItemNodeID(elementID)]
	return nodeID, ok
}

// Isoflow's upstream connector schema has no arbitrary data field. V1 kind,
// arrowheads, scale, grid, and Excalidraw fixedPoint therefore cannot be
// serialized losslessly. Explicit bend geometry can be represented natively as
// tile anchors, so resolved interior scene points are retained for that case.
func isoflowConnectorAnchors(input isoflowConnectorInput, source, target string) []entity.IsoflowConnectorAnchor {
	anchors := []entity.IsoflowConnectorAnchor{
		{ID: input.id + "-source", Ref: entity.IsoflowAnchorRef{Item: source}},
	}
	if !input.crossFrame && input.element.CustomData != nil && strings.TrimSpace(input.element.CustomData.ConnectorBends) != "" {
		previous := entity.IsoflowCoords{}
		hasPrevious := false
		bendIndex := 0
		for index := 1; index+1 < len(input.element.Points); index++ {
			point := input.element.Points[index]
			if len(point) < 2 {
				continue
			}
			tile := pixelToTile(input.element.X+point[0], input.element.Y+point[1])
			if hasPrevious && tile == previous {
				continue
			}
			hasPrevious = true
			previous = tile
			bendIndex++
			bendTile := tile
			anchors = append(anchors, entity.IsoflowConnectorAnchor{
				ID:  fmt.Sprintf("%s-bend-%d", input.id, bendIndex),
				Ref: entity.IsoflowAnchorRef{Tile: &bendTile},
			})
		}
	}
	return append(anchors, entity.IsoflowConnectorAnchor{
		ID: input.id + "-target", Ref: entity.IsoflowAnchorRef{Item: target},
	})
}
