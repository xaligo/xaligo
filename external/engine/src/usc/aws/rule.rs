use super::detail::{children, feature};
use crate::ent::model::aws::{
    condition::ConditionKind, feature::FeatureKind, rule::Rule, transform::TransformKind,
};
use crate::ent::model::document::DocumentSpec;
use crate::util::error::LayoutError;

pub(super) fn validate(doc: &DocumentSpec, index: usize, rule: &Rule) -> Result<(), LayoutError> {
    let mut evaluations = 0;
    let mut wildcards = 0;
    let mut single = std::collections::HashSet::new();
    let mut transforms = std::collections::HashSet::new();
    let mut actions = Vec::new();
    for i in children(doc, index) {
        match feature(doc, i) {
            Some(FeatureKind::Condition(c)) => {
                if rule.priority.is_none() {
                    return Err(LayoutError::new("default rule cannot have conditions"));
                }
                if !matches!(c.kind, ConditionKind::Header | ConditionKind::Query)
                    && !single.insert(super::condition::label(&c.kind))
                {
                    return Err(LayoutError::new(
                        "ALB rule duplicates a singleton condition",
                    ));
                }
                for m in children(doc, i) {
                    if let Some(FeatureKind::Match(m)) = feature(doc, m) {
                        evaluations += 1;
                        if !m.regex {
                            wildcards += m
                                .value
                                .chars()
                                .chain(m.key.chars())
                                .filter(|c| matches!(c, '*' | '?'))
                                .count();
                        }
                    }
                }
            }
            Some(FeatureKind::Transform(t)) => {
                if rule.priority.is_none() {
                    return Err(LayoutError::new("default rule cannot have transforms"));
                }
                if !transforms.insert(matches!(t.kind, TransformKind::Host)) {
                    return Err(LayoutError::new("ALB rule allows one transform per type"));
                }
            }
            Some(FeatureKind::Action(a)) => actions.push(a),
            _ => {
                return Err(LayoutError::new(
                    "ALB rule accepts conditions, transforms and actions",
                ));
            }
        }
    }
    if evaluations > 5 || wildcards > 5 || (rule.priority.is_some() && evaluations == 0) {
        return Err(LayoutError::new(
            "non-default ALB rule requires 1..5 matches and at most five wildcards",
        ));
    }
    if rule.priority == Some(0) || rule.priority.is_some_and(|p| p > 50000) {
        return Err(LayoutError::new(
            "rule priority must be 1..50000 or default",
        ));
    }
    actions.sort_by_key(|a| a.order);
    if actions.is_empty()
        || actions.len() > 2
        || actions.windows(2).any(|a| a[0].order == a[1].order)
    {
        return Err(LayoutError::new(
            "ALB rule requires a routing action and at most one authentication action, with unique order",
        ));
    }
    if !super::action::routing(&actions.last().unwrap().kind)
        || actions[..actions.len() - 1]
            .iter()
            .any(|a| super::action::routing(&a.kind))
    {
        return Err(LayoutError::new("exactly one routing action must be last"));
    }
    Ok(())
}
