package usecase

import (
	"context"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
)

type PreviewServer = repository.PreviewServer

// XaligoUsecase is the application boundary consumed by controllers and adapters.
type XaligoUsecase interface {
	ValidateRenderOptions(entity.RenderOptions) error
	Validate(context.Context, []byte) error
	Diagnose(context.Context, []byte) ([]entity.Diagnostic, error)
	Render(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderExcalidraw(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderSVG(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderPPTX(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderXYFlow(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderIsoflow(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	BuildPPTXPlan(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	NewPreviewServer(string, entity.PreviewOptions) (PreviewServer, error)
	ReadScene(string) (*entity.Scene, error)
	WriteScene(*entity.Scene, string) error
	ReadServiceList(string) ([]entity.ServiceEntry, error)
	LookupCatalogByID(string, int) (entity.CatalogEntry, error)
	SvgToDataURL(string) (string, error)
	FileID(string) string
	SVGBGColor(string) string
	ExportPptx(context.Context, entity.PptxExportOptions) error
}

type xaligoUsecase struct {
	excalidrawRepository repository.ExcalidrawRepository
	xaligoRepository     repository.XaligoRepository
	powerpointRepository repository.PowerpointRepository
	isoflowRepository    repository.IsoflowRepository
	svgRepository        repository.SVGRepository
	xyFlowRepository     repository.XYFlowRepository
}

// NewXaligoUsecase creates the application use case from explicit repository
// dependencies. Dependency construction belongs to composition roots.
func NewXaligoUsecase(
	excalidrawRepository repository.ExcalidrawRepository,
	xaligoRepository repository.XaligoRepository,
	powerpointRepository repository.PowerpointRepository,
	isoflowRepository repository.IsoflowRepository,
	svgRepository repository.SVGRepository,
	xyFlowRepository repository.XYFlowRepository,
) XaligoUsecase {
	return &xaligoUsecase{
		excalidrawRepository: excalidrawRepository,
		xaligoRepository:     xaligoRepository,
		powerpointRepository: powerpointRepository,
		isoflowRepository:    isoflowRepository,
		svgRepository:        svgRepository,
		xyFlowRepository:     xyFlowRepository,
	}
}

func (rcvr *xaligoUsecase) ValidateRenderOptions(opts entity.RenderOptions) error {
	return ValidateRenderOptions(opts)
}

func (rcvr *xaligoUsecase) Validate(ctx context.Context, input []byte) error {
	return Validate(ctx, input)
}

func (rcvr *xaligoUsecase) Diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	return Diagnose(ctx, input)
}

func (rcvr *xaligoUsecase) Render(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.render(ctx, input, opts)
}

func (rcvr *xaligoUsecase) RenderExcalidraw(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.renderExcalidraw(ctx, input, opts)
}

func (rcvr *xaligoUsecase) RenderSVG(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.renderSVG(ctx, input, opts)
}

func (rcvr *xaligoUsecase) RenderPPTX(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.renderPPTX(ctx, input, opts)
}

func (rcvr *xaligoUsecase) RenderXYFlow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.renderXYFlow(ctx, input, opts)
}

func (rcvr *xaligoUsecase) RenderIsoflow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.renderIsoflow(ctx, input, opts)
}

func (rcvr *xaligoUsecase) BuildPPTXPlan(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return rcvr.buildPPTXPlan(ctx, input, opts)
}

func (rcvr *xaligoUsecase) NewPreviewServer(path string, opts entity.PreviewOptions) (PreviewServer, error) {
	return repository.NewPreviewServer(
		path,
		opts,
		rcvr.RenderSVG,
		rcvr.ValidateRenderOptions,
		rcvr.Diagnose,
		rcvr.xaligoRepository.ReadSource,
	)
}

func (rcvr *xaligoUsecase) ReadScene(path string) (*entity.Scene, error) {
	return rcvr.excalidrawRepository.ReadScene(path)
}

func (rcvr *xaligoUsecase) WriteScene(scene *entity.Scene, path string) error {
	return rcvr.excalidrawRepository.WriteScene(scene, path)
}

func (rcvr *xaligoUsecase) ReadServiceList(path string) ([]entity.ServiceEntry, error) {
	return rcvr.xaligoRepository.ReadServiceList(path)
}

func (rcvr *xaligoUsecase) ExportPptx(ctx context.Context, opts entity.PptxExportOptions) error {
	return rcvr.powerpointRepository.WritePptxWithExporter(ctx, opts, nil)
}

func (rcvr *xaligoUsecase) LookupCatalogByID(csvPath string, id int) (entity.CatalogEntry, error) {
	return rcvr.xaligoRepository.LookupCatalogByID(csvPath, id)
}

func (rcvr *xaligoUsecase) SvgToDataURL(path string) (string, error) {
	return rcvr.excalidrawRepository.SvgToDataURL(path)
}

func (rcvr *xaligoUsecase) FileID(name string) string {
	return rcvr.excalidrawRepository.FileID(name)
}

func (rcvr *xaligoUsecase) SVGBGColor(dataURL string) string {
	return rcvr.excalidrawRepository.SVGBGColor(dataURL)
}
