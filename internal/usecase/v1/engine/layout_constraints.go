package engine

import (
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func normalizedOverflowV1EngineLayoutConstraints(node *entity.Node) entity.OverflowPolicy {
	if strings.EqualFold(strings.TrimSpace(node.Attr("overflow")), "visible") {
		return entity.OverflowVisible
	}
	return entity.OverflowError
}

func explicitSizeV1EngineLayoutConstraints(node *entity.Node, attribute string) (float64, bool) {
	value := strings.TrimSpace(node.Attr(attribute))
	if value == "" {
		return 0, false
	}
	return attrFloatV1EngineLayoutAttributes(value, 0), true
}

func flexibleAxisSpaceV1EngineLayoutConstraints(node *entity.Node, axis string, available, visibleFlexPool, fixed float64, flexCount int) (float64, error) {
	if math.IsNaN(available) || math.IsInf(available, 0) {
		return 0, newLayoutErrorV1EngineLayoutValidation(node, "gap and margins leave no usable %s (%.6g)", axis, available)
	}
	if math.IsNaN(fixed) || math.IsInf(fixed, 0) {
		return 0, newLayoutErrorV1EngineLayoutValidation(node, "fixed child %s total must be finite", axis)
	}
	if available <= geometryEpsilonV1EngineLayoutValidation {
		if normalizedOverflowV1EngineLayoutConstraints(node) != entity.OverflowVisible {
			return 0, newLayoutErrorV1EngineLayoutValidation(node, "gap and margins leave no usable %s (%.6g)", axis, available)
		}
		if flexCount == 0 {
			return 0, nil
		}
		if !isPositiveFiniteV1EngineLayoutConstraints(visibleFlexPool) {
			return 0, newLayoutErrorV1EngineLayoutValidation(node, "visible overflow flex pool must be positive and finite, got %.6g", visibleFlexPool)
		}
		return visibleFlexPool, nil
	}
	remaining := available - fixed
	if remaining < -geometryEpsilonV1EngineLayoutValidation {
		if normalizedOverflowV1EngineLayoutConstraints(node) != entity.OverflowVisible {
			return 0, newLayoutErrorV1EngineLayoutValidation(node, "fixed child %s total %.6g exceeds the available %.6g; set overflow=\"visible\" to allow it explicitly", axis, fixed, available)
		}
		if flexCount > 0 {
			return available, nil
		}
		return 0, nil
	}
	if flexCount > 0 && remaining <= geometryEpsilonV1EngineLayoutValidation {
		if normalizedOverflowV1EngineLayoutConstraints(node) == entity.OverflowVisible {
			return available, nil
		}
		return 0, newLayoutErrorV1EngineLayoutValidation(node, "fixed children leave no positive %s for %d flexible children", axis, flexCount)
	}
	if remaining < 0 {
		remaining = 0
	}
	return remaining, nil
}

func isPositiveFiniteV1EngineLayoutConstraints(value float64) bool {
	return value > 0 && !math.IsNaN(value) && !math.IsInf(value, 0)
}

func setContentBoxV1EngineLayoutConstraints(target *entity.Box, x, y, w, h float64) {
	target.ContentX = x
	target.ContentY = y
	target.ContentW = w
	target.ContentH = h
}

// layoutKids returns node's children that participate in layout,
