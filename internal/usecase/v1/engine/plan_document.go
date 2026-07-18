package engine

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
)

const documentPlanSchemaVersionV1EnginePlanDocument = 2

// BuildDocumentPlanJSONV1EnginePlanDocument converts the canonical scene into
// an ordered page document. By default every page frame becomes one page;
// combineFrames retains the legacy single-canvas result.
func BuildDocumentPlanJSONV1EnginePlanDocument(sceneJSON string, opt entity.PlanOptions, combineFrames bool) ([]byte, error) {
	// Reuse the single-page entry point for option validation so both plan
	// contracts reject the same invalid presentation settings.
	if _, err := BuildPlanJSONV1EnginePlanBuild(sceneJSON, opt); err != nil {
		return nil, err
	}
	var scene entity.PresentationScene
	if err := json.Unmarshal([]byte(sceneJSON), &scene); err != nil {
		return nil, fmt.Errorf("decode presentation scene: %w", err)
	}
	document := BuildDocumentPlanV1EnginePlanDocument(&scene, opt, combineFrames)
	return json.Marshal(document)
}

// BuildDocumentPlanV1EnginePlanDocument builds pages in source/frame order.
func BuildDocumentPlanV1EnginePlanDocument(scene *entity.PresentationScene, opt entity.PlanOptions, combineFrames bool) entity.DocumentPlan {
	combined := BuildPlanV1EnginePlanBuild(scene, opt)
	frames := pageFrameElementsV1EnginePlanDocument(scene)
	if combineFrames || len(frames) == 0 {
		return entity.DocumentPlan{
			SchemaVersion:   documentPlanSchemaVersionV1EnginePlanDocument,
			Pages:           []entity.DocumentPage{documentPageFromPlanV1EnginePlanDocument(combinedPageIDV1EnginePlanDocument(frames), combined)},
			Legend:          combined.Legend,
			ConnectorLegend: combined.ConnectorLegend,
		}
	}

	pages := make([]entity.DocumentPage, 0, len(frames))
	connectorLegend := make([]entity.ConnectorLegendEntry, 0, len(combined.ConnectorLegend))
	connectorIndex := 1
	for _, frame := range frames {
		pageScene := projectPageSceneV1EnginePlanDocument(scene, frame)
		pagePlan := BuildPlanV1EnginePlanBuild(&pageScene, opt)
		connectorIndex = renumberPageConnectorsV1EnginePlanDocument(&pagePlan, connectorIndex)
		connectorLegend = append(connectorLegend, pagePlan.ConnectorLegend...)
		pages = append(pages, documentPageFromPlanV1EnginePlanDocument(frame.CustomData.FrameID, pagePlan))
	}
	return entity.DocumentPlan{
		SchemaVersion:   documentPlanSchemaVersionV1EnginePlanDocument,
		Pages:           pages,
		Legend:          combined.Legend,
		ConnectorLegend: connectorLegend,
	}
}

// NormalizeDocumentPageSizesV1EnginePlanDocument applies the single common
// page size required by PPTX. Smaller frame pages are centred without scaling.
func NormalizeDocumentPageSizesV1EnginePlanDocument(document *entity.DocumentPlan) {
	if document == nil || len(document.Pages) < 2 {
		return
	}
	maxW, maxH := 0.0, 0.0
	for _, page := range document.Pages {
		maxW = math.Max(maxW, page.Slide.W)
		maxH = math.Max(maxH, page.Slide.H)
	}
	for pageIndex := range document.Pages {
		page := &document.Pages[pageIndex]
		offsetX := (maxW - page.Slide.W) / 2
		offsetY := (maxH - page.Slide.H) / 2
		for opIndex := range page.Ops {
			page.Ops[opIndex].X += offsetX
			page.Ops[opIndex].Y += offsetY
		}
		page.Slide.W = maxW
		page.Slide.H = maxH
	}
}

func pageFrameElementsV1EnginePlanDocument(scene *entity.PresentationScene) []*entity.Element {
	if scene == nil {
		return nil
	}
	frames := make([]*entity.Element, 0)
	for elementIndex := range scene.Elements {
		element := &scene.Elements[elementIndex]
		if element.IsDeleted || element.CustomData == nil || !element.CustomData.PageFrame {
			continue
		}
		frames = append(frames, element)
	}
	return frames
}

