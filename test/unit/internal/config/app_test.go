package config_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/config"
)

func TestNewResolvesDefaultConfigPaths(t *testing.T) {
	cfg := config.New()
	if cfg.ProjectRoot == "" || !filepath.IsAbs(cfg.ProjectRoot) {
		t.Fatalf("ProjectRoot = %q", cfg.ProjectRoot)
	}
	paths := []string{cfg.AssetDir(), cfg.OutputFramesDir(), cfg.ServiceCatalogCSVPath(), cfg.PptxExporterWASM}
	for _, path := range paths {
		if path == "" || !filepath.IsAbs(path) {
			t.Fatalf("path = %q, want absolute", path)
		}
	}
	if cfg.Legend.IconSize <= 0 || cfg.Legend.FontSize <= 0 || cfg.ItemIconSize <= 0 {
		t.Fatalf("config sizes = %#v", cfg)
	}
}

func TestNewUsesProjectLocalYAMLAndAbsolutePaths(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module temp\n"), 0644); err != nil {
		t.Fatal(err)
	}
	configDir := filepath.Join(dir, "etc", "resources", "aws")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatal(err)
	}
	absCatalog := filepath.Join(dir, "catalog.csv")
	yaml := []byte("paths:\n  asset_package: custom/assets\n  service_catalog_csv: " + absCatalog + "\n  output_frames: custom/output\n  pptx_exporter_wasm: custom/exporter.wasm\nlegend:\n  offset_x: 7\n  offset_y: 8\n  icon_size: 9\n  font_size: 10\nitem:\n  icon_size: 11\n")
	if err := os.WriteFile(filepath.Join(configDir, "app.yaml"), yaml, 0644); err != nil {
		t.Fatal(err)
	}
	oldWD, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(filepath.Join(dir, "etc")); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldWD) })

	cfg := config.New()
	wantRoot, err := filepath.EvalSymlinks(dir)
	if err != nil {
		t.Fatal(err)
	}
	gotRoot, err := filepath.EvalSymlinks(cfg.ProjectRoot)
	if err != nil {
		t.Fatal(err)
	}
	if gotRoot != wantRoot {
		t.Fatalf("ProjectRoot = %q, want %q", cfg.ProjectRoot, dir)
	}
	if !filepath.IsAbs(cfg.AssetDir()) || !strings.HasSuffix(filepath.ToSlash(cfg.AssetDir()), "/custom/assets") || cfg.ServiceCatalogCSVPath() != absCatalog || !filepath.IsAbs(cfg.PptxExporterWASM) || !strings.HasSuffix(filepath.ToSlash(cfg.PptxExporterWASM), "/custom/exporter.wasm") {
		t.Fatalf("paths = %#v", cfg)
	}
	if cfg.Legend.OffsetX != 7 || cfg.Legend.OffsetY != 8 || cfg.Legend.IconSize != 9 || cfg.Legend.FontSize != 10 || cfg.ItemIconSize != 11 {
		t.Fatalf("values = %#v", cfg)
	}
}

func TestNewUsesXaligoHome(t *testing.T) {
	home := t.TempDir()
	configDir := filepath.Join(home, "etc", "resources", "aws")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(configDir, "app.yaml"), []byte("paths:\n  pptx_exporter_wasm: runtime/exporter.wasm\n"), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("XALIGO_HOME", home)

	cfg := config.New()
	if cfg.ProjectRoot != home {
		t.Fatalf("ProjectRoot = %q, want %q", cfg.ProjectRoot, home)
	}
	if cfg.PptxExporterWASM != filepath.Join(home, "runtime", "exporter.wasm") {
		t.Fatalf("PptxExporterWASM = %q", cfg.PptxExporterWASM)
	}
}
