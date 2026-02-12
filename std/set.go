/*
File    : go-mix/std/set.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// This file implements built-in set manipulation methods for the Go-Mix language.
// It defines methods like insert_set, remove_set, contains_set, and values_set
// that can be called on set objects.
// These methods are registered as global builtins during package initialization.
package std

import (
	"io" // io.Writer is used for output in builtin functions
)

// setMethods is a slice of Builtin pointers representing the set manipulation functions.
// Each Builtin has a name (the method name) and a callback function that implements the behavior.
// These are appended to the global Builtins slice during package initialization.
var setMethods = []*Builtin{
	{Name: "insert_set", Callback: setInsert},     // Inserts a value into a set
	{Name: "remove_set", Callback: setRemove},     // Removes a value from a set
	{Name: "contains_set", Callback: setContains}, // Checks if a set contains a value
	{Name: "values_set", Callback: setValues},     // Returns an array of all values in a set
	{Name: "size_set", Callback: setSize},         // Returns the number of elements in a set
}

// init is a special Go function that runs when the package is initialized.
// It registers the set methods as global builtins by appending them to the Builtins slice.
func init() {
	Builtins = append(Builtins, setMethods...)
}

// setInsert adds a value to a set.
// Sets are mutable and only store unique values.
//
// Parameters:
//   - args[0]: The set to insert into
//   - args[1]: The value to insert (will be converted to string)
//
// Returns:
//   - The inserted value, or Error if wrong arguments
//
// Example:
//
//	var s = set{1, 2, 3};
//	insert_set(s, 4);     // s is now set{1, 2, 3, 4}
//	insert_set(s, 2);     // s remains set{1, 2, 3, 4} (2 already exists)
func setInsert(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].GetType() != SetType {
		return createError("ERROR: first argument to `insert_set` must be a set, got '%s'", args[0].GetType())
	}

	setObj := args[0].(*Set)
	valueStr := args[1].ToString()

	// Check if value already exists
	if !setObj.Elements[valueStr] {
		// New value, add to set
		setObj.Elements[valueStr] = true
		setObj.Values = append(setObj.Values, valueStr)
	}

	return args[1]
}

// setRemove deletes a value from a set.
// Sets are mutable, so this modifies the original set.
//
// Parameters:
//   - args[0]: The set to remove from
//   - args[1]: The value to remove (will be converted to string)
//
// Returns:
//   - Boolean true if value was removed, false if it didn't exist, or Error if wrong arguments
//
// Example:
//
//	var s = set{1, 2, 3};
//	remove_set(s, 2);  // Returns true, s is now set{1, 3}
//	remove_set(s, 5);  // Returns false (5 didn't exist)
func setRemove(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].GetType() != SetType {
		return createError("ERROR: first argument to `remove_set` must be a set, got '%s'", args[0].GetType())
	}

	setObj := args[0].(*Set)
	valueStr := args[1].ToString()

	// Check if value exists
	if !setObj.Elements[valueStr] {
		return &Boolean{Value: false}
	}

	// Remove from elements map
	delete(setObj.Elements, valueStr)

	// Remove from values list
	for i, v := range setObj.Values {
		if v == valueStr {
			setObj.Values = append(setObj.Values[:i], setObj.Values[i+1:]...)
			break
		}
	}

	return &Boolean{Value: true}
}

// setContains checks if a set contains a specific value.
//
// Parameters:
//   - args[0]: The set to check
//   - args[1]: The value to look for (will be converted to string)
//
// Returns:
//   - Boolean true if value exists, false otherwise, or Error if wrong arguments
//
// Example:
//
//	var s = set{1, 2, 3};
//	contains_set(s, 2);  // Returns true
//	contains_set(s, 5);  // Returns false
func setContains(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].GetType() != SetType {
		return createError("ERROR: first argument to `contains_set` must be a set, got '%s'", args[0].GetType())
	}

	setObj := args[0].(*Set)
	valueStr := args[1].ToString()

	return &Boolean{Value: setObj.Elements[valueStr]}
}

// setValues returns an array of all values in a set.
// The values are returned in the order they were inserted.
//
// Parameters:
//   - args[0]: The set to get values from
//
// Returns:
//   - Array of values, or Error if argument is not a set
//
// Example:
//
//	var s = set{1, 2, 3};
//	values_set(s);  // Returns [1, 2, 3]
func setValues(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].GetType() != SetType {
		return createError("ERROR: argument to `values_set` must be a set, got '%s'", args[0].GetType())
	}

	setObj := args[0].(*Set)
	valueObjects := make([]GoMixObject, len(setObj.Values))

	for i, value := range setObj.Values {
		valueObjects[i] = &String{Value: value}
	}

	return &Array{Elements: valueObjects}
}

// setSize returns the number of elements in a set.
// It takes one argument: the set.
// Returns an Integer object with the size.
//
// Example:
//
//	var s = set{1, 2, 3};
//	size_set(s);  // Returns 3
func setSize(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].GetType() != SetType {
		return createError("ERROR: argument to `size_set` must be a set, got '%s'", args[0].GetType())
	}

	setObj := args[0].(*Set)
	return &Integer{Value: int64(len(setObj.Values))}
}
