package usecase_test

import (
	"context"
	"encoding/json"
	"math"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
)

func TestFrameMetadataUsesDefaultRowGapInsetDespitePaddingAndContentMargins(t *testing.T) {
	document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="360" height="220" class="pa-2"
         margin-top="72" margin-right="24" margin-bottom="20" margin-left="24">
    <metadata align="right" width="90" key-width="30"><entry key="owner" value="platform" /></metadata>
    <rectangle id="content" title="Content" height="40" />
  </frame>
</frames></xaligo>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(document)
	if err != nil {
		t.Fatal(err)
	}
	frame := root.Children[0]
	metadata := frame.FrameMetadata
	if metadata == nil || len(metadata.Tags) != 2 {
		t.Fatalf("metadata = %#v", metadata)
	}
	const wantInset = 4.0
	if metadata.RowGap != wantInset {
		t.Fatalf("metadata row gap = %v, want default %v", metadata.RowGap, wantInset)
	}
	if got, want := metadata.Tags[0].Y, frame.Y+wantInset; math.Abs(got-want) > 1e-9 {
		t.Fatalf("metadata Y = %v, want one row-gap inside page edge %v", got, want)
	}
	if got, want := metadata.Tags[1].X+metadata.Tags[1].W, frame.X+frame.W-wantInset; math.Abs(got-want) > 1e-9 {
		t.Fatalf("right aligned row edge = %v, want one row-gap inside page edge %v", got, want)
	}
	if metadata.ReservedX != frame.X || metadata.ReservedW != frame.W || metadata.ReservedY != frame.Y || metadata.ReservedY+metadata.ReservedH != frame.ContentY {
		t.Fatalf("reserved strip = (%v,%v,%v,%v), frame/content = %#v", metadata.ReservedX, metadata.ReservedY, metadata.ReservedW, metadata.ReservedH, frame)
	}
}

