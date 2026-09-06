package v2

//go:generate env GOCACHE=/tmp/xaligo-go-generate-cache go run ../../../scripts/tool/gen_engine_abi.go ../../..

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
	"unicode/utf8"

	enginebridge "github.com/xaligo/xaligo/external/engine"
	"github.com/xaligo/xaligo/internal/entity"
)

type EngineUsecase interface {
	Resolve(context.Context, entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error)
	RenderSVG(context.Context, entity.EngineDocumentSpec) ([]byte, error)
	NormalizeSVG(context.Context, []byte) (entity.EngineSVG, error)
}

type engineUsecase struct{}

func NewEngineUsecase() EngineUsecase {
	return &engineUsecase{}
}

const (
	engineABIVersion            = uint16(5)
	engineOperationLayout       = byte(1)
	engineOperationSVG          = byte(2)
	engineOperationNormalizeSVG = byte(3)
	engineStatusOK              = byte(0)
	engineStatusError           = byte(1)
	engineMaxElements           = 10_000
	engineMaxDepth              = 128
	engineMaxRequest            = 16 * 1024 * 1024
	engineMaxResponse           = 32 * 1024 * 1024
	engineMaxSVG                = 2 * 1024 * 1024
	engineMaxRoutePoints        = 4096
)

var (
	engineRequestMagic  = [4]byte{'X', 'L', 'E', '2'}
	engineResponseMagic = [4]byte{'X', 'L', 'R', '2'}
)

type flatEngineElement struct {
	parent  int32
	element entity.EngineElementSpec
}

func (rcvr *engineUsecase) Resolve(ctx context.Context, spec entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error) {
	response, err := rcvr.execute(ctx, engineOperationLayout, spec)
	if err != nil {
		return entity.EngineResolvedDocument{}, err
	}
	resolved, err := decodeEngineLayout(response, spec)
	return resolved, enrichEngineDiagnostic(err, spec)
}

func (rcvr *engineUsecase) RenderSVG(ctx context.Context, spec entity.EngineDocumentSpec) ([]byte, error) {
	response, err := rcvr.execute(ctx, engineOperationSVG, spec)
	if err != nil {
		return nil, err
	}
	data, err := decodeEngineBytes(response, engineOperationSVG)
	return data, enrichEngineDiagnostic(err, spec)
}

func (rcvr *engineUsecase) NormalizeSVG(ctx context.Context, input []byte) (entity.EngineSVG, error) {
	request, err := encodeEngineSVGRequest(input)
	if err != nil {
		return entity.EngineSVG{}, err
	}
	response, err := rcvr.executeRequest(ctx, request)
	if err != nil {
		return entity.EngineSVG{}, err
	}
	return decodeEngineSVG(response)
}

func (rcvr *engineUsecase) execute(ctx context.Context, operation byte, spec entity.EngineDocumentSpec) ([]byte, error) {
	request, err := encodeEngineRequest(operation, spec)
	if err != nil {
		return nil, err
	}
	return rcvr.executeRequest(ctx, request)
}

func (rcvr *engineUsecase) executeRequest(ctx context.Context, request []byte) ([]byte, error) {
	if err := checkEngineContext(ctx); err != nil {
		return nil, err
	}
	if len(request) > engineMaxRequest {
		return nil, fmt.Errorf("Rust engine request size %d exceeds %d", len(request), engineMaxRequest)
	}
	if !enginebridge.Available() {
		return nil, enginebridge.ErrUnavailable
	}
	if version := enginebridge.ABIVersion(); version != uint32(engineABIVersion) {
		return nil, fmt.Errorf("Rust engine ABI version %d does not match Go ABI version %d", version, engineABIVersion)
	}
	response, err := enginebridge.ProcessContext(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("invoke Rust engine: %w", err)
	}
	if err := checkEngineContext(ctx); err != nil {
		return nil, err
	}
	if len(response) > engineMaxResponse {
		return nil, fmt.Errorf("Rust engine response size %d exceeds %d", len(response), engineMaxResponse)
	}
	return response, nil
}

func encodeEngineSVGRequest(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("SVG input must not be empty")
	}
	if len(input) > engineMaxSVG {
		return nil, fmt.Errorf("SVG input size %d exceeds %d", len(input), engineMaxSVG)
	}
	var output bytes.Buffer
	output.Grow(12 + len(input))
	output.Write(engineRequestMagic[:])
	writeUint16(&output, engineABIVersion)
	output.WriteByte(engineOperationNormalizeSVG)
	output.WriteByte(0)
	writeUint32(&output, uint32(len(input)))
	output.Write(input)
	return output.Bytes(), nil
}

func checkEngineContext(ctx context.Context) error {
	if ctx == nil {
		return nil
	}
	return ctx.Err()
}

