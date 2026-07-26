import JSZip from 'jszip';

import {
  ANCHOR_GROUP_MARKER,
  FRONT_LAYER_MARKER,
  planObjectName,
  planTextOverflow,
  type PlanOp,
  type PptxExportResult,
  type PptxOutputType,
} from '../entity/pptx';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';
import { PPTX_FALLBACK_PNG_BYTES } from '../share/pptx_runtime';

interface XmlObjectBlock {
  start: number;
  end: number;
  xml: string;
  groupId?: string;
}

interface XmlBounds {
  x: number;
  y: number;
  cx: number;
  cy: number;
}

const logger = NewEnvLogger('external/repository', 'pptx_package');
const ERPPGAOIP001 = NewMCode('ERPPGAOIP-001', 'Finalize PPTX package unnecessary branch');
const ERPPGAOIP002 = NewMCode('ERPPGAOIP-002', 'Finalize PPTX package missing slide branch');
const ERPPGAOIP003 = NewMCode('ERPPGAOIP-003', 'Finalize PPTX package completed');
const ERPPNSFI001 = NewMCode('ERPPNSFI-001', 'Normalize SVG fallback image completed');
const ERPPCPO001 = NewMCode('ERPPCPO-001', 'Convert PPTX output arraybuffer branch');
const ERPPCPO002 = NewMCode('ERPPCPO-002', 'Convert PPTX output base64 branch');
const ERPPCPO003 = NewMCode('ERPPCPO-003', 'Convert PPTX output blob branch');
const ERPPCPO004 = NewMCode('ERPPCPO-004', 'Convert PPTX output nodebuffer branch');
const ERPPCPO005 = NewMCode('ERPPCPO-005', 'Convert PPTX output uint8array branch');
const ERPPGSO001 = NewMCode('ERPPGSO-001', 'Group slide objects insufficient objects branch');
const ERPPGSO002 = NewMCode('ERPPGSO-002', 'Group slide objects missing bounds branch');
const ERPPGSO003 = NewMCode('ERPPGSO-003', 'Group slide objects completed');
const ERPPMAALOTF001 = NewMCode('ERPPMAALOTF-001', 'Move anchor and line objects to front empty branch');
const ERPPMAALOTF002 = NewMCode('ERPPMAALOTF-002', 'Move anchor and line objects to front missing tree branch');
const ERPPMAALOTF003 = NewMCode('ERPPMAALOTF-003', 'Move anchor and line objects to front completed');
const ERPPGB001 = NewMCode('ERPPGB-001', 'Group bounds missing branch');

export interface PptxPackageSlidePlan {
  slideNumber: number;
  ops: PlanOp[];
  normalizeStealthArrowheads?: boolean;
}

export async function finalizePptxPackage(bytes: Uint8Array, pages: PptxPackageSlidePlan[], compression: boolean): Promise<Uint8Array> {
  const zip = await JSZip.loadAsync(bytes);
  const normalizedFallbacks = await normalizeSvgFallbackImages(zip);
  if (!pages.some(requiresPackageFinalization) && normalizedFallbacks === 0) {
    logger.DEBUG(ERPPGAOIP001, 'branch no package finalization');
    return bytes;
  }

  let finalizedSlides = 0;
  let finalizedGroups = 0;
  for (const page of pages) {
    if (!requiresPackageFinalization(page)) continue;
    const slidePath = `ppt/slides/slide${page.slideNumber}.xml`;
    const slide = zip.file(slidePath);
    if (!slide) {
      logger.WARN(ERPPGAOIP002, 'branch missing slide', { slidePath });
      continue;
    }

    const groupIds = [...new Set(page.ops.map((op) => op.groupId).filter((id): id is string => !!id))];
    let xml = await slide.async('string');
    xml = applyTextOverflowPolicies(xml, page.ops);
    xml = applySlenderStealthArrowheads(xml);
    for (const groupId of groupIds.sort()) {
      xml = groupSlideObjects(xml, groupId);
    }
    xml = moveAnchorAndLineObjectsToFront(xml);
    zip.file(slidePath, xml);
    finalizedSlides++;
    finalizedGroups += groupIds.length;
  }
  if (finalizedSlides === 0 && normalizedFallbacks === 0) return bytes;
  const out = await zip.generateAsync({ type: 'uint8array', compression: compression ? 'DEFLATE' : 'STORE' });
  logger.DEBUG(ERPPGAOIP003, 'completed', {
    slides: finalizedSlides,
    groups: finalizedGroups,
    normalizedFallbacks,
    bytes: out.length,
  });
  return out;
}

