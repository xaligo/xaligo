package v2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

func validateFrontendAWSResource(node *frontendNode) error {
	if frontendAWSComponent(node) {
		return validateFrontendAWSComponent(node)
	}
	if frontendAWSLoadBalancer(node.tag) != "" && node.attrs["view"] != "" && node.attrs["view"] != "icon" {
		return fmt.Errorf("<%s> view must be icon or component", node.tag)
	}
	definition, ok := awsprofile.DefinitionForTag(node.tag)
	if !ok {
		return nil
	}
	if err := definition.ValidateAnnotations(node.attrs); err != nil {
		return err
	}
	if definition.Group != nil || definition.Boundary != nil {
		return nil
	}
	id := strings.TrimSpace(node.attrs["id"])
	if id == "" || strings.ContainsAny(id, " \t\r\n") {
		return fmt.Errorf("<%s> requires a non-empty whitespace-free id", node.tag)
	}
	if len(node.children) != 0 || strings.TrimSpace(node.text) != "" {
		return fmt.Errorf("<%s> must be empty; use label/detail attributes", node.tag)
	}
	for _, name := range []string{"icon", "icon-ref", "icon-id", "icon-width", "icon-height", "icon-scale", "icon-offset-x", "icon-offset-y", "side", "anchor", "offset"} {
		if _, exists := node.attrs[name]; exists {
			return fmt.Errorf("<%s> does not support %s; the AWS icon is fixed and this is not a boundary tag", node.tag, name)
		}
	}
	return nil
}

func applyFrontendAWSResource(node *frontendNode, element *entity.EngineElementSpec) error {
	if frontendAWSComponent(node) {
		return applyFrontendAWSComponent(node, element)
	}
	if !awsprofile.IsResourceTag(node.tag) {
		return nil
	}
	definition, _ := awsprofile.DefinitionForTag(node.tag)
	size, err := frontendNumber(node, "size", 48)
	if err != nil {
		return err
	}
	labelWidth, err := frontendNumber(node, "label-width", 160)
	if err != nil {
		return err
	}
	frontendSetNumberDefault(&element.Width, math.Max(size+12, labelWidth))
	label := share.WrapPresentationText(definition.Label(node.attrs), math.Max(1, *element.Width-8), frontendV1ItemLabelFontSize)
	height := size + 12
	element.Text = nil
	if label != "" {
		height += 4 + math.Ceil(float64(strings.Count(label, "\n")+1)*frontendV1ItemLabelFontSize*frontendV1ItemLabelLineHeight)
		element.Text = &entity.EngineTextSpec{Value: label, Role: string(entity.TextRoleItemLabel), FontFamily: "Helvetica", FontSize: frontendFloatPointer(frontendV1ItemLabelFontSize), LineHeight: frontendFloatPointer(frontendV1ItemLabelLineHeight)}
	}
	frontendSetNumberDefault(&element.Height, height)
	if *element.Width < size+12 || *element.Height < height {
		return fmt.Errorf("<%s> width/height cannot contain its icon and label; need at least %.1fx%.1f", node.tag, size+12, height)
	}
	element.Icon = &entity.EngineIconSpec{Ref: "catalog:" + strconv.Itoa(definition.CatalogID), Width: frontendFloatPointer(size), Height: frontendFloatPointer(size), MissingPolicy: entity.EngineIconMissingFallback}
	element.Visual.Shape = entity.EngineShapeNone
	return nil
}

func frontendAWSGroupProfiles() map[string]frontendV1GroupProfile {
	result := map[string]frontendV1GroupProfile{"capture": {stroke: "#F5A623", strokeWidth: 1, strokeStyle: entity.EngineLineDashed}}
	for _, definition := range awsprofile.Definitions() {
		if style := definition.Group; style != nil {
			result[definition.Tag] = frontendV1GroupProfile{stroke: style.Stroke, strokeWidth: float64(style.Width), strokeStyle: entity.EngineLineStyle(style.Style), icon: style.Icon}
		}
	}
	return result
}
