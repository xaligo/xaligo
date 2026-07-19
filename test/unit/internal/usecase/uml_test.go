package usecase_test

import (
	"context"
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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

func TestUMLCommonRelationAnchorsUseShapeProfiles(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="900" height="520"><uml id="state"><state-machine-diagram direction="right"><state id="source" title="Source"/><state id="target" title="Target"/><state id="alternate" title="Alternate"/><choice id="choice"/><transition src="source" dst="target" src-anchor="right-5" dst-anchor="left-2"/><transition src="source" dst="choice" dst-anchor="top-5"/><transition src="choice" dst="target" src-anchor="right-1"/><transition src="choice" dst="alternate" src-anchor="bottom-2"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var rectangleRelation, diamondDestination, diamondSource *entity.Element
	for i := range scene.Elements {
		element := &scene.Elements[i]
		if element.Type != "arrow" || element.CustomData == nil {
			continue
		}
		src := element.CustomData.UMLRelationSourceReference
		dst := element.CustomData.UMLRelationDestinationReference
		switch {
		case src == "source" && dst == "target":
			rectangleRelation = element
		case src == "source" && dst == "choice":
			diamondDestination = element
		case src == "choice" && dst == "target":
			diamondSource = element
		}
	}
	if rectangleRelation == nil || rectangleRelation.StartBinding == nil || rectangleRelation.EndBinding == nil {
		t.Fatalf("rectangle relation bindings not found")
	}
	if !sameFixedPointV1UsecaseUMLTest(rectangleRelation.StartBinding.FixedPoint, []float64{1, 0.9}) {
		t.Fatalf("rectangle src fixedPoint = %#v, want right-5", rectangleRelation.StartBinding.FixedPoint)
	}
	if !sameFixedPointV1UsecaseUMLTest(rectangleRelation.EndBinding.FixedPoint, []float64{0, 0.3}) {
		t.Fatalf("rectangle dst fixedPoint = %#v, want left-2", rectangleRelation.EndBinding.FixedPoint)
	}
	if diamondDestination == nil || diamondDestination.EndBinding == nil {
		t.Fatalf("diamond destination binding not found")
	}
	if !sameFixedPointV1UsecaseUMLTest(diamondDestination.EndBinding.FixedPoint, []float64{0.5, 0}) {
		t.Fatalf("diamond dst fixedPoint = %#v, want top vertex", diamondDestination.EndBinding.FixedPoint)
	}
	if diamondSource == nil || diamondSource.StartBinding == nil {
		t.Fatalf("diamond source binding not found")
	}
	if !sameFixedPointV1UsecaseUMLTest(diamondSource.StartBinding.FixedPoint, []float64{1, 0.5}) {
		t.Fatalf("diamond src fixedPoint = %#v, want right vertex", diamondSource.StartBinding.FixedPoint)
	}
}

func TestUMLCommonRelationAnchorsReachDrawPlanRoutes(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="900" height="520"><uml id="state"><state-machine-diagram direction="right"><state id="source" title="Source"/><state id="target" title="Target"/><state id="alternate" title="Alternate"/><choice id="choice"/><transition src="source" dst="target" src-anchor="right-5" dst-anchor="left-2"/><transition src="source" dst="choice" dst-anchor="top-5"/><transition src="choice" dst="target" src-anchor="right-1"/><transition src="choice" dst="alternate" src-anchor="bottom-2"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawPlan, err := newUsecase().BuildPPTXPlan(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("BuildPPTXPlan() error = %v", err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(rawPlan, &plan); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	shapes := map[string]entity.DrawOp{}
	lines := []entity.DrawOp{}
	for _, operation := range plan.Ops {
		if strings.HasPrefix(operation.ID, "conn-") && operation.Kind == "line" {
			lines = append(lines, operation)
			continue
		}
		if !operation.FrontLayer && strings.HasSuffix(operation.ID, "-rect") && (operation.Kind == "rect" || operation.Kind == "diamond") {
			shapes[operation.ID] = operation
		}
	}
	if len(lines) != 4 {
		t.Fatalf("connector lines = %#v, want 4", lines)
	}
	checkedDiamond := false
	checkedRectangle := false
	for _, line := range lines {
		if len(line.Points) < 2 {
			t.Fatalf("connector line = %#v", line)
		}
		points := absoluteLinePointsV1UsecaseUMLTest(line)
		srcShape, hasSource := endpointShapeV1UsecaseUMLTest(points[0], shapes)
		dstShape, hasDestination := endpointShapeV1UsecaseUMLTest(points[len(points)-1], shapes)
		if !hasSource || !hasDestination {
			t.Fatalf("connector %q endpoint shapes not found for %#v in %#v", line.ID, points, shapes)
		}
		if srcShape.Kind == "diamond" {
			if !pointIsDiamondVertexV1UsecaseUMLTest(points[0], srcShape) {
				t.Fatalf("diamond source point = %#v shape=%#v", points[0], srcShape)
			}
			checkedDiamond = true
		} else {
			if !pointIsRectangleSlotV1UsecaseUMLTest(points[0], srcShape) {
				t.Fatalf("rectangle source point = %#v shape=%#v", points[0], srcShape)
			}
			checkedRectangle = true
		}
		last := points[len(points)-1]
		if dstShape.Kind == "diamond" {
			if !pointIsDiamondVertexV1UsecaseUMLTest(last, dstShape) {
				t.Fatalf("diamond destination point = %#v shape=%#v", last, dstShape)
			}
			checkedDiamond = true
		} else {
			if !pointIsRectangleSlotV1UsecaseUMLTest(last, dstShape) {
				t.Fatalf("rectangle destination point = %#v shape=%#v", last, dstShape)
			}
			checkedRectangle = true
		}
	}
	if !checkedDiamond || !checkedRectangle {
		t.Fatalf("checked diamond=%t rectangle=%t lines=%#v shapes=%#v", checkedDiamond, checkedRectangle, lines, shapes)
	}
}

func sameFixedPointV1UsecaseUMLTest(got, want []float64) bool {
	return len(got) == 2 && len(want) == 2 && math.Abs(got[0]-want[0]) < 0.0001 && math.Abs(got[1]-want[1]) < 0.0001
}

func samePointV1UsecaseUMLTest(got struct{ X, Y float64 }, wantX, wantY float64) bool {
	return math.Abs(got.X-wantX) < 0.0001 && math.Abs(got.Y-wantY) < 0.0001
}

func absoluteLinePointsV1UsecaseUMLTest(operation entity.DrawOp) []struct{ X, Y float64 } {
	points := make([]struct{ X, Y float64 }, 0, len(operation.Points))
	for _, point := range operation.Points {
		points = append(points, struct{ X, Y float64 }{X: operation.X + point.X, Y: operation.Y + point.Y})
	}
	return points
}

func pointIsDiamondVertexV1UsecaseUMLTest(point struct{ X, Y float64 }, shape entity.DrawOp) bool {
	vertices := []struct{ X, Y float64 }{
		{X: shape.X + shape.W/2, Y: shape.Y},
		{X: shape.X + shape.W, Y: shape.Y + shape.H/2},
		{X: shape.X + shape.W/2, Y: shape.Y + shape.H},
		{X: shape.X, Y: shape.Y + shape.H/2},
	}
	for _, vertex := range vertices {
		if nearPointV1UsecaseUMLTest(point, vertex.X, vertex.Y, 0.09) {
			return true
		}
	}
	return false
}

func endpointShapeV1UsecaseUMLTest(point struct{ X, Y float64 }, shapes map[string]entity.DrawOp) (entity.DrawOp, bool) {
	for _, shape := range shapes {
		if shape.Kind == "diamond" && pointIsDiamondVertexV1UsecaseUMLTest(point, shape) {
			return shape, true
		}
		if shape.Kind != "diamond" && pointIsRectangleSlotV1UsecaseUMLTest(point, shape) {
			return shape, true
		}
	}
	return entity.DrawOp{}, false
}

func pointIsRectangleSlotV1UsecaseUMLTest(point struct{ X, Y float64 }, shape entity.DrawOp) bool {
	slots := []float64{0.1, 0.3, 0.5, 0.7, 0.9}
	for _, slot := range slots {
		if nearPointV1UsecaseUMLTest(point, shape.X+shape.W*slot, shape.Y, 0.09) ||
			nearPointV1UsecaseUMLTest(point, shape.X+shape.W, shape.Y+shape.H*slot, 0.09) ||
			nearPointV1UsecaseUMLTest(point, shape.X+shape.W*slot, shape.Y+shape.H, 0.09) ||
			nearPointV1UsecaseUMLTest(point, shape.X, shape.Y+shape.H*slot, 0.09) {
			return true
		}
	}
	return false
}

func nearPointV1UsecaseUMLTest(got struct{ X, Y float64 }, wantX, wantY, tolerance float64) bool {
	return math.Hypot(got.X-wantX, got.Y-wantY) <= tolerance
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

func TestUMLClassHidesRedundantContainerBorderAndTitle(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" title="Class Model" width="760" height="420"><uml id="classes" title="Domain Classes"><class-diagram><class id="user" title="User"><attribute>- id: int</attribute><operation>+ login()</operation></class></class-diagram></uml></frame></frames></xaligo>`)
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
		if element["type"] == "rectangle" && customData["xaligoUmlDiagramKind"] == "class-diagram" && customData["xaligoUmlElementKind"] == nil {
			t.Fatalf("class container border should not render: %#v", element)
		}
		if element["type"] == "text" && element["text"] == "Domain Classes" {
			t.Fatalf("class container title should not render when frame metadata can carry the title: %#v", element)
		}
	}
}

func TestUMLSequenceHidesRedundantContainerBorderAndTitle(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" title="Sequence Model" width="760" height="420"><uml id="sequence" title="Checkout Sequence"><sequence-diagram><participant id="user" title="User"/><lifeline id="api" title="API"/><message src="user" dst="api" order="1" title="submit()"/></sequence-diagram></uml></frame></frames></xaligo>`)
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
		if element["type"] == "rectangle" && customData["xaligoUmlDiagramKind"] == "sequence-diagram" && customData["xaligoUmlElementKind"] == nil {
			t.Fatalf("sequence container border should not render: %#v", element)
		}
		if element["type"] == "text" && element["text"] == "Checkout Sequence" {
			t.Fatalf("sequence container title should not render when frame metadata can carry the title: %#v", element)
		}
	}
}

func TestUMLStateMachineHidesRedundantContainerBorderAndTitle(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" title="State Machine" width="760" height="420"><uml id="state" title="Order State"><state-machine-diagram><initial id="start"/><state id="open" title="Open"/><final id="done"/><transition src="start" dst="open"/><transition src="open" dst="done"/></state-machine-diagram></uml></frame></frames></xaligo>`)
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
		if element["type"] == "rectangle" && customData["xaligoUmlDiagramKind"] == "state-machine-diagram" && customData["xaligoUmlElementKind"] == nil {
			t.Fatalf("state-machine container border should not render: %#v", element)
		}
		if element["type"] == "text" && element["text"] == "Order State" {
			t.Fatalf("state-machine container title should not render when frame metadata can carry the title: %#v", element)
		}
	}
}

func TestUMLStateMachineCanHideCompartmentElementNames(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="1200" height="520"><uml id="state"><state-machine-diagram direction="right" show-element-names="false"><state id="hidden" title="Hidden"><entry>prepare order</entry><do>pack order</do></state><state id="visible" title="Visible" show-element-names="true"><exit>visible exit</exit></state><choice id="choice" title="Hidden Choice"/><state id="accepted" title="Accepted"/><state id="rejected" title="Rejected"/><transition src="hidden" dst="choice"/><transition src="choice" dst="accepted" guard="yes"/><transition src="choice" dst="rejected" guard="no"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	sceneTexts := map[string]int{}
	for _, element := range scene.Elements {
		if element.Type == "text" {
			sceneTexts[element.Text]++
		}
	}
	for _, hidden := range []string{"entry", "do"} {
		if sceneTexts[hidden] != 0 {
			t.Fatalf("scene text %q rendered despite show-element-names=false: %#v", hidden, sceneTexts)
		}
	}
	for _, visible := range []string{"Hidden", "Hidden Choice", "Accepted", "Rejected", "Visible", "prepare order", "pack order", "exit", "visible exit"} {
		if sceneTexts[visible] == 0 {
			t.Fatalf("scene text %q missing: %#v", visible, sceneTexts)
		}
	}
	rawPlan, err := newUsecase().BuildPPTXPlan(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("BuildPPTXPlan() error = %v", err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(rawPlan, &plan); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	planTexts := map[string]int{}
	for _, operation := range plan.Ops {
		if operation.Kind == "text" {
			planTexts[operation.Text]++
		}
	}
	for _, hidden := range []string{"entry", "do"} {
		if planTexts[hidden] != 0 {
			t.Fatalf("plan text %q rendered despite show-element-names=false: %#v", hidden, planTexts)
		}
	}
	for _, visible := range []string{"Hidden", "Hidden Choice", "Accepted", "Rejected", "Visible", "prepare order", "pack order", "exit", "visible exit"} {
		if planTexts[visible] == 0 {
			t.Fatalf("plan text %q missing: %#v", visible, planTexts)
		}
	}
}

func TestUMLComponentPortsAttachToOwners(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="760" height="420"><uml id="components"><component-diagram direction="right"><component id="service" title="Order Service"/><port id="api" owner="service" side="right" title="api"/><component id="client" title="Client"/><port id="client-api" owner="client" side="left" title="api"/><assembly src="api" dst="client-api" title="uses"/></component-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var service, port *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil || element.Type == "text" {
			continue
		}
		switch {
		case element.CustomData.UMLLocalID == "service" && element.CustomData.UMLElementKind == "component":
			service = element
		case element.CustomData.UMLLocalID == "api" && element.CustomData.UMLElementKind == "port":
			port = element
		}
	}
	if service == nil || port == nil {
		t.Fatalf("component or port missing: service=%#v port=%#v", service, port)
	}
	if math.Abs((port.X+port.Width)-(service.X+service.Width)) > 0.1 {
		t.Fatalf("owned port is not attached to component right edge: component=%#v port=%#v", service, port)
	}
	if port.Y < service.Y-0.1 || port.Y+port.Height > service.Y+service.Height+0.1 {
		t.Fatalf("owned port is outside component vertical bounds: component=%#v port=%#v", service, port)
	}
}

func TestUMLComponentRendersComponentNotation(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="640" height="360"><uml id="components"><component-diagram><component id="service" title="Order Service"/></component-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	elements, _ := raw["elements"].([]any)
	headerCount, headerTextCount := 0, 0
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlComponentHeader"] == true {
			headerCount++
			if element["backgroundColor"] != "#08b8ea" {
				t.Fatalf("component header background = %q, want #08b8ea", element["backgroundColor"])
			}
		}
		if customData["xaligoUmlComponentHeaderContent"] == true {
			headerTextCount++
			if element["textAlign"] != "left" || element["strokeColor"] != "#ffffff" {
				t.Fatalf("component header text style = align %q color %q", element["textAlign"], element["strokeColor"])
			}
		}
	}
	if headerCount != 1 || headerTextCount != 1 {
		t.Fatalf("component notation counts = header %d header text %d, want 1, 1", headerCount, headerTextCount)
	}
}

