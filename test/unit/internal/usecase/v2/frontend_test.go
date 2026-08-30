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
		{version: "2", source: `<xaligo version="2"><frames><frame id="page" width="320" height="180" layout="horizontal"><item id="left" width="80">Left</item><item id="right" weight="1">Right</item><line id="flow" source="left" target="right" routing="orthogonal"/></frame></frames></xaligo>`},
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
		`<xaligo version="3"><frame id="page"/></xaligo>`,
		`<scene version="2"><item id="node"/></scene>`,
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
	if label.FontSize != 7.5 {
		t.Fatalf("text font size = %v, want 7.5pt", label.FontSize)
	}
}

func TestResolvedDocumentPlanProjectsV1GroupHeaderGeometry(t *testing.T) {
	document := entity.EngineResolvedDocument{Width: 200, Height: 160, Elements: []entity.EngineResolvedElement{{
		ID: "cloud", Concept: entity.EngineConceptGroup, X: 24, Y: 31, Width: 160, Height: 100,
		Visual: entity.EngineResolvedVisual{Shape: entity.EngineShapeRectangle, Fill: "#ffffff", Stroke: "#000000", StrokeWidth: 2, Opacity: 1, Visible: true},
		Text: entity.EngineResolvedText{
			Value: "AWS Cloud", Role: string(entity.TextRoleGroupHeader), FontFamily: "Helvetica", Color: "#000000", FontSize: 14, LineHeight: 1.25,
			X: 58, Y: 38, Width: 95, Height: 18,
		},
		IconRef: "group:AWS-Cloud-logo_32.svg", IconX: 22, IconY: 31, IconWidth: 32, IconHeight: 32,
	}}}
	plan, err := v2.BuildDocumentPlanWithIcons(document, 96, map[string]string{"group:AWS-Cloud-logo_32.svg": "data:image/svg+xml;base64,PHN2Zy8+"})
	if err != nil {
		t.Fatal(err)
	}
	ops := plan.Pages[0].Ops
	if len(ops) != 4 {
		t.Fatalf("ops = %#v", ops)
	}
	if border := ops[0]; border.Kind != "rect" || border.Y != 47.0/96.0 || border.H != 84.0/96.0 {
		t.Fatalf("border op = %#v", border)
	}
	if header := ops[1]; header.Kind != "polygon" || !header.FrontLayer || header.GroupID != "" || header.X != 22.0/96.0 || header.Y != 31.0/96.0 || len(header.Points) != 5 {
		t.Fatalf("header op = %#v", header)
	}
	if label := ops[3]; label.Align != "left" || label.Valign != "mid" || label.FontSize != 10.5 || label.TextLayout == nil || label.TextLayout.Wrap {
		t.Fatalf("header label op = %#v", label)
	}
}

func TestFrontendNormalizesV1ProfileAttributesForV2(t *testing.T) {
	source := []byte(`<xaligo version="2"><frame id="page" width="640" height="360"><row><col span="3"><generic-group id="group" title="Services" icon-id="100" row="1"><item id="200" name="API"/></generic-group></col><col span="9"><rectangle id="target" title="Target"><port id="in-a" side="left" title="A"/><port id="in-b" side="left" title="B"/></rectangle></col></row><connections><connection id="flow" src="200" dst="in-a" kind="traffic" color="#2563eb" stroke-style="dashed" arrow="triangle"/></connections></frame></xaligo>`)
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
	if group.Layout != entity.EngineLayoutAdaptiveGrid || group.Weight == nil || *group.Weight != 1 || group.Text == nil || group.Text.Value != "Services" || group.Icon == nil || group.Icon.Ref != "catalog:100" {
		t.Fatalf("V1 group was not normalized: %#v", group)
	}
	if group.Padding.Top == nil || *group.Padding.Top != 26 || group.Padding.Bottom == nil || *group.Padding.Bottom != 8 || group.Padding.Left == nil || *group.Padding.Left != 12 || group.Gap == nil || *group.Gap != 8 || group.Visual.Stroke != "#AAB7B8" || group.Visual.CornerRadius == nil || *group.Visual.CornerRadius != 0 || group.Line == nil || group.Line.Style != entity.EngineLineDashed {
		t.Fatalf("V1 group defaults were not normalized: %#v", group)
	}
	if group.Text.Role != string(entity.TextRoleGroupHeader) || group.Text.FontFamily != "Helvetica" || group.Text.FontSize == nil || *group.Text.FontSize != 14 || group.Text.LineHeight == nil || *group.Text.LineHeight != 1.25 {
		t.Fatalf("V1 group text profile was not normalized: %#v", group.Text)
	}
	if item.Text == nil || item.Text.Value != "API" || item.Icon == nil || item.Icon.Ref != "catalog:200" || len(ports) != 2 || ports[0].Port.Label != "A" || ports[1].Port.Label != "B" {
		t.Fatalf("V1 item/port was not normalized: item=%#v ports=%#v", item, ports)
	}
	if ports[0].Port.Anchor == nil || ports[1].Port.Anchor == nil || *ports[0].Port.Anchor != 1.0/3.0 || *ports[1].Port.Anchor != 2.0/3.0 {
		t.Fatalf("V1 ports were not distributed on their side: %#v", ports)
	}
	if ports[0].Width == nil || *ports[0].Width != 56 || ports[0].Height == nil || *ports[0].Height != 22 {
		t.Fatalf("V1 port dimensions were not defaulted: %#v", ports[0])
	}
	if line.Line.Routing != entity.EngineRoutingOrthogonal || line.Visual.Stroke != "#2563eb" || line.Line.Style != entity.EngineLineDashed || line.Line.TargetDecoration != entity.EngineDecorationTriangle {
		t.Fatalf("V1 connection was not normalized: %#v", line)
	}
	if line.Concept != entity.EngineConceptLine || len(frame.Children) != 2 {
		t.Fatalf("V1 connections wrapper was not flattened: %#v", frame.Children)
	}
}

