package excalidraw

import (
	"encoding/json"
	"testing"
)

func TestApplyThemeJSONDark(t *testing.T) {
	source := []byte(`{"type":"excalidraw","version":2,"elements":[{"type":"text","strokeColor":"#1e1e1e","backgroundColor":"#ffffff"},{"type":"rectangle","strokeColor":"#8C4FFF","backgroundColor":"transparent"}],"appState":{"viewBackgroundColor":"#ffffff"},"files":{}}`)
	got, err := ApplyThemeJSON(source, "dark")
	if err != nil {
		t.Fatal(err)
	}
	var scene file
	if err := json.Unmarshal(got, &scene); err != nil {
		t.Fatal(err)
	}
	if scene.AppState["viewBackgroundColor"] != "#111827" {
		t.Fatalf("background = %v", scene.AppState["viewBackgroundColor"])
	}
	if scene.Elements[0]["strokeColor"] != "#e5e7eb" || scene.Elements[0]["backgroundColor"] != "#111827" {
		t.Fatalf("neutral colors were not themed: %#v", scene.Elements[0])
	}
	if scene.Elements[1]["strokeColor"] != "#8C4FFF" {
		t.Fatalf("semantic color changed: %#v", scene.Elements[1])
	}
}

func TestApplyThemeJSONRejectsUnknownTheme(t *testing.T) {
	if _, err := ApplyThemeJSON([]byte(`{}`), "sepia"); err == nil {
		t.Fatal("expected invalid theme error")
	}
}

func TestApplyThemeJSONLightPreservesBytes(t *testing.T) {
	source := []byte(`{"appState":{}}`)
	got, err := ApplyThemeJSON(source, "")
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(source) {
		t.Fatalf("light theme changed scene: %s", got)
	}
}
