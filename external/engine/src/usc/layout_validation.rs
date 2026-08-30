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
    resolved_element_with_icon_limit(element, bounds, None)
}

fn resolved_element_with_icon_limit(
    element: &ElementSpec,
    bounds: Bounds,
    icon_limit: Option<f64>,
) -> Result<ResolvedElement, LayoutError> {
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
        resolved_foreground_geometry(element, bounds, &icon_ref, font_size, line_height, icon_limit);
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
    icon_limit: Option<f64>,
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

    if uses_v1_profile_group_header(element) && has_text {
        let header_x = bounds.x - 2.0;
        let label_width = v1_group_label_width(&element.text.value);
        icon_width = if has_icon {
            element
                .icon
                .width
                .map(|value| value * icon_scale)
                .unwrap_or(32.0)
        } else {
            0.0
        };
        icon_height = if has_icon {
            element
                .icon
                .height
                .map(|value| value * icon_scale)
                .unwrap_or(32.0)
        } else {
            0.0
        };
        let header_height = if has_icon {
            icon_height.max(20.0)
        } else {
            20.0
        };
        let header_y = bounds.y - header_height / 2.0;
        text_x = header_x + if has_icon { icon_width } else { 0.0 } + 4.0;
        text_y = header_y + (header_height - 18.0) / 2.0;
        text_width = label_width;
        text_height = 18.0;
        if has_icon {
            icon_x = header_x;
            icon_y = header_y + (header_height - icon_height) / 2.0;
        }
    } else if matches!(element.concept, Concept::Group | Concept::Frame) && has_text {
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
            .min(icon_limit.unwrap_or(f64::INFINITY))
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
            .min(icon_limit.unwrap_or(f64::INFINITY))
            .min((bounds.height - label_height - gap).max(1.0));
        let block_height = icon_height + gap + label_height;
        icon_x = bounds.x
            + (bounds.width - icon_width) / 2.0
            + element.icon.offset_x.unwrap_or(0.0);
        icon_y = bounds.y
            + if icon_limit.is_some() {
                0.0
            } else {
                (bounds.height - block_height).max(0.0) / 2.0
            }
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

const V1_GROUP_HEADER_ROLE: &str = "group-header";
const V1_GROUP_HEADER_GAP: f64 = 4.0;
const V1_GROUP_HEADER_BUCKET_HEIGHT: f64 = 64.0;

fn uses_v1_profile_group_header(element: &ElementSpec) -> bool {
    matches!(element.concept, Concept::Group | Concept::Capture)
        && element.text.role == V1_GROUP_HEADER_ROLE
}

fn resolved_uses_v1_profile_group_header(element: &ResolvedElement) -> bool {
    matches!(element.concept, Concept::Group | Concept::Capture)
        && element.text.role == V1_GROUP_HEADER_ROLE
}

fn v1_group_label_width(value: &str) -> f64 {
    (value.chars().map(v1_display_columns).sum::<f64>() * 9.6).ceil() + 8.0
}

fn v1_display_columns(value: char) -> f64 {
    let codepoint = value as u32;
    if value == '\t' {
        4.0
    } else if codepoint < 0x20 {
        0.0
    } else if (0x1100..=0x115f).contains(&codepoint)
        || matches!(codepoint, 0x2329 | 0x232a)
        || ((0x2e80..=0xa4cf).contains(&codepoint) && codepoint != 0x303f)
        || (0xac00..=0xd7a3).contains(&codepoint)
        || (0xf900..=0xfaff).contains(&codepoint)
        || (0xfe10..=0xfe19).contains(&codepoint)
        || (0xfe30..=0xfe6f).contains(&codepoint)
        || (0xff00..=0xff60).contains(&codepoint)
        || (0xffe0..=0xffe6).contains(&codepoint)
    {
        2.0
    } else {
        1.0
    }
}

fn v1_resolved_group_header_bounds(element: &ResolvedElement) -> Bounds {
    let has_icon = element.icon_width > 0.0 && element.icon_height > 0.0;
    let height = if has_icon {
        element.icon_height.max(20.0)
    } else {
        20.0
    };
    let x = if has_icon {
        element.icon_x
    } else {
        element.text.x - 4.0
    };
    let y = if has_icon {
        element.icon_y - (height - element.icon_height) / 2.0
    } else {
        element.text.y - (height - element.text.height) / 2.0
    };
    let tip = (height / 2.0).min(14.0);
    Bounds {
        x,
        y,
        width: element.text.x + element.text.width + 18.0 + tip - x,
        height,
    }
}

struct V1GroupHeaderObstacle {
    x: f64,
    width: f64,
    y: f64,
    height: f64,
    header: bool,
}

impl Clone for V1GroupHeaderObstacle {
    fn clone(&self) -> Self {
        *self
    }
}

impl Copy for V1GroupHeaderObstacle {}

struct V1GroupHeaderObstacleIndex {
    buckets: HashMap<i64, Vec<V1GroupHeaderObstacle>>,
}

impl V1GroupHeaderObstacleIndex {
    fn new() -> Self {
        Self {
            buckets: HashMap::new(),
        }
    }

    fn insert(&mut self, obstacle: V1GroupHeaderObstacle) {
        let first = v1_group_header_bucket(obstacle.y);
        let last = v1_group_header_bucket(obstacle.y + obstacle.height);
        for bucket in first..=last {
            self.buckets.entry(bucket).or_default().push(obstacle);
        }
    }

    fn avoid(&self, x: f64, width: f64, y: f64, height: f64) -> f64 {
        let minimum_y = y - V1_GROUP_HEADER_GAP;
        let maximum_y = y + height + V1_GROUP_HEADER_GAP;
        let first = v1_group_header_bucket(minimum_y);
        let last = v1_group_header_bucket(maximum_y);
        let mut adjusted = y;
        for bucket in first..=last {
            let Some(obstacles) = self.buckets.get(&bucket) else {
                continue;
            };
            for obstacle in obstacles {
                if horizontal_overlap(
                    x,
                    x + width,
                    obstacle.x,
                    obstacle.x + obstacle.width,
                ) <= 0.0
                {
                    continue;
                }
                if obstacle.header {
                    if obstacle.y < maximum_y && obstacle.y + obstacle.height > minimum_y {
                        adjusted = adjusted.max(
                            obstacle.y + obstacle.height + V1_GROUP_HEADER_GAP,
                        );
                    }
                } else if obstacle.y >= minimum_y && obstacle.y <= maximum_y {
                    adjusted = adjusted.max(obstacle.y + V1_GROUP_HEADER_GAP);
                }
            }
        }
        adjusted
    }
}

fn v1_group_header_bucket(y: f64) -> i64 {
    (y / V1_GROUP_HEADER_BUCKET_HEIGHT).floor() as i64
}

fn horizontal_overlap(a0: f64, a1: f64, b0: f64, b1: f64) -> f64 {
    a1.max(a0).min(b1.max(b0)) - a0.min(a1).max(b0.min(b1))
}

impl LayoutState<'_> {
    fn align_v1_profile_group_headers(&mut self) -> Result<(), LayoutError> {
        let mut enclosing_frames = vec![None; self.document.elements.len()];
        for (index, element) in self.document.elements.iter().enumerate() {
            enclosing_frames[index] = if element.concept == Concept::Frame {
                Some(index)
            } else {
                element.parent.and_then(|parent| enclosing_frames[parent])
            };
        }

        let mut metadata_floors: Vec<Option<f64>> =
            vec![None; self.document.elements.len()];
        for (index, resolved) in self.resolved.iter().enumerate() {
            let Some(element) = resolved.as_ref() else {
                continue;
            };
            if !element.text.role.starts_with("frame-metadata-") {
                continue;
            }
            let Some(frame) = enclosing_frames[index] else {
                continue;
            };
            let floor = element.y + element.height + 8.0;
            metadata_floors[frame] = Some(metadata_floors[frame].unwrap_or(floor).max(floor));
        }

        let mut obstacles = V1GroupHeaderObstacleIndex::new();
        for index in 0..self.resolved.len() {
            if index % 256 == 0 {
                crate::usc::cancel::check().map_err(LayoutError::new)?;
            }
            let Some(element) = self.resolved[index].as_ref() else {
                continue;
            };
            if !resolved_uses_v1_profile_group_header(element) {
                continue;
            }
            let header = v1_resolved_group_header_bounds(element);
            let mut y = header.y;
            for _ in 0..4 {
                let next = obstacles.avoid(header.x, header.width, y, header.height);
                if (next - y).abs() < 0.01 {
                    break;
                }
                y = next;
            }
            if let Some(frame) = enclosing_frames[index] {
                if let Some(floor) = metadata_floors[frame] {
                    y = y.max(floor);
                }
            }
            let delta = y - header.y;
            let element = self.resolved[index]
                .as_mut()
                .ok_or_else(|| LayoutError::new("group header geometry was not resolved"))?;
            element.text.y += delta;
            if element.icon_width > 0.0 && element.icon_height > 0.0 {
                element.icon_y += delta;
            }
            let header = v1_resolved_group_header_bounds(element);
            obstacles.insert(V1GroupHeaderObstacle {
                x: header.x,
                width: header.width,
                y: header.y,
                height: header.height,
                header: true,
            });
            let border_top = header.y + header.height / 2.0;
            let border_bottom = element.y + element.height;
            for border_y in [border_top, border_bottom] {
                obstacles.insert(V1GroupHeaderObstacle {
                    x: element.x,
                    width: element.width,
                    y: border_y,
                    height: 0.0,
                    header: false,
                });
            }
        }
        Ok(())
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
