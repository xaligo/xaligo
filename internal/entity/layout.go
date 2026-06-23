package entity

// Box is a resolved layout node shared by layout and output encoders.
type Box struct {
	ID       string
	Tag      string
	Label    string
	Attrs    map[string]string
	X        float64
	Y        float64
	W        float64
	H        float64
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
