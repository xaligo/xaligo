#[rustfmt::skip]
use super::{
    alb,
    listener,
    nlb,
};

pub enum Component {
    Alb(alb::Alb),
    Nlb(nlb::Nlb),
    Listener(listener::Listener),
    Feature(super::feature::Feature),
}
