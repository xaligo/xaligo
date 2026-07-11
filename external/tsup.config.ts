import { defineConfig } from 'tsup';

export default defineConfig({
  entry: {
    index: 'index.ts',
    cli: 'command.ts',
  },
  format: ['esm'],
  splitting: false,
  external: ['@resvg/resvg-js'],
  noExternal: ['pptxgenjs', 'jszip'],
  dts: true,
  clean: true,
  sourcemap: true,
});
