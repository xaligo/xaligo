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
	logger       = share.DefaultLogger()
	ICRRRFWUC001 = share.NewMCode("ICRRRFWUC-001", "Run render format with use case start")
	ICRRRFWUC002 = share.NewMCode("ICRRRFWUC-002", "Run render format with use case validate render options failed")
	ICRRRFWUC003 = share.NewMCode("ICRRRFWUC-003", "Run render format with use case normalize theme failed")
	ICRIRCWUC001 = share.NewMCode("ICRIRCWUC-001", "Init render command with use case start")
	ICRIRCWUC002 = share.NewMCode("ICRIRCWUC-002", "Init render command run start")
	ICRIRCWUC003 = share.NewMCode("ICRIRCWUC-003", "Init render command default output branch")
	ICRIRCWUC004 = share.NewMCode("ICRIRCWUC-004", "Init render command explicit output branch")
	ICRIRCWUC005 = share.NewMCode("ICRIRCWUC-005", "Init render command no compression branch")
	ICRIRCWUC006 = share.NewMCode("ICRIRCWUC-006", "Init render command compression branch")
	ICRIRCWUC007 = share.NewMCode("ICRIRCWUC-007", "Init render command return command")
	ICRRRF001    = share.NewMCode("ICRRRF-001", "Run render format start")
	ICRNRF001    = share.NewMCode("ICRNRF-001", "Normalize render format default branch")
	ICRNRF002    = share.NewMCode("ICRNRF-002", "Normalize render format explicit branch")
	ICRDRO001    = share.NewMCode("ICRDRO-001", "Default render output SVG branch")
	ICRDRO002    = share.NewMCode("ICRDRO-002", "Default render output PPTX branch")
	ICRRR001     = share.NewMCode("ICRRR-001", "Run render read input failed")
	ICRRR002     = share.NewMCode("ICRRR-002", "Run render read services failed")
	ICRRR003     = share.NewMCode("ICRRR-003", "Run render use case failed")
	ICRRR004     = share.NewMCode("ICRRR-004", "Run render write output failed")
	ICRRR005     = share.NewMCode("ICRRR-005", "Run render generated output")
)

type RenderController interface {
	Command() *cobra.Command
	RunFormat(opts entity.ControllerRenderOptions) error
	RunMarkdown(opts entity.ControllerRenderMarkdownOptions) error
}

// RenderControllerOption customizes a render controller dependency.
type RenderControllerOption func(*renderController)

type renderController struct {
	renderUsecase                usecase.RenderUsecase
	renderMarkdownFileOperations renderMarkdownFileOperations
}

func NewRenderController(renderUsecase usecase.RenderUsecase, options ...RenderControllerOption) RenderController {
	controller := &renderController{
		renderUsecase:                renderUsecase,
		renderMarkdownFileOperations: defaultRenderMarkdownFileOperations(),
	}
	for _, option := range options {
		if option != nil {
			option(controller)
		}
	}
	return controller
}

