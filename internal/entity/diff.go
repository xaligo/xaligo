package entity

// StructuralChangeKind classifies one semantic tree change.
type StructuralChangeKind string

const (
	StructuralChangeAdded    StructuralChangeKind = "added"
	StructuralChangeRemoved  StructuralChangeKind = "removed"
	StructuralChangeModified StructuralChangeKind = "modified"
)

// StructuralChange identifies one changed branch on one side of a document
// comparison. Modified nodes occur in both Before and After.
type StructuralChange struct {
	Kind     StructuralChangeKind `json:"kind"`
	Path     string               `json:"path"`
	Tag      string               `json:"tag"`
	Identity string               `json:"identity,omitempty"`
	Position Position             `json:"position"`
}

// StructuralDiff contains side-specific targets and a compact change summary.
// Added or removed subtrees are represented by their highest changed root.
type StructuralDiff struct {
	Before        []StructuralChange `json:"before"`
	After         []StructuralChange `json:"after"`
	AddedCount    int                `json:"addedCount"`
	RemovedCount  int                `json:"removedCount"`
	ModifiedCount int                `json:"modifiedCount"`
}

// DiffResult contains the two SVG projections of one structural comparison.
type DiffResult struct {
	RemovedImage []byte
	AddedImage   []byte
	Summary      StructuralDiff
}
