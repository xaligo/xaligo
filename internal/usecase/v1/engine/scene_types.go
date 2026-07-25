package engine

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"hash/fnv"
	"io/fs"
	"regexp"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

const excalidrawUpdatedV1EngineSceneTypes int64 = 1709000000000

// stableSceneSeedV1EngineSceneTypes derives Excalidraw metadata from a semantic
// element identity. Generated scenes therefore do not depend on wall-clock time
// or process-local pseudo-random state.
func stableSceneSeedV1EngineSceneTypes(parts ...string) int {
	hash := fnv.New32a()
	for _, part := range parts {
		_, _ = hash.Write([]byte{0})
		_, _ = hash.Write([]byte(part))
	}
	return int(hash.Sum32()%99999998) + 1
}

type fileV1EngineSceneTypes struct {
	Type     string           `json:"type"`
	Version  int              `json:"version"`
	Source   string           `json:"source"`
	Elements []map[string]any `json:"elements"`
	AppState map[string]any   `json:"appState"`
	Files    map[string]any   `json:"files"`
}

// groupDef holds visual style for an AWS architecture group tag.
type groupDefV1EngineSceneTypes struct {
	StrokeColor string
	StrokeStyle string
	StrokeWidth int
	IconFile    string // filename inside Architecture-Group-Icons dir, empty = no icon
}

// awsGroups maps xal tag names to their AWS group visual definitions.
var awsGroupsV1EngineSceneTypes = map[string]groupDefV1EngineSceneTypes{
	"aws-cloud":                     {"#000000", "solid", 2, "AWS-Cloud-logo_32.svg"},
	"aws-cloud-alt":                 {"#000000", "solid", 2, "AWS-Cloud_32.svg"},
	"region":                        {"#00A1C9", "dashed", 2, "Region_32.svg"},
	"availability-zone":             {"#00A1C9", "dashed", 2, ""},
	"security-group":                {"#CC0000", "dashed", 2, ""},
	"auto-scaling-group":            {"#E7601B", "dashed", 2, "Auto-Scaling-group_32.svg"},
	"vpc":                           {"#8C4FFF", "solid", 2, "Virtual-private-cloud-VPC_32.svg"},
	"private-subnet":                {"#00A1C9", "solid", 2, "Private-subnet_32.svg"},
	"public-subnet":                 {"#3F8624", "solid", 2, "Public-subnet_32.svg"},
	"server-contents":               {"#7A7C7F", "solid", 2, "Server-contents_32.svg"},
	"corporate-data-center":         {"#7A7C7F", "solid", 2, "Corporate-data-center_32.svg"},
	"ec2-instance-contents":         {"#E7601B", "solid", 2, "EC2-instance-contents_32.svg"},
	"spot-fleet":                    {"#E7601B", "solid", 2, "Spot-Fleet_32.svg"},
	"aws-account":                   {"#E7008A", "solid", 2, "AWS-Account_32.svg"},
	"aws-iot-greengrass-deployment": {"#3F8624", "solid", 2, "AWS-IoT-Greengrass-Deployment_32.svg"},
	"aws-iot-greengrass":            {"#3F8624", "solid", 2, ""},
	"elastic-beanstalk-container":   {"#E7601B", "solid", 2, ""},
	"aws-step-functions-workflow":   {"#E7008A", "solid", 2, ""},
	"generic-group":                 {"#AAB7B8", "dashed", 1, ""},
	"capture":                       {"#F5A623", "dashed", 1, ""},
}

const (
	groupIconSizeV1EngineSceneTypes           = 32
	groupHeaderLeftOverflowV1EngineSceneTypes = 2
	groupHeaderTextInsetV1EngineSceneTypes    = 4
	groupHeaderPadEndV1EngineSceneTypes       = 18
	groupHeaderTipMaxV1EngineSceneTypes       = 14
	groupHeaderBorderGapV1EngineSceneTypes    = 4
	groupFontSizeV1EngineSceneTypes           = 14
	groupTextHeightV1EngineSceneTypes         = groupFontSizeV1EngineSceneTypes + 4
	groupHeaderTextPadYV1EngineSceneTypes     = 1
	groupFontFamilyV1EngineSceneTypes         = 2 // Helvetica (normal)
	groupLabelCharWV1EngineSceneTypes         = 9.6
	itemFallbackIconColorV1EngineSceneTypes   = "#7758C1"
)