async function normalizeSvgFallbackImages(zip: JSZip): Promise<number> {
  let normalized = 0;
  const paths = Object.keys(zip.files)
    .filter((path) => /^ppt\/media\/.+[.]png$/.test(path))
    .sort();
  for (const path of paths) {
    const file = zip.file(path);
    if (!file) continue;
    const bytes = await file.async('uint8array');
    if (hasPNGSignature(bytes) || !hasSVGSignature(bytes)) continue;
    zip.file(path, PPTX_FALLBACK_PNG_BYTES);
    normalized++;
  }
  if (normalized > 0) logger.DEBUG(ERPPNSFI001, 'completed', { images: normalized });
  return normalized;
}

function hasPNGSignature(bytes: Uint8Array): boolean {
  const signature = PPTX_FALLBACK_PNG_BYTES.subarray(0, 8);
  return bytes.length >= signature.length
    && signature.every((value, index) => bytes[index] === value);
}

function hasSVGSignature(bytes: Uint8Array): boolean {
  const source = decodeSVGSource(bytes);
  if (source === undefined) return false;

  let offset = source.charCodeAt(0) === 0xfeff ? 1 : 0;
  for (;;) {
    while (/\s/.test(source[offset] ?? '')) offset++;
    if (source.startsWith('<?', offset)) {
      const end = source.indexOf('?>', offset + 2);
      if (end < 0) return false;
      offset = end + 2;
      continue;
    }
    if (source.startsWith('<!--', offset)) {
      const end = source.indexOf('-->', offset + 4);
      if (end < 0) return false;
      offset = end + 3;
      continue;
    }
    if (source.slice(offset, offset + 9).toLowerCase() === '<!doctype') {
      const end = markupDeclarationEnd(source, offset + 9);
      if (end < 0) return false;
      offset = end + 1;
      continue;
    }
    break;
  }
  return /^<(?:[A-Za-z_][\w.-]*:)?svg(?:[\s/>])/i.test(source.slice(offset));
}

function decodeSVGSource(bytes: Uint8Array): string | undefined {
  const isUTF16LE = (bytes[0] === 0xff && bytes[1] === 0xfe)
    || (bytes[0] === 0x3c && bytes[1] === 0);
  const isUTF16BE = (bytes[0] === 0xfe && bytes[1] === 0xff)
    || (bytes[0] === 0 && bytes[1] === 0x3c);
  if (isUTF16LE || isUTF16BE) {
    if (bytes.length % 2 !== 0) return undefined;
    let source = '';
    const chunk: number[] = [];
    for (let index = 0; index < bytes.length; index += 2) {
      const first = bytes[index] ?? 0;
      const second = bytes[index + 1] ?? 0;
      chunk.push(isUTF16LE ? first | (second << 8) : (first << 8) | second);
      if (chunk.length < 4096) continue;
      source += String.fromCharCode(...chunk);
      chunk.length = 0;
    }
    return source + String.fromCharCode(...chunk);
  }
  try {
    return new TextDecoder('utf-8', { fatal: true }).decode(bytes);
  } catch {
    return undefined;
  }
}

function markupDeclarationEnd(source: string, offset: number): number {
  let quote = '';
  let subsetDepth = 0;
  for (let index = offset; index < source.length; index++) {
    const char = source[index] ?? '';
    if (quote) {
      if (char === quote) quote = '';
      continue;
    }
    if (char === '"' || char === '\'') {
      quote = char;
      continue;
    }
    if (char === '[') subsetDepth++;
    else if (char === ']' && subsetDepth > 0) subsetDepth--;
    else if (char === '>' && subsetDepth === 0) return index;
  }
  return -1;
}

