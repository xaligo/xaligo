//go:build cgo && xaligo_engine

package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"math"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestV2SourceUsesRustSVGAndSharedResolvedPPTXPlan(t *testing.T) {
	renderer := usecase.NewRenderUsecase(
		repository.NewSceneRepository(), repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(), repository.NewSVGRepository(),
	)
	source := []byte(`<xaligo version="2"><frames><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><connection id="flow" source="left" target="right" routing="orthogonal"/></frame></frames></xaligo>`)
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
	if len(plan.Ops) == 0 || !bytes.Contains(svg, []byte(`>Left</text>`)) || !bytes.Contains(planJSON, []byte(`"id":"left"`)) {
		t.Fatalf("V2 projections diverged: svg=%s plan=%s", svg, planJSON)
	}
}

func TestV2SVGIsByteStable(t *testing.T) {
	renderer := usecase.NewRenderUsecase(
		repository.NewSceneRepository(), repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(), repository.NewSVGRepository(),
	)
	source := []byte(`<xaligo version="2"><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><connection id="flow" source="left" target="right" routing="orthogonal"/></frame></xaligo>`)
	first, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{Format: usecase.FormatSVG, PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	second, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{Format: usecase.FormatSVG, PxPerInch: 144})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, second) {
		t.Fatal("V2 SVG changed between identical geometry renders")
	}
}

func TestV2StructuredDiagnosticMapsBackToSourceSpan(t *testing.T) {
	source := []byte("<xaligo version=\"2\"><frame width=\"200\" height=\"100\">\n  <connection id=\"bad\" source=\"missing\" target=\"also-missing\"/>\n</frame></xaligo>")
	diagnostics, err := usecase.NewDiagnosticsUsecase().Diagnose(context.Background(), source)
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 1 || diagnostics[0].Code != "XAL-E2001" || diagnostics[0].Element != "bad" || diagnostics[0].Line != 2 {
		t.Fatalf("diagnostics = %#v", diagnostics)
	}
}

func engineFloat(value float64) *float64 {
	return &value
}

func engineUint16(value uint16) *uint16 {
	return &value
}

func engineBool(value bool) *bool {
	return &value
}

func TestRustStaticLibraryEngineThroughCgo(t *testing.T) {
	spec := entity.EngineDocumentSpec{
		Direction: entity.EngineDirectionVertical,
		Width:     120,
		Height:    100,
		Gap:       10,
		Elements: []entity.EngineElementSpec{
			{ID: "header", Width: engineFloat(120), Height: engineFloat(20)},
			{ID: "body", Width: engineFloat(100), Weight: engineFloat(1)},
		},
	}
	engine := v2.NewEngineUsecase()

	resolved, err := engine.Resolve(context.Background(), spec)
	if err != nil {
		t.Fatal(err)
	}
	if len(resolved.Elements) != 2 {
		t.Fatalf("resolved elements = %#v", resolved.Elements)
	}
	body := resolved.Elements[1]
	if body.ID != "body" || body.X != 0 || body.Y != 30 || body.Width != 100 || body.Height != 70 {
		t.Fatalf("resolved body = %#v", body)
	}

	svg, err := engine.RenderSVG(context.Background(), spec)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(svg, []byte(`<svg`)) || !bytes.Contains(svg, []byte(`id="body"`)) {
		t.Fatalf("Rust SVG projection = %s", svg)
	}
}

