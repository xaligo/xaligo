import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';
import { pptxPlanOpCount, pptxPlanPages } from '../entity/pptx';
import { exportPptxFromRequest } from '../usecase/pptx_exporter';
import { parsePptxExporterRequest } from '../usecase/pptx_exporter_request';

const logger = NewEnvLogger('external/controller', 'pptx_exporter');
const ECPE001 = NewMCode('ECPE-001', 'Run PPTX exporter start');
const ECPE002 = NewMCode('ECPE-002', 'Run PPTX exporter parse completed');
const ECPE003 = NewMCode('ECPE-003', 'Run PPTX exporter export completed');

// runPptxExporter is the external application entry boundary. Environment-
// specific command adapters own stdin/stdout; this controller parses their
// input and delegates PPTX creation through the external use case.
export async function runPptxExporter(input: string): Promise<Uint8Array> {
  logger.DEBUG(ECPE001, 'start', { bytes: input.length });
  const request = parsePptxExporterRequest(input);
  logger.DEBUG(ECPE002, 'parse completed', { pages: pptxPlanPages(request.plan).length, ops: pptxPlanOpCount(request.plan) });
  const pptx = await exportPptxFromRequest(request);
  logger.DEBUG(ECPE003, 'export completed', { bytes: pptx.length });
  return pptx;
}