func encodeEngineRequest(operation byte, spec entity.EngineDocumentSpec) ([]byte, error) {
	if operation != engineOperationLayout && operation != engineOperationSVG {
		return nil, fmt.Errorf("unsupported engine operation %d", operation)
	}
	layout, err := encodeDocumentLayout(spec)
	if err != nil {
		return nil, err
	}
	overflow, err := encodeEngineOverflow(spec.Overflow)
	if err != nil {
		return nil, err
	}
	elements, err := flattenEngineElements(spec.Elements)
	if err != nil {
		return nil, err
	}

	var documentFlags byte
	var padding [4]float64
	setDocumentNumber(&documentFlags, &padding, 0, spec.Padding.Top)
	setDocumentNumber(&documentFlags, &padding, 1, spec.Padding.Right)
	setDocumentNumber(&documentFlags, &padding, 2, spec.Padding.Bottom)
	setDocumentNumber(&documentFlags, &padding, 3, spec.Padding.Left)
	var columns uint16
	if spec.Columns != nil {
		documentFlags |= 1 << 4
		columns = *spec.Columns
	}

	var output bytes.Buffer
	output.Grow(72 + len(elements)*420)
	output.Write(engineRequestMagic[:])
	writeUint16(&output, engineABIVersion)
	output.WriteByte(operation)
	output.WriteByte(layout)
	writeFloat64(&output, spec.Width)
	writeFloat64(&output, spec.Height)
	writeFloat64(&output, spec.Gap)
	output.WriteByte(documentFlags)
	output.WriteByte(overflow)
	writeUint16(&output, columns)
	for _, value := range padding {
		writeFloat64(&output, value)
	}
	writeUint32(&output, uint32(len(elements)))
	for _, flat := range elements {
		if err := encodeEngineElement(&output, flat); err != nil {
			return nil, err
		}
		if output.Len() > engineMaxRequest {
			return nil, fmt.Errorf("engine request size %d exceeds %d", output.Len(), engineMaxRequest)
		}
	}
	return output.Bytes(), nil
}

func flattenEngineElements(roots []entity.EngineElementSpec) ([]flatEngineElement, error) {
	result := make([]flatEngineElement, 0, len(roots))
	var visit func([]entity.EngineElementSpec, int32, int) error
	visit = func(elements []entity.EngineElementSpec, parent int32, depth int) error {
		if depth > engineMaxDepth {
			return fmt.Errorf("engine element depth exceeds %d", engineMaxDepth)
		}
		for _, element := range elements {
			if len(result) >= engineMaxElements {
				return fmt.Errorf("engine element count exceeds %d", engineMaxElements)
			}
			index := int32(len(result))
			result = append(result, flatEngineElement{parent: parent, element: element})
			if err := visit(element.Children, index, depth+1); err != nil {
				return err
			}
		}
		return nil
	}
	if err := visit(roots, -1, 1); err != nil {
		return nil, err
	}
	return result, nil
}

