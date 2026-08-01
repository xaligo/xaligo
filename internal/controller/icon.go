package controller

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/entity"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

const maxIconSVGInputBytes = 2 * 1024 * 1024

type IconController interface {
	Command() *cobra.Command
}

type iconController struct {
	iconUsecase v2.IconUsecase
}

func NewIconController(iconUsecase v2.IconUsecase) IconController {
	return &iconController{iconUsecase: iconUsecase}
}

func (rcvr *iconController) Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "icon",
		Short: "Manage the embedded SVG icon registry",
		Long: `Register, retrieve, search, and remove namespaced SVG icons in the local
xaligo-assets.db registry. SVG input is normalized and safety-checked by the
Rust V2 engine before it is stored and indexed by SQLite FTS5.`,
	}
	command.AddCommand(rcvr.addCommand())
	command.AddCommand(rcvr.getCommand())
	command.AddCommand(rcvr.searchCommand())
	command.AddCommand(rcvr.removeCommand())
	command.AddCommand(rcvr.listCommand())
	command.AddCommand(rcvr.namespacesCommand())
	return command
}

func (rcvr *iconController) addCommand() *cobra.Command {
	var reference, description, license, source string
	var tags, aliases []string
	command := &cobra.Command{
		Use:   "add <icon.svg>",
		Short: "Add or update an SVG icon",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			data, err := readIconSVG(args[0])
			if err != nil {
				return err
			}
			stored, err := rcvr.iconUsecase.Put(command.Context(), entity.IconRegistration{
				Reference: reference, SVG: data, Description: description, Tags: tags,
				Aliases: aliases, License: license, Source: source,
			})
			if err != nil {
				return err
			}
			if _, err := fmt.Fprintf(command.OutOrStdout(), "%s\t%s\n", stored.Ref.String(), hex.EncodeToString(stored.Checksum[:])); err != nil {
				return fmt.Errorf("write icon add output: %w", err)
			}
			return nil
		},
	}
	command.Flags().StringVar(&reference, "name", "", "required icon identity as namespace:name")
	command.Flags().StringVar(&description, "description", "", "searchable icon description")
	command.Flags().StringSliceVar(&tags, "tag", nil, "searchable tag; repeat or use comma-separated values")
	command.Flags().StringSliceVar(&aliases, "alias", nil, "alias as name or namespace:name; repeat as needed")
	command.Flags().StringVar(&license, "license", "", "license identifier or notice")
	command.Flags().StringVar(&source, "source", "", "source attribution or URL")
	_ = command.MarkFlagRequired("name")
	return command
}

func readIconSVG(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open SVG icon %s: %w", path, err)
	}
	defer file.Close()
	data, err := io.ReadAll(io.LimitReader(file, maxIconSVGInputBytes+1))
	if err != nil {
		return nil, fmt.Errorf("read SVG icon %s: %w", path, err)
	}
	if len(data) > maxIconSVGInputBytes {
		return nil, fmt.Errorf("SVG icon %s exceeds %d bytes", path, maxIconSVGInputBytes)
	}
	return data, nil
}

func (rcvr *iconController) getCommand() *cobra.Command {
	var output string
	command := &cobra.Command{
		Use:   "get <namespace:name>",
		Short: "Write one normalized SVG icon",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			icon, err := rcvr.iconUsecase.Get(command.Context(), args[0])
			if err != nil {
				return err
			}
			if output == "" || output == "-" {
				if _, err := command.OutOrStdout().Write(icon.SVG); err != nil {
					return fmt.Errorf("write icon SVG output: %w", err)
				}
				return nil
			}
			if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
				return fmt.Errorf("create icon output directory: %w", err)
			}
			if err := os.WriteFile(output, icon.SVG, 0o644); err != nil {
				return fmt.Errorf("write icon SVG %s: %w", output, err)
			}
			return nil
		},
	}
	command.Flags().StringVarP(&output, "output", "o", "-", "output SVG path, or - for stdout")
	return command
}

func (rcvr *iconController) searchCommand() *cobra.Command {
	var limit int
	command := &cobra.Command{
		Use:   "search <FTS5 query>",
		Short: "Search icon names, descriptions, tags, and aliases",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			results, err := rcvr.iconUsecase.Search(command.Context(), strings.Join(args, " "), limit)
			if err != nil {
				return err
			}
			return writeIconSummaries(command, results)
		},
	}
	command.Flags().IntVar(&limit, "limit", 30, "maximum results (1-100)")
	return command
}

func (rcvr *iconController) removeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "remove <namespace:name>",
		Short: "Remove an icon and its tags and aliases",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, args []string) error {
			if err := rcvr.iconUsecase.Delete(command.Context(), args[0]); err != nil {
				return err
			}
			if _, err := fmt.Fprintln(command.OutOrStdout(), args[0]); err != nil {
				return fmt.Errorf("write icon remove output: %w", err)
			}
			return nil
		},
	}
}

func (rcvr *iconController) listCommand() *cobra.Command {
	var namespace string
	var limit int
	command := &cobra.Command{
		Use:   "list",
		Short: "List icons in stable identity order",
		RunE: func(command *cobra.Command, _ []string) error {
			results, err := rcvr.iconUsecase.List(command.Context(), namespace, limit)
			if err != nil {
				return err
			}
			return writeIconSummaries(command, results)
		},
	}
	command.Flags().StringVar(&namespace, "namespace", "", "only list one namespace")
	command.Flags().IntVar(&limit, "limit", 30, "maximum results (1-100)")
	return command
}

func (rcvr *iconController) namespacesCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "namespaces",
		Short: "List registered icon namespaces",
		RunE: func(command *cobra.Command, _ []string) error {
			namespaces, err := rcvr.iconUsecase.ListNamespaces(command.Context())
			if err != nil {
				return err
			}
			for _, namespace := range namespaces {
				if _, err := fmt.Fprintln(command.OutOrStdout(), namespace); err != nil {
					return fmt.Errorf("write icon namespace output: %w", err)
				}
			}
			return nil
		},
	}
}

func writeIconSummaries(command *cobra.Command, summaries []entity.IconSummary) error {
	for _, summary := range summaries {
		if _, err := fmt.Fprintf(command.OutOrStdout(), "%s\t%s\n", summary.Ref.String(), summary.Description); err != nil {
			return fmt.Errorf("write icon list output: %w", err)
		}
	}
	return nil
}
