'use strict';

const assert = require('node:assert/strict');
const crypto = require('node:crypto');
const fs = require('node:fs');
const os = require('node:os');
const path = require('node:path');
const test = require('node:test');

const { installBinary } = require('../../../../scripts/npm/install.cjs');

function checksum(contents) {
  return crypto.createHash('sha256').update(contents).digest('hex');
}

function fixture(t) {
  const directory = fs.mkdtempSync(path.join(os.tmpdir(), 'xaligo-install-test-'));
  t.after(() => fs.rmSync(directory, { recursive: true, force: true }));

  const name = 'xaligo-linux-amd64';
  return {
    binaryURL: `https://example.invalid/${name}`,
    checksumURL: `https://example.invalid/${name}.sha256`,
    destination: path.join(directory, name),
    directory,
    name,
  };
}

function fakeDownloader(files, calls) {
  return async (url, destination) => {
    calls.push(url);
    const contents = files.get(url);
    if (contents === undefined) throw new Error(`unexpected download: ${url}`);
    fs.writeFileSync(destination, contents, { mode: 0o600 });
  };
}

test('keeps an existing binary only when its published checksum matches', async (t) => {
  const paths = fixture(t);
  const binary = Buffer.from('verified binary');
  const calls = [];
  fs.writeFileSync(paths.destination, binary, { mode: 0o600 });

  const result = await installBinary({
    ...paths,
    downloadFile: fakeDownloader(new Map([
      [paths.checksumURL, `${checksum(binary)}  ${paths.name}\n`],
    ]), calls),
  });

  assert.equal(result.status, 'current');
  assert.deepEqual(calls, [paths.checksumURL]);
  assert.deepEqual(fs.readFileSync(paths.destination), binary);
  assert.equal(fs.statSync(paths.destination).mode & 0o777, 0o755);
});

test('replaces a mismatched existing binary with a verified download', async (t) => {
  const paths = fixture(t);
  const binary = Buffer.from('verified replacement');
  const calls = [];
  fs.writeFileSync(paths.destination, 'partial download');

  const result = await installBinary({
    ...paths,
    downloadFile: fakeDownloader(new Map([
      [paths.checksumURL, `${checksum(binary)}  ${paths.name}\n`],
      [paths.binaryURL, binary],
    ]), calls),
    renameFile(source, destination) {
      if (fs.existsSync(destination)) {
        const error = new Error(`destination exists: ${destination}`);
        error.code = 'EEXIST';
        throw error;
      }
      fs.renameSync(source, destination);
    },
  });

  assert.equal(result.status, 'installed');
  assert.deepEqual(calls, [paths.checksumURL, paths.binaryURL]);
  assert.deepEqual(fs.readFileSync(paths.destination), binary);
  assert.equal(fs.statSync(paths.destination).mode & 0o777, 0o755);
});

test('preserves the existing binary and removes temporary files after verification fails', async (t) => {
  const paths = fixture(t);
  const existing = Buffer.from('existing binary');
  const expected = Buffer.from('expected replacement');
  fs.writeFileSync(paths.destination, existing);

  await assert.rejects(
    installBinary({
      ...paths,
      downloadFile: fakeDownloader(new Map([
        [paths.checksumURL, `${checksum(expected)}  ${paths.name}\n`],
        [paths.binaryURL, Buffer.from('corrupted replacement')],
      ]), []),
    }),
    /checksum mismatch/,
  );

  assert.deepEqual(fs.readFileSync(paths.destination), existing);
  assert.deepEqual(fs.readdirSync(paths.directory), [paths.name]);
});

test('restores the existing binary when verified replacement cannot be installed', async (t) => {
  const paths = fixture(t);
  const existing = Buffer.from('existing binary');
  const replacement = Buffer.from('verified replacement');
  fs.writeFileSync(paths.destination, existing);
  let renameCalls = 0;

  await assert.rejects(
    installBinary({
      ...paths,
      downloadFile: fakeDownloader(new Map([
        [paths.checksumURL, `${checksum(replacement)}  ${paths.name}\n`],
        [paths.binaryURL, replacement],
      ]), []),
      renameFile(source, destination) {
        renameCalls += 1;
        if (renameCalls === 2) {
          const error = new Error('simulated install rename failure');
          error.code = 'EPERM';
          throw error;
        }
        fs.renameSync(source, destination);
      },
    }),
    /simulated install rename failure/,
  );

  assert.equal(renameCalls, 3);
  assert.deepEqual(fs.readFileSync(paths.destination), existing);
  assert.deepEqual(fs.readdirSync(paths.directory), [paths.name]);
});
