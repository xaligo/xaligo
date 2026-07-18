package engine

import (
	"fmt"
	"math"
	"strings"

	"github.com/xaligo/xaligo/internal/entity"
	"github.com/xaligo/xaligo/internal/share"
)

func validateFrameMetadataConnectionAnchorsV1EngineLayoutFrameMetadata(root *entity.Box, documentRoot *entity.Node) error {
	if root == nil || documentRoot == nil {
		return nil
	}
	frames := map[string]*entity.Box{}
	if root.Tag == "frames" {
		for _, frame := range root.Children {
			if frame == nil || frame.Tag != "frame" {
				continue
			}
			if frameID := strings.TrimSpace(frame.Attrs["id"]); frameID != "" {
				frames[frameID] = frame
			}
		}
	}
	for _, connection := range CollectConnectionNodesV1EngineSceneConnection(documentRoot) {
		if connection == nil || connection.Attr(internalConnectionCrossFrameAttrV1EngineParseDocument) != "true" {
			continue
		}
		for _, endpoint := range []string{"src", "dst"} {
			anchor, explicit := connectionFrameAnchorV1EngineSceneConnectionRoute(connection, endpoint)
			if !explicit {
				continue
			}
			frameID := strings.TrimSpace(connection.Attr(internalConnectionSrcFrameAttrV1EngineParseDocument))
			if endpoint == "dst" {
				frameID = strings.TrimSpace(connection.Attr(internalConnectionDstFrameAttrV1EngineParseDocument))
			}
			frame := frames[frameID]
			if frame == nil || frame.FrameMetadata == nil || frame.FrameMetadata.ReservedW <= 0 || frame.FrameMetadata.ReservedH <= 0 {
				continue
			}
			attribute := endpoint + "-frame-side"
			if strings.TrimSpace(connection.Attr(endpoint+"-frame-anchor")) != "" {
				attribute = endpoint + "-frame-anchor"
			}
			metadata := frame.FrameMetadata
			if string(anchor.side) == metadata.Position {
				return &entity.ParseError{Position: connection.Position, Err: fmt.Errorf("<connection %s=%q> selects frame %q metadata reservation edge %q", attribute, connection.Attr(attribute), frameID, metadata.Position)}
			}
			if !anchor.hasSlot || anchor.side != sideLeftV1EngineRouteTypes && anchor.side != sideRightV1EngineRouteTypes {
				continue
			}
			position := (float64(anchor.slot) + 0.5) / 5.0
			terminalY := frame.Y + frame.H*position
			reservedBottom := metadata.ReservedY + metadata.ReservedH
			if terminalY >= metadata.ReservedY-geometryEpsilonV1EngineLayoutValidation && terminalY <= reservedBottom+geometryEpsilonV1EngineLayoutValidation {
				return &entity.ParseError{Position: connection.Position, Err: fmt.Errorf("<connection %s=%q> places the frame terminal inside frame %q metadata reservation", attribute, connection.Attr(attribute), frameID)}
			}
		}
	}
	return nil
}

const (
	defaultFrameMetadataFontFamilyV1EngineLayoutFrameMetadata      = "virgil"
	defaultFrameMetadataFontSizeV1EngineLayoutFrameMetadata        = 12.0
	defaultFrameMetadataColorV1EngineLayoutFrameMetadata           = "#64748b"
	defaultFrameMetadataBackgroundColorV1EngineLayoutFrameMetadata = "transparent"
	defaultFrameMetadataKeyBackgroundV1EngineLayoutFrameMetadata   = "#f8fafc"
	defaultFrameMetadataBorderColorV1EngineLayoutFrameMetadata     = "#cbd5e1"
	defaultFrameMetadataGapV1EngineLayoutFrameMetadata             = 8.0
	defaultFrameMetadataRowGapV1EngineLayoutFrameMetadata          = 4.0
	defaultFrameMetadataContentGapV1EngineLayoutFrameMetadata      = 8.0
	minimumFrameMetadataCellWidthV1EngineLayoutFrameMetadata       = 1.0
)

type frameMetadataTagSpecV1EngineLayoutFrameMetadata struct {
	node          *entity.Node
	key           string
	value         string
	width         float64
	keyWidth      float64
	widthFixed    bool
	keyWidthFixed bool
	breakBefore   bool
}

type frameMetadataRowV1EngineLayoutFrameMetadata struct {
	start int
	end   int
	width float64
}

