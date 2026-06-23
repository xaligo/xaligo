package controller

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
)

var (
	logger       = share.DefaultLogger()
	ICRRRFWUC001 = share.NewMCode("ICRRRFWUC-001", "Run render format with use case start")
	ICRRRFWUC002 = share.NewMCode("ICRRRFWUC-002", "Run render format with use case validate render options failed")
	ICRRRFWUC003 = share.NewMCode("ICRRRFWUC-003", "Run render format with use case normalize theme failed")
	ICRRRFWUC004 = share.NewMCode("ICRRRFWUC-004", "Run render format with use case excalidraw branch")
	ICRRRFWUC005 = share.NewMCode("ICRRRFWUC-005", "Run render format with use case excalidraw service abbreviations failed")
	ICRRRFWUC006 = share.NewMCode("ICRRRFWUC-006", "Run render format with use case excalidraw services file branch")
	ICRRRFWUC007 = share.NewMCode("ICRRRFWUC-007", "Run render format with use case excalidraw render failed")
	ICRRRFWUC008 = share.NewMCode("ICRRRFWUC-008", "Run render format with use case add service batch branch")
	ICRRRFWUC009 = share.NewMCode("ICRRRFWUC-009", "Run render format with use case add service batch failed")
	ICRRRFWUC010 = share.NewMCode("ICRRRFWUC-010", "Run render format with use case excalidraw without services branch")
	ICRRRFWUC011 = share.NewMCode("ICRRRFWUC-011", "Run render format with use case svg branch")
	ICRRRFWUC012 = share.NewMCode("ICRRRFWUC-012", "Run render format with use case svg service abbreviations failed")
	ICRRRFWUC013 = share.NewMCode("ICRRRFWUC-013", "Run render format with use case svg services file branch")
	ICRRRFWUC014 = share.NewMCode("ICRRRFWUC-014", "Run render format with use case pptx branch")
	ICRRRFWUC015 = share.NewMCode("ICRRRFWUC-015", "Run render format with use case xyflow branch")
	ICRRRFWUC016 = share.NewMCode("ICRRRFWUC-016", "Run render format with use case xyflow service abbreviations failed")
	ICRRRFWUC017 = share.NewMCode("ICRRRFWUC-017", "Run render format with use case isoflow branch")
	ICRRRFWUC018 = share.NewMCode("ICRRRFWUC-018", "Run render format with use case isoflow service abbreviations failed")
	ICRRRFWUC019 = share.NewMCode("ICRRRFWUC-019", "Run render format with use case unknown format branch")
	ICRRRFWUC020 = share.NewMCode("ICRRRFWUC-020", "Run render format with use case apply theme failed")
	ICRRRFWUC021 = share.NewMCode("ICRRRFWUC-021", "Run render format with use case apply theme completed")
	ICRRRFWUC022 = share.NewMCode("ICRRRFWUC-022", "Run render format with use case SVG render failed")
	ICRRRFWUC023 = share.NewMCode("ICRRRFWUC-023", "Run render format with use case PPTX render failed")
	ICRRRFWUC024 = share.NewMCode("ICRRRFWUC-024", "Run render format with use case XYFlow render failed")
	ICRRRFWUC025 = share.NewMCode("ICRRRFWUC-025", "Run render format with use case Isoflow render failed")
	ICRIRC001    = share.NewMCode("ICRIRC-001", "Init render command start")
	ICRIRCWUC001 = share.NewMCode("ICRIRCWUC-001", "Init render command with use case start")
	ICRIRCWUC002 = share.NewMCode("ICRIRCWUC-002", "Init render command run start")
	ICRIRCWUC003 = share.NewMCode("ICRIRCWUC-003", "Init render command default output branch")
	ICRIRCWUC004 = share.NewMCode("ICRIRCWUC-004", "Init render command explicit output branch")
	ICRIRCWUC005 = share.NewMCode("ICRIRCWUC-005", "Init render command no compression branch")
	ICRIRCWUC006 = share.NewMCode("ICRIRCWUC-006", "Init render command compression branch")
	ICRIRCWUC007 = share.NewMCode("ICRIRCWUC-007", "Init render command return command")
	ICRR001      = share.NewMCode("ICRR-001", "Run render start")
	ICRRRF001    = share.NewMCode("ICRRRF-001", "Run render format start")
	ICRNRF001    = share.NewMCode("ICRNRF-001", "Normalize render format default branch")
	ICRNRF002    = share.NewMCode("ICRNRF-002", "Normalize render format explicit branch")
	ICRDRO001    = share.NewMCode("ICRDRO-001", "Default render output SVG branch")
	ICRDRO002    = share.NewMCode("ICRDRO-002", "Default render output PPTX branch")
	ICRDRO003    = share.NewMCode("ICRDRO-003", "Default render output XYFlow branch")
	ICRDRO004    = share.NewMCode("ICRDRO-004", "Default render output Isoflow branch")
	ICRDRO005    = share.NewMCode("ICRDRO-005", "Default render output Excalidraw branch")
	ICRRRI001    = share.NewMCode("ICRRRI-001", "Run render Isoflow read input failed")
	ICRRRI002    = share.NewMCode("ICRRRI-002", "Run render Isoflow render failed")
	ICRRRI003    = share.NewMCode("ICRRRI-003", "Run render Isoflow write output failed")
	ICRRRI004    = share.NewMCode("ICRRRI-004", "Run render Isoflow generated output")
	ICRRXYF001   = share.NewMCode("ICRRXYF-001", "Run render XYFlow read input failed")
	ICRRXYF002   = share.NewMCode("ICRRXYF-002", "Run render XYFlow render failed")
	ICRRXYF003   = share.NewMCode("ICRRXYF-003", "Run render XYFlow write output failed")
	ICRRXYF004   = share.NewMCode("ICRRXYF-004", "Run render XYFlow generated output")
	ICRSAM001    = share.NewMCode("ICRSAM-001", "Service abbreviation map empty services branch")
	ICRSAM002    = share.NewMCode("ICRSAM-002", "Service abbreviation map read services failed")
	ICRSAM003    = share.NewMCode("ICRSAM-003", "Service abbreviation map abbreviation branch")
	ICRSAM004    = share.NewMCode("ICRSAM-004", "Service abbreviation map completed")
	ICRRRE001    = share.NewMCode("ICRRRE-001", "Run render Excalidraw build failed")
	ICRRRE002    = share.NewMCode("ICRRRE-002", "Run render Excalidraw write output failed")
	ICRRRE003    = share.NewMCode("ICRRRE-003", "Run render Excalidraw generated output")
	ICRRRSVG001  = share.NewMCode("ICRRRSVG-001", "Run render SVG read input failed")
	ICRRRSVG002  = share.NewMCode("ICRRRSVG-002", "Run render SVG render failed")
	ICRRRSVG003  = share.NewMCode("ICRRRSVG-003", "Run render SVG write output failed")
	ICRRRSVG004  = share.NewMCode("ICRRRSVG-004", "Run render SVG generated output")
	ICRBEJWUC001 = share.NewMCode("ICRBEJWUC-001", "Build Excalidraw JSON with use case read input failed")
	ICRBEJWUC002 = share.NewMCode("ICRBEJWUC-002", "Build Excalidraw JSON with use case render branch")
	ICRATF001    = share.NewMCode("ICRATF-001", "Apply theme file read failed")
	ICRATF002    = share.NewMCode("ICRATF-002", "Apply theme file apply failed")
	ICRATF003    = share.NewMCode("ICRATF-003", "Apply theme file write failed")
	ICRATF004    = share.NewMCode("ICRATF-004", "Apply theme file completed")
)

