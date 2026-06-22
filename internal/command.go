package command

import (
	"os"

	"github.com/ryo-arima/xaligo/internal/controller"
	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
)

var (
	logger      = share.DefaultLogger()
	ICNRC001    = share.NewMCode("ICNRC-001", "New root command start")
	ICNRCWUC001 = share.NewMCode("ICNRCWUC-001", "New root command with use case nil use case branch")
	ICNRCWUC002 = share.NewMCode("ICNRCWUC-002", "New root command with use case configured")
	ICE001      = share.NewMCode("ICE-001", "Execute failed")
)

func NewRootCmd() *cobra.Command {
	logger.DEBUG(ICNRC001, "start")
	return NewRootCmdWithUseCase(usecase.New())
}

func NewRootCmdWithUseCase(uc usecase.API) *cobra.Command {
	if uc == nil {
		logger.DEBUG(ICNRCWUC001, "branch nil use case")
		uc = usecase.New()
	} else {
		logger.DEBUG(ICNRCWUC002, "branch configured use case")
	}
	var root = &cobra.Command{
		Use:   "xaligo",
		Short: "Vue-like DSL to Excalidraw layout generator",
		Long:  "xaligo renders a Vue-like layout DSL into an Excalidraw JSON file.",
	}

	root.AddCommand(controller.InitRenderCmdWithUseCase(uc))
	root.AddCommand(controller.InitValidateCmdWithUseCase(uc))
	root.AddCommand(controller.InitServeCmdWithUseCase(uc))
	root.AddCommand(controller.InitInitCmd())
	root.AddCommand(controller.InitVersionCmd())
	root.AddCommand(controller.InitAddCmd())
	root.AddCommand(controller.InitGenerateCmd())
	return root
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		logger.ERROR(ICE001, "execute failed", map[string]any{"error": err})
		os.Exit(1)
	}
}
