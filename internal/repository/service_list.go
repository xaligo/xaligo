package repository

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
)

var (
	IRSLRSL001   = share.NewMCode("IRSLRSL-001", "Read service list open failed")
	IRSLRSL002   = share.NewMCode("IRSLRSL-002", "Read service list parse failed")
	IRSLRSLFR001 = share.NewMCode("IRSLRSLFR-001", "Read service list from reader skip empty or comment branch")
	IRSLRSLFR002 = share.NewMCode("IRSLRSLFR-002", "Read service list from reader single column branch")
	IRSLRSLFR003 = share.NewMCode("IRSLRSLFR-003", "Read service list from reader two column ID branch")
	IRSLRSLFR004 = share.NewMCode("IRSLRSLFR-004", "Read service list from reader two column name branch")
	IRSLRSLFR005 = share.NewMCode("IRSLRSLFR-005", "Read service list from reader multi column ID branch")
	IRSLRSLFR006 = share.NewMCode("IRSLRSLFR-006", "Read service list from reader multi column name branch")
	IRSLRSLFR007 = share.NewMCode("IRSLRSLFR-007", "Read service list from reader append branch")
	IRSLRSLFR008 = share.NewMCode("IRSLRSLFR-008", "Read service list from reader scan failed")
)

// ReadServiceList reads a CSV/TXT service list from the given file path.
func ReadServiceList(path string) ([]entity.ServiceEntry, error) {
	f, err := os.Open(path)
	if err != nil {
		logger.ERROR(IRSLRSL001, "open failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("open service list %s: %w", path, err)
	}
	defer f.Close()
	entries, err := ReadServiceListFromReader(f)
	if err != nil {
		logger.ERROR(IRSLRSL002, "parse failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("read service list %s: %w", path, err)
	}
	return entries, nil
}

// ReadServiceListFromReader parses service list CSV content from an io.Reader.
// This is the Reader-based variant used by the WASM build to avoid file I/O.
//
// Format support:
//   - Lines beginning with '#' are comments and skipped.
//   - Single-column: service name only
//   - Two-column:    id,service_name  OR  service_name,category
//   - Three-column+: id,service_name,category,...
func ReadServiceListFromReader(r io.Reader) ([]entity.ServiceEntry, error) {
	var entries []entity.ServiceEntry
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			logger.DEBUG(IRSLRSLFR001, "branch skip empty or comment")
			continue
		}
		// Split up to 7 columns: id,正式名称,略語,サービス概要,用途,備考
		parts := strings.SplitN(line, ",", 7)
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}

		var entry entity.ServiceEntry
		switch len(parts) {
		case 1:
			logger.DEBUG(IRSLRSLFR002, "branch single column")
			entry.OfficialName = parts[0]
		case 2:
			// Could be "id,name" or "name,category"
			if id, err := strconv.Atoi(parts[0]); err == nil {
				logger.DEBUG(IRSLRSLFR003, "branch two column ID", map[string]any{"catalogID": id})
				entry.CatalogID = id
				entry.OfficialName = parts[1]
			} else {
				logger.DEBUG(IRSLRSLFR004, "branch two column name")
				entry.OfficialName = parts[0]
			}
		default:
			// 3+ columns: id,正式名称,略語,...
			if id, err := strconv.Atoi(parts[0]); err == nil {
				logger.DEBUG(IRSLRSLFR005, "branch multi column ID", map[string]any{"catalogID": id})
				entry.CatalogID = id
				entry.OfficialName = parts[1]
				if len(parts) >= 3 {
					entry.Abbreviation = parts[2]
				}
			} else {
				logger.DEBUG(IRSLRSLFR006, "branch multi column name")
				entry.OfficialName = parts[0]
				if len(parts) >= 2 {
					entry.Abbreviation = parts[1]
				}
			}
		}

		if entry.OfficialName != "" {
			logger.DEBUG(IRSLRSLFR007, "branch append", map[string]any{"catalogID": entry.CatalogID, "officialName": entry.OfficialName})
			entries = append(entries, entry)
		}
	}
	if err := scanner.Err(); err != nil {
		logger.ERROR(IRSLRSLFR008, "scan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("scan service list: %w", err)
	}
	return entries, nil
}
