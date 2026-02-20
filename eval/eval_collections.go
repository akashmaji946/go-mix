/*
File    : go-mix/eval/eval_collections.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

// evalArrayExpression evaluates array literal expressions to create array objects.
//
// This method processes array literals (e.g., [1, 2, 3]) by:
// 1. Evaluating each element expression in order
// 2. Collecting the results into a slice
// 3. Creating an Array object containing all evaluated elements
//
// Arrays in Go-Mix are heterogeneous - they can contain elements of different types.
// If any element evaluation produces an error, the error is returned immediately
// and array creation is aborted.
//
// Parameters:
//   - n: An ArrayExpressionNode containing the element expressions
//
// Returns:
//   - objects.GoMixObject: An Array object containing the evaluated elements,
//     or an Error if any element evaluation failed
//
// Example:
//
//	[1, 2, 3]              // Array of integers
//	["a", "b", "c"]        // Array of strings
//	[1, "two", 3.0, true]  // Mixed-type array
//	[x + 1, y * 2]         // Array with computed elements
func (e *Evaluator) evalArrayExpression(n *parser.ArrayExpressionNode) std.GoMixObject {
	elements := make([]std.GoMixObject, len(n.Elements))
	for i, elem := range n.Elements {
		evaluated := e.Eval(elem)
		if IsError(evaluated) {
			return evaluated
		}
		elements[i] = evaluated
	}
	return &std.Array{Elements: elements}
}

// evalMapExpression evaluates map literal expressions to create map objects.
//
// This method processes map literals (e.g., map{10: 20, "key": "value"}) by:
// 1. Evaluating each key expression in order
// 2. Evaluating each corresponding value expression
// 3. Converting keys to strings for storage (Go maps require hashable keys)
// 4. Creating a Map object with the key-value pairs
//
// Maps in Go-Mix:
// - Keys are converted to strings using ToString() for consistent hashing
// - Values can be of any type
// - Duplicate keys: Later values overwrite earlier ones
// - Empty maps are supported: map{}
//
// Parameters:
//   - n: A MapExpressionNode containing parallel slices of key and value expressions
//
// Returns:
//   - objects.GoMixObject: A Map object containing the evaluated key-value pairs,
//     or an Error if any key or value evaluation failed
//
// Example:
//
//	map{10: 20, 30: 40}                    // Integer keys
//	map{"name": "John", "age": 25}         // String keys
//	map{1: "one", 2: "two", 3: "three"}    // Mixed content
//	map{x: y, a+b: c*d}                    // Computed keys and values
func (e *Evaluator) evalMapExpression(n *parser.MapExpressionNode) std.GoMixObject {
	pairs := make(map[string]std.GoMixObject)
	keys := make([]string, 0, len(n.Keys))

	for i := range n.Keys {
		// Evaluate key
		keyObj := e.Eval(n.Keys[i])
		if IsError(keyObj) {
			return keyObj
		}

		// Evaluate value
		valueObj := e.Eval(n.Values[i])
		if IsError(valueObj) {
			return valueObj
		}

		// Convert key to string for map storage
		keyStr := keyObj.ToString()

		// Check if key already exists
		if _, exists := pairs[keyStr]; !exists {
			keys = append(keys, keyStr)
		}

		// Store the key-value pair
		pairs[keyStr] = valueObj
	}

	return &std.Map{
		Pairs: pairs,
		Keys:  keys,
	}
}

// evalSetExpression evaluates set literal expressions to create set objects.
//
// This method processes set literals by:
// 1. Evaluating each element expression
// 2. Converting elements to strings for uniqueness checking
// 3. Automatically removing duplicates
// 4. Creating a Set object with unique values
//
// Sets in Go-Mix:
// - Elements are converted to strings using ToString() for uniqueness
// - Duplicates are automatically removed
// - Order of first occurrence is preserved
// - Empty sets are supported: set{}
//
// Parameters:
//   - n: A SetExpressionNode containing a slice of element expressions
//
// Returns:
//   - objects.GoMixObject: A Set object containing unique evaluated elements,
//     or an Error if any element evaluation failed
//
// Example:
//
//	set{1, 2, 3}                    // Integer elements
//	set{"apple", "banana"}          // String elements
//	set{1, 2, 2, 3}                 // Duplicates removed -> set{1, 2, 3}
//	set{x, y, x+y}                  // Computed elements
func (e *Evaluator) evalSetExpression(n *parser.SetExpressionNode) std.GoMixObject {
	elements := make(map[string]bool)
	values := make([]string, 0)

	for _, elemExpr := range n.Elements {
		// Evaluate element
		elemObj := e.Eval(elemExpr)
		if IsError(elemObj) {
			return elemObj
		}

		// Convert element to string for uniqueness
		elemStr := elemObj.ToString()

		// Add only if not already present (ensures uniqueness)
		if !elements[elemStr] {
			elements[elemStr] = true
			values = append(values, elemStr)
		}
	}

	return &std.Set{
		Elements: elements,
		Values:   values,
	}
}

// evalRangeExpression evaluates range expressions to create Range objects.
//
// This method processes range expressions (e.g., 2...5) by:
// 1. Evaluating the start expression
// 2. Evaluating the end expression
// 3. Validating both are integers
// 4. Creating a Range object with the start and end values
//
// Ranges are inclusive on both ends, meaning 2...5 includes 2, 3, 4, and 5.
//
// Parameters:
//   - n: A RangeExpressionNode containing the start and end expressions
//
// Returns:
//   - objects.GoMixObject: A Range object, or an Error if:
//   - Start expression evaluation fails
//   - End expression evaluation fails
//   - Either operand is not an integer
//
// Example:
//
//	2...5        // Returns Range{Start: 2, End: 5}
//	var x = 1...10  // Creates a range from 1 to 10
func (e *Evaluator) evalRangeExpression(n *parser.RangeExpressionNode) std.GoMixObject {
	// Evaluate start expression
	start := e.Eval(n.Start)
	if IsError(start) {
		return start
	}

	// Evaluate end expression
	end := e.Eval(n.End)
	if IsError(end) {
		return end
	}

	// Validate both are integers
	if start.GetType() != std.IntegerType {
		return e.CreateError("ERROR: range start must be an integer, got '%s'", start.GetType())
	}
	if end.GetType() != std.IntegerType {
		return e.CreateError("ERROR: range end must be an integer, got '%s'", end.GetType())
	}

	// Create and return the Range object
	startVal := start.(*std.Integer).Value
	endVal := end.(*std.Integer).Value

	return &std.Range{
		Start: startVal,
		End:   endVal,
	}
}
