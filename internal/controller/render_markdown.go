package controller

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
)

var (
	ICRIRMC001 = share.NewMCode("ICRIRMC-001", "Init render markdown command start")
	ICRRM001   = share.NewMCode("ICRRM-001", "Run render markdown start")
	ICRRM002   = share.NewMCode("ICRRM-002", "Run render markdown read input failed")
	ICRRM003   = share.NewMCode("ICRRM-003", "Run render markdown read services failed")
	ICRRM004   = share.NewMCode("ICRRM-004", "Run render markdown create SVG directory failed")
	ICRRM005   = share.NewMCode("ICRRM-005", "Run render markdown embed code blocks failed")
	ICRRM006   = share.NewMCode("ICRRM-006", "Run render markdown write output failed")
	ICRRM007   = share.NewMCode("ICRRM-007", "Run render markdown generated output")
	IEXCB002   = share.NewMCode("IEXCB-002", "Embed xal code blocks render block failed")
)

// initRenderMarkdownCmd returns the `xaligo render markdown` subcommand, which
// scans a Markdown file for ```xal fenced code blocks, renders each block to
// SVG, writes the SVG files, and emits a new Markdown file with a Markdown
// image reference embedded at each code block's location.
func initRenderMarkdownCmd(rcvr *renderController) *cobra.Command {
	logger.DEBUG(ICRIRMC001, "start")
	var (
		output       string
		svgDir       string
		servicesFile string
		theme        string
		mode         string
		pxPerInch    float64
		paper        string
		orientation  string
	)

	cmd := &cobra.Command{
		Use:   "markdown <input.md>",
		Short: "Embed rendered ```xal code blocks as SVG images into a Markdown file",
		Long: `Read a Markdown file, render every fenced code block whose info string is
"xal" to SVG using the shared render pipeline, and write a new Markdown file
with a Markdown image reference (![](path/to/file.svg)) in place of each
original code block. Generated SVG files are written next to the input
Markdown file by default; use --svg-dir to write them elsewhere.

Use --paper and --orientation to fit each rendered diagram to a specific
physical page size, the same as 'render --format svg'.

Examples:
  xaligo render markdown docs/guide.md
  xaligo render markdown docs/guide.md -o docs/guide.embedded.md
  xaligo render markdown docs/guide.md --svg-dir docs/images
  xaligo render markdown docs/guide.md --paper A4 --orientation landscape`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return rcvr.RunMarkdown(entity.ControllerRenderMarkdownOptions{
				InputPath:    args[0],
				OutputPath:   output,
				SVGDir:       svgDir,
				ServicesFile: servicesFile,
				Theme:        theme,
				Mode:         mode,
				PxPerInch:    pxPerInch,
				Paper:        paper,
				Orientation:  orientation,
			})
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "output Markdown file path (default: <input-stem>.embedded.md next to the input)")
	cmd.Flags().StringVar(&svgDir, "svg-dir", "", "directory to write generated SVG files (default: the input Markdown file's directory)")
	cmd.Flags().StringVar(&servicesFile, "services", "", "optional services.csv for icon labels")
	cmd.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	cmd.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	cmd.Flags().Float64Var(&pxPerInch, "px-per-inch", 0, "pixels per inch for SVG layout scaling (default 96)")
	cmd.Flags().StringVar(&paper, "paper", "", "physical paper size to fit each diagram to: A5 | A4 | A3 | A2 | A1 | Letter | Legal | Tabloid (default: auto-fit)")
	cmd.Flags().StringVar(&orientation, "orientation", "", "page orientation: portrait | landscape (default: auto-fit)")
	return cmd
}

