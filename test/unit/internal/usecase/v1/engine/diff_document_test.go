package engine_test

import (
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestStructuralDiffIgnoresSourceFormattingAndImplicitV1(t *testing.T) {
	before := parseDiffDocument(t, `<frame width="320" height="180"><rectangle id="one" title="One" /></frame>`)
	after := parseDiffDocument(t, `
<frame height="180" version="1" width="320">
  <rectangle title="One" id="one"></rectangle>
</frame>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.AddedCount != 0 || diff.RemovedCount != 0 || diff.ModifiedCount != 0 {
		t.Fatalf("diff = %#v, want no structural changes", diff)
	}
}

func TestStructuralDiffInsertionDoesNotCascadeToFollowingSiblings(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1"><rectangle id="one" /><rectangle id="two" /></frame>`)
	after := parseDiffDocument(t, `<frame version="1"><rectangle id="new" /><rectangle id="one" /><rectangle id="two" /></frame>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.AddedCount != 1 || diff.RemovedCount != 0 || diff.ModifiedCount != 0 || len(diff.After) != 1 || diff.After[0].Identity != "id=new" {
		t.Fatalf("diff = %#v, want only id=new added", diff)
	}
}

func TestStructuralDiffReportsAttributeChangeOnBothSides(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1"><rectangle id="one" title="Old" /></frame>`)
	after := parseDiffDocument(t, `<frame version="1"><rectangle id="one" title="New" /></frame>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.ModifiedCount != 1 || len(diff.Before) != 1 || len(diff.After) != 1 {
		t.Fatalf("diff = %#v, want one modified pair", diff)
	}
	if diff.Before[0].Kind != entity.StructuralChangeModified || diff.After[0].Kind != entity.StructuralChangeModified {
		t.Fatalf("diff = %#v, want modified on both sides", diff)
	}
}

func TestStructuralDiffTreatsDistinctExplicitIdentitiesAsReplacement(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1"><rectangle id="removed" /></frame>`)
	after := parseDiffDocument(t, `<frame version="1"><rectangle id="added" /></frame>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.RemovedCount != 1 || diff.AddedCount != 1 || diff.ModifiedCount != 0 {
		t.Fatalf("diff = %#v, want one removed and one added", diff)
	}
}

func TestStructuralDiffDetectsReparentedIdentifiedNode(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1"><container id="left"><rectangle id="one" /></container><container id="right" /></frame>`)
	after := parseDiffDocument(t, `<frame version="1"><container id="left" /><container id="right"><rectangle id="one" /></container></frame>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.ModifiedCount != 1 || diff.Before[0].Identity != "id=one" || diff.After[0].Identity != "id=one" {
		t.Fatalf("diff = %#v, want reparented id=one", diff)
	}
}

func TestStructuralDiffMarksConnectionWhenNestedBendChanges(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1"><rectangle id="a"/><rectangle id="b"/><connection src="a" dst="b"><bend x="10" y="20"/></connection></frame>`)
	after := parseDiffDocument(t, `<frame version="1"><rectangle id="a"/><rectangle id="b"/><connection src="a" dst="b"><bend x="30" y="20"/></connection></frame>`)
	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)

	engine.MarkChangesV1EngineDiffDocument(before.Root, diff.Before, string(entity.StructuralChangeRemoved))
	engine.MarkChangesV1EngineDiffDocument(after.Root, diff.After, string(entity.StructuralChangeAdded))
	beforeConnection := before.Root.Children[2]
	afterConnection := after.Root.Children[2]
	if beforeConnection.Attrs["_xaligoDiffStatus"] != "removed" || afterConnection.Attrs["_xaligoDiffStatus"] != "added" {
		t.Fatalf("connection marks = %q/%q", beforeConnection.Attrs["_xaligoDiffStatus"], afterConnection.Attrs["_xaligoDiffStatus"])
	}
}

func parseDiffDocument(t *testing.T, source string) entity.Document {
	t.Helper()
	document, err := engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatal(err)
	}
	return document
}
