package usecase_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	projectrepository "github.com/xaligo/xaligo/internal/repository/project"
	"github.com/xaligo/xaligo/internal/usecase"
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

func TestProjectAnalyzeBuildsGenericConceptsFromComplexXAL(t *testing.T) {
	path := filepath.Join("..", "..", "..", "..", "docs", "src", "examples", "samples", "complex-hybrid-architecture.xal")
	source, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	project := usecase.NewProjectUsecase(newFakeProjectIndex())
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
	project := usecase.NewProjectUsecase(index)
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
