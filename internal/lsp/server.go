package lsp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode/utf16"

	"github.com/xaligo/xaligo/internal/entity"
)

// ProjectService is the shared application boundary required by the LSP
// adapter. usecase.ProjectUsecase satisfies it without another parser.
type ProjectService interface {
	Analyze(context.Context, string, []byte) (entity.ProjectAnalysis, error)
	IndexDocument(context.Context, string, []byte) (entity.ProjectAnalysis, bool, error)
	Search(context.Context, string, int) ([]entity.ProjectSearchResult, error)
}

type documentState struct {
	uri      string
	version  int
	source   []byte
	analysis entity.ProjectAnalysis
}

// Server serves one LSP 3.18 stdio session.
type Server struct {
	project     ProjectService
	transport   *transport
	documents   map[string]*documentState
	initialized bool
	shutdown    bool
}

func NewServer(project ProjectService) *Server {
	return &Server{project: project, documents: map[string]*documentState{}}
}

func (rcvr *Server) Serve(ctx context.Context, input io.Reader, output io.Writer) error {
	if rcvr.project == nil {
		return errors.New("LSP project service is required")
	}
	if input == nil || output == nil {
		return errors.New("LSP input and output are required")
	}
	rcvr.transport = newTransport(input, output)
	for {
		if err := contextError(ctx); err != nil {
			return err
		}
		payload, err := rcvr.transport.read()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		var message rpcMessage
		if err := json.Unmarshal(payload, &message); err != nil {
			if err := rcvr.transport.response(nil, nil, &rpcError{Code: -32700, Message: "Parse error"}); err != nil {
				return err
			}
			continue
		}
		stop, err := rcvr.handle(ctx, message)
		if err != nil {
			return err
		}
		if stop {
			if !rcvr.shutdown {
				return errors.New("LSP exit received before shutdown")
			}
			return nil
		}
	}
}

func (rcvr *Server) handle(ctx context.Context, message rpcMessage) (bool, error) {
	if message.JSONRPC != "2.0" || strings.TrimSpace(message.Method) == "" {
		if messageHasID(message) {
			return false, rcvr.transport.response(message.ID, nil, &rpcError{Code: -32600, Message: "Invalid Request"})
		}
		return false, nil
	}
	if message.Method == "exit" {
		return true, nil
	}
	if message.Method != "initialize" && !rcvr.initialized {
		if messageHasID(message) {
			return false, rcvr.transport.response(message.ID, nil, &rpcError{Code: -32002, Message: "Server not initialized"})
		}
		return false, nil
	}
	if rcvr.shutdown {
		if messageHasID(message) {
			return false, rcvr.transport.response(message.ID, nil, &rpcError{Code: -32600, Message: "Server is shutting down"})
		}
		return false, nil
	}

	result, err := rcvr.dispatch(ctx, message)
	if !messageHasID(message) {
		if err != nil {
			return false, rcvr.logError(err)
		}
		return false, nil
	}
	if err == nil {
		return false, rcvr.transport.response(message.ID, result, nil)
	}
	var protocol *protocolError
	if errors.As(err, &protocol) {
		return false, rcvr.transport.response(message.ID, nil, &rpcError{Code: protocol.code, Message: protocol.message, Data: protocol.data})
	}
	return false, rcvr.transport.response(message.ID, nil, &rpcError{Code: -32603, Message: "Internal error", Data: err.Error()})
}

func (rcvr *Server) dispatch(ctx context.Context, message rpcMessage) (any, error) {
	switch message.Method {
	case "initialize":
		if rcvr.initialized {
			return nil, &protocolError{code: -32600, message: "initialize may only be sent once"}
		}
		rcvr.initialized = true
		return initializeResult(), nil
	case "initialized", "$/cancelRequest", "workspace/didChangeConfiguration", "workspace/didChangeWatchedFiles":
		return nil, nil
	case "shutdown":
		rcvr.shutdown = true
		return nil, nil
	case "textDocument/didOpen":
		return nil, rcvr.didOpen(ctx, message.Params)
	case "textDocument/didChange":
		return nil, rcvr.didChange(ctx, message.Params)
	case "textDocument/didSave":
		return nil, rcvr.didSave(ctx, message.Params)
	case "textDocument/didClose":
		return nil, rcvr.didClose(message.Params)
	case "textDocument/documentSymbol":
		return rcvr.documentSymbols(ctx, message.Params)
	case "workspace/symbol":
		return rcvr.workspaceSymbols(ctx, message.Params)
	case "textDocument/diagnostic":
		return rcvr.documentDiagnostics(ctx, message.Params)
	case "textDocument/semanticTokens/full":
		return rcvr.semanticTokens(ctx, message.Params)
	case "textDocument/hover":
		return rcvr.hover(ctx, message.Params)
	case "textDocument/completion":
		return rcvr.completion(), nil
	case "textDocument/definition":
		return rcvr.definition(ctx, message.Params)
	case "textDocument/references":
		return rcvr.references(ctx, message.Params)
	default:
		return nil, &protocolError{code: -32601, message: "Method not found: " + message.Method}
	}
}

