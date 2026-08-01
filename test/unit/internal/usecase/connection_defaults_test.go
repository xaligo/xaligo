package usecase_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
)

func TestConnectionsChildAliasesOverrideParentSemanticDefaults(t *testing.T) {
	arrow := renderGroupedConnectionV1DefaultsTest(t,
		`stroke-width="2" end-arrowhead="diamond" coordinate-scale="2"`,
		`width="3" arrowhead="oval" scale="4"`,
	)
	custom, _ := arrow["customData"].(map[string]any)
	if arrow["strokeWidth"] != 3.0 {
		t.Fatalf("strokeWidth = %#v, want child width alias 3", arrow["strokeWidth"])
	}
	if custom["xaligoConnectorEndArrowhead"] != "oval" {
		t.Fatalf("end arrowhead = %#v, want child arrowhead alias oval", custom["xaligoConnectorEndArrowhead"])
	}
	if custom["xaligoConnectorScale"] != 4.0 {
		t.Fatalf("coordinate scale = %#v, want child scale alias 4", custom["xaligoConnectorScale"])
	}
}

func TestConnectionsChildCanonicalAttributesKeepAliasPrecedence(t *testing.T) {
	arrow := renderGroupedConnectionV1DefaultsTest(t,
		`stroke-width="2" end-arrowhead="diamond" coordinate-scale="2"`,
		`stroke-width="5" width="6" end-arrowhead="triangle" arrowhead="oval" coordinate-scale="7" scale="8"`,
	)
	custom, _ := arrow["customData"].(map[string]any)
	if arrow["strokeWidth"] != 5.0 {
		t.Fatalf("strokeWidth = %#v, want canonical stroke-width 5", arrow["strokeWidth"])
	}
	if custom["xaligoConnectorEndArrowhead"] != "triangle" {
		t.Fatalf("end arrowhead = %#v, want canonical end-arrowhead triangle", custom["xaligoConnectorEndArrowhead"])
	}
	if custom["xaligoConnectorScale"] != 7.0 {
		t.Fatalf("coordinate scale = %#v, want canonical coordinate-scale 7", custom["xaligoConnectorScale"])
	}
}

func TestV1RouteChildNoneSuppressesNonHeadlessParentAliasDefault(t *testing.T) {
	arrow := renderGroupedConnectionV1DefaultsTest(t,
		`kind="route" end-arrowhead="triangle"`,
		`arrowhead="none"`,
	)
	custom, _ := arrow["customData"].(map[string]any)
	if arrow["startArrowhead"] != nil || arrow["endArrowhead"] != nil {
		t.Fatalf("route Excalidraw arrowheads = %#v/%#v, want nil/nil", arrow["startArrowhead"], arrow["endArrowhead"])
	}
	if custom["xaligoConnectorStartArrowhead"] != "none" || custom["xaligoConnectorEndArrowhead"] != "none" {
		t.Fatalf("route logical arrowheads = %#v/%#v, want none/none", custom["xaligoConnectorStartArrowhead"], custom["xaligoConnectorEndArrowhead"])
	}
}

func renderGroupedConnectionV1DefaultsTest(t *testing.T, groupAttrs, childAttrs string) map[string]any {
	t.Helper()
	source := []byte(`<frame width="420" height="180" layout="horizontal">
  <rectangle id="source" title="Source" width="120" height="80" />
  <rectangle id="target" title="Target" width="120" height="80" />
  <connections ` + groupAttrs + `>
    <connection src="source" dst="target" ` + childAttrs + ` />
  </connections>
</frame>`)
	out, err := newUsecase().BuildScene(context.Background(), source, entity.RenderOptions{Theme: "light"})
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
