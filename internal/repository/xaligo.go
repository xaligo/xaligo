package repository

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
)

var (
	IRXRS001     = share.NewMCode("IRXRS-001", "Read xaligo source failed")
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
	IRILFC001    = share.NewMCode("IRILFC-001", "Load from CSV open failed")
	IRILFC002    = share.NewMCode("IRILFC-002", "Load from CSV EOF branch")
	IRILFC003    = share.NewMCode("IRILFC-003", "Load from CSV read row failed branch")
	IRILFC004    = share.NewMCode("IRILFC-004", "Load from CSV short row branch")
	IRILFC005    = share.NewMCode("IRILFC-005", "Load from CSV match branch")
	IRILFC006    = share.NewMCode("IRILFC-006", "Load from CSV empty base64 branch")
	IRILFC007    = share.NewMCode("IRILFC-007", "Load from CSV not found branch")
	IRILFCBI001  = share.NewMCode("IRILFCBI-001", "Load from CSV by ID open failed")
	IRILFCBI002  = share.NewMCode("IRILFCBI-002", "Load from CSV by ID EOF branch")
	IRILFCBI003  = share.NewMCode("IRILFCBI-003", "Load from CSV by ID skip row branch")
	IRILFCBI004  = share.NewMCode("IRILFCBI-004", "Load from CSV by ID name mismatch branch")
	IRILFCBI005  = share.NewMCode("IRILFCBI-005", "Load from CSV by ID empty base64 branch")
	IRILFCBI006  = share.NewMCode("IRILFCBI-006", "Load from CSV by ID found branch")
	IRILFCBI007  = share.NewMCode("IRILFCBI-007", "Load from CSV by ID not found branch")
	IRILCBI001   = share.NewMCode("IRILCBI-001", "Lookup catalog by ID open failed")
	IRILCBI002   = share.NewMCode("IRILCBI-002", "Lookup catalog by ID EOF branch")
	IRILCBI003   = share.NewMCode("IRILCBI-003", "Lookup catalog by ID skip row branch")
	IRILCBI004   = share.NewMCode("IRILCBI-004", "Lookup catalog by ID found branch")
	IRILCBI005   = share.NewMCode("IRILCBI-005", "Lookup catalog by ID not found branch")
	IRILCBIFS001 = share.NewMCode("IRILCBIFS-001", "Lookup catalog by ID FS open failed")
	IRILCBIFS002 = share.NewMCode("IRILCBIFS-002", "Lookup catalog by ID FS EOF branch")
	IRILCBIFS003 = share.NewMCode("IRILCBIFS-003", "Lookup catalog by ID FS skip row branch")
	IRILCBIFS004 = share.NewMCode("IRILCBIFS-004", "Lookup catalog by ID FS found branch")
	IRILCBIFS005 = share.NewMCode("IRILCBIFS-005", "Lookup catalog by ID FS not found branch")
)

type XaligoRepository interface {
	ReadSource(path string) ([]byte, error)
	ReadServiceList(path string) ([]entity.ServiceEntry, error)
	ReadServiceListFromReader(r io.Reader) ([]entity.ServiceEntry, error)
	LoadFromCSV(csvPath, svgFilename string) (string, error)
	LoadFromCSVByID(csvPath string, id int, name string) (string, string, error)
	LookupCatalogByID(csvPath string, id int) (entity.CatalogEntry, error)
	LookupCatalogByIDFS(fsys fs.FS, csvPath string, id int) (entity.CatalogEntry, error)
}

type xaligoRepository struct{}

func NewXaligoRepository() XaligoRepository {
	return &xaligoRepository{}
}

func (rcvr *xaligoRepository) ReadSource(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.ERROR(IRXRS001, "read failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("read xaligo source %s: %w", path, err)
	}
	return data, nil
}

