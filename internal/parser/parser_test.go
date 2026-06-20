package parser

import (
	"errors"
	"strings"
	"testing"
)

func TestParseStoresNodePositions(t *testing.T) {
	doc, err := Parse(strings.NewReader("<frame>\n  <item id=\"123\" />\n</frame>"))
	if err != nil {
		t.Fatal(err)
	}
	if got := doc.Root.Position; got.Line != 1 || got.Column != 1 || got.Offset != 0 {
		t.Fatalf("root position = %#v", got)
	}
	if got := doc.Root.Children[0].Position; got.Line != 2 || got.Column != 3 {
		t.Fatalf("item position = %#v", got)
	}
}

func TestParseValidationErrorHasPosition(t *testing.T) {
	_, err := Parse(strings.NewReader("<frame>\n  <item id=\"bad\" />\n</frame>"))
	var parseErr *Error
	if !errors.As(err, &parseErr) {
		t.Fatalf("error = %T %v, want *parser.Error", err, err)
	}
	if parseErr.Position.Line != 2 || parseErr.Position.Column != 3 {
		t.Fatalf("error position = %#v", parseErr.Position)
	}
}

func TestParseExpandsConnectionShorthands(t *testing.T) {
	doc, err := Parse(strings.NewReader(`<frame>
  <item id="1" name="web" />
  <item id="2" ref="db" />
  web --- db
  web ==> 2
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	if len(doc.Root.Children) != 4 {
		t.Fatalf("children = %d, want 4", len(doc.Root.Children))
	}
	route := doc.Root.Children[2]
	traffic := doc.Root.Children[3]
	if route.Tag != "connection" || route.Attr("src") != "1" || route.Attr("dst") != "2" || route.Attr("kind") != "route" {
		t.Fatalf("route = %#v", route)
	}
	if traffic.Attr("src") != "1" || traffic.Attr("dst") != "2" || traffic.Attr("kind") != "traffic" {
		t.Fatalf("traffic = %#v", traffic)
	}
	if route.Position.Line != 4 || route.Position.Column != 3 {
		t.Fatalf("route position = %#v", route.Position)
	}
}

func TestParseShorthandReportsUnknownReference(t *testing.T) {
	_, err := Parse(strings.NewReader(`<frame>
  <item id="1" name="web" />
  web --- missing
</frame>`))
	var parseErr *Error
	if !errors.As(err, &parseErr) {
		t.Fatalf("error = %T %v", err, err)
	}
	if parseErr.Position.Line != 3 || parseErr.Position.Column != 3 || !strings.Contains(err.Error(), `destination "missing"`) {
		t.Fatalf("error = %v at %#v", err, parseErr.Position)
	}
}

func TestParseRejectsInvalidGenericGroupIconID(t *testing.T) {
	_, err := Parse(strings.NewReader(`<frame><generic-group icon-id="router" /></frame>`))
	var parseErr *Error
	if !errors.As(err, &parseErr) || !strings.Contains(err.Error(), "positive catalog ID") {
		t.Fatalf("error = %T %v", err, err)
	}
}