func encodeEngineElement(output *bytes.Buffer, flat flatEngineElement) error {
	element := flat.element
	concept, err := encodeEngineConcept(element)
	if err != nil {
		return fmt.Errorf("engine element %q: %w", element.ID, err)
	}
	layout, err := encodeEngineLayout(element.Layout)
	if err != nil {
		return fmt.Errorf("engine element %q: %w", element.ID, err)
	}
	overflow, err := encodeEngineOverflow(element.Overflow)
	if err != nil {
		return fmt.Errorf("engine element %q: %w", element.ID, err)
	}
	shape, err := encodeEngineShape(element.Visual.Shape)
	if err != nil {
		return fmt.Errorf("engine element %q: %w", element.ID, err)
	}
	align, err := encodeEngineAlignment(element.Align)
	if err != nil {
		return fmt.Errorf("engine element %q: %w", element.ID, err)
	}
	justify, err := encodeEngineJustification(element.Justify)
	if err != nil {
		return fmt.Errorf("engine element %q: %w", element.ID, err)
	}

	var numbers [engineNumberFieldCount]float64
	var numberFlags uint64
	setEngineNumber(&numberFlags, &numbers, engineNumberX, element.X)
	setEngineNumber(&numberFlags, &numbers, engineNumberY, element.Y)
	setEngineNumber(&numberFlags, &numbers, engineNumberWidth, element.Width)
	setEngineNumber(&numberFlags, &numbers, engineNumberHeight, element.Height)
	setEngineNumber(&numberFlags, &numbers, engineNumberIntrinsicWidth, element.IntrinsicWidth)
	setEngineNumber(&numberFlags, &numbers, engineNumberIntrinsicHeight, element.IntrinsicHeight)
	setEngineNumber(&numberFlags, &numbers, engineNumberMinWidth, element.MinWidth)
	setEngineNumber(&numberFlags, &numbers, engineNumberMaxWidth, element.MaxWidth)
	setEngineNumber(&numberFlags, &numbers, engineNumberMinHeight, element.MinHeight)
	setEngineNumber(&numberFlags, &numbers, engineNumberMaxHeight, element.MaxHeight)
	setEngineNumber(&numberFlags, &numbers, engineNumberOffsetX, element.OffsetX)
	setEngineNumber(&numberFlags, &numbers, engineNumberOffsetY, element.OffsetY)
	setEngineNumber(&numberFlags, &numbers, engineNumberWeight, element.Weight)
	setEngineNumber(&numberFlags, &numbers, engineNumberGap, element.Gap)
	setEngineInsets(&numberFlags, &numbers, engineNumberMarginTop, element.Margin)
	setEngineInsets(&numberFlags, &numbers, engineNumberPaddingTop, element.Padding)
	setEngineNumber(&numberFlags, &numbers, engineNumberStrokeWidth, element.Visual.StrokeWidth)
	setEngineNumber(&numberFlags, &numbers, engineNumberCornerRadius, element.Visual.CornerRadius)
	setEngineNumber(&numberFlags, &numbers, engineNumberOpacity, element.Visual.Opacity)

	var booleansPresent uint16
	var booleansValue uint16
	setEngineBool(&booleansPresent, &booleansValue, engineBoolVisible, element.Visual.Visible)
	var layer int32
	if element.Visual.Layer != nil {
		booleansPresent |= 1 << engineBoolLayer
		layer = *element.Visual.Layer
	}

	var strings [engineStringFieldCount]string
	strings[0] = element.ID
	if element.Text != nil {
		strings[1] = element.Text.Value
		strings[2] = element.Text.FontFamily
		strings[3] = element.Text.Color
		strings[4] = element.Text.Role
		setEngineNumber(&numberFlags, &numbers, engineNumberFontSize, element.Text.FontSize)
		setEngineNumber(&numberFlags, &numbers, engineNumberLineHeight, element.Text.LineHeight)
		setEngineInsets(&numberFlags, &numbers, engineNumberTextPaddingTop, element.Text.Padding)
		setEngineBool(&booleansPresent, &booleansValue, engineBoolTextWrap, element.Text.Wrap)
		setEngineBool(&booleansPresent, &booleansValue, engineBoolTextFit, element.Text.Fit)
		setEngineBool(&booleansPresent, &booleansValue, engineBoolTextClip, element.Text.Clip)
	}
	var missingIcon byte
	if element.Icon != nil {
		strings[5] = element.Icon.Ref
		strings[6] = element.Icon.FallbackRef
		strings[7] = element.Icon.Color
		setEngineNumber(&numberFlags, &numbers, engineNumberIconWidth, element.Icon.Width)
		setEngineNumber(&numberFlags, &numbers, engineNumberIconHeight, element.Icon.Height)
		setEngineNumber(&numberFlags, &numbers, engineNumberIconScale, element.Icon.Scale)
		setEngineNumber(&numberFlags, &numbers, engineNumberIconOffsetX, element.Icon.OffsetX)
		setEngineNumber(&numberFlags, &numbers, engineNumberIconOffsetY, element.Icon.OffsetY)
		missingIcon, err = encodeEngineMissingIcon(element.Icon.MissingPolicy)
		if err != nil {
			return fmt.Errorf("engine element %q: %w", element.ID, err)
		}
	}
	strings[8] = element.Visual.Fill
	strings[9] = element.Visual.Stroke
	if component := element.AWS; component != nil {
		strings[engineStringAwsKind] = component.Kind
		strings[engineStringAwsDomain] = component.Domain
		strings[engineStringAwsProtocol] = component.Protocol
		strings[engineStringAwsMutualTls] = component.MutualTLS
		strings[engineStringAwsCertificate] = component.Certificate
		strings[engineStringAwsTrustStore] = component.TrustStore
		strings[engineStringAwsTargetGroup] = component.TargetGroup
		strings[engineStringAwsDetailLevel] = component.DetailLevel
		strings[engineStringAwsShow] = component.Show
		strings[engineStringAwsHide] = component.Hide
		strings[engineStringAwsType] = component.Type
		strings[engineStringAwsName] = component.Name
		strings[engineStringAwsValue] = component.Value
		strings[engineStringAwsAux] = component.Aux
		strings[engineStringAwsOrder] = component.Order
		if component.Port != nil {
			port := float64(*component.Port)
			setEngineNumber(&numberFlags, &numbers, engineNumberAwsPort, &port)
		}
		setEngineBool(&booleansPresent, &booleansValue, engineBoolAwsBackendTls, component.BackendTLS)
		setEngineBool(&booleansPresent, &booleansValue, engineBoolAwsBackendMtls, component.BackendMTLS)
		setEngineBool(&booleansPresent, &booleansValue, engineBoolAwsShowTitle, component.ShowTitle)
	}

	var side byte
	if element.Port != nil {
		side, err = encodeEngineSide(element.Port.Side)
		if err != nil {
			return fmt.Errorf("engine element %q: %w", element.ID, err)
		}
		strings[10] = element.Port.Label
		setEngineNumber(&numberFlags, &numbers, engineNumberPortAnchor, element.Port.Anchor)
		setEngineNumber(&numberFlags, &numbers, engineNumberPortOffset, element.Port.Offset)
		setEngineNumber(&numberFlags, &numbers, engineNumberPortSize, element.Port.Size)
		setEngineBool(&booleansPresent, &booleansValue, engineBoolPortVisible, element.Port.Visible)
	}

	var routing, sourceSide, targetSide, lineStyle, sourceDecoration, targetDecoration byte
	if element.Line != nil {
		strings[11] = element.Line.Source
		strings[12] = element.Line.Target
		strings[13] = element.Line.Label
		routing, err = encodeEngineRouting(element.Line.Routing)
		if err != nil {
			return fmt.Errorf("engine element %q: %w", element.ID, err)
		}
		sourceSide, err = encodeEngineSide(element.Line.SourceSide)
		if err != nil {
			return fmt.Errorf("engine element %q source: %w", element.ID, err)
		}
		targetSide, err = encodeEngineSide(element.Line.TargetSide)
		if err != nil {
			return fmt.Errorf("engine element %q target: %w", element.ID, err)
		}
		lineStyle, err = encodeEngineLineStyle(element.Line.Style)
		if err != nil {
			return fmt.Errorf("engine element %q: %w", element.ID, err)
		}
		sourceDecoration, err = encodeEngineDecoration(element.Line.SourceDecoration)
		if err != nil {
			return fmt.Errorf("engine element %q source: %w", element.ID, err)
		}
		targetDecoration, err = encodeEngineDecoration(element.Line.TargetDecoration)
		if err != nil {
			return fmt.Errorf("engine element %q target: %w", element.ID, err)
		}
		setEngineNumber(&numberFlags, &numbers, engineNumberSourceAnchor, element.Line.SourceAnchor)
		setEngineNumber(&numberFlags, &numbers, engineNumberTargetAnchor, element.Line.TargetAnchor)
		setEngineNumber(&numberFlags, &numbers, engineNumberObstacleMargin, element.Line.ObstacleMargin)
		setEngineNumber(&numberFlags, &numbers, engineNumberLabelPosition, element.Line.LabelPosition)
	}

	var columns, columnSpan, rowSpan uint16
	if element.Columns != nil {
		booleansPresent |= 1 << engineBoolColumns
		columns = *element.Columns
	}
	if element.ColumnSpan != nil {
		booleansPresent |= 1 << engineBoolColumnSpan
		columnSpan = *element.ColumnSpan
	}
	if element.RowSpan != nil {
		booleansPresent |= 1 << engineBoolRowSpan
		rowSpan = *element.RowSpan
	}

	writeInt32(output, flat.parent)
	writeUint64(output, numberFlags)
	writeUint16(output, booleansPresent)
	writeUint16(output, booleansValue)
	writeUint16(output, columns)
	writeUint16(output, columnSpan)
	writeUint16(output, rowSpan)
	writeUint16(output, 0)
	output.Write([]byte{
		concept, layout, overflow, shape, align, justify, side, routing,
		sourceSide, targetSide, lineStyle, sourceDecoration, targetDecoration, missingIcon,
	})
	writeInt32(output, layer)
	for _, value := range numbers {
		writeFloat64(output, value)
	}
	for _, value := range strings {
		if len(value) > math.MaxUint16 {
			return fmt.Errorf("engine element %q string field exceeds %d UTF-8 bytes", element.ID, math.MaxUint16)
		}
		writeUint16(output, uint16(len(value)))
	}
	for _, value := range strings {
		output.WriteString(value)
	}
	return nil
}

