package integration

import (
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/ryo-arima/xaligo/internal/usecase"
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
