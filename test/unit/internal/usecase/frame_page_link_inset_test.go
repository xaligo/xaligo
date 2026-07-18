package usecase_test

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestCrossFramePageTerminalsUseResolvedMetadataRowGapInset(t *testing.T) {
	tests := []struct {
		name                   string
		metadata               string
		rowGap                 float64
		sourceFrameAnchor      string
		sourceSide             string
		destinationFrameAnchor string
		destinationSide        string
	}{
		{
			name:                   "metadata absent uses default on right and left",
			rowGap:                 4,
			sourceFrameAnchor:      "right-2",
			sourceSide:             "right",
			destinationFrameAnchor: "left-4",
			destinationSide:        "left",
		},
		{
			name:                   "custom row gap insets top terminals",
			metadata:               `<metadata position="bottom" row-gap="12" width="100" key-width="30" />`,
			rowGap:                 12,
			sourceFrameAnchor:      "top-2",
			sourceSide:             "top",
			destinationFrameAnchor: "top-4",
			destinationSide:        "top",
		},
		{
			name:                   "custom row gap insets right and left terminals",
			metadata:               `<metadata position="top" row-gap="12" width="100" key-width="30" />`,
			rowGap:                 12,
			sourceFrameAnchor:      "right-2",
			sourceSide:             "right",
			destinationFrameAnchor: "left-4",
			destinationSide:        "left",
		},
		{
			name:                   "custom row gap insets bottom terminals",
			metadata:               `<metadata position="top" row-gap="12" width="100" key-width="30" />`,
			rowGap:                 12,
			sourceFrameAnchor:      "bottom-2",
			sourceSide:             "bottom",
			destinationFrameAnchor: "bottom-4",
			destinationSide:        "bottom",
		},
		{
			name:                   "explicit zero keeps terminals on physical edge",
			metadata:               `<metadata position="bottom" row-gap="0" width="100" key-width="30" />`,
			rowGap:                 0,
			sourceFrameAnchor:      "top-2",
			sourceSide:             "top",
			destinationFrameAnchor: "top-4",
			destinationSide:        "top",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="300" height="220" margin="60">
    ` + test.metadata + `
    <rectangle id="node" title="Source" width="80" height="40" />
    <connection src="node" dst="destination.node"
                src-frame-anchor="` + test.sourceFrameAnchor + `"
                dst-frame-anchor="` + test.destinationFrameAnchor + `" />
  </frame>
  <frame id="destination" width="300" height="220" margin="60">
    ` + test.metadata + `
    <rectangle id="node" title="Destination" width="80" height="40" />
  </frame>
</frames></xaligo>`)
			scene := renderPageLinkInsetScene(t, source)
			sourceStub, destinationStub := crossFrameStubsPageLinkInsetTest(t, scene.Elements)

			sourceFrame := sceneElementRect(t, scene.Elements, "paper-frame-source")
			destinationFrame := sceneElementRect(t, scene.Elements, "paper-frame-destination")
			sourceStart, sourceTerminal := sceneArrowEndpoints(t, sourceStub)
			destinationTerminal, destinationEnd := sceneArrowEndpoints(t, destinationStub)

			assertPageLinkTerminalInsetTest(t, sourceFrame, sourceTerminal, test.sourceSide, test.rowGap, 0.3)
			assertPageLinkTerminalInsetTest(t, destinationFrame, destinationTerminal, test.destinationSide, test.rowGap, 0.7)
			assertPageLinkFrameApproachInsetTest(t, sourceStub, test.sourceSide, false)
			assertPageLinkFrameApproachInsetTest(t, destinationStub, test.destinationSide, true)
			assertPageLinkLabelClearsStub(t, scene.Elements, sourceStub, sourceStart, sourceTerminal)
			assertPageLinkLabelClearsStub(t, scene.Elements, destinationStub, destinationTerminal, destinationEnd)
			assertPageLinkLabelGap(t, sceneTextRectByValue(t, scene.Elements, "to <destination>"), sourceTerminal, test.sourceSide, 4)
			assertPageLinkLabelGap(t, sceneTextRectByValue(t, scene.Elements, "from <source>"), destinationTerminal, test.destinationSide, 4)
		})
	}
}

func TestMetadataRowGapDoesNotMoveSameFrameConnectionEndpoints(t *testing.T) {
	baselineScene, baselineArrow := renderSameFrameConnectionPageLinkInsetTest(t, "")
	customScene, customArrow := renderSameFrameConnectionPageLinkInsetTest(t, `<metadata position="top" row-gap="12" width="120" key-width="40" />`)

	baselineStart, baselineEnd := sceneArrowEndpoints(t, baselineArrow)
	customStart, customEnd := sceneArrowEndpoints(t, customArrow)
	if !pointsEqualPageLinkInsetTest(baselineStart, customStart) || !pointsEqualPageLinkInsetTest(baselineEnd, customEnd) {
		t.Fatalf("same-frame endpoints changed with metadata row-gap: baseline=%#v -> %#v custom=%#v -> %#v", baselineStart, baselineEnd, customStart, customEnd)
	}

	baselineSource := boundElementRectPageLinkInsetTest(t, baselineScene.Elements, baselineArrow, "startBinding")
	baselineDestination := boundElementRectPageLinkInsetTest(t, baselineScene.Elements, baselineArrow, "endBinding")
	customSource := boundElementRectPageLinkInsetTest(t, customScene.Elements, customArrow, "startBinding")
	customDestination := boundElementRectPageLinkInsetTest(t, customScene.Elements, customArrow, "endBinding")
	if baselineSource != customSource || baselineDestination != customDestination {
		t.Fatalf("same-frame bound boxes changed with metadata row-gap: baseline=%#v/%#v custom=%#v/%#v", baselineSource, baselineDestination, customSource, customDestination)
	}
	if math.Abs(customStart[0]-(customSource[0]+customSource[2]+5)) > 1e-9 || math.Abs(customEnd[0]-(customDestination[0]-5)) > 1e-9 {
		t.Fatalf("same-frame endpoints lost their existing 5px binding gap: source=%#v start=%#v destination=%#v end=%#v", customSource, customStart, customDestination, customEnd)
	}
}

func TestCrossFrameZeroInsetCoincidentFrameEndpointIsRejected(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="300" height="220" margin-bottom="60">
    <metadata position="bottom" row-gap="0" width="100" key-width="30" />
    <connection src="source" dst="target.item"
                src-anchor="top-3" src-frame-anchor="top-3"
                dst-frame-anchor="left-3" />
  </frame>
  <frame id="target" width="300" height="220" margin="60">
    <rectangle id="item" title="Target" width="80" height="40" />
  </frame>
</frames></xaligo>`)
	const want = "coincides with the frame endpoint at zero page inset"
	tests := []struct {
		name string
		run  func() error
	}{
		{
			name: "validate",
			run: func() error {
				return usecase.Validate(context.Background(), source)
			},
		},
		{
			name: "render",
			run: func() error {
				_, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{Theme: "light"})
				return err
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.run()
			if err == nil || !strings.Contains(err.Error(), want) {
				t.Fatalf("%s error = %v, want source-positioned zero-inset coincidence diagnostic %q", test.name, err, want)
			}
		})
	}
}

