package usecase_test

import (
	"encoding/json"
	"strings"
	"testing"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestResolveConnectionStyleKinds(t *testing.T) {
	tests := []struct {
		name           string
		attrs          string
		kind           string
		color          string
		width          float64
		startArrowhead string
		endArrowhead   string
		stroke         string
	}{
		{name: "default", kind: "connection", color: "#1e1e1e", width: 1, startArrowhead: "none", endArrowhead: "stealth", stroke: "solid"},
		{name: "route", attrs: `kind="route"`, kind: "route", color: "#64748b", width: 1, startArrowhead: "none", endArrowhead: "none", stroke: "solid"},
		{name: "route without connectors", attrs: `kind="route" start-arrowhead="none" end-arrowhead="none"`, kind: "route", color: "#64748b", width: 1, startArrowhead: "none", endArrowhead: "none", stroke: "solid"},
		{name: "traffic", attrs: `kind="traffic"`, kind: "traffic", color: "#2563eb", width: 1, startArrowhead: "none", endArrowhead: "stealth", stroke: "solid"},
		{name: "overrides", attrs: `kind="traffic" color="#dc2626" stroke-width="3" stroke-style="dotted" start-arrowhead="oval" end-arrowhead="diamond"`, kind: "traffic", color: "#dc2626", width: 3, startArrowhead: "oval", endArrowhead: "diamond", stroke: "dotted"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arrow := buildConnectionArrow(t, tt.attrs)
			custom, _ := arrow["customData"].(map[string]any)
			if custom["xaligoConnectorKind"] != tt.kind || arrow["strokeColor"] != tt.color || arrow["strokeWidth"] != tt.width || custom["xaligoConnectorStartArrowhead"] != tt.startArrowhead || custom["xaligoConnectorEndArrowhead"] != tt.endArrowhead || arrow["strokeStyle"] != tt.stroke {
				t.Fatalf("connection arrow = %#v", arrow)
			}
		})
	}
}

func TestConnectionRoutingMetadata(t *testing.T) {
	arrow := buildConnectionArrow(t, `bends="80,40;100,40" coordinate-scale="2" grid="4"`)
	custom, _ := arrow["customData"].(map[string]any)
	if custom["xaligoConnectorBends"] != "80,40;100,40" || custom["xaligoConnectorScale"] != 2.0 || custom["xaligoConnectorGrid"] != 4.0 {
		t.Fatalf("connection routing metadata = %#v", custom)
	}
}

func TestConnectionRoutingMetadataFromChildTags(t *testing.T) {
	arrow := buildConnectionArrowSource(t, `<frame width="320" height="160"><item id="1" /><item id="2" /><connection src="1" dst="2" coordinate-scale="2" grid="4"><bend x="80" y="40" /><bend x="100" y="60" /></connection></frame>`)
	custom, _ := arrow["customData"].(map[string]any)
	if custom["xaligoConnectorBends"] != "80.000,40.000 100.000,60.000" || custom["xaligoConnectorScale"] != 2.0 || custom["xaligoConnectorGrid"] != 4.0 {
		t.Fatalf("connection routing metadata = %#v", custom)
	}
}

func TestConnectionRoutingMetadataFromConnectionsParent(t *testing.T) {
	arrow := buildConnectionArrowSource(t, `<frame width="320" height="160"><item id="1" /><item id="2" /><connections kind="traffic" color="#2563eb" coordinate-scale="2" grid="4"><connection src="1" dst="2"><bend x="80" y="40" /></connection></connections></frame>`)
	custom, _ := arrow["customData"].(map[string]any)
	if custom["xaligoConnectorKind"] != "traffic" || arrow["strokeColor"] != "#2563eb" || custom["xaligoConnectorScale"] != 2.0 || custom["xaligoConnectorGrid"] != 4.0 {
		t.Fatalf("connection routing metadata = %#v arrow=%#v", custom, arrow)
	}
}

