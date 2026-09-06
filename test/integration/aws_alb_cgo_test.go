//go:build cgo && xaligo_engine

package integration_test

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/repository"
	"github.com/xaligo/xaligo/internal/usecase"
	v2 "github.com/xaligo/xaligo/internal/usecase/v2"
)

const albTag = "aws-elastic-load-balancing-application-load-balancer"

func albSource(body, attrs string) []byte {
	return []byte(`<xaligo version="2"><frames><frame id="p" width="1800" height="2800"><` + albTag + ` id="alb" domain="api.example.test" ` + attrs + `>` + body + `</` + albTag + `></frame></frames></xaligo>`)
}
func albListener(body string) string {
	return `<aws-listener id="listener" protocol="HTTPS" port="443">` + body + `</aws-listener>`
}
func albRule(priority, body string) string {
	return `<aws-listener-rule id="rule" priority="` + priority + `">` + body + `</aws-listener-rule>`
}
func albCondition(field, value string) string {
	return `<aws-rule-condition id="condition" field="` + field + `"><aws-rule-match id="match" value="` + value + `"/></aws-rule-condition>`
}

const albForward = `<aws-rule-action id="forward" type="forward" target-group="app"/>`

func TestALBRulesHiddenDetailsKeepValidationAndConnectionIdentity(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	source := albSource(albListener(albRule("10", albCondition("path-pattern", "/private/*")+albForward)), "detail-level=\"detailed\"")
	source = bytes.Replace(source, []byte("</frame>"), []byte(`<rectangle id="client" width="120" height="60"/><connection id="entry" src="client" dst="rule" /></frame>`), 1)
	detailed, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	small := bytes.Replace(source, []byte(`detail-level="detailed"`), []byte(`detail-level="summary"`), 1)
	summary, err := renderer.RenderSVG(context.Background(), small, entity.RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(summary, []byte("/private/*")) || !bytes.Contains(detailed, []byte("/private/*")) {
		t.Fatal("summary exposed hidden rule details")
	}
	if svgElementGeometry(t, summary, "alb")[3] >= svgElementGeometry(t, detailed, "alb")[3] {
		t.Fatal("summary did not remove hidden layout space")
	}
	normalized, _, err := v2.NewFrontendUsecase().Lower(small)
	if err != nil {
		t.Fatal(err)
	}
	resolved, err := v2.NewEngineUsecase().Resolve(context.Background(), normalized)
	if err != nil {
		t.Fatal(err)
	}
	var anchor, edge entity.EngineResolvedElement
	for _, e := range resolved.Elements {
		if e.ID == "listener" {
			anchor = e
		}
		if e.ID == "entry" {
			edge = e
		}
	}
	if len(edge.Points) == 0 {
		t.Fatal("missing retained connection")
	}
	end := edge.Points[len(edge.Points)-1]
	near := func(a, b float64) bool { return math.Abs(a-b) < 0.01 }
	if !((near(end.X, anchor.X) || near(end.X, anchor.X+anchor.Width)) && near(end.Y, anchor.Y+anchor.Height/2) || (near(end.Y, anchor.Y) || near(end.Y, anchor.Y+anchor.Height)) && near(end.X, anchor.X+anchor.Width/2)) {
		t.Fatalf("collapsed endpoint=%#v, listener=%#v", end, anchor)
	}
	plan, err := renderer.BuildPPTXPlan(context.Background(), small, entity.RenderOptions{PxPerInch: 96})
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(plan, []byte("/private/*")) {
		t.Fatal("PPTX exposes hidden detail")
	}
	for _, bad := range [][]byte{
		bytes.Replace(small, []byte(`priority="10"`), []byte(`priority="default"`), 1),
		bytes.Replace(small, []byte(`field="path-pattern"`), []byte(`field="source-ip"`), 1),
	} {
		if _, err := renderer.RenderSVG(context.Background(), bad, entity.RenderOptions{}); err == nil {
			t.Fatal("summary bypassed validation")
		}
	}
}

func TestALBRulesRejectInvalidServiceCombinations(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	path := albCondition("path-pattern", "/api/*")
	cases := []struct{ name, body, attrs string }{
		{"unknown option", albListener(`<aws-option id="o" name="made-up-setting" value="true"/>`), ""},
		{"wrong owner", albListener(`<aws-option id="o" name="scheme" value="internal"/>`), ""},
		{"bad bool", albListener("") + `<aws-option id="o" name="deletion_protection.enabled" value="yes"/>`, ""},
		{"bad number", albListener("") + `<aws-option id="o" name="idle_timeout.timeout_seconds" value="0"/>`, ""},
		{"unknown display", albListener(""), `hide="secret-mode"`},
		{"conflicting display", albListener(""), `hide="tls" show="tls"`},
		{"unknown level", albListener(""), `detail-level="minimal"`},
		{"default conditions", albListener(albRule("default", path+albForward)), ""},
		{"nondefault no condition", albListener(albRule("10", albForward)), ""},
		{"zero priority", albListener(albRule("0", path+albForward)), ""},
		{"priority max", albListener(albRule("50001", path+albForward)), ""},
		{"no route action", albListener(albRule("10", path)), ""},
		{"two route actions", albListener(albRule("10", path+albForward+strings.Replace(albForward, `id="forward"`, `id="another" order="2"`, 1))), ""},
		{"regex on IP", albListener(albRule("10", strings.Replace(albCondition("source-ip", "10.0.0.0/8"), `value=`, `regex=`, 1)+albForward)), ""},
		{"header missing name", albListener(albRule("10", albCondition("http-header", "mobile")+albForward)), ""},
		{"too many conditions", albListener(albRule("10", path+strings.ReplaceAll(path, `id="`, `id="other-`)+albForward)), ""},
		{"target weight range", albListener(albRule("10", path+`<aws-rule-action id="f" type="forward"><aws-forward-target id="w" target-group="app" weight="1000"/></aws-rule-action>`)), ""},
		{"auth without configuration", albListener(albRule("10", path+`<aws-rule-action id="auth" type="jwt-validation" order="1"/><aws-rule-action id="forward" type="forward" order="2" target-group="app"/>`)), ""},
		{"redirect downgrade", albListener(albRule("default", `<aws-rule-action id="a" type="redirect"><aws-option id="s" name="status-code" value="HTTP_301"/><aws-option id="p" name="protocol" value="HTTP"/></aws-rule-action>`)), ""},
		{"fixed 3xx", albListener(albRule("default", `<aws-rule-action id="a" type="fixed-response"><aws-option id="s" name="status-code" value="302"/></aws-rule-action>`)), ""},
		{"lambda port", albListener("") + `<aws-target-group id="tg" name="fn" target-type="lambda" port="80"/>`, ""},
		{"grpc on HTTP", `<aws-listener id="http" protocol="HTTP" port="80" target-group="app"/><aws-target-group id="app" name="app" target-type="ip" protocol="HTTP" port="50051"><aws-option id="v" name="protocol-version" value="GRPC"/></aws-target-group>`, ""},
		{"HTTP TLS policy", `<aws-listener id="http" protocol="HTTP" port="80"><aws-option id="tls" name="tls-policy" value="policy"/></aws-listener>`, "detail-level=\"summary\""},
		{"slow-start algorithm", albListener("") + `<aws-target-group id="app" name="app" target-type="ip" protocol="HTTP" port="80"><aws-option id="a" name="load_balancing.algorithm.type" value="weighted_random"/><aws-option id="s" name="slow_start.duration_seconds" value="60"/></aws-target-group>`, ""},
		{"stickiness algorithm", albListener("") + `<aws-target-group id="app" name="app" target-type="ip" protocol="HTTP" port="80"><aws-option id="a" name="load_balancing.algorithm.type" value="weighted_random"/><aws-option id="s" name="stickiness.enabled" value="true"/></aws-target-group>`, ""},
		{"multiple Lambda targets", albListener("") + `<aws-target-group id="app" name="app" target-type="lambda"><aws-registered-target id="a" name="fn-a"/><aws-registered-target id="b" name="fn-b"/></aws-target-group>`, ""},
		{"Fargate instance target", albListener("") + `<aws-target-group id="app" name="app" target-type="instance" protocol="HTTP" port="80"><aws-target-service id="ecs" service="ecs" name="api"><aws-option id="launch" name="launch-type" value="fargate"/><aws-option id="network" name="network-mode" value="awsvpc"/></aws-target-service></aws-target-group>`, ""},
		{"EKS pod with NodePort", albListener("") + `<aws-target-group id="app" name="app" target-type="ip" protocol="HTTP" port="80"><aws-target-service id="eks" service="eks" name="api"><aws-option id="node" name="node-port" value="30080"/></aws-target-service></aws-target-group>`, ""},
		{"Lambda service in IP group", albListener("") + `<aws-target-group id="app" name="app" target-type="ip" protocol="HTTP" port="80"><aws-target-service id="fn" service="lambda" name="invoice"/></aws-target-group>`, ""},
		{"unknown target type", albListener("") + `<aws-target-group id="tg" name="fn" target-type="alb"/>`, ""},
		{"bad hierarchy", albListener(`<aws-rule-match id="m" value="*"/>`), ""},
		{"orphan rule", albRule("default", albForward), ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := renderer.RenderSVG(context.Background(), albSource(tc.body, tc.attrs), entity.RenderOptions{}); err == nil {
				t.Fatal("accepted invalid ALB input")
			}
		})
	}
	nlb := bytes.ReplaceAll(albSource(albListener(albRule("10", path+albForward)), ""), []byte(albTag), []byte("aws-elastic-load-balancing-network-load-balancer"))
	if _, err := renderer.RenderSVG(context.Background(), nlb, entity.RenderOptions{}); err == nil {
		t.Fatal("NLB accepted ALB routing")
	}
}

