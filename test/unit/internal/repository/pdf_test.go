package repository_test

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/base64"
	"errors"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"regexp"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
)

func TestPDFRepositoryExportsOnePDFPagePerRenderPage(t *testing.T) {
	pages := []entity.RenderPage{
		{ID: "overview", SVG: pdfTestSVG("#2563eb"), WidthPx: 192, HeightPx: 96},
		{ID: "detail", SVG: pdfTestSVG("#dc2626"), WidthPx: 96, HeightPx: 192},
	}

	output, err := repository.NewPDFRepository().Export(context.Background(), pages)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(string(output), "%PDF-") {
		t.Fatalf("PDF header missing: %q", output[:min(len(output), 16)])
	}
	pageObjects := regexp.MustCompile(`/Type\s*/Page\b`).FindAll(output, -1)
	if len(pageObjects) != len(pages) {
		t.Fatalf("PDF page objects = %d, want %d", len(pageObjects), len(pages))
	}
	for _, mediaBox := range []string{
		"/MediaBox[0 0 144 72]",
		"/MediaBox[0 0 72 144]",
	} {
		if !strings.Contains(string(output), mediaBox) {
			found := regexp.MustCompile(`/MediaBox\s*\[[^]]+\]`).FindAll(output, -1)
			t.Fatalf("PDF missing %q; media boxes = %q", mediaBox, found)
		}
	}
}

func TestPDFRepositoryRejectsInvalidInputs(t *testing.T) {
	tests := []struct {
		name  string
		pages []entity.RenderPage
		want  string
	}{
		{name: "empty", want: "at least one render page"},
		{name: "missing SVG", pages: []entity.RenderPage{{ID: "empty", WidthPx: 100, HeightPx: 50}}, want: "SVG is required"},
		{name: "zero width", pages: []entity.RenderPage{{ID: "zero", SVG: pdfTestSVG("#000000"), HeightPx: 50}}, want: "dimensions must be positive and finite"},
		{name: "non-finite height", pages: []entity.RenderPage{{ID: "nan", SVG: pdfTestSVG("#000000"), WidthPx: 100, HeightPx: math.NaN()}}, want: "dimensions must be positive and finite"},
		{name: "invalid SVG", pages: []entity.RenderPage{{ID: "invalid", SVG: []byte(`not-svg`), WidthPx: 100, HeightPx: 50}}, want: "parse SVG"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := repository.NewPDFRepository().Export(context.Background(), test.pages)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Export() error = %v, want containing %q", err, test.want)
			}
		})
	}
}

func TestPDFRepositoryHonorsCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := repository.NewPDFRepository().Export(ctx, []entity.RenderPage{{
		ID: "page", SVG: pdfTestSVG("#000000"), WidthPx: 100, HeightPx: 50,
	}})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Export() error = %v, want context.Canceled", err)
	}
}

func TestPDFRepositoryPreservesTspanAndDataImagesWithoutSystemFonts(t *testing.T) {
	embeddedSVG := base64.StdEncoding.EncodeToString([]byte(`
<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20">
  <rect x="1" y="1" width="18" height="18" fill="#00ff00"/>
</svg>`))
	// One opaque red PNG pixel. A raster image is included alongside the SVG
	// image so both PDF image paths are covered.
	pixel := image.NewRGBA(image.Rect(0, 0, 1, 1))
	pixel.Set(0, 0, color.RGBA{R: 255, A: 255})
	var pngData bytes.Buffer
	if err := png.Encode(&pngData, pixel); err != nil {
		t.Fatal(err)
	}
	embeddedPNG := base64.StdEncoding.EncodeToString(pngData.Bytes())
	svg := []byte(`
<svg xmlns="http://www.w3.org/2000/svg" width="200" height="100" viewBox="0 0 200 100">
  <rect x="0" y="0" width="200" height="100" fill="#ffffff"/>
  <text x="10" y="24" fill="#123456" font-family="font-that-is-not-installed" font-size="12">
    <tspan x="10" y="24">Alpha</tspan>
    <tspan x="10" dy="18" font-weight="700">Beta</tspan>
  </text>
  <image x="90" y="10" width="40" height="40" href="data:image/svg+xml;base64,` + embeddedSVG + `"/>
  <image x="150" y="10" width="20" height="20" href="data:image/png;base64,` + embeddedPNG + `"/>
</svg>`)

	output, err := repository.NewPDFRepository().Export(context.Background(), []entity.RenderPage{{
		ID: "text-and-images", SVG: svg, WidthPx: 200, HeightPx: 100,
	}})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(output, []byte("/Font")) {
		t.Fatal("PDF has no font resource; tspan text was not retained")
	}
	if !regexp.MustCompile(`/Subtype\s*/Image\b`).Match(output) {
		t.Fatal("PDF has no raster image object; data:image/png was not retained")
	}
	inflated := inflatePDFStreams(t, output)
	for _, want := range []string{" BT", "0 1 0 rg"} {
		if !strings.Contains(inflated, want) {
			t.Fatalf("inflated PDF streams do not contain %q; streams = %q", want, inflated)
		}
	}
}

func TestPDFRepositoryUsesUniformScaleForMismatchedSVGDimensions(t *testing.T) {
	output, err := repository.NewPDFRepository().Export(context.Background(), []entity.RenderPage{{
		ID: "mismatched", SVG: pdfTestSVG("#2563eb"), WidthPx: 96, HeightPx: 96,
	}})
	if err != nil {
		t.Fatal(err)
	}
	inflated := inflatePDFStreams(t, output)
	matrices := regexp.MustCompile(`(-?[0-9.]+) 0 0 (-?[0-9.]+) -?[0-9.]+ -?[0-9.]+ cm`).FindAllStringSubmatch(inflated, -1)
	if len(matrices) == 0 {
		t.Fatalf("no axis-aligned PDF transform found in streams: %q", inflated)
	}
	for _, matrix := range matrices {
		if matrix[1] != matrix[2] {
			t.Fatalf("non-uniform PDF transform %q distorted the SVG", matrix[0])
		}
	}
}

func inflatePDFStreams(t *testing.T, document []byte) string {
	t.Helper()
	streamPattern := regexp.MustCompile(`(?s)stream\r?\n(.*?)\r?\nendstream`)
	var inflated strings.Builder
	for _, match := range streamPattern.FindAllSubmatch(document, -1) {
		reader, err := zlib.NewReader(bytes.NewReader(match[1]))
		if err != nil {
			continue
		}
		data, readErr := io.ReadAll(reader)
		closeErr := reader.Close()
		if readErr != nil {
			t.Fatalf("read compressed PDF stream: %v", readErr)
		}
		if closeErr != nil {
			t.Fatalf("close compressed PDF stream: %v", closeErr)
		}
		inflated.Write(data)
	}
	return inflated.String()
}

func pdfTestSVG(color string) []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="192" height="96" viewBox="0 0 192 96">
  <rect x="0" y="0" width="192" height="96" fill="` + color + `"/>
</svg>`)
}
