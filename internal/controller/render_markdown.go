package controller

import (
	"context"
	"fmt"
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
	if err := os.MkdirAll(absSVGDir, 0o755); err != nil {
		logger.ERROR(ICRRM004, "create svg directory failed", map[string]any{"svgDir": absSVGDir, "error": err})
		return fmt.Errorf("create svg directory: %w", err)
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
	seen := map[string]string{}
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
			if writeErr := os.WriteFile(svgPath, artifact.Data, 0o644); writeErr != nil {
				return nil, fmt.Errorf("write SVG file %s: %w", svgPath, writeErr)
			}
			relPath, relErr := filepath.Rel(outputDir, svgPath)
			if relErr != nil {
				relPath = svgPath
			}
			lines = append(lines, fmt.Sprintf("![](%s)", filepath.ToSlash(relPath)), "")
		}
		return lines, nil
	})
	if err != nil {
		logger.ERROR(ICRRM005, "embed code blocks failed", map[string]any{"error": err})
		return err
	}

	if err := os.WriteFile(outputPath, []byte(embedded), 0o644); err != nil {
		logger.ERROR(ICRRM006, "write output failed", map[string]any{"output": outputPath, "error": err})
		return fmt.Errorf("write output file: %w", err)
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
