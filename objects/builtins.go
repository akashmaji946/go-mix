// Package objects - builtins.go
// This file defines the builtin functions available in the GoMix language.
// It includes common functions like print, println, printf, length, and tostring,
// as well as utility functions for error creation and type conversion.
// These builtins are registered globally and can be called from GoMix code.
/*
File    : go-mix/objects/builtins.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package objects

import (
	"fmt" // fmt is used for string formatting and printing
	"io"  // io.Writer is used for output operations in builtin functions
)

// CallbackFunc is the function signature for builtin functions.
// It takes an io.Writer for output (e.g., console) and a variadic list of GoMixObject arguments,
// returning a GoMixObject result (or an error if something goes wrong).
type CallbackFunc func(writer io.Writer, args ...GoMixObject) GoMixObject

// Builtin represents a builtin function with a name and its implementation callback.
// This struct is used to store and invoke builtin functions in the language.
type Builtin struct {
	Name     string       // The name of the builtin function (e.g., "print")
	Callback CallbackFunc // The function that implements the builtin behavior
}

// Builtins is a global slice of pointers to Builtin structs.
// It holds all the builtin functions available in the GoMix language.
// Functions are added to this slice during package initialization.
var Builtins = make([]*Builtin, 0)

// commonMethods is a slice of common builtin functions that are always available.
// These include printing functions, length calculation, and string conversion.
var commonMethods = []*Builtin{
	{
		Name:     "print", // Prints arguments without a newline
		Callback: print,
	},
	{
		Name:     "println", // Prints arguments with a newline
		Callback: println,
	},
	{
		Name:     "printf", // Prints formatted string with arguments
		Callback: printf,
	},
	{
		Name:     "length", // Returns the length of strings or arrays
		Callback: length,
	},
	{
		Name:     "tostring", // Converts an object to its string representation
		Callback: tostring,
	},
	{
		Name:     "range", // Creates an inclusive range from start to end
		Callback: rangeFunc,
	},
	{
		Name:     "typeof", // Returns the type of a GoMix object as a string
		Callback: typeofFunc,
	},
	{
		Name:     "size", // Alias for length - returns the size of strings, arrays, maps, or sets
		Callback: length,
	},
}

// init registers the common builtin methods by appending them to the global Builtins slice.
// This function runs automatically when the package is initialized.
func init() {
	Builtins = append(Builtins, commonMethods...)
}

// createError is a utility function to create an Error object with a formatted message.
// It takes a format string and variadic arguments, similar to fmt.Sprintf,
// and returns a pointer to an Error struct with the formatted message.
func createError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

// tostring converts a GoMixObject to its string representation, wrapped in quotes.
// It takes one argument: the object to convert.
// Returns a String object containing the quoted string representation of the input.
func tostring(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly one argument is provided
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	// Return the string representation wrapped in quotes
	return &String{Value: fmt.Sprintf("\"%s\"", args[0].ToString())}
}

// print outputs the string representations of its arguments to the writer without a trailing newline.
// It takes one or more arguments to print.
// Returns a Nil object. Flushes the writer if it supports syncing (e.g., for buffered output).
func print(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if at least one argument is provided
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1 or more", len(args))
	}
	// Build the output string by concatenating string representations with spaces
	res := ""
	for _, arg := range args {
		res += arg.ToString() + " "
	}
	// Remove the trailing space if there are arguments
	if len(args) > 0 {
		res = res[:len(res)-1]
	}
	// Print to the writer
	fmt.Fprint(writer, res)
	// Flush the writer if it has a Sync method (e.g., bufio.Writer)
	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}
	// Return nil as the result
	return &Nil{}
}

// println outputs the string representations of its arguments to the writer with a trailing newline.
// It takes one or more arguments to print.
// Returns a Nil object. Flushes the writer if it supports syncing.
func println(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if at least one argument is provided
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1 or more", len(args))
	}
	// Build the output string by concatenating string representations with spaces
	res := ""
	for _, arg := range args {
		res += arg.ToString() + " "
	}
	// Remove the trailing space if there are arguments
	if len(args) > 0 {
		res = res[:len(res)-1]
	}
	// Print to the writer with a newline
	fmt.Fprintln(writer, res)
	// Flush the writer if it has a Sync method
	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}
	// Return nil as the result
	return &Nil{}
}

// printf outputs a formatted string to the writer using Go's fmt.Printf style formatting.
// The first argument must be a format string, followed by arguments to format.
// It extracts raw values from GoMixObjects for formatting.
// Returns a Nil object. Flushes the writer if it supports syncing.
func printf(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if at least one argument (the format string) is provided
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1 or more", len(args))
	}
	// Ensure the first argument is a string (the format)
	if args[0].GetType() != StringType {
		return createError("ERROR: first argument to `printf` must be a string, got `%s`", args[0].GetType())
	}
	// Get the format string
	format := args[0].ToString()
	// Prepare arguments by extracting raw values
	arguments := make([]interface{}, len(args)-1)
	for i, arg := range args[1:] {
		val, err := ExtractValue(arg)
		if err != nil {
			return &Error{Message: err.Error()}
		}
		arguments[i] = val
	}
	// Print the formatted string to the writer
	fmt.Fprintf(writer, format, arguments...)
	// Flush the writer if it has a Sync method
	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}
	// Return nil as the result
	return &Nil{}
}

// length returns the length of a string, array, map, or set as an Integer object.
// It takes one argument: the string, array, map, or set to measure.
// For strings, returns the number of characters; for arrays, returns the number of elements;
// for maps, returns the number of key-value pairs; for sets, returns the number of unique values.
// Returns an error for unsupported types.
func length(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly one argument is provided
	if len(args) != 1 {
		return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
	}
	// Determine the type and calculate length accordingly
	switch args[0].GetType() {
	case StringType:
		// Return the length of the string
		return &Integer{Value: int64(len(args[0].ToString()))}
	case ArrayType:
		// Return the number of elements in the array
		return &Integer{Value: int64(len(args[0].(*Array).Elements))}
	case MapType:
		// Return the number of key-value pairs in the map
		return &Integer{Value: int64(len(args[0].(*Map).Keys))}
	case SetType:
		// Return the number of unique values in the set
		return &Integer{Value: int64(len(args[0].(*Set).Values))}
	default:
		// Return an error for unsupported types
		return &Error{Message: fmt.Sprintf("argument to `length` not supported, got '%s'", args[0].GetType())}
	}
}

// rangeFunc creates an inclusive range from start to end, similar to the ... operator.
// It takes two arguments: start and end (both must be integers).
// Returns a Range object that can be used in foreach loops or stored in variables.
// This provides a functional alternative to the ... operator syntax.
//
// Examples:
//
//	range(2, 5)    -> Range{Start: 2, End: 5}
//	range(1, 10)   -> Range{Start: 1, End: 10}
func rangeFunc(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly two arguments are provided
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}

	// Validate that both arguments are integers
	if args[0].GetType() != IntegerType {
		return createError("ERROR: first argument to `range` must be an integer, got '%s'", args[0].GetType())
	}
	if args[1].GetType() != IntegerType {
		return createError("ERROR: second argument to `range` must be an integer, got '%s'", args[1].GetType())
	}

	// Extract the integer values
	start := args[0].(*Integer).Value
	end := args[1].(*Integer).Value

	// Create and return the Range object
	return &Range{
		Start: start,
		End:   end,
	}
}

// typeofFunc returns the type of a GoMix object as a string.
// It takes one argument: the object whose type should be determined.
// Returns a String object containing the type name (e.g., "int", "string", "array", "func", etc.).
// This is useful for runtime type checking and debugging.
//
// Examples:
//
//	typeof(42)           -> "int"
//	typeof(3.14)         -> "float"
//	typeof("hello")      -> "string"
//	typeof(true)         -> "bool"
//	typeof(nil)          -> "nil"
//	typeof([1, 2, 3])    -> "array"
//	typeof(range(1, 5))  -> "range"
//	typeof(myFunc)       -> "func"
func typeofFunc(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly one argument is provided
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	// Get the type of the argument and return it as a string
	return &String{Value: string(args[0].GetType())}
}
