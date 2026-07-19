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
		{"object-diagram", `<object id="one"/><object id="two"/><link src="one" dst="two"/>`, 1},
		{"component-diagram", `<component id="one"/><interface id="two"/><realization src="one" dst="two"/>`, 1},
		{"deployment-diagram", `<artifact id="one"/><node id="two"/><deployment src="one" dst="two"/>`, 1},
		{"package-diagram", `<package id="one"/><package id="two"/><package-import src="one" dst="two"/>`, 1},
		{"composite-structure-diagram", `<structure id="one"/><part id="two" owner="one"/><port id="three" owner="one"/><connector src="two" dst="three"/>`, 1},
		{"profile-diagram", `<profile id="one"/><stereotype id="two"/><metaclass id="three"/><extension src="two" dst="three"/>`, 1},
		{"use-case-diagram", `<actor id="one"/><use-case id="two"/><association src="one" dst="two"/>`, 1},
		{"activity-diagram", `<action id="one"/><action id="two"/><control-flow src="one" dst="two"/>`, 1},
		{"state-machine-diagram", `<state id="one"/><state id="two"/><transition src="one" dst="two"/>`, 1},
		{"sequence-diagram", `<participant id="one"/><lifeline id="two"/><message src="one" dst="two" order="1"/>`, 1},
		{"communication-diagram", `<object id="one"/><object id="two"/><link src="one" dst="two"/><message src="one" dst="two" order="1"/>`, 2},
		{"interaction-overview-diagram", `<interaction id="one"/><interaction id="two"/><control-flow src="one" dst="two"/>`, 1},
		{"timing-diagram", `<lifeline id="one"/><time-state id="two" owner="one" from="0" to="1"/><time-state id="three" owner="one" from="1" to="2"/><transition src="two" dst="three"/>`, 1},
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

