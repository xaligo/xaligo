//go:build js

// Package main is the WebAssembly entry point for xaligo.
// Build with:
//
//	GOOS=js GOARCH=wasm go build -o xaligo.wasm ./cmd/wasm
//
// The resulting xaligo.wasm exposes the following functions on the global JS object:
//
//	xaligoRender(xal: string): { result?: string; error?: string }
//	  Converts a .xal DSL string into Excalidraw JSON.
//	  Uses the embedded service-catalog.csv and Architecture-Group-Icons assets.
//
//	xaligoRenderWithServices(xal: string, servicesCsv: string): { result?: string; error?: string }
//	  Same as xaligoRender but also parses a services.csv string and adds the
//	  service legend sidebar (abbreviation overrides from the CSV are applied).
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	xaligoapi "github.com/ryo-arima/xaligo"
	awsassets "github.com/ryo-arima/xaligo/etc/resources/aws"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/excalidraw"
	"github.com/ryo-arima/xaligo/internal/layout"
	"github.com/ryo-arima/xaligo/internal/model"
	"github.com/ryo-arima/xaligo/internal/parser"
	"github.com/ryo-arima/xaligo/internal/pptxplan"
	"github.com/ryo-arima/xaligo/internal/repository"
	xyflowrenderer "github.com/ryo-arima/xaligo/internal/xyflow"
)

func main() {
	js.Global().Set("xaligoRender", js.FuncOf(jsRender))
	js.Global().Set("xaligoRenderWithServices", js.FuncOf(jsRenderWithServices))
	js.Global().Set("xaligoBuildPptxPlan", js.FuncOf(jsBuildPptxPlan))
	js.Global().Set("xaligoDiagnose", js.FuncOf(jsDiagnose))
	js.Global().Set("xaligoRenderXYFlow", js.FuncOf(jsRenderXYFlow))

	// Keep the WASM module alive until the page unloads.
	<-make(chan struct{})
}

func jsRenderXYFlow(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return jsResult("", fmt.Errorf("xaligoRenderXYFlow: expected 1 argument (xal)"))
	}
	sceneJSON, err := renderXAL(args[0].String(), nil)
	if err != nil {
		return jsResult("", err)
	}
	out, err := xyflowrenderer.Render([]byte(sceneJSON))
	if err != nil {
		return jsResult("", err)
	}
	return jsResult(string(out), nil)
}

func jsDiagnose(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return jsResult("", fmt.Errorf("xaligoDiagnose: expected 1 argument (xal)"))
	}
	diagnostics, err := xaligoapi.Diagnose(context.Background(), []byte(args[0].String()))
	if err != nil {
		return jsResult("", err)
	}
	encoded, err := json.Marshal(diagnostics)
	if err != nil {
		return jsResult("", fmt.Errorf("encode diagnostics: %w", err))
	}
	return jsResult(string(encoded), nil)
}

// jsResult returns { result, error } objects back to JavaScript.
func jsResult(result string, err error) any {
	if err != nil {
		return map[string]any{"error": err.Error()}
	}
	return map[string]any{"result": result}
}

// renderXAL is the core conversion logic shared by both exported functions.
func renderXAL(xalSrc string, abbrevMap map[int]string) (string, error) {
	doc, err := parser.Parse(strings.NewReader(xalSrc))
	if err != nil {
		return "", fmt.Errorf("parse DSL: %w", err)
	}

	root, err := layout.Build(doc)
	if err != nil {
		return "", fmt.Errorf("build layout: %w", err)
	}

	var connections []*model.Node
	for _, child := range doc.Root.Children {
		if child.Tag == "connection" {
			connections = append(connections, child)
		}
	}

	out, err := excalidraw.BuildJSONWithFS(
		root,
		awsassets.Assets,
		awsassets.CatalogCSV,
		awsassets.GroupIconsDir,
		48.0,
		connections,
		abbrevMap,
	)
	if err != nil {
		return "", fmt.Errorf("build excalidraw: %w", err)
	}
	return string(out), nil
}

// jsRender handles xaligoRender(xal).
func jsRender(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return jsResult("", fmt.Errorf("xaligoRender: expected 1 argument (xal string)"))
	}
	result, err := renderXAL(args[0].String(), nil)
	return jsResult(result, err)
}

