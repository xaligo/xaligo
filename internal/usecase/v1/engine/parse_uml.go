package engine

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/xaligo/xaligo/internal/entity"
)

type umlDiagramSpecV1EngineParseUml struct {
	elements             map[string]bool
	relations            map[string]bool
	requiredElementKinds []umlRequiredElementsV1EngineParseUml
}

type umlRequiredElementsV1EngineParseUml struct {
	description string
	kinds       map[string]bool
	minimum     int
}

type umlSourceElementV1EngineParseUml struct {
	node      *entity.Node
	partition string
}

var umlDiagramSpecsV1EngineParseUml = map[string]umlDiagramSpecV1EngineParseUml{
	"class-diagram": umlSpecV1EngineParseUml("class,interface,enumeration", "association,aggregation,composition,generalization,realization,dependency",
		umlRequiredElementsV1EngineParseUml{"classifier", umlTagSetV1EngineParseUml("class,interface,enumeration"), 1}),
	"component-diagram": umlSpecV1EngineParseUml("component,interface,port,artifact", "dependency,realization,association,assembly,delegation",
		umlRequiredElementsV1EngineParseUml{"component", umlTagSetV1EngineParseUml("component"), 1}),
	"activity-diagram": umlSpecV1EngineParseUml("initial,final,activity,action,decision,merge,fork,join,object-node", "control-flow,object-flow",
		umlRequiredElementsV1EngineParseUml{"activity or action", umlTagSetV1EngineParseUml("activity,action"), 1}),
	"state-machine-diagram": umlSpecV1EngineParseUml("initial,final,state,history,choice,fork,join", "transition",
		umlRequiredElementsV1EngineParseUml{"state", umlTagSetV1EngineParseUml("state"), 1}),
	"sequence-diagram": umlSpecV1EngineParseUml("participant,lifeline", "message,return-message,create-message,destroy-message",
		umlRequiredElementsV1EngineParseUml{"participant or lifeline", umlTagSetV1EngineParseUml("participant,lifeline"), 1}),
}

var umlElementTagsV1EngineParseUml = map[string]bool{
	"element": true, "class": true, "interface": true, "enumeration": true,
	"component": true, "artifact": true, "port": true,
	"activity": true, "action": true, "decision": true, "merge": true, "fork": true, "join": true,
	"state": true, "initial": true, "final": true, "history": true, "choice": true,
	"participant": true, "lifeline": true, "object-node": true,
}

var umlRelationTagsV1EngineParseUml = map[string]bool{
	"relation": true, "association": true, "aggregation": true, "composition": true,
	"generalization": true, "realization": true, "dependency": true,
	"assembly": true, "delegation": true,
	"control-flow": true, "object-flow": true, "transition": true,
	"message": true, "return-message": true, "create-message": true, "destroy-message": true,
}

var umlCompartmentTagsV1EngineParseUml = map[string]bool{
	"compartment": true,
	"attribute":   true, "operation": true, "literal": true, "slot": true, "responsibility": true,
	"interface": true, "provided-interface": true, "required-interface": true, "property": true, "constraint": true,
	"entry": true, "do": true, "exit": true, "region": true, "note": true,
}

var umlElementCompartmentSpecsV1EngineParseUml = map[string]map[string]bool{
	"class":       umlTagSetV1EngineParseUml("attribute,operation,constraint,note"),
	"interface":   umlTagSetV1EngineParseUml("operation,constraint,note"),
	"enumeration": umlTagSetV1EngineParseUml("literal,operation,note"),
	"component":   umlTagSetV1EngineParseUml("interface,provided-interface,required-interface,property,constraint,note"),
	"artifact":    umlTagSetV1EngineParseUml("property,responsibility,note"),
	"activity":    umlTagSetV1EngineParseUml("responsibility,constraint,note"),
	"action":      umlTagSetV1EngineParseUml("responsibility,constraint,note"),
	"state":       umlTagSetV1EngineParseUml("entry,do,exit,region,note"),
}

func umlSpecV1EngineParseUml(elements, relations string, required ...umlRequiredElementsV1EngineParseUml) umlDiagramSpecV1EngineParseUml {
	return umlDiagramSpecV1EngineParseUml{
		elements:             umlTagSetV1EngineParseUml(elements),
		relations:            umlTagSetV1EngineParseUml(relations),
		requiredElementKinds: required,
	}
}