func TestALBRulesOrderAndSelectiveVisibility(t *testing.T) {
	late := albRule("20", albCondition("path-pattern", "/late/*")+albForward)
	early := strings.ReplaceAll(albRule("10", albCondition("host-header", "early.example.test")+albForward), `id="`, `id="early-`)
	def := strings.ReplaceAll(albRule("default", albForward), `id="`, `id="default-`)
	source := albSource(albListener(late+def+early), `detail-level="detailed"`)
	normalized, _, err := v2.NewFrontendUsecase().Lower(source)
	if err != nil {
		t.Fatal(err)
	}
	resolved, err := v2.NewEngineUsecase().Resolve(context.Background(), normalized)
	if err != nil {
		t.Fatal(err)
	}
	positions := map[string]float64{}
	for _, e := range resolved.Elements {
		positions[e.ID] = e.Y
	}
	if !(positions["listener::aws/rule-early-rule-priority"] < positions["listener::aws/rule-rule-priority"] && positions["listener::aws/rule-rule-priority"] < positions["listener::aws/rule-default-rule-priority"]) {
		t.Fatal("rules are not ordered by priority with default last", positions)
	}
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	source = bytes.Replace(source, []byte(`id="early-condition"`), []byte(`id="early-condition" visible="false"`), 1)
	output, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(output, []byte("early.example.test")) || !bytes.Contains(output, []byte("/late/*")) {
		t.Fatal("individual condition visibility leaked or suppressed sibling content")
	}
}

