package controller

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/mcp"
	"github.com/xaligo/xaligo/internal/share"
)

const defaultMCPAddress = "127.0.0.1:8081"

var (
	ICMCPCOMMAND001 = share.NewMCode("ICMCPCOMMAND-001", "Run MCP server")
	ICMCPHTTP001    = share.NewMCode("ICMCPHTTP-001", "Run MCP Streamable HTTP server")
)

type MCPController interface {
	Command() *cobra.Command
	RunStdio(context.Context, io.Reader, io.Writer) error
	RunHTTP(context.Context, string) error
}

type mcpController struct {
	server *mcp.Server
}

func NewMCPController(tools mcp.ToolService) MCPController {
	return &mcpController{server: mcp.NewServer(tools, resolvedVersion())}
}

func (rcvr *mcpController) Command() *cobra.Command {
	var stdio, streamableHTTP bool
	var address string
	command := &cobra.Command{
		Use:   "mcp",
		Short: "Expose xaligo operations through Model Context Protocol",
		Long: `Run the stateless MCP 2026-07-28 adapter through newline-delimited
stdio or a localhost-only Streamable HTTP POST endpoint.

The tools reuse xaligo's shared diagnostics, SVG rendering, project search,
and icon registry use cases. index_docs discovers only Markdown below docs/;
.xal source is analyzed only when a tool call supplies it explicitly.

Examples:
  xaligo mcp --stdio
  xaligo mcp --http --address 127.0.0.1:8081`,
		Args: cobra.NoArgs,
		RunE: func(command *cobra.Command, _ []string) error {
			if stdio && streamableHTTP {
				return errors.New("--stdio and --http are mutually exclusive")
			}
			if streamableHTTP {
				return rcvr.RunHTTP(command.Context(), address)
			}
			return rcvr.RunStdio(command.Context(), os.Stdin, os.Stdout)
		},
	}
	command.Flags().BoolVar(&stdio, "stdio", false, "serve newline-delimited MCP over stdin/stdout (default)")
	command.Flags().BoolVar(&streamableHTTP, "http", false, "serve Streamable HTTP POST at /mcp")
	command.Flags().StringVar(&address, "address", defaultMCPAddress, "localhost address for --http")
	return command
}

func (rcvr *mcpController) RunStdio(ctx context.Context, input io.Reader, output io.Writer) error {
	share.UseStderrForProtocol()
	logger.INFO(ICMCPCOMMAND001, "MCP stdio server started")
	return rcvr.server.ServeStdio(ctx, input, output)
}

func (rcvr *mcpController) RunHTTP(ctx context.Context, address string) error {
	if !isLoopbackMCPAddress(address) {
		return fmt.Errorf("MCP HTTP address %q must bind to localhost or a loopback IP", address)
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("listen for MCP HTTP on %s: %w", address, err)
	}
	defer listener.Close()
	mux := http.NewServeMux()
	mux.Handle("/mcp", mcp.NewHTTPHandler(rcvr.server))
	server := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      2 * time.Minute,
		IdleTimeout:       30 * time.Second,
	}
	logger.INFO(ICMCPHTTP001, "MCP HTTP server started", map[string]any{"address": listener.Addr().String(), "path": "/mcp"})
	completed := make(chan error, 1)
	go func() { completed <- server.Serve(listener) }()
	select {
	case err := <-completed:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		shutdownContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownContext); err != nil {
			return fmt.Errorf("shutdown MCP HTTP server: %w", err)
		}
		return ctx.Err()
	}
}

func isLoopbackMCPAddress(address string) bool {
	host, port, err := net.SplitHostPort(strings.TrimSpace(address))
	if err != nil || port == "" {
		return false
	}
	host = strings.TrimSuffix(strings.ToLower(host), ".")
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}
