use crate::cnf::engine::{
    ABI_VERSION, OPERATION_LAYOUT, OPERATION_NORMALIZE_SVG, OPERATION_SVG, RESPONSE_MAGIC,
    STATUS_ERROR, STATUS_OK,
};
use crate::ent::model::document::{Concept, Decoration, LineStyle, Shape};
use crate::ent::response::engine::EngineResponse;

pub(crate) trait Serialize {
    fn serialize(self) -> Vec<u8>;
}

impl Serialize for EngineResponse {
    fn serialize(self) -> Vec<u8> {
        let mut output = response_header(STATUS_OK);
        match self {
            EngineResponse::Layout(document) => {
                output.push(OPERATION_LAYOUT);
                output.extend_from_slice(&(document.elements.len() as u32).to_le_bytes());
                for element in document.elements {
                    let strings = [
                        element.id.as_str(),
                        element.text.value.as_str(),
                        element.text.font_family.as_str(),
                        element.text.color.as_str(),
                        element.text.role.as_str(),
                        element.icon_ref.as_str(),
                        element.visual.fill.as_str(),
                        element.visual.stroke.as_str(),
                        element.line.label.as_str(),
                    ];
                    output.extend_from_slice(
                        &element
                            .parent
                            .map_or(-1, |value| value as i32)
                            .to_le_bytes(),
                    );
                    output.push(encode_concept(element.concept));
                    output.push(encode_shape(element.visual.shape));
                    output.push(encode_line_style(element.line.style));
                    output.push(encode_decoration(element.line.source_decoration));
                    output.push(encode_decoration(element.line.target_decoration));
                    output.push(u8::from(element.visual.visible));
                    output.extend_from_slice(&0u16.to_le_bytes());
                    output.extend_from_slice(&element.visual.layer.to_le_bytes());
                    output.extend_from_slice(&(element.points.len() as u16).to_le_bytes());
                    output.extend_from_slice(&0u16.to_le_bytes());
                    for value in strings {
                        output.extend_from_slice(&(value.len() as u16).to_le_bytes());
                    }
                    for value in [
                        element.x,
                        element.y,
                        element.width,
                        element.height,
                        element.visual.stroke_width,
                        element.visual.corner_radius,
                        element.visual.opacity,
                        element.text.font_size,
                        element.text.line_height,
                        element.line.label_position,
                        element.text.x,
                        element.text.y,
                        element.text.width,
                        element.text.height,
                        element.icon_x,
                        element.icon_y,
                        element.icon_width,
                        element.icon_height,
                    ] {
                        output.extend_from_slice(&value.to_le_bytes());
                    }
                    for value in strings {
                        output.extend_from_slice(value.as_bytes());
                    }
                    for point in element.points {
                        output.extend_from_slice(&point.x.to_le_bytes());
                        output.extend_from_slice(&point.y.to_le_bytes());
                    }
                }
            }
            EngineResponse::Svg(svg) => {
                output.push(OPERATION_SVG);
                output.extend_from_slice(&(svg.len() as u32).to_le_bytes());
                output.extend_from_slice(&svg);
            }
            EngineResponse::NormalizedSvg(svg) => {
                let view_box = svg.view_box.as_bytes();
                output.push(OPERATION_NORMALIZE_SVG);
                output.extend_from_slice(&svg.width.to_le_bytes());
                output.extend_from_slice(&svg.height.to_le_bytes());
                output.extend_from_slice(&(view_box.len() as u16).to_le_bytes());
                output.extend_from_slice(&0u16.to_le_bytes());
                output.extend_from_slice(&(svg.data.len() as u32).to_le_bytes());
                output.extend_from_slice(view_box);
                output.extend_from_slice(&svg.data);
            }
        }
        output
    }
}

pub(crate) struct ErrorResponse<'a> {
    message: &'a str,
}

impl<'a> ErrorResponse<'a> {
    pub(crate) fn new(message: &'a str) -> Self {
        Self { message }
    }
}

impl Serialize for ErrorResponse<'_> {
    fn serialize(self) -> Vec<u8> {
        let bytes = self.message.as_bytes();
        let mut output = response_header(STATUS_ERROR);
        output.push(0);
        output.extend_from_slice(&(bytes.len() as u32).to_le_bytes());
        output.extend_from_slice(bytes);
        output
    }
}

fn encode_concept(value: Concept) -> u8 {
    match value {
        Concept::Frame => 1,
        Concept::Group => 2,
        Concept::Capture => 3,
        Concept::Item => 4,
        Concept::Port => 5,
        Concept::Line => 6,
        Concept::Text => 7,
        Concept::Spacer => 8,
    }
}

fn encode_shape(value: Shape) -> u8 {
    match value {
        Shape::Default | Shape::Rectangle => 1,
        Shape::Ellipse => 2,
        Shape::None => 3,
    }
}

fn encode_line_style(value: LineStyle) -> u8 {
    match value {
        LineStyle::Solid => 1,
        LineStyle::Dashed => 2,
        LineStyle::Dotted => 3,
    }
}

fn encode_decoration(value: Decoration) -> u8 {
    match value {
        Decoration::None => 1,
        Decoration::Arrow => 2,
        Decoration::Triangle => 3,
        Decoration::Diamond => 4,
        Decoration::Circle => 5,
    }
}

fn response_header(status: u8) -> Vec<u8> {
    let mut output = Vec::with_capacity(16);
    output.extend_from_slice(RESPONSE_MAGIC);
    output.extend_from_slice(&ABI_VERSION.to_le_bytes());
    output.push(status);
    output
}
