package usecase

import (
	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

type LayoutUsecase interface {
	Build(entity.Document) (*entity.Box, error)
	IsItemLike(string) bool
	IsBlank(string) bool
}

type layoutUsecase struct{}

func NewLayoutUsecase() LayoutUsecase {
	return &layoutUsecase{}
}

const (
	MinBoxWidth  = v1engine.MinBoxWidthV1EngineLayoutFlow
	MinBoxHeight = v1engine.MinBoxHeightV1EngineLayoutFlow

	GroupTopInset  = v1engine.GroupTopInsetV1EngineLayoutNode
	GroupSideInset = v1engine.GroupSideInsetV1EngineLayoutNode
)

func (rcvr *layoutUsecase) Build(doc entity.Document) (*entity.Box, error) {
	return v1engine.BuildV1EngineLayoutBuild(doc)
}

func (rcvr *layoutUsecase) IsItemLike(tag string) bool {
	return v1engine.IsItemLikeV1EngineLayoutAttributes(tag)
}

func (rcvr *layoutUsecase) IsBlank(tag string) bool {
	return v1engine.IsBlankV1EngineLayoutAttributes(tag)
}

// Build delegates resolved layout calculation to LayoutUsecase.
// Deprecated: use NewLayoutUsecase().Build instead.
func Build(doc entity.Document) (*entity.Box, error) {
	return NewLayoutUsecase().Build(doc)
}

// IsItemLike reports whether a tag participates in item-grid layout.
// Deprecated: use NewLayoutUsecase().IsItemLike instead.
func IsItemLike(tag string) bool {
	return NewLayoutUsecase().IsItemLike(tag)
}

// IsBlank reports whether a tag is an explicit empty layout slot.
// Deprecated: use NewLayoutUsecase().IsBlank instead.
func IsBlank(tag string) bool {
	return NewLayoutUsecase().IsBlank(tag)
}
