package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/objects"
)

// NodeVisitor: visitor pattern
// used to traverse the AST
type NodeVisitor interface {
	VisitExpressionNode(node ExpressionNode)
	VisitStatementNode(node StatementNode)
	VisitRootNode(node RootNode)

	VisitIntegerLiteralExpressionNode(node IntegerLiteralExpressionNode)
	VisitBooleanLiteralExpressionNode(node BooleanLiteralExpressionNode)
	// VisitIntegerLiteralExpressionNode(node IntegerLiteralExpressionNode)
	VisitFloatLiteralExpressionNode(node FloatLiteralExpressionNode)
	VisitStringLiteralExpressionNode(node StringLiteralExpressionNode)
	VisitNilLiteralExpressionNode(node NilLiteralExpressionNode)

	VisitBinaryExpressionNode(node BinaryExpressionNode)
	VisitUnaryExpressionNode(node UnaryExpressionNode)
	VisitBooleanExpressionNode(node BooleanExpressionNode)
	VisitParenthesizedExpressionNode(node ParenthesizedExpressionNode)
	VisitDeclarativeStatementNode(node DeclarativeStatementNode)
	VisitIdentifierExpressionNode(node IdentifierExpressionNode)
	VisitReturnStatementNode(node ReturnStatementNode)
	VisitBlockStatementNode(node BlockStatementNode)
	VisitAssignmentExpressionNode(node AssignmentExpressionNode)
	VisitIfExpressionNode(node IfExpressionNode)

	VisitFunctionStatementNode(node FunctionStatementNode)
	VisitCallExpressionNode(node CallExpressionNode)
}

// Node: base interface for all nodes of the AST
// Literal(): returns the string representation of the node
// Accept(): accepts a visitor
type Node interface {
	Literal() string
	Accept(visitor NodeVisitor)
}

// StatementNode: base interface for all statement nodes
// Node: every statement node is a node
// Statement(): returns the string representation of the node
type StatementNode interface {
	Node
	Statement()
}

// ExpressionNode: base interface for all expression nodes
// Node: every expression node is a node
// StatementNode: every expression is also a statement
// Expression(): returns the string representation of the node
type ExpressionNode interface {
	Node
	StatementNode
	Expression()
}

// RootNode: represents the root of the AST (the program node)
// Statements: list of statements in the program
// Value: value of the program (for expressions, or value of final expression below root)
type RootNode struct {
	Statements []StatementNode     // every line of code is a statement
	Value      objects.GoMixObject // (e.g. 2 + 3 * 4 + 2 => 16)
}

// RootNode.Literal(): string represenation of the root node's statement/expression
func (root *RootNode) Literal() string {
	res := ""
	for _, stmt := range root.Statements {
		res += stmt.Literal()
		res += ";"
	}
	return res
}

// RootNode.Accept(): accepts a visitor (eg PrintVisitor)
func (root *RootNode) Accept(visitor NodeVisitor) {
	visitor.VisitRootNode(*root)
}

// There can be many types of ExpressionNodes
// IntegerLiteralExpressionNode: represents a number literal
type IntegerLiteralExpressionNode struct {
	Token lexer.Token
	Value objects.GoMixObject
}

// NumberLiteralExpressionNode.Literal(): string represenation of the node
func (node *IntegerLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// NumberLiteralExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *IntegerLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitIntegerLiteralExpressionNode(*node)
}

// NumberLiteralExpressionNode.Statement(): every expression is also a statement
func (node *IntegerLiteralExpressionNode) Statement() {

}

// NumberLiteralExpressionNode.Expression(): every expression node is a node
func (node *IntegerLiteralExpressionNode) Expression() {

}

// FloatLiteralExpressionNode: represents a float literal
type FloatLiteralExpressionNode struct {
	Token lexer.Token
	Value objects.GoMixObject
}

// FloatLiteralExpressionNode.Literal(): string represenation of the node
func (node *FloatLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// FloatLiteralExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *FloatLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitFloatLiteralExpressionNode(*node)
}

// FloatLiteralExpressionNode.Statement(): every expression is also a statement
func (node *FloatLiteralExpressionNode) Statement() {

}

