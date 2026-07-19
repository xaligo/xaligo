package engine

import (
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

func polylineOpV1EnginePlanConnectorDraw(el *entity.Element, points []ptV1EngineRouteTypes, frame rectV1EngineRouteTypes, ppi float64, style connectorStyleV1EnginePlanConnectorStyle) (entity.DrawOp, bool) {
	if len(points) < 2 {
		return entity.DrawOp{}, false
	}
	inch := make([]ptV1EngineRouteTypes, len(points))
	for i, p := range points {
		inch[i] = ptV1EngineRouteTypes{X: (p.X - frame.X) / ppi, Y: (p.Y - frame.Y) / ppi}
	}
	minX, minY := inch[0].X, inch[0].Y
	maxX, maxY := inch[0].X, inch[0].Y
	for _, p := range inch {
		minX = math.Min(minX, p.X)
		minY = math.Min(minY, p.Y)
		maxX = math.Max(maxX, p.X)
		maxY = math.Max(maxY, p.Y)
	}
	w := math.Max(maxX-minX, 0.0001)
	h := math.Max(maxY-minY, 0.0001)
	rel := make([]entity.PtIn, len(inch))
	for i, p := range inch {
		rel[i] = entity.PtIn{X: p.X - minX, Y: p.Y - minY, MoveTo: i == 0}
	}
	ln := connectorLineV1EnginePlanConnectorDraw(el, style, ppi)
	applyPPTXArrowHeadExtensionV1EnginePlanConnectorDraw(&ln, ppi)
	return entity.DrawOp{
		ID:         el.ID,
		FrontLayer: true,
		Kind:       "line",
		X:          minX,
		Y:          minY,
		W:          w,
		H:          h,
		Points:     rel,
		Line:       &ln,
	}, true
}

func rawLineOpV1EnginePlanConnectorDraw(el *entity.Element, frame rectV1EngineRouteTypes, ppi float64, style connectorStyleV1EnginePlanConnectorStyle) (entity.DrawOp, bool) {
	startX := el.X - frame.X
	startY := el.Y - frame.Y
	points := el.Points
	if rawLineNeedsPolylineV1EnginePlanConnectorDraw(points) {
		absolute := make([]ptV1EngineRouteTypes, 0, len(points))
		for _, point := range points {
			if len(point) < 2 {
				continue
			}
			absolute = append(absolute, ptV1EngineRouteTypes{X: el.X + point[0], Y: el.Y + point[1]})
		}
		return polylineOpV1EnginePlanConnectorDraw(el, absolute, frame, ppi, style)
	}
	if len(points) == 0 {
		points = [][]float64{{0, 0}, {el.Width, el.Height}}
	}
	endPoint := points[len(points)-1]
	dx, dy := 0.0, 0.0
	if len(endPoint) >= 2 {
		dx, dy = endPoint[0], endPoint[1]
	}
	x := math.Min(startX, startX+dx) / ppi
	y := math.Min(startY, startY+dy) / ppi
	w := math.Abs(dx) / ppi
	h := math.Abs(dy) / ppi
	if w <= 0 && h <= 0 {
		return entity.DrawOp{}, false
	}
	ln := connectorLineV1EnginePlanConnectorDraw(el, style, ppi)
	applyPPTXArrowHeadExtensionV1EnginePlanConnectorDraw(&ln, ppi)
	return entity.DrawOp{
		ID:         el.ID,
		FrontLayer: true,
		Kind:       "line",
		X:          x,
		Y:          y,
		W:          w,
		H:          h,
		FlipH:      dx < 0,
		FlipV:      dy < 0,
		Line:       &ln,
	}, true
}

func rawLineNeedsPolylineV1EnginePlanConnectorDraw(points [][]float64) bool {
	if len(points) > 2 {
		return true
	}
	if len(points) < 2 || len(points[0]) < 2 {
		return false
	}
	return points[0][0] != 0 || points[0][1] != 0
}

func connectorLineV1EnginePlanConnectorDraw(el *entity.Element, style connectorStyleV1EnginePlanConnectorStyle, ppi float64) entity.LineStyle {
	base := linePropsV1EnginePlanStyle(el, ppi)
	kind := connectorKindV1EnginePlanConnectorDraw(el)
	beginHead, endHead := connectorArrowheadsV1EnginePlanConnectorDraw(el)
	width := base.Width
	switch kind {
	case "route":
		beginHead = "none"
		endHead = "none"
	case "traffic":
		if endHead == "" {
			endHead = style.Head
		}
	default:
		if endHead == "" && el.Type == "arrow" {
			endHead = style.Head
		}
	}
	if beginHead == "" {
		beginHead = "none"
	}
	if endHead == "" {
		endHead = "none"
	}
	widthExplicit := el.CustomData != nil && el.CustomData.ConnectorStyleSourceKnown && el.CustomData.ConnectorStrokeWidthExplicit
	if style.HasWidth && !widthExplicit {
		width = pxToPtV1EnginePlanStyle(style.Width, ppi)
	}
	base.Width = width
	base.BeginArrowType = beginHead
	base.EndArrowType = endHead
	return base
}

func applyPPTXArrowHeadExtensionV1EnginePlanConnectorDraw(line *entity.LineStyle, ppi float64) {
	if line == nil || ppi <= 0 {
		return
	}
	extension := pptxArrowHeadExtendPxV1EnginePlanBuild / ppi
	if line.BeginArrowType != "" && line.BeginArrowType != "none" {
		line.BeginArrowExtendIn = extension
	}
	if line.EndArrowType != "" && line.EndArrowType != "none" {
		line.EndArrowExtendIn = extension
	}
}

func connectorKindV1EnginePlanConnectorDraw(el *entity.Element) string {
	if el.CustomData == nil {
		return "connection"
	}
	switch el.CustomData.ConnectorKind {
	case "route", "traffic":
		return el.CustomData.ConnectorKind
	default:
		return "connection"
	}
}

func connectorKindPriorityV1EnginePlanConnectorDraw(kind string) int {
	switch kind {
	case "route":
		return 0
	case "traffic":
		return 2
	default:
		return 1
	}
}

func connectorArrowheadsV1EnginePlanConnectorDraw(el *entity.Element) (string, string) {
	if el.CustomData == nil {
		return "", ""
	}
	if !el.CustomData.ConnectorStyleSourceKnown {
		// Scenes created before source-style metadata was introduced stored only
		// resolved heads. Preserve their historical interpretation.
		return el.CustomData.ConnectorStartArrowhead, el.CustomData.ConnectorEndArrowhead
	}
	beginHead, endHead := "", ""
	if el.CustomData.ConnectorStartArrowheadExplicit {
		beginHead = el.CustomData.ConnectorStartArrowhead
	}
	if el.CustomData.ConnectorEndArrowheadExplicit {
		endHead = el.CustomData.ConnectorEndArrowhead
	}
	return beginHead, endHead
}
