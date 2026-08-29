use std::collections::BTreeMap;
use std::env;
use std::fs::{
    File,
    OpenOptions,
};
use std::io::{
    self,
    Write,
};
use std::panic::Location;
use std::path::Path;
use std::sync::{
    Mutex,
    OnceLock,
};
use std::time::{
    SystemTime,
    UNIX_EPOCH,
};

use crate::util::mcode::{
    LogLevel,
    MCode,
    MLOG_OUTPUT_FALLBACK,
    format_message,
};

pub type LogFields = BTreeMap<String, String>;

pub struct LoggerConfig {
    pub component: String,
    pub service: String,
    pub level: String,
    pub structured: bool,
    pub enable_caller: bool,
    pub output: String,
}

impl LoggerConfig {
    pub fn from_env(component: impl Into<String>, service: impl Into<String>) -> Self {
        Self {
            component: component.into(),
            service: service.into(),
            level: env::var("XALIGO_LOG_LEVEL").unwrap_or_default(),
            structured: env::var("XALIGO_LOG_STRUCTURED")
                .is_ok_and(|value| truthy(&value)),
            enable_caller: env::var("XALIGO_LOG_CALLER").is_ok_and(|value| truthy(&value)),
            output: env::var("XALIGO_LOG_OUTPUT").unwrap_or_default(),
        }
    }
}

pub struct Logger {
    config: LoggerConfig,
    level: LogLevel,
    output: Mutex<LogOutput>,
}

impl Logger {
    pub fn new(config: LoggerConfig) -> Self {
        let level = parse_log_level(&config.level);
        let output = open_log_output(&config.output);
        Self {
            config,
            level,
            output: Mutex::new(output),
        }
    }

    pub fn enabled(&self, level: LogLevel) -> bool {
        level.rank() >= self.level.rank()
    }

    #[track_caller]
    pub fn debug(&self, mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
        self.log(LogLevel::Debug, mcode, optional_message, fields);
    }

    #[track_caller]
    pub fn info(&self, mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
        self.log(LogLevel::Info, mcode, optional_message, fields);
    }

    #[track_caller]
    pub fn warn(&self, mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
        self.log(LogLevel::Warn, mcode, optional_message, fields);
    }

    #[track_caller]
    pub fn error(&self, mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
        self.log(LogLevel::Error, mcode, optional_message, fields);
    }

    #[track_caller]
    pub fn fatal(&self, mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
        self.log(LogLevel::Fatal, mcode, optional_message, fields);
    }

    #[track_caller]
    fn log(
        &self,
        level: LogLevel,
        mcode: MCode,
        optional_message: &str,
        fields: Option<&LogFields>,
    ) {
        if !self.enabled(level) {
            return;
        }
        let mut fields = fields.cloned().unwrap_or_default();
        let error = fields.remove("error");
        let caller = if self.config.enable_caller || level == LogLevel::Debug {
            let location = Location::caller();
            Some(LogCaller {
                file: Path::new(location.file())
                    .file_name()
                    .and_then(|value| value.to_str())
                    .unwrap_or("unknown"),
                function: module_path!(),
                line: location.line(),
            })
        } else {
            None
        };
        let entry = LogEntry {
            timestamp: utc_timestamp(),
            level,
            code: mcode.code,
            component: &self.config.component,
            service: &self.config.service,
            message: mcode.message_with(optional_message),
            fields: &fields,
            caller,
            error: error.as_deref(),
        };
        let line = if self.config.structured {
            render_structured(&entry)
        } else {
            render_text(&entry)
        };
        let mut output = self.output.lock().unwrap_or_else(|poisoned| poisoned.into_inner());
        let _ = writeln!(output, "{line}");
    }
}

static DEFAULT_LOGGER: OnceLock<Logger> = OnceLock::new();

pub fn default_logger() -> &'static Logger {
    DEFAULT_LOGGER.get_or_init(|| Logger::new(LoggerConfig::from_env("engine", "v2")))
}

#[track_caller]
pub fn debug(mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
    default_logger().debug(mcode, optional_message, fields);
}

#[track_caller]
pub fn info(mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
    default_logger().info(mcode, optional_message, fields);
}

#[track_caller]
pub fn warn(mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
    default_logger().warn(mcode, optional_message, fields);
}

#[track_caller]
pub fn error(mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
    default_logger().error(mcode, optional_message, fields);
}

#[track_caller]
pub fn fatal(mcode: MCode, optional_message: &str, fields: Option<&LogFields>) {
    default_logger().fatal(mcode, optional_message, fields);
}

struct LogEntry<'a> {
    timestamp: String,
    level: LogLevel,
    code: &'a str,
    component: &'a str,
    service: &'a str,
    message: String,
    fields: &'a LogFields,
    caller: Option<LogCaller<'a>>,
    error: Option<&'a str>,
}

struct LogCaller<'a> {
    file: &'a str,
    function: &'a str,
    line: u32,
}

enum LogOutput {
    Stdout(io::Stdout),
    Stderr(io::Stderr),
    File(File),
}

impl Write for LogOutput {
    fn write(&mut self, buffer: &[u8]) -> io::Result<usize> {
        match self {
            Self::Stdout(output) => output.write(buffer),
            Self::Stderr(output) => output.write(buffer),
            Self::File(output) => output.write(buffer),
        }
    }

    fn flush(&mut self) -> io::Result<()> {
        match self {
            Self::Stdout(output) => output.flush(),
            Self::Stderr(output) => output.flush(),
            Self::File(output) => output.flush(),
        }
    }
}

