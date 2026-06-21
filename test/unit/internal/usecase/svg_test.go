package usecase_test

import (
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

func TestRenderPlanScalesTextWithRequestedPPI(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 2, H: 1, Background: "ffffff"},
		Ops: []entity.DrawOp{{
			Kind: "text", X: 0.25, Y: 0.25, W: 1.5, H: 0.5,
			Text: "A&B <cloud>", FontSize: 12, FontFace: `A"B`,
		}},
	}

	out, err := usecase.RenderSVGPlan(plan, 144)
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
	_, err := usecase.RenderSVGPlan(entity.Plan{}, 96)
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
	out, err := usecase.RenderSVGPlan(plan, 96)
	if err != nil {
		t.Fatal(err)
	}
	svg := string(out)
	if !strings.Contains(svg, `marker-start="url(#xaligo-oval)"`) || !strings.Contains(svg, `marker-end="url(#xaligo-oval)"`) {
		t.Fatalf("circular markers missing:\n%s", svg)
	}
}
