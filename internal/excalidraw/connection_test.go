package excalidraw

import (
	"testing"

	"github.com/ryo-arima/xaligo/internal/model"
)

func TestResolveConnectionStyleKinds(t *testing.T) {
	tests := []struct {
		name           string
		attrs          map[string]string
		kind           string
		color          string
		width          float64
		startArrowhead string
		endArrowhead   string
		stroke         string
	}{
		{name: "default", attrs: map[string]string{}, kind: "connection", color: "#1e1e1e", width: 1, startArrowhead: "none", endArrowhead: "arrow", stroke: "solid"},
		{name: "route", attrs: map[string]string{"kind": "route"}, kind: "route", color: "#64748b", width: 1, startArrowhead: "oval", endArrowhead: "oval", stroke: "solid"},
		{name: "route without connectors", attrs: map[string]string{"kind": "route", "start-arrowhead": "none", "end-arrowhead": "none"}, kind: "route", color: "#64748b", width: 1, startArrowhead: "none", endArrowhead: "none", stroke: "solid"},
		{name: "traffic", attrs: map[string]string{"kind": "traffic"}, kind: "traffic", color: "#2563eb", width: 2, startArrowhead: "none", endArrowhead: "arrow", stroke: "solid"},
		{name: "overrides", attrs: map[string]string{"kind": "traffic", "color": "#dc2626", "stroke-width": "3", "stroke-style": "dotted", "start-arrowhead": "oval", "end-arrowhead": "diamond"}, kind: "traffic", color: "#dc2626", width: 3, startArrowhead: "oval", endArrowhead: "diamond", stroke: "dotted"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveConnectionStyle(&model.Node{Attrs: tt.attrs})
			if got.Kind != tt.kind || got.Color != tt.color || got.Width != tt.width || got.StartArrowhead != tt.startArrowhead || got.EndArrowhead != tt.endArrowhead || got.StrokeStyle != tt.stroke {
				t.Fatalf("resolveConnectionStyle() = %#v", got)
			}
		})
	}
}

func TestExtendConnectionPoint(t *testing.T) {
	if got := extendConnectionPoint([2]float64{10, 20}, "right", 25); got != ([2]float64{35, 20}) {
		t.Fatalf("right extension = %#v", got)
	}
	if got := extendConnectionPoint([2]float64{10, 20}, "top", 25); got != ([2]float64{10, -5}) {
		t.Fatalf("top extension = %#v", got)
	}
}
