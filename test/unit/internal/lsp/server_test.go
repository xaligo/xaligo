package lsp_test

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/lsp"
)

type fakeLSPProject struct {
	indexed int
}

func (rcvr *fakeLSPProject) Analyze(_ context.Context, uri string, source []byte) (entity.ProjectAnalysis, error) {
	analysis := entity.ProjectAnalysis{
		URI: uri, Kind: entity.ProjectDocumentXAL, Source: append([]byte(nil), source...),
		Symbols: []entity.ProjectSymbol{
			{Ordinal: 0, ParentOrdinal: -1, ID: "main", Name: "Main", Concept: entity.ProjectConceptFrame, SourceTag: "frame", Position: entity.Position{Offset: 0, Line: 1, Column: 1}},
			{Ordinal: 1, ParentOrdinal: 0, ID: "api", Name: "API", Concept: entity.ProjectConceptItem, SourceTag: "item", Position: entity.Position{Offset: 8, Line: 1, Column: 9}},
			{Ordinal: 2, ParentOrdinal: 1, ID: "port", Name: "Port", Concept: entity.ProjectConceptPort, SourceTag: "port", Position: entity.Position{Offset: 14, Line: 1, Column: 15}},
		},
	}
	if bytes.Contains(source, []byte("<connection")) {
		analysis.Symbols = append(analysis.Symbols, entity.ProjectSymbol{Ordinal: 3, ParentOrdinal: 0, ID: "flow", Name: "flow", Concept: entity.ProjectConceptLine, SourceTag: "connection", Source: "api", Target: "api", Position: entity.Position{Offset: 36, Line: 1, Column: 37}})
	}
	if bytes.Contains(source, []byte("bad")) {
		analysis.Diagnostics = []entity.Diagnostic{{Severity: "error", Message: "invalid source", Offset: 1, Line: 1, Column: 2}}
	}
	return analysis, nil
}

func TestServerCompletionDefinitionAndReferences(t *testing.T) {
	source := `<frame id="main"><item id="api"/><connection src="api" dst="api"/></frame>`
	position := strings.Index(source, `src="api"`) + len(`src="`)
	var input bytes.Buffer
	writeLSPInput(&input, request(1, "initialize", map[string]any{}))
	writeLSPInput(&input, map[string]any{"jsonrpc": "2.0", "method": "textDocument/didOpen", "params": map[string]any{"textDocument": map[string]any{"uri": "file:///tmp/diagram.xal", "version": 1, "text": source}}})
	params := map[string]any{"textDocument": map[string]any{"uri": "file:///tmp/diagram.xal"}, "position": map[string]any{"line": 0, "character": position}}
	writeLSPInput(&input, request(2, "textDocument/completion", params))
	writeLSPInput(&input, request(3, "textDocument/definition", params))
	writeLSPInput(&input, request(4, "textDocument/references", params))
	writeLSPInput(&input, request(5, "shutdown", map[string]any{}))
	writeLSPInput(&input, map[string]any{"jsonrpc": "2.0", "method": "exit"})
	var output bytes.Buffer
	if err := lsp.NewServer(&fakeLSPProject{}).Serve(context.Background(), &input, &output); err != nil {
		t.Fatal(err)
	}
	messages := readLSPOutput(t, output.Bytes())
	if got := len(messages[2]["result"].([]any)); got != 7 {
		t.Fatalf("completion count = %d", got)
	}
	if messages[3]["result"] == nil {
		t.Fatalf("definition = %#v", messages[3])
	}
	if got := len(messages[4]["result"].([]any)); got != 2 {
		t.Fatalf("references = %#v", messages[4])
	}
}

func (rcvr *fakeLSPProject) IndexDocument(ctx context.Context, uri string, source []byte) (entity.ProjectAnalysis, bool, error) {
	rcvr.indexed++
	analysis, err := rcvr.Analyze(ctx, uri, source)
	return analysis, true, err
}

func (*fakeLSPProject) Search(context.Context, string, int) ([]entity.ProjectSearchResult, error) {
	return []entity.ProjectSearchResult{{
		URI: "file:///docs/guide.md", ID: "heading-1", Name: "API Guide",
		Concept: entity.ProjectConceptText, SourceTag: "heading", Line: 1, Column: 1,
	}}, nil
}

var _ lsp.ProjectService = (*fakeLSPProject)(nil)

