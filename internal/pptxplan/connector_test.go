package pptxplan

import "testing"

func TestConnectorLinePreservesRouteAndTrafficStyles(t *testing.T) {
	global := resolveConnectorStyle("thin")
	route := connectorLine(&Element{
		Type: "arrow", StrokeColor: "#64748b", StrokeWidth: 1,
		CustomData: &CustomData{ConnectorKind: "route"},
	}, global)
	if route.BeginArrowType != "oval" || route.EndArrowType != "oval" || route.Width != 1 || route.Color != "64748B" {
		t.Fatalf("route line = %#v", route)
	}

	traffic := connectorLine(&Element{
		Type: "arrow", StrokeColor: "#2563eb", StrokeWidth: 2, StrokeStyle: "dotted",
		CustomData: &CustomData{ConnectorKind: "traffic", ConnectorStartArrowhead: "oval", ConnectorEndArrowhead: "diamond"},
	}, global)
	if traffic.BeginArrowType != "oval" || traffic.EndArrowType != "diamond" || traffic.Width != 2 || traffic.Color != "2563EB" || traffic.Dash != "dot" {
		t.Fatalf("traffic line = %#v", traffic)
	}
}
