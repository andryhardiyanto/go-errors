package errors

const (
	// Common validation error types
	ViolationErrorTypeRequired   ViolationErrorType = "REQUIRED"
	ViolationErrorTypeOneOf      ViolationErrorType = "ONEOF"
	ViolationErrorTypeUUID       ViolationErrorType = "UUID"
	ViolationErrorTypeMin        ViolationErrorType = "MIN"
	ViolationErrorTypeMax        ViolationErrorType = "MAX"
	ViolationErrorTypeEmail      ViolationErrorType = "EMAIL"
	ViolationErrorTypeDate       ViolationErrorType = "DATE"
	ViolationErrorTypeRequiredIf ViolationErrorType = "REQUIRED_IF"
)
