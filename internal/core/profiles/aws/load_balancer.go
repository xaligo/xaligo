package aws

import "fmt"

// These are annotation consistency checks, not a cloud deployment validator.
// Missing optional fields remain allowed in a partial architecture diagram.
func (d Definition) validateLoadBalancerAnnotations(attrs map[string]string) error {
	alb := d.Tag == "aws-elastic-load-balancing-application-load-balancer"
	nlb := d.Tag == "aws-elastic-load-balancing-network-load-balancer"
	if !alb && !nlb {
		return nil
	}
	invalid := func(message string) error { return fmt.Errorf("<%s> %s", d.Tag, message) }
	if nlb {
		for _, key := range []string{"trust-store", "mutual-tls-mode"} {
			if _, exists := attrs[key]; exists {
				return invalid("does not support " + key + "; backend mutual TLS uses a TCP passthrough listener")
			}
		}
	}
	protocol := attrs["listener-protocol"]
	secureProtocol := "TLS"
	if alb {
		secureProtocol = "HTTPS"
	}
	for _, key := range []string{"certificate", "tls-policy", "alpn-policy", "trust-store"} {
		if _, exists := attrs[key]; exists && protocol != "" && protocol != secureProtocol {
			return invalid(key + " requires a " + secureProtocol + " listener, not " + protocol)
		}
	}
	if alb {
		if _, exists := attrs["alpn-policy"]; exists {
			return invalid("alpn-policy is an NLB TLS listener annotation")
		}
		if mode := attrs["mutual-tls-mode"]; mode != "" && mode != "off" && protocol != "" && protocol != "HTTPS" {
			return invalid("mutual-tls-mode requires an HTTPS listener")
		}
		if _, exists := attrs["trust-store"]; exists && attrs["mutual-tls-mode"] != "" && attrs["mutual-tls-mode"] != "verify" {
			return invalid("trust-store is used with mutual-tls-mode=verify")
		}
		if attrs["target-type"] == "lambda" {
			for _, key := range []string{"target-port", "target-protocol"} {
				if _, exists := attrs[key]; exists {
					return invalid(key + " does not apply to Lambda targets")
				}
			}
		}
	}
	return nil
}
