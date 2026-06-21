package usecase

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
)

var connectionShorthandPattern = regexp.MustCompile(`^([A-Za-z0-9_.:-]+)\s*(---|==>)\s*([A-Za-z0-9_.:-]+)$`)

type Error struct {
	Position entity.Position
	Err      error
}

func (e *Error) Error() string {
	return fmt.Sprintf("line %d, column %d: %v", e.Position.Line, e.Position.Column, e.Err)
}

func (e *Error) Unwrap() error { return e.Err }

func Parse(r io.Reader) (entity.Document, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return entity.Document{}, fmt.Errorf("read DSL: %w", err)
	}
	dec := xml.NewDecoder(bytes.NewReader(data))
	var stack []*entity.Node
	var root *entity.Node

	for {
		offset := int(dec.InputOffset())
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return entity.Document{}, &Error{Position: positionAt(data, offset), Err: fmt.Errorf("parse xml-like token: %w", err)}
		}

		switch t := tok.(type) {
		case xml.StartElement:
			node := &entity.Node{Tag: t.Name.Local, Attrs: map[string]string{}, Position: positionAt(data, offset)}
			for _, a := range t.Attr {
				node.Attrs[a.Name.Local] = a.Value
			}
			if node.Tag == "item" {
				if err := validateItemNode(node); err != nil {
					return entity.Document{}, &Error{Position: node.Position, Err: fmt.Errorf("parse <item>: %w", err)}
				}
			}
			if node.Tag == "connection" {
				if err := validateConnectionNode(node); err != nil {
					return entity.Document{}, &Error{Position: node.Position, Err: fmt.Errorf("parse <connection>: %w", err)}
				}
			}
			if node.Tag == "generic-group" {
				if err := validateGenericGroupNode(node); err != nil {
					return entity.Document{}, &Error{Position: node.Position, Err: fmt.Errorf("parse <generic-group>: %w", err)}
				}
			}
			if len(stack) == 0 {
				root = node
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, node)
			}
			stack = append(stack, node)
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			text := strings.TrimSpace(string(t))
			cur := stack[len(stack)-1]
			cur.TextRuns = append(cur.TextRuns, entity.TextRun{Text: string(t), Position: positionAt(data, offset)})
			if text != "" {
				if cur.Text == "" {
					cur.Text = text
				} else {
					cur.Text += " " + text
				}
			}
		case xml.EndElement:
			if len(stack) == 0 {
				return entity.Document{}, &Error{Position: positionAt(data, offset), Err: fmt.Errorf("unexpected closing tag: %s", t.Name.Local)}
			}
			stack = stack[:len(stack)-1]
		}
	}

	if root == nil {
		return entity.Document{}, &Error{Position: entity.Position{Line: 1, Column: 1}, Err: fmt.Errorf("empty document")}
	}
	if root.Tag != "frame" {
		return entity.Document{}, &Error{Position: root.Position, Err: fmt.Errorf("root tag must be <frame>, got <%s>", root.Tag)}
	}
	if err := expandConnectionShorthands(root, data); err != nil {
		return entity.Document{}, err
	}

	return entity.Document{Root: root}, nil
}

func validateGenericGroupNode(node *entity.Node) error {
	iconID := strings.TrimSpace(node.Attrs["icon-id"])
	if iconID == "" {
		return nil
	}
	for _, ch := range iconID {
		if ch < '0' || ch > '9' {
			return fmt.Errorf("icon-id=%q must be a positive catalog ID", iconID)
		}
	}
	return nil
}

