use crate::ent::model::document::{
    Alignment,
    Concept,
    Decoration,
    DocumentSpec,
    ElementSpec,
    IconSpec,
    Insets,
    Justification,
    LayoutPolicy,
    LineSpec,
    LineStyle,
    MissingIconPolicy,
    Overflow,
    Point,
    PortSpec,
    ResolvedDocument,
    ResolvedElement,
    ResolvedLine,
    ResolvedText,
    ResolvedVisual,
    RoutingPolicy,
    Shape,
    Side,
    TextSpec,
    VisualSpec,
};
use crate::ent::model::svg::NormalizedSvg;
use crate::util::error::{
    LayoutError,
    SvgError,
};
use crate::util::mcode::{
    LogLevel,
    MCode,
};

macro_rules! impl_enum_equality {
    ($($target:ty),+ $(,)?) => {
        $(
            impl PartialEq for $target {
                fn eq(&self, other: &Self) -> bool {
                    std::mem::discriminant(self) == std::mem::discriminant(other)
                }
            }

            impl Eq for $target {}
        )+
    };
}

macro_rules! impl_struct_equality {
    ($target:ty, $first:ident $(, $field:ident)* $(,)?) => {
        impl PartialEq for $target {
            fn eq(&self, other: &Self) -> bool {
                self.$first == other.$first $(&& self.$field == other.$field)*
            }
        }
    };
}

impl_enum_equality!(
    Alignment,
    Concept,
    Decoration,
    Justification,
    LayoutPolicy,
    LineStyle,
    LogLevel,
    MissingIconPolicy,
    Overflow,
    RoutingPolicy,
    Shape,
    Side,
);

impl_struct_equality!(Insets, top, right, bottom, left,);
impl_struct_equality!(
    VisualSpec,
    shape,
    fill,
    stroke,
    stroke_width,
    corner_radius,
    opacity,
    visible,
    layer,
);
impl_struct_equality!(
    TextSpec,
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
impl_struct_equality!(
    IconSpec,
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
impl_struct_equality!(PortSpec, side, anchor, offset, size, visible, label,);
impl_struct_equality!(
    LineSpec,
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
impl_struct_equality!(
    ElementSpec,
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
impl_struct_equality!(
    DocumentSpec,
    layout,
    width,
    height,
    gap,
    padding,
    overflow,
    columns,
    elements,
);
impl_struct_equality!(Point, x, y,);
impl_struct_equality!(
    ResolvedVisual,
    shape,
    fill,
    stroke,
    stroke_width,
    corner_radius,
    opacity,
    visible,
    layer,
);
impl_struct_equality!(
    ResolvedText,
    value,
    font_family,
    color,
    role,
    font_size,
    line_height,
);
impl_struct_equality!(
    ResolvedLine,
    style,
    source_decoration,
    target_decoration,
    label,
    label_position,
);
impl_struct_equality!(
    ResolvedElement,
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
    line,
    points,
);
impl_struct_equality!(ResolvedDocument, width, height, elements,);
impl_struct_equality!(NormalizedSvg, data, view_box, width, height,);
impl_struct_equality!(LayoutError, message,);
impl Eq for LayoutError {}
impl_struct_equality!(SvgError, message,);
impl Eq for SvgError {}
impl_struct_equality!(MCode, code, message,);
impl Eq for MCode {}
