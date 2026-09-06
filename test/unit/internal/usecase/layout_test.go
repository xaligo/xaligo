package usecase_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/usecase"
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

func TestFramesLayoutPlacesPagesWithIDs(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frames gap="40">
  <frame id="overview" width="300" height="180"><blank /></frame>
  <frame id="detail" width="320" height="200"><blank /></frame>
</frames>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if root.Tag != "frames" || root.W != 660 || root.H != 200 {
		t.Fatalf("root = %#v", root)
	}
	if len(root.Children) != 2 {
		t.Fatalf("children = %#v", root.Children)
	}
	first, second := root.Children[0], root.Children[1]
	if first.ID != "overview" || second.ID != "detail" {
		t.Fatalf("frame IDs = %q %q", first.ID, second.ID)
	}
	if first.X != 0 || first.W != 300 || second.X != 340 || second.W != 320 {
		t.Fatalf("page positions = %#v %#v", first, second)
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

func TestStackLayoutReservesFixedHeightBeforeFlexRatios(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="300" height="300" gap="10">
  <rectangle id="fixed" width="100" height="80" />
  <rectangle id="one" width="100" row="1" />
  <rectangle id="two" width="100" row="2" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 3 {
		t.Fatalf("children = %#v", root.Children)
	}
	fixed, one, two := root.Children[0], root.Children[1], root.Children[2]
	if fixed.H != 80 {
		t.Fatalf("fixed height = %.3f, want 80", fixed.H)
	}
	if one.H < 66.6 || one.H > 66.7 || two.H < 133.3 || two.H > 133.4 {
		t.Fatalf("flex heights = %.3f %.3f, want about 66.667 and 133.333", one.H, two.H)
	}
	if one.Y != fixed.Y+fixed.H+10 || two.Y < one.Y+one.H+9.999 || two.Y > one.Y+one.H+10.001 {
		t.Fatalf("stack positions = fixed=%#v one=%#v two=%#v", fixed, one, two)
	}
}

func TestHorizontalLayoutReservesFixedWidthBeforeFlexRatios(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="500" height="120" layout="horizontal" gap="10">
  <rectangle id="fixed" width="100" height="60" />
  <rectangle id="one" col="1" height="60" />
  <rectangle id="three" col="3" height="60" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	fixed, one, three := root.Children[0], root.Children[1], root.Children[2]
	if fixed.W != 100 || one.W != 95 || three.W != 285 {
		t.Fatalf("widths = %.3f %.3f %.3f, want 100 95 285", fixed.W, one.W, three.W)
	}
	if one.X != 110 || three.X != 215 {
		t.Fatalf("positions = %.3f %.3f, want 110 215", one.X, three.X)
	}
}

func TestRowLayoutReservesFixedWidthBeforeGridShare(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="300" height="100">
  <row gap="10">
    <rectangle id="fixed" width="100" />
    <rectangle id="flex" />
  </row>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	fixed, flex := root.Children[0].Children[0], root.Children[0].Children[1]
	if fixed.W != 100 || flex.W != 190 || flex.X != 110 {
		t.Fatalf("row boxes = fixed=%#v flex=%#v, want widths 100/190 and flex x=110", fixed, flex)
	}
}

func TestBuildRejectsParentOverflowUnlessVisible(t *testing.T) {
	const strict = `<frame width="145" height="100"><rectangle id="wide" width="500" height="50" /></frame>`
	doc, err := usecase.Parse(strings.NewReader(strict))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := usecase.Build(doc); err == nil || !strings.Contains(err.Error(), "overflows parent <frame> content box") {
		t.Fatalf("Build() error = %v, want parent overflow diagnostic", err)
	}

	const visible = `<frame width="145" height="100" overflow="visible"><rectangle id="wide" width="500" height="50" /></frame>`
	doc, err = usecase.Parse(strings.NewReader(visible))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 1 || root.Children[0].W != 500 || root.Children[0].IntrinsicW != 500 || root.Overflow != "visible" || !root.Overflowed {
		t.Fatalf("visible overflow tree = %#v", root)
	}
}

func TestBuildRejectsImplicitStaggerOverflow(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="120" height="100">
  <generic-group id="stack" layout="staggered" content-width="20" content-height="20">
    <card /><card />
  </generic-group>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := usecase.Build(doc); err == nil || !strings.Contains(err.Error(), "staggered children require") {
		t.Fatalf("Build() error = %v, want stagger overflow diagnostic", err)
	}
}

func TestVisibleOverflowUsesOriginalExtentForFlexAfterFixedChildren(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="200" height="80" layout="horizontal" gap="10" overflow="visible">
  <rectangle id="fixed" width="240" height="60" />
  <rectangle id="flex" col="1" height="60" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 2 || root.Children[0].W != 240 || root.Children[1].W != 190 || root.Children[1].X != 250 || !root.Overflowed {
		t.Fatalf("visible fixed/flex layout = %#v", root)
	}
}

func TestVisibleOverflowAllowsGapsToConsumeAxisAndPreservesSourceOrder(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="100" height="80" layout="horizontal" gap="120" overflow="visible">
  <rectangle id="flex" col="1" height="60" />
  <rectangle id="fixed" width="40" height="60" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 2 || root.Children[0].W != 100 || root.Children[0].X != 0 || root.Children[1].W != 40 || root.Children[1].X != 220 || !root.Overflowed {
		t.Fatalf("visible gap overflow layout = %#v", root)
	}
}

func TestBuildRejectsOverlappingPortsOnSameSide(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="200" height="120">
  <rectangle id="service" width="120" height="80">
    <port id="first" side="top" width="80" height="20" />
    <port id="second" side="top" width="80" height="20" />
  </rectangle>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := usecase.Build(doc); err == nil || !strings.Contains(err.Error(), `port "second" overlaps port "first" on the top side`) {
		t.Fatalf("Build() error = %v, want port overlap diagnostic", err)
	}
}

func TestBlankAlignUsesDefaultWithoutWarning(t *testing.T) {
	if os.Getenv("XALIGO_TEST_BLANK_ALIGN") == "1" {
		doc, err := usecase.Parse(strings.NewReader(`<frame width="100" height="100"><blank /></frame>`))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := usecase.Build(doc); err != nil {
			t.Fatal(err)
		}
		return
	}

	logPath := filepath.Join(t.TempDir(), "layout.log")
	cmd := exec.Command(os.Args[0], "-test.run=^TestBlankAlignUsesDefaultWithoutWarning$")
	cmd.Env = append(os.Environ(),
		"XALIGO_TEST_BLANK_ALIGN=1",
		"XALIGO_LOG_LEVEL=WARN",
		"XALIGO_LOG_OUTPUT="+logPath,
	)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("helper process failed: %v\n%s", err, output)
	}
	data, err := os.ReadFile(logPath)
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "IULPA-001") {
		t.Fatalf("blank align emitted invalid-align warning: %s", data)
	}
}

