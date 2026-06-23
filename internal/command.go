package command

import (
	"os"

	"github.com/ryo-arima/xaligo/internal/config"
	"github.com/ryo-arima/xaligo/internal/controller"
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/ryo-arima/xaligo/internal/usecase"
	"github.com/spf13/cobra"
)

var (
	logger   = share.DefaultLogger()
	ICNRC001 = share.NewMCode("ICNRC-001", "New root command start")
	ICE001   = share.NewMCode("ICE-001", "Execute failed")
)

func NewRootCmd() *cobra.Command {
	logger.DEBUG(ICNRC001, "start")
	cfg := config.New()

	excalidrawRepository := repository.NewExcalidrawRepository()
	xaligoRepository := repository.NewXaligoRepository()
	powerpointRepository := repository.NewPowerpointRepository()
	isoflowRepository := repository.NewIsoflowRepository()
	svgRepository := repository.NewSVGRepository()
	xyFlowRepository := repository.NewXYFlowRepository()

	xaligoUsecase := usecase.NewXaligoUsecase(
		excalidrawRepository,
		xaligoRepository,
		powerpointRepository,
		isoflowRepository,
		svgRepository,
		xyFlowRepository,
	)

	addController := controller.NewAddController(cfg, xaligoUsecase)
	generateController := controller.NewGenerateController(xaligoUsecase)
	renderController := controller.NewRenderController(cfg, xaligoUsecase)
	validateController := controller.NewValidateController(xaligoUsecase)
	serveController := controller.NewServeController(xaligoUsecase)
	initController := controller.NewInitController()
	versionController := controller.NewVersionController()

	root := &cobra.Command{
		Use:   "xaligo",
		Short: "Vue-like DSL to Excalidraw layout generator",
		Long:  "xaligo renders a Vue-like layout DSL into an Excalidraw JSON file.",
	}

	root.AddCommand(renderController.Command())
	root.AddCommand(validateController.Command())
	root.AddCommand(serveController.Command())
	root.AddCommand(initController.Command())
	root.AddCommand(versionController.Command())
	root.AddCommand(addController.Command())
	root.AddCommand(generateController.Command())
	return root
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		logger.ERROR(ICE001, "execute failed", map[string]any{"error": err})
		os.Exit(1)
	}
}
