//go:build cgo && xaligo_engine

package integration_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func benchmarkComplexHybridSource(b *testing.B) []byte {
	b.Helper()
	path := filepath.Join("..", "..", "docs", "src", "examples", "samples", "complex-hybrid-architecture.xal")
	source, err := os.ReadFile(path)
	if err != nil {
		b.Fatal(err)
	}
	return source
}

func benchmarkRenderUsecase() usecase.RenderUsecase {
	return usecase.NewRenderUsecase(
		repository.NewSceneRepository(),
		repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(),
		repository.NewSVGRepository(),
	)
}

func BenchmarkComplexHybridProjectAnalyze(b *testing.B) {
	source := benchmarkComplexHybridSource(b)
	project := usecase.NewProjectUsecase(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(source)))
	b.ResetTimer()
	for range b.N {
		if _, err := project.Analyze(context.Background(), "file:///complex-hybrid-architecture.xal", source); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkComplexHybridV1RenderSVG(b *testing.B) {
	source := benchmarkComplexHybridSource(b)
	renderer := benchmarkRenderUsecase()
	options := entity.RenderOptions{
		Format:        usecase.FormatSVG,
		Theme:         "light",
		PxPerInch:     96,
		Mode:          usecase.ModeNetwork,
		CombineFrames: true,
	}
	b.ReportAllocs()
	b.SetBytes(int64(len(source)))
	b.ResetTimer()
	for range b.N {
		if _, err := renderer.RenderSVG(context.Background(), source, options); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkComplexHybridScaleRustV2Resolve(b *testing.B) {
	spec := benchmarkComplexHybridEngineSpec(b)
	engine := v2.NewEngineUsecase()
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, err := engine.Resolve(context.Background(), spec); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkComplexHybridScaleRustV2RenderSVG(b *testing.B) {
	spec := benchmarkComplexHybridEngineSpec(b)
	engine := v2.NewEngineUsecase()
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		if _, err := engine.RenderSVG(context.Background(), spec); err != nil {
			b.Fatal(err)
		}
	}
}

// benchmarkComplexHybridEngineSpec preserves the sample's generic concept
// count, hierarchy, and line count. It is intentionally not a V1-to-V2 visual
// compatibility adapter; that frontend remains roadmap work.
func benchmarkComplexHybridEngineSpec(b *testing.B) entity.EngineDocumentSpec {
	b.Helper()
	source := benchmarkComplexHybridSource(b)
	analysis, err := usecase.NewProjectUsecase(nil).Analyze(
		context.Background(),
		"file:///complex-hybrid-architecture.xal",
		source,
	)
	if err != nil {
		b.Fatal(err)
	}
	if len(analysis.Diagnostics) != 0 {
		b.Fatalf("sample diagnostics = %#v", analysis.Diagnostics)
	}
	children := make(map[int][]int, len(analysis.Symbols))
	for index, symbol := range analysis.Symbols {
		children[symbol.ParentOrdinal] = append(children[symbol.ParentOrdinal], index)
	}
	endpointIDs := make([]string, 0, len(analysis.Symbols))
	for index, symbol := range analysis.Symbols {
		if matchesEngineEndpoint(symbol.Concept) {
			endpointIDs = append(endpointIDs, benchmarkConceptID(symbol.Concept, index))
		}
	}
	if len(endpointIDs) < 2 {
		b.Fatal("sample requires at least two generic endpoints")
	}
	var lower func(int) entity.EngineElementSpec
	lower = func(index int) entity.EngineElementSpec {
		symbol := analysis.Symbols[index]
		element := entity.EngineElementSpec{
			ID:       benchmarkConceptID(symbol.Concept, index),
			Concept:  symbol.Concept,
			Overflow: entity.EngineOverflowVisible,
		}
		switch symbol.Concept {
		case entity.EngineConceptFrame:
			element.Width = engineFloat(1920)
			element.Height = engineFloat(1440)
			element.Layout = entity.EngineLayoutVertical
		case entity.EngineConceptGroup, entity.EngineConceptCapture:
			element.Weight = engineFloat(1)
			element.Layout = entity.EngineLayoutVertical
		case entity.EngineConceptPort:
			element.Port = &entity.EnginePortSpec{
				Side:   entity.EngineSideRight,
				Anchor: engineFloat(float64(index%5) / 4),
			}
		case entity.EngineConceptLine:
			element.Line = &entity.EngineLineSpec{
				Source:           endpointIDs[index%len(endpointIDs)],
				Target:           endpointIDs[(index+len(endpointIDs)/2)%len(endpointIDs)],
				Routing:          entity.EngineRoutingOrthogonal,
				TargetDecoration: entity.EngineDecorationArrow,
			}
		case entity.EngineConceptText:
			element.Text = &entity.EngineTextSpec{Value: symbol.Name}
		}
		for _, child := range children[index] {
			element.Children = append(element.Children, lower(child))
		}
		return element
	}
	roots := make([]entity.EngineElementSpec, 0, len(children[-1]))
	for _, index := range children[-1] {
		roots = append(roots, lower(index))
	}
	return entity.EngineDocumentSpec{
		Layout:   entity.EngineLayoutAbsolute,
		Width:    1920,
		Height:   1440,
		Overflow: entity.EngineOverflowVisible,
		Elements: roots,
	}
}

func matchesEngineEndpoint(concept entity.EngineConcept) bool {
	return concept == entity.EngineConceptGroup ||
		concept == entity.EngineConceptCapture ||
		concept == entity.EngineConceptItem ||
		concept == entity.EngineConceptPort
}

func benchmarkConceptID(concept entity.EngineConcept, ordinal int) string {
	return fmt.Sprintf("%s-%d", concept, ordinal)
}
