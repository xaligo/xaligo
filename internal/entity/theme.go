package entity

import (
	"fmt"
	"strings"
)

const (
	ThemeLight = "light"
	ThemeDark  = "dark"
)

// NormalizeTheme validates a renderer theme and supplies the default.
func NormalizeTheme(value string) (string, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	if value == "" {
		return ThemeLight, nil
	}
	if value != ThemeLight && value != ThemeDark {
		return "", fmt.Errorf("unknown theme %q; valid: light, dark", value)
	}
	return value, nil
}
