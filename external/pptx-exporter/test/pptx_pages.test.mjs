import assert from 'node:assert/strict';
import test from 'node:test';
import { inflateSync } from 'node:zlib';

import JSZip from 'jszip';

import { drawPlanToPptx, parsePptxExporterRequest } from '../dist/index.js';

const pixelPNG = 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNgYAAAAAMAASsJTYQAAAAASUVORK5CYII=';
const iconSVG = '<!DOCTYPE svg><!-- native SVG fallback regression --><svg xmlns="http://www.w3.org/2000/svg" width="32" height="32"><rect width="32" height="32" fill="#2563eb"/></svg>';
const iconSVGData = `data:image/svg+xml;base64,${Buffer.from(iconSVG).toString('base64')}`;
const utf16IconSVG = Buffer.concat([Buffer.from([0xff, 0xfe]), Buffer.from(iconSVG, 'utf16le')]);
const utf16IconSVGData = `data:image/svg+xml;base64,${utf16IconSVG.toString('base64')}`;

function textLayout(overflow = 'clip') {
  return {
    role: 'label',
    wrap: false,
    fit: 'none',
    overflow,
    clip: overflow === 'clip',
    lineHeight: 1.2,
    padding: { top: 0, right: 0, bottom: 0, left: 0 },
  };
}

function lineStyle() {
  return {
    color: '2563EB',
    width: 1,
    dash: 'solid',
    transparency: 0,
    beginArrowType: 'none',
    endArrowType: 'stealth',
  };
}

async function slideXML(zip, number) {
  const file = zip.file(`ppt/slides/slide${number}.xml`);
  assert.ok(file, `slide ${number} is missing`);
  return file.async('string');
}

function assertSlideTransformsWithinBounds(xml, widthIn, heightIn) {
  const emuPerInch = 914400;
  const maxX = Math.round(widthIn * emuPerInch);
  const maxY = Math.round(heightIn * emuPerInch);
  const transforms = xml.matchAll(/<a:off x="(-?\d+)" y="(-?\d+)"\/><a:ext cx="(\d+)" cy="(\d+)"\/>/g);
  let count = 0;
  for (const match of transforms) {
    const [, xText, yText, widthText, heightText] = match;
    const x = Number(xText);
    const y = Number(yText);
    const width = Number(widthText);
    const height = Number(heightText);
    assert.ok(x >= 0 && y >= 0, `negative legend transform: ${match[0]}`);
    assert.ok(x + width <= maxX + 2, `legend transform exceeds slide width: ${match[0]}`);
    assert.ok(y + height <= maxY + 2, `legend transform exceeds slide height: ${match[0]}`);
    count++;
  }
  assert.ok(count > 0, 'legend slide contains no transforms');
}

function assertValidPNG(bytes) {
  const signature = [0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a];
  assert.deepEqual([...bytes.subarray(0, signature.length)], signature);

  let offset = signature.length;
  let ihdr;
  const idat = [];
  let foundIEND = false;
  while (offset < bytes.length) {
    assert.ok(offset + 12 <= bytes.length, 'PNG chunk header is truncated');
    const view = new DataView(bytes.buffer, bytes.byteOffset + offset);
    const length = view.getUint32(0);
    const dataStart = offset + 8;
    const dataEnd = dataStart + length;
    assert.ok(dataEnd + 4 <= bytes.length, 'PNG chunk data is truncated');

    const type = new TextDecoder().decode(bytes.subarray(offset + 4, offset + 8));
    const expectedCRC = new DataView(bytes.buffer, bytes.byteOffset + dataEnd, 4).getUint32(0);
    assert.equal(crc32(bytes.subarray(offset + 4, dataEnd)), expectedCRC, `${type} CRC`);

    if (type === 'IHDR') ihdr = bytes.subarray(dataStart, dataEnd);
    if (type === 'IDAT') idat.push(bytes.subarray(dataStart, dataEnd));
    offset = dataEnd + 4;
    if (type !== 'IEND') continue;
    assert.equal(length, 0);
    assert.equal(offset, bytes.length, 'PNG contains bytes after IEND');
    foundIEND = true;
    break;
  }

  assert.equal(ihdr?.length, 13);
  assert.ok(idat.length > 0, 'PNG has no IDAT chunk');
  assert.ok(foundIEND, 'PNG has no IEND chunk');
  assert.ok(inflateSync(Buffer.concat(idat)).length > 0, 'PNG image data is empty');
}

