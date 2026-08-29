use super::*;

#[test]
fn formats_message_code_with_optional_detail() {
    let code = MCode::new("TEST1", "base");
    assert_eq!(
        format_message(LogLevel::Warn, code, "detail"),
        "[WARN][TEST1] base: detail"
    );
}