func setDocumentNumber(flags *byte, values *[4]float64, index int, value *float64) {
	if value == nil {
		return
	}
	*flags |= 1 << index
	values[index] = *value
}

func setEngineInsets(flags *uint64, values *[engineNumberFieldCount]float64, start int, insets entity.EngineInsets) {
	setEngineNumber(flags, values, start, insets.Top)
	setEngineNumber(flags, values, start+1, insets.Right)
	setEngineNumber(flags, values, start+2, insets.Bottom)
	setEngineNumber(flags, values, start+3, insets.Left)
}

func setEngineNumber(flags *uint64, values *[engineNumberFieldCount]float64, index int, value *float64) {
	if value == nil {
		return
	}
	*flags |= uint64(1) << index
	values[index] = *value
}

func setEngineBool(present *uint16, values *uint16, index int, value *bool) {
	if value == nil {
		return
	}
	*present |= 1 << index
	if *value {
		*values |= 1 << index
	}
}

func encodeDocumentLayout(spec entity.EngineDocumentSpec) (byte, error) {
	if spec.Layout != "" {
		return encodeEngineLayout(spec.Layout)
	}
	switch spec.Direction {
	case "", entity.EngineDirectionVertical:
		return 1, nil
	case entity.EngineDirectionHorizontal:
		return 2, nil
	default:
		return 0, fmt.Errorf("unsupported engine direction %q", spec.Direction)
	}
}

func encodeEngineConcept(element entity.EngineElementSpec) (byte, error) {
	concept := element.Concept
	if concept == "" {
		switch {
		case element.Line != nil:
			concept = entity.EngineConceptLine
		case element.Port != nil:
			concept = entity.EngineConceptPort
		case element.Text != nil:
			concept = entity.EngineConceptText
		default:
			concept = entity.EngineConceptItem
		}
	}
	switch concept {
	case entity.EngineConceptFrame:
		return 1, nil
	case entity.EngineConceptGroup:
		return 2, nil
	case entity.EngineConceptCapture:
		return 3, nil
	case entity.EngineConceptItem:
		return 4, nil
	case entity.EngineConceptPort:
		return 5, nil
	case entity.EngineConceptLine:
		return 6, nil
	case entity.EngineConceptText:
		return 7, nil
	case entity.EngineConceptSpacer:
		return 8, nil
	default:
		return 0, fmt.Errorf("unsupported concept %q", concept)
	}
}

