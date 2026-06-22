import type { PptxExporterRequest } from '../entity/pptx_exporter';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

const logger = NewEnvLogger('external/usecase', 'pptx_exporter_request');
const EUPERPPER001 = NewMCode('EUPERPPER-001', 'Parse PPTX exporter request empty input branch');
const EUPERVPER001 = NewMCode('EUPERVPER-001', 'Validate PPTX exporter request invalid object branch');
const EUPERVPER002 = NewMCode('EUPERVPER-002', 'Validate PPTX exporter request invalid plan branch');

export function parsePptxExporterRequest(input: string): PptxExporterRequest {
  if (input.trim() === '') {
    logger.ERROR(EUPERPPER001, 'branch empty input');
    throw new Error('PPTX exporter request JSON is required on stdin');
  }
  return validatePptxExporterRequest(JSON.parse(input));
}

function validatePptxExporterRequest(value: unknown): PptxExporterRequest {
  if (!isRecord(value)) {
    logger.ERROR(EUPERVPER001, 'branch invalid object');
    throw new Error('PPTX exporter request must be a JSON object');
  }
  const plan = value.plan;
  if (!isRecord(plan) || !isRecord(plan.slide) || !Array.isArray(plan.ops)) {
    logger.ERROR(EUPERVPER002, 'branch invalid plan');
    throw new Error('PPTX exporter request must contain plan.slide and plan.ops');
  }
  return value as unknown as PptxExporterRequest;
}

function isRecord(value: unknown): value is Record<string, unknown> {
  return typeof value === 'object' && value !== null;
}