import assert from 'node:assert/strict';
import test from 'node:test';

function globalSnapshot(name) {
  return {
    own: Object.prototype.hasOwnProperty.call(globalThis, name),
    value: globalThis[name],
  };
}

test('does not install browser globals in Node API consumers', async () => {
  const originalJSZip = Object.getOwnPropertyDescriptor(globalThis, 'JSZip');
  const sentinel = Object.freeze({ host: 'existing-jszip-global' });
  Object.defineProperty(globalThis, 'JSZip', {
    configurable: true,
    enumerable: true,
    value: sentinel,
    writable: false,
  });

  try {
    const before = new Map(
      ['document', 'Image', 'JSZip'].map((name) => [name, globalSnapshot(name)]),
    );
    const { drawPlanToPptx } = await import('../dist/index.js');
    await drawPlanToPptx({
      slide: { w: 4, h: 3, background: 'FFFFFF' },
      ops: [],
    });

    for (const [name, snapshot] of before) {
      assert.deepEqual(globalSnapshot(name), snapshot, `${name} global changed`);
    }
  } finally {
    if (originalJSZip) Object.defineProperty(globalThis, 'JSZip', originalJSZip);
    else delete globalThis.JSZip;
  }
});
