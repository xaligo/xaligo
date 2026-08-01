package usecase_test

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
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
	exampleDir := filepath.Join(root, "docs", "src", "examples", "samples")
	servicesCSV, err := os.ReadFile(filepath.Join(exampleDir, "services.csv"))
	if err != nil {
		t.Fatal(err)
	}
	onpremServicesCSV, err := os.ReadFile(filepath.Join(exampleDir, "onprem-access-services.csv"))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name      string
		file      string
		services  []byte
		wantScene string
		wantSVG   string
	}{
		{"sample", "sample.xal", servicesCSV, `"type": "excalidraw"`, `<svg`},
		{"line variants", "line-variants.xal", nil, `"type": "excalidraw"`, `<svg`},
		{"junctions", "junctions.xal", nil, `"type": "excalidraw"`, `<svg`},
		{"complex hybrid", "complex-hybrid-architecture.xal", nil, `"type": "excalidraw"`, `<svg`},
		{"onprem access", "onprem-access.xal", onpremServicesCSV, `"type": "excalidraw"`, `<svg`},
		{"tabler", "tabler.xal", nil, `"type": "excalidraw"`, `<svg`},
		{"yamaha", "yamaha-icons.xal", nil, `"type": "excalidraw"`, `<svg`},
		{"UML class", "uml-class.xal", nil, `"xaligoCrossFrame": true`, `Repository Contract`},
		{"canonical V1 envelope", "canonical-v1-envelope.xal", nil, `to \u003cdatabase-detail\u003e`, `from &lt;overview&gt;`},
		{"cross-frame page links", "page-links.xal", nil, `to \u003cservice-detail\u003e`, `from &lt;overview&gt;`},
		{"frame metadata", "frame-metadata.xal", nil, `"xaligoFrameMetadata": true`, `AWS Architecture`},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source, err := os.ReadFile(filepath.Join(exampleDir, tc.file))
			if err != nil {
				t.Fatal(err)
			}
			opts := entity.RenderOptions{Theme: "light", PxPerInch: 96, ServicesCSV: tc.services, Mode: usecase.ModeNetwork}
			scene, err := newUsecase().BuildScene(context.Background(), source, opts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(scene), tc.wantScene) {
				t.Fatalf("scene output missing %q", tc.wantScene)
			}
			svgOpts := opts
			svgOpts.CombineFrames = true
			svg, err := newUsecase().RenderSVG(context.Background(), source, svgOpts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(svg), tc.wantSVG) {
				t.Fatalf("svg output missing %q", tc.wantSVG)
			}
			plan, err := newUsecase().BuildPPTXPlan(context.Background(), source, opts)
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(plan), `"slide"`) {
				t.Fatalf("plan output missing slide")
			}
		})
	}
}

func TestRenderActivityCrossFrameTargetExample(t *testing.T) {
	root := repoRoot(t)
	source, err := os.ReadFile(filepath.Join(root, "docs", "src", "examples", "targets", "uml-activity-cross-frame.xal"))
	if err != nil {
		t.Fatal(err)
	}
	opts := entity.RenderOptions{Theme: "light", PxPerInch: 96, Mode: usecase.ModeNetwork}
	scene, err := newUsecase().BuildScene(context.Background(), source, opts)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"xaligoCrossFrame": true`, `to \u003cbank-detail\u003e`, `from \u003coverview\u003e`} {
		if !strings.Contains(string(scene), want) {
			t.Fatalf("scene output missing %q", want)
		}
	}
	svgOpts := opts
	svgOpts.CombineFrames = true
	svg, err := newUsecase().RenderSVG(context.Background(), source, svgOpts)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`to &lt;bank-detail&gt;`, `from &lt;overview&gt;`, `#04B79F`} {
		if !strings.Contains(string(svg), want) {
			t.Fatalf("svg output missing %q", want)
		}
	}
}
