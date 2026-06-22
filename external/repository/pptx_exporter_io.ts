import { Buffer } from 'node:buffer';
import process from 'node:process';

export async function readPptxExporterInput(): Promise<string> {
  const chunks: Buffer[] = [];
  for await (const chunk of process.stdin) {
    chunks.push(Buffer.isBuffer(chunk) ? chunk : Buffer.from(chunk));
  }
  return Buffer.concat(chunks).toString('utf8');
}

export function writePptxExporterOutput(bytes: Uint8Array): Promise<void> {
  return new Promise((resolve, reject) => {
    process.stdout.write(Buffer.from(bytes), (err) => {
      if (err) reject(err);
      else resolve();
    });
  });
}