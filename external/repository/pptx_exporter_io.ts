import { Buffer } from 'node:buffer';
import process from 'node:process';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

const logger = NewEnvLogger('external/repository', 'pptx_exporter_io');
const ERPEIRI001 = NewMCode('ERPEIRI-001', 'Read PPTX exporter input start');
const ERPEIRI002 = NewMCode('ERPEIRI-002', 'Read PPTX exporter input chunk branch');
const ERPEIRI003 = NewMCode('ERPEIRI-003', 'Read PPTX exporter input completed');
const ERPEIWO001 = NewMCode('ERPEIWO-001', 'Write PPTX exporter output start');
const ERPEIWO002 = NewMCode('ERPEIWO-002', 'Write PPTX exporter output failed');
const ERPEIWO003 = NewMCode('ERPEIWO-003', 'Write PPTX exporter output completed');

export async function readPptxExporterInput(): Promise<string> {
  logger.DEBUG(ERPEIRI001, 'start');
  const chunks: Buffer[] = [];
  for await (const chunk of process.stdin) {
    const buffer = Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk);
    logger.DEBUG(ERPEIRI002, 'branch chunk', { bytes: buffer.length });
    chunks.push(buffer);
  }
  const input = Buffer.concat(chunks).toString('utf8');
  logger.DEBUG(ERPEIRI003, 'completed', { bytes: input.length, chunks: chunks.length });
  return input;
}

export function writePptxExporterOutput(bytes: Uint8Array): Promise<void> {
  logger.DEBUG(ERPEIWO001, 'start', { bytes: bytes.length });
  return new Promise((resolve, reject) => {
    process.stdout.write(Buffer.from(bytes), (err) => {
      if (err) {
        logger.ERROR(ERPEIWO002, 'write failed', { error: err });
        reject(err);
      } else {
        logger.DEBUG(ERPEIWO003, 'completed', { bytes: bytes.length });
        resolve();
      }
    });
  });
}