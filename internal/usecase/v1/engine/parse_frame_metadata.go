package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const internalFrameContentVersionAttrV1EngineParseFrameMetadata = "_xaligoFrameContentVersion"

// normalizeFrameMetadataV1EngineParseFrameMetadata validates metadata attached
// to physical page frames and marks child-frame versions as page content
// versions. A legacy root frame keeps its version attribute exclusively as the
// V1 language version.
func normalizeFrameMetadataV1EngineParseFrameMetadata(root, envelope *entity.Node) error {
	if root == nil {
		return nil
	}
	if envelope == nil {
		envelope = root
	}
	if err := rejectMisplacedFrameContentVersionsV1EngineParseFrameMetadata(envelope, root); err != nil {
		return err
	}
	allowedMetadata := map[*entity.Node]bool{}
	switch root.Tag {
	case "frames":
		for _, frame := range root.Children {
			if frame.Tag != "frame" {
				continue
			}
			if err := normalizeFrameMetadataNodeV1EngineParseFrameMetadata(frame, true, allowedMetadata); err != nil {
				return err
			}
		}
	case "frame":
		if err := normalizeFrameMetadataNodeV1EngineParseFrameMetadata(root, false, allowedMetadata); err != nil {
			return err
		}
	}
	return rejectMisplacedFrameMetadataV1EngineParseFrameMetadata(root, allowedMetadata)
}

// rejectMisplacedFrameContentVersionsV1EngineParseFrameMetadata closes the
// streaming parser's intentionally deferred frame-under-frames exception. A
// content version belongs only to an identified physical page frame directly
// below the exact selected page <frames> node in the original envelope; a
// legacy root frame's version remains the V1 language version.
func rejectMisplacedFrameContentVersionsV1EngineParseFrameMetadata(envelope, pageRoot *entity.Node) error {
	allowed := map[*entity.Node]bool{}
	if envelope != nil {
		allowed[envelope] = true
	}
	if pageRoot != nil {
		allowed[pageRoot] = true
	}
	if pageRoot != nil && pageRoot.Tag == "frames" {
		for _, child := range pageRoot.Children {
			if child.Tag == "frame" {
				allowed[child] = true
			}
		}
	}

	var walk func(*entity.Node) error
	walk = func(node *entity.Node) error {
		if node == nil {
			return nil
		}
		if version, specified := node.Attrs["version"]; specified && node.Tag == "frame" && !allowed[node] {
			return &entity.ParseError{
				Position: node.Position,
				Err:      fmt.Errorf("<frame version=%q> is invalid: content version is only allowed on identified <frame> children directly under the document-root <frames>", version),
			}
		}
		for _, child := range node.Children {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(envelope)
}

func normalizeFrameMetadataNodeV1EngineParseFrameMetadata(frame *entity.Node, contentVersion bool, allowedMetadata map[*entity.Node]bool) error {
	if frame == nil {
		return nil
	}
	if contentVersion {
		if rawVersion, specified := frame.Attrs["version"]; specified {
			version := strings.TrimSpace(rawVersion)
			if version == "" {
				return &entity.ParseError{Position: frame.Position, Err: fmt.Errorf("<frame version> must be non-empty when used as page metadata")}
			}
			frame.Attrs[internalFrameContentVersionAttrV1EngineParseFrameMetadata] = version
		}
	}

	var metadata *entity.Node
	for _, child := range frame.Children {
		if child.Tag != "metadata" {
			continue
		}
		if metadata != nil {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<frame> may contain only one <metadata> child")}
		}
		metadata = child
		allowedMetadata[child] = true
	}
	if metadata == nil {
		return nil
	}
	if err := normalizeFrameMetadataStyleV1EngineParseFrameMetadata(metadata); err != nil {
		return &entity.ParseError{Position: metadata.Position, Err: err}
	}
	if strings.TrimSpace(metadata.Text) != "" {
		return &entity.ParseError{Position: metadata.Position, Err: fmt.Errorf("<metadata> may only contain <entry> children")}
	}
	for _, entry := range metadata.Children {
		if entry.Tag != "entry" {
			return &entity.ParseError{Position: entry.Position, Err: fmt.Errorf("<metadata> may only contain <entry> children, got <%s>", entry.Tag)}
		}
		key := strings.TrimSpace(entry.Attr("key"))
		if key == "" {
			return &entity.ParseError{Position: entry.Position, Err: fmt.Errorf("<entry> requires a non-empty key attribute")}
		}
		value := strings.TrimSpace(entry.Attr("value"))
		if value == "" {
			return &entity.ParseError{Position: entry.Position, Err: fmt.Errorf("<entry> requires a non-empty value attribute")}
		}
		if len(entry.Children) != 0 || strings.TrimSpace(entry.Text) != "" {
			return &entity.ParseError{Position: entry.Position, Err: fmt.Errorf("<entry> must be empty and declare key and value attributes")}
		}
		if rawBreakBefore, specified := entry.Attrs["break-before"]; specified {
			breakBefore := strings.ToLower(strings.TrimSpace(rawBreakBefore))
			if breakBefore != "true" && breakBefore != "false" {
				return &entity.ParseError{Position: entry.Position, Err: fmt.Errorf("<entry break-before=%q> must be true or false", rawBreakBefore)}
			}
			entry.Attrs["break-before"] = breakBefore
		}
		entry.Attrs["key"] = key
		entry.Attrs["value"] = value
	}
	return nil
}

func normalizeFrameMetadataStyleV1EngineParseFrameMetadata(metadata *entity.Node) error {
	if rawPosition, specified := metadata.Attrs["position"]; specified {
		position := strings.ToLower(strings.TrimSpace(rawPosition))
		if position != "top" && position != "bottom" {
			return fmt.Errorf("<metadata position=%q> must be top or bottom", rawPosition)
		}
		metadata.Attrs["position"] = position
	}
	if rawAlign, specified := metadata.Attrs["align"]; specified {
		align := strings.ToLower(strings.TrimSpace(rawAlign))
		if align != "left" && align != "center" && align != "right" {
			return fmt.Errorf("<metadata align=%q> must be left, center, or right", rawAlign)
		}
		metadata.Attrs["align"] = align
	}
	return normalizePresentationStyleV1EngineParsePresentationStyle(
		metadata,
		"color",
		"key-color",
		"background-color",
		"key-background-color",
		"border-color",
	)
}

func rejectMisplacedFrameMetadataV1EngineParseFrameMetadata(root *entity.Node, allowed map[*entity.Node]bool) error {
	var walk func(*entity.Node) error
	walk = func(node *entity.Node) error {
		if node == nil {
			return nil
		}
		if node.Tag == "metadata" && !allowed[node] {
			return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<metadata> must be a direct child of a page <frame>")}
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