func InitRenderCmd() *cobra.Command {
	logger.DEBUG(ICRIRC001, "start")
	return InitRenderCmdWithUseCase(nil)
}

func InitRenderCmdWithUseCase(uc usecase.API) *cobra.Command {
	logger.DEBUG(ICRIRCWUC001, "start")
	if uc == nil {
		uc = usecase.New()
	}
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
	)

	cmd := &cobra.Command{
		Use:   "render <input.xal>",
		Short: "Render xaligo DSL into an output format",
		Args:  cobra.ExactArgs(1),
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
			return RunRenderFormatWithUseCase(uc, entity.ControllerRenderOptions{
				InputPath:         input,
				OutputPath:        output,
				Format:            format,
				ServicesFile:      servicesFile,
				Title:             title,
				Author:            author,
				Company:           company,
				Subject:           subject,
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
				Stdout:            os.Stdout,
				Stderr:            os.Stderr,
			})
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "output file path")
	cmd.Flags().StringVar(&format, "format", "excalidraw", "output format: excalidraw | svg | pptx | xyflow | isoflow")
	cmd.Flags().StringVar(&servicesFile, "services", "", "optional services.csv for icon labels, Excalidraw legend, and PPTX legend slides")
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
	cmd.Flags().StringVar(&exporterWASM, "pptx-exporter-wasm", "", "path to the WASM PPTX exporter (default: external/wasm/xaligo.wasm or XALIGO_PPTX_EXPORTER_WASM)")
	cmd.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	cmd.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	logger.DEBUG(ICRIRCWUC007, "return command")
	return cmd
}

