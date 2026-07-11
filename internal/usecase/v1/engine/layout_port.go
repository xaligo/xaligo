package engine

import "github.com/xaligo/xaligo/internal/entity"

type portPlacementV1EngineLayoutPort struct {
	Node       *entity.Node
	X, Y, W, H float64
}

func rectanglesOverlapV1EngineLayoutPort(ax, ay, aw, ah, bx, by, bw, bh float64) bool {
	return ax < bx+bw-geometryEpsilonV1EngineLayoutValidation && ax+aw > bx+geometryEpsilonV1EngineLayoutValidation &&
		ay < by+bh-geometryEpsilonV1EngineLayoutValidation && ay+ah > by+geometryEpsilonV1EngineLayoutValidation
}

func portPositionV1EngineLayoutPort(x, y, w, h, portW, portH float64, side string, index, total int) (float64, float64) {
	slot := float64(index+1) / float64(total+1)
	switch side {
	case "right":
		return x + w - portW, y + h*slot - portH/2
	case "bottom":
		return x + w*slot - portW/2, y + h - portH
	case "left":
		return x, y + h*slot - portH/2
	default:
		return x + w*slot - portW/2, y
	}
}

func clampPortPositionV1EngineLayoutPort(portX, portY, parentX, parentY, parentW, parentH, portW, portH float64) (float64, float64) {
	return clampFloatV1EngineLayoutPort(portX, parentX, parentX+parentW-portW), clampFloatV1EngineLayoutPort(portY, parentY, parentY+parentH-portH)
}

func clampFloatV1EngineLayoutPort(value, min, max float64) float64 {
	if max < min {
		return min
	}
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
