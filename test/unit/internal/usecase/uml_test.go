package usecase_test

import (
	"context"
	"encoding/json"
	"math"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
)

func TestUMLMetadataAndRelationLabelsReachSharedOutputs(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="640" height="360"><uml id="sequence"><sequence-diagram><participant id="user" title="User"/><lifeline id="api" title="API"/><message src="user" dst="api" order="1" title="submit()"/></sequence-diagram></uml></frame></frames></xaligo>`)
	options := entity.RenderOptions{PxPerInch: 96}
	scene, err := newUsecase().RenderExcalidraw(context.Background(), source, options)
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	for _, want := range []string{
		`"xaligoUmlDiagramKind": "sequence-diagram"`,
		`"xaligoUmlElementKind": "participant"`,
		`"xaligoUmlReference": "user"`,
		`"xaligoUmlRelationKind": "message"`,
		`"xaligoUmlRelationSourceReference": "user"`,
		`"xaligoUmlRelationDestinationReference": "api"`,
		`"xaligoUmlMessageOrder": "1"`,
		`"text": "1: submit()"`,
	} {
		if !strings.Contains(string(scene), want) {
			t.Fatalf("scene missing %q: %s", want, scene)
		}
	}
	svg, err := newUsecase().RenderSVG(context.Background(), source, options)
	if err != nil {
		t.Fatalf("RenderSVG() error = %v", err)
	}
	if !strings.Contains(string(svg), "1: submit()") {
		t.Fatalf("SVG missing UML relation label: %s", svg)
	}
}

func TestUMLShapeKindsReachEditableSceneAndSharedPlan(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="activity"><activity-diagram direction="right"><action id="before"/><decision id="choice"/><action id="yes"/><action id="no"/><control-flow src="before" dst="choice"/><control-flow src="choice" dst="yes" guard="yes"/><control-flow src="choice" dst="no" guard="no"/></activity-diagram></uml></frame></frames></xaligo>`)
	options := entity.RenderOptions{PxPerInch: 96}
	scene, err := newUsecase().RenderExcalidraw(context.Background(), source, options)
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	if !strings.Contains(string(scene), `"type": "diamond"`) {
		t.Fatalf("scene missing UML diamond: %s", scene)
	}
	plan, err := newUsecase().BuildPPTXPlan(context.Background(), source, options)
	if err != nil {
		t.Fatalf("BuildPPTXPlan() error = %v", err)
	}
	if !strings.Contains(string(plan), `"kind":"diamond"`) {
		t.Fatalf("plan missing UML diamond: %s", plan)
	}
}

