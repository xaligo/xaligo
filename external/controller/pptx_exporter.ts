import { readPptxExporterInput, writePptxExporterOutput } from '../repository/pptx_exporter_io';
import { exportPptxFromRequest } from '../usecase/pptx_exporter';
import { parsePptxExporterRequest } from '../usecase/pptx_exporter_request';

export async function runPptxExporter(): Promise<void> {
  const request = parsePptxExporterRequest(await readPptxExporterInput());
  const pptx = await exportPptxFromRequest(request);
  await writePptxExporterOutput(pptx);
}