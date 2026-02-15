/*
File    : go-mix/std/common.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package std

// This file defines common types and utility functions used across
// multiple object types in the Go-Mix interpreter.
// It includes the definition of the Range type,
// which represents a sequence of integers defined by a start and end value.
// The file also contains helper functions for error handling and
// string conversion that are used by various builtin functions.
import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

// commonMethods is a slice of common builtin functions that are always available.
// These include printing functions, length calculation, and string conversion.
var commonMethods = []*Builtin{
	{Name: "print", Callback: print},     // Prints arguments without a newline
	{Name: "println", Callback: println}, // Prints arguments with a newline
	{Name: "printf", Callback: printf},   // Prints formatted string with arguments

	{Name: "length", Callback: length}, // Returns the length of strings or arrays
	{Name: "size", Callback: length},   // Alias for length - returns the size of strings, arrays, maps, or sets

	{Name: "to_string", Callback: tostring}, // Converts an object to its string representation
	// {Name: "string", Callback: tostring},                     // Alias for tostring - converts an object to a string
	{Name: "range", Callback: rangeFunc}, // Creates an inclusive range from start to end

	// constructors
	{Name: "array", Callback: arrayFunc}, // Converts any iterable to a new array
	{Name: "list", Callback: listFunc},   // Creates a new mutable list from arguments
	{Name: "tuple", Callback: tupleFunc}, // Creates a new immutable tuple from arguments

	// array methods
	{Name: "push", Callback: pushArray},       // Adds an element to the end of the array
	{Name: "pop", Callback: popArray},         // Removes and returns the last element of the array
	{Name: "shift", Callback: shiftArray},     // Removes and returns the first element of the array
	{Name: "unshift", Callback: unshiftArray}, // Adds an element to the beginning of the array
	{Name: "sort", Callback: sortArray},       // Sorts the elements of the array in-place
	{Name: "sorted", Callback: sortedArray},   // Returns a new sorted array
	{Name: "clone", Callback: cloneArray},     // Returns a shallow copy of the array
	{Name: "csort", Callback: csortArray},     // Custom sort for an array using a comparator
	{Name: "csorted", Callback: csortedArray}, // Returns a new sorted array using a comparator

	{Name: "find", Callback: findArray},   // Finds the first element matching a predicate
	{Name: "some", Callback: someArray},   // Checks if at least one element matches
	{Name: "every", Callback: everyArray}, // Checks if all elements match

	{Name: "reverse", Callback: reverseArray},   // Returns a new reversed array
	{Name: "contains", Callback: containsArray}, // Checks if a value exists in the array
	{Name: "replace", Callback: replaceArray},   // Returns the index of the first occurrence of a value, or -1 if not found
	{Name: "index", Callback: indexArray},       // Returns the index of the first occurrence of a value, or -1 if not found

	{Name: "json_string_to_map", Callback: jsonStringDecode}, // Converts a JSON string into a map
	{Name: "map_to_json_string", Callback: jsonStringEncode}, // Converts a map into a JSON string
	{Name: "json_encode", Callback: jsonStringEncode},        // Alias for

	{Name: "typeof", Callback: typeofFunc},     // Returns the type of a Go-Mix object as a string
	{Name: "addr", Callback: addrFunc},         // Returns the memory address of an object as an integer
	{Name: "is_same_ref", Callback: isSameRef}, // Checks if two objects point to the same memory address
}

// init registers the common builtin methods by appending them to the global Builtins slice.
// This function runs automatically when the package is initialized.
// It also registers the common package for import functionality.
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, commonMethods...)

	// Register as a package (for import functionality)
	commonPackage := &Package{
		Name:      "common",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range commonMethods {
		commonPackage.Functions[method.Name] = method
	}
	RegisterPackage(commonPackage)
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
func tostring(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly one argument is provided
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	// Return the string representation wrapped in quotes
	return &String{Value: args[0].ToString()}
}

// print outputs the string representations of its arguments to the writer without a trailing newline.
// It takes one or more arguments to print.
// Returns a Nil object. Flushes the writer if it supports syncing (e.g., for buffered output).
func print(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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
func println(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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
func printf(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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
func length(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly one argument is provided
	if len(args) != 1 {
		return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
	}
	// Determine the type and calculate length accordingly
	switch args[0].GetType() {
	case StringType:
		// Return the length of the string by directly accessing the Value field
		return &Integer{Value: int64(len(args[0].(*String).Value))}
	case ArrayType:
		// Return the number of elements in the array
		return &Integer{Value: int64(len(args[0].(*Array).Elements))}
	case MapType:
		// Return the number of key-value pairs in the map
		return &Integer{Value: int64(len(args[0].(*Map).Pairs))}
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
func rangeFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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

// typeofFunc returns the type of a Go-Mix object as a string.
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
func typeofFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly one argument is provided
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	// Get the type of the argument and return it as a string
	return &String{Value: string(args[0].GetType())}
}

// addrFunc returns the memory address of a GoMixObject as an integer.
// It uses Go's pointer representation to extract the address of the underlying value.
//
// Example:
//
//	var a = [1, 2];
//	println(addr(a)); // Prints the numeric memory address
func addrFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0] == nil {
		return &Integer{Value: 0}
	}

	rv := reflect.ValueOf(args[0])
	// Check if the underlying value is a pointer-like type that has an address
	if rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Map || rv.Kind() == reflect.Slice || rv.Kind() == reflect.Func {
		return &Integer{Value: int64(rv.Pointer())}
	}

	return createError("ERROR: could not determine memory address for type %s", args[0].GetType())
}

// isSameRef checks if two GoMixObjects point to the same memory address.
//
// Example:
//
//	var a = [1];
//	var b = a;
//	println(is_same_ref(a, b)); // true
func isSameRef(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: is_same_ref expects 2 arguments, got %d", len(args))
	}

	// Handle nil cases
	if args[0].GetType() == NilType && args[1].GetType() == NilType {
		return &Boolean{Value: true}
	}
	if args[0].GetType() == NilType || args[1].GetType() == NilType {
		return &Boolean{Value: false}
	}

	rv1 := reflect.ValueOf(args[0])
	rv2 := reflect.ValueOf(args[1])

	// Check if both are pointer-like types
	canHaveAddr1 := rv1.Kind() == reflect.Ptr || rv1.Kind() == reflect.Map || rv1.Kind() == reflect.Slice || rv1.Kind() == reflect.Func
	canHaveAddr2 := rv2.Kind() == reflect.Ptr || rv2.Kind() == reflect.Map || rv2.Kind() == reflect.Slice || rv2.Kind() == reflect.Func

	if !canHaveAddr1 || !canHaveAddr2 {
		return &Boolean{Value: false}
	}

	return &Boolean{Value: rv1.Pointer() == rv2.Pointer()}
}

// jsonStringDecode parses a JSON string into a Go-Mix Map.
//
// Parameters:
//   - args[0]: The JSON string to decode
//
// Returns:
//   - Map object, or Error if decoding fails
//
// Example:
//
//	var s = "{\"name\": \"John\", \"age\": 25}";
//	var m = json_string_decode(s);
func jsonStringDecode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: json_string_decode expects 1 argument (string)")
	}

	if args[0].GetType() != StringType {
		return createError("ERROR: argument to `json_string_decode` must be a string, got '%s'", args[0].GetType())
	}

	var data interface{}
	err := json.Unmarshal([]byte(args[0].ToString()), &data)
	if err != nil {
		return createError("ERROR: failed to decode JSON: %v", err)
	}

	return convertToGoMix(data)
}

// jsonStringEncode converts a Go-Mix object into a JSON string.
func jsonStringEncode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: json_string_encode expects 1 argument")
	}

	data := convertFromGoMix(args[0])
	bytes, err := json.Marshal(data)
	if err != nil {
		return createError("ERROR: failed to encode JSON: %v", err)
	}

	return &String{Value: string(bytes)}
}

func convertFromGoMix(obj GoMixObject) interface{} {
	switch obj.GetType() {
	case ArrayType:
		arr := obj.(*Array)
		res := make([]interface{}, len(arr.Elements))
		for i, e := range arr.Elements {
			res[i] = convertFromGoMix(e)
		}
		return res
	case MapType:
		m := obj.(*Map)
		res := make(map[string]interface{})
		for k, v := range m.Pairs {
			res[k] = convertFromGoMix(v)
		}
		return res
	case IntegerType:
		return obj.(*Integer).Value
	case FloatType:
		return obj.(*Float).Value
	case BooleanType:
		return obj.(*Boolean).Value
	case StringType:
		return obj.(*String).Value
	case NilType:
		return nil
	default:
		return obj.ToString()
	}
}

// convertToGoMix recursively converts Go native types from json.Unmarshal
// into Go-Mix internal objects.
func convertToGoMix(val interface{}) GoMixObject {
	switch v := val.(type) {
	case map[string]interface{}:
		m := &Map{
			Pairs: make(map[string]GoMixObject),
			Keys:  make([]string, 0, len(v)),
		}
		for k, rawVal := range v {
			m.Pairs[k] = convertToGoMix(rawVal)
			m.Keys = append(m.Keys, k)
		}
		return m
	case []interface{}:
		elements := make([]GoMixObject, len(v))
		for i, rawVal := range v {
			elements[i] = convertToGoMix(rawVal)
		}
		return &Array{Elements: elements}
	case string:
		return &String{Value: v}
	case bool:
		return &Boolean{Value: v}
	case float64:
		// Check if it's actually an integer
		if v == float64(int64(v)) {
			return &Integer{Value: int64(v)}
		}
		return &Float{Value: v}
	default:
		return &Nil{}
	}
}
