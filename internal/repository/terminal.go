package repository

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"github.com/xaligo/xaligo/internal/entity"
)

type TerminalRepository interface {
	Render(entity.EngineResolvedDocument, entity.RenderOptions) ([]byte, error)
}

type terminalRepository struct{}

func NewTerminalRepository() TerminalRepository { return &terminalRepository{} }

type terminalCell struct {
	rune rune
	mask uint8
	fg   string
}

type terminalGlyphs struct {
	horizontal  rune
	vertical    rune
	topLeft     rune
	topRight    rune
	bottomLeft  rune
	bottomRight rune
	teeRight    rune
	teeLeft     rune
	teeDown     rune
	teeUp       rune
	cross       rune
	arrowRight  rune
	arrowLeft   rune
	arrowUp     rune
	arrowDown   rune
}

const (
	terminalNorth uint8 = 1 << iota
	terminalEast
	terminalSouth
	terminalWest
)

func (rcvr *terminalRepository) Render(document entity.EngineResolvedDocument, opts entity.RenderOptions) ([]byte, error) {
	if document.Width <= 0 || document.Height <= 0 {
		return nil, fmt.Errorf("terminal document size must be positive")
	}
	opts = normalizeTerminalOptions(opts)
	var output string
	switch opts.TerminalLayout {
	case entity.TerminalLayoutDiagram:
		output = rcvr.renderDiagram(document, opts, opts.TerminalHeight)
	case entity.TerminalLayoutSemantic:
		output = boundTerminalText(rcvr.renderSemantic(document, opts), opts.TerminalWidth, opts.TerminalHeight, opts.TerminalStyle)
	case entity.TerminalLayoutHybrid:
		diagramHeight := max(10, opts.TerminalHeight*2/3)
		detailsHeight := max(3, opts.TerminalHeight-diagramHeight-1)
		details := boundTerminalText(rcvr.renderDetails(document, opts), opts.TerminalWidth, detailsHeight, opts.TerminalStyle)
		output = rcvr.renderDiagram(document, opts, diagramHeight) + "\n" + details
	default:
		return nil, fmt.Errorf("unknown terminal layout %q", opts.TerminalLayout)
	}
	return []byte(strings.TrimRight(output, "\n") + "\n"), nil
}

func normalizeTerminalOptions(opts entity.RenderOptions) entity.RenderOptions {
	if opts.TerminalStyle == "" {
		opts.TerminalStyle = entity.TerminalStyleUnicode
	}
	if opts.TerminalLayout == "" {
		opts.TerminalLayout = entity.TerminalLayoutDiagram
	}
	if opts.TerminalDetail == "" {
		opts.TerminalDetail = entity.TerminalDetailNormal
	}
	if opts.TerminalColor == "" {
		opts.TerminalColor = entity.TerminalColorNever
	}
	if opts.TerminalIcons == "" {
		opts.TerminalIcons = entity.TerminalIconsLabel
	}
	if opts.TerminalWidth <= 0 {
		opts.TerminalWidth = 100
	}
	if opts.TerminalHeight <= 0 {
		opts.TerminalHeight = 40
	}
	opts.TerminalWidth = min(max(opts.TerminalWidth, 20), 500)
	opts.TerminalHeight = min(max(opts.TerminalHeight, 8), 200)
	return opts
}

func terminalGlyphSet(style entity.TerminalStyle) terminalGlyphs {
	if style == entity.TerminalStyleASCII {
		return terminalGlyphs{'-', '|', '+', '+', '+', '+', '+', '+', '+', '+', '+', '>', '<', '^', 'v'}
	}
	return terminalGlyphs{'─', '│', '┌', '┐', '└', '┘', '├', '┤', '┬', '┴', '┼', '▶', '◀', '▲', '▼'}
}

