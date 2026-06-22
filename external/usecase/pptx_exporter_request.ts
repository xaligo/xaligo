import type { PptxExporterRequest } from '../entity/pptx_exporter';

export function parsePptxExporterRequest(input: string): PptxExporterRequest {
  if (input.trim() === '') throw new Error('PPTX exporter request JSON is required on stdin');
  return validatePptxExporterRequest(JSON.parse(input));
}

function validatePptxExporterRequest(value: unknown): PptxExporterRequest {
  if (!isRecord(value)) throw new Error('PPTX exporter request must be a JSON object');
  const plan = value.plan;
  if (!isRecord(plan) || !isRecord(plan.slide) || !Array.isArray(plan.ops)) {
    throw new Error('PPTX exporter request must contain plan.slide and plan.ops');
  }
  return value as unknown as PptxExporterRequest;
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return typeof value === 'object' && value !== null;
}