func TestStaggeredLayoutMarksDepthAndClampsSmallAreas(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="120" height="100">
  <generic-group id="stack" title="Stack" layout="staggered" content-width="20" content-height="20" align="middle-center" overflow="visible">
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

func TestRectanglePortsStayInsideAndShareSameSide(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="300" height="180">
  <rectangle id="svc" width="120" height="80">
    <port id="a" side="left" width="30" height="16" />
    <port id="b" side="left" width="30" height="16" />
    <port id="c" side="left" width="30" height="16" />
    <port id="d" side="bottom" width="40" height="18" x="-10" y="90" />
  </rectangle>
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
	rect := root.Children[0]
	if len(rect.Children) != 4 {
		t.Fatalf("ports = %#v", rect.Children)
	}
	for _, port := range rect.Children {
		if port.X < rect.X || port.Y < rect.Y || port.X+port.W > rect.X+rect.W || port.Y+port.H > rect.Y+rect.H {
			t.Fatalf("port outside rectangle: rect=%#v port=%#v", rect, port)
		}
	}
	byID := map[string]*entityBox{}
	for _, port := range rect.Children {
		byID[port.Attrs["id"]] = &entityBox{X: port.X, Y: port.Y, W: port.W, H: port.H}
	}
	first, second, third := byID["a"], byID["b"], byID["c"]
	if first == nil || second == nil || third == nil {
		t.Fatalf("left ports not found: %#v", rect.Children)
	}
	if first.X != rect.X || second.X != rect.X || third.X != rect.X {
		t.Fatalf("left ports should sit on inside left edge: %#v %#v %#v", first, second, third)
	}
	if !(first.Y < second.Y && second.Y < third.Y) {
		t.Fatalf("left ports should be ordered along the same side: y %.1f %.1f %.1f", first.Y, second.Y, third.Y)
	}
	explicit := byID["d"]
	if explicit == nil {
		t.Fatalf("explicit port not found: %#v", rect.Children)
	}
	if explicit.X != rect.X || explicit.Y+explicit.H != rect.Y+rect.H {
		t.Fatalf("explicit out-of-bounds port was not clamped to inside bottom-left: rect=%#v port=%#v", rect, explicit)
	}
}

func TestVPCEndpointsAttachToAndSlideAlongVPCBorder(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="400" height="240">
  <vpc id="network" title="VPC" width="300" height="180">
    <vpc-endpoint id="private-api" side="right" anchor="0.25" size="40" />
    <vpc-endpoint id="private-data" side="right" anchor="0.75" size="40" />
  </vpc>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	if len(root.Children) != 1 {
		t.Fatalf("root children = %#v", root.Children)
	}
	vpc := root.Children[0]
	if len(vpc.Children) != 2 {
		t.Fatalf("VPC children = %#v", vpc.Children)
	}
	first, second := vpc.Children[0], vpc.Children[1]
	if first.Tag != "vpc-endpoint" || second.Tag != "vpc-endpoint" {
		t.Fatalf("VPC endpoint tags = %q, %q", first.Tag, second.Tag)
	}
	wantX := vpc.X + vpc.W - first.W/2
	if first.X != wantX || second.X != wantX {
		t.Fatalf("endpoint x = %.1f, %.1f, want border-centered %.1f", first.X, second.X, wantX)
	}
	if first.Y != vpc.Y+0.25*(vpc.H-first.H) || second.Y != vpc.Y+0.75*(vpc.H-second.H) {
		t.Fatalf("endpoint y = %.1f, %.1f; VPC = %#v", first.Y, second.Y, vpc)
	}
	if second.Y <= first.Y || first.W != 40 || first.H != 40 {
		t.Fatalf("endpoint geometry = %#v, %#v", first, second)
	}
}

func TestBuildRejectsOverlappingVPCEndpoints(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="400" height="240">
  <vpc id="network" width="300" height="180">
    <vpc-endpoint id="first" side="right" anchor="0.5" />
    <vpc-endpoint id="second" side="right" anchor="0.5" />
  </vpc>
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := usecase.Build(doc); err == nil || !strings.Contains(err.Error(), `vpc-endpoint "second" overlaps "first"`) {
		t.Fatalf("Build() error = %v, want VPC endpoint overlap diagnostic", err)
	}
}

func TestLeafPaddingPreservesBorderAndSeparatesContentBox(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`
<frame width="240" height="140">
  <badge width="100" height="80" class="pa-1" />
</frame>`))
	if err != nil {
		t.Fatal(err)
	}
	root, err := usecase.Build(doc)
	if err != nil {
		t.Fatal(err)
	}
	leaf := root.Children[0]
	if leaf.W != 100 || leaf.H != 80 {
		t.Fatalf("leaf border = %.1fx%.1f, want 100x80", leaf.W, leaf.H)
	}
	if leaf.ContentX != leaf.X+8 || leaf.ContentY != leaf.Y+8 || leaf.ContentW != 84 || leaf.ContentH != 64 {
		t.Fatalf("leaf content box = (%v,%v,%v,%v), border=%#v", leaf.ContentX, leaf.ContentY, leaf.ContentW, leaf.ContentH, leaf)
	}
}

func TestBuildRejectsItemGridThatCannotFitMinimumCell(t *testing.T) {
	doc, err := usecase.Parse(strings.NewReader(`<frame width="40" height="30"><item id="1" /></frame>`))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := usecase.Build(doc); err == nil || !strings.Contains(err.Error(), "cannot fit 1 item slots") {
		t.Fatalf("Build() error = %v, want item-grid fit error", err)
	}
}

type entityBox struct {
	X float64
	Y float64
	W float64
	H float64
}
