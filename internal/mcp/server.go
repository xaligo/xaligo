package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Server handles stateless MCP 2026-07-28 requests.
type Server struct {
	tools      ToolService
	serverInfo map[string]string
}

func NewServer(tools ToolService, version string) *Server {
	if strings.TrimSpace(version) == "" {
		version = "dev"
	}
	return &Server{
		tools: tools,
		serverInfo: map[string]string{
			"name": "xaligo", "version": version,
		},
	}
}

func (rcvr *Server) Handle(ctx context.Context, payload []byte) HandleResult {
	if err := validateServer(rcvr); err != nil {
		return rcvr.errorResult(nil, -32603, "Internal error", err.Error(), http.StatusInternalServerError)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if len(payload) > maxMessageBytes {
		return rcvr.errorResult(nil, -32600, "Request exceeds the MCP message size limit", nil, http.StatusRequestEntityTooLarge)
	}
	message, err := decodeMessage(payload)
	if err != nil {
		if json.Valid(payload) {
			return rcvr.errorResult(nil, -32600, "Invalid Request", nil, http.StatusBadRequest)
		}
		return rcvr.errorResult(nil, -32700, "Parse error", nil, http.StatusBadRequest)
	}
	if message.JSONRPC != "2.0" || strings.TrimSpace(message.Method) == "" {
		return rcvr.errorResult(message.ID, -32600, "Invalid Request", nil, http.StatusBadRequest)
	}
	if len(message.ID) == 0 {
		return HandleResult{HTTPStatus: http.StatusAccepted, Notification: true}
	}
	if string(message.ID) == "null" || (!json.Valid(message.ID)) {
		return rcvr.errorResult(nil, -32600, "Request id must be a string or integer", nil, http.StatusBadRequest)
	}
	if !validRequestID(message.ID) {
		return rcvr.errorResult(nil, -32600, "Request id must be a string or integer", nil, http.StatusBadRequest)
	}

	version, metadataErr := requestProtocolMetadata(message.Params)
	if metadataErr != nil {
		return rcvr.errorResult(message.ID, -32602, metadataErr.Error(), nil, http.StatusBadRequest)
	}
	if version != ProtocolVersion {
		return rcvr.errorResult(message.ID, -32022, "Unsupported protocol version", map[string]any{
			"supported": []string{ProtocolVersion}, "requested": version,
		}, http.StatusBadRequest)
	}

	result, protocolErr, status := rcvr.dispatch(ctx, message)
	if protocolErr != nil {
		return rcvr.errorResult(message.ID, protocolErr.Code, protocolErr.Message, protocolErr.Data, status)
	}
	return rcvr.successResult(message.ID, result)
}

func decodeMessage(payload []byte) (rpcMessage, error) {
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	var message rpcMessage
	if err := decoder.Decode(&message); err != nil {
		return rpcMessage{}, err
	}
	if decoder.More() {
		return rpcMessage{}, errors.New("multiple MCP values")
	}
	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		return rpcMessage{}, errors.New("multiple MCP values")
	}
	return message, nil
}

func validRequestID(raw json.RawMessage) bool {
	if len(raw) == 0 || string(raw) == "null" {
		return false
	}
	if raw[0] == '"' {
		var value string
		return json.Unmarshal(raw, &value) == nil
	}
	value := string(raw)
	if strings.ContainsAny(value, ".eE") {
		return false
	}
	_, err := strconv.ParseInt(value, 10, 64)
	return err == nil
}

func requestProtocolMetadata(raw json.RawMessage) (string, error) {
	version, metadata, err := requestProtocolVersion(raw)
	if err != nil {
		return "", err
	}
	var capabilities map[string]any
	if err := json.Unmarshal(metadata["io.modelcontextprotocol/clientCapabilities"], &capabilities); err != nil || capabilities == nil {
		return "", errors.New("params._meta clientCapabilities is required")
	}
	return version, nil
}

func requestProtocolVersion(raw json.RawMessage) (string, map[string]json.RawMessage, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return "", nil, errors.New("params._meta is required")
	}
	var params map[string]json.RawMessage
	if err := json.Unmarshal(raw, &params); err != nil {
		return "", nil, errors.New("params must be an object")
	}
	var metadata map[string]json.RawMessage
	if err := json.Unmarshal(params["_meta"], &metadata); err != nil {
		return "", nil, errors.New("params._meta must be an object")
	}
	var version string
	if err := json.Unmarshal(metadata["io.modelcontextprotocol/protocolVersion"], &version); err != nil || strings.TrimSpace(version) == "" {
		return "", nil, errors.New("params._meta protocolVersion is required")
	}
	return version, metadata, nil
}

