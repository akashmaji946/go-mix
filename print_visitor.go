package main

import (
	"bytes"
	"fmt"

	"github.com/akashmaji946/go-mix/parser"
)

const INDENT_SIZE = 4

// indent indents the buffer by the indent size
func (p *PrintingVisitor) indent() {
	for i := 0; i < p.Indent; i++ {
		p.Buf.WriteString(" ")
	}
}

// PrintingVisitor is a visitor that prints the nodes
type PrintingVisitor struct {
	Indent int
	Buf    bytes.Buffer
}

// VisitRootNode visits the root node
func (p *PrintingVisitor) VisitRootNode(node parser.RootNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Root Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	for _, stmt := range node.Statements {
		stmt.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitExpressionNode visits the expression node
func (p *PrintingVisitor) VisitExpressionNode(node parser.ExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Expression Node (%s)\n", node.Literal()))

}

// VisitStatementNode visits the statement node
func (p *PrintingVisitor) VisitStatementNode(node parser.StatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Statement Node (%s)\n", node.Literal()))

}

// VisitNumberLiteralExpressionNode visits the number literal expression node
func (p *PrintingVisitor) VisitNumberLiteralExpressionNode(node parser.NumberLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Number Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))

}

// VisitBooleanLiteralExpressionNode visits the boolean literal expression node
func (p *PrintingVisitor) VisitBooleanLiteralExpressionNode(node parser.BooleanLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Boolean Node [%s] (%s => %t)\n", node.Literal(), node.Literal(), node.Value))
}

// VisitBinaryExpressionNode visits the binary expression node
func (p *PrintingVisitor) VisitBinaryExpressionNode(node parser.BinaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Binary Node [%s] (%s => %d)\n", node.Operation.Literal, node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitUnaryExpressionNode visits the unary expression node
func (p *PrintingVisitor) VisitUnaryExpressionNode(node parser.UnaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Unary Node [%s] (%s => %d)\n", node.Operation.Literal, node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitParenthesizedExpressionNode visits the parenthesized expression node
func (p *PrintingVisitor) VisitParenthesizedExpressionNode(node parser.ParenthesizedExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Parenthesized Node (%s)\n", node.Literal()))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitDeclarativeStatementNode visits the declarative statement node
func (p *PrintingVisitor) VisitDeclarativeStatementNode(node parser.DeclarativeStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Declarative Statement Node [%s](%s => %d)\n", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// String returns the string representation of the visitor
func (p *PrintingVisitor) String() string {
	return p.Buf.String()
}
