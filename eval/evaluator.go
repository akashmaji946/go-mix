package eval

import (
	"fmt"
	"testing"

	"github.com/akashmaji946/go-mix/function"
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
	"github.com/akashmaji946/go-mix/scope"
)

// Evaluates expressions in a repl
type Evaluator struct {
	Par      *parser.Parser
	Scp      *scope.Scope
	Builtins map[string]*objects.Builtin
}

// Evaluator constructor
func NewEvaluator() *Evaluator {
	ev := &Evaluator{
		Par:      nil,
		Scp:      scope.NewScope(nil),
		Builtins: make(map[string]*objects.Builtin),
	}
	for _, builtin := range objects.Builtins {
		ev.Builtins[builtin.Name] = builtin
	}
	return ev
}

func (e *Evaluator) IsBuiltin(name string) bool {
	_, ok := e.Builtins[name]
	return ok
}

func (e *Evaluator) InvokeBuiltin(name string, args ...objects.GoMixObject) objects.GoMixObject {

	if builtin, ok := e.Builtins[name]; ok {
		return builtin.Callback(args...)
	}
	return &objects.Nil{}
}

func IsError(obj objects.GoMixObject) bool {
	if obj != nil {
		return obj.GetType() == objects.ErrorType
	}
	return false
}

func CreateError(format string, a ...interface{}) *objects.Error {
	msg := fmt.Sprintf(format, a...)
	msg = fmt.Sprintf("%s", msg)
	return &objects.Error{Message: msg}
}

func UnwrapReturnValue(obj objects.GoMixObject) objects.GoMixObject {
	if retVal, isReturn := obj.(*objects.ReturnValue); isReturn {
		return retVal.Value
	}
	return obj
}

func AssertError(t *testing.T, obj objects.GoMixObject, expected string) {
	errObj, ok := obj.(*objects.Error)
	if !ok {
		t.Errorf("not error. got=%T (%+v)", obj, obj)
		return
	}
	if errObj.Message != expected {
		t.Errorf("wrong error message. expected=%q, got=%q", expected, errObj.Message)
	}
}

func AssertInteger(t *testing.T, obj objects.GoMixObject, expected int64) {
	result, ok := obj.(*objects.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
}

func AssertBoolean(t *testing.T, obj objects.GoMixObject, expected bool) {
	result, ok := obj.(*objects.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
	}

}

func AssertFloat(t *testing.T, obj objects.GoMixObject, expected float64) {
	result, ok := obj.(*objects.Float)
	if !ok {
		t.Errorf("object is not Float. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%f, want=%f", result.Value, expected)
	}

}

func AssertNil(t *testing.T, obj objects.GoMixObject) {
	if obj != nil {
		t.Errorf("object is not nil. got=%T (%+v)", obj, obj)
		return
	}
}

func AssertString(t *testing.T, obj objects.GoMixObject, expected string) {
	result, ok := obj.(*objects.String)
	if !ok {
		t.Errorf("object is not String. got=%T (%+v)", obj, obj)
		return
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%q, want=%q", result.Value, expected)
	}
}

// Evaluates the given node into a GoMixObject
// If the node is a statement a Nil object is returned
// Errors are GoMixObject instances as well,
// and they are designed to block the evaluation process.
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
		return e.registerFunction(n)
	case *parser.CallExpressionNode:
		return e.evalCallExpression(n)
	case *parser.AssignmentExpressionNode:
		return e.evalAssignmentExpression(n)
	default:
		return &objects.Nil{}
	}
}

func (e *Evaluator) SetParser(p *parser.Parser) {
	e.Par = p
}

func (e *Evaluator) registerFunction(n *parser.FunctionStatementNode) objects.GoMixObject {
	function := &function.Function{
		Name:   n.FuncName.Name,
		Params: n.FuncParams,
		Body:   &n.FuncBody,
		Scp:    e.Scp, // Reference the current scope directly, not a copy
	}
	// redeclared?
	name, has := e.Scp.Bind(n.FuncName.Name, function)
	if has && name != "" {
		return CreateError("function redeclaration found: (%s)", n.FuncName.Name)
	}
	e.Scp.Bind(n.FuncName.Name, function)
	return function
}

func (e *Evaluator) evalAssignmentExpression(n *parser.AssignmentExpressionNode) objects.GoMixObject {
	val := e.Eval(n.Right)
	if IsError(val) {
		return val
	}
	// Check if the variable exists in the current scope or any parent scope
	_, exists := e.Scp.LookUp(n.Left.Name)
	if !exists {
		return CreateError("identifier not found: (%s)", n.Left.Name)
	}
	// Check if it's a constant using the new IsConstant method
	if e.Scp.IsConstant(n.Left.Name) {
		return CreateError("can't assign to constant (%s)", n.Left.Name)
	}
	// Use Assign to update the variable in the scope where it was defined
	// This is essential for closures to work correctly
	e.Scp.Assign(n.Left.Name, val)

	return val
}

