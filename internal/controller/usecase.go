package controller

import (
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

func defaultUseCase(uc usecase.API) usecase.API {
	if uc != nil {
		return uc
	}
	return usecase.New(entity.UseCaseDependencies{})
}
