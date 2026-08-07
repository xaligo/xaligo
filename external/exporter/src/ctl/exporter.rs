use crate::ent::request::exporter::ExporterRequest;
use crate::usc::exporter;
use crate::util::error::Error;

pub fn export(input: &str) -> Result<Vec<u8>, Error> {
    if input.trim().is_empty() {
        return Err(Error::invalid("PPTX exporter request JSON is required on stdin"));
    }
    let request: ExporterRequest = serde_json::from_str(input)
        .map_err(|error| Error::invalid(format!("parse PPTX exporter request: {error}")))?;
    exporter::export(&request)
}
