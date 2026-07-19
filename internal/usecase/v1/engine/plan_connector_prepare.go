package engine

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

type preparedConnectorV1EnginePlanConnectorPrepare struct {
	el  *entity.Element
	req routeRequestV1EngineRouteTypes
}

type endpointV1EnginePlanConnectorPrepare struct {
	connID    string
	rect      rectV1EngineRouteTypes
	side      sideV1EngineRouteTypes
	oppCenter ptV1EngineRouteTypes
	isSrc     bool
	profile   string
}

type preparedResultV1EnginePlanConnectorPrepare struct {
	routed    []preparedConnectorV1EnginePlanConnectorPrepare
	raw       []*entity.Element
	gridRects map[string]anchorGridRectV1EnginePlanConnectorPrepare
}

type anchorGridRectV1EnginePlanConnectorPrepare struct {
	rect       rectV1EngineRouteTypes
	background string
}

func prepareConnectorsV1EnginePlanConnectorPrepare(connectors []*entity.Element, byID map[string]*entity.Element) preparedResultV1EnginePlanConnectorPrepare {
	raw := []*entity.Element{}
	gridRects := map[string]anchorGridRectV1EnginePlanConnectorPrepare{}
	groupKeys := []string{}
	groups := map[string][]endpointV1EnginePlanConnectorPrepare{}
	type item struct {
		el         *entity.Element
		src        rectV1EngineRouteTypes
		dst        rectV1EngineRouteTypes
		srcSide    sideV1EngineRouteTypes
		dstSide    sideV1EngineRouteTypes
		srcGap     float64
		dstGap     float64
		srcAnchor  *ptV1EngineRouteTypes
		dstAnchor  *ptV1EngineRouteTypes
		srcProfile string
		dstProfile string
	}
	itemKeys := []string{}
	items := map[string]item{}

	pushGroup := func(iconID string, s sideV1EngineRouteTypes, ep endpointV1EnginePlanConnectorPrepare) {
		key := iconID + "|" + string(s)
		if _, ok := groups[key]; !ok {
			groupKeys = append(groupKeys, key)
		}
		groups[key] = append(groups[key], ep)
	}

	for _, el := range connectors {
		srcIconID := ""
		if el.StartBinding != nil {
			srcIconID = el.StartBinding.ElementID
		}
		dstIconID := ""
		if el.EndBinding != nil {
			dstIconID = el.EndBinding.ElementID
		}
		srcEl := byID[srcIconID]
		dstEl := byID[dstIconID]
		src, srcOK := rectOfV1EnginePlanGeometry(srcEl)
		dst, dstOK := rectOfV1EnginePlanGeometry(dstEl)
		if !srcOK || !dstOK || el.ID == "" {
			raw = append(raw, el)
			continue
		}
		srcGrid := anchorGridForElementV1EnginePlanConnectorPrepare(srcIconID, srcEl, src, byID)
		dstGrid := anchorGridForElementV1EnginePlanConnectorPrepare(dstIconID, dstEl, dst, byID)
		srcSide, dstSide := inferSidesV1EnginePlanGeometry(src, dst)
		if el.StartBinding != nil {
			if s, ok := sideFromFixedPointV1EnginePlanObstacle(el.StartBinding.FixedPoint); ok {
				srcSide = s
			}
		}
		if el.EndBinding != nil {
			if s, ok := sideFromFixedPointV1EnginePlanObstacle(el.EndBinding.FixedPoint); ok {
				dstSide = s
			}
		}
		var srcFixedAnchor *ptV1EngineRouteTypes
		if el.StartBinding != nil {
			explicitAnchor := el.CustomData != nil && el.CustomData.ConnectorSrcAnchor
			if p, ok := anchorFromFixedPointV1EnginePlanObstacle(src, srcSide, el.StartBinding.FixedPoint, explicitAnchor); ok {
				srcFixedAnchor = &p
			}
		}
		var dstFixedAnchor *ptV1EngineRouteTypes
		if el.EndBinding != nil {
			explicitAnchor := el.CustomData != nil && el.CustomData.ConnectorDstAnchor
			if p, ok := anchorFromFixedPointV1EnginePlanObstacle(dst, dstSide, el.EndBinding.FixedPoint, explicitAnchor); ok {
				dstFixedAnchor = &p
			}
		}
		srcProfile := umlAnchorProfileForElementV1EnginePlanConnectorPrepare(srcEl)
		dstProfile := umlAnchorProfileForElementV1EnginePlanConnectorPrepare(dstEl)
		if srcProfile == "diamond" {
			p := anchorPointForDiamondV1EnginePlanConnectorPrepare(src, srcSide)
			srcFixedAnchor = &p
		}
		if dstProfile == "diamond" {
			p := anchorPointForDiamondV1EnginePlanConnectorPrepare(dst, dstSide)
			dstFixedAnchor = &p
		}
		if srcSide == sideBottomV1EngineRouteTypes && strings.HasSuffix(srcIconID, "-lbl") {
			src = inflateRectV1EnginePlanConnectorPrepare(src, anchorGridVisualPadPxV1EnginePlanBuild)
		}
		if dstSide == sideBottomV1EngineRouteTypes && strings.HasSuffix(dstIconID, "-lbl") {
			dst = inflateRectV1EnginePlanConnectorPrepare(dst, anchorGridVisualPadPxV1EnginePlanBuild)
		}
		srcCenter := ptV1EngineRouteTypes{X: src.X + src.W/2, Y: src.Y + src.H/2}
		dstCenter := ptV1EngineRouteTypes{X: dst.X + dst.W/2, Y: dst.Y + dst.H/2}

		if _, ok := items[el.ID]; !ok {
			itemKeys = append(itemKeys, el.ID)
		}
		srcGap := 0.0
		if el.StartBinding != nil {
			srcGap = el.StartBinding.Gap
		}
		dstGap := 0.0
		if el.EndBinding != nil {
			dstGap = el.EndBinding.Gap
		}
		items[el.ID] = item{el: el, src: src, dst: dst, srcSide: srcSide, dstSide: dstSide, srcGap: srcGap, dstGap: dstGap, srcAnchor: srcFixedAnchor, dstAnchor: dstFixedAnchor, srcProfile: srcProfile, dstProfile: dstProfile}
		if needsAnchorGridV1EnginePlanConnectorPrepare(srcEl) {
			gridRects[srcIconID] = srcGrid
		}
		if needsAnchorGridV1EnginePlanConnectorPrepare(dstEl) {
			gridRects[dstIconID] = dstGrid
		}
		pushGroup(srcIconID, srcSide, endpointV1EnginePlanConnectorPrepare{connID: el.ID, rect: src, side: srcSide, oppCenter: dstCenter, isSrc: true, profile: srcProfile})
		pushGroup(dstIconID, dstSide, endpointV1EnginePlanConnectorPrepare{connID: el.ID, rect: dst, side: dstSide, oppCenter: srcCenter, isSrc: false, profile: dstProfile})
	}

	anchors := map[string]*anchorPairV1EnginePlanConnectorPrepare{}
	for _, key := range groupKeys {
		assignGroupAnchorsV1EnginePlanConnectorPrepare(groups[key], anchors)
	}

	routed := make([]preparedConnectorV1EnginePlanConnectorPrepare, 0, len(itemKeys))
	for _, id := range itemKeys {
		it := items[id]
		req := routeRequestV1EngineRouteTypes{
			ID:         id,
			Kind:       connectorKindV1EnginePlanConnectorDraw(it.el),
			Src:        it.src,
			Dst:        it.dst,
			SrcSide:    it.srcSide,
			DstSide:    it.dstSide,
			SrcGap:     it.srcGap,
			DstGap:     it.dstGap,
			SrcProfile: it.srcProfile,
			DstProfile: it.dstProfile,
			HardAvoid:  it.el.CustomData != nil && it.el.CustomData.UMLDiagramKind == "component-diagram",
		}
		if it.el.CustomData != nil {
			scale := it.el.CustomData.ConnectorScale
			if scale <= 0 {
				scale = 1
			}
			req.Bends = parseConnectorBendsV1EnginePlanConnectorPrepare(it.el.CustomData.ConnectorBends, scale)
			req.Grid = it.el.CustomData.ConnectorGrid
		}
		if a := anchors[id]; a != nil {
			req.SrcAnchor = a.src
			req.DstAnchor = a.dst
			req.SrcLane = a.srcLane
			req.DstLane = a.dstLane
		}
		if it.srcAnchor != nil {
			req.SrcAnchor = it.srcAnchor
			req.SrcLane = 0
		}
		if it.dstAnchor != nil {
			req.DstAnchor = it.dstAnchor
			req.DstLane = 0
		}
		routed = append(routed, preparedConnectorV1EnginePlanConnectorPrepare{el: it.el, req: req})
	}
	return preparedResultV1EnginePlanConnectorPrepare{routed: routed, raw: raw, gridRects: gridRects}
}