// ReadServiceList reads a CSV/TXT service list from the given file path.
func (rcvr *xaligoRepository) ReadServiceList(path string) ([]entity.ServiceEntry, error) {
	f, err := os.Open(path)
	if err != nil {
		logger.ERROR(IRSLRSL001, "open failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("open service list %s: %w", path, err)
	}
	defer f.Close()
	entries, err := rcvr.ReadServiceListFromReader(f)
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
func (rcvr *xaligoRepository) ReadServiceListFromReader(r io.Reader) ([]entity.ServiceEntry, error) {
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

func newCatalogCSVReader(r io.Reader) *csv.Reader {
	reader := csv.NewReader(r)
	reader.Comment = '#'
	reader.FieldsPerRecord = -1
	return reader
}

// LoadFromCSV loads the base64 data URL for svgFilename from the service catalog CSV.
// The CSV is expected to have columns: id,category,service,svg_file,rel_path,base64
func (rcvr *xaligoRepository) LoadFromCSV(csvPath, svgFilename string) (string, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		logger.ERROR(IRILFC001, "open failed", map[string]any{"csvPath": csvPath, "error": err})
		return "", fmt.Errorf("open catalog CSV %s: %w", csvPath, err)
	}
	defer f.Close()

	r := newCatalogCSVReader(f)

	for {
		rec, err := r.Read()
		if err == io.EOF {
			logger.DEBUG(IRILFC002, "branch EOF")
			break
		}
		if err != nil {
			logger.WARN(IRILFC003, "branch read row failed", map[string]any{"error": err})
			continue
		}
		if len(rec) < 6 {
			logger.WARN(IRILFC004, "branch short row", map[string]any{"columns": len(rec)})
			continue
		}
		if strings.TrimSpace(rec[3]) == svgFilename {
			logger.DEBUG(IRILFC005, "branch match", map[string]any{"svgFilename": svgFilename})
			raw := strings.TrimSpace(rec[5])
			if raw == "" {
				logger.ERROR(IRILFC006, "branch empty base64", map[string]any{"svgFilename": svgFilename})
				return "", fmt.Errorf("empty base64 for %s", svgFilename)
			}
			return svgDataURL(raw), nil
		}
	}
	logger.ERROR(IRILFC007, "branch not found", map[string]any{"svgFilename": svgFilename})
	return "", fmt.Errorf("SVG %q not found in catalog CSV", svgFilename)
}

// LoadFromCSVByID loads the SVG path and data URL by catalog ID (and optionally name).
// Returns (svgRelPath, dataURL, error).
func (rcvr *xaligoRepository) LoadFromCSVByID(csvPath string, id int, name string) (string, string, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		logger.ERROR(IRILFCBI001, "open failed", map[string]any{"csvPath": csvPath, "error": err})
		return "", "", fmt.Errorf("open catalog CSV %s: %w", csvPath, err)
	}
	defer f.Close()

	r := newCatalogCSVReader(f)

	for {
		rec, err := r.Read()
		if err == io.EOF {
			logger.DEBUG(IRILFCBI002, "branch EOF")
			break
		}
		if err != nil {
			logger.WARN(IRILFCBI003, "branch skip row", map[string]any{"error": err})
			continue
		}
		if len(rec) < 6 {
			logger.WARN(IRILFCBI003, "branch skip row", map[string]any{"columns": len(rec)})
			continue
		}
		rowID, err := strconv.Atoi(strings.TrimSpace(rec[0]))
		if err != nil || rowID != id {
			logger.DEBUG(IRILFCBI003, "branch skip row", map[string]any{"id": id})
			continue
		}
		if name != "" {
			svc := strings.TrimSpace(rec[2])
			if !strings.EqualFold(svc, name) &&
				!strings.Contains(strings.ToLower(svc), strings.ToLower(name)) &&
				!strings.Contains(strings.ToLower(name), strings.ToLower(svc)) {
				logger.DEBUG(IRILFCBI004, "branch name mismatch", map[string]any{"name": name, "service": svc})
				continue
			}
		}
		relPath := strings.TrimSpace(rec[4])
		raw := strings.TrimSpace(rec[5])
		if raw == "" {
			logger.ERROR(IRILFCBI005, "branch empty base64", map[string]any{"id": id})
			return relPath, "", fmt.Errorf("empty base64 for catalog ID %d", id)
		}
		logger.DEBUG(IRILFCBI006, "branch found", map[string]any{"id": id})
		return relPath, svgDataURL(raw), nil
	}
	logger.ERROR(IRILFCBI007, "branch not found", map[string]any{"id": id})
	return "", "", fmt.Errorf("catalog ID %d not found in CSV", id)
}

// LookupCatalogByID finds the first entry with the given ID in service-catalog.csv
// and returns its data including the pre-encoded SVG data URL.
func (rcvr *xaligoRepository) LookupCatalogByID(csvPath string, id int) (entity.CatalogEntry, error) {
	f, err := os.Open(csvPath)
	if err != nil {
		logger.ERROR(IRILCBI001, "open failed", map[string]any{"csvPath": csvPath, "error": err})
		return entity.CatalogEntry{}, fmt.Errorf("open catalog CSV %s: %w", csvPath, err)
	}
	defer f.Close()
	r := newCatalogCSVReader(f)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			logger.DEBUG(IRILCBI002, "branch EOF")
			break
		}
		if err != nil {
			logger.WARN(IRILCBI003, "branch skip row", map[string]any{"error": err})
			continue
		}
		if len(rec) < 6 {
			logger.WARN(IRILCBI003, "branch skip row", map[string]any{"columns": len(rec)})
			continue
		}
		rowID, err := strconv.Atoi(strings.TrimSpace(rec[0]))
		if err != nil || rowID != id {
			logger.DEBUG(IRILCBI003, "branch skip row", map[string]any{"id": id})
			continue
		}
		raw := strings.TrimSpace(rec[5])
		logger.DEBUG(IRILCBI004, "branch found", map[string]any{"id": id})
		return entity.CatalogEntry{
			ID:       rowID,
			Category: strings.TrimSpace(rec[1]),
			Service:  strings.TrimSpace(rec[2]),
			SVGFile:  strings.TrimSpace(rec[3]),
			RelPath:  strings.TrimSpace(rec[4]),
			DataURL:  svgDataURL(raw),
		}, nil
	}
	logger.ERROR(IRILCBI005, "branch not found", map[string]any{"id": id})
	return entity.CatalogEntry{}, fmt.Errorf("catalog ID %d not found", id)
}

