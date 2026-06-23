package entity

import "fmt"

type ParseError struct {
	Position Position
	Err      error
}

func (rcvr *ParseError) Error() string {
	return fmt.Sprintf("line %d, column %d: %v", rcvr.Position.Line, rcvr.Position.Column, rcvr.Err)
}

func (rcvr *ParseError) Unwrap() error { return rcvr.Err }
