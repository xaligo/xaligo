package entity

import "fmt"

type ParseError struct {
	Position Position
	Err      error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d, column %d: %v", e.Position.Line, e.Position.Column, e.Err)
}

func (e *ParseError) Unwrap() error { return e.Err }
