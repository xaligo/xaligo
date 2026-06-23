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

	"github.com/ryo-arima/xaligo/internal/controller"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
)

type fakeUseCase struct {
	validateErr      error
	validateOptsErr  error
	renderErr        error
	previewErr       error
	lastRenderOpts   entity.RenderOptions
	lastPlanOpts     entity.RenderOptions
	lastPreviewPath  string
	lastPreviewOpts  entity.PreviewOptions
	renderExcalidraw []byte
	renderSVG        []byte
	renderXYFlow     []byte
	renderIsoflow    []byte
	planJSON         []byte
}

func (f *fakeUseCase) ValidateRenderOptions(opts entity.RenderOptions) error {
	if f.validateOptsErr != nil {
		return f.validateOptsErr
	}
	return usecase.ValidateRenderOptions(opts)
}

func (f *fakeUseCase) Validate(context.Context, []byte) error { return f.validateErr }

func (f *fakeUseCase) Diagnose(context.Context, []byte) ([]entity.Diagnostic, error) { return nil, nil }

func (f *fakeUseCase) Render(context.Context, []byte, entity.RenderOptions) ([]byte, error) {
	return []byte(`render`), f.renderErr
}

func (f *fakeUseCase) RenderExcalidraw(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	f.lastRenderOpts = opts
	if f.renderExcalidraw != nil {
		return f.renderExcalidraw, f.renderErr
	}
	return []byte(`{"type":"excalidraw","elements":[],"files":{}}`), f.renderErr
}

func (f *fakeUseCase) RenderSVG(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	f.lastRenderOpts = opts
	if f.renderSVG != nil {
		return f.renderSVG, f.renderErr
	}
	return []byte(`<svg></svg>`), f.renderErr
}

func (f *fakeUseCase) RenderPPTX(context.Context, []byte, entity.RenderOptions) ([]byte, error) {
	return []byte(`pptx`), f.renderErr
}

func (f *fakeUseCase) RenderXYFlow(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	f.lastRenderOpts = opts
	if f.renderXYFlow != nil {
		return f.renderXYFlow, f.renderErr
	}
	return []byte(`{"nodes":[],"edges":[]}`), f.renderErr
}

func (f *fakeUseCase) RenderIsoflow(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	f.lastRenderOpts = opts
	if f.renderIsoflow != nil {
		return f.renderIsoflow, f.renderErr
	}
	return []byte(`{"version":"3.3.0"}`), f.renderErr
}

func (f *fakeUseCase) BuildPPTXPlan(_ context.Context, _ []byte, opts entity.RenderOptions) ([]byte, error) {
	f.lastPlanOpts = opts
	if f.planJSON != nil {
		return f.planJSON, f.renderErr
	}
	return []byte(`{"slide":{"w":1,"h":1}}`), f.renderErr
}

func (f *fakeUseCase) NewPreviewServer(path string, opts entity.PreviewOptions) (usecase.PreviewServer, error) {
	f.lastPreviewPath = path
	f.lastPreviewOpts = opts
	if f.previewErr != nil {
		return nil, f.previewErr
	}
	return fakePreviewServer{}, nil
}

type fakePreviewServer struct{}

func (fakePreviewServer) Handler() http.Handler             { return http.NewServeMux() }
func (fakePreviewServer) Run(context.Context, string) error { return nil }
func (fakePreviewServer) Refresh() error                    { return nil }

