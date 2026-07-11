package usecase

import (
	"context"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
)

type ExportUsecase interface {
	ExportPptx(context.Context, entity.PptxExportOptions) error
}

type exportUsecase struct {
	powerpointRepository repository.PowerpointRepository
}

func NewExportUsecase(powerpointRepository repository.PowerpointRepository) ExportUsecase {
	return &exportUsecase{powerpointRepository: powerpointRepository}
}

func (rcvr *exportUsecase) ExportPptx(ctx context.Context, opts entity.PptxExportOptions) error {
	if err := checkContext(ctx); err != nil {
		return err
	}
	return rcvr.powerpointRepository.WritePptx(ctx, opts)
}
