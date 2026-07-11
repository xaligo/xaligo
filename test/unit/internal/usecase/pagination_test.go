package usecase_test

import (
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestSplitPlanPagesAssignsOpsToIntersectingTiles(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 20, H: 10, Background: "FFFFFF"},
		Ops: []entity.DrawOp{
			{ID: "left", Kind: "rect", X: 1, Y: 1, W: 2, H: 2},
			{ID: "right", Kind: "rect", X: 12, Y: 1, W: 2, H: 2},
			{ID: "line", Kind: "line", X: 5, Y: 5, Points: []entity.PtIn{{X: 0, Y: 0}, {X: 10, Y: 0}}},
		},
	}

	split := usecase.SplitPlanPages(plan, entity.PageSplitOptions{PageW: 10, PageH: 10})
	if split.Rows != 1 || split.Cols != 2 || len(split.Pages) != 2 {
		t.Fatalf("split grid = %dx%d pages=%d", split.Rows, split.Cols, len(split.Pages))
	}
	if got := split.Pages[0].OpIDs; !slices.Contains(got, "left") || !slices.Contains(got, "line") || slices.Contains(got, "right") {
		t.Fatalf("page 1 op IDs = %#v", got)
	}
	if got := split.Pages[1].OpIDs; !slices.Contains(got, "right") || !slices.Contains(got, "line") || slices.Contains(got, "left") {
		t.Fatalf("page 2 op IDs = %#v", got)
	}
}

func TestSplitPlanPagesAssignsStableUniqueIDsToAnonymousOps(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 10, H: 10},
		Ops: []entity.DrawOp{
			{Kind: "rect", X: 1, Y: 1, W: 1, H: 1},
			{Kind: "rect", X: 3, Y: 1, W: 1, H: 1},
			{ID: "named", Kind: "rect", X: 5, Y: 1, W: 1, H: 1},
			{ID: "named", Kind: "rect", X: 7, Y: 1, W: 1, H: 1},
		},
	}
	split, err := usecase.SplitPlanPagesChecked(plan, entity.PageSplitOptions{})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"op-1", "op-2", "named", "named#2"}
	if got := split.Pages[0].OpIDs; !slices.Equal(got, want) {
		t.Fatalf("op IDs = %#v, want %#v", got, want)
	}
}

func TestSplitPlanPagesCheckedRejectsNonFiniteGeometry(t *testing.T) {
	cases := []struct {
		name string
		plan entity.Plan
		opts entity.PageSplitOptions
		want string
	}{
		{"slide", entity.Plan{Slide: entity.PlanSlide{W: math.NaN(), H: 10}}, entity.PageSplitOptions{}, "slide size"},
		{"page", entity.Plan{Slide: entity.PlanSlide{W: 10, H: 10}}, entity.PageSplitOptions{PageW: math.Inf(1)}, "page width"},
		{"overlap", entity.Plan{Slide: entity.PlanSlide{W: 10, H: 10}}, entity.PageSplitOptions{Overlap: 10}, "overlap must be smaller"},
		{"op", entity.Plan{Slide: entity.PlanSlide{W: 10, H: 10}, Ops: []entity.DrawOp{{Kind: "rect", X: math.NaN(), W: 1, H: 1}}}, entity.PageSplitOptions{}, "draw op 1 x"},
		{"overflowing bounds", entity.Plan{Slide: entity.PlanSlide{W: 10, H: 10}, Ops: []entity.DrawOp{{Kind: "rect", X: math.MaxFloat64, W: math.MaxFloat64, H: 1}}}, entity.PageSplitOptions{}, "bounds must be finite"},
		{"too many pages", entity.Plan{Slide: entity.PlanSlide{W: 20000, H: 10}}, entity.PageSplitOptions{PageW: 1, PageH: 10}, "tile count is too large"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := usecase.SplitPlanPagesChecked(tc.plan, tc.opts)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v, want %q", err, tc.want)
			}
		})
	}
}

func TestSplitPlanPagesUsesOverlapAndClampsLastTile(t *testing.T) {
	plan := entity.Plan{Slide: entity.PlanSlide{W: 20, H: 6, Background: "FFFFFF"}}

	split := usecase.SplitPlanPages(plan, entity.PageSplitOptions{PageW: 8, PageH: 6, Overlap: 2})
	if split.Rows != 1 || split.Cols != 3 {
		t.Fatalf("split grid = %dx%d", split.Rows, split.Cols)
	}
	if split.Pages[0].Rect.X != 0 || split.Pages[1].Rect.X != 6 || split.Pages[2].Rect.X != 12 {
		t.Fatalf("page x positions = %.1f %.1f %.1f", split.Pages[0].Rect.X, split.Pages[1].Rect.X, split.Pages[2].Rect.X)
	}
	if right := split.Pages[2].Rect.X + split.Pages[2].Rect.W; right != 20 {
		t.Fatalf("last page right edge = %.1f, want 20", right)
	}
}

func TestSplitPlanPagesExpandsContentBeyondSlide(t *testing.T) {
	plan := entity.Plan{
		Slide: entity.PlanSlide{W: 10, H: 10, Background: "FFFFFF"},
		Ops: []entity.DrawOp{
			{ID: "outside", Kind: "rect", X: 11, Y: 2, W: 4, H: 3},
		},
	}

	split := usecase.SplitPlanPages(plan, entity.PageSplitOptions{PageW: 10, PageH: 10})
	if split.Content.W != 15 || split.Cols != 2 {
		t.Fatalf("content width/grid = %.1f/%d", split.Content.W, split.Cols)
	}
	if got := split.Pages[1].OpIDs; !slices.Contains(got, "outside") {
		t.Fatalf("outside op not assigned to expanded page: %#v", got)
	}
}