func initializeResult() map[string]any {
	return map[string]any{
		"serverInfo": map[string]any{"name": "xaligo", "version": "dev"},
		"capabilities": map[string]any{
			"positionEncoding": "utf-16",
			"textDocumentSync": map[string]any{
				"openClose": true, "change": 1,
				"save": map[string]any{"includeText": true},
			},
			"documentSymbolProvider":  true,
			"workspaceSymbolProvider": true,
			"hoverProvider":           true,
			"definitionProvider":      true,
			"referencesProvider":      true,
			"completionProvider": map[string]any{
				"triggerCharacters": []string{"<", " ", "\""},
			},
			"diagnosticProvider": map[string]any{
				"identifier": "xaligo", "interFileDependencies": false, "workspaceDiagnostics": false,
			},
			"semanticTokensProvider": map[string]any{
				"legend": map[string]any{
					"tokenTypes":     []string{"namespace", "class", "property", "event", "string", "operator"},
					"tokenModifiers": []string{},
				},
				"full": true,
			},
		},
	}
}

func (rcvr *Server) completion() []map[string]any {
	values := []struct{ label, detail, insert string }{
		{"frame", "Physical canvas", `<frame id="${1:main}" width="${2:1280}" height="${3:720}">$0</frame>`},
		{"row", "Horizontal layout", `<row gap="${1:16}">$0</row>`},
		{"col", "Vertical layout", `<col span="${1:12}">$0</col>`},
		{"item", "Icon or atomic item", `<item id="${1:id}" name="${2:label}" />`},
		{"rectangle", "Generic rectangle", `<rectangle id="${1:id}" title="${2:title}" />`},
		{"port", "Connection port", `<port id="${1:id}" side="${2|left,right,top,bottom|}" />`},
		{"connection", "Routed connection", `<connection src="${1:source}" dst="${2:target}" kind="${3|route,traffic|}" />`},
	}
	result := make([]map[string]any, 0, len(values))
	for _, value := range values {
		result = append(result, map[string]any{"label": value.label, "kind": 10, "detail": value.detail, "insertText": value.insert, "insertTextFormat": 2})
	}
	return result
}

func (rcvr *Server) definition(ctx context.Context, raw json.RawMessage) (any, error) {
	state, word, err := rcvr.documentWord(ctx, raw)
	if err != nil || word == "" {
		return nil, err
	}
	for _, symbol := range state.analysis.Symbols {
		if symbol.ID == word {
			return location{URI: state.uri, Range: rangeForSymbol(state.source, symbol)}, nil
		}
	}
	return nil, nil
}

func (rcvr *Server) references(ctx context.Context, raw json.RawMessage) (any, error) {
	state, word, err := rcvr.documentWord(ctx, raw)
	if err != nil || word == "" {
		return nil, err
	}
	result := make([]location, 0)
	for _, symbol := range state.analysis.Symbols {
		if symbol.ID == word || symbol.Source == word || symbol.Target == word {
			result = append(result, location{URI: state.uri, Range: rangeForSymbol(state.source, symbol)})
		}
	}
	return result, nil
}

func (rcvr *Server) documentWord(ctx context.Context, raw json.RawMessage) (*documentState, string, error) {
	var params struct {
		TextDocument struct {
			URI string `json:"uri"`
		} `json:"textDocument"`
		Position position `json:"position"`
	}
	if err := decodeParams(raw, &params); err != nil {
		return nil, "", err
	}
	state, err := rcvr.document(ctx, params.TextDocument.URI)
	if err != nil {
		return nil, "", err
	}
	return state, lspWordAt(state.source, params.Position), nil
}

