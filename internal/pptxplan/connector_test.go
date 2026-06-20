package pptxplan

import (
	"math"
	"testing"
)

func TestConnectorLinePreservesRouteAndTrafficStyles(t *testing.T) {
	global := resolveConnectorStyle("thin")
	route := connectorLine(&Element{
		Type: "arrow", StrokeColor: "#64748b", StrokeWidth: 1,
		CustomData: &CustomData{ConnectorKind: "route"},
	}, global)
	if route.BeginArrowType != "none" || route.EndArrowType != "stealth" || route.Width != 1 || route.Color != "64748B" {
		t.Fatalf("route line = %#v", route)
	}

	traffic := connectorLine(&Element{
		Type: "arrow", StrokeColor: "#2563eb", StrokeWidth: 2, StrokeStyle: "dotted",
		CustomData: &CustomData{ConnectorKind: "traffic", ConnectorStartArrowhead: "oval", ConnectorEndArrowhead: "diamond"},
	}, global)
	if traffic.BeginArrowType != "oval" || traffic.EndArrowType != "diamond" || traffic.Width != 1 || traffic.Color != "2563EB" || traffic.Dash != "dot" {
		t.Fatalf("traffic line = %#v", traffic)
	}
}

func TestLinePropsKeepsTransparentStrokeInvisible(t *testing.T) {
	line := lineProps(&Element{StrokeColor: "transparent", StrokeWidth: 1})
	if line.Transparency != 100 {
		t.Fatalf("transparent line = %#v", line)
	}
}

func TestConnectorLegendEntryIncludesLineIDAndEndpoints(t *testing.T) {
	line := connectorLine(&Element{Type: "arrow", StrokeColor: "#2563eb", StrokeWidth: 3, CustomData: &CustomData{ConnectorKind: "traffic"}}, resolveConnectorStyle("thin"))
	entry := connectorLegendEntry("L07", &Element{
		Type: "arrow", StartBinding: &Binding{ElementID: "src"}, EndBinding: &Binding{ElementID: "dst"},
		CustomData: &CustomData{ConnectorKind: "traffic"},
	}, line)
	if entry.ID != "L07" || entry.Kind != "traffic" || entry.Source != "src" || entry.Target != "dst" || entry.Line.Color != "2563EB" {
		t.Fatalf("connector legend entry = %#v", entry)
	}
}

func TestConnectorIDLabelUsesEndpointOrBendCandidate(t *testing.T) {
	path := routedPath{ID: "conn", Points: []pt{{X: 0, Y: 0}, {X: 10, Y: 0}, {X: 10, Y: 100}}}
	op, _, ok := connectorIDLabelOp("L01", path, []routedPath{path}, nil, nil, rect{}, 100, LineStyle{Color: "2563EB"})
	if !ok || op.Kind != "text" || op.Text != "L01" || op.Color != "2563EB" {
		t.Fatalf("connector id label = %#v, %v", op, ok)
	}
	if op.X > 0.3 || op.Y > 0.3 {
		t.Fatalf("connector id label was not placed near an endpoint or bend: %#v", op)
	}
}

func TestConnectorIDLabelAvoidsCrowdedCandidate(t *testing.T) {
	path := routedPath{ID: "conn", Points: []pt{{X: 0, Y: 0}, {X: 80, Y: 0}}}
	obstacles := []rect{{X: 6, Y: -20, W: 44, H: 40}}
	op, labelRect, ok := connectorIDLabelOp("L01", path, []routedPath{path}, obstacles, nil, rect{}, 100, LineStyle{Color: "2563EB"})
	if !ok {
		t.Fatal("connector id label was not generated")
	}
	if rectsOverlap(labelRect, obstacles[0]) {
		t.Fatalf("connector id label did not move away from crowded start: %#v", op)
	}
}

func TestRouteOneFallbackRemainsOrthogonal(t *testing.T) {
	req := routeRequest{ID: "conn", Src: rect{X: 0, Y: 0, W: 20, H: 20}, Dst: rect{X: 80, Y: 50, W: 20, H: 20}, SrcSide: sideLeft, DstSide: sideRight}
	path := routeOne(req, []rect{{X: -1000, Y: -1000, W: 2000, H: 2000}}, nil, defaultRouterOptions())
	for _, seg := range toSegments(path.Points) {
		if !isHorizontal(seg) && math.Abs(seg.A.X-seg.B.X) > eps {
			t.Fatalf("fallback produced diagonal segment: %#v", path.Points)
		}
	}
}

func TestOrthogonalizeEndpointStubsRemovesDiagonalFirstAndLastSegments(t *testing.T) {
	req := routeRequest{SrcSide: sideRight, DstSide: sideBottom}
	points := []pt{{X: 0, Y: 10}, {X: 8, Y: 4}, {X: 30, Y: 4}, {X: 40, Y: 20}, {X: 45, Y: 30}}
	got := orthogonalizeEndpointStubs(points, req)
	for _, seg := range []segment{{A: got[0], B: got[1]}, {A: got[len(got)-2], B: got[len(got)-1]}} {
		if !isHorizontal(seg) && math.Abs(seg.A.X-seg.B.X) > eps {
			t.Fatalf("endpoint stub remains diagonal: %#v", got)
		}
	}
}
