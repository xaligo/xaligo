package engine_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestDiffBoxHighlightBecomesDeferredPlanOverlay(t *testing.T) {
	document := parseDiffDocument(t, `<frame version="1" width="320" height="180"><rectangle id="one" title="One" /></frame>`)
	change := entity.StructuralChange{Position: document.Root.Children[0].Position}
	engine.MarkChangesV1EngineDiffDocument(document.Root, []entity.StructuralChange{change}, string(entity.StructuralChangeAdded))
	plan := buildDiffPlan(t, document)

	found := false
	for _, operation := range plan.Ops {
		if operation.Fill != nil && operation.Fill.Color == "DCFCE7" {
			found = true
			if operation.Fill.Transparency <= 0 || operation.Fill.Transparency >= 100 {
				t.Fatalf("highlight transparency = %v", operation.Fill.Transparency)
			}
		}
	}
	if !found {
		t.Fatalf("plan operations = %#v, want added highlight", plan.Ops)
	}
}

func TestDiffConnectionHighlightCopiesResolvedPath(t *testing.T) {
	document := parseDiffDocument(t, `<frame version="1" width="420" height="180" layout="horizontal"><rectangle id="a"/><rectangle id="b"/><connection src="a" dst="b" /></frame>`)
	connection := document.Root.Children[2]
	engine.MarkChangesV1EngineDiffDocument(document.Root, []entity.StructuralChange{{Position: connection.Position}}, string(entity.StructuralChangeRemoved))
	plan := buildDiffPlan(t, document)

	for index, operation := range plan.Ops {
		if operation.Line == nil || operation.Line.Color != "FCA5A5" {
			continue
		}
		if index+1 >= len(plan.Ops) || plan.Ops[index+1].Kind != "line" {
			t.Fatalf("highlight operation is not followed by original: %#v", plan.Ops[index:])
		}
		original := plan.Ops[index+1]
		if len(operation.Points) != len(original.Points) || operation.X != original.X || operation.Y != original.Y || operation.W != original.W || operation.H != original.H {
			t.Fatalf("highlight path = %#v, original = %#v", operation, original)
		}
		if operation.Line.BeginArrowType != "none" || operation.Line.EndArrowType != "none" {
			t.Fatalf("highlight arrowheads = %q/%q", operation.Line.BeginArrowType, operation.Line.EndArrowType)
		}
		return
	}
	t.Fatalf("plan operations = %#v, want removed connector highlight", plan.Ops)
}

func TestDiffEntryChangeHighlightsOnlyItsFrameMetadataTag(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1" width="320" height="180"><metadata><entry key="owner" value="platform" /><entry key="status" value="active" /></metadata></frame>`)
	after := parseDiffDocument(t, `<frame version="1" width="320" height="180"><metadata><entry key="owner" value="security" /><entry key="status" value="active" /></metadata></frame>`)
	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	engine.MarkChangesV1EngineDiffDocument(after.Root, diff.After, string(entity.StructuralChangeAdded))

	scene := buildDiffScene(t, after)
	highlights := frameMetadataDiffHighlights(t, scene)
	if len(highlights) != 1 {
		t.Fatalf("metadata highlights = %#v, want one changed tag", highlights)
	}
	highlight := highlights[0]
	if !strings.Contains(highlight.ID, "metadata-00") || strings.Contains(highlight.ID, "metadata-01") {
		t.Fatalf("highlight ID = %q, want first metadata tag only", highlight.ID)
	}
	if highlight.CustomData == nil || !highlight.CustomData.FrameMetadata || highlight.CustomData.DiffStatus != string(entity.StructuralChangeAdded) {
		t.Fatalf("highlight customData = %#v", highlight.CustomData)
	}

	plan := engine.BuildPlanV1EnginePlanBuild(&scene, entity.PlanOptions{PxPerInch: 96})
	highlightIndex := drawOpIndexByID(plan.Ops, highlight.ID)
	keyIndex := drawOpIndexContainingID(plan.Ops, "metadata-00-key")
	valueIndex := drawOpIndexContainingID(plan.Ops, "metadata-00-value")
	textIndex := drawOpIndexContainingID(plan.Ops, "metadata-00-key-content")
	if keyIndex < 0 || valueIndex < 0 || highlightIndex < 0 || textIndex < 0 || keyIndex >= highlightIndex || valueIndex >= highlightIndex || highlightIndex >= textIndex {
		t.Fatalf("metadata draw order key=%d value=%d highlight=%d text=%d; operations=%#v", keyIndex, valueIndex, highlightIndex, textIndex, plan.Ops)
	}
	op := plan.Ops[highlightIndex]
	if !op.FrontLayer || op.Fill == nil || op.Fill.Color != "DCFCE7" || op.Line == nil || op.Line.Color != "86EFAC" {
		t.Fatalf("highlight operation = %#v", op)
	}
}

func TestDiffMetadataStyleHighlightsEveryMetadataTag(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1" width="320" height="180"><metadata color="#111111"><entry key="owner" value="platform" /><entry key="status" value="active" /></metadata></frame>`)
	after := parseDiffDocument(t, `<frame version="1" width="320" height="180"><metadata color="#222222"><entry key="owner" value="platform" /><entry key="status" value="active" /></metadata></frame>`)
	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	engine.MarkChangesV1EngineDiffDocument(after.Root, diff.After, string(entity.StructuralChangeAdded))

	highlights := frameMetadataDiffHighlights(t, buildDiffScene(t, after))
	if len(highlights) != 2 {
		t.Fatalf("metadata highlights = %#v, want every metadata tag", highlights)
	}
	for _, highlight := range highlights {
		if highlight.CustomData == nil || !highlight.CustomData.FrameMetadata || !highlight.CustomData.DiffHighlight || highlight.CustomData.DiffStatus != string(entity.StructuralChangeAdded) {
			t.Fatalf("highlight customData = %#v", highlight.CustomData)
		}
	}
}

