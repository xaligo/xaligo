package controller_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/controller"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

type sequencedMarkdownUseCase struct {
	fakeUseCase
	results [][]entity.RenderArtifact
	errors  []error
	calls   int
}

func (rcvr *sequencedMarkdownUseCase) RenderArtifacts(_ context.Context, _ []byte, opts entity.RenderOptions) ([]entity.RenderArtifact, error) {
	rcvr.calls++
	rcvr.lastRenderOpts = opts
	index := rcvr.calls - 1
	if index < len(rcvr.errors) && rcvr.errors[index] != nil {
		return nil, rcvr.errors[index]
	}
	if index < len(rcvr.results) {
		return rcvr.results[index], nil
	}
	return []entity.RenderArtifact{{ID: "frame", Data: []byte(`<svg></svg>`)}}, nil
}

func TestRenderMarkdownCommandDefaults(t *testing.T) {
	cmd := newRenderController(&fakeUseCase{}).Command()
	markdownCmd, _, err := cmd.Find([]string{"markdown"})
	if err != nil {
		t.Fatal(err)
	}
	theme, err := markdownCmd.Flags().GetString("theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "light" {
		t.Fatalf("theme default = %q, want light", theme)
	}
	mode, err := markdownCmd.Flags().GetString("mode")
	if err != nil {
		t.Fatal(err)
	}
	if mode != "standard" {
		t.Fatalf("mode default = %q, want standard", mode)
	}
}

func TestRunMarkdownEmbedsSingleFrameSVG(t *testing.T) {
	dir := t.TempDir()
	input := "# Title\n\nSome text.\n\n```xal\n<frame></frame>\n```\n\nAfter.\n"
	inputPath := filepath.Join(dir, "guide.md")
	if err := os.WriteFile(inputPath, []byte(input), 0o644); err != nil {
		t.Fatal(err)
	}

	uc := &fakeUseCase{renderSVG: []byte(`<svg>ok</svg>`)}
	if err := newRenderController(uc).RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath, Theme: "light"}); err != nil {
		t.Fatalf("RunMarkdown() error = %v", err)
	}

	outputPath := filepath.Join(dir, "guide.embedded.md")
	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if !strings.Contains(string(output), "![](<guide-1.svg>)") {
		t.Fatalf("output missing image reference: %s", output)
	}
	if strings.Contains(string(output), "```xal") {
		t.Fatalf("output still contains xal code fence: %s", output)
	}
	if !strings.Contains(string(output), "Some text.") || !strings.Contains(string(output), "After.") {
		t.Fatalf("output lost surrounding content: %s", output)
	}

	svg, err := os.ReadFile(filepath.Join(dir, "guide-1.svg"))
	if err != nil {
		t.Fatalf("read svg: %v", err)
	}
	if string(svg) != "<svg>ok</svg>" {
		t.Fatalf("svg content = %q", svg)
	}
}

func TestRunMarkdownForwardsPaperAndOrientation(t *testing.T) {
	dir := t.TempDir()
	input := "```xal\n<frame></frame>\n```\n"
	inputPath := filepath.Join(dir, "guide.md")
	if err := os.WriteFile(inputPath, []byte(input), 0o644); err != nil {
		t.Fatal(err)
	}

	uc := &fakeUseCase{renderSVG: []byte(`<svg>ok</svg>`)}
	if err := newRenderController(uc).RunMarkdown(entity.ControllerRenderMarkdownOptions{
		InputPath: inputPath, Paper: "A4", Orientation: "landscape",
	}); err != nil {
		t.Fatalf("RunMarkdown() error = %v", err)
	}
	if uc.lastRenderOpts.PaperSize != "A4" || uc.lastRenderOpts.Orientation != "landscape" {
		t.Fatalf("lastRenderOpts = %#v, want PaperSize=A4 Orientation=landscape", uc.lastRenderOpts)
	}
}

func TestRunMarkdownEmbedsMultiFrameSVGs(t *testing.T) {
	dir := t.TempDir()
	input := "```xal\n<frames></frames>\n```\n"
	inputPath := filepath.Join(dir, "multi.md")
	if err := os.WriteFile(inputPath, []byte(input), 0o644); err != nil {
		t.Fatal(err)
	}

	uc := &fakeUseCase{renderArtifacts: []entity.RenderArtifact{
		{ID: "first", Data: []byte(`<svg>1</svg>`)},
		{ID: "second", Data: []byte(`<svg>2</svg>`)},
	}}
	if err := newRenderController(uc).RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath}); err != nil {
		t.Fatalf("RunMarkdown() error = %v", err)
	}

	output, err := os.ReadFile(filepath.Join(dir, "multi.embedded.md"))
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if !strings.Contains(string(output), "![](<multi-1-first.svg>)") || !strings.Contains(string(output), "![](<multi-1-second.svg>)") {
		t.Fatalf("output missing multi-frame image references: %s", output)
	}
	for _, name := range []string{"multi-1-first.svg", "multi-1-second.svg"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Fatalf("expected svg file %s: %v", name, err)
		}
	}
}

