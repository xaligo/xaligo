package repository_test

import (
	"encoding/json"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
)

func TestRenderBuildsNestedNodesAndEdges(t *testing.T) {
	scene := []byte(`{
  "elements": [
    {"id":"paper-frame","type":"frame","width":400,"height":200},
    {"id":"xyFlowGroup-rect","type":"rectangle","x":10,"y":10,"width":380,"height":180,"strokeColor":"#64748b","strokeWidth":1},
    {"id":"xyFlowGroup-icon","type":"image","x":10,"y":10,"width":32,"height":32,"fileId":"fg"},
    {"id":"a-item","type":"image","x":50,"y":80,"width":32,"height":32,"fileId":"fa"},
    {"id":"a-item-lbl","type":"text","text":"Web"},
    {"id":"b-item","type":"image","x":300,"y":80,"width":32,"height":32,"fileId":"fb"},
    {"id":"b-item-lbl","type":"text","text":"DB"},
    {"id":"edge","type":"arrow","strokeColor":"#2563eb","strokeWidth":2,
      "startBinding":{"elementId":"a-item","fixedPoint":[1,0.5]},
      "endBinding":{"elementId":"b-item-lbl","fixedPoint":[0,0.5]},
      "customData":{"xaligoConnectorKind":"traffic","xaligoConnectorEndArrowhead":"triangle"}}
  ],
  "files":{"fa":{"dataURL":"data:image/svg+xml;base64,QQ=="},"fg":{"dataURL":"data:image/svg+xml;base64,Rw=="}},
  "appState":{"viewBackgroundColor":"#ffffff"}
}`)
	out, err := repository.NewXYFlowRepository().Render(scene)
	if err != nil {
		t.Fatal(err)
	}
	var document entity.XYFlowDocument
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Nodes) != 3 || len(document.Edges) != 1 {
		t.Fatalf("nodes=%d edges=%d: %s", len(document.Nodes), len(document.Edges), out)
	}
	if document.Nodes[0].ID != "xyFlowGroup-rect" || document.Nodes[0].Data["icon"] != "data:image/svg+xml;base64,Rw==" {
		t.Fatalf("xyFlowGroup icon data = %#v", document.Nodes[0])
	}
	var item entity.XYFlowNode
	for _, node := range document.Nodes {
		if node.ID == "a-item" {
			item = node
		}
	}
	if item.ParentID != "xyFlowGroup-rect" || item.XYFlowPosition.X != 40 || item.XYFlowPosition.Y != 70 || item.Data["label"] != "Web" {
		t.Fatalf("item = %#v", item)
	}
	edge := document.Edges[0]
	if edge.Source != "a-item" || edge.Target != "b-item" || edge.SourceHandle != "right" || edge.TargetHandle != "left" || edge.ZIndex != 2 || edge.MarkerEnd == nil {
		t.Fatalf("edge = %#v", edge)
	}
}

