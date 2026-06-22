package usecase

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
)

var connectionShorthandPattern = regexp.MustCompile(`^([A-Za-z0-9_.:-]+)\s*(---|==>)\s*([A-Za-z0-9_.:-]+)$`)

var (
	IUPP001    = share.NewMCode("IUPP-001", "Parse read DSL failed")
	IUPP002    = share.NewMCode("IUPP-002", "Parse EOF branch")
	IUPP003    = share.NewMCode("IUPP-003", "Parse token failed")
	IUPP004    = share.NewMCode("IUPP-004", "Parse start element branch")
	IUPP005    = share.NewMCode("IUPP-005", "Parse item validation failed")
	IUPP006    = share.NewMCode("IUPP-006", "Parse connection validation failed")
	IUPP007    = share.NewMCode("IUPP-007", "Parse generic group validation failed")
	IUPP008    = share.NewMCode("IUPP-008", "Parse root assignment branch")
	IUPP009    = share.NewMCode("IUPP-009", "Parse child append branch")
	IUPP010    = share.NewMCode("IUPP-010", "Parse char data empty stack branch")
	IUPP011    = share.NewMCode("IUPP-011", "Parse text assignment branch")
	IUPP012    = share.NewMCode("IUPP-012", "Parse text append branch")
	IUPP013    = share.NewMCode("IUPP-013", "Parse unexpected closing tag branch")
	IUPP014    = share.NewMCode("IUPP-014", "Parse empty document branch")
	IUPP015    = share.NewMCode("IUPP-015", "Parse invalid root branch")
	IUPP016    = share.NewMCode("IUPP-016", "Parse expand connection shorthands failed")
	IUPVGGN001 = share.NewMCode("IUPVGGN-001", "Validate generic group empty icon ID branch")
	IUPVGGN002 = share.NewMCode("IUPVGGN-002", "Validate generic group invalid icon ID branch")
	IUPECS001  = share.NewMCode("IUPECS-001", "Expand connection shorthands item branch")
	IUPECS002  = share.NewMCode("IUPECS-002", "Expand connection shorthands item ID branch")
	IUPECS003  = share.NewMCode("IUPECS-003", "Expand connection shorthands empty alias branch")
	IUPECS004  = share.NewMCode("IUPECS-004", "Expand connection shorthands alias without ID branch")
	IUPECS005  = share.NewMCode("IUPECS-005", "Expand connection shorthands duplicate alias branch")
	IUPECS006  = share.NewMCode("IUPECS-006", "Expand connection shorthands collect failed")
	IUPECS007  = share.NewMCode("IUPECS-007", "Expand connection shorthands empty line branch")
	IUPECS008  = share.NewMCode("IUPECS-008", "Expand connection shorthands invalid shorthand branch")
	IUPECS009  = share.NewMCode("IUPECS-009", "Expand connection shorthands non shorthand branch")
	IUPECS010  = share.NewMCode("IUPECS-010", "Expand connection shorthands missing source branch")
	IUPECS011  = share.NewMCode("IUPECS-011", "Expand connection shorthands missing destination branch")
	IUPECS012  = share.NewMCode("IUPECS-012", "Expand connection shorthands traffic branch")
	IUPPA001   = share.NewMCode("IUPPA-001", "Position at negative offset branch")
	IUPPA002   = share.NewMCode("IUPPA-002", "Position at overflow offset branch")
	IUPVIN001  = share.NewMCode("IUPVIN-001", "Validate item spacer branch")
	IUPVIN002  = share.NewMCode("IUPVIN-002", "Validate item comma branch")
	IUPVIN003  = share.NewMCode("IUPVIN-003", "Validate item non numeric branch")
	IUPVCN001  = share.NewMCode("IUPVCN-001", "Validate connection missing source branch")
	IUPVCN002  = share.NewMCode("IUPVCN-002", "Validate connection missing destination branch")
	IUPVCN003  = share.NewMCode("IUPVCN-003", "Validate connection non numeric source branch")
	IUPVCN004  = share.NewMCode("IUPVCN-004", "Validate connection non numeric destination branch")
)