func (rcvr *renderController) Command() *cobra.Command {
	logger.DEBUG(ICRIRCWUC001, "start")
	var (
		output            string
		format            string
		servicesFile      string
		title             string
		author            string
		company           string
		subject           string
		compression       bool
		noCompression     bool
		pxPerInch         float64
		arrowStyle        string
		arrowStub         float64
		arrowMargin       float64
		paper             string
		orientation       string
		paperMargin       float64
		paperMarginTop    float64
		paperMarginRight  float64
		paperMarginBottom float64
		paperMarginLeft   float64
		exporterWASM      string
		theme             string
		mode              string
		svgLegendPosition string
		combineFrames     bool
	)

	cmd := &cobra.Command{
		Use:   "render <input.xal>",
		Short: "Render xaligo DSL into an output format",
		Long: `Render a .xal source file as SVG or PPTX.

Both formats share the same parser, layout, scene, routing, and draw-plan
pipeline, so geometry and theming remain consistent.

Identified child frames become separate SVG files or PPTX slides by default.
Pass --combine-frames to render them onto one canvas or slide instead.

Use 'xaligo render markdown <input.md>' to render every fenced ` + "```xal" + `
code block inside a Markdown file to SVG and embed the results as Markdown
image references.

Examples:
  xaligo render diagram.xal --format svg -o output/diagram.svg
  xaligo render diagram.xal --format pptx -o output/diagram.pptx --paper A3 --orientation landscape
  xaligo render diagram.xal --format svg -o output/diagram.svg --combine-frames`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			input := args[0]
			logger.DEBUG(ICRIRCWUC002, "start", map[string]any{"input": input, "format": format})
			format = normalizeRenderFormat(format)
			if output == "" {
				logger.DEBUG(ICRIRCWUC003, "branch default output", map[string]any{"format": format})
				output = defaultRenderOutput(format)
			} else {
				logger.DEBUG(ICRIRCWUC004, "branch explicit output", map[string]any{"output": output})
			}
			if noCompression {
				logger.DEBUG(ICRIRCWUC005, "branch no compression")
				compression = false
			} else {
				logger.DEBUG(ICRIRCWUC006, "branch compression", map[string]any{"compression": compression})
			}
			return rcvr.RunFormat(entity.ControllerRenderOptions{
				InputPath:         input,
				OutputPath:        output,
				Format:            format,
				ServicesFile:      servicesFile,
				Title:             title,
				Author:            author,
				Company:           company,
				Subject:           subject,
				CombineFrames:     combineFrames,
				Compression:       &compression,
				PxPerInch:         pxPerInch,
				ArrowStyle:        arrowStyle,
				ArrowStub:         arrowStub,
				ArrowMargin:       arrowMargin,
				Paper:             paper,
				Orientation:       orientation,
				PaperMargin:       paperMargin,
				PaperMarginTop:    paperMarginTop,
				PaperMarginRight:  paperMarginRight,
				PaperMarginBottom: paperMarginBottom,
				PaperMarginLeft:   paperMarginLeft,
				ExporterWASM:      exporterWASM,
				Theme:             theme,
				Mode:              mode,
				SVGLegendPosition: svgLegendPosition,
			})
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "output file path")
	cmd.Flags().StringVar(&format, "format", "svg", "output format: svg | pptx")
	cmd.Flags().StringVar(&servicesFile, "services", "", "optional services.csv for icon labels and output legends")
	cmd.Flags().BoolVar(&combineFrames, "combine-frames", false, "combine all frames on one SVG canvas or PPTX slide")
	cmd.Flags().StringVar(&title, "title", "", "optional PPTX title metadata")
	cmd.Flags().StringVar(&author, "author", "", "optional PPTX author metadata")
	cmd.Flags().StringVar(&company, "company", "", "optional PPTX company metadata")
	cmd.Flags().StringVar(&subject, "subject", "", "optional PPTX subject metadata")
	cmd.Flags().BoolVar(&compression, "compression", true, "compress PPTX output")
	cmd.Flags().BoolVar(&noCompression, "no-compression", false, "disable PPTX output compression")
	cmd.Flags().Float64Var(&pxPerInch, "px-per-inch", 0, "pixels per inch for PPTX layout scaling (default 96)")
	cmd.Flags().StringVar(&arrowStyle, "arrow-style", "", "connector arrow style: thin|standard|triangle|stealth|arrow|diamond|oval|none (default thin)")
	cmd.Flags().Float64Var(&arrowStub, "arrow-stub", 0, "stub length in px before the first/last bend (default 20)")
	cmd.Flags().Float64Var(&arrowMargin, "arrow-margin", 0, "clear margin in px reserved on both sides of each line (default 8)")
	cmd.Flags().StringVar(&paper, "paper", "", "slide paper size: A5 A4 A3 A2 A1 Letter Legal Tabloid (default: match .xal frame)")
	cmd.Flags().StringVar(&orientation, "orientation", "", "slide orientation: portrait | landscape (default: auto-fit)")
	cmd.Flags().Float64Var(&paperMargin, "paper-margin", 0, "paper margin in inches on all sides for paper fitting (default 0)")
	cmd.Flags().Float64Var(&paperMarginTop, "paper-margin-top", 0, "paper top margin in inches for paper fitting")
	cmd.Flags().Float64Var(&paperMarginRight, "paper-margin-right", 0, "paper right margin in inches for paper fitting")
	cmd.Flags().Float64Var(&paperMarginBottom, "paper-margin-bottom", 0, "paper bottom margin in inches for paper fitting")
	cmd.Flags().Float64Var(&paperMarginLeft, "paper-margin-left", 0, "paper left margin in inches for paper fitting")
	cmd.Flags().StringVar(&exporterWASM, "pptx-exporter-wasm", "", "path to the WASM PPTX exporter (default: external/exporter/wasm/xaligo.wasm or XALIGO_PPTX_EXPORTER_WASM)")
	cmd.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	cmd.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	cmd.Flags().StringVar(&svgLegendPosition, "svg-legend-position", "bottom", "SVG legend position when --services is provided: top | right | bottom | left")
	cmd.AddCommand(initRenderMarkdownCmd(rcvr))
	logger.DEBUG(ICRIRCWUC007, "return command")
	return cmd
}

