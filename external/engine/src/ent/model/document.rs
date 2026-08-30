pub enum Concept {
    Frame,
    Group,
    Capture,
    Item,
    Port,
    Line,
    Text,
    Spacer,
}

pub enum LayoutPolicy {
    Default,
    Vertical,
    Horizontal,
    Grid,
    AdaptiveGrid,
    Absolute,
    None,
}

pub enum Overflow {
    Error,
    Visible,
}

pub enum Alignment {
    Start,
    Center,
    End,
    Stretch,
}

pub enum Justification {
    Start,
    Center,
    End,
    SpaceBetween,
    SpaceEvenly,
}

pub enum Shape {
    Default,
    Rectangle,
    Ellipse,
    None,
}

pub enum Side {
    Auto,
    Top,
    Right,
    Bottom,
    Left,
}

pub enum RoutingPolicy {
    Orthogonal,
    Straight,
}

pub enum LineStyle {
    Solid,
    Dashed,
    Dotted,
}

pub enum Decoration {
    None,
    Arrow,
    Triangle,
    Diamond,
    Circle,
}

pub enum MissingIconPolicy {
    Error,
    Fallback,
    Hide,
}

pub struct Insets {
    pub top: Option<f64>,
    pub right: Option<f64>,
    pub bottom: Option<f64>,
    pub left: Option<f64>,
}

pub struct VisualSpec {
    pub shape: Shape,
    pub fill: String,
    pub stroke: String,
    pub stroke_width: Option<f64>,
    pub corner_radius: Option<f64>,
    pub opacity: Option<f64>,
    pub visible: Option<bool>,
    pub layer: Option<i32>,
}

pub struct TextSpec {
    pub value: String,
    pub font_family: String,
    pub color: String,
    pub role: String,
    pub font_size: Option<f64>,
    pub line_height: Option<f64>,
    pub wrap: Option<bool>,
    pub fit: Option<bool>,
    pub clip: Option<bool>,
    pub padding: Insets,
}

pub struct IconSpec {
    pub reference: String,
    pub fallback_reference: String,
    pub color: String,
    pub width: Option<f64>,
    pub height: Option<f64>,
    pub scale: Option<f64>,
    pub offset_x: Option<f64>,
    pub offset_y: Option<f64>,
    pub missing_policy: MissingIconPolicy,
}

pub struct PortSpec {
    pub side: Side,
    pub anchor: Option<f64>,
    pub offset: Option<f64>,
    pub size: Option<f64>,
    pub visible: Option<bool>,
    pub label: String,
}

pub struct LineSpec {
    pub source: String,
    pub target: String,
    pub source_side: Side,
    pub target_side: Side,
    pub source_anchor: Option<f64>,
    pub target_anchor: Option<f64>,
    pub routing: RoutingPolicy,
    pub obstacle_margin: Option<f64>,
    pub style: LineStyle,
    pub source_decoration: Decoration,
    pub target_decoration: Decoration,
    pub label: String,
    pub label_position: Option<f64>,
}

pub struct ElementSpec {
    pub parent: Option<usize>,
    pub id: String,
    pub concept: Concept,
    pub layout: LayoutPolicy,
    pub overflow: Overflow,
    pub align: Alignment,
    pub justify: Justification,
    pub x: Option<f64>,
    pub y: Option<f64>,
    pub width: Option<f64>,
    pub height: Option<f64>,
    pub intrinsic_width: Option<f64>,
    pub intrinsic_height: Option<f64>,
    pub min_width: Option<f64>,
    pub max_width: Option<f64>,
    pub min_height: Option<f64>,
    pub max_height: Option<f64>,
    pub offset_x: Option<f64>,
    pub offset_y: Option<f64>,
    pub weight: Option<f64>,
    pub gap: Option<f64>,
    pub margin: Insets,
    pub padding: Insets,
    pub columns: Option<u16>,
    pub column_span: Option<u16>,
    pub row_span: Option<u16>,
    pub visual: VisualSpec,
    pub text: TextSpec,
    pub icon: IconSpec,
    pub port: PortSpec,
    pub line: LineSpec,
}

pub struct DocumentSpec {
    pub layout: LayoutPolicy,
    pub width: f64,
    pub height: f64,
    pub gap: f64,
    pub padding: Insets,
    pub overflow: Overflow,
    pub columns: Option<u16>,
    pub elements: Vec<ElementSpec>,
}

pub struct Point {
    pub x: f64,
    pub y: f64,
}

pub struct ResolvedVisual {
    pub shape: Shape,
    pub fill: String,
    pub stroke: String,
    pub stroke_width: f64,
    pub corner_radius: f64,
    pub opacity: f64,
    pub visible: bool,
    pub layer: i32,
}

pub struct ResolvedText {
    pub value: String,
    pub font_family: String,
    pub color: String,
    pub role: String,
    pub font_size: f64,
    pub line_height: f64,
    pub x: f64,
    pub y: f64,
    pub width: f64,
    pub height: f64,
}

pub struct ResolvedLine {
    pub style: LineStyle,
    pub source_decoration: Decoration,
    pub target_decoration: Decoration,
    pub label: String,
    pub label_position: f64,
}

pub struct ResolvedElement {
    pub parent: Option<usize>,
    pub id: String,
    pub concept: Concept,
    pub x: f64,
    pub y: f64,
    pub width: f64,
    pub height: f64,
    pub visual: ResolvedVisual,
    pub text: ResolvedText,
    pub icon_ref: String,
    pub icon_x: f64,
    pub icon_y: f64,
    pub icon_width: f64,
    pub icon_height: f64,
    pub line: ResolvedLine,
    pub points: Vec<Point>,
}

pub struct ResolvedDocument {
    pub width: f64,
    pub height: f64,
    pub elements: Vec<ResolvedElement>,
}
