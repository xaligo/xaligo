package usecase

import (
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
)

type CatalogUsecase interface {
	ReadServiceList(string) ([]entity.ServiceEntry, error)
	LookupCatalogByID(string, int) (entity.CatalogEntry, error)
}

type catalogUsecase struct {
	xaligoRepository repository.XaligoRepository
}

func NewCatalogUsecase(xaligoRepository repository.XaligoRepository) CatalogUsecase {
	return &catalogUsecase{xaligoRepository: xaligoRepository}
}

func (rcvr *catalogUsecase) ReadServiceList(path string) ([]entity.ServiceEntry, error) {
	return rcvr.xaligoRepository.ReadServiceList(path)
}

func (rcvr *catalogUsecase) LookupCatalogByID(csvPath string, id int) (entity.CatalogEntry, error) {
	return rcvr.xaligoRepository.LookupCatalogByID(csvPath, id)
}
