#[rustfmt::skip]
use super::{
    drawing::{
        badge,
        part,
    },
    listener,
};
#[rustfmt::skip]
use crate::ent::model::{
    aws::{
        Component,
        listener::Listener,
    },
    document::*,
};
use crate::usc::layout::presentation_text_width;
use crate::util::error::LayoutError;

pub(super) fn compose(
    document: &mut DocumentSpec,
    index: usize,
    domain: &str,
    kind: &str,
    validate: fn(&Listener) -> Result<(), LayoutError>,
) -> Result<(), LayoutError> {
    let mut listeners = Vec::new();
    let mut ports = std::collections::HashSet::new();
    for (child_index, element) in document.elements.iter().enumerate() {
        if element.parent != Some(index) {
            continue;
        }
        let Some(Component::Listener(model)) = &element.aws else {
            return Err(LayoutError::new(
                "AWS load balancers accept only aws-listener children",
            ));
        };
        validate(model)?;
        if !ports.insert(model.port) {
            return Err(LayoutError::new(
                "AWS load balancer listener ports must be unique",
            ));
        }
        if element.weight.is_some()
            || element.width.is_some()
            || element.height.is_some()
            || element.x.is_some()
            || element.y.is_some()
        {
            return Err(LayoutError::new(
                "AWS listener geometry is controlled by its load balancer",
            ));
        }
        listeners.push((child_index, model.clone()));
    }
    if listeners.is_empty() || listeners.len() > 32 {
        return Err(LayoutError::new(
            "AWS load balancer components require 1..32 listeners",
        ));
    }
    if domain.len() > 1024 || domain.chars().any(char::is_control) {
        return Err(LayoutError::new(
            "AWS domain label must be single-line and at most 1024 bytes",
        ));
    }
    let label = if domain.is_empty() { kind } else { domain };
    let tag_width = presentation_text_width(label, 14.0).ceil() + 20.0;
    let card_width = listeners
        .iter()
        .map(|(_, model)| listener::measure(model).0)
        .fold(0.0_f64, f64::max);
    let owner = &mut document.elements[index];
    if owner.weight.is_some() {
        return Err(LayoutError::new(
            "AWS load balancer components use width/height, not weight",
        ));
    }
    let width = owner.width.unwrap_or(
        (12.0 + listeners.len().min(3) as f64 * (card_width + 12.0)).max(tag_width + 64.0),
    );
    if width < (card_width + 24.0).max(tag_width + 64.0) {
        return Err(LayoutError::new(
            "AWS load balancer width cannot contain its domain tag and listener",
        ));
    }
    let columns = (((width - 12.0) / (card_width + 12.0)).floor() as usize)
        .max(1)
        .min(listeners.len());
    let rows = listeners.len().div_ceil(columns);
    let row_heights: Vec<f64> = listeners
        .chunks(columns)
        .map(|row| {
            row.iter()
                .map(|(_, model)| listener::measure(model).1)
                .fold(0.0_f64, f64::max)
        })
        .collect();
    let minimum_height = 56.0 + row_heights.iter().sum::<f64>() + rows as f64 * 12.0;
    let height = owner.height.unwrap_or(minimum_height);
    if height < minimum_height {
        return Err(LayoutError::new(format!(
            "AWS load balancer height must be at least {minimum_height}"
        )));
    }
    let icon = owner.icon.clone();
    owner.icon.reference.clear();
    owner.text.value.clear();
    owner.concept = Concept::Group;
    owner.layout = LayoutPolicy::Absolute;
    owner.padding = Insets::default();
    owner.width = Some(width);
    owner.height = Some(height);
    owner.visual.shape = Shape::Rectangle;
    if owner.visual.fill.is_empty() {
        owner.visual.fill = "#f8fbff".into();
    }
    if owner.visual.stroke.is_empty() {
        owner.visual.stroke = "#b9cdf0".into();
    }
    owner.visual.stroke_width.get_or_insert(1.0);
    owner.visual.corner_radius.get_or_insert(12.0);
    let owner = owner.clone();
    let mut symbol = part(&owner, index, "icon", [12.0, 12.0, 32.0, 32.0]);
    symbol.concept = Concept::Item;
    symbol.icon = icon;
    symbol.icon.width = Some(32.0);
    symbol.icon.height = Some(32.0);
    let mut domain_tag = badge(
        &owner,
        index,
        "domain",
        [52.0, 14.0, tag_width, 28.0],
        label,
        false,
    );
    domain_tag.visual.fill = "#e8efff".into();
    domain_tag.text.color = "#172554".into();
    domain_tag.text.font_size = Some(14.0);
    document.elements.extend([symbol, domain_tag]);
    for (ordinal, (child_index, model)) in listeners.iter().enumerate() {
        if owner.visual.visible == Some(false) {
            document.elements[*child_index].visual.visible = Some(false);
        }
        document.elements[*child_index].visual.layer = owner.visual.layer;
        listener::compose(
            document,
            *child_index,
            model,
            12.0 + (ordinal % columns) as f64 * (card_width + 12.0),
            56.0 + row_heights[..ordinal / columns].iter().sum::<f64>()
                + (ordinal / columns) as f64 * 12.0,
        );
    }
    Ok(())
}
