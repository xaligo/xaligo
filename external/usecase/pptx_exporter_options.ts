import type { PptxExportOptions } from '../entity/pptx';
import type { PptxExporterOptions } from '../entity/pptx_exporter';

export function pptxExporterOptions(options: PptxExporterOptions | undefined): PptxExportOptions {
  const result: PptxExportOptions = {
    outputType: 'uint8array',
  };
  if (options?.title !== undefined) result.title = options.title;
  if (options?.author !== undefined) result.author = options.author;
  if (options?.company !== undefined) result.company = options.company;
  if (options?.subject !== undefined) result.subject = options.subject;
  if (options?.compression !== undefined) result.compression = options.compression;
  return result;
}