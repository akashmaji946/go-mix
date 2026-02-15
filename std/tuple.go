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

	{Name: "make_tuple", Callback: tupleFunc}, // Creates a new immutable tuple from arguments

	{Name: "peekback_tuple", Callback: peekbackTuple},   // Returns the last element of a tuple
	{Name: "peekfront_tuple", Callback: peekfrontTuple}, // Returns the first element of a tuple
	{Name: "contains_tuple", Callback: containsTuple},   // Checks if a value exists in the tuple

	{Name: "map_tuple", Callback: mapList},       // Applies a function to each element
	{Name: "filter_tuple", Callback: filterList}, // Filters elements based on a predicate
	{Name: "reduce_tuple", Callback: reduceList}, // Reduces the list to a single value using a binary function
	{Name: "find_tuple", Callback: findTuple},    // Finds the first element matching a predicate
	{Name: "some_tuple", Callback: someTuple},    // Checks if at least one element matches
	{Name: "every_tuple", Callback: everyTuple},  // Checks if all elements match

	{Name: "to_tuple", Callback: toTuple}, // Converts array/list to tuple

	{Name: "size_tuple", Callback: sizeTuple},   // Returns the number of elements in a tuple
	{Name: "length_tuple", Callback: sizeTuple}, // Returns the number of elements in a tuple (alias)
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

// toTuple converts an array or list to a tuple.
// Syntax: to_tuple(iterable)
func toTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: to_tuple expects 1 argument")
	}
	arg := args[0]
	switch arg.GetType() {
	case TupleType:
		return arg
	case ArrayType:
		a := arg.(*Array)
		newElements := make([]GoMixObject, len(a.Elements))
		copy(newElements, a.Elements)
		return &Tuple{Elements: newElements}
	case ListType:
		l := arg.(*List)
		newElements := make([]GoMixObject, len(l.Elements))
		copy(newElements, l.Elements)
		return &Tuple{Elements: newElements}
	default:
		return createError("ERROR: argument to `to_tuple` must be an array or list, got '%s'", arg.GetType())
	}
}

// findTuple returns the first element that satisfies the provided testing function.
// Syntax: find_tuple(tuple, function)
func findTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: find_tuple expects 2 arguments (tuple, function)")
	}
	if args[0].GetType() != TupleType {
		return createError("ERROR: first argument to `find_tuple` must be a tuple, got '%s'", args[0].GetType())
	}
	fn := args[1]
	if fn.GetType() != FunctionType {
		return createError("ERROR: second argument to `find_tuple` must be a function, got '%s'", fn.GetType())
	}

	tuple := args[0].(*Tuple)
	for _, elem := range tuple.Elements {
		res := rt.CallFunction(fn, elem)
		if IsTruthy(res) {
			return elem
		}
	}
	return &Nil{}
}

// someTuple tests whether at least one element in the tuple passes the test.
// Syntax: some_tuple(tuple, function)
func someTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: some_tuple expects 2 arguments (tuple, function)")
	}
	if args[0].GetType() != TupleType {
		return createError("ERROR: first argument to `some_tuple` must be a tuple, got '%s'", args[0].GetType())
	}
	fn := args[1]
	if fn.GetType() != FunctionType {
		return createError("ERROR: second argument to `some_tuple` must be a function, got '%s'", fn.GetType())
	}

	tuple := args[0].(*Tuple)
	for _, elem := range tuple.Elements {
		res := rt.CallFunction(fn, elem)
		if IsTruthy(res) {
			return &Boolean{Value: true}
		}
	}
	return &Boolean{Value: false}
}

// everyTuple tests whether all elements in the tuple pass the test.
// Syntax: every_tuple(tuple, function)
func everyTuple(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: every_tuple expects 2 arguments (tuple, function)")
	}
	if args[0].GetType() != TupleType {
		return createError("ERROR: first argument to `every_tuple` must be a tuple, got '%s'", args[0].GetType())
	}
	fn := args[1]
	if fn.GetType() != FunctionType {
		return createError("ERROR: second argument to `every_tuple` must be a function, got '%s'", fn.GetType())
	}

	tuple := args[0].(*Tuple)
	for _, elem := range tuple.Elements {
		res := rt.CallFunction(fn, elem)
		if !IsTruthy(res) {
			return &Boolean{Value: false}
		}
	}
	return &Boolean{Value: true}
}
