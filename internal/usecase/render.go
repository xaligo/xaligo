package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/share"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

// RenderUsecase owns format dispatch and the shared render pipeline.
type RenderUsecase interface {
	ValidateRenderOptions(entity.RenderOptions) error
	Render(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderExcalidraw(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderSVG(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderPPTX(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderXYFlow(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderIsoflow(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	BuildPPTXPlan(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	NewPreviewRepository(string, entity.PreviewOptions) (repository.PreviewRepository, error)
}

type renderUsecase struct {
	excalidrawRepository repository.ExcalidrawRepository
	xaligoRepository     repository.XaligoRepository
	powerpointRepository repository.PowerpointRepository
	isoflowRepository    repository.IsoflowRepository
	svgRepository        repository.SVGRepository
	xyFlowRepository     repository.XYFlowRepository
}

func NewRenderUsecase(
	excalidrawRepository repository.ExcalidrawRepository,
	xaligoRepository repository.XaligoRepository,
	powerpointRepository repository.PowerpointRepository,
	isoflowRepository repository.IsoflowRepository,
	svgRepository repository.SVGRepository,
	xyFlowRepository repository.XYFlowRepository,
) RenderUsecase {
	return &renderUsecase{
		excalidrawRepository: excalidrawRepository,
		xaligoRepository:     xaligoRepository,
		powerpointRepository: powerpointRepository,
		isoflowRepository:    isoflowRepository,
		svgRepository:        svgRepository,
		xyFlowRepository:     xyFlowRepository,
	}
}

const (
	ModeStandard = v1engine.ModeStandardV1EngineOptionRender
	ModeNetwork  = v1engine.ModeNetworkV1EngineOptionRender
	ModeAWS      = v1engine.ModeAWSV1EngineOptionRender

	FormatExcalidraw = v1engine.FormatExcalidrawV1EngineOptionRender
	FormatSVG        = v1engine.FormatSVGV1EngineOptionRender
	FormatPPTX       = v1engine.FormatPPTXV1EngineOptionRender
	FormatXYFlow     = v1engine.FormatXYFlowV1EngineOptionRender
	FormatIsoflow    = v1engine.FormatIsoflowV1EngineOptionRender

	SeverityError = v1engine.SeverityErrorV1EngineOptionRender
)

var ErrNotImplemented = v1engine.ErrNotImplementedV1EngineOptionRender

var (
	logger   = share.DefaultLogger()
	IURR001  = share.NewMCode("IURR-001", "Render context check failed")
	IURR002  = share.NewMCode("IURR-002", "Render validate render options failed")
	IURR003  = share.NewMCode("IURR-003", "Render default format branch")
	IURR004  = share.NewMCode("IURR-004", "Render excalidraw branch")
	IURR005  = share.NewMCode("IURR-005", "Render SVG branch")
	IURR006  = share.NewMCode("IURR-006", "Render PPTX branch")
	IURR007  = share.NewMCode("IURR-007", "Render XYFlow branch")
	IURR008  = share.NewMCode("IURR-008", "Render Isoflow branch")
	IURR009  = share.NewMCode("IURR-009", "Render unknown format branch")
	IURSO001 = share.NewMCode("IURSO-001", "Service options no services CSV branch")
	IURSO002 = share.NewMCode("IURSO-002", "Service options services CSV branch")
	IURSO003 = share.NewMCode("IURSO-003", "Service options read services CSV failed")
	IURSO005 = share.NewMCode("IURSO-005", "Service options legend validation failed")
	IURBS001 = share.NewMCode("IURBS-001", "Build scene context check failed")
	IURBS002 = share.NewMCode("IURBS-002", "Build scene validate render options failed")
	IURBS004 = share.NewMCode("IURBS-004", "Build scene embedded assets branch")
	IURBS005 = share.NewMCode("IURBS-005", "Build scene default embedded item size branch")
	IURBS006 = share.NewMCode("IURBS-006", "Build scene native assets branch")
	IURBS007 = share.NewMCode("IURBS-007", "Build scene parse DSL failed")
	IURBS008 = share.NewMCode("IURBS-008", "Build scene build layout failed")
	IURBS009 = share.NewMCode("IURBS-009", "Build scene service options failed")
	IURBS010 = share.NewMCode("IURBS-010", "Build scene build JSON failed")
	IURBS011 = share.NewMCode("IURBS-011", "Build scene apply theme failed")
)

func (rcvr *renderUsecase) ValidateRenderOptions(opts entity.RenderOptions) error {
	return v1engine.ValidateRenderOptionsV1EngineOptionRender(opts)
}

func (rcvr *renderUsecase) Render(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
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
	opts.Format = format
	switch format {
	case FormatExcalidraw:
		logger.DEBUG(IURR004, "branch excalidraw")
		return rcvr.RenderExcalidraw(ctx, input, opts)
	case FormatSVG:
		logger.DEBUG(IURR005, "branch svg")
		return rcvr.RenderSVG(ctx, input, opts)
	case FormatPPTX:
		logger.DEBUG(IURR006, "branch pptx")
		return rcvr.RenderPPTX(ctx, input, opts)
	case FormatXYFlow:
		logger.DEBUG(IURR007, "branch xyflow")
		return rcvr.RenderXYFlow(ctx, input, opts)
	case FormatIsoflow:
		logger.DEBUG(IURR008, "branch isoflow")
		return rcvr.RenderIsoflow(ctx, input, opts)
	default:
		logger.ERROR(IURR009, "branch unknown format", map[string]any{"format": format})
		return nil, fmt.Errorf("unknown render format %q", format)
	}
}

var IURRE001 = share.NewMCode("IURRE-001", "Render Excalidraw completed")

func (rcvr *renderUsecase) RenderExcalidraw(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := rcvr.buildScene(ctx, input, opts)
	if err == nil {
		logger.DEBUG(IURRE001, "completed")
	}
	return scene, err
}

var (
	IURRS001 = share.NewMCode("IURRS-001", "Render SVG build plan failed")
	IURRS002 = share.NewMCode("IURRS-002", "Render SVG decode plan failed")
)

func (rcvr *renderUsecase) RenderSVG(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	planJSON, err := rcvr.buildPlan(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRS001, "build plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("build SVG plan: %w", err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		logger.ERROR(IURRS002, "decode plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("decode SVG plan: %w", err)
	}
	return rcvr.svgRepository.Render(plan, opts.PxPerInch, opts.SVGLegendPosition)
}

var (
	IURBPP001 = share.NewMCode("IURBPP-001", "Build PPTX plan build scene failed")
	IURBPP002 = share.NewMCode("IURBPP-002", "Build PPTX plan build plan failed")
	IURRP001  = share.NewMCode("IURRP-001", "Render PPTX build plan failed")
	IURRP003  = share.NewMCode("IURRP-003", "Render PPTX export failed")
	IURRP004  = share.NewMCode("IURRP-004", "Render PPTX context check failed")
)

func (rcvr *renderUsecase) buildPlan(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, entries, err := rcvr.buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURBPP001, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	planJSON, err := v1engine.BuildPlanJSONV1EnginePlanBuild(string(scene), v1engine.ResolvePlanOptionsV1EngineOptionPlan(opts, entries))
	if err != nil {
		logger.ERROR(IURBPP002, "build plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("build draw plan: %w", err)
	}
	return planJSON, nil
}

// BuildPPTXPlan remains as the format-specific compatibility boundary for WASM.
func (rcvr *renderUsecase) BuildPPTXPlan(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.buildPlan(ctx, input, opts)
}

func (rcvr *renderUsecase) NewPreviewRepository(path string, opts entity.PreviewOptions) (repository.PreviewRepository, error) {
	return repository.NewPreviewRepository(
		path,
		opts,
		rcvr.RenderSVG,
		rcvr.ValidateRenderOptions,
		rcvr.diagnose,
		rcvr.xaligoRepository.ReadSource,
	)
}

func (rcvr *renderUsecase) diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	diagnostics := v1engine.DiagnoseV1EngineDiagnoseDocument(input)
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return diagnostics, nil
}

func (rcvr *renderUsecase) RenderPPTX(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	planJSON, err := rcvr.buildPlan(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRP001, "build plan failed", map[string]any{"error": err})
		return nil, err
	}
	data, err := rcvr.powerpointRepository.ExportPptxBytes(ctx, entity.PptxExportOptions{
		PlanJSON: planJSON, Title: opts.Title, Author: opts.Author,
		Company: opts.Company, Subject: opts.Subject, Compression: opts.Compression,
		ExporterWASM: opts.PPTXExporterWASM,
	})
	if err != nil {
		logger.ERROR(IURRP003, "export failed", map[string]any{"error": err})
		return nil, err
	}
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURRP004, "context check failed", map[string]any{"error": err})
		return nil, err
	}
	return data, nil
}

var (
	IURRXYF001 = share.NewMCode("IURRXYF-001", "Render XYFlow build scene failed")
	IURRXYF002 = share.NewMCode("IURRXYF-002", "Render XYFlow scene failed")
)

func (rcvr *renderUsecase) RenderXYFlow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := rcvr.buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRXYF001, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	out, err := rcvr.xyFlowRepository.Render(scene)
	if err != nil {
		logger.ERROR(IURRXYF002, "render scene failed", map[string]any{"error": err})
		return nil, err
	}
	return out, nil
}

var (
	IURRI001 = share.NewMCode("IURRI-001", "Render Isoflow load icons")
	IURRI003 = share.NewMCode("IURRI-003", "Render Isoflow build scene failed")
	IURRI004 = share.NewMCode("IURRI-004", "Render Isoflow with icons failed")
)

func (rcvr *renderUsecase) RenderIsoflow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := rcvr.buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRI003, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	logger.DEBUG(IURRI001, "load isoflow icons")
	icons, _ := rcvr.isoflowRepository.LoadIsoflowIcons(opts.Assets)
	out, err := rcvr.isoflowRepository.RenderWithIcons(scene, icons)
	if err != nil {
		logger.ERROR(IURRI004, "render with icons failed", map[string]any{"error": err})
		return nil, err
	}
	return out, nil
}

func (rcvr *renderUsecase) serviceOptions(opts entity.RenderOptions) ([]entity.ServiceEntry, map[int]string, error) {
	if len(bytes.TrimSpace(opts.ServicesCSV)) == 0 {
		logger.DEBUG(IURSO001, "branch no services csv", map[string]any{"abbreviations": len(opts.Abbreviations)})
		abbreviations, err := v1engine.ResolveServiceOptionsV1EngineOptionService(nil, opts.Abbreviations)
		return nil, abbreviations, err
	}

	logger.DEBUG(IURSO002, "branch services csv", map[string]any{"bytes": len(opts.ServicesCSV)})
	if err := v1engine.ValidateLegendCSVRowsV1EngineOptionService(opts.ServicesCSV); err != nil {
		logger.ERROR(IURSO005, "legend validation failed", map[string]any{"error": err})
		return nil, nil, err
	}
	entries, err := rcvr.xaligoRepository.ReadServiceListFromReader(bytes.NewReader(opts.ServicesCSV))
	if err != nil {
		logger.ERROR(IURSO003, "read services csv failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("read services CSV: %w", err)
	}
	abbreviations, err := v1engine.ResolveServiceOptionsV1EngineOptionService(entries, opts.Abbreviations)
	if err != nil {
		logger.ERROR(IURSO005, "legend validation failed", map[string]any{"error": err})
		return nil, nil, err
	}
	return entries, abbreviations, nil
}

// buildScene is the orchestration boundary around synchronous V1 engine
// stages. Cancellation checks, repository-backed option loading, and stage
// ordering intentionally remain here so future concurrency can be controlled
// by the usecase layer without changing calculation code.
func (rcvr *renderUsecase) buildScene(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, []entity.ServiceEntry, error) {
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}
	if err := v1engine.ValidateRenderOptionsV1EngineOptionRender(opts); err != nil {
		logger.ERROR(IURBS002, "validate render options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	theme, err := v1engine.NormalizeThemeV1EngineThemeApply(opts.Theme)
	if err != nil {
		return nil, nil, err
	}

	doc, err := v1engine.ParseV1EngineParseDocument(bytes.NewReader(input))
	if err != nil {
		logger.ERROR(IURBS007, "parse DSL failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("parse DSL: %w", err)
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	root, err := v1engine.BuildV1EngineLayoutBuild(doc)
	if err != nil {
		logger.ERROR(IURBS008, "build layout failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build layout: %w", err)
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	entries, abbreviations, err := rcvr.serviceOptions(opts)
	if err != nil {
		logger.ERROR(IURBS009, "service options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	connections := v1engine.CollectConnectionNodesV1EngineSceneConnection(doc.Root)
	dependencies := SceneDependencies{
		XaligoRepository:     rcvr.xaligoRepository,
		ExcalidrawRepository: rcvr.excalidrawRepository,
	}.core()

	var scene []byte
	if opts.Assets != nil {
		logger.DEBUG(IURBS004, "branch embedded assets")
		itemSize := opts.Assets.ItemIconSize
		if itemSize <= 0 {
			logger.DEBUG(IURBS005, "branch default embedded item size")
			itemSize = 32
		}
		scene, err = v1engine.BuildJSONWithFSV1EngineSceneBuild(root, opts.Assets.FS, opts.Assets.CatalogCSV, opts.Assets.GroupIconsDir, itemSize, connections, abbreviations, dependencies)
	} else {
		logger.DEBUG(IURBS006, "branch native assets")
		cfg := config.New()
		scene, err = v1engine.BuildJSONV1EngineSceneBuild(root, filepath.Join(cfg.AssetDir_, "Architecture-Group-Icons"), cfg.SvcCatalogCSV, cfg.ProjectRoot, cfg.ItemIconSize, connections, abbreviations, nil, dependencies)
	}
	if err != nil {
		logger.ERROR(IURBS010, "build JSON failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build excalidraw JSON: %w", err)
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	scene, err = v1engine.ApplyThemeJSONV1EngineThemeApply(scene, theme)
	if err != nil {
		logger.ERROR(IURBS011, "apply theme failed", map[string]any{"error": err})
		return nil, nil, err
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}
	return scene, entries, nil
}

func (rcvr *renderUsecase) checkRenderContext(ctx context.Context) error {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURBS001, "context check failed", map[string]any{"error": err})
		return err
	}
	return nil
}

// ValidateRenderOptions is kept as a source-compatible package boundary.
// Deprecated: construct RenderUsecase and call ValidateRenderOptions.
func ValidateRenderOptions(opts entity.RenderOptions) error {
	return v1engine.ValidateRenderOptionsV1EngineOptionRender(opts)
}
