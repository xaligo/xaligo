package controller

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
	"github.com/xaligo/xaligo/internal/usecase"
)

var (
	ICDIFF001 = share.NewMCode("ICDIFF-001", "Diff read before failed")
	ICDIFF002 = share.NewMCode("ICDIFF-002", "Diff read after failed")
	ICDIFF003 = share.NewMCode("ICDIFF-003", "Diff use case failed")
	ICDIFF004 = share.NewMCode("ICDIFF-004", "Diff output write failed")
	ICDIFF005 = share.NewMCode("ICDIFF-005", "Diff output generated")
)

type DiffController interface {
	Command() *cobra.Command
	Run(entity.ControllerDiffOptions) error
}

type diffController struct {
	diffUsecase usecase.DiffUsecase
}

func NewDiffController(diffUsecase usecase.DiffUsecase) DiffController {
	return &diffController{diffUsecase: diffUsecase}
}

type pendingDiffImage struct {
	target string
	data   []byte
	temp   string
}

func (rcvr *diffController) Command() *cobra.Command {
	var outputPrefix, theme, mode string
	var pxPerInch float64
	command := &cobra.Command{
		Use:   "diff <before.xal> <after.xal>",
		Short: "Render structural differences as removed and added SVG images",
		Args:  cobra.ExactArgs(2),
		RunE: func(command *cobra.Command, args []string) error {
			return rcvr.Run(entity.ControllerDiffOptions{
				BeforePath: args[0], AfterPath: args[1], OutputPrefix: outputPrefix,
				Theme: theme, Mode: mode, PxPerInch: pxPerInch, Stdout: os.Stdout,
			})
		},
	}
	command.Flags().StringVarP(&outputPrefix, "output", "o", "output", "output prefix; writes <prefix>-removed.svg and <prefix>-added.svg")
	command.Flags().StringVar(&theme, "theme", "light", "color theme: light | dark")
	command.Flags().StringVar(&mode, "mode", "standard", "rendering mode: standard | network | aws")
	command.Flags().Float64Var(&pxPerInch, "px-per-inch", 0, "pixels per inch for SVG layout scaling (default 96)")
	return command
}

func (rcvr *diffController) Run(opts entity.ControllerDiffOptions) error {
	before, err := os.ReadFile(opts.BeforePath)
	if err != nil {
		logger.ERROR(ICDIFF001, "read before failed", map[string]any{"path": opts.BeforePath, "error": err})
		return fmt.Errorf("read before file: %w", err)
	}
	after, err := os.ReadFile(opts.AfterPath)
	if err != nil {
		logger.ERROR(ICDIFF002, "read after failed", map[string]any{"path": opts.AfterPath, "error": err})
		return fmt.Errorf("read after file: %w", err)
	}
	result, err := rcvr.diffUsecase.Diff(context.Background(), before, after, entity.DiffOptions{
		Mode: entity.Mode(opts.Mode), Theme: opts.Theme, PxPerInch: opts.PxPerInch,
	})
	if err != nil {
		logger.ERROR(ICDIFF003, "use case failed", map[string]any{"error": err})
		return err
	}
	removedPath, addedPath := diffOutputPaths(opts.OutputPrefix)
	if err := writeDiffImagePair(removedPath, result.RemovedImage, addedPath, result.AddedImage); err != nil {
		logger.ERROR(ICDIFF004, "output write failed", map[string]any{"error": err})
		return err
	}
	logger.INFO(ICDIFF005, "generated", map[string]any{"removed": removedPath, "added": addedPath})
	if opts.Stdout != nil {
		fmt.Fprintf(opts.Stdout, "removed: %s\nadded: %s\nchanges: +%d -%d ~%d\n", removedPath, addedPath, result.Summary.AddedCount, result.Summary.RemovedCount, result.Summary.ModifiedCount)
	}
	return nil
}

func diffOutputPaths(prefix string) (string, string) {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" {
		prefix = "output"
	}
	if strings.EqualFold(filepath.Ext(prefix), ".svg") {
		prefix = strings.TrimSuffix(prefix, filepath.Ext(prefix))
	}
	return prefix + "-removed.svg", prefix + "-added.svg"
}

func writeDiffImagePair(removedPath string, removedImage []byte, addedPath string, addedImage []byte) error {
	if removedPath == addedPath {
		return fmt.Errorf("diff output paths must be distinct")
	}
	for _, path := range []string{removedPath, addedPath} {
		directory := filepath.Dir(path)
		if err := os.MkdirAll(directory, 0755); err != nil {
			return fmt.Errorf("create diff output directory %s: %w", directory, err)
		}
	}
	images := []pendingDiffImage{{target: removedPath, data: removedImage}, {target: addedPath, data: addedImage}}
	for index := range images {
		file, err := os.CreateTemp(filepath.Dir(images[index].target), ".xaligo-diff-*.svg")
		if err != nil {
			cleanupDiffTemps(images)
			return fmt.Errorf("create temporary diff image: %w", err)
		}
		images[index].temp = file.Name()
		if err := file.Chmod(0644); err != nil {
			_ = file.Close()
			cleanupDiffTemps(images)
			return fmt.Errorf("set temporary diff image permissions: %w", err)
		}
		if _, err := file.Write(images[index].data); err != nil {
			_ = file.Close()
			cleanupDiffTemps(images)
			return fmt.Errorf("write temporary diff image: %w", err)
		}
		if err := file.Close(); err != nil {
			cleanupDiffTemps(images)
			return fmt.Errorf("close temporary diff image: %w", err)
		}
	}
	for index := range images {
		if err := os.Rename(images[index].temp, images[index].target); err != nil {
			cleanupDiffTemps(images)
			return fmt.Errorf("replace diff output %s: %w", images[index].target, err)
		}
		images[index].temp = ""
	}
	return nil
}

func cleanupDiffTemps(images []pendingDiffImage) {
	for _, image := range images {
		if image.temp != "" {
			_ = os.Remove(image.temp)
		}
	}
}
