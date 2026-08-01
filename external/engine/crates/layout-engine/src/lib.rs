use std::collections::HashSet;
use std::error::Error;
use std::fmt::{Display, Formatter};

const MAX_ELEMENTS: usize = 10_000;
const MAX_ID_BYTES: usize = u16::MAX as usize;

#[derive(Clone, Copy, Debug, Eq, PartialEq)]
pub enum Direction {
    Vertical,
    Horizontal,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ElementSpec {
    pub id: String,
    pub width: Option<f64>,
    pub height: Option<f64>,
    pub weight: Option<f64>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct DocumentSpec {
    pub direction: Direction,
    pub width: f64,
    pub height: f64,
    pub gap: f64,
    pub elements: Vec<ElementSpec>,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ResolvedElement {
    pub id: String,
    pub x: f64,
    pub y: f64,
    pub width: f64,
    pub height: f64,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ResolvedDocument {
    pub width: f64,
    pub height: f64,
    pub elements: Vec<ResolvedElement>,
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct LayoutError {
    message: String,
}

impl LayoutError {
    fn new(message: impl Into<String>) -> Self {
        Self {
            message: message.into(),
        }
    }
}

impl Display for LayoutError {
    fn fmt(&self, formatter: &mut Formatter<'_>) -> std::fmt::Result {
        formatter.write_str(&self.message)
    }
}

impl Error for LayoutError {}

pub fn resolve(document: &DocumentSpec) -> Result<ResolvedDocument, LayoutError> {
    validate_positive_finite("document width", document.width)?;
    validate_positive_finite("document height", document.height)?;
    validate_non_negative_finite("document gap", document.gap)?;
    if document.elements.is_empty() {
        return Err(LayoutError::new("document requires at least one element"));
    }
    if document.elements.len() > MAX_ELEMENTS {
        return Err(LayoutError::new(format!(
            "document element count {} exceeds {MAX_ELEMENTS}",
            document.elements.len()
        )));
    }

    let (main_limit, cross_limit) = match document.direction {
        Direction::Vertical => (document.height, document.width),
        Direction::Horizontal => (document.width, document.height),
    };
    let mut identifiers = HashSet::with_capacity(document.elements.len());
    let mut fixed_total = 0.0;
    let mut weight_total = 0.0;

    for element in &document.elements {
        validate_element(element, document.direction, cross_limit)?;
        if !identifiers.insert(element.id.as_str()) {
            return Err(LayoutError::new(format!(
                "duplicate element id {:?}",
                element.id
            )));
        }
        if let Some(weight) = element.weight {
            weight_total += weight;
        } else {
            fixed_total += main_size(element, document.direction).expect("validated main size");
        }
    }

    let gap_total = document.gap * (document.elements.len().saturating_sub(1) as f64);
    let remaining = main_limit - fixed_total - gap_total;
    if remaining < 0.0 {
        return Err(LayoutError::new(format!(
            "fixed elements and gaps exceed available main-axis size by {}",
            format_number(-remaining)
        )));
    }
    if weight_total > 0.0 && remaining <= 0.0 {
        return Err(LayoutError::new(
            "flex elements require positive remaining main-axis size",
        ));
    }

    let mut cursor = 0.0;
    let mut resolved = Vec::with_capacity(document.elements.len());
    for element in &document.elements {
        let allocated_main = match element.weight {
            Some(weight) => remaining * weight / weight_total,
            None => main_size(element, document.direction).expect("validated main size"),
        };
        let allocated_cross = cross_size(element, document.direction).unwrap_or(cross_limit);
        let (x, y, width, height) = match document.direction {
            Direction::Vertical => (0.0, cursor, allocated_cross, allocated_main),
            Direction::Horizontal => (cursor, 0.0, allocated_main, allocated_cross),
        };
        resolved.push(ResolvedElement {
            id: element.id.clone(),
            x,
            y,
            width,
            height,
        });
        cursor += allocated_main + document.gap;
    }

    Ok(ResolvedDocument {
        width: document.width,
        height: document.height,
        elements: resolved,
    })
}

fn validate_element(
    element: &ElementSpec,
    direction: Direction,
    cross_limit: f64,
) -> Result<(), LayoutError> {
    if element.id.trim().is_empty() {
        return Err(LayoutError::new("element id must not be empty"));
    }
    if element.id.len() > MAX_ID_BYTES {
        return Err(LayoutError::new(format!(
            "element id exceeds {MAX_ID_BYTES} UTF-8 bytes"
        )));
    }
    if let Some(width) = element.width {
        validate_positive_finite(&format!("element {:?} width", element.id), width)?;
    }
    if let Some(height) = element.height {
        validate_positive_finite(&format!("element {:?} height", element.id), height)?;
    }
    if let Some(weight) = element.weight {
        validate_positive_finite(&format!("element {:?} weight", element.id), weight)?;
    }

    let main = main_size(element, direction);
    if element.weight.is_some() && main.is_some() {
        return Err(LayoutError::new(format!(
            "element {:?} cannot set both main-axis size and weight",
            element.id
        )));
    }
    if element.weight.is_none() && main.is_none() {
        return Err(LayoutError::new(format!(
            "element {:?} requires a main-axis size or weight",
            element.id
        )));
    }
    if let Some(cross) = cross_size(element, direction) {
        if cross > cross_limit {
            return Err(LayoutError::new(format!(
                "element {:?} cross-axis size {} exceeds container size {}",
                element.id,
                format_number(cross),
                format_number(cross_limit)
            )));
        }
    }
    Ok(())
}

fn main_size(element: &ElementSpec, direction: Direction) -> Option<f64> {
    match direction {
        Direction::Vertical => element.height,
        Direction::Horizontal => element.width,
    }
}

fn cross_size(element: &ElementSpec, direction: Direction) -> Option<f64> {
    match direction {
        Direction::Vertical => element.width,
        Direction::Horizontal => element.height,
    }
}

fn validate_positive_finite(name: &str, value: f64) -> Result<(), LayoutError> {
    if !value.is_finite() || value <= 0.0 {
        return Err(LayoutError::new(format!(
            "{name} must be finite and positive"
        )));
    }
    Ok(())
}

fn validate_non_negative_finite(name: &str, value: f64) -> Result<(), LayoutError> {
    if !value.is_finite() || value < 0.0 {
        return Err(LayoutError::new(format!(
            "{name} must be finite and non-negative"
        )));
    }
    Ok(())
}

fn format_number(value: f64) -> String {
    let rendered = format!("{value:.6}");
    rendered
        .trim_end_matches('0')
        .trim_end_matches('.')
        .to_owned()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn fixed(id: &str, width: Option<f64>, height: Option<f64>) -> ElementSpec {
        ElementSpec {
            id: id.to_owned(),
            width,
            height,
            weight: None,
        }
    }

    #[test]
    fn allocates_fixed_before_flexible_vertical_children() {
        let document = DocumentSpec {
            direction: Direction::Vertical,
            width: 200.0,
            height: 300.0,
            gap: 10.0,
            elements: vec![
                fixed("header", None, Some(40.0)),
                ElementSpec {
                    id: "body".to_owned(),
                    width: Some(160.0),
                    height: None,
                    weight: Some(1.0),
                },
                ElementSpec {
                    id: "footer".to_owned(),
                    width: None,
                    height: None,
                    weight: Some(2.0),
                },
            ],
        };

        let resolved = resolve(&document).expect("resolve document");
        assert_eq!(resolved.elements[0].height, 40.0);
        assert_eq!(resolved.elements[1].y, 50.0);
        assert_eq!(resolved.elements[1].height, 80.0);
        assert_eq!(resolved.elements[1].width, 160.0);
        assert_eq!(resolved.elements[2].y, 140.0);
        assert_eq!(resolved.elements[2].height, 160.0);
        assert_eq!(resolved.elements[2].width, 200.0);
    }

    #[test]
    fn rejects_conflicting_or_non_finite_values() {
        let conflicting = DocumentSpec {
            direction: Direction::Horizontal,
            width: 100.0,
            height: 40.0,
            gap: 0.0,
            elements: vec![ElementSpec {
                id: "bad".to_owned(),
                width: Some(20.0),
                height: None,
                weight: Some(1.0),
            }],
        };
        assert!(resolve(&conflicting)
            .expect_err("conflicting size and weight")
            .to_string()
            .contains("both main-axis size and weight"));

        let invalid = DocumentSpec {
            direction: Direction::Vertical,
            width: f64::NAN,
            height: 100.0,
            gap: 0.0,
            elements: vec![fixed("item", None, Some(20.0))],
        };
        assert!(resolve(&invalid)
            .expect_err("non-finite width")
            .to_string()
            .contains("document width"));
    }

    #[test]
    fn output_is_deterministic() {
        let document = DocumentSpec {
            direction: Direction::Horizontal,
            width: 240.0,
            height: 80.0,
            gap: 8.0,
            elements: vec![
                fixed("left", Some(40.0), Some(60.0)),
                ElementSpec {
                    id: "right".to_owned(),
                    width: None,
                    height: None,
                    weight: Some(1.0),
                },
            ],
        };
        assert_eq!(resolve(&document), resolve(&document));
    }
}