func Parse(r io.Reader) (entity.Document, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		logger.ERROR(IUPP001, "read DSL failed", map[string]any{"error": err})
		return entity.Document{}, fmt.Errorf("read DSL: %w", err)
	}
	dec := xml.NewDecoder(bytes.NewReader(data))
	var stack []*entity.Node
	var root *entity.Node

	for {
		offset := int(dec.InputOffset())
		tok, err := dec.Token()
		if err == io.EOF {
			logger.DEBUG(IUPP002, "branch EOF")
			break
		}
		if err != nil {
			logger.ERROR(IUPP003, "token failed", map[string]any{"offset": offset, "error": err})
			return entity.Document{}, &entity.ParseError{Position: positionAt(data, offset), Err: fmt.Errorf("parse xml-like token: %w", err)}
		}

		switch t := tok.(type) {
		case xml.StartElement:
			logger.DEBUG(IUPP004, "branch start element", map[string]any{"tag": t.Name.Local})
			node := &entity.Node{Tag: t.Name.Local, Attrs: map[string]string{}, Position: positionAt(data, offset)}
			for _, a := range t.Attr {
				node.Attrs[a.Name.Local] = a.Value
			}
			if node.Tag == "item" {
				if err := validateItemNode(node); err != nil {
					logger.ERROR(IUPP005, "item validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <item>: %w", err)}
				}
			}
			if node.Tag == "connection" {
				if err := validateConnectionNode(node); err != nil {
					logger.ERROR(IUPP006, "connection validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <connection>: %w", err)}
				}
			}
			if node.Tag == "generic-group" {
				if err := validateGenericGroupNode(node); err != nil {
					logger.ERROR(IUPP007, "generic group validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <generic-group>: %w", err)}
				}
			}
			if len(stack) == 0 {
				logger.DEBUG(IUPP008, "branch root assignment", map[string]any{"tag": node.Tag})
				root = node
			} else {
				logger.DEBUG(IUPP009, "branch child append", map[string]any{"tag": node.Tag})
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, node)
			}
			stack = append(stack, node)
		case xml.CharData:
			if len(stack) == 0 {
				logger.DEBUG(IUPP010, "branch empty stack char data")
				continue
			}
			text := strings.TrimSpace(string(t))
			cur := stack[len(stack)-1]
			cur.TextRuns = append(cur.TextRuns, entity.TextRun{Text: string(t), Position: positionAt(data, offset)})
			if text != "" {
				if cur.Text == "" {
					logger.DEBUG(IUPP011, "branch text assignment", map[string]any{"tag": cur.Tag})
					cur.Text = text
				} else {
					logger.DEBUG(IUPP012, "branch text append", map[string]any{"tag": cur.Tag})
					cur.Text += " " + text
				}
			}
		case xml.EndElement:
			if len(stack) == 0 {
				logger.DEBUG(IUPP013, "branch unexpected closing tag", map[string]any{"tag": t.Name.Local})
				return entity.Document{}, &entity.ParseError{Position: positionAt(data, offset), Err: fmt.Errorf("unexpected closing tag: %s", t.Name.Local)}
			}
			stack = stack[:len(stack)-1]
		}
	}

	if root == nil {
		logger.ERROR(IUPP014, "branch empty document")
		return entity.Document{}, &entity.ParseError{Position: entity.Position{Line: 1, Column: 1}, Err: fmt.Errorf("empty document")}
	}
	if root.Tag != "frame" {
		logger.ERROR(IUPP015, "branch invalid root", map[string]any{"tag": root.Tag})
		return entity.Document{}, &entity.ParseError{Position: root.Position, Err: fmt.Errorf("root tag must be <frame>, got <%s>", root.Tag)}
	}
	if err := expandConnectionShorthands(root, data); err != nil {
		logger.ERROR(IUPP016, "expand connection shorthands failed", map[string]any{"error": err})
		return entity.Document{}, err
	}

	return entity.Document{Root: root}, nil
}

