import pptxgen from 'pptxgenjs';
import JSZip from 'jszip';

/**
 * pptx.ts — PPTX drawing layer.
 *
 * This module performs NO geometry. Every layout calculation (bounds, paper
 * scaling, obstacle avoidance, arrow routing, anchoring, colour and coordinate
 * conversion) is done on the Go/WASM side, which returns a fully-resolved
 * {@link PptxPlan}. Here we only translate each plan op into the matching
 * PptxGenJS drawing call and write the file.
 */

export type PptxOutputType = 'arraybuffer' | 'base64' | 'blob' | 'nodebuffer' | 'uint8array';

/** Connector arrowhead style (passed through to the Go plan builder). */
export type ArrowStyle =
  | 'thin'
  | 'standard'
  | 'triangle'
  | 'stealth'
  | 'arrow'
  | 'diamond'
  | 'oval'
  | 'none';

/** Named paper size for the output slide (passed through to the Go plan builder). */
export type PaperSize =
  | 'A5'
  | 'A4'
  | 'A3'
  | 'A2'
  | 'A1'
  | 'Letter'
  | 'Legal'
  | 'Tabloid';

/** Slide orientation; when omitted the better-fitting orientation is chosen. */
export type PaperOrientation = 'portrait' | 'landscape';

export interface PptxExportOptions {
  // PPTX-write / metadata options handled on the TS side:
  title?: string;
  author?: string;
  company?: string;
  subject?: string;
  outputType?: PptxOutputType;
  compression?: boolean;

  // Geometry options forwarded to the Go plan builder:
  /** Shared renderer color theme. Default `'light'`. */
  theme?: 'light' | 'dark';
  /** Pixels per inch for layout scaling. Default 96. */
  pxPerInch?: number;
  /** Connector arrowhead style. Default `'thin'` (a slender stealth arrow). */
  arrowStyle?: ArrowStyle;
  /** Stub length (px) before the first/last bend. Default 20. */
  arrowStubPx?: number;
  /** Clear margin (px) reserved on both sides of every line. Default 8. */
  arrowMarginPx?: number;
  /** Named paper size for the slide; diagram is scaled and centred to fit. */
  paperSize?: PaperSize;
  /** Paper orientation; auto-fit when omitted. Ignored unless paperSize is set. */
  orientation?: PaperOrientation;
  /** Paper margin in inches applied to every side before fitting. */
  paperMargin?: number;
  /** Paper top margin in inches; overrides paperMargin for this side. */
  paperMarginTop?: number;
  /** Paper right margin in inches; overrides paperMargin for this side. */
  paperMarginRight?: number;
  /** Paper bottom margin in inches; overrides paperMargin for this side. */
  paperMarginBottom?: number;
  /** Paper left margin in inches; overrides paperMargin for this side. */
  paperMarginLeft?: number;
}

export type PptxExportResult = string | ArrayBuffer | Blob | Uint8Array;

// ---------------------------------------------------------------------------
// Plan model (mirrors internal/pptxplan output)
// ---------------------------------------------------------------------------

export interface PptxPlan {
  slide: PlanSlide;
  ops: PlanOp[];
  legend?: PlanLegendEntry[];
  connectorLegend?: PlanConnectorLegendEntry[];
}

interface PlanSlide {
  w: number;
  h: number;
  background: string;
}

interface PlanLegendEntry {
  catalogId: number;
  abbreviation: string;
  officialName: string;
  data?: string;
}

interface PlanConnectorLegendEntry {
  id: string;
  kind: string;
  label: string;
  description: string;
  source?: string;
  target?: string;
  line: PlanLine;
}

interface ConnectorLegendRow {
  ids: string[];
  kind: string;
  label: string;
  description: string;
  line: PlanLine;
}

interface PlanLine {
  color: string;
  width: number;
  dash: 'solid' | 'dash' | 'dot';
  transparency: number;
  beginArrowType?: ArrowHeadType;
  endArrowType?: ArrowHeadType;
}

interface PlanFill {
  color: string;
  transparency: number;
}

interface PlanPoint {
  x: number;
  y: number;
  moveTo?: boolean;
}

type PlanOpKind = 'rect' | 'ellipse' | 'polygon' | 'text' | 'image' | 'line';