func TestUMLActivitySwimlanesReachEditableScene(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="640"><uml id="activity"><activity-diagram direction="down" lanes="vertical" theme="xaligo"><partition id="customer" title="Customer"><initial id="start"/><action id="enter-pin" title="Enter PIN" tone="primary"/><final id="done"/></partition><partition id="atm" title="ATM"><action id="request-pin" title="Request PIN"/><decision id="pin-valid" title="PIN valid?"/></partition><control-flow src="start" dst="enter-pin"/><control-flow src="enter-pin" dst="request-pin"/><control-flow src="request-pin" dst="pin-valid"/><control-flow src="pin-valid" dst="request-pin" guard="invalid PIN" route="loop"/><control-flow src="pin-valid" dst="done" guard="valid"/></activity-diagram></uml></frame></frames></xaligo>`)
	scene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	for _, want := range []string{
		`"xaligoUmlPartition": true`,
		`"xaligoUmlPartitionHeader": true`,
		`"xaligoUmlPartitionId": "customer"`,
		`"xaligoUmlPartitionTitle": "Customer"`,
		`"xaligoUmlRoute": "loop"`,
		`"backgroundColor": "#08b8ea"`,
		`"strokeColor": "#052d6e"`,
	} {
		if !strings.Contains(string(scene), want) {
			t.Fatalf("scene missing %q: %s", want, scene)
		}
	}
}

func TestUMLActivityHidesRedundantContainerBorderAndTitle(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" title="Activity Process" width="760" height="420"><uml id="activity"><activity-diagram direction="down" lanes="vertical" theme="xaligo"><partition id="customer" title="Customer"><action id="one" title="One"/></partition><partition id="system" title="System"><action id="two" title="Two"/></partition><control-flow src="one" dst="two"/></activity-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene map[string]any
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	elements, _ := scene["elements"].([]any)
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if element["type"] == "rectangle" && customData["xaligoUmlDiagramKind"] == "activity-diagram" && customData["xaligoUmlElementKind"] == nil {
			t.Fatalf("activity container border should not render: %#v", element)
		}
		if element["type"] == "text" && element["text"] == "activity" {
			t.Fatalf("activity container title should not render when frame metadata carries the title: %#v", element)
		}
	}
}

func TestUMLActivityHorizontalSwimlanesReachEditableScene(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="760" height="420"><uml id="activity"><activity-diagram direction="right" lanes="horizontal" theme="xaligo"><partition id="customer" title="Customer"><initial id="start"/><action id="choose" title="Choose amount" tone="primary"/></partition><partition id="atm" title="ATM"><action id="read" title="Read card"/><final id="done"/></partition><control-flow src="start" dst="choose"/><control-flow src="choose" dst="read"/><control-flow src="read" dst="done"/></activity-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene map[string]any
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	elements, _ := scene["elements"].([]any)
	var header map[string]any
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlPartitionHeader"] == true && customData["xaligoUmlPartitionId"] == "customer" {
			header = element
			break
		}
	}
	if header == nil {
		t.Fatalf("horizontal partition header missing: %s", rawScene)
	}
	width, _ := header["width"].(float64)
	height, _ := header["height"].(float64)
	if !(height > width) {
		t.Fatalf("horizontal partition header = width %.1f height %.1f, want side header taller than wide", width, height)
	}
}

func TestUMLActivityElementsSupportCrossFramePageLinks(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames>
<frame id="overview" width="620" height="360"><uml id="start"><activity-diagram direction="right" lanes="horizontal"><partition id="actor" title="Actor"><action id="request" title="Request"/></partition></activity-diagram></uml><connection src="request" dst="detail.done" src-frame-side="right" dst-frame-side="left"/></frame>
<frame id="detail" width="620" height="360"><uml id="finish"><activity-diagram direction="right" lanes="horizontal"><partition id="system" title="System"><action id="done" title="Complete"/></partition></activity-diagram></uml></frame>
</frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene map[string]any
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	elements, _ := scene["elements"].([]any)
	crossFrameArrows := 0
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoCrossFrame"] != true {
			continue
		}
		crossFrameArrows++
		if customData["xaligoConnectorSourceFrameSide"] != "right" || customData["xaligoConnectorDestinationFrameSide"] != "left" {
			t.Fatalf("cross-frame activity sides = %#v", customData)
		}
	}
	if crossFrameArrows != 2 {
		t.Fatalf("cross-frame activity arrows = %d, want two page-local stubs: %s", crossFrameArrows, rawScene)
	}
}

func TestUMLGuardLabelsDoNotOverlapTheirConnectorSegment(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="420" height="320"><uml id="activity"><activity-diagram direction="down"><action id="one" title="One"/><action id="two" title="Two"/><control-flow src="one" dst="two" guard="next step"/></activity-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var arrow, label *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil || element.CustomData.UMLGuard != "next step" {
			continue
		}
		if element.Type == "arrow" {
			arrow = element
		} else if element.Type == "text" {
			label = element
		}
	}
	if arrow == nil || label == nil || len(arrow.Points) < 2 {
		t.Fatalf("guard arrow/label not found: %#v", scene.Elements)
	}
	for index := 0; index < len(arrow.Points)-1; index++ {
		start, end := arrow.Points[index], arrow.Points[index+1]
		if len(start) < 2 || len(end) < 2 || math.Abs(start[0]-end[0]) > 0.01 {
			continue
		}
		x := arrow.X + start[0]
		y1, y2 := arrow.Y+start[1], arrow.Y+end[1]
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		xIntersectsLabel := x >= label.X && x <= label.X+label.Width
		yIntersectsLabel := y2 >= label.Y && y1 <= label.Y+label.Height
		if xIntersectsLabel && yIntersectsLabel {
			t.Fatalf("guard label overlaps vertical connector segment: arrow=%#v label=%#v", arrow, label)
		}
	}
}

func TestUMLTimingAndOwnerMetadataReachEditableScene(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="timing"><timing-diagram><lifeline id="api"/><time-state id="busy" owner="api" from="10" to="20"><region>work</region></time-state><occurrence src="api" dst="busy" at="15" title="dispatch"/></timing-diagram></uml></frame></frames></xaligo>`)
	scene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	for _, want := range []string{
		`"xaligoUmlOwnerReference": "api"`,
		`"xaligoUmlCompartmentKinds": "region"`,
		`"xaligoUmlTimeFrom": "10"`,
		`"xaligoUmlTimeTo": "20"`,
		`"xaligoUmlOccurrenceAt": "15"`,
	} {
		if !strings.Contains(string(scene), want) {
			t.Fatalf("scene missing %q: %s", want, scene)
		}
	}
}