func layoutFrameMetadataV1EngineLayoutFrameMetadata(node *entity.Node, target *entity.Box, x, y, w, h float64) (float64, float64, float64, float64, error) {
	if node == nil || target == nil || node.Tag != "frame" {
		return x, y, w, h, nil
	}
	metadataNode := directFrameMetadataNodeV1EngineLayoutFrameMetadata(node)
	title := strings.TrimSpace(node.Attr("title"))
	contentVersion := strings.TrimSpace(node.Attr(internalFrameContentVersionAttrV1EngineParseFrameMetadata))
	if metadataNode == nil && title == "" && contentVersion == "" {
		return x, y, w, h, nil
	}
	if !isPositiveFiniteV1EngineLayoutConstraints(target.W) || !isPositiveFiniteV1EngineLayoutConstraints(target.H) || !isPositiveFiniteV1EngineLayoutConstraints(w) || !isPositiveFiniteV1EngineLayoutConstraints(h) {
		return 0, 0, 0, 0, newLayoutErrorV1EngineLayoutValidation(node, "frame metadata requires positive finite frame and content areas")
	}

	metadata := resolvedFrameMetadataStyleV1EngineLayoutFrameMetadata(metadataNode)
	tagSpecs := frameMetadataTagSpecsV1EngineLayoutFrameMetadata(node, metadataNode)
	if len(tagSpecs) == 0 {
		target.FrameMetadata = metadata
		return x, y, w, h, nil
	}

	gap := defaultFrameMetadataGapV1EngineLayoutFrameMetadata
	rowGap := defaultFrameMetadataRowGapV1EngineLayoutFrameMetadata
	if metadataNode != nil {
		gap = attrFloatV1EngineLayoutAttributes(metadataNode.Attr("gap"), gap)
		rowGap = attrFloatV1EngineLayoutAttributes(metadataNode.Attr("row-gap"), rowGap)
	}
	tagHeight := math.Ceil(metadata.FontSize*1.2) + 4
	rows := make([]frameMetadataRowV1EngineLayoutFrameMetadata, 0, 1)
	rowStart := 0
	rowWidth := 0.0
	finishRow := func() {
		if rowStart >= len(metadata.Tags) {
			return
		}
		rows = append(rows, frameMetadataRowV1EngineLayoutFrameMetadata{start: rowStart, end: len(metadata.Tags), width: rowWidth})
		rowStart = len(metadata.Tags)
		rowWidth = 0
	}
	for _, spec := range tagSpecs {
		tagWidth, keyWidth, err := resolveFrameMetadataTagWidthV1EngineLayoutFrameMetadata(spec, metadata.FontSize, target.W)
		if err != nil {
			return 0, 0, 0, 0, err
		}
		if spec.breakBefore && rowStart < len(metadata.Tags) {
			finishRow()
		}
		requiredWidth := tagWidth
		if rowStart < len(metadata.Tags) {
			requiredWidth = rowWidth + gap + tagWidth
		}
		if rowStart < len(metadata.Tags) && requiredWidth > target.W+geometryEpsilonV1EngineLayoutValidation {
			finishRow()
			requiredWidth = tagWidth
		}
		metadata.Tags = append(metadata.Tags, entity.FrameMetadataTag{
			Key: spec.key, Value: spec.value,
			W: tagWidth, H: tagHeight, KeyW: keyWidth,
			DiffStatus: frameMetadataTagDiffStatusV1EngineLayoutFrameMetadata(
				spec.node,
				metadataNode,
			),
		})
		rowWidth = requiredWidth
	}
	finishRow()
	bandHeight := float64(len(rows))*tagHeight + float64(len(rows)-1)*rowGap
	if bandHeight > target.H+geometryEpsilonV1EngineLayoutValidation {
		return 0, 0, 0, 0, newLayoutErrorV1EngineLayoutValidation(node, "frame metadata band %.6g exceeds the frame height %.6g", bandHeight, target.H)
	}
	bandY := target.Y
	if metadata.Position == "bottom" {
		bandY = target.Y + target.H - bandHeight
	}
	for rowIndex, row := range rows {
		offsetX := 0.0
		switch metadata.Align {
		case "center":
			offsetX = (target.W - row.width) / 2
		case "right":
			offsetX = target.W - row.width
		}
		cursorX := target.X + math.Max(0, offsetX)
		rowY := bandY + float64(rowIndex)*(tagHeight+rowGap)
		for tagIndex := row.start; tagIndex < row.end; tagIndex++ {
			metadata.Tags[tagIndex].X = cursorX
			metadata.Tags[tagIndex].Y = rowY
			cursorX += metadata.Tags[tagIndex].W + gap
		}
	}
	contentBottom := y + h
	if metadata.Position == "bottom" {
		maximumContentBottom := bandY - defaultFrameMetadataContentGapV1EngineLayoutFrameMetadata
		if maximumContentBottom < contentBottom {
			contentBottom = maximumContentBottom
		}
	} else {
		minimumContentY := bandY + bandHeight + defaultFrameMetadataContentGapV1EngineLayoutFrameMetadata
		if minimumContentY > y {
			y = minimumContentY
		}
	}
	h = contentBottom - y
	if !isPositiveFiniteV1EngineLayoutConstraints(h) {
		return 0, 0, 0, 0, newLayoutErrorV1EngineLayoutValidation(node, "frame metadata reserved strip %.6g leaves no positive content height", bandHeight)
	}
	target.FrameMetadata = metadata
	return x, y, w, h, nil
}