func umlTagSetV1EngineParseUml(value string) map[string]bool {
	result := map[string]bool{}
	for _, tag := range strings.Split(value, ",") {
		if tag = strings.TrimSpace(tag); tag != "" {
			result[tag] = true
		}
	}
	return result
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
	frameUMLIDs := map[*entity.Node]map[string]bool{}
	var walk func(*entity.Node, *entity.Node) error
	walk = func(node, frame *entity.Node) error {
		if node.Tag == "frame" {
			frame = node
		}
		if node.Tag == "uml" {
			if frame == nil {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<uml> must be inside a <frame>")}
			}
			rawID := node.Attr("id")
			id := strings.TrimSpace(rawID)
			if id != "" {
				if err := validateUMLIdentifierV1EngineParseUml("<uml> id", rawID); err != nil {
					return &entity.ParseError{Position: node.Position, Err: err}
				}
			}
			if frameUMLIDs[frame] == nil {
				frameUMLIDs[frame] = map[string]bool{}
			}
			if id != "" && frameUMLIDs[frame][id] {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("duplicate <uml id=%q> in frame", id)}
			}
			frameUMLIDs[frame][id] = true
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
	if len(uml.Children) != 1 {
		return &entity.ParseError{Position: uml.Position, Err: fmt.Errorf("<uml> requires exactly one supported UML diagram-kind child")}
	}
	diagram := uml.Children[0]
	spec, supported := umlDiagramSpecsV1EngineParseUml[diagram.Tag]
	if !supported {
		return &entity.ParseError{Position: diagram.Position, Err: fmt.Errorf("unsupported UML diagram kind <%s>", diagram.Tag)}
	}
	if strings.TrimSpace(uml.Attr("id")) == "" {
		return &entity.ParseError{Position: uml.Position, Err: fmt.Errorf("<uml> requires id")}
	}
	if strings.TrimSpace(uml.Attr("ref")) != "" {
		return &entity.ParseError{Position: uml.Position, Err: fmt.Errorf("<uml ref> is reserved; connect to an element with umlId/localId")}
	}
	if err := validateUMLDiagramAttributesV1EngineParseUml(diagram); err != nil {
		return &entity.ParseError{Position: diagram.Position, Err: err}
	}
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
	copyUMLDiagramLayoutAttrsV1EngineParseUml(uml, diagram)
	if uml.Attr("layout") == "" {
		if diagram.Attr("direction") == "right" || diagram.Tag == "sequence-diagram" {
			uml.Attrs["layout"] = "horizontal"
		} else {
			uml.Attrs["layout"] = "vertical"
		}
	}
	ids := map[string]bool{}
	elementKinds := map[string]string{}
	scopedIDs := map[string]string{}
	umlID := strings.TrimSpace(uml.Attr("id"))
	uml.Attrs["id"] = umlID
	var elements []*entity.Node
	var relations []*entity.Node
	normalizedElements := map[string]*entity.Node{}
	sourceElements := map[string]*entity.Node{}
	elementCounts := map[string]int{}
	partitions := map[string]string{}
	flattened, err := flattenUMLDiagramChildrenV1EngineParseUml(diagram, spec, partitions)
	if err != nil {
		return err
	}
	for _, entry := range flattened {
		child := entry.node
		switch {
		case umlElementTagsV1EngineParseUml[child.Tag]:
			rawID := child.Attr("id")
			id := strings.TrimSpace(rawID)
			if id == "" {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("UML <%s> requires id", child.Tag)}
			}
			if err := validateUMLIdentifierV1EngineParseUml(fmt.Sprintf("UML <%s> id", child.Tag), rawID); err != nil {
				return &entity.ParseError{Position: child.Position, Err: err}
			}
			if strings.TrimSpace(child.Attr("ref")) != "" {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("UML <%s ref> is reserved; use the generated umlId/localId reference", child.Tag)}
			}
			if ids[id] {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate UML element id %q", id)}
			}
			ids[id] = true
			child.Attrs["id"] = id
			elementKinds[id] = child.Tag
			sourceElements[id] = child
			elementCounts[child.Tag]++
			scopedIDs[id] = scopedUMLIDV1EngineParseUml(umlID, id)
			if err := validateUMLElementV1EngineParseUml(child); err != nil {
				return &entity.ParseError{Position: child.Position, Err: err}
			}
			normalized := normalizeUMLElementV1EngineParseUml(child, scopedIDs[id], diagram.Tag, umlID)
			if entry.partition != "" {
				normalized.Attrs["uml-partition-id"] = entry.partition
				normalized.Attrs["uml-partition-title"] = partitions[entry.partition]
			}
			normalizedElements[id] = normalized
			elements = append(elements, normalized)
		case umlRelationTagsV1EngineParseUml[child.Tag]:
			relations = append(relations, child)
		default:
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<%s> does not support UML child <%s>", diagram.Tag, child.Tag)}
		}
	}
	if err := validateUMLRequiredElementsV1EngineParseUml(diagram.Tag, spec, elementCounts); err != nil {
		return &entity.ParseError{Position: diagram.Position, Err: err}
	}
	ownerIDs := map[string]string{}
	for _, entry := range flattened {
		child := entry.node
		if !umlElementTagsV1EngineParseUml[child.Tag] {
			continue
		}
		owner, err := validateUMLOwnerV1EngineParseUml(diagram.Tag, child, elementKinds)
		if err != nil {
			return &entity.ParseError{Position: child.Position, Err: err}
		}
		if owner != "" {
			ownerIDs[child.Attr("id")] = owner
			normalizedElements[child.Attr("id")].Attrs["uml-owner-id"] = scopedIDs[owner]
			normalizedElements[child.Attr("id")].Attrs["uml-owner-ref"] = publicUMLRefV1EngineParseUml(umlID, owner)
		}
	}
	componentInterfaceLabels := map[string]map[string]bool{}
	for id, component := range normalizedElements {
		if component.Attr("uml-element-kind") != "component" {
			continue
		}
		labels := map[string]bool{}
		for _, line := range strings.Split(component.Attr("uml-component-interfaces"), "\n") {
			label := strings.TrimSpace(strings.SplitN(line, "\t", 2)[0])
			if label != "" {
				labels[label] = true
			}
		}
		componentInterfaceLabels[id] = labels
	}
	interfaceConnectionCounts := map[string]map[string]int{}
	for _, relation := range relations {
		src, dst := strings.TrimSpace(relation.Attr("src")), strings.TrimSpace(relation.Attr("dst"))
		relation.Attrs["src"] = src
		relation.Attrs["dst"] = dst
		if !ids[src] || !ids[dst] {
			return &entity.ParseError{Position: relation.Position, Err: fmt.Errorf("UML <%s> requires existing src and dst element IDs", relation.Tag)}
		}
		if err := validateUMLRelationEndpointsV1EngineParseUml(diagram.Tag, relation, elementKinds[src], elementKinds[dst]); err != nil {
			return &entity.ParseError{Position: relation.Position, Err: err}
		}
		if err := validateUMLRelationAttributesV1EngineParseUml(relation); err != nil {
			return &entity.ParseError{Position: relation.Position, Err: err}
		}
		if diagram.Tag == "component-diagram" && relation.Tag == "association" && elementKinds[src] == "component" && elementKinds[dst] == "component" {
			for _, id := range []string{src, dst} {
				component := normalizedElements[id]
				count, _ := strconv.Atoi(component.Attr("uml-component-connection-count"))
				component.Attrs["uml-component-connection-count"] = strconv.Itoa(count + 1)
			}
			var sharedLabels []string
			for label := range componentInterfaceLabels[src] {
				if componentInterfaceLabels[dst][label] {
					sharedLabels = append(sharedLabels, label)
				}
			}
			sort.Strings(sharedLabels)
			if len(sharedLabels) > 0 {
				if interfaceConnectionCounts[dst] == nil {
					interfaceConnectionCounts[dst] = map[string]int{}
				}
				label := sharedLabels[0]
				if interfaceConnectionCounts[dst][label] > 0 {
					component := normalizedElements[dst]
					extra, _ := strconv.Atoi(component.Attr("uml-component-interface-fanout-extra"))
					component.Attrs["uml-component-interface-fanout-extra"] = strconv.Itoa(extra + 1)
				}
				interfaceConnectionCounts[dst][label]++
			}
		}
	}
	if err := validateUMLRelationSetV1EngineParseUml(diagram.Tag, relations, sourceElements, elementKinds, ownerIDs); err != nil {
		return &entity.ParseError{Position: diagram.Position, Err: err}
	}
	for _, relation := range relations {
		src, dst := strings.TrimSpace(relation.Attr("src")), strings.TrimSpace(relation.Attr("dst"))
		connection := normalizeUMLRelationV1EngineParseUml(relation, scopedIDs[src], scopedIDs[dst], elementKinds[src], elementKinds[dst], diagram.Tag, umlID)
		if err := validateConnectionNodeV1EngineParseConnection(connection); err != nil {
			return &entity.ParseError{Position: relation.Position, Err: fmt.Errorf("invalid UML <%s>: %w", relation.Tag, err)}
		}
		frame.Children = append(frame.Children, connection)
	}
	uml.Children = elements
	return nil
}

