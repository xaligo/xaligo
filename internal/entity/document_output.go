package entity

// RenderPage is one renderer-neutral document page represented as SVG.
// WidthPx and HeightPx are the intrinsic rendered SVG canvas dimensions. PDF
// and spreadsheet encoders use them at 96 DPI without distorting the artwork.
type RenderPage struct {
	ID       string
	SVG      []byte
	WidthPx  float64
	HeightPx float64
}
