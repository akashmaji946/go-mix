/*
File    : go-mix/main/print_visitor.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package main

import (
	"bytes"
	"fmt"

	"github.com/akashmaji946/go-mix/parser"
)

const INDENT_SIZE = 4 // Number of spaces per indentation level

// indent writes the current indentation level to the buffer
func (p *PrintingVisitor) indent() {
	for i := 0; i < p.Indent; i++ {
		p.Buf.WriteString(" ")
	}
}

// PrintingVisitor is a visitor that prints AST nodes in a formatted tree structure
type PrintingVisitor struct {
	Indent int          // Current indentation level for formatting
	Buf    bytes.Buffer // Buffer to accumulate the formatted output
}

// VisitRootNode visits the root node and prints all statements with indentation
func (p *PrintingVisitor) VisitRootNode(node parser.RootNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %15s Node [%s] (%s => %v)\n", "Root",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	for _, stmt := range node.Statements {
		stmt.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitExpressionNode visits a generic expression node and prints its details
func (p *PrintingVisitor) VisitExpressionNode(node parser.ExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s)\n", "Expression",
		node.Literal(), node.Literal()))

}

// VisitStatementNode visits a generic statement node and prints its details
func (p *PrintingVisitor) VisitStatementNode(node parser.StatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s)\n", "Statement",
		node.Literal(), node.Literal()))

}

// VisitIntegerLiteralExpressionNode visits an integer literal node and prints its value
func (p *PrintingVisitor) VisitIntegerLiteralExpressionNode(node parser.IntegerLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Integer",
		node.Literal(), node.Literal(), node.Value.ToObject()))

}

// VisitBooleanLiteralExpressionNode visits a boolean literal node and prints its value
func (p *PrintingVisitor) VisitBooleanLiteralExpressionNode(node parser.BooleanLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Boolean",
		node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitBinaryExpressionNode visits a binary expression node and prints the operator with operands
func (p *PrintingVisitor) VisitBinaryExpressionNode(node parser.BinaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Binary",
		node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitUnaryExpressionNode visits a unary expression node and prints the operator with operand
func (p *PrintingVisitor) VisitUnaryExpressionNode(node parser.UnaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Unary",
		node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitParenthesizedExpressionNode visits a parenthesized expression node and prints the enclosed expression
func (p *PrintingVisitor) VisitParenthesizedExpressionNode(node parser.ParenthesizedExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %12s Node [%s] (%s => %v)\n", "Parenthesized",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitDeclarativeStatementNode visits a variable declaration node and prints the declaration details
func (p *PrintingVisitor) VisitDeclarativeStatementNode(node parser.DeclarativeStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s](%s => %v)\n", "Declaration",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitIdentifierExpressionNode visits an identifier node and prints its name
func (p *PrintingVisitor) VisitIdentifierExpressionNode(node parser.IdentifierExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Identifier",
		node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitReturnStatementNode visits a return statement node and prints the return expression
func (p *PrintingVisitor) VisitReturnStatementNode(node parser.ReturnStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s](%s => %v)\n", "Return",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Expr.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitBooleanExpressionNode visits a boolean comparison/logical expression node and prints the operator with operands
func (p *PrintingVisitor) VisitBooleanExpressionNode(node parser.BooleanExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Boolean",
		node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitBlockStatementNode visits a block statement node and prints all statements within
func (p *PrintingVisitor) VisitBlockStatementNode(node parser.BlockStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Block",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	for _, stmt := range node.Statements {
		stmt.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitAssignmentExpressionNode visits an assignment expression node and prints the assignment details
func (p *PrintingVisitor) VisitAssignmentExpressionNode(node parser.AssignmentExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Assignment",
		node.Operation.Literal, node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitIfExpressionNode visits an if-else expression node and prints the condition and branches
func (p *PrintingVisitor) VisitIfExpressionNode(node parser.IfExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "If",
		node.IfToken.Literal, node.Condition.Literal(), node.ConditionValue.ToObject()))
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

// VisitStringLiteralExpressionNode visits a string literal node and prints its value
func (p *PrintingVisitor) VisitStringLiteralExpressionNode(node parser.StringLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [\"%s\"] (\"%s\" => %v)\n", "String",
		node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitFloatLiteralExpressionNode visits a float literal node and prints its value
func (p *PrintingVisitor) VisitFloatLiteralExpressionNode(node parser.FloatLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %15s Node [%s] (%s => %v)\n", "Float",
		node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitNilLiteralExpressionNode visits a nil literal node and prints its value
func (p *PrintingVisitor) VisitNilLiteralExpressionNode(node parser.NilLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Nil",
		node.Literal(), node.Literal(), node.Value.ToObject()))
}

// VisitFunctionStatementNode visits a function declaration node and prints the function details
func (p *PrintingVisitor) VisitFunctionStatementNode(node parser.FunctionStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Function",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.FuncName.Accept(p)
	for _, param := range node.FuncParams {
		param.Accept(p)
	}
	node.FuncBody.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitCallExpressionNode visits a function call expression node and prints the call details
func (p *PrintingVisitor) VisitCallExpressionNode(node parser.CallExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Call",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.FunctionIdentifier.Accept(p)
	for _, arg := range node.Arguments {
		arg.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitForLoopStatementNode visits a for loop node and prints the loop components
func (p *PrintingVisitor) VisitForLoopStatementNode(node parser.ForLoopStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "For",
		node.Literal(), node.Literal(), node.Value.ToString()))
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

// VisitWhileLoopStatementNode visits a while loop node and prints the conditions and body
func (p *PrintingVisitor) VisitWhileLoopStatementNode(node parser.WhileLoopStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "While",
		node.Literal(), node.Literal(), node.Value.ToString()))
	p.Indent += INDENT_SIZE
	for _, cond := range node.Conditions {
		cond.Accept(p)
	}
	node.Body.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitArrayExpressionNode visits an array literal node and prints all elements
func (p *PrintingVisitor) VisitArrayExpressionNode(node parser.ArrayExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Array",
		node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	for _, elem := range node.Elements {
		elem.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitIndexExpressionNode visits an array index expression node and prints the array and index
func (p *PrintingVisitor) VisitIndexExpressionNode(node parser.IndexExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Index",
		node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Index.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitSliceExpressionNode visits an array slice expression node and prints the slice details
func (p *PrintingVisitor) VisitSliceExpressionNode(node parser.SliceExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Slice",
		node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	if node.Start != nil {
		node.Start.Accept(p)
	}
	if node.End != nil {
		node.End.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitRangeExpressionNode visits a range expression node and prints the range details
func (p *PrintingVisitor) VisitRangeExpressionNode(node parser.RangeExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Range",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.Start.Accept(p)
	node.End.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitForeachLoopStatementNode visits a foreach loop node and prints the loop details
func (p *PrintingVisitor) VisitForeachLoopStatementNode(node parser.ForeachLoopStatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Foreach",
		node.Literal(), node.Literal(), node.Value.ToString()))
	p.Indent += INDENT_SIZE
	node.Iterator.Accept(p)
	node.Iterable.Accept(p)
	node.Body.Accept(p)
	p.Indent -= INDENT_SIZE
}

// VisitMapExpressionNode visits a map literal node and prints all key-value pairs
func (p *PrintingVisitor) VisitMapExpressionNode(node parser.MapExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Map",
		node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	for i := range node.Keys {
		node.Keys[i].Accept(p)
		node.Values[i].Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitSetExpressionNode visits a set literal node and prints all elements
func (p *PrintingVisitor) VisitSetExpressionNode(node parser.SetExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %10s Node [%s] (%s => %v)\n", "Set",
		node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	for _, elem := range node.Elements {
		elem.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitStructDeclarationNode visits a struct declaration node and prints the struct details
func (p *PrintingVisitor) VisitStructDeclarationNode(node parser.StructDeclarationNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %15s Node [%s] (%s => %v)\n", "struct",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	for _, method := range node.Methods {
		method.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// VisitNewCallExpressionNode visits a struct instantiation node and prints the instantiation details
func (p *PrintingVisitor) VisitNewCallExpressionNode(node parser.NewCallExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting %15s Node [%s] (%s => %v)\n", "New",
		node.Literal(), node.Literal(), node.Value.ToObject()))
	p.Indent += INDENT_SIZE
	node.StructName.Accept(p)
	for _, arg := range node.Arguments {
		arg.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

// String returns the accumulated formatted output as a string
func (p *PrintingVisitor) String() string {
	return p.Buf.String()
}
