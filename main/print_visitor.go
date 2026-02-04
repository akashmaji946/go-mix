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
	p.Buf.WriteString(fmt.Sprintf("Visiting %15s Node [%s] (%s => %v)\n", "Root", node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	for _, stmt := range node.Statements {
		stmt.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitExpressionNode visits the expression node
func (p *PrintingVisitor) VisitExpressionNode(node parser.ExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s)\n", "Expression", node.Literal(), node.Literal()))

}

// VisitStatementNode visits the statement node
func (p *PrintingVisitor) VisitStatementNode(node parser.StatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s)\n", "Statement", node.Literal(), node.Literal()))

}

// VisitIntegerLiteralExpressionNode visits the number literal expression node
func (p *PrintingVisitor) VisitIntegerLiteralExpressionNode(node parser.IntegerLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Integer", node.Literal(), node.Literal(), node.Value.ToObject()))

}

// VisitBooleanLiteralExpressionNode visits the boolean literal expression node
func (p *PrintingVisitor) VisitBooleanLiteralExpressionNode(node parser.BooleanLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Boolean", node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitBinaryExpressionNode visits the binary expression node
func (p *PrintingVisitor) VisitBinaryExpressionNode(node parser.BinaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Binary", node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitUnaryExpressionNode visits the unary expression node
func (p *PrintingVisitor) VisitUnaryExpressionNode(node parser.UnaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Unary", node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitParenthesizedExpressionNode visits the parenthesized expression node
func (p *PrintingVisitor) VisitParenthesizedExpressionNode(node parser.ParenthesizedExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %12s Node [%s] (%s => %v)\n", "Parenthesized", node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitDeclarativeStatementNode visits the declarative statement node
func (p *PrintingVisitor) VisitDeclarativeStatementNode(node parser.DeclarativeStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s](%s => %v)\n", "Declaration", node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitIdentifierExpressionNode visits the identifier expression node
func (p *PrintingVisitor) VisitIdentifierExpressionNode(node parser.IdentifierExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Identifier", node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitReturnStatementNode visits the return statement node
func (p *PrintingVisitor) VisitReturnStatementNode(node parser.ReturnStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s](%s => %v)\n", "Return", node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitBooleanExpressionNode visits the boolean expression node
func (p *PrintingVisitor) VisitBooleanExpressionNode(node parser.BooleanExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Boolean", node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitBlockStatementNode visits the block statement node
func (p *PrintingVisitor) VisitBlockStatementNode(node parser.BlockStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Block", node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	for _, stmt := range node.Statements {
		stmt.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitAssignmentExpressionNode visits the assignment expression node
func (p *PrintingVisitor) VisitAssignmentExpressionNode(node parser.AssignmentExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Assignment", node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitIfExpressionNode visits the if expression node
func (p *PrintingVisitor) VisitIfExpressionNode(node parser.IfExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "If", node.IfToken.Literal, node.Condition.Literal(), node.ConditionValue.ToObject()))
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
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [\"%s\"] (\"%s\" => %v)\n", "String", node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitFloatLiteral visits the float literal node
func (p *PrintingVisitor) VisitFloatLiteralExpressionNode(node parser.FloatLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %15s Node [%s] (%s => %v)\n", "Float", node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitNilLiteral visits the nil literal node
func (p *PrintingVisitor) VisitNilLiteralExpressionNode(node parser.NilLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Nil", node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitFunctionStatementNode visits the function statement node
func (p *PrintingVisitor) VisitFunctionStatementNode(node parser.FunctionStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Function", node.Literal(), node.Literal(), node.Value.ToObject()))
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
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Call", node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.FunctionIdentifier.Accept(p)
	for _, arg := range node.Arguments {
		arg.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitForLoopStatementNode
func (p *PrintingVisitor) VisitForLoopStatementNode(node parser.ForLoopStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "For", node.Literal(), node.Literal(), node.Value.ToString()))
	p.Indent += INDENT_SIZE
	for _, init := range node.Initializers {
		init.Accept(p)
	}
	if node.Condition != nil {
		node.Condition.Accept(p)
	}
	for _, update := range node.Updates {
		update.Accept(p)
	}
	node.Body.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitWhileLoopStatementNode visits the while loop node
func (p *PrintingVisitor) VisitWhileLoopStatementNode(node parser.WhileLoopStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "While", node.Literal(), node.Literal(), node.Value.ToString()))
	p.Indent += INDENT_SIZE
	for _, cond := range node.Conditions {
		cond.Accept(p)
	}
	node.Body.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitArrayExpressionNode
func (p *PrintingVisitor) VisitArrayExpressionNode(node parser.ArrayExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Array", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	for _, elem := range node.Elements {
		elem.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitIndexExpressionNode
func (p *PrintingVisitor) VisitIndexExpressionNode(node parser.IndexExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Index", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Index.Accept(p)
	p.Indent -= INDENT_SIZE
}

// String returns the string representation of the visitor
func (p *PrintingVisitor) String() string {
	return p.Buf.String()
}
