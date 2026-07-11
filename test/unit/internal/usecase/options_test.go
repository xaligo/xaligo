package usecase_test

import (
	"math"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestValidateRenderOptionsAcceptsSupportedFormatsAndModes(t *testing.T) {
	for _, format := range []entity.Format{"", usecase.FormatExcalidraw, usecase.FormatSVG, usecase.FormatPPTX, usecase.FormatXYFlow, usecase.FormatIsoflow} {
		if err := usecase.ValidateRenderOptions(entity.RenderOptions{Mode: " network ", Format: format, Theme: "dark"}); err != nil {
			t.Fatalf("format %q error = %v", format, err)
		}
	}
	if err := usecase.ValidateRenderOptions(entity.RenderOptions{Assets: &entity.AssetSource{FS: fstest.MapFS{}, CatalogCSV: "catalog.csv", GroupIconsDir: "groups"}}); err != nil {
		t.Fatal(err)
	}
}

func TestValidateRenderOptionsRejectsInvalidValues(t *testing.T) {
	cases := []struct {
		name string
		opts entity.RenderOptions
		want string
	}{
		{"format", entity.RenderOptions{Format: "bad"}, "unknown render format"},
		{"mode", entity.RenderOptions{Mode: "bad"}, "unknown render mode"},
		{"future mode", entity.RenderOptions{Mode: "aws-2.5d"}, "renderer not implemented"},
		{"negative margin", entity.RenderOptions{PaperMarginLeftIn: -0.1}, "paper margin left"},
		{"NaN PPI", entity.RenderOptions{PxPerInch: math.NaN()}, "pixels per inch must be finite"},
		{"infinite PPI", entity.RenderOptions{PxPerInch: math.Inf(1)}, "pixels per inch must be finite"},
		{"negative PPI", entity.RenderOptions{PxPerInch: -96}, "pixels per inch must be non-negative"},
		{"NaN arrow stub", entity.RenderOptions{ArrowStubPx: math.NaN()}, "arrow stub must be finite"},
		{"infinite arrow margin", entity.RenderOptions{ArrowMarginPx: math.Inf(-1)}, "arrow margin must be finite"},
		{"NaN paper margin", entity.RenderOptions{PaperMarginTopIn: math.NaN()}, "paper margin top must be finite"},
		{"arrow style", entity.RenderOptions{ArrowStyle: "wide"}, "unknown arrow style"},
		{"paper size", entity.RenderOptions{PaperSize: "A0"}, "unknown paper size"},
		{"orientation", entity.RenderOptions{Orientation: "sideways"}, "unknown paper orientation"},
		{"margin without paper", entity.RenderOptions{PaperMarginIn: 0.5}, "require a paper size"},
		{"oversized paper margins", entity.RenderOptions{PaperSize: "A4", PaperMarginIn: 100}, "no positive content area"},
		{"oversized portrait horizontal margins", entity.RenderOptions{PaperSize: "A4", Orientation: "portrait", PaperMarginLeftIn: 5, PaperMarginRightIn: 5}, "no positive content area"},
		{"theme", entity.RenderOptions{Theme: "neon"}, "unknown theme"},
		{"asset fs", entity.RenderOptions{Assets: &entity.AssetSource{CatalogCSV: "catalog.csv", GroupIconsDir: "groups"}}, "filesystem"},
		{"asset paths", entity.RenderOptions{Assets: &entity.AssetSource{FS: fstest.MapFS{}, CatalogCSV: " ", GroupIconsDir: "groups"}}, "catalog and group icons"},
		{"asset item size", entity.RenderOptions{Assets: &entity.AssetSource{FS: fstest.MapFS{}, CatalogCSV: "catalog.csv", GroupIconsDir: "groups", ItemIconSize: math.NaN()}}, "asset item icon size must be finite"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := usecase.ValidateRenderOptions(tc.opts)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v, want %q", err, tc.want)
			}
		})
	}
}