// LookupCatalogByIDFS is the fs.FS-aware variant of LookupCatalogByID.
func (rcvr *xaligoRepository) LookupCatalogByIDFS(fsys fs.FS, csvPath string, id int) (entity.CatalogEntry, error) {
	f, err := fsys.Open(csvPath)
	if err != nil {
		logger.ERROR(IRILCBIFS001, "open failed", map[string]any{"csvPath": csvPath, "error": err})
		return entity.CatalogEntry{}, fmt.Errorf("open catalog CSV %s: %w", csvPath, err)
	}
	defer f.Close()
	r := newCatalogCSVReader(f)
	for {
		rec, err := r.Read()
		if err == io.EOF {
			logger.DEBUG(IRILCBIFS002, "branch EOF")
			break
		}
		if err != nil {
			logger.WARN(IRILCBIFS003, "branch skip row", map[string]any{"error": err})
			continue
		}
		if len(rec) < 6 {
			logger.WARN(IRILCBIFS003, "branch skip row", map[string]any{"columns": len(rec)})
			continue
		}
		rowID, err := strconv.Atoi(strings.TrimSpace(rec[0]))
		if err != nil || rowID != id {
			logger.DEBUG(IRILCBIFS003, "branch skip row", map[string]any{"id": id})
			continue
		}
		raw := strings.TrimSpace(rec[5])
		logger.DEBUG(IRILCBIFS004, "branch found", map[string]any{"id": id})
		return entity.CatalogEntry{
			ID:       rowID,
			Category: strings.TrimSpace(rec[1]),
			Service:  strings.TrimSpace(rec[2]),
			SVGFile:  strings.TrimSpace(rec[3]),
			RelPath:  strings.TrimSpace(rec[4]),
			DataURL:  svgDataURL(raw),
		}, nil
	}
	logger.ERROR(IRILCBIFS005, "branch not found", map[string]any{"id": id})
	return entity.CatalogEntry{}, fmt.Errorf("catalog ID %d not found", id)
}
