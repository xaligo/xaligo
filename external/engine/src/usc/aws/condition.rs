use super::detail::{children, feature};
use crate::ent::model::aws::{condition::*, feature::FeatureKind};
use crate::ent::model::document::DocumentSpec;
use crate::util::error::LayoutError;

pub(super) fn label(kind: &ConditionKind) -> &'static str {
    match kind {
        ConditionKind::Host => "Host",
        ConditionKind::Path => "Path",
        ConditionKind::Header => "Header",
        ConditionKind::Method => "Method",
        ConditionKind::Query => "Query",
        ConditionKind::SourceIp => "Source IP",
    }
}
pub(super) fn validate(
    doc: &DocumentSpec,
    index: usize,
    model: &Condition,
) -> Result<(), LayoutError> {
    let matches = children(doc, index);
    if matches.is_empty() || matches.len() > 3 {
        return Err(LayoutError::new(
            "ALB conditions require 1..3 aws-rule-match children",
        ));
    }
    if matches!(model.kind, ConditionKind::Header) && model.name.is_empty() {
        return Err(LayoutError::new("http-header condition requires name"));
    }
    if !matches!(model.kind, ConditionKind::Header) && !model.name.is_empty() {
        return Err(LayoutError::new(
            "condition name is only supported for http-header",
        ));
    }
    let mut regex_mode = None;
    for i in matches {
        let Some(FeatureKind::Match(m)) = feature(doc, i) else {
            return Err(LayoutError::new("condition accepts only aws-rule-match"));
        };
        if m.value.is_empty() {
            return Err(LayoutError::new("ALB match value cannot be empty"));
        }
        if m.regex
            && !matches!(
                model.kind,
                ConditionKind::Host | ConditionKind::Path | ConditionKind::Header
            )
        {
            return Err(LayoutError::new(
                "regex is supported only for host, path and header conditions",
            ));
        }
        if regex_mode.is_some_and(|v| v != m.regex) {
            return Err(LayoutError::new(
                "one condition cannot mix value and regex matching",
            ));
        }
        regex_mode = Some(m.regex);
        if !m.key.is_empty() && !matches!(model.kind, ConditionKind::Query) {
            return Err(LayoutError::new(
                "match key is only supported for query-string",
            ));
        }
        if matches!(model.kind, ConditionKind::SourceIp) {
            let valid = m.value.split_once('/').is_some_and(|(ip, prefix)| {
                ip.parse::<std::net::IpAddr>()
                    .ok()
                    .zip(prefix.parse::<u8>().ok())
                    .is_some_and(|(ip, prefix)| prefix <= if ip.is_ipv4() { 32 } else { 128 })
            });
            if !valid || m.value == "255.255.255.255/32" {
                return Err(LayoutError::new("source-ip requires an IPv4/IPv6 CIDR"));
            }
        }
        if matches!(model.kind, ConditionKind::Method) && m.value.contains(['*', '?']) {
            return Err(LayoutError::new(
                "HTTP method conditions cannot use wildcards",
            ));
        }
    }
    Ok(())
}