func (rcvr *terminalRepository) renderDiagram(document entity.EngineResolvedDocument, opts entity.RenderOptions, height int) string {
	width := opts.TerminalWidth
	grid := make([][]terminalCell, height)
	for row := range grid {
		grid[row] = make([]terminalCell, width)
		for col := range grid[row] {
			grid[row][col].rune = ' '
		}
	}
	mapPoint := func(x, y float64) (int, int) {
		return min(width-1, max(0, int(math.Round(x/document.Width*float64(width-1))))),
			min(height-1, max(0, int(math.Round(y/document.Height*float64(height-1)))))
	}
	glyphs := terminalGlyphSet(opts.TerminalStyle)
	for _, element := range document.Elements {
		if !element.Visual.Visible || element.Concept == entity.EngineConceptLine || element.Concept == entity.EngineConceptSpacer {
			continue
		}
		x1, y1 := mapPoint(element.X, element.Y)
		x2, y2 := mapPoint(element.X+element.Width, element.Y+element.Height)
		drawTerminalBox(grid, x1, y1, x2, y2, element.Visual.Stroke, glyphs)
	}
	for _, element := range document.Elements {
		if !element.Visual.Visible || element.Concept != entity.EngineConceptLine || len(element.Points) < 2 {
			continue
		}
		for index := 1; index < len(element.Points); index++ {
			x1, y1 := mapPoint(element.Points[index-1].X, element.Points[index-1].Y)
			x2, y2 := mapPoint(element.Points[index].X, element.Points[index].Y)
			drawTerminalSegment(grid, x1, y1, x2, y2, element.Visual.Stroke)
		}
		lastX, lastY := mapPoint(element.Points[len(element.Points)-1].X, element.Points[len(element.Points)-1].Y)
		prevX, prevY := mapPoint(element.Points[len(element.Points)-2].X, element.Points[len(element.Points)-2].Y)
		if element.Line.TargetDecoration != entity.EngineDecorationNone {
			grid[lastY][lastX].rune = terminalArrow(prevX, prevY, lastX, lastY, glyphs)
			grid[lastY][lastX].mask = 0
		}
	}
	for _, element := range document.Elements {
		if !element.Visual.Visible || element.Concept == entity.EngineConceptLine || element.Concept == entity.EngineConceptSpacer {
			continue
		}
		label := terminalElementLabel(element, opts)
		if label == "" {
			continue
		}
		x1, y1 := mapPoint(element.X, element.Y)
		x2, y2 := mapPoint(element.X+element.Width, element.Y+element.Height)
		row := min(height-1, max(0, (y1+y2)/2))
		if element.Concept == entity.EngineConceptFrame || element.Concept == entity.EngineConceptGroup || element.Concept == entity.EngineConceptCapture {
			row = min(height-1, y1+1)
		}
		available := max(0, x2-x1-1)
		label = terminalTruncate(label, available)
		writeTerminalText(grid, max(x1+1, (x1+x2-terminalDisplayWidth(label))/2), row, label, element.Text.Color)
	}
	return terminalGridString(grid, glyphs, opts.TerminalColor == entity.TerminalColorAlways)
}

func drawTerminalBox(grid [][]terminalCell, x1, y1, x2, y2 int, color string, glyphs terminalGlyphs) {
	if x2-x1 < 2 || y2-y1 < 2 {
		return
	}
	for x := x1 + 1; x < x2; x++ {
		setTerminalRune(grid, x, y1, glyphs.horizontal, color)
		setTerminalRune(grid, x, y2, glyphs.horizontal, color)
	}
	for y := y1 + 1; y < y2; y++ {
		setTerminalRune(grid, x1, y, glyphs.vertical, color)
		setTerminalRune(grid, x2, y, glyphs.vertical, color)
	}
	setTerminalRune(grid, x1, y1, glyphs.topLeft, color)
	setTerminalRune(grid, x2, y1, glyphs.topRight, color)
	setTerminalRune(grid, x1, y2, glyphs.bottomLeft, color)
	setTerminalRune(grid, x2, y2, glyphs.bottomRight, color)
}

func drawTerminalSegment(grid [][]terminalCell, x1, y1, x2, y2 int, color string) {
	x, y := x1, y1
	for x != x2 {
		next := x
		if x2 > x {
			next++
		} else {
			next--
		}
		connectTerminalCells(grid, x, y, next, y, color)
		x = next
	}
	for y != y2 {
		next := y
		if y2 > y {
			next++
		} else {
			next--
		}
		connectTerminalCells(grid, x, y, x, next, color)
		y = next
	}
}

