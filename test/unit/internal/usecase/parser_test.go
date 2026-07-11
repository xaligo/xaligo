package usecase_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestParseStoresNodePositions(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader("<frame>\n  <item id=\"123\" />\n</frame>"))
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
	_, err := usecase.Parse(strings.NewReader("<frame>\n  <item id=\"bad\" />\n</frame>"))
	var parseErr *entity.ParseError
	if !errors.As(err, &parseErr) {
		t.Fatalf("error = %T %v, want *entity.ParseError", err, err)
	}
	if parseErr.Position.Line != 2 || parseErr.Position.Column != 3 {
		t.Fatalf("error position = %#v", parseErr.Position)
	}
}

func TestParseExpandsConnectionShorthands(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame>
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
	if route.Tag != "connection" || route.Attr("src") != "web" || route.Attr("dst") != "db" || route.Attr("kind") != "route" {
		t.Fatalf("route = %#v", route)
	}
	if traffic.Attr("src") != "web" || traffic.Attr("dst") != "2" || traffic.Attr("kind") != "traffic" {
		t.Fatalf("traffic = %#v", traffic)
	}
	if route.Position.Line != 4 || route.Position.Column != 3 {
		t.Fatalf("route position = %#v", route.Position)
	}
}

func TestParseResolvesExplicitConnectionReferences(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame>
  <item id="27" name="app-primary" />
  <item id="27" name="app-standby" />
  <item id="110" ref="db" />
  <connection src="app-primary" dst="db" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	conn := doc.Root.Children[3]
	if conn.Attr("src") != "app-primary" || conn.Attr("dst") != "db" {
		t.Fatalf("connection attrs = %#v", conn.Attrs)
	}
	if conn.Attr("_xaligoConnectionSrcKey") == "" || conn.Attr("_xaligoConnectionDstKey") == "" {
		t.Fatalf("connection keys = %#v", conn.Attrs)
	}
}

func TestParseResolvesGroupedConnectionReferences(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame>
  <item id="1" name="web" />
  <item id="2" name="db" />
  <connections grid="8">
    <connection src="web" dst="db" />
  </connections>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	conn := doc.Root.Children[2].Children[0]
	if conn.Attr("_xaligoConnectionSrcKey") == "" || conn.Attr("_xaligoConnectionDstKey") == "" {
		t.Fatalf("connection keys = %#v", conn.Attrs)
	}
}

func TestParseFramesResolvesCrossFrameConnection(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frames>
  <frame id="page-a" width="320" height="180">
    <rectangle id="web" title="Web" />
    <connection src="web" dst="db" />
  </frame>
  <frame id="page-b" width="320" height="180">
    <rectangle id="db" title="DB" />
  </frame>
</frames>`))
	if err != nil {
		t.Fatal(err)
	}
	conn := doc.Root.Children[0].Children[1]
	if conn.Attr("_xaligoConnectionSrcFrame") != "page-a" || conn.Attr("_xaligoConnectionDstFrame") != "page-b" || conn.Attr("_xaligoConnectionCrossFrame") != "true" {
		t.Fatalf("cross-frame attrs = %#v", conn.Attrs)
	}
}

func TestParseFramesRequireFrameID(t *testing.T) {
	_, err := usecase.Parse(strings.NewReader(`<frames><frame><blank /></frame></frames>`))
	var parseErr *entity.ParseError
	if !errors.As(err, &parseErr) || !strings.Contains(err.Error(), "requires a non-empty id") {
		t.Fatalf("error = %T %v", err, err)
	}
}

func TestParseResolvesConnectionEndpointChildTags(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame>
  <item id="1" name="web" />
  <item id="2" name="db" />
  <connection>
    <src side="left" anchor="3">web</src>
    <dst anchor="right-5" ref="db" />
  </connection>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	conn := doc.Root.Children[2]
	if conn.Attr("src") != "web" || conn.Attr("dst") != "db" {
		t.Fatalf("connection endpoint attrs = %#v", conn.Attrs)
	}
	if conn.Attr("src-side") != "left" || conn.Attr("dst-side") != "right" {
		t.Fatalf("connection side attrs = %#v", conn.Attrs)
	}
	if conn.Attr("src-anchor") != "left-3" || conn.Attr("dst-anchor") != "right-5" {
		t.Fatalf("connection anchor attrs = %#v", conn.Attrs)
	}
	if conn.Attr("_xaligoConnectionSrcKey") == "" || conn.Attr("_xaligoConnectionDstKey") == "" {
		t.Fatalf("connection keys = %#v", conn.Attrs)
	}
}

func TestParseResolvesGroupConnectionReferences(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame>
  <generic-group id="edge" name="edge-group" title="Edge">
    <item id="1" />
  </generic-group>
  <item id="2" name="app" />
  <connection src="edge-group" dst="app" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	conn := doc.Root.Children[2]
	if conn.Attr("_xaligoConnectionSrcKey") == "" || conn.Attr("_xaligoConnectionDstKey") == "" {
		t.Fatalf("connection keys = %#v", conn.Attrs)
	}
	if conn.Attr("_xaligoConnectionSrcKey") == conn.Attr("_xaligoConnectionDstKey") {
		t.Fatalf("connection keys should differ = %#v", conn.Attrs)
	}
}

func TestParseShorthandReportsUnknownReference(t *testing.T) {
	_, err := usecase.Parse(strings.NewReader(`<frame>
  <item id="1" name="web" />
  web --- missing
</frame>`))
	var parseErr *entity.ParseError
	if !errors.As(err, &parseErr) {
		t.Fatalf("error = %T %v", err, err)
	}
	if parseErr.Position.Line != 3 || parseErr.Position.Column != 3 || !strings.Contains(err.Error(), `destination "missing"`) {
		t.Fatalf("error = %v at %#v", err, parseErr.Position)
	}
}

func TestParseRejectsInvalidGenericGroupIconID(t *testing.T) {
	_, err := usecase.Parse(strings.NewReader(`<frame><generic-group id="network" icon-id="router" /></frame>`))
	var parseErr *entity.ParseError
	if !errors.As(err, &parseErr) || !strings.Contains(err.Error(), "positive catalog ID") {
		t.Fatalf("error = %T %v", err, err)
	}
}
