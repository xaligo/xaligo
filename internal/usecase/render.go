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
	"github.com/ryo-arima/xaligo/internal/share"
)

var (
	logger     = share.DefaultLogger()
	IURR001    = share.NewMCode("IURR-001", "Render context check failed")
	IURR002    = share.NewMCode("IURR-002", "Render validate render options failed")
	IURR003    = share.NewMCode("IURR-003", "Render default format branch")
	IURR004    = share.NewMCode("IURR-004", "Render excalidraw branch")
	IURR005    = share.NewMCode("IURR-005", "Render SVG branch")
	IURR006    = share.NewMCode("IURR-006", "Render PPTX branch")
	IURR007    = share.NewMCode("IURR-007", "Render XYFlow branch")
	IURR008    = share.NewMCode("IURR-008", "Render Isoflow branch")
	IURR009    = share.NewMCode("IURR-009", "Render unknown format branch")
	IURRI001   = share.NewMCode("IURRI-001", "Render Isoflow embedded icons branch")
	IURRI002   = share.NewMCode("IURRI-002", "Render Isoflow native icons branch")
	IURBS001   = share.NewMCode("IURBS-001", "Build scene context check failed")
	IURBS002   = share.NewMCode("IURBS-002", "Build scene validate render options failed")
	IURBS003   = share.NewMCode("IURBS-003", "Build scene connection node branch")
	IURBS004   = share.NewMCode("IURBS-004", "Build scene embedded assets branch")
	IURBS005   = share.NewMCode("IURBS-005", "Build scene default embedded item size branch")
	IURBS006   = share.NewMCode("IURBS-006", "Build scene native assets branch")
	IURSO001   = share.NewMCode("IURSO-001", "Service options no services CSV branch")
	IURSO002   = share.NewMCode("IURSO-002", "Service options services CSV branch")
	IURSO003   = share.NewMCode("IURSO-003", "Service options read services CSV failed")
	IURSO004   = share.NewMCode("IURSO-004", "Service options service abbreviation branch")
	IURPO001   = share.NewMCode("IURPO-001", "Plan options legend entry branch")
	IURRE001   = share.NewMCode("IURRE-001", "Render Excalidraw completed")
	IURRS001   = share.NewMCode("IURRS-001", "Render SVG build plan failed")
	IURRS002   = share.NewMCode("IURRS-002", "Render SVG decode plan failed")
	IURBPP001  = share.NewMCode("IURBPP-001", "Build PPTX plan build scene failed")
	IURBPP002  = share.NewMCode("IURBPP-002", "Build PPTX plan build plan failed")
	IURRP001   = share.NewMCode("IURRP-001", "Render PPTX build plan failed")
	IURRP002   = share.NewMCode("IURRP-002", "Render PPTX create temp failed")
	IURRP003   = share.NewMCode("IURRP-003", "Render PPTX export failed")
	IURRP004   = share.NewMCode("IURRP-004", "Render PPTX context check failed")
	IURRP005   = share.NewMCode("IURRP-005", "Render PPTX read temp failed")
	IURRXYF001 = share.NewMCode("IURRXYF-001", "Render XYFlow build scene failed")
	IURRXYF002 = share.NewMCode("IURRXYF-002", "Render XYFlow scene failed")
	IURRI003   = share.NewMCode("IURRI-003", "Render Isoflow build scene failed")
	IURRI004   = share.NewMCode("IURRI-004", "Render Isoflow with icons failed")
	IURBS007   = share.NewMCode("IURBS-007", "Build scene parse DSL failed")
	IURBS008   = share.NewMCode("IURBS-008", "Build scene build layout failed")
	IURBS009   = share.NewMCode("IURBS-009", "Build scene service options failed")
	IURBS010   = share.NewMCode("IURBS-010", "Build scene build JSON failed")
	IURBS011   = share.NewMCode("IURBS-011", "Build scene apply theme failed")
)

// Render validates and renders a .xal document to the selected format.
func Render(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURR001, "context check failed", map[string]any{"error": err})
		return nil, err
	}
	if err := ValidateRenderOptions(opts); err != nil {
		logger.ERROR(IURR002, "validate render options failed", map[string]any{"format": opts.Format, "error": err})
		return nil, err
	}
	format := entity.Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	if format == "" {
		logger.DEBUG(IURR003, "branch default format")
		format = FormatExcalidraw
	}
	switch format {
	case FormatExcalidraw:
		logger.DEBUG(IURR004, "branch excalidraw")
		return RenderExcalidraw(ctx, input, opts)
	case FormatSVG:
		logger.DEBUG(IURR005, "branch svg")
		return RenderSVG(ctx, input, opts)
	case FormatPPTX:
		logger.DEBUG(IURR006, "branch pptx")
		return RenderPPTX(ctx, input, opts)
	case FormatXYFlow:
		logger.DEBUG(IURR007, "branch xyflow")
		return RenderXYFlow(ctx, input, opts)
	case FormatIsoflow:
		logger.DEBUG(IURR008, "branch isoflow")
		return RenderIsoflow(ctx, input, opts)
	default:
		logger.ERROR(IURR009, "branch unknown format", map[string]any{"format": format})
		return nil, fmt.Errorf("unknown render format %q", format)
	}
}

