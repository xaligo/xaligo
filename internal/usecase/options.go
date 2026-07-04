package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

const (
	ModeStandard entity.Mode = "standard"
	ModeNetwork  entity.Mode = "network"
	ModeAWS      entity.Mode = "aws"

	FormatExcalidraw entity.Format = "excalidraw"
	FormatSVG        entity.Format = "svg"
	FormatPPTX       entity.Format = "pptx"
	FormatXYFlow     entity.Format = "xyflow"
	FormatIsoflow    entity.Format = "isoflow"

	SeverityError entity.DiagnosticSeverity = "error"
)

var ErrNotImplemented = errors.New("renderer not implemented")

// ValidateRenderOptions validates mode, format, assets, and shared presentation
// values without parsing an input document.
func ValidateRenderOptions(opts entity.RenderOptions) error {
	if err := validateMode(opts.Mode); err != nil {
		return err
	}
	if opts.PaperMarginIn < 0 || opts.PaperMarginTopIn < 0 || opts.PaperMarginRightIn < 0 || opts.PaperMarginBottomIn < 0 || opts.PaperMarginLeftIn < 0 {
		return fmt.Errorf("paper margins must be non-negative")
	}
	if _, err := entity.NormalizeTheme(opts.Theme); err != nil {
		return err
	}
	if opts.Assets != nil {
		if opts.Assets.FS == nil {
			return fmt.Errorf("asset source filesystem is required")
		}
		if strings.TrimSpace(opts.Assets.CatalogCSV) == "" || strings.TrimSpace(opts.Assets.GroupIconsDir) == "" {
			return fmt.Errorf("asset source catalog and group icons directory are required")
		}
	}
	format := entity.Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	switch format {
	case "", FormatExcalidraw, FormatSVG, FormatPPTX, FormatXYFlow, FormatIsoflow:
		return nil
	default:
		return fmt.Errorf("unknown render format %q", format)
	}
}

func validateMode(mode entity.Mode) error {
	normalized := entity.Mode(strings.ToLower(strings.TrimSpace(string(mode))))
	switch normalized {
	case "", ModeStandard, ModeNetwork, ModeAWS:
		return nil
	case "aws-2.5d", "topology":
		return fmt.Errorf("mode %q: %w", normalized, ErrNotImplemented)
	default:
		return fmt.Errorf("unknown render mode %q", normalized)
	}
}

var (
	IURSO001 = share.NewMCode("IURSO-001", "Service options no services CSV branch")
	IURSO002 = share.NewMCode("IURSO-002", "Service options services CSV branch")
	IURSO003 = share.NewMCode("IURSO-003", "Service options read services CSV failed")
	IURSO004 = share.NewMCode("IURSO-004", "Service options service abbreviation branch")
	IURSO005 = share.NewMCode("IURSO-005", "Service options legend validation failed")
)

func (rcvr *xaligoUsecase) serviceOptions(opts entity.RenderOptions) ([]entity.ServiceEntry, map[int]string, error) {
	abbreviations := make(map[int]string, len(opts.Abbreviations))
	for id, value := range opts.Abbreviations {
		abbreviations[id] = value
	}
	if len(bytes.TrimSpace(opts.ServicesCSV)) == 0 {
		logger.DEBUG(IURSO001, "branch no services csv", map[string]any{"abbreviations": len(abbreviations)})
		return nil, abbreviations, nil
	}
	logger.DEBUG(IURSO002, "branch services csv", map[string]any{"bytes": len(opts.ServicesCSV)})
	if err := validateLegendCSVRows(opts.ServicesCSV); err != nil {
		logger.ERROR(IURSO005, "legend validation failed", map[string]any{"error": err})
		return nil, nil, err
	}
	entries, err := rcvr.xaligoRepository.ReadServiceListFromReader(bytes.NewReader(opts.ServicesCSV))
	if err != nil {
		logger.ERROR(IURSO003, "read services csv failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("read services CSV: %w", err)
	}
	if err := validateLegendEntries(entries); err != nil {
		logger.ERROR(IURSO005, "legend validation failed", map[string]any{"error": err})
		return nil, nil, err
	}
	for _, entry := range entries {
		if entry.CatalogID > 0 && entry.Abbreviation != "" {
			logger.DEBUG(IURSO004, "branch service abbreviation", map[string]any{"catalogID": entry.CatalogID})
			abbreviations[entry.CatalogID] = entry.Abbreviation
		}
	}
	return entries, abbreviations, nil
}

func validateLegendCSVRows(data []byte) error {
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		row := i + 1
		parts := strings.SplitN(line, ",", 3)
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		if len(parts) < 2 {
			return fmt.Errorf("legend row %d: catalog ID and official name are required", row)
		}
		if parts[0] == "" {
			return fmt.Errorf("legend row %d: catalog ID is required", row)
		}
		if parts[1] == "" {
			return fmt.Errorf("legend row %d: official name is required", row)
		}
	}
	return nil
}

func validateLegendEntries(entries []entity.ServiceEntry) error {
	seen := map[int]string{}
	for i, entry := range entries {
		row := i + 1
		if entry.CatalogID <= 0 {
			return fmt.Errorf("legend row %d: catalog ID is required", row)
		}
		if strings.TrimSpace(entry.OfficialName) == "" {
			return fmt.Errorf("legend row %d: official name is required", row)
		}
		if previous, ok := seen[entry.CatalogID]; ok {
			return fmt.Errorf("legend row %d: duplicate catalog ID %d already used by %q", row, entry.CatalogID, previous)
		}
		seen[entry.CatalogID] = entry.OfficialName
	}
	return nil
}
