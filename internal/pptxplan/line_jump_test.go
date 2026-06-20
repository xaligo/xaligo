package pptxplan

import (
	"strings"
	"testing"
)

func TestPathCrossingsFindsInteriorOrthogonalIntersections(t *testing.T) {
	below := []routedPath{{ID: "route", Points: []pt{{X: 0, Y: 50}, {X: 100, Y: 50}}}}
	top := routedPath{ID: "traffic", Points: []pt{{X: 40, Y: 0}, {X: 40, Y: 100}}}

	got := pathCrossings(top, below)
	if len(got) != 1 || got[0] != (pt{X: 40, Y: 50}) {
		t.Fatalf("pathCrossings() = %#v", got)
	}
}

func TestPathCrossingsIgnoresEndpointTouchesAndOverlaps(t *testing.T) {
	below := []routedPath{
		{ID: "touch", Points: []pt{{X: 0, Y: 50}, {X: 40, Y: 50}}},
		{ID: "overlap", Points: []pt{{X: 40, Y: 0}, {X: 40, Y: 100}}},
	}
	top := routedPath{ID: "top", Points: []pt{{X: 40, Y: 50}, {X: 80, Y: 50}}}

	if got := pathCrossings(top, below); len(got) != 0 {
		t.Fatalf("pathCrossings() = %#v, want none", got)
	}
}

func TestLineJumpMaskUsesSlideBackground(t *testing.T) {
	op := lineJumpMaskOp("jump", pt{X: 96, Y: 96}, rect{}, 96, "ABCDEF")
	if op.Kind != "rect" || !op.FrontLayer || op.Fill == nil || op.Fill.Color != "ABCDEF" || op.W != 6.0/96.0 || op.H != 6.0/96.0 {
		t.Fatalf("lineJumpMaskOp() = %#v", op)
	}
}

func TestBuildPlanAnchorGridUsesSlideBackgroundAndCoversLabel(t *testing.T) {
	opacity := 100.0
	files := map[string]SceneFile{
		"src-file": {DataURL: "data:image/svg+xml;base64,PHN2Zy8+"},
		"dst-file": {DataURL: "data:image/svg+xml;base64,PHN2Zy8+"},
	}
	scene := Scene{Elements: []Element{
		{ID: "src-icon", Type: "image", X: 0, Y: 0, Width: 32, Height: 32, BackgroundColor: "#F58536", Opacity: &opacity, FileID: "src-file"},
		{ID: "src-icon-lbl", Type: "text", X: -12, Y: 36, Width: 56, Height: 14, BackgroundColor: "transparent", Opacity: &opacity, Text: "SRC"},
		{ID: "dst-icon", Type: "image", X: 96, Y: 0, Width: 32, Height: 32, BackgroundColor: "#3366CC", Opacity: &opacity, FileID: "dst-file"},
		{ID: "dst-icon-lbl", Type: "text", X: 84, Y: 36, Width: 56, Height: 14, BackgroundColor: "transparent", Opacity: &opacity, Text: "DST"},
		{ID: "connector", Type: "arrow", StrokeColor: "#111111", Opacity: &opacity, StartBinding: &Binding{ElementID: "src-icon", FixedPoint: []float64{1, 0.5}}, EndBinding: &Binding{ElementID: "dst-icon", FixedPoint: []float64{0, 0.5}}},
	}, Files: files, AppState: &AppState{ViewBackgroundColor: "#FFFFFF"}}
	plan := BuildPlan(&scene, Options{PxPerInch: 96})
	seenSlideBackground := false
	tallGrid := false
	paddedGrid := false
	seenGroupedGrid := false
	seenGroupedIcon := false
	seenGroupedLabel := false
	wantCellW := (56.0 + anchorGridVisualPadPx*2) / anchorGrid / 96.0
	for _, op := range plan.Ops {
		if op.GroupID == anchorGroupID("src-icon") && strings.HasPrefix(op.ID, "src-icon-grid-") {
			seenGroupedGrid = true
		}
		if op.ID == "src-icon" && op.GroupID == anchorGroupID("src-icon") {
			seenGroupedIcon = true
		}
		if op.ID == "src-icon-lbl" && op.GroupID == anchorGroupID("src-icon") {
			seenGroupedLabel = true
		}
		if op.Kind != "rect" || op.Fill == nil {
			continue
		}
		if op.Fill.Color == "FFFFFF" && op.W < 0.3 && op.H < 0.3 {
			seenSlideBackground = true
			if op.H > 0.09 {
				tallGrid = true
			}
			if op.W >= wantCellW {
				paddedGrid = true
			}
		}
	}
	if !seenSlideBackground || !tallGrid || !paddedGrid {
		t.Fatalf("anchor grid slide background=%v tall=%v padded=%v ops=%#v", seenSlideBackground, tallGrid, paddedGrid, plan.Ops)
	}
	if !seenGroupedGrid || !seenGroupedIcon || !seenGroupedLabel {
		t.Fatalf("anchor group missing grid=%v icon=%v label=%v ops=%#v", seenGroupedGrid, seenGroupedIcon, seenGroupedLabel, plan.Ops)
	}
}

