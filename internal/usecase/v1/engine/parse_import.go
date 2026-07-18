package engine

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"go.yaml.in/yaml/v3"
)

func resolveTableImportsV1EngineParseImport(root, data *entity.Node, source *entity.ImportSource) error {
	definitions := map[string]*entity.Node{}
	if data != nil {
		for _, child := range data.Children {
			if child.Tag != "table-data" {
				continue
			}
			id := strings.TrimSpace(child.Attrs["id"])
			if id == "" {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<table-data> requires a non-empty id")}
			}
			if _, exists := definitions[id]; exists {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate <table-data id=%q>", id)}
			}
			if strings.TrimSpace(child.Attrs["src"]) != "" {
				if err := loadTableImportV1EngineParseImport(child, source); err != nil {
					return err
				}
			}
			definitions[id] = child
		}
	}
	return expandTableImportsV1EngineParseImport(root, definitions, source)
}

func expandTableImportsV1EngineParseImport(node *entity.Node, definitions map[string]*entity.Node, source *entity.ImportSource) error {
	if node == nil {
		return nil
	}
	if node.Tag == "table" {
		if reference := strings.TrimSpace(node.Attrs["data"]); reference != "" {
			definition, exists := definitions[reference]
			if !exists {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<table data=%q> references unknown table data", reference)}
			}
			if len(node.Children) > 0 || strings.TrimSpace(node.Text) != "" {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<table data=%q> may not also contain inline rows", reference)}
			}
			node.Children = cloneNodesV1EngineParseImport(definition.Children)
		}
		if strings.TrimSpace(node.Attrs["src"]) != "" {
			if len(node.Children) > 0 || strings.TrimSpace(node.Text) != "" {
				return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<table src=%q> may not also contain inline rows", node.Attrs["src"])}
			}
			if err := loadTableImportV1EngineParseImport(node, source); err != nil {
				return err
			}
		}
	}
	for _, child := range node.Children {
		if err := expandTableImportsV1EngineParseImport(child, definitions, source); err != nil {
			return err
		}
	}
	return nil
}

func loadTableImportV1EngineParseImport(node *entity.Node, source *entity.ImportSource) error {
	name := strings.TrimSpace(node.Attrs["src"])
	if name == "" {
		return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<%s> requires src", node.Tag)}
	}
	clean := path.Clean(strings.ReplaceAll(name, "\\", "/"))
	if clean == "." || !fs.ValidPath(clean) {
		return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<%s src=%q> must be a relative file path", node.Tag, name)}
	}
	if source == nil || source.FS == nil {
		return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("<%s src=%q> requires an import filesystem", node.Tag, name)}
	}
	contents, err := fs.ReadFile(source.FS, clean)
	if err != nil {
		return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("read table import %q: %w", name, err)}
	}
	format := strings.ToLower(strings.TrimSpace(node.Attrs["format"]))
	if format == "" {
		format = strings.TrimPrefix(strings.ToLower(path.Ext(clean)), ".")
	}
	columns, rows, err := decodeTableImportV1EngineParseImport(contents, format)
	if err != nil {
		return &entity.ParseError{Position: node.Position, Err: fmt.Errorf("decode table import %q: %w", name, err)}
	}
	node.Children = tableNodesV1EngineParseImport(columns, rows, node.Position)
	return nil
}

func decodeTableImportV1EngineParseImport(contents []byte, format string) ([]string, [][]string, error) {
	switch format {
	case "csv", "tsv":
		reader := csv.NewReader(strings.NewReader(string(contents)))
		if format == "tsv" {
			reader.Comma = '\t'
		}
		records, err := reader.ReadAll()
		if err != nil {
			return nil, nil, err
		}
		if len(records) == 0 || len(records[0]) == 0 {
			return nil, nil, fmt.Errorf("file must contain a header row")
		}
		return records[0], records[1:], nil
	case "json", "yaml", "yml":
		var records []map[string]any
		if format == "json" {
			if err := json.Unmarshal(contents, &records); err != nil {
				return nil, nil, err
			}
		} else if err := yaml.Unmarshal(contents, &records); err != nil {
			return nil, nil, err
		}
		if len(records) == 0 {
			return nil, nil, fmt.Errorf("file must contain at least one object")
		}
		columnSet := map[string]struct{}{}
		for _, record := range records {
			for key := range record {
				columnSet[key] = struct{}{}
			}
		}
		columns := make([]string, 0, len(columnSet))
		for key := range columnSet {
			columns = append(columns, key)
		}
		sort.Strings(columns)
		rows := make([][]string, 0, len(records))
		for _, record := range records {
			row := make([]string, len(columns))
			for index, column := range columns {
				if value, exists := record[column]; exists && value != nil {
					row[index] = fmt.Sprint(value)
				}
			}
			rows = append(rows, row)
		}
		return columns, rows, nil
	default:
		return nil, nil, fmt.Errorf("unsupported format %q; use csv, tsv, json, yaml, or yml", format)
	}
}

func tableNodesV1EngineParseImport(columns []string, rows [][]string, position entity.Position) []*entity.Node {
	nodes := []*entity.Node{{Tag: "header", Attrs: map[string]string{}, Position: position}}
	for _, column := range columns {
		nodes[0].Children = append(nodes[0].Children, &entity.Node{Tag: "cell", Attrs: map[string]string{}, Text: column, Position: position})
	}
	for _, values := range rows {
		row := &entity.Node{Tag: "row", Attrs: map[string]string{}, Position: position}
		for _, value := range values {
			row.Children = append(row.Children, &entity.Node{Tag: "cell", Attrs: map[string]string{}, Text: value, Position: position})
		}
		nodes = append(nodes, row)
	}
	return nodes
}

func cloneNodesV1EngineParseImport(nodes []*entity.Node) []*entity.Node {
	cloned := make([]*entity.Node, 0, len(nodes))
	for _, node := range nodes {
		copy := &entity.Node{Tag: node.Tag, Attrs: cloneAttrsV1EngineParseTable(node.Attrs), Text: node.Text, TextRuns: append([]entity.TextRun(nil), node.TextRuns...), Position: node.Position}
		copy.Children = cloneNodesV1EngineParseImport(node.Children)
		cloned = append(cloned, copy)
	}
	return cloned
}