func copyUMLDiagramLayoutAttrsV1EngineParseUml(uml, diagram *entity.Node) {
	for _, name := range []string{"grid", "component-width", "component-height"} {
		value, exists := diagram.Attrs[name]
		if _, inherited := uml.Attrs[name]; exists && !inherited {
			uml.Attrs[name] = strings.TrimSpace(value)
		}
	}
}

func flattenUMLDiagramChildrenV1EngineParseUml(diagram *entity.Node, spec umlDiagramSpecV1EngineParseUml, partitions map[string]string) ([]umlSourceElementV1EngineParseUml, error) {
	var flattened []umlSourceElementV1EngineParseUml
	for _, child := range diagram.Children {
		switch {
		case child.Tag == "partition":
			if diagram.Tag != "activity-diagram" {
				return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<%s> does not support UML child <partition>", diagram.Tag)}
			}
			id := strings.TrimSpace(child.Attr("id"))
			if id == "" {
				return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("UML <partition> requires id")}
			}
			if err := validateUMLIdentifierV1EngineParseUml("UML <partition> id", child.Attr("id")); err != nil {
				return nil, &entity.ParseError{Position: child.Position, Err: err}
			}
			if partitions[id] != "" {
				return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate UML partition id %q", id)}
			}
			title := strings.TrimSpace(child.Attr("title"))
			if title == "" {
				return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("UML <partition id=%q> requires title", id)}
			}
			partitions[id] = title
			for _, nested := range child.Children {
				if !umlElementTagsV1EngineParseUml[nested.Tag] {
					return nil, &entity.ParseError{Position: nested.Position, Err: fmt.Errorf("UML <partition> does not support child <%s>", nested.Tag)}
				}
				if !spec.elements[nested.Tag] {
					return nil, &entity.ParseError{Position: nested.Position, Err: fmt.Errorf("<%s> does not allow UML element <%s>", diagram.Tag, nested.Tag)}
				}
				if nestedLane := strings.TrimSpace(nested.Attr("lane")); nestedLane != "" && nestedLane != id {
					return nil, &entity.ParseError{Position: nested.Position, Err: fmt.Errorf("UML <%s lane=%q> conflicts with enclosing partition %q", nested.Tag, nestedLane, id)}
				}
				flattened = append(flattened, umlSourceElementV1EngineParseUml{node: nested, partition: id})
			}
		case umlElementTagsV1EngineParseUml[child.Tag]:
			if !spec.elements[child.Tag] {
				return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<%s> does not allow UML element <%s>", diagram.Tag, child.Tag)}
			}
			flattened = append(flattened, umlSourceElementV1EngineParseUml{node: child, partition: strings.TrimSpace(child.Attr("lane"))})
		case umlRelationTagsV1EngineParseUml[child.Tag]:
			if !spec.relations[child.Tag] {
				return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<%s> does not allow UML relation <%s>", diagram.Tag, child.Tag)}
			}
			flattened = append(flattened, umlSourceElementV1EngineParseUml{node: child})
		default:
			return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<%s> does not support UML child <%s>", diagram.Tag, child.Tag)}
		}
	}
	for _, entry := range flattened {
		if entry.partition == "" || partitions[entry.partition] != "" {
			continue
		}
		return nil, &entity.ParseError{Position: entry.node.Position, Err: fmt.Errorf("UML <%s lane=%q> references an unknown partition", entry.node.Tag, entry.partition)}
	}
	return flattened, nil
}

