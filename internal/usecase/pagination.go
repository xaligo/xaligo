package usecase

import (
	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

type PaginationUsecase interface {
	SplitPlanPages(entity.Plan, entity.PageSplitOptions) entity.PageSplit
	SplitPlanPagesChecked(entity.Plan, entity.PageSplitOptions) (entity.PageSplit, error)
}

type paginationUsecase struct{}

func NewPaginationUsecase() PaginationUsecase {
	return &paginationUsecase{}
}

func (rcvr *paginationUsecase) SplitPlanPages(plan entity.Plan, opts entity.PageSplitOptions) entity.PageSplit {
	return v1engine.SplitPlanPagesV1EnginePageSplit(plan, opts)
}

func (rcvr *paginationUsecase) SplitPlanPagesChecked(plan entity.Plan, opts entity.PageSplitOptions) (entity.PageSplit, error) {
	return v1engine.SplitPlanPagesCheckedV1EnginePageSplit(plan, opts)
}

// SplitPlanPages delegates renderer-neutral plan tiling to PaginationUsecase.
// Deprecated: use NewPaginationUsecase().SplitPlanPages instead.
func SplitPlanPages(plan entity.Plan, opts entity.PageSplitOptions) entity.PageSplit {
	return NewPaginationUsecase().SplitPlanPages(plan, opts)
}

// SplitPlanPagesChecked delegates validated renderer-neutral plan tiling to PaginationUsecase.
// Deprecated: use NewPaginationUsecase().SplitPlanPagesChecked instead.
func SplitPlanPagesChecked(plan entity.Plan, opts entity.PageSplitOptions) (entity.PageSplit, error) {
	return NewPaginationUsecase().SplitPlanPagesChecked(plan, opts)
}
