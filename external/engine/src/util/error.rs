use std::error::Error;
#[rustfmt::skip]
use std::fmt::{
    Display,
    Formatter,
};

pub struct LayoutError {
    pub(crate) message: String,
}

impl LayoutError {
    pub(crate) fn new(message: impl Into<String>) -> Self {
        Self {
            message: message.into(),
        }
    }
}

impl Display for LayoutError {
    fn fmt(&self, formatter: &mut Formatter<'_>) -> std::fmt::Result {
        formatter.write_str(&self.message)
    }
}

impl Error for LayoutError {}

pub struct SvgError {
    pub(crate) message: String,
}

impl SvgError {
    pub(crate) fn new(message: impl Into<String>) -> Self {
        Self {
            message: message.into(),
        }
    }
}

impl Display for SvgError {
    fn fmt(&self, formatter: &mut Formatter<'_>) -> std::fmt::Result {
        formatter.write_str(&self.message)
    }
}

impl Error for SvgError {}
