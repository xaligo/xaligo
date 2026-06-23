package usecase_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	awsassets "github.com/ryo-arima/xaligo/etc/resources/aws"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

const simpleXAL = `<frame width="240" height="120"><blank /></frame>`

func TestUseCaseAPIRendersStableFormats(t *testing.T) {
	uc := usecase.New()
	ctx := context.Background()
	if err := uc.ValidateRenderOptions(entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light"}); err != nil {
		t.Fatal(err)
	}
	if err := uc.Validate(ctx, []byte(simpleXAL)); err != nil {
		t.Fatal(err)
	}
	if diagnostics, err := uc.Diagnose(ctx, []byte(simpleXAL)); err != nil || len(diagnostics) != 0 {
		t.Fatalf("Diagnose() diagnostics=%#v err=%v", diagnostics, err)
	}

	checks := []struct {
		name   string
		format entity.Format
		call   func(context.Context, []byte, entity.RenderOptions) ([]byte, error)
		want   string
	}{
		{"Render default", "", uc.Render, `"type": "excalidraw"`},
		{"RenderExcalidraw", usecase.FormatExcalidraw, uc.RenderExcalidraw, `"type": "excalidraw"`},
		{"RenderSVG", usecase.FormatSVG, uc.RenderSVG, `<svg`},
		{"RenderXYFlow", usecase.FormatXYFlow, uc.RenderXYFlow, `"nodes"`},
		{"RenderIsoflow", usecase.FormatIsoflow, uc.RenderIsoflow, `"version": "3.3.0"`},
		{"BuildPPTXPlan", usecase.FormatPPTX, uc.BuildPPTXPlan, `"slide"`},
	}
	for _, check := range checks {
		t.Run(check.name, func(t *testing.T) {
			out, err := check.call(ctx, []byte(simpleXAL), entity.RenderOptions{Format: check.format, Theme: "light", PxPerInch: 96})
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(out), check.want) {
				t.Fatalf("output %q does not contain %q", out, check.want)
			}
		})
	}
}

func TestUseCaseRenderDispatcherBranches(t *testing.T) {
	ctx := context.Background()
	formats := []entity.Format{usecase.FormatSVG, usecase.FormatXYFlow, usecase.FormatIsoflow}
	for _, format := range formats {
		t.Run(string(format), func(t *testing.T) {
			out, err := usecase.Render(ctx, []byte(simpleXAL), entity.RenderOptions{Format: format, Theme: "light"})
			if err != nil {
				t.Fatal(err)
			}
			if len(out) == 0 {
				t.Fatal("Render output is empty")
			}
		})
	}
	if _, err := usecase.Render(ctx, []byte(simpleXAL), entity.RenderOptions{Format: "unknown", Theme: "light"}); err == nil || !strings.Contains(err.Error(), "unknown render format") {
		t.Fatalf("unknown format err = %v", err)
	}
	canceled, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := usecase.Render(canceled, []byte(simpleXAL), entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light"}); err == nil {
		t.Fatal("canceled Render error = nil")
	}
}

func TestUseCaseRenderPPTXExportErrorAfterPlanBuild(t *testing.T) {
	badWASM := filepath.Join(t.TempDir(), "bad.wasm")
	if err := os.WriteFile(badWASM, []byte("not wasm"), 0644); err != nil {
		t.Fatal(err)
	}
	_, err := usecase.RenderPPTX(context.Background(), []byte(simpleXAL), entity.RenderOptions{Format: usecase.FormatPPTX, Theme: "light", PPTXExporterWASM: badWASM})
	if err == nil || !strings.Contains(err.Error(), "run PPTX WASM exporter") {
		t.Fatalf("RenderPPTX err = %v", err)
	}
}