interface PlanOp {
  id?: string;
  groupId?: string;
  frontLayer?: boolean;
  kind: PlanOpKind;
  x: number;
  y: number;
  w: number;
  h: number;
  rotate?: number;
  line?: PlanLine;
  fill?: PlanFill;
  text?: string;
  color?: string;
  fontFace?: string;
  fontSize?: number;
  bold?: boolean;
  align?: string;
  valign?: string;
  data?: string;
  transparency?: number;
  points?: PlanPoint[];
  flipH?: boolean;
  flipV?: boolean;
}

type ArrowHeadType = 'none' | 'arrow' | 'diamond' | 'oval' | 'stealth' | 'triangle';

// 'custGeom' is a valid pptxgenjs shape (custom geometry / polyline) at runtime,
// but is missing from the published typings — cast the literal so addShape accepts it.
const CUST_GEOM = 'custGeom' as Parameters<pptxgen.Slide['addShape']>[0];

// ---------------------------------------------------------------------------
// Public API
// ---------------------------------------------------------------------------

/**
 * Serialise the geometry-affecting options into the JSON object consumed by the
 * Go plan builder (`internal/pptxplan.Options`). Metadata/write options stay on
 * the TS side and are not included here.
 */
export function pptxPlanOptionsJSON(options: PptxExportOptions = {}): string {
  const o: Record<string, unknown> = {};
  if (options.theme !== undefined) o.theme = options.theme;
  if (options.pxPerInch !== undefined) o.pxPerInch = options.pxPerInch;
  if (options.arrowStyle !== undefined) o.arrowStyle = options.arrowStyle;
  if (options.arrowStubPx !== undefined) o.arrowStubPx = options.arrowStubPx;
  if (options.arrowMarginPx !== undefined) o.arrowMarginPx = options.arrowMarginPx;
  if (options.paperSize !== undefined) o.paperSize = options.paperSize;
  if (options.orientation !== undefined) o.orientation = options.orientation;
  if (options.paperMargin !== undefined) o.paperMargin = options.paperMargin;
  if (options.paperMarginTop !== undefined) o.paperMarginTop = options.paperMarginTop;
  if (options.paperMarginRight !== undefined) o.paperMarginRight = options.paperMarginRight;
  if (options.paperMarginBottom !== undefined) o.paperMarginBottom = options.paperMarginBottom;
  if (options.paperMarginLeft !== undefined) o.paperMarginLeft = options.paperMarginLeft;
  return JSON.stringify(o);
}

/**
 * Draw a fully-resolved {@link PptxPlan} (produced by the Go plan builder) into
 * a PPTX file. The only work done here is dispatching each op to PptxGenJS.
 */
export async function drawPlanToPptx(
  plan: string | PptxPlan,
  options: PptxExportOptions = {},
): Promise<PptxExportResult> {
  const parsed: PptxPlan = typeof plan === 'string' ? (JSON.parse(plan) as PptxPlan) : plan;
  const pptx = new pptxgen();
  const layoutName = 'XALIGO_EXPORT';

  pptx.defineLayout({ name: layoutName, width: parsed.slide.w, height: parsed.slide.h });
  pptx.layout = layoutName;
  pptx.author = options.author ?? 'xaligo';
  pptx.company = options.company ?? '';
  pptx.subject = options.subject ?? 'xaligo PPTX export';
  pptx.title = options.title ?? 'xaligo export';

  const slide = pptx.addSlide();
  slide.background = { color: parsed.slide.background || 'FFFFFF' };

  for (const op of parsed.ops) {
    await drawOp(slide, pptx, op);
  }
  drawConnectorLegendSlide(pptx, parsed);
  await drawLegendSlides(pptx, parsed);

  const outputType = options.outputType ?? 'uint8array';
  const bytes = await pptx.write({
    outputType: 'uint8array',
    compression: options.compression ?? true,
  }) as Uint8Array;
  const grouped = await groupAnchorObjectsInPptx(bytes, parsed.ops, options.compression ?? true);
  return convertPptxOutput(grouped, outputType);
}