// abbrevMap is an optional catalog-ID → abbreviation override derived from services.csv.
// Pass nil to use only the built-in abbreviation table.
func RunRender(inputPath, outputPath string, abbrevMap map[int]string) error {
	logger.DEBUG(ICRR001, "start", map[string]any{"input": inputPath, "output": outputPath})
	return runRenderExcalidraw(usecase.New(), inputPath, outputPath, abbrevMap, string(usecase.ModeStandard), entity.ThemeLight)
}

// RunRenderFormat renders a .xal file into the requested output format. It is
// the public controller entry point for format-based rendering.
func RunRenderFormat(opts entity.ControllerRenderOptions) error {
	logger.DEBUG(ICRRRF001, "start", map[string]any{"format": opts.Format, "input": opts.InputPath, "output": opts.OutputPath})
	return RunRenderFormatWithUseCase(nil, opts)
}

func RunRenderFormatWithUseCase(uc usecase.API, opts entity.ControllerRenderOptions) error {
	if uc == nil {
		uc = usecase.New()
	}
	logger.DEBUG(ICRRRFWUC001, "start", map[string]any{"format": opts.Format, "input": opts.InputPath, "output": opts.OutputPath})
	if err := uc.ValidateRenderOptions(entity.RenderOptions{
		Mode: entity.Mode(opts.Mode), Format: entity.Format(opts.Format), Theme: opts.Theme,
		PaperMarginIn: opts.PaperMargin, PaperMarginTopIn: opts.PaperMarginTop, PaperMarginRightIn: opts.PaperMarginRight,
		PaperMarginBottomIn: opts.PaperMarginBottom, PaperMarginLeftIn: opts.PaperMarginLeft,
	}); err != nil {
		logger.ERROR(ICRRRFWUC002, "validate render options failed", map[string]any{"error": err})
		return err
	}
	theme, err := entity.NormalizeTheme(opts.Theme)
	if err != nil {
		logger.ERROR(ICRRRFWUC003, "normalize theme failed", map[string]any{"theme": opts.Theme, "error": err})
		return err
	}
	switch normalizeRenderFormat(opts.Format) {
	case "excalidraw":
		logger.DEBUG(ICRRRFWUC004, "branch excalidraw", map[string]any{"services": opts.ServicesFile != ""})
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			logger.ERROR(ICRRRFWUC005, "service abbreviations failed", map[string]any{"format": "excalidraw", "error": err})
			return err
		}
		if opts.ServicesFile != "" {
			logger.DEBUG(ICRRRFWUC006, "branch excalidraw services file", map[string]any{"servicesFile": opts.ServicesFile})
			warnServiceMismatch(opts.InputPath, opts.ServicesFile)
		}
		if err := runRenderExcalidraw(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme); err != nil {
			logger.ERROR(ICRRRFWUC007, "render excalidraw failed", map[string]any{"error": err})
			return err
		}
		if opts.ServicesFile != "" {
			logger.DEBUG(ICRRRFWUC008, "branch add service batch", map[string]any{"servicesFile": opts.ServicesFile})
			if err := RunAddServiceBatch(opts.OutputPath, opts.ServicesFile); err != nil {
				logger.ERROR(ICRRRFWUC009, "add service batch failed", map[string]any{"error": err})
				return err
			}
			return applyThemeFile(opts.OutputPath, theme)
		}
		logger.DEBUG(ICRRRFWUC010, "branch excalidraw without services")
		return nil
	case "svg":
		logger.DEBUG(ICRRRFWUC011, "branch svg", map[string]any{"services": opts.ServicesFile != ""})
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			logger.ERROR(ICRRRFWUC012, "service abbreviations failed", map[string]any{"format": "svg", "error": err})
			return err
		}
		if opts.ServicesFile != "" {
			logger.DEBUG(ICRRRFWUC013, "branch svg services file", map[string]any{"servicesFile": opts.ServicesFile})
			warnServiceMismatch(opts.InputPath, opts.ServicesFile)
		}
		return runRenderSVG(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme, opts.PxPerInch, opts.ArrowStyle, opts.ArrowStub, opts.ArrowMargin, opts.Paper, opts.Orientation, opts.PaperMargin, opts.PaperMarginTop, opts.PaperMarginRight, opts.PaperMarginBottom, opts.PaperMarginLeft)
	case "pptx":
		logger.DEBUG(ICRRRFWUC014, "branch pptx")
		return RunGeneratePptxWithUseCase(uc, entity.ControllerPptxGenerateOptions{
			XalPath:           opts.InputPath,
			Output:            opts.OutputPath,
			ServicesFile:      opts.ServicesFile,
			Title:             opts.Title,
			Author:            opts.Author,
			Company:           opts.Company,
			Subject:           opts.Subject,
			Compression:       opts.Compression,
			PxPerInch:         opts.PxPerInch,
			ArrowStyle:        opts.ArrowStyle,
			ArrowStub:         opts.ArrowStub,
			ArrowMargin:       opts.ArrowMargin,
			Paper:             opts.Paper,
			Orientation:       opts.Orientation,
			PaperMargin:       opts.PaperMargin,
			PaperMarginTop:    opts.PaperMarginTop,
			PaperMarginRight:  opts.PaperMarginRight,
			PaperMarginBottom: opts.PaperMarginBottom,
			PaperMarginLeft:   opts.PaperMarginLeft,
			ExporterWASM:      opts.ExporterWASM,
			Theme:             theme,
			Mode:              opts.Mode,
			Stdout:            opts.Stdout,
			Stderr:            opts.Stderr,
		})
	case "xyflow":
		logger.DEBUG(ICRRRFWUC015, "branch xyflow", map[string]any{"services": opts.ServicesFile != ""})
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			logger.ERROR(ICRRRFWUC016, "service abbreviations failed", map[string]any{"format": "xyflow", "error": err})
			return err
		}
		return runRenderXYFlow(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme)
	case "isoflow":
		logger.DEBUG(ICRRRFWUC017, "branch isoflow", map[string]any{"services": opts.ServicesFile != ""})
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			logger.ERROR(ICRRRFWUC018, "service abbreviations failed", map[string]any{"format": "isoflow", "error": err})
			return err
		}
		return runRenderIsoflow(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme)
	default:
		logger.ERROR(ICRRRFWUC019, "branch unknown format", map[string]any{"format": opts.Format})
		return fmt.Errorf("unknown render format %q; valid: excalidraw, svg, pptx, xyflow, isoflow", opts.Format)
	}
}

