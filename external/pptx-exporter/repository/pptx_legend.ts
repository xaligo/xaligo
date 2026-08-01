import pptxgen from 'pptxgenjs';

import type {
  ConnectorLegendRow,
  PlanConnectorLegendEntry,
  PlanLine,
  PptxPlan,
} from '../entity/pptx';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

type LegendLineOptions = pptxgen.ShapeLineProps;

const logger = NewEnvLogger('external/pptx-exporter/repository', 'pptx_legend');
const ERPLDCLS001 = NewMCode('ERPLDCLS-001', 'Draw connector legend slide empty branch');
const ERPLDCLS002 = NewMCode('ERPLDCLS-002', 'Draw connector legend slide page branch');
const ERPLDLS001 = NewMCode('ERPLDLS-001', 'Draw legend slides empty branch');
const ERPLDLS002 = NewMCode('ERPLDLS-002', 'Draw legend slides page branch');
const ERPLDLS003 = NewMCode('ERPLDLS-003', 'Draw legend slides image branch');
const ERPLGCLE001 = NewMCode('ERPLGCLE-001', 'Group connector legend entries existing branch');
const ERPLGCLE002 = NewMCode('ERPLGCLE-002', 'Group connector legend entries new branch');
const ERPLFCI001 = NewMCode('ERPLFCI-001', 'Format connector legend IDs non numeric branch');
const ERPLFCI002 = NewMCode('ERPLFCI-002', 'Format connector legend IDs range branch');
const ERPLFCI003 = NewMCode('ERPLFCI-003', 'Format connector legend IDs single branch');
const ERPLLO001 = NewMCode('ERPLLO-001', 'Legend line options default branch');