function drawConnectorLegendSlide(pptx: pptxgen, plan: PptxPlan): void {
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
      slide.addShape(pptx.ShapeType.line, {
        x: lineX,
        y: rowCenterY,
        w: sampleW,
        h: 0,
        line: lineOptions(entry.line),
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
  return `#${line.color}, ${line.width}pt, ${dash}, ${begin} → ${end}`;
}

async function drawLegendSlides(pptx: pptxgen, plan: PptxPlan): Promise<void> {
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

// ---------------------------------------------------------------------------
// Op dispatch
// ---------------------------------------------------------------------------

async function drawOp(slide: pptxgen.Slide, pptx: pptxgen, op: PlanOp): Promise<void> {
  switch (op.kind) {
    case 'rect':
    case 'ellipse':
      drawShape(slide, pptx, op);
      break;
    case 'polygon':
      drawPolygon(slide, op);
      break;
    case 'text':
      drawText(slide, op);
      break;
    case 'image':
      await drawImage(slide, op);
      break;
    case 'line':
      drawLine(slide, pptx, op);
      break;
  }
}

function drawPolygon(slide: pptxgen.Slide, op: PlanOp): void {
  if (!op.points || op.points.length < 3) return;
  slide.addShape(CUST_GEOM, {
    x: op.x,
    y: op.y,
    w: op.w,
    h: op.h,
    rotate: op.rotate ?? 0,
    points: op.points.map((p) => ({
      x: p.x,
      y: p.y,
      ...(p.moveTo ? { moveTo: true } : {}),
    })),
    ...objectNameOptions(op),
    line: lineOptions(op.line),
    fill: fillOptions(op.fill),
  });
}

function drawShape(slide: pptxgen.Slide, pptx: pptxgen, op: PlanOp): void {
  const shapeType = op.kind === 'ellipse' ? pptx.ShapeType.ellipse : pptx.ShapeType.rect;
  slide.addShape(shapeType, {
    x: op.x,
    y: op.y,
    w: op.w,
    h: op.h,
    rotate: op.rotate ?? 0,
    ...objectNameOptions(op),
    line: lineOptions(op.line),
    fill: fillOptions(op.fill),
  });
}

function drawText(slide: pptxgen.Slide, op: PlanOp): void {
  const text = op.text ?? '';
  if (text === '') return;
  slide.addText(text, {
    x: op.x,
    y: op.y,
    w: op.w,
    h: op.h,
    rotate: op.rotate ?? 0,
    color: op.color ?? '1E1E1E',
    fontFace: op.fontFace ?? 'Helvetica',
    fontSize: Math.max(1, op.fontSize ?? 9),
    bold: op.bold ?? false,
    fit: 'shrink',
    wrap: !isGroupHeaderLabelOp(op),
    margin: 0,
    breakLine: false,
    align: normalizeAlign(op.align),
    valign: normalizeValign(op.valign),
    ...objectNameOptions(op),
    fill: { color: 'FFFFFF', transparency: 100 },
    line: { color: 'FFFFFF', transparency: 100 },
  });
}

function isGroupHeaderLabelOp(op: PlanOp): boolean {
  return !!op.id && op.id.endsWith('-label') && !op.id.endsWith('-item-lbl') && !/^L\d+-label$/.test(op.id);
}

async function drawImage(slide: pptxgen.Slide, op: PlanOp): Promise<void> {
  if (!op.data) return;
  slide.addImage({
    x: op.x,
    y: op.y,
    w: op.w,
    h: op.h,
    data: await imageDataForPptx(op.data, op.w),
    rotate: op.rotate ?? 0,
    transparency: op.transparency ?? 0,
    ...objectNameOptions(op),
  });
}

const pngDataCache = new Map<string, string>();

// PptxGenJS cannot create a real PNG fallback for SVG data in Node. Viewers
// that do not consume Office's svgBlip then display a broken-image marker.
// Rasterise only in Node; browsers retain the native SVG path and its vector
// quality because they have Canvas available for fallback generation.
async function imageDataForPptx(data: string, widthInches: number): Promise<string> {
  const isNode =
    typeof process !== 'undefined' &&
    process.release?.name === 'node';
  if (!isNode || !data.startsWith('data:image/svg+xml;base64,')) return data;

  const targetWidth = Math.max(32, Math.round(widthInches * 192));
  const cacheKey = `${targetWidth}:${data}`;
  const cached = pngDataCache.get(cacheKey);
  if (cached) return cached;

  const encoded = data.slice(data.indexOf(',') + 1);
  const svg = Buffer.from(encoded, 'base64');
  const { Resvg } = await import('@resvg/resvg-js');
  const png = new Resvg(svg, {
    fitTo: { mode: 'width', value: targetWidth },
    font: { loadSystemFonts: false },
  }).render().asPng();
  const result = `data:image/png;base64,${png.toString('base64')}`;
  pngDataCache.set(cacheKey, result);
  return result;
}

function drawLine(slide: pptxgen.Slide, pptx: pptxgen, op: PlanOp): void {
  if (op.points && op.points.length >= 2) {
    slide.addShape(CUST_GEOM, {
      x: op.x,
      y: op.y,
      w: op.w,
      h: op.h,
      points: op.points.map((p) => ({
        x: p.x,
        y: p.y,
        ...(p.moveTo ? { moveTo: true } : {}),
      })),
      ...objectNameOptions(op),
      line: lineOptions(op.line),
    });
    return;
  }
  slide.addShape(pptx.ShapeType.line, {
    x: op.x,
    y: op.y,
    w: op.w,
    h: op.h,
    flipH: op.flipH ?? false,
    flipV: op.flipV ?? false,
    ...objectNameOptions(op),
    line: lineOptions(op.line),
  });
}

// ---------------------------------------------------------------------------
// Style mapping
// ---------------------------------------------------------------------------

function lineOptions(line: PlanLine | undefined) {
  if (!line) return { color: '1E1E1E', width: 1 };
  const opts: {
    color: string;
    width: number;
    dashType: 'solid' | 'dash' | 'sysDot';
    transparency: number;
    beginArrowType?: ArrowHeadType;
    endArrowType?: ArrowHeadType;
  } = {
    color: line.color,
    width: line.width,
    dashType: line.dash === 'dash' ? 'dash' : line.dash === 'dot' ? 'sysDot' : 'solid',
    transparency: line.transparency,
  };
  if (line.beginArrowType) opts.beginArrowType = line.beginArrowType;
  if (line.endArrowType) opts.endArrowType = line.endArrowType;
  return opts;
}

const ANCHOR_GROUP_MARKER = 'xaligo-anchor-group|';
const FRONT_LAYER_MARKER = 'xaligo-front-layer|';

function objectNameOptions(op: PlanOp): { objectName?: string } {
  const name = objectNameForOp(op);
  return name ? { objectName: name } : {};
}

function objectNameForOp(op: PlanOp): string | undefined {
  if (!op.id) return undefined;
  if (op.frontLayer) return `${FRONT_LAYER_MARKER}${op.id}`;
  if (!op.groupId) return op.id;
  return `${ANCHOR_GROUP_MARKER}${op.groupId}|${op.id}`;
}

async function groupAnchorObjectsInPptx(bytes: Uint8Array, ops: PlanOp[], compression: boolean): Promise<Uint8Array> {
  const groupIds = [...new Set(ops.map((op) => op.groupId).filter((id): id is string => !!id))];
  if (groupIds.length === 0) return bytes;

  const zip = await JSZip.loadAsync(bytes);
  const slidePath = 'ppt/slides/slide1.xml';
  const slide = zip.file(slidePath);
  if (!slide) return bytes;

  let xml = await slide.async('string');
  xml = applySlenderStealthArrowheads(xml);
  for (const groupId of groupIds.sort()) {
    xml = groupSlideObjects(xml, groupId);
  }
  xml = moveAnchorAndLineObjectsToFront(xml);
  zip.file(slidePath, xml);
  return zip.generateAsync({ type: 'uint8array', compression: compression ? 'DEFLATE' : 'STORE' });
}

interface XmlObjectBlock {
  start: number;
  end: number;
  xml: string;
  groupId?: string;
}

interface XmlBounds {
  x: number;
  y: number;
  cx: number;
  cy: number;
}

function groupSlideObjects(xml: string, groupId: string): string {
  const blocks = collectObjectBlocks(xml);
  const groupedIndexes = blocks
    .map((block, index) => (block.groupId === groupId ? index : -1))
    .filter((index) => index >= 0);
  if (groupedIndexes.length < 2) return xml;

  const groupedBlocks = groupedIndexes.map((index) => blocks[index]).filter((block): block is XmlObjectBlock => !!block);
  const bounds = groupBounds(groupedBlocks);
  if (!bounds) return xml;

  const groupedSet = new Set(groupedIndexes);
  const insertionIndex = groupedIndexes[0];
  const groupXML = groupShapeXML(groupId, groupedBlocks.map((block) => block.xml).join(''), bounds, nextNvPrID(xml));

  let out = '';
  let cursor = 0;
  blocks.forEach((block, index) => {
    if (!groupedSet.has(index)) {
      out += xml.slice(cursor, block.end);
      cursor = block.end;
      return;
    }
    out += xml.slice(cursor, block.start);
    cursor = block.end;
    if (index === insertionIndex) out += groupXML;
  });
  out += xml.slice(cursor);
  return out;
}

function collectObjectBlocks(xml: string): XmlObjectBlock[] {
  const blocks: XmlObjectBlock[] = [];
  const re = /<p:(grpSp|sp|pic)\b[\s\S]*?<\/p:\1>/g;
  let match: RegExpExecArray | null;
  while ((match = re.exec(xml))) {
    const blockXML = match[0];
    const block: XmlObjectBlock = {
      start: match.index,
      end: match.index + blockXML.length,
      xml: blockXML,
    };
    const groupId = groupIdFromObjectBlock(blockXML);
    if (groupId) block.groupId = groupId;
    blocks.push(block);
  }
  return blocks;
}

function groupIdFromObjectBlock(xml: string): string | undefined {
  const name = /<p:cNvPr\b[^>]*\bname="([^"]*)"/.exec(xml)?.[1];
  if (!name?.startsWith(ANCHOR_GROUP_MARKER)) return undefined;
  const rest = name.slice(ANCHOR_GROUP_MARKER.length);
  const separator = rest.indexOf('|');
  return separator >= 0 ? rest.slice(0, separator) : undefined;
}

function moveAnchorAndLineObjectsToFront(xml: string): string {
  const blocks = collectObjectBlocks(xml);
  const movingIndexes = blocks
    .map((block, index) => (isAnchorGroupBlock(block.xml) || isFrontLayerBlock(block.xml) ? index : -1))
    .filter((index) => index >= 0);
  if (movingIndexes.length === 0) return xml;

  const movingSet = new Set(movingIndexes);
  const movingXML = movingIndexes.map((index) => blocks[index]?.xml ?? '').join('');

  let out = '';
  let cursor = 0;
  blocks.forEach((block, index) => {
    if (!movingSet.has(index)) {
      out += xml.slice(cursor, block.end);
    } else {
      out += xml.slice(cursor, block.start);
    }
    cursor = block.end;
  });
  out += xml.slice(cursor);
  const spTreeClose = out.lastIndexOf('</p:spTree>');
  if (spTreeClose < 0) return out;
  return `${out.slice(0, spTreeClose)}${movingXML}${out.slice(spTreeClose)}`;
}

function isAnchorGroupBlock(xml: string): boolean {
  return /<p:grpSp\b/.test(xml) && /<p:cNvPr\b[^>]*\bname="xaligo anchor xaligo-anchor-/.test(xml);
}

function isFrontLayerBlock(xml: string): boolean {
  const name = /<p:cNvPr\b[^>]*\bname="([^"]*)"/.exec(xml)?.[1];
  return !!name?.startsWith(FRONT_LAYER_MARKER);
}

