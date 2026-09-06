package entity

// EngineDirection is the legacy main-axis selector retained for callers that
// construct a flat EngineDocumentSpec. New V2 callers should use Layout.
type EngineDirection string

const (
	EngineDirectionVertical   EngineDirection = "vertical"
	EngineDirectionHorizontal EngineDirection = "horizontal"
)

// EngineConcept is the closed, domain-neutral calculation vocabulary accepted
// by the V2 layout/router. Native AWS components expand to these concepts
// before calculation; other profiles lower directly.
type EngineConcept string

const (
	EngineConceptFrame   EngineConcept = "frame"
	EngineConceptGroup   EngineConcept = "group"
	EngineConceptCapture EngineConcept = "capture"
	EngineConceptItem    EngineConcept = "item"
	EngineConceptPort    EngineConcept = "port"
	EngineConceptLine    EngineConcept = "line"
	EngineConceptText    EngineConcept = "text"
	EngineConceptSpacer  EngineConcept = "spacer"
)

// EngineLayoutPolicy selects a generic child placement algorithm.
type EngineLayoutPolicy string

const (
	EngineLayoutVertical     EngineLayoutPolicy = "vertical"
	EngineLayoutHorizontal   EngineLayoutPolicy = "horizontal"
	EngineLayoutGrid         EngineLayoutPolicy = "grid"
	EngineLayoutAdaptiveGrid EngineLayoutPolicy = "adaptive-grid"
	EngineLayoutAbsolute     EngineLayoutPolicy = "absolute"
	EngineLayoutNone         EngineLayoutPolicy = "none"
)

type EngineOverflow string

const (
	EngineOverflowError   EngineOverflow = "error"
	EngineOverflowVisible EngineOverflow = "visible"
)

type EngineAlignment string

const (
	EngineAlignStart   EngineAlignment = "start"
	EngineAlignCenter  EngineAlignment = "center"
	EngineAlignEnd     EngineAlignment = "end"
	EngineAlignStretch EngineAlignment = "stretch"
)

type EngineJustification string

const (
	EngineJustifyStart        EngineJustification = "start"
	EngineJustifyCenter       EngineJustification = "center"
	EngineJustifyEnd          EngineJustification = "end"
	EngineJustifySpaceBetween EngineJustification = "space-between"
	EngineJustifySpaceEvenly  EngineJustification = "space-evenly"
)

type EngineShape string

const (
	EngineShapeRectangle EngineShape = "rectangle"
	EngineShapeEllipse   EngineShape = "ellipse"
	EngineShapeNone      EngineShape = "none"
)

type EngineSide string

const (
	EngineSideAuto   EngineSide = "auto"
	EngineSideTop    EngineSide = "top"
	EngineSideRight  EngineSide = "right"
	EngineSideBottom EngineSide = "bottom"
	EngineSideLeft   EngineSide = "left"
)

type EngineRoutingPolicy string

const (
	EngineRoutingOrthogonal EngineRoutingPolicy = "orthogonal"
	EngineRoutingStraight   EngineRoutingPolicy = "straight"
)

type EngineLineStyle string

const (
	EngineLineSolid  EngineLineStyle = "solid"
	EngineLineDashed EngineLineStyle = "dashed"
	EngineLineDotted EngineLineStyle = "dotted"
)

type EngineDecoration string

const (
	EngineDecorationNone     EngineDecoration = "none"
	EngineDecorationArrow    EngineDecoration = "arrow"
	EngineDecorationTriangle EngineDecoration = "triangle"
	EngineDecorationDiamond  EngineDecoration = "diamond"
	EngineDecorationCircle   EngineDecoration = "circle"
)

type EngineIconMissingPolicy string

const (
	EngineIconMissingError    EngineIconMissingPolicy = "error"
	EngineIconMissingFallback EngineIconMissingPolicy = "fallback"
	EngineIconMissingHide     EngineIconMissingPolicy = "hide"
)

// EngineInsets uses pointers so unset remains different from an explicit zero.
type EngineInsets struct {
	Top    *float64
	Right  *float64
	Bottom *float64
	Left   *float64
}

// EngineSourceSpan identifies a source range without exposing source contents
// to the Rust engine. Span IDs are stable within one request.
type EngineSourceSpan struct {
	ID     uint32
	File   string
	Offset int
	Line   int
	Column int
	Length int
}

// EngineParameterSource records where and how one concrete parameter was
// selected. Origin is a closed frontend vocabulary such as explicit, class,
// inherited, profile, or default.
type EngineParameterSource struct {
	Parameter string
	Origin    string
	SpanID    uint32
}

// EngineElementProvenance preserves source-authored identity for editor and
// project-intelligence adapters. It remains on the Go side and is never
// encoded into the domain-neutral Rust engine ABI.
type EngineElementProvenance struct {
	Tag       string
	Path      string
	Identity  string
	Name      string
	Detail    string
	SourceRef string
	TargetRef string
	Position  Position
}

type EngineDiagnostic struct {
	Code      string
	Severity  string
	Stage     string
	ElementID string
	Parameter string
	SpanID    uint32
	Message   string
}

