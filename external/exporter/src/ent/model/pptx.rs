use serde::Deserialize;

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct ExporterRequest {
    pub plan: DocumentPlan,
    #[serde(default)]
    pub options: ExportOptions,
}

#[derive(Default, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct ExportOptions {
    pub title: Option<String>,
    pub author: Option<String>,
    pub company: Option<String>,
    pub subject: Option<String>,
    pub compression: Option<bool>,
}

#[derive(Deserialize)]
#[serde(untagged)]
pub enum DocumentPlan {
    Current(CurrentPlan),
    Legacy(LegacyPlan),
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct CurrentPlan {
    pub schema_version: u32,
    pub pages: Vec<Page>,
    #[serde(default)]
    pub legend: Vec<LegendEntry>,
    #[serde(default)]
    pub connector_legend: Vec<ConnectorLegendEntry>,
}

#[derive(Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct LegacyPlan {
    pub slide: Slide,
    pub ops: Vec<Op>,
    #[serde(default)]
    pub legend: Vec<LegendEntry>,
    #[serde(default)]
    pub connector_legend: Vec<ConnectorLegendEntry>,
}

#[derive(Deserialize)]
pub struct Page {
    pub id: String,
    pub slide: Slide,
    pub ops: Vec<Op>,
}

#[derive(Clone, Deserialize)]
pub struct Slide {
    pub w: f64,
    pub h: f64,
    pub background: String,
}

#[derive(Clone, Default, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Line {
    pub color: String,
    pub width: f64,
    pub dash: String,
    pub transparency: f64,
    pub begin_arrow_type: Option<String>,
    pub end_arrow_type: Option<String>,
    pub begin_arrow_extend_in: Option<f64>,
    pub end_arrow_extend_in: Option<f64>,
}

#[derive(Clone, Deserialize)]
pub struct Fill {
    pub color: String,
    pub transparency: f64,
}

#[derive(Clone, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Point {
    pub x: f64,
    pub y: f64,
    #[serde(default)]
    pub move_to: bool,
}

#[derive(Clone, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct TextLayout {
    #[serde(default)]
    pub wrap: bool,
    pub fit: Option<String>,
    pub overflow: Option<String>,
    #[serde(default)]
    pub clip: bool,
    pub line_height: Option<f64>,
    #[serde(default)]
    pub padding: Padding,
}

#[derive(Clone, Default, Deserialize)]
pub struct Padding {
    #[serde(default)]
    pub top: f64,
    #[serde(default)]
    pub right: f64,
    #[serde(default)]
    pub bottom: f64,
    #[serde(default)]
    pub left: f64,
}

#[derive(Clone, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Op {
    pub id: Option<String>,
    pub group_id: Option<String>,
    #[serde(default)]
    pub front_layer: bool,
    pub kind: String,
    pub x: f64,
    pub y: f64,
    pub w: f64,
    pub h: f64,
    pub rotate: Option<f64>,
    pub line: Option<Line>,
    pub fill: Option<Fill>,
    pub text: Option<String>,
    pub color: Option<String>,
    pub font_face: Option<String>,
    pub font_size: Option<f64>,
    pub bold: Option<bool>,
    pub align: Option<String>,
    pub valign: Option<String>,
    pub text_layout: Option<TextLayout>,
    pub data: Option<String>,
    pub transparency: Option<f64>,
    pub points: Option<Vec<Point>>,
    #[serde(default)]
    pub flip_h: bool,
    #[serde(default)]
    pub flip_v: bool,
}

#[derive(Clone, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct LegendEntry {
    pub catalog_id: u32,
    pub abbreviation: String,
    pub official_name: String,
    pub data: Option<String>,
}

#[derive(Clone, Deserialize)]
pub struct ConnectorLegendEntry {
    pub id: String,
    pub kind: String,
    pub label: String,
    pub description: String,
    pub line: Line,
}

impl DocumentPlan {
    pub fn parts(&self) -> (&[Page], Option<(&Slide, &[Op])>, &[LegendEntry], &[ConnectorLegendEntry]) {
        match self {
            Self::Current(plan) => (&plan.pages, None, &plan.legend, &plan.connector_legend),
            Self::Legacy(plan) => (&[], Some((&plan.slide, &plan.ops)), &plan.legend, &plan.connector_legend),
        }
    }
}
