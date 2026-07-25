#!/usr/bin/env node
'use strict';

const crypto = require('node:crypto');
const fs = require('node:fs');
const https = require('node:https');
const path = require('node:path');
const { pipeline } = require('node:stream/promises');

const packageRoot = path.resolve(__dirname, '..', '..');
const packageJson = require(path.join(packageRoot, 'package.json'));
const downloadTimeoutMilliseconds = 30_000;
const maximumChecksumBytes = 4_096;

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

function openDownload(url, redirects = 0) {
  return new Promise((resolve, reject) => {
    const parsedURL = new URL(url);
    if (parsedURL.protocol !== 'https:') {
      reject(new Error(`refusing non-HTTPS download: ${url}`));
      return;
    }

    const request = https.get(parsedURL, (response) => {
      const status = response.statusCode ?? 0;
      if ([301, 302, 303, 307, 308].includes(status) && response.headers.location) {
        response.resume();
        if (redirects >= 5) {
          reject(new Error(`too many redirects while downloading ${url}`));
          return;
        }
        resolve(openDownload(new URL(response.headers.location, parsedURL).toString(), redirects + 1));
        return;
      }
      if (status < 200 || status >= 300) {
        response.resume();
        reject(new Error(`download failed with HTTP ${status}: ${url}`));
        return;
      }
      resolve(response);
    });
    request.setTimeout(downloadTimeoutMilliseconds, () => {
      request.destroy(new Error(`download timed out: ${url}`));
    });
    request.on('error', reject);
  });
}

async function download(url, destination, options = {}) {
  const response = await openDownload(url);
  const file = fs.createWriteStream(destination, {
    flags: 'wx',
    mode: options.mode ?? 0o600,
  });

  try {
    await pipeline(response, file);
  } catch (err) {
    safeUnlink(destination);
    throw err;
  }
}

function safeUnlink(filePath) {
  try {
    fs.unlinkSync(filePath);
  } catch (err) {
    if (!(err && err.code === 'ENOENT')) throw err;
  }
}

function sha256File(filePath) {
  const hash = crypto.createHash('sha256');
  const buffer = Buffer.allocUnsafe(64 * 1024);
  const file = fs.openSync(filePath, 'r');
  try {
    let bytesRead;
    do {
      bytesRead = fs.readSync(file, buffer, 0, buffer.length, null);
      if (bytesRead > 0) hash.update(buffer.subarray(0, bytesRead));
    } while (bytesRead > 0);
  } finally {
    fs.closeSync(file);
  }
  return hash.digest('hex');
}

function parseChecksum(contents, expectedName) {
  for (const line of contents.split(/\r?\n/)) {
    const match = line.match(/^([0-9a-f]{64})[ \t]+\*?(.+?)\s*$/i);
    if (match && match[2] === expectedName) return match[1].toLowerCase();
  }
  throw new Error(`invalid checksum asset for ${expectedName}`);
}

function checksumsEqual(left, right) {
  if (!/^[0-9a-f]{64}$/i.test(left) || !/^[0-9a-f]{64}$/i.test(right)) return false;
  return crypto.timingSafeEqual(Buffer.from(left, 'hex'), Buffer.from(right, 'hex'));
}

async function installBinary(options) {
  const {
    name,
    destination,
    binaryURL,
    checksumURL,
    downloadFile = download,
    renameFile = fs.renameSync,
  } = options;
  const nativeDir = path.dirname(destination);

  fs.mkdirSync(nativeDir, { recursive: true });
  const temporaryDir = fs.mkdtempSync(path.join(nativeDir, `.${name}.${process.pid}-`));
  const binaryTemporaryPath = path.join(temporaryDir, name);
  const checksumTemporaryPath = path.join(temporaryDir, `${name}.sha256`);
  try {
    await downloadFile(checksumURL, checksumTemporaryPath, { mode: 0o600 });
    if (fs.statSync(checksumTemporaryPath).size > maximumChecksumBytes) {
      throw new Error(`checksum asset is too large for ${name}`);
    }
    const expectedChecksum = parseChecksum(
      fs.readFileSync(checksumTemporaryPath, 'utf8'),
      name,
    );

    if (fs.existsSync(destination)) {
      const installedStat = fs.lstatSync(destination);
      if (!installedStat.isFile()) {
        throw new Error(`refusing to replace non-file native binary path: ${destination}`);
      }
      try {
        const installedChecksum = sha256File(destination);
        if (checksumsEqual(installedChecksum, expectedChecksum)) {
          fs.chmodSync(destination, 0o755);
          return { status: 'current', checksum: expectedChecksum };
        }
      } catch {
        // An unreadable existing artifact cannot be trusted; replace it below.
      }
    }

    await downloadFile(binaryURL, binaryTemporaryPath, { mode: 0o600 });
    const downloadedChecksum = sha256File(binaryTemporaryPath);
    if (!checksumsEqual(downloadedChecksum, expectedChecksum)) {
      throw new Error(
        `checksum mismatch for ${name}: expected ${expectedChecksum}, got ${downloadedChecksum}`,
      );
    }

    fs.chmodSync(binaryTemporaryPath, 0o755);
    let backupPath = '';
    if (fs.existsSync(destination)) {
      backupPath = path.join(
        nativeDir,
        `.${name}.backup-${process.pid}-${crypto.randomUUID()}`,
      );
      renameFile(destination, backupPath);
    }
    try {
      renameFile(binaryTemporaryPath, destination);
    } catch (installError) {
      if (backupPath) {
        try {
          renameFile(backupPath, destination);
          backupPath = '';
        } catch (restoreError) {
          throw new AggregateError(
            [installError, restoreError],
            `install failed and the previous binary could not be restored; recover it from ${backupPath}`,
          );
        }
      }
      throw installError;
    }
    if (backupPath) safeUnlink(backupPath);
    return { status: 'installed', checksum: expectedChecksum };
  } finally {
    fs.rmSync(temporaryDir, { recursive: true, force: true });
  }
}

async function main() {
  if (!supportedTargets.has(target())) {
    throw new Error(`unsupported platform/architecture: ${process.platform}/${process.arch}`);
  }

  const name = binaryName();
  const nativeDir = path.join(packageRoot, 'bin', 'native');
  const destination = path.join(nativeDir, name);

  if (process.env.XALIGO_NPM_SKIP_DOWNLOAD) return;
  if (isSourceCheckout()) return;

  const releaseURL = `https://github.com/xaligo/xaligo/releases/download/${releaseTag()}`;
  await installBinary({
    name,
    destination,
    binaryURL: `${releaseURL}/${name}`,
    checksumURL: `${releaseURL}/${name}.sha256`,
  });
}

if (require.main === module) {
  main().catch((err) => {
    console.error(`xaligo install failed: ${err instanceof Error ? err.message : String(err)}`);
    console.error('Set XALIGO_NPM_SKIP_DOWNLOAD=1 to skip native binary download.');
    process.exit(1);
  });
}

module.exports = {
  installBinary,
};
