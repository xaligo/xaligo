package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
)

var (
	logger  = share.DefaultLogger()
	IURR001 = share.NewMCode("IURR-001", "Render context check failed")
	IURR002 = share.NewMCode("IURR-002", "Render validate render options failed")
	IURR003 = share.NewMCode("IURR-003", "Render default format branch")
	IURR004 = share.NewMCode("IURR-004", "Render excalidraw branch")
	IURR005 = share.NewMCode("IURR-005", "Render SVG branch")
	IURR006 = share.NewMCode("IURR-006", "Render PPTX branch")
	IURR007 = share.NewMCode("IURR-007", "Render XYFlow branch")
	IURR008 = share.NewMCode("IURR-008", "Render Isoflow branch")
	IURR009 = share.NewMCode("IURR-009", "Render unknown format branch")
)

func (rcvr *xaligoUsecase) render(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
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

func (rcvr *xaligoUsecase) renderExcalidraw(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
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

func (rcvr *xaligoUsecase) renderSVG(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	planJSON, err := rcvr.BuildPPTXPlan(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRS001, "build plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("build SVG plan: %w", err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		logger.ERROR(IURRS002, "decode plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("decode SVG plan: %w", err)
	}
	return rcvr.svgRepository.Render(plan, opts.PxPerInch)
}

var (
	IURBPP001 = share.NewMCode("IURBPP-001", "Build PPTX plan build scene failed")
	IURBPP002 = share.NewMCode("IURBPP-002", "Build PPTX plan build plan failed")
	IURRP001  = share.NewMCode("IURRP-001", "Render PPTX build plan failed")
	IURRP003  = share.NewMCode("IURRP-003", "Render PPTX export failed")
	IURRP004  = share.NewMCode("IURRP-004", "Render PPTX context check failed")
)

func (rcvr *xaligoUsecase) buildPPTXPlan(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, entries, err := rcvr.buildScene(ctx, input, opts)
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

func (rcvr *xaligoUsecase) renderPPTX(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	planJSON, err := rcvr.BuildPPTXPlan(ctx, input, opts)
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

func (rcvr *xaligoUsecase) renderXYFlow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
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

func (rcvr *xaligoUsecase) renderIsoflow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
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