fn truthy(value: &str) -> bool {
    matches!(
        value.trim().to_ascii_lowercase().as_str(),
        "1" | "true" | "yes" | "y" | "on"
    )
}

fn parse_log_level(level: &str) -> LogLevel {
    match level.trim().to_ascii_uppercase().as_str() {
        "DEBUG" => LogLevel::Debug,
        "WARN" | "WARNING" => LogLevel::Warn,
        "ERROR" => LogLevel::Error,
        "FATAL" => LogLevel::Fatal,
        _ => LogLevel::Info,
    }
}

fn open_log_output(output: &str) -> LogOutput {
    match output.trim() {
        "stdout" => LogOutput::Stdout(io::stdout()),
        "" | "stderr" => LogOutput::Stderr(io::stderr()),
        path => match OpenOptions::new().create(true).append(true).open(path) {
            Ok(file) => LogOutput::File(file),
            Err(_) => {
                eprintln!(
                    "{}",
                    format_message(LogLevel::Warn, MLOG_OUTPUT_FALLBACK, "")
                );
                LogOutput::Stderr(io::stderr())
            }
        },
    }
}

fn render_text(entry: &LogEntry<'_>) -> String {
    let mut output = format!(
        "[{}] [{}] [{}] {}",
        entry.timestamp,
        entry.level.as_str(),
        entry.code,
        entry.message
    );
    if entry.level == LogLevel::Debug && !entry.fields.is_empty() {
        output.push(' ');
        output.push_str(&render_fields(entry.fields));
    }
    if let Some(error) = entry.error {
        output.push_str(" error=\"");
        output.push_str(&escape_text(error));
        output.push('"');
    }
    output
}

fn render_structured(entry: &LogEntry<'_>) -> String {
    let mut output = String::from("{");
    let mut first = true;
    push_json_string(&mut output, &mut first, "timestamp", &entry.timestamp);
    push_json_string(&mut output, &mut first, "level", entry.level.as_str());
    push_json_string(&mut output, &mut first, "code", entry.code);
    if !entry.component.is_empty() {
        push_json_string(&mut output, &mut first, "component", entry.component);
    }
    if !entry.service.is_empty() {
        push_json_string(&mut output, &mut first, "service", entry.service);
    }
    push_json_string(&mut output, &mut first, "message", &entry.message);
    if !entry.fields.is_empty() {
        push_json_key(&mut output, &mut first, "fields");
        output.push_str(&render_fields(entry.fields));
    }
    if let Some(caller) = &entry.caller {
        push_json_string(&mut output, &mut first, "file", caller.file);
        push_json_string(&mut output, &mut first, "function", caller.function);
        push_json_key(&mut output, &mut first, "line");
        output.push_str(&caller.line.to_string());
    }
    if let Some(error) = entry.error {
        push_json_string(&mut output, &mut first, "error", error);
    }
    output.push('}');
    output
}

fn render_fields(fields: &LogFields) -> String {
    let mut output = String::from("{");
    let mut first = true;
    for (key, value) in fields {
        push_json_string(&mut output, &mut first, key, value);
    }
    output.push('}');
    output
}

fn push_json_key(output: &mut String, first: &mut bool, key: &str) {
    if !*first {
        output.push(',');
    }
    *first = false;
    output.push('"');
    output.push_str(&escape_json(key));
    output.push_str("\":");
}

fn push_json_string(output: &mut String, first: &mut bool, key: &str, value: &str) {
    push_json_key(output, first, key);
    output.push('"');
    output.push_str(&escape_json(value));
    output.push('"');
}

fn escape_json(value: &str) -> String {
    let mut escaped = String::with_capacity(value.len());
    for character in value.chars() {
        match character {
            '"' => escaped.push_str("\\\""),
            '\\' => escaped.push_str("\\\\"),
            '\n' => escaped.push_str("\\n"),
            '\r' => escaped.push_str("\\r"),
            '\t' => escaped.push_str("\\t"),
            value if value <= '\u{001f}' => {
                escaped.push_str(&format!("\\u{:04x}", value as u32));
            }
            value => escaped.push(value),
        }
    }
    escaped
}

fn escape_text(value: &str) -> String {
    value.replace('\\', "\\\\").replace('"', "\\\"")
}

fn utc_timestamp() -> String {
    let duration = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap_or_default();
    let seconds = duration.as_secs() as i64;
    let days = seconds / 86_400;
    let seconds_of_day = seconds % 86_400;
    let (year, month, day) = civil_from_days(days);
    let hour = seconds_of_day / 3_600;
    let minute = (seconds_of_day % 3_600) / 60;
    let second = seconds_of_day % 60;
    format!(
        "{year:04}-{month:02}-{day:02}T{hour:02}:{minute:02}:{second:02}.{:09}Z",
        duration.subsec_nanos()
    )
}

fn civil_from_days(days_since_epoch: i64) -> (i64, i64, i64) {
    let shifted = days_since_epoch + 719_468;
    let era = if shifted >= 0 {
        shifted
    } else {
        shifted - 146_096
    } / 146_097;
    let day_of_era = shifted - era * 146_097;
    let year_of_era =
        (day_of_era - day_of_era / 1_460 + day_of_era / 36_524 - day_of_era / 146_096)
            / 365;
    let mut year = year_of_era + era * 400;
    let day_of_year = day_of_era - (365 * year_of_era + year_of_era / 4 - year_of_era / 100);
    let month_prime = (5 * day_of_year + 2) / 153;
    let day = day_of_year - (153 * month_prime + 2) / 5 + 1;
    let month = month_prime + if month_prime < 10 { 3 } else { -9 };
    if month <= 2 {
        year += 1;
    }
    (year, month, day)
}
