package repository_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/ryo-arima/xaligo/internal/repository"
)

func TestLoadIsoflowIconManifest(t *testing.T) {
	repo := repository.NewIsoflowRepository()
	manifest := `{"icons":{"lambda":{"dataURL":"data:image/svg+xml;base64,QQ=="},"empty":{}}}`
	path := filepath.Join(t.TempDir(), "icons.json")
	if err := os.WriteFile(path, []byte(manifest), 0644); err != nil {
		t.Fatal(err)
	}
	icons, err := repo.LoadIsoflowIconManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(icons) != 1 || icons["lambda"] != "data:image/svg+xml;base64,QQ==" {
		t.Fatalf("icons = %#v", icons)
	}
}

func TestLoadIsoflowIconManifestFSAndErrors(t *testing.T) {
	repo := repository.NewIsoflowRepository()
	fsys := fstest.MapFS{"icons.json": {Data: []byte(`{"icons":{"s3":{"dataURL":"data:image/svg+xml;base64,Uw=="}}}`)}}
	icons, err := repo.LoadIsoflowIconManifestFS(fsys, "icons.json")
	if err != nil {
		t.Fatal(err)
	}
	if icons["s3"] == "" {
		t.Fatalf("icons = %#v", icons)
	}
	if _, err := repo.LoadIsoflowIconManifestFS(fstest.MapFS{"bad.json": {Data: []byte(`{"icons":`)}}, "bad.json"); err == nil {
		t.Fatal("LoadIsoflowIconManifestFS invalid JSON error = nil")
	}
	if _, err := repo.LoadIsoflowIconManifest(filepath.Join(t.TempDir(), "missing.json")); err == nil {
		t.Fatal("LoadIsoflowIconManifest missing file error = nil")
	}
}
