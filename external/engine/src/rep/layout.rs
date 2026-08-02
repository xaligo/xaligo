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

pub fn resolve(document: &DocumentSpec) -> Result<ResolvedDocument, LayoutError> {
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

impl LayoutState<'_> {
    #[allow(clippy::too_many_arguments)]
    fn layout_children(
        &mut self,
        indices: &[usize],
        content: Bounds,
        policy: LayoutPolicy,
        gap: f64,
        overflow: Overflow,
        columns: u16,
        align: Alignment,
        justify: Justification,
    ) -> Result<(), LayoutError> {
        let flow = indices
            .iter()
            .copied()
            .filter(|index| {
                !matches!(
                    self.document.elements[*index].concept,
                    Concept::Port | Concept::Line
                )
            })
            .collect::<Vec<_>>();
        match policy {
            LayoutPolicy::Vertical => self.layout_stack(
                &flow,
                content,
                Axis::Vertical,
                gap,
                overflow,
                align,
                justify,
            )?,
            LayoutPolicy::Horizontal => self.layout_stack(
                &flow,
                content,
                Axis::Horizontal,
                gap,
                overflow,
                align,
                justify,
            )?,
            LayoutPolicy::Grid => {
                self.layout_grid(&flow, content, gap, overflow, columns, align)?
            }
            LayoutPolicy::Absolute | LayoutPolicy::None | LayoutPolicy::Default => {
                self.layout_absolute(&flow, content, overflow, align)?
            }
        }

        for index in indices
            .iter()
            .copied()
            .filter(|index| self.document.elements[*index].concept == Concept::Port)
        {
            self.layout_port(index)?;
        }
        if overflow == Overflow::Error {
            for index in indices.iter().copied().filter(|index| {
                !matches!(
                    self.document.elements[*index].concept,
                    Concept::Port | Concept::Line
                )
            }) {
                let child = self.resolved[index]
                    .as_ref()
                    .ok_or_else(|| LayoutError::new("child geometry was not resolved"))?;
                if !contains(content, bounds_of(child)) {
                    return Err(LayoutError::new(format!(
                        "element {:?} exceeds its parent content box",
                        child.id
                    )));
                }
            }
        }
        Ok(())
    }

    #[allow(clippy::too_many_arguments)]
    fn layout_stack(
        &mut self,
        indices: &[usize],
        content: Bounds,
        axis: Axis,
        configured_gap: f64,
        overflow: Overflow,
        align: Alignment,
        justify: Justification,
    ) -> Result<(), LayoutError> {
        if indices.is_empty() {
            return Ok(());
        }
        let main_limit = axis.main_size(content);
        let cross_limit = axis.cross_size(content);
        let mut fixed_total = 0.0;
        let mut weight_total = 0.0;
        let mut allocations = Vec::with_capacity(indices.len());
        for index in indices {
            let element = &self.document.elements[*index];
            let margin = element.margin.resolved();
            let main_margin = axis.main_before(margin) + axis.main_after(margin);
            let fixed = axis.explicit_size(element).or_else(|| {
                if element.weight.is_none() {
                    intrinsic_main(element, axis)
                } else {
                    None
                }
            });
            if element.weight.is_some() && axis.explicit_size(element).is_some() {
                return Err(LayoutError::new(format!(
                    "element {:?} cannot set both main-axis size and weight",
                    element.id
                )));
            }
            if let Some(size) = fixed {
                fixed_total += size + main_margin;
                allocations.push(MainAllocation::Fixed(size));
            } else {
                let weight = element.weight.unwrap_or(1.0);
                weight_total += weight;
                fixed_total += main_margin;
                allocations.push(MainAllocation::Flexible(weight));
            }
        }

        let base_gap_total = configured_gap * indices.len().saturating_sub(1) as f64;
        let remaining = main_limit - fixed_total - base_gap_total;
        if remaining < 0.0 && overflow == Overflow::Error {
            return Err(LayoutError::new(format!(
                "children and gaps exceed available main-axis size by {}",
                format_number(-remaining)
            )));
        }
        let flex_pool = if weight_total == 0.0 {
            0.0
        } else if remaining >= 0.0 {
            remaining
        } else {
            main_limit
        };
        let used_main = fixed_total + base_gap_total + flex_pool;
        let free = (main_limit - used_main).max(0.0);
        let mut gap = configured_gap;
        let mut cursor = axis.main_start(content);
        match justify {
            Justification::Center => cursor += free / 2.0,
            Justification::End => cursor += free,
            Justification::SpaceBetween if indices.len() > 1 && weight_total == 0.0 => {
                gap += free / (indices.len() - 1) as f64;
            }
            Justification::SpaceEvenly if weight_total == 0.0 => {
                gap += free / (indices.len() + 1) as f64;
                cursor += free / (indices.len() + 1) as f64;
            }
            _ => {}
        }

        for ((index, allocation), position) in indices.iter().zip(allocations).zip(0..indices.len())
        {
            let element = &self.document.elements[*index];
            let margin = element.margin.resolved();
            cursor += axis.main_before(margin);
            let main = match allocation {
                MainAllocation::Fixed(value) => value,
                MainAllocation::Flexible(weight) => flex_pool * weight / weight_total,
            };
            let cross_available =
                cross_limit - axis.cross_before(margin) - axis.cross_after(margin);
            let requested_cross = axis
                .cross_explicit_size(element)
                .or_else(|| intrinsic_cross(element, axis));
            let cross = requested_cross.unwrap_or(cross_available);
            if cross > cross_available && overflow == Overflow::Error {
                return Err(LayoutError::new(format!(
                    "element {:?} cross-axis size {} exceeds available size {}",
                    element.id,
                    format_number(cross),
                    format_number(cross_available)
                )));
            }
            let cross_free = (cross_available - cross).max(0.0);
            let cross_offset = match align {
                Alignment::Center => cross_free / 2.0,
                Alignment::End => cross_free,
                Alignment::Start | Alignment::Stretch => 0.0,
            };
            let bounds = axis.bounds(
                cursor,
                axis.cross_start(content) + axis.cross_before(margin) + cross_offset,
                main,
                cross,
            );
            self.place(*index, bounds)?;
            cursor += main + axis.main_after(margin);
            if position + 1 < indices.len() {
                cursor += gap;
            }
        }
        Ok(())
    }

    fn layout_grid(
        &mut self,
        indices: &[usize],
        content: Bounds,
        gap: f64,
        overflow: Overflow,
        columns: u16,
        align: Alignment,
    ) -> Result<(), LayoutError> {
        if indices.is_empty() {
            return Ok(());
        }
        validate_columns("grid columns", columns)?;
        let default_span = if indices.len() <= columns as usize {
            f64::from(columns) / indices.len() as f64
        } else {
            1.0
        };
        let mut placements = Vec::with_capacity(indices.len());
        let mut row = 0u16;
        let mut column = 0.0f64;
        let mut row_count = 1u16;
        for index in indices {
            let element = &self.document.elements[*index];
            let column_span = element.column_span.map(f64::from).unwrap_or(default_span);
            let row_span = element.row_span.unwrap_or(1);
            if column_span <= 0.0 || column_span > f64::from(columns) || row_span == 0 {
                return Err(LayoutError::new(format!(
                    "element {:?} has invalid grid span {}x{} for {} columns",
                    element.id, column_span, row_span, columns
                )));
            }
            if column + column_span > f64::from(columns) + f64::EPSILON {
                row += 1;
                column = 0.0;
            }
            placements.push((*index, row, column, row_span, column_span));
            row_count = row_count.max(row + row_span);
            column += column_span;
        }
        let cell_width =
            (content.width - gap * (columns.saturating_sub(1) as f64)) / f64::from(columns);
        let cell_height =
            (content.height - gap * (row_count.saturating_sub(1) as f64)) / f64::from(row_count);
        if cell_width <= 0.0 || cell_height <= 0.0 {
            return Err(LayoutError::new("grid gaps consume the content box"));
        }

        for (index, row, column, row_span, column_span) in placements {
            let element = &self.document.elements[index];
            let margin = element.margin.resolved();
            let cell = Bounds {
                x: content.x + column * (cell_width + gap),
                y: content.y + f64::from(row) * (cell_height + gap),
                width: cell_width * column_span + gap * (column_span - 1.0).max(0.0),
                height: cell_height * f64::from(row_span)
                    + gap * f64::from(row_span.saturating_sub(1)),
            };
            let available_width = cell.width - margin.left - margin.right;
            let available_height = cell.height - margin.top - margin.bottom;
            let width = element.width.unwrap_or(available_width);
            let height = element.height.unwrap_or(available_height);
            let free_x = (available_width - width).max(0.0);
            let free_y = (available_height - height).max(0.0);
            let factor = match align {
                Alignment::Center => 0.5,
                Alignment::End => 1.0,
                Alignment::Start | Alignment::Stretch => 0.0,
            };
            let bounds = Bounds {
                x: cell.x + margin.left + free_x * factor,
                y: cell.y + margin.top + free_y * factor,
                width,
                height,
            };
            if (width > available_width || height > available_height) && overflow == Overflow::Error
            {
                return Err(LayoutError::new(format!(
                    "element {:?} exceeds its grid cell",
                    element.id
                )));
            }
            self.place(index, bounds)?;
        }
        Ok(())
    }

    fn layout_absolute(
        &mut self,
        indices: &[usize],
        content: Bounds,
        overflow: Overflow,
        align: Alignment,
    ) -> Result<(), LayoutError> {
        for index in indices {
            let element = &self.document.elements[*index];
            let margin = element.margin.resolved();
            let available_width = content.width - margin.left - margin.right;
            let available_height = content.height - margin.top - margin.bottom;
            let width = element
                .width
                .or(element.intrinsic_width)
                .unwrap_or_else(|| default_absolute_width(element, available_width));
            let height = element
                .height
                .or(element.intrinsic_height)
                .unwrap_or_else(|| default_absolute_height(element, available_height));
            let default_x = match align {
                Alignment::Center => (available_width - width) / 2.0,
                Alignment::End => available_width - width,
                Alignment::Start | Alignment::Stretch => 0.0,
            };
            let bounds = Bounds {
                x: content.x + margin.left + element.x.unwrap_or(default_x),
                y: content.y + margin.top + element.y.unwrap_or(0.0),
                width,
                height,
            };
            if overflow == Overflow::Error && !contains(content, bounds) {
                return Err(LayoutError::new(format!(
                    "element {:?} exceeds its parent content box",
                    element.id
                )));
            }
            self.place(*index, bounds)?;
        }
        Ok(())
    }

    fn place(&mut self, index: usize, mut bounds: Bounds) -> Result<(), LayoutError> {
        let element = &self.document.elements[index];
        bounds.x += element.offset_x.unwrap_or(0.0);
        bounds.y += element.offset_y.unwrap_or(0.0);
        validate_resolved_bounds(element, bounds)?;
        let resolved = resolved_element(element, bounds)?;
        self.resolved[index] = Some(resolved);

        let child_indices = self.children[index].clone();
        if child_indices.is_empty() {
            return Ok(());
        }
        let content = bounds.content(element.padding.resolved(), &element.id)?;
        let layout = effective_layout(element.layout);
        let gap = element.gap.unwrap_or(DEFAULT_GAP);
        let columns = element.columns.unwrap_or(12);
        self.layout_children(
            &child_indices,
            content,
            layout,
            gap,
            element.overflow,
            columns,
            element.align,
            element.justify,
        )?;
        Ok(())
    }

    fn layout_port(&mut self, index: usize) -> Result<(), LayoutError> {
        let element = &self.document.elements[index];
        let parent_index = element.parent.ok_or_else(|| {
            LayoutError::new(format!("port {:?} requires an owning element", element.id))
        })?;
        let owner = self.resolved[parent_index]
            .as_ref()
            .ok_or_else(|| LayoutError::new("port owner was not resolved"))?;
        let owner_bounds = Bounds {
            x: owner.x,
            y: owner.y,
            width: owner.width,
            height: owner.height,
        };
        let size = element.port.size.unwrap_or(DEFAULT_PORT_SIZE);
        let width = element.width.unwrap_or(size);
        let height = element.height.unwrap_or(size);
        let anchor = element.port.anchor.unwrap_or(0.5);
        let offset = element.port.offset.unwrap_or(0.0);
        let (x, y) = if element.x.is_some() || element.y.is_some() {
            (
                owner_bounds.x + element.x.unwrap_or(0.0),
                owner_bounds.y + element.y.unwrap_or(0.0),
            )
        } else {
            match element.port.side {
                Side::Top | Side::Auto => (
                    owner_bounds.x + anchor * (owner_bounds.width - width) + offset,
                    owner_bounds.y,
                ),
                Side::Right => (
                    owner_bounds.x + owner_bounds.width - width,
                    owner_bounds.y + anchor * (owner_bounds.height - height) + offset,
                ),
                Side::Bottom => (
                    owner_bounds.x + anchor * (owner_bounds.width - width) + offset,
                    owner_bounds.y + owner_bounds.height - height,
                ),
                Side::Left => (
                    owner_bounds.x,
                    owner_bounds.y + anchor * (owner_bounds.height - height) + offset,
                ),
            }
        };
        let bounds = Bounds {
            x: x + element.offset_x.unwrap_or(0.0),
            y: y + element.offset_y.unwrap_or(0.0),
            width,
            height,
        };
        if !contains(owner_bounds, bounds) {
            return Err(LayoutError::new(format!(
                "port {:?} lies outside owner {:?}",
                element.id, owner.id
            )));
        }
        let mut resolved = resolved_element(element, bounds)?;
        resolved.visual.visible = element.port.visible.unwrap_or(resolved.visual.visible);
        if resolved.text.value.is_empty() {
            resolved.text.value = element.port.label.clone();
        }
        self.resolved[index] = Some(resolved);
        Ok(())
    }

    fn layout_lines(&mut self) -> Result<(), LayoutError> {
        let id_index = self
            .document
            .elements
            .iter()
            .enumerate()
            .map(|(index, element)| (element.id.as_str(), index))
            .collect::<HashMap<_, _>>();
        for index in 0..self.document.elements.len() {
            let element = &self.document.elements[index];
            if element.concept != Concept::Line {
                continue;
            }
            let source_index = *id_index.get(element.line.source.as_str()).ok_or_else(|| {
                LayoutError::new(format!(
                    "line {:?} source {:?} does not exist",
                    element.id, element.line.source
                ))
            })?;
            let target_index = *id_index.get(element.line.target.as_str()).ok_or_else(|| {
                LayoutError::new(format!(
                    "line {:?} target {:?} does not exist",
                    element.id, element.line.target
                ))
            })?;
            let source = self.resolved[source_index]
                .as_ref()
                .ok_or_else(|| LayoutError::new("line source was not resolved"))?;
            let target = self.resolved[target_index]
                .as_ref()
                .ok_or_else(|| LayoutError::new("line target was not resolved"))?;
            let source_bounds = bounds_of(source);
            let target_bounds = bounds_of(target);
            let source_side = resolve_auto_side(
                element.line.source_side,
                source_bounds.center(),
                target_bounds.center(),
            );
            let target_side = resolve_auto_side(
                element.line.target_side,
                target_bounds.center(),
                source_bounds.center(),
            );
            let start = anchor_point(
                source_bounds,
                source_side,
                element.line.source_anchor.unwrap_or(0.5),
            );
            let end = anchor_point(
                target_bounds,
                target_side,
                element.line.target_anchor.unwrap_or(0.5),
            );
            let points = match element.line.routing {
                RoutingPolicy::Straight => vec![start, end],
                RoutingPolicy::Orthogonal => self.route_orthogonal(
                    start,
                    end,
                    source_side,
                    source_index,
                    target_index,
                    element.line.obstacle_margin.unwrap_or(8.0),
                ),
            };
            let min_x = points
                .iter()
                .map(|point| point.x)
                .fold(f64::INFINITY, f64::min);
            let max_x = points
                .iter()
                .map(|point| point.x)
                .fold(f64::NEG_INFINITY, f64::max);
            let min_y = points
                .iter()
                .map(|point| point.y)
                .fold(f64::INFINITY, f64::min);
            let max_y = points
                .iter()
                .map(|point| point.y)
                .fold(f64::NEG_INFINITY, f64::max);
            let bounds = Bounds {
                x: min_x,
                y: min_y,
                width: max_x - min_x,
                height: max_y - min_y,
            };
            let mut resolved = resolved_element(element, bounds)?;
            resolved.points = points;
            self.resolved[index] = Some(resolved);
        }
        Ok(())
    }

    fn route_orthogonal(
        &self,
        start: Point,
        end: Point,
        source_side: Side,
        source_index: usize,
        target_index: usize,
        margin: f64,
    ) -> Vec<Point> {
        let horizontal_first = matches!(source_side, Side::Left | Side::Right);
        let mut candidates = Vec::new();
        let mid_x = (start.x + end.x) / 2.0;
        let mid_y = (start.y + end.y) / 2.0;
        candidates.push(deduplicate_points(if horizontal_first {
            vec![
                start,
                Point {
                    x: mid_x,
                    y: start.y,
                },
                Point { x: mid_x, y: end.y },
                end,
            ]
        } else {
            vec![
                start,
                Point {
                    x: start.x,
                    y: mid_y,
                },
                Point { x: end.x, y: mid_y },
                end,
            ]
        }));
        candidates.push(deduplicate_points(vec![
            start,
            Point {
                x: mid_x,
                y: start.y,
            },
            Point { x: mid_x, y: end.y },
            end,
        ]));
        candidates.push(deduplicate_points(vec![
            start,
            Point {
                x: start.x,
                y: mid_y,
            },
            Point { x: end.x, y: mid_y },
            end,
        ]));

        let mut endpoint_tree = HashSet::from([source_index, target_index]);
        for mut current in [source_index, target_index] {
            while let Some(parent) = self.document.elements[current].parent {
                endpoint_tree.insert(parent);
                current = parent;
            }
        }
        let obstacle_bounds = self
            .resolved
            .iter()
            .enumerate()
            .filter_map(|(index, resolved)| {
                if endpoint_tree.contains(&index) {
                    return None;
                }
                let resolved = resolved.as_ref()?;
                if matches!(resolved.concept, Concept::Line | Concept::Spacer)
                    || !resolved.visual.visible
                {
                    return None;
                }
                Some(bounds_of(resolved).expanded(margin))
            })
            .collect::<Vec<_>>();
        for obstacle in &obstacle_bounds {
            candidates.push(deduplicate_points(vec![
                start,
                Point {
                    x: obstacle.x - margin,
                    y: start.y,
                },
                Point {
                    x: obstacle.x - margin,
                    y: end.y,
                },
                end,
            ]));
            candidates.push(deduplicate_points(vec![
                start,
                Point {
                    x: obstacle.x + obstacle.width + margin,
                    y: start.y,
                },
                Point {
                    x: obstacle.x + obstacle.width + margin,
                    y: end.y,
                },
                end,
            ]));
            candidates.push(deduplicate_points(vec![
                start,
                Point {
                    x: start.x,
                    y: obstacle.y - margin,
                },
                Point {
                    x: end.x,
                    y: obstacle.y - margin,
                },
                end,
            ]));
            candidates.push(deduplicate_points(vec![
                start,
                Point {
                    x: start.x,
                    y: obstacle.y + obstacle.height + margin,
                },
                Point {
                    x: end.x,
                    y: obstacle.y + obstacle.height + margin,
                },
                end,
            ]));
        }
        candidates
            .into_iter()
            .enumerate()
            .min_by(|(left_index, left), (right_index, right)| {
                route_score(left, &obstacle_bounds)
                    .partial_cmp(&route_score(right, &obstacle_bounds))
                    .unwrap_or(std::cmp::Ordering::Equal)
                    .then_with(|| left_index.cmp(right_index))
            })
            .map(|(_, points)| points)
            .unwrap_or_else(|| vec![start, end])
    }
}

