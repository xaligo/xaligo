package mcp_test

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/mcp"
)

type fakeMCPTools struct {
	called chan struct{}
}

func (*fakeMCPTools) Tools() []mcp.Tool {
	return []mcp.Tool{{
		Name: "echo", Description: "Echo text",
		InputSchema: map[string]any{"type": "object", "properties": map[string]any{"text": map[string]any{"type": "string"}}},
	}}
}

func (rcvr *fakeMCPTools) Call(ctx context.Context, name string, arguments json.RawMessage) (mcp.ToolResult, error) {
	if name == "block" {
		if rcvr.called != nil {
			close(rcvr.called)
		}
		<-ctx.Done()
		return mcp.ToolResult{}, ctx.Err()
	}
	if name != "echo" {
		return mcp.ToolResult{}, mcp.ErrUnknownTool
	}
	var args struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal(arguments, &args); err != nil {
		return mcp.ToolResult{}, err
	}
	return mcp.ToolResult{
		ResultType:        "complete",
		Content:           []mcp.Content{{Type: "text", Text: args.Text}},
		StructuredContent: map[string]any{"text": args.Text},
	}, nil
}

func TestServerStdioDiscoveryToolsAndCall(t *testing.T) {
	server := mcp.NewServer(&fakeMCPTools{}, "test")
	var input bytes.Buffer
	writeMCPLine(&input, request(1, "server/discover", map[string]any{"_meta": metadata(mcp.ProtocolVersion)}))
	writeMCPLine(&input, request(2, "tools/list", map[string]any{"_meta": metadata(mcp.ProtocolVersion)}))
	writeMCPLine(&input, request(3, "tools/call", map[string]any{
		"_meta": metadata(mcp.ProtocolVersion), "name": "echo", "arguments": map[string]any{"text": "hello"},
	}))
	var output bytes.Buffer
	if err := server.ServeStdio(context.Background(), &input, &output); err != nil {
		t.Fatal(err)
	}
	messages := readMCPOutput(t, output.Bytes())
	if len(messages) != 3 {
		t.Fatalf("response count = %d, messages=%#v", len(messages), messages)
	}
	byID := indexResponses(messages)
	discovery := byID["1"]["result"].(map[string]any)
	if discovery["resultType"] != "complete" || discovery["supportedVersions"].([]any)[0] != mcp.ProtocolVersion {
		t.Fatalf("discovery result = %#v", discovery)
	}
	tools := byID["2"]["result"].(map[string]any)["tools"].([]any)
	if len(tools) != 1 || tools[0].(map[string]any)["name"] != "echo" {
		t.Fatalf("tools result = %#v", tools)
	}
	call := byID["3"]["result"].(map[string]any)
	if call["resultType"] != "complete" || call["structuredContent"].(map[string]any)["text"] != "hello" {
		t.Fatalf("tool call result = %#v", call)
	}
}

func TestServerRejectsMissingPerRequestCapabilities(t *testing.T) {
	server := mcp.NewServer(&fakeMCPTools{}, "test")
	params := map[string]any{"_meta": map[string]any{"io.modelcontextprotocol/protocolVersion": mcp.ProtocolVersion}}
	payload, _ := json.Marshal(request(1, "tools/list", params))
	result := server.Handle(context.Background(), payload)
	message := decodeMCPMessage(t, result.Body)
	protocolError := message["error"].(map[string]any)
	if result.HTTPStatus != http.StatusBadRequest || protocolError["code"] != float64(-32602) {
		t.Fatalf("missing capabilities result = status %d %#v", result.HTTPStatus, message)
	}
}

func TestServerStdioCancellationSuppressesResponse(t *testing.T) {
	tools := &fakeMCPTools{called: make(chan struct{})}
	server := mcp.NewServer(tools, "test")
	reader, writer := io.Pipe()
	var output bytes.Buffer
	completed := make(chan error, 1)
	go func() { completed <- server.ServeStdio(context.Background(), reader, &output) }()
	writeMCPLineWriter(t, writer, request(7, "tools/call", map[string]any{
		"_meta": metadata(mcp.ProtocolVersion), "name": "block", "arguments": map[string]any{},
	}))
	<-tools.called
	writeMCPLineWriter(t, writer, map[string]any{
		"jsonrpc": "2.0", "method": "notifications/cancelled", "params": map[string]any{"requestId": 7},
	})
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := <-completed; err != nil {
		t.Fatal(err)
	}
	if output.Len() != 0 {
		t.Fatalf("cancelled request wrote %q", output.String())
	}
}

