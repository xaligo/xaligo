package usecase

import v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"

type ThemeUsecase interface {
	NormalizeTheme(string) (string, error)
	ApplyThemeJSON([]byte, string) ([]byte, error)
}

type themeUsecase struct{}

func NewThemeUsecase() ThemeUsecase {
	return &themeUsecase{}
}

const (
	ThemeLight = v1engine.ThemeLightV1EngineThemeApply
	ThemeDark  = v1engine.ThemeDarkV1EngineThemeApply
)

func (rcvr *themeUsecase) NormalizeTheme(theme string) (string, error) {
	return v1engine.NormalizeThemeV1EngineThemeApply(theme)
}

func (rcvr *themeUsecase) ApplyThemeJSON(sceneJSON []byte, theme string) ([]byte, error) {
	return v1engine.ApplyThemeJSONV1EngineThemeApply(sceneJSON, theme)
}

// NormalizeTheme delegates theme normalization to ThemeUsecase.
// Deprecated: use NewThemeUsecase().NormalizeTheme instead.
func NormalizeTheme(theme string) (string, error) {
	return NewThemeUsecase().NormalizeTheme(theme)
}

// ApplyThemeJSON delegates canonical-scene theme application to ThemeUsecase.
// Deprecated: use NewThemeUsecase().ApplyThemeJSON instead.
func ApplyThemeJSON(sceneJSON []byte, theme string) ([]byte, error) {
	return NewThemeUsecase().ApplyThemeJSON(sceneJSON, theme)
}
