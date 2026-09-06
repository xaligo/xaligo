package share_test

import (
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/share"
)

func TestWrapPresentationTextRespectsMixedWidthAndExplicitLines(t *testing.T) {
	for _, value := range []string{"API / アプリケーション基盤", "ＡＢＣＤ and ABCD", "ｶﾀｶﾅ / カタカナ", "first line\nsecond line", "LongUnbrokenResourceIdentifier"} {
		wrapped := share.WrapPresentationText(value, 80, 12)
		for _, line := range strings.Split(wrapped, "\n") {
			if width := share.PresentationTextWidth(line, 12, false); width > 80 {
				t.Errorf("%q width=%v", line, width)
			}
		}
		compact := func(s string) string { return strings.NewReplacer("\n", "", " ", "").Replace(s) }
		if compact(wrapped) != compact(value) {
			t.Errorf("content changed: %q -> %q", value, wrapped)
		}
	}
	if got := share.WrapPresentationText("", 160, 12); got != "" {
		t.Fatal(got)
	}
}