var svgTintColorREV1EngineSceneTypes = regexp.MustCompile(`(?i)#[0-9a-f]{3,8}|currentColor`)

var svgCurrentColorREV1EngineSceneTypes = regexp.MustCompile(`(?i)currentColor`)

var (
	IUESW002V1EngineSceneTypes   = share.NewMCode("IUESW-002", "Walk generic group icon lookup warning")
	IUESRIA001V1EngineSceneTypes = share.NewMCode("IUESRIA-001", "Render icon at invalid item ID warning")
	IUESRIA002V1EngineSceneTypes = share.NewMCode("IUESRIA-002", "Render icon at catalog lookup warning")
	IUESRIA003V1EngineSceneTypes = share.NewMCode("IUESRIA-003", "Render icon at load SVG warning")
	IUESRC001V1EngineSceneTypes  = share.NewMCode("IUESRC-001", "Render connections invalid source or destination warning")
	IUESRC002V1EngineSceneTypes  = share.NewMCode("IUESRC-002", "Render connections source item not rendered warning")
	IUESRC003V1EngineSceneTypes  = share.NewMCode("IUESRC-003", "Render connections destination item not rendered warning")
	IUESRC004V1EngineSceneTypes  = share.NewMCode("IUESRC-004", "Render cross-frame connection frame not rendered warning")
)

// tintSVGDataURL makes a group header icon use the same semantic colour as
// its group border and title. White and transparent portions are preserved.
func tintSVGDataURLV1EngineSceneTypes(dataURL, color string) string {
	if !strings.HasPrefix(dataURL, share.SVGDataURLPrefix) {
		return dataURL
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(dataURL, share.SVGDataURLPrefix))
	if err != nil {
		return dataURL
	}
	tinted := svgTintColorREV1EngineSceneTypes.ReplaceAllStringFunc(string(raw), func(found string) string {
		switch strings.ToLower(found) {
		case "#fff", "#ffffff", "#ffffffff":
			return found
		default:
			return color
		}
	})
	return share.SVGDataURLFromBytes([]byte(tinted))
}

// normalizeItemSVGDataURL resolves SVG currentColor for item icons. PowerPoint
// does not reliably resolve currentColor when the WASM exporter embeds SVG
// directly, so Tabler-style line icons can appear blank unless they carry an
// explicit stroke color.
func normalizeItemSVGDataURLV1EngineSceneTypes(dataURL string) string {
	if !strings.HasPrefix(dataURL, share.SVGDataURLPrefix) {
		return dataURL
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(dataURL, share.SVGDataURLPrefix))
	if err != nil {
		return dataURL
	}
	normalized := svgCurrentColorREV1EngineSceneTypes.ReplaceAllString(string(raw), itemFallbackIconColorV1EngineSceneTypes)
	return share.SVGDataURLFromBytes([]byte(normalized))
}

// staggerFills are background fill colors for staggered AZ layers.
// Index = StaggerDepth (0 = front/white, 1/2 = progressively darker teal).
var staggerFillsV1EngineSceneTypes = []string{"#ffffff", "#c8e8e8", "#92cecd"}

// staggerBGColor returns the appropriate backgroundColor for a box.
// Boxes that participate in a staggered group get a solid fill so that
// overlapping back-layers are visually distinct.
func staggerBGColorV1EngineSceneTypes(b *entity.Box) string {
	if !b.InStagger {
		return "transparent"
	}
	idx := b.StaggerDepth
	if idx >= len(staggerFillsV1EngineSceneTypes) {
		idx = len(staggerFillsV1EngineSceneTypes) - 1
	}
	return staggerFillsV1EngineSceneTypes[idx]
}

const (
	itemMaxSizeV1EngineSceneTypes = 32.0
	// Item grids may shrink icons to preserve containment in compact groups;
	// labels use the shared shrink/clip text contract in the remaining cell.
	itemMinSizeV1EngineSceneTypes     = 8.0
	itemLabelFontPtV1EngineSceneTypes = 8.0
	itemLabelFontPxV1EngineSceneTypes = itemLabelFontPtV1EngineSceneTypes * 96.0 / 72.0
	itemLabelHV1EngineSceneTypes      = 14.0
	itemLabelWV1EngineSceneTypes      = 56.0 // text box width for item labels (wider than icon, centred on icon)
	itemLabelCharWV1EngineSceneTypes  = 6.2
	itemGapV1EngineSceneTypes         = 8.0
	// Mirrors the shared plan's visual anchor-grid expansion so groups reserve enough
	// top clearance before PPTX adds the grid around each item.
	itemAnchorGridVisualPadPxV1EngineSceneTypes = 6.0
	excalidrawAnchorGridV1EngineSceneTypes      = 5
	excalidrawAnchorPadPxV1EngineSceneTypes     = 2.0
	excalidrawAnchorCellGapPxV1EngineSceneTypes = 1.0
)

