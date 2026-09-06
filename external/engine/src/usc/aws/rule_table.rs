use super::{action, condition, detail, drawing, presentation, transform};
use crate::ent::model::{aws::feature::FeatureKind, document::*};

struct RuleRow {
    index: usize,
    id: String,
    priority: String,
    conditions: String,
    action: String,
    height: f64,
}

pub(super) struct Table {
    pub width: f64,
    pub height: f64,
    priority_width: f64,
    condition_width: f64,
    action_width: f64,
    rows: Vec<RuleRow>,
}

fn text_width(value: &str) -> f64 {
    crate::usc::layout::presentation_text_width(value, 12.0)
}

fn line_count(value: &str, width: f64) -> usize {
    ((text_width(value) / (width - 16.0).max(1.0)).ceil() as usize).max(1)
}

fn condition_text(doc: &DocumentSpec, index: usize, mask: u32) -> Option<String> {
    let FeatureKind::Condition(model) = detail::feature(doc, index)? else {
        return None;
    };
    if doc.elements[index].visual.visible == Some(false) || !presentation::has(mask, "conditions") {
        return None;
    }
    let mut values = Vec::new();
    for child in detail::children(doc, index) {
        if let Some(FeatureKind::Match(value)) = detail::feature(doc, child) {
            if doc.elements[child].visual.visible != Some(false) {
                values.push(format!(
                    "{}{}{}",
                    if value.regex { "regex:" } else { "" },
                    if value.key.is_empty() {
                        String::new()
                    } else {
                        format!("{}=", value.key)
                    },
                    value.value
                ));
            }
        }
    }
    let name = if model.name.is_empty() {
        String::new()
    } else {
        format!(" {}", model.name)
    };
    Some(format!(
        "{}{}: {}",
        condition::label(&model.kind),
        name,
        values.join(" OR ")
    ))
}

fn action_text(doc: &DocumentSpec, index: usize, mask: u32) -> Option<String> {
    if doc.elements[index].visual.visible == Some(false) {
        return None;
    }
    match detail::feature(doc, index)? {
        FeatureKind::Action(model) if presentation::has(mask, "actions") => {
            let mut value = format!("{}. {}", model.order, action::label(&model.kind));
            if !model.target_group.is_empty() {
                value.push_str(&format!(" → {}", model.target_group));
            }
            let weighted: Vec<_> = detail::children(doc, index)
                .into_iter()
                .filter_map(|child| match detail::feature(doc, child) {
                    Some(FeatureKind::ForwardTarget(target)) => {
                        Some(format!("{}(w={})", target.target_group, target.weight))
                    }
                    _ => None,
                })
                .collect();
            if !weighted.is_empty() {
                value.push_str(&format!(" → {}", weighted.join(" + ")));
            }
            let details: Vec<_> = detail::children(doc, index)
                .into_iter()
                .filter_map(|child| match detail::feature(doc, child) {
                    Some(FeatureKind::Option(option)) if presentation::has(mask, "options") => {
                        let name = super::option::DEFINITIONS[option.key].label;
                        Some(format!("{}={}", name, option.value))
                    }
                    Some(FeatureKind::JwtClaim(claim)) => {
                        Some(format!("{}={}", claim.name, claim.format))
                    }
                    _ => None,
                })
                .collect();
            if !details.is_empty() {
                value.push_str(&format!(" · {}", details.join(", ")));
            }
            Some(value)
        }
        FeatureKind::Transform(model) if presentation::has(mask, "transforms") => {
            let rewrites = detail::children(doc, index)
                .into_iter()
                .filter_map(|child| match detail::feature(doc, child) {
                    Some(FeatureKind::Rewrite(rewrite)) => {
                        Some(format!("{} → {}", rewrite.regex, rewrite.replacement))
                    }
                    _ => None,
                })
                .collect::<Vec<_>>()
                .join(", ");
            Some(if rewrites.is_empty() {
                transform::label(&model.kind).into()
            } else {
                format!("{}: {}", transform::label(&model.kind), rewrites)
            })
        }
        _ => None,
    }
}