func TestUMLComponentRendersBoundaryInterfaces(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="760" height="420"><uml id="components"><component-diagram><component id="service" title="Order Service"><interface>Ordering API</interface><interface>Order Store</interface></component></component-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	elements, _ := raw["elements"].([]any)
	interfaceSymbols, interfacePorts, portLabels := 0, 0, 0
	var component map[string]any
	var componentHeader map[string]any
	var interfacePort map[string]any
	var interfaceSymbol map[string]any
	var interfacePortElements []map[string]any
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		switch {
		case customData["xaligoUmlElementKind"] == "component":
			component = element
		case customData["xaligoUmlComponentHeader"] == true:
			componentHeader = element
		case customData["xaligoUmlComponentInterfaceSymbol"] == true:
			interfaceSymbols++
			interfaceSymbol = element
		case customData["xaligoUmlComponentInterfacePort"] == true:
			interfacePorts++
			interfacePort = element
			interfacePortElements = append(interfacePortElements, element)
		case customData["xaligoUmlComponentInterfacePortLabel"] == true:
			portLabels++
		}
	}
	if interfaceSymbols != 2 || interfacePorts != 2 || portLabels != 2 {
		t.Fatalf("boundary interface counts = symbols %d ports %d labels %d, want 2 each", interfaceSymbols, interfacePorts, portLabels)
	}
	if component == nil {
		t.Fatalf("component shape missing")
	}
	if componentHeader == nil {
		t.Fatalf("component header missing")
	}
	if component["backgroundColor"] != "#ffffff" {
		t.Fatalf("component background = %q, want #ffffff", component["backgroundColor"])
	}
	if interfacePort["type"] != "rectangle" || interfacePort["backgroundColor"] != "#ffffff" {
		t.Fatalf("interface port style = type %q background %q", interfacePort["type"], interfacePort["backgroundColor"])
	}
	componentX := component["x"].(float64)
	portX := interfacePort["x"].(float64)
	portW := interfacePort["width"].(float64)
	symbolX := interfaceSymbol["x"].(float64)
	symbolW := interfaceSymbol["width"].(float64)
	portProtrusion := componentX - portX
	if !(portProtrusion > 0 && portProtrusion < portW*0.25) {
		t.Fatalf("interface port protrusion = %.1f of width %.1f, want a small overlap outside component", portProtrusion, portW)
	}
	if gap := portX - (symbolX + symbolW); gap < 2 || gap > 6 {
		t.Fatalf("interface symbol gap = %.1f, want circle closer to port rectangle", gap)
	}
	headerBottom := componentHeader["y"].(float64) + componentHeader["height"].(float64)
	for _, portElement := range interfacePortElements {
		portY := portElement["y"].(float64)
		if portY < headerBottom+6 {
			t.Fatalf("interface port y = %.1f overlaps header bottom %.1f", portY, headerBottom)
		}
	}
}

func TestUMLComponentSampleUsesBoundaryInterfacesWithoutPorts(t *testing.T) {
	source, err := os.ReadFile("../../../../docs/src/examples/samples/uml-component.xal")
	if err != nil {
		t.Fatalf("os.ReadFile() error = %v", err)
	}
	if strings.Contains(string(source), "<port ") || strings.Contains(string(source), "<assembly ") {
		t.Fatalf("component sample should use component boundary interfaces without explicit ports or assemblies")
	}
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	elements, _ := raw["elements"].([]any)
	umlPorts, interfaceSymbols, interfacePorts := 0, 0, 0
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlElementKind"] == "port" {
			umlPorts++
		}
		if customData["xaligoUmlComponentInterfaceSymbol"] == true {
			interfaceSymbols++
		}
		if customData["xaligoUmlComponentInterfacePort"] == true {
			interfacePorts++
		}
	}
	if umlPorts != 0 || interfaceSymbols != 4 || interfacePorts != 4 {
		t.Fatalf("sample symbols = UML ports %d interface symbols %d interface ports %d, want 0, 4, 4", umlPorts, interfaceSymbols, interfacePorts)
	}
}

