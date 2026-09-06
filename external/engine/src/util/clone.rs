use crate::ent::model::document::{
    Alignment, Concept, Decoration, DocumentSpec, ElementSpec, IconSpec, Insets, Justification,
    LayoutPolicy, LineSpec, LineStyle, MissingIconPolicy, Overflow, Point, PortSpec,
    ResolvedDocument, ResolvedElement, ResolvedLine, ResolvedText, ResolvedVisual, RoutingPolicy,
    Shape, Side, TextSpec, VisualSpec,
};
use crate::ent::model::svg::NormalizedSvg;
#[rustfmt::skip]
use crate::ent::model::aws::{
    Component,
    alb::Alb,
    nlb::Nlb,
    listener::{
        Listener,
        Protocol,
        MutualTls,
    },
};
use crate::util::error::{LayoutError, SvgError};
use crate::util::mcode::{LogLevel, MCode};

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
    Protocol,
    MutualTls,
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
    aws,
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
impl_struct_clone!(Alb, domain, presentation);
impl_struct_clone!(Nlb, domain, presentation);
impl_struct_clone!(
    Listener,
    presentation,
    protocol,
    port,
    mutual_tls,
    certificate,
    trust_store,
    target_group,
    backend_tls,
    backend_mtls,
    show_title,
);

impl Clone for Component {
    fn clone(&self) -> Self {
        match self {
            Self::Alb(value) => Self::Alb(value.clone()),
            Self::Nlb(value) => Self::Nlb(value.clone()),
            Self::Listener(value) => Self::Listener(value.clone()),
            Self::Feature(value) => Self::Feature(value.clone()),
        }
    }
}
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
use crate::ent::model::aws::{
    action::*, condition::*, feature::*, option::OptionSetting, presentation::Presentation,
    rule::Rule, target_group::*, target_service::*, transform::*,
};
impl_struct_clone!(Presentation, level, show, hide);
impl_struct_clone!(Feature, kind, presentation);
impl_struct_clone!(Rule, priority);
impl_struct_clone!(Condition, kind, name);
impl_struct_clone!(Match, value, key, regex);
impl_struct_clone!(Action, kind, order, target_group);
impl_struct_clone!(ForwardTarget, target_group, weight);
impl_struct_clone!(JwtClaim, name, format);
impl_struct_clone!(Transform, kind);
impl_struct_clone!(Rewrite, regex, replacement);
impl_struct_clone!(TargetGroup, name, target_type, protocol, port);
impl_struct_clone!(Target, name, port, zone);
impl_struct_clone!(TargetService, kind, name, reference);
impl Clone for ServiceKind {
    fn clone(&self) -> Self {
        match self {
            Self::Ecs => Self::Ecs,
            Self::Eks => Self::Eks,
            Self::Ec2 => Self::Ec2,
            Self::Lambda => Self::Lambda,
            Self::Ip => Self::Ip,
        }
    }
}
impl_struct_clone!(OptionSetting, key, value, name);
impl Clone for ConditionKind {
    fn clone(&self) -> Self {
        match self {
            Self::Host => Self::Host,
            Self::Path => Self::Path,
            Self::Header => Self::Header,
            Self::Method => Self::Method,
            Self::Query => Self::Query,
            Self::SourceIp => Self::SourceIp,
        }
    }
}
impl Clone for ActionKind {
    fn clone(&self) -> Self {
        match self {
            Self::Forward => Self::Forward,
            Self::Redirect => Self::Redirect,
            Self::FixedResponse => Self::FixedResponse,
            Self::Oidc => Self::Oidc,
            Self::Cognito => Self::Cognito,
            Self::Jwt => Self::Jwt,
        }
    }
}
impl Clone for TransformKind {
    fn clone(&self) -> Self {
        match self {
            Self::Host => Self::Host,
            Self::Url => Self::Url,
        }
    }
}
impl Clone for FeatureKind {
    fn clone(&self) -> Self {
        match self {
            Self::JwtClaim(v) => Self::JwtClaim(v.clone()),
            Self::Rule(v) => Self::Rule(v.clone()),
            Self::Condition(v) => Self::Condition(v.clone()),
            Self::Match(v) => Self::Match(v.clone()),
            Self::Action(v) => Self::Action(v.clone()),
            Self::ForwardTarget(v) => Self::ForwardTarget(v.clone()),
            Self::Transform(v) => Self::Transform(v.clone()),
            Self::Rewrite(v) => Self::Rewrite(v.clone()),
            Self::TargetGroup(v) => Self::TargetGroup(v.clone()),
            Self::TargetService(v) => Self::TargetService(v.clone()),
            Self::Target(v) => Self::Target(v.clone()),
            Self::Option(v) => Self::Option(v.clone()),
        }
    }
}
