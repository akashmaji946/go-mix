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

// VisitIntegerLiteralExpressionNode visits the number literal expression node
func (p *PrintingVisitor) VisitIntegerLiteralExpressionNode(node parser.IntegerLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Integer Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))

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

// VisitIdentifierExpressionNode visits the identifier expression node
func (p *PrintingVisitor) VisitIdentifierExpressionNode(node parser.IdentifierExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Identifier Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))
}

// VisitReturnStatementNode visits the return statement node
func (p *PrintingVisitor) VisitReturnStatementNode(node parser.ReturnStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Return Statement Node [%s](%s => %d)\n", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitBooleanExpressionNode visits the boolean expression node
func (p *PrintingVisitor) VisitBooleanExpressionNode(node parser.BooleanExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Boolean Node [%s] (%s => %t)\n", node.Operation.Literal, node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitBlockStatementNode visits the block statement node
func (p *PrintingVisitor) VisitBlockStatementNode(node parser.BlockStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Block Statement Node (%s) => %d\n", node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	for _, stmt := range node.Statements {
		stmt.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitAssignmentExpressionNode visits the assignment expression node
func (p *PrintingVisitor) VisitAssignmentExpressionNode(node parser.AssignmentExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Assignment Node [%s] (%s => %d)\n", node.Operation.Literal, node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitIfExpressionNode visits the if expression node
func (p *PrintingVisitor) VisitIfExpressionNode(node parser.IfExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting If Expression Node [%s] (%s => %d)\n", node.IfToken.Literal, node.Condition.Literal(), node.ConditionValue))
	p.Indent += INDENT_SIZE
	node.Condition.Accept(p)
	if &node.ThenBlock != parser.EMPTY_BLOCK {
		node.ThenBlock.Accept(p)
	}
	if &node.ElseBlock != parser.EMPTY_BLOCK {
		node.ElseBlock.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitStringLiteral visits the string literal node
func (p *PrintingVisitor) VisitStringLiteralExpressionNode(node parser.StringLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting String Literal Node [\"%s\"] (\"%s\" => \"%s\")\n", node.Literal(), node.Literal(), node.Value))
}

// VisitFloatLiteral visits the float literal node
func (p *PrintingVisitor) VisitFloatLiteralExpressionNode(node parser.FloatLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Float Literal Node [%s] (%s => %f)\n", node.Literal(), node.Literal(), node.Value))
}

// VisitNilLiteral visits the nil literal node
func (p *PrintingVisitor) VisitNilLiteralExpressionNode(node parser.NilLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Nil Literal Node [%s] (%s => %s)\n", node.Literal(), node.Literal(), node.Value))
}

// VisitFunctionStatementNode visits the function statement node
func (p *PrintingVisitor) VisitFunctionStatementNode(node parser.FunctionStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Function Statement Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.FuncName.Accept(p)
	for _, param := range node.FuncParams {
		param.Accept(p)
	}
	node.FuncBody.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitCallExpressionNode visits the call expression node
func (p *PrintingVisitor) VisitCallExpressionNode(node parser.CallExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Call Expression Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.FunctionIdentifier.Accept(p)
	for _, arg := range node.Arguments {
		arg.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// String returns the string representation of the visitor
func (p *PrintingVisitor) String() string {
	return p.Buf.String()
}
