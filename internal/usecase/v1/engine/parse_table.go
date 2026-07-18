package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

var tableStyleAttributesV1EngineParseTable = []string{"color", "background-color", "border-color", "font-family", "font-size"}

func normalizeTablesV1EngineParseTable(root *entity.Node) error {
	if root == nil {
		return nil
	}
	if root.Tag == "table" {
		return normalizeTableV1EngineParseTable(root)
	}
	for _, child := range root.Children {
		if err := normalizeTablesV1EngineParseTable(child); err != nil {
			return err
		}
	}
	return nil
}

func normalizeTableV1EngineParseTable(table *entity.Node) error {
	if err := normalizeTableStyleV1EngineParseTable(table); err != nil {
		return &entity.ParseError{Position: table.Position, Err: err}
	}
	if _, exists := table.Attrs["gap"]; !exists {
		table.Attrs["gap"] = "0"
	}
	pipeRows, pipeAlignments, err := parsePipeTableV1EngineParseTable(table)
	if err != nil {
		return &entity.ParseError{Position: table.Position, Err: err}
	}

	var normalized []*entity.Node
	columnCount := 0
	hasHeader := false
	if len(pipeRows) > 0 {
		columnCount = len(pipeRows[0])
		hasHeader = true
		normalized = append(normalized, tableRowNodeV1EngineParseTable("table-header", pipeRows[0], pipeAlignments, table.Position))
		for _, cells := range pipeRows[1:] {
			if len(cells) != columnCount {
				return &entity.ParseError{Position: table.Position, Err: fmt.Errorf("<table> pipe row has %d cells, want %d", len(cells), columnCount)}
			}
			normalized = append(normalized, tableRowNodeV1EngineParseTable("table-row", cells, pipeAlignments, table.Position))
		}
	}

	for _, child := range table.Children {
		if child.Tag != "header" && child.Tag != "row" {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<table> may only contain pipe rows, <header>, and <row> children, got <%s>", child.Tag)}
		}
		if child.Tag == "header" && hasHeader {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<table> may define only one header across pipe and tag syntax")}
		}
		if child.Tag == "header" {
			hasHeader = true
		}
		row := &entity.Node{Tag: "table-row", Attrs: cloneAttrsV1EngineParseTable(child.Attrs), Position: child.Position}
		if _, exists := row.Attrs["gap"]; !exists {
			row.Attrs["gap"] = "0"
		}
		if child.Tag == "header" {
			row.Tag = "table-header"
			inheritTableHeaderStyleV1EngineParseTable(row.Attrs, table.Attrs)
		}
		inheritTableStyleV1EngineParseTable(row.Attrs, table.Attrs)
		if err := normalizeTableStyleV1EngineParseTable(row); err != nil {
			return &entity.ParseError{Position: child.Position, Err: err}
		}
		for _, cell := range child.Children {
			if cell.Tag != "cell" {
				return &entity.ParseError{Position: cell.Position, Err: fmt.Errorf("<%s> may only contain <cell> children, got <%s>", child.Tag, cell.Tag)}
			}
			cell.Tag = "table-cell"
			inheritTableStyleV1EngineParseTable(cell.Attrs, row.Attrs)
			if err := normalizeTableStyleV1EngineParseTable(cell); err != nil {
				return &entity.ParseError{Position: cell.Position, Err: err}
			}
			if align, exists := cell.Attrs["align"]; exists {
				switch strings.ToLower(strings.TrimSpace(align)) {
				case "left", "center", "right":
					cell.Attrs["align"] = "middle-" + strings.ToLower(strings.TrimSpace(align))
				case "top-left", "top-center", "top-right", "middle-left", "middle-center", "middle-right", "bottom-left", "bottom-center", "bottom-right":
				default:
					return &entity.ParseError{Position: cell.Position, Err: fmt.Errorf("<cell align=%q> must be left, center, right, or a vertical-horizontal alignment", align)}
				}
			}
			if row.Tag == "table-header" {
				cell.Attrs["_xaligoTableHeader"] = "true"
			}
			row.Children = append(row.Children, cell)
		}
		if len(row.Children) == 0 {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<%s> must contain at least one <cell>", child.Tag)}
		}
		if columnCount == 0 {
			columnCount = len(row.Children)
		} else if len(row.Children) != columnCount {
			return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<table> row has %d cells, want %d", len(row.Children), columnCount)}
		}
		normalized = append(normalized, row)
	}

	if len(normalized) == 0 {
		return &entity.ParseError{Position: table.Position, Err: fmt.Errorf("<table> must contain a pipe table or tagged rows")}
	}
	table.Children = normalized
	for _, row := range table.Children {
		if row.Tag == "table-header" {
			inheritTableHeaderStyleV1EngineParseTable(row.Attrs, table.Attrs)
		}
		inheritTableStyleV1EngineParseTable(row.Attrs, table.Attrs)
		for _, cell := range row.Children {
			inheritTableStyleV1EngineParseTable(cell.Attrs, row.Attrs)
		}
	}
	table.Text = ""
	table.TextRuns = nil
	return nil
}