func scopedUMLIDV1EngineParseUml(umlID, localID string) string {
	return "uml-" + hex.EncodeToString([]byte(umlID)) + "-" + hex.EncodeToString([]byte(localID))
}

func publicUMLRefV1EngineParseUml(umlID, localID string) string {
	return umlID + "/" + localID
}

func validateUMLIdentifierV1EngineParseUml(description, value string) error {
	if strings.ContainsAny(value, "./") {
		return fmt.Errorf("%s=%q must not contain '.' or '/' because they delimit frame and UML references", description, value)
	}
	if strings.IndexFunc(value, unicode.IsSpace) >= 0 {
		return fmt.Errorf("%s=%q must not contain whitespace", description, value)
	}
	return nil
}

func validateUMLRequiredElementsV1EngineParseUml(diagramKind string, spec umlDiagramSpecV1EngineParseUml, counts map[string]int) error {
	for _, requirement := range spec.requiredElementKinds {
		count := 0
		for kind := range requirement.kinds {
			count += counts[kind]
		}
		if count < requirement.minimum {
			return fmt.Errorf("<%s> requires at least %d %s element(s)", diagramKind, requirement.minimum, requirement.description)
		}
	}
	return nil
}

func validateUMLOwnerV1EngineParseUml(diagramKind string, element *entity.Node, elementKinds map[string]string) (string, error) {
	owner := strings.TrimSpace(element.Attr("owner"))
	if diagramKind == "component-diagram" && element.Tag == "port" {
		if owner == "" {
			return "", fmt.Errorf("UML <%s> requires owner", element.Tag)
		}
		ownerKind, exists := elementKinds[owner]
		if !exists {
			return "", fmt.Errorf("UML <%s owner=%q> references an unknown element", element.Tag, owner)
		}
		if ownerKind != "component" {
			return "", fmt.Errorf("UML <%s owner=%q> does not allow owner kind <%s>", element.Tag, owner, ownerKind)
		}
		return owner, nil
	}
	if owner == "" {
		return "", nil
	}
	return "", fmt.Errorf("UML <%s> does not allow owner", element.Tag)
}

