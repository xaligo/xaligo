package integration

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

func TestRockyWASMBuilderCopiesTsupEntrypoints(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	tsupConfig := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "tsup.config.ts"))
	dockerfile := readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile"))

	entryBlockStart := strings.Index(tsupConfig, "entry: {")
	if entryBlockStart < 0 {
		t.Fatal("tsup config does not contain an entry block")
	}
	entryBlock := tsupConfig[entryBlockStart+len("entry: {"):]
	entryBlockEnd := strings.Index(entryBlock, "}")
	if entryBlockEnd < 0 {
		t.Fatal("tsup config entry block is not closed")
	}
	entryPattern := regexp.MustCompile(`(?m)^\s*[A-Za-z0-9_-]+\s*:\s*['\"]([^'\"]+[.]ts)['\"]\s*,?\s*$`)
	entryMatches := entryPattern.FindAllStringSubmatch(entryBlock[:entryBlockEnd], -1)
	if len(entryMatches) == 0 {
		t.Fatal("tsup config entry block does not contain TypeScript entrypoints")
	}

	buildStart := strings.Index(dockerfile, "npm run build:pptx-exporter-wasm")
	if buildStart < 0 {
		t.Fatal("Rocky Dockerfile does not invoke build:pptx-exporter-wasm")
	}
	copySources := rockyWASMBuilderCopySources(dockerfile[:buildStart])
	for _, match := range entryMatches {
		required := filepath.ToSlash(filepath.Join("external", match[1]))
		if !copySources[required] && !copySources["external"] {
			t.Errorf("Rocky wasm-builder does not COPY tsup entrypoint %q before building", required)
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
