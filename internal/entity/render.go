package entity

import (
	"fmt"
	"io/fs"
	"strings"
)

type Mode string
type Format string
type DiagnosticSeverity string

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

// AssetSource describes an embedded or virtual asset tree.
type AssetSource struct {
	FS               fs.FS
	CatalogCSV       string
	GroupIconsDir    string
	IsoflowIconsJSON string
	ItemIconSize     float64
}

// RenderOptions contains renderer-independent presentation and output options.
type RenderOptions struct {
	Mode          Mode           `json:"mode,omitempty"`
	Format        Format         `json:"format,omitempty"`
	Theme         string         `json:"theme,omitempty"`
	ServicesCSV   []byte         `json:"-"`
	Abbreviations map[int]string `json:"-"`
	Assets        *AssetSource   `json:"-"`

	PxPerInch           float64 `json:"pxPerInch,omitempty"`
	ArrowStyle          string  `json:"arrowStyle,omitempty"`
	ArrowStubPx         float64 `json:"arrowStubPx,omitempty"`
	ArrowMarginPx       float64 `json:"arrowMarginPx,omitempty"`
	PaperSize           string  `json:"paperSize,omitempty"`
	Orientation         string  `json:"orientation,omitempty"`
	PaperMarginIn       float64 `json:"paperMargin,omitempty"`
	PaperMarginTopIn    float64 `json:"paperMarginTop,omitempty"`
	PaperMarginRightIn  float64 `json:"paperMarginRight,omitempty"`
	PaperMarginBottomIn float64 `json:"paperMarginBottom,omitempty"`
	PaperMarginLeftIn   float64 `json:"paperMarginLeft,omitempty"`

	Title            string `json:"title,omitempty"`
	Author           string `json:"author,omitempty"`
	Company          string `json:"company,omitempty"`
	Subject          string `json:"subject,omitempty"`
	Compression      *bool  `json:"compression,omitempty"`
	PPTXExporterWASM string `json:"pptxExporterWasm,omitempty"`
}

type Diagnostic struct {
	Severity DiagnosticSeverity `json:"severity"`
	Message  string             `json:"message"`
	Offset   int                `json:"offset,omitempty"`
	Line     int                `json:"line,omitempty"`
	Column   int                `json:"column,omitempty"`
}

type DiagnosticsError struct {
	Diagnostics []Diagnostic
}

func (e *DiagnosticsError) Error() string {
	if len(e.Diagnostics) == 0 {
		return "validation failed"
	}
	d := e.Diagnostics[0]
	if d.Line > 0 {
		return fmt.Sprintf("line %d, column %d: %s", d.Line, d.Column, d.Message)
	}
	return d.Message
}