func lspWordAt(source []byte, at position) string {
	lines := strings.Split(string(source), "\n")
	if at.Line < 0 || at.Line >= len(lines) {
		return ""
	}
	line := []rune(lines[at.Line])
	index := min(max(at.Character, 0), len(line))
	isWord := func(value rune) bool {
		return value == '-' || value == '_' || value == '.' || value == ':' || value >= '0' && value <= '9' || value >= 'A' && value <= 'Z' || value >= 'a' && value <= 'z'
	}
	start, end := index, index
	for start > 0 && isWord(line[start-1]) {
		start--
	}
	for end < len(line) && isWord(line[end]) {
		end++
	}
	return string(line[start:end])
}

func (rcvr *Server) didOpen(ctx context.Context, raw json.RawMessage) error {
	var params struct {
		TextDocument struct {
			URI     string `json:"uri"`
			Version int    `json:"version"`
			Text    string `json:"text"`
		} `json:"textDocument"`
	}
	if err := decodeParams(raw, &params); err != nil {
		return err
	}
	if strings.TrimSpace(params.TextDocument.URI) == "" {
		return invalidParams("textDocument.uri is required")
	}
	state := &documentState{
		uri: params.TextDocument.URI, version: params.TextDocument.Version,
		source: []byte(params.TextDocument.Text),
	}
	rcvr.documents[state.uri] = state
	return rcvr.analyzeAndPublish(ctx, state)
}

func (rcvr *Server) didChange(ctx context.Context, raw json.RawMessage) error {
	var params struct {
		TextDocument struct {
			URI     string `json:"uri"`
			Version int    `json:"version"`
		} `json:"textDocument"`
		ContentChanges []struct {
			Text string `json:"text"`
		} `json:"contentChanges"`
	}
	if err := decodeParams(raw, &params); err != nil {
		return err
	}
	state := rcvr.documents[params.TextDocument.URI]
	if state == nil {
		return invalidParams("textDocument/didChange requires an open document")
	}
	if len(params.ContentChanges) == 0 {
		return invalidParams("contentChanges must contain a full document change")
	}
	state.version = params.TextDocument.Version
	state.source = []byte(params.ContentChanges[len(params.ContentChanges)-1].Text)
	return rcvr.analyzeAndPublish(ctx, state)
}

func (rcvr *Server) didSave(ctx context.Context, raw json.RawMessage) error {
	var params struct {
		TextDocument struct {
			URI string `json:"uri"`
		} `json:"textDocument"`
		Text *string `json:"text,omitempty"`
	}
	if err := decodeParams(raw, &params); err != nil {
		return err
	}
	state := rcvr.documents[params.TextDocument.URI]
	if state == nil {
		return invalidParams("textDocument/didSave requires an open document")
	}
	if params.Text != nil {
		state.source = []byte(*params.Text)
		if err := rcvr.analyzeAndPublish(ctx, state); err != nil {
			return err
		}
	}
	if strings.EqualFold(filepath.Ext(lspURIPath(state.uri)), ".xal") {
		analysis, _, err := rcvr.project.IndexDocument(ctx, state.uri, state.source)
		if err != nil {
			return fmt.Errorf("index saved .xal document: %w", err)
		}
		state.analysis = analysis
	}
	return nil
}

func (rcvr *Server) didClose(raw json.RawMessage) error {
	var params textDocumentParams
	if err := decodeParams(raw, &params); err != nil {
		return err
	}
	delete(rcvr.documents, params.TextDocument.URI)
	return rcvr.transport.notification("textDocument/publishDiagnostics", map[string]any{
		"uri": params.TextDocument.URI, "diagnostics": []any{},
	})
}

func (rcvr *Server) analyzeAndPublish(ctx context.Context, state *documentState) error {
	analysis, err := rcvr.project.Analyze(ctx, state.uri, state.source)
	if err != nil {
		return err
	}
	state.analysis = analysis
	return rcvr.transport.notification("textDocument/publishDiagnostics", map[string]any{
		"uri": state.uri, "version": state.version,
		"diagnostics": lspDiagnostics(state.source, analysis.Diagnostics),
	})
}

func (rcvr *Server) documentSymbols(ctx context.Context, raw json.RawMessage) ([]documentSymbol, error) {
	state, err := rcvr.documentForParams(ctx, raw)
	if err != nil {
		return nil, err
	}
	return lspDocumentSymbols(state.source, state.analysis.Symbols), nil
}

