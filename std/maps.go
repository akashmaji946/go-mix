/*
File    : go-mix/std/maps.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// This file implements built-in map manipulation methods for the Go-Mix language.
// It defines methods like keys_map, insert_map, remove_map, contain_map, and enumerate_map
// that can be called on map objects.
// These methods are registered as global builtins during package initialization.
package std

import (
	"io" // io.Writer is used for output in builtin functions
)

// mapMethods is a slice of Builtin pointers representing the map manipulation functions.
// Each Builtin has a name (the method name) and a callback function that implements the behavior.
// These are appended to the global Builtins slice during package initialization.
var mapMethods = []*Builtin{
	{Name: "make_map", Callback: mapMake},           // Creates a new  map
	{Name: "keys_map", Callback: mapKeys},           // Returns an array of all keys in a map
	{Name: "insert_map", Callback: mapInsert},       // Inserts or updates a key-value pair in a map
	{Name: "remove_map", Callback: mapRemove},       // Removes a key-value pair from a map
	{Name: "contain_map", Callback: mapContain},     // Checks if a map contains a key
	{Name: "enumerate_map", Callback: mapEnumerate}, // Returns array of [key, value] pairs

	{Name: "size_map", Callback: mapSize},   // Returns the number of key-value pairs in a map
	{Name: "length_map", Callback: mapSize}, // Applies a function to each key-value pair, returning a new map

	{Name: "values_map", Callback: mapValues}, // Returns an array of all values in a map

}

// init is a special Go function that runs when the package is initialized.
// It registers the map methods as global builtins by appending them to the Builtins slice.
// It also registers the map package for import functionality.
//
// Import Examples:
//
//	import "maps"
//	var m = maps.make_map("name", "John", "age", 25)
//	var keys = maps.keys_map(m)
//
// Or with alias:
//
//	import "maps" as m
//	var map1 = m.make_map("a", 1, "b", 2)
//	var k = m.keys_map(map1)
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, mapMethods...)

	// Register as a package (for import functionality)
	mapsPackage := &Package{
		Name:      "maps",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range mapMethods {
		mapsPackage.Functions[method.Name] = method
	}
	RegisterPackage(mapsPackage)
}

// mapMake creates a new map object.
// Parameters:
//   - args: Optional key-value pairs to initialize the map (e.g. "name", "John", "age", 25)
//
// Returns:
//   - A new Map object initialized with the provided key-value pairs, or Error if arguments are invalid
//
// Example:
//
//	var m = make_map("name", "John", "age", 25); // Creates map{"name": "John", "age": 25}
func mapMake(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args)%2 != 0 {
		return createError("ERROR: make_map expects an even number of arguments (key-value pairs)")
	}

	m := &Map{
		Pairs: make(map[string]GoMixObject),
		Keys:  make([]string, 0, len(args)/2),
	}

	for i := 0; i < len(args); i += 2 {
		keyStr := args[i].ToString()
		value := args[i+1]

		// Check if key already exists
		if _, exists := m.Pairs[keyStr]; !exists {
			// New key, add to keys list
			m.Keys = append(m.Keys, keyStr)
		}

		m.Pairs[keyStr] = value
	}

	return m
}

// mapSize returns the number of key-value pairs in a map.
//
// Parameters:
//   - args[0]: The map to check
//
// Returns:
//   - Integer representing the number of key-value pairs, or Error if argument is not a map
//
// Example:
//
//	var m = make_map("name", "John", "age", 25);
//	size_map(m);  // Returns 2
func mapSize(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].GetType() != MapType {
		return createError("ERROR: argument to `size_map` must be a map, got '%s'", args[0].GetType())
	}

	mapObj := args[0].(*Map)
	return &Integer{Value: int64(len(mapObj.Pairs))}
}

// mapKeys returns an array of all keys in a map.
// The keys are returned in the order they were inserted.
//
// Parameters:
//   - args[0]: The map to get keys from
//
// Returns:
//   - Array of string keys, or Error if argument is not a map
//
// Example:
//
//	var m = map{"name": "John", "age": 25};
//	keys_map(m);  // Returns ["name", "age"]
func mapKeys(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].GetType() != MapType {
		return createError("ERROR: argument to `keys_map` must be a map, got '%s'", args[0].GetType())
	}

	mapObj := args[0].(*Map)
	keyObjects := make([]GoMixObject, len(mapObj.Keys))

	for i, key := range mapObj.Keys {
		keyObjects[i] = &String{Value: key}
	}

	return &Array{Elements: keyObjects}
}

// mapInsert adds or updates a key-value pair in a map.
// Maps are mutable, so this modifies the original map.
//
// Parameters:
//   - args[0]: The map to insert into
//   - args[1]: The key (will be converted to string)
//   - args[2]: The value to insert
//
// Returns:
//   - The inserted value, or Error if wrong arguments
//
// Example:
//
//	var m = map{"name": "John"};
//	insert_map(m, "age", 25);     // m is now map{"name": "John", "age": 25}
//	insert_map(m, "name", "Jane"); // m is now map{"name": "Jane", "age": 25}
func mapInsert(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: wrong number of arguments. got=%d, want=3", len(args))
	}

	if args[0].GetType() != MapType {
		return createError("ERROR: first argument to `insert_map` must be a map, got '%s'", args[0].GetType())
	}

	mapObj := args[0].(*Map)
	keyStr := args[1].ToString()
	value := args[2]

	// Check if key already exists
	if _, exists := mapObj.Pairs[keyStr]; !exists {
		// New key, add to keys list
		mapObj.Keys = append(mapObj.Keys, keyStr)
	}

	// Insert or update the value
	mapObj.Pairs[keyStr] = value

	return value
}

// mapRemove deletes a key-value pair from a map.
// Maps are mutable, so this modifies the original map.
//
// Parameters:
//   - args[0]: The map to remove from
//   - args[1]: The key to remove (will be converted to string)
//
// Returns:
//   - The removed value if key existed, nil otherwise, or Error if wrong arguments
//
// Example:
//
//	var m = map{"name": "John", "age": 25};
//	remove_map(m, "age");  // Returns 25, m is now map{"name": "John"}
//	remove_map(m, "city"); // Returns nil (key didn't exist)
func mapRemove(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].GetType() != MapType {
		return createError("ERROR: first argument to `remove_map` must be a map, got '%s'", args[0].GetType())
	}

	mapObj := args[0].(*Map)
	keyStr := args[1].ToString()

	// Check if key exists
	value, exists := mapObj.Pairs[keyStr]
	if !exists {
		return &Nil{}
	}

	// Remove from pairs
	delete(mapObj.Pairs, keyStr)

	// Remove from keys list
	for i, k := range mapObj.Keys {
		if k == keyStr {
			mapObj.Keys = append(mapObj.Keys[:i], mapObj.Keys[i+1:]...)
			break
		}
	}

	return value
}

// mapContain checks if a map contains a specific key.
//
// Parameters:
//   - args[0]: The map to check
//   - args[1]: The key to look for (will be converted to string)
//
// Returns:
//   - Boolean true if key exists, false otherwise, or Error if wrong arguments
//
// Example:
//
//	var m = map{"name": "John", "age": 25};
//	contain_map(m, "name");  // Returns true
//	contain_map(m, "city");  // Returns false
func mapContain(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: wrong number of arguments. got=%d, want=2", len(args))
	}

	if args[0].GetType() != MapType {
		return createError("ERROR: first argument to `contain_map` must be a map, got '%s'", args[0].GetType())
	}

	mapObj := args[0].(*Map)
	keyStr := args[1].ToString()

	_, exists := mapObj.Pairs[keyStr]
	return &Boolean{Value: exists}
}

// mapEnumerate returns an array of [key, value] pairs from a map.
// The pairs are returned in the order the keys were inserted.
//
// Parameters:
//   - args[0]: The map to enumerate
//
// Returns:
//   - Array of arrays, where each inner array is [key, value], or Error if argument is not a map
//
// Example:
//
//	var m = map{"name": "John", "age": 25};
//	enumerate_map(m);  // Returns [["name", "John"], ["age", 25]]
func mapEnumerate(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].GetType() != MapType {
		return createError("ERROR: argument to `enumerate_map` must be a map, got '%s'", args[0].GetType())
	}

	mapObj := args[0].(*Map)
	pairs := make([]GoMixObject, len(mapObj.Keys))

	for i, key := range mapObj.Keys {
		pair := &Array{
			Elements: []GoMixObject{
				&String{Value: key},
				mapObj.Pairs[key],
			},
		}
		pairs[i] = pair
	}

	return &Array{Elements: pairs}
}

// mapValues returns an array of all values in a map.
// The values may not be returned in the order their keys were inserted.
//
// Parameters:
//   - args[0]: The map to get values from
//
// Returns:
//   - Array of values, or Error if argument is not a map
//
// Example:
//
//	var m = map{"name": "John", "age": 25};
//	values_map(m);  // Returns ["John", 25]
func mapValues(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].GetType() != MapType {
		return createError("ERROR: argument to `values_map` must be a map, got '%s'", args[0].GetType())
	}

	mapObj := args[0].(*Map)
	valueObjects := make([]GoMixObject, len(mapObj.Keys))

	for i, key := range mapObj.Keys {
		valueObjects[i] = mapObj.Pairs[key]
	}

	return &Array{Elements: valueObjects}
}
