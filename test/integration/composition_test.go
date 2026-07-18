package integration

import (
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

func newUsecase() usecase.RenderUsecase {
	return usecase.NewRenderUsecase(
		repository.NewExcalidrawRepository(),
		repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(),
		repository.NewIsoflowRepository(),
		repository.NewSVGRepository(),
		repository.NewXYFlowRepository(),
		repository.NewPDFRepository(),
		repository.NewSpreadsheetRepository(),
	)
}