func encodeEngineLayout(value entity.EngineLayoutPolicy) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineLayoutVertical:
		return 1, nil
	case entity.EngineLayoutHorizontal:
		return 2, nil
	case entity.EngineLayoutGrid:
		return 3, nil
	case entity.EngineLayoutAbsolute:
		return 4, nil
	case entity.EngineLayoutNone:
		return 5, nil
	case entity.EngineLayoutAdaptiveGrid:
		return 6, nil
	default:
		return 0, fmt.Errorf("unsupported layout policy %q", value)
	}
}

func encodeEngineOverflow(value entity.EngineOverflow) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineOverflowError:
		return 1, nil
	case entity.EngineOverflowVisible:
		return 2, nil
	default:
		return 0, fmt.Errorf("unsupported overflow policy %q", value)
	}
}

func encodeEngineShape(value entity.EngineShape) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineShapeRectangle:
		return 1, nil
	case entity.EngineShapeEllipse:
		return 2, nil
	case entity.EngineShapeNone:
		return 3, nil
	default:
		return 0, fmt.Errorf("unsupported shape %q", value)
	}
}

func encodeEngineAlignment(value entity.EngineAlignment) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineAlignStart:
		return 1, nil
	case entity.EngineAlignCenter:
		return 2, nil
	case entity.EngineAlignEnd:
		return 3, nil
	case entity.EngineAlignStretch:
		return 4, nil
	default:
		return 0, fmt.Errorf("unsupported alignment %q", value)
	}
}

func encodeEngineJustification(value entity.EngineJustification) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineJustifyStart:
		return 1, nil
	case entity.EngineJustifyCenter:
		return 2, nil
	case entity.EngineJustifyEnd:
		return 3, nil
	case entity.EngineJustifySpaceBetween:
		return 4, nil
	case entity.EngineJustifySpaceEvenly:
		return 5, nil
	default:
		return 0, fmt.Errorf("unsupported justification %q", value)
	}
}

func encodeEngineSide(value entity.EngineSide) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineSideAuto:
		return 1, nil
	case entity.EngineSideTop:
		return 2, nil
	case entity.EngineSideRight:
		return 3, nil
	case entity.EngineSideBottom:
		return 4, nil
	case entity.EngineSideLeft:
		return 5, nil
	default:
		return 0, fmt.Errorf("unsupported side %q", value)
	}
}

func encodeEngineRouting(value entity.EngineRoutingPolicy) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineRoutingOrthogonal:
		return 1, nil
	case entity.EngineRoutingStraight:
		return 2, nil
	default:
		return 0, fmt.Errorf("unsupported routing policy %q", value)
	}
}

func encodeEngineLineStyle(value entity.EngineLineStyle) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineLineSolid:
		return 1, nil
	case entity.EngineLineDashed:
		return 2, nil
	case entity.EngineLineDotted:
		return 3, nil
	default:
		return 0, fmt.Errorf("unsupported line style %q", value)
	}
}

func encodeEngineDecoration(value entity.EngineDecoration) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineDecorationNone:
		return 1, nil
	case entity.EngineDecorationArrow:
		return 2, nil
	case entity.EngineDecorationTriangle:
		return 3, nil
	case entity.EngineDecorationDiamond:
		return 4, nil
	case entity.EngineDecorationCircle:
		return 5, nil
	default:
		return 0, fmt.Errorf("unsupported decoration %q", value)
	}
}

func encodeEngineMissingIcon(value entity.EngineIconMissingPolicy) (byte, error) {
	switch value {
	case "":
		return 0, nil
	case entity.EngineIconMissingError:
		return 1, nil
	case entity.EngineIconMissingFallback:
		return 2, nil
	case entity.EngineIconMissingHide:
		return 3, nil
	default:
		return 0, fmt.Errorf("unsupported missing-icon policy %q", value)
	}
}

