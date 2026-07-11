package engine

import (
	"math"
	"sort"
)

// ── Gutter extraction ────────────────────────────────────────────────────────

func projectXV1EngineRouteGutter(rects []rectV1EngineRouteTypes) [][2]float64 {
	out := make([][2]float64, len(rects))
	for i, r := range rects {
		out[i] = [2]float64{r.X, r.X + r.W}
	}
	return out
}

func projectYV1EngineRouteGutter(rects []rectV1EngineRouteTypes) [][2]float64 {
	out := make([][2]float64, len(rects))
	for i, r := range rects {
		out[i] = [2]float64{r.Y, r.Y + r.H}
	}
	return out
}

func gutterCentersV1EngineRouteGutter(intervals [][2]float64, lo, hi float64) []float64 {
	if hi-lo < epsV1EngineRouteTypes {
		return nil
	}
	sorted := make([][2]float64, len(intervals))
	copy(sorted, intervals)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i][0] < sorted[j][0] })
	merged := [][2]float64{}
	for _, iv := range sorted {
		if len(merged) > 0 && iv[0] <= merged[len(merged)-1][1] {
			if iv[1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = iv[1]
			}
		} else {
			merged = append(merged, iv)
		}
	}
	centers := []float64{}
	cursor := lo
	for _, m := range merged {
		start, end := m[0], m[1]
		if start > cursor {
			gapStart := math.Max(cursor, lo)
			gapEnd := math.Min(start, hi)
			if gapEnd > gapStart {
				centers = append(centers, (gapStart+gapEnd)/2)
			}
		}
		cursor = math.Max(cursor, end)
	}
	if cursor < hi {
		centers = append(centers, (math.Max(cursor, lo)+hi)/2)
	}
	return centers
}

func dedupeV1EngineRouteGutter(values []float64) []float64 {
	out := []float64{}
	for _, v := range values {
		dup := false
		for _, u := range out {
			if math.Abs(u-v) < alignTolV1EngineRouteTypes {
				dup = true
				break
			}
		}
		if !dup {
			out = append(out, v)
		}
	}
	return out
}
