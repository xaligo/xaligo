package entity

// OverflowPolicy controls containment of a box's direct children.
type OverflowPolicy string

const (
	OverflowError   OverflowPolicy = "error"
	OverflowVisible OverflowPolicy = "visible"
)

// Box is a resolved layout node shared by layout and output encoders.
type Box struct {
	ID       string
	Tag      string
	Label    string
	Attrs    map[string]string
	Position Position
	X        float64
	Y        float64
	W        float64
	H        float64

	// ContentX/Y/W/H describe the resolved area in which direct children are
	// allowed to participate in layout. Keeping the content box beside the
	// border box makes the parent/child constraint explicit and lets every
	// renderer consume the same resolved geometry.
	ContentX float64
	ContentY float64
	ContentW float64
	ContentH float64
	// IntrinsicW/H record a source-requested or content-derived natural size
	// when one is known. Zero means that intrinsic measurement is deferred.
	IntrinsicW float64
	IntrinsicH float64
	// Overflow is the non-inherited policy for direct children. Overflowed is
	// set when visible overflow was actually required by resolved geometry.
	Overflow   OverflowPolicy
	Overflowed bool

	Children []*Box

	StaggerDepth int
	IsStaggerBg  bool
	InStagger    bool
}

// Spacing stores resolved edge spacing in pixels.
type Spacing struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}
