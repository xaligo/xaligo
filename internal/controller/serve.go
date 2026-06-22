package controller

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
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

func InitServeCmd() *cobra.Command {
	logger.DEBUG(ICSISC001, "start")
	return InitServeCmdWithUseCase(nil)
}

func InitServeCmdWithUseCase(uc usecase.API) *cobra.Command {
	logger.DEBUG(ICSISCWUC001, "start")
	if uc == nil {
		uc = usecase.New()
	}
	var address, mode, theme string
	var poll time.Duration
	cmd := &cobra.Command{
		Use:   "serve <input.xal>",
		Short: "Serve a live SVG preview and reload it when the source changes",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServeWithUseCase(uc, cmd.Context(), entity.ControllerServeOptions{
				InputPath: args[0], Address: address, Mode: mode, Theme: theme,
				PollInterval: poll,
			})
		},
	}
	cmd.Flags().StringVar(&address, "address", "127.0.0.1:8080", "HTTP listen address")
	cmd.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	cmd.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	cmd.Flags().DurationVar(&poll, "poll-interval", 500*time.Millisecond, "source file polling interval")
	return cmd
}

func RunServe(ctx context.Context, opts entity.ControllerServeOptions) error {
	logger.DEBUG(ICSRS001, "start", map[string]any{"input": opts.InputPath, "address": opts.Address})
	return RunServeWithUseCase(nil, ctx, opts)
}

func RunServeWithUseCase(uc usecase.API, ctx context.Context, opts entity.ControllerServeOptions) error {
	if uc == nil {
		uc = usecase.New()
	}
	server, err := uc.NewPreviewServer(opts.InputPath, entity.PreviewOptions{
		Render:       entity.RenderOptions{Mode: entity.Mode(opts.Mode), Format: usecase.FormatSVG, Theme: opts.Theme},
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
