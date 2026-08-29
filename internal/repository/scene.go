package repository

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strings"

	"github.com/xaligo/xaligo/internal/share"
)

const svgDataURLPrefix = "data:image/svg+xml;base64,"

var (
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

// SceneRepository provides the filesystem-backed asset operations required by
// the transitional V1 presentation-scene builder. The scene is an internal
// intermediate for SVG and PPTX; this repository does not expose a scene-file
// output format.
type SceneRepository interface {
	SvgToDataURL(path string) (string, error)
	SvgToDataURLFS(fsys fs.FS, path string) (string, error)
	FileID(name string) string
	SVGBGColor(dataURL string) string
}

type sceneRepository struct{}

func NewSceneRepository() SceneRepository {
	return &sceneRepository{}
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
func (rcvr *sceneRepository) SvgToDataURL(path string) (string, error) {
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
func (rcvr *sceneRepository) SvgToDataURLFS(fsys fs.FS, path string) (string, error) {
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
func (rcvr *sceneRepository) FileID(name string) string {
	sum := md5.Sum([]byte(name))
	return fmt.Sprintf("%x", sum)[:16]
}

var svgFillRe = regexp.MustCompile(`(?i)fill[=:]["']?(#[0-9a-fA-F]{3,8}|[a-zA-Z]+)`)

// SVGBGColor extracts the dominant fill colour from an SVG data URL.
// Falls back to "transparent" if none is found.
func (rcvr *sceneRepository) SVGBGColor(dataURL string) string {
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
