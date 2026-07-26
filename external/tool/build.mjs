import { spawnSync } from 'node:child_process';
import {
  copyFileSync,
  mkdirSync,
  mkdtempSync,
  readFileSync,
  readdirSync,
  rmSync,
  statSync,
  writeFileSync,
} from 'node:fs';
import { createRequire } from 'node:module';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const toolDirectory = path.dirname(fileURLToPath(import.meta.url));
const externalRoot = path.resolve(toolDirectory, '..');
const tsconfigPath = path.join(externalRoot, 'tsconfig.json');
const distDirectory = path.join(externalRoot, 'dist');
const require = createRequire(import.meta.url);
const typescriptPackagePath = require.resolve('typescript/package.json');
const tscPath = path.join(path.dirname(typescriptPackagePath), 'lib', 'tsc.js');
const supportedExternalModules = new Set(['jszip', 'pptxgenjs', 'node:fs']);
const sourceMapReferencePattern = /^\s*\/\/[#@]\s*sourceMappingURL=.*$/gm;
const staticRequirePattern = /\brequire\(\s*(['"])([^'"]+)\1\s*\)/g;

const argumentsSet = new Set(process.argv.slice(2));
for (const argument of argumentsSet) {
  if (argument !== '--watch') throw new Error(`unsupported build argument: ${argument}`);
}

if (argumentsSet.has('--watch')) {
  try {
    build();
  } catch (error) {
    console.error(error);
  }
  watch();
} else {
  build();
}

function build() {
  const startedAt = Date.now();
  const sourcePaths = findTypeScriptSources(externalRoot);
  rmSync(distDirectory, { recursive: true, force: true });
  runTypeScript(['--project', tsconfigPath, '--emitDeclarationOnly']);

  const workDirectory = mkdtempSync(path.join(externalRoot, '.xaligo-build-'));
  try {
    const sourceDirectory = path.join(workDirectory, 'source');
    const commonJSDirectory = path.join(workDirectory, 'commonjs');
    copyCommonJSSources(sourcePaths, sourceDirectory);
    writeCommonJSTSConfig(sourceDirectory, commonJSDirectory);
    runTypeScript(['--project', path.join(sourceDirectory, 'tsconfig.json')]);

    const vendorRuntime = createVendorRuntime();
    const indexBundle = createBundle(
      commonJSDirectory,
      sourceDirectory,
      'index.js',
      'index.js',
      vendorRuntime,
      true,
    );
    const commandBundle = createBundle(
      commonJSDirectory,
      sourceDirectory,
      'command.js',
      'cli.js',
      vendorRuntime,
      false,
    );
    writeBundle(indexBundle);
    writeBundle(commandBundle);
  } finally {
    rmSync(workDirectory, { recursive: true, force: true });
  }
  console.log(`[xaligo-build] built TypeScript bundles in ${Date.now() - startedAt}ms`);
}

function findTypeScriptSources(directory) {
  const sources = [];
  for (const entry of readdirSync(directory, { withFileTypes: true })) {
    if (entry.isDirectory()) {
      if (ignoredBuildDirectory(entry.name)) continue;
      sources.push(...findTypeScriptSources(path.join(directory, entry.name)));
      continue;
    }
    if (entry.isFile() && entry.name.endsWith('.ts')) {
      sources.push(path.join(directory, entry.name));
    }
  }
  return sources.sort();
}

function ignoredBuildDirectory(name) {
  return name === 'dist'
    || name === 'node_modules'
    || name === 'wasm'
    || name.startsWith('.xaligo-build-');
}

function copyCommonJSSources(sourcePaths, destinationRoot) {
  for (const sourcePath of sourcePaths) {
    const relativePath = path.relative(externalRoot, sourcePath);
    const destinationPath = path.join(destinationRoot, relativePath);
    mkdirSync(path.dirname(destinationPath), { recursive: true });
    copyFileSync(sourcePath, destinationPath);
  }
  writeFileSync(
    path.join(destinationRoot, 'package.json'),
    `${JSON.stringify({ type: 'commonjs' }, null, 2)}\n`,
  );
}

function writeCommonJSTSConfig(sourceDirectory, commonJSDirectory) {
  const sourceConfig = JSON.parse(readFileSync(tsconfigPath, 'utf8'));
  const config = {
    ...sourceConfig,
    compilerOptions: {
      ...sourceConfig.compilerOptions,
      module: 'Node16',
      moduleResolution: 'Node16',
      declaration: false,
      declarationMap: false,
      inlineSourceMap: false,
      inlineSources: true,
      sourceMap: true,
      outDir: commonJSDirectory,
      rootDir: '.',
    },
    include: ['**/*.ts'],
    exclude: ['node_modules', 'dist'],
  };
  writeFileSync(
    path.join(sourceDirectory, 'tsconfig.json'),
    `${JSON.stringify(config, null, 2)}\n`,
  );
}

function runTypeScript(argumentsList) {
  const result = spawnSync(process.execPath, [tscPath, ...argumentsList], {
    cwd: externalRoot,
    stdio: 'inherit',
  });
  if (result.error) throw result.error;
  if (result.status !== 0) {
    throw new Error(`TypeScript compiler failed with status ${result.status ?? 'unknown'}`);
  }
}

function createVendorRuntime() {
  const jszipEntry = require.resolve('jszip');
  const pptxgenEntry = require.resolve('pptxgenjs');
  const jszipPath = path.resolve(path.dirname(jszipEntry), '..', 'dist', 'jszip.min.js');
  const pptxgenPath = path.join(path.dirname(pptxgenEntry), 'pptxgen.min.js');
  const jszipSource = stripSourceMapReference(readFileSync(jszipPath, 'utf8'));
  const pptxgenSource = stripSourceMapReference(readFileSync(pptxgenPath, 'utf8'));

  return [
    {
      code: `// Generated by external/tool/build.mjs. Do not edit.
const __xaligoImmediateTasks = new Set();
let __xaligoImmediateSequence = 0;
if (typeof globalThis.setImmediate !== 'function') {
  globalThis.setImmediate = (callback, ...args) => {
    const id = ++__xaligoImmediateSequence;
    __xaligoImmediateTasks.add(id);
    const schedule = typeof globalThis.queueMicrotask === 'function'
      ? globalThis.queueMicrotask
      : (task) => void Promise.resolve().then(task);
    schedule(() => {
      if (!__xaligoImmediateTasks.delete(id)) return;
      callback(...args);
    });
    return id;
  };
}
if (typeof globalThis.clearImmediate !== 'function') {
  globalThis.clearImmediate = (id) => {
    __xaligoImmediateTasks.delete(id);
  };
}

const JSZip = (() => {
  const vendorModule = { exports: {} };
  (function (module, exports, define, window, global, self) {
`,
    },
    {
      code: jszipSource,
      sourceMap: createIdentitySourceMap('vendor/jszip.min.js', jszipSource),
    },
    {
      code: `
  }).call(
    globalThis,
    vendorModule,
    vendorModule.exports,
    undefined,
    undefined,
    globalThis,
    globalThis,
  );
  if (typeof vendorModule.exports !== 'function') throw new Error('failed to initialize JSZip');
  return vendorModule.exports;
})();

const PptxGenJS = (() => {
  let vendorExport;
  (function (module, exports, define, window, global, self) {
`,
    },
    {
      code: pptxgenSource,
      sourceMap: createIdentitySourceMap('vendor/pptxgenjs.min.js', pptxgenSource),
    },
    {
      code: `
    vendorExport = PptxGenJS;
  }).call(globalThis, undefined, undefined, undefined, undefined, globalThis, globalThis);
  if (typeof vendorExport !== 'function') throw new Error('failed to initialize PptxGenJS');
  return vendorExport;
})();
`,
    },
  ];
}

function stripSourceMapReference(source) {
  return source.replace(sourceMapReferencePattern, '').trimEnd();
}

function createBundle(
  commonJSDirectory,
  sourceDirectory,
  entryID,
  outputFile,
  vendorRuntime,
  exportEntry,
) {
  const { modules, resolutions } = collectModules(commonJSDirectory, sourceDirectory, entryID);
  const entrySource = readFileSync(path.join(commonJSDirectory, entryID), 'utf8');
  const footer = exportEntry
    ? createExportFooter(entryID, exportedNames(entrySource))
    : `__xaligoRequire(${JSON.stringify(entryID)});\n`;
  const writer = createBundleWriter(outputFile);

  for (const chunk of vendorRuntime) {
    writer.append(chunk.code, chunk.sourceMap);
  }
  writer.append(`
const __xaligoFactories = Object.create(null);
`);
  for (const [moduleID, module] of [...modules.entries()]
    .sort(([left], [right]) => left.localeCompare(right))) {
    writer.append(
      `__xaligoFactories[${JSON.stringify(moduleID)}] = function (module, exports, require) {\n`,
    );
    writer.append(module.source, module.sourceMap);
    writer.append('\n};\n');
  }
  writer.append(`const __xaligoResolutions = Object.assign(
  Object.create(null),
  ${JSON.stringify(Object.fromEntries([...resolutions.entries()].sort()))},
);
const __xaligoModuleCache = Object.create(null);

function __xaligoRequire(request, fromModuleID = '') {
  let moduleID = request;
  if (request.startsWith('.')) {
    moduleID = __xaligoResolutions[\`\${fromModuleID}\\0\${request}\`];
    if (!moduleID) {
      throw new Error(\`bundled module resolution not found: \${request} from \${fromModuleID}\`);
    }
  }
  if (moduleID === 'jszip') return JSZip;
  if (moduleID === 'pptxgenjs') return PptxGenJS;
  if (moduleID.startsWith('node:')) {
    if (typeof globalThis.process?.getBuiltinModule === 'function') {
      return globalThis.process.getBuiltinModule(moduleID);
    }
    throw new Error(\`Dynamic require of "\${moduleID}" is not supported\`);
  }
  if (__xaligoModuleCache[moduleID]) return __xaligoModuleCache[moduleID].exports;
  const factory = __xaligoFactories[moduleID];
  if (!factory) throw new Error(\`bundled module not found: \${moduleID}\`);
  const module = { exports: {} };
  __xaligoModuleCache[moduleID] = module;
  factory(module, module.exports, (nextRequest) => __xaligoRequire(nextRequest, moduleID));
  return module.exports;
}

${footer}
//# sourceMappingURL=${outputFile}.map
`);
  return writer.finish();
}

function collectModules(commonJSDirectory, sourceDirectory, entryID) {
  const modules = new Map();
  const resolutions = new Map();
  visit(entryID);
  return { modules, resolutions };

  function visit(moduleID) {
    if (modules.has(moduleID)) return;
    const modulePath = path.join(commonJSDirectory, ...moduleID.split('/'));
    const dependencies = [];
    const source = stripSourceMapReference(readFileSync(modulePath, 'utf8'));
    for (const match of source.matchAll(staticRequirePattern)) {
      const request = match[2];
      if (request === undefined) throw new Error(`invalid require in ${moduleID}`);
      if (request.startsWith('.')) {
        const resolved = resolveCommonJSModule(commonJSDirectory, moduleID, request);
        dependencies.push(resolved);
        resolutions.set(`${moduleID}\0${request}`, resolved);
        continue;
      }
      if (!supportedExternalModules.has(request)) {
        throw new Error(`unsupported external module ${JSON.stringify(request)} in ${moduleID}`);
      }
    }
    modules.set(moduleID, {
      source,
      sourceMap: readModuleSourceMap(modulePath, sourceDirectory),
    });
    for (const dependency of dependencies) visit(dependency);
  }
}

function readModuleSourceMap(modulePath, sourceDirectory) {
  const mapPath = `${modulePath}.map`;
  const sourceMap = JSON.parse(readFileSync(mapPath, 'utf8'));
  if (sourceMap.version !== 3 || !Array.isArray(sourceMap.sources)) {
    throw new Error(`invalid TypeScript source map: ${mapPath}`);
  }
  if (!Array.isArray(sourceMap.sourcesContent)
      || sourceMap.sourcesContent.length !== sourceMap.sources.length) {
    throw new Error(`TypeScript source map lacks inline sources: ${mapPath}`);
  }

  const sourceRoot = sourceMap.sourceRoot ?? '';
  sourceMap.sources = sourceMap.sources.map((source) => {
    const sourcePath = path.resolve(path.dirname(mapPath), sourceRoot, source);
    const relativeSource = path.relative(sourceDirectory, sourcePath);
    if (
      relativeSource.startsWith('..')
      || path.isAbsolute(relativeSource)
      || relativeSource === ''
    ) {
      throw new Error(`source map path escapes the TypeScript source root: ${sourcePath}`);
    }
    const repositorySource = path.join(externalRoot, relativeSource);
    return path.relative(distDirectory, repositorySource).split(path.sep).join('/');
  });
  delete sourceMap.file;
  delete sourceMap.sourceRoot;
  return sourceMap;
}

function createIdentitySourceMap(sourceName, source) {
  const lineCount = source.split('\n').length;
  return {
    version: 3,
    sources: [sourceName],
    names: [],
    mappings: Array.from(
      { length: lineCount },
      (_, index) => (index === 0 ? 'AAAA' : 'AACA'),
    ).join(';'),
    sourcesContent: [source],
  };
}

function createBundleWriter(outputFile) {
  const chunks = [];
  const sections = [];
  let line = 0;
  let column = 0;

  return {
    append(code, sourceMap) {
      if (sourceMap) {
        sections.push({ offset: { line, column }, map: sourceMap });
      }
      chunks.push(code);
      const lastNewline = code.lastIndexOf('\n');
      if (lastNewline < 0) {
        column += code.length;
      } else {
        line += code.match(/\n/g)?.length ?? 0;
        column = code.length - lastNewline - 1;
      }
      if (sourceMap) {
        if (column !== 0) this.append('\n');
        sections.push({
          offset: { line, column },
          map: { version: 3, sources: [], names: [], mappings: 'A' },
        });
      }
    },
    finish() {
      return {
        code: chunks.join(''),
        map: {
          version: 3,
          file: outputFile,
          // Node's source-map loader normalizes this field before it parses
          // indexed sections, so retain the empty compatibility field.
          sources: [],
          sourceRoot: '',
          sections,
        },
        outputFile,
      };
    },
  };
}

function writeBundle(bundle) {
  writeFileSync(path.join(distDirectory, bundle.outputFile), bundle.code);
  writeFileSync(
    path.join(distDirectory, `${bundle.outputFile}.map`),
    `${JSON.stringify(bundle.map)}\n`,
  );
}

function resolveCommonJSModule(commonJSDirectory, fromID, request) {
  const fromPath = path.join(commonJSDirectory, ...fromID.split('/'));
  const basePath = path.resolve(path.dirname(fromPath), request);
  const candidates = path.extname(basePath)
    ? [basePath]
    : [`${basePath}.js`, path.join(basePath, 'index.js')];
  for (const candidate of candidates) {
    try {
      if (!statSync(candidate).isFile()) continue;
    } catch {
      continue;
    }
    const relativePath = path.relative(commonJSDirectory, candidate);
    if (relativePath.startsWith('..') || path.isAbsolute(relativePath)) {
      throw new Error(`module ${JSON.stringify(request)} escapes the build root`);
    }
    return relativePath.split(path.sep).join('/');
  }
  throw new Error(`cannot resolve ${JSON.stringify(request)} from ${fromID}`);
}

function exportedNames(entrySource) {
  const names = new Set();
  for (const match of entrySource.matchAll(/\bexports\.([A-Za-z_$][\w$]*)\s*=/g)) {
    if (match[1] !== '__esModule') names.add(match[1]);
  }
  for (const match of entrySource.matchAll(
    /Object\.defineProperty\(exports,\s*['"]([A-Za-z_$][\w$]*)['"]/g,
  )) {
    if (match[1] !== '__esModule') names.add(match[1]);
  }
  if (names.size === 0) throw new Error('index entrypoint has no named exports');
  return [...names].sort();
}

function createExportFooter(entryID, names) {
  const exports = names
    .map((name) => `export const ${name} = __xaligoEntry.${name};`)
    .join('\n');
  return `const __xaligoEntry = __xaligoRequire(${JSON.stringify(entryID)});\n${exports}\n`;
}

function watch() {
  let timer;
  let snapshot = watchSnapshot();
  // Poll the small source set so directory watchers do not exhaust file
  // descriptors on constrained development hosts.
  const poller = setInterval(poll, 250);

  function poll() {
    let nextSnapshot;
    try {
      nextSnapshot = watchSnapshot();
    } catch (error) {
      console.error(error);
      return;
    }
    if (nextSnapshot === snapshot) return;
    snapshot = nextSnapshot;
    clearTimeout(timer);
    timer = setTimeout(() => {
      try {
        build();
      } catch (error) {
        console.error(error);
      }
    }, 100);
  }

  const close = () => {
    clearTimeout(timer);
    clearInterval(poller);
    process.exitCode = 0;
  };
  process.once('SIGINT', close);
  process.once('SIGTERM', close);
  console.log('[xaligo-build] watching TypeScript sources');
}

function watchSnapshot() {
  return [
    ...findTypeScriptSources(externalRoot),
    path.join(externalRoot, 'package.json'),
    tsconfigPath,
  ].map((watchedFile) => {
    const metadata = statSync(watchedFile, { bigint: true });
    return `${watchedFile}\0${metadata.mtimeNs}\0${metadata.size}`;
  }).join('\n');
}
