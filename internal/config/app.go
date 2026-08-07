package config

import (
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v3"
)

const DefaultServePort = 8080

// appYAML mirrors the structure of etc/app.yaml.
type appYAML struct {
	Paths struct {
		AssetPackage      string `yaml:"asset_package"`
		ServiceCatalogCSV string `yaml:"service_catalog_csv"`
		PptxExporterWASM  string `yaml:"pptx_exporter_wasm"`
		AssetsDB          string `yaml:"assets_db"`
		ProjectDB         string `yaml:"project_db"`
	} `yaml:"paths"`
	Item struct {
		IconSize float64 `yaml:"icon_size"`
	} `yaml:"item"`
	Serve struct {
		Port int `yaml:"port"`
	} `yaml:"serve"`
}

// ServeConfig holds live-preview server defaults.
type ServeConfig struct {
	Port int
}

// Config holds application-wide configuration resolved from etc/app.yaml.
type Config struct {
	ProjectRoot      string
	AssetDir_        string  // absolute path to Asset-Package
	SvcCatalogCSV    string  // absolute path to service-catalog.csv
	PptxExporterWASM string  // absolute path to the PPTX WASM exporter
	AssetsDB         string  // absolute path to the embedded SVG registry
	ProjectDB        string  // absolute path to the durable project search index
	ItemIconSize     float64 // default max icon size for <item> elements (px)
	Serve            ServeConfig
}

// New loads etc/app.yaml from the project root and returns a resolved Config.
// All paths fall back to sensible defaults when the config file is absent.
func New() *Config {
	root := findProjectRoot()

	def := appYAML{}
	def.Paths.AssetPackage = "etc/resources/aws/svg"
	def.Paths.ServiceCatalogCSV = "etc/resources/aws/service-catalog.csv"
	def.Paths.PptxExporterWASM = "external/exporter/wasm/xaligo.wasm"
	def.Paths.AssetsDB = "xaligo-assets.db"
	def.Paths.ProjectDB = ".xaligo/project.db"
	def.Item.IconSize = 32.0
	def.Serve.Port = DefaultServePort

	yamlPath := filepath.Join(root, "etc", "resources", "aws", "app.yaml")
	if data, err := os.ReadFile(yamlPath); err == nil {
		_ = yaml.Unmarshal(data, &def)
	}
	if def.Serve.Port < 1 || def.Serve.Port > 65535 {
		def.Serve.Port = DefaultServePort
	}

	abs := func(rel string) string {
		if filepath.IsAbs(rel) {
			return rel
		}
		return filepath.Join(root, rel)
	}

	return &Config{
		ProjectRoot:      root,
		AssetDir_:        abs(def.Paths.AssetPackage),
		SvcCatalogCSV:    abs(def.Paths.ServiceCatalogCSV),
		PptxExporterWASM: abs(def.Paths.PptxExporterWASM),
		AssetsDB:         abs(def.Paths.AssetsDB),
		ProjectDB:        abs(def.Paths.ProjectDB),
		ItemIconSize:     def.Item.IconSize,
		Serve:            ServeConfig{Port: def.Serve.Port},
	}
}

// AssetDir returns the absolute path to the Asset-Package directory.
func (rcvr *Config) AssetDir() string { return rcvr.AssetDir_ }

// ServiceCatalogCSVPath returns the absolute path to service-catalog.csv.
func (rcvr *Config) ServiceCatalogCSVPath() string { return rcvr.SvcCatalogCSV }

// findProjectRoot walks up from cwd until it finds go.mod, then returns that dir.
func findProjectRoot() string {
	if home := os.Getenv("XALIGO_HOME"); home != "" {
		if abs, err := filepath.Abs(home); err == nil {
			return abs
		}
		return home
	}

	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	dir := cwd
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	if executable, err := os.Executable(); err == nil {
		binDir := filepath.Dir(executable)
		candidates := []string{
			filepath.Clean(filepath.Join(binDir, "..", "lib", "xaligo")),
			filepath.Clean(filepath.Join(binDir, "..", "share", "xaligo")),
		}
		for _, candidate := range candidates {
			if isRuntimeRoot(candidate) {
				return candidate
			}
		}
	}
	return cwd
}

func isRuntimeRoot(root string) bool {
	info, err := os.Stat(filepath.Join(root, "etc", "resources", "aws", "app.yaml"))
	return err == nil && !info.IsDir()
}
