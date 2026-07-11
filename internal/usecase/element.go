package usecase

import v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"

type ElementUsecase interface {
	MakeText(id string, x, y, w, h float64, text string, fontSize int, color string, bold bool, textAlign string, seed int) map[string]interface{}
	MakeImage(id string, x, y, w, h float64, fileID string, backgroundColor string, seed int) map[string]interface{}
}

type elementUsecase struct{}

func NewElementUsecase() ElementUsecase {
	return &elementUsecase{}
}

func (rcvr *elementUsecase) MakeText(id string, x, y, w, h float64, text string, fontSize int, color string, bold bool, textAlign string, seed int) map[string]interface{} {
	return v1engine.MakeTextV1EngineElementBuild(id, x, y, w, h, text, fontSize, color, bold, textAlign, seed)
}

func (rcvr *elementUsecase) MakeImage(id string, x, y, w, h float64, fileID string, backgroundColor string, seed int) map[string]interface{} {
	return v1engine.MakeImageV1EngineElementBuild(id, x, y, w, h, fileID, backgroundColor, seed)
}

// MakeText delegates Excalidraw-compatible text element construction to ElementUsecase.
// Deprecated: use NewElementUsecase().MakeText instead.
func MakeText(id string, x, y, w, h float64, text string, fontSize int, color string, bold bool, textAlign string, seed int) map[string]interface{} {
	return NewElementUsecase().MakeText(id, x, y, w, h, text, fontSize, color, bold, textAlign, seed)
}

// MakeImage delegates Excalidraw-compatible image element construction to ElementUsecase.
// Deprecated: use NewElementUsecase().MakeImage instead.
func MakeImage(id string, x, y, w, h float64, fileID string, backgroundColor string, seed int) map[string]interface{} {
	return NewElementUsecase().MakeImage(id, x, y, w, h, fileID, backgroundColor, seed)
}