func TestLineJumpBackgroundUsesTopmostOpaqueFill(t *testing.T) {
	opacity := 100.0
	elements := []*Element{
		{Type: "rectangle", X: 0, Y: 0, Width: 100, Height: 100, BackgroundColor: "#112233", Opacity: &opacity},
		{Type: "rectangle", X: 20, Y: 20, Width: 60, Height: 60, BackgroundColor: "#AABBCC", Opacity: &opacity},
	}
	if got := lineJumpBackground(pt{X: 50, Y: 50}, elements, "FFFFFF"); got != "AABBCC" {
		t.Fatalf("lineJumpBackground() = %q", got)
	}
}

func TestLineJumpBackgroundIgnoresTransparentFill(t *testing.T) {
	opacity := 50.0
	elements := []*Element{{Type: "rectangle", X: 0, Y: 0, Width: 100, Height: 100, BackgroundColor: "#112233", Opacity: &opacity}}
	if got := lineJumpBackground(pt{X: 50, Y: 50}, elements, "FFFFFF"); got != "FFFFFF" {
		t.Fatalf("lineJumpBackground() = %q", got)
	}
}

func TestPathBorderCrossingsFindsGroupFrameIntersection(t *testing.T) {
	path := routedPath{ID: "traffic", Points: []pt{{X: 0, Y: 50}, {X: 100, Y: 50}}}
	borders := []segment{{A: pt{X: 40, Y: 0}, B: pt{X: 40, Y: 100}}}
	got := pathBorderCrossings(path, borders)
	if len(got) != 1 || got[0] != (pt{X: 40, Y: 50}) {
		t.Fatalf("pathBorderCrossings() = %#v", got)
	}
}

func TestGroupBorderMaskIsOpaqueWhite(t *testing.T) {
	op := groupBorderMaskOp("border", pt{X: 96, Y: 96}, rect{}, 96)
	if op.Kind != "rect" || !op.FrontLayer || op.Fill == nil || op.Fill.Color != "FFFFFF" || op.Fill.Transparency != 0 || op.W != 8.0/96.0 {
		t.Fatalf("groupBorderMaskOp() = %#v", op)
	}
}

func TestBuildPlanDrawsDeferredHeaderAfterGroupBorders(t *testing.T) {
	opacity := 100.0
	scene := Scene{Elements: []Element{
		{ID: "paper-frame", Type: "frame", Width: 200, Height: 100},
		{ID: "parent-header-bg", Type: "line", X: 0, Y: 0, Width: 80, Height: 40, StrokeColor: "#AAB7B8", BackgroundColor: "#FFFFFF", Opacity: &opacity, Points: [][]float64{{0, 0}, {66, 0}, {80, 20}, {66, 40}, {0, 40}, {0, 0}}, CustomData: &CustomData{GroupHeader: true}},
		{ID: "child-rect", Type: "rectangle", X: 20, Y: 0, Width: 100, Height: 80, StrokeColor: "#00A1C9", BackgroundColor: "transparent", Opacity: &opacity, CustomData: &CustomData{GroupBorder: true}},
	}}
	plan := BuildPlan(&scene, Options{PxPerInch: 96})
	if len(plan.Ops) < 2 || plan.Ops[0].Kind != "rect" || plan.Ops[1].Kind != "polygon" {
		t.Fatalf("draw order = %#v", plan.Ops)
	}
}