function applySlenderStealthArrowheads(xml: string): string {
  return xml.replace(/<a:(headEnd|tailEnd)\b([^>]*\btype="stealth"[^>]*)\/>/g, (_match, tag: string, attrs: string) => {
    const width = /\bw="/.test(attrs) ? '' : ' w="sm"';
    const length = /\blen="/.test(attrs) ? '' : ' len="lg"';
    return `<a:${tag}${attrs}${width}${length}/>`;
  });
}

function groupBounds(blocks: XmlObjectBlock[]): XmlBounds | undefined {
  let minX = Number.POSITIVE_INFINITY;
  let minY = Number.POSITIVE_INFINITY;
  let maxX = Number.NEGATIVE_INFINITY;
  let maxY = Number.NEGATIVE_INFINITY;
  for (const block of blocks) {
    const bounds = objectBounds(block.xml);
    if (!bounds) continue;
    minX = Math.min(minX, bounds.x);
    minY = Math.min(minY, bounds.y);
    maxX = Math.max(maxX, bounds.x + bounds.cx);
    maxY = Math.max(maxY, bounds.y + bounds.cy);
  }
  if (!Number.isFinite(minX) || !Number.isFinite(minY) || !Number.isFinite(maxX) || !Number.isFinite(maxY)) return undefined;
  return { x: minX, y: minY, cx: maxX - minX, cy: maxY - minY };
}

