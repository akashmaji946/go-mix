/*
File    : go-mix/parser/test_visitor.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package parser - test_visitor.go
// This file defines the TestingVisitor type, which is a visitor implementation used for testing the AST traversal of the parser.
// The TestingVisitor asserts that the nodes visited during traversal match an expected sequence of nodes provided in advance.
// It uses the testify/assert package to perform assertions and will fail tests if the actual traversal does not match expectations.
package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestingVisitor is a visitor that asserts the expected nodes
// The expected nodes are given in the in-order (like) traversal order
type TestingVisitor struct {
	ExpectedNodes []Node     // List of expected nodes in traversal order
	Ptr           int        // Current position pointer in the expected nodes list
	T             *testing.T // Testing instance for assertions
}

// VisitRootNode visits the root node and recursively visits all statements
func (v *TestingVisitor) VisitRootNode(node RootNode) {
	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

// VisitExpressionNode visits a generic expression node (no-op implementation)
func (v *TestingVisitor) VisitExpressionNode(node ExpressionNode) {
}

// VisitStatementNode visits a generic statement node (no-op implementation)
func (v *TestingVisitor) VisitStatementNode(node StatementNode) {
}

// VisitIntegerLiteralExpressionNode visits an integer literal node and asserts its value matches expected
func (v *TestingVisitor) VisitIntegerLiteralExpressionNode(node IntegerLiteralExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*IntegerLiteralExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Value, exp.Value)
	}
	v.Ptr++
}

// VisitBooleanLiteralExpressionNode visits a boolean literal node and asserts its value and token type match expected
func (v *TestingVisitor) VisitBooleanLiteralExpressionNode(node BooleanLiteralExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*BooleanLiteralExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Value, exp.Value)
		assert.Equal(v.T, node.Token.Type, exp.Token.Type)
	}
	v.Ptr++
}

func (v *TestingVisitor) VisitCharLiteralExpressionNode(node CharLiteralExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*CharLiteralExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Value, exp.Value)
	}
	v.Ptr++
}

// VisitBinaryExpressionNode visits a binary expression node and asserts the operator matches expected
func (v *TestingVisitor) VisitBinaryExpressionNode(node BinaryExpressionNode) {
	node.Left.Accept(v)
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*BinaryExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Operation.Literal, exp.Operation.Literal)
	v.Ptr++

	node.Right.Accept(v)
}

// VisitUnaryExpressionNode visits a unary expression node and asserts the operator matches expected
func (v *TestingVisitor) VisitUnaryExpressionNode(node UnaryExpressionNode) {
	node.Right.Accept(v)
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*UnaryExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Operation.Literal, exp.Operation.Literal)
	assert.Equal(v.T, node.Operation.Type, exp.Operation.Type)
	v.Ptr++
}

// VisitParenthesizedExpressionNode visits a parenthesized expression node and asserts its type
func (v *TestingVisitor) VisitParenthesizedExpressionNode(node ParenthesizedExpressionNode) {
	node.Expr.Accept(v)
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ParenthesizedExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++
}

// VisitDeclarativeStatementNode visits a variable declaration node and asserts the keyword and identifier match expected
func (v *TestingVisitor) VisitDeclarativeStatementNode(node DeclarativeStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	currNode, ok := curr.(*DeclarativeStatementNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.VarToken.Literal, currNode.VarToken.Literal)
	assert.Equal(v.T, node.Identifier.Name, currNode.Identifier.Name)
	v.Ptr++

	node.Expr.Accept(v)
}

// VisitIdentifierExpressionNode visits an identifier node and asserts its name matches expected
func (v *TestingVisitor) VisitIdentifierExpressionNode(node IdentifierExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*IdentifierExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Name, exp.Name)
	}
	v.Ptr++
}

// VisitReturnStatementNode visits a return statement node and asserts its type
func (v *TestingVisitor) VisitReturnStatementNode(node ReturnStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ReturnStatementNode)
	assert.True(v.T, ok)
	v.Ptr++
}

// VisitBooleanExpressionNode visits a boolean comparison/logical expression node and asserts its type
func (v *TestingVisitor) VisitBooleanExpressionNode(node BooleanExpressionNode) {
	node.Left.Accept(v)
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*BooleanExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++
	node.Right.Accept(v)
}

// VisitBlockStatementNode visits a block statement node and recursively visits all statements within
func (v *TestingVisitor) VisitBlockStatementNode(node BlockStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*BlockStatementNode)
	assert.True(v.T, ok)
	v.Ptr++

	for _, stmt := range node.Statements {
		stmt.Accept(v)
	}
}

// VisitAssignmentExpressionNode visits an assignment expression node and asserts the operator and operands match expected
func (v *TestingVisitor) VisitAssignmentExpressionNode(node AssignmentExpressionNode) {
	node.Left.Accept(v)
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	exp, ok := curr.(*AssignmentExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Operation.Literal, exp.Operation.Literal)
	}
	v.Ptr++
	node.Right.Accept(v)
}

// VisitIfExpressionNode visits an if-else expression node and asserts the if token matches expected
func (v *TestingVisitor) VisitIfExpressionNode(node IfExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
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

// VisitStringLiteralExpressionNode visits a string literal node and asserts its value matches expected
func (v *TestingVisitor) VisitStringLiteralExpressionNode(node StringLiteralExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*StringLiteralExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Value, curr.(*StringLiteralExpressionNode).Value)
	}
	v.Ptr++
}

// VisitFloatLiteralExpressionNode visits a float literal node and asserts its value matches expected
func (v *TestingVisitor) VisitFloatLiteralExpressionNode(node FloatLiteralExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*FloatLiteralExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Value, curr.(*FloatLiteralExpressionNode).Value)
	}
	v.Ptr++
}

// VisitNilLiteralExpressionNode visits a nil literal node and asserts its value matches expected
func (v *TestingVisitor) VisitNilLiteralExpressionNode(node NilLiteralExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*NilLiteralExpressionNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Value, curr.(*NilLiteralExpressionNode).Value)
	}
	v.Ptr++
}

// VisitFunctionStatementNode visits a function declaration node and asserts the function name matches expected
func (v *TestingVisitor) VisitFunctionStatementNode(node FunctionStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
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

// VisitCallExpressionNode visits a function call expression node and asserts the function identifier matches expected
func (v *TestingVisitor) VisitCallExpressionNode(node CallExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
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

// VisitForLoopStatementNode visits a for loop node and recursively visits initializers, condition, updates, and body
func (v *TestingVisitor) VisitForLoopStatementNode(node ForLoopStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ForLoopStatementNode)
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

// VisitArrayExpressionNode visits an array literal node and recursively visits all elements
func (v *TestingVisitor) VisitArrayExpressionNode(node ArrayExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ArrayExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++

	for _, elem := range node.Elements {
		elem.Accept(v)
	}
}

// VisitIndexExpressionNode visits an array index expression node and visits the array and index
func (v *TestingVisitor) VisitIndexExpressionNode(node IndexExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*IndexExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++

	node.Left.Accept(v)
	node.Index.Accept(v)
}

// VisitWhileLoopStatementNode visits a while loop node and recursively visits conditions and body
func (v *TestingVisitor) VisitWhileLoopStatementNode(node WhileLoopStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*WhileLoopStatementNode)
	assert.True(v.T, ok)
	v.Ptr++

	// Visit conditions
	for _, cond := range node.Conditions {
		cond.Accept(v)
	}
	// Visit body
	node.Body.Accept(v)
}

// VisitSliceExpressionNode visits an array slice expression node and visits the array, start, and end indices
func (v *TestingVisitor) VisitSliceExpressionNode(node SliceExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*SliceExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++

	node.Left.Accept(v)
	if node.Start != nil {
		node.Start.Accept(v)
	}
	if node.End != nil {
		node.End.Accept(v)
	}
}

// VisitRangeExpressionNode visits a range expression node and visits the start and end expressions
func (v *TestingVisitor) VisitRangeExpressionNode(node RangeExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*RangeExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++

	node.Start.Accept(v)
	node.End.Accept(v)
}

// VisitForeachLoopStatementNode visits a foreach loop node and visits the iterator, iterable, and body
func (v *TestingVisitor) VisitForeachLoopStatementNode(node ForeachLoopStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ForeachLoopStatementNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.Iterator.Name, curr.(*ForeachLoopStatementNode).Iterator.Name)
	v.Ptr++

	node.Iterable.Accept(v)
	node.Body.Accept(v)
}

// VisitMapExpressionNode visits a map literal node and recursively visits all keys and values
func (v *TestingVisitor) VisitMapExpressionNode(node MapExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*MapExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++

	// Visit all key-value pairs
	for i := range node.Keys {
		node.Keys[i].Accept(v)
		node.Values[i].Accept(v)
	}
}

// VisitSetExpressionNode visits a set literal node and recursively visits all elements
func (v *TestingVisitor) VisitSetExpressionNode(node SetExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*SetExpressionNode)
	assert.True(v.T, ok)
	v.Ptr++

	// Visit all elements
	for _, elem := range node.Elements {
		elem.Accept(v)
	}
}

// VisitStructDeclarationNode visits a struct declaration node and asserts the struct name matches expected, then visits all methods
func (v *TestingVisitor) VisitStructDeclarationNode(node StructDeclarationNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*StructDeclarationNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.StructName.Literal(), curr.(*StructDeclarationNode).StructName.Literal())
	v.Ptr++

	for _, field := range node.Fields {
		field.Accept(v)
	}

	for _, method := range node.Methods {
		method.Accept(v)
	}
}

// VisitNewCallExpressionNode visits a struct instantiation node and asserts the struct name matches expected, then visits all arguments
func (v *TestingVisitor) VisitNewCallExpressionNode(node NewCallExpressionNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*NewCallExpressionNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.StructName.Literal(), curr.(*NewCallExpressionNode).StructName.Literal())
	v.Ptr++

	for _, arg := range node.Arguments {
		arg.Accept(v)
	}
}

func (v *TestingVisitor) VisitBreakStatementNode(node BreakStatementNode) {
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	v.Ptr++
	// No specific assertions for break node other than type which is handled by Accept
}

func (v *TestingVisitor) VisitContinueStatementNode(node ContinueStatementNode) {
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	v.Ptr++
	// No specific assertions for continue node other than type which is handled by Accept
}

// VisitImportStatementNode visits an import statement node
func (v *TestingVisitor) VisitImportStatementNode(node ImportStatementNode) {
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*ImportStatementNode)
	assert.True(v.T, ok)
	if ok {
		assert.Equal(v.T, node.Name, curr.(*ImportStatementNode).Name)
	}
	v.Ptr++
}

// VisitSwitchStatementNode visits a switch statement node and recursively visits expression, cases, and default
func (v *TestingVisitor) VisitSwitchStatementNode(node SwitchStatementNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*SwitchStatementNode)
	assert.True(v.T, ok)
	v.Ptr++

	// Visit the switch expression
	node.Expression.Accept(v)

	// Visit all case clauses
	for _, caseNode := range node.Cases {
		caseNode.Value.Accept(v)
		caseNode.Body.Accept(v)
	}

	// Visit default clause if present
	if node.Default != nil {
		node.Default.Body.Accept(v)
	}
}

// String returns the string representation of the visitor (empty string)
func (v *TestingVisitor) String() string {
	return ""
}

// VisitEnumDeclarationNode visits an enum declaration node and asserts the enum name matches expected, then visits all members
func (v *TestingVisitor) VisitEnumDeclarationNode(node EnumDeclarationNode) {
	// Check bounds before accessing ExpectedNodes
	if v.Ptr >= len(v.ExpectedNodes) {
		return
	}
	// assert on type
	curr := v.ExpectedNodes[v.Ptr]
	_, ok := curr.(*EnumDeclarationNode)
	assert.True(v.T, ok)
	assert.Equal(v.T, node.EnumName.Literal(), curr.(*EnumDeclarationNode).EnumName.Literal())
	v.Ptr++

	for _, member := range node.Members {
		member.Accept(v)
	}

}