func validateUMLDiagramAttributesV1EngineParseUml(diagram *entity.Node) error {
	if direction := strings.ToLower(strings.TrimSpace(diagram.Attr("direction"))); direction != "" {
		if direction != "right" && direction != "down" {
			return fmt.Errorf("<%s direction=%q> must be right or down", diagram.Tag, direction)
		}
		diagram.Attrs["direction"] = direction
	}
	if lanes := strings.ToLower(strings.TrimSpace(diagram.Attr("lanes"))); lanes != "" {
		if diagram.Tag != "activity-diagram" {
			return fmt.Errorf("<%s> does not allow lanes", diagram.Tag)
		}
		if lanes != "vertical" && lanes != "horizontal" {
			return fmt.Errorf("<%s lanes=%q> must be vertical or horizontal", diagram.Tag, lanes)
		}
		diagram.Attrs["lanes"] = lanes
	}
	if theme := strings.ToLower(strings.TrimSpace(diagram.Attr("theme"))); theme != "" {
		if diagram.Tag != "activity-diagram" {
			return fmt.Errorf("<%s> does not allow theme", diagram.Tag)
		}
		if theme != "xaligo" {
			return fmt.Errorf("<%s theme=%q> must be xaligo", diagram.Tag, theme)
		}
		diagram.Attrs["theme"] = theme
	}
	return nil
}

func validateUMLElementV1EngineParseUml(element *entity.Node) error {
	allowed := umlElementCompartmentSpecsV1EngineParseUml[element.Tag]
	for _, compartment := range element.Children {
		if !umlCompartmentTagsV1EngineParseUml[compartment.Tag] {
			return fmt.Errorf("UML <%s> does not allow compartment <%s>", element.Tag, compartment.Tag)
		}
		if allowed == nil || (!allowed[compartment.Tag] && compartment.Tag != "compartment") {
			return fmt.Errorf("UML <%s> does not allow compartment <%s>", element.Tag, compartment.Tag)
		}
		if strings.TrimSpace(compartment.Text) == "" && strings.TrimSpace(compartment.Attr("title")) == "" && strings.TrimSpace(compartment.Attr("name")) == "" {
			return fmt.Errorf("UML compartment <%s> must contain text, title, or name", compartment.Tag)
		}
		if len(compartment.Children) != 0 {
			return fmt.Errorf("UML compartment <%s> must not contain child elements", compartment.Tag)
		}
	}
	for _, attribute := range []string{"from", "to"} {
		if _, exists := element.Attrs[attribute]; exists {
			return fmt.Errorf("UML <%s> does not allow %s", element.Tag, attribute)
		}
	}
	return nil
}

