package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type endpointRefV1EngineParseReference struct {
	key      string
	frameID  string
	position entity.Position
}

type connectionRefV1EngineParseReference struct {
	node    *entity.Node
	frameID string
}

func validateConnectionReferencesV1EngineParseReference(root *entity.Node) error {
	qualified := map[string]endpointRefV1EngineParseReference{}
	local := map[string]map[string]endpointRefV1EngineParseReference{}
	ambiguous := map[string]map[string]bool{}
	frameRefs := map[string]endpointRefV1EngineParseReference{}
	endpointFrameByKey := map[string]string{}
	var connections []connectionRefV1EngineParseReference

	var walk func(node, parent, grandparent *entity.Node, currentFrameID string) error
	walk = func(node, parent, grandparent *entity.Node, currentFrameID string) error {
		if node.Tag == "frame" {
			currentFrameID = strings.TrimSpace(node.Attrs["id"])
			if strings.Contains(currentFrameID, ".") {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<frame id=%q> must not contain '.' because dots delimit frameId.id references", currentFrameID)}
			}
		}
		if nodeConnectableByIDV1EngineParseNode(node) {
			id := strings.TrimSpace(node.Attrs["id"])
			if node.Tag != "item" && strings.Contains(id, ".") {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<%s id=%q> must not contain '.' because dots delimit frameId.id references", node.Tag, id)}
			}
			key := strings.TrimSpace(node.Attrs[internalConnectionKeyAttrV1EngineParseDocument])
			refFrameID := currentFrameID
			if node.Tag == "frame" {
				refFrameID = id
			}
			ref := endpointRefV1EngineParseReference{key: key, frameID: refFrameID, position: node.Position}
			if key != "" && refFrameID != "" {
				endpointFrameByKey[key] = refFrameID
			}
			if node.Tag == "frame" {
				if _, exists := frameRefs[id]; exists {
					return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate frame id %q", id)}
				}
				frameRefs[id] = ref
			} else if id != "" {
				if local[refFrameID] == nil {
					local[refFrameID] = map[string]endpointRefV1EngineParseReference{}
					ambiguous[refFrameID] = map[string]bool{}
				}
				for index, token := range []string{id, strings.TrimSpace(node.Attrs["name"]), strings.TrimSpace(node.Attrs["ref"])} {
					if token == "" {
						continue
					}
					if _, exists := local[refFrameID][token]; exists && index != 0 {
						return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate connection reference %q in frame %q", token, refFrameID)}
					}
					if _, exists := local[refFrameID][token]; exists {
						if node.Tag != "item" {
							return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate frame reference id %q", token)}
						}
						delete(local[refFrameID], token)
						delete(qualified, refFrameID+"."+token)
						ambiguous[refFrameID][token] = true
						continue
					}
					if ambiguous[refFrameID][token] {
						continue
					}
					local[refFrameID][token] = ref
					qualified[refFrameID+"."+token] = ref
				}
			}
		}
		if node.Tag == "connections" && (parent == nil || parent.Tag != "frame") {
			return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<connections> must be a direct child of <frame>")}
		}
		if node.Tag == "connections" {
			for _, child := range node.Children {
				if child.Tag != "connection" {
					return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<connections> may only contain <connection> children, got <%s>", child.Tag)}
				}
			}
		}
		if node.Tag == "connection" {
			direct := parent != nil && parent.Tag == "frame"
			grouped := parent != nil && parent.Tag == "connections" && grandparent != nil && grandparent.Tag == "frame"
			if !direct && !grouped {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<connection> must be a direct child of <frame> or <connections>")}
			}
			connections = append(connections, connectionRefV1EngineParseReference{node: node, frameID: currentFrameID})
		}
		for _, child := range node.Children {
			if err := walk(child, node, parent, currentFrameID); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(root, nil, nil, ""); err != nil {
		return err
	}

	for _, connection := range connections {
		conn := connection.node
		for _, endpoint := range []struct{ attr, keyAttr string }{{"src", internalConnectionSrcKeyAttrV1EngineParseDocument}, {"dst", internalConnectionDstKeyAttrV1EngineParseDocument}} {
			token := strings.TrimSpace(conn.Attrs[endpoint.attr])
			if key := strings.TrimSpace(conn.Attrs[endpoint.keyAttr]); key != "" {
				setConnectionEndpointFrameV1EngineParseReference(conn, endpoint.attr, endpointFrameByKey[key])
				continue
			}
			if ambiguousEndpointV1EngineParseReference(token, connection.frameID, ambiguous) {
				return &entity.ParseError{Position: conn.Position, Err: fmt.Errorf("<connection %s=%q> is ambiguous because endpoint id=%q appears 2 times; use a unique name or ref", endpoint.attr, token, token)}
			}
			ref, exists := resolveEndpointV1EngineParseReference(token, connection.frameID, local, qualified, frameRefs)
			if !exists {
				return &entity.ParseError{Position: conn.Position, Err: fmt.Errorf("<connection %s=%q> does not match any connection endpoint id/name/ref in frame %q; use frameId.id for a cross-frame reference", endpoint.attr, token, connection.frameID)}
			}
			conn.Attrs[endpoint.keyAttr] = ref.key
			setConnectionEndpointFrameV1EngineParseReference(conn, endpoint.attr, ref.frameID)
		}
		if conn.Attrs[internalConnectionSrcFrameAttrV1EngineParseDocument] != conn.Attrs[internalConnectionDstFrameAttrV1EngineParseDocument] {
			conn.Attrs[internalConnectionCrossFrameAttrV1EngineParseDocument] = "true"
		}
		if conn.Attrs[internalConnectionCrossFrameAttrV1EngineParseDocument] != "true" {
			for _, attribute := range []string{"src-frame-side", "src-frame-anchor", "dst-frame-side", "dst-frame-anchor"} {
				if strings.TrimSpace(conn.Attrs[attribute]) == "" {
					continue
				}
				return &entity.ParseError{Position: conn.Position, Err: fmt.Errorf("<connection %s=%q> is only valid for a cross-frame connection", attribute, conn.Attrs[attribute])}
			}
		}
	}
	return nil
}

func ambiguousEndpointV1EngineParseReference(token, currentFrameID string, ambiguous map[string]map[string]bool) bool {
	if strings.Contains(token, ".") {
		parts := strings.SplitN(token, ".", 2)
		return ambiguous[parts[0]][parts[1]]
	}
	return ambiguous[currentFrameID][token]
}

func resolveEndpointV1EngineParseReference(token, currentFrameID string, local map[string]map[string]endpointRefV1EngineParseReference, qualified, frames map[string]endpointRefV1EngineParseReference) (endpointRefV1EngineParseReference, bool) {
	if strings.Contains(token, ".") {
		ref, exists := qualified[token]
		return ref, exists
	}
	if ref, exists := local[currentFrameID][token]; exists {
		return ref, true
	}
	ref, exists := frames[token]
	return ref, exists
}

func setConnectionEndpointFrameV1EngineParseReference(conn *entity.Node, endpoint, frameID string) {
	if conn == nil || frameID == "" {
		return
	}
	if endpoint == "src" {
		conn.Attrs[internalConnectionSrcFrameAttrV1EngineParseDocument] = frameID
	} else if endpoint == "dst" {
		conn.Attrs[internalConnectionDstFrameAttrV1EngineParseDocument] = frameID
	}
}
