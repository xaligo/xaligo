#[rustfmt::skip]
use super::{
    alb,
    nlb,
};

use crate::cnf::engine::MAX_ELEMENTS;
#[rustfmt::skip]
use crate::ent::model::{
    aws::Component,
    document::DocumentSpec,
};
use crate::util::error::LayoutError;

// Service-specific composition runs once, before the shared layout/router.
// The output contains only ordinary groups/items/text, shared by SVG and PPTX.
pub(crate) fn compose(source: &DocumentSpec) -> Result<DocumentSpec, LayoutError> {
    super::detail::validate(source)?;
    let mut document = source.clone();
    let count = document.elements.len();
    for index in 0..count {
        crate::usc::cancel::check().map_err(LayoutError::new)?;
        match document.elements[index].aws.clone() {
            Some(Component::Alb(model)) => alb::compose(&mut document, index, &model)?,
            Some(Component::Nlb(model)) => nlb::compose(&mut document, index, &model)?,
            Some(Component::Listener(_)) => {
                let parent = source.elements[index]
                    .parent
                    .and_then(|p| source.elements.get(p));
                if !parent
                    .is_some_and(|p| matches!(p.aws, Some(Component::Alb(_) | Component::Nlb(_))))
                {
                    return Err(LayoutError::new(format!(
                        "AWS listener {:?} must belong to an ALB or NLB",
                        source.elements[index].id
                    )));
                }
            }
            None => {}
            Some(Component::Feature(_)) => {}
        }
        if document.elements.len() > MAX_ELEMENTS {
            return Err(LayoutError::new(
                "AWS component expansion exceeds the element limit",
            ));
        }
    }
    for element in &mut document.elements {
        element.aws = None;
    }
    Ok(document)
}