func validateUMLRelationEndpointsV1EngineParseUml(diagramKind string, relation *entity.Node, srcKind, dstKind string) error {
	relationKind := relation.Tag
	require := func(source, destination map[string]bool) error {
		if !source[srcKind] || !destination[dstKind] {
			return fmt.Errorf("UML <%s> does not allow %s -> %s endpoints", relationKind, srcKind, dstKind)
		}
		return nil
	}
	requireSameKind := func(allowed map[string]bool) error {
		if !allowed[srcKind] || srcKind != dstKind {
			return fmt.Errorf("UML <%s> requires equal endpoint kinds, got %s -> %s", relationKind, srcKind, dstKind)
		}
		return nil
	}
	switch diagramKind {
	case "class-diagram":
		classifiers := umlTagSetV1EngineParseUml("class,interface,enumeration")
		switch relationKind {
		case "association", "dependency":
			return require(classifiers, classifiers)
		case "aggregation", "composition":
			return require(umlTagSetV1EngineParseUml("class"), umlTagSetV1EngineParseUml("class"))
		case "generalization":
			return requireSameKind(classifiers)
		case "realization":
			return require(umlTagSetV1EngineParseUml("class"), umlTagSetV1EngineParseUml("interface"))
		}
	case "component-diagram":
		componentKinds := umlTagSetV1EngineParseUml("component,interface,port,artifact")
		switch relationKind {
		case "dependency", "association":
			return require(componentKinds, componentKinds)
		case "realization":
			return require(umlTagSetV1EngineParseUml("component"), umlTagSetV1EngineParseUml("interface"))
		case "assembly":
			if err := require(umlTagSetV1EngineParseUml("port,interface"), umlTagSetV1EngineParseUml("port,interface")); err != nil {
				return err
			}
			if srcKind != "port" && dstKind != "port" {
				return fmt.Errorf("UML <assembly> requires at least one port endpoint")
			}
			return nil
		case "delegation":
			return require(umlTagSetV1EngineParseUml("port"), umlTagSetV1EngineParseUml("component,port"))
		}
	case "activity-diagram":
		switch relationKind {
		case "control-flow":
			nodes := umlTagSetV1EngineParseUml("initial,final,activity,action,decision,merge,fork,join")
			return validateUMLFlowEndpointsV1EngineParseUml(relation, srcKind, dstKind, nodes)
		case "object-flow":
			nodes := umlTagSetV1EngineParseUml("activity,action,object-node")
			if err := require(nodes, nodes); err != nil {
				return err
			}
			if srcKind != "object-node" && dstKind != "object-node" {
				return fmt.Errorf("UML <object-flow> requires at least one object-node endpoint")
			}
			return nil
		}
	case "state-machine-diagram":
		nodes := umlTagSetV1EngineParseUml("initial,final,state,history,choice,fork,join")
		return validateUMLFlowEndpointsV1EngineParseUml(relation, srcKind, dstKind, nodes)
	case "sequence-diagram":
		if err := validateUMLMessageOrderV1EngineParseUml(relation); err != nil {
			return err
		}
		participants := umlTagSetV1EngineParseUml("participant,lifeline")
		if err := require(participants, participants); err != nil {
			return err
		}
		if (relationKind == "create-message" || relationKind == "destroy-message") && relation.Attr("src") == relation.Attr("dst") {
			return fmt.Errorf("UML <%s> does not allow a self message", relationKind)
		}
		return nil
	}
	return fmt.Errorf("UML <%s> has no endpoint contract for <%s>", diagramKind, relationKind)
}

func validateUMLFlowEndpointsV1EngineParseUml(relation *entity.Node, srcKind, dstKind string, allowed map[string]bool) error {
	if !allowed[srcKind] || !allowed[dstKind] {
		return fmt.Errorf("UML <%s> does not allow %s -> %s endpoints", relation.Tag, srcKind, dstKind)
	}
	if srcKind == "final" {
		return fmt.Errorf("UML <%s> must not leave a final node", relation.Tag)
	}
	if dstKind == "initial" {
		return fmt.Errorf("UML <%s> must not enter an initial node", relation.Tag)
	}
	return nil
}

func validateUMLMessageOrderV1EngineParseUml(relation *entity.Node) error {
	order := strings.TrimSpace(relation.Attr("order"))
	if order == "" {
		return fmt.Errorf("UML <%s> requires order", relation.Tag)
	}
	for _, part := range strings.Split(order, ".") {
		if len(part) > 1 && part[0] == '0' {
			return fmt.Errorf("UML <%s order=%q> must use canonical positive dot-separated integers without leading zeroes", relation.Tag, order)
		}
		value, err := strconv.Atoi(part)
		if err != nil || value <= 0 {
			return fmt.Errorf("UML <%s order=%q> must be positive dot-separated integers", relation.Tag, order)
		}
	}
	return nil
}

func validateUMLRelationAttributesV1EngineParseUml(relation *entity.Node) error {
	for _, attribute := range []string{"kind", "stroke-style", "arrowhead", "start-arrowhead", "end-arrowhead"} {
		if _, exists := relation.Attrs[attribute]; exists {
			return fmt.Errorf("UML <%s> does not allow %s; connector semantics are derived from the relation kind", relation.Tag, attribute)
		}
	}
	message := relation.Tag == "message" || relation.Tag == "return-message" || relation.Tag == "create-message" || relation.Tag == "destroy-message"
	if _, exists := relation.Attrs["order"]; exists && !message {
		return fmt.Errorf("UML <%s> does not allow order", relation.Tag)
	}
	if _, exists := relation.Attrs["guard"]; exists {
		switch relation.Tag {
		case "control-flow", "object-flow", "transition":
		default:
			return fmt.Errorf("UML <%s> does not allow guard", relation.Tag)
		}
	}
	for _, attribute := range []string{"src-multiplicity", "dst-multiplicity"} {
		if _, exists := relation.Attrs[attribute]; !exists {
			continue
		}
		switch relation.Tag {
		case "association", "aggregation", "composition":
		default:
			return fmt.Errorf("UML <%s> does not allow %s", relation.Tag, attribute)
		}
	}
	for _, attribute := range []string{"at", "from", "to"} {
		if _, exists := relation.Attrs[attribute]; exists {
			return fmt.Errorf("UML <%s> does not allow %s", relation.Tag, attribute)
		}
	}
	return nil
}

