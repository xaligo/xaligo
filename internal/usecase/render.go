package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryo-arima/xaligo/internal/config"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
)

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

func RenderExcalidraw(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	return scene, err
}

func RenderSVG(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	planJSON, err := BuildPPTXPlan(ctx, input, opts)
	if err != nil {
		return nil, fmt.Errorf("build SVG plan: %w", err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		return nil, fmt.Errorf("decode SVG plan: %w", err)
	}
	return RenderSVGPlan(plan, opts.PxPerInch)
}

// BuildPPTXPlan produces the shared resolved draw plan used by native SVG,
// native PPTX, and the TypeScript PptxGenJS adapter.
func BuildPPTXPlan(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, entries, err := buildScene(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	planJSON, err := BuildPlanJSON(string(scene), planOptions(opts, entries))
	if err != nil {
		return nil, fmt.Errorf("build PPTX plan: %w", err)
	}
	return planJSON, nil
}

func RenderPPTX(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	planJSON, err := BuildPPTXPlan(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	tmp, err := os.CreateTemp("", "xaligo-*.pptx")
	if err != nil {
		return nil, fmt.Errorf("create temporary PPTX: %w", err)
	}
	path := tmp.Name()
	_ = tmp.Close()
	defer os.Remove(path)
	if err := repository.ExportPptx(entity.PptxExportOptions{
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
	return RenderXYFlowScene(scene)
}

func RenderIsoflow(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	var icons map[string]string
	if opts.Assets != nil && opts.Assets.IsoflowIconsJSON != "" {
		icons, _ = LoadIsoflowIconManifestFS(opts.Assets.FS, opts.Assets.IsoflowIconsJSON)
	} else {
		cfg := config.New()
		icons, _ = LoadIsoflowIconManifest(filepath.Join(cfg.ProjectRoot, "etc", "resources", "aws", "isoflow-icons.json"))
	}
	return RenderIsoflowWithIcons(scene, icons)
}

func buildScene(ctx context.Context, input []byte, opts RenderOptions) ([]byte, []entity.ServiceEntry, error) {
	if err := checkContext(ctx); err != nil {
		return nil, nil, err
	}
	if err := ValidateRenderOptions(opts); err != nil {
		return nil, nil, err
	}
	theme, _ := entity.NormalizeTheme(opts.Theme)
	doc, err := Parse(bytes.NewReader(input))
	if err != nil {
		return nil, nil, fmt.Errorf("parse DSL: %w", err)
	}
	root, err := Build(doc)
	if err != nil {
		return nil, nil, fmt.Errorf("build layout: %w", err)
	}
	entries, abbreviations, err := serviceOptions(opts)
	if err != nil {
		return nil, nil, err
	}
	var connections []*entity.Node
	for _, child := range doc.Root.Children {
		if child.Tag == "connection" {
			connections = append(connections, child)
		}
	}
	var scene []byte
	if opts.Assets != nil {
		itemSize := opts.Assets.ItemIconSize
		if itemSize <= 0 {
			itemSize = 32
		}
		scene, err = BuildJSONWithFS(root, opts.Assets.FS, opts.Assets.CatalogCSV, opts.Assets.GroupIconsDir, itemSize, connections, abbreviations)
	} else {
		cfg := config.New()
		scene, err = BuildJSON(root, filepath.Join(cfg.AssetDir_, "Architecture-Group-Icons"), cfg.SvcCatalogCSV, cfg.ProjectRoot, cfg.ItemIconSize, connections, abbreviations, nil)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("build excalidraw JSON: %w", err)
	}
	scene, err = ApplyThemeJSON(scene, theme)
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

func planOptions(opts RenderOptions, entries []entity.ServiceEntry) entity.PptxOptions {
	legend := make([]entity.LegendEntry, 0, len(entries))
	for _, entry := range entries {
		if entry.CatalogID > 0 && entry.OfficialName != "" {
			legend = append(legend, entity.LegendEntry{CatalogID: entry.CatalogID, Abbreviation: entry.Abbreviation, OfficialName: entry.OfficialName})
		}
	}
	return entity.PptxOptions{
		Theme: opts.Theme, PxPerInch: opts.PxPerInch, ArrowStyle: opts.ArrowStyle,
		ArrowStubPx: opts.ArrowStubPx, ArrowMargin: opts.ArrowMarginPx,
		PaperSize: opts.PaperSize, Orientation: opts.Orientation,
		PaperMargin: opts.PaperMarginIn, PaperMarginTop: opts.PaperMarginTopIn, PaperMarginRight: opts.PaperMarginRightIn,
		PaperMarginBottom: opts.PaperMarginBottomIn, PaperMarginLeft: opts.PaperMarginLeftIn, LegendEntries: legend,
	}
}
