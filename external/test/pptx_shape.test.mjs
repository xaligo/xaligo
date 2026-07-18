import assert from 'node:assert/strict';
import test from 'node:test';

import JSZip from 'jszip';

import { drawPlanToPptx } from '../dist/index.js';

test('renders diamond plan operations as PowerPoint diamond shapes', async () => {
  const bytes = await drawPlanToPptx({
    slide: { w: 2, h: 2, background: 'FFFFFF' },
    ops: [{
      id: 'uml-decision',
      kind: 'diamond',
      x: 0.5,
      y: 0.5,
      w: 1,
      h: 1,
      line: { color: '1E1E1E', width: 1, dash: 'solid', transparency: 0 },
      fill: { color: 'FFFFFF', transparency: 0 },
    }],
  }, { outputType: 'uint8array', compression: false });

  const zip = await JSZip.loadAsync(bytes);
  const xml = await zip.file('ppt/slides/slide1.xml').async('string');
  const diamond = [...xml.matchAll(/<p:sp\b[\s\S]*?<\/p:sp>/g)]
    .map((match) => match[0])
    .find((block) => block.includes('name="uml-decision"'));

  assert.ok(diamond);
  assert.match(diamond, /<a:prstGeom prst="diamond">/);
});
