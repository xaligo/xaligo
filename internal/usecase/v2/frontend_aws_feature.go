package v2

import (
	"fmt"
	"strconv"
	"strings"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
)

func frontendAWSFeatureKind(tag string) string { return awsprofile.NativeChildKind(tag) }

func validateFrontendAWSFeature(node *frontendNode, allowed string) error {
	switch frontendAWSFeatureKind(node.tag) {
	case "rule":
		allowed += "priority "
	case "condition":
		allowed += "field name "
	case "match":
		allowed += "value key regex "
	case "action":
		allowed += "type order target-group "
	case "forward-target":
		allowed += "target-group weight "
	case "jwt-claim":
		allowed += "name format "
	case "transform":
		allowed += "type "
	case "rewrite":
		allowed += "regex replace "
	case "target-group":
		allowed += "name target-type protocol port "
	case "target-service":
		allowed += "service name ref "
	case "target":
		allowed += "name port zone "
	case "option":
		allowed += "name value key "
	}
	for name := range node.attrs {
		if !strings.Contains(allowed, " "+name+" ") {
			return fmt.Errorf("<%s> does not support %s", node.tag, name)
		}
	}
	return nil
}

func applyFrontendAWSFeature(node *frontendNode, element *entity.EngineElementSpec, model *entity.EngineAWSComponentSpec) error {
	model.Kind = frontendAWSFeatureKind(node.tag)
	switch model.Kind {
	case "rule":
		model.Value = node.attrs["priority"]
	case "condition":
		model.Type, model.Name = node.attrs["field"], node.attrs["name"]
	case "match":
		model.Value, model.Name = node.attrs["value"], node.attrs["key"]
		model.Type = "value"
		if value, ok := node.attrs["regex"]; ok {
			if _, both := node.attrs["value"]; both {
				return fmt.Errorf("<aws-rule-match> use value or regex, not both")
			}
			model.Type, model.Value = "regex", value
		}
	case "action":
		model.Type, model.Order, model.Value = node.attrs["type"], node.attrs["order"], node.attrs["target-group"]
	case "forward-target":
		model.Value, model.Order = node.attrs["target-group"], node.attrs["weight"]
		element.Weight = nil // AWS forwarding weight is not a generic layout weight.
	case "jwt-claim":
		model.Name, model.Type = node.attrs["name"], node.attrs["format"]
	case "transform":
		model.Type = node.attrs["type"]
	case "rewrite":
		model.Value, model.Aux = node.attrs["regex"], node.attrs["replace"]
	case "target-group":
		model.Name, model.Type, model.Value, model.Order = node.attrs["name"], node.attrs["target-type"], node.attrs["protocol"], node.attrs["port"]
	case "target":
		model.Name, model.Order, model.Value = node.attrs["name"], node.attrs["port"], node.attrs["zone"]
	case "target-service":
		model.Name, model.Type, model.Value = node.attrs["name"], node.attrs["service"], node.attrs["ref"]
		if d, ok := awsprofile.DefinitionForTag(awsprofile.TargetServiceIconTag(model.Type)); ok {
			element.Icon = &entity.EngineIconSpec{Ref: "catalog:" + strconv.Itoa(d.CatalogID), Width: frontendFloatPointer(24), Height: frontendFloatPointer(24), MissingPolicy: entity.EngineIconMissingError}
		}
	case "option":
		model.Type, model.Value, model.Name = node.attrs["name"], node.attrs["value"], node.attrs["key"]
	}
	element.AWS = model
	return nil
}
