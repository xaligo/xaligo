use super::*;

fn entry<'a>(fields: &'a LogFields) -> LogEntry<'a> {
    LogEntry {
        timestamp: "2026-08-02T00:00:00.000000000Z".to_owned(),
        level: LogLevel::Debug,
        code: "TEST1",
        component: "engine",
        service: "v2",
        message: "message".to_owned(),
        fields,
        caller: Some(LogCaller {
            file: "logger.rs",
            function: "tests",
            line: 1,
        }),
        error: Some("boom"),
    }
}

#[test]
fn structured_output_matches_go_logger_field_contract() {
    let fields = LogFields::from([("answer".to_owned(), "42".to_owned())]);
    let output = render_structured(&entry(&fields));
    assert!(output.contains("\"level\":\"DEBUG\""));
    assert!(output.contains("\"code\":\"TEST1\""));
    assert!(output.contains("\"fields\":{\"answer\":\"42\"}"));
    assert!(output.contains("\"error\":\"boom\""));
}

#[test]
fn text_output_keeps_message_code_and_debug_fields() {
    let fields = LogFields::from([("answer".to_owned(), "42".to_owned())]);
    let output = render_text(&entry(&fields));
    assert!(output.contains("[DEBUG] [TEST1] message"));
    assert!(output.contains("{\"answer\":\"42\"}"));
    assert!(output.contains("error=\"boom\""));
}

#[test]
fn parses_go_compatible_environment_values() {
    assert!(truthy(" YES "));
    assert_eq!(parse_log_level("warning"), LogLevel::Warn);
    assert_eq!(parse_log_level("unknown"), LogLevel::Info);
    assert_eq!(civil_from_days(0), (1970, 1, 1));
}
