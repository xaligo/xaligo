package share

import (
	"strings"
	"sync"
)

// MCode represents a structured message code used in logs.
type MCode struct {
	Code    string
	Message string
}

var (
	mcodeMu       sync.RWMutex
	maxCodeLength int
)

// NewMCode creates a message code and updates the shared padding width.
func NewMCode(code, message string) MCode {
	mcode := MCode{Code: code, Message: message}
	RegisterMCodes(mcode)
	return mcode
}

// Mcode returns the provided message code unchanged.
func Mcode(mcode MCode) MCode { return mcode }

// RegisterMCodes records message codes for aligned human-readable logging.
func RegisterMCodes(mcodes ...MCode) {
	mcodeMu.Lock()
	defer mcodeMu.Unlock()
	for _, mcode := range mcodes {
		if len(mcode.Code) > maxCodeLength {
			maxCodeLength = len(mcode.Code)
		}
	}
}

// PaddedCode returns Code padded to the longest registered code length.
func (rcvr MCode) PaddedCode() string {
	mcodeMu.RLock()
	width := maxCodeLength
	mcodeMu.RUnlock()
	if len(rcvr.Code) >= width {
		return rcvr.Code
	}
	return rcvr.Code + strings.Repeat(" ", width-len(rcvr.Code))
}

// GetMaxCodeLength returns the longest registered message code length.
func GetMaxCodeLength() int {
	mcodeMu.RLock()
	defer mcodeMu.RUnlock()
	return maxCodeLength
}

var (
	MSYS1 = NewMCode("MSYS1", "System start")
	MSYS2 = NewMCode("MSYS2", "System error")
	MLOG1 = NewMCode("MLOG1", "Logger created")
	MLOG2 = NewMCode("MLOG2", "Logger output fallback")
)
