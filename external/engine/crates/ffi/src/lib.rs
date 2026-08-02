use std::panic::{catch_unwind, AssertUnwindSafe};
use std::ptr;

use xaligo_layout_engine::{
    resolve, Alignment, Concept, Decoration, DocumentSpec, ElementSpec, IconSpec, Insets,
    LayoutPolicy, LineSpec, LineStyle, MissingIconPolicy, Overflow, PortSpec, ResolvedDocument,
    RoutingPolicy, Shape, Side, TextSpec, VisualSpec,
};

pub const ABI_VERSION: u16 = 2;

const REQUEST_MAGIC: &[u8; 4] = b"XLE2";
const RESPONSE_MAGIC: &[u8; 4] = b"XLR2";
const OPERATION_LAYOUT: u8 = 1;
const OPERATION_SVG: u8 = 2;
const OPERATION_NORMALIZE_SVG: u8 = 3;
const STATUS_OK: u8 = 0;
const STATUS_ERROR: u8 = 1;
const MAX_REQUEST_BYTES: usize = 16 * 1024 * 1024;
const MAX_ELEMENTS: usize = 10_000;
const NUMERIC_FIELD_COUNT: usize = 43;
const STRING_FIELD_COUNT: usize = 14;

const BOOL_VISIBLE: u16 = 1 << 0;
const BOOL_TEXT_WRAP: u16 = 1 << 1;
const BOOL_TEXT_FIT: u16 = 1 << 2;
const BOOL_TEXT_CLIP: u16 = 1 << 3;
const BOOL_PORT_VISIBLE: u16 = 1 << 4;
const BOOL_LAYER: u16 = 1 << 5;
const BOOL_COLUMNS: u16 = 1 << 6;
const BOOL_COLUMN_SPAN: u16 = 1 << 7;
const BOOL_ROW_SPAN: u16 = 1 << 8;
const BOOL_KNOWN: u16 = BOOL_VISIBLE
    | BOOL_TEXT_WRAP
    | BOOL_TEXT_FIT
    | BOOL_TEXT_CLIP
    | BOOL_PORT_VISIBLE
    | BOOL_LAYER
    | BOOL_COLUMNS
    | BOOL_COLUMN_SPAN
    | BOOL_ROW_SPAN;

const FFI_OK: i32 = 0;
const FFI_NULL_OUTPUT: i32 = 1;
const FFI_NULL_INPUT: i32 = 2;
const FFI_PANIC: i32 = 3;

#[repr(C)]
pub struct XaligoEngineBuffer {
    pub data: *mut u8,
    pub len: usize,
    pub capacity: usize,
}

impl XaligoEngineBuffer {
    const fn empty() -> Self {
        Self {
            data: ptr::null_mut(),
            len: 0,
            capacity: 0,
        }
    }

    fn from_vec(mut value: Vec<u8>) -> Self {
        let buffer = Self {
            data: value.as_mut_ptr(),
            len: value.len(),
            capacity: value.capacity(),
        };
        std::mem::forget(value);
        buffer
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn xaligo_engine_abi_version() -> u32 {
    ABI_VERSION as u32
}

/// Runs one engine operation through the versioned binary protocol.
///
/// # Safety
///
/// `output` must point to writable memory for one `XaligoEngineBuffer`. When
/// `input_len` is non-zero, `input` must point to a readable allocation of at
/// least that length. A successful output must be released exactly once with
/// `xaligo_engine_buffer_free`.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn xaligo_engine_process(
    input: *const u8,
    input_len: usize,
    output: *mut XaligoEngineBuffer,
) -> i32 {
    if output.is_null() {
        return FFI_NULL_OUTPUT;
    }
    ptr::write(output, XaligoEngineBuffer::empty());
    if input.is_null() && input_len != 0 {
        return FFI_NULL_INPUT;
    }

    let result = catch_unwind(AssertUnwindSafe(|| {
        let request = if input_len == 0 {
            &[]
        } else {
            std::slice::from_raw_parts(input, input_len)
        };
        process_request(request)
    }));

    match result {
        Ok(response) => {
            ptr::write(output, XaligoEngineBuffer::from_vec(response));
            FFI_OK
        }
        Err(_) => FFI_PANIC,
    }
}

/// Releases a buffer returned by `xaligo_engine_process`.
///
/// # Safety
///
/// The buffer must be empty or be an unchanged value returned by a successful
/// call to `xaligo_engine_process`, and it must not have been released before.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn xaligo_engine_buffer_free(buffer: XaligoEngineBuffer) {
    if buffer.data.is_null() {
        return;
    }
    drop(Vec::from_raw_parts(
        buffer.data,
        buffer.len,
        buffer.capacity,
    ));
}

