package engine_test

import (
	"strings"
	"testing"
	"testing/fstest"

	"github.com/xaligo/xaligo/internal/entity"
	engine "github.com/xaligo/xaligo/internal/usecase/v1/engine"
)

func TestParseWithImportsLoadsReusableCSVTableV1EngineParseImport(t *testing.T) {
	input := `<xaligo version="1"><data><table-data id="services" src="data/services.csv" /></data><frames><frame id="main"><table data="services" /></frame></frames></xaligo>`
	doc, err := engine.ParseWithImportsV1EngineParseDocument(strings.NewReader(input), &entity.ImportSource{FS: fstest.MapFS{
		"data/services.csv": {Data: []byte("name,port\nAPI,8080\nDB,5432\n")},
	}})
	if err != nil {
		t.Fatalf("ParseWithImports() error = %v", err)
	}
	table := doc.Root.Children[0].Children[0]
	if len(table.Children) != 3 || table.Children[0].Tag != "table-header" || table.Children[2].Children[1].Text != "5432" {
		t.Fatalf("normalized imported table = %#v", table)
	}
}

func TestParseWithImportsLoadsDirectJSONTableV1EngineParseImport(t *testing.T) {
	input := `<xaligo version="1"><frames><frame id="main"><table src="rows.json" /></frame></frames></xaligo>`
	doc, err := engine.ParseWithImportsV1EngineParseDocument(strings.NewReader(input), &entity.ImportSource{FS: fstest.MapFS{
		"rows.json": {Data: []byte(`[{"name":"API","port":8080}]`)},
	}})
	if err != nil {
		t.Fatalf("ParseWithImports() error = %v", err)
	}
	table := doc.Root.Children[0].Children[0]
	if table.Children[0].Children[0].Text != "name" || table.Children[1].Children[1].Text != "8080" {
		t.Fatalf("JSON table = %#v", table)
	}
}

func TestParseWithImportsRejectsUnsafeAndMissingSourcesV1EngineParseImport(t *testing.T) {
	for _, source := range []string{"../secret.csv", "/tmp/secret.csv"} {
		input := `<xaligo version="1"><frames><frame id="main"><table src="` + source + `" /></frame></frames></xaligo>`
		if _, err := engine.ParseWithImportsV1EngineParseDocument(strings.NewReader(input), &entity.ImportSource{FS: fstest.MapFS{}}); err == nil || !strings.Contains(err.Error(), "relative file path") {
			t.Fatalf("src %q error = %v, want relative path error", source, err)
		}
	}
	input := `<xaligo version="1"><frames><frame id="main"><table src="rows.csv" /></frame></frames></xaligo>`
	if _, err := engine.ParseV1EngineParseDocument(strings.NewReader(input)); err == nil || !strings.Contains(err.Error(), "requires an import filesystem") {
		t.Fatalf("missing filesystem error = %v", err)
	}
}

func TestParseWithImportsLoadsCompositeSQLSchemaV1EngineParseImport(t *testing.T) {
	input := `<xaligo version="1"><data><database-schema id="app" src="schema.sql" dialect="postgresql" /></data><frames><frame id="erd"><database data="app" /></frame></frames></xaligo>`
	sql := `CREATE TABLE roles (
  tenant_id bigint NOT NULL,
  id bigint NOT NULL,
  PRIMARY KEY (tenant_id, id)
);
CREATE TABLE users (
  id bigint PRIMARY KEY,
  tenant_id bigint NOT NULL,
  role_id bigint NOT NULL,
  CONSTRAINT fk_role FOREIGN KEY (tenant_id, role_id)
    REFERENCES roles (tenant_id, id) ON DELETE CASCADE
);`
	doc, err := engine.ParseWithImportsV1EngineParseDocument(strings.NewReader(input), &entity.ImportSource{FS: fstest.MapFS{
		"schema.sql": {Data: []byte(sql)},
	}})
	if err != nil {
		t.Fatalf("ParseWithImports() error = %v", err)
	}
	database := doc.Root.Children[0].Children[0]
	if len(database.Children) != 2 || database.Children[0].Attr("id") != "roles" || database.Children[1].Attr("id") != "users" {
		t.Fatalf("imported database = %#v", database)
	}
	connection := doc.Root.Children[0].Children[len(doc.Root.Children[0].Children)-1]
	if connection.Attr("_xaligoDatabaseForeignKey") != "tenant_id,role_id" || connection.Attr("_xaligoDatabaseOnDelete") != "cascade" {
		t.Fatalf("imported SQL relation = %#v", connection.Attrs)
	}
}

func TestParseWithImportsRejectsUnsupportedSQLDialectV1EngineParseImport(t *testing.T) {
	input := `<xaligo version="1"><data><database-schema id="app" src="schema.sql" dialect="oracle" /></data><frames><frame id="erd"><database data="app" /></frame></frames></xaligo>`
	_, err := engine.ParseWithImportsV1EngineParseDocument(strings.NewReader(input), &entity.ImportSource{FS: fstest.MapFS{"schema.sql": {Data: []byte("CREATE TABLE t (id bigint);")}}})
	if err == nil || !strings.Contains(err.Error(), "postgresql, mysql, or sqlite") {
		t.Fatalf("ParseWithImports() error = %v, want dialect error", err)
	}
}
