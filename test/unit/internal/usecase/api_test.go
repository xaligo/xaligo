package usecase_test

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

const simpleXAL = `<frame version="1" width="240" height="120"><blank /></frame>`

const multiFrameXAL = `<xaligo version="1">
  <data></data>
  <frames gap="32">
    <frame id="overview" width="240" height="120"><rectangle id="source" title="Overview" width="100" height="60" /></frame>
    <frame id="detail" width="300" height="160"><rectangle id="destination" title="Detail" width="120" height="70" /></frame>
  </frames>
</xaligo>`

func TestRenderArtifactsUsesOneSVGPerFrame(t *testing.T) {
	uc := newUsecase()
	artifacts, err := uc.RenderArtifacts(context.Background(), []byte(multiFrameXAL), entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	if len(artifacts) != 2 || artifacts[0].ID != "overview" || artifacts[1].ID != "detail" {
		t.Fatalf("artifacts = %#v", artifacts)
	}
	if !bytes.Contains(artifacts[0].Data, []byte("Overview")) || bytes.Contains(artifacts[0].Data, []byte(">Detail<")) {
		t.Fatalf("overview SVG contains wrong frame content: %s", artifacts[0].Data)
	}
	if bytes.Contains(artifacts[0].Data, []byte(`width="240" height="120" fill="none"`)) {
		t.Fatalf("overview SVG must not draw its page-frame outline: %s", artifacts[0].Data)
	}
	if _, err := uc.RenderSVG(context.Background(), []byte(multiFrameXAL), entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light"}); err == nil || !strings.Contains(err.Error(), "RenderArtifacts") {
		t.Fatalf("RenderSVG multi-frame error = %v", err)
	}
	combined, err := uc.RenderArtifacts(context.Background(), []byte(multiFrameXAL), entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light", CombineFrames: true})
	if err != nil || len(combined) != 1 || !bytes.Contains(combined[0].Data, []byte("Overview")) || !bytes.Contains(combined[0].Data, []byte("Detail")) {
		t.Fatalf("combined artifacts = %#v err=%v", combined, err)
	}
	if bytes.Contains(combined[0].Data, []byte(`width="240" height="120" fill="none"`)) || bytes.Contains(combined[0].Data, []byte(`width="300" height="160" fill="none"`)) {
		t.Fatalf("combined SVG must not draw page-frame outlines: %s", combined[0].Data)
	}
}

func TestPageDocumentFormatsUseFrameOrder(t *testing.T) {
	uc := newUsecase()
	planJSON, err := uc.BuildPPTXPlan(context.Background(), []byte(multiFrameXAL), entity.RenderOptions{Format: usecase.FormatPPTX, Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var document entity.DocumentPlan
	if err := json.Unmarshal(planJSON, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Pages) != 2 || document.Pages[0].ID != "overview" || document.Pages[1].ID != "detail" {
		t.Fatalf("PPTX document = %#v", document)
	}
	if document.Pages[0].Slide != document.Pages[1].Slide {
		t.Fatalf("PPTX pages must share one slide size: %#v", document.Pages)
	}

	pdf, err := uc.RenderPDF(context.Background(), []byte(multiFrameXAL), entity.RenderOptions{Format: usecase.FormatPDF, Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(pdf, []byte("%PDF-")) || !bytes.Contains(pdf, []byte("/Type/Pages/Count 2")) {
		t.Fatalf("PDF does not contain two pages: %q", pdf)
	}

	workbook, err := uc.RenderExcel(context.Background(), []byte(multiFrameXAL), entity.RenderOptions{Format: usecase.FormatExcel, Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	workbookXML := zipEntryString(t, workbook, "xl/workbook.xml")
	if !strings.Contains(workbookXML, `name="overview"`) || !strings.Contains(workbookXML, `name="detail"`) {
		t.Fatalf("workbook sheets = %s", workbookXML)
	}
	if zipEntryString(t, workbook, "xl/media/image1.svg") == "" || zipEntryString(t, workbook, "xl/media/image2.svg") == "" {
		t.Fatal("workbook frame SVG media is missing")
	}
}

func zipEntryString(t *testing.T, archive []byte, name string) string {
	t.Helper()
	reader, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range reader.File {
		if file.Name != name {
			continue
		}
		entry, err := file.Open()
		if err != nil {
			t.Fatal(err)
		}
		defer entry.Close()
		data, err := io.ReadAll(entry)
		if err != nil {
			t.Fatal(err)
		}
		return string(data)
	}
	t.Fatalf("ZIP entry %s not found", name)
	return ""
}

func TestRenderSVGLoadsInjectedTableImport(t *testing.T) {
	input := []byte(`<xaligo version="1"><data><table-data id="rows" src="rows.csv" /></data><frames><frame id="main" width="500" height="300"><table data="rows" /></frame></frames></xaligo>`)
	output, err := newUsecase().RenderSVG(context.Background(), input, entity.RenderOptions{
		Format: usecase.FormatSVG,
		Imports: &entity.ImportSource{FS: fstest.MapFS{
			"rows.csv": {Data: []byte("name,value\nAPI,8080\n")},
		}},
	})
	if err != nil {
		t.Fatalf("RenderSVG() error = %v", err)
	}
	if !strings.Contains(string(output), "API") || !strings.Contains(string(output), "8080") {
		t.Fatalf("SVG does not contain imported cells: %s", output)
	}
}

type fakePPTXExporter struct {
	seen entity.PptxExportOptions
}

type fixedDocumentSVGRepository struct {
	data []byte
}

func (rcvr *fixedDocumentSVGRepository) Render(entity.Plan, float64, string) ([]byte, error) {
	return bytes.Clone(rcvr.data), nil
}

type capturingSpreadsheetRepository struct {
	pages []entity.RenderPage
}

func (rcvr *capturingSpreadsheetRepository) Export(_ context.Context, pages []entity.RenderPage) ([]byte, error) {
	rcvr.pages = append([]entity.RenderPage(nil), pages...)
	return []byte("spreadsheet"), nil
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
		repository.NewPDFRepository(),
		repository.NewSpreadsheetRepository(),
	)
}

func TestDocumentContainersUseRenderedSVGIntrinsicDimensions(t *testing.T) {
	svgRepository := &fixedDocumentSVGRepository{data: []byte(`<?xml version="1.0"?><svg xmlns="http://www.w3.org/2000/svg" width="640" height="360" viewBox="0 0 640 360"></svg>`)}
	spreadsheetRepository := &capturingSpreadsheetRepository{}
	uc := usecase.NewRenderUsecase(
		repository.NewExcalidrawRepository(),
		repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(),
		repository.NewIsoflowRepository(),
		svgRepository,
		repository.NewXYFlowRepository(),
		nil,
		spreadsheetRepository,
	)
	if _, err := uc.RenderExcel(context.Background(), []byte(simpleXAL), entity.RenderOptions{Format: usecase.FormatExcel, Theme: "light"}); err != nil {
		t.Fatal(err)
	}
	if len(spreadsheetRepository.pages) != 1 {
		t.Fatalf("spreadsheet pages = %#v", spreadsheetRepository.pages)
	}
	page := spreadsheetRepository.pages[0]
	if page.WidthPx != 640 || page.HeightPx != 360 {
		t.Fatalf("render page dimensions = %gx%g, want intrinsic SVG 640x360", page.WidthPx, page.HeightPx)
	}
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
	if diagnostics, err := diagnosticsUsecase.Diagnose(ctx, []byte(simpleXAL)); err != nil || len(diagnostics) != 1 || diagnostics[0].Severity != usecase.SeverityWarning {
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
		{"RenderPDF", usecase.FormatPDF, uc.RenderPDF, `%PDF-`},
		{"RenderExcel", usecase.FormatExcel, uc.RenderExcel, `PK`},
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
	formats := []entity.Format{usecase.FormatSVG, usecase.FormatPDF, usecase.FormatExcel, usecase.FormatXYFlow, usecase.FormatIsoflow}
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
    <connection src="web" dst="detail.db" />
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
	var sourceStub map[string]any
	var destinationStub map[string]any
	for _, element := range scene.Elements {
		switch element["id"] {
		case "paper-frame":
			if element["strokeColor"] != "transparent" {
				t.Fatalf("document frame outline must be transparent: %#v", element)
			}
		case "paper-frame-overview":
			foundOverviewFrame = true
			if element["strokeColor"] != "transparent" {
				t.Fatalf("overview page frame outline must be transparent: %#v", element)
			}
		case "paper-frame-detail":
			foundDetailFrame = true
			if element["strokeColor"] != "transparent" {
				t.Fatalf("detail page frame outline must be transparent: %#v", element)
			}
		}
		if element["type"] == "text" {
			if element["text"] == "to <detail>" {
				foundToLabel = true
			}
			if element["text"] == "from <overview>" {
				foundFromLabel = true
			}
		}
		if custom, _ := element["customData"].(map[string]any); custom["xaligoCrossFrame"] == true {
			foundCrossArrows++
			if element["startBinding"] != nil && element["endBinding"] == nil {
				sourceStub = element
			}
			if element["startBinding"] == nil && element["endBinding"] != nil {
				destinationStub = element
			}
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
	if sourceStub == nil || destinationStub == nil {
		t.Fatalf("cross-frame directions missing: source=%#v destination=%#v", sourceStub, destinationStub)
	}
	overviewFrame := sceneElementRect(t, scene.Elements, "paper-frame-overview")
	detailFrame := sceneElementRect(t, scene.Elements, "paper-frame-detail")
	sourceStart, sourceTerminal := sceneArrowEndpoints(t, sourceStub)
	destinationTerminal, destinationEnd := sceneArrowEndpoints(t, destinationStub)
	assertCrossFrameStubGeometry(t, scene.Elements, sourceStub, sourceStart, sourceTerminal, overviewFrame, "startBinding")
	assertCrossFrameStubGeometry(t, scene.Elements, destinationStub, destinationTerminal, destinationEnd, detailFrame, "endBinding")
	assertPageLinkLabelClearsStub(t, scene.Elements, sourceStub, sourceStart, sourceTerminal)
	assertPageLinkLabelClearsStub(t, scene.Elements, destinationStub, destinationTerminal, destinationEnd)
	if side := sceneFrameSideAtPoint(overviewFrame, sourceTerminal); side != "top" {
		t.Fatalf("equal-distance source tie selected %q, want stable top edge", side)
	}
	if side := sceneFrameSideAtPoint(detailFrame, destinationTerminal); side != "left" {
		t.Fatalf("equal-distance destination tie selected %q, want remote-facing left edge", side)
	}

	svgOut, err := newUsecase().RenderSVG(context.Background(), input, entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light", CombineFrames: true})
	if err != nil {
		t.Fatal(err)
	}
	svg := string(svgOut)
	for _, label := range []string{"to &lt;detail&gt;", "from &lt;overview&gt;"} {
		if !strings.Contains(svg, label) {
			t.Fatalf("SVG cross-frame label %q missing:\n%s", label, svg)
		}
	}
	if count := strings.Count(svg, `fill="none" stroke="#1E1E1E"`); count < 2 {
		t.Fatalf("SVG visible cross-frame paths = %d, want at least two:\n%s", count, svg)
	}
}

func TestRenderExcalidrawCrossFrameAutomaticSideUsesEveryNearestEdge(t *testing.T) {
	input := []byte(`<xaligo version="1"><data></data><frames gap="48">
  <frame id="top-page" width="200" height="200" content-width="40" content-height="40" align="top-center">
    <rectangle id="node" title="Top" width="40" height="40" />
    <connection src="node" dst="sink.target"><bend x="200" y="100" /></connection>
  </frame>
  <frame id="right-page" width="200" height="200" content-width="40" content-height="40" align="middle-right">
    <rectangle id="node" title="Right" width="40" height="40" />
    <connection src="node" dst="sink.target" />
  </frame>
  <frame id="bottom-page" width="200" height="200" content-width="40" content-height="40" align="bottom-center">
    <rectangle id="node" title="Bottom" width="40" height="40" />
    <connection src="node" dst="sink.target" />
  </frame>
  <frame id="left-page" width="200" height="200" content-width="40" content-height="40" align="middle-left">
    <rectangle id="node" title="Left" width="40" height="40" />
    <connection src="node" dst="sink.target" />
  </frame>
	<frame id="corner-page" width="100" height="100" class="pl-2 pt-1">
	  <rectangle id="node" title="Corner" width="10" height="10" />
	  <connection src="node" dst="sink.target" />
	</frame>
  <frame id="sink" width="200" height="200">
    <rectangle id="target" title="Target" width="80" height="80" />
  </frame>
</frames></xaligo>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	wantSides := map[string]string{
		"top-page":    "top",
		"right-page":  "right",
		"bottom-page": "bottom",
		"left-page":   "left",
		"corner-page": "top",
	}
	found := map[string]bool{}
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		frameID, _ := custom["xaligoSourceFrame"].(string)
		wantSide, expected := wantSides[frameID]
		if !expected || element["startBinding"] == nil || element["endBinding"] != nil {
			continue
		}
		start, terminal := sceneArrowEndpoints(t, element)
		frame := sceneElementRect(t, scene.Elements, "paper-frame-"+frameID)
		assertCrossFrameStubGeometry(t, scene.Elements, element, start, terminal, frame, "startBinding")
		if side := sceneFrameSideAtPoint(frame, terminal); side != wantSide {
			t.Fatalf("source frame %q page-link side = %q, want %q", frameID, side, wantSide)
		}
		if frameID == "corner-page" && len(sceneArrowPoints(t, element)) != 4 {
			t.Fatalf("corner page-link points = %#v, want an orthogonal border dogleg", element["points"])
		}
		found[frameID] = true
	}
	for frameID := range wantSides {
		if !found[frameID] {
			t.Fatalf("source page-link for frame %q missing: %#v", frameID, scene.Elements)
		}
	}
}

func TestRenderExcalidrawCrossFrameSmallPagesKeepStubsVisible(t *testing.T) {
	input := []byte(`<xaligo version="1"><data></data><frames gap="16">
  <frame id="source-page-with-long-id" width="48.000000001" height="48.000000001">
    <rectangle id="node" title="A" width="48.000000001" height="48.000000001" />
    <connection src="node" dst="destination-page-with-long-id.node" />
  </frame>
  <frame id="destination-page-with-long-id" width="49" height="49">
    <rectangle id="node" title="B" width="49" height="49" />
  </frame>
</frames></xaligo>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	stubCount := 0
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] != true {
			continue
		}
		stubCount++
		start, end := sceneArrowEndpoints(t, element)
		frameID := custom["xaligoSourceFrame"].(string)
		bindingName := "startBinding"
		if element["startBinding"] == nil {
			frameID = custom["xaligoDestinationFrame"].(string)
			bindingName = "endBinding"
		}
		assertCrossFrameStubGeometry(t, scene.Elements, element, start, end, sceneElementRect(t, scene.Elements, "paper-frame-"+frameID), bindingName)
	}
	if stubCount != 2 {
		t.Fatalf("small-page cross-frame stubs = %d, want 2: %#v", stubCount, scene.Elements)
	}
	wantLabelFrames := map[string]string{
		"to <destination-page-with-long-id>": "source-page-with-long-id",
		"from <source-page-with-long-id>":    "destination-page-with-long-id",
	}
	for label, frameID := range wantLabelFrames {
		var labelRect [4]float64
		for _, element := range scene.Elements {
			if element["text"] == label {
				labelRect = sceneElementRect(t, scene.Elements, element["id"].(string))
				break
			}
		}
		frame := sceneElementRect(t, scene.Elements, "paper-frame-"+frameID)
		if labelRect[2] <= 0 || labelRect[0] < frame[0]-1e-9 || labelRect[1] < frame[1]-1e-9 || labelRect[0]+labelRect[2] > frame[0]+frame[2]+1e-9 || labelRect[1]+labelRect[3] > frame[1]+frame[3]+1e-9 {
			t.Fatalf("small-page label %q rect %#v escapes frame %#v", label, labelRect, frame)
		}
	}
	svgOut, err := newUsecase().RenderSVG(context.Background(), input, entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light", CombineFrames: true})
	if err != nil {
		t.Fatal(err)
	}
	if count := strings.Count(string(svgOut), `fill="none" stroke="#1E1E1E"`); count < 2 {
		t.Fatalf("small-page SVG visible cross-frame paths = %d, want at least 2:\n%s", count, svgOut)
	}
}

func TestRenderExcalidrawCrossFrameExplicitSidesOverrideNearestEdge(t *testing.T) {
	input := []byte(`<frames gap="48">
  <frame id="overview" width="320" height="180">
    <rectangle id="web" title="Web" width="120" height="80" />
    <connection src="web" dst="detail.db" src-side="bottom" dst-anchor="right-2" />
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
	var sourceStub, destinationStub map[string]any
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] != true {
			continue
		}
		if element["startBinding"] != nil && element["endBinding"] == nil {
			sourceStub = element
		}
		if element["startBinding"] == nil && element["endBinding"] != nil {
			destinationStub = element
		}
	}
	if sourceStub == nil || destinationStub == nil {
		t.Fatalf("cross-frame stubs missing: %#v", scene.Elements)
	}
	overviewFrame := sceneElementRect(t, scene.Elements, "paper-frame-overview")
	detailFrame := sceneElementRect(t, scene.Elements, "paper-frame-detail")
	_, sourceTerminal := sceneArrowEndpoints(t, sourceStub)
	destinationTerminal, _ := sceneArrowEndpoints(t, destinationStub)
	if math.Abs(sourceTerminal[1]-(overviewFrame[1]+overviewFrame[3])) > 1e-9 {
		t.Fatalf("source terminal = %#v, want explicit bottom frame edge %#v", sourceTerminal, overviewFrame)
	}
	if math.Abs(destinationTerminal[0]-(detailFrame[0]+detailFrame[2])) > 1e-9 {
		t.Fatalf("destination terminal = %#v, want explicit right frame edge %#v", destinationTerminal, detailFrame)
	}
	assertSceneBindingFixedPoint(t, sourceStub, "startBinding", [2]float64{0.5, 1})
	assertSceneBindingFixedPoint(t, destinationStub, "endBinding", [2]float64{1, 0.3})
}

func TestRenderExcalidrawCrossFrameBoundaryAnchorsAndNearbyLabels(t *testing.T) {
	input := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="overview" width="320" height="180">
    <rectangle id="web" title="Web" width="120" height="80" />
    <connection src="web" dst="detail.db"
                src-anchor="right-2" src-frame-anchor="bottom-4"
                dst-anchor="left-4" dst-frame-anchor="top-2" />
  </frame>
  <frame id="detail" width="320" height="180">
    <rectangle id="db" title="DB" width="120" height="80" />
  </frame>
</frames></xaligo>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var sourceStub, destinationStub map[string]any
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] != true {
			continue
		}
		if element["startBinding"] != nil {
			sourceStub = element
		} else {
			destinationStub = element
		}
	}
	if sourceStub == nil || destinationStub == nil {
		t.Fatalf("cross-frame stubs missing: %#v", scene.Elements)
	}
	overviewFrame := sceneElementRect(t, scene.Elements, "paper-frame-overview")
	detailFrame := sceneElementRect(t, scene.Elements, "paper-frame-detail")
	sourceStart, sourceTerminal := sceneArrowEndpoints(t, sourceStub)
	destinationTerminal, destinationEnd := sceneArrowEndpoints(t, destinationStub)
	wantSourceTerminal := [2]float64{overviewFrame[0] + overviewFrame[2]*0.7, overviewFrame[1] + overviewFrame[3]}
	wantDestinationTerminal := [2]float64{detailFrame[0] + detailFrame[2]*0.3, detailFrame[1]}
	if math.Abs(sourceTerminal[0]-wantSourceTerminal[0]) > 1e-9 || math.Abs(sourceTerminal[1]-wantSourceTerminal[1]) > 1e-9 {
		t.Fatalf("source frame terminal = %#v, want %#v", sourceTerminal, wantSourceTerminal)
	}
	if math.Abs(destinationTerminal[0]-wantDestinationTerminal[0]) > 1e-9 || math.Abs(destinationTerminal[1]-wantDestinationTerminal[1]) > 1e-9 {
		t.Fatalf("destination frame terminal = %#v, want %#v", destinationTerminal, wantDestinationTerminal)
	}
	assertSceneBindingFixedPoint(t, sourceStub, "startBinding", [2]float64{1, 0.3})
	assertSceneBindingFixedPoint(t, destinationStub, "endBinding", [2]float64{0, 0.7})
	assertCrossFrameEndpointAndTerminalApproaches(t, sourceStub, "right", "bottom", false)
	assertCrossFrameEndpointAndTerminalApproaches(t, destinationStub, "left", "top", true)
	assertPageLinkLabelClearsStub(t, scene.Elements, sourceStub, sourceStart, sourceTerminal)
	assertPageLinkLabelClearsStub(t, scene.Elements, destinationStub, destinationTerminal, destinationEnd)
	assertPageLinkLabelGap(t, sceneTextRectByValue(t, scene.Elements, "to <detail>"), sourceTerminal, "bottom", 4)
	assertPageLinkLabelGap(t, sceneTextRectByValue(t, scene.Elements, "from <overview>"), destinationTerminal, "top", 4)

	artifacts, err := newUsecase().RenderArtifacts(context.Background(), input, entity.RenderOptions{Format: usecase.FormatSVG, Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	if len(artifacts) != 2 {
		t.Fatalf("SVG artifacts = %#v", artifacts)
	}
	for _, artifact := range artifacts {
		svg := string(artifact.Data)
		if !strings.Contains(svg, `width="320" height="180" viewBox="0 0 320 180"`) || !strings.Contains(svg, `transform="translate(0 0)" clip-path="url(#xaligo-slide-clip)"`) {
			t.Fatalf("frame %q is not cropped to its logical page:\n%s", artifact.ID, svg)
		}
	}
}

func TestRenderExcalidrawCrossFrameSideDoesNotOverrideAutomaticItemSide(t *testing.T) {
	input := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="overview" width="200" height="200" content-width="40" content-height="40" align="top-center">
    <rectangle id="web" title="Web" width="40" height="40" />
    <connection src="web" dst="detail.db" src-frame-side="bottom" />
  </frame>
  <frame id="detail" width="200" height="200"><rectangle id="db" width="80" height="80" /></frame>
</frames></xaligo>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var sourceStub map[string]any
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] == true && element["startBinding"] != nil {
			sourceStub = element
			break
		}
	}
	if sourceStub == nil {
		t.Fatalf("source page-link stub missing: %#v", scene.Elements)
	}
	frame := sceneElementRect(t, scene.Elements, "paper-frame-overview")
	_, terminal := sceneArrowEndpoints(t, sourceStub)
	if side := sceneFrameSideAtPoint(frame, terminal); side != "bottom" {
		t.Fatalf("source frame terminal side = %q, want bottom", side)
	}
	assertSceneBindingFixedPoint(t, sourceStub, "startBinding", [2]float64{0.5, 0})
	assertCrossFrameEndpointAndTerminalApproaches(t, sourceStub, "top", "bottom", false)
}

func TestRenderExcalidrawExactCoincidentFrameAnchorKeepsVisibleStub(t *testing.T) {
	input := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="overview" width="100" height="100">
    <rectangle id="web" title="Web" width="100" height="100" />
    <connection src="web" dst="detail.db" src-anchor="top-3" src-frame-anchor="top-3" />
  </frame>
  <frame id="detail" width="100" height="100"><rectangle id="db" width="80" height="80" /></frame>
</frames></xaligo>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var sourceStub map[string]any
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] == true && element["startBinding"] != nil {
			sourceStub = element
			break
		}
	}
	if sourceStub == nil {
		t.Fatalf("source page-link stub missing: %#v", scene.Elements)
	}
	start, terminal := sceneArrowEndpoints(t, sourceStub)
	if math.Abs(start[0]-terminal[0]) > 1e-9 || math.Abs(start[1]-terminal[1]) > 1e-9 {
		t.Fatalf("explicit coincident terminal moved: start=%#v terminal=%#v", start, terminal)
	}
	points := sceneArrowPoints(t, sourceStub)
	if len(points) < 3 || sceneNumber(t, sourceStub["width"]) <= 0 && sceneNumber(t, sourceStub["height"]) <= 0 {
		t.Fatalf("coincident frame anchor produced invisible stub: %#v", sourceStub)
	}
	assertCrossFrameEndpointAndTerminalApproaches(t, sourceStub, "top", "top", false)
}

func assertCrossFrameEndpointAndTerminalApproaches(t *testing.T, stub map[string]any, endpointSide, frameSide string, frameAtStart bool) {
	t.Helper()
	points := sceneArrowPoints(t, stub)
	if len(points) < 2 {
		t.Fatalf("cross-frame stub %q points = %#v", stub["id"], points)
	}
	first, second := points[0], points[1]
	penultimate, last := points[len(points)-2], points[len(points)-1]
	if frameAtStart {
		assertSegmentPerpendicularToSide(t, first, second, frameSide)
		assertSegmentPerpendicularToSide(t, penultimate, last, endpointSide)
		return
	}
	assertSegmentPerpendicularToSide(t, first, second, endpointSide)
	assertSegmentPerpendicularToSide(t, penultimate, last, frameSide)
}

func assertSegmentPerpendicularToSide(t *testing.T, start, end [2]float64, side string) {
	t.Helper()
	dx := math.Abs(end[0] - start[0])
	dy := math.Abs(end[1] - start[1])
	if side == "left" || side == "right" {
		if dx <= 1e-9 || dy > 1e-9 {
			t.Fatalf("segment %#v -> %#v is not perpendicular to %s", start, end, side)
		}
		return
	}
	if dy <= 1e-9 || dx > 1e-9 {
		t.Fatalf("segment %#v -> %#v is not perpendicular to %s", start, end, side)
	}
}

func sceneTextRectByValue(t *testing.T, elements []map[string]any, value string) [4]float64 {
	t.Helper()
	for _, element := range elements {
		if element["type"] == "text" && element["text"] == value {
			return [4]float64{sceneNumber(t, element["x"]), sceneNumber(t, element["y"]), sceneNumber(t, element["width"]), sceneNumber(t, element["height"])}
		}
	}
	t.Fatalf("scene text %q not found", value)
	return [4]float64{}
}

func assertPageLinkLabelGap(t *testing.T, label [4]float64, terminal [2]float64, side string, want float64) {
	t.Helper()
	var normalGap, tangentGap float64
	switch side {
	case "top":
		normalGap = label[1] - terminal[1]
		tangentGap = math.Min(math.Abs(label[0]-terminal[0]), math.Abs(terminal[0]-(label[0]+label[2])))
	case "bottom":
		normalGap = terminal[1] - (label[1] + label[3])
		tangentGap = math.Min(math.Abs(label[0]-terminal[0]), math.Abs(terminal[0]-(label[0]+label[2])))
	case "left":
		normalGap = label[0] - terminal[0]
		tangentGap = math.Min(math.Abs(label[1]-terminal[1]), math.Abs(terminal[1]-(label[1]+label[3])))
	case "right":
		normalGap = terminal[0] - (label[0] + label[2])
		tangentGap = math.Min(math.Abs(label[1]-terminal[1]), math.Abs(terminal[1]-(label[1]+label[3])))
	}
	if math.Abs(normalGap-want) > 1e-9 || tangentGap < want-1e-9 {
		t.Fatalf("page-link label %#v gaps normal=%.3f tangent=%.3f, want normal %.3f and tangent >= %.3f from terminal %#v on %s", label, normalGap, tangentGap, want, want, terminal, side)
	}
}

func sceneElementRect(t *testing.T, elements []map[string]any, id string) [4]float64 {
	t.Helper()
	for _, element := range elements {
		if element["id"] == id {
			return [4]float64{sceneNumber(t, element["x"]), sceneNumber(t, element["y"]), sceneNumber(t, element["width"]), sceneNumber(t, element["height"])}
		}
	}
	t.Fatalf("scene element %q not found", id)
	return [4]float64{}
}

func sceneArrowEndpoints(t *testing.T, arrow map[string]any) ([2]float64, [2]float64) {
	t.Helper()
	points := sceneArrowPoints(t, arrow)
	return points[0], points[len(points)-1]
}

func sceneArrowPoints(t *testing.T, arrow map[string]any) [][2]float64 {
	t.Helper()
	origin := [2]float64{sceneNumber(t, arrow["x"]), sceneNumber(t, arrow["y"])}
	points, ok := arrow["points"].([]any)
	if !ok || len(points) < 2 {
		t.Fatalf("arrow points = %#v", arrow["points"])
	}
	absolute := make([][2]float64, 0, len(points))
	for _, point := range points {
		coordinates, ok := point.([]any)
		if !ok || len(coordinates) < 2 {
			t.Fatalf("arrow point = %#v", point)
		}
		absolute = append(absolute, [2]float64{origin[0] + sceneNumber(t, coordinates[0]), origin[1] + sceneNumber(t, coordinates[1])})
	}
	return absolute
}

func assertCrossFrameStubGeometry(t *testing.T, elements []map[string]any, stub map[string]any, start, end [2]float64, frame [4]float64, bindingName string) {
	t.Helper()
	points := sceneArrowPoints(t, stub)
	totalLength := 0.0
	for index := 1; index < len(points); index++ {
		dx := math.Abs(points[index][0] - points[index-1][0])
		dy := math.Abs(points[index][1] - points[index-1][1])
		totalLength += dx + dy
		if dx > 1e-9 && dy > 1e-9 {
			t.Fatalf("cross-frame stub %q segment is diagonal: %#v", stub["id"], points)
		}
		if index >= 2 && math.Abs(points[index][0]-points[index-2][0]) <= 1e-9 && math.Abs(points[index][1]-points[index-2][1]) <= 1e-9 {
			t.Fatalf("cross-frame stub %q backtracks over the previous segment: %#v", stub["id"], points)
		}
	}
	if totalLength <= 1e-9 {
		t.Fatalf("cross-frame stub %q has zero length: start=%#v end=%#v", stub["id"], start, end)
	}
	terminal := end
	if bindingName == "endBinding" {
		terminal = start
	}
	side := sceneFrameSideAtPoint(frame, terminal)
	if side == "" {
		t.Fatalf("cross-frame terminal %#v is not on physical frame edge %#v", terminal, frame)
	}
	if len(points) > 2 {
		firstDX := math.Abs(points[1][0] - points[0][0])
		firstDY := math.Abs(points[1][1] - points[0][1])
		lastDX := math.Abs(points[len(points)-1][0] - points[len(points)-2][0])
		lastDY := math.Abs(points[len(points)-1][1] - points[len(points)-2][1])
		if side == "left" || side == "right" {
			if firstDX <= 1e-9 || firstDY > 1e-9 || lastDX <= 1e-9 || lastDY > 1e-9 {
				t.Fatalf("cross-frame stub %q must enter/leave a vertical border perpendicularly: %#v", stub["id"], points)
			}
		} else if firstDY <= 1e-9 || firstDX > 1e-9 || lastDY <= 1e-9 || lastDX > 1e-9 {
			t.Fatalf("cross-frame stub %q must enter/leave a horizontal border perpendicularly: %#v", stub["id"], points)
		}
	}
	binding, _ := stub[bindingName].(map[string]any)
	elementID, _ := binding["elementId"].(string)
	endpoint := sceneElementRect(t, elements, elementID)
	distances := map[string]float64{
		"top":    math.Max(0, endpoint[1]-frame[1]),
		"right":  math.Max(0, frame[0]+frame[2]-(endpoint[0]+endpoint[2])),
		"bottom": math.Max(0, frame[1]+frame[3]-(endpoint[1]+endpoint[3])),
		"left":   math.Max(0, endpoint[0]-frame[0]),
	}
	minimum := math.Min(math.Min(distances["top"], distances["right"]), math.Min(distances["bottom"], distances["left"]))
	if distances[side] > minimum+1e-9 {
		t.Fatalf("cross-frame stub %q uses %s edge at distance %.2f, want nearest distance %.2f: endpoint=%#v frame=%#v", stub["id"], side, distances[side], minimum, endpoint, frame)
	}
}

func assertPageLinkLabelClearsStub(t *testing.T, elements []map[string]any, stub map[string]any, start, end [2]float64) {
	t.Helper()
	id, _ := stub["id"].(string)
	label := sceneElementRect(t, elements, id+"-label")
	bindingName := "startBinding"
	if stub[bindingName] == nil {
		bindingName = "endBinding"
	}
	binding, _ := stub[bindingName].(map[string]any)
	elementID, _ := binding["elementId"].(string)
	endpoint := sceneElementRect(t, elements, elementID)
	if label[0] < endpoint[0]+endpoint[2] && endpoint[0] < label[0]+label[2] && label[1] < endpoint[1]+endpoint[3] && endpoint[1] < label[1]+label[3] {
		t.Fatalf("page-link label %#v overlaps endpoint %#v for stub %q", label, endpoint, id)
	}
	points := sceneArrowPoints(t, stub)
	const epsilon = 1e-9
	for index := 1; index < len(points); index++ {
		segmentStart, segmentEnd := points[index-1], points[index]
		if math.Abs(segmentEnd[1]-segmentStart[1]) <= epsilon {
			minimumX, maximumX := math.Min(segmentStart[0], segmentEnd[0]), math.Max(segmentStart[0], segmentEnd[0])
			if segmentStart[1] >= label[1]-epsilon && segmentStart[1] <= label[1]+label[3]+epsilon && maximumX >= label[0]-epsilon && minimumX <= label[0]+label[2]+epsilon {
				t.Fatalf("page-link label %#v overlaps horizontal stub %q from %#v to %#v", label, id, start, end)
			}
			continue
		}
		minimumY, maximumY := math.Min(segmentStart[1], segmentEnd[1]), math.Max(segmentStart[1], segmentEnd[1])
		if segmentStart[0] >= label[0]-epsilon && segmentStart[0] <= label[0]+label[2]+epsilon && maximumY >= label[1]-epsilon && minimumY <= label[1]+label[3]+epsilon {
			t.Fatalf("page-link label %#v overlaps vertical stub %q from %#v to %#v", label, id, start, end)
		}
	}
}

func sceneFrameSideAtPoint(frame [4]float64, point [2]float64) string {
	const epsilon = 1e-9
	switch {
	case math.Abs(point[1]-frame[1]) <= epsilon:
		return "top"
	case math.Abs(point[0]-(frame[0]+frame[2])) <= epsilon:
		return "right"
	case math.Abs(point[1]-(frame[1]+frame[3])) <= epsilon:
		return "bottom"
	case math.Abs(point[0]-frame[0]) <= epsilon:
		return "left"
	default:
		return ""
	}
}

func assertSceneBindingFixedPoint(t *testing.T, arrow map[string]any, bindingName string, want [2]float64) {
	t.Helper()
	binding, _ := arrow[bindingName].(map[string]any)
	fixedPoint, ok := binding["fixedPoint"].([]any)
	if !ok || len(fixedPoint) != 2 || math.Abs(sceneNumber(t, fixedPoint[0])-want[0]) > 1e-9 || math.Abs(sceneNumber(t, fixedPoint[1])-want[1]) > 1e-9 {
		t.Fatalf("%s fixedPoint = %#v, want %#v", bindingName, binding["fixedPoint"], want)
	}
}

func sceneNumber(t *testing.T, value any) float64 {
	t.Helper()
	number, ok := value.(float64)
	if !ok {
		t.Fatalf("scene number = %#v (%T)", value, value)
	}
	return number
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
	if !strings.Contains(svg, `<path d="M 171 71 L 185 71"`) {
		t.Fatalf("SVG explicit center anchor path missing:\n%s", svg)
	}
	if !strings.Contains(svg, `<path d="M 171 79 L 195 79 L 195 71 L 249 71"`) {
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
