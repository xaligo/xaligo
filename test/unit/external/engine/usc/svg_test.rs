use super::*;
#[rustfmt::skip]
use crate::ent::model::document::{
    ResolvedLine,
    ResolvedText,
    ResolvedVisual,
};

fn resolved_element(concept: Concept) -> ResolvedElement {
    ResolvedElement {
        parent: None,
        id: "api<&\"".to_owned(),
        concept,
        x: 10.0,
        y: 20.0,
        width: 80.0,
        height: 40.0,
        visual: ResolvedVisual {
            shape: if concept == Concept::Line {
                Shape::None
            } else {
                Shape::Rectangle
            },
            fill: "#ffffff".to_owned(),
            stroke: "#1e1e1e".to_owned(),
            stroke_width: 1.5,
            corner_radius: 4.0,
            opacity: 1.0,
            visible: true,
            layer: 0,
        },
        text: ResolvedText {
            value: String::new(),
            font_family: "sans-serif".to_owned(),
            color: "#111111".to_owned(),
            role: String::new(),
            font_size: 14.0,
            line_height: 1.2,
            x: 10.0,
            y: 20.0,
            width: 80.0,
            height: 40.0,
        },
        icon_ref: String::new(),
        icon_x: 0.0,
        icon_y: 0.0,
        icon_width: 0.0,
        icon_height: 0.0,
        line: ResolvedLine {
            style: LineStyle::Solid,
            source_decoration: Decoration::None,
            target_decoration: Decoration::None,
            label: String::new(),
            label_position: 0.5,
        },
        points: Vec::new(),
    }
}

