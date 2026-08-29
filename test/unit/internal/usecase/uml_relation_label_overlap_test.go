package usecase_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
)

// TestUMLRelationLabelsAvoidUnrelatedClassifierBoxes is a regression test for a
// bug where a UML relation label (e.g. "specializes") rendered on top of an
// unrelated classifier box's header text, because relation-label placement
// only avoided its own connection's two endpoints and never the other
// classifier/component boxes sharing the same frame.
func TestUMLRelationLabelsAvoidUnrelatedClassifierBoxes(t *testing.T) {
	root := repoRoot(t)
	source, err := os.ReadFile(filepath.Join(root, "docs", "src", "examples", "samples", "uml-class.xal"))
	if err != nil {
		t.Fatal(err)
	}
	out, err := newUsecase().BuildScene(context.Background(), source, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}

	var labels, headers []map[string]any
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if _, ok := custom["xaligoUmlRelationLabel"]; ok && element["type"] == "text" {
			labels = append(labels, element)
			continue
		}
		if custom["xaligoUmlClassHeaderContent"] == true {
			headers = append(headers, element)
		}
	}
	if len(labels) == 0 {
		t.Fatalf("no UML relation labels found: %s", out)
	}
	if len(headers) == 0 {
		t.Fatalf("no UML classifier headers found: %s", out)
	}

	for _, label := range labels {
		labelRect := frameMetadataElementRectEdgeTest(t, label)
		for _, header := range headers {
			headerRect := frameMetadataElementRectEdgeTest(t, header)
			if rectanglesOverlapFrameMetadataTest(labelRect, headerRect) {
				t.Fatalf("relation label %q rect %#v overlaps unrelated classifier header %q rect %#v", label["text"], labelRect, header["id"], headerRect)
			}
		}
	}
}
