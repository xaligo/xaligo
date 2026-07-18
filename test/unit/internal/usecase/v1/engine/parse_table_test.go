package engine_test

import (
	"strings"
	"testing"

	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestPipeTableStylesInheritAndOverrideV1EngineParseTable(t *testing.T) {
	source := `<frame width="640" height="360"><table color="#112233" background-color="#fefefe" border-color="#445566" font-family="nunito" font-size="16" header-color="#ffffff" header-background-color="#2563eb" header-font-family="cascadia" header-font-size="18">
| Name | Port |
|:-----|-----:|
| API  | 8080 |
<row background-color="#eeeeee"><cell color="#abcdef" font-family="helvetica">DB</cell><cell>5432</cell></row>
</table></frame>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	table := document.Root.Children[0]
	headerCell := table.Children[0].Children[0]
	if headerCell.Attr("color") != "#ffffff" || headerCell.Attr("background-color") != "#2563eb" || headerCell.Attr("font-family") != "cascadia" || headerCell.Attr("font-size") != "18" {
		t.Fatalf("header style = %#v", headerCell.Attrs)
	}
	pipeCell := table.Children[1].Children[0]
	if pipeCell.Attr("color") != "#112233" || pipeCell.Attr("background-color") != "#fefefe" || pipeCell.Attr("border-color") != "#445566" || pipeCell.Attr("font-family") != "nunito" || pipeCell.Attr("font-size") != "16" {
		t.Fatalf("pipe cell style = %#v", pipeCell.Attrs)
	}
	taggedCell := table.Children[2].Children[0]
	if taggedCell.Attr("color") != "#abcdef" || taggedCell.Attr("background-color") != "#eeeeee" || taggedCell.Attr("font-family") != "helvetica" {
		t.Fatalf("tagged cell style = %#v", taggedCell.Attrs)
	}
}

func TestPipeTableRejectsInvalidStyleV1EngineParseTable(t *testing.T) {
	for _, source := range []string{
		`<frame><table color="red">| A |` + "\n" + `|---|</table></frame>`,
		`<frame><table header-background-color="#xyzxyz">| A |` + "\n" + `|---|</table></frame>`,
		`<frame><table font-family="unknown">| A |` + "\n" + `|---|</table></frame>`,
	} {
		if _, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source)); err == nil {
			t.Fatalf("Parse(%q) succeeded, want style error", source)
		}
	}
}
