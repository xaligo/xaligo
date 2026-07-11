package usecase_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestUsecaseRootFilesStartWithCompleteComponent(t *testing.T) {
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve test file path")
	}
	repositoryRoot := filepath.Clean(filepath.Join(filepath.Dir(testFile), "../../../.."))
	usecaseDir := filepath.Join(repositoryRoot, "internal", "usecase")
	entries, err := os.ReadDir(usecaseDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		t.Run(name, func(t *testing.T) {
			file, err := parser.ParseFile(token.NewFileSet(), filepath.Join(usecaseDir, name), nil, 0)
			if err != nil {
				t.Fatal(err)
			}
			interfaceName, concreteName, constructorName := leadingUsecaseComponent(file)
			if !strings.HasSuffix(interfaceName, "Usecase") || !ast.IsExported(interfaceName) {
				t.Fatalf("first component declaration = %q, want exported XxxUsecase interface", interfaceName)
			}
			if concreteName != lowerFirst(interfaceName) {
				t.Fatalf("concrete type = %q, want %q", concreteName, lowerFirst(interfaceName))
			}
			if constructorName != "New"+interfaceName {
				t.Fatalf("constructor = %q, want %q", constructorName, "New"+interfaceName)
			}
		})
	}
}

func leadingUsecaseComponent(file *ast.File) (string, string, string) {
	var typeNames []string
	constructor := ""
	for _, declaration := range file.Decls {
		switch declaration := declaration.(type) {
		case *ast.GenDecl:
			if declaration.Tok != token.TYPE {
				continue
			}
			for _, specification := range declaration.Specs {
				if typed, ok := specification.(*ast.TypeSpec); ok {
					typeNames = append(typeNames, typed.Name.Name)
					if len(typeNames) == 2 {
						break
					}
				}
			}
		case *ast.FuncDecl:
			if declaration.Recv == nil && strings.HasPrefix(declaration.Name.Name, "New") {
				constructor = declaration.Name.Name
			}
		}
		if len(typeNames) >= 2 && constructor != "" {
			break
		}
	}
	if len(typeNames) < 2 {
		return "", "", constructor
	}
	return typeNames[0], typeNames[1], constructor
}

func lowerFirst(value string) string {
	if value == "" {
		return ""
	}
	return strings.ToLower(value[:1]) + value[1:]
}