func TestFrontendAppliesCatalogLabelsWithoutReparsingSource(t *testing.T) {
	source := []byte(`<xaligo version="2"><frames><frame id="page" width="320" height="180" item-size="32"><generic-group id="services" title="Services"><item id="200" name="fallback"/><item id="201" label="Explicit"/><item id="202" name="fallback"/></generic-group></frame></frames></xaligo>`)
	spec, _, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	v2.ApplyCatalogLabels(&spec, map[int]string{200: "network_user", 201: "catalog value", 202: "CloudWatch Event Event Based"})
	items := spec.Elements[0].Children[0].Children
	if items[0].Text == nil || items[0].Text.Value != "network_u\nser" || items[0].Height == nil || *items[0].Height != 63 || items[0].Width == nil || *items[0].Width != 56 {
		t.Fatalf("catalog fallback = %#v", items[0])
	}
	if items[1].Text == nil || items[1].Text.Value != "Explicit" {
		t.Fatalf("explicit label was overwritten: %#v", items[1])
	}
	if items[2].Text == nil || items[2].Text.Value != "CloudWatc\nh Event\nEvent\nBased" || items[2].Height == nil || *items[2].Height != 90 {
		t.Fatalf("long catalog label = %#v", items[2])
	}
}

func TestFrontendComposesV1FrameMetadataWithoutChangingContentBounds(t *testing.T) {
	source := []byte(`<xaligo version="2"><frames><frame id="page" title="Architecture" version="2026.08" width="640" height="360" class="pa-3"><row><item id="node" width="40" height="40"/></row></frame></frames></xaligo>`)
	spec, _, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	frame := spec.Elements[0]
	if frame.Text != nil || frame.Layout != entity.EngineLayoutVertical || frame.Padding.Top == nil || *frame.Padding.Top != 4 || frame.Padding.Bottom == nil || *frame.Padding.Bottom != 24 || frame.Gap == nil || *frame.Gap != 8 || len(frame.Children) != 2 {
		t.Fatalf("metadata frame = %#v", frame)
	}
	metadata, content := frame.Children[0], frame.Children[1]
	if metadata.Concept != entity.EngineConceptGroup || metadata.Visual.Visible == nil || *metadata.Visual.Visible || metadata.Height == nil || *metadata.Height != 19 || len(metadata.Children) != 6 {
		t.Fatalf("metadata row = %#v", metadata)
	}
	if metadata.Children[0].Text == nil || metadata.Children[0].Text.Value != "id" || metadata.Children[1].Text == nil || metadata.Children[1].Text.Value != "page" || metadata.Children[5].Text == nil || metadata.Children[5].Text.Value != "2026.08" {
		t.Fatalf("metadata cells = %#v", metadata.Children)
	}
	if content.Concept != entity.EngineConceptGroup || content.Weight == nil || *content.Weight != 1 || content.Margin.Left == nil || *content.Margin.Left != 24 || content.Margin.Right == nil || *content.Margin.Right != 24 || len(content.Children) != 1 || content.Children[0].ID == "" {
		t.Fatalf("metadata content wrapper = %#v", content)
	}
}