func TestUMLConnectedComponentSampleRendersBoundaryInterfaces(t *testing.T) {
	source, err := os.ReadFile("../../../../docs/src/examples/samples/uml-component-connected.xal")
	if err != nil {
		t.Fatalf("os.ReadFile() error = %v", err)
	}
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	elements, _ := raw["elements"].([]any)
	interfaceSymbols, circles, callerSockets, associations := 0, 0, 0, 0
	lastInterfaceIndex, associationIndex := -1, -1
	interfaceIDs := map[string]bool{}
	associationStartID, associationEndID := "", ""
	var callerSocketElement map[string]any
	for index, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlComponentInterfaceSymbol"] == true {
			interfaceSymbols++
			lastInterfaceIndex = index
			id, _ := element["id"].(string)
			interfaceIDs[id] = true
			if customData["xaligoUmlComponentInterfaceCircle"] == true {
				circles++
			}
		}
		if customData["xaligoUmlComponentCallerSocket"] == true {
			callerSockets++
			callerSocketElement = element
		}
		if customData["xaligoUmlRelationKind"] == "association" {
			associations++
			associationIndex = index
			startBinding, _ := element["startBinding"].(map[string]any)
			endBinding, _ := element["endBinding"].(map[string]any)
			associationStartID, _ = startBinding["elementId"].(string)
			associationEndID, _ = endBinding["elementId"].(string)
		}
	}
	if interfaceSymbols != 2 || circles != 2 || callerSockets != 1 || associations != 1 {
		t.Fatalf("connected sample counts = interface symbols %d circles %d caller sockets %d associations %d, want 2, 2, 1, and 1", interfaceSymbols, circles, callerSockets, associations)
	}
	if associationIndex <= lastInterfaceIndex {
		t.Fatalf("association connector index = %d, last interface index = %d; connector should render in front", associationIndex, lastInterfaceIndex)
	}
	if interfaceIDs[associationStartID] || !interfaceIDs[associationEndID] {
		t.Fatalf("association binding = %q -> %q, want component anchor -> interface circle IDs %#v", associationStartID, associationEndID, interfaceIDs)
	}
	if callerSocketElement == nil {
		t.Fatalf("caller socket element missing")
	}
	if callerSocketElement["height"].(float64) <= callerSocketElement["width"].(float64) {
		t.Fatalf("caller socket size = %.1fx%.1f, want vertical semicircle for horizontal caller line", callerSocketElement["width"].(float64), callerSocketElement["height"].(float64))
	}
}

func TestUMLComplexConnectedComponentSampleBindsInterfaces(t *testing.T) {
	source, err := os.ReadFile("../../../../docs/src/examples/samples/uml-component-connected-complex.xal")
	if err != nil {
		t.Fatalf("os.ReadFile() error = %v", err)
	}
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	elements, _ := raw["elements"].([]any)
	interfaceIDs := map[string]bool{}
	callerSockets := 0
	associationBindings := 0
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		id, _ := element["id"].(string)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlComponentInterfaceSymbol"] == true {
			interfaceIDs[id] = true
		}
		if customData["xaligoUmlComponentCallerSocket"] == true {
			callerSockets++
		}
		if customData["xaligoUmlRelationKind"] != "association" {
			continue
		}
		startBinding, _ := element["startBinding"].(map[string]any)
		endBinding, _ := element["endBinding"].(map[string]any)
		startID, _ := startBinding["elementId"].(string)
		endID, _ := endBinding["elementId"].(string)
		endFixedPoint, _ := endBinding["fixedPoint"].([]any)
		if interfaceIDs[startID] || !interfaceIDs[endID] {
			t.Fatalf("association binds %q -> %q, want component anchor -> interface circle", startID, endID)
		}
		if len(endFixedPoint) < 2 || endFixedPoint[0] != 0.0 || endFixedPoint[1] != 0.5 {
			t.Fatalf("association end fixed point = %#v, want left-center of interface circle", endFixedPoint)
		}
		points, _ := element["points"].([]any)
		if len(points) < 2 {
			t.Fatalf("association points = %#v, want at least two points", points)
		}
		prev, _ := points[len(points)-2].([]any)
		last, _ := points[len(points)-1].([]any)
		if len(prev) < 2 || len(last) < 2 {
			t.Fatalf("association final points = %#v -> %#v, want coordinate pairs", prev, last)
		}
		prevY, _ := prev[1].(float64)
		lastY, _ := last[1].(float64)
		if math.Abs(prevY-lastY) > 0.1 {
			t.Fatalf("association final segment = %#v -> %#v, want horizontal approach into interface circle", prev, last)
		}
		prevX, _ := prev[0].(float64)
		lastX, _ := last[0].(float64)
		if prevX >= lastX {
			t.Fatalf("association final segment = %#v -> %#v, want approach from left of interface circle", prev, last)
		}
		associationBindings++
	}
	if len(interfaceIDs) != 8 || callerSockets != 4 || associationBindings != 4 {
		t.Fatalf("complex sample counts = interface symbols %d caller sockets %d associations %d, want 8, 4, and 4", len(interfaceIDs), callerSockets, associationBindings)
	}
	svg, err := newUsecase().RenderSVG(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderSVG() error = %v", err)
	}
	paths := svgUMLComponentAssociationPathsV1UMLTest(t, string(svg))
	if len(paths) != 4 {
		t.Fatalf("SVG association paths = %d, want 4", len(paths))
	}
	for _, path := range paths {
		prev := path[len(path)-2]
		last := path[len(path)-1]
		if math.Abs(prev.y-last.y) > 0.1 {
			t.Fatalf("SVG association final segment = %#v -> %#v, want horizontal approach into interface circle", prev, last)
		}
		if prev.x >= last.x {
			t.Fatalf("SVG association final segment = %#v -> %#v, want approach from left of interface circle", prev, last)
		}
	}
}

func TestUMLComponentMultipleCallersRenderSeparateInterfaceCircles(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="980" height="520"><uml id="components"><component-diagram grid="3"><component id="web" title="Web"><interface>Shared API</interface></component><component id="worker" title="Worker"><interface>Shared API</interface></component><component id="api" title="API"><interface>Shared API</interface><interface>Admin API</interface></component><association src="web" dst="api"/><association src="worker" dst="api"/></component-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	elements, _ := raw["elements"].([]any)
	interfaceCircles := map[string]map[string]any{}
	associationEndIDs := map[string]bool{}
	associations := 0
	minAPICircleX := math.Inf(1)
	maxAPICircleX := math.Inf(-1)
	multiTrunks := 0
	multiPortStems := 0
	multiCircleStems := 0
	maxCircleStemWidth := 0.0
	portStemWidth := 0.0
	for _, rawElement := range elements {
		element, _ := rawElement.(map[string]any)
		id, _ := element["id"].(string)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlComponentInterfaceCircle"] == true {
			interfaceCircles[id] = element
			if customData["xaligoUmlComponentOwnerLocalId"] == "api" {
				x, _ := element["x"].(float64)
				minAPICircleX = math.Min(minAPICircleX, x)
				maxAPICircleX = math.Max(maxAPICircleX, x)
			}
		}
		if customData["xaligoUmlComponentInterfaceStem"] == true {
			width, _ := element["width"].(float64)
			height, _ := element["height"].(float64)
			switch {
			case strings.HasSuffix(id, "-multi-trunk"):
				multiTrunks++
				if math.Abs(width) > 0.1 || height <= 0 {
					t.Fatalf("multi trunk %q width %.1f height %.1f, want vertical trunk", id, width, height)
				}
			case strings.HasSuffix(id, "-multi-port-stem"):
				multiPortStems++
				if width <= 0 || math.Abs(height) > 0.1 {
					t.Fatalf("multi port stem %q width %.1f height %.1f, want horizontal stem", id, width, height)
				}
				portStemWidth = width
			case strings.Contains(id, "-multi-circle-stem-"):
				multiCircleStems++
				if width <= 0 || math.Abs(height) > 0.1 {
					t.Fatalf("multi circle stem %q width %.1f height %.1f, want horizontal stem", id, width, height)
				}
				maxCircleStemWidth = math.Max(maxCircleStemWidth, width)
			}
		}
		if customData["xaligoUmlRelationKind"] != "association" {
			continue
		}
		associations++
		endBinding, _ := element["endBinding"].(map[string]any)
		endID, _ := endBinding["elementId"].(string)
		endFixedPoint, _ := endBinding["fixedPoint"].([]any)
		if len(endFixedPoint) < 2 || endFixedPoint[0] != 0.0 || endFixedPoint[1] != 0.5 {
			t.Fatalf("association end fixed point = %#v, want left-center of interface circle", endFixedPoint)
		}
		circle := interfaceCircles[endID]
		if circle == nil {
			t.Fatalf("association end %q is not an interface circle", endID)
		}
		points, _ := element["points"].([]any)
		prev, _ := points[len(points)-2].([]any)
		last, _ := points[len(points)-1].([]any)
		prevX, _ := prev[0].(float64)
		lastX, _ := last[0].(float64)
		prevY, _ := prev[1].(float64)
		lastY, _ := last[1].(float64)
		if math.Abs(prevY-lastY) > 0.1 || prevX >= lastX {
			t.Fatalf("association final segment = %#v -> %#v, want horizontal approach from left", prev, last)
		}
		absoluteEndY := element["y"].(float64) + lastY
		circleCenterY := circle["y"].(float64) + circle["height"].(float64)/2
		if math.Abs(absoluteEndY-circleCenterY) > 0.1 {
			t.Fatalf("association end y = %.1f, circle center y = %.1f", absoluteEndY, circleCenterY)
		}
		associationEndIDs[endID] = true
	}
	if associations != 2 || len(associationEndIDs) != 2 {
		t.Fatalf("associations = %d distinct destination circles = %d, want 2 and 2", associations, len(associationEndIDs))
	}
	if len(interfaceCircles) != 5 {
		t.Fatalf("interface circles = %d, want four component interfaces plus one extra destination circle", len(interfaceCircles))
	}
	if math.Abs(maxAPICircleX-minAPICircleX) > 0.1 {
		t.Fatalf("API interface circle x range = %.1f..%.1f, want single and multiple circles horizontally aligned", minAPICircleX, maxAPICircleX)
	}
	if multiTrunks != 1 || multiPortStems != 1 || multiCircleStems != 2 {
		t.Fatalf("multi interface stems = trunks %d port stems %d circle stems %d, want 1, 1, and 2", multiTrunks, multiPortStems, multiCircleStems)
	}
	if maxCircleStemWidth <= 0 || portStemWidth <= 0 || maxCircleStemWidth > portStemWidth*1.25 {
		t.Fatalf("multi interface stem widths = circle %.1f port %.1f, want compact circle branches with a visible port stem", maxCircleStemWidth, portStemWidth)
	}
}