func TestUMLSequenceOrderControlsVerticalMessageAnchors(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b" order="1"/><message src="b" dst="b" order="2"/><return-message src="b" dst="a" order="3"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	anchors := map[string][2]float64{}
	for _, element := range scene.Elements {
		if element.Type != "arrow" || element.CustomData == nil || element.CustomData.UMLMessageOrder == "" || element.StartBinding == nil || element.EndBinding == nil {
			continue
		}
		anchors[element.CustomData.UMLMessageOrder] = [2]float64{element.StartBinding.FixedPoint[1], element.EndBinding.FixedPoint[1]}
	}
	if len(anchors) != 3 {
		t.Fatalf("sequence anchors = %#v", anchors)
	}
	if !(anchors["1"][0] < anchors["2"][0] && anchors["2"][0] < anchors["3"][0]) {
		t.Fatalf("message anchors are not ordered top-to-bottom: %#v", anchors)
	}
	if math.Abs(anchors["2"][0]-anchors["2"][1]) < 0.01 {
		t.Fatalf("self-message endpoints should form a loop: %#v", anchors["2"])
	}
}

func TestUMLAggregationAndCompositionRemainHeadlessAtDestination(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="classes"><class-diagram><class id="whole"/><class id="aggregate"/><class id="composite"/><aggregation src="whole" dst="aggregate"/><composition src="whole" dst="composite"/></class-diagram></uml></frame></frames></xaligo>`)
	options := entity.RenderOptions{PxPerInch: 96}
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, options)
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	relationStyles := map[string][2]string{}
	for _, element := range scene.Elements {
		if element.Type != "arrow" || element.CustomData == nil || element.CustomData.UMLRelationKind == "" {
			continue
		}
		relationStyles[element.CustomData.UMLRelationKind] = [2]string{
			element.CustomData.ConnectorStartArrowhead,
			element.CustomData.ConnectorEndArrowhead,
		}
		if element.StrokeColor != "#052d6e" || element.StrokeWidth < 1.3 {
			t.Fatalf("UML %s line style = color %q width %.2f, want xaligo activity palette", element.CustomData.UMLRelationKind, element.StrokeColor, element.StrokeWidth)
		}
		if !element.CustomData.ConnectorStartArrowheadExplicit || !element.CustomData.ConnectorEndArrowheadExplicit {
			t.Fatalf("UML %s arrowheads must be explicit: %#v", element.CustomData.UMLRelationKind, element.CustomData)
		}
	}
	for _, kind := range []string{"aggregation", "composition"} {
		if got := relationStyles[kind]; got != [2]string{"diamond", "none"} {
			t.Fatalf("UML %s arrowheads = %#v, want diamond/none", kind, got)
		}
	}

	rawPlan, err := newUsecase().BuildPPTXPlan(context.Background(), source, options)
	if err != nil {
		t.Fatalf("BuildPPTXPlan() error = %v", err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(rawPlan, &plan); err != nil {
		t.Fatalf("json.Unmarshal(plan) error = %v", err)
	}
	lineCount := 0
	for _, operation := range plan.Ops {
		if operation.Kind != "line" || operation.Line == nil {
			continue
		}
		if operation.Line.BeginArrowType == "none" && operation.Line.EndArrowType == "none" {
			continue
		}
		lineCount++
		if operation.Line.BeginArrowType != "diamond" || operation.Line.EndArrowType != "none" {
			t.Fatalf("plan UML arrowheads = %q/%q, want diamond/none", operation.Line.BeginArrowType, operation.Line.EndArrowType)
		}
	}
	if lineCount != 2 {
		t.Fatalf("plan UML line count = %d, want 2", lineCount)
	}
}

func TestUMLClassStereotypeAndModifiersReachEditableScene(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="640" height="360"><uml id="classes"><class-diagram><class id="repository" title="Repository" stereotype="service" abstract="true" static="true"><attribute>- store: Store</attribute><operation>+ find(id): Entity</operation></class></class-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var classElement *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData != nil && element.CustomData.UMLStereotype == "service" {
			classElement = element
			break
		}
	}
	if classElement == nil || classElement.CustomData.UMLAbstract != "true" || classElement.CustomData.UMLStatic != "true" || classElement.CustomData.UMLCompartmentKinds != "attribute,operation" {
		t.Fatalf("class UML metadata missing: %#v", scene.Elements)
	}
	if classElement.StrokeColor != "#052d6e" || classElement.BackgroundColor != "#e8f7fd" || classElement.StrokeWidth < 1.3 {
		t.Fatalf("class style = stroke %q background %q width %.2f, want xaligo activity palette", classElement.StrokeColor, classElement.BackgroundColor, classElement.StrokeWidth)
	}
	if classElement.Width > 260 || classElement.Height > 180 {
		t.Fatalf("class box size = %.1fx%.1f, want compact Lucid-like classifier", classElement.Width, classElement.Height)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	text := ""
	var textElement map[string]any
	for _, rawElement := range raw["elements"].([]any) {
		element, _ := rawElement.(map[string]any)
		if element["containerId"] == classElement.ID && element["type"] == "text" {
			text, _ = element["text"].(string)
			textElement = element
			break
		}
	}
	wantText := "<<service>>\n{abstract, static} Repository\n- store: Store\n+ find(id): Entity"
	if text != wantText {
		t.Fatalf("class label = %q, want %q", text, wantText)
	}
	if textElement["strokeColor"] != "#052d6e" {
		t.Fatalf("class text color = %q, want xaligo text color", textElement["strokeColor"])
	}
	foundDivider := false
	for _, rawElement := range raw["elements"].([]any) {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlClassHeaderDivider"] == true {
			foundDivider = true
			break
		}
	}
	if !foundDivider {
		t.Fatalf("class header divider missing: %s", rawScene)
	}
	xyflow, err := newUsecase().RenderXYFlow(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderXYFlow() error = %v", err)
	}
	for _, want := range []string{`"umlStereotype": "service"`, `"umlAbstract": "true"`, `"umlStatic": "true"`} {
		if !strings.Contains(string(xyflow), want) {
			t.Fatalf("XYFlow missing %q: %s", want, xyflow)
		}
	}
}

func TestUMLSequenceOrderAnchorsRemainTopToBottomForEveryConnectionSide(t *testing.T) {
	tests := []struct {
		name  string
		sides string
	}{
		{name: "left and right", sides: `src-side="left" dst-side="right"`},
		{name: "top and bottom", sides: `src-side="top" dst-side="bottom"`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b" order="1" ` + test.sides + `/><message src="a" dst="b" order="2" ` + test.sides + `/><message src="a" dst="b" order="3" ` + test.sides + `/></sequence-diagram></uml></frame></frames></xaligo>`)
			rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
			if err != nil {
				t.Fatalf("RenderExcalidraw() error = %v", err)
			}
			var scene entity.PresentationScene
			if err := json.Unmarshal(rawScene, &scene); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			anchors := map[string][2][]float64{}
			for _, element := range scene.Elements {
				if element.Type != "arrow" || element.CustomData == nil || element.CustomData.UMLMessageOrder == "" || element.StartBinding == nil || element.EndBinding == nil {
					continue
				}
				anchors[element.CustomData.UMLMessageOrder] = [2][]float64{
					element.StartBinding.FixedPoint,
					element.EndBinding.FixedPoint,
				}
			}
			if len(anchors) != 3 {
				t.Fatalf("sequence anchors = %#v", anchors)
			}
			for endpoint := 0; endpoint < 2; endpoint++ {
				first, middle, last := anchors["1"][endpoint], anchors["2"][endpoint], anchors["3"][endpoint]
				if len(first) != 2 || len(middle) != 2 || len(last) != 2 {
					t.Fatalf("sequence fixed points = %#v", anchors)
				}
				if !(first[1] < middle[1] && middle[1] < last[1]) {
					t.Fatalf("endpoint %d anchors are not ordered top-to-bottom: %#v", endpoint, anchors)
				}
				for _, point := range [][]float64{first, middle, last} {
					if point[0] != 0 && point[0] != 1 {
						t.Fatalf("endpoint %d anchor is not on a vertical edge: %#v", endpoint, point)
					}
				}
			}
		})
	}
}

