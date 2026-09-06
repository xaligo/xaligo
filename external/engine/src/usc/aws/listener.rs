#[rustfmt::skip]
use super::drawing::{
    badge,
    part,
};
#[rustfmt::skip]
use crate::ent::model::{
    aws::listener::{
        Listener,
        MutualTls,
        Protocol,
    },
    document::*,
};
use crate::util::error::LayoutError;

pub(crate) fn validate(model: &Listener) -> Result<(), LayoutError> {
    if model.port == 0 {
        return Err(LayoutError::new("AWS listener port must be 1..65535"));
    }
    if !model.certificate.is_empty() && !tls(model.protocol) {
        return Err(LayoutError::new(
            "AWS listener certificate requires HTTPS or TLS",
        ));
    }
    if model.backend_mtls == Some(true) && model.backend_tls != Some(true) {
        return Err(LayoutError::new(
            "backend-mtls=true requires backend-tls=true",
        ));
    }
    for value in [&model.certificate, &model.trust_store, &model.target_group] {
        if value.chars().any(char::is_control) || value.len() > 2048 {
            return Err(LayoutError::new(
                "AWS listener references must be single-line and at most 2048 bytes",
            ));
        }
    }
    Ok(())
}

pub(super) fn tls(protocol: Protocol) -> bool {
    matches!(protocol, Protocol::Https | Protocol::Tls)
}

pub(super) fn protocol_name(protocol: Protocol) -> &'static str {
    match protocol {
        Protocol::Http => "HTTP",
        Protocol::Https => "HTTPS",
        Protocol::Tcp => "TCP",
        Protocol::Tls => "TLS",
        Protocol::Udp => "UDP",
        Protocol::TcpUdp => "TCP_UDP",
        Protocol::Quic => "QUIC",
        Protocol::TcpQuic => "TCP_QUIC",
    }
}

fn port_label(model: &Listener) -> String {
    format!("{} :{}", protocol_name(model.protocol), model.port)
}

fn header_label(model: &Listener) -> String {
    if model.show_title == Some(false) {
        port_label(model)
    } else {
        format!(
            "Listener · {}:{}",
            protocol_name(model.protocol),
            model.port
        )
    }
}

fn references(model: &Listener) -> &'static str {
    match (model.certificate.is_empty(), model.trust_store.is_empty()) {
        (false, false) => "Cert · CA trust store",
        (false, true) => "Cert",
        (true, false) => "CA trust store",
        _ => "",
    }
}

fn target_label(model: &Listener) -> String {
    if model.target_group.is_empty() {
        return String::new();
    }
    let name = model
        .target_group
        .split("targetgroup/")
        .last()
        .unwrap_or(&model.target_group)
        .split('/')
        .next()
        .unwrap_or("");
    let mut label = format!("→ {name}");
    let mut shortened = false;
    while crate::usc::layout::presentation_text_width(&label, 12.0) > 172.0 {
        label.pop();
        shortened = true;
    }
    if shortened {
        label.push('…');
    }
    label
}

// Measure visible content only; target TLS/mTLS remains metadata, not extra badges.
pub(super) fn measure(model: &Listener) -> (f64, f64) {
    let width = [
        crate::usc::layout::presentation_text_width(&header_label(model), 13.0) + 20.0,
        crate::usc::layout::presentation_text_width(references(model), 12.0) + 20.0,
        crate::usc::layout::presentation_text_width(&target_label(model), 12.0) + 20.0,
        124.0,
    ]
    .into_iter()
    .fold(0.0_f64, f64::max)
    .ceil();
    let height = 104.0
        + if references(model).is_empty() {
            0.0
        } else {
            20.0
        }
        + if model.target_group.is_empty() {
            0.0
        } else {
            24.0
        };
    (width, height)
}

pub(super) fn compose(document: &mut DocumentSpec, index: usize, model: &Listener, x: f64, y: f64) {
    let (width, height) = measure(model);
    let owner = &mut document.elements[index];
    owner.concept = Concept::Group;
    owner.layout = LayoutPolicy::Absolute;
    owner.padding = Insets::default();
    owner.margin = Insets::default();
    owner.x = Some(x);
    owner.y = Some(y);
    owner.width = Some(width);
    owner.height = Some(height);
    owner.text.value.clear();
    owner.icon.reference.clear();
    owner.visual.shape = Shape::Rectangle;
    owner.visual.fill = "#ffffff".into();
    owner.visual.stroke = "#c7d8f4".into();
    owner.visual.stroke_width = Some(1.0);
    owner.visual.corner_radius = Some(8.0);
    let owner = owner.clone();
    let mut top = 10.0;
    let mut heading = part(&owner, index, "header", [10.0, top, width - 20.0, 28.0]);
    heading.text.value = header_label(model);
    heading.text.font_size = Some(13.0);
    let mut accent = part(&owner, index, "accent", [8.0, 0.0, width - 16.0, 3.0]);
    accent.visual.shape = Shape::Rectangle;
    accent.visual.fill = "#8b5cf6".into();
    top += 28.0;
    document.elements.extend([heading, accent]);
    for (name, label, active) in [
        (
            "tls",
            if tls(model.protocol) {
                "TLS ON"
            } else {
                "TLS OFF"
            },
            tls(model.protocol),
        ),
        (
            "mtls",
            match model.mutual_tls {
                MutualTls::Off => "mTLS OFF",
                MutualTls::Verify => "mTLS ON",
                MutualTls::Passthrough => "mTLS PASS",
            },
            !matches!(model.mutual_tls, MutualTls::Off),
        ),
    ] {
        let mut item = badge(
            &owner,
            index,
            name,
            [10.0, top, width - 20.0, 24.0],
            label,
            active,
        );
        item.text.font_size = Some(12.0);
        document.elements.push(item);
        top += 28.0;
    }
    if !references(model).is_empty() {
        let mut item = part(&owner, index, "references", [10.0, top, width - 20.0, 20.0]);
        item.text.value = references(model).into();
        item.text.font_size = Some(12.0);
        item.text.color = "#6d28d9".into();
        document.elements.push(item);
        top += 20.0;
    }
    if !model.target_group.is_empty() {
        let mut target = part(
            &owner,
            index,
            "target-group",
            [10.0, top, width - 20.0, 24.0],
        );
        target.text.value = target_label(model);
        target.text.font_size = Some(12.0);
        document.elements.push(target);
    }
}
