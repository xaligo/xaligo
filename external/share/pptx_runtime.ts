import { NewEnvLogger } from './logger';
import { NewMCode } from './mcode';

const logger = NewEnvLogger('external/share', 'pptx_runtime');
const ESPRS001 = NewMCode('ESPRS-001', 'Install PPTX runtime shims start');
const ESPRS002 = NewMCode('ESPRS-002', 'Install PPTX runtime set immediate branch');
const ESPRS003 = NewMCode('ESPRS-003', 'Install PPTX runtime set image branch');
const ESPRS004 = NewMCode('ESPRS-004', 'Install PPTX runtime set document branch');

// Modern Office renders the paired native SVG. This valid transparent PNG keeps
// the required fallback relationship well-formed where rasterization is absent.
export const PPTX_FALLBACK_PNG_BYTES = new Uint8Array([
  137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82,
  0, 0, 0, 1, 0, 0, 0, 1, 8, 4, 0, 0, 0, 181, 28, 12, 2,
  0, 0, 0, 11, 73, 68, 65, 84, 120, 218, 99, 96, 96, 0, 0, 0,
  3, 0, 1, 43, 9, 77, 132, 0, 0, 0, 0, 73, 69, 78, 68, 174,
  66, 96, 130,
]);
export const PPTX_FALLBACK_PNG_DATA =
  `data:image/png;base64,${bytesToBase64(PPTX_FALLBACK_PNG_BYTES)}`;

export function installPptxRuntimeShims(): void {
  logger.DEBUG(ESPRS001, 'start');
  const globals = globalThis as Record<string, unknown>;
  if (globals.setImmediate === undefined) logger.DEBUG(ESPRS002, 'branch set immediate');
  globals.setImmediate ??= (callback: (...args: unknown[]) => void, ...args: unknown[]) => {
    if (typeof queueMicrotask === 'function') queueMicrotask(() => callback(...args));
    else void Promise.resolve().then(() => callback(...args));
    return 0;
  };
  globals.clearImmediate ??= () => undefined;
  if (globals.Image === undefined) logger.DEBUG(ESPRS003, 'branch set image');
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
  if (globals.document === undefined) logger.DEBUG(ESPRS004, 'branch set document');
  globals.document ??= {
    createElement() {
      return {
        width: 1,
        height: 1,
        getContext() {
          return { drawImage() { /* fallback image shim */ } };
        },
        toDataURL() {
          return PPTX_FALLBACK_PNG_DATA;
        },
      };
    },
  };
}

function bytesToBase64(bytes: Uint8Array): string {
  const alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/';
  let output = '';
  for (let index = 0; index < bytes.length; index += 3) {
    const first = bytes[index] ?? 0;
    const second = bytes[index + 1] ?? 0;
    const third = bytes[index + 2] ?? 0;
    output += alphabet[first >> 2];
    output += alphabet[((first & 0x03) << 4) | (second >> 4)];
    output += index + 1 < bytes.length
      ? alphabet[((second & 0x0f) << 2) | (third >> 6)]
      : '=';
    output += index + 2 < bytes.length ? alphabet[third & 0x3f] : '=';
  }
  return output;
}
