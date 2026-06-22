import type { PptxExportOptions, PptxExportResult } from '../entity/pptx';
import type { PptxExporterRequest } from '../entity/pptx_exporter';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';
import { drawPlanToPptx, pptxPlanOptionsJSON } from './pptx';
import { exportPptxFromRequest } from './pptx_exporter';
import { parsePptxExporterRequest } from './pptx_exporter_request';

const logger = NewEnvLogger('external/usecase', 'api');
const EUA001 = NewMCode('EUA-001', 'Render PPTX plan start');
const EUA002 = NewMCode('EUA-002', 'Render PPTX exporter request start');
const EUA003 = NewMCode('EUA-003', 'Render PPTX exporter request JSON start');

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
export {
  GetMaxCodeLength,
  MCode,
  Mcode,
  MLOG1,
  MLOG2,
  MSYS1,
  MSYS2,
  NewMCode,
  RegisterMCodes,
} from '../share/mcode';
export { LogLevel, NewEnvLogger, NewLogger, newLogger } from '../share/logger';
export type { LogEntry, Logger, LoggerConfig } from '../share/logger';

export async function renderPptxPlan(
  plan: string | PptxExporterRequest['plan'],
  options?: PptxExportOptions,
): Promise<PptxExportResult> {
  logger.DEBUG(EUA001, 'start', { planType: typeof plan });
  return drawPlanToPptx(plan, options);
}

export async function renderPptxExporterRequest(request: PptxExporterRequest): Promise<Uint8Array> {
  logger.DEBUG(EUA002, 'start', { ops: request.plan.ops.length });
  return exportPptxFromRequest(request);
}

export async function renderPptxExporterRequestJSON(input: string): Promise<Uint8Array> {
  logger.DEBUG(EUA003, 'start', { bytes: input.length });
  return exportPptxFromRequest(parsePptxExporterRequest(input));
}