func TestFrameMetadataReservesTopMarginBand(t *testing.T) {
	document, err := usecase.Parse(strings.NewReader(`<xaligo version="1">
  <data></data>
  <frames>
    <frame id="aws-architecture" title="AWS architecture overview" version="1.0.0"
           width="640" height="360" margin="20">
      <metadata position="top" font-family="helvetica" font-size="16"
                color="#112233" key-color="#334455"
                background-color="#ffffff" key-background-color="#eeeeee"
                border-color="#667788" gap="10" row-gap="6">
        <entry key="owner" value="platform" />
        <entry key="release" value="production" width="220" key-width="72" />
      </metadata>
      <rectangle id="service" title="Service" height="80" />
    </frame>
  </frames>
</xaligo>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(document)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 1 {
		t.Fatalf("frames = %#v", root.Children)
	}
	frame := root.Children[0]
	if frame.FrameMetadata == nil {
		t.Fatal("FrameMetadata = nil")
	}
	metadata := frame.FrameMetadata
	if metadata.Position != "top" || metadata.FontFamily != "helvetica" || metadata.FontSize != 16 {
		t.Fatalf("metadata style = %#v", metadata)
	}
	if metadata.RowGap != 6 {
		t.Fatalf("metadata row gap = %v, want custom value 6", metadata.RowGap)
	}
	if metadata.Color != "#112233" || metadata.KeyColor != "#334455" || metadata.BackgroundColor != "#ffffff" || metadata.KeyBackgroundColor != "#eeeeee" || metadata.BorderColor != "#667788" {
		t.Fatalf("metadata colors = %#v", metadata)
	}
	if len(metadata.Tags) != 5 {
		t.Fatalf("metadata tags = %#v, want id/title/version plus two entries", metadata.Tags)
	}
	wantKeys := []string{"id", "title", "version", "owner", "release"}
	for index, want := range wantKeys {
		if metadata.Tags[index].Key != want {
			t.Fatalf("tag[%d].Key = %q, want %q", index, metadata.Tags[index].Key, want)
		}
	}
	if got, want := metadata.Tags[0].Value, "aws-architecture"; got != want {
		t.Fatalf("id value = %q, want %q", got, want)
	}
	if metadata.Tags[1].W <= metadata.Tags[0].W {
		t.Fatalf("auto widths id=%v title=%v, title should be wider", metadata.Tags[0].W, metadata.Tags[1].W)
	}
	fixed := metadata.Tags[4]
	if fixed.W != 220 || fixed.KeyW != 72 {
		t.Fatalf("fixed tag = %#v, want width=220 key-width=72", fixed)
	}
	wantHeight := math.Ceil(16*1.2) + 4
	for _, tag := range metadata.Tags {
		if tag.H != wantHeight {
			t.Fatalf("tag %q height = %v, want %v", tag.Key, tag.H, wantHeight)
		}
	}
	if len(frame.Children) != 1 || frame.Children[0].Y < frame.ContentY || frame.Children[0].Y < metadata.Tags[len(metadata.Tags)-1].Y+metadata.Tags[len(metadata.Tags)-1].H {
		t.Fatalf("content overlaps metadata: frame=%#v child=%#v metadata=%#v", frame, frame.Children, metadata.Tags)
	}
	if got, want := metadata.Tags[0].Y, frame.Y+metadata.RowGap; math.Abs(got-want) > 1e-9 {
		t.Fatalf("top metadata Y = %v, want row-gap inset %v", got, want)
	}
	if got, want := metadata.Tags[0].X, frame.X+metadata.RowGap; math.Abs(got-want) > 1e-9 {
		t.Fatalf("left aligned metadata edge = %v, want row-gap inset %v", got, want)
	}
	if metadata.ReservedX != frame.X || metadata.ReservedY != frame.Y || metadata.ReservedW != frame.W || math.Abs(metadata.ReservedY+metadata.ReservedH-frame.ContentY) > 1e-9 {
		t.Fatalf("top reserved strip = (%v,%v,%v,%v), frame/content = %#v", metadata.ReservedX, metadata.ReservedY, metadata.ReservedW, metadata.ReservedH, frame)
	}
}

func TestFrameMetadataBottomPositionUsesCustomRowGapAndFontSizedHeight(t *testing.T) {
	document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="release" title="Release" version="2026.07" width="420" height="280"
         margin-top="12" margin-right="12" margin-bottom="80" margin-left="12">
    <metadata position="bottom" font-size="24" width="180" key-width="60" row-gap="9" />
    <rectangle id="content" title="Content" height="80" />
  </frame>
</frames></xaligo>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(document)
	if err != nil {
		t.Fatal(err)
	}
	frame := root.Children[0]
	metadata := frame.FrameMetadata
	if metadata == nil || metadata.Position != "bottom" || len(metadata.Tags) != 3 {
		t.Fatalf("metadata = %#v", metadata)
	}
	if metadata.RowGap != 9 {
		t.Fatalf("metadata row gap = %v, want custom value 9", metadata.RowGap)
	}
	wantHeight := math.Ceil(24*1.2) + 4
	for _, tag := range metadata.Tags {
		if tag.H != wantHeight || tag.W != 180 || tag.KeyW != 60 {
			t.Fatalf("tag = %#v, want height=%v width=180 key-width=60", tag, wantHeight)
		}
		if tag.Y < frame.ContentY+frame.ContentH {
			t.Fatalf("bottom tag overlaps content: tag=%#v content=(%v,%v)", tag, frame.ContentY, frame.ContentH)
		}
	}
	last := metadata.Tags[len(metadata.Tags)-1]
	if got, want := last.Y+last.H, frame.Y+frame.H-metadata.RowGap; math.Abs(got-want) > 1e-9 {
		t.Fatalf("bottom metadata edge = %v, want row-gap inset %v", got, want)
	}
	if got, want := metadata.Tags[0].X, frame.X+metadata.RowGap; math.Abs(got-want) > 1e-9 {
		t.Fatalf("left metadata edge = %v, want row-gap inset %v", got, want)
	}
	if got, want := frame.ContentY+frame.ContentH, metadata.Tags[0].Y-8; math.Abs(got-want) > 1e-9 {
		t.Fatalf("content bottom = %v, want fixed content gap boundary %v", got, want)
	}
	if metadata.ReservedX != frame.X || metadata.ReservedW != frame.W || math.Abs(metadata.ReservedY-(frame.ContentY+frame.ContentH)) > 1e-9 || math.Abs(metadata.ReservedY+metadata.ReservedH-(frame.Y+frame.H)) > 1e-9 {
		t.Fatalf("bottom reserved strip = (%v,%v,%v,%v), frame/content = %#v", metadata.ReservedX, metadata.ReservedY, metadata.ReservedW, metadata.ReservedH, frame)
	}
}

func TestFrameMetadataUsesDefaultRowGapInsetAndAlignsWrappedRows(t *testing.T) {
	tests := []struct {
		align         string
		wantFirstRowX float64
		wantLastRowX  float64
	}{
		{align: "left", wantFirstRowX: 4, wantLastRowX: 4},
		{align: "center", wantFirstRowX: 40, wantLastRowX: 130},
		{align: "right", wantFirstRowX: 76, wantLastRowX: 256},
	}
	for _, test := range tests {
		t.Run(test.align, func(t *testing.T) {
			document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="340" height="240"
         margin-top="64" margin-right="20" margin-bottom="20" margin-left="20">
    <metadata align="` + test.align + `" width="80" key-width="24" gap="10">
      <entry key="a" value="1" />
      <entry key="b" value="2" />
      <entry key="c" value="3" />
    </metadata>
    <rectangle id="content" title="Content" height="40" />
  </frame>
</frames></xaligo>`))
			if err != nil {
				t.Fatal(err)
			}
			root, err := usecase.Build(document)
			if err != nil {
				t.Fatal(err)
			}
			frame := root.Children[0]
			metadata := frame.FrameMetadata
			if metadata == nil || metadata.Align != test.align || len(metadata.Tags) != 4 {
				t.Fatalf("metadata = %#v", metadata)
			}
			if metadata.RowGap != 4 {
				t.Fatalf("metadata row gap = %v, want default value 4", metadata.RowGap)
			}
			if got := frame.ContentY - frame.Y; math.Abs(got-64) > 1e-9 {
				t.Fatalf("content Y offset = %v, want existing top margin 64", got)
			}
			if metadata.Tags[0].Y+metadata.Tags[0].H > frame.ContentY+1e-9 {
				t.Fatalf("metadata is not inside the top reserved strip: tag=%#v contentY=%v", metadata.Tags[0], frame.ContentY)
			}
			if got, want := metadata.Tags[0].Y, frame.Y+metadata.RowGap; math.Abs(got-want) > 1e-9 {
				t.Fatalf("metadata top = %v, want row-gap inset %v", got, want)
			}
			if metadata.ReservedX != frame.X || metadata.ReservedY != frame.Y || metadata.ReservedW != frame.W || math.Abs(metadata.ReservedH-64) > 1e-9 {
				t.Fatalf("reserved strip = (%v,%v,%v,%v), want full-width 64px top margin", metadata.ReservedX, metadata.ReservedY, metadata.ReservedW, metadata.ReservedH)
			}
			if got := metadata.Tags[0].X - frame.X; math.Abs(got-test.wantFirstRowX) > 1e-9 {
				t.Fatalf("first row X = %v, want %v", got, test.wantFirstRowX)
			}
			if got := metadata.Tags[3].X - frame.X; math.Abs(got-test.wantLastRowX) > 1e-9 {
				t.Fatalf("last row X = %v, want %v", got, test.wantLastRowX)
			}
			for _, tag := range metadata.Tags {
				if tag.X < frame.X+metadata.RowGap-1e-9 || tag.X+tag.W > frame.X+frame.W-metadata.RowGap+1e-9 {
					t.Fatalf("tag %#v leaves row-gap-inset usable width [%v,%v]", tag, frame.X+metadata.RowGap, frame.X+frame.W-metadata.RowGap)
				}
			}
			if metadata.Tags[0].Y != metadata.Tags[2].Y || metadata.Tags[3].Y <= metadata.Tags[2].Y {
				t.Fatalf("greedy rows = %#v, want three tags then one tag", metadata.Tags)
			}
			if metadata.FontSize != 12 || metadata.Color != "#64748b" || metadata.KeyColor != "#64748b" || metadata.BackgroundColor != "transparent" || metadata.KeyBackgroundColor != "#f8fafc" || metadata.BorderColor != "#cbd5e1" {
				t.Fatalf("subtle defaults = %#v", metadata)
			}
		})
	}
}