func RenderExcalidraw(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	if err == nil {
		logger.DEBUG(IURRE001, "completed")
	}
	return scene, err
}

func RenderSVG(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	planJSON, err := BuildPPTXPlan(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRS001, "build plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("build SVG plan: %w", err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		logger.ERROR(IURRS002, "decode plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("decode SVG plan: %w", err)
	}
	return RenderSVGPlan(plan, opts.PxPerInch)
}

// BuildPPTXPlan produces the shared resolved draw plan used by native SVG,
// native PPTX, and the TypeScript PptxGenJS adapter.
func BuildPPTXPlan(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, entries, err := buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURBPP001, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	planJSON, err := BuildPlanJSON(string(scene), planOptions(opts, entries))
	if err != nil {
		logger.ERROR(IURBPP002, "build plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("build PPTX plan: %w", err)
	}
	return planJSON, nil
}

func RenderPPTX(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	planJSON, err := BuildPPTXPlan(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRP001, "build plan failed", map[string]any{"error": err})
		return nil, err
	}
	tmp, err := os.CreateTemp("", "xaligo-*.pptx")
	if err != nil {
		logger.ERROR(IURRP002, "create temp failed", map[string]any{"error": err})
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
		logger.ERROR(IURRP003, "export failed", map[string]any{"error": err})
		return nil, err
	}
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURRP004, "context check failed", map[string]any{"error": err})
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		logger.ERROR(IURRP005, "read temp failed", map[string]any{"path": path, "error": err})
		return nil, err
	}
	return data, nil
}

func RenderXYFlow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRXYF001, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	out, err := RenderXYFlowScene(scene)
	if err != nil {
		logger.ERROR(IURRXYF002, "render scene failed", map[string]any{"error": err})
		return nil, err
	}
	return out, nil
}

func RenderIsoflow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRI003, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	var icons map[string]string
	if opts.Assets != nil && opts.Assets.IsoflowIconsJSON != "" {
		logger.DEBUG(IURRI001, "branch embedded isoflow icons", map[string]any{"path": opts.Assets.IsoflowIconsJSON})
		icons, _ = LoadIsoflowIconManifestFS(opts.Assets.FS, opts.Assets.IsoflowIconsJSON)
	} else {
		logger.DEBUG(IURRI002, "branch native isoflow icons")
		cfg := config.New()
		icons, _ = LoadIsoflowIconManifest(filepath.Join(cfg.ProjectRoot, "etc", "resources", "aws", "isoflow-icons.json"))
	}
	out, err := RenderIsoflowWithIcons(scene, icons)
	if err != nil {
		logger.ERROR(IURRI004, "render with icons failed", map[string]any{"error": err})
		return nil, err
	}
	return out, nil
}

func buildScene(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, []entity.ServiceEntry, error) {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURBS001, "context check failed", map[string]any{"error": err})
		return nil, nil, err
	}
	if err := ValidateRenderOptions(opts); err != nil {
		logger.ERROR(IURBS002, "validate render options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	theme, _ := entity.NormalizeTheme(opts.Theme)
	doc, err := Parse(bytes.NewReader(input))
	if err != nil {
		logger.ERROR(IURBS007, "parse DSL failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("parse DSL: %w", err)
	}
	root, err := Build(doc)
	if err != nil {
		logger.ERROR(IURBS008, "build layout failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build layout: %w", err)
	}
	entries, abbreviations, err := serviceOptions(opts)
	if err != nil {
		logger.ERROR(IURBS009, "service options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	var connections []*entity.Node
	for _, child := range doc.Root.Children {
		if child.Tag == "connection" {
			logger.DEBUG(IURBS003, "branch connection node", map[string]any{"tag": child.Tag})
			connections = append(connections, child)
		}
	}
	var scene []byte
	if opts.Assets != nil {
		logger.DEBUG(IURBS004, "branch embedded assets")
		itemSize := opts.Assets.ItemIconSize
		if itemSize <= 0 {
			logger.DEBUG(IURBS005, "branch default embedded item size")
			itemSize = 32
		}
		scene, err = BuildJSONWithFS(root, opts.Assets.FS, opts.Assets.CatalogCSV, opts.Assets.GroupIconsDir, itemSize, connections, abbreviations)
	} else {
		logger.DEBUG(IURBS006, "branch native assets")
		cfg := config.New()
		scene, err = BuildJSON(root, filepath.Join(cfg.AssetDir_, "Architecture-Group-Icons"), cfg.SvcCatalogCSV, cfg.ProjectRoot, cfg.ItemIconSize, connections, abbreviations, nil)
	}
	if err != nil {
		logger.ERROR(IURBS010, "build JSON failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build excalidraw JSON: %w", err)
	}
	scene, err = ApplyThemeJSON(scene, theme)
	if err != nil {
		logger.ERROR(IURBS011, "apply theme failed", map[string]any{"error": err})
	}
	return scene, entries, err
}
