//go:build cgo && xaligo_engine

package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestV2SourceUsesRustSVGAndSharedResolvedPPTXPlan(t *testing.T) {
	renderer := usecase.NewRenderUsecase(
		repository.NewSceneRepository(), repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository(),
	)
	source := []byte(`<scene version="2" width="320" height="180" layout="absolute"><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><line id="flow" source="left" target="right" routing="orthogonal"/></frame></scene>`)
	svg, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{Format: usecase.FormatSVG, PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(source, []byte(`<scene version="2"`)) || bytes.Contains(source, []byte(`<xaligo version="2"`)) {
		t.Fatal("V2 sample does not use the reject-safe native document root")
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
		repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository(),
	)
	source := []byte(`<scene version="2" width="320" height="180" layout="absolute"><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><line id="flow" source="left" target="right" routing="orthogonal"/></frame></scene>`)
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

func TestComplexHybridV2CompatibilityProjection(t *testing.T) {
	source, err := os.ReadFile(filepath.Join("..", "..", "docs", "src", "examples", "samples", "complex-hybrid-architecture-v2.xal"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(source, []byte(`<scene version="2"`)) || bytes.Contains(source, []byte(`<xaligo version="2"`)) {
		t.Fatal("complex V2 sample does not use the reject-safe native document root")
	}
	renderer := usecase.NewRenderUsecase(
		repository.NewSceneRepository(), repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository(),
	)
	svg, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{Format: usecase.FormatSVG, PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	if got := bytes.Count(svg, []byte(`<polyline`)); got != 36 {
		t.Fatalf("V2 connections = %d, want 36", got)
	}
	if got := bytes.Count(svg, []byte(`<image`)); got != 39 {
		t.Fatalf("V2 embedded catalog icons = %d, want 39", got)
	}
	if !bytes.Contains(svg, []byte(`r="5" fill="#ffffff"`)) || !bytes.Contains(svg, []byte(`r="3.5"`)) {
		t.Fatal("V2 line jumps or junctions were not rendered")
	}
	for _, token := range [][]byte{
		[]byte(`>Complex Hybrid Architecture</text>`),
		[]byte(`>R53</text>`),
		[]byte(`data-icon="catalog:200025"`),
		[]byte(`stroke="#2563eb"`),
		[]byte(`stroke-dasharray="8 5"`),
	} {
		if !bytes.Contains(svg, token) {
			t.Fatalf("V2 compatibility SVG does not contain %q", token)
		}
	}
	planJSON, err := renderer.BuildPPTXPlan(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		t.Fatal(err)
	}
	var imageCount, lineCount int
	var employeeIcon *entity.DrawOp
	for index := range plan.Ops {
		op := &plan.Ops[index]
		switch op.Kind {
		case "image":
			imageCount++
			if !op.FrontLayer {
				t.Fatalf("PPTX image %q is not projected above connections", op.ID)
			}
		case "line":
			lineCount++
		}
		if op.ID == "200025-icon" {
			employeeIcon = op
		}
	}
	if imageCount != 39 || lineCount != 36 {
		t.Fatalf("V2 PPTX plan has %d images and %d lines, want 39 and 36", imageCount, lineCount)
	}
	if employeeIcon == nil {
		t.Fatal("V2 PPTX plan does not contain the employee icon")
	}
	svgEmployeeIcon := svgElementGeometry(t, svg, "200025-icon")
	for index, got := range []float64{employeeIcon.X, employeeIcon.Y, employeeIcon.W, employeeIcon.H} {
		if want := svgEmployeeIcon[index] / 96; math.Abs(got-want) > 0.00001 {
			t.Fatalf("employee icon PPTX geometry = (%v, %v, %v, %v), SVG geometry = %v", employeeIcon.X, employeeIcon.Y, employeeIcon.W, employeeIcon.H, svgEmployeeIcon)
		}
	}
	for id, want := range map[string][4]float64{
		"onprem-hq":          {24, 47, 462, 1369},
		"aws-cloud":          {510, 47, 1386, 1369},
		"aws-edge":           {522, 83, 1362, 111.364},
		"region-apne1":       {522, 214.364, 1362, 1189.636},
		"prod-vpc":           {534, 376.331, 1338, 1015.669},
		"vpc-edge-security":  {546, 416.331, 1314, 86.152},
		"az-apne1a":          {546, 518.482, 649, 861.518},
		"public-subnet-a":    {558, 562.482, 625, 257.839},
		"elastic-app-tier-a": {570, 880.322, 601, 201.839},
	} {
		got := svgElementGeometry(t, svg, id)
		for index := range want {
			if math.Abs(got[index]-want[index]) > 0.002 {
				t.Fatalf("%s geometry = %v, want %v", id, got, want)
			}
		}
	}
	var root struct {
		XMLName xml.Name
	}
	if err := xml.Unmarshal(svg, &root); err != nil {
		t.Fatal(err)
	}
	if root.XMLName.Local != "svg" {
		t.Fatalf("SVG root = %q", root.XMLName.Local)
	}
}

func svgElementGeometry(t *testing.T, svg []byte, id string) [4]float64 {
	t.Helper()
	decoder := xml.NewDecoder(bytes.NewReader(svg))
	for {
		token, err := decoder.Token()
		if err != nil {
			t.Fatalf("find SVG element %q: %v", id, err)
		}
		start, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		attrs := make(map[string]string, len(start.Attr))
		for _, attr := range start.Attr {
			attrs[attr.Name.Local] = attr.Value
		}
		if attrs["id"] != id {
			continue
		}
		var geometry [4]float64
		for index, name := range []string{"x", "y", "width", "height"} {
			geometry[index], err = strconv.ParseFloat(attrs[name], 64)
			if err != nil {
				t.Fatalf("SVG element %q %s=%q: %v", id, name, attrs[name], err)
			}
		}
		return geometry
	}
}

func TestV2StructuredDiagnosticMapsBackToSourceSpan(t *testing.T) {
	source := []byte("<scene version=\"2\" width=\"200\" height=\"100\">\n  <line id=\"bad\" source=\"missing\" target=\"also-missing\"/>\n</scene>")
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
