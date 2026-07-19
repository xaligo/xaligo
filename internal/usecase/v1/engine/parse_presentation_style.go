package engine

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

var presentationColorPatternV1EngineParsePresentationStyle = regexp.MustCompile(`^#[0-9a-fA-F]{6}$`)

// normalizePresentationStyleV1EngineParsePresentationStyle validates the
// renderer-neutral color and font attributes shared by tables and frame
// metadata. Numeric font-size validation remains in the common layout stage.
func normalizePresentationStyleV1EngineParsePresentationStyle(node *entity.Node, colorAttributes ...string) error {
	if node == nil {
		return nil
	}
	for _, name := range colorAttributes {
		value, exists := node.Attrs[name]
		if !exists {
			continue
		}
		value = strings.TrimSpace(value)
		if !presentationColorPatternV1EngineParsePresentationStyle.MatchString(value) && value != "transparent" {
			return fmt.Errorf("<%s %s=%q> must be #RRGGBB or transparent", node.Tag, name, value)
		}
		node.Attrs[name] = strings.ToLower(value)
	}
	if family, exists := node.Attrs["font-family"]; exists {
		family = strings.ToLower(strings.TrimSpace(family))
		switch family {
		case "virgil", "helvetica", "cascadia", "assistant", "excalifont", "nunito", "lilita-one", "comic-shanns", "liberation-sans":
			node.Attrs["font-family"] = family
		default:
			return fmt.Errorf("<%s font-family=%q> is not a supported font family", node.Tag, family)
		}
	}
	return nil
}