// FloatLiteralExpressionNode.Expression(): every expression node is a node
func (node *FloatLiteralExpressionNode) Expression() {

}

// BooleanLiteralExpressionNode: represents a boolean literal
type BooleanLiteralExpressionNode struct {
	Token lexer.Token
	Value objects.GoMixObject
}

// BooleanLiteralExpressionNode.Literal(): string represenation of the node
func (node *BooleanLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// BooleanLiteralExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BooleanLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitBooleanLiteralExpressionNode(*node)
}

// BooleanLiteralExpressionNode.Statement(): every expression is also a statement
func (node *BooleanLiteralExpressionNode) Statement() {

}

// BooleanLiteralExpressionNode.Expression(): every expression node is a node
func (node *BooleanLiteralExpressionNode) Expression() {

}

// BinaryExpressionNode: represents an expression with an operator
type BinaryExpressionNode struct {
	Operation lexer.Token
	Left      ExpressionNode
	Right     ExpressionNode
	Value     objects.GoMixObject
}

// BinaryExpressionNode.Literal(): string represenation of the node
func (node *BinaryExpressionNode) Literal() string {
	return node.Left.Literal() + node.Operation.Literal + node.Right.Literal()
}

// BinaryExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BinaryExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitBinaryExpressionNode(*node)
}

// BinaryExpressionNode.Statement(): every expression is also a statement
func (node *BinaryExpressionNode) Statement() {

}

// BinaryExpressionNode.Expression(): every expression node is a node
func (node *BinaryExpressionNode) Expression() {

}

// UnaryExpressionNode: represents an expression with an operator
type UnaryExpressionNode struct {
	Operation lexer.Token
	Right     ExpressionNode
	Value     objects.GoMixObject
}

// UnaryExpressionNode.Literal(): string represenation of the node
func (node *UnaryExpressionNode) Literal() string {
	return node.Operation.Literal + node.Right.Literal()
}

// UnaryExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *UnaryExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitUnaryExpressionNode(*node)
}

// UnaryExpressionNode.Statement(): every expression is also a statement
func (node *UnaryExpressionNode) Statement() {

}

// UnaryExpressionNode.Expression(): every expression node is a node
func (node *UnaryExpressionNode) Expression() {

}

// BooleanExpressionNode: represents an expression with a boolean operator
type BooleanExpressionNode struct {
	Operation lexer.Token
	Left      ExpressionNode
	Right     ExpressionNode
	Value     objects.GoMixObject
}

// ParenthesizedExpressionNode: represents an expression in parentheses
type ParenthesizedExpressionNode struct {
	Expr  ExpressionNode
	Value objects.GoMixObject
}

// ParenthesizedExpressionNode.Literal(): string represenation of the node
func (node *ParenthesizedExpressionNode) Literal() string {
	return "(" + node.Expr.Literal() + ")"
}

// ParenthesizedExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *ParenthesizedExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitParenthesizedExpressionNode(*node)
}

// ParenthesizedExpressionNode.Statement(): every expression is also a statement
func (node *ParenthesizedExpressionNode) Statement() {

}

// ParenthesizedExpressionNode.Expression(): every expression node is a node
func (node *ParenthesizedExpressionNode) Expression() {

}

// DeclarativeStatementNode: represents a declarative statement
type DeclarativeStatementNode struct {
	VarToken   lexer.Token
	Identifier lexer.Token
	Expr       ExpressionNode
	Value      objects.GoMixObject
}

// DeclarativeStatementNode.Literal(): string represenation of the node
func (node *DeclarativeStatementNode) Literal() string {
	return node.VarToken.Literal + " " + node.Identifier.Literal + " = " + node.Expr.Literal()
}

// DeclarativeStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *DeclarativeStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitDeclarativeStatementNode(*node)
}

// DeclarativeStatementNode.Statement(): every expression is also a statement
func (node *DeclarativeStatementNode) Statement() {

}

// DeclarativeStatementNode.Expression(): we should not have it
// why? because declarative statement is not an expression
// func (node *DeclarativeStatementNode) Expression() {
// }