func normalizeTableStyleV1EngineParseTable(node *entity.Node) error {
	if err := normalizePresentationStyleV1EngineParsePresentationStyle(node, "color", "background-color", "border-color"); err != nil {
		return err
	}
	if node.Tag == "table" {
		header := &entity.Node{Tag: "table header style", Attrs: map[string]string{}}
		for _, name := range tableStyleAttributesV1EngineParseTable {
			if value, exists := node.Attrs["header-"+name]; exists {
				header.Attrs[name] = value
			}
		}
		if err := normalizeTableStyleV1EngineParseTable(header); err != nil {
			return err
		}
		for name, value := range header.Attrs {
			node.Attrs["header-"+name] = value
		}
	}
	return nil
}

func inheritTableStyleV1EngineParseTable(target, source map[string]string) {
	for _, name := range tableStyleAttributesV1EngineParseTable {
		if _, exists := target[name]; !exists {
			if value, available := source[name]; available {
				target[name] = value
			}
		}
	}
}

func inheritTableHeaderStyleV1EngineParseTable(target, table map[string]string) {
	for _, name := range tableStyleAttributesV1EngineParseTable {
		if _, exists := target[name]; !exists {
			if value, available := table["header-"+name]; available {
				target[name] = value
			}
		}
	}
}

func parsePipeTableV1EngineParseTable(table *entity.Node) ([][]string, []string, error) {
	var source strings.Builder
	for _, run := range table.TextRuns {
		source.WriteString(run.Text)
	}
	lines := strings.Split(source.String(), "\n")
	var raw [][]string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, "|") || !strings.HasSuffix(line, "|") {
			return nil, nil, fmt.Errorf("<table> text must use pipe rows beginning and ending with '|'")
		}
		cells, err := splitPipeRowV1EngineParseTable(line)
		if err != nil {
			return nil, nil, err
		}
		raw = append(raw, cells)
	}
	if len(raw) == 0 {
		return nil, nil, nil
	}
	if len(raw) < 2 || len(raw[0]) != len(raw[1]) {
		return nil, nil, fmt.Errorf("<table> pipe syntax requires a header and matching separator row")
	}
	alignments := make([]string, len(raw[1]))
	for i, separator := range raw[1] {
		value := strings.TrimSpace(separator)
		left, right := strings.HasPrefix(value, ":"), strings.HasSuffix(value, ":")
		core := strings.Trim(value, ":")
		if len(core) < 3 || strings.Trim(core, "-") != "" {
			return nil, nil, fmt.Errorf("<table> separator cell %q must contain at least three hyphens", separator)
		}
		switch {
		case left && right:
			alignments[i] = "center"
		case right:
			alignments[i] = "right"
		default:
			alignments[i] = "left"
		}
	}
	return append([][]string{raw[0]}, raw[2:]...), alignments, nil
}

func splitPipeRowV1EngineParseTable(line string) ([]string, error) {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(strings.TrimSuffix(line, "|"), "|")
	var cells []string
	var cell strings.Builder
	escaped := false
	for _, character := range line {
		if escaped {
			if character != '|' && character != '\\' {
				cell.WriteRune('\\')
			}
			cell.WriteRune(character)
			escaped = false
			continue
		}
		if character == '\\' {
			escaped = true
			continue
		}
		if character == '|' {
			cells = append(cells, strings.TrimSpace(cell.String()))
			cell.Reset()
			continue
		}
		cell.WriteRune(character)
	}
	if escaped {
		return nil, fmt.Errorf("<table> pipe row ends with an incomplete escape")
	}
	cells = append(cells, strings.TrimSpace(cell.String()))
	return cells, nil
}

func tableRowNodeV1EngineParseTable(tag string, values, alignments []string, position entity.Position) *entity.Node {
	row := &entity.Node{Tag: tag, Attrs: map[string]string{"gap": "0"}, Position: position}
	for i, value := range values {
		attrs := map[string]string{}
		if tag == "table-header" {
			attrs["_xaligoTableHeader"] = "true"
		}
		if i < len(alignments) && alignments[i] != "" {
			attrs["align"] = "middle-" + alignments[i]
		}
		row.Children = append(row.Children, &entity.Node{Tag: "table-cell", Attrs: attrs, Text: value, Position: position})
	}
	return row
}

func cloneAttrsV1EngineParseTable(source map[string]string) map[string]string {
	cloned := make(map[string]string, len(source))
	for key, value := range source {
		cloned[key] = value
	}
	return cloned
}
