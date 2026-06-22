package usecase

import (
	"context"

	"github.com/ryo-arima/xaligo/internal/entity"
)

// API is the application boundary consumed by controllers and adapters.
type API interface {
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
}

type service struct{}

// New creates a usecase service for controllers and adapters.
func New() API {
	return &service{}
}

func (s *service) ValidateRenderOptions(opts entity.RenderOptions) error {
	return ValidateRenderOptions(opts)
}

func (s *service) Validate(ctx context.Context, input []byte) error {
	return Validate(ctx, input)
}

func (s *service) Diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	return Diagnose(ctx, input)
}

func (s *service) Render(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return Render(ctx, input, opts)
}

func (s *service) RenderExcalidraw(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return RenderExcalidraw(ctx, input, opts)
}

func (s *service) RenderSVG(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return RenderSVG(ctx, input, opts)
}

func (s *service) RenderPPTX(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return RenderPPTX(ctx, input, opts)
}

func (s *service) RenderXYFlow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return RenderXYFlow(ctx, input, opts)
}

func (s *service) RenderIsoflow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return RenderIsoflow(ctx, input, opts)
}

func (s *service) BuildPPTXPlan(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	return BuildPPTXPlan(ctx, input, opts)
}

func (s *service) NewPreviewServer(path string, opts entity.PreviewOptions) (PreviewServer, error) {
	return NewPreviewServer(path, opts)
}
