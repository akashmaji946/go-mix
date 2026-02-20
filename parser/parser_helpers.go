package parser

import "github.com/akashmaji946/go-mix/std"

// toFloat64 converts a GoMixObject to float64.
// This helper function is used for mixed-type arithmetic operations.
//
// Parameters:
//
//	obj - The object to convert (Integer or Float)
//
// Returns:
//
//	The float64 representation of the value, or 0 if not a number
func toFloat64(obj std.GoMixObject) float64 {
	if obj.GetType() == std.IntegerType {
		return float64(obj.(*std.Integer).Value)
	} else if obj.GetType() == std.FloatType {
		return obj.(*std.Float).Value
	}
	return 0
}