function requiresPackageFinalization(page: PptxPackageSlidePlan): boolean {
  return page.normalizeStealthArrowheads === true
    || page.ops.some((op) => !!op.groupId)
    || page.ops.some((op) => op.kind === 'text' && !!op.text && !!planObjectName(op))
    || page.ops.some((op) => op.kind === 'line' || op.frontLayer);
}

function applyTextOverflowPolicies(xml: string, ops: PlanOp[]): string {
  const policies = new Map<string, 'clip' | 'overflow'>();
  for (const op of ops) {
    if (op.kind !== 'text' || !op.text) continue;
    const objectName = planObjectName(op);
    if (!objectName) continue;
    policies.set(objectName, planTextOverflow(op) === 'clip' ? 'clip' : 'overflow');
  }
  if (policies.size === 0) return xml;

  return xml.replace(/<p:sp\b[\s\S]*?<\/p:sp>/g, (block) => {
    const objectName = pptxObjectName(block);
    const overflow = objectName ? policies.get(objectName) : undefined;
    if (!overflow || !/<p:txBody\b/.test(block)) return block;
    return block.replace(/<a:bodyPr\b([^>]*?)(\/?)>/, (_match, attrs: string, selfClose: string) => {
      const preserved = attrs
        .replace(/\s+(?:horzOverflow|vertOverflow)="[^"]*"/g, '')
        .replace(/\s+$/, '');
      return `<a:bodyPr${preserved} horzOverflow="${overflow}" vertOverflow="${overflow}"${selfClose}>`;
    });
  });
}

export function convertPptxOutput(bytes: Uint8Array, outputType: PptxOutputType): PptxExportResult {
  switch (outputType) {
    case 'arraybuffer':
      logger.DEBUG(ERPPCPO001, 'branch arraybuffer');
      return toArrayBuffer(bytes);
    case 'base64':
      logger.DEBUG(ERPPCPO002, 'branch base64');
      return bytesToBase64(bytes);
    case 'blob':
      logger.DEBUG(ERPPCPO003, 'branch blob');
      return new Blob([toArrayBuffer(bytes)], { type: 'application/vnd.openxmlformats-officedocument.presentationml.presentation' });
    case 'nodebuffer':
      logger.DEBUG(ERPPCPO004, 'branch nodebuffer');
      return Buffer.from(bytes);
    case 'uint8array':
    default:
      logger.DEBUG(ERPPCPO005, 'branch uint8array');
      return bytes;
  }
}

function groupSlideObjects(xml: string, groupId: string): string {
  const blocks = collectObjectBlocks(xml);
  const groupedIndexes = blocks
    .map((block, index) => (block.groupId === groupId ? index : -1))
    .filter((index) => index >= 0);
  if (groupedIndexes.length < 2) {
    logger.DEBUG(ERPPGSO001, 'branch insufficient objects', { groupId, objects: groupedIndexes.length });
    return xml;
  }

  const groupedBlocks = groupedIndexes.map((index) => blocks[index]).filter((block): block is XmlObjectBlock => !!block);
  const bounds = groupBounds(groupedBlocks);
  if (!bounds) {
    logger.WARN(ERPPGSO002, 'branch missing bounds', { groupId });
    return xml;
  }

  const groupedSet = new Set(groupedIndexes);
  const insertionIndex = groupedIndexes[0];
  const groupXML = groupShapeXML(groupId, groupedBlocks.map((block) => block.xml).join(''), bounds, nextNvPrID(xml));

  let out = '';
  let cursor = 0;
  blocks.forEach((block, index) => {
    if (!groupedSet.has(index)) {
      out += xml.slice(cursor, block.end);
      cursor = block.end;
      return;
    }
    out += xml.slice(cursor, block.start);
    cursor = block.end;
    if (index === insertionIndex) out += groupXML;
  });
  out += xml.slice(cursor);
  logger.DEBUG(ERPPGSO003, 'completed', { groupId, objects: groupedBlocks.length });
  return out;
}

