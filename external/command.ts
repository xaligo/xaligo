import { exportPptxFromRequest } from './usecase/pptx_exporter';
import { parsePptxExporterRequest } from './usecase/pptx_exporter_request';

declare const Javy: {
  IO: {
    readSync(fd: number, buffer: Uint8Array): number;
    writeSync(fd: number, buffer: Uint8Array): void;
  };
};

main().catch((err: unknown) => {
  const message = err instanceof Error ? err.message : String(err);
  writeAll(2, new TextEncoder().encode(`${message}\n`));
  throw err;
});

async function main(): Promise<void> {
  installPptxWasiShims();
  const request = parsePptxExporterRequest(readAllText(0));
  const pptx = await exportPptxFromRequest(request);
  writeAll(1, pptx);
}

function installPptxWasiShims(): void {
  const globals = globalThis as Record<string, unknown>;
  globals.setImmediate ??= (callback: (...args: unknown[]) => void, ...args: unknown[]) => {
    if (typeof queueMicrotask === 'function') queueMicrotask(() => callback(...args));
    else void Promise.resolve().then(() => callback(...args));
    return 0;
  };
  globals.clearImmediate ??= () => undefined;
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
    if (bytesRead === 0) break;
    chunks.push(buffer.subarray(0, bytesRead));
    totalBytes += bytesRead;
  }
  const input = new Uint8Array(totalBytes);
  let offset = 0;
  for (const chunk of chunks) {
    input.set(chunk, offset);
    offset += chunk.length;
  }
  return new TextDecoder().decode(input);
}

function writeAll(fd: number, bytes: Uint8Array): void {
  let offset = 0;
  while (offset < bytes.length) {
    const chunk = bytes.subarray(offset, Math.min(offset + 8192, bytes.length));
    Javy.IO.writeSync(fd, chunk);
    offset += chunk.length;
  }
}
