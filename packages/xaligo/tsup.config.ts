import { defineConfig } from 'tsup';

export default defineConfig({
  entry: ['src/index.ts', 'src/cli.ts'],
  format: ['esm', 'cjs'],
  external: ['@resvg/resvg-js'],
  dts: true,
  clean: true,
  sourcemap: true,
});
