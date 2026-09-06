use crate::ent::model::document::*;

// Small renderer-neutral parts. No SVG strings or exporter-specific geometry.
pub(super) fn part(
    owner: &ElementSpec,
    parent: usize,
    name: &str,
    bounds: [f64; 4],
) -> ElementSpec {
    ElementSpec {
        aws: None,
        parent: Some(parent),
        id: format!("{}::aws/{name}", owner.id),
        concept: Concept::Text,
        layout: LayoutPolicy::Absolute,
        overflow: Overflow::Error,
        align: Alignment::Start,
        justify: Justification::Start,
        x: Some(bounds[0]),
        y: Some(bounds[1]),
        width: Some(bounds[2]),
        height: Some(bounds[3]),
        intrinsic_width: None,
        intrinsic_height: None,
        min_width: None,
        max_width: None,
        min_height: None,
        max_height: None,
        offset_x: None,
        offset_y: None,
        weight: None,
        gap: Some(0.0),
        margin: Insets::default(),
        padding: Insets::default(),
        columns: None,
        column_span: None,
        row_span: None,
        visual: VisualSpec {
            shape: Shape::None,
            fill: "none".into(),
            stroke: "none".into(),
            stroke_width: Some(1.0),
            corner_radius: Some(0.0),
            opacity: owner.visual.opacity,
            visible: owner.visual.visible,
            layer: owner.visual.layer,
        },
        text: TextSpec {
            value: String::new(),
            font_family: "Helvetica".into(),
            color: "#172554".into(),
            role: String::new(),
            font_size: Some(14.0),
            line_height: Some(1.25),
            wrap: Some(true),
            fit: Some(false),
            clip: Some(false),
            padding: Insets::default(),
        },
        icon: IconSpec {
            reference: String::new(),
            fallback_reference: String::new(),
            color: String::new(),
            width: None,
            height: None,
            scale: None,
            offset_x: None,
            offset_y: None,
            missing_policy: MissingIconPolicy::Error,
        },
        port: PortSpec {
            side: Side::Auto,
            anchor: None,
            offset: None,
            size: None,
            visible: None,
            label: String::new(),
        },
        line: LineSpec {
            source: String::new(),
            target: String::new(),
            source_side: Side::Auto,
            target_side: Side::Auto,
            source_anchor: None,
            target_anchor: None,
            routing: RoutingPolicy::Orthogonal,
            obstacle_margin: None,
            style: LineStyle::Solid,
            source_decoration: Decoration::None,
            target_decoration: Decoration::None,
            label: String::new(),
            label_position: None,
        },
    }
}

pub(super) fn badge(
    owner: &ElementSpec,
    parent: usize,
    name: &str,
    bounds: [f64; 4],
    label: &str,
    active: bool,
) -> ElementSpec {
    let mut element = part(owner, parent, name, bounds);
    element.visual.shape = Shape::Rectangle;
    element.visual.corner_radius = Some(8.0);
    element.visual.fill = if active { "#dcfce7" } else { "#e8eef6" }.into();
    element.text.color = if active { "#166534" } else { "#475569" }.into();
    element.text.value = label.into();
    element.text.padding.left = Some(10.0);
    element.text.padding.right = Some(8.0);
    element
}
