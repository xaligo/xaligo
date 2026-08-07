use crate::cnf::engine::MAX_SVG_BYTES;
#[rustfmt::skip]
use crate::ent::model::document::{
    Concept,
    Decoration,
    LineStyle,
    Point,
    ResolvedDocument,
    ResolvedElement,
    Shape,
};
use crate::ent::model::svg::NormalizedSvg;
use crate::util::error::SvgError;

pub(crate) fn normalize(input: &[u8]) -> Result<NormalizedSvg, SvgError> {
    if input.is_empty() {
        return Err(SvgError::new("SVG input must not be empty"));
    }
    if input.len() > MAX_SVG_BYTES {
        return Err(SvgError::new(format!(
            "SVG input size {} exceeds {MAX_SVG_BYTES}",
            input.len()
        )));
    }
    let text =
        std::str::from_utf8(input).map_err(|_| SvgError::new("SVG input is not valid UTF-8"))?;
    validate_svg_safety(text)?;

    let options = usvg::Options {
        resources_dir: None,
        image_href_resolver: usvg::ImageHrefResolver {
            resolve_data: Box::new(|_, _, _| None),
            resolve_string: Box::new(|_, _| None),
        },
        ..Default::default()
    };
    let tree = usvg::Tree::from_str(text, &options)
        .map_err(|error| SvgError::new(format!("parse SVG: {error}")))?;
    let width = f64::from(tree.size().width());
    let height = f64::from(tree.size().height());
    let view_box = format!("0 0 {} {}", format_number(width), format_number(height));
    let writer = usvg::WriteOptions {
        indent: usvg::Indent::None,
        attributes_indent: usvg::Indent::None,
        ..Default::default()
    };
    let data = tree.to_string(&writer).into_bytes();
    if data.len() > MAX_SVG_BYTES {
        return Err(SvgError::new(format!(
            "normalized SVG size {} exceeds {MAX_SVG_BYTES}",
            data.len()
        )));
    }
    Ok(NormalizedSvg {
        data,
        view_box,
        width,
        height,
    })
}

fn validate_svg_safety(text: &str) -> Result<(), SvgError> {
    let lowercase = text.to_ascii_lowercase();
    if lowercase.contains("<!doctype") || lowercase.contains("<!entity") {
        return Err(SvgError::new(
            "SVG DTD and entity declarations are forbidden",
        ));
    }
    let document = usvg::roxmltree::Document::parse(text)
        .map_err(|error| SvgError::new(format!("parse SVG XML: {error}")))?;
    let root = document.root_element();
    if !root.tag_name().name().eq_ignore_ascii_case("svg") {
        return Err(SvgError::new("SVG root element must be <svg>"));
    }

    for node in document.descendants().filter(|node| node.is_element()) {
        let element = node.tag_name().name().to_ascii_lowercase();
        if matches!(
            element.as_str(),
            "script"
                | "style"
                | "foreignobject"
                | "iframe"
                | "object"
                | "embed"
                | "image"
                | "audio"
                | "video"
                | "canvas"
                | "text"
        ) {
            return Err(SvgError::new(format!(
                "SVG element <{element}> is forbidden"
            )));
        }
        for attribute in node.attributes() {
            let name = attribute.name().to_ascii_lowercase();
            let value = attribute.value().trim();
            let value_lowercase = value.to_ascii_lowercase();
            if name.starts_with("on") {
                return Err(SvgError::new(format!(
                    "SVG event attribute {name:?} is forbidden"
                )));
            }
            if name == "href" && !value.starts_with('#') {
                return Err(SvgError::new("SVG href must reference a local fragment"));
            }
            if value_lowercase.contains("javascript:")
                || value_lowercase.contains("file:")
                || value_lowercase.contains("http:")
                || value_lowercase.contains("https:")
                || value_lowercase.contains("data:")
                || value_lowercase.contains("@import")
                || contains_external_url(&value_lowercase)
            {
                return Err(SvgError::new(format!(
                    "SVG attribute {name:?} contains an external or executable reference"
                )));
            }
        }
    }
    Ok(())
}