// jsRenderWithServices handles xaligoRenderWithServices(xal, servicesCsv).
// servicesCsv is the text content of a services.csv file (same format used by
// the --services flag of the CLI command `xaligo generate excalidraw`).
func jsRenderWithServices(_ js.Value, args []js.Value) any {
	if len(args) < 2 {
		return jsResult("", fmt.Errorf("xaligoRenderWithServices: expected 2 arguments (xal, servicesCsv)"))
	}
	xal := args[0].String()
	csvContent := args[1].String()

	_, abbrevMap, err := parseServicesCsv(csvContent)
	if err != nil {
		return jsResult("", fmt.Errorf("parse servicesCsv: %w", err))
	}

	result, err := renderXAL(xal, abbrevMap)
	return jsResult(result, err)
}

// jsBuildPptxPlan handles xaligoBuildPptxPlan(xal, servicesCsv, optionsJson).
//
// It renders the .xal to an Excalidraw scene (applying services.csv overrides
// when provided) and then builds a fully-resolved PPTX draw plan in Go — every
// geometry calculation (bounds, paper scaling, routing, anchoring, colour and
// coordinate conversion) happens here, so the JS side only issues PptxGenJS
// drawing calls. servicesCsv may be empty. optionsJson is a JSON object of
// pptxplan.Options (may be empty/"" for defaults).
//
// Returns { result: <plan JSON> } or { error }.
func jsBuildPptxPlan(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return jsResult("", fmt.Errorf("xaligoBuildPptxPlan: expected at least 1 argument (xal)"))
	}
	xal := args[0].String()

	var entries []entity.ServiceEntry
	var abbrevMap map[int]string
	if len(args) >= 2 && !args[1].IsUndefined() && !args[1].IsNull() {
		if csv := args[1].String(); strings.TrimSpace(csv) != "" {
			parsedEntries, m, err := parseServicesCsv(csv)
			if err != nil {
				return jsResult("", fmt.Errorf("parse servicesCsv: %w", err))
			}
			entries = parsedEntries
			abbrevMap = m
		}
	}

	var opts pptxplan.Options
	if len(args) >= 3 && !args[2].IsUndefined() && !args[2].IsNull() {
		if optJSON := args[2].String(); strings.TrimSpace(optJSON) != "" {
			if err := json.Unmarshal([]byte(optJSON), &opts); err != nil {
				return jsResult("", fmt.Errorf("parse options: %w", err))
			}
		}
	}

	sceneJSON, err := renderXAL(xal, abbrevMap)
	if err != nil {
		return jsResult("", err)
	}
	themedSceneJSON, err := excalidraw.ApplyThemeJSON([]byte(sceneJSON), opts.Theme)
	if err != nil {
		return jsResult("", err)
	}
	opts.LegendEntries = legendEntries(entries)
	planJSON, err := pptxplan.BuildPlanJSON(string(themedSceneJSON), opts)
	if err != nil {
		return jsResult("", fmt.Errorf("build pptx plan: %w", err))
	}
	return jsResult(string(planJSON), nil)
}

// parseServicesCsv parses the in-memory content of a services.csv into a
// catalog-ID → abbreviation map (same format as repository.ReadServiceList).
func parseServicesCsv(content string) ([]entity.ServiceEntry, map[int]string, error) {
	entries, err := repository.ReadServiceListFromReader(strings.NewReader(content))
	if err != nil {
		return nil, nil, err
	}
	m := make(map[int]string, len(entries))
	for _, e := range entries {
		if e.CatalogID > 0 && e.Abbreviation != "" {
			m[e.CatalogID] = e.Abbreviation
		}
	}
	return entries, m, nil
}

func legendEntries(entries []entity.ServiceEntry) []pptxplan.LegendEntry {
	out := make([]pptxplan.LegendEntry, 0, len(entries))
	for _, e := range entries {
		if e.CatalogID <= 0 || e.OfficialName == "" {
			continue
		}
		out = append(out, pptxplan.LegendEntry{
			CatalogID:    e.CatalogID,
			Abbreviation: e.Abbreviation,
			OfficialName: e.OfficialName,
		})
	}
	return out
}
