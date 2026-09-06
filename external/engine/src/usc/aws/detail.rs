use super::{
    action, condition, drawing, option, option_table, presentation, rule, rule_table, target_group,
    target_service, transform,
};
use crate::ent::model::{
    aws::{
        feature::{Feature, FeatureKind},
        listener::{MutualTls, Protocol},
        presentation::Presentation,
        Component,
    },
    document::*,
};
use crate::util::error::LayoutError;

pub(super) fn children(doc: &DocumentSpec, index: usize) -> Vec<usize> {
    doc.elements
        .iter()
        .enumerate()
        .filter_map(|(i, e)| (e.parent == Some(index)).then_some(i))
        .collect()
}
pub(super) fn feature(doc: &DocumentSpec, index: usize) -> Option<&FeatureKind> {
    match doc.elements[index].aws.as_ref() {
        Some(Component::Feature(f)) => Some(&f.kind),
        _ => None,
    }
}
pub(super) fn option_value<'a>(doc: &'a DocumentSpec, index: usize, name: &str) -> Option<&'a str> {
    children(doc, index)
        .into_iter()
        .find_map(|i| match feature(doc, i) {
            Some(FeatureKind::Option(o)) if option::DEFINITIONS[o.key].name == name => {
                Some(o.value.as_str())
            }
            _ => None,
        })
}
fn display(component: &Component) -> &Presentation {
    match component {
        Component::Alb(m) => &m.presentation,
        Component::Nlb(m) => &m.presentation,
        Component::Listener(m) => &m.presentation,
        Component::Feature(m) => &m.presentation,
    }
}
fn descendant(doc: &DocumentSpec, index: usize, root: usize) -> bool {
    let mut node = Some(index);
    while let Some(i) = node {
        if i == root {
            return true;
        }
        node = doc.elements[i].parent;
    }
    false
}
pub(super) fn required(doc: &DocumentSpec, root: usize) -> bool {
    doc.elements.iter().enumerate().any(|(i, e)| {
        descendant(doc, i, root)
            && e.aws.as_ref().is_some_and(|a| {
                matches!(a, Component::Feature(_)) || presentation::configured(display(a))
            })
    })
}
fn owner_name(doc: &DocumentSpec, index: usize) -> &'static str {
    match doc.elements[index].aws.as_ref() {
        Some(Component::Alb(_)) => "alb",
        Some(Component::Nlb(_)) => "nlb",
        Some(Component::Listener(_)) => "listener",
        Some(Component::Feature(Feature {
            kind: FeatureKind::Action(a),
            ..
        })) => action::owner(&a.kind),
        Some(Component::Feature(Feature {
            kind: FeatureKind::TargetGroup(_),
            ..
        })) => "tg",
        Some(Component::Feature(Feature {
            kind: FeatureKind::Target(_),
            ..
        })) => "target",
        Some(Component::Feature(Feature {
            kind: FeatureKind::TargetService(m),
            ..
        })) => target_service::owner(&m.kind),
        _ => "",
    }
}
fn listener_protocol(doc: &DocumentSpec, index: usize) -> Protocol {
    let mut i = Some(index);
    while let Some(n) = i {
        if let Some(Component::Listener(m)) = &doc.elements[n].aws {
            return m.protocol;
        }
        i = doc.elements[n].parent;
    }
    Protocol::Http
}
pub(super) fn validate(doc: &DocumentSpec) -> Result<(), LayoutError> {
    for (index, element) in doc.elements.iter().enumerate() {
        let Some(component) = &element.aws else {
            continue;
        };
        let p = display(component);
        if !matches!(p.level.as_str(), "" | "summary" | "standard" | "detailed")
            || p.show & p.hide != 0
            || (p.show | p.hide) >> presentation::PARTS.len() != 0
        {
            return Err(LayoutError::new("invalid AWS presentation parameters"));
        }
        let parent = element.parent.and_then(|i| doc.elements[i].aws.as_ref());
        match component {
            Component::Alb(_) | Component::Nlb(_) => {
                let domain = match component {
                    Component::Alb(m) => m.domain.as_str(),
                    Component::Nlb(m) => m.domain.as_str(),
                    _ => "",
                };
                if domain.len() > 1024 || domain.chars().any(char::is_control) {
                    return Err(LayoutError::new(
                        "AWS domain must be single-line and at most 1024 bytes",
                    ));
                }
                let mut ports = std::collections::HashSet::new();
                for i in children(doc, index) {
                    match &doc.elements[i].aws {
                        Some(Component::Listener(m)) => {
                            if matches!(component, Component::Alb(_)) {
                                super::alb::validate(m)?;
                            } else {
                                super::nlb::validate(m)?;
                            }
                            if !ports.insert(m.port) {
                                return Err(LayoutError::new("AWS listener ports must be unique"));
                            }
                        }
                        Some(Component::Feature(Feature {
                            kind: FeatureKind::Option(_) | FeatureKind::TargetGroup(_),
                            ..
                        })) if matches!(component, Component::Alb(_)) => {}
                        _ => {
                            return Err(LayoutError::new(
                                "load balancer accepts listeners and ALB options/target groups",
                            ));
                        }
                    }
                }
                if ports.is_empty() || ports.len() > 32 {
                    return Err(LayoutError::new("load balancer requires 1..32 listeners"));
                }
            }
            Component::Listener(m) => {
                if !matches!(parent, Some(Component::Alb(_) | Component::Nlb(_))) {
                    return Err(LayoutError::new("listener requires ALB/NLB parent"));
                }
                if matches!(parent, Some(Component::Alb(_))) {
                    target_group::validate_reference(
                        doc,
                        index,
                        &m.target_group,
                        m.protocol,
                        false,
                    )?;
                    if !matches!(m.protocol, Protocol::Https)
                        && [
                            "tls-policy",
                            "sni-certificates",
                            "mtls-ignore-expiry",
                            "mtls-advertise-ca",
                            "trust-store-ca-bundle",
                            "trust-store-crls",
                        ]
                        .iter()
                        .any(|name| option_value(doc, index, name).is_some())
                    {
                        return Err(LayoutError::new(
                            "TLS and client trust options require an HTTPS listener",
                        ));
                    }
                }
                let mut priorities = std::collections::HashSet::new();
                for i in children(doc, index) {
                    match feature(doc, i) {
                        Some(FeatureKind::Rule(r)) => {
                            if !priorities.insert(r.priority) {
                                return Err(LayoutError::new(
                                    "listener rule priorities must be unique (including default)",
                                ));
                            }
                            if r.priority.is_none() && !m.target_group.is_empty() {
                                return Err(LayoutError::new(
                                    "use target-group shorthand or an explicit default rule, not both",
                                ));
                            }
                        }
                        Some(FeatureKind::Option(_)) => {}
                        _ => return Err(LayoutError::new("listener accepts rules and options")),
                    }
                }
            }
            Component::Feature(f) => {
                let mut root = element.parent;
                let mut alb = false;
                while let Some(i) = root {
                    if matches!(doc.elements[i].aws, Some(Component::Alb(_))) {
                        alb = true;
                        break;
                    }
                    if matches!(doc.elements[i].aws, Some(Component::Nlb(_))) {
                        break;
                    }
                    root = doc.elements[i].parent;
                }
                if !alb {
                    return Err(LayoutError::new(
                        "ALB feature requires an ALB component ancestor (not NLB)",
                    ));
                }
                let correct = match &f.kind {
                    FeatureKind::Rule(_) => matches!(parent, Some(Component::Listener(_))),
                    FeatureKind::Condition(_)
                    | FeatureKind::Action(_)
                    | FeatureKind::Transform(_) => matches!(
                        parent,
                        Some(Component::Feature(Feature {
                            kind: FeatureKind::Rule(_),
                            ..
                        }))
                    ),
                    FeatureKind::Match(_) => matches!(
                        parent,
                        Some(Component::Feature(Feature {
                            kind: FeatureKind::Condition(_) | FeatureKind::JwtClaim(_),
                            ..
                        }))
                    ),
                    FeatureKind::JwtClaim(_) => {
                        matches!(parent,Some(Component::Feature(Feature {kind:FeatureKind::Action(a),..})) if matches!(a.kind,crate::ent::model::aws::action::ActionKind::Jwt))
                    }
                    FeatureKind::ForwardTarget(_) => matches!(
                        parent,
                        Some(Component::Feature(Feature {
                            kind: FeatureKind::Action(_),
                            ..
                        }))
                    ),
                    FeatureKind::Rewrite(_) => matches!(
                        parent,
                        Some(Component::Feature(Feature {
                            kind: FeatureKind::Transform(_),
                            ..
                        }))
                    ),
                    FeatureKind::TargetGroup(_) => matches!(parent, Some(Component::Alb(_))),
                    FeatureKind::TargetService(_) => matches!(
                        parent,
                        Some(Component::Feature(Feature {
                            kind: FeatureKind::TargetGroup(_),
                            ..
                        }))
                    ),
                    FeatureKind::Target(_) => matches!(
                        parent,
                        Some(Component::Feature(Feature {
                            kind: FeatureKind::TargetGroup(_) | FeatureKind::TargetService(_),
                            ..
                        }))
                    ),
                    FeatureKind::Option(_) => {
                        !owner_name(doc, element.parent.unwrap_or(index)).is_empty()
                    }
                };
                if !correct {
                    return Err(LayoutError::new(format!(
                        "invalid parent for AWS feature {:?}",
                        element.id
                    )));
                }
                match &f.kind {
                    FeatureKind::Rule(m) => rule::validate(doc, index, m)?,
                    FeatureKind::Condition(m) => condition::validate(doc, index, m)?,
                    FeatureKind::Action(m) => {
                        action::validate(doc, index, m, listener_protocol(doc, index))?
                    }
                    FeatureKind::JwtClaim(m) => {
                        if matches!(m.name.as_str(), "exp" | "iss" | "nbf" | "iat") {
                            return Err(LayoutError::new(
                                "JWT standard claims are validated automatically",
                            ));
                        }
                        let items = children(doc, index);
                        if m.name.is_empty()
                            || !matches!(
                                m.format.as_str(),
                                "single-string" | "string-array" | "space-separated-values"
                            )
                            || items.is_empty()
                            || items.len() > 10
                            || m.format == "single-string" && items.len() != 1
                        {
                            return Err(LayoutError::new(
                                "invalid JWT claim name, format or value count",
                            ));
                        }
                        for i in items {
                            if !matches!(feature(doc,i),Some(FeatureKind::Match(value)) if !value.value.is_empty() && value.value.len()<=256 && !value.regex && value.key.is_empty() && (m.format!="space-separated-values" || !value.value.contains(' ')))
                            {
                                return Err(LayoutError::new(
                                    "JWT claim requires literal values without spaces for space-separated-values",
                                ));
                            }
                        }
                    }
                    FeatureKind::Transform(_) => transform::validate(doc, index)?,
                    FeatureKind::TargetGroup(m) => target_group::validate(doc, index, m)?,
                    FeatureKind::TargetService(m) => target_service::validate(doc, index, m)?,
                    FeatureKind::Target(m) if m.name.is_empty() => {
                        return Err(LayoutError::new("registered target requires name"));
                    }
                    FeatureKind::Option(m) => {
                        option::validate(m, owner_name(doc, element.parent.unwrap()))?
                    }
                    _ => {}
                }
                let allowed_children: &[&str] = match f.kind {
                    FeatureKind::Rule(_) => &["rule"], // Child ownership is checked above.
                    FeatureKind::Condition(_)
                    | FeatureKind::Transform(_)
                    | FeatureKind::JwtClaim(_) => &["nested"],
                    FeatureKind::Action(_)
                    | FeatureKind::TargetGroup(_)
                    | FeatureKind::TargetService(_)
                    | FeatureKind::Target(_) => &["options"],
                    _ => &[],
                };
                if allowed_children.is_empty() && !children(doc, index).is_empty() {
                    return Err(LayoutError::new("AWS feature must be empty"));
                }
                if children(doc, index)
                    .iter()
                    .any(|i| doc.elements[*i].aws.is_none())
                {
                    return Err(LayoutError::new(
                        "AWS feature children must be dedicated AWS tags",
                    ));
                }
            }
        }
        if !matches!(component, Component::Alb(_) | Component::Nlb(_))
            && (element.width.is_some()
                || element.height.is_some()
                || element.x.is_some()
                || element.y.is_some()
                || element.weight.is_some())
        {
            return Err(LayoutError::new(
                "AWS child geometry is controlled by the native component",
            ));
        }
        let mut options = std::collections::HashSet::new();
        for i in children(doc, index) {
            if let Some(FeatureKind::Option(o)) = feature(doc, i) {
                if !options.insert((o.key, o.name.as_str())) {
                    return Err(LayoutError::new("duplicate ALB option"));
                }
            }
        }
    }
    Ok(())
}

