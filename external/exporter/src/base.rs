use crate::ctl::exporter;
use crate::util::error::Error;

pub fn export(input: &str) -> Result<Vec<u8>, Error> {
    exporter::export(input)
}
