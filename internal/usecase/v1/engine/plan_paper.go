package engine

import (
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

// paperSizeIn lists portrait paper dimensions in inches (width × height).
// Landscape swaps them.
var paperSizeInV1EnginePlanPaper = map[string][2]float64{
	"A5":      {5.83, 8.27},
	"A4":      {8.27, 11.69},
	"A3":      {11.69, 16.54},
	"A2":      {16.54, 23.39},
	"A1":      {23.39, 33.11},
	"Letter":  {8.5, 11},
	"Legal":   {8.5, 14},
	"Tabloid": {11, 17},
}

type paperMarginsV1EnginePlanPaper struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

// resolvePaper returns the target slide size (inches) for a named paper size,
// auto-selecting the orientation that fits the diagram best when not forced.
func resolvePaperV1EnginePlanPaper(size, orientation string, contentWIn, contentHIn float64, margins paperMarginsV1EnginePlanPaper) (w, h float64, ok bool) {
	base, found := paperSizeInV1EnginePlanPaper[size]
	if !found {
		return 0, 0, false
	}
	pw, ph := base[0], base[1]
	lw, lh := base[1], base[0]
	switch orientation {
	case "portrait":
		return pw, ph, true
	case "landscape":
		return lw, lh, true
	}
	availPW, availPH := margins.availableV1EnginePlanPaper(pw, ph)
	availLW, availLH := margins.availableV1EnginePlanPaper(lw, lh)
	scaleP := math.Min(availPW/contentWIn, availPH/contentHIn)
	scaleL := math.Min(availLW/contentWIn, availLH/contentHIn)
	if scaleL >= scaleP {
		return lw, lh, true
	}
	return pw, ph, true
}

func resolvePaperMarginsV1EnginePlanPaper(opt entity.PlanOptions) paperMarginsV1EnginePlanPaper {
	all := math.Max(0, opt.PaperMargin)
	margins := paperMarginsV1EnginePlanPaper{Top: all, Right: all, Bottom: all, Left: all}
	if opt.PaperMarginTop > 0 {
		margins.Top = opt.PaperMarginTop
	}
	if opt.PaperMarginRight > 0 {
		margins.Right = opt.PaperMarginRight
	}
	if opt.PaperMarginBottom > 0 {
		margins.Bottom = opt.PaperMarginBottom
	}
	if opt.PaperMarginLeft > 0 {
		margins.Left = opt.PaperMarginLeft
	}
	return margins
}

func (rcvr paperMarginsV1EnginePlanPaper) availableV1EnginePlanPaper(w, h float64) (float64, float64) {
	const minPaperContentIn = 0.01
	return math.Max(minPaperContentIn, w-rcvr.Left-rcvr.Right), math.Max(minPaperContentIn, h-rcvr.Top-rcvr.Bottom)
}
