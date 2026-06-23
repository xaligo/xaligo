package usecase_test

import (
	"encoding/json"
	"strings"
	"testing"

	awsassets "github.com/ryo-arima/xaligo/etc/resources/aws"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

func TestResolveConnectionStyleKinds(t *testing.T) {
	tests := []struct {
		name           string
		attrs          string
		kind           string
		color          string
		width          float64
		startArrowhead string
		endArrowhead   string
		stroke         string
	}{
		{name: "default", kind: "connection", color: "#1e1e1e", width: 1, startArrowhead: "none", endArrowhead: "stealth", stroke: "solid"},
		{name: "route", attrs: `kind="route"`, kind: "route", color: "#64748b", width: 1, startArrowhead: "none", endArrowhead: "stealth", stroke: "solid"},
		{name: "route without connectors", attrs: `kind="route" start-arrowhead="none" end-arrowhead="none"`, kind: "route", color: "#64748b", width: 1, startArrowhead: "none", endArrowhead: "none", stroke: "solid"},
		{name: "traffic", attrs: `kind="traffic"`, kind: "traffic", color: "#2563eb", width: 1, startArrowhead: "none", endArrowhead: "stealth", stroke: "solid"},
		{name: "overrides", attrs: `kind="traffic" color="#dc2626" stroke-width="3" stroke-style="dotted" start-arrowhead="oval" end-arrowhead="diamond"`, kind: "traffic", color: "#dc2626", width: 3, startArrowhead: "oval", endArrowhead: "diamond", stroke: "dotted"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrow := buildConnectionArrow(t, tt.attrs)
			custom, _ := arrow["customData"].(map[string]any)
			if custom["xaligoConnectorKind"] != tt.kind || arrow["strokeColor"] != tt.color || arrow["strokeWidth"] != tt.width || custom["xaligoConnectorStartArrowhead"] != tt.startArrowhead || custom["xaligoConnectorEndArrowhead"] != tt.endArrowhead || arrow["strokeStyle"] != tt.stroke {
				t.Fatalf("connection arrow = %#v", arrow)
			}
		})
	}
}

func buildConnectionArrow(t *testing.T, attrs string) map[string]any {
	t.Helper()
	doc, err := usecase.Parse(strings.NewReader(`<frame width="320" height="160"><item id="1" /><item id="2" /><connection src="1" dst="2" ` + attrs + ` /></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, doc.Root.Children[2:], nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	for _, element := range scene.Elements {
		if element["type"] == "arrow" {
			return element
		}
	}
	t.Fatalf("arrow not found: %#v", scene.Elements)
	return nil
}
