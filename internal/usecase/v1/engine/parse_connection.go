package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

// validateConnectionNode validates immediately knowable <connection> and
// <connections> attributes. Endpoint presence is checked after child
// <src>/<dst> tags have been parsed.
func validateConnectionNodeV1EngineParseConnection(node *entity.Node) error {
	if node.Tag == "connection" {
		if err := validateConnectionSideAttrsV1EngineParseConnection(node); err != nil {
			return err
		}
	}
	return validateConnectionEnumAttrsV1EngineParseConnection(node)
}

func validateConnectionEnumAttrsV1EngineParseConnection(node *entity.Node) error {
	if color := strings.TrimSpace(node.Attrs["color"]); color != "" && !isConnectionColorV1EngineParseConnection(color) {
		return fmt.Errorf("<%s color=%q> must be a six-digit hexadecimal color such as #2563eb", node.Tag, color)
	}
	if err := validateConnectionEnumAttrV1EngineParseConnection(node, "kind", "connection", "route", "traffic"); err != nil {
		return err
	}
	if err := validateConnectionEnumAttrV1EngineParseConnection(node, "stroke-style", "solid", "dashed", "dotted"); err != nil {
		return err
	}
	for _, name := range []string{"start-arrowhead", "end-arrowhead", "arrowhead"} {
		if err := validateConnectionEnumAttrV1EngineParseConnection(node, name, "none", "arrow", "triangle", "stealth", "diamond", "oval"); err != nil {
			return err
		}
	}
	// When a V1 connector has an arrowhead, every renderer emits the small size.
	// Reject larger values rather than accepting syntax that the V1 scene and
	// shared plan cannot preserve.
	return validateConnectionEnumAttrV1EngineParseConnection(node, "arrowhead-size", "s")
}

