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
        crate::usc::cancel::check().map_err(LayoutError::new)?;
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
        crate::usc::cancel::check().map_err(LayoutError::new)?;
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
        crate::usc::cancel::check().map_err(LayoutError::new)?;
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

}