func TestConnectionCanBindToGroupByID(t *testing.T) {
	arrow := buildConnectionArrowSource(t, `<frame width="400" height="220">
  <generic-group id="edge" title="Edge">
    <item id="1" />
  </generic-group>
  <item id="2" name="app" />
  <connection src="edge" dst="app" />
</frame>`)
	start, _ := arrow["startBinding"].(map[string]any)
	end, _ := arrow["endBinding"].(map[string]any)
	startID, _ := start["elementId"].(string)
	endID, _ := end["elementId"].(string)
	if !strings.HasSuffix(startID, "-rect") {
		t.Fatalf("start binding = %#v, want group rectangle", start)
	}
	if !strings.Contains(endID, "-item") {
		t.Fatalf("end binding = %#v, want item image or label", end)
	}
}

func TestConnectionCanBindToRectangleAndPortByID(t *testing.T) {
	arrow := buildConnectionArrowSource(t, `<frame width="400" height="220">
  <rectangle id="svc" title="Service" width="160" height="90">
    <port id="svc-out" side="right" title="out" />
  </rectangle>
  <item id="2" name="app" />
  <connection src="svc-out" dst="app" />
</frame>`)
	start, _ := arrow["startBinding"].(map[string]any)
	startID, _ := start["elementId"].(string)
	if !strings.HasSuffix(startID, "-rect") {
		t.Fatalf("start binding = %#v, want port rectangle", start)
	}
}

func TestConnectionEndpointChildTagsSetAnchorSides(t *testing.T) {
	arrow := buildConnectionArrowSource(t, `<frame width="320" height="160">
  <item id="1" name="web" />
  <item id="2" name="db" />
  <connection>
    <src anchor="top-2">web</src>
    <dst side="left" anchor="5">db</dst>
  </connection>
</frame>`)
	start, _ := arrow["startBinding"].(map[string]any)
	end, _ := arrow["endBinding"].(map[string]any)
	startFP, _ := start["fixedPoint"].([]any)
	endFP, _ := end["fixedPoint"].([]any)
	if len(startFP) != 2 || startFP[0] != 0.25 || startFP[1] != 0.0 {
		t.Fatalf("start fixedPoint = %#v, want top anchor 2 of 5", start["fixedPoint"])
	}
	if len(endFP) != 2 || endFP[0] != 0.0 || endFP[1] != 1.0 {
		t.Fatalf("end fixedPoint = %#v, want left anchor 5 of 5", end["fixedPoint"])
	}
}

func TestExcalidrawItemAnchorsGroupAndCoverConnections(t *testing.T) {
	scene := buildConnectionScene(t, `<frame width="320" height="160"><item id="1" /><item id="2" /><connection src="1" dst="2" /></frame>`)
	indexByID := map[string]int{}
	byID := map[string]map[string]any{}
	for i, element := range scene.Elements {
		id, _ := element["id"].(string)
		if id == "" {
			continue
		}
		indexByID[id] = i
		byID[id] = element
	}
	arrowIndex := -1
	for i, element := range scene.Elements {
		if element["type"] == "arrow" {
			arrowIndex = i
			break
		}
	}
	if arrowIndex < 0 {
		t.Fatalf("arrow not found: %#v", scene.Elements)
	}
	for _, prefix := range []string{"frame-0", "frame-1"} {
		icon := byID[prefix+"-item"]
		label := byID[prefix+"-item-lbl"]
		var firstBg map[string]any
		gridCells := 0
		for _, element := range scene.Elements {
			id, _ := element["id"].(string)
			if strings.HasPrefix(id, prefix+"-anchor-bg-") {
				gridCells++
				if firstBg == nil {
					firstBg = element
				}
			}
		}
		if firstBg == nil || icon == nil || label == nil {
			t.Fatalf("anchor elements missing for %s: bg=%#v icon=%#v label=%#v", prefix, firstBg, icon, label)
		}
		if gridCells != 25 {
			t.Fatalf("anchor grid cells for %s = %d, want 25", prefix, gridCells)
		}
		custom, _ := firstBg["customData"].(map[string]any)
		if custom["xaligoAnchorBackground"] != true {
			t.Fatalf("anchor background customData = %#v", custom)
		}
		if firstBg["backgroundColor"] != "#ffffff" {
			t.Fatalf("anchor background color = %#v", firstBg["backgroundColor"])
		}
		bgGroup := firstGroupID(firstBg)
		if bgGroup == "" || firstGroupID(icon) != bgGroup || firstGroupID(label) != bgGroup {
			t.Fatalf("anchor group mismatch for %s: bg=%#v icon=%#v label=%#v", prefix, firstBg["groupIds"], icon["groupIds"], label["groupIds"])
		}
		firstBgID := firstBg["id"].(string)
		if !(arrowIndex < indexByID[firstBgID] && indexByID[firstBgID] < indexByID[prefix+"-item"] && indexByID[prefix+"-item"] < indexByID[prefix+"-item-lbl"]) {
			t.Fatalf("anchor layer order for %s: arrow=%d bg=%d icon=%d label=%d", prefix, arrowIndex, indexByID[firstBgID], indexByID[prefix+"-item"], indexByID[prefix+"-item-lbl"])
		}
	}
}

