package repository_test

import (
	"encoding/json"
	"testing"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/repository"
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
