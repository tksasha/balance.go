package errors

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrUnknown  = errors.New("unknown")
)
