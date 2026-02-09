/*
File    : go-mix/eval/evaluator_expressions.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"fmt"

	"github.com/akashmaji946/go-mix/function"
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
)

// Eval is the main evaluation dispatcher that converts AST nodes into runtime objects.
//
// This method serves as the central hub of the evaluation process, routing each node type
// to its appropriate evaluation handler. It implements a type switch pattern to handle:
// - Literal expressions: Return their corresponding object values directly
// - Unary/Binary expressions: Compute and return results
// - Boolean expressions: Evaluate comparisons and logical operations
// - Control flow: Handle if-else, loops, and return statements
// - Function operations: Handle declarations and calls
// - Variable operations: Handle declarations, lookups, and assignments
// - Array operations: Handle array literals, indexing, and slicing
//
// The evaluation process is recursive - complex expressions are broken down into
// simpler sub-expressions that are evaluated in turn.
//
// Parameters:
//   - n: The AST node to evaluate (can be any type implementing parser.Node)
//
// Returns:
//   - objects.GoMixObject: The result of evaluating the node. For statements, typically
//     returns Nil unless there's an error or return value. For expressions, returns the
//     computed value. Errors halt evaluation and are propagated up the call stack.
//
// Example flow:
//
//	RootNode -> evalStatements -> Eval(each statement) -> specific eval methods
func (e *Evaluator) Eval(n parser.Node) objects.GoMixObject {
	switch n := n.(type) {
	case *parser.RootNode:
		result := e.evalStatements(n.Statements)
		return UnwrapReturnValue(result)
	case *parser.BooleanLiteralExpressionNode:
		return n.Value
	case *parser.IntegerLiteralExpressionNode:
		return n.Value
	case *parser.StringLiteralExpressionNode:
		return n.Value
	case *parser.FloatLiteralExpressionNode:
		return n.Value
	case *parser.NilLiteralExpressionNode:
		return &objects.Nil{}
	case *parser.BinaryExpressionNode:
		return e.evalBinaryExpression(n)
	case *parser.UnaryExpressionNode:
		return e.evalUnaryExpression(n)
	case *parser.BooleanExpressionNode:
		return e.evalBooleanExpression(n)
	case *parser.ParenthesizedExpressionNode:
		return e.Eval(n.Expr)
	case *parser.IfExpressionNode:
		return e.evalConditionalExpression(n)
	case *parser.DeclarativeStatementNode:
		return e.evalDeclarativeStatement(n)
	case *parser.ReturnStatementNode:
		return e.evalReturnStatement(n)
	case *parser.BlockStatementNode:
		return e.evalBlockStatement(n)
	case *parser.IdentifierExpressionNode:
		return e.evalIdentifierExpression(n)
	case *parser.FunctionStatementNode:
		return e.RegisterFunction(n)
	case *parser.CallExpressionNode:
		return e.evalCallExpression(n)
	case *parser.AssignmentExpressionNode:
		return e.evalAssignmentExpression(n)
	case *parser.ForLoopStatementNode:
		return e.evalForLoop(n)
	case *parser.WhileLoopStatementNode:
		return e.evalWhileLoop(n)
	case *parser.ArrayExpressionNode:
		return e.evalArrayExpression(n)
	case *parser.MapExpressionNode:
		return e.evalMapExpression(n)
	case *parser.SetExpressionNode:
		return e.evalSetExpression(n)
	case *parser.IndexExpressionNode:
		return e.evalIndexExpression(n)
	case *parser.SliceExpressionNode:
		return e.evalSliceExpression(n)
	case *parser.RangeExpressionNode:
		return e.evalRangeExpression(n)
	case *parser.ForeachLoopStatementNode:
		return e.evalForeachLoop(n)
	case *parser.StructDeclarationNode:
		return e.evalStructDeclaration(n)
	case *parser.NewCallExpressionNode:
		return e.evalNewCallExpression(n)
	case *parser.BreakStatementNode:
		return &objects.Break{}
	case *parser.ContinueStatementNode:
		return &objects.Continue{}
	default:
		return &objects.Nil{}
	}
}

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
func (e *Evaluator) evalAssignmentExpression(n *parser.AssignmentExpressionNode) objects.GoMixObject {

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

func (e *Evaluator) evalCompoundAssignment(n *parser.AssignmentExpressionNode) objects.GoMixObject {
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
		case objects.ArrayType:
			return e.evalArrayIndexAssignment(container, index, newVal)
		case objects.ListType:
			return e.evalListIndexAssignment(container, index, newVal)
		case objects.MapType:
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
			if leftObj.GetType() != objects.ObjectType {
				return e.CreateError("ERROR: member access can only be done on struct instances")
			}
			inst := leftObj.(*objects.GoMixObjectInstance)
			ident, ok := binNode.Right.(*parser.IdentifierExpressionNode)
			if !ok {
				return e.CreateError("ERROR: invalid member assignment target")
			}
			leftVal, ok := inst.Fields[ident.Name]
			if !ok {
				return e.CreateError("ERROR: field (%s) not found in struct instance", ident.Name)
			}
			newVal := e.evaluateBinaryOp(n.Operation, binOpType, leftVal, rightVal)
			if IsError(newVal) {
				return newVal
			}
			inst.Fields[ident.Name] = newVal
			return newVal
		}
	}

	return e.CreateError("ERROR: invalid assignment target")
}

// evalIdentifierAssignment handles assignment to an identifier (variable).
func (e *Evaluator) evalIdentifierAssignment(ident *parser.IdentifierExpressionNode, val objects.GoMixObject) objects.GoMixObject {
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
func (e *Evaluator) evalIndexAssignment(indexNode *parser.IndexExpressionNode, val objects.GoMixObject) objects.GoMixObject {
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
	case objects.ArrayType:
		return e.evalArrayIndexAssignment(container, index, val)
	case objects.ListType:
		return e.evalListIndexAssignment(container, index, val)
	case objects.MapType:
		return e.evalMapIndexAssignment(container, index, val)
	default:
		return e.CreateError("ERROR: index assignment not supported for type '%s'", container.GetType())
	}
}

// evalArrayIndexAssignment handles assignment to an array element.
func (e *Evaluator) evalArrayIndexAssignment(container, index, val objects.GoMixObject) objects.GoMixObject {
	arr := container.(*objects.Array)

	// Check if index is an integer
	if index.GetType() != objects.IntegerType {
		return e.CreateError("ERROR: array index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*objects.Integer).Value
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
func (e *Evaluator) evalListIndexAssignment(container, index, val objects.GoMixObject) objects.GoMixObject {
	list := container.(*objects.List)

	// Check if index is an integer
	if index.GetType() != objects.IntegerType {
		return e.CreateError("ERROR: list index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*objects.Integer).Value
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
func (e *Evaluator) evalMapIndexAssignment(container, index, val objects.GoMixObject) objects.GoMixObject {
	m := container.(*objects.Map)

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
func (e *Evaluator) evalMemberAssignment(node *parser.BinaryExpressionNode, val objects.GoMixObject) objects.GoMixObject {
	left := e.Eval(node.Left)
	if IsError(left) {
		return left
	}

	if left.GetType() != objects.ObjectType {
		return e.CreateError("ERROR: member assignment can only be done on struct instances, got %s", left.GetType())
	}

	inst := left.(*objects.GoMixObjectInstance)

	ident, ok := node.Right.(*parser.IdentifierExpressionNode)
	if !ok {
		return e.CreateError("ERROR: invalid member assignment target")
	}

	inst.Fields[ident.Name] = val
	return val
}

// evalCallExpression evaluates function call expressions for both builtin and user-defined functions.
//
// This method handles the complete function call process:
// 1. Checks if the function name is a builtin (print, len, push, etc.)
//   - If builtin: evaluates arguments and invokes the builtin directly
//
// 2. For user-defined functions:
//   - Looks up the function in the scope chain
//   - Validates it's actually a function object
//   - Checks argument count matches parameter count
//   - Creates a new call-site scope with the function's captured scope as parent
//   - Binds arguments to parameters in the call-site scope
//   - Evaluates the function body in the new scope
//   - Unwraps return values and handles closure scope updates
//
// The scope handling is critical for closures: functions capture their defining scope,
// and when they return other functions, those returned functions get an updated scope
// that includes variables from the call site.
//
// Parameters:
//   - n: A CallExpressionNode containing the function identifier and argument expressions
//
// Returns:
//   - objects.GoMixObject: The function's return value, or an Error object if:
//   - Function not found
//   - Identifier is not a function
//   - Wrong number of arguments provided
//
// Example:
//
//	print("Hello");           // Builtin function call
//	add(5, 3);                // User-defined function call
//	makeCounter()(10);        // Closure returning a function
func (e *Evaluator) evalCallExpression(n *parser.CallExpressionNode) objects.GoMixObject {

	funcName := n.FunctionIdentifier.Name

	// Support method calls: obj.method()
	if dotIdx := indexOfDot(funcName); dotIdx > 0 {
		objName := funcName[:dotIdx]
		methodName := funcName[dotIdx+1:]
		// fmt.Printf("DEBUG: Method call detected: obj='%s', method='%s'\n", objName, methodName)
		objVal, ok := e.Scp.LookUp(objName)
		if !ok {
			return e.createError(n.FunctionIdentifier.Token, "ERROR: object not found: (%s)", objName)
		}
		inst, ok := objVal.(*objects.GoMixObjectInstance)
		if !ok {
			return e.CreateError("ERROR: (%s) is not a struct instance", objName)
		}
		// Prepare named parameters for method call
		params := make([]NamedParameter, len(n.Arguments))
		for i, arg := range n.Arguments {
			evaluated := e.Eval(arg)
			if IsError(evaluated) {
				return evaluated
			}
			params[i] = NamedParameter{Name: "", Value: evaluated}
		}
		result := e.callFunctionOnObject(methodName, inst, params...)
		// fmt.Printf("DEBUG: Method call result type=%s\n", result.GetType())
		return result
	}

	// look for builtin name
	if ok := e.IsBuiltin(funcName); ok {
		args := make([]objects.GoMixObject, len(n.Arguments))
		for i, arg := range n.Arguments {
			args[i] = e.Eval(arg)
			if IsError(args[i]) {
				return args[i]
			}
		}
		rv := e.InvokeBuiltin(funcName, args...)
		return rv
	}

	// lookup for function name
	obj, ok := e.Scp.LookUp(funcName)
	if !ok {
		return e.createError(n.FunctionIdentifier.Token, "ERROR: function not found: (%s)", funcName)
	}
	if obj.GetType() != objects.FunctionType {
		return e.createError(n.FunctionIdentifier.Token, "ERROR: not a function: (%s)", funcName)
	}
	functionObject := obj.(*function.Function)

	// Validate argument count
	expectedArgs := len(functionObject.Params)
	actualArgs := len(n.Arguments)
	if actualArgs != expectedArgs {
		return e.CreateError("ERROR: wrong number of arguments: expected %d, got %d", expectedArgs, actualArgs)
	}

	// Create a new scope with the function's captured scope as parent
	var parentScope *scope.Scope
	if functionObject.Scp != nil {
		parentScope = functionObject.Scp
	} else {
		parentScope = e.Scp
	}
	callSiteScope := scope.NewScope(parentScope)

	for i, param := range functionObject.Params {
		val := e.Eval(n.Arguments[i])
		if IsError(val) {
			return val
		}
		callSiteScope.Bind(param.Name, val)
	}
	oldScope := e.Scp
	e.Scp = callSiteScope
	result := e.Eval(functionObject.Body)
	e.Scp = oldScope

	// Unwrap return value if present
	if retVal, isReturn := result.(*objects.ReturnValue); isReturn {
		returnVal := retVal.Value
		// If returning a function, update its captured scope to the current scope
		// This is essential for closures to work correctly
		// Only copy if the call site scope has variables not in the function's existing scope
		if fn, isFunc := returnVal.(*function.Function); isFunc {
			if len(callSiteScope.Variables) > len(fn.Scp.Variables) {
				fn.Scp = callSiteScope.Copy()
			}
		}
		return returnVal
	}
	return result

}

// evalIdentifierExpression resolves an identifier to its value by searching the scope chain.
//
// This method performs variable lookup by searching through the scope hierarchy:
// 1. Checks the current scope for the identifier
// 2. If not found, recursively searches parent scopes
// 3. Returns the bound value if found, or an error if not found
//
// The scope chain lookup enables lexical scoping and closures, allowing inner
// functions to access variables from outer scopes.
//
// Parameters:
//   - n: An IdentifierExpressionNode containing the variable name to look up
//
// Returns:
//   - objects.GoMixObject: The value bound to the identifier, or an Error object
//     if the identifier is not found in any scope in the chain
//
// Example:
//
//	var x = 10;
//	func inner() { return x; }  // Looks up 'x' in parent scope
//	inner();  // Returns 10
func (e *Evaluator) evalIdentifierExpression(n *parser.IdentifierExpressionNode) objects.GoMixObject {

	val, ok := e.Scp.LookUp(n.Name)
	if !ok {
		return e.createError(n.Token, "ERROR: identifier not found: (%s)", n.Name)
	}
	return val
}

// evalBlockStatement evaluates a sequence of statements within a block.
//
// This method processes statement blocks (code between { and }) by delegating to
// evalStatements. Blocks are used in function bodies, if-else branches, and loops.
// The method returns the result of the last statement in the block, or stops early
// if a return statement or error is encountered.
//
// Note: This method does NOT create a new scope - scope creation is handled by
// the constructs that use blocks (functions, loops, etc.).
//
// Parameters:
//   - n: A BlockStatementNode containing a list of statements to evaluate
//
// Returns:
//   - objects.GoMixObject: The result of the last statement, a ReturnValue if a
//     return statement was encountered, or an Error if evaluation failed
//
// Example:
//
//	{
//	    var x = 10;
//	    var y = 20;
//	    x + y;  // Block returns 30
//	}
func (e *Evaluator) evalBlockStatement(n *parser.BlockStatementNode) objects.GoMixObject {
	return e.evalStatements(n.Statements)
}

// evalReturnStatement evaluates a return statement and wraps the result for propagation.
//
// This method handles the 'return' keyword by:
// 1. Evaluating the return expression to get the value to return
// 2. Wrapping the value in a ReturnValue object for special handling
//
// The ReturnValue wrapper is used to signal that evaluation should stop and
// propagate the return value up through nested blocks and function calls.
// The wrapper is unwrapped by evalCallExpression when returning from a function.
//
// Parameters:
//   - n: A ReturnStatementNode containing the expression to return
//
// Returns:
//   - objects.GoMixObject: A ReturnValue wrapper containing the evaluated expression,
//     or an Error object if the expression evaluation failed
//
// Example:
//
//	func add(a, b) {
//	    return a + b;  // Evaluates a + b, wraps in ReturnValue
//	}
func (e *Evaluator) evalReturnStatement(n *parser.ReturnStatementNode) objects.GoMixObject {
	val := e.Eval(n.Expr)
	if IsError(val) {
		return val
	}
	return &objects.ReturnValue{Value: val}
}

// evalDeclarativeStatement handles variable declarations with var, const, and let keywords.
//
// This method processes variable declarations by:
// 1. Evaluating the initialization expression to get the initial value
// 2. Checking for redeclaration conflicts in the current scope
// 3. Binding the variable to its value in the current scope
// 4. Recording special properties based on the declaration keyword:
//   - 'const': Marks the variable as immutable (stored in Consts map)
//   - 'let': Marks the variable as type-safe and records its type (stored in LetVars and LetTypes)
//   - 'var': Standard mutable variable with no type restrictions
//
// The distinction between declaration types affects later assignment operations:
// - const variables cannot be reassigned
// - let variables can only be assigned values of the same type
// - var variables can be reassigned to any type
//
// Parameters:
//   - n: A DeclarativeStatementNode containing the keyword, identifier, and initialization expression
//
// Returns:
//   - objects.GoMixObject: The initialized value on success, or an Error object if:
//   - The initialization expression fails
//   - The variable is already declared in the current scope
//
// Example:
//
//	var x = 10;      // Mutable, any type
//	const PI = 3.14; // Immutable
//	let name = "Go"; // Type-safe (must remain string)
func (e *Evaluator) evalDeclarativeStatement(n *parser.DeclarativeStatementNode) objects.GoMixObject {
	// fmt.Printf("DEBUG: evalDeclarativeStatement for '%s', expr type=%T\n", n.Identifier.Name, n.Expr)
	val := e.Eval(n.Expr)
	// fmt.Printf("DEBUG: evalDeclarativeStatement result type=%s\n", val.GetType())
	if IsError(val) {
		return val
	}

	// redeclared?
	_, has := e.Scp.Bind(n.Identifier.Name, val)
	if has {
		return e.CreateError("ERROR: identifier redeclaration found: (%s)", n.Identifier.Name)
	}

	if n.VarToken.Type == lexer.CONST_KEY {
		e.Scp.Consts[n.Identifier.Name] = true
	} else if n.VarToken.Type == lexer.LET_KEY {
		e.Scp.LetVars[n.Identifier.Name] = true
		e.Scp.LetTypes[n.Identifier.Name] = val.GetType()
	}
	e.Scp.Bind(n.Identifier.Name, val)
	return val
}

// evalConditionalExpression evaluates if-else conditional expressions.
//
// This method implements conditional branching by:
// 1. Evaluating the condition expression
// 2. Validating that the condition is a boolean type
// 3. Executing the then-block if the condition is true
// 4. Executing the else-block if the condition is false
//
// Both branches are represented as BlockStatementNodes, and the method returns
// the result of whichever branch is executed. If the condition is not a boolean,
// an error is returned.
//
// Parameters:
//   - n: An IfExpressionNode containing the condition and both branch blocks
//
// Returns:
//   - objects.GoMixObject: The result of the executed branch (then or else),
//     or an Error object if the condition is not a boolean type
//
// Example:
//
//	if (x > 10) {
//	    return "big";
//	} else {
//	    return "small";
//	}
func (e *Evaluator) evalConditionalExpression(n *parser.IfExpressionNode) objects.GoMixObject {
	condition := e.Eval(n.Condition)
	if IsError(condition) {
		return condition
	}

	if condition.GetType() != objects.BooleanType {
		return e.CreateError("ERROR: conditional expression must be (bool)")
	}
	if condition.(*objects.Boolean).Value {
		return e.Eval(&n.ThenBlock)
	}
	return e.Eval(&n.ElseBlock)
}

// evalStatements evaluates a sequence of statements in order, with early termination support.
//
// This method processes a list of statements sequentially and implements two important
// control flow behaviors:
//  1. Error propagation: If any statement produces an error, evaluation stops immediately
//     and the error is returned
//  2. Return handling: If any statement produces a ReturnValue, evaluation stops and the
//     return value is propagated (used to exit from functions early)
//
// For normal execution, the method continues through all statements and returns the
// result of the last one. If the list is empty, returns Nil.
//
// Parameters:
//   - stmts: A slice of StatementNode objects to evaluate in sequence
//
// Returns:
//   - objects.GoMixObject: The result of the last statement, a ReturnValue if a return
//     was encountered, an Error if any statement failed, or Nil for an empty list
//
// Example:
//
//	var x = 10;
//	var y = 20;
//	return x + y;  // Stops here, returns 30
//	var z = 30;    // Never executed
func (e *Evaluator) evalStatements(stmts []parser.StatementNode) objects.GoMixObject {
	var result objects.GoMixObject = &objects.Nil{}
	for _, stmt := range stmts {
		result = e.Eval(stmt)

		if IsError(result) {
			return result
		}
		// Stop evaluation if we hit a return statement
		if _, isReturn := result.(*objects.ReturnValue); isReturn {
			return result
		}
		// Stop evaluation if we hit break or continue
		if result.GetType() == objects.BreakType || result.GetType() == objects.ContinueType {
			return result
		}
	}
	return result
}

// evalBinaryExpression evaluates binary arithmetic and bitwise operations.
//
// This method handles infix operators that take two operands:
//
// Arithmetic operators (work with integers and floats):
//   - Addition (+): Returns integer if both operands are integers, otherwise float
//   - Subtraction (-): Same type promotion rules as addition
//   - Multiplication (*): Same type promotion rules as addition
//   - Division (/): Same type promotion rules as addition
//   - Modulo (%): Only works with integers
//
// Bitwise operators (only work with integers):
//   - Bitwise AND (&): Performs bit-by-bit AND operation
//   - Bitwise OR (|): Performs bit-by-bit OR operation
//   - Bitwise XOR (^): Performs bit-by-bit exclusive OR operation
//   - Left shift (<<): Shifts bits left by the right operand amount
//   - Right shift (>>): Shifts bits right by the right operand amount
//
// Type handling:
// - If both operands are integers, the result is an integer
// - If either operand is a float, both are converted to float and result is float
// - Bitwise operations require both operands to be integers
//
// Parameters:
//   - n: A BinaryExpressionNode containing the operator and left/right operands
//
// Returns:
//   - objects.GoMixObject: The computed result (Integer or Float), or an Error if:
//   - Either operand is not a number
//   - Operator is not supported for the operand types
//   - Bitwise operation attempted on non-integer types
//
// Example:
//
//	5 + 3      // Returns Integer(8)
//	5.0 + 3    // Returns Float(8.0)
//	10 % 3     // Returns Integer(1)
//	5 & 3      // Returns Integer(1) - bitwise AND
func (e *Evaluator) evalBinaryExpression(n *parser.BinaryExpressionNode) objects.GoMixObject {
	left := e.Eval(n.Left)

	if IsError(left) {
		return left
	}

	// we prioritize the dot (.) member access operator in the parser,
	if n.Operation.Type == lexer.DOT_OP {

		if left.GetType() != objects.ObjectType {
			return e.CreateError("ERROR: member access operator (.) can only be used on struct instances, got (%s)", left.GetType())
		}
		structInstance := left.(*objects.GoMixObjectInstance)

		// Handle Index Access on Field/Method (e.g. this.q[0])
		if indexNode, ok := n.Right.(*parser.IndexExpressionNode); ok {
			container := e.evalMemberAccess(structInstance, indexNode.Left)
			if IsError(container) {
				return container
			}
			index := e.Eval(indexNode.Index)
			if IsError(index) {
				return index
			}
			return e.getIndexValue(container, index)
		}

		return e.evalMemberAccess(structInstance, n.Right)
	}

	right := e.Eval(n.Right)
	if IsError(right) {
		return right
	}

	return e.evaluateBinaryOp(n.Operation, n.Operation.Type, left, right)
}

func (e *Evaluator) evaluateBinaryOp(token lexer.Token, opType lexer.TokenType, left, right objects.GoMixObject) objects.GoMixObject {
	err := e.createError(token, "ERROR: operator (%s) not implemented for (%s) and (%s)", token.Literal, left.GetType(), right.GetType())

	if left.GetType() == objects.StringType && right.GetType() == objects.StringType {
		if opType == lexer.PLUS_OP {
			return &objects.String{Value: left.(*objects.String).Value + right.(*objects.String).Value}
		}
		return err
	}

	if left.GetType() != objects.IntegerType && left.GetType() != objects.FloatType {
		return err
	}
	if right.GetType() != objects.IntegerType && right.GetType() != objects.FloatType {
		return err
	}

	leftType := left.GetType()
	rightType := right.GetType()

	switch opType {
	case lexer.PLUS_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value + right.(*objects.Integer).Value}
		}
		return &objects.Float{Value: toFloat64(left) + toFloat64(right)}
	case lexer.MINUS_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value - right.(*objects.Integer).Value}
		}
		return &objects.Float{Value: toFloat64(left) - toFloat64(right)}
	case lexer.MUL_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value * right.(*objects.Integer).Value}
		}
		return &objects.Float{Value: toFloat64(left) * toFloat64(right)}
	case lexer.DIV_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value / right.(*objects.Integer).Value}
		}
		return &objects.Float{Value: toFloat64(left) / toFloat64(right)}
	case lexer.MOD_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value % right.(*objects.Integer).Value}
		}
		return err
	case lexer.BIT_AND_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value & right.(*objects.Integer).Value}
		}
		return err
	case lexer.BIT_OR_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value | right.(*objects.Integer).Value}
		}
		return err
	case lexer.BIT_XOR_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value ^ right.(*objects.Integer).Value}
		}
		return err
	case lexer.BIT_LEFT_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value << right.(*objects.Integer).Value}
		}
		return err
	case lexer.BIT_RIGHT_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value >> right.(*objects.Integer).Value}
		}
		return err
	}
	return err
}

func (e *Evaluator) callFunctionOnObject(name string, obj *objects.GoMixObjectInstance, args ...NamedParameter) objects.GoMixObject {

	initMethodInterface, exists := obj.Struct.Methods[name]
	if !exists {
		return e.CreateError("ERROR: method (%s) not found in struct (%s)", name, obj.Struct.GetName())
	}

	initMethod, ok := initMethodInterface.(*function.Function)
	if !ok {
		return e.CreateError("ERROR: method (%s) not found in struct (%s)", name, obj.Struct.GetName())
	}

	// Create a new scope for the method call with the struct instance's scope as parent
	methodScope := scope.NewScope(e.Scp)

	// Bind the struct instance to a special variable (e.g., "self") in the method scope
	methodScope.Bind("this", obj)
	for _, arg := range args {
		methodScope.Bind(arg.Name, arg.Value)
	}

	// Save the current scope and switch to the method scope for evaluation
	oldScope := e.Scp
	e.Scp = methodScope
	res := e.Eval(initMethod.Body)
	e.Scp = oldScope
	if res.GetType() == objects.ErrorType {
		return res
	}
	return UnwrapReturnValue(res)
}

// toFloat64 is a helper function that converts numeric objects to float64.
//
// This function enables mixed-type arithmetic by converting both Integer and Float
// objects to the float64 primitive type. It's used when performing arithmetic
// operations where at least one operand is a float, ensuring type consistency.
//
// Parameters:
//   - obj: A GoMixObject that must be either an Integer or Float type
//
// Returns:
//   - float64: The numeric value as a float64. Integers are converted to their
//     floating-point equivalent, floats are returned as-is.
//
// Example:
//
//	toFloat64(&objects.Integer{Value: 5})   // Returns 5.0
//	toFloat64(&objects.Float{Value: 3.14})  // Returns 3.14
func toFloat64(obj objects.GoMixObject) float64 {
	if obj.GetType() == objects.IntegerType {
		return float64(obj.(*objects.Integer).Value)
	}
	return obj.(*objects.Float).Value
}

// evalUnaryExpression evaluates unary prefix operations on a single operand.
//
// This method handles operators that appear before their operand:
//
// Logical operator:
//   - NOT (!): Inverts a boolean value (true -> false, false -> true)
//     Only works with boolean operands
//
// Bitwise operator:
//   - Bitwise NOT (~): Inverts all bits in an integer
//     Only works with integer operands
//
// Arithmetic operators:
//   - Negation (-): Returns the negative of a number
//     Works with both integers and floats
//   - Unary plus (+): Returns the number unchanged (identity operation)
//     Works with both integers and floats
//
// Parameters:
//   - n: A UnaryExpressionNode containing the operator and the operand expression
//
// Returns:
//   - objects.GoMixObject: The result of applying the operator, or an Error if:
//   - The operator is not supported for the operand type
//   - Type mismatch (e.g., ! on a number, ~ on a float)
//
// Example:
//
//	!true      // Returns Boolean(false)
//	-5         // Returns Integer(-5)
//	~10        // Returns Integer(-11) - bitwise NOT
//	+3.14      // Returns Float(3.14)
func (e *Evaluator) evalUnaryExpression(n *parser.UnaryExpressionNode) objects.GoMixObject {
	right := e.Eval(n.Right)
	if IsError(right) {
		return right
	}

	err := e.createError(n.Operation, "ERROR: operator (%s) not implemented for (%s)", n.Operation.Literal, right.GetType())

	switch n.Operation.Type {
	case lexer.NOT_OP:
		if right.GetType() != objects.BooleanType {
			return err
		}
		return &objects.Boolean{Value: !right.(*objects.Boolean).Value}
	case lexer.BIT_NOT_OP:
		if right.GetType() == objects.IntegerType {
			return &objects.Integer{Value: ^right.(*objects.Integer).Value}
		}
		return err
	case lexer.MINUS_OP:
		if right.GetType() == objects.IntegerType {
			return &objects.Integer{Value: -right.(*objects.Integer).Value}
		} else if right.GetType() == objects.FloatType {
			return &objects.Float{Value: -right.(*objects.Float).Value}
		}
		return err
	case lexer.PLUS_OP:
		if right.GetType() == objects.IntegerType {
			return right
		} else if right.GetType() == objects.FloatType {
			return right
		}
		return err
	}
	return &objects.Nil{}
}

// evalBooleanExpression evaluates comparison and logical operations that produce boolean results.
//
// This method handles operators that compare values or combine boolean expressions:
//
// Equality operators (work with any types):
//   - Equal (==): Compares string representations for equality
//   - Not equal (!=): Compares string representations for inequality
//
// Comparison operators (work with numbers):
//   - Greater than (>): Returns true if left > right
//   - Less than (<): Returns true if left < right
//   - Greater than or equal (>=): Returns true if left >= right
//   - Less than or equal (<=): Returns true if left <= right
//     For mixed integer/float comparisons, both are converted to float
//
// Logical operators (work with booleans):
//   - AND (&&): Returns true only if both operands are true
//   - OR (||): Returns true if at least one operand is true
//
// Type handling:
// - Equality operators convert both sides to strings for comparison
// - Comparison operators work with integers and floats (with automatic type promotion)
// - Logical operators require both operands to be booleans
//
// Parameters:
//   - n: A BooleanExpressionNode containing the operator and left/right operands
//
// Returns:
//   - objects.GoMixObject: A Boolean object with the comparison result
//
// Example:
//
//	5 > 3           // Returns Boolean(true)
//	"hi" == "hi"    // Returns Boolean(true)
//	true && false   // Returns Boolean(false)
//	10 >= 10.0      // Returns Boolean(true) - mixed types
func (e *Evaluator) evalBooleanExpression(n *parser.BooleanExpressionNode) objects.GoMixObject {
	// Handle short-circuiting for logical operators
	if n.Operation.Type == lexer.AND_OP {
		left := e.Eval(n.Left)
		if IsError(left) {
			return left
		}
		if left.GetType() != objects.BooleanType {
			return e.createError(n.Operation, "ERROR: left operand of '&&' must be a boolean, got %s", left.GetType())
		}
		if !left.(*objects.Boolean).Value {
			return &objects.Boolean{Value: false} // short-circuit
		}
		// if left is true, the result is the boolean value of the right side
		right := e.Eval(n.Right)
		if IsError(right) {
			return right
		}
		if right.GetType() != objects.BooleanType {
			return e.createError(n.Operation, "ERROR: right operand of '&&' must be a boolean, got %s", right.GetType())
		}
		return right // it's already a boolean object
	}

	if n.Operation.Type == lexer.OR_OP {
		left := e.Eval(n.Left)
		if IsError(left) {
			return left
		}
		if left.GetType() != objects.BooleanType {
			return e.createError(n.Operation, "ERROR: left operand of '||' must be a boolean, got %s", left.GetType())
		}
		if left.(*objects.Boolean).Value {
			return &objects.Boolean{Value: true} // short-circuit
		}
		// if left is false, the result is the boolean value of the right side
		right := e.Eval(n.Right)
		if IsError(right) {
			return right
		}
		if right.GetType() != objects.BooleanType {
			return e.createError(n.Operation, "ERROR: right operand of '||' must be a boolean, got %s", right.GetType())
		}
		return right // it's already a boolean object
	}

	// For other operators, evaluate both sides
	left := e.Eval(n.Left)
	if IsError(left) {
		return left
	}
	right := e.Eval(n.Right)
	if IsError(right) {
		return right
	}
	switch n.Operation.Type {
	case lexer.EQ_OP:
		return &objects.Boolean{Value: left.ToString() == right.ToString()}
	case lexer.NE_OP:
		return &objects.Boolean{Value: left.ToString() != right.ToString()}
	case lexer.GT_OP:
		if left.GetType() == objects.IntegerType && right.GetType() == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value > right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) > toFloat64(right)}
	case lexer.LT_OP:
		if left.GetType() == objects.IntegerType && right.GetType() == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value < right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) < toFloat64(right)}
	case lexer.GE_OP:
		if left.GetType() == objects.IntegerType && right.GetType() == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value >= right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) >= toFloat64(right)}
	case lexer.LE_OP:
		if left.GetType() == objects.IntegerType && right.GetType() == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value <= right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) <= toFloat64(right)}
	}
	return &objects.Nil{}
}

// evalForLoop evaluates for loop statements with comprehensive scope management.
//
// This method implements the classic for loop with three parts:
// 1. Initializers: Executed once before the loop starts (e.g., var i = 0)
// 2. Condition: Checked before each iteration (e.g., i < 10)
// 3. Updates: Executed after each iteration (e.g., i = i + 1)
//
// Scope management (critical for correct variable scoping):
// - Loop scope: Created for the entire loop, contains initializer variables
// - Iteration scope: Created fresh for each iteration, contains body variables
// - This two-level scoping ensures:
//   - Initializer variables persist across iterations
//   - Body variables are fresh each iteration
//   - Updates can access and modify initializer variables
//
// Control flow:
// - Loop continues while condition evaluates to true
// - Stops immediately on error or return statement
// - If no condition is provided, loops indefinitely (until return/error)
//
// Parameters:
//   - n: A ForLoopStatementNode containing initializers, condition, updates, and body
//
// Returns:
//   - objects.GoMixObject: The result of the last iteration's body, a ReturnValue if
//     a return was encountered, or an Error if evaluation failed
//
// Example:
//
//	for (var i = 0; i < 5; i = i + 1) {
//	    print(i);  // Prints 0, 1, 2, 3, 4
//	}
func (e *Evaluator) evalForLoop(n *parser.ForLoopStatementNode) objects.GoMixObject {
	// Create a new scope for the entire for loop (for initializers and loop variables)
	loopScope := scope.NewScope(e.Scp)
	oldScope := e.Scp
	e.Scp = loopScope

	// Evaluate initializers in the loop scope
	for _, init := range n.Initializers {
		result := e.Eval(init)
		if IsError(result) {
			e.Scp = oldScope
			return result
		}
	}

	// Loop execution
	var result objects.GoMixObject = &objects.Nil{}
	for {
		// Evaluate condition if present
		if n.Condition != nil {
			condition := e.Eval(n.Condition)
			if IsError(condition) {
				e.Scp = oldScope
				return condition
			}

			// Check if condition is false
			if condition.GetType() != objects.BooleanType {
				e.Scp = oldScope
				return e.CreateError("ERROR: for loop condition must be (bool)")
			}
			if !condition.(*objects.Boolean).Value {
				break
			}
		}

		// Create a new scope for each iteration of the loop body
		// This ensures variables declared in the body are scoped to that iteration
		iterationScope := scope.NewScope(loopScope)
		e.Scp = iterationScope

		// Execute loop body
		result = e.Eval(&n.Body)

		// Restore to loop scope after body execution
		e.Scp = loopScope

		if IsError(result) {
			e.Scp = oldScope
			return result
		}

		// Stop if we hit a return statement
		if _, isReturn := result.(*objects.ReturnValue); isReturn {
			e.Scp = oldScope
			return result
		}

		if result.GetType() == objects.BreakType {
			result = &objects.Nil{}
			break
		}

		if result.GetType() == objects.ContinueType {
			result = &objects.Nil{}
			// continue to updates
		}

		// Evaluate updates in the loop scope (not iteration scope)
		for _, update := range n.Updates {
			updateResult := e.Eval(update)
			if IsError(updateResult) {
				e.Scp = oldScope
				return updateResult
			}

			if result.GetType() == objects.BreakType {
				e.Scp = oldScope
				return &objects.Nil{}
			}

			if result.GetType() == objects.ContinueType {
				continue
			}
		}
	}

	// Restore the original scope
	e.Scp = oldScope
	return result
}

// evalArrayExpression evaluates array literal expressions to create array objects.
//
// This method processes array literals (e.g., [1, 2, 3]) by:
// 1. Evaluating each element expression in order
// 2. Collecting the results into a slice
// 3. Creating an Array object containing all evaluated elements
//
// Arrays in GoMix are heterogeneous - they can contain elements of different types.
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
func (e *Evaluator) evalArrayExpression(n *parser.ArrayExpressionNode) objects.GoMixObject {
	elements := make([]objects.GoMixObject, len(n.Elements))
	for i, elem := range n.Elements {
		evaluated := e.Eval(elem)
		if IsError(evaluated) {
			return evaluated
		}
		elements[i] = evaluated
	}
	return &objects.Array{Elements: elements}
}

// evalMapExpression evaluates map literal expressions to create map objects.
//
// This method processes map literals (e.g., map{10: 20, "key": "value"}) by:
// 1. Evaluating each key expression in order
// 2. Evaluating each corresponding value expression
// 3. Converting keys to strings for storage (Go maps require hashable keys)
// 4. Creating a Map object with the key-value pairs
//
// Maps in GoMix:
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
func (e *Evaluator) evalMapExpression(n *parser.MapExpressionNode) objects.GoMixObject {
	pairs := make(map[string]objects.GoMixObject)
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

	return &objects.Map{
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
// Sets in GoMix:
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
func (e *Evaluator) evalSetExpression(n *parser.SetExpressionNode) objects.GoMixObject {
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

	return &objects.Set{
		Elements: elements,
		Values:   values,
	}
}

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
func (e *Evaluator) evalIndexExpression(n *parser.IndexExpressionNode) objects.GoMixObject {
	left := e.Eval(n.Left)
	if IsError(left) {
		return left
	}

	index := e.Eval(n.Index)
	if IsError(index) {
		return index
	}

	// Handle map indexing
	if left.GetType() == objects.MapType {
		mapObj := left.(*objects.Map)
		keyStr := index.ToString()

		if value, exists := mapObj.Pairs[keyStr]; exists {
			return value
		}
		// Return nil if key doesn't exist
		return &objects.Nil{}
	}

	// Handle range indexing
	if left.GetType() == objects.RangeType {
		return e.evalRangeIndexExpression(left, index)
	}

	// Handle array, list, and tuple indexing
	leftType := left.GetType()
	if leftType != objects.ArrayType && leftType != objects.ListType && leftType != objects.TupleType {
		return e.CreateError("ERROR: index operator not supported for type '%s'", leftType)
	}

	// Check if index is an integer
	if index.GetType() != objects.IntegerType {
		return e.CreateError("ERROR: index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*objects.Integer).Value
	var length int64
	var elements []objects.GoMixObject

	// Get elements based on type
	switch leftType {
	case objects.ArrayType:
		arr := left.(*objects.Array)
		elements = arr.Elements
		length = int64(len(arr.Elements))
	case objects.ListType:
		list := left.(*objects.List)
		elements = list.Elements
		length = int64(len(list.Elements))
	case objects.TupleType:
		tuple := left.(*objects.Tuple)
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
// Returns the value at the specified index in the range sequence.
// Example: range(1, 5)[0] returns 1, range(1, 5)[2] returns 3
func (e *Evaluator) evalRangeIndexExpression(left, index objects.GoMixObject) objects.GoMixObject {
	r := left.(*objects.Range)

	// Check if index is an integer
	if index.GetType() != objects.IntegerType {
		return e.CreateError("ERROR: range index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*objects.Integer).Value
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

	return &objects.Integer{Value: value}
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
func (e *Evaluator) evalSliceExpression(n *parser.SliceExpressionNode) objects.GoMixObject {
	left := e.Eval(n.Left)
	if IsError(left) {
		return left
	}

	// Check if left is an array, list, or tuple
	leftType := left.GetType()
	if leftType != objects.ArrayType && leftType != objects.ListType && leftType != objects.TupleType {
		return e.CreateError("ERROR: slice operator not supported for type '%s'", leftType)
	}

	var elements []objects.GoMixObject
	var length int64

	// Get elements based on type
	switch leftType {
	case objects.ArrayType:
		arr := left.(*objects.Array)
		elements = arr.Elements
		length = int64(len(arr.Elements))
	case objects.ListType:
		list := left.(*objects.List)
		elements = list.Elements
		length = int64(len(list.Elements))
	case objects.TupleType:
		tuple := left.(*objects.Tuple)
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
		if startObj.GetType() != objects.IntegerType {
			return e.CreateError("ERROR: slice start index must be an integer, got '%s'", startObj.GetType())
		}
		start = startObj.(*objects.Integer).Value
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
		if endObj.GetType() != objects.IntegerType {
			return e.CreateError("ERROR: slice end index must be an integer, got '%s'", endObj.GetType())
		}
		end = endObj.(*objects.Integer).Value
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
	slicedElements := make([]objects.GoMixObject, end-start)
	copy(slicedElements, elements[start:end])

	return &objects.Array{Elements: slicedElements}
}

func (e *Evaluator) getIndexValue(container, index objects.GoMixObject) objects.GoMixObject {
	if container.GetType() == objects.MapType {
		mapObj := container.(*objects.Map)
		keyStr := index.ToString()
		if value, exists := mapObj.Pairs[keyStr]; exists {
			return value
		}
		return &objects.Nil{}
	}

	leftType := container.GetType()
	if leftType != objects.ArrayType && leftType != objects.ListType {
		return e.CreateError("ERROR: index operator not supported for type '%s'", leftType)
	}

	if index.GetType() != objects.IntegerType {
		return e.CreateError("ERROR: index must be an integer, got '%s'", index.GetType())
	}

	idx := index.(*objects.Integer).Value
	var length int64
	var elements []objects.GoMixObject

	if leftType == objects.ArrayType {
		arr := container.(*objects.Array)
		elements = arr.Elements
		length = int64(len(arr.Elements))
	} else {
		list := container.(*objects.List)
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

// evalWhileLoop evaluates while loop statements with multiple condition support.
//
// This method implements while loops with the following features:
// 1. Supports multiple conditions that are implicitly AND-ed together
// 2. Creates a loop scope for the entire while loop
// 3. Creates a fresh iteration scope for each loop iteration
// 4. Continues looping while all conditions evaluate to true
// 5. Stops on error, return statement, or when any condition becomes false
//
// Scope management (similar to for loops):
// - Loop scope: Created for the entire loop, persists across iterations
// - Iteration scope: Created fresh for each iteration, contains body variables
// - This ensures variables declared in the loop body don't leak between iterations
//
// Condition evaluation:
// - All conditions must be boolean expressions
// - Conditions are evaluated in order before each iteration
// - If any condition is false, the loop terminates
// - If all conditions are true, the body executes
//
// Parameters:
//   - n: A WhileLoopStatementNode containing the condition expressions and body
//
// Returns:
//   - objects.GoMixObject: The result of the last iteration's body, a ReturnValue if
//     a return was encountered, or an Error if evaluation failed
//
// Example:
//
//	var i = 0;
//	while (i < 5) {
//	    print(i);
//	    i = i + 1;
//	}
//
//	// Multiple conditions (AND-ed together):
//	while (x > 0, y < 10) {
//	    // Continues only while both conditions are true
//	}
func (e *Evaluator) evalWhileLoop(n *parser.WhileLoopStatementNode) objects.GoMixObject {
	// Create a new scope for the entire while loop
	loopScope := scope.NewScope(e.Scp)
	oldScope := e.Scp
	e.Scp = loopScope

	var result objects.GoMixObject = &objects.Nil{}

	for {
		// Evaluate all conditions (they should be AND-ed together)
		allTrue := true
		for _, cond := range n.Conditions {
			condition := e.Eval(cond)
			if IsError(condition) {
				e.Scp = oldScope
				return condition
			}

			if condition.GetType() != objects.BooleanType {
				e.Scp = oldScope
				return e.CreateError("ERROR: while loop condition must be (bool)")
			}

			if !condition.(*objects.Boolean).Value {
				allTrue = false
				break
			}
		}

		if !allTrue {
			break
		}

		// Create a new scope for each iteration of the loop body
		// This ensures variables declared in the body are scoped to that iteration
		iterationScope := scope.NewScope(loopScope)
		e.Scp = iterationScope

		// Execute loop body
		result = e.Eval(&n.Body)

		// Restore to loop scope after body execution
		e.Scp = loopScope

		if IsError(result) {
			e.Scp = oldScope
			return result
		}

		// Stop if we hit a return statement
		if _, isReturn := result.(*objects.ReturnValue); isReturn {
			e.Scp = oldScope
			return result
		}

		if result.GetType() == objects.BreakType {
			result = &objects.Nil{}
			break
		}

		if result.GetType() == objects.ContinueType {
			result = &objects.Nil{}
			continue
		}
	}

	// Restore the original scope
	e.Scp = oldScope
	return result
}

func (e *Evaluator) evalStructDeclaration(n *parser.StructDeclarationNode) objects.GoMixObject {
	// Create a new struct type with the given name and fields
	s := &objects.GoMixStruct{
		Name:       n.StructName.Name,
		Methods:    make(map[string]objects.FunctionInterface),
		FieldNodes: make([]interface{}, len(n.Fields)),
	}

	for i, f := range n.Fields {
		s.FieldNodes[i] = f
	}

	for _, m := range n.Methods {
		method := &function.Function{
			Name:   m.FuncName.Name,
			Params: m.FuncParams,
			Body:   &m.FuncBody,
			Scp:    e.Scp, // Capture the current scope for closures
		}
		if err := s.Add(method); err != nil {
			return e.CreateError("ERROR: struct method '%s' already defined", method.Name)
		}
	}

	e.Types[s.Name] = s
	return s
}

func (e *Evaluator) evalNewCallExpression(n *parser.NewCallExpressionNode) objects.GoMixObject {
	// Look up the struct type by name
	s, exists := e.Types[n.StructName.Name]
	if !exists {
		return e.CreateError("ERROR: struct type '%s' not defined", n.StructName.Name)
	}

	inst := objects.NewStructInstance(s)

	// Initialize fields from struct definition
	for _, fieldNode := range s.FieldNodes {
		decl := fieldNode.(*parser.DeclarativeStatementNode)
		val := e.Eval(decl.Expr)
		if IsError(val) {
			return val
		}
		inst.Fields[decl.Identifier.Name] = val
	}

	initMethod, hasInit := s.GetConstructor()
	if hasInit {
		// Cast to Function to access Body and Params directly
		fn, ok := initMethod.(*function.Function)
		if !ok {
			return e.CreateError("ERROR: constructor method is not a valid function")
		}

		if len(n.Arguments) != len(fn.Params) {
			return e.CreateError("ERROR: constructor for struct '%s' expects %d arguments, got %d", s.Name, len(fn.Params), len(n.Arguments))
		}

		// Save the current scope before creating a new one
		oldScope := e.Scp

		// Create a new scope for the constructor call
		constructorScope := scope.NewScope(e.Scp)
		constructorScope.Bind("this", inst) // Set 'this' to the new instance

		// Evaluate the constructor with the given arguments
		for i, arg := range n.Arguments {
			argValue := e.Eval(arg)
			if IsError(argValue) {
				e.Scp = oldScope
				return argValue
			}
			constructorScope.Bind(fn.Params[i].Name, argValue)
		}

		// Switch to the constructor scope
		e.Scp = constructorScope

		// Execute the constructor body
		result := e.Eval(fn.Body)
		if IsError(result) {
			e.Scp = oldScope
			return result
		}

		// Restore the original scope
		e.Scp = oldScope
	}
	return inst
}

// createError creates an error object with line and column information from the token
func (e *Evaluator) createError(token lexer.Token, format string, args ...interface{}) objects.GoMixObject {
	return &objects.Error{
		Message: fmt.Sprintf("[%d:%d] %s", token.Line, token.Column, fmt.Sprintf(format, args...)),
	}
}

func (e *Evaluator) evalMemberAccess(structInstance *objects.GoMixObjectInstance, node parser.ExpressionNode) objects.GoMixObject {
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
		if res.GetType() == objects.ErrorType {
			return res
		}
		return res
	}

	// Handle Field Access
	if ident, ok := node.(*parser.IdentifierExpressionNode); ok {
		fieldName := ident.Name
		if val, ok := structInstance.Fields[fieldName]; ok {
			return val
		}
		return e.CreateError("ERROR: field (%s) not found in struct instance", fieldName)
	}

	return e.CreateError("ERROR: member access operator (.) must be followed by a function call or identifier")
}