func TestUMLCrossFramePublicReferenceReachesGraphOutputs(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames>
<frame id="overview" width="720" height="420"><rectangle id="caller" title="Caller"/><connection src="caller" dst="detail.order"/></frame>
<frame id="detail" width="720" height="420"><uml id="domain"><class-diagram><class id="order" title="Order"/></class-diagram></uml></frame>
</frames></xaligo>`)
	options := entity.RenderOptions{PxPerInch: 96, Theme: "light"}

	rawXYFlow, err := newUsecase().RenderXYFlow(context.Background(), source, options)
	if err != nil {
		t.Fatalf("RenderXYFlow() error = %v", err)
	}
	var xyflow entity.XYFlowDocument
	if err := json.Unmarshal(rawXYFlow, &xyflow); err != nil {
		t.Fatalf("json.Unmarshal(XYFlow) error = %v", err)
	}
	if len(xyflow.Edges) != 1 || xyflow.Edges[0].Source == "" || xyflow.Edges[0].Target == "" || xyflow.Edges[0].Source == xyflow.Edges[0].Target {
		t.Fatalf("XYFlow cross-frame UML edge = %#v", xyflow.Edges)
	}
	if crossFrame, _ := xyflow.Edges[0].Data["crossFrame"].(bool); !crossFrame {
		t.Fatalf("XYFlow edge missing cross-frame metadata: %#v", xyflow.Edges[0])
	}

	rawIsoflow, err := newUsecase().RenderIsoflow(context.Background(), source, options)
	if err != nil {
		t.Fatalf("RenderIsoflow() error = %v", err)
	}
	var isoflow entity.IsoflowDocument
	if err := json.Unmarshal(rawIsoflow, &isoflow); err != nil {
		t.Fatalf("json.Unmarshal(Isoflow) error = %v", err)
	}
	if len(isoflow.Views) != 1 || len(isoflow.Views[0].Connectors) != 1 {
		t.Fatalf("Isoflow cross-frame UML connectors = %#v", isoflow.Views)
	}
	anchors := isoflow.Views[0].Connectors[0].Anchors
	if len(anchors) < 2 || anchors[0].Ref.Item == "" || anchors[len(anchors)-1].Ref.Item == "" || anchors[0].Ref.Item == anchors[len(anchors)-1].Ref.Item {
		t.Fatalf("Isoflow cross-frame UML anchors = %#v", anchors)
	}
}