func writeTempXAL(t *testing.T, dir string) string {
	t.Helper()
	path := filepath.Join(dir, "input.xal")
	if err := os.WriteFile(path, []byte(`<frame width="120" height="80"><blank /></frame>`), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestControllerCommandInitializers(t *testing.T) {
	commands := []*cobra.Command{
		controller.InitAddCmd(),
		controller.InitGenerateCmd(),
		controller.InitInitCmd(),
		controller.InitServeCmd(),
		controller.InitServeCmdWithUseCase(&fakeUseCase{}),
		controller.InitValidateCmd(),
		controller.InitValidateCmdWithUseCase(&fakeUseCase{}),
		controller.InitVersionCmd(),
	}
	for _, cmd := range commands {
		if cmd.Use == "" || cmd.Short == "" {
			t.Fatalf("command missing metadata: %#v", cmd)
		}
	}
}

func TestRunValidateWithUseCase(t *testing.T) {
	input := writeTempXAL(t, t.TempDir())
	if err := controller.RunValidate(input, nil); err != nil {
		t.Fatal(err)
	}
	if err := controller.RunValidateWithUseCase(&fakeUseCase{}, input, nil); err != nil {
		t.Fatal(err)
	}
	if err := controller.RunValidateWithUseCase(&fakeUseCase{validateErr: errors.New("invalid")}, input, nil); err == nil || !strings.Contains(err.Error(), "invalid") {
		t.Fatalf("validation error = %v", err)
	}
	if err := controller.RunValidateWithUseCase(&fakeUseCase{}, filepath.Join(t.TempDir(), "missing.xal"), nil); err == nil {
		t.Fatal("missing input error = nil")
	}
}

func TestRunRenderFormatWithUseCaseWritesFormats(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	services := filepath.Join(dir, "services.csv")
	if err := os.WriteFile(services, []byte("27,Amazon EC2,EC2\n"), 0644); err != nil {
		t.Fatal(err)
	}
	formats := []string{"excalidraw", "svg", "xyflow", "isoflow"}
	for _, format := range formats {
		t.Run(format, func(t *testing.T) {
			output := filepath.Join(dir, "out."+format)
			fake := &fakeUseCase{}
			err := controller.RunRenderFormatWithUseCase(fake, entity.ControllerRenderOptions{
				InputPath: input, OutputPath: output, Format: format, ServicesFile: services, Theme: "light", Mode: "standard",
			})
			if err != nil {
				t.Fatal(err)
			}
			data, err := os.ReadFile(output)
			if err != nil {
				t.Fatal(err)
			}
			if len(data) == 0 || fake.lastRenderOpts.Abbreviations[27] != "EC2" {
				t.Fatalf("data=%q opts=%#v", data, fake.lastRenderOpts)
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
		"excalidraw": "output.excalidraw",
		"svg":        "output.svg",
		"xyflow":     "output.xyflow.json",
		"isoflow":    "output.isoflow.json",
	}
	for format, output := range formats {
		t.Run(format, func(t *testing.T) {
			cmd := controller.InitRenderCmdWithUseCase(&fakeUseCase{})
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
		cmd := controller.InitRenderCmdWithUseCase(fake)
		cmd.SetArgs([]string{input, "--format", "svg", "--output", explicit, "--compression", "--theme", "dark", "--mode", "network", "--px-per-inch", "120", "--arrow-style", "standard", "--arrow-stub", "22", "--arrow-margin", "11", "--paper", "A4", "--orientation", "landscape", "--paper-margin-left", "0.25"})
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
		if fake.lastRenderOpts.Theme != "dark" || fake.lastRenderOpts.Mode != entity.Mode("network") || fake.lastRenderOpts.PxPerInch != 120 || fake.lastRenderOpts.ArrowStyle != "standard" || fake.lastRenderOpts.PaperMarginLeftIn != 0.25 {
			t.Fatalf("explicit render opts = %#v", fake.lastRenderOpts)
		}
		if _, err := os.Stat(explicit); err != nil {
			t.Fatalf("explicit output was not created: %v", err)
		}
	}
}

func TestRunRenderFormatWithUseCaseErrors(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	if err := controller.RunRenderFormatWithUseCase(&fakeUseCase{validateOptsErr: errors.New("bad options")}, entity.ControllerRenderOptions{InputPath: input, Format: "svg", Theme: "light"}); err == nil {
		t.Fatal("ValidateRenderOptions error = nil")
	}
	if err := controller.RunRenderFormatWithUseCase(&fakeUseCase{}, entity.ControllerRenderOptions{InputPath: input, Format: "unknown", Theme: "light"}); err == nil {
		t.Fatal("unknown format error = nil")
	}
	if err := controller.RunRenderFormatWithUseCase(&fakeUseCase{}, entity.ControllerRenderOptions{InputPath: filepath.Join(dir, "missing.xal"), OutputPath: filepath.Join(dir, "out.svg"), Format: "svg", Theme: "light"}); err == nil {
		t.Fatal("missing input error = nil")
	}
	if err := controller.RunRenderFormatWithUseCase(&fakeUseCase{}, entity.ControllerRenderOptions{InputPath: input, Format: "pptx", Theme: "light"}); err == nil || !strings.Contains(err.Error(), "--output") {
		t.Fatalf("pptx missing output error = %v", err)
	}
	missingServices := filepath.Join(dir, "missing-services.csv")
	for _, format := range []string{"excalidraw", "svg", "xyflow", "isoflow"} {
		t.Run(format+" services", func(t *testing.T) {
			err := controller.RunRenderFormatWithUseCase(&fakeUseCase{}, entity.ControllerRenderOptions{InputPath: input, OutputPath: filepath.Join(dir, format+".out"), Format: format, Theme: "light", ServicesFile: missingServices})
			if err == nil || !strings.Contains(err.Error(), "read services") {
				t.Fatalf("missing services err = %v", err)
			}
		})
	}
	for _, format := range []string{"excalidraw", "svg", "xyflow", "isoflow"} {
		t.Run(format+" render", func(t *testing.T) {
			err := controller.RunRenderFormatWithUseCase(&fakeUseCase{renderErr: errors.New("render failed")}, entity.ControllerRenderOptions{InputPath: input, OutputPath: filepath.Join(dir, format+"-render.out"), Format: format, Theme: "light"})
			if err == nil || !strings.Contains(err.Error(), "render failed") {
				t.Fatalf("render err = %v", err)
			}
		})
	}
	for _, format := range []string{"excalidraw", "svg", "xyflow", "isoflow"} {
		t.Run(format+" write", func(t *testing.T) {
			err := controller.RunRenderFormatWithUseCase(&fakeUseCase{}, entity.ControllerRenderOptions{InputPath: input, OutputPath: dir, Format: format, Theme: "light"})
			if err == nil || !strings.Contains(err.Error(), "write output file") {
				t.Fatalf("write err = %v", err)
			}
		})
	}
}

func TestRunServeWithUseCase(t *testing.T) {
	fake := &fakeUseCase{}
	if err := controller.RunServe(context.Background(), entity.ControllerServeOptions{InputPath: filepath.Join(t.TempDir(), "missing.xal"), Theme: "light"}); err == nil {
		t.Fatal("RunServe missing input error = nil")
	}
	err := controller.RunServeWithUseCase(fake, context.Background(), entity.ControllerServeOptions{InputPath: "input.xal", Theme: "light", PollInterval: time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}
	if fake.lastPreviewPath != "input.xal" || fake.lastPreviewOpts.Render.Format != usecase.FormatSVG {
		t.Fatalf("preview path=%q opts=%#v", fake.lastPreviewPath, fake.lastPreviewOpts)
	}
	if err := controller.RunServeWithUseCase(&fakeUseCase{previewErr: errors.New("preview failed")}, nil, entity.ControllerServeOptions{InputPath: "input.xal", Theme: "light"}); err == nil {
		t.Fatal("preview creation error = nil")
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
	if !strings.Contains(string(data), `<frame width="1122" height="794"`) || !strings.Contains(string(data), "<aws-cloud") {
		t.Fatalf("generated XAL = %s", data)
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
	if err := controller.RunInit(initDir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(initDir, "sample.xal")); err != nil {
		t.Fatal(err)
	}
}

func TestRunAddServiceBatchAddsIconsAndLegend(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "diagram.excalidraw")
	scene := entity.NewScene()
	scene.Elements = append(scene.Elements, map[string]interface{}{
		"id": "paper-frame", "type": "frame", "x": float64(0), "y": float64(0), "width": float64(300), "height": float64(200),
	})
	if err := repository.WriteScene(scene, target); err != nil {
		t.Fatal(err)
	}
	services := filepath.Join(dir, "services.csv")
	if err := os.WriteFile(services, []byte("27,Amazon EC2,EC2\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := controller.RunAddServiceBatch(target, services); err != nil {
		t.Fatal(err)
	}
	updated, err := repository.ReadScene(target)
	if err != nil {
		t.Fatal(err)
	}
	if len(updated.Files) != 1 || len(updated.Elements) < 5 {
		t.Fatalf("updated scene files=%d elements=%d scene=%#v", len(updated.Files), len(updated.Elements), updated)
	}
	foundLegend := false
	for _, element := range updated.Elements {
		id, _ := element["id"].(string)
		if strings.Contains(id, "-lg-ico") {
			foundLegend = true
			break
		}
	}
	if !foundLegend {
		t.Fatalf("legend icon not found in %#v", updated.Elements)
	}
	if err := controller.RunAddServiceBatch(target, filepath.Join(dir, "missing.csv")); err == nil {
		t.Fatal("missing service list error = nil")
	}
}

func TestAddServiceCommandSingleModeFindsIcon(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "diagram.excalidraw")
	scene := entity.NewScene()
	scene.Elements = append(scene.Elements, map[string]interface{}{
		"id": "box", "type": "rectangle", "x": float64(10), "y": float64(20), "width": float64(300), "height": float64(200),
	})
	if err := repository.WriteScene(scene, target); err != nil {
		t.Fatal(err)
	}
	cmd := controller.InitAddCmd()
	cmd.SetArgs([]string{"service", "--file", target, "--name", "Amazon EC2", "--category", "Arch_Compute", "--size", "48", "--no-legend"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	updated, err := repository.ReadScene(target)
	if err != nil {
		t.Fatal(err)
	}
	if len(updated.Files) != 1 || len(updated.Elements) < 3 {
		t.Fatalf("updated scene files=%d elements=%d", len(updated.Files), len(updated.Elements))
	}

	withLegend := filepath.Join(dir, "diagram-with-legend.excalidraw")
	legendScene := entity.NewScene()
	legendScene.Elements = append(legendScene.Elements, map[string]interface{}{
		"id": "paper-frame", "type": "frame", "x": float64(50), "y": float64(10), "width": float64(300), "height": float64(80),
	})
	if err := repository.WriteScene(legendScene, withLegend); err != nil {
		t.Fatal(err)
	}
	cmd = controller.InitAddCmd()
	cmd.SetArgs([]string{"service", "--file", withLegend, "--name", "Amazon EC2", "--category", "Arch_Compute", "--size", "48"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	legendUpdated, err := repository.ReadScene(withLegend)
	if err != nil {
		t.Fatal(err)
	}
	foundLeftLegend := false
	for _, element := range legendUpdated.Elements {
		id, _ := element["id"].(string)
		x, _ := element["x"].(float64)
		if strings.Contains(id, "-lg-ico") && x < 50 {
			foundLeftLegend = true
			break
		}
	}
	if !foundLeftLegend {
		t.Fatalf("left legend icon not found: %#v", legendUpdated.Elements)
	}

	noFrame := filepath.Join(dir, "diagram-no-frame.excalidraw")
	noFrameScene := entity.NewScene()
	if err := repository.WriteScene(noFrameScene, noFrame); err != nil {
		t.Fatal(err)
	}
	cmd = controller.InitAddCmd()
	cmd.SetArgs([]string{"service", "--file", noFrame, "--name", "Amazon Simple Storage Service", "--size", "48", "--no-legend"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
	noFrameUpdated, err := repository.ReadScene(noFrame)
	if err != nil {
		t.Fatal(err)
	}
	if len(noFrameUpdated.Files) != 1 || len(noFrameUpdated.Elements) < 2 {
		t.Fatalf("no-frame updated files=%d elements=%d", len(noFrameUpdated.Files), len(noFrameUpdated.Elements))
	}
	foundS3Label := false
	for _, element := range noFrameUpdated.Elements {
		if element["type"] == "text" {
			if text, _ := element["text"].(string); text == "S3" {
				foundS3Label = true
			}
		}
	}
	if !foundS3Label {
		t.Fatalf("S3 label not found: %#v", noFrameUpdated.Elements)
	}
}

func TestRunGeneratePptxWithUseCaseReachesPlanBuild(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	services := filepath.Join(dir, "services.csv")
	if err := os.WriteFile(services, []byte("27,Amazon EC2,EC2\n"), 0644); err != nil {
		t.Fatal(err)
	}
	badWASM := filepath.Join(dir, "bad.wasm")
	if err := os.WriteFile(badWASM, []byte("not wasm"), 0644); err != nil {
		t.Fatal(err)
	}
	fake := &fakeUseCase{planJSON: []byte(`{"slide":{"w":1,"h":1}}`)}
	err := controller.RunGeneratePptxWithUseCase(fake, entity.ControllerPptxGenerateOptions{
		XalPath: input, Output: filepath.Join(dir, "out.pptx"), ServicesFile: services, Theme: "dark", Mode: "network", ExporterWASM: badWASM,
		PxPerInch: 120, ArrowStyle: "orthogonal", ArrowStub: 24, ArrowMargin: 12, Paper: "A3", Orientation: "landscape",
		PaperMargin: 0.5, PaperMarginTop: 0.25, PaperMarginRight: 0.3, PaperMarginBottom: 0.35, PaperMarginLeft: 0.4,
	})
	if err == nil || !strings.Contains(err.Error(), "run PPTX WASM exporter") {
		t.Fatalf("err = %v", err)
	}
	if fake.lastPlanOpts.Theme != "dark" || fake.lastPlanOpts.Mode != entity.Mode("network") || fake.lastPlanOpts.PxPerInch != 120 || fake.lastPlanOpts.PaperMarginLeftIn != 0.4 || fake.lastPlanOpts.ServicesCSV == nil {
		t.Fatalf("plan opts = %#v", fake.lastPlanOpts)
	}
	if err := controller.RunGeneratePptx(entity.ControllerPptxGenerateOptions{}); err == nil || !strings.Contains(err.Error(), "--xal") {
		t.Fatalf("RunGeneratePptx missing xal err = %v", err)
	}
	if err := controller.RunGeneratePptx(entity.ControllerPptxGenerateOptions{XalPath: input, Output: filepath.Join(dir, "real.pptx"), ServicesFile: services, Theme: "light", ExporterWASM: badWASM}); err == nil || !strings.Contains(err.Error(), "run PPTX WASM exporter") {
		t.Fatalf("RunGeneratePptx real planner err = %v", err)
	}
	if err := controller.RunGeneratePptxWithUseCase(fake, entity.ControllerPptxGenerateOptions{XalPath: input, Output: filepath.Join(dir, "out.pptx"), PxPerInch: -1}); err == nil || !strings.Contains(err.Error(), "px-per-inch") {
		t.Fatalf("negative px err = %v", err)
	}
	if err := controller.RunGeneratePptxWithUseCase(fake, entity.ControllerPptxGenerateOptions{XalPath: input, Output: filepath.Join(dir, "out.pptx"), PaperMargin: -1}); err == nil || !strings.Contains(err.Error(), "paper margins") {
		t.Fatalf("negative margin err = %v", err)
	}
}

func TestRunRenderWritesExcalidraw(t *testing.T) {
	dir := t.TempDir()
	input := writeTempXAL(t, dir)
	output := filepath.Join(dir, "out.excalidraw")
	if err := controller.RunRender(input, output, nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"type": "excalidraw"`) {
		t.Fatalf("render output = %s", data)
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
