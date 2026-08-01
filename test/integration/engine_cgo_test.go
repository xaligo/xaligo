//go:build cgo && xaligo_engine

package integration_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func engineFloat(value float64) *float64 {
	return &value
}

func TestRustStaticLibraryEngineThroughCgo(t *testing.T) {
	spec := entity.EngineDocumentSpec{
		Direction: entity.EngineDirectionVertical,
		Width:     120,
		Height:    100,
		Gap:       10,
		Elements: []entity.EngineElementSpec{
			{ID: "header", Width: engineFloat(120), Height: engineFloat(20)},
			{ID: "body", Width: engineFloat(100), Weight: engineFloat(1)},
		},
	}
	engine := v2.NewEngineUsecase()

	resolved, err := engine.Resolve(context.Background(), spec)
	if err != nil {
		t.Fatal(err)
	}
	if len(resolved.Elements) != 2 {
		t.Fatalf("resolved elements = %#v", resolved.Elements)
	}
	body := resolved.Elements[1]
	if body.ID != "body" || body.X != 0 || body.Y != 30 || body.Width != 100 || body.Height != 70 {
		t.Fatalf("resolved body = %#v", body)
	}

	svg, err := engine.RenderSVG(context.Background(), spec)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(svg, []byte(`<svg`)) || !bytes.Contains(svg, []byte(`id="body"`)) {
		t.Fatalf("Rust SVG projection = %s", svg)
	}
}

func TestRustSVGNormalizationThroughCgo(t *testing.T) {
	engine := v2.NewEngineUsecase()
	normalized, err := engine.NormalizeSVG(context.Background(), []byte(
		`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="8"><path d="M0 0h16v8z"/></svg>`,
	))
	if err != nil {
		t.Fatal(err)
	}
	if normalized.Width != 16 || normalized.Height != 8 || normalized.ViewBox != "0 0 16 8" || !bytes.Contains(normalized.Data, []byte(`<svg`)) {
		t.Fatalf("normalized SVG = %#v %s", normalized, normalized.Data)
	}
	for _, input := range [][]byte{
		[]byte(`<svg xmlns="http://www.w3.org/2000/svg"><script>alert(1)</script></svg>`),
		[]byte(`<svg xmlns="http://www.w3.org/2000/svg"><image href="https://example.com/icon.png"/></svg>`),
		[]byte(`<svg`),
		make([]byte, 2*1024*1024+1),
	} {
		if _, err := engine.NormalizeSVG(context.Background(), input); err == nil {
			t.Fatalf("unsafe or oversized SVG was accepted: %q", input[:min(len(input), 80)])
		}
	}
}
