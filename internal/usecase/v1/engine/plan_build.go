package engine

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

// Plan building resolves a renderer-shared physical draw plan from the
// Excalidraw-compatible canonical scene.

const (
	defaultPxPerInchV1EnginePlanBuild        = 96.0
	anchorGridV1EnginePlanBuild              = 5
	anchorGridPadPxV1EnginePlanBuild         = 4.0
	anchorGridOuterMarginPxV1EnginePlanBuild = 2.0
	anchorGridVisualPadPxV1EnginePlanBuild   = anchorGridPadPxV1EnginePlanBuild + anchorGridOuterMarginPxV1EnginePlanBuild
	pptxArrowHeadExtendPxV1EnginePlanBuild   = anchorGridVisualPadPxV1EnginePlanBuild + 2.0
	// Keep the mask smaller than the default 8 px lane gap so a jump does not
	// accidentally erase a nearby parallel lane.
	lineJumpSizePxV1EnginePlanBuild         = 6.0
	groupBorderMaskSizePxV1EnginePlanBuild  = 8.0
	connectorLabelWidthPxV1EnginePlanBuild  = 22.0
	connectorLabelHeightPxV1EnginePlanBuild = 12.0
	// 5.5pt at the default 96 PPI, expressed in layout pixels so it scales
	// together with connector-label geometry at non-default PPI values.
	connectorLabelFontPxV1EnginePlanBuild = 5.5 * defaultPxPerInchV1EnginePlanBuild / 72.0
)

// BuildPlanJSONV1EnginePlanBuild parses the Excalidraw-compatible canonical scene and returns the
// renderer-shared draw plan as JSON, applying every geometry option in opt.
func BuildPlanJSONV1EnginePlanBuild(sceneJSON string, opt entity.PlanOptions) ([]byte, error) {
	if err := ValidateRenderOptionsV1EngineOptionRender(entity.RenderOptions{
		Format: FormatPPTXV1EngineOptionRender, Theme: opt.Theme,
		PxPerInch: opt.PxPerInch, ArrowStyle: opt.ArrowStyle, ArrowStubPx: opt.ArrowStubPx, ArrowMarginPx: opt.ArrowMargin,
		PaperSize: opt.PaperSize, Orientation: opt.Orientation,
		PaperMarginIn: opt.PaperMargin, PaperMarginTopIn: opt.PaperMarginTop, PaperMarginRightIn: opt.PaperMarginRight,
		PaperMarginBottomIn: opt.PaperMarginBottom, PaperMarginLeftIn: opt.PaperMarginLeft,
	}); err != nil {
		return nil, fmt.Errorf("invalid plan options: %w", err)
	}
	opt.ArrowStyle = strings.TrimSpace(opt.ArrowStyle)
	opt.PaperSize = strings.TrimSpace(opt.PaperSize)
	opt.Orientation = strings.TrimSpace(opt.Orientation)
	var scene entity.PresentationScene
	if err := json.Unmarshal([]byte(sceneJSON), &scene); err != nil {
		return nil, err
	}
	plan := BuildPlanV1EnginePlanBuild(&scene, opt)
	return json.Marshal(plan)
}

