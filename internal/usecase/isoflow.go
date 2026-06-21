// Package isoflow serializes the shared scene for isometric / 2.5D integrations.
package usecase

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/fs"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
)

const (
	documentVersion     = "3.3.0"
	tileSize            = 100.0
	defaultLabelHeight  = 60.0
	iconLabelToken      = "__XALIGO_LABEL__"
	maxEmbeddedLabelLen = 12
)

type colorRegistry struct {
	ids    map[string]string
	colors []entity.IsoflowColor
}

// Render converts the shared Excalidraw scene into an initial Isoflow document.
func RenderIsoflowScene(sceneJSON []byte) ([]byte, error) {
	return RenderIsoflowWithIcons(sceneJSON, nil)
}

// RenderWithIcons converts the shared Excalidraw scene and applies optional
// Isoflow-specific icon data URLs keyed by Excalidraw file ID.
func RenderIsoflowWithIcons(sceneJSON []byte, iconOverrides map[string]string) ([]byte, error) {
	var scene entity.PptxScene
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		return nil, fmt.Errorf("decode resolved scene for Isoflow: %w", err)
	}
	labels := itemLabels(scene.Elements)
	colors := colorRegistry{ids: map[string]string{}, colors: []entity.IsoflowColor{}}
	rectangles := groupRectangles(scene.Elements, &colors)

	itemIDs := map[string]bool{}
	modelItems := []entity.IsoflowModelItem{}
	viewItems := []entity.IsoflowViewItem{}
	iconsByID := map[string]entity.IsoflowIcon{}
	placer := newTilePlacer()
	for _, element := range scene.Elements {
		if element.IsDeleted || element.Type != "image" || !strings.HasSuffix(element.ID, "-item") {
			continue
		}
		label := labels[element.ID]
		if label == "" {
			label = strings.TrimSuffix(element.ID, "-item")
		}
		item := entity.IsoflowModelItem{ID: element.ID, Name: label}
		if file, ok := scene.Files[element.FileID]; ok && file.DataURL != "" {
			item.IsoflowIcon = element.FileID
			iconURL := file.DataURL
			if override := iconOverrides[element.FileID]; override != "" {
				iconURL = embedIconLabel(override, label)
			}
			iconsByID[element.FileID] = entity.IsoflowIcon{ID: element.FileID, Name: label, URL: iconURL, Collection: "xaligo", IsIsometric: true}
		}
		modelItems = append(modelItems, item)
		viewItems = append(viewItems, entity.IsoflowViewItem{ID: element.ID, Tile: placer.place(pixelToTile(element.X+element.Width/2, element.Y+element.Height/2)), LabelHeight: defaultLabelHeight})
		itemIDs[element.ID] = true
	}

	connectors := []entity.IsoflowConnector{}
	for _, element := range scene.Elements {
		if element.IsDeleted || (element.Type != "arrow" && element.Type != "line") || element.StartBinding == nil || element.EndBinding == nil {
			continue
		}
		source := ItemNodeID(element.StartBinding.ElementID)
		target := ItemNodeID(element.EndBinding.ElementID)
		if !itemIDs[source] || !itemIDs[target] {
			continue
		}
		connectors = append(connectors, entity.IsoflowConnector{
			ID:           element.ID,
			IsoflowColor: colors.idFor(element.StrokeColor),
			Width:        PositiveWidth(element.StrokeWidth),
			Style:        isoflowConnectorStyle(element.StrokeStyle),
			Anchors: []entity.IsoflowConnectorAnchor{
				{ID: element.ID + "-source", Ref: entity.IsoflowAnchorRef{Item: source}},
				{ID: element.ID + "-target", Ref: entity.IsoflowAnchorRef{Item: target}},
			},
		})
	}

	icons := make([]entity.IsoflowIcon, 0, len(iconsByID))
	for _, icon := range iconsByID {
		icons = append(icons, icon)
	}
	sort.Slice(icons, func(i, j int) bool { return icons[i].ID < icons[j].ID })
	document := entity.IsoflowDocument{
		Version: documentVersion,
		Title:   "xaligo export",
		Items:   modelItems,
		Views: []entity.IsoflowView{
			{ID: "main", Name: "Main", Items: viewItems, Rectangles: rectangles, Connectors: connectors},
		},
		Icons:     icons,
		Colors:    colors.colors,
		FitToView: true,
	}
	return json.MarshalIndent(document, "", "  ")
}

