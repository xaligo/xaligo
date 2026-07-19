package usecase_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestRenderIsoflowPreservesNonItemConnectionEndpoints(t *testing.T) {
	source := []byte(`<frame width="640" height="480" gap="16">
  <generic-group id="group" title="AWS Group" height="120"><item id="27" name="ec2" /></generic-group>
  <rectangle id="service" title="Service" height="180">
    <port id="service-out" side="right" title="out" width="64" height="24" />
  </rectangle>
  <connection src="ec2" dst="group" />
  <connection src="group" dst="service" kind="route" />
  <connection src="service" dst="service-out" kind="traffic">
    <bend x="260" y="220" />
    <bend x="420" y="220" />
  </connection>
</frame>`)

	document := renderIsoflowDocument(t, source)
	view := document.Views[0]
	viewItemIDs := isoflowViewItemIDs(view.Items)
	if len(document.Items) == 0 || !strings.HasSuffix(document.Items[0].ID, "-item") || document.Items[0].IsoflowIcon == "" {
		t.Fatalf("existing item output/order changed: %#v", document.Items)
	}
	if !viewItemIDs[document.Items[0].ID] {
		t.Fatalf("existing item %q lost its view position: %#v", document.Items[0].ID, view.Items)
	}
	for _, name := range []string{"AWS Group", "Service", "out"} {
		modelID := isoflowModelItemIDByName(t, document.Items, name)
		if !viewItemIDs[modelID] {
			t.Fatalf("model item %q (%q) is not positioned in the view: %#v", name, modelID, view.Items)
		}
	}
	if len(view.Connectors) != 3 {
		t.Fatalf("connectors = %#v, want 3", view.Connectors)
	}
	for _, connector := range view.Connectors {
		if len(connector.Anchors) < 2 {
			t.Fatalf("connector %q anchors = %#v", connector.ID, connector.Anchors)
		}
		for _, endpoint := range []entity.IsoflowConnectorAnchor{connector.Anchors[0], connector.Anchors[len(connector.Anchors)-1]} {
			if !viewItemIDs[endpoint.Ref.Item] {
				t.Fatalf("connector %q references missing view item %q", connector.ID, endpoint.Ref.Item)
			}
		}
	}
	if len(view.Rectangles) == 0 {
		t.Fatal("AWS group rectangle was removed while adding its endpoint node")
	}
	foundNativeBend := false
	for _, connector := range view.Connectors {
		for _, anchor := range connector.Anchors[1 : len(connector.Anchors)-1] {
			if anchor.Ref.Tile != nil {
				foundNativeBend = true
			}
		}
	}
	if !foundNativeBend {
		t.Fatalf("explicit bend was not represented by an Isoflow tile anchor: %#v", view.Connectors)
	}
}

func TestRenderIsoflowDeduplicatesCrossFrameConnectionAndEmitsUMLShapeEndpoint(t *testing.T) {
	source := []byte(`<frames gap="80">
  <frame id="left" width="400" height="300">
    <rectangle id="left-node" title="Left Node" width="160" height="100" />
	<connection src="left" dst="left-node" />
	<connection src="left-node" dst="right.activity/remote" kind="traffic" />
  </frame>
	<frame id="right" width="400" height="300">
		<uml id="activity"><activity-diagram>
			<action id="remote" title="Remote Action" />
		</activity-diagram></uml>
  </frame>
</frames>`)

	document := renderIsoflowDocument(t, source)
	view := document.Views[0]
	leftFrameID := isoflowModelItemIDByName(t, document.Items, "left")
	leftNodeID := isoflowModelItemIDByName(t, document.Items, "Left Node")
	rightNodeID := isoflowModelItemIDByName(t, document.Items, "Remote Action")
	viewItemIDs := isoflowViewItemIDs(view.Items)
	for _, id := range []string{leftFrameID, leftNodeID, rightNodeID} {
		if !viewItemIDs[id] {
			t.Fatalf("endpoint node %q is not positioned in view: %#v", id, view.Items)
		}
	}
	if len(view.Connectors) != 2 {
		t.Fatalf("connectors = %#v, want local connector plus one deduplicated cross-frame connector", view.Connectors)
	}
	crossFrameCount := 0
	for _, connector := range view.Connectors {
		if len(connector.Anchors) != 2 {
			continue
		}
		sourceID := connector.Anchors[0].Ref.Item
		targetID := connector.Anchors[1].Ref.Item
		if sourceID == leftNodeID && targetID == rightNodeID {
			crossFrameCount++
		}
	}
	if crossFrameCount != 1 {
		t.Fatalf("cross-frame logical connector count = %d, connectors=%#v", crossFrameCount, view.Connectors)
	}
}

