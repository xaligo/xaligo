package engine

// connectorStyle is the resolved arrowhead + width for connectors.
type connectorStyleV1EnginePlanConnectorStyle struct {
	Head     string
	Width    float64
	HasWidth bool
}

func resolveConnectorStyleV1EnginePlanConnectorStyle(style string) connectorStyleV1EnginePlanConnectorStyle {
	switch style {
	case "standard":
		return connectorStyleV1EnginePlanConnectorStyle{Head: "triangle", Width: 1.5, HasWidth: true}
	case "triangle":
		return connectorStyleV1EnginePlanConnectorStyle{Head: "triangle"}
	case "stealth":
		return connectorStyleV1EnginePlanConnectorStyle{Head: "stealth"}
	case "arrow":
		return connectorStyleV1EnginePlanConnectorStyle{Head: "arrow"}
	case "diamond":
		return connectorStyleV1EnginePlanConnectorStyle{Head: "diamond"}
	case "oval":
		return connectorStyleV1EnginePlanConnectorStyle{Head: "oval"}
	case "none":
		return connectorStyleV1EnginePlanConnectorStyle{Head: "none"}
	default: // "thin" and unset → slender stealth head on a thin line
		return connectorStyleV1EnginePlanConnectorStyle{Head: "stealth", Width: 1, HasWidth: true}
	}
}
