package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/akashmaji946/go-mix/lexer"
)

func TestParser_Parse_OneNumberExpression(t *testing.T) {

	src := `12`
	par := NewParser(src)
	root := par.Parse()
	// root should not be nil
	assert.NotNil(t, root)
	// optional: print the root

	// must: root has 1 statement
	assert.Equal(t, 1, len(root.Statements))

	exp, can := root.Statements[0].(*NumberLiteralExpressionNode)
	assert.True(t, can)
	assert.Equal(t, "12", exp.Literal())
	assert.Equal(t, 12, exp.Value)
}

func TestParser_Parse_AddExpression(t *testing.T) {

	src := `12 + 13`
	par := NewParser(src)
	root := par.Parse()
	// root should not be nil
	assert.NotNil(t, root)
	// optional: print the root

	// must: root has 1 statement
	assert.Equal(t, 1, len(root.Statements))

	exp, can := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, can)
	left, can := exp.Left.(*NumberLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*NumberLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "12", left.Literal())
	assert.Equal(t, 12, left.Value)
	assert.Equal(t, "13", right.Literal())
	assert.Equal(t, 13, right.Value)
	assert.Equal(t, "12+13", exp.Literal())
	assert.Equal(t, 25, exp.Value)
}

func TestParser_Parse_MulExpression(t *testing.T) {

	src := `12 * 13`
	par := NewParser(src)
	root := par.Parse()
	// root should not be nil
	assert.NotNil(t, root)
	// optional: print the root

	// must: root has 1 statement
	assert.Equal(t, 1, len(root.Statements))

	exp, can := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, can)
	left, can := exp.Left.(*NumberLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*NumberLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "12", left.Literal())
	assert.Equal(t, 12, left.Value)
	assert.Equal(t, "13", right.Literal())
	assert.Equal(t, 13, right.Value)
	assert.Equal(t, "12*13", exp.Literal())
	assert.Equal(t, 156, exp.Value)
}

func TestParser_Parse_DivExpression(t *testing.T) {

	src := `26 / 13`
	par := NewParser(src)
	root := par.Parse()
	// root should not be nil
	assert.NotNil(t, root)
	// optional: print the root

	// must: root has 1 statement
	assert.Equal(t, 1, len(root.Statements))

	exp, can := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, can)
	left, can := exp.Left.(*NumberLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*NumberLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "26", left.Literal())
	assert.Equal(t, 26, left.Value)
	assert.Equal(t, "13", right.Literal())
	assert.Equal(t, 13, right.Value)
	assert.Equal(t, "26/13", exp.Literal())
	assert.Equal(t, 2, exp.Value)
}

func TestParser_Parse_SubExpression(t *testing.T) {

	src := `28 - 13 * 2`
	par := NewParser(src)
	root := par.Parse()
	// root should not be nil
	assert.NotNil(t, root)
	// optional: print the root

	// must: root has 1 statement
	assert.Equal(t, 1, len(root.Statements))

	exp, can := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, can)
	left, can := exp.Left.(*NumberLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*BinaryExpressionNode)
	assert.True(t, can)
	rightLeft, can := right.Left.(*NumberLiteralExpressionNode)
	assert.True(t, can)
	rightRight, can := right.Right.(*NumberLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "28", left.Literal())
	assert.Equal(t, 28, left.Value)
	assert.Equal(t, "13", rightLeft.Literal())
	assert.Equal(t, 13, rightLeft.Value)
	assert.Equal(t, "2", rightRight.Literal())
	assert.Equal(t, 2, rightRight.Value)
	assert.Equal(t, "13*2", right.Literal())
	assert.Equal(t, 26, right.Value)
	assert.Equal(t, "28-13*2", exp.Literal())
	assert.Equal(t, 2, exp.Value)
}

func TestParser_Parse_FullyExpandedExpression(t *testing.T) {
	src := `26 + 13 * 2 - 12 / 2 - 6 + 6 - 4 * 2 + 100 - 9`
	par := NewParser(src)
	root := par.Parse()

	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// ((((((((26 + (13*2)) - (12/2)) - 6) + 6) - (4*2)) + 100) - 9)

	// level 1: - 9
	exp1, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp1.Operation.Type)

	right9, ok := exp1.Right.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 9, right9.Value)

	// level 2: + 100
	exp2, ok := exp1.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, exp2.Operation.Type)

	right100, ok := exp2.Right.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 100, right100.Value)

	// level 3: - (4*2)
	exp3, ok := exp2.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp3.Operation.Type)

	mul4x2, ok := exp3.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MUL_OP, mul4x2.Operation.Type)

	n4, ok := mul4x2.Left.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 4, n4.Value)

	n2, ok := mul4x2.Right.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 2, n2.Value)

	// level 4: + 6
	exp4, ok := exp3.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, exp4.Operation.Type)

	right6b, ok := exp4.Right.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 6, right6b.Value)

	// level 5: - 6
	exp5, ok := exp4.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp5.Operation.Type)

	right6a, ok := exp5.Right.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 6, right6a.Value)

	// level 6: - (12/2)
	exp6, ok := exp5.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp6.Operation.Type)

	div12by2, ok := exp6.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.DIV_OP, div12by2.Operation.Type)

	n12, ok := div12by2.Left.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 12, n12.Value)

	n2b, ok := div12by2.Right.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 2, n2b.Value)

	// level 7: + (13*2)
	exp7, ok := exp6.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, exp7.Operation.Type)

	mul13x2, ok := exp7.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MUL_OP, mul13x2.Operation.Type)

	n13, ok := mul13x2.Left.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 13, n13.Value)

	n2c, ok := mul13x2.Right.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 2, n2c.Value)

	// level 8: 26
	n26, ok := exp7.Left.(*NumberLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 26, n26.Value)

	// final sanity checks
	assert.Equal(t, "26+13*2-12/2-6+6-4*2+100-9", exp1.Literal())
	assert.Equal(t, 129, exp1.Value)
}

func TestParser_Parse_UnaryExpression(t *testing.T) {
	src := `!true`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*UnaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.NOT_OP, exp.Operation.Type)
	assert.Equal(t, "!true", exp.Literal())
	assert.Equal(t, 0, exp.Value)
}

func TestParser_Parse_UnaryExpressionSimple(t *testing.T) {
	src := `!!!!!!false`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*UnaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.NOT_OP, exp.Operation.Type)
	assert.Equal(t, "!!!!!!false", exp.Literal())
	assert.Equal(t, 0, exp.Value)
}
