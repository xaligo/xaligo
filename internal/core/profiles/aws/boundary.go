// Package aws contains declarative AWS authoring-profile data.
package aws

// BoundaryAttachment describes an AWS resource whose icon is attached to a
// semantic container boundary instead of participating in normal child flow.
// Layout engines consume these values through generic port/boundary concepts.
type BoundaryAttachment struct {
	Tag         string
	ParentTag   string
	CatalogID   int
	DefaultSide string
	DefaultSize float64
}

const (
	VPCEndpointTag       = "vpc-endpoint"
	VPCEndpointCatalogID = 1579
)

// BoundaryAttachmentForTag returns profile data for one semantic AWS boundary
// resource. The returned value is independent from the registry.
func BoundaryAttachmentForTag(tag string) (BoundaryAttachment, bool) {
	i, ok := definitionIndex[tag]
	if !ok || definitions[i].Boundary == nil {
		return BoundaryAttachment{}, false
	}
	return *definitions[i].Boundary, true
}
