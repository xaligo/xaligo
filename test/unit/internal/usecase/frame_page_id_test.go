package usecase_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestFramePageIdentityDoesNotCollapsePunctuation(t *testing.T) {
	input := []byte(`<xaligo version="1">
  <data></data>
  <frames gap="32">
    <frame id="a/b" width="240" height="120"><rectangle id="left" title="Left" width="100" height="60" /><connection src="left" dst="a:b.right" /></frame>
    <frame id="a:b" width="240" height="120"><rectangle id="right" title="Right" width="100" height="60" /></frame>
  </frames>
</xaligo>`)

	uc := newUsecase()
	sceneJSON, err := uc.RenderExcalidraw(context.Background(), input, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		t.Fatal(err)
	}
	seenIDs := map[string]bool{}
	pageFrameIDs := map[string]string{}
	for _, element := range scene.Elements {
		id, _ := element["id"].(string)
		if id == "" {
			t.Fatalf("scene element has no ID: %#v", element)
		}
		if seenIDs[id] {
			t.Fatalf("duplicate scene element ID %q", id)
		}
		seenIDs[id] = true
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoPageFrame"] == true {
			frameID, _ := custom["xaligoFrameID"].(string)
			pageFrameIDs[frameID] = id
		}
	}
	if pageFrameIDs["a/b"] != "paper-frame-a/b" || pageFrameIDs["a:b"] != "paper-frame-a:b" {
		t.Fatalf("page frame IDs = %#v", pageFrameIDs)
	}

	planJSON, err := uc.BuildPPTXPlan(context.Background(), input, entity.RenderOptions{Format: usecase.FormatPPTX, Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var document entity.DocumentPlan
	if err := json.Unmarshal(planJSON, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Pages) != 2 || document.Pages[0].ID != "a/b" || document.Pages[1].ID != "a:b" {
		t.Fatalf("document pages = %#v", document.Pages)
	}
	assertFramePageContent(t, document.Pages[0], "Left", "Right", "to <a:b>")
	assertFramePageContent(t, document.Pages[1], "Right", "Left", "from <a/b>")
}

func assertFramePageContent(t *testing.T, page entity.DocumentPage, wantText, unwantedText, wantPageLink string) {
	t.Helper()
	rectangles := 0
	foundText := false
	foundPageLink := false
	for _, op := range page.Ops {
		if op.Kind == "rect" {
			rectangles++
		}
		if op.Kind == "text" && op.Text == wantText {
			foundText = true
		}
		if op.Kind == "text" && op.Text == wantPageLink {
			foundPageLink = true
		}
		if op.Kind == "text" && op.Text == unwantedText {
			t.Fatalf("page %q contains text from another frame: %#v", page.ID, op)
		}
	}
	if rectangles != 1 || !foundText || !foundPageLink {
		t.Fatalf("page %q rectangles=%d found %q=%t page link %q=%t ops=%#v", page.ID, rectangles, wantText, foundText, wantPageLink, foundPageLink, page.Ops)
	}
}