function objectBounds(xml: string): XmlBounds | undefined {
  const off = /<a:off\b[^>]*\bx="(-?\d+)"[^>]*\by="(-?\d+)"/.exec(xml);
  const ext = /<a:ext\b[^>]*\bcx="(\d+)"[^>]*\bcy="(\d+)"/.exec(xml);
  if (!off || !ext || !off[1] || !off[2] || !ext[1] || !ext[2]) return undefined;
  return { x: Number(off[1]), y: Number(off[2]), cx: Number(ext[1]), cy: Number(ext[2]) };
}

function groupShapeXML(groupId: string, children: string, bounds: XmlBounds, id: number): string {
  const name = xmlAttr(`xaligo anchor ${groupId}`);
  return `<p:grpSp><p:nvGrpSpPr><p:cNvPr id="${id}" name="${name}"/><p:cNvGrpSpPr/><p:nvPr/></p:nvGrpSpPr><p:grpSpPr><a:xfrm><a:off x="${bounds.x}" y="${bounds.y}"/><a:ext cx="${bounds.cx}" cy="${bounds.cy}"/><a:chOff x="${bounds.x}" y="${bounds.y}"/><a:chExt cx="${bounds.cx}" cy="${bounds.cy}"/></a:xfrm></p:grpSpPr>${children}</p:grpSp>`;
}

