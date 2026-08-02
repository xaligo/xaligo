package entity

import "crypto/sha256"

// ProjectDocumentKind identifies source families stored in the local project
// index. Diagram calculations continue to use generic ProjectConcept values.
type ProjectDocumentKind string

const (
	ProjectDocumentXAL      ProjectDocumentKind = "xal"
	ProjectDocumentMarkdown ProjectDocumentKind = "markdown"
)

// ProjectConcept is the compact, domain-neutral vocabulary shared by RAG,
// LSP, and MCP project inspection.
type ProjectConcept string

const (
	ProjectConceptFrame   ProjectConcept = "frame"
	ProjectConceptGroup   ProjectConcept = "group"
	ProjectConceptCapture ProjectConcept = "capture"
	ProjectConceptItem    ProjectConcept = "item"
	ProjectConceptPort    ProjectConcept = "port"
	ProjectConceptLine    ProjectConcept = "line"
	ProjectConceptText    ProjectConcept = "text"
	ProjectConceptSpacer  ProjectConcept = "spacer"
)

// ProjectSymbol is one compact semantic row derived from a parsed document.
// ParentOrdinal refers to another symbol in the same Analysis, or -1.
type ProjectSymbol struct {
	Ordinal       int            `json:"ordinal"`
	ParentOrdinal int            `json:"parentOrdinal"`
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Detail        string         `json:"detail,omitempty"`
	Concept       ProjectConcept `json:"concept"`
	SourceTag     string         `json:"sourceTag,omitempty"`
	Source        string         `json:"source,omitempty"`
	Target        string         `json:"target,omitempty"`
	Position      Position       `json:"position"`
}

// ProjectAnalysis is produced once per source revision and reused by editor,
// search, and agent adapters.
type ProjectAnalysis struct {
	URI         string              `json:"uri"`
	Kind        ProjectDocumentKind `json:"kind"`
	Checksum    [sha256.Size]byte   `json:"-"`
	Source      []byte              `json:"-"`
	Symbols     []ProjectSymbol     `json:"symbols"`
	Diagnostics []Diagnostic        `json:"diagnostics,omitempty"`
}

// ProjectSearchResult is a stable FTS result projected from one concept row.
type ProjectSearchResult struct {
	URI       string         `json:"uri"`
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Detail    string         `json:"detail,omitempty"`
	Concept   ProjectConcept `json:"concept"`
	SourceTag string         `json:"sourceTag,omitempty"`
	Line      int            `json:"line"`
	Column    int            `json:"column"`
	Score     float64        `json:"score"`
}

// ProjectIndexStats describes one deterministic incremental indexing pass.
type ProjectIndexStats struct {
	Root        string `json:"root"`
	Scanned     int    `json:"scanned"`
	Indexed     int    `json:"indexed"`
	Unchanged   int    `json:"unchanged"`
	Removed     int    `json:"removed"`
	Diagnostics int    `json:"diagnostics"`
}
