package aws

// NativeChildKind is the shared authoring vocabulary, separate from icon tags.
func NativeChildKind(tag string) string {
	switch tag {
	case "aws-listener-rule":
		return "rule"
	case "aws-rule-condition":
		return "condition"
	case "aws-rule-match":
		return "match"
	case "aws-rule-action":
		return "action"
	case "aws-forward-target":
		return "forward-target"
	case "aws-jwt-claim":
		return "jwt-claim"
	case "aws-rule-transform":
		return "transform"
	case "aws-rule-rewrite":
		return "rewrite"
	case "aws-target-group":
		return "target-group"
	case "aws-target-service":
		return "target-service"
	case "aws-registered-target":
		return "target"
	case "aws-option":
		return "option"
	}
	return ""
}

// TargetServiceIconTag maps logical workload kinds to existing official assets.
func TargetServiceIconTag(kind string) string {
	switch kind {
	case "ecs":
		return "aws-elastic-container-service"
	case "eks":
		return "aws-elastic-kubernetes-service"
	case "ec2":
		return "aws-ec2"
	case "lambda":
		return "aws-lambda"
	}
	return ""
}
