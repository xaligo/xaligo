// Package entity contains the renderer-shared presentation scene and flat draw
// plan consumed by xaligo output encoders.
//
// All geometry — bounds detection, paper-size scaling and centring, obstacle
// collection, connector anchoring and orthogonal routing, coordinate conversion,
// text layout, and colour normalisation — is resolved before an encoder consumes
// Plan. SVG and PPTX therefore serialize the same drawing decisions.
package entity

// PresentationScene is the canonical scene subset used by plan and model
// encoders. Its JSON remains Excalidraw-compatible for editable output.
type PresentationScene struct {
	Elements []Element            `json:"elements"`
	Files    map[string]SceneFile `json:"files"`
	AppState *AppState            `json:"appState"`
}

// PptxScene is kept as a source-compatible alias for older callers.
// Deprecated: use PresentationScene.
type PptxScene = PresentationScene

type AppState struct {
	ViewBackgroundColor string `json:"viewBackgroundColor"`
}

type SceneFile struct {
	DataURL string `json:"dataURL"`
}

type Binding struct {
	ElementID  string    `json:"elementId"`
	FixedPoint []float64 `json:"fixedPoint"`
	Gap        float64   `json:"gap"`
}

type Element struct {
	ID              string      `json:"id"`
	Type            string      `json:"type"`
	X               float64     `json:"x"`
	Y               float64     `json:"y"`
	Width           float64     `json:"width"`
	Height          float64     `json:"height"`
	Angle           float64     `json:"angle"`
	Opacity         *float64    `json:"opacity"`
	StrokeColor     string      `json:"strokeColor"`
	BackgroundColor string      `json:"backgroundColor"`
	StrokeWidth     float64     `json:"strokeWidth"`
	StrokeStyle     string      `json:"strokeStyle"`
	Text            string      `json:"text"`
	RawText         string      `json:"rawText"`
	FontSize        *float64    `json:"fontSize"`
	FontFamily      *int        `json:"fontFamily"`
	FontStyle       string      `json:"fontStyle"`
	TextAlign       string      `json:"textAlign"`
	VerticalAlign   string      `json:"verticalAlign"`
	LineHeight      *float64    `json:"lineHeight"`
	FileID          string      `json:"fileId"`
	Points          [][]float64 `json:"points"`
	IsDeleted       bool        `json:"isDeleted"`
	StartBinding    *Binding    `json:"startBinding"`
	EndBinding      *Binding    `json:"endBinding"`
	CustomData      *CustomData `json:"customData"`
}

