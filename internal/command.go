package command

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/controller"
	"github.com/xaligo/xaligo/internal/core/profiles/builtin"
	"github.com/xaligo/xaligo/internal/repository"
	iconrepository "github.com/xaligo/xaligo/internal/repository/icon"
	projectrepository "github.com/xaligo/xaligo/internal/repository/project"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

var (
	logger   = share.DefaultLogger()
	ICNRC001 = share.NewMCode("ICNRC-001", "New root command start")
	ICE001   = share.NewMCode("ICE-001", "Execute failed")
)

func NewRootCmd() *cobra.Command {
	logger.DEBUG(ICNRC001, "start")
	cfg := config.New()

	sceneRepository := repository.NewSceneRepository()
	xaligoRepository := repository.NewXaligoRepository()
	powerpointRepository := repository.NewPowerpointRepository()
	svgRepository := repository.NewSVGRepository()

	renderUsecase := usecase.NewRenderUsecase(
		sceneRepository,
		xaligoRepository,
		powerpointRepository,
		svgRepository,
	)
	diagnosticsUsecase := usecase.NewDiagnosticsUsecase()
	diffUsecase := usecase.NewDiffUsecase(xaligoRepository, sceneRepository, svgRepository)
	engineUsecase := v2.NewEngineUsecase()
	iconRegistryRepository := iconrepository.NewRegistryRepository(cfg.AssetsDB)
	iconUsecase := v2.NewIconUsecase(iconRegistryRepository, engineUsecase, builtin.IconRegistrations()...)
	projectIndexRepository := projectrepository.NewIndexRepository(cfg.ProjectDB)
	projectUsecase := usecase.NewProjectUsecase(projectIndexRepository)

	generateController := controller.NewGenerateController()
	renderController := controller.NewRenderController(renderUsecase)
	validateController := controller.NewValidateController(diagnosticsUsecase)
	serveController := controller.NewServeController(cfg, renderUsecase)
	initController := controller.NewInitController()
	versionController := controller.NewVersionController()
	diffController := controller.NewDiffController(diffUsecase)
	iconController := controller.NewIconController(iconUsecase)
	ragController := controller.NewRAGController(projectUsecase, cfg.ProjectRoot)
	lspController := controller.NewLSPController(projectUsecase)

	root := &cobra.Command{
		Use:   "xaligo",
		Short: "Diagram-as-code renderer for SVG, PPTX, and Markdown",
		Long: `xaligo is a diagram-as-code engine for architecture, network, and UML
diagrams. It parses the Vue-style .xal XML DSL once and pushes the result
through one shared parser -> layout -> draw-plan pipeline. The supported
outputs are SVG and PPTX, plus Markdown documents that embed rendered SVGs.

Use 'xaligo <command> --help' for full details, flags, and examples for any
subcommand below.`,
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return errors.Join(iconRegistryRepository.Close(), projectUsecase.Close())
		},
	}

	root.AddCommand(renderController.Command())
	root.AddCommand(validateController.Command())
	root.AddCommand(serveController.Command())
	root.AddCommand(initController.Command())
	root.AddCommand(versionController.Command())
	root.AddCommand(generateController.Command())
	root.AddCommand(diffController.Command())
	root.AddCommand(iconController.Command())
	root.AddCommand(ragController.Command())
	root.AddCommand(lspController.Command())
	return root
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		logger.ERROR(ICE001, "execute failed", map[string]any{"error": err})
		os.Exit(1)
	}
}
