package engine

import (
	"fmt"

	"github.com/xaligo/xaligo/internal/entity"
)

const diffStatusAttrV1EngineDiffDocument = "_xaligoDiffStatus"

type diffNodeDescriptorV1EngineDiffDocument struct {
	node   *entity.Node
	parent *entity.Node
	index  int
	depth  int
	path   string
}

type diffNodePairV1EngineDiffDocument struct {
	before *entity.Node
	after  *entity.Node
}

type diffDocumentStateV1EngineDiffDocument struct {
	beforeDescriptors []*diffNodeDescriptorV1EngineDiffDocument
	afterDescriptors  []*diffNodeDescriptorV1EngineDiffDocument
	beforeByNode      map[*entity.Node]*diffNodeDescriptorV1EngineDiffDocument
	afterByNode       map[*entity.Node]*diffNodeDescriptorV1EngineDiffDocument
	beforeToAfter     map[*entity.Node]*entity.Node
	afterToBefore     map[*entity.Node]*entity.Node
}

// DiffDocumentsV1EngineDiffDocument compares parsed V1 document trees rather
// than source bytes or generated scene IDs.
func DiffDocumentsV1EngineDiffDocument(before, after entity.Document) entity.StructuralDiff {
	state := newDiffDocumentStateV1EngineDiffDocument(before.Root, after.Root)
	if before.Root == nil || after.Root == nil || before.Root.Tag != after.Root.Tag {
		return state.classifyV1EngineDiffClassify()
	}
	state.addMatchV1EngineDiffDocument(before.Root, after.Root)
	state.matchGlobalIdentitiesV1EngineDiffMatch()
	state.matchDescendantsV1EngineDiffMatch()
	return state.classifyV1EngineDiffClassify()
}

// MarkChangesV1EngineDiffDocument annotates the smallest renderable branches
// for one side. Connection descendants are promoted to their owning connector.
func MarkChangesV1EngineDiffDocument(root *entity.Node, changes []entity.StructuralChange, status string) {
	if root == nil || len(changes) == 0 {
		return
	}
	selected := make(map[int]struct{}, len(changes))
	for _, change := range changes {
		selected[change.Position.Offset] = struct{}{}
	}
	var markConnections func(*entity.Node)
	markConnections = func(node *entity.Node) {
		if node == nil {
			return
		}
		if node.Tag == "connection" {
			setDiffStatusV1EngineDiffDocument(node, status)
		}
		for _, child := range node.Children {
			markConnections(child)
		}
	}
	var visit func(*entity.Node, *entity.Node, bool)
	visit = func(node, connection *entity.Node, covered bool) {
		if node == nil {
			return
		}
		if node.Tag == "connection" {
			connection = node
			if covered {
				setDiffStatusV1EngineDiffDocument(connection, status)
			}
		}
		_, changed := selected[node.Position.Offset]
		if changed {
			switch {
			case connection != nil:
				setDiffStatusV1EngineDiffDocument(connection, status)
			case node.Tag == "connections":
				markConnections(node)
			case !covered:
				setDiffStatusV1EngineDiffDocument(node, status)
				covered = true
			}
		}
		for _, child := range node.Children {
			visit(child, connection, covered)
		}
	}
	visit(root, nil, false)
}

func setDiffStatusV1EngineDiffDocument(node *entity.Node, status string) {
	if node.Attrs == nil {
		node.Attrs = map[string]string{}
	}
	node.Attrs[diffStatusAttrV1EngineDiffDocument] = status
}

func newDiffDocumentStateV1EngineDiffDocument(before, after *entity.Node) *diffDocumentStateV1EngineDiffDocument {
	state := &diffDocumentStateV1EngineDiffDocument{
		beforeByNode:  map[*entity.Node]*diffNodeDescriptorV1EngineDiffDocument{},
		afterByNode:   map[*entity.Node]*diffNodeDescriptorV1EngineDiffDocument{},
		beforeToAfter: map[*entity.Node]*entity.Node{},
		afterToBefore: map[*entity.Node]*entity.Node{},
	}
	state.beforeDescriptors = flattenDiffNodesV1EngineDiffDocument(before, state.beforeByNode)
	state.afterDescriptors = flattenDiffNodesV1EngineDiffDocument(after, state.afterByNode)
	return state
}

func flattenDiffNodesV1EngineDiffDocument(root *entity.Node, byNode map[*entity.Node]*diffNodeDescriptorV1EngineDiffDocument) []*diffNodeDescriptorV1EngineDiffDocument {
	result := []*diffNodeDescriptorV1EngineDiffDocument{}
	var visit func(*entity.Node, *entity.Node, int, int, string)
	visit = func(node, parent *entity.Node, index, depth int, parentPath string) {
		if node == nil {
			return
		}
		segment := fmt.Sprintf("%s[%d]", node.Tag, index+1)
		if identity := diffNodeIdentityV1EngineDiffFingerprint(node); identity != "" {
			segment += "{" + identity + "}"
		}
		path := parentPath + "/" + segment
		descriptor := &diffNodeDescriptorV1EngineDiffDocument{node: node, parent: parent, index: index, depth: depth, path: path}
		result = append(result, descriptor)
		byNode[node] = descriptor
		for childIndex, child := range node.Children {
			visit(child, node, childIndex, depth+1, path)
		}
	}
	visit(root, nil, 0, 0, "")
	return result
}

func (state *diffDocumentStateV1EngineDiffDocument) addMatchV1EngineDiffDocument(before, after *entity.Node) bool {
	if before == nil || after == nil || state.beforeToAfter[before] != nil || state.afterToBefore[after] != nil {
		return false
	}
	state.beforeToAfter[before] = after
	state.afterToBefore[after] = before
	return true
}
