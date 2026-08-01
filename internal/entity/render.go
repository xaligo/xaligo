package entity

import "io/fs"

type Mode string
type Format string
type DiagnosticSeverity string

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

	Title            string `json:"title,omitempty"`
	Author           string `json:"author,omitempty"`
	Company          string `json:"company,omitempty"`
	Subject          string `json:"subject,omitempty"`
	Compression      *bool  `json:"compression,omitempty"`
	PPTXExporterWASM string `json:"pptxExporterWasm,omitempty"`
}