func (rcvr *Server) workspaceSymbols(ctx context.Context, raw json.RawMessage) ([]symbolInformation, error) {
	var params struct {
		Query string `json:"query"`
	}
	if err := decodeParams(raw, &params); err != nil {
		return nil, err
	}
	query := strings.TrimSpace(params.Query)
	seen := map[string]bool{}
	var symbols []symbolInformation
	for _, state := range rcvr.documents {
		for _, symbol := range state.analysis.Symbols {
			if query != "" && !projectSymbolMatches(symbol, query) {
				continue
			}
			key := state.uri + "\x00" + symbol.ID
			if seen[key] {
				continue
			}
			seen[key] = true
			symbols = append(symbols, lspSymbolInformation(state.uri, state.source, symbol, state.analysis.Symbols))
		}
	}
	if query != "" {
		results, err := rcvr.project.Search(ctx, lspFTSQuery(query), 100)
		if err != nil {
			return nil, err
		}
		for _, result := range results {
			key := result.URI + "\x00" + result.ID
			if seen[key] {
				continue
			}
			seen[key] = true
			symbols = append(symbols, symbolInformation{
				Name: result.Name, Kind: lspSymbolKind(result.Concept),
				Location:      location{URI: result.URI, Range: rangeFromLineColumn(result.Line, result.Column, result.Name)},
				ContainerName: result.SourceTag,
			})
		}
	}
	sort.SliceStable(symbols, func(left, right int) bool {
		if symbols[left].Location.URI != symbols[right].Location.URI {
			return symbols[left].Location.URI < symbols[right].Location.URI
		}
		if symbols[left].Location.Range.Start.Line != symbols[right].Location.Range.Start.Line {
			return symbols[left].Location.Range.Start.Line < symbols[right].Location.Range.Start.Line
		}
		return symbols[left].Name < symbols[right].Name
	})
	return symbols, nil
}

func (rcvr *Server) documentDiagnostics(ctx context.Context, raw json.RawMessage) (map[string]any, error) {
	state, err := rcvr.documentForParams(ctx, raw)
	if err != nil {
		return nil, err
	}
	return map[string]any{"kind": "full", "items": lspDiagnostics(state.source, state.analysis.Diagnostics)}, nil
}

func (rcvr *Server) semanticTokens(ctx context.Context, raw json.RawMessage) (map[string]any, error) {
	state, err := rcvr.documentForParams(ctx, raw)
	if err != nil {
		return nil, err
	}
	type token struct {
		line, character, length, kind int
	}
	var tokens []token
	for _, symbol := range state.analysis.Symbols {
		if symbol.SourceTag == "" || symbol.SourceTag == "heading" {
			continue
		}
		position := lspEntityPosition(state.source, symbol.Position)
		tokens = append(tokens, token{
			line: position.Line, character: position.Character + 1,
			length: utf16Length(symbol.SourceTag), kind: lspSemanticKind(symbol.Concept),
		})
	}
	sort.SliceStable(tokens, func(left, right int) bool {
		if tokens[left].line != tokens[right].line {
			return tokens[left].line < tokens[right].line
		}
		return tokens[left].character < tokens[right].character
	})
	data := make([]int, 0, len(tokens)*5)
	previousLine, previousCharacter := 0, 0
	for index, current := range tokens {
		deltaLine := current.line - previousLine
		deltaCharacter := current.character
		if index > 0 && deltaLine == 0 {
			deltaCharacter -= previousCharacter
		}
		data = append(data, deltaLine, deltaCharacter, current.length, current.kind, 0)
		previousLine, previousCharacter = current.line, current.character
	}
	return map[string]any{"data": data}, nil
}

func (rcvr *Server) hover(ctx context.Context, raw json.RawMessage) (any, error) {
	var params struct {
		TextDocument struct {
			URI string `json:"uri"`
		} `json:"textDocument"`
		Position position `json:"position"`
	}
	if err := decodeParams(raw, &params); err != nil {
		return nil, err
	}
	state, err := rcvr.document(ctx, params.TextDocument.URI)
	if err != nil {
		return nil, err
	}
	var candidate *entity.ProjectSymbol
	for index := range state.analysis.Symbols {
		symbol := &state.analysis.Symbols[index]
		symbolRange := rangeForSymbol(state.source, *symbol)
		if !positionInRange(params.Position, symbolRange) {
			continue
		}
		if candidate == nil || symbol.Position.Offset > candidate.Position.Offset {
			candidate = symbol
		}
	}
	if candidate == nil {
		return nil, nil
	}
	value := fmt.Sprintf("**%s** `%s`\n\nID: `%s`", candidate.Concept, candidate.SourceTag, candidate.ID)
	if candidate.Detail != "" {
		value += "\n\n" + candidate.Detail
	}
	return map[string]any{
		"contents": map[string]any{"kind": "markdown", "value": value},
		rangeKey:   rangeForSymbol(state.source, *candidate),
	}, nil
}

