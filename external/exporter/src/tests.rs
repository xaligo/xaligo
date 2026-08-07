use std::io::{Cursor, Read};

use zip::ZipArchive;

use crate::export;

fn package_file(bytes: &[u8], path: &str) -> Vec<u8> {
    let mut archive = ZipArchive::new(Cursor::new(bytes)).expect("open PPTX package");
    let mut file = archive.by_name(path).expect("find PPTX package part");
    let mut output = Vec::new();
    file.read_to_end(&mut output)
        .expect("read PPTX package part");
    output
}

#[test]
fn exports_pages_shapes_groups_text_and_legends() {
    let input = r#"{"plan":{"schemaVersion":2,"pages":[{"id":"one","slide":{"w":4,"h":3,"background":"FFFFFF"},"ops":[{"id":"decision","kind":"diamond","x":0.5,"y":0.5,"w":1,"h":1},{"id":"shape&1","groupId":"group&1","kind":"rect","x":2,"y":0.5,"w":1,"h":0.5},{"id":"label&1","groupId":"group&1","kind":"text","x":2,"y":1,"w":1,"h":0.2,"text":"label","textLayout":{"wrap":false,"overflow":"clip","clip":true,"padding":{}}}]},{"id":"two","slide":{"w":4,"h":3,"background":"F8FAFC"},"ops":[]}],"connectorLegend":[{"id":"L01","kind":"connection","label":"Line","description":"Shared","line":{"color":"2563EB","width":1,"dash":"solid","transparency":0,"endArrowType":"stealth"}}]},"options":{"compression":false}}"#;
    let bytes = export(input).expect("export PPTX");
    let first =
        String::from_utf8(package_file(&bytes, "ppt/slides/slide1.xml")).expect("slide XML");
    let legend =
        String::from_utf8(package_file(&bytes, "ppt/slides/slide3.xml")).expect("legend XML");
    assert!(first.contains("prst=\"diamond\""));
    assert!(first.contains("name=\"xaligo anchor group&amp;1\""));
    assert!(first.contains("horzOverflow=\"clip\" vertOverflow=\"clip\""));
    assert!(legend.contains("<a:t>Line Legend</a:t>"));
}

#[test]
fn embeds_svg_with_png_fallback() {
    let input = r#"{"plan":{"slide":{"w":4,"h":3,"background":"FFFFFF"},"ops":[{"id":"icon","kind":"image","x":0.5,"y":0.5,"w":1,"h":1,"data":"data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciLz4="}]}}"#;
    let bytes = export(input).expect("export PPTX");
    assert_eq!(
        package_file(&bytes, "ppt/media/image1_2.svg"),
        br#"<svg xmlns="http://www.w3.org/2000/svg"/>"#
    );
    assert!(package_file(&bytes, "ppt/media/image1_2.png").starts_with(&[137, 80, 78, 71]));
    let relationships = String::from_utf8(package_file(&bytes, "ppt/slides/_rels/slide1.xml.rels"))
        .expect("relationships XML");
    assert!(relationships.contains("../media/image1_2.svg"));
}

#[test]
fn rejects_empty_and_mixed_size_plans() {
    assert!(
        export("")
            .expect_err("empty request must fail")
            .to_string()
            .contains("required")
    );
    let mixed = r#"{"plan":{"schemaVersion":2,"pages":[{"id":"one","slide":{"w":4,"h":3,"background":"FFFFFF"},"ops":[]},{"id":"two","slide":{"w":5,"h":3,"background":"FFFFFF"},"ops":[]}]}}"#;
    assert!(
        export(mixed)
            .expect_err("mixed sizes must fail")
            .to_string()
            .contains("one slide size")
    );
}
