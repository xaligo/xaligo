package controller_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/controller"
	"github.com/xaligo/xaligo/internal/mcp"
)

type fakeControllerMCPTools struct{}

func (*fakeControllerMCPTools) Tools() []mcp.Tool { return nil }

func (*fakeControllerMCPTools) Call(context.Context, string, json.RawMessage) (mcp.ToolResult, error) {
	return mcp.ToolResult{}, mcp.ErrUnknownTool
}

func TestMCPCommandRejectsNonLoopbackHTTPBinding(t *testing.T) {
	command := controller.NewMCPController(&fakeControllerMCPTools{}).Command()
	command.SetArgs([]string{"--http", "--address", "0.0.0.0:8081"})
	err := command.Execute()
	if err == nil || !strings.Contains(err.Error(), "must bind to localhost") {
		t.Fatalf("non-loopback bind error = %v", err)
	}
}

func TestMCPCommandRejectsBothTransports(t *testing.T) {
	command := controller.NewMCPController(&fakeControllerMCPTools{}).Command()
	command.SetArgs([]string{"--stdio", "--http"})
	err := command.Execute()
	if err == nil || !strings.Contains(err.Error(), "mutually exclusive") {
		t.Fatalf("transport conflict error = %v", err)
	}
}
