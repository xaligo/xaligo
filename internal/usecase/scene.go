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

	"github.com/ryo-arima/xaligo/internal/config"
	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
	"github.com/ryo-arima/xaligo/internal/share"
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
	var connections []*entity.Node
	for _, child := range doc.Root.Children {
		if child.Tag == "connection" {
			logger.DEBUG(IURBS003, "branch connection node", map[string]any{"tag": child.Tag})
			connections = append(connections, child)
		}
	}
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
)

var svgTintColorRE = regexp.MustCompile(`(?i)#[0-9a-f]{3,8}|currentColor`)

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
	// catalog ID → bounding rect (x,y,w,h) and element ID of the image / label elements.
	// Populated during renderItemGrid → renderIconAt, used for edge-based connections.
	itemImgRects := map[int][4]float64{}
	itemLblRects := map[int][4]float64{}
	itemImgIDs := map[int]string{}
	itemLblIDs := map[int]string{}

	walk(root, &elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, r, root, itemGroups, ancestorBoxes, deps)
	for ancID, items := range itemGroups {
		renderItemGrid(items, ancestorBoxes[ancID], &elements, files, catalogCSV, projectRoot, fsys, itemIconSize, r, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps)
	}
	renderConnections(connections, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, &elements, r)
	elements = orderGroupHeaderLayers(elements)

	out := file{
		Type:     "excalidraw",
		Version:  2,
		Source:   "https://github.com/ryo-arima/xaligo",
		Elements: elements,
		AppState: map[string]any{
			"gridSize":            20,
			"viewBackgroundColor": "#ffffff",
		},
		Files: files,
	}
	return json.MarshalIndent(out, "", "  ")
}

// orderGroupHeaderLayers keeps title tags above every group border while
// preserving connector priority. Header icons and labels are placed last so
// neither nested borders nor connectors can make the title unreadable.
func orderGroupHeaderLayers(elements []map[string]any) []map[string]any {
	base := make([]map[string]any, 0, len(elements))
	headShapes := make([]map[string]any, 0)
	connectors := make([]map[string]any, 0)
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
		typ, _ := el["type"].(string)
		if typ == "arrow" || typ == "line" {
			connectors = append(connectors, el)
			continue
		}
		base = append(base, el)
	}
	ordered := append(base, headShapes...)
	ordered = append(ordered, connectors...)
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

