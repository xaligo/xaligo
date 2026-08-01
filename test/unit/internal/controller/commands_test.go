package controller_test

import (
	"context"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/controller"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

type fakeUseCase struct {
	validateErr      error
	validateOptsErr  error
	renderErr        error
	previewErr       error
	renderCalls      int
	lastRenderOpts   entity.RenderOptions
	lastPlanOpts     entity.RenderOptions
	lastPreviewPath  string
	lastPreviewOpts  entity.PreviewOptions
	lastPreviewAddr  string
	renderExcalidraw []byte
	renderSVG        []byte
	renderArtifacts  []entity.RenderArtifact
	planJSON         []byte
}

func (rcvr *fakeUseCase) RenderArtifacts(_ context.Context, _ []byte, opts entity.RenderOptions) ([]entity.RenderArtifact, error) {
	rcvr.renderCalls++
	rcvr.lastRenderOpts = opts
	if rcvr.renderErr != nil {
		return nil, rcvr.renderErr
	}
	if rcvr.renderArtifacts != nil {
		return rcvr.renderArtifacts, nil
	}
	data := rcvr.renderSVG
	if data == nil {
		data = []byte(`<svg></svg>`)
	}
	return []entity.RenderArtifact{{ID: "frame", Data: data}}, nil
}

func (rcvr *fakeUseCase) ValidateRenderOptions(opts entity.RenderOptions) error {
	if rcvr.validateOptsErr != nil {
		return rcvr.validateOptsErr
	}
	return usecase.ValidateRenderOptions(opts)
}

func (rcvr *fakeUseCase) Validate(context.Context, []byte) error { return rcvr.validateErr }

func (rcvr *fakeUseCase) ValidateWithImports(context.Context, []byte, *entity.ImportSource) error {
	return rcvr.validateErr
}

func (rcvr *fakeUseCase) Diagnose(context.Context, []byte) ([]entity.Diagnostic, error) {
	return nil, nil
}

func (rcvr *fakeUseCase) DiagnoseWithImports(context.Context, []byte, *entity.ImportSource) ([]entity.Diagnostic, error) {
	return nil, nil
}

func (rcvr *fakeUseCase) Render(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	rcvr.renderCalls++
	rcvr.lastRenderOpts = opts
	switch opts.Format {
	case usecase.FormatSVG:
		if rcvr.renderSVG != nil {
			return rcvr.renderSVG, rcvr.renderErr
		}
		return []byte(`<svg></svg>`), rcvr.renderErr
	case usecase.FormatPPTX:
		return []byte(`pptx`), rcvr.renderErr
	default:
		return nil, errors.New("unsupported format")
	}
}

func (rcvr *fakeUseCase) BuildScene(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	rcvr.lastRenderOpts = opts
	if rcvr.renderExcalidraw != nil {
		return rcvr.renderExcalidraw, rcvr.renderErr
	}
	return []byte(`{"type":"excalidraw","elements":[],"files":{}}`), rcvr.renderErr
}

func (rcvr *fakeUseCase) RenderSVG(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	rcvr.lastRenderOpts = opts
	if rcvr.renderSVG != nil {
		return rcvr.renderSVG, rcvr.renderErr
	}
	return []byte(`<svg></svg>`), rcvr.renderErr
}

func (rcvr *fakeUseCase) RenderPPTX(context.Context, []byte, entity.RenderOptions) ([]byte, error) {
	return []byte(`pptx`), rcvr.renderErr
}

func (rcvr *fakeUseCase) BuildPPTXPlan(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	rcvr.lastPlanOpts = opts
	if rcvr.planJSON != nil {
		return rcvr.planJSON, rcvr.renderErr
	}
	return []byte(`{"slide":{"w":1,"h":1}}`), rcvr.renderErr
}

func (rcvr *fakeUseCase) NewPreviewRepository(path string, opts entity.PreviewOptions) (repository.PreviewRepository, error) {
	rcvr.lastPreviewPath = path
	rcvr.lastPreviewOpts = opts
	if rcvr.previewErr != nil {
		return nil, rcvr.previewErr
	}
	return fakePreviewRepository{usecase: rcvr}, nil
}

func newGenerateController(_ *fakeUseCase) controller.GenerateController {
	return controller.NewGenerateController()
}

func newRenderController(uc *fakeUseCase) controller.RenderController {
	return controller.NewRenderController(uc)
}

func newServeController(uc *fakeUseCase, port ...int) controller.ServeController {
	servePort := config.DefaultServePort
	if len(port) > 0 {
		servePort = port[0]
	}
	return controller.NewServeController(&config.Config{
		Serve: config.ServeConfig{Port: servePort},
	}, uc)
}

type fakePreviewRepository struct {
	usecase *fakeUseCase
}

func (fakePreviewRepository) Handler() http.Handler { return http.NewServeMux() }
func (rcvr fakePreviewRepository) Run(_ context.Context, address string) error {
	rcvr.usecase.lastPreviewAddr = address
	return nil
}
func (fakePreviewRepository) Refresh() error { return nil }

func writeTempXAL(t *testing.T, dir string) string {
	t.Helper()
	path := filepath.Join(dir, "input.xal")
	if err := os.WriteFile(path, []byte(`<frame width="120" height="80"><blank /></frame>`), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestControllerCommandInitializers(t *testing.T) {
	uc := &fakeUseCase{}
	commands := []*cobra.Command{
		newGenerateController(uc).Command(),
		controller.NewInitController().Command(),
		newServeController(uc).Command(),
		controller.NewValidateController(uc).Command(),
		controller.NewVersionController().Command(),
	}
	for _, cmd := range commands {
		if cmd.Use == "" || cmd.Short == "" {
			t.Fatalf("command missing metadata: %#v", cmd)
		}
	}
}

func TestRunValidateWithUseCase(t *testing.T) {
	input := writeTempXAL(t, t.TempDir())
	if err := controller.NewValidateController(&fakeUseCase{}).Run(input, nil); err != nil {
		t.Fatal(err)
	}
	if err := controller.NewValidateController(&fakeUseCase{validateErr: errors.New("invalid")}).Run(input, nil); err == nil || !strings.Contains(err.Error(), "invalid") {
		t.Fatalf("validation error = %v", err)
	}
	if err := controller.NewValidateController(&fakeUseCase{}).Run(filepath.Join(t.TempDir(), "missing.xal"), nil); err == nil {
		t.Fatal("missing input error = nil")
	}
}

func TestRunValidateResolvesTableImportRelativeToInput(t *testing.T) {
	directory := t.TempDir()
	input := filepath.Join(directory, "diagram.xal")
	if err := os.WriteFile(input, []byte(`<xaligo version="1"><data><table-data id="rows" src="rows.csv" /></data><frames><frame id="main"><table data="rows" /></frame></frames></xaligo>`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directory, "rows.csv"), []byte("name,value\nAPI,8080\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := controller.NewValidateController(usecase.NewDiagnosticsUsecase()).Run(input, nil); err != nil {
		t.Fatalf("validate imported table: %v", err)
	}
}

func TestRunRenderFormatWithUseCaseWritesFormats(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	services := filepath.Join(dir, "services.csv")
	if err := os.WriteFile(services, []byte("27,Amazon EC2,EC2\n"), 0644); err != nil {
		t.Fatal(err)
	}
	formats := []string{"svg", "pptx"}
	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			output := filepath.Join(dir, "out."+format)
			fake := &fakeUseCase{}
			err := newRenderController(fake).RunFormat(entity.ControllerRenderOptions{
				InputPath: input, OutputPath: output, Format: format, ServicesFile: services, Theme: "light", Mode: "standard",
			})
			if err != nil {
				t.Fatal(err)
			}
			data, err := os.ReadFile(output)
			if err != nil {
				t.Fatal(err)
			}
			if len(data) == 0 || fake.renderCalls != 1 || fake.lastRenderOpts.Format != entity.Format(format) {
				t.Fatalf("data=%q opts=%#v", data, fake.lastRenderOpts)
			}
			if !strings.Contains(string(fake.lastRenderOpts.ServicesCSV), "Amazon EC2") {
				t.Fatalf("services CSV was not forwarded: %#v", fake.lastRenderOpts)
			}
		})
	}
}

func TestRenderCommandUsesDefaultOutputs(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	oldWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldWD) })

	formats := map[string]string{
		"svg":  "output.svg",
		"pptx": "output.pptx",
	}
	for format, output := range formats {
		t.Run(format, func(t *testing.T) {
			cmd := newRenderController(&fakeUseCase{}).Command()
			cmd.SetArgs([]string{input, "--format", format, "--no-compression"})
			if err := cmd.Execute(); err != nil {
				t.Fatal(err)
			}
			if _, err := os.Stat(filepath.Join(dir, output)); err != nil {
				t.Fatalf("default output %s was not created: %v", output, err)
			}
		})

		explicit := filepath.Join(dir, "explicit.svg")
		fake := &fakeUseCase{}
		cmd := newRenderController(fake).Command()
		cmd.SetArgs([]string{input, "--format", "svg", "--output", explicit, "--compression", "--combine-frames", "--theme", "dark", "--mode", "network", "--px-per-inch", "120", "--arrow-style", "standard", "--arrow-stub", "22", "--arrow-margin", "11", "--paper", "A4", "--orientation", "landscape", "--paper-margin-left", "0.25", "--svg-legend-position", "left"})
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if fake.lastRenderOpts.Theme != "dark" || fake.lastRenderOpts.Mode != entity.Mode("network") || fake.lastRenderOpts.PxPerInch != 120 || fake.lastRenderOpts.ArrowStyle != "standard" || fake.lastRenderOpts.PaperMarginLeftIn != 0.25 || fake.lastRenderOpts.SVGLegendPosition != "left" || !fake.lastRenderOpts.CombineFrames {
			t.Fatalf("explicit render opts = %#v", fake.lastRenderOpts)
		}
		if _, err := os.Stat(explicit); err != nil {
			t.Fatalf("explicit output was not created: %v", err)
		}
	}
}

