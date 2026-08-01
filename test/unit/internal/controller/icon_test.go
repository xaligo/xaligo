package controller_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/controller"
	"github.com/xaligo/xaligo/internal/entity"
)

type fakeIconUsecase struct {
	registration entity.IconRegistration
	searchQuery  string
	putCalls     int
}

func (rcvr *fakeIconUsecase) Put(_ context.Context, registration entity.IconRegistration) (entity.Icon, error) {
	rcvr.putCalls++
	rcvr.registration = registration
	return entity.Icon{
		Ref: entity.IconRef{Namespace: "builtin", Name: "database"},
		SVG: registration.SVG, Checksum: sha256.Sum256(registration.SVG),
	}, nil
}

func TestIconControllerRejectsOversizedSVGBeforeUsecase(t *testing.T) {
	path := filepath.Join(t.TempDir(), "oversized.svg")
	if err := os.WriteFile(path, make([]byte, 2*1024*1024+1), 0o644); err != nil {
		t.Fatal(err)
	}
	fake := &fakeIconUsecase{}
	command := controller.NewIconController(fake).Command()
	command.SetArgs([]string{"add", path, "--name", "builtin:oversized"})
	if err := command.Execute(); err == nil || !strings.Contains(err.Error(), "exceeds 2097152 bytes") {
		t.Fatalf("oversized SVG error = %v", err)
	}
	if fake.putCalls != 0 {
		t.Fatalf("usecase Put calls = %d", fake.putCalls)
	}
}

func (rcvr *fakeIconUsecase) Get(context.Context, string) (entity.Icon, error) {
	return entity.Icon{SVG: []byte(`<svg/>`)}, nil
}

func (rcvr *fakeIconUsecase) Search(_ context.Context, query string, _ int) ([]entity.IconSummary, error) {
	rcvr.searchQuery = query
	return []entity.IconSummary{{Ref: entity.IconRef{Namespace: "builtin", Name: "database"}, Description: "Database"}}, nil
}

func (rcvr *fakeIconUsecase) Delete(context.Context, string) error { return nil }

func (rcvr *fakeIconUsecase) List(context.Context, string, int) ([]entity.IconSummary, error) {
	return []entity.IconSummary{{Ref: entity.IconRef{Namespace: "builtin", Name: "database"}}}, nil
}

func (rcvr *fakeIconUsecase) ListNamespaces(context.Context) ([]string, error) {
	return []string{"builtin"}, nil
}

func TestIconControllerCommandsAndIO(t *testing.T) {
	fake := &fakeIconUsecase{}
	command := controller.NewIconController(fake).Command()
	seen := map[string]bool{}
	for _, child := range command.Commands() {
		seen[child.Name()] = true
	}
	for _, name := range []string{"add", "get", "search", "remove", "list", "namespaces"} {
		if !seen[name] {
			t.Fatalf("icon subcommand %q missing from %#v", name, seen)
		}
	}

	path := filepath.Join(t.TempDir(), "database.svg")
	if err := os.WriteFile(path, []byte(`<svg/>`), 0o644); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	command.SetOut(&output)
	command.SetErr(&output)
	command.SetArgs([]string{"add", path, "--name", "builtin:database", "--tag", "storage", "--alias", "db"})
	if err := command.Execute(); err != nil {
		t.Fatal(err)
	}
	if fake.registration.Reference != "builtin:database" || len(fake.registration.Tags) != 1 || len(fake.registration.Aliases) != 1 || !strings.Contains(output.String(), "builtin:database") {
		t.Fatalf("registration=%#v output=%q", fake.registration, output.String())
	}

	output.Reset()
	command = controller.NewIconController(fake).Command()
	command.SetOut(&output)
	command.SetArgs([]string{"search", "database", "OR", "storage"})
	if err := command.Execute(); err != nil {
		t.Fatal(err)
	}
	if fake.searchQuery != "database OR storage" || !strings.Contains(output.String(), "builtin:database") {
		t.Fatalf("query=%q output=%q", fake.searchQuery, output.String())
	}

	output.Reset()
	command = controller.NewIconController(fake).Command()
	command.SetOut(&output)
	command.SetArgs([]string{"get", "builtin:database"})
	if err := command.Execute(); err != nil {
		t.Fatal(err)
	}
	if output.String() != `<svg/>` {
		t.Fatalf("get output = %q", output.String())
	}
}
