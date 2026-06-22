package controller

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
)

var (
	ICVALIDATEIVC001    = share.NewMCode("ICVALIDATEIVC-001", "Init validate command start")
	ICVALIDATEIVCWUC001 = share.NewMCode("ICVALIDATEIVCWUC-001", "Init validate command with use case start")
	ICVALIDATERV001     = share.NewMCode("ICVALIDATERV-001", "Run validate start")
	ICVALIDATERVWUC001  = share.NewMCode("ICVALIDATERVWUC-001", "Run validate with use case read input failed")
	ICVALIDATERVWUC002  = share.NewMCode("ICVALIDATERVWUC-002", "Run validate with use case validation failed")
	ICVALIDATERVWUC003  = share.NewMCode("ICVALIDATERVWUC-003", "Run validate with use case stdout branch")
	ICVALIDATERVWUC004  = share.NewMCode("ICVALIDATERVWUC-004", "Run validate with use case nil stdout branch")
)

func InitValidateCmd() *cobra.Command {
	logger.DEBUG(ICVALIDATEIVC001, "start")
	return InitValidateCmdWithUseCase(nil)
}

func InitValidateCmdWithUseCase(uc usecase.API) *cobra.Command {
	logger.DEBUG(ICVALIDATEIVCWUC001, "start")
	if uc == nil {
		uc = usecase.New()
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
	logger.DEBUG(ICVALIDATERV001, "start", map[string]any{"input": inputPath})
	return RunValidateWithUseCase(nil, inputPath, stdout)
}

func RunValidateWithUseCase(uc usecase.API, inputPath string, stdout io.Writer) error {
	if uc == nil {
		uc = usecase.New()
	}
	input, err := os.ReadFile(inputPath)
	if err != nil {
		logger.ERROR(ICVALIDATERVWUC001, "read input failed", map[string]any{"input": inputPath, "error": err})
		return fmt.Errorf("open input file: %w", err)
	}
	if err := uc.Validate(context.Background(), input); err != nil {
		logger.ERROR(ICVALIDATERVWUC002, "validation failed", map[string]any{"input": inputPath, "error": err})
		return err
	}
	if stdout != nil {
		logger.DEBUG(ICVALIDATERVWUC003, "branch stdout provided")
		logger.INFO(ICVALIDATERVWUC003, "valid", map[string]any{"input": inputPath})
	} else {
		logger.DEBUG(ICVALIDATERVWUC004, "branch nil stdout")
	}
	return nil
}
