#[rustfmt::skip]
use crate::cnf::engine::{
    OPERATION_LAYOUT,
    OPERATION_SVG,
};
use crate::ent::request::engine::EngineRequest;
use crate::ent::response::engine::EngineResponse;
#[rustfmt::skip]
use crate::usc::{
    layout,
    svg,
};

pub(crate) fn execute(request: EngineRequest) -> Result<EngineResponse, String> {
    crate::usc::cancel::check()?;
    match request {
        EngineRequest::Document {
            operation,
            document,
        } => {
            let resolved = layout::resolve(&document).map_err(|error| error.to_string())?;
            match operation {
                OPERATION_LAYOUT => Ok(EngineResponse::Layout(resolved)),
                OPERATION_SVG => Ok(EngineResponse::Svg(svg::render(&resolved))),
                _ => Err(format!("unsupported engine operation {operation}")),
            }
        }
        EngineRequest::NormalizeSvg(input) => svg::normalize(&input)
            .map(EngineResponse::NormalizedSvg)
            .map_err(|error| error.to_string()),
    }
}
