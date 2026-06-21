package usecase

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
)

const (
	ThemeLight = entity.ThemeLight
	ThemeDark  = entity.ThemeDark
)

// NormalizeTheme validates a renderer theme and supplies the default.
func NormalizeTheme(theme string) (string, error) {
	return entity.NormalizeTheme(theme)
}

// ApplyThemeJSON applies presentation colors to a rendered Excalidraw scene.
// Service colors and embedded icons are intentionally preserved; only the
// neutral canvas, text, borders and white surfaces are remapped.
func ApplyThemeJSON(sceneJSON []byte, theme string) ([]byte, error) {
	normalized, err := NormalizeTheme(theme)
	if err != nil {
		return nil, err
	}
	if normalized == ThemeLight {
		return sceneJSON, nil
	}

	var scene file
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		return nil, fmt.Errorf("decode excalidraw scene for theme: %w", err)
	}
	if scene.AppState == nil {
		scene.AppState = map[string]any{}
	}
	scene.AppState["viewBackgroundColor"] = "#111827"

	for _, element := range scene.Elements {
		if color, ok := element["strokeColor"].(string); ok {
			element["strokeColor"] = darkStrokeColor(color)
		}
		if color, ok := element["backgroundColor"].(string); ok {
			element["backgroundColor"] = darkFillColor(color)
		}
	}
	return json.MarshalIndent(scene, "", "  ")
}

func darkStrokeColor(color string) string {
	switch strings.ToLower(strings.TrimSpace(color)) {
	case "#000", "#000000", "#1e1e1e":
		return "#e5e7eb"
	case "#bbb", "#bbbbbb", "#aab7b8":
		return "#94a3b8"
	default:
		return color
	}
}

func darkFillColor(color string) string {
	switch strings.ToLower(strings.TrimSpace(color)) {
	case "#fff", "#ffffff":
		return "#111827"
	default:
		return color
	}
}