func TestRunRenderFormatWritesOneSVGFilePerFrame(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	output := filepath.Join(dir, "diagram.svg")
	fake := &fakeUseCase{renderArtifacts: []entity.RenderArtifact{
		{ID: "overview", Data: []byte(`<svg id="overview"/>`)},
		{ID: "service/detail", Data: []byte(`<svg id="detail"/>`)},
	}}
	if err := newRenderController(fake).RunFormat(entity.ControllerRenderOptions{
		InputPath: input, OutputPath: output, Format: "svg", Theme: "light",
	}); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"diagram-overview.svg", "diagram-service-detail.svg"} {
		if data, err := os.ReadFile(filepath.Join(dir, name)); err != nil || len(data) == 0 {
			t.Fatalf("artifact %s data=%q err=%v", name, data, err)
		}
	}
	if _, err := os.Stat(output); !os.IsNotExist(err) {
		t.Fatalf("combined SVG should not be written by default: %v", err)
	}
}

func TestRunRenderFormatRejectsCollidingSVGFrameNames(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	fake := &fakeUseCase{renderArtifacts: []entity.RenderArtifact{
		{ID: "Overview", Data: []byte(`<svg/>`)},
		{ID: "overview", Data: []byte(`<svg/>`)},
	}}
	err := newRenderController(fake).RunFormat(entity.ControllerRenderOptions{
		InputPath: input, OutputPath: filepath.Join(dir, "diagram.svg"), Format: "svg", Theme: "light",
	})
	if err == nil || !strings.Contains(err.Error(), "same output") {
		t.Fatalf("collision error = %v", err)
	}
}