type EngineDiagnosticError struct {
	Diagnostic EngineDiagnostic
}

func (rcvr *EngineDiagnosticError) Error() string {
	return rcvr.Diagnostic.Message
}

type EngineVisualSpec struct {
	Shape        EngineShape
	Fill         string
	Stroke       string
	StrokeWidth  *float64
	CornerRadius *float64
	Opacity      *float64
	Visible      *bool
	Layer        *int32
}

type EngineTextSpec struct {
	Value      string
	FontFamily string
	Color      string
	Role       string
	FontSize   *float64
	LineHeight *float64
	Wrap       *bool
	Fit        *bool
	Clip       *bool
	Padding    EngineInsets
}

type EngineIconSpec struct {
	Ref           string
	FallbackRef   string
	Color         string
	Width         *float64
	Height        *float64
	Scale         *float64
	OffsetX       *float64
	OffsetY       *float64
	MissingPolicy EngineIconMissingPolicy
}

type EnginePortSpec struct {
	Side    EngineSide
	Anchor  *float64
	Offset  *float64
	Size    *float64
	Visible *bool
	Label   string
}

type EngineLineSpec struct {
	Source           string
	Target           string
	SourceSide       EngineSide
	TargetSide       EngineSide
	SourceAnchor     *float64
	TargetAnchor     *float64
	Routing          EngineRoutingPolicy
	ObstacleMargin   *float64
	Style            EngineLineStyle
	SourceDecoration EngineDecoration
	TargetDecoration EngineDecoration
	Label            string
	LabelPosition    *float64
}

// EngineElementSpec is one unresolved generic concept. Children are lowered
// to parent-indexed ABI records. Raw source tags and plugin identifiers stay on
// the Go side; the optional AWS field carries closed native composition models.
// Width, Height, and Weight retain source compatibility with the original flat
// V2 prototype.
type EngineElementSpec struct {
	ID         string
	Concept    EngineConcept
	SpanID     uint32
	Sources    []EngineParameterSource
	Provenance *EngineElementProvenance

	X               *float64
	Y               *float64
	Width           *float64
	Height          *float64
	IntrinsicWidth  *float64
	IntrinsicHeight *float64
	MinWidth        *float64
	MaxWidth        *float64
	MinHeight       *float64
	MaxHeight       *float64
	OffsetX         *float64
	OffsetY         *float64
	Weight          *float64

	Margin  EngineInsets
	Padding EngineInsets

	Layout     EngineLayoutPolicy
	Gap        *float64
	Columns    *uint16
	ColumnSpan *uint16
	RowSpan    *uint16
	Align      EngineAlignment
	Justify    EngineJustification
	Overflow   EngineOverflow

	Visual EngineVisualSpec
	Text   *EngineTextSpec
	Icon   *EngineIconSpec
	Port   *EnginePortSpec
	Line   *EngineLineSpec
	AWS    *EngineAWSComponentSpec

	Children []EngineElementSpec
}

// EngineDocumentSpec is the typed Go-side request for generic allocation.
// Direction and Gap are retained as compatibility fields; Layout is the native
// V2 policy and takes precedence when set.
type EngineDocumentSpec struct {
	Direction EngineDirection
	Layout    EngineLayoutPolicy
	Width     float64
	Height    float64
	Gap       float64
	Padding   EngineInsets
	Overflow  EngineOverflow
	Columns   *uint16
	Elements  []EngineElementSpec
	Spans     []EngineSourceSpan
}

type EnginePoint struct {
	X float64
	Y float64
}

type EngineResolvedVisual struct {
	Shape        EngineShape
	Fill         string
	Stroke       string
	StrokeWidth  float64
	CornerRadius float64
	Opacity      float64
	Visible      bool
	Layer        int32
}

type EngineResolvedText struct {
	Value      string
	FontFamily string
	Color      string
	Role       string
	FontSize   float64
	LineHeight float64
	X          float64
	Y          float64
	Width      float64
	Height     float64
}

type EngineResolvedLine struct {
	Style            EngineLineStyle
	SourceDecoration EngineDecoration
	TargetDecoration EngineDecoration
	Label            string
	LabelPosition    float64
}

// EngineResolvedElement contains immutable geometry and renderer-neutral draw
// parameters calculated by the Rust engine. Elements remain in deterministic
// pre-order and ParentID reconstructs their hierarchy.
type EngineResolvedElement struct {
	ID         string
	ParentID   string
	Concept    EngineConcept
	X          float64
	Y          float64
	Width      float64
	Height     float64
	Visual     EngineResolvedVisual
	Text       EngineResolvedText
	Line       EngineResolvedLine
	IconRef    string
	IconX      float64
	IconY      float64
	IconWidth  float64
	IconHeight float64
	Points     []EnginePoint
}

// EngineResolvedDocument is the immutable resolved result consumed by output
// projections. Width and Height retain the requested document dimensions.
type EngineResolvedDocument struct {
	Width    float64
	Height   float64
	Elements []EngineResolvedElement
}

// EngineSVG is a safe canonical SVG returned by the Rust SVG engine.
type EngineSVG struct {
	Data    []byte
	ViewBox string
	Width   float64
	Height  float64
}
