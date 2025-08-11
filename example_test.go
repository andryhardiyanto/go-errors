package errors

import (
	"fmt"
	"testing"
)

// Example penggunaan basic error
func ExampleNew() {
	err := New(404, "Not found", "NOT_FOUND")
	fmt.Println(err.Error())
}

// Example penggunaan predefined errors
func ExampleErrorNotFound() {
	err := ErrorNotFound
	fmt.Printf("Code: %d, Message: %s, Type: %s\n", err.Code, err.Message, err.Type)
}

// Example wrapping existing error
func ExampleWrap() {
	originalErr := fmt.Errorf("database connection failed")
	err := Wrap(originalErr)

	fmt.Println(err.Error())
	fmt.Println("Unwrapped:", err.Unwrap().Error())
}

// Example validation errors
func ExampleViolations() {
	violations := []ValidationError{
		{Type: ViolationErrorTypeRequired, Field: "email", Message: "Email is required"},
		{Type: ViolationErrorTypeEmail, Field: "email", Message: "Invalid email format"},
	}

	err := Violations(violations)
	fmt.Printf("Violations count: %d\n", len(err.Violations))
}

// Example default error
func ExampleDefaultError() {
	err := DefaultError()
	fmt.Printf("Default error: %s\n", err.Error())
}

func TestErrorInterface(t *testing.T) {
	err := New(404, "Not found", "NOT_FOUND")

	// Test error interface
	if err.Error() == "" {
		t.Error("Error() should not return empty string")
	}

	// Test Is method - same error type should return true
	anotherNotFound := New(404, "Not found", "NOT_FOUND")
	if !err.Is(anotherNotFound) {
		t.Error("Is() should return true for same error type")
	}

	// Test Is method - different error type should return false
	differentErr := New(400, "Bad request", "BAD_REQUEST")
	isResult := err.Is(differentErr)
	t.Logf("err.Type: %s, differentErr.Type: %s, Is result: %v", err.Type, differentErr.Type, isResult)
	if isResult {
		t.Error("Is() should return false for different error type")
	}

	// Test Unwrap method
	originalErr := fmt.Errorf("original error")
	wrappedErr := Wrap(originalErr)
	if wrappedErr.Unwrap() != originalErr {
		t.Error("Unwrap() should return the original error")
	}
}

func TestPredefinedErrors(t *testing.T) {
	// Test predefined errors
	if ErrorNotFound.Code != 404 {
		t.Errorf("Expected ErrorNotFound.Code to be 404, got %d", ErrorNotFound.Code)
	}

	if ErrorNotFound.Type != "NOT_FOUND" {
		t.Errorf("Expected ErrorNotFound.Type to be 'NOT_FOUND', got %s", ErrorNotFound.Type)
	}

	if ErrorUnauthorized.Code != 401 {
		t.Errorf("Expected ErrorUnauthorized.Code to be 401, got %d", ErrorUnauthorized.Code)
	}
}

func TestValidationErrors(t *testing.T) {
	violations := []ValidationError{
		{Type: ViolationErrorTypeRequired, Field: "email", Message: "Email is required"},
		{Type: ViolationErrorTypeEmail, Field: "email", Message: "Invalid email format"},
	}

	err := Violations(violations)

	if len(err.Violations) != 2 {
		t.Errorf("Expected 2 violations, got %d", len(err.Violations))
	}

	if err.Violations[0].Type != ViolationErrorTypeRequired {
		t.Errorf("Expected first violation type to be REQUIRED, got %s", err.Violations[0].Type)
	}
}
