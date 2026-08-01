import assert from 'node:assert/strict';
import { spawnSync } from 'node:child_process';
import { readFileSync } from 'node:fs';
import test from 'node:test';

for (const outputFile of ['index.js', 'cli.js']) {
  test(`${outputFile} has stable embedded source maps`, () => {
    const code = readFileSync(new URL(`../dist/${outputFile}`, import.meta.url), 'utf8');
    const sourceMap = JSON.parse(
      readFileSync(new URL(`../dist/${outputFile}.map`, import.meta.url), 'utf8'),
    );
    const serializedMap = JSON.stringify(sourceMap);
    const mappedSources = sourceMap.sections.flatMap((section) => section.map.sources ?? []);
    const mappedSourceContents = sourceMap.sections.flatMap(
      (section) => section.map.sourcesContent ?? [],
    );

    assert.equal(sourceMap.version, 3);
    assert.equal(sourceMap.file, outputFile);
    assert.deepEqual(sourceMap.sources, []);
    assert.ok(sourceMap.sections.length > 2);
    assert.ok(mappedSources.includes('vendor/jszip.min.js'));
    assert.ok(mappedSources.includes('vendor/pptxgenjs.min.js'));
    assert.ok(mappedSources.some((source) => source.endsWith('.ts')));
    assert.ok(mappedSourceContents.length > 0);
    assert.ok(mappedSourceContents.every((content) => typeof content === 'string'));
    assert.equal(serializedMap.includes('.xaligo-build-'), false);
    assert.equal(serializedMap.includes(process.cwd()), false);
    assert.equal((code.match(/sourceMappingURL=/g) ?? []).length, 1);
    assert.ok(code.trimEnd().endsWith(`//# sourceMappingURL=${outputFile}.map`));

    for (let index = 1; index < sourceMap.sections.length; index += 1) {
      const previous = sourceMap.sections[index - 1].offset;
      const current = sourceMap.sections[index].offset;
      assert.ok(
        current.line > previous.line
          || (current.line === previous.line && current.column > previous.column),
        `source-map sections ${index - 1} and ${index} overlap`,
      );
    }
  });
}

test('Node stack traces resolve bundled code to TypeScript sources', () => {
  const script = `
    import { drawPlanToPptx } from ${JSON.stringify(
      new URL('../dist/index.js', import.meta.url).href,
    )};
    await drawPlanToPptx(null);
  `;
  const result = spawnSync(
    process.execPath,
    ['--enable-source-maps', '--input-type=module', '--eval', script],
    { encoding: 'utf8' },
  );

  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /entity[/\\]pptx[.]ts:\d+:\d+/);
  assert.doesNotMatch(result.stderr, /dist[/\\]index[.]js:\d+:\d+/);
});
