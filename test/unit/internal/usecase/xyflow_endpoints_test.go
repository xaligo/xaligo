package usecase_test

import (
	"context"
	"encoding/json"
	"testing"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestRenderXYFlowPreservesV1EndpointKinds(t *testing.T) {
	source := []byte(`<frames gap="48">
  <frame id="overview" width="640" height="360" layout="horizontal" gap="20">
    <generic-group id="group" title="Group" width="240" height="300">
      <item id="1" name="service" />
    </generic-group>
    <rectangle id="box" title="Box" width="280" height="300">
      <port id="port" title="Port" side="left" width="40" height="24" />
    </rectangle>
    <connections>
      <connection src="group" dst="box" kind="traffic" src-anchor="right-2" dst-anchor="left-4" scale="1" grid="8">
        <bend x="250" y="120" />
      </connection>
      <connection src="port" dst="service" />
      <connection src="overview" dst="box" />
      <connection src="box" dst="remote" kind="route" src-anchor="right-5" dst-anchor="left-1" scale="2" grid="10">
        <bend x="330" y="140" />
      </connection>
    </connections>
  </frame>
  <frame id="detail" width="640" height="360">
    <rectangle id="remote" title="Remote" width="200" height="120" />
  </frame>
</frames>`)
	out, err := newUsecase().RenderXYFlow(context.Background(), source, entity.RenderOptions{
		Format: usecase.FormatXYFlow,
		Theme:  "light",
		Assets: &entity.AssetSource{
			FS:            awsassets.Assets,
			CatalogCSV:    awsassets.CatalogCSV,
			GroupIconsDir: awsassets.GroupIconsDir,
			ItemIconSize:  32,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	var document entity.XYFlowDocument
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Nodes) != 7 {
		t.Fatalf("nodes=%d, want two frames, group, two rectangles, port, and item: %s", len(document.Nodes), out)
	}
	if len(document.Edges) != 4 {
		t.Fatalf("edges=%d, want three local edges and one deduplicated cross-frame edge: %s", len(document.Edges), out)
	}
	if document.Width != 1328 || document.Height != 360 {
		t.Fatalf("multi-frame document size = %.0fx%.0f, want 1328x360", document.Width, document.Height)
	}

	overview := xyFlowNodeByID(t, document.Nodes, "paper-frame-overview")
	detail := xyFlowNodeByID(t, document.Nodes, "paper-frame-detail")
	group := xyFlowNodeBySize(t, document.Nodes, "xyFlowGroup", 240, 300)
	box := xyFlowNodeBySize(t, document.Nodes, "xyFlowGroup", 280, 300)
	port := xyFlowNodeBySize(t, document.Nodes, "xyFlowGroup", 40, 24)
	remote := xyFlowNodeBySize(t, document.Nodes, "xyFlowGroup", 200, 120)
	item := xyFlowNodeBySize(t, document.Nodes, "xaligoItem", 32, 32)
	if group.ParentID != overview.ID || box.ParentID != overview.ID || port.ParentID != box.ID || item.ParentID != group.ID || remote.ParentID != detail.ID {
		t.Fatalf("unexpected XYFlow nesting: overview=%#v detail=%#v group=%#v box=%#v port=%#v item=%#v remote=%#v", overview, detail, group, box, port, item, remote)
	}

	wantPairs := map[[2]string]bool{
		{group.ID, box.ID}:    false,
		{port.ID, item.ID}:    false,
		{overview.ID, box.ID}: false,
		{box.ID, remote.ID}:   false,
	}
	var metadataEdge, crossFrameEdge entity.XYFlowEdge
	for _, edge := range document.Edges {
		pair := [2]string{edge.Source, edge.Target}
		if _, expected := wantPairs[pair]; !expected {
			t.Fatalf("unexpected edge endpoints %#v: %s", edge, out)
		}
		wantPairs[pair] = true
		if pair == [2]string{group.ID, box.ID} {
			metadataEdge = edge
		}
		if pair == [2]string{box.ID, remote.ID} {
			crossFrameEdge = edge
		}
	}
	for pair, found := range wantPairs {
		if !found {
			t.Fatalf("edge %q -> %q is missing: %s", pair[0], pair[1], out)
		}
	}
	if metadataEdge.Data["kind"] != "traffic" || metadataEdge.Data["bends"] != "250.000,120.000" || metadataEdge.Data["scale"] != float64(1) || metadataEdge.Data["grid"] != float64(8) || metadataEdge.Data["sourceAnchorExplicit"] != true || metadataEdge.Data["targetAnchorExplicit"] != true {
		t.Fatalf("local edge metadata = %#v", metadataEdge.Data)
	}
	if !equalXYFlowFixedPoint(metadataEdge.Data["sourceFixedPoint"], []float64{1, 0.3}) || !equalXYFlowFixedPoint(metadataEdge.Data["targetFixedPoint"], []float64{0, 0.7}) {
		t.Fatalf("local edge fixed points = %#v", metadataEdge.Data)
	}
	if crossFrameEdge.Data["crossFrame"] != true || crossFrameEdge.Data["sourceFrame"] != "overview" || crossFrameEdge.Data["targetFrame"] != "detail" {
		t.Fatalf("cross-frame edge metadata = %#v", crossFrameEdge.Data)
	}
	if crossFrameEdge.Data["bends"] != "330.000,140.000" || crossFrameEdge.Data["scale"] != float64(2) || crossFrameEdge.Data["grid"] != float64(10) || crossFrameEdge.Data["sourceAnchorExplicit"] != true || crossFrameEdge.Data["targetAnchorExplicit"] != true {
		t.Fatalf("cross-frame routing metadata = %#v", crossFrameEdge.Data)
	}
	if crossFrameEdge.SourceHandle != "right" || crossFrameEdge.TargetHandle != "left" || !equalXYFlowFixedPoint(crossFrameEdge.Data["sourceFixedPoint"], []float64{1, 0.9}) || !equalXYFlowFixedPoint(crossFrameEdge.Data["targetFixedPoint"], []float64{0, 0.1}) {
		t.Fatalf("cross-frame bindings were not reassembled: %#v", crossFrameEdge)
	}
}

func TestRenderXYFlowUsesBoxTreeParentsThroughPureLayoutAndVisibleOverflow(t *testing.T) {
	source := []byte(`<frames>
  <frame id="page" width="320" height="300">
    <container width="320" height="300" overflow="visible">
      <panel id="outer" title="Outer" width="120" height="100" overflow="visible">
        <rectangle id="inner" title="Inner" width="120" height="100">
          <port id="inner-port" title="Port" side="right" width="32" height="20" />
        </rectangle>
      </panel>
      <generic-group id="items" title="Items" width="120" height="100">
        <item id="1" name="service" />
      </generic-group>
    </container>
  </frame>
</frames>`)
	options := entity.RenderOptions{
		Format: usecase.FormatXYFlow,
		Theme:  "light",
		Assets: &entity.AssetSource{
			FS:            awsassets.Assets,
			CatalogCSV:    awsassets.CatalogCSV,
			GroupIconsDir: awsassets.GroupIconsDir,
			ItemIconSize:  32,
		},
	}
	out, err := newUsecase().RenderXYFlow(context.Background(), source, options)
	if err != nil {
		t.Fatal(err)
	}
	var document entity.XYFlowDocument
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	page := xyFlowNodeByID(t, document.Nodes, "paper-frame-page")
	outer := xyFlowNodeByID(t, document.Nodes, "page-0-0-rect")
	inner := xyFlowNodeByID(t, document.Nodes, "page-0-0-0-rect")
	port := xyFlowNodeByID(t, document.Nodes, "page-0-0-0-0-rect")
	items := xyFlowNodeByID(t, document.Nodes, "page-0-1-rect")
	item := xyFlowNodeByID(t, document.Nodes, "page-0-1-0-item")
	if outer.ParentID != page.ID {
		t.Fatalf("pure layout container became a semantic parent: page=%#v outer=%#v", page, outer)
	}
	if inner.ParentID != outer.ID {
		t.Fatalf("equal-size overflowing child parent=%q, want %q: %s", inner.ParentID, outer.ID, out)
	}
	if inner.Width != outer.Width || inner.Height != outer.Height {
		t.Fatalf("test precondition failed: outer=%#v inner=%#v", outer, inner)
	}
	if inner.XYFlowPosition.X <= 0 || inner.XYFlowPosition.Y <= 0 {
		t.Fatalf("test precondition failed: child does not visibly overflow parent: %#v", inner)
	}
	if inner.Extent != "" {
		t.Fatalf("overflowing child extent=%q, want no XYFlow parent clamp", inner.Extent)
	}
	if port.ParentID != inner.ID || items.ParentID != page.ID || item.ParentID != items.ID {
		t.Fatalf("unexpected port/item parents: page=%#v inner=%#v port=%#v items=%#v item=%#v", page, inner, port, items, item)
	}

	options.Format = usecase.FormatExcalidraw
	sceneJSON, err := newUsecase().RenderExcalidraw(context.Background(), source, options)
	if err != nil {
		t.Fatal(err)
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		t.Fatal(err)
	}
	wantSemantic := map[string]struct {
		parent string
		kind   string
	}{
		"paper-frame-page":  {kind: "frame"},
		"page-0-0-rect":     {parent: "paper-frame-page", kind: "group"},
		"page-0-0-0-rect":   {parent: "page-0-0-rect", kind: "rectangle"},
		"page-0-0-0-0-rect": {parent: "page-0-0-0-rect", kind: "port"},
		"page-0-1-rect":     {parent: "paper-frame-page", kind: "group"},
		"page-0-1-0-item":   {parent: "page-0-1-rect", kind: "item"},
	}
	for _, element := range scene.Elements {
		want, ok := wantSemantic[element.ID]
		if !ok {
			continue
		}
		if element.CustomData == nil || element.CustomData.SemanticParentElementID != want.parent || element.CustomData.SemanticElementKind != want.kind {
			t.Fatalf("element %q semantic metadata=%#v, want parent=%q kind=%q", element.ID, element.CustomData, want.parent, want.kind)
		}
		delete(wantSemantic, element.ID)
	}
	if len(wantSemantic) != 0 {
		t.Fatalf("semantic scene elements are missing: %#v", wantSemantic)
	}
}

func xyFlowNodeByID(t *testing.T, nodes []entity.XYFlowNode, id string) entity.XYFlowNode {
	t.Helper()
	for _, node := range nodes {
		if node.ID == id {
			return node
		}
	}
	t.Fatalf("XYFlow node %q is missing", id)
	return entity.XYFlowNode{}
}

func xyFlowNodeBySize(t *testing.T, nodes []entity.XYFlowNode, nodeType string, width, height float64) entity.XYFlowNode {
	t.Helper()
	for _, node := range nodes {
		if node.Type == nodeType && node.Width == width && node.Height == height {
			return node
		}
	}
	t.Fatalf("XYFlow %s node %.0fx%.0f is missing", nodeType, width, height)
	return entity.XYFlowNode{}
}

func equalXYFlowFixedPoint(value any, want []float64) bool {
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
