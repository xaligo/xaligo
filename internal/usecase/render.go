package usecase

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/share"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
	v2usecase "github.com/xaligo/xaligo/internal/usecase/v2"
	"github.com/yuin/goldmark"
	goldmarkast "github.com/yuin/goldmark/ast"
	goldmarktext "github.com/yuin/goldmark/text"
)

// RenderUsecase owns format dispatch and the shared render pipeline.
type RenderUsecase interface {
	ValidateRenderOptions(entity.RenderOptions) error
	Render(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderArtifacts(context.Context, []byte, entity.RenderOptions) ([]entity.RenderArtifact, error)
	BuildScene(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderSVG(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderPPTX(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderTerminal(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	BuildPPTXPlan(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	NewPreviewRepository(string, entity.PreviewOptions) (repository.PreviewRepository, error)
}

type renderUsecase struct {
	sceneRepository      repository.SceneRepository
	xaligoRepository     repository.XaligoRepository
	powerpointRepository repository.PowerpointRepository
	svgRepository        repository.SVGRepository
	terminalRepository   repository.TerminalRepository
	v2Frontend           v2usecase.FrontendUsecase
	v2Engine             v2usecase.EngineUsecase
}

func NewRenderUsecase(
	sceneRepository repository.SceneRepository,
	xaligoRepository repository.XaligoRepository,
	powerpointRepository repository.PowerpointRepository,
	svgRepository repository.SVGRepository,
	terminalRepository repository.TerminalRepository,
) RenderUsecase {
	return &renderUsecase{
		sceneRepository:      sceneRepository,
		xaligoRepository:     xaligoRepository,
		powerpointRepository: powerpointRepository,
		svgRepository:        svgRepository,
		terminalRepository:   terminalRepository,
		v2Frontend:           v2usecase.NewFrontendUsecase(),
		v2Engine:             v2usecase.NewEngineUsecase(),
	}
}

const (
	ModeStandard = v1engine.ModeStandardV1EngineOptionRender
	ModeNetwork  = v1engine.ModeNetworkV1EngineOptionRender
	ModeAWS      = v1engine.ModeAWSV1EngineOptionRender

	FormatSVG                    = v1engine.FormatSVGV1EngineOptionRender
	FormatPPTX                   = v1engine.FormatPPTXV1EngineOptionRender
	FormatTerminal entity.Format = "terminal"

	SeverityError   = v1engine.SeverityErrorV1EngineOptionRender
	SeverityWarning = v1engine.SeverityWarningV1EngineOptionRender
)

var (
	logger   = share.DefaultLogger()
	IURR001  = share.NewMCode("IURR-001", "Render context check failed")
	IURR002  = share.NewMCode("IURR-002", "Render validate render options failed")
	IURR003  = share.NewMCode("IURR-003", "Render default format branch")
	IURR005  = share.NewMCode("IURR-005", "Render SVG branch")
	IURR006  = share.NewMCode("IURR-006", "Render PPTX branch")
	IURR010  = share.NewMCode("IURR-010", "Render terminal branch")
	IURR009  = share.NewMCode("IURR-009", "Render unknown format branch")
	IURSO001 = share.NewMCode("IURSO-001", "Service options no services CSV branch")
	IURSO002 = share.NewMCode("IURSO-002", "Service options services CSV branch")
	IURSO003 = share.NewMCode("IURSO-003", "Service options read services CSV failed")
	IURSO005 = share.NewMCode("IURSO-005", "Service options legend validation failed")
	IURBS001 = share.NewMCode("IURBS-001", "Build scene context check failed")
	IURBS002 = share.NewMCode("IURBS-002", "Build scene validate render options failed")
	IURBS004 = share.NewMCode("IURBS-004", "Build scene embedded assets branch")
	IURBS005 = share.NewMCode("IURBS-005", "Build scene default embedded item size branch")
	IURBS006 = share.NewMCode("IURBS-006", "Build scene native assets branch")
	IURBS007 = share.NewMCode("IURBS-007", "Build scene parse DSL failed")
	IURBS008 = share.NewMCode("IURBS-008", "Build scene build layout failed")
	IURBS009 = share.NewMCode("IURBS-009", "Build scene service options failed")
	IURBS010 = share.NewMCode("IURBS-010", "Build scene build JSON failed")
	IURBS011 = share.NewMCode("IURBS-011", "Build scene apply theme failed")
)

func (rcvr *renderUsecase) ValidateRenderOptions(opts entity.RenderOptions) error {
	return ValidateRenderOptions(opts)
}

func validateRenderOptions(opts entity.RenderOptions) error {
	if entity.Format(strings.ToLower(strings.TrimSpace(string(opts.Format)))) == FormatTerminal {
		if err := validateTerminalRenderOptions(opts); err != nil {
			return err
		}
		opts.Format = FormatSVG
	}
	return v1engine.ValidateRenderOptionsV1EngineOptionRender(opts)
}

func validateTerminalRenderOptions(opts entity.RenderOptions) error {
	if opts.TerminalStyle != "" && opts.TerminalStyle != entity.TerminalStyleUnicode && opts.TerminalStyle != entity.TerminalStyleASCII {
		return fmt.Errorf("unknown terminal style %q", opts.TerminalStyle)
	}
	if opts.TerminalLayout != "" && opts.TerminalLayout != entity.TerminalLayoutDiagram && opts.TerminalLayout != entity.TerminalLayoutSemantic && opts.TerminalLayout != entity.TerminalLayoutHybrid {
		return fmt.Errorf("unknown terminal layout %q", opts.TerminalLayout)
	}
	if opts.TerminalDetail != "" && opts.TerminalDetail != entity.TerminalDetailCompact && opts.TerminalDetail != entity.TerminalDetailNormal && opts.TerminalDetail != entity.TerminalDetailFull {
		return fmt.Errorf("unknown terminal detail %q", opts.TerminalDetail)
	}
	if opts.TerminalColor != "" && opts.TerminalColor != entity.TerminalColorAuto && opts.TerminalColor != entity.TerminalColorAlways && opts.TerminalColor != entity.TerminalColorNever {
		return fmt.Errorf("unknown terminal color %q", opts.TerminalColor)
	}
	if opts.TerminalIcons != "" && opts.TerminalIcons != entity.TerminalIconsLabel && opts.TerminalIcons != entity.TerminalIconsSymbol && opts.TerminalIcons != entity.TerminalIconsNone {
		return fmt.Errorf("unknown terminal icons %q", opts.TerminalIcons)
	}
	if opts.TerminalWidth < 0 || opts.TerminalHeight < 0 {
		return fmt.Errorf("terminal dimensions must not be negative")
	}
	return nil
}

func (rcvr *renderUsecase) Render(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURR001, "context check failed", map[string]any{"error": err})
		return nil, err
	}
	if err := rcvr.ValidateRenderOptions(opts); err != nil {
		logger.ERROR(IURR002, "validate render options failed", map[string]any{"format": opts.Format, "error": err})
		return nil, err
	}
	format := entity.Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	if format == "" {
		logger.DEBUG(IURR003, "branch default format")
		format = FormatSVG
	}
	opts.Format = format
	switch format {
	case FormatSVG:
		logger.DEBUG(IURR005, "branch svg")
		return rcvr.RenderSVG(ctx, input, opts)
	case FormatPPTX:
		logger.DEBUG(IURR006, "branch pptx")
		return rcvr.RenderPPTX(ctx, input, opts)
	case FormatTerminal:
		logger.DEBUG(IURR010, "branch terminal")
		return rcvr.RenderTerminal(ctx, input, opts)
	default:
		logger.ERROR(IURR009, "branch unknown format", map[string]any{"format": format})
		return nil, fmt.Errorf("unknown render format %q", format)
	}
}

func (rcvr *renderUsecase) RenderTerminal(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	version, err := renderDocumentVersion(input)
	if err != nil {
		return nil, err
	}
	if version != "2" {
		return nil, fmt.Errorf("terminal output is available only for V2 documents")
	}
	spec, _, err := rcvr.v2Frontend.Lower(input)
	if err != nil {
		return nil, fmt.Errorf("lower V2 document: %w", err)
	}
	resolved, err := rcvr.v2Engine.Resolve(ctx, spec)
	if err != nil {
		return nil, fmt.Errorf("resolve V2 document: %w", err)
	}
	data, err := rcvr.terminalRepository.Render(resolved, opts)
	if err != nil {
		return nil, fmt.Errorf("render V2 terminal output: %w", err)
	}
	return data, nil
}

var IURBS012 = share.NewMCode("IURBS-012", "Build presentation scene completed")

// BuildScene exposes the transitional V1 presentation scene to internal
// callers and tests. It is not a supported output format; SVG and PPTX are the
// only public render projections.
func (rcvr *renderUsecase) BuildScene(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := rcvr.buildScene(ctx, input, opts)
	if err == nil {
		logger.DEBUG(IURBS012, "completed")
	}
	return scene, err
}

var (
	IURRS001 = share.NewMCode("IURRS-001", "Render SVG build plan failed")
)

func (rcvr *renderUsecase) RenderSVG(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	artifacts, err := rcvr.RenderArtifacts(ctx, input, opts)
	if err != nil {
		return nil, err
	}
	if len(artifacts) != 1 {
		return nil, fmt.Errorf("SVG render produced %d frame files; use RenderArtifacts or set CombineFrames", len(artifacts))
	}
	return artifacts[0].Data, nil
}

// RenderArtifacts renders one SVG artifact per XAL frame. Formats represented
// by one container file continue to use Render.
func (rcvr *renderUsecase) RenderArtifacts(ctx context.Context, input []byte, opts entity.RenderOptions) ([]entity.RenderArtifact, error) {
	format := entity.Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	if format == "" {
		format = FormatSVG
	}
	if format != FormatSVG {
		return nil, fmt.Errorf("render artifacts is only available for SVG, got %q", format)
	}
	opts.Format = FormatSVG
	version, err := renderDocumentVersion(input)
	if err != nil {
		return nil, err
	}
	if version == "2" {
		spec, _, err := rcvr.v2Frontend.Lower(input)
		if err != nil {
			return nil, fmt.Errorf("lower V2 document: %w", err)
		}
		data, err := rcvr.v2Engine.RenderSVG(ctx, spec)
		if err != nil {
			return nil, fmt.Errorf("render V2 SVG: %w", err)
		}
		data, err = embedV2CatalogIcons(data, opts)
		if err != nil {
			return nil, fmt.Errorf("embed V2 SVG icons: %w", err)
		}
		return []entity.RenderArtifact{{ID: "v2", Data: data}}, nil
	}
	document, err := rcvr.buildDocumentPlan(ctx, input, opts, false)
	if err != nil {
		logger.ERROR(IURRS001, "build document plan failed", map[string]any{"error": err})
		return nil, fmt.Errorf("build SVG document plan: %w", err)
	}
	artifacts := make([]entity.RenderArtifact, 0, len(document.Pages))
	for _, page := range document.Pages {
		data, renderErr := rcvr.svgRepository.Render(entity.Plan{
			Slide: page.Slide, Ops: page.Ops, Legend: document.Legend,
		}, opts.PxPerInch, opts.SVGLegendPosition)
		if renderErr != nil {
			return nil, fmt.Errorf("render SVG frame %q: %w", page.ID, renderErr)
		}
		artifacts = append(artifacts, entity.RenderArtifact{ID: page.ID, Data: data})
	}
	return artifacts, nil
}

var (
	v2CatalogRectPattern = regexp.MustCompile(`<rect\b[^>]*\bdata-icon="catalog:([0-9]+)"[^>]*/>`)
	v2SVGAttrPattern     = regexp.MustCompile(`(?:^|\s)(x|y|width|height)="([^"]+)"`)
	v2CatalogCache       sync.Map
)

func embedV2CatalogIcons(svg []byte, opts entity.RenderOptions) ([]byte, error) {
	matches := v2CatalogRectPattern.FindAllSubmatch(svg, -1)
	if len(matches) == 0 {
		return svg, nil
	}
	catalog, err := readV2Catalog(opts)
	if err != nil {
		return nil, err
	}
	var images strings.Builder
	for _, match := range matches {
		id, _ := strconv.Atoi(string(match[1]))
		dataURL := catalog[id]
		if dataURL == "" {
			continue
		}
		attrs := map[string]float64{}
		for _, attr := range v2SVGAttrPattern.FindAllSubmatch(match[0], -1) {
			attrs[string(attr[1])], _ = strconv.ParseFloat(string(attr[2]), 64)
		}
		size := min(40.0, attrs["width"], attrs["height"])
		if size <= 0 {
			continue
		}
		x := attrs["x"] + (attrs["width"]-size)/2
		y := attrs["y"] + (attrs["height"]-size)/2
		fmt.Fprintf(&images, `<image x="%g" y="%g" width="%g" height="%g" preserveAspectRatio="xMidYMid meet" href="%s"/>`, x, y, size, size, html.EscapeString(dataURL))
	}
	if images.Len() == 0 {
		return svg, nil
	}
	end := bytes.LastIndex(svg, []byte("</svg>"))
	if end < 0 {
		return nil, fmt.Errorf("Rust SVG has no closing element")
	}
	output := make([]byte, 0, len(svg)+images.Len())
	output = append(output, svg[:end]...)
	output = append(output, images.String()...)
	output = append(output, svg[end:]...)
	return output, nil
}

func readV2Catalog(opts entity.RenderOptions) (map[int]string, error) {
	var data []byte
	var err error
	cacheKey := ""
	if opts.Assets != nil && opts.Assets.FS != nil {
		data, err = fs.ReadFile(opts.Assets.FS, opts.Assets.CatalogCSV)
	} else {
		cacheKey = config.New().SvcCatalogCSV
		if cached, ok := v2CatalogCache.Load(cacheKey); ok {
			return cached.(map[int]string), nil
		}
		data, err = os.ReadFile(cacheKey)
	}
	if err != nil {
		return nil, fmt.Errorf("read service catalog: %w", err)
	}
	records, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parse service catalog: %w", err)
	}
	catalog := make(map[int]string, len(records))
	for _, record := range records[1:] {
		if len(record) < 6 {
			continue
		}
		id, parseErr := strconv.Atoi(strings.TrimSpace(record[0]))
		if parseErr == nil {
			catalog[id] = strings.TrimSpace(record[5])
		}
	}
	if cacheKey != "" {
		v2CatalogCache.Store(cacheKey, catalog)
	}
	return catalog, nil
}

var (
	IURBPP001 = share.NewMCode("IURBPP-001", "Build PPTX plan build scene failed")
	IURBPP002 = share.NewMCode("IURBPP-002", "Build PPTX plan build plan failed")
	IURRP001  = share.NewMCode("IURRP-001", "Render PPTX build plan failed")
	IURRP003  = share.NewMCode("IURRP-003", "Render PPTX export failed")
	IURRP004  = share.NewMCode("IURRP-004", "Render PPTX context check failed")
)

func (rcvr *renderUsecase) buildDocumentPlan(ctx context.Context, input []byte, opts entity.RenderOptions, uniformPages bool) (entity.DocumentPlan, error) {
	version, err := renderDocumentVersion(input)
	if err != nil {
		return entity.DocumentPlan{}, err
	}
	if version == "2" {
		spec, _, err := rcvr.v2Frontend.Lower(input)
		if err != nil {
			return entity.DocumentPlan{}, fmt.Errorf("lower V2 document: %w", err)
		}
		resolved, err := rcvr.v2Engine.Resolve(ctx, spec)
		if err != nil {
			return entity.DocumentPlan{}, fmt.Errorf("resolve V2 document: %w", err)
		}
		icons, err := resolvedV2CatalogIcons(resolved, opts)
		if err != nil {
			return entity.DocumentPlan{}, fmt.Errorf("resolve V2 plan icons: %w", err)
		}
		document, err := v2usecase.BuildDocumentPlanWithIcons(resolved, opts.PxPerInch, icons)
		if err != nil {
			return entity.DocumentPlan{}, fmt.Errorf("build V2 document plan: %w", err)
		}
		if uniformPages {
			v1engine.NormalizeDocumentPageSizesV1EnginePlanDocument(&document)
		}
		return document, nil
	}
	scene, entries, err := rcvr.buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURBPP001, "build scene failed", map[string]any{"error": err})
		return entity.DocumentPlan{}, err
	}
	planJSON, err := v1engine.BuildDocumentPlanJSONV1EnginePlanDocument(
		string(scene), v1engine.ResolvePlanOptionsV1EngineOptionPlan(opts, entries), opts.CombineFrames,
	)
	if err != nil {
		logger.ERROR(IURBPP002, "build document plan failed", map[string]any{"error": err})
		return entity.DocumentPlan{}, fmt.Errorf("build document draw plan: %w", err)
	}
	var document entity.DocumentPlan
	if err := json.Unmarshal(planJSON, &document); err != nil {
		return entity.DocumentPlan{}, fmt.Errorf("decode document draw plan: %w", err)
	}
	if uniformPages {
		v1engine.NormalizeDocumentPageSizesV1EnginePlanDocument(&document)
	}
	return document, nil
}

func resolvedV2CatalogIcons(document entity.EngineResolvedDocument, opts entity.RenderOptions) (map[string]string, error) {
	wanted := make(map[int]string)
	for _, element := range document.Elements {
		const prefix = "catalog:"
		if !strings.HasPrefix(element.IconRef, prefix) {
			continue
		}
		id, err := strconv.Atoi(strings.TrimPrefix(element.IconRef, prefix))
		if err == nil {
			wanted[id] = element.IconRef
		}
	}
	if len(wanted) == 0 {
		return nil, nil
	}
	catalog, err := readV2Catalog(opts)
	if err != nil {
		return nil, err
	}
	icons := make(map[string]string, len(wanted))
	for id, ref := range wanted {
		if data := catalog[id]; data != "" {
			icons[ref] = data
		}
	}
	return icons, nil
}

func renderDocumentVersion(input []byte) (string, error) {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	for {
		token, err := decoder.Token()
		if err != nil {
			return "", fmt.Errorf("inspect XAL document root: %w", err)
		}
		start, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		version := ""
		for _, attr := range start.Attr {
			if attr.Name.Local == "version" {
				version = strings.TrimSpace(attr.Value)
				break
			}
		}
		switch start.Name.Local {
		case "scene":
			if version != "2" {
				return "", fmt.Errorf("<scene> requires version=\"2\"")
			}
			return "2", nil
		case "xaligo":
			if version == "" || version == "1" {
				return "1", nil
			}
			return "", fmt.Errorf("<xaligo> accepts only version=\"1\"; V2 uses <scene version=\"2\">")
		case "frame", "frames":
			return "1", nil
		default:
			return "", fmt.Errorf("unsupported XAL document root <%s>", start.Name.Local)
		}
	}
}

// BuildPPTXPlan remains as the format-specific compatibility boundary for WASM.
func (rcvr *renderUsecase) BuildPPTXPlan(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	document, err := rcvr.buildDocumentPlan(ctx, input, opts, true)
	if err != nil {
		return nil, err
	}
	// Preserve the established single-page JSON boundary for callers that do
	// not need a document wrapper. Multi-frame output uses schemaVersion 2.
	if len(document.Pages) == 1 {
		page := document.Pages[0]
		return json.Marshal(entity.Plan{
			Slide: page.Slide, Ops: page.Ops,
			Legend: document.Legend, ConnectorLegend: document.ConnectorLegend,
		})
	}
	return json.Marshal(document)
}

func (rcvr *renderUsecase) NewPreviewRepository(path string, opts entity.PreviewOptions) (repository.PreviewRepository, error) {
	renderPreview := rcvr.renderSVGPreview
	if opts.Kind == entity.PreviewKindHTML {
		renderPreview = rcvr.renderMarkdownPreview
	}
	return repository.NewPreviewRepository(
		path,
		opts,
		renderPreview,
		rcvr.ValidateRenderOptions,
		rcvr.diagnose,
		rcvr.xaligoRepository.ReadSource,
	)
}

// renderSVGPreview renders a .xal source onto one combined SVG canvas for
// the live preview server.
func (rcvr *renderUsecase) renderSVGPreview(ctx context.Context, input []byte, renderOpts entity.RenderOptions) ([]byte, error) {
	renderOpts.CombineFrames = true
	return rcvr.RenderSVG(ctx, input, renderOpts)
}

// renderMarkdownPreview renders every ```xal code block in a Markdown source
// to an isolated SVG image and converts the result to a standalone HTML
// document for the live preview server.
func (rcvr *renderUsecase) renderMarkdownPreview(ctx context.Context, input []byte, renderOpts entity.RenderOptions) ([]byte, error) {
	blockIndex := 0
	placeholderPrefix := "xaligo-diagram-placeholder"
	for strings.Contains(string(input), placeholderPrefix) {
		placeholderPrefix += "-safe"
	}
	diagrams := make([]markdownPreviewDiagram, 0)
	embedded, err := EmbedXalCodeBlocks(string(input), func(xal string) ([]string, error) {
		blockIndex++
		artifacts, renderErr := rcvr.RenderArtifacts(ctx, []byte(xal), renderOpts)
		if renderErr != nil {
			return nil, fmt.Errorf("render xal code block %d: %w", blockIndex, renderErr)
		}
		lines := make([]string, 0, len(artifacts)*2)
		for artifactIndex, artifact := range artifacts {
			placeholder := fmt.Sprintf("%s-%d-%d", placeholderPrefix, blockIndex, artifactIndex+1)
			alt := "xaligo diagram"
			if strings.TrimSpace(artifact.ID) != "" {
				alt += " " + artifact.ID
			}
			diagrams = append(diagrams, markdownPreviewDiagram{
				placeholder: placeholder,
				alt:         alt,
				svg:         append([]byte(nil), artifact.Data...),
			})
			lines = append(lines, "", placeholder, "")
		}
		return lines, nil
	})
	if err != nil {
		return nil, err
	}
	return renderMarkdownHTMLDocument(embedded, diagrams)
}

// EmbedXalCodeBlocks scans Markdown source for fenced code blocks whose info
// string is exactly "xal" (``` or ~~~ fences, up to 3 leading spaces of
// indentation per CommonMark) and replaces each one with the lines returned
// by renderBlock for that block's body. Every other line is preserved as-is.
// Shared by the `render markdown` file-output flow and the Markdown live
// preview flow.
func EmbedXalCodeBlocks(source string, renderBlock func(xal string) ([]string, error)) (string, error) {
	lines := strings.Split(source, "\n")
	output := make([]string, 0, len(lines))
	lineIndex := 0
	for lineIndex < len(lines) {
		line := lines[lineIndex]
		fenceChar, fenceLen, info, isFence := parseMarkdownFenceOpen(line)
		if !isFence {
			output = append(output, line)
			lineIndex++
			continue
		}
		bodyStart := lineIndex + 1
		closeIndex := findMarkdownFenceClose(lines, bodyStart, fenceChar, fenceLen)
		if info != "xal" {
			if closeIndex == -1 {
				output = append(output, lines[lineIndex:]...)
				break
			}
			output = append(output, lines[lineIndex:closeIndex+1]...)
			lineIndex = closeIndex + 1
			continue
		}
		if closeIndex == -1 {
			return "", fmt.Errorf("unterminated ```xal code fence starting at line %d", lineIndex+1)
		}
		body := strings.Join(lines[bodyStart:closeIndex], "\n")
		replacement, err := renderBlock(body)
		if err != nil {
			return "", err
		}
		output = append(output, replacement...)
		lineIndex = closeIndex + 1
	}
	return strings.Join(output, "\n"), nil
}

// parseMarkdownFenceOpen reports whether line opens a fenced code block (at
// most 3 leading spaces, 3+ backticks or tildes), returning the fence
// character, fence length, and trimmed info string.
func parseMarkdownFenceOpen(line string) (fenceChar byte, fenceLen int, info string, ok bool) {
	trimmed := strings.TrimLeft(line, " ")
	if len(line)-len(trimmed) > 3 {
		return 0, 0, "", false
	}
	if len(trimmed) < 3 {
		return 0, 0, "", false
	}
	fenceChar = trimmed[0]
	if fenceChar != '`' && fenceChar != '~' {
		return 0, 0, "", false
	}
	for fenceLen < len(trimmed) && trimmed[fenceLen] == fenceChar {
		fenceLen++
	}
	if fenceLen < 3 {
		return 0, 0, "", false
	}
	return fenceChar, fenceLen, strings.TrimSpace(trimmed[fenceLen:]), true
}

// findMarkdownFenceClose returns the index of the first line at or after
// start that closes a fence opened with fenceChar/fenceLen, or -1 if none is
// found.
func findMarkdownFenceClose(lines []string, start int, fenceChar byte, fenceLen int) int {
	for index := start; index < len(lines); index++ {
		trimmed := strings.TrimLeft(lines[index], " ")
		if len(lines[index])-len(trimmed) > 3 {
			continue
		}
		count := 0
		for count < len(trimmed) && trimmed[count] == fenceChar {
			count++
		}
		if count >= fenceLen && strings.TrimSpace(trimmed[count:]) == "" {
			return index
		}
	}
	return -1
}

type markdownPreviewDiagram struct {
	placeholder string
	alt         string
	svg         []byte
}

// renderMarkdownHTMLDocument converts Markdown source into a standalone HTML
// document. Raw Markdown HTML remains disabled. Rendered diagrams are inserted
// afterwards as data-URL images so every SVG has its own document scope and
// duplicate SVG IDs cannot collide.
func renderMarkdownHTMLDocument(source string, diagrams []markdownPreviewDiagram) ([]byte, error) {
	converter := goldmark.New()
	sourceBytes := []byte(source)
	document := converter.Parser().Parse(goldmarktext.NewReader(sourceBytes))
	if err := goldmarkast.Walk(document, func(node goldmarkast.Node, entering bool) (goldmarkast.WalkStatus, error) {
		if !entering {
			return goldmarkast.WalkContinue, nil
		}
		image, ok := node.(*goldmarkast.Image)
		if !ok {
			return goldmarkast.WalkContinue, nil
		}
		image.Destination = rewriteMarkdownImageDestination(image.Destination)
		return goldmarkast.WalkContinue, nil
	}); err != nil {
		return nil, fmt.Errorf("inspect markdown images: %w", err)
	}
	var htmlBody bytes.Buffer
	if err := converter.Renderer().Render(&htmlBody, sourceBytes, document); err != nil {
		return nil, fmt.Errorf("convert markdown to HTML: %w", err)
	}
	body := htmlBody.String()
	for _, diagram := range diagrams {
		needle := "<p>" + diagram.placeholder + "</p>"
		if !strings.Contains(body, needle) {
			return nil, fmt.Errorf("embed rendered diagram: placeholder %q was not preserved", diagram.placeholder)
		}
		dataURL := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(diagram.svg)
		replacement := `<figure class="xaligo-diagram"><img src="` + dataURL +
			`" alt="` + html.EscapeString(diagram.alt) + `" loading="lazy"></figure>`
		body = strings.Replace(body, needle, replacement, 1)
	}
	var doc bytes.Buffer
	doc.WriteString(markdownPreviewHTMLHeader)
	doc.WriteString(body)
	doc.WriteString(markdownPreviewHTMLFooter)
	return doc.Bytes(), nil
}

func rewriteMarkdownImageDestination(destination []byte) []byte {
	raw := strings.TrimSpace(string(destination))
	parsed, err := url.Parse(raw)
	if err != nil {
		return []byte("#xaligo-blocked-relative-image")
	}
	if parsed.IsAbs() || parsed.Host != "" || strings.HasPrefix(raw, "//") ||
		parsed.Path == "" || strings.HasPrefix(parsed.Path, "/") {
		return destination
	}
	slashPath := strings.ReplaceAll(parsed.Path, `\`, "/")
	cleanPath := path.Clean(slashPath)
	if cleanPath == "." || cleanPath == ".." || strings.HasPrefix(cleanPath, "../") {
		return []byte("#xaligo-blocked-relative-image")
	}
	parsed.Path = "/assets/" + cleanPath
	parsed.RawPath = ""
	return []byte(parsed.String())
}

const markdownPreviewHTMLHeader = `<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>xaligo markdown preview</title><style>
body{margin:0;padding:24px;background:#fff;color:#111827;font-family:system-ui,sans-serif;line-height:1.6}
.xaligo-diagram{margin:8px 0}.xaligo-diagram img{max-width:100%;height:auto;border:0;box-shadow:none;display:block}
pre{background:#f3f4f6;padding:12px;border-radius:6px;overflow:auto}
code{background:#f3f4f6;padding:2px 4px;border-radius:4px}
</style></head><body>
`

const markdownPreviewHTMLFooter = `
</body></html>`

func (rcvr *renderUsecase) diagnose(ctx context.Context, input []byte) ([]entity.Diagnostic, error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	diagnostics := v1engine.DiagnoseV1EngineDiagnoseDocument(input)
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	return diagnostics, nil
}

func (rcvr *renderUsecase) RenderPPTX(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	planJSON, err := rcvr.BuildPPTXPlan(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRP001, "build plan failed", map[string]any{"error": err})
		return nil, err
	}
	data, err := rcvr.powerpointRepository.ExportPptxBytes(ctx, entity.PptxExportOptions{
		PlanJSON: planJSON, Title: opts.Title, Author: opts.Author,
		Company: opts.Company, Subject: opts.Subject, Compression: opts.Compression,
	})
	if err != nil {
		logger.ERROR(IURRP003, "export failed", map[string]any{"error": err})
		return nil, err
	}
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURRP004, "context check failed", map[string]any{"error": err})
		return nil, err
	}
	return data, nil
}

func (rcvr *renderUsecase) serviceOptions(opts entity.RenderOptions) ([]entity.ServiceEntry, map[int]string, error) {
	if len(bytes.TrimSpace(opts.ServicesCSV)) == 0 {
		logger.DEBUG(IURSO001, "branch no services csv")
		abbreviations, err := v1engine.ResolveServiceOptionsV1EngineOptionService(nil, nil)
		return nil, abbreviations, err
	}

	logger.DEBUG(IURSO002, "branch services csv", map[string]any{"bytes": len(opts.ServicesCSV)})
	if err := v1engine.ValidateLegendCSVRowsV1EngineOptionService(opts.ServicesCSV); err != nil {
		logger.ERROR(IURSO005, "legend validation failed", map[string]any{"error": err})
		return nil, nil, err
	}
	entries, err := rcvr.xaligoRepository.ReadServiceListFromReader(bytes.NewReader(opts.ServicesCSV))
	if err != nil {
		logger.ERROR(IURSO003, "read services csv failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("read services CSV: %w", err)
	}
	abbreviations, err := v1engine.ResolveServiceOptionsV1EngineOptionService(entries, nil)
	if err != nil {
		logger.ERROR(IURSO005, "legend validation failed", map[string]any{"error": err})
		return nil, nil, err
	}
	return entries, abbreviations, nil
}

// buildScene is the orchestration boundary around synchronous V1 engine
// stages. Cancellation checks, repository-backed option loading, and stage
// ordering intentionally remain here so future concurrency can be controlled
// by the usecase layer without changing calculation code.
func (rcvr *renderUsecase) buildScene(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, []entity.ServiceEntry, error) {
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}
	if err := v1engine.ValidateRenderOptionsV1EngineOptionRender(opts); err != nil {
		logger.ERROR(IURBS002, "validate render options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	theme, err := v1engine.NormalizeThemeV1EngineThemeApply(opts.Theme)
	if err != nil {
		return nil, nil, err
	}

	doc, err := v1engine.ParseWithImportsV1EngineParseDocument(bytes.NewReader(input), opts.Imports)
	if err != nil {
		logger.ERROR(IURBS007, "parse DSL failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("parse DSL: %w", err)
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	root, err := v1engine.BuildV1EngineLayoutBuild(doc)
	if err != nil {
		logger.ERROR(IURBS008, "build layout failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build layout: %w", err)
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	entries, abbreviations, err := rcvr.serviceOptions(opts)
	if err != nil {
		logger.ERROR(IURBS009, "service options failed", map[string]any{"error": err})
		return nil, nil, err
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	connections := v1engine.CollectConnectionNodesV1EngineSceneConnection(doc.Root)
	dependencies := SceneDependencies{
		XaligoRepository: rcvr.xaligoRepository,
		SceneRepository:  rcvr.sceneRepository,
	}.core()

	var scene []byte
	if opts.Assets != nil {
		logger.DEBUG(IURBS004, "branch embedded assets")
		itemSize := opts.Assets.ItemIconSize
		if itemSize <= 0 {
			logger.DEBUG(IURBS005, "branch default embedded item size")
			itemSize = 32
		}
		scene, err = v1engine.BuildJSONWithFSV1EngineSceneBuild(root, opts.Assets.FS, opts.Assets.CatalogCSV, opts.Assets.GroupIconsDir, itemSize, connections, abbreviations, dependencies)
	} else {
		logger.DEBUG(IURBS006, "branch native assets")
		cfg := config.New()
		scene, err = v1engine.BuildJSONV1EngineSceneBuild(root, filepath.Join(cfg.AssetDir_, "Architecture-Group-Icons"), cfg.SvcCatalogCSV, cfg.ProjectRoot, cfg.ItemIconSize, connections, abbreviations, nil, dependencies)
	}
	if err != nil {
		logger.ERROR(IURBS010, "build JSON failed", map[string]any{"error": err})
		return nil, nil, fmt.Errorf("build presentation scene: %w", err)
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}

	scene, err = v1engine.ApplyThemeJSONV1EngineThemeApply(scene, theme)
	if err != nil {
		logger.ERROR(IURBS011, "apply theme failed", map[string]any{"error": err})
		return nil, nil, err
	}
	if err := rcvr.checkRenderContext(ctx); err != nil {
		return nil, nil, err
	}
	return scene, entries, nil
}

func (rcvr *renderUsecase) checkRenderContext(ctx context.Context) error {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURBS001, "context check failed", map[string]any{"error": err})
		return err
	}
	return nil
}

// ValidateRenderOptions is kept as a source-compatible package boundary.
// Deprecated: construct RenderUsecase and call ValidateRenderOptions.
func ValidateRenderOptions(opts entity.RenderOptions) error {
	return validateRenderOptions(opts)
}