func TestRenderKeepsEdgesForEveryEmittedNodeKind(t *testing.T) {
	scene := []byte(`{
  "elements": [
    {"id":"paper-frame","type":"frame","width":600,"height":400},
    {"id":"page-a","type":"frame","x":0,"y":0,"width":600,"height":400},
    {"id":"group-rect","type":"rectangle","x":20,"y":20,"width":250,"height":300,"customData":{"xaligoGroupBorder":true}},
    {"id":"box-rect","type":"rectangle","x":300,"y":50,"width":250,"height":200},
    {"id":"port-rect","type":"rectangle","x":300,"y":120,"width":40,"height":30},
    {"id":"service-item","type":"image","x":100,"y":100,"width":32,"height":32,"fileId":"service"},
    {"id":"service-item-lbl","type":"text","text":"Service"},
    {"id":"service-anchor-bg-00-00","type":"rectangle","x":95,"y":95,"width":8,"height":8,"customData":{"xaligoAnchorBackground":true}},
    {"id":"group-to-box","type":"arrow","strokeColor":"#2563eb","strokeWidth":2,
      "startBinding":{"elementId":"group-rect","fixedPoint":[1,0.3]},
      "endBinding":{"elementId":"box-rect","fixedPoint":[0,0.7]},
      "customData":{"xaligoConnectorKind":"traffic","xaligoConnectorStartArrowhead":"oval","xaligoConnectorEndArrowhead":"triangle","xaligoConnectorBends":"270,100 290,100","xaligoConnectorScale":2,"xaligoConnectorGrid":8,"xaligoConnectorSrcAnchor":true,"xaligoConnectorDstAnchor":true}},
    {"id":"port-to-frame","type":"arrow","strokeColor":"#1e1e1e","strokeWidth":1,
      "startBinding":{"elementId":"port-rect","fixedPoint":[0,0.5]},
      "endBinding":{"elementId":"page-a","fixedPoint":[1,0.5]}},
    {"id":"item-to-group","type":"arrow","strokeColor":"#1e1e1e","strokeWidth":1,
      "startBinding":{"elementId":"service-item-lbl","fixedPoint":[1,0.5]},
      "endBinding":{"elementId":"group-rect","fixedPoint":[0,0.5]}},
    {"id":"cross-source-stub","type":"arrow","strokeColor":"#64748b","strokeWidth":1,
      "startBinding":{"elementId":"box-rect","fixedPoint":[1,0.4]},
      "customData":{"xaligoConnectorKind":"route","xaligoCrossFrame":true,"xaligoSourceFrame":"a","xaligoDestinationFrame":"b","xaligoConnectorLogicalId":"logical-cross","xaligoConnectorSourceElementId":"box-rect","xaligoConnectorDestinationElementId":"service-item-lbl"}},
    {"id":"cross-target-stub","type":"arrow","strokeColor":"#64748b","strokeWidth":1,
      "endBinding":{"elementId":"service-item-lbl","fixedPoint":[0,0.6]},
      "customData":{"xaligoConnectorKind":"route","xaligoCrossFrame":true,"xaligoSourceFrame":"a","xaligoDestinationFrame":"b","xaligoConnectorLogicalId":"logical-cross","xaligoConnectorSourceElementId":"box-rect","xaligoConnectorDestinationElementId":"service-item-lbl"}}
  ],
  "files":{"service":{"dataURL":"data:image/svg+xml;base64,Uw=="}},
  "appState":{"viewBackgroundColor":"#ffffff"}
}`)
	out, err := repository.NewXYFlowRepository().Render(scene)
	if err != nil {
		t.Fatal(err)
	}
	var document entity.XYFlowDocument
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Nodes) != 5 {
		t.Fatalf("nodes=%d, want frame/group/rectangle/port/item only: %s", len(document.Nodes), out)
	}
	nodeIDs := map[string]bool{}
	for _, node := range document.Nodes {
		nodeIDs[node.ID] = true
	}
	if nodeIDs["service-anchor-bg-00-00"] {
		t.Fatalf("anchor background was emitted as a node: %s", out)
	}
	if len(document.Edges) != 4 {
		t.Fatalf("edges=%d, want three local and one logical cross-frame edge: %s", len(document.Edges), out)
	}
	edges := map[string]entity.XYFlowEdge{}
	for _, edge := range document.Edges {
		edges[edge.ID] = edge
		if !nodeIDs[edge.Source] || !nodeIDs[edge.Target] {
			t.Fatalf("edge %q points outside emitted nodes: %#v", edge.ID, edge)
		}
	}
	metadata := edges["group-to-box"]
	if metadata.Source != "group-rect" || metadata.Target != "box-rect" || metadata.SourceHandle != "right" || metadata.TargetHandle != "left" {
		t.Fatalf("group-to-box edge = %#v", metadata)
	}
	if metadata.Data["bends"] != "270,100 290,100" || metadata.Data["scale"] != float64(2) || metadata.Data["grid"] != float64(8) || metadata.Data["sourceAnchorExplicit"] != true || metadata.Data["targetAnchorExplicit"] != true {
		t.Fatalf("connector metadata = %#v", metadata.Data)
	}
	if got := metadata.Data["sourceFixedPoint"]; !equalFloatSlice(got, []float64{1, 0.3}) {
		t.Fatalf("sourceFixedPoint = %#v", got)
	}
	if got := metadata.Data["targetFixedPoint"]; !equalFloatSlice(got, []float64{0, 0.7}) {
		t.Fatalf("targetFixedPoint = %#v", got)
	}
	logical := edges["logical-cross"]
	if logical.Source != "box-rect" || logical.Target != "service-item" || logical.SourceHandle != "right" || logical.TargetHandle != "left" || logical.Data["crossFrame"] != true {
		t.Fatalf("logical cross-frame edge = %#v", logical)
	}
	if _, exists := edges["cross-source-stub"]; exists {
		t.Fatalf("source stub was not deduplicated: %s", out)
	}
	if _, exists := edges["cross-target-stub"]; exists {
		t.Fatalf("target stub was not deduplicated: %s", out)
	}
}

