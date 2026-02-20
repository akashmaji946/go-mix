/*
File    : go-mix/eval/evaluator_helpers.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"strings"
	"testing"

	"github.com/akashmaji946/go-mix/std"
)

// IsError checks if a GoMixObject represents an error condition.
//
// This helper function is used throughout the evaluator to detect error objects
// and enable early termination of evaluation. When an error is detected, it should
// be propagated up the call stack rather than continuing evaluation.
//
// The function includes a nil check to safely handle cases where the object
// might be nil (though this should rarely occur in normal operation).
//
// Parameters:
//   - obj: The GoMixObject to check (can be nil)
//
// Returns:
//   - bool: true if the object is non-nil and has type ErrorType, false otherwise
//
// Example usage:
//
//	result := e.Eval(node)
//	if IsError(result) {
//	    return result  // Propagate error up
//	}
//	// Continue with normal evaluation
func IsError(obj std.GoMixObject) bool {
	if obj != nil {
		return obj.GetType() == std.ErrorType
	}
	return false
}

// UnwrapReturnValue extracts the actual value from a ReturnValue wrapper.
//
// This helper function is used to unwrap return values after function execution
// completes. During evaluation, return statements create ReturnValue wrappers to
// signal early termination. Once we've exited the function context, we need to
// extract the actual returned value.
//
// If the object is not a ReturnValue (i.e., it's a normal value), it's returned
// unchanged. This makes the function safe to call on any object.
//
// Parameters:
//   - obj: The GoMixObject to potentially unwrap
//
// Returns:
//   - objects.GoMixObject: The unwrapped value if obj is a ReturnValue,
//     otherwise returns obj unchanged
//
// Example usage:
//
//	result := e.evalStatements(statements)
//	return UnwrapReturnValue(result)  // Extract value from ReturnValue wrapper
//
// Example flow:
//
//	func add(a, b) { return a + b; }  // Creates ReturnValue(Integer(8))
//	add(5, 3)                          // UnwrapReturnValue extracts Integer(8)
func UnwrapReturnValue(obj std.GoMixObject) std.GoMixObject {
	if retVal, isReturn := obj.(*std.ReturnValue); isReturn {
		return retVal.Value
	}
	return obj
}

// AssertError is a test helper function that validates error objects and their messages.
//
// This function performs two critical checks for testing error conditions:
// 1. Verifies that the object is actually an Error type (not some other type)
// 2. Checks that the error message contains the expected substring
//
// The substring matching (rather than exact matching) allows tests to focus on
// the key error information without being brittle to formatting changes like
// line numbers or exact wording.
//
// If either check fails, the test is marked as failed with a descriptive message.
//
// Parameters:
//   - t: The testing.T instance for reporting test failures
//   - obj: The GoMixObject to check (should be an Error)
//   - expected: A substring that should appear in the error message
//
// Example usage in tests:
//
//	result := ev.Eval(node)
//	AssertError(t, result, "identifier not found")
//	// Passes if result is Error with message containing "identifier not found"
func AssertError(t *testing.T, obj std.GoMixObject, expected string) {
	errObj, ok := obj.(*std.Error)
	if !ok {
		t.Errorf("not error. got=%T (%+v)", obj, obj)
		return
	}
	if !strings.Contains(errObj.Message, expected) {
		t.Errorf("wrong error message. expected to contain=%q, got=%q", expected, errObj.Message)
	}
}

// AssertInteger is a test helper function that validates integer objects and their values.
//
// This function performs two checks for testing integer results:
// 1. Verifies that the object is an Integer type (not Float, String, etc.)
// 2. Checks that the integer value exactly matches the expected value
//
// If either check fails, the test is marked as failed with a descriptive message
// showing what was expected versus what was received.
//
// Parameters:
//   - t: The testing.T instance for reporting test failures
//   - obj: The GoMixObject to check (should be an Integer)
//   - expected: The expected integer value (int64)
//
// Example usage in tests:
//
//	result := ev.Eval(parseExpression("5 + 3"))
//	AssertInteger(t, result, 8)
//	// Passes if result is Integer with value 8
func AssertInteger(t *testing.T, obj std.GoMixObject, expected int64) {
	result, ok := obj.(*std.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
}

// AssertBoolean is a test helper function that validates boolean objects and their values.
//
// This function performs two checks for testing boolean results:
// 1. Verifies that the object is a Boolean type (not Integer, String, etc.)
// 2. Checks that the boolean value exactly matches the expected value
//
// If either check fails, the test is marked as failed with a descriptive message
// showing what was expected versus what was received.
//
// Parameters:
//   - t: The testing.T instance for reporting test failures
//   - obj: The GoMixObject to check (should be a Boolean)
//   - expected: The expected boolean value (true or false)
//
// Example usage in tests:
//
//	result := ev.Eval(parseExpression("5 > 3"))
//	AssertBoolean(t, result, true)
//	// Passes if result is Boolean with value true
func AssertBoolean(t *testing.T, obj std.GoMixObject, expected bool) {
	result, ok := obj.(*std.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
	}

}

// AssertFloat is a test helper function that validates float objects and their values.
//
// This function performs two checks for testing floating-point results:
// 1. Verifies that the object is a Float type (not Integer, String, etc.)
// 2. Checks that the float value exactly matches the expected value
//
// Note: This uses exact equality (==) for float comparison, which may not be
// suitable for all floating-point tests due to precision issues. For tests
// requiring approximate equality, consider using a tolerance-based comparison.
//
// If either check fails, the test is marked as failed with a descriptive message
// showing what was expected versus what was received.
//
// Parameters:
//   - t: The testing.T instance for reporting test failures
//   - obj: The GoMixObject to check (should be a Float)
//   - expected: The expected float64 value
//
// Example usage in tests:
//
//	result := ev.Eval(parseExpression("3.14 + 2.86"))
//	AssertFloat(t, result, 6.0)
//	// Passes if result is Float with value 6.0
func AssertFloat(t *testing.T, obj std.GoMixObject, expected float64) {
	result, ok := obj.(*std.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
	}

}

// AssertNil is a test helper function that validates that an object is nil.
//
// This function checks that the provided object is nil, which is useful for
// testing statements that don't produce a value or for verifying cleanup
// operations. Note that this checks for Go nil, not the Go-Mix Nil object type.
//
// If the object is not nil, the test is marked as failed with a descriptive
// message showing the actual type and value received.
//
// Parameters:
//   - t: The testing.T instance for reporting test failures
//   - obj: The GoMixObject to check (should be nil)
//
// Example usage in tests:
//
//	result := someOperation()
//	AssertNil(t, result)
//	// Passes if result is nil (not a Nil object, but actual Go nil)
//
// Note: To check for Go-Mix Nil objects (the language's nil value), compare
// the type instead: obj.GetType() == objects.NilType
func AssertNil(t *testing.T, obj std.GoMixObject) {
	if obj != nil {
		t.Errorf("object is not nil. got=%T (%+v)", obj, obj)
		return
	}
}

// AssertString is a test helper function that validates string objects and their values.
//
// This function performs two checks for testing string results:
// 1. Verifies that the object is a String type (not Integer, Boolean, etc.)
// 2. Checks that the string value exactly matches the expected value
//
// String comparison is case-sensitive and requires an exact match. If either
// check fails, the test is marked as failed with a descriptive message showing
// what was expected versus what was received.
//
// Parameters:
//   - t: The testing.T instance for reporting test failures
//   - obj: The GoMixObject to check (should be a String)
//   - expected: The expected string value
//
// Example usage in tests:
//
//	result := ev.Eval(parseExpression("\"Hello\" + \" World\""))
//	AssertString(t, result, "Hello World")
//	// Passes if result is String with value "Hello World"
func AssertString(t *testing.T, obj std.GoMixObject, expected string) {
	result, ok := obj.(*std.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q", result.Value, expected)
	}
}

// IndexOfDot finds the index of the first period (.) character in a string.
//
// This helper function is used by the evaluator to detect method calls in
// identifier names (e.g., "obj.method"). It scans the string from left to right.
//
// Parameters:
//   - s: The string to search
//
// Returns:
//   - int: The index of the first dot, or -1 if no dot is found
func IndexOfDot(s string) int {
	for i, c := range s {
		if c == '.' {
			return i
		}
	}
	return -1
}

// switchToFloat64 converts a GoMixObject to float64 for numeric comparisons.
func switchToFloat64(obj std.GoMixObject) float64 {
	if obj.GetType() == std.IntegerType {
		return float64(obj.(*std.Integer).Value)
	} else if obj.GetType() == std.FloatType {
		return obj.(*std.Float).Value
	}
	return 0
}
