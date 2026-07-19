package engine

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

var connectionShorthandPatternV1EngineParseShorthand = regexp.MustCompile(`^([A-Za-z0-9_.:-]+)\s*(---|==>)\s*([A-Za-z0-9_.:-]+)$`)

func expandConnectionShorthandsV1EngineParseShorthand(root *entity.Node, data []byte) error {
	var expandInFrame func(*entity.Node) error
	expandInFrame = func(frame *entity.Node) error {
		for _, run := range frame.TextRuns {
			lineOffset := 0
			for _, line := range strings.SplitAfter(run.Text, "\n") {
				withoutNewline := strings.TrimSuffix(line, "\n")
				trimmed := strings.TrimSpace(withoutNewline)
				leading := len(withoutNewline) - len(strings.TrimLeft(withoutNewline, " \t\r"))
				position := positionAtV1EngineParseDocument(data, run.Position.Offset+lineOffset+leading)
				lineOffset += len(line)
				if trimmed == "" {
					loggerV1EngineSharedLogging.DEBUG(IUPECS007V1EngineParseDocument, "branch empty line")
					continue
				}
				match := connectionShorthandPatternV1EngineParseShorthand.FindStringSubmatch(trimmed)
				if match == nil {
					if strings.Contains(trimmed, "---") || strings.Contains(trimmed, "==>") {
						loggerV1EngineSharedLogging.ERROR(IUPECS008V1EngineParseDocument, "branch invalid shorthand", map[string]any{"line": trimmed})
						return &entity.ParseError{Position: position, Err: fmt.Errorf("invalid connection shorthand %q; expected 'source --- destination' or 'source ==> destination'", trimmed)}
					}
					loggerV1EngineSharedLogging.DEBUG(IUPECS009V1EngineParseDocument, "branch non shorthand", map[string]any{"line": trimmed})
					continue
				}
				kind := "route"
				if match[2] == "==>" {
					loggerV1EngineSharedLogging.DEBUG(IUPECS012V1EngineParseDocument, "branch traffic")
					kind = "traffic"
				}
				frame.Children = append(frame.Children, &entity.Node{
					Tag: "connection",
					Attrs: map[string]string{
						"src":  match[1],
						"dst":  match[3],
						"kind": kind,
					},
					Position: position,
				})
			}
		}
		for _, child := range frame.Children {
			if child.Tag == "frame" {
				if err := expandInFrame(child); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return expandInFrame(root)
}
