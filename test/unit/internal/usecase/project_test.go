package usecase_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	projectrepository "github.com/xaligo/xaligo/internal/repository/project"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

type fakeProjectIndex struct {
	checksums map[string][32]byte
	analyses  map[string]entity.ProjectAnalysis
	roots     map[string]string
}

func newFakeProjectIndex() *fakeProjectIndex {
	return &fakeProjectIndex{
		checksums: map[string][32]byte{},
		analyses:  map[string]entity.ProjectAnalysis{},
		roots:     map[string]string{},
	}
}

func (rcvr *fakeProjectIndex) Put(_ context.Context, root string, analysis entity.ProjectAnalysis) (bool, error) {
	changed := rcvr.checksums[analysis.URI] != analysis.Checksum
	rcvr.checksums[analysis.URI] = analysis.Checksum
	rcvr.analyses[analysis.URI] = analysis
	rcvr.roots[analysis.URI] = root
	return changed, nil
}

func (rcvr *fakeProjectIndex) Search(context.Context, string, int) ([]entity.ProjectSearchResult, error) {
	return nil, nil
}

func (rcvr *fakeProjectIndex) Symbols(_ context.Context, uri string) ([]entity.ProjectSymbol, error) {
	return rcvr.analyses[uri].Symbols, nil
}

func (rcvr *fakeProjectIndex) Prune(_ context.Context, root string, keep []string) (int, error) {
	kept := map[string]bool{}
	for _, uri := range keep {
		kept[uri] = true
	}
	removed := 0
	for uri, candidateRoot := range rcvr.roots {
		if candidateRoot == root && !kept[uri] {
			delete(rcvr.roots, uri)
			delete(rcvr.analyses, uri)
			delete(rcvr.checksums, uri)
			removed++
		}
	}
	return removed, nil
}

func (*fakeProjectIndex) Close() error { return nil }

var _ projectrepository.IndexRepository = (*fakeProjectIndex)(nil)

type fakeProjectEngine struct {
	resolveCalls int
	resolveErr   error
}

func (rcvr *fakeProjectEngine) Resolve(context.Context, entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error) {
	rcvr.resolveCalls++
	return entity.EngineResolvedDocument{}, rcvr.resolveErr
}

func (*fakeProjectEngine) RenderSVG(context.Context, entity.EngineDocumentSpec) ([]byte, error) {
	return nil, nil
}

func (*fakeProjectEngine) NormalizeSVG(context.Context, []byte) (entity.EngineSVG, error) {
	return entity.EngineSVG{}, nil
}

var _ v2.EngineUsecase = (*fakeProjectEngine)(nil)

func newProjectUsecase(index projectrepository.IndexRepository, engine v2.EngineUsecase) usecase.ProjectUsecase {
	if engine == nil {
		engine = &fakeProjectEngine{}
	}
	return usecase.NewProjectUsecase(index, v2.NewFrontendUsecase(), engine)
}

func TestProjectAnalyzeBuildsGenericConceptsFromComplexXAL(t *testing.T) {
	path := filepath.Join("..", "..", "..", "..", "docs", "src", "examples", "samples", "complex-hybrid-architecture.xal")
	source, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	project := newProjectUsecase(newFakeProjectIndex(), nil)
	analysis, err := project.Analyze(context.Background(), "file:///complex-hybrid-architecture.xal", source)
	if err != nil {
		t.Fatal(err)
	}
	if analysis.Kind != entity.ProjectDocumentXAL || len(analysis.Diagnostics) != 0 {
		t.Fatalf("analysis kind=%q diagnostics=%#v", analysis.Kind, analysis.Diagnostics)
	}
	counts := map[entity.ProjectConcept]int{}
	for _, symbol := range analysis.Symbols {
		counts[symbol.Concept]++
		if symbol.SourceTag == "aws-cloud" && symbol.Concept != entity.ProjectConceptGroup {
			t.Fatalf("domain tag did not normalize to group: %#v", symbol)
		}
	}
	if counts[entity.ProjectConceptFrame] != 1 || counts[entity.ProjectConceptPort] != 5 || counts[entity.ProjectConceptLine] != 36 {
		t.Fatalf("generic concept counts = %#v", counts)
	}
}

