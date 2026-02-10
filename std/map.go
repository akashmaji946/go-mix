/*
File    : go-mix/std/map.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// This file implements built-in map manipulation methods for the GoMix language.
// It defines methods like keys_map, insert_map, remove_map, contain_map, and enumerate_map
// that can be called on map objects.
// These methods are registered as global builtins during package initialization.
package std

import (
	"encoding/json"
	"io" // io.Writer is used for output in builtin functions
)

// mapMethods is a slice of Builtin pointers representing the map manipulation functions.
// Each Builtin has a name (the method name) and a callback function that implements the behavior.
// These are appended to the global Builtins slice during package initialization.
var mapMethods = []*Builtin{
	{Name: "keys_map", Callback: mapKeys},                        // Returns an array of all keys in a map
	{Name: "insert_map", Callback: mapInsert},                    // Inserts or updates a key-value pair in a map
	{Name: "remove_map", Callback: mapRemove},                    // Removes a key-value pair from a map
	{Name: "contain_map", Callback: mapContain},                  // Checks if a map contains a key
	{Name: "enumerate_map", Callback: mapEnumerate},              // Returns array of [key, value] pairs
	{Name: "json_string_decode_map", Callback: jsonStringDecode}, // Decodes a JSON string into a map
	{Name: "json_string_encode_map", Callback: jsonStringEncode}, // Encodes a GoMix object to JSON string
}

// init is a special Go function that runs when the package is initialized.
// It registers the map methods as global builtins by appending them to the Builtins slice.
func init() {
	Builtins = append(Builtins, mapMethods...)
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

// jsonStringDecode parses a JSON string into a GoMix Map.
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

// jsonStringEncode converts a GoMix object into a JSON string.
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
// into GoMix internal objects.
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
