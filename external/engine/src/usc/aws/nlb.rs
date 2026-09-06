#[rustfmt::skip]
use crate::ent::model::{
    aws::{
        listener::{
            Listener,
            MutualTls,
            Protocol,
        },
        nlb::Nlb,
    },
    document::DocumentSpec,
};
use crate::util::error::LayoutError;

pub(crate) fn validate(listener: &Listener) -> Result<(), LayoutError> {
    if matches!(listener.protocol, Protocol::Http | Protocol::Https) {
        return Err(LayoutError::new(
            "NLB listeners do not use HTTP or HTTPS protocols",
        ));
    }
    if !matches!(listener.mutual_tls, MutualTls::Off) || !listener.trust_store.is_empty() {
        return Err(LayoutError::new(
            "NLB does not terminate mTLS or attach an ALB trust store; use TCP passthrough and backend-mtls",
        ));
    }
    if listener.backend_mtls == Some(true)
        && (!matches!(listener.protocol, Protocol::Tcp) || listener.backend_tls != Some(true))
    {
        return Err(LayoutError::new(
            "backend-mtls requires NLB TCP passthrough with backend-tls=true",
        ));
    }
    super::listener::validate(listener)
}

pub(crate) fn compose(
    document: &mut DocumentSpec,
    index: usize,
    model: &Nlb,
) -> Result<(), LayoutError> {
    if super::detail::required(document, index) {
        return super::detail::compose(document, index);
    }
    super::load_balancer::compose(document, index, &model.domain, "NLB", validate)
}
