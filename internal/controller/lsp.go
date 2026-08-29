package controller

import (
	"context"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/lsp"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
)

var ICLSPCOMMAND001 = share.NewMCode("ICLSPCOMMAND-001", "Run LSP stdio server")

type LSPController interface {
	Command() *cobra.Command
	Run(context.Context, io.Reader, io.Writer) error
}

type lspController struct {
	projectUsecase usecase.ProjectUsecase
}

func NewLSPController(projectUsecase usecase.ProjectUsecase) LSPController {
	return &lspController{projectUsecase: projectUsecase}
}

func (rcvr *lspController) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "lsp",
		Short: "Run the xaligo language server over stdio",
		Long: `Run the Language Server Protocol 3.18 adapter over standard input and
standard output. The server uses Content-Length framed JSON-RPC and provides
full-document synchronization, push and pull diagnostics, document/workspace
symbols, semantic tokens, and hover details for .xal concepts.

The server reuses the shared project parser and diagnostic analysis. Saving an
open .xal document updates its explicit concept rows in the local project
database; it does not broaden the docs-only initial RAG corpus.`,
		Args: cobra.NoArgs,
		RunE: func(command *cobra.Command, _ []string) error {
			return rcvr.Run(command.Context(), os.Stdin, os.Stdout)
		},
	}
}

func (rcvr *lspController) Run(ctx context.Context, input io.Reader, output io.Writer) error {
	share.UseStderrForProtocol()
	logger.INFO(ICLSPCOMMAND001, "LSP server started")
	return lsp.NewServer(rcvr.projectUsecase).Serve(ctx, input, output)
}
