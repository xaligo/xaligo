package repository

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strings"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
)

const svgDataURLPrefix = "data:image/svg+xml;base64,"

var (
	IRSRS001     = share.NewMCode("IRSRS-001", "Read scene read failed")
	IRSRS002     = share.NewMCode("IRSRS-002", "Read scene parse failed")
	IRSRS003     = share.NewMCode("IRSRS-003", "Read scene initialize files branch")
	IRSRS004     = share.NewMCode("IRSRS-004", "Read scene initialize elements branch")
	IRSWSC001    = share.NewMCode("IRSWSC-001", "Write scene marshal failed")
	IRSWSC002    = share.NewMCode("IRSWSC-002", "Write scene write failed")
	IRISDU001    = share.NewMCode("IRISDU-001", "SVG data URL empty or data branch")
	IRISTDU001   = share.NewMCode("IRISTDU-001", "SVG to data URL read failed")
	IRISTDUFS001 = share.NewMCode("IRISTDUFS-001", "SVG to data URL FS open failed")
	IRISTDUFS002 = share.NewMCode("IRISTDUFS-002", "SVG to data URL FS read failed")
	IRISBGC001   = share.NewMCode("IRISBGC-001", "SVG background color decode branch")
	IRISBGC002   = share.NewMCode("IRISBGC-002", "SVG background color transparent fallback branch")
	IRISBGC003   = share.NewMCode("IRISBGC-003", "SVG background color candidate branch")
	IRISBFC001   = share.NewMCode("IRISBFC-001", "SVG background fill rejected branch")
	IRISBFC002   = share.NewMCode("IRISBFC-002", "SVG background fill accepted branch")
)

type ExcalidrawRepository interface {
	ReadScene(path string) (*entity.Scene, error)
	WriteScene(scene *entity.Scene, path string) error
	SvgToDataURL(path string) (string, error)
	SvgToDataURLFS(fsys fs.FS, path string) (string, error)
	FileID(name string) string
	SVGBGColor(dataURL string) string
}

type excalidrawRepository struct{}

func NewExcalidrawRepository() ExcalidrawRepository {
	return &excalidrawRepository{}
}

// ReadScene reads a .excalidraw JSON file and returns the parsed Scene.
func (rcvr *excalidrawRepository) ReadScene(path string) (*entity.Scene, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.ERROR(IRSRS001, "read failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("read scene %s: %w", path, err)
	}
	var scene entity.Scene
	if err := json.Unmarshal(data, &scene); err != nil {
		logger.ERROR(IRSRS002, "parse failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("parse scene %s: %w", path, err)
	}
	if scene.Files == nil {
		logger.DEBUG(IRSRS003, "branch initialize files", map[string]any{"path": path})
		scene.Files = map[string]map[string]interface{}{}
	}
	if scene.Elements == nil {
		logger.DEBUG(IRSRS004, "branch initialize elements", map[string]any{"path": path})
		scene.Elements = []map[string]interface{}{}
	}
	return &scene, nil
}

// WriteScene serialises the Scene and writes it back to path.
func (rcvr *excalidrawRepository) WriteScene(scene *entity.Scene, path string) error {
	data, err := json.MarshalIndent(scene, "", "  ")
	if err != nil {
		logger.ERROR(IRSWSC001, "marshal failed", map[string]any{"error": err})
		return fmt.Errorf("marshal scene: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		logger.ERROR(IRSWSC002, "write failed", map[string]any{"path": path, "error": err})
		return fmt.Errorf("write scene %s: %w", path, err)
	}
	return nil
}

func svgDataURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" || strings.HasPrefix(raw, "data:") {
		logger.DEBUG(IRISDU001, "branch empty or data URL")
		return raw
	}
	return svgDataURLPrefix + raw
}

// SvgToDataURL reads an SVG file and returns it as a base64 data URL.
func (rcvr *excalidrawRepository) SvgToDataURL(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.ERROR(IRISTDU001, "read failed", map[string]any{"path": path, "error": err})
		return "", fmt.Errorf("read SVG %s: %w", path, err)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return svgDataURL(encoded), nil
}

// SvgToDataURLFS is the fs.FS-aware variant of SvgToDataURL.
// It reads the SVG file from fsys instead of the OS filesystem.
func (rcvr *excalidrawRepository) SvgToDataURLFS(fsys fs.FS, path string) (string, error) {
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

// FileID derives a stable 16-char hex ID from a file path (MD5-based).
func (rcvr *excalidrawRepository) FileID(name string) string {
	sum := md5.Sum([]byte(name))
	return fmt.Sprintf("%x", sum)[:16]
}

var svgFillRe = regexp.MustCompile(`(?i)fill[=:]["']?(#[0-9a-fA-F]{3,8}|[a-zA-Z]+)`)

// SVGBGColor extracts the dominant fill colour from an SVG data URL.
// Falls back to "transparent" if none is found.
func (rcvr *excalidrawRepository) SVGBGColor(dataURL string) string {
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
