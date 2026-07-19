package engine

import (
	"sort"

	"github.com/xaligo/xaligo/internal/entity"
)

func (state *diffDocumentStateV1EngineDiffDocument) classifyV1EngineDiffClassify() entity.StructuralDiff {
	result := entity.StructuralDiff{}
	for _, descriptor := range state.beforeDescriptors {
		if state.beforeToAfter[descriptor.node] != nil || (descriptor.parent != nil && state.beforeToAfter[descriptor.parent] == nil) {
			continue
		}
		result.Before = append(result.Before, structuralChangeV1EngineDiffClassify(entity.StructuralChangeRemoved, descriptor))
		result.RemovedCount++
	}
	for _, descriptor := range state.afterDescriptors {
		if state.afterToBefore[descriptor.node] != nil || (descriptor.parent != nil && state.afterToBefore[descriptor.parent] == nil) {
			continue
		}
		result.After = append(result.After, structuralChangeV1EngineDiffClassify(entity.StructuralChangeAdded, descriptor))
		result.AddedCount++
	}
	for _, beforeDescriptor := range state.beforeDescriptors {
		afterNode := state.beforeToAfter[beforeDescriptor.node]
		if afterNode == nil {
			continue
		}
		afterDescriptor := state.afterByNode[afterNode]
		beforeIsRoot := beforeDescriptor.parent == nil
		afterIsRoot := afterDescriptor.parent == nil
		if diffOwnFingerprintV1EngineDiffFingerprint(beforeDescriptor.node, beforeIsRoot) == diffOwnFingerprintV1EngineDiffFingerprint(afterNode, afterIsRoot) && !state.nodeMovedV1EngineDiffClassify(beforeDescriptor, afterDescriptor) {
			continue
		}
		result.Before = append(result.Before, structuralChangeV1EngineDiffClassify(entity.StructuralChangeModified, beforeDescriptor))
		result.After = append(result.After, structuralChangeV1EngineDiffClassify(entity.StructuralChangeModified, afterDescriptor))
		result.ModifiedCount++
	}
	sortStructuralChangesV1EngineDiffClassify(result.Before)
	sortStructuralChangesV1EngineDiffClassify(result.After)
	return result
}

func (state *diffDocumentStateV1EngineDiffDocument) nodeMovedV1EngineDiffClassify(before, after *diffNodeDescriptorV1EngineDiffDocument) bool {
	if before.parent == nil || after.parent == nil {
		return before.parent != after.parent
	}
	if state.beforeToAfter[before.parent] != after.parent {
		return true
	}
	beforeRank, afterRank := 0, 0
	for _, sibling := range before.parent.Children {
		counterpart := state.beforeToAfter[sibling]
		if counterpart == nil || state.afterByNode[counterpart].parent != after.parent {
			continue
		}
		if sibling == before.node {
			break
		}
		beforeRank++
	}
	for _, sibling := range after.parent.Children {
		counterpart := state.afterToBefore[sibling]
		if counterpart == nil || state.beforeByNode[counterpart].parent != before.parent {
			continue
		}
		if sibling == after.node {
			break
		}
		afterRank++
	}
	return beforeRank != afterRank
}

func structuralChangeV1EngineDiffClassify(kind entity.StructuralChangeKind, descriptor *diffNodeDescriptorV1EngineDiffDocument) entity.StructuralChange {
	return entity.StructuralChange{
		Kind: kind, Path: descriptor.path, Tag: descriptor.node.Tag,
		Identity: diffNodeIdentityV1EngineDiffFingerprint(descriptor.node), Position: descriptor.node.Position,
	}
}

func sortStructuralChangesV1EngineDiffClassify(changes []entity.StructuralChange) {
	sort.SliceStable(changes, func(i, j int) bool {
		return changes[i].Position.Offset < changes[j].Position.Offset
	})
}
