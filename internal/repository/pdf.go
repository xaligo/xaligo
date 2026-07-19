//go:build !js

package repository

import (
	"bytes"
	"context"
	"fmt"
	"math"

	"github.com/tdewolff/canvas"
	canvaspdf "github.com/tdewolff/canvas/renderers/pdf"
	"github.com/xaligo/xaligo/internal/entity"
)

const pdfPixelsPerInch = 96.0

// PDFRepository encodes renderer-neutral SVG pages as a PDF document.
type PDFRepository interface {
	Export(context.Context, []entity.RenderPage) ([]byte, error)
}

type pdfRepository struct{}

func NewPDFRepository() PDFRepository {
	return &pdfRepository{}
}

// Export renders each RenderPage as one PDF page, preserving input order.
func (rcvr *pdfRepository) Export(ctx context.Context, pages []entity.RenderPage) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("export PDF: %w", err)
	}
	if len(pages) == 0 {
		return nil, fmt.Errorf("export PDF: at least one render page is required")
	}
	composer, err := newPDFSVGComposer()
	if err != nil {
		return nil, fmt.Errorf("export PDF: %w", err)
	}

	var output bytes.Buffer
	var renderer *canvaspdf.PDF
	for index, page := range pages {
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("export PDF page %d: %w", index+1, err)
		}
		if err := validatePDFRenderPage(page, index); err != nil {
			return nil, err
		}

		drawing, err := composer.parse(page.SVG)
		if err != nil {
			return nil, fmt.Errorf("export PDF page %d (%q): parse SVG: %w", index+1, page.ID, err)
		}
		drawingWidth, drawingHeight := drawing.Size()
		if !positiveFinitePDFDimension(drawingWidth) || !positiveFinitePDFDimension(drawingHeight) {
			return nil, fmt.Errorf("export PDF page %d (%q): parsed SVG dimensions must be positive and finite", index+1, page.ID)
		}

		pageWidth := page.WidthPx * 25.4 / pdfPixelsPerInch
		pageHeight := page.HeightPx * 25.4 / pdfPixelsPerInch
		if !positiveFinitePDFDimension(pageWidth) || !positiveFinitePDFDimension(pageHeight) {
			return nil, fmt.Errorf("export PDF page %d (%q): converted page dimensions must be positive and finite", index+1, page.ID)
		}
		if renderer == nil {
			renderer = canvaspdf.New(&output, pageWidth, pageHeight, nil)
		} else {
			renderer.NewPage(pageWidth, pageHeight)
		}
		scale := math.Min(pageWidth/drawingWidth, pageHeight/drawingHeight)
		offsetX := (pageWidth - drawingWidth*scale) / 2
		offsetY := (pageHeight - drawingHeight*scale) / 2
		view := canvas.Identity.Translate(offsetX, offsetY).Scale(scale, scale)
		drawing.RenderViewTo(renderer, view)
	}

	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("export PDF: %w", err)
	}
	if err := renderer.Close(); err != nil {
		return nil, fmt.Errorf("export PDF: finalize document: %w", err)
	}
	return output.Bytes(), nil
}

func validatePDFRenderPage(page entity.RenderPage, index int) error {
	if len(bytes.TrimSpace(page.SVG)) == 0 {
		return fmt.Errorf("export PDF page %d (%q): SVG is required", index+1, page.ID)
	}
	if !positiveFinitePDFDimension(page.WidthPx) || !positiveFinitePDFDimension(page.HeightPx) {
		return fmt.Errorf("export PDF page %d (%q): dimensions must be positive and finite", index+1, page.ID)
	}
	return nil
}

func positiveFinitePDFDimension(value float64) bool {
	return value > 0 && !math.IsNaN(value) && !math.IsInf(value, 0)
}
