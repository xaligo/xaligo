package repository

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryo-arima/xaligo/internal/config"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

const pptxExporterWasmRel = "external/wasm/xaligo.wasm"

var (
	logger     = share.DefaultLogger()
	IRPEPWX001 = share.NewMCode("IRPEPWX-001", "Export PPTX with exporter generated output")
	IRPEPWX002 = share.NewMCode("IRPEPWX-002", "Export PPTX bytes generated output")
)

type PptxExporter interface {
	Export(ctx context.Context, requestJSON []byte) (stdout []byte, stderr []byte, err error)
}

type PowerpointRepository interface {
	WritePptx(opts entity.PptxExportOptions) error
	WritePptxWithExporter(ctx context.Context, opts entity.PptxExportOptions, exporter PptxExporter) error
	ExportPptxBytes(ctx context.Context, opts entity.PptxExportOptions) ([]byte, error)
	ExportPptxBytesWithExporter(ctx context.Context, opts entity.PptxExportOptions, exporter PptxExporter) ([]byte, error)
}

type powerpointRepository struct{}

func NewPowerpointRepository() PowerpointRepository {
	return &powerpointRepository{}
}

type WASMPptxExporter struct {
	Path string
}

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

// WritePptx invokes the WASM PPTX exporter and writes the returned PPTX bytes.
func (rcvr *powerpointRepository) WritePptx(opts entity.PptxExportOptions) error {
	return rcvr.WritePptxWithExporter(context.Background(), opts, WASMPptxExporter{Path: opts.ExporterWASM})
}

func (rcvr *powerpointRepository) WritePptxWithExporter(ctx context.Context, opts entity.PptxExportOptions, exporter PptxExporter) error {
	pptxBytes, err := rcvr.ExportPptxBytesWithExporter(ctx, opts, exporter)
	if err != nil {
		return err
	}
	if opts.Output == "" {
		return fmt.Errorf("output path is required")
	}
	if err := os.WriteFile(opts.Output, pptxBytes, 0644); err != nil {
		return fmt.Errorf("write PPTX output %s: %w", opts.Output, err)
	}
	logger.INFO(IRPEPWX001, "generated", map[string]any{"output": opts.Output})
	return nil
}

func (rcvr *powerpointRepository) ExportPptxBytes(ctx context.Context, opts entity.PptxExportOptions) ([]byte, error) {
	return rcvr.ExportPptxBytesWithExporter(ctx, opts, WASMPptxExporter{Path: opts.ExporterWASM})
}

func (rcvr *powerpointRepository) ExportPptxBytesWithExporter(ctx context.Context, opts entity.PptxExportOptions, exporter PptxExporter) ([]byte, error) {
	if len(bytes.TrimSpace(opts.PlanJSON)) == 0 {
		return nil, fmt.Errorf("PPTX plan JSON is required")
	}
	if exporter == nil {
		exporter = WASMPptxExporter{Path: opts.ExporterWASM}
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
		return nil, fmt.Errorf("encode PPTX WASM request: %w", err)
	}

	pptxBytes, stderr, err := exporter.Export(ctx, reqJSON)
	if len(stderr) > 0 {
		if opts.Stderr != nil {
			_, _ = opts.Stderr.Write(stderr)
		}
	}
	if err != nil {
		return nil, err
	}
	if len(pptxBytes) == 0 {
		return nil, fmt.Errorf("PPTX WASM exporter produced no output")
	}
	logger.INFO(IRPEPWX002, "generated")
	return pptxBytes, nil
}

func (rcvr WASMPptxExporter) Export(ctx context.Context, requestJSON []byte) ([]byte, []byte, error) {
	wasmPath, err := resolvePptxExporterWASM(rcvr.Path)
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
