package usecase_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/usecase"
)

func TestEmbedXalCodeBlocksReplacesFencedBlocks(t *testing.T) {
	source := "# Title\n\nIntro.\n\n```xal\n<frame></frame>\n```\n\nOutro.\n"
	var seen []string
	result, err := usecase.EmbedXalCodeBlocks(source, func(xal string) ([]string, error) {
		seen = append(seen, xal)
		return []string{"<svg>rendered</svg>"}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(seen) != 1 || seen[0] != "<frame></frame>" {
		t.Fatalf("seen = %#v", seen)
	}
	if !strings.Contains(result, "<svg>rendered</svg>") || !strings.Contains(result, "Intro.") || !strings.Contains(result, "Outro.") {
		t.Fatalf("result = %q", result)
	}
	if strings.Contains(result, "```xal") {
		t.Fatalf("result still contains the original fence: %q", result)
	}
}

func TestEmbedXalCodeBlocksIgnoresOtherFences(t *testing.T) {
	source := "```go\nfmt.Println(1)\n```\n"
	called := false
	result, err := usecase.EmbedXalCodeBlocks(source, func(string) ([]string, error) {
		called = true
		return nil, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if called {
		t.Fatal("renderBlock should not be called for a non-xal fence")
	}
	if result != source {
		t.Fatalf("result = %q, want unchanged source", result)
	}
}

func TestEmbedXalCodeBlocksUnterminatedFenceErrors(t *testing.T) {
	source := "```xal\n<frame></frame>\n"
	_, err := usecase.EmbedXalCodeBlocks(source, func(xal string) ([]string, error) {
		return []string{xal}, nil
	})
	if err == nil || !strings.Contains(err.Error(), "unterminated") {
		t.Fatalf("err = %v", err)
	}
}

func TestEmbedXalCodeBlocksPropagatesRenderError(t *testing.T) {
	source := "```xal\nbad\n```\n"
	_, err := usecase.EmbedXalCodeBlocks(source, func(string) ([]string, error) {
		return nil, fmt.Errorf("render failed")
	})
	if err == nil || !strings.Contains(err.Error(), "render failed") {
		t.Fatalf("err = %v", err)
	}
}
