package controller

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
)

var (
	ICGIGC001  = share.NewMCode("ICGIGC-001", "Init generate command start")
	ICGIGXC001 = share.NewMCode("ICGIGXC-001", "Init generate XAL command start")
	ICGRG001   = share.NewMCode("ICGRG-001", "Run generate unknown paper size branch")
	ICGRG002   = share.NewMCode("ICGRG-002", "Run generate landscape branch")
	ICGRG003   = share.NewMCode("ICGRG-003", "Run generate invalid orientation branch")
	ICGRG004   = share.NewMCode("ICGRG-004", "Run generate invalid AZ layout branch")
	ICGRG005   = share.NewMCode("ICGRG-005", "Run generate invalid spacing branch")
	ICGRG006   = share.NewMCode("ICGRG-006", "Run generate invalid start branch")
	ICGRG007   = share.NewMCode("ICGRG-007", "Run generate value out of range branch")
	ICGRG008   = share.NewMCode("ICGRG-008", "Run generate write output failed")
	ICGRG009   = share.NewMCode("ICGRG-009", "Run generate generated output")
)

var paperSizes = map[string][2]int{
	"A5":      {559, 794},
	"A4":      {794, 1122},
	"A3":      {1122, 1587},
	"A2":      {1587, 2245},
	"A1":      {2245, 3179},
	"Letter":  {816, 1056},
	"Legal":   {816, 1344},
	"Tabloid": {1056, 1632},
}

type GenerateController interface {
	Command() *cobra.Command
	RunPptx(opts entity.ControllerPptxGenerateOptions) error
}

type generateController struct {
	renderUsecase usecase.RenderUsecase
	exportUsecase usecase.ExportUsecase
}

func NewGenerateController(renderUsecase usecase.RenderUsecase, exportUsecase usecase.ExportUsecase) GenerateController {
	return &generateController{renderUsecase: renderUsecase, exportUsecase: exportUsecase}
}

// Command returns the `xaligo generate` parent command with subcommands:
//   - xaligo generate xal … generate an AWS hierarchy .xal
//
// Format conversion belongs to `xaligo render --format ...`.
func (rcvr *generateController) Command() *cobra.Command {
	logger.DEBUG(ICGIGC001, "start")
	parent := &cobra.Command{
		Use:   "generate",
		Short: "Generate source files",
		Long: `Generate starter source files.

Currently provides one subcommand:
  xal   generate a .xal file with an AWS infrastructure hierarchy

Format conversion (Excalidraw, SVG, PPTX, PDF, Excel, XYFlow, Isoflow) is a
separate concern; use 'xaligo render' on the generated .xal file for that.`,
	}
	parent.AddCommand(initGenerateXalCmd())
	return parent
}

// ── xaligo generate xal ──────────────────────────────────────────────────────

