package controller_test

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/controller"
	"github.com/xaligo/xaligo/internal/entity"
)

type fakeDiffUsecase struct {
	result entity.DiffResult
	err    error
	opts   entity.DiffOptions
}

func (rcvr *fakeDiffUsecase) Diff(_ context.Context, _, _ []byte, opts entity.DiffOptions) (entity.DiffResult, error) {
	rcvr.opts = opts
	return rcvr.result, rcvr.err
}

func TestDiffControllerWritesBothImages(t *testing.T) {
	directory := t.TempDir()
	beforePath := filepath.Join(directory, "before.xal")
	afterPath := filepath.Join(directory, "after.xal")
	for _, path := range []string{beforePath, afterPath} {
		if err := os.WriteFile(path, []byte(`<frame version="1" />`), 0644); err != nil {
			t.Fatal(err)
		}
	}
	fake := &fakeDiffUsecase{result: entity.DiffResult{
		RemovedImage: []byte("removed-svg"), AddedImage: []byte("added-svg"),
		Summary: entity.StructuralDiff{AddedCount: 2, RemovedCount: 1, ModifiedCount: 3},
	}}
	var stdout bytes.Buffer
	prefix := filepath.Join(directory, "nested", "architecture.svg")
	err := controller.NewDiffController(fake).Run(entity.ControllerDiffOptions{
		BeforePath: beforePath, AfterPath: afterPath, OutputPrefix: prefix,
		Theme: "dark", Mode: "network", PxPerInch: 144, Stdout: &stdout,
	})
	if err != nil {
		t.Fatal(err)
	}
	removedPath := filepath.Join(directory, "nested", "architecture-removed.svg")
	addedPath := filepath.Join(directory, "nested", "architecture-added.svg")
	for path, want := range map[string]string{removedPath: "removed-svg", addedPath: "added-svg"} {
		data, err := os.ReadFile(path)
		if err != nil || string(data) != want {
			t.Fatalf("read %s = %q, %v", path, data, err)
		}
	}
	if fake.opts.Theme != "dark" || fake.opts.Mode != "network" || fake.opts.PxPerInch != 144 {
		t.Fatalf("diff options = %#v", fake.opts)
	}
	if !strings.Contains(stdout.String(), "changes: +2 -1 ~3") {
		t.Fatalf("stdout = %q", stdout.String())
	}
}

func TestDiffControllerDoesNotWriteWhenUsecaseFails(t *testing.T) {
	directory := t.TempDir()
	beforePath := filepath.Join(directory, "before.xal")
	afterPath := filepath.Join(directory, "after.xal")
	for _, path := range []string{beforePath, afterPath} {
		if err := os.WriteFile(path, []byte(`<frame version="1" />`), 0644); err != nil {
			t.Fatal(err)
		}
	}
	prefix := filepath.Join(directory, "failed")
	err := controller.NewDiffController(&fakeDiffUsecase{err: errors.New("diff failed")}).Run(entity.ControllerDiffOptions{
		BeforePath: beforePath, AfterPath: afterPath, OutputPrefix: prefix,
	})
	if err == nil || !strings.Contains(err.Error(), "diff failed") {
		t.Fatalf("Run() error = %v", err)
	}
	for _, path := range []string{prefix + "-removed.svg", prefix + "-added.svg"} {
		if _, statErr := os.Stat(path); !os.IsNotExist(statErr) {
			t.Fatalf("unexpected output %s: %v", path, statErr)
		}
	}
}

func TestDiffCommandRequiresTwoInputs(t *testing.T) {
	command := controller.NewDiffController(&fakeDiffUsecase{}).Command()
	if err := command.Args(command, []string{"before.xal"}); err == nil {
		t.Fatal("one-argument validation error = nil")
	}
	if err := command.Args(command, []string{"before.xal", "after.xal"}); err != nil {
		t.Fatal(err)
	}
}
