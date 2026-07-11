package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func anchorGridOpsV1EnginePlanAnchor(id string, r rectV1EngineRouteTypes, frame rectV1EngineRouteTypes, ppi float64, background string) []entity.DrawOp {
	cellW := r.W / float64(anchorGridV1EnginePlanBuild)
	cellH := r.H / float64(anchorGridV1EnginePlanBuild)
	ops := make([]entity.DrawOp, 0, anchorGridV1EnginePlanBuild*anchorGridV1EnginePlanBuild)
	baseID := anchorBaseIDV1EnginePlanAnchor(id)
	for i := 0; i < anchorGridV1EnginePlanBuild; i++ {
		for j := 0; j < anchorGridV1EnginePlanBuild; j++ {
			cx := r.X + float64(i)*cellW
			cy := r.Y + float64(j)*cellH
			ops = append(ops, entity.DrawOp{
				ID:   fmt.Sprintf("%s-grid-%02d-%02d", baseID, i, j),
				Kind: "rect",
				X:    (cx - frame.X) / ppi,
				Y:    (cy - frame.Y) / ppi,
				W:    cellW / ppi,
				H:    cellH / ppi,
				Fill: &entity.FillStyle{Color: background, Transparency: 0},
				Line: &entity.LineStyle{Color: background, Width: pxToPtV1EnginePlanStyle(0.25, ppi), Dash: "solid", Transparency: 0},
			})
		}
	}
	return ops
}

func anchorGroupsV1EnginePlanAnchor(grids map[string]anchorGridRectV1EnginePlanConnectorPrepare) map[string]string {
	out := map[string]string{}
	for id := range grids {
		baseID := anchorBaseIDV1EnginePlanAnchor(id)
		out[baseID] = anchorGroupIDV1EnginePlanAnchor(baseID)
	}
	return out
}

func applyAnchorGroupV1EnginePlanAnchor(op *entity.DrawOp, elementID string, groups map[string]string) {
	if op == nil || elementID == "" {
		return
	}
	baseID := anchorBaseIDV1EnginePlanAnchor(elementID)
	groupID, ok := groups[baseID]
	if !ok {
		return
	}
	op.GroupID = groupID
}

func anchorBaseIDV1EnginePlanAnchor(id string) string {
	return strings.TrimSuffix(id, "-lbl")
}

func anchorGroupIDV1EnginePlanAnchor(baseID string) string {
	return "xaligo-anchor-" + baseID
}
