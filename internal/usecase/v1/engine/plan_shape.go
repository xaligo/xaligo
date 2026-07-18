package engine

import "github.com/xaligo/xaligo/internal/entity"

func shapeOpV1EnginePlanShape(el *entity.Element, frame rectV1EngineRouteTypes, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPosV1EnginePlanGeometry(el, frame, ppi)
	if !ok {
		return entity.DrawOp{}, false
	}
	kind := "rect"
	if el.Type == "ellipse" {
		kind = "ellipse"
	} else if el.Type == "diamond" {
		kind = "diamond"
	}
	ln := linePropsV1EnginePlanStyle(el, ppi)
	fl := fillPropsV1EnginePlanStyle(el.BackgroundColor, opacityToTransparencyV1EnginePlanStyle(el.Opacity))
	return entity.DrawOp{
		ID:         el.ID,
		Kind:       kind,
		X:          p.X,
		Y:          p.Y,
		W:          p.W,
		H:          p.H,
		Rotate:     el.Angle,
		Line:       &ln,
		Fill:       &fl,
		FrontLayer: el.CustomData != nil && el.CustomData.FrameMetadata,
	}, true
}

func polygonOpV1EnginePlanShape(el *entity.Element, frame rectV1EngineRouteTypes, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPosV1EnginePlanGeometry(el, frame, ppi)
	if !ok || len(el.Points) < 3 {
		return entity.DrawOp{}, false
	}
	points := make([]entity.PtIn, 0, len(el.Points))
	for i, point := range el.Points {
		if len(point) < 2 {
			continue
		}
		points = append(points, entity.PtIn{X: point[0] / ppi, Y: point[1] / ppi, MoveTo: i == 0})
	}
	if len(points) < 3 {
		return entity.DrawOp{}, false
	}
	ln := linePropsV1EnginePlanStyle(el, ppi)
	fl := fillPropsV1EnginePlanStyle(el.BackgroundColor, opacityToTransparencyV1EnginePlanStyle(el.Opacity))
	return entity.DrawOp{Kind: "polygon", X: p.X, Y: p.Y, W: p.W, H: p.H, Rotate: el.Angle, Points: points, Line: &ln, Fill: &fl}, true
}
