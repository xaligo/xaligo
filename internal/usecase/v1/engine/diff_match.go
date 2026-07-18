package engine

import (
	"sort"

	"github.com/xaligo/xaligo/internal/entity"
)

func (state *diffDocumentStateV1EngineDiffDocument) matchGlobalIdentitiesV1EngineDiffMatch() {
	for _, identityKind := range []string{"name", "ref", "id", "connection"} {
		before := uniqueDiffIdentityNodesV1EngineDiffMatch(state.beforeDescriptors, state.beforeToAfter, identityKind)
		after := uniqueDiffIdentityNodesV1EngineDiffMatch(state.afterDescriptors, state.afterToBefore, identityKind)
		for key, beforeNode := range before {
			if afterNode := after[key]; afterNode != nil {
				state.addMatchV1EngineDiffDocument(beforeNode, afterNode)
			}
		}
	}
}

func uniqueDiffIdentityNodesV1EngineDiffMatch(descriptors []*diffNodeDescriptorV1EngineDiffDocument, matches map[*entity.Node]*entity.Node, kind string) map[string]*entity.Node {
	counts := map[string]int{}
	nodes := map[string]*entity.Node{}
	for _, descriptor := range descriptors {
		if matches[descriptor.node] != nil {
			continue
		}
		key := diffIdentityKeyV1EngineDiffFingerprint(descriptor.node, kind)
		if key == "" {
			continue
		}
		counts[key]++
		nodes[key] = descriptor.node
	}
	for key, count := range counts {
		if count != 1 {
			delete(nodes, key)
		}
	}
	return nodes
}

func (state *diffDocumentStateV1EngineDiffDocument) matchDescendantsV1EngineDiffMatch() {
	queue := make([]diffNodePairV1EngineDiffDocument, 0, len(state.beforeToAfter))
	for before, after := range state.beforeToAfter {
		queue = append(queue, diffNodePairV1EngineDiffDocument{before: before, after: after})
	}
	sort.Slice(queue, func(i, j int) bool {
		return state.beforeByNode[queue[i].before].depth < state.beforeByNode[queue[j].before].depth
	})
	seen := map[*entity.Node]bool{}
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		if seen[pair.before] {
			continue
		}
		seen[pair.before] = true
		for _, childPair := range state.matchChildrenV1EngineDiffMatch(pair.before.Children, pair.after.Children) {
			if state.addMatchV1EngineDiffDocument(childPair.before, childPair.after) {
				queue = append(queue, childPair)
			}
		}
	}
}

func (state *diffDocumentStateV1EngineDiffDocument) matchChildrenV1EngineDiffMatch(before, after []*entity.Node) []diffNodePairV1EngineDiffDocument {
	pairs := []diffNodePairV1EngineDiffDocument{}
	remainingBefore := unmatchedDiffChildrenV1EngineDiffMatch(before, state.beforeToAfter)
	remainingAfter := unmatchedDiffChildrenV1EngineDiffMatch(after, state.afterToBefore)
	for _, identityKind := range []string{"name", "ref", "id", "connection"} {
		matched, nextBefore, nextAfter := matchLocalIdentitiesV1EngineDiffMatch(remainingBefore, remainingAfter, identityKind)
		pairs = append(pairs, matched...)
		remainingBefore, remainingAfter = nextBefore, nextAfter
	}
	matched, remainingBefore, remainingAfter := matchExactSubtreesV1EngineDiffMatch(remainingBefore, remainingAfter)
	pairs = append(pairs, matched...)
	return append(pairs, alignDiffChildrenV1EngineDiffMatch(remainingBefore, remainingAfter)...)
}

func unmatchedDiffChildrenV1EngineDiffMatch(children []*entity.Node, matches map[*entity.Node]*entity.Node) []*entity.Node {
	result := make([]*entity.Node, 0, len(children))
	for _, child := range children {
		if matches[child] == nil {
			result = append(result, child)
		}
	}
	return result
}

func matchLocalIdentitiesV1EngineDiffMatch(before, after []*entity.Node, kind string) ([]diffNodePairV1EngineDiffDocument, []*entity.Node, []*entity.Node) {
	beforeUnique := uniqueDiffNodesByKindV1EngineDiffMatch(before, kind)
	afterUnique := uniqueDiffNodesByKindV1EngineDiffMatch(after, kind)
	usedBefore := map[*entity.Node]bool{}
	usedAfter := map[*entity.Node]bool{}
	pairs := []diffNodePairV1EngineDiffDocument{}
	for key, beforeNode := range beforeUnique {
		if afterNode := afterUnique[key]; afterNode != nil {
			pairs = append(pairs, diffNodePairV1EngineDiffDocument{before: beforeNode, after: afterNode})
			usedBefore[beforeNode] = true
			usedAfter[afterNode] = true
		}
	}
	return pairs, filterDiffNodesV1EngineDiffMatch(before, usedBefore), filterDiffNodesV1EngineDiffMatch(after, usedAfter)
}

