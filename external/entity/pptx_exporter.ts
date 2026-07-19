import type { PptxExportOptions, PptxPlanInput } from './pptx';

export interface PptxExporterRequest {
  plan: PptxPlanInput;
  options?: PptxExporterOptions;
}

export type PptxExporterOptions = Pick<
  PptxExportOptions,
  'title' | 'author' | 'company' | 'subject' | 'compression'
>;
