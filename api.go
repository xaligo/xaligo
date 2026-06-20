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
	"github.com/ryo-arima/xaligo/internal/layout"
	"github.com/ryo-arima/xaligo/internal/model"
	"github.com/ryo-arima/xaligo/internal/parser"
	"github.com/ryo-arima/xaligo/internal/pptxplan"
	"github.com/ryo-arima/xaligo/internal/repository"
	svgrenderer "github.com/ryo-arima/xaligo/internal/svg"
)

type Mode string
type Format string

const (
	ModeStandard Mode = "standard"
	ModeNetwork  Mode = "network"
	ModeAWS      Mode = "aws"

	FormatExcalidraw Format = "excalidraw"
	FormatSVG        Format = "svg"
	FormatPPTX       Format = "pptx"
	FormatXYFlow     Format = "xyflow"
	FormatIsoflow    Format = "isoflow"
)

var ErrNotImplemented = errors.New("renderer not implemented")

// RenderOptions contains renderer-independent presentation and routing options.
// Abbreviations is primarily useful to adapters that already parsed a
// services.csv file; ServicesCSV is the convenient in-memory equivalent.
type RenderOptions struct {
	Mode          Mode
	Format        Format
	Theme         string
	ServicesCSV   []byte
	Abbreviations map[int]string
	PxPerInch     float64
	ArrowStyle    string
	ArrowStubPx   float64
	ArrowMarginPx float64
	PaperSize     string
	Orientation   string

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

func RenderXYFlow(context.Context, []byte, RenderOptions) ([]byte, error) {
	return nil, fmt.Errorf("xyflow: %w", ErrNotImplemented)
}

func RenderIsoflow(context.Context, []byte, RenderOptions) ([]byte, error) {
	return nil, fmt.Errorf("isoflow: %w", ErrNotImplemented)
}

// Validate runs the same parser and layout validation used by Render.
func Validate(ctx context.Context, input []byte) error {
	if err := checkContext(ctx); err != nil {
		return err
	}
	doc, err := parser.Parse(bytes.NewReader(input))
	if err != nil {
		return fmt.Errorf("parse DSL: %w", err)
	}
	if _, err := layout.Build(doc); err != nil {
		return fmt.Errorf("build layout: %w", err)
	}
	return checkContext(ctx)
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
		PaperSize: opts.PaperSize, Orientation: opts.Orientation, LegendEntries: legend,
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
