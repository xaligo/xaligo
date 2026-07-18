package usecase_test

import (
	"context"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestDiffUsecaseRendersRemovedAndAddedHighlights(t *testing.T) {
	before := []byte(`<frame version="1" width="420" height="180" layout="horizontal"><rectangle id="one" title="One" /><rectangle id="removed" title="Removed" /></frame>`)
	after := []byte(`<frame version="1" width="420" height="180" layout="horizontal"><rectangle id="one" title="One" /><rectangle id="added" title="Added" /></frame>`)

	result, err := newDiffUsecase().Diff(context.Background(), before, after, entity.DiffOptions{Theme: "light", PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	if result.Summary.RemovedCount != 1 || result.Summary.AddedCount != 1 || result.Summary.ModifiedCount != 0 {
		t.Fatalf("summary = %#v", result.Summary)
	}
	removedSVG, addedSVG := string(result.RemovedImage), string(result.AddedImage)
	if !strings.Contains(removedSVG, `fill="#FEE2E2"`) || strings.Contains(removedSVG, `fill="#DCFCE7"`) {
		t.Fatalf("removed SVG does not contain only red highlight: %s", removedSVG)
	}
	if !strings.Contains(addedSVG, `fill="#DCFCE7"`) || strings.Contains(addedSVG, `fill="#FEE2E2"`) {
		t.Fatalf("added SVG does not contain only green highlight: %s", addedSVG)
	}
}

func TestDiffUsecaseMarksBothSidesOfModification(t *testing.T) {
	before := []byte(`<frame version="1" width="320" height="180"><rectangle id="one" title="Old" /></frame>`)
	after := []byte(`<frame version="1" width="320" height="180"><rectangle id="one" title="New" /></frame>`)
	result, err := newDiffUsecase().Diff(context.Background(), before, after, entity.DiffOptions{Theme: "dark"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Summary.ModifiedCount != 1 || !strings.Contains(string(result.RemovedImage), `fill="#FEE2E2"`) || !strings.Contains(string(result.AddedImage), `fill="#DCFCE7"`) {
		t.Fatalf("result = %#v", result.Summary)
	}
	if !strings.Contains(string(result.AddedImage), `fill="#111827"`) {
		t.Fatalf("dark added SVG missing dark background: %s", result.AddedImage)
	}
}

func TestDiffUsecaseRendersFrameMetadataEntryHighlights(t *testing.T) {
	before := []byte(`<frame version="1" width="320" height="180"><metadata><entry key="owner" value="platform" /></metadata></frame>`)
	after := []byte(`<frame version="1" width="320" height="180"><metadata><entry key="owner" value="security" /></metadata></frame>`)

	result, err := newDiffUsecase().Diff(context.Background(), before, after, entity.DiffOptions{Theme: "light", PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	if result.Summary.ModifiedCount != 1 {
		t.Fatalf("summary = %#v", result.Summary)
	}
	if !strings.Contains(string(result.RemovedImage), `fill="#FEE2E2"`) {
		t.Fatalf("removed metadata highlight is missing: %s", result.RemovedImage)
	}
	if !strings.Contains(string(result.AddedImage), `fill="#DCFCE7"`) {
		t.Fatalf("added metadata highlight is missing: %s", result.AddedImage)
	}
}

func TestDiffUsecaseReturnsSideSpecificErrorsAndHonorsContext(t *testing.T) {
	valid := []byte(`<frame version="1" width="320" height="180"><blank /></frame>`)
	if _, err := newDiffUsecase().Diff(context.Background(), valid, []byte(`<frame version="1">`), entity.DiffOptions{}); err == nil || !strings.Contains(err.Error(), "parse after DSL") {
		t.Fatalf("after parse error = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := newDiffUsecase().Diff(ctx, valid, valid, entity.DiffOptions{}); err == nil {
		t.Fatal("canceled Diff() error = nil")
	}
}

func TestDiffUsecaseWithoutChangesProducesUnhighlightedImages(t *testing.T) {
	input := []byte(`<frame version="1" width="320" height="180"><rectangle id="one" /></frame>`)
	result, err := newDiffUsecase().Diff(context.Background(), input, input, entity.DiffOptions{})
	if err != nil {
		t.Fatal(err)
	}
	for _, image := range [][]byte{result.RemovedImage, result.AddedImage} {
		if strings.Contains(string(image), "FEE2E2") || strings.Contains(string(image), "DCFCE7") {
			t.Fatalf("unexpected highlight in %s", image)
		}
	}
}

func newDiffUsecase() usecase.DiffUsecase {
	return usecase.NewDiffUsecase(
		repository.NewXaligoRepository(), repository.NewExcalidrawRepository(), repository.NewSVGRepository(),
	)
}
