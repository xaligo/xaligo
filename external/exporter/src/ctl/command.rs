use std::io::{Read, Write};

use base64::Engine;

fn main() {
    if let Err(error) = run() {
        eprintln!("{error}");
        std::process::exit(1);
    }
}

fn run() -> Result<(), Box<dyn std::error::Error>> {
    let mut input = String::new();
    std::io::stdin().read_to_string(&mut input)?;
    let bytes = xaligo_pptx_exporter::export(&input)?;
    let encoded = base64::engine::general_purpose::STANDARD.encode(bytes);
    std::io::stdout().write_all(encoded.as_bytes())?;
    Ok(())
}
