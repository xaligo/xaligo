package engine

import (
	"fmt"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

func normalizeDatabasesV1EngineParseDatabase(root, data *entity.Node) error {
	if root == nil {
		return nil
	}
	schemas := map[string]*entity.Node{}
	if data != nil {
		for _, child := range data.Children {
			if child.Tag != "database-schema" {
				continue
			}
			id := strings.TrimSpace(child.Attr("id"))
			if id == "" {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<database-schema> requires a non-empty id")}
			}
			if _, exists := schemas[id]; exists {
				return &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate database schema id %q", id)}
			}
			schemas[id] = child
		}
	}

	var walk func(*entity.Node) error
	walk = func(node *entity.Node) error {
		if node.Tag == "frame" {
			return normalizeFrameDatabasesV1EngineParseDatabase(node, schemas)
		}
		for _, child := range node.Children {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(root)
}

func normalizeFrameDatabasesV1EngineParseDatabase(frame *entity.Node, schemas map[string]*entity.Node) error {
	var generatedConnections []*entity.Node
	var walk func(*entity.Node) error
	walk = func(node *entity.Node) error {
		if node.Tag == "database" {
			connections, err := normalizeDatabaseV1EngineParseDatabase(node, schemas)
			if err != nil {
				return err
			}
			generatedConnections = append(generatedConnections, connections...)
		}
		for _, child := range node.Children {
			if err := walk(child); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(frame); err != nil {
		return err
	}
	frame.Children = append(frame.Children, generatedConnections...)
	return nil
}

func normalizeDatabaseV1EngineParseDatabase(database *entity.Node, schemas map[string]*entity.Node) ([]*entity.Node, error) {
	if ref := strings.TrimSpace(database.Attr("data")); ref != "" {
		schema, exists := schemas[ref]
		if !exists {
			return nil, &entity.ParseError{Position: database.Position, Err: fmt.Errorf("<database data=%q> does not match a <database-schema id>", ref)}
		}
		if len(database.Children) != 0 {
			return nil, &entity.ParseError{Position: database.Position, Err: fmt.Errorf("<database data=%q> must not also contain inline entities", ref)}
		}
		for _, child := range schema.Children {
			database.Children = append(database.Children, cloneNodeV1EngineParseDatabase(child))
		}
	}
	if _, exists := database.Attrs["layout"]; !exists {
		database.Attrs["layout"] = "horizontal"
	}
	entities := map[string]*entity.Node{}
	columns := map[string]map[string]bool{}
	for _, child := range database.Children {
		if child.Tag != "entity" {
			return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<database> may only contain <entity> children, got <%s>", child.Tag)}
		}
		id := strings.TrimSpace(child.Attr("id"))
		if id == "" {
			return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("<entity> requires a non-empty id")}
		}
		if _, exists := entities[id]; exists {
			return nil, &entity.ParseError{Position: child.Position, Err: fmt.Errorf("duplicate entity id %q", id)}
		}
		entities[id] = child
		columns[id] = map[string]bool{}
		for _, definition := range child.Children {
			if definition.Tag == "column" {
				name := strings.TrimSpace(definition.Attr("name"))
				if name == "" || columns[id][name] {
					return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("<entity id=%q> column names must be non-empty and unique", id)}
				}
				columns[id][name] = true
			}
		}
	}
	for id, entityNode := range entities {
		var primaryKey *entity.Node
		for _, definition := range entityNode.Children {
			if definition.Tag != "primary-key" {
				continue
			}
			if primaryKey != nil {
				return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("<entity id=%q> may define only one <primary-key>", id)}
			}
			primaryKey = definition
			for _, column := range splitDatabaseColumnsV1EngineParseDatabase(definition.Attr("columns")) {
				if !columns[id][column] {
					return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("primary key %s references missing column %s", id, column)}
				}
				for _, candidate := range entityNode.Children {
					if candidate.Tag == "column" && candidate.Attr("name") == column {
						candidate.Attrs["primary-key"] = "true"
					}
				}
			}
			if len(splitDatabaseColumnsV1EngineParseDatabase(definition.Attr("columns"))) == 0 {
				return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("<primary-key> requires one or more columns")}
			}
		}
	}

	var connections []*entity.Node
	for id, entityNode := range entities {
		var rows []*entity.Node
		rows = append(rows, tableRowNodeV1EngineParseTable("table-header", []string{"Column", "Type", "Constraints"}, []string{"left", "left", "left"}, entityNode.Position))
		for _, definition := range entityNode.Children {
			switch definition.Tag {
			case "column":
				constraints := databaseColumnConstraintsV1EngineParseDatabase(definition)
				rows = append(rows, tableRowNodeV1EngineParseTable("table-row", []string{definition.Attr("name"), definition.Attr("type"), constraints}, []string{"left", "left", "left"}, definition.Position))
			case "foreign-key":
				fromColumns := splitDatabaseColumnsV1EngineParseDatabase(definition.Attr("columns"))
				toEntity, toColumns, err := parseDatabaseReferenceV1EngineParseDatabase(definition.Attr("references"))
				if err != nil || len(fromColumns) == 0 || len(fromColumns) != len(toColumns) {
					return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("<foreign-key> requires equally sized columns and entity.column references")}
				}
				for index, fromColumn := range fromColumns {
					if !columns[id][fromColumn] || !columns[toEntity][toColumns[index]] {
						return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("foreign key %s.%s references missing column %s.%s", id, fromColumn, toEntity, toColumns[index])}
					}
				}
				attrs := map[string]string{"src": id, "dst": toEntity, "_xaligoDatabaseForeignKey": strings.Join(fromColumns, ","), "_xaligoDatabaseReferences": toEntity + "." + strings.Join(toColumns, ",")}
				for _, action := range []string{"on-delete", "on-update"} {
					if value := strings.TrimSpace(definition.Attr(action)); value != "" {
						value = strings.ToLower(value)
						switch value {
						case "cascade", "restrict", "no-action", "set-null", "set-default":
						default:
							return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("<foreign-key %s=%q> must be cascade, restrict, no-action, set-null, or set-default", action, value)}
						}
						key := "_xaligoDatabaseOnDelete"
						if action == "on-update" {
							key = "_xaligoDatabaseOnUpdate"
						}
						attrs[key] = value
					}
				}
				connections = append(connections, &entity.Node{Tag: "connection", Attrs: attrs, Position: definition.Position})
			case "primary-key":
				continue
			default:
				return nil, &entity.ParseError{Position: definition.Position, Err: fmt.Errorf("<entity> may only contain <column>, <primary-key>, and <foreign-key> children, got <%s>", definition.Tag)}
			}
		}
		entityNode.Children = rows
		entityNode.TextRuns = nil
		entityNode.Text = ""
		entityNode.Attrs["title"] = firstNonEmptyV1EngineParseDatabase(entityNode.Attr("name"), id)
		delete(entityNode.Attrs, "name")
		entityNode.Attrs["gap"] = "0"
	}
	return connections, nil
}

