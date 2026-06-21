#!/usr/bin/env node

import { runPptxCli } from './controller/cli';

runPptxCli(process.argv.slice(2)).catch((err: unknown) => {
  console.error(err instanceof Error ? err.message : String(err));
  process.exitCode = 1;
});
