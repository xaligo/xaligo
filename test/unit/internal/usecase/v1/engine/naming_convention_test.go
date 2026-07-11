package engine_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"unicode"
)

func TestV1EnginePackageScopeNamesIncludeFileSuffix(t *testing.T) {
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve test file path")
	}
	repositoryRoot := filepath.Clean(filepath.Join(filepath.Dir(testFile), "../../../../../.."))
	engineDir := filepath.Join(repositoryRoot, "internal", "usecase", "v1", "engine")
	entries, err := os.ReadDir(engineDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			path := filepath.Join(engineDir, name)
			file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
			if file.Name.Name != "engine" {
				t.Errorf("package = %q, want engine", file.Name.Name)
			}
			if name == "doc.go" {
				return
			}
			wantSuffix := "V1Engine" + camelFileBase(strings.TrimSuffix(name, ".go"))
			for _, identifier := range packageDeclarations(file) {
				if identifier.Name == "init" || identifier.Name == "_" {
					continue
				}
				if !strings.HasSuffix(identifier.Name, wantSuffix) {
					t.Errorf("package-scope identifier %q must end in %q", identifier.Name, wantSuffix)
				}
			}
		})
	}
}

func packageDeclarations(file *ast.File) []*ast.Ident {
	var identifiers []*ast.Ident
	for _, declaration := range file.Decls {
		switch declaration := declaration.(type) {
		case *ast.FuncDecl:
			identifiers = append(identifiers, declaration.Name)
		case *ast.GenDecl:
			for _, specification := range declaration.Specs {
				switch specification := specification.(type) {
				case *ast.TypeSpec:
					identifiers = append(identifiers, specification.Name)
				case *ast.ValueSpec:
					identifiers = append(identifiers, specification.Names...)
				}
			}
		}
	}
	return identifiers
}

func camelFileBase(base string) string {
	parts := strings.FieldsFunc(base, func(character rune) bool {
		return !unicode.IsLetter(character) && !unicode.IsDigit(character)
	})
	var result strings.Builder
	for _, part := range parts {
		characters := []rune(part)
		if len(characters) == 0 {
			continue
		}
		result.WriteRune(unicode.ToUpper(characters[0]))
		result.WriteString(string(characters[1:]))
	}
	return result.String()
}
