// Package repository contains output adapters for the shared resolved scene.
package repository

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
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

func (rcvr *xyFlowRepository) Render(sceneJSON []byte) ([]byte, error) {
	var scene entity.PptxScene
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		return nil, fmt.Errorf("decode resolved scene for XYFlow: %w", err)
	}
	labels := map[string]string{}
	groupIcons := map[string]string{}
	for _, element := range scene.Elements {
		if element.Type == "text" && strings.HasSuffix(element.ID, "-item-lbl") {
			labels[strings.TrimSuffix(element.ID, "-lbl")] = element.Text
		}
		if element.Type == "image" && strings.HasSuffix(element.ID, "-icon") {
			if file, ok := scene.Files[element.FileID]; ok && file.DataURL != "" {
				groupIcons[strings.TrimSuffix(element.ID, "-icon")+"-rect"] = file.DataURL
			}
		}
	}

	groups := collectGroups(scene.Elements)
	nodes := make([]entity.XYFlowNode, 0, len(groups))
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
		if candidate.parent != "" {
			node.Extent = "parent"
		}
		nodes = append(nodes, node)
	}

	itemIDs := map[string]bool{}
	for _, element := range scene.Elements {
		if element.IsDeleted || element.Type != "image" || !strings.HasSuffix(element.ID, "-item") {
			continue
		}
		parent := smallestContainingGroup(groups, element)
		position := entity.XYFlowPosition{X: element.X, Y: element.Y}
		if parent != "" {
			if candidate, ok := groupByID(groups, parent); ok {
				position.X -= candidate.element.X
				position.Y -= candidate.element.Y
			}
		}
		data := map[string]any{"kind": "item", "label": labels[element.ID], "fileId": element.FileID}
		if file, ok := scene.Files[element.FileID]; ok && file.DataURL != "" {
			data["image"] = file.DataURL
		}
		node := entity.XYFlowNode{ID: element.ID, Type: "xaligoItem", XYFlowPosition: position, Width: element.Width, Height: element.Height, ParentID: parent, Data: data}
		if parent != "" {
			node.Extent = "parent"
		}
		nodes = append(nodes, node)
		itemIDs[element.ID] = true
	}

	edges := []entity.XYFlowEdge{}
	for _, element := range scene.Elements {
		if element.IsDeleted || (element.Type != "arrow" && element.Type != "line") || element.StartBinding == nil || element.EndBinding == nil {
			continue
		}
		source := share.ItemNodeID(element.StartBinding.ElementID)
		target := share.ItemNodeID(element.EndBinding.ElementID)
		if !itemIDs[source] || !itemIDs[target] {
			continue
		}
		kind, startHead, endHead := connectorData(element)
		color := xyFlowNormalizedColor(element.StrokeColor, "#1e1e1e")
		edge := entity.XYFlowEdge{
			ID: element.ID, Source: source, Target: target,
			SourceHandle: bindingSide(element.StartBinding.FixedPoint), TargetHandle: bindingSide(element.EndBinding.FixedPoint),
			Type: "smoothstep", ZIndex: edgeZIndex(kind),
			Data:        map[string]any{"kind": kind, "startArrowhead": startHead, "endArrowhead": endHead},
			Style:       map[string]any{"stroke": color, "strokeWidth": share.PositiveWidth(element.StrokeWidth), "strokeDasharray": cssStrokeDash(element.StrokeStyle)},
			MarkerStart: marker(startHead, color), MarkerEnd: marker(endHead, color),
		}
		edges = append(edges, edge)
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
		if element.IsDeleted || element.ID == "paper-frame" || strings.HasSuffix(element.ID, "-header-bg") || element.Type != "rectangle" || element.Width <= 0 || element.Height <= 0 {
			continue
		}
		groups = append(groups, xyFlowGroup{element: element})
	}
	for i := range groups {
		groups[i].parent = smallestContainingGroupExcluding(groups, groups[i].element, groups[i].element.ID)
	}
	return groups
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