func decodeEngineLayout(input []byte, spec entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error) {
	reader, err := decodeEngineResponseHeader(input, engineOperationLayout)
	if err != nil {
		return entity.EngineResolvedDocument{}, err
	}
	count, err := readUint32(reader)
	if err != nil {
		return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine layout count: %w", err)
	}
	if count > engineMaxElements {
		return entity.EngineResolvedDocument{}, fmt.Errorf("Rust engine layout count %d exceeds %d", count, engineMaxElements)
	}
	elements := make([]entity.EngineResolvedElement, 0, count)
	for index := range int(count) {
		parent, err := readInt32(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine parent index: %w", err)
		}
		if parent < -1 || parent >= int32(index) {
			return entity.EngineResolvedDocument{}, fmt.Errorf("Rust engine parent index %d is invalid for element %d", parent, index)
		}
		conceptCode, err := reader.ReadByte()
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine concept: %w", err)
		}
		shapeCode, err := reader.ReadByte()
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine shape: %w", err)
		}
		lineStyleCode, err := reader.ReadByte()
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine line style: %w", err)
		}
		sourceDecorationCode, err := reader.ReadByte()
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine source decoration: %w", err)
		}
		targetDecorationCode, err := reader.ReadByte()
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine target decoration: %w", err)
		}
		visible, err := reader.ReadByte()
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine visibility: %w", err)
		}
		if visible > 1 {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine visibility: invalid value %d", visible)
		}
		reserved, err := readUint16(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine reserved field: %w", err)
		}
		if reserved != 0 {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine reserved field is %d", reserved)
		}
		layer, err := readInt32(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine layer: %w", err)
		}
		pointCount, err := readUint16(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine route point count: %w", err)
		}
		if pointCount > engineMaxRoutePoints {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine route point count %d exceeds %d", pointCount, engineMaxRoutePoints)
		}
		reserved, err = readUint16(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine route reserved field: %w", err)
		}
		if reserved != 0 {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine route reserved field is %d", reserved)
		}
		var stringLengths [engineResolvedStringFieldCount]uint16
		for stringIndex := range stringLengths {
			stringLengths[stringIndex], err = readUint16(reader)
			if err != nil {
				return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine string length: %w", err)
			}
		}
		var numbers [engineResolvedNumberFieldCount]float64
		for numberIndex := range numbers {
			numbers[numberIndex], err = readFloat64(reader)
			if err != nil {
				return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine number %d: %w", numberIndex, err)
			}
			if !isFinite(numbers[numberIndex]) {
				return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine number %d is non-finite", numberIndex)
			}
		}
		var strings [engineResolvedStringFieldCount]string
		for stringIndex, length := range stringLengths {
			value := make([]byte, length)
			if _, err := io.ReadFull(reader, value); err != nil {
				return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine string: %w", err)
			}
			if !utf8.Valid(value) {
				return entity.EngineResolvedDocument{}, errors.New("Rust engine returned invalid UTF-8")
			}
			strings[stringIndex] = string(value)
		}
		points := make([]entity.EnginePoint, pointCount)
		for pointIndex := range points {
			points[pointIndex].X, err = readFloat64(reader)
			if err != nil {
				return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine route x: %w", err)
			}
			if !isFinite(points[pointIndex].X) {
				return entity.EngineResolvedDocument{}, errors.New("decode Rust engine route x: non-finite value")
			}
			points[pointIndex].Y, err = readFloat64(reader)
			if err != nil {
				return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine route y: %w", err)
			}
			if !isFinite(points[pointIndex].Y) {
				return entity.EngineResolvedDocument{}, errors.New("decode Rust engine route y: non-finite value")
			}
		}
		concept, err := decodeEngineConcept(conceptCode)
		if err != nil {
			return entity.EngineResolvedDocument{}, err
		}
		shape, err := decodeEngineShape(shapeCode)
		if err != nil {
			return entity.EngineResolvedDocument{}, err
		}
		lineStyle, err := decodeEngineLineStyle(lineStyleCode)
		if err != nil {
			return entity.EngineResolvedDocument{}, err
		}
		sourceDecoration, err := decodeEngineDecoration(sourceDecorationCode)
		if err != nil {
			return entity.EngineResolvedDocument{}, err
		}
		targetDecoration, err := decodeEngineDecoration(targetDecorationCode)
		if err != nil {
			return entity.EngineResolvedDocument{}, err
		}
		parentID := ""
		if parent >= 0 {
			parentID = elements[parent].ID
		}
		elements = append(elements, entity.EngineResolvedElement{
			ID:       strings[0],
			ParentID: parentID,
			Concept:  concept,
			X:        numbers[0],
			Y:        numbers[1],
			Width:    numbers[2],
			Height:   numbers[3],
			Visual: entity.EngineResolvedVisual{
				Shape: shape, Fill: strings[6], Stroke: strings[7], StrokeWidth: numbers[4],
				CornerRadius: numbers[5], Opacity: numbers[6], Visible: visible == 1, Layer: layer,
			},
			Text: entity.EngineResolvedText{
				Value: strings[1], FontFamily: strings[2], Color: strings[3], Role: strings[4],
				FontSize: numbers[7], LineHeight: numbers[8],
				X: numbers[10], Y: numbers[11], Width: numbers[12], Height: numbers[13],
			},
			Line: entity.EngineResolvedLine{
				Style: lineStyle, SourceDecoration: sourceDecoration, TargetDecoration: targetDecoration,
				Label: strings[8], LabelPosition: numbers[9],
			},
			IconRef: strings[5], IconX: numbers[14], IconY: numbers[15],
			IconWidth: numbers[16], IconHeight: numbers[17], Points: points,
		})
	}
	if reader.Len() != 0 {
		return entity.EngineResolvedDocument{}, fmt.Errorf("Rust engine layout response has %d trailing bytes", reader.Len())
	}
	return entity.EngineResolvedDocument{Width: spec.Width, Height: spec.Height, Elements: elements}, nil
}

func decodeEngineConcept(value byte) (entity.EngineConcept, error) {
	switch value {
	case 1:
		return entity.EngineConceptFrame, nil
	case 2:
		return entity.EngineConceptGroup, nil
	case 3:
		return entity.EngineConceptCapture, nil
	case 4:
		return entity.EngineConceptItem, nil
	case 5:
		return entity.EngineConceptPort, nil
	case 6:
		return entity.EngineConceptLine, nil
	case 7:
		return entity.EngineConceptText, nil
	case 8:
		return entity.EngineConceptSpacer, nil
	default:
		return "", fmt.Errorf("unsupported Rust engine concept %d", value)
	}
}

