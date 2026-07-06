#!/usr/bin/env node
'use strict';

const fs = require('node:fs');
const https = require('node:https');
const path = require('node:path');

const packageRoot = path.resolve(__dirname, '..', '..');
const packageJson = require(path.join(packageRoot, 'package.json'));

const supportedTargets = new Set([
  'darwin/amd64',
  'darwin/arm64',
  'linux/amd64',
  'linux/arm64',
  'win32/amd64',
  'win32/arm64',
]);

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

function target() {
  return `${process.platform}/${goArch()}`;
}

function binaryPlatform() {
  return process.platform === 'win32' ? 'windows' : process.platform;
}

function isSourceCheckout() {
  return fs.existsSync(path.join(packageRoot, '.git'));
}

function binaryName() {
  const arch = goArch();
  const suffix = process.platform === 'win32' ? '.exe' : '';
  return `xaligo-${binaryPlatform()}-${arch}${suffix}`;
}

function releaseTag() {
  if (process.env.XALIGO_NPM_RELEASE_TAG) return process.env.XALIGO_NPM_RELEASE_TAG;
  if (packageJson.xaligo?.releaseTag) return packageJson.xaligo.releaseTag;
  return `v${String(packageJson.version).split('+')[0]}`;
}

function download(url, destination, redirects = 0) {
  return new Promise((resolve, reject) => {
    https.get(url, (response) => {
      const status = response.statusCode ?? 0;
      if ([301, 302, 303, 307, 308].includes(status) && response.headers.location) {
        response.resume();
        if (redirects >= 5) {
          reject(new Error(`too many redirects while downloading ${url}`));
          return;
        }
        resolve(download(new URL(response.headers.location, url).toString(), destination, redirects + 1));
        return;
      }
      if (status < 200 || status >= 300) {
        response.resume();
        reject(new Error(`download failed with HTTP ${status}: ${url}`));
        return;
      }
      const file = fs.createWriteStream(destination, { mode: 0o755 });
      response.pipe(file);
      file.on('finish', () => {
        file.close(() => resolve());
      });
      file.on('error', reject);
    }).on('error', reject);
  });
}

async function main() {
  if (!supportedTargets.has(target())) {
    throw new Error(`unsupported platform/architecture: ${process.platform}/${process.arch}`);
  }

  const name = binaryName();
  const nativeDir = path.join(packageRoot, 'bin', 'native');
  const destination = path.join(nativeDir, name);

  if (fs.existsSync(destination)) return;
  if (process.env.XALIGO_NPM_SKIP_DOWNLOAD) return;
  if (isSourceCheckout()) return;

  fs.mkdirSync(nativeDir, { recursive: true });
  const url = `https://github.com/xaligo/xaligo/releases/download/${releaseTag()}/${name}`;
  await download(url, destination);
  fs.chmodSync(destination, 0o755);
}

main().catch((err) => {
  console.error(`xaligo install failed: ${err instanceof Error ? err.message : String(err)}`);
  console.error('Set XALIGO_NPM_SKIP_DOWNLOAD=1 to skip native binary download.');
  process.exit(1);
});
