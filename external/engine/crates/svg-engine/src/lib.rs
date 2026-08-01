use std::error::Error;
use std::fmt::{Display, Formatter};

use xaligo_layout_engine::ResolvedDocument;

const MAX_SVG_BYTES: usize = 2 * 1024 * 1024;

#[derive(Clone, Debug, PartialEq)]
pub struct NormalizedSvg {
    pub data: Vec<u8>,
    pub view_box: String,
    pub width: f64,
    pub height: f64,
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct SvgError {
    message: String,
}

impl SvgError {
    fn new(message: impl Into<String>) -> Self {
        Self {
            message: message.into(),
        }
    }
}

impl Display for SvgError {
    fn fmt(&self, formatter: &mut Formatter<'_>) -> std::fmt::Result {
        formatter.write_str(&self.message)
    }
}

impl Error for SvgError {}

pub fn normalize(input: &[u8]) -> Result<NormalizedSvg, SvgError> {
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

    let mut options = usvg::Options::default();
    options.resources_dir = None;
    options.image_href_resolver = usvg::ImageHrefResolver {
        resolve_data: Box::new(|_, _, _| None),
        resolve_string: Box::new(|_, _| None),
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

pub fn render(document: &ResolvedDocument) -> Vec<u8> {
    let mut output = String::new();
    output.push_str(r#"<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 "#);
    output.push_str(&format_number(document.width));
    output.push(' ');
    output.push_str(&format_number(document.height));
    output.push_str(r#"" width=""#);
    output.push_str(&format_number(document.width));
    output.push_str(r#"" height=""#);
    output.push_str(&format_number(document.height));
    output.push_str(r#"">"#);
    for element in &document.elements {
        output.push_str(r#"<rect id=""#);
        output.push_str(&escape_xml(&element.id));
        output.push_str(r#"" x=""#);
        output.push_str(&format_number(element.x));
        output.push_str(r#"" y=""#);
        output.push_str(&format_number(element.y));
        output.push_str(r#"" width=""#);
        output.push_str(&format_number(element.width));
        output.push_str(r#"" height=""#);
        output.push_str(&format_number(element.height));
        output.push_str(r##"" fill="none" stroke="#1e1e1e"/>"##);
    }
    output.push_str("</svg>");
    output.into_bytes()
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
    use xaligo_layout_engine::{ResolvedDocument, ResolvedElement};

    #[test]
    fn renders_deterministic_safe_svg() {
        let document = ResolvedDocument {
            width: 200.0,
            height: 100.0,
            elements: vec![ResolvedElement {
                id: "api<&\"".to_owned(),
                x: 10.0,
                y: 20.0,
                width: 80.0,
                height: 40.0,
            }],
        };
        let svg = String::from_utf8(render(&document)).expect("UTF-8 SVG");
        assert_eq!(
            svg,
            r##"<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 100" width="200" height="100"><rect id="api&lt;&amp;&quot;" x="10" y="20" width="80" height="40" fill="none" stroke="#1e1e1e"/></svg>"##
        );
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
