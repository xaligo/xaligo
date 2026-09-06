package share

import (
	"strings"
	"unicode"
)

// WrapPresentationText inserts explicit line breaks before rendering so SVG
// and presentation encoders consume the same lines and reserved text height.
func WrapPresentationText(value string, width, fontSize float64) string {
	var result []string
	for _, paragraph := range strings.Split(value, "\n") {
		line := ""
		for _, character := range paragraph {
			candidate := line + string(character)
			if line != "" && PresentationTextWidth(candidate, fontSize, false) > width {
				if split := strings.LastIndex(line, " "); split > 0 {
					result = append(result, line[:split])
					line = strings.TrimLeft(line[split+1:]+string(character), " ")
				} else {
					result = append(result, line)
					line = string(character)
				}
			} else {
				line = candidate
			}
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

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
		case isPresentationFullWidth(character):
			units += 1.0
		case unicode.IsSpace(character):
			units += 0.33
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

// isPresentationFullWidth reports the East Asian wide/full-width ranges used
// by the presentation text estimator. In particular, the half-width forms at
// U+FF61..U+FF9F must not inherit a full em merely because their code points
// are adjacent to the full-width forms.
func isPresentationFullWidth(character rune) bool {
	return character >= 0x1100 && (character <= 0x115f ||
		character == 0x2329 || character == 0x232a ||
		(character >= 0x2e80 && character <= 0xa4cf && character != 0x303f) ||
		(character >= 0xac00 && character <= 0xd7a3) ||
		(character >= 0xf900 && character <= 0xfaff) ||
		(character >= 0xfe10 && character <= 0xfe19) ||
		(character >= 0xfe30 && character <= 0xfe6f) ||
		(character >= 0xff00 && character <= 0xff60) ||
		(character >= 0xffe0 && character <= 0xffe6) ||
		(character >= 0x20000 && character <= 0x2fffd) ||
		(character >= 0x30000 && character <= 0x3fffd))
}
