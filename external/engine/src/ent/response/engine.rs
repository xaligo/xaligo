use crate::ent::model::document::ResolvedDocument;
use crate::ent::model::svg::NormalizedSvg;

pub(crate) enum EngineResponse {
    Layout(ResolvedDocument),
    Svg(Vec<u8>),
    NormalizedSvg(NormalizedSvg),
}