func TestServerLifecycleDiagnosticsSymbolsTokensHoverAndSave(t *testing.T) {
	project := &fakeLSPProject{}
	var input bytes.Buffer
	writeLSPInput(&input, map[string]any{"jsonrpc": "2.0", "id": 1, "method": "initialize", "params": map[string]any{}})
	writeLSPInput(&input, map[string]any{
		"jsonrpc": "2.0", "method": "textDocument/didOpen",
		"params": map[string]any{"textDocument": map[string]any{
			"uri": "file:///tmp/diagram.xal", "version": 1, "text": "<frame><item><port/></item></frame>",
		}},
	})
	writeLSPInput(&input, request(2, "textDocument/documentSymbol", documentParams()))
	writeLSPInput(&input, request(3, "textDocument/semanticTokens/full", documentParams()))
	writeLSPInput(&input, request(4, "textDocument/diagnostic", documentParams()))
	writeLSPInput(&input, request(5, "textDocument/hover", map[string]any{
		"textDocument": map[string]any{"uri": "file:///tmp/diagram.xal"},
		"position":     map[string]any{"line": 0, "character": 10},
	}))
	writeLSPInput(&input, request(6, "workspace/symbol", map[string]any{"query": "api"}))
	writeLSPInput(&input, map[string]any{"jsonrpc": "2.0", "method": "textDocument/didSave", "params": documentParams()})
	writeLSPInput(&input, request(7, "shutdown", map[string]any{}))
	writeLSPInput(&input, map[string]any{"jsonrpc": "2.0", "method": "exit"})

	var output bytes.Buffer
	if err := lsp.NewServer(project).Serve(context.Background(), &input, &output); err != nil {
		t.Fatal(err)
	}
	messages := readLSPOutput(t, output.Bytes())
	if len(messages) != 8 {
		t.Fatalf("message count = %d, messages=%#v", len(messages), messages)
	}
	if project.indexed != 1 {
		t.Fatalf("saved .xal index calls = %d", project.indexed)
	}
	initialize := messages[0]["result"].(map[string]any)
	capabilities := initialize["capabilities"].(map[string]any)
	if capabilities["positionEncoding"] != "utf-16" || capabilities["documentSymbolProvider"] != true {
		t.Fatalf("initialize capabilities = %#v", capabilities)
	}
	if messages[1]["method"] != "textDocument/publishDiagnostics" {
		t.Fatalf("didOpen notification = %#v", messages[1])
	}
	documentSymbols := messages[2]["result"].([]any)
	root := documentSymbols[0].(map[string]any)
	child := root["children"].([]any)[0].(map[string]any)
	if child["name"] != "API" || len(child["children"].([]any)) != 1 {
		t.Fatalf("nested document symbols = %#v", documentSymbols)
	}
	tokens := messages[3]["result"].(map[string]any)["data"].([]any)
	if len(tokens) != 15 {
		t.Fatalf("semantic token data = %#v", tokens)
	}
	if messages[4]["result"].(map[string]any)["kind"] != "full" {
		t.Fatalf("pull diagnostics = %#v", messages[4])
	}
	if messages[5]["result"] == nil {
		t.Fatalf("hover result = %#v", messages[5])
	}
	workspace := messages[6]["result"].([]any)
	if len(workspace) != 2 {
		t.Fatalf("workspace symbols = %#v", workspace)
	}
	if result, exists := messages[7]["result"]; !exists || result != nil {
		t.Fatalf("shutdown result = %#v", messages[7])
	}
}

func TestServerPublishesDiagnosticsAfterFullDocumentChange(t *testing.T) {
	var input bytes.Buffer
	writeLSPInput(&input, request(1, "initialize", map[string]any{}))
	writeLSPInput(&input, map[string]any{
		"jsonrpc": "2.0", "method": "textDocument/didOpen",
		"params": map[string]any{"textDocument": map[string]any{
			"uri": "file:///tmp/diagram.xal", "version": 1, "text": "good",
		}},
	})
	writeLSPInput(&input, map[string]any{
		"jsonrpc": "2.0", "method": "textDocument/didChange",
		"params": map[string]any{
			"textDocument":   map[string]any{"uri": "file:///tmp/diagram.xal", "version": 2},
			"contentChanges": []any{map[string]any{"text": "bad"}},
		},
	})
	writeLSPInput(&input, request(2, "shutdown", map[string]any{}))
	writeLSPInput(&input, map[string]any{"jsonrpc": "2.0", "method": "exit"})
	var output bytes.Buffer
	if err := lsp.NewServer(&fakeLSPProject{}).Serve(context.Background(), &input, &output); err != nil {
		t.Fatal(err)
	}
	messages := readLSPOutput(t, output.Bytes())
	diagnostics := messages[2]["params"].(map[string]any)["diagnostics"].([]any)
	if len(diagnostics) != 1 || diagnostics[0].(map[string]any)["message"] != "invalid source" {
		t.Fatalf("changed diagnostics = %#v", diagnostics)
	}
}

func TestServerRejectsMissingContentLength(t *testing.T) {
	err := lsp.NewServer(&fakeLSPProject{}).Serve(context.Background(), strings.NewReader("\r\n{}"), io.Discard)
	if err == nil || !strings.Contains(err.Error(), "Content-Length") {
		t.Fatalf("missing header error = %v", err)
	}
}

func request(id int, method string, params any) map[string]any {
	return map[string]any{"jsonrpc": "2.0", "id": id, "method": method, "params": params}
}

func documentParams() map[string]any {
	return map[string]any{"textDocument": map[string]any{"uri": "file:///tmp/diagram.xal"}}
}

func writeLSPInput(output *bytes.Buffer, message any) {
	payload, _ := json.Marshal(message)
	fmt.Fprintf(output, "Content-Length: %d\r\n\r\n", len(payload))
	output.Write(payload)
}

func readLSPOutput(t *testing.T, data []byte) []map[string]any {
	t.Helper()
	reader := bufio.NewReader(bytes.NewReader(data))
	var messages []map[string]any
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF && line == "" {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		name, rawLength, found := strings.Cut(strings.TrimSpace(line), ":")
		if !found || name != "Content-Length" {
			t.Fatalf("invalid output header %q", line)
		}
		length, err := strconv.Atoi(strings.TrimSpace(rawLength))
		if err != nil {
			t.Fatal(err)
		}
		blank, err := reader.ReadString('\n')
		if err != nil || strings.TrimSpace(blank) != "" {
			t.Fatalf("invalid output separator %q: %v", blank, err)
		}
		payload := make([]byte, length)
		if _, err := io.ReadFull(reader, payload); err != nil {
			t.Fatal(err)
		}
		var message map[string]any
		if err := json.Unmarshal(payload, &message); err != nil {
			t.Fatal(err)
		}
		messages = append(messages, message)
	}
	return messages
}
