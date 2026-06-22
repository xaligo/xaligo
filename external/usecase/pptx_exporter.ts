import type { PptxExporterRequest } from '../entity/pptx_exporter';
import { drawPlanToPptx } from './pptx';
import { pptxExporterOptions } from './pptx_exporter_options';

export async function exportPptxFromRequest(request: PptxExporterRequest): Promise<Uint8Array> {
  const out = await drawPlanToPptx(request.plan, pptxExporterOptions(request.options));
  return out as Uint8Array;
}