package repository_test

import (
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestRenderPlanScalesTextWithRequestedPPI(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 2, H: 1, Background: "ffffff"},
		Ops: []entity.DrawOp{{
			Kind: "text", X: 0.25, Y: 0.25, W: 1.5, H: 0.5,
			Text: "A&B <cloud>", FontSize: 12, FontFace: `A"B`,
		}},
	}

	out, err := repository.NewSVGRepository().Render(plan, 144, "")
	if err != nil {
		t.Fatalf("RenderSVGPlan() error = %v", err)
	}
	svg := string(out)
	for _, want := range []string{
		`width="288" height="144"`,
		`font-size="24"`,
		`font-family="A&#34;B"`,
		`A&amp;B &lt;cloud&gt;`,
	} {
		if !strings.Contains(svg, want) {
			t.Errorf("SVG missing %q:\n%s", want, svg)
		}
	}
}

func TestRenderPlanRejectsInvalidSlideSize(t *testing.T) {
	_, err := repository.NewSVGRepository().Render(entity.Plan{}, 96, "")
	if err == nil {
		t.Fatal("RenderSVGPlan() error = nil, want invalid slide size error")
	}
}

func TestRenderPlanUsesCircularRouteMarkers(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 2, H: 1, Background: "ffffff"},
		Ops: []entity.DrawOp{{
			Kind: "line", X: 0.25, Y: 0.5, W: 1.5,
			Line: &entity.LineStyle{Color: "64748B", Width: 1, BeginArrowType: "oval", EndArrowType: "oval"},
		}},
	}
	out, err := repository.NewSVGRepository().Render(plan, 96, "")
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	if !strings.Contains(svg, `marker-start="url(#xaligo-oval)"`) || !strings.Contains(svg, `marker-end="url(#xaligo-oval)"`) {
		t.Fatalf("circular markers missing:\n%s", svg)
	}
}

func TestRenderPlanDrawsServiceLegendAtRequestedPosition(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 2, H: 1, Background: "ffffff"},
		Legend: []entity.LegendEntry{{
			CatalogID:    27,
			Abbreviation: "EC2",
			OfficialName: "Amazon EC2",
			Data:         "data:image/svg+xml;base64,QQ==",
		}},
	}

	out, err := repository.NewSVGRepository().Render(plan, 96, "left")
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	for _, want := range []string{
		`id="xaligo-svg-legend"`,
		`width="496" height="96"`,
		`transform="translate(304 0)"`,
		`EC2`,
		`Amazon EC2`,
	} {
		if !strings.Contains(svg, want) {
			t.Fatalf("SVG missing %q:\n%s", want, svg)
		}
	}
}

func TestRenderPlanKeepsLayoutFontSizeAtNonDefaultPPI(t *testing.T) {
	fontSize := 12.0
	scene := entity.PresentationScene{Elements: []entity.Element{
		{ID: "paper-frame", Type: "frame", Width: 200, Height: 100},
		{ID: "text", Type: "text", X: 10, Y: 10, Width: 100, Height: 30, Text: "PPI", FontSize: &fontSize},
	}}
	plan := usecase.BuildPlan(&scene, entity.PlanOptions{PxPerInch: 144})

	out, err := repository.NewSVGRepository().Render(plan, 144, "")
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	for _, want := range []string{`width="200" height="100"`, `font-size="12"`} {
		if !strings.Contains(svg, want) {
			t.Fatalf("SVG missing %q:\n%s", want, svg)
		}
	}
}

func TestRenderPlanKeepsLayoutStrokeWidthAtNonDefaultPPI(t *testing.T) {
	opacity := 100.0
	scene := entity.PresentationScene{Elements: []entity.Element{
		{ID: "paper-frame", Type: "frame", Width: 200, Height: 100},
		{ID: "box", Type: "rectangle", X: 10, Y: 10, Width: 100, Height: 50, StrokeColor: "#1e1e1e", StrokeWidth: 2, Opacity: &opacity},
	}}
	plan := usecase.BuildPlan(&scene, entity.PlanOptions{PxPerInch: 144})

	out, err := repository.NewSVGRepository().Render(plan, 144, "")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `stroke-width="2"`) {
		t.Fatalf("SVG did not restore the 2px layout stroke at 144 PPI:\n%s", out)
	}
}