func normalizeRenderFormat(format string) string {
	format = strings.TrimSpace(strings.ToLower(format))
	if format == "" {
		logger.DEBUG(ICRNRF001, "branch default format")
		return "excalidraw"
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
	case "xyflow":
		logger.DEBUG(ICRDRO003, "branch xyflow output")
		return "output.xyflow.json"
	case "isoflow":
		logger.DEBUG(ICRDRO004, "branch isoflow output")
		return "output.isoflow.json"
	default:
		logger.DEBUG(ICRDRO005, "branch excalidraw output")
		return "output.excalidraw"
	}
}

func runRenderIsoflow(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		logger.ERROR(ICRRRI001, "read input failed", map[string]any{"input": inputPath, "error": err})
		return fmt.Errorf("read input file: %w", err)
	}
	out, err := uc.RenderIsoflow(context.Background(), input, entity.RenderOptions{
		Mode: entity.Mode(mode), Theme: theme, Abbreviations: abbrevMap,
	})
	if err != nil {
		logger.ERROR(ICRRRI002, "render failed", map[string]any{"error": err})
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		logger.ERROR(ICRRRI003, "write output failed", map[string]any{"output": outputPath, "error": err})
		return fmt.Errorf("write output file: %w", err)
	}
	logger.INFO(ICRRRI004, "generated", map[string]any{"output": outputPath})
	return nil
}

func runRenderXYFlow(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		logger.ERROR(ICRRXYF001, "read input failed", map[string]any{"input": inputPath, "error": err})
		return fmt.Errorf("read input file: %w", err)
	}
	out, err := uc.RenderXYFlow(context.Background(), input, entity.RenderOptions{
		Mode: entity.Mode(mode), Theme: theme, Abbreviations: abbrevMap,
	})
	if err != nil {
		logger.ERROR(ICRRXYF002, "render failed", map[string]any{"error": err})
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		logger.ERROR(ICRRXYF003, "write output failed", map[string]any{"output": outputPath, "error": err})
		return fmt.Errorf("write output file: %w", err)
	}
	logger.INFO(ICRRXYF004, "generated", map[string]any{"output": outputPath})
	return nil
}

