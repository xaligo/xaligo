package controller

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
)

func InitRenderCmd() *cobra.Command {
	return InitRenderCmdWithUseCase(nil)
}

func InitRenderCmdWithUseCase(uc usecase.API) *cobra.Command {
	uc = defaultUseCase(uc)
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
			format = normalizeRenderFormat(format)
			if output == "" {
				output = defaultRenderOutput(format)
			}
			if noCompression {
				compression = false
			}
			return RunRenderFormatWithUseCase(uc, RenderOptions{
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
	cmd.Flags().StringVar(&exporterWASM, "pptx-exporter-wasm", "", "path to the WASM PPTX exporter (default: external/wasm/pptx_exporter.wasm or XALIGO_PPTX_EXPORTER_WASM)")
	cmd.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	cmd.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	return cmd
}

// abbrevMap is an optional catalog-ID → abbreviation override derived from services.csv.
// Pass nil to use only the built-in abbreviation table.
func RunRender(inputPath, outputPath string, abbrevMap map[int]string) error {
	return runRenderExcalidraw(defaultUseCase(nil), inputPath, outputPath, abbrevMap, string(usecase.ModeStandard), entity.ThemeLight)
}

// RunRenderFormat renders a .xal file into the requested output format. It is
// the public controller entry point for format-based rendering.
func RunRenderFormat(opts RenderOptions) error {
	return RunRenderFormatWithUseCase(nil, opts)
}

func RunRenderFormatWithUseCase(uc usecase.API, opts RenderOptions) error {
	uc = defaultUseCase(uc)
	if err := uc.ValidateRenderOptions(usecase.RenderOptions{
		Mode: usecase.Mode(opts.Mode), Format: usecase.Format(opts.Format), Theme: opts.Theme,
		PaperMarginIn: opts.PaperMargin, PaperMarginTopIn: opts.PaperMarginTop, PaperMarginRightIn: opts.PaperMarginRight,
		PaperMarginBottomIn: opts.PaperMarginBottom, PaperMarginLeftIn: opts.PaperMarginLeft,
	}); err != nil {
		return err
	}
	theme, err := entity.NormalizeTheme(opts.Theme)
	if err != nil {
		return err
	}
	switch normalizeRenderFormat(opts.Format) {
	case "excalidraw":
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			return err
		}
		if opts.ServicesFile != "" {
			warnServiceMismatch(opts.InputPath, opts.ServicesFile)
		}
		if err := runRenderExcalidraw(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme); err != nil {
			return err
		}
		if opts.ServicesFile != "" {
			if err := RunAddServiceBatch(opts.OutputPath, opts.ServicesFile); err != nil {
				return err
			}
			return applyThemeFile(opts.OutputPath, theme)
		}
		return nil
	case "svg":
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			return err
		}
		if opts.ServicesFile != "" {
			warnServiceMismatch(opts.InputPath, opts.ServicesFile)
		}
		return runRenderSVG(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme, opts.PxPerInch, opts.ArrowStyle, opts.ArrowStub, opts.ArrowMargin, opts.Paper, opts.Orientation, opts.PaperMargin, opts.PaperMarginTop, opts.PaperMarginRight, opts.PaperMarginBottom, opts.PaperMarginLeft)
	case "pptx":
		return RunGeneratePptxWithUseCase(uc, PptxGenerateOptions{
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
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			return err
		}
		return runRenderXYFlow(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme)
	case "isoflow":
		abbrevMap, err := serviceAbbrevMap(opts.ServicesFile)
		if err != nil {
			return err
		}
		return runRenderIsoflow(uc, opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme)
	default:
		return fmt.Errorf("unknown render format %q; valid: excalidraw, svg, pptx, xyflow, isoflow", opts.Format)
	}
}

func normalizeRenderFormat(format string) string {
	format = strings.TrimSpace(strings.ToLower(format))
	if format == "" {
		return "excalidraw"
	}
	return format
}

func defaultRenderOutput(format string) string {
	switch normalizeRenderFormat(format) {
	case "svg":
		return "output.svg"
	case "pptx":
		return "output.pptx"
	case "xyflow":
		return "output.xyflow.json"
	case "isoflow":
		return "output.isoflow.json"
	default:
		return "output.excalidraw"
	}
}

func runRenderIsoflow(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}
	out, err := uc.RenderIsoflow(context.Background(), input, usecase.RenderOptions{
		Mode: usecase.Mode(mode), Theme: theme, Abbreviations: abbrevMap,
	})
	if err != nil {
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}
	fmt.Printf("generated: %s\n", outputPath)
	return nil
}

func runRenderXYFlow(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}
	out, err := uc.RenderXYFlow(context.Background(), input, usecase.RenderOptions{
		Mode: usecase.Mode(mode), Theme: theme, Abbreviations: abbrevMap,
	})
	if err != nil {
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}
	fmt.Printf("generated: %s\n", outputPath)
	return nil
}

func serviceAbbrevMap(servicesFile string) (map[int]string, error) {
	if servicesFile == "" {
		return nil, nil
	}
	entries, err := repository.ReadServiceList(servicesFile)
	if err != nil {
		return nil, fmt.Errorf("read services %s: %w", servicesFile, err)
	}
	abbrevMap := make(map[int]string, len(entries))
	for _, e := range entries {
		if e.CatalogID > 0 && e.Abbreviation != "" {
			abbrevMap[e.CatalogID] = e.Abbreviation
		}
	}
	return abbrevMap, nil
}

func runRenderExcalidraw(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string) error {
	out, err := buildExcalidrawJSONWithUseCase(uc, inputPath, abbrevMap, mode, theme)
	if err != nil {
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	fmt.Printf("generated: %s\n", outputPath)
	return nil
}

func runRenderSVG(uc usecase.API, inputPath, outputPath string, abbrevMap map[int]string, mode, theme string, pxPerInch float64, arrowStyle string, arrowStub, arrowMargin float64, paper, orientation string, paperMargin, paperMarginTop, paperMarginRight, paperMarginBottom, paperMarginLeft float64) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}
	out, err := uc.RenderSVG(context.Background(), input, usecase.RenderOptions{
		Mode: usecase.Mode(mode), Theme: theme, Abbreviations: abbrevMap, PxPerInch: pxPerInch,
		ArrowStyle: arrowStyle, ArrowStubPx: arrowStub, ArrowMarginPx: arrowMargin,
		PaperSize: paper, Orientation: orientation,
		PaperMarginIn: paperMargin, PaperMarginTopIn: paperMarginTop, PaperMarginRightIn: paperMarginRight,
		PaperMarginBottomIn: paperMarginBottom, PaperMarginLeftIn: paperMarginLeft,
	})
	if err != nil {
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}
	fmt.Printf("generated: %s\n", outputPath)
	return nil
}

func buildExcalidrawJSON(inputPath string, abbrevMap map[int]string, mode, theme string) ([]byte, error) {
	return buildExcalidrawJSONWithUseCase(defaultUseCase(nil), inputPath, abbrevMap, mode, theme)
}

func buildExcalidrawJSONWithUseCase(uc usecase.API, inputPath string, abbrevMap map[int]string, mode, theme string) ([]byte, error) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	return uc.RenderExcalidraw(context.Background(), input, usecase.RenderOptions{Mode: usecase.Mode(mode), Theme: theme, Abbreviations: abbrevMap})
}

func applyThemeFile(path, theme string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read themed output file: %w", err)
	}
	data, err = usecase.ApplyThemeJSON(data, theme)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write themed output file: %w", err)
	}
	return nil
}
