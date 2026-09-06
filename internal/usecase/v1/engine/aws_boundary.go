package engine

import (
	"fmt"
	"strconv"
	"strings"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
)

func awsBoundaryAttachmentV1EngineAwsBoundary(tag string) (awsprofile.BoundaryAttachment, bool) {
	return awsprofile.BoundaryAttachmentForTag(tag)
}

func isAWSBoundaryAttachmentV1EngineAwsBoundary(tag string) bool {
	_, ok := awsBoundaryAttachmentV1EngineAwsBoundary(tag)
	return ok
}

func awsBoundaryCatalogIDV1EngineAwsBoundary(tag string) string {
	definition, ok := awsBoundaryAttachmentV1EngineAwsBoundary(tag)
	if !ok {
		return ""
	}
	return strconv.Itoa(definition.CatalogID)
}

func validateAWSBoundaryAttachmentsV1EngineAwsBoundary(root *entity.Node) error {
	var walk func(*entity.Node, *entity.Node) error
	walk = func(node, parent *entity.Node) error {
		if node == nil {
			return nil
		}
		if err := validateAWSResourceV1EngineAwsResource(node); err != nil {
			return &entity.ParseError{Position: node.Position, Err: err}
		}
		if definition, ok := awsBoundaryAttachmentV1EngineAwsBoundary(node.Tag); ok {
			if parent == nil || parent.Tag != definition.ParentTag {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<%s> must be a direct child of <%s>", node.Tag, definition.ParentTag)}
			}
			if len(node.Children) != 0 || strings.TrimSpace(node.Text) != "" {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<%s> must be empty", node.Tag)}
			}
		}
		for _, child := range node.Children {
			if err := walk(child, node); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(root, nil)
}