func TestDiffEntryKeyAddAndRemoveHighlightTheAffectedMetadataTag(t *testing.T) {
	tests := []struct {
		name       string
		before     string
		after      string
		beforeSide bool
		wantStatus entity.StructuralChangeKind
		wantTag    string
		wantFill   string
		wantLine   string
	}{
		{
			name:       "key changed",
			before:     `<frame version="1" width="320" height="180"><metadata><entry key="owner" value="platform" /></metadata></frame>`,
			after:      `<frame version="1" width="320" height="180"><metadata><entry key="team" value="platform" /></metadata></frame>`,
			wantStatus: entity.StructuralChangeAdded,
			wantTag:    "metadata-00",
			wantFill:   "DCFCE7",
			wantLine:   "86EFAC",
		},
		{
			name:       "entry added",
			before:     `<frame version="1" width="320" height="180"><metadata><entry key="owner" value="platform" /></metadata></frame>`,
			after:      `<frame version="1" width="320" height="180"><metadata><entry key="owner" value="platform" /><entry key="status" value="active" /></metadata></frame>`,
			wantStatus: entity.StructuralChangeAdded,
			wantTag:    "metadata-01",
			wantFill:   "DCFCE7",
			wantLine:   "86EFAC",
		},
		{
			name:       "entry removed",
			before:     `<frame version="1" width="320" height="180"><metadata><entry key="owner" value="platform" /><entry key="status" value="active" /></metadata></frame>`,
			after:      `<frame version="1" width="320" height="180"><metadata><entry key="owner" value="platform" /></metadata></frame>`,
			beforeSide: true,
			wantStatus: entity.StructuralChangeRemoved,
			wantTag:    "metadata-01",
			wantFill:   "FEE2E2",
			wantLine:   "FCA5A5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			before := parseDiffDocument(t, test.before)
			after := parseDiffDocument(t, test.after)
			diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
			document := after
			changes := diff.After
			if test.beforeSide {
				document = before
				changes = diff.Before
			}
			engine.MarkChangesV1EngineDiffDocument(document.Root, changes, string(test.wantStatus))
			scene := buildDiffScene(t, document)
			highlights := frameMetadataDiffHighlights(t, scene)
			if len(highlights) != 1 || !strings.Contains(highlights[0].ID, test.wantTag) {
				t.Fatalf("metadata highlights = %#v, want only %s", highlights, test.wantTag)
			}
			plan := engine.BuildPlanV1EnginePlanBuild(&scene, entity.PlanOptions{PxPerInch: 96})
			index := drawOpIndexByID(plan.Ops, highlights[0].ID)
			if index < 0 || plan.Ops[index].Fill == nil || plan.Ops[index].Fill.Color != test.wantFill || plan.Ops[index].Line == nil || plan.Ops[index].Line.Color != test.wantLine {
				t.Fatalf("highlight operation index=%d operations=%#v", index, plan.Ops)
			}
		})
	}
}

func buildDiffPlan(t *testing.T, document entity.Document) entity.Plan {
	t.Helper()
	scene := buildDiffScene(t, document)
	return engine.BuildPlanV1EnginePlanBuild(&scene, entity.PlanOptions{PxPerInch: 96})
}

func buildDiffScene(t *testing.T, document entity.Document) entity.PresentationScene {
	t.Helper()
	root, err := engine.BuildV1EngineLayoutBuild(document)
	if err != nil {
		t.Fatal(err)
	}
	sceneJSON, err := engine.BuildJSONV1EngineSceneBuild(root, "", "", "", 32, engine.CollectConnectionNodesV1EngineSceneConnection(document.Root), nil, nil, engine.SceneDependenciesV1EngineSceneTypes{})
	if err != nil {
		t.Fatal(err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		t.Fatal(err)
	}
	return scene
}

func frameMetadataDiffHighlights(t *testing.T, scene entity.PresentationScene) []entity.Element {
	t.Helper()
	result := make([]entity.Element, 0)
	for _, element := range scene.Elements {
		if element.CustomData != nil && element.CustomData.FrameMetadata && element.CustomData.DiffHighlight {
			result = append(result, element)
		}
	}
	return result
}

func drawOpIndexByID(operations []entity.DrawOp, id string) int {
	for index, operation := range operations {
		if operation.ID == id {
			return index
		}
	}
	return -1
}

func drawOpIndexContainingID(operations []entity.DrawOp, fragment string) int {
	for index, operation := range operations {
		if strings.Contains(operation.ID, fragment) {
			return index
		}
	}
	return -1
}

func TestDiffHighlightInternalAttributeIsNotCompared(t *testing.T) {
	before := parseDiffDocument(t, `<frame version="1"><rectangle id="one" /></frame>`)
	after := parseDiffDocument(t, `<frame version="1"><rectangle id="one" /></frame>`)
	before.Root.Children[0].Attrs["_xaligoDiffStatus"] = "removed"
	diff := engine.DiffDocumentsV1EngineDiffDocument(before, after)
	if diff.AddedCount+diff.RemovedCount+diff.ModifiedCount != 0 {
		t.Fatalf("diff = %#v", diff)
	}
}