pub fn process_request(input: &[u8]) -> Vec<u8> {
    match decode_request(input).and_then(execute) {
        Ok(response) => encode_success(response),
        Err(message) => encode_error(&message),
    }
}

enum EngineResponse {
    Layout(ResolvedDocument),
    Svg(Vec<u8>),
    NormalizedSvg(xaligo_svg_engine::NormalizedSvg),
}

enum EngineRequest {
    Document {
        operation: u8,
        document: DocumentSpec,
    },
    NormalizeSvg(Vec<u8>),
}

fn execute(request: EngineRequest) -> Result<EngineResponse, String> {
    match request {
        EngineRequest::Document {
            operation,
            document,
        } => {
            let resolved = resolve(&document).map_err(|error| error.to_string())?;
            match operation {
                OPERATION_LAYOUT => Ok(EngineResponse::Layout(resolved)),
                OPERATION_SVG => Ok(EngineResponse::Svg(xaligo_svg_engine::render(&resolved))),
                _ => Err(format!("unsupported engine operation {operation}")),
            }
        }
        EngineRequest::NormalizeSvg(svg) => xaligo_svg_engine::normalize(&svg)
            .map(EngineResponse::NormalizedSvg)
            .map_err(|error| error.to_string()),
    }
}

fn decode_request(input: &[u8]) -> Result<EngineRequest, String> {
    if input.len() > MAX_REQUEST_BYTES {
        return Err(format!(
            "engine request size {} exceeds {MAX_REQUEST_BYTES}",
            input.len()
        ));
    }
    let mut decoder = Decoder::new(input);
    if decoder.read_exact(4)? != REQUEST_MAGIC {
        return Err("invalid engine request magic".to_owned());
    }
    let version = decoder.read_u16()?;
    if version != ABI_VERSION {
        return Err(format!(
            "unsupported engine ABI version {version}; expected {ABI_VERSION}"
        ));
    }
    let operation = decoder.read_u8()?;
    let discriminator = decoder.read_u8()?;
    if operation == OPERATION_NORMALIZE_SVG {
        if discriminator != 0 {
            return Err("invalid SVG normalization request flags".to_owned());
        }
        let length = decoder.read_u32()? as usize;
        let svg = decoder.read_exact(length)?.to_vec();
        if !decoder.is_empty() {
            return Err("engine request has trailing bytes".to_owned());
        }
        return Ok(EngineRequest::NormalizeSvg(svg));
    }
    if operation != OPERATION_LAYOUT && operation != OPERATION_SVG {
        return Err(format!("unsupported engine operation {operation}"));
    }
    let layout = decode_layout(discriminator)?;
    let width = decoder.read_f64()?;
    let height = decoder.read_f64()?;
    let gap = decoder.read_f64()?;
    let document_flags = decoder.read_u8()?;
    if document_flags & !0b0001_1111 != 0 {
        return Err("invalid engine document flags".to_owned());
    }
    let overflow = decode_overflow(decoder.read_u8()?)?;
    let columns_value = decoder.read_u16()?;
    let mut padding_values = [0.0; 4];
    for value in &mut padding_values {
        *value = decoder.read_f64()?;
    }
    let padding = Insets {
        top: option_from_u8_flag(document_flags, 0, padding_values[0]),
        right: option_from_u8_flag(document_flags, 1, padding_values[1]),
        bottom: option_from_u8_flag(document_flags, 2, padding_values[2]),
        left: option_from_u8_flag(document_flags, 3, padding_values[3]),
    };
    let columns = if document_flags & (1 << 4) != 0 {
        Some(columns_value)
    } else {
        None
    };
    let count = decoder.read_u32()? as usize;
    if count > MAX_ELEMENTS {
        return Err(format!(
            "engine element count {count} exceeds {MAX_ELEMENTS}"
        ));
    }
    let mut elements = Vec::with_capacity(count);
    for index in 0..count {
        elements.push(decode_element(&mut decoder, index)?);
    }
    if !decoder.is_empty() {
        return Err("engine request has trailing bytes".to_owned());
    }
    Ok(EngineRequest::Document {
        operation,
        document: DocumentSpec {
            layout,
            width,
            height,
            gap,
            padding,
            overflow,
            columns,
            elements,
        },
    })
}

