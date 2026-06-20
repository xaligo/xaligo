package pptxplan

import "testing"

func TestCollectContainerBorderPathsSkipsInvisibleAndPaperFrames(t *testing.T) {
	elements := []*Element{
		{ID: "paper-frame", Type: "frame", Width: 200, Height: 100, StrokeColor: "#bbb"},
		{ID: "visible", Type: "rectangle", X: 10, Y: 20, Width: 100, Height: 50, StrokeColor: "#00A1C9"},
		{ID: "hidden", Type: "rectangle", Width: 20, Height: 20, StrokeColor: "transparent"},
		{ID: "icon", Type: "image", Width: 20, Height: 20, StrokeColor: "#000000"},
	}
	paths := collectContainerBorderPaths(elements)
	if len(paths) != 4 {
		t.Fatalf("border paths = %d, want 4", len(paths))
	}
	if got := paths[0][0]; got.A != (pt{X: 10, Y: 20}) || got.B != (pt{X: 110, Y: 20}) {
		t.Fatalf("top border = %#v", got)
	}
}

func TestFrameBorderReservationPrefersMargin(t *testing.T) {
	border := segment{A: pt{X: 0, Y: 10}, B: pt{X: 200, Y: 10}}
	near := []pt{{X: 20, Y: 12}, {X: 180, Y: 12}}
	clear := []pt{{X: 20, Y: 18}, {X: 180, Y: 18}}
	reserved := [][]segment{{border}}

	nearScore := scorePath(near, nil, reserved, 8)
	clearScore := scorePath(clear, nil, reserved, 8)
	if clearScore >= nearScore {
		t.Fatalf("clear path score %.2f should be lower than near-border score %.2f", clearScore, nearScore)
	}
}
