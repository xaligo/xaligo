import type { WasmResult } from '../entity/api';

export interface WasmGateway {
  diagnose(xal: string): WasmResult;
  renderXYFlow(xal: string): WasmResult;
  renderIsoflow(xal: string): WasmResult;
  render(xal: string): WasmResult;
  renderWithServices(xal: string, servicesCsv: string): WasmResult;
  buildPptxPlan(xal: string, servicesCsv: string, optionsJson: string): WasmResult;
}

declare global {
  var Go: new () => {
    importObject: WebAssembly.Imports;
    run(instance: WebAssembly.Instance): Promise<void>;
  };
  function xaligoRender(xal: string): WasmResult;
  function xaligoRenderWithServices(xal: string, servicesCsv: string): WasmResult;
  function xaligoBuildPptxPlan(xal: string, servicesCsv: string, optionsJson: string): WasmResult;
  function xaligoDiagnose(xal: string): WasmResult;
  function xaligoRenderXYFlow(xal: string): WasmResult;
  function xaligoRenderIsoflow(xal: string): WasmResult;
}

let initialized = false;

export async function createWasmGateway(wasmUrl?: string): Promise<WasmGateway> {
  await initializeWasm(wasmUrl);
  return {
    diagnose: (xal: string) => globalThis.xaligoDiagnose(xal),
    renderXYFlow: (xal: string) => globalThis.xaligoRenderXYFlow(xal),
    renderIsoflow: (xal: string) => globalThis.xaligoRenderIsoflow(xal),
    render: (xal: string) => globalThis.xaligoRender(xal),
    renderWithServices: (xal: string, servicesCsv: string) => globalThis.xaligoRenderWithServices(xal, servicesCsv),
    buildPptxPlan: (xal: string, servicesCsv: string, optionsJson: string) => globalThis.xaligoBuildPptxPlan(xal, servicesCsv, optionsJson),
  };
}

async function initializeWasm(wasmUrl?: string): Promise<void> {
  if (initialized) return;

  if (typeof globalThis.Go === 'undefined') {
    const { createRequire } = await import('node:module');
    const require = createRequire(import.meta.url);
    require('../wasm/wasm_exec.js');
  }

  const resolvedUrl = wasmUrl ?? new URL('../wasm/xaligo.wasm', import.meta.url).toString();
  const go = new globalThis.Go();
  const wasmBytes = await readWasmBytes(resolvedUrl);
  const result = await WebAssembly.instantiate(wasmBytes, go.importObject);
  void go.run(result.instance);
  await waitFor(() => typeof globalThis.xaligoRender === 'function', 5000);
  initialized = true;
}

async function readWasmBytes(resolvedUrl: string): Promise<ArrayBuffer> {
  const isNode = typeof process !== 'undefined' && !!process.versions?.node;
  if (!isNode && typeof globalThis.fetch !== 'undefined') {
    const resp = await fetch(resolvedUrl);
    if (!resp.ok) throw new Error(`Failed to fetch xaligo.wasm: ${resp.statusText}`);
    return resp.arrayBuffer();
  }

  const { readFile } = await import('node:fs/promises');
  const { fileURLToPath } = await import('node:url');
  const filePath = resolvedUrl.startsWith('file:') ? fileURLToPath(resolvedUrl) : resolvedUrl;
  return (await readFile(filePath)).buffer;
}

function waitFor(predicate: () => boolean, timeoutMs: number): Promise<void> {
  return new Promise((resolve, reject) => {
    if (predicate()) { resolve(); return; }
    const start = Date.now();
    const interval = setInterval(() => {
      if (predicate()) {
        clearInterval(interval);
        resolve();
      } else if (Date.now() - start > timeoutMs) {
        clearInterval(interval);
        reject(new Error('xaligo WASM did not initialise within the timeout'));
      }
    }, 10);
  });
}
