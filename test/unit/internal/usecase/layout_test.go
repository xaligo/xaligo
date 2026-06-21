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
