package engine

import "github.com/xaligo/xaligo/internal/entity"

func imageOpV1EnginePlanImage(el *entity.Element, files map[string]entity.SceneFile, frame rectV1EngineRouteTypes, ppi float64) (entity.DrawOp, bool) {
	p, ok := toPosV1EnginePlanGeometry(el, frame, ppi)
	if !ok || el.FileID == "" {
		return entity.DrawOp{}, false
	}
	f, ok := files[el.FileID]
	if !ok || f.DataURL == "" {
		return entity.DrawOp{}, false
	}
	return entity.DrawOp{
		ID:           el.ID,
		Kind:         "image",
		X:            p.X,
		Y:            p.Y,
		W:            p.W,
		H:            p.H,
		Rotate:       el.Angle,
		Data:         f.DataURL,
		Transparency: opacityToTransparencyV1EnginePlanStyle(el.Opacity),
	}, true
}