enum Axis {
    Vertical,
    Horizontal,
}

impl Clone for Axis {
    fn clone(&self) -> Self {
        *self
    }
}

impl Copy for Axis {}

impl Axis {
    fn main_size(self, bounds: Bounds) -> f64 {
        match self {
            Self::Vertical => bounds.height,
            Self::Horizontal => bounds.width,
        }
    }

    fn cross_size(self, bounds: Bounds) -> f64 {
        match self {
            Self::Vertical => bounds.width,
            Self::Horizontal => bounds.height,
        }
    }

    fn main_start(self, bounds: Bounds) -> f64 {
        match self {
            Self::Vertical => bounds.y,
            Self::Horizontal => bounds.x,
        }
    }

    fn cross_start(self, bounds: Bounds) -> f64 {
        match self {
            Self::Vertical => bounds.x,
            Self::Horizontal => bounds.y,
        }
    }

    fn explicit_size(self, element: &ElementSpec) -> Option<f64> {
        match self {
            Self::Vertical => element.height,
            Self::Horizontal => element.width,
        }
    }

    fn cross_explicit_size(self, element: &ElementSpec) -> Option<f64> {
        match self {
            Self::Vertical => element.width,
            Self::Horizontal => element.height,
        }
    }

    fn main_before(self, margin: ResolvedInsets) -> f64 {
        match self {
            Self::Vertical => margin.top,
            Self::Horizontal => margin.left,
        }
    }