func TestFrontendExplicitValuesOverrideV1ProfileDefaults(t *testing.T) {
	source := []byte(`<xaligo version="2"><frames><frame id="page"><region id="custom" title="Custom" padding="0" fill="#abcdef" stroke="#123456" stroke-width="4" stroke-style="solid" icon-width="48"><blank/></region></frame></frames></xaligo>`)
	spec, _, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	region := spec.Elements[0].Children[0]
	if region.Padding.Top == nil || *region.Padding.Top != 0 || region.Padding.Right == nil || *region.Padding.Right != 0 {
		t.Fatalf("explicit zero padding was not preserved: %#v", region.Padding)
	}
	if region.Visual.Fill != "#abcdef" || region.Visual.Stroke != "#123456" || region.Visual.StrokeWidth == nil || *region.Visual.StrokeWidth != 4 || region.Line != nil {
		t.Fatalf("explicit visual values did not override the profile: %#v", region)
	}
	if region.Icon == nil || region.Icon.Ref != "group:Region_32.svg" || region.Icon.Width == nil || *region.Icon.Width != 48 || region.Icon.Height == nil || *region.Icon.Height != 32 {
		t.Fatalf("explicit icon width or profile icon default was lost: %#v", region.Icon)
	}
}

func TestFrontendV2RejectsMultipleFramesUntilDocumentLayoutSupportsPages(t *testing.T) {
	for _, source := range []string{
		`<xaligo version="2"><frames><frame id="one"/><frame id="two"/></frames></xaligo>`,
		`<xaligo version="2"><frame id="one"/><frame id="two"/></xaligo>`,
	} {
		if _, _, err := v2.NewFrontendUsecase().Lower([]byte(source)); err == nil {
			t.Fatalf("multiple V2 content roots accepted: %s", source)
		}
	}
}

func TestFrontendLowersNativeV2ParametersWithoutV1Aliases(t *testing.T) {
	source := []byte(`<xaligo version="2"><frames><frame id="page" width="640" height="360" layout="vertical" padding="24"><group id="services" weight="1" layout="horizontal" align="center" justify="space-evenly" padding-top="44" padding-x="12" padding-bottom="12" shape="none"><item id="api" label="API" icon-ref="catalog:200" icon-width="40" icon-height="40"/><item id="worker" name="endpoint-only" icon-ref="catalog:201"><port id="worker-in" side="left"/></item></group><line id="flow" source="api" target="worker-in" routing="orthogonal" source-side="right" target-side="left" source-anchor="0.3" target-anchor="0.7" source-decoration="circle" target-decoration="arrow"/></frame></frames></xaligo>`)
	spec, version, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	if version != "2" || spec.Layout != entity.EngineLayoutVertical || len(spec.Elements) != 1 {
		t.Fatalf("native document = %#v", spec)
	}
	frame := spec.Elements[0]
	if frame.Padding.Top == nil || *frame.Padding.Top != 24 {
		t.Fatalf("native frame = %#v", frame)
	}
	group := frame.Children[0]
	if group.Align != entity.EngineAlignCenter || group.Justify != entity.EngineJustifySpaceEvenly || group.Padding.Top == nil || *group.Padding.Top != 44 || group.Padding.Left == nil || *group.Padding.Left != 12 || group.Visual.Shape != entity.EngineShapeNone {
		t.Fatalf("native group = %#v", group)
	}
	if group.Children[0].Text == nil || group.Children[0].Text.Value != "API" || group.Children[0].Icon.Width == nil || *group.Children[0].Icon.Width != 40 {
		t.Fatalf("native item = %#v", group.Children[0])
	}
	if group.Children[1].Text != nil {
		t.Fatalf("name must remain identity-only in V2: %#v", group.Children[1].Text)
	}
	if port := group.Children[1].Children[0].Port; port == nil || port.Anchor == nil || *port.Anchor != 0.5 {
		t.Fatalf("V2 authoring profile port default was not normalized: %#v", port)
	}
	line := frame.Children[1].Line
	if line.SourceSide != entity.EngineSideRight || line.TargetSide != entity.EngineSideLeft || line.SourceAnchor == nil || *line.SourceAnchor != 0.3 || line.TargetAnchor == nil || *line.TargetAnchor != 0.7 || line.SourceDecoration != entity.EngineDecorationCircle || line.TargetDecoration != entity.EngineDecorationArrow {
		t.Fatalf("native line = %#v", line)
	}
}
