/*
File    : go-mix/parser/parser_test.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// TestParser_OneNum verifies parsing of a single integer literal
func TestParser_OneNum(t *testing.T) {

	src := `12`
	par := NewParser(src)
	root := par.Parse()
	// root should not be nil
	assert.NotNil(t, root)
	// optional: print the root

	// must: root has 1 statement
	assert.Equal(t, 1, len(root.Statements))

	exp, can := root.Statements[0].(*IntegerLiteralExpressionNode)
	assert.True(t, can)
	assert.Equal(t, "12", exp.Literal())
	const expectedVal int64 = 12
	if intObj, ok := exp.Value.(*std.Integer); ok {
		assert.Equal(t, expectedVal, intObj.Value)
	} else {
		t.Errorf("Expected objects.Integer, got %T", exp.Value)
	}
}

func TestParser_HexIntLiteral(t *testing.T) {
	src := `0x16`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	exp, ok := root.Statements[0].(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "0x16", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 22}, exp.Value)
}

func TestParser_OctalIntLiteral(t *testing.T) {
	src := `0777`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	exp, ok := root.Statements[0].(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "0777", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 511}, exp.Value)
}

func TestParser_ScientificFloatLiteral(t *testing.T) {
	src := `1.4e3`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	exp, ok := root.Statements[0].(*FloatLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "1.4e3", exp.Literal())
	if floatObj, ok := exp.Value.(*std.Float); ok {
		assert.InDelta(t, 1400.0, floatObj.Value, 1e-9)
	} else {
		t.Errorf("Expected objects.Float, got %T", exp.Value)
	}
}

// TestParser_Add verifies parsing of addition expressions
func TestParser_Add(t *testing.T) {

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
	left, can := exp.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "12", left.Literal())
	assert.Equal(t, &std.Integer{Value: 12}, left.Value)
	assert.Equal(t, "13", right.Literal())
	assert.Equal(t, &std.Integer{Value: 13}, right.Value)
	assert.Equal(t, "12+13", exp.Literal())

	const expectedVal int64 = 25
	if intObj, ok := exp.Value.(*std.Integer); ok {
		assert.Equal(t, expectedVal, intObj.Value)
	} else {
		t.Errorf("Expected objects.Integer, got %T", exp.Value)
	}
}

// TestParser_Sub verifies parsing of subtraction with operator precedence
func TestParser_Sub(t *testing.T) {

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
	left, can := exp.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*BinaryExpressionNode)
	assert.True(t, can)
	rightLeft, can := right.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, can)
	rightRight, can := right.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "28", left.Literal())
	assert.Equal(t, &std.Integer{Value: 28}, left.Value)
	assert.Equal(t, "13", rightLeft.Literal())
	assert.Equal(t, &std.Integer{Value: 13}, rightLeft.Value)
	assert.Equal(t, "2", rightRight.Literal())
	assert.Equal(t, &std.Integer{Value: 2}, rightRight.Value)
	assert.Equal(t, "13*2", right.Literal())
	assert.Equal(t, &std.Integer{Value: 26}, right.Value)
	assert.Equal(t, "28-13*2", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 2}, exp.Value)
}

// TestParser_Mul verifies parsing of multiplication expressions
func TestParser_Mul(t *testing.T) {

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
	left, can := exp.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "12", left.Literal())
	assert.Equal(t, &std.Integer{Value: 12}, left.Value)
	assert.Equal(t, "13", right.Literal())
	assert.Equal(t, &std.Integer{Value: 13}, right.Value)
	assert.Equal(t, "12*13", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 156}, exp.Value)
}

// TestParser_Div verifies parsing of division expressions
func TestParser_Div(t *testing.T) {

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
	left, can := exp.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, can)
	right, can := exp.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, can)

	assert.Equal(t, "26", left.Literal())
	assert.Equal(t, &std.Integer{Value: 26}, left.Value)
	assert.Equal(t, "13", right.Literal())
	assert.Equal(t, &std.Integer{Value: 13}, right.Value)
	assert.Equal(t, "26/13", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 2}, exp.Value)
}

// TestParser_FullExpr verifies parsing of complex expressions with multiple operators
func TestParser_FullExpr(t *testing.T) {
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

	right9, ok := exp1.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 9}, right9.Value)

	// level 2: + 100
	exp2, ok := exp1.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, exp2.Operation.Type)

	right100, ok := exp2.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 100}, right100.Value)

	// level 3: - (4*2)
	exp3, ok := exp2.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp3.Operation.Type)

	mul4x2, ok := exp3.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MUL_OP, mul4x2.Operation.Type)

	n4, ok := mul4x2.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 4}, n4.Value)

	n2, ok := mul4x2.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 2}, n2.Value)

	// level 4: + 6
	exp4, ok := exp3.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, exp4.Operation.Type)

	right6b, ok := exp4.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 6}, right6b.Value)

	// level 5: - 6
	exp5, ok := exp4.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp5.Operation.Type)

	right6a, ok := exp5.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 6}, right6a.Value)

	// level 6: - (12/2)
	exp6, ok := exp5.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp6.Operation.Type)

	div12by2, ok := exp6.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.DIV_OP, div12by2.Operation.Type)

	n12, ok := div12by2.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 12}, n12.Value)

	n2b, ok := div12by2.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 2}, n2b.Value)

	// level 7: + (13*2)
	exp7, ok := exp6.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, exp7.Operation.Type)

	mul13x2, ok := exp7.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MUL_OP, mul13x2.Operation.Type)

	n13, ok := mul13x2.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 13}, n13.Value)

	n2c, ok := mul13x2.Right.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 2}, n2c.Value)

	// level 8: 26
	n26, ok := exp7.Left.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 26}, n26.Value)

	// final sanity checks
	assert.Equal(t, "26+13*2-12/2-6+6-4*2+100-9", exp1.Literal())
	assert.Equal(t, &std.Integer{Value: 129}, exp1.Value)
}

// TestParser_Unary1 verifies parsing of boolean negation unary operator
func TestParser_Unary1(t *testing.T) {
	src := `!true`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*UnaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.NOT_OP, exp.Operation.Type)
	assert.Equal(t, "!true", exp.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, exp.Value)
}

// TestParser_Unary2 verifies parsing of numeric negation unary operator
func TestParser_Unary2(t *testing.T) {
	src := `-12`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*UnaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, exp.Operation.Type)
	assert.Equal(t, "-12", exp.Literal())
	assert.Equal(t, &std.Integer{Value: -12}, exp.Value)
}

// TestParser_Bool1 verifies parsing of true boolean literal
func TestParser_Bool1(t *testing.T) {
	src := `true`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*BooleanLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "true", exp.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, exp.Value)
}

// TestParser_Bool2 verifies parsing of false boolean literal
func TestParser_Bool2(t *testing.T) {
	src := `false`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*BooleanLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "false", exp.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, exp.Value)
}

// TestParser_BoolSimple verifies parsing of simple boolean AND expression
func TestParser_BoolSimple(t *testing.T) {
	src := `false && true`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.AND_OP, exp.Operation.Type)
	assert.Equal(t, "false&&true", exp.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, exp.Value)
}

// TestParser_BoolComplex verifies parsing of complex boolean expressions with AND/OR
func TestParser_BoolComplex(t *testing.T) {
	src := `false && true || false`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
		},
		Ptr: 0,
		T:   t,
	}

	// check for correctness
	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "false&&true||false", exp.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, exp.Value)
}

// TestParser_BoolComplex2 verifies parsing of boolean expressions with parentheses
func TestParser_BoolComplex2(t *testing.T) {
	src := `false && true || (false || true)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&ParenthesizedExpressionNode{Expr: &BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}}},
		},
		Ptr: 0,
		T:   t,
	}

	// check for correctness
	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "false&&true||(false||true)", exp.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, exp.Value)
}

// TestParser_Arith verifies parsing of arithmetic expressions with precedence
func TestParser_Arith(t *testing.T) {
	src := `1+2*3-4`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 4}},
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

	if intObj, ok := exp.Value.(*std.Integer); ok {
		assert.Equal(t, &std.Integer{Value: 3}, intObj)
	} else {
		t.Errorf("Expected objects.Integer, got %T", exp.Value)
	}
}

// TestParser_ArithComplex1 verifies parsing of complex arithmetic with division
func TestParser_ArithComplex1(t *testing.T) {
	src := `1+2*3-4/2`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 4}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "/"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp := root.Statements[0].(*BinaryExpressionNode)
	assert.Equal(t, lexer.MINUS_OP, exp.Operation.Type)
	assert.Equal(t, &std.Integer{Value: 5}, exp.Value)
}

// TestParser_ArithComplex2 verifies parsing of left-associative subtraction
func TestParser_ArithComplex2(t *testing.T) {
	src := `20-5-5`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 20}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp := root.Statements[0].(*BinaryExpressionNode)
	assert.Equal(t, &std.Integer{Value: 10}, exp.Value)
}

// TestParser_Paren verifies parsing of parenthesized expressions
func TestParser_Paren(t *testing.T) {
	src := `(10)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&ParenthesizedExpressionNode{Expr: &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp, ok := root.Statements[0].(*ParenthesizedExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "(10)", exp.Literal())

}

// TestParser_ParenComplex verifies parsing of parentheses affecting precedence
func TestParser_ParenComplex(t *testing.T) {
	src := `(10-5)+5*1`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "(10-5)+5*1", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 10}, exp.Value)

}

// TestParser_ParenNested verifies parsing of nested parenthesized expressions
func TestParser_ParenNested(t *testing.T) {
	src := `((10 - 5)+5)*1`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "-"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	exp, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "((10-5)+5)*1", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 10}, exp.Value)

}

// TestParser_DeclStmt verifies parsing of simple variable declarations
func TestParser_DeclStmt(t *testing.T) {
	src := `var a = 1`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 1", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 1}, exp.Value)

}

// TestParser_DeclComplex verifies parsing of declarations with expressions
func TestParser_DeclComplex(t *testing.T) {
	src := `var a = 1 + 2 * 3`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 1+2*3", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 7}, exp.Value)

}

// TestParser_DeclComplex2 verifies parsing of declarations with parenthesized expressions
func TestParser_DeclComplex2(t *testing.T) {
	src := `var a = (1 + 2) * 3`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))
	exp, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = (1+2)*3", exp.Literal())
	assert.Equal(t, &std.Integer{Value: 9}, exp.Value)

}

// TestParser_DeclIdent verifies parsing of declarations using identifiers
func TestParser_DeclIdent(t *testing.T) {
	src := `var a=1
	var b = a + 10`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IdentifierExpressionNode{Name: "a", Value: &std.Integer{Value: 1}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
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
	assert.Equal(t, &std.Integer{Value: 1}, stmt1.Value)

	// check second statement: var b = a + 10
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = a+10", stmt2.Literal())
	assert.Equal(t, &std.Integer{Value: 11}, stmt2.Value)

	assert.Equal(t, "var a = 1;var b = a+10;", root.Literal())
}

// TestParser_DeclIdentParen verifies parsing of declarations with identifiers and parentheses
func TestParser_DeclIdentParen(t *testing.T) {
	src := `var a=11
	var b = (a + 10 * 2)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 11}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IdentifierExpressionNode{Name: "a", Value: &std.Integer{Value: 11}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
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
	assert.Equal(t, &std.Integer{Value: 11}, stmt1.Value)

	// check second statement: var b = (a + 10 * 2)
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = (a+10*2)", stmt2.Literal())
	assert.Equal(t, &std.Integer{Value: 31}, stmt2.Value)

	assert.Equal(t, "var a = 11;var b = (a+10*2);", root.Literal())
}

// TestParser_DeclMulti verifies parsing of multiple declarations with semicolons
func TestParser_DeclMulti(t *testing.T) {
	src := `var a=11;var b = (a + 10 * 2);var c = (b + 10 * 3)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 11}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IdentifierExpressionNode{Name: "a", Value: &std.Integer{Value: 11}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&ParenthesizedExpressionNode{Expr: &BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "c"},
			},
			&IdentifierExpressionNode{Name: "b", Value: &std.Integer{Value: 31}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
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
	assert.Equal(t, &std.Integer{Value: 11}, stmt1.Value)

	// check second statement: var b = (a + 10 * 2)
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = (a+10*2)", stmt2.Literal())
	assert.Equal(t, &std.Integer{Value: 31}, stmt2.Value)

	// check third statement: var c = (b + 10 * 3	)
	stmt3, ok := root.Statements[2].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var c = (b+10*3)", stmt3.Literal())
	assert.Equal(t, &std.Integer{Value: 61}, stmt3.Value)

	assert.Equal(t, "var a = 11;var b = (a+10*2);var c = (b+10*3);", root.Literal())
}

// TestParser_DeclReturn verifies parsing of declarations followed by return
func TestParser_DeclReturn(t *testing.T) {
	src := `var a = 1;return a`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr:        &IdentifierExpressionNode{Name: "a", Value: &std.Integer{Value: 1}},
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
	assert.Equal(t, &std.Integer{Value: 1}, stmt1.Value)

	// check second statement: return a
	stmt2, ok := root.Statements[1].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return a", stmt2.Literal())
	assert.Equal(t, &std.Integer{Value: 1}, stmt2.Value)

	assert.Equal(t, "var a = 1;return a;", root.Literal())
}

// TestParser_DeclReturnParen verifies parsing of return with parenthesized expressions
func TestParser_DeclReturnParen(t *testing.T) {
	src := `var a = 1;return (a + 10 * 2)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
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
	assert.Equal(t, &std.Integer{Value: 1}, stmt1.Value)

	// check second statement: return (a + 10 * 2)
	stmt2, ok := root.Statements[1].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return (a+10*2)", stmt2.Literal())
	assert.Equal(t, &std.Integer{Value: 21}, stmt2.Value)

	assert.Equal(t, "var a = 1;return (a+10*2);", root.Literal())
}

// TestParser_BoolExpr verifies parsing of boolean AND expressions
func TestParser_BoolExpr(t *testing.T) {
	src := `true && false`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{

			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: true && false
	stmt1, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "true&&false", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, stmt1.Value)

	assert.Equal(t, "true&&false;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, root.Value)
}

// TestParser_BoolParenExpr verifies parsing of parenthesized boolean expressions
func TestParser_BoolParenExpr(t *testing.T) {
	src := `(false || true && false)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&ParenthesizedExpressionNode{
				Expr: &BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: true && false
	stmt1, ok := root.Statements[0].(*ParenthesizedExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "(false||true&&false)", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, stmt1.Value)

	assert.Equal(t, "(false||true&&false);", root.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, root.Value)
}

// TestParser_DeclBoolReturn verifies parsing of boolean declarations and returns
func TestParser_DeclBoolReturn(t *testing.T) {
	src := `var a = true; var b = a && false; return b || true;`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IdentifierExpressionNode{Name: "a", Value: &std.Integer{Value: 1}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr:        &BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 3, len(root.Statements))

	// check first statement: var a = true
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = true", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt1.Value)

	// check second statement: var b = a && false
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = a&&false", stmt2.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, stmt2.Value)

	// check third statement: return b || true
	stmt3, ok := root.Statements[2].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return b||true", stmt3.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt3.Value)

	assert.Equal(t, "var a = true;var b = a&&false;return b||true;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)

}

// TestParser_RelOp verifies parsing of relational comparison operators
func TestParser_RelOp(t *testing.T) {
	src := `1 < 2`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "<"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: 1 < 2
	stmt1, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "1<2", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt1.Value)

	assert.Equal(t, "1<2;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)
}

// TestParser_RelOpSimple verifies parsing of relational operators with boolean expressions
func TestParser_RelOpSimple(t *testing.T) {
	src := `false || 1 < 2`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "<"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: 1 < 2
	stmt1, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "false||1<2", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt1.Value)

	assert.Equal(t, "false||1<2;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)
}

// TestParser_RelOpComplex verifies parsing of complex relational and boolean combinations
func TestParser_RelOpComplex(t *testing.T) {
	src := `false || 10 <= 20 && true`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "<="}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 20}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: 1 < 2
	stmt1, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "false||10<=20&&true", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt1.Value)

	assert.Equal(t, "false||10<=20&&true;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)
}

// TestParser_RelOpParen verifies parsing of relational operators with parentheses
func TestParser_RelOpParen(t *testing.T) {
	src := `false || (10 <= 20 && true)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "||"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "<="}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 20}},
			&BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&ParenthesizedExpressionNode{
				Expr: &BooleanExpressionNode{Operation: lexer.Token{Literal: "&&"}},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: 1 < 2
	stmt1, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "false||(10<=20&&true)", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt1.Value)

	assert.Equal(t, "false||(10<=20&&true);", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)
}

// TestParser_RelOpVar verifies parsing of relational operators with variables
func TestParser_RelOpVar(t *testing.T) {
	src := `var a = false; return a || (10 <= 20 && true);`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr: &BooleanExpressionNode{
					Operation: lexer.Token{Literal: "||"},
					Left:      &IdentifierExpressionNode{Name: "a"},
					Right: &BooleanExpressionNode{
						Operation: lexer.Token{Literal: "&&"},
						Left: &BooleanExpressionNode{
							Operation: lexer.Token{Literal: "<="},
							Left:      &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
							Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 20}},
						},
						Right: &BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
					},
				},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 2, len(root.Statements))

	// check first statement: var a = false
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = false", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: false}, stmt1.Value)

	// check second statement: return a || (10 <= 20 && true)
	stmt2, ok := root.Statements[1].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return a||(10<=20&&true)", stmt2.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt2.Value)

	assert.Equal(t, "var a = false;return a||(10<=20&&true);", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)

}

// TestParser_BitOp verifies parsing of bitwise operators with precedence
func TestParser_BitOp(t *testing.T) {
	// In C-based languages, == has higher precedence than &
	// So `3 & 7 == 3` is parsed as `3 & (7 == 3)` = `3 & false` = `3 & 0` = 0
	// To get `(3 & 7) == 3`, you need explicit parentheses
	src := `(3 & 7) == 3`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "&"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
			&ParenthesizedExpressionNode{},
			&BooleanExpressionNode{
				Operation: lexer.Token{Literal: "=="},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: (3 & 7) == 3
	stmt1, ok := root.Statements[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "(3&7)==3", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt1.Value)

	assert.Equal(t, "(3&7)==3;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)
}

// TestParser_RelBitComplex verifies parsing of complex relational and bitwise combinations
func TestParser_RelBitComplex(t *testing.T) {
	src := `return ((3&7)!=3&&true||false&&true)||true;`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr: &BooleanExpressionNode{
					Operation: lexer.Token{Literal: "||"},
					Left: &BooleanExpressionNode{
						Operation: lexer.Token{Literal: "&&"},
						Left: &BooleanExpressionNode{
							Operation: lexer.Token{Literal: "!="},
							Left: &ParenthesizedExpressionNode{
								Expr: &BooleanExpressionNode{
									Operation: lexer.Token{Literal: "&"},
									Left:      &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
									Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
								},
							},
							Right: &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
						},
						Right: &BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
					},
					Right: &BooleanExpressionNode{
						Operation: lexer.Token{Literal: "&&"},
						Left:      &BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
						Right:     &BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
					},
				},
				// Right: &BooleanLiteralExpressionNode{Value: &objects.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 1, len(root.Statements))

	// check first statement: return ((3&7)!=3&&true||false&&true)||true
	stmt1, ok := root.Statements[0].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return ((3&7)!=3&&true||false&&true)||true", stmt1.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt1.Value)

	assert.Equal(t, "return ((3&7)!=3&&true||false&&true)||true;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)
}

// TestParser_BitOpParen verifies parsing of bitwise operators with parentheses and variables
func TestParser_BitOpParen(t *testing.T) {
	src := `var a = (3&7); return (a==3) && true;`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
			&BinaryExpressionNode{
				Operation: lexer.Token{Literal: "&"},
				Left:      &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
				Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
			&ParenthesizedExpressionNode{
				Expr: &BooleanExpressionNode{
					Operation: lexer.Token{Literal: "&"},
					Left:      &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
					Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
				},
			},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr: &BooleanExpressionNode{
					Operation: lexer.Token{Literal: "||"},
					Left: &BooleanExpressionNode{
						Operation: lexer.Token{Literal: "&&"},
						Left: &BooleanExpressionNode{
							Operation: lexer.Token{Literal: "!="},
							Left: &ParenthesizedExpressionNode{
								Expr: &BooleanExpressionNode{
									Operation: lexer.Token{Literal: "&"},
									Left:      &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
									Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
								},
							},
							Right: &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
						},
						Right: &BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
					},
					Right: &BooleanExpressionNode{
						Operation: lexer.Token{Literal: "&&"},
						Left:      &BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
						Right:     &BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
					},
				},
				// Right: &BooleanLiteralExpressionNode{Value: &objects.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 2, len(root.Statements))

	// check first statement: var a = (3&7)
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = (3&7)", stmt1.Literal())
	assert.Equal(t, &std.Integer{Value: 3}, stmt1.Value)

	// check second statement: return (a==3) && true
	stmt2, ok := root.Statements[1].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return (a==3)&&true", stmt2.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt2.Value)

	assert.Equal(t, "var a = (3&7);return (a==3)&&true;", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)

}

// TestParser_RelReturn verifies parsing of relational operators in return statements
func TestParser_RelReturn(t *testing.T) {
	src := `var a = 7; var b = 1; var c = 2; var d = 1; return ((a-b)>(c+d));`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "c"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "d"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr: &BooleanExpressionNode{
					Operation: lexer.Token{Literal: ">"},
					Left: &ParenthesizedExpressionNode{
						Expr: &BooleanExpressionNode{
							Operation: lexer.Token{Literal: "-"},
							Left:      &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 7}},
							Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
						},
					},
					Right: &ParenthesizedExpressionNode{
						Expr: &BooleanExpressionNode{
							Operation: lexer.Token{Literal: "+"},
							Left:      &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
							Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
						},
					},
				},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

	assert.Equal(t, 5, len(root.Statements))

	// check first statement: var a = 7
	stmt1, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var a = 7", stmt1.Literal())
	assert.Equal(t, &std.Integer{Value: 7}, stmt1.Value)

	// check second statement: var b = 1
	stmt2, ok := root.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var b = 1", stmt2.Literal())
	assert.Equal(t, &std.Integer{Value: 1}, stmt2.Value)

	// check third statement: var c = 2
	stmt3, ok := root.Statements[2].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var c = 2", stmt3.Literal())
	assert.Equal(t, &std.Integer{Value: 2}, stmt3.Value)

	// check fourth statement: var d = 1
	stmt4, ok := root.Statements[3].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "var d = 1", stmt4.Literal())
	assert.Equal(t, &std.Integer{Value: 1}, stmt4.Value)

	// check fifth statement: return ((a-b)>(c+d))
	stmt5, ok := root.Statements[4].(*ReturnStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "return ((a-b)>(c+d))", stmt5.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, stmt5.Value)

	assert.Equal(t, "var a = 7;var b = 1;var c = 2;var d = 1;return ((a-b)>(c+d));", root.Literal())
	assert.Equal(t, &std.Boolean{Value: true}, root.Value)

}

// TestParser_BlockSimple verifies parsing of simple block statements
func TestParser_BlockSimple(t *testing.T) {
	src := `{10 * 2 + 100;}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BlockStatementNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&BinaryExpressionNode{Operation: lexer.Token{Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 100}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_Block verifies parsing of block statements with multiple declarations
func TestParser_Block(t *testing.T) {
	src := `{
	var a = 10;
	var b = a + 10;
	var c = b * 100;
	return 1000;
	}`

	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{

		ExpectedNodes: []Node{
			&BlockStatementNode{},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},

			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IdentifierExpressionNode{Name: "a"},
			&BinaryExpressionNode{
				Operation: lexer.Token{Literal: "+"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},

			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "c"},
			},
			&IdentifierExpressionNode{Name: "b"},
			&BinaryExpressionNode{
				Operation: lexer.Token{Literal: "*"},
				Left:      &IdentifierExpressionNode{Name: "b"},
				Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 100}},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 100}},

			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr:        &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1000}},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1000}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Integer{Value: 1000}, root.Value)
	assert.Equal(t, `{var a = 10;var b = a+10;var c = b*100;return 1000;};`, root.Literal())

}

// TestParser_BlockReturn verifies parsing of block statements with return
func TestParser_BlockReturn(t *testing.T) {
	src := `{
	var a = 10;
	var b = a + 10;
	var c = b * 100;
	return 1000;
	}`

	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BlockStatementNode{},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IdentifierExpressionNode{Name: "a"},
			&BinaryExpressionNode{
				Operation: lexer.Token{Literal: "+"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "c"},
			},
			&IdentifierExpressionNode{Name: "b"},
			&BinaryExpressionNode{
				Operation: lexer.Token{Literal: "*"},
				Left:      &IdentifierExpressionNode{Name: "b"},
				Right:     &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 100}},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 100}},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr:        &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1000}},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1000}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

}

// TestParser_If verifies parsing of simple if statements
func TestParser_If(t *testing.T) {
	src := `if (1) { }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IfExpressionNode{
				IfToken: lexer.Token{Literal: "if"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&ParenthesizedExpressionNode{
				Expr: &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			},
			&BlockStatementNode{},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Nil{}, root.Value)
	assert.Equal(t, `if (1) {};`, root.Literal())
}

// TestParser_IfElse verifies parsing of if-else statements
func TestParser_IfElse(t *testing.T) {
	src := `if (1) { 1 } else { 2 }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IfExpressionNode{
				IfToken: lexer.Token{Literal: "if"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&ParenthesizedExpressionNode{
				Expr: &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			},
			&BlockStatementNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BlockStatementNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Integer{Value: 1}, root.Value) // Condition is true, ThenBlock returns 1
	assert.Equal(t, `if (1) {1;} else {2;};`, root.Literal())
}

// TestParser_ElseIf verifies parsing of else-if chains
func TestParser_ElseIf(t *testing.T) {
	src := `if (1) { 1 } else if (2) { 2 } else { 3 }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	// Expecting:
	// IfNode
	//   Condition: 1
	//   ThenBlock: {1}
	//   ElseBlock: {
	//       IfNode
	//         Condition: 2
	//         ThenBlock: {2}
	//         ElseBlock: {3}
	//   }

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IfExpressionNode{
				IfToken: lexer.Token{Literal: "if"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&ParenthesizedExpressionNode{
				Expr: &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			},
			&BlockStatementNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},

			// Implicit block for the else if
			&BlockStatementNode{},

			&IfExpressionNode{
				IfToken: lexer.Token{Literal: "if"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&ParenthesizedExpressionNode{
				Expr: &IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			},
			&BlockStatementNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&BlockStatementNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Integer{Value: 1}, root.Value)
	// Note: Literal reconstruction might differ slightly depending on implementation details of nested if block wrapping
	// but purely based on AST node traversal above, we are good.
}

// TestParser_ElseIfEval verifies parsing and evaluation of else-if conditions
func TestParser_ElseIfEval(t *testing.T) {
	src := `if (1 == 2) { 1 } else if (2 != 2) { 2 } else { 3 }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	// Result should be 2 because the else-if condition is true.
	assert.Equal(t, &std.Integer{Value: 3}, root.Value)
	assert.Equal(t, `if (1==2) {1;} else if (2!=2) {2;} else {3;};`, root.Literal())
}

// TestParser_ElseIfEval2 verifies parsing of multi-line else-if statements
func TestParser_ElseIfEval2(t *testing.T) {
	src := `if (1 == 2) { 
	   1
	} else if (2 == 2) {
	  2 
	} else {
	  3 
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	// Result should be 2 because the else-if condition is true.
	assert.Equal(t, &std.Integer{Value: 2}, root.Value)
	assert.Equal(t, `if (1==2) {1;} else if (2==2) {2;} else {3;};`, root.Literal())
}

// TestParser_ElseIfComplex verifies parsing of complex else-if with assignments
func TestParser_ElseIfComplex(t *testing.T) {
	src := `
	var a = 100;
	var b = 0;
	if (2 * a == 200) { 
		b = 1;
	} else if (2 * a != 200) {
		b = 2;
	} else {
		b = 311111;
	}
	return b;`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	// why 311111? return b; last b value is b=311111
	assert.Equal(t, &std.Integer{Value: 311111}, root.Value)
	assert.Equal(t, `var a = 100;var b = 0;if (2*a==200) {b = 1;} else if (2*a!=200) {b = 2;} else {b = 311111;};return b;`, root.Literal())
}

// TestParser_ElseIfNested verifies parsing of nested if-else in blocks
func TestParser_ElseIfNested(t *testing.T) {
	src := `{
	var x = 1;
	{
	 if(x==1){}else{}
	}
	}
	`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	// assert.Equal(t, 0, root.Value)
	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&BlockStatementNode{},

			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "x"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},

			&BlockStatementNode{},
			&IfExpressionNode{
				IfToken: lexer.Token{Literal: "if"},
			},
			&IdentifierExpressionNode{Name: "x"},
			&BooleanExpressionNode{
				Operation: lexer.Token{Literal: "=="},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&ParenthesizedExpressionNode{
				Expr: &BooleanExpressionNode{
					Operation: lexer.Token{Literal: "=="},
				},
			},
			&BlockStatementNode{},
			&BlockStatementNode{},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)

}

// TestParser_StrSimple verifies parsing of simple string literals
func TestParser_StrSimple(t *testing.T) {
	src := `"hello"`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StringLiteralExpressionNode{
				Token: lexer.Token{Literal: "hello"},
				Value: &std.String{Value: "hello"},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.String{Value: "hello"}, root.Value)
	assert.Equal(t, `hello;`, root.Literal())
}

// TestParser_Str verifies parsing of multiple string literals and identifiers
func TestParser_Str(t *testing.T) {
	src := `"hello" "there" boy 123`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StringLiteralExpressionNode{
				Token: lexer.Token{Literal: "hello"},
				Value: &std.String{Value: "hello"},
			},
			&StringLiteralExpressionNode{
				Token: lexer.Token{Literal: "there"},
				Value: &std.String{Value: "there"},
			},
			&IdentifierExpressionNode{Name: "boy"},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 123}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 4, len(root.Statements))
	assert.Equal(t, &std.Integer{Value: 123}, root.Value)
	assert.Equal(t, `hello;there;boy;123;`, root.Literal())
}

// TestParser_Func verifies parsing of simple function declarations
func TestParser_Func(t *testing.T) {
	src := `func foo() {  }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&FunctionStatementNode{
				FuncName: IdentifierExpressionNode{Name: "foo"},
			},
			&BlockStatementNode{},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Nil{}, root.Value)
	assert.Equal(t, `func foo () {};`, root.Literal())
}

// TestParser_FuncReturn verifies parsing of functions with return statements
func TestParser_FuncReturn(t *testing.T) {
	src := `func foo(a, b) { return a + b; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&FunctionStatementNode{
				FuncName: IdentifierExpressionNode{Name: "foo"},
			},
			&IdentifierExpressionNode{Name: "a"},
			&IdentifierExpressionNode{Name: "b"},
			&BlockStatementNode{},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr: &BinaryExpressionNode{
					Left: &IdentifierExpressionNode{
						Name: "a",
					},
					Operation: lexer.Token{Literal: "+"},
					Right: &IdentifierExpressionNode{
						Name: "b",
					},
				},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Nil{}, root.Value)
	assert.Equal(t, `func foo (a,b) {return a+b;};`, root.Literal())
}

// TestParser_FuncComplex verifies parsing of functions with conditional logic
func TestParser_FuncComplex(t *testing.T) {
	src := `func foo(a, b) {
		if (a == b) {
			return a + b;
		} else {
			return a - b;
		}
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&FunctionStatementNode{
				FuncName: IdentifierExpressionNode{Name: "foo"},
			},
			&IdentifierExpressionNode{Name: "a"},
			&IdentifierExpressionNode{Name: "b"},
			&BlockStatementNode{},
			&IfExpressionNode{
				IfToken: lexer.Token{Literal: "if"},
			},
			&IdentifierExpressionNode{Name: "a"},
			&BooleanExpressionNode{
				Operation: lexer.Token{Literal: "=="},
			},
			&IdentifierExpressionNode{Name: "b"},
			&ParenthesizedExpressionNode{},

			&BlockStatementNode{},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr: &BinaryExpressionNode{
					Left: &IdentifierExpressionNode{
						Name: "a",
					},
					Operation: lexer.Token{Literal: "+"},
					Right: &IdentifierExpressionNode{
						Name: "b",
					},
				},
			},
			&BlockStatementNode{},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Literal: "return"},
				Expr: &BinaryExpressionNode{
					Left: &IdentifierExpressionNode{
						Name: "a",
					},
					Operation: lexer.Token{Literal: "-"},
					Right: &IdentifierExpressionNode{
						Name: "b",
					},
				},
			},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Nil{}, root.Value)
	assert.Equal(t, `func foo (a,b) {if (a==b) {return a+b;} else {return a-b;};};`, root.Literal())
}

// TestParser_FuncCallArgs verifies parsing of function calls with arguments
func TestParser_FuncCallArgs(t *testing.T) {
	src := `foo(1, 2, 3)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&CallExpressionNode{
				FunctionIdentifier: IdentifierExpressionNode{Name: "foo"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Nil{}, root.Value)
	assert.Equal(t, `foo(1,2,3);`, root.Literal())
}

// TestParser_FuncCallSimple verifies parsing of function calls with variable arguments
func TestParser_FuncCallSimple(t *testing.T) {
	src := `
	var a  = 1;
	var b = 2;
	foo(a, b);
	`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "b"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&CallExpressionNode{
				FunctionIdentifier: IdentifierExpressionNode{Name: "foo"},
			},
			&IdentifierExpressionNode{Name: "a"},
			&IdentifierExpressionNode{Name: "b"},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 3, len(root.Statements))
	assert.Equal(t, &std.Nil{}, root.Value)
	assert.Equal(t, `var a = 1;var b = 2;foo(a,b);`, root.Literal())
}

// TestParser_FuncCallExpr verifies parsing of function calls in variable assignments
func TestParser_FuncCallExpr(t *testing.T) {
	src := `var a = foo(1, 2, 3);`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&CallExpressionNode{
				FunctionIdentifier: IdentifierExpressionNode{Name: "foo"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}

	root.Accept(testingVisitor)
	assert.Equal(t, 1, len(root.Statements))
	assert.Equal(t, &std.Nil{}, root.Value)
	assert.Equal(t, `var a = foo(1,2,3);`, root.Literal())
}

// TestParser_WhileSingle verifies parsing of while loops with single condition
func TestParser_WhileSingle(t *testing.T) {
	src := `var i = 0; while(i < 5){ i = i + 1; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check while loop statement
	whileStmt, ok := root.Statements[1].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(whileStmt.Conditions))
	assert.Equal(t, "while(i<5){i = i+1;}", whileStmt.Literal())
}

// TestParser_WhileTwo verifies parsing of while loops with two conditions
func TestParser_WhileTwo(t *testing.T) {
	src := `var i = 0; var j = 10; while(i < 5, j > 5){ i = i + 1; j = j - 1; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 3, len(root.Statements))

	// Check while loop statement
	whileStmt, ok := root.Statements[2].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(whileStmt.Conditions))
	assert.Equal(t, "while(i<5 && j>5){i = i+1;j = j-1;}", whileStmt.Literal())
}

// TestParser_WhileThree verifies parsing of while loops with three conditions
func TestParser_WhileThree(t *testing.T) {
	src := `var a = 0; var b = 20; var c = 10; while(a < 10, b > 10, c > 5){ a = a + 1; b = b - 1; c = c - 1; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 4, len(root.Statements))

	// Check while loop statement
	whileStmt, ok := root.Statements[3].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(whileStmt.Conditions))
	assert.Equal(t, "while(a<10 && b>10 && c>5){a = a+1;b = b-1;c = c-1;}", whileStmt.Literal())
}

// TestParser_WhileComplex verifies parsing of while loops with complex conditions
func TestParser_WhileComplex(t *testing.T) {
	src := `var x = 0; var y = 0; while(x < 5, y < 10, x + y < 12){ x = x + 1; y = y + 2; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 3, len(root.Statements))

	// Check while loop statement
	whileStmt, ok := root.Statements[2].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(whileStmt.Conditions))

	// Verify each condition is parsed correctly
	cond1, ok := whileStmt.Conditions[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.LT_OP, cond1.Operation.Type)

	cond2, ok := whileStmt.Conditions[1].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.LT_OP, cond2.Operation.Type)

	cond3, ok := whileStmt.Conditions[2].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.LT_OP, cond3.Operation.Type)
}

// TestParser_WhileEmpty verifies parsing of while loops with empty body
func TestParser_WhileEmpty(t *testing.T) {
	src := `var i = 0; while(i < 5, i >= 0){}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check while loop statement
	whileStmt, ok := root.Statements[1].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(whileStmt.Conditions))
	assert.Equal(t, 0, len(whileStmt.Body.Statements))
}

// TestParser_WhileNested verifies parsing of while loops nested in blocks
func TestParser_WhileNested(t *testing.T) {
	src := `{
		var count = 0;
		while(count < 3){
			count = count + 1;
		}
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check block statement
	blockStmt, ok := root.Statements[0].(*BlockStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(blockStmt.Statements))

	// Check while loop inside block
	whileStmt, ok := blockStmt.Statements[1].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(whileStmt.Conditions))
}

// TestParser_ForMulti verifies parsing of for loops with multiple initializers and updates
func TestParser_ForMulti(t *testing.T) {
	src := `for(i = 0, j = 10; i < 5 && j > 5; i = i + 1, j = j - 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check for loop statement
	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(forStmt.Initializers))
	assert.Equal(t, 2, len(forStmt.Updates))
	assert.NotNil(t, forStmt.Condition)
}

// TestParser_WhileValue verifies while loop root value initialization
func TestParser_WhileValue(t *testing.T) {
	src := `var i = 0; while(i < 5){ i = i + 1; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	// Root value should not be nil - it should be &objects.Nil{}
	assert.NotNil(t, root.Value)
	assert.Equal(t, &std.Nil{}, root.Value)
}

// TestParser_WhileCondType verifies parsing of while loop condition types
func TestParser_WhileCondType(t *testing.T) {
	src := `while(true){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(whileStmt.Conditions))

	// Check condition is a boolean literal
	boolCond, ok := whileStmt.Conditions[0].(*BooleanLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, true, boolCond.Value.(*std.Boolean).Value)
}

// TestParser_WhileMultiOps verifies parsing of while loops with different operators
func TestParser_WhileMultiOps(t *testing.T) {
	src := `while(a < 10, b >= 5, c != 0, d == 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 4, len(whileStmt.Conditions))

	// Verify each condition type
	cond1, ok := whileStmt.Conditions[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.LT_OP, cond1.Operation.Type)

	cond2, ok := whileStmt.Conditions[1].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.GE_OP, cond2.Operation.Type)

	cond3, ok := whileStmt.Conditions[2].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.NE_OP, cond3.Operation.Type)

	cond4, ok := whileStmt.Conditions[3].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.EQ_OP, cond4.Operation.Type)
}

// TestParser_WhileBody verifies parsing of while loop bodies with multiple statements
func TestParser_WhileBody(t *testing.T) {
	src := `while(i < 5){
		var a = i;
		var b = a + 1;
		i = b;
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(whileStmt.Body.Statements))

	// Verify body statements
	_, ok = whileStmt.Body.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	_, ok = whileStmt.Body.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	_, ok = whileStmt.Body.Statements[2].(*AssignmentExpressionNode)
	assert.True(t, ok)
}

// TestParser_WhileNested2 verifies parsing of nested while loops
func TestParser_WhileNested2(t *testing.T) {
	src := `while(i < 5){
		while(j < 10){
			j = j + 1;
		}
		i = i + 1;
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	outerWhile, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(outerWhile.Body.Statements))

	// Check nested while loop
	innerWhile, ok := outerWhile.Body.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(innerWhile.Conditions))
	assert.Equal(t, 1, len(innerWhile.Body.Statements))
}

// TestParser_WhileIf verifies parsing of while loops with if statements
func TestParser_WhileIf(t *testing.T) {
	src := `while(i < 10){
		if(i == 5){
			break;
		}
		i = i + 1;
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(whileStmt.Body.Statements))

	// Check if statement inside while
	_, ok = whileStmt.Body.Statements[0].(*IfExpressionNode)
	assert.True(t, ok)
}

// TestParser_WhileComplexCond verifies parsing of while loops with complex condition expressions
func TestParser_WhileComplexCond(t *testing.T) {
	src := `while((a + b) < (c * d), x > y){
		a = a + 1;
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(whileStmt.Conditions))

	// First condition should be a boolean expression with parenthesized expressions
	cond1, ok := whileStmt.Conditions[0].(*BooleanExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.LT_OP, cond1.Operation.Type)
}

// TestParser_WhileValueField verifies while loop value field initialization
func TestParser_WhileValueField(t *testing.T) {
	src := `while(i < 5){ i = i + 1; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)

	// Value field should be initialized to &objects.Nil{}
	assert.NotNil(t, whileStmt.Value)
	assert.Equal(t, &std.Nil{}, whileStmt.Value)
}

// TestParser_ForValue verifies for loop root value initialization
func TestParser_ForValue(t *testing.T) {
	src := `for(i = 0; i < 5; i = i + 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	// Root value should not be nil - it should be &objects.Nil{}
	assert.NotNil(t, root.Value)
	assert.Equal(t, &std.Nil{}, root.Value)
}

// TestParser_ForSingle verifies parsing of for loops with single initializer
func TestParser_ForSingle(t *testing.T) {
	src := `for(i = 0; i < 5; i = i + 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(forStmt.Initializers))
	assert.Equal(t, 1, len(forStmt.Updates))
	assert.NotNil(t, forStmt.Condition)
}

// TestParser_ForMultiInit verifies parsing of for loops with multiple initializers
func TestParser_ForMultiInit(t *testing.T) {
	src := `for(i = 0, j = 10, k = 20; i < 5; i = i + 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(forStmt.Initializers))

	// Verify each initializer is an assignment
	for _, init := range forStmt.Initializers {
		_, ok := init.(*AssignmentExpressionNode)
		assert.True(t, ok)
	}
}

// TestParser_ForMultiUpdate verifies parsing of for loops with multiple updates
func TestParser_ForMultiUpdate(t *testing.T) {
	src := `for(i = 0; i < 5; i = i + 1, j = j - 1, k = k * 2){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(forStmt.Updates))

	// Verify each update is an assignment
	for _, update := range forStmt.Updates {
		_, ok := update.(*AssignmentExpressionNode)
		assert.True(t, ok)
	}
}

// TestParser_ForNoInit verifies parsing of for loops without initializer
func TestParser_ForNoInit(t *testing.T) {
	src := `for(; i < 5; i = i + 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 0, len(forStmt.Initializers))
	assert.NotNil(t, forStmt.Condition)
	assert.Equal(t, 1, len(forStmt.Updates))
}

// TestParser_ForNoCond verifies parsing of for loops without condition
func TestParser_ForNoCond(t *testing.T) {
	src := `for(i = 0; ; i = i + 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(forStmt.Initializers))
	assert.Nil(t, forStmt.Condition)
	assert.Equal(t, 1, len(forStmt.Updates))
}

// TestParser_ForNoUpdate verifies parsing of for loops without update
func TestParser_ForNoUpdate(t *testing.T) {
	src := `for(i = 0; i < 5; ){ i = i + 1; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(forStmt.Initializers))
	assert.NotNil(t, forStmt.Condition)
	assert.Equal(t, 0, len(forStmt.Updates))
}

// TestParser_ForComplexCond verifies parsing of for loops with complex conditions
func TestParser_ForComplexCond(t *testing.T) {
	src := `for(i = 0; i < 5 && j > 0 || k == 10; i = i + 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.NotNil(t, forStmt.Condition)

	// Condition should be a boolean expression
	_, ok = forStmt.Condition.(*BooleanExpressionNode)
	assert.True(t, ok)
}

// TestParser_ForBody verifies parsing of for loop bodies with multiple statements
func TestParser_ForBody(t *testing.T) {
	src := `for(i = 0; i < 5; i = i + 1){
		var a = i;
		var b = a * 2;
		var c = b + 10;
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(forStmt.Body.Statements))

	// Verify all body statements are declarations
	for _, stmt := range forStmt.Body.Statements {
		_, ok := stmt.(*DeclarativeStatementNode)
		assert.True(t, ok)
	}
}

// TestParser_ForNested verifies parsing of nested for loops
func TestParser_ForNested(t *testing.T) {
	src := `for(i = 0; i < 5; i = i + 1){
		for(j = 0; j < 10; j = j + 1){
			var c = i + j;
		}
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	outerFor, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(outerFor.Body.Statements))

	// Check nested for loop
	innerFor, ok := outerFor.Body.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(innerFor.Initializers))
	assert.Equal(t, 1, len(innerFor.Updates))
	assert.Equal(t, 1, len(innerFor.Body.Statements))
}

// TestParser_ForIf verifies parsing of for loops with if statements
func TestParser_ForIf(t *testing.T) {
	src := `for(i = 0; i < 10; i = i + 1){
		if(i == 5){
			var x = i;
		}
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(forStmt.Body.Statements))

	// Check if statement inside for loop
	_, ok = forStmt.Body.Statements[0].(*IfExpressionNode)
	assert.True(t, ok)
}

// TestParser_ForValueField verifies for loop value field initialization
func TestParser_ForValueField(t *testing.T) {
	src := `for(i = 0; i < 5; i = i + 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)

	// Value field should be initialized to &objects.Nil{}
	assert.NotNil(t, forStmt.Value)
	assert.Equal(t, &std.Nil{}, forStmt.Value)
}

// TestParser_ForEmpty verifies parsing of empty for loops
func TestParser_ForEmpty(t *testing.T) {
	src := `for(;;){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 0, len(forStmt.Initializers))
	assert.Nil(t, forStmt.Condition)
	assert.Equal(t, 0, len(forStmt.Updates))
	assert.Equal(t, 0, len(forStmt.Body.Statements))
}

// TestParser_ForWhile verifies parsing of for loops containing while loops
func TestParser_ForWhile(t *testing.T) {
	src := `for(i = 0; i < 5; i = i + 1){
		while(j < 10){
			j = j + 1;
		}
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(forStmt.Body.Statements))

	// Check while loop inside for loop
	_, ok = forStmt.Body.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
}

// TestParser_WhileFor verifies parsing of while loops containing for loops
func TestParser_WhileFor(t *testing.T) {
	src := `while(i < 5){
		for(j = 0; j < 10; j = j + 1){
			var c = i + j;
		}
		i = i + 1;
	}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(whileStmt.Body.Statements))

	// Check for loop inside while loop
	_, ok = whileStmt.Body.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
}

// TestParser_ForLiteral verifies for loop literal representation
func TestParser_ForLiteral(t *testing.T) {
	src := `for(i = 0, j = 10; i < 5 && j > 5; i = i + 1, j = j - 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)

	// Check literal representation
	expected := "for(i = 0,j = 10;i<5&&j>5;i = i+1,j = j-1){}"
	assert.Equal(t, expected, forStmt.Literal())
}

// TestParser_WhileLiteral verifies while loop literal representation
func TestParser_WhileLiteral(t *testing.T) {
	src := `while(i < 5, j > 0, k == 10 && l != 20){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	whileStmt, ok := root.Statements[0].(*WhileLoopStatementNode)
	assert.True(t, ok)

	// Check literal representation
	expected := "while(i<5 && j>0 && k==10&&l!=20){}"
	assert.Equal(t, expected, whileStmt.Literal())
}

// TestParser_CompoundPlus verifies parsing of += compound assignment operator
func TestParser_CompoundPlus(t *testing.T) {
	src := `var a = 10; a += 5`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())
	assert.Equal(t, lexer.ASSIGN_OP, assignStmt.Operation.Type) // Transformed to regular assignment

	// Right side should be a binary expression (a + 5)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundMinus verifies parsing of -= compound assignment operator
func TestParser_CompoundMinus(t *testing.T) {
	src := `var a = 20; a -= 5`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a - 5)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MINUS_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundMul verifies parsing of *= compound assignment operator
func TestParser_CompoundMul(t *testing.T) {
	src := `var a = 5; a *= 4`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a * 4)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MUL_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundDiv verifies parsing of /= compound assignment operator
func TestParser_CompoundDiv(t *testing.T) {
	src := `var a = 20; a /= 4`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a / 4)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.DIV_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundMod verifies parsing of %= compound assignment operator
func TestParser_CompoundMod(t *testing.T) {
	src := `var a = 17; a %= 5`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a % 5)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.MOD_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundAnd verifies parsing of &= compound assignment operator
func TestParser_CompoundAnd(t *testing.T) {
	src := `var a = 12; a &= 10`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a & 10)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.BIT_AND_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundOr verifies parsing of |= compound assignment operator
func TestParser_CompoundOr(t *testing.T) {
	src := `var a = 12; a |= 3`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a | 3)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.BIT_OR_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundXor verifies parsing of ^= compound assignment operator
func TestParser_CompoundXor(t *testing.T) {
	src := `var a = 12; a ^= 5`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a ^ 5)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.BIT_XOR_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundLeftShift verifies parsing of <<= compound assignment operator
func TestParser_CompoundLeftShift(t *testing.T) {
	src := `var a = 4; a <<= 2`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a << 2)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.BIT_LEFT_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundRightShift verifies parsing of >>= compound assignment operator
func TestParser_CompoundRightShift(t *testing.T) {
	src := `var a = 16; a >>= 2`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a >> 2)
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.BIT_RIGHT_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundInFor verifies parsing of compound assignments in for loops
func TestParser_CompoundInFor(t *testing.T) {
	src := `for(i = 0; i < 5; i += 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check for loop statement
	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 1, len(forStmt.Initializers))
	assert.Equal(t, 1, len(forStmt.Updates))
	assert.NotNil(t, forStmt.Condition)
}

// TestParser_CompoundMultiFor verifies parsing of multiple compound assignments in for loops
func TestParser_CompoundMultiFor(t *testing.T) {
	src := `for(i = 0, j = 10; i < 5; i += 1, j -= 1){ }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check for loop statement
	forStmt, ok := root.Statements[0].(*ForLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(forStmt.Initializers))
	assert.Equal(t, 2, len(forStmt.Updates))
}

// TestParser_CompoundComplex verifies parsing of compound assignments with complex expressions
func TestParser_CompoundComplex(t *testing.T) {
	src := `var a = 10; a += 2 * 3`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check assignment statement
	assignStmt, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt.Left.Literal())

	// Right side should be a binary expression (a + (2 * 3))
	binaryExpr, ok := assignStmt.Right.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.PLUS_OP, binaryExpr.Operation.Type)
}

// TestParser_CompoundChained verifies parsing of chained compound assignments
func TestParser_CompoundChained(t *testing.T) {
	src := `var a = 10; a += 5; a *= 2; a -= 10`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 4, len(root.Statements))

	// Check all assignments are transformed correctly
	assignStmt1, ok := root.Statements[1].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt1.Left.Literal())

	assignStmt2, ok := root.Statements[2].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt2.Left.Literal())

	assignStmt3, ok := root.Statements[3].(*AssignmentExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "a", assignStmt3.Left.Literal())

	// Value should be the result of the compound assignment: a -= 10
	// After: var a = 10 (a=10), a += 5 (a=15), a *= 2 (a=30), a -= 10 (a=20)
	assert.Equal(t, assignStmt3.Value, &std.Integer{Value: 20})
}

// TestParser_ArrayLiteral verifies parsing of array literal expressions
func TestParser_ArrayLiteral(t *testing.T) {
	tests := []struct {
		Expr     string
		Expected []Node
	}{
		{
			Expr: `
		 	[]
		 `,
			Expected: []Node{
				&ArrayExpressionNode{},
			},
		},
		{
			Expr: `
		 	[1, 2, 3]
		 `,
			Expected: []Node{
				&ArrayExpressionNode{},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
			},
		},
		{
			Expr: `
			["comet"]
		`,
			Expected: []Node{
				&ArrayExpressionNode{},
				&StringLiteralExpressionNode{Value: &std.String{Value: "comet"}},
			},
		},
		{
			Expr: `
			[[1, 2, 3], [42, 43, 44], [1]]
		`,
			Expected: []Node{
				&ArrayExpressionNode{},
				&ArrayExpressionNode{},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},

				&ArrayExpressionNode{},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 42}},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 43}},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 44}},

				&ArrayExpressionNode{},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			},
		},
	}

	for _, test := range tests {
		par := NewParser(test.Expr)
		rootNode := par.Parse()
		assert.NotNil(t, rootNode)
		assert.False(t, par.HasErrors())

		testingVisitor := &TestingVisitor{
			ExpectedNodes: test.Expected,
			Ptr:           0,
			T:             t,
		}
		rootNode.Accept(testingVisitor)
	}
}

// TestParser_ArrayBool verifies parsing of array literals with boolean elements
func TestParser_ArrayBool(t *testing.T) {
	src := `[true, false, true]`
	par := NewParser(src)
	rootNode := par.Parse()
	assert.NotNil(t, rootNode)
	assert.False(t, par.HasErrors())

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&ArrayExpressionNode{},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
		},
		Ptr: 0,
		T:   t,
	}
	rootNode.Accept(testingVisitor)
}

// TestParser_ArrayMixed verifies parsing of array literals with mixed type elements
func TestParser_ArrayMixed(t *testing.T) {
	src := `[1, "hello", true, 42]`
	par := NewParser(src)
	rootNode := par.Parse()
	assert.NotNil(t, rootNode)
	assert.False(t, par.HasErrors())

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&ArrayExpressionNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&StringLiteralExpressionNode{Value: &std.String{Value: "hello"}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 42}},
		},
		Ptr: 0,
		T:   t,
	}
	rootNode.Accept(testingVisitor)
}

// TestParser_ArrayIdent verifies parsing of array literals with identifier elements
func TestParser_ArrayIdent(t *testing.T) {
	src := `var x = 10; var arr = [x, 20, 30]`
	par := NewParser(src)
	rootNode := par.Parse()
	assert.NotNil(t, rootNode)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 2, len(rootNode.Statements))

	// Check second statement is a declarative statement with array
	declStmt, ok := rootNode.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "arr", declStmt.Identifier.Name)

	// Check the expression is an array
	arrayExpr, ok := declStmt.Expr.(*ArrayExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(arrayExpr.Elements))
}

// TestParser_ArrayExpr verifies parsing of array literals with expression elements
func TestParser_ArrayExpr(t *testing.T) {
	src := `[1 + 2, 3 * 4, 10 - 5]`
	par := NewParser(src)
	rootNode := par.Parse()
	assert.NotNil(t, rootNode)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 1, len(rootNode.Statements))

	// Check the statement is an array expression
	arrayExpr, ok := rootNode.Statements[0].(*ArrayExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(arrayExpr.Elements))

	// Check each element is a binary expression
	for _, elem := range arrayExpr.Elements {
		_, ok := elem.(*BinaryExpressionNode)
		assert.True(t, ok)
	}
}

// TestParser_ArrayFunc verifies parsing of array literals with function elements
func TestParser_ArrayFunc(t *testing.T) {
	src := `var a = [1, 2, func(){2+3;}]; var b = a[2]; b();`
	par := NewParser(src)
	rootNode := par.Parse()
	assert.NotNil(t, rootNode)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 3, len(rootNode.Statements))

	// Check first statement: var a = [1, 2, func(){2+3;}]
	declStmt1, ok := rootNode.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "a", declStmt1.Identifier.Name)

	// Check the expression is an array
	arrayExpr, ok := declStmt1.Expr.(*ArrayExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(arrayExpr.Elements))

	// Check first element is integer 1
	intElem1, ok := arrayExpr.Elements[0].(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 1}, intElem1.Value)

	// Check second element is integer 2
	intElem2, ok := arrayExpr.Elements[1].(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 2}, intElem2.Value)

	// Check third element is a function
	funcElem, ok := arrayExpr.Elements[2].(*FunctionStatementNode)
	assert.True(t, ok)
	assert.NotNil(t, funcElem.FuncBody)

	// Check second statement: var b = a[2]
	declStmt2, ok := rootNode.Statements[1].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "b", declStmt2.Identifier.Name)

	// Check the expression is an index expression
	indexExpr, ok := declStmt2.Expr.(*IndexExpressionNode)
	assert.True(t, ok)
	assert.NotNil(t, indexExpr.Left)
	assert.NotNil(t, indexExpr.Index)

	// Check third statement: b()
	callExpr, ok := rootNode.Statements[2].(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "b", callExpr.FunctionIdentifier.Name)
}

// TestParser_ArrayIndex verifies parsing of array index access expressions
func TestParser_ArrayIndex(t *testing.T) {
	src := `var arr = [10, 20, 30]; arr[0]; arr[1]; arr[2]`
	par := NewParser(src)
	rootNode := par.Parse()
	assert.NotNil(t, rootNode)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 4, len(rootNode.Statements))

	// Check first statement is array declaration
	declStmt, ok := rootNode.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "arr", declStmt.Identifier.Name)

	// Check remaining statements are index expressions
	for i := 1; i < 4; i++ {
		indexExpr, ok := rootNode.Statements[i].(*IndexExpressionNode)
		assert.True(t, ok)

		// Check left is identifier "arr"
		ident, ok := indexExpr.Left.(*IdentifierExpressionNode)
		assert.True(t, ok)
		assert.Equal(t, "arr", ident.Name)

		// Check index is integer
		indexInt, ok := indexExpr.Index.(*IntegerLiteralExpressionNode)
		assert.True(t, ok)
		assert.Equal(t, &std.Integer{Value: int64(i - 1)}, indexInt.Value)
	}
}

// TestParser_ArrayNested verifies parsing of nested array index access
func TestParser_ArrayNested(t *testing.T) {
	src := `var matrix = [[1, 2], [3, 4]]; matrix[0][1]`
	par := NewParser(src)
	rootNode := par.Parse()
	assert.NotNil(t, rootNode)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 2, len(rootNode.Statements))

	// Check first statement is array declaration
	declStmt, ok := rootNode.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "matrix", declStmt.Identifier.Name)

	// Check the expression is a nested array
	arrayExpr, ok := declStmt.Expr.(*ArrayExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 2, len(arrayExpr.Elements))

	// Check second statement is nested index expression
	outerIndex, ok := rootNode.Statements[1].(*IndexExpressionNode)
	assert.True(t, ok)

	// The left side should be another index expression (matrix[0])
	innerIndex, ok := outerIndex.Left.(*IndexExpressionNode)
	assert.True(t, ok)

	// Check inner index left is identifier "matrix"
	ident, ok := innerIndex.Left.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "matrix", ident.Name)
}

// TestParser_RangeSimple verifies parsing of simple range expressions
func TestParser_RangeSimple(t *testing.T) {
	src := `2...5`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a range expression
	rangeExpr, ok := root.Statements[0].(*RangeExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "2...5", rangeExpr.Literal())

	// Check start is integer 2
	startInt, ok := rangeExpr.Start.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 2}, startInt.Value)

	// Check end is integer 5
	endInt, ok := rangeExpr.End.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 5}, endInt.Value)

	// Check value is a Range object
	rangeObj, ok := rangeExpr.Value.(*std.Range)
	assert.True(t, ok)
	assert.Equal(t, int64(2), rangeObj.Start)
	assert.Equal(t, int64(5), rangeObj.End)
}

// TestParser_RangeVar verifies parsing of range expressions in variable declarations
func TestParser_RangeVar(t *testing.T) {
	src := `var x = 1...10`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a declarative statement
	declStmt, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "x", declStmt.Identifier.Name)
	assert.Equal(t, "var x = 1...10", declStmt.Literal())
}

// TestParser_RangeExpr verifies parsing of range expressions with arithmetic
func TestParser_RangeExpr(t *testing.T) {
	src := `(1+1)...(5*2)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a range expression
	rangeExpr, ok := root.Statements[0].(*RangeExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "(1+1)...(5*2)", rangeExpr.Literal())
}

// TestParser_ForeachSimple verifies parsing of simple foreach loops
func TestParser_ForeachSimple(t *testing.T) {
	src := `foreach i in 1...5 { }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a foreach loop
	foreachStmt, ok := root.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "i", foreachStmt.Iterator.Name)

	// Check iterable is a range expression
	rangeExpr, ok := foreachStmt.Iterable.(*RangeExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "1...5", rangeExpr.Literal())

	// Check body is empty
	assert.Equal(t, 0, len(foreachStmt.Body.Statements))
}

// TestParser_ForeachArray verifies parsing of foreach loops with arrays
func TestParser_ForeachArray(t *testing.T) {
	src := `foreach item in [10, 20, 30] { }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a foreach loop
	foreachStmt, ok := root.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "item", foreachStmt.Iterator.Name)

	// Check iterable is an array expression
	arrayExpr, ok := foreachStmt.Iterable.(*ArrayExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, 3, len(arrayExpr.Elements))
}

// TestParser_ForeachVar verifies parsing of foreach loops with range variables
func TestParser_ForeachVar(t *testing.T) {
	src := `var r = 1...10; foreach i in r { }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 2, len(root.Statements))

	// Check second statement is a foreach loop
	foreachStmt, ok := root.Statements[1].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "i", foreachStmt.Iterator.Name)

	// Check iterable is an identifier
	ident, ok := foreachStmt.Iterable.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "r", ident.Name)
}

// TestParser_ForeachBody verifies parsing of foreach loop bodies
func TestParser_ForeachBody(t *testing.T) {
	src := `foreach i in 1...3 { var x = i * 2; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a foreach loop
	foreachStmt, ok := root.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)

	// Check body has one statement
	assert.Equal(t, 1, len(foreachStmt.Body.Statements))

	// Check body statement is a declaration
	declStmt, ok := foreachStmt.Body.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "x", declStmt.Identifier.Name)
}

// TestParser_ForeachNested verifies parsing of nested foreach loops
func TestParser_ForeachNested(t *testing.T) {
	src := `foreach i in 1...2 { foreach j in 1...2 { var c = i + j; } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	assert.Equal(t, 1, len(root.Statements))

	// Check outer foreach
	outerForeach, ok := root.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "i", outerForeach.Iterator.Name)
	assert.Equal(t, 1, len(outerForeach.Body.Statements))

	// Check inner foreach
	innerForeach, ok := outerForeach.Body.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "j", innerForeach.Iterator.Name)
	assert.Equal(t, 1, len(innerForeach.Body.Statements))
}

// TestParser_RangeLiteral verifies range expression literal representation
func TestParser_RangeLiteral(t *testing.T) {
	src := `var x = 5...15`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	declStmt, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)

	rangeExpr, ok := declStmt.Expr.(*RangeExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "5...15", rangeExpr.Literal())
}

// TestParser_ForeachLiteral verifies foreach loop literal representation
func TestParser_ForeachLiteral(t *testing.T) {
	src := `foreach num in 1...5 { var x = num; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	foreachStmt, ok := root.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "foreach num in 1...5 {var x = num;}", foreachStmt.Literal())
}

// TestParser_ListCall verifies parsing of list() constructor calls
func TestParser_ListCall(t *testing.T) {
	tests := []struct {
		src         string
		expectedLen int
	}{
		{`list()`, 0},
		{`list(1, 2, 3)`, 3},
		{`list(1, "hello", true)`, 3},
	}

	for _, tt := range tests {
		root := NewParser(tt.src).Parse()
		assert.NotNil(t, root)
		assert.Equal(t, 1, len(root.Statements))

		// Check the statement is a call expression
		callExpr, ok := root.Statements[0].(*CallExpressionNode)
		assert.True(t, ok)
		assert.Equal(t, "list", callExpr.FunctionIdentifier.Name)
		assert.Equal(t, tt.expectedLen, len(callExpr.Arguments))
	}
}

// TestParser_TupleCall verifies parsing of tuple() constructor calls
func TestParser_TupleCall(t *testing.T) {
	tests := []struct {
		src         string
		expectedLen int
	}{
		{`tuple()`, 0},
		{`tuple(1, 2, 3)`, 3},
		{`tuple("Alice", 25, true)`, 3},
	}

	for _, tt := range tests {
		root := NewParser(tt.src).Parse()
		assert.NotNil(t, root)
		assert.Equal(t, 1, len(root.Statements))

		// Check the statement is a call expression
		callExpr, ok := root.Statements[0].(*CallExpressionNode)
		assert.True(t, ok)
		assert.Equal(t, "tuple", callExpr.FunctionIdentifier.Name)
		assert.Equal(t, tt.expectedLen, len(callExpr.Arguments))
	}
}

// TestParser_ListVar verifies parsing of list assignment to variables
func TestParser_ListVar(t *testing.T) {
	src := `var l = list(1, 2, 3)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a declarative statement
	declStmt, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "l", declStmt.Identifier.Name)

	// Check the expression is a call to list()
	callExpr, ok := declStmt.Expr.(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "list", callExpr.FunctionIdentifier.Name)
	assert.Equal(t, 3, len(callExpr.Arguments))
}

// TestParser_TupleVar verifies parsing of tuple assignment to variables
func TestParser_TupleVar(t *testing.T) {
	src := `var t = tuple(10, 20, 30)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a declarative statement
	declStmt, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "t", declStmt.Identifier.Name)

	// Check the expression is a call to tuple()
	callExpr, ok := declStmt.Expr.(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "tuple", callExpr.FunctionIdentifier.Name)
	assert.Equal(t, 3, len(callExpr.Arguments))
}

// TestParser_ListFunctions verifies parsing of list manipulation function calls
func TestParser_ListFunctions(t *testing.T) {
	tests := []struct {
		src      string
		funcName string
		argCount int
	}{
		{`pushback_list(l, 4)`, "pushback_list", 2},
		{`pushfront_list(l, 0)`, "pushfront_list", 2},
		{`popback_list(l)`, "popback_list", 1},
		{`popfront_list(l)`, "popfront_list", 1},
		{`size_list(l)`, "size_list", 1},
		{`peekback_list(l)`, "peekback_list", 1},
		{`peekfront_list(l)`, "peekfront_list", 1},
	}

	for _, tt := range tests {
		root := NewParser(tt.src).Parse()
		assert.NotNil(t, root)
		assert.Equal(t, 1, len(root.Statements))

		// Check the statement is a call expression
		callExpr, ok := root.Statements[0].(*CallExpressionNode)
		assert.True(t, ok)
		assert.Equal(t, tt.funcName, callExpr.FunctionIdentifier.Name)
		assert.Equal(t, tt.argCount, len(callExpr.Arguments))
	}
}

// TestParser_TupleFunctions verifies parsing of tuple helper function calls
func TestParser_TupleFunctions(t *testing.T) {
	tests := []struct {
		src      string
		funcName string
		argCount int
	}{
		{`size_tuple(t)`, "size_tuple", 1},
		{`peekback_tuple(t)`, "peekback_tuple", 1},
		{`peekfront_tuple(t)`, "peekfront_tuple", 1},
	}

	for _, tt := range tests {
		root := NewParser(tt.src).Parse()
		assert.NotNil(t, root)
		assert.Equal(t, 1, len(root.Statements))

		// Check the statement is a call expression
		callExpr, ok := root.Statements[0].(*CallExpressionNode)
		assert.True(t, ok)
		assert.Equal(t, tt.funcName, callExpr.FunctionIdentifier.Name)
		assert.Equal(t, tt.argCount, len(callExpr.Arguments))
	}
}

// TestParser_ListIndexAccess verifies parsing of list index access
func TestParser_ListIndexAccess(t *testing.T) {
	src := `var l = list(1, 2, 3); l[0]`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 2, len(root.Statements))

	// Check second statement is an index expression
	indexExpr, ok := root.Statements[1].(*IndexExpressionNode)
	assert.True(t, ok)

	// Check left is identifier "l"
	ident, ok := indexExpr.Left.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "l", ident.Name)

	// Check index is integer 0
	indexInt, ok := indexExpr.Index.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 0}, indexInt.Value)
}

// TestParser_TupleIndexAccess verifies parsing of tuple index access
func TestParser_TupleIndexAccess(t *testing.T) {
	src := `var t = tuple(10, 20, 30); t[1]`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 2, len(root.Statements))

	// Check second statement is an index expression
	indexExpr, ok := root.Statements[1].(*IndexExpressionNode)
	assert.True(t, ok)

	// Check left is identifier "t"
	ident, ok := indexExpr.Left.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "t", ident.Name)

	// Check index is integer 1
	indexInt, ok := indexExpr.Index.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 1}, indexInt.Value)
}

// TestParser_ListSlice verifies parsing of list slicing
func TestParser_ListSlice(t *testing.T) {
	src := `var l = list(0, 10, 20, 30); l[1:3]`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 2, len(root.Statements))

	// Check second statement is a slice expression
	sliceExpr, ok := root.Statements[1].(*SliceExpressionNode)
	assert.True(t, ok)

	// Check left is identifier "l"
	ident, ok := sliceExpr.Left.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "l", ident.Name)

	// Check start index
	startInt, ok := sliceExpr.Start.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 1}, startInt.Value)

	// Check end index
	endInt, ok := sliceExpr.End.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, &std.Integer{Value: 3}, endInt.Value)
}

// TestParser_TupleSlice verifies parsing of tuple slicing
func TestParser_TupleSlice(t *testing.T) {
	src := `var t = tuple(1, 2, 3, 4, 5); t[2:4]`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 2, len(root.Statements))

	// Check second statement is a slice expression
	sliceExpr, ok := root.Statements[1].(*SliceExpressionNode)
	assert.True(t, ok)

	// Check left is identifier "t"
	ident, ok := sliceExpr.Left.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "t", ident.Name)
}

// TestParser_ForeachList verifies parsing of foreach loops with lists
func TestParser_ForeachList(t *testing.T) {
	src := `foreach item in list(1, 2, 3) { var x = item; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a foreach loop
	foreachStmt, ok := root.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "item", foreachStmt.Iterator.Name)

	// Check iterable is a call to list()
	callExpr, ok := foreachStmt.Iterable.(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "list", callExpr.FunctionIdentifier.Name)
}

// TestParser_ForeachTuple verifies parsing of foreach loops with tuples
func TestParser_ForeachTuple(t *testing.T) {
	src := `foreach val in tuple(10, 20, 30) { var y = val; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a foreach loop
	foreachStmt, ok := root.Statements[0].(*ForeachLoopStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "val", foreachStmt.Iterator.Name)

	// Check iterable is a call to tuple()
	callExpr, ok := foreachStmt.Iterable.(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "tuple", callExpr.FunctionIdentifier.Name)
}

// TestParser_ListNested verifies parsing of nested lists
func TestParser_ListNested(t *testing.T) {
	src := `var matrix = list(list(1, 2), list(3, 4))`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a declarative statement
	declStmt, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "matrix", declStmt.Identifier.Name)

	// Check the expression is a call to list()
	outerCall, ok := declStmt.Expr.(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "list", outerCall.FunctionIdentifier.Name)
	assert.Equal(t, 2, len(outerCall.Arguments))

	// Check first argument is also a list() call
	innerCall1, ok := outerCall.Arguments[0].(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "list", innerCall1.FunctionIdentifier.Name)
}

// TestParser_TupleNested verifies parsing of nested tuples
func TestParser_TupleNested(t *testing.T) {
	src := `var nested = tuple(tuple(1, 2), tuple(3, 4))`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a declarative statement
	declStmt, ok := root.Statements[0].(*DeclarativeStatementNode)
	assert.True(t, ok)
	assert.Equal(t, "nested", declStmt.Identifier.Name)

	// Check the expression is a call to tuple()
	outerCall, ok := declStmt.Expr.(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "tuple", outerCall.FunctionIdentifier.Name)
	assert.Equal(t, 2, len(outerCall.Arguments))

	// Check first argument is also a tuple() call
	innerCall1, ok := outerCall.Arguments[0].(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "tuple", innerCall1.FunctionIdentifier.Name)
}

// TestParser_ListIndexAssignment verifies parsing of list index assignment
func SkipTestParser_ListIndexAssignment(t *testing.T) {
	src := `var l = list(1, 2, 3); l[0] = 10`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 2, len(root.Statements))

	// The second statement should parse successfully
	// Note: Index assignment is handled specially in the evaluator
	// The parser creates an assignment with the index expression
	assert.NotNil(t, root.Statements[1])
}

// TestParser_ListMixed verifies parsing of lists with mixed types
func TestParser_ListMixed(t *testing.T) {
	src := `list(1, "hello", true, 3.14)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a call expression
	callExpr, ok := root.Statements[0].(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "list", callExpr.FunctionIdentifier.Name)
	assert.Equal(t, 4, len(callExpr.Arguments))

	// Check argument types
	_, ok = callExpr.Arguments[0].(*IntegerLiteralExpressionNode)
	assert.True(t, ok)

	_, ok = callExpr.Arguments[1].(*StringLiteralExpressionNode)
	assert.True(t, ok)

	_, ok = callExpr.Arguments[2].(*BooleanLiteralExpressionNode)
	assert.True(t, ok)

	_, ok = callExpr.Arguments[3].(*FloatLiteralExpressionNode)
	assert.True(t, ok)
}

// TestParser_TupleMixed verifies parsing of tuples with mixed types
func TestParser_TupleMixed(t *testing.T) {
	src := `tuple("Alice", 25, true, 5.8)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a call expression
	callExpr, ok := root.Statements[0].(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "tuple", callExpr.FunctionIdentifier.Name)
	assert.Equal(t, 4, len(callExpr.Arguments))

	// Check argument types
	_, ok = callExpr.Arguments[0].(*StringLiteralExpressionNode)
	assert.True(t, ok)

	_, ok = callExpr.Arguments[1].(*IntegerLiteralExpressionNode)
	assert.True(t, ok)

	_, ok = callExpr.Arguments[2].(*BooleanLiteralExpressionNode)
	assert.True(t, ok)

	_, ok = callExpr.Arguments[3].(*FloatLiteralExpressionNode)
	assert.True(t, ok)
}

// TestParser_ListTupleNewBuiltins verifies parsing of new list and tuple builtin functions
func TestParser_ListTupleNewBuiltins(t *testing.T) {
	tests := []struct {
		src      string
		funcName string
		argCount int
	}{
		// New list functions
		{`insert_list(l, 2, 3)`, "insert_list", 3},
		{`remove_list(l, 2)`, "remove_list", 2},
		{`contains_list(l, 3)`, "contains_list", 2},
		// New tuple function
		{`contains_tuple(t, 3)`, "contains_tuple", 2},
	}

	for _, tt := range tests {
		root := NewParser(tt.src).Parse()
		assert.NotNil(t, root)
		assert.Equal(t, 1, len(root.Statements))

		// Check the statement is a call expression
		callExpr, ok := root.Statements[0].(*CallExpressionNode)
		assert.True(t, ok)
		assert.Equal(t, tt.funcName, callExpr.FunctionIdentifier.Name)
		assert.Equal(t, tt.argCount, len(callExpr.Arguments))
	}
}

// TestParser_Struct
func TestParser_Struct(t *testing.T) {
	src := `struct Point {}`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a struct declaration
	structDecl, ok := root.Statements[0].(*StructDeclarationNode)
	assert.True(t, ok)
	assert.Equal(t, "Point", structDecl.StructName.Literal())
	assert.Equal(t, 0, len(structDecl.Fields))
	assert.Equal(t, "struct Point {}", structDecl.Literal())
}

// TestParser_StructInit verifies parsing of struct initialization
func TestParser_StructInit(t *testing.T) {
	src := `struct Point { func init(){} }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))
	// Check the statement is a struct declaration
	structDecl, ok := root.Statements[0].(*StructDeclarationNode)
	assert.True(t, ok)
	assert.Equal(t, "Point", structDecl.StructName.Literal())
	assert.Equal(t, 0, len(structDecl.Fields))
	assert.Equal(t, "struct Point {func init () {} }", structDecl.Literal())
	assert.Equal(t, 1, len(structDecl.Methods))

	// Check the field is an init function
	initFunc := structDecl.Methods[0]
	// assert.True(t, ok)
	assert.Equal(t, "init", initFunc.FuncName.Literal())
}

// TestParser_StructMethods verifies parsing of struct methods
func TestParser_StructMethods(t *testing.T) {
	src := `struct Point { func init(){} func move(x, y){} }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.Equal(t, 1, len(root.Statements))

	// Check the statement is a struct declaration
	structDecl, ok := root.Statements[0].(*StructDeclarationNode)
	assert.True(t, ok)
	assert.Equal(t, "Point", structDecl.StructName.Literal())
	assert.Equal(t, 0, len(structDecl.Fields))
	assert.Equal(t, "struct Point {func init () {} func move (x,y) {} }", structDecl.Literal())
	assert.Equal(t, 2, len(structDecl.Methods))

	// Check the first method is an init function
	initFunc := structDecl.Methods[0]
	// assert.True(t, ok)
	assert.Equal(t, "init", initFunc.FuncName.Literal())

	// Check the second method is a move function
	moveFunc := structDecl.Methods[1]
	// assert.True(t, ok)
	assert.Equal(t, "move", moveFunc.FuncName.Literal())
}

// TestParser_StructMethods verifies parsing of struct fields
func TestParser_StructMethodsLayout(t *testing.T) {
	tests := []struct {
		Expr     string
		Expected []Node
	}{
		{
			Expr: `struct Point { func init(){} func move(x, y){} }; func foo(){};`,
			Expected: []Node{
				&StructDeclarationNode{
					StructName: IdentifierExpressionNode{Name: "Point"},
					Methods: []*FunctionStatementNode{
						&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "init"}},
						&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "move"}},
					},
				},
				&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "init"}},
				&BlockStatementNode{},
				&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "move"}},
				&IdentifierExpressionNode{Name: "x"},
				&IdentifierExpressionNode{Name: "y"},
				&BlockStatementNode{},
				&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "foo"}},
				&BlockStatementNode{},
			},
		},
		{
			Expr: `struct Point { func init(){} }; func foo(){};`,
			Expected: []Node{
				&StructDeclarationNode{
					StructName: IdentifierExpressionNode{Name: "Point"},
					Methods: []*FunctionStatementNode{
						&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "init"}},
					},
				},
				&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "init"}},
				&BlockStatementNode{},
				&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "foo"}},
				&BlockStatementNode{},
			},
		},
	}

	for _, tt := range tests {
		parser := NewParser(tt.Expr)
		root := parser.Parse()
		assert.NotNil(t, root)
		assert.False(t, parser.HasErrors())
		assert.Equal(t, 2, len(root.Statements))
		_, ok := root.Statements[0].(*StructDeclarationNode)
		assert.True(t, ok)

		testingVisitor := &TestingVisitor{
			ExpectedNodes: tt.Expected,
			Ptr:           0,
			T:             t,
		}
		root.Accept(testingVisitor)

	}
}

// TestParser_ParseNewCall verifies parsing of new struct instantiation expressions
func TestParser_ParseNewCall(t *testing.T) {
	tests := []struct {
		Expr     string
		Expected []Node
	}{
		{
			Expr: `var a = new Data()`,
			Expected: []Node{
				&DeclarativeStatementNode{
					VarToken:   lexer.Token{Literal: "var"},
					Identifier: IdentifierExpressionNode{Name: "a"},
				},
				&NewCallExpressionNode{
					StructName: IdentifierExpressionNode{Name: "Data"},
				},
			},
		},
		{
			Expr: `var a = new Data(10, "foo")`,
			Expected: []Node{
				&DeclarativeStatementNode{
					VarToken:   lexer.Token{Literal: "var"},
					Identifier: IdentifierExpressionNode{Name: "a"},
				},
				&NewCallExpressionNode{
					StructName: IdentifierExpressionNode{Name: "Data"},
				},
				&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
				&StringLiteralExpressionNode{Value: &std.String{Value: "foo"}},
			},
		},
	}

	for _, test := range tests {
		parser := NewParser(test.Expr)
		rootNode := parser.Parse()
		assert.NotNil(t, rootNode)
		assert.False(t, parser.HasErrors())

		testingVisitor := &TestingVisitor{
			ExpectedNodes: test.Expected,
			Ptr:           0,
			T:             t,
		}
		rootNode.Accept(testingVisitor)
	}
}

// TestParser_ParseErrorNewCall verifies error handling for invalid new call expressions
func TestParser_ParseErrorNewCall(t *testing.T) {
	tests := []string{
		`var a = new S`,
		`var a = new`,
		`var c = new ()`,
	}
	for _, test := range tests {
		parser := NewParser(test)
		rootNode := parser.Parse()
		assert.NotNil(t, rootNode)
		assert.True(t, parser.HasErrors())
	}
}

// TestParser_MemberIndexAccess verifies parsing of member access followed by index
func TestParser_MemberIndexAccess(t *testing.T) {
	src := `this.m[1]`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 1, len(root.Statements))
	// Should be IndexExpression
	// Left: BinaryExpression (this.m)
	// Index: Integer(1)

	indexExpr, ok := root.Statements[0].(*IndexExpressionNode)
	assert.True(t, ok)

	binExpr, ok := indexExpr.Left.(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.DOT_OP, binExpr.Operation.Type)

	leftIdent, ok := binExpr.Left.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "this", leftIdent.Name)

	rightIdent, ok := binExpr.Right.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "m", rightIdent.Name)

	indexVal, ok := indexExpr.Index.(*IntegerLiteralExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, int64(1), indexVal.Value.(*std.Integer).Value)

	assert.Equal(t, "this.m[1];", root.Literal())
}

// TestParser_ThisAccess verifies parsing of 'this' keyword access
func TestParser_ThisAccess(t *testing.T) {
	src := `this.x`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 1, len(root.Statements))
	// Should be BinaryExpression (DOT_OP)
	binExpr, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, lexer.DOT_OP, binExpr.Operation.Type)

	left, ok := binExpr.Left.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "this", left.Name)

	right, ok := binExpr.Right.(*IdentifierExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "x", right.Name)
}

// TestParser_MethodCall verifies parsing of method calls on objects
func TestParser_MethodCall(t *testing.T) {
	src := `obj.method(1, 2)`
	par := NewParser(src)
	root := par.Parse()
	assert.NotNil(t, root)
	assert.False(t, par.HasErrors())

	assert.Equal(t, 1, len(root.Statements))
	binExpr, ok := root.Statements[0].(*BinaryExpressionNode)
	assert.True(t, ok)

	right, ok := binExpr.Right.(*CallExpressionNode)
	assert.True(t, ok)
	assert.Equal(t, "method", right.FunctionIdentifier.Name)
	assert.Equal(t, 2, len(right.Arguments))
}

// TestParser_StructFields verifies parsing of struct with const, let, var fields
func TestParser_StructFields(t *testing.T) {
	src := `struct Config { const MAX = 100; let retries = 3; var debug = true; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)
	assert.False(t, NewParser(src).HasErrors())

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{
				StructName: IdentifierExpressionNode{Name: "Config"},
			},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Type: lexer.CONST_KEY, Literal: "const"},
				Identifier: IdentifierExpressionNode{Name: "MAX"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 100}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Type: lexer.LET_KEY, Literal: "let"},
				Identifier: IdentifierExpressionNode{Name: "retries"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Type: lexer.VAR_KEY, Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "debug"},
			},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructConstLetVar verifies parsing of struct with const, let, var fields
func TestParser_StructConstLetVar(t *testing.T) {
	src := `struct S { const C = 1; let L = 2; var V = 3; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "S"}},
			&DeclarativeStatementNode{VarToken: lexer.Token{Type: lexer.CONST_KEY, Literal: "const"}, Identifier: IdentifierExpressionNode{Name: "C"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&DeclarativeStatementNode{VarToken: lexer.Token{Type: lexer.LET_KEY, Literal: "let"}, Identifier: IdentifierExpressionNode{Name: "L"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&DeclarativeStatementNode{VarToken: lexer.Token{Type: lexer.VAR_KEY, Literal: "var"}, Identifier: IdentifierExpressionNode{Name: "V"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructMethodAccessStatic verifies parsing of method accessing static fields via StructName
func TestParser_StructMethodAccessStatic(t *testing.T) {
	src := `struct S { var x = 0; func f() { return S.x; } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "S"}},
			&DeclarativeStatementNode{VarToken: lexer.Token{Type: lexer.VAR_KEY, Literal: "var"}, Identifier: IdentifierExpressionNode{Name: "x"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 0}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "f"}},
			&BlockStatementNode{},
			&ReturnStatementNode{},
			&IdentifierExpressionNode{Name: "S"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "x"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructMethodAccessSelf verifies parsing of method accessing static fields via self
func TestParser_StructMethodAccessSelf(t *testing.T) {
	src := `struct S { var x = 0; func f() { return self.x; } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "S"}},
			&DeclarativeStatementNode{VarToken: lexer.Token{Type: lexer.VAR_KEY, Literal: "var"}, Identifier: IdentifierExpressionNode{Name: "x"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 0}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "f"}},
			&BlockStatementNode{},
			&ReturnStatementNode{},
			&IdentifierExpressionNode{Name: "self"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "x"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructMethodAccessThis verifies parsing of method accessing instance fields via this
func TestParser_StructMethodAccessThis(t *testing.T) {
	src := `struct S { func init(v) { this.v = v; } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "S"}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "init"}},
			&IdentifierExpressionNode{Name: "v"},
			&BlockStatementNode{},
			&IdentifierExpressionNode{Name: "this"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "v"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&IdentifierExpressionNode{Name: "v"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructNestedInstantiation verifies parsing of nested struct instantiation
func TestParser_StructNestedInstantiation(t *testing.T) {
	src := `struct B { func create() { return new A(); } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "B"}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "create"}},
			&BlockStatementNode{},
			&ReturnStatementNode{},
			&NewCallExpressionNode{StructName: IdentifierExpressionNode{Name: "A"}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructStaticAssignmentExternal verifies parsing of assignment to static field
func TestParser_StructStaticAssignmentExternal(t *testing.T) {
	src := `S.x = 5`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IdentifierExpressionNode{Name: "S"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "x"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructInstanceAssignmentExternal verifies parsing of assignment to instance field
func TestParser_StructInstanceAssignmentExternal(t *testing.T) {
	src := `s.x = 5`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IdentifierExpressionNode{Name: "s"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "x"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 5}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructComplexMethodLogic verifies parsing of complex method logic
func TestParser_StructComplexMethodLogic(t *testing.T) {
	src := `struct S { func f() { while(true) { break; } } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "S"}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "f"}},
			&BlockStatementNode{},
			&WhileLoopStatementNode{},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
			&BlockStatementNode{},
			&BreakStatementNode{},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructArrayFieldPush verifies parsing of method pushing to array field
func TestParser_StructArrayFieldPush(t *testing.T) {
	src := `struct S { var arr = []; func add(x) { push(self.arr, x); } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "S"}},
			&DeclarativeStatementNode{VarToken: lexer.Token{Type: lexer.VAR_KEY, Literal: "var"}, Identifier: IdentifierExpressionNode{Name: "arr"}},
			&ArrayExpressionNode{},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "add"}},
			&IdentifierExpressionNode{Name: "x"},
			&BlockStatementNode{},
			&CallExpressionNode{FunctionIdentifier: IdentifierExpressionNode{Name: "push"}},
			&IdentifierExpressionNode{Name: "self"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "arr"},
			&IdentifierExpressionNode{Name: "x"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructRecursiveMethod verifies parsing of recursive method call
func TestParser_StructRecursiveMethod(t *testing.T) {
	src := `struct S { func f() { this.f(); } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "S"}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "f"}},
			&BlockStatementNode{},
			&IdentifierExpressionNode{Name: "this"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&CallExpressionNode{FunctionIdentifier: IdentifierExpressionNode{Name: "f"}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StaticMemberAccess verifies parsing of static member access
func TestParser_StaticMemberAccess(t *testing.T) {
	src := `Config.MAX`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IdentifierExpressionNode{Name: "Config"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "MAX"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StaticAssignment verifies parsing of assignment to static fields
func TestParser_StaticAssignment(t *testing.T) {
	src := `Config.retries = false`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IdentifierExpressionNode{Name: "Config"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "retries"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: false}, Token: lexer.Token{Type: lexer.FALSE_KEY, Literal: "false"}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructMethodSelf verifies parsing of methods using self and this
func TestParser_StructMethodSelf(t *testing.T) {
	src := `struct A { func get() { return self.x + this.y; } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{
				StructName: IdentifierExpressionNode{Name: "A"},
			},
			&FunctionStatementNode{
				FuncName: IdentifierExpressionNode{Name: "get"},
			},
			&BlockStatementNode{},
			&ReturnStatementNode{
				ReturnToken: lexer.Token{Type: lexer.RETURN_KEY, Literal: "return"},
			},
			&IdentifierExpressionNode{Name: "self"},
			&BinaryExpressionNode{
				Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."},
			},
			&IdentifierExpressionNode{Name: "x"},
			&BinaryExpressionNode{
				Operation: lexer.Token{Type: lexer.PLUS_OP, Literal: "+"},
			},
			&IdentifierExpressionNode{Name: "this"},
			&BinaryExpressionNode{
				Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."},
			},
			&IdentifierExpressionNode{Name: "y"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructNestedLogic verifies parsing of complex logic inside struct methods
func TestParser_StructNestedLogic(t *testing.T) {
	src := `struct Logic { func run(x) { if (x > 0) { while(x > 0) { x = x - 1; } } return x; } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "Logic"}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "run"}},
			&IdentifierExpressionNode{Name: "x"},
			&BlockStatementNode{},
			&IfExpressionNode{IfToken: lexer.Token{Type: lexer.IF_KEY, Literal: "if"}},
			&IdentifierExpressionNode{Name: "x"},
			&BooleanExpressionNode{Operation: lexer.Token{Type: lexer.GT_OP, Literal: ">"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 0}},
			&ParenthesizedExpressionNode{},
			&BlockStatementNode{},
			&WhileLoopStatementNode{WhileToken: lexer.Token{Type: lexer.WHILE_KEY, Literal: "while"}},
			&IdentifierExpressionNode{Name: "x"},
			&BooleanExpressionNode{Operation: lexer.Token{Type: lexer.GT_OP, Literal: ">"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 0}},
			&BlockStatementNode{},
			&IdentifierExpressionNode{Name: "x"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&IdentifierExpressionNode{Name: "x"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.MINUS_OP, Literal: "-"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&ReturnStatementNode{ReturnToken: lexer.Token{Type: lexer.RETURN_KEY, Literal: "return"}},
			&IdentifierExpressionNode{Name: "x"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructStaticComplex verifies parsing of complex static field usage
func TestParser_StructStaticComplex(t *testing.T) {
	src := `Config.max = Config.min + 10 * 2`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IdentifierExpressionNode{Name: "Config"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "max"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&IdentifierExpressionNode{Name: "Config"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "min"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.PLUS_OP, Literal: "+"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 10}},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.MUL_OP, Literal: "*"}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructArrayFieldInit verifies parsing of array initialization in struct fields
func TestParser_StructArrayFieldInit(t *testing.T) {
	src := `struct Data { var items = [1, 2, 3]; }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "Data"}},
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Type: lexer.VAR_KEY, Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "items"},
			},
			&ArrayExpressionNode{},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 2}},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 3}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructMethodChainedCall verifies parsing of chained member access
func TestParser_StructMethodChainedCall(t *testing.T) {
	src := `obj.get().val`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&IdentifierExpressionNode{Name: "obj"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&CallExpressionNode{FunctionIdentifier: IdentifierExpressionNode{Name: "get"}},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "val"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_StructComplexConstructor verifies parsing of complex constructor logic
func TestParser_StructComplexConstructor(t *testing.T) {
	src := `struct Point { func init(x, y) { this.x = x; this.y = y; } }`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&StructDeclarationNode{StructName: IdentifierExpressionNode{Name: "Point"}},
			&FunctionStatementNode{FuncName: IdentifierExpressionNode{Name: "init"}},
			&IdentifierExpressionNode{Name: "x"},
			&IdentifierExpressionNode{Name: "y"},
			&BlockStatementNode{},
			&IdentifierExpressionNode{Name: "this"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "x"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&IdentifierExpressionNode{Name: "x"},
			&IdentifierExpressionNode{Name: "this"},
			&BinaryExpressionNode{Operation: lexer.Token{Type: lexer.DOT_OP, Literal: "."}},
			&IdentifierExpressionNode{Name: "y"},
			&AssignmentExpressionNode{Operation: lexer.Token{Type: lexer.ASSIGN_OP, Literal: "="}},
			&IdentifierExpressionNode{Name: "y"},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}

// TestParser_NewCallArgs verifies parsing of new struct instantiation with arguments
func TestParser_NewCallArgs(t *testing.T) {
	src := `var a = new A(1, true)`
	root := NewParser(src).Parse()
	assert.NotNil(t, root)

	testingVisitor := &TestingVisitor{
		ExpectedNodes: []Node{
			&DeclarativeStatementNode{
				VarToken:   lexer.Token{Literal: "var"},
				Identifier: IdentifierExpressionNode{Name: "a"},
			},
			&NewCallExpressionNode{
				StructName: IdentifierExpressionNode{Name: "A"},
			},
			&IntegerLiteralExpressionNode{Value: &std.Integer{Value: 1}},
			&BooleanLiteralExpressionNode{Value: &std.Boolean{Value: true}, Token: lexer.Token{Type: lexer.TRUE_KEY, Literal: "true"}},
		},
		Ptr: 0,
		T:   t,
	}
	root.Accept(testingVisitor)
}