func TestRustGenericConceptTreeThroughCgo(t *testing.T) {
	spec := entity.EngineDocumentSpec{
		Layout: entity.EngineLayoutAbsolute,
		Width:  400,
		Height: 240,
		Elements: []entity.EngineElementSpec{
			{
				ID:      "canvas",
				Concept: entity.EngineConceptFrame,
				Width:   engineFloat(400),
				Height:  engineFloat(240),
				Layout:  entity.EngineLayoutHorizontal,
				Gap:     engineFloat(20),
				Padding: entity.EngineInsets{
					Top: engineFloat(20), Right: engineFloat(20), Bottom: engineFloat(20), Left: engineFloat(20),
				},
				Children: []entity.EngineElementSpec{
					{
						ID:      "left",
						Concept: entity.EngineConceptGroup,
						Weight:  engineFloat(1),
						Padding: entity.EngineInsets{
							Top: engineFloat(10), Right: engineFloat(10), Bottom: engineFloat(10), Left: engineFloat(10),
						},
						Children: []entity.EngineElementSpec{
							{
								ID:      "source",
								Concept: entity.EngineConceptItem,
								Height:  engineFloat(80),
								Text:    &entity.EngineTextSpec{Value: "Source"},
								Icon:    &entity.EngineIconSpec{Ref: "builtin:service"},
								Children: []entity.EngineElementSpec{
									{
										ID:      "source-out",
										Concept: entity.EngineConceptPort,
										Port: &entity.EnginePortSpec{
											Side: entity.EngineSideRight, Anchor: engineFloat(0), Label: "out",
										},
									},
								},
							},
						},
					},
					{
						ID:      "right",
						Concept: entity.EngineConceptCapture,
						Weight:  engineFloat(1),
						Padding: entity.EngineInsets{
							Top: engineFloat(10), Right: engineFloat(10), Bottom: engineFloat(10), Left: engineFloat(10),
						},
						Children: []entity.EngineElementSpec{
							{ID: "target", Concept: entity.EngineConceptItem, Height: engineFloat(80), Text: &entity.EngineTextSpec{Value: "Target"}},
							{ID: "caption", Concept: entity.EngineConceptText, Text: &entity.EngineTextSpec{Value: "Generic endpoint"}},
						},
					},
					{
						ID:      "flow",
						Concept: entity.EngineConceptLine,
						Line: &entity.EngineLineSpec{
							Source: "source-out", Target: "target", Routing: entity.EngineRoutingOrthogonal,
							TargetDecoration: entity.EngineDecorationArrow, Label: "calls",
						},
					},
				},
			},
		},
	}
	engine := v2.NewEngineUsecase()
	resolved, err := engine.Resolve(context.Background(), spec)
	if err != nil {
		t.Fatal(err)
	}
	if len(resolved.Elements) != 8 {
		t.Fatalf("resolved concepts = %#v", resolved.Elements)
	}
	if got := resolved.Elements[2]; got.ID != "source" || got.ParentID != "left" || got.Concept != entity.EngineConceptItem {
		t.Fatalf("resolved source = %#v", got)
	}
	if got := resolved.Elements[3]; got.ParentID != "source" || got.Concept != entity.EngineConceptPort || got.Y != resolved.Elements[2].Y {
		t.Fatalf("resolved port = %#v", got)
	}
	if got := resolved.Elements[2].IconRef; got != "builtin:service" {
		t.Fatalf("resolved icon = %q", got)
	}
	line := resolved.Elements[7]
	if line.Concept != entity.EngineConceptLine || line.ParentID != "canvas" || len(line.Points) < 2 || line.Line.TargetDecoration != entity.EngineDecorationArrow {
		t.Fatalf("resolved line = %#v", line)
	}

	svg, err := engine.RenderSVG(context.Background(), spec)
	if err != nil {
		t.Fatal(err)
	}
	for _, token := range [][]byte{[]byte(`data-concept="frame"`), []byte(`data-concept="capture"`), []byte(`id="caption" data-owner="caption" data-concept="text"`), []byte(`data-icon="builtin:service"`), []byte(`<polyline`), []byte(`calls`)} {
		if !bytes.Contains(svg, token) {
			t.Fatalf("Rust SVG does not contain %q: %s", token, svg)
		}
	}
}

func TestRustGridAndTypedOptionalValuesThroughCgo(t *testing.T) {
	spec := entity.EngineDocumentSpec{
		Layout:  entity.EngineLayoutGrid,
		Width:   240,
		Height:  100,
		Gap:     4,
		Columns: engineUint16(12),
		Elements: []entity.EngineElementSpec{
			{
				ID: "hidden-slot", Concept: entity.EngineConceptItem, ColumnSpan: engineUint16(4),
				Visual: entity.EngineVisualSpec{Opacity: engineFloat(0), Visible: engineBool(false)},
			},
			{ID: "middle", Concept: entity.EngineConceptSpacer, ColumnSpan: engineUint16(4)},
			{ID: "visible", Concept: entity.EngineConceptItem, ColumnSpan: engineUint16(4)},
		},
	}
	resolved, err := v2.NewEngineUsecase().Resolve(context.Background(), spec)
	if err != nil {
		t.Fatal(err)
	}
	if resolved.Elements[0].Visual.Opacity != 0 || resolved.Elements[0].Visual.Visible {
		t.Fatalf("explicit zero/false were not preserved: %#v", resolved.Elements[0].Visual)
	}
	if resolved.Elements[2].X <= resolved.Elements[0].X {
		t.Fatalf("grid positions = %#v", resolved.Elements)
	}

	invalid := spec
	invalid.Elements = []entity.EngineElementSpec{{ID: "bad", Concept: entity.EngineConceptItem, ColumnSpan: engineUint16(0)}}
	if _, err := v2.NewEngineUsecase().Resolve(context.Background(), invalid); err == nil {
		t.Fatal("explicit zero column span was accepted")
	}
	invalid.Elements = []entity.EngineElementSpec{{ID: "bad", Concept: entity.EngineConceptItem, Width: engineFloat(math.Inf(1))}}
	if _, err := v2.NewEngineUsecase().Resolve(context.Background(), invalid); err == nil {
		t.Fatal("non-finite element width was accepted")
	}
}

func TestRustSVGNormalizationThroughCgo(t *testing.T) {
	engine := v2.NewEngineUsecase()
	normalized, err := engine.NormalizeSVG(context.Background(), []byte(
		`<svg xmlns="http://www.w3.org/2000/svg" width="16" height="8"><path d="M0 0h16v8z"/></svg>`,
	))
	if err != nil {
		t.Fatal(err)
	}
	if normalized.Width != 16 || normalized.Height != 8 || normalized.ViewBox != "0 0 16 8" || !bytes.Contains(normalized.Data, []byte(`<svg`)) {
		t.Fatalf("normalized SVG = %#v %s", normalized, normalized.Data)
	}
	for _, input := range [][]byte{
		[]byte(`<svg xmlns="http://www.w3.org/2000/svg"><script>alert(1)</script></svg>`),
		[]byte(`<svg xmlns="http://www.w3.org/2000/svg"><image href="https://example.com/icon.png"/></svg>`),
		[]byte(`<svg`),
		make([]byte, 2*1024*1024+1),
	} {
		if _, err := engine.NormalizeSVG(context.Background(), input); err == nil {
			t.Fatalf("unsafe or oversized SVG was accepted: %q", input[:min(len(input), 80)])
		}
	}
}
