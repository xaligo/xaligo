fn decode_aws_component(
    strings: &[String; STRING_FIELD_COUNT],
    port: Option<f64>,
    backend_tls: Option<bool>,
    backend_mtls: Option<bool>,
    show_title: Option<bool>,
) -> Result<Option<crate::ent::model::aws::Component>, String> {
    use crate::cnf::engine_abi::*;
    #[rustfmt::skip]
    use crate::ent::model::aws::{
        Component,
        alb::Alb,
        listener::{
            Listener,
            MutualTls,
            Protocol,
        },
        nlb::Nlb,
    };
    let kind = strings[STRING_AWS_KIND].as_str();
    let presentation = crate::usc::aws::presentation::parse(
        &strings[STRING_AWS_DETAIL_LEVEL],
        &strings[STRING_AWS_SHOW],
        &strings[STRING_AWS_HIDE],
    )?;
    if kind.is_empty()
        && (strings[STRING_AWS_DETAIL_LEVEL..]
            .iter()
            .any(|s| !s.is_empty()))
    {
        return Err("AWS feature fields require an AWS component kind".into());
    }
    if matches!(kind, "alb" | "nlb" | "listener")
        && strings[STRING_AWS_TYPE..].iter().any(|s| !s.is_empty())
    {
        return Err("AWS feature fields require a feature component kind".into());
    }
    if kind != "listener"
        && (strings[STRING_AWS_PROTOCOL..=STRING_AWS_TARGET_GROUP]
            .iter()
            .any(|s| !s.is_empty())
            || port.is_some()
            || backend_tls.is_some()
            || backend_mtls.is_some()
            || show_title.is_some())
    {
        return Err("AWS listener fields require the listener component kind".to_owned());
    }
    if kind != "alb" && kind != "nlb" && !strings[STRING_AWS_DOMAIN].is_empty() {
        return Err("AWS domain requires an ALB or NLB component".to_owned());
    }
    match kind {
        "" => Ok(None),
        "alb" => Ok(Some(Component::Alb(Alb {
            domain: strings[STRING_AWS_DOMAIN].clone(),
            presentation,
        }))),
        "nlb" => Ok(Some(Component::Nlb(Nlb {
            domain: strings[STRING_AWS_DOMAIN].clone(),
            presentation,
        }))),
        "listener" => {
            let number = port.ok_or_else(|| "AWS listener port is required".to_owned())?;
            if !number.is_finite() || number.fract() != 0.0 || !(1.0..=65535.0).contains(&number) {
                return Err("AWS listener port must be an integer in 1..65535".to_owned());
            }
            let protocol = match strings[STRING_AWS_PROTOCOL].as_str() {
                "HTTP" => Protocol::Http,
                "HTTPS" => Protocol::Https,
                "TCP" => Protocol::Tcp,
                "TLS" => Protocol::Tls,
                "UDP" => Protocol::Udp,
                "TCP_UDP" => Protocol::TcpUdp,
                "QUIC" => Protocol::Quic,
                "TCP_QUIC" => Protocol::TcpQuic,
                value => return Err(format!("unsupported AWS listener protocol {value:?}")),
            };
            let mutual_tls = match strings[STRING_AWS_MUTUAL_TLS].as_str() {
                "" | "off" => MutualTls::Off,
                "verify" => MutualTls::Verify,
                "passthrough" => MutualTls::Passthrough,
                value => return Err(format!("unsupported AWS listener mTLS mode {value:?}")),
            };
            Ok(Some(Component::Listener(Listener {
                presentation,
                protocol,
                port: number as u16,
                mutual_tls,
                certificate: strings[STRING_AWS_CERTIFICATE].clone(),
                trust_store: strings[STRING_AWS_TRUST_STORE].clone(),
                target_group: strings[STRING_AWS_TARGET_GROUP].clone(),
                backend_tls,
                backend_mtls,
                show_title,
            })))
        }
        _ => decode_aws_feature(strings, presentation).map(|f| Some(Component::Feature(f))),
    }
}

include!("deserialize_aws_feature.rs");