func TestUseCaseRenderFunctionsReportBuildSceneErrors(t *testing.T) {
	badInput := []byte(`<frame><item id="abc" /></frame>`)
	cases := []struct {
		name string
		call func(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	}{
		{"RenderExcalidraw", usecase.RenderExcalidraw},
		{"RenderSVG", usecase.RenderSVG},
		{"BuildPPTXPlan", usecase.BuildPPTXPlan},
		{"RenderPPTX", usecase.RenderPPTX},
		{"RenderXYFlow", usecase.RenderXYFlow},
		{"RenderIsoflow", usecase.RenderIsoflow},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.call(context.Background(), badInput, entity.RenderOptions{Theme: "light"})
			if err == nil || !strings.Contains(err.Error(), "positive integer") {
				t.Fatalf("err = %v", err)
			}
		})
	}
}

func TestUseCaseRenderIsoflowUsesEmbeddedManifest(t *testing.T) {
	out, err := usecase.RenderIsoflow(context.Background(), []byte(simpleXAL), entity.RenderOptions{
		Format: usecase.FormatIsoflow,
		Theme:  "light",
		Assets: &entity.AssetSource{
			FS:               awsassets.Assets,
			CatalogCSV:       awsassets.CatalogCSV,
			GroupIconsDir:    awsassets.GroupIconsDir,
			IsoflowIconsJSON: awsassets.IsoflowIconsJSON,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `"version": "3.3.0"`) {
		t.Fatalf("isoflow output = %s", out)
	}
}

func TestRenderExcalidrawStaggeredBackgrounds(t *testing.T) {
	input := []byte(`<frame width="600" height="300"><aws-cloud title="AWS"><region title="Region"><vpc title="VPC" layout="staggered"><availability-zone title="AZ 1"><blank /></availability-zone><availability-zone title="AZ 2"><blank /></availability-zone><availability-zone title="AZ 3"><blank /></availability-zone><availability-zone title="AZ 4"><blank /></availability-zone><availability-zone title="AZ 5"><blank /></availability-zone></vpc></region></aws-cloud></frame>`)
	out, err := usecase.RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	text := string(out)
	for _, color := range []string{`#ffffff`, `#c8e8e8`, `#92cecd`} {
		if !strings.Contains(text, color) {
			t.Fatalf("staggered color %s missing from %s", color, text)
		}
	}
}

func TestUseCaseAPIRenderPPTXHonorsCanceledContext(t *testing.T) {
	uc := usecase.New()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := uc.RenderPPTX(ctx, []byte(simpleXAL), entity.RenderOptions{Format: usecase.FormatPPTX}); err == nil {
		t.Fatal("RenderPPTX canceled context error = nil")
	}
}

func TestUseCaseAPINewPreviewServer(t *testing.T) {
	uc := usecase.New()
	if _, err := uc.NewPreviewServer("", entity.PreviewOptions{}); err == nil {
		t.Fatal("NewPreviewServer empty path error = nil")
	}
	path := filepath.Join(t.TempDir(), "diagram.xal")
	if err := os.WriteFile(path, []byte(simpleXAL), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := uc.NewPreviewServer(path, entity.PreviewOptions{Render: entity.RenderOptions{Theme: "light"}})
	if err != nil {
		t.Fatal(err)
	}
	if server.Handler() == nil {
		t.Fatal("preview handler is nil")
	}
	if err := server.Refresh(); err != nil {
		t.Fatal(err)
	}
}

func TestBuildPPTXPlanUsesServiceLegend(t *testing.T) {
	planJSON, err := usecase.BuildPPTXPlan(context.Background(), []byte(`<frame width="240" height="120"><item id="27" /></frame>`), entity.RenderOptions{
		Format: usecase.FormatPPTX,
		Theme:  "light",
		ServicesCSV: []byte(strings.Join([]string{
			"27,Amazon EC2,EC2,Virtual server,Application tier,",
		}, "\n")),
	})
	if err != nil {
		t.Fatal(err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		t.Fatal(err)
	}
	if len(plan.Legend) != 1 || plan.Legend[0].CatalogID != 27 || plan.Legend[0].Abbreviation != "EC2" {
		t.Fatalf("legend = %#v", plan.Legend)
	}
}
