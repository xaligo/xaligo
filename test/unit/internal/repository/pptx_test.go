package repository_test

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
)

type fakePptxExporter struct {
	stdout []byte
	stderr []byte
	err    error
	seen   []byte
}

func (rcvr *fakePptxExporter) Export(_ context.Context, _ string, requestJSON []byte) ([]byte, []byte, error) {
	rcvr.seen = append([]byte(nil), requestJSON...)
	return rcvr.stdout, rcvr.stderr, rcvr.err
}

func TestPowerpointRepositoryBuildsExporterRequest(t *testing.T) {
	compression := false
	exporter := &fakePptxExporter{
		stdout: []byte("pptx-bytes"),
		stderr: []byte("exporter warning\n"),
	}
	repo := repository.NewPowerpointRepository(repository.WithPowerpointExportFunc(exporter.Export))

	written, err := repo.ExportPptxBytes(context.Background(), entity.PptxExportOptions{
		PlanJSON:    []byte(`{"slides":[{"name":"main"}]}`),
		Title:       "Example",
		Author:      "xaligo",
		Compression: &compression,
	})
	if err != nil {
		t.Fatal(err)
	}

	if string(written) != "pptx-bytes" {
		t.Fatalf("written PPTX = %q", written)
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

func TestPowerpointRepositoryReturnsPptxBytes(t *testing.T) {
	exporter := &fakePptxExporter{stdout: []byte("pptx-bytes")}
	repo := repository.NewPowerpointRepository(repository.WithPowerpointExportFunc(exporter.Export))
	pptxBytes, err := repo.ExportPptxBytes(context.Background(), entity.PptxExportOptions{
		PlanJSON: []byte(`{"slides":[{"name":"main"}]}`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if string(pptxBytes) != "pptx-bytes" {
		t.Fatalf("PPTX bytes = %q", pptxBytes)
	}
	if strings.Contains(string(exporter.seen), "output") {
		t.Fatalf("request JSON includes output path: %s", exporter.seen)
	}
}

func TestPowerpointRepositoryReturnsExporterError(t *testing.T) {
	exporter := &fakePptxExporter{err: errors.New("native exporter failed")}
	repo := repository.NewPowerpointRepository(repository.WithPowerpointExportFunc(exporter.Export))
	_, err := repo.ExportPptxBytes(context.Background(), entity.PptxExportOptions{PlanJSON: []byte(`{"slides":[]}`)})
	if err == nil || err.Error() != "native exporter failed" {
		t.Fatalf("err = %v", err)
	}
}

func TestPowerpointRepositoryRejectsEmptyExporterOutput(t *testing.T) {
	exporter := &fakePptxExporter{}
	repo := repository.NewPowerpointRepository(repository.WithPowerpointExportFunc(exporter.Export))
	_, err := repo.ExportPptxBytes(context.Background(), entity.PptxExportOptions{PlanJSON: []byte(`{"slides":[]}`)})
	if err == nil || err.Error() != "Rust PPTX exporter produced no output" {
		t.Fatalf("err = %v", err)
	}
}
