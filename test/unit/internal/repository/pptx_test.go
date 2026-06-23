package repository_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
)

type fakePptxExporter struct {
	stdout []byte
	stderr []byte
	err    error
	seen   []byte
}

func (f *fakePptxExporter) Export(_ context.Context, requestJSON []byte) ([]byte, []byte, error) {
	f.seen = append([]byte(nil), requestJSON...)
	return f.stdout, f.stderr, f.err
}

func TestExportPptxWithExporterWritesPptxOutput(t *testing.T) {
	output := filepath.Join(t.TempDir(), "out.pptx")
	compression := false
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	exporter := &fakePptxExporter{
		stdout: []byte("pptx-bytes"),
		stderr: []byte("exporter warning\n"),
	}

	err := repository.ExportPptxWithExporter(context.Background(), entity.PptxExportOptions{
		PlanJSON:    []byte(`{"slides":[{"name":"main"}]}`),
		Output:      output,
		Title:       "Example",
		Author:      "xaligo",
		Compression: &compression,
		Stdout:      stdout,
		Stderr:      stderr,
	}, exporter)
	if err != nil {
		t.Fatal(err)
	}

	written, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if string(written) != "pptx-bytes" {
		t.Fatalf("written PPTX = %q", written)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q", stdout.String())
	}
	if stderr.String() != "exporter warning\n" {
		t.Fatalf("stderr = %q", stderr.String())
	}

	var request struct {
		Plan    json.RawMessage `json:"plan"`
		Options struct {
			Title       string `json:"title"`
			Author      string `json:"author"`
			Compression *bool  `json:"compression"`
		} `json:"options"`
	}
	if err := json.Unmarshal(exporter.seen, &request); err != nil {
		t.Fatalf("request JSON = %s: %v", exporter.seen, err)
	}
	if string(request.Plan) != `{"slides":[{"name":"main"}]}` {
		t.Fatalf("plan = %s", request.Plan)
	}
	if request.Options.Title != "Example" || request.Options.Author != "xaligo" || request.Options.Compression == nil || *request.Options.Compression {
		t.Fatalf("options = %#v", request.Options)
	}
}

func TestExportPptxWithExporterReturnsExporterError(t *testing.T) {
	exporter := &fakePptxExporter{err: errors.New("wasm failed")}
	err := repository.ExportPptxWithExporter(context.Background(), entity.PptxExportOptions{
		PlanJSON: []byte(`{"slides":[]}`),
		Output:   filepath.Join(t.TempDir(), "out.pptx"),
	}, exporter)
	if err == nil || err.Error() != "wasm failed" {
		t.Fatalf("err = %v", err)
	}
}

func TestExportPptxWithExporterRejectsEmptyExporterOutput(t *testing.T) {
	exporter := &fakePptxExporter{}
	err := repository.ExportPptxWithExporter(context.Background(), entity.PptxExportOptions{
		PlanJSON: []byte(`{"slides":[]}`),
		Output:   filepath.Join(t.TempDir(), "out.pptx"),
	}, exporter)
	if err == nil || err.Error() != "PPTX WASM exporter produced no output" {
		t.Fatalf("err = %v", err)
	}
}

func TestExportPptxUsesWASMExporterAndReportsMissingPath(t *testing.T) {
	t.Setenv("XALIGO_PPTX_EXPORTER_WASM", filepath.Join(t.TempDir(), "missing-env.wasm"))
	err := repository.ExportPptx(entity.PptxExportOptions{
		PlanJSON:     []byte(`{"slides":[]}`),
		Output:       filepath.Join(t.TempDir(), "out.pptx"),
		ExporterWASM: filepath.Join(t.TempDir(), "missing-explicit.wasm"),
	})
	if err == nil || (!strings.Contains(err.Error(), "PPTX WASM exporter not found") && !strings.Contains(err.Error(), "produced no output")) {
		t.Fatalf("err = %v", err)
	}
}

func TestWASMPptxExporterReportsInvalidWASM(t *testing.T) {
	wasmPath := filepath.Join(t.TempDir(), "bad.wasm")
	if err := os.WriteFile(wasmPath, []byte("not wasm"), 0644); err != nil {
		t.Fatal(err)
	}
	exporter := repository.WASMPptxExporter{Path: wasmPath}
	_, _, err := exporter.Export(context.Background(), []byte(`{}`))
	if err == nil || !strings.Contains(err.Error(), "run PPTX WASM exporter") {
		t.Fatalf("err = %v", err)
	}
}