type CustomData struct {
	PageFrame                        bool        `json:"xaligoPageFrame,omitempty"`
	FrameID                          string      `json:"xaligoFrameID,omitempty"`
	FrameMetadata                    bool        `json:"xaligoFrameMetadata,omitempty"`
	FrameMetadataContent             bool        `json:"xaligoFrameMetadataContent,omitempty"`
	FrameMetadataReserved            bool        `json:"xaligoFrameMetadataReserved,omitempty"`
	ConnectorKind                    string      `json:"xaligoConnectorKind"`
	ConnectorStartArrowhead          string      `json:"xaligoConnectorStartArrowhead"`
	ConnectorEndArrowhead            string      `json:"xaligoConnectorEndArrowhead"`
	ConnectorStyleSourceKnown        bool        `json:"xaligoConnectorStyleSourceKnown,omitempty"`
	ConnectorStartArrowheadExplicit  bool        `json:"xaligoConnectorStartArrowheadExplicit,omitempty"`
	ConnectorEndArrowheadExplicit    bool        `json:"xaligoConnectorEndArrowheadExplicit,omitempty"`
	ConnectorStrokeWidthExplicit     bool        `json:"xaligoConnectorStrokeWidthExplicit,omitempty"`
	ConnectorBends                   string      `json:"xaligoConnectorBends,omitempty"`
	ConnectorScale                   float64     `json:"xaligoConnectorScale,omitempty"`
	ConnectorGrid                    float64     `json:"xaligoConnectorGrid,omitempty"`
	ConnectorSrcAnchor               bool        `json:"xaligoConnectorSrcAnchor,omitempty"`
	ConnectorDstAnchor               bool        `json:"xaligoConnectorDstAnchor,omitempty"`
	ConnectorCrossFrame              bool        `json:"xaligoCrossFrame,omitempty"`
	ConnectorSourceFrame             string      `json:"xaligoSourceFrame,omitempty"`
	ConnectorDestinationFrame        string      `json:"xaligoDestinationFrame,omitempty"`
	ConnectorSourceFrameSide         string      `json:"xaligoConnectorSourceFrameSide,omitempty"`
	ConnectorDestinationFrameSide    string      `json:"xaligoConnectorDestinationFrameSide,omitempty"`
	ConnectorSourceFrameAnchor       string      `json:"xaligoConnectorSourceFrameAnchor,omitempty"`
	ConnectorDestinationFrameAnchor  string      `json:"xaligoConnectorDestinationFrameAnchor,omitempty"`
	ConnectorLogicalID               string      `json:"xaligoConnectorLogicalId,omitempty"`
	ConnectorSourceElementID         string      `json:"xaligoConnectorSourceElementId,omitempty"`
	ConnectorDestinationElementID    string      `json:"xaligoConnectorDestinationElementId,omitempty"`
	Junction                         bool        `json:"xaligoJunction,omitempty"`
	GroupBorder                      bool        `json:"xaligoGroupBorder,omitempty"`
	GroupHeader                      bool        `json:"xaligoGroupHeader,omitempty"`
	GroupHeaderContent               bool        `json:"xaligoGroupHeaderContent,omitempty"`
	AnchorBackground                 bool        `json:"xaligoAnchorBackground,omitempty"`
	AnchorContent                    bool        `json:"xaligoAnchorContent,omitempty"`
	SemanticParentElementID          string      `json:"xaligoSemanticParentElementId,omitempty"`
	SemanticElementKind              string      `json:"xaligoSemanticElementKind,omitempty"`
	UMLID                            string      `json:"xaligoUmlId,omitempty"`
	UMLLocalID                       string      `json:"xaligoUmlLocalId,omitempty"`
	UMLReference                     string      `json:"xaligoUmlReference,omitempty"`
	UMLDiagramKind                   string      `json:"xaligoUmlDiagramKind,omitempty"`
	UMLElementKind                   string      `json:"xaligoUmlElementKind,omitempty"`
	UMLStereotype                    string      `json:"xaligoUmlStereotype,omitempty"`
	UMLAbstract                      string      `json:"xaligoUmlAbstract,omitempty"`
	UMLStatic                        string      `json:"xaligoUmlStatic,omitempty"`
	UMLOwnerID                       string      `json:"xaligoUmlOwnerId,omitempty"`
	UMLOwnerReference                string      `json:"xaligoUmlOwnerReference,omitempty"`
	UMLCompartmentKinds              string      `json:"xaligoUmlCompartmentKinds,omitempty"`
	UMLTimeFrom                      string      `json:"xaligoUmlTimeFrom,omitempty"`
	UMLTimeTo                        string      `json:"xaligoUmlTimeTo,omitempty"`
	UMLRelationKind                  string      `json:"xaligoUmlRelationKind,omitempty"`
	UMLRelationLabel                 string      `json:"xaligoUmlRelationLabel,omitempty"`
	UMLRelationSourceReference       string      `json:"xaligoUmlRelationSourceReference,omitempty"`
	UMLRelationDestinationReference  string      `json:"xaligoUmlRelationDestinationReference,omitempty"`
	UMLMessageOrder                  string      `json:"xaligoUmlMessageOrder,omitempty"`
	UMLMessageMode                   string      `json:"xaligoUmlMessageMode,omitempty"`
	UMLSequenceLifeline              bool        `json:"xaligoUmlSequenceLifeline,omitempty"`
	UMLSequenceLifelineHeader        bool        `json:"xaligoUmlSequenceLifelineHeader,omitempty"`
	UMLSequenceLifelineOwner         string      `json:"xaligoUmlSequenceLifelineOwner,omitempty"`
	UMLSequenceActivation            bool        `json:"xaligoUmlSequenceActivation,omitempty"`
	UMLSequenceActivationOwner       string      `json:"xaligoUmlSequenceActivationOwner,omitempty"`
	UMLSequenceStop                  bool        `json:"xaligoUmlSequenceStop,omitempty"`
	UMLSequenceStopOwner             string      `json:"xaligoUmlSequenceStopOwner,omitempty"`
	UMLGuard                         string      `json:"xaligoUmlGuard,omitempty"`
	UMLSourceMultiplicity            string      `json:"xaligoUmlSourceMultiplicity,omitempty"`
	UMLDestinationMultiplicity       string      `json:"xaligoUmlDestinationMultiplicity,omitempty"`
	UMLComponentInterfaceDestination any         `json:"xaligoUmlComponentInterfaceDestination,omitempty"`
	UMLOccurrenceAt                  string      `json:"xaligoUmlOccurrenceAt,omitempty"`
	UMLDurationFrom                  string      `json:"xaligoUmlDurationFrom,omitempty"`
	UMLDurationTo                    string      `json:"xaligoUmlDurationTo,omitempty"`
	PortLabel                        bool        `json:"xaligoPortLabel,omitempty"`
	CrossFrameLabel                  bool        `json:"xaligoCrossFrameLabel,omitempty"`
	DiffHighlight                    bool        `json:"xaligoDiffHighlight,omitempty"`
	DiffStatus                       string      `json:"xaligoDiffStatus,omitempty"`
	TextLayout                       *TextLayout `json:"xaligoTextLayout,omitempty"`
}