func TestFrameMetadataSupportsExplicitRowBreak(t *testing.T) {
	tests := []struct {
		name      string
		breakAttr string
		wantRows  int
		wantLastY bool
	}{
		{name: "automatic single row", wantRows: 1},
		{name: "explicit break", breakAttr: ` break-before="true"`, wantRows: 2, wantLastY: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="500" height="180" margin="20">
    <metadata width="80" key-width="24" gap="10" row-gap="7">
      <entry key="a" value="1" />
      <entry key="b" value="2"` + test.breakAttr + ` />
      <entry key="c" value="3" />
    </metadata>
    <blank />
  </frame>
</frames></xaligo>`))
			if err != nil {
				t.Fatal(err)
			}
			root, err := usecase.Build(document)
			if err != nil {
				t.Fatal(err)
			}
			frame := root.Children[0]
			metadata := frame.FrameMetadata
			if metadata.RowGap != 7 {
				t.Fatalf("metadata row gap = %v, want custom value 7", metadata.RowGap)
			}
			tags := metadata.Tags
			if got, want := tags[0].X, frame.X+metadata.RowGap; math.Abs(got-want) > 1e-9 {
				t.Fatalf("metadata left edge = %v, want row-gap inset %v", got, want)
			}
			if got, want := tags[0].Y, frame.Y+metadata.RowGap; math.Abs(got-want) > 1e-9 {
				t.Fatalf("metadata top edge = %v, want row-gap inset %v", got, want)
			}
			rows := map[float64]bool{}
			for _, tag := range tags {
				rows[tag.Y] = true
			}
			if len(rows) != test.wantRows {
				t.Fatalf("row count = %d, want %d: %#v", len(rows), test.wantRows, tags)
			}
			if got := tags[len(tags)-1].Y > tags[0].Y; got != test.wantLastY {
				t.Fatalf("last tag wrapped = %v, want %v: %#v", got, test.wantLastY, tags)
			}
			if test.wantLastY {
				if got, want := tags[len(tags)-1].Y-tags[0].Y, tags[0].H+metadata.RowGap; math.Abs(got-want) > 1e-9 {
					t.Fatalf("row advance = %v, want tag height plus row gap %v", got, want)
				}
			}
		})
	}
}

func TestFrameMetadataWrapsAgainstRowGapInsetUsableWidth(t *testing.T) {
	tests := []struct {
		name            string
		frameWidth      string
		wantUsableWidth float64
		sameRow         bool
	}{
		{name: "exact usable width fits", frameWidth: "190", wantUsableWidth: 170, sameRow: true},
		{name: "one pixel below usable width wraps", frameWidth: "189", wantUsableWidth: 169, sameRow: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="` + test.frameWidth + `" height="140" margin-top="64">
    <metadata width="80" key-width="24" gap="10" row-gap="10">
      <entry key="a" value="1" />
    </metadata>
    <blank />
  </frame>
</frames></xaligo>`))
			if err != nil {
				t.Fatal(err)
			}
			root, err := usecase.Build(document)
			if err != nil {
				t.Fatal(err)
			}
			frame := root.Children[0]
			metadata := frame.FrameMetadata
			if metadata == nil || len(metadata.Tags) != 2 || metadata.RowGap != 10 {
				t.Fatalf("metadata = %#v, want two tags and row-gap 10", metadata)
			}
			if got, want := frame.W-2*metadata.RowGap, test.wantUsableWidth; math.Abs(got-want) > 1e-9 {
				t.Fatalf("usable width = %v, want frame width minus two row gaps %v", got, want)
			}
			if got, want := metadata.Tags[0].X, frame.X+metadata.RowGap; math.Abs(got-want) > 1e-9 {
				t.Fatalf("first tag X = %v, want row-gap inset %v", got, want)
			}
			if got, want := metadata.Tags[0].Y, frame.Y+metadata.RowGap; math.Abs(got-want) > 1e-9 {
				t.Fatalf("first tag Y = %v, want row-gap inset %v", got, want)
			}
			if got := metadata.Tags[0].Y == metadata.Tags[1].Y; got != test.sameRow {
				t.Fatalf("tags share row = %v, want %v: usable-width=%v tags=%#v", got, test.sameRow, frame.W-2*metadata.RowGap, metadata.Tags)
			}
			if test.sameRow {
				if got, want := metadata.Tags[1].X+metadata.Tags[1].W, frame.X+frame.W-metadata.RowGap; math.Abs(got-want) > 1e-9 {
					t.Fatalf("exact-fit row edge = %v, want usable-width edge %v", got, want)
				}
			} else {
				if got, want := metadata.Tags[1].X, frame.X+metadata.RowGap; math.Abs(got-want) > 1e-9 {
					t.Fatalf("wrapped tag X = %v, want row-gap inset %v", got, want)
				}
				if got, want := metadata.Tags[1].Y-metadata.Tags[0].Y, metadata.Tags[0].H+metadata.RowGap; math.Abs(got-want) > 1e-9 {
					t.Fatalf("wrapped row advance = %v, want tag height plus row gap %v", got, want)
				}
			}
		})
	}
}

