// Package xyflow serializes the shared resolved scene for React Flow / XYFlow.
package xyflow

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/ryo-arima/xaligo/internal/pptxplan"
)

type Document struct {
	Nodes      []Node   `json:"nodes"`
	Edges      []Edge   `json:"edges"`
	Viewport   Viewport `json:"viewport"`
	Width      float64  `json:"width"`
	Height     float64  `json:"height"`
	Background string   `json:"background"`
}

type Viewport struct {
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Zoom float64 `json:"zoom"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Node struct {
	ID       string         `json:"id"`
	Type     string         `json:"type"`
	Position Position       `json:"position"`
	Width    float64        `json:"width"`
	Height   float64        `json:"height"`
	ParentID string         `json:"parentId,omitempty"`
	Extent   string         `json:"extent,omitempty"`
	Data     map[string]any `json:"data"`
	Style    map[string]any `json:"style,omitempty"`
}

type Edge struct {
	ID           string         `json:"id"`
	Source       string         `json:"source"`
	Target       string         `json:"target"`
	SourceHandle string         `json:"sourceHandle,omitempty"`
	TargetHandle string         `json:"targetHandle,omitempty"`
	Type         string         `json:"type"`
	ZIndex       int            `json:"zIndex,omitempty"`
	Data         map[string]any `json:"data"`
	Style        map[string]any `json:"style,omitempty"`
	MarkerStart  *Marker        `json:"markerStart,omitempty"`
	MarkerEnd    *Marker        `json:"markerEnd,omitempty"`
}

type Marker struct {
	Type  string `json:"type"`
	Color string `json:"color,omitempty"`
}

type group struct {
	element pptxplan.Element
	parent  string
}

func Render(sceneJSON []byte) ([]byte, error) {
	var scene pptxplan.Scene
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
	nodes := make([]Node, 0, len(groups))
	for _, candidate := range groups {
		element := candidate.element
		position := Position{X: element.X, Y: element.Y}
		if candidate.parent != "" {
			if parent, ok := groupByID(groups, candidate.parent); ok {
				position.X -= parent.element.X
				position.Y -= parent.element.Y
			}
		}
		data := map[string]any{"kind": "group"}
		if icon := groupIcons[element.ID]; icon != "" {
			data["icon"] = icon
		}
		node := Node{
			ID: element.ID, Type: "group", Position: position,
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
		position := Position{X: element.X, Y: element.Y}
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
		node := Node{ID: element.ID, Type: "xaligoItem", Position: position, Width: element.Width, Height: element.Height, ParentID: parent, Data: data}
		if parent != "" {
			node.Extent = "parent"
		}
		nodes = append(nodes, node)
		itemIDs[element.ID] = true
	}

	edges := []Edge{}
	for _, element := range scene.Elements {
		if element.IsDeleted || (element.Type != "arrow" && element.Type != "line") || element.StartBinding == nil || element.EndBinding == nil {
			continue
		}
		source := itemNodeID(element.StartBinding.ElementID)
		target := itemNodeID(element.EndBinding.ElementID)
		if !itemIDs[source] || !itemIDs[target] {
			continue
		}
		kind, startHead, endHead := connectorData(element)
		color := normalizedColor(element.StrokeColor, "#1e1e1e")
		edge := Edge{
			ID: element.ID, Source: source, Target: target,
			SourceHandle: bindingSide(element.StartBinding.FixedPoint), TargetHandle: bindingSide(element.EndBinding.FixedPoint),
			Type: "smoothstep", ZIndex: edgeZIndex(kind),
			Data:        map[string]any{"kind": kind, "startArrowhead": startHead, "endArrowhead": endHead},
			Style:       map[string]any{"stroke": color, "strokeWidth": width(element.StrokeWidth), "strokeDasharray": cssStrokeDash(element.StrokeStyle)},
			MarkerStart: marker(startHead, color), MarkerEnd: marker(endHead, color),
		}
		edges = append(edges, edge)
	}

	width, height := sceneSize(scene.Elements)
	background := "#ffffff"
	if scene.AppState != nil && scene.AppState.ViewBackgroundColor != "" {
		background = scene.AppState.ViewBackgroundColor
	}
	document := Document{Nodes: nodes, Edges: edges, Viewport: Viewport{Zoom: 1}, Width: width, Height: height, Background: background}
	return json.MarshalIndent(document, "", "  ")
}

func collectGroups(elements []pptxplan.Element) []group {
	groups := []group{}
	for _, element := range elements {
		if element.IsDeleted || element.ID == "paper-frame" || strings.HasSuffix(element.ID, "-header-bg") || element.Type != "rectangle" || element.Width <= 0 || element.Height <= 0 {
			continue
		}
		groups = append(groups, group{element: element})
	}
	for i := range groups {
		groups[i].parent = smallestContainingGroupExcluding(groups, groups[i].element, groups[i].element.ID)
	}
	return groups
}

func smallestContainingGroup(groups []group, element pptxplan.Element) string {
	return smallestContainingGroupExcluding(groups, element, "")
}

func smallestContainingGroupExcluding(groups []group, element pptxplan.Element, exclude string) string {
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

func contains(container, element pptxplan.Element) bool {
	return element.X >= container.X && element.Y >= container.Y && element.X+element.Width <= container.X+container.Width && element.Y+element.Height <= container.Y+container.Height
}

func groupByID(groups []group, id string) (group, bool) {
	for _, candidate := range groups {
		if candidate.element.ID == id {
			return candidate, true
		}
	}
	return group{}, false
}

func itemNodeID(bindingID string) string {
	return strings.TrimSuffix(bindingID, "-lbl")
}

func connectorData(element pptxplan.Element) (kind, start, end string) {
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

func marker(arrowhead, color string) *Marker {
	switch arrowhead {
	case "arrow", "triangle", "stealth":
		return &Marker{Type: "arrowclosed", Color: color}
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

func width(value float64) float64 {
	if value <= 0 {
		return 1
	}
	return value
}

func normalizedColor(value, fallback string) string {
	if value == "" || value == "transparent" {
		return fallback
	}
	return value
}

func sceneSize(elements []pptxplan.Element) (float64, float64) {
	for _, element := range elements {
		if element.ID == "paper-frame" {
			return element.Width, element.Height
		}
	}
	return 1280, 720
}
