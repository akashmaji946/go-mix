package parser

import "github.com/akashmaji946/go-mix/lexer"

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
	VisitParenthesizedExpressionNode(node ParenthesizedExpressionNode)
	VisitDeclarativeStatementNode(node DeclarativeStatementNode)
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
//  why? because declarative statement is not an expression
// func (node *DeclarativeStatementNode) Expression() {
// }