func TestExcalidrawConnectionsOffsetOverlappingLanesAndUseSmallHeads(t *testing.T) {
	scene := buildConnectionScene(t, `<frame width="320" height="160">
  <item id="1" />
  <item id="2" />
  <connection src="1" dst="2" arrowhead-size="l" />
  <connection src="1" dst="2" arrowhead-size="l" />
</frame>`)
	arrows := []map[string]any{}
	for _, element := range scene.Elements {
		if element["type"] == "arrow" {
			arrows = append(arrows, element)
		}
	}
	if len(arrows) != 2 {
		t.Fatalf("arrow count = %d, want 2", len(arrows))
	}
	if arrows[0]["endArrowheadSize"] != "s" || arrows[1]["endArrowheadSize"] != "s" {
		t.Fatalf("arrowhead sizes = %#v / %#v", arrows[0]["endArrowheadSize"], arrows[1]["endArrowheadSize"])
	}
	first, _ := json.Marshal(arrows[0]["points"])
	second, _ := json.Marshal(arrows[1]["points"])
	if string(first) == string(second) {
		t.Fatalf("overlapping arrows were not offset: %s", first)
	}
}

func buildConnectionArrow(t *testing.T, attrs string) map[string]any {
	t.Helper()
	return buildConnectionArrowSource(t, `<frame width="320" height="160"><item id="1" /><item id="2" /><connection src="1" dst="2" `+attrs+` /></frame>`)
}

func buildConnectionArrowSource(t *testing.T, source string) map[string]any {
	t.Helper()
	scene := buildConnectionScene(t, source)
	for _, element := range scene.Elements {
		if element["type"] == "arrow" {
			return element
		}
	}
	t.Fatalf("arrow not found: %#v", scene.Elements)
	return nil
}

func buildConnectionScene(t *testing.T, source string) sceneFile {
	t.Helper()
	doc, err := usecase.Parse(strings.NewReader(source))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, testConnectionNodes(doc.Root), nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	return scene
}

func firstGroupID(element map[string]any) string {
	groups, _ := element["groupIds"].([]any)
	if len(groups) == 0 {
		return ""
	}
	group, _ := groups[0].(string)
	return group
}

func testConnectionNodes(root *entity.Node) []*entity.Node {
	if root == nil {
		return nil
	}
	out := []*entity.Node{}
	for _, child := range root.Children {
		switch child.Tag {
		case "connection":
			out = append(out, child)
		case "connections":
			defaults := map[string]string{}
			for _, name := range []string{
				"arrowhead-size", "kind", "color", "stroke-width", "width", "stroke-style",
				"start-arrowhead", "end-arrowhead", "arrowhead", "scale", "coordinate-scale", "grid",
			} {
				if value := strings.TrimSpace(child.Attrs[name]); value != "" {
					defaults[name] = value
				}
			}
			for _, conn := range child.Children {
				if conn.Tag != "connection" {
					continue
				}
				clone := *conn
				clone.Attrs = map[string]string{}
				for key, value := range defaults {
					clone.Attrs[key] = value
				}
				for key, value := range conn.Attrs {
					clone.Attrs[key] = value
				}
				out = append(out, &clone)
			}
		}
	}
	return out
}
