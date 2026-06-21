import type { PptxExportOptions, PptxExportResult, PptxPlan } from '../entity/pptx';
import { createPptxFromPlan } from '../repository/pptx';

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

export async function drawPlanToPptx(
  plan: string | PptxPlan,
  options: PptxExportOptions = {},
): Promise<PptxExportResult> {
  const parsed: PptxPlan = typeof plan === 'string' ? (JSON.parse(plan) as PptxPlan) : plan;
  return createPptxFromPlan(parsed, options);
}
