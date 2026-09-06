package v2_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestEveryAWSExampleLowersWithSemanticTagAndCatalogIcon(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "../../../../.."))
	for _, definition := range awsprofile.Definitions() {
		t.Run(definition.Tag, func(t *testing.T) {
			source, err := os.ReadFile(filepath.Join(root, "docs/src/examples/samples/aws", definition.Tag, "sample.xal"))
			if err != nil {
				t.Fatal(err)
			}
			document, _, err := v2.NewFrontendUsecase().LowerWithProvenance(source)
			if err != nil {
				t.Fatal(err)
			}
			var found *entity.EngineElementSpec
			var visit func([]entity.EngineElementSpec)
			visit = func(elements []entity.EngineElementSpec) {
				for i := range elements {
					e := &elements[i]
					if e.Provenance != nil && e.Provenance.Tag == definition.Tag {
						found = e
					}
					visit(e.Children)
				}
			}
			visit(document.Elements)
			if found == nil {
				t.Fatal("dedicated tag was lost")
			}
			want := entity.EngineConceptItem
			if definition.Group != nil || found.AWS != nil {
				want = entity.EngineConceptGroup
			}
			if definition.Boundary != nil {
				want = entity.EngineConceptPort
			}
			if found.Concept != want {
				t.Fatalf("concept=%s, want %s", found.Concept, want)
			}
			if definition.Group == nil && (found.Icon == nil || found.Icon.Ref != fmt.Sprintf("catalog:%d", definition.CatalogID)) {
				t.Fatalf("icon=%#v", found.Icon)
			}
			// The same editable sample grammar must remain valid in standalone V1.
			v1Source := strings.Replace(string(source), `version="2"`, `version="1"`, 1)
			parsed, err := usecase.Parse(strings.NewReader(v1Source))
			if found.AWS != nil {
				if err == nil || !strings.Contains(err.Error(), "requires XAL version 2") {
					t.Fatalf("native component V1 boundary: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if _, err := usecase.Build(parsed); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestAWSResourceValidationParity(t *testing.T) {
	for _, body := range []string{
		`<aws-ec2/>`, `<aws-ec2 id="bad id"/>`, `<aws-ec2 id="a" state="invalid"/>`,
		`<aws-ec2 id="a" size="NaN"/>`, `<aws-ec2 id="a" label-width="0"/>`,
		`<aws-ec2 id="a" icon-id="1"/>`, `<aws-ec2 id="a"><aws-s3 id="b"/></aws-ec2>`,
		`<aws-ec2 id="a" side="right"/>`, `<aws-ec2 id="a" height="8"/>`, `<aws-ec2 id="a" width="8"/>`,
		`<aws-rds id="a" multi-az="yes"/>`, `<vpc id="a" cidr="not-a-prefix"/>`,
		`<aws-elastic-load-balancing-network-load-balancer id="a" trust-store="ca"/>`,
		`<aws-elastic-load-balancing-network-load-balancer id="a" listener-port="65536"/>`,
		`<aws-elastic-load-balancing-application-load-balancer id="a" listener-protocol="TCP"/>`,
		`<aws-elastic-load-balancing-application-load-balancer id="a" listener-protocol="HTTP" mutual-tls-mode="verify"/>`,
	} {
		t.Run(body, func(t *testing.T) {
			for _, version := range []string{"1", "2"} {
				source := `<xaligo version="` + version + `"><frames><frame id="p" width="600" height="400">` + body + `</frame></frames></xaligo>`
				var err error
				if version == "1" {
					_, err = usecase.Parse(strings.NewReader(source))
				} else {
					_, _, err = v2.NewFrontendUsecase().Lower([]byte(source))
				}
				if err == nil {
					t.Errorf("V%s accepted invalid AWS resource", version)
				}
			}
		})
	}
}

func TestAWSBoundaryDefaultAndExplicitSidesShareAutoDistribution(t *testing.T) {
	source := `<xaligo version="2"><frames><frame id="p"><vpc id="v"><vpc-endpoint id="a"/><vpc-endpoint id="b" side="right"/></vpc></frame></frames></xaligo>`
	document, _, err := v2.NewFrontendUsecase().Lower([]byte(source))
	if err != nil {
		t.Fatal(err)
	}
	ports := document.Elements[0].Children[0].Children
	if *ports[0].Port.Anchor != 1.0/3 || *ports[1].Port.Anchor != 2.0/3 {
		t.Fatal("default and explicit right sides did not share the same edge")
	}
}

func TestAWSResourceSizeDoesNotInheritLegacyItemGridDefaults(t *testing.T) {
	source := `<xaligo version="2"><frames><frame id="p" item-size="24"><aws-ec2 id="instance" label="API基盤" size="64"/><item id="27"/></frame></frames></xaligo>`
	document, _, err := v2.NewFrontendUsecase().Lower([]byte(source))
	if err != nil {
		t.Fatal(err)
	}
	resource := document.Elements[0].Children[0]
	legacy := document.Elements[0].Children[1]
	if *resource.Width != 160 || *resource.Icon.Width != 64 || *legacy.Icon.Width != 24 {
		t.Fatalf("dedicated and legacy size defaults were mixed: %#v %#v", resource, legacy)
	}
}
