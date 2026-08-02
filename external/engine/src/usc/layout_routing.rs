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
        crate::usc::cancel::check().map_err(LayoutError::new)?;
        let id_index = self
            .document
            .elements
            .iter()
            .enumerate()
            .map(|(index, element)| (element.id.as_str(), index))
            .collect::<HashMap<_, _>>();
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
