// Package pptxplan turns an Excalidraw scene (the JSON produced by the xaligo
// renderer) into a flat, fully-resolved PPTX draw plan.
//
// All geometry — bounds detection, paper-size scaling and centring, obstacle
// collection, connector anchoring and orthogonal arrow routing, pixel→inch
// conversion and colour normalisation — happens here on the Go side. The
// TypeScript layer only iterates the resulting Plan and issues the matching
// PptxGenJS drawing calls; it performs no geometry of its own.
package pptxplan

// ── Input: Excalidraw scene (subset of fields used by the exporter) ──────────

type Scene struct {
	Elements []Element            `json:"elements"`
	Files    map[string]SceneFile `json:"files"`
	AppState *AppState            `json:"appState"`
}

type AppState struct {
	ViewBackgroundColor string `json:"viewBackgroundColor"`
}

type SceneFile struct {
	DataURL string `json:"dataURL"`
}

type Binding struct {
	ElementID  string    `json:"elementId"`
	FixedPoint []float64 `json:"fixedPoint"`
	Gap        float64   `json:"gap"`
}

type Element struct {
	ID              string      `json:"id"`
	Type            string      `json:"type"`
	X               float64     `json:"x"`
	Y               float64     `json:"y"`
	Width           float64     `json:"width"`
	Height          float64     `json:"height"`
	Angle           float64     `json:"angle"`
	Opacity         *float64    `json:"opacity"`
	StrokeColor     string      `json:"strokeColor"`
	BackgroundColor string      `json:"backgroundColor"`
	StrokeWidth     float64     `json:"strokeWidth"`
	StrokeStyle     string      `json:"strokeStyle"`
	Text            string      `json:"text"`
	RawText         string      `json:"rawText"`
	FontSize        *float64    `json:"fontSize"`
	FontFamily      *int        `json:"fontFamily"`
	FontStyle       string      `json:"fontStyle"`
	TextAlign       string      `json:"textAlign"`
	VerticalAlign   string      `json:"verticalAlign"`
	FileID          string      `json:"fileId"`
	Points          [][]float64 `json:"points"`
	IsDeleted       bool        `json:"isDeleted"`
	StartBinding    *Binding    `json:"startBinding"`
	EndBinding      *Binding    `json:"endBinding"`
	CustomData      *CustomData `json:"customData"`
}

type CustomData struct {
	ConnectorKind           string `json:"xaligoConnectorKind"`
	ConnectorStartArrowhead string `json:"xaligoConnectorStartArrowhead"`
	ConnectorEndArrowhead   string `json:"xaligoConnectorEndArrowhead"`
	Junction                bool   `json:"xaligoJunction,omitempty"`
}

// ── Options driving the calculations ─────────────────────────────────────────

// Options collects every parameter that influences the geometry of the plan.
// They originate from the CLI / Go controller and are passed verbatim to the
// WASM plan builder as JSON.
type Options struct {
	Theme         string        `json:"theme,omitempty"`
	PxPerInch     float64       `json:"pxPerInch"`
	ArrowStyle    string        `json:"arrowStyle"`
	ArrowStubPx   float64       `json:"arrowStubPx"`
	ArrowMargin   float64       `json:"arrowMarginPx"`
	PaperSize     string        `json:"paperSize"`
	Orientation   string        `json:"orientation"`
	LegendEntries []LegendEntry `json:"legendEntries"`
}

// ── Output: the PPTX draw plan ───────────────────────────────────────────────

// Plan is the complete, ordered list of drawing operations plus slide metadata.
// Every coordinate is already in inches and every colour is a 6-hex string.
type Plan struct {
	Slide  PlanSlide     `json:"slide"`
	Ops    []DrawOp      `json:"ops"`
	Legend []LegendEntry `json:"legend,omitempty"`
}

type LegendEntry struct {
	CatalogID    int    `json:"catalogId"`
	Abbreviation string `json:"abbreviation"`
	OfficialName string `json:"officialName"`
	Data         string `json:"data,omitempty"`
}

type PlanSlide struct {
	W          float64 `json:"w"`
	H          float64 `json:"h"`
	Background string  `json:"background"`
}

// DrawOp is a single PptxGenJS drawing call. Kind selects the dispatch:
// "rect" | "ellipse" | "text" | "image" | "line".
type DrawOp struct {
	Kind   string  `json:"kind"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	W      float64 `json:"w"`
	H      float64 `json:"h"`
	Rotate float64 `json:"rotate,omitempty"`

	// rect / ellipse / line
	Line *LineStyle `json:"line,omitempty"`
	Fill *FillStyle `json:"fill,omitempty"`

	// text
	Text     string  `json:"text,omitempty"`
	Color    string  `json:"color,omitempty"`
	FontFace string  `json:"fontFace,omitempty"`
	FontSize float64 `json:"fontSize,omitempty"`
	Bold     bool    `json:"bold,omitempty"`
	Align    string  `json:"align,omitempty"`
	Valign   string  `json:"valign,omitempty"`

	// image
	Data         string  `json:"data,omitempty"`
	Transparency float64 `json:"transparency,omitempty"`

	// line / polyline (points are relative to the op's x/y bbox origin, inches)
	Points []PtIn `json:"points,omitempty"`
	FlipH  bool   `json:"flipH,omitempty"`
	FlipV  bool   `json:"flipV,omitempty"`
}

type LineStyle struct {
	Color          string  `json:"color"`
	Width          float64 `json:"width"`
	Dash           string  `json:"dash"`
	Transparency   float64 `json:"transparency"`
	BeginArrowType string  `json:"beginArrowType,omitempty"`
	EndArrowType   string  `json:"endArrowType,omitempty"`
}

type FillStyle struct {
	Color        string  `json:"color"`
	Transparency float64 `json:"transparency"`
}

// PtIn is an inch-space point for a polyline op. MoveTo marks the first vertex.
type PtIn struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	MoveTo bool    `json:"moveTo,omitempty"`
}
