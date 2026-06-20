package xaligo

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

var apiTestXAL = []byte(`<frame width="320" height="180"><box title="API" /></frame>`)

func TestRenderPublicFormats(t *testing.T) {
	excal, err := RenderExcalidraw(context.Background(), apiTestXAL, RenderOptions{Theme: "dark"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(excal, []byte(`"viewBackgroundColor": "#111827"`)) {
		t.Fatalf("dark Excalidraw theme missing")
	}
	svg, err := Render(context.Background(), apiTestXAL, RenderOptions{Format: FormatSVG})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(svg, []byte(`<svg`)) {
		t.Fatalf("SVG output missing root element")
	}
}

func TestValidatePublicAPI(t *testing.T) {
	if err := Validate(context.Background(), apiTestXAL); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Validate(ctx, apiTestXAL); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled Validate error = %v", err)
	}
}

func TestFutureFormatReturnsStableSentinel(t *testing.T) {
	_, err := Render(context.Background(), apiTestXAL, RenderOptions{Format: FormatXYFlow})
	if !errors.Is(err, ErrNotImplemented) {
		t.Fatalf("XYFlow error = %v", err)
	}
}