function collectObjectBlocks(xml: string): XmlObjectBlock[] {
  const blocks: XmlObjectBlock[] = [];
  const re = /<p:(grpSp|sp|pic)\b[\s\S]*?<\/p:\1>/g;
  let match: RegExpExecArray | null;
  while ((match = re.exec(xml))) {
    const blockXML = match[0];
    const block: XmlObjectBlock = {
      start: match.index,
      end: match.index + blockXML.length,
      xml: blockXML,
    };
    const groupId = groupIdFromObjectBlock(blockXML);
    if (groupId) block.groupId = groupId;
    blocks.push(block);
  }
  return blocks;
}

function groupIdFromObjectBlock(xml: string): string | undefined {
  const name = pptxObjectName(xml);
  if (!name?.startsWith(ANCHOR_GROUP_MARKER)) return undefined;
  const rest = name.slice(ANCHOR_GROUP_MARKER.length);
  const separator = rest.indexOf('|');
  return separator >= 0 ? rest.slice(0, separator) : undefined;
}

function moveAnchorAndLineObjectsToFront(xml: string): string {
  const blocks = collectObjectBlocks(xml);
  const movingIndexes = blocks
    .map((block, index) => (isAnchorGroupBlock(block.xml) || isFrontLayerBlock(block.xml) ? index : -1))
    .filter((index) => index >= 0);
  if (movingIndexes.length === 0) {
    logger.DEBUG(ERPPMAALOTF001, 'branch empty');
    return xml;
  }

  const movingSet = new Set(movingIndexes);
  const lineAndMaskXML = movingIndexes
    .map((index) => blocks[index])
    .filter((block): block is XmlObjectBlock => !!block && !isAnchorGroupBlock(block.xml))
    .map((block) => block.xml)
    .join('');
  const anchorXML = movingIndexes
    .map((index) => blocks[index])
    .filter((block): block is XmlObjectBlock => !!block && isAnchorGroupBlock(block.xml))
    .map((block) => block.xml)
    .join('');

  let out = '';
  let cursor = 0;
  blocks.forEach((block, index) => {
    if (!movingSet.has(index)) {
      out += xml.slice(cursor, block.end);
    } else {
      out += xml.slice(cursor, block.start);
    }
    cursor = block.end;
  });
  out += xml.slice(cursor);
  const spTreeClose = out.lastIndexOf('</p:spTree>');
  if (spTreeClose < 0) {
    logger.WARN(ERPPMAALOTF002, 'branch missing tree');
    return out;
  }
  logger.DEBUG(ERPPMAALOTF003, 'completed', { objects: movingIndexes.length });
  return `${out.slice(0, spTreeClose)}${lineAndMaskXML}${anchorXML}${out.slice(spTreeClose)}`;
}

function isAnchorGroupBlock(xml: string): boolean {
  return /<p:grpSp\b/.test(xml) && /<p:cNvPr\b[^>]*\bname="xaligo anchor xaligo-anchor-/.test(xml);
}

function isFrontLayerBlock(xml: string): boolean {
  const name = pptxObjectName(xml);
  return !!name?.startsWith(FRONT_LAYER_MARKER);
}

// PptxGenJS currently escapes objectName once while constructing DrawingML
// and the XML serializer escapes that value again. Decode exactly those two
// layers before comparing with PlanOp IDs; otherwise IDs containing XML
// metacharacters silently miss grouping, layering, and text overflow policies.
function pptxObjectName(xml: string): string | undefined {
  const encoded = /<p:cNvPr\b[^>]*\bname="([^"]*)"/.exec(xml)?.[1];
  if (encoded === undefined) return undefined;
  return decodeXmlAttr(decodeXmlAttr(encoded));
}

