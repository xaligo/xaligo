//go:build cgo && sqlite_fts5

package icon_test

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"sync"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xaligo/xaligo/internal/entity"
	iconrepository "github.com/xaligo/xaligo/internal/repository/icon"
)

func iconFloat(value float64) *float64 {
	return &value
}

func testIcon() entity.Icon {
	svg := []byte(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M0 0h24v24z"/></svg>`)
	return entity.Icon{
		Ref:         entity.IconRef{Namespace: "builtin", Name: "database"},
		Description: "Relational database storage",
		SVG:         svg,
		ViewBox:     "0 0 24 24",
		Width:       iconFloat(24),
		Height:      iconFloat(24),
		Checksum:    sha256.Sum256(svg),
		License:     "MIT",
		Source:      "test",
		Tags:        []string{"database", "storage"},
		Aliases:     []entity.IconRef{{Namespace: "builtin", Name: "db"}},
	}
}

func TestRegistryCRUDSearchAndFTSReplacement(t *testing.T) {
	ctx := context.Background()
	path := filepath.Join(t.TempDir(), "xaligo-assets.db")
	repository := iconrepository.NewRegistryRepository(path)
	t.Cleanup(func() { _ = repository.Close() })

	stored, err := repository.Put(ctx, testIcon())
	if err != nil {
		t.Fatal(err)
	}
	if stored.ID == 0 || stored.CreatedAt.IsZero() || stored.UpdatedAt.IsZero() {
		t.Fatalf("stored icon metadata = %#v", stored)
	}
	byAlias, err := repository.Get(ctx, entity.IconRef{Namespace: "builtin", Name: "db"})
	if err != nil {
		t.Fatal(err)
	}
	if byAlias.Ref != stored.Ref || string(byAlias.SVG) != string(stored.SVG) || len(byAlias.Tags) != 2 || len(byAlias.Aliases) != 1 {
		t.Fatalf("icon resolved by alias = %#v", byAlias)
	}

	for _, query := range []string{"relational", "storage", "db"} {
		results, err := repository.Search(ctx, query, 10)
		if err != nil || len(results) != 1 || results[0].Ref != stored.Ref {
			t.Fatalf("Search(%q) = %#v, %v", query, results, err)
		}
	}

	updated := testIcon()
	updated.Description = "Primary data service"
	updated.Tags = []string{"primary"}
	updated.Aliases = []entity.IconRef{{Namespace: "builtin", Name: "data"}}
	storedAgain, err := repository.Put(ctx, updated)
	if err != nil {
		t.Fatal(err)
	}
	if storedAgain.ID != stored.ID || !storedAgain.CreatedAt.Equal(stored.CreatedAt) {
		t.Fatalf("update changed stable identity: before=%#v after=%#v", stored, storedAgain)
	}
	stale, err := repository.Search(ctx, "storage", 10)
	if err != nil || len(stale) != 0 {
		t.Fatalf("stale FTS results = %#v, %v", stale, err)
	}
	current, err := repository.Search(ctx, "primary", 10)
	if err != nil || len(current) != 1 {
		t.Fatalf("updated FTS results = %#v, %v", current, err)
	}
	if _, err := repository.Get(ctx, entity.IconRef{Namespace: "builtin", Name: "db"}); !errors.Is(err, iconrepository.ErrNotFound) {
		t.Fatalf("removed alias error = %v", err)
	}

	if _, err := repository.Search(ctx, `database'; DROP TABLE icons; --`, 10); err == nil {
		t.Fatal("malformed FTS query unexpectedly succeeded")
	}
	listed, err := repository.List(ctx, "builtin", 10)
	if err != nil || len(listed) != 1 {
		t.Fatalf("registry changed by SQL metacharacters: %#v, %v", listed, err)
	}
	namespaces, err := repository.ListNamespaces(ctx)
	if err != nil || len(namespaces) != 1 || namespaces[0] != "builtin" {
		t.Fatalf("namespaces = %#v, %v", namespaces, err)
	}

	if err := repository.Delete(ctx, stored.Ref); err != nil {
		t.Fatal(err)
	}
	if _, err := repository.Get(ctx, stored.Ref); !errors.Is(err, iconrepository.ErrNotFound) {
		t.Fatalf("deleted icon error = %v", err)
	}
	results, err := repository.Search(ctx, "primary", 10)
	if err != nil || len(results) != 0 {
		t.Fatalf("deleted FTS results = %#v, %v", results, err)
	}
}

func TestRegistryRejectsFutureSchemaVersion(t *testing.T) {
	path := filepath.Join(t.TempDir(), "future.db")
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`PRAGMA user_version = 999`); err != nil {
		t.Fatal(err)
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	repository := iconrepository.NewRegistryRepository(path)
	t.Cleanup(func() { _ = repository.Close() })
	if _, err := repository.ListNamespaces(context.Background()); err == nil {
		t.Fatal("future schema version was accepted")
	}
}

func TestRegistryRetriesAfterCanceledOpen(t *testing.T) {
	repository := iconrepository.NewRegistryRepository(filepath.Join(t.TempDir(), "retry.db"))
	t.Cleanup(func() { _ = repository.Close() })
	canceled, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := repository.ListNamespaces(canceled); !errors.Is(err, context.Canceled) {
		t.Fatalf("canceled open error = %v", err)
	}
	namespaces, err := repository.ListNamespaces(context.Background())
	if err != nil || len(namespaces) != 0 {
		t.Fatalf("retry namespaces = %#v, %v", namespaces, err)
	}
}

func TestRegistryAllowsConcurrentReadersWithSerializedWrites(t *testing.T) {
	ctx := context.Background()
	repository := iconrepository.NewRegistryRepository(filepath.Join(t.TempDir(), "concurrent.db"))
	t.Cleanup(func() { _ = repository.Close() })
	if _, err := repository.Put(ctx, testIcon()); err != nil {
		t.Fatal(err)
	}

	errorsCh := make(chan error, 32)
	var group sync.WaitGroup
	group.Add(1)
	go func() {
		defer group.Done()
		for index := range 12 {
			icon := testIcon()
			icon.Description = fmt.Sprintf("Database revision %d", index)
			icon.Tags = []string{fmt.Sprintf("revision-%d", index)}
			if _, err := repository.Put(ctx, icon); err != nil {
				errorsCh <- err
				return
			}
		}
	}()
	for range 3 {
		group.Add(1)
		go func() {
			defer group.Done()
			for range 12 {
				if _, err := repository.Get(ctx, entity.IconRef{Namespace: "builtin", Name: "database"}); err != nil {
					errorsCh <- err
					return
				}
				if _, err := repository.List(ctx, "builtin", 10); err != nil {
					errorsCh <- err
					return
				}
			}
		}()
	}
	group.Wait()
	close(errorsCh)
	for err := range errorsCh {
		t.Errorf("concurrent registry operation: %v", err)
	}
}
