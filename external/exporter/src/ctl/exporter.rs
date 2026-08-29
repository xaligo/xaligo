use std::panic::{AssertUnwindSafe, catch_unwind};
use std::ptr;

use crate::ent::request::exporter::ExporterRequest;
use crate::usc::exporter;
use crate::util::error::Error;

pub fn export(input: &str) -> Result<Vec<u8>, Error> {
    if input.trim().is_empty() {
        return Err(Error::invalid("PPTX exporter request JSON is required on stdin"));
    }
    let request: ExporterRequest = serde_json::from_str(input)
        .map_err(|error| Error::invalid(format!("parse PPTX exporter request: {error}")))?;
    exporter::export(&request)
}

#[repr(C)]
pub struct XaligoExporterBuffer {
    pub data: *mut u8,
    pub len: usize,
    pub capacity: usize,
}

impl XaligoExporterBuffer {
    const fn empty() -> Self {
        Self { data: ptr::null_mut(), len: 0, capacity: 0 }
    }

    fn from_vec(mut value: Vec<u8>) -> Self {
        let buffer = Self { data: value.as_mut_ptr(), len: value.len(), capacity: value.capacity() };
        std::mem::forget(value);
        buffer
    }
}

#[unsafe(no_mangle)]
pub extern "C" fn xaligo_exporter_abi_version() -> u32 { 1 }

/// # Safety
/// `output` must be writable. A non-empty input must point to `input_len`
/// readable bytes. Successful output must be freed exactly once.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn xaligo_exporter_process(
    input: *const u8,
    input_len: usize,
    output: *mut XaligoExporterBuffer,
) -> i32 {
    if output.is_null() { return 2; }
    unsafe { ptr::write(output, XaligoExporterBuffer::empty()) };
    if input.is_null() && input_len != 0 { return 1; }
    let result = catch_unwind(AssertUnwindSafe(|| {
        let bytes = if input_len == 0 { &[] } else { unsafe { std::slice::from_raw_parts(input, input_len) } };
        let source = std::str::from_utf8(bytes).map_err(|error| Error::invalid(format!("PPTX exporter request is not UTF-8: {error}")))?;
        export(source)
    }));
    match result {
        Ok(Ok(bytes)) => { unsafe { ptr::write(output, XaligoExporterBuffer::from_vec(bytes)) }; 0 }
        Ok(Err(_)) => 3,
        Err(_) => 4,
    }
}

/// # Safety
/// `buffer` must be empty or an unchanged successful process result.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn xaligo_exporter_buffer_free(buffer: XaligoExporterBuffer) {
    if !buffer.data.is_null() {
        drop(unsafe { Vec::from_raw_parts(buffer.data, buffer.len, buffer.capacity) });
    }
}
