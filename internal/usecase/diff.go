package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/share"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

var (
	IUD001 = share.NewMCode("IUD-001", "Diff context check failed")
	IUD002 = share.NewMCode("IUD-002", "Diff render options validation failed")
	IUD003 = share.NewMCode("IUD-003", "Diff before parse failed")
	IUD004 = share.NewMCode("IUD-004", "Diff after parse failed")
	IUD005 = share.NewMCode("IUD-005", "Diff side render failed")
)

type DiffUsecase interface {
	Diff(context.Context, []byte, []byte, entity.DiffOptions) (entity.DiffResult, error)
}

type diffUsecase struct {
	xaligoRepository repository.XaligoRepository
	sceneRepository  repository.SceneRepository
	svgRepository    repository.SVGRepository
}

func NewDiffUsecase(xaligoRepository repository.XaligoRepository, sceneRepository repository.SceneRepository, svgRepository repository.SVGRepository) DiffUsecase {
	return &diffUsecase{
		xaligoRepository: xaligoRepository, sceneRepository: sceneRepository, svgRepository: svgRepository,
	}
}

func (rcvr *diffUsecase) Diff(ctx context.Context, beforeInput, afterInput []byte, opts entity.DiffOptions) (entity.DiffResult, error) {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IUD001, "context check failed", map[string]any{"error": err})
		return entity.DiffResult{}, err
	}
	renderOptions := entity.RenderOptions{
		Mode: opts.Mode, Format: FormatSVG, Theme: opts.Theme, PxPerInch: opts.PxPerInch, Assets: opts.Assets,
	}
	if err := v1engine.ValidateRenderOptionsV1EngineOptionRender(renderOptions); err != nil {
		logger.ERROR(IUD002, "render options validation failed", map[string]any{"error": err})
		return entity.DiffResult{}, err
	}
	theme, err := v1engine.NormalizeThemeV1EngineThemeApply(renderOptions.Theme)
	if err != nil {
		return entity.DiffResult{}, err
	}
	renderOptions.Theme = theme

	before, err := v1engine.ParseV1EngineParseDocument(bytes.NewReader(beforeInput))
	if err != nil {
		logger.ERROR(IUD003, "before parse failed", map[string]any{"error": err})
		return entity.DiffResult{}, fmt.Errorf("parse before DSL: %w", err)
	}
	after, err := v1engine.ParseV1EngineParseDocument(bytes.NewReader(afterInput))
	if err != nil {
		logger.ERROR(IUD004, "after parse failed", map[string]any{"error": err})
		return entity.DiffResult{}, fmt.Errorf("parse after DSL: %w", err)
	}
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IUD001, "context check failed", map[string]any{"error": err})
		return entity.DiffResult{}, err
	}

	summary := v1engine.DiffDocumentsV1EngineDiffDocument(before, after)
	v1engine.MarkChangesV1EngineDiffDocument(before.Root, summary.Before, string(entity.StructuralChangeRemoved))
	v1engine.MarkChangesV1EngineDiffDocument(after.Root, summary.After, string(entity.StructuralChangeAdded))

	removedImage, err := rcvr.renderDiffDocument(ctx, "before", before, renderOptions)
	if err != nil {
		logger.ERROR(IUD005, "before render failed", map[string]any{"error": err})
		return entity.DiffResult{}, err
	}
	addedImage, err := rcvr.renderDiffDocument(ctx, "after", after, renderOptions)
	if err != nil {
		logger.ERROR(IUD005, "after render failed", map[string]any{"error": err})
		return entity.DiffResult{}, err
	}
	return entity.DiffResult{RemovedImage: removedImage, AddedImage: addedImage, Summary: summary}, nil
}

func (rcvr *diffUsecase) renderDiffDocument(ctx context.Context, side string, document entity.Document, opts entity.RenderOptions) ([]byte, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	root, err := v1engine.BuildV1EngineLayoutBuild(document)
	if err != nil {
		return nil, fmt.Errorf("build %s layout: %w", side, err)
	}
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	dependencies := SceneDependencies{
		XaligoRepository: rcvr.xaligoRepository, SceneRepository: rcvr.sceneRepository,
	}.core()
	connections := v1engine.CollectConnectionNodesV1EngineSceneConnection(document.Root)
	var sceneJSON []byte
	if opts.Assets != nil {
		itemSize := opts.Assets.ItemIconSize
		if itemSize <= 0 {
			itemSize = 32
		}
		sceneJSON, err = v1engine.BuildJSONWithFSV1EngineSceneBuild(root, opts.Assets.FS, opts.Assets.CatalogCSV, opts.Assets.GroupIconsDir, itemSize, connections, nil, dependencies)
	} else {
		cfg := config.New()
		sceneJSON, err = v1engine.BuildJSONV1EngineSceneBuild(root, filepath.Join(cfg.AssetDir_, "Architecture-Group-Icons"), cfg.SvcCatalogCSV, cfg.ProjectRoot, cfg.ItemIconSize, connections, nil, nil, dependencies)
	}
	if err != nil {
		return nil, fmt.Errorf("build %s scene: %w", side, err)
	}
	sceneJSON, err = v1engine.ApplyThemeJSONV1EngineThemeApply(sceneJSON, opts.Theme)
	if err != nil {
		return nil, fmt.Errorf("apply %s theme: %w", side, err)
	}
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		return nil, fmt.Errorf("decode %s scene: %w", side, err)
	}
	plan := v1engine.BuildPlanV1EnginePlanBuild(&scene, v1engine.ResolvePlanOptionsV1EngineOptionPlan(opts, nil))
	image, err := rcvr.svgRepository.Render(plan, opts.PxPerInch, "")
	if err != nil {
		return nil, fmt.Errorf("render %s SVG: %w", side, err)
	}
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return image, nil
}
