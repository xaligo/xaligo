include!("mod.rs");
mod base;

#[rustfmt::skip]
pub use ctl::engine::{
    XaligoEngineBuffer,
    xaligo_engine_abi_version,
    xaligo_engine_buffer_free,
    xaligo_engine_process,
};

// Keep the PPTX C ABI in the same Rust archive as the engine so the Go binary
// links exactly one Rust runtime.
pub use xaligo_pptx_exporter::{
    XaligoExporterBuffer, xaligo_exporter_abi_version, xaligo_exporter_buffer_free,
    xaligo_exporter_process,
};
