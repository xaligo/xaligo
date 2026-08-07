package integration

import (
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
)

func newUsecase() usecase.RenderUsecase {
	return usecase.NewRenderUsecase(
		repository.NewSceneRepository(),
		repository.NewXaligoRepository(),
		repository.NewPowerpointRepository(),
		repository.NewSVGRepository(),
		repository.NewTerminalRepository(),
	)
}
