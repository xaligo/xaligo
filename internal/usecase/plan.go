package usecase

import (
	"github.com/xaligo/xaligo/internal/entity"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

type PlanUsecase interface {
	BuildPlanJSON(string, entity.PlanOptions) ([]byte, error)
	BuildPlan(*entity.PresentationScene, entity.PlanOptions) entity.Plan
}

type planUsecase struct{}

func NewPlanUsecase() PlanUsecase {
	return &planUsecase{}
}

func (rcvr *planUsecase) BuildPlanJSON(sceneJSON string, opt entity.PlanOptions) ([]byte, error) {
	return v1engine.BuildPlanJSONV1EnginePlanBuild(sceneJSON, opt)
}

func (rcvr *planUsecase) BuildPlan(scene *entity.PresentationScene, opt entity.PlanOptions) entity.Plan {
	return v1engine.BuildPlanV1EnginePlanBuild(scene, opt)
}

// BuildPlanJSON delegates renderer-neutral draw-plan construction to PlanUsecase.
// Deprecated: use NewPlanUsecase().BuildPlanJSON instead.
func BuildPlanJSON(sceneJSON string, opt entity.PlanOptions) ([]byte, error) {
	return NewPlanUsecase().BuildPlanJSON(sceneJSON, opt)
}

// BuildPlan delegates renderer-neutral draw-plan construction to PlanUsecase.
// Deprecated: use NewPlanUsecase().BuildPlan instead.
func BuildPlan(scene *entity.PresentationScene, opt entity.PlanOptions) entity.Plan {
	return NewPlanUsecase().BuildPlan(scene, opt)
}
