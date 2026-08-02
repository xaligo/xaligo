pub const ABI_VERSION: u16 = 2;

pub(crate) const REQUEST_MAGIC: &[u8; 4] = b"XLE2";
pub(crate) const RESPONSE_MAGIC: &[u8; 4] = b"XLR2";

pub(crate) const OPERATION_LAYOUT: u8 = 1;
pub(crate) const OPERATION_SVG: u8 = 2;
pub(crate) const OPERATION_NORMALIZE_SVG: u8 = 3;
pub(crate) const STATUS_OK: u8 = 0;
pub(crate) const STATUS_ERROR: u8 = 1;

pub(crate) const MAX_REQUEST_BYTES: usize = 16 * 1024 * 1024;
pub(crate) const MAX_SVG_BYTES: usize = 2 * 1024 * 1024;
pub(crate) const MAX_ELEMENTS: usize = 10_000;
pub(crate) const MAX_ID_BYTES: usize = u16::MAX as usize;
pub(crate) const MAX_COLUMNS: u16 = 256;
pub(crate) const MAX_DEPTH: usize = 128;

pub(crate) const DEFAULT_GAP: f64 = 16.0;
pub(crate) const DEFAULT_ITEM_SIZE: f64 = 48.0;
pub(crate) const DEFAULT_PORT_SIZE: f64 = 8.0;
pub(crate) const DEFAULT_FONT_SIZE: f64 = 14.0;
pub(crate) const DEFAULT_LINE_HEIGHT: f64 = 1.2;

pub(crate) const NUMERIC_FIELD_COUNT: usize = 43;
pub(crate) const STRING_FIELD_COUNT: usize = 14;

pub(crate) const BOOL_VISIBLE: u16 = 1 << 0;
pub(crate) const BOOL_TEXT_WRAP: u16 = 1 << 1;
pub(crate) const BOOL_TEXT_FIT: u16 = 1 << 2;
pub(crate) const BOOL_TEXT_CLIP: u16 = 1 << 3;
pub(crate) const BOOL_PORT_VISIBLE: u16 = 1 << 4;
pub(crate) const BOOL_LAYER: u16 = 1 << 5;
pub(crate) const BOOL_COLUMNS: u16 = 1 << 6;
pub(crate) const BOOL_COLUMN_SPAN: u16 = 1 << 7;
pub(crate) const BOOL_ROW_SPAN: u16 = 1 << 8;
pub(crate) const BOOL_KNOWN: u16 = BOOL_VISIBLE
    | BOOL_TEXT_WRAP
    | BOOL_TEXT_FIT
    | BOOL_TEXT_CLIP
    | BOOL_PORT_VISIBLE
    | BOOL_LAYER
    | BOOL_COLUMNS
    | BOOL_COLUMN_SPAN
    | BOOL_ROW_SPAN;

pub(crate) const FFI_OK: i32 = 0;
pub(crate) const FFI_NULL_OUTPUT: i32 = 1;
pub(crate) const FFI_NULL_INPUT: i32 = 2;
pub(crate) const FFI_PANIC: i32 = 3;
