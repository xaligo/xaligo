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

func TestRockyWASMBuilderCopiesTypeScriptBuildInputs(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	packageJSON := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "pptx-exporter", "package.json"))
	buildTool := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "pptx-exporter", "tool", "build.mjs"))
	dockerfile := readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile"))

	buildStart := strings.Index(dockerfile, "npm run build:pptx-exporter-wasm")
	if buildStart < 0 {
		t.Fatal("Rocky Dockerfile does not invoke build:pptx-exporter-wasm")
	}
	copySources := rockyWASMBuilderCopySources(dockerfile[:buildStart])
	if !strings.Contains(packageJSON, `"build": "node tool/build.mjs"`) {
		t.Fatal("external build script does not invoke tool/build.mjs")
	}
	for _, entrypoint := range []string{"index.ts", "command.ts"} {
		compiledEntrypoint := strings.TrimSuffix(entrypoint, ".ts") + ".js"
		if !strings.Contains(buildTool, "'"+compiledEntrypoint+"'") {
			t.Fatalf("external build script does not contain TypeScript entrypoint %q", entrypoint)
		}
		required := filepath.ToSlash(filepath.Join("external", "pptx-exporter", entrypoint))
		if !copySources[required] && !copySources["external/pptx-exporter"] {
			t.Errorf("Rocky wasm-builder does not COPY TypeScript entrypoint %q before building", required)
		}
	}
	if !copySources["external/pptx-exporter/tool"] {
		t.Error("Rocky wasm-builder does not COPY external/pptx-exporter/tool before building")
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

func TestNPMDependencyGraphExcludesRemovedNativeTools(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	for _, path := range []string{
		filepath.Join(repositoryRoot, "package.json"),
		filepath.Join(repositoryRoot, "external", "pptx-exporter", "package.json"),
	} {
		var manifest struct {
			Dependencies    map[string]string `json:"dependencies"`
			DevDependencies map[string]string `json:"devDependencies"`
			Overrides       map[string]string `json:"overrides"`
		}
		if err := json.Unmarshal([]byte(readIntegrationFile(t, path)), &manifest); err != nil {
			t.Fatalf("parse %s: %v", path, err)
		}
		for _, dependency := range []string{"@resvg/resvg-js", "esbuild"} {
			if _, ok := manifest.Dependencies[dependency]; ok {
				t.Errorf("%s still declares dependency %q", path, dependency)
			}
			if _, ok := manifest.DevDependencies[dependency]; ok {
				t.Errorf("%s still declares dev dependency %q", path, dependency)
			}
			if _, ok := manifest.Overrides[dependency]; ok {
				t.Errorf("%s still overrides dependency %q", path, dependency)
			}
		}
	}

	var lock struct {
		Packages map[string]json.RawMessage `json:"packages"`
	}
	lockPath := filepath.Join(repositoryRoot, "package-lock.json")
	if err := json.Unmarshal([]byte(readIntegrationFile(t, lockPath)), &lock); err != nil {
		t.Fatalf("parse %s: %v", lockPath, err)
	}
	for packagePath := range lock.Packages {
		if packagePath == "node_modules/esbuild" ||
			strings.HasPrefix(packagePath, "node_modules/@esbuild/") ||
			packagePath == "node_modules/@resvg/resvg-js" ||
			strings.HasPrefix(packagePath, "node_modules/@resvg/resvg-js-") {
			t.Errorf("%s still contains removed package %q", lockPath, packagePath)
		}
	}
}

func TestDistributionsIncludeBundledPptxLicenses(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	notice := readIntegrationFile(t, filepath.Join(repositoryRoot, "THIRD_PARTY_LICENSES"))
	var lock struct {
		Packages map[string]struct {
			Version string `json:"version"`
		} `json:"packages"`
	}
	lockPath := filepath.Join(repositoryRoot, "package-lock.json")
	if err := json.Unmarshal([]byte(readIntegrationFile(t, lockPath)), &lock); err != nil {
		t.Fatalf("parse %s: %v", lockPath, err)
	}
	for _, dependency := range []struct {
		packageName string
		heading     string
	}{
		{packageName: "pptxgenjs", heading: "PptxGenJS"},
		{packageName: "jszip", heading: "JSZip"},
		{packageName: "pako", heading: "pako"},
		{packageName: "lie", heading: "lie"},
		{packageName: "immediate", heading: "immediate"},
		{packageName: "setimmediate", heading: "setimmediate"},
	} {
		packagePath := "node_modules/" + dependency.packageName
		version := lock.Packages[packagePath].Version
		if version == "" {
			t.Fatalf("resolve bundled dependency version for %q", packagePath)
		}
		heading := dependency.heading + " " + version
		if !strings.Contains(notice, heading) {
			t.Errorf("third-party license notice does not contain %q", heading)
		}
	}

	var manifest struct {
		Files []string `json:"files"`
	}
	manifestPath := filepath.Join(repositoryRoot, "package.json")
	if err := json.Unmarshal([]byte(readIntegrationFile(t, manifestPath)), &manifest); err != nil {
		t.Fatalf("parse %s: %v", manifestPath, err)
	}
	if !containsIntegrationString(manifest.Files, "THIRD_PARTY_LICENSES") {
		t.Error("npm package does not include THIRD_PARTY_LICENSES")
	}

	debScript := readIntegrationFile(t, filepath.Join(repositoryRoot, "scripts", "build", "build-deb.sh"))
	rpmScript := readIntegrationFile(t, filepath.Join(repositoryRoot, "scripts", "build", "build-rpm.sh"))
	if !strings.Contains(
		debScript,
		`cat LICENSE THIRD_PARTY_LICENSES > "$WORK_DIR/usr/share/doc/${PACKAGE_NAME}/copyright"`,
	) {
		t.Error("Debian copyright file does not include bundled license terms")
	}
	if !strings.Contains(debScript, "install -m 0644 THIRD_PARTY_LICENSES") {
		t.Error("Debian package does not install THIRD_PARTY_LICENSES")
	}
	if !strings.Contains(rpmScript, "install -m 0644 THIRD_PARTY_LICENSES") ||
		!strings.Contains(rpmScript, "%doc /usr/share/doc/%{name}/THIRD_PARTY_LICENSES") {
		t.Error("RPM package does not install THIRD_PARTY_LICENSES as documentation")
	}
	if !strings.Contains(rpmScript, "License: MIT AND Zlib") {
		t.Error("RPM metadata does not declare bundled license terms")
	}
}

func TestDockerToolchainsMatchRepositoryRequirements(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	goMod := readIntegrationFile(t, filepath.Join(repositoryRoot, "go.mod"))
	cargoManifest := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "engine", "Cargo.toml"))
	packageJSON := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "pptx-exporter", "package.json"))
	dockerfiles := []string{
		readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile")),
		readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "ubuntu.Dockerfile")),
	}

	goVersion := requiredVersion(t, goMod, `(?m)^toolchain go([0-9.]+)$`, "Go toolchain")
	rustVersion := requiredVersion(t, cargoManifest, `(?m)^rust-version = "([0-9.]+)"$`, "Rust toolchain")
	nodeVersion := requiredNodeMajor(t, packageJSON)
	for index, dockerfile := range dockerfiles {
		if !strings.Contains(dockerfile, "ARG GO_VERSION="+goVersion) {
			t.Errorf("Dockerfile %d does not pin required Go version %s", index, goVersion)
		}
		if !strings.Contains(dockerfile, "FROM node:"+nodeVersion+"-") {
			t.Errorf("Dockerfile %d does not use required Node major %s", index, nodeVersion)
		}
		if !strings.Contains(dockerfile, "ARG RUST_VERSION="+rustVersion+".") {
			t.Errorf("Dockerfile %d does not pin required Rust series %s", index, rustVersion)
		}
	}
}