// RunRenderFormat renders a .xal file into the requested output format. It is
// the public controller entry point for format-based rendering.
func (rcvr *renderController) RunFormat(opts entity.ControllerRenderOptions) error {
	logger.DEBUG(ICRRRF001, "start", map[string]any{"format": opts.Format, "input": opts.InputPath, "output": opts.OutputPath})
	logger.DEBUG(ICRRRFWUC001, "start", map[string]any{"format": opts.Format, "input": opts.InputPath, "output": opts.OutputPath})
	format := entity.Format(normalizeRenderFormat(opts.Format))
	if format == usecase.FormatPPTX && opts.OutputPath == "" {
		return fmt.Errorf("--output is required")
	}
	renderOpts := entity.RenderOptions{
		Mode: entity.Mode(opts.Mode), Format: format, Theme: opts.Theme,
		CombineFrames: opts.CombineFrames,
		Title:         opts.Title, Author: opts.Author, Company: opts.Company, Subject: opts.Subject, Compression: opts.Compression,
		PxPerInch: opts.PxPerInch, ArrowStyle: opts.ArrowStyle, ArrowStubPx: opts.ArrowStub, ArrowMarginPx: opts.ArrowMargin,
		PaperSize: opts.Paper, Orientation: opts.Orientation,
		PaperMarginIn: opts.PaperMargin, PaperMarginTopIn: opts.PaperMarginTop, PaperMarginRightIn: opts.PaperMarginRight,
		PaperMarginBottomIn: opts.PaperMarginBottom, PaperMarginLeftIn: opts.PaperMarginLeft,
		SVGLegendPosition: opts.SVGLegendPosition, PPTXExporterWASM: opts.ExporterWASM,
	}
	if err := rcvr.renderUsecase.ValidateRenderOptions(renderOpts); err != nil {
		logger.ERROR(ICRRRFWUC002, "validate render options failed", map[string]any{"error": err})
		return err
	}
	theme, err := entity.NormalizeTheme(opts.Theme)
	if err != nil {
		logger.ERROR(ICRRRFWUC003, "normalize theme failed", map[string]any{"theme": opts.Theme, "error": err})
		return err
	}
	renderOpts.Theme = theme

	if opts.ServicesFile != "" {
		renderOpts.ServicesCSV, err = os.ReadFile(opts.ServicesFile)
		if err != nil {
			logger.ERROR(ICRRR002, "read services failed", map[string]any{"servicesFile": opts.ServicesFile, "error": err})
			return fmt.Errorf("read services %s: %w", opts.ServicesFile, err)
		}
	}

	return runRender(rcvr.renderUsecase, opts.InputPath, opts.OutputPath, renderOpts)
}