// RunMarkdown renders every ```xal code block in a Markdown file to SVG and
// writes a new Markdown file with the rendered SVGs embedded as image
// references in place of each original code block.
func (rcvr *renderController) RunMarkdown(opts entity.ControllerRenderMarkdownOptions) error {
	logger.DEBUG(ICRRM001, "start", map[string]any{"input": opts.InputPath, "output": opts.OutputPath})
	input, err := os.ReadFile(opts.InputPath)
	if err != nil {
		logger.ERROR(ICRRM002, "read input failed", map[string]any{"input": opts.InputPath, "error": err})
		return fmt.Errorf("read input file: %w", err)
	}
	absInputPath, err := filepath.Abs(opts.InputPath)
	if err != nil {
		return fmt.Errorf("resolve input file path: %w", err)
	}

	outputPath := opts.OutputPath
	if outputPath == "" {
		outputPath = defaultRenderMarkdownOutput(opts.InputPath)
	}
	absOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		return fmt.Errorf("resolve output file path: %w", err)
	}

	svgDir := opts.SVGDir
	if svgDir == "" {
		svgDir = filepath.Dir(absInputPath)
	}
	absSVGDir, err := filepath.Abs(svgDir)
	if err != nil {
		return fmt.Errorf("resolve svg directory: %w", err)
	}
	theme, err := entity.NormalizeTheme(opts.Theme)
	if err != nil {
		return err
	}
	renderOpts := entity.RenderOptions{
		Format:      usecase.FormatSVG,
		Mode:        entity.Mode(strings.TrimSpace(opts.Mode)),
		Theme:       theme,
		PxPerInch:   opts.PxPerInch,
		PaperSize:   opts.Paper,
		Orientation: opts.Orientation,
		Imports:     &entity.ImportSource{FS: os.DirFS(filepath.Dir(absInputPath))},
	}
	if opts.ServicesFile != "" {
		servicesCSV, readErr := os.ReadFile(opts.ServicesFile)
		if readErr != nil {
			logger.ERROR(ICRRM003, "read services failed", map[string]any{"servicesFile": opts.ServicesFile, "error": readErr})
			return fmt.Errorf("read services %s: %w", opts.ServicesFile, readErr)
		}
		renderOpts.ServicesCSV = servicesCSV
	}
	if err := rcvr.renderUsecase.ValidateRenderOptions(renderOpts); err != nil {
		return err
	}

	mdStem := strings.TrimSuffix(filepath.Base(opts.InputPath), filepath.Ext(opts.InputPath))
	outputDir := filepath.Dir(absOutputPath)
	seen := map[string]string{
		strings.ToLower(filepath.Clean(absOutputPath)): "Markdown output",
	}
	pendingOutputs := make([]renderMarkdownPendingOutput, 0)
	blockIndex := 0

	embedded, err := usecase.EmbedXalCodeBlocks(string(input), func(xal string) ([]string, error) {
		blockIndex++
		artifacts, renderErr := rcvr.renderUsecase.RenderArtifacts(context.Background(), []byte(xal), renderOpts)
		if renderErr != nil {
			logger.ERROR(IEXCB002, "render block failed", map[string]any{"block": blockIndex, "error": renderErr})
			return nil, fmt.Errorf("render xal code block %d: %w", blockIndex, renderErr)
		}
		lines := make([]string, 0, len(artifacts)*2)
		for artifactIndex, artifact := range artifacts {
			name := renderMarkdownSVGFileName(mdStem, blockIndex, artifactIndex, artifact.ID, len(artifacts))
			svgPath := filepath.Join(absSVGDir, name)
			collisionKey := strings.ToLower(filepath.Clean(svgPath))
			if prior, exists := seen[collisionKey]; exists {
				return nil, fmt.Errorf("code block %d frame %q and %s resolve to the same SVG output %s", blockIndex, artifact.ID, prior, svgPath)
			}
			seen[collisionKey] = fmt.Sprintf("block %d frame %q", blockIndex, artifact.ID)
			pendingOutputs = append(pendingOutputs, renderMarkdownPendingOutput{path: svgPath, data: artifact.Data})
			relPath, relErr := rcvr.renderMarkdownFileOperations.relativePath(outputDir, svgPath)
			if relErr != nil {
				return nil, fmt.Errorf(
					"make SVG output %s relative to Markdown output directory %s: %w; "+
						"place both outputs on the same filesystem volume",
					svgPath, outputDir, relErr,
				)
			}
			lines = append(lines, renderMarkdownImageReference(relPath), "")
		}
		return lines, nil
	})
	if err != nil {
		logger.ERROR(ICRRM005, "embed code blocks failed", map[string]any{"error": err})
		return err
	}

	pendingOutputs = append(pendingOutputs, renderMarkdownPendingOutput{path: absOutputPath, data: []byte(embedded)})
	if err := writeRenderMarkdownOutputs(pendingOutputs, rcvr.renderMarkdownFileOperations); err != nil {
		logger.ERROR(ICRRM006, "write output failed", map[string]any{"output": outputPath, "error": err})
		return fmt.Errorf("write Markdown output set: %w", err)
	}
	logger.INFO(ICRRM007, "generated", map[string]any{"input": opts.InputPath, "output": outputPath, "svgDir": svgDir, "blocks": blockIndex})
	return nil
}

