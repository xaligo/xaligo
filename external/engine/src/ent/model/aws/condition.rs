pub enum ConditionKind {
    Host,
    Path,
    Header,
    Method,
    Query,
    SourceIp,
}
pub struct Condition {
    pub kind: ConditionKind,
    pub name: String,
}
pub struct Match {
    pub value: String,
    pub key: String,
    pub regex: bool,
}
