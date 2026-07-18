package command

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/controller"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
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
	pdfRepository := repository.NewPDFRepository()
	spreadsheetRepository := repository.NewSpreadsheetRepository()

	renderUsecase := usecase.NewRenderUsecase(
		excalidrawRepository,
		xaligoRepository,
		powerpointRepository,
		isoflowRepository,
		svgRepository,
		xyFlowRepository,
		pdfRepository,
		spreadsheetRepository,
	)
	sceneIOUsecase := usecase.NewSceneIOUsecase(excalidrawRepository)
	catalogUsecase := usecase.NewCatalogUsecase(xaligoRepository)
	exportUsecase := usecase.NewExportUsecase(powerpointRepository)
	diagnosticsUsecase := usecase.NewDiagnosticsUsecase()
	elementUsecase := usecase.NewElementUsecase()
	themeUsecase := usecase.NewThemeUsecase()
	diffUsecase := usecase.NewDiffUsecase(xaligoRepository, excalidrawRepository, svgRepository)

	addController := controller.NewAddController(cfg, sceneIOUsecase, catalogUsecase, elementUsecase)
	generateController := controller.NewGenerateController(renderUsecase, exportUsecase)
	renderController := controller.NewRenderController(cfg, renderUsecase, catalogUsecase, sceneIOUsecase, themeUsecase, elementUsecase)
	validateController := controller.NewValidateController(diagnosticsUsecase)
	serveController := controller.NewServeController(renderUsecase)
	initController := controller.NewInitController()
	versionController := controller.NewVersionController()
	diffController := controller.NewDiffController(diffUsecase)

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
	root.AddCommand(diffController.Command())
	return root
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		logger.ERROR(ICE001, "execute failed", map[string]any{"error": err})
		os.Exit(1)
	}
}
