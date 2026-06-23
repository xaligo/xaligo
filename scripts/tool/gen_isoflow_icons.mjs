import fs from 'node:fs';
import path from 'node:path';

const root = path.resolve(import.meta.dirname, '../..');
const catalogPath = path.join(root, 'etc/resources/aws/service-catalog.csv');
const defaultOutDir = path.join(root, 'etc/resources/aws/svg/Isoflow-Icons');
const defaultManifestPath = path.join(root, 'etc/resources/aws/isoflow-icons.json');

const args = parseArgs(process.argv.slice(2));
const outDir = path.resolve(args.outDir ?? defaultOutDir);
const manifestPath = path.resolve(args.manifest ?? defaultManifestPath);
const ids = collectIDs(args);

if (ids.size === 0) {
  console.error('No IDs selected. Use --xal <file>, --services <csv>, or --ids <id,id,...>.');
  process.exit(1);
}

fs.mkdirSync(outDir, { recursive: true });

const catalog = parseCSV(fs.readFileSync(catalogPath, 'utf8'));
const header = catalog.shift();
if (!header || header.length < 6) throw new Error('service-catalog.csv header is invalid');

const icons = {};
let generated = 0;
for (const row of catalog) {
  if (row.length < 6) continue;
  const id = Number.parseInt(row[0], 10);
  if (!ids.has(id)) continue;

  const category = row[1];
  const service = row[2];
  const sourceFile = row[3];
  const sourceRelPath = row[4];
  const dataURL = normalizeDataURL(row[5]);
  const sourceSVG = readSourceSVG(sourceRelPath, dataURL);
  const rawAccent = dominantColor(sourceSVG) ?? '#7c3aed';
  const accent = normalizeAccent(rawAccent);
  const topFill = '#ffffff';
  const slug = slugify(service || sourceFile || String(id));
  const svgFile = `isoflow-${id}-${slug}.svg`;
  const relPath = `etc/resources/aws/svg/Isoflow-Icons/${svgFile}`;
  const svg = renderIsoflowIcon({ service, accent, topFill, sourceDataURL: dataURL });

  fs.writeFileSync(path.join(outDir, svgFile), svg);
  icons[`item-cat-${id}`] = {
    catalogId: id,
    service,
    category,
    sourceFile,
    svgFile,
    relPath,
    dataURL: `data:image/svg+xml;base64,${Buffer.from(svg).toString('base64')}`,
  };
  generated++;
}

const manifest = {
  version: 1,
  kind: 'xaligo-isoflow-icons',
  generatedAt: new Date().toISOString(),
  icons,
};

fs.writeFileSync(manifestPath, `${JSON.stringify(manifest, null, 2)}\n`);
console.log(`Generated ${generated} Isoflow icons into ${path.relative(root, outDir)}`);
console.log(`Wrote ${path.relative(root, manifestPath)}`);

function parseArgs(values) {
  const result = { ids: [], xal: [], services: [] };
  for (let i = 0; i < values.length; i++) {
    const value = values[i];
    if (value === '--ids') result.ids.push(...String(values[++i] ?? '').split(','));
    else if (value === '--xal') result.xal.push(values[++i]);
    else if (value === '--services') result.services.push(values[++i]);
    else if (value === '--out-dir') result.outDir = values[++i];
    else if (value === '--manifest') result.manifest = values[++i];
    else if (value === '--help' || value === '-h') {
      printHelp();
      process.exit(0);
    } else {
      throw new Error(`unknown argument: ${value}`);
    }
  }
  return result;
}

function printHelp() {
  console.log(`Usage:
  npm run gen:isoflow-icons -- --xal examples/sample.xal
  npm run gen:isoflow-icons -- --services examples/services.csv
  npm run gen:isoflow-icons -- --ids 27,117,1020

Options:
  --xal <file>       Collect <item id="N"> values from a .xal file.
  --services <csv>   Collect catalog IDs from a services.csv file.
  --ids <list>       Comma-separated catalog IDs.
  --out-dir <dir>    Output SVG directory. Default: etc/resources/aws/svg/Isoflow-Icons
  --manifest <file>  Output manifest. Default: etc/resources/aws/isoflow-icons.json`);
}

