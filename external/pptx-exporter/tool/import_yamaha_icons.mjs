import fs from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';

const require = createRequire(new URL('../package.json', import.meta.url));
const JSZip = require('jszip');

const root = path.resolve(import.meta.dirname, '../..');
const sourceURL = 'https://network.yamaha.com/support/download/tool/network_iconset.zip';
const sourceFile = process.argv[2];
const outDir = path.join(root, 'etc/resources/aws/svg/Yamaha-Network-Icons');
const catalogPath = path.join(root, 'etc/resources/aws/service-catalog.csv');
const indexPath = path.join(root, 'etc/resources/aws/service-index.csv');
const firstID = 200000;

const zipBytes = sourceFile
  ? fs.readFileSync(path.resolve(sourceFile))
  : Buffer.from(await (await fetch(sourceURL)).arrayBuffer());
const archive = await JSZip.loadAsync(zipBytes);

function csv(value) {
  const text = String(value);
  return /[",\r\n]/.test(text) ? `"${text.replaceAll('"', '""')}"` : text;
}

fs.mkdirSync(outDir, { recursive: true });
for (const file of fs.readdirSync(outDir)) {
  if (file.endsWith('.svg')) fs.rmSync(path.join(outDir, file));
}

const svgEntries = Object.values(archive.files)
  .filter((entry) => !entry.dir && /^network_iconset\/[^/]+\/[^/]+\.svg$/i.test(entry.name))
  .sort((a, b) => a.name.localeCompare(b.name, 'en'));

const catalogRows = [];
const indexRows = [];
for (const [offset, entry] of svgEntries.entries()) {
  const id = firstID + offset;
  const name = path.posix.basename(entry.name, '.svg');
  const file = `${name}.svg`;
  const relPath = `etc/resources/aws/svg/Yamaha-Network-Icons/${file}`;
  const svg = await entry.async('nodebuffer');
  fs.writeFileSync(path.join(outDir, file), svg);
  const dataURL = `data:image/svg+xml;base64,${svg.toString('base64')}`;
  catalogRows.push([id, 'Yamaha/Network Icon Set', name, file, relPath, dataURL].map(csv).join(','));
  indexRows.push(`${id},${csv(name)}`);
}

const catalog = fs.readFileSync(catalogPath, 'utf8').replaceAll('\r\n', '\n');
const baseCatalog = catalog.split('\n').filter((line) => {
  const id = Number.parseInt(line.split(',', 1)[0], 10);
  return !Number.isFinite(id) || id < firstID;
});
while (baseCatalog.at(-1) === '') baseCatalog.pop();
fs.writeFileSync(catalogPath, `${baseCatalog.join('\n')}\n${catalogRows.join('\n')}\n`);

const index = fs.readFileSync(indexPath, 'utf8').replaceAll('\r\n', '\n');
const baseIndex = index.split('\n').filter((line) => {
  const id = Number.parseInt(line.split(',', 1)[0], 10);
  return !Number.isFinite(id) || id < firstID;
});
while (baseIndex.at(-1) === '') baseIndex.pop();
fs.writeFileSync(indexPath, `${baseIndex.join('\n')}\n${indexRows.join('\n')}\n`);

const attribution = `Yamaha Network Diagram Icons\n\nSource: ${sourceURL}\nCreator: Yamaha Corporation\nLicense: Creative Commons Attribution-NoDerivatives 4.0 International\nLicense URL: https://creativecommons.org/licenses/by-nd/4.0/\n\nThe SVG files in this directory are redistributed unchanged.\n`;
fs.writeFileSync(path.join(outDir, 'ATTRIBUTION.txt'), attribution);
console.log(`Imported ${svgEntries.length} Yamaha icons as catalog IDs ${firstID}-${firstID + svgEntries.length - 1}`);
