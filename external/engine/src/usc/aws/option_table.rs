use super::{
    detail::{children, feature},
    drawing, option, presentation,
};
use crate::ent::model::{aws::feature::FeatureKind, document::*};

// A setting remains a source-addressable row. Generic shapes/text cells are
// shared by SVG and PPTX; no renderer-specific table implementation is needed.
pub(super) struct TableRow {
    index: usize,
    keys: Vec<String>,
    values: Vec<String>,
    height: f64,
    active: bool,
}
pub(super) struct Table {
    pub width: f64,
    pub height: f64,
    key_width: f64,
    value_width: f64,
    rows: Vec<TableRow>,
}
fn width(value: &str) -> f64 {
    crate::usc::layout::presentation_text_width(value, 12.0)
}
fn wrap(value: &str, limit: f64) -> Vec<String> {
    let mut lines = Vec::new();
    let mut line = String::new();
    for c in value.chars() {
        line.push(c);
        if width(&line) > limit {
            line.pop();
            lines.push(line);
            line = c.to_string();
        }
    }
    if !line.is_empty() {
        lines.push(line);
    }
    lines
}
pub(super) fn measure(doc: &DocumentSpec, index: usize, inherited: u32) -> Table {
    let mut cells = Vec::new();
    if presentation::has(inherited, "options") {
        for i in children(doc, index) {
            if let Some(FeatureKind::Option(o)) = feature(doc, i) {
                if doc.elements[i].visual.visible == Some(false) {
                    continue;
                }
                let mask = match &doc.elements[i].aws {
                    Some(crate::ent::model::aws::Component::Feature(f)) => {
                        presentation::resolve(&f.presentation, inherited)
                    }
                    _ => inherited,
                };
                let key = if o.name.is_empty() {
                    option::DEFINITIONS[o.key].label.into()
                } else {
                    format!("{} {}", option::DEFINITIONS[o.key].label, o.name)
                };
                let value = if presentation::has(mask, "values") {
                    match o.value.as_str() {
                        "true" => "ON",
                        "false" => "OFF",
                        v => v,
                    }
                    .to_string()
                } else {
                    String::new()
                };
                cells.push((
                    i,
                    key,
                    value,
                    matches!(o.value.as_str(), "true" | "on" | "healthy"),
                ));
            }
        }
    }
    if cells.is_empty() {
        return Table {
            width: 0.0,
            height: 0.0,
            key_width: 0.0,
            value_width: 0.0,
            rows: Vec::new(),
        };
    }
    let key_width = cells
        .iter()
        .map(|(_, k, _, _)| width(k))
        .fold(width("Setting"), f64::max)
        .min(168.0)
        .ceil()
        + 16.0;
    let value_width = if cells.iter().all(|(_, _, v, _)| v.is_empty()) {
        0.0
    } else {
        cells
            .iter()
            .map(|(_, _, v, _)| width(v))
            .fold(width("Value"), f64::max)
            .min(240.0)
            .ceil()
            + 16.0
    };
    let rows: Vec<_> = cells
        .into_iter()
        .map(|(index, key, value, active)| {
            let keys = wrap(&key, key_width - 16.0);
            let values = if value_width > 0.0 {
                wrap(&value, value_width - 16.0)
            } else {
                Vec::new()
            };
            let height = 12.0 + 18.0 * keys.len().max(values.len()).max(1) as f64;
            TableRow {
                index,
                keys,
                values,
                height,
                active,
            }
        })
        .collect();
    Table {
        width: key_width + value_width,
        height: 28.0 + rows.iter().map(|r| r.height).sum::<f64>(),
        key_width,
        value_width,
        rows,
    }
}
fn cell(
    doc: &mut DocumentSpec,
    owner: &ElementSpec,
    parent: usize,
    key: &str,
    bounds: [f64; 4],
    lines: &[String],
    fill: &str,
    active: bool,
) {
    let mut bg = drawing::part(owner, parent, key, bounds);
    bg.visual.shape = Shape::Rectangle;
    bg.visual.fill = fill.into();
    bg.visual.stroke = "#c7d8f4".into();
    doc.elements.push(bg);
    for (i, line) in lines.iter().enumerate() {
        let mut text = drawing::part(
            owner,
            parent,
            &format!("{key}-{i}"),
            [
                bounds[0] + 8.0,
                bounds[1] + 6.0 + i as f64 * 18.0,
                bounds[2] - 16.0,
                18.0,
            ],
        );
        text.text.value = line.clone();
        text.text.font_size = Some(12.0);
        if active {
            text.text.color = "#166534".into();
        }
        doc.elements.push(text);
    }
}
pub(super) fn draw(doc: &mut DocumentSpec, table: &Table, parent: usize, x: f64, y: f64) {
    if table.rows.is_empty() {
        return;
    }
    let owner = doc.elements[parent].clone();
    cell(
        doc,
        &owner,
        parent,
        "settings-header-key",
        [x, y, table.key_width, 28.0],
        &["Setting".into()],
        "#e8efff",
        false,
    );
    if table.value_width > 0.0 {
        cell(
            doc,
            &owner,
            parent,
            "settings-header-value",
            [x + table.key_width, y, table.value_width, 28.0],
            &["Value".into()],
            "#e8efff",
            false,
        );
    }
    let mut top = y + 28.0;
    for row in &table.rows {
        let e = &mut doc.elements[row.index];
        e.concept = Concept::Group;
        e.layout = LayoutPolicy::Absolute;
        e.padding = Insets::default();
        e.margin = Insets::default();
        e.x = Some(x);
        e.y = Some(top);
        e.width = Some(table.width);
        e.height = Some(row.height);
        e.visual.shape = Shape::None;
        e.visual.fill = "none".into();
        e.visual.stroke = "none".into();
        e.visual.visible = owner.visual.visible;
        e.visual.layer = owner.visual.layer;
        e.text.value.clear();
        e.icon.reference.clear();
        let owner = e.clone();
        cell(
            doc,
            &owner,
            row.index,
            "key",
            [0.0, 0.0, table.key_width, row.height],
            &row.keys,
            "#f8fafc",
            false,
        );
        if table.value_width > 0.0 {
            cell(
                doc,
                &owner,
                row.index,
                "value",
                [table.key_width, 0.0, table.value_width, row.height],
                &row.values,
                if row.active && !row.values.is_empty() {
                    "#dcfce7"
                } else {
                    "#ffffff"
                },
                row.active,
            );
        }
        top += row.height;
    }
}
