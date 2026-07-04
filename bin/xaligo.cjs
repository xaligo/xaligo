#!/usr/bin/env node
'use strict';

const { spawn } = require('node:child_process');
const fs = require('node:fs');
const path = require('node:path');

const packageRoot = path.resolve(__dirname, '..');

function goArch() {
  switch (process.arch) {
    case 'x64':
      return 'amd64';
    case 'arm64':
      return 'arm64';
    default:
      return process.arch;
  }
}

const executable = process.platform === 'win32'
  ? `xaligo-${process.platform}-${goArch()}.exe`
  : `xaligo-${process.platform}-${goArch()}`;
const binary = path.join(__dirname, 'native', executable);

if (!fs.existsSync(binary)) {
  console.error(`xaligo native binary is not installed for ${process.platform}/${process.arch}.`);
  console.error('Reinstall the package, or run npm run build:npm-binaries from a source checkout.');
  process.exit(1);
}

const child = spawn(binary, process.argv.slice(2), {
  stdio: 'inherit',
  env: {
    ...process.env,
    XALIGO_HOME: process.env.XALIGO_HOME || packageRoot,
  },
});

child.on('error', (err) => {
  console.error(err.message);
  process.exit(1);
});

child.on('exit', (code, signal) => {
  if (signal) {
    process.kill(process.pid, signal);
    return;
  }
  process.exit(code ?? 0);
});