func validateUMLRelationSetV1EngineParseUml(diagramKind string, relations []*entity.Node, _ map[string]*entity.Node, elementKinds, _ map[string]string) error {
	orders := map[string]bool{}
	messageCount := 0
	incoming := map[string]int{}
	outgoing := map[string]int{}
	for _, relation := range relations {
		src := strings.TrimSpace(relation.Attr("src"))
		dst := strings.TrimSpace(relation.Attr("dst"))
		if relation.Tag == "message" || relation.Tag == "return-message" || relation.Tag == "create-message" || relation.Tag == "destroy-message" {
			order := strings.TrimSpace(relation.Attr("order"))
			if orders[order] {
				return fmt.Errorf("<%s> contains duplicate UML message order %q", diagramKind, order)
			}
			orders[order] = true
			messageCount++
		}
		if relation.Tag == "control-flow" || relation.Tag == "object-flow" || (diagramKind == "state-machine-diagram" && relation.Tag == "transition") {
			outgoing[src]++
			incoming[dst]++
		}
	}
	if diagramKind == "sequence-diagram" {
		orderedMessages := make([]*entity.Node, 0, messageCount)
		for _, relation := range relations {
			if relation.Tag == "message" || relation.Tag == "return-message" || relation.Tag == "create-message" || relation.Tag == "destroy-message" {
				orderedMessages = append(orderedMessages, relation)
			}
		}
		sort.SliceStable(orderedMessages, func(i, j int) bool {
			return compareUMLMessageOrderV1EngineParseUml(orderedMessages[i].Attr("order"), orderedMessages[j].Attr("order")) < 0
		})
		for index, relation := range orderedMessages {
			position := float64(index+1) / float64(len(orderedMessages)+1)
			relation.Attrs["uml-sequence-position"] = strconv.FormatFloat(position, 'f', 6, 64)
		}
	}
	if diagramKind == "activity-diagram" || diagramKind == "state-machine-diagram" {
		for id, kind := range elementKinds {
			if err := validateUMLControlNodeDegreeV1EngineParseUml(kind, incoming[id], outgoing[id]); err != nil {
				return fmt.Errorf("UML <%s id=%q>: %w", kind, id, err)
			}
		}
	}
	return nil
}

func compareUMLMessageOrderV1EngineParseUml(left, right string) int {
	leftParts, rightParts := strings.Split(left, "."), strings.Split(right, ".")
	length := len(leftParts)
	if len(rightParts) < length {
		length = len(rightParts)
	}
	for index := 0; index < length; index++ {
		leftValue, _ := strconv.Atoi(leftParts[index])
		rightValue, _ := strconv.Atoi(rightParts[index])
		if leftValue < rightValue {
			return -1
		}
		if leftValue > rightValue {
			return 1
		}
	}
	if len(leftParts) < len(rightParts) {
		return -1
	}
	if len(leftParts) > len(rightParts) {
		return 1
	}
	return 0
}

func validateUMLControlNodeDegreeV1EngineParseUml(kind string, incoming, outgoing int) error {
	switch kind {
	case "initial":
		if outgoing < 1 {
			return fmt.Errorf("requires at least one outgoing flow")
		}
	case "final":
		if incoming < 1 {
			return fmt.Errorf("requires at least one incoming flow")
		}
	case "decision", "fork", "choice":
		if incoming < 1 || outgoing < 2 {
			return fmt.Errorf("requires at least one incoming and two outgoing flows")
		}
	case "merge", "join":
		if incoming < 2 || outgoing < 1 {
			return fmt.Errorf("requires at least two incoming and one outgoing flow")
		}
	case "history":
		if outgoing < 1 {
			return fmt.Errorf("requires at least one outgoing transition")
		}
	}
	return nil
}

