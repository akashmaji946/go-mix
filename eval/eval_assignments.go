/*
File    : go-mix/eval/eval_assignments.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
)

// evalAssignmentExpression evaluates variable assignment expressions with comprehensive validation.
//
// This method handles the assignment operator (=) and performs several critical checks:
// 1. Evaluates the right-hand side expression to get the value to assign
// 2. Verifies the variable exists in the current scope or any parent scope
// 3. Prevents assignment to constants (declared with 'const')
// 4. Enforces type safety for 'let' variables (must match declared type)
// 5. Updates the variable in its defining scope (essential for closures)
//
// The method uses Scope.Assign() rather than Bind() to ensure the variable is updated
// in the scope where it was originally defined, not the current scope. This is crucial
// for closures to work correctly.
//
// Parameters:
//   - n: An AssignmentExpressionNode containing the target identifier and value expression
//
// Returns:
//   - objects.GoMixObject: The assigned value on success, or an Error object if:
//   - The variable doesn't exist
//   - Attempting to assign to a constant
//   - Type mismatch for 'let' variables
//
// Example:
//
//		let x = 10;  // Declares x as integer
//		x = 20;      // Valid: same type
//		x = "hi";    // Error: type mismatch
//		const y = 5;
//		y = 10;      // Error: can't assign to constant
//	 	var z = 15;
//	 	z += 5;     // Valid compound assignment
func (e *Evaluator) evalAssignmentExpression(n *parser.AssignmentExpressionNode) std.GoMixObject {

	if n.Operation.Type != lexer.ASSIGN_OP {
		return e.evalCompoundAssignment(n)
	}

	// Evaluate the right-hand side
	rightVal := e.Eval(n.Right)
	if IsError(rightVal) {
		return rightVal
	}

	// Check if left is an identifier or index expression
	if identNode, ok := n.Left.(*parser.IdentifierExpressionNode); ok {
		// Handle identifier assignment
		return e.evalIdentifierAssignment(identNode, rightVal)
	}

	if indexNode, ok := n.Left.(*parser.IndexExpressionNode); ok {
		// Handle index assignment (e.g., a[0] = 11, map["key"] = value)
		return e.evalIndexAssignment(indexNode, rightVal)
	}

	// Handle member assignment (obj.field = val)
	if binNode, ok := n.Left.(*parser.BinaryExpressionNode); ok {
		if binNode.Operation.Type == lexer.DOT_OP {
			return e.evalMemberAssignment(binNode, rightVal)
		}
	}

	// Should not reach here if parser is correct
	return e.CreateError("ERROR: invalid assignment target")
}

// evalCompoundAssignment handles compound assignment operators (+=, -=, *=, /=, etc.) with type validation.
//
// This method evaluates the right-hand side expression, retrieves the current value of the left-hand side,
// performs the appropriate binary operation based on the compound operator, and then assigns the result back
// to the left-hand side. It supports identifiers, index expressions, and member access as assignment targets.
// The method also includes comprehensive error handling for unsupported operators, type mismatches, and invalid assignment targets.
//
// Parameters:
//   - n: The AssignmentExpressionNode representing the compound assignment
//
// Returns:
//   - objects.GoMixObject: The result of the assignment (the new value), or an Error object
func (e *Evaluator) evalCompoundAssignment(n *parser.AssignmentExpressionNode) std.GoMixObject {
	var binOpType lexer.TokenType
	switch n.Operation.Type {
	case lexer.PLUS_ASSIGN:
		binOpType = lexer.PLUS_OP
	case lexer.MINUS_ASSIGN:
		binOpType = lexer.MINUS_OP
	case lexer.MUL_ASSIGN:
		binOpType = lexer.MUL_OP
	case lexer.DIV_ASSIGN:
		binOpType = lexer.DIV_OP
	case lexer.MOD_ASSIGN:
		binOpType = lexer.MOD_OP
	case lexer.BIT_AND_ASSIGN:
		binOpType = lexer.BIT_AND_OP
	case lexer.BIT_OR_ASSIGN:
		binOpType = lexer.BIT_OR_OP
	case lexer.BIT_XOR_ASSIGN:
		binOpType = lexer.BIT_XOR_OP
	case lexer.BIT_LEFT_ASSIGN:
		binOpType = lexer.BIT_LEFT_OP
	case lexer.BIT_RIGHT_ASSIGN:
		binOpType = lexer.BIT_RIGHT_OP
	default:
		return e.createError(n.Operation, "ERROR: unknown compound assignment operator: %s", n.Operation.Literal)
	}

	rightVal := e.Eval(n.Right)
	if IsError(rightVal) {
		return rightVal
	}

	// 1. Identifier
	if identNode, ok := n.Left.(*parser.IdentifierExpressionNode); ok {
		leftVal := e.evalIdentifierExpression(identNode)
		if IsError(leftVal) {
			return leftVal
		}
		newVal := e.evaluateBinaryOp(n.Operation, binOpType, leftVal, rightVal)
		if IsError(newVal) {
			return newVal
		}
		return e.evalIdentifierAssignment(identNode, newVal)
	}

	// 2. Index Expression
	if indexNode, ok := n.Left.(*parser.IndexExpressionNode); ok {
		container := e.Eval(indexNode.Left)
		if IsError(container) {
			return container
		}
		index := e.Eval(indexNode.Index)
		if IsError(index) {
			return index
		}

		leftVal := e.getIndexValue(container, index)
		if IsError(leftVal) {
			return leftVal
		}

		newVal := e.evaluateBinaryOp(n.Operation, binOpType, leftVal, rightVal)
		if IsError(newVal) {
			return newVal
		}

		switch container.GetType() {
		case std.ArrayType:
			return e.evalArrayIndexAssignment(container, index, newVal)
		case std.ListType:
			return e.evalListIndexAssignment(container, index, newVal)
		case std.MapType:
			return e.evalMapIndexAssignment(container, index, newVal)
		default:
			return e.CreateError("ERROR: index assignment not supported for type '%s'", container.GetType())
		}
	}

	// 3. Member Access
	if binNode, ok := n.Left.(*parser.BinaryExpressionNode); ok {
		if binNode.Operation.Type == lexer.DOT_OP {
			leftObj := e.Eval(binNode.Left)
			if IsError(leftObj) {
				return leftObj
			}
			if leftObj.GetType() == std.StructType {
				s := leftObj.(*std.GoMixStruct)
				ident, ok := binNode.Right.(*parser.IdentifierExpressionNode)
				if !ok {
					return e.CreateError("ERROR: invalid member assignment target")
				}
				if s.ConstFields[ident.Name] {
					return e.CreateError("ERROR: can't assign to constant field (%s) in struct (%s)", ident.Name, s.Name)
				}
				leftVal, ok := s.ClassFields[ident.Name]
				if !ok {
					return e.CreateError("ERROR: class field (%s) not found", ident.Name)
				}
				newVal := e.evaluateBinaryOp(n.Operation, binOpType, leftVal, rightVal)
				if IsError(newVal) {
					return newVal
				}
				if s.LetFields[ident.Name] {
					expectedType := s.LetTypes[ident.Name]
					if newVal.GetType() != expectedType {
						return e.CreateError("ERROR: can't assign `%s` to field (%s) of type `%s` in struct (%s)", newVal.GetType(), ident.Name, expectedType, s.Name)
					}
				}
				s.ClassFields[ident.Name] = newVal
				return newVal
			}

			if leftObj.GetType() != std.ObjectType {
				return e.CreateError("ERROR: member access can only be done on struct instances")
			}
			inst := leftObj.(*std.GoMixObjectInstance)
			ident, ok := binNode.Right.(*parser.IdentifierExpressionNode)
			if !ok {
				return e.CreateError("ERROR: invalid member assignment target")
			}
			leftVal, ok := inst.InstanceFields[ident.Name]
			if !ok {
				return e.CreateError("ERROR: field (%s) not found in struct instance", ident.Name)
			}
			newVal := e.evaluateBinaryOp(n.Operation, binOpType, leftVal, rightVal)
			if IsError(newVal) {
				return newVal
			}
			inst.InstanceFields[ident.Name] = newVal
			return newVal
		}
	}

	return e.CreateError("ERROR: invalid assignment target")
}

// evalIdentifierAssignment handles assignment to an identifier (variable).
//
// This method performs the necessary checks to ensure that the variable exists,
// is not a constant, and if it's a 'let' variable, that the assigned value matches the declared type.
// It then updates the variable in its defining scope using Scope.Assign(),
// which is crucial for closures to work correctly.
//
// Parameters:
//   - ident: The identifier node representing the variable to assign to
//   - val: The value to assign
//
// Returns:
//   - objects.GoMixObject: The assigned value on success, or an Error if validation fails
func (e *Evaluator) evalIdentifierAssignment(ident *parser.IdentifierExpressionNode, val std.GoMixObject) std.GoMixObject {
	// Check if the variable exists in the current scope or any parent scope
	_, exists := e.Scp.LookUp(ident.Name)
	if !exists {
		return e.createError(ident.Token, "ERROR: identifier not found: (%s)", ident.Name)
	}

	// Check if it's a constant using the new IsConstant method
	if e.Scp.IsConstant(ident.Name) {
		return e.createError(ident.Token, "ERROR: can't assign to constant (%s)", ident.Name)
	}

	// Check if it's a let variable and if the type is compatible
	if e.Scp.IsLetVariable(ident.Name) {
		expectedType, ok := e.Scp.GetLetType(ident.Name)
		if !ok {
			return e.createError(ident.Token, "ERROR: let variable type not found: (%s)", ident.Name)
		}
		if val.GetType() != expectedType {
			return e.createError(ident.Token, "ERROR: can't assign `%s` to variable (%s) of type `%s`", val.GetType(), ident.Name, expectedType)
		}
	}

	// Use Assign to update the variable in the scope where it was defined
	// This is essential for closures to work correctly
	e.Scp.Assign(ident.Name, val)

	return val
}

// evalIndexAssignment handles assignment to an indexed element (e.g., a[0] = 11, map["key"] = value).
//
// This method evaluates the container and index expressions, retrieves the current value at that index
// (if needed for validation, though mostly handled by specific implementations), performs the assignment,
// and updates the container accordingly. It delegates to specific handlers for arrays, lists, and maps.
//
// Parameters:
//   - indexNode: The IndexExpressionNode representing the target (container[index])
//   - val: The value to assign
//
// Returns:
//   - objects.GoMixObject: The assigned value on success, or an Error if the operation fails
func (e *Evaluator) evalIndexAssignment(indexNode *parser.IndexExpressionNode, val std.GoMixObject) std.GoMixObject {
	// Evaluate the container (array, list, map, etc.)
	container := e.Eval(indexNode.Left)
	if IsError(container) {
		return container
	}

	// Evaluate the index/key
	index := e.Eval(indexNode.Index)
	if IsError(index) {
		return index
	}

	// Handle different container types
	switch container.GetType() {
	case std.ArrayType:
		return e.evalArrayIndexAssignment(container, index, val)
	case std.ListType:
		return e.evalListIndexAssignment(container, index, val)
	case std.MapType:
		return e.evalMapIndexAssignment(container, index, val)
	default:
		return e.CreateError("ERROR: index assignment not supported for type '%s'", container.GetType())
	}
}

// evalArrayIndexAssignment handles assignment to an array element.
//
// This method checks that the index is an integer, supports negative indexing (Python-style),
// performs bounds checking, and assigns the value to the specified index in the array.
//
// Parameters:
//   - container: The Array object
//   - index: The index object (must be Integer)
//   - val: The value to assign
//
// Returns:
//   - objects.GoMixObject: The assigned value on success, or an Error if index is invalid/out of bounds
func (e *Evaluator) evalArrayIndexAssignment(container, index, val std.GoMixObject) std.GoMixObject {
	arr := container.(*std.Array)

	// Check if index is an integer
	if index.GetType() != std.IntegerType {
		return e.CreateError("ERROR: array index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*std.Integer).Value
	length := int64(len(arr.Elements))

	// Handle negative indices (Python-style)
	if idx < 0 {
		idx = length + idx
	}

	// Bounds checking
	if idx < 0 || idx >= length {
		return e.CreateError("ERROR: index out of bounds: index %d, length %d", idx, length)
	}

	// Assign the value
	arr.Elements[idx] = val
	return val
}

// evalListIndexAssignment handles assignment to a list element.
//
// This method checks that the index is an integer, supports negative indexing (Python-style),
// performs bounds checking, and assigns the value to the specified index in the list.
//
// Parameters:
//   - container: The List object
//   - index: The index object (must be Integer)
//   - val: The value to assign
//
// Returns:
//   - objects.GoMixObject: The assigned value on success, or an Error if index is invalid/out of bounds
func (e *Evaluator) evalListIndexAssignment(container, index, val std.GoMixObject) std.GoMixObject {
	list := container.(*std.List)

	// Check if index is an integer
	if index.GetType() != std.IntegerType {
		return e.CreateError("ERROR: list index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*std.Integer).Value
	length := int64(len(list.Elements))

	// Handle negative indices (Python-style)
	if idx < 0 {
		idx = length + idx
	}

	// Bounds checking
	if idx < 0 || idx >= length {
		return e.CreateError("ERROR: index out of bounds: index %d, length %d", idx, length)
	}

	// Assign the value
	list.Elements[idx] = val
	return val
}

// evalMapIndexAssignment handles assignment to a map key.
//
// This method assigns a value to a specific key in a map. It converts the index
// object to a string key, checks if the key exists to maintain insertion order
// (if it's a new key), and then updates the map's pairs.
//
// Parameters:
//   - container: The Map object to update
//   - index: The key object (converted to string)
//   - val: The value to assign
//
// Returns:
//   - objects.GoMixObject: The assigned value
func (e *Evaluator) evalMapIndexAssignment(container, index, val std.GoMixObject) std.GoMixObject {
	m := container.(*std.Map)

	// Convert index to string key
	keyStr := index.ToString()

	// Check if key already exists
	_, exists := m.Pairs[keyStr]
	if !exists {
		// New key - add to keys list to maintain insertion order
		m.Keys = append(m.Keys, keyStr)
	}

	// Assign the value
	m.Pairs[keyStr] = val
	return val
}

// evalMemberAssignment handles assignment to a struct member (e.g., obj.field = val).
//
// This method evaluates the left side to get the struct instance or type, validates
// the target field, checks for const/let constraints, and performs the assignment.
// It supports assignment to both instance fields (on objects) and static fields (on struct types).
//
// Parameters:
//   - node: The BinaryExpressionNode representing the member access (DOT_OP)
//   - val: The value to assign
//
// Returns:
//   - objects.GoMixObject: The assigned value, or an Error if the target is invalid or immutable
func (e *Evaluator) evalMemberAssignment(node *parser.BinaryExpressionNode, val std.GoMixObject) std.GoMixObject {
	left := e.Eval(node.Left)
	if IsError(left) {
		return left
	}

	if left.GetType() == std.StructType {
		s := left.(*std.GoMixStruct)
		ident, ok := node.Right.(*parser.IdentifierExpressionNode)
		if !ok {
			return e.CreateError("ERROR: invalid member assignment target")
		}
		if s.ConstFields[ident.Name] {
			return e.CreateError("ERROR: can't assign to constant field (%s) in struct (%s)", ident.Name, s.Name)
		}
		if s.LetFields[ident.Name] {
			expectedType := s.LetTypes[ident.Name]
			if val.GetType() != expectedType {
				return e.CreateError("ERROR: can't assign `%s` to field (%s) of type `%s` in struct (%s)", val.GetType(), ident.Name, expectedType, s.Name)
			}
		}
		s.ClassFields[ident.Name] = val
		return val
	}

	if left.GetType() != std.ObjectType {
		return e.CreateError("ERROR: member assignment can only be done on struct instances, got %s", left.GetType())
	}

	inst := left.(*std.GoMixObjectInstance)

	ident, ok := node.Right.(*parser.IdentifierExpressionNode)
	if !ok {
		return e.CreateError("ERROR: invalid member assignment target")
	}

	inst.InstanceFields[ident.Name] = val
	return val
}
