package repository

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/xaligo/xaligo/internal/config"
)

const pptxExporterWasmRel = "external/wasm/xaligo.wasm"

func exportPptxWithWASM(ctx context.Context, path string, requestJSON []byte) ([]byte, []byte, error) {
	wasmPath, err := resolvePptxExporterWASM(path)
	if err != nil {
		return nil, nil, err
	}
	stdout, stderr, err := runPptxExporterWASM(ctx, wasmPath, requestJSON)
	if err != nil {
		return nil, stderr, err
	}
	pptxBytes, decodeErr := base64.StdEncoding.DecodeString(strings.TrimSpace(string(stdout)))
	if decodeErr != nil {
		return nil, stderr, fmt.Errorf("decode PPTX WASM base64 output: %w", decodeErr)
	}
	return pptxBytes, stderr, nil
}

func runPptxExporterWASM(ctx context.Context, wasmPath string, stdin []byte) ([]byte, []byte, error) {
	runtime := wazero.NewRuntime(ctx)
	defer runtime.Close(ctx)

	if _, err := wasi_snapshot_preview1.Instantiate(ctx, runtime); err != nil {
		return nil, nil, fmt.Errorf("instantiate WASI imports: %w", err)
	}
	wasmBytes, err := os.ReadFile(wasmPath)
	if err != nil {
		return nil, nil, fmt.Errorf("read PPTX WASM exporter %s: %w", wasmPath, err)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cfg := wazero.NewModuleConfig().
		WithName("xaligo-pptx-exporter").
		WithArgs(wasmPath, "pptx-exporter").
		WithStdin(bytes.NewReader(stdin)).
		WithStdout(&stdout).
		WithStderr(&stderr)

	if _, err := runtime.InstantiateWithConfig(ctx, wasmBytes, cfg); err != nil {
		return nil, stderr.Bytes(), fmt.Errorf("run PPTX WASM exporter: %w", err)
	}
	return stdout.Bytes(), stderr.Bytes(), nil
}

func resolvePptxExporterWASM(explicit string) (string, error) {
	candidates := []string{}
	if explicit != "" {
		candidates = append(candidates, explicit)
	}
	if env := os.Getenv("XALIGO_PPTX_EXPORTER_WASM"); env != "" {
		candidates = append(candidates, env)
	}
	if cfgPath := config.New().PptxExporterWASM; cfgPath != "" {
		candidates = append(candidates, cfgPath)
	}
	for _, base := range searchBases() {
		candidates = append(candidates, filepath.Join(base, pptxExporterWasmRel))
	}
	seen := map[string]bool{}
	for _, candidate := range candidates {
		if candidate == "" || seen[candidate] {
			continue
		}
		seen[candidate] = true
		if st, err := os.Stat(candidate); err == nil && !st.IsDir() {
			return candidate, nil
		}
	}
	return "", fmt.Errorf("PPTX WASM exporter not found; configure paths.pptx_exporter_wasm, provide %s, or set XALIGO_PPTX_EXPORTER_WASM", pptxExporterWasmRel)
}

func searchBases() []string {
	var bases []string
	if wd, err := os.Getwd(); err == nil {
		bases = append(bases, wd)
	}
	if exe, err := os.Executable(); err == nil {
		dir := filepath.Dir(exe)
		bases = append(bases, dir, filepath.Dir(dir))
	}

	var out []string
	seen := map[string]bool{}
	for _, base := range bases {
		for dir := base; ; dir = filepath.Dir(dir) {
			if !seen[dir] {
				out = append(out, dir)
				seen[dir] = true
			}
			next := filepath.Dir(dir)
			if next == dir {
				break
			}
		}
	}
	return out
}
