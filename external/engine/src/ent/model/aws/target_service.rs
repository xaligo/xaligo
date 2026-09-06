// Logical workload associated with a target group, not an AWS target type.
pub enum ServiceKind {
    Ecs,
    Eks,
    Ec2,
    Lambda,
    Ip,
}
pub struct TargetService {
    pub kind: ServiceKind,
    pub name: String,
    pub reference: String,
}
