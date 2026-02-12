/*
File    : go-mix/std/format.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - format.go
// This file defines the type conversion builtin functions for the Go-Mix language.
// It provides functions for converting between integers, floats, booleans, and strings.
package std

import (
	"io"
	"strconv"
)

var formatMethods = []*Builtin{
	{Name: "to_int", Callback: toInt},       // Converts a value to an integer
	{Name: "to_float", Callback: toFloat},   // Converts a value to a float
	{Name: "to_bool", Callback: toBool},     // Converts a value to a boolean
	{Name: "to_string", Callback: toString}, // Converts a value to a string
	{Name: "to_char", Callback: toChar},     // Converts a value to a character
}

// init registers the format methods as global builtins.
func init() {
	Builtins = append(Builtins, formatMethods...)
}

// toInt converts a value to an integer.
//
// Syntax: to_int(value)
//
// Usage:
//   - If value is an integer, returns it as is.
//   - If value is a float, truncates it to an integer.
//   - If value is a boolean, returns 1 for true and 0 for false.
//   - If value is a string, parses it as an integer (supports 0x and 0 prefixes).
//   - If value is a character, returns its Unicode code point.
//
// Example:
//
//	to_int("123"); // Returns 123
//	to_int(3.14);  // Returns 3
//	to_int(true);  // Returns 1
func toInt(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: to_int expects 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch arg.GetType() {
	case IntegerType:
		return arg
	case FloatType:
		return &Integer{Value: int64(arg.(*Float).Value)}
	case BooleanType:
		if arg.(*Boolean).Value {
			return &Integer{Value: 1}
		}
		return &Integer{Value: 0}
	case StringType:
		val, err := strconv.ParseInt(arg.ToString(), 0, 64)
		if err != nil {
			return createError("ERROR: could not convert string to int: %v", err)
		}
		return &Integer{Value: val}
	case CharType:
		return &Integer{Value: int64(arg.(*Char).Value)}
	default:
		return createError("ERROR: cannot convert %s to int", arg.GetType())
	}
}

// toFloat converts a value to a floating-point number.
//
// Syntax: to_float(value)
//
// Usage:
//   - If value is a float, returns it as is.
//   - If value is an integer, converts it to a float.
//   - If value is a boolean, returns 1.0 for true and 0.0 for false.
//   - If value is a string, parses it as a float.
//
// Example:
//
//	to_float("3.14"); // Returns 3.14
//	to_float(123);    // Returns 123.0
//	to_float(false);  // Returns 0.0
func toFloat(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: to_float expects 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch arg.GetType() {
	case FloatType:
		return arg
	case IntegerType:
		return &Float{Value: float64(arg.(*Integer).Value)}
	case BooleanType:
		if arg.(*Boolean).Value {
			return &Float{Value: 1.0}
		}
		return &Float{Value: 0.0}
	case StringType:
		val, err := strconv.ParseFloat(arg.ToString(), 64)
		if err != nil {
			return createError("ERROR: could not convert string to float: %v", err)
		}
		return &Float{Value: val}
	default:
		return createError("ERROR: cannot convert %s to float", arg.GetType())
	}
}

// toBool converts a value to a boolean.
//
// Syntax: to_bool(value)
//
// Usage:
//   - If value is a boolean, returns it as is.
//   - If value is an integer, returns false for 0 and true otherwise.
//   - If value is a float, returns false for 0.0 and true otherwise.
//   - If value is a string, parses it using standard boolean rules (e.g., "true", "false", "1", "0").
//   - If value is nil, returns false.
//
// Example:
//
//	to_bool(1);       // Returns true
//	to_bool("false"); // Returns false
//	to_bool(nil);     // Returns false
func toBool(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: to_bool expects 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch arg.GetType() {
	case BooleanType:
		return arg
	case IntegerType:
		return &Boolean{Value: arg.(*Integer).Value != 0}
	case FloatType:
		return &Boolean{Value: arg.(*Float).Value != 0.0}
	case StringType:
		val, err := strconv.ParseBool(arg.ToString())
		if err != nil {
			return createError("ERROR: could not convert string to bool: %v", err)
		}
		return &Boolean{Value: val}
	case NilType:
		return &Boolean{Value: false}
	default:
		return &Boolean{Value: true}
	}
}

// toString converts a value to its string representation.
//
// Syntax: to_string(value)
//
// Example:
//
//	to_string(123);   // Returns "123"
//	to_string(true);  // Returns "true"
func toString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: to_string expects 1 argument, got %d", len(args))
	}
	return &String{Value: args[0].ToString()}
}

// toChar converts a value to a character.
//
// Syntax: to_char(value)
//
// Usage:
//   - If value is a character, returns it as is.
//   - If value is an integer, treats it as a Unicode code point.
//   - If value is a float, truncates it and treats as a Unicode code point.
//   - If value is a string, returns the first character of the string.
//
// Example:
//
//	to_char(65);    // Returns 'A'
//	to_char("abc"); // Returns 'a'
func toChar(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: to_char expects 1 argument, got %d", len(args))
	}

	arg := args[0]
	switch arg.GetType() {
	case CharType:
		return arg
	case IntegerType:
		return &Char{Value: rune(arg.(*Integer).Value)}
	case FloatType:
		return &Char{Value: rune(arg.(*Float).Value)}
	case StringType:
		s := arg.ToString()
		if len(s) == 0 {
			return createError("ERROR: cannot convert empty string to char")
		}
		return &Char{Value: []rune(s)[0]}
	default:
		return createError("ERROR: cannot convert %s to char", arg.GetType())
	}
}