fn decode_element(decoder: &mut Decoder<'_>, index: usize) -> Result<ElementSpec, String> {
    let parent_value = decoder.read_i32()?;
    let parent = match parent_value {
        -1 => None,
        value if value >= 0 && (value as usize) < index => Some(value as usize),
        value => {
            return Err(format!(
                "invalid parent index {value} for engine element {index}"
            ))
        }
    };
    let numeric_flags = decoder.read_u64()?;
    if numeric_flags >> NUMERIC_FIELD_COUNT != 0 {
        return Err("invalid engine numeric flags".to_owned());
    }
    let booleans_present = decoder.read_u16()?;
    let booleans_value = decoder.read_u16()?;
    if booleans_present & !BOOL_KNOWN != 0 || booleans_value & !booleans_present != 0 {
        return Err("invalid engine boolean flags".to_owned());
    }
    let columns_value = decoder.read_u16()?;
    let column_span_value = decoder.read_u16()?;
    let row_span_value = decoder.read_u16()?;
    if decoder.read_u16()? != 0 {
        return Err("invalid engine element reserved field".to_owned());
    }
    let concept = decode_concept(decoder.read_u8()?)?;
    let layout = decode_layout(decoder.read_u8()?)?;
    let overflow = decode_overflow(decoder.read_u8()?)?;
    let shape = decode_shape(decoder.read_u8()?)?;
    let align = decode_alignment(decoder.read_u8()?)?;
    let justify = decode_justification(decoder.read_u8()?)?;
    let side = decode_side(decoder.read_u8()?)?;
    let routing = decode_routing(decoder.read_u8()?)?;
    let source_side = decode_side(decoder.read_u8()?)?;
    let target_side = decode_side(decoder.read_u8()?)?;
    let line_style = decode_line_style(decoder.read_u8()?)?;
    let source_decoration = decode_decoration(decoder.read_u8()?)?;
    let target_decoration = decode_decoration(decoder.read_u8()?)?;
    let missing_policy = decode_missing_icon(decoder.read_u8()?)?;
    let layer_value = decoder.read_i32()?;
    let mut numbers = [0.0; NUMERIC_FIELD_COUNT];
    for number in &mut numbers {
        *number = decoder.read_f64()?;
    }
    let mut lengths = [0usize; STRING_FIELD_COUNT];
    for length in &mut lengths {
        *length = decoder.read_u16()? as usize;
    }
    let mut strings: [String; STRING_FIELD_COUNT] = std::array::from_fn(|_| String::new());
    for (slot, length) in strings.iter_mut().zip(lengths) {
        let bytes = decoder.read_exact(length)?;
        *slot = std::str::from_utf8(bytes)
            .map_err(|_| "engine element string is not valid UTF-8".to_owned())?
            .to_owned();
    }
    let number = |field: usize| option_from_u64_flag(numeric_flags, field, numbers[field]);
    Ok(ElementSpec {
        parent,
        id: strings[0].clone(),
        concept,
        layout,
        overflow,
        align,
        justify,
        x: number(0),
        y: number(1),
        width: number(2),
        height: number(3),
        intrinsic_width: number(4),
        intrinsic_height: number(5),
        min_width: number(6),
        max_width: number(7),
        min_height: number(8),
        max_height: number(9),
        offset_x: number(10),
        offset_y: number(11),
        weight: number(12),
        gap: number(13),
        margin: Insets {
            top: number(14),
            right: number(15),
            bottom: number(16),
            left: number(17),
        },
        padding: Insets {
            top: number(18),
            right: number(19),
            bottom: number(20),
            left: number(21),
        },
        columns: option_from_bool_flag(booleans_present, BOOL_COLUMNS, columns_value),
        column_span: option_from_bool_flag(booleans_present, BOOL_COLUMN_SPAN, column_span_value),
        row_span: option_from_bool_flag(booleans_present, BOOL_ROW_SPAN, row_span_value),
        visual: VisualSpec {
            shape,
            fill: strings[8].clone(),
            stroke: strings[9].clone(),
            stroke_width: number(22),
            corner_radius: number(23),
            opacity: number(24),
            visible: decode_optional_bool(booleans_present, booleans_value, BOOL_VISIBLE),
            layer: if booleans_present & BOOL_LAYER != 0 {
                Some(layer_value)
            } else {
                None
            },
        },
        text: TextSpec {
            value: strings[1].clone(),
            font_family: strings[2].clone(),
            color: strings[3].clone(),
            role: strings[4].clone(),
            font_size: number(25),
            line_height: number(26),
            wrap: decode_optional_bool(booleans_present, booleans_value, BOOL_TEXT_WRAP),
            fit: decode_optional_bool(booleans_present, booleans_value, BOOL_TEXT_FIT),
            clip: decode_optional_bool(booleans_present, booleans_value, BOOL_TEXT_CLIP),
            padding: Insets {
                top: number(27),
                right: number(28),
                bottom: number(29),
                left: number(30),
            },
        },
        icon: IconSpec {
            reference: strings[5].clone(),
            fallback_reference: strings[6].clone(),
            color: strings[7].clone(),
            width: number(31),
            height: number(32),
            scale: number(33),
            offset_x: number(34),
            offset_y: number(35),
            missing_policy,
        },
        port: PortSpec {
            side,
            anchor: number(36),
            offset: number(37),
            size: number(38),
            visible: decode_optional_bool(booleans_present, booleans_value, BOOL_PORT_VISIBLE),
            label: strings[10].clone(),
        },
        line: LineSpec {
            source: strings[11].clone(),
            target: strings[12].clone(),
            source_side,
            target_side,
            source_anchor: number(39),
            target_anchor: number(40),
            routing,
            obstacle_margin: number(41),
            style: line_style,
            source_decoration,
            target_decoration,
            label: strings[13].clone(),
            label_position: number(42),
        },
    })
}