#[test]
fn renders_deterministic_safe_generic_svg() {
    let document = ResolvedDocument {
        width: 200.0,
        height: 100.0,
        elements: vec![resolved_element(Concept::Item)],
    };
    let svg = String::from_utf8(render(&document)).expect("UTF-8 SVG");
    assert!(svg.starts_with(r#"<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 100""#));
    assert!(svg.contains(r#"id="api&lt;&amp;&quot;""#));
    assert!(svg.contains(r#"data-concept="item""#));
    usvg::roxmltree::Document::parse(&svg).expect("generated SVG must be valid XML");
    assert_eq!(render(&document), render(&document));
}

#[test]
fn renders_shape_line_style_and_opacity_as_valid_xml() {
    let mut group = resolved_element(Concept::Group);
    group.line.style = LineStyle::Dashed;
    group.visual.opacity = 0.5;
    let svg = String::from_utf8(render(&ResolvedDocument {
        width: 200.0,
        height: 100.0,
        elements: vec![group],
    }))
    .expect("UTF-8 SVG");
    assert!(svg.contains(r#"stroke-dasharray="8 5" opacity="0.5""#));
    usvg::roxmltree::Document::parse(&svg).expect("generated SVG must be valid XML");
}

#[test]
fn renders_v1_profile_group_header_over_the_aligned_border() {
    let mut group = resolved_element(Concept::Group);
    group.id = "cloud".to_owned();
    group.x = 24.0;
    group.y = 31.0;
    group.width = 160.0;
    group.height = 100.0;
    group.visual.stroke_width = 2.0;
    group.text.value = "AWS Cloud".to_owned();
    group.text.role = "group-header".to_owned();
    group.text.x = 58.0;
    group.text.y = 38.0;
    group.text.width = 79.0;
    group.text.height = 18.0;
    group.icon_ref = "group:AWS-Cloud-logo_32.svg".to_owned();
    group.icon_x = 22.0;
    group.icon_y = 31.0;
    group.icon_width = 32.0;
    group.icon_height = 32.0;

    let svg = String::from_utf8(render(&ResolvedDocument {
        width: 200.0,
        height: 160.0,
        elements: vec![group],
    }))
    .expect("UTF-8 SVG");
    assert!(svg.contains(r#"id="cloud" data-concept="group" x="24" y="47" width="160" height="84""#));
    assert!(svg.contains(r##"<polygon points="22,31 141,31 155,47 141,63 22,63" fill="#ffffff" stroke="#1e1e1e" stroke-width="1"/>"##));
    assert!(svg.contains(r#"id="cloud-icon" data-owner="cloud" data-concept="group" data-icon="group:AWS-Cloud-logo_32.svg" x="22" y="31" width="32" height="32""#));
    assert!(svg.contains(r#"data-owner="cloud" data-concept="group" x="58" y="47" text-anchor="start""#));
    usvg::roxmltree::Document::parse(&svg).expect("generated SVG must be valid XML");
}

#[test]
fn text_with_a_background_uses_distinct_shape_and_text_ids() {
    let mut text = resolved_element(Concept::Text);
    text.id = "metadata".to_owned();
    text.text.value = "value".to_owned();
    let svg = String::from_utf8(render(&ResolvedDocument {
        width: 200.0,
        height: 100.0,
        elements: vec![text],
    }))
    .expect("UTF-8 SVG");
    assert_eq!(svg.matches(r#"id="metadata""#).count(), 1);
    assert_eq!(svg.matches(r#"id="metadata-text""#).count(), 1);
    usvg::roxmltree::Document::parse(&svg).expect("generated SVG must be valid XML");
}

#[test]
fn renders_line_route_and_decoration() {
    let mut line = resolved_element(Concept::Line);
    line.points = vec![Point { x: 10.0, y: 20.0 }, Point { x: 90.0, y: 20.0 }];
    line.line.target_decoration = Decoration::Arrow;
    line.line.label = "calls".to_owned();
    let svg = String::from_utf8(render(&ResolvedDocument {
        width: 100.0,
        height: 50.0,
        elements: vec![line],
    }))
    .expect("UTF-8 SVG");
    assert!(svg.contains("<polyline"));
    assert!(svg.contains("<path"));
    assert!(svg.contains("calls"));
    usvg::roxmltree::Document::parse(&svg).expect("generated line SVG must be valid XML");
}

#[test]
fn renders_lines_behind_resolved_item_foreground() {
    let mut item = resolved_element(Concept::Item);
    item.id = "item".to_owned();
    item.text.value = "API".to_owned();
    item.icon_ref = "builtin:service".to_owned();
    item.icon_x = 34.0;
    item.icon_y = 22.0;
    item.icon_width = 32.0;
    item.icon_height = 24.0;

    let mut line = resolved_element(Concept::Line);
    line.id = "line".to_owned();
    line.points = vec![Point { x: 0.0, y: 40.0 }, Point { x: 100.0, y: 40.0 }];

    let svg = String::from_utf8(render(&ResolvedDocument {
        width: 100.0,
        height: 80.0,
        elements: vec![item, line],
    }))
    .expect("UTF-8 SVG");
    let line_position = svg.find("<polyline").expect("line");
    let icon_position = svg.find(r#"id="item-icon""#).expect("icon anchor");
    let text_position = svg.find(">API</text>").expect("item label");
    assert!(line_position < icon_position);
    assert!(line_position < text_position);
}

#[test]
fn renders_icon_item_labels_from_the_top_of_the_resolved_text_box() {
    let mut item = resolved_element(Concept::Item);
    item.text.value = "first\nsecond".to_owned();
    item.text.y = 50.0;
    item.text.height = 27.0;
    item.text.font_size = 8.0 * 96.0 / 72.0;
    item.text.line_height = 1.25;
    item.icon_ref = "catalog:1".to_owned();
    let svg = String::from_utf8(render(&ResolvedDocument {
        width: 200.0,
        height: 100.0,
        elements: vec![item],
    }))
    .expect("UTF-8 SVG");
    assert!(svg.contains(r#"y="60.666667" text-anchor="middle" font-family="#));
    assert!(!svg.contains("dominant-baseline"));
}

#[test]
fn renders_line_jumps_and_junctions() {
    let mut horizontal = resolved_element(Concept::Line);
    horizontal.id = "horizontal".to_owned();
    horizontal.points = vec![Point { x: 10.0, y: 50.0 }, Point { x: 90.0, y: 50.0 }];
    let mut vertical = resolved_element(Concept::Line);
    vertical.id = "vertical".to_owned();
    vertical.points = vec![Point { x: 50.0, y: 10.0 }, Point { x: 50.0, y: 90.0 }];
    let mut branch_a = resolved_element(Concept::Line);
    branch_a.id = "branch-a".to_owned();
    branch_a.points = vec![Point { x: 10.0, y: 50.0 }, Point { x: 10.0, y: 90.0 }];
    let mut branch_b = resolved_element(Concept::Line);
    branch_b.id = "branch-b".to_owned();
    branch_b.points = vec![Point { x: 10.0, y: 50.0 }, Point { x: 30.0, y: 90.0 }];
    let svg = String::from_utf8(render(&ResolvedDocument {
        width: 100.0,
        height: 100.0,
        elements: vec![horizontal, vertical, branch_a, branch_b],
    }))
    .expect("UTF-8 SVG");
    assert!(svg.contains(r##"r="5" fill="#ffffff""##));
    assert!(svg.contains(r#"r="3.5""#));
    usvg::roxmltree::Document::parse(&svg).expect("generated SVG must be valid XML");
}

#[test]
fn normalizes_safe_vector_svg() {
    let normalized = normalize(
        br##"<svg xmlns="http://www.w3.org/2000/svg" width="24" height="12" viewBox="0 0 24 12"><path id="p" d="M0 0h24v12H0z" fill="#123456"/></svg>"##,
    )
    .expect("normalize SVG");
    assert_eq!(normalized.width, 24.0);
    assert_eq!(normalized.height, 12.0);
    assert_eq!(normalized.view_box, "0 0 24 12");
    let data = String::from_utf8(normalized.data).expect("UTF-8 normalized SVG");
    assert!(data.starts_with("<svg"));
    assert!(data.contains("id=\"p\""));
}

#[test]
fn rejects_active_or_external_svg_content() {
    for input in [
        br#"<svg xmlns="http://www.w3.org/2000/svg"><script>alert(1)</script></svg>"#.as_slice(),
        br#"<svg xmlns="http://www.w3.org/2000/svg"><image href="file:///tmp/private.png"/></svg>"#.as_slice(),
        br#"<svg xmlns="http://www.w3.org/2000/svg" onload="alert(1)"/>"#.as_slice(),
        br#"<!DOCTYPE svg><svg xmlns="http://www.w3.org/2000/svg"/>"#.as_slice(),
    ] {
        assert!(normalize(input).is_err(), "accepted unsafe SVG: {input:?}");
    }
}