func combinedPageIDV1EnginePlanDocument(frames []*entity.Element) string {
	if len(frames) == 1 && strings.TrimSpace(frames[0].CustomData.FrameID) != "" {
		return frames[0].CustomData.FrameID
	}
	return "combined"
}

func documentPageFromPlanV1EnginePlanDocument(id string, plan entity.Plan) entity.DocumentPage {
	if strings.TrimSpace(id) == "" {
		id = "frame"
	}
	return entity.DocumentPage{ID: id, Slide: plan.Slide, Ops: plan.Ops}
}

func projectPageSceneV1EnginePlanDocument(scene *entity.PresentationScene, frame *entity.Element) entity.PresentationScene {
	projected := entity.PresentationScene{
		Elements: make([]entity.Element, 0, len(scene.Elements)),
		Files:    scene.Files,
		AppState: scene.AppState,
	}
	projected.Elements = append(projected.Elements, entity.Element{
		ID: "paper-frame", Type: "frame",
		X: frame.X, Y: frame.Y, Width: frame.Width, Height: frame.Height,
		StrokeColor: frame.StrokeColor, BackgroundColor: "transparent",
	})
	owners := semanticFrameOwnersV1EnginePlanDocument(scene)
	for _, element := range scene.Elements {
		if element.IsDeleted || element.ID == "paper-frame" {
			continue
		}
		if element.ID == frame.ID || elementBelongsToPageV1EnginePlanDocument(element, frame, owners) {
			projected.Elements = append(projected.Elements, element)
		}
	}
	return projected
}

func semanticFrameOwnersV1EnginePlanDocument(scene *entity.PresentationScene) map[string]string {
	owners := map[string]string{}
	parents := map[string]string{}
	for _, element := range scene.Elements {
		if element.CustomData == nil {
			continue
		}
		if element.CustomData.PageFrame {
			owners[element.ID] = element.CustomData.FrameID
		}
		if element.CustomData.SemanticParentElementID != "" {
			parents[element.ID] = element.CustomData.SemanticParentElementID
		}
	}
	for elementID := range parents {
		visited := map[string]bool{}
		parentID := elementID
		for parentID != "" && !visited[parentID] {
			visited[parentID] = true
			if frameID := owners[parentID]; frameID != "" {
				owners[elementID] = frameID
				break
			}
			parentID = parents[parentID]
		}
	}
	return owners
}

func elementBelongsToPageV1EnginePlanDocument(element entity.Element, frame *entity.Element, owners map[string]string) bool {
	frameID := strings.TrimSpace(frame.CustomData.FrameID)
	if element.CustomData != nil {
		if explicit := strings.TrimSpace(element.CustomData.FrameID); explicit != "" {
			return explicit == frameID
		}
	}
	if owner := strings.TrimSpace(owners[element.ID]); owner != "" {
		return owner == frameID
	}
	// Generated icon, text, connector, and mask elements do not all have a
	// semantic parent. Their centre is nevertheless page-local because cross-
	// frame connections are already represented by two independent stubs.
	centreX := element.X + element.Width/2
	centreY := element.Y + element.Height/2
	const epsilon = 1e-7
	return centreX >= frame.X-epsilon && centreX <= frame.X+frame.Width+epsilon &&
		centreY >= frame.Y-epsilon && centreY <= frame.Y+frame.Height+epsilon
}

func renumberPageConnectorsV1EnginePlanDocument(plan *entity.Plan, start int) int {
	if plan == nil || len(plan.ConnectorLegend) == 0 {
		return start
	}
	replacements := make(map[string]string, len(plan.ConnectorLegend))
	for legendIndex := range plan.ConnectorLegend {
		oldID := plan.ConnectorLegend[legendIndex].ID
		newID := fmt.Sprintf("L%02d", start)
		start++
		replacements[oldID] = newID
		plan.ConnectorLegend[legendIndex].ID = newID
	}
	for opIndex := range plan.Ops {
		op := &plan.Ops[opIndex]
		if op.Kind != "text" || op.TextLayout == nil || op.TextLayout.Role != entity.TextRoleConnectorLabel {
			continue
		}
		if replacement := replacements[op.Text]; replacement != "" {
			op.Text = replacement
		}
	}
	return start
}