func TestUMLStateMachineFinalRendersFinalDot(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="760" height="420"><uml id="state"><state-machine-diagram><initial id="start"/><state id="open" title="Open"/><final id="done"/><transition src="start" dst="open"/><transition src="open" dst="done"/></state-machine-diagram></uml></frame></frames></xaligo>`)
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
		if customData["xaligoUmlFinalDot"] == true {
			return
		}
	}
	t.Fatalf("state-machine final dot missing: %s", rawScene)
}

func TestUMLStateMachinePseudostatesKeepCompactProportions(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="520"><uml id="state"><state-machine-diagram direction="right"><initial id="start"/><state id="open" title="Open"/><choice id="choice" title="Choice"/><state id="done-state" title="Done"/><final id="done"/><transition src="start" dst="open"/><transition src="open" dst="choice"/><transition src="choice" dst="done-state" guard="ok"/><transition src="choice" dst="open" guard="retry"/><transition src="done-state" dst="done"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	for _, element := range scene.Elements {
		if element.CustomData == nil || element.CustomData.UMLElementKind == "" {
			continue
		}
		switch element.CustomData.UMLElementKind {
		case "initial", "final", "choice":
			if math.Abs(element.Width-element.Height) > 0.01 {
				t.Fatalf("%s shape is not square: width %.1f height %.1f", element.CustomData.UMLElementKind, element.Width, element.Height)
			}
			if element.Width > 90 || element.Height > 90 {
				t.Fatalf("%s shape is too large: width %.1f height %.1f", element.CustomData.UMLElementKind, element.Width, element.Height)
			}
		}
	}
}

func TestUMLStateMachineRowsSeparateBranches(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="620"><uml id="state"><state-machine-diagram direction="right"><initial id="start" row="1"/><state id="paid" title="Paid" row="1"/><final id="done" row="1"/><state id="refund" title="Refund" row="2"/><state id="cancelled" title="Cancelled" row="3"/><transition src="start" dst="paid"/><transition src="paid" dst="done"/><transition src="paid" dst="refund"/><transition src="refund" dst="cancelled"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	positions := map[string]entity.Element{}
	for _, element := range scene.Elements {
		if element.CustomData == nil || element.CustomData.UMLLocalID == "" || element.CustomData.UMLElementKind == "" {
			continue
		}
		positions[element.CustomData.UMLLocalID] = element
	}
	paid, hasPaid := positions["paid"]
	refund, hasRefund := positions["refund"]
	cancelled, hasCancelled := positions["cancelled"]
	if !hasPaid || !hasRefund || !hasCancelled {
		t.Fatalf("state-machine row elements missing: %#v", positions)
	}
	if !(refund.Y > paid.Y+100 && cancelled.Y > refund.Y+100) {
		t.Fatalf("state-machine rows not separated: paid y=%.1f refund y=%.1f cancelled y=%.1f", paid.Y, refund.Y, cancelled.Y)
	}
	paidCenter := paid.X + paid.Width/2
	refundCenter := refund.X + refund.Width/2
	cancelledCenter := cancelled.X + cancelled.Width/2
	if math.Abs(paidCenter-refundCenter) > 2 || math.Abs(refundCenter-cancelledCenter) > 2 {
		t.Fatalf("related state-machine rows not kept near: paid=%.1f refund=%.1f cancelled=%.1f", paidCenter, refundCenter, cancelledCenter)
	}
	if paid.StrokeColor != "#052d6e" || paid.BackgroundColor != "#ffffff" || refund.StrokeColor != "#052d6e" || refund.BackgroundColor != "#ffffff" {
		t.Fatalf("state-machine palette does not match UML class theme: paid=%q/%q refund=%q/%q", paid.StrokeColor, paid.BackgroundColor, refund.StrokeColor, refund.BackgroundColor)
	}
}

func TestUMLStateMachineGridColumnsAlignRelatedRows(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="1080" height="620"><uml id="state"><state-machine-diagram direction="right"><initial id="start" row="1" col="1"/><state id="paid" title="Paid" row="1" col="2"/><state id="shipped" title="Shipped" row="1" col="4"/><final id="done" row="1" col="5"/><state id="return" title="Return" row="2" col="4"/><state id="cancelled" title="Cancelled" row="3" col="4"/><transition src="start" dst="paid"/><transition src="paid" dst="shipped"/><transition src="shipped" dst="done"/><transition src="shipped" dst="return"/><transition src="return" dst="cancelled"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	positions := map[string]entity.Element{}
	for _, element := range scene.Elements {
		if element.CustomData == nil || element.CustomData.UMLLocalID == "" || element.CustomData.UMLElementKind == "" {
			continue
		}
		positions[element.CustomData.UMLLocalID] = element
	}
	shipped, hasShipped := positions["shipped"]
	returned, hasReturn := positions["return"]
	cancelled, hasCancelled := positions["cancelled"]
	if !hasShipped || !hasReturn || !hasCancelled {
		t.Fatalf("state-machine grid elements missing: %#v", positions)
	}
	shippedCenter := shipped.X + shipped.Width/2
	returnCenter := returned.X + returned.Width/2
	cancelledCenter := cancelled.X + cancelled.Width/2
	if math.Abs(shippedCenter-returnCenter) > 2 || math.Abs(returnCenter-cancelledCenter) > 2 {
		t.Fatalf("related state-machine columns not aligned: shipped=%.1f return=%.1f cancelled=%.1f", shippedCenter, returnCenter, cancelledCenter)
	}
}

func TestUMLStateMachineContainerRowsAndColumnsPlaceStates(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="1080" height="620"><uml id="state"><state-machine-diagram direction="right"><container><row><col><initial id="start"/></col><col><state id="paid" title="Paid"/></col><col><state id="shipped" title="Shipped"/></col></row><row><col></col><col></col><col><state id="return" title="Return"/></col></row><row><col></col><col></col><col><state id="cancelled" title="Cancelled"/></col></row></container><transition src="start" dst="paid"/><transition src="paid" dst="shipped"/><transition src="shipped" dst="return"/><transition src="return" dst="cancelled"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	positions := umlElementPositionsV1UMLTest(t, rawScene)
	paid, hasPaid := positions["paid"]
	shipped, hasShipped := positions["shipped"]
	returned, hasReturn := positions["return"]
	cancelled, hasCancelled := positions["cancelled"]
	if !hasPaid || !hasShipped || !hasReturn || !hasCancelled {
		t.Fatalf("container row/col state-machine elements missing: %#v", positions)
	}
	shippedCenter := shipped.X + shipped.Width/2
	returnCenter := returned.X + returned.Width/2
	cancelledCenter := cancelled.X + cancelled.Width/2
	if !(returned.Y > shipped.Y+100 && cancelled.Y > returned.Y+100) {
		t.Fatalf("container rows not separated: shipped y=%.1f return y=%.1f cancelled y=%.1f", shipped.Y, returned.Y, cancelled.Y)
	}
	if math.Abs(shippedCenter-returnCenter) > 2 || math.Abs(returnCenter-cancelledCenter) > 2 || !(shipped.X > paid.X) {
		t.Fatalf("container columns not aligned: paid=%.1f shipped=%.1f return=%.1f cancelled=%.1f", paid.X+paid.Width/2, shippedCenter, returnCenter, cancelledCenter)
	}
}

func TestUMLStateMachineConnectorsAvoidIntermediateStates(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="360"><uml id="state"><state-machine-diagram direction="right"><state id="left" title="Left" row="1" col="1"/><state id="middle" title="Middle" row="1" col="2"/><state id="right" title="Right" row="1" col="3"/><transition src="left" dst="right" event="skip middle"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var middle, arrow *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.UMLLocalID == "middle" && element.Type != "text" {
			middle = element
		}
		if element.Type == "arrow" && element.CustomData.UMLRelationSourceReference == "left" && element.CustomData.UMLRelationDestinationReference == "right" {
			arrow = element
		}
	}
	if middle == nil || arrow == nil || len(arrow.Points) < 2 {
		t.Fatalf("intermediate state or arrow missing: middle=%#v arrow=%#v", middle, arrow)
	}
	for index := 0; index < len(arrow.Points)-1; index++ {
		start, end := absoluteArrowSegmentV1UMLTest(arrow, index)
		if segmentIntersectsRectV1UMLTest(start[0], start[1], end[0], end[1], middle.X, middle.Y, middle.Width, middle.Height) {
			t.Fatalf("state-machine connector crosses intermediate state: segment=%#v->%#v middle=%#v arrow=%#v", start, end, middle, arrow)
		}
	}
}

func TestUMLStateMachineBentConnectorsAvoidIntermediateStates(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="520"><uml id="state"><state-machine-diagram direction="right"><state id="left" title="Left" row="1" col="1"/><state id="middle" title="Middle" row="1" col="2"/><state id="right" title="Right" row="1" col="3"/><transition src="left" dst="right" event="skip middle"><bend x="480" y="300"/></transition></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	middle, arrow := umlStateAndArrowV1UMLTest(t, rawScene, "middle", "left", "right")
	assertArrowDoesNotCrossRectV1UMLTest(t, arrow, middle)
}

func TestUMLStateMachineBentRouteAvoidanceStaysInsideFrame(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="520"><uml id="state"><state-machine-diagram direction="right"><state id="left" title="Left" row="1" col="1"/><state id="middle" title="Middle" row="1" col="2"/><state id="right" title="Right" row="1" col="3"/><state id="cancelled" title="Cancelled" row="3" col="3"/><transition src="left" dst="cancelled" event="outside"><bend x="920" y="170"/><bend x="920" y="430"/></transition></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	middle, arrow := umlStateAndArrowV1UMLTest(t, rawScene, "middle", "left", "cancelled")
	assertArrowDoesNotCrossRectV1UMLTest(t, arrow, middle)
	assertArrowInsideFrameV1UMLTest(t, arrow, 0, 0, 960, 520)
}

func TestUMLStateMachineDistantConnectorsUseOuterDetours(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="1200" height="620"><uml id="state"><state-machine-diagram direction="right"><container><row><col><state id="source" title="Source"/></col><col><state id="top-a" title="Top A"/></col><col><state id="top-b" title="Top B"/></col><col><state id="top-c" title="Top C"/></col><col><state id="top-d" title="Top D"/></col></row><row><col/><col><state id="mid-a" title="Mid A"/></col><col><state id="mid-b" title="Mid B"/></col><col><state id="mid-c" title="Mid C"/></col><col/></row><row><col/><col/><col/><col/><col><state id="destination" title="Destination"/></col></row></container><transition src="source" dst="destination" event="far"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	_, arrow := umlStateAndArrowV1UMLTest(t, rawScene, "mid-b", "source", "destination")
	minY := math.Inf(1)
	maxY := math.Inf(-1)
	for _, point := range arrow.Points {
		y := arrow.Y + point[1]
		minY = math.Min(minY, y)
		maxY = math.Max(maxY, y)
	}
	if minY >= 80 && maxY <= 540 {
		t.Fatalf("distant state-machine connector should use an outer detour: arrow=%#v minY=%.1f maxY=%.1f", arrow, minY, maxY)
	}
}

func TestUMLStateMachineSampleSVGRoutesStayInsideFrameAndAvoidStates(t *testing.T) {
	source, err := os.ReadFile(filepath.Join(repoRoot(t), "docs", "src", "examples", "samples", "uml-state-machine.xal"))
	if err != nil {
		t.Fatal(err)
	}
	svg, err := newUsecase().RenderSVG(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderSVG() error = %v", err)
	}
	paths := svgConnectorPathsV1UMLTest(t, string(svg))
	stateRects := svgStateBodyRectsV1UMLTest(t, string(svg))
	assertSVGConnectorPathsInsideFrameV1UMLTest(t, paths, 0, 0, 1680, 900)
	assertSVGConnectorPathsAvoidRectsV1UMLTest(t, paths, stateRects)
}

func TestUMLStateMachineConceptLabelsReachEditableScene(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="520"><uml id="state"><state-machine-diagram direction="right"><initial id="start"/><state id="processing" title="Processing"><entry>reserve stock</entry><do>pack order</do><internal>timeout / notify operator</internal><exit>publish event</exit><region>fulfilment</region></state><choice id="result" title="Result"/><final id="done"/><transition src="start" dst="processing" event="created"/><transition src="processing" dst="result" event="paymentCaptured" guard="stock available" action="ship"/><transition src="result" dst="done" guard="ok"/><transition src="result" dst="processing" guard="retry"/></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	if !strings.Contains(string(rawScene), `"text": "paymentCaptured [stock available] / ship"`) || !strings.Contains(string(rawScene), `"xaligoUmlAction": "ship"`) {
		t.Fatalf("transition concept label or metadata missing: %s", rawScene)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var choice entity.Element
	for _, element := range scene.Elements {
		if element.CustomData != nil && element.CustomData.UMLLocalID == "result" && element.CustomData.UMLElementKind == "choice" {
			choice = element
		}
	}
	for _, want := range []string{"Processing", "entry", "reserve stock", "do", "pack order", "internal", "timeout / notify operator", "exit", "publish event", "region", "fulfilment"} {
		if !strings.Contains(string(rawScene), `"text": "`+want+`"`) {
			t.Fatalf("state concept text %q missing: %s", want, rawScene)
		}
	}
	for _, marker := range []string{"xaligoUmlStateHeader", "xaligoUmlStateRowDivider", "xaligoUmlStateColumnDivider"} {
		if !strings.Contains(string(rawScene), marker) {
			t.Fatalf("state structured marker %q missing: %s", marker, rawScene)
		}
	}
	if choice.BackgroundColor != "#ffffff" || choice.StrokeColor != "#052d6e" {
		t.Fatalf("choice palette = %q/%q, want white body with class-theme stroke", choice.StrokeColor, choice.BackgroundColor)
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

func TestUMLRelationBendsReachEditableScene(t *testing.T) {
	for _, test := range []struct {
		name         string
		source       []byte
		relationKind string
	}{
		{
			name:         "class dependency",
			relationKind: "dependency",
			source:       []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="360"><uml id="class"><class-diagram><class id="service" title="Service"/><class id="repo" title="Repository"/><dependency src="service" dst="repo" title="uses"><bend x="360" y="80"/><bend x="360" y="260"/></dependency></class-diagram></uml></frame></frames></xaligo>`),
		},
		{
			name:         "activity control flow",
			relationKind: "control-flow",
			source:       []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="360"><uml id="activity"><activity-diagram direction="right"><initial id="start"/><action id="review" title="Review"/><final id="done"/><control-flow src="start" dst="review"><bend x="240" y="80"/></control-flow><control-flow src="review" dst="done"><bend x="480" y="280"/></control-flow></activity-diagram></uml></frame></frames></xaligo>`),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			rawScene, err := newUsecase().RenderExcalidraw(context.Background(), test.source, entity.RenderOptions{PxPerInch: 96})
			if err != nil {
				t.Fatalf("RenderExcalidraw() error = %v", err)
			}
			if !strings.Contains(string(rawScene), `"xaligoConnectorBends":`) {
				t.Fatalf("UML relation bend metadata missing: %s", rawScene)
			}
			var scene entity.PresentationScene
			if err := json.Unmarshal(rawScene, &scene); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			var arrow *entity.Element
			for index := range scene.Elements {
				element := &scene.Elements[index]
				if element.Type == "arrow" && element.CustomData != nil && element.CustomData.UMLRelationKind == test.relationKind && element.CustomData.ConnectorBends != "" {
					arrow = element
					break
				}
			}
			if arrow == nil || len(arrow.Points) < 4 {
				t.Fatalf("bent UML relation arrow not found: %#v", scene.Elements)
			}
		})
	}
}

func TestUMLRelationLabelsAvoidEndpointItems(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="360"><uml id="state"><state-machine-diagram direction="right"><state id="left" title="Left" row="1" col="1"/><state id="right" title="Right" row="1" col="2"/><transition src="left" dst="right" event="label" action="that should avoid boxes"><bend x="360" y="180"/></transition></state-machine-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var left, right, label *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil {
			continue
		}
		switch {
		case element.CustomData.UMLLocalID == "left" && element.Type != "text":
			left = element
		case element.CustomData.UMLLocalID == "right" && element.Type != "text":
			right = element
		case element.Type == "text" && element.CustomData.UMLRelationLabel != "":
			label = element
		}
	}
	if left == nil || right == nil || label == nil {
		t.Fatalf("state label test elements missing: left=%#v right=%#v label=%#v", left, right, label)
	}
	for _, item := range []*entity.Element{left, right} {
		if rectsOverlapV1UMLTest(label.X, label.Y, label.Width, label.Height, item.X, item.Y, item.Width, item.Height) {
			t.Fatalf("UML relation label overlaps endpoint item: label=%#v item=%#v", label, item)
		}
	}
}

func rectsOverlapV1UMLTest(ax, ay, aw, ah, bx, by, bw, bh float64) bool {
	return ax < bx+bw && ax+aw > bx && ay < by+bh && ay+ah > by
}

func umlElementPositionsV1UMLTest(t *testing.T, rawScene []byte) map[string]entity.Element {
	t.Helper()
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	positions := map[string]entity.Element{}
	for _, element := range scene.Elements {
		if element.CustomData == nil || element.CustomData.UMLLocalID == "" || element.CustomData.UMLElementKind == "" || element.Type == "text" {
			continue
		}
		positions[element.CustomData.UMLLocalID] = element
	}
	return positions
}

func umlStateAndArrowV1UMLTest(t *testing.T, rawScene []byte, stateID, src, dst string) (*entity.Element, *entity.Element) {
	t.Helper()
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var state, arrow *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.UMLLocalID == stateID && element.Type != "text" {
			state = element
		}
		if element.Type == "arrow" && element.CustomData.UMLRelationSourceReference == src && element.CustomData.UMLRelationDestinationReference == dst {
			arrow = element
		}
	}
	if state == nil || arrow == nil || len(arrow.Points) < 2 {
		t.Fatalf("state or arrow missing: state=%#v arrow=%#v", state, arrow)
	}
	return state, arrow
}

func assertArrowDoesNotCrossRectV1UMLTest(t *testing.T, arrow, rect *entity.Element) {
	t.Helper()
	for index := 0; index < len(arrow.Points)-1; index++ {
		start, end := absoluteArrowSegmentV1UMLTest(arrow, index)
		if segmentIntersectsRectV1UMLTest(start[0], start[1], end[0], end[1], rect.X, rect.Y, rect.Width, rect.Height) {
			t.Fatalf("connector crosses state: segment=%#v->%#v rect=%#v arrow=%#v", start, end, rect, arrow)
		}
	}
}

func assertArrowInsideFrameV1UMLTest(t *testing.T, arrow *entity.Element, frameX, frameY, frameW, frameH float64) {
	t.Helper()
	for index, point := range arrow.Points {
		x := arrow.X + point[0]
		y := arrow.Y + point[1]
		if x < frameX || x > frameX+frameW || y < frameY || y > frameY+frameH {
			t.Fatalf("arrow point %d is outside frame: point=(%.1f, %.1f) frame=(%.1f, %.1f, %.1f, %.1f) arrow=%#v", index, x, y, frameX, frameY, frameW, frameH, arrow)
		}
	}
}

type svgPointV1UMLTest struct {
	x float64
	y float64
}

type svgRectV1UMLTest struct {
	x float64
	y float64
	w float64
	h float64
}

func svgConnectorPathsV1UMLTest(t *testing.T, svg string) [][]svgPointV1UMLTest {
	t.Helper()
	pathRE := regexp.MustCompile(`<path d="([^"]+)"[^>]*marker-end="url\(#xaligo-triangle\)"`)
	pointRE := regexp.MustCompile(`[ML] (-?[0-9]+(?:\.[0-9]+)?) (-?[0-9]+(?:\.[0-9]+)?)`)
	matches := pathRE.FindAllStringSubmatch(svg, -1)
	paths := make([][]svgPointV1UMLTest, 0, len(matches))
	for _, match := range matches {
		pointMatches := pointRE.FindAllStringSubmatch(match[1], -1)
		points := make([]svgPointV1UMLTest, 0, len(pointMatches))
		for _, pointMatch := range pointMatches {
			x, err := strconv.ParseFloat(pointMatch[1], 64)
			if err != nil {
				t.Fatalf("parse SVG path x %q: %v", pointMatch[1], err)
			}
			y, err := strconv.ParseFloat(pointMatch[2], 64)
			if err != nil {
				t.Fatalf("parse SVG path y %q: %v", pointMatch[2], err)
			}
			points = append(points, svgPointV1UMLTest{x: x, y: y})
		}
		if len(points) >= 2 {
			paths = append(paths, points)
		}
	}
	if len(paths) == 0 {
		t.Fatalf("no SVG connector paths found")
	}
	return paths
}

