package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

var umlDiagramKindsV1EngineParseUml = map[string]bool{
	"class-diagram": true, "object-diagram": true, "component-diagram": true,
	"deployment-diagram": true, "package-diagram": true, "composite-structure-diagram": true,
	"profile-diagram": true, "use-case-diagram": true, "activity-diagram": true,
	"state-machine-diagram": true, "sequence-diagram": true, "communication-diagram": true,
	"interaction-overview-diagram": true, "timing-diagram": true,
}

var umlElementTagsV1EngineParseUml = map[string]bool{
	"element": true, "class": true, "interface": true, "enumeration": true, "object": true,
	"component": true, "node": true, "artifact": true, "package": true, "part": true, "port": true,
	"profile": true, "stereotype": true, "metaclass": true, "actor": true, "use-case": true,
	"activity": true, "action": true, "decision": true, "merge": true, "fork": true, "join": true,
	"state": true, "initial": true, "final": true, "history": true, "choice": true,
	"participant": true, "lifeline": true, "interaction": true, "time-state": true,
}

var umlRelationTagsV1EngineParseUml = map[string]bool{
	"relation": true, "association": true, "aggregation": true, "composition": true,
	"generalization": true, "realization": true, "dependency": true, "include": true,
	"extend": true, "control-flow": true, "object-flow": true, "transition": true,
	"message": true, "link": true, "occurrence": true, "duration": true,
}

