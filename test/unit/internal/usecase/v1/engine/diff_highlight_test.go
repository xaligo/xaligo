package engine_test

import (
	"encoding/json"
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

func buildDiffPlan(t *testing.T, document entity.Document) entity.Plan {
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
	return engine.BuildPlanV1EnginePlanBuild(&scene, entity.PlanOptions{PxPerInch: 96})
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
