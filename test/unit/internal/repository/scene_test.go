package repository_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
)

func TestReadSceneInitializesMissingCollections(t *testing.T) {
	repo := repository.NewExcalidrawRepository()
	path := filepath.Join(t.TempDir(), "scene.excalidraw")
	if err := os.WriteFile(path, []byte(`{"type":"excalidraw","version":2}`), 0644); err != nil {
		t.Fatal(err)
	}
	scene, err := repo.ReadScene(path)
	if err != nil {
		t.Fatal(err)
	}
	if scene.Files == nil || scene.Elements == nil {
		t.Fatalf("scene collections were not initialized: %#v", scene)
	}
}

func TestWriteSceneRoundTrip(t *testing.T) {
	repo := repository.NewExcalidrawRepository()
	path := filepath.Join(t.TempDir(), "scene.excalidraw")
	scene := entity.NewScene()
	scene.Elements = append(scene.Elements, map[string]interface{}{"id": "box", "type": "rectangle"})
	if err := repo.WriteScene(scene, path); err != nil {
		t.Fatal(err)
	}
	read, err := repo.ReadScene(path)
	if err != nil {
		t.Fatal(err)
	}
	if read.Type != "excalidraw" || len(read.Elements) != 1 || read.Elements[0]["id"] != "box" {
		t.Fatalf("read scene = %#v", read)
	}
}

func TestReadSceneAndWriteSceneErrors(t *testing.T) {
	repo := repository.NewExcalidrawRepository()
	dir := t.TempDir()
	badJSON := filepath.Join(dir, "bad.excalidraw")
	if err := os.WriteFile(badJSON, []byte(`{"elements":`), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := repo.ReadScene(filepath.Join(dir, "missing.excalidraw")); err == nil {
		t.Fatal("ReadScene missing file error = nil")
	}
	if _, err := repo.ReadScene(badJSON); err == nil {
		t.Fatal("ReadScene invalid JSON error = nil")
	}
	if err := repo.WriteScene(&entity.Scene{Files: map[string]map[string]interface{}{"bad": {"fn": func() {}}}}, filepath.Join(dir, "out.excalidraw")); err == nil {
		t.Fatal("WriteScene marshal error = nil")
	}
	if err := repo.WriteScene(entity.NewScene(), filepath.Join(dir, "missing", "out.excalidraw")); err == nil {
		t.Fatal("WriteScene write error = nil")
	}
	var decoded entity.Scene
	data, err := os.ReadFile(badJSON)
	if err != nil {
		t.Fatal(err)
	}
	if json.Unmarshal(data, &decoded) == nil {
		t.Fatal("test fixture should be invalid JSON")
	}
}
