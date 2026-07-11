package usecase_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

const simpleXAL = `<frame width="240" height="120"><blank /></frame>`

type fakePPTXExporter struct {
	seen entity.PptxExportOptions
}

func newUsecase() usecase.RenderUsecase {
	return newUsecaseWithPPTX(repository.NewPowerpointRepository())
}

func newUsecaseWithPPTX(powerpointRepository repository.PowerpointRepository) usecase.RenderUsecase {
	return usecase.NewRenderUsecase(
		repository.NewExcalidrawRepository(),
		repository.NewXaligoRepository(),
		powerpointRepository,
		repository.NewIsoflowRepository(),
		repository.NewSVGRepository(),
		repository.NewXYFlowRepository(),
	)
}

func newSceneDependencies() usecase.SceneDependencies {
	return usecase.SceneDependencies{
		XaligoRepository:     repository.NewXaligoRepository(),
		ExcalidrawRepository: repository.NewExcalidrawRepository(),
	}
}

func (rcvr *fakePPTXExporter) WritePptx(_ context.Context, _ entity.PptxExportOptions) error {
	return nil
}

func (rcvr *fakePPTXExporter) ExportPptxBytes(_ context.Context, opts entity.PptxExportOptions) ([]byte, error) {
	rcvr.seen = opts
	return []byte("pptx-from-fake"), nil
}