func TestCrossFrameDefaultInsetCoincidenceShiftsAutomaticTerminalTangent(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="300" height="220" margin-top="4">
    <rectangle id="node" title="Source" width="300" height="40" />
    <connection src="node" dst="target.node"
                src-anchor="top-3" src-frame-side="top" />
  </frame>
  <frame id="target" width="300" height="220" margin="60">
    <rectangle id="node" title="Target" width="80" height="40" />
  </frame>
</frames></xaligo>`)
	scene := renderPageLinkInsetScene(t, source)
	sourceStub, _ := crossFrameStubsPageLinkInsetTest(t, scene.Elements)
	frame := sceneElementRect(t, scene.Elements, "paper-frame-source")
	endpointRect := boundElementRectPageLinkInsetTest(t, scene.Elements, sourceStub, "startBinding")
	start, terminal := sceneArrowEndpoints(t, sourceStub)
	wantStart := [2]float64{endpointRect[0] + endpointRect[2]*0.5, endpointRect[1]}
	if !pointsEqualPageLinkInsetTest(start, wantStart) {
		t.Fatalf("explicit endpoint anchor moved: start=%#v want=%#v endpoint=%#v", start, wantStart, endpointRect)
	}
	if side := sceneFrameSideAtPageLinkInset(frame, terminal, 4); side != "top" {
		t.Fatalf("automatically shifted terminal %#v is not on the default top inset of frame %#v", terminal, frame)
	}
	tangentShift := math.Abs(terminal[0] - start[0])
	if tangentShift <= 1e-9 || tangentShift > 24+1e-9 || math.Abs(terminal[1]-start[1]) > 1e-9 {
		t.Fatalf("automatic coincident terminal did not shift only along the top-edge tangent by up to 24px: start=%#v terminal=%#v", start, terminal)
	}
	points := sceneArrowPoints(t, sourceStub)
	totalLength := 0.0
	for index := 1; index < len(points); index++ {
		dx := math.Abs(points[index][0] - points[index-1][0])
		dy := math.Abs(points[index][1] - points[index-1][1])
		totalLength += dx + dy
		if dx > 1e-9 && dy > 1e-9 {
			t.Fatalf("automatic coincidence stub has diagonal segment %#v -> %#v: %#v", points[index-1], points[index], points)
		}
		if index >= 2 && pointsEqualPageLinkInsetTest(points[index], points[index-2]) {
			t.Fatalf("automatic coincidence stub immediately backtracks %#v -> %#v -> %#v: %#v", points[index-2], points[index-1], points[index], points)
		}
	}
	if totalLength <= 1e-9 {
		t.Fatalf("automatic coincidence produced a zero-length stub: %#v", sourceStub)
	}
	assertCrossFrameEndpointAndTerminalApproaches(t, sourceStub, "top", "top", false)
	assertPageLinkLabelGap(t, sceneTextRectByValue(t, scene.Elements, "to <target>"), terminal, "top", 4)
}

func TestCrossFrameMetadataInsetRejectsTerminalInsideReservedStrip(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="160" height="50">
    <metadata position="bottom" row-gap="12" width="100" key-width="30" />
    <connection src="source" dst="target" src-frame-side="top" />
  </frame>
  <frame id="target" width="160" height="100" />
</frames></xaligo>`)
	assertValidateAndRenderPageLinkInsetErrorTest(t, source, 4, "page terminal", "metadata reservation")
}

