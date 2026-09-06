import test from 'node:test';
import assert from 'node:assert/strict';
import fs from 'node:fs';
import path from 'node:path';
import {fileURLToPath} from 'node:url';
import {buildDesigns} from '../../../../scripts/tool/aws_design_model.mjs';
import {loadBalancerDesigns} from '../../../../scripts/tool/aws_load_balancer_designs.mjs';

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
const base = path.join(root, 'docs/src/examples/samples/aws');
const read = file => JSON.parse(fs.readFileSync(path.join(base, file), 'utf8'));
const catalog = read('catalog.json');
const cfn = read('research/cloudformation-models.json');
const api = read('research/api-models.json');
const designs = buildDesigns(catalog, cfn, api);

test('all AWS tags have an explicit, reproducible research classification and sources', () => {
  assert.equal(designs.length, 877);
  assert.deepEqual(JSON.parse(JSON.stringify(designs.map(({fields, ...d}) => d))), read('research/designs.json'));
  for (const d of designs) {
    assert.notEqual(d.kind, 'unmapped', d.tag);
    assert.ok(d.sources.length, d.tag);
    for (const source of d.sources) assert.equal(new URL(source).protocol, 'https:');
    for (const resource of d.resources) assert.ok(cfn.models[resource], resource);
    for (const nested of d.nested) assert.ok(cfn.models[nested.model], nested.model);
    for (const service of d.apis) assert.ok(api.services[service], service);
    if (d.primary) for (const field of d.fields) assert.equal(typeof field.type, 'string', d.primary + '.' + field.name);
  }
});

test('every editable AWS source has its SVG pair and research README', () => {
  let sources = 0;
  for (const d of catalog) {
    const directory = path.join(base, d.tag);
    const readme = fs.readFileSync(path.join(directory, 'README.md'), 'utf8');
    assert.ok(readme.includes('<!-- aws-functional-research:start -->'), d.tag);
    for (const filename of fs.readdirSync(directory).filter(file => file.endsWith('.xal'))) {
      const source = fs.readFileSync(path.join(directory, filename), 'utf8');
      assert.equal((source.match(/<frame\b/g) ?? []).length, 1, 'V2 supports one frame per source');
      assert.ok(!source.includes('undefined'), d.tag + '/' + filename);
      for (const match of source.matchAll(/<rectangle[^>]*height="(\d+)"[^>]*font-size="16"[^>]*title="([^"]*)"/g)) {
        assert.ok(match[2].split('&#10;').length * 19.2 + 24 <= Number(match[1]), 'Text exceeds card: ' + d.tag + '/' + filename);
      }
      assert.ok(fs.existsSync(path.join(directory, filename.replace(/\.xal$/, '.svg'))), d.tag + '/' + filename);
      sources++;
    }
  }
  assert.equal(sources, 906); // Includes 25 native ALB detail/service examples.
});

test('load balancer capabilities are not inferred from the shared ELBv2 schema', () => {
  const alb = loadBalancerDesigns['aws-elastic-load-balancing-application-load-balancer'];
  const nlb = loadBalancerDesigns['aws-elastic-load-balancing-network-load-balancer'];
  assert.ok(alb.cards.flat().some(line => line.includes('MutualAuthentication.Mode = verify')));
  assert.ok(nlb.cards.flat().some(line => line.includes('do NOT terminate mTLS')));
  assert.ok(nlb.cards.flat().some(line => line.includes('QUIC, TCP_QUIC')));
  assert.ok(designs.find(d => d.tag.endsWith('-classic-load-balancer')).primary.startsWith('AWS::ElasticLoadBalancing::'));
  for (const tag of ['aws-workdocs', 'aws-private-5g', 'aws-panorama', 'aws-lookout-for-metrics']) assert.equal(designs.find(d => d.tag === tag).lifecycle, 'retired', tag);
});

test('every closed native ALB option is documented and has an editable example', () => {
  const implementation = fs.readFileSync(path.join(root, 'external/engine/src/usc/aws/option.rs'), 'utf8');
  const names = [...implementation.matchAll(/name: "([^"]+)",[\s\S]*?label: "([^"]+)",[\s\S]*?owners: "([^"]+)",/g)].map(m => m[1]);
  assert.equal(names.length, 138);
  const docs = fs.readFileSync(path.join(root, 'docs/src/xal/alb-options.md'), 'utf8');
  const directory = path.join(base, 'aws-elastic-load-balancing-application-load-balancer');
  const examples = fs.readdirSync(directory).filter(f => f.endsWith('.xal')).map(f => fs.readFileSync(path.join(directory, f), 'utf8')).join('\n');
  for (const name of names) {
    assert.ok(docs.includes('`' + name + '`'), 'Undocumented option: ' + name);
    assert.ok(examples.includes('name="' + name + '"'), 'Missing option example: ' + name);
  }
});