fn option_from_u8_flag(flags: u8, bit: u8, value: f64) -> Option<f64> {
    if flags & (1 << bit) != 0 {
        Some(value)
    } else {
        None
    }
}

fn option_from_u64_flag(flags: u64, bit: usize, value: f64) -> Option<f64> {
    if flags & (1 << bit) != 0 {
        Some(value)
    } else {
        None
    }
}

fn option_from_bool_flag<T>(flags: u16, bit: u16, value: T) -> Option<T> {
    if flags & bit != 0 {
        Some(value)
    } else {
        None
    }
}

fn decode_optional_bool(present: u16, values: u16, bit: u16) -> Option<bool> {
    if present & bit == 0 {
        None
    } else {
        Some(values & bit != 0)
    }
}

fn decode_concept(value: u8) -> Result<Concept, String> {
    match value {
        1 => Ok(Concept::Frame),
        2 => Ok(Concept::Group),
        3 => Ok(Concept::Capture),
        4 => Ok(Concept::Item),
        5 => Ok(Concept::Port),
        6 => Ok(Concept::Line),
        7 => Ok(Concept::Text),
        8 => Ok(Concept::Spacer),
        _ => Err(format!("unsupported engine concept {value}")),
    }
}

fn decode_layout(value: u8) -> Result<LayoutPolicy, String> {
    match value {
        0 => Ok(LayoutPolicy::Default),
        1 => Ok(LayoutPolicy::Vertical),
        2 => Ok(LayoutPolicy::Horizontal),
        3 => Ok(LayoutPolicy::Grid),
        4 => Ok(LayoutPolicy::Absolute),
        5 => Ok(LayoutPolicy::None),
        _ => Err(format!("unsupported layout policy {value}")),
    }
}

