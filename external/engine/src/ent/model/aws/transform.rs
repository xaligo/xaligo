pub enum TransformKind {
    Host,
    Url,
}
pub struct Transform {
    pub kind: TransformKind,
}
pub struct Rewrite {
    pub regex: String,
    pub replacement: String,
}
