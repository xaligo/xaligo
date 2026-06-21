import type {
  IsoflowDocument,
  WasmResult,
  XaligoDiagnostic,
  XaligoWasm,
  XYFlowDocument,
} from '../entity/api';
import type { PptxExportOptions, PptxExportResult } from '../entity/pptx';
import { createWasmGateway, type WasmGateway } from '../repository/wasm';
import { drawPlanToPptx, pptxPlanOptionsJSON } from './pptx';

export { drawPlanToPptx, pptxPlanOptionsJSON } from './pptx';
export type {
  IsoflowColor,
  IsoflowConnector,
  IsoflowConnectorAnchor,
  IsoflowCoords,
  IsoflowDocument,
  IsoflowIcon,
  IsoflowModelItem,
  IsoflowRectangle,
  IsoflowView,
  IsoflowViewItem,
  XaligoDiagnostic,
  XaligoWasm,
  XYFlowDocument,
  XYFlowEdge,
  XYFlowNode,
} from '../entity/api';
export type {
  ArrowStyle,
  PaperOrientation,
  PaperSize,
  PptxExportOptions,
  PptxExportResult,
  PptxOutputType,
} from '../entity/pptx';

let instance: XaligoWasm | undefined;

export async function loadXaligo(wasmUrl?: string): Promise<XaligoWasm> {
  if (instance) return instance;
  instance = createXaligoUseCase(await createWasmGateway(wasmUrl));
  return instance;
}

export function createXaligoUseCase(wasm: WasmGateway): XaligoWasm {
  return {
    diagnose(xal: string): Promise<XaligoDiagnostic[]> {
      const res = ensureResult(wasm.diagnose(xal), 'xaligoDiagnose');
      return Promise.resolve(JSON.parse(res) as XaligoDiagnostic[]);
    },

    renderXYFlow(xal: string): Promise<XYFlowDocument> {
      const res = ensureResult(wasm.renderXYFlow(xal), 'xaligoRenderXYFlow');
      return Promise.resolve(JSON.parse(res) as XYFlowDocument);
    },

    renderIsoflow(xal: string): Promise<IsoflowDocument> {
      const res = ensureResult(wasm.renderIsoflow(xal), 'xaligoRenderIsoflow');
      return Promise.resolve(JSON.parse(res) as IsoflowDocument);
    },

    render(xal: string): Promise<string> {
      return Promise.resolve(ensureResult(wasm.render(xal), 'xaligoRender'));
    },

    renderWithServices(xal: string, servicesCsv: string): Promise<string> {
      return Promise.resolve(ensureResult(wasm.renderWithServices(xal, servicesCsv), 'xaligoRenderWithServices'));
    },

    async renderPptx(xal: string, options?: PptxExportOptions): Promise<PptxExportResult> {
      const optsJson = pptxPlanOptionsJSON(options);
      const plan = ensureResult(wasm.buildPptxPlan(xal, '', optsJson), 'xaligoBuildPptxPlan');
      return drawPlanToPptx(plan, options);
    },

    async renderWithServicesPptx(
      xal: string,
      servicesCsv: string,
      options?: PptxExportOptions,
    ): Promise<PptxExportResult> {
      const optsJson = pptxPlanOptionsJSON(options);
      const plan = ensureResult(wasm.buildPptxPlan(xal, servicesCsv, optsJson), 'xaligoBuildPptxPlan');
      return drawPlanToPptx(plan, options);
    },
  };
}

function ensureResult(res: WasmResult, functionName: string): string {
  if (res.error) throw new Error(res.error);
  if (!res.result) throw new Error(`${functionName} returned empty result`);
  return res.result;
}
