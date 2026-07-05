package usecase

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

var connectionShorthandPattern = regexp.MustCompile(`^([A-Za-z0-9_.:-]+)\s*(---|==>)\s*([A-Za-z0-9_.:-]+)$`)

const (
	internalConnectionKeyAttr    = "_xaligoConnectionKey"
	internalConnectionSrcKeyAttr = "_xaligoConnectionSrcKey"
	internalConnectionDstKeyAttr = "_xaligoConnectionDstKey"
)

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
	IUPVCN005  = share.NewMCode("IUPVCN-005", "Validate connection nested branch")
	IUPVCN006  = share.NewMCode("IUPVCN-006", "Validate connection missing endpoint item branch")
	IUPVCN007  = share.NewMCode("IUPVCN-007", "Validate connection ambiguous endpoint item branch")
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
			if isConnectableFrameTag(node.Tag) {
				if err := validateConnectableFrameNode(node); err != nil {
					logger.ERROR(IUPP007, "connectable frame validation failed", map[string]any{"error": err})
					return entity.Document{}, &entity.ParseError{Position: node.Position, Err: err}
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
	assignConnectionKeys(root)
	if err := expandConnectionShorthands(root, data); err != nil {
		logger.ERROR(IUPP016, "expand connection shorthands failed", map[string]any{"error": err})
		return entity.Document{}, err
	}
	if err := normalizeConnectionEndpointTags(root); err != nil {
		logger.ERROR(IUPP006, "connection endpoint tag normalization failed", map[string]any{"error": err})
		return entity.Document{}, err
	}
	if err := validateConnectionReferences(root); err != nil {
		logger.ERROR(IUPP006, "connection reference validation failed", map[string]any{"error": err})
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

func validateConnectableFrameNode(node *entity.Node) error {
	id := strings.TrimSpace(node.Attrs["id"])
	if id == "" {
		return fmt.Errorf("<%s> requires a non-empty id attribute", node.Tag)
	}
	if strings.ContainsAny(id, " \t\r\n") {
		return fmt.Errorf("<%s id=%q> must not contain whitespace", node.Tag, id)
	}
	return nil
}

func assignConnectionKeys(root *entity.Node) {
	next := 1
	var walk func(*entity.Node)
	walk = func(node *entity.Node) {
		if nodeConnectableByID(node) {
			node.Attrs[internalConnectionKeyAttr] = node.Tag + "-" + strconv.Itoa(next)
			next++
		}
		for _, child := range node.Children {
			walk(child)
		}
	}
	walk(root)
}

func nodeConnectableByID(node *entity.Node) bool {
	if node == nil || strings.TrimSpace(node.Attrs["id"]) == "" {
		return false
	}
	return node.Tag == "item" || isConnectableFrameTag(node.Tag)
}

func isConnectableFrameTag(tag string) bool {
	if tag == "rectangle" || tag == "port" {
		return true
	}
	_, isGroup := awsGroups[tag]
	return isGroup
}

func expandConnectionShorthands(root *entity.Node, data []byte) error {
	aliases := map[string]string{}
	duplicateAliases := map[string]bool{}
	var collect func(*entity.Node) error
	collect = func(node *entity.Node) error {
		if nodeConnectableByID(node) {
			logger.DEBUG(IUPECS001, "branch item", map[string]any{"tag": node.Tag})
			id := strings.TrimSpace(node.Attrs["id"])
			key := strings.TrimSpace(node.Attrs[internalConnectionKeyAttr])
			if id != "" {
				logger.DEBUG(IUPECS002, "branch item ID", map[string]any{"id": id})
				if _, exists := aliases[id]; exists {
					duplicateAliases[id] = true
				} else {
					aliases[id] = key
				}
			}
			for _, key := range []string{"name", "ref"} {
				alias := strings.TrimSpace(node.Attrs[key])
				if alias == "" {
					logger.DEBUG(IUPECS003, "branch empty alias", map[string]any{"key": key})
					continue
				}
				if _, exists := aliases[alias]; exists || duplicateAliases[alias] {
					logger.DEBUG(IUPECS005, "branch duplicate alias", map[string]any{"alias": alias, "id": id})
					return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate connection reference %q", alias)}
				}
				aliases[alias] = node.Attrs[internalConnectionKeyAttr]
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
			if !ok || src == "" || duplicateAliases[match[1]] {
				logger.ERROR(IUPECS010, "branch missing source", map[string]any{"source": match[1]})
				return &entity.ParseError{Position: position, Err: fmt.Errorf("connection shorthand source %q does not match an <item> or group id/name/ref", match[1])}
			}
			dst, ok := aliases[match[3]]
			if !ok || dst == "" || duplicateAliases[match[3]] {
				logger.ERROR(IUPECS011, "branch missing destination", map[string]any{"destination": match[3]})
				return &entity.ParseError{Position: position, Err: fmt.Errorf("connection shorthand destination %q does not match an <item> or group id/name/ref", match[3])}
			}
			kind := "route"
			if match[2] == "==>" {
				logger.DEBUG(IUPECS012, "branch traffic")
				kind = "traffic"
			}
			root.Children = append(root.Children, &entity.Node{
				Tag: "connection",
				Attrs: map[string]string{
					"src":                        match[1],
					"dst":                        match[3],
					"kind":                       kind,
					internalConnectionSrcKeyAttr: src,
					internalConnectionDstKeyAttr: dst,
				},
				Position: position,
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

// validateConnectionNode validates immediately knowable <connection> attributes.
// Endpoint presence is checked after child <src>/<dst> tags have been parsed.
func validateConnectionNode(node *entity.Node) error {
	return validateConnectionSideAttrs(node)
}

func normalizeConnectionEndpointTags(root *entity.Node) error {
	var walk func(node *entity.Node) error
	walk = func(node *entity.Node) error {
		if node == nil {
			return nil
		}
		if node.Tag == "connection" {
			if err := normalizeConnectionEndpointTag(node, "src"); err != nil {
				return err
			}
			if err := normalizeConnectionEndpointTag(node, "dst"); err != nil {
				return err
			}
			if err := validateConnectionEndpointPresence(node); err != nil {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <connection>: %w", err)}
			}
			if err := validateConnectionSideAttrs(node); err != nil {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <connection>: %w", err)}
			}
		}
		for _, child := range node.Children {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(root)
}

func normalizeConnectionEndpointTag(conn *entity.Node, endpoint string) error {
	for _, child := range conn.Children {
		if strings.ToLower(strings.TrimSpace(child.Tag)) != endpoint {
			continue
		}
		if strings.TrimSpace(conn.Attrs[endpoint]) == "" {
			if token := connectionEndpointTagToken(child); token != "" {
				conn.Attrs[endpoint] = token
			}
		}
		sideValue := strings.TrimSpace(child.Attrs["side"])
		anchorValue := connectionEndpointTagAnchor(child)
		if spec, ok, err := parseConnectionAnchorSpec(sideValue, anchorValue); err != nil {
			return &entity.ParseError{Position: child.Position, Err: err}
		} else if ok {
			sideAttr := endpoint + "-side"
			if strings.TrimSpace(conn.Attrs[sideAttr]) == "" {
				conn.Attrs[sideAttr] = string(spec.side)
			}
			anchorAttr := endpoint + "-anchor"
			if spec.hasSlot && strings.TrimSpace(conn.Attrs[anchorAttr]) == "" {
				conn.Attrs[anchorAttr] = spec.String()
			}
		}
	}
	return nil
}

func connectionEndpointTagToken(node *entity.Node) string {
	for _, attr := range []string{"id", "ref", "name", "target"} {
		if value := strings.TrimSpace(node.Attrs[attr]); value != "" {
			return value
		}
	}
	return strings.TrimSpace(node.Text)
}

func connectionEndpointTagAnchor(node *entity.Node) string {
	for _, attr := range []string{"anchor", "position", "slot"} {
		if value := strings.TrimSpace(node.Attrs[attr]); value != "" {
			return value
		}
	}
	return ""
}

func validateConnectionEndpointPresence(node *entity.Node) error {
	if strings.TrimSpace(node.Attrs["src"]) == "" {
		logger.ERROR(IUPVCN001, "branch missing source")
		return fmt.Errorf("<connection> requires a src attribute")
	}
	if strings.TrimSpace(node.Attrs["dst"]) == "" {
		logger.ERROR(IUPVCN002, "branch missing destination")
		return fmt.Errorf("<connection> requires a dst attribute")
	}
	return nil
}

func validateConnectionSideAttrs(node *entity.Node) error {
	for _, endpoint := range []string{"src", "dst"} {
		sideAttr := endpoint + "-side"
		anchorAttr := endpoint + "-anchor"
		sideValue := strings.TrimSpace(node.Attrs[sideAttr])
		anchorValue := strings.TrimSpace(node.Attrs[anchorAttr])
		if sideValue != "" || anchorValue != "" {
			spec, ok, err := parseConnectionAnchorSpec(sideValue, anchorValue)
			if err != nil {
				return err
			}
			if ok {
				node.Attrs[sideAttr] = string(spec.side)
				if spec.hasSlot {
					node.Attrs[anchorAttr] = spec.String()
				}
			}
		}
	}
	return nil
}

type connectionAnchorSpec struct {
	side    side
	slot    int
	hasSlot bool
}

func (spec connectionAnchorSpec) String() string {
	return string(spec.side) + "-" + strconv.Itoa(spec.slot+1)
}

func parseConnectionAnchorSpec(sideValue, anchorValue string) (connectionAnchorSpec, bool, error) {
	sideValue = strings.TrimSpace(sideValue)
	anchorValue = strings.TrimSpace(anchorValue)
	var spec connectionAnchorSpec
	if sideValue != "" {
		s, ok := normalizeConnectionSide(sideValue)
		if !ok {
			return spec, false, fmt.Errorf("<connection %s=%q> must be one of top, right, bottom, or left", "side", sideValue)
		}
		spec.side = s
	}
	if anchorValue == "" {
		return spec, sideValue != "", nil
	}
	if anchor, ok := parseFullConnectionAnchor(anchorValue); ok {
		if sideValue != "" && spec.side != anchor.side {
			return spec, false, fmt.Errorf("<connection anchor=%q> conflicts with side=%q", anchorValue, sideValue)
		}
		return anchor, true, nil
	}
	if s, ok := normalizeConnectionSide(anchorValue); ok {
		if sideValue != "" && spec.side != s {
			return spec, false, fmt.Errorf("<connection anchor=%q> conflicts with side=%q", anchorValue, sideValue)
		}
		return connectionAnchorSpec{side: s}, true, nil
	}
	if sideValue == "" {
		return spec, false, fmt.Errorf("<connection anchor=%q> must be SIDE-POSITION or be paired with side", anchorValue)
	}
	slot, ok := parseConnectionAnchorSlot(anchorValue, spec.side)
	if !ok {
		return spec, false, fmt.Errorf("<connection anchor=%q> position must be 1, 2, 3, 4, 5, start, near, center, far, or end", anchorValue)
	}
	spec.slot = slot
	spec.hasSlot = true
	return spec, true, nil
}

func parseFullConnectionAnchor(value string) (connectionAnchorSpec, bool) {
	normalized := strings.ToLower(strings.TrimSpace(value))
	normalized = strings.ReplaceAll(normalized, "_", "-")
	normalized = strings.ReplaceAll(normalized, ":", "-")
	normalized = strings.ReplaceAll(normalized, ".", "-")
	parts := strings.Split(normalized, "-")
	if len(parts) >= 2 {
		s, ok := normalizeConnectionSide(parts[0])
		if !ok {
			return connectionAnchorSpec{}, false
		}
		slot, ok := parseConnectionAnchorSlot(strings.Join(parts[1:], "-"), s)
		if !ok {
			return connectionAnchorSpec{}, false
		}
		return connectionAnchorSpec{side: s, slot: slot, hasSlot: true}, true
	}
	for _, prefix := range []string{"top", "right", "bottom", "left"} {
		if strings.HasPrefix(normalized, prefix) && len(normalized) > len(prefix) {
			s, _ := normalizeConnectionSide(prefix)
			slot, ok := parseConnectionAnchorSlot(normalized[len(prefix):], s)
			if ok {
				return connectionAnchorSpec{side: s, slot: slot, hasSlot: true}, true
			}
		}
	}
	return connectionAnchorSpec{}, false
}

func parseConnectionAnchorSlot(value string, s side) (int, bool) {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
	case "0":
		return 0, true
	case "1":
		return 0, true
	case "2":
		return 1, true
	case "3":
		return 2, true
	case "4":
		return 3, true
	case "5":
		return 4, true
	case "start", "near":
		return 0, true
	case "middle", "mid", "center", "centre":
		return 2, true
	case "far", "end":
		return 4, true
	case "left", "west":
		if s == sideTop || s == sideBottom {
			return 0, true
		}
	case "right", "east":
		if s == sideTop || s == sideBottom {
			return 4, true
		}
	case "top", "north":
		if s == sideLeft || s == sideRight {
			return 0, true
		}
	case "bottom", "south":
		if s == sideLeft || s == sideRight {
			return 4, true
		}
	}
	return 0, false
}

func normalizeConnectionSide(value string) (side, bool) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "top", "north", "n":
		return sideTop, true
	case "right", "east", "e":
		return sideRight, true
	case "bottom", "south", "s":
		return sideBottom, true
	case "left", "west", "w":
		return sideLeft, true
	default:
		return "", false
	}
}

func validateConnectionReferences(root *entity.Node) error {
	type endpointRef struct {
		key      string
		position entity.Position
	}
	endpointsByID := map[string][]endpointRef{}
	endpointsByAlias := map[string]endpointRef{}
	frameIDs := map[string]endpointRef{}
	var connections []*entity.Node

	var walk func(node, parent, grandparent *entity.Node) error
	walk = func(node, parent, grandparent *entity.Node) error {
		if nodeConnectableByID(node) {
			id := strings.TrimSpace(node.Attrs["id"])
			key := strings.TrimSpace(node.Attrs[internalConnectionKeyAttr])
			if id != "" {
				ref := endpointRef{key: key, position: node.Position}
				if isConnectableFrameTag(node.Tag) {
					if _, exists := frameIDs[id]; exists {
						logger.ERROR(IUPVCN007, "branch duplicate frame endpoint ID", map[string]any{"id": id})
						return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate frame reference id %q", id)}
					}
					frameIDs[id] = ref
				}
				endpointsByID[id] = append(endpointsByID[id], ref)
				for _, attr := range []string{"name", "ref"} {
					alias := strings.TrimSpace(node.Attrs[attr])
					if alias == "" {
						continue
					}
					if _, exists := endpointsByAlias[alias]; exists {
						logger.ERROR(IUPVCN007, "branch duplicate endpoint alias", map[string]any{"alias": alias})
						return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate connection reference %q", alias)}
					}
					endpointsByAlias[alias] = ref
				}
			}
		}
		if node.Tag == "connections" && parent != root {
			logger.ERROR(IUPVCN005, "branch nested connections", map[string]any{"tag": parent.Tag})
			return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<connections> must be a direct child of <frame>")}
		}
		if node.Tag == "connection" {
			grouped := parent != nil && parent.Tag == "connections" && grandparent == root
			if parent != root && !grouped {
				logger.ERROR(IUPVCN005, "branch nested connection", map[string]any{"tag": parent.Tag})
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<connection> must be a direct child of <frame> or <connections>")}
			}
			connections = append(connections, node)
		}
		for _, child := range node.Children {
			if err := walk(child, node, parent); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(root, nil, nil); err != nil {
		return err
	}

	for _, conn := range connections {
		for _, endpoint := range []struct {
			attr    string
			keyAttr string
		}{
			{"src", internalConnectionSrcKeyAttr},
			{"dst", internalConnectionDstKeyAttr},
		} {
			token := strings.TrimSpace(conn.Attrs[endpoint.attr])
			if strings.TrimSpace(conn.Attrs[endpoint.keyAttr]) != "" {
				continue
			}
			if refs := endpointsByID[token]; len(refs) > 0 {
				if len(refs) > 1 {
					logger.ERROR(IUPVCN007, "branch ambiguous endpoint item", map[string]any{"attr": endpoint.attr, "id": token, "count": len(refs)})
					return &entity.ParseError{Position: conn.Position, Err: fmt.Errorf("<connection %s=%q> is ambiguous because endpoint id=%q appears %d times; use a unique name or ref", endpoint.attr, token, token, len(refs))}
				}
				conn.Attrs[endpoint.keyAttr] = refs[0].key
				continue
			}
			if ref, ok := endpointsByAlias[token]; ok {
				conn.Attrs[endpoint.keyAttr] = ref.key
				continue
			}
			logger.ERROR(IUPVCN006, "branch missing endpoint item", map[string]any{"attr": endpoint.attr, "token": token})
			return &entity.ParseError{Position: conn.Position, Err: fmt.Errorf("<connection %s=%q> does not match any <item> or group id/name/ref", endpoint.attr, token)}
		}
	}
	return nil
}
