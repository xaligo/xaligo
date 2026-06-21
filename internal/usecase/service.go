package usecase

import "context"

// API is the application boundary consumed by controllers and adapters.
type API interface {
	ValidateRenderOptions(RenderOptions) error
	Validate(context.Context, []byte) error
	Diagnose(context.Context, []byte) ([]Diagnostic, error)
	Render(context.Context, []byte, RenderOptions) ([]byte, error)
	RenderExcalidraw(context.Context, []byte, RenderOptions) ([]byte, error)
	RenderSVG(context.Context, []byte, RenderOptions) ([]byte, error)
	RenderPPTX(context.Context, []byte, RenderOptions) ([]byte, error)
	RenderXYFlow(context.Context, []byte, RenderOptions) ([]byte, error)
	RenderIsoflow(context.Context, []byte, RenderOptions) ([]byte, error)
	BuildPPTXPlan(context.Context, []byte, RenderOptions) ([]byte, error)
	NewPreviewServer(string, PreviewOptions) (*PreviewServer, error)
}

// Dependencies groups usecase collaborators for constructor injection.
type Dependencies struct{}

// Service implements API.
type Service struct{}

// New creates a usecase service. Dependencies is intentionally small for now;
// add collaborators here instead of letting controllers import lower layers.
func New(_ Dependencies) API {
	return &Service{}
}

func (s *Service) ValidateRenderOptions(opts RenderOptions) error {
	return ValidateRenderOptions(opts)
}

func (s *Service) Validate(ctx context.Context, input []byte) error {
	return Validate(ctx, input)
}

func (s *Service) Diagnose(ctx context.Context, input []byte) ([]Diagnostic, error) {
	return Diagnose(ctx, input)
}

func (s *Service) Render(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	return Render(ctx, input, opts)
}

func (s *Service) RenderExcalidraw(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	return RenderExcalidraw(ctx, input, opts)
}

func (s *Service) RenderSVG(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	return RenderSVG(ctx, input, opts)
}

func (s *Service) RenderPPTX(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	return RenderPPTX(ctx, input, opts)
}

func (s *Service) RenderXYFlow(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	return RenderXYFlow(ctx, input, opts)
}

func (s *Service) RenderIsoflow(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	return RenderIsoflow(ctx, input, opts)
}

func (s *Service) BuildPPTXPlan(ctx context.Context, input []byte, opts RenderOptions) ([]byte, error) {
	return BuildPPTXPlan(ctx, input, opts)
}

func (s *Service) NewPreviewServer(path string, opts PreviewOptions) (*PreviewServer, error) {
	return NewPreviewServer(path, opts)
}
