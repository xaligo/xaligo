use crate::ent::model::aws::presentation::Presentation;

pub const PARTS: &[&str] = &[
    "title",
    "protocol",
    "tls",
    "mtls",
    "certificate",
    "trust-store",
    "target-group",
    "icon",
    "domain",
    "rules",
    "conditions",
    "actions",
    "transforms",
    "options",
    "targets",
    "priority",
    "values",
    "weights",
    "name",
    "type",
    "services",
];
pub fn parse(level: &str, show: &str, hide: &str) -> Result<Presentation, String> {
    if !matches!(level, "" | "summary" | "standard" | "detailed") {
        return Err("detail-level must be summary, standard, or detailed".into());
    }
    fn bits(input: &str) -> Result<u32, String> {
        let mut mask = 0;
        for name in input.split(',').filter(|s| !s.is_empty()) {
            let index = PARTS
                .iter()
                .position(|p| *p == name.trim())
                .ok_or_else(|| format!("unknown AWS display part {name:?}"))?;
            mask |= 1 << index;
        }
        Ok(mask)
    }
    let show = bits(show)?;
    let hide = bits(hide)?;
    if show & hide != 0 {
        return Err("AWS show and hide cannot contain the same part".into());
    }
    Ok(Presentation {
        level: level.into(),
        show,
        hide,
    })
}
pub(super) fn resolve(p: &Presentation, inherited: u32) -> u32 {
    let base = match p.level.as_str() {
        "detailed" => (1 << PARTS.len()) - 1,
        "standard" => ((1 << PARTS.len()) - 1) & !(1 << 13),
        "summary" => ["protocol", "tls", "mtls", "icon", "domain", "name", "type"]
            .iter()
            .fold(0, |m, n| m | bit(n)),
        _ => inherited,
    };
    (base | p.show) & !p.hide
}
pub(super) fn bit(name: &str) -> u32 {
    1 << PARTS.iter().position(|p| *p == name).unwrap()
}
pub(super) fn has(mask: u32, name: &str) -> bool {
    mask & bit(name) != 0
}
pub(super) fn configured(p: &Presentation) -> bool {
    !p.level.is_empty() || p.show != 0 || p.hide != 0
}
