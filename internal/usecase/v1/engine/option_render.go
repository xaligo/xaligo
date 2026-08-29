package engine

import (
	"fmt"
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const (
	ModeStandardV1EngineOptionRender entity.Mode = "standard"
	ModeNetworkV1EngineOptionRender  entity.Mode = "network"
	ModeAWSV1EngineOptionRender      entity.Mode = "aws"

	FormatSVGV1EngineOptionRender  entity.Format = "svg"
	FormatPPTXV1EngineOptionRender entity.Format = "pptx"

	SeverityErrorV1EngineOptionRender   entity.DiagnosticSeverity = "error"
	SeverityWarningV1EngineOptionRender entity.DiagnosticSeverity = "warning"
)

// ValidateRenderOptionsV1EngineOptionRender validates mode, format, assets, and shared presentation
// values without parsing an input document.
func ValidateRenderOptionsV1EngineOptionRender(opts entity.RenderOptions) error {
	if err := validateModeV1EngineOptionRender(opts.Mode); err != nil {
		return err
	}
	for _, option := range []struct {
		name  string
		value float64
	}{
		{"pixels per inch", opts.PxPerInch},
		{"arrow stub", opts.ArrowStubPx},
		{"arrow margin", opts.ArrowMarginPx},
		{"paper margin", opts.PaperMarginIn},
		{"paper margin top", opts.PaperMarginTopIn},
		{"paper margin right", opts.PaperMarginRightIn},
		{"paper margin bottom", opts.PaperMarginBottomIn},
		{"paper margin left", opts.PaperMarginLeftIn},
	} {
		if err := validateOptionalNonNegativeV1EngineOptionRender(option.name, option.value); err != nil {
			return err
		}
	}
	if err := validatePresentationOptionsV1EngineOptionRender(opts); err != nil {
		return err
	}
	switch strings.ToLower(strings.TrimSpace(opts.SVGLegendPosition)) {
	case "", "top", "right", "bottom", "left":
	default:
		return fmt.Errorf("unknown SVG legend position %q; valid: top, right, bottom, left", opts.SVGLegendPosition)
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
		if err := validateOptionalNonNegativeV1EngineOptionRender("asset item icon size", opts.Assets.ItemIconSize); err != nil {
			return err
		}
	}
	format := entity.Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	switch format {
	case "", FormatSVGV1EngineOptionRender, FormatPPTXV1EngineOptionRender:
		return nil
	default:
		return fmt.Errorf("unknown render format %q", format)
	}
}

func validatePresentationOptionsV1EngineOptionRender(opts entity.RenderOptions) error {
	switch strings.TrimSpace(opts.ArrowStyle) {
	case "", "thin", "standard", "triangle", "stealth", "arrow", "diamond", "oval", "none":
	default:
		return fmt.Errorf("unknown arrow style %q; valid: thin, standard, triangle, stealth, arrow, diamond, oval, none", opts.ArrowStyle)
	}

	orientation := strings.TrimSpace(opts.Orientation)
	switch orientation {
	case "", "portrait", "landscape":
	default:
		return fmt.Errorf("unknown paper orientation %q; valid: portrait, landscape", opts.Orientation)
	}

	paper := strings.TrimSpace(opts.PaperSize)
	base, hasPaper := paperSizeInV1EnginePlanPaper[paper]
	if paper != "" && !hasPaper {
		return fmt.Errorf("unknown paper size %q", opts.PaperSize)
	}
	hasMargins := opts.PaperMarginIn > 0 || opts.PaperMarginTopIn > 0 || opts.PaperMarginRightIn > 0 || opts.PaperMarginBottomIn > 0 || opts.PaperMarginLeftIn > 0
	if !hasPaper {
		if hasMargins {
			return fmt.Errorf("paper margins require a paper size")
		}
		return nil
	}

	top, right, bottom, left := opts.PaperMarginIn, opts.PaperMarginIn, opts.PaperMarginIn, opts.PaperMarginIn
	if opts.PaperMarginTopIn > 0 {
		top = opts.PaperMarginTopIn
	}
	if opts.PaperMarginRightIn > 0 {
		right = opts.PaperMarginRightIn
	}
	if opts.PaperMarginBottomIn > 0 {
		bottom = opts.PaperMarginBottomIn
	}
	if opts.PaperMarginLeftIn > 0 {
		left = opts.PaperMarginLeftIn
	}
	fits := func(width, height float64) bool {
		return left+right < width && top+bottom < height
	}
	portraitOK := fits(base[0], base[1])
	landscapeOK := fits(base[1], base[0])
	valid := portraitOK || landscapeOK
	if orientation == "portrait" {
		valid = portraitOK
	} else if orientation == "landscape" {
		valid = landscapeOK
	}
	if !valid {
		return fmt.Errorf("paper margins leave no positive content area on %s %s", paper, orientation)
	}
	return nil
}

// A zero value means "use the renderer default" for these options. Explicit
// negative and non-finite values cannot be allowed to reach geometry or JSON
// encoding, where they otherwise fail later than `validate`.
func validateOptionalNonNegativeV1EngineOptionRender(name string, value float64) error {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return fmt.Errorf("%s must be finite", name)
	}
	if value < 0 {
		return fmt.Errorf("%s must be non-negative", name)
	}
	return nil
}

func validateModeV1EngineOptionRender(mode entity.Mode) error {
	normalized := entity.Mode(strings.ToLower(strings.TrimSpace(string(mode))))
	switch normalized {
	case "", ModeStandardV1EngineOptionRender, ModeNetworkV1EngineOptionRender, ModeAWSV1EngineOptionRender:
		return nil
	default:
		return fmt.Errorf("unknown render mode %q", normalized)
	}
}
