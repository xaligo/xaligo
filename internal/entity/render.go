package entity

import "io/fs"

type Mode string
type Format string
type DiagnosticSeverity string

type TerminalStyle string
type TerminalLayout string
type TerminalDetail string
type TerminalColor string
type TerminalIcons string

const (
	TerminalStyleUnicode TerminalStyle = "unicode"
	TerminalStyleASCII   TerminalStyle = "ascii"

	TerminalLayoutDiagram  TerminalLayout = "diagram"
	TerminalLayoutSemantic TerminalLayout = "semantic"
	TerminalLayoutHybrid   TerminalLayout = "hybrid"

	TerminalDetailCompact TerminalDetail = "compact"
	TerminalDetailNormal  TerminalDetail = "normal"
	TerminalDetailFull    TerminalDetail = "full"

	TerminalColorAuto   TerminalColor = "auto"
	TerminalColorAlways TerminalColor = "always"
	TerminalColorNever  TerminalColor = "never"

	TerminalIconsLabel  TerminalIcons = "label"
	TerminalIconsSymbol TerminalIcons = "symbol"
	TerminalIconsNone   TerminalIcons = "none"
)

// AssetSource describes an embedded or virtual asset tree.
type AssetSource struct {
	FS            fs.FS
	CatalogCSV    string
	GroupIconsDir string
	ItemIconSize  float64
}

// ImportSource describes files referenced relative to the input document.
type ImportSource struct {
	FS fs.FS
}

// RenderOptions contains renderer-independent presentation and output options.
type RenderOptions struct {
	Mode          Mode          `json:"mode,omitempty"`
	Format        Format        `json:"format,omitempty"`
	Theme         string        `json:"theme,omitempty"`
	ServicesCSV   []byte        `json:"-"`
	Assets        *AssetSource  `json:"-"`
	Imports       *ImportSource `json:"-"`
	CombineFrames bool          `json:"combineFrames,omitempty"`

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
	SVGLegendPosition   string  `json:"svgLegendPosition,omitempty"`

	TerminalStyle  TerminalStyle  `json:"terminalStyle,omitempty"`
	TerminalLayout TerminalLayout `json:"terminalLayout,omitempty"`
	TerminalDetail TerminalDetail `json:"terminalDetail,omitempty"`
	TerminalColor  TerminalColor  `json:"terminalColor,omitempty"`
	TerminalIcons  TerminalIcons  `json:"terminalIcons,omitempty"`
	TerminalWidth  int            `json:"terminalWidth,omitempty"`
	TerminalHeight int            `json:"terminalHeight,omitempty"`
	TerminalFocus  string         `json:"terminalFocus,omitempty"`

	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	Company     string `json:"company,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Compression *bool  `json:"compression,omitempty"`
}
