package engine

import (
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

func connectorDiffHighlightOpV1EnginePlanDiffHighlight(op entity.DrawOp, status string) (entity.DrawOp, bool) {
	if op.Kind != "line" || op.Line == nil {
		return entity.DrawOp{}, false
	}
	var color string
	switch status {
	case string(entity.StructuralChangeAdded):
		color = "86EFAC"
	case string(entity.StructuralChangeRemoved):
		color = "FCA5A5"
	default:
		return entity.DrawOp{}, false
	}
	highlight := op
	line := *op.Line
	line.Color = color
	line.Width = math.Max(4, line.Width*3)
	line.Transparency = 25
	line.BeginArrowType = "none"
	line.EndArrowType = "none"
	highlight.ID = op.ID + "-diff-highlight"
	highlight.Line = &line
	return highlight, true
}

func connectorDiffStatusV1EnginePlanDiffHighlight(element *entity.Element) string {
	if element == nil || element.CustomData == nil {
		return ""
	}
	return element.CustomData.DiffStatus
}
