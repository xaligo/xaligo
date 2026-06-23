package entity_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/entity"
)

func TestNewSceneDefaults(t *testing.T) {
	scene := entity.NewScene()
	if scene.Type != "excalidraw" || scene.Version != 2 || scene.Source == "" {
		t.Fatalf("scene metadata = %#v", scene)
	}
	if scene.Elements == nil || scene.AppState == nil || scene.Files == nil {
		t.Fatalf("scene collections = %#v", scene)
	}
	if scene.AppState["viewBackgroundColor"] != "#ffffff" {
		t.Fatalf("app state = %#v", scene.AppState)
	}
}

func TestServiceShortLabels(t *testing.T) {
	if got := entity.ShortLabel(entity.ServiceEntry{OfficialName: "Amazon Simple Storage Service"}); got != "S3" {
		t.Fatalf("ShortLabel(S3) = %q", got)
	}
	if got := entity.ShortLabel(entity.ServiceEntry{OfficialName: "Amazon DynamoDB", Abbreviation: "DDBX"}); got != "DDBX" {
		t.Fatalf("ShortLabel override = %q", got)
	}
	if got := entity.ItemShortName("AWS Custom Service"); got != "Custom Service" {
		t.Fatalf("ItemShortName custom = %q", got)
	}
}

func TestNormalizeTheme(t *testing.T) {
	if got, err := entity.NormalizeTheme(" "); err != nil || got != entity.ThemeLight {
		t.Fatalf("NormalizeTheme default = %q, %v", got, err)
	}
	if got, err := entity.NormalizeTheme(" DARK "); err != nil || got != entity.ThemeDark {
		t.Fatalf("NormalizeTheme dark = %q, %v", got, err)
	}
	if _, err := entity.NormalizeTheme("blue"); err == nil || !strings.Contains(err.Error(), "unknown theme") {
		t.Fatalf("NormalizeTheme invalid err = %v", err)
	}
}

func TestDiagnosticsErrorMessages(t *testing.T) {
	if got := (&entity.DiagnosticsError{}).Error(); got != "validation failed" {
		t.Fatalf("empty diagnostics error = %q", got)
	}
	lineErr := (&entity.DiagnosticsError{Diagnostics: []entity.Diagnostic{{Line: 2, Column: 4, Message: "bad tag"}}}).Error()
	if !strings.Contains(lineErr, "line 2, column 4") || !strings.Contains(lineErr, "bad tag") {
		t.Fatalf("line diagnostics error = %q", lineErr)
	}
	plainErr := (&entity.DiagnosticsError{Diagnostics: []entity.Diagnostic{{Message: "bad layout"}}}).Error()
	if plainErr != "bad layout" {
		t.Fatalf("plain diagnostics error = %q", plainErr)
	}
}

func TestNodeAttrAndParseErrorUnwrap(t *testing.T) {
	if got := (*entity.Node)(nil).Attr("missing"); got != "" {
		t.Fatalf("nil Attr = %q", got)
	}
	node := &entity.Node{Attrs: map[string]string{"title": "Dashboard"}}
	if got := node.Attr("title"); got != "Dashboard" {
		t.Fatalf("Attr(title) = %q", got)
	}
	inner := errors.New("bad syntax")
	parseErr := &entity.ParseError{Position: entity.Position{Line: 3, Column: 5}, Err: inner}
	if !errors.Is(parseErr, inner) || !strings.Contains(parseErr.Error(), "line 3, column 5") {
		t.Fatalf("parse error = %v", parseErr)
	}
}
