package v2_test

import (
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestFrontendLowersV1AndV2ToOneGenericContract(t *testing.T) {
	for _, version := range []string{"1", "2"} {
		source := []byte(`<xaligo version="` + version + `"><frames><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><connection id="flow" source="left" target="right" routing="orthogonal"/></frame></frames></xaligo>`)
		spec, gotVersion, err := v2.NewFrontendUsecase().Lower(source)
		if err != nil {
			t.Fatalf("version %s: %v", version, err)
		}
		if gotVersion != version || spec.Width != 320 || spec.Height != 180 || len(spec.Elements) != 1 {
			t.Fatalf("version %s spec = %#v", version, spec)
		}
		frame := spec.Elements[0]
		if frame.Concept != entity.EngineConceptFrame || len(frame.Children) != 3 || frame.Children[2].Concept != entity.EngineConceptLine {
			t.Fatalf("version %s frame = %#v", version, frame)
		}
		if frame.SpanID == 0 || len(spec.Spans) != 4 || spec.Spans[0].Line < 1 {
			t.Fatalf("version %s spans = %#v", version, spec.Spans)
		}
	}
}

func TestResolvedDocumentBuildsRendererNeutralPlan(t *testing.T) {
	document := entity.EngineResolvedDocument{Width: 192, Height: 96, Elements: []entity.EngineResolvedElement{{
		ID: "node", Concept: entity.EngineConceptItem, X: 0, Y: 0, Width: 96, Height: 48,
		Visual: entity.EngineResolvedVisual{Shape: entity.EngineShapeRectangle, Fill: "#ffffff", Stroke: "#000000", StrokeWidth: 1, Opacity: 1, Visible: true},
		Text:   entity.EngineResolvedText{Value: "Node", Color: "#111111", FontSize: 12, LineHeight: 1.2},
	}}}
	plan, err := v2.BuildDocumentPlan(document, 96)
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Pages) != 1 || plan.Pages[0].Slide.W != 2 || len(plan.Pages[0].Ops) != 2 {
		t.Fatalf("plan = %#v", plan)
	}
}