function crc32(bytes) {
  let crc = 0xffffffff;
  for (const byte of bytes) {
    crc ^= byte;
    for (let bit = 0; bit < 8; bit++) {
      crc = (crc >>> 1) ^ ((crc & 1) ? 0xedb88320 : 0);
    }
  }
  return (crc ^ 0xffffffff) >>> 0;
}

test('renders schemaVersion 2 pages before one shared set of legend slides', async () => {
  const bytes = await drawPlanToPptx({
    schemaVersion: 2,
    pages: [
      {
        id: 'overview',
        slide: { w: 4, h: 3, background: 'FFFFFF' },
        ops: [{ id: 'overview-shape', kind: 'rect', x: 0.4, y: 0.4, w: 1, h: 1 }],
      },
      {
        id: 'detail',
        slide: { w: 4, h: 3, background: 'F8FAFC' },
        ops: [
          {
            id: 'detail-line', frontLayer: true, kind: 'line',
            x: 0.5, y: 1.8, w: 2, h: 0,
            points: [{ x: 0, y: 0, moveTo: true }, { x: 2, y: 0 }],
            line: lineStyle(),
          },
          { id: 'detail-shape', groupId: 'detail-group', kind: 'rect', x: 0.5, y: 0.5, w: 1, h: 0.7 },
          {
            id: 'detail-label', groupId: 'detail-group', kind: 'text',
            x: 0.5, y: 1.2, w: 1, h: 0.3, text: 'Detail page', fontSize: 10,
            textLayout: textLayout(),
          },
        ],
      },
    ],
    connectorLegend: [{
      id: 'L01',
      kind: 'connection',
      label: 'Connection line',
      description: 'Shared connector legend',
      line: lineStyle(),
    }],
    legend: [{ catalogId: 27, abbreviation: 'EC2', officialName: 'Amazon EC2', data: pixelPNG }],
  }, { outputType: 'uint8array', compression: false });

  const zip = await JSZip.loadAsync(bytes);
  const slidePaths = Object.keys(zip.files)
    .filter((path) => /^ppt\/slides\/slide\d+[.]xml$/.test(path))
    .sort();
  assert.equal(slidePaths.length, 4);

  const overview = await slideXML(zip, 1);
  const detail = await slideXML(zip, 2);
  const connectorLegend = await slideXML(zip, 3);
  const serviceLegend = await slideXML(zip, 4);

  assert.match(overview, /name="overview-shape"/);
  assert.doesNotMatch(overview, /detail-shape/);
  assert.match(detail, /detail-shape/);
  assert.doesNotMatch(detail, /overview-shape/);
  assert.match(connectorLegend, /<a:t>Line Legend<\/a:t>/);
  assert.match(serviceLegend, /<a:t>Legend<\/a:t>/);
  assert.equal([overview, detail, connectorLegend, serviceLegend]
    .filter((xml) => xml.includes('<a:t>Line Legend</a:t>')).length, 1);
  assert.equal([overview, detail, connectorLegend, serviceLegend]
    .filter((xml) => xml.includes('<a:t>Legend</a:t>')).length, 1);
  assertSlideTransformsWithinBounds(connectorLegend, 4, 3);
  assertSlideTransformsWithinBounds(serviceLegend, 4, 3);

  // Package post-processing must run for every diagram page, not only slide 1.
  assert.match(detail, /<p:grpSp><p:nvGrpSpPr><p:cNvPr[^>]*name="xaligo anchor detail-group"/);
  assert.match(detail, /<a:bodyPr\b[^>]*horzOverflow="clip" vertOverflow="clip"/);
  assert.match(detail, /<a:tailEnd\b[^>]*type="stealth"[^>]*w="sm"[^>]*len="lg"/);
  assert.match(connectorLegend, /<a:tailEnd\b[^>]*type="stealth"[^>]*w="sm"[^>]*len="lg"/);
  assert.ok(detail.indexOf('xaligo anchor detail-group') < detail.indexOf('xaligo-front-layer|detail-line'));
});