func TestProjectAnalyzeBuildsGenericConceptsFromComplexV2XAL(t *testing.T) {
	path := filepath.Join("..", "..", "..", "..", "docs", "src", "examples", "samples", "complex-hybrid-architecture-v2.xal")
	source, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	engine := &fakeProjectEngine{}
	project := newProjectUsecase(newFakeProjectIndex(), engine)
	analysis, err := project.Analyze(context.Background(), "file:///complex-hybrid-architecture-v2.xal", source)
	if err != nil {
		t.Fatal(err)
	}
	if analysis.Kind != entity.ProjectDocumentXAL || len(analysis.Diagnostics) != 0 {
		t.Fatalf("analysis kind=%q diagnostics=%#v", analysis.Kind, analysis.Diagnostics)
	}
	counts := map[entity.ProjectConcept]int{}
	wantRow := "xaligo[0]/frames[0]/frame#complex-hybrid-architecture/row[0]"
	wantConnection := "xaligo[0]/frames[0]/frame#complex-hybrid-architecture/connections[1]/connection[0]"
	foundRow, foundConnection := false, false
	for _, symbol := range analysis.Symbols {
		counts[symbol.Concept]++
		foundRow = foundRow || symbol.ID == wantRow
		foundConnection = foundConnection || symbol.ID == wantConnection
		if symbol.SourceTag == "aws-cloud" && symbol.Concept != entity.ProjectConceptGroup {
			t.Fatalf("domain tag did not normalize to group: %#v", symbol)
		}
	}
	if counts[entity.ProjectConceptFrame] != 1 || counts[entity.ProjectConceptPort] != 5 || counts[entity.ProjectConceptLine] != 38 {
		t.Fatalf("generic concept counts = %#v", counts)
	}
	if engine.resolveCalls != 1 {
		t.Fatalf("V2 resolve calls = %d, want 1", engine.resolveCalls)
	}
	if !foundRow || !foundConnection {
		t.Fatalf("V2 anonymous semantic paths = row %t connection %t", foundRow, foundConnection)
	}
}

func TestProjectAnalyzeMapsV2EngineDiagnosticsToSource(t *testing.T) {
	source := []byte(`<xaligo version="2">
<frames><frame id="page" width="320" height="180">
<item id="api">API</item>
</frame></frames></xaligo>`)
	engine := &fakeProjectEngine{resolveErr: &entity.EngineDiagnosticError{Diagnostic: entity.EngineDiagnostic{
		Code: "XAL-E2001", Severity: "error", Stage: "calculate", ElementID: "api", SpanID: 2, Message: "invalid API geometry",
	}}}
	analysis, err := newProjectUsecase(newFakeProjectIndex(), engine).Analyze(context.Background(), "file:///diagnostic.xal", source)
	if err != nil {
		t.Fatal(err)
	}
	if len(analysis.Symbols) != 2 || len(analysis.Diagnostics) != 1 {
		t.Fatalf("V2 analysis = symbols %#v diagnostics %#v", analysis.Symbols, analysis.Diagnostics)
	}
	diagnostic := analysis.Diagnostics[0]
	if diagnostic.Code != "XAL-E2001" || diagnostic.Element != "api" || diagnostic.Line != 3 || diagnostic.Column != 1 {
		t.Fatalf("source diagnostic = %#v", diagnostic)
	}
}

func TestProjectIndexDocumentKeepsLastGoodV2Symbols(t *testing.T) {
	const uri = "file:///project-v2.xal"
	index := newFakeProjectIndex()
	engine := &fakeProjectEngine{}
	project := newProjectUsecase(index, engine)
	valid := []byte(`<xaligo version="2"><frames><frame id="page" width="320" height="180"><item id="api">API</item></frame></frames></xaligo>`)
	first, changed, err := project.IndexDocument(context.Background(), uri, valid)
	if err != nil {
		t.Fatal(err)
	}
	if !changed || len(first.Diagnostics) != 0 || len(first.Symbols) != 2 {
		t.Fatalf("first V2 index = changed %t analysis %#v", changed, first)
	}
	invalid := []byte(`<xaligo version="2"><frames><frame id="page">`)
	second, changed, err := project.IndexDocument(context.Background(), uri, invalid)
	if err != nil {
		t.Fatal(err)
	}
	if changed || len(second.Diagnostics) != 1 {
		t.Fatalf("invalid V2 index = changed %t diagnostics %#v", changed, second.Diagnostics)
	}
	persisted := index.analyses[uri]
	if persisted.Checksum != first.Checksum || len(persisted.Symbols) != len(first.Symbols) {
		t.Fatalf("last-good V2 symbols were replaced: %#v", persisted)
	}
}

func TestProjectRAGIndexSeedsOnlyDocsMarkdown(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "docs", "nested"), 0o755); err != nil {
		t.Fatal(err)
	}
	files := map[string]string{
		filepath.Join(root, "docs", "guide.md"):                "# Guide\nDatabase architecture",
		filepath.Join(root, "docs", "nested", "more.markdown"): "# More\nRouting",
		filepath.Join(root, "docs", "sample.xal"):              `<xaligo version="1"><frames><frame id="sample"><blank /></frame></frames></xaligo>`,
		filepath.Join(root, "README.md"):                       "# Outside docs",
	}
	for path, source := range files {
		if err := os.WriteFile(path, []byte(source), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	index := newFakeProjectIndex()
	project := newProjectUsecase(index, nil)
	first, err := project.Index(context.Background(), root)
	if err != nil {
		t.Fatal(err)
	}
	if first.Scanned != 2 || first.Indexed != 2 || first.Unchanged != 0 || len(index.analyses) != 2 {
		t.Fatalf("first docs-only index = %#v analyses=%d", first, len(index.analyses))
	}
	for _, analysis := range index.analyses {
		if analysis.Kind != entity.ProjectDocumentMarkdown {
			t.Fatalf("non-Markdown source was seeded: %#v", analysis)
		}
	}
	second, err := project.Index(context.Background(), filepath.Join(root, "docs"))
	if err != nil {
		t.Fatal(err)
	}
	if second.Indexed != 0 || second.Unchanged != 2 {
		t.Fatalf("incremental index = %#v", second)
	}
}
