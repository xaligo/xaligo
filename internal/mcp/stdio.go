package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
)

type stdioWriter struct {
	output io.Writer
	mu     sync.Mutex
}

func (rcvr *stdioWriter) write(payload []byte) error {
	if len(payload) == 0 {
		return nil
	}
	rcvr.mu.Lock()
	defer rcvr.mu.Unlock()
	if _, err := rcvr.output.Write(append(append([]byte(nil), payload...), '\n')); err != nil {
		return fmt.Errorf("write MCP stdio response: %w", err)
	}
	return nil
}

type activeRequest struct {
	cancelled bool
	cancel    context.CancelFunc
}

// ServeStdio serves newline-delimited MCP JSON-RPC until input reaches EOF.
func (rcvr *Server) ServeStdio(ctx context.Context, input io.Reader, output io.Writer) error {
	if err := validateServer(rcvr); err != nil {
		return err
	}
	if input == nil || output == nil {
		return errors.New("MCP stdio input and output are required")
	}
	writer := &stdioWriter{output: output}
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 64*1024), maxMessageBytes)
	active := map[string]*activeRequest{}
	var activeMu sync.Mutex
	var requests sync.WaitGroup
	var writeErr error
	var writeErrMu sync.Mutex

	for scanner.Scan() {
		payload := append([]byte(nil), scanner.Bytes()...)
		message, decodeErr := decodeMessage(payload)
		if decodeErr != nil {
			if err := writer.write(rcvr.Handle(ctx, payload).Body); err != nil {
				return err
			}
			continue
		}
		if message.Method == "notifications/cancelled" {
			key := cancelledRequestKey(message.Params)
			activeMu.Lock()
			if request := active[key]; request != nil {
				request.cancelled = true
				request.cancel()
			}
			activeMu.Unlock()
			continue
		}
		if len(message.ID) == 0 {
			_ = rcvr.Handle(ctx, payload)
			continue
		}
		key := requestIDKey(message.ID)
		requestContext, cancel := context.WithCancel(ctx)
		state := &activeRequest{cancel: cancel}
		activeMu.Lock()
		if _, duplicate := active[key]; duplicate {
			activeMu.Unlock()
			cancel()
			if err := writer.write(protocolErrorBody(message.ID, -32600, "Duplicate in-flight request id", nil)); err != nil {
				return err
			}
			continue
		}
		active[key] = state
		activeMu.Unlock()

		requests.Add(1)
		go func() {
			defer requests.Done()
			defer cancel()
			result := rcvr.Handle(requestContext, payload)
			activeMu.Lock()
			cancelled := state.cancelled
			delete(active, key)
			activeMu.Unlock()
			if cancelled || len(result.Body) == 0 {
				return
			}
			if err := writer.write(result.Body); err != nil {
				writeErrMu.Lock()
				writeErr = errors.Join(writeErr, err)
				writeErrMu.Unlock()
			}
		}()
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read MCP stdio message: %w", err)
	}
	requests.Wait()
	writeErrMu.Lock()
	defer writeErrMu.Unlock()
	return writeErr
}

func cancelledRequestKey(raw json.RawMessage) string {
	var params struct {
		RequestID json.RawMessage `json:"requestId"`
	}
	if json.Unmarshal(raw, &params) != nil {
		return ""
	}
	return requestIDKey(params.RequestID)
}