func svgUMLComponentAssociationPathsV1UMLTest(t *testing.T, svg string) [][]svgPointV1UMLTest {
	t.Helper()
	pathRE := regexp.MustCompile(`<path d="([^"]+)"[^>]*stroke="#1E1E1E"`)
	pointRE := regexp.MustCompile(`[ML] (-?[0-9]+(?:\.[0-9]+)?) (-?[0-9]+(?:\.[0-9]+)?)`)
	matches := pathRE.FindAllStringSubmatch(svg, -1)
	paths := make([][]svgPointV1UMLTest, 0, len(matches))
	for _, match := range matches {
		pointMatches := pointRE.FindAllStringSubmatch(match[1], -1)
		points := make([]svgPointV1UMLTest, 0, len(pointMatches))
		for _, pointMatch := range pointMatches {
			x, err := strconv.ParseFloat(pointMatch[1], 64)
			if err != nil {
				t.Fatalf("parse SVG component path x %q: %v", pointMatch[1], err)
			}
			y, err := strconv.ParseFloat(pointMatch[2], 64)
			if err != nil {
				t.Fatalf("parse SVG component path y %q: %v", pointMatch[2], err)
			}
			points = append(points, svgPointV1UMLTest{x: x, y: y})
		}
		if len(points) >= 2 {
			paths = append(paths, points)
		}
	}
	if len(paths) == 0 {
		t.Fatalf("no SVG component association paths found")
	}
	return paths
}