func expandConnectionShorthands(root *entity.Node, data []byte) error {
	aliases := map[string]string{}
	var collect func(*entity.Node) error
	collect = func(node *entity.Node) error {
		if node.Tag == "item" {
			id := strings.TrimSpace(node.Attrs["id"])
			if id != "" {
				aliases[id] = id
			}
			for _, key := range []string{"name", "ref"} {
				alias := strings.TrimSpace(node.Attrs[key])
				if alias == "" {
					continue
				}
				if id == "" {
					return &Error{Position: node.Position, Err: fmt.Errorf("<item %s=%q> requires a non-empty id", key, alias)}
				}
				if previous, exists := aliases[alias]; exists && previous != id {
					return &Error{Position: node.Position, Err: fmt.Errorf("duplicate item reference %q", alias)}
				}
				aliases[alias] = id
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
		return err
	}

	for _, run := range root.TextRuns {
		lineOffset := 0
		for _, line := range strings.SplitAfter(run.Text, "\n") {
			withoutNewline := strings.TrimSuffix(line, "\n")
			trimmed := strings.TrimSpace(withoutNewline)
			leading := len(withoutNewline) - len(strings.TrimLeft(withoutNewline, " \t\r"))
			position := positionAt(data, run.Position.Offset+lineOffset+leading)
			lineOffset += len(line)
			if trimmed == "" {
				continue
			}
			match := connectionShorthandPattern.FindStringSubmatch(trimmed)
			if match == nil {
				if strings.Contains(trimmed, "---") || strings.Contains(trimmed, "==>") {
					return &Error{Position: position, Err: fmt.Errorf("invalid connection shorthand %q; expected 'source --- destination' or 'source ==> destination'", trimmed)}
				}
				continue
			}
			src, ok := aliases[match[1]]
			if !ok || src == "" {
				return &Error{Position: position, Err: fmt.Errorf("connection shorthand source %q does not match an <item name=...>, <item ref=...>, or item ID", match[1])}
			}
			dst, ok := aliases[match[3]]
			if !ok || dst == "" {
				return &Error{Position: position, Err: fmt.Errorf("connection shorthand destination %q does not match an <item name=...>, <item ref=...>, or item ID", match[3])}
			}
			kind := "route"
			if match[2] == "==>" {
				kind = "traffic"
			}
			root.Children = append(root.Children, &entity.Node{
				Tag: "connection", Attrs: map[string]string{"src": src, "dst": dst, "kind": kind}, Position: position,
			})
		}
	}
	return nil
}

func positionAt(data []byte, offset int) entity.Position {
	if offset < 0 {
		offset = 0
	}
	if offset > len(data) {
		offset = len(data)
	}
	prefix := data[:offset]
	line := bytes.Count(prefix, []byte{'\n'}) + 1
	lastNewline := bytes.LastIndexByte(prefix, '\n')
	column := offset - lastNewline
	return entity.Position{Offset: offset, Line: line, Column: column}
}

// validateItemNode ensures <item> carries at most one numeric id attribute.
// An empty (or absent) id is allowed — the item acts as a layout spacer.
func validateItemNode(node *entity.Node) error {
	id, ok := node.Attrs["id"]
	if !ok || strings.TrimSpace(id) == "" {
		return nil // spacer item — no id required
	}
	if strings.Contains(id, ",") {
		return fmt.Errorf("<item id=%q> must contain a single ID; use separate <item> tags for multiple services", id)
	}
	for _, ch := range strings.TrimSpace(id) {
		if ch < '0' || ch > '9' {
			return fmt.Errorf("<item id=%q> must be a positive integer", id)
		}
	}
	return nil
}

// validateConnectionNode ensures <connection> carries numeric src and dst attributes.
func validateConnectionNode(node *entity.Node) error {
	src, hasSrc := node.Attrs["src"]
	dst, hasDst := node.Attrs["dst"]
	if !hasSrc || strings.TrimSpace(src) == "" {
		return fmt.Errorf("<connection> requires a src attribute")
	}
	if !hasDst || strings.TrimSpace(dst) == "" {
		return fmt.Errorf("<connection> requires a dst attribute")
	}
	for _, ch := range strings.TrimSpace(src) {
		if ch < '0' || ch > '9' {
			return fmt.Errorf("<connection src=%q> must be a positive integer", src)
		}
	}
	for _, ch := range strings.TrimSpace(dst) {
		if ch < '0' || ch > '9' {
			return fmt.Errorf("<connection dst=%q> must be a positive integer", dst)
		}
	}
	return nil
}
