package repository_test

import (
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
)

func terminalDocument() entity.EngineResolvedDocument {
	visible := entity.EngineResolvedVisual{Visible: true, Shape: entity.EngineShapeRectangle, Stroke: "#2563eb", Fill: "#ffffff", Opacity: 1}
	return entity.EngineResolvedDocument{Width: 400, Height: 200, Elements: []entity.EngineResolvedElement{
		{ID: "system", Concept: entity.EngineConceptGroup, X: 10, Y: 10, Width: 380, Height: 180, Visual: visible, Text: entity.EngineResolvedText{Value: "System"}},
		{ID: "client", ParentID: "system", Concept: entity.EngineConceptItem, X: 40, Y: 70, Width: 100, Height: 60, Visual: visible, Text: entity.EngineResolvedText{Value: "Client"}, IconRef: "catalog:1"},
		{ID: "api", ParentID: "system", Concept: entity.EngineConceptItem, X: 260, Y: 70, Width: 100, Height: 60, Visual: visible, Text: entity.EngineResolvedText{Value: "API 日本語"}},
		{ID: "request", ParentID: "system", Concept: entity.EngineConceptLine, Visual: visible,
			Line:   entity.EngineResolvedLine{TargetDecoration: entity.EngineDecorationArrow, Label: "request"},
			Points: []entity.EnginePoint{{X: 140, Y: 100}, {X: 260, Y: 100}}},
	}}
}

func TestTerminalRepositoryRendersAllLayouts(t *testing.T) {
	renderer := repository.NewTerminalRepository()
	for _, layout := range []entity.TerminalLayout{entity.TerminalLayoutDiagram, entity.TerminalLayoutSemantic, entity.TerminalLayoutHybrid} {
		t.Run(string(layout), func(t *testing.T) {
			output, err := renderer.Render(terminalDocument(), entity.RenderOptions{
				TerminalLayout: layout, TerminalWidth: 60, TerminalHeight: 20,
				TerminalStyle: entity.TerminalStyleUnicode, TerminalFocus: "api",
			})
			if err != nil {
				t.Fatal(err)
			}
			text := string(output)
			if !strings.Contains(text, "API") {
				t.Fatalf("terminal %s output misses API: %q", layout, text)
			}
			if layout == entity.TerminalLayoutHybrid && !strings.Contains(text, "Selected: API") {
				t.Fatalf("hybrid output misses selected detail: %q", text)
			}
		})
	}
}

func TestTerminalRepositoryASCIIContainsOnlyASCII(t *testing.T) {
	output, err := repository.NewTerminalRepository().Render(terminalDocument(), entity.RenderOptions{
		TerminalLayout: entity.TerminalLayoutHybrid, TerminalStyle: entity.TerminalStyleASCII,
		TerminalIcons: entity.TerminalIconsSymbol, TerminalWidth: 60, TerminalHeight: 20,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, value := range output {
		if value > 127 {
			t.Fatalf("ASCII output contains non-ASCII rune %q: %q", value, output)
		}
	}
}

func TestTerminalRepositoryANSIColorIsExplicit(t *testing.T) {
	renderer := repository.NewTerminalRepository()
	plain, err := renderer.Render(terminalDocument(), entity.RenderOptions{TerminalColor: entity.TerminalColorNever})
	if err != nil {
		t.Fatal(err)
	}
	colored, err := renderer.Render(terminalDocument(), entity.RenderOptions{TerminalColor: entity.TerminalColorAlways})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(plain), "\x1b[") || !strings.Contains(string(colored), "\x1b[") {
		t.Fatalf("ANSI policy not honored: plain=%q colored=%q", plain, colored)
	}
}
