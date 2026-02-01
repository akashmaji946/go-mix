package eval

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/parser"
)

type Evaluator struct {
	parser *parser.Parser
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (e *Evaluator) Eval(n parser.Node) GoMixObject {
	switch n := n.(type) {
	case *parser.RootNode:
		return e.evalStatements(n.Statements)
	case *parser.BooleanLiteralExpressionNode:
		return &Boolean{Value: n.Value}
	case *parser.IntegerLiteralExpressionNode:
		return &Integer{Value: int64(n.Value)}
	case *parser.StringLiteralExpressionNode:
		return &String{Value: n.Value}
	case *parser.FloatLiteralExpressionNode:
		return &Float{Value: n.Value}
	case *parser.NilLiteralExpressionNode:
		return &Nil{}
	case *parser.BinaryExpressionNode:
		return e.evalBinaryExpression(n)
	case *parser.UnaryExpressionNode:
		return e.evalUnaryExpression(n)
	case *parser.BooleanExpressionNode:
		return e.evalBooleanExpression(n)
	case *parser.ParenthesizedExpressionNode:
		return e.Eval(n.Expr)
	default:
		panic("not implemented")
	}
}

func (e *Evaluator) evalStatements(stmts []parser.StatementNode) GoMixObject {
	var result GoMixObject
	for _, stmt := range stmts {
		result = e.Eval(stmt)
	}
	return result
}

func (e *Evaluator) evalBinaryExpression(n *parser.BinaryExpressionNode) GoMixObject {
	left := e.Eval(n.Left)
	right := e.Eval(n.Right)

	if left.GetType() != IntegerType && left.GetType() != FloatType {
		panic("not implemented")
	}
	if right.GetType() != IntegerType && right.GetType() != FloatType {
		panic("not implemented")
	}

	leftType := left.GetType()
	rightType := right.GetType()

	switch n.Operation.Type {
	case lexer.PLUS_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value + right.(*Integer).Value}
		}
		return &Float{Value: toFloat64(left) + toFloat64(right)}
	case lexer.MINUS_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value - right.(*Integer).Value}
		}
		return &Float{Value: toFloat64(left) - toFloat64(right)}
	case lexer.MUL_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value * right.(*Integer).Value}
		}
		return &Float{Value: toFloat64(left) * toFloat64(right)}
	case lexer.DIV_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value / right.(*Integer).Value}
		}
		return &Float{Value: toFloat64(left) / toFloat64(right)}
	case lexer.MOD_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value % right.(*Integer).Value}
		}
		panic("not implemented")
	case lexer.BIT_AND_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value & right.(*Integer).Value}
		}
		panic("not implemented")
	case lexer.BIT_OR_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value | right.(*Integer).Value}
		}
		panic("not implemented")
	case lexer.BIT_XOR_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value ^ right.(*Integer).Value}
		}
		panic("not implemented")

	case lexer.BIT_LEFT_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value << right.(*Integer).Value}
		}
		panic("not implemented")
	case lexer.BIT_RIGHT_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Integer{Value: left.(*Integer).Value >> right.(*Integer).Value}
		}
		panic("not implemented")
	}
	return &Nil{}
}

func toFloat64(obj GoMixObject) float64 {
	if obj.GetType() == IntegerType {
		return float64(obj.(*Integer).Value)
	}
	return obj.(*Float).Value
}

func (e *Evaluator) evalUnaryExpression(n *parser.UnaryExpressionNode) GoMixObject {
	right := e.Eval(n.Right)

	switch n.Operation.Type {
	case lexer.NOT_OP:
		if right.GetType() != BooleanType {
			panic("not implemented")
		}
		return &Boolean{Value: !right.(*Boolean).Value}
	case lexer.BIT_NOT_OP:
		if right.GetType() == IntegerType {
			return &Integer{Value: ^right.(*Integer).Value}
		}
		panic("not implemented")
	case lexer.MINUS_OP:
		if right.GetType() == IntegerType {
			return &Integer{Value: -right.(*Integer).Value}
		} else if right.GetType() == FloatType {
			return &Float{Value: -right.(*Float).Value}
		}
		panic("not implemented")
	}
	return &Nil{}
}

func (e *Evaluator) evalBooleanExpression(n *parser.BooleanExpressionNode) GoMixObject {
	left := e.Eval(n.Left)
	right := e.Eval(n.Right)

	leftType := left.GetType()
	rightType := right.GetType()

	switch n.Operation.Type {
	case lexer.EQ_OP:
		return &Boolean{Value: left.ToString() == right.ToString()}
	case lexer.NE_OP:
		return &Boolean{Value: left.ToString() != right.ToString()}
	case lexer.GT_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Boolean{Value: left.(*Integer).Value > right.(*Integer).Value}
		}
		return &Boolean{Value: toFloat64(left) > toFloat64(right)}
	case lexer.LT_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Boolean{Value: left.(*Integer).Value < right.(*Integer).Value}
		}
		return &Boolean{Value: toFloat64(left) < toFloat64(right)}
	case lexer.GE_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Boolean{Value: left.(*Integer).Value >= right.(*Integer).Value}
		}
		return &Boolean{Value: toFloat64(left) >= toFloat64(right)}
	case lexer.LE_OP:
		if leftType == IntegerType && rightType == IntegerType {
			return &Boolean{Value: left.(*Integer).Value <= right.(*Integer).Value}
		}
		return &Boolean{Value: toFloat64(left) <= toFloat64(right)}
	case lexer.AND_OP:
		return &Boolean{Value: left.(*Boolean).Value && right.(*Boolean).Value}
	case lexer.OR_OP:
		return &Boolean{Value: left.(*Boolean).Value || right.(*Boolean).Value}
	}
	return &Nil{}
}
