package errors

import (
	"runtime"
	"strings"
)

// captureStackTrace captures the current stack trace and returns it as a slice of strings
func captureStackTrace() []string {
	const stackSize = 4096
	stackBuf := make([]byte, stackSize)
	n := runtime.Stack(stackBuf, false)
	stackFrames := string(stackBuf[:n])

	lines := strings.Split(stackFrames, "\n")
	result := make([]string, 0, len(lines))

	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			result = append(result, trimmed)
		}
	}

	return result
}

// New creates a new error with the provided code, message, and error type.
func New(code int64, message, errorType string) *Error {
	e := &Error{
		Type:        errorType,
		Code:        code,
		Violations:  make([]ValidationError, 0),
		Message:     message,
		StackTraces: make([]string, 0),
	}

	if len(e.StackTraces) == 0 {
		e.StackTraces = append(e.StackTraces, captureStackTrace()...)
	}

	return e
}

// Wrap wraps an existing error with a default error, setting the error type, code, and message.
func Wrap(err error) *Error {
	e := DefaultError()
	e.Err = err

	if len(e.StackTraces) == 0 {
		e.StackTraces = captureStackTrace()
	}

	return e
}

// Violations returns a validation error with a 422 status code, "UNPROCESSABLE_ENTITY" type, and the provided validation violations.
func Violations(violations []ValidationError) *Error {
	e := ErrorUnprocessableEntity
	e.Violations = violations

	if len(e.StackTraces) == 0 {
		e.StackTraces = captureStackTrace()
	}

	return e
}

// DefaultError returns a default error with a 500 status code, "INTERNAL_SERVER_ERROR" type, and a generic error message.
func DefaultError() *Error {
	return &Error{
		Type:        "INTERNAL_SERVER_ERROR",
		Code:        500,
		Message:     "An internal server error occurred",
		Violations:  make([]ValidationError, 0),
		StackTraces: captureStackTrace(),
	}
}
