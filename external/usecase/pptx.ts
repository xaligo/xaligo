import type { PptxExportOptions, PptxExportResult, PptxPlan } from '../entity/pptx';
import { createPptxFromPlan } from '../repository/pptx';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

const logger = NewEnvLogger('external/usecase', 'pptx');
const EUPPOPJ001 = NewMCode('EUPPOPJ-001', 'PPTX plan options JSON theme branch');
const EUPPOPJ002 = NewMCode('EUPPOPJ-002', 'PPTX plan options JSON px per inch branch');
const EUPPOPJ003 = NewMCode('EUPPOPJ-003', 'PPTX plan options JSON arrow style branch');
const EUPPOPJ004 = NewMCode('EUPPOPJ-004', 'PPTX plan options JSON arrow stub branch');
const EUPPOPJ005 = NewMCode('EUPPOPJ-005', 'PPTX plan options JSON arrow margin branch');
const EUPPOPJ006 = NewMCode('EUPPOPJ-006', 'PPTX plan options JSON paper size branch');
const EUPPOPJ007 = NewMCode('EUPPOPJ-007', 'PPTX plan options JSON orientation branch');
const EUPPOPJ008 = NewMCode('EUPPOPJ-008', 'PPTX plan options JSON paper margin branch');
const EUPPOPJ009 = NewMCode('EUPPOPJ-009', 'PPTX plan options JSON paper margin top branch');
const EUPPOPJ010 = NewMCode('EUPPOPJ-010', 'PPTX plan options JSON paper margin right branch');
const EUPPOPJ011 = NewMCode('EUPPOPJ-011', 'PPTX plan options JSON paper margin bottom branch');
const EUPPOPJ012 = NewMCode('EUPPOPJ-012', 'PPTX plan options JSON paper margin left branch');
const EUPDPTP001 = NewMCode('EUPDPTP-001', 'Draw plan to PPTX parse string branch');
const EUPDPTP002 = NewMCode('EUPDPTP-002', 'Draw plan to PPTX use object branch');

export function pptxPlanOptionsJSON(options: PptxExportOptions = {}): string {
  const o: Record<string, unknown> = {};
  if (options.theme !== undefined) {
    logger.DEBUG(EUPPOPJ001, 'branch theme');
    o.theme = options.theme;
  }
  if (options.pxPerInch !== undefined) {
    logger.DEBUG(EUPPOPJ002, 'branch px per inch');
    o.pxPerInch = options.pxPerInch;
  }
  if (options.arrowStyle !== undefined) {
    logger.DEBUG(EUPPOPJ003, 'branch arrow style');
    o.arrowStyle = options.arrowStyle;
  }
  if (options.arrowStubPx !== undefined) {
    logger.DEBUG(EUPPOPJ004, 'branch arrow stub');
    o.arrowStubPx = options.arrowStubPx;
  }
  if (options.arrowMarginPx !== undefined) {
    logger.DEBUG(EUPPOPJ005, 'branch arrow margin');
    o.arrowMarginPx = options.arrowMarginPx;
  }
  if (options.paperSize !== undefined) {
    logger.DEBUG(EUPPOPJ006, 'branch paper size');
    o.paperSize = options.paperSize;
  }
  if (options.orientation !== undefined) {
    logger.DEBUG(EUPPOPJ007, 'branch orientation');
    o.orientation = options.orientation;
  }
  if (options.paperMargin !== undefined) {
    logger.DEBUG(EUPPOPJ008, 'branch paper margin');
    o.paperMargin = options.paperMargin;
  }
  if (options.paperMarginTop !== undefined) {
    logger.DEBUG(EUPPOPJ009, 'branch paper margin top');
    o.paperMarginTop = options.paperMarginTop;
  }
  if (options.paperMarginRight !== undefined) {
    logger.DEBUG(EUPPOPJ010, 'branch paper margin right');
    o.paperMarginRight = options.paperMarginRight;
  }
  if (options.paperMarginBottom !== undefined) {
    logger.DEBUG(EUPPOPJ011, 'branch paper margin bottom');
    o.paperMarginBottom = options.paperMarginBottom;
  }
  if (options.paperMarginLeft !== undefined) {
    logger.DEBUG(EUPPOPJ012, 'branch paper margin left');
    o.paperMarginLeft = options.paperMarginLeft;
  }
  return JSON.stringify(o);
}

export async function drawPlanToPptx(
  plan: string | PptxPlan,
  options: PptxExportOptions = {},
): Promise<PptxExportResult> {
  const parsed: PptxPlan = typeof plan === 'string' ? parsePlanJSON(plan) : plan;
  if (typeof plan !== 'string') logger.DEBUG(EUPDPTP002, 'branch plan object');
  return createPptxFromPlan(parsed, options);
}

function parsePlanJSON(plan: string): PptxPlan {
  logger.DEBUG(EUPDPTP001, 'branch plan string');
  return JSON.parse(plan) as PptxPlan;
}
