package engine_test

import (
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestStructuralDiffIgnoresSourceFormattingAndImplicitV1(t *testing.T) {
	before := parseDiffDocument(t, `<frame id="root" width="320" height="180"><rectangle id="one" title="One" /></frame>`)
	after := parseDiffDocument(t, `
<frame height="180" id="root" version="1" width="320">
  <rectangle title="One" id="one"></rectangle>
</frame>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.AddedCount != 0 || diff.RemovedCount != 0 || diff.ModifiedCount != 0 {
		t.Fatalf("diff = %#v, want no structural changes", diff)
	}
}

func TestStructuralDiffIgnoresCanonicalDocumentRootVersion(t *testing.T) {
	before := parseDiffDocument(t, `<xaligo><frames><frame id="page" /></frames></xaligo>`)
	after := parseDiffDocument(t, `<xaligo version="1"><frames><frame id="page" /></frames></xaligo>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.AddedCount != 0 || diff.RemovedCount != 0 || diff.ModifiedCount != 0 {
		t.Fatalf("diff = %#v, want no structural changes", diff)
	}
}

func TestStructuralDiffIncludesIdentifiedChildFrameVersionEvenWhenValueIsOne(t *testing.T) {
	before := parseDiffDocument(t, `<xaligo version="1"><frames><frame id="page" /></frames></xaligo>`)
	after := parseDiffDocument(t, `<xaligo version="1"><frames><frame id="page" version="1" /></frames></xaligo>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.AddedCount != 0 || diff.RemovedCount != 0 || diff.ModifiedCount != 1 {
		t.Fatalf("diff = %#v, want one modified frame", diff)
	}
	if len(diff.Before) != 1 || len(diff.After) != 1 || diff.Before[0].Identity != "id=page" || diff.After[0].Identity != "id=page" {
		t.Fatalf("diff = %#v, want only identified child frame id=page modified", diff)
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

func TestStructuralDiffShorthandChangeDoesNotMarkWholeFrame(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1"><rectangle id="a"/><rectangle id="b"/><rectangle id="c"/>a --- b</frame>`)
	after := parseDiffDocument(t, `<frame version="1"><rectangle id="a"/><rectangle id="b"/><rectangle id="c"/>a --- c</frame>`)

	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.RemovedCount != 1 || diff.AddedCount != 1 || diff.ModifiedCount != 0 {
		t.Fatalf("diff = %#v, want one shorthand connector replacement", diff)
	}
	if len(diff.Before) != 1 || diff.Before[0].Tag != "connection" || len(diff.After) != 1 || diff.After[0].Tag != "connection" {
		t.Fatalf("diff = %#v, want only connection changes", diff)
	}
}

func TestStructuralDiffDetectsFrameMetadataChanges(t *testing.T) {
	tests := []struct {
		name        string
		before      string
		after       string
		wantTag     string
		wantAdded   int
		wantRemoved int
		wantChanged int
	}{
		{
			name:        "metadata style",
			before:      `<frame version="1"><metadata color="#111111"><entry key="owner" value="platform" /></metadata></frame>`,
			after:       `<frame version="1"><metadata color="#222222"><entry key="owner" value="platform" /></metadata></frame>`,
			wantTag:     "metadata",
			wantChanged: 1,
		},
		{
			name:        "entry key",
			before:      `<frame version="1"><metadata><entry key="owner" value="platform" /></metadata></frame>`,
			after:       `<frame version="1"><metadata><entry key="team" value="platform" /></metadata></frame>`,
			wantTag:     "entry",
			wantChanged: 1,
		},
		{
			name:        "entry value",
			before:      `<frame version="1"><metadata><entry key="owner" value="platform" /></metadata></frame>`,
			after:       `<frame version="1"><metadata><entry key="owner" value="security" /></metadata></frame>`,
			wantTag:     "entry",
			wantChanged: 1,
		},
		{
			name:      "entry added",
			before:    `<frame version="1"><metadata><entry key="owner" value="platform" /></metadata></frame>`,
			after:     `<frame version="1"><metadata><entry key="owner" value="platform" /><entry key="status" value="active" /></metadata></frame>`,
			wantTag:   "entry",
			wantAdded: 1,
		},
		{
			name:        "entry removed",
			before:      `<frame version="1"><metadata><entry key="owner" value="platform" /><entry key="status" value="active" /></metadata></frame>`,
			after:       `<frame version="1"><metadata><entry key="owner" value="platform" /></metadata></frame>`,
			wantTag:     "entry",
			wantRemoved: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before := parseDiffDocument(t, test.before)
			after := parseDiffDocument(t, test.after)
			diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
			if diff.AddedCount != test.wantAdded || diff.RemovedCount != test.wantRemoved || diff.ModifiedCount != test.wantChanged {
				t.Fatalf("diff = %#v", diff)
			}
			changes := append(append([]entity.StructuralChange{}, diff.Before...), diff.After...)
			if len(changes) == 0 {
				t.Fatal("changes = nil")
			}
			for _, change := range changes {
				if change.Tag != test.wantTag {
					t.Fatalf("change = %#v, want tag %q", change, test.wantTag)
				}
			}
		})
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
