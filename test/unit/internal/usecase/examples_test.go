package usecase_test

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", "..", ".."))
}

func TestRenderExamplesThroughPublicUseCases(t *testing.T) {
	root := repoRoot(t)
	servicesCSV, err := os.ReadFile(filepath.Join(root, "examples", "services.csv"))
	if err != nil {
		t.Fatal(err)
	}
	onpremServicesCSV, err := os.ReadFile(filepath.Join(root, "examples", "onprem-access-services.csv"))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name       string
		file       string
		services   []byte
		wantScene  string
		wantSVG    string
		wantXYFlow string
		wantIso    string
	}{
		{"sample", "sample.xal", servicesCSV, `"type": "excalidraw"`, `<svg`, `"nodes"`, `"version": "3.3.0"`},
		{"line variants", "line-variants.xal", nil, `"type": "excalidraw"`, `<svg`, `"edges"`, `"connectors"`},
		{"junctions", "junctions.xal", nil, `"type": "excalidraw"`, `<svg`, `"edges"`, `"connectors"`},
		{"complex hybrid", "complex-hybrid-architecture.xal", nil, `"type": "excalidraw"`, `<svg`, `"nodes"`, `"version": "3.3.0"`},
		{"onprem access", "onprem-access.xal", onpremServicesCSV, `"type": "excalidraw"`, `<svg`, `"nodes"`, `"version": "3.3.0"`},
		{"tabler", "tabler.xal", nil, `"type": "excalidraw"`, `<svg`, `"nodes"`, `"version": "3.3.0"`},
		{"yamaha", "yamaha-icons.xal", nil, `"type": "excalidraw"`, `<svg`, `"nodes"`, `"version": "3.3.0"`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source, err := os.ReadFile(filepath.Join(root, "examples", tc.file))
			if err != nil {
				t.Fatal(err)
			}
			opts := entity.RenderOptions{Theme: "light", PxPerInch: 96, ServicesCSV: tc.services, Mode: usecase.ModeNetwork}
			scene, err := usecase.RenderExcalidraw(context.Background(), source, opts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(scene), tc.wantScene) {
				t.Fatalf("scene output missing %q", tc.wantScene)
			}
			svg, err := usecase.RenderSVG(context.Background(), source, opts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(svg), tc.wantSVG) {
				t.Fatalf("svg output missing %q", tc.wantSVG)
			}
			xyflow, err := usecase.RenderXYFlow(context.Background(), source, opts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(xyflow), tc.wantXYFlow) {
				t.Fatalf("xyflow output missing %q", tc.wantXYFlow)
			}
			isoflow, err := usecase.RenderIsoflow(context.Background(), source, opts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(isoflow), tc.wantIso) {
				t.Fatalf("isoflow output missing %q", tc.wantIso)
			}
			plan, err := usecase.BuildPPTXPlan(context.Background(), source, opts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(plan), `"slide"`) {
				t.Fatalf("plan output missing slide")
			}
		})
	}
}