//go:build cgo && xaligo_engine

package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestNativeAWSComponentsShareSVGAndPPTXParts(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	for _, tag := range []string{"aws-elastic-load-balancing-application-load-balancer", "aws-elastic-load-balancing-network-load-balancer"} {
		source, err := os.ReadFile(filepath.Join("..", "..", "docs", "src", "examples", "samples", "aws", tag, "sample.xal"))
		if err != nil {
			t.Fatal(err)
		}
		svg, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{Format: usecase.FormatSVG, PxPerInch: 96})
		if err != nil {
			t.Fatal(err)
		}
		planJSON, err := renderer.BuildPPTXPlan(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
		if err != nil {
			t.Fatal(err)
		}
		var plan entity.Plan
		if err := json.Unmarshal(planJSON, &plan); err != nil {
			t.Fatal(err)
		}
		if len(plan.Ops) < 20 {
			t.Fatalf("missing native parts: %d", len(plan.Ops))
		}
		for _, value := range []string{"TLS OFF", "TLS ON", "mTLS OFF", "component::aws/domain", ".example.test"} {
			if !bytes.Contains(svg, []byte(value)) || !bytes.Contains(planJSON, []byte(value)) {
				t.Errorf("%s: missing shared part %q", tag, value)
			}
		}
		if strings.HasSuffix(tag, "-network-load-balancer") {
			for _, value := range []string{"Target TLS ON", "Target mTLS ON", "Target mTLS OFF"} {
				if bytes.Contains(svg, []byte(value)) || bytes.Contains(planJSON, []byte(value)) {
					t.Errorf("duplicate target security badge: %s", value)
				}
			}
		}
	}
}

func TestNativeAWSInvalidServiceParametersFailThroughABI(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	for _, input := range []struct{ tag, listener string }{
		{"application", `protocol="TCP" port="443"`},
		{"application", `protocol="HTTPS" port="443" mtls="verify"`},
		{"network", `protocol="TLS" port="443" mtls="verify" trust-store="ca"`},
		{"network", `protocol="TLS" port="443" backend-tls="true" backend-mtls="true"`},
		{"network", `protocol="TCP" port="443" certificate="server"`},
	} {
		tag := "aws-elastic-load-balancing-" + input.tag + "-load-balancer"
		source := `<xaligo version="2"><frames><frame id="p" width="800" height="600"><` + tag + ` id="lb"><aws-listener id="listener" ` + input.listener + `/></` + tag + `></frame></frames></xaligo>`
		if _, err := renderer.RenderSVG(context.Background(), []byte(source), entity.RenderOptions{Format: usecase.FormatSVG}); err == nil {
			t.Errorf("accepted invalid %s listener %s", input.tag, input.listener)
		}
	}
}

func TestNativeAWSHiddenListenerTitleSharesSVGAndPPTX(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	source, err := os.ReadFile(filepath.Join("..", "..", "docs", "src", "examples", "samples", "aws", "aws-elastic-load-balancing-network-load-balancer", "hidden-title.xal"))
	if err != nil {
		t.Fatal(err)
	}
	for _, show := range []bool{false, true} {
		input := source
		if show {
			input = bytes.ReplaceAll(source, []byte(` show-title="false"`), nil)
		}
		svg, err := renderer.RenderSVG(context.Background(), input, entity.RenderOptions{Format: usecase.FormatSVG, PxPerInch: 96})
		if err != nil {
			t.Fatal(err)
		}
		plan, err := renderer.BuildPPTXPlan(context.Background(), input, entity.RenderOptions{PxPerInch: 96})
		if err != nil {
			t.Fatal(err)
		}
		for _, output := range [][]byte{svg, plan} {
			if bytes.Contains(output, []byte("tcp::aws/title")) != show {
				t.Fatalf("title visibility not preserved: %v", show)
			}
			for _, label := range []string{"TLS OFF", "mTLS OFF"} {
				if !bytes.Contains(output, []byte(label)) {
					t.Fatalf("title option hid security badge: %s", label)
				}
			}
		}
	}
}

func TestAWSComponentSamplesRenderWithBothEngines(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	options := entity.RenderOptions{Format: usecase.FormatSVG, PxPerInch: 96, CombineFrames: true}
	for _, definition := range awsprofile.Definitions() {
		t.Run(definition.Tag, func(t *testing.T) {
			directory := filepath.Join("..", "..", "docs", "src", "examples", "samples", "aws", definition.Tag)
			files, err := filepath.Glob(filepath.Join(directory, "*.xal"))
			if err != nil || len(files) == 0 {
				t.Fatalf("missing AWS sources: %v", err)
			}
			for _, filename := range files {
				source, err := os.ReadFile(filename)
				if err != nil {
					t.Fatal(err)
				}
				for _, version := range []string{"1", "2"} {
					input := bytes.Replace(source, []byte(`version="2"`), []byte(`version="`+version+`"`), 1)
					svg, err := renderer.RenderSVG(context.Background(), input, options)
					if version == "1" && bytes.Contains(source, []byte("<aws-listener")) {
						if err == nil || !strings.Contains(err.Error(), "requires XAL version 2") {
							t.Fatalf("native component V1 boundary: %v", err)
						}
						continue
					}
					if err != nil {
						t.Fatalf("V%s render: %v", version, err)
					}
					if !bytes.Contains(svg, []byte("<svg")) || bytes.Contains(svg, []byte("NaN")) {
						t.Fatalf("V%s invalid SVG", version)
					}
					if definition.CatalogID != 0 && !bytes.Contains(svg, []byte("data:image/svg+xml;base64,")) {
						t.Fatalf("V%s missing embedded AWS icon", version)
					}
					if version == "2" {
						stored, err := os.ReadFile(strings.TrimSuffix(filename, ".xal") + ".svg")
						if err != nil {
							t.Fatal(err)
						}
						if !bytes.Equal(bytes.TrimSpace(stored), bytes.TrimSpace(svg)) {
							for i := 0; i < len(stored) && i < len(svg); i++ {
								if stored[i] != svg[i] {
									t.Logf("first difference at %d: stored=%q rendered=%q", i, string(stored[max(0, i-50):min(len(stored), i+100)]), string(svg[max(0, i-50):min(len(svg), i+100)]))
									break
								}
							}
							t.Fatal("sample.svg is stale; run npm run generate:aws-samples -- --render")
						}
					}
				}
			}
		})
	}
}

func TestAWSDedicatedResourceIDsAreUnique(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	for _, version := range []string{"1", "2"} {
		source := `<xaligo version="` + version + `"><frames><frame id="p"><aws-ec2 id="same"/><aws-s3 id="same"/></frame></frames></xaligo>`
		_, err := renderer.RenderSVG(context.Background(), []byte(source), entity.RenderOptions{Format: usecase.FormatSVG})
		if err == nil || !strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			t.Fatalf("V%s duplicate IDs: %v", version, err)
		}
	}
}
