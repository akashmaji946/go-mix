package eval

import (
	"fmt"
	"testing"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/objects"
	"github.com/akashmaji946/go-mix/parser"
)

type Evaluator struct {
	parser *parser.Parser
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func IsError(obj objects.GoMixObject) bool {
	if obj != nil {
		return obj.GetType() == objects.ErrorType
	}
	return false
}

func CreateError(format string, a ...interface{}) *objects.Error {
	msg := fmt.Sprintf(format, a...)
	msg = fmt.Sprintf("[ERROR]: %s", msg)
	return &objects.Error{Message: msg}
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

func (e *Evaluator) Eval(n parser.Node) objects.GoMixObject {
	switch n := n.(type) {
	case *parser.RootNode:
		return e.evalStatements(n.Statements)
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
	case *parser.CallExpressionNode:
		return e.evalCallExpression(n)
	default:
		return CreateError("Not implemented.")
	}
}

func (e *Evaluator) SetParser(p *parser.Parser) {
	e.parser = p
}

func (e *Evaluator) evalCallExpression(n *parser.CallExpressionNode) objects.GoMixObject {
	return &objects.Nil{}
}

func (e *Evaluator) evalIdentifierExpression(n *parser.IdentifierExpressionNode) objects.GoMixObject {
	if val, ok := e.parser.Env[n.Name]; ok {
		return val
	}
	return &objects.Nil{}
}

func (e *Evaluator) evalBlockStatement(n *parser.BlockStatementNode) objects.GoMixObject {
	return e.evalStatements(n.Statements)
}

func (e *Evaluator) evalReturnStatement(n *parser.ReturnStatementNode) objects.GoMixObject {
	val := e.Eval(n.Expr)
	if IsError(val) {
		return val
	}
	return val
}

func (e *Evaluator) evalDeclarativeStatement(n *parser.DeclarativeStatementNode) objects.GoMixObject {
	val := e.Eval(n.Expr)
	if IsError(val) {
		return val
	}
	e.parser.Env[n.Identifier.Literal] = val
	return val
}

func (e *Evaluator) evalConditionalExpression(n *parser.IfExpressionNode) objects.GoMixObject {
	condition := e.Eval(n.Condition)
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
	}
	return result
}

func (e *Evaluator) evalBinaryExpression(n *parser.BinaryExpressionNode) objects.GoMixObject {
	left := e.Eval(n.Left)
	right := e.Eval(n.Right)

	err := CreateError("Operator (%s) not implemented for (%s) and (%s)", n.Operation.Literal, left.GetType(), right.GetType())

	if left.GetType() != objects.IntegerType && left.GetType() != objects.FloatType {
		// panic("not implemented")
		return err
	}
	if right.GetType() != objects.IntegerType && right.GetType() != objects.FloatType {
		// panic("not implemented")
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

	err := CreateError("Operator (%s) not implemented for (%s)", n.Operation.Literal, right.GetType())

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
