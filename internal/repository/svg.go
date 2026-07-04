package repository

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const svgDefaultPxPerInch = 96.0

type SVGRepository interface {
	Render(plan entity.Plan, pxPerInch float64, legendPosition string) ([]byte, error)
}

type svgRepository struct{}

func NewSVGRepository() SVGRepository { return &svgRepository{} }

// RenderPlan converts the shared draw plan into an SVG document.
func (rcvr *svgRepository) Render(plan entity.Plan, pxPerInch float64, legendPosition string) ([]byte, error) {
	if pxPerInch <= 0 {
		pxPerInch = svgDefaultPxPerInch
	}
	if plan.Slide.W <= 0 || plan.Slide.H <= 0 {
		return nil, fmt.Errorf("SVG slide size must be positive")
	}

	w := plan.Slide.W * pxPerInch
	h := plan.Slide.H * pxPerInch
	layout := svgLegendLayout(w, h, plan.Legend, legendPosition)

	var b bytes.Buffer
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%s" height="%s" viewBox="0 0 %s %s">`+"\n", num(layout.canvasW), num(layout.canvasH), num(layout.canvasW), num(layout.canvasH))
	b.WriteString(`<defs><marker id="xaligo-arrow" markerWidth="10" markerHeight="10" refX="9" refY="5" orient="auto-start-reverse" markerUnits="strokeWidth"><path d="M 0 0 L 10 5 L 0 10 z" fill="context-stroke"/></marker><marker id="xaligo-oval" markerWidth="7" markerHeight="7" refX="3.5" refY="3.5" orient="auto" markerUnits="strokeWidth"><circle cx="3.5" cy="3.5" r="2.5" fill="context-stroke"/></marker></defs>` + "\n")
	fmt.Fprintf(&b, `<rect x="0" y="0" width="%s" height="%s" fill="#%s"/>`+"\n", num(layout.canvasW), num(layout.canvasH), color(plan.Slide.Background, "FFFFFF"))

	if layout.diagramX != 0 || layout.diagramY != 0 {
		fmt.Fprintf(&b, `<g transform="translate(%s %s)">`+"\n", num(layout.diagramX), num(layout.diagramY))
	}
	for _, op := range plan.Ops {
		writeOp(&b, op, pxPerInch)
	}
	if layout.diagramX != 0 || layout.diagramY != 0 {
		b.WriteString("</g>\n")
	}
	writeLegend(&b, plan.Legend, layout)
	b.WriteString("</svg>\n")
	return b.Bytes(), nil
}

type svgLegendBox struct {
	canvasW  float64
	canvasH  float64
	diagramX float64
	diagramY float64
	legendX  float64
	legendY  float64
	legendW  float64
	legendH  float64
	visible  bool
}

func svgLegendLayout(w, h float64, entries []entity.LegendEntry, position string) svgLegendBox {
	layout := svgLegendBox{canvasW: w, canvasH: h}
	if len(entries) == 0 {
		return layout
	}
	position = strings.ToLower(strings.TrimSpace(position))
	if position == "" {
		position = "bottom"
	}
	const (
		gap       = 24.0
		pad       = 16.0
		rowH      = 38.0
		sideW     = 280.0
		bottomMin = 360.0
	)
	rows := float64(len(entries))
	layout.visible = true
	switch position {
	case "top":
		layout.legendW = math.Max(bottomMin, w)
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = math.Max(w, layout.legendW)
		layout.canvasH = h + gap + layout.legendH
		layout.diagramY = layout.legendH + gap
	case "right":
		layout.legendW = sideW
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = w + gap + layout.legendW
		layout.canvasH = math.Max(h, layout.legendH)
		layout.legendX = w + gap
	case "left":
		layout.legendW = sideW
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = w + gap + layout.legendW
		layout.canvasH = math.Max(h, layout.legendH)
		layout.diagramX = layout.legendW + gap
	default:
		layout.legendW = math.Max(bottomMin, w)
		layout.legendH = pad*2 + rows*rowH
		layout.canvasW = math.Max(w, layout.legendW)
		layout.canvasH = h + gap + layout.legendH
		layout.legendY = h + gap
	}
	return layout
}

