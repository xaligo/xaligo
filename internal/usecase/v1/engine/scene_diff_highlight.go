package engine

import (
	"fmt"
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

const (
	diffAddedFillV1EngineSceneDiffHighlight     = "#DCFCE7"
	diffAddedStrokeV1EngineSceneDiffHighlight   = "#86EFAC"
	diffRemovedFillV1EngineSceneDiffHighlight   = "#FEE2E2"
	diffRemovedStrokeV1EngineSceneDiffHighlight = "#FCA5A5"
)

func appendDiffBoxHighlightsV1EngineSceneDiffHighlight(root *entity.Box, elements *[]map[string]any) {
	if root == nil {
		return
	}
	status := root.Attrs[diffStatusAttrV1EngineDiffDocument]
	if status == string(entity.StructuralChangeAdded) || status == string(entity.StructuralChangeRemoved) {
		fill, stroke := diffHighlightColorsV1EngineSceneDiffHighlight(status)
		if root.W > 0 && root.H > 0 && !math.IsNaN(root.W) && !math.IsNaN(root.H) && !math.IsInf(root.W, 0) && !math.IsInf(root.H, 0) {
			id := fmt.Sprintf("diff-%s-%s-%d", status, sanitizeElementIDV1EngineSceneConnectionRoute(root.ID), root.Position.Offset)
			seed := stableSceneSeedV1EngineSceneTypes(id)
			*elements = append(*elements, map[string]any{
				"id": id, "type": "rectangle",
				"x": root.X, "y": root.Y, "width": root.W, "height": root.H,
				"angle": 0, "strokeColor": stroke, "backgroundColor": fill,
				"fillStyle": "solid", "strokeWidth": 2, "strokeStyle": "solid",
				"roughness": 0, "opacity": 55,
				"groupIds": []string{}, "roundness": map[string]any{"type": 3},
				"seed": seed, "version": 1, "versionNonce": seed,
				"isDeleted": false, "boundElements": nil,
				"updated": excalidrawUpdatedV1EngineSceneTypes, "link": nil, "locked": false,
				"customData": map[string]any{"xaligoDiffHighlight": true, "xaligoDiffStatus": status},
			})
		}
	}
	for _, child := range root.Children {
		appendDiffBoxHighlightsV1EngineSceneDiffHighlight(child, elements)
	}
}

func applyConnectionDiffStatusV1EngineSceneDiffHighlight(customData map[string]any, connection *entity.Node) {
	if customData == nil || connection == nil {
		return
	}
	status := connection.Attrs[diffStatusAttrV1EngineDiffDocument]
	if status == string(entity.StructuralChangeAdded) || status == string(entity.StructuralChangeRemoved) {
		customData["xaligoDiffStatus"] = status
	}
}

func diffHighlightColorsV1EngineSceneDiffHighlight(status string) (string, string) {
	if status == string(entity.StructuralChangeRemoved) {
		return diffRemovedFillV1EngineSceneDiffHighlight, diffRemovedStrokeV1EngineSceneDiffHighlight
	}
	return diffAddedFillV1EngineSceneDiffHighlight, diffAddedStrokeV1EngineSceneDiffHighlight
}
