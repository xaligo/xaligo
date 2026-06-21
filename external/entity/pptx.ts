export type PptxOutputType = 'arraybuffer' | 'base64' | 'blob' | 'nodebuffer' | 'uint8array';

export type ArrowStyle =
  | 'thin'
  | 'standard'
  | 'triangle'
  | 'stealth'
  | 'arrow'
  | 'diamond'
  | 'oval'
  | 'none';

export type PaperSize =
  | 'A5'
  | 'A4'
  | 'A3'
  | 'A2'
  | 'A1'
  | 'Letter'
  | 'Legal'
  | 'Tabloid';

export type PaperOrientation = 'portrait' | 'landscape';

export interface PptxExportOptions {
  title?: string;
  author?: string;
  company?: string;
  subject?: string;
  outputType?: PptxOutputType;
  compression?: boolean;
  theme?: 'light' | 'dark';
  pxPerInch?: number;
  arrowStyle?: ArrowStyle;
  arrowStubPx?: number;
  arrowMarginPx?: number;
  paperSize?: PaperSize;
  orientation?: PaperOrientation;
  paperMargin?: number;
  paperMarginTop?: number;
  paperMarginRight?: number;
  paperMarginBottom?: number;
  paperMarginLeft?: number;
}

export type PptxExportResult = string | ArrayBuffer | Blob | Uint8Array;

export interface PptxPlan {
  slide: PlanSlide;
  ops: PlanOp[];
  legend?: PlanLegendEntry[];
  connectorLegend?: PlanConnectorLegendEntry[];
}

export interface PlanSlide {
  w: number;
  h: number;
  background: string;
}

export interface PlanLegendEntry {
  catalogId: number;
  abbreviation: string;
  officialName: string;
  data?: string;
}

export interface PlanConnectorLegendEntry {
  id: string;
  kind: string;
  label: string;
  description: string;
  source?: string;
  target?: string;
  line: PlanLine;
}

export interface ConnectorLegendRow {
  ids: string[];
  kind: string;
  label: string;
  description: string;
  line: PlanLine;
}

export interface PlanLine {
  color: string;
  width: number;
  dash: 'solid' | 'dash' | 'dot';
  transparency: number;
  beginArrowType?: ArrowHeadType;
  endArrowType?: ArrowHeadType;
}

export interface PlanFill {
  color: string;
  transparency: number;
}

export interface PlanPoint {
  x: number;
  y: number;
  moveTo?: boolean;
}

export type PlanOpKind = 'rect' | 'ellipse' | 'polygon' | 'text' | 'image' | 'line';

export interface PlanOp {
  id?: string;
  groupId?: string;
  frontLayer?: boolean;
  kind: PlanOpKind;
  x: number;
  y: number;
  w: number;
  h: number;
  rotate?: number;
  line?: PlanLine;
  fill?: PlanFill;
  text?: string;
  color?: string;
  fontFace?: string;
  fontSize?: number;
  bold?: boolean;
  align?: string;
  valign?: string;
  data?: string;
  transparency?: number;
  points?: PlanPoint[];
  flipH?: boolean;
  flipV?: boolean;
}

export type ArrowHeadType = 'none' | 'arrow' | 'diamond' | 'oval' | 'stealth' | 'triangle';

export const ANCHOR_GROUP_MARKER = 'xaligo-anchor-group|';
export const FRONT_LAYER_MARKER = 'xaligo-front-layer|';
