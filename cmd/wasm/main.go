//go:build js

// Package main adapts the shared render use case to synchronous JavaScript globals.
// Parsing, layout, validation, and rendering remain owned by internal/usecase.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"syscall/js"

	awsassets "github.com/ryo-arima/xaligo/etc/resources/aws"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

var embeddedAssets = &entity.AssetSource{
	FS: awsassets.Assets, CatalogCSV: awsassets.CatalogCSV,
	GroupIconsDir: awsassets.GroupIconsDir, IsoflowIconsJSON: awsassets.IsoflowIconsJSON,
	ItemIconSize: 48,
}

var xaligoUsecase = usecase.NewXaligoUsecase(
	repository.NewExcalidrawRepository(),
	repository.NewXaligoRepository(),
	repository.NewPowerpointRepository(),
	repository.NewIsoflowRepository(),
	repository.NewSVGRepository(),
	repository.NewXYFlowRepository(),
)

func main() {
	js.Global().Set("xaligoRender", js.FuncOf(jsRender))
	js.Global().Set("xaligoRenderWithServices", js.FuncOf(jsRenderWithServices))
	js.Global().Set("xaligoBuildPptxPlan", js.FuncOf(jsBuildPptxPlan))
	js.Global().Set("xaligoDiagnose", js.FuncOf(jsDiagnose))
	js.Global().Set("xaligoRenderXYFlow", js.FuncOf(jsRenderXYFlow))
	js.Global().Set("xaligoRenderIsoflow", js.FuncOf(jsRenderIsoflow))
	<-make(chan struct{})
}

func jsRender(_ js.Value, args []js.Value) any {
	return renderResult("xaligoRender", args, usecase.FormatExcalidraw, nil)
}

func jsRenderWithServices(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return jsResult(nil, fmt.Errorf("xaligoRenderWithServices: expected 2 arguments (xal, servicesCsv)"))
	}
	return renderResult("xaligoRenderWithServices", args, usecase.FormatExcalidraw, []byte(args[1].String()))
}

func jsRenderXYFlow(_ js.Value, args []js.Value) any {
	return renderResult("xaligoRenderXYFlow", args, usecase.FormatXYFlow, nil)
}

func jsRenderIsoflow(_ js.Value, args []js.Value) any {
	return renderResult("xaligoRenderIsoflow", args, usecase.FormatIsoflow, nil)
}

func renderResult(name string, args []js.Value, format entity.Format, servicesCSV []byte) any {
	if len(args) < 1 {
		return jsResult(nil, fmt.Errorf("%s: expected 1 argument (xal)", name))
	}
	out, err := xaligoUsecase.Render(context.Background(), []byte(args[0].String()), entity.RenderOptions{
		Format: format, ServicesCSV: servicesCSV, Assets: embeddedAssets,
	})
	return jsResult(out, err)
}

func jsDiagnose(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return jsResult(nil, fmt.Errorf("xaligoDiagnose: expected 1 argument (xal)"))
	}
	diagnostics, err := usecase.Diagnose(context.Background(), []byte(args[0].String()))
	if err != nil {
		return jsResult(nil, err)
	}
	encoded, err := json.Marshal(diagnostics)
	return jsResult(encoded, err)
}

func jsBuildPptxPlan(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return jsResult(nil, fmt.Errorf("xaligoBuildPptxPlan: expected at least 1 argument (xal)"))
	}
	opts := entity.RenderOptions{Assets: embeddedAssets}
	if len(args) >= 2 && !args[1].IsUndefined() && !args[1].IsNull() {
		opts.ServicesCSV = []byte(args[1].String())
	}
	if len(args) >= 3 && !args[2].IsUndefined() && !args[2].IsNull() && args[2].String() != "" {
		if err := json.Unmarshal([]byte(args[2].String()), &opts); err != nil {
			return jsResult(nil, fmt.Errorf("parse options: %w", err))
		}
		opts.Assets = embeddedAssets
	}
	out, err := xaligoUsecase.BuildPPTXPlan(context.Background(), []byte(args[0].String()), opts)
	return jsResult(out, err)
}

func jsResult(result []byte, err error) any {
	if err != nil {
		return map[string]any{"error": err.Error()}
	}
	return map[string]any{"result": string(result)}
}
