// Package repository contains output adapters for the shared resolved scene.
package repository

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

type XYFlowRepository interface {
	Render(sceneJSON []byte) ([]byte, error)
}

type xyFlowRepository struct{}

func NewXYFlowRepository() XYFlowRepository { return &xyFlowRepository{} }

type xyFlowGroup struct {
	element entity.Element
	parent  string
}

type xyFlowLogicalEdge struct {
	element       entity.Element
	sourceElement string
	targetElement string
	sourceBinding *entity.Binding
	targetBinding *entity.Binding
}

func (rcvr *xyFlowRepository) Render(sceneJSON []byte) ([]byte, error) {
	var scene entity.PresentationScene
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		return nil, fmt.Errorf("decode resolved scene for XYFlow: %w", err)
	}
	labels := xyFlowNodeLabels(scene.Elements)
	groupIcons := map[string]string{}
	for _, element := range scene.Elements {
		if element.Type == "image" && strings.HasSuffix(element.ID, "-icon") {
			if file, ok := scene.Files[element.FileID]; ok && file.DataURL != "" {
				groupIcons[strings.TrimSuffix(element.ID, "-icon")+"-rect"] = file.DataURL
			}
		}
	}

	groups := collectGroups(scene.Elements)
	nodes := make([]entity.XYFlowNode, 0, len(groups))
	elementNodes := map[string]string{}
	for _, candidate := range groups {
		element := candidate.element
		position := entity.XYFlowPosition{X: element.X, Y: element.Y}
		if candidate.parent != "" {
			if parent, ok := groupByID(groups, candidate.parent); ok {
				position.X -= parent.element.X
				position.Y -= parent.element.Y
			}
		}
		data := map[string]any{"kind": "xyFlowGroup"}
		if label := labels[element.ID]; label != "" {
			data["label"] = label
		}
		if semanticKind := xyFlowSemanticElementKind(element); semanticKind != "" {
			data["semanticKind"] = semanticKind
		}
		addXYFlowUMLNodeData(data, element)
		if icon := groupIcons[element.ID]; icon != "" {
			data["icon"] = icon
		}
		node := entity.XYFlowNode{
			ID: element.ID, Type: "xyFlowGroup", XYFlowPosition: position,
			Width: element.Width, Height: element.Height, ParentID: candidate.parent,
			Data: data,
			Style: map[string]any{
				"borderColor": element.StrokeColor, "backgroundColor": element.BackgroundColor,
				"borderWidth": element.StrokeWidth, "borderStyle": cssDash(element.StrokeStyle),
			},
		}
		if candidate.parent != "" && xyFlowElementInsideParent(groups, candidate.parent, element) {
			node.Extent = "parent"
		}
		nodes = append(nodes, node)
		registerXYFlowElementNode(elementNodes, element.ID, node.ID)
	}

	for _, element := range scene.Elements {
		semanticKind := xyFlowSemanticElementKind(element)
		legacyItem := semanticKind == "" && strings.HasSuffix(element.ID, "-item")
		if element.IsDeleted || element.Type != "image" || (semanticKind != "item" && !legacyItem) {
			continue
		}
		parent := ""
		if semanticKind == "item" {
			parent = xyFlowSemanticParent(groups, element)
		} else {
			parent = smallestContainingGroup(groups, element)
		}
		position := entity.XYFlowPosition{X: element.X, Y: element.Y}
		if parent != "" {
			if candidate, ok := groupByID(groups, parent); ok {
				position.X -= candidate.element.X
				position.Y -= candidate.element.Y
			}
		}
		data := map[string]any{"kind": "item", "label": labels[element.ID], "fileId": element.FileID}
		if semanticKind != "" {
			data["semanticKind"] = semanticKind
		}
		if file, ok := scene.Files[element.FileID]; ok && file.DataURL != "" {
			data["image"] = file.DataURL
		}
		node := entity.XYFlowNode{ID: element.ID, Type: "xaligoItem", XYFlowPosition: position, Width: element.Width, Height: element.Height, ParentID: parent, Data: data}
		if parent != "" && xyFlowElementInsideParent(groups, parent, element) {
			node.Extent = "parent"
		}
		nodes = append(nodes, node)
		registerXYFlowElementNode(elementNodes, element.ID, node.ID)
		registerXYFlowElementNode(elementNodes, element.ID+"-lbl", node.ID)
	}

	edges := []entity.XYFlowEdge{}
	logicalEdges := map[string]*xyFlowLogicalEdge{}
	logicalEdgeOrder := []string{}
	for _, element := range scene.Elements {
		if element.IsDeleted || (element.Type != "arrow" && element.Type != "line") {
			continue
		}
		if element.CustomData != nil && element.CustomData.ConnectorCrossFrame {
			logicalID := strings.TrimSpace(element.CustomData.ConnectorLogicalID)
			if logicalID == "" {
				continue
			}
			logical := logicalEdges[logicalID]
			if logical == nil {
				logical = &xyFlowLogicalEdge{
					element:       element,
					sourceElement: element.CustomData.ConnectorSourceElementID,
					targetElement: element.CustomData.ConnectorDestinationElementID,
				}
				logicalEdges[logicalID] = logical
				logicalEdgeOrder = append(logicalEdgeOrder, logicalID)
			}
			collectXYFlowLogicalBinding(logical, element.StartBinding)
			collectXYFlowLogicalBinding(logical, element.EndBinding)
			continue
		}
		if element.StartBinding == nil || element.EndBinding == nil {
			continue
		}
		source := xyFlowNodeForElement(elementNodes, element.StartBinding.ElementID)
		target := xyFlowNodeForElement(elementNodes, element.EndBinding.ElementID)
		if source == "" || target == "" {
			continue
		}
		edges = append(edges, buildXYFlowEdge(element, element.ID, source, target, element.StartBinding, element.EndBinding))
	}
	for _, logicalID := range logicalEdgeOrder {
		logical := logicalEdges[logicalID]
		source := xyFlowNodeForElement(elementNodes, logical.sourceElement)
		target := xyFlowNodeForElement(elementNodes, logical.targetElement)
		if source == "" || target == "" {
			continue
		}
		edges = append(edges, buildXYFlowEdge(logical.element, logicalID, source, target, logical.sourceBinding, logical.targetBinding))
	}

	width, height := sceneSize(scene.Elements)
	background := "#ffffff"
	if scene.AppState != nil && scene.AppState.ViewBackgroundColor != "" {
		background = scene.AppState.ViewBackgroundColor
	}
	document := entity.XYFlowDocument{Nodes: nodes, Edges: edges, XYFlowViewport: entity.XYFlowViewport{Zoom: 1}, Width: width, Height: height, Background: background}
	return json.MarshalIndent(document, "", "  ")
}

