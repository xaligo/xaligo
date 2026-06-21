package usecase_test

import (
	"encoding/json"
	"testing"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

func TestBuildPlanPreservesConnectorStylesAndLegend(t *testing.T) {
	opacity := 100.0
	scene := entity.PptxScene{Elements: []entity.Element{
		{ID: "src-item", Type: "image", X: 0, Y: 0, Width: 32, Height: 32, Opacity: &opacity, FileID: "src-file"},
		{ID: "src-item-lbl", Type: "text", X: -12, Y: 36, Width: 56, Height: 14, Text: "SRC", Opacity: &opacity},
		{ID: "dst-item", Type: "image", X: 160, Y: 0, Width: 32, Height: 32, Opacity: &opacity, FileID: "dst-file"},
		{ID: "dst-item-lbl", Type: "text", X: 148, Y: 36, Width: 56, Height: 14, Text: "DST", Opacity: &opacity},
		{ID: "route", Type: "arrow", StrokeColor: "#64748b", StrokeWidth: 1, Opacity: &opacity,
			StartBinding: &entity.Binding{ElementID: "src-item", FixedPoint: []float64{1, 0.5}},
			EndBinding:   &entity.Binding{ElementID: "dst-item", FixedPoint: []float64{0, 0.5}},
			CustomData:   &entity.CustomData{ConnectorKind: "route"}},
		{ID: "traffic", Type: "arrow", StrokeColor: "#2563eb", StrokeWidth: 2, StrokeStyle: "dotted", Opacity: &opacity,
			StartBinding: &entity.Binding{ElementID: "src-item", FixedPoint: []float64{1, 0.5}},
			EndBinding:   &entity.Binding{ElementID: "dst-item", FixedPoint: []float64{0, 0.5}},
			CustomData:   &entity.CustomData{ConnectorKind: "traffic", ConnectorStartArrowhead: "oval", ConnectorEndArrowhead: "diamond"}},
	}, AppState: &entity.AppState{ViewBackgroundColor: "#FFFFFF"}}

	plan := usecase.BuildPlan(&scene, entity.PptxOptions{PxPerInch: 96, ArrowStyle: "thin"})
	if len(plan.ConnectorLegend) != 2 {
		t.Fatalf("connector legend = %#v", plan.ConnectorLegend)
	}
	traffic := plan.ConnectorLegend[1]
	if traffic.Kind != "traffic" || traffic.Source != "src-item" || traffic.Target != "dst-item" || traffic.Line.Color != "2563EB" || traffic.Line.Dash != "dot" || traffic.Line.BeginArrowType != "oval" || traffic.Line.EndArrowType != "diamond" {
		t.Fatalf("traffic legend = %#v", traffic)
	}
}

func TestBuildPlanPaperMarginsInsetFittedContent(t *testing.T) {
	scene := entity.PptxScene{Elements: []entity.Element{
		{ID: "paper-frame", Type: "frame", Width: 100, Height: 50},
		{ID: "box", Type: "rectangle", Width: 100, Height: 50, StrokeColor: "#000000", BackgroundColor: "#ffffff"},
	}}
	plan := usecase.BuildPlan(&scene, entity.PptxOptions{
		PxPerInch: 100, PaperSize: "A4", Orientation: "portrait",
		PaperMarginTop: 2, PaperMarginRight: 1, PaperMarginBottom: 2, PaperMarginLeft: 1,
	})
	if plan.Slide.W != 8.27 || plan.Slide.H != 11.69 {
		t.Fatalf("slide = %.2fx%.2f, want A4 portrait", plan.Slide.W, plan.Slide.H)
	}
	var box *entity.DrawOp
	for i := range plan.Ops {
		if plan.Ops[i].Kind == "rect" && plan.Ops[i].W > 0 && plan.Ops[i].H > 0 {
			box = &plan.Ops[i]
			break
		}
	}
	if box == nil {
		t.Fatal("box op was not generated")
	}
	if box.X < 0.99 || box.Y < 1.99 {
		t.Fatalf("box was not inset by paper margins: %#v", box)
	}
}

func TestBuildPlanDrawsHeaderTagAboveConnectors(t *testing.T) {
	opacity := 100.0
	scene := entity.PptxScene{Elements: []entity.Element{
		{ID: "paper-frame", Type: "frame", Width: 200, Height: 100},
		{ID: "group-rect", Type: "rectangle", X: 0, Y: 0, Width: 160, Height: 80, StrokeColor: "#00A1C9", BackgroundColor: "transparent", Opacity: &opacity, CustomData: &entity.CustomData{GroupBorder: true}},
		{ID: "raw-line", Type: "line", X: 0, Y: 20, Width: 160, Height: 0, StrokeColor: "#334155", Opacity: &opacity, Points: [][]float64{{0, 0}, {160, 0}}},
		{ID: "header-bg", Type: "line", X: 0, Y: 8, Width: 80, Height: 24, StrokeColor: "#00A1C9", BackgroundColor: "#FFFFFF", Opacity: &opacity, Points: [][]float64{{0, 0}, {68, 0}, {80, 12}, {68, 24}, {0, 24}, {0, 0}}, CustomData: &entity.CustomData{GroupHeader: true}},
	}}
	plan := usecase.BuildPlan(&scene, entity.PptxOptions{PxPerInch: 96})
	lineIndex, headerIndex := -1, -1
	for i, op := range plan.Ops {
		if op.ID == "raw-line" && op.Kind == "line" {
			lineIndex = i
		}
		if op.Kind == "polygon" {
			headerIndex = i
		}
	}
	if lineIndex < 0 || headerIndex < 0 || headerIndex < lineIndex {
		t.Fatalf("draw order line=%d header=%d ops=%#v", lineIndex, headerIndex, plan.Ops)
	}
}

func TestBuildPlanJSONEncodesPlan(t *testing.T) {
	sceneJSON := `{"elements":[{"id":"paper-frame","type":"frame","width":200,"height":100}],"appState":{"viewBackgroundColor":"#ffffff"},"files":{}}`
	out, err := usecase.BuildPlanJSON(sceneJSON, entity.PptxOptions{PxPerInch: 100})
	if err != nil {
		t.Fatal(err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(out, &plan); err != nil {
		t.Fatal(err)
	}
	if plan.Slide.W != 2 || plan.Slide.H != 1 || plan.Slide.Background != "FFFFFF" {
		t.Fatalf("plan slide = %#v", plan.Slide)
	}
}
