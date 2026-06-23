import pptxgen from 'pptxgenjs';

import type {
  ArrowHeadType,
  PlanFill,
  PlanLine,
  PlanOp,
  PptxExportOptions,
  PptxExportResult,
  PptxPlan,
} from '../entity/pptx';
import { ANCHOR_GROUP_MARKER, FRONT_LAYER_MARKER } from '../entity/pptx';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';
import { imageDataForPptx } from './pptx_image';
import { drawConnectorLegendSlide, drawLegendSlides } from './pptx_legend';
import { convertPptxOutput, groupAnchorObjectsInPptx } from './pptx_package';

const logger = NewEnvLogger('external/repository', 'pptx');
const ERPCPFP001 = NewMCode('ERPCPFP-001', 'Create PPTX from plan start');
const ERPCPFP002 = NewMCode('ERPCPFP-002', 'Create PPTX from plan write completed');
const ERPCPFP003 = NewMCode('ERPCPFP-003', 'Create PPTX from plan grouping completed');
const ERPCPFP004 = NewMCode('ERPCPFP-004', 'Create PPTX from plan completed');
const ERPDO001 = NewMCode('ERPDO-001', 'Draw op dispatch branch');
const ERPDP001 = NewMCode('ERPDP-001', 'Draw polygon skipped branch');
const ERPDT001 = NewMCode('ERPDT-001', 'Draw text skipped branch');
const ERPDI001 = NewMCode('ERPDI-001', 'Draw image skipped branch');
const ERPDL001 = NewMCode('ERPDL-001', 'Draw line custom geometry branch');
const ERPDL002 = NewMCode('ERPDL-002', 'Draw line fallback branch');
const ERPFLO001 = NewMCode('ERPFLO-001', 'Line options default branch');
const ERPFOP001 = NewMCode('ERPFOP-001', 'Fill options default branch');

// 'custGeom' is a valid pptxgenjs shape (custom geometry / polyline) at runtime,
// but is missing from the published typings — cast the literal so addShape accepts it.
const CUST_GEOM = 'custGeom' as Parameters<pptxgen.Slide['addShape']>[0];

export async function createPptxFromPlan(
  parsed: PptxPlan,
  options: PptxExportOptions = {},
): Promise<PptxExportResult> {
  logger.DEBUG(ERPCPFP001, 'start', { ops: parsed.ops.length, outputType: options.outputType ?? 'uint8array' });
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
  logger.DEBUG(ERPCPFP002, 'write completed', { bytes: bytes.length });
  const grouped = await groupAnchorObjectsInPptx(bytes, parsed.ops, options.compression ?? true);
  logger.DEBUG(ERPCPFP003, 'grouping completed', { bytes: grouped.length });
  const result = convertPptxOutput(grouped, outputType);
  logger.DEBUG(ERPCPFP004, 'completed', { outputType });
  return result;
}

// ---------------------------------------------------------------------------
// Op dispatch
// ---------------------------------------------------------------------------

async function drawOp(slide: pptxgen.Slide, pptx: pptxgen, op: PlanOp): Promise<void> {
  logger.DEBUG(ERPDO001, 'branch op', { kind: op.kind, id: op.id });
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
  if (!op.points || op.points.length < 3) {
    logger.WARN(ERPDP001, 'branch skipped polygon', { id: op.id, points: op.points?.length ?? 0 });
    return;
  }
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
  if (text === '') {
    logger.DEBUG(ERPDT001, 'branch skipped empty text', { id: op.id });
    return;
  }
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
  if (!op.data) {
    logger.WARN(ERPDI001, 'branch skipped image without data', { id: op.id });
    return;
  }
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

function drawLine(slide: pptxgen.Slide, pptx: pptxgen, op: PlanOp): void {
  if (op.points && op.points.length >= 2) {
    logger.DEBUG(ERPDL001, 'branch custom geometry', { id: op.id, points: op.points.length });
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
  logger.DEBUG(ERPDL002, 'branch fallback line', { id: op.id });
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
  if (!line) {
    logger.DEBUG(ERPFLO001, 'branch default line');
    return { color: '1E1E1E', width: 1 };
  }
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

function fillOptions(fill: PlanFill | undefined) {
  if (!fill) {
    logger.DEBUG(ERPFOP001, 'branch default fill');
    return { color: 'FFFFFF', transparency: 100 };
  }
  return { color: fill.color, transparency: fill.transparency };
}

function normalizeAlign(align: string | undefined): 'left' | 'center' | 'right' {
  return align === 'center' || align === 'right' ? align : 'left';
}

function normalizeValign(align: string | undefined): 'top' | 'middle' | 'bottom' {
  return align === 'middle' || align === 'bottom' ? align : 'top';
}
