package engine_test

import (
	"strings"
	"testing"

	v1engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestAllUMLDiagramKindsNormalizeAndBuildV1EngineParseUML(t *testing.T) {
	tests := []struct {
		kind          string
		body          string
		relationCount int
	}{
		{"class-diagram", `<class id="one"/><class id="two"/><association src="one" dst="two"/>`, 1},
		{"component-diagram", `<component id="one"/><interface id="two"/><realization src="one" dst="two"/>`, 1},
		{"activity-diagram", `<action id="one"/><action id="two"/><control-flow src="one" dst="two"/>`, 1},
		{"state-machine-diagram", `<state id="one"/><state id="two"/><transition src="one" dst="two"/>`, 1},
		{"sequence-diagram", `<participant id="one"/><lifeline id="two"/><message src="one" dst="two" order="1"/>`, 1},
	}
	for _, test := range tests {
		t.Run(test.kind, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main" width="640" height="360"><uml id="view"><` + test.kind + `>` + test.body + `</` + test.kind + `></uml></frame></frames></xaligo>`
			document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			frame := document.Root.Children[0]
			uml := frame.Children[0]
			if uml.Attr("uml-kind") != test.kind || len(uml.Children) < 2 || uml.Children[0].Tag != "rectangle" {
				t.Fatalf("normalized UML = %#v", uml)
			}
			if len(frame.Children) != 1+test.relationCount || frame.Children[1].Tag != "connection" {
				t.Fatalf("normalized relation = %#v", frame.Children)
			}
			if _, err := v1engine.BuildV1EngineLayoutBuild(document); err != nil {
				t.Fatalf("Build() error = %v", err)
			}
		})
	}
}

func TestUnsupportedUMLDiagramKindsV1EngineParseUML(t *testing.T) {
	unsupportedKinds := []string{
		"object-diagram",
		"deployment-diagram",
		"package-diagram",
		"composite-structure-diagram",
		"profile-diagram",
		"use-case-diagram",
		"communication-diagram",
		"interaction-overview-diagram",
		"timing-diagram",
	}
	for _, kind := range unsupportedKinds {
		t.Run(kind, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="view"><` + kind + `><element id="one"/></` + kind + `></uml></frame></frames></xaligo>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), "unsupported UML diagram kind") {
				t.Fatalf("Parse() error = %v, want unsupported UML diagram kind", err)
			}
		})
	}
}

func TestUMLModelDataReferenceV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data><uml-model id="domain"><class id="user" title="User"/><class id="role" title="Role"/><association src="user" dst="role"/></uml-model></data><frames><frame id="main"><uml id="domain-view"><class-diagram data="domain"/></uml></frame></frames></xaligo>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	if got := document.Root.Children[0].Children[0].Children[0].Attr("uml-element-kind"); got != "class" {
		t.Fatalf("element kind = %q", got)
	}
}

