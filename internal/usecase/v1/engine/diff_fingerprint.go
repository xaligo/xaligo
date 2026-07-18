package engine

import (
	"fmt"
	"sort"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func diffNodeIdentityV1EngineDiffFingerprint(node *entity.Node) string {
	for _, key := range []string{"name", "ref", "id"} {
		if value := strings.TrimSpace(node.Attrs[key]); value != "" {
			return key + "=" + value
		}
	}
	if node.Tag == "connection" {
		src, dst := strings.TrimSpace(node.Attrs["src"]), strings.TrimSpace(node.Attrs["dst"])
		if src != "" || dst != "" {
			return "src=" + src + ",dst=" + dst
		}
	}
	return ""
}

func diffIdentityKeyV1EngineDiffFingerprint(node *entity.Node, kind string) string {
	if node == nil {
		return ""
	}
	switch kind {
	case "name", "ref", "id":
		value := strings.TrimSpace(node.Attrs[kind])
		if value == "" {
			return ""
		}
		return node.Tag + "|" + kind + "=" + value
	case "connection":
		if node.Tag != "connection" {
			return ""
		}
		src, dst := strings.TrimSpace(node.Attrs["src"]), strings.TrimSpace(node.Attrs["dst"])
		if src == "" || dst == "" {
			return ""
		}
		kindValue := strings.TrimSpace(node.Attrs["kind"])
		if kindValue == "" {
			kindValue = "connection"
		}
		return "connection=" + kindValue + ":" + src + "->" + dst
	default:
		return ""
	}
}

func diffOwnFingerprintV1EngineDiffFingerprint(node *entity.Node, documentRoot bool) string {
	if node == nil {
		return ""
	}
	keys := make([]string, 0, len(node.Attrs))
	for key := range node.Attrs {
		if strings.HasPrefix(key, "_xaligo") {
			continue
		}
		if key == "version" && documentRoot {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var result strings.Builder
	writeDiffFingerprintPartV1EngineDiffFingerprint(&result, node.Tag)
	for _, key := range keys {
		writeDiffFingerprintPartV1EngineDiffFingerprint(&result, key)
		writeDiffFingerprintPartV1EngineDiffFingerprint(&result, canonicalDiffAttributeV1EngineDiffFingerprint(node, key))
	}
	writeDiffFingerprintPartV1EngineDiffFingerprint(&result, comparableDiffNodeTextV1EngineDiffFingerprint(node))
	return result.String()
}

func diffAttributeDistanceV1EngineDiffFingerprint(before, after *entity.Node) int {
	distance := 0
	keys := map[string]struct{}{}
	for key := range before.Attrs {
		if !strings.HasPrefix(key, "_xaligo") {
			keys[key] = struct{}{}
		}
	}
	for key := range after.Attrs {
		if !strings.HasPrefix(key, "_xaligo") {
			keys[key] = struct{}{}
		}
	}
	for key := range keys {
		if canonicalDiffAttributeV1EngineDiffFingerprint(before, key) != canonicalDiffAttributeV1EngineDiffFingerprint(after, key) {
			distance++
		}
	}
	if comparableDiffNodeTextV1EngineDiffFingerprint(before) != comparableDiffNodeTextV1EngineDiffFingerprint(after) {
		distance++
	}
	return distance
}

func canonicalDiffAttributeV1EngineDiffFingerprint(node *entity.Node, key string) string {
	return strings.TrimSpace(node.Attrs[key])
}

func canonicalDiffTextV1EngineDiffFingerprint(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

func comparableDiffNodeTextV1EngineDiffFingerprint(node *entity.Node) string {
	if node == nil || node.Tag == "frame" || node.Tag == "frames" {
		return ""
	}
	return canonicalDiffTextV1EngineDiffFingerprint(node.Text)
}

func diffSubtreeFingerprintV1EngineDiffFingerprint(node *entity.Node) string {
	var result strings.Builder
	writeDiffFingerprintPartV1EngineDiffFingerprint(&result, diffOwnFingerprintV1EngineDiffFingerprint(node, false))
	for _, child := range node.Children {
		writeDiffFingerprintPartV1EngineDiffFingerprint(&result, diffSubtreeFingerprintV1EngineDiffFingerprint(child))
	}
	return result.String()
}

func writeDiffFingerprintPartV1EngineDiffFingerprint(result *strings.Builder, value string) {
	fmt.Fprintf(result, "%d:%s;", len(value), value)
}