func connectTerminalCells(grid [][]terminalCell, x1, y1, x2, y2 int, color string) {
	if x2 > x1 {
		grid[y1][x1].mask |= terminalEast
		grid[y2][x2].mask |= terminalWest
	}
	if x2 < x1 {
		grid[y1][x1].mask |= terminalWest
		grid[y2][x2].mask |= terminalEast
	}
	if y2 > y1 {
		grid[y1][x1].mask |= terminalSouth
		grid[y2][x2].mask |= terminalNorth
	}
	if y2 < y1 {
		grid[y1][x1].mask |= terminalNorth
		grid[y2][x2].mask |= terminalSouth
	}
	grid[y1][x1].fg, grid[y2][x2].fg = color, color
}

func terminalArrow(x1, y1, x2, y2 int, glyphs terminalGlyphs) rune {
	if abs(x2-x1) >= abs(y2-y1) {
		if x2 >= x1 {
			return glyphs.arrowRight
		}
		return glyphs.arrowLeft
	}
	if y2 >= y1 {
		return glyphs.arrowDown
	}
	return glyphs.arrowUp
}

func (rcvr *terminalRepository) renderSemantic(document entity.EngineResolvedDocument, opts entity.RenderOptions) string {
	children := make(map[string][]entity.EngineResolvedElement)
	for _, element := range document.Elements {
		if element.Concept != entity.EngineConceptLine && element.Concept != entity.EngineConceptSpacer {
			children[element.ParentID] = append(children[element.ParentID], element)
		}
	}
	var output strings.Builder
	roots := children[""]
	for index, root := range roots {
		rcvr.writeSemanticElement(&output, root, children, "", index == len(roots)-1, opts)
	}
	flows := terminalFlows(document, opts.TerminalStyle)
	if len(flows) > 0 {
		output.WriteString("\nFlows\n")
		for _, flow := range flows {
			fmt.Fprintf(&output, "  %s\n", flow)
		}
	}
	return output.String()
}

func (rcvr *terminalRepository) writeSemanticElement(output *strings.Builder, element entity.EngineResolvedElement, children map[string][]entity.EngineResolvedElement, prefix string, last bool, opts entity.RenderOptions) {
	branch, childPrefix := "├─ ", prefix+"│  "
	if opts.TerminalStyle == entity.TerminalStyleASCII {
		branch, childPrefix = "|-- ", prefix+"|   "
		if last {
			branch, childPrefix = "`-- ", prefix+"    "
		}
	} else if last {
		branch, childPrefix = "└─ ", prefix+"   "
	}
	if prefix == "" {
		branch = ""
	}
	label := terminalElementLabel(element, opts)
	if label == "" {
		label = element.ID
	}
	fmt.Fprintf(output, "%s%s%s", prefix, branch, label)
	if opts.TerminalDetail != entity.TerminalDetailCompact {
		fmt.Fprintf(output, " [%s]", element.Concept)
	}
	if opts.TerminalDetail == entity.TerminalDetailFull {
		fmt.Fprintf(output, " (%.0f,%.0f %.0fx%.0f)", element.X, element.Y, element.Width, element.Height)
	}
	output.WriteByte('\n')
	nested := children[element.ID]
	for index, child := range nested {
		rcvr.writeSemanticElement(output, child, children, childPrefix, index == len(nested)-1, opts)
	}
}

func (rcvr *terminalRepository) renderDetails(document entity.EngineResolvedDocument, opts entity.RenderOptions) string {
	var output strings.Builder
	focus := strings.TrimSpace(opts.TerminalFocus)
	if focus != "" {
		for _, element := range document.Elements {
			if element.ID == focus {
				fmt.Fprintf(&output, "Selected: %s\n  ID       %s\n  Concept  %s\n  Parent   %s\n  Bounds   %.0f,%.0f %.0fx%.0f\n",
					terminalElementLabel(element, opts), element.ID, element.Concept, terminalFallback(element.ParentID, "root"), element.X, element.Y, element.Width, element.Height)
				return output.String()
			}
		}
		fmt.Fprintf(&output, "Selected: %s (not found)\n", focus)
		return output.String()
	}
	output.WriteString("Components\n")
	count := 0
	for _, element := range document.Elements {
		if element.Concept == entity.EngineConceptLine || element.Concept == entity.EngineConceptSpacer {
			continue
		}
		fmt.Fprintf(&output, "  %-18s %s\n", terminalTruncate(element.ID, 18), terminalElementLabel(element, opts))
		count++
		if opts.TerminalDetail == entity.TerminalDetailCompact && count >= 8 {
			output.WriteString("  …\n")
			break
		}
	}
	flows := terminalFlows(document, opts.TerminalStyle)
	if len(flows) > 0 {
		output.WriteString("Flows\n")
		for _, flow := range flows {
			fmt.Fprintf(&output, "  %s\n", flow)
		}
	}
	return output.String()
}

