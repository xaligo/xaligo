use super::{
    action, condition, option, presentation::Presentation, rule, target_group, target_service,
    transform,
};
pub enum FeatureKind {
    Rule(rule::Rule),
    Condition(condition::Condition),
    Match(condition::Match),
    Action(action::Action),
    ForwardTarget(action::ForwardTarget),
    JwtClaim(action::JwtClaim),
    Transform(transform::Transform),
    Rewrite(transform::Rewrite),
    TargetGroup(target_group::TargetGroup),
    TargetService(target_service::TargetService),
    Target(target_group::Target),
    Option(option::OptionSetting),
}
pub struct Feature {
    pub kind: FeatureKind,
    pub presentation: Presentation,
}