// ── Options driving the calculations ─────────────────────────────────────────

// PlanOptions collects every parameter that influences the geometry of the plan.
// They originate from the CLI / Go controller and are resolved by the Go plan
// builder before the resulting plan crosses an encoder boundary.
type PlanOptions struct {
	Theme             string        `json:"theme,omitempty"`
	PxPerInch         float64       `json:"pxPerInch"`
	ArrowStyle        string        `json:"arrowStyle"`
	ArrowStubPx       float64       `json:"arrowStubPx"`
	ArrowMargin       float64       `json:"arrowMarginPx"`
	PaperSize         string        `json:"paperSize"`
	Orientation       string        `json:"orientation"`
	PaperMargin       float64       `json:"paperMargin"`
	PaperMarginTop    float64       `json:"paperMarginTop"`
	PaperMarginRight  float64       `json:"paperMarginRight"`
	PaperMarginBottom float64       `json:"paperMarginBottom"`
	PaperMarginLeft   float64       `json:"paperMarginLeft"`
	LegendEntries     []LegendEntry `json:"legendEntries"`
}

// PptxOptions is kept as a source-compatible alias for older callers.
// Deprecated: use PlanOptions.
type PptxOptions = PlanOptions

// ── Output: renderer-shared physical draw plan ───────────────────────────────

// Plan is the complete, ordered list of drawing operations plus slide metadata.
// Every coordinate is already in inches and every colour is a 6-hex string.
type Plan struct {
	Slide           PlanSlide              `json:"slide"`
	Ops             []DrawOp               `json:"ops"`
	Legend          []LegendEntry          `json:"legend,omitempty"`
	ConnectorLegend []ConnectorLegendEntry `json:"connectorLegend,omitempty"`
}

// DocumentPlan is an ordered, page-oriented drawing plan. Each page is derived
// from one XAL frame unless CombineFrames was explicitly requested.
type DocumentPlan struct {
	SchemaVersion   int                    `json:"schemaVersion"`
	Pages           []DocumentPage         `json:"pages"`
	Legend          []LegendEntry          `json:"legend,omitempty"`
	ConnectorLegend []ConnectorLegendEntry `json:"connectorLegend,omitempty"`
}

// DocumentPage is one physical output page. PPTX maps it to a slide, PDF to a
// page, Excel to a worksheet, and SVG to a separate artifact.
type DocumentPage struct {
	ID    string    `json:"id"`
	Slide PlanSlide `json:"slide"`
	Ops   []DrawOp  `json:"ops"`
}

// RenderArtifact is a named output emitted by formats that can produce more
// than one file, currently SVG.
type RenderArtifact struct {
	ID   string
	Data []byte
}

type LegendEntry struct {
	CatalogID    int    `json:"catalogId"`
	Abbreviation string `json:"abbreviation"`
	OfficialName string `json:"officialName"`
	Data         string `json:"data,omitempty"`
}

