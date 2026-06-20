import pptxgen from 'pptxgenjs';

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
}

export type PptxExportResult = string | ArrayBuffer | Blob | Uint8Array;

// ---------------------------------------------------------------------------
// Plan model (mirrors internal/pptxplan output)
// ---------------------------------------------------------------------------

export interface PptxPlan {
  slide: PlanSlide;
  ops: PlanOp[];
  legend?: PlanLegendEntry[];
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

type PlanOpKind = 'rect' | 'ellipse' | 'text' | 'image' | 'line';

interface PlanOp {
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
  await drawLegendSlides(pptx, parsed);

  return pptx.write({
    outputType: options.outputType ?? 'uint8array',
    compression: options.compression ?? true,
  }) as Promise<PptxExportResult>;
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

function drawShape(slide: pptxgen.Slide, pptx: pptxgen, op: PlanOp): void {
  const shapeType = op.kind === 'ellipse' ? pptx.ShapeType.ellipse : pptx.ShapeType.rect;
  slide.addShape(shapeType, {
    x: op.x,
    y: op.y,
    w: op.w,
    h: op.h,
    rotate: op.rotate ?? 0,
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
    margin: 0,
    breakLine: false,
    align: normalizeAlign(op.align),
    valign: normalizeValign(op.valign),
    fill: { color: 'FFFFFF', transparency: 100 },
    line: { color: 'FFFFFF', transparency: 100 },
  });
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