// BuildPlanV1EnginePlanBuild converts a parsed scene into a draw plan.
func BuildPlanV1EnginePlanBuild(scene *entity.PresentationScene, opt entity.PlanOptions) entity.Plan {
	elements := make([]*entity.Element, 0, len(scene.Elements))
	for i := range scene.Elements {
		if !scene.Elements[i].IsDeleted {
			elements = append(elements, &scene.Elements[i])
		}
	}

	basePPI := opt.PxPerInch
	if basePPI <= 0 || math.IsNaN(basePPI) || math.IsInf(basePPI, 0) {
		basePPI = defaultPxPerInchV1EnginePlanBuild
	}
	contentFrame := findPaperFrameV1EnginePlanGeometry(elements)
	if contentFrame == nil {
		cb := contentBoundsV1EnginePlanGeometry(elements)
		contentFrame = &cb
	}

	contentWIn := contentFrame.W / basePPI
	contentHIn := contentFrame.H / basePPI
	paperMargins := resolvePaperMarginsV1EnginePlanPaper(opt)
	paperW, paperH, hasPaper := resolvePaperV1EnginePlanPaper(opt.PaperSize, opt.Orientation, contentWIn, contentHIn, paperMargins)

	frame := *contentFrame
	ppi := basePPI
	layoutW := contentWIn
	layoutH := contentHIn
	if hasPaper {
		availableW, availableH := paperMargins.availableV1EnginePlanPaper(paperW, paperH)
		scale := math.Min(availableW/contentWIn, availableH/contentHIn)
		ppi = basePPI / scale
		offsetXIn := paperMargins.Left + (availableW-contentWIn*scale)/2
		offsetYIn := paperMargins.Top + (availableH-contentHIn*scale)/2
		frame = rectV1EngineRouteTypes{
			X: contentFrame.X - offsetXIn*ppi,
			Y: contentFrame.Y - offsetYIn*ppi,
			W: contentFrame.W,
			H: contentFrame.H,
		}
		layoutW = paperW
		layoutH = paperH
	}

	style := resolveConnectorStyleV1EnginePlanConnectorStyle(opt.ArrowStyle)
	background := normalizeColorV1EnginePlanStyle(backgroundColorV1EnginePlanLegend(scene), "FFFFFF")
	stubPx := 20.0
	if opt.ArrowStubPx > 0 {
		stubPx = opt.ArrowStubPx
	}
	marginPx := 8.0
	if opt.ArrowMargin > 0 {
		marginPx = opt.ArrowMargin
	}

	elementsByID := map[string]*entity.Element{}
	for _, el := range elements {
		if el.ID != "" {
			elementsByID[el.ID] = el
		}
	}

	obstacles := collectObstaclesV1EnginePlanObstacle(elements)

	connectors := []*entity.Element{}
	for _, el := range elements {
		if (el.Type == "arrow" || el.Type == "line") && (el.CustomData == nil || !el.CustomData.GroupHeader) {
			connectors = append(connectors, el)
		}
	}
	prepared := prepareConnectorsV1EnginePlanConnectorPrepare(connectors, elementsByID)

	ops := []entity.DrawOp{}
	diffAreaHighlights := []entity.DrawOp{}
	frameMetadataShapes := []entity.DrawOp{}
	frameMetadataTexts := []entity.DrawOp{}

	// 1) Anchor grids first → behind the icons drawn on top.
	gridIDs := make([]string, 0, len(prepared.gridRects))
	for id := range prepared.gridRects {
		gridIDs = append(gridIDs, id)
	}
	sort.Strings(gridIDs)
	for _, id := range gridIDs {
		grid := prepared.gridRects[id]
		ops = append(ops, anchorGridOpsV1EnginePlanAnchor(id, grid.rect, frame, ppi, background)...)
	}

	// 2) Containers/shapes in scene order. Group title tags are deferred until
	// after every group border so a nested child border cannot cover a parent tag.
	headerShapes := []*entity.Element{}
	for _, el := range elements {
		if el.ID == "paper-frame" || (el.CustomData != nil && el.CustomData.PageFrame) {
			continue
		}
		if el.CustomData != nil && el.CustomData.Junction {
			continue
		}
		if el.CustomData != nil && el.CustomData.GroupHeader {
			headerShapes = append(headerShapes, el)
			continue
		}
		switch el.Type {
		case "frame", "rectangle", "ellipse", "diamond":
			if op, ok := shapeOpV1EnginePlanShape(el, frame, ppi); ok {
				if el.CustomData != nil && el.CustomData.FrameMetadata {
					frameMetadataShapes = append(frameMetadataShapes, op)
				} else if el.CustomData != nil && el.CustomData.DiffHighlight {
					diffAreaHighlights = append(diffAreaHighlights, op)
				} else {
					ops = append(ops, op)
				}
			}
		}
	}
	// 3) Connectors above containers but below icons/labels.
	for _, el := range prepared.raw {
		if op, ok := rawLineOpV1EnginePlanConnectorDraw(el, frame, ppi, style); ok {
			if highlight, highlighted := connectorDiffHighlightOpV1EnginePlanDiffHighlight(op, connectorDiffStatusV1EnginePlanDiffHighlight(el)); highlighted {
				ops = append(ops, highlight)
			}
			ops = append(ops, op)
		}
	}

	ordered := make([]preparedConnectorV1EnginePlanConnectorPrepare, len(prepared.routed))
	copy(ordered, prepared.routed)
	sort.Slice(ordered, func(i, j int) bool {
		pi := connectorKindPriorityV1EnginePlanConnectorDraw(ordered[i].req.Kind)
		pj := connectorKindPriorityV1EnginePlanConnectorDraw(ordered[j].req.Kind)
		if pi != pj {
			return pi < pj
		}
		return ordered[i].req.ID < ordered[j].req.ID
	})
	reqs := make([]routeRequestV1EngineRouteTypes, len(ordered))
	for i, pc := range ordered {
		reqs[i] = pc.req
	}
	junctions := applyRouteJunctionsV1EnginePlanJunction(reqs, stubPx)
	rOpt := defaultRouterOptionsV1EngineRouteTypes()
	rOpt.Stub = stubPx
	rOpt.LineMargin = marginPx
	rOpt.Reserved = collectContainerBorderPathsV1EnginePlanObstacle(elements)
	groupBorders := collectGroupBorderPathsV1EnginePlanObstacle(elements)
	routed := routeConnectionsV1EngineRouteBuild(reqs, obstacles, rOpt)
	elByConn := map[string]*entity.Element{}
	for _, pc := range ordered {
		elByConn[pc.req.ID] = pc.el
	}
	connectorLabels := []entity.DrawOp{}
	connectorLabelRects := []rectV1EngineRouteTypes{}
	connectorLegend := []entity.ConnectorLegendEntry{}
	for i, path := range routed {
		el := elByConn[path.ID]
		if el == nil {
			continue
		}
		connectorID := fmt.Sprintf("L%02d", i+1)
		for maskIndex, crossing := range pathBorderCrossingsV1EnginePlanObstacle(path, groupBorders) {
			ops = append(ops, groupBorderMaskOpV1EnginePlanJunction(fmt.Sprintf("%s-border-mask-%02d", path.ID, maskIndex), crossing, frame, ppi))
		}
		for maskIndex, crossing := range pathCrossingsV1EngineRouteGeometry(path, routed[:i]) {
			maskColor := lineJumpBackgroundV1EnginePlanJunction(crossing, elements, background)
			ops = append(ops, lineJumpMaskOpV1EnginePlanJunction(fmt.Sprintf("%s-jump-mask-%02d", path.ID, maskIndex), crossing, frame, ppi, maskColor))
		}
		if op, ok := polylineOpV1EnginePlanConnectorDraw(el, path.Points, frame, ppi, style); ok {
			if highlight, highlighted := connectorDiffHighlightOpV1EnginePlanDiffHighlight(op, connectorDiffStatusV1EnginePlanDiffHighlight(el)); highlighted {
				ops = append(ops, highlight)
			}
			ops = append(ops, op)
		}
		line := connectorLineV1EnginePlanConnectorDraw(el, style, ppi)
		if el.CustomData == nil || el.CustomData.UMLRelationKind == "" {
			if op, labelRect, ok := connectorIDLabelOpV1EnginePlanConnectorLabel(connectorID, path, routed, obstacles, connectorLabelRects, frame, ppi, line); ok {
				connectorLabels = append(connectorLabels, op)
				connectorLabelRects = append(connectorLabelRects, labelRect)
			}
		}
		connectorLegend = append(connectorLegend, connectorLegendEntryV1EnginePlanLegend(connectorID, el, line))
	}
	for _, junction := range junctions {
		el := elByConn[junction.ConnectorID]
		if el == nil {
			continue
		}
		ops = append(ops, junctionOpV1EnginePlanJunction(junction.ConnectorID, junction.Point, frame, ppi, connectorLineV1EnginePlanConnectorDraw(el, style, ppi)))
	}
	for _, el := range headerShapes {
		if op, ok := polygonOpV1EnginePlanShape(el, frame, ppi); ok {
			ops = append(ops, op)
		}
	}

	anchorGroupIDs := anchorGroupsV1EnginePlanAnchor(prepared.gridRects)

	// 4) Icons and labels on top so routed lines never visually cover them.
	for _, el := range elements {
		if el.ID == "paper-frame" {
			continue
		}
		switch el.Type {
		case "text":
			if op, ok := textOpV1EnginePlanText(el, frame, ppi); ok {
				applyAnchorGroupV1EnginePlanAnchor(&op, el.ID, anchorGroupIDs)
				if el.CustomData != nil && el.CustomData.FrameMetadata {
					frameMetadataTexts = append(frameMetadataTexts, op)
				} else {
					ops = append(ops, op)
				}
			}
		case "image":
			if op, ok := imageOpV1EnginePlanImage(el, scene.Files, frame, ppi); ok {
				applyAnchorGroupV1EnginePlanAnchor(&op, el.ID, anchorGroupIDs)
				ops = append(ops, op)
			}
		}
	}
	ops = append(ops, diffAreaHighlights...)
	ops = append(ops, connectorLabels...)
	ops = append(ops, frameMetadataShapes...)
	ops = append(ops, frameMetadataTexts...)

	return entity.Plan{
		Slide: entity.PlanSlide{
			W:          layoutW,
			H:          layoutH,
			Background: background,
		},
		Ops:             ops,
		Legend:          buildLegendV1EnginePlanLegend(scene, opt.LegendEntries),
		ConnectorLegend: connectorLegend,
	}
}
