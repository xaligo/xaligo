impl LayoutState<'_> {
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
        let boundary_icon = is_boundary_icon_port(element, owner_bounds, bounds);
        if !contains(owner_bounds, bounds) && !boundary_icon {
            return Err(LayoutError::new(format!(
                "port {:?} lies outside owner {:?}",
                element.id, owner.id
            )));
        }
        for previous_index in 0..index {
            let previous = &self.document.elements[previous_index];
            if previous.concept != Concept::Port || previous.parent != element.parent {
                continue;
            }
            let Some(previous_resolved) = self.resolved[previous_index].as_ref() else {
                continue;
            };
            let previous_bounds = bounds_of(previous_resolved);
            if (boundary_icon || is_boundary_icon_port(previous, owner_bounds, previous_bounds))
                && port_bounds_overlap(bounds, previous_bounds)
            {
                return Err(LayoutError::new(format!(
                    "port {:?} overlaps sibling port {:?}",
                    element.id, previous.id
                )));
            }
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
        crate::usc::cancel::check().map_err(LayoutError::new)?;
        let id_index = self
            .document
            .elements
            .iter()
            .enumerate()
            .map(|(index, element)| (element.id.as_str(), index))
            .collect::<HashMap<_, _>>();
        let mut lane_counts = HashMap::<(usize, usize), usize>::new();
        for index in 0..self.document.elements.len() {
            crate::usc::cancel::check().map_err(LayoutError::new)?;
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
            let source_bounds = connection_bounds_of(source);
            let target_bounds = connection_bounds_of(target);
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
            let start = connection_anchor_point(
                source,
                source_side,
                element.line.source_anchor.unwrap_or(0.5),
            );
            let end = connection_anchor_point(
                target,
                target_side,
                element.line.target_anchor.unwrap_or(0.5),
            );
            let mut points = match element.line.routing {
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
            let lane_key = (source_index, target_index);
            let lane_index = lane_counts.entry(lane_key).or_insert(0);
            if *lane_index > 0 {
                points = separate_parallel_lane(points, *lane_index);
            }
            *lane_index += 1;
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
        let route_corridor = Bounds {
            x: start.x.min(end.x) - margin,
            y: start.y.min(end.y) - margin,
            width: (start.x - end.x).abs() + margin * 2.0,
            height: (start.y - end.y).abs() + margin * 2.0,
        };
        for obstacle in obstacle_bounds
            .iter()
            .filter(|obstacle| rectangles_overlap(route_corridor, **obstacle))
        {
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

// Icon-only items have a wider text/layout slot. Connect to their visible icon,
// not to the invisible slot edge. This is a generic visual rule, not a domain
// resource lookup; group and port boundary geometry is unchanged.
fn connection_bounds_of(element: &ResolvedElement) -> Bounds {
    if element.concept == Concept::Item
        && element.visual.shape == Shape::None
        && !element.icon_ref.is_empty()
        && element.icon_width > 0.0
        && element.icon_height > 0.0
    {
        return Bounds {
            x: element.icon_x,
            y: element.icon_y,
            width: element.icon_width,
            height: element.icon_height,
        };
    }
    bounds_of(element)
}

fn connection_anchor_point(element: &ResolvedElement, side: Side, anchor: f64) -> Point {
    let mut bounds = connection_bounds_of(element);
    // A bottom connection must leave below the label, not strike through it.
    // Other sides bind directly to the visible icon.
    if side == Side::Bottom
        && element.concept == Concept::Item
        && element.visual.shape == Shape::None
        && !element.icon_ref.is_empty()
        && !element.text.value.is_empty()
    {
        bounds = Bounds {
            x: element.text.x,
            y: element.text.y,
            width: element.text.width,
            height: element.text.height,
        };
    }
    anchor_point(bounds, side, anchor)
}

fn is_boundary_icon_port(element: &ElementSpec, owner: Bounds, port: Bounds) -> bool {
    if element.visual.shape != Shape::None
        || (element.icon.reference.is_empty() && element.icon.fallback_reference.is_empty())
    {
        return false;
    }
    const EPSILON: f64 = 1e-9;
    match element.port.side {
        Side::Top | Side::Auto => {
            port.x + EPSILON >= owner.x
                && port.x + port.width <= owner.x + owner.width + EPSILON
                && port.y < owner.y - EPSILON
                && port.y + port.height > owner.y + EPSILON
                && port.y + port.height < owner.y + owner.height - EPSILON
        }
        Side::Right => {
            port.y + EPSILON >= owner.y
                && port.y + port.height <= owner.y + owner.height + EPSILON
                && port.x < owner.x + owner.width - EPSILON
                && port.x + port.width > owner.x + owner.width + EPSILON
                && port.x > owner.x + EPSILON
        }
        Side::Bottom => {
            port.x + EPSILON >= owner.x
                && port.x + port.width <= owner.x + owner.width + EPSILON
                && port.y < owner.y + owner.height - EPSILON
                && port.y + port.height > owner.y + owner.height + EPSILON
                && port.y > owner.y + EPSILON
        }
        Side::Left => {
            port.y + EPSILON >= owner.y
                && port.y + port.height <= owner.y + owner.height + EPSILON
                && port.x < owner.x - EPSILON
                && port.x + port.width > owner.x + EPSILON
                && port.x + port.width < owner.x + owner.width - EPSILON
        }
    }
}

fn port_bounds_overlap(left: Bounds, right: Bounds) -> bool {
    const EPSILON: f64 = 1e-9;
    left.x < right.x + right.width - EPSILON
        && left.x + left.width > right.x + EPSILON
        && left.y < right.y + right.height - EPSILON
        && left.y + left.height > right.y + EPSILON
}

fn separate_parallel_lane(points: Vec<Point>, lane_index: usize) -> Vec<Point> {
    if points.len() < 2 {
        return points;
    }
    let magnitude = ((lane_index + 1) / 2) as f64 * 6.0;
    let offset = if lane_index % 2 == 1 { magnitude } else { -magnitude };
    let start = points[0];
    let end = points[points.len() - 1];
    let horizontal = (end.x - start.x).abs() >= (end.y - start.y).abs();
    let mut separated = Vec::with_capacity(points.len() + 4);
    separated.push(start);
    if horizontal {
        separated.push(Point { x: start.x, y: start.y + offset });
        for point in points.iter().skip(1).take(points.len().saturating_sub(2)) {
            separated.push(Point { x: point.x, y: point.y + offset });
        }
        separated.push(Point { x: end.x, y: end.y + offset });
    } else {
        separated.push(Point { x: start.x + offset, y: start.y });
        for point in points.iter().skip(1).take(points.len().saturating_sub(2)) {
            separated.push(Point { x: point.x + offset, y: point.y });
        }
        separated.push(Point { x: end.x + offset, y: end.y });
    }
    separated.push(end);
    deduplicate_points(separated)
}
