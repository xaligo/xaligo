package v2

import (
	"fmt"
	"math"

	"github.com/xaligo/xaligo/internal/entity"
)

// BuildDocumentPlan projects one immutable V2 result into the existing
// renderer-neutral document plan. It performs unit conversion only; geometry,
// text sizing, and routes remain owned by the Rust calculation result.
func BuildDocumentPlan(document entity.EngineResolvedDocument, ppi float64) (entity.DocumentPlan, error) {
	if ppi <= 0 || math.IsNaN(ppi) || math.IsInf(ppi, 0) {
		ppi = 96
	}
	page := entity.DocumentPage{
		ID: "v2", Slide: entity.PlanSlide{W: document.Width / ppi, H: document.Height / ppi, Background: "FFFFFF", CropToSlide: true},
	}
	for _, element := range document.Elements {
		ops, err := resolvedElementOps(element, ppi)
		if err != nil {
			return entity.DocumentPlan{}, err
		}
		page.Ops = append(page.Ops, ops...)
	}
	return entity.DocumentPlan{SchemaVersion: 2, Pages: []entity.DocumentPage{page}}, nil
}

func resolvedElementOps(element entity.EngineResolvedElement, ppi float64) ([]entity.DrawOp, error) {
	if !element.Visual.Visible {
		return nil, nil
	}
	base := entity.DrawOp{ID: element.ID, X: element.X / ppi, Y: element.Y / ppi, W: element.Width / ppi, H: element.Height / ppi}
	line := &entity.LineStyle{Color: planColor(element.Visual.Stroke, "1F2937"), Width: element.Visual.StrokeWidth * 72 / ppi, Dash: string(element.Line.Style)}
	fill := &entity.FillStyle{Color: planColor(element.Visual.Fill, "FFFFFF"), Transparency: (1 - element.Visual.Opacity) * 100}
	ops := make([]entity.DrawOp, 0, 2)
	switch element.Concept {
	case entity.EngineConceptLine:
		base.Kind = "line"
		base.Line = line
		base.Points = make([]entity.PtIn, len(element.Points))
		if len(element.Points) > 0 {
			base.X, base.Y = element.Points[0].X/ppi, element.Points[0].Y/ppi
			for index, point := range element.Points {
				base.Points[index] = entity.PtIn{X: point.X/ppi - base.X, Y: point.Y/ppi - base.Y, MoveTo: index == 0}
			}
		}
		line.BeginArrowType = planDecoration(element.Line.SourceDecoration)
		line.EndArrowType = planDecoration(element.Line.TargetDecoration)
		ops = append(ops, base)
	case entity.EngineConceptSpacer:
		return nil, nil
	default:
		switch element.Visual.Shape {
		case entity.EngineShapeEllipse:
			base.Kind = "ellipse"
		case entity.EngineShapeNone:
			base.Kind = ""
		default:
			base.Kind = "rect"
		}
		if base.Kind != "" {
			base.Line, base.Fill = line, fill
			ops = append(ops, base)
		}
	}
	if element.Text.Value != "" {
		text := entity.DrawOp{
			ID: element.ID + "-text", GroupID: element.ID, Kind: "text",
			X: element.X / ppi, Y: element.Y / ppi, W: element.Width / ppi, H: element.Height / ppi,
			Text: element.Text.Value, Color: planColor(element.Text.Color, "111827"),
			FontFace: element.Text.FontFamily, FontSize: element.Text.FontSize, Align: "center", Valign: "mid",
			TextLayout: &entity.TextLayout{Role: entity.TextRole(element.Text.Role), Wrap: true, LineHeight: element.Text.LineHeight},
		}
		ops = append(ops, text)
	}
	if len(ops) == 0 {
		return nil, fmt.Errorf("resolved element %q has no projectable representation", element.ID)
	}
	return ops, nil
}

func planColor(value, fallback string) string {
	if len(value) == 7 && value[0] == '#' {
		return value[1:]
	}
	if len(value) == 6 {
		return value
	}
	return fallback
}

func planDecoration(value entity.EngineDecoration) string {
	switch value {
	case entity.EngineDecorationArrow:
		return "arrow"
	case entity.EngineDecorationTriangle:
		return "triangle"
	case entity.EngineDecorationDiamond:
		return "diamond"
	case entity.EngineDecorationCircle:
		return "oval"
	default:
		return ""
	}
}
