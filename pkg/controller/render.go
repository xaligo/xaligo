package controller

import (
	"context"
	"fmt"
	"os"
	"strings"

	xaligoapi "github.com/ryo-arima/xaligo"
	"github.com/ryo-arima/xaligo/internal/excalidraw"
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/spf13/cobra"
)

func InitRenderCmd() *cobra.Command {
	var (
		output        string
		format        string
		servicesFile  string
		title         string
		author        string
		company       string
		subject       string
		compression   bool
		noCompression bool
		pxPerInch     float64
		arrowStyle    string
		arrowStub     float64
		arrowMargin   float64
		paper         string
		orientation   string
		exporterWASM  string
		theme         string
		mode          string
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
			return RunRenderFormat(RenderOptions{
				InputPath:    input,
				OutputPath:   output,
				Format:       format,
				ServicesFile: servicesFile,
				Title:        title,
				Author:       author,
				Company:      company,
				Subject:      subject,
				Compression:  &compression,
				PxPerInch:    pxPerInch,
				ArrowStyle:   arrowStyle,
				ArrowStub:    arrowStub,
				ArrowMargin:  arrowMargin,
				Paper:        paper,
				Orientation:  orientation,
				ExporterWASM: exporterWASM,
				Theme:        theme,
				Mode:         mode,
				Stdout:       os.Stdout,
				Stderr:       os.Stderr,
			})
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "", "output file path")
	cmd.Flags().StringVar(&format, "format", "excalidraw", "output format: excalidraw | svg | pptx")
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
	cmd.Flags().StringVar(&exporterWASM, "pptx-exporter-wasm", "", "path to the WASM PPTX exporter (default: packages/xaligo/wasm/pptx_exporter.wasm or XALIGO_PPTX_EXPORTER_WASM)")
	cmd.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	cmd.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	return cmd
}

// abbrevMap is an optional catalog-ID → abbreviation override derived from services.csv.
// Pass nil to use only the built-in abbreviation table.
func RunRender(inputPath, outputPath string, abbrevMap map[int]string) error {
	return runRenderExcalidraw(inputPath, outputPath, abbrevMap, string(xaligoapi.ModeStandard), excalidraw.ThemeLight)
}

type RenderOptions struct {
	InputPath    string
	OutputPath   string
	Format       string
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
	ExporterWASM string
	Theme        string
	Mode         string
	Stdout       *os.File
	Stderr       *os.File
}

// RunRenderFormat renders a .xal file into the requested output format. It is
// the public controller entry point for format-based rendering.
func RunRenderFormat(opts RenderOptions) error {
	if err := xaligoapi.ValidateRenderOptions(xaligoapi.RenderOptions{Mode: xaligoapi.Mode(opts.Mode), Format: xaligoapi.Format(opts.Format), Theme: opts.Theme}); err != nil {
		return err
	}
	theme, err := excalidraw.NormalizeTheme(opts.Theme)
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
		if err := runRenderExcalidraw(opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme); err != nil {
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
		return runRenderSVG(opts.InputPath, opts.OutputPath, abbrevMap, opts.Mode, theme, opts.PxPerInch, opts.ArrowStyle, opts.ArrowStub, opts.ArrowMargin, opts.Paper, opts.Orientation)
	case "pptx":
		return RunGeneratePptx(PptxGenerateOptions{
			XalPath:      opts.InputPath,
			Output:       opts.OutputPath,
			ServicesFile: opts.ServicesFile,
			Title:        opts.Title,
			Author:       opts.Author,
			Company:      opts.Company,
			Subject:      opts.Subject,
			Compression:  opts.Compression,
			PxPerInch:    opts.PxPerInch,
			ArrowStyle:   opts.ArrowStyle,
			ArrowStub:    opts.ArrowStub,
			ArrowMargin:  opts.ArrowMargin,
			Paper:        opts.Paper,
			Orientation:  opts.Orientation,
			ExporterWASM: opts.ExporterWASM,
			Theme:        theme,
			Mode:         opts.Mode,
			Stdout:       opts.Stdout,
			Stderr:       opts.Stderr,
		})
	default:
		return fmt.Errorf("unknown render format %q; valid: excalidraw, svg, pptx", opts.Format)
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
	default:
		return "output.excalidraw"
	}
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

func runRenderExcalidraw(inputPath, outputPath string, abbrevMap map[int]string, mode, theme string) error {
	out, err := buildExcalidrawJSON(inputPath, abbrevMap, mode, theme)
	if err != nil {
		return err
	}
	if err := os.WriteFile(outputPath, out, 0644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	fmt.Printf("generated: %s\n", outputPath)
	return nil
}

func runRenderSVG(inputPath, outputPath string, abbrevMap map[int]string, mode, theme string, pxPerInch float64, arrowStyle string, arrowStub, arrowMargin float64, paper, orientation string) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}
	out, err := xaligoapi.RenderSVG(context.Background(), input, xaligoapi.RenderOptions{
		Mode: xaligoapi.Mode(mode), Theme: theme, Abbreviations: abbrevMap, PxPerInch: pxPerInch,
		ArrowStyle: arrowStyle, ArrowStubPx: arrowStub, ArrowMarginPx: arrowMargin,
		PaperSize: paper, Orientation: orientation,
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
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	return xaligoapi.RenderExcalidraw(context.Background(), input, xaligoapi.RenderOptions{Mode: xaligoapi.Mode(mode), Theme: theme, Abbreviations: abbrevMap})
}

func applyThemeFile(path, theme string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read themed output file: %w", err)
	}
	data, err = excalidraw.ApplyThemeJSON(data, theme)
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write themed output file: %w", err)
	}
	return nil
}
