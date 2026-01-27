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

func TestParser_Parse_ArithmeticExpression(t *testing.T) {
	src := `1+2*3-4`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&NumberLiteralExpressionNode{Value: 1},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 2},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 3},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&NumberLiteralExpressionNode{Value: 4},
		},
		Ptr: 0,
		T:   t,
	}

	// check for correctness
	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)

	assert.Equal(t, lexer.MINUS_OP, exp.Operation.Type)
	assert.Equal(t, "1+2*3-4", exp.Literal())
	assert.Equal(t, 3, exp.Value)
}

func TestParser_Parse_ArithmeticExpression_Complex1(t *testing.T) {
	src := `1+2*3-4/2`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&NumberLiteralExpressionNode{Value: 1},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 2},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 3},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&NumberLiteralExpressionNode{Value: 4},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "/"}},
			&NumberLiteralExpressionNode{Value: 2},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp := root.Statements[0].(*BinaryExpressionNode)
	assert.Equal(t, lexer.MINUS_OP, exp.Operation.Type)
	assert.Equal(t, 5, exp.Value)
}

func TestParser_Parse_ArithmeticExpression_Complex2(t *testing.T) {
	src := `20-5-5`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&NumberLiteralExpressionNode{Value: 20},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&NumberLiteralExpressionNode{Value: 5},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&NumberLiteralExpressionNode{Value: 5},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp := root.Statements[0].(*BinaryExpressionNode)
	assert.Equal(t, 10, exp.Value)
}

func TestParser_Parse_ParenthesizedExpression(t *testing.T) {
	src := `(10)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&NumberLiteralExpressionNode{Value: 10},
			&ParenthesizedExpressionNode{Expr: &NumberLiteralExpressionNode{Value: 10}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp, ok := root.Statements[0].(*ParenthesizedExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "(10)", exp.Literal())

}

func TestParser_Parse_ParenthesizedExpression_Complex(t *testing.T) {
	src := `(10-5)+5*1`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&NumberLiteralExpressionNode{Value: 10},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&NumberLiteralExpressionNode{Value: 5},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 5},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 1},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "(10-5)+5*1", exp.Literal())
	assert.Equal(t, 10, exp.Value)

}

func TestParser_Parse_ParenthesizedExpressionComplex(t *testing.T) {
	src := `((10 - 5)+5)*1`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&NumberLiteralExpressionNode{Value: 10},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&NumberLiteralExpressionNode{Value: 5},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 5},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 1},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "((10-5)+5)*1", exp.Literal())
	assert.Equal(t, 10, exp.Value)

}

func TestParser_ParseDeclarativeStatement(t *testing.T) {
	src := `var a = 1`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 1},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 1", exp.Literal())
	assert.Equal(t, 1, exp.Value)

}

func TestParser_ParseDeclarativeStatement_Complex(t *testing.T) {
	src := `var a = 1 + 2 * 3`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 1},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 2},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 3},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 1+2*3", exp.Literal())
	assert.Equal(t, 7, exp.Value)

}

func TestParser_ParseDeclarativeStatement_Complex2(t *testing.T) {
	src := `var a = (1 + 2) * 3`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 1},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 2},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 3},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = (1+2)*3", exp.Literal())
	assert.Equal(t, 9, exp.Value)

}

func TestParser_ParseDeclarativeStatement_Identifier(t *testing.T) {
	src := `var a=1
	var b = a + 10`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 1},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "b"},
			},
			&IdentifierExpressionNode{Name: "a", Value: 1},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 10},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 2, len(root.Statements))

	// check first statement: var a = 1
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 1", stmt1.Literal())
	assert.Equal(t, 1, stmt1.Value)

	// check second statement: var b = a + 10
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = a+10", stmt2.Literal())
	assert.Equal(t, 11, stmt2.Value)

	assert.Equal(t, "var a = 1;var b = a+10;", root.Literal())
}

func TestParser_ParseDeclarativeStatement_Identifier_With_ParenthesizedExpression(t *testing.T) {
	src := `var a=11
	var b = (a + 10 * 2)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 11},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "b"},
			},
			&IdentifierExpressionNode{Name: "a", Value: 11},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 10},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 2},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 2, len(root.Statements))

	// check first statement: var a = 11
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 11", stmt1.Literal())
	assert.Equal(t, 11, stmt1.Value)

	// check second statement: var b = (a + 10 * 2)
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = (a+10*2)", stmt2.Literal())
	assert.Equal(t, 31, stmt2.Value)

	assert.Equal(t, "var a = 11;var b = (a+10*2);", root.Literal())
}

func TestParser_ParseDeclarativeStatement_Identifier_With_ParenthesizedExpressionAndComma(t *testing.T) {
	src := `var a=11;var b = (a + 10 * 2);var c = (b + 10 * 3)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 11},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "b"},
			},
			&IdentifierExpressionNode{Name: "a", Value: 11},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 10},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 2},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "c"},
			},
			&IdentifierExpressionNode{Name: "b", Value: 31},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&NumberLiteralExpressionNode{Value: 10},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&NumberLiteralExpressionNode{Value: 3},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 3, len(root.Statements))

	// check first statement: var a = 11
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 11", stmt1.Literal())
	assert.Equal(t, 11, stmt1.Value)

	// check second statement: var b = (a + 10 * 2)
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = (a+10*2)", stmt2.Literal())
	assert.Equal(t, 31, stmt2.Value)

	// check third statement: var c = (b + 10 * 3	)
	stmt3, ok := root.Statements[2].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var c = (b+10*3)", stmt3.Literal())
	assert.Equal(t, 61, stmt3.Value)

	assert.Equal(t, "var a = 11;var b = (a+10*2);var c = (b+10*3);", root.Literal())
}

func TestParser_ParseDeclarativeStatement_Identifier_With_ReturnStatement(t *testing.T) {
	src := `var a = 1;return a`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 1},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr:        &IdentifierExpressionNode{Name: "a", Value: 1},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 2, len(root.Statements))

	// check first statement: var a = 1
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 1", stmt1.Literal())
	assert.Equal(t, 1, stmt1.Value)

	// check second statement: return a
	stmt2, ok := root.Statements[1].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return a", stmt2.Literal())
	assert.Equal(t, 1, stmt2.Value)

	assert.Equal(t, "var a = 1;return a;", root.Literal())
}

func TestParser_ParseDeclarativeStatement_Identifier_With_ReturnStatement_With_ParenthesizedExpression(t *testing.T) {
	src := `var a = 1;return (a + 10 * 2)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: lexer.Token{Literal: "a"},
			},
			&NumberLiteralExpressionNode{Value: 1},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr:        &ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 2, len(root.Statements))

	// check first statement: var a = 1
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 1", stmt1.Literal())
	assert.Equal(t, 1, stmt1.Value)

	// check second statement: return (a + 10 * 2)
	stmt2, ok := root.Statements[1].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return (a+10*2)", stmt2.Literal())
	assert.Equal(t, 21, stmt2.Value)

	assert.Equal(t, "var a = 1;return (a+10*2);", root.Literal())
}
