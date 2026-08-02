use std::panic::{AssertUnwindSafe, catch_unwind};
use std::ptr;

use crate::base::process_request;
use crate::cnf::engine::{
    ABI_VERSION, FFI_NULL_INPUT, FFI_NULL_OUTPUT, FFI_OK, FFI_PANIC,
};

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
    unsafe { ptr::write(output, XaligoEngineBuffer::empty()) };
    if input.is_null() && input_len != 0 {
        return FFI_NULL_INPUT;
    }

    let result = catch_unwind(AssertUnwindSafe(|| {
        let request = if input_len == 0 {
            &[]
        } else {
            unsafe { std::slice::from_raw_parts(input, input_len) }
        };
        process_request(request)
    }));

    match result {
        Ok(response) => {
            unsafe { ptr::write(output, XaligoEngineBuffer::from_vec(response)) };
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
    drop(unsafe { Vec::from_raw_parts(buffer.data, buffer.len, buffer.capacity) });
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::cnf::engine::{
        NUMERIC_FIELD_COUNT, OPERATION_LAYOUT, OPERATION_NORMALIZE_SVG, OPERATION_SVG,
        REQUEST_MAGIC, RESPONSE_MAGIC, STATUS_ERROR, STATUS_OK, STRING_FIELD_COUNT,
    };

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
