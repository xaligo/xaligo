import type { ArrowStyle, PaperOrientation, PaperSize } from './pptx';

export interface CliArgs {
  xal?: string;
  services?: string;
  output?: string;
  title?: string;
  author?: string;
  company?: string;
  subject?: string;
  compression?: boolean;
  pxPerInch?: number;
  arrowStyle?: ArrowStyle;
  arrowStub?: number;
  arrowMargin?: number;
  paper?: PaperSize;
  orientation?: PaperOrientation;
  paperMargin?: number;
  paperMarginTop?: number;
  paperMarginRight?: number;
  paperMarginBottom?: number;
  paperMarginLeft?: number;
  help?: boolean;
}
