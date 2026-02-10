/*
File    : go-mix/std/common.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package std

// This file defines common types and utility functions used across
// multiple object types in the GoMix interpreter.
// It includes the definition of the Range type,
// which represents a sequence of integers defined by a start and end value.
// The file also contains helper functions for error handling and
// string conversion that are used by various builtin functions.
import (
	"fmt"
	"io"
)

// commonMethods is a slice of common builtin functions that are always available.
// These include printing functions, length calculation, and string conversion.
var commonMethods = []*Builtin{
	{Name: "print", Callback: print},       // Prints arguments without a newline
	{Name: "println", Callback: println},   // Prints arguments with a newline
	{Name: "printf", Callback: printf},     // Prints formatted string with arguments
	{Name: "length", Callback: length},     // Returns the length of strings or arrays
	{Name: "tostring", Callback: tostring}, // Converts an object to its string representation
	{Name: "range", Callback: rangeFunc},   // Creates an inclusive range from start to end
	{Name: "typeof", Callback: typeofFunc}, // Returns the type of a GoMix object as a string
	{Name: "size", Callback: length},       // Alias for length - returns the size of strings, arrays, maps, or sets
	{Name: "array", Callback: arrayFunc},   // Converts any iterable to a new array
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
	return &String{Value: fmt.Sprintf("%s", args[0].ToString())}
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
		res += arg.ToString() + ""
	}
	// Remove the trailing space if there are arguments
	if len(args) > 0 {
		res = res[:len(res)]
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

// length returns the length of a string, array, map, set, list, or tuple as an Integer object.
// It takes one argument: the string, array, map, set, list, or tuple to measure.
// For strings, returns the number of characters; for arrays, returns the number of elements;
// for maps, returns the number of key-value pairs; for sets, returns the number of unique values;
// for lists and tuples, returns the number of elements.
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
	case ListType:
		// Return the number of elements in the list
		return &Integer{Value: int64(len(args[0].(*List).Elements))}
	case TupleType:
		// Return the number of elements in the tuple
		return &Integer{Value: int64(len(args[0].(*Tuple).Elements))}
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

// arrayFunc creates a new array from arguments or converts an iterable to an array.
// It accepts zero or more arguments:
//   - 0 arguments: returns an empty array
//   - 1 iterable argument: converts the iterable to a new array
//   - 1 non-iterable argument: wraps it in an array
//   - Multiple arguments: creates an array containing all arguments
//
// Examples:
//
//	array()                    -> []
//	array(1, 2, 3)             -> [1, 2, 3]
//	array([1, 2, 3])           -> [1, 2, 3]
//	array(list(1, 2, 3))       -> [1, 2, 3]
//	array(tuple(1, 2, 3))      -> [1, 2, 3]
//	array(set{1, 2, 3})        -> [1, 2, 3]
//	array(map{"a": 1, "b": 2}) -> [1, 2]
//	array(42)                  -> [42]
func arrayFunc(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Handle 0 arguments: return empty array
	if len(args) == 0 {
		return &Array{Elements: []GoMixObject{}}
	}

	// Handle multiple arguments: create array from all arguments
	if len(args) > 1 {
		elements := make([]GoMixObject, len(args))
		copy(elements, args)
		return &Array{Elements: elements}
	}

	// Handle single argument
	arg := args[0]
	argType := arg.GetType()

	// Handle iterable types (convert to array)
	switch argType {
	case ArrayType:
		// Create a new array with copied elements from the input array
		arr := arg.(*Array)
		elements := make([]GoMixObject, len(arr.Elements))
		copy(elements, arr.Elements)
		return &Array{Elements: elements}

	case ListType:
		// Convert list elements to array
		list := arg.(*List)
		elements := make([]GoMixObject, len(list.Elements))
		copy(elements, list.Elements)
		return &Array{Elements: elements}

	case TupleType:
		// Convert tuple elements to array
		tuple := arg.(*Tuple)
		elements := make([]GoMixObject, len(tuple.Elements))
		copy(elements, tuple.Elements)
		return &Array{Elements: elements}

	case MapType:
		// Convert map values to array (in key insertion order)
		m := arg.(*Map)
		elements := make([]GoMixObject, len(m.Keys))
		for i, key := range m.Keys {
			elements[i] = m.Pairs[key]
		}
		return &Array{Elements: elements}

	case SetType:
		// Convert set values to array (in insertion order)
		s := arg.(*Set)
		elements := make([]GoMixObject, len(s.Values))
		for i, val := range s.Values {
			elements[i] = &String{Value: val}
		}
		return &Array{Elements: elements}

	case RangeType:
		// Convert range to array of integers
		r := arg.(*Range)
		start := r.Start
		end := r.End

		// Calculate size and direction
		var size int
		if start <= end {
			size = int(end - start + 1)
		} else {
			size = int(start - end + 1)
		}

		elements := make([]GoMixObject, size)
		if start <= end {
			// Ascending range
			for i := int64(0); i <= end-start; i++ {
				elements[i] = &Integer{Value: start + i}
			}
		} else {
			// Descending range
			for i := int64(0); i <= start-end; i++ {
				elements[i] = &Integer{Value: start - i}
			}
		}
		return &Array{Elements: elements}

	default:
		// Non-iterable single argument: wrap it in an array
		return &Array{Elements: []GoMixObject{arg}}
	}
}
