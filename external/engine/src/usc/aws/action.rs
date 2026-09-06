use super::detail::{children, feature, option_value};
use crate::ent::model::aws::{action::*, feature::FeatureKind, listener::Protocol};
use crate::ent::model::document::DocumentSpec;
use crate::util::error::LayoutError;

pub(super) fn routing(kind: &ActionKind) -> bool {
    matches!(
        kind,
        ActionKind::Forward | ActionKind::Redirect | ActionKind::FixedResponse
    )
}
pub(super) fn owner(kind: &ActionKind) -> &'static str {
    match kind {
        ActionKind::Forward => "forward",
        ActionKind::Redirect => "redirect",
        ActionKind::FixedResponse => "fixed",
        ActionKind::Oidc => "oidc",
        ActionKind::Cognito => "cognito",
        ActionKind::Jwt => "jwt",
    }
}
pub(super) fn label(kind: &ActionKind) -> &'static str {
    match kind {
        ActionKind::Forward => "Forward",
        ActionKind::Redirect => "Redirect",
        ActionKind::FixedResponse => "Response",
        ActionKind::Oidc => "OIDC",
        ActionKind::Cognito => "Cognito",
        ActionKind::Jwt => "JWT",
    }
}
pub(super) fn validate(
    doc: &DocumentSpec,
    index: usize,
    action: &Action,
    protocol: Protocol,
) -> Result<(), LayoutError> {
    if action.order == 0 || action.order > 50000 {
        return Err(LayoutError::new("action order must be 1..50000"));
    }
    if !routing(&action.kind) && !matches!(protocol, Protocol::Https) {
        return Err(LayoutError::new("ALB authentication requires HTTPS"));
    }
    if !matches!(action.kind, ActionKind::Forward) && !action.target_group.is_empty() {
        return Err(LayoutError::new("only forward actions accept target-group"));
    }
    let targets: Vec<_> = children(doc, index)
        .into_iter()
        .filter(|i| matches!(feature(doc, *i), Some(FeatureKind::ForwardTarget(_))))
        .collect();
    if matches!(action.kind, ActionKind::Forward) {
        let has_auth = doc.elements[index].parent.is_some_and(|parent| {
            children(doc, parent).iter().any(
                |i| matches!(feature(doc,*i),Some(FeatureKind::Action(a)) if !routing(&a.kind)),
            )
        });
        super::target_group::validate_reference(
            doc,
            index,
            &action.target_group,
            protocol,
            has_auth,
        )?;
        if targets.len() > 5
            || targets.is_empty() && action.target_group.is_empty()
            || !targets.is_empty() && !action.target_group.is_empty()
        {
            return Err(LayoutError::new(
                "forward requires a target-group or 1..5 weighted targets, not both",
            ));
        }
        let mut identities = std::collections::HashSet::new();
        for i in targets {
            if let Some(FeatureKind::ForwardTarget(t)) = feature(doc, i) {
                if t.target_group.is_empty()
                    || t.weight > 999
                    || !identities.insert(&t.target_group)
                {
                    return Err(LayoutError::new(
                        "forward targets require unique references and weights 0..999",
                    ));
                }
                super::target_group::validate_reference(
                    doc,
                    index,
                    &t.target_group,
                    protocol,
                    has_auth,
                )?;
            }
        }
    } else if !targets.is_empty() {
        return Err(LayoutError::new(
            "weighted targets require a forward action",
        ));
    }
    let required: &[&str] = match action.kind {
        ActionKind::Redirect => &["status-code"],
        ActionKind::FixedResponse => &["status-code"],
        ActionKind::Oidc => &[
            "issuer",
            "authorization-endpoint",
            "token-endpoint",
            "user-info-endpoint",
            "client-id",
        ],
        ActionKind::Cognito => &["user-pool", "user-pool-domain", "client-id"],
        ActionKind::Jwt => &["jwt-issuer", "jwks-endpoint"],
        _ => &[],
    };
    for key in required {
        if option_value(doc, index, key).is_none() {
            return Err(LayoutError::new(format!(
                "ALB {} requires option {key}",
                label(&action.kind)
            )));
        }
    }
    if matches!(action.kind, ActionKind::Redirect) {
        if matches!(protocol, Protocol::Https)
            && option_value(doc, index, "protocol") == Some("HTTP")
        {
            return Err(LayoutError::new("HTTPS cannot redirect to HTTP"));
        }
        if ["protocol", "host", "port", "path"]
            .iter()
            .all(|key| option_value(doc, index, key).is_none())
        {
            return Err(LayoutError::new(
                "redirect must modify protocol, host, port or path",
            ));
        }
        if option_value(doc, index, "port")
            .is_some_and(|p| p != "#{port}" && p.parse::<u16>().map_or(true, |n| n == 0))
        {
            return Err(LayoutError::new(
                "redirect port must be 1..65535 or #{port}",
            ));
        }
        if option_value(doc, index, "path").is_some_and(|p| !p.starts_with('/')) {
            return Err(LayoutError::new("redirect path must start with /"));
        }
    }
    if matches!(action.kind, ActionKind::FixedResponse)
        && option_value(doc, index, "message-body").is_some_and(|v| v.len() > 1024)
    {
        return Err(LayoutError::new("fixed response body exceeds 1024 bytes"));
    }
    if matches!(action.kind, ActionKind::Jwt)
        && !option_value(doc, index, "jwks-endpoint")
            .unwrap_or("")
            .starts_with("https://")
    {
        return Err(LayoutError::new("JWKS endpoint must use HTTPS"));
    }
    let mut claims = std::collections::HashSet::new();
    let mut extras = 0;
    for i in children(doc, index) {
        if let Some(FeatureKind::JwtClaim(c)) = feature(doc, i) {
            if !claims.insert(&c.name) || claims.len() > 10 {
                return Err(LayoutError::new(
                    "JWT requires unique additional claims (up to 10)",
                ));
            }
        }
        if let Some(FeatureKind::Option(o)) = feature(doc, i) {
            if super::option::DEFINITIONS[o.key].name == "auth-parameter" {
                extras += 1;
            }
        }
    }
    if extras > 10 {
        return Err(LayoutError::new(
            "authentication accepts up to 10 extra parameters",
        ));
    }
    Ok(())
}
