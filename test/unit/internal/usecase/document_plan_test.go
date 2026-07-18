package usecase_test

import (
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestDocumentPlanUsesOnePagePerFrameByDefault(t *testing.T) {
	scene := documentPlanScene(t)
	document := v1engine.BuildDocumentPlanV1EnginePlanDocument(&scene, entity.PlanOptions{PxPerInch: 100}, false)
	if document.SchemaVersion != 2 {
		t.Fatalf("SchemaVersion = %d, want 2", document.SchemaVersion)
	}
	if len(document.Pages) != 2 {
		t.Fatalf("pages = %d, want 2", len(document.Pages))
	}
	if document.Pages[0].ID != "source" || document.Pages[1].ID != "destination" {
		t.Fatalf("page IDs = %q, %q", document.Pages[0].ID, document.Pages[1].ID)
	}
	assertDocumentPageRectCount(t, document.Pages[0], 1)
	assertDocumentPageRectCount(t, document.Pages[1], 1)
	if document.Pages[0].Slide.W != 2 || document.Pages[1].Slide.W != 1.5 {
		t.Fatalf("page widths = %v, %v", document.Pages[0].Slide.W, document.Pages[1].Slide.W)
	}
	if !document.Pages[0].Slide.CropToSlide || !document.Pages[1].Slide.CropToSlide {
		t.Fatalf("page slides must crop to their logical frame: %#v", document.Pages)
	}
}

func TestDocumentPlanCombineFramesRetainsSingleCanvas(t *testing.T) {
	scene := documentPlanScene(t)
	document := v1engine.BuildDocumentPlanV1EnginePlanDocument(&scene, entity.PlanOptions{PxPerInch: 100}, true)
	if len(document.Pages) != 1 || document.Pages[0].ID != "combined" {
		t.Fatalf("combined pages = %#v", document.Pages)
	}
	if document.Pages[0].Slide.CropToSlide {
		t.Fatalf("combined compatibility canvas must retain marker-safe overflow: %#v", document.Pages[0].Slide)
	}
	assertDocumentPageRectCount(t, document.Pages[0], 2)
}

func TestDocumentPlanLegacySingleFrameCropsToLogicalSlideByDefault(t *testing.T) {
	scene := entity.PresentationScene{Elements: []entity.Element{
		{ID: "paper-frame", Type: "frame", Width: 200, Height: 100},
	}}
	document := v1engine.BuildDocumentPlanV1EnginePlanDocument(&scene, entity.PlanOptions{PxPerInch: 100}, false)
	if len(document.Pages) != 1 || !document.Pages[0].Slide.CropToSlide {
		t.Fatalf("legacy single-frame page must crop to its logical slide: %#v", document.Pages)
	}
}

func TestNormalizeDocumentPageSizesCentresSmallerPages(t *testing.T) {
	document := entity.DocumentPlan{Pages: []entity.DocumentPage{
		{ID: "small", Slide: entity.PlanSlide{W: 2, H: 1}, Ops: []entity.DrawOp{{ID: "shape", X: 0.25, Y: 0.25}}},
		{ID: "large", Slide: entity.PlanSlide{W: 4, H: 3}},
	}}
	v1engine.NormalizeDocumentPageSizesV1EnginePlanDocument(&document)
	if document.Pages[0].Slide.W != 4 || document.Pages[0].Slide.H != 3 {
		t.Fatalf("small page slide = %#v", document.Pages[0].Slide)
	}
	if document.Pages[0].Ops[0].X != 1.25 || document.Pages[0].Ops[0].Y != 1.25 {
		t.Fatalf("centred op = %#v", document.Pages[0].Ops[0])
	}
}

func documentPlanScene(t *testing.T) entity.PresentationScene {
	t.Helper()
	return entity.PresentationScene{
		AppState: &entity.AppState{ViewBackgroundColor: "#ffffff"},
		Files:    map[string]entity.SceneFile{},
		Elements: []entity.Element{
			{ID: "paper-frame", Type: "frame", X: 0, Y: 0, Width: 500, Height: 200},
			{ID: "paper-frame-source", Type: "frame", X: 0, Y: 0, Width: 200, Height: 100, StrokeColor: "#bbbbbb", CustomData: &entity.CustomData{PageFrame: true, FrameID: "source"}},
			{ID: "source-rect", Type: "rectangle", X: 20, Y: 20, Width: 40, Height: 30, StrokeColor: "#000000", BackgroundColor: "#ffffff"},
			{ID: "paper-frame-destination", Type: "frame", X: 300, Y: 0, Width: 150, Height: 120, StrokeColor: "#bbbbbb", CustomData: &entity.CustomData{PageFrame: true, FrameID: "destination"}},
			{ID: "destination-rect", Type: "rectangle", X: 320, Y: 20, Width: 40, Height: 30, StrokeColor: "#000000", BackgroundColor: "#ffffff"},
		},
	}
}

func assertDocumentPageRectCount(t *testing.T, page entity.DocumentPage, want int) {
	t.Helper()
	count := 0
	for _, op := range page.Ops {
		if op.Kind == "rect" {
			count++
		}
	}
	if count != want {
		t.Fatalf("page %q rect count=%d want=%d ops=%#v", page.ID, count, want, page.Ops)
	}
}
