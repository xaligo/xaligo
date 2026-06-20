package excalidraw

import (
	"encoding/json"
	"strings"
	"testing"

	awsassets "github.com/ryo-arima/xaligo/etc/resources/aws"
	"github.com/ryo-arima/xaligo/internal/layout"
	"github.com/ryo-arima/xaligo/internal/parser"
)

func TestGenericGroupCatalogIcon(t *testing.T) {
	doc, err := parser.Parse(strings.NewReader(`<frame width="400" height="200"><generic-group title="Network" icon-id="200036" /></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := layout.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	var scene file
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	if _, ok := scene.Files["group-cat-200036"]; !ok {
		t.Fatalf("catalog group icon file missing: %v", scene.Files)
	}
	var border, icon, label, header map[string]any
	for _, element := range scene.Elements {
		switch element["id"] {
		case "frame-0-rect":
			border = element
		case "frame-0-icon":
			icon = element
		case "frame-0-label":
			label = element
		case "frame-0-header-bg":
			header = element
		}
	}
	if icon == nil || icon["fileId"] != "group-cat-200036" {
		t.Fatalf("group icon = %#v", icon)
	}
	if icon["width"].(float64) != groupIconSize || icon["height"].(float64) != groupIconSize {
		t.Fatalf("group icon size = %vx%v", icon["width"], icon["height"])
	}
	if label == nil || label["x"].(float64) != icon["x"].(float64)+float64(groupIconSize)+4 {
		t.Fatalf("label/icon positions: label=%#v icon=%#v", label, icon)
	}
	borderY := icon["y"].(float64) + icon["height"].(float64)/2
	if border == nil || border["y"].(float64) != borderY {
		t.Fatalf("border and icon centers do not align: border=%#v icon=%#v", border, icon)
	}
	if label["y"].(float64)+label["height"].(float64)/2 != borderY {
		t.Fatalf("label and icon centers do not align: label=%#v icon=%#v", label, icon)
	}
	if header == nil || header["type"] != "line" || header["backgroundColor"] != "#ffffff" || header["strokeColor"] != "#AAB7B8" {
		t.Fatalf("group header background = %#v", header)
	}
	if points, ok := header["points"].([]any); !ok || len(points) != 6 {
		t.Fatalf("group header is not a closed tag polygon: %#v", header["points"])
	}
	if header["y"].(float64)+header["height"].(float64)/2 != borderY {
		t.Fatalf("header and content centers do not align: header=%#v icon=%#v", header, icon)
	}
	if header["x"].(float64) != icon["x"].(float64) || header["y"].(float64) != icon["y"].(float64) || header["height"].(float64) != icon["height"].(float64) {
		t.Fatalf("header and icon bounds do not align: header=%#v icon=%#v", header, icon)
	}
	if !(header["x"].(float64) <= icon["x"].(float64) && header["y"].(float64) <= icon["y"].(float64) && header["x"].(float64)+header["width"].(float64) >= label["x"].(float64)+label["width"].(float64)) {
		t.Fatalf("header does not cover icon and label: header=%#v icon=%#v label=%#v", header, icon, label)
	}
	wantRight := float64(groupHeaderPadEnd) + min(float64(groupHeaderTipMax), header["height"].(float64)/2)
	if got := header["x"].(float64) + header["width"].(float64) - label["x"].(float64) - label["width"].(float64); got != wantRight {
		t.Fatalf("header right extent = %v, want %v", got, wantRight)
	}
}

func TestGroupHeaderKeepsConservativeTextSpare(t *testing.T) {
	doc, err := parser.Parse(strings.NewReader(`<frame width="800" height="240"><aws-cloud title="AWS Cloud"><region title="ap-northeast-1"><generic-group title="Application Subnet upper lane" /></region></aws-cloud></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := layout.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	var scene file
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	for _, element := range scene.Elements {
		label, ok := element["text"].(string)
		if !ok || label != "Application Subnet upper lane" {
			continue
		}
		width := element["width"].(float64)
		if want := float64(len(label)) * 9.2; width < want {
			t.Fatalf("group label width = %v, want at least %v", width, want)
		}
		return
	}
	t.Fatal("group label not found")
}

func TestItemLabelHeightExpandsForWrappedCatalogLabel(t *testing.T) {
	doc, err := parser.Parse(strings.NewReader(`<frame width="400" height="240"><generic-group title="Network"><item id="200013" /></generic-group></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := layout.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	var scene file
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	for _, element := range scene.Elements {
		if element["id"] != "frame-0-0-item-lbl" {
			continue
		}
		if element["text"] != "l2_switch_general" {
			t.Fatalf("label text = %v", element["text"])
		}
		if got := element["height"].(float64); got <= itemLabelH {
			t.Fatalf("wrapped label height = %v, want greater than %v", got, itemLabelH)
		}
		return
	}
	t.Fatal("item label not found")
}

func TestItemLabelHeightAccountsForPPTXWrapping(t *testing.T) {
	if got := itemLabelHeight("desktop_pc"); got <= itemLabelH {
		t.Fatalf("desktop_pc label height = %v, want greater than %v", got, itemLabelH)
	}
}

func TestGroupWithoutIconKeepsBalancedVerticalPadding(t *testing.T) {
	doc, err := parser.Parse(strings.NewReader(`<frame width="400" height="200"><generic-group title="Compact" /></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := layout.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	var scene file
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var label, header map[string]any
	for _, element := range scene.Elements {
		switch element["id"] {
		case "frame-0-label":
			label = element
		case "frame-0-header-bg":
			header = element
		}
	}
	if label == nil || header == nil {
		t.Fatalf("compact header/label missing: header=%#v label=%#v", header, label)
	}
	top := label["y"].(float64) - header["y"].(float64)
	bottom := header["y"].(float64) + header["height"].(float64) - label["y"].(float64) - label["height"].(float64)
	if top != bottom || top != float64(groupHeaderTextPadY) {
		t.Fatalf("compact header vertical padding top=%v bottom=%v: header=%#v label=%#v", top, bottom, header, label)
	}
}

func TestItemOnlyGroupReservesAnchorGridClearanceBelowHeader(t *testing.T) {
	doc, err := parser.Parse(strings.NewReader(`<frame width="320" height="180"><private-subnet title="Private"><item id="27" /></private-subnet></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := layout.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	var scene file
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	var header, icon map[string]any
	for _, element := range scene.Elements {
		switch element["id"] {
		case "frame-0-header-bg":
			header = element
		case "frame-0-0-item":
			icon = element
		}
	}
	if header == nil || icon == nil {
		t.Fatalf("header/icon missing: header=%#v icon=%#v", header, icon)
	}
	headerBottom := header["y"].(float64) + header["height"].(float64)
	anchorGridTop := icon["y"].(float64) - itemAnchorGridVisualPadPx
	if anchorGridTop < headerBottom {
		t.Fatalf("anchor grid top overlaps header: gridTop=%v headerBottom=%v header=%#v icon=%#v", anchorGridTop, headerBottom, header, icon)
	}
}

func TestAvoidGroupHeaderBorderOverlapNudgesPastSiblingBottom(t *testing.T) {
	elements := []map[string]any{{
		"id": "sibling-rect", "type": "rectangle",
		"x": 100.0, "y": 100.0, "width": 300.0, "height": 80.0,
		"customData": map[string]any{"xaligoGroupBorder": true},
	}}
	got := avoidGroupHeaderBorderOverlap(98, 178, 180, 32, "own-rect", elements)
	want := 180.0 + float64(groupHeaderBorderGap)
	if got != want {
		t.Fatalf("adjusted header y = %v, want %v", got, want)
	}
}

func TestAvoidGroupHeaderBorderOverlapIgnoresOwnBorder(t *testing.T) {
	elements := []map[string]any{{
		"id": "own-rect", "type": "rectangle",
		"x": 100.0, "y": 200.0, "width": 300.0, "height": 80.0,
		"customData": map[string]any{"xaligoGroupBorder": true},
	}}
	got := avoidGroupHeaderBorderOverlap(98, 184, 180, 32, "own-rect", elements)
	if got != 184.0 {
		t.Fatalf("own border changed header y = %v", got)
	}
}

func TestAlignGroupBorderTopToHeaderPreservesHeaderCenter(t *testing.T) {
	elements := []map[string]any{{
		"id": "own-rect", "type": "rectangle",
		"x": 100.0, "y": 200.0, "width": 300.0, "height": 120.0,
		"customData": map[string]any{"xaligoGroupBorder": true},
	}}
	alignGroupBorderTopToHeader("own-rect", 216, 320, elements)
	if elements[0]["y"].(float64) != 216 || elements[0]["height"].(float64) != 104 {
		t.Fatalf("aligned border = %#v", elements[0])
	}
}
