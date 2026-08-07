package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const (
	maxXALSourceBytes = 4 * 1024 * 1024
	maxIconSVGBytes   = 2 * 1024 * 1024
)

type DiagnosticsService interface {
	Diagnose(context.Context, []byte) ([]entity.Diagnostic, error)
}

type RenderService interface {
	RenderArtifacts(context.Context, []byte, entity.RenderOptions) ([]entity.RenderArtifact, error)
}

type ProjectService interface {
	Analyze(context.Context, string, []byte) (entity.ProjectAnalysis, error)
	Index(context.Context, string) (entity.ProjectIndexStats, error)
	Search(context.Context, string, int) ([]entity.ProjectSearchResult, error)
	Symbols(context.Context, string) ([]entity.ProjectSymbol, error)
}

type IconService interface {
	Put(context.Context, entity.IconRegistration) (entity.Icon, error)
	Get(context.Context, string) (entity.Icon, error)
	Search(context.Context, string, int) ([]entity.IconSummary, error)
	Delete(context.Context, string) error
	ListNamespaces(context.Context) ([]string, error)
}

type toolService struct {
	diagnostics DiagnosticsService
	render      RenderService
	project     ProjectService
	icons       IconService
	projectRoot string
}

func NewToolService(
	diagnostics DiagnosticsService,
	render RenderService,
	project ProjectService,
	icons IconService,
	projectRoot string,
) ToolService {
	return &toolService{
		diagnostics: diagnostics, render: render, project: project, icons: icons,
		projectRoot: projectRoot,
	}
}

func (rcvr *toolService) Tools() []Tool {
	readOnly := map[string]any{"readOnlyHint": true, "destructiveHint": false, "idempotentHint": true, "openWorldHint": false}
	writeSafe := map[string]any{"readOnlyHint": false, "destructiveHint": false, "idempotentHint": true, "openWorldHint": false}
	destructive := map[string]any{"readOnlyHint": false, "destructiveHint": true, "idempotentHint": true, "openWorldHint": false}
	return []Tool{
		newTool("get_icon", "Get SVG icon", "Get one normalized SVG and its registry metadata by namespace:name.", map[string]any{
			"reference": stringProperty("Stable icon identity as namespace:name"),
		}, []string{"reference"}, readOnly),
		newTool("index_docs", "Index documentation", "Incrementally index only Markdown files below the configured docs/ directory. This never discovers .xal files implicitly.", map[string]any{}, nil, writeSafe),
		newTool("inspect_xal", "Inspect XAL concepts", "Analyze explicitly supplied .xal source once and return generic concepts and diagnostics without adding it to the initial RAG corpus.", map[string]any{
			"source": stringProperty("Complete .xal source text"),
			"uri":    stringProperty("Optional stable URI used only to identify this explicit analysis"),
		}, []string{"source"}, readOnly),
		newTool("list_icon_namespaces", "List icon namespaces", "List the registered SVG icon namespaces in deterministic order.", map[string]any{}, nil, readOnly),
		newTool("register_icon", "Register SVG icon", "Validate, normalize, and add or update a namespaced SVG icon in the local registry.", map[string]any{
			"reference":   stringProperty("Stable icon identity as namespace:name"),
			"svg":         stringProperty("Complete SVG markup"),
			"description": stringProperty("Searchable description"),
			"tags":        stringArrayProperty("Searchable tags"),
			"aliases":     stringArrayProperty("Aliases as name or namespace:name"),
			"license":     stringProperty("License identifier or notice"),
			"source":      stringProperty("Source attribution or URL"),
		}, []string{"reference", "svg"}, writeSafe),
		newTool("remove_icon", "Remove SVG icon", "Remove one icon and its aliases, tags, and search row from the local registry.", map[string]any{
			"reference": stringProperty("Stable icon identity as namespace:name"),
		}, []string{"reference"}, destructive),
		newTool("render_svg", "Render XAL as SVG", "Render explicitly supplied .xal source through the shared SVG pipeline and return one artifact per frame unless combineFrames is true.", map[string]any{
			"source":        stringProperty("Complete .xal source text"),
			"mode":          enumProperty("Visual mode", "standard", "network", "aws"),
			"theme":         enumProperty("Color theme", "light", "dark"),
			"combineFrames": map[string]any{"type": "boolean", "description": "Combine every frame on one SVG canvas"},
			"pxPerInch":     map[string]any{"type": "number", "exclusiveMinimum": 0, "description": "Optional output pixels per inch"},
		}, []string{"source"}, readOnly),
		newTool("search_icons", "Search SVG icons", "Search icon names, descriptions, tags, and aliases with an FTS5 query.", map[string]any{
			"query": stringProperty("FTS5 search expression"),
			"limit": limitProperty(),
		}, []string{"query"}, readOnly),
		newTool("search_project", "Search project knowledge", "Search indexed docs Markdown and any explicitly stored concept rows. Initial indexing contains only docs/ Markdown.", map[string]any{
			"query": stringProperty("FTS5 search expression"),
			"limit": limitProperty(),
		}, []string{"query"}, readOnly),
		newTool("project_symbols", "Get project symbols", "Return the indexed semantic symbol tree for one document URI.", map[string]any{
			"uri": stringProperty("Indexed document URI"),
		}, []string{"uri"}, readOnly),
		newTool("validate_xal", "Validate XAL", "Run shared parser, layout, and connection diagnostics for explicitly supplied .xal source.", map[string]any{
			"source": stringProperty("Complete .xal source text"),
		}, []string{"source"}, readOnly),
	}
}

