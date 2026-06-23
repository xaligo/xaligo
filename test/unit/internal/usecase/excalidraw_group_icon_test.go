package usecase_test

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	awsassets "github.com/ryo-arima/xaligo/etc/resources/aws"
	"github.com/ryo-arima/xaligo/internal/share"
	"github.com/ryo-arima/xaligo/internal/usecase"
)

func TestGenericGroupCatalogIcon(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame width="400" height="200"><generic-group title="Network" icon-id="200036" /></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
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
	if icon["width"].(float64) != 32 || icon["height"].(float64) != 32 {
		t.Fatalf("group icon size = %vx%v", icon["width"], icon["height"])
	}
	if label == nil || label["x"].(float64) != icon["x"].(float64)+32+4 {
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
	wantRight := 18.0 + min(14.0, header["height"].(float64)/2)
	if got := header["x"].(float64) + header["width"].(float64) - label["x"].(float64) - label["width"].(float64); got != wantRight {
		t.Fatalf("header right extent = %v, want %v", got, wantRight)
	}
}

func TestGroupHeaderKeepsConservativeTextSpare(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame width="800" height="240"><aws-cloud title="AWS Cloud"><region title="ap-northeast-1"><generic-group title="Application Subnet upper lane" /></region></aws-cloud></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
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
	doc, err := usecase.Parse(strings.NewReader(`<frame width="400" height="240"><generic-group title="Network"><item id="200013" /></generic-group></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
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
		if got := element["height"].(float64); got <= 14.0 {
			t.Fatalf("wrapped label height = %v, want greater than 14", got)
		}
		return
	}
	t.Fatal("item label not found")
}

func TestTablerItemIconCurrentColorIsResolved(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame width="200" height="160"><generic-group title="Network"><item id="104915" /></generic-group></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	file, ok := scene.Files["item-cat-104915"].(map[string]any)
	if !ok {
		t.Fatalf("item file missing: %#v", scene.Files)
	}
	dataURL, ok := file["dataURL"].(string)
	if !ok || !strings.HasPrefix(dataURL, share.SVGDataURLPrefix) {
		t.Fatalf("item dataURL = %#v", file["dataURL"])
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(dataURL, share.SVGDataURLPrefix))
	if err != nil {
		t.Fatal(err)
	}
	svg := string(raw)
	if strings.Contains(strings.ToLower(svg), "currentcolor") {
		t.Fatalf("item SVG still contains currentColor: %s", svg)
	}
	if !strings.Contains(svg, "#7758C1") {
		t.Fatalf("item SVG does not contain fallback icon color: %s", svg)
	}
}

func TestGroupWithoutIconKeepsBalancedVerticalPadding(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame width="400" height="200"><generic-group title="Compact" /></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
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
	if top != bottom || top != 1.0 {
		t.Fatalf("compact header vertical padding top=%v bottom=%v: header=%#v label=%#v", top, bottom, header, label)
	}
}

func TestItemOnlyGroupReservesAnchorGridClearanceBelowHeader(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame width="320" height="180"><private-subnet title="Private"><item id="27" /></private-subnet></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	out, err := usecase.BuildJSONWithFS(root, awsassets.Assets, awsassets.CatalogCSV, awsassets.GroupIconsDir, 32, nil, nil, newSceneDependencies())
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
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
	anchorGridTop := icon["y"].(float64) - 6.0
	if anchorGridTop < headerBottom {
		t.Fatalf("anchor grid top overlaps header: gridTop=%v headerBottom=%v header=%#v icon=%#v", anchorGridTop, headerBottom, header, icon)
	}
}