func TestRunRenderFormatWithUseCaseErrors(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	if err := newRenderController(&fakeUseCase{validateOptsErr: errors.New("bad options")}).RunFormat(entity.ControllerRenderOptions{InputPath: input, Format: "svg", Theme: "light"}); err == nil {
		t.Fatal("ValidateRenderOptions error = nil")
	}
	if err := newRenderController(&fakeUseCase{}).RunFormat(entity.ControllerRenderOptions{InputPath: input, Format: "unknown", Theme: "light"}); err == nil {
		t.Fatal("unknown format error = nil")
	}
	if err := newRenderController(&fakeUseCase{}).RunFormat(entity.ControllerRenderOptions{InputPath: filepath.Join(dir, "missing.xal"), OutputPath: filepath.Join(dir, "out.svg"), Format: "svg", Theme: "light"}); err == nil {
		t.Fatal("missing input error = nil")
	}
	if err := newRenderController(&fakeUseCase{}).RunFormat(entity.ControllerRenderOptions{InputPath: input, Format: "pptx", Theme: "light"}); err == nil || !strings.Contains(err.Error(), "--output") {
		t.Fatalf("pptx missing output error = %v", err)
	}
	for _, format := range []string{"excalidraw", "pdf", "excel", "xlsx", "xyflow", "isoflow"} {
		t.Run(format+" removed", func(t *testing.T) {
			err := newRenderController(&fakeUseCase{}).RunFormat(entity.ControllerRenderOptions{InputPath: input, OutputPath: filepath.Join(dir, format+".out"), Format: format, Theme: "light"})
			if err == nil || !strings.Contains(err.Error(), "unknown render format") {
				t.Fatalf("removed format error = %v", err)
			}
		})
	}
	missingServices := filepath.Join(dir, "missing-services.csv")
	for _, format := range []string{"svg", "pptx"} {
		t.Run(format+" services", func(t *testing.T) {
			err := newRenderController(&fakeUseCase{}).RunFormat(entity.ControllerRenderOptions{InputPath: input, OutputPath: filepath.Join(dir, format+".out"), Format: format, Theme: "light", ServicesFile: missingServices})
			if err == nil || !strings.Contains(err.Error(), "read services") {
				t.Fatalf("missing services err = %v", err)
			}
		})
	}
	for _, format := range []string{"svg", "pptx"} {
		t.Run(format+" render", func(t *testing.T) {
			err := newRenderController(&fakeUseCase{renderErr: errors.New("render failed")}).RunFormat(entity.ControllerRenderOptions{InputPath: input, OutputPath: filepath.Join(dir, format+"-render.out"), Format: format, Theme: "light"})
			if err == nil || !strings.Contains(err.Error(), "render failed") {
				t.Fatalf("render err = %v", err)
			}
		})
	}
	for _, format := range []string{"svg", "pptx"} {
		t.Run(format+" write", func(t *testing.T) {
			err := newRenderController(&fakeUseCase{}).RunFormat(entity.ControllerRenderOptions{InputPath: input, OutputPath: dir, Format: format, Theme: "light"})
			if err == nil || !strings.Contains(err.Error(), "write output file") {
				t.Fatalf("write err = %v", err)
			}
		})
	}
}

