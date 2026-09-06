use super::detail::{children, feature, option_value};
use crate::ent::model::aws::{
    feature::FeatureKind, listener::Protocol, target_group::TargetGroup, Component,
};
use crate::ent::model::document::DocumentSpec;
use crate::util::error::LayoutError;
pub(super) fn validate(
    doc: &DocumentSpec,
    index: usize,
    model: &TargetGroup,
) -> Result<(), LayoutError> {
    if model.name.is_empty() || !matches!(model.target_type.as_str(), "instance" | "ip" | "lambda")
    {
        return Err(LayoutError::new(
            "ALB target group requires name and target-type=instance|ip|lambda",
        ));
    }
    if model.target_type == "lambda" {
        if !model.protocol.is_empty() || model.port.is_some() {
            return Err(LayoutError::new(
                "Lambda target groups do not have protocol/port",
            ));
        }
        if option_value(doc, index, "protocol-version").is_some()
            || option_value(doc, index, "target-control-port").is_some()
            || option_value(doc, index, "target-ip-address-type") == Some("ipv6")
        {
            return Err(LayoutError::new(
                "Lambda targets do not support protocol versions, optimizer ports or IPv6",
            ));
        }
        let targets: Vec<_> = doc
            .elements
            .iter()
            .enumerate()
            .filter_map(|(i, _)| (super::target_service::group(doc, i) == Some(index)).then_some(i))
            .filter_map(|i| match feature(doc, i) {
                Some(FeatureKind::Target(t)) => Some(t),
                _ => None,
            })
            .collect();
        if targets.len() > 1 || targets.iter().any(|t| t.port.is_some()) {
            return Err(LayoutError::new(
                "Lambda target groups accept one function without a port",
            ));
        }
        if children(doc, index)
            .iter()
            .filter(|i| matches!(feature(doc, **i), Some(FeatureKind::TargetService(_))))
            .count()
            > 1
        {
            return Err(LayoutError::new(
                "Lambda target group accepts one logical function service",
            ));
        }
    } else if !matches!(model.protocol.as_str(), "HTTP" | "HTTPS")
        || model.port.is_none_or(|p| p == 0)
    {
        return Err(LayoutError::new(
            "instance/ip target group requires HTTP/HTTPS and port 1..65535",
        ));
    }
    let algorithm = option_value(doc, index, "load_balancing.algorithm.type");
    if algorithm == Some("weighted_random")
        && option_value(doc, index, "stickiness.enabled") == Some("true")
    {
        return Err(LayoutError::new(
            "weighted_random does not support sticky sessions",
        ));
    }
    if algorithm.is_some_and(|a| a != "round_robin")
        && option_value(doc, index, "target-control-port").is_some()
    {
        return Err(LayoutError::new("Target Optimizer requires round_robin"));
    }
    if option_value(doc, index, "load_balancing.algorithm.anomaly_mitigation") == Some("on")
        && option_value(doc, index, "load_balancing.algorithm.type") != Some("weighted_random")
    {
        return Err(LayoutError::new(
            "anomaly mitigation requires weighted_random",
        ));
    }
    if let Some(value) = option_value(doc, index, "slow_start.duration_seconds") {
        if value.parse::<u32>().is_ok_and(|v| v > 0 && v < 30) {
            return Err(LayoutError::new("slow start must be zero or 30..900"));
        }
        if value != "0" && algorithm.is_some_and(|a| a != "round_robin") {
            return Err(LayoutError::new("slow start requires round_robin"));
        }
    }
    Ok(())
}

// Resolve only locally declared groups. External references remain valid for partial diagrams.
pub(super) fn validate_reference(
    doc: &DocumentSpec,
    index: usize,
    reference: &str,
    protocol: Protocol,
    has_auth: bool,
) -> Result<(), LayoutError> {
    if reference.is_empty() {
        return Ok(());
    }
    let mut root = Some(index);
    while let Some(i) = root {
        if matches!(doc.elements[i].aws, Some(Component::Alb(_))) {
            break;
        }
        root = doc.elements[i].parent;
    }
    let Some(root) = root else {
        return Ok(());
    };
    let groups: Vec<_> = children(doc,root).into_iter().filter(|i|matches!(feature(doc,*i),Some(FeatureKind::TargetGroup(m)) if doc.elements[*i].id==reference || m.name==reference)).collect();
    if groups.len() > 1 {
        return Err(LayoutError::new("ambiguous local target-group reference"));
    }
    if groups.first().is_some_and(|i| {
        matches!(
            option_value(doc, *i, "protocol-version"),
            Some("HTTP2" | "GRPC")
        )
    }) && (!matches!(protocol, Protocol::Https) || has_auth)
    {
        return Err(LayoutError::new(
            "HTTP2/GRPC targets require HTTPS and forward-only rules",
        ));
    }
    Ok(())
}
