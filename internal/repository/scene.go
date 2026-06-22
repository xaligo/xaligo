package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ryo-arima/xaligo/internal/entity"
	"github.com/ryo-arima/xaligo/internal/share"
)

var (
	IRSRS001  = share.NewMCode("IRSRS-001", "Read scene read failed")
	IRSRS002  = share.NewMCode("IRSRS-002", "Read scene parse failed")
	IRSRS003  = share.NewMCode("IRSRS-003", "Read scene initialize files branch")
	IRSRS004  = share.NewMCode("IRSRS-004", "Read scene initialize elements branch")
	IRSWSC001 = share.NewMCode("IRSWSC-001", "Write scene marshal failed")
	IRSWSC002 = share.NewMCode("IRSWSC-002", "Write scene write failed")
)

// ReadScene reads a .excalidraw JSON file and returns the parsed Scene.
func ReadScene(path string) (*entity.Scene, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.ERROR(IRSRS001, "read failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("read scene %s: %w", path, err)
	}
	var scene entity.Scene
	if err := json.Unmarshal(data, &scene); err != nil {
		logger.ERROR(IRSRS002, "parse failed", map[string]any{"path": path, "error": err})
		return nil, fmt.Errorf("parse scene %s: %w", path, err)
	}
	if scene.Files == nil {
		logger.DEBUG(IRSRS003, "branch initialize files", map[string]any{"path": path})
		scene.Files = map[string]map[string]interface{}{}
	}
	if scene.Elements == nil {
		logger.DEBUG(IRSRS004, "branch initialize elements", map[string]any{"path": path})
		scene.Elements = []map[string]interface{}{}
	}
	return &scene, nil
}

// WriteScene serialises the Scene and writes it back to path.
func WriteScene(scene *entity.Scene, path string) error {
	data, err := json.MarshalIndent(scene, "", "  ")
	if err != nil {
		logger.ERROR(IRSWSC001, "marshal failed", map[string]any{"error": err})
		return fmt.Errorf("marshal scene: %w", err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		logger.ERROR(IRSWSC002, "write failed", map[string]any{"path": path, "error": err})
		return fmt.Errorf("write scene %s: %w", path, err)
	}
	return nil
}
