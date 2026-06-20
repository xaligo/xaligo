package controller

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	xaligo "github.com/ryo-arima/xaligo"
	"github.com/ryo-arima/xaligo/preview"
	"github.com/spf13/cobra"
)

func InitServeCmd() *cobra.Command {
	var address, mode, theme string
	var poll time.Duration
	cmd := &cobra.Command{
		Use:   "serve <input.xal>",
		Short: "Serve a live SVG preview and reload it when the source changes",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServe(cmd.Context(), ServeOptions{
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

type ServeOptions struct {
	InputPath    string
	Address      string
	Mode         string
	Theme        string
	PollInterval time.Duration
}

func RunServe(ctx context.Context, opts ServeOptions) error {
	server, err := preview.New(opts.InputPath, preview.Options{
		Render:       xaligo.RenderOptions{Mode: xaligo.Mode(opts.Mode), Format: xaligo.FormatSVG, Theme: opts.Theme},
		PollInterval: opts.PollInterval,
	})
	if err != nil {
		return err
	}
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()
	address := opts.Address
	if address == "" {
		address = "127.0.0.1:8080"
	}
	fmt.Printf("preview: http://%s\n", address)
	fmt.Printf("watching: %s\n", opts.InputPath)
	return server.Run(ctx, address)
}
