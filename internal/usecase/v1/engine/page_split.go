package engine

import (
	"fmt"
	"math"
	"strconv"

	"github.com/xaligo/xaligo/internal/entity"
)

// SplitPlanPagesV1EnginePageSplit calculates page-sized tiles for a resolved draw plan. It does
// not mutate the plan or remove draw operations; renderers can use the returned
// metadata to crop/export pages consistently.
func SplitPlanPagesV1EnginePageSplit(plan entity.Plan, opts entity.PageSplitOptions) entity.PageSplit {
	split, err := SplitPlanPagesCheckedV1EnginePageSplit(plan, opts)
	if err == nil {
		return split
	}
	// Compatibility API: never expose NaN/Inf to callers that cannot receive an
	// error. New integrations should use SplitPlanPagesChecked.
	rect := entity.PageRect{W: 0.01, H: 0.01}
	return entity.PageSplit{Content: rect, PageW: rect.W, PageH: rect.H, Rows: 1, Cols: 1, Pages: []entity.PlanPage{{Index: 1, Row: 1, Col: 1, Rect: rect}}}
}

// SplitPlanPagesCheckedV1EnginePageSplit validates all page and plan geometry before tiling.
func SplitPlanPagesCheckedV1EnginePageSplit(plan entity.Plan, opts entity.PageSplitOptions) (entity.PageSplit, error) {
	if err := validatePageSplitInputV1EnginePageSplit(plan, opts); err != nil {
		return entity.PageSplit{}, err
	}
	content := planContentRectV1EnginePageSplit(plan)
	pageW, pageH := opts.PageW, opts.PageH
	if pageW <= 0 {
		pageW = content.W
	}
	if pageH <= 0 {
		pageH = content.H
	}
	pageW = math.Max(pageW, 0.01)
	pageH = math.Max(pageH, 0.01)
	overlap := math.Max(0, opts.Overlap)
	overlap = math.Min(overlap, math.Min(pageW, pageH)-0.01)
	if overlap < 0 {
		overlap = 0
	}

	stepX := pageW - overlap
	stepY := pageH - overlap
	cols, err := pageCountV1EnginePageSplit(content.W, pageW, stepX)
	if err != nil {
		return entity.PageSplit{}, fmt.Errorf("page columns: %w", err)
	}
	rows, err := pageCountV1EnginePageSplit(content.H, pageH, stepY)
	if err != nil {
		return entity.PageSplit{}, fmt.Errorf("page rows: %w", err)
	}
	const maxPageTiles = 10000
	if rows > maxPageTiles/cols {
		return entity.PageSplit{}, fmt.Errorf("page split exceeds the %d tile limit", maxPageTiles)
	}
	opBounds := planOpBoundsV1EnginePageSplit(plan.Ops)

	pages := make([]entity.PlanPage, 0, rows*cols)
	index := 1
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			rect := entity.PageRect{
				X: content.X + float64(col)*stepX,
				Y: content.Y + float64(row)*stepY,
				W: pageW,
				H: pageH,
			}
			if right := content.X + content.W; rect.X+rect.W > right {
				rect.X = math.Max(content.X, right-rect.W)
			}
			if bottom := content.Y + content.H; rect.Y+rect.H > bottom {
				rect.Y = math.Max(content.Y, bottom-rect.H)
			}
			page := entity.PlanPage{Index: index, Row: row + 1, Col: col + 1, Rect: rect}
			for _, op := range opBounds {
				if rectIntersectsV1EnginePageSplit(rect, op.rect) {
					page.OpIDs = append(page.OpIDs, op.id)
				}
			}
			pages = append(pages, page)
			index++
		}
	}

	return entity.PageSplit{
		Content: content,
		PageW:   pageW,
		PageH:   pageH,
		Overlap: overlap,
		Rows:    rows,
		Cols:    cols,
		Pages:   pages,
	}, nil
}

