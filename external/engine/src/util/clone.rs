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

macro_rules! impl_copy_clone {
    ($($target:ty),+ $(,)?) => {
        $(
            impl Clone for $target {
                fn clone(&self) -> Self {
                    *self
                }
            }

            impl Copy for $target {}
        )+
    };
}

macro_rules! impl_struct_clone {
    ($target:ty, $($field:ident),+ $(,)?) => {
        impl Clone for $target {
            fn clone(&self) -> Self {
                Self {
                    $($field: self.$field.clone(),)+
                }
            }
        }
    };
}

impl_copy_clone!(
    Alignment,
    Concept,
    Decoration,
    Insets,
    Justification,
    LayoutPolicy,
    LineStyle,
    LogLevel,
    MCode,
    MissingIconPolicy,
    Overflow,
    Point,
    RoutingPolicy,
    Shape,
    Side,
);

impl_struct_clone!(
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
impl_struct_clone!(
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
impl_struct_clone!(
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
impl_struct_clone!(PortSpec, side, anchor, offset, size, visible, label,);
impl_struct_clone!(
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
impl_struct_clone!(
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
impl_struct_clone!(
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
impl_struct_clone!(
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
impl_struct_clone!(
    ResolvedText,
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
impl_struct_clone!(
    ResolvedLine,
    style,
    source_decoration,
    target_decoration,
    label,
    label_position,
);
impl_struct_clone!(
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
    icon_x,
    icon_y,
    icon_width,
    icon_height,
    line,
    points,
);
impl_struct_clone!(ResolvedDocument, width, height, elements,);
impl_struct_clone!(NormalizedSvg, data, view_box, width, height,);
impl_struct_clone!(LayoutError, message,);
impl_struct_clone!(SvgError, message,);
