package usecase_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
)

func TestFrameMetadataReservedStripOverridesVisibleOverflow(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{
			name: "top item and label",
			source: `<xaligo version="1"><frames>
  <frame id="page" title="Page" width="360" height="200" margin-top="60" overflow="visible">
    <metadata position="top" width="120" key-width="40" />
    <item id="27" name="intruder" dy="-80" />
  </frame>
</frames></xaligo>`,
		},
		{
			name: "bottom box and text",
			source: `<xaligo version="1"><frames>
  <frame id="page" title="Page" width="360" height="200" margin-bottom="60" overflow="visible">
    <metadata position="bottom" width="120" key-width="40" />
    <rectangle id="intruder" title="Intruder" height="170" />
  </frame>
</frames></xaligo>`,
		},
		{
			name: "bottom item and label",
			source: `<xaligo version="1"><frames>
  <frame id="page" title="Page" width="360" height="200" margin-bottom="60" overflow="visible">
    <metadata position="bottom" width="120" key-width="40" />
    <item id="27" name="intruder" dy="100" />
  </frame>
</frames></xaligo>`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := newUsecase().BuildScene(context.Background(), []byte(test.source), entity.RenderOptions{Theme: "light"})
			if err != nil {
				if !strings.Contains(err.Error(), "metadata reserved strip") {
					t.Fatalf("BuildScene() error = %v, want metadata reserved-strip rejection", err)
				}
				return
			}
			var scene sceneFile
			if err := json.Unmarshal(out, &scene); err != nil {
				t.Fatal(err)
			}
			reserved, ok := frameMetadataReservedStripEdgeTest(t, scene.Elements, "page")
			if !ok {
				t.Fatalf("frame metadata reserved strip is missing: %s", out)
			}
			ordinaryElements := 0
			for _, element := range scene.Elements {
				custom, _ := element["customData"].(map[string]any)
				if custom["xaligoFrameMetadata"] == true || custom["xaligoFrameMetadataReserved"] == true || custom["xaligoPageFrame"] == true {
					continue
				}
				typeName, _ := element["type"].(string)
				if typeName != "rectangle" && typeName != "image" && typeName != "text" {
					continue
				}
				ordinaryElements++
				rectangle := frameMetadataElementRectEdgeTest(t, element)
				if rectanglesOverlapFrameMetadataTest(rectangle, reserved) {
					t.Fatalf("overflow=visible ordinary %s %q rect %#v overlaps metadata reserved strip %#v", typeName, element["id"], rectangle, reserved)
				}
			}
			if ordinaryElements == 0 {
				t.Fatalf("render succeeded without the ordinary box/text/item under test: %s", out)
			}
		})
	}
}

func TestNarrowCrossFramePageLinkLabelsStayOutsideMetadataReservedStrip(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="30">
  <frame id="a" title="A" width="200" height="70">
    <metadata width="192" key-width="50" />
    <rectangle id="s" title="S" />
    <connection src="s" dst="b.d" />
  </frame>
  <frame id="b" title="B" width="200" height="70">
    <metadata width="192" key-width="50" />
    <rectangle id="d" title="D" />
  </frame>
</frames></xaligo>`)
	scene := renderFrameMetadataEdgeScene(t, source)
	reserved := map[string][4]float64{}
	for _, frameID := range []string{"a", "b"} {
		strip, ok := frameMetadataReservedStripEdgeTest(t, scene.Elements, frameID)
		if !ok {
			t.Fatalf("frame %q metadata reserved strip is missing", frameID)
		}
		reserved[frameID] = strip
	}
	wanted := map[string]string{"to <b>": "a", "from <a>": "b"}
	found := map[string]bool{}
	for _, element := range scene.Elements {
		label, _ := element["text"].(string)
		frameID, ok := wanted[label]
		if !ok {
			continue
		}
		found[label] = true
		rectangle := frameMetadataElementRectEdgeTest(t, element)
		if rectanglesOverlapFrameMetadataTest(rectangle, reserved[frameID]) {
			t.Fatalf("narrow page-link label %q rect %#v overlaps frame %q metadata reserved strip %#v", label, rectangle, frameID, reserved[frameID])
		}
	}
	for label := range wanted {
		if !found[label] {
			t.Fatalf("narrow page-link label %q is missing: %#v", label, scene.Elements)
		}
	}
}

func TestNarrowUMLRelationLabelStaysOutsideMetadataReservedStrip(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames>
  <frame id="page" title="Page" width="200" height="70" layout="horizontal" gap="20">
    <metadata width="192" key-width="50" />
    <rectangle id="source" title="Source" />
    <rectangle id="destination" title="Destination" />
    <connection src="source" dst="destination" uml-relation-label="calls" />
  </frame>
</frames></xaligo>`)
	scene := renderFrameMetadataEdgeScene(t, source)
	reserved, ok := frameMetadataReservedStripEdgeTest(t, scene.Elements, "page")
	if !ok {
		t.Fatal("frame metadata reserved strip is missing")
	}
	for _, element := range scene.Elements {
		if element["text"] != "calls" {
			continue
		}
		label := frameMetadataElementRectEdgeTest(t, element)
		if rectanglesOverlapFrameMetadataTest(label, reserved) {
			t.Fatalf("narrow UML relation label rect %#v overlaps metadata reserved strip %#v", label, reserved)
		}
		return
	}
	t.Fatalf("narrow UML relation label is missing: %#v", scene.Elements)
}

func renderFrameMetadataEdgeScene(t *testing.T, source []byte) sceneFile {
	t.Helper()
	out, err := newUsecase().BuildScene(context.Background(), source, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	return scene
}

func frameMetadataReservedStripEdgeTest(t *testing.T, elements []map[string]any, frameID string) ([4]float64, bool) {
	t.Helper()
	for _, element := range elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoFrameMetadataReserved"] != true || custom["xaligoFrameID"] != frameID {
			continue
		}
		return frameMetadataElementRectEdgeTest(t, element), true
	}
	return [4]float64{}, false
}

func frameMetadataElementRectEdgeTest(t *testing.T, element map[string]any) [4]float64 {
	t.Helper()
	return [4]float64{
		sceneNumber(t, element["x"]),
		sceneNumber(t, element["y"]),
		sceneNumber(t, element["width"]),
		sceneNumber(t, element["height"]),
	}
}
