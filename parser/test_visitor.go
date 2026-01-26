package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestingVisitor is a visitor that asserts the expected nodes
type TestingVisitor struct {
	expectedNodes []Node
	ptr           int
	t             *testing.T
}

// TestingVisitor.VisitRootNode visits the root node
func (v *TestingVisitor) VisitRootNode(node RootNode) {
	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

// TestingVisitor.VisitExpressionNode visits the expression node
func (v *TestingVisitor) VisitExpressionNode(node ExpressionNode) {
}

// TestingVisitor.VisitStatementNode visits the statement node
func (v *TestingVisitor) VisitStatementNode(node StatementNode) {
}

// TestingVisitor.VisitNumberLiteralExpressionNode visits the number literal expression node
func (v *TestingVisitor) VisitNumberLiteralExpressionNode(node NumberLiteralExpressionNode) {
	// assert on type
	curr := v.expectedNodes[v.ptr]
	exp, ok := curr.(*NumberLiteralExpressionNode)
	assert.True(v.t, ok)
	assert.Equal(v.t, node.Value, exp.Value)
	v.ptr++
}

// TestingVisitor.VisitBooleanLiteralExpressionNode visits the boolean literal expression node
func (v *TestingVisitor) VisitBooleanLiteralExpressionNode(node BooleanLiteralExpressionNode) {
	// assert on type
	curr := v.expectedNodes[v.ptr]
	exp, ok := curr.(*BooleanLiteralExpressionNode)
	assert.True(v.t, ok)
	assert.Equal(v.t, node.Value, exp.Value)
	assert.Equal(v.t, node.Token.Type, exp.Token.Type)
	v.ptr++
}

// TestingVisitor.VisitBinaryExpressionNode visits the binary expression node
func (v *TestingVisitor) VisitBinaryExpressionNode(node BinaryExpressionNode) {
	node.Left.Accept(v)
	// assert on type
	curr := v.expectedNodes[v.ptr]
	exp, ok := curr.(*BinaryExpressionNode)
	assert.True(v.t, ok)
	assert.Equal(v.t, node.Operation.Literal, exp.Operation.Literal)
	v.ptr++

	node.Right.Accept(v)
}

// TestingVisitor.VisitUnaryExpressionNode visits the unary expression node
func (v *TestingVisitor) VisitUnaryExpressionNode(node UnaryExpressionNode) {
	node.Right.Accept(v)
	// assert on type
	curr := v.expectedNodes[v.ptr]
	exp, ok := curr.(*UnaryExpressionNode)
	assert.True(v.t, ok)
	assert.Equal(v.t, node.Operation.Literal, exp.Operation.Literal)
	assert.Equal(v.t, node.Operation.Type, exp.Operation.Type)
	v.ptr++
}

// TestingVisitor.VisitParenthesizedExpressionNode visits the parenthesized expression node
func (v *TestingVisitor) VisitParenthesizedExpressionNode(node ParenthesizedExpressionNode) {
	node.Expr.Accept(v)
	// assert on type
	curr := v.expectedNodes[v.ptr]
	_, ok := curr.(*ParenthesizedExpressionNode)
	assert.True(v.t, ok)
	v.ptr++
}

// TestingVisitor.String returns the string representation of the visitor
func (v *TestingVisitor) String() string {
	return ""
}
