import assert from 'node:assert/strict';
import test from 'node:test';

import JSZip from 'jszip';

import { drawPlanToPptx } from '../dist/index.js';

function textLayout(overflow) {
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

async function slideXML(ops) {
  const bytes = await drawPlanToPptx({
    slide: { w: 4, h: 1, background: 'FFFFFF' },
    ops,
  }, { outputType: 'uint8array', compression: false });
  const zip = await JSZip.loadAsync(bytes);
  return zip.file('ppt/slides/slide1.xml').async('string');
}

function shapeBlock(xml, encodedName) {
  return [...xml.matchAll(/<p:sp\b[\s\S]*?<\/p:sp>/g)]
    .map((match) => match[0])
    .find((block) => block.includes(`name="${encodedName}"`));
}

test('writes exact DrawingML clip and visible overflow without forcing shrink', async () => {
  const xml = await slideXML([
    {
      id: 'clip-text', kind: 'text', x: 0, y: 0, w: 1, h: 0.2,
      text: 'clipped long text', fontSize: 12, textLayout: textLayout('clip'),
    },
    {
      id: 'visible-text', kind: 'text', x: 1, y: 0, w: 1, h: 0.2,
      text: 'visible long text', fontSize: 12, textLayout: textLayout('visible'),
    },
  ]);

  const clipped = shapeBlock(xml, 'clip-text');
  const visible = shapeBlock(xml, 'visible-text');
  assert.match(clipped, /<a:bodyPr\b[^>]*horzOverflow="clip" vertOverflow="clip"/);
  assert.doesNotMatch(clipped, /<a:normAutofit\b/);
  assert.match(visible, /<a:bodyPr\b[^>]*horzOverflow="overflow" vertOverflow="overflow"/);
  assert.doesNotMatch(visible, /<a:normAutofit\b/);
});

test('finds text and anchor groups whose plan IDs contain XML metacharacters', async () => {
  const xml = await slideXML([
    {
      id: 'shape&1', groupId: 'group&1', kind: 'rect',
      x: 0, y: 0, w: 0.5, h: 0.5,
    },
    {
      id: 'label&1', groupId: 'group&1', kind: 'text',
      x: 0, y: 0.5, w: 0.5, h: 0.2, text: 'label',
      textLayout: textLayout('clip'),
    },
  ]);

  assert.match(xml, /<p:grpSp><p:nvGrpSpPr><p:cNvPr[^>]*name="xaligo anchor group&amp;1"/);
  assert.match(xml, /<a:bodyPr\b[^>]*horzOverflow="clip" vertOverflow="clip"/);
});
