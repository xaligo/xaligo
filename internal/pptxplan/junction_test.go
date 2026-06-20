package pptxplan

import "testing"

func TestApplyRouteJunctionsSharesFanOutStub(t *testing.T) {
	src := rect{X: 10, Y: 20, W: 20, H: 20}
	requests := []routeRequest{
		{ID: "a", Kind: "route", Src: src, Dst: rect{X: 100, Y: 10, W: 10, H: 10}, SrcSide: sideRight, DstSide: sideLeft, SrcGap: 5},
		{ID: "b", Kind: "route", Src: src, Dst: rect{X: 100, Y: 50, W: 10, H: 10}, SrcSide: sideRight, DstSide: sideLeft, SrcGap: 5},
		{ID: "traffic", Kind: "traffic", Src: src, Dst: rect{X: 100, Y: 80, W: 10, H: 10}, SrcSide: sideRight, DstSide: sideLeft},
	}

	junctions := applyRouteJunctions(requests, 20)
	if len(junctions) != 1 {
		t.Fatalf("junction count = %d, want 1", len(junctions))
	}
	wantAnchor := pt{X: 30, Y: 30}
	if requests[0].SrcAnchor == nil || *requests[0].SrcAnchor != wantAnchor || requests[1].SrcAnchor == nil || *requests[1].SrcAnchor != wantAnchor {
		t.Fatalf("route anchors = %#v, %#v", requests[0].SrcAnchor, requests[1].SrcAnchor)
	}
	if requests[2].SrcAnchor != nil {
		t.Fatalf("traffic anchor unexpectedly changed: %#v", requests[2].SrcAnchor)
	}
	if junctions[0].Point != (pt{X: 55, Y: 30}) {
		t.Fatalf("junction point = %#v", junctions[0].Point)
	}
}
