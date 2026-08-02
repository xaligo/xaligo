use crate::ent::request::engine::decode_request;
use crate::ent::response::engine::{encode_error, encode_success};
use crate::usc::engine::execute;

pub(crate) fn process_request(input: &[u8]) -> Vec<u8> {
    match decode_request(input).and_then(execute) {
        Ok(response) => encode_success(response),
        Err(message) => encode_error(&message),
    }
}