func TestRenderPlanIncludesVisibleTextGlyphsInCanvasBounds(t *testing.T) {
	textOp := entity.DrawOp{
		ID: "visible", Kind: "text", X: 0.8, Y: 0.1, W: 0.1, H: 0.2,
		Text: "visible overflow", FontSize: 12, FontFace: "Helvetica",
		TextLayout: &entity.TextLayout{
			Role: entity.TextRoleLabel, Wrap: false, Fit: entity.TextFitNone,
			Overflow: entity.TextOverflowVisible, Clip: false, LineHeight: 1.2,
		},
	}
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 1, H: 0.5, Background: "ffffff"},
		Ops:   []entity.DrawOp{textOp},
	}
	out, err := repository.NewSVGRepository().Render(plan, 96, "")
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	match := regexp.MustCompile(`<svg[^>]* width="([0-9.]+)"`).FindStringSubmatch(svg)
	if len(match) != 2 {
		t.Fatalf("SVG canvas width not found:\n%s", svg)
	}
	width, err := strconv.ParseFloat(match[1], 64)
	if err != nil || width <= 96 {
		t.Fatalf("visible glyph overflow was cropped: canvas width=%q\n%s", match[1], svg)
	}
	if strings.Contains(svg, `<clipPath`) {
		t.Fatalf("visible text unexpectedly received a clip path:\n%s", svg)
	}

	clipped := textOp
	clipped.ID = "clipped"
	clipped.TextLayout = &entity.TextLayout{
		Role: entity.TextRoleLabel, Wrap: false, Fit: entity.TextFitNone,
		Overflow: entity.TextOverflowClip, Clip: true, LineHeight: 1.2,
	}
	plan.Ops = []entity.DrawOp{clipped}
	out, err = repository.NewSVGRepository().Render(plan, 96, "")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), `<svg xmlns="http://www.w3.org/2000/svg" width="96"`) {
		t.Fatalf("clipped text should not expand the canvas:\n%s", out)
	}
}

func TestRenderPlanWrapsShrinksAndClipsTextFromSharedContract(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 1, H: 0.6, Background: "ffffff"},
		Ops: []entity.DrawOp{
			{
				ID: "long-label", Kind: "text", X: 0.05, Y: 0.05, W: 0.3, H: 0.15,
				Text: "Excalidraw / XYFlow / Isoflow", FontSize: 12,
				TextLayout: &entity.TextLayout{
					Role: entity.TextRoleLabel, Wrap: false, Fit: entity.TextFitShrink,
					Overflow: entity.TextOverflowClip, Clip: true, LineHeight: 1.2,
				},
			},
			{
				ID: "wrapped-label", Kind: "text", X: 0.4, Y: 0.05, W: 0.5, H: 0.5,
				Text: "alpha beta gamma", FontSize: 9,
				TextLayout: &entity.TextLayout{
					Role: entity.TextRoleLabel, Wrap: true, Fit: entity.TextFitNone,
					Overflow: entity.TextOverflowClip, Clip: true, LineHeight: 1.2,
				},
			},
			{
				ID: "visible-label", Kind: "text", X: 0.05, Y: 0.4, W: 0.2, H: 0.1,
				Text: "visible overflow", FontSize: 9,
				TextLayout: &entity.TextLayout{
					Role: entity.TextRoleLabel, Wrap: false, Fit: entity.TextFitNone,
					Overflow: entity.TextOverflowVisible, Clip: true, LineHeight: 1.2,
				},
			},
		},
	}

	out, err := repository.NewSVGRepository().Render(plan, 96, "")
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	if strings.Count(svg, `<clipPath id="xaligo-text-clip-`) != 2 {
		t.Fatalf("text clip paths missing:\n%s", svg)
	}
	if strings.Count(svg, "<tspan") < 3 {
		t.Fatalf("wrapped tspans missing:\n%s", svg)
	}
	match := regexp.MustCompile(`id="xaligo-text-clip-[^"]+"[\s\S]*?<text[^>]*font-size="([0-9.]+)"`).FindStringSubmatch(svg)
	if len(match) != 2 {
		t.Fatalf("shrunk font size not found:\n%s", svg)
	}
	got, err := strconv.ParseFloat(match[1], 64)
	if err != nil || got >= 16 {
		t.Fatalf("long label font size = %q, want smaller than 16px", match[1])
	}
}