func validatePageSplitInputV1EnginePageSplit(plan entity.Plan, opts entity.PageSplitOptions) error {
	finite := func(value float64) bool { return !math.IsNaN(value) && !math.IsInf(value, 0) }
	if !finite(plan.Slide.W) || !finite(plan.Slide.H) || plan.Slide.W <= 0 || plan.Slide.H <= 0 {
		return fmt.Errorf("page split slide size must be positive and finite")
	}
	for name, value := range map[string]float64{"page width": opts.PageW, "page height": opts.PageH, "page overlap": opts.Overlap} {
		if !finite(value) {
			return fmt.Errorf("%s must be finite", name)
		}
		if value < 0 {
			return fmt.Errorf("%s must be non-negative", name)
		}
	}
	for index, op := range plan.Ops {
		for name, value := range map[string]float64{"x": op.X, "y": op.Y, "width": op.W, "height": op.H} {
			if !finite(value) {
				return fmt.Errorf("draw op %d %s must be finite", index+1, name)
			}
		}
		if op.W < 0 || op.H < 0 {
			return fmt.Errorf("draw op %d size must be non-negative", index+1)
		}
		if !finite(op.X+op.W) || !finite(op.Y+op.H) {
			return fmt.Errorf("draw op %d bounds must be finite", index+1)
		}
		for pointIndex, point := range op.Points {
			if !finite(point.X) || !finite(point.Y) {
				return fmt.Errorf("draw op %d point %d must be finite", index+1, pointIndex+1)
			}
			if !finite(op.X+point.X) || !finite(op.Y+point.Y) {
				return fmt.Errorf("draw op %d point %d absolute position must be finite", index+1, pointIndex+1)
			}
		}
	}
	pageW := opts.PageW
	pageH := opts.PageH
	if pageW == 0 {
		pageW = plan.Slide.W
	}
	if pageH == 0 {
		pageH = plan.Slide.H
	}
	if opts.Overlap >= math.Min(pageW, pageH) {
		return fmt.Errorf("page overlap must be smaller than page width and height")
	}
	return nil
}

func pageCountV1EnginePageSplit(contentSize, pageSize, step float64) (int, error) {
	if contentSize <= pageSize {
		return 1, nil
	}
	count := math.Ceil((contentSize-pageSize)/step) + 1
	if math.IsNaN(count) || math.IsInf(count, 0) || count > 10000 {
		return 0, fmt.Errorf("tile count is too large")
	}
	return int(count), nil
}

func planContentRectV1EnginePageSplit(plan entity.Plan) entity.PageRect {
	rect := entity.PageRect{W: math.Max(plan.Slide.W, 0.01), H: math.Max(plan.Slide.H, 0.01)}
	for _, op := range planOpBoundsV1EnginePageSplit(plan.Ops) {
		rect = unionPageRectV1EnginePageSplit(rect, op.rect)
	}
	return rect
}

type planOpRectV1EnginePageSplit struct {
	id   string
	rect entity.PageRect
}

func planOpBoundsV1EnginePageSplit(ops []entity.DrawOp) []planOpRectV1EnginePageSplit {
	out := make([]planOpRectV1EnginePageSplit, 0, len(ops))
	seen := map[string]int{}
	for i, op := range ops {
		id := uniqueOpIDV1EnginePageSplit(op.ID, i)
		seen[id]++
		if seen[id] > 1 {
			id += "#" + strconv.Itoa(seen[id])
		}
		rect := entity.PageRect{X: op.X, Y: op.Y, W: math.Max(0, op.W), H: math.Max(0, op.H)}
		switch op.Kind {
		case "line", "polygon":
			if len(op.Points) > 0 {
				rect = pointsPageRectV1EnginePageSplit(op)
			}
		}
		if rect.W <= 0 {
			rect.W = 0.001
		}
		if rect.H <= 0 {
			rect.H = 0.001
		}
		out = append(out, planOpRectV1EnginePageSplit{id: id, rect: rect})
	}
	return out
}

func uniqueOpIDV1EnginePageSplit(id string, index int) string {
	if id != "" {
		return id
	}
	return "op-" + strconv.Itoa(index+1)
}

func pointsPageRectV1EnginePageSplit(op entity.DrawOp) entity.PageRect {
	minX, minY := math.Inf(1), math.Inf(1)
	maxX, maxY := math.Inf(-1), math.Inf(-1)
	for _, point := range op.Points {
		x := op.X + point.X
		y := op.Y + point.Y
		minX = math.Min(minX, x)
		minY = math.Min(minY, y)
		maxX = math.Max(maxX, x)
		maxY = math.Max(maxY, y)
	}
	return entity.PageRect{X: minX, Y: minY, W: maxX - minX, H: maxY - minY}
}

func unionPageRectV1EnginePageSplit(a, b entity.PageRect) entity.PageRect {
	minX := math.Min(a.X, b.X)
	minY := math.Min(a.Y, b.Y)
	maxX := math.Max(a.X+a.W, b.X+b.W)
	maxY := math.Max(a.Y+a.H, b.Y+b.H)
	return entity.PageRect{X: minX, Y: minY, W: maxX - minX, H: maxY - minY}
}

func rectIntersectsV1EnginePageSplit(a, b entity.PageRect) bool {
	return a.X < b.X+b.W && a.X+a.W > b.X && a.Y < b.Y+b.H && a.Y+a.H > b.Y
}