func serviceAbbrevMap(servicesFile string) (map[int]string, error) {
	if servicesFile == "" {
		logger.DEBUG(ICRSAM001, "branch empty services file")
		return nil, nil
	}
	entries, err := repository.ReadServiceList(servicesFile)
	if err != nil {
		logger.ERROR(ICRSAM002, "read services failed", map[string]any{"servicesFile": servicesFile, "error": err})
		return nil, fmt.Errorf("read services %s: %w", servicesFile, err)
	}
	abbrevMap := make(map[int]string, len(entries))
	for _, e := range entries {
		if e.CatalogID > 0 && e.Abbreviation != "" {
			logger.DEBUG(ICRSAM003, "branch abbreviation", map[string]any{"catalogID": e.CatalogID})
			abbrevMap[e.CatalogID] = e.Abbreviation
		}
	}
	logger.DEBUG(ICRSAM004, "completed", map[string]any{"servicesFile": servicesFile, "entries": len(entries), "abbreviations": len(abbrevMap)})
	return abbrevMap, nil
}

func runRenderExcalidraw(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string) error {
	out, err := buildExcalidrawJSONWithUseCase(uc, inputPath, abbrevMap, mode, theme)
	if err != nil {
		logger.ERROR(ICRRRE001, "build failed", map[string]any{"error": err})
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		logger.ERROR(ICRRRE002, "write output failed", map[string]any{"output": outputPath, "error": err})
		return fmt.Errorf("write output file: %w", err)
	}

	logger.INFO(ICRRRE003, "generated", map[string]any{"output": outputPath})
	return nil
}

func runRenderSVG(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string, pxPerInch float64, arrowStyle string, arrowStub, arrowMargin float64, paper, orientation string, paperMargin, paperMarginTop, paperMarginRight, paperMarginBottom, paperMarginLeft float64) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		logger.ERROR(ICRRRSVG001, "read input failed", map[string]any{"input": inputPath, "error": err})
		return fmt.Errorf("read input file: %w", err)
	}
	out, err := uc.RenderSVG(context.Background(), input, entity.RenderOptions{
		Mode: entity.Mode(mode), Theme: theme, Abbreviations: abbrevMap, PxPerInch: pxPerInch,
		ArrowStyle: arrowStyle, ArrowStubPx: arrowStub, ArrowMarginPx: arrowMargin,
		PaperSize: paper, Orientation: orientation,
		PaperMarginIn: paperMargin, PaperMarginTopIn: paperMarginTop, PaperMarginRightIn: paperMarginRight,
		PaperMarginBottomIn: paperMarginBottom, PaperMarginLeftIn: paperMarginLeft,
	})
	if err != nil {
		logger.ERROR(ICRRRSVG002, "render failed", map[string]any{"error": err})
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		logger.ERROR(ICRRRSVG003, "write output failed", map[string]any{"output": outputPath, "error": err})
		return fmt.Errorf("write output file: %w", err)
	}
	logger.INFO(ICRRRSVG004, "generated", map[string]any{"output": outputPath})
	return nil
}

func buildExcalidrawJSONWithUseCase(uc usecase.API, inputPath string, abbrevMap map[int]string, mode, theme string) ([]byte, error) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		logger.ERROR(ICRBEJWUC001, "read input failed", map[string]any{"input": inputPath, "error": err})
		return nil, fmt.Errorf("read input file: %w", err)
	}
	logger.DEBUG(ICRBEJWUC002, "branch render excalidraw", map[string]any{"input": inputPath})
	return uc.RenderExcalidraw(context.Background(), input, entity.RenderOptions{Mode: entity.Mode(mode), Theme: theme, Abbreviations: abbrevMap})
}

func applyThemeFile(path, theme string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.ERROR(ICRATF001, "read failed", map[string]any{"path": path, "error": err})
		return fmt.Errorf("read themed output file: %w", err)
	}
	data, err = usecase.ApplyThemeJSON(data, theme)
	if err != nil {
		logger.ERROR(ICRATF002, "apply failed", map[string]any{"theme": theme, "error": err})
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		logger.ERROR(ICRATF003, "write failed", map[string]any{"path": path, "error": err})
		return fmt.Errorf("write themed output file: %w", err)
	}
	logger.DEBUG(ICRATF004, "completed", map[string]any{"path": path, "theme": theme})
	return nil
}
