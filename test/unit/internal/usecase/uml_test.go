package usecase_test

import (
	"context"
	"encoding/json"
	"math"
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
