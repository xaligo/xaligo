import { readPptxExporterInput, writePptxExporterOutput } from '../repository/pptx_exporter_io';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';
import { exportPptxFromRequest } from '../usecase/pptx_exporter';
import { parsePptxExporterRequest } from '../usecase/pptx_exporter_request';

const logger = NewEnvLogger('external/controller', 'pptx_exporter');
const ECPE001 = NewMCode('ECPE-001', 'Run PPTX exporter start');
const ECPE002 = NewMCode('ECPE-002', 'Run PPTX exporter read completed');
const ECPE003 = NewMCode('ECPE-003', 'Run PPTX exporter export completed');
const ECPE004 = NewMCode('ECPE-004', 'Run PPTX exporter completed');

export async function runPptxExporter(): Promise<void> {
  logger.DEBUG(ECPE001, 'start');
  const input = await readPptxExporterInput();
  logger.DEBUG(ECPE002, 'read completed', { bytes: input.length });
  const request = parsePptxExporterRequest(input);
  const pptx = await exportPptxFromRequest(request);
  logger.DEBUG(ECPE003, 'export completed', { bytes: pptx.length });
  await writePptxExporterOutput(pptx);
  logger.DEBUG(ECPE004, 'completed');
}