func TestNativePackageBuildLinksRustEngineAndSQLiteFTS(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	common := readIntegrationFile(t, filepath.Join(repositoryRoot, "scripts", "build", "common.sh"))
	for _, required := range []string{
		"cargo build",
		"CGO_ENABLED=1",
		"xaligo_engine sqlite_fts5 sqlite_omit_load_extension",
		"external/engine/lib/libxaligo_engine.a",
	} {
		if !strings.Contains(common, required) {
			t.Errorf("native package build does not contain %q", required)
		}
	}
	if strings.Contains(common, "CGO_ENABLED=0") {
		t.Error("native package build still disables cgo")
	}
}

func TestReleaseBuildsEachNativeTargetWithRustEngine(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	workflow := readIntegrationFile(t, filepath.Join(repositoryRoot, ".github", "workflows", "release.yml"))
	for _, target := range []string{
		"darwin-amd64",
		"darwin-arm64",
		"linux-amd64",
		"linux-arm64",
		"windows-amd64",
		"windows-arm64",
	} {
		if !strings.Contains(workflow, "target: "+target) {
			t.Errorf("release workflow does not build native target %q", target)
		}
	}
	for _, required := range []string{
		"make test-engine",
		"NPM_SKIP_WASM: '1'",
		"NPM_PACKAGE_TARGETS: none",
		"pattern: xaligo-native-*",
	} {
		if !strings.Contains(workflow, required) {
			t.Errorf("release workflow does not contain %q", required)
		}
	}
}

func containsIntegrationString(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
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
		t.Fatalf("parse external/pptx-exporter/package.json: %v", err)
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
