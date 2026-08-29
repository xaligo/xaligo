package integration

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

var apiTestXAL = []byte(`<frame width="320" height="180"><box title="API" /></frame>`)

func TestRenderFormats(t *testing.T) {
	uc := newUsecase()
	excal, err := newUsecase().BuildScene(context.Background(), apiTestXAL, entity.RenderOptions{Theme: "dark"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(excal, []byte(`"viewBackgroundColor": "#111827"`)) {
		t.Fatalf("dark internal scene theme missing")
	}
	svg, err := uc.Render(context.Background(), apiTestXAL, entity.RenderOptions{Format: usecase.FormatSVG})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(svg, []byte(`<svg`)) {
		t.Fatalf("SVG output missing root element")
	}
}

func TestValidateUseCase(t *testing.T) {
	if err := usecase.Validate(context.Background(), apiTestXAL); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := usecase.Validate(ctx, apiTestXAL); !errors.Is(err, context.Canceled) {
		t.Fatalf("cancelled Validate error = %v", err)
	}
}

func TestDiagnoseReturnsSourcePosition(t *testing.T) {
	source := []byte("<frame>\n  <connection src=\"bad\" dst=\"2\" />\n</frame>")
	diagnostics, err := usecase.Diagnose(context.Background(), source)
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 1 || diagnostics[0].Line != 2 || diagnostics[0].Column != 3 {
		t.Fatalf("diagnostics = %#v", diagnostics)
	}
	if err := usecase.Validate(context.Background(), source); err == nil || !strings.Contains(err.Error(), "line 2, column 3") {
		t.Fatalf("Validate error = %v", err)
	}
}

func TestRenderWithVirtualAssetSource(t *testing.T) {
	assets := &entity.AssetSource{
		FS: fstest.MapFS{}, CatalogCSV: "catalog.csv", GroupIconsDir: "groups", ItemIconSize: 48,
	}
	out, err := newUsecase().Render(context.Background(), apiTestXAL, entity.RenderOptions{Format: usecase.FormatSVG, Assets: assets})
	if err != nil || !bytes.Contains(out, []byte(`<svg`)) {
		t.Fatalf("virtual asset render = %s, err = %v", out, err)
	}
}

func TestBuildPPTXPlanUseCase(t *testing.T) {
	out, err := newUsecase().BuildPPTXPlan(context.Background(), apiTestXAL, entity.RenderOptions{Theme: "dark", PaperSize: "A4"})
	if err != nil || !bytes.Contains(out, []byte(`"slide"`)) || !bytes.Contains(out, []byte(`"ops"`)) {
		t.Fatalf("PPTX plan = %s, err = %v", out, err)
	}
}
