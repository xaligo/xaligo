#[rustfmt::skip]
use std::collections::{
    HashMap,
    HashSet,
};

#[rustfmt::skip]
use crate::cnf::engine::{
    DEFAULT_FONT_SIZE,
    DEFAULT_GAP,
    DEFAULT_ITEM_SIZE,
    DEFAULT_LINE_HEIGHT,
    DEFAULT_PORT_SIZE,
    MAX_COLUMNS,
    MAX_DEPTH,
    MAX_ELEMENTS,
    MAX_ID_BYTES,
};
use crate::ent::model::document::{
    Alignment,
    Concept,
    DocumentSpec,
    ElementSpec,
    Insets,
    Justification,
    LayoutPolicy,
    MissingIconPolicy,
    Overflow,
    Point,
    ResolvedDocument,
    ResolvedElement,
    ResolvedLine,
    ResolvedText,
    ResolvedVisual,
    RoutingPolicy,
    Shape,
    Side,
};
use crate::util::error::LayoutError;

impl Insets {
    fn resolved(self) -> ResolvedInsets {
        ResolvedInsets {
            top: self.top.unwrap_or(0.0),
            right: self.right.unwrap_or(0.0),
            bottom: self.bottom.unwrap_or(0.0),
            left: self.left.unwrap_or(0.0),
        }
    }
}

struct ResolvedInsets {
    top: f64,
    right: f64,
    bottom: f64,
    left: f64,
}

impl Clone for ResolvedInsets {
    fn clone(&self) -> Self {
        *self
    }
}

impl Copy for ResolvedInsets {}

struct Bounds {
    x: f64,
    y: f64,
    width: f64,
    height: f64,
}

impl Clone for Bounds {
    fn clone(&self) -> Self {
        *self
    }
}

impl Copy for Bounds {}

impl Bounds {
    fn content(self, padding: ResolvedInsets, owner: &str) -> Result<Self, LayoutError> {
        let width = self.width - padding.left - padding.right;
        let height = self.height - padding.top - padding.bottom;
        if width <= 0.0 || height <= 0.0 {
            return Err(LayoutError::new(format!(
                "element {owner:?} padding consumes its content box"
            )));
        }
        Ok(Self {
            x: self.x + padding.left,
            y: self.y + padding.top,
            width,
            height,
        })
    }

    fn center(self) -> Point {
        Point {
            x: self.x + self.width / 2.0,
            y: self.y + self.height / 2.0,
        }
    }

    fn expanded(self, amount: f64) -> Self {
        Self {
            x: self.x - amount,
            y: self.y - amount,
            width: self.width + amount * 2.0,
            height: self.height + amount * 2.0,
        }
    }
}

pub(crate) fn resolve(document: &DocumentSpec) -> Result<ResolvedDocument, LayoutError> {
    crate::usc::cancel::check().map_err(LayoutError::new)?;
    validate_document(document)?;
    let mut children = vec![Vec::new(); document.elements.len()];
    let mut roots = Vec::new();
    for (index, element) in document.elements.iter().enumerate() {
        if let Some(parent) = element.parent {
            children[parent].push(index);
        } else {
            roots.push(index);
        }
    }

    let document_bounds = Bounds {
        x: 0.0,
        y: 0.0,
        width: document.width,
        height: document.height,
    };
    let content = document_bounds.content(document.padding.resolved(), "document")?;
    let mut state = LayoutState {
        document,
        children,
        resolved: vec![None; document.elements.len()],
    };
    state.layout_children(
        &roots,
        content,
        effective_layout(document.layout),
        document.gap,
        document.overflow,
        document.columns.unwrap_or(12),
        Alignment::Stretch,
        Justification::Start,
    )?;
    state.align_v1_profile_group_headers()?;
    state.layout_lines()?;

    let elements = state
        .resolved
        .into_iter()
        .enumerate()
        .map(|(index, value)| {
            value.ok_or_else(|| {
                LayoutError::new(format!(
                    "element {:?} was not resolved",
                    document.elements[index].id
                ))
            })
        })
        .collect::<Result<Vec<_>, _>>()?;
    Ok(ResolvedDocument {
        width: document.width,
        height: document.height,
        elements,
    })
}

struct LayoutState<'a> {
    document: &'a DocumentSpec,
    children: Vec<Vec<usize>>,
    resolved: Vec<Option<ResolvedElement>>,
}


include!("layout_flow.rs");
include!("layout_geometry.rs");
include!("layout_routing.rs");
include!("layout_validation.rs");
