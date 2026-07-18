package engine_test

import (
	"strings"
	"testing"

	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestAllUMLDiagramKindsNormalizeAndBuildV1EngineParseUML(t *testing.T) {
	kinds := []string{
		"class-diagram", "object-diagram", "component-diagram", "deployment-diagram",
		"package-diagram", "composite-structure-diagram", "profile-diagram", "use-case-diagram",
		"activity-diagram", "state-machine-diagram", "sequence-diagram", "communication-diagram",
		"interaction-overview-diagram", "timing-diagram",
	}
	for _, kind := range kinds {
		t.Run(kind, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main" width="640" height="360"><uml id="view"><` + kind + `><element id="one" title="One"><compartment>value</compartment></element><element id="two" title="Two"/><relation src="one" dst="two" title="uses"/></` + kind + `></uml></frame></frames></xaligo>`
			document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			frame := document.Root.Children[0]
			uml := frame.Children[0]
			if uml.Attr("uml-kind") != kind || len(uml.Children) != 2 || uml.Children[0].Tag != "rectangle" {
				t.Fatalf("normalized UML = %#v", uml)
			}
			if len(frame.Children) != 2 || frame.Children[1].Tag != "connection" || frame.Children[1].Attr("uml-relation-kind") != "relation" {
				t.Fatalf("normalized relation = %#v", frame.Children)
			}
			if _, err := v1engine.BuildV1EngineLayoutBuild(document); err != nil {
				t.Fatalf("Build() error = %v", err)
			}
		})
	}
}

func TestUMLModelDataReferenceV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data><uml-model id="domain"><class id="user" title="User"/><class id="role" title="Role"/><association src="user" dst="role"/></uml-model></data><frames><frame id="main"><uml><class-diagram data="domain"/></uml></frame></frames></xaligo>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if got := document.Root.Children[0].Children[0].Children[0].Attr("uml-element-kind"); got != "class" {
		t.Fatalf("element kind = %q", got)
	}
}

func TestUMLRejectsInvalidStructureV1EngineParseUML(t *testing.T) {
	tests := []struct {
		name, body, want string
	}{
		{"missing kind", `<uml/>`, "exactly one"},
		{"multiple kinds", `<uml><class-diagram><class id="a"/></class-diagram><sequence-diagram><participant id="b"/></sequence-diagram></uml>`, "exactly one"},
		{"unknown kind", `<uml><flowchart><element id="a"/></flowchart></uml>`, "exactly one"},
		{"empty diagram", `<uml><class-diagram/></uml>`, "must contain UML elements"},
		{"missing element id", `<uml><class-diagram><class/></class-diagram></uml>`, "requires id"},
		{"bad reference", `<uml><class-diagram><class id="a"/><association src="a" dst="missing"/></class-diagram></uml>`, "existing src and dst"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main">` + test.body + `</frame></frames></xaligo>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}
