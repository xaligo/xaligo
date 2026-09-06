package v2_test

import (
	"strings"
	"testing"

	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

func TestNativeAWSFrontendKeepsTypedParametersWithoutGeometry(t *testing.T) {
	source := `<xaligo version="2"><frames><frame id="p"><aws-elastic-load-balancing-network-load-balancer id="nlb" domain="API基盤.example.test"><aws-listener id="tcp" protocol="TCP" port="443" backend-tls="true" backend-mtls="false" target-group="app"/></aws-elastic-load-balancing-network-load-balancer></frame></frames></xaligo>`
	document, _, err := v2.NewFrontendUsecase().LowerWithProvenance([]byte(source))
	if err != nil {
		t.Fatal(err)
	}
	lb := document.Elements[0].Children[0]
	listener := lb.Children[0]
	if lb.AWS == nil || lb.AWS.Kind != "nlb" || lb.AWS.Domain != "API基盤.example.test" || lb.Width != nil || lb.Height != nil {
		t.Fatalf("native lowering: %#v", lb)
	}
	if listener.AWS == nil || listener.AWS.Kind != "listener" || listener.ID != "tcp" || *listener.AWS.Port != 443 || !*listener.AWS.BackendTLS || *listener.AWS.BackendMTLS {
		t.Fatalf("listener lowering: %#v", listener)
	}
	if listener.AWS.MutualTLS != "" || listener.Width != nil || listener.Text != nil || listener.Provenance.Tag != "aws-listener" {
		t.Fatalf("frontend must not precompute native design: %#v", listener)
	}
	if _, _, err := v2.NewFrontendUsecase().Lower([]byte(strings.Replace(source, `version="2"`, `version="1"`, 1))); err == nil || !strings.Contains(err.Error(), "requires XAL version 2") {
		t.Fatalf("V1 normalization accepted native component: %v", err)
	}
}

func TestNativeAWSFrontendRejectsInvalidStructureAndUnsupportedAttributes(t *testing.T) {
	const tag = "aws-elastic-load-balancing-application-load-balancer"
	for _, body := range []string{
		`<` + tag + ` id="alb"><aws-listener id="a" protocol="HTTPS" port="443" show-title="invalid"/></` + tag + `>`,
		`<aws-listener id="orphan" protocol="HTTPS" port="443"/>`,
		`<` + tag + ` id="alb" domain="api.example.test"/>`,
		`<` + tag + ` id="alb"><rectangle id="wrong"/></` + tag + `>`,
		`<` + tag + ` id="alb"><aws-listener protocol="HTTPS" port="443"/></` + tag + `>`,
		`<` + tag + ` id="alb"><aws-listener id="a" protocol="BOGUS" port="443"/></` + tag + `>`,
		`<` + tag + ` id="alb"><aws-listener id="a" protocol="HTTPS" port="0"/></` + tag + `>`,
		`<` + tag + ` id="alb"><aws-listener id="a" protocol="HTTPS" port="65536"/></` + tag + `>`,
		`<` + tag + ` id="alb"><aws-listener id="a" protocol="HTTPS" port="443" tls="off"/></` + tag + `>`,
		`<` + tag + ` id="alb"><aws-listener id="a" protocol="HTTPS" port="443" width="100"/></` + tag + `>`,
		`<` + tag + ` id="alb"><aws-listener id="a" protocol="HTTPS" port="443" backend-tls="invalid"/></` + tag + `>`,
		`<` + tag + ` id="alb" view="icon"><aws-listener id="a" protocol="HTTPS" port="443"/></` + tag + `>`,
		`<` + tag + ` id="alb" scheme="internet-facing"><aws-listener id="a" protocol="HTTPS" port="443"/></` + tag + `>`,
	} {
		source := `<xaligo version="2"><frames><frame id="p">` + body + `</frame></frames></xaligo>`
		if _, _, err := v2.NewFrontendUsecase().Lower([]byte(source)); err == nil {
			t.Errorf("accepted %s", body)
		}
	}
	if _, _, err := v2.NewFrontendUsecase().Lower([]byte(strings.Replace(`<xaligo version="2"><frames><frame id="p"><TAG id="alb"/></frame></frames></xaligo>`, "TAG", tag, 1))); err != nil {
		t.Fatalf("legacy icon form: %v", err)
	}
}

func TestNativeAWSListenerTitleVisibilityIsOptional(t *testing.T) {
	for _, value := range []string{"", "true", "false"} {
		attr := ""
		if value != "" {
			attr = ` show-title="` + value + `"`
		}
		source := `<xaligo version="2"><frames><frame id="p"><aws-elastic-load-balancing-network-load-balancer id="lb"><aws-listener id="l" protocol="TCP" port="443"` + attr + `/></aws-elastic-load-balancing-network-load-balancer></frame></frames></xaligo>`
		document, _, err := v2.NewFrontendUsecase().Lower([]byte(source))
		if err != nil {
			t.Fatal(err)
		}
		got := document.Elements[0].Children[0].Children[0].AWS.ShowTitle
		if value == "" {
			if got != nil {
				t.Fatal("unset title flag must stay unset")
			}
		} else if got == nil || *got != (value == "true") {
			t.Fatalf("show-title=%s: %v", value, got)
		}
	}
}

func TestNativeALBFeaturesKeepTypedFieldsAndZeroForwardWeight(t *testing.T) {
	const source = `<xaligo version="2"><frames><frame id="p"><aws-elastic-load-balancing-application-load-balancer id="alb" detail-level="summary" show="rules" hide="certificate"><aws-listener id="l" protocol="HTTPS" port="443"><aws-listener-rule id="r" priority="default"><aws-rule-action id="a" type="forward"><aws-forward-target id="canary" target-group="app" weight="0" visible="false"/></aws-rule-action></aws-listener-rule></aws-listener></aws-elastic-load-balancing-application-load-balancer></frame></frames></xaligo>`
	document, _, err := v2.NewFrontendUsecase().LowerWithProvenance([]byte(source))
	if err != nil {
		t.Fatal(err)
	}
	lb := document.Elements[0].Children[0]
	if lb.AWS.DetailLevel != "summary" || lb.AWS.Show != "rules" || lb.AWS.Hide != "certificate" {
		t.Fatalf("presentation lost: %#v", lb.AWS)
	}
	rule := lb.Children[0].Children[0]
	target := rule.Children[0].Children[0]
	if rule.AWS.Kind != "rule" || rule.AWS.Value != "default" || target.AWS.Kind != "forward-target" || target.AWS.Value != "app" || target.AWS.Order != "0" || target.Weight != nil {
		t.Fatalf("typed routing fields lost: %#v / %#v", rule.AWS, target)
	}
	if target.Provenance.Tag != "aws-forward-target" || target.ID != "canary" {
		t.Fatal("source identity lost")
	}
	for _, invalid := range []string{
		strings.Replace(source, `weight="0"`, `weight="0" port="80"`, 1),
		strings.Replace(source, `version="2"`, `version="1"`, 1),
	} {
		if _, _, err := v2.NewFrontendUsecase().Lower([]byte(invalid)); err == nil {
			t.Fatal("accepted foreign attributes or V1 native rules")
		}
	}
}
