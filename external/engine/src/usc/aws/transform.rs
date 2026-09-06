use super::detail::{children, feature};
use crate::ent::model::aws::{feature::FeatureKind, transform::*};
use crate::ent::model::document::DocumentSpec;
use crate::util::error::LayoutError;
pub(super) fn label(kind: &TransformKind) -> &'static str {
    match kind {
        TransformKind::Host => "Rewrite host",
        TransformKind::Url => "Rewrite URL",
    }
}
pub(super) fn validate(doc: &DocumentSpec, index: usize) -> Result<(), LayoutError> {
    let entries = children(doc, index);
    if entries.is_empty() {
        return Err(LayoutError::new("transform requires aws-rule-rewrite"));
    }
    for i in entries {
        if !matches!(feature(doc,i),Some(FeatureKind::Rewrite(r)) if !r.regex.is_empty()) {
            return Err(LayoutError::new(
                "transform requires regex/replace rewrite entries",
            ));
        }
    }
    Ok(())
}
