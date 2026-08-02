use crate::cnf::engine::{
    ABI_VERSION,
    BOOL_COLUMN_SPAN,
    BOOL_COLUMNS,
    BOOL_KNOWN,
    BOOL_LAYER,
    BOOL_PORT_VISIBLE,
    BOOL_ROW_SPAN,
    BOOL_TEXT_CLIP,
    BOOL_TEXT_FIT,
    BOOL_TEXT_WRAP,
    BOOL_VISIBLE,
    MAX_ELEMENTS,
    MAX_REQUEST_BYTES,
    NUMERIC_FIELD_COUNT,
    OPERATION_LAYOUT,
    OPERATION_NORMALIZE_SVG,
    OPERATION_SVG,
    REQUEST_MAGIC,
    STRING_FIELD_COUNT,
};
use crate::ent::model::document::{
    Alignment,
    Concept,
    Decoration,
    DocumentSpec,
    ElementSpec,
    IconSpec,
    Insets,
    Justification,
    LayoutPolicy,
    LineSpec,
    LineStyle,
    MissingIconPolicy,
    Overflow,
    PortSpec,
    RoutingPolicy,
    Shape,
    Side,
    TextSpec,
    VisualSpec,
};
use crate::ent::request::engine::EngineRequest;

pub(crate) trait Deserialize: Sized {
    fn deserialize(input: &[u8]) -> Result<Self, String>;
}

impl Deserialize for EngineRequest {
    fn deserialize(input: &[u8]) -> Result<Self, String> {
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

fn decode_element(decoder: &mut Decoder<'_>, index: usize) -> Result<ElementSpec, String> {
    let parent_value = decoder.read_i32()?;
    let parent = match parent_value {
        -1 => None,
        value if value >= 0 && (value as usize) < index => Some(value as usize),
        value => {
            return Err(format!(
                "invalid parent index {value} for engine element {index}"
            ));
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

fn decode_justification(value: u8) -> Result<Justification, String> {
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
