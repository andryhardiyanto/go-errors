# Go Errors Library

A simple and efficient error handling library for Go applications with support for structured errors, stack traces, and validation errors.

## Features

- 📚 **Stack Traces**: Automatic stack trace capture
- 🔧 **Structured Errors**: Well-defined error types with codes
- ✅ **Validation Errors**: Support for field validation errors
- 🔗 **Error Wrapping**: Compatible with Go 1.13+ error wrapping
- 🚀 **Predefined Errors**: Common HTTP error types ready to use
- 🎯 **Simple API**: Clean and straightforward interface

## Installation

```bash
go get github.com/andryhardiyanto/go-errors
```

## Quick Start

```go
package main

import (
    "fmt"
    errors "github.com/andryhardiyanto/go-errors"
)

func main() {
    // Create custom error
    err := errors.New(404, "User not found", "NOT_FOUND")
    fmt.Println(err.Error())

    // Use predefined errors
    err = errors.ErrorNotFound
    fmt.Printf("Code: %d, Message: %s\n", err.Code, err.Message)
    
    // Other predefined errors
    badRequestErr := errors.ErrorBadRequest
    unauthorizedErr := errors.ErrorUnauthorized
    conflictErr := errors.ErrorConflict
    
    // Wrap existing error
    originalErr := fmt.Errorf("database error")
    wrappedErr := errors.Wrap(originalErr)
    fmt.Println(wrappedErr.Error())
    
    // Validation errors
    violations := []errors.ValidationError{
        {Type: errors.ViolationErrorTypeRequired, Field: "email", Message: "Email is required"},
        {Type: errors.ViolationErrorTypeEmail, Field: "email", Message: "Invalid email format"},
        {Type: errors.ViolationErrorTypeMin, Field: "age", Message: "Age must be at least 18"},
    }
    validationErr := errors.Violations(violations)
    fmt.Printf("Validation errors: %d\n", len(validationErr.Violations))
    
    // Default error
    defaultErr := errors.DefaultError()
    fmt.Println(defaultErr.Error())
}
```

## API Reference

### Creating Errors

#### `New(code int64, message string, errorType string) *Error`
Creates a new error with the specified code, message, and type.

```go
err := errors.New(400, "Invalid input", "BAD_REQUEST")
```

#### `Wrap(err error) *Error`
Wraps an existing error with stack trace information.

```go
originalErr := fmt.Errorf("connection failed")
wrappedErr := errors.Wrap(originalErr)
```

#### `DefaultError() *Error`
Returns a default internal server error.

```go
err := errors.DefaultError()
```

### Predefined Errors

| Variable | Code | Type | Message |
|----------|------|------|---------|
| `ErrorBadRequest` | 400 | BAD_REQUEST | Bad request |
| `ErrorUnauthorized` | 401 | UNAUTHORIZED | Unauthorized |
| `ErrorForbidden` | 403 | FORBIDDEN | Forbidden |
| `ErrorNotFound` | 404 | NOT_FOUND | Not found |
| `ErrorConflict` | 409 | CONFLICT | Conflict |
| `ErrorUnprocessableEntity` | 422 | UNPROCESSABLE_ENTITY | Unprocessable entity |
| `ErrorInternalServerError` | 500 | INTERNAL_SERVER_ERROR | Internal server error |
| `ErrorPanic` | 500 | PANIC | Panic |

### Validation Errors

```go
violations := []errors.ValidationError{
    {Type: errors.ViolationErrorTypeRequired, Field: "email", Message: "Email is required"},
    {Type: errors.ViolationErrorTypeEmail, Field: "email", Message: "Invalid email format"},
}

err := errors.Violations(violations)
fmt.Printf("Validation errors: %d\n", len(err.Violations))
```

#### Validation Error Types