func (rcvr *toolService) Call(ctx context.Context, name string, arguments json.RawMessage) (ToolResult, error) {
	if err := ctx.Err(); err != nil {
		return ToolResult{}, err
	}
	switch name {
	case "validate_xal":
		return rcvr.validateXAL(ctx, arguments)
	case "inspect_xal":
		return rcvr.inspectXAL(ctx, arguments)
	case "render_svg":
		return rcvr.renderSVG(ctx, arguments)
	case "index_docs":
		return rcvr.indexDocs(ctx, arguments)
	case "search_project":
		return rcvr.searchProject(ctx, arguments)
	case "project_symbols":
		return rcvr.projectSymbols(ctx, arguments)
	case "search_icons":
		return rcvr.searchIcons(ctx, arguments)
	case "get_icon":
		return rcvr.getIcon(ctx, arguments)
	case "register_icon":
		return rcvr.registerIcon(ctx, arguments)
	case "remove_icon":
		return rcvr.removeIcon(ctx, arguments)
	case "list_icon_namespaces":
		return rcvr.listIconNamespaces(ctx, arguments)
	default:
		return ToolResult{}, ErrUnknownTool
	}
}

func (rcvr *toolService) projectSymbols(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		URI string `json:"uri"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	args.URI = strings.TrimSpace(args.URI)
	if args.URI == "" {
		return ToolResult{}, errors.New("project document URI must not be empty")
	}
	if rcvr.project == nil {
		return ToolResult{}, errors.New("project service is unavailable")
	}
	symbols, err := rcvr.project.Symbols(ctx, args.URI)
	if err != nil {
		return ToolResult{}, err
	}
	return structuredToolResult(map[string]any{"uri": args.URI, "symbols": symbols})
}

func (rcvr *toolService) validateXAL(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Source string `json:"source"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if err := validateSource(args.Source); err != nil {
		return ToolResult{}, err
	}
	if rcvr.diagnostics == nil {
		return ToolResult{}, errors.New("diagnostics service is unavailable")
	}
	diagnostics, err := rcvr.diagnostics.Diagnose(ctx, []byte(args.Source))
	if err != nil {
		return ToolResult{}, err
	}
	valid := true
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == entity.DiagnosticSeverity("error") {
			valid = false
			break
		}
	}
	return structuredToolResult(map[string]any{"valid": valid, "diagnostics": diagnostics})
}

