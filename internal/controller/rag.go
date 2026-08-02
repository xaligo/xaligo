package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
)

var (
	ICRAGCOMMAND001 = share.NewMCode("ICRAGCOMMAND-001", "Initialize RAG command")
	ICRAGINDEX001   = share.NewMCode("ICRAGINDEX-001", "Index documentation corpus")
	ICRAGSEARCH001  = share.NewMCode("ICRAGSEARCH-001", "Search project knowledge index")
	ICRAGWATCH001   = share.NewMCode("ICRAGWATCH-001", "Watch documentation corpus")
)

// RAGController adapts the project concept index to the CLI command family.
type RAGController interface {
	Command() *cobra.Command
	RunIndex(context.Context, string, io.Writer) (entity.ProjectIndexStats, error)
	RunSearch(context.Context, string, int, bool, io.Writer) error
	RunWatch(context.Context, string, time.Duration, io.Writer) error
}

type ragController struct {
	projectUsecase usecase.ProjectUsecase
	defaultRoot    string
}

func NewRAGController(projectUsecase usecase.ProjectUsecase, defaultRoot string) RAGController {
	return &ragController{projectUsecase: projectUsecase, defaultRoot: defaultRoot}
}

func (rcvr *ragController) Command() *cobra.Command {
	logger.DEBUG(ICRAGCOMMAND001, "start")
	command := &cobra.Command{
		Use:   "rag",
		Short: "Index and search the local documentation knowledge base",
		Long: `Build and query xaligo's local SQLite/FTS5 knowledge index.

The initial RAG corpus is deliberately limited to Markdown files below docs/.
Sample .xal files and other repository content are not registered implicitly.
The same database can hold explicitly analyzed .xal concepts for editor and
agent requests without adding them to the initial documentation corpus.`,
	}
	command.AddCommand(rcvr.indexCommand())
	command.AddCommand(rcvr.searchCommand())
	command.AddCommand(rcvr.watchCommand())
	return command
}

func (rcvr *ragController) RunIndex(ctx context.Context, root string, output io.Writer) (entity.ProjectIndexStats, error) {
	if rcvr.projectUsecase == nil {
		return entity.ProjectIndexStats{}, errors.New("project use case is required")
	}
	if strings.TrimSpace(root) == "" {
		root = rcvr.defaultRoot
	}
	stats, err := rcvr.projectUsecase.Index(ctx, root)
	if err != nil {
		return entity.ProjectIndexStats{}, err
	}
	logger.INFO(ICRAGINDEX001, "documentation indexed", map[string]any{
		"root": stats.Root, "scanned": stats.Scanned, "indexed": stats.Indexed,
		"unchanged": stats.Unchanged, "removed": stats.Removed,
	})
	if output != nil {
		if err := writeRAGJSON(output, stats); err != nil {
			return entity.ProjectIndexStats{}, err
		}
	}
	return stats, nil
}

func (rcvr *ragController) RunSearch(ctx context.Context, query string, limit int, asJSON bool, output io.Writer) error {
	if rcvr.projectUsecase == nil {
		return errors.New("project use case is required")
	}
	results, err := rcvr.projectUsecase.Search(ctx, query, limit)
	if err != nil {
		return err
	}
	logger.DEBUG(ICRAGSEARCH001, "search complete", map[string]any{"query": query, "results": len(results)})
	if output == nil {
		return nil
	}
	if asJSON {
		return writeRAGJSON(output, results)
	}
	for _, result := range results {
		detail := strings.Join(strings.Fields(result.Detail), " ")
		if _, err := fmt.Fprintf(output, "%s:%d:%d\t%s\t%s\t%s\n",
			result.URI, result.Line, result.Column, result.Concept, result.Name, detail); err != nil {
			return fmt.Errorf("write RAG search result: %w", err)
		}
	}
	return nil
}

func (rcvr *ragController) RunWatch(ctx context.Context, root string, interval time.Duration, output io.Writer) error {
	if interval < 100*time.Millisecond {
		return errors.New("RAG watch interval must be at least 100ms")
	}
	if _, err := rcvr.RunIndex(ctx, root, output); err != nil {
		return err
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			stats, err := rcvr.RunIndex(ctx, root, nil)
			if err != nil {
				return err
			}
			if stats.Indexed > 0 || stats.Removed > 0 {
				logger.INFO(ICRAGWATCH001, "documentation index changed", map[string]any{
					"indexed": stats.Indexed, "removed": stats.Removed,
				})
				if output != nil {
					if err := writeRAGJSON(output, stats); err != nil {
						return err
					}
				}
			}
		}
	}
}

func (rcvr *ragController) indexCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "index [project-root|docs-root]",
		Short: "Index Markdown files below docs/ incrementally",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			root := ""
			if len(args) == 1 {
				root = args[0]
			}
			_, err := rcvr.RunIndex(command.Context(), root, command.OutOrStdout())
			return err
		},
	}
}

func (rcvr *ragController) searchCommand() *cobra.Command {
	var limit int
	var asJSON bool
	command := &cobra.Command{
		Use:   "search <FTS5 query>",
		Short: "Search indexed documentation and explicit concepts",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			return rcvr.RunSearch(command.Context(), strings.Join(args, " "), limit, asJSON, command.OutOrStdout())
		},
	}
	command.Flags().IntVar(&limit, "limit", 30, "maximum results (1-100)")
	command.Flags().BoolVar(&asJSON, "json", false, "write one JSON array instead of tab-separated rows")
	return command
}

func (rcvr *ragController) watchCommand() *cobra.Command {
	var interval time.Duration
	command := &cobra.Command{
		Use:   "watch [project-root|docs-root]",
		Short: "Continuously refresh the Markdown documentation index",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			root := ""
			if len(args) == 1 {
				root = args[0]
			}
			return rcvr.RunWatch(command.Context(), root, interval, command.OutOrStdout())
		},
	}
	command.Flags().DurationVar(&interval, "interval", 2*time.Second, "poll interval")
	return command
}

func writeRAGJSON(output io.Writer, value any) error {
	encoder := json.NewEncoder(output)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		return fmt.Errorf("write RAG JSON: %w", err)
	}
	return nil
}
