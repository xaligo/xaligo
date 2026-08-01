use std::collections::HashSet;
use std::panic::{catch_unwind, AssertUnwindSafe};
use std::ptr;

use xaligo_layout_engine::{resolve, Direction, DocumentSpec, ElementSpec, ResolvedDocument};

pub const ABI_VERSION: u16 = 1;

const REQUEST_MAGIC: &[u8; 4] = b"XLE2";
const RESPONSE_MAGIC: &[u8; 4] = b"XLR2";
const OPERATION_LAYOUT: u8 = 1;
const OPERATION_SVG: u8 = 2;
const OPERATION_NORMALIZE_SVG: u8 = 3;
const STATUS_OK: u8 = 0;
const STATUS_ERROR: u8 = 1;
const MAX_REQUEST_BYTES: usize = 16 * 1024 * 1024;

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
    let direction = match discriminator {
        1 => Direction::Vertical,
        2 => Direction::Horizontal,
        value => return Err(format!("unsupported layout direction {value}")),
    };
    let width = decoder.read_f64()?;
    let height = decoder.read_f64()?;
    let gap = decoder.read_f64()?;
    let count = decoder.read_u32()? as usize;
    if count > 10_000 {
        return Err(format!("engine element count {count} exceeds 10000"));
    }

    let mut identifiers = HashSet::with_capacity(count);
    let mut elements = Vec::with_capacity(count);
    for _ in 0..count {
        let id_length = decoder.read_u16()? as usize;
        let flags = decoder.read_u8()?;
        let reserved = decoder.read_u8()?;
        if reserved != 0 || flags & !0b0000_0111 != 0 {
            return Err("invalid engine element flags".to_owned());
        }
        let width_value = decoder.read_f64()?;
        let height_value = decoder.read_f64()?;
        let weight_value = decoder.read_f64()?;
        let id_bytes = decoder.read_exact(id_length)?;
        let id = std::str::from_utf8(id_bytes)
            .map_err(|_| "engine element id is not valid UTF-8".to_owned())?
            .to_owned();
        if !identifiers.insert(id.clone()) {
            return Err(format!("duplicate element id {id:?}"));
        }
        elements.push(ElementSpec {
            id,
            width: option_from_flag(flags, 0, width_value),
            height: option_from_flag(flags, 1, height_value),
            weight: option_from_flag(flags, 2, weight_value),
        });
    }
    if !decoder.is_empty() {
        return Err("engine request has trailing bytes".to_owned());
    }
    Ok(EngineRequest::Document {
        operation,
        document: DocumentSpec {
            direction,
            width,
            height,
            gap,
            elements,
        },
    })
}

fn option_from_flag(flags: u8, bit: u8, value: f64) -> Option<f64> {
    if flags & (1 << bit) != 0 {
        Some(value)
    } else {
        None
    }
}

fn encode_success(response: EngineResponse) -> Vec<u8> {
    let mut output = response_header(STATUS_OK);
    match response {
        EngineResponse::Layout(document) => {
            output.push(OPERATION_LAYOUT);
            output.extend_from_slice(&(document.elements.len() as u32).to_le_bytes());
            for element in document.elements {
                let id = element.id.as_bytes();
                output.extend_from_slice(&(id.len() as u16).to_le_bytes());
                output.extend_from_slice(&0u16.to_le_bytes());
                output.extend_from_slice(&element.x.to_le_bytes());
                output.extend_from_slice(&element.y.to_le_bytes());
                output.extend_from_slice(&element.width.to_le_bytes());
                output.extend_from_slice(&element.height.to_le_bytes());
                output.extend_from_slice(id);
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

    fn request(operation: u8, direction: u8, width: f64, height: f64) -> Vec<u8> {
        let mut input = Vec::new();
        input.extend_from_slice(REQUEST_MAGIC);
        input.extend_from_slice(&ABI_VERSION.to_le_bytes());
        input.push(operation);
        input.push(direction);
        input.extend_from_slice(&width.to_le_bytes());
        input.extend_from_slice(&height.to_le_bytes());
        input.extend_from_slice(&4.0f64.to_le_bytes());
        input.extend_from_slice(&2u32.to_le_bytes());
        input.extend(element("fixed", 0b0000_0011, 40.0, 20.0, 0.0));
        input.extend(element("flex", 0b0000_0101, 50.0, 0.0, 1.0));
        input
    }

    fn element(id: &str, flags: u8, width: f64, height: f64, weight: f64) -> Vec<u8> {
        let mut bytes = Vec::new();
        bytes.extend_from_slice(&(id.len() as u16).to_le_bytes());
        bytes.push(flags);
        bytes.push(0);
        bytes.extend_from_slice(&width.to_le_bytes());
        bytes.extend_from_slice(&height.to_le_bytes());
        bytes.extend_from_slice(&weight.to_le_bytes());
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
        let layout = process_request(&request(OPERATION_LAYOUT, 1, 100.0, 80.0));
        assert_eq!(&layout[0..4], RESPONSE_MAGIC);
        assert_eq!(layout[6], STATUS_OK);
        assert_eq!(layout[7], OPERATION_LAYOUT);

        let svg = process_request(&request(OPERATION_SVG, 1, 100.0, 80.0));
        assert_eq!(svg[6], STATUS_OK);
        assert_eq!(svg[7], OPERATION_SVG);
        assert!(String::from_utf8_lossy(&svg).contains("<svg"));
        assert!(String::from_utf8_lossy(&svg).contains("id=\"flex\""));
    }

    #[test]
    fn malformed_and_unknown_abi_requests_return_typed_errors() {
        let truncated = process_request(b"XLE2");
        assert_eq!(truncated[6], STATUS_ERROR);
        assert!(String::from_utf8_lossy(&truncated).contains("truncated"));

        let mut unknown = request(OPERATION_LAYOUT, 1, 100.0, 80.0);
        unknown[4..6].copy_from_slice(&99u16.to_le_bytes());
        let response = process_request(&unknown);
        assert_eq!(response[6], STATUS_ERROR);
        assert!(String::from_utf8_lossy(&response).contains("ABI version 99"));
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

        let request = request(OPERATION_LAYOUT, 1, 100.0, 80.0);
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