fn decode_overflow(value: u8) -> Result<Overflow, String> {
    match value {
        0 | 1 => Ok(Overflow::Error),
        2 => Ok(Overflow::Visible),
        _ => Err(format!("unsupported overflow policy {value}")),
    }
}

fn decode_shape(value: u8) -> Result<Shape, String> {
    match value {
        0 => Ok(Shape::Default),
        1 => Ok(Shape::Rectangle),
        2 => Ok(Shape::Ellipse),
        3 => Ok(Shape::None),
        _ => Err(format!("unsupported shape {value}")),
    }
}

fn decode_alignment(value: u8) -> Result<Alignment, String> {
    match value {
        0 | 4 => Ok(Alignment::Stretch),
        1 => Ok(Alignment::Start),
        2 => Ok(Alignment::Center),
        3 => Ok(Alignment::End),
        _ => Err(format!("unsupported alignment {value}")),
    }
}

fn decode_justification(value: u8) -> Result<xaligo_layout_engine::Justification, String> {
    use xaligo_layout_engine::Justification;
    match value {
        0 | 1 => Ok(Justification::Start),
        2 => Ok(Justification::Center),
        3 => Ok(Justification::End),
        4 => Ok(Justification::SpaceBetween),
        5 => Ok(Justification::SpaceEvenly),
        _ => Err(format!("unsupported justification {value}")),
    }
}

fn decode_side(value: u8) -> Result<Side, String> {
    match value {
        0 | 1 => Ok(Side::Auto),
        2 => Ok(Side::Top),
        3 => Ok(Side::Right),
        4 => Ok(Side::Bottom),
        5 => Ok(Side::Left),
        _ => Err(format!("unsupported side {value}")),
    }
}

fn decode_routing(value: u8) -> Result<RoutingPolicy, String> {
    match value {
        0 | 1 => Ok(RoutingPolicy::Orthogonal),
        2 => Ok(RoutingPolicy::Straight),
        _ => Err(format!("unsupported routing policy {value}")),
    }
}

fn decode_line_style(value: u8) -> Result<LineStyle, String> {
    match value {
        0 | 1 => Ok(LineStyle::Solid),
        2 => Ok(LineStyle::Dashed),
        3 => Ok(LineStyle::Dotted),
        _ => Err(format!("unsupported line style {value}")),
    }
}

fn decode_decoration(value: u8) -> Result<Decoration, String> {
    match value {
        0 | 1 => Ok(Decoration::None),
        2 => Ok(Decoration::Arrow),
        3 => Ok(Decoration::Triangle),
        4 => Ok(Decoration::Diamond),
        5 => Ok(Decoration::Circle),
        _ => Err(format!("unsupported decoration {value}")),
    }
}

fn decode_missing_icon(value: u8) -> Result<MissingIconPolicy, String> {
    match value {
        0 | 2 => Ok(MissingIconPolicy::Fallback),
        1 => Ok(MissingIconPolicy::Error),
        3 => Ok(MissingIconPolicy::Hide),
        _ => Err(format!("unsupported missing-icon policy {value}")),
    }
}