func initGenerateXalCmd() *cobra.Command {
	logger.DEBUG(ICGIGXC001, "start")
	var (
		nClouds     int
		nAccounts   int
		nRegions    int
		nAZs        int
		azLayout    string
		nSubnets    int
		spacingMode string
		startMode   string
		paper       string
		orientation string
		output      string
	)

	cmd := &cobra.Command{
		Use:   "xal",
		Short: "Generate a .xal file with AWS infrastructure hierarchy",
		Long: `Generate a .xal file for a standard AWS architecture usecase.

The hierarchy is: Cloud > Account > Region > VPC > AZ > Subnet.
Subnets alternate between public and private (1st=public, 2nd=private, ...).

Parameters correspond to options in generate_aws_frames.py:
  --az-layout   grid | staggered (default: grid)
  --spacing     vertical | horizontal | both (default: both)
  --start       top (vertical stack) | left (horizontal side-by-side) (default: top)
  --paper       A5 | A4 | A3 | A2 | A1 | Letter | Legal | Tabloid (default: A4)
  --orientation portrait | landscape (default: landscape)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunGenerate(
				nClouds, nAccounts, nRegions, nAZs,
				azLayout, nSubnets, spacingMode, startMode,
				paper, orientation, output,
			)
		},
	}

	cmd.Flags().IntVar(&nClouds, "clouds", 1, "number of AWS clouds (1–2)")
	cmd.Flags().IntVar(&nAccounts, "accounts", 1, "number of AWS accounts per cloud (1–3)")
	cmd.Flags().IntVar(&nRegions, "regions", 1, "number of regions per account (1–2)")
	cmd.Flags().IntVar(&nAZs, "azs", 2, "number of availability zones per region (1–3)")
	cmd.Flags().StringVar(&azLayout, "az-layout", "grid", "AZ layout: grid or staggered")
	cmd.Flags().IntVar(&nSubnets, "subnets", 2, "number of subnets per AZ (1–4)")
	cmd.Flags().StringVar(&spacingMode, "spacing", "both", "spacing mode: vertical | horizontal | both")
	cmd.Flags().StringVar(&startMode, "start", "top", "start mode: top (vertical) | left (horizontal)")
	cmd.Flags().StringVar(&paper, "paper", "A4", "paper size: A5 A4 A3 A2 A1 Letter Legal Tabloid")
	cmd.Flags().StringVar(&orientation, "orientation", "landscape", "page orientation: portrait | landscape")
	cmd.Flags().StringVarP(&output, "output", "o", "", "output .xal file path")

	_ = cmd.MarkFlagRequired("output")

	return cmd
}

// ── RunGenerate ──────────────────────────────────────────────────────────────

// RunGenerate validates parameters and writes the generated .xal to output.
func RunGenerate(
	nClouds, nAccounts, nRegions, nAZs int,
	azLayout string, nSubnets int,
	spacingMode, startMode, paper, orientation, output string,
) error {
	// ── validate ────────────────────────────────────────────────────────────
	size, ok := paperSizes[paper]
	if !ok {
		logger.ERROR(ICGRG001, "branch unknown paper size", map[string]any{"paper": paper})
		return fmt.Errorf("unknown paper size %q; valid: A5 A4 A3 A2 A1 Letter Legal Tabloid", paper)
	}
	W, H := size[0], size[1]
	if strings.EqualFold(orientation, "landscape") {
		logger.DEBUG(ICGRG002, "branch landscape")
		W, H = H, W
	} else if !strings.EqualFold(orientation, "portrait") {
		logger.ERROR(ICGRG003, "branch invalid orientation", map[string]any{"orientation": orientation})
		return fmt.Errorf("orientation must be portrait or landscape")
	}
	if azLayout != "grid" && azLayout != "staggered" {
		logger.ERROR(ICGRG004, "branch invalid AZ layout", map[string]any{"azLayout": azLayout})
		return fmt.Errorf("az-layout must be grid or staggered")
	}
	if spacingMode != "vertical" && spacingMode != "horizontal" && spacingMode != "both" {
		logger.ERROR(ICGRG005, "branch invalid spacing", map[string]any{"spacing": spacingMode})
		return fmt.Errorf("spacing must be vertical, horizontal, or both")
	}
	if startMode != "top" && startMode != "left" {
		logger.ERROR(ICGRG006, "branch invalid start", map[string]any{"start": startMode})
		return fmt.Errorf("start must be top or left")
	}
	for _, pair := range [][2]int{{nClouds, 2}, {nAccounts, 3}, {nRegions, 2}, {nAZs, 3}, {nSubnets, 4}} {
		if pair[0] < 1 || pair[0] > pair[1] {
			logger.ERROR(ICGRG007, "branch value out of range", map[string]any{"value": pair[0], "max": pair[1]})
			return fmt.Errorf("value %d out of valid range (1–%d)", pair[0], pair[1])
		}
	}

	// ── generate ────────────────────────────────────────────────────────────
	xal := buildXAL(W, H, nClouds, nAccounts, nRegions, nAZs, azLayout, nSubnets, spacingMode, startMode)

	if err := os.WriteFile(output, []byte(xal), 0644); err != nil {
		logger.ERROR(ICGRG008, "write output failed", map[string]any{"output": output, "error": err})
		return fmt.Errorf("write output file: %w", err)
	}
	logger.INFO(ICGRG009, "generated", map[string]any{"output": output})
	return nil
}

// RunGeneratePptx builds a resolved Go PPTX plan, then asks the repository layer
// to invoke the WASM exporter that turns the plan into PPTX bytes.
func (rcvr *generateController) RunPptx(opts entity.ControllerPptxGenerateOptions) error {
	return runGeneratePptx(rcvr.renderUsecase, rcvr.exportUsecase, opts)
}

func runGeneratePptx(renderUsecase usecase.RenderUsecase, exportUsecase usecase.ExportUsecase, opts entity.ControllerPptxGenerateOptions) error {
	if opts.XalPath == "" {
		return fmt.Errorf("--xal is required")
	}
	if opts.Output == "" {
		return fmt.Errorf("--output is required")
	}
	if opts.PxPerInch < 0 {
		return fmt.Errorf("--px-per-inch must be positive")
	}
	if opts.PaperMargin < 0 || opts.PaperMarginTop < 0 || opts.PaperMarginRight < 0 || opts.PaperMarginBottom < 0 || opts.PaperMarginLeft < 0 {
		return fmt.Errorf("paper margins must be non-negative")
	}
	planJSON, err := buildPptxPlanJSON(renderUsecase, opts)
	if err != nil {
		return err
	}
	return exportUsecase.ExportPptx(context.Background(), entity.PptxExportOptions{
		PlanJSON:     planJSON,
		Output:       opts.Output,
		Title:        opts.Title,
		Author:       opts.Author,
		Company:      opts.Company,
		Subject:      opts.Subject,
		Compression:  opts.Compression,
		ExporterWASM: opts.ExporterWASM,
		Stdout:       opts.Stdout,
		Stderr:       opts.Stderr,
	})
}

func buildPptxPlanJSON(renderUsecase usecase.RenderUsecase, opts entity.ControllerPptxGenerateOptions) ([]byte, error) {
	if err := renderUsecase.ValidateRenderOptions(entity.RenderOptions{
		Mode: entity.Mode(opts.Mode), Format: usecase.FormatPPTX, Theme: opts.Theme,
		CombineFrames: opts.CombineFrames,
		PxPerInch:     opts.PxPerInch, ArrowStyle: opts.ArrowStyle, ArrowStubPx: opts.ArrowStub, ArrowMarginPx: opts.ArrowMargin,
		PaperSize: opts.Paper, Orientation: opts.Orientation,
		PaperMarginIn: opts.PaperMargin, PaperMarginTopIn: opts.PaperMarginTop, PaperMarginRightIn: opts.PaperMarginRight,
		PaperMarginBottomIn: opts.PaperMarginBottom, PaperMarginLeftIn: opts.PaperMarginLeft,
	}); err != nil {
		return nil, err
	}
	input, err := os.ReadFile(opts.XalPath)
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}
	var servicesCSV []byte
	if opts.ServicesFile != "" {
		servicesCSV, err = os.ReadFile(opts.ServicesFile)
		if err != nil {
			return nil, fmt.Errorf("read services %s: %w", opts.ServicesFile, err)
		}
	}
	return renderUsecase.BuildPPTXPlan(context.Background(), input, entity.RenderOptions{
		Mode:                entity.Mode(opts.Mode),
		Format:              usecase.FormatPPTX,
		Theme:               opts.Theme,
		CombineFrames:       opts.CombineFrames,
		PxPerInch:           opts.PxPerInch,
		ArrowStyle:          opts.ArrowStyle,
		ArrowStubPx:         opts.ArrowStub,
		ArrowMarginPx:       opts.ArrowMargin,
		PaperSize:           opts.Paper,
		Orientation:         opts.Orientation,
		PaperMarginIn:       opts.PaperMargin,
		PaperMarginTopIn:    opts.PaperMarginTop,
		PaperMarginRightIn:  opts.PaperMarginRight,
		PaperMarginBottomIn: opts.PaperMarginBottom,
		PaperMarginLeftIn:   opts.PaperMarginLeft,
		ServicesCSV:         servicesCSV,
	})
}

// ── xal builder ─────────────────────────────────────────────────────────────

type xalBuilder struct {
	sb          strings.Builder
	startMode   string
	spacingMode string
	azLayout    string // "grid" or "staggered"
	usedIDs     map[string]int
}

func (rcvr *xalBuilder) ind(level int) string {
	return strings.Repeat("  ", level)
}

func (rcvr *xalBuilder) group(tag, title string, level int, fn func()) {
	rcvr.sb.WriteString(fmt.Sprintf("%s<%s id=%q title=%q>\n", rcvr.ind(level), tag, rcvr.generatedID(tag, title), title))
	fn()
	rcvr.sb.WriteString(fmt.Sprintf("%s</%s>\n", rcvr.ind(level), tag))
}

func (rcvr *xalBuilder) groupAttrs(tag, title, extraAttrs string, level int, fn func()) {
	if extraAttrs != "" {
		rcvr.sb.WriteString(fmt.Sprintf("%s<%s id=%q title=%q %s>\n", rcvr.ind(level), tag, rcvr.generatedID(tag, title), title, extraAttrs))
	} else {
		rcvr.sb.WriteString(fmt.Sprintf("%s<%s id=%q title=%q>\n", rcvr.ind(level), tag, rcvr.generatedID(tag, title), title))
	}
	fn()
	rcvr.sb.WriteString(fmt.Sprintf("%s</%s>\n", rcvr.ind(level), tag))
}

func (rcvr *xalBuilder) leaf(tag, title string, level int) {
	rcvr.sb.WriteString(fmt.Sprintf("%s<%s id=%q title=%q />\n", rcvr.ind(level), tag, rcvr.generatedID(tag, title), title))
}

func (rcvr *xalBuilder) generatedID(tag, title string) string {
	base := generatedXALID(tag, title)
	if rcvr.usedIDs == nil {
		rcvr.usedIDs = map[string]int{}
	}
	rcvr.usedIDs[base]++
	if rcvr.usedIDs[base] == 1 {
		return base
	}
	return fmt.Sprintf("%s-%d", base, rcvr.usedIDs[base])
}

func generatedXALID(tag, title string) string {
	var slug strings.Builder
	lastDash := false
	for _, r := range strings.ToLower(title) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			slug.WriteRune(r)
			lastDash = false
		default:
			if !lastDash {
				slug.WriteByte('-')
				lastDash = true
			}
		}
	}
	value := strings.Trim(slug.String(), "-")
	if value == "" {
		return tag
	}
	return tag + "-" + value
}

func (rcvr *xalBuilder) spacingClass() string {
	switch rcvr.spacingMode {
	case "vertical":
		return "pt-2 pb-2"
	case "horizontal":
		return "pl-2 pr-2"
	default:
		return "pa-2"
	}
}

func (rcvr *xalBuilder) many(level, n int, fn func(i, level int)) {
	if rcvr.startMode == "left" && n > 1 {
		span := 12 / n
		rcvr.sb.WriteString(fmt.Sprintf("%s<row gap=\"16\">\n", rcvr.ind(level)))
		for i := range n {
			rcvr.sb.WriteString(fmt.Sprintf("%s  <col span=\"%d\" class=%q>\n", rcvr.ind(level), span, rcvr.spacingClass()))
			fn(i, level+2)
			rcvr.sb.WriteString(fmt.Sprintf("%s  </col>\n", rcvr.ind(level)))
		}
		rcvr.sb.WriteString(fmt.Sprintf("%s</row>\n", rcvr.ind(level)))
	} else {
		for i := range n {
			fn(i, level)
		}
	}
}

func buildXAL(W, H, nClouds, nAccounts, nRegions, nAZs int, azLayout string, nSubnets int, spacingMode, startMode string) string {
	b := &xalBuilder{startMode: startMode, spacingMode: spacingMode, azLayout: azLayout}

	b.sb.WriteString(fmt.Sprintf(
		"<!-- xaligo generate xal: clouds=%d accounts=%d regions=%d azs=%d az-layout=%s subnets=%d spacing=%s start=%s -->\n",
		nClouds, nAccounts, nRegions, nAZs, azLayout, nSubnets, spacingMode, startMode,
	))
	if azLayout == "staggered" {
		b.sb.WriteString("<!-- az-layout=staggered: AZs are rendered with depth offset in the excalidraw output -->\n")
	}
	b.sb.WriteString("<xaligo version=\"1\">\n")
	b.sb.WriteString("  <frames>\n")
	b.sb.WriteString(fmt.Sprintf("    <frame id=\"overview\" width=\"%d\" height=\"%d\" class=\"pa-4\">\n", W, H))

	b.many(3, nClouds, func(ci, level int) {
		b.group("aws-cloud", fmt.Sprintf("AWS Cloud %d", ci+1), level, func() {
			b.many(level+1, nAccounts, func(ai, level int) {
				b.group("aws-account", fmt.Sprintf("Account %d", ai+1), level, func() {
					b.many(level+1, nRegions, func(ri, level int) {
						b.group("region", fmt.Sprintf("Region %d", ri+1), level, func() {
							vpcAttr := ""
							if b.azLayout == "staggered" && nAZs >= 2 {
								vpcAttr = `layout="staggered"`
							}
							b.groupAttrs("vpc", fmt.Sprintf("VPC %d", ri+1), vpcAttr, level+1, func() {
								b.many(level+2, nAZs, func(zi, level int) {
									b.group("availability-zone", fmt.Sprintf("AZ %d", zi+1), level, func() {
										b.many(level+1, nSubnets, func(si, level int) {
											if si%2 == 0 {
												b.leaf("public-subnet", fmt.Sprintf("Public Subnet %d", si/2+1), level)
											} else {
												b.leaf("private-subnet", fmt.Sprintf("Private Subnet %d", si/2+1), level)
											}
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})

	b.sb.WriteString("    </frame>\n")
	b.sb.WriteString("  </frames>\n")
	b.sb.WriteString("</xaligo>\n")
	return b.sb.String()
}