func collectGroups(elements []entity.Element) []xyFlowGroup {
	groups := []xyFlowGroup{}
	for _, element := range elements {
		if element.IsDeleted || element.ID == "paper-frame" || strings.HasSuffix(element.ID, "-header-bg") || element.Width <= 0 || element.Height <= 0 {
			continue
		}
		semanticKind := xyFlowSemanticElementKind(element)
		if semanticKind != "" {
			if !xyFlowSemanticGroupKind(semanticKind) {
				continue
			}
			groups = append(groups, xyFlowGroup{element: element, parent: strings.TrimSpace(element.CustomData.SemanticParentElementID)})
			continue
		}
		if !xyFlowGenericNodeElement(element) || (element.CustomData != nil && element.CustomData.AnchorBackground) {
			continue
		}
		groups = append(groups, xyFlowGroup{element: element})
	}
	for i := range groups {
		if xyFlowSemanticElementKind(groups[i].element) == "" {
			groups[i].parent = smallestContainingGroupExcluding(groups, groups[i].element, groups[i].element.ID)
			continue
		}
		if groups[i].parent != "" {
			if _, ok := groupByID(groups, groups[i].parent); !ok {
				groups[i].parent = ""
			}
		}
	}
	return groups
}

func xyFlowGenericNodeElement(element entity.Element) bool {
	switch element.Type {
	case "rectangle", "frame":
		return true
	case "ellipse", "diamond":
		return element.CustomData != nil && strings.TrimSpace(element.CustomData.UMLElementKind) != ""
	default:
		return false
	}
}

func xyFlowNodeLabels(elements []entity.Element) map[string]string {
	labels := map[string]string{}
	for _, element := range elements {
		if element.IsDeleted || element.Type != "text" || strings.TrimSpace(element.Text) == "" {
			continue
		}
		label := strings.TrimSpace(element.Text)
		switch {
		case strings.HasSuffix(element.ID, "-item-lbl"):
			labels[strings.TrimSuffix(element.ID, "-lbl")] = label
		case strings.HasSuffix(element.ID, "-label"):
			labels[strings.TrimSuffix(element.ID, "-label")+"-rect"] = label
		case strings.HasSuffix(element.ID, "-text"):
			labels[strings.TrimSuffix(element.ID, "-text")+"-rect"] = label
		}
	}
	return labels
}

