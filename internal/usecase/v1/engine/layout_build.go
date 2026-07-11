package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

var (
	IULB001V1EngineLayoutBuild = share.NewMCode("IULB-001", "Build nil root branch")
	IULB002V1EngineLayoutBuild = share.NewMCode("IULB-002", "Build layout root branch")
)

func BuildV1EngineLayoutBuild(doc entity.Document) (*entity.Box, error) {
	if doc.Root == nil {
		loggerV1EngineSharedLogging.DEBUG(IULB001V1EngineLayoutBuild, "branch nil root")
		return nil, fmt.Errorf("document root is nil")
	}
	if err := validateLayoutDocumentV1EngineLayoutValidation(doc.Root); err != nil {
		return nil, err
	}
	if doc.Root.Tag == "frames" {
		root, err := buildFramesV1EngineLayoutBuild(doc.Root)
		if err != nil {
			return nil, err
		}
		if err := validateResolvedGeometryV1EngineLayoutValidation(root); err != nil {
			return nil, err
		}
		if err := validateResolvedItemGridsV1EngineLayoutItemGrid(root); err != nil {
			return nil, err
		}
		return root, nil
	}
	w := attrFloatV1EngineLayoutAttributes(doc.Root.Attr("width"), 1280)
	h := attrFloatV1EngineLayoutAttributes(doc.Root.Attr("height"), 720)
	root := &entity.Box{ID: "frame", Tag: "frame", Label: "frame", Position: doc.Root.Position, X: 0, Y: 0, W: w, H: h}
	loggerV1EngineSharedLogging.DEBUG(IULB002V1EngineLayoutBuild, "branch layout root", map[string]any{"width": w, "height": h})
	if err := layoutNodeV1EngineLayoutNode(doc.Root, root, 0, 0, w, h); err != nil {
		return nil, err
	}
	if err := validateResolvedGeometryV1EngineLayoutValidation(root); err != nil {
		return nil, err
	}
	if err := validateResolvedItemGridsV1EngineLayoutItemGrid(root); err != nil {
		return nil, err
	}
	return root, nil
}

func buildFramesV1EngineLayoutBuild(rootNode *entity.Node) (*entity.Box, error) {
	gap := attrFloatV1EngineLayoutAttributes(rootNode.Attr("gap"), 48)
	horizontal := rootNode.Attr("layout") != "vertical"
	root := &entity.Box{ID: "frames", Tag: "frames", Label: "frames", Position: rootNode.Position, X: 0, Y: 0, Attrs: rootNode.Attrs, Overflow: normalizedOverflowV1EngineLayoutConstraints(rootNode)}
	curX, curY := 0.0, 0.0
	maxW, maxH := 0.0, 0.0
	for i, child := range rootNode.Children {
		if child.Tag != "frame" {
			continue
		}
		w := attrFloatV1EngineLayoutAttributes(child.Attr("width"), attrFloatV1EngineLayoutAttributes(rootNode.Attr("width"), 1280))
		h := attrFloatV1EngineLayoutAttributes(child.Attr("height"), attrFloatV1EngineLayoutAttributes(rootNode.Attr("height"), 720))
		id := strings.TrimSpace(child.Attrs["id"])
		boxID := childIDV1EngineLayoutAttributes(root.ID, i)
		if id != "" {
			boxID = sanitizeElementIDV1EngineSceneConnectionRoute(id)
		}
		cb := &entity.Box{ID: boxID, Tag: child.Tag, Label: labelOfV1EngineLayoutAttributes(child), Position: child.Position}
		if err := layoutNodeV1EngineLayoutNode(child, cb, curX, curY, w, h); err != nil {
			return nil, err
		}
		root.Children = append(root.Children, cb)
		if horizontal {
			curX += w + gap
			maxH = maxFloatV1EngineLayoutBuild(maxH, h)
		} else {
			curY += h + gap
			maxW = maxFloatV1EngineLayoutBuild(maxW, w)
		}
	}
	if horizontal {
		root.W = maxFloatV1EngineLayoutBuild(0, curX-gap)
		root.H = maxH
	} else {
		root.W = maxW
		root.H = maxFloatV1EngineLayoutBuild(0, curY-gap)
	}
	root.ContentX = root.X
	root.ContentY = root.Y
	root.ContentW = root.W
	root.ContentH = root.H
	return root, nil
}

func maxFloatV1EngineLayoutBuild(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
