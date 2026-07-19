package repository_test

import (
	"bytes"
	"context"
	"errors"
	"image"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xuri/excelize/v2"
)

func TestSpreadsheetRepositoryExportsOneWorksheetPerRenderPage(t *testing.T) {
	longID := strings.Repeat("a", 40)
	ids := []string{
		"overview",
		`bad:/?*[]\name`,
		longID,
		longID,
		"OVERVIEW",
		"'quoted'",
		"",
	}
	pages := make([]entity.RenderPage, 0, len(ids))
	for _, id := range ids {
		pages = append(pages, entity.RenderPage{
			ID: id, SVG: spreadsheetTestSVG(), WidthPx: 320, HeightPx: 180,
		})
	}

	output, err := repository.NewSpreadsheetRepository().Export(context.Background(), pages)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(output, []byte("PK")) {
		t.Fatalf("XLSX ZIP header missing: %q", output[:min(len(output), 8)])
	}

	workbook, err := excelize.OpenReader(bytes.NewReader(output))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := workbook.Close(); err != nil {
			t.Errorf("Close() error = %v", err)
		}
	}()
	wantNames := []string{
		"overview",
		"bad_______name",
		strings.Repeat("a", 31),
		strings.Repeat("a", 27) + " (2)",
		"OVERVIEW (2)",
		"_quoted_",
		"Frame 7",
	}
	if names := workbook.GetSheetList(); !reflect.DeepEqual(names, wantNames) {
		t.Fatalf("worksheet names = %#v, want %#v", names, wantNames)
	}
	for _, sheet := range wantNames {
		pictures, err := workbook.GetPictures(sheet, "A1")
		if err != nil {
			t.Fatalf("GetPictures(%q) error = %v", sheet, err)
		}
		if len(pictures) != 1 || pictures[0].Extension != ".svg" || !bytes.HasPrefix(pictures[0].File, []byte("<svg")) {
			t.Fatalf("GetPictures(%q) = %#v", sheet, pictures)
		}
		config, format, err := image.DecodeConfig(bytes.NewReader(pictures[0].File))
		if err != nil || format != "svg" || pictures[0].Format == nil {
			t.Fatalf("DecodeConfig(%q) = %#v, %q, %v", sheet, config, format, err)
		}
		width := int(math.Round(float64(config.Width) * pictures[0].Format.ScaleX))
		height := int(math.Round(float64(config.Height) * pictures[0].Format.ScaleY))
		if width != 320 || height != 180 {
			t.Fatalf("picture size on %q = %dx%d, want 320x180", sheet, width, height)
		}
	}
}

func TestSpreadsheetRepositoryRejectsInvalidInputs(t *testing.T) {
	tests := []struct {
		name  string
		pages []entity.RenderPage
		want  string
	}{
		{name: "empty", want: "at least one render page"},
		{name: "missing SVG", pages: []entity.RenderPage{{ID: "empty", WidthPx: 100, HeightPx: 50}}, want: "SVG is required"},
		{name: "negative width", pages: []entity.RenderPage{{ID: "negative", SVG: spreadsheetTestSVG(), WidthPx: -1, HeightPx: 50}}, want: "dimensions must be positive and finite"},
		{name: "non-finite height", pages: []entity.RenderPage{{ID: "infinite", SVG: spreadsheetTestSVG(), WidthPx: 100, HeightPx: math.Inf(1)}}, want: "dimensions must be positive and finite"},
		{name: "invalid SVG", pages: []entity.RenderPage{{ID: "invalid", SVG: []byte(`<svg`), WidthPx: 100, HeightPx: 50}}, want: "parse SVG"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := repository.NewSpreadsheetRepository().Export(context.Background(), test.pages)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Export() error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestSpreadsheetRepositoryHonorsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := repository.NewSpreadsheetRepository().Export(ctx, []entity.RenderPage{{
		ID: "page", SVG: spreadsheetTestSVG(), WidthPx: 100, HeightPx: 50,
	}})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Export() error = %v, want context.Canceled", err)
	}
}

func spreadsheetTestSVG() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="320" height="180" viewBox="0 0 320 180">
  <rect x="0" y="0" width="320" height="180" fill="#2563eb"/>
</svg>`)
}