// IdentifierExpressionNode: represents an identifier expression
type IdentifierExpressionNode struct {
	Name  string
	Value objects.GoMixObject
}

// IdentifierExpressionNode.Literal(): string represenation of the node
func (node *IdentifierExpressionNode) Literal() string {
	return node.Name
}

// IdentifierExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *IdentifierExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitIdentifierExpressionNode(*node)
}

// IdentifierExpressionNode.Statement(): every expression is also a statement
func (node *IdentifierExpressionNode) Statement() {

}

// IdentifierExpressionNode.Expression(): every expression node is a node
func (node *IdentifierExpressionNode) Expression() {

}

// ReturnStatementNode():
type ReturnStatementNode struct {
	ReturnToken lexer.Token
	Expr        ExpressionNode
	Value       objects.GoMixObject
}

// ReturnStatementNode.Literal(): string represenation of the node
func (node *ReturnStatementNode) Literal() string {
	return node.ReturnToken.Literal + " " + node.Expr.Literal()
}

// ReturnStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *ReturnStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitReturnStatementNode(*node)
}

// ReturnStatementNode.Statement(): every expression is also a statement
func (node *ReturnStatementNode) Statement() {

}

// ReturnStatementNode.Expression(): every expression node is a node
func (node *ReturnStatementNode) Expression() {

}

// BlockStatementNode: represents a block of statements
type BlockStatementNode struct {
	Statements []Node
	Value      objects.GoMixObject
}

// BlockStatementNode.Literal(): string represenation of the node
func (node *BlockStatementNode) Literal() string {
	str := "{"
	for _, stmt := range node.Statements {
		str += stmt.Literal()
		str += ";"
	}
	str += "}"
	return str
}

// BlockStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BlockStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitBlockStatementNode(*node)
}

// BlockStatementNode.Statement(): every expression is also a statement
func (node *BlockStatementNode) Statement() {

}

// BlockStatementNode.Expression(): every expression node is a node
func (node *BlockStatementNode) Expression() {

}

// BooleanExpressionNode.Literal(): string represenation of the node
func (node *BooleanExpressionNode) Literal() string {
	return node.Left.Literal() + node.Operation.Literal + node.Right.Literal()
}

// BooleanExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BooleanExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitBooleanExpressionNode(*node)
}

// BooleanExpressionNode.Statement(): every expression is also a statement
func (node *BooleanExpressionNode) Statement() {

}

// BooleanExpressionNode.Expression(): every expression node is a node
func (node *BooleanExpressionNode) Expression() {

}

// AssignmentExpressionNode: represents an assignment expression
type AssignmentExpressionNode struct {
	Operation lexer.Token
	Left      string
	Right     ExpressionNode
	Value     objects.GoMixObject
}

// AssignmentExpressionNode.Literal(): string represenation of the node
func (node *AssignmentExpressionNode) Literal() string {
	return node.Left + " " + node.Operation.Literal + " " + node.Right.Literal()
}

// AssignmentExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *AssignmentExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitAssignmentExpressionNode(*node)
}

// AssignmentExpressionNode.Statement(): every expression is also a statement
func (node *AssignmentExpressionNode) Statement() {

}

// AssignmentExpressionNode.Expression(): every expression node is a node
func (node *AssignmentExpressionNode) Expression() {

}

// if expression node
type IfExpressionNode struct {
	IfToken        lexer.Token
	Condition      ExpressionNode
	ConditionValue objects.GoMixObject
	ThenBlock      BlockStatementNode
	ElseBlock      BlockStatementNode
}

// IfExpressionNode.Literal(): string represenation of the node
func (node *IfExpressionNode) Literal() string {
	res := node.IfToken.Literal + " " + node.Condition.Literal() + " " + node.ThenBlock.Literal()
	if len(node.ElseBlock.Statements) > 0 {
		if len(node.ElseBlock.Statements) == 1 {
			if nestedIf, ok := node.ElseBlock.Statements[0].(*IfExpressionNode); ok {
				return res + " else " + nestedIf.Literal()
			}
		}
		return res + " else " + node.ElseBlock.Literal()
	}
	return res
}

// IfExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *IfExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitIfExpressionNode(*node)
}