func TestFrameMetadataRejectsRowGapWithoutPositivePageWidth(t *testing.T) {
	document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="20" height="120">
    <metadata row-gap="10"><entry key="owner" value="platform" /></metadata>
    <blank />
  </frame>
</frames></xaligo>`))
	if err != nil {
		t.Fatal(err)
	}
	_, err = usecase.Build(document)
	if err == nil || !strings.Contains(err.Error(), "row-gap") || !strings.Contains(err.Error(), "positive page width") {
		t.Fatalf("Build() error = %v, want row-gap positive page width diagnostic", err)
	}
}

func TestFrameMetadataZeroRowGapKeepsEdgeAlignedFullWidth(t *testing.T) {
	document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="200" height="120" margin-top="40">
    <metadata row-gap="0" width="200" key-width="50" />
    <blank />
  </frame>
</frames></xaligo>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(document)
	if err != nil {
		t.Fatal(err)
	}
	frame := root.Children[0]
	metadata := frame.FrameMetadata
	if metadata == nil || metadata.RowGap != 0 || len(metadata.Tags) != 1 {
		t.Fatalf("metadata = %#v, want one tag with zero row gap", metadata)
	}
	tag := metadata.Tags[0]
	if tag.X != frame.X || tag.Y != frame.Y || tag.W != frame.W {
		t.Fatalf("zero-row-gap tag = %#v, want edge-aligned full frame width %#v", tag, frame)
	}
	if metadata.ReservedX != frame.X || metadata.ReservedY != frame.Y || metadata.ReservedW != frame.W || metadata.ReservedY+metadata.ReservedH != frame.ContentY {
		t.Fatalf("zero-row-gap reserved strip = (%v,%v,%v,%v), frame/content = %#v", metadata.ReservedX, metadata.ReservedY, metadata.ReservedW, metadata.ReservedH, frame)
	}
}

func TestFrameMetadataRejectsInvalidRowGap(t *testing.T) {
	for _, test := range []struct {
		name   string
		value  string
		needle string
	}{
		{name: "negative", value: "-1", needle: "zero or greater"},
		{name: "not finite", value: "NaN", needle: "finite number"},
	} {
		t.Run(test.name, func(t *testing.T) {
			document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="200" height="120">
    <metadata row-gap="` + test.value + `" />
    <blank />
  </frame>
</frames></xaligo>`))
			if err != nil {
				t.Fatal(err)
			}
			_, err = usecase.Build(document)
			if err == nil || !strings.Contains(err.Error(), "row-gap") || !strings.Contains(err.Error(), test.needle) {
				t.Fatalf("Build() error = %v, want row-gap %q diagnostic", err, test.needle)
			}
		})
	}
}