func normalizeUMLElementV1EngineParseUml(source *entity.Node, scopedID, diagramKind, umlID string) *entity.Node {
	attrs := cloneAttrsV1EngineParseTable(source.Attrs)
	displayName := attrs["name"]
	attrs["uml-local-id"] = source.Attr("id")
	attrs["id"] = scopedID
	attrs["ref"] = publicUMLRefV1EngineParseUml(umlID, source.Attr("id"))
	attrs["uml-id"] = umlID
	attrs["uml-ref"] = attrs["ref"]
	attrs["uml-diagram-kind"] = diagramKind
	attrs["uml-element-kind"] = source.Tag
	if strings.TrimSpace(attrs["font-family"]) == "" {
		attrs["font-family"] = "helvetica"
	}
	if strings.TrimSpace(attrs["font-size"]) == "" {
		attrs["font-size"] = "14"
	}
	if attrs["title"] == "" {
		attrs["title"] = displayName
	}
	if attrs["title"] == "" {
		attrs["title"] = source.Text
	}
	if attrs["title"] == "" {
		attrs["title"] = source.Attr("id")
	}
	// UML names are display text, not frame-level connection aliases. Public
	// endpoint references are exclusively uml-id/local-id, so retaining name on
	// the normalized rectangle would create false collisions between diagrams.
	delete(attrs, "name")
	var compartments []string
	var compartmentKinds []string
	var componentInterfaces []string
	for _, child := range source.Children {
		label := strings.TrimSpace(child.Attr("title"))
		if label == "" {
			label = strings.TrimSpace(child.Attr("name"))
		}
		if label == "" {
			label = strings.TrimSpace(child.Text)
		}
		if label != "" {
			if diagramKind == "component-diagram" && source.Tag == "component" && (child.Tag == "interface" || child.Tag == "provided-interface" || child.Tag == "required-interface") {
				componentInterfaces = append(componentInterfaces, label+"\t"+strings.TrimSpace(child.Attr("description")))
				continue
			}
			compartments = append(compartments, label)
			compartmentKinds = append(compartmentKinds, child.Tag)
		}
	}
	if len(compartments) > 0 {
		attrs["title"] += "\n────────\n" + strings.Join(compartments, "\n")
		attrs["uml-compartment-kinds"] = strings.Join(compartmentKinds, ",")
	}
	if len(componentInterfaces) > 0 {
		attrs["uml-component-interfaces"] = strings.Join(componentInterfaces, "\n")
	}
	return &entity.Node{Tag: "rectangle", Attrs: attrs, Position: source.Position}
}

func normalizeUMLRelationV1EngineParseUml(source *entity.Node, scopedSrc, scopedDst, srcKind, dstKind, diagramKind, umlID string) *entity.Node {
	attrs := cloneAttrsV1EngineParseTable(source.Attrs)
	attrs["src"] = scopedSrc
	attrs["dst"] = scopedDst
	attrs["uml-id"] = umlID
	attrs["uml-diagram-kind"] = diagramKind
	attrs["uml-relation-kind"] = source.Tag
	attrs["uml-relation-label"] = umlRelationLabelV1EngineParseUml(source)
	attrs["uml-src-ref"] = publicUMLRefV1EngineParseUml(umlID, source.Attr("src"))
	attrs["uml-dst-ref"] = publicUMLRefV1EngineParseUml(umlID, source.Attr("dst"))
	attrs["uml-src-kind"] = srcKind
	attrs["uml-dst-kind"] = dstKind
	for _, attribute := range []string{"order", "guard", "route", "src-multiplicity", "dst-multiplicity", "at", "from", "to"} {
		if value := strings.TrimSpace(source.Attr(attribute)); value != "" {
			attrs["uml-"+attribute] = value
		}
	}
	switch source.Tag {
	case "dependency", "realization", "return-message":
		attrs["stroke-style"] = "dashed"
		attrs["end-arrowhead"] = "triangle"
	case "generalization", "control-flow", "object-flow", "transition", "message", "create-message", "destroy-message", "delegation":
		attrs["end-arrowhead"] = "triangle"
	case "aggregation", "composition":
		attrs["start-arrowhead"] = "diamond"
		attrs["end-arrowhead"] = "none"
	case "association", "relation", "assembly":
		attrs["end-arrowhead"] = "none"
	}
	return &entity.Node{Tag: "connection", Attrs: attrs, Position: source.Position}
}

func umlRelationLabelV1EngineParseUml(source *entity.Node) string {
	label := strings.TrimSpace(source.Attr("label"))
	if label == "" {
		label = strings.TrimSpace(source.Attr("title"))
	}
	if label == "" {
		label = strings.TrimSpace(source.Text)
	}
	if guard := strings.TrimSpace(source.Attr("guard")); guard != "" {
		label = strings.TrimSpace(label + " [" + guard + "]")
	}
	if order := strings.TrimSpace(source.Attr("order")); order != "" {
		label = strings.TrimSpace(order + ": " + label)
	}
	if sourceMultiplicity, destinationMultiplicity := strings.TrimSpace(source.Attr("src-multiplicity")), strings.TrimSpace(source.Attr("dst-multiplicity")); sourceMultiplicity != "" || destinationMultiplicity != "" {
		label = strings.TrimSpace(label + " " + sourceMultiplicity + " → " + destinationMultiplicity)
	}
	return label
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
