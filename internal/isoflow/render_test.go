package isoflow

import (
	"encoding/json"
	"math"
	"testing"
)

func TestRenderBuildsIsometricDocument(t *testing.T) {
	scene := []byte(`{
  "elements": [
    {"id":"paper-frame","type":"frame","width":400,"height":200},
    {"id":"group-rect","type":"rectangle","x":10,"y":10,"width":380,"height":180,"strokeColor":"#64748b","strokeWidth":1},
    {"id":"a-item","type":"image","x":50,"y":80,"width":32,"height":32,"fileId":"fa"},
    {"id":"a-item-lbl","type":"text","text":"Web"},
    {"id":"b-item","type":"image","x":300,"y":80,"width":32,"height":32,"fileId":"fb"},
    {"id":"b-item-lbl","type":"text","text":"DB"},
    {"id":"edge","type":"arrow","strokeColor":"#2563eb","strokeWidth":2,
      "startBinding":{"elementId":"a-item","fixedPoint":[1,0.5]},
      "endBinding":{"elementId":"b-item-lbl","fixedPoint":[0,0.5]},
      "customData":{"xaligoConnectorKind":"traffic","xaligoConnectorEndArrowhead":"stealth"}}
  ],
  "files":{"fa":{"dataURL":"data:image/svg+xml;base64,QQ=="}},
  "appState":{"viewBackgroundColor":"#ffffff"}
}`)
	out, err := Render(scene)
	if err != nil {
		t.Fatal(err)
	}
	var document Document
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	if document.Version != documentVersion {
		t.Fatalf("version = %q", document.Version)
	}
	if document.Title == "" || len(document.Icons) != 1 || len(document.Colors) == 0 {
		t.Fatalf("document = %#v", document)
	}
	if len(document.Items) != 2 || len(document.Views) != 1 {
		t.Fatalf("items=%d views=%d: %s", len(document.Items), len(document.Views), out)
	}
	view := document.Views[0]
	if len(view.Items) != 2 || len(view.Rectangles) != 0 || len(view.Connectors) != 1 {
		t.Fatalf("view = %#v", view)
	}
	var item ModelItem
	for _, candidate := range document.Items {
		if candidate.ID == "a-item" {
			item = candidate
		}
	}
	if item.Name != "Web" || item.Icon != "fa" {
		t.Fatalf("item = %#v", item)
	}
	viewItem := view.Items[0]
	if viewItem.ID != "a-item" || viewItem.Tile.X != 2 || viewItem.Tile.Y != 2 || viewItem.LabelHeight != defaultLabelHeight {
		t.Fatalf("view item = %#v", viewItem)
	}
	connector := view.Connectors[0]
	if len(connector.Anchors) != 2 || connector.Anchors[0].Ref.Item != "a-item" || connector.Anchors[1].Ref.Item != "b-item" || connector.Style != "SOLID" {
		t.Fatalf("connector = %#v", connector)
	}
}

func TestRenderKeepsOneTileAroundEachItem(t *testing.T) {
	scene := []byte(`{
  "elements": [
    {"id":"a-item","type":"image","x":50,"y":50,"width":32,"height":32,"fileId":"fa"},
    {"id":"a-item-lbl","type":"text","text":"A"},
    {"id":"b-item","type":"image","x":150,"y":50,"width":32,"height":32,"fileId":"fb"},
    {"id":"b-item-lbl","type":"text","text":"B"},
    {"id":"c-item","type":"image","x":50,"y":150,"width":32,"height":32,"fileId":"fc"},
    {"id":"c-item-lbl","type":"text","text":"C"}
  ],
  "files":{}
}`)
	out, err := Render(scene)
	if err != nil {
		t.Fatal(err)
	}
	var document Document
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	items := document.Views[0].Items
	for i := range items {
		for j := i + 1; j < len(items); j++ {
			dx := math.Abs(items[i].Tile.X - items[j].Tile.X)
			dy := math.Abs(items[i].Tile.Y - items[j].Tile.Y)
			if dx <= 1 && dy <= 1 {
				t.Fatalf("items are adjacent: %#v %#v", items[i], items[j])
			}
		}
	}
}

func TestRenderWithGeneratedIconsPreservesServiceNameLabels(t *testing.T) {
	scene := []byte(`{
  "elements": [
    {"id":"lambda-item","type":"image","x":50,"y":50,"width":32,"height":32,"fileId":"lambda"},
    {"id":"lambda-item-lbl","type":"text","text":"AWS Lambda"}
  ],
  "files":{"lambda":{"dataURL":"data:image/svg+xml;base64,QQ=="}}
}`)
	out, err := RenderWithIcons(scene, map[string]string{"lambda": "data:image/svg+xml;base64,PHN2Zz5fX1hBTElHT19MQUJFTF9fPC9zdmc+"})
	if err != nil {
		t.Fatal(err)
	}
	var document Document
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Items) != 1 || document.Items[0].Name != "AWS Lambda" {
		t.Fatalf("items = %#v", document.Items)
	}
}
