import pptxgen from 'pptxgenjs';

import type {
  ConnectorLegendRow,
  PlanConnectorLegendEntry,
  PlanLine,
  PptxPlan,
} from '../entity/pptx';
import { imageDataForPptx } from './pptx_image';

type LegendLineOptions = pptxgen.ShapeLineProps;

export function drawConnectorLegendSlide(pptx: pptxgen, plan: PptxPlan): void {
  const rows = groupConnectorLegendEntries(plan.connectorLegend ?? []);
  if (rows.length === 0) return;

  const slideW = plan.slide.w;
  const slideH = plan.slide.h;
  const marginX = 0.55;
  const marginTop = 0.42;
  const marginBottom = 0.35;
  const titleH = 0.35;
  const headerH = 0.24;
  const rowH = 0.31;
  const rowsPerSlide = Math.max(1, Math.floor((slideH - marginTop - marginBottom - titleH - headerH - 0.18) / rowH));
  const sampleW = 0.82;
  const lineX = marginX;
  const typeX = lineX + sampleW + 0.15;
  const typeW = 0.92;
  const idX = typeX + typeW + 0.15;
  const idW = 1.45;
  const styleX = idX + idW + 0.15;
  const styleW = 2.15;
  const descriptionX = styleX + styleW + 0.18;
  const descriptionW = Math.max(1, slideW - marginX - descriptionX);
  const rowTextH = 0.18;

  for (let start = 0; start < rows.length; start += rowsPerSlide) {
    const pageRows = rows.slice(start, start + rowsPerSlide);
    const slide = pptx.addSlide();
    slide.background = { color: plan.slide.background || 'FFFFFF' };

    slide.addText('Line Legend', {
      x: marginX,
      y: marginTop,
      w: slideW - marginX * 2,
      h: titleH,
      fontFace: 'Helvetica',
      fontSize: 16,
      bold: true,
      color: '1E1E1E',
      margin: 0,
    });

    const headerY = marginTop + titleH;
    slide.addText('Line', { x: lineX, y: headerY, w: sampleW, h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });
    slide.addText('Type', { x: typeX, y: headerY, w: typeW, h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });
    slide.addText('ID', { x: idX, y: headerY, w: idW, h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });
    slide.addText('Style', { x: styleX, y: headerY, w: styleW, h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });
    slide.addText('Description', { x: descriptionX, y: headerY, w: descriptionW, h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });

    for (const [i, entry] of pageRows.entries()) {
      const y = marginTop + titleH + headerH + 0.06 + i * rowH;
      const rowCenterY = y + rowH / 2;
      const rowTextY = rowCenterY - rowTextH / 2;
      const line: pptxgen.ShapeLineProps = lineOptions(entry.line);
      slide.addShape(pptx.ShapeType.line, {
        x: lineX,
        y: rowCenterY,
        w: sampleW,
        h: 0,
        line,
      });
      slide.addText(entry.label || entry.kind, {
        x: typeX,
        y: rowTextY,
        w: typeW,
        h: rowTextH,
        fontFace: 'Helvetica',
        fontSize: 6.5,
        bold: true,
        color: '1E1E1E',
        margin: 0,
        fit: 'shrink',
      });
      slide.addText(formatConnectorLegendIDs(entry.ids), {
        x: idX,
        y: rowTextY,
        w: idW,
        h: rowTextH,
        fontFace: 'Helvetica',
        fontSize: 6.2,
        bold: true,
        color: entry.line.color || '1E1E1E',
        margin: 0,
        fit: 'shrink',
        breakLine: false,
      });
      slide.addText(lineStyleSummary(entry.line), {
        x: styleX,
        y: rowTextY,
        w: styleW,
        h: rowTextH,
        fontFace: 'Helvetica',
        fontSize: 6.2,
        color: '444444',
        margin: 0,
        fit: 'shrink',
        breakLine: false,
      });
      slide.addText(entry.description || '', {
        x: descriptionX,
        y: rowTextY,
        w: descriptionW,
        h: rowTextH,
        fontFace: 'Helvetica',
        fontSize: 6.2,
        color: '444444',
        margin: 0,
        fit: 'shrink',
        breakLine: false,
      });
    }
  }
}