func TestUMLComponentDimensionsAndFanoutNormalizeV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data></data><frames><frame id="main" width="1400" height="700"><uml id="view"><component-diagram grid="4" component-width="280" component-height="180"><component id="one"><interface>Shared API</interface></component><component id="two"><interface>Shared API</interface></component><component id="three"><interface>Shared API</interface></component><component id="hub" height="240"><interface>Shared API</interface><interface>Admin API</interface></component><association src="one" dst="hub"/><association src="two" dst="hub"/><association src="three" dst="hub"/></component-diagram></uml></frame></frames></xaligo>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	uml := document.Root.Children[0].Children[0]
	if uml.Attr("component-width") != "280" || uml.Attr("component-height") != "180" {
		t.Fatalf("component diagram dimensions = width %q height %q", uml.Attr("component-width"), uml.Attr("component-height"))
	}
	hub := uml.Children[3]
	if hub.Attr("height") != "240" || hub.Attr("uml-component-interface-fanout-extra") != "2" {
		t.Fatalf("hub normalized attrs = %#v, want height 240 and fanout extra 2", hub.Attrs)
	}
	if _, err := v1engine.BuildV1EngineLayoutBuild(document); err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	invalid := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="view"><component-diagram component-height="0"><component id="one"/></component-diagram></uml></frame></frames></xaligo>`
	document, err = v1engine.ParseV1EngineParseDocument(strings.NewReader(invalid))
	if err != nil {
		t.Fatalf("Parse(invalid numeric input) error = %v", err)
	}
	if _, err := v1engine.BuildV1EngineLayoutBuild(document); err == nil || !strings.Contains(err.Error(), "component-height") || !strings.Contains(err.Error(), "greater than zero") {
		t.Fatalf("Build(invalid component-height) error = %v", err)
	}
}

func TestUMLRejectsInvalidStructureV1EngineParseUML(t *testing.T) {
	tests := []struct {
		name, body, want string
	}{
		{"missing kind", `<uml id="view"/>`, "exactly one"},
		{"multiple kinds", `<uml id="view"><class-diagram><class id="a"/></class-diagram><sequence-diagram><participant id="b"/></sequence-diagram></uml>`, "exactly one"},
		{"unknown kind", `<uml id="view"><flowchart><element id="a"/></flowchart></uml>`, "unsupported UML diagram kind"},
		{"empty diagram", `<uml id="view"><class-diagram/></uml>`, "must contain UML elements"},
		{"missing element id", `<uml id="view"><class-diagram><class/></class-diagram></uml>`, "requires id"},
		{"bad reference", `<uml id="view"><class-diagram><class id="a"/><association src="a" dst="missing"/></class-diagram></uml>`, "existing src and dst"},
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

func TestUMLDiagramSpecificValidationV1EngineParseUML(t *testing.T) {
	tests := []struct {
		name, diagram, want string
	}{
		{"class rejects action", `<class-diagram><action id="a"/></class-diagram>`, "does not allow UML element"},
		{"realization targets interface", `<class-diagram><class id="a"/><class id="b"/><realization src="a" dst="b"/></class-diagram>`, "does not allow class -> class"},
		{"message requires order", `<sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b"/></sequence-diagram>`, "requires order"},
		{"message order syntax", `<sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b" order="1.a"/></sequence-diagram>`, "dot-separated integers"},
		{"message order leading zero", `<sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b" order="01"/></sequence-diagram>`, "without leading zeroes"},
		{"state rejects attribute", `<state-machine-diagram><state id="a"><attribute>x</attribute></state></state-machine-diagram>`, "does not allow compartment"},
		{"component requires component", `<component-diagram><interface id="api"/></component-diagram>`, "requires at least 1 component"},
		{"component realization targets interface", `<component-diagram><component id="c"/><component id="d"/><realization src="c" dst="d"/></component-diagram>`, "does not allow component -> component"},
		{"component assembly needs port", `<component-diagram><interface id="a"/><interface id="b"/><component id="c"/><assembly src="a" dst="b"/></component-diagram>`, "requires at least one port endpoint"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="view">` + test.diagram + `</uml></frame></frames></xaligo>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestUMLLocalIDsAreScopedByComponentV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="left"><class-diagram><class id="same" name="Shared display name"/></class-diagram></uml><uml id="right"><class-diagram><class id="same" name="Shared display name"/></class-diagram></uml></frame></frames></xaligo>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	left := document.Root.Children[0].Children[0].Children[0].Attr("id")
	right := document.Root.Children[0].Children[1].Children[0].Attr("id")
	if left != "uml-6c656674-73616d65" || right != "uml-7269676874-73616d65" || left == right {
		t.Fatalf("scoped IDs = %q, %q", left, right)
	}
	for _, uml := range document.Root.Children[0].Children {
		element := uml.Children[0]
		if element.Attr("title") != "Shared display name" || element.Attr("name") != "" {
			t.Fatalf("normalized UML display attrs = %#v", element.Attrs)
		}
	}
}

