package engine

import (
	"fmt"
	"io/fs"
	"path"
	"regexp"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

var createTablePatternV1EngineParseDatabaseImport = regexp.MustCompile(`(?is)CREATE\s+TABLE\s+(?:IF\s+NOT\s+EXISTS\s+)?([^\s(]+)\s*\((.*?)\)\s*;`)
var lineCommentPatternV1EngineParseDatabaseImport = regexp.MustCompile(`(?m)--[^\n]*$`)

func resolveDatabaseImportsV1EngineParseDatabaseImport(data *entity.Node, source *entity.ImportSource) error {
	if data == nil {
		return nil
	}
	for _, schema := range data.Children {
		if schema.Tag != "database-schema" || strings.TrimSpace(schema.Attr("src")) == "" {
			continue
		}
		if len(schema.Children) != 0 {
			return &entity.ParseError{Position: schema.Position, Err: fmt.Errorf("<database-schema src=%q> must not also contain inline entities", schema.Attr("src"))}
		}
		name := strings.TrimSpace(schema.Attr("src"))
		clean := path.Clean(strings.ReplaceAll(name, "\\", "/"))
		if clean == "." || !fs.ValidPath(clean) {
			return &entity.ParseError{Position: schema.Position, Err: fmt.Errorf("<database-schema src=%q> must be a relative file path", name)}
		}
		if source == nil || source.FS == nil {
			return &entity.ParseError{Position: schema.Position, Err: fmt.Errorf("<database-schema src=%q> requires an import filesystem", name)}
		}
		format := strings.ToLower(strings.TrimSpace(schema.Attr("format")))
		if format == "" {
			format = strings.TrimPrefix(strings.ToLower(path.Ext(clean)), ".")
		}
		if format != "sql" {
			return &entity.ParseError{Position: schema.Position, Err: fmt.Errorf("<database-schema format=%q> must be sql", format)}
		}
		dialect := strings.ToLower(strings.TrimSpace(schema.Attr("dialect")))
		if dialect == "postgres" {
			dialect = "postgresql"
		}
		if dialect != "" && dialect != "postgresql" && dialect != "mysql" && dialect != "sqlite" {
			return &entity.ParseError{Position: schema.Position, Err: fmt.Errorf("<database-schema dialect=%q> must be postgresql, mysql, or sqlite", dialect)}
		}
		contents, err := fs.ReadFile(source.FS, clean)
		if err != nil {
			return &entity.ParseError{Position: schema.Position, Err: fmt.Errorf("read database import %q: %w", name, err)}
		}
		entities, err := parseSQLSchemaV1EngineParseDatabaseImport(string(contents), schema.Position)
		if err != nil {
			return &entity.ParseError{Position: schema.Position, Err: fmt.Errorf("parse database import %q: %w", name, err)}
		}
		schema.Children = entities
	}
	return nil
}

func parseSQLSchemaV1EngineParseDatabaseImport(sql string, position entity.Position) ([]*entity.Node, error) {
	sql = lineCommentPatternV1EngineParseDatabaseImport.ReplaceAllString(sql, "")
	matches := createTablePatternV1EngineParseDatabaseImport.FindAllStringSubmatch(sql, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("SQL must contain at least one CREATE TABLE statement ending with ';'")
	}
	entities := make([]*entity.Node, 0, len(matches))
	for _, match := range matches {
		id := normalizeSQLIdentifierV1EngineParseDatabaseImport(match[1])
		entityNode := &entity.Node{Tag: "entity", Attrs: map[string]string{"id": id, "name": id}, Position: position}
		definitions := splitSQLDefinitionsV1EngineParseDatabaseImport(match[2])
		for _, definition := range definitions {
			upper := strings.ToUpper(strings.TrimSpace(definition))
			switch {
			case strings.HasPrefix(upper, "PRIMARY KEY") || strings.Contains(upper, " PRIMARY KEY") && strings.HasPrefix(upper, "CONSTRAINT "):
				columns, err := sqlColumnsInParensV1EngineParseDatabaseImport(definition, "PRIMARY KEY")
				if err != nil {
					return nil, err
				}
				entityNode.Children = append(entityNode.Children, &entity.Node{Tag: "primary-key", Attrs: map[string]string{"columns": strings.Join(columns, ",")}, Position: position})
			case strings.HasPrefix(upper, "FOREIGN KEY") || strings.Contains(upper, " FOREIGN KEY") && strings.HasPrefix(upper, "CONSTRAINT "):
				foreignKey, err := parseSQLForeignKeyV1EngineParseDatabaseImport(definition, position)
				if err != nil {
					return nil, err
				}
				entityNode.Children = append(entityNode.Children, foreignKey)
			case strings.HasPrefix(upper, "UNIQUE"), strings.HasPrefix(upper, "CHECK"), strings.HasPrefix(upper, "CONSTRAINT"):
				continue
			default:
				column, foreignKey, err := parseSQLColumnV1EngineParseDatabaseImport(definition, position)
				if err != nil {
					return nil, err
				}
				entityNode.Children = append(entityNode.Children, column)
				if foreignKey != nil {
					entityNode.Children = append(entityNode.Children, foreignKey)
				}
			}
		}
		entities = append(entities, entityNode)
	}
	return entities, nil
}

func splitSQLDefinitionsV1EngineParseDatabaseImport(body string) []string {
	var definitions []string
	start, depth := 0, 0
	quote := rune(0)
	for index, character := range body {
		if quote != 0 {
			if character == quote {
				quote = 0
			}
			continue
		}
		if character == '\'' || character == '"' || character == '`' {
			quote = character
		} else if character == '(' {
			depth++
		} else if character == ')' {
			depth--
		} else if character == ',' && depth == 0 {
			definitions = append(definitions, strings.TrimSpace(body[start:index]))
			start = index + 1
		}
	}
	if value := strings.TrimSpace(body[start:]); value != "" {
		definitions = append(definitions, value)
	}
	return definitions
}

func parseSQLColumnV1EngineParseDatabaseImport(definition string, position entity.Position) (*entity.Node, *entity.Node, error) {
	fields := strings.Fields(definition)
	if len(fields) < 2 {
		return nil, nil, fmt.Errorf("invalid SQL column definition %q", definition)
	}
	name := normalizeSQLIdentifierV1EngineParseDatabaseImport(fields[0])
	constraintIndex := len(fields)
	for index := 1; index < len(fields); index++ {
		switch strings.ToUpper(fields[index]) {
		case "PRIMARY", "NOT", "NULL", "UNIQUE", "DEFAULT", "REFERENCES", "CHECK", "COLLATE", "CONSTRAINT", "GENERATED":
			constraintIndex = index
			index = len(fields)
		}
	}
	columnType := strings.Join(fields[1:constraintIndex], " ")
	if columnType == "" {
		return nil, nil, fmt.Errorf("SQL column %q requires a type", name)
	}
	upper := strings.ToUpper(definition)
	attrs := map[string]string{"name": name, "type": columnType}
	if strings.Contains(upper, "PRIMARY KEY") {
		attrs["primary-key"] = "true"
	}
	if strings.Contains(upper, "NOT NULL") {
		attrs["nullable"] = "false"
	}
	if strings.Contains(upper, " UNIQUE") {
		attrs["unique"] = "true"
	}
	column := &entity.Node{Tag: "column", Attrs: attrs, Position: position}
	referencePattern := regexp.MustCompile(`(?i)REFERENCES\s+([^\s(]+)\s*\(([^)]+)\)`)
	match := referencePattern.FindStringSubmatch(definition)
	if len(match) == 0 {
		return column, nil, nil
	}
	foreignKey := &entity.Node{Tag: "foreign-key", Attrs: map[string]string{"columns": name, "references": normalizeSQLIdentifierV1EngineParseDatabaseImport(match[1]) + "." + normalizeSQLIdentifierV1EngineParseDatabaseImport(match[2])}, Position: position}
	applySQLReferentialActionsV1EngineParseDatabaseImport(foreignKey, definition)
	return column, foreignKey, nil
}

func parseSQLForeignKeyV1EngineParseDatabaseImport(definition string, position entity.Position) (*entity.Node, error) {
	local, err := sqlColumnsInParensV1EngineParseDatabaseImport(definition, "FOREIGN KEY")
	if err != nil {
		return nil, err
	}
	referencePattern := regexp.MustCompile(`(?i)REFERENCES\s+([^\s(]+)\s*\(([^)]+)\)`)
	match := referencePattern.FindStringSubmatch(definition)
	if len(match) == 0 {
		return nil, fmt.Errorf("foreign key requires REFERENCES table(columns)")
	}
	table := normalizeSQLIdentifierV1EngineParseDatabaseImport(match[1])
	referenced := splitSQLIdentifierListV1EngineParseDatabaseImport(match[2])
	qualified := make([]string, len(referenced))
	for index, column := range referenced {
		qualified[index] = table + "." + column
	}
	node := &entity.Node{Tag: "foreign-key", Attrs: map[string]string{"columns": strings.Join(local, ","), "references": strings.Join(qualified, ",")}, Position: position}
	applySQLReferentialActionsV1EngineParseDatabaseImport(node, definition)
	return node, nil
}

func sqlColumnsInParensV1EngineParseDatabaseImport(definition, keyword string) ([]string, error) {
	upper := strings.ToUpper(definition)
	index := strings.Index(upper, keyword)
	if index < 0 {
		return nil, fmt.Errorf("missing %s", keyword)
	}
	open := strings.Index(definition[index+len(keyword):], "(")
	if open < 0 {
		return nil, fmt.Errorf("%s requires a column list", keyword)
	}
	open += index + len(keyword)
	close := strings.Index(definition[open+1:], ")")
	if close < 0 {
		return nil, fmt.Errorf("%s has an unterminated column list", keyword)
	}
	return splitSQLIdentifierListV1EngineParseDatabaseImport(definition[open+1 : open+1+close]), nil
}

func splitSQLIdentifierListV1EngineParseDatabaseImport(value string) []string {
	parts := strings.Split(value, ",")
	columns := make([]string, 0, len(parts))
	for _, part := range parts {
		columns = append(columns, normalizeSQLIdentifierV1EngineParseDatabaseImport(part))
	}
	return columns
}

func normalizeSQLIdentifierV1EngineParseDatabaseImport(value string) string {
	value = strings.TrimSpace(value)
	if index := strings.LastIndex(value, "."); index >= 0 {
		value = value[index+1:]
	}
	return strings.Trim(value, "`\"[] ")
}

func applySQLReferentialActionsV1EngineParseDatabaseImport(node *entity.Node, definition string) {
	upper := strings.ToUpper(definition)
	for _, action := range []struct{ sql, attr string }{{"ON DELETE", "on-delete"}, {"ON UPDATE", "on-update"}} {
		index := strings.Index(upper, action.sql)
		if index < 0 {
			continue
		}
		value := strings.TrimSpace(definition[index+len(action.sql):])
		fields := strings.Fields(value)
		if len(fields) > 0 {
			resolved := strings.ToLower(fields[0])
			if resolved == "set" || resolved == "no" {
				if len(fields) > 1 {
					resolved += "-" + strings.ToLower(fields[1])
				}
			}
			node.Attrs[action.attr] = resolved
		}
	}
}
