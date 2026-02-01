package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
)

// NodeVisitor: visitor pattern
// used to traverse the AST
type NodeVisitor interface {
	VisitExpressionNode(node ExpressionNode)
	VisitStatementNode(node StatementNode)
	VisitRootNode(node RootNode)
	VisitNumberLiteralExpressionNode(node NumberLiteralExpressionNode)
	VisitBooleanLiteralExpressionNode(node BooleanLiteralExpressionNode)
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
	VisitStringLiteral(StringLiteral)
	VisitFunctionStatementNode(node FunctionStatementNode)
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
	Statements []StatementNode // every line of code is a statement
	Value      int             // (e.g. 2 + 3 * 4 + 2 => 16)
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
// NumberLiteralExpressionNode: represents a number literal
type NumberLiteralExpressionNode struct {
	Token lexer.Token
	Value int
}

// NumberLiteralExpressionNode.Literal(): string represenation of the node
func (node *NumberLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// NumberLiteralExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *NumberLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitNumberLiteralExpressionNode(*node)
}

// NumberLiteralExpressionNode.Statement(): every expression is also a statement
func (node *NumberLiteralExpressionNode) Statement() {

}

// NumberLiteralExpressionNode.Expression(): every expression node is a node
func (node *NumberLiteralExpressionNode) Expression() {

}

// BooleanLiteralExpressionNode: represents a boolean literal
type BooleanLiteralExpressionNode struct {
	Token lexer.Token
	Value bool
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
	Value     int
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
	Value     int
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
	Value     bool
}

// ParenthesizedExpressionNode: represents an expression in parentheses
type ParenthesizedExpressionNode struct {
	Expr  ExpressionNode
	Value int
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
	Value      int
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
	Value int
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
	Value       int
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
	Value      int
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
	Value     int
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
	ConditionValue int
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

var EMPTY_BLOCK = &BlockStatementNode{
	Statements: []Node{},
}

func NewIfStatement() *IfExpressionNode {
	return &IfExpressionNode{
		Condition:      &BinaryExpressionNode{},
		ThenBlock:      *EMPTY_BLOCK,
		ElseBlock:      *EMPTY_BLOCK,
		ConditionValue: 0,
		IfToken:        lexer.Token{},
	}
}

type StringLiteral struct {
	Token lexer.Token
	Value string
}

func (node *StringLiteral) Literal() string {
	return node.Value
}

func (node *StringLiteral) Accept(visitor NodeVisitor) {
	visitor.VisitStringLiteral(*node)
}

func (node *StringLiteral) Statement() {

}

func (node *StringLiteral) Expression() {

}

type FunctionStatementNode struct {
	FuncToken  lexer.Token
	FuncName   IdentifierExpressionNode
	FuncParams []*IdentifierExpressionNode
	FuncBody   BlockStatementNode
	Value      int
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
		FuncName:   IdentifierExpressionNode{Name: "foo", Value: 0},
		FuncParams: make([]*IdentifierExpressionNode, 0),
		FuncBody:   *EMPTY_BLOCK,
	}
}
