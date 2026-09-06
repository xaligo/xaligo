use crate::ent::model::aws::{
    alb::Alb,
    listener::{Listener, MutualTls, Protocol},
    nlb::Nlb,
    Component,
};

#[test]
fn alb_option_visibility_removes_geometry_without_mutating_source() {
    use crate::ent::model::aws::{
        feature::{Feature, FeatureKind},
        option::OptionSetting,
    };
    let mut input = aws_document(false, vec![aws_listener(Protocol::Https, 443)]);
    if let Some(Component::Alb(a)) = &mut input.elements[0].aws {
        a.presentation = crate::usc::aws::presentation::parse("detailed", "", "").unwrap();
    }
    let mut option = element("idle-timeout", Concept::Group, Some(0));
    option.aws = Some(Component::Feature(Feature {
        presentation: crate::usc::aws::presentation::parse("", "", "").unwrap(),
        kind: FeatureKind::Option(OptionSetting {
            key: crate::usc::aws::option::key("idle_timeout.timeout_seconds").unwrap(),
            value: "60".into(),
            name: String::new(),
        }),
    }));
    input.elements.push(option);
    let detailed = resolve(&input).unwrap();
    let before = detailed.elements.iter().find(|e| e.id == "lb").unwrap();
    assert!(detailed
        .elements
        .iter()
        .any(|e| e.text.value == "Idle timeout (s)"));
    input.elements[2].visual.visible = Some(false);
    let hidden = resolve(&input).unwrap();
    let after = hidden.elements.iter().find(|e| e.id == "lb").unwrap();
    assert!(after.width < before.width || after.height < before.height);
    assert!(!hidden
        .elements
        .iter()
        .any(|e| e.text.value.contains("Idle timeout")));
    assert!(input.elements[2].aws.is_some());
    // Hidden annotations still enforce the closed schema and ranges.
    if let Some(Component::Feature(f)) = &mut input.elements[2].aws {
        if let FeatureKind::Option(o) = &mut f.kind {
            o.value = "0".into();
        }
    }
    assert!(resolve(&input).is_err());
}

#[test]
fn alb_presentation_rejects_unknown_and_conflicting_controls() {
    let parse = crate::usc::aws::presentation::parse;
    assert!(parse("summary", "rules, actions", "options").is_ok());
    assert!(parse("compact", "", "").is_err());
    assert!(parse("standard", "rules", "rules").is_err());
    assert!(parse("detailed", "typo", "").is_err());
}

fn aws_listener(protocol: Protocol, port: u16) -> Listener {
    Listener {
        presentation: crate::usc::aws::presentation::parse("", "", "").unwrap(),
        protocol,
        port,
        mutual_tls: MutualTls::Off,
        certificate: String::new(),
        trust_store: String::new(),
        target_group: "app".into(),
        backend_tls: None,
        backend_mtls: None,
        show_title: None,
    }
}

fn aws_document(nlb: bool, models: Vec<Listener>) -> DocumentSpec {
    let mut owner = element("lb", Concept::Group, None);
    owner.aws = Some(if nlb {
        Component::Nlb(Nlb {
            domain: "api.example.test".into(),
            presentation: crate::usc::aws::presentation::parse("", "", "").unwrap(),
        })
    } else {
        Component::Alb(Alb {
            domain: "api.example.test".into(),
            presentation: crate::usc::aws::presentation::parse("", "", "").unwrap(),
        })
    });
    let mut elements = vec![owner];
    for (i, model) in models.into_iter().enumerate() {
        let mut child = element(&format!("listener-{i}"), Concept::Group, Some(0));
        child.aws = Some(Component::Listener(model));
        elements.push(child);
    }
    let mut input = document(elements);
    input.width = 1800.0;
    input.height = 1400.0;
    input
}

#[test]
fn aws_compact_alb_keeps_listener_identity_and_has_no_watermark() {
    let input = aws_document(
        false,
        vec![
            aws_listener(Protocol::Http, 80),
            aws_listener(Protocol::Https, 443),
            aws_listener(Protocol::Https, 8443),
        ],
    );
    let resolved = resolve(&input).unwrap();
    assert!(resolved.elements[0].width < 600.0);
    assert_eq!(resolved.elements[1].id, "listener-0");
    assert!(resolved.elements[2].x >= resolved.elements[1].x + resolved.elements[1].width + 12.0);
    let values: Vec<_> = resolved
        .elements
        .iter()
        .map(|e| e.text.value.as_str())
        .collect();
    assert!(values.contains(&"TLS OFF"));
    assert!(values.contains(&"TLS ON"));
    assert!(values.contains(&"mTLS OFF"));
    assert!(values.contains(&"Listener · HTTPS:443"), "{values:?}");
    assert!(!values
        .iter()
        .any(|s| s.contains("watermark") || s.contains("TERMINATION")));
    assert_eq!(
        input.elements.len(),
        4,
        "composition must not mutate source"
    );
    assert!(input.elements[0].aws.is_some());
}