func TestRunMarkdownCustomOutputAndSVGDir(t *testing.T) {
	dir := t.TempDir()
	svgDir := filepath.Join(dir, "images")
	input := "```xal\n<frame></frame>\n```\n"
	inputPath := filepath.Join(dir, "doc.md")
	if err := os.WriteFile(inputPath, []byte(input), 0o644); err != nil {
		t.Fatal(err)
	}

	outputPath := filepath.Join(dir, "out.md")
	uc := &fakeUseCase{}
	if err := newRenderController(uc).RunMarkdown(entity.ControllerRenderMarkdownOptions{
		InputPath: inputPath, OutputPath: outputPath, SVGDir: svgDir,
	}); err != nil {
		t.Fatalf("RunMarkdown() error = %v", err)
	}

	output, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if !strings.Contains(string(output), "![](<images/doc-1.svg>)") {
		t.Fatalf("output missing relative image reference: %s", output)
	}
	if _, err := os.Stat(filepath.Join(svgDir, "doc-1.svg")); err != nil {
		t.Fatalf("expected svg in custom dir: %v", err)
	}
}

func TestRunMarkdownRejectsUnterminatedFence(t *testing.T) {
	dir := t.TempDir()
	input := "```xal\n<frame></frame>\n"
	inputPath := filepath.Join(dir, "broken.md")
	if err := os.WriteFile(inputPath, []byte(input), 0o644); err != nil {
		t.Fatal(err)
	}

	err := newRenderController(&fakeUseCase{}).RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath})
	if err == nil || !strings.Contains(err.Error(), "unterminated") {
		t.Fatalf("RunMarkdown() error = %v, want unterminated fence error", err)
	}
}

func TestRunMarkdownIgnoresNonXalFences(t *testing.T) {
	dir := t.TempDir()
	input := "```go\nfmt.Println(\"hi\")\n```\n"
	inputPath := filepath.Join(dir, "code.md")
	if err := os.WriteFile(inputPath, []byte(input), 0o644); err != nil {
		t.Fatal(err)
	}

	uc := &fakeUseCase{}
	if err := newRenderController(uc).RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath}); err != nil {
		t.Fatalf("RunMarkdown() error = %v", err)
	}
	if uc.renderCalls != 0 {
		t.Fatalf("renderCalls = %d, want 0 for non-xal fence", uc.renderCalls)
	}
	output, err := os.ReadFile(filepath.Join(dir, "code.embedded.md"))
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if string(output) != input {
		t.Fatalf("output = %q, want unchanged %q", output, input)
	}
}

