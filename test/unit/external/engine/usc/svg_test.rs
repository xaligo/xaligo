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
        },
        icon_ref: String::new(),
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
