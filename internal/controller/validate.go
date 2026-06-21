package controller

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
)

func InitValidateCmd() *cobra.Command {
	return InitValidateCmdWithUseCase(usecase.New(usecase.Dependencies{}))
}

func InitValidateCmdWithUseCase(uc usecase.API) *cobra.Command {
	if uc == nil {
		uc = usecase.New(usecase.Dependencies{})
	}
	cmd := &cobra.Command{
		Use:   "validate <input.xal>",
		Short: "Validate xaligo DSL syntax and layout",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunValidateWithUseCase(uc, args[0], os.Stdout)
		},
	}
	return cmd
}

func RunValidate(inputPath string, stdout io.Writer) error {
	return RunValidateWithUseCase(usecase.New(usecase.Dependencies{}), inputPath, stdout)
}

func RunValidateWithUseCase(uc usecase.API, inputPath string, stdout io.Writer) error {
	if uc == nil {
		uc = usecase.New(usecase.Dependencies{})
	}
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("open input file: %w", err)
	}
	if err := uc.Validate(context.Background(), input); err != nil {
		return err
	}
	if stdout != nil {
		_, _ = fmt.Fprintf(stdout, "valid: %s\n", inputPath)
	}
	return nil
}
