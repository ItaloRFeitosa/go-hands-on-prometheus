package errs

import (
	"errors"
)

type Error struct {
	Type    Type   `json:"type"`
	Context string `json:"context"`
	Code    string `json:"code"`
	Message string `json:"message"`

	template string
	err      error
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Is(target error) bool {
	if targetError, ok := target.(Error); ok {
		return targetError.Code == e.Code && targetError.Context == e.Context && targetError.Type == e.Type
	}

	return errors.Is(e.err, target)
}

func AsError(err error) (Error, bool) {
	asError, ok := err.(Error)

	return asError, ok
}
