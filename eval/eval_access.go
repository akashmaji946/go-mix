/*
File    : go-mix/eval/eval_access.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/function"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

// evalIndexExpression evaluates array, map, list, and tuple element access using bracket notation.
//
// This method implements indexing for arrays, maps, lists, and tuples:
//
// Array/List/Tuple indexing:
// 1. Validates that the index is an integer
// 2. Supports negative indices (Python-style):
//   - Negative indices count from the end: -1 is last element, -2 is second-to-last, etc.
//
// 3. Performs bounds checking to prevent out-of-range access
//
// Map indexing:
// 1. Converts the index to a string key using ToString()
// 2. Looks up the value in the map
// 3. Returns nil if the key doesn't exist
//
// Parameters:
//   - n: An IndexExpressionNode containing the array/map/list/tuple and index expressions
//
// Returns:
//   - objects.GoMixObject: The element at the specified index/key, or an Error if:
//   - Left operand is not an array, map, list, or tuple
//   - Index is not an integer (for arrays/lists/tuples)
//   - Index is out of bounds
//
// Example:
//
//	var arr = [10, 20, 30];
//	arr[0]    // Returns 10 (first element)
//	arr[-1]   // Returns 30 (last element)
//
//	var l = list(1, 2, 3);
//	l[0]      // Returns 1
//	l[-1]     // Returns 3
//
//	var t = tuple("a", "b", "c");
//	t[1]      // Returns "b"
//
//	var m = map{"name": "John", "age": 25};
//	m["name"]  // Returns "John"
//	m["city"]  // Returns nil (key doesn't exist)
func (e *Evaluator) evalIndexExpression(n *parser.IndexExpressionNode) std.GoMixObject {
	left := e.Eval(n.Left)
	if IsError(left) {
		return left
	}

	index := e.Eval(n.Index)
	if IsError(index) {
		return index
	}

	// Handle map indexing
	if left.GetType() == std.MapType {
		mapObj := left.(*std.Map)
		keyStr := index.ToString()

		if value, exists := mapObj.Pairs[keyStr]; exists {
			return value
		}
		// Return nil if key doesn't exist
		return &std.Nil{}
	}

	// Handle range indexing
	if left.GetType() == std.RangeType {
		return e.evalRangeIndexExpression(left, index)
	}

	// Handle array, list, and tuple indexing
	leftType := left.GetType()
	if leftType != std.ArrayType && leftType != std.ListType && leftType != std.TupleType {
		return e.CreateError("ERROR: index operator not supported for type '%s'", leftType)
	}

	// Check if index is an integer
	if index.GetType() != std.IntegerType {
		return e.CreateError("ERROR: index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*std.Integer).Value
	var length int64
	var elements []std.GoMixObject

	// Get elements based on type
	switch leftType {
	case std.ArrayType:
		arr := left.(*std.Array)
		elements = arr.Elements
		length = int64(len(arr.Elements))
	case std.ListType:
		list := left.(*std.List)
		elements = list.Elements
		length = int64(len(list.Elements))
	case std.TupleType:
		tuple := left.(*std.Tuple)
		elements = tuple.Elements
		length = int64(len(tuple.Elements))
	}

	// Handle negative indices (Python-style)
	if idx < 0 {
		idx = length + idx
	}

	// Bounds checking
	if idx < 0 || idx >= length {
		return e.CreateError("ERROR: index out of bounds: index %d, length %d", idx, length)
	}

	return elements[idx]
}

// evalRangeIndexExpression evaluates index access on range objects.
//
// This method calculates the value at a specific index within a range sequence without generating
// the entire sequence. It supports both ascending and descending ranges and negative indexing.
//
// Parameters:
//   - left: The Range object
//   - index: The index object (must be Integer)
//
// Returns:
//   - objects.GoMixObject: The integer value at the specified index, or an Error if invalid
//
// Example:
//
//	range(1, 5)[0]  // Returns 1
//	range(1, 5)[-1] // Returns 5
func (e *Evaluator) evalRangeIndexExpression(left, index std.GoMixObject) std.GoMixObject {
	r := left.(*std.Range)

	// Check if index is an integer
	if index.GetType() != std.IntegerType {
		return e.CreateError("ERROR: range index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*std.Integer).Value
	start := r.Start
	end := r.End

	// Calculate the size and direction of the range
	var size int64
	if start <= end {
		size = end - start + 1
	} else {
		size = start - end + 1
	}

	// Handle negative indices (Python-style)
	if idx < 0 {
		idx = size + idx
	}

	// Bounds checking
	if idx < 0 || idx >= size {
		return e.CreateError("ERROR: range index out of bounds: index %d, size %d", idx, size)
	}

	// Calculate the value at the index
	var value int64
	if start <= end {
		// Ascending range
		value = start + idx
	} else {
		// Descending range
		value = start - idx
	}

	return &std.Integer{Value: value}
}

// evalSliceExpression evaluates array, list, and tuple slicing operations to extract sub-sequences.
//
// This method implements Python-style slicing with the syntax arr[start:end]:
// 1. Evaluates the array/list/tuple expression
// 2. Determines the start index (defaults to 0 if omitted)
// 3. Determines the end index (defaults to length if omitted)
// 4. Handles negative indices for both start and end
// 5. Clamps indices to valid range [0, length]
// 6. Creates a new array containing elements from start (inclusive) to end (exclusive)
//
// Index handling:
// - Omitted start: Defaults to 0 (beginning)
// - Omitted end: Defaults to length (end)
// - Negative indices: Count from end (-1 is last element position)
// - Out-of-range indices: Clamped to valid range (no error)
// - If start > end after processing: Returns empty array
//
// Note: Slicing always returns an array, even for lists and tuples (as per requirements).
//
// Parameters:
//   - n: A SliceExpressionNode containing the array/list/tuple, optional start, and optional end expressions
//
// Returns:
//   - objects.GoMixObject: A new Array containing the sliced elements, or an Error if:
//   - Left operand is not an array, list, or tuple
//   - Start or end index is not an integer
//
// Example:
//
//	var arr = [10, 20, 30, 40, 50];
//	arr[1:3]    // Returns [20, 30]
//	arr[:2]     // Returns [10, 20]
//	arr[2:]     // Returns [30, 40, 50]
//
//	var l = list(1, 2, 3, 4, 5);
//	l[1:3]      // Returns [2, 3] (array, not list)
//
//	var t = tuple("a", "b", "c", "d");
//	t[1:-1]     // Returns ["b", "c"] (array, not tuple)
func (e *Evaluator) evalSliceExpression(n *parser.SliceExpressionNode) std.GoMixObject {
	left := e.Eval(n.Left)
	if IsError(left) {
		return left
	}

	// Check if left is an array, list, or tuple
	leftType := left.GetType()
	if leftType != std.ArrayType && leftType != std.ListType && leftType != std.TupleType {
		return e.CreateError("ERROR: slice operator not supported for type '%s'", leftType)
	}

	var elements []std.GoMixObject
	var length int64

	// Get elements based on type
	switch leftType {
	case std.ArrayType:
		arr := left.(*std.Array)
		elements = arr.Elements
		length = int64(len(arr.Elements))
	case std.ListType:
		list := left.(*std.List)
		elements = list.Elements
		length = int64(len(list.Elements))
	case std.TupleType:
		tuple := left.(*std.Tuple)
		elements = tuple.Elements
		length = int64(len(tuple.Elements))
	}

	// Determine start index
	var start int64 = 0
	if n.Start != nil {
		startObj := e.Eval(n.Start)
		if IsError(startObj) {
			return startObj
		}
		if startObj.GetType() != std.IntegerType {
			return e.CreateError("ERROR: slice start index must be an integer, got '%s'", startObj.GetType())
		}
		start = startObj.(*std.Integer).Value
		// Handle negative start index
		if start < 0 {
			start = length + start
		}
		// Clamp to valid range
		if start < 0 {
			start = 0
		}
		if start > length {
			start = length
		}
	}

	// Determine end index
	var end int64 = length
	if n.End != nil {
		endObj := e.Eval(n.End)
		if IsError(endObj) {
			return endObj
		}
		if endObj.GetType() != std.IntegerType {
			return e.CreateError("ERROR: slice end index must be an integer, got '%s'", endObj.GetType())
		}
		end = endObj.(*std.Integer).Value
		// Handle negative end index
		if end < 0 {
			end = length + end
		}
		// Clamp to valid range
		if end < 0 {
			end = 0
		}
		if end > length {
			end = length
		}
	}

	// Ensure start <= end
	if start > end {
		start = end
	}

	// Create the sliced array (always returns array, even for lists/tuples)
	slicedElements := make([]std.GoMixObject, end-start)
	copy(slicedElements, elements[start:end])

	return &std.Array{Elements: slicedElements}
}

// getIndexValue retrieves a value from a container (array, list, or map) at a given index.
//
// This helper method abstracts index access for compound assignment operations.
// It handles type checking, index validation, and value retrieval.
//
// Parameters:
//   - container: The collection object (Array, List, or Map)
//   - index: The index or key to access
//
// Returns:
//   - objects.GoMixObject: The value at the index, or an Error if invalid
func (e *Evaluator) getIndexValue(container, index std.GoMixObject) std.GoMixObject {
	if container.GetType() == std.MapType {
		mapObj := container.(*std.Map)
		keyStr := index.ToString()
		if value, exists := mapObj.Pairs[keyStr]; exists {
			return value
		}
		return &std.Nil{}
	}

	leftType := container.GetType()
	if leftType != std.ArrayType && leftType != std.ListType {
		return e.CreateError("ERROR: index operator not supported for type '%s'", leftType)
	}

	if index.GetType() != std.IntegerType {
		return e.CreateError("ERROR: index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*std.Integer).Value
	var length int64
	var elements []std.GoMixObject

	if leftType == std.ArrayType {
		arr := container.(*std.Array)
		elements = arr.Elements
		length = int64(len(arr.Elements))
	} else {
		list := container.(*std.List)
		elements = list.Elements
		length = int64(len(list.Elements))
	}

	if idx < 0 {
		idx = length + idx
	}

	if idx < 0 || idx >= length {
		return e.CreateError("ERROR: index out of bounds: index %d, length %d", idx, length)
	}

	return elements[idx]
}

// evalMemberAccess evaluates member access (dot operator) on a struct instance.
//
// This method handles accessing fields or calling methods on an object instance.
// It distinguishes between:
// - Method calls: Dispatches to callFunctionOnObject
// - Field access: Looks up instance fields, then static fields
//
// Parameters:
//   - structInstance: The object instance being accessed
//   - node: The expression to the right of the dot (Identifier or CallExpression)
//
// Returns:
//   - objects.GoMixObject: The field value or method return value
func (e *Evaluator) evalMemberAccess(structInstance *std.GoMixObjectInstance, node parser.ExpressionNode) std.GoMixObject {
	// Handle Method Call
	if fn, ok := node.(*parser.CallExpressionNode); ok {
		methodName := fn.FunctionIdentifier.Name
		methodInterface, exists := structInstance.Struct.Methods[methodName]
		if !exists {
			return e.CreateError("ERROR: method (%s) does not exist in struct (%s)", methodName, structInstance.Struct.GetName())
		}
		method, ok := methodInterface.(*function.Function)
		if !ok {
			return e.CreateError("ERROR: method (%s) not found in struct (%s)", methodName, structInstance.Struct.GetName())
		}
		params := make([]NamedParameter, len(fn.Arguments))
		if len(fn.Arguments) != len(method.Params) {
			return e.CreateError("ERROR: wrong number of arguments for method (%s): expected %d, got %d", methodName, len(method.Params), len(fn.Arguments))
		}
		for i, arg := range fn.Arguments {
			params[i] = NamedParameter{
				Name:  method.Params[i].Name,
				Value: e.Eval(arg),
			}
			if IsError(params[i].Value) {
				return params[i].Value
			}
		}

		res := e.callFunctionOnObject(methodName, structInstance, params...)
		if res.GetType() == std.ErrorType {
			return res
		}
		return res
	}

	// Handle Field Access
	if ident, ok := node.(*parser.IdentifierExpressionNode); ok {
		fieldName := ident.Name
		if val, ok := structInstance.InstanceFields[fieldName]; ok {
			return val
		}
		if val, ok := structInstance.Struct.ClassFields[fieldName]; ok {
			return val
		}
		if val, ok := structInstance.Struct.ClassFields[fieldName]; ok {
			return val
		}
		return e.CreateError("ERROR: field (%s) not found in struct instance", fieldName)
	}

	return e.CreateError("ERROR: member access operator (.) must be followed by a function call or identifier")
}

// evalStructMemberAccess evaluates member access on a struct type (static access).
//
// This method handles accessing static fields on the struct type itself.
//
// Parameters:
//   - s: The struct type definition
//   - node: The identifier expression for the field
//
// Returns:
//   - objects.GoMixObject: The static field value
func (e *Evaluator) evalStructMemberAccess(s *std.GoMixStruct, node parser.ExpressionNode) std.GoMixObject {
	// Handle Field Access
	if ident, ok := node.(*parser.IdentifierExpressionNode); ok {
		fieldName := ident.Name
		if val, ok := s.ClassFields[fieldName]; ok {
			return val
		}
		return e.CreateError("ERROR: class field (%s) not found in struct (%s)", fieldName, s.Name)
	}
	return e.CreateError("ERROR: invalid member access on struct")
}

// evalPackageMemberAccess evaluates member access on a package (e.g., math.abs).
//
// This method handles accessing functions from an imported package. Note that this
// is primarily used for direct access (e.g., getting a reference to the function).
// Actual function calls are handled through evalCallExpression which has special
// logic for package.function() calls.
//
// Parameters:
//   - pkg: The package object
//   - node: The identifier expression for the function name
//
// Returns:
//   - objects.GoMixObject: An error if the function is not found, or nil otherwise
//     (The actual return value is handled through CallExpression evaluation)
func (e *Evaluator) evalPackageMemberAccess(pkg *std.Package, node parser.ExpressionNode) std.GoMixObject {
	// Handle Function Access
	if ident, ok := node.(*parser.IdentifierExpressionNode); ok {
		funcName := ident.Name
		if _, ok := pkg.Functions[funcName]; ok {
			// Function exists - the actual call will be handled in evalCallExpression
			// We return nil here since this path is only hit for non-call accesses
			return &std.Nil{}
		}
		return e.CreateError("ERROR: function '%s' not found in package '%s'", funcName, pkg.Name)
	}
	return e.CreateError("ERROR: invalid member access on package")
}

// evalEnumAccessExpression evaluates enum member access expressions.
//
// This method handles accessing enum members like Color.RED or Status.ACTIVE
// by looking up the enum type and retrieving the member's value.
//
// Parameters:
//   - n: An EnumAccessExpressionNode containing the enum name and member name
//
// Returns:
//   - objects.GoMixObject: The enum member's value, or an Error if:
//   - The enum type is not found
//   - The member is not found in the enum
//
// Example:
//
//	Color.RED      // Returns the integer value of RED
//	Status.ACTIVE  // Returns the integer value of ACTIVE
func (e *Evaluator) evalEnumAccessExpression(n *parser.EnumAccessExpressionNode) std.GoMixObject {
	// Look up the enum type
	enumObj, ok := e.Scp.LookUp(n.EnumName.Name)
	if !ok {
		return e.CreateError("ERROR: enum type '%s' not found", n.EnumName.Name)
	}

	// Check if it's an enum type
	enumType, ok := enumObj.(*std.GoMixEnum)
	if !ok {
		return e.CreateError("ERROR: '%s' is not an enum type", n.EnumName.Name)
	}

	// Look up the member
	memberValue, ok := enumType.Members[n.MemberName.Name]
	if !ok {
		return e.CreateError("ERROR: enum member '%s' not found in enum '%s'", n.MemberName.Name, n.EnumName.Name)
	}

	return memberValue
}
