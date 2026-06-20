// Package xaligo exposes the stable, renderer-independent API used by the CLI,
// preview servers, editors and other Go integrations.
package xaligo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryo-arima/xaligo/internal/config"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/excalidraw"
	isoflowrenderer "github.com/ryo-arima/xaligo/internal/isoflow"
	"github.com/ryo-arima/xaligo/internal/layout"
	"github.com/ryo-arima/xaligo/internal/model"
	"github.com/ryo-arima/xaligo/internal/parser"
	"github.com/ryo-arima/xaligo/internal/pptxplan"
	"github.com/ryo-arima/xaligo/internal/repository"
	svgrenderer "github.com/ryo-arima/xaligo/internal/svg"
	xyflowrenderer "github.com/ryo-arima/xaligo/internal/xyflow"
)

type Mode string
type Format string
type DiagnosticSeverity string

const (
	ModeStandard Mode = "standard"
	ModeNetwork  Mode = "network"
	ModeAWS      Mode = "aws"

	FormatExcalidraw Format = "excalidraw"
	FormatSVG        Format = "svg"
	FormatPPTX       Format = "pptx"
	FormatXYFlow     Format = "xyflow"
	FormatIsoflow    Format = "isoflow"

	SeverityError DiagnosticSeverity = "error"
)

var ErrNotImplemented = errors.New("renderer not implemented")

type Diagnostic struct {
	Severity DiagnosticSeverity `json:"severity"`
	Message  string             `json:"message"`
	Offset   int                `json:"offset,omitempty"`
	Line     int                `json:"line,omitempty"`
	Column   int                `json:"column,omitempty"`
}

type DiagnosticsError struct {
	Diagnostics []Diagnostic
}

func (e *DiagnosticsError) Error() string {
	if len(e.Diagnostics) == 0 {
		return "validation failed"
	}
	d := e.Diagnostics[0]
	if d.Line > 0 {
		return fmt.Sprintf("line %d, column %d: %s", d.Line, d.Column, d.Message)
	}
	return d.Message
}

// RenderOptions contains renderer-independent presentation and routing options.
// Abbreviations is primarily useful to adapters that already parsed a
// services.csv file; ServicesCSV is the convenient in-memory equivalent.
type RenderOptions struct {
	Mode                Mode
	Format              Format
	Theme               string
	ServicesCSV         []byte
	Abbreviations       map[int]string
	PxPerInch           float64
	ArrowStyle          string
	ArrowStubPx         float64
	ArrowMarginPx       float64
	PaperSize           string
	Orientation         string
	PaperMarginIn       float64
	PaperMarginTopIn    float64
	PaperMarginRightIn  float64
	PaperMarginBottomIn float64
	PaperMarginLeftIn   float64

	Title            string
	Author           string
	Company          string
	Subject          string
	Compression      *bool
	PPTXExporterWASM string
}

// Render validates and renders a .xal document to the selected format.
func Render(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	if err := ValidateRenderOptions(opts); err != nil {
		return nil, err
	}
	format := Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	if format == "" {
		format = FormatExcalidraw
	}
	switch format {
	case FormatExcalidraw:
		return RenderExcalidraw(ctx, input, opts)
	case FormatSVG:
		return RenderSVG(ctx, input, opts)
	case FormatPPTX:
		return RenderPPTX(ctx, input, opts)
	case FormatXYFlow:
		return RenderXYFlow(ctx, input, opts)
	case FormatIsoflow:
		return RenderIsoflow(ctx, input, opts)
	default:
		return nil, fmt.Errorf("unknown render format %q", format)
	}
}

// ValidateRenderOptions validates mode, format and shared presentation values
// without parsing an input document.
func ValidateRenderOptions(opts RenderOptions) error {
	if err := validateMode(opts.Mode); err != nil {
		return err
	}
	if opts.PaperMarginIn < 0 || opts.PaperMarginTopIn < 0 || opts.PaperMarginRightIn < 0 || opts.PaperMarginBottomIn < 0 || opts.PaperMarginLeftIn < 0 {
		return fmt.Errorf("paper margins must be non-negative")
	}
	if _, err := excalidraw.NormalizeTheme(opts.Theme); err != nil {
		return err
	}
	format := Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	switch format {
	case "", FormatExcalidraw, FormatSVG, FormatPPTX, FormatXYFlow, FormatIsoflow:
		return nil
	default:
		return fmt.Errorf("unknown render format %q", format)
	}
}

func RenderExcalidraw(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	return scene, err
}

func RenderSVG(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, entries, err := buildScene(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	planJSON, err := pptxplan.BuildPlanJSON(string(scene), planOptions(opts, entries))
	if err != nil {
		return nil, fmt.Errorf("build SVG plan: %w", err)
	}
	var plan pptxplan.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		return nil, fmt.Errorf("decode SVG plan: %w", err)
	}
	return svgrenderer.RenderPlan(plan, opts.PxPerInch)
}