func decodeEngineShape(value byte) (entity.EngineShape, error) {
	switch value {
	case 1:
		return entity.EngineShapeRectangle, nil
	case 2:
		return entity.EngineShapeEllipse, nil
	case 3:
		return entity.EngineShapeNone, nil
	default:
		return "", fmt.Errorf("unsupported Rust engine shape %d", value)
	}
}

func decodeEngineLineStyle(value byte) (entity.EngineLineStyle, error) {
	switch value {
	case 1:
		return entity.EngineLineSolid, nil
	case 2:
		return entity.EngineLineDashed, nil
	case 3:
		return entity.EngineLineDotted, nil
	default:
		return "", fmt.Errorf("unsupported Rust engine line style %d", value)
	}
}

func decodeEngineDecoration(value byte) (entity.EngineDecoration, error) {
	switch value {
	case 1:
		return entity.EngineDecorationNone, nil
	case 2:
		return entity.EngineDecorationArrow, nil
	case 3:
		return entity.EngineDecorationTriangle, nil
	case 4:
		return entity.EngineDecorationDiamond, nil
	case 5:
		return entity.EngineDecorationCircle, nil
	default:
		return "", fmt.Errorf("unsupported Rust engine decoration %d", value)
	}
}

func decodeEngineBytes(input []byte, operation byte) ([]byte, error) {
	reader, err := decodeEngineResponseHeader(input, operation)
	if err != nil {
		return nil, err
	}
	length, err := readUint32(reader)
	if err != nil {
		return nil, fmt.Errorf("decode Rust engine payload length: %w", err)
	}
	if uint64(length) > uint64(engineMaxResponse) {
		return nil, fmt.Errorf("Rust engine payload size %d exceeds %d", length, engineMaxResponse)
	}
	if uint64(length) != uint64(reader.Len()) {
		return nil, fmt.Errorf("Rust engine payload length %d does not match %d available bytes", length, reader.Len())
	}
	payload := make([]byte, length)
	if _, err := io.ReadFull(reader, payload); err != nil {
		return nil, fmt.Errorf("decode Rust engine payload: %w", err)
	}
	return payload, nil
}

func decodeEngineSVG(input []byte) (entity.EngineSVG, error) {
	reader, err := decodeEngineResponseHeader(input, engineOperationNormalizeSVG)
	if err != nil {
		return entity.EngineSVG{}, err
	}
	width, err := readFloat64(reader)
	if err != nil {
		return entity.EngineSVG{}, fmt.Errorf("decode normalized SVG width: %w", err)
	}
	height, err := readFloat64(reader)
	if err != nil {
		return entity.EngineSVG{}, fmt.Errorf("decode normalized SVG height: %w", err)
	}
	viewBoxLength, err := readUint16(reader)
	if err != nil {
		return entity.EngineSVG{}, fmt.Errorf("decode normalized SVG viewBox length: %w", err)
	}
	reserved, err := readUint16(reader)
	if err != nil {
		return entity.EngineSVG{}, fmt.Errorf("decode normalized SVG reserved field: %w", err)
	}
	if reserved != 0 {
		return entity.EngineSVG{}, fmt.Errorf("normalized SVG reserved field is %d", reserved)
	}
	svgLength, err := readUint32(reader)
	if err != nil {
		return entity.EngineSVG{}, fmt.Errorf("decode normalized SVG length: %w", err)
	}
	totalLength := uint64(viewBoxLength) + uint64(svgLength)
	if totalLength != uint64(reader.Len()) {
		return entity.EngineSVG{}, fmt.Errorf("normalized SVG payload length %d does not match %d available bytes", totalLength, reader.Len())
	}
	viewBox := make([]byte, viewBoxLength)
	if _, err := io.ReadFull(reader, viewBox); err != nil {
		return entity.EngineSVG{}, fmt.Errorf("decode normalized SVG viewBox: %w", err)
	}
	data := make([]byte, svgLength)
	if _, err := io.ReadFull(reader, data); err != nil {
		return entity.EngineSVG{}, fmt.Errorf("decode normalized SVG data: %w", err)
	}
	return entity.EngineSVG{Data: data, ViewBox: string(viewBox), Width: width, Height: height}, nil
}

func decodeEngineResponseHeader(input []byte, expectedOperation byte) (*bytes.Reader, error) {
	reader := bytes.NewReader(input)
	var magic [4]byte
	if _, err := io.ReadFull(reader, magic[:]); err != nil {
		return nil, fmt.Errorf("decode Rust engine response magic: %w", err)
	}
	if magic != engineResponseMagic {
		return nil, fmt.Errorf("invalid Rust engine response magic %q", magic)
	}
	version, err := readUint16(reader)
	if err != nil {
		return nil, fmt.Errorf("decode Rust engine response ABI: %w", err)
	}
	if version != engineABIVersion {
		return nil, fmt.Errorf("Rust engine response ABI version %d does not match %d", version, engineABIVersion)
	}
	status, err := reader.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("decode Rust engine response status: %w", err)
	}
	operation, err := reader.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("decode Rust engine response operation: %w", err)
	}
	if status == engineStatusError {
		diagnostic, decodeErr := decodeEngineError(reader)
		if decodeErr != nil {
			return nil, decodeErr
		}
		return nil, &entity.EngineDiagnosticError{Diagnostic: diagnostic}
	}
	if status != engineStatusOK {
		return nil, fmt.Errorf("unsupported Rust engine response status %d", status)
	}
	if operation != expectedOperation {
		return nil, fmt.Errorf("Rust engine response operation %d does not match %d", operation, expectedOperation)
	}
	return reader, nil
}

