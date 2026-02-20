/*
File    : go-mix/eval/eval_expressions.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package eval

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/std"
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
func (e *Evaluator) Eval(n parser.Node) std.GoMixObject {
	switch n := n.(type) {
	case *parser.RootNode:
		result := e.evalStatements(n.Statements)
		return UnwrapReturnValue(result)
	case *parser.BooleanLiteralExpressionNode:
		return n.Value
	case *parser.IntegerLiteralExpressionNode:
		return n.Value
	case *parser.CharLiteralExpressionNode:
		return n.Value
	case *parser.StringLiteralExpressionNode:
		return n.Value
	case *parser.FloatLiteralExpressionNode:
		return n.Value
	case *parser.NilLiteralExpressionNode:
		return &std.Nil{}
	case *parser.BinaryExpressionNode:
		return e.evalBinaryExpression(n)
	case *parser.UnaryExpressionNode:
		return e.evalUnaryExpression(n)
	case *parser.BooleanExpressionNode:
		return e.evalBooleanExpression(n)
	case *parser.ParenthesizedExpressionNode:
		return e.Eval(n.Expr)
	case *parser.IfExpressionNode:
		return e.evalIfExpression(n)
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
		return &std.Break{}
	case *parser.ContinueStatementNode:
		return &std.Continue{}
	case *parser.ImportStatementNode:
		return e.evalImportStatement(n)
	case *parser.EnumDeclarationNode:
		return e.evalEnumDeclaration(n)
	case *parser.EnumAccessExpressionNode:
		return e.evalEnumAccessExpression(n)
	case *parser.SwitchStatementNode:
		return e.evalSwitchStatement(*n)
	default:
		return &std.Nil{}
	}
}

// evaluateBinaryOp performs the actual computation for binary operations.
//
// This helper method handles arithmetic (+, -, *, /, %), bitwise (&, |, ^, <<, >>),
// and string concatenation operations. It performs type checking and promotion
// (int vs float) before executing the operation.
//
// Parameters:
//   - token: The operator token (for error reporting)
//   - opType: The type of binary operator
//   - left: The left operand
//   - right: The right operand
//
// Returns:
//   - objects.GoMixObject: The result of the operation, or an Error if types are incompatible
func (e *Evaluator) evaluateBinaryOp(token lexer.Token, opType lexer.TokenType, left, right std.GoMixObject) std.GoMixObject {
	err := e.createError(token, "ERROR: operator (%s) not implemented for (%s) and (%s)", token.Literal, left.GetType(), right.GetType())

	if opType == lexer.PLUS_OP {
		if left.GetType() == std.StringType || right.GetType() == std.StringType {
			return &std.String{Value: left.ToString() + right.ToString()}
		}
	}

	if left.GetType() == std.StringType && right.GetType() == std.StringType {
		if opType == lexer.PLUS_OP {
			return &std.String{Value: left.(*std.String).Value + right.(*std.String).Value}
		}
		return err
	}

	if left.GetType() != std.IntegerType && left.GetType() != std.FloatType {
		return err
	}
	if right.GetType() != std.IntegerType && right.GetType() != std.FloatType {
		return err
	}

	leftType := left.GetType()
	rightType := right.GetType()

	switch opType {
	case lexer.PLUS_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value + right.(*std.Integer).Value}
		}
		return &std.Float{Value: toFloat64(left) + toFloat64(right)}
	case lexer.MINUS_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value - right.(*std.Integer).Value}
		}
		return &std.Float{Value: toFloat64(left) - toFloat64(right)}
	case lexer.MUL_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value * right.(*std.Integer).Value}
		}
		return &std.Float{Value: toFloat64(left) * toFloat64(right)}
	case lexer.DIV_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value / right.(*std.Integer).Value}
		}
		return &std.Float{Value: toFloat64(left) / toFloat64(right)}
	case lexer.MOD_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value % right.(*std.Integer).Value}
		}
		return err
	case lexer.BIT_AND_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value & right.(*std.Integer).Value}
		}
		return err
	case lexer.BIT_OR_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value | right.(*std.Integer).Value}
		}
		return err
	case lexer.BIT_XOR_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value ^ right.(*std.Integer).Value}
		}
		return err
	case lexer.BIT_LEFT_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value << right.(*std.Integer).Value}
		}
		return err
	case lexer.BIT_RIGHT_OP:
		if leftType == std.IntegerType && rightType == std.IntegerType {
			return &std.Integer{Value: left.(*std.Integer).Value >> right.(*std.Integer).Value}
		}
		return err
	}
	return err
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
func (e *Evaluator) evalIdentifierExpression(n *parser.IdentifierExpressionNode) std.GoMixObject {

	val, ok := e.Scp.LookUp(n.Name)
	if !ok {
		return e.createError(n.Token, "ERROR: identifier not found: (%s)", n.Name)
	}
	return val
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
func (e *Evaluator) evalBinaryExpression(n *parser.BinaryExpressionNode) std.GoMixObject {
	left := e.Eval(n.Left)

	if IsError(left) {
		return left
	}

	// we prioritize the dot (.) member access operator in the parser,
	if n.Operation.Type == lexer.DOT_OP {

		if left.GetType() == std.StructType {
			return e.evalStructMemberAccess(left.(*std.GoMixStruct), n.Right)
		}

		// Handle enum member access (e.g., Color.RED)
		if left.GetType() == std.EnumType {
			enumType := left.(*std.GoMixEnum)
			ident, ok := n.Right.(*parser.IdentifierExpressionNode)
			if !ok {
				return e.CreateError("ERROR: enum member access must be an identifier")
			}
			memberValue, exists := enumType.Members[ident.Name]
			if !exists {
				return e.CreateError("ERROR: enum member '%s' not found in enum '%s'", ident.Name, enumType.Name)
			}
			return memberValue
		}

		// Handle package member access (e.g., math.abs or math.abs(...))
		if left.GetType() == std.PackageType {
			pkg := left.(*std.Package)

			// If the right side is a call expression, invoke the package function
			if callNode, ok := n.Right.(*parser.CallExpressionNode); ok {
				funcName := callNode.FunctionIdentifier.Name
				fn, exists := pkg.Functions[funcName]
				if !exists {
					return e.createError(callNode.FunctionIdentifier.Token, "ERROR: function '%s' not found in package '%s'", funcName, pkg.Name)
				}

				// Evaluate arguments
				args := make([]std.GoMixObject, len(callNode.Arguments))
				for i, arg := range callNode.Arguments {
					args[i] = e.Eval(arg)
					if IsError(args[i]) {
						return args[i]
					}
				}

				// Call the package function
				return fn.Callback(e, e.Writer, args...)
			}

			// For non-call access, just validate the function exists
			return e.evalPackageMemberAccess(pkg, n.Right)
		}

		if left.GetType() != std.ObjectType {
			return e.CreateError("ERROR: member access operator (.) can only be used on struct instances, packages, or types, got (%s)", left.GetType())
		}
		structInstance := left.(*std.GoMixObjectInstance)

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
func (e *Evaluator) evalUnaryExpression(n *parser.UnaryExpressionNode) std.GoMixObject {
	right := e.Eval(n.Right)
	if IsError(right) {
		return right
	}

	err := e.createError(n.Operation, "ERROR: operator (%s) not implemented for (%s)", n.Operation.Literal, right.GetType())

	switch n.Operation.Type {
	case lexer.NOT_OP:
		if right.GetType() != std.BooleanType {
			return err
		}
		return &std.Boolean{Value: !right.(*std.Boolean).Value}
	case lexer.BIT_NOT_OP:
		if right.GetType() == std.IntegerType {
			return &std.Integer{Value: ^right.(*std.Integer).Value}
		}
		return err
	case lexer.MINUS_OP:
		if right.GetType() == std.IntegerType {
			return &std.Integer{Value: -right.(*std.Integer).Value}
		} else if right.GetType() == std.FloatType {
			return &std.Float{Value: -right.(*std.Float).Value}
		}
		return err
	case lexer.PLUS_OP:
		if right.GetType() == std.IntegerType {
			return right
		} else if right.GetType() == std.FloatType {
			return right
		}
		return err
	}
	return &std.Nil{}
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
func (e *Evaluator) evalBooleanExpression(n *parser.BooleanExpressionNode) std.GoMixObject {
	// Handle short-circuiting for logical operators
	if n.Operation.Type == lexer.AND_OP {
		left := e.Eval(n.Left)
		if IsError(left) {
			return left
		}
		if left.GetType() != std.BooleanType {
			return e.createError(n.Operation, "ERROR: left operand of '&&' must be a boolean, got %s", left.GetType())
		}
		if !left.(*std.Boolean).Value {
			return &std.Boolean{Value: false} // short-circuit
		}
		// if left is true, the result is the boolean value of the right side
		right := e.Eval(n.Right)
		if IsError(right) {
			return right
		}
		if right.GetType() != std.BooleanType {
			return e.createError(n.Operation, "ERROR: right operand of '&&' must be a boolean, got %s", right.GetType())
		}
		return right // it's already a boolean object
	}

	if n.Operation.Type == lexer.OR_OP {
		left := e.Eval(n.Left)
		if IsError(left) {
			return left
		}
		if left.GetType() != std.BooleanType {
			return e.createError(n.Operation, "ERROR: left operand of '||' must be a boolean, got %s", left.GetType())
		}
		if left.(*std.Boolean).Value {
			return &std.Boolean{Value: true} // short-circuit
		}
		// if left is false, the result is the boolean value of the right side
		right := e.Eval(n.Right)
		if IsError(right) {
			return right
		}
		if right.GetType() != std.BooleanType {
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
		return &std.Boolean{Value: left.ToString() == right.ToString()}
	case lexer.NE_OP:
		return &std.Boolean{Value: left.ToString() != right.ToString()}
	case lexer.STRICT_EQ_OP:
		return &std.Boolean{Value: StrictEqual(left, right)}
	case lexer.STRICT_NE_OP:
		return &std.Boolean{Value: !StrictEqual(left, right)}
	case lexer.GT_OP:
		if left.GetType() == std.IntegerType && right.GetType() == std.IntegerType {
			return &std.Boolean{Value: left.(*std.Integer).Value > right.(*std.Integer).Value}
		}
		return &std.Boolean{Value: toFloat64(left) > toFloat64(right)}
	case lexer.LT_OP:
		if left.GetType() == std.IntegerType && right.GetType() == std.IntegerType {
			return &std.Boolean{Value: left.(*std.Integer).Value < right.(*std.Integer).Value}
		}
		return &std.Boolean{Value: toFloat64(left) < toFloat64(right)}
	case lexer.GE_OP:
		if left.GetType() == std.IntegerType && right.GetType() == std.IntegerType {
			return &std.Boolean{Value: left.(*std.Integer).Value >= right.(*std.Integer).Value}
		}
		return &std.Boolean{Value: toFloat64(left) >= toFloat64(right)}
	case lexer.LE_OP:
		if left.GetType() == std.IntegerType && right.GetType() == std.IntegerType {
			return &std.Boolean{Value: left.(*std.Integer).Value <= right.(*std.Integer).Value}
		}
		return &std.Boolean{Value: toFloat64(left) <= toFloat64(right)}
	}
	return &std.Nil{}
}
