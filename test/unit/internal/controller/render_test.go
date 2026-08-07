package controller_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
)

func TestRenderCommandDefaults(t *testing.T) {
	cmd := newRenderController(&fakeUseCase{}).Command()
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		t.Fatal(err)
	}
	if format != "svg" {
		t.Fatalf("format default = %q, want svg", format)
	}
	theme, err := cmd.Flags().GetString("theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "light" {
		t.Fatalf("theme default = %q, want light", theme)
	}
}

func TestRunRenderTerminalWritesStdoutAndForwardsOptions(t *testing.T) {
	dir := t.TempDir()
	input := filepath.Join(dir, "diagram.xal")
	if err := os.WriteFile(input, []byte(`<scene version="2"><item id="api">API</item></scene>`), 0o644); err != nil {
		t.Fatal(err)
	}
	uc := &fakeUseCase{}
	var output bytes.Buffer
	err := newRenderController(uc).RunFormat(entity.ControllerRenderOptions{
		InputPath: input, OutputPath: "-", Format: "terminal", Theme: "light",
		TerminalStyle: "ascii", TerminalLayout: "hybrid", TerminalDetail: "full",
		TerminalColor: "never", TerminalIcons: "symbol", TerminalWidth: 88,
		TerminalHeight: 30, TerminalFocus: "api", Stdout: &output,
	})
	if err != nil {
		t.Fatal(err)
	}
	if output.String() != "terminal\n" {
		t.Fatalf("stdout = %q", output.String())
	}
	if uc.lastRenderOpts.TerminalLayout != entity.TerminalLayoutHybrid || uc.lastRenderOpts.TerminalWidth != 88 || uc.lastRenderOpts.TerminalFocus != "api" {
		t.Fatalf("terminal options = %#v", uc.lastRenderOpts)
	}
}

func TestRunRenderFormatRejectsUnknownFormat(t *testing.T) {
	err := newRenderController(&fakeUseCase{}).RunFormat(entity.ControllerRenderOptions{Format: "unknown", Theme: "light"})
	if err == nil || !strings.Contains(err.Error(), "unknown render format") {
		t.Fatalf("RunRenderFormat() error = %v, want unknown format", err)
	}
}