func TestCrossFrameDefaultInsetRejectsTerminalOutsideNarrowFrame(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="3" height="100">
    <connection src="source" dst="target" src-frame-side="left" />
  </frame>
  <frame id="target" width="100" height="100" />
</frames></xaligo>`)
	assertValidateAndRenderPageLinkInsetErrorTest(t, source, 3, "page terminal", "outside frame")
}

func TestCrossFrameDefaultInsetAllowsUnusedNarrowFrameDimension(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="3" height="100">
    <connection src="source" dst="target" src-frame-side="top" />
  </frame>
  <frame id="target" width="100" height="100" />
</frames></xaligo>`)
	if err := usecase.Validate(context.Background(), source); err != nil {
		t.Fatalf("Validate() rejected a top terminal whose 4px normal inset fits the frame height: %v", err)
	}
	scene := renderPageLinkInsetScene(t, source)
	sourceStub, _ := crossFrameStubsPageLinkInsetTest(t, scene.Elements)
	frame := sceneElementRect(t, scene.Elements, "paper-frame-source")
	_, terminal := sceneArrowEndpoints(t, sourceStub)
	if side := sceneFrameSideAtPageLinkInset(frame, terminal, 4); side != "top" {
		t.Fatalf("top page terminal = %#v, want 4px top inset in narrow frame %#v", terminal, frame)
	}
}

func TestCrossFrameDefaultInsetAutomaticSideRemapsFromNarrowDimension(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="3" height="100">
    <connection src="source" dst="target" />
  </frame>
  <frame id="target" width="100" height="100" />
</frames></xaligo>`)
	if err := usecase.Validate(context.Background(), source); err != nil {
		t.Fatalf("Validate() rejected an automatic terminal with safe top/bottom alternatives: %v", err)
	}
	scene := renderPageLinkInsetScene(t, source)
	sourceStub, _ := crossFrameStubsPageLinkInsetTest(t, scene.Elements)
	frame := sceneElementRect(t, scene.Elements, "paper-frame-source")
	_, terminal := sceneArrowEndpoints(t, sourceStub)
	side := sceneFrameSideAtPageLinkInset(frame, terminal, 4)
	if side != "top" && side != "bottom" {
		t.Fatalf("automatic page terminal = %#v on %q, want a safe top/bottom side for narrow frame %#v", terminal, side, frame)
	}
}

func TestCrossFrameDefaultInsetRejectsAutomaticFrameWithoutSafeSide(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="3" height="3">
    <connection src="source" dst="target" />
  </frame>
  <frame id="target" width="100" height="100" />
</frames></xaligo>`)
	assertValidateAndRenderPageLinkInsetErrorTest(t, source, 3, "page terminal", "no safe terminal side")
}