func TestRenderPrefersSemanticParentsForEqualAndOverflowingNodes(t *testing.T) {
	scene := []byte(`{
  "elements": [
    {"id":"paper-frame","type":"frame","width":240,"height":240},
    {"id":"page","type":"frame","x":0,"y":0,"width":100,"height":100,
      "customData":{"xaligoSemanticElementKind":"frame"}},
    {"id":"group","type":"rectangle","x":0,"y":0,"width":100,"height":100,
      "customData":{"xaligoSemanticElementKind":"group","xaligoSemanticParentElementId":"page"}},
    {"id":"child","type":"rectangle","x":80,"y":80,"width":100,"height":100,
      "customData":{"xaligoSemanticElementKind":"rectangle","xaligoSemanticParentElementId":"group"}},
    {"id":"service-item","type":"image","x":170,"y":170,"width":32,"height":32,
      "customData":{"xaligoSemanticElementKind":"item","xaligoSemanticParentElementId":"child"}}
  ],
  "files":{},
  "appState":{"viewBackgroundColor":"#ffffff"}
}`)
	out, err := repository.NewXYFlowRepository().Render(scene)
	if err != nil {
		t.Fatal(err)
	}
	var document entity.XYFlowDocument
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Nodes) != 4 {
		t.Fatalf("nodes=%d, want page/group/child/item: %s", len(document.Nodes), out)
	}
	wantParents := map[string]string{
		"page":         "",
		"group":        "page",
		"child":        "group",
		"service-item": "child",
	}
	wantKinds := map[string]string{
		"page":         "frame",
		"group":        "group",
		"child":        "rectangle",
		"service-item": "item",
	}
	for _, node := range document.Nodes {
		if want, ok := wantParents[node.ID]; !ok {
			t.Fatalf("unexpected node %#v", node)
		} else if node.ParentID != want {
			t.Fatalf("node %q parent=%q, want %q: %s", node.ID, node.ParentID, want, out)
		}
		if node.Data["semanticKind"] != wantKinds[node.ID] {
			t.Fatalf("node %q semantic kind=%#v, want %q", node.ID, node.Data["semanticKind"], wantKinds[node.ID])
		}
	}
	child := xyFlowNodeByID(document.Nodes, "child")
	if child.XYFlowPosition.X != 80 || child.XYFlowPosition.Y != 80 {
		t.Fatalf("overflowing child position = %#v", child.XYFlowPosition)
	}
	if child.Extent != "" {
		t.Fatalf("overflowing child extent=%q, want no parent clamp", child.Extent)
	}
}

func xyFlowNodeByID(nodes []entity.XYFlowNode, id string) entity.XYFlowNode {
	for _, node := range nodes {
		if node.ID == id {
			return node
		}
	}
	return entity.XYFlowNode{}
}

func equalFloatSlice(value any, want []float64) bool {
	got, ok := value.([]any)
	if !ok || len(got) != len(want) {
		return false
	}
	for i := range want {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