function nextNvPrID(xml: string): number {
  let max = 1;
  const re = /<p:cNvPr\b[^>]*\bid="(\d+)"/g;
  let match: RegExpExecArray | null;
  while ((match = re.exec(xml))) {
    const id = match[1];
    if (id) max = Math.max(max, Number(id));
  }
  return max + 1;
}

function xmlAttr(value: string): string {
  return value
    .replace(/&/g, '&amp;')
    .replace(/"/g, '&quot;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
}

function convertPptxOutput(bytes: Uint8Array, outputType: PptxOutputType): PptxExportResult {
  switch (outputType) {
    case 'arraybuffer':
      return toArrayBuffer(bytes);
    case 'base64':
      return bytesToBase64(bytes);
    case 'blob':
      return new Blob([toArrayBuffer(bytes)], { type: 'application/vnd.openxmlformats-officedocument.presentationml.presentation' });
    case 'nodebuffer':
      return Buffer.from(bytes);
    case 'uint8array':
    default:
      return bytes;
  }
}

function toArrayBuffer(bytes: Uint8Array): ArrayBuffer {
  const out = new ArrayBuffer(bytes.byteLength);
  new Uint8Array(out).set(bytes);
  return out;
}

function bytesToBase64(bytes: Uint8Array): string {
  if (typeof Buffer !== 'undefined') return Buffer.from(bytes).toString('base64');
  let binary = '';
  for (const byte of bytes) binary += String.fromCharCode(byte);
  return btoa(binary);
}

function fillOptions(fill: PlanFill | undefined) {
  if (!fill) return { color: 'FFFFFF', transparency: 100 };
  return { color: fill.color, transparency: fill.transparency };
}

function normalizeAlign(align: string | undefined): 'left' | 'center' | 'right' {
  return align === 'center' || align === 'right' ? align : 'left';
}

function normalizeValign(align: string | undefined): 'top' | 'middle' | 'bottom' {
  return align === 'middle' || align === 'bottom' ? align : 'top';
}
