//go:build !js

package repository

import "testing"

func TestPDFSVGComposerHonorsGeneratedRectangularTextClip(t *testing.T) {
	composer, err := newPDFSVGComposer()
	if err != nil {
		t.Fatal(err)
	}
	drawing, err := composer.parse([]byte(`
<svg xmlns="http://www.w3.org/2000/svg" width="100" height="50" viewBox="0 0 100 50">
  <defs><clipPath id="label-clip" clipPathUnits="userSpaceOnUse"><rect x="10" y="5" width="20" height="20"/></clipPath></defs>
  <text x="10" y="20" font-family="not-installed" font-size="12" clip-path="url(#label-clip)"><tspan x="10" y="20">This label is wider than its box</tspan></text>
</svg>`))
	if err != nil {
		t.Fatal(err)
	}
	if drawing.Empty() {
		t.Fatal("clipped text was dropped")
	}
	bounds := drawing.Bounds()
	if bounds.X0 < 10-0.01 || bounds.X1 > 30+0.01 {
		t.Fatalf("clipped text bounds = %+v, want within x=[10,30]", bounds)
	}
}

func TestPDFSVGComposerPlacesOverlaysWithoutViewBox(t *testing.T) {
	composer, err := newPDFSVGComposer()
	if err != nil {
		t.Fatal(err)
	}
	drawing, err := composer.parse([]byte(`
<svg xmlns="http://www.w3.org/2000/svg" width="100px" height="50px">
  <text x="10" y="20" font-family="not-installed" font-size="10"><tspan x="10" y="20">No viewBox</tspan></text>
</svg>`))
	if err != nil {
		t.Fatal(err)
	}
	width, height := drawing.Size()
	if width != 100 || height != 50 {
		t.Fatalf("drawing size = %gx%g, want 100x50", width, height)
	}
	bounds := drawing.Bounds()
	if bounds.X0 < 9.99 || bounds.X0 > 10.01 || bounds.X1 > 100 {
		t.Fatalf("text bounds = %+v; overlay was not placed in root pixel coordinates", bounds)
	}
}
