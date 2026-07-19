package integration

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var versionNumberPattern = regexp.MustCompile(`[0-9]+(?:\.[0-9]+)*`)

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

func TestRockyWASMBuilderCopySourcesExist(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	dockerfile := readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile"))

	buildStart := strings.Index(dockerfile, "npm run build:pptx-exporter-wasm")
	if buildStart < 0 {
		t.Fatal("Rocky Dockerfile does not invoke build:pptx-exporter-wasm")
	}
	for source := range rockyWASMBuilderCopySources(dockerfile[:buildStart]) {
		if _, err := os.Stat(filepath.Join(repositoryRoot, filepath.FromSlash(source))); err != nil {
			t.Errorf("Rocky wasm-builder COPY source %q: %v", source, err)
		}
	}
}

func TestDockerToolchainsMatchRepositoryRequirements(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	goMod := readIntegrationFile(t, filepath.Join(repositoryRoot, "go.mod"))
	packageJSON := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "package.json"))
	dockerfiles := []string{
		readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile")),
		readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "ubuntu.Dockerfile")),
	}

	goVersion := requiredVersion(t, goMod, `(?m)^toolchain go([0-9.]+)$`, "Go toolchain")
	nodeVersion := requiredNodeMajor(t, packageJSON)
	for index, dockerfile := range dockerfiles {
		if !strings.Contains(dockerfile, "ARG GO_VERSION="+goVersion) {
			t.Errorf("Dockerfile %d does not pin required Go version %s", index, goVersion)
		}
		if !strings.Contains(dockerfile, "FROM node:"+nodeVersion+"-") {
			t.Errorf("Dockerfile %d does not use required Node major %s", index, nodeVersion)
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

func requiredVersion(t *testing.T, source, pattern, requirement string) string {
	t.Helper()
	matches := regexp.MustCompile(pattern).FindStringSubmatch(source)
	if len(matches) != 2 {
		t.Fatalf("resolve %s", requirement)
	}
	return matches[1]
}

func requiredNodeMajor(t *testing.T, packageJSON string) string {
	t.Helper()
	var manifest struct {
		Engines struct {
			Node string `json:"node"`
		} `json:"engines"`
	}
	if err := json.Unmarshal([]byte(packageJSON), &manifest); err != nil {
		t.Fatalf("parse external/package.json: %v", err)
	}
	version := versionNumberPattern.FindString(manifest.Engines.Node)
	if version == "" {
		t.Fatal("resolve required Node version")
	}
	return strings.SplitN(version, ".", 2)[0]
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
