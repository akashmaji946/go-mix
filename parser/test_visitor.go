package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestingVisitor is a visitor that asserts the expected nodes
// The expected nodes are given in the in-order (like) traversal order
// The change in order of flat nodes (in expected nodes list) in the test
// should be failing comparison (in actual nodes list)
type TestingVisitor struct {
	ExpectedNodes []Node
	Ptr           int
	T             *testing.T
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
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*NumberLiteralExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Value, exp.Value)
	v.Ptr++
}

// TestingVisitor.VisitBooleanLiteralExpressionNode visits the boolean literal expression node
func (v *TestingVisitor) VisitBooleanLiteralExpressionNode(node BooleanLiteralExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*BooleanLiteralExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Value, exp.Value)
	assert.Equal(v.T, node.Token.Type, exp.Token.Type)
	v.Ptr++
}

// TestingVisitor.VisitBinaryExpressionNode visits the binary expression node
func (v *TestingVisitor) VisitBinaryExpressionNode(node BinaryExpressionNode) {
	node.Left.Accept(v)
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*BinaryExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Operation.Literal, exp.Operation.Literal)
	v.Ptr++

	node.Right.Accept(v)
}

// TestingVisitor.VisitUnaryExpressionNode visits the unary expression node
func (v *TestingVisitor) VisitUnaryExpressionNode(node UnaryExpressionNode) {
	node.Right.Accept(v)
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*UnaryExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Operation.Literal, exp.Operation.Literal)
	assert.Equal(v.T, node.Operation.Type, exp.Operation.Type)
	v.Ptr++
}

// TestingVisitor.VisitParenthesizedExpressionNode visits the parenthesized expression node
func (v *TestingVisitor) VisitParenthesizedExpressionNode(node ParenthesizedExpressionNode) {
	node.Expr.Accept(v)
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ParenthesizedExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++
}

// TestingVisitor.VisitDeclarativeStatementNode visits the declarative statement node
func (v *TestingVisitor) VisitDeclarativeStatementNode(node DeclarativeStatementNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	currNode, ok := curr.(*DeclarativeStatementNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.VarToken.Literal, currNode.VarToken.Literal)
	assert.Equal(v.T, node.Identifier.Literal, currNode.Identifier.Literal)
	v.Ptr++

	node.Expr.Accept(v)
}

// TestingVisitor.VisitIdentifierExpressionNode visits the identifier expression node
func (v *TestingVisitor) VisitIdentifierExpressionNode(node IdentifierExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*IdentifierExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Name, curr.(*IdentifierExpressionNode).Name)
	v.Ptr++
}

// TestingVisitor.VisitReturnStatementNode visits the return statement node
func (v *TestingVisitor) VisitReturnStatementNode(node ReturnStatementNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ReturnStatementNode)
	assert.True(v.T, ok)
	v.Ptr++
}

// TestingVisitor.String returns the string representation of the visitor
func (v *TestingVisitor) String() string {
	return ""
}
