import JSZip from 'jszip';

import {
  ANCHOR_GROUP_MARKER,
  FRONT_LAYER_MARKER,
  type PlanOp,
  type PptxExportResult,
  type PptxOutputType,
} from '../entity/pptx';
import { NewEnvLogger } from '../share/logger';
import { NewMCode } from '../share/mcode';

const logger = NewEnvLogger('external/repository', 'pptx_package');
const ERPPGAOIP001 = NewMCode('ERPPGAOIP-001', 'Group anchor objects in PPTX no groups branch');
const ERPPGAOIP002 = NewMCode('ERPPGAOIP-002', 'Group anchor objects in PPTX missing slide branch');
const ERPPGAOIP003 = NewMCode('ERPPGAOIP-003', 'Group anchor objects in PPTX completed');
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

export async function groupAnchorObjectsInPptx(bytes: Uint8Array, ops: PlanOp[], compression: boolean): Promise<Uint8Array> {
  const groupIds = [...new Set(ops.map((op) => op.groupId).filter((id): id is string => !!id))];
  if (groupIds.length === 0) {
    logger.DEBUG(ERPPGAOIP001, 'branch no groups');
    return bytes;
  }

  const zip = await JSZip.loadAsync(bytes);
  const slidePath = 'ppt/slides/slide1.xml';
  const slide = zip.file(slidePath);
  if (!slide) {
    logger.WARN(ERPPGAOIP002, 'branch missing slide', { slidePath });
    return bytes;
  }

  let xml = await slide.async('string');
  xml = applySlenderStealthArrowheads(xml);
  for (const groupId of groupIds.sort()) {
    xml = groupSlideObjects(xml, groupId);
  }
  xml = moveAnchorAndLineObjectsToFront(xml);
  zip.file(slidePath, xml);
  const out = await zip.generateAsync({ type: 'uint8array', compression: compression ? 'DEFLATE' : 'STORE' });
  logger.DEBUG(ERPPGAOIP003, 'completed', { groups: groupIds.length, bytes: out.length });
  return out;
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
  const name = /<p:cNvPr\b[^>]*\bname="([^"]*)"/.exec(xml)?.[1];
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
  const movingXML = movingIndexes.map((index) => blocks[index]?.xml ?? '').join('');

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
  return `${out.slice(0, spTreeClose)}${movingXML}${out.slice(spTreeClose)}`;
}

function isAnchorGroupBlock(xml: string): boolean {
  return /<p:grpSp\b/.test(xml) && /<p:cNvPr\b[^>]*\bname="xaligo anchor xaligo-anchor-/.test(xml);
}

function isFrontLayerBlock(xml: string): boolean {
  const name = /<p:cNvPr\b[^>]*\bname="([^"]*)"/.exec(xml)?.[1];
  return !!name?.startsWith(FRONT_LAYER_MARKER);
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
