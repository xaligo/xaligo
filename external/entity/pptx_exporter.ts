import type { PptxExportOptions, PptxPlan } from './pptx';

export interface PptxExporterRequest {
  plan: PptxPlan;
  options?: PptxExporterOptions;
}

export type PptxExporterOptions = Pick<
  PptxExportOptions,
  'title' | 'author' | 'company' | 'subject' | 'compression'
>;