func TestCrossFrameMetadataInsetAllowsSafeSelectedSide(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="160" height="50">
    <metadata position="bottom" row-gap="12" width="100" key-width="30" />
    <connection src="source" dst="target" src-frame-anchor="left-1" />
  </frame>
  <frame id="target" width="160" height="100" />
</frames></xaligo>`)
	if err := usecase.Validate(context.Background(), source); err != nil {
		t.Fatalf("Validate() rejected a safe left terminal because the unused top inset line enters metadata: %v", err)
	}
	scene := renderPageLinkInsetScene(t, source)
	sourceStub, _ := crossFrameStubsPageLinkInsetTest(t, scene.Elements)
	frame := sceneElementRect(t, scene.Elements, "paper-frame-source")
	_, terminal := sceneArrowEndpoints(t, sourceStub)
	assertPageLinkTerminalInsetTest(t, frame, terminal, "left", 12, 0.1)
}

func TestCrossFrameMetadataInsetAutomaticSideRemapsFromReservedOppositeSide(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames layout="vertical" gap="48">
  <frame id="source" width="160" height="50">
    <metadata position="bottom" row-gap="12" width="100" key-width="30" />
    <connection src="source" dst="target" />
  </frame>
  <frame id="target" width="160" height="100" />
</frames></xaligo>`)
	if err := usecase.Validate(context.Background(), source); err != nil {
		t.Fatalf("Validate() rejected an automatic terminal with safe left/right alternatives: %v", err)
	}
	scene := renderPageLinkInsetScene(t, source)
	sourceStub, _ := crossFrameStubsPageLinkInsetTest(t, scene.Elements)
	frame := sceneElementRect(t, scene.Elements, "paper-frame-source")
	reserved := metadataReservedRectPageLinkInsetTest(t, scene.Elements, "source")
	_, terminal := sceneArrowEndpoints(t, sourceStub)
	side := sceneFrameSideAtPageLinkInset(frame, terminal, 12)
	if side != "left" && side != "right" {
		t.Fatalf("automatic page terminal = %#v on %q, want a safe left/right side outside metadata %#v", terminal, side, reserved)
	}
	if pointInsideRectPageLinkInsetTest(terminal, reserved) {
		t.Fatalf("automatic page terminal %#v entered metadata reservation %#v", terminal, reserved)
	}
}

func TestCrossFrameZeroInsetOwningFrameSideEquivalentFrameAnchorIsRejected(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="160" height="100">
    <metadata position="bottom" row-gap="0" width="100" key-width="30" />
    <connection src="source" dst="target"
                src-side="top" src-frame-anchor="top-3" />
  </frame>
  <frame id="target" width="160" height="100" />
</frames></xaligo>`)
	assertValidateAndRenderPageLinkInsetErrorTest(t, source, 4, "coincides with the frame endpoint at zero page inset")
}

func TestCrossFrameZeroInsetOwningFrameAutomaticSideEquivalentFrameAnchorIsRejected(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="160" height="100">
    <metadata position="bottom" row-gap="0" width="100" key-width="30" />
    <connection src="source" dst="target" src-frame-anchor="right-3" />
  </frame>
  <frame id="target" width="160" height="100" />
</frames></xaligo>`)
	assertValidateAndRenderPageLinkInsetErrorTest(t, source, 4, "coincides with the frame endpoint at zero page inset")
}

