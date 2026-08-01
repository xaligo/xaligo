import { pptxPlanOpCount, pptxPlanPages, type PptxExportOptions } from '../entity/pptx';
import type { PptxExporterRequest } from '../entity/pptx_exporter';
import type { PptxExporterOptions } from '../entity/pptx_exporter';
import { createPptxFromPlan } from '../repository/pptx';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

const logger = NewEnvLogger('external/pptx-exporter/usecase', 'pptx_exporter');
const EUPE001 = NewMCode('EUPE-001', 'Export PPTX from request start');
const EUPE002 = NewMCode('EUPE-002', 'Export PPTX from request completed');

export async function exportPptxFromRequest(request: PptxExporterRequest): Promise<Uint8Array> {
  logger.DEBUG(EUPE001, 'start', { pages: pptxPlanPages(request.plan).length, ops: pptxPlanOpCount(request.plan) });
  const out = await createPptxFromPlan(request.plan, pptxExporterOptions(request.options));
  const bytes = out as Uint8Array;
  logger.DEBUG(EUPE002, 'completed', { bytes: bytes.length });
  return bytes;
}

function pptxExporterOptions(options: PptxExporterOptions | undefined): PptxExportOptions {
  const result: PptxExportOptions = { outputType: 'uint8array' };
  if (options?.title !== undefined) result.title = options.title;
  if (options?.author !== undefined) result.author = options.author;
  if (options?.company !== undefined) result.company = options.company;
  if (options?.subject !== undefined) result.subject = options.subject;
  if (options?.compression !== undefined) result.compression = options.compression;
  return result;
}
