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

func TestFrontendNormalizesV1ProfileAttributesForV2(t *testing.T) {
	source := []byte(`<xaligo version="2"><frame id="page" width="640" height="360"><row><col span="3"><generic-group id="group" title="Services" icon-id="100"><item id="200" name="API"/></generic-group></col><col span="9"><rectangle id="target" title="Target"><port id="in" side="left" title="IN"/></rectangle></col></row><connections><connection id="flow" src="200" dst="in" color="#2563eb" stroke-style="dashed" arrow="triangle"/></connections></frame></xaligo>`)
	spec, _, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	frame := spec.Elements[0]
	row := frame.Children[0]
	left, right := row.Children[0], row.Children[1]
	group := left.Children[0]
	item := group.Children[0]
	port := right.Children[0].Children[0]
	line := frame.Children[1]
	if row.Layout != entity.EngineLayoutHorizontal || left.Weight == nil || *left.Weight != 3 || right.Weight == nil || *right.Weight != 9 {
		t.Fatalf("V1 column layout was not normalized: %#v", row)
	}
	if group.Layout != entity.EngineLayoutVertical || group.Text == nil || group.Text.Value != "Services" || group.Icon == nil || group.Icon.Ref != "catalog:100" {
		t.Fatalf("V1 group was not normalized: %#v", group)
	}
	if item.Text == nil || item.Text.Value != "API" || item.Icon == nil || item.Icon.Ref != "catalog:200" || port.Port.Label != "IN" {
		t.Fatalf("V1 item/port was not normalized: item=%#v port=%#v", item, port)
	}
	if line.Line.Routing != entity.EngineRoutingOrthogonal || line.Visual.Stroke != "#2563eb" || line.Line.Style != entity.EngineLineDashed || line.Line.TargetDecoration != entity.EngineDecorationTriangle {
		t.Fatalf("V1 connection was not normalized: %#v", line)
	}
	if line.Concept != entity.EngineConceptLine || len(frame.Children) != 2 {
		t.Fatalf("V1 connections wrapper was not flattened: %#v", frame.Children)
	}
}