#[test]
fn aws_tcp_passthrough_distinguishes_backend_authentication() {
    let mut listener = aws_listener(Protocol::Tcp, 443);
    listener.backend_tls = Some(true);
    listener.backend_mtls = Some(true);
    let resolved = resolve(&aws_document(true, vec![listener])).unwrap();
    let values: Vec<_> = resolved
        .elements
        .iter()
        .map(|e| e.text.value.as_str())
        .collect();
    for expected in ["TLS OFF", "mTLS OFF", "Listener · TCP:443"] {
        assert!(values.contains(&expected));
    }
    assert!(!values.iter().any(|value| value.starts_with("Target ")));
    let tls = resolved
        .elements
        .iter()
        .find(|e| e.id == "listener-0::aws/tls")
        .unwrap();
    let mtls = resolved
        .elements
        .iter()
        .find(|e| e.id == "listener-0::aws/mtls")
        .unwrap();
    assert_eq!(tls.x, mtls.x);
    assert!(mtls.y >= tls.y + tls.height);
}

#[test]
fn aws_service_specific_security_combinations_are_checked() {
    for (nlb, mut model) in [
        (false, aws_listener(Protocol::Tcp, 443)),
        (true, aws_listener(Protocol::Https, 443)),
    ] {
        assert!(resolve(&aws_document(nlb, vec![model.clone()])).is_err());
        model.port = 0;
        assert!(resolve(&aws_document(nlb, vec![model])).is_err());
    }
    let mut model = aws_listener(Protocol::Https, 443);
    model.mutual_tls = MutualTls::Verify;
    assert!(resolve(&aws_document(false, vec![model.clone()])).is_err());
    model.trust_store = "client-ca".into();
    assert!(resolve(&aws_document(false, vec![model.clone()])).is_ok());
    model.mutual_tls = MutualTls::Passthrough;
    assert!(resolve(&aws_document(false, vec![model])).is_err());
    let mut model = aws_listener(Protocol::Tls, 443);
    model.backend_tls = Some(true);
    model.backend_mtls = Some(true);
    assert!(resolve(&aws_document(true, vec![model])).is_err());
    let mut model = aws_listener(Protocol::Tcp, 443);
    model.certificate = "cert".into();
    assert!(resolve(&aws_document(true, vec![model])).is_err());
}

#[test]
fn aws_component_validates_ownership_size_duplicates_and_visibility() {
    let mut input = aws_document(false, vec![aws_listener(Protocol::Https, 443)]);
    input.elements[1].parent = None;
    assert!(resolve(&input).is_err());
    input.elements[1].parent = Some(0);
    input.elements[0].width = Some(100.0);
    assert!(resolve(&input).is_err());
    input.elements[0].width = Some(350.0);
    input.elements[0].visual.visible = Some(false);
    assert!(resolve(&input)
        .unwrap()
        .elements
        .iter()
        .all(|e| !e.visual.visible));
    let mut duplicate = input.elements[1].clone();
    duplicate.id = "second".into();
    input.elements.push(duplicate);
    assert!(resolve(&input).is_err());
}

#[test]
fn aws_domain_width_and_listener_wrapping_are_deterministic() {
    let mut input = aws_document(
        false,
        (0..4)
            .map(|n| aws_listener(Protocol::Https, 443 + n))
            .collect(),
    );
    input.elements[0].aws = Some(Component::Alb(Alb {
        domain: "API基盤.example.test".into(),
        presentation: crate::usc::aws::presentation::parse("", "", "").unwrap(),
    }));
    input.elements[0].width = Some(360.0);
    let resolved = resolve(&input).unwrap();
    let domain = resolved
        .elements
        .iter()
        .find(|e| e.id == "lb::aws/domain")
        .unwrap();
    assert_eq!(
        domain.width,
        presentation_text_width("API基盤.example.test", 14.0).ceil() + 20.0
    );
    assert_eq!(resolved.elements[3].y - resolved.elements[1].y, 140.0);
    assert!(resolved.elements[4].y + resolved.elements[4].height < resolved.elements[0].height);
}

#[test]
fn aws_auto_size_tracks_content_and_hidden_title_without_stretching() {
    let mut model = aws_listener(Protocol::Https, 443);
    model.mutual_tls = MutualTls::Verify;
    model.trust_store = "ca".into();
    let visible = resolve(&aws_document(false, vec![model.clone()])).unwrap();
    for name in ["tls", "mtls"] {
        let part = visible
            .elements
            .iter()
            .find(|e| e.id == format!("listener-0::aws/{name}"))
            .unwrap();
        assert_eq!(part.visual.fill, "#dcfce7");
    }
    assert!(visible
        .elements
        .iter()
        .any(|e| e.text.value == "Listener · HTTPS:443"));
    model.show_title = Some(false);
    let hidden = resolve(&aws_document(false, vec![model.clone()])).unwrap();
    assert!(!hidden
        .elements
        .iter()
        .any(|e| e.text.value.contains("Listener")));
    assert!(hidden.elements.iter().any(|e| e.text.value == "HTTPS :443"));
    assert_eq!(visible.elements[1].height, hidden.elements[1].height);
    assert_eq!(visible.elements[0].height, hidden.elements[0].height);
    let mut fixed = aws_document(false, vec![model]);
    fixed.elements[0].width = Some(600.0);
    fixed.elements[0].height = Some(400.0);
    let expanded = resolve(&fixed).unwrap();
    assert_eq!(expanded.elements[0].width, 600.0);
    assert_eq!(expanded.elements[0].height, 400.0);
    assert_eq!(expanded.elements[1].width, hidden.elements[1].width);
    assert_eq!(expanded.elements[1].height, hidden.elements[1].height);
}