// paperSizeNames maps (short-side, long-side) → paper name for reverse lookup.
var paperSizeNamesV1EngineSceneTypes = map[[2]int]string{
	{559, 794}:   "A5",
	{794, 1122}:  "A4",
	{1122, 1587}: "A3",
	{1587, 2245}: "A2",
	{2245, 3179}: "A1",
	{816, 1056}:  "Letter",
	{816, 1344}:  "Legal",
	{1056, 1632}: "Tabloid",
}

// detectPaperName returns e.g. "A4 landscape" / "A4 portrait" from box dimensions.
func detectPaperNameV1EngineSceneTypes(w, h float64) string {
	wi, hi := int(w), int(h)
	short, long := wi, hi
	orientation := "portrait"
	if wi > hi {
		short, long = hi, wi
		orientation = "landscape"
	}
	if name, ok := paperSizeNamesV1EngineSceneTypes[[2]int{short, long}]; ok {
		return name + " " + orientation
	}
	return fmt.Sprintf("%d×%d", wi, hi)
}

// SceneDependenciesV1EngineSceneTypes defines the synchronous asset operations needed while a
// scene is calculated. The use-case orchestration layer adapts repositories to
// these function ports; the V1 engine neither owns repository implementations
// nor controls their execution model.
type SceneDependenciesV1EngineSceneTypes struct {
	LookupCatalogByID   func(csvPath string, id int) (entity.CatalogEntry, error)
	LookupCatalogByIDFS func(fsys fs.FS, csvPath string, id int) (entity.CatalogEntry, error)
	SVGToDataURL        func(path string) (string, error)
	SVGToDataURLFS      func(fsys fs.FS, path string) (string, error)
	FileID              func(name string) string
	SVGBGColor          func(dataURL string) string
}

func (d SceneDependenciesV1EngineSceneTypes) lookupCatalogByIDV1EngineSceneTypes(fsys fs.FS, csvPath string, id int) (entity.CatalogEntry, error) {
	if fsys != nil {
		if d.LookupCatalogByIDFS == nil {
			return entity.CatalogEntry{}, fmt.Errorf("embedded catalog lookup dependency is not configured")
		}
		return d.LookupCatalogByIDFS(fsys, csvPath, id)
	}
	if d.LookupCatalogByID == nil {
		return entity.CatalogEntry{}, fmt.Errorf("catalog lookup dependency is not configured")
	}
	return d.LookupCatalogByID(csvPath, id)
}

func (d SceneDependenciesV1EngineSceneTypes) svgToDataURLV1EngineSceneTypes(fsys fs.FS, path string) (string, error) {
	if fsys != nil {
		if d.SVGToDataURLFS == nil {
			return "", fmt.Errorf("embedded SVG conversion dependency is not configured")
		}
		return d.SVGToDataURLFS(fsys, path)
	}
	if d.SVGToDataURL == nil {
		return "", fmt.Errorf("SVG conversion dependency is not configured")
	}
	return d.SVGToDataURL(path)
}

func (d SceneDependenciesV1EngineSceneTypes) fileIDV1EngineSceneTypes(name string) string {
	if d.FileID != nil {
		return d.FileID(name)
	}
	// Match the current adapter's stable ID so an omitted optional port cannot
	// produce an invalid or non-deterministic Excalidraw file reference.
	sum := md5.Sum([]byte(name))
	return fmt.Sprintf("%x", sum)[:16]
}

func (d SceneDependenciesV1EngineSceneTypes) svgBGColorV1EngineSceneTypes(dataURL string) string {
	if d.SVGBGColor == nil {
		return "transparent"
	}
	return d.SVGBGColor(dataURL)
}

type resolvedConnectionStyleV1EngineSceneTypes struct {
	Kind                     string
	Color                    string
	Width                    float64
	WidthExplicit            bool
	StrokeStyle              string
	StartArrowhead           string
	StartArrowheadExplicit   bool
	EndArrowhead             string
	EndArrowheadExplicit     bool
	ExcalidrawStartArrowhead any
	ExcalidrawEndArrowhead   any
}
