import fs from 'node:fs';
import path from 'node:path';
const root = path.resolve(import.meta.dirname, '../..');
const pkgDir = path.join(root, 'node_modules/@tabler/icons');
const metadata = JSON.parse(fs.readFileSync(path.join(pkgDir, 'icons.json'), 'utf8'));
const nodes = JSON.parse(fs.readFileSync(path.join(pkgDir, 'tabler-nodes-outline.json'), 'utf8'));
const outDir = path.join(root, 'etc/resources/aws/svg/Tabler-Icons');
const catalogPath = path.join(root, 'etc/resources/aws/service-catalog.csv');
const indexPath = path.join(root, 'etc/resources/aws/service-index.csv');
const firstID = 100000;

function escapeXML(value) {
  return String(value)
    .replaceAll('&', '&amp;')
    .replaceAll('"', '&quot;')
    .replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;');
}

function renderNode([tag, attrs]) {
  const attributes = Object.entries(attrs ?? {})
    .map(([key, value]) => `${key}="${escapeXML(value)}"`)
    .join(' ');
  return `  <${tag}${attributes ? ` ${attributes}` : ''} />`;
}

function renderSVG(name, iconNodes) {
  return [
    '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">',
    `  <title>${escapeXML(name)}</title>`,
    ...iconNodes.map(renderNode),
    '</svg>',
    '',
  ].join('\n');
}

function csv(value) {
  const text = String(value);
  return /[",\r\n]/.test(text) ? `"${text.replaceAll('"', '""')}"` : text;
}

fs.mkdirSync(outDir, { recursive: true });
for (const file of fs.readdirSync(outDir)) {
  if (file.endsWith('.svg')) fs.rmSync(path.join(outDir, file));
}

const names = Object.keys(nodes).filter((name) => metadata[name]).sort();
const catalogRows = [];
const indexRows = [];
for (const [offset, name] of names.entries()) {
  const id = firstID + offset;
  const file = `${name}.svg`;
  const relPath = `etc/resources/aws/svg/Tabler-Icons/${file}`;
  const svg = renderSVG(name, nodes[name]);
  fs.writeFileSync(path.join(outDir, file), svg);
  const dataURL = `data:image/svg+xml;base64,${Buffer.from(svg).toString('base64')}`;
  const category = `Tabler/${metadata[name].category || 'Uncategorized'}`;
  catalogRows.push([id, category, name, file, relPath, dataURL].map(csv).join(','));
  indexRows.push(`${id},${csv(name)}`);
}

const catalog = fs.readFileSync(catalogPath, 'utf8').replaceAll('\r\n', '\n');
const awsCatalog = catalog.split('\n').filter((line) => {
  const id = Number.parseInt(line.split(',', 1)[0], 10);
  return !Number.isFinite(id) || id < firstID || id >= 200000;
});
while (awsCatalog.at(-1) === '') awsCatalog.pop();
fs.writeFileSync(catalogPath, `${awsCatalog.join('\n')}\n${catalogRows.join('\n')}\n`);

const index = fs.readFileSync(indexPath, 'utf8').replaceAll('\r\n', '\n');
const awsIndex = index.split('\n').filter((line) => {
  const id = Number.parseInt(line.split(',', 1)[0], 10);
  return !Number.isFinite(id) || id < firstID || id >= 200000;
});
while (awsIndex.at(-1) === '') awsIndex.pop();
fs.writeFileSync(indexPath, `${awsIndex.join('\n')}\n${indexRows.join('\n')}\n`);

fs.copyFileSync(path.join(pkgDir, 'LICENSE'), path.join(outDir, 'LICENSE'));
console.log(`Imported ${names.length} Tabler icons as catalog IDs ${firstID}-${firstID + names.length - 1}`);
