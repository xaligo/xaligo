package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ryo-arima/xaligo/internal/config"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

const pptxExporterWasmRel = "external/wasm/pptx_exporter.wasm"

type pptxWasmRequest struct {
	Plan    json.RawMessage `json:"plan"`
	Options pptxWasmOptions `json:"options,omitempty"`
}

type pptxWasmOptions struct {
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	Company     string `json:"company,omitempty"`
	Subject     string `json:"subject,omitempty"`
	Compression *bool  `json:"compression,omitempty"`
}

// ExportPptx invokes the WASM PPTX exporter and writes the returned PPTX bytes.
func ExportPptx(opts entity.PptxExportOptions) error {
	if len(bytes.TrimSpace(opts.PlanJSON)) == 0 {
		return fmt.Errorf("PPTX plan JSON is required")
	}
	if opts.Output == "" {
		return fmt.Errorf("output path is required")
	}
	req := pptxWasmRequest{
		Plan: json.RawMessage(opts.PlanJSON),
		Options: pptxWasmOptions{
			Title:       opts.Title,
			Author:      opts.Author,
			Company:     opts.Company,
			Subject:     opts.Subject,
			Compression: opts.Compression,
		},
	}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("encode PPTX WASM request: %w", err)
	}

	wasmPath, err := resolvePptxExporterWASM(opts.ExporterWASM)
	if err != nil {
		return err
	}

	pptxBytes, stderr, err := runPptxExporterWASM(wasmPath, reqJSON)
	if len(stderr) > 0 {
		if opts.Stderr != nil {
			_, _ = opts.Stderr.Write(stderr)
		}
	}
	if err != nil {
		return err
	}
	if len(pptxBytes) == 0 {
		return fmt.Errorf("PPTX WASM exporter produced no output")
	}

	if err := os.WriteFile(opts.Output, pptxBytes, 0644); err != nil {
		return fmt.Errorf("write PPTX output %s: %w", opts.Output, err)
	}
	if opts.Stdout != nil {
		_, _ = fmt.Fprintf(opts.Stdout, "generated: %s\n", opts.Output)
	}
	return nil
}

func runPptxExporterWASM(wasmPath string, stdin []byte) ([]byte, []byte, error) {
	ctx := context.Background()
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
		WithArgs(wasmPath).
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