func TestRunServeWithUseCase(t *testing.T) {
	fake := &fakeUseCase{}
	if err := newServeController(&fakeUseCase{previewErr: errors.New("missing")}).Run(context.Background(), entity.ControllerServeOptions{InputPath: filepath.Join(t.TempDir(), "missing.xal"), Theme: "light"}); err == nil {
		t.Fatal("RunServe missing input error = nil")
	}
	err := newServeController(fake).Run(context.Background(), entity.ControllerServeOptions{InputPath: "input.xal", Theme: "light", PollInterval: time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}
	if fake.lastPreviewPath != "input.xal" || fake.lastPreviewOpts.Render.Format != usecase.FormatSVG {
		t.Fatalf("preview path=%q opts=%#v", fake.lastPreviewPath, fake.lastPreviewOpts)
	}
	if fake.lastPreviewAddr != "127.0.0.1:8080" {
		t.Fatalf("preview address = %q, want 127.0.0.1:8080", fake.lastPreviewAddr)
	}
	if err := newServeController(&fakeUseCase{previewErr: errors.New("preview failed")}).Run(nil, entity.ControllerServeOptions{InputPath: "input.xal", Theme: "light"}); err == nil {
		t.Fatal("preview creation error = nil")
	}
}

func TestRunServeUsesConfiguredAndExplicitPorts(t *testing.T) {
	configured := &fakeUseCase{}
	if err := newServeController(configured, 9090).Run(context.Background(), entity.ControllerServeOptions{
		InputPath: "input.xal", Theme: "light", PollInterval: time.Millisecond,
	}); err != nil {
		t.Fatal(err)
	}
	if configured.lastPreviewAddr != "127.0.0.1:9090" {
		t.Fatalf("configured preview address = %q, want 127.0.0.1:9090", configured.lastPreviewAddr)
	}

	overridden := &fakeUseCase{}
	if err := newServeController(overridden, 9090).Run(context.Background(), entity.ControllerServeOptions{
		InputPath: "input.xal", Address: "0.0.0.0:7777", Port: 9191,
		Theme: "light", PollInterval: time.Millisecond,
	}); err != nil {
		t.Fatal(err)
	}
	if overridden.lastPreviewAddr != "0.0.0.0:9191" {
		t.Fatalf("overridden preview address = %q, want 0.0.0.0:9191", overridden.lastPreviewAddr)
	}

	invalid := &fakeUseCase{}
	err := newServeController(invalid).Run(context.Background(), entity.ControllerServeOptions{
		InputPath: "input.xal", Port: 65536, Theme: "light",
	})
	if err == nil || !strings.Contains(err.Error(), "serve port must be between 1 and 65535") {
		t.Fatalf("invalid port error = %v", err)
	}
	if invalid.lastPreviewPath != "" {
		t.Fatalf("preview was created before port validation: %q", invalid.lastPreviewPath)
	}
}

func TestServeCommandPortOverridesAddressPort(t *testing.T) {
	fake := &fakeUseCase{}
	cmd := newServeController(fake, 9090).Command()
	if flag := cmd.Flags().Lookup("port"); flag == nil || flag.DefValue != "9090" {
		t.Fatalf("port flag = %#v, want configured default 9090", flag)
	}
	cmd.SetArgs([]string{"input.xal", "--address", "0.0.0.0:7777", "--port", "9191"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	if fake.lastPreviewAddr != "0.0.0.0:9191" {
		t.Fatalf("preview address = %q, want 0.0.0.0:9191", fake.lastPreviewAddr)
	}

	invalid := newServeController(&fakeUseCase{}).Command()
	invalid.SetArgs([]string{"input.xal", "--port", "0"})
	if err := invalid.Execute(); err == nil || !strings.Contains(err.Error(), "serve port must be between 1 and 65535") {
		t.Fatalf("explicit zero port error = %v", err)
	}
}

func TestRunServeDetectsMarkdownAndForwardsPaperOptions(t *testing.T) {
	fake := &fakeUseCase{}
	err := newServeController(fake).Run(context.Background(), entity.ControllerServeOptions{
		InputPath: "guide.md", Theme: "light", Paper: "A4", Orientation: "landscape", PollInterval: time.Millisecond,
	})
	if err != nil {
		t.Fatal(err)
	}
	if fake.lastPreviewOpts.Kind != entity.PreviewKindHTML {
		t.Fatalf("preview kind = %q, want html", fake.lastPreviewOpts.Kind)
	}
	if fake.lastPreviewOpts.Render.PaperSize != "A4" || fake.lastPreviewOpts.Render.Orientation != "landscape" {
		t.Fatalf("preview render opts = %#v", fake.lastPreviewOpts.Render)
	}

	xalFake := &fakeUseCase{}
	if err := newServeController(xalFake).Run(context.Background(), entity.ControllerServeOptions{
		InputPath: "diagram.xal", Theme: "light", PollInterval: time.Millisecond,
	}); err != nil {
		t.Fatal(err)
	}
	if xalFake.lastPreviewOpts.Kind != entity.PreviewKindSVG && xalFake.lastPreviewOpts.Kind != "" {
		t.Fatalf("preview kind = %q, want svg or empty", xalFake.lastPreviewOpts.Kind)
	}
}

func TestRunGenerateAndInit(t *testing.T) {
	dir := t.TempDir()
	generated := filepath.Join(dir, "generated.xal")
	if err := controller.RunGenerate(1, 1, 1, 1, "grid", 2, "both", "top", "A4", "landscape", generated); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(generated)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `<xaligo version="1">`) ||
		!strings.Contains(string(data), `<frame id="overview" width="1122" height="794"`) ||
		!strings.Contains(string(data), "<aws-cloud") {
		t.Fatalf("generated XAL = %s", data)
	}
	diagnostics, err := usecase.Diagnose(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 0 {
		t.Fatalf("generated XAL diagnostics = %#v, want none", diagnostics)
	}

	leftVertical := filepath.Join(dir, "left-vertical.xal")
	if err := controller.RunGenerate(2, 1, 1, 2, "staggered", 2, "vertical", "left", "A5", "portrait", leftVertical); err != nil {
		t.Fatal(err)
	}
	leftVerticalData, err := os.ReadFile(leftVertical)
	if err != nil {
		t.Fatal(err)
	}
	leftVerticalText := string(leftVerticalData)
	if !strings.Contains(leftVerticalText, `<row gap="16">`) || !strings.Contains(leftVerticalText, `class="pt-2 pb-2"`) || !strings.Contains(leftVerticalText, `layout="staggered"`) || !strings.Contains(leftVerticalText, `az-layout=staggered`) {
		t.Fatalf("left vertical XAL = %s", leftVerticalData)
	}

	leftHorizontal := filepath.Join(dir, "left-horizontal.xal")
	if err := controller.RunGenerate(1, 2, 1, 1, "grid", 2, "horizontal", "left", "A5", "portrait", leftHorizontal); err != nil {
		t.Fatal(err)
	}
	leftHorizontalData, err := os.ReadFile(leftHorizontal)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(leftHorizontalData), `class="pl-2 pr-2"`) {
		t.Fatalf("left horizontal XAL = %s", leftHorizontalData)
	}

	initDir := filepath.Join(dir, "starter")
	if err := controller.NewInitController().Run(initDir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(initDir, "sample.xal")); err != nil {
		t.Fatal(err)
	}
	initialized, err := os.ReadFile(filepath.Join(initDir, "sample.xal"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(initialized), `<xaligo version="1">`) ||
		!strings.Contains(string(initialized), `<frame id="overview"`) {
		t.Fatalf("initialized XAL = %s", initialized)
	}
	diagnostics, err = usecase.Diagnose(context.Background(), initialized)
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 0 {
		t.Fatalf("initialized XAL diagnostics = %#v, want none", diagnostics)
	}
}

func TestGenerateCommandUsesDocumentedDefaults(t *testing.T) {
	output := filepath.Join(t.TempDir(), "generated.xal")
	cmd := controller.NewGenerateController().Command()
	cmd.SetArgs([]string{"xal", "--output", output})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	for _, want := range []string{
		`clouds=1 accounts=1 regions=1 azs=2 az-layout=grid subnets=2 spacing=both start=top`,
		`<frame id="overview" width="1122" height="794"`,
		`<availability-zone id="availability-zone-az-2"`,
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("generated XAL missing %q:\n%s", want, text)
		}
	}
}

func TestRunGenerateRejectsInvalidOptions(t *testing.T) {
	cases := []struct {
		name string
		args []interface{}
	}{
		{"paper", []interface{}{1, 1, 1, 1, "grid", 2, "both", "top", "bad", "portrait"}},
		{"orientation", []interface{}{1, 1, 1, 1, "grid", 2, "both", "top", "A4", "sideways"}},
		{"az layout", []interface{}{1, 1, 1, 1, "free", 2, "both", "top", "A4", "portrait"}},
		{"spacing", []interface{}{1, 1, 1, 1, "grid", 2, "diagonal", "top", "A4", "portrait"}},
		{"start", []interface{}{1, 1, 1, 1, "grid", 2, "both", "middle", "A4", "portrait"}},
		{"range", []interface{}{0, 1, 1, 1, "grid", 2, "both", "top", "A4", "portrait"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := controller.RunGenerate(
				tc.args[0].(int), tc.args[1].(int), tc.args[2].(int), tc.args[3].(int),
				tc.args[4].(string), tc.args[5].(int), tc.args[6].(string), tc.args[7].(string),
				tc.args[8].(string), tc.args[9].(string), filepath.Join(t.TempDir(), "out.xal"),
			)
			if err == nil {
				t.Fatal("RunGenerate error = nil")
			}
		})
	}
	missingDirOutput := filepath.Join(t.TempDir(), "missing", "out.xal")
	if err := controller.RunGenerate(1, 1, 1, 1, "grid", 2, "both", "top", "A4", "portrait", missingDirOutput); err == nil || !strings.Contains(err.Error(), "write output file") {
		t.Fatalf("write error = %v", err)
	}
}