func svgStateBodyRectsV1UMLTest(t *testing.T, svg string) []svgRectV1UMLTest {
	t.Helper()
	rectRE := regexp.MustCompile(`<rect x="([0-9]+(?:\.[0-9]+)?)" y="([0-9]+(?:\.[0-9]+)?)" width="([0-9]+(?:\.[0-9]+)?)" height="([0-9]+(?:\.[0-9]+)?)" fill="#FFFFFF" fill-opacity="1" stroke="#052D6E" stroke-width="1\.35"`)
	matches := rectRE.FindAllStringSubmatch(svg, -1)
	rects := make([]svgRectV1UMLTest, 0, len(matches))
	for _, match := range matches {
		values := make([]float64, 4)
		for i := range values {
			value, err := strconv.ParseFloat(match[i+1], 64)
			if err != nil {
				t.Fatalf("parse SVG state rect value %q: %v", match[i+1], err)
			}
			values[i] = value
		}
		rects = append(rects, svgRectV1UMLTest{x: values[0], y: values[1], w: values[2], h: values[3]})
	}
	if len(rects) == 0 {
		t.Fatalf("no SVG state body rectangles found")
	}
	return rects
}

func assertSVGConnectorPathsInsideFrameV1UMLTest(t *testing.T, paths [][]svgPointV1UMLTest, frameX, frameY, frameW, frameH float64) {
	t.Helper()
	for pathIndex, path := range paths {
		for pointIndex, point := range path {
			if point.x < frameX || point.x > frameX+frameW || point.y < frameY || point.y > frameY+frameH {
				t.Fatalf("SVG connector point outside frame: path=%d point=%d value=(%.1f, %.1f) frame=(%.1f, %.1f, %.1f, %.1f) path=%#v", pathIndex, pointIndex, point.x, point.y, frameX, frameY, frameW, frameH, path)
			}
		}
	}
}

func assertSVGConnectorPathsAvoidRectsV1UMLTest(t *testing.T, paths [][]svgPointV1UMLTest, rects []svgRectV1UMLTest) {
	t.Helper()
	for pathIndex, path := range paths {
		for segmentIndex := 0; segmentIndex < len(path)-1; segmentIndex++ {
			start := path[segmentIndex]
			end := path[segmentIndex+1]
			for rectIndex, rect := range rects {
				if segmentIntersectsRectV1UMLTest(start.x, start.y, end.x, end.y, rect.x, rect.y, rect.w, rect.h) {
					t.Fatalf("SVG connector crosses state body: path=%d segment=%d rect=%d start=(%.1f, %.1f) end=(%.1f, %.1f) rect=(%.1f, %.1f, %.1f, %.1f) path=%#v", pathIndex, segmentIndex, rectIndex, start.x, start.y, end.x, end.y, rect.x, rect.y, rect.w, rect.h, path)
				}
			}
		}
	}
}

func absoluteArrowSegmentV1UMLTest(arrow *entity.Element, index int) ([2]float64, [2]float64) {
	start := arrow.Points[index]
	end := arrow.Points[index+1]
	return [2]float64{arrow.X + start[0], arrow.Y + start[1]}, [2]float64{arrow.X + end[0], arrow.Y + end[1]}
}