// IfExpressionNode.Statement(): every expression is also a statement
func (node *IfExpressionNode) Statement() {

}

// IfExpressionNode.Expression(): every expression node is a node
func (node *IfExpressionNode) Expression() {

}

// empty block statement node
var EMPTY_BLOCK = &BlockStatementNode{
	Statements: []Node{},
}

// new if statement node
func NewIfStatement() *IfExpressionNode {
	return &IfExpressionNode{
		Condition:      &BinaryExpressionNode{},
		ThenBlock:      *EMPTY_BLOCK,
		ElseBlock:      *EMPTY_BLOCK,
		ConditionValue: &objects.Nil{},
		IfToken:        lexer.Token{},
	}
}

// string literal node
type StringLiteralExpressionNode struct {
	Token lexer.Token
	Value objects.GoMixObject
}

// StringLiteral.Literal(): string represenation of the node
func (node *StringLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// StringLiteral.Accept(): accepts a visitor (eg PrintVisitor)
func (node *StringLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitStringLiteralExpressionNode(*node)
}

// StringLiteral.Statement(): every expression is also a statement
func (node *StringLiteralExpressionNode) Statement() {

}

// StringLiteral.Expression(): every expression node is a node
func (node *StringLiteralExpressionNode) Expression() {

}

// null literal node
type NilLiteralExpressionNode struct {
	Token lexer.Token
	Value objects.GoMixObject
}

// NullLiteral.Literal(): string represenation of the node
func (node *NilLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// NullLiteral.Accept(): accepts a visitor (eg PrintVisitor)
func (node *NilLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitNilLiteralExpressionNode(*node)
}

// NullLiteral.Statement(): every expression is also a statement
func (node *NilLiteralExpressionNode) Statement() {

}

// NullLiteral.Expression(): every expression node is a node
func (node *NilLiteralExpressionNode) Expression() {

}

type FunctionStatementNode struct {
	FuncToken  lexer.Token
	FuncName   IdentifierExpressionNode
	FuncParams []*IdentifierExpressionNode
	FuncBody   BlockStatementNode
	Value      objects.GoMixObject
}

// FunctionStatementNode.Literal(): string represenation of the node
func (node *FunctionStatementNode) Literal() string {

	funcParams := ""
	for _, param := range node.FuncParams {
		funcParams += param.Literal() + ","
	}
	if len(funcParams) > 0 {
		funcParams = funcParams[:len(funcParams)-1]
	}
	return node.FuncToken.Literal + " " + node.FuncName.Literal() + " (" + funcParams + ") " + node.FuncBody.Literal()
}

// FunctionStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *FunctionStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitFunctionStatementNode(*node)
}

// FunctionStatementNode.Statement(): every expression is also a statement
func (node *FunctionStatementNode) Statement() {

}

// FunctionStatementNode.Expression(): every expression node is a node
func (node *FunctionStatementNode) Expression() {

}

func NewFunctionStatementNode() *FunctionStatementNode {
	return &FunctionStatementNode{
		FuncToken:  lexer.Token{Type: lexer.FUNC_KEY, Literal: "func"},
		FuncName:   IdentifierExpressionNode{Name: "foo", Value: &objects.Nil{}},
		FuncParams: make([]*IdentifierExpressionNode, 0),
		FuncBody:   *EMPTY_BLOCK,
	}
}

type CallExpressionNode struct {
	FunctionIdentifier IdentifierExpressionNode
	Arguments          []ExpressionNode
	Value              objects.GoMixObject
}

// CallExpressionNode.Literal(): string represenation of the node
func (node *CallExpressionNode) Literal() string {
	args := ""
	for _, arg := range node.Arguments {
		args += arg.Literal() + ","
	}
	if len(args) > 0 {
		args = args[:len(args)-1]
	}
	return node.FunctionIdentifier.Literal() + "(" + args + ")"
}

// CallExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *CallExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitCallExpressionNode(*node)
}

// CallExpressionNode.Statement(): every expression is also a statement
func (node *CallExpressionNode) Statement() {

}

// CallExpressionNode.Expression(): every expression node is a node
func (node *CallExpressionNode) Expression() {

}
