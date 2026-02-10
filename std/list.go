/*
File    : go-mix/std/list.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// This file implements built-in list manipulation methods for the GoMix language.
// It defines methods for creating, modifying, and querying mutable list objects.
// These methods are registered as global builtins during package initialization.
package std

import (
	"io" // io.Writer is used for output in builtin functions
)

// listMethods is a slice of Builtin pointers representing the list manipulation functions.
// Each Builtin has a name (the method name) and a callback function that implements the behavior.
// These are appended to the global Builtins slice during package initialization.
var listMethods = []*Builtin{
	{Name: "list", Callback: listFunc},                // Creates a new mutable list from arguments
	{Name: "pushback_list", Callback: pushbackList},   // Appends an element to the end of a list
	{Name: "pushfront_list", Callback: pushfrontList}, // Prepends an element to the start of a list
	{Name: "popback_list", Callback: popbackList},     // Removes and returns the last element of a list
	{Name: "popfront_list", Callback: popfrontList},   // Removes and returns the first element of a list
	{Name: "size_list", Callback: sizeList},           // Returns the number of elements in a list
	{Name: "peekback_list", Callback: peekbackList},   // Returns the last element without removing it
	{Name: "peekfront_list", Callback: peekfrontList}, // Returns the first element without removing it
	{Name: "insert_list", Callback: insertList},       // Inserts an element at a specific index
	{Name: "remove_list", Callback: removeList},       // Removes an element at a specific index
	{Name: "contains_list", Callback: containsList},   // Checks if a value exists in the list
}

// init registers the list methods by appending them to the global Builtins slice.
// This function runs automatically when the package is initialized.
func init() {
	Builtins = append(Builtins, listMethods...)
}

// normalizeIndex converts a potentially negative index to a positive index.
// Supports Python-style negative indexing where -1 is the last element.
// Returns the normalized index and true if valid, or -1 and false if out of bounds.
func normalizeIndex(index int64, length int) (int, bool) {
	var actualIndex int
	if index < 0 {
		actualIndex = length + int(index)
	} else {
		actualIndex = int(index)
	}

	if actualIndex < 0 || actualIndex >= length {
		return -1, false
	}
	return actualIndex, true
}

// listFunc creates a new mutable list from the provided arguments.
// It takes zero or more arguments of any type and returns a List object.
// Lists are heterogeneous and mutable, allowing in-place modifications.
//
// Examples:
//
//	list()                    -> list()
//	list(1, 2, 3)            -> list(1, 2, 3)
//	list("a", 1, true)       -> list(a, 1, true)
func listFunc(writer io.Writer, args ...GoMixObject) GoMixObject {
	elements := make([]GoMixObject, len(args))
	copy(elements, args)
	return &List{Elements: elements}
}

// pushbackList appends an element to the end of a list (in-place mutation).
// It takes two arguments: the list and the value to append.
// Returns the modified list.
//
// Examples:
//
//	var l = list(1, 2, 3);
//	pushback_list(l, 4);     -> list(1, 2, 3, 4)
func pushbackList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: first argument to `pushback_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	list.Elements = append(list.Elements, args[1])
	return list
}

// pushfrontList prepends an element to the start of a list (in-place mutation).
// It takes two arguments: the list and the value to prepend.
// Returns the modified list.
//
// Examples:
//
//	var l = list(2, 3, 4);
//	pushfront_list(l, 1);    -> list(1, 2, 3, 4)
func pushfrontList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: first argument to `pushfront_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	list.Elements = append([]GoMixObject{args[1]}, list.Elements...)
	return list
}

// popbackList removes and returns the last element of a list (in-place mutation).
// It takes one argument: the list.
// Returns the removed element, or an error if the list is empty.
//
// Examples:
//
//	var l = list(1, 2, 3);
//	popback_list(l);         -> 3 (list becomes list(1, 2))
func popbackList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: argument to `popback_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	if len(list.Elements) == 0 {
		return createError("ERROR: cannot pop from empty list")
	}

	lastElement := list.Elements[len(list.Elements)-1]
	list.Elements = list.Elements[:len(list.Elements)-1]
	return lastElement
}

// popfrontList removes and returns the first element of a list (in-place mutation).
// It takes one argument: the list.
// Returns the removed element, or an error if the list is empty.
//
// Examples:
//
//	var l = list(1, 2, 3);
//	popfront_list(l);        -> 1 (list becomes list(2, 3))
func popfrontList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: argument to `popfront_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	if len(list.Elements) == 0 {
		return createError("ERROR: cannot pop from empty list")
	}

	firstElement := list.Elements[0]
	list.Elements = list.Elements[1:]
	return firstElement
}

// sizeList returns the number of elements in a list.
// It takes one argument: the list.
// Returns an Integer object with the size.
//
// Examples:
//
//	var l = list(1, 2, 3);
//	size_list(l);            -> 3
func sizeList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: argument to `size_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	return &Integer{Value: int64(len(list.Elements))}
}

// peekbackList returns the last element of a list without removing it.
// It takes one argument: the list.
// Returns the last element, or an error if the list is empty.
//
// Examples:
//
//	var l = list(1, 2, 3);
//	peekback_list(l);        -> 3 (list unchanged)
func peekbackList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: argument to `peekback_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	if len(list.Elements) == 0 {
		return createError("ERROR: cannot peek from empty list")
	}

	return list.Elements[len(list.Elements)-1]
}

// peekfrontList returns the first element of a list without removing it.
// It takes one argument: the list.
// Returns the first element, or an error if the list is empty.
//
// Examples:
//
//	var l = list(1, 2, 3);
//	peekfront_list(l);       -> 1 (list unchanged)
func peekfrontList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: argument to `peekfront_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	if len(list.Elements) == 0 {
		return createError("ERROR: cannot peek from empty list")
	}

	return list.Elements[0]
}

// insertList inserts an element at a specific index in a list (in-place mutation).
// It takes three arguments: the list, the index, and the value to insert.
// Supports negative indexing where -1 is after the last element.
// Returns the modified list, or an error if index is out of bounds.
//
// Examples:
//
//	var l = list(1, 2, 4);
//	insert_list(l, 2, 3);    -> list(1, 2, 3, 4)
//	insert_list(l, -1, 5);   -> list(1, 2, 3, 4, 5)
//	insert_list(l, 0, 0);    -> list(0, 1, 2, 3, 4, 5)
func insertList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: wrong number of arguments. got=%d, want=3", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: first argument to `insert_list` must be a list, got '%s'", args[0].GetType())
	}
	if args[1].GetType() != IntegerType {
		return createError("ERROR: second argument to `insert_list` must be an integer, got '%s'", args[1].GetType())
	}

	list := args[0].(*List)
	index := args[1].(*Integer).Value
	value := args[2]

	length := len(list.Elements)

	// Normalize negative index
	var actualIndex int
	if index < 0 {
		actualIndex = length + int(index) + 1 // For insertion, -1 means after last element
	} else {
		actualIndex = int(index)
	}

	// Check bounds (allow insertion at length for appending)
	if actualIndex < 0 || actualIndex > length {
		return createError("ERROR: list index out of bounds: index=%d, length=%d", index, length)
	}

	// Insert the element
	if actualIndex == length {
		// Append to end
		list.Elements = append(list.Elements, value)
	} else {
		// Insert in middle
		list.Elements = append(list.Elements[:actualIndex+1], list.Elements[actualIndex:]...)
		list.Elements[actualIndex] = value
	}

	return list
}

// removeList removes an element at a specific index from a list (in-place mutation).
// It takes two arguments: the list and the index.
// Supports negative indexing where -1 is the last element.
// Returns the removed element, or an error if index is out of bounds.
//
// Examples:
//
//	var l = list(1, 2, 3, 4);
//	remove_list(l, 2);       -> 3 (list becomes list(1, 2, 4))
//	remove_list(l, -1);      -> 4 (list becomes list(1, 2))
//	remove_list(l, 0);       -> 1 (list becomes list(2))
func removeList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: first argument to `remove_list` must be a list, got '%s'", args[0].GetType())
	}
	if args[1].GetType() != IntegerType {
		return createError("ERROR: second argument to `remove_list` must be an integer, got '%s'", args[1].GetType())
	}

	list := args[0].(*List)
	index := args[1].(*Integer).Value
	length := len(list.Elements)

	if length == 0 {
		return createError("ERROR: cannot remove from empty list")
	}

	// Normalize and validate index
	actualIndex, valid := normalizeIndex(index, length)
	if !valid {
		return createError("ERROR: list index out of bounds: index=%d, length=%d", index, length)
	}

	// Get the element to return
	removedElement := list.Elements[actualIndex]

	// Remove the element
	list.Elements = append(list.Elements[:actualIndex], list.Elements[actualIndex+1:]...)

	return removedElement
}

// containsList checks if a value exists in a list.
// It takes two arguments: the list and the value to search for.
// Returns a Boolean true if the value is found, false otherwise.
// Comparison is done using the ToString() representation of objects.
//
// Examples:
//
//	var l = list(1, 2, 3, 4);
//	contains_list(l, 3);     -> true
//	contains_list(l, 5);     -> false
//	contains_list(l, "2");   -> false (type matters)
func containsList(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].GetType() != ListType {
		return createError("ERROR: first argument to `contains_list` must be a list, got '%s'", args[0].GetType())
	}

	list := args[0].(*List)
	searchValue := args[1]

	// Search for the value in the list
	for _, elem := range list.Elements {
		// Compare both type and value
		if elem.GetType() == searchValue.GetType() && elem.ToString() == searchValue.ToString() {
			return &Boolean{Value: true}
		}
	}

	return &Boolean{Value: false}
}
