/*
File    : go-mix/std/arrays.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// This file implements built-in array manipulation methods for the Go-Mix language.
// It defines methods like push, pop, shift, and unshift that can be called on array objects.
// These methods are registered as global builtins during package initialization.
package std

import (
	"io" // io.Writer is used for output in builtin functions, though not directly in this file
	"sort"
)

// arrayMethods is a slice of Builtin pointers representing the array manipulation functions.
// Each Builtin has a name (the method name) and a callback function that implements the behavior.
// These are appended to the global Builtins slice during package initialization.
var arrayMethods = []*Builtin{
	{Name: "push_array", Callback: push},          // Adds an element to the end of the array
	{Name: "pop_array", Callback: pop},            // Removes and returns the last element of the array
	{Name: "shift_array", Callback: shift},        // Removes and returns the first element of the array
	{Name: "unshift_array", Callback: unshift},    // Adds an element to the beginning of the array
	{Name: "sort_array", Callback: sortArray},     // Sorts the elements of the array in-place
	{Name: "sorted_array", Callback: sortedArray}, // Returns a new sorted array
	{Name: "clone_array", Callback: cloneArray},   // Returns a shallow copy of the array
	{Name: "csort_array", Callback: csort},        // Custom sort for an array using a comparator
	{Name: "csorted_array", Callback: csorted},    // Returns a new sorted array using a comparator

	{Name: "find_array", Callback: findArray},   // Finds the first element matching a predicate
	{Name: "some_array", Callback: someArray},   // Checks if at least one element matches
	{Name: "every_array", Callback: everyArray}, // Checks if all elements match

	{Name: "reverse_array", Callback: reverseArray}, // Returns a new reversed array

	{Name: "to_array", Callback: toArray},         // Converts list/tuple to array
	{Name: "map_array", Callback: mapArray},       // Applies a function to each element
	{Name: "filter_array", Callback: filterArray}, // Filters elements based on a predicate
	{Name: "reduce_array", Callback: reduceArray}, // Accumulates a value across an array
}

// init is a special Go function that runs when the package is initialized.
// It registers the array methods as global builtins by appending them to the Builtins slice.
// It also registers the arrays package for import functionality.
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, arrayMethods...)

	// Register as a package (for import functionality)
	arraysPackage := &Package{
		Name:      "arrays",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range arrayMethods {
		arraysPackage.Functions[method.Name] = method
	}
	RegisterPackage(arraysPackage)
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
func push(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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
	// Append the new element in-place
	arr.Elements = append(arr.Elements, args[1])

	// Return the modified array
	return arr
}

// toArray converts a list or tuple to an array.
// Syntax: to_array(iterable)
func toArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: to_array expects 1 argument")
	}
	arg := args[0]
	switch arg.GetType() {
	case ArrayType:
		return arg
	case ListType:
		l := arg.(*List)
		newElements := make([]GoMixObject, len(l.Elements))
		copy(newElements, l.Elements)
		return &Array{Elements: newElements}
	case TupleType:
		t := arg.(*Tuple)
		newElements := make([]GoMixObject, len(t.Elements))
		copy(newElements, t.Elements)
		return &Array{Elements: newElements}
	default:
		return createError("ERROR: argument to `to_array` must be a list or tuple, got '%s'", arg.GetType())
	}
}

// sortArray sorts the elements of an array in-place.
// Currently performs lexicographical sorting based on ToString().
//
// Syntax: sort_array(array)
// Syntax: sort_array(array, [reverse])
func sortArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) < 1 || len(args) > 2 {
		return createError("ERROR: sort_array expects 1 or 2 arguments (array, [reverse])")
	}
	if args[0].GetType() != ArrayType {
		return createError("ERROR: argument to `sort_array` must be an array")
	}

	reverse := false
	if len(args) == 2 {
		if args[1].GetType() != BooleanType {
			return createError("ERROR: second argument to `sort_array` must be a boolean")
		}
		reverse = args[1].(*Boolean).Value
	}

	arr := args[0].(*Array)

	sort.Slice(arr.Elements, func(i, j int) bool {
		// Basic implementation: compare string representations
		// You could expand this to check types (e.g., numeric sort for ints)
		if arr.Elements[i].GetType() == IntegerType && arr.Elements[j].GetType() == IntegerType {
			v1 := arr.Elements[i].(*Integer).Value
			v2 := arr.Elements[j].(*Integer).Value
			if reverse {
				return v1 > v2
			}
			return v1 < v2
		}
		s1 := arr.Elements[i].ToString()
		s2 := arr.Elements[j].ToString()
		if reverse {
			return s1 > s2
		}
		return s1 < s2
	})

	return arr
}

// sortedArray returns a new array with elements sorted.
// It does not modify the original array.
//
// Syntax: sorted_array(array)
// Syntax: sorted_array(array, [reverse])
func sortedArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) < 1 || len(args) > 2 {
		return createError("ERROR: sorted_array expects 1 or 2 arguments (array, [reverse])")
	}
	if args[0].GetType() != ArrayType {
		return createError("ERROR: argument to `sorted_array` must be an array, got '%s'", args[0].GetType())
	}

	reverse := false
	if len(args) == 2 {
		if args[1].GetType() != BooleanType {
			return createError("ERROR: second argument to `sorted_array` must be a boolean")
		}
		reverse = args[1].(*Boolean).Value
	}

	arr := args[0].(*Array)
	newElements := make([]GoMixObject, len(arr.Elements))
	copy(newElements, arr.Elements)

	sort.Slice(newElements, func(i, j int) bool {
		if newElements[i].GetType() == IntegerType && newElements[j].GetType() == IntegerType {
			v1 := newElements[i].(*Integer).Value
			v2 := newElements[j].(*Integer).Value
			if reverse {
				return v1 > v2
			}
			return v1 < v2
		}
		s1 := newElements[i].ToString()
		s2 := newElements[j].ToString()
		if reverse {
			return s1 > s2
		}
		return s1 < s2
	})

	return &Array{Elements: newElements}
}

// cloneArray returns a shallow copy of the array.
//
// Syntax: clone_array(array)
func cloneArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: clone_array expects 1 argument (array)")
	}
	if args[0].GetType() != ArrayType {
		return createError("ERROR: argument to `clone_array` must be an array, got '%s'", args[0].GetType())
	}

	arr := args[0].(*Array)
	newElements := make([]GoMixObject, len(arr.Elements))
	copy(newElements, arr.Elements)

	return &Array{Elements: newElements}
}

// csort performs an in-place sort of an array using a custom comparator.
//
// Syntax: csort(array, comparator)
func csort(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: csort expects 2 arguments (array, comparator)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `csort` must be an array")
	}
	cmp := args[1]
	if cmp.GetType() != FunctionType {
		return createError("ERROR: second argument to `csort` must be a function")
	}

	var sortErr GoMixObject
	sort.Slice(arr.Elements, func(i, j int) bool {
		if sortErr != nil {
			return false
		}
		// Call the Go-Mix comparator function
		res := rt.CallFunction(cmp, arr.Elements[i], arr.Elements[j])
		if res.GetType() == ErrorType {
			sortErr = res
			return false
		}
		if b, ok := res.(*Boolean); ok {
			return b.Value
		}
		return false
	})

	if sortErr != nil {
		return sortErr
	}

	return arr
}

// csorted returns a new array with elements sorted using a custom comparator.
// It does not modify the original array.
//
// Syntax: csorted(array, comparator)
func csorted(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: csorted expects 2 arguments (array, comparator)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `csorted` must be an array")
	}
	cmp := args[1]
	if cmp.GetType() != FunctionType {
		return createError("ERROR: second argument to `csorted` must be a function")
	}

	// Create a shallow copy
	newElements := make([]GoMixObject, len(arr.Elements))
	copy(newElements, arr.Elements)

	var sortErr GoMixObject
	sort.Slice(newElements, func(i, j int) bool {
		if sortErr != nil {
			return false
		}
		// Call the Go-Mix comparator function
		res := rt.CallFunction(cmp, newElements[i], newElements[j])
		if res.GetType() == ErrorType {
			sortErr = res
			return false
		}
		if b, ok := res.(*Boolean); ok {
			return b.Value
		}
		return false
	})

	if sortErr != nil {
		return sortErr
	}

	return &Array{Elements: newElements}
}

// mapArray applies a function to each element of an array and returns a new array.
//
// Syntax: map_array(array, function)
func mapArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: map_array expects 2 arguments (array, function)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `map_array` must be an array, got '%s'", args[0].GetType())
	}
	fn := args[1]
	if fn.GetType() != FunctionType {
		return createError("ERROR: second argument to `map_array` must be a function, got '%s'", fn.GetType())
	}

	newElements := make([]GoMixObject, len(arr.Elements))
	for i, elem := range arr.Elements {
		res := rt.CallFunction(fn, elem)
		if res.GetType() == ErrorType {
			return res
		}
		newElements[i] = res
	}

	return &Array{Elements: newElements}
}

// filterArray returns a new array containing elements that satisfy a predicate.
//
// Syntax: filter_array(array, function)
func filterArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: filter_array expects 2 arguments (array, function)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `filter_array` must be an array, got '%s'", args[0].GetType())
	}
	fn := args[1]
	if fn.GetType() != FunctionType {
		return createError("ERROR: second argument to `filter_array` must be a function, got '%s'", fn.GetType())
	}

	newElements := []GoMixObject{}
	for _, elem := range arr.Elements {
		res := rt.CallFunction(fn, elem)
		if res.GetType() == ErrorType {
			return res
		}

		// Check if the result is a boolean true
		isMatch := false
		if b, ok := res.(*Boolean); ok {
			isMatch = b.Value
		} else if i, ok := res.(*Integer); ok {
			isMatch = i.Value != 0
		}

		if isMatch {
			newElements = append(newElements, elem)
		}
	}

	return &Array{Elements: newElements}
}

// reduceArray accumulates a value by applying a function to each element of an array.
//
// Syntax: reduce_array(array, function, initial)
func reduceArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: reduce_array expects 3 arguments (array, function, initial)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `reduce_array` must be an array, got '%s'", args[0].GetType())
	}
	fn := args[1]
	if fn.GetType() != FunctionType {
		return createError("ERROR: second argument to `reduce_array` must be a function, got '%s'", fn.GetType())
	}

	accumulator := args[2]
	for _, elem := range arr.Elements {
		res := rt.CallFunction(fn, accumulator, elem)
		if res.GetType() == ErrorType {
			return res
		}
		accumulator = res
	}

	return accumulator
}

// findArray returns the first element that satisfies the provided testing function.
// If no values satisfy the testing function, nil is returned.
func findArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: find_array expects 2 arguments (array, function)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `find_array` must be an array")
	}
	fn := args[1]

	for _, elem := range arr.Elements {
		res := rt.CallFunction(fn, elem)
		if IsTruthy(res) {
			return elem
		}
	}
	return &Nil{}
}

// someArray tests whether at least one element in the array passes the test.
func someArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: some_array expects 2 arguments (array, function)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `some_array` must be an array")
	}
	fn := args[1]

	for _, elem := range arr.Elements {
		res := rt.CallFunction(fn, elem)
		if IsTruthy(res) {
			return &Boolean{Value: true}
		}
	}
	return &Boolean{Value: false}
}

// everyArray tests whether all elements in the array pass the test.
func everyArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: every_array expects 2 arguments (array, function)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: first argument to `every_array` must be an array")
	}
	fn := args[1]

	for _, elem := range arr.Elements {
		res := rt.CallFunction(fn, elem)
		if !IsTruthy(res) {
			return &Boolean{Value: false}
		}
	}
	return &Boolean{Value: true}
}

// reverseArray returns a new array with the elements in reverse order.
func reverseArray(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: reverse_array expects 1 argument (array)")
	}
	arr, ok := args[0].(*Array)
	if !ok {
		return createError("ERROR: argument to `reverse_array` must be an array")
	}

	n := len(arr.Elements)
	newElements := make([]GoMixObject, n)
	for i, elem := range arr.Elements {
		newElements[n-1-i] = elem
	}
	return &Array{Elements: newElements}
}

// IsTruthy is a helper to determine the boolean value of a GoMixObject.
func IsTruthy(obj GoMixObject) bool {
	switch v := obj.(type) {
	case *Boolean:
		return v.Value
	case *Integer:
		return v.Value != 0
	case *Nil:
		return false
	case *String:
		return len(v.Value) > 0
	default:
		return true
	}
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
func pop(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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
func shift(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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
func unshift(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
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
	// Prepend the new element in-place
	arr.Elements = append([]GoMixObject{args[1]}, arr.Elements...)

	// Return the modified array
	return arr
}
