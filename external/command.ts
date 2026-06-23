import { exportPptxFromRequest } from './usecase/pptx_exporter';
import { parsePptxExporterRequest } from './usecase/pptx_exporter_request';
import { NewEnvLogger } from './share/logger';
import { NewMCode } from './share/mcode';

const logger = NewEnvLogger('external', 'command');
const ECM001 = NewMCode('ECM-001', 'Main start');
const ECM002 = NewMCode('ECM-002', 'Main export failed');
const ECM003 = NewMCode('ECM-003', 'Main completed');
const ECIPWS001 = NewMCode('ECIPWS-001', 'Install PPTX WASI shims start');
const ECIPWS002 = NewMCode('ECIPWS-002', 'Install PPTX WASI shims set immediate branch');
const ECIPWS003 = NewMCode('ECIPWS-003', 'Install PPTX WASI shims set image branch');
const ECIPWS004 = NewMCode('ECIPWS-004', 'Install PPTX WASI shims set document branch');
const ECRAT001 = NewMCode('ECRAT-001', 'Read all text EOF branch');
const ECRAT002 = NewMCode('ECRAT-002', 'Read all text completed');
const ECWA001 = NewMCode('ECWA-001', 'Write all chunk branch');
const ECWA002 = NewMCode('ECWA-002', 'Write all completed');

declare const Javy: {
  IO: {
    readSync(fd: number, buffer: Uint8Array): number;
    writeSync(fd: number, buffer: Uint8Array): void;
  };
};

main().catch((err: unknown) => {
  const message = err instanceof Error ? err.message : String(err);
  logger.ERROR(ECM002, 'main failed', { error: message });
  writeAll(2, new TextEncoder().encode(`${message}\n`));
  throw err;
});

async function main(): Promise<void> {
  logger.DEBUG(ECM001, 'start');
  installPptxWasiShims();
  const request = parsePptxExporterRequest(readAllText(0));
  const pptx = await exportPptxFromRequest(request);
  writeAll(1, new TextEncoder().encode(bytesToBase64(pptx)));
  logger.DEBUG(ECM003, 'completed', { bytes: pptx.length });
}

function installPptxWasiShims(): void {
  logger.DEBUG(ECIPWS001, 'start');
  const globals = globalThis as Record<string, unknown>;
  if (globals.setImmediate === undefined) logger.DEBUG(ECIPWS002, 'branch set immediate');
  globals.setImmediate ??= (callback: (...args: unknown[]) => void, ...args: unknown[]) => {
    if (typeof queueMicrotask === 'function') queueMicrotask(() => callback(...args));
    else void Promise.resolve().then(() => callback(...args));
    return 0;
  };
  globals.clearImmediate ??= () => undefined;
  if (globals.Image === undefined) logger.DEBUG(ECIPWS003, 'branch set image');
  globals.Image ??= class {
    width = 1;
    height = 1;
    onload?: () => void;
    onerror?: (err: unknown) => void;
    set src(_value: string) {
      this.onload?.();
    }
    get src(): string {
      return '';
    }
  };
  if (globals.document === undefined) logger.DEBUG(ECIPWS004, 'branch set document');
  globals.document ??= {
    createElement() {
      return {
        width: 1,
        height: 1,
        getContext() {
          return { drawImage() { /* preview shim */ } };
        },
        toDataURL() {
          return 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADElEQVR4nGNgYGBgAAAABAABJzQnCgAAAABJRU5ErkJggg==';
        },
      };
    },
  };
}

function readAllText(fd: number): string {
  const chunkSize = 8192;
  const chunks: Uint8Array[] = [];
  let totalBytes = 0;
  for (;;) {
    const buffer = new Uint8Array(chunkSize);
    const bytesRead = Javy.IO.readSync(fd, buffer);
    if (bytesRead === 0) {
      logger.DEBUG(ECRAT001, 'branch EOF', { totalBytes });
      break;
    }
    chunks.push(buffer.subarray(0, bytesRead));
    totalBytes += bytesRead;
  }
  const input = new Uint8Array(totalBytes);
  let offset = 0;
  for (const chunk of chunks) {
    input.set(chunk, offset);
    offset += chunk.length;
  }
  logger.DEBUG(ECRAT002, 'completed', { totalBytes, chunks: chunks.length });
  return new TextDecoder().decode(input);
}

function writeAll(fd: number, bytes: Uint8Array): void {
  let offset = 0;
  while (offset < bytes.length) {
    const chunk = bytes.subarray(offset, Math.min(offset + 8192, bytes.length));
    logger.DEBUG(ECWA001, 'branch chunk', { fd, offset, bytes: chunk.length });
    Javy.IO.writeSync(fd, chunk);
    offset += chunk.length;
  }
  logger.DEBUG(ECWA002, 'completed', { fd, bytes: bytes.length });
}

function bytesToBase64(bytes: Uint8Array): string {
  const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/';
  let output = '';
  for (let i = 0; i < bytes.length; i += 3) {
    const a = bytes[i] ?? 0;
    const b = bytes[i + 1] ?? 0;
    const c = bytes[i + 2] ?? 0;
    output += alphabet[a >> 2];
    output += alphabet[((a & 0x03) << 4) | (b >> 4)];
    output += i + 1 < bytes.length ? alphabet[((b & 0x0f) << 2) | (c >> 6)] : '=';
    output += i + 2 < bytes.length ? alphabet[c & 0x3f] : '=';
  }
  return output;
}