    fn main_after(self, margin: ResolvedInsets) -> f64 {
        match self {
            Self::Vertical => margin.bottom,
            Self::Horizontal => margin.right,
        }
    }

    fn cross_before(self, margin: ResolvedInsets) -> f64 {
        match self {
            Self::Vertical => margin.left,
            Self::Horizontal => margin.top,
        }
    }

    fn cross_after(self, margin: ResolvedInsets) -> f64 {
        match self {
            Self::Vertical => margin.right,
            Self::Horizontal => margin.bottom,
        }
    }

    fn bounds(self, main: f64, cross: f64, main_size: f64, cross_size: f64) -> Bounds {
        match self {
            Self::Vertical => Bounds {
                x: cross,
                y: main,
                width: cross_size,
                height: main_size,
            },
            Self::Horizontal => Bounds {
                x: main,
                y: cross,
                width: main_size,
                height: cross_size,
            },
        }
    }
}

enum MainAllocation {
    Fixed(f64),
    Flexible(f64),
}

fn validate_document(document: &DocumentSpec) -> Result<(), LayoutError> {
    validate_positive_finite("document width", document.width)?;
    validate_positive_finite("document height", document.height)?;
    validate_non_negative_finite("document gap", document.gap)?;
    validate_insets("document padding", document.padding)?;
    if document.elements.is_empty() {
        return Err(LayoutError::new("document requires at least one element"));
    }
    if document.elements.len() > MAX_ELEMENTS {
        return Err(LayoutError::new(format!(
            "document element count {} exceeds {MAX_ELEMENTS}",
            document.elements.len()
        )));
    }
    if let Some(columns) = document.columns {
        validate_columns("document columns", columns)?;
    }
    let mut identifiers = HashSet::with_capacity(document.elements.len());
    let mut depths = vec![0usize; document.elements.len()];
    for (index, element) in document.elements.iter().enumerate() {
        validate_element(index, element, document.elements.len())?;
        depths[index] = element.parent.map_or(1, |parent| depths[parent] + 1);
        if depths[index] > MAX_DEPTH {
            return Err(LayoutError::new(format!(
                "element {:?} depth exceeds {MAX_DEPTH}",
                element.id
            )));
        }
        if let Some(parent) = element.parent {
            let parent_element = &document.elements[parent];
            if matches!(
                parent_element.concept,
                Concept::Port | Concept::Line | Concept::Text | Concept::Spacer
            ) {
                return Err(LayoutError::new(format!(
                    "element {:?} cannot be a child of {:?}",
                    element.id, parent_element.id
                )));
            }
        }
        if !identifiers.insert(element.id.as_str()) {
            return Err(LayoutError::new(format!(
                "duplicate element id {:?}",
                element.id
            )));
        }
    }
    Ok(())
}

