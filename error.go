package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// captureStackTrace captures the current stack trace using runtime.Callers
// skip parameter indicates how many stack frames to skip (0 = current function, 1 = caller, etc.)
func captureStackTrace(skip int) []string {
	const maxFrames = 32
	pcs := make([]uintptr, maxFrames)

	// Skip additional frames: skip + 1 (for captureStackTrace itself)
	n := runtime.Callers(skip+2, pcs)
	if n == 0 {
		return []string{}
	}

	frames := runtime.CallersFrames(pcs[:n])
	result := make([]string, 0, n)

	for {
		frame, more := frames.Next()

		// Skip internal runtime frames and this package's internal frames
		if isRelevantFrame(frame) {
			result = append(result, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
		}

		if !more {
			break
		}
	}
	return result
}

// isRelevantFrame determines if a stack frame is relevant for error reporting
func isRelevantFrame(frame runtime.Frame) bool {
	// Skip runtime internal functions
	if strings.HasPrefix(frame.Function, "runtime.") {
		return false
	}

	// Skip testing framework functions
	if strings.HasPrefix(frame.Function, "testing.") {
		return false
	}

	// Include all other frames
	return true
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
		e.StackTraces = append(e.StackTraces, captureStackTrace(1)...)
	}

	return e
}

// Wrap wraps an existing error with a default error, setting the error type, code, and message.
func Wrap(err error) *Error {
	e := &Error{
		Type:        "INTERNAL_SERVER_ERROR",
		Code:        500,
		Message:     "An internal server error occurred",
		Violations:  make([]ValidationError, 0),
		StackTraces: captureStackTrace(1),
		Err:         err,
	}

	return e
}

// Violations returns a validation error with a 422 status code, "UNPROCESSABLE_ENTITY" type, and the provided validation violations.
func Violations(violations []ValidationError) *Error {
	e := &Error{
		Type:        "UNPROCESSABLE_ENTITY",
		Code:        422,
		Message:     "Unprocessable entity",
		Violations:  violations,
		StackTraces: captureStackTrace(1),
	}

	return e
}

// Factory functions for common errors - these capture stack trace when called, not during package init
func ErrorBadRequest() *Error {
	e := &Error{
		Type:        "BAD_REQUEST",
		Code:        400,
		Violations:  make([]ValidationError, 0),
		Message:     "Bad request",
		StackTraces: captureStackTrace(1),
	}
	return e
}

func ErrorUnauthorized() *Error {
	e := &Error{
		Type:        "UNAUTHORIZED",
		Code:        401,
		Violations:  make([]ValidationError, 0),
		Message:     "Unauthorized",
		StackTraces: captureStackTrace(1),
	}
	return e
}

func ErrorForbidden() *Error {
	e := &Error{
		Type:        "FORBIDDEN",
		Code:        403,
		Violations:  make([]ValidationError, 0),
		Message:     "Forbidden",
		StackTraces: captureStackTrace(1),
	}
	return e
}

func ErrorNotFound() *Error {
	e := &Error{
		Type:        "NOT_FOUND",
		Code:        404,
		Violations:  make([]ValidationError, 0),
		Message:     "Not found",
		StackTraces: captureStackTrace(1),
	}
	return e
}

func ErrorConflict() *Error {
	e := &Error{
		Type:        "CONFLICT",
		Code:        409,
		Violations:  make([]ValidationError, 0),
		Message:     "Conflict",
		StackTraces: captureStackTrace(1),
	}
	return e
}

func ErrorUnprocessableEntity() *Error {
	e := &Error{
		Type:        "UNPROCESSABLE_ENTITY",
		Code:        422,
		Violations:  make([]ValidationError, 0),
		Message:     "Unprocessable entity",
		StackTraces: captureStackTrace(1),
	}
	return e
}

func ErrorInternalServerError() *Error {
	e := &Error{
		Type:        "INTERNAL_SERVER_ERROR",
		Code:        500,
		Violations:  make([]ValidationError, 0),
		Message:     "Internal server error",
		StackTraces: captureStackTrace(1),
	}
	return e
}

func ErrorPanic() *Error {
	e := &Error{
		Type:        "PANIC",
		Code:        500,
		Violations:  make([]ValidationError, 0),
		Message:     "Panic",
		StackTraces: captureStackTrace(1),
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
		StackTraces: captureStackTrace(1),
	}
}