func runRender(renderUsecase usecase.RenderUsecase, inputPath, outputPath string, opts entity.RenderOptions) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		logger.ERROR(ICRRR001, "read input failed", map[string]any{"input": inputPath, "error": err})
		return fmt.Errorf("read input file: %w", err)
	}
	absInputPath, err := filepath.Abs(inputPath)
	if err != nil {
		return fmt.Errorf("resolve input file path: %w", err)
	}
	opts.Imports = &entity.ImportSource{FS: os.DirFS(filepath.Dir(absInputPath))}
	if opts.Format == usecase.FormatSVG {
		artifacts, renderErr := renderUsecase.RenderArtifacts(context.Background(), input, opts)
		if renderErr != nil {
			logger.ERROR(ICRRR003, "render failed", map[string]any{"format": opts.Format, "error": renderErr})
			return renderErr
		}
		return writeRenderArtifacts(outputPath, opts.Format, artifacts)
	}
	out, err := renderUsecase.Render(context.Background(), input, opts)
	if err != nil {
		logger.ERROR(ICRRR003, "render failed", map[string]any{"format": opts.Format, "error": err})
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		logger.ERROR(ICRRR004, "write output failed", map[string]any{"output": outputPath, "error": err})
		return fmt.Errorf("write output file: %w", err)
	}
	logger.INFO(ICRRR005, "generated", map[string]any{"format": opts.Format, "output": outputPath})
	return nil
}

func writeRenderArtifacts(outputPath string, format entity.Format, artifacts []entity.RenderArtifact) error {
	if len(artifacts) == 0 {
		return fmt.Errorf("render produced no artifacts")
	}
	paths := make([]string, len(artifacts))
	if len(artifacts) == 1 {
		paths[0] = outputPath
	} else {
		ext := filepath.Ext(outputPath)
		stem := strings.TrimSuffix(filepath.Base(outputPath), ext)
		if ext == "" {
			ext = ".svg"
		}
		seen := map[string]string{}
		for index, artifact := range artifacts {
			id := safeRenderArtifactID(artifact.ID)
			if id == "" {
				id = fmt.Sprintf("frame-%d", index+1)
			}
			path := filepath.Join(filepath.Dir(outputPath), stem+"-"+id+ext)
			collisionKey := strings.ToLower(filepath.Clean(path))
			if prior := seen[collisionKey]; prior != "" {
				return fmt.Errorf("frame IDs %q and %q resolve to the same output %s", prior, artifact.ID, path)
			}
			seen[collisionKey] = artifact.ID
			paths[index] = path
		}
	}
	for index, artifact := range artifacts {
		if err := os.WriteFile(paths[index], artifact.Data, 0o644); err != nil {
			logger.ERROR(ICRRR004, "write output failed", map[string]any{"output": paths[index], "error": err})
			return fmt.Errorf("write output file: %w", err)
		}
		logger.INFO(ICRRR005, "generated", map[string]any{"format": format, "output": paths[index], "frame": artifact.ID})
	}
	return nil
}

func safeRenderArtifactID(id string) string {
	var builder strings.Builder
	previousDash := false
	for _, char := range strings.TrimSpace(id) {
		valid := char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' || char == '_' || char == '-'
		if valid {
			builder.WriteRune(char)
			previousDash = false
			continue
		}
		if !previousDash {
			builder.WriteByte('-')
			previousDash = true
		}
	}
	return strings.Trim(builder.String(), "-")
}

func normalizeRenderFormat(format string) string {
	format = strings.TrimSpace(strings.ToLower(format))
	if format == "" {
		logger.DEBUG(ICRNRF001, "branch default format")
		return "svg"
	}
	logger.DEBUG(ICRNRF002, "branch explicit format", map[string]any{"format": format})
	return format
}

func defaultRenderOutput(format string) string {
	switch normalizeRenderFormat(format) {
	case "svg":
		logger.DEBUG(ICRDRO001, "branch svg output")
		return "output.svg"
	case "pptx":
		logger.DEBUG(ICRDRO002, "branch pptx output")
		return "output.pptx"
	default:
		return "output"
	}
}
