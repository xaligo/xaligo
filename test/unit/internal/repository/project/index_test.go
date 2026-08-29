//go:build cgo && sqlite_fts5

package project_test

import (
	"context"
	"crypto/sha256"
	"path/filepath"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	projectrepository "github.com/xaligo/xaligo/internal/repository/project"
)

func projectAnalysis(uri, detail string) entity.ProjectAnalysis {
	source := []byte(detail)
	return entity.ProjectAnalysis{
		URI: uri, Kind: entity.ProjectDocumentMarkdown,
		Checksum: sha256.Sum256(source), Source: source,
		Symbols: []entity.ProjectSymbol{{
			Ordinal: 0, ParentOrdinal: -1, ID: "database", Name: "Database",
			Detail: detail, Concept: entity.ProjectConceptText, SourceTag: "heading",
			Position: entity.Position{Line: 4, Column: 1},
		}},
	}
}

func TestProjectIndexIncrementalSearchReplacementAndPrune(t *testing.T) {
	ctx := context.Background()
	repository := projectrepository.NewIndexRepository(filepath.Join(t.TempDir(), "project.db"))
	t.Cleanup(func() { _ = repository.Close() })
	uri := "file:///project/docs/database.md"
	root := "file:///project/docs"

	changed, err := repository.Put(ctx, root, projectAnalysis(uri, "Relational database storage"))
	if err != nil || !changed {
		t.Fatalf("initial Put changed=%v err=%v", changed, err)
	}
	changed, err = repository.Put(ctx, root, projectAnalysis(uri, "Relational database storage"))
	if err != nil || changed {
		t.Fatalf("unchanged Put changed=%v err=%v", changed, err)
	}
	results, err := repository.Search(ctx, "relational", 10)
	if err != nil || len(results) != 1 || results[0].Line != 4 {
		t.Fatalf("search results=%#v err=%v", results, err)
	}
	symbols, err := repository.Symbols(ctx, uri)
	if err != nil || len(symbols) != 1 || symbols[0].ID != "database" {
		t.Fatalf("symbols=%#v err=%v", symbols, err)
	}

	changed, err = repository.Put(ctx, root, projectAnalysis(uri, "Vector search index"))
	if err != nil || !changed {
		t.Fatalf("updated Put changed=%v err=%v", changed, err)
	}
	stale, err := repository.Search(ctx, "relational", 10)
	if err != nil || len(stale) != 0 {
		t.Fatalf("stale search=%#v err=%v", stale, err)
	}
	current, err := repository.Search(ctx, "vector", 10)
	if err != nil || len(current) != 1 {
		t.Fatalf("current search=%#v err=%v", current, err)
	}
	removed, err := repository.Prune(ctx, root, nil)
	if err != nil || removed != 1 {
		t.Fatalf("Prune removed=%d err=%v", removed, err)
	}
	current, err = repository.Search(ctx, "vector", 10)
	if err != nil || len(current) != 0 {
		t.Fatalf("search after prune=%#v err=%v", current, err)
	}
}