func terminalFlows(document entity.EngineResolvedDocument, style entity.TerminalStyle) []string {
	var nodes []entity.EngineResolvedElement
	for _, element := range document.Elements {
		if element.Concept != entity.EngineConceptLine && element.Concept != entity.EngineConceptSpacer {
			nodes = append(nodes, element)
		}
	}
	flows := make([]string, 0)
	for _, line := range document.Elements {
		if line.Concept != entity.EngineConceptLine || len(line.Points) < 2 {
			continue
		}
		source := nearestTerminalElement(nodes, line.Points[0])
		target := nearestTerminalElement(nodes, line.Points[len(line.Points)-1])
		arrow := "──▶"
		if line.Line.TargetDecoration == entity.EngineDecorationNone {
			arrow = "───"
		}
		if style == entity.TerminalStyleASCII {
			arrow = "-->"
			if line.Line.TargetDecoration == entity.EngineDecorationNone {
				arrow = "---"
			}
		}
		label := ""
		if line.Line.Label != "" {
			label = " (" + line.Line.Label + ")"
		}
		flows = append(flows, fmt.Sprintf("%s %s %s%s", source, arrow, target, label))
	}
	return flows
}

func nearestTerminalElement(elements []entity.EngineResolvedElement, point entity.EnginePoint) string {
	best, distance, area := "?", math.MaxFloat64, math.MaxFloat64
	for _, element := range elements {
		dx := math.Max(math.Max(element.X-point.X, 0), point.X-(element.X+element.Width))
		dy := math.Max(math.Max(element.Y-point.Y, 0), point.Y-(element.Y+element.Height))
		d := math.Hypot(dx, dy)
		elementArea := element.Width * element.Height
		if d < distance-0.001 || math.Abs(d-distance) <= 0.001 && elementArea < area {
			best, distance, area = element.ID, d, elementArea
		}
	}
	return best
}

func terminalElementLabel(element entity.EngineResolvedElement, opts entity.RenderOptions) string {
	label := strings.TrimSpace(element.Text.Value)
	if label == "" && opts.TerminalIcons != entity.TerminalIconsNone {
		label = element.ID
	}
	if opts.TerminalStyle == entity.TerminalStyleASCII {
		label = terminalASCII(label)
	}
	if opts.TerminalIcons == entity.TerminalIconsSymbol && element.IconRef != "" {
		if opts.TerminalStyle == entity.TerminalStyleASCII {
			return "* " + label
		}
		return "◆ " + label
	}
	return label
}

func boundTerminalText(value string, width, height int, style entity.TerminalStyle) string {
	lines := strings.Split(strings.TrimRight(value, "\n"), "\n")
	truncated := len(lines) > height
	if truncated {
		lines = lines[:height]
	}
	for index := range lines {
		lines[index] = terminalTruncate(lines[index], width)
	}
	if truncated && len(lines) > 0 {
		marker := "…"
		if style == entity.TerminalStyleASCII {
			marker = "..."
		}
		lines[len(lines)-1] = marker
	}
	return strings.Join(lines, "\n")
}

func terminalASCII(value string) string {
	var output strings.Builder
	for _, r := range value {
		if r >= 0x20 && r <= 0x7e {
			output.WriteRune(r)
		} else {
			output.WriteByte('?')
		}
	}
	return output.String()
}

