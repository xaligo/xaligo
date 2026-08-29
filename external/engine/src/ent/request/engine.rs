use crate::ent::model::document::DocumentSpec;

pub(crate) enum EngineRequest {
    Document {
        operation: u8,
        document: DocumentSpec,
    },
    NormalizeSvg(Vec<u8>),
}