func (rcvr *Server) dispatch(ctx context.Context, message rpcMessage) (any, *rpcError, int) {
	switch message.Method {
	case "server/discover":
		return map[string]any{
			"resultType":        "complete",
			"supportedVersions": []string{ProtocolVersion},
			"capabilities":      map[string]any{"tools": map[string]any{"listChanged": false}},
			"instructions":      "Validate, inspect, render, and search xaligo diagrams and documentation. Initial RAG indexing is restricted to Markdown below docs/.",
			"ttlMs":             300000,
			"cacheScope":        "public",
			"_meta":             rcvr.responseMeta(),
		}, nil, http.StatusOK
	case "tools/list":
		tools := rcvr.tools.Tools()
		sort.SliceStable(tools, func(left, right int) bool { return tools[left].Name < tools[right].Name })
		return map[string]any{
			"resultType": "complete", "tools": tools,
			"ttlMs": 300000, "cacheScope": "public", "_meta": rcvr.responseMeta(),
		}, nil, http.StatusOK
	case "tools/call":
		var params struct {
			Name      string          `json:"name"`
			Arguments json.RawMessage `json:"arguments"`
			Meta      json.RawMessage `json:"_meta"`
		}
		if err := json.Unmarshal(message.Params, &params); err != nil || strings.TrimSpace(params.Name) == "" {
			return nil, &rpcError{Code: -32602, Message: "tools/call requires a tool name and object arguments"}, http.StatusOK
		}
		if len(params.Arguments) == 0 {
			params.Arguments = json.RawMessage(`{}`)
		}
		if !json.Valid(params.Arguments) || params.Arguments[0] != '{' {
			return nil, &rpcError{Code: -32602, Message: "tools/call arguments must be an object"}, http.StatusOK
		}
		result, err := rcvr.tools.Call(ctx, params.Name, params.Arguments)
		if errors.Is(err, ErrUnknownTool) {
			return nil, &rpcError{Code: -32602, Message: "Unknown tool: " + params.Name}, http.StatusOK
		}
		if err != nil {
			result = ToolResult{
				ResultType: "complete", IsError: true,
				Content: []Content{{Type: "text", Text: err.Error()}},
			}
		}
		result.ResultType = "complete"
		if result.Content == nil {
			result.Content = []Content{}
		}
		result.Meta = rcvr.responseMeta()
		return result, nil, http.StatusOK
	default:
		return nil, &rpcError{Code: -32601, Message: "Method not found: " + message.Method}, http.StatusNotFound
	}
}

func (rcvr *Server) successResult(id json.RawMessage, result any) HandleResult {
	body, err := json.Marshal(map[string]any{"jsonrpc": "2.0", "id": id, "result": result})
	if err != nil {
		return rcvr.errorResult(id, -32603, "Internal error", err.Error(), http.StatusInternalServerError)
	}
	return HandleResult{Body: body, HTTPStatus: http.StatusOK}
}

func (rcvr *Server) errorResult(id json.RawMessage, code int, message string, data any, status int) HandleResult {
	if len(id) == 0 {
		id = json.RawMessage("null")
	}
	body, _ := json.Marshal(map[string]any{
		"jsonrpc": "2.0", "id": id,
		"error": rpcError{Code: code, Message: message, Data: data},
	})
	return HandleResult{Body: body, HTTPStatus: status}
}

func (rcvr *Server) responseMeta() map[string]any {
	return map[string]any{"io.modelcontextprotocol/serverInfo": rcvr.serverInfo}
}

func decodeRequestName(raw json.RawMessage) string {
	var params struct {
		Name string `json:"name"`
		URI  string `json:"uri"`
	}
	if json.Unmarshal(raw, &params) != nil {
		return ""
	}
	if params.Name != "" {
		return params.Name
	}
	return params.URI
}

func protocolErrorBody(id json.RawMessage, code int, message string, data any) []byte {
	server := NewServer(nil, "dev")
	return server.errorResult(id, code, message, data, http.StatusBadRequest).Body
}

func validateServer(server *Server) error {
	if server == nil || server.tools == nil {
		return fmt.Errorf("MCP tool service is required")
	}
	return nil
}
