//go:build cgo && xaligo_engine && sqlite_fts5

package integration_test

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"

	"github.com/xaligo/xaligo/internal/core/profiles/builtin"
	"github.com/xaligo/xaligo/internal/entity"
	iconrepository "github.com/xaligo/xaligo/internal/repository/icon"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestIconUsecaseValidatesWithRustAndPersistsInSQLite(t *testing.T) {
	registry := iconrepository.NewRegistryRepository(filepath.Join(t.TempDir(), "xaligo-assets.db"))
	t.Cleanup(func() { _ = registry.Close() })
	icons := v2.NewIconUsecase(registry, v2.NewEngineUsecase())
	ctx := context.Background()

	stored, err := icons.Put(ctx, entity.IconRegistration{
		Reference:   "Builtin:Database",
		SVG:         []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="24" height="12"><path d="M0 0h24v12z"/></svg>`),
		Description: "Generic database",
		Tags:        []string{"Storage", "Database"},
		Aliases:     []string{"db"},
		License:     "MIT",
		Source:      "builtin",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stored.Ref.String() != "builtin:database" || stored.ViewBox != "0 0 24 12" || !bytes.Contains(stored.SVG, []byte(`<svg`)) {
		t.Fatalf("stored icon = %#v %s", stored, stored.SVG)
	}
	byAlias, err := icons.Get(ctx, "builtin:db")
	if err != nil || byAlias.Ref != stored.Ref {
		t.Fatalf("alias result = %#v, %v", byAlias, err)
	}
	results, err := icons.Search(ctx, "storage", 10)
	if err != nil || len(results) != 1 || results[0].Ref != stored.Ref {
		t.Fatalf("search results = %#v, %v", results, err)
	}

	if _, err := icons.Put(ctx, entity.IconRegistration{
		Reference: "builtin:unsafe",
		SVG:       []byte(`<svg xmlns="http://www.w3.org/2000/svg"><script>alert(1)</script></svg>`),
	}); err == nil {
		t.Fatal("unsafe icon was persisted")
	}
	listed, err := icons.List(ctx, "builtin", 10)
	if err != nil || len(listed) != 1 {
		t.Fatalf("unsafe registration changed the registry: %#v, %v", listed, err)
	}
}

func TestBuiltinIconCatalogSeedsThroughRustValidation(t *testing.T) {
	registry := iconrepository.NewRegistryRepository(filepath.Join(t.TempDir(), "xaligo-assets.db"))
	t.Cleanup(func() { _ = registry.Close() })
	icons := v2.NewIconUsecase(registry, v2.NewEngineUsecase(), builtin.IconRegistrations()...)
	ctx := context.Background()

	listed, err := icons.List(ctx, "builtin", 100)
	if err != nil {
		t.Fatal(err)
	}
	if len(listed) != 13 {
		t.Fatalf("builtin icon count = %d, want 13", len(listed))
	}
	if listed[0].Ref.String() != "builtin:application" || listed[len(listed)-1].Ref.String() != "builtin:user" {
		t.Fatalf("builtin icon order = %s ... %s", listed[0].Ref.String(), listed[len(listed)-1].Ref.String())
	}
	terminal, err := icons.Get(ctx, "builtin:cli")
	if err != nil {
		t.Fatal(err)
	}
	if terminal.Ref.String() != "builtin:terminal" || terminal.ViewBox != "0 0 24 24" || terminal.Checksum == [32]byte{} {
		t.Fatalf("terminal icon = %#v", terminal)
	}
	results, err := icons.Search(ctx, "command", 10)
	if err != nil || len(results) != 1 || results[0].Ref != terminal.Ref {
		t.Fatalf("command search = %#v, %v", results, err)
	}
	namespaces, err := icons.ListNamespaces(ctx)
	if err != nil || len(namespaces) != 1 || namespaces[0] != "builtin" {
		t.Fatalf("builtin namespaces = %#v, %v", namespaces, err)
	}
}
