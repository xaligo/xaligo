package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// appYAML mirrors the structure of etc/app.yaml.
type appYAML struct {
	Paths struct {
		AssetPackage      string `yaml:"asset_package"`
		ServiceCatalogCSV string `yaml:"service_catalog_csv"`
		OutputFrames      string `yaml:"output_frames"`
		PptxExporterWASM  string `yaml:"pptx_exporter_wasm"`
	} `yaml:"paths"`
	Legend struct {
		OffsetX  float64 `yaml:"offset_x"`
		OffsetY  float64 `yaml:"offset_y"`
		IconSize int     `yaml:"icon_size"`
		FontSize int     `yaml:"font_size"`
	} `yaml:"legend"`
	Item struct {
		IconSize float64 `yaml:"icon_size"`
	} `yaml:"item"`
}

// LegendConfig holds resolved legend defaults.
type LegendConfig struct {
	OffsetX  float64
	OffsetY  float64
	IconSize int
	FontSize int
}

// Config holds application-wide configuration resolved from etc/app.yaml.
type Config struct {
	ProjectRoot      string
	AssetDir_        string // absolute path to Asset-Package
	OutFramesDir     string // absolute path to generated frames output dir
	SvcCatalogCSV    string // absolute path to service-catalog.csv
	PptxExporterWASM string // absolute path to the PPTX WASM exporter
	Legend           LegendConfig
	ItemIconSize     float64 // default max icon size for <item> elements (px)
}

// New loads etc/app.yaml from the project root and returns a resolved Config.
// All paths fall back to sensible defaults when the config file is absent.
func New() *Config {
	root := findProjectRoot()

	def := appYAML{}
	def.Paths.AssetPackage = "etc/resources/aws/svg"
	def.Paths.ServiceCatalogCSV = "etc/resources/aws/service-catalog.csv"
	def.Paths.OutputFrames = "output/aws-frames"
	def.Paths.PptxExporterWASM = "external/wasm/xaligo.wasm"
	def.Legend.OffsetX = 120
	def.Legend.OffsetY = 0
	def.Legend.IconSize = 32
	def.Legend.FontSize = 12
	def.Item.IconSize = 32.0

	yamlPath := filepath.Join(root, "etc", "resources", "aws", "app.yaml")
	if data, err := os.ReadFile(yamlPath); err == nil {
		_ = yaml.Unmarshal(data, &def)
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
		OutFramesDir:     abs(def.Paths.OutputFrames),
		SvcCatalogCSV:    abs(def.Paths.ServiceCatalogCSV),
		PptxExporterWASM: abs(def.Paths.PptxExporterWASM),
		ItemIconSize:     def.Item.IconSize,
		Legend: LegendConfig{
			OffsetX:  def.Legend.OffsetX,
			OffsetY:  def.Legend.OffsetY,
			IconSize: def.Legend.IconSize,
			FontSize: def.Legend.FontSize,
		},
	}
}

// AssetDir returns the absolute path to the Asset-Package directory.
func (rcvr *Config) AssetDir() string { return rcvr.AssetDir_ }

// OutputFramesDir returns the absolute path to the frames output directory.
func (rcvr *Config) OutputFramesDir() string { return rcvr.OutFramesDir }

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