func (rcvr *toolService) inspectXAL(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Source string `json:"source"`
		URI    string `json:"uri"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if err := validateSource(args.Source); err != nil {
		return ToolResult{}, err
	}
	if rcvr.project == nil {
		return ToolResult{}, errors.New("project service is unavailable")
	}
	if strings.TrimSpace(args.URI) == "" {
		args.URI = "memory://mcp/document.xal"
	}
	analysis, err := rcvr.project.Analyze(ctx, args.URI, []byte(args.Source))
	if err != nil {
		return ToolResult{}, err
	}
	return structuredToolResult(map[string]any{
		"uri": analysis.URI, "kind": analysis.Kind,
		"symbols": analysis.Symbols, "diagnostics": analysis.Diagnostics,
	})
}

func (rcvr *toolService) renderSVG(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Source        string  `json:"source"`
		Mode          string  `json:"mode"`
		Theme         string  `json:"theme"`
		CombineFrames bool    `json:"combineFrames"`
		PxPerInch     float64 `json:"pxPerInch"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if err := validateSource(args.Source); err != nil {
		return ToolResult{}, err
	}
	if rcvr.render == nil {
		return ToolResult{}, errors.New("render service is unavailable")
	}
	if args.Mode == "" {
		args.Mode = "standard"
	}
	if args.Theme == "" {
		args.Theme = "light"
	}
	artifacts, err := rcvr.render.RenderArtifacts(ctx, []byte(args.Source), entity.RenderOptions{
		Format: entity.Format("svg"), Mode: entity.Mode(args.Mode), Theme: args.Theme,
		CombineFrames: args.CombineFrames, PxPerInch: args.PxPerInch,
	})
	if err != nil {
		return ToolResult{}, err
	}
	values := make([]map[string]any, 0, len(artifacts))
	for _, artifact := range artifacts {
		values = append(values, map[string]any{"id": artifact.ID, "svg": string(artifact.Data)})
	}
	return structuredToolResult(map[string]any{"artifacts": values})
}

func (rcvr *toolService) indexDocs(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct{}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if rcvr.project == nil {
		return ToolResult{}, errors.New("project service is unavailable")
	}
	stats, err := rcvr.project.Index(ctx, rcvr.projectRoot)
	if err != nil {
		return ToolResult{}, err
	}
	return structuredToolResult(stats)
}

func (rcvr *toolService) searchProject(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Query string `json:"query"`
		Limit int    `json:"limit"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if rcvr.project == nil {
		return ToolResult{}, errors.New("project service is unavailable")
	}
	args.Limit = defaultLimit(args.Limit)
	results, err := rcvr.project.Search(ctx, args.Query, args.Limit)
	if err != nil {
		return ToolResult{}, err
	}
	return structuredToolResult(results)
}

func (rcvr *toolService) searchIcons(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Query string `json:"query"`
		Limit int    `json:"limit"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if rcvr.icons == nil {
		return ToolResult{}, errors.New("icon service is unavailable")
	}
	results, err := rcvr.icons.Search(ctx, args.Query, defaultLimit(args.Limit))
	if err != nil {
		return ToolResult{}, err
	}
	values := make([]map[string]any, 0, len(results))
	for _, result := range results {
		values = append(values, iconSummaryValue(result))
	}
	return structuredToolResult(values)
}