func needsAnchorGridV1EnginePlanConnectorPrepare(el *entity.Element) bool {
	if el == nil {
		return false
	}
	return el.Type == "image" || el.Type == "text"
}

func parseConnectorBendsV1EnginePlanConnectorPrepare(value string, scale float64) []ptV1EngineRouteTypes {
	if scale <= 0 {
		scale = 1
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	replacer := strings.NewReplacer(";", " ", "|", " ", "\n", " ", "\t", " ")
	tokens := strings.Fields(replacer.Replace(value))
	points := make([]ptV1EngineRouteTypes, 0, len(tokens))
	for _, token := range tokens {
		parts := strings.Split(token, ",")
		if len(parts) != 2 {
			continue
		}
		x, xErr := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		y, yErr := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		if xErr != nil || yErr != nil {
			continue
		}
		points = append(points, ptV1EngineRouteTypes{X: x * scale, Y: y * scale})
	}
	return points
}

func anchorGridForElementV1EnginePlanConnectorPrepare(id string, el *entity.Element, base rectV1EngineRouteTypes, byID map[string]*entity.Element) anchorGridRectV1EnginePlanConnectorPrepare {
	grid := anchorGridRectV1EnginePlanConnectorPrepare{rect: base, background: el.BackgroundColor}
	if strings.HasSuffix(id, "-lbl") {
		if imageEl := byID[strings.TrimSuffix(id, "-lbl")]; imageEl != nil {
			if imageRect, ok := rectOfV1EnginePlanGeometry(imageEl); ok {
				grid.rect = unionRectV1EnginePlanConnectorPrepare(grid.rect, imageRect)
				if grid.background == "" || grid.background == "transparent" {
					grid.background = imageEl.BackgroundColor
				}
			}
		}
		grid.rect = inflateRectV1EnginePlanConnectorPrepare(grid.rect, anchorGridVisualPadPxV1EnginePlanBuild)
		return grid
	}
	if labelEl := byID[id+"-lbl"]; labelEl != nil {
		if labelRect, ok := rectOfV1EnginePlanGeometry(labelEl); ok {
			grid.rect = unionRectV1EnginePlanConnectorPrepare(grid.rect, labelRect)
		}
	}
	grid.rect = inflateRectV1EnginePlanConnectorPrepare(grid.rect, anchorGridVisualPadPxV1EnginePlanBuild)
	return grid
}

func inflateRectV1EnginePlanConnectorPrepare(r rectV1EngineRouteTypes, pad float64) rectV1EngineRouteTypes {
	return rectV1EngineRouteTypes{X: r.X - pad, Y: r.Y - pad, W: r.W + pad*2, H: r.H + pad*2}
}

func unionRectV1EnginePlanConnectorPrepare(a, b rectV1EngineRouteTypes) rectV1EngineRouteTypes {
	minX := math.Min(a.X, b.X)
	minY := math.Min(a.Y, b.Y)
	maxX := math.Max(a.X+a.W, b.X+b.W)
	maxY := math.Max(a.Y+a.H, b.Y+b.H)
	return rectV1EngineRouteTypes{X: minX, Y: minY, W: maxX - minX, H: maxY - minY}
}

type anchorPairV1EnginePlanConnectorPrepare struct {
	src     *ptV1EngineRouteTypes
	dst     *ptV1EngineRouteTypes
	srcLane float64
	dstLane float64
}

func assignGroupAnchorsV1EnginePlanConnectorPrepare(eps []endpointV1EnginePlanConnectorPrepare, anchors map[string]*anchorPairV1EnginePlanConnectorPrepare) {
	if len(eps) == 0 {
		return
	}
	s := eps[0].side
	horizontal := s == sideTopV1EngineRouteTypes || s == sideBottomV1EngineRouteTypes
	sort.SliceStable(eps, func(i, j int) bool {
		if horizontal {
			return eps[i].oppCenter.X < eps[j].oppCenter.X
		}
		return eps[i].oppCenter.Y < eps[j].oppCenter.Y
	})
	meanFrac := 0.0
	for _, ep := range eps {
		meanFrac += edgeFractionV1EnginePlanConnectorPrepare(ep, horizontal)
	}
	meanFrac /= float64(len(eps))
	slots := assignSlotsV1EnginePlanConnectorPrepare(len(eps), meanFrac)
	for k, ep := range eps {
		p := anchorPointForProfileV1EnginePlanConnectorPrepare(ep.rect, s, slots[k], ep.profile)
		entry := anchors[ep.connID]
		if entry == nil {
			entry = &anchorPairV1EnginePlanConnectorPrepare{}
			anchors[ep.connID] = entry
		}
		pc := p
		lane := 0.0
		if len(eps) > 1 {
			lane = float64(k) - float64(len(eps)-1)/2
		}
		if ep.isSrc {
			entry.src = &pc
			entry.srcLane = lane
		} else {
			entry.dst = &pc
			entry.dstLane = lane
		}
	}
}

func edgeFractionV1EnginePlanConnectorPrepare(ep endpointV1EnginePlanConnectorPrepare, horizontal bool) float64 {
	var f float64
	if horizontal {
		w := ep.rect.W
		if w == 0 {
			w = 1
		}
		f = (ep.oppCenter.X - ep.rect.X) / w
	} else {
		h := ep.rect.H
		if h == 0 {
			h = 1
		}
		f = (ep.oppCenter.Y - ep.rect.Y) / h
	}
	return math.Max(0, math.Min(1, f))
}

func assignSlotsV1EnginePlanConnectorPrepare(n int, frac float64) []int {
	slots := anchorGridV1EnginePlanBuild
	if n <= 0 {
		return nil
	}
	if n >= slots {
		out := make([]int, n)
		for k := 0; k < n; k++ {
			out[k] = int(math.Round(float64(k*(slots-1)) / float64(n-1)))
		}
		return out
	}
	center := frac * float64(slots-1)
	start := int(math.Round(center - float64(n-1)/2))
	if start < 0 {
		start = 0
	}
	if start > slots-n {
		start = slots - n
	}
	out := make([]int, n)
	for k := 0; k < n; k++ {
		out[k] = start + k
	}
	return out
}

func anchorPointV1EnginePlanConnectorPrepare(r rectV1EngineRouteTypes, s sideV1EngineRouteTypes, slot int) ptV1EngineRouteTypes {
	return anchorPointForProfileV1EnginePlanConnectorPrepare(r, s, slot, "")
}

func anchorPointForProfileV1EnginePlanConnectorPrepare(r rectV1EngineRouteTypes, s sideV1EngineRouteTypes, slot int, profile string) ptV1EngineRouteTypes {
	if profile == "diamond" {
		return anchorPointForDiamondV1EnginePlanConnectorPrepare(r, s)
	}
	colX := r.X + (float64(slot)+0.5)*(r.W/float64(anchorGridV1EnginePlanBuild))
	rowY := r.Y + (float64(slot)+0.5)*(r.H/float64(anchorGridV1EnginePlanBuild))
	switch s {
	case sideTopV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: colX, Y: r.Y}
	case sideBottomV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: colX, Y: r.Y + r.H}
	case sideLeftV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: r.X, Y: rowY}
	default:
		return ptV1EngineRouteTypes{X: r.X + r.W, Y: rowY}
	}
}

func anchorPointForDiamondV1EnginePlanConnectorPrepare(r rectV1EngineRouteTypes, s sideV1EngineRouteTypes) ptV1EngineRouteTypes {
	switch s {
	case sideTopV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: r.X + r.W/2, Y: r.Y}
	case sideBottomV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: r.X + r.W/2, Y: r.Y + r.H}
	case sideLeftV1EngineRouteTypes:
		return ptV1EngineRouteTypes{X: r.X, Y: r.Y + r.H/2}
	default:
		return ptV1EngineRouteTypes{X: r.X + r.W, Y: r.Y + r.H/2}
	}
}

func umlAnchorProfileForElementV1EnginePlanConnectorPrepare(el *entity.Element) string {
	if el == nil || el.CustomData == nil || el.CustomData.UMLDiagramKind == "" || el.CustomData.UMLDiagramKind == "sequence-diagram" {
		return ""
	}
	switch el.Type {
	case "diamond":
		return "diamond"
	case "rectangle":
		return "rectangle"
	default:
		return ""
	}
}

// ── Obstacles + side inference ───────────────────────────────────────────────