func TestUseCaseAPIRendersStableFormats(t *testing.T) {
	uc := newUsecase()
	diagnosticsUsecase := usecase.NewDiagnosticsUsecase()
	ctx := context.Background()
	if err := uc.ValidateRenderOptions(entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light"}); err != nil {
		t.Fatal(err)
	}
	if err := diagnosticsUsecase.Validate(ctx, []byte(simpleXAL)); err != nil {
		t.Fatal(err)
	}
	if diagnostics, err := diagnosticsUsecase.Diagnose(ctx, []byte(simpleXAL)); err != nil || len(diagnostics) != 0 {
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
	uc := newUsecase()
	ctx := context.Background()
	formats := []entity.Format{usecase.FormatSVG, usecase.FormatXYFlow, usecase.FormatIsoflow}
	for _, format := range formats {
		t.Run(string(format), func(t *testing.T) {
			out, err := uc.Render(ctx, []byte(simpleXAL), entity.RenderOptions{Format: format, Theme: "light"})
			if err != nil {
				t.Fatal(err)
			}
			if len(out) == 0 {
				t.Fatal("Render output is empty")
			}
		})
	}
	if _, err := uc.Render(ctx, []byte(simpleXAL), entity.RenderOptions{Format: "unknown", Theme: "light"}); err == nil || !strings.Contains(err.Error(), "unknown render format") {
		t.Fatalf("unknown format err = %v", err)
	}
	canceled, cancel := context.WithCancel(ctx)
	cancel()
	if _, err := uc.Render(canceled, []byte(simpleXAL), entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light"}); err == nil {
		t.Fatal("canceled Render error = nil")
	}
}

func TestUseCaseRenderPPTXExportErrorAfterPlanBuild(t *testing.T) {
	badWASM := filepath.Join(t.TempDir(), "bad.wasm")
	if err := os.WriteFile(badWASM, []byte("not wasm"), 0644); err != nil {
		t.Fatal(err)
	}
	_, err := newUsecase().RenderPPTX(context.Background(), []byte(simpleXAL), entity.RenderOptions{Format: usecase.FormatPPTX, Theme: "light", PPTXExporterWASM: badWASM})
	if err == nil || !strings.Contains(err.Error(), "run PPTX WASM exporter") {
		t.Fatalf("RenderPPTX err = %v", err)
	}
}

func TestUseCaseRenderPPTXUsesInjectedExporter(t *testing.T) {
	exporter := &fakePPTXExporter{}
	uc := newUsecaseWithPPTX(exporter)
	compression := false
	opts := entity.RenderOptions{
		Format:           usecase.FormatPPTX,
		Theme:            "light",
		Title:            "Injected",
		Compression:      &compression,
		PPTXExporterWASM: "custom.wasm",
	}
	out, err := uc.RenderPPTX(context.Background(), []byte(simpleXAL), opts)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "pptx-from-fake" {
		t.Fatalf("RenderPPTX output = %q", out)
	}
	if exporter.seen.Title != "Injected" || exporter.seen.ExporterWASM != "custom.wasm" || exporter.seen.Compression == nil || *exporter.seen.Compression {
		t.Fatalf("exporter opts = %#v", exporter.seen)
	}
	if !strings.Contains(string(exporter.seen.PlanJSON), `"slide"`) {
		t.Fatalf("exporter plan = %s", exporter.seen.PlanJSON)
	}
	out, err = uc.Render(context.Background(), []byte(simpleXAL), opts)
	if err != nil || string(out) != "pptx-from-fake" {
		t.Fatalf("Render(PPTX) output=%q err=%v", out, err)
	}
}

func TestUseCaseRenderFunctionsReportBuildSceneErrors(t *testing.T) {
	uc := newUsecase()
	badInput := []byte(`<frame><item id="abc" /></frame>`)
	cases := []struct {
		name string
		call func(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	}{
		{"RenderExcalidraw", uc.RenderExcalidraw},
		{"RenderSVG", uc.RenderSVG},
		{"BuildPPTXPlan", uc.BuildPPTXPlan},
		{"RenderPPTX", uc.RenderPPTX},
		{"RenderXYFlow", uc.RenderXYFlow},
		{"RenderIsoflow", uc.RenderIsoflow},
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
	out, err := newUsecase().RenderIsoflow(context.Background(), []byte(simpleXAL), entity.RenderOptions{
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
	input := []byte(`<frame width="600" height="300"><aws-cloud id="cloud" title="AWS"><region id="region" title="Region"><vpc id="vpc" title="VPC" layout="staggered"><availability-zone id="az1" title="AZ 1"><blank /></availability-zone><availability-zone id="az2" title="AZ 2"><blank /></availability-zone><availability-zone id="az3" title="AZ 3"><blank /></availability-zone><availability-zone id="az4" title="AZ 4"><blank /></availability-zone><availability-zone id="az5" title="AZ 5"><blank /></availability-zone></vpc></region></aws-cloud></frame>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
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

func TestRenderExcalidrawFramesAndCrossFrameLabels(t *testing.T) {
	input := []byte(`<frames gap="48">
  <frame id="overview" width="320" height="180">
    <rectangle id="web" title="Web" width="120" height="80" />
    <connection src="web" dst="db" />
  </frame>
  <frame id="detail" width="320" height="180">
    <rectangle id="db" title="DB" width="120" height="80" />
  </frame>
</frames>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	foundOverviewFrame := false
	foundDetailFrame := false
	foundToLabel := false
	foundFromLabel := false
	foundCrossArrows := 0
	logicalID := ""
	for _, element := range scene.Elements {
		switch element["id"] {
		case "paper-frame-overview":
			foundOverviewFrame = true
		case "paper-frame-detail":
			foundDetailFrame = true
		}
		if element["type"] == "text" {
			if element["text"] == "to detail" {
				foundToLabel = true
			}
			if element["text"] == "from overview" {
				foundFromLabel = true
			}
		}
		if custom, _ := element["customData"].(map[string]any); custom["xaligoCrossFrame"] == true {
			foundCrossArrows++
			gotLogicalID, _ := custom["xaligoConnectorLogicalId"].(string)
			sourceElementID, _ := custom["xaligoConnectorSourceElementId"].(string)
			destinationElementID, _ := custom["xaligoConnectorDestinationElementId"].(string)
			if gotLogicalID == "" || sourceElementID == "" || destinationElementID == "" {
				t.Fatalf("cross-frame logical metadata missing: %#v", custom)
			}
			if logicalID == "" {
				logicalID = gotLogicalID
			} else if logicalID != gotLogicalID {
				t.Fatalf("cross-frame stubs have different logical IDs: %q and %q", logicalID, gotLogicalID)
			}
		}
	}
	if !foundOverviewFrame || !foundDetailFrame || !foundToLabel || !foundFromLabel || foundCrossArrows != 2 {
		t.Fatalf("frames/cross labels missing: overview=%v detail=%v to=%v from=%v arrows=%d elements=%#v", foundOverviewFrame, foundDetailFrame, foundToLabel, foundFromLabel, foundCrossArrows, scene.Elements)
	}
}

func TestUseCaseAPIRenderPPTXHonorsCanceledContext(t *testing.T) {
	uc := newUsecase()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := uc.RenderPPTX(ctx, []byte(simpleXAL), entity.RenderOptions{Format: usecase.FormatPPTX}); err == nil {
		t.Fatal("RenderPPTX canceled context error = nil")
	}
}

func TestUseCaseAPINewPreviewRepository(t *testing.T) {
	uc := newUsecase()
	if _, err := uc.NewPreviewRepository("", entity.PreviewOptions{}); err == nil {
		t.Fatal("NewPreviewRepository empty path error = nil")
	}
	path := filepath.Join(t.TempDir(), "diagram.xal")
	if err := os.WriteFile(path, []byte(simpleXAL), 0644); err != nil {
		t.Fatal(err)
	}
	server, err := uc.NewPreviewRepository(path, entity.PreviewOptions{Render: entity.RenderOptions{Theme: "light"}})
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
	planJSON, err := newUsecase().BuildPPTXPlan(context.Background(), []byte(`<frame width="240" height="120"><item id="27" /></frame>`), entity.RenderOptions{
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

func TestRenderSVGDrawsServiceLegend(t *testing.T) {
	out, err := newUsecase().RenderSVG(context.Background(), []byte(`<frame width="240" height="120"><item id="27" /></frame>`), entity.RenderOptions{
		Theme:             "light",
		SVGLegendPosition: "right",
		ServicesCSV: []byte(strings.Join([]string{
			"27,Amazon EC2,EC2,Virtual server,Application tier,",
		}, "\n")),
	})
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	for _, want := range []string{`id="xaligo-svg-legend"`, `EC2`, `Amazon EC2`} {
		if !strings.Contains(svg, want) {
			t.Fatalf("SVG missing %q:\n%s", want, svg)
		}
	}
}

func TestRenderExcalidrawCarriesSharedPortTextLayout(t *testing.T) {
	out, err := newUsecase().RenderExcalidraw(context.Background(), []byte(`<frame width="240" height="120">
  <rectangle id="service" width="180" height="80">
    <port id="input" side="left" width="80" title="long input port" />
  </rectangle>
</frame>`), entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	for _, element := range scene.Elements {
		if element.ID != "frame-0-0-text" {
			continue
		}
		if element.CustomData == nil || !element.CustomData.PortLabel || element.CustomData.TextLayout == nil {
			t.Fatalf("port custom data = %#v", element.CustomData)
		}
		layout := element.CustomData.TextLayout
		if layout.Role != entity.TextRolePortLabel || !layout.Wrap || layout.Fit != entity.TextFitShrink || layout.Overflow != entity.TextOverflowClip || !layout.Clip {
			t.Fatalf("port text layout = %#v", layout)
		}
		return
	}
	t.Fatalf("port text element not found: %#v", scene.Elements)
}

func TestRenderSVGContainsLongPortText(t *testing.T) {
	out, err := newUsecase().RenderSVG(context.Background(), []byte(`<frame width="240" height="120">
  <rectangle id="service" width="180" height="80">
    <port id="input" side="left" width="48" title="very-long-port-label" />
  </rectangle>
</frame>`), entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	if !strings.Contains(svg, `clipPath id="xaligo-text-clip-`) || !strings.Contains(svg, "<tspan") {
		t.Fatalf("long port text is not wrapped/clipped by the shared contract:\n%s", svg)
	}
}

func TestRenderSVGPreservesExplicitCenterAnchor(t *testing.T) {
	out, err := newUsecase().RenderSVG(context.Background(), []byte(`<frame width="420" height="160" layout="horizontal" gap="80" class="pa-4" item-size="40">
  <item id="27" name="src" />
  <item id="110" name="dst-a" />
  <item id="117" name="dst-b" />
  <connection src="src" dst="dst-a" src-anchor="right-3" dst-anchor="left-3" />
  <connection src="src" dst="dst-b" />
</frame>`), entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	if !strings.Contains(svg, `<path d="M 187 71 L 185 71"`) {
		t.Fatalf("SVG explicit center anchor path missing:\n%s", svg)
	}
	if !strings.Contains(svg, `<path d="M 187 79 L 187 71 L 233 71"`) {
		t.Fatalf("SVG automatic fanout path missing:\n%s", svg)
	}
}

func TestBuildPPTXPlanValidatesServiceLegend(t *testing.T) {
	cases := []struct {
		name     string
		services string
		want     string
	}{
		{"missing catalog ID", "Amazon EC2,EC2", "catalog ID is required"},
		{"missing official name", "27,,EC2", "official name is required"},
		{"duplicate catalog ID", "27,Amazon EC2,EC2\n27,Amazon EC2 Duplicate,EC2B", "duplicate catalog ID"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := newUsecase().BuildPPTXPlan(context.Background(), []byte(`<frame width="240" height="120"><item id="27" /></frame>`), entity.RenderOptions{
				Format:      usecase.FormatPPTX,
				Theme:       "light",
				ServicesCSV: []byte(tc.services),
			})
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v, want %q", err, tc.want)
			}
		})
	}
}