test('continues to accept the legacy single-page plan contract', () => {
  const request = parsePptxExporterRequest(JSON.stringify({
    plan: {
      slide: { w: 4, h: 3, background: 'FFFFFF' },
      ops: [],
    },
  }));
  assert.equal(request.plan.slide.w, 4);
});

test('validates schemaVersion 2 pages and rejects mixed PowerPoint slide sizes', async () => {
  const request = parsePptxExporterRequest(JSON.stringify({
    plan: {
      schemaVersion: 2,
      pages: [{ id: 'main', slide: { w: 4, h: 3, background: 'FFFFFF' }, ops: [] }],
    },
  }));
  assert.equal(request.plan.pages[0].id, 'main');

  assert.throws(
    () => parsePptxExporterRequest(JSON.stringify({ plan: { schemaVersion: 2, pages: [] } })),
    /must contain non-empty pages/,
  );
  assert.throws(
    () => parsePptxExporterRequest(JSON.stringify({ plan: { schemaVersion: 3, pages: [] } })),
    /schemaVersion 2 pages plan or legacy/,
  );

  await assert.rejects(
    drawPlanToPptx({
      schemaVersion: 2,
      pages: [
        { id: 'one', slide: { w: 4, h: 3, background: 'FFFFFF' }, ops: [] },
        { id: 'two', slide: { w: 5, h: 3, background: 'FFFFFF' }, ops: [] },
      ],
    }),
    /must use one slide size/,
  );
});

test('preserves SVG plan and legend images as native PowerPoint media', async () => {
  const bytes = await drawPlanToPptx({
    slide: { w: 4, h: 3, background: 'FFFFFF' },
    ops: [
      {
        id: 'native-svg',
        kind: 'image',
        x: 0.5,
        y: 0.5,
        w: 0.5,
        h: 0.5,
        data: iconSVGData,
      },
      {
        id: 'native-utf16-svg',
        kind: 'image',
        x: 1.25,
        y: 0.5,
        w: 0.5,
        h: 0.5,
        data: utf16IconSVGData,
      },
    ],
    legend: [{
      catalogId: 27,
      abbreviation: 'EC2',
      officialName: 'Amazon EC2',
      data: iconSVGData,
    }],
  }, { outputType: 'uint8array', compression: false });

  const zip = await JSZip.loadAsync(bytes);
  const svgPaths = Object.keys(zip.files)
    .filter((path) => /^ppt\/media\/.+[.]svg$/.test(path))
    .sort();
  assert.equal(svgPaths.length, 3);
  const actualSVGs = [];
  for (const path of svgPaths) {
    actualSVGs.push(Buffer.from(await zip.file(path).async('uint8array')).toString('hex'));
  }
  assert.deepEqual(
    actualSVGs.sort(),
    [Buffer.from(iconSVG), Buffer.from(iconSVG), utf16IconSVG]
      .map((source) => source.toString('hex'))
      .sort(),
  );

  const pngPaths = Object.keys(zip.files)
    .filter((path) => /^ppt\/media\/.+[.]png$/.test(path))
    .sort();
  assert.equal(pngPaths.length, 3);
  for (const path of pngPaths) {
    const png = await zip.file(path).async('uint8array');
    assertValidPNG(png);
  }

  for (const slideNumber of [1, 2]) {
    const slide = await slideXML(zip, slideNumber);
    assert.match(slide, /<asvg:svgBlip\b/);
    const relationships = await zip
      .file(`ppt/slides/_rels/slide${slideNumber}.xml.rels`)
      .async('string');
    assert.match(relationships, /Target="\.\.\/media\/[^"]+[.]svg"/);
  }
});
