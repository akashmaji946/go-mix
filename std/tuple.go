/*
File    : go-mix/std/tuple.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// This file implements built-in tuple manipulation methods for the Go-Mix language.
// It defines methods for creating and querying immutable tuple objects.
// These methods are registered as global builtins during package initialization.
package std

import (
	"io" // io.Writer is used for output in builtin functions
)

// tupleMethods is a slice of Builtin pointers representing the tuple manipulation functions.
// Each Builtin has a name (the method name) and a callback function that implements the behavior.
// These are appended to the global Builtins slice during package initialization.
var tupleMethods = []*Builtin{
	{
		Name:     "tuple", // Creates a new immutable tuple from arguments
		Callback: tupleFunc,
	},
	{
		Name:     "size_tuple", // Returns the number of elements in a tuple
		Callback: sizeTuple,
	},
	{
		Name:     "peekback_tuple", // Returns the last element of a tuple
		Callback: peekbackTuple,
	},
	{
		Name:     "peekfront_tuple", // Returns the first element of a tuple
		Callback: peekfrontTuple,
	},
	{
		Name:     "contains_tuple", // Checks if a value exists in the tuple
		Callback: containsTuple,
	},
}

// init registers the tuple methods by appending them to the global Builtins slice.
// This function runs automatically when the package is initialized.
// It also registers the tuple package for import functionality.
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, tupleMethods...)

	// Register as a package (for import functionality)
	tuplePackage := &Package{
		Name:      "tuple",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range tupleMethods {
		tuplePackage.Functions[method.Name] = method
	}
	RegisterPackage(tuplePackage)
}

// tupleFunc creates a new immutable tuple from the provided arguments.
// It takes zero or more arguments of any type and returns a Tuple object.
// Tuples are heterogeneous and immutable, preventing modifications after creation.
//
// Examples:
//
//	tuple()                   -> tuple()
//	tuple(1, 2, 3)           -> tuple(1, 2, 3)
//	tuple("a", 1, true)      -> tuple(a, 1, true)
func tupleFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	elements := make([]GoMixObject, len(args))
	copy(elements, args)
	return &Tuple{Elements: elements}
}

// sizeTuple returns the number of elements in a tuple.
// It takes one argument: the tuple.
// Returns an Integer object with the size.
//
// Examples:
//
//	var t = tuple(1, 2, 3);
//	size_tuple(t);           -> 3
func sizeTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != TupleType {
		return createError("ERROR: argument to `size_tuple` must be a tuple, got '%s'", args[0].GetType())
	}

	tuple := args[0].(*Tuple)
	return &Integer{Value: int64(len(tuple.Elements))}
}

// peekbackTuple returns the last element of a tuple.
// It takes one argument: the tuple.
// Returns the last element, or an error if the tuple is empty.
//
// Examples:
//
//	var t = tuple(1, 2, 3);
//	peekback_tuple(t);       -> 3
func peekbackTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != TupleType {
		return createError("ERROR: argument to `peekback_tuple` must be a tuple, got '%s'", args[0].GetType())
	}

	tuple := args[0].(*Tuple)
	if len(tuple.Elements) == 0 {
		return createError("ERROR: cannot peek from empty tuple")
	}

	return tuple.Elements[len(tuple.Elements)-1]
}

// peekfrontTuple returns the first element of a tuple.
// It takes one argument: the tuple.
// Returns the first element, or an error if the tuple is empty.
//
// Examples:
//
//	var t = tuple(1, 2, 3);
//	peekfront_tuple(t);      -> 1
func peekfrontTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].GetType() != TupleType {
		return createError("ERROR: argument to `peekfront_tuple` must be a tuple, got '%s'", args[0].GetType())
	}

	tuple := args[0].(*Tuple)
	if len(tuple.Elements) == 0 {
		return createError("ERROR: cannot peek from empty tuple")
	}

	return tuple.Elements[0]
}

// containsTuple checks if a value exists in a tuple.
// It takes two arguments: the tuple and the value to search for.
// Returns a Boolean true if the value is found, false otherwise.
// Comparison is done using the ToString() representation of objects.
//
// Examples:
//
//	var t = tuple(1, 2, 3, 4);
//	contains_tuple(t, 3);     -> true
//	contains_tuple(t, 5);     -> false
//	contains_tuple(t, "2");   -> false (type matters)
func containsTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].GetType() != TupleType {
		return createError("ERROR: first argument to `contains_tuple` must be a tuple, got '%s'", args[0].GetType())
	}

	tuple := args[0].(*Tuple)
	searchValue := args[1]

	// Search for the value in the tuple
	for _, elem := range tuple.Elements {
		// Compare both type and value
		if elem.GetType() == searchValue.GetType() && elem.ToString() == searchValue.ToString() {
			return &Boolean{Value: true}
		}
	}

	return &Boolean{Value: false}
}