| Constant | Value | Description |
|----------|-------|-------------|
| `ViolationErrorTypeRequired` | REQUIRED | Field is required |
| `ViolationErrorTypeOneOf` | ONEOF | Value must be one of specified options |
| `ViolationErrorTypeUUID` | UUID | Value must be a valid UUID |
| `ViolationErrorTypeMin` | MIN | Value is below minimum threshold |
| `ViolationErrorTypeMax` | MAX | Value exceeds maximum threshold |
| `ViolationErrorTypeEmail` | EMAIL | Value must be a valid email format |
| `ViolationErrorTypeDate` | DATE | Value must be a valid date format |
| `ViolationErrorTypeRequiredIf` | REQUIRED_IF | Field is required under certain conditions |

### Error Methods

#### `Error() string`
Returns the error message as a string.

#### `Is(target error) bool`
Checks if the error is of the same type as the target error.

```go
err1 := errors.New(404, "Not found", "NOT_FOUND")
err2 := errors.New(404, "Not found", "NOT_FOUND")

if err1.Is(err2) {
    fmt.Println("Same error type")
}
```

#### `Unwrap() error`
Returns the wrapped error, if any.

```go
originalErr := fmt.Errorf("original error")
wrappedErr := errors.Wrap(originalErr)

if unwrapped := wrappedErr.Unwrap(); unwrapped != nil {
    fmt.Println("Original:", unwrapped.Error())
}
```

## Error Structure

```go
type Error struct {
    Type        string            `json:"type"`
    Code        int64             `json:"code"`
    Message     string            `json:"message"`
    Violations  []ValidationError `json:"violations,omitempty"`
    Err         error             `json:"-"`
    StackTraces []string          `json:"stack_traces,omitempty"`
}
```

## Examples

### Using Predefined Errors

```go
// HTTP 400 - Bad Request
if invalidInput {
    return errors.ErrorBadRequest
}

// HTTP 401 - Unauthorized
if !isAuthenticated {
    return errors.ErrorUnauthorized
}

// HTTP 403 - Forbidden
if !hasPermission {
    return errors.ErrorForbidden
}

// HTTP 404 - Not Found
if user == nil {
    return errors.ErrorNotFound
}

// HTTP 409 - Conflict
if emailExists {
    return errors.ErrorConflict
}

// HTTP 422 - Unprocessable Entity
if hasValidationErrors {
    return errors.ErrorUnprocessableEntity
}

// HTTP 500 - Internal Server Error
if dbError != nil {
    return errors.ErrorInternalServerError
}
```

### Validation Error Examples

```go
// Single validation error
violations := []errors.ValidationError{
    {Type: errors.ViolationErrorTypeRequired, Field: "username", Message: "Username is required"},
}

// Multiple validation errors
violations := []errors.ValidationError{
    {Type: errors.ViolationErrorTypeRequired, Field: "email", Message: "Email is required"},
    {Type: errors.ViolationErrorTypeEmail, Field: "email", Message: "Invalid email format"},
    {Type: errors.ViolationErrorTypeMin, Field: "password", Message: "Password must be at least 8 characters"},
    {Type: errors.ViolationErrorTypeMax, Field: "bio", Message: "Bio cannot exceed 500 characters"},
    {Type: errors.ViolationErrorTypeUUID, Field: "user_id", Message: "Invalid user ID format"},
    {Type: errors.ViolationErrorTypeDate, Field: "birth_date", Message: "Invalid date format"},
    {Type: errors.ViolationErrorTypeOneOf, Field: "gender", Message: "Gender must be male, female, or other"},
    {Type: errors.ViolationErrorTypeRequiredIf, Field: "company", Message: "Company is required for business accounts"},
}

err := errors.Violations(violations)
```

### Error Handling in HTTP Handlers

```go
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        writeErrorResponse(w, errors.ErrorBadRequest)
        return
    }

    if violations := validateUser(user); len(violations) > 0 {
        err := errors.Violations(violations)
        writeErrorResponse(w, err)
        return
    }

    if err := userService.Create(user); err != nil {
        if errors.Is(err, ErrUserExists) {
            writeErrorResponse(w, errors.ErrorConflict)
            return
        }
        writeErrorResponse(w, errors.ErrorInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func writeErrorResponse(w http.ResponseWriter, err *errors.Error) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(int(err.Code))
    json.NewEncoder(w).Encode(err)
}
```

## Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## Support

If you encounter any issues or have questions, please open an issue on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.