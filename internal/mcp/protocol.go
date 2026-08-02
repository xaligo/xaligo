// Package mcp implements xaligo's Model Context Protocol adapter.
package mcp

import (
	"context"
	"encoding/json"
	"errors"
)

const (
	// ProtocolVersion is the stateless MCP revision implemented by xaligo.
	ProtocolVersion = "2026-07-28"
	maxMessageBytes = 8 * 1024 * 1024
)

var ErrUnknownTool = errors.New("unknown MCP tool")

type rpcMessage struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Tool describes one deterministic MCP tool entry.
type Tool struct {
	Name         string         `json:"name"`
	Title        string         `json:"title,omitempty"`
	Description  string         `json:"description"`
	InputSchema  map[string]any `json:"inputSchema"`
	OutputSchema map[string]any `json:"outputSchema,omitempty"`
	Annotations  map[string]any `json:"annotations,omitempty"`
}

// Content is one unstructured MCP tool result block.
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ToolResult is the complete result of a tool invocation.
type ToolResult struct {
	ResultType        string    `json:"resultType"`
	Content           []Content `json:"content"`
	StructuredContent any       `json:"structuredContent,omitempty"`
	IsError           bool      `json:"isError,omitempty"`
	Meta              any       `json:"_meta,omitempty"`
}

// ToolService owns the MCP-visible tool catalog and invocation adapter.
type ToolService interface {
	Tools() []Tool
	Call(context.Context, string, json.RawMessage) (ToolResult, error)
}

// HandleResult is one transport-neutral JSON-RPC handling outcome.
type HandleResult struct {
	Body         []byte
	HTTPStatus   int
	Notification bool
}

func requestIDKey(id json.RawMessage) string {
	return string(id)
}