func segmentIntersectsRectV1UMLTest(x1, y1, x2, y2, rx, ry, rw, rh float64) bool {
	const tolerance = 0.5
	rx -= tolerance
	ry -= tolerance
	rw += tolerance * 2
	rh += tolerance * 2
	if math.Abs(x1-x2) <= tolerance {
		x := x1
		if x < rx || x > rx+rw {
			return false
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		return y2 >= ry && y1 <= ry+rh
	}
	if math.Abs(y1-y2) <= tolerance {
		y := y1
		if y < ry || y > ry+rh {
			return false
		}
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		return x2 >= rx && x1 <= rx+rw
	}
	return false
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
	var selfMessage *entity.Element
	var selfMessageLabel *entity.Element
	var coveringActivation *entity.Element
	selfMessagePoints := [][]float64{}
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil || element.CustomData.UMLMessageOrder == "" {
			continue
		}
		if element.Type == "arrow" && element.StartBinding != nil && element.EndBinding != nil {
			anchors[element.CustomData.UMLMessageOrder] = [2]float64{element.StartBinding.FixedPoint[1], element.EndBinding.FixedPoint[1]}
			if element.CustomData.UMLMessageOrder == "2" {
				selfMessage = element
				selfMessagePoints = element.Points
			}
		}
		if element.Type == "text" && element.CustomData.UMLMessageOrder == "2" {
			selfMessageLabel = element
		}
		if element.CustomData.UMLSequenceActivation && element.CustomData.UMLSequenceActivationOwner == "b" {
			coveringActivation = element
		}
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
	if len(selfMessagePoints) < 4 || selfMessagePoints[0][0] <= 0 || selfMessagePoints[1][0] < 96 || math.Abs(selfMessagePoints[len(selfMessagePoints)-1][0]-selfMessagePoints[0][0]) > 0.01 {
		t.Fatalf("self-message should route as a right-side loop, got %#v", selfMessagePoints)
	}
	if selfMessage == nil || selfMessageLabel == nil || coveringActivation == nil {
		t.Fatalf("self-message, label, or activation missing: arrow=%#v label=%#v activation=%#v", selfMessage, selfMessageLabel, coveringActivation)
	}
	startX := selfMessage.X + selfMessagePoints[0][0]
	endX := selfMessage.X + selfMessagePoints[len(selfMessagePoints)-1][0]
	activationRight := coveringActivation.X + coveringActivation.Width
	if math.Abs(startX-activationRight) > 0.01 || math.Abs(endX-activationRight) > 0.01 {
		t.Fatalf("self-message endpoints should align with activation right edge: start=%v end=%v activation=%#v", startX, endX, coveringActivation)
	}
	startY := selfMessage.Y + selfMessagePoints[0][1]
	endY := selfMessage.Y + selfMessagePoints[len(selfMessagePoints)-1][1]
	activationTop := coveringActivation.Y
	activationBottom := coveringActivation.Y + coveringActivation.Height
	if startY < activationTop || startY > activationBottom || endY < activationTop || endY > activationBottom {
		t.Fatalf("self-message endpoints should stay within activation vertical range: startY=%v endY=%v activation=%#v", startY, endY, coveringActivation)
	}
	loopTopY := selfMessage.Y + selfMessagePoints[0][1]
	if selfMessageLabel.Y+selfMessageLabel.Height > loopTopY-4 {
		t.Fatalf("self-message label overlaps loop: arrow=%#v label=%#v", selfMessage, selfMessageLabel)
	}
}

func TestUMLSequenceSelfMessageSVGAlignsWithActivationBar(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="customer"/><lifeline id="session" title="Session"/><message src="customer" dst="session" order="1" title="checkout"/><message src="session" dst="session" order="1.1" title="validateCart()"/></sequence-diagram></uml></frame></frames></xaligo>`)
	svg, err := newUsecase().RenderSVG(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderSVG() error = %v", err)
	}
	output := string(svg)
	lifelineIndex := strings.Index(output, `stroke-dasharray="8 6"`)
	activationIndex := strings.Index(output, `width="16"`)
	if lifelineIndex < 0 || activationIndex < 0 || lifelineIndex > activationIndex {
		t.Fatalf("activation bar should render in front of dashed lifeline: %s", output)
	}
	pathMatch := regexp.MustCompile(`<path d="M ([0-9.]+) [0-9.]+ L [0-9.]+ [0-9.]+ L [0-9.]+ [0-9.]+ L ([0-9.]+) [0-9.]+"[^>]*marker-end="url\(#xaligo-triangle\)"`).FindStringSubmatch(output)
	if pathMatch == nil {
		t.Fatalf("SVG missing self-message loop: %s", output)
	}
	startX, err := strconv.ParseFloat(pathMatch[1], 64)
	if err != nil {
		t.Fatalf("ParseFloat(start x) error = %v", err)
	}
	endX, err := strconv.ParseFloat(pathMatch[2], 64)
	if err != nil {
		t.Fatalf("ParseFloat(end x) error = %v", err)
	}
	activationMatches := regexp.MustCompile(`<rect x="([0-9.]+)" y="[0-9.]+" width="([0-9.]+)" height="[0-9.]+"`).FindAllStringSubmatch(output, -1)
	foundActivation := false
	for _, match := range activationMatches {
		activationX, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			t.Fatalf("ParseFloat(activation x) error = %v", err)
		}
		activationWidth, err := strconv.ParseFloat(match[2], 64)
		if err != nil {
			t.Fatalf("ParseFloat(activation width) error = %v", err)
		}
		activationRight := activationX + activationWidth
		if math.Abs(startX-activationRight) <= 0.01 && math.Abs(endX-activationRight) <= 0.01 {
			foundActivation = true
			break
		}
	}
	if !foundActivation {
		t.Fatalf("self-message endpoints should align with an activation right edge: start=%v end=%v activations=%#v", startX, endX, activationMatches)
	}
}

func TestUMLSequenceCallerActivationCoversChildMessageStart(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="customer"/><lifeline id="api"/><lifeline id="session"/><message src="customer" dst="api" order="1" title="checkout"/><create-message src="api" dst="session" order="1.1" title="create"/><return-message src="session" dst="api" order="1.2" title="created"/><return-message src="api" dst="customer" order="2" title="ok"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var childMessage *entity.Element
	var callerActivation *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil {
			continue
		}
		if element.Type == "arrow" && element.CustomData.UMLMessageOrder == "1.1" {
			childMessage = element
		}
		if element.CustomData.UMLSequenceActivation && element.CustomData.UMLSequenceActivationOwner == "api" && element.CustomData.UMLMessageOrder == "1" {
			callerActivation = element
		}
	}
	if childMessage == nil || callerActivation == nil || len(childMessage.Points) == 0 {
		t.Fatalf("child message or caller activation missing: message=%#v activation=%#v", childMessage, callerActivation)
	}
	startY := childMessage.Y + childMessage.Points[0][1]
	if startY < callerActivation.Y || startY > callerActivation.Y+callerActivation.Height {
		t.Fatalf("child message start should stay within caller activation range: startY=%v activation=%#v", startY, callerActivation)
	}
}

func TestUMLSequenceActivationCoversReturnAndCleanupMessages(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="520"><uml id="sequence"><sequence-diagram><participant id="customer"/><lifeline id="api"/><lifeline id="session"/><message src="customer" dst="api" order="1" title="checkout"/><create-message src="api" dst="session" order="1.1" title="create"/><message src="session" dst="session" order="1.2" title="validate"/><return-message src="session" dst="api" order="1.3" title="receipt"/><message src="api" dst="session" order="1.4" title="release"/><return-message src="api" dst="customer" order="2" title="ok"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	activations := map[string]*entity.Element{}
	messages := map[string]*entity.Element{}
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.UMLSequenceActivation {
			key := element.CustomData.UMLSequenceActivationOwner + ":" + element.CustomData.UMLMessageOrder
			activations[key] = element
		}
		if element.Type == "arrow" && element.CustomData.UMLMessageOrder != "" {
			messages[element.CustomData.UMLMessageOrder] = element
		}
	}
	apiActivation := activations["api:1"]
	sessionActivation := activations["session:1.1"]
	if apiActivation == nil || sessionActivation == nil {
		t.Fatalf("missing activations: %#v", activations)
	}
	assertMessageEndpointWithinActivation := func(order string, pointIndex int, activation *entity.Element) {
		t.Helper()
		message := messages[order]
		if message == nil || len(message.Points) <= pointIndex {
			t.Fatalf("missing message %s: %#v", order, messages)
		}
		y := message.Y + message.Points[pointIndex][1]
		if y < activation.Y || y > activation.Y+activation.Height {
			t.Fatalf("message %s point %d should be within activation: y=%v activation=%#v", order, pointIndex, y, activation)
		}
	}
	assertMessageEndpointWithinActivation("1.4", 0, apiActivation)
	assertMessageEndpointWithinActivation("2", 0, apiActivation)
	assertMessageEndpointWithinActivation("1.3", 0, sessionActivation)
}

func TestUMLSequenceSuppressesContainedActivationBars(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="960" height="520"><uml id="sequence"><sequence-diagram><participant id="customer"/><lifeline id="session"/><message src="customer" dst="session" order="1" title="checkout"/><message src="session" dst="session" order="1.1" title="validate"/><return-message src="session" dst="customer" order="2" title="ok"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	activations := 0
	var activation *entity.Element
	var selfMessage *entity.Element
	for index := range scene.Elements {
		element := &scene.Elements[index]
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.UMLSequenceActivation && element.CustomData.UMLSequenceActivationOwner == "session" {
			activations++
			activation = element
		}
		if element.Type == "arrow" && element.CustomData.UMLMessageOrder == "1.1" {
			selfMessage = element
		}
	}
	if activations != 1 || activation == nil || selfMessage == nil || len(selfMessage.Points) == 0 {
		t.Fatalf("expected one covering activation and self-message, activations=%d activation=%#v message=%#v", activations, activation, selfMessage)
	}
	y := selfMessage.Y + selfMessage.Points[0][1]
	if y < activation.Y || y > activation.Y+activation.Height {
		t.Fatalf("self-message should stay inside covering activation: y=%v activation=%#v", y, activation)
	}
}

func TestUMLSequenceMessagesRenderActivationBars(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="user" title="User"/><lifeline id="api" title="API"/><lifeline id="worker" title="Worker"/><message src="user" dst="api" order="1" title="submit()"/><create-message src="api" dst="worker" order="2" title="create"/><return-message src="worker" dst="api" order="3" title="ok"/><destroy-message src="api" dst="worker" order="4" title="destroy"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	activations := map[string]int{}
	for _, element := range scene.Elements {
		if element.CustomData == nil || !element.CustomData.UMLSequenceActivation {
			continue
		}
		activations[element.CustomData.UMLSequenceActivationOwner]++
		if element.Width < 8 || element.Width > 18 || element.Height < 36 {
			t.Fatalf("activation bar has unexpected geometry: %#v", element)
		}
	}
	if activations["api"] != 1 || activations["worker"] != 2 || activations["user"] != 0 {
		t.Fatalf("activation bars = %#v", activations)
	}
}

func TestUMLSequenceMessagesUseResponseAndStopNotation(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="user" title="User"/><lifeline id="api" title="API"/><lifeline id="worker" title="Worker"/><message src="user" dst="api" order="1" title="submit()"/><create-message src="api" dst="worker" order="2" title="create"/><return-message src="worker" dst="api" order="3" title="ok"/><destroy-message src="api" dst="worker" order="4" title="destroy"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	returnDashed := false
	stopMarks := 0
	for _, element := range scene.Elements {
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.UMLRelationKind == "return-message" && element.Type == "arrow" && element.StrokeStyle == "dashed" {
			returnDashed = true
		}
		if element.CustomData.UMLSequenceStop {
			stopMarks++
			if element.CustomData.UMLSequenceStopOwner != "worker" || element.StrokeColor != "#052d6e" {
				t.Fatalf("destroy stop marker = %#v", element)
			}
		}
	}
	if !returnDashed || stopMarks != 2 {
		t.Fatalf("return dashed = %t, stop marks = %d, scene = %#v", returnDashed, stopMarks, scene.Elements)
	}
}

func TestUMLSequenceMessagesDistinguishSyncAndAsyncNotation(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="client" title="Client"/><lifeline id="filterChain" title="FilterChain"/><lifeline id="filter" title="Filter"/><message src="client" dst="filterChain" order="1" title="request"/><message src="filterChain" dst="filter" order="2" title="doFilter" mode="async"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	modes := map[string]string{}
	heads := map[string]string{}
	for _, element := range scene.Elements {
		if element.Type != "arrow" || element.CustomData == nil || element.CustomData.UMLMessageOrder == "" {
			continue
		}
		modes[element.CustomData.UMLMessageOrder] = element.CustomData.UMLMessageMode
		heads[element.CustomData.UMLMessageOrder] = element.CustomData.ConnectorEndArrowhead
	}
	if modes["1"] != "sync" || heads["1"] != "triangle" || modes["2"] != "async" || heads["2"] != "arrow" {
		t.Fatalf("sequence message notation modes = %#v heads = %#v", modes, heads)
	}
}

func TestUMLSequenceParticipantsRenderAsLifelines(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420"><uml id="sequence"><sequence-diagram><participant id="client" title="Client"/><lifeline id="filterChain" title="FilterChain"/><lifeline id="filter" title="Filter"/><message src="client" dst="filterChain" order="1" title="request"/><message src="filterChain" dst="filter" order="2" title="doFilter" mode="async"/></sequence-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	headers := 0
	lines := 0
	for _, element := range scene.Elements {
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.UMLSequenceLifelineHeader {
			headers++
			if element.Height > 40 || element.StrokeColor != "#052d6e" || element.BackgroundColor != "#08b8ea" {
				t.Fatalf("lifeline header too tall: %#v", element)
			}
		}
		if element.CustomData.UMLSequenceLifeline {
			lines++
			if element.Type != "line" || element.StrokeStyle != "dashed" || element.StrokeColor != "#052d6e" || element.Width > 1 {
				t.Fatalf("lifeline line = %#v", element)
			}
		}
		if (element.CustomData.UMLElementKind == "participant" || element.CustomData.UMLElementKind == "lifeline") && element.Height > 120 {
			t.Fatalf("sequence participant rendered as a full box: %#v", element)
		}
	}
	if headers != 3 || lines != 3 {
		t.Fatalf("lifeline headers = %d lines = %d, want 3 each", headers, lines)
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
	if classElement.StrokeColor != "#052d6e" || classElement.BackgroundColor != "#ffffff" || classElement.StrokeWidth < 1.3 {
		t.Fatalf("class style = stroke %q background %q width %.2f, want xaligo activity palette", classElement.StrokeColor, classElement.BackgroundColor, classElement.StrokeWidth)
	}
	if classElement.Width > 260 || classElement.Height > 180 {
		t.Fatalf("class box size = %.1fx%.1f, want compact Lucid-like classifier", classElement.Width, classElement.Height)
	}
	var raw map[string]any
	if err := json.Unmarshal(rawScene, &raw); err != nil {
		t.Fatalf("json.Unmarshal(raw map) error = %v", err)
	}
	foundHeader := false
	foundHeaderText := false
	foundAttributeText := false
	foundOperationText := false
	foundBodyDivider := false
	for _, rawElement := range raw["elements"].([]any) {
		element, _ := rawElement.(map[string]any)
		customData, _ := element["customData"].(map[string]any)
		if customData["xaligoUmlClassHeader"] == true && element["backgroundColor"] == "#08b8ea" {
			foundHeader = true
		}
		if customData["xaligoUmlClassHeaderContent"] == true && element["strokeColor"] == "#ffffff" {
			if element["text"] == "<<service>>\n{abstract, static} Repository" {
				foundHeaderText = true
			}
		}
		if customData["xaligoUmlClassAttributeContent"] == true && element["strokeColor"] == "#052d6e" {
			if element["text"] == "- store: Store" {
				foundAttributeText = true
			}
		}
		if customData["xaligoUmlClassOperationContent"] == true && element["strokeColor"] == "#052d6e" {
			if element["text"] == "+ find(id): Entity" {
				foundOperationText = true
			}
		}
		if customData["xaligoUmlClassBodyDivider"] == true {
			foundBodyDivider = true
		}
	}
	if !foundHeader || !foundHeaderText || !foundAttributeText || !foundOperationText || !foundBodyDivider {
		t.Fatalf("class compartment rendering missing header=%t headerText=%t attributeText=%t operationText=%t bodyDivider=%t: %s", foundHeader, foundHeaderText, foundAttributeText, foundOperationText, foundBodyDivider, rawScene)
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

func TestUMLClassDiagramSupportsPackageGroups(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="900" height="520"><uml id="classes"><class-diagram grid="1"><package id="identity" title="Identity" grid="2"><class id="user" title="User"><attribute>- id: int</attribute><operation>+ login()</operation></class><class id="session" title="Session"><attribute>- token: string</attribute></class><association src="user" dst="session" label="1 -> *"/></package><package id="billing" title="Billing"><class id="invoice" title="Invoice"><attribute>- id: UUID</attribute><operation>+ issue()</operation></class></package></class-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	classCount := 0
	packageCount := 0
	foundRelation := false
	var packageWidths []float64
	for _, element := range scene.Elements {
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.UMLElementKind == "class" {
			classCount++
		}
		if element.CustomData.GroupBorder {
			packageCount++
			packageWidths = append(packageWidths, element.Width)
		}
		if element.CustomData.UMLRelationKind == "association" {
			foundRelation = true
		}
	}
	if classCount != 3 || packageCount != 2 || !foundRelation {
		t.Fatalf("package class diagram scene = classes %d packages %d relation %t: %#v", classCount, packageCount, foundRelation, scene.Elements)
	}
	for _, width := range packageWidths {
		if width < 700 {
			t.Fatalf("package width %.1f is not optimized to its grid cell: %#v", width, scene.Elements)
		}
	}
}

func TestUMLClassPackagesAutoBalanceToFrameArea(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="1200" height="520"><uml id="classes"><class-diagram><package id="identity" title="Identity"><class id="user" title="User" /></package><package id="billing" title="Billing"><class id="invoice" title="Invoice" /></package><package id="shipping" title="Shipping"><class id="shipment" title="Shipment" /></package></class-diagram></uml></frame></frames></xaligo>`)
	rawScene, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderExcalidraw() error = %v", err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(rawScene, &scene); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	var packageElements []entity.Element
	for _, element := range scene.Elements {
		if element.CustomData != nil && element.CustomData.GroupBorder {
			packageElements = append(packageElements, element)
		}
	}
	if len(packageElements) != 3 {
		t.Fatalf("package count = %d, want 3: %#v", len(packageElements), scene.Elements)
	}
	baseY := packageElements[0].Y
	baseWidth := packageElements[0].Width
	baseHeight := packageElements[0].Height
	for _, element := range packageElements {
		if math.Abs(element.Y-baseY) > 1 || math.Abs(element.Width-baseWidth) > 1 || math.Abs(element.Height-baseHeight) > 1 {
			t.Fatalf("packages should share one balanced row and equal size: %#v", packageElements)
		}
		if element.Height < 430 {
			t.Fatalf("package height %.1f should use the frame height: %#v", element.Height, packageElements)
		}
	}
}

