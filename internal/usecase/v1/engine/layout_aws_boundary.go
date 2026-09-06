package engine

import (
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type awsBoundaryPlacementV1EngineLayoutAwsBoundary struct {
	node       *entity.Node
	x, y, size float64
}

func hasAWSBoundaryAttachmentChildrenV1EngineLayoutAwsBoundary(node *entity.Node) bool {
	if node == nil {
		return false
	}
	for _, child := range node.Children {
		if isAWSBoundaryAttachmentV1EngineAwsBoundary(child.Tag) {
			return true
		}
	}
	return false
}

func layoutAWSBoundaryAttachmentsV1EngineLayoutAwsBoundary(node *entity.Node, target *entity.Box) error {
	if node == nil || target == nil {
		return nil
	}
	autoBySide := map[string][]*entity.Node{}
	for _, child := range node.Children {
		definition, ok := awsBoundaryAttachmentV1EngineAwsBoundary(child.Tag)
		if !ok || strings.TrimSpace(child.Attr("anchor")) != "" {
			continue
		}
		side := normalizedAWSBoundarySideV1EngineLayoutAwsBoundary(child, definition.DefaultSide)
		autoBySide[side] = append(autoBySide[side], child)
	}
	autoAnchors := map[*entity.Node]float64{}
	for _, children := range autoBySide {
		for index, child := range children {
			autoAnchors[child] = float64(index+1) / float64(len(children)+1)
		}
	}

	placements := make([]awsBoundaryPlacementV1EngineLayoutAwsBoundary, 0, len(node.Children))
	for sourceIndex, child := range node.Children {
		definition, ok := awsBoundaryAttachmentV1EngineAwsBoundary(child.Tag)
		if !ok {
			continue
		}
		side := normalizedAWSBoundarySideV1EngineLayoutAwsBoundary(child, definition.DefaultSide)
		size := attrFloatV1EngineLayoutAttributes(child.Attr("size"), definition.DefaultSize)
		anchor, hasAnchor := attrFloatOKV1EngineLayoutAttributes(child.Attr("anchor"))
		if !hasAnchor {
			anchor = autoAnchors[child]
		}
		offset := attrFloatV1EngineLayoutAttributes(child.Attr("offset"), 0)
		x, y := awsBoundaryPositionV1EngineLayoutAwsBoundary(target, side, anchor, offset, size)
		if !awsBoundaryTangentContainsV1EngineLayoutAwsBoundary(target, side, x, y, size) {
			return newLayoutErrorV1EngineLayoutValidation(child, "<%s id=%q> offset places the icon beyond the <%s> %s edge", child.Tag, child.Attr("id"), node.Tag, side)
		}
		for _, previous := range placements {
			if rectanglesOverlapV1EngineLayoutPort(x, y, size, size, previous.x, previous.y, previous.size, previous.size) {
				return newLayoutErrorV1EngineLayoutValidation(child, "%s %q overlaps %q on the <%s> boundary", child.Tag, child.Attr("id"), previous.node.Attr("id"), node.Tag)
			}
		}
		placements = append(placements, awsBoundaryPlacementV1EngineLayoutAwsBoundary{node: child, x: x, y: y, size: size})
		box := &entity.Box{
			ID:  childIDV1EngineLayoutAttributes(target.ID, sourceIndex) + "-boundary",
			Tag: child.Tag, Attrs: child.Attrs, Position: child.Position,
			X: x, Y: y, W: size, H: size, IntrinsicW: size, IntrinsicH: size,
		}
		setContentBoxV1EngineLayoutConstraints(box, x, y, size, size)
		target.Children = append(target.Children, box)
	}
	return nil
}

func normalizedAWSBoundarySideV1EngineLayoutAwsBoundary(node *entity.Node, fallback string) string {
	return normalizedAWSBoundarySideValueV1EngineLayoutAwsBoundary(node.Attr("side"), fallback)
}

func normalizedAWSBoundarySideValueV1EngineLayoutAwsBoundary(raw, fallback string) string {
	side := strings.ToLower(strings.TrimSpace(raw))
	if side == "" {
		return fallback
	}
	return side
}

func awsBoundaryPositionV1EngineLayoutAwsBoundary(parent *entity.Box, side string, anchor, offset, size float64) (float64, float64) {
	switch side {
	case "top":
		return parent.X + anchor*(parent.W-size) + offset, parent.Y - size/2
	case "bottom":
		return parent.X + anchor*(parent.W-size) + offset, parent.Y + parent.H - size/2
	case "left":
		return parent.X - size/2, parent.Y + anchor*(parent.H-size) + offset
	default:
		return parent.X + parent.W - size/2, parent.Y + anchor*(parent.H-size) + offset
	}
}

func awsBoundaryTangentContainsV1EngineLayoutAwsBoundary(parent *entity.Box, side string, x, y, size float64) bool {
	if parent == nil || !isPositiveFiniteV1EngineLayoutConstraints(size) {
		return false
	}
	const epsilon = geometryEpsilonV1EngineLayoutValidation
	if side == "top" || side == "bottom" {
		return x+epsilon >= parent.X && x+size <= parent.X+parent.W+epsilon
	}
	return y+epsilon >= parent.Y && y+size <= parent.Y+parent.H+epsilon
}

func resolvedAWSBoundaryAttachmentV1EngineLayoutAwsBoundary(box, parent *entity.Box) bool {
	if box == nil || parent == nil {
		return false
	}
	definition, ok := awsBoundaryAttachmentV1EngineAwsBoundary(box.Tag)
	if !ok {
		return false
	}
	side := normalizedAWSBoundarySideValueV1EngineLayoutAwsBoundary(box.Attrs["side"], definition.DefaultSide)
	if parent.Tag != definition.ParentTag || !awsBoundaryTangentContainsV1EngineLayoutAwsBoundary(parent, side, box.X, box.Y, box.W) || math.Abs(box.W-box.H) > geometryEpsilonV1EngineLayoutValidation {
		return false
	}
	centerX := box.X + box.W/2
	centerY := box.Y + box.H/2
	switch side {
	case "top":
		return math.Abs(centerY-parent.Y) <= geometryEpsilonV1EngineLayoutValidation
	case "bottom":
		return math.Abs(centerY-(parent.Y+parent.H)) <= geometryEpsilonV1EngineLayoutValidation
	case "left":
		return math.Abs(centerX-parent.X) <= geometryEpsilonV1EngineLayoutValidation
	case "right":
		return math.Abs(centerX-(parent.X+parent.W)) <= geometryEpsilonV1EngineLayoutValidation
	default:
		return false
	}
}
