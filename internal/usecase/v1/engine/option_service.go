package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func ValidateLegendCSVRowsV1EngineOptionService(data []byte) error {
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

func ResolveServiceOptionsV1EngineOptionService(entries []entity.ServiceEntry, base map[int]string) (map[int]string, error) {
	abbreviations := make(map[int]string, len(base)+len(entries))
	for id, value := range base {
		abbreviations[id] = value
	}
	if err := validateLegendEntriesV1EngineOptionService(entries); err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.CatalogID > 0 && entry.Abbreviation != "" {
			abbreviations[entry.CatalogID] = entry.Abbreviation
		}
	}
	return abbreviations, nil
}

func validateLegendEntriesV1EngineOptionService(entries []entity.ServiceEntry) error {
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