func splitDatabaseColumnsV1EngineParseDatabase(value string) []string {
	var columns []string
	for _, column := range strings.Split(value, ",") {
		if column = strings.TrimSpace(column); column != "" {
			columns = append(columns, column)
		}
	}
	return columns
}

func parseDatabaseReferenceV1EngineParseDatabase(value string) (string, []string, error) {
	parts := splitDatabaseColumnsV1EngineParseDatabase(value)
	entityID := ""
	columns := make([]string, 0, len(parts))
	for _, part := range parts {
		qualified := strings.Split(part, ".")
		if len(qualified) != 2 || qualified[0] == "" || qualified[1] == "" {
			return "", nil, fmt.Errorf("invalid reference %q", part)
		}
		if entityID == "" {
			entityID = qualified[0]
		} else if entityID != qualified[0] {
			return "", nil, fmt.Errorf("foreign key references multiple entities")
		}
		columns = append(columns, qualified[1])
	}
	return entityID, columns, nil
}

func databaseColumnConstraintsV1EngineParseDatabase(column *entity.Node) string {
	var values []string
	if column.Attr("primary-key") == "true" {
		values = append(values, "PK")
	}
	if column.Attr("nullable") == "false" {
		values = append(values, "NOT NULL")
	}
	if column.Attr("unique") == "true" {
		values = append(values, "UNIQUE")
	}
	if value := strings.TrimSpace(column.Attr("default")); value != "" {
		values = append(values, "DEFAULT "+value)
	}
	return strings.Join(values, ", ")
}

func cloneNodeV1EngineParseDatabase(source *entity.Node) *entity.Node {
	clone := &entity.Node{Tag: source.Tag, Attrs: cloneAttrsV1EngineParseTable(source.Attrs), Text: source.Text, Position: source.Position, TextRuns: append([]entity.TextRun(nil), source.TextRuns...)}
	for _, child := range source.Children {
		clone.Children = append(clone.Children, cloneNodeV1EngineParseDatabase(child))
	}
	return clone
}

func firstNonEmptyV1EngineParseDatabase(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}
