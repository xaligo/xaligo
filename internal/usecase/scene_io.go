package usecase

import (
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
)

type SceneIOUsecase interface {
	ReadScene(string) (*entity.Scene, error)
	WriteScene(*entity.Scene, string) error
	SvgToDataURL(string) (string, error)
	FileID(string) string
	SVGBGColor(string) string
}

type sceneIOUsecase struct {
	excalidrawRepository repository.ExcalidrawRepository
}

func NewSceneIOUsecase(excalidrawRepository repository.ExcalidrawRepository) SceneIOUsecase {
	return &sceneIOUsecase{excalidrawRepository: excalidrawRepository}
}

func (rcvr *sceneIOUsecase) ReadScene(path string) (*entity.Scene, error) {
	return rcvr.excalidrawRepository.ReadScene(path)
}

func (rcvr *sceneIOUsecase) WriteScene(scene *entity.Scene, path string) error {
	return rcvr.excalidrawRepository.WriteScene(scene, path)
}

func (rcvr *sceneIOUsecase) SvgToDataURL(path string) (string, error) {
	return rcvr.excalidrawRepository.SvgToDataURL(path)
}

func (rcvr *sceneIOUsecase) FileID(name string) string {
	return rcvr.excalidrawRepository.FileID(name)
}

func (rcvr *sceneIOUsecase) SVGBGColor(dataURL string) string {
	return rcvr.excalidrawRepository.SVGBGColor(dataURL)
}
