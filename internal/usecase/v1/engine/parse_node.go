package engine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func validateGenericGroupNodeV1EngineParseNode(node *entity.Node) error {
	iconID := strings.TrimSpace(node.Attrs["icon-id"])
	if iconID == "" {
		loggerV1EngineSharedLogging.DEBUG(IUPVGGN001V1EngineParseDocument, "branch empty icon ID")
		return nil
	}
	if !isPositiveCatalogIDV1EngineParseNode(iconID) {
		loggerV1EngineSharedLogging.ERROR(IUPVGGN002V1EngineParseDocument, "branch non-positive icon ID", map[string]any{"iconID": iconID})
		return fmt.Errorf("icon-id=%q must be a positive catalog ID", iconID)
	}
	return nil
}

func validateConnectableFrameNodeV1EngineParseNode(node *entity.Node) error {
	id := strings.TrimSpace(node.Attrs["id"])
	if id == "" {
		return fmt.Errorf("<%s> requires a non-empty id attribute", node.Tag)
	}
	if strings.ContainsAny(id, " \t\r\n") {
		return fmt.Errorf("<%s id=%q> must not contain whitespace", node.Tag, id)
	}
	return nil
}

func validateFrameHierarchyV1EngineParseNode(root *entity.Node) error {
	if root == nil || root.Tag != "frames" {
		return nil
	}
	ids := map[string]entity.Position{}
	for _, child := range root.Children {
		if child.Tag != "frame" {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<frames> may only contain <frame> children, got <%s>", child.Tag)}
		}
		id := strings.TrimSpace(child.Attrs["id"])
		if id == "" {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<frame> inside <frames> requires a non-empty id attribute")}
		}
		if strings.ContainsAny(id, " \t\r\n") {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<frame id=%q> must not contain whitespace", id)}
		}
		if _, exists := ids[id]; exists {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate frame id %q", id)}
		}
		ids[id] = child.Position
	}
	return nil
}

func validateRootVersionV1EngineParseNode(root *entity.Node) error {
	if root == nil {
		return nil
	}
	version, specified := root.Attrs["version"]
	if !specified || version == "1" {
		return nil
	}
	return fmt.Errorf("<%s version=%q> is not supported by the V1 engine; use version=\"1\"", root.Tag, version)
}

func normalizeEnvelopeV1EngineParseNode(envelope *entity.Node) (*entity.Node, *entity.Node, error) {
	var dataNode *entity.Node
	var framesNode *entity.Node
	for _, child := range envelope.Children {
		switch child.Tag {
		case "metadata", "imports", "styles":
			continue
		case "data":
			if dataNode != nil {
				return nil, nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<xaligo> may contain only one <data> child")}
			}
			dataNode = child
		case "frames":
			if framesNode != nil {
				return nil, nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<xaligo> must contain exactly one <frames> child")}
			}
			framesNode = child
		default:
			return nil, nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<xaligo> may only contain <metadata>, <imports>, <data>, <styles>, and <frames> children, got <%s>", child.Tag)}
		}
	}
	if framesNode == nil {
		return nil, nil, &entity.ParseError{Position: envelope.Position, Err: fmt.Errorf("<xaligo> must contain exactly one <frames> child")}
	}
	if _, exists := framesNode.Attrs["version"]; exists {
		return nil, nil, &entity.ParseError{Position: framesNode.Position, Err: fmt.Errorf("<frames> inside <xaligo> must not declare version; version belongs on <xaligo>")}
	}
	if version, exists := envelope.Attrs["version"]; exists {
		framesNode.Attrs["version"] = version
	}
	return framesNode, dataNode, nil
}

func implicitV1VersionWarningV1EngineParseNode(root *entity.Node) string {
	return fmt.Sprintf("V1 document version is implicit; add version=\"1\" to the <%s> root", root.Tag)
}

func legacyV1RootWarningV1EngineParseNode(root *entity.Node) string {
	return fmt.Sprintf("legacy V1 <%s> root is deprecated; wrap identified frames in <xaligo version=\"1\"><frames>...</frames></xaligo>", root.Tag)
}

func validateNestedVersionV1EngineParseNode(node, parent *entity.Node) error {
	if node == nil {
		return nil
	}
	version, specified := node.Attrs["version"]
	if !specified {
		return nil
	}
	if node.Tag == "frame" && parent != nil && parent.Tag == "frames" {
		return nil
	}
	return fmt.Errorf("<%s version=%q> is invalid: version is only allowed on the V1 document root", node.Tag, version)
}

func assignConnectionKeysV1EngineParseNode(root *entity.Node) {
	next := 1
	var walk func(*entity.Node)
	walk = func(node *entity.Node) {
		if nodeConnectableByIDV1EngineParseNode(node) {
			node.Attrs[internalConnectionKeyAttrV1EngineParseDocument] = node.Tag + "-" + strconv.Itoa(next)
			next++
		}
		for _, child := range node.Children {
			walk(child)
		}
	}
	walk(root)
}

func nodeConnectableByIDV1EngineParseNode(node *entity.Node) bool {
	if node == nil || strings.TrimSpace(node.Attrs["id"]) == "" {
		return false
	}
	return node.Tag == "item" || node.Tag == "frame" || node.Tag == "entity" || node.Tag == "table" || node.Tag == "database" || isConnectableFrameTagV1EngineParseNode(node.Tag)
}

func isConnectableFrameTagV1EngineParseNode(tag string) bool {
	if tag == "rectangle" || tag == "port" {
		return true
	}
	_, isGroup := awsGroupsV1EngineSceneTypes[tag]
	return isGroup
}

// validateItemNode ensures <item> carries at most one numeric id attribute.
// An empty (or absent) id is allowed — the item acts as a layout spacer.
func validateItemNodeV1EngineParseNode(node *entity.Node) error {
	for _, attr := range []string{"dx", "dy"} {
		value := strings.TrimSpace(node.Attrs[attr])
		if value == "" {
			continue
		}
		if _, err := strconv.ParseFloat(value, 64); err != nil {
			return fmt.Errorf("<item %s=%q> must be a number", attr, value)
		}
	}

	id, ok := node.Attrs["id"]
	if !ok || strings.TrimSpace(id) == "" {
		loggerV1EngineSharedLogging.DEBUG(IUPVIN001V1EngineParseDocument, "branch spacer item")
		return nil // spacer item — no id required
	}
	if strings.Contains(id, ",") {
		loggerV1EngineSharedLogging.DEBUG(IUPVIN002V1EngineParseDocument, "branch comma ID", map[string]any{"id": id})
		return fmt.Errorf("<item id=%q> must contain a single ID; use separate <item> tags for multiple services", id)
	}
	if !isPositiveCatalogIDV1EngineParseNode(strings.TrimSpace(id)) {
		loggerV1EngineSharedLogging.DEBUG(IUPVIN003V1EngineParseDocument, "branch non-positive ID", map[string]any{"id": id})
		return fmt.Errorf("<item id=%q> must be a positive integer", id)
	}
	return nil
}

func isPositiveCatalogIDV1EngineParseNode(value string) bool {
	if value == "" {
		return false
	}
	for _, character := range value {
		if character < '0' || character > '9' {
			return false
		}
	}
	parsed, err := strconv.ParseUint(value, 10, 31)
	return err == nil && parsed > 0
}
