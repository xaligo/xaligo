package entity

import "crypto/sha256"

// ProjectDocumentKind identifies source families stored in the local project
// index. Diagram calculations continue to use generic ProjectConcept values.
type ProjectDocumentKind string

const (
	ProjectDocumentXAL      ProjectDocumentKind = "xal"
	ProjectDocumentMarkdown ProjectDocumentKind = "markdown"
)

// ProjectConcept reuses the engine's closed, domain-neutral vocabulary for
// RAG, LSP, and MCP inspection. The alias avoids a second concept mapping when
// a parsed .xal tree is later lowered into EngineDocumentSpec.
type ProjectConcept = EngineConcept

const (
	ProjectConceptFrame   = EngineConceptFrame
	ProjectConceptGroup   = EngineConceptGroup
	ProjectConceptCapture = EngineConceptCapture
	ProjectConceptItem    = EngineConceptItem
	ProjectConceptPort    = EngineConceptPort
	ProjectConceptLine    = EngineConceptLine
	ProjectConceptText    = EngineConceptText
	ProjectConceptSpacer  = EngineConceptSpacer
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
