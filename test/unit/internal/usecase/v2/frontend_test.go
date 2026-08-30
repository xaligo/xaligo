package v2_test

import (
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestFrontendLowersV1AndV2ToOneGenericContract(t *testing.T) {
	cases := []struct {
		version string
		source  string
	}{
		{version: "1", source: `<xaligo version="1"><frames><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><connection id="flow" source="left" target="right" routing="orthogonal"/></frame></frames></xaligo>`},
		{version: "2", source: `<scene version="2" width="320" height="180" layout="absolute"><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><line id="flow" source="left" target="right" routing="orthogonal"/></frame></scene>`},
	}
	for _, test := range cases {
		source := []byte(test.source)
		spec, gotVersion, err := v2.NewFrontendUsecase().Lower(source)
		if err != nil {
			t.Fatalf("version %s: %v", test.version, err)
		}
		if gotVersion != test.version || spec.Width != 320 || spec.Height != 180 || len(spec.Elements) != 1 {
			t.Fatalf("version %s spec = %#v", test.version, spec)
		}
		frame := spec.Elements[0]
		if frame.Concept != entity.EngineConceptFrame || len(frame.Children) != 3 || frame.Children[2].Concept != entity.EngineConceptLine {
			t.Fatalf("version %s frame = %#v", test.version, frame)
		}
		if frame.SpanID == 0 || len(spec.Spans) != 4 || spec.Spans[0].Line < 1 {
			t.Fatalf("version %s spans = %#v", test.version, spec.Spans)
		}
	}
}

func TestFrontendRejectsContradictoryDocumentRootAndVersion(t *testing.T) {
	for _, source := range []string{
		`<xaligo version="2"><frame id="page"/></xaligo>`,
		`<scene version="1"><item id="node"/></scene>`,
		`<scene><item id="node"/></scene>`,
	} {
		if _, _, err := v2.NewFrontendUsecase().Lower([]byte(source)); err == nil {
			t.Fatalf("Lower(%q) error = nil", source)
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

func TestResolvedDocumentPlanUsesResolvedIconAndTextGeometry(t *testing.T) {
	document := entity.EngineResolvedDocument{Width: 192, Height: 96, Elements: []entity.EngineResolvedElement{{
		ID: "node", Concept: entity.EngineConceptItem, X: 20, Y: 10, Width: 72, Height: 52,
		Visual: entity.EngineResolvedVisual{Shape: entity.EngineShapeNone, Visible: true},
		Text: entity.EngineResolvedText{
			Value: "Node", FontSize: 10, LineHeight: 1.2,
			X: 20, Y: 50, Width: 72, Height: 12,
		},
		IconRef: "catalog:27", IconX: 40, IconY: 14, IconWidth: 32, IconHeight: 32,
	}}}
	plan, err := v2.BuildDocumentPlanWithIcons(document, 96, map[string]string{"catalog:27": "data:image/svg+xml;base64,PHN2Zy8+"})
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Pages) != 1 || len(plan.Pages[0].Ops) != 2 {
		t.Fatalf("plan = %#v", plan)
	}
	icon, label := plan.Pages[0].Ops[0], plan.Pages[0].Ops[1]
	if icon.Kind != "image" || icon.X != 40.0/96.0 || icon.Y != 14.0/96.0 || icon.W != 32.0/96.0 || !icon.FrontLayer {
		t.Fatalf("icon op = %#v", icon)
	}
	if label.Kind != "text" || label.Y != 50.0/96.0 || label.H != 12.0/96.0 || !label.FrontLayer {
		t.Fatalf("text op = %#v", label)
	}
}

func TestFrontendNormalizesV1ProfileAttributesForV2(t *testing.T) {
	source := []byte(`<xaligo version="1"><frame id="page" width="640" height="360"><row><col span="3"><generic-group id="group" title="Services" icon-id="100" row="1"><item id="200" name="API"/></generic-group></col><col span="9"><rectangle id="target" title="Target"><port id="in-a" side="left" title="A"/><port id="in-b" side="left" title="B"/></rectangle></col></row><connections><connection id="flow" src="200" dst="in-a" kind="traffic" color="#2563eb" stroke-style="dashed" arrow="triangle"/></connections></frame></xaligo>`)
	spec, _, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	frame := spec.Elements[0]
	row := frame.Children[0]
	left, right := row.Children[0], row.Children[1]
	group := left.Children[0]
	item := group.Children[0]
	ports := right.Children[0].Children
	line := frame.Children[1]
	if row.Layout != entity.EngineLayoutHorizontal || left.Weight == nil || *left.Weight != 3 || right.Weight == nil || *right.Weight != 9 {
		t.Fatalf("V1 column layout was not normalized: %#v", row)
	}
	if group.Layout != entity.EngineLayoutVertical || group.Weight == nil || *group.Weight != 1 || group.Text == nil || group.Text.Value != "Services" || group.Icon == nil || group.Icon.Ref != "catalog:100" {
		t.Fatalf("V1 group was not normalized: %#v", group)
	}
	if item.Text == nil || item.Text.Value != "API" || item.Icon == nil || item.Icon.Ref != "catalog:200" || len(ports) != 2 || ports[0].Port.Label != "A" || ports[1].Port.Label != "B" {
		t.Fatalf("V1 item/port was not normalized: item=%#v ports=%#v", item, ports)
	}
	if ports[0].Port.Anchor == nil || ports[1].Port.Anchor == nil || *ports[0].Port.Anchor != 1.0/3.0 || *ports[1].Port.Anchor != 2.0/3.0 {
		t.Fatalf("V1 ports were not distributed on their side: %#v", ports)
	}
	if line.Line.Routing != entity.EngineRoutingOrthogonal || line.Visual.Stroke != "#2563eb" || line.Line.Style != entity.EngineLineDashed || line.Line.TargetDecoration != entity.EngineDecorationTriangle {
		t.Fatalf("V1 connection was not normalized: %#v", line)
	}
	if line.Concept != entity.EngineConceptLine || len(frame.Children) != 2 {
		t.Fatalf("V1 connections wrapper was not flattened: %#v", frame.Children)
	}
}

func TestFrontendLowersNativeV2ParametersWithoutV1Aliases(t *testing.T) {
	source := []byte(`<scene version="2" width="640" height="360" layout="vertical" padding="24"><group id="services" weight="1" layout="horizontal" align="center" justify="space-evenly" padding-top="44" padding-x="12" padding-bottom="12" shape="none"><item id="api" label="API" icon-ref="catalog:200" icon-width="40" icon-height="40"/><item id="worker" name="endpoint-only" icon-ref="catalog:201"><port id="worker-in" side="left"/></item></group><line id="flow" source="api" target="worker-in" routing="orthogonal" source-side="right" target-side="left" source-anchor="0.3" target-anchor="0.7" source-decoration="circle" target-decoration="arrow"/></scene>`)
	spec, version, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	if version != "2" || spec.Layout != entity.EngineLayoutVertical || spec.Padding.Top == nil || *spec.Padding.Top != 24 {
		t.Fatalf("native document = %#v", spec)
	}
	group := spec.Elements[0]
	if group.Align != entity.EngineAlignCenter || group.Justify != entity.EngineJustifySpaceEvenly || group.Padding.Top == nil || *group.Padding.Top != 44 || group.Padding.Left == nil || *group.Padding.Left != 12 || group.Visual.Shape != entity.EngineShapeNone {
		t.Fatalf("native group = %#v", group)
	}
	if group.Children[0].Text == nil || group.Children[0].Text.Value != "API" || group.Children[0].Icon.Width == nil || *group.Children[0].Icon.Width != 40 {
		t.Fatalf("native item = %#v", group.Children[0])
	}
	if group.Children[1].Text != nil {
		t.Fatalf("name must remain identity-only in V2: %#v", group.Children[1].Text)
	}
	if port := group.Children[1].Children[0].Port; port == nil || port.Anchor != nil {
		t.Fatalf("native V2 port default must remain unresolved for the engine: %#v", port)
	}
	line := spec.Elements[1].Line
	if line.SourceSide != entity.EngineSideRight || line.TargetSide != entity.EngineSideLeft || line.SourceAnchor == nil || *line.SourceAnchor != 0.3 || line.TargetAnchor == nil || *line.TargetAnchor != 0.7 || line.SourceDecoration != entity.EngineDecorationCircle || line.TargetDecoration != entity.EngineDecorationArrow {
		t.Fatalf("native line = %#v", line)
	}
}
