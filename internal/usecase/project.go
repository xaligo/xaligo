package usecase

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	projectrepository "github.com/xaligo/xaligo/internal/repository/project"
	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

// ProjectUsecase owns source analysis and the durable concept index shared by
// RAG, LSP workspace features, and MCP project tools.
type ProjectUsecase interface {
	Analyze(context.Context, string, []byte) (entity.ProjectAnalysis, error)
	IndexDocument(context.Context, string, []byte) (entity.ProjectAnalysis, bool, error)
	Index(context.Context, string) (entity.ProjectIndexStats, error)
	Search(context.Context, string, int) ([]entity.ProjectSearchResult, error)
	Symbols(context.Context, string) ([]entity.ProjectSymbol, error)
	Close() error
}

type projectUsecase struct {
	index projectrepository.IndexRepository
}

func NewProjectUsecase(index projectrepository.IndexRepository) ProjectUsecase {
	return &projectUsecase{index: index}
}

const (
	projectMaxSourceBytes = 8 * 1024 * 1024
	projectMaxQueryBytes  = 2 * 1024
	projectMaxDetailBytes = 16 * 1024
)

func (rcvr *projectUsecase) Analyze(ctx context.Context, uri string, source []byte) (entity.ProjectAnalysis, error) {
	if err := checkContext(ctx); err != nil {
		return entity.ProjectAnalysis{}, err
	}
	uri = strings.TrimSpace(uri)
	if uri == "" {
		return entity.ProjectAnalysis{}, errors.New("project document URI must not be empty")
	}
	if len(source) > projectMaxSourceBytes {
		return entity.ProjectAnalysis{}, fmt.Errorf("project source size %d exceeds %d", len(source), projectMaxSourceBytes)
	}
	analysis := entity.ProjectAnalysis{
		URI:      uri,
		Checksum: sha256.Sum256(source),
		Source:   append([]byte(nil), source...),
	}
	extension := strings.ToLower(filepath.Ext(projectURIPath(uri)))
	switch extension {
	case ".md", ".markdown":
		analysis.Kind = entity.ProjectDocumentMarkdown
		analysis.Symbols = projectMarkdownSymbols(uri, source)
	default:
		analysis.Kind = entity.ProjectDocumentXAL
		imports := projectImportSource(uri)
		document, diagnostics := v1engine.AnalyzeWithImportsV1EngineDiagnoseDocument(source, imports)
		analysis.Diagnostics = diagnostics
		if document.Envelope != nil {
			analysis.Symbols = projectXALSymbols(document.Envelope)
		}
	}
	if err := checkContext(ctx); err != nil {
		return entity.ProjectAnalysis{}, err
	}
	return analysis, nil
}

func (rcvr *projectUsecase) IndexDocument(ctx context.Context, uri string, source []byte) (entity.ProjectAnalysis, bool, error) {
	if rcvr.index == nil {
		return entity.ProjectAnalysis{}, false, errors.New("project index repository is required")
	}
	analysis, err := rcvr.Analyze(ctx, uri, source)
	if err != nil {
		return entity.ProjectAnalysis{}, false, err
	}
	rootURI := projectDocumentRootURI(uri)
	changed, err := rcvr.index.Put(ctx, rootURI, analysis)
	if err != nil {
		return entity.ProjectAnalysis{}, false, err
	}
	return analysis, changed, nil
}

func (rcvr *projectUsecase) Index(ctx context.Context, root string) (entity.ProjectIndexStats, error) {
	if rcvr.index == nil {
		return entity.ProjectIndexStats{}, errors.New("project index repository is required")
	}
	root = strings.TrimSpace(root)
	if root == "" {
		root = "."
	}
	absolute, err := filepath.Abs(root)
	if err != nil {
		return entity.ProjectIndexStats{}, fmt.Errorf("resolve project root: %w", err)
	}
	info, err := os.Stat(absolute)
	if err != nil {
		return entity.ProjectIndexStats{}, fmt.Errorf("inspect project root: %w", err)
	}
	if !info.IsDir() {
		return entity.ProjectIndexStats{}, fmt.Errorf("project root %s is not a directory", absolute)
	}
	docsRoot, err := projectDocsRoot(absolute)
	if err != nil {
		return entity.ProjectIndexStats{}, err
	}
	absolute = docsRoot
	rootURI := projectPathURI(absolute)
	stats := entity.ProjectIndexStats{Root: absolute}
	var seen []string
	err = filepath.WalkDir(absolute, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if err := checkContext(ctx); err != nil {
			return err
		}
		if entry.IsDir() {
			if path != absolute && projectIgnoredDirectory(entry.Name()) {
				return filepath.SkipDir
			}
			return nil
		}
		if !projectRAGExtension(path) {
			return nil
		}
		fileInfo, err := entry.Info()
		if err != nil {
			return fmt.Errorf("inspect project file %s: %w", path, err)
		}
		if fileInfo.Size() > projectMaxSourceBytes {
			return fmt.Errorf("project source %s size %d exceeds %d", path, fileInfo.Size(), projectMaxSourceBytes)
		}
		source, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read project source %s: %w", path, err)
		}
		uri := projectPathURI(path)
		analysis, err := rcvr.Analyze(ctx, uri, source)
		if err != nil {
			return fmt.Errorf("analyze project source %s: %w", path, err)
		}
		changed, err := rcvr.index.Put(ctx, rootURI, analysis)
		if err != nil {
			return fmt.Errorf("index project source %s: %w", path, err)
		}
		stats.Scanned++
		stats.Diagnostics += len(analysis.Diagnostics)
		if changed {
			stats.Indexed++
		} else {
			stats.Unchanged++
		}
		seen = append(seen, uri)
		return nil
	})
	if err != nil {
		return entity.ProjectIndexStats{}, fmt.Errorf("walk project root %s: %w", absolute, err)
	}
	stats.Removed, err = rcvr.index.Prune(ctx, rootURI, seen)
	if err != nil {
		return entity.ProjectIndexStats{}, err
	}
	return stats, nil
}