func addXYFlowUMLNodeData(data map[string]any, element entity.Element) {
	if element.CustomData == nil || strings.TrimSpace(element.CustomData.UMLElementKind) == "" {
		return
	}
	data["shape"] = element.Type
	for key, value := range map[string]string{
		"umlId":               element.CustomData.UMLID,
		"umlLocalId":          element.CustomData.UMLLocalID,
		"umlReference":        element.CustomData.UMLReference,
		"umlDiagramKind":      element.CustomData.UMLDiagramKind,
		"umlElementKind":      element.CustomData.UMLElementKind,
		"umlOwnerId":          element.CustomData.UMLOwnerID,
		"umlOwnerReference":   element.CustomData.UMLOwnerReference,
		"umlCompartmentKinds": element.CustomData.UMLCompartmentKinds,
		"umlTimeFrom":         element.CustomData.UMLTimeFrom,
		"umlTimeTo":           element.CustomData.UMLTimeTo,
	} {
		if value = strings.TrimSpace(value); value != "" {
			data[key] = value
		}
	}
}

func xyFlowSemanticElementKind(element entity.Element) string {
	if element.CustomData == nil {
		return ""
	}
	return strings.TrimSpace(element.CustomData.SemanticElementKind)
}

func xyFlowSemanticGroupKind(kind string) bool {
	switch kind {
	case "frame", "group", "rectangle", "port":
		return true
	default:
		return false
	}
}

func xyFlowSemanticParent(groups []xyFlowGroup, element entity.Element) string {
	if element.CustomData == nil {
		return ""
	}
	parent := strings.TrimSpace(element.CustomData.SemanticParentElementID)
	if parent == "" {
		return ""
	}
	if _, ok := groupByID(groups, parent); !ok {
		return ""
	}
	return parent
}

func xyFlowElementInsideParent(groups []xyFlowGroup, parentID string, element entity.Element) bool {
	parent, ok := groupByID(groups, parentID)
	return ok && contains(parent.element, element)
}

func registerXYFlowElementNode(elementNodes map[string]string, elementID, nodeID string) {
	if elementID == "" || nodeID == "" {
		return
	}
	if _, exists := elementNodes[elementID]; !exists {
		elementNodes[elementID] = nodeID
	}
}

func xyFlowNodeForElement(elementNodes map[string]string, elementID string) string {
	if nodeID := elementNodes[elementID]; nodeID != "" {
		return nodeID
	}
	return elementNodes[share.ItemNodeID(elementID)]
}

func collectXYFlowLogicalBinding(logical *xyFlowLogicalEdge, binding *entity.Binding) {
	if logical == nil || binding == nil {
		return
	}
	switch binding.ElementID {
	case logical.sourceElement:
		logical.sourceBinding = binding
	case logical.targetElement:
		logical.targetBinding = binding
	}
}

func buildXYFlowEdge(element entity.Element, id, source, target string, sourceBinding, targetBinding *entity.Binding) entity.XYFlowEdge {
	kind, startHead, endHead := connectorData(element)
	color := xyFlowNormalizedColor(element.StrokeColor, "#1e1e1e")
	data := map[string]any{"kind": kind, "startArrowhead": startHead, "endArrowhead": endHead}
	if element.CustomData != nil {
		if element.CustomData.ConnectorBends != "" {
			data["bends"] = element.CustomData.ConnectorBends
		}
		if element.CustomData.ConnectorScale > 0 {
			data["scale"] = element.CustomData.ConnectorScale
		}
		if element.CustomData.ConnectorGrid > 0 {
			data["grid"] = element.CustomData.ConnectorGrid
		}
		if element.CustomData.ConnectorSrcAnchor {
			data["sourceAnchorExplicit"] = true
		}
		if element.CustomData.ConnectorDstAnchor {
			data["targetAnchorExplicit"] = true
		}
		if element.CustomData.ConnectorCrossFrame {
			data["crossFrame"] = true
			data["sourceFrame"] = element.CustomData.ConnectorSourceFrame
			data["targetFrame"] = element.CustomData.ConnectorDestinationFrame
		}
		addXYFlowUMLEdgeData(data, element.CustomData)
	}
	if sourceBinding != nil && len(sourceBinding.FixedPoint) >= 2 {
		data["sourceFixedPoint"] = append([]float64(nil), sourceBinding.FixedPoint...)
	}
	if targetBinding != nil && len(targetBinding.FixedPoint) >= 2 {
		data["targetFixedPoint"] = append([]float64(nil), targetBinding.FixedPoint...)
	}
	return entity.XYFlowEdge{
		ID: id, Source: source, Target: target,
		SourceHandle: bindingHandle(sourceBinding), TargetHandle: bindingHandle(targetBinding),
		Type: "smoothstep", ZIndex: edgeZIndex(kind),
		Data:        data,
		Style:       map[string]any{"stroke": color, "strokeWidth": share.PositiveWidth(element.StrokeWidth), "strokeDasharray": cssStrokeDash(element.StrokeStyle)},
		MarkerStart: marker(startHead, color), MarkerEnd: marker(endHead, color),
	}
}

