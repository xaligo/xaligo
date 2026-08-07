//go:build cgo && xaligo_exporter

package exporter

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/../engine/lib -lxaligo_engine
#cgo linux LDFLAGS: -lutil -lrt -lpthread -lm -ldl
#cgo windows LDFLAGS: -lkernel32 -lntdll -luserenv -lws2_32 -ldbghelp
#cgo windows,arm64 LDFLAGS: -lunwind
#include "xaligo_exporter.h"
*/
import "C"

import (
	"context"
	"fmt"
	"math"
	"unsafe"
)

func Available() bool                      { return true }
func ABIVersion() uint32                   { return uint32(C.xaligo_exporter_abi_version()) }
func Process(input []byte) ([]byte, error) { return process(input) }
func ProcessContext(ctx context.Context, input []byte) ([]byte, error) {
	if ctx != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
	return process(input)
}

func process(input []byte) ([]byte, error) {
	var pointer *C.uint8_t
	if len(input) > 0 {
		pointer = (*C.uint8_t)(unsafe.Pointer(&input[0]))
	}
	var output C.XaligoExporterBuffer
	status := int32(C.xaligo_exporter_process(pointer, C.size_t(len(input)), &output))
	if output.data != nil {
		defer C.xaligo_exporter_buffer_free(output)
	}
	if status != 0 {
		return nil, fmt.Errorf("Rust PPTX exporter C ABI failed with status %d", status)
	}
	if output.len > C.size_t(math.MaxInt32) {
		return nil, fmt.Errorf("Rust PPTX response is too large: %d bytes", uint64(output.len))
	}
	if output.len == 0 {
		return []byte{}, nil
	}
	if output.data == nil {
		return nil, fmt.Errorf("Rust PPTX exporter returned a null response")
	}
	return C.GoBytes(unsafe.Pointer(output.data), C.int(output.len)), nil
}
