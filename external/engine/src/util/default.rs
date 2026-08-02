use crate::ent::model::document::{
    Insets,
    Point,
};

impl Insets {
    const fn empty() -> Self {
        Self {
            top: None,
            right: None,
            bottom: None,
            left: None,
        }
    }
}

impl Default for Insets {
    fn default() -> Self {
        Self::empty()
    }
}

impl Point {
    const fn origin() -> Self {
        Self { x: 0.0, y: 0.0 }
    }
}

impl Default for Point {
    fn default() -> Self {
        Self::origin()
    }
}
