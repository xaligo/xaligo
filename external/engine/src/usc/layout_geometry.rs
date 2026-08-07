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

fn rectangles_overlap(left: Bounds, right: Bounds) -> bool {
    left.x <= right.x + right.width
        && left.x + left.width >= right.x
        && left.y <= right.y + right.height
        && left.y + left.height >= right.y
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

fn format_number(value: f64) -> String {
    let rendered = format!("{value:.6}");
    rendered
        .trim_end_matches('0')
        .trim_end_matches('.')
        .to_owned()
}
