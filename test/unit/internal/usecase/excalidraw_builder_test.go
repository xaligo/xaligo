package usecase_test

import (
	"encoding/json"
	"strings"
	"testing"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestMakeTextAndImageBuildExpectedElements(t *testing.T) {
	text := usecase.MakeText("label", 1.2, 2.6, 30.4, 10.2, "Hello", 14, "#111111", true, "center", 42)
	if text["type"] != "text" || text["id"] != "label" || text["x"] != float64(1) || text["y"] != float64(3) || text["fontStyle"] != "bold" || text["textAlign"] != "center" {
		t.Fatalf("text = %#v", text)
	}
	image := usecase.MakeImage("icon", 10, 20, 32, 32, "file-id", "#ffffff", 77)
	if image["type"] != "image" || image["fileId"] != "file-id" || image["status"] != "saved" || image["backgroundColor"] != "#ffffff" {
		t.Fatalf("image = %#v", image)
	}
}

func TestRectangleRendersPortsAndText(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame width="320" height="200">
  <rectangle id="service-box" title="Service" width="180" height="100" font-size="18">
    <port id="service-in" side="left" title="in" font-size="9" />
    <port id="service-out" side="right" title="out" font-size="10" />
  </rectangle>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	texts := map[string]float64{}
	rects := 0
	for _, el := range scene.Elements {
		if el["type"] == "rectangle" {
			rects++
		}
		if el["type"] == "text" {
			texts[el["text"].(string)] = el["fontSize"].(float64)
		}
	}
	if rects != 3 {
		t.Fatalf("rectangle count = %d, want body plus two ports", rects)
	}
	if texts["Service"] != 18 || texts["in"] != 9 || texts["out"] != 10 {
		t.Fatalf("texts/font sizes = %#v", texts)
	}
}
