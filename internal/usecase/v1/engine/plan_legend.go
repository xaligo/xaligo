package engine

import (
	"fmt"

	"github.com/xaligo/xaligo/internal/entity"
)

func buildLegendV1EnginePlanLegend(scene *entity.PresentationScene, entries []entity.LegendEntry) []entity.LegendEntry {
	if scene == nil || len(entries) == 0 {
		return nil
	}
	out := make([]entity.LegendEntry, 0, len(entries))
	seen := map[int]bool{}
	for _, entry := range entries {
		if entry.CatalogID <= 0 || seen[entry.CatalogID] {
			continue
		}
		file, ok := scene.Files[fmt.Sprintf("item-cat-%d", entry.CatalogID)]
		if !ok || file.DataURL == "" {
			continue
		}
		if entry.Abbreviation == "" {
			entry.Abbreviation = entry.OfficialName
		}
		entry.Data = file.DataURL
		out = append(out, entry)
		seen[entry.CatalogID] = true
	}
	return out
}

func connectorLegendEntryV1EnginePlanLegend(id string, el *entity.Element, line entity.LineStyle) entity.ConnectorLegendEntry {
	kind := connectorKindV1EnginePlanConnectorDraw(el)
	entry := entity.ConnectorLegendEntry{ID: id, Kind: kind, Line: line, Source: bindingElementIDV1EnginePlanLegend(el.StartBinding), Target: bindingElementIDV1EnginePlanLegend(el.EndBinding)}
	switch kind {
	case "route":
		entry.Label = "Route line"
		entry.Description = "Network route, path, or logical reachability"
	case "traffic":
		entry.Label = "Traffic line"
		entry.Description = "Application, data, or operational communication"
	default:
		entry.Label = "Connection line"
		entry.Description = "Custom connector"
	}
	return entry
}

func bindingElementIDV1EnginePlanLegend(binding *entity.Binding) string {
	if binding == nil {
		return ""
	}
	return binding.ElementID
}

func backgroundColorV1EnginePlanLegend(scene *entity.PresentationScene) string {
	if scene.AppState != nil {
		return scene.AppState.ViewBackgroundColor
	}
	return ""
}
