use std::collections::{HashMap, HashSet};
use std::fmt::Write as FmtWrite;
use std::io::{Cursor, Read, Write};

use base64::Engine;
use pptx::presentation::Presentation;
use pptx::slide::SlideLayoutRef;
use zip::write::SimpleFileOptions;
use zip::{CompressionMethod, ZipArchive, ZipWriter};

use crate::cnf::exporter::{DEFAULT_BACKGROUND, EMU_PER_INCH};
use crate::ent::model::pptx::{
    ConnectorLegendEntry, Fill, LegendEntry, Line, Op, Padding, Slide, TextLayout,
};
use crate::util::error::Error;

pub fn slide_xml(slide: &Slide, ops: &[Op]) -> Result<String, Error> {
    let mut xml = format!(r#"<?xml version="1.0" encoding="UTF-8" standalone="yes"?><p:sld xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:p="http://schemas.openxmlformats.org/presentationml/2006/main"><p:cSld><p:bg><p:bgPr><a:solidFill><a:srgbClr val="{}"/></a:solidFill><a:effectLst/></p:bgPr></p:bg><p:spTree><p:nvGrpSpPr><p:cNvPr id="1" name=""/><p:cNvGrpSpPr/><p:nvPr/></p:nvGrpSpPr><p:grpSpPr><a:xfrm><a:off x="0" y="0"/><a:ext cx="0" cy="0"/><a:chOff x="0" y="0"/><a:chExt cx="0" cy="0"/></a:xfrm></p:grpSpPr>"#, color(&slide.background, DEFAULT_BACKGROUND));
    let mut grouped = HashSet::new();
    let mut next_id = 2;
    for front in [false, true] {
        for op in ops.iter().filter(|op| op.front_layer == front) {
            if let Some(group) = &op.group_id {
                if grouped.insert(group.clone()) {
                    let members: Vec<&Op> = ops.iter().filter(|item| item.group_id.as_ref() == Some(group)).collect();
                    write_group(&mut xml, &mut next_id, group, &members)?;
                }
            } else {
                write_op(&mut xml, next_id, op)?;
                next_id += 1;
            }
        }
    }
    xml.push_str("</p:spTree></p:cSld><p:clrMapOvr><a:masterClrMapping/></p:clrMapOvr></p:sld>");
    Ok(xml)
}

pub fn add_legends(
    presentation: &mut Presentation,
    layout: &SlideLayoutRef,
    slide: &Slide,
    legend: &[LegendEntry],
    connector_legend: &[ConnectorLegendEntry],
) -> Result<Vec<Vec<Op>>, Error> {
    let mut pages = Vec::new();
    if !connector_legend.is_empty() {
        let mut ops = vec![text_op("Line Legend", 0.35, 0.25, slide.w - 0.7, 0.35, 16.0, true)];
        for (index, entry) in connector_legend.iter().enumerate() {
            let y = 0.8 + index as f64 * 0.3;
            ops.push(Op { id: Some(format!("legend-line-{index}")), kind: "line".into(), x: 0.35, y: y + 0.1, w: 0.8, h: 0.0, line: Some(entry.line.clone()), ..empty_op() });
            ops.push(text_op(&format!("{}  {}  {}", entry.id, entry.label, entry.description), 1.3, y, slide.w - 1.65, 0.22, 7.0, false));
        }
        add_page(presentation, layout, slide, &ops)?;
        pages.push(ops);
    }
    let visible: Vec<&LegendEntry> = legend.iter().filter(|item| item.data.is_some() && !item.official_name.is_empty()).collect();
    if !visible.is_empty() {
        let mut ops = vec![text_op("Legend", 0.35, 0.25, slide.w - 0.7, 0.35, 16.0, true)];
        for (index, entry) in visible.iter().enumerate() {
            let column = index % 4;
            let row = index / 4;
            let width = (slide.w - 0.7) / 4.0;
            let x = 0.35 + column as f64 * width;
            let y = 0.8 + row as f64 * 0.3;
            ops.push(Op { id: Some(format!("legend-icon-{index}")), kind: "image".into(), x, y, w: 0.2, h: 0.2, data: entry.data.clone(), ..empty_op() });
            ops.push(text_op(&format!("{}  {}", entry.abbreviation, entry.official_name), x + 0.25, y, width - 0.28, 0.22, 6.5, false));
        }
        add_page(presentation, layout, slide, &ops)?;
        pages.push(ops);
    }
    Ok(pages)
}

pub fn finalize_package(bytes: Vec<u8>, pages: &[Vec<Op>], compression: bool) -> Result<Vec<u8>, Error> {
    let mut archive = ZipArchive::new(Cursor::new(bytes)).map_err(package_error)?;
    let mut files = Vec::new();
    for index in 0..archive.len() {
        let mut file = archive.by_index(index).map_err(package_error)?;
        let mut data = Vec::new();
        file.read_to_end(&mut data).map_err(package_error)?;
        files.push((file.name().to_string(), data));
    }
    let mut replacements = HashMap::new();
    let mut additions = Vec::new();
    if let Some((_, content_types)) = files.iter().find(|(name, _)| name == "[Content_Types].xml") {
        let original = String::from_utf8_lossy(content_types);
        let mut defaults = String::new();
        if !original.contains("Extension=\"png\"") {
            defaults.push_str(r#"<Default Extension="png" ContentType="image/png"/>"#);
        }
        if !original.contains("Extension=\"svg\"") {
            defaults.push_str(r#"<Default Extension="svg" ContentType="image/svg+xml"/>"#);
        }
        replacements.insert("[Content_Types].xml".into(), original.replace("</Types>", &format!("{defaults}</Types>")).into_bytes());
    }
    for (page_index, ops) in pages.iter().enumerate() {
        let images = image_bindings(ops);
        if images.is_empty() { continue; }
        let rel_path = format!("ppt/slides/_rels/slide{}.xml.rels", page_index + 1);
        let original = files.iter().find(|(name, _)| name == &rel_path).map(|(_, data)| String::from_utf8_lossy(data).into_owned())
            .unwrap_or_else(|| r#"<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships"></Relationships>"#.into());
        let mut relationships = String::new();
        for (shape_id, op) in images {
            let (mime, data) = decode_data_uri(op.data.as_deref().unwrap_or_default())?;
            let stem = format!("image{}_{}", page_index + 1, shape_id);
            if mime.contains("svg") {
                let fallback = render_svg_fallback(&data, op)?;
                additions.push((format!("ppt/media/{stem}.svg"), data));
                additions.push((format!("ppt/media/{stem}.png"), fallback));
                relationships.push_str(&format!(r#"<Relationship Id="rIdImg{shape_id}" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="../media/{stem}.png"/>"#));
                relationships.push_str(&format!(r#"<Relationship Id="rIdSvg{shape_id}" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="../media/{stem}.svg"/>"#));
            } else {
                additions.push((format!("ppt/media/{stem}.png"), data));
                relationships.push_str(&format!(r#"<Relationship Id="rIdImg{shape_id}" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/image" Target="../media/{stem}.png"/>"#));
            }
        }
        replacements.insert(rel_path, original.replace("</Relationships>", &format!("{relationships}</Relationships>" )).into_bytes());
    }
    let mut output = Cursor::new(Vec::new());
    let mut writer = ZipWriter::new(&mut output);
    let method = if compression { CompressionMethod::Deflated } else { CompressionMethod::Stored };
    let options = SimpleFileOptions::default().compression_method(method);
    for (name, data) in files {
        writer.start_file(&name, options).map_err(package_error)?;
        writer.write_all(replacements.get(&name).unwrap_or(&data)).map_err(package_error)?;
    }
    for (name, data) in additions {
        writer.start_file(name, options).map_err(package_error)?;
        writer.write_all(&data).map_err(package_error)?;
    }
    writer.finish().map_err(package_error)?;
    Ok(output.into_inner())
}

fn add_page(presentation: &mut Presentation, layout: &SlideLayoutRef, slide: &Slide, ops: &[Op]) -> Result<(), Error> {
    let reference = presentation.add_slide(layout)?;
    *presentation.slide_xml_mut(&reference)? = slide_xml(slide, ops)?.into_bytes();
    Ok(())
}

fn write_group(xml: &mut String, next_id: &mut usize, group: &str, ops: &[&Op]) -> Result<(), Error> {
    let min_x = ops.iter().map(|op| op.x).fold(f64::INFINITY, f64::min);
    let min_y = ops.iter().map(|op| op.y).fold(f64::INFINITY, f64::min);
    let max_x = ops.iter().map(|op| op.x + op.w).fold(f64::NEG_INFINITY, f64::max);
    let max_y = ops.iter().map(|op| op.y + op.h).fold(f64::NEG_INFINITY, f64::max);
    write!(xml, r#"<p:grpSp><p:nvGrpSpPr><p:cNvPr id="{}" name="xaligo anchor {}"/><p:cNvGrpSpPr/><p:nvPr/></p:nvGrpSpPr><p:grpSpPr><a:xfrm><a:off x="{}" y="{}"/><a:ext cx="{}" cy="{}"/><a:chOff x="{}" y="{}"/><a:chExt cx="{}" cy="{}"/></a:xfrm></p:grpSpPr>"#, *next_id, escape(group), emu(min_x), emu(min_y), emu(max_x-min_x), emu(max_y-min_y), emu(min_x), emu(min_y), emu(max_x-min_x), emu(max_y-min_y)).map_err(xml_error)?;
    *next_id += 1;
    for op in ops { write_op(xml, *next_id, op)?; *next_id += 1; }
    xml.push_str("</p:grpSp>");
    Ok(())
}

fn write_op(xml: &mut String, id: usize, op: &Op) -> Result<(), Error> {
    match op.kind.as_str() {
        "text" if op.text.as_deref().unwrap_or_default().is_empty() => Ok(()),
        "text" => write_text(xml, id, op),
        "image" if op.data.is_some() => write_image(xml, id, op),
        "rect" | "ellipse" | "diamond" | "polygon" | "line" => write_shape(xml, id, op),
        _ => Ok(()),
    }
}

fn write_shape(xml: &mut String, id: usize, op: &Op) -> Result<(), Error> {
    let geometry = if let Some(points) = &op.points {
        let mut path = format!(r#"<a:custGeom><a:avLst/><a:gdLst/><a:ahLst/><a:cxnLst/><a:rect l="l" t="t" r="r" b="b"/><a:pathLst><a:path w="{}" h="{}">"#, emu(op.w).max(1), emu(op.h).max(1));
        for (index, point) in points.iter().enumerate() {
            let tag = if index == 0 || point.move_to { "moveTo" } else { "lnTo" };
            write!(path, "<a:{tag}><a:pt x=\"{}\" y=\"{}\"/></a:{tag}>", emu(point.x), emu(point.y)).map_err(xml_error)?;
        }
        if op.kind == "polygon" { path.push_str("<a:close/>"); }
        path.push_str("</a:path></a:pathLst></a:custGeom>");
        path
    } else {
        let preset = match op.kind.as_str() { "ellipse" => "ellipse", "diamond" => "diamond", "line" => "line", _ => "rect" };
        format!(r#"<a:prstGeom prst="{preset}"><a:avLst/></a:prstGeom>"#)
    };
    let name = object_name(op).unwrap_or_else(|| format!("Shape {id}"));
    write!(xml, r#"<p:sp><p:nvSpPr><p:cNvPr id="{id}" name="{}"/><p:cNvSpPr/><p:nvPr/></p:nvSpPr><p:spPr><a:xfrm{}><a:off x="{}" y="{}"/><a:ext cx="{}" cy="{}"/></a:xfrm>{geometry}{}{}</p:spPr></p:sp>"#, escape(&name), rotation(op), emu(op.x), emu(op.y), emu(op.w), emu(op.h), fill_xml(op.fill.as_ref()), line_xml(op.line.as_ref())).map_err(xml_error)?;
    Ok(())
}

fn write_text(xml: &mut String, id: usize, op: &Op) -> Result<(), Error> {
    let layout = op.text_layout.as_ref();
    let padding = layout.map(|item| &item.padding).cloned().unwrap_or_default();
    let overflow = if layout.and_then(|item| item.overflow.as_deref()) == Some("visible") || layout.is_some_and(|item| !item.clip) { "overflow" } else { "clip" };
    let align = match op.align.as_deref() { Some("center") => "ctr", Some("right") => "r", _ => "l" };
    let anchor = match op.valign.as_deref() { Some("middle") => "ctr", Some("bottom") => "b", _ => "t" };
    let size = (op.font_size.unwrap_or(9.0).max(1.0) * 100.0).round() as i64;
    write!(xml, r#"<p:sp><p:nvSpPr><p:cNvPr id="{id}" name="{}"/><p:cNvSpPr txBox="1"/><p:nvPr/></p:nvSpPr><p:spPr><a:xfrm{}><a:off x="{}" y="{}"/><a:ext cx="{}" cy="{}"/></a:xfrm><a:prstGeom prst="rect"><a:avLst/></a:prstGeom><a:noFill/><a:ln><a:noFill/></a:ln></p:spPr><p:txBody><a:bodyPr wrap="{}" anchor="{anchor}" horzOverflow="{overflow}" vertOverflow="{overflow}" lIns="{}" tIns="{}" rIns="{}" bIns="{}"/><a:lstStyle/><a:p><a:pPr algn="{align}"/><a:r><a:rPr lang="en-US" sz="{size}" b="{}"><a:solidFill><a:srgbClr val="{}"/></a:solidFill><a:latin typeface="{}"/></a:rPr><a:t>{}</a:t></a:r><a:endParaRPr lang="en-US" sz="{size}"/></a:p></p:txBody></p:sp>"#, escape(&object_name(op).unwrap_or_else(|| format!("TextBox {id}"))), rotation(op), emu(op.x), emu(op.y), emu(op.w), emu(op.h), if layout.is_some_and(|item| item.wrap) { "square" } else { "none" }, emu(padding.left), emu(padding.top), emu(padding.right), emu(padding.bottom), if op.bold.unwrap_or(false) { 1 } else { 0 }, color(op.color.as_deref().unwrap_or("1E1E1E"), "1E1E1E"), escape(op.font_face.as_deref().unwrap_or("Helvetica")), escape(op.text.as_deref().unwrap_or_default())).map_err(xml_error)?;
    Ok(())
}

fn write_image(xml: &mut String, id: usize, op: &Op) -> Result<(), Error> {
    let name = object_name(op).unwrap_or_else(|| format!("Picture {id}"));
    let extension = if op.data.as_deref().is_some_and(|data| data.starts_with("data:image/svg")) { format!(r#"<a:extLst><a:ext uri="{{28A0092B-C50C-407E-A947-70E740481C1C}}"><asvg:svgBlip xmlns:asvg="http://schemas.microsoft.com/office/drawing/2016/SVG/main" r:embed="rIdSvg{id}"/></a:ext></a:extLst>"#) } else { String::new() };
    write!(xml, r#"<p:pic><p:nvPicPr><p:cNvPr id="{id}" name="{}"/><p:cNvPicPr><a:picLocks noChangeAspect="1"/></p:cNvPicPr><p:nvPr/></p:nvPicPr><p:blipFill><a:blip r:embed="rIdImg{id}">{extension}</a:blip><a:stretch><a:fillRect/></a:stretch></p:blipFill><p:spPr><a:xfrm{}><a:off x="{}" y="{}"/><a:ext cx="{}" cy="{}"/></a:xfrm><a:prstGeom prst="rect"><a:avLst/></a:prstGeom></p:spPr></p:pic>"#, escape(&name), rotation(op), emu(op.x), emu(op.y), emu(op.w), emu(op.h)).map_err(xml_error)?;
    Ok(())
}

fn image_bindings(ops: &[Op]) -> Vec<(usize, &Op)> {
    let mut bindings = Vec::new();
    let mut grouped = HashSet::new();
    let mut next_id = 2;
    for front in [false, true] {
        for op in ops.iter().filter(|op| op.front_layer == front) {
            if let Some(group) = &op.group_id {
                if grouped.insert(group.clone()) {
                    next_id += 1;
                    for member in ops.iter().filter(|item| item.group_id.as_ref() == Some(group)) {
                        if member.kind == "image" && member.data.is_some() { bindings.push((next_id, member)); }
                        next_id += 1;
                    }
                }
            } else {
                if op.kind == "image" && op.data.is_some() { bindings.push((next_id, op)); }
                next_id += 1;
            }
        }
    }
    bindings
}

fn render_svg_fallback(svg: &[u8], op: &Op) -> Result<Vec<u8>, Error> {
    let options = resvg::usvg::Options::default();
    let tree = resvg::usvg::Tree::from_data(svg, &options).map_err(|error| Error::invalid(format!("parse PPTX SVG fallback: {error}")))?;
    let width = (op.w.max(0.01) * 96.0).round().clamp(1.0, 4096.0) as u32;
    let height = (op.h.max(0.01) * 96.0).round().clamp(1.0, 4096.0) as u32;
    let mut pixmap = resvg::tiny_skia::Pixmap::new(width, height).ok_or_else(|| Error::invalid("allocate PPTX SVG fallback"))?;
    let sx = width as f32 / tree.size().width();
    let sy = height as f32 / tree.size().height();
    resvg::render(&tree, resvg::tiny_skia::Transform::from_scale(sx, sy), &mut pixmap.as_mut());
    pixmap.encode_png().map_err(|error| Error::invalid(format!("encode PPTX SVG fallback: {error}")))
}

fn text_op(text: &str, x: f64, y: f64, w: f64, h: f64, size: f64, bold: bool) -> Op {
    Op { kind: "text".into(), x, y, w, h, text: Some(text.into()), font_size: Some(size), bold: Some(bold), text_layout: Some(TextLayout { wrap: false, fit: Some("shrink".into()), overflow: Some("clip".into()), clip: true, line_height: Some(1.2), padding: Padding::default() }), ..empty_op() }
}

fn empty_op() -> Op { Op { id: None, group_id: None, front_layer: false, kind: String::new(), x: 0.0, y: 0.0, w: 0.0, h: 0.0, rotate: None, line: None, fill: None, text: None, color: None, font_face: None, font_size: None, bold: None, align: None, valign: None, text_layout: None, data: None, transparency: None, points: None, flip_h: false, flip_v: false } }
fn line_xml(line: Option<&Line>) -> String { let Some(line)=line else{return r#"<a:ln w="12700"><a:solidFill><a:srgbClr val="1E1E1E"/></a:solidFill></a:ln>"#.into()}; let dash=match line.dash.as_str(){"dash"=>"dash","dot"=>"sysDot",_=>"solid"}; format!(r#"<a:ln w="{}"><a:solidFill><a:srgbClr val="{}"><a:alpha val="{}"/></a:srgbClr></a:solidFill><a:prstDash val="{dash}"/>{}{}</a:ln>"#,(line.width.max(0.0)*12700.0).round() as i64,color(&line.color,"1E1E1E"),alpha(line.transparency),arrow("headEnd",line.begin_arrow_type.as_deref()),arrow("tailEnd",line.end_arrow_type.as_deref())) }
fn arrow(tag:&str,value:Option<&str>)->String{let kind=match value{Some("arrow")=>"arrow",Some("diamond")=>"diamond",Some("oval")=>"oval",Some("stealth")=>"stealth",Some("triangle")=>"triangle",_=>"none"};format!("<a:{tag} type=\"{kind}\" w=\"sm\" len=\"lg\"/>")}
fn fill_xml(fill:Option<&Fill>)->String{fill.map_or_else(||"<a:noFill/>".into(),|fill|format!(r#"<a:solidFill><a:srgbClr val="{}"><a:alpha val="{}"/></a:srgbClr></a:solidFill>"#,color(&fill.color,"FFFFFF"),alpha(fill.transparency)))}
fn object_name(op:&Op)->Option<String>{let id=op.id.as_ref()?;if op.front_layer{return Some(format!("xaligo-front-layer|{id}"))}op.group_id.as_ref().map_or_else(||Some(id.clone()),|group|Some(format!("xaligo-anchor-group|{group}|{id}")))}
fn decode_data_uri(value:&str)->Result<(String,Vec<u8>),Error>{let (header,payload)=value.split_once(',').ok_or_else(||Error::invalid("PPTX image must be a data URI"))?;let mime=header.strip_prefix("data:").unwrap_or("image/png").split(';').next().unwrap_or("image/png").to_string();let bytes=if header.ends_with(";base64"){base64::engine::general_purpose::STANDARD.decode(payload).map_err(|error|Error::invalid(format!("decode PPTX image: {error}")))?}else{payload.as_bytes().to_vec()};Ok((mime,bytes))}
fn rotation(op:&Op)->String{op.rotate.filter(|v|*v!=0.0).map_or_else(String::new,|v|format!(" rot=\"{}\"",(v*60000.0).round() as i64))}
fn emu(v:f64)->i64{(v*EMU_PER_INCH).round() as i64} fn alpha(v:f64)->i64{((100.0-v.clamp(0.0,100.0))*1000.0).round() as i64}
fn color<'a>(v:&'a str,f:&'a str)->&'a str{if v.len()==6&&v.bytes().all(|b|b.is_ascii_hexdigit()){v}else{f}} fn escape(v:&str)->String{v.replace('&',"&amp;").replace('<',"&lt;").replace('>',"&gt;").replace('"',"&quot;").replace('\'',"&apos;")}
fn xml_error(error:std::fmt::Error)->Error{Error::invalid(format!("write PPTX XML: {error}"))} fn package_error(error:impl std::fmt::Display)->Error{Error::invalid(format!("write PPTX package: {error}"))}
