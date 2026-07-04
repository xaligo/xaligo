package integration

import (
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

func newUsecase() usecase.XaligoUsecase {
	return usecase.NewXaligoUsecase(
		repository.NewExcalidrawRepository(),
		repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(),
		repository.NewIsoflowRepository(),
		repository.NewSVGRepository(),
		repository.NewXYFlowRepository(),
	)
}
