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
    let fill = if element.visual.fill.is_empty() {
        default_fill.to_owned()
    } else {
        element.visual.fill.clone()
    };
    let stroke = if element.visual.stroke.is_empty() {
        default_stroke.to_owned()
    } else {
        element.visual.stroke.clone()
    };
    let font_size = element.text.font_size.unwrap_or(DEFAULT_FONT_SIZE);
    let line_height = element.text.line_height.unwrap_or(DEFAULT_LINE_HEIGHT);
    let (text_x, text_y, text_width, text_height, icon_x, icon_y, icon_width, icon_height) =
        resolved_foreground_geometry(element, bounds, &icon_ref, font_size, line_height);
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
            fill,
            stroke: stroke.clone(),
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
                if matches!(element.concept, Concept::Group | Concept::Frame) {
                    stroke
                } else {
                    "#0f172a".to_owned()
                }
            } else {
                element.text.color.clone()
            },
            role: element.text.role.clone(),
            font_size,
            line_height,
            x: text_x,
            y: text_y,
            width: text_width,
            height: text_height,
        },
        icon_ref,
        icon_x,
        icon_y,
        icon_width,
        icon_height,
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

fn resolved_foreground_geometry(
    element: &ElementSpec,
    bounds: Bounds,
    icon_ref: &str,
    font_size: f64,
    line_height: f64,
) -> (f64, f64, f64, f64, f64, f64, f64, f64) {
    let mut text_x = bounds.x;
    let mut text_y = bounds.y;
    let mut text_width = bounds.width;
    let mut text_height = bounds.height;
    let mut icon_x = 0.0;
    let mut icon_y = 0.0;
    let mut icon_width = 0.0;
    let mut icon_height = 0.0;
    let has_text = !element.text.value.is_empty();
    let has_icon = !icon_ref.is_empty();
    let icon_scale = element.icon.scale.unwrap_or(1.0);

    if matches!(element.concept, Concept::Group | Concept::Frame) && has_text {
        let header_height = 28.0_f64.min(bounds.height);
        let tip = (header_height / 2.0).min(10.0);
        let icon_space = if has_icon { 26.0 } else { 0.0 };
        let header_width = ((element.text.value.chars().count() as f64 * font_size * 0.62)
            + 28.0
            + icon_space)
            .min(bounds.width)
            .max(header_height);
        text_x = bounds.x + 4.0 + icon_space;
        text_y = bounds.y;
        text_width = (header_width - tip - 4.0 - icon_space).max(1.0);
        text_height = header_height;
        if has_icon {
            icon_width = element
                .icon
                .width
                .map(|value| value * icon_scale)
                .unwrap_or(20.0)
                .min(20.0)
                .min(header_height);
            icon_height = element
                .icon
                .height
                .map(|value| value * icon_scale)
                .unwrap_or(20.0)
                .min(20.0)
                .min(header_height);
            icon_x = bounds.x + 4.0 + element.icon.offset_x.unwrap_or(0.0);
            icon_y = bounds.y
                + (header_height - icon_height) / 2.0
                + element.icon.offset_y.unwrap_or(0.0);
        }
    } else if has_icon {
        icon_width = element
            .icon
            .width
            .map(|value| value * icon_scale)
            .unwrap_or(32.0)
            .min(bounds.width);
        let line_count = element.text.value.lines().count().max(1) as f64;
        let label_height = if has_text {
            font_size * line_height * line_count
        } else {
            0.0
        };
        let gap = if has_text { 4.0 } else { 0.0 };
        icon_height = element
            .icon
            .height
            .map(|value| value * icon_scale)
            .unwrap_or(32.0)
            .min((bounds.height - label_height - gap).max(1.0));
        let block_height = icon_height + gap + label_height;
        icon_x = bounds.x
            + (bounds.width - icon_width) / 2.0
            + element.icon.offset_x.unwrap_or(0.0);
        icon_y = bounds.y
            + (bounds.height - block_height).max(0.0) / 2.0
            + element.icon.offset_y.unwrap_or(0.0);
        if has_text {
            text_y = icon_y + icon_height + gap;
            text_height = label_height;
        }
    }

    (
        text_x,
        text_y,
        text_width,
        text_height,
        icon_x,
        icon_y,
        icon_width,
        icon_height,
    )
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
