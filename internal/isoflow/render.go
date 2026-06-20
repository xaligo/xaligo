// Package isoflow serializes the shared scene for isometric / 2.5D integrations.
package isoflow

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/fs"
	"math"
	"os"
	"strings"

	"github.com/ryo-arima/xaligo/internal/pptxplan"
)

const (
	documentVersion     = "3.3.0"
	tileSize            = 100.0
	defaultLabelHeight  = 60.0
	iconLabelToken      = "__XALIGO_LABEL__"
	svgDataURLPrefix    = "data:image/svg+xml;base64,"
	maxEmbeddedLabelLen = 12
)

type Document struct {
	Version     string      `json:"version"`
	Title       string      `json:"title"`
	Description string      `json:"description,omitempty"`
	Items       []ModelItem `json:"items"`
	Views       []View      `json:"views"`
	Icons       []Icon      `json:"icons"`
	Colors      []Color     `json:"colors"`
	FitToView   bool        `json:"fitToView,omitempty"`
}

type ModelItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
}

type View struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Items       []ViewItem  `json:"items"`
	Rectangles  []Rectangle `json:"rectangles,omitempty"`
	Connectors  []Connector `json:"connectors,omitempty"`
}

type ViewItem struct {
	ID          string  `json:"id"`
	Tile        Coords  `json:"tile"`
	LabelHeight float64 `json:"labelHeight"`
}

type Rectangle struct {
	ID    string `json:"id"`
	Color string `json:"color,omitempty"`
	From  Coords `json:"from"`
	To    Coords `json:"to"`
}

type Connector struct {
	ID      string            `json:"id"`
	Color   string            `json:"color,omitempty"`
	Width   float64           `json:"width,omitempty"`
	Style   string            `json:"style,omitempty"`
	Anchors []ConnectorAnchor `json:"anchors"`
}

type ConnectorAnchor struct {
	ID  string    `json:"id"`
	Ref AnchorRef `json:"ref"`
}

type AnchorRef struct {
	Item string `json:"item,omitempty"`
}

type Icon struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Collection  string `json:"collection,omitempty"`
	IsIsometric bool   `json:"isIsometric,omitempty"`
}

type Color struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Coords struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type colorRegistry struct {
	ids    map[string]string
	colors []Color
}

type IconManifest struct {
	Icons map[string]IconManifestEntry `json:"icons"`
}

type IconManifestEntry struct {
	DataURL string `json:"dataURL"`
}

// Render converts the shared Excalidraw scene into an initial Isoflow document.
func Render(sceneJSON []byte) ([]byte, error) {
	return RenderWithIcons(sceneJSON, nil)
}

// RenderWithIcons converts the shared Excalidraw scene and applies optional
// Isoflow-specific icon data URLs keyed by Excalidraw file ID.
func RenderWithIcons(sceneJSON []byte, iconOverrides map[string]string) ([]byte, error) {
	var scene pptxplan.Scene
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		return nil, fmt.Errorf("decode resolved scene for Isoflow: %w", err)
	}
	labels := itemLabels(scene.Elements)
	colors := colorRegistry{ids: map[string]string{}, colors: []Color{}}
	rectangles := []Rectangle{}

	itemIDs := map[string]bool{}
	modelItems := []ModelItem{}
	viewItems := []ViewItem{}
	iconsByID := map[string]Icon{}
	placer := newTilePlacer()
	for _, element := range scene.Elements {
		if element.IsDeleted || element.Type != "image" || !strings.HasSuffix(element.ID, "-item") {
			continue
		}
		label := labels[element.ID]
		if label == "" {
			label = strings.TrimSuffix(element.ID, "-item")
		}
		item := ModelItem{ID: element.ID, Name: label}
		if file, ok := scene.Files[element.FileID]; ok && file.DataURL != "" {
			item.Icon = element.FileID
			iconURL := file.DataURL
			if override := iconOverrides[element.FileID]; override != "" {
				iconURL = embedIconLabel(override, label)
			}
			iconsByID[element.FileID] = Icon{ID: element.FileID, Name: label, URL: iconURL, Collection: "xaligo", IsIsometric: true}
		}
		modelItems = append(modelItems, item)
		viewItems = append(viewItems, ViewItem{ID: element.ID, Tile: placer.place(pixelToTile(element.X+element.Width/2, element.Y+element.Height/2)), LabelHeight: defaultLabelHeight})
		itemIDs[element.ID] = true
	}

	connectors := []Connector{}
	for _, element := range scene.Elements {
		if element.IsDeleted || (element.Type != "arrow" && element.Type != "line") || element.StartBinding == nil || element.EndBinding == nil {
			continue
		}
		source := itemNodeID(element.StartBinding.ElementID)
		target := itemNodeID(element.EndBinding.ElementID)
		if !itemIDs[source] || !itemIDs[target] {
			continue
		}
		connectors = append(connectors, Connector{
			ID:    element.ID,
			Color: colors.idFor(element.StrokeColor),
			Width: width(element.StrokeWidth),
			Style: connectorStyle(element.StrokeStyle),
			Anchors: []ConnectorAnchor{
				{ID: element.ID + "-source", Ref: AnchorRef{Item: source}},
				{ID: element.ID + "-target", Ref: AnchorRef{Item: target}},
			},
		})
	}

	icons := make([]Icon, 0, len(iconsByID))
	for _, icon := range iconsByID {
		icons = append(icons, icon)
	}
	document := Document{
		Version: documentVersion,
		Title:   "xaligo export",
		Items:   modelItems,
		Views: []View{
			{ID: "main", Name: "Main", Items: viewItems, Rectangles: rectangles, Connectors: connectors},
		},
		Icons:     icons,
		Colors:    colors.colors,
		FitToView: true,
	}
	return json.MarshalIndent(document, "", "  ")
}

