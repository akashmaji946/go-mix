package eval

import (
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
	default:
		panic("not implemented")
	}
}

func (e *Evaluator) evalStatements(stmts []parser.StatementNode) objects.GoMixObject {
	var result objects.GoMixObject
	for _, stmt := range stmts {
		result = e.Eval(stmt)
	}
	return result
}

func (e *Evaluator) evalBinaryExpression(n *parser.BinaryExpressionNode) objects.GoMixObject {
	left := e.Eval(n.Left)
	right := e.Eval(n.Right)

	if left.GetType() != objects.IntegerType && left.GetType() != objects.FloatType {
		panic("not implemented")
	}
	if right.GetType() != objects.IntegerType && right.GetType() != objects.FloatType {
		panic("not implemented")
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
		panic("not implemented")
	case lexer.BIT_AND_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value & right.(*objects.Integer).Value}
		}
		panic("not implemented")
	case lexer.BIT_OR_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value | right.(*objects.Integer).Value}
		}
		panic("not implemented")
	case lexer.BIT_XOR_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value ^ right.(*objects.Integer).Value}
		}
		panic("not implemented")

	case lexer.BIT_LEFT_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value << right.(*objects.Integer).Value}
		}
		panic("not implemented")
	case lexer.BIT_RIGHT_OP:
		if leftType == objects.IntegerType && rightType == objects.IntegerType {
			return &objects.Integer{Value: left.(*objects.Integer).Value >> right.(*objects.Integer).Value}
		}
		panic("not implemented")
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

	switch n.Operation.Type {
	case lexer.NOT_OP:
		if right.GetType() != objects.BooleanType {
			panic("not implemented")
		}
		return &objects.Boolean{Value: !right.(*objects.Boolean).Value}
	case lexer.BIT_NOT_OP:
		if right.GetType() == objects.IntegerType {
			return &objects.Integer{Value: ^right.(*objects.Integer).Value}
		}
		panic("not implemented")
	case lexer.MINUS_OP:
		if right.GetType() == objects.IntegerType {
			return &objects.Integer{Value: -right.(*objects.Integer).Value}
		} else if right.GetType() == objects.FloatType {
			return &objects.Float{Value: -right.(*objects.Float).Value}
		}
		panic("not implemented")
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
