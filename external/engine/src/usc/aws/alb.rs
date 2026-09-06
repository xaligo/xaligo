#[rustfmt::skip]
use crate::ent::model::{
    aws::{
        alb::Alb,
        listener::{
            Listener,
            MutualTls,
            Protocol,
        },
    },
    document::DocumentSpec,
};
use crate::util::error::LayoutError;

pub(crate) fn validate(listener: &Listener) -> Result<(), LayoutError> {
    if !matches!(listener.protocol, Protocol::Http | Protocol::Https) {
        return Err(LayoutError::new("ALB listeners support HTTP or HTTPS"));
    }
    if !matches!(listener.mutual_tls, MutualTls::Off)
        && !matches!(listener.protocol, Protocol::Https)
    {
        return Err(LayoutError::new("ALB mTLS requires an HTTPS listener"));
    }
    if !listener.trust_store.is_empty() && !matches!(listener.mutual_tls, MutualTls::Verify) {
        return Err(LayoutError::new("ALB trust-store requires mtls=verify"));
    }
    if matches!(listener.mutual_tls, MutualTls::Verify) && listener.trust_store.is_empty() {
        return Err(LayoutError::new(
            "ALB mtls=verify requires a trust-store reference",
        ));
    }
    if listener.backend_mtls == Some(true) {
        return Err(LayoutError::new(
            "ALB cannot present a client certificate for backend mTLS",
        ));
    }
    super::listener::validate(listener)
}

pub(crate) fn compose(
    document: &mut DocumentSpec,
    index: usize,
    model: &Alb,
) -> Result<(), LayoutError> {
    if super::detail::required(document, index) {
        return super::detail::compose(document, index);
    }
    super::load_balancer::compose(document, index, &model.domain, "ALB", validate)
}
