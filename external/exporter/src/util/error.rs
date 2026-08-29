use std::fmt;

#[derive(Debug)]
pub struct Error(String);

impl Error {
    pub fn invalid(message: impl Into<String>) -> Self {
        Self(message.into())
    }
}

impl fmt::Display for Error {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        formatter.write_str(&self.0)
    }
}

impl std::error::Error for Error {}

impl From<pptx::PptxError> for Error {
    fn from(error: pptx::PptxError) -> Self {
        Self(format!("create PPTX package: {error}"))
    }
}
