pub struct TargetGroup {
    pub name: String,
    pub target_type: String,
    pub protocol: String,
    pub port: Option<u16>,
}
pub struct Target {
    pub name: String,
    pub port: Option<u16>,
    pub zone: String,
}
