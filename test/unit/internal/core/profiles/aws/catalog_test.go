package aws_test

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"testing"

	awsassets "github.com/xaligo/xaligo/etc/resources/aws"
	awsprofile "github.com/xaligo/xaligo/internal/core/profiles/aws"
)

func TestDedicatedTagsCoverEveryBundledAWSIcon(t *testing.T) {
	file, err := awsassets.Assets.Open(awsassets.CatalogCSV)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		t.Fatal(err)
	}
	covered, tags := map[int]string{}, map[string]bool{}
	definitions := awsprofile.Definitions()
	for _, definition := range definitions {
		if tags[definition.Tag] {
			t.Fatalf("duplicate tag %s", definition.Tag)
		}
		tags[definition.Tag] = true
		if definition.Name == "" || definition.Scope == "" || definition.Description == "" {
			t.Fatalf("incomplete profile: %#v", definition)
		}
		for _, id := range definition.CatalogIDs {
			if previous := covered[id]; previous != "" {
				t.Fatalf("catalog %d covered by both %s and %s", id, previous, definition.Tag)
			}
			covered[id] = definition.Tag
		}
		attrs := map[string]string{}
		for _, parameter := range definition.Parameters {
			attrs[parameter.Name] = parameter.Example
		}
		if err := definition.ValidateAnnotations(attrs); err != nil {
			t.Fatalf("invalid example for %s: %v", definition.Tag, err)
		}
	}
	count := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(row[4], "/Tabler-Icons/") || strings.Contains(row[4], "/Yamaha-Network-Icons/") {
			continue
		}
		id, _ := strconv.Atoi(row[0])
		count++
		if covered[id] == "" {
			t.Errorf("missing AWS catalog %d: %s", id, row[2])
		}
		delete(covered, id)
	}
	if len(covered) != 0 {
		t.Errorf("non-AWS catalog entries referenced: %v", covered)
	}
	if count != 1875 || len(definitions) != 877 {
		t.Fatalf("unexpected coverage: %d entries / %d tags", count, len(definitions))
	}
}

func TestAWSProfilesAreIndependentValues(t *testing.T) {
	definition, _ := awsprofile.DefinitionForTag("aws-ec2")
	definition.CatalogIDs[0] = -1
	definition.Parameters[1].Values[0] = "broken"
	fresh, _ := awsprofile.DefinitionForTag("aws-ec2")
	if fresh.CatalogIDs[0] < 0 || fresh.Parameters[1].Values[0] != "running" {
		t.Fatal("caller mutated registry")
	}
	group, _ := awsprofile.DefinitionForTag("vpc")
	group.Group.Stroke = "broken"
	fresh, _ = awsprofile.DefinitionForTag("vpc")
	if fresh.Group.Stroke == "broken" {
		t.Fatal("caller mutated group registry")
	}
}

func TestAWSParameterValidationAndLabels(t *testing.T) {
	for _, test := range []struct{ tag, attr, value string }{
		{"aws-ec2", "state", "flying"}, {"aws-ec2", "size", "NaN"}, {"aws-ec2", "label-width", "-1"},
		{"aws-lambda", "memory-mb", "1.5"}, {"aws-rds", "multi-az", "maybe"}, {"vpc", "cidr", "10.0.0.0/99"},
		{"vpc-endpoint", "endpoint-type", "nat"}, {"aws-s3-bucket", "show-details", "yes"},
	} {
		d, ok := awsprofile.DefinitionForTag(test.tag)
		if !ok {
			t.Fatal(test.tag)
		}
		if d.ValidateAnnotations(map[string]string{test.attr: test.value}) == nil {
			t.Errorf("accepted %s %s=%s", test.tag, test.attr, test.value)
		}
	}
	d, _ := awsprofile.DefinitionForTag("aws-ec2")
	attrs := map[string]string{"label": "API / 基盤", "instance-type": "t3.micro"}
	if got := d.Label(attrs); got != "API / 基盤\ninstance-type: t3.micro" {
		t.Fatal(got)
	}
	attrs["show-details"] = "false"
	if got := d.Label(attrs); got != "API / 基盤" {
		t.Fatal(got)
	}
}

func TestLoadBalancerAnnotationCapabilities(t *testing.T) {
	alb, _ := awsprofile.DefinitionForTag("aws-elastic-load-balancing-application-load-balancer")
	nlb, _ := awsprofile.DefinitionForTag("aws-elastic-load-balancing-network-load-balancer")
	for _, d := range []awsprofile.Definition{alb, nlb} {
		for _, port := range []string{"0", "65536", "-1", "443.5", "NaN"} {
			if d.ValidateAnnotations(map[string]string{"listener-port": port}) == nil {
				t.Errorf("%s accepted port %s", d.Tag, port)
			}
		}
		for _, port := range []string{"1", "443", "65535"} {
			if err := d.ValidateAnnotations(map[string]string{"listener-port": port}); err != nil {
				t.Error(err)
			}
		}
	}
	for _, attrs := range []map[string]string{
		{"trust-store": "ca"}, {"mutual-tls-mode": "verify"},
		{"listener-protocol": "HTTPS"}, {"listener-protocol": "TCP", "certificate": "server"},
	} {
		if nlb.ValidateAnnotations(attrs) == nil {
			t.Errorf("NLB accepted unsupported annotations: %v", attrs)
		}
	}
	for _, attrs := range []map[string]string{
		{"listener-protocol": "TLS"}, {"listener-protocol": "HTTP", "mutual-tls-mode": "verify"},
		{"mutual-tls-mode": "passthrough", "trust-store": "ca"}, {"target-type": "lambda", "target-port": "443"},
		{"alpn-policy": "HTTP2Preferred"},
	} {
		if alb.ValidateAnnotations(attrs) == nil {
			t.Errorf("ALB accepted unsupported annotations: %v", attrs)
		}
	}
	if err := alb.ValidateAnnotations(map[string]string{"listener-protocol": "HTTPS", "listener-port": "443", "mutual-tls-mode": "verify", "trust-store": "ca"}); err != nil {
		t.Fatal(err)
	}
	for _, protocol := range []string{"TCP", "TLS", "UDP", "TCP_UDP", "QUIC", "TCP_QUIC"} {
		if err := nlb.ValidateAnnotations(map[string]string{"listener-protocol": protocol}); err != nil {
			t.Fatal(err)
		}
	}
}