fn validate_element(index: usize, element: &ElementSpec, count: usize) -> Result<(), LayoutError> {
    if element.id.trim().is_empty() {
        return Err(LayoutError::new("element id must not be empty"));
    }
    if element.id.len() > MAX_ID_BYTES {
        return Err(LayoutError::new(format!(
            "element id exceeds {MAX_ID_BYTES} UTF-8 bytes"
        )));
    }
    if let Some(parent) = element.parent {
        if parent >= index || parent >= count {
            return Err(LayoutError::new(format!(
                "element {:?} has invalid parent index {parent}",
                element.id
            )));
        }
    }
    if element.concept == Concept::Port && element.parent.is_none() {
        return Err(LayoutError::new(format!(
            "port {:?} requires an owning element",
            element.id
        )));
    }
    if element.concept == Concept::Line
        && (element.line.source.trim().is_empty() || element.line.target.trim().is_empty())
    {
        return Err(LayoutError::new(format!(
            "line {:?} requires source and target IDs",
            element.id
        )));
    }
    for (name, value) in [
        ("x", element.x),
        ("y", element.y),
        ("offset x", element.offset_x),
        ("offset y", element.offset_y),
        ("port offset", element.port.offset),
        ("icon offset x", element.icon.offset_x),
        ("icon offset y", element.icon.offset_y),
    ] {
        validate_optional_finite(&format!("element {:?} {name}", element.id), value)?;
    }
    for (name, value) in [
        ("width", element.width),
        ("height", element.height),
        ("intrinsic width", element.intrinsic_width),
        ("intrinsic height", element.intrinsic_height),
        ("minimum width", element.min_width),
        ("maximum width", element.max_width),
        ("minimum height", element.min_height),
        ("maximum height", element.max_height),
        ("weight", element.weight),
        ("stroke width", element.visual.stroke_width),
        ("font size", element.text.font_size),
        ("line height", element.text.line_height),
        ("icon width", element.icon.width),
        ("icon height", element.icon.height),
        ("icon scale", element.icon.scale),
        ("port size", element.port.size),
    ] {
        if let Some(value) = value {
            validate_positive_finite(&format!("element {:?} {name}", element.id), value)?;
        }
    }
    for (name, value) in [
        ("gap", element.gap),
        ("corner radius", element.visual.corner_radius),
        ("obstacle margin", element.line.obstacle_margin),
    ] {
        if let Some(value) = value {
            validate_non_negative_finite(&format!("element {:?} {name}", element.id), value)?;
        }
    }
    for (name, value) in [
        ("opacity", element.visual.opacity),
        ("port anchor", element.port.anchor),
        ("source anchor", element.line.source_anchor),
        ("target anchor", element.line.target_anchor),
        ("label position", element.line.label_position),
    ] {
        if let Some(value) = value {
            validate_unit_interval(&format!("element {:?} {name}", element.id), value)?;
        }
    }
    validate_insets(&format!("element {:?} margin", element.id), element.margin)?;
    validate_insets(
        &format!("element {:?} padding", element.id),
        element.padding,
    )?;
    validate_insets(
        &format!("element {:?} text padding", element.id),
        element.text.padding,
    )?;
    if let (Some(minimum), Some(maximum)) = (element.min_width, element.max_width) {
        if minimum > maximum {
            return Err(LayoutError::new(format!(
                "element {:?} minimum width exceeds maximum width",
                element.id
            )));
        }
    }
    if let (Some(minimum), Some(maximum)) = (element.min_height, element.max_height) {
        if minimum > maximum {
            return Err(LayoutError::new(format!(
                "element {:?} minimum height exceeds maximum height",
                element.id
            )));
        }
    }
    if let Some(columns) = element.columns {
        validate_columns(&format!("element {:?} columns", element.id), columns)?;
    }
    for (name, value) in [
        ("fill", element.visual.fill.as_str()),
        ("stroke", element.visual.stroke.as_str()),
        ("text color", element.text.color.as_str()),
        ("icon color", element.icon.color.as_str()),
    ] {
        validate_safe_paint(&format!("element {:?} {name}", element.id), value)?;
    }
    Ok(())
}

