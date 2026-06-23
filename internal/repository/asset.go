package repository

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ryo-arima/xaligo/internal/config"
	"github.com/ryo-arima/xaligo/internal/entity"
)

type IsoflowRepository interface {
	Render(sceneJSON []byte) ([]byte, error)
	RenderWithIcons(sceneJSON []byte, iconOverrides map[string]string) ([]byte, error)
	LoadIsoflowIcons(assets *entity.AssetSource) (map[string]string, error)
	LoadIsoflowIconManifest(path string) (map[string]string, error)
	LoadIsoflowIconManifestFS(fsys fs.FS, path string) (map[string]string, error)
}

type isoflowRepository struct{}

func NewIsoflowRepository() IsoflowRepository {
	return &isoflowRepository{}
}

func (rcvr *isoflowRepository) LoadIsoflowIcons(assets *entity.AssetSource) (map[string]string, error) {
	if assets != nil && assets.IsoflowIconsJSON != "" {
		return rcvr.LoadIsoflowIconManifestFS(assets.FS, assets.IsoflowIconsJSON)
	}
	cfg := config.New()
	return rcvr.LoadIsoflowIconManifest(filepath.Join(cfg.ProjectRoot, "etc", "resources", "aws", "isoflow-icons.json"))
}

func (rcvr *isoflowRepository) LoadIsoflowIconManifest(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return decodeIsoflowIconManifest(data)
}

func (rcvr *isoflowRepository) LoadIsoflowIconManifestFS(fsys fs.FS, path string) (map[string]string, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, err
	}
	return decodeIsoflowIconManifest(data)
}

func decodeIsoflowIconManifest(data []byte) (map[string]string, error) {
	var manifest entity.IsoflowIconManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("decode Isoflow icon manifest: %w", err)
	}
	overrides := map[string]string{}
	for id, entry := range manifest.Icons {
		if entry.DataURL != "" {
			overrides[id] = entry.DataURL
		}
	}
	return overrides, nil
}