const rangeKey = "range"

func (rcvr *Server) documentForParams(ctx context.Context, raw json.RawMessage) (*documentState, error) {
	var params textDocumentParams
	if err := decodeParams(raw, &params); err != nil {
		return nil, err
	}
	return rcvr.document(ctx, params.TextDocument.URI)
}

func (rcvr *Server) document(ctx context.Context, uri string) (*documentState, error) {
	if state := rcvr.documents[uri]; state != nil {
		return state, nil
	}
	path := lspURIPath(uri)
	if path == "" {
		return nil, invalidParams("document is not open and URI is not a local file")
	}
	source, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read LSP document %s: %w", uri, err)
	}
	analysis, err := rcvr.project.Analyze(ctx, uri, source)
	if err != nil {
		return nil, err
	}
	return &documentState{uri: uri, source: source, analysis: analysis}, nil
}

func (rcvr *Server) logError(err error) error {
	return rcvr.transport.notification("window/logMessage", map[string]any{"type": 1, "message": err.Error()})
}

type textDocumentParams struct {
	TextDocument struct {
		URI string `json:"uri"`
	} `json:"textDocument"`
}

type position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type lspRange struct {
	Start position `json:"start"`
	End   position `json:"end"`
}

type diagnostic struct {
	Range    lspRange `json:"range"`
	Severity int      `json:"severity"`
	Source   string   `json:"source"`
	Message  string   `json:"message"`
}

type documentSymbol struct {
	Name           string           `json:"name"`
	Detail         string           `json:"detail,omitempty"`
	Kind           int              `json:"kind"`
	Range          lspRange         `json:"range"`
	SelectionRange lspRange         `json:"selectionRange"`
	Children       []documentSymbol `json:"children,omitempty"`
}

type location struct {
	URI   string   `json:"uri"`
	Range lspRange `json:"range"`
}

type symbolInformation struct {
	Name          string   `json:"name"`
	Kind          int      `json:"kind"`
	Location      location `json:"location"`
	ContainerName string   `json:"containerName,omitempty"`
}

func lspDiagnostics(source []byte, values []entity.Diagnostic) []diagnostic {
	result := make([]diagnostic, 0, len(values))
	for _, value := range values {
		start := lspEntityPosition(source, entity.Position{Offset: value.Offset, Line: value.Line, Column: value.Column})
		severity := 1
		if value.Severity == "warning" {
			severity = 2
		}
		result = append(result, diagnostic{
			Range:    lspRange{Start: start, End: position{Line: start.Line, Character: start.Character + 1}},
			Severity: severity, Source: "xaligo", Message: value.Message,
		})
	}
	return result
}

func lspDocumentSymbols(source []byte, values []entity.ProjectSymbol) []documentSymbol {
	nodes := make([]documentSymbol, len(values))
	byOrdinal := make(map[int]int, len(values))
	for index, value := range values {
		rangeValue := rangeForSymbol(source, value)
		nodes[index] = documentSymbol{
			Name: value.Name, Detail: value.Detail, Kind: lspSymbolKind(value.Concept),
			Range: rangeValue, SelectionRange: rangeValue,
		}
		byOrdinal[value.Ordinal] = index
	}
	children := make(map[int][]int, len(values))
	var roots []int
	for index, value := range values {
		if parentIndex, exists := byOrdinal[value.ParentOrdinal]; exists {
			children[parentIndex] = append(children[parentIndex], index)
		} else {
			roots = append(roots, index)
		}
	}
	var build func(int) documentSymbol
	build = func(index int) documentSymbol {
		result := nodes[index]
		for _, child := range children[index] {
			resolvedChild := build(child)
			result.Children = append(result.Children, resolvedChild)
			if positionBefore(result.Range.End, resolvedChild.Range.End) {
				result.Range.End = resolvedChild.Range.End
			}
		}
		return result
	}
	result := make([]documentSymbol, 0, len(roots))
	for _, root := range roots {
		result = append(result, build(root))
	}
	return result
}

