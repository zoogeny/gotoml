package gotoml

import (
	"fmt"
)

const (
	// parse erros
	CouldNotParse = iota

	// get errors
	KeyNotFound
	InvalidType
)

type ParseError struct {
	Reason     int
	LineNumber int
}

type GetError struct {
	Reason        int
	RequestedKey  string
	RequestedType string
	ActualValue   string
}

func (e *GetError) Error() string {
	switch e.Reason {
	case KeyNotFound:
		return fmt.Sprintf("Could not find key %s of type %s", e.RequestedKey, e.RequestedType)
	case InvalidType:
		return fmt.Sprintf("Could not parse key %s as type %s", e.RequestedKey, e.RequestedType)
	}
	return "Unknown reason for GetError"
}

func NewKeyNotFoundError(k string, t string) *GetError {
	return &GetError{KeyNotFound, k, t, ""}
}

func NewInvalidTypeError(k string, v string, t string) *GetError {
	return &GetError{InvalidType, k, t, v}
}
