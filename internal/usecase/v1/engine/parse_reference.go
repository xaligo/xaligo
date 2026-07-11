package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func validateConnectionReferencesV1EngineParseReference(root *entity.Node) error {
	type endpointRef struct {
		key      string
		frameID  string
		position entity.Position
	}
	endpointsByID := map[string][]endpointRef{}
	endpointsByAlias := map[string]endpointRef{}
	endpointFrameByKey := map[string]string{}
	frameIDs := map[string]endpointRef{}
	var connections []*entity.Node

	var walk func(node, parent, grandparent *entity.Node, currentFrameID string) error
	walk = func(node, parent, grandparent *entity.Node, currentFrameID string) error {
		if node.Tag == "frame" {
			currentFrameID = strings.TrimSpace(node.Attrs["id"])
		}
		if nodeConnectableByIDV1EngineParseNode(node) {
			id := strings.TrimSpace(node.Attrs["id"])
			key := strings.TrimSpace(node.Attrs[internalConnectionKeyAttrV1EngineParseDocument])
			if id != "" {
				refFrameID := currentFrameID
				if node.Tag == "frame" {
					refFrameID = id
				}
				ref := endpointRef{key: key, frameID: refFrameID, position: node.Position}
				if key != "" && refFrameID != "" {
					endpointFrameByKey[key] = refFrameID
				}
				if node.Tag == "frame" || isConnectableFrameTagV1EngineParseNode(node.Tag) {
					if _, exists := frameIDs[id]; exists {
						loggerV1EngineSharedLogging.ERROR(IUPVCN007V1EngineParseDocument, "branch duplicate frame endpoint ID", map[string]any{"id": id})
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
						loggerV1EngineSharedLogging.ERROR(IUPVCN007V1EngineParseDocument, "branch duplicate endpoint alias", map[string]any{"alias": alias})
						return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate connection reference %q", alias)}
					}
					endpointsByAlias[alias] = ref
				}
			}
		}
		if node.Tag == "connections" && (parent == nil || parent.Tag != "frame") {
			parentTag := ""
			if parent != nil {
				parentTag = parent.Tag
			}
			loggerV1EngineSharedLogging.ERROR(IUPVCN005V1EngineParseDocument, "branch nested connections", map[string]any{"tag": parentTag})
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
				parentTag := ""
				if parent != nil {
					parentTag = parent.Tag
				}
				loggerV1EngineSharedLogging.ERROR(IUPVCN005V1EngineParseDocument, "branch nested connection", map[string]any{"tag": parentTag})
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<connection> must be a direct child of <frame> or <connections>")}
			}
			connections = append(connections, node)
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

	for _, conn := range connections {
		for _, endpoint := range []struct {
			attr    string
			keyAttr string
		}{
			{"src", internalConnectionSrcKeyAttrV1EngineParseDocument},
			{"dst", internalConnectionDstKeyAttrV1EngineParseDocument},
		} {
			token := strings.TrimSpace(conn.Attrs[endpoint.attr])
			if key := strings.TrimSpace(conn.Attrs[endpoint.keyAttr]); key != "" {
				setConnectionEndpointFrameV1EngineParseReference(conn, endpoint.attr, endpointFrameByKey[key])
				continue
			}
			if refs := endpointsByID[token]; len(refs) > 0 {
				if len(refs) > 1 {
					loggerV1EngineSharedLogging.ERROR(IUPVCN007V1EngineParseDocument, "branch ambiguous endpoint item", map[string]any{"attr": endpoint.attr, "id": token, "count": len(refs)})
					return &entity.ParseError{Position: conn.Position, Err: fmt.Errorf("<connection %s=%q> is ambiguous because endpoint id=%q appears %d times; use a unique name or ref", endpoint.attr, token, token, len(refs))}
				}
				conn.Attrs[endpoint.keyAttr] = refs[0].key
				setConnectionEndpointFrameV1EngineParseReference(conn, endpoint.attr, refs[0].frameID)
				continue
			}
			if ref, ok := endpointsByAlias[token]; ok {
				conn.Attrs[endpoint.keyAttr] = ref.key
				setConnectionEndpointFrameV1EngineParseReference(conn, endpoint.attr, ref.frameID)
				continue
			}
			loggerV1EngineSharedLogging.ERROR(IUPVCN006V1EngineParseDocument, "branch missing endpoint item", map[string]any{"attr": endpoint.attr, "token": token})
			return &entity.ParseError{Position: conn.Position, Err: fmt.Errorf("<connection %s=%q> does not match any connection endpoint id/name/ref", endpoint.attr, token)}
		}
		if strings.TrimSpace(conn.Attrs[internalConnectionSrcFrameAttrV1EngineParseDocument]) != "" &&
			strings.TrimSpace(conn.Attrs[internalConnectionDstFrameAttrV1EngineParseDocument]) != "" &&
			conn.Attrs[internalConnectionSrcFrameAttrV1EngineParseDocument] != conn.Attrs[internalConnectionDstFrameAttrV1EngineParseDocument] {
			conn.Attrs[internalConnectionCrossFrameAttrV1EngineParseDocument] = "true"
		}
	}
	return nil
}

func setConnectionEndpointFrameV1EngineParseReference(conn *entity.Node, endpoint, frameID string) {
	if conn == nil || frameID == "" {
		return
	}
	switch endpoint {
	case "src":
		conn.Attrs[internalConnectionSrcFrameAttrV1EngineParseDocument] = frameID
	case "dst":
		conn.Attrs[internalConnectionDstFrameAttrV1EngineParseDocument] = frameID
	}
}