pub(super) fn measure(doc: &DocumentSpec, listener: usize, mask: u32) -> Table {
    if !presentation::has(mask, "rules") {
        return Table::empty();
    }
    let mut values = Vec::new();
    let mut rule_indexes = detail::children(doc, listener);
    rule_indexes.sort_by_key(|index| match detail::feature(doc, *index) {
        Some(FeatureKind::Rule(rule)) => rule.priority.unwrap_or(u16::MAX),
        _ => u16::MAX,
    });
    for index in rule_indexes {
        let Some(FeatureKind::Rule(rule)) = detail::feature(doc, index) else {
            continue;
        };
        if doc.elements[index].visual.visible == Some(false) {
            continue;
        }
        let child_mask = match doc.elements[index].aws.as_ref() {
            Some(crate::ent::model::aws::Component::Feature(feature)) => {
                presentation::resolve(&feature.presentation, mask)
            }
            _ => mask,
        };
        let conditions = detail::children(doc, index)
            .into_iter()
            .filter_map(|child| condition_text(doc, child, child_mask))
            .collect::<Vec<_>>()
            .join(" AND ");
        let conditions = if conditions.is_empty() {
            "Otherwise".into()
        } else {
            conditions
        };
        let actions = detail::children(doc, index)
            .into_iter()
            .filter_map(|child| action_text(doc, child, child_mask))
            .collect::<Vec<_>>()
            .join("; ");
        values.push((
            index,
            doc.elements[index].id.clone(),
            rule.priority.map_or("Default".into(), |p| format!("#{p}")),
            conditions,
            actions,
        ));
    }
    if values.is_empty() {
        return Table::empty();
    }
    let priority_width = 72.0;
    let condition_width = values
        .iter()
        .map(|(_, _, _, v, _)| text_width(v))
        .fold(text_width("Conditions"), f64::max)
        .min(280.0)
        .ceil()
        + 16.0;
    let action_width = values
        .iter()
        .map(|(_, _, _, _, v)| text_width(v))
        .fold(text_width("Action"), f64::max)
        .min(220.0)
        .ceil()
        + 16.0;
    let rows = values
        .into_iter()
        .map(|(index, id, priority, conditions, action)| {
            let lines =
                line_count(&conditions, condition_width).max(line_count(&action, action_width));
            RuleRow {
                index,
                id,
                priority,
                conditions,
                action,
                height: 14.0 + lines as f64 * 18.0,
            }
        })
        .collect::<Vec<_>>();
    Table {
        width: priority_width + condition_width + action_width,
        height: 54.0 + rows.iter().map(|r| r.height).sum::<f64>(),
        priority_width,
        condition_width,
        action_width,
        rows,
    }
}

impl Table {
    fn empty() -> Self {
        Self {
            width: 0.0,
            height: 0.0,
            priority_width: 0.0,
            condition_width: 0.0,
            action_width: 0.0,
            rows: Vec::new(),
        }
    }
}

fn cell(
    doc: &mut DocumentSpec,
    owner: &ElementSpec,
    parent: usize,
    key: &str,
    bounds: [f64; 4],
    text: &str,
    fill: &str,
) {
    let mut cell = drawing::part(owner, parent, key, bounds);
    cell.visual.shape = Shape::Rectangle;
    cell.visual.fill = fill.into();
    cell.visual.stroke = "#c7d8f4".into();
    cell.text.value = text.into();
    cell.text.font_size = Some(12.0);
    cell.text.padding.left = Some(6.0);
    cell.text.padding.right = Some(6.0);
    doc.elements.push(cell);
}

pub(super) fn draw(
    doc: &mut DocumentSpec,
    table: &Table,
    parent: usize,
    x: f64,
    y: f64,
    aliases: &mut Vec<(String, String)>,
) {
    if table.rows.is_empty() {
        return;
    }
    let owner = doc.elements[parent].clone();
    cell(
        doc,
        &owner,
        parent,
        "rules-title",
        [x, y, table.width, 26.0],
        "Listener rules",
        "#ede9fe",
    );
    let top = y + 26.0;
    cell(
        doc,
        &owner,
        parent,
        "rules-priority",
        [x, top, table.priority_width, 28.0],
        "Priority",
        "#e8efff",
    );
    cell(
        doc,
        &owner,
        parent,
        "rules-conditions",
        [x + table.priority_width, top, table.condition_width, 28.0],
        "Conditions",
        "#e8efff",
    );
    cell(
        doc,
        &owner,
        parent,
        "rules-action",
        [
            x + table.priority_width + table.condition_width,
            top,
            table.action_width,
            28.0,
        ],
        "Action",
        "#e8efff",
    );
    let mut top = top + 28.0;
    for row in &table.rows {
        cell(
            doc,
            &owner,
            parent,
            &format!("rule-{}-priority", row.id),
            [x, top, table.priority_width, row.height],
            &row.priority,
            "#ffffff",
        );
        cell(
            doc,
            &owner,
            parent,
            &format!("rule-{}-conditions", row.id),
            [
                x + table.priority_width,
                top,
                table.condition_width,
                row.height,
            ],
            &row.conditions,
            "#ffffff",
        );
        cell(
            doc,
            &owner,
            parent,
            &format!("rule-{}-action", row.id),
            [
                x + table.priority_width + table.condition_width,
                top,
                table.action_width,
                row.height,
            ],
            &row.action,
            "#ffffff",
        );
        detail::collapse(doc, row.index, &owner.id, aliases);
        top += row.height;
    }
}