function decodeXmlAttr(value: string): string {
  return value
    .replace(/&#x([0-9a-f]+);/gi, (match, digits: string) => decodeXmlCodePoint(match, digits, 16))
    .replace(/&#([0-9]+);/g, (match, digits: string) => decodeXmlCodePoint(match, digits, 10))
    .replace(/&quot;/g, '"')
    .replace(/&apos;|&#39;/g, "'")
    .replace(/&lt;/g, '<')
    .replace(/&gt;/g, '>')
    .replace(/&amp;/g, '&');
}

function decodeXmlCodePoint(source: string, digits: string, radix: number): string {
  const value = Number.parseInt(digits, radix);
  return Number.isInteger(value) && value >= 0 && value <= 0x10ffff
    ? String.fromCodePoint(value)
    : source;
}

function applySlenderStealthArrowheads(xml: string): string {
  return xml.replace(/<a:(headEnd|tailEnd)\b([^>]*\btype="stealth"[^>]*)\/>/g, (_match, tag: string, attrs: string) => {
    const width = /\bw="/.test(attrs) ? '' : ' w="sm"';
    const length = /\blen="/.test(attrs) ? '' : ' len="lg"';
    return `<a:${tag}${attrs}${width}${length}/>`;
  });
}

function groupBounds(blocks: XmlObjectBlock[]): XmlBounds | undefined {
  let minX = Number.POSITIVE_INFINITY;
  let minY = Number.POSITIVE_INFINITY;
  let maxX = Number.NEGATIVE_INFINITY;
  let maxY = Number.NEGATIVE_INFINITY;
  for (const block of blocks) {
    const bounds = objectBounds(block.xml);
    if (!bounds) continue;
    minX = Math.min(minX, bounds.x);
    minY = Math.min(minY, bounds.y);
    maxX = Math.max(maxX, bounds.x + bounds.cx);
    maxY = Math.max(maxY, bounds.y + bounds.cy);
  }
  if (!Number.isFinite(minX) || !Number.isFinite(minY) || !Number.isFinite(maxX) || !Number.isFinite(maxY)) {
    logger.WARN(ERPPGB001, 'branch missing bounds', { blocks: blocks.length });
    return undefined;
  }
  return { x: minX, y: minY, cx: maxX - minX, cy: maxY - minY };
}

function objectBounds(xml: string): XmlBounds | undefined {
  const off = /<a:off\b[^>]*\bx="(-?\d+)"[^>]*\by="(-?\d+)"/.exec(xml);
  const ext = /<a:ext\b[^>]*\bcx="(\d+)"[^>]*\bcy="(\d+)"/.exec(xml);
  if (!off || !ext || !off[1] || !off[2] || !ext[1] || !ext[2]) return undefined;
  return { x: Number(off[1]), y: Number(off[2]), cx: Number(ext[1]), cy: Number(ext[2]) };
}

function groupShapeXML(groupId: string, children: string, bounds: XmlBounds, id: number): string {
  const name = xmlAttr(`xaligo anchor ${groupId}`);
  return `<p:grpSp><p:nvGrpSpPr><p:cNvPr id="${id}" name="${name}"/><p:cNvGrpSpPr/><p:nvPr/></p:nvGrpSpPr><p:grpSpPr><a:xfrm><a:off x="${bounds.x}" y="${bounds.y}"/><a:ext cx="${bounds.cx}" cy="${bounds.cy}"/><a:chOff x="${bounds.x}" y="${bounds.y}"/><a:chExt cx="${bounds.cx}" cy="${bounds.cy}"/></a:xfrm></p:grpSpPr>${children}</p:grpSp>`;
}

function nextNvPrID(xml: string): number {
  let max = 1;
  const re = /<p:cNvPr\b[^>]*\bid="(\d+)"/g;
  let match: RegExpExecArray | null;
  while ((match = re.exec(xml))) {
    const id = match[1];
    if (id) max = Math.max(max, Number(id));
  }
  return max + 1;
}

function xmlAttr(value: string): string {
  return value
    .replace(/&/g, '&amp;')
    .replace(/"/g, '&quot;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
}

function toArrayBuffer(bytes: Uint8Array): ArrayBuffer {
  const out = new ArrayBuffer(bytes.byteLength);
  new Uint8Array(out).set(bytes);
  return out;
}

function bytesToBase64(bytes: Uint8Array): string {
  if (typeof Buffer !== 'undefined') return Buffer.from(bytes).toString('base64');
  let binary = '';
  for (const byte of bytes) binary += String.fromCharCode(byte);
  return btoa(binary);
}