func (rcvr *projectUsecase) Search(ctx context.Context, query string, limit int) ([]entity.ProjectSearchResult, error) {
	if rcvr.index == nil {
		return nil, errors.New("project index repository is required")
	}
	query = strings.TrimSpace(query)
	if query == "" {
		return nil, errors.New("project search query must not be empty")
	}
	if len(query) > projectMaxQueryBytes {
		return nil, fmt.Errorf("project search query exceeds %d UTF-8 bytes", projectMaxQueryBytes)
	}
	return rcvr.index.Search(ctx, query, limit)
}

func (rcvr *projectUsecase) Symbols(ctx context.Context, uri string) ([]entity.ProjectSymbol, error) {
	if rcvr.index == nil {
		return nil, errors.New("project index repository is required")
	}
	uri = strings.TrimSpace(uri)
	if uri == "" {
		return nil, errors.New("project document URI must not be empty")
	}
	return rcvr.index.Symbols(ctx, uri)
}

func (rcvr *projectUsecase) Close() error {
	if rcvr.index == nil {
		return nil
	}
	return rcvr.index.Close()
}

func projectXALSymbols(root *entity.Node) []entity.ProjectSymbol {
	type walker struct {
		symbols []entity.ProjectSymbol
	}
	state := &walker{}
	var visit func(*entity.Node, int, string, int)
	visit = func(node *entity.Node, parentOrdinal int, parentPath string, sibling int) {
		if node == nil {
			return
		}
		identity := projectNodeIdentity(node)
		segment := fmt.Sprintf("%s[%d]", node.Tag, sibling)
		if identity != "" {
			segment = node.Tag + "#" + identity
		}
		path := segment
		if parentPath != "" {
			path = parentPath + "/" + segment
		}
		childParent := parentOrdinal
		if concept, include := projectNodeConcept(node); include {
			ordinal := len(state.symbols)
			semanticID := identity
			if semanticID == "" {
				semanticID = path
			}
			state.symbols = append(state.symbols, entity.ProjectSymbol{
				Ordinal:       ordinal,
				ParentOrdinal: parentOrdinal,
				ID:            semanticID,
				Name:          projectNodeName(node, semanticID),
				Detail:        projectNodeDetail(node),
				Concept:       concept,
				SourceTag:     node.Tag,
				Source:        strings.TrimSpace(node.Attr("src")),
				Target:        strings.TrimSpace(node.Attr("dst")),
				Position:      node.Position,
			})
			childParent = ordinal
		}
		for index, child := range node.Children {
			visit(child, childParent, path, index)
		}
	}
	visit(root, -1, "", 0)
	return state.symbols
}

func projectNodeConcept(node *entity.Node) (entity.ProjectConcept, bool) {
	switch node.Tag {
	case "xaligo", "scene", "frames", "data", "connections", "bends":
		return "", false
	case "frame":
		return entity.ProjectConceptFrame, true
	case "capture":
		return entity.ProjectConceptCapture, true
	case "port":
		return entity.ProjectConceptPort, true
	case "connection", "line":
		return entity.ProjectConceptLine, true
	case "text", "label", "title":
		return entity.ProjectConceptText, true
	case "spacer", "blank":
		return entity.ProjectConceptSpacer, true
	case "item", "rectangle", "table", "database", "uml-class", "uml-component", "uml-state", "uml-activity":
		return entity.ProjectConceptItem, true
	case "bend":
		return "", false
	default:
		if len(node.Children) > 0 {
			return entity.ProjectConceptGroup, true
		}
		return entity.ProjectConceptItem, true
	}
}

func projectNodeIdentity(node *entity.Node) string {
	for _, key := range []string{"id", "name", "ref"} {
		if value := strings.TrimSpace(node.Attr(key)); value != "" {
			return value
		}
	}
	return ""
}