func walk(b *entity.Box, elements *[]map[string]any, files map[string]any, svgGroupDir string, catalogCSV string, projectRoot string, fsys fs.FS, r *rand.Rand, visibleAncestor *entity.Box, itemGroups map[string][]*entity.Box, ancestorBoxes map[string]*entity.Box, deps SceneDependencies) {
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

	if b.Tag != "frame" && (b.W < MinBoxWidth || b.H < MinBoxHeight) {
		logger.WARN(IUESW001, "skipping too small element", map[string]any{"label": b.Label, "tag": b.Tag, "width": b.W, "height": b.H, "minWidth": MinBoxWidth, "minHeight": MinBoxHeight})
		// 子の item も同じ visibleAncestor に結び付けて収集
		for _, c := range b.Children {
			if IsItemLike(c.Tag) {
				key := visibleAncestor.ID
				itemGroups[key] = append(itemGroups[key], c)
				ancestorBoxes[key] = visibleAncestor
			} else {
				walk(c, elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, r, visibleAncestor, itemGroups, ancestorBoxes, deps)
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
			*elements = append(*elements, map[string]any{
				"id": rectID, "type": "rectangle",
				"x": b.X, "y": b.Y, "width": b.W, "height": b.H,
				"angle":       0,
				"strokeColor": genStroke, "backgroundColor": "transparent",
				"fillStyle": "hachure", "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": map[string]any{"type": 3},
				"seed": r.Intn(99999999), "version": 1,
				"versionNonce":  r.Intn(99999999),
				"isDeleted":     false,
				"boundElements": []map[string]any{{"type": "text", "id": textID}},
				"updated":       updated, "link": nil, "locked": false,
			})
			*elements = append(*elements, map[string]any{
				"id": textID, "type": "text",
				"x": b.X + 12, "y": b.Y + 12,
				// fontFamily=1 (Virgil 20px): ~10px/rune
				"width": textWidth(b.Label, 10.0), "height": 24,
				"angle":       0,
				"strokeColor": "#1e1e1e", "backgroundColor": "transparent",
				"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
				"roughness": 0, "opacity": 100,
				"groupIds": []string{}, "roundness": nil,
				"seed": r.Intn(99999999), "version": 1,
				"versionNonce": r.Intn(99999999),
				"isDeleted":    false, "boundElements": nil,
				"updated": updated, "link": nil, "locked": false,
				"text": b.Label, "fontSize": 20, "fontFamily": 1,
				"textAlign": "left", "verticalAlign": "top",
				"containerId": rectID, "originalText": b.Label, "lineHeight": 1.2,
			})
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
		walk(c, elements, files, svgGroupDir, catalogCSV, projectRoot, fsys, r, nextVisible, itemGroups, ancestorBoxes, deps)
	}
}

// isLayoutTag reports whether a tag is a pure layout container
// (<row>, <col>, <container>) that should not render any visible border or label.
func isLayoutTag(tag string) bool {
	return tag == "row" || tag == "col" || tag == "container" || IsBlank(tag)
}

// textWidth estimates the rendered width of a string in pixels.
// charW: approximate pixel width per rune (font-specific).
func textWidth(s string, charW float64) float64 {
	return math.Ceil(float64(len([]rune(s)))*charW) + 8
}

func itemLabelHeight(label string) float64 {
	lines := 1
	for _, line := range strings.Split(label, "\n") {
		lineRunes := len([]rune(line))
		wrapped := int(math.Ceil(float64(lineRunes) * itemLabelCharW / itemLabelW))
		if wrapped < 1 {
			wrapped = 1
		}
		lines += wrapped - 1
	}
	lineH := itemLabelFontPx * 1.25
	return math.Max(itemLabelH, math.Ceil(float64(lines)*lineH))
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
func renderItemGrid(items []*entity.Box, ancestor *entity.Box, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, maxSize float64, r *rand.Rand, itemImgRects map[int][4]float64, itemLblRects map[int][4]float64, itemImgIDs map[int]string, itemLblIDs map[int]string, abbrevMap map[int]string, deps SceneDependencies) {
	if catalogCSV == "" || len(items) == 0 || ancestor == nil {
		return
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
		return
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
		renderIconAt(item.ID, item.Attrs["id"], iconX, iconY, iconSize, elements, files, catalogCSV, projectRoot, fsys, r, itemImgRects, itemLblRects, itemImgIDs, itemLblIDs, abbrevMap, deps)
	}
}

func groupHeaderHeightForItems(_ *entity.Box) float64 {
	headerH := float64(groupTextHeight + groupHeaderTextPadY*2)
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
// and element ID of the image and label elements, keyed by the catalog integer ID.
func renderIconAt(boxID, idAttr string, iconX, iconY, iconSize float64, elements *[]map[string]any, files map[string]any, catalogCSV string, projectRoot string, fsys fs.FS, r *rand.Rand, itemImgRects map[int][4]float64, itemLblRects map[int][4]float64, itemImgIDs map[int]string, itemLblIDs map[int]string, abbrevMap map[int]string, deps SceneDependencies) {
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
	// Record bounding rects and element IDs for edge-based connection arrows.
	if itemImgRects != nil {
		itemImgRects[id] = [4]float64{iconX, iconY, iconSize, iconSize}
		itemImgIDs[id] = iconID
	}
	*elements = append(*elements, map[string]any{
		"id": iconID, "type": "image",
		"x": iconX, "y": iconY,
		"width": iconSize, "height": iconSize,
		"fileId": fid, "status": "saved",
		"scale":       []int{1, 1},
		"strokeColor": "transparent", "backgroundColor": deps.ExcalidrawRepository.SVGBGColor(ce.DataURL),
		"fillStyle": "solid", "strokeWidth": 1, "strokeStyle": "solid",
		"roughness": 0, "opacity": 100, "angle": 0,
		"groupIds": []string{}, "roundness": nil,
		"seed": seed, "version": 1, "versionNonce": seed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false, "frameId": nil,
	})
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
	// Record label bounding rect for bottom-side connection binding.
	if itemLblRects != nil {
		itemLblRects[id] = [4]float64{labelX, labelY, itemLabelW, labelH}
		itemLblIDs[id] = iconID + "-lbl"
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
		"groupIds": []string{}, "roundness": nil,
		"seed": textSeed, "version": 1, "versionNonce": textSeed,
		"isDeleted": false, "boundElements": nil,
		"updated": updated, "link": nil, "locked": false, "frameId": nil,
		"text": label, "rawText": label, "originalText": label,
		"fontSize": itemLabelFontPx, "fontFamily": 4,
		"textAlign": "center", "verticalAlign": "top",
		"containerId": nil, "lineHeight": 1.25,
	})
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

// renderConnections generates elbowed arrow elements for each <connection> node and
// updates the boundElements of the bound source/destination elements — required by
// Excalidraw so that the application recognises the binding relationship.
//
// src/dst are catalog integer IDs; the corresponding item rects and element IDs
// must already be populated in itemImgRects/itemLblRects/itemImgIDs/itemLblIDs by renderIconAt.
// Arrows start/end at the actual element edge; when the connection exits/enters from the
// bottom the label text element is used instead of the image element.
func renderConnections(connections []*entity.Node, itemImgRects map[int][4]float64, itemLblRects map[int][4]float64, itemImgIDs map[int]string, itemLblIDs map[int]string, elements *[]map[string]any, r *rand.Rand) {
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

	for i, conn := range orderedConnections {
		srcIDStr := strings.TrimSpace(conn.Attrs["src"])
		dstIDStr := strings.TrimSpace(conn.Attrs["dst"])
		srcID, err1 := strconv.Atoi(srcIDStr)
		dstID, err2 := strconv.Atoi(dstIDStr)
		if err1 != nil || err2 != nil {
			logger.WARN(IUESRC001, "invalid connection source or destination", map[string]any{"src": srcIDStr, "dst": dstIDStr, "srcError": err1, "dstError": err2})
			continue
		}
		srcImgRect, srcOk := itemImgRects[srcID]
		dstImgRect, dstOk := itemImgRects[dstID]
		if !srcOk {
			logger.WARN(IUESRC002, "source item not found or not rendered", map[string]any{"src": srcID})
			continue
		}
		if !dstOk {
			logger.WARN(IUESRC003, "destination item not found or not rendered", map[string]any{"dst": dstID})
			continue
		}

		// Determine exit/entry side from image-center to image-center.
		srcCx := srcImgRect[0] + srcImgRect[2]/2
		srcCy := srcImgRect[1] + srcImgRect[3]/2
		dstCx := dstImgRect[0] + dstImgRect[2]/2
		dstCy := dstImgRect[1] + dstImgRect[3]/2
		srcSide, dstSide := connectionSide(srcCx, srcCy, dstCx, dstCy)

		// Choose element: bottom edge → label text box; other edges → image element.
		var srcElemID string
		var srcRect [4]float64
		if srcSide == "bottom" {
			if lblRect, ok := itemLblRects[srcID]; ok {
				srcRect = lblRect
				srcElemID = itemLblIDs[srcID]
			} else {
				srcRect = srcImgRect
				srcElemID = itemImgIDs[srcID]
			}
		} else {
			srcRect = srcImgRect
			srcElemID = itemImgIDs[srcID]
		}

		var dstElemID string
		var dstRect [4]float64
		if dstSide == "bottom" {
			if lblRect, ok := itemLblRects[dstID]; ok {
				dstRect = lblRect
				dstElemID = itemLblIDs[dstID]
			} else {
				dstRect = dstImgRect
				dstElemID = itemImgIDs[dstID]
			}
		} else {
			dstRect = dstImgRect
			dstElemID = itemImgIDs[dstID]
		}

		srcEdge := rectEdgePoint(srcRect, srcSide)
		dstEdge := rectEdgePoint(dstRect, dstSide)
		dx := dstEdge[0] - srcEdge[0]
		dy := dstEdge[1] - srcEdge[1]

		srcFP := fixedPointForSide(srcSide)
		dstFP := fixedPointForSide(dstSide)

		// seed は src/dst/index から決定論的に計算し、再生成しても描画ばらつきが出ないようにする。
		seed := srcID*1_000_000 + dstID*1_000 + i + 1
		connID := fmt.Sprintf("conn-%d-%d-%d", srcID, dstID, i)

		style := resolveConnectionStyle(conn)
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

		// arrowhead-size 属性: "s" / "m" / "l"。未指定時は最小 "s" を使用する。
		ahSize := strings.TrimSpace(conn.Attrs["arrowhead-size"])
		if ahSize != "s" && ahSize != "m" && ahSize != "l" {
			ahSize = "s"
		}

		*elements = append(*elements, map[string]any{
			"id": connID, "type": "arrow",
			"x": srcEdge[0], "y": srcEdge[1],
			"width": math.Abs(dx), "height": math.Abs(dy),
			"angle":       0,
			"strokeColor": style.Color, "backgroundColor": "transparent",
			"fillStyle": "solid", "strokeWidth": style.Width, "strokeStyle": style.StrokeStyle,
			"roughness": 0, "opacity": 100,
			"groupIds": []string{}, "roundness": map[string]any{"type": 2},
			"seed": seed, "version": 1, "versionNonce": seed,
			"isDeleted": false, "boundElements": nil,
			"updated": updated, "link": nil, "locked": false, "frameId": nil,
			"points":             [][]float64{{0, 0}, {dx, dy}},
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
			"endArrowheadSize":   ahSize,
			"startArrowheadSize": ahSize,
			"elbowed":            true,
			"customData": map[string]any{
				"xaligoConnectorKind":           style.Kind,
				"xaligoConnectorStartArrowhead": style.StartArrowhead,
				"xaligoConnectorEndArrowhead":   style.EndArrowhead,
			},
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
