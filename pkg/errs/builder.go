package errs

import (
	"fmt"
)

type BuilderContext interface {
	WithContext(string) BuilderType
}

type BuilderType interface {
	AsInvalidOperation() BuilderCode
	AsNotFound() BuilderCode
	AsInternal() BuilderCode
	AsInvalidParams() BuilderCode
}

type BuilderCode interface {
	WithCode(string) BuilderError
}

type BuilderError interface {
	WithError(error) error
	WithTemplate(string) BuilderTemplateWithArgs
	WithMessage(string) error
	Is(error) bool
	Error() string
}

type BuilderTemplateWithArgs interface {
	WithArgs(args ...any) error
	Is(error) bool
	Error() string
}

func Builder() BuilderContext {
	return Error{}
}

func New() Error {
	return Error{}
}

func (e Error) AsInvalidOperation() BuilderCode {
	return e.withType(InvalidOperationType)
}

func (e Error) AsInvalidParams() BuilderCode {
	return e.withType(InvalidParamsType)
}

func (e Error) AsNotFound() BuilderCode {
	return e.withType(NotFoundType)
}

func (e Error) AsInternal() BuilderCode {
	return e.withType(InternalType)
}

func (e Error) withType(t Type) Error {
	e.Type = t
	return e
}

func (e Error) WithContext(context string) BuilderType {
	e.Context = context
	return e
}

func (e Error) WithCode(code string) BuilderError {
	return e.withCode(code)
}

func (e Error) withCode(code string) Error {
	e.Code = code
	return e.withError(fmt.Errorf(code))
}

func (e Error) WithMessage(msg string) error {
	return e.WithError(fmt.Errorf(msg))
}

func (e Error) WithTemplate(template string) BuilderTemplateWithArgs {
	e.template = template
	return e.withError(fmt.Errorf(e.template))
}

func (e Error) WithArgs(args ...any) error {
	return e.withError(fmt.Errorf(e.template, args...))
}

func (e Error) WithError(err error) error {
	return e.withError(err)
}

func (e Error) withError(err error) Error {
	e.err = err
	e.Message = err.Error()
	return e
}