fn validate_resolved_bounds(element: &ElementSpec, bounds: Bounds) -> Result<(), LayoutError> {
    for (name, value) in [
        ("x", bounds.x),
        ("y", bounds.y),
        ("width", bounds.width),
        ("height", bounds.height),
    ] {
        if !value.is_finite() {
            return Err(LayoutError::new(format!(
                "element {:?} resolved {name} must be finite",
                element.id
            )));
        }
    }
    if element.concept != Concept::Line && (bounds.width <= 0.0 || bounds.height <= 0.0) {
        return Err(LayoutError::new(format!(
            "element {:?} resolved size must be positive",
            element.id
        )));
    }
    for (name, value, minimum, maximum) in [
        ("width", bounds.width, element.min_width, element.max_width),
        (
            "height",
            bounds.height,
            element.min_height,
            element.max_height,
        ),
    ] {
        if minimum.is_some_and(|limit| value < limit) || maximum.is_some_and(|limit| value > limit)
        {
            return Err(LayoutError::new(format!(
                "element {:?} resolved {name} {} violates its constraints",
                element.id,
                format_number(value)
            )));
        }
    }
    Ok(())
}

fn resolved_element(element: &ElementSpec, bounds: Bounds) -> Result<ResolvedElement, LayoutError> {
    validate_resolved_bounds(element, bounds)?;
    let shape = match element.visual.shape {
        Shape::Default => default_shape(element.concept),
        value => value,
    };
    let (default_fill, default_stroke) = default_colors(element.concept);
    let visible = element
        .visual
        .visible
        .unwrap_or(element.concept != Concept::Spacer);
    let icon_ref = if !element.icon.reference.is_empty() {
        element.icon.reference.clone()
    } else if !element.icon.fallback_reference.is_empty() {
        element.icon.fallback_reference.clone()
    } else {
        if element.icon.missing_policy == MissingIconPolicy::Error
            && element.concept == Concept::Item
        {
            return Err(LayoutError::new(format!(
                "element {:?} requires an icon",
                element.id
            )));
        }
        String::new()
    };
    Ok(ResolvedElement {
        parent: element.parent,
        id: element.id.clone(),
        concept: element.concept,
        x: bounds.x,
        y: bounds.y,
        width: bounds.width,
        height: bounds.height,
        visual: ResolvedVisual {
            shape,
            fill: if element.visual.fill.is_empty() {
                default_fill.to_owned()
            } else {
                element.visual.fill.clone()
            },
            stroke: if element.visual.stroke.is_empty() {
                default_stroke.to_owned()
            } else {
                element.visual.stroke.clone()
            },
            stroke_width: element.visual.stroke_width.unwrap_or(1.5),
            corner_radius: element.visual.corner_radius.unwrap_or(4.0),
            opacity: element.visual.opacity.unwrap_or(1.0),
            visible,
            layer: element.visual.layer.unwrap_or(0),
        },
        text: ResolvedText {
            value: element.text.value.clone(),
            font_family: if element.text.font_family.is_empty() {
                "sans-serif".to_owned()
            } else {
                element.text.font_family.clone()
            },
            color: if element.text.color.is_empty() {
                "#0f172a".to_owned()
            } else {
                element.text.color.clone()
            },
            role: element.text.role.clone(),
            font_size: element.text.font_size.unwrap_or(DEFAULT_FONT_SIZE),
            line_height: element.text.line_height.unwrap_or(DEFAULT_LINE_HEIGHT),
        },
        icon_ref,
        line: ResolvedLine {
            style: element.line.style,
            source_decoration: element.line.source_decoration,
            target_decoration: element.line.target_decoration,
            label: element.line.label.clone(),
            label_position: element.line.label_position.unwrap_or(0.5),
        },
        points: Vec::new(),
    })
}

