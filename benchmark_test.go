package errors

import (
	"fmt"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = New(404, "Not found", "NOT_FOUND")
	}
}

func BenchmarkWrap(b *testing.B) {
	originalErr := fmt.Errorf("original error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Wrap(originalErr)
	}
}

func BenchmarkValidations(b *testing.B) {
	violations := []ValidationError{
		{Type: ViolationErrorTypeRequired, Field: "email", Message: "Email is required"},
		{Type: ViolationErrorTypeEmail, Field: "email", Message: "Invalid email format"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Violations(violations)
	}
}

func BenchmarkErrorString(b *testing.B) {
	err := New(404, "Test error", "NOT_FOUND")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func BenchmarkErrorIs(b *testing.B) {
	err1 := New(404, "Not found", "NOT_FOUND")
	err2 := New(404, "Not found", "NOT_FOUND")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = err1.Is(err2)
	}
}

func BenchmarkCaptureStackTrace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = captureStackTrace()
	}
}

func BenchmarkDefaultError(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DefaultError()
	}
}
