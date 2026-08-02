#[derive(Clone, Debug, PartialEq)]
pub struct NormalizedSvg {
    pub data: Vec<u8>,
    pub view_box: String,
    pub width: f64,
    pub height: f64,
}