func TestUMLClassPackageSVGKeepsReadableCompartmentText(t *testing.T) {
	source := []byte(`<xaligo version="1"><data></data><frames><frame id="main" width="1440" height="760"><uml id="classes"><class-diagram><package id="domain" title="Domain"><class id="order" title="Order" stereotype="aggregate-root"><attribute>- id: UUID</attribute><attribute>- status: OrderStatus</attribute><operation>+ confirm()</operation><operation>+ total(): Money</operation></class><class id="customer" title="Customer"><attribute>- id: UUID</attribute><attribute>- name: String</attribute><operation>+ placeOrder(): Order</operation></class><class id="premium" title="PremiumCustomer" abstract="true"><attribute>- discountRate: Decimal</attribute><operation>+ calculateDiscount(): Money</operation></class></package><association src="customer" dst="order" src-multiplicity="1" dst-multiplicity="0..*"/></class-diagram></uml></frame></frames></xaligo>`)
	svg, err := newUsecase().RenderSVG(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatalf("RenderSVG() error = %v", err)
	}
	fontSizePattern := regexp.MustCompile(`font-size="([0-9]+(?:\.[0-9]+)?)"`)
	for _, match := range fontSizePattern.FindAllStringSubmatch(string(svg), -1) {
		fontSize, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			t.Fatalf("ParseFloat(%q) error = %v", match[1], err)
		}
		if fontSize < 12 {
			t.Fatalf("SVG contains unreadably small font-size %.3f: %s", fontSize, svg)
		}
	}
	if !strings.Contains(string(svg), "+ calculateDiscount(): Money") {
		t.Fatalf("SVG should keep long operation on one readable line: %s", svg)
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