func TestHTTPHandlerValidatesTransportHeadersAndOrigin(t *testing.T) {
	server := mcp.NewServer(&fakeMCPTools{}, "test")
	handler := mcp.NewHTTPHandler(server)
	payload, _ := json.Marshal(request(1, "tools/list", map[string]any{"_meta": metadata(mcp.ProtocolVersion)}))

	httpRequest := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8081/mcp", bytes.NewReader(payload))
	setMCPHTTPHeaders(httpRequest, "tools/list", "")
	httpRequest.Header.Set("Origin", "http://localhost:3000")
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httpRequest)
	if response.Code != http.StatusOK {
		t.Fatalf("valid HTTP status = %d body=%s", response.Code, response.Body.String())
	}

	badOrigin := httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8081/mcp", bytes.NewReader(payload))
	setMCPHTTPHeaders(badOrigin, "tools/list", "")
	badOrigin.Header.Set("Origin", "https://example.com")
	badResponse := httptest.NewRecorder()
	handler.ServeHTTP(badResponse, badOrigin)
	if badResponse.Code != http.StatusForbidden {
		t.Fatalf("bad origin status = %d body=%s", badResponse.Code, badResponse.Body.String())
	}
}

func TestHTTPHandlerRejectsMismatchedNameAndUnsupportedVersion(t *testing.T) {
	server := mcp.NewServer(&fakeMCPTools{}, "test")
	handler := mcp.NewHTTPHandler(server)
	callPayload, _ := json.Marshal(request(1, "tools/call", map[string]any{
		"_meta": metadata(mcp.ProtocolVersion), "name": "echo", "arguments": map[string]any{"text": "hello"},
	}))
	httpRequest := httptest.NewRequest(http.MethodPost, "http://localhost/mcp", bytes.NewReader(callPayload))
	setMCPHTTPHeaders(httpRequest, "tools/call", "other")
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httpRequest)
	message := decodeMCPMessage(t, response.Body.Bytes())
	if response.Code != http.StatusBadRequest || message["error"].(map[string]any)["code"] != float64(-32020) {
		t.Fatalf("header mismatch = status %d %#v", response.Code, message)
	}

	unsupported := "2099-01-01"
	unsupportedPayload, _ := json.Marshal(request(2, "tools/list", map[string]any{"_meta": metadata(unsupported)}))
	httpRequest = httptest.NewRequest(http.MethodPost, "http://localhost/mcp", bytes.NewReader(unsupportedPayload))
	setMCPHTTPHeaders(httpRequest, "tools/list", "")
	httpRequest.Header.Set("MCP-Protocol-Version", unsupported)
	response = httptest.NewRecorder()
	handler.ServeHTTP(response, httpRequest)
	message = decodeMCPMessage(t, response.Body.Bytes())
	if response.Code != http.StatusBadRequest || message["error"].(map[string]any)["code"] != float64(-32022) {
		t.Fatalf("unsupported version = status %d %#v", response.Code, message)
	}
}

func metadata(version string) map[string]any {
	return map[string]any{
		"io.modelcontextprotocol/protocolVersion":    version,
		"io.modelcontextprotocol/clientCapabilities": map[string]any{},
		"io.modelcontextprotocol/clientInfo":         map[string]any{"name": "test", "version": "1"},
	}
}

func request(id int, method string, params any) map[string]any {
	return map[string]any{"jsonrpc": "2.0", "id": id, "method": method, "params": params}
}

func writeMCPLine(output *bytes.Buffer, value any) {
	payload, _ := json.Marshal(value)
	output.Write(payload)
	output.WriteByte('\n')
}

func writeMCPLineWriter(t *testing.T, output io.Writer, value any) {
	t.Helper()
	payload, _ := json.Marshal(value)
	if _, err := output.Write(append(payload, '\n')); err != nil {
		t.Fatal(err)
	}
}

func readMCPOutput(t *testing.T, data []byte) []map[string]any {
	t.Helper()
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var messages []map[string]any
	for scanner.Scan() {
		messages = append(messages, decodeMCPMessage(t, scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	return messages
}

func decodeMCPMessage(t *testing.T, data []byte) map[string]any {
	t.Helper()
	var message map[string]any
	if err := json.Unmarshal(data, &message); err != nil {
		t.Fatalf("decode %q: %v", data, err)
	}
	return message
}

func indexResponses(messages []map[string]any) map[string]map[string]any {
	indexed := make(map[string]map[string]any, len(messages))
	for _, message := range messages {
		indexed[strings.TrimSuffix(strings.TrimSuffix(jsonNumber(message["id"]), ".0"), ".0")] = message
	}
	return indexed
}

func jsonNumber(value any) string {
	payload, _ := json.Marshal(value)
	return string(payload)
}

func setMCPHTTPHeaders(request *http.Request, method, name string) {
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json, text/event-stream")
	request.Header.Set("MCP-Protocol-Version", mcp.ProtocolVersion)
	request.Header.Set("Mcp-Method", method)
	if name != "" {
		request.Header.Set("Mcp-Name", name)
	}
}
