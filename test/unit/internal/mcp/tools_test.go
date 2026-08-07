package mcp_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/mcp"
)

type fakeDiagnostics struct{}

func (*fakeDiagnostics) Diagnose(context.Context, []byte) ([]entity.Diagnostic, error) {
	return []entity.Diagnostic{{Severity: entity.DiagnosticSeverity("warning"), Message: "legacy root"}}, nil
}

type fakeRender struct {
	options entity.RenderOptions
}

func (rcvr *fakeRender) RenderArtifacts(_ context.Context, _ []byte, options entity.RenderOptions) ([]entity.RenderArtifact, error) {
	rcvr.options = options
	return []entity.RenderArtifact{{ID: "main", Data: []byte("<svg/>")}}, nil
}

type fakeProject struct {
	indexedRoot string
	analyzed    int
}

func (rcvr *fakeProject) Analyze(_ context.Context, uri string, _ []byte) (entity.ProjectAnalysis, error) {
	rcvr.analyzed++
	return entity.ProjectAnalysis{URI: uri, Kind: entity.ProjectDocumentXAL, Symbols: []entity.ProjectSymbol{{ID: "main", Concept: entity.ProjectConceptFrame}}}, nil
}

func (rcvr *fakeProject) Index(_ context.Context, root string) (entity.ProjectIndexStats, error) {
	rcvr.indexedRoot = root
	return entity.ProjectIndexStats{Root: root, Scanned: 2, Indexed: 2}, nil
}

func (*fakeProject) Search(context.Context, string, int) ([]entity.ProjectSearchResult, error) {
	return []entity.ProjectSearchResult{{URI: "docs/guide.md", Name: "Guide", Concept: entity.ProjectConceptText}}, nil
}

func (*fakeProject) Symbols(context.Context, string) ([]entity.ProjectSymbol, error) {
	return []entity.ProjectSymbol{{ID: "api", Name: "API", Concept: entity.ProjectConceptItem}}, nil
}

type fakeIcons struct {
	put     entity.IconRegistration
	deleted string
}

func (rcvr *fakeIcons) Put(_ context.Context, registration entity.IconRegistration) (entity.Icon, error) {
	rcvr.put = registration
	return entity.Icon{Ref: entity.IconRef{Namespace: "custom", Name: "router"}, ViewBox: "0 0 24 24", SVG: []byte("<svg/>")}, nil
}

func (*fakeIcons) Get(context.Context, string) (entity.Icon, error) {
	return entity.Icon{Ref: entity.IconRef{Namespace: "builtin", Name: "database"}, SVG: []byte("<svg/>")}, nil
}

func (*fakeIcons) Search(context.Context, string, int) ([]entity.IconSummary, error) {
	return []entity.IconSummary{{Ref: entity.IconRef{Namespace: "builtin", Name: "database"}}}, nil
}

func (rcvr *fakeIcons) Delete(_ context.Context, reference string) error {
	rcvr.deleted = reference
	return nil
}

func (*fakeIcons) ListNamespaces(context.Context) ([]string, error) {
	return []string{"builtin", "custom"}, nil
}

func TestToolServiceKeepsDocsIndexSeparateFromExplicitXALAnalysis(t *testing.T) {
	project := &fakeProject{}
	service := mcp.NewToolService(&fakeDiagnostics{}, &fakeRender{}, project, &fakeIcons{}, "/project")

	inspect, err := service.Call(context.Background(), "inspect_xal", json.RawMessage(`{"source":"<frame/>"}`))
	if err != nil {
		t.Fatal(err)
	}
	if project.analyzed != 1 || project.indexedRoot != "" {
		t.Fatalf("explicit inspect analyzed=%d indexedRoot=%q", project.analyzed, project.indexedRoot)
	}
	if inspect.StructuredContent.(map[string]any)["uri"] != "memory://mcp/document.xal" {
		t.Fatalf("inspect result = %#v", inspect.StructuredContent)
	}

	index, err := service.Call(context.Background(), "index_docs", json.RawMessage(`{}`))
	if err != nil {
		t.Fatal(err)
	}
	if project.indexedRoot != "/project" || index.StructuredContent.(entity.ProjectIndexStats).Scanned != 2 {
		t.Fatalf("docs index root=%q result=%#v", project.indexedRoot, index.StructuredContent)
	}
}

func TestToolServiceReturnsProjectSymbols(t *testing.T) {
	service := mcp.NewToolService(&fakeDiagnostics{}, &fakeRender{}, &fakeProject{}, &fakeIcons{}, "/project")
	result, err := service.Call(context.Background(), "project_symbols", json.RawMessage(`{"uri":"file:///diagram.xal"}`))
	if err != nil {
		t.Fatal(err)
	}
	content := result.StructuredContent.(map[string]any)
	if content["uri"] != "file:///diagram.xal" || len(content["symbols"].([]entity.ProjectSymbol)) != 1 {
		t.Fatalf("project symbols = %#v", content)
	}
}

func TestToolServiceRendersSVGAndManagesIcons(t *testing.T) {
	render := &fakeRender{}
	icons := &fakeIcons{}
	service := mcp.NewToolService(&fakeDiagnostics{}, render, &fakeProject{}, icons, "/project")

	result, err := service.Call(context.Background(), "render_svg", json.RawMessage(`{"source":"<frame/>","mode":"network","combineFrames":true}`))
	if err != nil {
		t.Fatal(err)
	}
	if render.options.Format != entity.Format("svg") || render.options.Mode != entity.Mode("network") || !render.options.CombineFrames {
		t.Fatalf("render options = %#v", render.options)
	}
	artifacts := result.StructuredContent.(map[string]any)["artifacts"].([]map[string]any)
	if len(artifacts) != 1 || artifacts[0]["svg"] != "<svg/>" {
		t.Fatalf("render result = %#v", result.StructuredContent)
	}

	_, err = service.Call(context.Background(), "register_icon", json.RawMessage(`{"reference":"custom:router","svg":"<svg/>","tags":["network"]}`))
	if err != nil {
		t.Fatal(err)
	}
	if icons.put.Reference != "custom:router" || len(icons.put.Tags) != 1 {
		t.Fatalf("icon registration = %#v", icons.put)
	}
	if _, err := service.Call(context.Background(), "remove_icon", json.RawMessage(`{"reference":"custom:router"}`)); err != nil {
		t.Fatal(err)
	}
	if icons.deleted != "custom:router" {
		t.Fatalf("deleted icon = %q", icons.deleted)
	}
}

func TestToolServiceRejectsUnknownArguments(t *testing.T) {
	service := mcp.NewToolService(&fakeDiagnostics{}, &fakeRender{}, &fakeProject{}, &fakeIcons{}, "/project")
	if _, err := service.Call(context.Background(), "validate_xal", json.RawMessage(`{"source":"<frame/>","extra":true}`)); err == nil {
		t.Fatal("expected strict argument validation")
	}
}
