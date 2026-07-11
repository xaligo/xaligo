import pptxgen from 'pptxgenjs';

import type {
  ArrowHeadType,
  PlanFill,
  PlanLine,
  PlanOp,
  PlanTextLayout,
  PptxExportOptions,
  PptxExportResult,
  PptxPlan,
} from '../entity/pptx';
import { planObjectName, planTextOverflow } from '../entity/pptx';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';
import { imageDataForPptx } from './pptx_image';
import { drawConnectorLegendSlide, drawLegendSlides } from './pptx_legend';
import { convertPptxOutput, finalizePptxPackage } from './pptx_package';

const logger = NewEnvLogger('external/repository', 'pptx');
const ERPCPFP001 = NewMCode('ERPCPFP-001', 'Create PPTX from plan start');
const ERPCPFP002 = NewMCode('ERPCPFP-002', 'Create PPTX from plan write completed');
const ERPCPFP003 = NewMCode('ERPCPFP-003', 'Create PPTX from plan package finalization completed');
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
  const finalized = await finalizePptxPackage(bytes, parsed.ops, options.compression ?? true);
  logger.DEBUG(ERPCPFP003, 'package finalization completed', { bytes: finalized.length });
  const result = convertPptxOutput(finalized, outputType);
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
  const layout = resolvedTextLayout(op);
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
    fit: pptxTextFit(layout, !!planObjectName(op)),
    wrap: layout.wrap,
    margin: textMargin(layout),
    lineSpacingMultiple: positiveLineHeight(layout.lineHeight),
    breakLine: false,
    align: normalizeAlign(op.align),
    valign: normalizeValign(op.valign),
    ...objectNameOptions(op),
    fill: { color: 'FFFFFF', transparency: 100 },
    line: { color: 'FFFFFF', transparency: 100 },
  });
}

function resolvedTextLayout(op: PlanOp): PlanTextLayout {
  if (op.textLayout) {
    const overflow = planTextOverflow(op);
    return { ...op.textLayout, overflow, clip: overflow === 'clip' };
  }
  // Compatibility for plan JSON generated before the renderer-neutral text
  // contract. New plans never need to infer semantics from object names.
  const groupHeader = isLegacyGroupHeaderLabelOp(op);
  return {
    role: groupHeader ? 'group-header' : 'label',
    wrap: !groupHeader,
    fit: 'shrink',
    overflow: 'clip',
    clip: true,
    lineHeight: 1.2,
    padding: { top: 0, right: 0, bottom: 0, left: 0 },
  };
}

function pptxTextFit(layout: PlanTextLayout, hasObjectName: boolean): 'none' | 'shrink' {
  if (layout.fit === 'shrink') return 'shrink';
  // Named text objects receive exact DrawingML overflow attributes during
  // package finalization. An unnamed clipped object cannot be found reliably,
  // so keep shrink as a compatibility containment fallback for that edge case.
  return (layout.overflow === 'clip' || layout.clip) && !hasObjectName ? 'shrink' : 'none';
}

function isLegacyGroupHeaderLabelOp(op: PlanOp): boolean {
  return !!op.id && op.id.endsWith('-label') && !op.id.endsWith('-item-lbl') && !/^L\d+-label$/.test(op.id);
}

function textMargin(layout: PlanTextLayout): number | [number, number, number, number] {
  const padding = layout.padding ?? { top: 0, right: 0, bottom: 0, left: 0 };
  const values: [number, number, number, number] = [
    Math.max(0, padding.top * 72),
    Math.max(0, padding.right * 72),
    Math.max(0, padding.bottom * 72),
    Math.max(0, padding.left * 72),
  ];
  return values.every((value) => value === values[0]) ? values[0] : values;
}

function positiveLineHeight(value: number | undefined): number {
  return value && Number.isFinite(value) && value > 0 ? value : 1.2;
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
    const geom = lineGeometryForPptx(op);
    slide.addShape(CUST_GEOM, {
      x: geom.x,
      y: geom.y,
      w: geom.w,
      h: geom.h,
      points: geom.points.map((p) => ({
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
  const geom = fallbackLineGeometryForPptx(op);
  slide.addShape(pptx.ShapeType.line, {
    x: geom.x,
    y: geom.y,
    w: geom.w,
    h: geom.h,
    flipH: geom.flipH,
    flipV: geom.flipV,
    ...objectNameOptions(op),
    line: lineOptions(op.line),
  });
}

function lineGeometryForPptx(op: PlanOp): { x: number; y: number; w: number; h: number; points: NonNullable<PlanOp['points']> } {
  const abs = (op.points ?? []).map((p) => ({ x: op.x + p.x, y: op.y + p.y, moveTo: p.moveTo }));
  extendEndpoint(abs, 0, 1, op.line?.beginArrowExtendIn);
  extendEndpoint(abs, abs.length - 1, abs.length - 2, op.line?.endArrowExtendIn);
  const bounds = pointBounds(abs);
  return {
    x: bounds.x,
    y: bounds.y,
    w: bounds.w,
    h: bounds.h,
    points: abs.map((p) => ({ x: p.x - bounds.x, y: p.y - bounds.y, ...(p.moveTo ? { moveTo: true } : {}) })),
  };
}

function fallbackLineGeometryForPptx(op: PlanOp): { x: number; y: number; w: number; h: number; flipH: boolean; flipV: boolean } {
  let start = { x: op.x, y: op.y };
  let end = { x: op.x + op.w, y: op.y + op.h };
  if (op.flipH) [start.x, end.x] = [end.x, start.x];
  if (op.flipV) [start.y, end.y] = [end.y, start.y];
  const points = [start, end];
  extendEndpoint(points, 0, 1, op.line?.beginArrowExtendIn);
  extendEndpoint(points, 1, 0, op.line?.endArrowExtendIn);
  const x = Math.min(points[0].x, points[1].x);
  const y = Math.min(points[0].y, points[1].y);
  return {
    x,
    y,
    w: Math.abs(points[1].x - points[0].x),
    h: Math.abs(points[1].y - points[0].y),
    flipH: points[1].x < points[0].x,
    flipV: points[1].y < points[0].y,
  };
}

function extendEndpoint(points: Array<{ x: number; y: number }>, endpointIndex: number, neighborIndex: number, extension: number | undefined): void {
  if (!extension || extension <= 0) return;
  const endpoint = points[endpointIndex];
  const neighbor = points[neighborIndex];
  if (!endpoint || !neighbor) return;
  const dx = endpoint.x - neighbor.x;
  const dy = endpoint.y - neighbor.y;
  const len = Math.hypot(dx, dy);
  if (len <= 0) return;
  endpoint.x += (dx / len) * extension;
  endpoint.y += (dy / len) * extension;
}

function pointBounds(points: Array<{ x: number; y: number }>): { x: number; y: number; w: number; h: number } {
  let minX = points[0]?.x ?? 0;
  let minY = points[0]?.y ?? 0;
  let maxX = minX;
  let maxY = minY;
  for (const p of points) {
    minX = Math.min(minX, p.x);
    minY = Math.min(minY, p.y);
    maxX = Math.max(maxX, p.x);
    maxY = Math.max(maxY, p.y);
  }
  return { x: minX, y: minY, w: Math.max(0.0001, maxX - minX), h: Math.max(0.0001, maxY - minY) };
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
  const name = planObjectName(op);
  return name ? { objectName: name } : {};
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
