package xaligo

import (
	"bytes"
	"context"
	"errors"
	"strings"
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

func TestDiagnoseReturnsSourcePosition(t *testing.T) {
	source := []byte("<frame>\n  <connection src=\"bad\" dst=\"2\" />\n</frame>")
	diagnostics, err := Diagnose(context.Background(), source)
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 1 || diagnostics[0].Line != 2 || diagnostics[0].Column != 3 {
		t.Fatalf("diagnostics = %#v", diagnostics)
	}
	if err := Validate(context.Background(), source); err == nil || !strings.Contains(err.Error(), "line 2, column 3") {
		t.Fatalf("Validate error = %v", err)
	}
}

func TestXYFlowPublicFormat(t *testing.T) {
	out, err := Render(context.Background(), apiTestXAL, RenderOptions{Format: FormatXYFlow})
	if err != nil || !bytes.Contains(out, []byte(`"nodes"`)) || !bytes.Contains(out, []byte(`"edges"`)) {
		t.Fatalf("XYFlow output = %s, err = %v", out, err)
	}
}

func TestIsoflowPublicFormat(t *testing.T) {
	out, err := Render(context.Background(), apiTestXAL, RenderOptions{Format: FormatIsoflow})
	if err != nil || !bytes.Contains(out, []byte(`"version": "3.3.0"`)) || !bytes.Contains(out, []byte(`"items"`)) || !bytes.Contains(out, []byte(`"views"`)) {
		t.Fatalf("Isoflow output = %s, err = %v", out, err)
	}
}