func finalizeFrameMetadataReservedStripV1EngineLayoutFrameMetadata(target *entity.Box, contentY, contentH float64) {
	if target == nil || target.FrameMetadata == nil || len(target.FrameMetadata.Tags) == 0 {
		return
	}
	metadata := target.FrameMetadata
	metadata.ReservedX = target.X
	metadata.ReservedW = target.W
	if metadata.Position == "bottom" {
		metadata.ReservedY = contentY + contentH
		metadata.ReservedH = math.Max(0, target.Y+target.H-metadata.ReservedY)
		return
	}
	metadata.ReservedY = target.Y
	metadata.ReservedH = math.Max(0, contentY-target.Y)
}

func frameMetadataTagDiffStatusV1EngineLayoutFrameMetadata(source, metadata *entity.Node) string {
	if source != nil && source.Tag == "entry" {
		if status := validFrameMetadataDiffStatusV1EngineLayoutFrameMetadata(source.Attr(diffStatusAttrV1EngineDiffDocument)); status != "" {
			return status
		}
	}
	if metadata != nil {
		return validFrameMetadataDiffStatusV1EngineLayoutFrameMetadata(metadata.Attr(diffStatusAttrV1EngineDiffDocument))
	}
	return ""
}

func validFrameMetadataDiffStatusV1EngineLayoutFrameMetadata(status string) string {
	status = strings.TrimSpace(status)
	if status == string(entity.StructuralChangeAdded) || status == string(entity.StructuralChangeRemoved) {
		return status
	}
	return ""
}

func directFrameMetadataNodeV1EngineLayoutFrameMetadata(frame *entity.Node) *entity.Node {
	if frame == nil {
		return nil
	}
	for _, child := range frame.Children {
		if child.Tag == "metadata" {
			return child
		}
	}
	return nil
}

func resolvedFrameMetadataStyleV1EngineLayoutFrameMetadata(node *entity.Node) *entity.FrameMetadata {
	metadata := &entity.FrameMetadata{
		Position:           "top",
		Align:              "left",
		FontFamily:         defaultFrameMetadataFontFamilyV1EngineLayoutFrameMetadata,
		FontSize:           defaultFrameMetadataFontSizeV1EngineLayoutFrameMetadata,
		Color:              defaultFrameMetadataColorV1EngineLayoutFrameMetadata,
		KeyColor:           defaultFrameMetadataColorV1EngineLayoutFrameMetadata,
		BackgroundColor:    defaultFrameMetadataBackgroundColorV1EngineLayoutFrameMetadata,
		KeyBackgroundColor: defaultFrameMetadataKeyBackgroundV1EngineLayoutFrameMetadata,
		BorderColor:        defaultFrameMetadataBorderColorV1EngineLayoutFrameMetadata,
	}
	if node == nil {
		return metadata
	}
	if value := strings.TrimSpace(node.Attr("position")); value != "" {
		metadata.Position = value
	}
	if value := strings.TrimSpace(node.Attr("align")); value != "" {
		metadata.Align = value
	}
	if value := strings.TrimSpace(node.Attr("font-family")); value != "" {
		metadata.FontFamily = value
	}
	metadata.FontSize = attrFloatV1EngineLayoutAttributes(node.Attr("font-size"), metadata.FontSize)
	if value := strings.TrimSpace(node.Attr("color")); value != "" {
		metadata.Color = value
		metadata.KeyColor = value
	}
	if value := strings.TrimSpace(node.Attr("key-color")); value != "" {
		metadata.KeyColor = value
	}
	if value := strings.TrimSpace(node.Attr("background-color")); value != "" {
		metadata.BackgroundColor = value
	}
	if value := strings.TrimSpace(node.Attr("key-background-color")); value != "" {
		metadata.KeyBackgroundColor = value
	}
	if value := strings.TrimSpace(node.Attr("border-color")); value != "" {
		metadata.BorderColor = value
	}
	return metadata
}

