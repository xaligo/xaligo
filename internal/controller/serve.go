package controller

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
)

var (
	ICSISC001    = share.NewMCode("ICSISC-001", "Init serve command start")
	ICSISCWUC001 = share.NewMCode("ICSISCWUC-001", "Init serve command with use case start")
	ICSRS001     = share.NewMCode("ICSRS-001", "Run serve start")
	ICSRSWUC001  = share.NewMCode("ICSRSWUC-001", "Run serve with use case create preview server failed")
	ICSRSWUC002  = share.NewMCode("ICSRSWUC-002", "Run serve with use case nil context branch")
	ICSRSWUC003  = share.NewMCode("ICSRSWUC-003", "Run serve with use case default address branch")
	ICSRSWUC004  = share.NewMCode("ICSRSWUC-004", "Run serve with use case explicit address branch")
	ICSRSWUC005  = share.NewMCode("ICSRSWUC-005", "Run serve with use case preview URL")
	ICSRSWUC006  = share.NewMCode("ICSRSWUC-006", "Run serve with use case watching source")
)

type ServeController interface {
	Command() *cobra.Command
	Run(ctx context.Context, opts entity.ControllerServeOptions) error
}

type serveController struct {
	renderUsecase usecase.RenderUsecase
}

func NewServeController(renderUsecase usecase.RenderUsecase) ServeController {
	return &serveController{renderUsecase: renderUsecase}
}

func (rcvr *serveController) Command() *cobra.Command {
	logger.DEBUG(ICSISCWUC001, "start")
	var address, mode, theme, paper, orientation string
	var poll time.Duration
	cmd := &cobra.Command{
		Use:   "serve <input.xal|input.md>",
		Short: "Serve a live preview and reload it when the source changes",
		Long: "Serve a live-reloading preview of a .xal or Markdown (.md) source file over\n" +
			"HTTP.\n\n" +
			"The source file is polled at --poll-interval; when it changes, the preview\n" +
			"is re-rendered and the open browser page reloads automatically.\n\n" +
			"A .xal source renders every frame onto one combined SVG canvas, the same as\n" +
			"'render --combine-frames'. A .md source renders every fenced ```xal\n" +
			"code block to SVG and previews the full Markdown document with the\n" +
			"diagrams embedded inline, the same as 'render markdown'.\n\n" +
			"Use --paper and --orientation to preview how a diagram fits a specific\n" +
			"physical page size.\n\n" +
			"Press Ctrl+C to stop the server.\n\n" +
			"Examples:\n" +
			"  xaligo serve diagram.xal\n" +
			"  xaligo serve diagram.xal --address 0.0.0.0:9090 --mode network --theme dark\n" +
			"  xaligo serve guide.md\n" +
			"  xaligo serve diagram.xal --paper A4 --orientation landscape",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return rcvr.Run(cmd.Context(), entity.ControllerServeOptions{
				InputPath: args[0], Address: address, Mode: mode, Theme: theme,
				PollInterval: poll, Paper: paper, Orientation: orientation,
			})
		},
	}
	cmd.Flags().StringVar(&address, "address", "127.0.0.1:8080", "HTTP listen address")
	cmd.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	cmd.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	cmd.Flags().DurationVar(&poll, "poll-interval", 500*time.Millisecond, "source file polling interval")
	cmd.Flags().StringVar(&paper, "paper", "", "physical paper size to preview: A5 | A4 | A3 | A2 | A1 | Letter | Legal | Tabloid (default: auto-fit)")
	cmd.Flags().StringVar(&orientation, "orientation", "", "page orientation: portrait | landscape (default: auto-fit)")
	return cmd
}

func (rcvr *serveController) Run(ctx context.Context, opts entity.ControllerServeOptions) error {
	logger.DEBUG(ICSRS001, "start", map[string]any{"input": opts.InputPath, "address": opts.Address})
	absInputPath, err := filepath.Abs(opts.InputPath)
	if err != nil {
		return fmt.Errorf("resolve input file path: %w", err)
	}
	kind := entity.PreviewKindSVG
	ext := strings.ToLower(filepath.Ext(opts.InputPath))
	if ext == ".md" || ext == ".markdown" {
		kind = entity.PreviewKindHTML
	}
	server, err := rcvr.renderUsecase.NewPreviewRepository(opts.InputPath, entity.PreviewOptions{
		Kind: kind,
		Render: entity.RenderOptions{
			Mode: entity.Mode(opts.Mode), Format: usecase.FormatSVG, Theme: opts.Theme,
			PaperSize: opts.Paper, Orientation: opts.Orientation,
			Imports: &entity.ImportSource{FS: os.DirFS(filepath.Dir(absInputPath))},
		},
		PollInterval: opts.PollInterval,
	})
	if err != nil {
		logger.ERROR(ICSRSWUC001, "create preview server failed", map[string]any{"input": opts.InputPath, "error": err})
		return err
	}
	if ctx == nil {
		logger.DEBUG(ICSRSWUC002, "branch nil context")
		ctx = context.Background()
	}
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	address := opts.Address
	if address == "" {
		logger.DEBUG(ICSRSWUC003, "branch default address")
		address = "127.0.0.1:8080"
	} else {
		logger.DEBUG(ICSRSWUC004, "branch explicit address", map[string]any{"address": address})
	}
	logger.INFO(ICSRSWUC005, "preview", map[string]any{"url": "http://" + address})
	logger.INFO(ICSRSWUC006, "watching", map[string]any{"input": opts.InputPath})
	return server.Run(ctx, address)
}