func validateGenericGroupNode(node *entity.Node) error {
	iconID := strings.TrimSpace(node.Attrs["icon-id"])
	if iconID == "" {
		logger.DEBUG(IUPVGGN001, "branch empty icon ID")
		return nil
	}
	for _, ch := range iconID {
		if ch < '0' || ch > '9' {
			logger.ERROR(IUPVGGN002, "branch invalid icon ID", map[string]any{"iconID": iconID})
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
			logger.DEBUG(IUPECS001, "branch item", map[string]any{"tag": node.Tag})
			id := strings.TrimSpace(node.Attrs["id"])
			if id != "" {
				logger.DEBUG(IUPECS002, "branch item ID", map[string]any{"id": id})
				aliases[id] = id
			}
			for _, key := range []string{"name", "ref"} {
				alias := strings.TrimSpace(node.Attrs[key])
				if alias == "" {
					logger.DEBUG(IUPECS003, "branch empty alias", map[string]any{"key": key})
					continue
				}
				if id == "" {
					logger.DEBUG(IUPECS004, "branch alias without ID", map[string]any{"key": key, "alias": alias})
					return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<item %s=%q> requires a non-empty id", key, alias)}
				}
				if previous, exists := aliases[alias]; exists && previous != id {
					logger.DEBUG(IUPECS005, "branch duplicate alias", map[string]any{"alias": alias, "previous": previous, "id": id})
					return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate item reference %q", alias)}
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
		logger.ERROR(IUPECS006, "collect failed", map[string]any{"error": err})
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
				logger.DEBUG(IUPECS007, "branch empty line")
				continue
			}
			match := connectionShorthandPattern.FindStringSubmatch(trimmed)
			if match == nil {
				if strings.Contains(trimmed, "---") || strings.Contains(trimmed, "==>") {
					logger.ERROR(IUPECS008, "branch invalid shorthand", map[string]any{"line": trimmed})
					return &entity.ParseError{Position: position, Err: fmt.Errorf("invalid connection shorthand %q; expected 'source --- destination' or 'source ==> destination'", trimmed)}
				}
				logger.DEBUG(IUPECS009, "branch non shorthand", map[string]any{"line": trimmed})
				continue
			}
			src, ok := aliases[match[1]]
			if !ok || src == "" {
				logger.ERROR(IUPECS010, "branch missing source", map[string]any{"source": match[1]})
				return &entity.ParseError{Position: position, Err: fmt.Errorf("connection shorthand source %q does not match an <item name=...>, <item ref=...>, or item ID", match[1])}
			}
			dst, ok := aliases[match[3]]
			if !ok || dst == "" {
				logger.ERROR(IUPECS011, "branch missing destination", map[string]any{"destination": match[3]})
				return &entity.ParseError{Position: position, Err: fmt.Errorf("connection shorthand destination %q does not match an <item name=...>, <item ref=...>, or item ID", match[3])}
			}
			kind := "route"
			if match[2] == "==>" {
				logger.DEBUG(IUPECS012, "branch traffic")
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
		logger.DEBUG(IUPPA001, "branch negative offset", map[string]any{"offset": offset})
		offset = 0
	}
	if offset > len(data) {
		logger.DEBUG(IUPPA002, "branch overflow offset", map[string]any{"offset": offset, "length": len(data)})
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
		logger.DEBUG(IUPVIN001, "branch spacer item")
		return nil // spacer item — no id required
	}
	if strings.Contains(id, ",") {
		logger.DEBUG(IUPVIN002, "branch comma ID", map[string]any{"id": id})
		return fmt.Errorf("<item id=%q> must contain a single ID; use separate <item> tags for multiple services", id)
	}
	for _, ch := range strings.TrimSpace(id) {
		if ch < '0' || ch > '9' {
			logger.DEBUG(IUPVIN003, "branch non numeric ID", map[string]any{"id": id})
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
		logger.ERROR(IUPVCN001, "branch missing source")
		return fmt.Errorf("<connection> requires a src attribute")
	}
	if !hasDst || strings.TrimSpace(dst) == "" {
		logger.ERROR(IUPVCN002, "branch missing destination")
		return fmt.Errorf("<connection> requires a dst attribute")
	}
	for _, ch := range strings.TrimSpace(src) {
		if ch < '0' || ch > '9' {
			logger.DEBUG(IUPVCN003, "branch non numeric source", map[string]any{"src": src})
			return fmt.Errorf("<connection src=%q> must be a positive integer", src)
		}
	}
	for _, ch := range strings.TrimSpace(dst) {
		if ch < '0' || ch > '9' {
			logger.DEBUG(IUPVCN004, "branch non numeric destination", map[string]any{"dst": dst})
			return fmt.Errorf("<connection dst=%q> must be a positive integer", dst)
		}
	}
	return nil
}
