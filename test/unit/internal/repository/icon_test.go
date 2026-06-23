package repository_test

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/ryo-arima/xaligo/internal/repository"
)

const testSVG = `<svg xmlns="http://www.w3.org/2000/svg"><rect fill="#123456"/></svg>`

func TestSVGBGColorSkipsYamahaHiddenReferenceFills(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", "..", ".."))
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

func TestSvgToDataURLAndFileID(t *testing.T) {
	dir := t.TempDir()
	svgPath := filepath.Join(dir, "icon.svg")
	if err := os.WriteFile(svgPath, []byte(testSVG), 0644); err != nil {
		t.Fatal(err)
	}
	dataURL, err := repository.SvgToDataURL(svgPath)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(dataURL, "data:image/svg+xml;base64,") || repository.SVGBGColor(dataURL) != "#123456" {
		t.Fatalf("unexpected data URL %q", dataURL)
	}
	if len(repository.FileID(svgPath)) != 16 || repository.FileID(svgPath) != repository.FileID(svgPath) {
		t.Fatalf("FileID is not stable 16-char hex")
	}
}

func TestSvgToDataURLFS(t *testing.T) {
	fsys := fstest.MapFS{"icons/icon.svg": {Data: []byte(testSVG)}}
	dataURL, err := repository.SvgToDataURLFS(fsys, "icons/icon.svg")
	if err != nil {
		t.Fatal(err)
	}
	if got := repository.SVGBGColor(dataURL); got != "#123456" {
		t.Fatalf("SVGBGColor() = %q, want #123456", got)
	}
}

func TestLoadFromCSVAndLookupCatalog(t *testing.T) {
	csvPath := filepath.Join(t.TempDir(), "catalog.csv")
	encoded := base64.StdEncoding.EncodeToString([]byte(testSVG))
	content := strings.Join([]string{
		"# comment",
		"bad,row",
		"1,Compute,Amazon EC2,ec2.svg,Architecture-Service-Icons/ec2.svg," + encoded,
		"2,Database,Amazon RDS,rds.svg,Architecture-Service-Icons/rds.svg," + encoded,
	}, "\n")
	if err := os.WriteFile(csvPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	dataURL, err := repository.LoadFromCSV(csvPath, "ec2.svg")
	if err != nil {
		t.Fatal(err)
	}
	if got := repository.SVGBGColor(dataURL); got != "#123456" {
		t.Fatalf("LoadFromCSV color = %q, want #123456", got)
	}

	relPath, dataURL, err := repository.LoadFromCSVByID(csvPath, 2, "RDS")
	if err != nil {
		t.Fatal(err)
	}
	if relPath != "Architecture-Service-Icons/rds.svg" || !strings.HasPrefix(dataURL, "data:image/svg+xml;base64,") {
		t.Fatalf("LoadFromCSVByID() = %q %q", relPath, dataURL)
	}

	entry, err := repository.LookupCatalogByID(csvPath, 1)
	if err != nil {
		t.Fatal(err)
	}
	if entry.ID != 1 || entry.Category != "Compute" || entry.Service != "Amazon EC2" || entry.SVGFile != "ec2.svg" || repository.SVGBGColor(entry.DataURL) != "#123456" {
		t.Fatalf("entry = %#v", entry)
	}
}

func TestLookupCatalogByIDFS(t *testing.T) {
	encoded := base64.StdEncoding.EncodeToString([]byte(testSVG))
	fsys := fstest.MapFS{"catalog.csv": {Data: []byte(strings.Join([]string{
		"bad,row",
		"abc,Skip,Bad,bad.svg,bad.svg," + encoded,
		"6,Skip,Wrong,wrong.svg,wrong.svg," + encoded,
		"7,Network,Amazon VPC,vpc.svg,vpc.svg," + encoded,
	}, "\n"))}}
	entry, err := repository.LookupCatalogByIDFS(fsys, "catalog.csv", 7)
	if err != nil {
		t.Fatal(err)
	}
	if entry.ID != 7 || entry.Service != "Amazon VPC" || repository.SVGBGColor(entry.DataURL) != "#123456" {
		t.Fatalf("entry = %#v", entry)
	}
	if _, err := repository.LookupCatalogByIDFS(fsys, "catalog.csv", 99); err == nil {
		t.Fatal("LookupCatalogByIDFS missing ID error = nil")
	}
	if _, err := repository.LookupCatalogByIDFS(fsys, "missing.csv", 7); err == nil {
		t.Fatal("LookupCatalogByIDFS missing file error = nil")
	}
}

func TestCatalogLookupErrors(t *testing.T) {
	csvPath := filepath.Join(t.TempDir(), "catalog.csv")
	if err := os.WriteFile(csvPath, []byte("1,Compute,Amazon EC2,ec2.svg,ec2.svg,\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := repository.LoadFromCSV(csvPath, "missing.svg"); err == nil {
		t.Fatal("LoadFromCSV missing icon error = nil")
	}
	if _, err := repository.LoadFromCSV(csvPath, "ec2.svg"); err == nil {
		t.Fatal("LoadFromCSV empty base64 error = nil")
	}
	if _, _, err := repository.LoadFromCSVByID(csvPath, 1, "different service"); err == nil {
		t.Fatal("LoadFromCSVByID name mismatch error = nil")
	}
	if _, err := repository.LookupCatalogByID(csvPath, 99); err == nil {
		t.Fatal("LookupCatalogByID missing ID error = nil")
	}
	if _, err := repository.SvgToDataURL(filepath.Join(t.TempDir(), "missing.svg")); err == nil {
		t.Fatal("SvgToDataURL missing file error = nil")
	}
	if _, err := repository.SvgToDataURLFS(fstest.MapFS{}, "missing.svg"); err == nil {
		t.Fatal("SvgToDataURLFS missing file error = nil")
	}
	if got := repository.SVGBGColor("not-a-data-url"); got != "transparent" {
		t.Fatalf("SVGBGColor invalid data = %q", got)
	}
}
