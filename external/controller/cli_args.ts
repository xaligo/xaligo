import type { CliArgs } from '../entity/cli';
import type { ArrowStyle, PaperOrientation, PaperSize } from '../entity/pptx';

const ARROW_STYLES: readonly ArrowStyle[] = [
  'thin', 'standard', 'triangle', 'stealth', 'arrow', 'diamond', 'oval', 'none',
];

const PAPER_SIZES: readonly PaperSize[] = [
  'A5', 'A4', 'A3', 'A2', 'A1', 'Letter', 'Legal', 'Tabloid',
];

export function parseArgs(argv: string[]): CliArgs {
  const out: CliArgs = {};
  for (let i = 0; i < argv.length; i += 1) {
    const arg = argv[i];
    switch (arg) {
      case '--xal':
        out.xal = nextValue(argv, i, arg);
        i += 1;
        break;
      case '--services':
        out.services = nextValue(argv, i, arg);
        i += 1;
        break;
      case '-o':
      case '--output':
        out.output = nextValue(argv, i, arg);
        i += 1;
        break;
      case '--title':
        out.title = nextValue(argv, i, arg);
        i += 1;
        break;
      case '--author':
        out.author = nextValue(argv, i, arg);
        i += 1;
        break;
      case '--company':
        out.company = nextValue(argv, i, arg);
        i += 1;
        break;
      case '--subject':
        out.subject = nextValue(argv, i, arg);
        i += 1;
        break;
      case '--compression':
        out.compression = parseBoolean(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--no-compression':
        out.compression = false;
        break;
      case '--px-per-inch':
        out.pxPerInch = parsePositiveNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--arrow-style':
        out.arrowStyle = parseArrowStyle(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--arrow-stub':
        out.arrowStub = parsePositiveNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--arrow-margin':
        out.arrowMargin = parsePositiveNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--paper':
        out.paper = parsePaper(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--orientation':
        out.orientation = parseOrientation(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--paper-margin':
        out.paperMargin = parseNonNegativeNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--paper-margin-top':
        out.paperMarginTop = parseNonNegativeNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--paper-margin-right':
        out.paperMarginRight = parseNonNegativeNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--paper-margin-bottom':
        out.paperMarginBottom = parseNonNegativeNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '--paper-margin-left':
        out.paperMarginLeft = parseNonNegativeNumber(nextValue(argv, i, arg), arg);
        i += 1;
        break;
      case '-h':
      case '--help':
        out.help = true;
        break;
      default:
        throw new Error(`unknown option: ${arg}`);
    }
  }
  return out;
}

export function printHelp(): void {
  console.log(`Usage:
  xaligo-pptx --xal <file.xal> -o <out.pptx> [options]

Options:
  --services <services.csv>
  --title <title>
  --author <author>
  --company <company>
  --subject <subject>
  --compression true|false
  --no-compression
  --px-per-inch <number>
  --arrow-style <thin|standard|triangle|stealth|arrow|diamond|oval|none>  (default: thin)
  --arrow-stub <number>   stub length in px before the first/last bend (default: 20)
  --arrow-margin <number> clear margin in px reserved on both sides of each line (default: 8)
  --paper <A5|A4|A3|A2|A1|Letter|Legal|Tabloid>  size the slide to this paper and fit the diagram
  --orientation <portrait|landscape>             paper orientation (default: auto-fit)
  --paper-margin <number>                        paper margin in inches on all sides (default: 0)
  --paper-margin-top <number>                    paper top margin in inches
  --paper-margin-right <number>                  paper right margin in inches
  --paper-margin-bottom <number>                 paper bottom margin in inches
  --paper-margin-left <number>                   paper left margin in inches
`);
}

function nextValue(argv: string[], index: number, option: string): string {
  const value = argv[index + 1];
  if (!value || value.startsWith('-')) {
    throw new Error(`${option} requires a value`);
  }
  return value;
}

function parseBoolean(value: string, option: string): boolean {
  if (value === 'true') return true;
  if (value === 'false') return false;
  throw new Error(`${option} must be true or false`);
}

function parsePositiveNumber(value: string, option: string): number {
  const parsed = Number(value);
  if (!Number.isFinite(parsed) || parsed <= 0) {
    throw new Error(`${option} must be a positive number`);
  }
  return parsed;
}

function parseNonNegativeNumber(value: string, option: string): number {
  const parsed = Number(value);
  if (!Number.isFinite(parsed) || parsed < 0) {
    throw new Error(`${option} must be a non-negative number`);
  }
  return parsed;
}

function parseArrowStyle(value: string, option: string): ArrowStyle {
  if ((ARROW_STYLES as readonly string[]).includes(value)) {
    return value as ArrowStyle;
  }
  throw new Error(`${option} must be one of: ${ARROW_STYLES.join(', ')}`);
}

function parsePaper(value: string, option: string): PaperSize {
  if ((PAPER_SIZES as readonly string[]).includes(value)) {
    return value as PaperSize;
  }
  throw new Error(`${option} must be one of: ${PAPER_SIZES.join(', ')}`);
}

function parseOrientation(value: string, option: string): PaperOrientation {
  if (value === 'portrait' || value === 'landscape') {
    return value;
  }
  throw new Error(`${option} must be portrait or landscape`);
}
