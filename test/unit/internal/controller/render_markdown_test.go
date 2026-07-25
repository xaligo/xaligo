package controller_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
)

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
	if !strings.Contains(string(output), "![](guide-1.svg)") {
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
	if !strings.Contains(string(output), "![](multi-1-first.svg)") || !strings.Contains(string(output), "![](multi-1-second.svg)") {
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
	if !strings.Contains(string(output), "![](images/doc-1.svg)") {
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