function collectIDs(options) {
  const ids = new Set();
  for (const raw of options.ids) addID(ids, raw);
  for (const file of options.xal) {
    const source = fs.readFileSync(path.resolve(file), 'utf8');
    for (const match of source.matchAll(/<item\b[^>]*\bid\s*=\s*["'](\d+)["']/gi)) addID(ids, match[1]);
  }
  for (const file of options.services) {
    const rows = parseCSV(fs.readFileSync(path.resolve(file), 'utf8'));
    for (const row of rows) addID(ids, row[0]);
  }
  return ids;
}

function addID(ids, value) {
  const id = Number.parseInt(String(value).trim(), 10);
  if (Number.isFinite(id)) ids.add(id);
}

function parseCSV(text) {
  const rows = [];
  let row = [];
  let field = '';
  let quoted = false;
  for (let i = 0; i < text.length; i++) {
    const char = text[i];
    if (quoted) {
      if (char === '"' && text[i + 1] === '"') {
        field += '"';
        i++;
      } else if (char === '"') quoted = false;
      else field += char;
      continue;
    }
    if (char === '"') quoted = true;
    else if (char === ',') {
      row.push(field);
      field = '';
    } else if (char === '\n') {
      row.push(field);
      rows.push(row);
      row = [];
      field = '';
    } else if (char !== '\r') field += char;
  }
  if (field !== '' || row.length > 0) {
    row.push(field);
    rows.push(row);
  }
  return rows.filter((candidate) => candidate.length > 0 && !String(candidate[0] ?? '').startsWith('#'));
}

function normalizeDataURL(value) {
  const trimmed = String(value ?? '').trim();
  if (trimmed.startsWith('data:')) return trimmed;
  return `data:image/svg+xml;base64,${trimmed}`;
}

function readSourceSVG(relPath, dataURL) {
  const sourcePath = path.join(root, relPath);
  if (relPath && fs.existsSync(sourcePath)) return fs.readFileSync(sourcePath, 'utf8');
  const prefix = 'data:image/svg+xml;base64,';
  if (dataURL.startsWith(prefix)) return Buffer.from(dataURL.slice(prefix.length), 'base64').toString('utf8');
  return '';
}

function dominantColor(svg) {
  const matches = svg.match(/#[0-9a-fA-F]{6}\b/g) ?? [];
  return matches.find((color) => !['#ffffff', '#000000'].includes(color.toLowerCase()));
}

function renderIsoflowIcon({ service, accent, topFill, sourceDataURL }) {
  const sideLeft = mix(accent, '#111827', 0.38);
  const sideRight = mix(accent, '#111827', 0.18);
  const shadow = mix(accent, '#111827', 0.62);
  return `<svg xmlns="http://www.w3.org/2000/svg" width="160" height="124" viewBox="0 0 160 124" role="img">
  <title>${escapeXML(service)}</title>
  <defs>
    <filter id="soft-shadow" x="-20%" y="-20%" width="140%" height="150%">
      <feDropShadow dx="0" dy="7" stdDeviation="5" flood-color="${escapeXML(shadow)}" flood-opacity="0.30" />
    </filter>
    <filter id="glyph-shadow" x="-40%" y="-40%" width="180%" height="180%">
      <feDropShadow dx="0" dy="1" stdDeviation="1.1" flood-color="#0f172a" flood-opacity="0.36" />
    </filter>
  </defs>
  <g filter="url(#soft-shadow)">
    <polygon points="24,38 80,72 80,124 24,90" fill="${escapeXML(sideLeft)}" />
    <polygon points="136,38 80,72 80,124 136,90" fill="${escapeXML(sideRight)}" />
    <polygon points="80,6 136,38 80,72 24,38" fill="${escapeXML(topFill)}" />
    <image href="${escapeXML(sourceDataURL)}" x="50" y="19" width="60" height="38" preserveAspectRatio="xMidYMid meet" filter="url(#glyph-shadow)" />
    <polyline points="24,38 80,72 136,38" fill="none" stroke="rgba(255,255,255,0.45)" stroke-width="1.4" />
    <polyline points="24,90 80,124 136,90" fill="none" stroke="rgba(15,23,42,0.28)" stroke-width="2" />
    <line x1="80" y1="72" x2="80" y2="124" stroke="rgba(15,23,42,0.24)" stroke-width="1.6" />
  </g>
</svg>
`;
}

function mix(a, b, ratio) {
  const ca = hexToRGB(a);
  const cb = hexToRGB(b);
  const value = ca.map((channel, i) => Math.round(channel * (1 - ratio) + cb[i] * ratio));
  return `#${value.map((channel) => channel.toString(16).padStart(2, '0')).join('')}`;
}

function normalizeAccent(color) {
  const rgb = hexToRGB(color);
  const luminance = (0.2126 * rgb[0] + 0.7152 * rgb[1] + 0.0722 * rgb[2]) / 255;
  if (luminance > 0.78) return mix(color, '#334155', 0.42);
  if (luminance < 0.18) return mix(color, '#ffffff', 0.24);
  return color;
}

function hexToRGB(value) {
  const hex = value.replace('#', '');
  return [0, 2, 4].map((offset) => Number.parseInt(hex.slice(offset, offset + 2), 16));
}

function slugify(value) {
  return String(value).toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/^-|-$/g, '').slice(0, 80) || 'icon';
}

function escapeXML(value) {
  return String(value ?? '')
    .replaceAll('&', '&amp;')
    .replaceAll('"', '&quot;')
    .replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;');
}