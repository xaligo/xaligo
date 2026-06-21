package repository_test

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ryo-arima/xaligo/internal/repository"
)

func TestSVGBGColorSkipsYamahaHiddenReferenceFills(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	svgPath := filepath.Join(root, "etc", "resources", "aws", "svg", "Yamaha-Network-Icons", "router_general.svg")
	raw, err := os.ReadFile(svgPath)
	if err != nil {
		t.Fatal(err)
	}
	dataURL := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(raw)
	if got := repository.SVGBGColor(dataURL); got != "#7758c1" {
		t.Fatalf("SVGBGColor(router_general) = %q, want #7758c1", got)
	}
}
