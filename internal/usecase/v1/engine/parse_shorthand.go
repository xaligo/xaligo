package engine

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

var connectionShorthandPatternV1EngineParseShorthand = regexp.MustCompile(`^([A-Za-z0-9_.:-]+)\s*(---|==>)\s*([A-Za-z0-9_.:-]+)$`)

func expandConnectionShorthandsV1EngineParseShorthand(root *entity.Node, data []byte) error {
	aliases := map[string]string{}
	duplicateAliases := map[string]bool{}
	var collect func(*entity.Node) error
	collect = func(node *entity.Node) error {
		if nodeConnectableByIDV1EngineParseNode(node) {
			loggerV1EngineSharedLogging.DEBUG(IUPECS001V1EngineParseDocument, "branch item", map[string]any{"tag": node.Tag})
			id := strings.TrimSpace(node.Attrs["id"])
			key := strings.TrimSpace(node.Attrs[internalConnectionKeyAttrV1EngineParseDocument])
			if id != "" {
				loggerV1EngineSharedLogging.DEBUG(IUPECS002V1EngineParseDocument, "branch item ID", map[string]any{"id": id})
				if _, exists := aliases[id]; exists {
					duplicateAliases[id] = true
				} else {
					aliases[id] = key
				}
			}
			for _, key := range []string{"name", "ref"} {
				alias := strings.TrimSpace(node.Attrs[key])
				if alias == "" {
					loggerV1EngineSharedLogging.DEBUG(IUPECS003V1EngineParseDocument, "branch empty alias", map[string]any{"key": key})
					continue
				}
				if _, exists := aliases[alias]; exists || duplicateAliases[alias] {
					loggerV1EngineSharedLogging.DEBUG(IUPECS005V1EngineParseDocument, "branch duplicate alias", map[string]any{"alias": alias, "id": id})
					return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate connection reference %q", alias)}
				}
				aliases[alias] = node.Attrs[internalConnectionKeyAttrV1EngineParseDocument]
			}
		}
		for _, child := range node.Children {
			if err := collect(child); err != nil {
				return err
			}
		}
		return nil
	}
	if err := collect(root); err != nil {
		loggerV1EngineSharedLogging.ERROR(IUPECS006V1EngineParseDocument, "collect failed", map[string]any{"error": err})
		return err
	}

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
				src, ok := aliases[match[1]]
				if !ok || src == "" || duplicateAliases[match[1]] {
					loggerV1EngineSharedLogging.ERROR(IUPECS010V1EngineParseDocument, "branch missing source", map[string]any{"source": match[1]})
					return &entity.ParseError{Position: position, Err: fmt.Errorf("connection shorthand source %q does not match any connection endpoint id/name/ref", match[1])}
				}
				dst, ok := aliases[match[3]]
				if !ok || dst == "" || duplicateAliases[match[3]] {
					loggerV1EngineSharedLogging.ERROR(IUPECS011V1EngineParseDocument, "branch missing destination", map[string]any{"destination": match[3]})
					return &entity.ParseError{Position: position, Err: fmt.Errorf("connection shorthand destination %q does not match any connection endpoint id/name/ref", match[3])}
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
						internalConnectionSrcKeyAttrV1EngineParseDocument: src,
						internalConnectionDstKeyAttrV1EngineParseDocument: dst,
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