func (rcvr *toolService) getIcon(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Reference string `json:"reference"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if rcvr.icons == nil {
		return ToolResult{}, errors.New("icon service is unavailable")
	}
	icon, err := rcvr.icons.Get(ctx, args.Reference)
	if err != nil {
		return ToolResult{}, err
	}
	aliases := make([]string, 0, len(icon.Aliases))
	for _, alias := range icon.Aliases {
		aliases = append(aliases, alias.String())
	}
	return structuredToolResult(map[string]any{
		"reference": icon.Ref.String(), "description": icon.Description,
		"svg": string(icon.SVG), "viewBox": icon.ViewBox,
		"width": icon.Width, "height": icon.Height, "license": icon.License,
		"source": icon.Source, "tags": icon.Tags, "aliases": aliases,
	})
}

func (rcvr *toolService) registerIcon(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Reference   string   `json:"reference"`
		SVG         string   `json:"svg"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
		Aliases     []string `json:"aliases"`
		License     string   `json:"license"`
		Source      string   `json:"source"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if len(args.SVG) == 0 || len(args.SVG) > maxIconSVGBytes {
		return ToolResult{}, fmt.Errorf("svg must contain 1-%d UTF-8 bytes", maxIconSVGBytes)
	}
	if rcvr.icons == nil {
		return ToolResult{}, errors.New("icon service is unavailable")
	}
	icon, err := rcvr.icons.Put(ctx, entity.IconRegistration{
		Reference: args.Reference, SVG: []byte(args.SVG), Description: args.Description,
		Tags: args.Tags, Aliases: args.Aliases, License: args.License, Source: args.Source,
	})
	if err != nil {
		return ToolResult{}, err
	}
	return structuredToolResult(iconSummaryValue(entity.IconSummary{
		Ref: icon.Ref, Description: icon.Description, ViewBox: icon.ViewBox,
		Width: icon.Width, Height: icon.Height, License: icon.License, Source: icon.Source,
	}))
}

func (rcvr *toolService) removeIcon(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct {
		Reference string `json:"reference"`
	}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if rcvr.icons == nil {
		return ToolResult{}, errors.New("icon service is unavailable")
	}
	if err := rcvr.icons.Delete(ctx, args.Reference); err != nil {
		return ToolResult{}, err
	}
	return structuredToolResult(map[string]any{"removed": args.Reference})
}

func (rcvr *toolService) listIconNamespaces(ctx context.Context, raw json.RawMessage) (ToolResult, error) {
	var args struct{}
	if err := decodeArguments(raw, &args); err != nil {
		return ToolResult{}, err
	}
	if rcvr.icons == nil {
		return ToolResult{}, errors.New("icon service is unavailable")
	}
	namespaces, err := rcvr.icons.ListNamespaces(ctx)
	if err != nil {
		return ToolResult{}, err
	}
	return structuredToolResult(namespaces)
}

func newTool(name, title, description string, properties map[string]any, required []string, annotations map[string]any) Tool {
	schema := map[string]any{
		"$schema":              "https://json-schema.org/draft/2020-12/schema",
		"type":                 "object",
		"properties":           properties,
		"additionalProperties": false,
	}
	if len(required) > 0 {
		schema["required"] = required
	}
	return Tool{Name: name, Title: title, Description: description, InputSchema: schema, Annotations: annotations}
}

func stringProperty(description string) map[string]any {
	return map[string]any{"type": "string", "description": description}
}

func stringArrayProperty(description string) map[string]any {
	return map[string]any{
		"type": "array", "description": description,
		"items": map[string]any{"type": "string"},
	}
}

func enumProperty(description string, values ...string) map[string]any {
	return map[string]any{"type": "string", "description": description, "enum": values}
}

func limitProperty() map[string]any {
	return map[string]any{"type": "integer", "minimum": 1, "maximum": 100, "default": 30}
}

func decodeArguments(raw json.RawMessage, destination any) error {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		return fmt.Errorf("invalid tool arguments: %w", err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		return errors.New("invalid tool arguments: multiple JSON values")
	}
	return nil
}

func validateSource(source string) error {
	if len(source) == 0 {
		return errors.New("source must not be empty")
	}
	if len(source) > maxXALSourceBytes {
		return fmt.Errorf("source exceeds %d UTF-8 bytes", maxXALSourceBytes)
	}
	return nil
}

func structuredToolResult(value any) (ToolResult, error) {
	serialized, err := json.Marshal(value)
	if err != nil {
		return ToolResult{}, fmt.Errorf("encode tool result: %w", err)
	}
	return ToolResult{
		ResultType:        "complete",
		Content:           []Content{{Type: "text", Text: string(serialized)}},
		StructuredContent: value,
	}, nil
}

func defaultLimit(value int) int {
	if value == 0 {
		return 30
	}
	return value
}

func iconSummaryValue(summary entity.IconSummary) map[string]any {
	return map[string]any{
		"reference": summary.Ref.String(), "description": summary.Description,
		"viewBox": summary.ViewBox, "width": summary.Width, "height": summary.Height,
		"license": summary.License, "source": summary.Source,
	}
}