func RenderPPTX(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, entries, err := buildScene(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	planJSON, err := pptxplan.BuildPlanJSON(string(scene), planOptions(opts, entries))
	if err != nil {
		return nil, fmt.Errorf("build PPTX plan: %w", err)
	}
	tmp, err := os.CreateTemp("", "xaligo-*.pptx")
	if err != nil {
		return nil, fmt.Errorf("create temporary PPTX: %w", err)
	}
	path := tmp.Name()
	_ = tmp.Close()
	defer os.Remove(path)
	if err := repository.ExportPptx(repository.PptxExportOptions{
		PlanJSON: planJSON, Output: path, Title: opts.Title, Author: opts.Author,
		Company: opts.Company, Subject: opts.Subject, Compression: opts.Compression,
		ExporterWASM: opts.PPTXExporterWASM,
	}); err != nil {
		return nil, err
	}
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return os.ReadFile(path)
}

func RenderXYFlow(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	return xyflowrenderer.Render(scene)
}

func RenderIsoflow(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	cfg := config.New()
	icons, _ := isoflowrenderer.LoadIconManifest(filepath.Join(cfg.ProjectRoot, "etc", "resources", "aws", "isoflow-icons.json"))
	return isoflowrenderer.RenderWithIcons(scene, icons)
}

// Validate runs the same parser and layout validation used by Render.
func Validate(ctx context.Context, input []byte) error {
	diagnostics, err := Diagnose(ctx, input)
	if err != nil {
		return err
	}
	if len(diagnostics) > 0 {
		return &DiagnosticsError{Diagnostics: diagnostics}
	}
	return nil
}

// Diagnose validates a document and returns editor-friendly source positions.
func Diagnose(ctx context.Context, input []byte) ([]Diagnostic, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	doc, err := parser.Parse(bytes.NewReader(input))
	if err != nil {
		return []Diagnostic{diagnosticFromError(err)}, nil
	}
	if _, err := layout.Build(doc); err != nil {
		return []Diagnostic{{Severity: SeverityError, Message: err.Error()}}, nil
	}
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

func diagnosticFromError(err error) Diagnostic {
	diagnostic := Diagnostic{Severity: SeverityError, Message: err.Error()}
	var positioned *parser.Error
	if errors.As(err, &positioned) {
		diagnostic.Message = positioned.Err.Error()
		diagnostic.Offset = positioned.Position.Offset
		diagnostic.Line = positioned.Position.Line
		diagnostic.Column = positioned.Position.Column
	}
	return diagnostic
}

func buildScene(ctx context.Context, input []byte, opts RenderOptions) ([]byte, []entity.ServiceEntry, error) {
	if err := checkContext(ctx); err != nil {
		return nil, nil, err
	}
	if err := validateMode(opts.Mode); err != nil {
		return nil, nil, err
	}
	theme, err := excalidraw.NormalizeTheme(opts.Theme)
	if err != nil {
		return nil, nil, err
	}
	doc, err := parser.Parse(bytes.NewReader(input))
	if err != nil {
		return nil, nil, fmt.Errorf("parse DSL: %w", err)
	}
	root, err := layout.Build(doc)
	if err != nil {
		return nil, nil, fmt.Errorf("build layout: %w", err)
	}
	entries, abbreviations, err := serviceOptions(opts)
	if err != nil {
		return nil, nil, err
	}
	var connections []*model.Node
	for _, child := range doc.Root.Children {
		if child.Tag == "connection" {
			connections = append(connections, child)
		}
	}
	cfg := config.New()
	scene, err := excalidraw.BuildJSON(root, filepath.Join(cfg.AssetDir_, "Architecture-Group-Icons"), cfg.SvcCatalogCSV, cfg.ProjectRoot, cfg.ItemIconSize, connections, abbreviations, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("build excalidraw JSON: %w", err)
	}
	scene, err = excalidraw.ApplyThemeJSON(scene, theme)
	return scene, entries, err
}

func serviceOptions(opts RenderOptions) ([]entity.ServiceEntry, map[int]string, error) {
	abbreviations := make(map[int]string, len(opts.Abbreviations))
	for id, value := range opts.Abbreviations {
		abbreviations[id] = value
	}
	if len(bytes.TrimSpace(opts.ServicesCSV)) == 0 {
		return nil, abbreviations, nil
	}
	entries, err := repository.ReadServiceListFromReader(bytes.NewReader(opts.ServicesCSV))
	if err != nil {
		return nil, nil, fmt.Errorf("read services CSV: %w", err)
	}
	for _, entry := range entries {
		if entry.CatalogID > 0 && entry.Abbreviation != "" {
			abbreviations[entry.CatalogID] = entry.Abbreviation
		}
	}
	return entries, abbreviations, nil
}

func planOptions(opts RenderOptions, entries []entity.ServiceEntry) pptxplan.Options {
	legend := make([]pptxplan.LegendEntry, 0, len(entries))
	for _, entry := range entries {
		if entry.CatalogID > 0 && entry.OfficialName != "" {
			legend = append(legend, pptxplan.LegendEntry{CatalogID: entry.CatalogID, Abbreviation: entry.Abbreviation, OfficialName: entry.OfficialName})
		}
	}
	return pptxplan.Options{
		Theme: opts.Theme, PxPerInch: opts.PxPerInch, ArrowStyle: opts.ArrowStyle,
		ArrowStubPx: opts.ArrowStubPx, ArrowMargin: opts.ArrowMarginPx,
		PaperSize: opts.PaperSize, Orientation: opts.Orientation,
		PaperMargin: opts.PaperMarginIn, PaperMarginTop: opts.PaperMarginTopIn, PaperMarginRight: opts.PaperMarginRightIn,
		PaperMarginBottom: opts.PaperMarginBottomIn, PaperMarginLeft: opts.PaperMarginLeftIn, LegendEntries: legend,
	}
}

func validateMode(mode Mode) error {
	normalized := Mode(strings.ToLower(strings.TrimSpace(string(mode))))
	switch normalized {
	case "", ModeStandard, ModeNetwork, ModeAWS:
		return nil
	case "aws-2.5d", "topology":
		return fmt.Errorf("mode %q: %w", normalized, ErrNotImplemented)
	default:
		return fmt.Errorf("unknown render mode %q", normalized)
	}
}

func checkContext(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	return ctx.Err()
}
