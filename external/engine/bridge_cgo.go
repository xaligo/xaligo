//go:build cgo && xaligo_engine

package engine

/*
#cgo CFLAGS: -I${SRCDIR}/include
#cgo LDFLAGS: -L${SRCDIR}/lib -lxaligo_engine
#cgo linux LDFLAGS: -lutil -lrt -lpthread -lm -ldl
#cgo windows LDFLAGS: -lkernel32 -lntdll -luserenv -lws2_32 -ldbghelp
#cgo windows,arm64 LDFLAGS: -lunwind
#include "xaligo_engine.h"
*/
import "C"

import (
	"context"
	"fmt"
	"math"
	"unsafe"
)

// Available reports whether this build contains the Rust engine static
// library.
func Available() bool {
	return true
}

// ABIVersion returns the version exported by the linked Rust library.
func ABIVersion() uint32 {
	return uint32(C.xaligo_engine_abi_version())
}

// Process synchronously transfers one versioned request to Rust and copies the
// Rust-owned response into Go memory before releasing it through the C ABI.
func Process(input []byte) ([]byte, error) {
	return process(input, nil)
}

// ProcessContext performs the same synchronous transfer while a small C-owned
// atomic handle lets Rust calculation loops observe context cancellation.
func ProcessContext(ctx context.Context, input []byte) ([]byte, error) {
	if ctx == nil {
		return Process(input)
	}
	cancel := C.xaligo_engine_cancel_new()
	if cancel == nil {
		return nil, fmt.Errorf("allocate Rust engine cancellation handle")
	}
	defer C.xaligo_engine_cancel_free(cancel)
	done := make(chan struct{})
	watcherDone := make(chan struct{})
	go func() {
		defer close(watcherDone)
		select {
		case <-ctx.Done():
			C.xaligo_engine_cancel_set(cancel)
		case <-done:
		}
	}()
	result, err := process(input, cancel)
	close(done)
	<-watcherDone
	return result, err
}

func process(input []byte, cancel *C.XaligoEngineCancel) ([]byte, error) {
	var inputPointer *C.uint8_t
	if len(input) > 0 {
		inputPointer = (*C.uint8_t)(unsafe.Pointer(&input[0]))
	}

	var output C.XaligoEngineBuffer
	status := int32(C.xaligo_engine_process_with_cancel(inputPointer, C.size_t(len(input)), cancel, &output))
	if output.data != nil {
		defer C.xaligo_engine_buffer_free(output)
	}
	if status != 0 {
		return nil, fmt.Errorf("Rust engine C ABI failed with status %d", status)
	}
	if output.len > C.size_t(math.MaxInt32) {
		return nil, fmt.Errorf("Rust engine response is too large: %d bytes", uint64(output.len))
	}
	if output.len == 0 {
		return []byte{}, nil
	}
	if output.data == nil {
		return nil, fmt.Errorf("Rust engine returned a null response for %d bytes", uint64(output.len))
	}
	return C.GoBytes(unsafe.Pointer(output.data), C.int(output.len)), nil
}