func TestRunMarkdownEscapesGeneratedImageDestination(t *testing.T) {
	dir := t.TempDir()
	inputPath := filepath.Join(dir, "My Guide (draft).md")
	if err := os.WriteFile(inputPath, []byte("```xal\n<frame></frame>\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := newRenderController(&fakeUseCase{}).RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath}); err != nil {
		t.Fatalf("RunMarkdown() error = %v", err)
	}
	output, err := os.ReadFile(filepath.Join(dir, "My Guide (draft).embedded.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(output), "![](<My%20Guide%20%28draft%29-1.svg>)") {
		t.Fatalf("output contains an unsafe image destination: %s", output)
	}
}

func TestRunMarkdownLateRenderFailurePreservesExistingOutputSet(t *testing.T) {
	dir := t.TempDir()
	inputPath := filepath.Join(dir, "guide.md")
	input := "```xal\n<frame id=\"first\"></frame>\n```\n\n```xal\n<frame id=\"second\"></frame>\n```\n"
	if err := os.WriteFile(inputPath, []byte(input), 0o644); err != nil {
		t.Fatal(err)
	}
	markdownPath := filepath.Join(dir, "guide.embedded.md")
	svgPath := filepath.Join(dir, "guide-1.svg")
	if err := os.WriteFile(markdownPath, []byte("old markdown"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(svgPath, []byte("old svg"), 0o644); err != nil {
		t.Fatal(err)
	}

	uc := &sequencedMarkdownUseCase{
		results: [][]entity.RenderArtifact{
			{{ID: "first", Data: []byte("new svg")}},
		},
		errors: []error{nil, errors.New("second block failed")},
	}
	markdownController := controller.NewRenderController(
		config.New(), uc, uc, uc, usecase.NewThemeUsecase(), usecase.NewElementUsecase(),
	)
	err := markdownController.RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath})
	if err == nil || !strings.Contains(err.Error(), "second block failed") {
		t.Fatalf("RunMarkdown() error = %v, want second block failure", err)
	}
	for path, want := range map[string]string{
		markdownPath: "old markdown",
		svgPath:      "old svg",
	} {
		got, readErr := os.ReadFile(path)
		if readErr != nil {
			t.Fatalf("read preserved output %s: %v", path, readErr)
		}
		if string(got) != want {
			t.Fatalf("output %s = %q, want %q", path, got, want)
		}
	}
	if _, statErr := os.Stat(filepath.Join(dir, "guide-2.svg")); !os.IsNotExist(statErr) {
		t.Fatalf("guide-2.svg should not exist after a failed render, stat error = %v", statErr)
	}
}

func TestRunMarkdownReportsFailedRollbackAndPreservesBackup(t *testing.T) {
	dir := t.TempDir()
	inputPath := filepath.Join(dir, "guide.md")
	if err := os.WriteFile(inputPath, []byte("```xal\n<frame></frame>\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	svgPath := filepath.Join(dir, "guide-1.svg")
	markdownPath := filepath.Join(dir, "guide.embedded.md")
	if err := os.WriteFile(svgPath, []byte("old svg"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(markdownPath, []byte("old markdown"), 0o644); err != nil {
		t.Fatal(err)
	}

	var svgBackup string
	remove := func(path string) error {
		if path == svgPath && svgBackup != "" {
			return errors.New("remove installed svg denied")
		}
		return os.Remove(path)
	}
	rename := func(oldPath string, newPath string) error {
		switch {
		case oldPath == svgPath && strings.Contains(filepath.Base(newPath), ".xaligo-markdown-backup-"):
			svgBackup = newPath
		case newPath == markdownPath &&
			strings.Contains(filepath.Base(oldPath), ".xaligo-markdown-") &&
			!strings.Contains(filepath.Base(oldPath), ".xaligo-markdown-backup-"):
			return errors.New("install markdown denied")
		case oldPath == svgBackup && newPath == svgPath:
			return errors.New("restore svg denied")
		}
		return os.Rename(oldPath, newPath)
	}
	uc := &fakeUseCase{renderSVG: []byte("new svg")}
	markdownController := controller.NewRenderController(
		config.New(), uc, uc, uc, usecase.NewThemeUsecase(), usecase.NewElementUsecase(),
		controller.WithRenderMarkdownFileOperations(controller.RenderMarkdownFileOperations{
			Remove: remove,
			Rename: rename,
		}),
	)

	err := markdownController.RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath})
	if err == nil {
		t.Fatal("RunMarkdown() error = nil, want install and rollback errors")
	}
	for _, text := range []string{
		"install markdown denied",
		"rollback Markdown output set",
		"remove installed svg denied",
		"restore existing output " + svgPath,
		"previous output is preserved at " + svgBackup,
	} {
		if !strings.Contains(err.Error(), text) {
			t.Fatalf("RunMarkdown() error = %q, want %q", err, text)
		}
	}
	if svgBackup == "" {
		t.Fatal("SVG backup path was not captured")
	}
	backup, readErr := os.ReadFile(svgBackup)
	if readErr != nil {
		t.Fatalf("read preserved SVG backup: %v", readErr)
	}
	if string(backup) != "old svg" {
		t.Fatalf("SVG backup = %q, want old svg", backup)
	}
	svg, readErr := os.ReadFile(svgPath)
	if readErr != nil {
		t.Fatalf("read installed SVG left by failed removal: %v", readErr)
	}
	if string(svg) != "new svg" {
		t.Fatalf("installed SVG = %q, want new svg", svg)
	}
	markdown, readErr := os.ReadFile(markdownPath)
	if readErr != nil {
		t.Fatalf("read restored Markdown output: %v", readErr)
	}
	if string(markdown) != "old markdown" {
		t.Fatalf("Markdown output = %q, want old markdown", markdown)
	}
	artifacts, globErr := filepath.Glob(filepath.Join(dir, ".xaligo-markdown-*"))
	if globErr != nil {
		t.Fatal(globErr)
	}
	if len(artifacts) != 1 || artifacts[0] != svgBackup {
		t.Fatalf("transaction artifacts = %v, want only preserved backup %s", artifacts, svgBackup)
	}
}

func TestRunMarkdownRejectsImageReferenceAcrossFilesystemVolumes(t *testing.T) {
	dir := t.TempDir()
	inputPath := filepath.Join(dir, "guide.md")
	if err := os.WriteFile(inputPath, []byte("```xal\n<frame></frame>\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	uc := &fakeUseCase{renderSVG: []byte("new svg")}
	markdownController := controller.NewRenderController(
		config.New(), uc, uc, uc, usecase.NewThemeUsecase(), usecase.NewElementUsecase(),
		controller.WithRenderMarkdownFileOperations(controller.RenderMarkdownFileOperations{
			RelativePath: func(string, string) (string, error) {
				return "", errors.New(`Rel: can't make D:\diagrams relative to C:\docs`)
			},
		}),
	)

	err := markdownController.RunMarkdown(entity.ControllerRenderMarkdownOptions{InputPath: inputPath})
	if err == nil || !strings.Contains(err.Error(), "same filesystem volume") {
		t.Fatalf("RunMarkdown() error = %v, want cross-volume guidance", err)
	}
	if !strings.Contains(err.Error(), "make SVG output") {
		t.Fatalf("RunMarkdown() error = %v, want relative-path context", err)
	}
	for _, path := range []string{
		filepath.Join(dir, "guide-1.svg"),
		filepath.Join(dir, "guide.embedded.md"),
	} {
		if _, statErr := os.Stat(path); !os.IsNotExist(statErr) {
			t.Fatalf("output %s must not be written after cross-volume failure, stat error = %v", path, statErr)
		}
	}
}
