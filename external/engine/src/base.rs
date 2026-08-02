use crate::ent::request::engine::EngineRequest;
use crate::usc::engine::execute;
use crate::util::deserialize::Deserialize;
use crate::util::logger::debug;
#[rustfmt::skip]
use crate::util::mcode::{
    MENGINE_COMPLETE,
    MENGINE_PROCESS,
};
#[rustfmt::skip]
use crate::util::serialize::{
    ErrorResponse,
    Serialize,
};

pub(crate) fn process_request(input: &[u8]) -> Vec<u8> {
    debug(MENGINE_PROCESS, "", None);
    match EngineRequest::deserialize(input).and_then(execute) {
        Ok(response) => {
            debug(MENGINE_COMPLETE, "", None);
            response.serialize()
        }
        Err(message) => ErrorResponse::new(&message).serialize(),
    }
}
