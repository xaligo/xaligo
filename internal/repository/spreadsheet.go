//go:build !js

package repository

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"math"
	"strings"
	"unicode"
	"unicode/utf16"

	_ "github.com/tdewolff/canvas/svg"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xuri/excelize/v2"
)

const spreadsheetSheetNameLimit = 31

// SpreadsheetRepository encodes renderer-neutral SVG pages as an XLSX workbook.
type SpreadsheetRepository interface {
	Export(context.Context, []entity.RenderPage) ([]byte, error)
}

type spreadsheetRepository struct{}

func NewSpreadsheetRepository() SpreadsheetRepository {
	return &spreadsheetRepository{}
}

// Export places each RenderPage SVG at A1 on a dedicated worksheet.
func (rcvr *spreadsheetRepository) Export(ctx context.Context, pages []entity.RenderPage) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("export spreadsheet: %w", err)
	}
	if len(pages) == 0 {
		return nil, fmt.Errorf("export spreadsheet: at least one render page is required")
	}

	workbook := excelize.NewFile()
	closed := false
	defer func() {
		if !closed {
			_ = workbook.Close()
		}
	}()

	defaultSheet := workbook.GetSheetName(0)
	usedNames := make([]string, 0, len(pages))
	for index, page := range pages {
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("export spreadsheet page %d: %w", index+1, err)
		}
		if err := validateSpreadsheetRenderPage(page, index); err != nil {
			return nil, err
		}

		sheetName := uniqueSpreadsheetSheetName(page.ID, index, usedNames)
		usedNames = append(usedNames, sheetName)
		if index == 0 {
			if err := workbook.SetSheetName(defaultSheet, sheetName); err != nil {
				return nil, fmt.Errorf("export spreadsheet page %d (%q): rename worksheet: %w", index+1, page.ID, err)
			}
		} else if _, err := workbook.NewSheet(sheetName); err != nil {
			return nil, fmt.Errorf("export spreadsheet page %d (%q): create worksheet: %w", index+1, page.ID, err)
		}

		svg, config, err := prepareSpreadsheetSVG(page.SVG)
		if err != nil {
			return nil, fmt.Errorf("export spreadsheet page %d (%q): %w", index+1, page.ID, err)
		}
		if err := ctx.Err(); err != nil {
			return nil, fmt.Errorf("export spreadsheet page %d: %w", index+1, err)
		}
		if err := workbook.AddPictureFromBytes(sheetName, "A1", &excelize.Picture{
			Extension: ".svg",
			File:      svg,
			Format: &excelize.GraphicOptions{
				AltText:         sheetName,
				Name:            sheetName,
				LockAspectRatio: true,
				Positioning:     "oneCell",
				ScaleX:          page.WidthPx / float64(config.Width),
				ScaleY:          page.HeightPx / float64(config.Height),
			},
		}); err != nil {
			return nil, fmt.Errorf("export spreadsheet page %d (%q): add SVG: %w", index+1, page.ID, err)
		}
	}

	workbook.SetActiveSheet(0)
	output, err := workbook.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("export spreadsheet: finalize workbook: %w", err)
	}
	if err := workbook.Close(); err != nil {
		return nil, fmt.Errorf("export spreadsheet: close workbook: %w", err)
	}
	closed = true
	return bytes.Clone(output.Bytes()), nil
}

func validateSpreadsheetRenderPage(page entity.RenderPage, index int) error {
	if len(bytes.TrimSpace(page.SVG)) == 0 {
		return fmt.Errorf("export spreadsheet page %d (%q): SVG is required", index+1, page.ID)
	}
	if !positiveFiniteSpreadsheetDimension(page.WidthPx) || !positiveFiniteSpreadsheetDimension(page.HeightPx) {
		return fmt.Errorf("export spreadsheet page %d (%q): dimensions must be positive and finite", index+1, page.ID)
	}
	return nil
}

func positiveFiniteSpreadsheetDimension(value float64) bool {
	return value > 0 && !math.IsNaN(value) && !math.IsInf(value, 0)
}

func prepareSpreadsheetSVG(data []byte) ([]byte, image.Config, error) {
	trimmed := bytes.TrimSpace(data)
	start := bytes.Index(trimmed, []byte("<svg"))
	if start < 0 {
		return nil, image.Config{}, fmt.Errorf("parse SVG: expected SVG tag")
	}
	svg := bytes.Clone(trimmed[start:])
	config, format, err := image.DecodeConfig(bytes.NewReader(svg))
	if err != nil {
		return nil, image.Config{}, fmt.Errorf("parse SVG: %w", err)
	}
	if format != "svg" || config.Width <= 0 || config.Height <= 0 {
		return nil, image.Config{}, fmt.Errorf("parse SVG: dimensions must be positive")
	}
	return svg, config, nil
}

func uniqueSpreadsheetSheetName(id string, index int, used []string) string {
	base := sanitizeSpreadsheetSheetName(id, index)
	if !containsSpreadsheetSheetName(used, base) {
		return base
	}
	for sequence := 2; ; sequence++ {
		suffix := fmt.Sprintf(" (%d)", sequence)
		candidate := truncateSpreadsheetSheetName(base, spreadsheetSheetNameLimit-spreadsheetSheetNameLength(suffix)) + suffix
		if !containsSpreadsheetSheetName(used, candidate) {
			return candidate
		}
	}
}

func sanitizeSpreadsheetSheetName(id string, index int) string {
	name := strings.TrimSpace(id)
	if name == "" {
		name = fmt.Sprintf("Frame %d", index+1)
	}
	runes := []rune(name)
	for position, value := range runes {
		if unicode.IsControl(value) || strings.ContainsRune(`:\/?*[]`, value) {
			runes[position] = '_'
		}
	}
	if len(runes) > 0 && runes[0] == '\'' {
		runes[0] = '_'
	}
	if len(runes) > 0 && runes[len(runes)-1] == '\'' {
		runes[len(runes)-1] = '_'
	}
	name = strings.TrimSpace(string(runes))
	if name == "" {
		name = fmt.Sprintf("Frame %d", index+1)
	}
	return truncateSpreadsheetSheetName(name, spreadsheetSheetNameLimit)
}

func containsSpreadsheetSheetName(names []string, candidate string) bool {
	for _, name := range names {
		if strings.EqualFold(name, candidate) {
			return true
		}
	}
	return false
}

func truncateSpreadsheetSheetName(value string, maxLength int) string {
	if maxLength <= 0 {
		return ""
	}
	length := 0
	runes := make([]rune, 0, len(value))
	for _, current := range value {
		currentLength := utf16.RuneLen(current)
		if currentLength < 0 || length+currentLength > maxLength {
			break
		}
		runes = append(runes, current)
		length += currentLength
	}
	return string(runes)
}

func spreadsheetSheetNameLength(value string) int {
	return len(utf16.Encode([]rune(value)))
}
