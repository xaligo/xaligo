package engine

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const (
	ThemeLightV1EngineThemeApply = entity.ThemeLight
	ThemeDarkV1EngineThemeApply  = entity.ThemeDark
)

// NormalizeThemeV1EngineThemeApply validates a renderer theme and supplies the default.
func NormalizeThemeV1EngineThemeApply(theme string) (string, error) {
	return entity.NormalizeTheme(theme)
}

// ApplyThemeJSONV1EngineThemeApply applies presentation colors to a rendered Excalidraw scene.
// Service colors and embedded icons are intentionally preserved; only the
// neutral canvas, text, borders and white surfaces are remapped.
func ApplyThemeJSONV1EngineThemeApply(sceneJSON []byte, theme string) ([]byte, error) {
	normalized, err := NormalizeThemeV1EngineThemeApply(theme)
	if err != nil {
		return nil, err
	}
	if normalized == ThemeLightV1EngineThemeApply {
		return sceneJSON, nil
	}

	var scene fileV1EngineSceneTypes
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		return nil, fmt.Errorf("decode excalidraw scene for theme: %w", err)
	}
	if scene.AppState == nil {
		scene.AppState = map[string]any{}
	}
	scene.AppState["viewBackgroundColor"] = "#111827"

	for _, element := range scene.Elements {
		if color, ok := element["strokeColor"].(string); ok {
			element["strokeColor"] = darkStrokeColorV1EngineThemeApply(color)
		}
		if color, ok := element["backgroundColor"].(string); ok {
			element["backgroundColor"] = darkFillColorV1EngineThemeApply(color)
		}
	}
	return json.MarshalIndent(scene, "", "  ")
}

func darkStrokeColorV1EngineThemeApply(color string) string {
	switch strings.ToLower(strings.TrimSpace(color)) {
	case "#000", "#000000", "#1e1e1e":
		return "#e5e7eb"
	case "#bbb", "#bbbbbb", "#aab7b8":
		return "#94a3b8"
	default:
		return color
	}
}

func darkFillColorV1EngineThemeApply(color string) string {
	switch strings.ToLower(strings.TrimSpace(color)) {
	case "#fff", "#ffffff":
		return "#111827"
	default:
		return color
	}
}
