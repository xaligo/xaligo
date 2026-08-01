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
