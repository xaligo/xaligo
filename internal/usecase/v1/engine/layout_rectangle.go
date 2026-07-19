package engine

import (
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func layoutLeafV1EngineLayoutRectangle(node *entity.Node, target *entity.Box, x, y, w, h float64) {
	target.X = x
	target.Y = y
	target.W = w
	target.H = h
	setContentBoxV1EngineLayoutConstraints(target, x, y, w, h)
}

func layoutRectangleV1EngineLayoutRectangle(node *entity.Node, target *entity.Box, x, y, w, h float64) error {
	layoutLeafV1EngineLayoutRectangle(node, target, x, y, w, h)
	portsBySide := map[string][]*entity.Node{}
	for _, child := range node.Children {
		if !isPortNodeV1EngineLayoutRectangle(child) {
			continue
		}
		side := strings.ToLower(strings.TrimSpace(child.Attr("side")))
		if side != "right" && side != "bottom" && side != "left" {
			side = "top"
		}
		portsBySide[side] = append(portsBySide[side], child)
	}
	for _, side := range []string{"top", "right", "bottom", "left"} {
		ports := portsBySide[side]
		placed := make([]portPlacementV1EngineLayoutPort, 0, len(ports))
		for i, port := range ports {
			portW := attrFloatV1EngineLayoutAttributes(firstAttrV1EngineLayoutAttributes(port, "width", "w"), 48)
			portH := attrFloatV1EngineLayoutAttributes(firstAttrV1EngineLayoutAttributes(port, "height", "h"), 20)
			if normalizedOverflowV1EngineLayoutConstraints(node) != entity.OverflowVisible && (portW > w+geometryEpsilonV1EngineLayoutValidation || portH > h+geometryEpsilonV1EngineLayoutValidation) {
				return newLayoutErrorV1EngineLayoutValidation(port, "port size %.6gx%.6g overflows rectangle %.6gx%.6g", portW, portH, w, h)
			}
			portX, portY := portPositionV1EngineLayoutPort(x, y, w, h, portW, portH, side, i, len(ports))
			if px, ok := attrFloatOKV1EngineLayoutAttributes(port.Attr("x")); ok {
				portX = x + px
			}
			if py, ok := attrFloatOKV1EngineLayoutAttributes(port.Attr("y")); ok {
				portY = y + py
			}
			portX, portY = clampPortPositionV1EngineLayoutPort(portX, portY, x, y, w, h, portW, portH)
			for _, previous := range placed {
				if rectanglesOverlapV1EngineLayoutPort(portX, portY, portW, portH, previous.X, previous.Y, previous.W, previous.H) {
					return newLayoutErrorV1EngineLayoutValidation(port, "port %q overlaps port %q on the %s side of rectangle %q", port.Attr("id"), previous.Node.Attr("id"), side, node.Attr("id"))
				}
			}
			placed = append(placed, portPlacementV1EngineLayoutPort{Node: port, X: portX, Y: portY, W: portW, H: portH})
			cb := &entity.Box{
				ID: childIDV1EngineLayoutAttributes(target.ID, len(target.Children)), Tag: port.Tag, Label: labelOfV1EngineLayoutAttributes(port),
				Attrs: port.Attrs, Position: port.Position, IntrinsicW: portW, IntrinsicH: portH,
			}
			layoutLeafV1EngineLayoutRectangle(port, cb, portX, portY, portW, portH)
			target.Children = append(target.Children, cb)
		}
	}
	return nil
}

func isPortNodeV1EngineLayoutRectangle(node *entity.Node) bool {
	return node != nil && (node.Tag == "port" || node.Attr("uml-element-kind") == "port")
}
