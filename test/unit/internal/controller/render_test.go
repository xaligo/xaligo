package controller_test

import (
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/controller"
)

func TestRenderCommandDefaults(t *testing.T) {
	cmd := controller.InitRenderCmd()
	format, err := cmd.Flags().GetString("format")
	if err != nil {
		t.Fatal(err)
	}
	if format != "excalidraw" {
		t.Fatalf("format default = %q, want excalidraw", format)
	}
	theme, err := cmd.Flags().GetString("theme")
	if err != nil {
		t.Fatal(err)
	}
	if theme != "light" {
		t.Fatalf("theme default = %q, want light", theme)
	}
}

func TestRunRenderFormatRejectsUnknownFormat(t *testing.T) {
	err := controller.RunRenderFormat(controller.RenderOptions{Format: "unknown", Theme: "light"})
	if err == nil || !strings.Contains(err.Error(), "unknown render format") {
		t.Fatalf("RunRenderFormat() error = %v, want unknown format", err)
	}
}