func decodeEngineError(reader *bytes.Reader) (entity.EngineDiagnostic, error) {
	length, err := readUint32(reader)
	if err != nil {
		return entity.EngineDiagnostic{}, fmt.Errorf("decode Rust engine error length: %w", err)
	}
	if uint64(length) != uint64(reader.Len()) {
		return entity.EngineDiagnostic{}, fmt.Errorf("Rust engine error length %d does not match %d available bytes", length, reader.Len())
	}
	message := make([]byte, length)
	if _, err := io.ReadFull(reader, message); err != nil {
		return entity.EngineDiagnostic{}, fmt.Errorf("decode Rust engine error: %w", err)
	}
	return entity.EngineDiagnostic{
		Code: "XAL-E2001", Severity: "error", Stage: "calculate",
		Message: "Rust engine: " + string(message),
	}, nil
}

func enrichEngineDiagnostic(err error, spec entity.EngineDocumentSpec) error {
	if err == nil {
		return nil
	}
	var diagnosticErr *entity.EngineDiagnosticError
	if !errors.As(err, &diagnosticErr) {
		return err
	}
	for _, element := range flattenEngineDiagnosticElements(spec.Elements) {
		if element.ID != "" && strings.Contains(diagnosticErr.Diagnostic.Message, fmt.Sprintf("%q", element.ID)) {
			diagnosticErr.Diagnostic.ElementID = element.ID
			diagnosticErr.Diagnostic.SpanID = element.SpanID
			break
		}
	}
	return diagnosticErr
}

// DiagnosticFromEngineError projects a structured engine failure onto the
// source span retained by the frontend. Non-engine failures are returned to
// the caller by reporting ok=false.
func DiagnosticFromEngineError(spec entity.EngineDocumentSpec, err error) (diagnostic entity.Diagnostic, ok bool) {
	var engineErr *entity.EngineDiagnosticError
	if !errors.As(err, &engineErr) {
		return entity.Diagnostic{}, false
	}
	diagnostic = entity.Diagnostic{
		Code: engineErr.Diagnostic.Code, Severity: entity.DiagnosticSeverity(engineErr.Diagnostic.Severity),
		Stage: engineErr.Diagnostic.Stage, Element: engineErr.Diagnostic.ElementID,
		Parameter: engineErr.Diagnostic.Parameter, Message: engineErr.Diagnostic.Message,
	}
	if diagnostic.Severity == "" {
		diagnostic.Severity = "error"
	}
	for _, span := range spec.Spans {
		if span.ID == engineErr.Diagnostic.SpanID {
			diagnostic.Offset, diagnostic.Line, diagnostic.Column = span.Offset, span.Line, span.Column
			break
		}
	}
	return diagnostic, true
}

func flattenEngineDiagnosticElements(roots []entity.EngineElementSpec) []entity.EngineElementSpec {
	result := make([]entity.EngineElementSpec, 0, len(roots))
	var walk func([]entity.EngineElementSpec)
	walk = func(elements []entity.EngineElementSpec) {
		for _, element := range elements {
			result = append(result, element)
			walk(element.Children)
		}
	}
	walk(roots)
	return result
}

func isFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

func writeUint16(output *bytes.Buffer, value uint16) {
	var encoded [2]byte
	binary.LittleEndian.PutUint16(encoded[:], value)
	output.Write(encoded[:])
}

func writeUint32(output *bytes.Buffer, value uint32) {
	var encoded [4]byte
	binary.LittleEndian.PutUint32(encoded[:], value)
	output.Write(encoded[:])
}

func writeUint64(output *bytes.Buffer, value uint64) {
	var encoded [8]byte
	binary.LittleEndian.PutUint64(encoded[:], value)
	output.Write(encoded[:])
}

func writeInt32(output *bytes.Buffer, value int32) {
	writeUint32(output, uint32(value))
}

func writeFloat64(output *bytes.Buffer, value float64) {
	writeUint64(output, math.Float64bits(value))
}

func readUint16(reader *bytes.Reader) (uint16, error) {
	var encoded [2]byte
	if _, err := io.ReadFull(reader, encoded[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(encoded[:]), nil
}

func readUint32(reader *bytes.Reader) (uint32, error) {
	var encoded [4]byte
	if _, err := io.ReadFull(reader, encoded[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(encoded[:]), nil
}

func readInt32(reader *bytes.Reader) (int32, error) {
	value, err := readUint32(reader)
	return int32(value), err
}

func readFloat64(reader *bytes.Reader) (float64, error) {
	var encoded [8]byte
	if _, err := io.ReadFull(reader, encoded[:]); err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(encoded[:])), nil
}
