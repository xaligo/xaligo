pub enum LogLevel {
    Debug,
    Info,
    Warn,
    Error,
    Fatal,
}

impl LogLevel {
    pub const fn as_str(&self) -> &'static str {
        match self {
            Self::Debug => "DEBUG",
            Self::Info => "INFO",
            Self::Warn => "WARN",
            Self::Error => "ERROR",
            Self::Fatal => "FATAL",
        }
    }

    pub(crate) const fn rank(&self) -> u8 {
        match self {
            Self::Debug => 0,
            Self::Info => 1,
            Self::Warn => 2,
            Self::Error => 3,
            Self::Fatal => 4,
        }
    }
}

pub struct MCode {
    pub code: &'static str,
    pub message: &'static str,
}

impl MCode {
    pub const fn new(code: &'static str, message: &'static str) -> Self {
        Self { code, message }
    }

    pub fn message_with(&self, optional_message: &str) -> String {
        match (self.message.is_empty(), optional_message.is_empty()) {
            (false, false) => format!("{}: {optional_message}", self.message),
            (false, true) => self.message.to_owned(),
            (true, false) => optional_message.to_owned(),
            (true, true) => String::new(),
        }
    }
}

pub const MLOG_OUTPUT_FALLBACK: MCode =
    MCode::new("MLOG2", "Logger output fallback to stderr");
pub const MENGINE_PROCESS: MCode = MCode::new("MENG1", "Engine request processing started");
pub const MENGINE_COMPLETE: MCode = MCode::new("MENG2", "Engine request processing completed");

pub fn format_message(level: LogLevel, mcode: MCode, optional_message: &str) -> String {
    format!(
        "[{}][{}] {}",
        level.as_str(),
        mcode.code,
        mcode.message_with(optional_message)
    )
}