export function drawConnectorLegendSlide(pptx: pptxgen, plan: PptxPlan): number {
  const rows = groupConnectorLegendEntries(plan.connectorLegend ?? []);
  if (rows.length === 0) {
    logger.DEBUG(ERPLDCLS001, 'branch empty');
    return 0;
  }

  const slideW = plan.slide.w;
  const slideH = plan.slide.h;
  const vertical = legendVerticalLayout(slideH, 0.31);
  const columns = legendColumns(slideW, [0.13, 0.15, 0.14, 0.25, 0.33], 0.15);
  const [lineColumn, typeColumn, idColumn, styleColumn, descriptionColumn] = columns;
  if (!lineColumn || !typeColumn || !idColumn || !styleColumn || !descriptionColumn) return 0;
  const rowTextH = Math.min(0.18, vertical.rowH * 0.72);
  let slideCount = 0;

  for (let start = 0; start < rows.length; start += vertical.rowsPerPage) {
    const pageRows = rows.slice(start, start + vertical.rowsPerPage);
    logger.DEBUG(ERPLDCLS002, 'branch page', { start, rows: pageRows.length });
    const slide = pptx.addSlide();
    slideCount++;
    slide.background = { color: plan.slide.background || 'FFFFFF' };

    slide.addText('Line Legend', {
      x: lineColumn.x,
      y: vertical.marginTop,
      w: columnsWidth(columns),
      h: vertical.titleH,
      fontFace: 'Helvetica',
      fontSize: 16,
      bold: true,
      color: '1E1E1E',
      margin: 0,
      fit: 'shrink',
    });

    const headerY = vertical.marginTop + vertical.titleH;
    const headerOptions = { fontSize: 7, bold: true, color: '666666', margin: 0, fit: 'shrink' as const };
    slide.addText('Line', { x: lineColumn.x, y: headerY, w: lineColumn.w, h: vertical.headerH, ...headerOptions });
    slide.addText('Type', { x: typeColumn.x, y: headerY, w: typeColumn.w, h: vertical.headerH, ...headerOptions });
    slide.addText('ID', { x: idColumn.x, y: headerY, w: idColumn.w, h: vertical.headerH, ...headerOptions });
    slide.addText('Style', { x: styleColumn.x, y: headerY, w: styleColumn.w, h: vertical.headerH, ...headerOptions });
    slide.addText('Description', { x: descriptionColumn.x, y: headerY, w: descriptionColumn.w, h: vertical.headerH, ...headerOptions });

    for (const [i, entry] of pageRows.entries()) {
      const y = vertical.bodyTop + i * vertical.rowH;
      const rowCenterY = y + vertical.rowH / 2;
      const rowTextY = rowCenterY - rowTextH / 2;
      const line: pptxgen.ShapeLineProps = lineOptions(entry.line);
      slide.addShape(pptx.ShapeType.line, {
        x: lineColumn.x,
        y: rowCenterY,
        w: lineColumn.w,
        h: 0,
        line,
      });
      slide.addText(entry.label || entry.kind, {
        x: typeColumn.x,
        y: rowTextY,
        w: typeColumn.w,
        h: rowTextH,
        fontFace: 'Helvetica',
        fontSize: 6.5,
        bold: true,
        color: '1E1E1E',
        margin: 0,
        fit: 'shrink',
      });
      slide.addText(formatConnectorLegendIDs(entry.ids), {
        x: idColumn.x,
        y: rowTextY,
        w: idColumn.w,
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
        x: styleColumn.x,
        y: rowTextY,
        w: styleColumn.w,
        h: rowTextH,
        fontFace: 'Helvetica',
        fontSize: 6.2,
        color: '444444',
        margin: 0,
        fit: 'shrink',
        breakLine: false,
      });
      slide.addText(entry.description || '', {
        x: descriptionColumn.x,
        y: rowTextY,
        w: descriptionColumn.w,
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
  return slideCount;
}

export async function drawLegendSlides(pptx: pptxgen, plan: PptxPlan): Promise<number> {
  const entries = (plan.legend ?? []).filter((e) => e.data && e.officialName);
  if (entries.length === 0) {
    logger.DEBUG(ERPLDLS001, 'branch empty');
    return 0;
  }

  const slideW = plan.slide.w;
  const slideH = plan.slide.h;
  const vertical = legendVerticalLayout(slideH, 0.32);
  const colsPerSlide = 4;
  const columns = legendColumns(slideW, [1, 1, 1, 1], 0);
  const entriesPerSlide = vertical.rowsPerPage * colsPerSlide;
  let slideCount = 0;

  for (let start = 0; start < entries.length; start += entriesPerSlide) {
    const pageEntries = entries.slice(start, start + entriesPerSlide);
    logger.DEBUG(ERPLDLS002, 'branch page', { start, entries: pageEntries.length });
    const slide = pptx.addSlide();
    slideCount++;
    slide.background = { color: plan.slide.background || 'FFFFFF' };

    slide.addText('Legend', {
      x: columns[0]?.x ?? 0,
      y: vertical.marginTop,
      w: columnsWidth(columns),
      h: vertical.titleH,
      fontFace: 'Helvetica',
      fontSize: 16,
      bold: true,
      color: '1E1E1E',
      margin: 0,
      fit: 'shrink',
    });

    for (const column of columns) {
      const fields = serviceLegendFields(column);
      const y = vertical.marginTop + vertical.titleH;
      const headerOptions = { h: vertical.headerH, fontSize: 7, bold: true, color: '666666', margin: 0, fit: 'shrink' as const };
      slide.addText('Icon', { x: fields.icon.x, y, w: fields.icon.w, ...headerOptions });
      slide.addText('Abbr.', { x: fields.abbreviation.x, y, w: fields.abbreviation.w, ...headerOptions });
      slide.addText('Official name', { x: fields.officialName.x, y, w: fields.officialName.w, ...headerOptions });
    }

    for (const [i, entry] of pageEntries.entries()) {
      const col = Math.floor(i / vertical.rowsPerPage);
      const row = i % vertical.rowsPerPage;
      const column = columns[col];
      if (!column) continue;
      const fields = serviceLegendFields(column);
      const y = vertical.bodyTop + row * vertical.rowH;
      const textH = Math.min(0.22, vertical.rowH * 0.7);
      const textY = y + (vertical.rowH - textH) / 2;
      if (entry.data) {
        logger.DEBUG(ERPLDLS003, 'branch image', { catalogId: entry.catalogId });
        const imageSize = Math.min(0.2, fields.icon.w, vertical.rowH * 0.7);
        slide.addImage({
          data: entry.data,
          x: fields.icon.x + (fields.icon.w - imageSize) / 2,
          y: y + (vertical.rowH - imageSize) / 2,
          w: imageSize,
          h: imageSize,
        });
      }
      slide.addText(entry.abbreviation || String(entry.catalogId), {
        x: fields.abbreviation.x,
        y: textY,
        w: fields.abbreviation.w,
        h: textH,
        fontFace: 'Helvetica',
        fontSize: 7,
        bold: true,
        color: '1E1E1E',
        margin: 0,
        fit: 'shrink',
      });
      slide.addText(entry.officialName, {
        x: fields.officialName.x,
        y: textY,
        w: fields.officialName.w,
        h: textH,
        fontFace: 'Helvetica',
        fontSize: 6.5,
        color: '1E1E1E',
        margin: 0,
        fit: 'shrink',
        breakLine: false,
      });
    }
  }
  return slideCount;
}

interface LegendColumn {
  x: number;
  w: number;
}

interface LegendVerticalLayout {
  marginTop: number;
  titleH: number;
  headerH: number;
  bodyTop: number;
  rowH: number;
  rowsPerPage: number;
}

function legendColumns(slideW: number, weights: number[], maximumGap: number): LegendColumn[] {
  const marginX = Math.min(0.55, slideW * 0.08);
  const contentW = Math.max(Number.EPSILON, slideW - marginX * 2);
  const gap = Math.min(maximumGap, contentW * 0.015);
  const availableW = Math.max(Number.EPSILON, contentW - gap * Math.max(0, weights.length - 1));
  const weightTotal = weights.reduce((sum, weight) => sum + weight, 0);
  const columns: LegendColumn[] = [];
  let x = marginX;
  for (const [index, weight] of weights.entries()) {
    const isLast = index === weights.length - 1;
    const w = isLast ? Math.max(Number.EPSILON, slideW - marginX - x) : availableW * weight / weightTotal;
    columns.push({ x, w });
    x += w + gap;
  }
  return columns;
}

function columnsWidth(columns: LegendColumn[]): number {
  const first = columns[0];
  const last = columns[columns.length - 1];
  if (!first || !last) return 0;
  return last.x + last.w - first.x;
}

function legendVerticalLayout(slideH: number, maximumRowH: number): LegendVerticalLayout {
  const marginTop = Math.min(0.42, slideH * 0.08);
  const marginBottom = Math.min(0.35, slideH * 0.07);
  const titleH = Math.min(0.35, slideH * 0.12);
  const headerH = Math.min(0.24, slideH * 0.08);
  const bodyGap = Math.min(0.06, slideH * 0.02);
  const usableH = Math.max(Number.EPSILON, slideH - marginTop - marginBottom - titleH - headerH - bodyGap);
  const rowH = Math.min(maximumRowH, usableH);
  return {
    marginTop,
    titleH,
    headerH,
    bodyTop: marginTop + titleH + headerH + bodyGap,
    rowH,
    rowsPerPage: Math.max(1, Math.floor(usableH / rowH)),
  };
}

function serviceLegendFields(column: LegendColumn): {
  icon: LegendColumn;
  abbreviation: LegendColumn;
  officialName: LegendColumn;
} {
  const gap = Math.min(0.08, column.w * 0.03);
  const iconW = Math.min(0.38, column.w * 0.22);
  const abbreviationW = Math.min(0.72, column.w * 0.28);
  const abbreviationX = column.x + iconW + gap;
  const officialNameX = abbreviationX + abbreviationW + gap;
  return {
    icon: { x: column.x, w: iconW },
    abbreviation: { x: abbreviationX, w: abbreviationW },
    officialName: { x: officialNameX, w: Math.max(Number.EPSILON, column.x + column.w - officialNameX) },
  };
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
      logger.DEBUG(ERPLGCLE001, 'branch existing', { id: entry.id });
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
    logger.DEBUG(ERPLGCLE002, 'branch new', { id: entry.id, kind: entry.kind });
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
      logger.DEBUG(ERPLFCI001, 'branch non numeric', { id: current.id });
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
      logger.DEBUG(ERPLFCI002, 'branch range', { start: current.id, end: last.id });
      ranges.push(`${format(current.number, current.width)} - ${format(last.number, last.width)}`);
    } else {
      logger.DEBUG(ERPLFCI003, 'branch single', { id: current.id });
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
  if (!line) {
    logger.DEBUG(ERPLLO001, 'branch default line');
    return { color: '1E1E1E', width: 1 };
  }

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