fn encode_success(response: EngineResponse) -> Vec<u8> {
    let mut output = response_header(STATUS_OK);
    match response {
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

fn encode_error(message: &str) -> Vec<u8> {
    let bytes = message.as_bytes();
    let mut output = response_header(STATUS_ERROR);
    output.push(0);
    output.extend_from_slice(&(bytes.len() as u32).to_le_bytes());
    output.extend_from_slice(bytes);
    output
}

fn response_header(status: u8) -> Vec<u8> {
    let mut output = Vec::with_capacity(16);
    output.extend_from_slice(RESPONSE_MAGIC);
    output.extend_from_slice(&ABI_VERSION.to_le_bytes());
    output.push(status);
    output
}

struct Decoder<'a> {
    input: &'a [u8],
    offset: usize,
}

impl<'a> Decoder<'a> {
    fn new(input: &'a [u8]) -> Self {
        Self { input, offset: 0 }
    }

    fn read_exact(&mut self, length: usize) -> Result<&'a [u8], String> {
        let end = self
            .offset
            .checked_add(length)
            .ok_or_else(|| "engine request length overflow".to_owned())?;
        if end > self.input.len() {
            return Err("truncated engine request".to_owned());
        }
        let value = &self.input[self.offset..end];
        self.offset = end;
        Ok(value)
    }

    fn read_u8(&mut self) -> Result<u8, String> {
        Ok(self.read_exact(1)?[0])
    }

    fn read_u16(&mut self) -> Result<u16, String> {
        let bytes: [u8; 2] = self
            .read_exact(2)?
            .try_into()
            .map_err(|_| "invalid u16".to_owned())?;
        Ok(u16::from_le_bytes(bytes))
    }

    fn read_u32(&mut self) -> Result<u32, String> {
        let bytes: [u8; 4] = self
            .read_exact(4)?
            .try_into()
            .map_err(|_| "invalid u32".to_owned())?;
        Ok(u32::from_le_bytes(bytes))
    }

    fn read_i32(&mut self) -> Result<i32, String> {
        let bytes: [u8; 4] = self
            .read_exact(4)?
            .try_into()
            .map_err(|_| "invalid i32".to_owned())?;
        Ok(i32::from_le_bytes(bytes))
    }

    fn read_u64(&mut self) -> Result<u64, String> {
        let bytes: [u8; 8] = self
            .read_exact(8)?
            .try_into()
            .map_err(|_| "invalid u64".to_owned())?;
        Ok(u64::from_le_bytes(bytes))
    }

    fn read_f64(&mut self) -> Result<f64, String> {
        let bytes: [u8; 8] = self
            .read_exact(8)?
            .try_into()
            .map_err(|_| "invalid f64".to_owned())?;
        Ok(f64::from_le_bytes(bytes))
    }

    fn is_empty(&self) -> bool {
        self.offset == self.input.len()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn document_request(operation: u8, width: f64, elements: Vec<Vec<u8>>) -> Vec<u8> {
        let mut input = Vec::new();
        input.extend_from_slice(REQUEST_MAGIC);
        input.extend_from_slice(&ABI_VERSION.to_le_bytes());
        input.push(operation);
        input.push(1);
        input.extend_from_slice(&width.to_le_bytes());
        input.extend_from_slice(&80.0f64.to_le_bytes());
        input.extend_from_slice(&4.0f64.to_le_bytes());
        input.push(0);
        input.push(0);
        input.extend_from_slice(&0u16.to_le_bytes());
        for _ in 0..4 {
            input.extend_from_slice(&0.0f64.to_le_bytes());
        }
        input.extend_from_slice(&(elements.len() as u32).to_le_bytes());
        for element in elements {
            input.extend(element);
        }
        input
    }

    fn element(id: &str, parent: i32, flags: u64, values: &[(usize, f64)]) -> Vec<u8> {
        let mut numbers = [0.0; NUMERIC_FIELD_COUNT];
        for (index, value) in values {
            numbers[*index] = *value;
        }
        let mut bytes = Vec::new();
        bytes.extend_from_slice(&parent.to_le_bytes());
        bytes.extend_from_slice(&flags.to_le_bytes());
        bytes.extend_from_slice(&0u16.to_le_bytes());
        bytes.extend_from_slice(&0u16.to_le_bytes());
        for _ in 0..4 {
            bytes.extend_from_slice(&0u16.to_le_bytes());
        }
        bytes.extend_from_slice(&[4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);
        bytes.extend_from_slice(&0i32.to_le_bytes());
        for number in numbers {
            bytes.extend_from_slice(&number.to_le_bytes());
        }
        bytes.extend_from_slice(&(id.len() as u16).to_le_bytes());
        for _ in 1..STRING_FIELD_COUNT {
            bytes.extend_from_slice(&0u16.to_le_bytes());
        }
        bytes.extend_from_slice(id.as_bytes());
        bytes
    }

    fn normalize_request(svg: &[u8]) -> Vec<u8> {
        let mut input = Vec::new();
        input.extend_from_slice(REQUEST_MAGIC);
        input.extend_from_slice(&ABI_VERSION.to_le_bytes());
        input.push(OPERATION_NORMALIZE_SVG);
        input.push(0);
        input.extend_from_slice(&(svg.len() as u32).to_le_bytes());
        input.extend_from_slice(svg);
        input
    }

    #[test]
    fn layout_and_svg_operations_share_resolution() {
        let fixed = element("fixed", -1, (1 << 2) | (1 << 3), &[(2, 40.0), (3, 20.0)]);
        let flex = element("flex", -1, (1 << 2) | (1 << 12), &[(2, 50.0), (12, 1.0)]);
        let layout = process_request(&document_request(
            OPERATION_LAYOUT,
            100.0,
            vec![fixed.clone(), flex.clone()],
        ));
        assert_eq!(&layout[0..4], RESPONSE_MAGIC);
        assert_eq!(layout[6], STATUS_OK);
        assert_eq!(layout[7], OPERATION_LAYOUT);

        let svg = process_request(&document_request(OPERATION_SVG, 100.0, vec![fixed, flex]));
        assert_eq!(svg[6], STATUS_OK);
        assert_eq!(svg[7], OPERATION_SVG);
        assert!(String::from_utf8_lossy(&svg).contains("<svg"));
        assert!(String::from_utf8_lossy(&svg).contains("id=\"flex\""));
    }

    #[test]
    fn malformed_unknown_and_non_finite_requests_return_typed_errors() {
        let truncated = process_request(b"XLE2");
        assert_eq!(truncated[6], STATUS_ERROR);
        assert!(String::from_utf8_lossy(&truncated).contains("truncated"));

        let mut unknown = document_request(
            OPERATION_LAYOUT,
            100.0,
            vec![element("item", -1, 1 << 3, &[(3, 20.0)])],
        );
        unknown[4..6].copy_from_slice(&99u16.to_le_bytes());
        let response = process_request(&unknown);
        assert_eq!(response[6], STATUS_ERROR);
        assert!(String::from_utf8_lossy(&response).contains("ABI version 99"));

        let non_finite = process_request(&document_request(
            OPERATION_LAYOUT,
            f64::NAN,
            vec![element("item", -1, 1 << 3, &[(3, 20.0)])],
        ));
        assert_eq!(non_finite[6], STATUS_ERROR);
        assert!(String::from_utf8_lossy(&non_finite).contains("document width"));
    }

    #[test]
    fn normalizes_svg_and_rejects_active_content_through_abi() {
        let safe = process_request(&normalize_request(
            br#"<svg xmlns="http://www.w3.org/2000/svg" width="8" height="4"><path d="M0 0h8v4z"/></svg>"#,
        ));
        assert_eq!(safe[6], STATUS_OK);
        assert_eq!(safe[7], OPERATION_NORMALIZE_SVG);
        assert!(String::from_utf8_lossy(&safe).contains("<svg"));

        let unsafe_svg = process_request(&normalize_request(
            br#"<svg xmlns="http://www.w3.org/2000/svg"><script/></svg>"#,
        ));
        assert_eq!(unsafe_svg[6], STATUS_ERROR);
        assert!(String::from_utf8_lossy(&unsafe_svg).contains("forbidden"));
    }

    #[test]
    fn c_abi_returns_owned_buffer_and_rejects_invalid_pointers() {
        assert_eq!(xaligo_engine_abi_version(), ABI_VERSION as u32);
        assert_eq!(
            unsafe { xaligo_engine_process(ptr::null(), 1, ptr::null_mut()) },
            FFI_NULL_OUTPUT
        );

        let request = document_request(
            OPERATION_LAYOUT,
            100.0,
            vec![element("item", -1, 1 << 3, &[(3, 20.0)])],
        );
        let mut output = XaligoEngineBuffer::empty();
        let status = unsafe {
            xaligo_engine_process(request.as_ptr(), request.len(), &mut output as *mut _)
        };
        assert_eq!(status, FFI_OK);
        assert!(!output.data.is_null());
        assert!(output.len > 8);
        let response = unsafe { std::slice::from_raw_parts(output.data, output.len) };
        assert_eq!(&response[0..4], RESPONSE_MAGIC);
        unsafe { xaligo_engine_buffer_free(output) };
    }
}
