import { Buffer } from 'node:buffer';
import { readFile, writeFile } from 'fs/promises';
import process from 'node:process';

import { loadXaligo, type PptxExportOptions } from '../usecase/api';
import type { CliArgs } from '../entity/cli';
import { parseArgs, printHelp } from './cli_args';

export async function runPptxCli(argv: string[]): Promise<void> {
  const args = parseArgs(argv);
  if (args.help) {
    printHelp();
    return;
  }
  if (!args.xal || !args.output) {
    printHelp();
    process.exitCode = 1;
    return;
  }

  const xal = await readFile(args.xal, 'utf8');
  const services = args.services ? await readFile(args.services, 'utf8') : undefined;
  const xaligo = await loadXaligo();
  const options = pptxOptions(args);
  const pptx = services
    ? await xaligo.renderWithServicesPptx(xal, services, options)
    : await xaligo.renderPptx(xal, options);

  await writeFile(args.output, Buffer.from(pptx as Uint8Array));
  console.log(`generated: ${args.output}`);
}

function pptxOptions(args: CliArgs): PptxExportOptions {
  const options: PptxExportOptions = {
    outputType: 'uint8array',
  };
  const title = args.title ?? args.output;
  if (title !== undefined) options.title = title;
  if (args.author !== undefined) options.author = args.author;
  if (args.company !== undefined) options.company = args.company;
  if (args.subject !== undefined) options.subject = args.subject;
  if (args.compression !== undefined) options.compression = args.compression;
  if (args.pxPerInch !== undefined) options.pxPerInch = args.pxPerInch;
  if (args.arrowStyle !== undefined) options.arrowStyle = args.arrowStyle;
  if (args.arrowStub !== undefined) options.arrowStubPx = args.arrowStub;
  if (args.arrowMargin !== undefined) options.arrowMarginPx = args.arrowMargin;
  if (args.paper !== undefined) options.paperSize = args.paper;
  if (args.orientation !== undefined) options.orientation = args.orientation;
  if (args.paperMargin !== undefined) options.paperMargin = args.paperMargin;
  if (args.paperMarginTop !== undefined) options.paperMarginTop = args.paperMarginTop;
  if (args.paperMarginRight !== undefined) options.paperMarginRight = args.paperMarginRight;
  if (args.paperMarginBottom !== undefined) options.paperMarginBottom = args.paperMarginBottom;
  if (args.paperMarginLeft !== undefined) options.paperMarginLeft = args.paperMarginLeft;
  return options;
}
