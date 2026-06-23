import type { PptxExporterRequest } from '../entity/pptx_exporter';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';
import { drawPlanToPptx } from './pptx';
import { pptxExporterOptions } from './pptx_exporter_options';

const logger = NewEnvLogger('external/usecase', 'pptx_exporter');
const EUPE001 = NewMCode('EUPE-001', 'Export PPTX from request start');
const EUPE002 = NewMCode('EUPE-002', 'Export PPTX from request completed');

export async function exportPptxFromRequest(request: PptxExporterRequest): Promise<Uint8Array> {
  logger.DEBUG(EUPE001, 'start', { ops: request.plan.ops.length });
  const out = await drawPlanToPptx(request.plan, pptxExporterOptions(request.options));
  const bytes = out as Uint8Array;
  logger.DEBUG(EUPE002, 'completed', { bytes: bytes.length });
  return bytes;
}