func normalizeUMLDiagramsV1EngineParseUml(root, data *entity.Node) error {
	models := map[string]*entity.Node{}
	if data != nil {
		for _, child := range data.Children {
			if child.Tag != "uml-model" {
				continue
			}
			id := strings.TrimSpace(child.Attr("id"))
			if id == "" {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<uml-model> requires id")}
			}
			if models[id] != nil {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate <uml-model id=%q>", id)}
			}
			models[id] = child
		}
	}
	var walk func(*entity.Node, *entity.Node) error
	walk = func(node, frame *entity.Node) error {
		if node.Tag == "frame" {
			frame = node
		}
		if node.Tag == "uml" {
			if frame == nil {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<uml> must be inside a <frame>")}
			}
			return normalizeUMLComponentV1EngineParseUml(node, frame, models)
		}
		for _, child := range node.Children {
			if err := walk(child, frame); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(root, nil)
}

func normalizeUMLComponentV1EngineParseUml(uml, frame *entity.Node, models map[string]*entity.Node) error {
	if len(uml.Children) != 1 || !umlDiagramKindsV1EngineParseUml[uml.Children[0].Tag] {
		return &entity.ParseError{Position: uml.Position, Err: fmt.Errorf("<uml> requires exactly one supported UML diagram-kind child")}
	}
	diagram := uml.Children[0]
	if reference := strings.TrimSpace(diagram.Attr("data")); reference != "" {
		model := models[reference]
		if model == nil {
			return &entity.ParseError{Position: diagram.Position, Err: fmt.Errorf("UML data reference %q is not defined", reference)}
		}
		if len(diagram.Children) > 0 {
			return &entity.ParseError{Position: diagram.Position, Err: fmt.Errorf("UML diagram cannot combine data=%q with inline elements", reference)}
		}
		diagram.Children = cloneUMLChildrenV1EngineParseUml(model.Children)
	}
	if len(diagram.Children) == 0 {
		return &entity.ParseError{Position: diagram.Position, Err: fmt.Errorf("<%s> must contain UML elements", diagram.Tag)}
	}
	if uml.Attr("title") == "" {
		uml.Attrs["title"] = strings.TrimSuffix(diagram.Tag, "-diagram")
	}
	uml.Attrs["uml-kind"] = diagram.Tag
	if uml.Attr("layout") == "" {
		if diagram.Attr("direction") == "right" || diagram.Tag == "sequence-diagram" || diagram.Tag == "timing-diagram" {
			uml.Attrs["layout"] = "horizontal"
		} else {
			uml.Attrs["layout"] = "vertical"
		}
	}
	ids := map[string]bool{}
	var elements []*entity.Node
	var relations []*entity.Node
	for _, child := range diagram.Children {
		switch {
		case umlElementTagsV1EngineParseUml[child.Tag]:
			id := strings.TrimSpace(child.Attr("id"))
			if id == "" {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("UML <%s> requires id", child.Tag)}
			}
			if ids[id] {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate UML element id %q", id)}
			}
			ids[id] = true
			elements = append(elements, normalizeUMLElementV1EngineParseUml(child))
		case umlRelationTagsV1EngineParseUml[child.Tag]:
			relations = append(relations, child)
		default:
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<%s> does not support UML child <%s>", diagram.Tag, child.Tag)}
		}
	}
	for _, relation := range relations {
		src, dst := strings.TrimSpace(relation.Attr("src")), strings.TrimSpace(relation.Attr("dst"))
		if !ids[src] || !ids[dst] {
			return &entity.ParseError{Position: relation.Position, Err: fmt.Errorf("UML <%s> requires existing src and dst element IDs", relation.Tag)}
		}
		connection := normalizeUMLRelationV1EngineParseUml(relation)
		if err := validateConnectionNodeV1EngineParseConnection(connection); err != nil {
			return &entity.ParseError{Position: relation.Position, Err: fmt.Errorf("invalid UML <%s>: %w", relation.Tag, err)}
		}
		frame.Children = append(frame.Children, connection)
	}
	uml.Children = elements
	return nil
}

func normalizeUMLElementV1EngineParseUml(source *entity.Node) *entity.Node {
	attrs := cloneAttrsV1EngineParseTable(source.Attrs)
	attrs["uml-element-kind"] = source.Tag
	if attrs["title"] == "" {
		attrs["title"] = attrs["name"]
	}
	if attrs["title"] == "" {
		attrs["title"] = source.Text
	}
	if attrs["title"] == "" {
		attrs["title"] = source.Attr("id")
	}
	var compartments []string
	for _, child := range source.Children {
		label := strings.TrimSpace(child.Attr("title"))
		if label == "" {
			label = strings.TrimSpace(child.Attr("name"))
		}
		if label == "" {
			label = strings.TrimSpace(child.Text)
		}
		if label != "" {
			compartments = append(compartments, label)
		}
	}
	if len(compartments) > 0 {
		attrs["title"] += "\n────────\n" + strings.Join(compartments, "\n")
	}
	return &entity.Node{Tag: "rectangle", Attrs: attrs, Position: source.Position}
}

func normalizeUMLRelationV1EngineParseUml(source *entity.Node) *entity.Node {
	attrs := cloneAttrsV1EngineParseTable(source.Attrs)
	attrs["uml-relation-kind"] = source.Tag
	if attrs["label"] == "" {
		attrs["label"] = source.Attr("title")
	}
	switch source.Tag {
	case "dependency", "realization":
		attrs["stroke-style"] = "dashed"
		attrs["end-arrowhead"] = "triangle"
	case "generalization", "include", "extend", "control-flow", "object-flow", "transition", "message":
		attrs["end-arrowhead"] = "triangle"
	case "aggregation", "composition":
		attrs["start-arrowhead"] = "diamond"
	case "association", "relation", "link", "occurrence", "duration":
		attrs["end-arrowhead"] = "none"
	}
	return &entity.Node{Tag: "connection", Attrs: attrs, Position: source.Position}
}

func cloneUMLChildrenV1EngineParseUml(source []*entity.Node) []*entity.Node {
	result := make([]*entity.Node, 0, len(source))
	for _, child := range source {
		copy := &entity.Node{Tag: child.Tag, Attrs: cloneAttrsV1EngineParseTable(child.Attrs), Text: child.Text, TextRuns: append([]entity.TextRun(nil), child.TextRuns...), Position: child.Position}
		copy.Children = cloneUMLChildrenV1EngineParseUml(child.Children)
		result = append(result, copy)
	}
	return result
}
