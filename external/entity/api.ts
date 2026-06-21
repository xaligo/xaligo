import type { PptxExportOptions, PptxExportResult } from './pptx';

export interface WasmResult {
  result?: string;
  error?: string;
}

export interface XaligoDiagnostic {
  severity: 'error';
  message: string;
  offset?: number;
  line?: number;
  column?: number;
}

export interface XYFlowNode {
  id: string;
  type: string;
  position: { x: number; y: number };
  width: number;
  height: number;
  parentId?: string;
  extent?: 'parent';
  data: Record<string, unknown>;
  style?: Record<string, unknown>;
}

export interface XYFlowEdge {
  id: string;
  source: string;
  target: string;
  sourceHandle?: string;
  targetHandle?: string;
  type: string;
  zIndex?: number;
  data: Record<string, unknown>;
  style?: Record<string, unknown>;
  markerStart?: { type: string; color?: string };
  markerEnd?: { type: string; color?: string };
}

export interface XYFlowDocument {
  nodes: XYFlowNode[];
  edges: XYFlowEdge[];
  viewport: { x: number; y: number; zoom: number };
  width: number;
  height: number;
  background: string;
}

export interface IsoflowCoords { x: number; y: number; }

export interface IsoflowModelItem {
  id: string;
  name: string;
  description?: string;
  icon?: string;
}

export interface IsoflowViewItem {
  id: string;
  tile: IsoflowCoords;
}

export interface IsoflowRectangle {
  id: string;
  color?: string;
  from: IsoflowCoords;
  to: IsoflowCoords;
}

export interface IsoflowConnectorAnchor {
  id: string;
  ref: { item?: string; anchor?: string; tile?: IsoflowCoords };
}

export interface IsoflowConnector {
  id: string;
  description?: string;
  color?: string;
  width?: number;
  style?: 'SOLID' | 'DOTTED' | 'DASHED';
  anchors: IsoflowConnectorAnchor[];
}

export interface IsoflowView {
  id: string;
  lastUpdated?: string;
  name: string;
  description?: string;
  items: IsoflowViewItem[];
  rectangles?: IsoflowRectangle[];
  connectors?: IsoflowConnector[];
}

export interface IsoflowIcon {
  id: string;
  name: string;
  url: string;
  collection?: string;
  isIsometric?: boolean;
}

export interface IsoflowColor {
  id: string;
  value: string;
}

export interface IsoflowDocument {
  version: string;
  title: string;
  description?: string;
  items: IsoflowModelItem[];
  views: IsoflowView[];
  icons: IsoflowIcon[];
  colors: IsoflowColor[];
  fitToView?: boolean;
}

export interface XaligoWasm {
  diagnose(xal: string): Promise<XaligoDiagnostic[]>;
  renderXYFlow(xal: string): Promise<XYFlowDocument>;
  renderIsoflow(xal: string): Promise<IsoflowDocument>;
  render(xal: string): Promise<string>;
  renderWithServices(xal: string, servicesCsv: string): Promise<string>;
  renderPptx(xal: string, options?: PptxExportOptions): Promise<PptxExportResult>;
  renderWithServicesPptx(
    xal: string,
    servicesCsv: string,
    options?: PptxExportOptions,
  ): Promise<PptxExportResult>;
}