struct Row {
    key: String,
    text: String,
    font: f64,
    fill: &'static str,
}
struct Card {
    index: usize,
    width: f64,
    height: f64,
    rows: Vec<Row>,
    children: Vec<Card>,
    hidden: Vec<usize>,
    columns: usize,
    root: bool,
    mask: u32,
    header: f64,
    table: option_table::Table,
    rules: rule_table::Table,
}
fn category(c: &Component) -> Option<&'static str> {
    match c {
        Component::Feature(Feature {
            kind: FeatureKind::Rule(_),
            ..
        }) => Some("rules"),
        Component::Feature(Feature {
            kind: FeatureKind::Condition(_),
            ..
        }) => Some("conditions"),
        Component::Feature(Feature {
            kind: FeatureKind::Action(_),
            ..
        }) => Some("actions"),
        Component::Feature(Feature {
            kind: FeatureKind::Transform(_),
            ..
        }) => Some("transforms"),
        Component::Feature(Feature {
            kind: FeatureKind::TargetGroup(_) | FeatureKind::Target(_),
            ..
        }) => Some("targets"),
        Component::Feature(Feature {
            kind: FeatureKind::TargetService(_),
            ..
        }) => Some("services"),
        Component::Feature(Feature {
            kind: FeatureKind::Option(_),
            ..
        }) => Some("options"),
        Component::Feature(Feature {
            kind: FeatureKind::Match(_) | FeatureKind::Rewrite(_),
            ..
        }) => Some("values"),
        _ => None,
    }
}
fn add(rows: &mut Vec<Row>, key: &str, value: String, font: f64, fill: &'static str) {
    // Width-aware wrapping is resolved here, identically for SVG and PPTX.
    let mut line = String::new();
    for c in value.chars() {
        line.push(c);
        if crate::usc::layout::presentation_text_width(&line, font) > 260.0 {
            line.pop();
            rows.push(Row {
                key: format!("{key}-{}", rows.len()),
                text: line,
                font,
                fill,
            });
            line = c.to_string();
        }
    }
    if !line.is_empty() {
        rows.push(Row {
            key: format!("{key}-{}", rows.len()),
            text: line,
            font,
            fill,
        });
    }
}
fn build(
    doc: &DocumentSpec,
    index: usize,
    inherited: u32,
    root: bool,
) -> Result<Card, LayoutError> {
    let element = &doc.elements[index];
    let component = element
        .aws
        .as_ref()
        .ok_or_else(|| LayoutError::new("missing AWS child model"))?;
    let mask = presentation::resolve(display(component), inherited);
    let has = |name| presentation::has(mask, name);
    let mut rows = Vec::new();
    match component {
        Component::Listener(m) => {
            let title = has("title") && m.show_title != Some(false);
            if title || has("protocol") {
                add(
                    &mut rows,
                    "header",
                    format!(
                        "{}{}",
                        if title { "Listener · " } else { "" },
                        if has("protocol") {
                            format!("{} :{}", super::listener::protocol_name(m.protocol), m.port)
                        } else {
                            String::new()
                        }
                    ),
                    if has("protocol") { 16.0 } else { 12.0 },
                    "none",
                );
            }
            if has("tls") {
                let active = super::listener::tls(m.protocol);
                add(
                    &mut rows,
                    "tls",
                    if active { "TLS ON" } else { "TLS OFF" }.into(),
                    12.0,
                    if active { "#dcfce7" } else { "#e8eef6" },
                );
            }
            if has("mtls") {
                add(
                    &mut rows,
                    "mtls",
                    match m.mutual_tls {
                        MutualTls::Off => "mTLS OFF",
                        MutualTls::Verify => "mTLS ON",
                        MutualTls::Passthrough => "mTLS PASS",
                    }
                    .into(),
                    12.0,
                    if matches!(m.mutual_tls, MutualTls::Off) {
                        "#e8eef6"
                    } else {
                        "#dcfce7"
                    },
                );
            }
            let mut refs = Vec::new();
            if has("certificate") && !m.certificate.is_empty() {
                refs.push(format!("Cert {}", m.certificate));
            }
            if has("trust-store") && !m.trust_store.is_empty() {
                refs.push(format!("CA {}", m.trust_store));
            }
            if has("target-group") && !m.target_group.is_empty() {
                refs.push(format!("→ {}", m.target_group));
            }
            if !refs.is_empty() {
                add(&mut rows, "references", refs.join(" · "), 12.0, "none");
            }
        }
        Component::Feature(f) => match &f.kind {
            FeatureKind::Rule(m) => {
                if has("priority") {
                    add(
                        &mut rows,
                        "priority",
                        m.priority
                            .map_or("Default".into(), |p| format!("#{p} · AND")),
                        13.0,
                        "#ede9fe",
                    );
                }
            }
            FeatureKind::Condition(m) => {
                if has("type") {
                    add(
                        &mut rows,
                        "condition",
                        format!(
                            "{}{} · OR",
                            condition::label(&m.kind),
                            if m.name.is_empty() {
                                String::new()
                            } else {
                                format!(" {}", m.name)
                            }
                        ),
                        12.0,
                        "#eff6ff",
                    );
                }
            }
            FeatureKind::Match(m) => {
                if has("values") {
                    add(
                        &mut rows,
                        "match",
                        format!(
                            "{}{}{}",
                            if m.regex { "regex: " } else { "" },
                            if m.key.is_empty() {
                                String::new()
                            } else {
                                format!("{}=", m.key)
                            },
                            m.value
                        ),
                        12.0,
                        "none",
                    );
                }
            }
            FeatureKind::Action(m) => {
                if has("type") {
                    add(
                        &mut rows,
                        "action",
                        format!(
                            "{}. {}{}",
                            m.order,
                            action::label(&m.kind),
                            if has("target-group") && !m.target_group.is_empty() {
                                format!(" · → {}", m.target_group)
                            } else {
                                String::new()
                            }
                        ),
                        13.0,
                        "#fff7ed",
                    );
                }
                if !has("type") && has("target-group") && !m.target_group.is_empty() {
                    add(
                        &mut rows,
                        "target-group",
                        format!("→ {}", m.target_group),
                        12.0,
                        "none",
                    );
                }
            }
            FeatureKind::ForwardTarget(m) => {
                let value = format!(
                    "{}{}",
                    if has("target-group") {
                        format!("→ {}", m.target_group)
                    } else {
                        String::new()
                    },
                    if has("weights") {
                        format!(" · w={}", m.weight)
                    } else {
                        String::new()
                    }
                );
                add(&mut rows, "target", value, 12.0, "#fff7ed");
            }
            FeatureKind::JwtClaim(m) => {
                if has("name") {
                    add(
                        &mut rows,
                        "claim",
                        format!("{} · {}", m.name, m.format),
                        12.0,
                        "#eff6ff",
                    );
                }
            }
            FeatureKind::Transform(m) => {
                if has("type") {
                    add(
                        &mut rows,
                        "transform",
                        transform::label(&m.kind).into(),
                        12.0,
                        "#f3e8ff",
                    );
                }
            }
            FeatureKind::Rewrite(m) => {
                if has("values") {
                    add(
                        &mut rows,
                        "rewrite",
                        format!("{} → {}", m.regex, m.replacement),
                        12.0,
                        "none",
                    );
                }
            }
            FeatureKind::TargetGroup(m) => {
                if has("name") || has("protocol") {
                    let mut header = Vec::new();
                    if has("name") {
                        header.push(format!("TG / {}", m.name));
                    }
                    if has("protocol") && (!m.protocol.is_empty() || m.port.is_some()) {
                        header.push(format!(
                            "{}{}",
                            m.protocol,
                            m.port.map_or(String::new(), |p| format!(":{p}"))
                        ));
                    }
                    if has("protocol") && !m.target_type.is_empty() {
                        header.push(m.target_type.clone());
                    }
                    add(&mut rows, "header", header.join(" · "), 14.0, "#fff7ed");
                }
            }
            FeatureKind::Target(m) => {
                if has("name") {
                    add(
                        &mut rows,
                        "name",
                        format!(
                            "{}{}{}",
                            m.name,
                            m.port.map_or(String::new(), |p| format!(":{p}")),
                            if has("values") && !m.zone.is_empty() {
                                format!(" · {}", m.zone)
                            } else {
                                String::new()
                            }
                        ),
                        12.0,
                        "none",
                    );
                }
            }
            FeatureKind::TargetService(m) => {
                let title = format!(
                    "{}{}{}",
                    if has("type") {
                        format!("{} / ", target_service::label(&m.kind))
                    } else {
                        String::new()
                    },
                    if has("name") { m.name.as_str() } else { "" },
                    if has("values") && !m.reference.is_empty() {
                        format!(" · ↗ {}", m.reference)
                    } else {
                        String::new()
                    }
                );
                add(&mut rows, "service", title, 14.0, "none");
            }
            FeatureKind::Option(_) => {}
        },
        _ => {}
    }
    let mut list = children(doc, index);
    if matches!(component, Component::Listener(_)) {
        list.sort_by_key(|i| match feature(doc, *i) {
            Some(FeatureKind::Rule(r)) => u32::from(r.priority.unwrap_or(u16::MAX)),
            _ => u32::MAX,
        });
    } else if matches!(feature(doc, index), Some(FeatureKind::Rule(_))) {
        list.sort_by_key(|i| match feature(doc, *i) {
            Some(FeatureKind::Condition(_)) => 0,
            Some(FeatureKind::Transform(_)) => 1,
            Some(FeatureKind::Action(a)) => 2 + u32::from(a.order),
            _ => u32::MAX,
        });
    }
    let mut children_cards = Vec::new();
    let mut hidden = Vec::new();
    for i in list {
        let child = &doc.elements[i];
        let gate = child
            .aws
            .as_ref()
            .and_then(category)
            .is_none_or(|part| has(part));
        if !gate || child.visual.visible == Some(false) {
            hidden.push(i);
        } else if matches!(component, Component::Listener(_))
            && matches!(feature(doc, i), Some(FeatureKind::Rule(_)))
        {
            // Listener rules are rendered as a compact table below the header.
        } else if !matches!(feature(doc, i), Some(FeatureKind::Option(_))) {
            children_cards.push(build(doc, i, mask, false)?);
        }
    }
    let table = option_table::measure(doc, index, mask);
    let rules = if matches!(component, Component::Listener(_)) {
        rule_table::measure(doc, index, mask)
    } else {
        rule_table::measure(doc, index, 0)
    };
    let text_width = rows
        .iter()
        .map(|r| crate::usc::layout::presentation_text_width(&r.text, r.font) + 20.0)
        .fold(80.0_f64, f64::max)
        .ceil();
    let largest = children_cards
        .iter()
        .map(|c| c.width)
        .fold(0.0_f64, f64::max);
    let service_icon = matches!(feature(doc, index), Some(FeatureKind::TargetService(_)))
        && has("icon")
        && !element.icon.reference.is_empty();
    let mut width = (text_width + if service_icon { 32.0 } else { 0.0 })
        .max(largest + 20.0)
        .max(rules.width + 20.0)
        .max(table.width + 20.0);
    let mut header = rows.len() as f64 * 24.0 + 16.0;
    let mut columns = 1;
    if root {
        let (domain, kind) = match component {
            Component::Alb(m) => (m.domain.as_str(), "ALB"),
            Component::Nlb(m) => (m.domain.as_str(), "NLB"),
            _ => ("", ""),
        };
        let label = if domain.is_empty() { kind } else { domain };
        let header_width = (if has("domain") {
            crate::usc::layout::presentation_text_width(label, 14.0) + 20.0
        } else {
            0.0
        }) + (if has("icon") { 52.0 } else { 12.0 })
            + 12.0;
        columns = if children_cards.iter().any(|card| card.rules.height > 0.0) {
            1
        } else {
            children_cards.len().clamp(1, 3)
        };
        let packed_width = children_cards
            .chunks(columns)
            .map(|row| {
                row.iter().map(|card| card.width).sum::<f64>()
                    + 12.0 * (row.len().saturating_sub(1) as f64)
                    + 24.0
            })
            .fold(0.0_f64, f64::max);
        width = packed_width
            .max(header_width)
            .max(table.width + 24.0)
            .max(96.0);
        if let Some(explicit) = element.width {
            if explicit < header_width.max(largest + 24.0).max(table.width + 24.0) {
                return Err(LayoutError::new(
                    "ALB width cannot contain visible header/cards",
                ));
            }
            width = explicit;
            columns = (((width - 12.0) / (largest + 12.0)).floor() as usize)
                .max(1)
                .min(children_cards.len().max(1));
        }
        header = if has("icon") || has("domain") {
            56.0
        } else {
            12.0
        };
    }
    let content: f64 = children_cards
        .chunks(columns)
        .map(|row| row.iter().map(|c| c.height).fold(0.0_f64, f64::max) + 8.0)
        .sum();
    let mut height = header
        + content
        + rules.height
        + table.height
        + if children_cards.is_empty() && table.height == 0.0 && rules.height == 0.0 {
            0.0
        } else {
            4.0
        };
    if root {
        if element.weight.is_some() {
            return Err(LayoutError::new("AWS component does not support weight"));
        }
        if let Some(explicit) = element.height {
            if explicit < height {
                return Err(LayoutError::new("ALB height cannot contain visible cards"));
            }
            height = explicit;
        }
    }
    Ok(Card {
        index,
        width,
        height,
        rows,
        children: children_cards,
        hidden,
        columns,
        root,
        mask,
        header,
        table,
        rules,
    })
}
pub(super) fn collapse(
    doc: &mut DocumentSpec,
    index: usize,
    anchor: &str,
    aliases: &mut Vec<(String, String)>,
) {
    let nested = children(doc, index);
    let e = &mut doc.elements[index];
    aliases.push((e.id.clone(), anchor.into()));
    e.concept = Concept::Group;
    e.layout = LayoutPolicy::Absolute;
    e.padding = Insets::default();
    e.margin = Insets::default();
    e.x = Some(0.0);
    e.y = Some(0.0);
    e.width = Some(1.0);
    e.height = Some(1.0);
    e.text.value.clear();
    e.icon.reference.clear();
    e.visual.visible = Some(false);
    for i in nested {
        collapse(doc, i, anchor, aliases);
    }
}
fn draw(
    doc: &mut DocumentSpec,
    card: &Card,
    x: Option<f64>,
    y: Option<f64>,
    aliases: &mut Vec<(String, String)>,
) {
    let original = doc.elements[card.index].clone();
    let e = &mut doc.elements[card.index];
    e.concept = Concept::Group;
    e.layout = LayoutPolicy::Absolute;
    e.padding = Insets::default();
    if !card.root {
        e.margin = Insets::default();
        e.x = x;
        e.y = y;
    }
    e.width = Some(card.width);
    e.height = Some(card.height);
    e.text.value.clear();
    e.icon.reference.clear();
    e.visual.shape = Shape::Rectangle;
    if !card.root || e.visual.fill.is_empty() {
        e.visual.fill = if card.root { "#f8fbff" } else { "#ffffff" }.into();
    }
    if !card.root || e.visual.stroke.is_empty() {
        e.visual.stroke = "#c7d8f4".into();
    }
    e.visual.stroke_width.get_or_insert(1.0);
    e.visual.corner_radius.get_or_insert(8.0);
    let owner = e.clone();
    let service_icon = matches!(
        feature(doc, card.index),
        Some(FeatureKind::TargetService(_))
    ) && presentation::has(card.mask, "icon")
        && !original.icon.reference.is_empty();
    if service_icon {
        let mut icon = drawing::part(&owner, card.index, "service-icon", [8.0, 8.0, 24.0, 24.0]);
        icon.concept = Concept::Item;
        icon.icon = original.icon.clone();
        icon.icon.width = Some(24.0);
        icon.icon.height = Some(24.0);
        doc.elements.push(icon);
    }
    if card.root {
        let component = original.aws.as_ref().unwrap();
        let (domain, kind) = match component {
            Component::Alb(m) => (m.domain.as_str(), "ALB"),
            Component::Nlb(m) => (m.domain.as_str(), "NLB"),
            _ => ("", ""),
        };
        if presentation::has(card.mask, "icon") {
            let mut icon = drawing::part(&owner, card.index, "icon", [12.0, 12.0, 32.0, 32.0]);
            icon.concept = Concept::Item;
            icon.icon = original.icon.clone();
            icon.icon.width = Some(32.0);
            icon.icon.height = Some(32.0);
            doc.elements.push(icon);
        }
        if presentation::has(card.mask, "domain") {
            let label = if domain.is_empty() { kind } else { domain };
            let left = if presentation::has(card.mask, "icon") {
                52.0
            } else {
                12.0
            };
            let mut tag = drawing::badge(
                &owner,
                card.index,
                "domain",
                [
                    left,
                    14.0,
                    crate::usc::layout::presentation_text_width(label, 14.0).ceil() + 20.0,
                    28.0,
                ],
                label,
                false,
            );
            tag.text.font_size = Some(14.0);
            tag.visual.fill = "#e8efff".into();
            doc.elements.push(tag);
        }
    }
    for (i, row) in card.rows.iter().enumerate() {
        let inset = if service_icon && i == 0 { 40.0 } else { 8.0 };
        let mut part = drawing::part(
            &owner,
            card.index,
            &row.key,
            [inset, 8.0 + i as f64 * 24.0, card.width - inset - 8.0, 24.0],
        );
        part.text.value = row.text.clone();
        part.text.font_size = Some(row.font);
        if row.fill != "none" {
            part.visual.shape = Shape::Rectangle;
            part.visual.fill = row.fill.into();
            part.visual.corner_radius = Some(5.0);
            part.text.padding.left = Some(4.0);
            part.text.padding.right = Some(4.0);
        }
        if row.fill == "#dcfce7" {
            part.text.color = "#166534".into();
        }
        doc.elements.push(part);
    }
    for i in &card.hidden {
        collapse(doc, *i, &owner.id, aliases);
    }
    let mut top = card.header;
    for row in card.children.chunks(card.columns) {
        let mut left = if card.root { 12.0 } else { 10.0 };
        for c in row {
            if owner.visual.visible == Some(false) {
                doc.elements[c.index].visual.visible = Some(false);
            }
            doc.elements[c.index].visual.layer = owner.visual.layer;
            draw(doc, c, Some(left), Some(top), aliases);
            if card.root {
                left += c.width + 12.0;
            }
        }
        top += row.iter().map(|c| c.height).fold(0.0_f64, f64::max) + 8.0;
    }
    rule_table::draw(
        doc,
        &card.rules,
        card.index,
        if card.root { 12.0 } else { 10.0 },
        top,
        aliases,
    );
    top += card.rules.height;
    option_table::draw(
        doc,
        &card.table,
        card.index,
        if card.root { 12.0 } else { 10.0 },
        top,
    );
}
pub(super) fn compose(doc: &mut DocumentSpec, index: usize) -> Result<(), LayoutError> {
    let base = presentation::parse("standard", "", "").unwrap();
    let card = build(doc, index, presentation::resolve(&base, 0), true)?;
    let mut aliases = Vec::new();
    draw(doc, &card, None, None, &mut aliases);
    for e in &mut doc.elements {
        let mut remapped = false;
        for (id, anchor) in &aliases {
            if e.line.source == *id {
                e.line.source = anchor.clone();
                remapped = true;
            }
            if e.line.target == *id {
                e.line.target = anchor.clone();
                remapped = true;
            }
        }
        if remapped && e.line.source == e.line.target {
            e.visual.visible = Some(false);
        }
    }
    Ok(())
}