func LoadIconManifest(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return decodeIconManifest(data)
}

func LoadIconManifestFS(fsys fs.FS, path string) (map[string]string, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, err
	}
	return decodeIconManifest(data)
}

func decodeIconManifest(data []byte) (map[string]string, error) {
	var manifest IconManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("decode Isoflow icon manifest: %w", err)
	}
	overrides := map[string]string{}
	for id, entry := range manifest.Icons {
		if entry.DataURL != "" {
			overrides[id] = entry.DataURL
		}
	}
	return overrides, nil
}

func embedIconLabel(dataURL, label string) string {
	if !strings.HasPrefix(dataURL, svgDataURLPrefix) {
		return dataURL
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(dataURL, svgDataURLPrefix))
	if err != nil {
		return dataURL
	}
	text := string(raw)
	if !strings.Contains(text, iconLabelToken) {
		return dataURL
	}
	label = strings.TrimSpace(label)
	if len([]rune(label)) > maxEmbeddedLabelLen {
		runes := []rune(label)
		label = string(runes[:maxEmbeddedLabelLen-1]) + "…"
	}
	text = strings.ReplaceAll(text, iconLabelToken, escapeXML(label))
	return svgDataURLPrefix + base64.StdEncoding.EncodeToString([]byte(text))
}

func escapeXML(value string) string {
	value = strings.ReplaceAll(value, "&", "&amp;")
	value = strings.ReplaceAll(value, "\"", "&quot;")
	value = strings.ReplaceAll(value, "<", "&lt;")
	value = strings.ReplaceAll(value, ">", "&gt;")
	return value
}

func itemLabels(elements []pptxplan.Element) map[string]string {
	labels := map[string]string{}
	for _, element := range elements {
		if element.Type == "text" && strings.HasSuffix(element.ID, "-item-lbl") {
			labels[strings.TrimSuffix(element.ID, "-lbl")] = element.Text
		}
	}
	return labels
}

func itemNodeID(bindingID string) string {
	return strings.TrimSuffix(bindingID, "-lbl")
}

func normalizedColor(color, fallback string) string {
	if strings.HasPrefix(color, "#") && len(color) == 7 {
		return color
	}
	return fallback
}

func width(value float64) float64 {
	if value > 0 {
		return value
	}
	return 1
}

func connectorStyle(style string) string {
	switch style {
	case "dashed":
		return "DASHED"
	case "dotted":
		return "DOTTED"
	default:
		return "SOLID"
	}
}

func pixelToTile(x, y float64) Coords {
	return Coords{X: math.Round(x/tileSize) * 2, Y: math.Round(y/tileSize) * 2}
}

type tilePlacer struct {
	occupied map[tileKey]bool
}

type tileKey struct {
	x int
	y int
}

func newTilePlacer() *tilePlacer {
	return &tilePlacer{occupied: map[tileKey]bool{}}
}

func (placer *tilePlacer) place(preferred Coords) Coords {
	base := tileKey{x: int(math.Round(preferred.X)), y: int(math.Round(preferred.Y))}
	for radius := 0; radius < 512; radius++ {
		for dy := -radius; dy <= radius; dy++ {
			for dx := -radius; dx <= radius; dx++ {
				if max(abs(dx), abs(dy)) != radius {
					continue
				}
				candidate := tileKey{x: base.x + dx, y: base.y + dy}
				if placer.available(candidate) {
					placer.reserve(candidate)
					return Coords{X: float64(candidate.x), Y: float64(candidate.y)}
				}
			}
		}
	}
	placer.reserve(base)
	return Coords{X: float64(base.x), Y: float64(base.y)}
}

func (placer *tilePlacer) available(key tileKey) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if placer.occupied[tileKey{x: key.x + dx, y: key.y + dy}] {
				return false
			}
		}
	}
	return true
}

func (placer *tilePlacer) reserve(key tileKey) {
	placer.occupied[key] = true
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (registry *colorRegistry) idFor(color string) string {
	value := normalizedColor(color, "#999999")
	if id, ok := registry.ids[value]; ok {
		return id
	}
	id := fmt.Sprintf("color%d", len(registry.colors)+1)
	registry.ids[value] = id
	registry.colors = append(registry.colors, Color{ID: id, Value: value})
	return id
}
