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
	return BuildDocumentPlanWithIcons(document, ppi, nil)
}

// BuildDocumentPlanWithIcons projects resolved geometry and already-resolved
// icon assets without asking the output adapter to calculate placement.
func BuildDocumentPlanWithIcons(document entity.EngineResolvedDocument, ppi float64, icons map[string]string) (entity.DocumentPlan, error) {
	if ppi <= 0 || math.IsNaN(ppi) || math.IsInf(ppi, 0) {
		ppi = 96
	}
	page := entity.DocumentPage{
		ID: "v2", Slide: entity.PlanSlide{W: document.Width / ppi, H: document.Height / ppi, Background: "FFFFFF", CropToSlide: true},
	}
	for _, element := range document.Elements {
		ops, err := resolvedElementOps(element, ppi, icons)
		if err != nil {
			return entity.DocumentPlan{}, err
		}
		page.Ops = append(page.Ops, ops...)
	}
	return entity.DocumentPlan{SchemaVersion: 2, Pages: []entity.DocumentPage{page}}, nil
}

func resolvedElementOps(element entity.EngineResolvedElement, ppi float64, icons map[string]string) ([]entity.DrawOp, error) {
	if !element.Visual.Visible {
		return nil, nil
	}
	base := entity.DrawOp{ID: element.ID, X: element.X / ppi, Y: element.Y / ppi, W: element.Width / ppi, H: element.Height / ppi}
	line := &entity.LineStyle{Color: planColor(element.Visual.Stroke, "1F2937"), Width: element.Visual.StrokeWidth * 72 / ppi, Dash: string(element.Line.Style)}
	fill := &entity.FillStyle{Color: planColor(element.Visual.Fill, "FFFFFF"), Transparency: (1 - element.Visual.Opacity) * 100}
	header, hasV1GroupHeader := resolvedV1GroupHeaderGeometry(element)
	ops := make([]entity.DrawOp, 0, 3)
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
		if hasV1GroupHeader {
			borderY := header.y + header.height/2
			base.Y = borderY / ppi
			base.H = math.Max(0, element.Y+element.Height-borderY) / ppi
		}
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
	if hasV1GroupHeader {
		tip := math.Min(14, header.height/2)
		ops = append(ops, entity.DrawOp{
			ID: element.ID + "-header", FrontLayer: true, Kind: "polygon",
			X: header.x / ppi, Y: header.y / ppi, W: header.width / ppi, H: header.height / ppi,
			Points: []entity.PtIn{
				{X: 0, Y: 0, MoveTo: true},
				{X: (header.width - tip) / ppi, Y: 0},
				{X: header.width / ppi, Y: header.height / 2 / ppi},
				{X: (header.width - tip) / ppi, Y: header.height / ppi},
				{X: 0, Y: header.height / ppi},
			},
			Line: &entity.LineStyle{Color: planColor(element.Visual.Stroke, "1F2937"), Width: 72 / ppi, Dash: string(entity.EngineLineSolid)},
			Fill: fill,
		})
	}
	if data := icons[element.IconRef]; data != "" && element.IconWidth > 0 && element.IconHeight > 0 {
		ops = append(ops, entity.DrawOp{
			ID: element.ID + "-icon", GroupID: element.ID, FrontLayer: true, Kind: "image",
			X: element.IconX / ppi, Y: element.IconY / ppi,
			W: element.IconWidth / ppi, H: element.IconHeight / ppi, Data: data,
		})
	}
	if element.Text.Value != "" {
		textX, textY, textWidth, textHeight := element.Text.X, element.Text.Y, element.Text.Width, element.Text.Height
		valign := "mid"
		align := "center"
		wrap := true
		if element.Concept == entity.EngineConceptItem && element.IconRef != "" {
			valign = "top"
		}
		if hasV1GroupHeader {
			align = "left"
			wrap = false
		}
		if textWidth <= 0 || textHeight <= 0 {
			textX, textY, textWidth, textHeight = element.X, element.Y, element.Width, element.Height
		}
		text := entity.DrawOp{
			ID: element.ID + "-text", GroupID: element.ID, FrontLayer: true, Kind: "text",
			X: textX / ppi, Y: textY / ppi, W: textWidth / ppi, H: textHeight / ppi,
			Text: element.Text.Value, Color: planColor(element.Text.Color, "111827"),
			FontFace: element.Text.FontFamily, FontSize: element.Text.FontSize * 72 / ppi, Align: align, Valign: valign,
			TextLayout: &entity.TextLayout{Role: entity.TextRole(element.Text.Role), Wrap: wrap, LineHeight: element.Text.LineHeight},
		}
		ops = append(ops, text)
	}
	if len(ops) == 0 {
		if element.Visual.Shape == entity.EngineShapeNone {
			return nil, nil
		}
		return nil, fmt.Errorf("resolved element %q has no projectable representation", element.ID)
	}
	return ops, nil
}

type v1GroupHeaderGeometry struct {
	x      float64
	y      float64
	width  float64
	height float64
}

const v1GroupHeaderEndPadding = 4.0

func resolvedV1GroupHeaderGeometry(element entity.EngineResolvedElement) (v1GroupHeaderGeometry, bool) {
	if element.Text.Role != string(entity.TextRoleGroupHeader) || (element.Concept != entity.EngineConceptGroup && element.Concept != entity.EngineConceptCapture) {
		return v1GroupHeaderGeometry{}, false
	}
	hasIcon := element.IconWidth > 0 && element.IconHeight > 0
	height := 20.0
	if hasIcon {
		height = math.Max(height, element.IconHeight)
	}
	x := element.Text.X - 4
	y := element.Text.Y - (height-element.Text.Height)/2
	if hasIcon {
		x = element.IconX
		y = element.IconY - (height-element.IconHeight)/2
	}
	tip := math.Min(14, height/2)
	return v1GroupHeaderGeometry{
		x: x, y: y, width: element.Text.X + element.Text.Width + v1GroupHeaderEndPadding + tip - x, height: height,
	}, true
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
