package controller

import "testing"

func TestNormalizeRenderFormat(t *testing.T) {
	tests := map[string]string{
		"":           "excalidraw",
		" SVG ":      "svg",
		"Excalidraw": "excalidraw",
		"PPTX":       "pptx",
	}
	for input, want := range tests {
		if got := normalizeRenderFormat(input); got != want {
			t.Errorf("normalizeRenderFormat(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestDefaultRenderOutput(t *testing.T) {
	tests := map[string]string{
		"":           "output.excalidraw",
		"excalidraw": "output.excalidraw",
		"svg":        "output.svg",
		"pptx":       "output.pptx",
	}
	for format, want := range tests {
		if got := defaultRenderOutput(format); got != want {
			t.Errorf("defaultRenderOutput(%q) = %q, want %q", format, got, want)
		}
	}
}
