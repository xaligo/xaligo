use super::detail::{children, feature, option_value};
use crate::ent::model::{
    aws::{
        feature::FeatureKind,
        target_service::{ServiceKind, TargetService},
    },
    document::DocumentSpec,
};
use crate::util::error::LayoutError;

pub(super) fn owner(kind: &ServiceKind) -> &'static str {
    match kind {
        ServiceKind::Ecs => "ecs",
        ServiceKind::Eks => "eks",
        ServiceKind::Ec2 => "ec2",
        ServiceKind::Lambda => "lambda",
        ServiceKind::Ip => "ip-service",
    }
}
pub(super) fn label(kind: &ServiceKind) -> &'static str {
    match kind {
        ServiceKind::Ecs => "ECS",
        ServiceKind::Eks => "EKS",
        ServiceKind::Ec2 => "EC2",
        ServiceKind::Lambda => "Lambda",
        ServiceKind::Ip => "IP service",
    }
}
pub(super) fn group(doc: &DocumentSpec, index: usize) -> Option<usize> {
    let mut i = doc.elements[index].parent;
    while let Some(n) = i {
        if matches!(feature(doc, n), Some(FeatureKind::TargetGroup(_))) {
            return Some(n);
        }
        i = doc.elements[n].parent;
    }
    None
}
pub(super) fn validate(
    doc: &DocumentSpec,
    index: usize,
    model: &TargetService,
) -> Result<(), LayoutError> {
    if model.name.is_empty() {
        return Err(LayoutError::new("target service requires name"));
    }
    let Some(parent) = group(doc, index) else {
        return Err(LayoutError::new(
            "target service requires target-group parent",
        ));
    };
    let Some(FeatureKind::TargetGroup(tg)) = feature(doc, parent) else {
        return Err(LayoutError::new("missing target group"));
    };
    let valid = match model.kind {
        ServiceKind::Lambda => tg.target_type == "lambda",
        ServiceKind::Ip => tg.target_type == "ip",
        ServiceKind::Ec2 | ServiceKind::Eks => matches!(tg.target_type.as_str(), "ip" | "instance"),
        ServiceKind::Ecs => {
            let launch = option_value(doc, index, "launch-type");
            let network = option_value(doc, index, "network-mode");
            if launch == Some("fargate") && network.is_some_and(|n| n != "awsvpc") {
                return Err(LayoutError::new("Fargate requires awsvpc networking"));
            }
            if launch == Some("fargate") || network == Some("awsvpc") {
                tg.target_type == "ip"
            } else if matches!(network, Some("bridge" | "host")) {
                tg.target_type == "instance"
            } else {
                matches!(tg.target_type.as_str(), "ip" | "instance")
            }
        }
    };
    if !valid {
        return Err(LayoutError::new(
            "target service and target-type are incompatible",
        ));
    }
    if matches!(model.kind, ServiceKind::Eks)
        && tg.target_type != "instance"
        && option_value(doc, index, "node-port").is_some()
    {
        return Err(LayoutError::new("EKS node-port requires instance targets"));
    }
    if children(doc, index).iter().any(|i| {
        !matches!(
            feature(doc, *i),
            Some(FeatureKind::Target(_) | FeatureKind::Option(_))
        )
    }) {
        return Err(LayoutError::new(
            "target service accepts registered targets and options",
        ));
    }
    Ok(())
}
