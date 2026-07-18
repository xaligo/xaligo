package engine

import (
	"fmt"
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const (
	frameMetadataTextPaddingXV1EngineSceneFrameMetadata = 4.0
	frameMetadataTextPaddingYV1EngineSceneFrameMetadata = 2.0
)

type frameMetadataSceneGeometryV1EngineSceneFrameMetadata struct {
	Position string
	Rects    [][4]float64
}

func appendFrameMetadataV1EngineSceneFrameMetadata(root *entity.Box, elements *[]map[string]any) map[string]frameMetadataSceneGeometryV1EngineSceneFrameMetadata {
	occupied := map[string]frameMetadataSceneGeometryV1EngineSceneFrameMetadata{}
	if root == nil || elements == nil {
		return occupied
	}
	if root.Tag == "frames" {
		for _, frame := range root.Children {
			if frame == nil || frame.Tag != "frame" || frame.FrameMetadata == nil {
				continue
			}
			frameID := strings.TrimSpace(frame.Attrs["id"])
			if frameID == "" {
				frameID = frame.ID
			}
			occupied[frameID] = frameMetadataSceneGeometryV1EngineSceneFrameMetadata{
				Position: frame.FrameMetadata.Position,
				Rects:    appendFrameMetadataBoxV1EngineSceneFrameMetadata(frame, frameID, pageFrameElementIDV1EngineSceneWalk(frame), elements),
			}
		}
		return occupied
	}
	if root.Tag == "frame" && root.FrameMetadata != nil {
		frameID := strings.TrimSpace(root.Attrs["id"])
		occupied[frameID] = frameMetadataSceneGeometryV1EngineSceneFrameMetadata{
			Position: root.FrameMetadata.Position,
			Rects:    appendFrameMetadataBoxV1EngineSceneFrameMetadata(root, frameID, "paper-frame", elements),
		}
	}
	return occupied
}

func appendFrameMetadataBoxV1EngineSceneFrameMetadata(frame *entity.Box, frameID, semanticParent string, elements *[]map[string]any) [][4]float64 {
	metadata := frame.FrameMetadata
	if metadata == nil {
		return nil
	}
	fontFamily := fontFamilyV1EngineSceneWalk(metadata.FontFamily)
	occupied := make([][4]float64, 0, len(metadata.Tags))
	for index, tag := range metadata.Tags {
		baseID := fmt.Sprintf("%s-metadata-%02d", frame.ID, index)
		keyID := baseID + "-key"
		valueID := baseID + "-value"
		keyTextID := baseID + "-key-content"
		valueTextID := baseID + "-value-content"
		keyWidth := math.Max(0.1, math.Min(tag.W-0.1, tag.KeyW))
		valueWidth := math.Max(0.1, tag.W-keyWidth)
		keyRect := [4]float64{tag.X, tag.Y, keyWidth, tag.H}
		valueRect := [4]float64{tag.X + keyWidth, tag.Y, valueWidth, tag.H}
		occupied = append(occupied, [4]float64{tag.X, tag.Y, tag.W, tag.H})

		*elements = append(*elements,
			frameMetadataRectangleV1EngineSceneFrameMetadata(keyID, keyTextID, keyRect, metadata.KeyBackgroundColor, metadata.BorderColor, frameID, semanticParent),
			frameMetadataRectangleV1EngineSceneFrameMetadata(valueID, valueTextID, valueRect, metadata.BackgroundColor, metadata.BorderColor, frameID, semanticParent),
		)
		*elements = append(*elements,
			frameMetadataTextV1EngineSceneFrameMetadata(keyTextID, keyID, tag.Key, keyRect, metadata.KeyColor, metadata.FontSize, fontFamily, frameID, semanticParent),
			frameMetadataTextV1EngineSceneFrameMetadata(valueTextID, valueID, tag.Value, valueRect, metadata.Color, metadata.FontSize, fontFamily, frameID, semanticParent),
		)
		if tag.DiffStatus != "" {
			*elements = append(*elements, frameMetadataDiffHighlightV1EngineSceneFrameMetadata(
				fmt.Sprintf("diff-%s-%s", tag.DiffStatus, baseID),
				[4]float64{tag.X, tag.Y, tag.W, tag.H},
				tag.DiffStatus,
				frameID,
				semanticParent,
			))
		}
	}
	return occupied
}

func frameMetadataDiffHighlightV1EngineSceneFrameMetadata(id string, rect [4]float64, status, frameID, semanticParent string) map[string]any {
	fill, stroke := diffHighlightColorsV1EngineSceneDiffHighlight(status)
	seed := stableSceneSeedV1EngineSceneTypes(id)
	customData := frameMetadataCustomDataV1EngineSceneFrameMetadata(frameID, semanticParent, false)
	customData["xaligoDiffHighlight"] = true
	customData["xaligoDiffStatus"] = status
	return map[string]any{
		"id": id, "type": "rectangle",
		"x": rect[0], "y": rect[1], "width": rect[2], "height": rect[3],
		"angle":       0,
		"strokeColor": stroke, "backgroundColor": fill,
		"fillStyle": "solid", "strokeWidth": 2, "strokeStyle": "solid",
		"roughness": 0, "opacity": 55,
		"groupIds": []string{}, "roundness": map[string]any{"type": 3},
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": excalidrawUpdatedV1EngineSceneTypes, "link": nil, "locked": false, "frameId": nil,
		"customData": customData,
	}
}

func frameMetadataRectangleV1EngineSceneFrameMetadata(id, textID string, rect [4]float64, backgroundColor, borderColor, frameID, semanticParent string) map[string]any {
	seed := stableSceneSeedV1EngineSceneTypes(id)
	customData := frameMetadataCustomDataV1EngineSceneFrameMetadata(frameID, semanticParent, false)
	return map[string]any{
		"id": id, "type": "rectangle",
		"x": rect[0], "y": rect[1], "width": rect[2], "height": rect[3],
		"angle":       0,
		"strokeColor": borderColor, "backgroundColor": backgroundColor,
		"fillStyle": "solid", "strokeWidth": 0.75, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": []map[string]any{{"type": "text", "id": textID}},
		"updated": excalidrawUpdatedV1EngineSceneTypes, "link": nil, "locked": false, "frameId": nil,
		"customData": customData,
	}
}

func frameMetadataTextV1EngineSceneFrameMetadata(id, containerID, value string, rect [4]float64, color string, fontSize float64, fontFamily int, frameID, semanticParent string) map[string]any {
	seed := stableSceneSeedV1EngineSceneTypes(id)
	textX := rect[0] + frameMetadataTextPaddingXV1EngineSceneFrameMetadata
	textY := rect[1] + frameMetadataTextPaddingYV1EngineSceneFrameMetadata
	textW := math.Max(0.1, rect[2]-frameMetadataTextPaddingXV1EngineSceneFrameMetadata*2)
	textH := math.Max(0.1, rect[3]-frameMetadataTextPaddingYV1EngineSceneFrameMetadata*2)
	customData := frameMetadataCustomDataV1EngineSceneFrameMetadata(frameID, semanticParent, true)
	customData["xaligoTextLayout"] = sceneTextLayoutV1EngineSceneBuild(entity.TextRoleLabel, false, 1.2)
	return map[string]any{
		"id": id, "type": "text",
		"x": textX, "y": textY, "width": textW, "height": textH,
		"angle":       0,
		"strokeColor": color, "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": excalidrawUpdatedV1EngineSceneTypes, "link": nil, "locked": false, "frameId": nil,
		"text": value, "rawText": value, "originalText": value,
		"fontSize": fontSize, "fontFamily": fontFamily,
		"textAlign": "center", "verticalAlign": "middle",
		"containerId": containerID, "lineHeight": 1.2,
		"customData": customData,
	}
}

func frameMetadataCustomDataV1EngineSceneFrameMetadata(frameID, semanticParent string, content bool) map[string]any {
	customData := map[string]any{
		"xaligoFrameMetadata":       true,
		"xaligoSemanticElementKind": "frame-metadata",
	}
	if content {
		customData["xaligoFrameMetadataContent"] = true
	}
	if frameID != "" {
		customData["xaligoFrameID"] = frameID
	}
	if semanticParent != "" {
		customData["xaligoSemanticParentElementId"] = semanticParent
	}
	return customData
}
