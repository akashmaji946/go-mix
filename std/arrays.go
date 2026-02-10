/*
File    : go-mix/std/arrays.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// This file implements built-in array manipulation methods for the GoMix language.
// It defines methods like push, pop, shift, and unshift that can be called on array objects.
// These methods are registered as global builtins during package initialization.
package std

import (
	"io" // io.Writer is used for output in builtin functions, though not directly in this file
)

// arrayMethods is a slice of Builtin pointers representing the array manipulation functions.
// Each Builtin has a name (the method name) and a callback function that implements the behavior.
// These are appended to the global Builtins slice during package initialization.
var arrayMethods = []*Builtin{
	{Name: "push", Callback: push},       // Adds an element to the end of the array
	{Name: "pop", Callback: pop},         // Removes and returns the last element of the array
	{Name: "shift", Callback: shift},     // Removes and returns the first element of the array
	{Name: "unshift", Callback: unshift}, // Adds an element to the beginning of the array
}

// init is a special Go function that runs when the package is initialized.
// It registers the array methods as global builtins by appending them to the Builtins slice.
func init() {
	Builtins = append(Builtins, arrayMethods...)
}

// push adds an element to the end of an array and returns the modified array.
// It takes two arguments: the array to modify and the element to add.
// If the arguments are invalid, it returns an error object.
//
// Syntax: push(array, element)
//
// Usage:
//
//	Appends the specified element to the end of the array.
//	Returns a new array containing all previous elements plus the new one.
//
// Example:
//
//	var a = [1, 2];
//	a = push(a, 3); // a is now [1, 2, 3]
func push(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly 2 arguments are provided
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	// Ensure the first argument is an array
	if args[0].GetType() != ArrayType {
		return createError("ERROR: first argument to `push` must be an array, got '%s'", args[0].GetType())
	}

	// Type assert to *Array
	arr := args[0].(*Array)
	// Create a new slice with space for the additional element
	newElements := make([]GoMixObject, len(arr.Elements)+1)
	// Copy existing elements
	copy(newElements, arr.Elements)
	// Add the new element at the end
	newElements[len(arr.Elements)] = args[1]

	// Return the modified array
	return &Array{Elements: newElements}
}

// pop removes and returns the last element from an array.
// It takes one argument: the array to modify.
// If the array is empty or arguments are invalid, it returns an error.
//
// Syntax: pop(array)
//
// Usage:
//
//	Removes the last element from the provided array and returns it.
//	This operation modifies the original array in-place.
//
// Example:
//
//	var a = [1, 2, 3];
//	var x = pop(a); // x is 3, a is now [1, 2]
func pop(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly 1 argument is provided
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	// Ensure the argument is an array
	if args[0].GetType() != ArrayType {
		return createError("ERROR: argument to `pop` must be an array, got '%s'", args[0].GetType())
	}

	// Type assert to *Array
	arr := args[0].(*Array)
	// Check if the array is empty
	if len(arr.Elements) == 0 {
		return createError("ERROR: cannot pop from empty array")
	}

	// Get the last element before removal
	lastElement := arr.Elements[len(arr.Elements)-1]

	// Modify the array by removing the last element (slice up to second last)
	arr.Elements = arr.Elements[:len(arr.Elements)-1]

	// Return the removed element
	return lastElement
}

// shift removes and returns the first element from an array.
// It takes one argument: the array to modify.
// If the array is empty or arguments are invalid, it returns an error.
//
// Syntax: shift(array)
//
// Usage:
//
//	Removes the first element from the provided array and returns it.
//	This operation modifies the original array in-place.
//
// Example:
//
//	var a = [1, 2, 3];
//	var x = shift(a); // x is 1, a is now [2, 3]
func shift(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly 1 argument is provided
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	// Ensure the argument is an array
	if args[0].GetType() != ArrayType {
		return createError("ERROR: argument to `shift` must be an array, got '%s'", args[0].GetType())
	}

	// Type assert to *Array
	arr := args[0].(*Array)
	// Check if the array is empty
	if len(arr.Elements) == 0 {
		return createError("ERROR: cannot shift from empty array")
	}

	// Get the first element before removal
	firstElement := arr.Elements[0]

	// Modify the array by removing the first element (slice from index 1)
	arr.Elements = arr.Elements[1:]

	// Return the removed element
	return firstElement
}

// unshift adds an element to the beginning of an array and returns the modified array.
// It takes two arguments: the array to modify and the element to add at the front.
// If the arguments are invalid, it returns an error object.
//
// Syntax: unshift(array, element)
//
// Usage:
//
//	Prepends the specified element to the beginning of the array.
//	Returns a new array containing the new element followed by all previous elements.
//
// Example:
//
//	var a = [2, 3];
//	a = unshift(a, 1); // a is now [1, 2, 3]
func unshift(writer io.Writer, args ...GoMixObject) GoMixObject {
	// Check if exactly 2 arguments are provided
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	// Ensure the first argument is an array
	if args[0].GetType() != ArrayType {
		return createError("ERROR: first argument to `unshift` must be an array, got '%s'", args[0].GetType())
	}

	// Type assert to *Array
	arr := args[0].(*Array)
	// Create a new slice with space for the additional element at the front
	newElements := make([]GoMixObject, len(arr.Elements)+1)
	// Place the new element at index 0
	newElements[0] = args[1]
	// Copy the existing elements starting from index 1
	copy(newElements[1:], arr.Elements)

	// Return the modified array
	return &Array{Elements: newElements}
}