func TestFrameMetadataKeepsManualKeyWidthWhenAutoValueIsClamped(t *testing.T) {
	document, err := usecase.Parse(strings.NewReader(`<xaligo version="1"><frames>
  <frame id="page" width="200" height="160">
    <metadata key-width="150"><entry key="owner" value="a very long owner name that cannot fit" /></metadata>
    <blank />
  </frame>
</frames></xaligo>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(document)
	if err != nil {
		t.Fatal(err)
	}
	tag := root.Children[0].FrameMetadata.Tags[1]
	if tag.W != 192 || tag.KeyW != 150 || tag.X != 4 {
		t.Fatalf("clamped tag = %#v, want row-gap-inset width=192, X=4, and manual key-width=150", tag)
	}
}

func TestFrameMetadataValidation(t *testing.T) {
	valid := `<xaligo version="1"><frames><frame id="page" title="Page" version="1.2.3"><metadata><entry key="owner" value="platform" /></metadata><blank /></frame></frames></xaligo>`
	if _, err := usecase.Parse(strings.NewReader(valid)); err != nil {
		t.Fatalf("Parse(valid) error = %v", err)
	}

	tests := []struct {
		name   string
		body   string
		needle string
	}{
		{name: "position", body: `<metadata position="left" />`, needle: "top or bottom"},
		{name: "font", body: `<metadata font-family="unknown" />`, needle: "supported font family"},
		{name: "color", body: `<metadata color="red" />`, needle: "#RRGGBB or transparent"},
		{name: "missing key", body: `<metadata><entry value="platform" /></metadata>`, needle: "non-empty key"},
		{name: "missing value", body: `<metadata><entry key="owner" /></metadata>`, needle: "non-empty value"},
		{name: "unknown child", body: `<metadata><property key="owner" value="platform" /></metadata>`, needle: "may only contain <entry>"},
		{name: "duplicate metadata", body: `<metadata /><metadata />`, needle: "only one <metadata>"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<xaligo version="1"><frames><frame id="page" title="Page">` + test.body + `<blank /></frame></frames></xaligo>`
			_, err := usecase.Parse(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.needle) {
				t.Fatalf("Parse() error = %v, want %q", err, test.needle)
			}
		})
	}
}

func TestFrameMetadataFlowsThroughSceneAndPagePlan(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="32">
  <frame id="overview" title="Architecture Overview" version="1.0.0" width="520" height="300" margin="16">
    <metadata font-family="cascadia" font-size="18" color="#112233"
              background-color="#fefefe" key-background-color="#ddeeff"
              border-color="#445566"><entry key="owner" value="platform" /></metadata>
    <rectangle id="overview-content" title="Overview content" height="100" />
  </frame>
  <frame id="detail" title="Service Detail" version="2.0.0" width="520" height="300" margin="16">
    <metadata position="bottom"><entry key="owner" value="service-team" /></metadata>
    <rectangle id="detail-content" title="Detail content" height="100" />
  </frame>
</frames></xaligo>`)
	renderer := newUsecase()
	sceneJSON, err := renderer.BuildScene(context.Background(), source, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene struct {
		Elements []map[string]any `json:"elements"`
	}
	if err := json.Unmarshal(sceneJSON, &scene); err != nil {
		t.Fatal(err)
	}
	metadataRects, metadataTexts := 0, 0
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoPageFrame"] == true && element["strokeColor"] != "transparent" {
			t.Fatalf("page frame has visible stroke: %#v", element)
		}
		if custom["xaligoFrameMetadata"] != true {
			continue
		}
		if custom["xaligoFrameID"] != "overview" && custom["xaligoFrameID"] != "detail" {
			t.Fatalf("metadata frame ownership = %#v", custom)
		}
		switch element["type"] {
		case "rectangle":
			metadataRects++
		case "text":
			metadataTexts++
			if custom["xaligoFrameID"] == "overview" {
				if element["fontFamily"] != float64(3) || element["fontSize"] != float64(18) {
					t.Fatalf("overview metadata typography = %#v", element)
				}
			}
		}
		if custom["xaligoSemanticElementKind"] != "frame-metadata" {
			t.Fatalf("metadata semantic kind = %#v", custom)
		}
	}
	if metadataRects != 16 || metadataTexts != 16 {
		t.Fatalf("metadata scene counts rect=%d text=%d, want 16 each", metadataRects, metadataTexts)
	}

	planJSON, err := renderer.BuildPPTXPlan(context.Background(), source, entity.RenderOptions{Format: usecase.FormatPPTX, Theme: "light", PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	var document entity.DocumentPlan
	if err := json.Unmarshal(planJSON, &document); err != nil {
		t.Fatal(err)
	}
	if len(document.Pages) != 2 {
		t.Fatalf("pages = %#v", document.Pages)
	}
	assertFrameMetadataPage(t, document.Pages[0], []string{"overview", "Architecture Overview", "1.0.0", "platform"}, []string{"detail", "Service Detail", "2.0.0", "service-team"})
	assertFrameMetadataPage(t, document.Pages[1], []string{"detail", "Service Detail", "2.0.0", "service-team"}, []string{"overview", "Architecture Overview", "1.0.0", "platform"})
}

func TestFrameMetadataTransparentTextIsOmittedFromPhysicalPlan(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames>
  <frame id="page" title="Hidden metadata" width="360" height="200" margin="32">
    <metadata color="transparent" key-color="transparent"><entry key="owner" value="platform" /></metadata>
    <rectangle id="content" title="Visible content" height="40" />
  </frame>
</frames></xaligo>`)
	planJSON, err := newUsecase().BuildPPTXPlan(context.Background(), source, entity.RenderOptions{Format: usecase.FormatPPTX, Theme: "light", PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		t.Fatal(err)
	}
	foundVisibleContent := false
	for _, op := range plan.Ops {
		if op.Kind == "text" && op.Text == "Visible content" {
			foundVisibleContent = true
		}
		if op.Kind == "text" && strings.Contains(op.ID, "metadata") {
			t.Fatalf("transparent metadata text reached physical plan: %#v", op)
		}
	}
	if !foundVisibleContent {
		t.Fatalf("visible content text is missing: %#v", plan.Ops)
	}
}

