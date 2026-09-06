pub enum ActionKind {
    Forward,
    Redirect,
    FixedResponse,
    Oidc,
    Cognito,
    Jwt,
}
pub struct Action {
    pub kind: ActionKind,
    pub order: u16,
    pub target_group: String,
}
pub struct ForwardTarget {
    pub target_group: String,
    pub weight: u16,
}
pub struct JwtClaim {
    pub name: String,
    pub format: String,
}