fn effective_layout(policy: LayoutPolicy) -> LayoutPolicy {
    match policy {
        LayoutPolicy::Default => LayoutPolicy::Vertical,
        value => value,
    }
}

fn intrinsic_main(element: &ElementSpec, axis: Axis) -> Option<f64> {
    let explicit = match axis {
        Axis::Vertical => element.intrinsic_height,
        Axis::Horizontal => element.intrinsic_width,
    };
    explicit
        .or_else(|| icon_size(element, axis))
        .or_else(|| match element.concept {
            Concept::Item => Some(DEFAULT_ITEM_SIZE),
            Concept::Text => Some(match axis {
                Axis::Vertical => measured_text_size(element).1,
                Axis::Horizontal => measured_text_size(element).0,
            }),
            Concept::Spacer => Some(1.0),
            _ => None,
        })
}

fn intrinsic_cross(element: &ElementSpec, axis: Axis) -> Option<f64> {
    let explicit = match axis {
        Axis::Vertical => element.intrinsic_width,
        Axis::Horizontal => element.intrinsic_height,
    };
    explicit.or_else(|| icon_cross_size(element, axis))
}

fn icon_size(element: &ElementSpec, axis: Axis) -> Option<f64> {
    let scale = element.icon.scale.unwrap_or(1.0);
    match axis {
        Axis::Vertical => element.icon.height.map(|value| value * scale),
        Axis::Horizontal => element.icon.width.map(|value| value * scale),
    }
}

fn icon_cross_size(element: &ElementSpec, axis: Axis) -> Option<f64> {
    let scale = element.icon.scale.unwrap_or(1.0);
    match axis {
        Axis::Vertical => element.icon.width.map(|value| value * scale),
        Axis::Horizontal => element.icon.height.map(|value| value * scale),
    }
}

fn measured_text_size(element: &ElementSpec) -> (f64, f64) {
    let font_size = element.text.font_size.unwrap_or(DEFAULT_FONT_SIZE);
    let line_height = element.text.line_height.unwrap_or(DEFAULT_LINE_HEIGHT);
    let padding = element.text.padding.resolved();
    let lines = element.text.value.lines().collect::<Vec<_>>();
    let line_count = lines.len().max(1) as f64;
    let longest = lines
        .iter()
        .map(|line| line.chars().count())
        .max()
        .unwrap_or(0) as f64;
    (
        (longest * font_size * 0.6 + padding.left + padding.right).max(1.0),
        (line_count * font_size * line_height + padding.top + padding.bottom).max(1.0),
    )
}

fn default_absolute_width(element: &ElementSpec, available: f64) -> f64 {
    if let Some(width) = element.icon.width {
        return width * element.icon.scale.unwrap_or(1.0);
    }
    match element.concept {
        Concept::Item => DEFAULT_ITEM_SIZE,
        Concept::Text => measured_text_size(element).0,
        Concept::Spacer => 1.0,
        _ => available,
    }
}

fn default_absolute_height(element: &ElementSpec, available: f64) -> f64 {
    if let Some(height) = element.icon.height {
        return height * element.icon.scale.unwrap_or(1.0);
    }
    match element.concept {
        Concept::Item => DEFAULT_ITEM_SIZE,
        Concept::Text => measured_text_size(element).1,
        Concept::Spacer => 1.0,
        _ => available,
    }
}

fn default_shape(concept: Concept) -> Shape {
    match concept {
        Concept::Frame | Concept::Group | Concept::Capture | Concept::Item | Concept::Port => {
            Shape::Rectangle
        }
        Concept::Line | Concept::Text | Concept::Spacer => Shape::None,
    }
}

fn default_colors(concept: Concept) -> (&'static str, &'static str) {
    match concept {
        Concept::Frame => ("#ffffff", "#475569"),
        Concept::Group => ("#f8fafc", "#64748b"),
        Concept::Capture => ("#fff7ed", "#f97316"),
        Concept::Item => ("#ffffff", "#334155"),
        Concept::Port => ("#334155", "#334155"),
        Concept::Line => ("none", "#334155"),
        Concept::Text | Concept::Spacer => ("none", "none"),
    }
}

fn bounds_of(element: &ResolvedElement) -> Bounds {
    Bounds {
        x: element.x,
        y: element.y,
        width: element.width,
        height: element.height,
    }
}

fn contains(parent: Bounds, child: Bounds) -> bool {
    const EPSILON: f64 = 1e-9;
    child.x + EPSILON >= parent.x
        && child.y + EPSILON >= parent.y
        && child.x + child.width <= parent.x + parent.width + EPSILON
        && child.y + child.height <= parent.y + parent.height + EPSILON
}

