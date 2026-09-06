// Read public SDK model data without extracting or executing the downloaded archive.
import fs from 'node:fs';
import path from 'node:path';
import zlib from 'node:zlib';
import crypto from 'node:crypto';
import { fileURLToPath } from 'node:url';

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
if (!process.argv[2]) throw new Error('Usage: import_aws_api_research.mjs <botocore source tar.gz>');
const bytes = fs.readFileSync(process.argv[2]);
const tar = zlib.gunzipSync(bytes);
const inputs = new Map();
for (let offset = 0; offset + 512 <= tar.length;) {
  const header = tar.subarray(offset, offset + 512);
  const name = header.subarray(0, 100).toString().replace(/\0.*$/s, '');
  if (!name) break;
  const size = parseInt(header.subarray(124, 136).toString().replace(/\0.*$/s, '').trim(), 8);
  if (!Number.isSafeInteger(size) || size < 0 || offset + 512 + size > tar.length) throw new Error('Invalid tar entry');
  const match = name.match(/^botocore-[^/]+\/botocore\/data\/([^/]+)\/(\d{4}-\d{2}-\d{2})\/service-2\.json$/);
  if (match && (!inputs.has(match[1]) || inputs.get(match[1]).date < match[2])) {
    inputs.set(match[1], { date: match[2], model: JSON.parse(tar.subarray(offset + 512, offset + 512 + size)) });
  }
  offset += 512 + Math.ceil(size / 512) * 512;
}
const services = {};
for (const [name, { date, model }] of [...inputs].sort(([a], [b]) => a.localeCompare(b, 'en'))) {
  const typeOf = (shapeName, depth = 0) => {
    const shape = model.shapes[shapeName];
    if (!shape) return shapeName;
    if (shape.type === 'list' && depth < 3) return `List<${typeOf(shape.member.shape, depth + 1)}>`;
    if (shape.type === 'map' && depth < 3) return `Map<${typeOf(shape.value.shape, depth + 1)}>`;
    return shape.type === 'structure' ? shapeName : shape.type;
  };
  const fieldsOf = shapeName => {
    const shape = model.shapes[shapeName] ?? {};
    return Object.entries(shape.members ?? {}).map(([field, ref]) => ({name: field, type: typeOf(ref.shape), required: (shape.required ?? []).includes(field), ...(model.shapes[ref.shape]?.enum ? {values: model.shapes[ref.shape].enum} : {})}));
  };
  services[name] = {
    name: model.metadata.serviceFullName, id: model.metadata.serviceId, version: date,
    source: `https://github.com/boto/botocore/blob/develop/botocore/data/${name}/${date}/service-2.json`,
    operations: Object.entries(model.operations).map(([operation, spec]) => ({name: operation, input: fieldsOf(spec.input?.shape), output: fieldsOf(spec.output?.shape)})),
  };
}
const directory = path.join(root, 'docs/src/examples/samples/aws/research');
fs.mkdirSync(directory, {recursive: true});
const snapshot = {source: 'https://github.com/boto/botocore', archive: 'https://codeload.github.com/boto/botocore/tar.gz/refs/heads/develop', retrieved: '2026-09-06', sha256: crypto.createHash('sha256').update(bytes).digest('hex'), services};
fs.writeFileSync(path.join(directory, 'api-models.json'), JSON.stringify(snapshot) + '\n');
console.log(`Imported ${Object.keys(services).length} public AWS SDK service models.`);
