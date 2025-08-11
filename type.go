package errors

type (
	ViolationErrorType string
	ValidationError    struct {
		Type    ViolationErrorType
		Field   string
		Message string
	}

	Error struct {
		Type        string
		Code        int64
		Message     string
		Violations  []ValidationError
		Err         error
		StackTraces []string
	}
)

// Error implements the error interface
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Message
}

// Unwrap returns the wrapped error, implementing the errors.Unwrap interface
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

// Is reports whether any error in err's chain matches target
func (e *Error) Is(target error) bool {
	if e == nil {
		return target == nil
	}

	if targetErr, ok := target.(*Error); ok {
		return e.Type == targetErr.Type
	}

	// Check if the underlying error matches
	if e.Err != nil {
		return e.Err == target
	}

	return false
}