func defaultRenderMarkdownOutput(inputPath string) string {
	ext := filepath.Ext(inputPath)
	stem := strings.TrimSuffix(inputPath, ext)
	if ext == "" {
		ext = ".md"
	}
	return stem + ".embedded" + ext
}

func renderMarkdownSVGFileName(mdStem string, blockIndex, artifactIndex int, artifactID string, artifactCount int) string {
	if artifactCount <= 1 {
		return fmt.Sprintf("%s-%d.svg", mdStem, blockIndex)
	}
	safeID := safeRenderArtifactID(artifactID)
	if safeID == "" {
		safeID = fmt.Sprintf("frame-%d", artifactIndex+1)
	}
	return fmt.Sprintf("%s-%d-%s.svg", mdStem, blockIndex, safeID)
}

type renderMarkdownPendingOutput struct {
	path string
	data []byte
}

type renderMarkdownStagedOutput struct {
	target    string
	temporary string
	backup    string
	installed bool
}

// RenderMarkdownFileOperations provides the filesystem operations whose
// failures must be simulated to verify output-set rollback behavior.
type RenderMarkdownFileOperations struct {
	Remove       func(string) error
	Rename       func(string, string) error
	RelativePath func(string, string) (string, error)
}

type renderMarkdownFileOperations struct {
	remove       func(string) error
	rename       func(string, string) error
	relativePath func(string, string) (string, error)
}

func defaultRenderMarkdownFileOperations() renderMarkdownFileOperations {
	return renderMarkdownFileOperations{
		remove:       os.Remove,
		rename:       os.Rename,
		relativePath: filepath.Rel,
	}
}

// WithRenderMarkdownFileOperations overrides selected Markdown output
// filesystem operations while retaining native defaults for nil callbacks.
func WithRenderMarkdownFileOperations(operations RenderMarkdownFileOperations) RenderControllerOption {
	return func(controller *renderController) {
		if operations.Remove != nil {
			controller.renderMarkdownFileOperations.remove = operations.Remove
		}
		if operations.Rename != nil {
			controller.renderMarkdownFileOperations.rename = operations.Rename
		}
		if operations.RelativePath != nil {
			controller.renderMarkdownFileOperations.relativePath = operations.RelativePath
		}
	}
}

func renderMarkdownImageReference(path string) string {
	normalized := filepath.ToSlash(path)
	escaped := (&url.URL{Path: normalized}).EscapedPath()
	return fmt.Sprintf("![](<%s>)", escaped)
}

