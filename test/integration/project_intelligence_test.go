//go:build cgo && xaligo_engine && sqlite_fts5

package integration_test

import (
	"context"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	projectrepository "github.com/xaligo/xaligo/internal/repository/project"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestV2ProjectAnalysisFeedsRAGIndex(t *testing.T) {
	path := filepath.Join("..", "..", "docs", "src", "examples", "samples", "complex-hybrid-architecture-v2.xal")
	source, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	absolute, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	uri := (&url.URL{Scheme: "file", Path: filepath.ToSlash(absolute)}).String()
	index := projectrepository.NewIndexRepository(filepath.Join(t.TempDir(), "project.db"))
	project := usecase.NewProjectUsecase(index, v2.NewFrontendUsecase(), v2.NewEngineUsecase())
	t.Cleanup(func() {
		if closeErr := project.Close(); closeErr != nil {
			t.Errorf("close project index: %v", closeErr)
		}
	})

	analysis, changed, err := project.IndexDocument(context.Background(), uri, source)
	if err != nil {
		t.Fatal(err)
	}
	if !changed || len(analysis.Diagnostics) != 0 || len(analysis.Symbols) == 0 {
		t.Fatalf("V2 indexed analysis = changed %t symbols %d diagnostics %#v", changed, len(analysis.Symbols), analysis.Diagnostics)
	}
	symbols, err := project.Symbols(context.Background(), uri)
	if err != nil {
		t.Fatal(err)
	}
	if len(symbols) != len(analysis.Symbols) {
		t.Fatalf("stored V2 symbols = %d, want %d", len(symbols), len(analysis.Symbols))
	}
	results, err := project.Search(context.Background(), "Headquarters", 10)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, result := range results {
		if result.URI == uri && result.ID == "onprem-hq" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("V2 concept missing from RAG search: %#v", results)
	}
}