func (e *Evaluator) evalCallExpression(n *parser.CallExpressionNode) objects.GoMixObject {

	// look for builtin name
	funcName := n.FunctionIdentifier.Name
	if ok := e.IsBuiltin(funcName); ok {
		args := make([]objects.GoMixObject, len(n.Arguments))
		for i, arg := range n.Arguments {
			args[i] = e.Eval(arg)
		}
		rv := e.InvokeBuiltin(funcName, args...)
		if rv.GetType() != objects.ErrorType {
			fmt.Println("")
		}
		return rv
	}

	// lookup for function name
	obj, ok := e.Scp.LookUp(funcName)
	if !ok {
		return CreateError("function not found: (%s)", funcName)
	}
	if obj.GetType() != objects.FunctionType {
		return CreateError("not a function: (%s)", funcName)
	}
	functionObject := obj.(*function.Function)

	// Validate argument count
	expectedArgs := len(functionObject.Params)
	actualArgs := len(n.Arguments)
	if actualArgs != expectedArgs {
		return CreateError("wrong number of arguments: expected %d, got %d", expectedArgs, actualArgs)
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
		callSiteScope.Bind(param.Name, e.Eval(n.Arguments[i]))
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

func (e *Evaluator) evalIdentifierExpression(n *parser.IdentifierExpressionNode) objects.GoMixObject {
	// if val, ok := e.parser.Env[n.Name]; ok {
	// 	return val
	// }
	// return &objects.Nil{}
	val, ok := e.Scp.LookUp(n.Name)
	if !ok {
		return CreateError("identifier not found: (%s)", n.Name)
	}
	return val
}

func (e *Evaluator) evalBlockStatement(n *parser.BlockStatementNode) objects.GoMixObject {
	return e.evalStatements(n.Statements)
}

func (e *Evaluator) evalReturnStatement(n *parser.ReturnStatementNode) objects.GoMixObject {
	val := e.Eval(n.Expr)
	if IsError(val) {
		return val
	}
	return &objects.ReturnValue{Value: val}
}

func (e *Evaluator) evalDeclarativeStatement(n *parser.DeclarativeStatementNode) objects.GoMixObject {
	val := e.Eval(n.Expr)
	if IsError(val) {
		return val
	}
	// redeclared?
	_, has := e.Scp.Bind(n.Identifier.Name, val)
	if has {
		return CreateError("identifier redeclaration found: (%s)", n.Identifier.Name)
	}

	if n.VarToken.Type == lexer.CONST_KEY {
		e.Scp.Consts[n.Identifier.Name] = true
	}
	e.Scp.Bind(n.Identifier.Name, val)
	return val
}

func (e *Evaluator) evalConditionalExpression(n *parser.IfExpressionNode) objects.GoMixObject {
	condition := e.Eval(n.Condition)
	if IsError(condition) {
		return condition
	}

	if condition.GetType() != objects.BooleanType {
		return CreateError("Conditional expression must be (bool)")
	}
	if condition.(*objects.Boolean).Value {
		return e.Eval(&n.ThenBlock)
	}
	return e.Eval(&n.ElseBlock)
}

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
	}
	return result
}

func (e *Evaluator) evalBinaryExpression(n *parser.BinaryExpressionNode) objects.GoMixObject {
	left := e.Eval(n.Left)
	right := e.Eval(n.Right)

	if IsError(left) {
		return left
	}
	if IsError(right) {
		return right
	}

	err := CreateError("operator (%s) not implemented for (%s) and (%s)", n.Operation.Literal, left.GetType(), right.GetType())

	if left.GetType() != objects.IntegerType && left.GetType() != objects.FloatType {
		return err
	}
	if right.GetType() != objects.IntegerType && right.GetType() != objects.FloatType {
		return err
	}

	leftType := left.GetType()
	rightType := right.GetType()

	switch n.Operation.Type {
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
	return &objects.Nil{}
}

func toFloat64(obj objects.GoMixObject) float64 {
	if obj.GetType() == objects.IntegerType {
		return float64(obj.(*objects.Integer).Value)
	}
	return obj.(*objects.Float).Value
}

func (e *Evaluator) evalUnaryExpression(n *parser.UnaryExpressionNode) objects.GoMixObject {
	right := e.Eval(n.Right)
	if IsError(right) {
		return right
	}

	err := CreateError("operator (%s) not implemented for (%s)", n.Operation.Literal, right.GetType())

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

func (e *Evaluator) evalBooleanExpression(n *parser.BooleanExpressionNode) objects.GoMixObject {
	left := e.Eval(n.Left)
	right := e.Eval(n.Right)

	leftType := left.GetType()
	rightType := right.GetType()

	switch n.Operation.Type {
	case lexer.EQ_OP:
		return &objects.Boolean{Value: left.ToString() == right.ToString()}
	case lexer.NE_OP:
		return &objects.Boolean{Value: left.ToString() != right.ToString()}
	case lexer.GT_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value > right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) > toFloat64(right)}
	case lexer.LT_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value < right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) < toFloat64(right)}
	case lexer.GE_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value >= right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) >= toFloat64(right)}
	case lexer.LE_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Boolean{Value: left.(*objects.Integer).Value <= right.(*objects.Integer).Value}
		}
		return &objects.Boolean{Value: toFloat64(left) <= toFloat64(right)}
	case lexer.AND_OP:
		return &objects.Boolean{Value: left.(*objects.Boolean).Value && right.(*objects.Boolean).Value}
	case lexer.OR_OP:
		return &objects.Boolean{Value: left.(*objects.Boolean).Value || right.(*objects.Boolean).Value}
	}
	return &objects.Nil{}
}
