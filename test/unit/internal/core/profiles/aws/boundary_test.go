package aws_test

import (
	"testing"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
)

func TestVPCEndpointBoundaryProfile(t *testing.T) {
	definition, ok := awsprofile.BoundaryAttachmentForTag(awsprofile.VPCEndpointTag)
	if !ok {
		t.Fatal("VPC endpoint boundary profile is missing")
	}
	if definition.ParentTag != "vpc" || definition.CatalogID != awsprofile.VPCEndpointCatalogID || definition.DefaultSide != "right" || definition.DefaultSize != 48 {
		t.Fatalf("VPC endpoint profile = %#v", definition)
	}
	if _, ok := awsprofile.BoundaryAttachmentForTag("unknown"); ok {
		t.Fatal("unknown AWS boundary tag resolved")
	}
}
