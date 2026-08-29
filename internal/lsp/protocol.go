// Package lsp implements the xaligo Language Server Protocol stdio adapter.
package lsp

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
)

const (
	maxHeaderBytes  = 16 * 1024
	maxContentBytes = 8 * 1024 * 1024
)

type rpcMessage struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcNotification struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type protocolError struct {
	code    int
	message string
	data    any
}

func (rcvr *protocolError) Error() string { return rcvr.message }

func invalidParams(message string) error {
	return &protocolError{code: -32602, message: message}
}

type transport struct {
	reader *bufio.Reader
	writer io.Writer
	mu     sync.Mutex
}

func newTransport(input io.Reader, output io.Writer) *transport {
	return &transport{reader: bufio.NewReader(input), writer: output}
}

func (rcvr *transport) read() ([]byte, error) {
	contentLength := -1
	headerBytes := 0
	for {
		line, err := rcvr.reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) && line == "" && headerBytes == 0 {
				return nil, io.EOF
			}
			return nil, fmt.Errorf("read LSP header: %w", err)
		}
		headerBytes += len(line)
		if headerBytes > maxHeaderBytes {
			return nil, fmt.Errorf("LSP header exceeds %d bytes", maxHeaderBytes)
		}
		line = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
		if line == "" {
			break
		}
		name, value, found := strings.Cut(line, ":")
		if !found {
			return nil, fmt.Errorf("invalid LSP header line %q", line)
		}
		if strings.EqualFold(strings.TrimSpace(name), "Content-Length") {
			parsed, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil || parsed < 0 {
				return nil, fmt.Errorf("invalid LSP Content-Length %q", strings.TrimSpace(value))
			}
			contentLength = parsed
		}
	}
	if contentLength < 0 {
		return nil, errors.New("LSP Content-Length header is required")
	}
	if contentLength > maxContentBytes {
		return nil, fmt.Errorf("LSP message size %d exceeds %d", contentLength, maxContentBytes)
	}
	payload := make([]byte, contentLength)
	if _, err := io.ReadFull(rcvr.reader, payload); err != nil {
		return nil, fmt.Errorf("read LSP payload: %w", err)
	}
	return payload, nil
}

func (rcvr *transport) response(id json.RawMessage, result any, responseError *rpcError) error {
	if len(id) == 0 {
		id = json.RawMessage("null")
	}
	if responseError != nil {
		return rcvr.write(map[string]any{"jsonrpc": "2.0", "id": id, "error": responseError})
	}
	// JSON-RPC requires a successful response to contain result, including the
	// explicit null result returned by LSP shutdown.
	return rcvr.write(map[string]any{"jsonrpc": "2.0", "id": id, "result": result})
}

func (rcvr *transport) notification(method string, params any) error {
	return rcvr.write(rpcNotification{JSONRPC: "2.0", Method: method, Params: params})
}

func (rcvr *transport) write(message any) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("encode LSP message: %w", err)
	}
	var framed bytes.Buffer
	fmt.Fprintf(&framed, "Content-Length: %d\r\n\r\n", len(payload))
	framed.Write(payload)
	rcvr.mu.Lock()
	defer rcvr.mu.Unlock()
	if _, err := rcvr.writer.Write(framed.Bytes()); err != nil {
		return fmt.Errorf("write LSP message: %w", err)
	}
	return nil
}

func messageHasID(message rpcMessage) bool {
	return len(message.ID) > 0 && string(message.ID) != "null"
}