func TestALBFeatureSamplesShareVisibleSVGAndPPTX(t *testing.T) {
	renderer := usecase.NewRenderUsecase(repository.NewSceneRepository(), repository.NewXaligoRepository(), repository.NewPowerpointRepository(), repository.NewSVGRepository(), repository.NewTerminalRepository())
	for file, expected := range map[string][]string{
		"conditions":             {"Host", "Path", "Header", "Method", "Query", "Source IP"},
		"rewrites":               {"Rewrite host", "Rewrite URL", "/v2/$1"},
		"authentication-oidc":    {"OIDC", "Issuer", "Authorize"},
		"authentication-cognito": {"Cognito", "User pool"},
		"authentication-jwt":     {"JWT", "single-string", "string-array", "space-separated-values"},
		"redirect-response":      {"Redirect", "Response", "HTTP_301", "503"},
		"weighted-forward":       {"w=90", "w=10", "Group stickiness", "ON"},
		"target-types":           {"TG /", "ip-app", "instance-app", "lambda-app"},
		"grpc":                   {"GRPC", "/orders.v1.Service/*"},
		"ecs-fargate":            {"ECS / 注文API", "Launch type", "fargate", "Network mode", "awsvpc"},
		"ecs-ec2":                {"ECS / API on EC2", "bridge", "32768", "32769"},
		"eks-targets":            {"EKS / orders / pod mode", "EKS / orders / node mode", "NodePort"},
		"target-services":        {"EC2 / Web fleet", "Lambda / Invoice", "IP service / Private"},
	} {
		t.Run(file, func(t *testing.T) {
			source, err := os.ReadFile(filepath.Join("..", "..", "docs", "src", "examples", "samples", "aws", albTag, file+".xal"))
			if err != nil {
				t.Fatal(err)
			}
			svg, err := renderer.RenderSVG(context.Background(), source, entity.RenderOptions{})
			if err != nil {
				t.Fatal(err)
			}
			plan, err := renderer.BuildPPTXPlan(context.Background(), source, entity.RenderOptions{PxPerInch: 96})
			if err != nil {
				t.Fatal(err)
			}
			for _, text := range expected {
				if !bytes.Contains(svg, []byte(text)) || !bytes.Contains(plan, []byte(text)) {
					t.Error(fmt.Sprintf("missing visible shared content %q", text))
				}
			}
		})
	}
}
