/*
File    : go-mix/objects/objects.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package objects defines the core data types and interfaces for the GoMix programming language.
// It provides implementations for primitive types (integers, floats, strings, booleans, nil, errors),
// composite types (arrays), and utility types (return values). All types implement the GoMixObject
// interface, which allows for type checking, string representation, and object inspection.
// This package also includes utility functions for extracting raw values from objects.
package objects

import (
	"fmt" // fmt is used for string formatting in ToString and ToObject methods
)

// GoMixType represents the type of a GoMix object as a string constant.
// These constants are used to identify the type of objects in the language,
// enabling type checking and polymorphic behavior across different object types.
type GoMixType string

const (
	// IntegerType represents 64-bit integer values
	IntegerType GoMixType = "int"
	// FloatType represents 64-bit floating-point values
	FloatType GoMixType = "float"
	// StringType represents string values
	StringType GoMixType = "string"
	// BooleanType represents boolean (true/false) values
	BooleanType GoMixType = "bool"
	// NilType represents null or undefined values
	NilType GoMixType = "nil"
	// ErrorType represents error objects with messages
	ErrorType GoMixType = "error"

	// FunctionType represents function objects (defined elsewhere)
	FunctionType GoMixType = "func"
	// ArrayType represents arrays of GoMix objects
	ArrayType GoMixType = "array"
	// RangeType represents range objects (inclusive ranges)
	RangeType GoMixType = "range"
)

// GoMixObject is the core interface that all GoMix objects must implement.
// It provides methods for type identification, string representation for display,
// and object inspection for debugging or serialization purposes.
type GoMixObject interface {
	// GetType returns the GoMixType of the object, used for type checking
	GetType() GoMixType
	// ToString returns a human-readable string representation of the object's value
	ToString() string
	// ToObject returns a detailed string representation including type information,
	// useful for debugging and object inspection
	ToObject() string
}

// ExtractValue extracts the raw Go value from a GoMixObject.
// This utility function is used when interfacing with Go's standard library
// or when performing operations that require native Go types.
// It returns the underlying value (e.g., int64 for Integer) or an error for unsupported types.
func ExtractValue(obj GoMixObject) (interface{}, error) {
	switch obj.GetType() {
	case IntegerType:
		// Extract the int64 value from an Integer object
		return obj.(*Integer).Value, nil
	case FloatType:
		// Extract the float64 value from a Float object
		return obj.(*Float).Value, nil
	case StringType:
		// Extract the string value from a String object
		return obj.(*String).Value, nil
	case BooleanType:
		// Extract the bool value from a Boolean object
		return obj.(*Boolean).Value, nil
	case NilType:
		// Extract the nil value from a Nil object
		return obj.(*Nil).Value, nil
	case ErrorType:
		// Extract the error message from an Error object
		return obj.(*Error).Message, nil
	default:
		// Return an error for unsupported types like functions or arrays
		return nil, fmt.Errorf("unsupported type: %s", obj.GetType())
	}
}

// The following section defines the concrete implementations of GoMixObject for each type.
// Each struct represents a GoMix data type and implements the GoMixObject interface.

// Integer represents a 64-bit signed integer value in GoMix.
// It wraps an int64 and provides methods for type identification and string conversion.
type Integer struct {
	Value int64 // The underlying integer value
}

// GetType returns the type of the Integer object
func (i *Integer) GetType() GoMixType {
	return IntegerType
}

// ToString returns the string representation of the integer value (e.g., "42")
func (i *Integer) ToString() string {
	return fmt.Sprintf("%d", i.Value)
}

// ToObject returns a detailed representation including type info (e.g., "<int(42)>")
func (i *Integer) ToObject() string {
	return fmt.Sprintf("<int(%d)>", i.Value)
}

// Float represents a 64-bit floating-point value in GoMix.
// It wraps a float64 and provides methods for type identification and string conversion.
type Float struct {
	Value float64 // The underlying floating-point value
}

// GetType returns the type of the Float object
func (f *Float) GetType() GoMixType {
	return FloatType
}

// ToString returns the string representation of the float value (e.g., "3.140000")
func (f *Float) ToString() string {
	return fmt.Sprintf("%f", f.Value)
}

// ToObject returns a detailed representation including type info (e.g., "<float(3.140000)>")
func (f *Float) ToObject() string {
	return fmt.Sprintf("<float(%f)>", f.Value)
}

// String represents a string value in GoMix.
// It wraps a Go string and provides methods for type identification and string conversion.
type String struct {
	Value string // The underlying string value
}

// GetType returns the type of the String object
func (s *String) GetType() GoMixType {
	return StringType
}

// ToString returns the string value itself (e.g., "hello")
func (s *String) ToString() string {
	return s.Value
}

// ToObject returns a detailed representation including type info (e.g., "<string(hello)>")
func (s *String) ToObject() string {
	return fmt.Sprintf("<string(%s)>", s.Value)
}

// Boolean represents a boolean value in GoMix.
// It wraps a Go bool and provides methods for type identification and string conversion.
type Boolean struct {
	Value bool // The underlying boolean value
}

// GetType returns the type of the Boolean object
func (b *Boolean) GetType() GoMixType {
	return BooleanType
}

// ToString returns the string representation of the boolean value (e.g., "true" or "false")
func (b *Boolean) ToString() string {
	return fmt.Sprintf("%t", b.Value)
}

// ToObject returns a detailed representation including type info (e.g., "<bool(true)>")
func (b *Boolean) ToObject() string {
	return fmt.Sprintf("<bool(%t)>", b.Value)
}

// Nil represents a null or undefined value in GoMix.
// It wraps an interface{} (which is typically nil) and provides methods for type identification.
type Nil struct {
	Value interface{} // The underlying value, usually nil
}

// GetType returns the type of the Nil object
func (n *Nil) GetType() GoMixType {
	return NilType
}

// ToString returns the string "nil"
func (n *Nil) ToString() string {
	return "nil"
}

// ToObject returns a detailed representation "<nil()>"
func (n *Nil) ToObject() string {
	return "<nil()>"
}

// Error represents an error object in GoMix.
// It wraps an error message as a string and provides methods for type identification and display.
type Error struct {
	Message string // The error message
}

// GetType returns the type of the Error object
func (e *Error) GetType() GoMixType {
	return ErrorType
}

// ToString returns the error message as a string
func (e *Error) ToString() string {
	return fmt.Sprintf("%s", e.Message)
}

// ToObject returns a detailed representation including type info (e.g., "<error(message)>")
func (e *Error) ToObject() string {
	return fmt.Sprintf("<error(%s)>", e.Message)
}

// ReturnValue wraps a value returned from a function in GoMix.
// It holds a GoMixObject and delegates type and string methods to the wrapped value.
// This is used to distinguish return values from regular expressions in the evaluator.
type ReturnValue struct {
	Value GoMixObject // The wrapped GoMix object returned from a function
}

// GetType returns the type of the wrapped value
func (r *ReturnValue) GetType() GoMixType {
	return r.Value.GetType()
}

// ToString returns the string representation of the wrapped value
func (r *ReturnValue) ToString() string {
	return r.Value.ToString()
}

// ToObject returns the object representation of the wrapped value
func (r *ReturnValue) ToObject() string {
	return r.Value.ToObject()
}

// Array represents an array of GoMix objects in GoMix.
// It holds a slice of GoMixObject elements and provides methods for type identification,
// string representation (as a comma-separated list), and object inspection.
type Array struct {
	Elements []GoMixObject // The slice of GoMix objects in the array
}

// GetType returns the type of the Array object
func (a *Array) GetType() GoMixType {
	return ArrayType
}

// ToString returns a string representation of the array as "[elem1, elem2, ...]"
// Each element is converted to its string representation using ToString()
func (a *Array) ToString() string {
	result := "["
	for i, elem := range a.Elements {
		if i > 0 {
			result += ", "
		}
		result += elem.ToString()
	}
	result += "]"
	return result
}

// ToObject returns a detailed representation of the array as "<array([elem1, elem2, ...])>"
// Each element is converted to its object representation using ToObject()
func (a *Array) ToObject() string {
	result := "<array(["
	for i, elem := range a.Elements {
		if i > 0 {
			result += ", "
		}
		result += elem.ToObject()
	}
	result += "])>"
	return result
}

// Range represents an inclusive range of integers in GoMix.
// It holds start and end values and provides methods for type identification
// and string representation. Ranges are used for iteration in foreach loops
// and can be created using the ... operator (e.g., 2...5).
type Range struct {
	Start int64 // The start value of the range (inclusive)
	End   int64 // The end value of the range (inclusive)
}

// GetType returns the type of the Range object
func (r *Range) GetType() GoMixType {
	return RangeType
}

// ToString returns a string representation of the range as "range(start,end)"
func (r *Range) ToString() string {
	return fmt.Sprintf("range(%d,%d)", r.Start, r.End)
}

// ToObject returns a detailed representation of the range as "<range(start,end)>"
func (r *Range) ToObject() string {
	return fmt.Sprintf("<range(%d,%d)>", r.Start, r.End)
}