// groupRectangles preserves xaligo's visible container boundaries in Isoflow.
// The GroupBorder marker is intentionally required: ordinary Excalidraw
// rectangles may be decorations and must not become Isoflow floor regions.
func groupRectangles(elements []entity.Element, colors *colorRegistry) []entity.IsoflowRectangle {
	rectangles := []entity.IsoflowRectangle{}
	for _, element := range elements {
		if element.IsDeleted || element.Type != "rectangle" || element.Width <= 0 || element.Height <= 0 ||
			element.CustomData == nil || !element.CustomData.GroupBorder || element.StrokeColor == "transparent" {
			continue
		}
		from := pixelToTile(element.X, element.Y)
		to := pixelToTile(element.X+element.Width, element.Y+element.Height)
		if from.X == to.X || from.Y == to.Y {
			continue
		}
		if from.X > to.X {
			from.X, to.X = to.X, from.X
		}
		if from.Y > to.Y {
			from.Y, to.Y = to.Y, from.Y
		}
		rectangles = append(rectangles, entity.IsoflowRectangle{
			ID: element.ID, IsoflowColor: colors.idFor(element.StrokeColor), From: from, To: to,
		})
	}
	return rectangles
}

func LoadIsoflowIconManifest(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return decodeIconManifest(data)
}

func LoadIsoflowIconManifestFS(fsys fs.FS, path string) (map[string]string, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, err
	}
	return decodeIconManifest(data)
}

func decodeIconManifest(data []byte) (map[string]string, error) {
	var manifest entity.IsoflowIconManifest
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
	return svgDataURLFromBytes([]byte(text))
}

func escapeXML(value string) string {
	value = strings.ReplaceAll(value, "&", "&amp;")
	value = strings.ReplaceAll(value, "\"", "&quot;")
	value = strings.ReplaceAll(value, "<", "&lt;")
	value = strings.ReplaceAll(value, ">", "&gt;")
	return value
}

func itemLabels(elements []entity.Element) map[string]string {
	labels := map[string]string{}
	for _, element := range elements {
		if element.Type == "text" && strings.HasSuffix(element.ID, "-item-lbl") {
			labels[strings.TrimSuffix(element.ID, "-lbl")] = element.Text
		}
	}
	return labels
}

func isoflowNormalizedColor(color, fallback string) string {
	if strings.HasPrefix(color, "#") && len(color) == 7 {
		return color
	}
	return fallback
}

func isoflowConnectorStyle(style string) string {
	switch style {
	case "dashed":
		return "DASHED"
	case "dotted":
		return "DOTTED"
	default:
		return "SOLID"
	}
}

func pixelToTile(x, y float64) entity.IsoflowCoords {
	return entity.IsoflowCoords{X: math.Round(x/tileSize) * 2, Y: math.Round(y/tileSize) * 2}
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

func (placer *tilePlacer) place(preferred entity.IsoflowCoords) entity.IsoflowCoords {
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
					return entity.IsoflowCoords{X: float64(candidate.x), Y: float64(candidate.y)}
				}
			}
		}
	}
	placer.reserve(base)
	return entity.IsoflowCoords{X: float64(base.x), Y: float64(base.y)}
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
	value := isoflowNormalizedColor(color, "#999999")
	if id, ok := registry.ids[value]; ok {
		return id
	}
	id := fmt.Sprintf("color%d", len(registry.colors)+1)
	registry.ids[value] = id
	registry.colors = append(registry.colors, entity.IsoflowColor{ID: id, Value: value})
	return id
}