func addXYFlowUMLEdgeData(data map[string]any, customData *entity.CustomData) {
	if customData == nil || strings.TrimSpace(customData.UMLRelationKind) == "" {
		return
	}
	for key, value := range map[string]string{
		"umlRelationKind":            customData.UMLRelationKind,
		"umlRelationLabel":           customData.UMLRelationLabel,
		"umlSourceReference":         customData.UMLRelationSourceReference,
		"umlDestinationReference":    customData.UMLRelationDestinationReference,
		"umlMessageOrder":            customData.UMLMessageOrder,
		"umlGuard":                   customData.UMLGuard,
		"umlSourceMultiplicity":      customData.UMLSourceMultiplicity,
		"umlDestinationMultiplicity": customData.UMLDestinationMultiplicity,
		"umlOccurrenceAt":            customData.UMLOccurrenceAt,
		"umlDurationFrom":            customData.UMLDurationFrom,
		"umlDurationTo":              customData.UMLDurationTo,
	} {
		if value = strings.TrimSpace(value); value != "" {
			data[key] = value
		}
	}
}

func bindingHandle(binding *entity.Binding) string {
	if binding == nil {
		return ""
	}
	return bindingSide(binding.FixedPoint)
}

func smallestContainingGroup(groups []xyFlowGroup, element entity.Element) string {
	return smallestContainingGroupExcluding(groups, element, "")
}

func smallestContainingGroupExcluding(groups []xyFlowGroup, element entity.Element, exclude string) string {
	best, bestArea := "", math.Inf(1)
	for _, candidate := range groups {
		container := candidate.element
		if container.ID == exclude || !contains(container, element) {
			continue
		}
		area := container.Width * container.Height
		if area < bestArea && area > element.Width*element.Height {
			best, bestArea = container.ID, area
		}
	}
	return best
}

func contains(container, element entity.Element) bool {
	return element.X >= container.X && element.Y >= container.Y && element.X+element.Width <= container.X+container.Width && element.Y+element.Height <= container.Y+container.Height
}

func groupByID(groups []xyFlowGroup, id string) (xyFlowGroup, bool) {
	for _, candidate := range groups {
		if candidate.element.ID == id {
			return candidate, true
		}
	}
	return xyFlowGroup{}, false
}

func connectorData(element entity.Element) (kind, start, end string) {
	kind = "connection"
	if element.CustomData != nil {
		if element.CustomData.ConnectorKind != "" {
			kind = element.CustomData.ConnectorKind
		}
		start = element.CustomData.ConnectorStartArrowhead
		end = element.CustomData.ConnectorEndArrowhead
	}
	return
}

func bindingSide(point []float64) string {
	if len(point) < 2 {
		return ""
	}
	if point[1] < 0.01 {
		return "top"
	}
	if point[1] > 0.99 {
		return "bottom"
	}
	if point[0] < 0.01 {
		return "left"
	}
	if point[0] > 0.99 {
		return "right"
	}
	return ""
}

func marker(arrowhead, color string) *entity.XYFlowMarker {
	switch arrowhead {
	case "arrow", "triangle", "stealth":
		return &entity.XYFlowMarker{Type: "arrowclosed", Color: color}
	default:
		return nil
	}
}

func edgeZIndex(kind string) int {
	if kind == "traffic" {
		return 2
	}
	if kind == "route" {
		return 0
	}
	return 1
}

func cssDash(style string) string {
	if style == "dashed" || style == "dotted" {
		return style
	}
	return "solid"
}

func cssStrokeDash(style string) string {
	if style == "dashed" {
		return "8 6"
	}
	if style == "dotted" {
		return "2 5"
	}
	return ""
}

func xyFlowNormalizedColor(value, fallback string) string {
	if value == "" || value == "transparent" {
		return fallback
	}
	return value
}

func sceneSize(elements []entity.Element) (float64, float64) {
	for _, element := range elements {
		if element.ID == "paper-frame" {
			return element.Width, element.Height
		}
	}
	return 1280, 720
}