func frameMetadataTagSpecsV1EngineLayoutFrameMetadata(frame, metadata *entity.Node) []frameMetadataTagSpecV1EngineLayoutFrameMetadata {
	defaultWidth, defaultWidthFixed := explicitFrameMetadataSizeV1EngineLayoutFrameMetadata(metadata, "width")
	defaultKeyWidth, defaultKeyWidthFixed := explicitFrameMetadataSizeV1EngineLayoutFrameMetadata(metadata, "key-width")
	specs := make([]frameMetadataTagSpecV1EngineLayoutFrameMetadata, 0, 3)
	appendTag := func(source *entity.Node, key, value string, allowSizeOverride bool) {
		value = strings.TrimSpace(value)
		if value == "" {
			return
		}
		width, widthFixed := defaultWidth, defaultWidthFixed
		keyWidth, keyWidthFixed := defaultKeyWidth, defaultKeyWidthFixed
		breakBefore := false
		if source != nil && allowSizeOverride {
			if resolved, fixed := explicitFrameMetadataSizeV1EngineLayoutFrameMetadata(source, "width"); fixed {
				width, widthFixed = resolved, true
			}
			if resolved, fixed := explicitFrameMetadataSizeV1EngineLayoutFrameMetadata(source, "key-width"); fixed {
				keyWidth, keyWidthFixed = resolved, true
			}
			breakBefore = source.Attr("break-before") == "true"
		}
		specs = append(specs, frameMetadataTagSpecV1EngineLayoutFrameMetadata{
			node: source, key: key, value: value,
			width: width, keyWidth: keyWidth,
			widthFixed: widthFixed, keyWidthFixed: keyWidthFixed,
			breakBefore: breakBefore,
		})
	}
	appendTag(frame, "id", frame.Attr("id"), false)
	appendTag(frame, "title", frame.Attr("title"), false)
	appendTag(frame, "version", frame.Attr(internalFrameContentVersionAttrV1EngineParseFrameMetadata), false)
	if metadata != nil {
		for _, entry := range metadata.Children {
			appendTag(entry, entry.Attr("key"), entry.Attr("value"), true)
		}
	}
	return specs
}

func explicitFrameMetadataSizeV1EngineLayoutFrameMetadata(node *entity.Node, attribute string) (float64, bool) {
	if node == nil {
		return 0, false
	}
	value := strings.TrimSpace(node.Attr(attribute))
	if value == "" {
		return 0, false
	}
	return attrFloatV1EngineLayoutAttributes(value, 0), true
}

func resolveFrameMetadataTagWidthV1EngineLayoutFrameMetadata(spec frameMetadataTagSpecV1EngineLayoutFrameMetadata, fontSize, availableWidth float64) (float64, float64, error) {
	autoKeyWidth := math.Ceil(share.PresentationTextWidth(spec.key, fontSize, false)) + 8
	autoValueWidth := math.Ceil(share.PresentationTextWidth(spec.value, fontSize, false)) + 8
	tagWidth := autoKeyWidth + autoValueWidth
	keyWidth := autoKeyWidth
	if spec.keyWidthFixed {
		keyWidth = spec.keyWidth
		tagWidth = keyWidth + autoValueWidth
	}
	if spec.widthFixed {
		tagWidth = spec.width
		if !spec.keyWidthFixed {
			keyWidth = tagWidth * autoKeyWidth / (autoKeyWidth + autoValueWidth)
		}
	}
	if spec.widthFixed && tagWidth > availableWidth+geometryEpsilonV1EngineLayoutValidation {
		return 0, 0, newLayoutErrorV1EngineLayoutValidation(spec.node, "metadata width %.6g exceeds the available frame width %.6g", tagWidth, availableWidth)
	}
	if !spec.widthFixed && tagWidth > availableWidth {
		tagWidth = availableWidth
		if !spec.keyWidthFixed {
			keyWidth = tagWidth * autoKeyWidth / (autoKeyWidth + autoValueWidth)
		}
	}
	if !isPositiveFiniteV1EngineLayoutConstraints(tagWidth) || keyWidth < minimumFrameMetadataCellWidthV1EngineLayoutFrameMetadata || tagWidth-keyWidth < minimumFrameMetadataCellWidthV1EngineLayoutFrameMetadata {
		return 0, 0, newLayoutErrorV1EngineLayoutValidation(spec.node, "metadata width %.6g and key-width %.6g must leave positive key and value cells", tagWidth, keyWidth)
	}
	return tagWidth, keyWidth, nil
}