func writeRenderMarkdownOutputs(outputs []renderMarkdownPendingOutput, operations renderMarkdownFileOperations) error {
	staged := make([]renderMarkdownStagedOutput, 0, len(outputs))
	cleanupTemporary := func() error {
		failures := make([]error, 0)
		for index := range staged {
			temporary := staged[index].temporary
			if temporary == "" {
				continue
			}
			if err := operations.remove(temporary); err != nil && !os.IsNotExist(err) {
				failures = append(failures, fmt.Errorf("remove staged output %s: %w", temporary, err))
				continue
			}
			staged[index].temporary = ""
		}
		return errors.Join(failures...)
	}
	rollback := func() error {
		failures := make([]error, 0)
		for index := len(staged) - 1; index >= 0; index-- {
			output := &staged[index]
			if output.installed {
				if err := operations.remove(output.target); err != nil && !os.IsNotExist(err) {
					failures = append(failures, fmt.Errorf("remove newly installed output %s: %w", output.target, err))
				} else {
					output.installed = false
				}
			}
			if output.backup != "" {
				backup := output.backup
				if err := operations.rename(backup, output.target); err != nil {
					failures = append(failures, fmt.Errorf(
						"restore existing output %s from backup %s: %w; previous output is preserved at %s",
						output.target, backup, err, backup,
					))
				} else {
					output.backup = ""
				}
			}
		}
		if err := cleanupTemporary(); err != nil {
			failures = append(failures, err)
		}
		return errors.Join(failures...)
	}
	failStaging := func(primary error, additional ...error) error {
		failures := []error{primary}
		for _, err := range additional {
			if err != nil {
				failures = append(failures, err)
			}
		}
		if err := cleanupTemporary(); err != nil {
			failures = append(failures, err)
		}
		return errors.Join(failures...)
	}
	failTransaction := func(primary error, additional ...error) error {
		failures := []error{primary}
		for _, err := range additional {
			if err != nil {
				failures = append(failures, err)
			}
		}
		if err := rollback(); err != nil {
			failures = append(failures, fmt.Errorf("rollback Markdown output set: %w", err))
		}
		return errors.Join(failures...)
	}

	for _, output := range outputs {
		directory := filepath.Dir(output.path)
		if err := os.MkdirAll(directory, 0o755); err != nil {
			return failStaging(fmt.Errorf("create output directory %s: %w", directory, err))
		}
		info, err := os.Stat(output.path)
		if err == nil && info.IsDir() {
			return failStaging(fmt.Errorf("output path is a directory: %s", output.path))
		}
		if err != nil && !os.IsNotExist(err) {
			return failStaging(fmt.Errorf("inspect output path %s: %w", output.path, err))
		}

		file, err := os.CreateTemp(directory, ".xaligo-markdown-*")
		if err != nil {
			return failStaging(fmt.Errorf("create staged output for %s: %w", output.path, err))
		}
		temporary := file.Name()
		staged = append(staged, renderMarkdownStagedOutput{target: output.path, temporary: temporary})
		if err := file.Chmod(0o644); err != nil {
			closeErr := file.Close()
			return failStaging(
				fmt.Errorf("set staged output permissions for %s: %w", output.path, err),
				wrapRenderMarkdownOutputError("close staged output "+temporary, closeErr),
			)
		}
		if _, err := file.Write(output.data); err != nil {
			closeErr := file.Close()
			return failStaging(
				fmt.Errorf("write staged output for %s: %w", output.path, err),
				wrapRenderMarkdownOutputError("close staged output "+temporary, closeErr),
			)
		}
		if err := file.Sync(); err != nil {
			closeErr := file.Close()
			return failStaging(
				fmt.Errorf("sync staged output for %s: %w", output.path, err),
				wrapRenderMarkdownOutputError("close staged output "+temporary, closeErr),
			)
		}
		if err := file.Close(); err != nil {
			return failStaging(fmt.Errorf("close staged output for %s: %w", output.path, err))
		}
	}

	for index := range staged {
		if _, err := os.Lstat(staged[index].target); err == nil {
			backupFile, createErr := os.CreateTemp(filepath.Dir(staged[index].target), ".xaligo-markdown-backup-*")
			if createErr != nil {
				return failTransaction(fmt.Errorf("create backup path for %s: %w", staged[index].target, createErr))
			}
			backup := backupFile.Name()
			if closeErr := backupFile.Close(); closeErr != nil {
				removeErr := operations.remove(backup)
				return failTransaction(
					fmt.Errorf("close backup path for %s: %w", staged[index].target, closeErr),
					wrapRenderMarkdownOutputError("remove unused backup path "+backup, removeErr),
				)
			}
			if removeErr := operations.remove(backup); removeErr != nil {
				return failTransaction(fmt.Errorf(
					"prepare backup path for %s: %w; unused backup placeholder remains at %s",
					staged[index].target, removeErr, backup,
				))
			}
			if renameErr := operations.rename(staged[index].target, backup); renameErr != nil {
				return failTransaction(fmt.Errorf("back up existing output %s: %w", staged[index].target, renameErr))
			}
			staged[index].backup = backup
		} else if !os.IsNotExist(err) {
			return failTransaction(fmt.Errorf("inspect existing output %s: %w", staged[index].target, err))
		}
	}

	for index := range staged {
		if err := operations.rename(staged[index].temporary, staged[index].target); err != nil {
			return failTransaction(fmt.Errorf("install output %s: %w", staged[index].target, err))
		}
		staged[index].temporary = ""
		staged[index].installed = true
	}
	cleanupFailures := make([]error, 0)
	for index := range staged {
		output := &staged[index]
		if output.backup != "" {
			backup := output.backup
			if err := operations.remove(backup); err != nil && !os.IsNotExist(err) {
				cleanupFailures = append(cleanupFailures, fmt.Errorf(
					"Markdown output set was installed, but removing the previous output backup %s for %s failed: %w; previous output is preserved at %s",
					backup, output.target, err, backup,
				))
			} else {
				output.backup = ""
			}
		}
	}
	return errors.Join(cleanupFailures...)
}

func wrapRenderMarkdownOutputError(context string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", context, err)
}
