package controller

import (
	"context"
	"fmt"
	"io"
	"os"

	xaligoapi "github.com/ryo-arima/xaligo"
	"github.com/spf13/cobra"
)

func InitValidateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate <input.xal>",
		Short: "Validate xaligo DSL syntax and layout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunValidate(args[0], os.Stdout)
		},
	}
	return cmd
}

func RunValidate(inputPath string, stdout io.Writer) error {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("open input file: %w", err)
	}
	if err := xaligoapi.Validate(context.Background(), input); err != nil {
		return err
	}
	if stdout != nil {
		_, _ = fmt.Fprintf(stdout, "valid: %s\n", inputPath)
	}
	return nil
}
