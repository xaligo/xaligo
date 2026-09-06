package engine

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

func isAWSResourceV1EngineAwsResource(tag string) bool { return awsprofile.IsResourceTag(tag) }

func validateAWSResourceV1EngineAwsResource(node *entity.Node) error {
	loadBalancer := node.Tag == "aws-elastic-load-balancing-application-load-balancer" || node.Tag == "aws-elastic-load-balancing-network-load-balancer"
	if node.Tag == "aws-listener" || awsprofile.NativeChildKind(node.Tag) != "" || loadBalancer && (node.Attrs["view"] == "component" || node.Attrs["domain"] != "" || len(node.Children) > 0) {
		return fmt.Errorf("<%s> component requires XAL version 2", node.Tag)
	}
	definition, ok := awsprofile.DefinitionForTag(node.Tag)
	if !ok {
		return nil
	}
	if err := definition.ValidateAnnotations(node.Attrs); err != nil {
		return err
	}
	if definition.Group != nil || definition.Boundary != nil {
		return nil
	}
	if len(node.Children) != 0 || strings.TrimSpace(node.Text) != "" {
		return fmt.Errorf("<%s> must be empty; use label/detail attributes", node.Tag)
	}
	if err := validateConnectableFrameNodeV1EngineParseNode(node); err != nil {
		return err
	}
	for _, name := range []string{"icon", "icon-ref", "icon-id", "icon-width", "icon-height", "icon-scale", "icon-offset-x", "icon-offset-y", "side", "anchor", "offset"} {
		if _, exists := node.Attrs[name]; exists {
			return fmt.Errorf("<%s> does not support %s; the AWS icon is fixed and this is not a boundary tag", node.Tag, name)
		}
	}
	size := attrFloatV1EngineLayoutAttributes(node.Attr("size"), 48)
	width := attrFloatV1EngineLayoutAttributes(node.Attr("width"), math.Max(size+12, attrFloatV1EngineLayoutAttributes(node.Attr("label-width"), 160)))
	if node.Attr("width") == "" {
		node.Attrs["width"] = strconv.FormatFloat(width, 'f', -1, 64)
	}
	label := share.WrapPresentationText(definition.Label(node.Attrs), math.Max(1, width-8), itemLabelFontPxV1EngineSceneTypes)
	height := size + 12
	if label != "" {
		height = size + 4 + math.Ceil(float64(strings.Count(label, "\n")+1)*itemLabelFontPxV1EngineSceneTypes*1.25) + 12
	}
	if node.Attr("height") == "" {
		node.Attrs["height"] = strconv.FormatFloat(height, 'f', -1, 64)
	}
	if width < size+12 || attrFloatV1EngineLayoutAttributes(node.Attr("height"), 0) < height {
		return fmt.Errorf("<%s> width/height cannot contain its icon and label; need at least %.1fx%.1f", node.Tag, size+12, height)
	}
	return nil
}

func awsResourceLabelV1EngineAwsResource(tag string, attrs map[string]string, width float64) string {
	definition, _ := awsprofile.DefinitionForTag(tag)
	return share.WrapPresentationText(definition.Label(attrs), math.Max(1, width-8), itemLabelFontPxV1EngineSceneTypes)
}

func awsGroupLabelV1EngineAwsResource(node *entity.Node) (string, bool) {
	definition, ok := awsprofile.DefinitionForTag(node.Tag)
	if !ok || definition.Group == nil {
		return "", false
	}
	// Preserve legacy title/text/tag fallback when no new annotation is present.
	for _, parameter := range definition.Parameters {
		if node.Attrs[parameter.Name] != "" {
			return definition.Label(node.Attrs), true
		}
	}
	if node.Attrs["detail"] != "" || node.Attrs["label"] != "" {
		return definition.Label(node.Attrs), true
	}
	return "", false
}

func awsGroupDefinitionsV1EngineAwsResource() map[string]groupDefV1EngineSceneTypes {
	result := map[string]groupDefV1EngineSceneTypes{"capture": {StrokeColor: "#F5A623", StrokeStyle: "dashed", StrokeWidth: 1}}
	for _, definition := range awsprofile.Definitions() {
		if style := definition.Group; style != nil {
			result[definition.Tag] = groupDefV1EngineSceneTypes{StrokeColor: style.Stroke, StrokeStyle: style.Style, StrokeWidth: style.Width, IconFile: style.Icon}
		}
	}
	return result
}