type ConnectorLegendEntry struct {
	ID          string    `json:"id"`
	Kind        string    `json:"kind"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Source      string    `json:"source,omitempty"`
	Target      string    `json:"target,omitempty"`
	Line        LineStyle `json:"line"`
}

type PlanSlide struct {
	W          float64 `json:"w"`
	H          float64 `json:"h"`
	Background string  `json:"background"`
	// CropToSlide makes page-oriented encoders treat W and H as hard bounds.
	CropToSlide bool `json:"cropToSlide,omitempty"`
}

// DrawOp is a single encoder drawing instruction. Kind selects the dispatch:
// "rect" | "ellipse" | "diamond" | "polygon" | "text" | "image" | "line".
type DrawOp struct {
	ID         string  `json:"id,omitempty"`
	GroupID    string  `json:"groupId,omitempty"`
	FrontLayer bool    `json:"frontLayer,omitempty"`
	Kind       string  `json:"kind"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	W          float64 `json:"w"`
	H          float64 `json:"h"`
	Rotate     float64 `json:"rotate,omitempty"`

	// rect / ellipse / line
	Line *LineStyle `json:"line,omitempty"`
	Fill *FillStyle `json:"fill,omitempty"`

	// text
	Text     string  `json:"text,omitempty"`
	Color    string  `json:"color,omitempty"`
	FontFace string  `json:"fontFace,omitempty"`
	FontSize float64 `json:"fontSize,omitempty"`
	Bold     bool    `json:"bold,omitempty"`
	Align    string  `json:"align,omitempty"`
	Valign   string  `json:"valign,omitempty"`
	// TextLayout is renderer-neutral text overflow and placement metadata.
	// FontSize remains points while padding follows the plan's inch coordinates.
	TextLayout *TextLayout `json:"textLayout,omitempty"`

	// image
	Data         string  `json:"data,omitempty"`
	Transparency float64 `json:"transparency,omitempty"`

	// line / polyline (points are relative to the op's x/y bbox origin, inches)
	Points []PtIn `json:"points,omitempty"`
	FlipH  bool   `json:"flipH,omitempty"`
	FlipV  bool   `json:"flipV,omitempty"`
}

// TextRole describes the semantic purpose of a text operation without relying
// on renderer-specific object IDs.
type TextRole string

const (
	TextRoleLabel          TextRole = "label"
	TextRoleGroupHeader    TextRole = "group-header"
	TextRoleItemLabel      TextRole = "item-label"
	TextRolePortLabel      TextRole = "port-label"
	TextRoleConnectorLabel TextRole = "connector-label"
)

// TextFit controls how a renderer handles text larger than its content box.
type TextFit string

const (
	TextFitNone   TextFit = "none"
	TextFitShrink TextFit = "shrink"
)

// TextOverflow controls whether glyphs remain visible outside their text box.
type TextOverflow string

const (
	TextOverflowVisible TextOverflow = "visible"
	TextOverflowClip    TextOverflow = "clip"
)

// TextPadding stores text-box insets in the containing model's coordinate
// space. Canonical scenes use layout pixels; BuildPlan converts them to plan
// inches before SVG/PPTX encoding.
type TextPadding struct {
	Top    float64 `json:"top"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
	Left   float64 `json:"left"`
}

// TextLayout is the common text-box contract interpreted by SVG and PPTX
// encoders. LineHeight is a multiplier of FontSize.
type TextLayout struct {
	Role     TextRole     `json:"role,omitempty"`
	Wrap     bool         `json:"wrap"`
	Fit      TextFit      `json:"fit,omitempty"`
	Overflow TextOverflow `json:"overflow,omitempty"`
	// Clip is retained for plan JSON compatibility. New consumers should use
	// Overflow; producers keep both fields consistent during the migration.
	Clip       bool        `json:"clip"`
	LineHeight float64     `json:"lineHeight,omitempty"`
	Padding    TextPadding `json:"padding"`
}

type LineStyle struct {
	Color string `json:"color"`
	// Width is stored in points, matching FontSize. SVG converts it back to
	// output pixels with the same PPI transform used for plan geometry.
	Width              float64 `json:"width"`
	Dash               string  `json:"dash"`
	Transparency       float64 `json:"transparency"`
	BeginArrowType     string  `json:"beginArrowType,omitempty"`
	EndArrowType       string  `json:"endArrowType,omitempty"`
	BeginArrowExtendIn float64 `json:"beginArrowExtendIn,omitempty"`
	EndArrowExtendIn   float64 `json:"endArrowExtendIn,omitempty"`
}

type FillStyle struct {
	Color        string  `json:"color"`
	Transparency float64 `json:"transparency"`
}

// PtIn is an inch-space point for a polyline op. MoveTo marks the first vertex.
type PtIn struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	MoveTo bool    `json:"moveTo,omitempty"`
}
