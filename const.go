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

var (
	// Common errors
	ErrorBadRequest          = New(400, "Bad request", "BAD_REQUEST")
	ErrorUnauthorized        = New(401, "Unauthorized", "UNAUTHORIZED")
	ErrorForbidden           = New(403, "Forbidden", "FORBIDDEN")
	ErrorNotFound            = New(404, "Not found", "NOT_FOUND")
	ErrorConflict            = New(409, "Conflict", "CONFLICT")
	ErrorUnprocessableEntity = New(422, "Unprocessable entity", "UNPROCESSABLE_ENTITY")
	ErrorInternalServerError = New(500, "Internal server error", "INTERNAL_SERVER_ERROR")
	ErrorPanic               = New(500, "Panic", "PANIC")
)
