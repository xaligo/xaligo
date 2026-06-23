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

func TestBuildPlanResolvesConnectorArrowStyles(t *testing.T) {
	opacity := 100.0
	scene := entity.PptxScene{Elements: []entity.Element{
		{ID: "src", Type: "image", X: 0, Y: 0, Width: 32, Height: 32, Opacity: &opacity, FileID: "src-file"},
		{ID: "dst", Type: "image", X: 160, Y: 0, Width: 32, Height: 32, Opacity: &opacity, FileID: "dst-file"},
		{ID: "connector", Type: "arrow", StrokeColor: "#64748b", StrokeWidth: 3, Opacity: &opacity,
			StartBinding: &entity.Binding{ElementID: "src", FixedPoint: []float64{1, 0.5}},
			EndBinding:   &entity.Binding{ElementID: "dst", FixedPoint: []float64{0, 0.5}},
			CustomData:   &entity.CustomData{ConnectorKind: "route"}},
	}, AppState: &entity.AppState{ViewBackgroundColor: "#FFFFFF"}}

	cases := []struct {
		style string
		head  string
		width float64
	}{
		{"standard", "triangle", 1.5},
		{"triangle", "triangle", 3},
		{"stealth", "stealth", 3},
		{"arrow", "arrow", 3},
		{"diamond", "diamond", 3},
		{"oval", "oval", 3},
		{"none", "none", 3},
		{"", "stealth", 1},
	}
	for _, tc := range cases {
		t.Run(tc.style, func(t *testing.T) {
			plan := usecase.BuildPlan(&scene, entity.PptxOptions{PxPerInch: 96, ArrowStyle: tc.style})
			if len(plan.ConnectorLegend) != 1 {
				t.Fatalf("connector legend = %#v", plan.ConnectorLegend)
			}
			line := plan.ConnectorLegend[0].Line
			if line.EndArrowType != tc.head || line.Width != tc.width {
				t.Fatalf("style %q line = %#v", tc.style, line)
			}
		})
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

func TestBuildPlanResolvesPaperVariantsAndLegendFilters(t *testing.T) {
	scene := entity.PptxScene{
		Elements: []entity.Element{{ID: "paper-frame", Type: "frame", Width: 300, Height: 100}},
		Files: map[string]entity.SceneFile{
			"item-cat-27": {DataURL: "data:image/svg+xml;base64,QQ=="},
		},
	}
	plan := usecase.BuildPlan(&scene, entity.PptxOptions{
		PxPerInch: 100,
		PaperSize: "A5",
		LegendEntries: []entity.LegendEntry{
			{CatalogID: 0, OfficialName: "Ignored"},
			{CatalogID: 27, OfficialName: "Amazon EC2"},
			{CatalogID: 27, OfficialName: "Duplicate"},
			{CatalogID: 999, OfficialName: "Missing File"},
		},
	})
	if plan.Slide.W != 8.27 || plan.Slide.H != 5.83 {
		t.Fatalf("auto paper slide = %.2fx%.2f", plan.Slide.W, plan.Slide.H)
	}
	if len(plan.Legend) != 1 || plan.Legend[0].CatalogID != 27 || plan.Legend[0].Abbreviation != "Amazon EC2" || plan.Legend[0].Data == "" {
		t.Fatalf("legend = %#v", plan.Legend)
	}

	invalidPaper := usecase.BuildPlan(&scene, entity.PptxOptions{PxPerInch: 100, PaperSize: "Nope"})
	if invalidPaper.Slide.W != 3 || invalidPaper.Slide.H != 1 {
		t.Fatalf("invalid paper fallback slide = %.2fx%.2f", invalidPaper.Slide.W, invalidPaper.Slide.H)
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

func TestBuildPlanConvertsStylesAndFallbacks(t *testing.T) {
	opacity := 40.0
	fontFamily := 1
	fontSize := 20.0
	scene := entity.PptxScene{
		Elements: []entity.Element{
			{ID: "paper-frame", Type: "frame", Width: 300, Height: 200},
			{ID: "ellipse", Type: "ellipse", X: 10, Y: 10, Width: 40, Height: 30, StrokeColor: "transparent", BackgroundColor: "transparent", StrokeWidth: 0, Opacity: &opacity},
			{ID: "poly", Type: "line", X: 60, Y: 10, Width: 40, Height: 40, StrokeColor: "not-a-color", BackgroundColor: "#abcdef", StrokeStyle: "dotted", Points: [][]float64{{0, 0}, {40, 0}, {20, 40}, {0, 0}}, Opacity: &opacity, CustomData: &entity.CustomData{GroupHeader: true}},
			{ID: "text", Type: "text", X: 10, Y: 60, Width: 120, Height: 30, RawText: "Fallback Text", StrokeColor: "bad", FontSize: &fontSize, FontFamily: &fontFamily, FontStyle: "bold", TextAlign: "right", VerticalAlign: "bottom"},
			{ID: "image", Type: "image", X: 160, Y: 60, Width: 40, Height: 40, FileID: "img", Angle: 15, Opacity: &opacity},
			{ID: "missing-image", Type: "image", X: 210, Y: 60, Width: 40, Height: 40, FileID: "missing"},
		},
		Files:    map[string]entity.SceneFile{"img": {DataURL: "data:image/svg+xml;base64,QQ=="}},
		AppState: &entity.AppState{ViewBackgroundColor: "not-a-color"},
	}
	plan := usecase.BuildPlan(&scene, entity.PptxOptions{PxPerInch: 100})
	if plan.Slide.Background != "FFFFFF" {
		t.Fatalf("slide background = %q", plan.Slide.Background)
	}
	var ellipse, polygon, text, image *entity.DrawOp
	for i := range plan.Ops {
		if plan.Ops[i].Kind == "ellipse" {
			ellipse = &plan.Ops[i]
		}
		if plan.Ops[i].Kind == "polygon" {
			polygon = &plan.Ops[i]
		}
		switch plan.Ops[i].ID {
		case "text":
			text = &plan.Ops[i]
		case "image":
			image = &plan.Ops[i]
		case "missing-image":
			t.Fatalf("missing image should not produce an op: %#v", plan.Ops[i])
		}
	}
	if ellipse == nil || ellipse.Kind != "ellipse" || ellipse.Line.Transparency != 100 || ellipse.Fill.Transparency != 100 {
		t.Fatalf("ellipse = %#v", ellipse)
	}
	if polygon == nil || polygon.Kind != "polygon" || polygon.Line.Color != "1E1E1E" || polygon.Line.Dash != "dot" || polygon.Fill.Color != "ABCDEF" {
		t.Fatalf("polygon = %#v", polygon)
	}
	if text == nil || text.Text != "Fallback Text" || text.FontFace != "Virgil" || text.Bold != true || text.Align != "right" || text.Valign != "bottom" || text.Color != "1E1E1E" {
		t.Fatalf("text = %#v", text)
	}
	if image == nil || image.Kind != "image" || image.Data == "" || image.Transparency != 60 || image.Rotate != 15 {
		t.Fatalf("image = %#v", image)
	}
}
