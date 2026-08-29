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

func TestNPMDependencyGraphExcludesRemovedNativeTools(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	for _, path := range []string{filepath.Join(repositoryRoot, "package.json")} {
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
	cargoLock := readIntegrationFile(t, filepath.Join(repositoryRoot, "external", "exporter", "Cargo.lock"))
	if !strings.Contains(cargoLock, "name = \"pptx\"\nversion = \"0.1.0\"") {
		t.Fatal("Rust exporter does not lock pptx 0.1.0")
	}
	if !strings.Contains(notice, "pptx 0.1.0") {
		t.Error("third-party license notice does not contain pptx 0.1.0")
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
	dockerfiles := []string{
		readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile")),
		readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "ubuntu.Dockerfile")),
	}

	goVersion := requiredVersion(t, goMod, `(?m)^toolchain go([0-9.]+)$`, "Go toolchain")
	rustVersion := requiredVersion(t, cargoManifest, `(?m)^rust-version = "([0-9.]+)"$`, "Rust toolchain")
	for index, dockerfile := range dockerfiles {
		if !strings.Contains(dockerfile, "ARG GO_VERSION="+goVersion) {
			t.Errorf("Dockerfile %d does not pin required Go version %s", index, goVersion)
		}
		if !strings.Contains(dockerfile, "ARG RUST_VERSION="+rustVersion+".") {
			t.Errorf("Dockerfile %d does not pin required Rust series %s", index, rustVersion)
		}
	}
}

func TestRockyDockerToolchainUsesCompatibleCurlPackage(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	rocky := readIntegrationFile(t, filepath.Join(repositoryRoot, "docker", "rocky.Dockerfile"))
	if !strings.Contains(rocky, "\n    curl-minimal \\\n") {
		t.Error("Rocky Dockerfile does not install curl-minimal")
	}
	if strings.Contains(rocky, "\n    curl \\\n") {
		t.Error("Rocky Dockerfile installs full curl, which conflicts with the base image's curl-minimal package")
	}
}

func TestNativePackageBuildLinksRustEngineAndSQLiteFTS(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	common := readIntegrationFile(t, filepath.Join(repositoryRoot, "scripts", "build", "common.sh"))
	for _, required := range []string{
		"cargo build",
		"CGO_ENABLED=1",
		"xaligo_engine xaligo_exporter sqlite_fts5 sqlite_omit_load_extension",
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

func TestReleaseProvisionsWindowsARM64CGOToolchain(t *testing.T) {
	repositoryRoot := integrationRepositoryRoot(t)
	workflow := readIntegrationFile(t, filepath.Join(repositoryRoot, ".github", "workflows", "release.yml"))
	for _, required := range []string{
		"name: Install LLVM-MinGW for Windows ARM64",
		"Get-FileHash -Algorithm SHA256",
		"cgo_cc: aarch64-w64-mingw32-clang",
		"CC: ${{ matrix.cgo_cc || '' }}",
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
		t.Fatalf("parse external/exporter/package.json: %v", err)
	}
	version := versionNumberPattern.FindString(manifest.Engines.Node)
	if version == "" {
		t.Fatal("resolve required Node version")
	}
	return strings.SplitN(version, ".", 2)[0]
}
