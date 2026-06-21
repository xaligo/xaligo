import { defineConfig } from 'tsup';

export default defineConfig({
  entry: {
    index: 'usecase/api.ts',
    cli: 'command.ts',
  },
  format: ['esm'],
  external: ['@resvg/resvg-js'],
  dts: true,
  clean: true,
  sourcemap: true,
});
