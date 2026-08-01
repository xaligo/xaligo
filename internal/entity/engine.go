package entity

// EngineDirection selects the main axis used by the generic V2 engine.
type EngineDirection string

const (
	EngineDirectionVertical   EngineDirection = "vertical"
	EngineDirectionHorizontal EngineDirection = "horizontal"
)

// EngineElementSpec is one domain-neutral element passed to the Rust engine.
// Pointer fields preserve the distinction between an unset value and zero.
type EngineElementSpec struct {
	ID     string
	Width  *float64
	Height *float64
	Weight *float64
}

// EngineDocumentSpec is the typed Go-side request for generic allocation.
type EngineDocumentSpec struct {
	Direction EngineDirection
	Width     float64
	Height    float64
	Gap       float64
	Elements  []EngineElementSpec
}

// EngineResolvedElement contains geometry calculated by the Rust engine.
type EngineResolvedElement struct {
	ID     string
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// EngineResolvedDocument is the immutable resolved result consumed by output
// projections. Width and Height retain the requested document dimensions.
type EngineResolvedDocument struct {
	Width    float64
	Height   float64
	Elements []EngineResolvedElement
}
