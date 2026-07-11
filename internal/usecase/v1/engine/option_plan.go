package engine

import "github.com/xaligo/xaligo/internal/entity"

// ResolvePlanOptionsV1EngineOptionPlan combines renderer options with service metadata already
// loaded by the orchestration layer. It performs no I/O.
func ResolvePlanOptionsV1EngineOptionPlan(opts entity.RenderOptions, entries []entity.ServiceEntry) entity.PlanOptions {
	legend := make([]entity.LegendEntry, 0, len(entries))
	for _, entry := range entries {
		if entry.CatalogID > 0 && entry.OfficialName != "" {
			legend = append(legend, entity.LegendEntry{
				CatalogID:    entry.CatalogID,
				Abbreviation: entry.Abbreviation,
				OfficialName: entry.OfficialName,
			})
		}
	}
	return entity.PlanOptions{
		Theme: opts.Theme, PxPerInch: opts.PxPerInch, ArrowStyle: opts.ArrowStyle,
		ArrowStubPx: opts.ArrowStubPx, ArrowMargin: opts.ArrowMarginPx,
		PaperSize: opts.PaperSize, Orientation: opts.Orientation,
		PaperMargin: opts.PaperMarginIn, PaperMarginTop: opts.PaperMarginTopIn, PaperMarginRight: opts.PaperMarginRightIn,
		PaperMarginBottom: opts.PaperMarginBottomIn, PaperMarginLeft: opts.PaperMarginLeftIn, LegendEntries: legend,
	}
}