func TestUMLStrictProfileValidationV1EngineParseUML(t *testing.T) {
	tests := []struct {
		name, diagram, want string
	}{
		{"neutral element rejected", `<class-diagram><element id="a"/></class-diagram>`, "does not allow UML element <element>"},
		{"default deny compartment", `<sequence-diagram><participant id="a"><note>not allowed</note></participant></sequence-diagram>`, "does not allow compartment <note>"},
		{"nested compartment content rejected", `<class-diagram><class id="a"><attribute title="name"><operation>hidden</operation></attribute></class></class-diagram>`, "must not contain child elements"},
		{"reserved connector style", `<class-diagram><class id="a"/><class id="b"/><association src="a" dst="b" end-arrowhead="triangle"/></class-diagram>`, "connector semantics are derived"},
		{"invalid id dot", `<class-diagram><class id="a.b"/></class-diagram>`, "must not contain '.' or '/'"},
		{"invalid id slash", `<class-diagram><class id="a/b"/></class-diagram>`, "must not contain '.' or '/'"},
		{"invalid id whitespace", `<class-diagram><class id="a b"/></class-diagram>`, "must not contain whitespace"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="view">` + test.diagram + `</uml></frame></frames></xaligo>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestUMLOwnerValidationV1EngineParseUML(t *testing.T) {
	tests := []struct {
		name, diagram, want string
	}{
		{"owner forbidden elsewhere", `<class-diagram><class id="a" owner="a"/></class-diagram>`, "does not allow owner"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="view">` + test.diagram + `</uml></frame></frames></xaligo>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestUMLMessageAndCommunicationValidationV1EngineParseUML(t *testing.T) {
	tests := []struct {
		name, diagram, want string
	}{
		{"duplicate order", `<sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b" order="1"/><return-message src="b" dst="a" order="1"/></sequence-diagram>`, "duplicate UML message order"},
		{"create self message", `<sequence-diagram><lifeline id="a"/><create-message src="a" dst="a" order="1"/></sequence-diagram>`, "does not allow a self message"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="view">` + test.diagram + `</uml></frame></frames></xaligo>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestUMLTimingAndControlFlowValidationV1EngineParseUML(t *testing.T) {
	tests := []struct {
		name, diagram, want string
	}{
		{"final has outgoing", `<activity-diagram><action id="a"/><final id="f"/><control-flow src="f" dst="a"/></activity-diagram>`, "must not leave a final node"},
		{"decision needs two branches", `<activity-diagram><action id="a"/><decision id="d"/><action id="b"/><control-flow src="a" dst="d"/><control-flow src="d" dst="b"/></activity-diagram>`, "two outgoing flows"},
		{"object flow needs object node", `<activity-diagram><action id="a"/><action id="b"/><object-flow src="a" dst="b"/></activity-diagram>`, "requires at least one object-node"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="view">` + test.diagram + `</uml></frame></frames></xaligo>`
			_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("Parse() error = %v, want %q", err, test.want)
			}
		})
	}
}

func TestUMLPublicReferencesSupportNormalAndCrossFrameConnectionsV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data></data><frames>
<frame id="overview"><rectangle id="caller"/><connection src="caller" dst="detail.domain/order"/></frame>
<frame id="detail"><uml id="domain"><class-diagram><class id="order"/></class-diagram></uml></frame>
</frames></xaligo>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	umlElement := document.Root.Children[1].Children[0].Children[0]
	if got := umlElement.Attr("ref"); got != "domain/order" {
		t.Fatalf("public UML ref = %q", got)
	}
	connection := document.Root.Children[0].Children[1]
	if connection.Attr("_xaligoConnectionCrossFrame") != "true" {
		t.Fatalf("cross-frame connection attrs = %#v", connection.Attrs)
	}

	missingFrameID := `<xaligo version="1"><data></data><frames>
<frame id="overview"><rectangle id="caller"/><connection src="caller" dst="domain/order"/></frame>
<frame id="detail"><uml id="domain"><class-diagram><class id="order"/></class-diagram></uml></frame>
</frames></xaligo>`
	_, err = v1engine.ParseV1EngineParseDocument(strings.NewReader(missingFrameID))
	if err == nil || !strings.Contains(err.Error(), "use frameId.id for a cross-frame reference") {
		t.Fatalf("missing cross-frame frame ID error = %v", err)
	}
}

func TestUMLIDsAreUniqueWithinFrameV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="same"><class-diagram><class id="a"/></class-diagram></uml><uml id="same"><class-diagram><class id="b"/></class-diagram></uml></frame></frames></xaligo>`
	_, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err == nil || !strings.Contains(err.Error(), `duplicate <uml id="same">`) {
		t.Fatalf("Parse() error = %v", err)
	}
}

func TestUMLRegionCompartmentIsNotParsedAsAnArchitectureRegionV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="state"><state-machine-diagram><state id="active"><region>Nested behavior</region></state></state-machine-diagram></uml></frame></frames></xaligo>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	label := document.Root.Children[0].Children[0].Children[0].Attr("title")
	if !strings.Contains(label, "Nested behavior") {
		t.Fatalf("normalized state label = %q", label)
	}
}
