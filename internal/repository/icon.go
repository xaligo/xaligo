package repository

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
)

const svgDataURLPrefix = "data:image/svg+xml;base64,"

var (
	IRISDU001    = share.NewMCode("IRISDU-001", "SVG data URL empty or data branch")
	IRISTDU001   = share.NewMCode("IRISTDU-001", "SVG to data URL read failed")
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
	IRISTDUFS001 = share.NewMCode("IRISTDUFS-001", "SVG to data URL FS open failed")
	IRISTDUFS002 = share.NewMCode("IRISTDUFS-002", "SVG to data URL FS read failed")
	IRISBGC001   = share.NewMCode("IRISBGC-001", "SVG background color decode branch")
	IRISBGC002   = share.NewMCode("IRISBGC-002", "SVG background color transparent fallback branch")
	IRISBGC003   = share.NewMCode("IRISBGC-003", "SVG background color candidate branch")
	IRISBFC001   = share.NewMCode("IRISBFC-001", "SVG background fill rejected branch")
	IRISBFC002   = share.NewMCode("IRISBFC-002", "SVG background fill accepted branch")
)

func svgDataURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" || strings.HasPrefix(raw, "data:") {
		logger.DEBUG(IRISDU001, "branch empty or data URL")
		return raw
	}
	return svgDataURLPrefix + raw
}

// SvgToDataURL reads an SVG file and returns it as a base64 data URL.
func SvgToDataURL(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.ERROR(IRISTDU001, "read failed", map[string]any{"path": path, "error": err})
		return "", fmt.Errorf("read SVG %s: %w", path, err)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return svgDataURL(encoded), nil
}

// FileID derives a stable 16-char hex ID from a file path (MD5-based).
func FileID(name string) string {
	sum := md5.Sum([]byte(name))
	return fmt.Sprintf("%x", sum)[:16]
}

// LoadFromCSV loads the base64 data URL for svgFilename from the service catalog CSV.
// The CSV is expected to have columns: id,category,service,svg_file,rel_path,base64
func LoadFromCSV(csvPath, svgFilename string) (string, error) {
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
		// columns: id, category, service, svg_file, rel_path, base64
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
func LoadFromCSVByID(csvPath string, id int, name string) (string, string, error) {
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
		// Optional: verify name matches service column (case-insensitive)
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
func LookupCatalogByID(csvPath string, id int) (entity.CatalogEntry, error) {
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
// It opens csvPath inside fsys instead of the OS filesystem.
// Use this in contexts where assets are embedded (e.g. the WASM build).
func LookupCatalogByIDFS(fsys fs.FS, csvPath string, id int) (entity.CatalogEntry, error) {
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

// SvgToDataURLFS is the fs.FS-aware variant of SvgToDataURL.
// It reads the SVG file from fsys instead of the OS filesystem.
func SvgToDataURLFS(fsys fs.FS, path string) (string, error) {
	f, err := fsys.Open(path)
	if err != nil {
		logger.ERROR(IRISTDUFS001, "open failed", map[string]any{"path": path, "error": err})
		return "", fmt.Errorf("open SVG %s: %w", path, err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		logger.ERROR(IRISTDUFS002, "read failed", map[string]any{"path": path, "error": err})
		return "", fmt.Errorf("read SVG %s: %w", path, err)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return svgDataURL(encoded), nil
}

var svgFillRe = regexp.MustCompile(`(?i)fill[=:]["']?(#[0-9a-fA-F]{3,8}|[a-zA-Z]+)`)

// SVGBGColor extracts the dominant fill colour from an SVG data URL.
// Falls back to "transparent" if none is found.
func SVGBGColor(dataURL string) string {
	var svgBytes []byte
	if strings.HasPrefix(dataURL, svgDataURLPrefix) {
		b64 := dataURL[len(svgDataURLPrefix):]
		decoded, err := base64.StdEncoding.DecodeString(b64)
		if err == nil {
			logger.DEBUG(IRISBGC001, "branch decoded")
			svgBytes = decoded
		}
	}
	if len(svgBytes) == 0 {
		logger.DEBUG(IRISBGC002, "branch transparent fallback")
		return "transparent"
	}
	matches := svgFillRe.FindAllSubmatch(svgBytes, -1)
	for _, m := range matches {
		color := strings.ToLower(strings.Trim(string(m[1]), `"'`))
		if svgBackgroundFillCandidate(color) {
			logger.DEBUG(IRISBGC003, "branch candidate", map[string]any{"color": color})
			return color
		}
	}
	logger.DEBUG(IRISBGC002, "branch transparent fallback")
	return "transparent"
}

func svgBackgroundFillCandidate(color string) bool {
	switch color {
	case "", "none", "transparent", "white", "#ffffff", "#fff", "#ffffffff":
		logger.DEBUG(IRISBFC001, "branch rejected", map[string]any{"color": color})
		return false
	case "#231815", "#6e6e6e":
		logger.DEBUG(IRISBFC001, "branch rejected", map[string]any{"color": color})
		return false
	default:
		logger.DEBUG(IRISBFC002, "branch accepted", map[string]any{"color": color})
		return true
	}
}