fn resolve_auto_side(side: Side, from: Point, to: Point) -> Side {
    if side != Side::Auto {
        return side;
    }
    let dx = to.x - from.x;
    let dy = to.y - from.y;
    if dx.abs() >= dy.abs() {
        if dx >= 0.0 {
            Side::Right
        } else {
            Side::Left
        }
    } else if dy >= 0.0 {
        Side::Bottom
    } else {
        Side::Top
    }
}

fn anchor_point(bounds: Bounds, side: Side, anchor: f64) -> Point {
    match side {
        Side::Top => Point {
            x: bounds.x + bounds.width * anchor,
            y: bounds.y,
        },
        Side::Right => Point {
            x: bounds.x + bounds.width,
            y: bounds.y + bounds.height * anchor,
        },
        Side::Bottom => Point {
            x: bounds.x + bounds.width * anchor,
            y: bounds.y + bounds.height,
        },
        Side::Left => Point {
            x: bounds.x,
            y: bounds.y + bounds.height * anchor,
        },
        Side::Auto => bounds.center(),
    }
}

fn deduplicate_points(points: Vec<Point>) -> Vec<Point> {
    let mut result = Vec::with_capacity(points.len());
    for point in points {
        if result
            .last()
            .is_none_or(|previous: &Point| previous.x != point.x || previous.y != point.y)
        {
            result.push(point);
        }
    }
    result
}

fn route_score(points: &[Point], obstacles: &[Bounds]) -> f64 {
    let intersections = points
        .windows(2)
        .map(|segment| {
            obstacles
                .iter()
                .filter(|obstacle| segment_intersects_rect(segment[0], segment[1], **obstacle))
                .count()
        })
        .sum::<usize>();
    let length = points
        .windows(2)
        .map(|segment| (segment[1].x - segment[0].x).abs() + (segment[1].y - segment[0].y).abs())
        .sum::<f64>();
    intersections as f64 * 1_000_000_000.0 + length
}

fn segment_intersects_rect(start: Point, end: Point, rect: Bounds) -> bool {
    if start.x == end.x {
        start.x > rect.x
            && start.x < rect.x + rect.width
            && start.y.max(end.y) > rect.y
            && start.y.min(end.y) < rect.y + rect.height
    } else if start.y == end.y {
        start.y > rect.y
            && start.y < rect.y + rect.height
            && start.x.max(end.x) > rect.x
            && start.x.min(end.x) < rect.x + rect.width
    } else {
        false
    }
}

fn validate_columns(name: &str, value: u16) -> Result<(), LayoutError> {
    if value == 0 || value > MAX_COLUMNS {
        return Err(LayoutError::new(format!(
            "{name} must be between 1 and {MAX_COLUMNS}"
        )));
    }
    Ok(())
}

fn validate_insets(name: &str, value: Insets) -> Result<(), LayoutError> {
    for (side, value) in [
        ("top", value.top),
        ("right", value.right),
        ("bottom", value.bottom),
        ("left", value.left),
    ] {
        if let Some(value) = value {
            validate_non_negative_finite(&format!("{name} {side}"), value)?;
        }
    }
    Ok(())
}

