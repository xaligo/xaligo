import type { PptxExportOptions, PptxExportResult } from '../entity/pptx';
import type { PptxExporterRequest } from '../entity/pptx_exporter';
import { drawPlanToPptx, pptxPlanOptionsJSON } from './pptx';
import { exportPptxFromRequest } from './pptx_exporter';
import { parsePptxExporterRequest } from './pptx_exporter_request';

export { drawPlanToPptx, pptxPlanOptionsJSON } from './pptx';
export { exportPptxFromRequest } from './pptx_exporter';
export { parsePptxExporterRequest } from './pptx_exporter_request';
export type {
  ArrowStyle,
  PaperOrientation,
  PaperSize,
  PptxExportOptions,
  PptxExportResult,
  PptxOutputType,
} from '../entity/pptx';
export type { PptxExporterOptions, PptxExporterRequest } from '../entity/pptx_exporter';

export async function renderPptxPlan(
  plan: string | PptxExporterRequest['plan'],
  options?: PptxExportOptions,
): Promise<PptxExportResult> {
  return drawPlanToPptx(plan, options);
}

export async function renderPptxExporterRequest(request: PptxExporterRequest): Promise<Uint8Array> {
  return exportPptxFromRequest(request);
}

export async function renderPptxExporterRequestJSON(input: string): Promise<Uint8Array> {
  return exportPptxFromRequest(parsePptxExporterRequest(input));
}
