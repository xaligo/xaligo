package share

import "strings"

func ItemNodeID(bindingID string) string {
	return strings.TrimSuffix(bindingID, "-lbl")
}

func PositiveWidth(value float64) float64 {
	if value > 0 {
		return value
	}
	return 1
}