fn contains_external_url(value: &str) -> bool {
    let mut remainder = value;
    while let Some(index) = remainder.find("url(") {
        remainder = &remainder[index + 4..];
        let candidate = remainder.trim_start_matches(|character: char| {
            character.is_ascii_whitespace() || character == '\'' || character == '"'
        });
        if !candidate.starts_with('#') {
            return true;
        }
        let Some(end) = remainder.find(')') else {
            return true;
        };
        remainder = &remainder[end + 1..];
    }
    false
}

pub(crate) fn render(document: &ResolvedDocument) -> Vec<u8> {
    const DOCUMENT_OVERHEAD_BYTES: usize = 160;
    const ELEMENT_ESTIMATE_BYTES: usize = 256;
    let mut output = String::with_capacity(
        DOCUMENT_OVERHEAD_BYTES.saturating_add(
            document
                .elements
                .len()
                .saturating_mul(ELEMENT_ESTIMATE_BYTES),
        ),
    );
    output.push_str(r#"<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 "#);
    output.push_str(&format_number(document.width));
    output.push(' ');
    output.push_str(&format_number(document.height));
    output.push_str(r#"" width=""#);
    output.push_str(&format_number(document.width));
    output.push_str(r#"" height=""#);
    output.push_str(&format_number(document.height));
    output.push_str(r#"">"#);

    if document
        .elements
        .windows(2)
        .all(|pair| pair[0].visual.layer <= pair[1].visual.layer)
    {
        for element in &document.elements {
            render_element(&mut output, element);
        }
    } else {
        let mut ordered = (0..document.elements.len()).collect::<Vec<_>>();
        ordered.sort_by_key(|index| (document.elements[*index].visual.layer, *index));
        for index in ordered {
            render_element(&mut output, &document.elements[index]);
        }
    }
    render_line_crossings(&mut output, document);
    render_junctions(&mut output, document);
    output.push_str("</svg>");
    output.into_bytes()
}

fn render_line_crossings(output: &mut String, document: &ResolvedDocument) {
    let lines = document
        .elements
        .iter()
        .filter(|element| element.concept == Concept::Line && element.visual.visible)
        .collect::<Vec<_>>();
    for later_index in 1..lines.len() {
        let later = lines[later_index];
        for earlier in &lines[..later_index] {
            for later_segment in later.points.windows(2) {
                for earlier_segment in earlier.points.windows(2) {
                    let Some((point, horizontal)) = orthogonal_crossing(
                        later_segment[0], later_segment[1], earlier_segment[0], earlier_segment[1],
                    ) else {
                        continue;
                    };
                    output.push_str(r#"<circle cx=""#);
                    output.push_str(&format_number(point.x));
                    output.push_str(r#"" cy=""#);
                    output.push_str(&format_number(point.y));
                    output.push_str(r##"" r="5" fill="#ffffff"/>"##);
                    output.push_str(r#"<path d="M "#);
                    if horizontal {
                        output.push_str(&format_number(point.x - 6.0));
                        output.push(' ');
                        output.push_str(&format_number(point.y));
                        output.push_str(" Q ");
                        output.push_str(&format_number(point.x));
                        output.push(' ');
                        output.push_str(&format_number(point.y - 6.0));
                        output.push(' ');
                        output.push_str(&format_number(point.x + 6.0));
                        output.push(' ');
                        output.push_str(&format_number(point.y));
                    } else {
                        output.push_str(&format_number(point.x));
                        output.push(' ');
                        output.push_str(&format_number(point.y - 6.0));
                        output.push_str(" Q ");
                        output.push_str(&format_number(point.x + 6.0));
                        output.push(' ');
                        output.push_str(&format_number(point.y));
                        output.push(' ');
                        output.push_str(&format_number(point.x));
                        output.push(' ');
                        output.push_str(&format_number(point.y + 6.0));
                    }
                    output.push_str(r#"" fill="none" stroke=""#);
                    output.push_str(&escape_xml(&later.visual.stroke));
                    output.push_str(r#"" stroke-width=""#);
                    output.push_str(&format_number(later.visual.stroke_width));
                    output.push_str(r#""/>"#);
                }
            }
        }
    }
}

fn orthogonal_crossing(a: Point, b: Point, c: Point, d: Point) -> Option<(Point, bool)> {
    let ab_horizontal = a.y == b.y && a.x != b.x;
    let ab_vertical = a.x == b.x && a.y != b.y;
    let cd_horizontal = c.y == d.y && c.x != d.x;
    let cd_vertical = c.x == d.x && c.y != d.y;
    let (horizontal_start, horizontal_end, vertical_start, vertical_end, later_horizontal) =
        if ab_horizontal && cd_vertical {
            (a, b, c, d, true)
        } else if ab_vertical && cd_horizontal {
            (c, d, a, b, false)
        } else {
            return None;
        };
    let x = vertical_start.x;
    let y = horizontal_start.y;
    if x <= horizontal_start.x.min(horizontal_end.x)
        || x >= horizontal_start.x.max(horizontal_end.x)
        || y <= vertical_start.y.min(vertical_end.y)
        || y >= vertical_start.y.max(vertical_end.y)
    {
        return None;
    }
    Some((Point { x, y }, later_horizontal))
}

fn render_junctions(output: &mut String, document: &ResolvedDocument) {
    let mut endpoints = Vec::<(Point, usize, &str)>::new();
    for line in document.elements.iter().filter(|element| {
        element.concept == Concept::Line && element.visual.visible && element.points.len() >= 2
    }) {
        for point in [line.points[0], line.points[line.points.len() - 1]] {
            if let Some(existing) = endpoints
                .iter_mut()
                .find(|(candidate, _, _)| candidate.x == point.x && candidate.y == point.y)
            {
                existing.1 += 1;
            } else {
                endpoints.push((point, 1, line.visual.stroke.as_str()));
            }
        }
    }
    for (point, count, color) in endpoints {
        if count < 3 {
            continue;
        }
        output.push_str(r#"<circle cx=""#);
        output.push_str(&format_number(point.x));
        output.push_str(r#"" cy=""#);
        output.push_str(&format_number(point.y));
        output.push_str(r#"" r="3.5" fill=""#);
        output.push_str(&escape_xml(color));
        output.push_str(r#""/>"#);
    }
}

fn render_element(output: &mut String, element: &ResolvedElement) {
    if !element.visual.visible {
        return;
    }
    if element.concept == Concept::Line {
        render_line(output, element);
    } else {
        render_shape(output, element);
        if matches!(element.concept, Concept::Group | Concept::Frame)
            && !element.text.value.is_empty()
        {
            render_group_header(output, element);
        } else {
            render_text(output, element, &element.text.value);
        }
    }
}

fn render_group_header(output: &mut String, element: &ResolvedElement) {
    let height = 28.0_f64.min(element.height);
    let width = ((element.text.value.chars().count() as f64 * element.text.font_size * 0.62)
        + 28.0)
        .min(element.width)
        .max(height);
    let tip = (height / 2.0).min(10.0);
    polygon(
        output,
        &[
            Point { x: element.x, y: element.y },
            Point { x: element.x + width - tip, y: element.y },
            Point { x: element.x + width, y: element.y + height / 2.0 },
            Point { x: element.x + width - tip, y: element.y + height },
            Point { x: element.x, y: element.y + height },
        ],
        &element.visual.stroke,
    );
    let mut label = element.clone();
    label.x = element.x + 4.0;
    label.y = element.y;
    label.width = width - tip - 4.0;
    label.height = height;
    render_text(output, &label, &element.text.value);
}

fn render_shape(output: &mut String, element: &ResolvedElement) {
    if element.visual.shape == Shape::None {
        return;
    }
    match element.visual.shape {
        Shape::Rectangle | Shape::Default => {
            output.push_str(r#"<rect"#);
            common_attributes(output, element);
            output.push_str(r#" x=""#);
            output.push_str(&format_number(element.x));
            output.push_str(r#"" y=""#);
            output.push_str(&format_number(element.y));
            output.push_str(r#"" width=""#);
            output.push_str(&format_number(element.width));
            output.push_str(r#"" height=""#);
            output.push_str(&format_number(element.height));
            if element.visual.corner_radius > 0.0 {
                output.push_str(r#"" rx=""#);
                output.push_str(&format_number(element.visual.corner_radius));
            }
            paint_attributes(output, element);
            output.push_str("/>");
        }
        Shape::Ellipse => {
            output.push_str(r#"<ellipse"#);
            common_attributes(output, element);
            output.push_str(r#" cx=""#);
            output.push_str(&format_number(element.x + element.width / 2.0));
            output.push_str(r#"" cy=""#);
            output.push_str(&format_number(element.y + element.height / 2.0));
            output.push_str(r#"" rx=""#);
            output.push_str(&format_number(element.width / 2.0));
            output.push_str(r#"" ry=""#);
            output.push_str(&format_number(element.height / 2.0));
            paint_attributes(output, element);
            output.push_str("/>");
        }
        Shape::None => {}
    }
}

fn common_attributes(output: &mut String, element: &ResolvedElement) {
    output.push_str(r#" id=""#);
    output.push_str(&escape_xml(&element.id));
    output.push_str(r#"" data-concept=""#);
    output.push_str(concept_name(element.concept));
    if let Some(parent) = element.parent {
        output.push_str(r#"" data-parent-index=""#);
        output.push_str(&parent.to_string());
    }
    if !element.icon_ref.is_empty() {
        output.push_str(r#"" data-icon=""#);
        output.push_str(&escape_xml(&element.icon_ref));
    }
    output.push('"');
}

fn paint_attributes(output: &mut String, element: &ResolvedElement) {
    output.push_str(r#"" fill=""#);
    output.push_str(&escape_xml(&element.visual.fill));
    output.push_str(r#"" stroke=""#);
    output.push_str(&escape_xml(&element.visual.stroke));
    output.push_str(r#"" stroke-width=""#);
    output.push_str(&format_number(element.visual.stroke_width));
    if element.visual.opacity < 1.0 {
        output.push_str(r#"" opacity=""#);
        output.push_str(&format_number(element.visual.opacity));
    }
    output.push('"');
}

fn render_text(output: &mut String, element: &ResolvedElement, value: &str) {
    if value.is_empty() {
        return;
    }
    let lines = value.lines().collect::<Vec<_>>();
    let line_height = element.text.font_size * element.text.line_height;
    let first_y = element.y + element.height / 2.0
        - line_height * (lines.len().saturating_sub(1) as f64) / 2.0;
    output.push_str("<text");
    if element.concept == Concept::Text {
        output.push_str(r#" id=""#);
        output.push_str(&escape_xml(&element.id));
        output.push('"');
    }
    output.push_str(r#" data-owner=""#);
    output.push_str(&escape_xml(&element.id));
    output.push_str(r#"" data-concept=""#);
    output.push_str(concept_name(element.concept));
    output.push_str(r#"" x=""#);
    output.push_str(&format_number(element.x + element.width / 2.0));
    output.push_str(r#"" y=""#);
    output.push_str(&format_number(first_y));
    output.push_str(r#"" text-anchor="middle" dominant-baseline="middle" font-family=""#);
    output.push_str(&escape_xml(&element.text.font_family));
    output.push_str(r#"" font-size=""#);
    output.push_str(&format_number(element.text.font_size));
    output.push_str(r#"" fill=""#);
    output.push_str(&escape_xml(&element.text.color));
    output.push_str(r#"">"#);
    for (index, line) in lines.iter().enumerate() {
        if index == 0 {
            output.push_str(&escape_xml(line));
        } else {
            output.push_str(r#"<tspan x=""#);
            output.push_str(&format_number(element.x + element.width / 2.0));
            output.push_str(r#"" dy=""#);
            output.push_str(&format_number(line_height));
            output.push_str(r#"">"#);
            output.push_str(&escape_xml(line));
            output.push_str("</tspan>");
        }
    }
    output.push_str("</text>");
}

fn render_line(output: &mut String, element: &ResolvedElement) {
    if element.points.len() < 2 {
        return;
    }
    output.push_str(r#"<polyline"#);
    common_attributes(output, element);
    output.push_str(r#" points=""#);
    for (index, point) in element.points.iter().enumerate() {
        if index > 0 {
            output.push(' ');
        }
        output.push_str(&format_number(point.x));
        output.push(',');
        output.push_str(&format_number(point.y));
    }
    output.push_str(r#"" fill="none" stroke=""#);
    output.push_str(&escape_xml(&element.visual.stroke));
    output.push_str(r#"" stroke-width=""#);
    output.push_str(&format_number(element.visual.stroke_width));
    output.push_str(r#"" stroke-linejoin="round" stroke-linecap="round""#);
    match element.line.style {
        LineStyle::Solid => {}
        LineStyle::Dashed => output.push_str(r#" stroke-dasharray="8 5""#),
        LineStyle::Dotted => output.push_str(r#" stroke-dasharray="2 4""#),
    }
    if element.visual.opacity < 1.0 {
        output.push_str(r#" opacity=""#);
        output.push_str(&format_number(element.visual.opacity));
        output.push('"');
    }
    output.push_str("/>");
    render_decoration(
        output,
        element.points[0],
        element.points[1],
        element.line.source_decoration,
        element,
    );
    let end = element.points.len() - 1;
    render_decoration(
        output,
        element.points[end],
        element.points[end - 1],
        element.line.target_decoration,
        element,
    );
    if !element.line.label.is_empty() {
        let label_point = point_along_path(&element.points, element.line.label_position);
        let mut label_owner = element.clone();
        label_owner.x = label_point.x - 1.0;
        label_owner.y = label_point.y - 1.0;
        label_owner.width = 2.0;
        label_owner.height = 2.0;
        render_text(output, &label_owner, &element.line.label);
    }
}

fn render_decoration(
    output: &mut String,
    endpoint: Point,
    neighbor: Point,
    decoration: Decoration,
    element: &ResolvedElement,
) {
    if decoration == Decoration::None {
        return;
    }
    let dx = endpoint.x - neighbor.x;
    let dy = endpoint.y - neighbor.y;
    let length = (dx * dx + dy * dy).sqrt();
    if length == 0.0 {
        return;
    }
    let ux = dx / length;
    let uy = dy / length;
    let px = -uy;
    let py = ux;
    let size = 7.0 + element.visual.stroke_width;
    let back = Point {
        x: endpoint.x - ux * size,
        y: endpoint.y - uy * size,
    };
    match decoration {
        Decoration::Arrow => {
            output.push_str(r#"<path d="M "#);
            output.push_str(&point_pair(Point {
                x: back.x + px * size * 0.55,
                y: back.y + py * size * 0.55,
            }));
            output.push_str(" L ");
            output.push_str(&point_pair(endpoint));
            output.push_str(" L ");
            output.push_str(&point_pair(Point {
                x: back.x - px * size * 0.55,
                y: back.y - py * size * 0.55,
            }));
            output.push_str(r#"" fill="none" stroke=""#);
            output.push_str(&escape_xml(&element.visual.stroke));
            output.push_str(r#"" stroke-width=""#);
            output.push_str(&format_number(element.visual.stroke_width));
            output.push_str(r#""/>"#);
        }
        Decoration::Triangle => {
            polygon(
                output,
                &[
                    endpoint,
                    Point {
                        x: back.x + px * size * 0.6,
                        y: back.y + py * size * 0.6,
                    },
                    Point {
                        x: back.x - px * size * 0.6,
                        y: back.y - py * size * 0.6,
                    },
                ],
                &element.visual.stroke,
            );
        }
        Decoration::Diamond => {
            polygon(
                output,
                &[
                    endpoint,
                    Point {
                        x: endpoint.x - ux * size + px * size * 0.55,
                        y: endpoint.y - uy * size + py * size * 0.55,
                    },
                    Point {
                        x: endpoint.x - ux * size * 2.0,
                        y: endpoint.y - uy * size * 2.0,
                    },
                    Point {
                        x: endpoint.x - ux * size - px * size * 0.55,
                        y: endpoint.y - uy * size - py * size * 0.55,
                    },
                ],
                &element.visual.stroke,
            );
        }
        Decoration::Circle => {
            output.push_str(r#"<circle cx=""#);
            output.push_str(&format_number(endpoint.x - ux * size * 0.55));
            output.push_str(r#"" cy=""#);
            output.push_str(&format_number(endpoint.y - uy * size * 0.55));
            output.push_str(r#"" r=""#);
            output.push_str(&format_number(size * 0.55));
            output.push_str(r#"" fill=""#);
            output.push_str(&escape_xml(&element.visual.stroke));
            output.push_str(r#""/>"#);
        }
        Decoration::None => {}
    }
}

fn polygon(output: &mut String, points: &[Point], color: &str) {
    output.push_str(r#"<polygon points=""#);
    for (index, point) in points.iter().enumerate() {
        if index > 0 {
            output.push(' ');
        }
        output.push_str(&point_pair(*point));
    }
    output.push_str(r#"" fill=""#);
    output.push_str(&escape_xml(color));
    output.push_str(r#""/>"#);
}

fn point_pair(point: Point) -> String {
    format!("{},{}", format_number(point.x), format_number(point.y))
}

fn point_along_path(points: &[Point], position: f64) -> Point {
    let lengths = points
        .windows(2)
        .map(|segment| (segment[1].x - segment[0].x).abs() + (segment[1].y - segment[0].y).abs())
        .collect::<Vec<_>>();
    let total = lengths.iter().sum::<f64>();
    if total == 0.0 {
        return points[0];
    }
    let mut remaining = total * position;
    for (index, length) in lengths.iter().enumerate() {
        if remaining <= *length {
            let ratio = if *length == 0.0 {
                0.0
            } else {
                remaining / *length
            };
            return Point {
                x: points[index].x + (points[index + 1].x - points[index].x) * ratio,
                y: points[index].y + (points[index + 1].y - points[index].y) * ratio,
            };
        }
        remaining -= *length;
    }
    *points.last().expect("non-empty path")
}

fn concept_name(concept: Concept) -> &'static str {
    match concept {
        Concept::Frame => "frame",
        Concept::Group => "group",
        Concept::Capture => "capture",
        Concept::Item => "item",
        Concept::Port => "port",
        Concept::Line => "line",
        Concept::Text => "text",
        Concept::Spacer => "spacer",
    }
}

fn escape_xml(value: &str) -> String {
    let mut escaped = String::with_capacity(value.len());
    for character in value.chars() {
        match character {
            '&' => escaped.push_str("&amp;"),
            '<' => escaped.push_str("&lt;"),
            '>' => escaped.push_str("&gt;"),
            '"' => escaped.push_str("&quot;"),
            '\'' => escaped.push_str("&apos;"),
            _ => escaped.push(character),
        }
    }
    escaped
}

fn format_number(value: f64) -> String {
    let rendered = format!("{value:.6}");
    let trimmed = rendered.trim_end_matches('0').trim_end_matches('.');
    if trimmed == "-0" {
        "0".to_owned()
    } else {
        trimmed.to_owned()
    }
}

#[cfg(test)]
mod tests {
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
}