func TestCrossFrameNarrowMetadataSafeRangeKeepsAutomaticCoincidenceInsideFrame(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" width="120" height="56" margin-left="4">
    <metadata position="top" row-gap="4" width="80" key-width="24" />
    <rectangle id="node" title="Source" width="40" height="20" />
    <connection src="node" dst="target.node"
                src-anchor="left-3" src-frame-side="left" />
  </frame>
  <frame id="target" width="120" height="80" margin="16">
    <rectangle id="node" title="Target" width="40" height="20" />
  </frame>
</frames></xaligo>`)
	if err := usecase.Validate(context.Background(), source); err != nil {
		assertValidateAndRenderPageLinkInsetErrorTest(t, source, 5, "page terminal")
		return
	}

	scene := renderPageLinkInsetScene(t, source)
	sourceStub, _ := crossFrameStubsPageLinkInsetTest(t, scene.Elements)
	frame := sceneElementRect(t, scene.Elements, "paper-frame-source")
	reserved := metadataReservedRectPageLinkInsetTest(t, scene.Elements, "source")
	start, terminal := sceneArrowEndpoints(t, sourceStub)
	if math.Abs(start[0]-terminal[0]) > 1e-9 || math.Abs(start[1]-terminal[1]) <= 1e-9 {
		t.Fatalf("narrow-range coincidence avoidance did not shift only along the left-edge tangent: start=%#v terminal=%#v", start, terminal)
	}
	if terminal[1] <= reserved[1]+reserved[3]+1e-9 || terminal[1] >= frame[1]+frame[3]-1e-9 {
		t.Fatalf("automatic coincidence terminal %#v is not in the non-reserved frame interval below %#v within %#v", terminal, reserved, frame)
	}
	for _, point := range sceneArrowPoints(t, sourceStub) {
		if point[0] < frame[0]-1e-9 || point[0] > frame[0]+frame[2]+1e-9 || point[1] < frame[1]-1e-9 || point[1] > frame[1]+frame[3]+1e-9 {
			t.Fatalf("automatic coincidence avoidance moved page-link point %#v outside frame %#v: %#v", point, frame, sourceStub)
		}
		if pointInsideRectPageLinkInsetTest(point, reserved) {
			t.Fatalf("automatic coincidence avoidance moved page-link point %#v inside metadata reservation %#v: %#v", point, reserved, sourceStub)
		}
	}
}

func renderPageLinkInsetScene(t *testing.T, source []byte) sceneFile {
	t.Helper()
	out, err := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	return scene
}

func crossFrameStubsPageLinkInsetTest(t *testing.T, elements []map[string]any) (map[string]any, map[string]any) {
	t.Helper()
	var sourceStub, destinationStub map[string]any
	for _, element := range elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] != true {
			continue
		}
		if element["startBinding"] != nil && element["endBinding"] == nil {
			sourceStub = element
		}
		if element["startBinding"] == nil && element["endBinding"] != nil {
			destinationStub = element
		}
	}
	if sourceStub == nil || destinationStub == nil {
		t.Fatalf("cross-frame stubs missing: source=%#v destination=%#v elements=%#v", sourceStub, destinationStub, elements)
	}
	return sourceStub, destinationStub
}

func assertPageLinkTerminalInsetTest(t *testing.T, frame [4]float64, terminal [2]float64, side string, inset, anchorFraction float64) {
	t.Helper()
	if got := sceneFrameSideAtPageLinkInset(frame, terminal, inset); got != side {
		t.Fatalf("page-link terminal %#v side = %q, want %q at inset %v from frame %#v", terminal, got, side, inset, frame)
	}
	var gotParallel, wantParallel float64
	if side == "top" || side == "bottom" {
		gotParallel = terminal[0]
		wantParallel = frame[0] + frame[2]*anchorFraction
	} else {
		gotParallel = terminal[1]
		wantParallel = frame[1] + frame[3]*anchorFraction
	}
	if math.Abs(gotParallel-wantParallel) > 1e-9 {
		t.Fatalf("page-link terminal %#v moved explicit %s anchor parallel coordinate: got %v, want %v", terminal, side, gotParallel, wantParallel)
	}
}

func assertPageLinkFrameApproachInsetTest(t *testing.T, stub map[string]any, side string, terminalAtStart bool) {
	t.Helper()
	points := sceneArrowPoints(t, stub)
	if terminalAtStart {
		assertSegmentPerpendicularToSide(t, points[0], points[1], side)
		return
	}
	assertSegmentPerpendicularToSide(t, points[len(points)-2], points[len(points)-1], side)
}

func renderSameFrameConnectionPageLinkInsetTest(t *testing.T, metadata string) (sceneFile, map[string]any) {
	t.Helper()
	source := []byte(`<xaligo version="1"><frames>
  <frame id="page" width="400" height="240" layout="horizontal" gap="40"
         margin-top="80" margin-right="20" margin-bottom="20" margin-left="20">
    ` + metadata + `
    <rectangle id="source" title="Source" width="100" height="60" />
    <rectangle id="destination" title="Destination" width="100" height="60" />
    <connection src="source" dst="destination" src-side="right" dst-side="left" />
  </frame>
</frames></xaligo>`)
	scene := renderPageLinkInsetScene(t, source)
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if element["type"] == "arrow" && custom["xaligoCrossFrame"] != true && element["startBinding"] != nil && element["endBinding"] != nil {
			return scene, element
		}
	}
	t.Fatalf("same-frame connection arrow missing: %#v", scene.Elements)
	return sceneFile{}, nil
}

func boundElementRectPageLinkInsetTest(t *testing.T, elements []map[string]any, arrow map[string]any, bindingName string) [4]float64 {
	t.Helper()
	binding, _ := arrow[bindingName].(map[string]any)
	elementID, _ := binding["elementId"].(string)
	return sceneElementRect(t, elements, elementID)
}

func pointsEqualPageLinkInsetTest(left, right [2]float64) bool {
	return math.Abs(left[0]-right[0]) <= 1e-9 && math.Abs(left[1]-right[1]) <= 1e-9
}

func assertValidateAndRenderPageLinkInsetErrorTest(t *testing.T, source []byte, wantLine int, needles ...string) {
	t.Helper()
	validationErr := usecase.Validate(context.Background(), source)
	if validationErr == nil {
		t.Fatalf("Validate() error = nil, want positioned page-link diagnostic containing %q", needles)
	}
	var diagnosticsErr *entity.DiagnosticsError
	if !errors.As(validationErr, &diagnosticsErr) || len(diagnosticsErr.Diagnostics) == 0 {
		t.Fatalf("Validate() error = %T %v, want DiagnosticsError", validationErr, validationErr)
	}
	diagnostic := diagnosticsErr.Diagnostics[0]
	if diagnostic.Line != wantLine || diagnostic.Column == 0 {
		t.Fatalf("Validate() diagnostic position = line %d, column %d, want line %d with a non-zero column: %#v", diagnostic.Line, diagnostic.Column, wantLine, diagnostic)
	}
	for _, needle := range needles {
		if !strings.Contains(validationErr.Error(), needle) {
			t.Fatalf("Validate() error = %v, want %q", validationErr, needle)
		}
	}

	_, renderErr := newUsecase().RenderExcalidraw(context.Background(), source, entity.RenderOptions{Theme: "light"})
	if renderErr == nil {
		t.Fatalf("RenderExcalidraw() error = nil, want positioned page-link diagnostic containing %q", needles)
	}
	var positionedErr *entity.ParseError
	if !errors.As(renderErr, &positionedErr) || positionedErr.Position.Line != wantLine || positionedErr.Position.Column == 0 {
		t.Fatalf("RenderExcalidraw() error = %T %v, want ParseError on line %d with a non-zero column", renderErr, renderErr, wantLine)
	}
	for _, needle := range needles {
		if !strings.Contains(renderErr.Error(), needle) {
			t.Fatalf("RenderExcalidraw() error = %v, want %q", renderErr, needle)
		}
	}
}

func metadataReservedRectPageLinkInsetTest(t *testing.T, elements []map[string]any, frameID string) [4]float64 {
	t.Helper()
	for _, element := range elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoFrameMetadataReserved"] != true || custom["xaligoFrameID"] != frameID {
			continue
		}
		return [4]float64{
			sceneNumber(t, element["x"]), sceneNumber(t, element["y"]),
			sceneNumber(t, element["width"]), sceneNumber(t, element["height"]),
		}
	}
	t.Fatalf("frame %q metadata reservation missing: %#v", frameID, elements)
	return [4]float64{}
}

func pointInsideRectPageLinkInsetTest(point [2]float64, rect [4]float64) bool {
	const epsilon = 1e-9
	return point[0] > rect[0]+epsilon && point[0] < rect[0]+rect[2]-epsilon &&
		point[1] > rect[1]+epsilon && point[1] < rect[1]+rect[3]-epsilon
}
