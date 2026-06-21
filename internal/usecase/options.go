package usecase

import (
	"fmt"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
)

// ValidateRenderOptions validates mode, format, assets, and shared presentation
// values without parsing an input document.
func ValidateRenderOptions(opts RenderOptions) error {
	if err := validateMode(opts.Mode); err != nil {
		return err
	}
	if opts.PaperMarginIn < 0 || opts.PaperMarginTopIn < 0 || opts.PaperMarginRightIn < 0 || opts.PaperMarginBottomIn < 0 || opts.PaperMarginLeftIn < 0 {
		return fmt.Errorf("paper margins must be non-negative")
	}
	if _, err := entity.NormalizeTheme(opts.Theme); err != nil {
		return err
	}
	if opts.Assets != nil {
		if opts.Assets.FS == nil {
			return fmt.Errorf("asset source filesystem is required")
		}
		if strings.TrimSpace(opts.Assets.CatalogCSV) == "" || strings.TrimSpace(opts.Assets.GroupIconsDir) == "" {
			return fmt.Errorf("asset source catalog and group icons directory are required")
		}
	}
	format := Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	switch format {
	case "", FormatExcalidraw, FormatSVG, FormatPPTX, FormatXYFlow, FormatIsoflow:
		return nil
	default:
		return fmt.Errorf("unknown render format %q", format)
	}
}

func validateMode(mode Mode) error {
	normalized := Mode(strings.ToLower(strings.TrimSpace(string(mode))))
	switch normalized {
	case "", ModeStandard, ModeNetwork, ModeAWS:
		return nil
	case "aws-2.5d", "topology":
		return fmt.Errorf("mode %q: %w", normalized, ErrNotImplemented)
	default:
		return fmt.Errorf("unknown render mode %q", normalized)
	}
}
