package engine

import (
	"math"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

var (
	IURBS003V1EngineSceneConnection = share.NewMCode("IURBS-003", "Build scene connection node branch")
)

func CollectConnectionNodesV1EngineSceneConnection(root *entity.Node) []*entity.Node {
	if root == nil {
		return nil
	}
	connections := []*entity.Node{}
	var walk func(*entity.Node)
	walk = func(node *entity.Node) {
		if node.Tag == "frame" {
			for _, child := range node.Children {
				switch child.Tag {
				case "connection":
					loggerV1EngineSharedLogging.DEBUG(IURBS003V1EngineSceneConnection, "branch connection node", map[string]any{"tag": child.Tag})
					connections = append(connections, child)
				case "connections":
					defaults := connectionGroupDefaultsV1EngineParseConnection(child)
					for _, grouped := range child.Children {
						if grouped.Tag != "connection" {
							continue
						}
						loggerV1EngineSharedLogging.DEBUG(IURBS003V1EngineSceneConnection, "branch grouped connection node", map[string]any{"tag": grouped.Tag})
						connections = append(connections, connectionWithDefaultsV1EngineParseConnection(grouped, defaults))
					}
				}
			}
		}
		for _, child := range node.Children {
			walk(child)
		}
	}
	walk(root)
	return connections
}

// connectionSide determines which edge exits src (srcSide) and enters dst (dstSide)
// based on the direction between their center points.
// Returns "top", "bottom", "left", or "right".
func connectionSideV1EngineSceneConnection(srcCx, srcCy, dstCx, dstCy float64) (srcSide, dstSide string) {
	dx := dstCx - srcCx
	dy := dstCy - srcCy
	if math.Abs(dx) >= math.Abs(dy) {
		if dx >= 0 {
			return "right", "left"
		}
		return "left", "right"
	}
	if dy >= 0 {
		return "bottom", "top"
	}
	return "top", "bottom"
}

func connectionBendPointsV1EngineSceneConnection(conn *entity.Node) []ptV1EngineRouteTypes {
	scale := 1.0
	if value, ok := positiveFloatAttrV1EngineSceneConnectionRoute(conn, "coordinate-scale", "scale"); ok {
		scale = value
	}
	return parseConnectorBendsV1EnginePlanConnectorPrepare(connectionBendsV1EngineSceneConnectionRoute(conn), scale)
}

func sideTowardPointV1EngineSceneConnection(rect [4]float64, point ptV1EngineRouteTypes) string {
	cx := rect[0] + rect[2]/2
	cy := rect[1] + rect[3]/2
	dx := point.X - cx
	dy := point.Y - cy
	if math.Abs(dx) >= math.Abs(dy) {
		if dx >= 0 {
			return "right"
		}
		return "left"
	}
	if dy >= 0 {
		return "bottom"
	}
	return "top"
}

// rectEdgePoint returns the midpoint of the named edge of a rectangle.
// rect = [x, y, w, h]; side is "top", "bottom", "left", or "right".
func rectEdgePointV1EngineSceneConnection(rect [4]float64, side string) [2]float64 {
	x, y, w, h := rect[0], rect[1], rect[2], rect[3]
	cx := x + w/2
	cy := y + h/2
	switch side {
	case "top":
		return [2]float64{cx, y}
	case "bottom":
		return [2]float64{cx, y + h}
	case "left":
		return [2]float64{x, cy}
	default: // "right"
		return [2]float64{x + w, cy}
	}
}

// fixedPointForSide returns the normalized [x, y] fixedPoint on an element's bounding box
// that corresponds to the given side. This matches Excalidraw's binding coordinate system:
// [0,0]=top-left, [1,1]=bottom-right; each side midpoint:
//
//	top=[0.5,0], bottom=[0.5,1], left=[0,0.5], right=[1,0.5]
func fixedPointForSideV1EngineSceneConnection(side string) [2]float64 {
	switch side {
	case "top":
		return [2]float64{0.5, 0}
	case "bottom":
		return [2]float64{0.5, 1}
	case "left":
		return [2]float64{0, 0.5}
	default: // "right"
		return [2]float64{1, 0.5}
	}
}

func fixedPointForAnchorV1EngineSceneConnection(anchor connectionAnchorSpecV1EngineParseConnection) [2]float64 {
	pos := (float64(anchor.slot) + 0.5) / 5.0
	switch anchor.side {
	case sideTopV1EngineRouteTypes:
		return [2]float64{pos, 0}
	case sideBottomV1EngineRouteTypes:
		return [2]float64{pos, 1}
	case sideLeftV1EngineRouteTypes:
		return [2]float64{0, pos}
	default:
		return [2]float64{1, pos}
	}
}

func rectFixedPointV1EngineSceneConnection(rect [4]float64, fp [2]float64) [2]float64 {
	return [2]float64{rect[0] + rect[2]*fp[0], rect[1] + rect[3]*fp[1]}
}
