use xaligo_layout_engine::ResolvedDocument;

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
}