export async function drawLegendSlides(pptx: pptxgen, plan: PptxPlan): Promise<void> {
  const entries = (plan.legend ?? []).filter((e) => e.data && e.officialName);
  if (entries.length === 0) return;

  const slideW = plan.slide.w;
  const slideH = plan.slide.h;
  const marginX = 0.55;
  const marginTop = 0.42;
  const marginBottom = 0.35;
  const titleH = 0.35;
  const headerH = 0.24;
  const rowH = 0.32;
  const usableH = slideH - marginTop - marginBottom - titleH - headerH;
  const rowsPerCol = Math.max(1, Math.floor(usableH / rowH));
  const colsPerSlide = 4;
  const entriesPerSlide = rowsPerCol * colsPerSlide;

  for (let start = 0; start < entries.length; start += entriesPerSlide) {
    const pageEntries = entries.slice(start, start + entriesPerSlide);
    const cols = colsPerSlide;
    const colW = (slideW - marginX * 2) / cols;
    const slide = pptx.addSlide();
    slide.background = { color: plan.slide.background || 'FFFFFF' };

    slide.addText('Legend', {
      x: marginX,
      y: marginTop,
      w: slideW - marginX * 2,
      h: titleH,
      fontFace: 'Helvetica',
      fontSize: 16,
      bold: true,
      color: '1E1E1E',
      margin: 0,
    });

    for (let col = 0; col < cols; col++) {
      const x = marginX + col * colW;
      const y = marginTop + titleH;
      slide.addText('Icon', { x, y, w: 0.38, h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });
      slide.addText('Abbr.', { x: x + 0.46, y, w: 0.55, h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });
      slide.addText('Official name', { x: x + 1.05, y, w: Math.max(0.5, colW - 1.08), h: headerH, fontSize: 7, bold: true, color: '666666', margin: 0 });
    }

    for (const [i, entry] of pageEntries.entries()) {
      const col = Math.floor(i / rowsPerCol);
      const row = i % rowsPerCol;
      const x = marginX + col * colW;
      const y = marginTop + titleH + headerH + row * rowH;
      if (entry.data) {
        slide.addImage({ data: await imageDataForPptx(entry.data, 0.2), x, y: y + 0.04, w: 0.2, h: 0.2 });
      }
      slide.addText(entry.abbreviation || String(entry.catalogId), {
        x: x + 0.28,
        y: y + 0.03,
        w: 0.72,
        h: 0.22,
        fontFace: 'Helvetica',
        fontSize: 7,
        bold: true,
        color: '1E1E1E',
        margin: 0,
        fit: 'shrink',
      });
      slide.addText(entry.officialName, {
        x: x + 1.05,
        y: y + 0.03,
        w: Math.max(0.5, colW - 1.1),
        h: 0.22,
        fontFace: 'Helvetica',
        fontSize: 6.5,
        color: '1E1E1E',
        margin: 0,
        fit: 'shrink',
        breakLine: false,
      });
    }
  }
}

function groupConnectorLegendEntries(entries: PlanConnectorLegendEntry[]): ConnectorLegendRow[] {
  const rows: ConnectorLegendRow[] = [];
  const byKey = new Map<string, ConnectorLegendRow>();

  for (const entry of entries) {
    const key = JSON.stringify({
      kind: entry.kind,
      label: entry.label,
      description: entry.description,
      line: entry.line,
    });
    const existing = byKey.get(key);
    if (existing) {
      existing.ids.push(entry.id);
      continue;
    }
    const row: ConnectorLegendRow = {
      ids: [entry.id],
      kind: entry.kind,
      label: entry.label,
      description: entry.description,
      line: entry.line,
    };
    rows.push(row);
    byKey.set(key, row);
  }

  return rows;
}

function formatConnectorLegendIDs(ids: string[]): string {
  const parsed = ids
    .map((id) => {
      const match = /^L(\d+)$/.exec(id);
      const numberText = match?.[1];
      return numberText ? { id, number: Number(numberText), width: numberText.length } : { id, number: Number.NaN, width: 0 };
    })
    .sort((a, b) => {
      if (Number.isNaN(a.number) || Number.isNaN(b.number)) return a.id.localeCompare(b.id);
      return a.number - b.number;
    });
  const ranges: string[] = [];
  const format = (number: number, width: number) => `L${String(number).padStart(width, '0')}`;

  for (let i = 0; i < parsed.length;) {
    const current = parsed[i];
    if (!current) break;
    if (Number.isNaN(current.number)) {
      ranges.push(current.id);
      i++;
      continue;
    }
    let end = i;
    while (end + 1 < parsed.length) {
      const lastInRange = parsed[end];
      const next = parsed[end + 1];
      if (!lastInRange || !next || Number.isNaN(next.number) || next.number !== lastInRange.number + 1) break;
      end++;
    }
    const last = parsed[end];
    if (!last) break;
    if (end > i) {
      ranges.push(`${format(current.number, current.width)} - ${format(last.number, last.width)}`);
    } else {
      ranges.push(format(current.number, current.width));
    }
    i = end + 1;
  }

  return ranges.join(', ');
}

function lineStyleSummary(line: PlanLine): string {
  const dash = line.dash === 'dash' ? 'dashed' : line.dash === 'dot' ? 'dotted' : 'solid';
  const begin = line.beginArrowType && line.beginArrowType !== 'none' ? line.beginArrowType : 'none';
  const end = line.endArrowType && line.endArrowType !== 'none' ? line.endArrowType : 'none';
  return `#${line.color}, ${line.width}pt, ${dash}, ${begin} -> ${end}`;
}

function lineOptions(line: PlanLine | undefined): LegendLineOptions {
  if (!line) return { color: '1E1E1E', width: 1 };

  const options: LegendLineOptions = {
    color: line.color,
    width: line.width,
  };
  options.dashType = line.dash === 'dash' ? 'dash' : line.dash === 'dot' ? 'sysDot' : 'solid';
  if (line.transparency !== undefined) options.transparency = line.transparency;
  if (line.beginArrowType !== undefined) options.beginArrowType = line.beginArrowType;
  if (line.endArrowType !== undefined) options.endArrowType = line.endArrowType;
  return options;
}