func TestFrameMetadataStaysAbovePageLinksAndKeepsTheirLabelsClear(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="overview" title="Overview" version="1.0" width="420" height="240" margin="16">
    <metadata width="170" key-width="58">
      <entry key="owner" value="platform" />
      <entry key="stage" value="production" />
      <entry key="region" value="east" />
      <entry key="review" value="approved" />
      <entry key="team" value="service" />
    </metadata>
    <rectangle id="web" title="Web" width="100" height="30" />
    <connection src="web" dst="detail.db" src-side="top" dst-side="top" />
  </frame>
  <frame id="detail" title="Detail" version="2.0" width="420" height="240" margin="16">
    <metadata width="170" key-width="58">
      <entry key="owner" value="data" />
      <entry key="stage" value="production" />
      <entry key="region" value="west" />
      <entry key="review" value="approved" />
      <entry key="team" value="database" />
    </metadata>
    <rectangle id="db" title="DB" width="100" height="30" />
  </frame>
</frames></xaligo>`)
	out, err := newUsecase().BuildScene(context.Background(), source, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}

	metadataByFrame := map[string][][4]float64{}
	reservedByFrame := map[string][4]float64{}
	lastPageLinkIndex := -1
	firstMetadataIndex := len(scene.Elements)
	for index, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoFrameMetadataReserved"] == true {
			frameID, _ := custom["xaligoFrameID"].(string)
			reservedByFrame[frameID] = [4]float64{
				sceneNumber(t, element["x"]), sceneNumber(t, element["y"]),
				sceneNumber(t, element["width"]), sceneNumber(t, element["height"]),
			}
		}
		if custom["xaligoCrossFrame"] == true || element["text"] == "to <detail>" || element["text"] == "from <overview>" {
			lastPageLinkIndex = index
		}
		if custom["xaligoFrameMetadata"] != true {
			continue
		}
		if index < firstMetadataIndex {
			firstMetadataIndex = index
		}
		if element["type"] != "rectangle" {
			continue
		}
		frameID, _ := custom["xaligoFrameID"].(string)
		metadataByFrame[frameID] = append(metadataByFrame[frameID], [4]float64{
			sceneNumber(t, element["x"]), sceneNumber(t, element["y"]),
			sceneNumber(t, element["width"]), sceneNumber(t, element["height"]),
		})
	}
	if lastPageLinkIndex < 0 || firstMetadataIndex <= lastPageLinkIndex {
		t.Fatalf("metadata layer=%d, last page-link layer=%d", firstMetadataIndex, lastPageLinkIndex)
	}

	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		frameID, _ := custom["xaligoFrameID"].(string)
		if element["text"] == "to <detail>" {
			frameID = "overview"
		} else if element["text"] == "from <overview>" {
			frameID = "detail"
		} else {
			continue
		}
		label := [4]float64{
			sceneNumber(t, element["x"]), sceneNumber(t, element["y"]),
			sceneNumber(t, element["width"]), sceneNumber(t, element["height"]),
		}
		for _, metadata := range metadataByFrame[frameID] {
			if rectanglesOverlapFrameMetadataTest(label, metadata) {
				t.Fatalf("page-link label %#v overlaps frame %q metadata %#v", label, frameID, metadata)
			}
		}
		if reserved := reservedByFrame[frameID]; reserved[2] <= 0 || reserved[3] <= 0 || rectanglesOverlapFrameMetadataTest(label, reserved) {
			t.Fatalf("page-link label %#v is not clear of frame %q reserved strip %#v", label, frameID, reserved)
		}
	}
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] != true {
			continue
		}
		frameID, _ := custom["xaligoFrameID"].(string)
		points := sceneArrowPoints(t, element)
		for index := 1; index < len(points); index++ {
			for _, metadata := range metadataByFrame[frameID] {
				if segmentCrossesRectangleInteriorFrameMetadataTest(points[index-1], points[index], metadata) {
					t.Fatalf("page-link segment %#v -> %#v crosses frame %q metadata %#v", points[index-1], points[index], frameID, metadata)
				}
			}
			if reserved := reservedByFrame[frameID]; segmentCrossesRectangleInteriorFrameMetadataTest(points[index-1], points[index], reserved) {
				t.Fatalf("page-link segment %#v -> %#v crosses frame %q reserved strip %#v", points[index-1], points[index], frameID, reserved)
			}
		}
	}
}

func TestFrameMetadataReservedStripExcludesItemsGroupHeadersAndLocalRoutes(t *testing.T) {
	source := `<xaligo version="1"><frames><frame id="page" title="Page" width="420" height="280" margin-top="72" margin-right="16" margin-bottom="16" margin-left="16">
  <metadata width="120" key-width="40"><entry key="owner" value="platform" /></metadata>
  <generic-group id="network" title="Network" height="96"><item id="27" name="web" /></generic-group>
  <item id="110" name="database" />
  <connection src="network" dst="database" bends="210,12" uml-relation-label="calls" />
</frame></frames></xaligo>`
	out, err := newUsecase().BuildScene(context.Background(), []byte(source), entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	reserved := [4]float64{}
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoFrameMetadataReserved"] == true {
			reserved = [4]float64{sceneNumber(t, element["x"]), sceneNumber(t, element["y"]), sceneNumber(t, element["width"]), sceneNumber(t, element["height"])}
			break
		}
	}
	if reserved[2] != 420 || reserved[3] != 72 {
		t.Fatalf("reserved strip = %#v, want full-width 72px top strip", reserved)
	}
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoFrameMetadata"] == true || custom["xaligoFrameMetadataReserved"] == true || custom["xaligoPageFrame"] == true {
			continue
		}
		typeName, _ := element["type"].(string)
		if typeName == "arrow" {
			points := sceneArrowPoints(t, element)
			for index := 1; index < len(points); index++ {
				if segmentCrossesRectangleInteriorFrameMetadataTest(points[index-1], points[index], reserved) {
					t.Fatalf("local connector segment %#v -> %#v crosses reserved strip %#v: %#v", points[index-1], points[index], reserved, element)
				}
			}
			continue
		}
		if typeName != "rectangle" && typeName != "line" && typeName != "image" && typeName != "text" {
			continue
		}
		rect := [4]float64{sceneNumber(t, element["x"]), sceneNumber(t, element["y"]), sceneNumber(t, element["width"]), sceneNumber(t, element["height"])}
		if rectanglesOverlapFrameMetadataTest(rect, reserved) {
			t.Fatalf("ordinary %s %q rect %#v overlaps reserved strip %#v", typeName, element["id"], rect, reserved)
		}
	}
}

func TestFrameMetadataKeepsNestedGroupHeadersSeparated(t *testing.T) {
	source := `<xaligo version="1"><frames><frame id="page" title="Page" version="1.0" width="640" height="360">
  <aws-cloud id="cloud" title="AWS Cloud">
    <generic-group id="edge" title="Global Edge" icon-id="104915" />
  </aws-cloud>
</frame></frames></xaligo>`
	out, err := newUsecase().BuildScene(context.Background(), []byte(source), entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	headers := make([][4]float64, 0, 2)
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoGroupHeader"] != true {
			continue
		}
		headers = append(headers, [4]float64{
			sceneNumber(t, element["x"]),
			sceneNumber(t, element["y"]),
			sceneNumber(t, element["width"]),
			sceneNumber(t, element["height"]),
		})
	}
	if len(headers) != 2 {
		t.Fatalf("group headers = %#v, want outer and nested headers", headers)
	}
	if headers[1][1] < headers[0][1] {
		headers[0], headers[1] = headers[1], headers[0]
	}
	wantMinimumY := headers[0][1] + headers[0][3] + 4
	if headers[1][1] < wantMinimumY {
		t.Fatalf("nested group header y = %v, want at least %v after outer header %#v", headers[1][1], wantMinimumY, headers[0])
	}
}

func TestFrameMetadataReservedStripExcludesPhysicalPlanLinesAndLabels(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames><frame id="page" title="Page" width="420" height="260" margin-top="72" margin-right="16" margin-bottom="16" margin-left="16">
  <metadata width="120" key-width="40"><entry key="owner" value="platform" /></metadata>
  <rectangle id="source" title="Source" width="120" height="44" />
  <rectangle id="destination" title="Destination" width="120" height="44" />
  <connection src="source" dst="destination" bends="210,12" uml-relation-label="calls" />
</frame></frames></xaligo>`)
	planJSON, err := newUsecase().BuildPPTXPlan(context.Background(), source, entity.RenderOptions{Format: usecase.FormatPPTX, Theme: "light", PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	var plan entity.Plan
	if err := json.Unmarshal(planJSON, &plan); err != nil {
		t.Fatal(err)
	}
	reservedBottom := 72.0 / 96.0
	foundConnector := false
	foundUMLLabel := false
	for _, op := range plan.Ops {
		if strings.Contains(op.ID, "metadata") {
			continue
		}
		if op.Kind == "line" && len(op.Points) > 0 {
			foundConnector = true
			for _, point := range op.Points {
				if y := op.Y + point.Y; y < reservedBottom-1e-9 {
					t.Fatalf("physical connector %q enters reserved strip: y=%v, boundary=%v, op=%#v", op.ID, y, reservedBottom, op)
				}
			}
		}
		if op.Kind == "text" {
			if op.Text == "calls" {
				foundUMLLabel = true
			}
			if op.Y < reservedBottom-1e-9 {
				t.Fatalf("physical text %q enters reserved strip: %#v", op.Text, op)
			}
		}
	}
	if !foundConnector || !foundUMLLabel {
		t.Fatalf("missing connector/UML label in plan: %#v", plan.Ops)
	}
}

func TestFrameMetadataReservedStripOverridesUnsafeLocalEndpointSide(t *testing.T) {
	source := `<xaligo version="1"><frames><frame id="page" title="Page" width="360" height="220" margin-top="64" margin-right="16" margin-bottom="16" margin-left="16" align="top-left">
  <metadata width="120" key-width="40" />
  <item id="27" name="source" />
  <item id="110" name="destination" />
  <connection src="source" dst="destination" src-side="top" dst-side="top" />
</frame></frames></xaligo>`
	out, err := newUsecase().BuildScene(context.Background(), []byte(source), entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}
	reserved := [4]float64{0, 0, 360, 64}
	found := false
	for _, element := range scene.Elements {
		if element["type"] != "arrow" {
			continue
		}
		found = true
		points := sceneArrowPoints(t, element)
		for index := 1; index < len(points); index++ {
			if segmentCrossesRectangleInteriorFrameMetadataTest(points[index-1], points[index], reserved) {
				t.Fatalf("unsafe explicit endpoint side entered reserved strip: %#v -> %#v, arrow=%#v", points[index-1], points[index], element)
			}
		}
	}
	if !found {
		t.Fatal("local connector arrow not found")
	}
}

func TestFrameMetadataPageLinksAvoidFullUsableWidthTagsAndFrameEndpoints(t *testing.T) {
	source := []byte(`<xaligo version="1"><frames gap="48">
  <frame id="source" title="Source" width="300" height="180">
    <metadata width="292" key-width="72" />
    <rectangle id="local" title="Local" width="80" height="30" />
    <connection src="source" dst="target.item" src-side="top" dst-side="top" />
  </frame>
  <frame id="target" title="Target" width="300" height="180">
    <metadata width="292" key-width="72" />
    <rectangle id="item" title="Item" width="80" height="30" />
  </frame>
</frames></xaligo>`)
	out, err := newUsecase().BuildScene(context.Background(), source, entity.RenderOptions{Theme: "light"})
	if err != nil {
		t.Fatal(err)
	}
	var scene sceneFile
	if err := json.Unmarshal(out, &scene); err != nil {
		t.Fatal(err)
	}

	metadataByFrame := map[string][][4]float64{}
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoFrameMetadata"] != true || element["type"] != "rectangle" {
			continue
		}
		frameID, _ := custom["xaligoFrameID"].(string)
		metadataByFrame[frameID] = append(metadataByFrame[frameID], [4]float64{
			sceneNumber(t, element["x"]), sceneNumber(t, element["y"]),
			sceneNumber(t, element["width"]), sceneNumber(t, element["height"]),
		})
	}
	crossFrameArrows := 0
	for _, element := range scene.Elements {
		custom, _ := element["customData"].(map[string]any)
		if custom["xaligoCrossFrame"] != true {
			continue
		}
		crossFrameArrows++
		frameID, _ := custom["xaligoFrameID"].(string)
		points := sceneArrowPoints(t, element)
		for index := 1; index < len(points); index++ {
			for _, metadata := range metadataByFrame[frameID] {
				if segmentCrossesRectangleInteriorFrameMetadataTest(points[index-1], points[index], metadata) {
					t.Fatalf("frame-endpoint page-link segment %#v -> %#v crosses frame %q full-width metadata %#v", points[index-1], points[index], frameID, metadata)
				}
			}
		}
	}
	if crossFrameArrows != 2 {
		t.Fatalf("cross-frame arrows = %d, want two page-local stubs", crossFrameArrows)
	}
}

