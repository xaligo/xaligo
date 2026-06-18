package repository

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// PptxExportOptions contains the inputs passed to the Node/PptxGenJS exporter.
type PptxExportOptions struct {
	XalPath      string
	Output       string
	ServicesFile string
	Title        string
	Author       string
	Company      string
	Subject      string
	Compression  *bool
	PxPerInch    float64
	ArrowStyle   string
	ArrowStub    float64
	ArrowMargin  float64
	Paper        string
	Orientation  string
	Stdout       io.Writer
	Stderr       io.Writer
}

// ExportPptx shells out to the Node/PptxGenJS exporter.
func ExportPptx(opts PptxExportOptions) error {
	cmdName, args, err := pptxExporterCommand()
	if err != nil {
		return err
	}
	args = append(args, "--xal", opts.XalPath, "-o", opts.Output)
	if opts.ServicesFile != "" {
		args = append(args, "--services", opts.ServicesFile)
	}
	if opts.Title != "" {
		args = append(args, "--title", opts.Title)
	}
	if opts.Author != "" {
		args = append(args, "--author", opts.Author)
	}
	if opts.Company != "" {
		args = append(args, "--company", opts.Company)
	}
	if opts.Subject != "" {
		args = append(args, "--subject", opts.Subject)
	}
	if opts.Compression != nil {
		args = append(args, "--compression", fmt.Sprintf("%t", *opts.Compression))
	}
	if opts.PxPerInch > 0 {
		args = append(args, "--px-per-inch", fmt.Sprintf("%g", opts.PxPerInch))
	}
	if opts.ArrowStyle != "" {
		args = append(args, "--arrow-style", opts.ArrowStyle)
	}
	if opts.ArrowStub > 0 {
		args = append(args, "--arrow-stub", fmt.Sprintf("%g", opts.ArrowStub))
	}
	if opts.ArrowMargin > 0 {
		args = append(args, "--arrow-margin", fmt.Sprintf("%g", opts.ArrowMargin))
	}
	if opts.Paper != "" {
		args = append(args, "--paper", opts.Paper)
	}
	if opts.Orientation != "" {
		args = append(args, "--orientation", opts.Orientation)
	}

	cmd := exec.Command(cmdName, args...)
	cmd.Stdout = opts.Stdout
	cmd.Stderr = opts.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run PPTX exporter: %w", err)
	}
	return nil
}

func pptxExporterCommand() (string, []string, error) {
	if cliPath, ok := findPptxExporterCLI(); ok {
		if _, err := exec.LookPath("node"); err != nil {
			return "", nil, fmt.Errorf("node executable not found; install Node.js to use `xaligo generate pptx`")
		}
		return "node", []string{cliPath}, nil
	}
	if binPath, err := exec.LookPath("xaligo-pptx"); err == nil {
		return binPath, nil, nil
	}
	return "", nil, fmt.Errorf("PPTX exporter not found; run `npm run build --workspace packages/xaligo` and `make build-wasm`, or install the npm package so `xaligo-pptx` is on PATH")
}

func findPptxExporterCLI() (string, bool) {
	const rel = "packages/xaligo/dist/cli.mjs"
	var bases []string
	if wd, err := os.Getwd(); err == nil {
		bases = append(bases, wd)
	}
	if exe, err := os.Executable(); err == nil {
		dir := filepath.Dir(exe)
		bases = append(bases, dir, filepath.Dir(dir))
	}

	seen := map[string]bool{}
	for _, base := range bases {
		for dir := base; ; dir = filepath.Dir(dir) {
			candidate := filepath.Join(dir, rel)
			if !seen[candidate] {
				seen[candidate] = true
				if st, err := os.Stat(candidate); err == nil && !st.IsDir() {
					return candidate, true
				}
			}
			next := filepath.Dir(dir)
			if next == dir {
				break
			}
		}
	}
	return "", false
}
