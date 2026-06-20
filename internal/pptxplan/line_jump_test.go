package pptxplan

import "testing"

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
	op := lineJumpMaskOp(pt{X: 96, Y: 96}, rect{}, 96, "ABCDEF")
	if op.Kind != "rect" || op.Fill == nil || op.Fill.Color != "ABCDEF" || op.W != 6.0/96.0 || op.H != 6.0/96.0 {
		t.Fatalf("lineJumpMaskOp() = %#v", op)
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