func projectNodeName(node *entity.Node, fallback string) string {
	for _, key := range []string{"title", "name", "ref", "id"} {
		if value := strings.TrimSpace(node.Attr(key)); value != "" {
			return value
		}
	}
	if value := strings.TrimSpace(node.Text); value != "" {
		return projectTruncate(value, 256)
	}
	if fallback != "" {
		return fallback
	}
	return node.Tag
}

func projectNodeDetail(node *entity.Node) string {
	keys := make([]string, 0, len(node.Attrs))
	for key := range node.Attrs {
		if !strings.HasPrefix(key, "_") {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys)+2)
	parts = append(parts, node.Tag)
	for _, key := range keys {
		value := strings.TrimSpace(node.Attrs[key])
		if value != "" {
			parts = append(parts, key+"="+value)
		}
	}
	if text := strings.TrimSpace(node.Text); text != "" {
		parts = append(parts, text)
	}
	return projectTruncate(strings.Join(parts, " "), projectMaxDetailBytes)
}

func projectMarkdownSymbols(uri string, source []byte) []entity.ProjectSymbol {
	lines := strings.Split(string(source), "\n")
	type heading struct {
		line  int
		level int
		name  string
	}
	var headings []heading
	for index, line := range lines {
		trimmed := strings.TrimLeft(line, " \t")
		level := 0
		for level < len(trimmed) && level < 6 && trimmed[level] == '#' {
			level++
		}
		if level == 0 || level >= len(trimmed) || trimmed[level] != ' ' {
			continue
		}
		name := strings.TrimSpace(trimmed[level:])
		if name != "" {
			headings = append(headings, heading{line: index, level: level, name: name})
		}
	}
	if len(headings) == 0 {
		name := filepath.Base(projectURIPath(uri))
		if name == "." || name == "" {
			name = uri
		}
		return []entity.ProjectSymbol{{
			Ordinal: 0, ParentOrdinal: -1, ID: "document", Name: name,
			Detail:  projectTruncate(strings.TrimSpace(string(source)), projectMaxDetailBytes),
			Concept: entity.ProjectConceptText, SourceTag: "markdown",
			Position: entity.Position{Line: 1, Column: 1},
		}}
	}
	symbols := make([]entity.ProjectSymbol, 0, len(headings))
	parents := [7]int{-1, -1, -1, -1, -1, -1, -1}
	for index, current := range headings {
		end := len(lines)
		if index+1 < len(headings) {
			end = headings[index+1].line
		}
		parent := -1
		for level := current.level - 1; level >= 1; level-- {
			if parents[level] >= 0 {
				parent = parents[level]
				break
			}
		}
		ordinal := len(symbols)
		symbols = append(symbols, entity.ProjectSymbol{
			Ordinal: ordinal, ParentOrdinal: parent,
			ID: fmt.Sprintf("heading-%d", current.line+1), Name: current.name,
			Detail:  projectTruncate(strings.TrimSpace(strings.Join(lines[current.line+1:end], "\n")), projectMaxDetailBytes),
			Concept: entity.ProjectConceptText, SourceTag: "heading",
			Position: entity.Position{Line: current.line + 1, Column: 1},
		})
		parents[current.level] = ordinal
		for level := current.level + 1; level < len(parents); level++ {
			parents[level] = -1
		}
	}
	return symbols
}

func projectImportSource(uri string) *entity.ImportSource {
	path := projectURIPath(uri)
	if path == "" {
		return nil
	}
	directory := filepath.Dir(path)
	if info, err := os.Stat(directory); err != nil || !info.IsDir() {
		return nil
	}
	return &entity.ImportSource{FS: os.DirFS(directory)}
}

func projectRAGExtension(path string) bool {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".md", ".markdown":
		return true
	default:
		return false
	}
}

func projectIgnoredDirectory(name string) bool {
	switch name {
	case ".git", ".xaligo", ".bin", "node_modules", "target", "dist", "book", "output":
		return true
	default:
		return false
	}
}

func projectDocsRoot(root string) (string, error) {
	if filepath.Base(filepath.Clean(root)) == "docs" {
		return root, nil
	}
	candidate := filepath.Join(root, "docs")
	info, err := os.Stat(candidate)
	if err != nil {
		return "", fmt.Errorf("inspect RAG documentation root %s: %w", candidate, err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("RAG documentation root %s is not a directory", candidate)
	}
	return candidate, nil
}

func projectPathURI(path string) string {
	absolute, err := filepath.Abs(path)
	if err != nil {
		absolute = path
	}
	return (&url.URL{Scheme: "file", Path: filepath.ToSlash(absolute)}).String()
}

func projectURIPath(uri string) string {
	parsed, err := url.Parse(uri)
	if err == nil && parsed.Scheme == "file" {
		return filepath.FromSlash(parsed.Path)
	}
	if err == nil && parsed.Scheme != "" {
		return ""
	}
	return uri
}

func projectDocumentRootURI(uri string) string {
	path := projectURIPath(uri)
	if path == "" {
		return uri
	}
	return projectPathURI(filepath.Dir(path))
}

func projectTruncate(value string, limit int) string {
	if len(value) <= limit {
		return value
	}
	return value[:limit]
}
