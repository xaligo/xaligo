package v2

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"

	enginebridge "github.com/xaligo/xaligo/external/engine"
	"github.com/xaligo/xaligo/internal/entity"
)

type EngineUsecase interface {
	Resolve(context.Context, entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error)
	RenderSVG(context.Context, entity.EngineDocumentSpec) ([]byte, error)
}

type engineUsecase struct{}

func NewEngineUsecase() EngineUsecase {
	return &engineUsecase{}
}

const (
	engineABIVersion      = uint16(1)
	engineOperationLayout = byte(1)
	engineOperationSVG    = byte(2)
	engineStatusOK        = byte(0)
	engineStatusError     = byte(1)
	engineMaxElements     = 10_000
	engineMaxResponse     = 32 * 1024 * 1024
)

var (
	engineRequestMagic  = [4]byte{'X', 'L', 'E', '2'}
	engineResponseMagic = [4]byte{'X', 'L', 'R', '2'}
)

func (rcvr *engineUsecase) Resolve(ctx context.Context, spec entity.EngineDocumentSpec) (entity.EngineResolvedDocument, error) {
	response, err := rcvr.execute(ctx, engineOperationLayout, spec)
	if err != nil {
		return entity.EngineResolvedDocument{}, err
	}
	return decodeEngineLayout(response, spec)
}

func (rcvr *engineUsecase) RenderSVG(ctx context.Context, spec entity.EngineDocumentSpec) ([]byte, error) {
	response, err := rcvr.execute(ctx, engineOperationSVG, spec)
	if err != nil {
		return nil, err
	}
	return decodeEngineBytes(response, engineOperationSVG)
}

func (rcvr *engineUsecase) execute(ctx context.Context, operation byte, spec entity.EngineDocumentSpec) ([]byte, error) {
	if err := checkEngineContext(ctx); err != nil {
		return nil, err
	}
	if !enginebridge.Available() {
		return nil, enginebridge.ErrUnavailable
	}
	if version := enginebridge.ABIVersion(); version != uint32(engineABIVersion) {
		return nil, fmt.Errorf("Rust engine ABI version %d does not match Go ABI version %d", version, engineABIVersion)
	}
	request, err := encodeEngineRequest(operation, spec)
	if err != nil {
		return nil, err
	}
	response, err := enginebridge.Process(request)
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
	if len(spec.Elements) > engineMaxElements {
		return nil, fmt.Errorf("engine element count %d exceeds %d", len(spec.Elements), engineMaxElements)
	}
	direction, err := encodeEngineDirection(spec.Direction)
	if err != nil {
		return nil, err
	}

	var output bytes.Buffer
	output.Grow(36 + len(spec.Elements)*32)
	output.Write(engineRequestMagic[:])
	writeUint16(&output, engineABIVersion)
	output.WriteByte(operation)
	output.WriteByte(direction)
	writeFloat64(&output, spec.Width)
	writeFloat64(&output, spec.Height)
	writeFloat64(&output, spec.Gap)
	writeUint32(&output, uint32(len(spec.Elements)))
	for _, element := range spec.Elements {
		identifier := []byte(element.ID)
		if len(identifier) > math.MaxUint16 {
			return nil, fmt.Errorf("engine element id exceeds %d UTF-8 bytes", math.MaxUint16)
		}
		writeUint16(&output, uint16(len(identifier)))
		var flags byte
		if element.Width != nil {
			flags |= 1 << 0
		}
		if element.Height != nil {
			flags |= 1 << 1
		}
		if element.Weight != nil {
			flags |= 1 << 2
		}
		output.WriteByte(flags)
		output.WriteByte(0)
		writeFloat64(&output, optionalFloat64(element.Width))
		writeFloat64(&output, optionalFloat64(element.Height))
		writeFloat64(&output, optionalFloat64(element.Weight))
		output.Write(identifier)
	}
	return output.Bytes(), nil
}

func encodeEngineDirection(direction entity.EngineDirection) (byte, error) {
	switch direction {
	case entity.EngineDirectionVertical:
		return 1, nil
	case entity.EngineDirectionHorizontal:
		return 2, nil
	default:
		return 0, fmt.Errorf("unsupported engine direction %q", direction)
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
	for range count {
		identifierLength, err := readUint16(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine element id length: %w", err)
		}
		reserved, err := readUint16(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine element reserved field: %w", err)
		}
		if reserved != 0 {
			return entity.EngineResolvedDocument{}, fmt.Errorf("Rust engine element reserved field is %d", reserved)
		}
		x, err := readFloat64(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine element x: %w", err)
		}
		y, err := readFloat64(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine element y: %w", err)
		}
		width, err := readFloat64(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine element width: %w", err)
		}
		height, err := readFloat64(reader)
		if err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine element height: %w", err)
		}
		identifier := make([]byte, identifierLength)
		if _, err := io.ReadFull(reader, identifier); err != nil {
			return entity.EngineResolvedDocument{}, fmt.Errorf("decode Rust engine element id: %w", err)
		}
		elements = append(elements, entity.EngineResolvedElement{
			ID: string(identifier), X: x, Y: y, Width: width, Height: height,
		})
	}
	if reader.Len() != 0 {
		return entity.EngineResolvedDocument{}, fmt.Errorf("Rust engine layout response has %d trailing bytes", reader.Len())
	}
	return entity.EngineResolvedDocument{
		Width: spec.Width, Height: spec.Height, Elements: elements,
	}, nil
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
		message, decodeErr := decodeEngineError(reader)
		if decodeErr != nil {
			return nil, decodeErr
		}
		return nil, errors.New(message)
	}
	if status != engineStatusOK {
		return nil, fmt.Errorf("unsupported Rust engine response status %d", status)
	}
	if operation != expectedOperation {
		return nil, fmt.Errorf("Rust engine response operation %d does not match %d", operation, expectedOperation)
	}
	return reader, nil
}

func decodeEngineError(reader *bytes.Reader) (string, error) {
	length, err := readUint32(reader)
	if err != nil {
		return "", fmt.Errorf("decode Rust engine error length: %w", err)
	}
	if uint64(length) != uint64(reader.Len()) {
		return "", fmt.Errorf("Rust engine error length %d does not match %d available bytes", length, reader.Len())
	}
	message := make([]byte, length)
	if _, err := io.ReadFull(reader, message); err != nil {
		return "", fmt.Errorf("decode Rust engine error: %w", err)
	}
	return "Rust engine: " + string(message), nil
}

func optionalFloat64(value *float64) float64 {
	if value == nil {
		return 0
	}
	return *value
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

func writeFloat64(output *bytes.Buffer, value float64) {
	var encoded [8]byte
	binary.LittleEndian.PutUint64(encoded[:], math.Float64bits(value))
	output.Write(encoded[:])
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

func readFloat64(reader *bytes.Reader) (float64, error) {
	var encoded [8]byte
	if _, err := io.ReadFull(reader, encoded[:]); err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(encoded[:])), nil
}
