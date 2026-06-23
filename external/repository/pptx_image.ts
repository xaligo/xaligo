const pngDataCache = new Map<string, string>();

import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

const logger = NewEnvLogger('external/repository', 'pptx_image');
const ERPIFP001 = NewMCode('ERPIFP-001', 'Image data for PPTX passthrough branch');
const ERPIFP002 = NewMCode('ERPIFP-002', 'Image data for PPTX cache hit branch');
const ERPIFP003 = NewMCode('ERPIFP-003', 'Image data for PPTX rasterize start');
const ERPIFP004 = NewMCode('ERPIFP-004', 'Image data for PPTX rasterize completed');

export async function imageDataForPptx(data: string, widthInches: number): Promise<string> {
  const isNode =
    typeof process !== 'undefined' &&
    process.release?.name === 'node';
  if (!isNode || !data.startsWith('data:image/svg+xml;base64,')) {
    logger.DEBUG(ERPIFP001, 'branch passthrough', { isNode, dataPrefix: data.slice(0, 32) });
    return data;
  }

  const targetWidth = Math.max(32, Math.round(widthInches * 192));
  const cacheKey = `${targetWidth}:${data}`;
  const cached = pngDataCache.get(cacheKey);
  if (cached) {
    logger.DEBUG(ERPIFP002, 'branch cache hit', { targetWidth });
    return cached;
  }

  logger.DEBUG(ERPIFP003, 'start rasterize', { targetWidth });
  const encoded = data.slice(data.indexOf(',') + 1);
  const svg = Buffer.from(encoded, 'base64');
  const { Resvg } = await import('@resvg/resvg-js');
  const png = new Resvg(svg, {
    fitTo: { mode: 'width', value: targetWidth },
    font: { loadSystemFonts: false },
  }).render().asPng();
  const result = `data:image/png;base64,${png.toString('base64')}`;
  pngDataCache.set(cacheKey, result);
  logger.DEBUG(ERPIFP004, 'completed rasterize', { targetWidth, bytes: result.length });
  return result;
}
