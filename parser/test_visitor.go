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

// TestingVisitor.VisitIntegerLiteralExpressionNode visits the number literal expression node
func (v *TestingVisitor) VisitIntegerLiteralExpressionNode(node IntegerLiteralExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*IntegerLiteralExpressionNode)
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
	assert.Equal(v.T, node.Identifier.Name, currNode.Identifier.Name)
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

// TestingVisitor.VisitBooleanExpressionNode visits the boolean expression node
func (v *TestingVisitor) VisitBooleanExpressionNode(node BooleanExpressionNode) {
	node.Left.Accept(v)
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*BooleanExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++
	node.Right.Accept(v)
}

// TestingVisitor.VisitBlockStatementNode visits the block statement node
func (v *TestingVisitor) VisitBlockStatementNode(node BlockStatementNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*BlockStatementNode)
	assert.True(v.T, ok)
	v.Ptr++

	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

// TestingVisitor.VisitAssignmentExpressionNode visits the assignment expression node
func (v *TestingVisitor) VisitAssignmentExpressionNode(node AssignmentExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*AssignmentExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Operation.Literal, curr.(*AssignmentExpressionNode).Operation.Literal)
	assert.Equal(v.T, node.Left, curr.(*AssignmentExpressionNode).Left)
	assert.Equal(v.T, node.Right.Literal(), curr.(*AssignmentExpressionNode).Right.Literal())
	v.Ptr++
}

// TestingVisitor.VisitIfExpressionNode visits the if expression node
func (v *TestingVisitor) VisitIfExpressionNode(node IfExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*IfExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.IfToken.Literal, curr.(*IfExpressionNode).IfToken.Literal)
	v.Ptr++
	node.Condition.Accept(v)
	(&node.ThenBlock).Accept(v)
	// (&node.ElseBlock).Accept(v)
}

// TestingVisitor.VisitStringLiteral visits the string literal node
func (v *TestingVisitor) VisitStringLiteralExpressionNode(node StringLiteralExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*StringLiteralExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Value, curr.(*StringLiteralExpressionNode).Value)
	v.Ptr++
}

// TestingVisitor.VisitFloatLiteralExpressionNode visits the float literal node
func (v *TestingVisitor) VisitFloatLiteralExpressionNode(node FloatLiteralExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*FloatLiteralExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Value, curr.(*FloatLiteralExpressionNode).Value)
	v.Ptr++
}

// TestingVisitor.VisitNilLiteralExpressionNode visits the nil literal node
func (v *TestingVisitor) VisitNilLiteralExpressionNode(node NilLiteralExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*NilLiteralExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Value, curr.(*NilLiteralExpressionNode).Value)
	v.Ptr++
}

// TestingVisitor.VisitFunctionStatementNode visits the function statement node
func (v *TestingVisitor) VisitFunctionStatementNode(node FunctionStatementNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*FunctionStatementNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.FuncName.Literal(), curr.(*FunctionStatementNode).FuncName.Literal())
	v.Ptr++

	for _, param := range node.FuncParams {
		param.Accept(v)
	}
	node.FuncBody.Accept(v)
}

// TestingVisitor.VisitCallExpressionNode visits the call expression node
func (v *TestingVisitor) VisitCallExpressionNode(node CallExpressionNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*CallExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.FunctionIdentifier.Literal(), curr.(*CallExpressionNode).FunctionIdentifier.Literal())
	v.Ptr++

	for _, arg := range node.Arguments {
		arg.Accept(v)
	}
}

// TestingVisitor.VisitForLoopNode visits the for loop node
func (v *TestingVisitor) VisitForLoopNode(node ForLoopNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ForLoopNode)
	assert.True(v.T, ok)
	v.Ptr++

	// Visit initializers
	for _, init := range node.Initializers {
		init.Accept(v)
	}
	// Visit condition
	if node.Condition != nil {
		node.Condition.Accept(v)
	}
	// Visit updates
	for _, update := range node.Updates {
		update.Accept(v)
	}
	// Visit body
	node.Body.Accept(v)
}

// TestingVisitor.VisitWhileLoopNode visits the while loop node
func (v *TestingVisitor) VisitWhileLoopNode(node WhileLoopNode) {
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*WhileLoopNode)
	assert.True(v.T, ok)
	v.Ptr++

	// Visit conditions
	for _, cond := range node.Conditions {
		cond.Accept(v)
	}
	// Visit body
	node.Body.Accept(v)
}

// TestingVisitor.String returns the string representation of the visitor
func (v *TestingVisitor) String() string {
	return ""
}
