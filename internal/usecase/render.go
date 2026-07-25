package usecase

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"math"
	"net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/config"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/share"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
	"github.com/yuin/goldmark"
	goldmarkast "github.com/yuin/goldmark/ast"
	goldmarktext "github.com/yuin/goldmark/text"
)

// RenderUsecase owns format dispatch and the shared render pipeline.
type RenderUsecase interface {
	ValidateRenderOptions(entity.RenderOptions) error
	Render(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderArtifacts(context.Context, []byte, entity.RenderOptions) ([]entity.RenderArtifact, error)
	RenderExcalidraw(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderSVG(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderPPTX(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderPDF(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderExcel(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderXYFlow(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	RenderIsoflow(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	BuildPPTXPlan(context.Context, []byte, entity.RenderOptions) ([]byte, error)
	NewPreviewRepository(string, entity.PreviewOptions) (repository.PreviewRepository, error)
}

type renderUsecase struct {
	excalidrawRepository  repository.ExcalidrawRepository
	xaligoRepository      repository.XaligoRepository
	powerpointRepository  repository.PowerpointRepository
	isoflowRepository     repository.IsoflowRepository
	svgRepository         repository.SVGRepository
	xyFlowRepository      repository.XYFlowRepository
	pdfRepository         repository.PDFRepository
	spreadsheetRepository repository.SpreadsheetRepository
}

func NewRenderUsecase(
	excalidrawRepository repository.ExcalidrawRepository,
	xaligoRepository repository.XaligoRepository,
	powerpointRepository repository.PowerpointRepository,
	isoflowRepository repository.IsoflowRepository,
	svgRepository repository.SVGRepository,
	xyFlowRepository repository.XYFlowRepository,
	pdfRepository repository.PDFRepository,
	spreadsheetRepository repository.SpreadsheetRepository,
) RenderUsecase {
	return &renderUsecase{
		excalidrawRepository:  excalidrawRepository,
		xaligoRepository:      xaligoRepository,
		powerpointRepository:  powerpointRepository,
		isoflowRepository:     isoflowRepository,
		svgRepository:         svgRepository,
		xyFlowRepository:      xyFlowRepository,
		pdfRepository:         pdfRepository,
		spreadsheetRepository: spreadsheetRepository,
	}
}

const (
	ModeStandard = v1engine.ModeStandardV1EngineOptionRender
	ModeNetwork  = v1engine.ModeNetworkV1EngineOptionRender
	ModeAWS      = v1engine.ModeAWSV1EngineOptionRender

	FormatExcalidraw = v1engine.FormatExcalidrawV1EngineOptionRender
	FormatSVG        = v1engine.FormatSVGV1EngineOptionRender
	FormatPPTX       = v1engine.FormatPPTXV1EngineOptionRender
	FormatPDF        = v1engine.FormatPDFV1EngineOptionRender
	FormatExcel      = v1engine.FormatExcelV1EngineOptionRender
	FormatXYFlow     = v1engine.FormatXYFlowV1EngineOptionRender
	FormatIsoflow    = v1engine.FormatIsoflowV1EngineOptionRender

	SeverityError   = v1engine.SeverityErrorV1EngineOptionRender
	SeverityWarning = v1engine.SeverityWarningV1EngineOptionRender
)

var ErrNotImplemented = v1engine.ErrNotImplementedV1EngineOptionRender

var (
	logger   = share.DefaultLogger()
	IURR001  = share.NewMCode("IURR-001", "Render context check failed")
	IURR002  = share.NewMCode("IURR-002", "Render validate render options failed")
	IURR003  = share.NewMCode("IURR-003", "Render default format branch")
	IURR004  = share.NewMCode("IURR-004", "Render excalidraw branch")
	IURR005  = share.NewMCode("IURR-005", "Render SVG branch")
	IURR006  = share.NewMCode("IURR-006", "Render PPTX branch")
	IURR007  = share.NewMCode("IURR-007", "Render XYFlow branch")
	IURR008  = share.NewMCode("IURR-008", "Render Isoflow branch")
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
	return v1engine.ValidateRenderOptionsV1EngineOptionRender(opts)
}

func (rcvr *renderUsecase) Render(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	if err := checkContext(ctx); err != nil {
		logger.ERROR(IURR001, "context check failed", map[string]any{"error": err})
		return nil, err
	}
	if err := ValidateRenderOptions(opts); err != nil {
		logger.ERROR(IURR002, "validate render options failed", map[string]any{"format": opts.Format, "error": err})
		return nil, err
	}
	format := entity.Format(strings.ToLower(strings.TrimSpace(string(opts.Format))))
	if format == "" {
		logger.DEBUG(IURR003, "branch default format")
		format = FormatExcalidraw
	}
	opts.Format = format
	if format == FormatExcalidraw && containsUMLRenderUsecase(input) {
		return nil, fmt.Errorf("UML Excalidraw export is disabled; use svg, pdf, pptx, excel, xyflow, or isoflow instead")
	}
	switch format {
	case FormatExcalidraw:
		logger.DEBUG(IURR004, "branch excalidraw")
		return rcvr.RenderExcalidraw(ctx, input, opts)
	case FormatSVG:
		logger.DEBUG(IURR005, "branch svg")
		return rcvr.RenderSVG(ctx, input, opts)
	case FormatPPTX:
		logger.DEBUG(IURR006, "branch pptx")
		return rcvr.RenderPPTX(ctx, input, opts)
	case FormatPDF:
		return rcvr.RenderPDF(ctx, input, opts)
	case FormatExcel:
		return rcvr.RenderExcel(ctx, input, opts)
	case FormatXYFlow:
		logger.DEBUG(IURR007, "branch xyflow")
		return rcvr.RenderXYFlow(ctx, input, opts)
	case FormatIsoflow:
		logger.DEBUG(IURR008, "branch isoflow")
		return rcvr.RenderIsoflow(ctx, input, opts)
	default:
		logger.ERROR(IURR009, "branch unknown format", map[string]any{"format": format})
		return nil, fmt.Errorf("unknown render format %q", format)
	}
}

func containsUMLRenderUsecase(input []byte) bool {
	decoder := xml.NewDecoder(bytes.NewReader(input))
	for {
		token, err := decoder.Token()
		if err != nil {
			return false
		}
		start, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		if start.Name.Local == "uml" {
			return true
		}
	}
}

var IURRE001 = share.NewMCode("IURRE-001", "Render Excalidraw completed")

func (rcvr *renderUsecase) RenderExcalidraw(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := rcvr.buildScene(ctx, input, opts)
	if err == nil {
		logger.DEBUG(IURRE001, "completed")
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
	IURBPP001 = share.NewMCode("IURBPP-001", "Build PPTX plan build scene failed")
	IURBPP002 = share.NewMCode("IURBPP-002", "Build PPTX plan build plan failed")
	IURRP001  = share.NewMCode("IURRP-001", "Render PPTX build plan failed")
	IURRP003  = share.NewMCode("IURRP-003", "Render PPTX export failed")
	IURRP004  = share.NewMCode("IURRP-004", "Render PPTX context check failed")
)

func (rcvr *renderUsecase) buildDocumentPlan(ctx context.Context, input []byte, opts entity.RenderOptions, uniformPages bool) (entity.DocumentPlan, error) {
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
.xaligo-diagram{margin:8px 0}.xaligo-diagram img{max-width:100%;height:auto;box-shadow:0 4px 16px #0002;display:block}
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
		ExporterWASM: opts.PPTXExporterWASM,
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

// RenderPDF maps each document page to one PDF page.
func (rcvr *renderUsecase) RenderPDF(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	if rcvr.pdfRepository == nil {
		return nil, fmt.Errorf("PDF repository is not available in this runtime")
	}
	document, err := rcvr.buildDocumentPlan(ctx, input, opts, false)
	if err != nil {
		return nil, err
	}
	pages, err := rcvr.renderDocumentPages(document, opts)
	if err != nil {
		return nil, err
	}
	return rcvr.pdfRepository.Export(ctx, pages)
}

// RenderExcel maps each document page to one worksheet with the page SVG
// embedded as a vector image.
func (rcvr *renderUsecase) RenderExcel(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	if rcvr.spreadsheetRepository == nil {
		return nil, fmt.Errorf("Excel repository is not available in this runtime")
	}
	document, err := rcvr.buildDocumentPlan(ctx, input, opts, false)
	if err != nil {
		return nil, err
	}
	pages, err := rcvr.renderDocumentPages(document, opts)
	if err != nil {
		return nil, err
	}
	return rcvr.spreadsheetRepository.Export(ctx, pages)
}

func (rcvr *renderUsecase) renderDocumentPages(document entity.DocumentPlan, opts entity.RenderOptions) ([]entity.RenderPage, error) {
	pxPerInch := opts.PxPerInch
	if pxPerInch <= 0 {
		pxPerInch = 96
	}
	pages := make([]entity.RenderPage, 0, len(document.Pages))
	for _, page := range document.Pages {
		svg, err := rcvr.svgRepository.Render(entity.Plan{
			Slide: page.Slide, Ops: page.Ops, Legend: document.Legend,
		}, pxPerInch, opts.SVGLegendPosition)
		if err != nil {
			return nil, fmt.Errorf("render document page %q as SVG: %w", page.ID, err)
		}
		widthPx, heightPx, err := intrinsicSVGDimensionsRenderDocument(svg)
		if err != nil {
			return nil, fmt.Errorf("resolve document page %q SVG dimensions: %w", page.ID, err)
		}
		pages = append(pages, entity.RenderPage{
			ID: page.ID, SVG: svg,
			WidthPx: widthPx, HeightPx: heightPx,
		})
	}
	return pages, nil
}

func intrinsicSVGDimensionsRenderDocument(data []byte) (float64, float64, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return 0, 0, fmt.Errorf("SVG root element is missing")
		}
		if err != nil {
			return 0, 0, fmt.Errorf("parse SVG root: %w", err)
		}
		start, ok := token.(xml.StartElement)
		if !ok {
			continue
		}
		if start.Name.Local != "svg" {
			return 0, 0, fmt.Errorf("root element is <%s>, expected <svg>", start.Name.Local)
		}
		width, height := "", ""
		for _, attribute := range start.Attr {
			switch attribute.Name.Local {
			case "width":
				width = attribute.Value
			case "height":
				height = attribute.Value
			}
		}
		parsedWidth, err := parseSVGPixelDimensionRenderDocument(width)
		if err != nil {
			return 0, 0, fmt.Errorf("width: %w", err)
		}
		parsedHeight, err := parseSVGPixelDimensionRenderDocument(height)
		if err != nil {
			return 0, 0, fmt.Errorf("height: %w", err)
		}
		return parsedWidth, parsedHeight, nil
	}
}

func parseSVGPixelDimensionRenderDocument(value string) (float64, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, fmt.Errorf("dimension is missing")
	}
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("dimension %q must be a unitless pixel number: %w", value, err)
	}
	if parsed <= 0 || math.IsNaN(parsed) || math.IsInf(parsed, 0) {
		return 0, fmt.Errorf("dimension %q must be positive and finite", value)
	}
	return parsed, nil
}

var (
	IURRXYF001 = share.NewMCode("IURRXYF-001", "Render XYFlow build scene failed")
	IURRXYF002 = share.NewMCode("IURRXYF-002", "Render XYFlow scene failed")
)

func (rcvr *renderUsecase) RenderXYFlow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := rcvr.buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRXYF001, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	out, err := rcvr.xyFlowRepository.Render(scene)
	if err != nil {
		logger.ERROR(IURRXYF002, "render scene failed", map[string]any{"error": err})
		return nil, err
	}
	return out, nil
}

var (
	IURRI001 = share.NewMCode("IURRI-001", "Render Isoflow load icons")
	IURRI003 = share.NewMCode("IURRI-003", "Render Isoflow build scene failed")
	IURRI004 = share.NewMCode("IURRI-004", "Render Isoflow with icons failed")
)

func (rcvr *renderUsecase) RenderIsoflow(ctx context.Context, input []byte, opts entity.RenderOptions) ([]byte, error) {
	scene, _, err := rcvr.buildScene(ctx, input, opts)
	if err != nil {
		logger.ERROR(IURRI003, "build scene failed", map[string]any{"error": err})
		return nil, err
	}
	logger.DEBUG(IURRI001, "load isoflow icons")
	icons, _ := rcvr.isoflowRepository.LoadIsoflowIcons(opts.Assets)
	out, err := rcvr.isoflowRepository.RenderWithIcons(scene, icons)
	if err != nil {
		logger.ERROR(IURRI004, "render with icons failed", map[string]any{"error": err})
		return nil, err
	}
	return out, nil
}

func (rcvr *renderUsecase) serviceOptions(opts entity.RenderOptions) ([]entity.ServiceEntry, map[int]string, error) {
	if len(bytes.TrimSpace(opts.ServicesCSV)) == 0 {
		logger.DEBUG(IURSO001, "branch no services csv", map[string]any{"abbreviations": len(opts.Abbreviations)})
		abbreviations, err := v1engine.ResolveServiceOptionsV1EngineOptionService(nil, opts.Abbreviations)
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
	abbreviations, err := v1engine.ResolveServiceOptionsV1EngineOptionService(entries, opts.Abbreviations)
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
		XaligoRepository:     rcvr.xaligoRepository,
		ExcalidrawRepository: rcvr.excalidrawRepository,
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
		return nil, nil, fmt.Errorf("build excalidraw JSON: %w", err)
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
	return v1engine.ValidateRenderOptionsV1EngineOptionRender(opts)
}
