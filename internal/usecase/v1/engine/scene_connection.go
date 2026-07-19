package engine

import (
	"math"
	"strings"

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

func fixedPointForUMLProfileV1EngineSceneConnection(conn *entity.Node, endpoint, side string, rect, otherRect [4]float64) ([2]float64, bool) {
	if conn == nil || strings.TrimSpace(conn.Attr("uml-diagram-kind")) == "" || conn.Attr("uml-diagram-kind") == "sequence-diagram" {
		return [2]float64{}, false
	}
	profile := umlEndpointAnchorProfileV1EngineSceneConnection(conn.Attr("uml-" + endpoint + "-kind"))
	switch profile {
	case "diamond":
		return fixedPointForSideV1EngineSceneConnection(side), true
	case "rectangle":
		return fixedPointForSideSlotV1EngineSceneConnection(side, slotForUMLRectangleEndpointV1EngineSceneConnection(side, rect, otherRect)), true
	default:
		return [2]float64{}, false
	}
}

func umlEndpointAnchorProfileV1EngineSceneConnection(kind string) string {
	switch strings.TrimSpace(kind) {
	case "choice", "decision", "merge", "history":
		return "diamond"
	case "state", "class", "interface", "enumeration", "object", "component", "artifact", "node", "package", "structure", "collaboration", "part", "activity", "action", "object-node", "interaction", "time-state":
		return "rectangle"
	default:
		return ""
	}
}

func slotForUMLRectangleEndpointV1EngineSceneConnection(side string, rect, otherRect [4]float64) int {
	fraction := 0.5
	switch side {
	case "top", "bottom":
		if rect[2] > 0 {
			fraction = (otherRect[0] + otherRect[2]/2 - rect[0]) / rect[2]
		}
	case "left", "right":
		if rect[3] > 0 {
			fraction = (otherRect[1] + otherRect[3]/2 - rect[1]) / rect[3]
		}
	}
	fraction = math.Max(0, math.Min(1, fraction))
	return int(math.Round(fraction * float64(anchorGridV1EnginePlanBuild-1)))
}

func fixedPointForSideSlotV1EngineSceneConnection(side string, slot int) [2]float64 {
	if slot < 0 {
		slot = 0
	}
	if slot >= anchorGridV1EnginePlanBuild {
		slot = anchorGridV1EnginePlanBuild - 1
	}
	position := (float64(slot) + 0.5) / float64(anchorGridV1EnginePlanBuild)
	switch side {
	case "top":
		return [2]float64{position, 0}
	case "bottom":
		return [2]float64{position, 1}
	case "left":
		return [2]float64{0, position}
	default:
		return [2]float64{1, position}
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
