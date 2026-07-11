package engine_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestV1ParseRejectsNonPositiveCatalogIDs(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   string
	}{
		{name: "item zero", source: `<frame><item id="0" /></frame>`, want: "positive integer"},
		{name: "item padded zero", source: `<frame><item id="000" /></frame>`, want: "positive integer"},
		{name: "item above signed 32-bit", source: `<frame><item id="2147483648" /></frame>`, want: "positive integer"},
		{name: "group icon zero", source: `<frame><generic-group id="group" icon-id="0" /></frame>`, want: "positive catalog ID"},
		{name: "group icon padded zero", source: `<frame><generic-group id="group" icon-id="000" /></frame>`, want: "positive catalog ID"},
		{name: "group icon above signed 32-bit", source: `<frame><generic-group id="group" icon-id="2147483648" /></frame>`, want: "positive catalog ID"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(test.source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestV1BuildValidatesPortSide(t *testing.T) {
	tests := []struct {
		name    string
		side    string
		wantErr bool
	}{
		{name: "omitted", side: ""},
		{name: "top", side: "top"},
		{name: "right", side: "right"},
		{name: "bottom", side: "bottom"},
		{name: "left", side: "left"},
		{name: "invalid", side: "front", wantErr: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sideAttr := ""
			if test.side != "" {
				sideAttr = ` side="` + test.side + `"`
			}
			source := `<frame width="240" height="120"><rectangle id="rect"><port id="port"` + sideAttr + ` /></rectangle></frame>`
			err := parseAndBuildV1EngineDSLValidationTest(source)
			if test.wantErr {
				if err == nil || !strings.Contains(err.Error(), "must be top, right, bottom, or left") {
					t.Fatalf("Build() error = %v, want invalid port side", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Build() error = %v", err)
			}
		})
	}
}

func TestV1BuildValidatesLayoutByTag(t *testing.T) {
	valid := []string{
		`<frame width="240" height="120" layout="vertical"><blank /></frame>`,
		`<frame width="240" height="120" layout="horizontal"><blank /></frame>`,
		`<frames layout="vertical"><frame id="one" width="240" height="120"><blank /></frame></frames>`,
		`<frames layout="horizontal"><frame id="one" width="240" height="120"><blank /></frame></frames>`,
		`<frame width="240" height="120"><generic-group id="group" layout="staggered" /></frame>`,
		`<frame width="240" height="120"><custom-group layout="horizontal"><card /><panel /></custom-group></frame>`,
		`<frame width="240" height="120"><custom-group layout="staggered"><card /><panel /></custom-group></frame>`,
	}
	for _, source := range valid {
		if err := parseAndBuildV1EngineDSLValidationTest(source); err != nil {
			t.Errorf("Build(%s) error = %v", source, err)
		}
	}

	invalid := []struct {
		name   string
		source string
		want   string
	}{
		{name: "unknown frame layout", source: `<frame layout="diagonal"><blank /></frame>`, want: "must be vertical or horizontal"},
		{name: "staggered frame", source: `<frame layout="staggered"><blank /></frame>`, want: "only supported on AWS/group tags"},
		{name: "staggered container", source: `<frame><container layout="staggered"><blank /></container></frame>`, want: "only supported on AWS/group tags"},
		{name: "row layout", source: `<frame><row layout="horizontal"><col /></row></frame>`, want: "is not supported on <row>"},
		{name: "unknown group layout", source: `<frame><generic-group id="group" layout="diagonal" /></frame>`, want: "vertical, horizontal, or staggered"},
	}
	for _, test := range invalid {
		t.Run(test.name, func(t *testing.T) {
			err := parseAndBuildV1EngineDSLValidationTest(test.source)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Build() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestV1ParseRejectsNonV1RootVersion(t *testing.T) {
	valid := []string{
		`<frame><blank /></frame>`,
		`<frame version="1"><blank /></frame>`,
		`<frames version="1"><frame id="one"><blank /></frame></frames>`,
	}
	for _, source := range valid {
		if _, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source)); err != nil {
			t.Errorf("Parse(%s) error = %v", source, err)
		}
	}

	invalid := []struct {
		name   string
		source string
		want   string
	}{
		{name: "V2 frame", source: `<frame version="2"><blank /></frame>`, want: `version="2"`},
		{name: "empty version", source: `<frame version=""><blank /></frame>`, want: "omit version"},
		{name: "noncanonical V1 version", source: `<frame version="1.0"><blank /></frame>`, want: "omit version"},
		{name: "V2 scene root", source: `<scene version="2" />`, want: "root tag must be <frame> or <frames>"},
		{name: "versioned child frame", source: `<frames><frame id="one" version="1"><blank /></frame></frames>`, want: "only allowed on the V1 document root"},
		{name: "versioned nested tag", source: `<frame><custom version="2" /></frame>`, want: "only allowed on the V1 document root"},
	}
	for _, test := range invalid {
		t.Run(test.name, func(t *testing.T) {
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(test.source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestV1ParseNormalizesDocumentedAnchorAliases(t *testing.T) {
	tests := []struct {
		alias string
		want  string
	}{
		{alias: "start", want: "right-1"},
		{alias: "near", want: "right-2"},
		{alias: "center", want: "right-3"},
		{alias: "far", want: "right-4"},
		{alias: "end", want: "right-5"},
	}
	for _, test := range tests {
		t.Run(test.alias, func(t *testing.T) {
			source := `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" src-side="right" src-anchor="` + test.alias + `" /></frame>`
			document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err != nil {
				t.Fatal(err)
			}
			if got := document.Root.Children[2].Attr("src-anchor"); got != test.want {
				t.Fatalf("src-anchor = %q, want %q", got, test.want)
			}
		})
	}
}

func TestV1ParseRejectsAnchorPositionZero(t *testing.T) {
	source := `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" src-side="right" src-anchor="0" /></frame>`
	_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err == nil || !strings.Contains(err.Error(), "position must be 1, 2, 3, 4, 5") {
		t.Fatalf("Parse() error = %v, want anchor position error", err)
	}
}

func TestV1ParseValidatesConnectionEnums(t *testing.T) {
	valid := []string{
		`kind="connection"`,
		`kind="route"`,
		`kind="traffic"`,
		`color="#2563eb"`,
		`stroke-style="solid"`,
		`stroke-style="dashed"`,
		`stroke-style="dotted"`,
		`start-arrowhead="none" end-arrowhead="arrow" arrowhead="stealth"`,
		`start-arrowhead="triangle" end-arrowhead="diamond" arrowhead="oval"`,
		`arrowhead-size="s"`,
	}
	for _, attrs := range valid {
		source := `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" ` + attrs + ` /></frame>`
		if _, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source)); err != nil {
			t.Errorf("Parse(%q) error = %v", attrs, err)
		}
	}

	invalid := []struct {
		name  string
		attrs string
		want  string
	}{
		{name: "kind", attrs: `kind="unknown"`, want: "connection, route, traffic"},
		{name: "stroke style", attrs: `stroke-style="double"`, want: "solid, dashed, dotted"},
		{name: "start arrowhead", attrs: `start-arrowhead="dot"`, want: "none, arrow, triangle, stealth, diamond, oval"},
		{name: "end arrowhead", attrs: `end-arrowhead="dot"`, want: "none, arrow, triangle, stealth, diamond, oval"},
		{name: "arrowhead alias", attrs: `arrowhead="dot"`, want: "none, arrow, triangle, stealth, diamond, oval"},
		{name: "medium arrowhead size", attrs: `arrowhead-size="m"`, want: "must be one of s"},
		{name: "large arrowhead size", attrs: `arrowhead-size="l"`, want: "must be one of s"},
		{name: "unknown arrowhead size", attrs: `arrowhead-size="xl"`, want: "must be one of s"},
		{name: "named color", attrs: `color="red"`, want: "six-digit hexadecimal color"},
		{name: "short color", attrs: `color="#abc"`, want: "six-digit hexadecimal color"},
		{name: "missing color hash", attrs: `color="112233"`, want: "six-digit hexadecimal color"},
		{name: "alpha color", attrs: `color="#11223344"`, want: "six-digit hexadecimal color"},
	}
	for _, test := range invalid {
		t.Run(test.name, func(t *testing.T) {
			source := `<frame><item id="1" /><item id="2" /><connection src="1" dst="2" ` + test.attrs + ` /></frame>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}

	grouped := `<frame><item id="1" /><item id="2" /><connections kind="unknown"><connection src="1" dst="2" /></connections></frame>`
	if _, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(grouped)); err == nil || !strings.Contains(err.Error(), "connection, route, traffic") {
		t.Fatalf("Parse(<connections>) error = %v, want inherited kind validation", err)
	}
}

func TestV1ParseRejectsNonHeadlessEffectiveRoutes(t *testing.T) {
	tests := []struct {
		name       string
		connection string
		wantHead   string
		wantLine   int
	}{
		{
			name:       "direct start arrowhead",
			connection: `<connection src="1" dst="2" kind="route" start-arrowhead="oval" />`,
			wantHead:   `start-arrowhead="oval"`,
			wantLine:   4,
		},
		{
			name:       "direct canonical end arrowhead",
			connection: `<connection src="1" dst="2" kind="route" end-arrowhead="triangle" />`,
			wantHead:   `end-arrowhead/arrowhead="triangle"`,
			wantLine:   4,
		},
		{
			name:       "direct end arrowhead alias",
			connection: `<connection src="1" dst="2" kind="route" arrowhead="diamond" />`,
			wantHead:   `end-arrowhead/arrowhead="diamond"`,
			wantLine:   4,
		},
		{
			name: "inherited route and arrowhead",
			connection: `<connections kind="route" end-arrowhead="triangle">
    <connection src="1" dst="2" />
  </connections>`,
			wantHead: `end-arrowhead/arrowhead="triangle"`,
			wantLine: 5,
		},
		{
			name: "child route inherits parent arrowhead",
			connection: `<connections end-arrowhead="triangle">
    <connection src="1" dst="2" kind="route" />
  </connections>`,
			wantHead: `end-arrowhead/arrowhead="triangle"`,
			wantLine: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<frame>
  <item id="1" />
  <item id="2" />
  ` + test.connection + `
</frame>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), `kind="route"> must be headless`) || !strings.Contains(err.Error(), test.wantHead) {
				t.Fatalf("Parse() error = %v, want headless error containing %q", err, test.wantHead)
			}
			var parseErr *entity.ParseError
			if !errors.As(err, &parseErr) || parseErr.Position.Line != test.wantLine || parseErr.Position.Column == 0 {
				t.Fatalf("Parse() positioned error = %#v, want line %d", parseErr, test.wantLine)
			}
		})
	}
}

func TestV1ParseAllowsHeadlessEffectiveRoutesAfterAliasOverrides(t *testing.T) {
	sources := []string{
		`<frame><item id="1" /><item id="2" /><connection src="1" dst="2" kind="route" start-arrowhead="none" end-arrowhead="none" /></frame>`,
		`<frame><item id="1" /><item id="2" /><connection src="1" dst="2" kind="route" end-arrowhead="none" arrowhead="oval" /></frame>`,
		`<frame><item id="1" /><item id="2" /><connections kind="route" end-arrowhead="triangle"><connection src="1" dst="2" arrowhead="none" /></connections></frame>`,
		`<frame><item id="1" /><item id="2" /><connections kind="route" arrowhead="diamond"><connection src="1" dst="2" end-arrowhead="none" /></connections></frame>`,
		`<frame><item id="1" /><item id="2" /><connections kind="route" start-arrowhead="oval"><connection src="1" dst="2" start-arrowhead="none" /></connections></frame>`,
		`<frame><item id="1" /><item id="2" /><connections kind="route" end-arrowhead="triangle"><connection src="1" dst="2" kind="traffic" /></connections></frame>`,
	}

	for _, source := range sources {
		if _, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source)); err != nil {
			t.Errorf("Parse(%q) error = %v", source, err)
		}
	}
}

func TestV1ParseRejectsNonConnectionGroupChild(t *testing.T) {
	source := `<frame><connections><connecton src="one" dst="two" /></connections></frame>`
	_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err == nil || !strings.Contains(err.Error(), "may only contain <connection> children") {
		t.Fatalf("Parse() error = %v, want connection-group child error", err)
	}
}

func parseAndBuildV1EngineDSLValidationTest(source string) error {
	document, err := v1engine.ParseV1EngineParseDocument(bytes.NewBufferString(source))
	if err != nil {
		return err
	}
	_, err = v1engine.BuildV1EngineLayoutBuild(document)
	return err
}