func lspSymbolInformation(uri string, source []byte, symbol entity.ProjectSymbol, all []entity.ProjectSymbol) symbolInformation {
	container := ""
	for _, candidate := range all {
		if candidate.Ordinal == symbol.ParentOrdinal {
			container = candidate.Name
			break
		}
	}
	return symbolInformation{
		Name: symbol.Name, Kind: lspSymbolKind(symbol.Concept),
		Location:      location{URI: uri, Range: rangeForSymbol(source, symbol)},
		ContainerName: container,
	}
}

func rangeForSymbol(source []byte, symbol entity.ProjectSymbol) lspRange {
	start := lspEntityPosition(source, symbol.Position)
	if symbol.SourceTag != "" && symbol.SourceTag != "heading" {
		start.Character++
	}
	length := utf16Length(symbol.SourceTag)
	if length == 0 {
		length = max(1, utf16Length(symbol.Name))
	}
	return lspRange{Start: start, End: position{Line: start.Line, Character: start.Character + length}}
}

func rangeFromLineColumn(line, column int, name string) lspRange {
	start := position{Line: max(0, line-1), Character: max(0, column-1)}
	return lspRange{Start: start, End: position{Line: start.Line, Character: start.Character + max(1, utf16Length(name))}}
}

func positionInRange(value position, valueRange lspRange) bool {
	return !positionBefore(value, valueRange.Start) && positionBefore(value, valueRange.End)
}

func positionBefore(left, right position) bool {
	return left.Line < right.Line || (left.Line == right.Line && left.Character < right.Character)
}

func lspEntityPosition(source []byte, value entity.Position) position {
	line := max(0, value.Line-1)
	character := max(0, value.Column-1)
	if value.Offset >= 0 && value.Offset <= len(source) {
		prefix := source[:value.Offset]
		lineStart := 0
		if index := strings.LastIndexByte(string(prefix), '\n'); index >= 0 {
			lineStart = index + 1
		}
		character = utf16Length(string(prefix[lineStart:]))
	}
	return position{Line: line, Character: character}
}

func lspSymbolKind(concept entity.ProjectConcept) int {
	switch concept {
	case entity.ProjectConceptFrame:
		return 2
	case entity.ProjectConceptGroup:
		return 3
	case entity.ProjectConceptCapture:
		return 24
	case entity.ProjectConceptItem:
		return 19
	case entity.ProjectConceptPort:
		return 11
	case entity.ProjectConceptLine:
		return 24
	case entity.ProjectConceptText:
		return 15
	case entity.ProjectConceptSpacer:
		return 13
	default:
		return 19
	}
}

func lspSemanticKind(concept entity.ProjectConcept) int {
	switch concept {
	case entity.ProjectConceptFrame, entity.ProjectConceptGroup:
		return 0
	case entity.ProjectConceptItem:
		return 1
	case entity.ProjectConceptPort:
		return 2
	case entity.ProjectConceptCapture, entity.ProjectConceptLine:
		return 3
	case entity.ProjectConceptText:
		return 4
	default:
		return 5
	}
}

func projectSymbolMatches(symbol entity.ProjectSymbol, query string) bool {
	needle := strings.ToLower(query)
	haystack := strings.ToLower(strings.Join([]string{symbol.ID, symbol.Name, symbol.Detail, symbol.SourceTag}, " "))
	return strings.Contains(haystack, needle)
}

func lspFTSQuery(query string) string {
	parts := strings.Fields(query)
	quoted := make([]string, 0, len(parts))
	for _, part := range parts {
		quoted = append(quoted, `"`+strings.ReplaceAll(part, `"`, `""`)+`"`)
	}
	return strings.Join(quoted, " AND ")
}

func lspURIPath(uri string) string {
	parsed, err := url.Parse(uri)
	if err != nil || parsed.Scheme != "file" {
		return ""
	}
	return filepath.FromSlash(parsed.Path)
}

func utf16Length(value string) int {
	return len(utf16.Encode([]rune(value)))
}

func decodeParams(raw json.RawMessage, destination any) error {
	if len(raw) == 0 || string(raw) == "null" {
		return invalidParams("params are required")
	}
	if err := json.Unmarshal(raw, destination); err != nil {
		return invalidParams("invalid params: " + err.Error())
	}
	return nil
}

func contextError(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	return ctx.Err()
}