func uniqueDiffNodesByKindV1EngineDiffMatch(nodes []*entity.Node, kind string) map[string]*entity.Node {
	counts := map[string]int{}
	result := map[string]*entity.Node{}
	for _, node := range nodes {
		key := diffIdentityKeyV1EngineDiffFingerprint(node, kind)
		if key == "" {
			continue
		}
		counts[key]++
		result[key] = node
	}
	for key, count := range counts {
		if count != 1 {
			delete(result, key)
		}
	}
	return result
}

func filterDiffNodesV1EngineDiffMatch(nodes []*entity.Node, used map[*entity.Node]bool) []*entity.Node {
	result := make([]*entity.Node, 0, len(nodes))
	for _, node := range nodes {
		if !used[node] {
			result = append(result, node)
		}
	}
	return result
}

func matchExactSubtreesV1EngineDiffMatch(before, after []*entity.Node) ([]diffNodePairV1EngineDiffDocument, []*entity.Node, []*entity.Node) {
	afterByFingerprint := map[string][]*entity.Node{}
	for _, node := range after {
		fingerprint := diffSubtreeFingerprintV1EngineDiffFingerprint(node)
		afterByFingerprint[fingerprint] = append(afterByFingerprint[fingerprint], node)
	}
	usedBefore := map[*entity.Node]bool{}
	usedAfter := map[*entity.Node]bool{}
	pairs := []diffNodePairV1EngineDiffDocument{}
	for _, beforeNode := range before {
		fingerprint := diffSubtreeFingerprintV1EngineDiffFingerprint(beforeNode)
		candidates := afterByFingerprint[fingerprint]
		if len(candidates) == 0 {
			continue
		}
		afterNode := candidates[0]
		afterByFingerprint[fingerprint] = candidates[1:]
		pairs = append(pairs, diffNodePairV1EngineDiffDocument{before: beforeNode, after: afterNode})
		usedBefore[beforeNode] = true
		usedAfter[afterNode] = true
	}
	return pairs, filterDiffNodesV1EngineDiffMatch(before, usedBefore), filterDiffNodesV1EngineDiffMatch(after, usedAfter)
}

func alignDiffChildrenV1EngineDiffMatch(before, after []*entity.Node) []diffNodePairV1EngineDiffDocument {
	const gapCost = 100
	rows, cols := len(before)+1, len(after)+1
	dp := make([][]int, rows)
	action := make([][]byte, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
		action[i] = make([]byte, cols)
	}
	for i := 1; i < rows; i++ {
		dp[i][0], action[i][0] = i*gapCost, 'd'
	}
	for j := 1; j < cols; j++ {
		dp[0][j], action[0][j] = j*gapCost, 'a'
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			deleteCost := dp[i-1][j] + gapCost
			addCost := dp[i][j-1] + gapCost
			matchCost := dp[i-1][j-1] + diffNodeMatchCostV1EngineDiffMatch(before[i-1], after[j-1])
			dp[i][j], action[i][j] = deleteCost, 'd'
			if addCost < dp[i][j] {
				dp[i][j], action[i][j] = addCost, 'a'
			}
			if matchCost <= dp[i][j] {
				dp[i][j], action[i][j] = matchCost, 'm'
			}
		}
	}
	pairs := []diffNodePairV1EngineDiffDocument{}
	for i, j := len(before), len(after); i > 0 || j > 0; {
		switch action[i][j] {
		case 'm':
			if before[i-1].Tag == after[j-1].Tag {
				pairs = append(pairs, diffNodePairV1EngineDiffDocument{before: before[i-1], after: after[j-1]})
			}
			i, j = i-1, j-1
		case 'd':
			i--
		default:
			j--
		}
	}
	return pairs
}

func diffNodeMatchCostV1EngineDiffMatch(before, after *entity.Node) int {
	if before.Tag != after.Tag {
		return 250
	}
	beforeIdentity := diffNodeIdentityV1EngineDiffFingerprint(before)
	afterIdentity := diffNodeIdentityV1EngineDiffFingerprint(after)
	if beforeIdentity != afterIdentity && (beforeIdentity != "" || afterIdentity != "") {
		return 250
	}
	if diffOwnFingerprintV1EngineDiffFingerprint(before, false) == diffOwnFingerprintV1EngineDiffFingerprint(after, false) {
		return 10
	}
	cost := 60 + 15*diffAttributeDistanceV1EngineDiffFingerprint(before, after)
	if cost > 180 {
		return 180
	}
	return cost
}
