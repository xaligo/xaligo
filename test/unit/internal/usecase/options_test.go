package usecase_test

import (
	"strings"
	"testing"
	"testing/fstest"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/usecase"
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
		{"margin", entity.RenderOptions{PaperMarginLeftIn: -0.1}, "paper margins"},
		{"theme", entity.RenderOptions{Theme: "neon"}, "unknown theme"},
		{"asset fs", entity.RenderOptions{Assets: &entity.AssetSource{CatalogCSV: "catalog.csv", GroupIconsDir: "groups"}}, "filesystem"},
		{"asset paths", entity.RenderOptions{Assets: &entity.AssetSource{FS: fstest.MapFS{}, CatalogCSV: " ", GroupIconsDir: "groups"}}, "catalog and group icons"},
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
