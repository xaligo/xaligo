use std::fmt::{Debug, Formatter};

use crate::ent::model::document::{
    Alignment, Concept, Decoration, DocumentSpec, ElementSpec, IconSpec, Insets, Justification,
    LayoutPolicy, LineSpec, LineStyle, MissingIconPolicy, Overflow, Point, PortSpec,
    ResolvedDocument, ResolvedElement, ResolvedLine, ResolvedText, ResolvedVisual, RoutingPolicy,
    Shape, Side, TextSpec, VisualSpec,
};
use crate::ent::model::svg::NormalizedSvg;
use crate::util::error::{LayoutError, SvgError};
use crate::util::mcode::{LogLevel, MCode};

macro_rules! impl_enum_debug {
    ($target:ty, $($variant:path => $name:literal),+ $(,)?) => {
        impl Debug for $target {
            fn fmt(&self, formatter: &mut Formatter<'_>) -> std::fmt::Result {
                match self {
                    $($variant => formatter.write_str($name),)+
                }
            }
        }
    };
}

macro_rules! impl_struct_debug {
    ($target:ty, $name:literal, $($field:ident),+ $(,)?) => {
        impl Debug for $target {
            fn fmt(&self, formatter: &mut Formatter<'_>) -> std::fmt::Result {
                let mut value = formatter.debug_struct($name);
                $(value.field(stringify!($field), &self.$field);)+
                value.finish()
            }
        }
    };
}

impl_enum_debug!(
    Concept,
    Concept::Frame => "Frame",
    Concept::Group => "Group",
    Concept::Capture => "Capture",
    Concept::Item => "Item",
    Concept::Port => "Port",
    Concept::Line => "Line",
    Concept::Text => "Text",
    Concept::Spacer => "Spacer",
);
impl_enum_debug!(
    LayoutPolicy,
    LayoutPolicy::Default => "Default",
    LayoutPolicy::Vertical => "Vertical",
    LayoutPolicy::Horizontal => "Horizontal",
    LayoutPolicy::Grid => "Grid",
    LayoutPolicy::AdaptiveGrid => "AdaptiveGrid",
    LayoutPolicy::Absolute => "Absolute",
    LayoutPolicy::None => "None",
);
impl_enum_debug!(
    Overflow,
    Overflow::Error => "Error",
    Overflow::Visible => "Visible",
);
impl_enum_debug!(
    Alignment,
    Alignment::Start => "Start",
    Alignment::Center => "Center",
    Alignment::End => "End",
    Alignment::Stretch => "Stretch",
);
impl_enum_debug!(
    Justification,
    Justification::Start => "Start",
    Justification::Center => "Center",
    Justification::End => "End",
    Justification::SpaceBetween => "SpaceBetween",
    Justification::SpaceEvenly => "SpaceEvenly",
);
impl_enum_debug!(
    Shape,
    Shape::Default => "Default",
    Shape::Rectangle => "Rectangle",
    Shape::Ellipse => "Ellipse",
    Shape::None => "None",
);
impl_enum_debug!(
    Side,
    Side::Auto => "Auto",
    Side::Top => "Top",
    Side::Right => "Right",
    Side::Bottom => "Bottom",
    Side::Left => "Left",
);
impl_enum_debug!(
    RoutingPolicy,
    RoutingPolicy::Orthogonal => "Orthogonal",
    RoutingPolicy::Straight => "Straight",
);
impl_enum_debug!(
    LineStyle,
    LineStyle::Solid => "Solid",
    LineStyle::Dashed => "Dashed",
    LineStyle::Dotted => "Dotted",
);
impl_enum_debug!(
    Decoration,
    Decoration::None => "None",
    Decoration::Arrow => "Arrow",
    Decoration::Triangle => "Triangle",
    Decoration::Diamond => "Diamond",
    Decoration::Circle => "Circle",
);
impl_enum_debug!(
    MissingIconPolicy,
    MissingIconPolicy::Error => "Error",
    MissingIconPolicy::Fallback => "Fallback",
    MissingIconPolicy::Hide => "Hide",
);
impl_enum_debug!(
    LogLevel,
    LogLevel::Debug => "Debug",
    LogLevel::Info => "Info",
    LogLevel::Warn => "Warn",
    LogLevel::Error => "Error",
    LogLevel::Fatal => "Fatal",
);

impl_struct_debug!(Insets, "Insets", top, right, bottom, left,);
impl_struct_debug!(
    VisualSpec,
    "VisualSpec",
    shape,
    fill,
    stroke,
    stroke_width,
    corner_radius,
    opacity,
    visible,
    layer,
);
impl_struct_debug!(
    TextSpec,
    "TextSpec",
    value,
    font_family,
    color,
    role,
    font_size,
    line_height,
    wrap,
    fit,
    clip,
    padding,
);
impl_struct_debug!(
    IconSpec,
    "IconSpec",
    reference,
    fallback_reference,
    color,
    width,
    height,
    scale,
    offset_x,
    offset_y,
    missing_policy,
);
impl_struct_debug!(PortSpec, "PortSpec", side, anchor, offset, size, visible, label,);
impl_struct_debug!(
    LineSpec,
    "LineSpec",
    source,
    target,
    source_side,
    target_side,
    source_anchor,
    target_anchor,
    routing,
    obstacle_margin,
    style,
    source_decoration,
    target_decoration,
    label,
    label_position,
);
impl_struct_debug!(
    ElementSpec,
    "ElementSpec",
    parent,
    id,
    concept,
    layout,
    overflow,
    align,
    justify,
    x,
    y,
    width,
    height,
    intrinsic_width,
    intrinsic_height,
    min_width,
    max_width,
    min_height,
    max_height,
    offset_x,
    offset_y,
    weight,
    gap,
    margin,
    padding,
    columns,
    column_span,
    row_span,
    visual,
    text,
    icon,
    port,
    line,
);
impl_struct_debug!(
    DocumentSpec,
    "DocumentSpec",
    layout,
    width,
    height,
    gap,
    padding,
    overflow,
    columns,
    elements,
);
impl_struct_debug!(Point, "Point", x, y,);
impl_struct_debug!(
    ResolvedVisual,
    "ResolvedVisual",
    shape,
    fill,
    stroke,
    stroke_width,
    corner_radius,
    opacity,
    visible,
    layer,
);
impl_struct_debug!(
    ResolvedText,
    "ResolvedText",
    value,
    font_family,
    color,
    role,
    font_size,
    line_height,
    x,
    y,
    width,
    height,
);
impl_struct_debug!(
    ResolvedLine,
    "ResolvedLine",
    style,
    source_decoration,
    target_decoration,
    label,
    label_position,
);
impl_struct_debug!(
    ResolvedElement,
    "ResolvedElement",
    parent,
    id,
    concept,
    x,
    y,
    width,
    height,
    visual,
    text,
    icon_ref,
    icon_x,
    icon_y,
    icon_width,
    icon_height,
    line,
    points,
);
impl_struct_debug!(
    ResolvedDocument,
    "ResolvedDocument",
    width,
    height,
    elements,
);
impl_struct_debug!(
    NormalizedSvg,
    "NormalizedSvg",
    data,
    view_box,
    width,
    height,
);
impl_struct_debug!(LayoutError, "LayoutError", message,);
impl_struct_debug!(SvgError, "SvgError", message,);
impl_struct_debug!(MCode, "MCode", code, message,);