fn validate_optional_finite(name: &str, value: Option<f64>) -> Result<(), LayoutError> {
    if value.is_some_and(|value| !value.is_finite()) {
        return Err(LayoutError::new(format!("{name} must be finite")));
    }
    Ok(())
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

fn validate_unit_interval(name: &str, value: f64) -> Result<(), LayoutError> {
    if !value.is_finite() || !(0.0..=1.0).contains(&value) {
        return Err(LayoutError::new(format!(
            "{name} must be finite and between 0 and 1"
        )));
    }
    Ok(())
}

fn validate_safe_paint(name: &str, value: &str) -> Result<(), LayoutError> {
    if value.is_empty() {
        return Ok(());
    }
    let lowercase = value.to_ascii_lowercase();
    if lowercase.contains("url")
        || lowercase.contains("javascript")
        || lowercase.contains("expression")
        || lowercase.contains("@import")
        || value.chars().any(|character| {
            !character.is_ascii_alphanumeric()
                && !matches!(
                    character,
                    '#' | '(' | ')' | ',' | '.' | '%' | '-' | '_' | ' '
                )
        })
    {
        return Err(LayoutError::new(format!(
            "{name} contains an unsafe paint value"
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
    #[rustfmt::skip]
    use crate::ent::model::document::{
        Decoration,
        IconSpec,
        LineSpec,
        LineStyle,
        PortSpec,
        TextSpec,
        VisualSpec,
    };

    fn empty_text() -> TextSpec {
        TextSpec {
            value: String::new(),
            font_family: String::new(),
            color: String::new(),
            role: String::new(),
            font_size: None,
            line_height: None,
            wrap: None,
            fit: None,
            clip: None,
            padding: Insets::default(),
        }
    }

    fn element(id: &str, concept: Concept, parent: Option<usize>) -> ElementSpec {
        ElementSpec {
            parent,
            id: id.to_owned(),
            concept,
            layout: LayoutPolicy::Default,
            overflow: Overflow::Error,
            align: Alignment::Stretch,
            justify: Justification::Start,
            x: None,
            y: None,
            width: None,
            height: None,
            intrinsic_width: None,
            intrinsic_height: None,
            min_width: None,
            max_width: None,
            min_height: None,
            max_height: None,
            offset_x: None,
            offset_y: None,
            weight: None,
            gap: None,
            margin: Insets::default(),
            padding: Insets::default(),
            columns: None,
            column_span: None,
            row_span: None,
            visual: VisualSpec {
                shape: Shape::Default,
                fill: String::new(),
                stroke: String::new(),
                stroke_width: None,
                corner_radius: None,
                opacity: None,
                visible: None,
                layer: None,
            },
            text: empty_text(),
            icon: IconSpec {
                reference: String::new(),
                fallback_reference: String::new(),
                color: String::new(),
                width: None,
                height: None,
                scale: None,
                offset_x: None,
                offset_y: None,
                missing_policy: MissingIconPolicy::Fallback,
            },
            port: PortSpec {
                side: Side::Auto,
                anchor: None,
                offset: None,
                size: None,
                visible: None,
                label: String::new(),
            },
            line: LineSpec {
                source: String::new(),
                target: String::new(),
                source_side: Side::Auto,
                target_side: Side::Auto,
                source_anchor: None,
                target_anchor: None,
                routing: RoutingPolicy::Orthogonal,
                obstacle_margin: None,
                style: LineStyle::Solid,
                source_decoration: Decoration::None,
                target_decoration: Decoration::None,
                label: String::new(),
                label_position: None,
            },
        }
    }

    fn document(elements: Vec<ElementSpec>) -> DocumentSpec {
        DocumentSpec {
            layout: LayoutPolicy::Vertical,
            width: 200.0,
            height: 300.0,
            gap: 10.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements,
        }
    }

    #[test]
    fn allocates_fixed_before_flexible_vertical_children() {
        let mut header = element("header", Concept::Item, None);
        header.height = Some(40.0);
        let mut body = element("body", Concept::Item, None);
        body.width = Some(160.0);
        body.weight = Some(1.0);
        let mut footer = element("footer", Concept::Item, None);
        footer.weight = Some(2.0);
        let resolved = resolve(&document(vec![header, body, footer])).expect("resolve document");
        assert_eq!(resolved.elements[0].height, 40.0);
        assert_eq!(resolved.elements[1].y, 50.0);
        assert_eq!(resolved.elements[1].height, 80.0);
        assert_eq!(resolved.elements[1].width, 160.0);
        assert_eq!(resolved.elements[2].y, 140.0);
        assert_eq!(resolved.elements[2].height, 160.0);
        assert_eq!(resolved.elements[2].width, 200.0);
    }

    #[test]
    fn resolves_nested_grid_ports_and_orthogonal_lines() {
        let mut frame = element("frame", Concept::Frame, None);
        frame.width = Some(200.0);
        frame.height = Some(300.0);
        frame.layout = LayoutPolicy::Grid;
        frame.columns = Some(12);
        frame.gap = Some(4.0);
        let mut left = element("left", Concept::Item, Some(0));
        left.column_span = Some(6);
        let mut port = element("left-out", Concept::Port, Some(1));
        port.port.side = Side::Right;
        let mut right = element("right", Concept::Capture, Some(0));
        right.column_span = Some(6);
        let mut line = element("flow", Concept::Line, Some(0));
        line.line.source = "left-out".to_owned();
        line.line.target = "right".to_owned();
        line.line.target_decoration = Decoration::Arrow;
        let resolved = resolve(&DocumentSpec {
            layout: LayoutPolicy::Absolute,
            width: 200.0,
            height: 300.0,
            gap: 0.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements: vec![frame, left, port, right, line],
        })
        .expect("resolve generic concepts");
        assert_eq!(resolved.elements[1].parent, Some(0));
        assert_eq!(resolved.elements[2].concept, Concept::Port);
        assert!(resolved.elements[4].points.len() >= 2);
        assert_eq!(
            resolved.elements[4].line.target_decoration,
            Decoration::Arrow
        );
    }

    #[test]
    fn rejects_non_finite_and_conflicting_constraints() {
        let mut invalid = element("bad", Concept::Item, None);
        invalid.width = Some(f64::NAN);
        assert!(resolve(&document(vec![invalid]))
            .expect_err("non-finite width")
            .to_string()
            .contains("width"));

        let mut constrained = element("bad", Concept::Item, None);
        constrained.min_width = Some(20.0);
        constrained.max_width = Some(10.0);
        assert!(resolve(&document(vec![constrained]))
            .expect_err("conflicting constraints")
            .to_string()
            .contains("minimum width"));

        let mut unsafe_paint = element("bad", Concept::Item, None);
        unsafe_paint.visual.fill = "url(https://example.com/pixel.svg)".to_owned();
        assert!(resolve(&document(vec![unsafe_paint]))
            .expect_err("unsafe paint")
            .to_string()
            .contains("unsafe paint"));
    }

    #[test]
    fn output_is_deterministic() {
        let mut left = element("left", Concept::Item, None);
        left.width = Some(40.0);
        let mut right = element("right", Concept::Item, None);
        right.weight = Some(1.0);
        let document = DocumentSpec {
            layout: LayoutPolicy::Horizontal,
            width: 240.0,
            height: 80.0,
            gap: 8.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements: vec![left, right],
        };
        assert_eq!(resolve(&document), resolve(&document));
    }

    #[test]
    fn missing_icon_hide_preserves_the_item_slot() {
        let mut item = element("service", Concept::Item, None);
        item.height = Some(48.0);
        item.icon.missing_policy = MissingIconPolicy::Hide;
        let resolved = resolve(&document(vec![item])).expect("resolve hidden icon slot");
        assert!(resolved.elements[0].visual.visible);
        assert!(resolved.elements[0].icon_ref.is_empty());
        assert_eq!(resolved.elements[0].height, 48.0);
    }

    #[test]
    fn validates_root_offsets_but_places_ports_against_the_owner_border() {
        let mut shifted = element("shifted", Concept::Item, None);
        shifted.height = Some(20.0);
        shifted.offset_x = Some(1.0);
        assert!(resolve(&document(vec![shifted]))
            .expect_err("root offset overflow")
            .to_string()
            .contains("parent content box"));

        let mut owner = element("owner", Concept::Item, None);
        owner.width = Some(100.0);
        owner.height = Some(100.0);
        owner.padding = Insets {
            top: Some(20.0),
            right: Some(20.0),
            bottom: Some(20.0),
            left: Some(20.0),
        };
        let mut port = element("owner-out", Concept::Port, Some(0));
        port.port.side = Side::Right;
        let resolved = resolve(&DocumentSpec {
            layout: LayoutPolicy::Absolute,
            width: 100.0,
            height: 100.0,
            gap: 0.0,
            padding: Insets::default(),
            overflow: Overflow::Error,
            columns: None,
            elements: vec![owner, port],
        })
        .expect("port on padded owner border");
        assert_eq!(resolved.elements[1].x, 92.0);
    }
}
