package aws

import (
	"fmt"
	"math"
	"net/netip"
	"slices"
	"strconv"
	"strings"
)

// Parameter describes diagram annotations, not AWS API/provisioning settings.
type Parameter struct {
	Name, Type, Description, Example string
	Values                           []string
}

type GroupStyle struct {
	Stroke, Style, Icon string
	Width               int
}

// Definition is a declarative authoring profile. Layout engines only receive
// generic item/group/port geometry; they never interpret AWS service names.
type Definition struct {
	Tag, Name, Kind, Category, Scope, Description string
	CatalogID                                     int
	CatalogIDs                                    []int
	Parameters                                    []Parameter
	Group                                         *GroupStyle
	Boundary                                      *BoundaryAttachment
}

var definitionIndex = buildDefinitionIndex()

func buildDefinitionIndex() map[string]int {
	result := make(map[string]int, len(definitions))
	for i, definition := range definitions {
		result[definition.Tag] = i
	}
	return result
}

func cloneDefinition(d Definition) Definition {
	d.CatalogIDs = slices.Clone(d.CatalogIDs)
	d.Parameters = slices.Clone(d.Parameters)
	for i := range d.Parameters {
		d.Parameters[i].Values = slices.Clone(d.Parameters[i].Values)
	}
	if d.Group != nil {
		value := *d.Group
		d.Group = &value
	}
	if d.Boundary != nil {
		value := *d.Boundary
		d.Boundary = &value
	}
	return d
}

func Definitions() []Definition {
	result := make([]Definition, len(definitions))
	for i, definition := range definitions {
		result[i] = cloneDefinition(definition)
	}
	return result
}

func DefinitionForTag(tag string) (Definition, bool) {
	i, ok := definitionIndex[tag]
	if !ok {
		return Definition{}, false
	}
	return cloneDefinition(definitions[i]), true
}

func IsResourceTag(tag string) bool {
	i, ok := definitionIndex[tag]
	return ok && definitions[i].Group == nil && definitions[i].Boundary == nil
}

// ValidateAnnotations checks the declarative parameter schema without imposing
// physical deployment constraints on a logical architecture diagram.
func (d Definition) ValidateAnnotations(attrs map[string]string) error {
	for _, parameter := range d.Parameters {
		value, exists := attrs[parameter.Name]
		if !exists {
			continue
		}
		value = strings.TrimSpace(value)
		valid := value != ""
		switch parameter.Type {
		case "enum":
			valid = slices.Contains(parameter.Values, value)
		case "boolean":
			valid = value == "true" || value == "false"
		case "integer":
			number, err := strconv.ParseUint(value, 10, 32)
			valid = err == nil && number <= math.MaxInt32
		case "port":
			number, err := strconv.ParseUint(value, 10, 16)
			valid = err == nil && number > 0
		case "cidr":
			_, err := netip.ParsePrefix(value)
			valid = err == nil
		}
		if !valid {
			return fmt.Errorf("<%s> %s=%q is not a valid %s (%s)", d.Tag, parameter.Name, value, parameter.Type, parameter.Description)
		}
	}
	if err := d.validateLoadBalancerAnnotations(attrs); err != nil {
		return err
	}
	if value, exists := attrs["show-details"]; exists && value != "true" && value != "false" {
		return fmt.Errorf("<%s> show-details must be true or false", d.Tag)
	}
	for _, name := range []string{"size", "label-width"} {
		if value, exists := attrs[name]; exists {
			number, err := strconv.ParseFloat(value, 64)
			if err != nil || number <= 0 || math.IsNaN(number) || math.IsInf(number, 0) {
				return fmt.Errorf("<%s> %s must be positive and finite", d.Tag, name)
			}
		}
	}
	return nil
}

// Label formats only explicitly supplied annotations, in schema order. No
// defaults imply that real cloud configuration has been inspected or applied.
func (d Definition) Label(attrs map[string]string) string {
	label := d.Name
	for _, key := range []string{"label", "title", "name"} {
		if value, exists := attrs[key]; exists {
			label = strings.TrimSpace(value)
			break
		}
	}
	if attrs["show-details"] == "false" {
		return label
	}
	parts := []string{label}
	for _, parameter := range d.Parameters {
		if value := strings.TrimSpace(attrs[parameter.Name]); value != "" {
			parts = append(parts, parameter.Name+": "+value)
		}
	}
	if detail := strings.TrimSpace(attrs["detail"]); detail != "" {
		parts = append(parts, detail)
	}
	separator := "\n"
	if d.Group != nil {
		separator = " · "
	}
	return strings.TrimSpace(strings.Join(parts, separator))
}
