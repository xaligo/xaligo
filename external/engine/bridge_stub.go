//go:build !cgo || !xaligo_engine

package engine

func Available() bool {
	return false
}

func ABIVersion() uint32 {
	return 0
}

func Process([]byte) ([]byte, error) {
	return nil, ErrUnavailable
}