func writeLegend(b *bytes.Buffer, entries []entity.LegendEntry, layout svgLegendBox) {
	if !layout.visible {
		return
	}
	const (
		pad      = 16.0
		rowH     = 38.0
		iconSize = 24.0
	)
	fmt.Fprintf(b, `<g id="xaligo-svg-legend" transform="translate(%s %s)">`+"\n", num(layout.legendX), num(layout.legendY))
	fmt.Fprintf(b, `<rect x="0" y="0" width="%s" height="%s" fill="#FFFFFF" stroke="#CBD5E1" stroke-width="1"/>`+"\n", num(layout.legendW), num(layout.legendH))
	for i, entry := range entries {
		y := pad + float64(i)*rowH
		if entry.Data != "" {
			fmt.Fprintf(b, `<image x="%s" y="%s" width="%s" height="%s" href="%s"/>`+"\n", num(pad), num(y), num(iconSize), num(iconSize), attr(entry.Data))
		}
		textX := pad + iconSize + 10
		fmt.Fprintf(b, `<text x="%s" y="%s" fill="#0F172A" font-family="Arial" font-size="12" font-weight="700">%s</text>`+"\n", num(textX), num(y+11), text(entry.Abbreviation))
		fmt.Fprintf(b, `<text x="%s" y="%s" fill="#475569" font-family="Arial" font-size="10">%s</text>`+"\n", num(textX), num(y+25), text(entry.OfficialName))
	}
	b.WriteString("</g>\n")
}

func writeOp(b *bytes.Buffer, op entity.DrawOp, ppi float64) {
	x, y, w, h := op.X*ppi, op.Y*ppi, op.W*ppi, op.H*ppi
	transform := rotateAttr(op.Rotate, x+w/2, y+h/2)
	switch op.Kind {
	case "rect":
		fmt.Fprintf(b, `<rect x="%s" y="%s" width="%s" height="%s"%s%s%s/>`+"\n", num(x), num(y), num(w), num(h), fillAttrs(op.Fill), lineAttrs(op.Line), transform)
	case "ellipse":
		fmt.Fprintf(b, `<ellipse cx="%s" cy="%s" rx="%s" ry="%s"%s%s%s/>`+"\n", num(x+w/2), num(y+h/2), num(w/2), num(h/2), fillAttrs(op.Fill), lineAttrs(op.Line), transform)
	case "polygon":
		writePolygon(b, op, x, y, w, h, ppi, transform)
	case "text":
		writeText(b, op, x, y, w, h, ppi, transform)
	case "image":
		if op.Data == "" {
			return
		}
		fmt.Fprintf(b, `<image x="%s" y="%s" width="%s" height="%s" href="%s" opacity="%s"%s/>`+"\n", num(x), num(y), num(w), num(h), attr(op.Data), num(opacity(op.Transparency)), transform)
	case "line":
		writeLine(b, op, x, y, w, h, ppi)
	}
}

func writePolygon(b *bytes.Buffer, op entity.DrawOp, x, y, w, h, ppi float64, transform string) {
	points := absolutePoints(op, x, y, w, h, ppi)
	if len(points) < 3 {
		return
	}
	var value strings.Builder
	for i, p := range points {
		if i > 0 {
			value.WriteByte(' ')
		}
		fmt.Fprintf(&value, "%s,%s", num(p.x), num(p.y))
	}
	fmt.Fprintf(b, `<polygon points="%s"%s%s%s/>`+"\n", value.String(), fillAttrs(op.Fill), lineAttrs(op.Line), transform)
}

