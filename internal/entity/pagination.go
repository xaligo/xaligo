package entity

// PageSplitOptions controls shared page tiling for large rendered plans.
// Values are in the same inch coordinate space as Plan and DrawOp.
type PageSplitOptions struct {
	PageW   float64 `json:"pageW,omitempty"`
	PageH   float64 `json:"pageH,omitempty"`
	Overlap float64 `json:"overlap,omitempty"`
}

// PageSplit describes how a plan can be split into page-sized tiles.
type PageSplit struct {
	Content PageRect   `json:"content"`
	PageW   float64    `json:"pageW"`
	PageH   float64    `json:"pageH"`
	Overlap float64    `json:"overlap,omitempty"`
	Rows    int        `json:"rows"`
	Cols    int        `json:"cols"`
	Pages   []PlanPage `json:"pages"`
}

// PlanPage is one page-sized tile in a split plan.
type PlanPage struct {
	Index int      `json:"index"`
	Row   int      `json:"row"`
	Col   int      `json:"col"`
	Rect  PageRect `json:"rect"`
	OpIDs []string `json:"opIds,omitempty"`
}

// PageRect is an inch-space rectangle.
type PageRect struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}
