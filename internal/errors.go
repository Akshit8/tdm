package internal

import "fmt"

// ErrorCode defines supported error codes.
type ErrorCode uint

// Error represents application error,
// can be user generated or wrap any other thrown error
// it includes a code for determining what triggered the error.
type Error struct {
	original error
	msg      string
	code     ErrorCode
}

const (
	// ErrorCodeUnknown ...
	ErrorCodeUnknown ErrorCode = iota

	// ErrorCodeNotFound ...
	ErrorCodeNotFound

	// ErrorCodeInvalidArgument ...
	ErrorCodeInvalidArgument
)

// WrapErrorf returns a wrapped error.
func WrapError(original error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		original: original,
		code:     code,
		msg:      fmt.Sprintf(format, a...),
	}
}

// NewErrorf creates a new error
func NewErrorf(code ErrorCode, format string, a ...interface{}) error {
	return WrapError(nil, code, format, a...)
}

// Error returns the message, when wrapping errors the wrapped error is returned.
func (e *Error) Error() string {
	if e.original != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.original)
	}

	return e.msg
}

// Unwrap returns the wrapped error, if any.
func (e *Error) Unwrap() error {
	return e.original
}

// Code returns the code representing this error.
func (e *Error) Code() ErrorCode {
	return e.code
}