// validateEffectiveRouteArrowheadsV1EngineParseConnection applies connection
// group defaults before enforcing V1's headless route contract. Scene
// construction uses the same connectionWithDefaults helper, so child semantic
// alias overrides remain identical in Parse, Validate, and Render.
func validateEffectiveRouteArrowheadsV1EngineParseConnection(root *entity.Node) error {
	var walk func(*entity.Node) error
	walk = func(node *entity.Node) error {
		if node == nil {
			return nil
		}
		if node.Tag == "frame" {
			for _, child := range node.Children {
				switch child.Tag {
				case "connection":
					if err := validateEffectiveRouteArrowheadsForConnectionV1EngineParseConnection(child); err != nil {
						return err
					}
				case "connections":
					defaults := connectionGroupDefaultsV1EngineParseConnection(child)
					for _, grouped := range child.Children {
						if grouped.Tag != "connection" {
							continue
						}
						effective := connectionWithDefaultsV1EngineParseConnection(grouped, defaults)
						if err := validateEffectiveRouteArrowheadsForConnectionV1EngineParseConnection(effective); err != nil {
							return err
						}
					}
				}
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

func validateEffectiveRouteArrowheadsForConnectionV1EngineParseConnection(conn *entity.Node) error {
	if conn == nil || !strings.EqualFold(strings.TrimSpace(conn.Attrs["kind"]), "route") {
		return nil
	}
	startArrowhead := strings.ToLower(strings.TrimSpace(conn.Attrs["start-arrowhead"]))
	endArrowhead := strings.ToLower(strings.TrimSpace(conn.Attrs["end-arrowhead"]))
	if endArrowhead == "" {
		endArrowhead = strings.ToLower(strings.TrimSpace(conn.Attrs["arrowhead"]))
	}

	effectiveHeads := make([]string, 0, 2)
	if startArrowhead != "" && startArrowhead != "none" {
		effectiveHeads = append(effectiveHeads, fmt.Sprintf("start-arrowhead=%q", startArrowhead))
	}
	if endArrowhead != "" && endArrowhead != "none" {
		effectiveHeads = append(effectiveHeads, fmt.Sprintf("end-arrowhead/arrowhead=%q", endArrowhead))
	}
	if len(effectiveHeads) == 0 {
		return nil
	}
	return &entity.ParseError{
		Position: conn.Position,
		Err:      fmt.Errorf("<connection kind=\"route\"> must be headless; effective %s must be \"none\"", strings.Join(effectiveHeads, " and ")),
	}
}

func connectionGroupDefaultsV1EngineParseConnection(group *entity.Node) map[string]string {
	defaults := map[string]string{}
	if group == nil {
		return defaults
	}
	for _, name := range []string{
		"arrowhead-size", "kind", "color", "stroke-width", "width", "stroke-style",
		"start-arrowhead", "end-arrowhead", "arrowhead", "scale", "coordinate-scale", "grid",
	} {
		if value := strings.TrimSpace(group.Attrs[name]); value != "" {
			defaults[name] = value
		}
	}
	return defaults
}

func connectionWithDefaultsV1EngineParseConnection(conn *entity.Node, defaults map[string]string) *entity.Node {
	if conn == nil || len(defaults) == 0 {
		return conn
	}
	// These pairs are alternate spellings of one semantic setting. If the
	// child supplies either spelling, inheriting the parent's other spelling
	// would incorrectly win in downstream canonical-first resolution.
	semanticAliases := [][]string{
		{"stroke-width", "width"},
		{"end-arrowhead", "arrowhead"},
		{"coordinate-scale", "scale"},
	}
	suppressedDefaults := map[string]struct{}{}
	for _, aliases := range semanticAliases {
		for _, name := range aliases {
			if _, exists := conn.Attrs[name]; !exists {
				continue
			}
			for _, alias := range aliases {
				suppressedDefaults[alias] = struct{}{}
			}
			break
		}
	}

	clone := *conn
	clone.Attrs = map[string]string{}
	for key, value := range defaults {
		if _, suppressed := suppressedDefaults[key]; suppressed {
			continue
		}
		clone.Attrs[key] = value
	}
	for key, value := range conn.Attrs {
		clone.Attrs[key] = value
	}
	return &clone
}

func isConnectionColorV1EngineParseConnection(value string) bool {
	if len(value) != 7 || value[0] != '#' {
		return false
	}
	for _, character := range value[1:] {
		if character >= '0' && character <= '9' || character >= 'a' && character <= 'f' || character >= 'A' && character <= 'F' {
			continue
		}
		return false
	}
	return true
}

func validateConnectionEnumAttrV1EngineParseConnection(node *entity.Node, name string, allowed ...string) error {
	value := strings.TrimSpace(node.Attrs[name])
	if value == "" {
		return nil
	}
	for _, candidate := range allowed {
		if strings.EqualFold(value, candidate) {
			return nil
		}
	}
	return fmt.Errorf("<%s %s=%q> must be one of %s", node.Tag, name, value, strings.Join(allowed, ", "))
}

func normalizeConnectionEndpointTagsV1EngineParseConnection(root *entity.Node) error {
	var walk func(node *entity.Node) error
	walk = func(node *entity.Node) error {
		if node == nil {
			return nil
		}
		if node.Tag == "connection" {
			if err := normalizeConnectionEndpointTagV1EngineParseConnection(node, "src"); err != nil {
				return err
			}
			if err := normalizeConnectionEndpointTagV1EngineParseConnection(node, "dst"); err != nil {
				return err
			}
			if err := validateConnectionEndpointPresenceV1EngineParseConnection(node); err != nil {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("parse <connection>: %w", err)}
			}
			if err := validateConnectionSideAttrsV1EngineParseConnection(node); err != nil {
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

func normalizeConnectionEndpointTagV1EngineParseConnection(conn *entity.Node, endpoint string) error {
	for _, child := range conn.Children {
		if strings.ToLower(strings.TrimSpace(child.Tag)) != endpoint {
			continue
		}
		if strings.TrimSpace(conn.Attrs[endpoint]) == "" {
			if token := connectionEndpointTagTokenV1EngineParseConnection(child); token != "" {
				conn.Attrs[endpoint] = token
			}
		}
		sideValue := strings.TrimSpace(child.Attrs["side"])
		anchorValue := connectionEndpointTagAnchorV1EngineParseConnection(child)
		if spec, ok, err := parseConnectionAnchorSpecV1EngineParseConnection(sideValue, anchorValue); err != nil {
			return &entity.ParseError{Position: child.Position, Err: err}
		} else if ok {
			sideAttr := endpoint + "-side"
			if strings.TrimSpace(conn.Attrs[sideAttr]) == "" {
				conn.Attrs[sideAttr] = string(spec.side)
			}
			anchorAttr := endpoint + "-anchor"
			if spec.hasSlot && strings.TrimSpace(conn.Attrs[anchorAttr]) == "" {
				conn.Attrs[anchorAttr] = spec.StringV1EngineParseConnection()
			}
		}
		frameSideValue := strings.TrimSpace(child.Attrs["frame-side"])
		frameAnchorValue := strings.TrimSpace(child.Attrs["frame-anchor"])
		if spec, ok, err := parseConnectionAnchorSpecV1EngineParseConnection(frameSideValue, frameAnchorValue); err != nil {
			return &entity.ParseError{Position: child.Position, Err: err}
		} else if ok {
			frameSideAttr := endpoint + "-frame-side"
			if strings.TrimSpace(conn.Attrs[frameSideAttr]) == "" {
				conn.Attrs[frameSideAttr] = string(spec.side)
			}
			frameAnchorAttr := endpoint + "-frame-anchor"
			if spec.hasSlot && strings.TrimSpace(conn.Attrs[frameAnchorAttr]) == "" {
				conn.Attrs[frameAnchorAttr] = spec.StringV1EngineParseConnection()
			}
		}
	}
	return nil
}

func connectionEndpointTagTokenV1EngineParseConnection(node *entity.Node) string {
	for _, attr := range []string{"id", "ref", "name", "target"} {
		if value := strings.TrimSpace(node.Attrs[attr]); value != "" {
			return value
		}
	}
	return strings.TrimSpace(node.Text)
}

func connectionEndpointTagAnchorV1EngineParseConnection(node *entity.Node) string {
	for _, attr := range []string{"anchor", "position", "slot"} {
		if value := strings.TrimSpace(node.Attrs[attr]); value != "" {
			return value
		}
	}
	return ""
}

func validateConnectionEndpointPresenceV1EngineParseConnection(node *entity.Node) error {
	if strings.TrimSpace(node.Attrs["src"]) == "" {
		loggerV1EngineSharedLogging.ERROR(IUPVCN001V1EngineParseDocument, "branch missing source")
		return fmt.Errorf("<connection> requires a src attribute")
	}
	if strings.TrimSpace(node.Attrs["dst"]) == "" {
		loggerV1EngineSharedLogging.ERROR(IUPVCN002V1EngineParseDocument, "branch missing destination")
		return fmt.Errorf("<connection> requires a dst attribute")
	}
	return nil
}

func validateConnectionSideAttrsV1EngineParseConnection(node *entity.Node) error {
	for _, endpoint := range []string{"src", "dst"} {
		for _, attributes := range [][2]string{
			{endpoint + "-side", endpoint + "-anchor"},
			{endpoint + "-frame-side", endpoint + "-frame-anchor"},
		} {
			sideAttr, anchorAttr := attributes[0], attributes[1]
			sideValue := strings.TrimSpace(node.Attrs[sideAttr])
			anchorValue := strings.TrimSpace(node.Attrs[anchorAttr])
			if sideValue == "" && anchorValue == "" {
				continue
			}
			spec, ok, err := parseConnectionAnchorSpecV1EngineParseConnection(sideValue, anchorValue)
			if err != nil {
				return err
			}
			if !ok {
				continue
			}
			node.Attrs[sideAttr] = string(spec.side)
			if spec.hasSlot {
				node.Attrs[anchorAttr] = spec.StringV1EngineParseConnection()
			}
		}
	}
	return nil
}

type connectionAnchorSpecV1EngineParseConnection struct {
	side    sideV1EngineRouteTypes
	slot    int
	hasSlot bool
}

func (spec connectionAnchorSpecV1EngineParseConnection) StringV1EngineParseConnection() string {
	return string(spec.side) + "-" + strconv.Itoa(spec.slot+1)
}

func parseConnectionAnchorSpecV1EngineParseConnection(sideValue, anchorValue string) (connectionAnchorSpecV1EngineParseConnection, bool, error) {
	sideValue = strings.TrimSpace(sideValue)
	anchorValue = strings.TrimSpace(anchorValue)
	var spec connectionAnchorSpecV1EngineParseConnection
	if sideValue != "" {
		s, ok := normalizeConnectionSideV1EngineParseConnection(sideValue)
		if !ok {
			return spec, false, fmt.Errorf("<connection %s=%q> must be one of top, right, bottom, or left", "side", sideValue)
		}
		spec.side = s
	}
	if anchorValue == "" {
		return spec, sideValue != "", nil
	}
	if anchor, ok := parseFullConnectionAnchorV1EngineParseConnection(anchorValue); ok {
		if sideValue != "" && spec.side != anchor.side {
			return spec, false, fmt.Errorf("<connection anchor=%q> conflicts with side=%q", anchorValue, sideValue)
		}
		return anchor, true, nil
	}
	if s, ok := normalizeConnectionSideV1EngineParseConnection(anchorValue); ok {
		if sideValue != "" && spec.side != s {
			return spec, false, fmt.Errorf("<connection anchor=%q> conflicts with side=%q", anchorValue, sideValue)
		}
		return connectionAnchorSpecV1EngineParseConnection{side: s}, true, nil
	}
	if sideValue == "" {
		return spec, false, fmt.Errorf("<connection anchor=%q> must be SIDE-POSITION or be paired with side", anchorValue)
	}
	slot, ok := parseConnectionAnchorSlotV1EngineParseConnection(anchorValue, spec.side)
	if !ok {
		return spec, false, fmt.Errorf("<connection anchor=%q> position must be 1, 2, 3, 4, 5, start, near, center, far, or end", anchorValue)
	}
	spec.slot = slot
	spec.hasSlot = true
	return spec, true, nil
}

func parseFullConnectionAnchorV1EngineParseConnection(value string) (connectionAnchorSpecV1EngineParseConnection, bool) {
	normalized := strings.ToLower(strings.TrimSpace(value))
	normalized = strings.ReplaceAll(normalized, "_", "-")
	normalized = strings.ReplaceAll(normalized, ":", "-")
	normalized = strings.ReplaceAll(normalized, ".", "-")
	parts := strings.Split(normalized, "-")
	if len(parts) >= 2 {
		s, ok := normalizeConnectionSideV1EngineParseConnection(parts[0])
		if !ok {
			return connectionAnchorSpecV1EngineParseConnection{}, false
		}
		slot, ok := parseConnectionAnchorSlotV1EngineParseConnection(strings.Join(parts[1:], "-"), s)
		if !ok {
			return connectionAnchorSpecV1EngineParseConnection{}, false
		}
		return connectionAnchorSpecV1EngineParseConnection{side: s, slot: slot, hasSlot: true}, true
	}
	for _, prefix := range []string{"top", "right", "bottom", "left"} {
		if strings.HasPrefix(normalized, prefix) && len(normalized) > len(prefix) {
			s, _ := normalizeConnectionSideV1EngineParseConnection(prefix)
			slot, ok := parseConnectionAnchorSlotV1EngineParseConnection(normalized[len(prefix):], s)
			if ok {
				return connectionAnchorSpecV1EngineParseConnection{side: s, slot: slot, hasSlot: true}, true
			}
		}
	}
	return connectionAnchorSpecV1EngineParseConnection{}, false
}

func parseConnectionAnchorSlotV1EngineParseConnection(value string, s sideV1EngineRouteTypes) (int, bool) {
	value = strings.ToLower(strings.TrimSpace(value))
	switch value {
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
	case "start":
		return 0, true
	case "near":
		return 1, true
	case "middle", "mid", "center", "centre":
		return 2, true
	case "far":
		return 3, true
	case "end":
		return 4, true
	case "left", "west":
		if s == sideTopV1EngineRouteTypes || s == sideBottomV1EngineRouteTypes {
			return 0, true
		}
	case "right", "east":
		if s == sideTopV1EngineRouteTypes || s == sideBottomV1EngineRouteTypes {
			return 4, true
		}
	case "top", "north":
		if s == sideLeftV1EngineRouteTypes || s == sideRightV1EngineRouteTypes {
			return 0, true
		}
	case "bottom", "south":
		if s == sideLeftV1EngineRouteTypes || s == sideRightV1EngineRouteTypes {
			return 4, true
		}
	}
	return 0, false
}

func normalizeConnectionSideV1EngineParseConnection(value string) (sideV1EngineRouteTypes, bool) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "top", "north", "n":
		return sideTopV1EngineRouteTypes, true
	case "right", "east", "e":
		return sideRightV1EngineRouteTypes, true
	case "bottom", "south", "s":
		return sideBottomV1EngineRouteTypes, true
	case "left", "west", "w":
		return sideLeftV1EngineRouteTypes, true
	default:
		return "", false
	}
}
