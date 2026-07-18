package integration

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestRockyWASMBuilderCopiesTypeScriptEntrypoints(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	packageJSON := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "package.json"))
	dockerfile := readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile"))

	buildStart := strings.Index(dockerfile, "npm run build:pptx-exporter-wasm")
	if buildStart < 0 {
		t.Fatal("Rocky Dockerfile does not invoke build:pptx-exporter-wasm")
	}
	copySources := rockyWASMBuilderCopySources(dockerfile[:buildStart])
	for _, entrypoint := range []string{"index.ts", "command.ts"} {
		if !strings.Contains(packageJSON, "esbuild "+entrypoint) {
			t.Fatalf("external build script does not contain TypeScript entrypoint %q", entrypoint)
		}
		required := filepath.ToSlash(filepath.Join("external", entrypoint))
		if !copySources[required] && !copySources["external"] {
			t.Errorf("Rocky wasm-builder does not COPY TypeScript entrypoint %q before building", required)
		}
	}
}

func integrationRepositoryRoot(t *testing.T) string {
	t.Helper()
	_, testFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve integration test path")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(testFile), "../.."))
}

func readIntegrationFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func rockyWASMBuilderCopySources(dockerfile string) map[string]bool {
	sources := map[string]bool{}
	normalized := strings.ReplaceAll(dockerfile, "\\\n", " ")
	for _, line := range strings.Split(normalized, "\n") {
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) < 3 || fields[0] != "COPY" {
			continue
		}
		for _, source := range fields[1 : len(fields)-1] {
			if strings.HasPrefix(source, "--") {
				continue
			}
			cleaned := filepath.ToSlash(filepath.Clean(strings.TrimPrefix(source, "./")))
			sources[cleaned] = true
		}
	}
	return sources
}