func rectanglesOverlapFrameMetadataTest(left, right [4]float64) bool {
	return left[0] < right[0]+right[2] && left[0]+left[2] > right[0] && left[1] < right[1]+right[3] && left[1]+left[3] > right[1]
}

func segmentCrossesRectangleInteriorFrameMetadataTest(start, end [2]float64, rectangle [4]float64) bool {
	left, top := rectangle[0], rectangle[1]
	right, bottom := left+rectangle[2], top+rectangle[3]
	if math.Abs(start[0]-end[0]) <= 1e-9 {
		minimumY, maximumY := math.Min(start[1], end[1]), math.Max(start[1], end[1])
		return start[0] > left && start[0] < right && maximumY > top && minimumY < bottom
	}
	if math.Abs(start[1]-end[1]) <= 1e-9 {
		minimumX, maximumX := math.Min(start[0], end[0]), math.Max(start[0], end[0])
		return start[1] > top && start[1] < bottom && maximumX > left && minimumX < right
	}
	return true
}

func assertFrameMetadataPage(t *testing.T, page entity.DocumentPage, wanted, unwanted []string) {
	t.Helper()
	texts := map[string]bool{}
	metadataFrontOps := 0
	for _, op := range page.Ops {
		if op.Kind == "text" {
			texts[op.Text] = true
		}
		if op.FrontLayer && strings.Contains(op.ID, "metadata") {
			metadataFrontOps++
		}
	}
	for _, value := range wanted {
		if !texts[value] {
			t.Fatalf("page %q is missing metadata value %q: %#v", page.ID, value, page.Ops)
		}
	}
	for _, value := range unwanted {
		if texts[value] {
			t.Fatalf("page %q contains metadata from another frame %q", page.ID, value)
		}
	}
	if metadataFrontOps == 0 {
		t.Fatalf("page %q has no front-layer metadata ops: %#v", page.ID, page.Ops)
	}
}
