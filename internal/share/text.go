package share

import "unicode"

// PresentationTextWidth estimates a single-line presentation-text width in
// layout pixels. Layout and physical encoders use the same estimate so auto
// sized boxes do not trigger encoder-only shrink-to-fit behavior.
func PresentationTextWidth(value string, fontSize float64, bold bool) float64 {
	factor := 1.0
	if bold {
		factor = 1.05
	}
	units := 0.0
	for _, character := range value {
		switch {
		case unicode.Is(unicode.Mn, character), unicode.Is(unicode.Me, character):
			continue
		case unicode.IsSpace(character):
			units += 0.33
		case character >= 0x1100:
			units += 1.0
		case unicode.IsPunct(character):
			units += 0.42
		case unicode.IsUpper(character):
			units += 0.62
		default:
			units += 0.55
		}
	}
	return units * fontSize * factor
}