func writeText(b *bytes.Buffer, op entity.DrawOp, x, y, w, h, ppi float64, transform string) {
	if op.Text == "" {
		return
	}
	anchor := "start"
	textX := x
	if op.Align == "center" {
		anchor = "middle"
		textX = x + w/2
	} else if op.Align == "right" {
		anchor = "end"
		textX = x + w
	}
	fontSize := op.FontSize * ppi / 72.0
	textY := y + fontSize
	if op.Valign == "middle" {
		textY = y + h/2 + fontSize*0.35
	} else if op.Valign == "bottom" {
		textY = y + h - 2
	}
	weight := ""
	if op.Bold {
		weight = ` font-weight="700"`
	}
	fmt.Fprintf(b, `<text x="%s" y="%s" fill="#%s" font-family="%s" font-size="%s" text-anchor="%s"%s%s>%s</text>`+"\n",
		num(textX), num(textY), color(op.Color, "1E1E1E"), attr(op.FontFace), num(fontSize), anchor, weight, transform, text(op.Text))
}

func writeLine(b *bytes.Buffer, op entity.DrawOp, x, y, w, h, ppi float64) {
	points := absolutePoints(op, x, y, w, h, ppi)
	if len(points) < 2 {
		return
	}
	var d strings.Builder
	for i, p := range points {
		cmd := "L"
		if i == 0 {
			cmd = "M"
		}
		fmt.Fprintf(&d, "%s %s %s ", cmd, num(p.x), num(p.y))
	}
	marker := ""
	if op.Line != nil && op.Line.BeginArrowType != "" && op.Line.BeginArrowType != "none" {
		marker += ` marker-start="url(#` + markerID(op.Line.BeginArrowType) + `)"`
	}
	if op.Line != nil && op.Line.EndArrowType != "" && op.Line.EndArrowType != "none" {
		marker += ` marker-end="url(#` + markerID(op.Line.EndArrowType) + `)"`
	}
	fmt.Fprintf(b, `<path d="%s" fill="none"%s%s/>`+"\n", strings.TrimSpace(d.String()), lineAttrs(op.Line), marker)
}

func markerID(arrowType string) string {
	if arrowType == "oval" {
		return "xaligo-oval"
	}
	return "xaligo-arrow"
}

type point struct{ x, y float64 }

func absolutePoints(op entity.DrawOp, x, y, w, h, ppi float64) []point {
	if len(op.Points) > 0 {
		out := make([]point, 0, len(op.Points))
		for _, p := range op.Points {
			out = append(out, point{x + p.X*ppi, y + p.Y*ppi})
		}
		return out
	}
	x1, x2 := x, x+w
	y1, y2 := y, y+h
	if op.FlipH {
		x1, x2 = x2, x1
	}
	if op.FlipV {
		y1, y2 = y2, y1
	}
	return []point{{x1, y1}, {x2, y2}}
}

func fillAttrs(fill *entity.FillStyle) string {
	if fill == nil || fill.Transparency >= 100 {
		return ` fill="none"`
	}
	return fmt.Sprintf(` fill="#%s" fill-opacity="%s"`, color(fill.Color, "FFFFFF"), num(opacity(fill.Transparency)))
}

func lineAttrs(line *entity.LineStyle) string {
	if line == nil || line.Transparency >= 100 || line.Width <= 0 {
		return ` stroke="none"`
	}
	dash := ""
	if line.Dash == "dash" {
		dash = ` stroke-dasharray="8 6"`
	} else if line.Dash == "dot" {
		dash = ` stroke-dasharray="2 5" stroke-linecap="round"`
	}
	return fmt.Sprintf(` stroke="#%s" stroke-width="%s" stroke-opacity="%s"%s`, color(line.Color, "1E1E1E"), num(line.Width), num(opacity(line.Transparency)), dash)
}

func rotateAttr(deg, cx, cy float64) string {
	if math.Abs(deg) < 0.0001 {
		return ""
	}
	return fmt.Sprintf(` transform="rotate(%s %s %s)"`, num(deg), num(cx), num(cy))
}

func opacity(transparency float64) float64 {
	return math.Max(0, math.Min(1, 1-transparency/100.0))
}

func color(v, fallback string) string {
	v = strings.TrimPrefix(strings.TrimSpace(v), "#")
	if len(v) == 6 {
		return strings.ToUpper(v)
	}
	return fallback
}

func num(v float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.3f", v), "0"), ".")
}

func attr(v string) string {
	var b bytes.Buffer
	_ = xml.EscapeText(&b, []byte(v))
	return b.String()
}

func text(v string) string {
	return attr(v)
}
