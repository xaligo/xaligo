package usecase_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

// <capture> is a structural annotation group tag: it draws a border (and title
// band) around normally laid-out child content without any AWS/architectural
// semantics, and it is connectable by id exactly like any other group tag,
// including cross-frame page links.

func TestCaptureGroupRendersAnnotationBorder(t *testing.T) {
	input := []byte(`<frame width="400" height="200"><capture id="hot-path" title="Hot Path"><rectangle id="inner" title="Inner" /></capture></frame>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var border map[string]any
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoGroupBorder"] == true {
			border = element
			break
		}
	}
	if border == nil {
		t.Fatalf("capture border rectangle missing: %#v", scene.Elements)
	}
	if border["strokeColor"] != "#F5A623" || border["strokeStyle"] != "dashed" || border["strokeWidth"] != float64(1) {
		t.Fatalf("capture border style = %#v", border)
	}
	sceneTextRectByValue(t, scene.Elements, "Hot Path")
}

func TestCaptureGroupConnectableWithinSameFrame(t *testing.T) {
	input := []byte(`<frame width="400" height="240">
  <container gap="24">
    <capture id="hot-path" title="Hot Path"><rectangle id="a" title="A" /></capture>
    <capture id="slow-path" title="Slow Path"><rectangle id="b" title="B" /></capture>
  </container>
  <connection src="hot-path" dst="slow-path" />
</frame>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var arrow map[string]any
	for _, element := range scene.Elements {
		if element["type"] != "arrow" {
			continue
		}
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] == true {
			continue
		}
		arrow = element
	}
	if arrow == nil {
		t.Fatalf("routed connection between captures missing: %#v", scene.Elements)
	}
	if arrow["startBinding"] == nil || arrow["endBinding"] == nil {
		t.Fatalf("connection is not bound to both capture borders: %#v", arrow)
	}
}

func TestCaptureGroupSupportsCrossFramePageLink(t *testing.T) {
	input := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="overview" width="320" height="180">
    <capture id="hot-path" title="Hot Path"><rectangle id="a" title="A" width="120" height="80" /></capture>
    <connection src="hot-path" dst="detail.slow-path" />
  </frame>
  <frame id="detail" width="320" height="180">
    <capture id="slow-path" title="Slow Path"><rectangle id="b" title="B" width="120" height="80" /></capture>
  </frame>
</frames></xaligo>`)
	out, err := newUsecase().RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var sourceStub, destinationStub map[string]any
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] != true {
			continue
		}
		if element["startBinding"] != nil {
			sourceStub = element
		} else {
			destinationStub = element
		}
	}
	if sourceStub == nil || destinationStub == nil {
		t.Fatalf("cross-frame capture stubs missing: %#v", scene.Elements)
	}
	sceneTextRectByValue(t, scene.Elements, "to <detail>")
	sceneTextRectByValue(t, scene.Elements, "from <overview>")
}

func TestValidateCaptureRequiresID(t *testing.T) {
	err := usecase.Validate(context.Background(), []byte(`<frame><capture title="g" /></frame>`))
	if err == nil {
		t.Fatal("Validate() error = nil, want missing id error")
	}
	if got := err.Error(); got == "" {
		t.Fatal("Validate() error message is empty")
	}
}