func TestRenderIsoflowPreservesUMLShapeEndpoints(t *testing.T) {
	tests := []struct {
		name           string
		source         string
		endpointNames  []string
		connectorPairs [][2]string
	}{
		{
			name: "activity final ellipse",
			source: `<xaligo version="1"><data></data><frames><frame id="main" width="640" height="360">
	<uml id="activity"><activity-diagram direction="right">
		<initial id="start" title="Start"/><action id="active" title="Active"/><final id="done" title="Done"/>
		<control-flow src="start" dst="active"/><control-flow src="active" dst="done"/>
	</activity-diagram></uml>
</frame></frames></xaligo>`,
			endpointNames:  []string{"Start", "Active", "Done"},
			connectorPairs: [][2]string{{"Start", "Active"}, {"Active", "Done"}},
		},
		{
			name: "activity decision diamond",
			source: `<xaligo version="1"><data></data><frames><frame id="main" width="720" height="420">
  <uml id="activity"><activity-diagram direction="right">
    <action id="validate" title="Validate"/><decision id="decision" title="Valid?"/>
    <action id="accept" title="Accept"/><action id="reject" title="Reject"/>
    <control-flow src="validate" dst="decision"/>
    <control-flow src="decision" dst="accept" guard="yes"/>
    <control-flow src="decision" dst="reject" guard="no"/>
  </activity-diagram></uml>
</frame></frames></xaligo>`,
			endpointNames: []string{"Validate", "Valid?", "Accept", "Reject"},
			connectorPairs: [][2]string{
				{"Validate", "Valid?"}, {"Valid?", "Accept"}, {"Valid?", "Reject"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			document := renderIsoflowDocument(t, []byte(test.source))
			view := document.Views[0]
			viewItemIDs := isoflowViewItemIDs(view.Items)
			modelIDs := make(map[string]string, len(test.endpointNames))
			for _, name := range test.endpointNames {
				modelID := isoflowModelItemIDByName(t, document.Items, name)
				if !viewItemIDs[modelID] {
					t.Fatalf("UML endpoint %q is not positioned in the Isoflow view: %#v", name, view.Items)
				}
				modelIDs[name] = modelID
			}
			if len(view.Connectors) != len(test.connectorPairs) {
				t.Fatalf("connectors = %#v, want %d", view.Connectors, len(test.connectorPairs))
			}
			for _, pair := range test.connectorPairs {
				if !isoflowHasConnectorBetween(view.Connectors, modelIDs[pair[0]], modelIDs[pair[1]]) {
					t.Fatalf("UML connector %q -> %q is missing: %#v", pair[0], pair[1], view.Connectors)
				}
			}
		})
	}
}

func renderIsoflowDocument(t *testing.T, source []byte) entity.IsoflowDocument {
	t.Helper()
	out, err := newUsecase().RenderIsoflow(context.Background(), source, entity.RenderOptions{Format: usecase.FormatIsoflow, Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var document entity.IsoflowDocument
	if err := json.Unmarshal(out, &document); err != nil {
		t.Fatalf("decode Isoflow output: %v\n%s", err, out)
	}
	if len(document.Views) != 1 {
		t.Fatalf("views = %#v", document.Views)
	}
	if !document.FitToScreen {
		t.Fatal("Isoflow output did not request fitToScreen")
	}
	iconIDs := make(map[string]bool, len(document.Icons))
	for _, icon := range document.Icons {
		if icon.ID == "" || icon.URL == "" {
			t.Fatalf("invalid Isoflow icon: %#v", icon)
		}
		iconIDs[icon.ID] = true
	}
	for _, item := range document.Items {
		if item.IsoflowIcon == "" || !iconIDs[item.IsoflowIcon] {
			t.Fatalf("Isoflow item %q references missing icon %q", item.ID, item.IsoflowIcon)
		}
	}
	return document
}

func isoflowModelItemIDByName(t *testing.T, items []entity.IsoflowModelItem, name string) string {
	t.Helper()
	for _, item := range items {
		if strings.TrimSpace(item.Name) == name {
			return item.ID
		}
	}
	t.Fatalf("model item named %q is missing: %#v", name, items)
	return ""
}

func isoflowViewItemIDs(items []entity.IsoflowViewItem) map[string]bool {
	ids := make(map[string]bool, len(items))
	for _, item := range items {
		ids[item.ID] = true
	}
	return ids
}

func isoflowHasConnectorBetween(connectors []entity.IsoflowConnector, source, target string) bool {
	for _, connector := range connectors {
		if len(connector.Anchors) < 2 {
			continue
		}
		if connector.Anchors[0].Ref.Item == source && connector.Anchors[len(connector.Anchors)-1].Ref.Item == target {
			return true
		}
	}
	return false
}