func TestUMLActivityPartitionsNormalizeV1EngineParseUML(t *testing.T) {
	source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="atm"><activity-diagram direction="down" lanes="vertical" theme="xaligo"><partition id="customer" title="Customer"><initial id="start"/><action id="enter-pin" title="Enter PIN"/></partition><partition id="atm-lane" title="ATM"><action id="request-pin" title="Request PIN"/><decision id="valid"/></partition><control-flow src="start" dst="enter-pin"/><control-flow src="enter-pin" dst="request-pin"/><control-flow src="request-pin" dst="valid"/><control-flow src="valid" dst="request-pin" guard="invalid" route="loop"/><control-flow src="valid" dst="enter-pin" guard="valid"/></activity-diagram></uml></frame></frames></xaligo>`
	document, err := v1engine.ParseV1EngineParseDocument(strings.NewReader(source))
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}
	uml := document.Root.Children[0].Children[0]
	if uml.Attr("uml-kind") != "activity-diagram" || uml.Attr("layout") != "vertical" {
		t.Fatalf("normalized UML attrs = %#v", uml.Attrs)
	}
	first := uml.Children[0]
	if first.Attr("uml-partition-id") != "customer" || first.Attr("uml-partition-title") != "Customer" {
		t.Fatalf("partition metadata = %#v", first.Attrs)
	}
	frame := document.Root.Children[0]
	if len(frame.Children) != 6 {
		t.Fatalf("normalized frame children = %#v", frame.Children)
	}
	if got := frame.Children[4].Attr("uml-route"); got != "loop" {
		t.Fatalf("loop route = %q, attrs = %#v", got, frame.Children[4].Attrs)
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
		{"class rejects actor", `<class-diagram><actor id="a"/></class-diagram>`, "does not allow UML element"},
		{"object rejects association", `<object-diagram><object id="a"/><object id="b"/><association src="a" dst="b"/></object-diagram>`, "does not allow UML relation"},
		{"include requires use cases", `<use-case-diagram><actor id="a"/><use-case id="b"/><include src="a" dst="b"/></use-case-diagram>`, "does not allow actor -> use-case"},
		{"realization targets interface", `<class-diagram><class id="a"/><class id="b"/><realization src="a" dst="b"/></class-diagram>`, "does not allow class -> class"},
		{"deployment targets node", `<deployment-diagram><node id="node"/><artifact id="a"/><artifact id="b"/><deployment src="a" dst="b"/></deployment-diagram>`, "does not allow artifact -> artifact"},
		{"message requires order", `<sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b"/></sequence-diagram>`, "requires order"},
		{"message order syntax", `<communication-diagram><object id="a"/><object id="b"/><message src="a" dst="b" order="1.a"/></communication-diagram>`, "dot-separated integers"},
		{"message order leading zero", `<sequence-diagram><participant id="a"/><lifeline id="b"/><message src="a" dst="b" order="01"/></sequence-diagram>`, "without leading zeroes"},
		{"state rejects attribute", `<state-machine-diagram><state id="a"><attribute>x</attribute></state></state-machine-diagram>`, "does not allow compartment"},
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
	source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="left"><object-diagram><object id="same" name="Shared display name"/></object-diagram></uml><uml id="right"><object-diagram><object id="same" name="Shared display name"/></object-diagram></uml></frame></frames></xaligo>`
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
		{"neutral relation rejected", `<object-diagram><object id="a"/><object id="b"/><relation src="a" dst="b"/></object-diagram>`, "does not allow UML relation <relation>"},
		{"required component", `<component-diagram><interface id="api"/></component-diagram>`, "requires at least 1 component"},
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
		{"component port owner required", `<component-diagram><component id="c"/><port id="p"/></component-diagram>`, "requires owner"},
		{"component port owner kind", `<component-diagram><component id="c"/><interface id="i"/><port id="p" owner="i"/></component-diagram>`, "does not allow owner kind <interface>"},
		{"composite part owner required", `<composite-structure-diagram><structure id="s"/><part id="p"/></composite-structure-diagram>`, "requires owner"},
		{"time state owner kind", `<timing-diagram><lifeline id="l"/><time-state id="a" owner="a" from="0" to="1"/></timing-diagram>`, "does not allow owner kind <time-state>"},
		{"use case boundary owner", `<use-case-diagram><actor id="a"/><use-case id="u" owner="a"/></use-case-diagram>`, "does not allow owner kind <actor>"},
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
		{"communication needs link", `<communication-diagram><object id="a"/><object id="b"/><message src="a" dst="b" order="1"/></communication-diagram>`, "requires at least one link and one message"},
		{"message needs matching link", `<communication-diagram><object id="a"/><object id="b"/><object id="c"/><link src="a" dst="b"/><message src="b" dst="c" order="1"/></communication-diagram>`, "requires a matching <link>"},
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
		{"time range required", `<timing-diagram><lifeline id="l"/><time-state id="s" owner="l"/></timing-diagram>`, "requires from"},
		{"time range ordered", `<timing-diagram><lifeline id="l"/><time-state id="s" owner="l" from="2" to="1"/></timing-diagram>`, "requires from < to"},
		{"time ranges overlap", `<timing-diagram><lifeline id="l"/><time-state id="a" owner="l" from="0" to="2"/><time-state id="b" owner="l" from="1" to="3"/></timing-diagram>`, "overlaps"},
		{"occurrence at required", `<timing-diagram><lifeline id="l"/><time-state id="s" owner="l" from="0" to="1"/><occurrence src="l" dst="s"/></timing-diagram>`, "requires at"},
		{"occurrence outside interval", `<timing-diagram><lifeline id="l"/><time-state id="s" owner="l" from="0" to="1"/><occurrence src="l" dst="s" at="2"/></timing-diagram>`, "is outside"},
		{"duration range pair", `<timing-diagram><lifeline id="l"/><time-state id="a" owner="l" from="0" to="1"/><time-state id="b" owner="l" from="1" to="2"/><duration src="a" dst="b" from="0"/></timing-diagram>`, "requires from and to together"},
		{"duration range ordered", `<timing-diagram><lifeline id="l"/><time-state id="a" owner="l" from="0" to="1"/><time-state id="b" owner="l" from="1" to="2"/><duration src="a" dst="b" from="2" to="1"/></timing-diagram>`, "requires from < to"},
		{"transition owner mismatch", `<timing-diagram><lifeline id="a"/><lifeline id="b"/><time-state id="one" owner="a" from="0" to="1"/><time-state id="two" owner="b" from="1" to="2"/><transition src="one" dst="two"/></timing-diagram>`, "owned by the same lifeline"},
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
	source := `<xaligo version="1"><data></data><frames><frame id="main"><uml id="same"><class-diagram><class id="a"/></class-diagram></uml><uml id="same"><object-diagram><object id="b"/></object-diagram></uml></frame></frames></xaligo>`
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