func TestBuildPlanDrawsHeaderTagAboveConnectors(t *testing.T) {
	opacity := 100.0
	scene := Scene{Elements: []Element{
		{ID: "paper-frame", Type: "frame", Width: 200, Height: 100},
		{ID: "group-rect", Type: "rectangle", X: 0, Y: 0, Width: 160, Height: 80, StrokeColor: "#00A1C9", BackgroundColor: "transparent", Opacity: &opacity, CustomData: &CustomData{GroupBorder: true}},
		{ID: "raw-line", Type: "line", X: 0, Y: 20, Width: 160, Height: 0, StrokeColor: "#334155", Opacity: &opacity, Points: [][]float64{{0, 0}, {160, 0}}},
		{ID: "header-bg", Type: "line", X: 0, Y: 8, Width: 80, Height: 24, StrokeColor: "#00A1C9", BackgroundColor: "#FFFFFF", Opacity: &opacity, Points: [][]float64{{0, 0}, {68, 0}, {80, 12}, {68, 24}, {0, 24}, {0, 0}}, CustomData: &CustomData{GroupHeader: true}},
	}}
	plan := BuildPlan(&scene, Options{PxPerInch: 96})
	lineIndex, headerIndex := -1, -1
	for i, op := range plan.Ops {
		if op.Kind == "line" {
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

func TestCollectObstaclesIncludesGroupHeaderTag(t *testing.T) {
	elements := []*Element{
		{ID: "header-bg", Type: "line", X: 10, Y: 20, Width: 80, Height: 24, CustomData: &CustomData{GroupHeader: true}},
		{ID: "connector", Type: "line", X: 0, Y: 0, Width: 100, Height: 0},
	}
	got := collectObstacles(elements)
	if len(got) != 1 || got[0] != (rect{X: 10, Y: 20, W: 80, H: 24}) {
		t.Fatalf("collectObstacles() = %#v", got)
	}
}

func TestSeparateExactOverlapsOffsetsInternalTrunk(t *testing.T) {
	points := []pt{{X: 0, Y: 0}, {X: 20, Y: 0}, {X: 20, Y: 80}, {X: 100, Y: 80}, {X: 100, Y: 100}}
	placed := [][]segment{{{A: pt{X: 20, Y: 40}, B: pt{X: 20, Y: 70}}}}
	got := separateExactOverlaps(points, placed, nil, defaultRouterOptions())
	if overlap := exactOverlapLength(toSegments(got), placed); overlap > eps {
		t.Fatalf("overlap remains %.2f in %#v", overlap, got)
	}
	if got[0] != points[0] || got[len(got)-1] != points[len(points)-1] {
		t.Fatalf("endpoints changed: %#v", got)
	}
}

func TestSeparateExactOverlapsEscapesOverlappingFirstStub(t *testing.T) {
	points := []pt{{X: 100, Y: 20}, {X: 80, Y: 20}, {X: 80, Y: 100}}
	placed := [][]segment{{{A: pt{X: 60, Y: 20}, B: pt{X: 100, Y: 20}}}}
	got := separateExactOverlaps(points, placed, nil, defaultRouterOptions())
	if overlap := exactOverlapLength(toSegments(got), placed); overlap > 2.01 {
		t.Fatalf("first-stub overlap remains %.2f in %#v", overlap, got)
	}
	if got[0] != points[0] || got[len(got)-1] != points[len(points)-1] {
		t.Fatalf("endpoints changed: %#v", got)
	}
}

// TestSeparateExactOverlapsEscapesStubInCrowdedCorner reproduces the WAF→IGW
// stub that shared the Admin→IAM approach lane at the top of the complex sample.
// Both jog directions intrude an icon's full clearance halo, so the escape only
// succeeds because the stub jog uses a lighter margin than the routing clearance.
func TestSeparateExactOverlapsEscapesStubInCrowdedCorner(t *testing.T) {
	points := []pt{{X: 100, Y: 20}, {X: 80, Y: 20}, {X: 80, Y: 100}}
	placed := [][]segment{{{A: pt{X: 60, Y: 20}, B: pt{X: 100, Y: 20}}}}
	// One halo above (blocks the up jog) and one below (blocks the down jog) at
	// full 12 px clearance, but neither rect reaches the lighter 4 px stub margin.
	obstacles := []rect{
		{X: 85, Y: -6, W: 13, H: 10},
		{X: 85, Y: 36, W: 13, H: 10},
	}
	got := separateExactOverlaps(points, placed, obstacles, defaultRouterOptions())
	if overlap := exactOverlapLength(toSegments(got), placed); overlap > 2.01 {
		t.Fatalf("stub overlap remains %.2f in %#v", overlap, got)
	}
	if got[0] != points[0] || got[len(got)-1] != points[len(points)-1] {
		t.Fatalf("endpoints changed: %#v", got)
	}
}

func TestSeparateObstacleHitsOffsetsInternalSegment(t *testing.T) {
	points := []pt{{X: 0, Y: -20}, {X: 0, Y: 0}, {X: 80, Y: 0}, {X: 80, Y: 40}}
	obstacles := []rect{{X: 20, Y: -5, W: 20, H: 10}}
	got := separateObstacleHits(points, nil, obstacles, defaultRouterOptions())
	if hits := obstacleHitCount(got, obstacles); hits > eps {
		t.Fatalf("obstacle hit remains %.2f in %#v", hits, got)
	}
	if got[0] != points[0] || got[len(got)-1] != points[len(points)-1] {
		t.Fatalf("endpoints changed: %#v", got)
	}
}

func TestSeparateObstacleHitsOffsetsVisualMarginContact(t *testing.T) {
	points := []pt{{X: 0, Y: -20}, {X: 0, Y: 0}, {X: 80, Y: 0}, {X: 80, Y: 40}}
	obstacles := []rect{{X: 20, Y: 2, W: 20, H: 10}}
	visualObstacles := inflateRects(obstacles, 4)
	got := separateObstacleHits(points, nil, visualObstacles, defaultRouterOptions())
	if hits := obstacleHitCount(got, visualObstacles); hits > eps {
		t.Fatalf("visual-margin hit remains %.2f in %#v", hits, got)
	}
	if got[0] != points[0] || got[len(got)-1] != points[len(points)-1] {
		t.Fatalf("endpoints changed: %#v", got)
	}
}