func terminalGridString(grid [][]terminalCell, glyphs terminalGlyphs, color bool) string {
	var output strings.Builder
	activeColor := ""
	for _, row := range grid {
		end := len(row)
		for end > 0 && row[end-1].rune == ' ' && row[end-1].mask == 0 {
			end--
		}
		for index := 0; index < end; index++ {
			cell := row[index]
			r := cell.rune
			if cell.mask != 0 {
				r = terminalMaskRune(cell.mask, glyphs)
			}
			if color && cell.fg != activeColor {
				if cell.fg == "" {
					output.WriteString("\x1b[0m")
				} else {
					output.WriteString(terminalANSI(cell.fg))
				}
				activeColor = cell.fg
			}
			output.WriteRune(r)
		}
		if color && activeColor != "" {
			output.WriteString("\x1b[0m")
			activeColor = ""
		}
		output.WriteByte('\n')
	}
	return strings.TrimRight(output.String(), "\n")
}

func terminalMaskRune(mask uint8, glyphs terminalGlyphs) rune {
	switch mask {
	case terminalEast | terminalWest:
		return glyphs.horizontal
	case terminalNorth | terminalSouth:
		return glyphs.vertical
	case terminalEast | terminalSouth:
		return glyphs.topLeft
	case terminalWest | terminalSouth:
		return glyphs.topRight
	case terminalEast | terminalNorth:
		return glyphs.bottomLeft
	case terminalWest | terminalNorth:
		return glyphs.bottomRight
	case terminalNorth | terminalEast | terminalSouth:
		return glyphs.teeRight
	case terminalNorth | terminalWest | terminalSouth:
		return glyphs.teeLeft
	case terminalEast | terminalSouth | terminalWest:
		return glyphs.teeDown
	case terminalEast | terminalNorth | terminalWest:
		return glyphs.teeUp
	default:
		return glyphs.cross
	}
}

func setTerminalRune(grid [][]terminalCell, x, y int, value rune, color string) {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return
	}
	grid[y][x].rune, grid[y][x].fg = value, color
}

func writeTerminalText(grid [][]terminalCell, x, y int, value, color string) {
	for _, r := range value {
		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
			break
		}
		grid[y][x] = terminalCell{rune: r, fg: color}
		width := terminalRuneWidth(r)
		for offset := 1; offset < width && x+offset < len(grid[y]); offset++ {
			grid[y][x+offset] = terminalCell{rune: ' ', fg: color}
		}
		x += width
	}
}

func terminalTruncate(value string, width int) string {
	if width <= 0 {
		return ""
	}
	if terminalDisplayWidth(value) <= width {
		return value
	}
	marker := "..."
	if width <= len(marker) {
		return marker[:width]
	}
	runes := []rune(value)
	for len(runes) > 0 && terminalDisplayWidth(string(runes))+len(marker) > width {
		runes = runes[:len(runes)-1]
	}
	return string(runes) + marker
}

func terminalDisplayWidth(value string) int {
	width := 0
	for _, r := range value {
		width += terminalRuneWidth(r)
	}
	return width
}

func terminalRuneWidth(r rune) int {
	if r == 0 || unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) || unicode.Is(unicode.Cf, r) {
		return 0
	}
	if r >= 0x1100 && (r <= 0x115f || r == 0x2329 || r == 0x232a ||
		(r >= 0x2e80 && r <= 0xa4cf && r != 0x303f) ||
		(r >= 0xac00 && r <= 0xd7a3) || (r >= 0xf900 && r <= 0xfaff) ||
		(r >= 0xfe10 && r <= 0xfe19) || (r >= 0xfe30 && r <= 0xfe6f) ||
		(r >= 0xff00 && r <= 0xff60) || (r >= 0xffe0 && r <= 0xffe6) ||
		(r >= 0x1f300 && r <= 0x1faff) || (r >= 0x20000 && r <= 0x3fffd)) {
		return 2
	}
	return 1
}

func terminalANSI(color string) string {
	value := strings.TrimPrefix(strings.TrimSpace(color), "#")
	if len(value) != 6 {
		return ""
	}
	var r, g, b int
	if _, err := fmt.Sscanf(value, "%02x%02x%02x", &r, &g, &b); err != nil {
		return ""
	}
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func terminalFallback(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}
func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
