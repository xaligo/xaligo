const pngDataCache = new Map<string, string>();

export async function imageDataForPptx(data: string, widthInches: number): Promise<string> {
  const isNode =
    typeof process !== 'undefined' &&
    process.release?.name === 'node';
  if (!isNode || !data.startsWith('data:image/svg+xml;base64,')) return data;

  const targetWidth = Math.max(32, Math.round(widthInches * 192));
  const cacheKey = `${targetWidth}:${data}`;
  const cached = pngDataCache.get(cacheKey);
  if (cached) return cached;

  const encoded = data.slice(data.indexOf(',') + 1);
  const svg = Buffer.from(encoded, 'base64');
  const { Resvg } = await import('@resvg/resvg-js');
  const png = new Resvg(svg, {
    fitTo: { mode: 'width', value: targetWidth },
    font: { loadSystemFonts: false },
  }).render().asPng();
  const result = `data:image/png;base64,${png.toString('base64')}`;
  pngDataCache.set(cacheKey, result);
  return result;
}
