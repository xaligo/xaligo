#[rustfmt::skip]
use std::panic::{
    AssertUnwindSafe,
    catch_unwind,
};
use std::ptr;
use std::sync::atomic::{
    AtomicU8,
    Ordering,
};

use crate::base::process_request;
#[rustfmt::skip]
use crate::cnf::engine::{
    ABI_VERSION,
    FFI_NULL_INPUT,
    FFI_NULL_OUTPUT,
    FFI_OK,
    FFI_PANIC,
};

#[repr(C)]
pub struct XaligoEngineBuffer {
    pub data: *mut u8,
    pub len: usize,
    pub capacity: usize,
}

#[repr(C)]
pub struct XaligoEngineCancel {
    cancelled: AtomicU8,
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
    unsafe { xaligo_engine_process_with_cancel(input, input_len, ptr::null(), output) }
}

#[unsafe(no_mangle)]
pub extern "C" fn xaligo_engine_cancel_new() -> *mut XaligoEngineCancel {
    Box::into_raw(Box::new(XaligoEngineCancel {
        cancelled: AtomicU8::new(0),
    }))
}

/// Marks a cancellation handle as cancelled.
///
/// # Safety
///
/// `cancel` must be null or an allocation returned by
/// `xaligo_engine_cancel_new` that has not been freed.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn xaligo_engine_cancel_set(cancel: *mut XaligoEngineCancel) {
    if let Some(cancel) = unsafe { cancel.as_ref() } {
        cancel.cancelled.store(1, Ordering::Relaxed);
    }
}

/// Releases a cancellation handle.
///
/// # Safety
///
/// `cancel` must be null or an allocation returned by
/// `xaligo_engine_cancel_new`, and it must be released exactly once after all
/// concurrent process/cancel calls using it have completed.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn xaligo_engine_cancel_free(cancel: *mut XaligoEngineCancel) {
    if !cancel.is_null() {
        drop(unsafe { Box::from_raw(cancel) });
    }
}

/// Runs one engine operation with optional cooperative cancellation.
///
/// # Safety
///
/// The input and output requirements match `xaligo_engine_process`. `cancel`
/// must be null or point to a live `XaligoEngineCancel` for the entire call.
#[unsafe(no_mangle)]
pub unsafe extern "C" fn xaligo_engine_process_with_cancel(
    input: *const u8,
    input_len: usize,
    cancel: *const XaligoEngineCancel,
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
        let token = if cancel.is_null() {
            ptr::null()
        } else {
            unsafe { &(*cancel).cancelled as *const AtomicU8 }
        };
        crate::usc::cancel::with_token(token, || process_request(request))
    }));
    crate::usc::cancel::clear();

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
