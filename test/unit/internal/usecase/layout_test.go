package usecase_test

import (
	"strings"
	"testing"

	"github.com/ryo-arima/xaligo/internal/usecase"
)

func TestContentAreaAlignsChildren(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="400" height="300">
  <container content-width="200" content-height="100" align="bottom-right">
    <blank />
  </container>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 1 || len(root.Children[0].Children) != 1 {
		t.Fatalf("unexpected tree shape: %#v", root.Children)
	}
	child := root.Children[0].Children[0]
	if child.X != 200 || child.Y != 200 || child.W != 200 || child.H != 100 {
		t.Fatalf("child = x %.1f y %.1f w %.1f h %.1f, want 200 200 200 100", child.X, child.Y, child.W, child.H)
	}
}

func TestFrameMarginKeepsPaperSizeAndInsetsContent(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="400" height="300" margin="20">
  <blank />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if root.X != 0 || root.Y != 0 || root.W != 400 || root.H != 300 {
		t.Fatalf("root = x %.1f y %.1f w %.1f h %.1f, want 0 0 400 300", root.X, root.Y, root.W, root.H)
	}
	if len(root.Children) != 1 {
		t.Fatalf("children = %d, want 1", len(root.Children))
	}
	child := root.Children[0]
	if child.X != 20 || child.Y != 20 || child.W != 360 || child.H != 260 {
		t.Fatalf("child = x %.1f y %.1f w %.1f h %.1f, want 20 20 360 260", child.X, child.Y, child.W, child.H)
	}
}

func TestBlankTagsAreItemLike(t *testing.T) {
	for _, tag := range []string{"spacer", "blank"} {
		if !usecase.IsBlank(tag) {
			t.Fatalf("%s should be blank", tag)
		}
		if !usecase.IsItemLike(tag) {
			t.Fatalf("%s should be item-like", tag)
		}
	}
}

func TestHorizontalLayoutUsesColumnWeightsAndMargins(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="420" height="120" layout="horizontal" gap="20">
  <card title="A" col="1" class="mx-1" />
  <card title="B" col="2" margin-left="10" margin-right="6" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 2 {
		t.Fatalf("children = %#v", root.Children)
	}
	first, second := root.Children[0], root.Children[1]
	if first.X != 8 || first.W < 120 || first.W > 124 {
		t.Fatalf("first = x %.1f w %.1f", first.X, first.W)
	}
	if second.X <= first.X+first.W || second.W <= first.W*1.8 {
		t.Fatalf("second = x %.1f w %.1f", second.X, second.W)
	}
}

func TestStaggeredLayoutMarksDepthAndClampsSmallAreas(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="120" height="100">
  <generic-group title="Stack" layout="staggered" content-width="20" content-height="20" align="middle-center">
    <card title="Front" />
    <card title="Back" />
  </generic-group>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	group := root.Children[0]
	if len(group.Children) != 2 {
		t.Fatalf("group children = %#v", group.Children)
	}
	if !group.Children[0].InStagger || group.Children[0].StaggerDepth != 1 || !group.Children[0].IsStaggerBg {
		t.Fatalf("back child = %#v", group.Children[0])
	}
	if !group.Children[1].InStagger || group.Children[1].StaggerDepth != 0 || group.Children[1].IsStaggerBg {
		t.Fatalf("front child = %#v", group.Children[1])
	}
	for _, child := range group.Children {
		if child.W < usecase.MinBoxWidth || child.H < usecase.MinBoxHeight {
			t.Fatalf("child was not clamped: %#v", child)
		}
	}
}

func TestSpacingClassesAndDirectMarginsAreApplied(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="300" height="180" class="pa-1 px-2 py-3 pt-4 pr-5 pb-6 pl-7 ma-1 mt-2 mr-3 mb-4 ml-5" margin="10" margin-top="11" margin-right="12" margin-bottom="13" margin-left="14">
  <container class="pa-x" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 1 {
		t.Fatalf("children = %#v", root.Children)
	}
	child := root.Children[0]
	if child.X <= 0 || child.Y <= 0 || child.W >= root.W || child.H >= root.H {
		t.Fatalf("spacing was not applied: root=%#v child=%#v", root, child)
	}
}
