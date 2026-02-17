/*
File    : go-mix/std/json.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

package std

import (
	"encoding/json"
	"io"
)

var jsonMethods = []*Builtin{
	{Name: "parse_json", Callback: jsonParse},
	{Name: "stringify_json", Callback: jsonStringify},
}

func init() {
	jsonPackage := &Package{
		Name:      "json",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range jsonMethods {
		jsonPackage.Functions[method.Name] = method
	}
	RegisterPackage(jsonPackage)
}

// jsonParse parses a JSON string into a Go-Mix Map.
func jsonParse(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: json.parse expects 1 argument (string)")
	}

	if args[0].GetType() != StringType {
		return createError("ERROR: argument to `json.parse` must be a string, got '%s'", args[0].GetType())
	}

	var data interface{}
	err := json.Unmarshal([]byte(args[0].ToString()), &data)
	if err != nil {
		return createError("ERROR: failed to decode JSON: %v", err)
	}

	return convertToGoMix(data)
}

// jsonStringify converts a Go-Mix object into a JSON string.
func jsonStringify(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: json.stringify expects 1 argument")
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
		if v == float64(int64(v)) {
			return &Integer{Value: int64(v)}
		}
		return &Float{Value: v}
	default:
		return &Nil{}
	}
}
