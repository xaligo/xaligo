package usecase

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/fs"
	"math"
	"math/rand"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/share"
)

var (
	IURBS001 = share.NewMCode("IURBS-001", "Build scene context check failed")
	IURBS002 = share.NewMCode("IURBS-002", "Build scene validate render options failed")
	IURBS003 = share.NewMCode("IURBS-003", "Build scene connection node branch")
	IURBS004 = share.NewMCode("IURBS-004", "Build scene embedded assets branch")
	IURBS005 = share.NewMCode("IURBS-005", "Build scene default embedded item size branch")
	IURBS006 = share.NewMCode("IURBS-006", "Build scene native assets branch")
	IURBS007 = share.NewMCode("IURBS-007", "Build scene parse DSL failed")
	IURBS008 = share.NewMCode("IURBS-008", "Build scene build layout failed")
	IURBS009 = share.NewMCode("IURBS-009", "Build scene service options failed")
	IURBS010 = share.NewMCode("IURBS-010", "Build scene build JSON failed")
	IURBS011 = share.NewMCode("IURBS-011", "Build scene apply theme failed")
)

func (rcvr *xaligoUsecase) buildScene(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, []entity.ServiceEntry, error) {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURBS001, "context check failed", map[string]any{"error": err})
		return nil, nil, err
	}
	if err := ValidateRenderOptions(opts); err != nil {
		logger.ERROR(IURBS002, "validate render options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	theme, _ := entity.NormalizeTheme(opts.Theme)
	doc, err := Parse(bytes.NewReader(input))
	if err != nil {
		logger.ERROR(IURBS007, "parse DSL failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("parse DSL: %w", err)
	}
	root, err := Build(doc)
	if err != nil {
		logger.ERROR(IURBS008, "build layout failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build layout: %w", err)
	}
	entries, abbreviations, err := rcvr.serviceOptions(opts)
	if err != nil {
		logger.ERROR(IURBS009, "service options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	connections := collectConnectionNodes(doc.Root)
	var scene []byte
	if opts.Assets != nil {
		logger.DEBUG(IURBS004, "branch embedded assets")
		itemSize := opts.Assets.ItemIconSize
		if itemSize <= 0 {
			logger.DEBUG(IURBS005, "branch default embedded item size")
			itemSize = 32
		}
		scene, err = BuildJSONWithFS(root, opts.Assets.FS, opts.Assets.CatalogCSV, opts.Assets.GroupIconsDir, itemSize, connections, abbreviations, rcvr.sceneDependencies())
	} else {
		logger.DEBUG(IURBS006, "branch native assets")
		cfg := config.New()
		scene, err = BuildJSON(root, filepath.Join(cfg.AssetDir_, "Architecture-Group-Icons"), cfg.SvcCatalogCSV, cfg.ProjectRoot, cfg.ItemIconSize, connections, abbreviations, nil, rcvr.sceneDependencies())
	}
	if err != nil {
		logger.ERROR(IURBS010, "build JSON failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build excalidraw JSON: %w", err)
	}
	scene, err = ApplyThemeJSON(scene, theme)
	if err != nil {
		logger.ERROR(IURBS011, "apply theme failed", map[string]any{"error": err})
	}
	return scene, entries, err
}

func collectConnectionNodes(root *entity.Node) []*entity.Node {
	if root == nil {
		return nil
	}
	connections := []*entity.Node{}
	for _, child := range root.Children {
		switch child.Tag {
		case "connection":
			logger.DEBUG(IURBS003, "branch connection node", map[string]any{"tag": child.Tag})
			connections = append(connections, child)
		case "connections":
			defaults := connectionGroupDefaults(child)
			for _, grouped := range child.Children {
				if grouped.Tag != "connection" {
					continue
				}
				logger.DEBUG(IURBS003, "branch grouped connection node", map[string]any{"tag": grouped.Tag})
				connections = append(connections, connectionWithDefaults(grouped, defaults))
			}
		}
	}
	return connections
}

func connectionGroupDefaults(group *entity.Node) map[string]string {
	defaults := map[string]string{}
	if group == nil {
		return defaults
	}
	for _, name := range []string{
		"arrowhead-size", "kind", "color", "stroke-width", "width", "stroke-style",
		"start-arrowhead", "end-arrowhead", "arrowhead", "scale", "coordinate-scale", "grid",
	} {
		if value := strings.TrimSpace(group.Attrs[name]); value != "" {
			defaults[name] = value
		}
	}
	return defaults
}

func connectionWithDefaults(conn *entity.Node, defaults map[string]string) *entity.Node {
	if conn == nil || len(defaults) == 0 {
		return conn
	}
	clone := *conn
	clone.Attrs = map[string]string{}
	for key, value := range defaults {
		clone.Attrs[key] = value
	}
	for key, value := range conn.Attrs {
		clone.Attrs[key] = value
	}
	return &clone
}

type file struct {
	Type     string           `json:"type"`
	Version  int              `json:"version"`
	Source   string           `json:"source"`
	Elements []map[string]any `json:"elements"`
	AppState map[string]any   `json:"appState"`
	Files    map[string]any   `json:"files"`
}

// groupDef holds visual style for an AWS architecture group tag.
type groupDef struct {
	StrokeColor string
	StrokeStyle string
	StrokeWidth int
	IconFile    string // filename inside Architecture-Group-Icons dir, empty = no icon
}

// awsGroups maps xal tag names to their AWS group visual definitions.
var awsGroups = map[string]groupDef{
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
}

const (
	groupIconSize           = 32
	groupHeaderLeftOverflow = 2
	groupHeaderTextInset    = 4
	groupHeaderPadEnd       = 18
	groupHeaderTipMax       = 14
	groupHeaderBorderGap    = 4
	groupFontSize           = 14
	groupTextHeight         = groupFontSize + 4
	groupHeaderTextPadY     = 1
	groupFontFamily         = 2 // Helvetica (normal)
	groupLabelCharW         = 9.6
	itemFallbackIconColor   = "#7758C1"
)

var svgTintColorRE = regexp.MustCompile(`(?i)#[0-9a-f]{3,8}|currentColor`)
var svgCurrentColorRE = regexp.MustCompile(`(?i)currentColor`)

var (
	IUESW001   = share.NewMCode("IUESW-001", "Walk skip too small warning")
	IUESW002   = share.NewMCode("IUESW-002", "Walk generic group icon lookup warning")
	IUESRIA001 = share.NewMCode("IUESRIA-001", "Render icon at invalid item ID warning")
	IUESRIA002 = share.NewMCode("IUESRIA-002", "Render icon at catalog lookup warning")
	IUESRIA003 = share.NewMCode("IUESRIA-003", "Render icon at load SVG warning")
	IUESRC001  = share.NewMCode("IUESRC-001", "Render connections invalid source or destination warning")
	IUESRC002  = share.NewMCode("IUESRC-002", "Render connections source item not rendered warning")
	IUESRC003  = share.NewMCode("IUESRC-003", "Render connections destination item not rendered warning")
)

// tintSVGDataURL makes a group header icon use the same semantic colour as
// its group border and title. White and transparent portions are preserved.
func tintSVGDataURL(dataURL, color string) string {
	if !strings.HasPrefix(dataURL, share.SVGDataURLPrefix) {
		return dataURL
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(dataURL, share.SVGDataURLPrefix))
	if err != nil {
		return dataURL
	}
	tinted := svgTintColorRE.ReplaceAllStringFunc(string(raw), func(found string) string {
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
func normalizeItemSVGDataURL(dataURL string) string {
	if !strings.HasPrefix(dataURL, share.SVGDataURLPrefix) {
		return dataURL
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(dataURL, share.SVGDataURLPrefix))
	if err != nil {
		return dataURL
	}
	normalized := svgCurrentColorRE.ReplaceAllString(string(raw), itemFallbackIconColor)
	return share.SVGDataURLFromBytes([]byte(normalized))
}

// staggerFills are background fill colors for staggered AZ layers.
// Index = StaggerDepth (0 = front/white, 1/2 = progressively darker teal).
var staggerFills = []string{"#ffffff", "#c8e8e8", "#92cecd"}

// staggerBGColor returns the appropriate backgroundColor for a box.
// Boxes that participate in a staggered group get a solid fill so that
// overlapping back-layers are visually distinct.
func staggerBGColor(b *entity.Box) string {
	if !b.InStagger {
		return "transparent"
	}
	idx := b.StaggerDepth
	if idx >= len(staggerFills) {
		idx = len(staggerFills) - 1
	}
	return staggerFills[idx]
}

const (
	itemMaxSize     = 32.0
	itemMinSize     = 16.0
	itemLabelFontPt = 8.0
	itemLabelFontPx = itemLabelFontPt * 96.0 / 72.0
	itemLabelH      = 14.0
	itemLabelW      = 56.0 // text box width for item labels (wider than icon, centred on icon)
	itemLabelCharW  = 6.2
	itemGap         = 8.0
	// Mirrors pptxplan's visual anchor-grid expansion so groups reserve enough
	// top clearance before PPTX adds the grid around each item.
	itemAnchorGridVisualPadPx = 6.0
	excalidrawAnchorGrid      = 5
	excalidrawAnchorPadPx     = 2.0
	excalidrawAnchorCellGapPx = 1.0
)

// paperSizeNames maps (short-side, long-side) → paper name for reverse lookup.
var paperSizeNames = map[[2]int]string{
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
func detectPaperName(w, h float64) string {
	wi, hi := int(w), int(h)
	short, long := wi, hi
	orientation := "portrait"
	if wi > hi {
		short, long = hi, wi
		orientation = "landscape"
	}
	if name, ok := paperSizeNames[[2]int{short, long}]; ok {
		return name + " " + orientation
	}
	return fmt.Sprintf("%d×%d", wi, hi)
}

// BuildJSONWithFS is a convenience wrapper for WASM / embedded builds.
// It uses fsys (typically an embed.FS) for all asset reads instead of the OS
// filesystem.  catalogCSV and svgGroupDir are resolved relative to the root
// of fsys (e.g. "service-catalog.csv" and "svg/Architecture-Group-Icons").
type SceneDependencies struct {
	XaligoRepository     repository.XaligoRepository
	ExcalidrawRepository repository.ExcalidrawRepository
}

func (rcvr *xaligoUsecase) sceneDependencies() SceneDependencies {
	return SceneDependencies{
		XaligoRepository:     rcvr.xaligoRepository,
		ExcalidrawRepository: rcvr.excalidrawRepository,
	}
}

func BuildJSONWithFS(root *entity.Box, fsys fs.FS, catalogCSV, svgGroupDir string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, deps SceneDependencies) ([]byte, error) {
	return BuildJSON(root, svgGroupDir, catalogCSV, "", itemIconSize, connections, abbrevMap, fsys, deps)
}

// BuildJSON converts a entity.Box layout tree into Excalidraw JSON.
// svgGroupDir:  absolute path to Architecture-Group-Icons/ (or FS-relative path when fsys≠nil)
// catalogCSV:   absolute path to service-catalog.csv (or FS-relative path when fsys≠nil)
// projectRoot:  project root directory (used to resolve rel_path from catalog; ignored when fsys≠nil)
// itemIconSize: default maximum icon size (px) for <item> elements.
// connections:  <connection> nodes extracted from the DSL (may be nil).
// abbrevMap:    optional catalog-ID → abbreviation map derived from services.csv.
// fsys:         when non-nil, all asset reads go through this fs.FS (WASM / embedded mode).
func BuildJSON(root *entity.Box, svgGroupDir string, catalogCSV string, projectRoot string, itemIconSize float64, connections []*entity.Node, abbrevMap map[int]string, fsys fs.FS, deps SceneDependencies) ([]byte, error) {
	if root == nil {
		return nil, fmt.Errorf("root layout is nil")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	updated := time.Now().UnixMilli()

	// Outermost Excalidraw frame element representing the paper size.
	frameElem := map[string]any{
		"id": "paper-frame", "type": "frame",
		"x": root.X, "y": root.Y, "width": root.W, "height": root.H,
		"angle":       0,
		"name":        detectPaperName(root.W, root.H),
		"strokeColor": "#bbb", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{}, "roundness": nil,
		"seed": r.Intn(99999999), "version": 1,
		"versionNonce": r.Intn(99999999),
		"isDeleted":    false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false,
	}

	var elements []map[string]any
	elements = append(elements, frameElem)
	files := map[string]any{}

	// 2パス: 1) item を visibleAncestorID ごとに収集, 2) グリッド一括描画
	itemGroups := map[string][]*entity.Box{}
	ancestorBoxes := map[string]*entity.Box{}
	// <frame item-size="N"> overrides the global itemIconSize.
	if v := root.Attrs["item-size"]; v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil && f > 0 {
			itemIconSize = f
		}
	}

	// itemImgRects / itemLblRects / itemImgIDs / itemLblIDs:
	// connection key → bounding rect (x,y,w,h) and element ID of the image / label elements.
	// Populated during renderItemGrid → renderIconAt, used for edge-based connections.
	itemImgRects := map[string][4]float64{}
	itemLblRects := map[string][4]float64{}
	itemImgIDs := map[string]string{}
	itemLblIDs := map[string]string{}

	walk(root, &elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, r, root, itemGroups, ancestorBoxes, itemImgRects, itemImgIDs, deps)
	ancestorIDs := make([]string, 0, len(itemGroups))
	for ancID := range itemGroups {
		ancestorIDs = append(ancestorIDs, ancID)
	}
	sort.Strings(ancestorIDs)
	for _, ancID := range ancestorIDs {
		items := itemGroups[ancID]
		if err := renderItemGrid(items, ancestorBoxes[ancID], &elements, files, catalogCSV, projectRoot, fsys, itemIconSize, r, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps); err != nil {
			return nil, err
		}
	}
	renderConnections(connections, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, &elements, r)
	elements = orderSceneLayers(elements)

	out := file{
		Type:     "excalidraw",
		Version:  2,
		Source:   "https://github.com/xaligo/xaligo",
		Elements: elements,
		AppState: map[string]any{
			"gridSize":            20,
			"viewBackgroundColor": "#ffffff",
		},
		Files: files,
	}
	return json.MarshalIndent(out, "", "  ")
}

// orderSceneLayers keeps connectors below readable content while preserving
// group headers above nested borders.
func orderSceneLayers(elements []map[string]any) []map[string]any {
	base := make([]map[string]any, 0, len(elements))
	headShapes := make([]map[string]any, 0)
	connectors := make([]map[string]any, 0)
	anchorBackgrounds := make([]map[string]any, 0)
	anchorContent := make([]map[string]any, 0)
	headContent := make([]map[string]any, 0)
	for _, el := range elements {
		custom, _ := el["customData"].(map[string]any)
		if isHeader, _ := custom["xaligoGroupHeader"].(bool); isHeader {
			headShapes = append(headShapes, el)
			continue
		}
		if isContent, _ := custom["xaligoGroupHeaderContent"].(bool); isContent {
			headContent = append(headContent, el)
			continue
		}
		if isAnchorBackground, _ := custom["xaligoAnchorBackground"].(bool); isAnchorBackground {
			anchorBackgrounds = append(anchorBackgrounds, el)
			continue
		}
		if isAnchorContent, _ := custom["xaligoAnchorContent"].(bool); isAnchorContent {
			anchorContent = append(anchorContent, el)
			continue
		}
		typ, _ := el["type"].(string)
		if isJunction, _ := custom["xaligoJunction"].(bool); typ == "arrow" || typ == "line" || isJunction {
			connectors = append(connectors, el)
			continue
		}
		base = append(base, el)
	}
	ordered := append(base, headShapes...)
	ordered = append(ordered, connectors...)
	ordered = append(ordered, anchorBackgrounds...)
	ordered = append(ordered, anchorContent...)
	return append(ordered, headContent...)
}

func avoidGroupHeaderBorderOverlap(x, y, w, h float64, ownBorderID string, elements []map[string]any) float64 {
	adjustedY := y
	for pass := 0; pass < 4; pass++ {
		nextY := adjustedY
		for _, el := range elements {
			if id, _ := el["id"].(string); id == ownBorderID {
				continue
			}
			custom, _ := el["customData"].(map[string]any)
			if isBorder, _ := custom["xaligoGroupBorder"].(bool); !isBorder {
				continue
			}
			bx, okX := el["x"].(float64)
			by, okY := el["y"].(float64)
			bw, okW := el["width"].(float64)
			bh, okH := el["height"].(float64)
			if !okX || !okY || !okW || !okH || horizontalOverlap(x, x+w, bx, bx+bw) <= 0 {
				continue
			}
			for _, lineY := range []float64{by, by + bh} {
				if lineY >= adjustedY-float64(groupHeaderBorderGap) && lineY <= adjustedY+h+float64(groupHeaderBorderGap) {
					nextY = math.Max(nextY, lineY+float64(groupHeaderBorderGap))
				}
			}
		}
		if math.Abs(nextY-adjustedY) < 0.01 {
			break
		}
		adjustedY = nextY
	}
	return adjustedY
}

func horizontalOverlap(a0, a1, b0, b1 float64) float64 {
	return math.Max(0, math.Min(math.Max(a0, a1), math.Max(b0, b1))-math.Max(math.Min(a0, a1), math.Min(b0, b1)))
}

func alignGroupBorderTopToHeader(borderID string, topY, bottomY float64, elements []map[string]any) {
	for i := range elements {
		id, _ := elements[i]["id"].(string)
		if id != borderID {
			continue
		}
		if topY <= bottomY-MinBoxHeight {
			elements[i]["y"] = topY
			elements[i]["height"] = bottomY - topY
		}
		return
	}
}

func walk(b *entity.Box, elements *[]map[string]any, files map[string]any, svgGroupDir string, catalogCSV string, projectRoot string, fsys fs.FS, r *rand.Rand, visibleAncestor *entity.Box, itemGroups map[string][]*entity.Box, ancestorBoxes map[string]*entity.Box, itemImgRects map[string][4]float64, itemImgIDs map[string]string, deps SceneDependencies) {
	if IsItemLike(b.Tag) {
		// 描画はしない: visibleAncestor に結び付けて収集のみ (<item> / <spacer> 共通)
		key := visibleAncestor.ID
		itemGroups[key] = append(itemGroups[key], b)
		ancestorBoxes[key] = visibleAncestor
		return
	}

	// selfVisible=false のとき: 自身の描画 (枠・アイコン・ラベル) はスキップするが
	// 子要素の描画は継続する (親子関係なく個別に制御可能)。
	selfVisible := b.Attrs["visible"] != "false"

	if b.Tag != "frame" && b.Tag != "port" && (b.W < MinBoxWidth || b.H < MinBoxHeight) {
		logger.WARN(IUESW001, "skipping too small element", map[string]any{"label": b.Label, "tag": b.Tag, "width": b.W, "height": b.H, "minWidth": MinBoxWidth, "minHeight": MinBoxHeight})
		// 子の item も同じ visibleAncestor に結び付けて収集
		for _, c := range b.Children {
			if IsItemLike(c.Tag) {
				key := visibleAncestor.ID
				itemGroups[key] = append(itemGroups[key], c)
				ancestorBoxes[key] = visibleAncestor
			} else {
				walk(c, elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, r, visibleAncestor, itemGroups, ancestorBoxes, itemImgRects, itemImgIDs, deps)
			}
		}
		return
	}

	if b.Tag != "frame" && selfVisible {
		updated := time.Now().UnixMilli()

		noBorder := b.Attrs["border"] == "none"

		if gd, isGroup := awsGroups[b.Tag]; isGroup {
			// ── AWS group border ────────────────────────────────────
			rectID := fmt.Sprintf("%s-rect", b.ID)
			groupStroke := gd.StrokeColor
			if noBorder {
				groupStroke = "transparent"
			}
			*elements = append(*elements, map[string]any{
				"id": rectID, "type": "rectangle",
				"x": b.X, "y": b.Y, "width": b.W, "height": b.H,
				"angle":       0,
				"strokeColor": groupStroke, "backgroundColor": staggerBGColor(b),
				"fillStyle":   "solid",
				"strokeWidth": gd.StrokeWidth, "strokeStyle": gd.StrokeStyle,
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": r.Intn(99999999), "version": 1,
				"versionNonce": r.Intn(99999999),
				"isDeleted":    false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"customData": map[string]any{"xaligoGroupBorder": true},
			})
			registerConnectionEndpoint(b, rectID, [4]float64{b.X, b.Y, b.W, b.H}, itemImgRects, itemImgIDs)

			// ── Group icon ──────────────────────────────────────────
			headerX := b.X - groupHeaderLeftOverflow
			textX := headerX + groupHeaderTextInset
			var iconDataURL, iconFileID, iconBackground string
			if b.Tag == "generic-group" && strings.TrimSpace(b.Attrs["icon-id"]) != "" {
				catalogID, _ := strconv.Atoi(strings.TrimSpace(b.Attrs["icon-id"]))
				var entry entity.CatalogEntry
				var err error
				if fsys != nil {
					entry, err = deps.XaligoRepository.LookupCatalogByIDFS(fsys, catalogCSV, catalogID)
				} else {
					entry, err = deps.XaligoRepository.LookupCatalogByID(catalogCSV, catalogID)
				}
				if err == nil && entry.DataURL == "" && entry.RelPath != "" && projectRoot != "" {
					entry.DataURL, err = deps.ExcalidrawRepository.SvgToDataURL(filepath.Join(projectRoot, entry.RelPath))
				}
				if err != nil {
					logger.WARN(IUESW002, "generic group icon lookup failed", map[string]any{"catalogID": catalogID, "error": err})
				} else {
					iconDataURL = entry.DataURL
					iconFileID = fmt.Sprintf("group-cat-%d", catalogID)
					iconBackground = deps.ExcalidrawRepository.SVGBGColor(entry.DataURL)
				}
			} else if gd.IconFile != "" && svgGroupDir != "" {
				iconPath := filepath.Join(svgGroupDir, gd.IconFile)
				var err error
				if fsys != nil {
					// In embedded mode, use forward slashes even on Windows.
					iconPath = svgGroupDir + "/" + gd.IconFile
					iconDataURL, err = deps.ExcalidrawRepository.SvgToDataURLFS(fsys, iconPath)
				} else {
					iconDataURL, err = deps.ExcalidrawRepository.SvgToDataURL(iconPath)
				}
				if err != nil {
					iconDataURL = ""
				}
				iconFileID = deps.ExcalidrawRepository.FileID(gd.IconFile)
				iconBackground = "transparent"
			}
			if iconDataURL != "" {
				iconDataURL = tintSVGDataURL(iconDataURL, gd.StrokeColor)
				iconBackground = "transparent"
				textX = headerX + float64(groupIconSize) + groupHeaderTextInset
			}
			lblW := textWidth(b.Label, groupLabelCharW)
			headerBackground := staggerBGColor(b)
			if headerBackground == "transparent" {
				headerBackground = "#ffffff"
			}
			// Extend the opaque header mask beyond the group's left border so the
			// vertical border cannot show through beside a catalog icon.
			headerH := float64(groupTextHeight + groupHeaderTextPadY*2)
			if iconDataURL != "" {
				headerH = float64(groupIconSize)
			}
			headerTip := math.Min(groupHeaderTipMax, headerH/2)
			headerW := textX + lblW + groupHeaderPadEnd + headerTip - headerX
			headerY := avoidGroupHeaderBorderOverlap(headerX, b.Y-headerH/2, headerW, headerH, rectID, *elements)
			alignGroupBorderTopToHeader(rectID, headerY+headerH/2, b.Y+b.H, *elements)
			*elements = append(*elements, map[string]any{
				"id": fmt.Sprintf("%s-header-bg", b.ID), "type": "line",
				"x": headerX, "y": headerY,
				"width": headerW, "height": headerH,
				"points": [][]float64{{0, 0}, {headerW - headerTip, 0}, {headerW, headerH / 2}, {headerW - headerTip, headerH}, {0, headerH}, {0, 0}},
				"angle":  0, "strokeColor": gd.StrokeColor, "backgroundColor": headerBackground,
				"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": r.Intn(99999999), "version": 1, "versionNonce": r.Intn(99999999),
				"isDeleted": false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"customData": map[string]any{"xaligoGroupHeader": true},
			})
			if iconDataURL != "" {
				*elements = append(*elements, map[string]any{
					"id": fmt.Sprintf("%s-icon", b.ID), "type": "image",
					"x": headerX, "y": headerY + (headerH-float64(groupIconSize))/2,
					"width": float64(groupIconSize), "height": float64(groupIconSize),
					"fileId": iconFileID, "status": "saved", "scale": []int{1, 1},
					"strokeColor": "transparent", "backgroundColor": iconBackground,
					"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
					"roughness": 0, "opacity": 100, "angle": 0,
					"version": 1, "versionNonce": r.Intn(99999999),
					"isDeleted": false, "groupIds": []string{},
					"frameId": nil, "boundElements": nil,
					"updated": updated, "link": nil, "locked": false,
					"customData": map[string]any{"xaligoGroupHeaderContent": true},
				})
				if _, exists := files[iconFileID]; !exists {
					files[iconFileID] = map[string]any{
						"mimeType": "image/svg+xml", "id": iconFileID, "dataURL": iconDataURL,
						"created": updated, "lastRetrieved": updated,
					}
				}
			}

			// ── AWS group label ─────────────────────────────────────
			textY := headerY + (headerH-float64(groupTextHeight))/2
			// groupFontFamily=2 (Helvetica 14px): ~7.5px/rune
			*elements = append(*elements, map[string]any{
				"id": fmt.Sprintf("%s-label", b.ID), "type": "text",
				"x": textX, "y": textY,
				"width": lblW, "height": float64(groupTextHeight),
				"angle":       0,
				"strokeColor": gd.StrokeColor, "backgroundColor": "transparent",
				"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": r.Intn(99999999), "version": 1,
				"versionNonce": r.Intn(99999999),
				"isDeleted":    false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"text": b.Label, "fontSize": groupFontSize, "fontFamily": groupFontFamily,
				"textAlign": "left", "verticalAlign": "middle",
				"containerId": nil, "originalText": b.Label, "lineHeight": 1.25,
				"customData": map[string]any{"xaligoGroupHeaderContent": true},
			})
		} else if !isLayoutTag(b.Tag) {
			// ── Generic tag: rectangle + label ──────────────────────
			rectID := fmt.Sprintf("%s-rect", b.ID)
			textID := fmt.Sprintf("%s-text", b.ID)
			genStroke := "#1e1e1e"
			if noBorder {
				genStroke = "transparent"
			}
			backgroundColor := "transparent"
			fillStyle := "hachure"
			roundness := map[string]any{"type": 3}
			if b.Tag == "rectangle" {
				fillStyle = "solid"
			}
			if b.Tag == "port" {
				backgroundColor = "#ffffff"
				fillStyle = "solid"
				roundness = nil
			}
			boundElements := any(nil)
			if b.Label != "" {
				boundElements = []map[string]any{{"type": "text", "id": textID}}
			}
			*elements = append(*elements, map[string]any{
				"id": rectID, "type": "rectangle",
				"x": b.X, "y": b.Y, "width": b.W, "height": b.H,
				"angle": 0, "strokeColor": genStroke, "backgroundColor": backgroundColor,
				"fillStyle": fillStyle, "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": roundness,
				"seed": r.Intn(99999999), "version": 1,
				"versionNonce":  r.Intn(99999999),
				"isDeleted":     false,
				"boundElements": boundElements,
				"updated":       updated, "link": nil, "locked": false,
			})
			if b.Tag == "rectangle" || b.Tag == "port" {
				registerConnectionEndpoint(b, rectID, [4]float64{b.X, b.Y, b.W, b.H}, itemImgRects, itemImgIDs)
			}
			if b.Label != "" {
				fontSize := attrFloat(b.Attrs["font-size"], 20)
				textX, textY := b.X+12, b.Y+12
				textW, textH := textWidth(b.Label, fontSize*0.5), math.Ceil(fontSize*1.2)
				textAlign, verticalAlign := "left", "top"
				if b.Tag == "rectangle" || b.Tag == "port" {
					textX, textY = b.X+4, b.Y+2
					textW, textH = math.Max(1, b.W-8), math.Max(1, b.H-4)
					textAlign, verticalAlign = "center", "middle"
				}
				*elements = append(*elements, map[string]any{
					"id": textID, "type": "text",
					"x": textX, "y": textY,
					"width": textW, "height": textH,
					"angle":       0,
					"strokeColor": "#1e1e1e", "backgroundColor": "transparent",
					"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
					"roughness": 0, "opacity": 100,
					"groupIds": []string{}, "roundness": nil,
					"seed": r.Intn(99999999), "version": 1,
					"versionNonce": r.Intn(99999999),
					"isDeleted":    false, "boundElements": nil,
					"updated": updated, "link": nil, "locked": false,
					"text": b.Label, "fontSize": fontSize, "fontFamily": 1,
					"textAlign": textAlign, "verticalAlign": verticalAlign,
					"containerId": rectID, "originalText": b.Label, "lineHeight": 1.2,
				})
			}
		}
	}

	// Stagger background layers: render border + label only, skip children.
	if b.IsStaggerBg {
		return
	}
	// 非表示要素は visibleAncestor を引き継ぐ (子の item が正しい親に紐付くよう)
	nextVisible := b
	if !selfVisible {
		nextVisible = visibleAncestor
	}
	for _, c := range b.Children {
		walk(c, elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, r, nextVisible, itemGroups, ancestorBoxes, itemImgRects, itemImgIDs, deps)
	}
}

func registerConnectionEndpoint(b *entity.Box, elementID string, rect [4]float64, endpointRects map[string][4]float64, endpointIDs map[string]string) {
	if b == nil || endpointRects == nil || endpointIDs == nil || elementID == "" {
		return
	}
	key := strings.TrimSpace(b.Attrs[internalConnectionKeyAttr])
	if key == "" {
		return
	}
	endpointRects[key] = rect
	endpointIDs[key] = elementID
}

// isLayoutTag reports whether a tag is a pure layout container
// (<row>, <col>, <container>) that should not render any visible border or label.
func isLayoutTag(tag string) bool {
	return tag == "row" || tag == "col" || tag == "container" || IsBlank(tag)
}

// textWidth estimates the rendered width of a string in pixels.
// charW: approximate pixel width per rune (font-specific).
func textWidth(s string, charW float64) float64 {
	return math.Ceil(displayColumns(s)*charW) + 8
}

func itemLabelHeight(label string) float64 {
	lines := 1
	for _, line := range strings.Split(label, "\n") {
		wrapped := int(math.Ceil(displayColumns(line) * itemLabelCharW / itemLabelW))
		if wrapped < 1 {
			wrapped = 1
		}
		lines += wrapped - 1
	}
	lineH := itemLabelFontPx * 1.25
	return math.Max(itemLabelH, math.Ceil(float64(lines)*lineH))
}

func displayColumns(s string) float64 {
	cols := 0.0
	for _, r := range s {
		cols += runeColumns(r)
	}
	return cols
}

func runeColumns(r rune) float64 {
	switch {
	case r == '\t':
		return 4
	case r < 0x20:
		return 0
	case r >= 0x1100 && (r <= 0x115F ||
		r == 0x2329 || r == 0x232A ||
		(r >= 0x2E80 && r <= 0xA4CF && r != 0x303F) ||
		(r >= 0xAC00 && r <= 0xD7A3) ||
		(r >= 0xF900 && r <= 0xFAFF) ||
		(r >= 0xFE10 && r <= 0xFE19) ||
		(r >= 0xFE30 && r <= 0xFE6F) ||
		(r >= 0xFF00 && r <= 0xFF60) ||
		(r >= 0xFFE0 && r <= 0xFFE6)):
		return 2
	default:
		return 1
	}
}

// parseItemAlign parses an align attribute value (e.g. "top-left", "middle-center")
// into vertical ("top"/"middle"/"bottom") and horizontal ("left"/"center"/"right") parts.
// Defaults to "middle" / "center" when absent or unrecognised.
func parseItemAlign(align string) (vert, horiz string) {
	vert, horiz = "middle", "center"
	parts := strings.SplitN(strings.ToLower(strings.TrimSpace(align)), "-", 2)
	if len(parts) == 2 {
		if parts[0] == "top" || parts[0] == "middle" || parts[0] == "bottom" {
			vert = parts[0]
		}
		if parts[1] == "left" || parts[1] == "center" || parts[1] == "right" || parts[1] == "spread" {
			horiz = parts[1]
		}
	}
	return
}

// renderItemGrid lays out all items collected under the same visibleAncestor as
// a compact grid within the ancestor's content area.
func renderItemGrid(items []*entity.Box, ancestor *entity.Box, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, maxSize float64, r *rand.Rand, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, abbrevMap map[int]string, deps SceneDependencies) error {
	if catalogCSV == "" || len(items) == 0 || ancestor == nil {
		return nil
	}
	nItems := len(items)
	vert, horiz := parseItemAlign(ancestor.Attrs["align"])

	var areaX, areaY, areaW, areaH float64

	if _, isGroup := awsGroups[ancestor.Tag]; isGroup {
		// When a group's children are ALL items, go used layoutRow (no GroupTopInset).
		// In that case we must also skip the topInset here so icons aren't pushed off-screen.
		allItemChildren := true
		for _, ch := range ancestor.Children {
			if !IsItemLike(ch.Tag) {
				allItemChildren = false
				break
			}
		}

		if allItemChildren {
			// Reserve the group tag band plus the PPTX anchor-grid expansion. Without
			// this, the later PPTX-only anchor grid can extend into the tag area.
			topClearance := groupHeaderHeightForItems(ancestor)/2 + itemAnchorGridVisualPadPx + float64(groupHeaderBorderGap)
			if topClearance < itemGap {
				topClearance = itemGap
			}
			areaX = ancestor.X + GroupSideInset
			areaY = ancestor.Y + topClearance
			areaW = ancestor.W - GroupSideInset*2
			areaH = ancestor.H - topClearance - itemGap
		} else {
			// Content area: below the header row.
			areaX = ancestor.X + GroupSideInset
			areaY = ancestor.Y + GroupTopInset + itemGap
			areaW = ancestor.W - GroupSideInset*2
			areaH = ancestor.H - GroupTopInset - itemGap*2
		}
	} else {
		// 汎用コンテナ (frame, container, col など).
		areaX = ancestor.X + itemGap
		areaY = ancestor.Y + itemGap
		areaW = ancestor.W - itemGap*2
		areaH = ancestor.H - itemGap*2
	}

	labelBoxH := estimateMaxItemLabelHeight(items, catalogCSV, fsys, abbrevMap, deps.XaligoRepository)
	cols, rows, iconSize := chooseItemGrid(nItems, areaW, areaH, maxSize, labelBoxH)
	if cols <= 0 || rows <= 0 {
		return nil
	}
	cellW := iconSize
	cellH := iconSize + 4 + labelBoxH
	totalW := cellW*float64(cols) + itemGap*float64(cols-1)
	totalH := cellH*float64(rows) + itemGap*float64(rows-1)

	startX, stepX := gridAxis(areaX, areaW, totalW, cellW, cols, horiz)
	startY, stepY := gridAxis(areaY, areaH, totalH, cellH, rows, vert)

	for i, item := range items {
		col := i % cols
		row := i / cols
		iconX := startX + float64(col)*stepX + math.Max(0, (cellW-iconSize)/2)
		iconY := startY + float64(row)*stepY
		if strings.TrimSpace(item.Attrs["id"]) != "" {
			dx, dy, err := itemIconOffset(item)
			if err != nil {
				return err
			}
			iconX += dx
			iconY += dy
			if err := validateItemIconBounds(item, ancestor, iconX, iconY, iconSize); err != nil {
				return err
			}
		}
		connectionKey := strings.TrimSpace(item.Attrs[internalConnectionKeyAttr])
		if connectionKey == "" {
			connectionKey = item.ID
		}
		renderIconAt(item.ID, connectionKey, item.Attrs["id"], iconX, iconY, iconSize, elements, files, catalogCSV, projectRoot, fsys, r, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps)
	}
	return nil
}

func itemIconOffset(item *entity.Box) (float64, float64, error) {
	if item == nil {
		return 0, 0, nil
	}
	dx, err := parseOptionalFloatAttr(item, "dx")
	if err != nil {
		return 0, 0, err
	}
	dy, err := parseOptionalFloatAttr(item, "dy")
	if err != nil {
		return 0, 0, err
	}
	return dx, dy, nil
}

func parseOptionalFloatAttr(item *entity.Box, attr string) (float64, error) {
	value := strings.TrimSpace(item.Attrs[attr])
	if value == "" {
		return 0, nil
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("<item %s=%q> must be a number", attr, value)
	}
	return parsed, nil
}

func validateItemIconBounds(item *entity.Box, ancestor *entity.Box, x, y, size float64) error {
	const epsilon = 1e-6
	if item == nil || ancestor == nil {
		return nil
	}
	minX := ancestor.X
	minY := ancestor.Y
	maxX := ancestor.X + ancestor.W
	maxY := ancestor.Y + ancestor.H
	if x+epsilon < minX || y+epsilon < minY || x+size > maxX+epsilon || y+size > maxY+epsilon {
		return fmt.Errorf("<item id=%q> icon offset moves icon outside parent %q bounds: icon=(%.1f,%.1f,%.1f,%.1f), parent=(%.1f,%.1f,%.1f,%.1f)",
			strings.TrimSpace(item.Attrs["id"]), ancestor.Tag, x, y, size, size, minX, minY, ancestor.W, ancestor.H)
	}
	return nil
}

func groupHeaderHeightForItems(ancestor *entity.Box) float64 {
	headerH := float64(groupTextHeight + groupHeaderTextPadY*2)
	if ancestor == nil {
		return headerH
	}
	if ancestor.Tag == "generic-group" && strings.TrimSpace(ancestor.Attrs["icon-id"]) != "" {
		return float64(groupIconSize)
	}
	if gd, ok := awsGroups[ancestor.Tag]; ok && gd.IconFile != "" {
		return float64(groupIconSize)
	}
	return headerH
}

func estimateMaxItemLabelHeight(items []*entity.Box, catalogCSV string, fsys fs.FS, abbrevMap map[int]string, catalogRepo repository.XaligoRepository) float64 {
	maxH := itemLabelH
	for _, item := range items {
		id, err := strconv.Atoi(strings.TrimSpace(item.Attrs["id"]))
		if err != nil {
			continue
		}
		label := ""
		if abbrevMap != nil {
			label = abbrevMap[id]
		}
		if label == "" {
			var ce entity.CatalogEntry
			if fsys != nil {
				ce, err = catalogRepo.LookupCatalogByIDFS(fsys, catalogCSV, id)
			} else {
				ce, err = catalogRepo.LookupCatalogByID(catalogCSV, id)
			}
			if err != nil {
				continue
			}
			label = entity.ItemShortName(ce.Service)
		}
		maxH = math.Max(maxH, itemLabelHeight(label))
	}
	return maxH
}

func chooseItemGrid(n int, areaW, areaH, maxSize float64, labelBoxH float64) (cols int, rows int, iconSize float64) {
	if n <= 0 || areaW <= 0 || areaH <= 0 {
		return 0, 0, 0
	}
	bestScore := -1.0
	for c := 1; c <= n; c++ {
		r := int(math.Ceil(float64(n) / float64(c)))
		cellW := (areaW - itemGap*float64(c-1)) / float64(c)
		cellH := (areaH - itemGap*float64(r-1)) / float64(r)
		size := math.Min(cellW, cellH-4-labelBoxH)
		size = math.Min(size, maxSize)
		if size < itemMinSize {
			continue
		}
		usedW := size*float64(c) + itemGap*float64(c-1)
		usedH := (size+4+labelBoxH)*float64(r) + itemGap*float64(r-1)
		if usedW-areaW > 1e-6 || usedH-areaH > 1e-6 {
			continue
		}
		aspectPenalty := math.Abs(float64(c)/float64(r) - areaW/math.Max(1, areaH))
		score := size*100 - aspectPenalty
		if score > bestScore {
			bestScore = score
			cols = c
			rows = r
			iconSize = size
		}
	}
	if cols == 0 {
		cols = n
		rows = 1
		iconSize = itemMinSize
	}
	return cols, rows, iconSize
}

func gridAxis(areaStart, areaSize, totalSize, cellSize float64, count int, align string) (start, step float64) {
	if count <= 1 {
		return areaStart + math.Max(0, (areaSize-cellSize)/2), 0
	}
	switch align {
	case "left", "top":
		return areaStart, cellSize + itemGap
	case "right", "bottom":
		return areaStart + math.Max(0, areaSize-totalSize), cellSize + itemGap
	case "spread":
		gap := (areaSize - cellSize*float64(count)) / float64(count+1)
		if gap < itemGap {
			gap = itemGap
		}
		return areaStart + gap, cellSize + gap
	default:
		return areaStart + math.Max(0, (areaSize-totalSize)/2), cellSize + itemGap
	}
}

// renderIconAt draws a single service icon (image + label) at an explicit position.
// itemImgRects/itemLblRects/itemImgIDs/itemLblIDs are populated with the bounding rect
// and element ID of the image and label elements, keyed by the unique item connection key.
func renderIconAt(boxID, connectionKey, idAttr string, iconX, iconY, iconSize float64, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, r *rand.Rand, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, abbrevMap map[int]string, deps SceneDependencies) {
	if catalogCSV == "" {
		return
	}
	idAttr = strings.TrimSpace(idAttr)
	if idAttr == "" {
		return
	}

	// 1:1 — id は単一の整数
	id, err := strconv.Atoi(idAttr)
	if err != nil {
		logger.WARN(IUESRIA001, "item ID must be a single integer", map[string]any{"id": idAttr, "error": err})
		return
	}
	var ce entity.CatalogEntry
	if fsys != nil {
		ce, err = deps.XaligoRepository.LookupCatalogByIDFS(fsys, catalogCSV, id)
	} else {
		ce, err = deps.XaligoRepository.LookupCatalogByID(catalogCSV, id)
	}
	if err != nil {
		logger.WARN(IUESRIA002, "catalog lookup failed", map[string]any{"id": id, "error": err})
		return
	}
	if ce.DataURL == "" && ce.RelPath != "" && projectRoot != "" {
		svgPath := filepath.Join(projectRoot, ce.RelPath)
		if du, err2 := deps.ExcalidrawRepository.SvgToDataURL(svgPath); err2 == nil {
			ce.DataURL = du
		} else {
			logger.WARN(IUESRIA003, "cannot load SVG", map[string]any{"id": id, "path": svgPath, "error": err2})
		}
	}
	if ce.DataURL == "" {
		return
	}
	ce.DataURL = normalizeItemSVGDataURL(ce.DataURL)

	updated := time.Now().UnixMilli()
	fid := fmt.Sprintf("item-cat-%d", id)
	if _, exists := files[fid]; !exists {
		files[fid] = map[string]any{
			"mimeType": "image/svg+xml", "id": fid,
			"dataURL": ce.DataURL,
			"created": updated, "lastRetrieved": updated,
		}
	}
	seed := r.Intn(99999999)
	iconID := fmt.Sprintf("%s-item", boxID)
	anchorGroupID := fmt.Sprintf("%s-anchor", boxID)
	var label string
	if abbrevMap != nil {
		label = abbrevMap[id]
	}
	if label == "" {
		label = entity.ItemShortName(ce.Service)
	}
	labelH := itemLabelHeight(label)
	labelY := iconY + iconSize + 4
	labelX := iconX + (iconSize-itemLabelW)/2 // centre label on icon
	anchorX := iconX - excalidrawAnchorPadPx
	anchorY := iconY - excalidrawAnchorPadPx
	anchorW := iconSize + excalidrawAnchorPadPx*2
	anchorH := iconSize + excalidrawAnchorPadPx*2
	// Record bounding rects and element IDs for edge-based connection arrows.
	if itemImgRects != nil {
		itemImgRects[connectionKey] = [4]float64{iconX, iconY, iconSize, iconSize}
		itemImgIDs[connectionKey] = iconID
	}
	appendExcalidrawAnchorGrid(elements, boxID, anchorGroupID, anchorX, anchorY, anchorW, anchorH, seed+1, updated)
	*elements = append(*elements, map[string]any{
		"id": iconID, "type": "image",
		"x": iconX, "y": iconY,
		"width": iconSize, "height": iconSize,
		"fileId": fid, "status": "saved",
		"scale":       []int{1, 1},
		"strokeColor": "transparent", "backgroundColor": deps.ExcalidrawRepository.SVGBGColor(ce.DataURL),
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100, "angle": 0,
		"groupIds": []string{anchorGroupID}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false, "frameId": nil,
		"customData": map[string]any{"xaligoAnchorContent": true},
	})
	// Record label bounding rect for bottom-side connection binding.
	if itemLblRects != nil {
		itemLblRects[connectionKey] = [4]float64{labelX, labelY, itemLabelW, labelH}
		itemLblIDs[connectionKey] = iconID + "-lbl"
	}
	textSeed := r.Intn(99999999)
	*elements = append(*elements, map[string]any{
		"id": iconID + "-lbl", "type": "text",
		"x": labelX, "y": labelY,
		"width": itemLabelW, "height": labelH,
		"angle":       0,
		"strokeColor": "#1e1e1e", "backgroundColor": "transparent",
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100,
		"groupIds": []string{anchorGroupID}, "roundness": nil,
		"seed": textSeed, "version": 1, "versionNonce": textSeed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false, "frameId": nil,
		"text": label, "rawText": label, "originalText": label,
		"fontSize": itemLabelFontPx, "fontFamily": 4,
		"textAlign": "center", "verticalAlign": "top",
		"containerId": nil, "lineHeight": 1.25,
		"customData": map[string]any{"xaligoAnchorContent": true},
	})
}

func appendExcalidrawAnchorGrid(elements *[]map[string]any, boxID, groupID string, x, y, w, h float64, seed int, updated int64) {
	if elements == nil || excalidrawAnchorGrid <= 0 {
		return
	}
	cellW := w / float64(excalidrawAnchorGrid)
	cellH := h / float64(excalidrawAnchorGrid)
	for row := 0; row < excalidrawAnchorGrid; row++ {
		for col := 0; col < excalidrawAnchorGrid; col++ {
			cellX := x + float64(col)*cellW + excalidrawAnchorCellGapPx
			cellY := y + float64(row)*cellH + excalidrawAnchorCellGapPx
			cellWidth := math.Max(1, cellW-excalidrawAnchorCellGapPx*2)
			cellHeight := math.Max(1, cellH-excalidrawAnchorCellGapPx*2)
			cellSeed := seed + row*excalidrawAnchorGrid + col
			*elements = append(*elements, map[string]any{
				"id": fmt.Sprintf("%s-anchor-bg-%02d-%02d", boxID, row, col), "type": "rectangle",
				"x": cellX, "y": cellY,
				"width": cellWidth, "height": cellHeight,
				"angle":       0,
				"strokeColor": "#ffffff", "backgroundColor": "#ffffff",
				"fillStyle": "solid", "strokeWidth": 0, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{groupID}, "roundness": nil,
				"seed": cellSeed, "version": 1, "versionNonce": cellSeed,
				"isDeleted": false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false, "frameId": nil,
				"customData": map[string]any{"xaligoAnchorBackground": true},
			})
		}
	}
}

// connectionSide determines which edge exits src (srcSide) and enters dst (dstSide)
// based on the direction between their center points.
// Returns "top", "bottom", "left", or "right".
func connectionSide(srcCx, srcCy, dstCx, dstCy float64) (srcSide, dstSide string) {
	dx := dstCx - srcCx
	dy := dstCy - srcCy
	if math.Abs(dx) >= math.Abs(dy) {
		if dx >= 0 {
			return "right", "left"
		}
		return "left", "right"
	}
	if dy >= 0 {
		return "bottom", "top"
	}
	return "top", "bottom"
}

func connectionBendPoints(conn *entity.Node) []pt {
	scale := 1.0
	if value, ok := positiveFloatAttr(conn, "coordinate-scale", "scale"); ok {
		scale = value
	}
	return parseConnectorBends(connectionBends(conn), scale)
}

func sideTowardPoint(rect [4]float64, point pt) string {
	cx := rect[0] + rect[2]/2
	cy := rect[1] + rect[3]/2
	dx := point.X - cx
	dy := point.Y - cy
	if math.Abs(dx) >= math.Abs(dy) {
		if dx >= 0 {
			return "right"
		}
		return "left"
	}
	if dy >= 0 {
		return "bottom"
	}
	return "top"
}

// rectEdgePoint returns the midpoint of the named edge of a rectangle.
// rect = [x, y, w, h]; side is "top", "bottom", "left", or "right".
func rectEdgePoint(rect [4]float64, side string) [2]float64 {
	x, y, w, h := rect[0], rect[1], rect[2], rect[3]
	cx := x + w/2
	cy := y + h/2
	switch side {
	case "top":
		return [2]float64{cx, y}
	case "bottom":
		return [2]float64{cx, y + h}
	case "left":
		return [2]float64{x, cy}
	default: // "right"
		return [2]float64{x + w, cy}
	}
}

// fixedPointForSide returns the normalized [x, y] fixedPoint on an element's bounding box
// that corresponds to the given side. This matches Excalidraw's binding coordinate system:
// [0,0]=top-left, [1,1]=bottom-right; each side midpoint:
//
//	top=[0.5,0], bottom=[0.5,1], left=[0,0.5], right=[1,0.5]
func fixedPointForSide(side string) [2]float64 {
	switch side {
	case "top":
		return [2]float64{0.5, 0}
	case "bottom":
		return [2]float64{0.5, 1}
	case "left":
		return [2]float64{0, 0.5}
	default: // "right"
		return [2]float64{1, 0.5}
	}
}

func fixedPointForAnchor(anchor connectionAnchorSpec) [2]float64 {
	pos := (float64(anchor.slot) + 0.5) / 5.0
	switch anchor.side {
	case sideTop:
		return [2]float64{pos, 0}
	case sideBottom:
		return [2]float64{pos, 1}
	case sideLeft:
		return [2]float64{0, pos}
	default:
		return [2]float64{1, pos}
	}
}

func rectFixedPoint(rect [4]float64, fp [2]float64) [2]float64 {
	return [2]float64{rect[0] + rect[2]*fp[0], rect[1] + rect[3]*fp[1]}
}

// renderConnections generates elbowed arrow elements for each <connection> node and
// updates the boundElements of the bound source/destination elements — required by
// Excalidraw so that the application recognises the binding relationship.
//
// src/dst are resolved item connection keys; the corresponding item rects and element IDs
// must already be populated in itemImgRects/itemLblRects/itemImgIDs/itemLblIDs by renderIconAt.
// Arrows start/end at the actual element edge; when the connection exits/enters from the
// bottom the label text element is used instead of the image element.
func renderConnections(connections []*entity.Node, itemImgRects map[string][4]float64, itemLblRects map[string][4]float64, itemImgIDs map[string]string, itemLblIDs map[string]string, elements *[]map[string]any, r *rand.Rand) {
	if len(connections) == 0 {
		return
	}
	updated := time.Now().UnixMilli()

	// boundMap accumulates the arrow binding entries that must be written back
	// into each referenced element's boundElements array.
	// key = element ID, value = slice of {"type":"arrow","id":<arrowID>}
	boundMap := map[string][]map[string]any{}
	type junctionCandidate struct {
		edge  [2]float64
		side  string
		color string
		count int
		seed  int
	}
	junctionCandidates := map[string]*junctionCandidate{}

	orderedConnections := append([]*entity.Node(nil), connections...)
	sort.SliceStable(orderedConnections, func(i, j int) bool {
		return connectionKindPriority(connectionKind(orderedConnections[i])) < connectionKindPriority(connectionKind(orderedConnections[j]))
	})
	obstacles := excalidrawRouteObstacles(*elements)
	placed := [][]segment{}
	routePaths := map[string][]pt{}

	for i, conn := range orderedConnections {
		srcIDStr := strings.TrimSpace(conn.Attrs["src"])
		dstIDStr := strings.TrimSpace(conn.Attrs["dst"])
		srcKey := strings.TrimSpace(conn.Attrs[internalConnectionSrcKeyAttr])
		dstKey := strings.TrimSpace(conn.Attrs[internalConnectionDstKeyAttr])
		if srcKey == "" || dstKey == "" {
			logger.WARN(IUESRC001, "invalid connection source or destination", map[string]any{"src": srcIDStr, "dst": dstIDStr, "srcKey": srcKey, "dstKey": dstKey})
			continue
		}
		srcImgRect, srcOk := itemImgRects[srcKey]
		dstImgRect, dstOk := itemImgRects[dstKey]
		if !srcOk {
			logger.WARN(IUESRC002, "source item not found or not rendered", map[string]any{"src": srcIDStr, "key": srcKey})
			continue
		}
		if !dstOk {
			logger.WARN(IUESRC003, "destination item not found or not rendered", map[string]any{"dst": dstIDStr, "key": dstKey})
			continue
		}

		// Determine exit/entry side from image-center to image-center.
		srcCx := srcImgRect[0] + srcImgRect[2]/2
		srcCy := srcImgRect[1] + srcImgRect[3]/2
		dstCx := dstImgRect[0] + dstImgRect[2]/2
		dstCy := dstImgRect[1] + dstImgRect[3]/2
		srcSide, dstSide := connectionSide(srcCx, srcCy, dstCx, dstCy)
		bends := connectionBendPoints(conn)
		srcAnchor, hasSrcAnchor := connectionEndpointAnchor(conn, "src")
		dstAnchor, hasDstAnchor := connectionEndpointAnchor(conn, "dst")
		if hasSrcAnchor {
			srcSide = string(srcAnchor.side)
		} else if explicit, ok := connectionEndpointSide(conn, "src"); ok {
			srcSide = string(explicit)
		} else if len(bends) > 0 {
			srcSide = sideTowardPoint(srcImgRect, bends[0])
		}
		if hasDstAnchor {
			dstSide = string(dstAnchor.side)
		} else if explicit, ok := connectionEndpointSide(conn, "dst"); ok {
			dstSide = string(explicit)
		} else if len(bends) > 0 {
			dstSide = sideTowardPoint(dstImgRect, bends[len(bends)-1])
		}

		// Choose element: bottom edge → label text box; other edges → image element.
		var srcElemID string
		var srcRect [4]float64
		if srcSide == "bottom" {
			if lblRect, ok := itemLblRects[srcKey]; ok {
				srcRect = lblRect
				srcElemID = itemLblIDs[srcKey]
			} else {
				srcRect = srcImgRect
				srcElemID = itemImgIDs[srcKey]
			}
		} else {
			srcRect = srcImgRect
			srcElemID = itemImgIDs[srcKey]
		}

		var dstElemID string
		var dstRect [4]float64
		if dstSide == "bottom" {
			if lblRect, ok := itemLblRects[dstKey]; ok {
				dstRect = lblRect
				dstElemID = itemLblIDs[dstKey]
			} else {
				dstRect = dstImgRect
				dstElemID = itemImgIDs[dstKey]
			}
		} else {
			dstRect = dstImgRect
			dstElemID = itemImgIDs[dstKey]
		}

		srcFP := fixedPointForSide(srcSide)
		if hasSrcAnchor {
			srcFP = fixedPointForAnchor(srcAnchor)
		}
		dstFP := fixedPointForSide(dstSide)
		if hasDstAnchor {
			dstFP = fixedPointForAnchor(dstAnchor)
		}
		srcEdge := rectFixedPoint(srcRect, srcFP)
		dstEdge := rectFixedPoint(dstRect, dstFP)
		dx := dstEdge[0] - srcEdge[0]
		dy := dstEdge[1] - srcEdge[1]
		style := resolveConnectionStyle(conn)
		routePoints := excalidrawConnectionPoints(conn, srcRect, dstRect, srcSide, dstSide, style.Kind, obstacles, placed, routePaths)
		if style.Kind == "route" {
			routePaths[routePairKey(excalidrawRouteRequest(conn, srcRect, dstRect, srcSide, dstSide, style.Kind), false)] = append([]pt(nil), routePoints...)
		}
		placed = append(placed, toSegments(routePoints))
		minX, minY, maxX, maxY := srcEdge[0], srcEdge[1], srcEdge[0], srcEdge[1]
		points := make([][]float64, 0, len(routePoints))
		for _, p := range routePoints {
			minX = math.Min(minX, p.X)
			minY = math.Min(minY, p.Y)
			maxX = math.Max(maxX, p.X)
			maxY = math.Max(maxY, p.Y)
			points = append(points, []float64{p.X - srcEdge[0], p.Y - srcEdge[1]})
		}
		// seed は src/dst/index から決定論的に計算し、再生成しても描画ばらつきが出ないようにする。
		seed := stableConnectionSeed(srcKey, dstKey, i)
		connID := fmt.Sprintf("conn-%s-%s-%d", sanitizeElementID(srcKey), sanitizeElementID(dstKey), i)

		if style.Kind == "route" {
			for _, endpoint := range []struct {
				id   string
				edge [2]float64
				side string
				seed int
			}{{srcElemID, srcEdge, srcSide, seed}, {dstElemID, dstEdge, dstSide, seed + 1}} {
				key := endpoint.id + "|" + endpoint.side
				candidate := junctionCandidates[key]
				if candidate == nil {
					candidate = &junctionCandidate{edge: endpoint.edge, side: endpoint.side, color: style.Color, seed: endpoint.seed}
					junctionCandidates[key] = candidate
				}
				candidate.count++
			}
		}

		customData := map[string]any{
			"xaligoConnectorKind":           style.Kind,
			"xaligoConnectorStartArrowhead": style.StartArrowhead,
			"xaligoConnectorEndArrowhead":   style.EndArrowhead,
		}
		if hasSrcAnchor {
			customData["xaligoConnectorSrcAnchor"] = true
		}
		if hasDstAnchor {
			customData["xaligoConnectorDstAnchor"] = true
		}
		if bends := strings.TrimSpace(connectionBends(conn)); bends != "" {
			customData["xaligoConnectorBends"] = bends
		}
		if scale, ok := positiveFloatAttr(conn, "coordinate-scale", "scale"); ok {
			customData["xaligoConnectorScale"] = scale
		}
		if grid, ok := positiveFloatAttr(conn, "grid"); ok {
			customData["xaligoConnectorGrid"] = grid
		}
		if hasSrcAnchor {
			customData["xaligoConnectorStartAnchor"] = true
		}
		if hasDstAnchor {
			customData["xaligoConnectorEndAnchor"] = true
		}

		*elements = append(*elements, map[string]any{
			"id": connID, "type": "arrow",
			"x": srcEdge[0], "y": srcEdge[1],
			"width": math.Max(math.Abs(dx), maxX-minX), "height": math.Max(math.Abs(dy), maxY-minY),
			"angle":       0,
			"strokeColor": style.Color, "backgroundColor": "transparent",
			"fillStyle": "solid", "strokeWidth": style.Width, "strokeStyle": style.StrokeStyle,
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": map[string]any{"type": 2},
			"seed": seed, "version": 1, "versionNonce": seed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false, "frameId": nil,
			"points":             points,
			"lastCommittedPoint": nil,
			"startBinding": map[string]any{
				"elementId":  srcElemID,
				"focus":      0.0,
				"gap":        5.0,
				"fixedPoint": []float64{srcFP[0], srcFP[1]},
			},
			"endBinding": map[string]any{
				"elementId":  dstElemID,
				"focus":      0.0,
				"gap":        5.0,
				"fixedPoint": []float64{dstFP[0], dstFP[1]},
			},
			"startArrowhead":     style.ExcalidrawStartArrowhead,
			"endArrowhead":       style.ExcalidrawEndArrowhead,
			"endArrowheadSize":   "s",
			"startArrowheadSize": "s",
			"elbowed":            true,
			"customData":         customData,
		})

		// Register this arrow in boundMap for both endpoints.
		entry := map[string]any{"type": "arrow", "id": connID}
		boundMap[srcElemID] = append(boundMap[srcElemID], entry)
		boundMap[dstElemID] = append(boundMap[dstElemID], entry)
	}

	junctionKeys := make([]string, 0, len(junctionCandidates))
	for key, candidate := range junctionCandidates {
		if candidate.count >= 2 {
			junctionKeys = append(junctionKeys, key)
		}
	}
	sort.Strings(junctionKeys)
	for i, key := range junctionKeys {
		candidate := junctionCandidates[key]
		point := extendConnectionPoint(candidate.edge, candidate.side, 25)
		const diameter = 8.0
		*elements = append(*elements, map[string]any{
			"id": fmt.Sprintf("junction-%d", i), "type": "ellipse",
			"x": point[0] - diameter/2, "y": point[1] - diameter/2,
			"width": diameter, "height": diameter, "angle": 0,
			"strokeColor": candidate.color, "backgroundColor": candidate.color,
			"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
			"roughness": 0, "opacity": 100, "groupIds": []string{}, "roundness": nil,
			"seed": candidate.seed, "version": 1, "versionNonce": candidate.seed,
			"isDeleted": false, "boundElements": nil, "updated": updated,
			"link": nil, "locked": false, "frameId": nil,
			"customData": map[string]any{"xaligoJunction": true},
		})
	}

	// Second pass: write back boundElements into each referenced element so that
	// Excalidraw recognises the binding relationship.
	if len(boundMap) == 0 {
		return
	}
	for idx := range *elements {
		elem := (*elements)[idx]
		id, _ := elem["id"].(string)
		if entries, ok := boundMap[id]; ok {
			// Merge with any existing bound elements (e.g. text containerId refs).
			existing, _ := elem["boundElements"].([]map[string]any)
			elem["boundElements"] = append(existing, entries...)
			(*elements)[idx] = elem
		}
	}
}

func firstNonEmptyAttr(node *entity.Node, names ...string) string {
	if node == nil {
		return ""
	}
	for _, name := range names {
		if value := strings.TrimSpace(node.Attrs[name]); value != "" {
			return value
		}
	}
	return ""
}

func connectionEndpointSide(conn *entity.Node, endpoint string) (side, bool) {
	if conn == nil {
		return "", false
	}
	return normalizeConnectionSide(conn.Attrs[endpoint+"-side"])
}

func connectionEndpointAnchor(conn *entity.Node, endpoint string) (connectionAnchorSpec, bool) {
	if conn == nil {
		return connectionAnchorSpec{}, false
	}
	spec, ok, err := parseConnectionAnchorSpec(conn.Attrs[endpoint+"-side"], conn.Attrs[endpoint+"-anchor"])
	if err != nil || !ok || !spec.hasSlot {
		return connectionAnchorSpec{}, false
	}
	return spec, true
}

func excalidrawConnectionPoints(conn *entity.Node, srcRect, dstRect [4]float64, srcSide, dstSide, kind string, obstacles []rect, placed [][]segment, routePaths map[string][]pt) []pt {
	req := excalidrawRouteRequest(conn, srcRect, dstRect, srcSide, dstSide, kind)
	opt := defaultRouterOptions()
	local := filterObstacles(obstacles, req)
	path := routeOne(req, local, placed, opt)
	followedRoute := false
	if req.Kind == "traffic" {
		if base, ok := matchingRoutePath(req, routePaths); ok {
			path.Points = trafficAlongsideRoute(base, path.Points, opt.LaneGap)
			path.Points = separateExactOverlaps(path.Points, placed, local, opt)
			followedRoute = true
		} else {
			path.Points = separateExactOverlaps(path.Points, placed, local, opt)
		}
	} else if req.Kind != "route" {
		path.Points = separateExactOverlaps(path.Points, placed, local, opt)
	}
	visualMargin := math.Min(opt.LineMargin, opt.Clearance) / 2
	path.Points = separateObstacleHits(path.Points, placed, inflateRects(local, visualMargin), opt)
	if len(req.Bends) == 0 {
		path.Points = rerouteEndpointApproach(path.Points, req, opt)
	}
	path.Points = separatePinnedExactOverlaps(path.Points, placed, local, opt)
	if followedRoute {
		path.Points = restoreDestinationApproach(path.Points, req.DstSide, opt.Stub)
	}
	return enforceOrthogonalPolyline(path.Points)
}

func separatePinnedExactOverlaps(points []pt, placed [][]segment, obstacles []rect, opt routerOptions) []pt {
	if len(points) < 3 || len(placed) == 0 || opt.LaneGap <= 0 {
		return points
	}
	best := append([]pt(nil), points...)
	bestOverlap := exactOverlapLength(toSegments(best), placed)
	if bestOverlap <= eps {
		return best
	}
	inflated := inflateRects(obstacles, math.Min(opt.LineMargin, opt.Clearance)/2)
	bestScore := scorePath(best, inflated, placed, opt.LineMargin)
	for _, offset := range []float64{opt.LaneGap, -opt.LaneGap, opt.LaneGap * 2, -opt.LaneGap * 2} {
		shifted := offsetPolyline(points, offset)
		candidate := []pt{points[0]}
		appendTarget := shifted[1]
		candidate = appendOrthogonalLeg(candidate, points[0], appendTarget)
		if len(shifted) > 3 {
			candidate = append(candidate, shifted[2:len(shifted)-1]...)
		}
		candidate = appendOrthogonalLeg(candidate, shifted[len(shifted)-2], points[len(points)-1])
		candidate = simplifyRouteCandidate(candidate)
		candidate = enforceOrthogonalPolyline(candidate)
		if obstacleHitCount(candidate, inflated) > 0 {
			continue
		}
		overlap := exactOverlapLength(toSegments(candidate), placed)
		score := scorePath(candidate, inflated, placed, opt.LineMargin)
		if overlap < bestOverlap-eps || (math.Abs(overlap-bestOverlap) < eps && score < bestScore) {
			best, bestOverlap, bestScore = candidate, overlap, score
		}
	}
	return best
}

func excalidrawRouteRequest(conn *entity.Node, srcRect, dstRect [4]float64, srcSide, dstSide, kind string) routeRequest {
	src := rect{X: srcRect[0], Y: srcRect[1], W: srcRect[2], H: srcRect[3]}
	dst := rect{X: dstRect[0], Y: dstRect[1], W: dstRect[2], H: dstRect[3]}
	req := routeRequest{
		ID:      firstNonEmptyAttr(conn, "src") + "-" + firstNonEmptyAttr(conn, "dst"),
		Kind:    kind,
		Src:     src,
		Dst:     dst,
		SrcSide: side(srcSide),
		DstSide: side(dstSide),
		SrcGap:  5,
		DstGap:  5,
	}
	if anchor, ok := connectionEndpointAnchor(conn, "src"); ok {
		fp := fixedPointForAnchor(anchor)
		req.SrcAnchor = &pt{X: src.X + src.W*fp[0], Y: src.Y + src.H*fp[1]}
	}
	if anchor, ok := connectionEndpointAnchor(conn, "dst"); ok {
		fp := fixedPointForAnchor(anchor)
		req.DstAnchor = &pt{X: dst.X + dst.W*fp[0], Y: dst.Y + dst.H*fp[1]}
	}
	if scale, ok := positiveFloatAttr(conn, "coordinate-scale", "scale"); ok {
		req.Bends = parseConnectorBends(connectionBends(conn), scale)
	} else {
		req.Bends = parseConnectorBends(connectionBends(conn), 1)
	}
	if grid, ok := positiveFloatAttr(conn, "grid"); ok {
		req.Grid = grid
	}
	return req
}

func excalidrawRouteObstacles(elements []map[string]any) []rect {
	obstacles := make([]rect, 0)
	for _, el := range elements {
		custom, _ := el["customData"].(map[string]any)
		isAnchorContent, _ := custom["xaligoAnchorContent"].(bool)
		isHeader, _ := custom["xaligoGroupHeader"].(bool)
		isHeaderContent, _ := custom["xaligoGroupHeaderContent"].(bool)
		if !isAnchorContent && !isHeader && !isHeaderContent {
			continue
		}
		r, ok := elementRect(el)
		if !ok {
			continue
		}
		obstacles = append(obstacles, r)
	}
	return obstacles
}

func elementRect(el map[string]any) (rect, bool) {
	x, okX := el["x"].(float64)
	y, okY := el["y"].(float64)
	w, okW := el["width"].(float64)
	h, okH := el["height"].(float64)
	if !okX || !okY || !okW || !okH || w <= 0 || h <= 0 {
		return rect{}, false
	}
	return rect{X: x, Y: y, W: w, H: h}, true
}

func connectionBends(conn *entity.Node) string {
	if conn == nil {
		return ""
	}
	points := connectionChildBends(conn)
	if len(points) > 0 {
		return strings.Join(points, " ")
	}
	return firstNonEmptyAttr(conn, "bends", "points", "via")
}

func connectionChildBends(node *entity.Node) []string {
	if node == nil {
		return nil
	}
	points := []string{}
	for _, child := range node.Children {
		switch strings.ToLower(strings.TrimSpace(child.Tag)) {
		case "bend", "point", "via", "waypoint":
			if point, ok := connectionPointString(child); ok {
				points = append(points, point)
			}
		case "bends", "points", "path":
			points = append(points, connectionChildBends(child)...)
		}
	}
	return points
}

func connectionPointString(node *entity.Node) (string, bool) {
	x, xOK := floatAttr(node, "x")
	y, yOK := floatAttr(node, "y")
	if xOK && yOK {
		return fmtFloat(x) + "," + fmtFloat(y), true
	}
	parts := strings.Split(strings.TrimSpace(node.Text), ",")
	if len(parts) != 2 {
		return "", false
	}
	x, xErr := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	y, yErr := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if xErr != nil || yErr != nil {
		return "", false
	}
	return fmtFloat(x) + "," + fmtFloat(y), true
}

func floatAttr(node *entity.Node, name string) (float64, bool) {
	if node == nil {
		return 0, false
	}
	value := strings.TrimSpace(node.Attrs[name])
	if value == "" {
		return 0, false
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}
	return parsed, true
}

func positiveFloatAttr(node *entity.Node, names ...string) (float64, bool) {
	if node == nil {
		return 0, false
	}
	for _, name := range names {
		value := strings.TrimSpace(node.Attrs[name])
		if value == "" {
			continue
		}
		parsed, err := strconv.ParseFloat(value, 64)
		if err == nil && parsed > 0 {
			return parsed, true
		}
	}
	return 0, false
}

func stableConnectionSeed(srcKey, dstKey string, index int) int {
	seed := 2166136261
	for _, r := range srcKey + "|" + dstKey + "|" + strconv.Itoa(index) {
		seed ^= int(r)
		seed *= 16777619
	}
	if seed < 0 {
		seed = -seed
	}
	return seed%99999999 + 1
}

func sanitizeElementID(s string) string {
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			b.WriteRune(r)
			continue
		}
		b.WriteByte('-')
	}
	out := b.String()
	if out == "" {
		return "endpoint"
	}
	return out
}

func extendConnectionPoint(point [2]float64, side string, distance float64) [2]float64 {
	switch side {
	case "top":
		point[1] -= distance
	case "bottom":
		point[1] += distance
	case "left":
		point[0] -= distance
	default:
		point[0] += distance
	}
	return point
}

type resolvedConnectionStyle struct {
	Kind                     string
	Color                    string
	Width                    float64
	StrokeStyle              string
	StartArrowhead           string
	EndArrowhead             string
	ExcalidrawStartArrowhead any
	ExcalidrawEndArrowhead   any
}

func connectionKind(conn *entity.Node) string {
	kind := strings.ToLower(strings.TrimSpace(conn.Attrs["kind"]))
	if kind == "route" || kind == "traffic" {
		return kind
	}
	return "connection"
}

func connectionKindPriority(kind string) int {
	switch kind {
	case "route":
		return 0
	case "traffic":
		return 2
	default:
		return 1
	}
}

func resolveConnectionStyle(conn *entity.Node) resolvedConnectionStyle {
	kind := connectionKind(conn)
	style := resolvedConnectionStyle{
		Kind: kind, Color: "#1e1e1e", Width: 1, StrokeStyle: "solid",
		StartArrowhead: "none", EndArrowhead: "stealth",
		ExcalidrawStartArrowhead: nil, ExcalidrawEndArrowhead: "arrow",
	}
	switch kind {
	case "route":
		style.Color = "#64748b"
		style.EndArrowhead = "none"
		style.ExcalidrawEndArrowhead = nil
	case "traffic":
		style.Color = "#2563eb"
	}
	if color := strings.TrimSpace(conn.Attrs["color"]); color != "" {
		style.Color = color
	}
	widthValue := strings.TrimSpace(conn.Attrs["stroke-width"])
	if widthValue == "" {
		widthValue = strings.TrimSpace(conn.Attrs["width"])
	}
	if width, err := strconv.ParseFloat(widthValue, 64); err == nil && width > 0 {
		style.Width = width
	}
	if strokeStyle := strings.ToLower(strings.TrimSpace(conn.Attrs["stroke-style"])); strokeStyle == "solid" || strokeStyle == "dashed" || strokeStyle == "dotted" {
		style.StrokeStyle = strokeStyle
	}
	endArrowhead := strings.TrimSpace(conn.Attrs["end-arrowhead"])
	if endArrowhead == "" {
		endArrowhead = strings.TrimSpace(conn.Attrs["arrowhead"])
	}
	style.StartArrowhead, style.ExcalidrawStartArrowhead = resolveArrowhead(conn.Attrs["start-arrowhead"], style.StartArrowhead, style.ExcalidrawStartArrowhead)
	style.EndArrowhead, style.ExcalidrawEndArrowhead = resolveArrowhead(endArrowhead, style.EndArrowhead, style.ExcalidrawEndArrowhead)
	return style
}

func resolveArrowhead(value, current string, currentExcalidraw any) (string, any) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "none":
		return "none", nil
	case "arrow", "triangle", "diamond":
		value = strings.ToLower(strings.TrimSpace(value))
		return value, value
	case "stealth":
		return "stealth", "arrow"
	case "oval":
		return "oval", "dot"
	default:
		return current, currentExcalidraw
	}
}
