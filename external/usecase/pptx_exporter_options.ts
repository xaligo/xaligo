import type { PptxExportOptions } from '../entity/pptx';
import type { PptxExporterOptions } from '../entity/pptx_exporter';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

const logger = NewEnvLogger('external/usecase', 'pptx_exporter_options');
const EUPEOPEO001 = NewMCode('EUPEOPEO-001', 'PPTX exporter options title branch');
const EUPEOPEO002 = NewMCode('EUPEOPEO-002', 'PPTX exporter options author branch');
const EUPEOPEO003 = NewMCode('EUPEOPEO-003', 'PPTX exporter options company branch');
const EUPEOPEO004 = NewMCode('EUPEOPEO-004', 'PPTX exporter options subject branch');
const EUPEOPEO005 = NewMCode('EUPEOPEO-005', 'PPTX exporter options compression branch');

export function pptxExporterOptions(options: PptxExporterOptions | undefined): PptxExportOptions {
  const result: PptxExportOptions = {
    outputType: 'uint8array',
  };
  if (options?.title !== undefined) {
    logger.DEBUG(EUPEOPEO001, 'branch title');
    result.title = options.title;
  }
  if (options?.author !== undefined) {
    logger.DEBUG(EUPEOPEO002, 'branch author');
    result.author = options.author;
  }
  if (options?.company !== undefined) {
    logger.DEBUG(EUPEOPEO003, 'branch company');
    result.company = options.company;
  }
  if (options?.subject !== undefined) {
    logger.DEBUG(EUPEOPEO004, 'branch subject');
    result.subject = options.subject;
  }
  if (options?.compression !== undefined) {
    logger.DEBUG(EUPEOPEO005, 'branch compression', { compression: options.compression });
    result.compression = options.compression;
  }
  return result;
}