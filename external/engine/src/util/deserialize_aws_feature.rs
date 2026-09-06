fn decode_aws_feature(
    s: &[String; STRING_FIELD_COUNT],
    presentation: crate::ent::model::aws::presentation::Presentation,
) -> Result<crate::ent::model::aws::feature::Feature, String> {
    use crate::cnf::engine_abi::*;
    use crate::ent::model::aws::{
        action::*, condition::*, feature::*, option::OptionSetting, rule::Rule, target_group::*,
        target_service::*, transform::*,
    };
    let value = &s[STRING_AWS_VALUE];
    let name = &s[STRING_AWS_NAME];
    let subtype = s[STRING_AWS_TYPE].as_str();
    let order = &s[STRING_AWS_ORDER];
    let allowed: &[usize] = match s[STRING_AWS_KIND].as_str() {
        "rule" => &[STRING_AWS_VALUE],
        "condition" | "jwt-claim" => &[STRING_AWS_TYPE, STRING_AWS_NAME],
        "target-service" => &[STRING_AWS_TYPE, STRING_AWS_NAME, STRING_AWS_VALUE],
        "match" | "option" => &[STRING_AWS_TYPE, STRING_AWS_NAME, STRING_AWS_VALUE],
        "action" => &[STRING_AWS_TYPE, STRING_AWS_ORDER, STRING_AWS_VALUE],
        "forward-target" => &[STRING_AWS_ORDER, STRING_AWS_VALUE],
        "transform" => &[STRING_AWS_TYPE],
        "rewrite" => &[STRING_AWS_VALUE, STRING_AWS_AUX],
        "target-group" => &[
            STRING_AWS_NAME,
            STRING_AWS_TYPE,
            STRING_AWS_VALUE,
            STRING_AWS_ORDER,
        ],
        "target" => &[STRING_AWS_NAME, STRING_AWS_VALUE, STRING_AWS_ORDER],
        _ => return Err("unsupported AWS component kind".into()),
    };
    for slot in [
        STRING_AWS_TYPE,
        STRING_AWS_NAME,
        STRING_AWS_VALUE,
        STRING_AWS_AUX,
        STRING_AWS_ORDER,
    ] {
        if !allowed.contains(&slot) && !s[slot].is_empty() {
            return Err("AWS feature contains inapplicable typed fields".into());
        }
    }
    fn number(value: &str, default: Option<u16>, max: u16, zero: bool) -> Result<u16, String> {
        if value.is_empty() {
            if let Some(n) = default {
                return Ok(n);
            }
        }
        value
            .parse::<u16>()
            .ok()
            .filter(|n| *n <= max && (zero || *n > 0))
            .ok_or_else(|| "invalid AWS integer parameter".into())
    }
    let kind = match s[STRING_AWS_KIND].as_str() {
        "rule" => FeatureKind::Rule(Rule {
            priority: if value == "default" {
                None
            } else {
                Some(number(value, None, 50000, false)?)
            },
        }),
        "condition" => FeatureKind::Condition(Condition {
            kind: match subtype {
                "host-header" => ConditionKind::Host,
                "path-pattern" => ConditionKind::Path,
                "http-header" => ConditionKind::Header,
                "http-request-method" => ConditionKind::Method,
                "query-string" => ConditionKind::Query,
                "source-ip" => ConditionKind::SourceIp,
                _ => return Err("invalid ALB condition field".into()),
            },
            name: name.clone(),
        }),
        "match" => FeatureKind::Match(Match {
            value: value.clone(),
            key: name.clone(),
            regex: match subtype {
                "value" => false,
                "regex" => true,
                _ => return Err("invalid ALB match mode".into()),
            },
        }),
        "action" => FeatureKind::Action(Action {
            kind: match subtype {
                "forward" => ActionKind::Forward,
                "redirect" => ActionKind::Redirect,
                "fixed-response" => ActionKind::FixedResponse,
                "authenticate-oidc" => ActionKind::Oidc,
                "authenticate-cognito" => ActionKind::Cognito,
                "jwt-validation" => ActionKind::Jwt,
                _ => return Err("invalid ALB action type".into()),
            },
            order: number(order, Some(1), 50000, false)?,
            target_group: value.clone(),
        }),
        "forward-target" => FeatureKind::ForwardTarget(ForwardTarget {
            target_group: value.clone(),
            weight: number(order, Some(1), 999, true)?,
        }),
        "jwt-claim" => FeatureKind::JwtClaim(JwtClaim {
            name: name.clone(),
            format: subtype.into(),
        }),
        "transform" => FeatureKind::Transform(Transform {
            kind: match subtype {
                "host-header-rewrite" => TransformKind::Host,
                "url-rewrite" => TransformKind::Url,
                _ => return Err("invalid ALB transform type".into()),
            },
        }),
        "rewrite" => FeatureKind::Rewrite(Rewrite {
            regex: value.clone(),
            replacement: s[STRING_AWS_AUX].clone(),
        }),
        "target-group" => FeatureKind::TargetGroup(TargetGroup {
            name: name.clone(),
            target_type: subtype.into(),
            protocol: value.clone(),
            port: if order.is_empty() {
                None
            } else {
                Some(number(order, None, 65535, false)?)
            },
        }),
        "target-service" => FeatureKind::TargetService(TargetService {
            kind: match subtype {
                "ecs" => ServiceKind::Ecs,
                "eks" => ServiceKind::Eks,
                "ec2" => ServiceKind::Ec2,
                "lambda" => ServiceKind::Lambda,
                "ip" => ServiceKind::Ip,
                _ => return Err("target service must be ecs, eks, ec2, lambda or ip".into()),
            },
            name: name.clone(),
            reference: value.clone(),
        }),
        "target" => FeatureKind::Target(Target {
            name: name.clone(),
            zone: value.clone(),
            port: if order.is_empty() {
                None
            } else {
                Some(number(order, None, 65535, false)?)
            },
        }),
        "option" => FeatureKind::Option(OptionSetting {
            key: crate::usc::aws::option::key(subtype)?,
            value: value.clone(),
            name: name.clone(),
        }),
        _ => return Err("unsupported AWS component kind".into()),
    };
    for v in [
        &s[STRING_AWS_VALUE],
        &s[STRING_AWS_NAME],
        &s[STRING_AWS_AUX],
    ] {
        if v.len() > 2048 || v.chars().any(char::is_control) {
            return Err("AWS feature values must be single-line and at most 2048 bytes".into());
        }
    }
    Ok(Feature { kind, presentation })
}
