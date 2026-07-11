package usecase

import (
	"io/fs"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

type SceneUsecase interface {
	BuildJSONWithFS(root *entity.Box, fsys fs.FS, catalogCSV, svgGroupDir string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string) ([]byte, error)
	BuildJSON(root *entity.Box, svgGroupDir, catalogCSV, projectRoot string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, fsys fs.FS) ([]byte, error)
}

type sceneUsecase struct {
	xaligoRepository     repository.XaligoRepository
	excalidrawRepository repository.ExcalidrawRepository
}

func NewSceneUsecase(xaligoRepository repository.XaligoRepository, excalidrawRepository repository.ExcalidrawRepository) SceneUsecase {
	return &sceneUsecase{
		xaligoRepository:     xaligoRepository,
		excalidrawRepository: excalidrawRepository,
	}
}

func (rcvr *sceneUsecase) BuildJSONWithFS(root *entity.Box, fsys fs.FS, catalogCSV, svgGroupDir string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string) ([]byte, error) {
	return v1engine.BuildJSONWithFSV1EngineSceneBuild(root, fsys, catalogCSV, svgGroupDir, itemIconSize, connections, abbrevMap, rcvr.engineDependencies())
}

func (rcvr *sceneUsecase) BuildJSON(root *entity.Box, svgGroupDir, catalogCSV, projectRoot string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, fsys fs.FS) ([]byte, error) {
	return v1engine.BuildJSONV1EngineSceneBuild(root, svgGroupDir, catalogCSV, projectRoot, itemIconSize, connections, abbrevMap, fsys, rcvr.engineDependencies())
}

func (rcvr *sceneUsecase) engineDependencies() v1engine.SceneDependenciesV1EngineSceneTypes {
	return SceneDependencies{
		XaligoRepository:     rcvr.xaligoRepository,
		ExcalidrawRepository: rcvr.excalidrawRepository,
	}.core()
}

// SceneDependencies is the compatibility boundary accepted by the public
// scene helpers. Repositories are adapted here to the V1 engine's synchronous
// function ports; the V1 engine never imports the repository layer.
type SceneDependencies struct {
	XaligoRepository     repository.XaligoRepository
	ExcalidrawRepository repository.ExcalidrawRepository
}

func (d SceneDependencies) core() v1engine.SceneDependenciesV1EngineSceneTypes {
	ports := v1engine.SceneDependenciesV1EngineSceneTypes{}
	if d.XaligoRepository != nil {
		ports.LookupCatalogByID = d.XaligoRepository.LookupCatalogByID
		ports.LookupCatalogByIDFS = d.XaligoRepository.LookupCatalogByIDFS
	}
	if d.ExcalidrawRepository != nil {
		ports.SVGToDataURL = d.ExcalidrawRepository.SvgToDataURL
		ports.SVGToDataURLFS = d.ExcalidrawRepository.SvgToDataURLFS
		ports.FileID = d.ExcalidrawRepository.FileID
		ports.SVGBGColor = d.ExcalidrawRepository.SVGBGColor
	}
	return ports
}

func BuildJSONWithFS(root *entity.Box, fsys fs.FS, catalogCSV, svgGroupDir string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, deps SceneDependencies) ([]byte, error) {
	return NewSceneUsecase(deps.XaligoRepository, deps.ExcalidrawRepository).BuildJSONWithFS(root, fsys, catalogCSV, svgGroupDir, itemIconSize, connections, abbrevMap)
}

func BuildJSON(root *entity.Box, svgGroupDir, catalogCSV, projectRoot string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, fsys fs.FS, deps SceneDependencies) ([]byte, error) {
	return NewSceneUsecase(deps.XaligoRepository, deps.ExcalidrawRepository).BuildJSON(root, svgGroupDir, catalogCSV, projectRoot, itemIconSize, connections, abbrevMap, fsys)
}
