package v2

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func frontendAWSLoadBalancer(tag string) string {
	switch tag {
	case "aws-elastic-load-balancing-application-load-balancer":
		return "alb"
	case "aws-elastic-load-balancing-network-load-balancer":
		return "nlb"
	}
	return ""
}

func frontendAWSComponent(node *frontendNode) bool {
	return node != nil && (node.tag == "aws-listener" || frontendAWSFeatureKind(node.tag) != "" || frontendAWSLoadBalancer(node.tag) != "" && (node.attrs["view"] == "component" || node.attrs["domain"] != "" || len(node.children) > 0))
}

func validateFrontendAWSComponent(node *frontendNode) error {
	if id := strings.TrimSpace(node.attrs["id"]); id == "" || strings.ContainsAny(id, " \t\r\n") {
		return fmt.Errorf("<%s> requires a non-empty whitespace-free id", node.tag)
	}
	if strings.TrimSpace(node.text) != "" {
		return fmt.Errorf("<%s> does not accept text content", node.tag)
	}
	allowed := " id visible detail-level show hide "
	if kind := frontendAWSFeatureKind(node.tag); kind != "" {
		return validateFrontendAWSFeature(node, allowed)
	}
	if node.tag == "aws-listener" {
		allowed += "protocol port mtls certificate trust-store target-group backend-tls backend-mtls show-title "
		for _, child := range node.children {
			if child.tag != "aws-listener-rule" && child.tag != "aws-option" {
				return fmt.Errorf("<aws-listener> accepts aws-listener-rule and aws-option children")
			}
		}
		switch node.attrs["protocol"] {
		case "HTTP", "HTTPS", "TCP", "TLS", "UDP", "TCP_UDP", "QUIC", "TCP_QUIC":
		default:
			return fmt.Errorf("<aws-listener> requires a supported protocol")
		}
		port, err := frontendOptionalUint16(node, "port")
		if err != nil || port == nil || *port == 0 {
			return fmt.Errorf("<aws-listener> port must be an integer in 1..65535")
		}
		switch node.attrs["mtls"] {
		case "", "off", "verify", "passthrough":
		default:
			return fmt.Errorf("<aws-listener> mtls must be off, verify, or passthrough")
		}
	} else {
		allowed += "domain view width height x y dx dy margin margin-top margin-right margin-bottom margin-left class fill stroke stroke-width corner-radius opacity layer "
		if view := node.attrs["view"]; view != "" && view != "component" {
			return fmt.Errorf("<%s> with listeners/domain requires view=component", node.tag)
		}
		listeners := 0
		for _, child := range node.children {
			if child.tag == "aws-listener" {
				listeners++
			}
		}
		if listeners < 1 || listeners > 32 {
			return fmt.Errorf("<%s> component requires 1..32 aws-listener children", node.tag)
		}
		for _, child := range node.children {
			if child.tag != "aws-listener" && child.tag != "aws-target-group" && child.tag != "aws-option" {
				return fmt.Errorf("<%s> accepts aws-listener, aws-target-group and aws-option children", node.tag)
			}
		}
	}
	for name := range node.attrs {
		if !strings.Contains(allowed, " "+name+" ") {
			return fmt.Errorf("<%s> component does not support %s", node.tag, name)
		}
	}
	return nil
}

func applyFrontendAWSComponent(node *frontendNode, element *entity.EngineElementSpec) error {
	model := &entity.EngineAWSComponentSpec{Kind: frontendAWSLoadBalancer(node.tag), Domain: strings.TrimSpace(node.attrs["domain"])}
	model.DetailLevel, model.Show, model.Hide = node.attrs["detail-level"], node.attrs["show"], node.attrs["hide"]
	element.Concept = entity.EngineConceptGroup
	element.Layout = entity.EngineLayoutAbsolute
	element.Text = nil
	element.Padding = entity.EngineInsets{}
	if frontendAWSFeatureKind(node.tag) != "" {
		return applyFrontendAWSFeature(node, element, model)
	}
	if node.tag == "aws-listener" {
		model.Kind, model.Protocol, model.MutualTLS = "listener", node.attrs["protocol"], node.attrs["mtls"]
		model.Certificate, model.TrustStore, model.TargetGroup = node.attrs["certificate"], node.attrs["trust-store"], node.attrs["target-group"]
		var err error
		if model.ShowTitle, err = frontendOptionalBool(node, "show-title"); err != nil {
			return err
		}
		if model.Port, err = frontendOptionalUint16(node, "port"); err != nil {
			return err
		}
		if model.BackendTLS, err = frontendOptionalBool(node, "backend-tls"); err != nil {
			return err
		}
		if model.BackendMTLS, err = frontendOptionalBool(node, "backend-mtls"); err != nil {
			return err
		}
	}
	element.AWS = model
	return nil
}
