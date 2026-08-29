#[cfg(test)]
mod tests {
    use super::*;
    #[rustfmt::skip]
    use crate::ent::model::document::{
        Decoration,
        IconSpec,
        LineSpec,
        LineStyle,
        PortSpec,
        TextSpec,
        VisualSpec,
    };

    fn empty_text() -> TextSpec {
        TextSpec {
            value: String::new(),
            font_family: String::new(),
            color: String::new(),
            role: String::new(),
            font_size: None,
            line_height: None,
            wrap: None,
            fit: None,
            clip: None,
            padding: Insets::default(),
        }
    }

    fn element(id: &str, concept: Concept, parent: Option<usize>) -> ElementSpec {
        ElementSpec {
            parent,
            id: id.to_owned(),
            concept,
            layout: LayoutPolicy::Default,
            overflow: Overflow::Error,
            align: Alignment::Stretch,
            justify: Justification::Start,
            x: None,
            y: None,
            width: None,
            height: None,
            intrinsic_width: None,
            intrinsic_height: None,
            min_width: None,
            max_width: None,
            min_height: None,
            max_height: None,
            offset_x: None,
            offset_y: None,
            weight: None,
            gap: None,
            margin: Insets::default(),
            padding: Insets::default(),
            columns: None,
            column_span: None,
            row_span: None,
            visual: VisualSpec {
                shape: Shape::Default,
                fill: String::new(),
                stroke: String::new(),
                stroke_width: None,
                corner_radius: None,
                opacity: None,
                visible: None,
                layer: None,
            },
            text: empty_text(),
            icon: IconSpec {
                reference: String::new(),
                fallback_reference: String::new(),
                color: String::new(),
                width: None,
                height: None,
                scale: None,
                offset_x: None,
                offset_y: None,
                missing_policy: MissingIconPolicy::Fallback,
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

    fn document(elements: Vec<ElementSpec>) -> DocumentSpec {
        DocumentSpec {
            layout: LayoutPolicy::Vertical,
            width: 200.0,
            height: 300.0,
            gap: 10.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements,
        }
    }

    #[test]
    fn allocates_fixed_before_flexible_vertical_children() {
        let mut header = element("header", Concept::Item, None);
        header.height = Some(40.0);
        let mut body = element("body", Concept::Item, None);
        body.width = Some(160.0);
        body.weight = Some(1.0);
        let mut footer = element("footer", Concept::Item, None);
        footer.weight = Some(2.0);
        let resolved = resolve(&document(vec![header, body, footer])).expect("resolve document");
        assert_eq!(resolved.elements[0].height, 40.0);
        assert_eq!(resolved.elements[1].y, 50.0);
        assert_eq!(resolved.elements[1].height, 80.0);
        assert_eq!(resolved.elements[1].width, 160.0);
        assert_eq!(resolved.elements[2].y, 140.0);
        assert_eq!(resolved.elements[2].height, 160.0);
        assert_eq!(resolved.elements[2].width, 200.0);
    }

    #[test]
    fn resolves_nested_grid_ports_and_orthogonal_lines() {
        let mut frame = element("frame", Concept::Frame, None);
        frame.width = Some(200.0);
        frame.height = Some(300.0);
        frame.layout = LayoutPolicy::Grid;
        frame.columns = Some(12);
        frame.gap = Some(4.0);
        let mut left = element("left", Concept::Item, Some(0));
        left.column_span = Some(6);
        let mut port = element("left-out", Concept::Port, Some(1));
        port.port.side = Side::Right;
        let mut right = element("right", Concept::Capture, Some(0));
        right.column_span = Some(6);
        let mut line = element("flow", Concept::Line, Some(0));
        line.line.source = "left-out".to_owned();
        line.line.target = "right".to_owned();
        line.line.target_decoration = Decoration::Arrow;
        let resolved = resolve(&DocumentSpec {
            layout: LayoutPolicy::Absolute,
            width: 200.0,
            height: 300.0,
            gap: 0.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements: vec![frame, left, port, right, line],
        })
        .expect("resolve generic concepts");
        assert_eq!(resolved.elements[1].parent, Some(0));
        assert_eq!(resolved.elements[2].concept, Concept::Port);
        assert!(resolved.elements[4].points.len() >= 2);
        assert_eq!(
            resolved.elements[4].line.target_decoration,
            Decoration::Arrow
        );
    }

    #[test]
    fn rejects_non_finite_and_conflicting_constraints() {
        let mut invalid = element("bad", Concept::Item, None);
        invalid.width = Some(f64::NAN);
        assert!(resolve(&document(vec![invalid]))
            .expect_err("non-finite width")
            .to_string()
            .contains("width"));

        let mut constrained = element("bad", Concept::Item, None);
        constrained.min_width = Some(20.0);
        constrained.max_width = Some(10.0);
        assert!(resolve(&document(vec![constrained]))
            .expect_err("conflicting constraints")
            .to_string()
            .contains("minimum width"));

        let mut unsafe_paint = element("bad", Concept::Item, None);
        unsafe_paint.visual.fill = "url(https://example.com/pixel.svg)".to_owned();
        assert!(resolve(&document(vec![unsafe_paint]))
            .expect_err("unsafe paint")
            .to_string()
            .contains("unsafe paint"));
    }

    #[test]
    fn output_is_deterministic() {
        let mut left = element("left", Concept::Item, None);
        left.width = Some(40.0);
        let mut right = element("right", Concept::Item, None);
        right.weight = Some(1.0);
        let document = DocumentSpec {
            layout: LayoutPolicy::Horizontal,
            width: 240.0,
            height: 80.0,
            gap: 8.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements: vec![left, right],
        };
        assert_eq!(resolve(&document), resolve(&document));
    }

    #[test]
    fn missing_icon_hide_preserves_the_item_slot() {
        let mut item = element("service", Concept::Item, None);
        item.height = Some(48.0);
        item.icon.missing_policy = MissingIconPolicy::Hide;
        let resolved = resolve(&document(vec![item])).expect("resolve hidden icon slot");
        assert!(resolved.elements[0].visual.visible);
        assert!(resolved.elements[0].icon_ref.is_empty());
        assert_eq!(resolved.elements[0].height, 48.0);
    }

    #[test]
    fn validates_root_offsets_but_places_ports_against_the_owner_border() {
        let mut shifted = element("shifted", Concept::Item, None);
        shifted.height = Some(20.0);
        shifted.offset_x = Some(1.0);
        assert!(resolve(&document(vec![shifted]))
            .expect_err("root offset overflow")
            .to_string()
            .contains("parent content box"));

        let mut owner = element("owner", Concept::Item, None);
        owner.width = Some(100.0);
        owner.height = Some(100.0);
        owner.padding = Insets {
            top: Some(20.0),
            right: Some(20.0),
            bottom: Some(20.0),
            left: Some(20.0),
        };
        let mut port = element("owner-out", Concept::Port, Some(0));
        port.port.side = Side::Right;
        let resolved = resolve(&DocumentSpec {
            layout: LayoutPolicy::Absolute,
            width: 100.0,
            height: 100.0,
            gap: 0.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements: vec![owner, port],
        })
        .expect("port on padded owner border");
        assert_eq!(resolved.elements[1].x, 92.0);
    }
}
